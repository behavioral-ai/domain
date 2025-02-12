package collective

import (
	"github.com/behavioral-ai/domain/common"
)

type host struct {
	EntryId int           `json:"entry-id"`
	Created Timestamp     `json:"created-ts"`
	Origin  common.Origin `json:"origin"`
}

func _ExampleStore() {

	//Output:
	//fail
}
