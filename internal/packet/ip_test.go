package packet

import (
	"net"
	"reflect"
	"testing"
)

func TestParseIP(t *testing.T) {
	parsed := net.ParseIP("1.1.1.1").To4()
	expected := net.IP{1, 1, 1, 1}

	if !reflect.DeepEqual(parsed, expected) {
		t.Errorf("\nExp: %#v\nGot: %#v\n", expected, parsed)
	}
}
