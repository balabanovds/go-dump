package ospf

import (
	"reflect"
	"testing"

	"github.com/balabanovds/go-dump/internal/util"
)

func TestParseOspfHeader(t *testing.T) {
	tests := []struct {
		hex string
		exp Header
	}{
		{
			hex: p1HeaderHex,
			exp: p1Header,
		},
	}

	for _, c := range tests {
		data := util.HexString2Bytes(c.hex)
		got := parseHeader(data)
		if !reflect.DeepEqual(got, c.exp) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}

}
