package ospf

import (
	"reflect"
	"testing"
)

func TestParseLS2(t *testing.T) {
	var tests = []struct {
		data []byte
		exp  LS2
	}{
		{
			data: p22LSData,
			exp:  p22LS2,
		},
	}

	for _, c := range tests {
		got := parseLS2(c.data)

		// as we test only LS1 link we have to nil LSA
		c.exp.LSA = LSA{}

		if !reflect.DeepEqual(c.exp, got) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}
}
