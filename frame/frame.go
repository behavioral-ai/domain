package frame

import "net/http"

type Frame interface {
	Uri() string
	Headers() http.Header
	Load(b []byte) error
	Save() []byte
}
