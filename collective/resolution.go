package collective

// Thing
type thing struct {
	Name string `json:"name"` // Uuid
	Cn   string `json:"cn"`
}

// Relation
type relation struct {
	Thing1 string `json:"thing1"`
	Thing2 string `json:"thing2"`
}

// ResolutionFunc - data store function
type resolutionFunc func(method, name string, body []byte, version int) ([]byte, error)
