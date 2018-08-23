package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/itsfunky/funky"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	log.Println("Sleeping...")
	time.Sleep(5000 * time.Millisecond)
	log.Println("Hello!")
	io.WriteString(w, "Hello World")
}

func main() {
	funky.Handle(http.HandlerFunc(handler))
}
