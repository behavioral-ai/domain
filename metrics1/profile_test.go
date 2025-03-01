package metrics1

import (
	"encoding/json"
	"fmt"
)

func ExampleNewTrafficProfile() {
	p := NewTrafficProfile()
	buf, err := json.Marshal(&p)

	fmt.Printf("test: json.Marshal() -> [err:%v] [buf:%v]\n", err, string(buf))

	//Output:
	//fail
}

func ExampleNewTrafficProfile2() {
	p := NewTrafficProfile()
	buf, err := json.Marshal(&p)

	fmt.Printf("test: json.Marshal() -> [err:%v] [buf:%v]\n", err, string(buf))

	//Output:
	//fail
}

/*
	for i,day := range p.Week {
		switch i {
		case 0:

		}
		if day.
	}

*/
