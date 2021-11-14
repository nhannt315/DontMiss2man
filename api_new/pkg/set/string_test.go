package set_test

import (
	"reflect"
	"testing"

	"github.com/nhannt315/real_estate_api/pkg/set"
)

func TestStringSet_Add(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name           string
		args           args
		wantAlreadySet bool
	}{
		{
			"val is not existing",
			args{val: "0"},
			false,
		},
		{
			"val is not existing",
			args{val: "1"},
			false,
		},
		{
			"val is existing",
			args{val: "1"},
			true,
		},
		{
			"val is existing",
			args{val: "0"},
			true,
		},
	}
	s := set.NewStringSet()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAlreadySet := s.Add(tt.args.val); gotAlreadySet != tt.wantAlreadySet {
				t.Errorf("Add() = %v, want %v", gotAlreadySet, tt.wantAlreadySet)
			}
		})
	}
}

func TestStringSet_Remove(t *testing.T) {
	tests := []struct {
		name       string
		initVals   []string
		removeVal  string
		expectRes  bool
		expectVals []string
	}{
		{"removed first(multi value sets)",
			[]string{"9", "5", "7", "3", "8"}, "9",
			false, []string{"5", "7", "3", "8"}},
		{"removed middle(multi value sets)",
			[]string{"9", "5", "7", "3", "8"}, "7",
			false, []string{"9", "5", "3", "8"}},
		{"removed last(multi value sets)",
			[]string{"9", "5", "7", "3", "8"}, "8",
			false, []string{"9", "5", "7", "3"}},
		{"removed only(one value sets)",
			[]string{"9"}, "9",
			false, []string{}},

		{"not removed(multi value sets)",
			[]string{"9", "5", "7", "3", "8"}, "4",
			true, []string{"9", "5", "7", "3", "8"}},
		{"not removed(one value sets)",
			[]string{"9"}, "4",
			true, []string{"9"}},
		{"not removed(empty sets)",
			[]string{}, "4",
			true, []string{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			s := set.NewStringSet()
			for _, v := range test.initVals {
				s.Add(v)
			}

			res := s.Remove(test.removeVal)
			if test.expectRes != res {
				t.Errorf("invalid remove result. %t != %t",
					test.expectRes, res)
			}
			if !reflect.DeepEqual(test.expectVals, s.Values()) {
				t.Errorf("invalid values. %+v != %+v",
					test.expectVals, s.Values())
			}
			if s.Contains(test.removeVal) {
				t.Errorf("removed value found.")
			}
		})
	}
}

