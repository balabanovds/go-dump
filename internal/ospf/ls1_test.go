package ospf

import (
	"reflect"
	"testing"

	"github.com/balabanovds/go-dump/internal/util"
)

func TestParseLS1(t *testing.T) {
	var tests = []struct {
		hex string
		exp LS1
	}{
		{
			hex: p1LsaHex,
			exp: p1LS1,
		},
		{
			hex: p21LsaHex,
			exp: p21LS1,
		},
	}

	for _, c := range tests {
		data := util.HexString2Bytes(c.hex)[lsaHeaderLength:]
		got := parseLS1(data)

		// as we test only LS1 link we have to nil LSA
		c.exp.LSA = LSA{}

		if !reflect.DeepEqual(c.exp, got) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}
}
