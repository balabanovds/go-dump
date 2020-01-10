package ospf

import (
	"reflect"
	"testing"
)

func TestParseLS1(t *testing.T) {
	var tests = []struct {
		data []byte
		exp  LS1
	}{
		{
			data: p1LSData,
			exp:  p1LS1,
		},
		{
			data: p21LSData,
			exp:  p21LS1,
		},
	}

	for _, c := range tests {
		got := parseLS1(c.data)

		// as we test only LS1 link we have to nil LSA
		c.exp.LSA = LSA{}

		if !reflect.DeepEqual(c.exp, got) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}
}