func TestStringSet_Values(t *testing.T) {
	type fields struct {
		vals []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"val not existing",
			fields{
				vals: []string{"1"},
			},
			[]string{"1"},
		},
		{
			"val existing, single value",
			fields{
				vals: []string{"1"},
			},
			[]string{"1"},
		},
		{
			"val existing, multiple values",
			fields{
				vals: []string{"1", "2", "1", "2", "3"},
			},
			[]string{"1", "2", "3"},
		},
		{
			"nil",
			fields{
				vals: nil,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := set.NewStringSet()
			if tt.fields.vals != nil {
				s.AddAll(tt.fields.vals...)
			} else {
				s = nil
			}

			if got := s.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_SortedValues(t *testing.T) {
	type fields struct {
		vals []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"single value",
			fields{
				vals: []string{"1"},
			},
			[]string{"1"},
		},
		{
			"multiple values",
			fields{
				vals: []string{"3", "1", "2", "1", "2"},
			},
			[]string{"1", "2", "3"},
		},
		{
			"nil",
			fields{
				vals: nil,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := set.NewStringSet()
			if tt.fields.vals != nil {
				s.AddAll(tt.fields.vals...)
			} else {
				s = nil
			}

			if got := s.SortedValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_Contains(t *testing.T) {
	type fields struct {
		vals []string
	}
	tests := []struct {
		name    string
		fields  fields
		valWant map[string]bool
	}{
		{
			"single value",
			fields{
				vals: []string{"1"},
			},
			map[string]bool{
				"1":  true,
				"0":  false,
				"10": false,
			},
		},
		{
			"multiple values",
			fields{
				vals: []string{"1", "2", "1", "2", "3"},
			},
			map[string]bool{
				"1":  true,
				"2":  true,
				"3":  true,
				"0":  false,
				"10": false,
			},
		},
		{
			"nil",
			fields{
				vals: nil,
			},
			map[string]bool{
				"1":  false,
				"0":  false,
				"10": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := set.NewStringSet()
			if tt.fields.vals != nil {
				s.AddAll(tt.fields.vals...)
			} else {
				s = nil
			}

			for v, w := range tt.valWant {
				if got := s.Contains(v); !reflect.DeepEqual(got, w) {
					t.Errorf("Contains(%v) = %v, want %v", v, got, w)
				}
			}
		})
	}
}

func TestStringSet_ContainsAll(t *testing.T) {
	type fields struct {
		vals []string
	}
	type valsKey struct {
		vals [10]string
	}
	newVals := func(vs ...string) valsKey {
		var key valsKey
		for i, v := range vs {
			key.vals[i] = v
		}
		return key
	}
	valsToStrings := func(key valsKey) []string {
		s := make([]string, 0)
		for _, v := range key.vals {
			if v != "" {
				s = append(s, v)
			}
		}
		return s
	}
	newSet := func(vs ...string) *set.StringSet {
		s := set.NewStringSet()
		s.AddAll(vs...)
		return s
	}
	tests := []struct {
		name    string
		fields  fields
		valWant map[valsKey]struct {
			exists bool
			noVals *set.StringSet
		}
	}{
		{
			"single value",
			fields{
				vals: []string{"1"},
			},
			map[valsKey]struct {
				exists bool
				noVals *set.StringSet
			}{
				newVals("1"):           {true, nil},
				newVals("1", "2", "3"): {false, newSet("2", "3")},
				newVals("5", "3", "8"): {false, newSet("5", "3", "8")},
			},
		},
		{
			"multiple value",
			fields{
				vals: []string{"1", "2", "1", "2", "3"},
			},
			map[valsKey]struct {
				exists bool
				noVals *set.StringSet
			}{
				newVals("1", "2", "1", "2", "3"): {true, nil},
				newVals("3", "2", "1"):           {true, nil},
				newVals("1", "5", "7", "3"):      {false, newSet("5", "7")},
				newVals("8", "5", "4", "6"):      {false, newSet("8", "5", "4", "6")},
			},
		},
		{
			"nil",
			fields{
				vals: nil,
			},
			map[valsKey]struct {
				exists bool
				noVals *set.StringSet
			}{
				newVals():              {true, nil},
				newVals("1", "2", "3"): {false, newSet("1", "2", "3")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := set.NewStringSet()
			if tt.fields.vals != nil {
				s.AddAll(tt.fields.vals...)
			} else {
				s = nil
			}

			for key, w := range tt.valWant {
				got, gotNoVals := s.ContainsAll(valsToStrings(key)...)
				if got != w.exists {
					t.Errorf("ContainsAll(%+v) = %v, want %v", key, got, w.exists)
				}
				if gotNoVals.String() != w.noVals.String() {
					t.Errorf("ContainsAll(%+v) no values = %v, want %v", key, gotNoVals.String(), w.noVals.String())
				}
			}
		})
	}
}

func TestStringSet_IsEmpty(t *testing.T) {
	type fields struct {
		vals []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"single value",
			fields{
				vals: []string{"1"},
			},
			false,
		},
		{
			"multiple value",
			fields{
				vals: []string{"1", "2", "3", "2"},
			},
			false,
		},
		{
			"empty",
			fields{
				vals: []string{},
			},
			true,
		},
		{
			"nil",
			fields{
				vals: nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := set.NewStringSet()
			if tt.fields.vals != nil {
				s.AddAll(tt.fields.vals...)
			} else {
				s = nil
			}

			if got := s.IsEmpty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_Length(t *testing.T) {
	type fields struct {
		vals []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"single value",
			fields{
				vals: []string{"1"},
			},
			1,
		},
		{
			"multiple value",
			fields{
				vals: []string{"1", "2", "3", "2"},
			},
			3,
		},
		{
			"empty",
			fields{
				vals: []string{},
			},
			0,
		},
		{
			"nil",
			fields{
				vals: nil,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := set.NewStringSet()
			if tt.fields.vals != nil {
				s.AddAll(tt.fields.vals...)
			} else {
				s = nil
			}

			if got := s.Length(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_String(t *testing.T) {
	type fields struct {
		vals []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"single value",
			fields{
				vals: []string{"1"},
			},
			"[1]",
		},
		{
			"multiple value",
			fields{
				vals: []string{"1", "5", "3", "5", "4"},
			},
			"[1 5 3 4]",
		},
		{
			"empty",
			fields{
				vals: []string{},
			},
			"[]",
		},
		{
			"nil",
			fields{
				vals: nil,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := set.NewStringSet()
			if tt.fields.vals != nil {
				s.AddAll(tt.fields.vals...)
			} else {
				s = nil
			}

			if got := s.String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
