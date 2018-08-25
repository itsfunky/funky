package aws

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/aws/aws-lambda-go/lambdacontext"
)

// LambdaHandler represents the handler passed into the AWS Lambda SDK.
type LambdaHandler struct {
	http.Handler
}

// Invoke calls the handler.
func (h LambdaHandler) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	r, w, err := rpcToHTTP(payload)
	if err != nil {
		return nil, err
	}

	h.Handler.ServeHTTP(w, r)

	resp := w.Result()
	return ioutil.ReadAll(resp.Body)
}

// rpcToHTTP translates a request into appropriate HTTP request/response
// handlers.
func rpcToHTTP(in []byte) (*http.Request, *httptest.ResponseRecorder, error) {
	r, err := http.NewRequest("LAMBDA", lambdacontext.FunctionName, bytes.NewReader(in))
	if err != nil {
		return nil, nil, err
	}

	// TODO: Create custom recorder that implements http.ResponseWriter.
	w := httptest.NewRecorder()

	return r, w, nil
}
