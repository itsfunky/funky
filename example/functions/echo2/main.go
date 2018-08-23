package main

import (
	"net/http"

	"github.com/itsfunky/funky"
	"github.com/itsfunky/funky/example/lib"
)

func main() {
	funky.Handle(http.HandlerFunc(lib.MetadataHandler))
}
