package collective

import (
	"errors"
	"fmt"
	"net/http"
)

// fileResolution - is read only and returns "not found" on gets
func fileResolution(method, name, _ string, _ []byte, version int) ([]byte, error) {
	// file resolution is read only
	if method == http.MethodPut {
		return nil, nil
	}
	return nil, errors.New(fmt.Sprintf("error: not found, name \"%v\" version \"%v\"", name, version))
}
