package collective

import "net/http"

func httpResolution(method, name, author string, body []byte, version int) ([]byte, error) {
	req, _ := http.NewRequest(method, "", nil)
	_, err := http.DefaultClient.Do(req)
	return nil, err
}
