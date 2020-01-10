package ospf

import (
	"reflect"
	"testing"

	"github.com/balabanovds/go-dump/internal/util"
)

func TestParseLSUpdate(t *testing.T) {
	var tests = []struct {
		hex string
		exp LSUpdate
	}{
		{
			hex: p1LSUpdateHex,
			exp: p1LSUpdate,
		},
		{
			hex: p2LSUpdateHex,
			exp: p2LSUpdate,
		},
	}

	for _, c := range tests {
		data := util.HexString2Bytes(c.hex)
		got := parseLSUpdatePacket(data)

		if !reflect.DeepEqual(c.exp, got) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}
}
