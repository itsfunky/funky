package serve

import (
	"context"
	"log"
	"net/http"
	"strings"
	"sync"
)

var (
	routes = map[string]route{}
	mutex  = &sync.Mutex{}
)

type route struct {
	path string
	runner
}

func findOrCreateRoute(ctx context.Context, path string) (route, error) {
	// Ensure two quick calls don't create two runners/routes on the same path.
	mutex.Lock()
	defer mutex.Unlock()

	if r, ok := routes[path]; ok {
		return r, nil
	}

	runner, err := createRunner(ctx, strings.TrimPrefix(path, "/"))
	if err != nil {
		return route{}, err
	}

	routes[path] = route{
		path:   path,
		runner: runner,
	}

	return routes[path], nil
}

// Serve creates an http server for locally invoking functions.
func Serve() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		route, err := findOrCreateRoute(ctx, path)
		if err != nil {
			log.Println("Failed to create funky runner: ", err)
			return
		}

		if err := route.runner(w, req); err != nil {
			log.Println(err)
		}
	})

	log.Println("Now serving on port 8080...")

	// TODO: Configurable listening address.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
