package collective

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
	"sync"
)

type content struct {
	body []byte
}

type contentT struct {
	m *sync.Map
}

func newContentCache() *contentT {
	c := new(contentT)
	c.m = new(sync.Map)
	return c
}

func (c *contentT) get(name string, version int) ([]byte, *messaging.Status) {
	if name == "" || version <= 0 {
		return nil, messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: name \"%v\" or version \"%v\" is empty", name, version)))
	}
	key := ResolutionKey{Name: name, Version: version}
	value, ok := c.m.Load(key)
	if !ok {
		return nil, messaging.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error: name \"%v\" version \"%v\"", name, version)))
	}
	if value1, ok1 := value.(content); ok1 {
		return value1.body, messaging.StatusOK()
	}
	return nil, messaging.StatusOK()
}

func (c *contentT) put(name string, body []byte, version int) *messaging.Status {
	if name == "" || body == nil || version <= 0 {
		return messaging.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error: name \"%v\", body \"%v\", or version \"%v\" is empty", name, body, version)))
	}
	c.m.Store(ResolutionKey{Name: name, Version: version}, content{body: body})
	return messaging.StatusOK()
}
