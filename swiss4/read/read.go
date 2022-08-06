package read

import (
	"io"
	"net/http"
	"os"
)

func FromWeb(url string) (io.ReadCloser, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func FromFile(filename string) (io.ReadCloser, error) {
	data, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ResourceHash(resourceLoc string) (io.ReadCloser, error) {
	f, err := os.Open(resourceLoc)
	if err != nil {
		return nil, err
	}
	return f, nil
}
