package url

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParser(t *testing.T) {
	const rawUrl = "https://foo.com/go"

	url, err := Parse(rawUrl)

	if err != nil {
		t.Fatalf("Parse(%q) err=%q, want nil", rawUrl, err)
	}

	if got, want := url.Scheme, "https"; got != want {
		t.Fatalf("Parse(%q).Scheme = %q; want = %q", rawUrl, got, want)
	}

	if got, want := url.Host, "foo.com"; got != want {
		t.Fatalf("Parse(%q).Host = %q; want = %q", rawUrl, got, want)
	}

	if got, want := url.Path, "go"; got != want {
		t.Fatalf("Parse(%q).Path = %q; want = %q", rawUrl, got, want)
	}
}

var tests = []struct {
	name     string
	in       string
	port     string
	hostName string
}{
	{"with port", "foo.com:80", "80", "foo.com"},
	{"with empty port", "foo.com:", "", "foo.com"},
	{"without port", "foo.com", "", "foo.com"},
	{"ip with port", "1.2.3.4:90", "90", "1.2.3.4"},
	{"ip without port", "1.2.3.4", "", "1.2.3.4"},
}

func TestParserPort(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s/%s", tt.name, tt.in), func(t *testing.T) {
			url := &URL{Host: tt.in}
			if got := url.Port(); tt.port != got {
				t.Errorf("want: %q; got: %q", tt.port, got)
			}
		})
	}
}

func TestParserHostName(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s/%s", tt.name, tt.in), func(t *testing.T) {
			url := &URL{Host: tt.in}
			if got := url.Hostname(); tt.hostName != got {
				t.Errorf("want: %q; got: %q", tt.hostName, got)
			}
		})
	}
}

func TestURLString(t *testing.T) {
	u := &url.URL{Scheme: "http", Host: "foo.com", Path: "go"}

	if want, got := "http://foo.com/go", u.String(); want != got {
		t.Fatalf("%#v.String\nwant = %q; got = %q", u, want, got)
	}
}
