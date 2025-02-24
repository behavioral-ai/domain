package collective

// Thing
type thing struct {
	Name   string `json:"name"` // Uuid
	Cn     string `json:"cn"`
	Author string `json:"author"`
}

// Relation
type relation struct {
	Author string `json:"author"`
	Thing1 string `json:"thing1"`
	Thing2 string `json:"thing2"`
}
