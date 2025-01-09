package ingress1

import "net/http"

// How to do:
// identity
// create a frame
// select a frame
// load a frames attributes

type Frame interface {
	Uri() string
	Headers() http.Header
	Load(b []byte) error
	Save() []byte
}

// Query - how to select a frame
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

type frame struct {
	uri    string
	data   int
	desc   string
	output bool
}

func (f *frame) method() {

}

type Frame2 struct {
	T   any    // A struct type
	Uri string // Identity, not for selection but only for reference
}

func match(s *Frame, t *Frame) bool {
	return false

}
