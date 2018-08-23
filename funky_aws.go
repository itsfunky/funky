// +build aws

package funky

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/itsfunky/funky/aws"
)

// Handle sets up and starts the handler with AWS Lambda.
func Handle(handler http.Handler) {
	fmt.Println("Built with AWS type handler.")
	lh := aws.LambdaHandler{Handler: handler}
	lambda.StartHandler(lh)
}
