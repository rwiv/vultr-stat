package json

import (
	"bytes"
	"encoding/json"
	"io"
)

func ReadReader[T any](r io.ReadCloser, v *T) (*T, error) {
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func ReadReaderSlice[T any](r io.ReadCloser, v []T) ([]T, error) {
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func ToJsonReader(v any) (io.Reader, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(b), nil
}

func ToJsonBytes(v any) ([]byte, error) {
	r, err := ToJsonReader(v)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}
