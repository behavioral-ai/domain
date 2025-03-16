package content

import (
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"net/url"
	"strconv"
)

func httpResolution(method, nsName, author string, value []byte, version int) ([]byte, *messaging.Status) {
	if method == http.MethodGet {
		resp, status := get(nsName, version)
		if !status.OK() {
			return nil, status
		}
		if resp != nil {
		}
		return nil, status
	} else {
		resp, status := put(nsName, author, value, version)
		if !status.OK() {
			return nil, status
		}
		if resp != nil {
		}
		return nil, status
	}
}

func get(nsName string, version int) (*http.Response, *messaging.Status) {
	v := make(url.Values)
	v.Set(NsNameKey, nsName)
	v.Set(VersionKey, strconv.Itoa(version))
	http.NewRequest(http.MethodGet, "https://domain/res?"+v.Encode(), nil)
	//resp, status := http2.Do(req)
	return nil, messaging.StatusNotFound()
}

func put(nsName, author string, value []byte, version int) (*http.Response, *messaging.Status) {
	//req, _ := http.NewRequest(http.MethodPut, "", io.Nnil)
	//resp,status := http2.Do(req)
	return nil, messaging.StatusOK()
}
