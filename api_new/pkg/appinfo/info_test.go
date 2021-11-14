package appinfo

import (
	"reflect"
	"testing"
)

func TestGetValue(t *testing.T) {
	tests := []struct {
		name      string
		getter    func() interface{}
		expectVal interface{}
	}{
		{"PackageRootPath", func() interface{} {
			return PackageRootPath()
		}, "github.com/moneyforward/consumption_tax_api"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			v := test.getter()
			if !reflect.DeepEqual(test.expectVal, v) {
				t.Errorf("invalid value. <%+v> != <%+v>",
					test.expectVal, v)
			}
		})
	}
}
