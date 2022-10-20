package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Header)
		writer.Write([]byte("ok"))
	})
	http.ListenAndServe(":8080", nil)
}
