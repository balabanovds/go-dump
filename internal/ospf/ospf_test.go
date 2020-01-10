package ospf

import (
	"reflect"
	"testing"

	"github.com/balabanovds/go-dump/internal/util"
)

func TestParseOSPF(t *testing.T) {
	var tests = []struct {
		hex string
		exp Packet
	}{
		{
			hex: p1OspfHex,
			exp: p1Ospf,
		},
		{
			hex: p2OspfHex,
			exp: p2Ospf,
		},
	}

	for _, c := range tests {
		data := util.HexString2Bytes(c.hex)
		got, err := ParsePacket(data)

		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(c.exp, got) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}
}
