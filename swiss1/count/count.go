package count

import (
	"io"
	"strconv"
)

func FromReader(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "0", nil
	}
	n := len(data)
	return strconv.Itoa(n), nil
}
