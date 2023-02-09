package backend

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* A request body is deleted when is read for the first time,
thus it is important to rebuild it after reading it as a Log. */

func ReqBodyReader(req *http.Request) {
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
		fmt.Println("bodyBytes: ", string(bodyBytes[:]))
	}
	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}
