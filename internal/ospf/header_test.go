package ospf

import (
	"reflect"
	"testing"
)

func TestParseOspfHeader(t *testing.T) {
	tests := []struct {
		data []byte
		exp  Header
	}{
		{
			data: p1HeaderData,
			exp:  p1Header,
		},
	}

	for _, c := range tests {
		got := parseHeader(c.data)
		if !reflect.DeepEqual(got, c.exp) {
			t.Errorf("\nexp: %#v\ngot: %#v\n", c.exp, got)
		}
	}

}
