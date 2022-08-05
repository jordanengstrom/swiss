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
	// FIXME
	data, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
