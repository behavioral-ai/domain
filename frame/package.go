package frame

import (
	"errors"
)

const (
	PkgPath = "github/behavioral-ai/domain/frame"
)

// Query - how to find a frame
func Query[T any](template string, terms Map) (t []T, err error) {
	return t, nil
}

func Get[T any](uri string) (t T, err error) {
	if uri == "" {
		return t, errors.New("error: Frame uri is empty")
	}

	return t, err
}

func Update(f Frame) error {
	return nil
}

func Insert(f Frame) error {
	// Create the uri based on class + unique id
	// Need a function for this
	return nil
}
