package ospf

import (
	"reflect"
	"testing"

	"github.com/balabanovds/go-dump/internal/util"
)

func TestParseLSA(t *testing.T) {
	var tests = []struct {
		hex string
		exp LSA
	}{
		{
			hex: p1LsaHex,
			exp: p1LSA,
		},
		{
			hex: p21LsaHex,
			exp: p21LSA,
		},
		{
			hex: p22LsaHex,
			exp: p22LSA,
		},
	}

	for _, c := range tests {
		data := util.HexString2Bytes(c.hex)
		got := parseLSA(data)

		if !reflect.DeepEqual(c.exp, got) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}
}
