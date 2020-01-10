package ospf

import (
	"reflect"
	"testing"
)

func TestParseLSA(t *testing.T) {
	var tests = []struct {
		data []byte
		exp  LSA
	}{
		{
			data: p1LSAData,
			exp:  p1LSA,
		},
		{
			data: p21LSAData,
			exp:  p21LSA,
		},
		{
			data: p22LSAData,
			exp:  p22LSA,
		},
	}

	for _, c := range tests {
		got := parseLSA(c.data)

		if !reflect.DeepEqual(c.exp, got) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}
}
