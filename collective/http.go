package collective

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

func httpResolution(method, name, author string, body []byte, version int) ([]byte, *messaging.Status) {
	//req, _ := http.NewRequest(method, "", nil)
	//_, _ := http.DefaultClient.Do(req)
	if method == http.MethodPut {
		return nil, messaging.StatusOK()
	}
	return nil, messaging.StatusNotFound()
}
