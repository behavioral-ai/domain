package collective

// ResolutionFunc - data store function
type resolutionFunc func(name string, version int) ([]byte, error)

func fileResolution(name string, version int) ([]byte, error) {
	return nil, nil
}

func httpResolution(name string, version int) ([]byte, error) {
	return nil, nil
}
