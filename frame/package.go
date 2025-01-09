package frame

import "net/http"

const (
	PkgPath = "github/behavioral-ai/domain/frame"
)

// Query - how to find a frame
func Query(class string, h http.Header) (Frame, error) {
	return nil, nil
}

func Select(uri string, h http.Header) (Frame, error) {
	return nil, nil
}

func Update(uri string, h http.Header) error {
	return nil
}

func Insert(class string, f Frame) error {
	// Create the uri based on class + unique id
	return nil
}
