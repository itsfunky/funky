package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/itsfunky/funky"
)

// MetadataHandler responds with the function's metadata.
func MetadataHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(funky.Metadata())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(b))
}
