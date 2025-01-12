package frame

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	UrnScheme  = "urn"
	FileScheme = "file"
	Wildcard   = "*"
)

type Map map[string]string

func NewMap() Map {
	return make(Map)
}

func (m Map) Match(k, v string) bool {
	if k == "" {
		return false
	}
	if v2, ok := m[k]; ok {
		return v2 == v || v2 == Wildcard
	}
	return false
}

type Frame interface {
	Uri() string
	Terms() Map
	Load(b []byte) error
	Save() []byte
}

func NewUri(scheme, module, opClass, frmClass string) string {
	uid, _ := uuid.NewUUID()
	return fmt.Sprintf("%v:%v:%v.%v.%v", scheme, module, opClass, frmClass, uid)
}

func NewUriTemplate(scheme, module, opClass, frmClass string) string {
	if module == "" {
		module = Wildcard
	}
	if opClass == "" {
		opClass = Wildcard
	}
	if frmClass == "" {
		frmClass = Wildcard
	}
	return fmt.Sprintf("%v:%v:%v.%v", scheme, module, opClass, frmClass)
}
