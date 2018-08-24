[![GoDoc](https://godoc.org/github.com/itsfunky/funky?status.svg)](https://godoc.org/github.com/itsfunky/funky)
[![Go Report Card](https://goreportcard.com/badge/github.com/itsfunky/funky)](https://goreportcard.com/report/github.com/itsfunky/funky)

# Funky

Funky lets you develop and build lambdas/functions for multiple cloud services, while allowing you to test and iterate quickly when developing locally.

## Example

```go
package main

import (
  "io"
  "net/http"
  
  "github.com/itsfunky/funky"
)

func main() {
  funky.Handle(http.HandlerFunc(func (w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "Hello World")
  }))
}
```

[Full Examples →](example)

## Roadmap

- [ ] Command Line
  - [ ] Build Functions
  - [ ] Delete Functions
  - [ ] Deploy Functions
    - [ ] Smart Deploy (only changed functions)
  - [ ] Invoke Functions (remote)
  - [ ] Invoke Functions (local)
  - [ ] List Functions
  - [ ] View Logs
    - [ ] Live Tail
  - [ ] Metrics (remote)
  - [ ] Metrics (local)
  - [ ] Serve HTTP (local)
    - [x] Basic Serving
    - [ ] HTTP Event Serving
- [ ] Local
  - [x] Basic RPC Service
  - [ ] Extended Options
- [ ] AWS
  - [x] Basic Invocations
  - [ ] Extended Invocation Support
- [ ] Google
  - [ ] JS Shim
  - [ ] Foreground Function Support
  - [ ] Background Function Support
