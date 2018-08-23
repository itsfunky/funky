package local

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

// Service represents the RPC service that will be invoke funky handlers.
type Service struct {
	handler http.Handler
}

// Invoke is an RPC handler that is called when there is an incoming request.
func (svc Service) Invoke(in *Request, out *Response) error {
	r, w, err := rpcToHTTP(in)
	if err != nil {
		return err
	}

	svc.handler.ServeHTTP(w, r)

	resp := w.Result()
	out.Payload, err = ioutil.ReadAll(resp.Body)
	return err
}

// rpcToHTTP translates the RPC request into appropriate HTTP request/response
// handlers.
func rpcToHTTP(in *Request) (*http.Request, *httptest.ResponseRecorder, error) {
	r, err := http.NewRequest("LOCAL", os.Getenv("FUNKY_FUNCTION_NAME"), bytes.NewReader(in.Payload))
	if err != nil {
		return nil, nil, err
	}

	// TODO: Create custom recorder that implements http.ResponseWriter.
	w := httptest.NewRecorder()

	return r, w, nil
}
