package collective

import "fmt"

func ExampleAddNotification() {
	urn := Urn("agent:thing/path/resource")
	addNotification(urn, func(thing, event Urn) {
		fmt.Printf("Notify 1 -> [%v] [%v]\n", thing, event)
	})
	fn := getNotification(urn)

	fmt.Printf("test: GetNotification() -> [%v]\n", fn != nil)

	sendNotification(urn, EventChanged)

	addNotification(urn, func(thing, event Urn) {
		fmt.Printf("Notify 2 -> [%v] [%v]\n", thing, event)
	})

	sendNotification(urn, "event:updated")

	//Output:
	//test: GetNotification() -> [true]
	//Notify 1 -> [agent:thing/path/resource] [event:changed]
	//Notify 1 -> [agent:thing/path/resource] [event:updated]
	//Notify 2 -> [agent:thing/path/resource] [event:updated]

}
