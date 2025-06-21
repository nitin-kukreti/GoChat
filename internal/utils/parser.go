package utils

import (
	"encoding/json"
	"io"
)

func BodyParser[T any](r io.Reader) (T, error) {
	var body T
	if err := json.NewDecoder(r).Decode(&body); err != nil {
		return body, err
	}
	return body, nil
}
