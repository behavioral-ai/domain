package collective

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

const (
	internalError           = "Internal Error"
	contextDeadlineExceeded = "context deadline exceeded"
)

type HttpExchange func(r *http.Request) (*http.Response, error)

var (
	authority []string
	exchange  HttpExchange
	client    *http.Client
)

// Initialize -
// TODO - add authentication and authorization
func Initialize(hosts []string, e HttpExchange) error {
	if len(hosts) == 0 {
		return errors.New("error: hosts configuration is empty")
	}
	copy(authority, hosts)
	exchange = e
	if e == nil {
		initClient()
		exchange = wrapper
	}
	return nil
}

// TODO : add timeout and failover
func wrapper(req *http.Request) (resp *http.Response, err error) {

	return do(req)
}

// do - process an HTTP request
func do(req *http.Request) (resp *http.Response, err error) {
	if req == nil {
		return &http.Response{StatusCode: http.StatusInternalServerError}, errors.New("invalid argument : request is nil")
	}
	resp, err = client.Do(req)
	if err != nil {
		if resp == nil {
			return serverErrorResponse(), err
		}
		// check for deadline exceeded error
		if req.Context() != nil && req.Context().Err() == context.DeadlineExceeded {
			resp.StatusCode = http.StatusGatewayTimeout
			err = errors.New(contextDeadlineExceeded)
		}
	}
	return
}

func serverErrorResponse() *http.Response {
	resp := new(http.Response)
	resp.StatusCode = http.StatusInternalServerError
	resp.Status = internalError
	return resp
}

func initClient() {
	t, ok := http.DefaultTransport.(*http.Transport)
	if ok {
		// Used clone instead of assignment due to presence of sync.Mutex fields
		var transport = t.Clone()
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		transport.MaxIdleConns = 200
		transport.MaxIdleConnsPerHost = 100
		client = &http.Client{Transport: transport, Timeout: time.Second * 5}
	} else {
		client = &http.Client{Transport: http.DefaultTransport, Timeout: time.Second * 5}
	}
}
