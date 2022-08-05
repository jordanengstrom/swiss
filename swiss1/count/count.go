package count

import (
	"io"
)

func FromReader(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	n := len(data)
	return n, nil
}
