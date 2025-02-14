package collective

import (
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/domain/testrsc"
)

func fileResolver(name string, _ int) ([]byte, error) {
	return iox.ReadFile(path(name))
}

func path(name string) string {
	switch name {
	case ResiliencyGradient:
		return testrsc.ResiliencyGradient
	case ResiliencySaturation:
		return testrsc.ResiliencySaturation
	}
	return name
}
