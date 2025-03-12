package testrsc

import (
	"embed"
	"github.com/behavioral-ai/core/iox"
)

//go:embed files
var f embed.FS

func init() {
	iox.Mount(f)
}
