package testrsc

import (
	"embed"
	"github.com/behavioral-ai/core/io"
)

//go:embed files
var f embed.FS

func init() {
	io.Mount(f)
}
