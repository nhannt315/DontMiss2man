package strings

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestMaskedString(t *testing.T) {

	maskedStr1 := MaskedStringWithString("abc")
	type tData struct {
		Name string
		Note MaskedString
	}
	data1 := tData{
		Name: "TestX",
		Note: maskedStr1,
	}

	tests := []struct {
		name   string
		str    func(t *testing.T) string
		expect string
	}{
		{"String()", func(t *testing.T) string {
			return maskedStr1.String()
		}, `******`},
		{"%s", func(t *testing.T) string {
			return fmt.Sprintf("x %s", maskedStr1)
		}, `x ******`},
		{"%v", func(t *testing.T) string {
			return fmt.Sprintf("x %v", maskedStr1)
		}, `x ******`},
		{"%+v", func(t *testing.T) string {
			return fmt.Sprintf("x %+v", maskedStr1)
		}, `x ******`},

		{"struct %v", func(t *testing.T) string {
			return fmt.Sprintf("x %v", data1)
		}, `x {TestX ******}`},
		{"struct %+v", func(t *testing.T) string {
			return fmt.Sprintf("x %+v", data1)
		}, `x {Name:TestX Note:******}`},

		{"*struct %v", func(t *testing.T) string {
			return fmt.Sprintf("x %v", &data1)
		}, `x &{TestX ******}`},
		{"*struct %+v", func(t *testing.T) string {
			return fmt.Sprintf("x %+v", &data1)
		}, `x &{Name:TestX Note:******}`},

		{"json", func(t *testing.T) string {
			var d tData
			j := `{"Name":"TestJson", "Note":"abcdef"}`
			err := json.Unmarshal([]byte(j), &d)
			if err != nil {
				t.Errorf("json Unmarshal: %+v", err)
			}
			if d.Name != "TestJson" {
				t.Errorf("invalid data: %s", d.Name)
			}
			if d.Note.PlaneString() != "abcdef" {
				t.Errorf("invalid data: %s", d.Note.PlaneString())
			}

			maskedJson, err := json.Marshal(d)
			if err != nil {
				t.Errorf("json Marshal: %+v", err)
			}
			return string(maskedJson)
		}, `{"Name":"TestJson","Note":"******"}`},

		{"yaml", func(t *testing.T) string {
			var d tData
			j := `name: TestYaml
note: abcdef`
			err := yaml.Unmarshal([]byte(j), &d)
			if err != nil {
				t.Errorf("yaml Unmarshal: %+v", err)
			}
			if d.Name != "TestYaml" {
				t.Errorf("invalid data: %s", d.Name)
			}
			if d.Note.PlaneString() != "abcdef" {
				t.Errorf("invalid data: %s", d.Note.PlaneString())
			}

			maskedYaml, err := yaml.Marshal(d)
			if err != nil {
				t.Errorf("yaml Marshal: %+v", err)
			}
			return string(maskedYaml)
		}, `name: TestYaml
note: '******'
`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			s := test.str(t)
			if test.expect != s {
				t.Errorf("invalid string. <%s> != <%s>",
					test.expect, s)
			}
		})
	}
}

func TestMaskedString_Value(t *testing.T) {
	planText := "abc"
	maskedStr := MaskedStringWithString(planText)

	tests := []struct {
		name    string
		m       MaskedString
		want    driver.Value
		wantErr bool
	}{
		{
			name: "正常系",
			m:    maskedStr,
			want: driver.Value(planText),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("MaskedString.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaskedString.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
