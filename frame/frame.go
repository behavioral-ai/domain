package frame

import (
	"fmt"
	"github.com/google/uuid"
)

type Header map[string]string

type Frame interface {
	Uri() string
	Headers() Header
	Load(b []byte) error
	Save() []byte
}

func NewUri(module string, opClass string, frmClass string) string {
	uid, _ := uuid.NewUUID()
	return fmt.Sprintf("%v:%v.%v.%v", module, opClass, frmClass, uid)
}
