package ospf

import (
	"reflect"
	"testing"
)

func TestParseLSUpdate(t *testing.T) {
	var tests = []struct {
		data []byte
		exp  LSUpdate
	}{
		{
			data: p1LSUpdateData,
			exp:  p1LSUpdate,
		},
		{
			data: p2LSUpdateData,
			exp:  p2LSUpdate,
		},
	}

	for _, c := range tests {
		got := parseLSUpdatePacket(c.data)

		if !reflect.DeepEqual(c.exp, got) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}
}
