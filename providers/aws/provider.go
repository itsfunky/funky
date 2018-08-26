package aws

import (
	"github.com/itsfunky/funky/providers"
)

func init() {
	providers.Register("aws", providers.Provider{
		Name:    "aws",
		Compile: "GOOS=linux GOARCH=amd64 go build -o main *.go",
	})
}
