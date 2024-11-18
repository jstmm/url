package url

import "testing"

func TestParse(t *testing.T) {
	const raw_url = "https://foo.com/go"

	u, err := Parse(raw_url)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", raw_url, err)
	}
	want := "https"
	if got := u.Scheme; got != want {
		t.Errorf("Parse(%q).Scheme = %q; want %q", raw_url, got, want)
	}

	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", raw_url, got, want)
	}

	if got, want := u.Path, "go"; got != want {
		t.Errorf("Parse(%q).Path = %q; want %q", raw_url, got, want)
	}
}

func TestURLPortWithPort(t *testing.T)      { testPort(t, "foo.com:80", "80") }
func TestURLPortWithEmptyPort(t *testing.T) { testPort(t, "foo.com:", "") }
func TestURLPortWithoutPort(t *testing.T)   { testPort(t, "foo.com", "") }
func TestURLPortIPWithPort(t *testing.T)    { testPort(t, "1.2.3.4:90", "90") }
func TestURLPortIPWithoutPort(t *testing.T) { testPort(t, "1.2.3.4:", "") }

func testPort(t *testing.T, in, wantPort string) {
	t.Helper()
	u := &URL{Host: in}
	if got := u.Port(); got != wantPort {
		t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
	}
}
