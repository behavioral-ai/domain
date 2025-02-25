package collective

import (
	"errors"
	"github.com/behavioral-ai/core/messaging"
	"net/http"
)

type appender struct{}

// newAppender -
func newHttpAppender() Appender {
	a := new(appender)
	return a
}

// Thing - append a thing
func (a *appender) Thing(name, author string, related []string) *messaging.Status {
	return messaging.NewStatusError(http.StatusBadRequest, errors.New("error: not implemented"), "", "uri")
}

// Relation - append a relation
func (a *appender) Relation(name1, name2, author string) *messaging.Status {
	return messaging.NewStatusError(http.StatusBadRequest, errors.New("error: not implemented"), "", "uri")
}

// Frame - append a frame
func (a *appender) Frame(name, author string, contains []string, version int) *messaging.Status {
	return messaging.NewStatusError(http.StatusBadRequest, errors.New("error: not implemented"), "", "uri")
}

// Guidance - append guidance
func (a *appender) Guidance(name, author, text string, related []string) *messaging.Status {
	return messaging.NewStatusError(http.StatusBadRequest, errors.New("error: not implemented"), "", "uri")
}
