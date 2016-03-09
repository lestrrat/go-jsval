package jsval

import (
	"encoding/json"
	"errors"
	"time"
)

var ErrInvalidMaybeValue = errors.New("invalid Maybe value")

// Maybe is an interface that can be used for struct fields which
// want to differentiate between initialized and uninitialized state.
// For example, a string field, if uninitialized, will contain the zero
// value of "", but that empty string *could* be a valid value for
// our validation purposes.
//
// To differentiate between an uninitialized string and an empty string,
// you should wrap it with a wrapper that implements the Maybe interface
// and JSVal will do its best to figure this out
type Maybe interface {
	// Valid should return true if this value has been properly initialized.
	// If this returns false, JSVal will treat as if the field is has not been
	// provided at all.
	Valid() bool

	// Value should return whatever the underlying value is.
	Value() interface{}

	// Set sets a value to this Maybe value, and turns on the Valid flag.
	// An error may be returned if the value could not be set (e.g.
	// you provided a value with the wrong type)
	Set(interface{}) error

	// Reset clears the Maybe value, and sets the Valid flag to false.
	Reset()
}

type ValidFlag bool

func (v *ValidFlag) Reset() {
	*v = false
}

func (v ValidFlag) Valid() bool {
	return bool(v)
}

type MaybeBool struct {
	ValidFlag
	Bool bool
}

func (v *MaybeBool) Set(x interface{}) error {
	s, ok := x.(bool)
	if !ok {
		return ErrInvalidMaybeValue
	}
	v.ValidFlag = true
	v.Bool = s
	return nil
}

func (v MaybeBool) Value() interface{} {
	return v.Bool
}

func (v MaybeBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Bool)
}

func (v *MaybeBool) UnmarshalJSON(data []byte) error {
	var in bool
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Set(in)
}

type MaybeFloat struct {
	ValidFlag
	Float float64
}

func (v *MaybeFloat) Set(x interface{}) error {
	s, ok := x.(float64)
	if !ok {
		return ErrInvalidMaybeValue
	}
	v.ValidFlag = true
	v.Float = s
	return nil
}

func (v MaybeFloat) Value() interface{} {
	return v.Float
}

func (v MaybeFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Float)
}

func (v *MaybeFloat) UnmarshalJSON(data []byte) error {
	var in float64
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Set(in)
}

type MaybeInt struct {
	ValidFlag
	Int int64
}

func (v *MaybeInt) Set(x interface{}) error {
	s, ok := x.(int64)
	if !ok {
		return ErrInvalidMaybeValue
	}
	v.ValidFlag = true
	v.Int = s
	return nil
}

func (v MaybeInt) Value() interface{} {
	return v.Int
}

func (v MaybeInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Int)
}

func (v *MaybeInt) UnmarshalJSON(data []byte) error {
	var in int64
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Set(in)
}

type MaybeString struct {
	ValidFlag
	String string
}

func (v *MaybeString) Set(x interface{}) error {
	s, ok := x.(string)
	if !ok {
		return ErrInvalidMaybeValue
	}
	v.ValidFlag = true
	v.String = s
	return nil
}

func (v MaybeString) Value() interface{} {
	return v.String
}

func (v MaybeString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String)
}

func (v *MaybeString) UnmarshalJSON(data []byte) error {
	var in string
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Set(in)
}

type MaybeTime struct {
	ValidFlag
	Time time.Time
}

func (v *MaybeTime) Set(x interface{}) error {
	s, ok := x.(time.Time)
	if !ok {
		return ErrInvalidMaybeValue
	}
	v.ValidFlag = true
	v.Time = s
	return nil
}

func (v MaybeTime) Value() interface{} {
	return v.Time
}

func (v MaybeTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Time)
}

func (v *MaybeTime) UnmarshalJSON(data []byte) error {
	var in time.Time
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}
	return v.Set(in)
}