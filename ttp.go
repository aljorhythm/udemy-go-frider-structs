package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

// HTTPMain func
func HTTPMain() {
	response, err := http.Get("http://www.google.com")

	if err != nil {
		os.Exit(-1)
	} else {
		defer response.Body.Close()
		// io.Copy(os.Stdout, response.Body)
		io.Copy(logWriter{}, response.Body)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
