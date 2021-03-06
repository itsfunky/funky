package serve

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/itsfunky/funky/internal"
	"github.com/itsfunky/funky/providers"
	"github.com/itsfunky/funky/providers/local"
)

type runner func(http.ResponseWriter, *http.Request) error

func getFreePort() (int, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	port := l.Addr().(*net.TCPAddr).Port

	if err := l.Close(); err != nil {
		return 0, err
	}

	return port, nil
}

func createRunnerCommand(ctx context.Context, function string, port int) *exec.Cmd {
	cmd := internal.CreateCommand(ctx, providers.Provider{Name: "local"}, function, internal.RunAction)
	cmd.Env = append(cmd.Env, fmt.Sprintf("FUNKY_SERVER_PORT=%d", port))
	cmd.Stdout = createLogWriter(function, os.Stdout)
	cmd.Stderr = createLogWriter(function, os.Stderr)

	return cmd
}

func startRunnerRPC(port int) (client *rpc.Client, err error) {
	// Ping the runner every ~100ms for ~5s.
	for i := 1; i <= 50; i++ {
		time.Sleep(100 * time.Millisecond)

		client, err = rpc.Dial("tcp", "localhost:"+strconv.Itoa(port))
		if err == nil {
			return client, nil
		}
	}

	return nil, err
}

func createRunner(ctx context.Context, path string) (runner, error) {
	if _, err := os.Stat(filepath.Join("functions", path, "main.go")); os.IsNotExist(err) {
		return func(w http.ResponseWriter, _ *http.Request) error {
			w.WriteHeader(http.StatusNotImplemented)
			return nil
		}, nil
	}

	port, err := getFreePort()
	if err != nil {
		return nil, err
	}

	cmd := createRunnerCommand(ctx, path, port)
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	client, err := startRunnerRPC(port)
	if err != nil {
		return nil, err
	}

	mutex := &sync.Mutex{}

	return func(w http.ResponseWriter, req *http.Request) error {
		// We use Mutex to simulate 1 request per container.
		mutex.Lock()
		defer mutex.Unlock()

		// TODO: Populate these accordingly
		rpcReq := &local.Request{}
		rpcResp := &local.Response{}

		if err := client.Call("Service.Invoke", rpcReq, rpcResp); err != nil {
			// TODO: Make a nice error, fall through for now.
			return err
		}

		if _, err := w.Write(rpcResp.Payload); err != nil {
			return err
		}

		return nil
	}, nil
}
