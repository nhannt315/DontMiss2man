package set

import (
	"fmt"
	"sort"
)

// StringSet represents a multiple sets of string
type StringSet struct {
	lookup map[string]struct{}
	vals   []string
}

// NewStringSet returns new instance of StringSet
func NewStringSet() *StringSet {
	return &StringSet{
		lookup: make(map[string]struct{}),
		vals:   make([]string, 0),
	}
}

// Add adds the given `val` to the set.
// Returning true if the `val` is already set before.
// This method is not thead-safe.
func (s *StringSet) Add(val string) (alreadySet bool) {
	if _, found := s.lookup[val]; !found { // lookup value in the set by val
		s.lookup[val] = struct{}{} // mark val is set to the set
		s.vals = append(s.vals, val)
		return
	}

	return true
}

// AddAll adds the given `vals` to the set.
// Returning true if the `vals` is already set before.
// This method is not thead-safe.
func (s *StringSet) AddAll(vals ...string) (alreadySet bool) {
	for _, v := range vals {
		alreadySet = s.Add(v) || alreadySet
	}
	return
}

// Remove remove the given `val` from the set.
// Returning true if the `val` is not found by the set.
// This method is not thead-safe.
func (s *StringSet) Remove(val string) (notValue bool) {
	if _, found := s.lookup[val]; !found {
		return true // 指定された値なし
	}
	// 指定された値の削除
	delete(s.lookup, val)
	for i, v := range s.vals {
		if v == val {
			// 削除
			afterVals := s.vals[(i + 1):]
			s.vals = append(s.vals[:i], afterVals...)
			break
		}
	}
	return false
}

// RemoveAll remove the given `vals` from the set.
// Returning true if the `vals` is not found by the set.
// This method is not thead-safe.
func (s *StringSet) RemoveAll(vals ...string) (notValue bool) {
	for _, v := range vals {
		notValue = s.Remove(v) || notValue
	}
	return
}

// Values returns all values in set.
// This method is not thead-safe.
func (s *StringSet) Values() []string {
	if s == nil {
		return nil
	}
	return s.vals
}

// SortedValues sort and returns all values in set.
func (s *StringSet) SortedValues() []string {
	if s == nil {
		return nil
	}
	sorted := make([]string, len(s.vals))
	copy(sorted, s.vals)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	return sorted
}

// Contains reports whether value is within the set.
func (s *StringSet) Contains(value string) (exists bool) {
	if s == nil {
		return false
	}
	_, exists = s.lookup[value]
	return
}

// ContainsAll reports whether all value is within the set.
func (s *StringSet) ContainsAll(vals ...string) (allExists bool, noExistsValueSet *StringSet) {
	if len(vals) == 0 {
		return true, nil
	}
	if s == nil {
		// nil なら全ての値なし
		noExistsValueSet = NewStringSet()
		noExistsValueSet.AddAll(vals...)
		return false, noExistsValueSet
	}

	allExists = true
	for _, v := range vals {
		if _, exists := s.lookup[v]; !exists {
			// 存在しなかった値の保持
			if noExistsValueSet == nil {
				noExistsValueSet = NewStringSet()
			}
			allExists = false
			noExistsValueSet.Add(v)
		}
	}
	return
}

// IsEmpty reports whether the set is empty or not.
func (s *StringSet) IsEmpty() bool {
	return (s == nil) || (len(s.vals) == 0)
}

// Length returns value cont
func (s *StringSet) Length() int {
	if s == nil {
		return 0
	}
	return len(s.vals)
}

// String to string
func (s *StringSet) String() string {
	if s == nil {
		return ""
	}
	return fmt.Sprintf("%+v", s.vals)
}
