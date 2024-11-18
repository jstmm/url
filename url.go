package url

import (
	"errors"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

// Parse parses raw_url into a URL structure.
func Parse(raw_url string) (*URL, error) {
	i := strings.Index(raw_url, "://")
	if i < 0 {
		return nil, errors.New("missing scheme")
	}
	scheme, host, path := raw_url[:i], raw_url[i+3:], ""
	if i = strings.Index(host, "/"); i >= 0 {
		host, path = host[:i], host[i+1:]
	}
	return &URL{scheme, host, path}, nil
}

func (u *URL) Port() string {
	i := strings.Index(u.Host, ":")
	if i < 0 {
		return ""
	} else {
		return u.Host[i+1:]
	}
}
