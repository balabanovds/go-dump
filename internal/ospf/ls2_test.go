package ospf

import (
	"reflect"
	"testing"

	"github.com/balabanovds/go-dump/internal/util"
)

func TestParseLS2(t *testing.T) {
	var tests = []struct {
		hex string
		exp LS2
	}{
		{
			hex: p22LsaHex,
			exp: p22LS2,
		},
	}

	for _, c := range tests {
		data := util.HexString2Bytes(c.hex)[lsaHeaderLength:]
		got := parseLS2(data)

		// as we test only LS1 link we have to nil LSA
		c.exp.LSA = LSA{}

		if !reflect.DeepEqual(c.exp, got) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}
}
