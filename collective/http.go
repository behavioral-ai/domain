package collective

import (
	"github.com/behavioral-ai/core/messaging"
)

func httpResolution(method, name, author string, body []byte, version int) ([]byte, *messaging.Status) {
	//req, _ := http.NewRequest(method, "", nil)
	//_, _ := http.DefaultClient.Do(req)
	return nil, messaging.StatusOK()
}
