package collective

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"io/fs"
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

func loadContent(c *contentT, dir string) error {
	fileSystem := iox.DirFS(dir)
	return fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.Index(path, ".json") != -1 {
			buf, err1 := fs.ReadFile(fileSystem, path)
			if err1 != nil {
				return err1
			}
			k, err2 := parseResolutionKey(string(buf))
			if err2 != nil {
				return err2
			}
			err2 = c.put(k.Name, buf, k.Version)
			if err2 != nil {
				return err2
			}
		}
		return nil
	})
}
