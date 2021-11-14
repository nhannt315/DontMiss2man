package logs

import "time"

// Field is the addition log information.
// it contains two method key and value, of essential to configure logger's field
type Field interface {
	Key() string
	Value() interface{}
}

// FindField フィールド一覧からキーの一致するフィールドを検索/取得する。
func FindField(fds []Field, key string) (fld Field, ok bool) {
	for _, f := range fds {
		if f.Key() == key {
			return f, true
		}
	}
	return nil, false
}

const (
	// LevelFieldKey is level field's key name
	LevelFieldKey = "severity"
	// TimeFieldKey is timestamp field's key name
	TimeFieldKey = "timestamp"
	// StackFieldKey is stack field's key name
	StackFieldKey = "stack"
)

// int64Field is used to create field in int64 type
type int64Field struct {
	key   string
	value int64
}

// NewInt64Field creates int64 type field
func NewInt64Field(key string, value int64) Field {
	return &int64Field{key: key, value: value}
}

// Key returns int64 type field's key
func (fld *int64Field) Key() string {
	return fld.key
}

// Value returns int64 type field's value
func (fld *int64Field) Value() interface{} {
	return fld.value
}

// uint64Field is used to create field in uint64 type
type uint64Field struct {
	key   string
	value uint64
}

// NewUint64Field creates uint64 type field
func NewUint64Field(key string, value uint64) Field {
	return &uint64Field{key: key, value: value}
}

// Key returns uint64 type field's key
func (fld *uint64Field) Key() string {
	return fld.key
}

// Value returns uint64 type field's value
func (fld *uint64Field) Value() interface{} {
	return fld.value
}

// stringField is used to create field in string type
type stringField struct {
	key   string
	value string
}

// NewStringField creates string type field
func NewStringField(key string, value string) Field {
	return &stringField{key: key, value: value}
}

// Key returns string type field's key
func (fld *stringField) Key() string {
	return fld.key
}

// Value returns string type field's value
func (fld *stringField) Value() interface{} {
	return fld.value
}

// uint64Field is used to create field in uint64 type
type durationField struct {
	key   string
	value time.Duration
}

// NewDurationField creates time.Duration type field
func NewDurationField(key string, value time.Duration) Field {
	return &durationField{key: key, value: value}
}

// Key returns uint64 type field's key
func (fld *durationField) Key() string {
	return fld.key
}

// Value returns uint64 type field's value
func (fld *durationField) Value() interface{} {
	return fld.value
}
