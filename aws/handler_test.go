package aws

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRequest() []byte {
	return []byte("{\"test\":123}")
}

func TestRPCToHTTPSuccess(t *testing.T) {
	lambdacontext.FunctionName = "foobar"
	r, _, err := rpcToHTTP(createRequest())
	assert.Equal(t, "LAMBDA", r.Method, "Request method should be 'LAMBDA'.")
	assert.Equal(t, "foobar", r.URL.String(), "Request URL should be 'foobar'.")
	assert.NoError(t, err, "Should not return an error.")
}

func TestRPCToHTTPFailure(t *testing.T) {
	lambdacontext.FunctionName = ":"
	_, _, err := rpcToHTTP(createRequest())
	assert.Error(t, err, "Should return an error for invalid url.")
}

func TestServiceInvokeSuccess(t *testing.T) {
	lambdacontext.FunctionName = "foobar"
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		in, err := ioutil.ReadAll(req.Body)
		require.NoError(t, err, "Reading from request body should not error.")
		assert.Equal(t, "{\"test\":123}", string(in), "Payload should match expected input.")

		_, err = io.WriteString(w, "Hello World")
		require.NoError(t, err, "Writing to response body should not error.")
	})
	svc := LambdaHandler{Handler: handler}

	out, err := svc.Invoke(context.Background(), createRequest())
	assert.NoError(t, err, "Invoking the service should not error.")
	assert.Equal(t, "Hello World", string(out), "Payload should match expected output.")
}

func TestServiceInvokeFailure(t *testing.T) {
	lambdacontext.FunctionName = ":"
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {})
	svc := LambdaHandler{Handler: handler}

	_, err := svc.Invoke(context.Background(), createRequest())
	assert.Error(t, err, "Should return an error for invalid url..")
}
