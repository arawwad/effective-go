package url

import (
	"errors"
	"fmt"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

// Parse parses raw url into a URL struct
func Parse(url string) (*URL, error) {

	i := strings.Index(url, "://")

	if i < 0 {
		return nil, errors.New("missing scheme")
	}

	scheme, rest := url[:i], strings.Split(url[i+3:], "/")

	var host, path string

	host = rest[0]
	if len(rest) > 0 {
		path = rest[1]
	}

	return &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}, nil
}

func (url *URL) Port() string {
	index := strings.Index(url.Host, ":")

	if index < 0 {
		return ""
	}
	return url.Host[index+1:]
}
func (url *URL) Hostname() string {
	index := strings.Index(url.Host, ":")

	if index < 0 {
		return url.Host
	}
	return url.Host[:index]
}

func (url *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", url.Scheme, url.Host, url.Path)
}
