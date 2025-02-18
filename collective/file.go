package collective

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// fileResolution - is read only and returns "not found" on gets
func fileResolution(method, name, _ string, _ []byte, version int) ([]byte, error) {
	// file resolution is read only
	if method == http.MethodPut {
		return nil, nil
	}
	return nil, errors.New(fmt.Sprintf("error: not found, name \"%v\" version \"%v\"", name, version))
}

func parseResolutionKey(s string) (ResolutionKey, error) {
	token := "resolution-key"
	t1 := strings.Index(s, token)
	if t1 < 0 {
		return ResolutionKey{}, errors.New("error: resolution-key object not found")
	}
	t1 += len(token)

	token = "name"
	t2 := t1 + 2
	t2 += strings.Index(s[t1:], token)
	if t2 < t1+2 {
		return ResolutionKey{}, errors.New("error: name value not found")
	}
	t2 += len(token)

	token = "version"
	t3 := t2 + 2
	t3 += strings.Index(s[t2:], token)
	if t3 < t2+2 {
		return ResolutionKey{}, errors.New("error: version value not found")
	}
	t3 += len(token)

	// parse name
	k := ResolutionKey{}
	tokens := strings.Split(s[t2:], "\"")
	k.Name = tokens[1]

	// parse version
	tokens = strings.Split(s[t3:], "\r")
	i, err := strconv.Atoi(strings.Trim(tokens[0], " "))
	if err != nil {
		return ResolutionKey{}, err
	}
	k.Version = i
	return k, nil
}

func load(dir string) error {
	//fileSystem := os.DirFS(dirroot)
	//fs.ReadDir(dir)

	return nil
}
