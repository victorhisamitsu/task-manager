package httphandler

import (
	"encoding/json"
	"io"
)

func ReadBody[T any](b io.ReadCloser, bodyRequest *T) error {
	body, err := io.ReadAll(b)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, bodyRequest)
	if err != nil {
		return err
	}
	return nil
}
