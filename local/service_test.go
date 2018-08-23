package local

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRequest() *Request {
	return &Request{
		Payload: []byte("test"),
	}
}

func TestRPCToHTTPSuccess(t *testing.T) {
	err := os.Setenv("FUNKY_FUNCTION_NAME", "foobar")
	require.NoError(t, err, "Setting the environment variable should not error.")

	r, _, err := rpcToHTTP(createRequest())
	assert.Equal(t, "LOCAL", r.Method, "Request method should be 'LOCAL'.")
	assert.Equal(t, "foobar", r.URL.String(), "Request URL should be 'foobar'.")
	assert.NoError(t, err, "Should not return an error.")
}

func TestRPCToHTTPFailure(t *testing.T) {
	err := os.Setenv("FUNKY_FUNCTION_NAME", ":")
	require.NoError(t, err, "Setting the environment variable should not error.")

	_, _, err = rpcToHTTP(createRequest())
	assert.Error(t, err, "Should return an error for invalid url.")
}

func TestServiceInvokeSuccess(t *testing.T) {
	err := os.Setenv("FUNKY_FUNCTION_NAME", "foobar")
	require.NoError(t, err, "Setting the environment variable should not error.")

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		in, err := ioutil.ReadAll(req.Body)
		require.NoError(t, err, "Reading request body should not error.")
		assert.Equal(t, "test", string(in), "Payload should match expected input.")

		_, err = io.WriteString(w, "Hello World")
		require.NoError(t, err, "Writing to response body should not error.")
	})
	svc := Service{handler: handler}

	resp := Response{}
	err = svc.Invoke(createRequest(), &resp)
	assert.NoError(t, err, "Invoking the service should not error.")
	assert.Equal(t, "Hello World", string(resp.Payload), "Payload should match expected output.")
}

func TestServiceInvokeFailure(t *testing.T) {
	err := os.Setenv("FUNKY_FUNCTION_NAME", ":")
	require.NoError(t, err, "Setting the environment variable should not error.")

	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {})
	svc := Service{handler: handler}

	err = svc.Invoke(createRequest(), &Response{})
	assert.Error(t, err, "Should return an error for invalid url..")
}
