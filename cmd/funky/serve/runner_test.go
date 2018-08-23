package serve

import (
	"context"
	"net"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFreePort(t *testing.T) {
	port, err := getFreePort()
	assert.NotEqual(t, 0, port, "Port should not equal 0.")
	assert.NoError(t, err, "Should not return an error.")
}

func TestCreateRunnerCommand(t *testing.T) {
	cmd := createRunnerCommand(context.Background(), "foobar", 12345)
	assert.Equal(t, filepath.Join("functions", "foobar"), cmd.Dir, "Dir should point to the function's folder.")

	gotFunkyEnvs := map[string]string{}
	expectedFunkyEnvs := map[string]string{
		"FUNKY_FUNCTION_NAME": "foobar",
		"FUNKY_SERVER_PORT":   "12345",
	}

	for _, v := range cmd.Env {
		if strings.HasPrefix(v, "FUNKY_") {
			kv := strings.SplitN(v, "=", 2)
			gotFunkyEnvs[kv[0]] = kv[1]
		}
	}

	assert.Equal(t, expectedFunkyEnvs, gotFunkyEnvs, "Env should include both Environment variables and Funky variables.")
}

func TestStartRunnerRPCSuccess(t *testing.T) {
	port, err := getFreePort()
	require.NoError(t, err, "Getting a free port should not error.")

	lis, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	require.NoError(t, err, "Starting an RPC server should not error.")

	go func(t *testing.T) {
		_, err = lis.Accept()
		require.NoError(t, err, "Listening should not error.")
	}(t)

	_, err = startRunnerRPC(port)
	assert.NoError(t, err, "Runner should be able to dial into the server.")
}

func TestStartRunnerRPCFailure(t *testing.T) {
	port, err := getFreePort()
	require.NoError(t, err, "Getting a free port should not error.")

	_, err = startRunnerRPC(port)
	assert.Error(t, err, "Error should be returned client could not dial into port.")
}
