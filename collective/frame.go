package collective

type join struct {
	Thing1 Urn
	Thing2 Urn
}

type Frame struct {
	Id      thing  `json:"id"`
	Things  []join `json:"things"`
	Version int    `json:"version"` // System generated auto incrementing version
}
