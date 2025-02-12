package collective

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/domain/common"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

type entry struct {
	Name      Urn `json:"name"` // Uuid
	CreatedTS string
	//Partition string
	Content []byte
	Version int
}

type host struct {
	EntryId int           `json:"entry-id"`
	Created Timestamp     `json:"created-ts"`
	Origin  common.Origin `json:"origin"`
}

var (
	storeVersion = 0
	sm           sync.Mutex
	store        []entry
)

func newVersion(fragment string) string {
	version++
	v := strconv.Itoa(version)
	if fragment == "" {
		return v
	}
	return fragment + ":" + v
}

func partition(name Urn) string {
	i := strings.Index(string(name), "#")
	if i >= 0 {
		return string(name)[i+1:]
	}
	return ""
}

func put(name Urn, body []byte) *aspect.Status {
	sm.Lock()
	defer sm.Unlock()
	version++
	store = append(store, entry{Name: name, Content: body, Version: version})
	return aspect.StatusOK()
}

func get(name Urn, values url.Values, fragment string) (body []byte, status *aspect.Status) {

	sm.Lock()
	defer sm.Unlock()

	return nil, aspect.StatusOK()
}
