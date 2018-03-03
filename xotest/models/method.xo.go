// Package models contains the types for schema 'AllergyNew'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// Method is the 'method' enum type from schema 'AllergyNew'.
type Method uint16

const (
	// MethodSlt is the 'slt' Method.
	MethodSlt = Method(1)

	// MethodInj is the 'inj' Method.
	MethodInj = Method(2)

	// MethodDefault is the 'default' Method.
	MethodDefault = Method(3)
)

// String returns the string value of the Method.
func (m Method) String() string {
	var enumVal string

	switch m {
	case MethodSlt:
		enumVal = "slt"

	case MethodInj:
		enumVal = "inj"

	case MethodDefault:
		enumVal = "default"
	}

	return enumVal
}

// MarshalText marshals Method into text.
func (m Method) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// UnmarshalText unmarshals Method from text.
func (m *Method) UnmarshalText(text []byte) error {
	switch string(text) {
	case "slt":
		*m = MethodSlt

	case "inj":
		*m = MethodInj

	case "default":
		*m = MethodDefault

	default:
		return errors.New("invalid Method")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for Method.
func (m Method) Value() (driver.Value, error) {
	return m.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for Method.
func (m *Method) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid Method")
	}

	return m.UnmarshalText(buf)
}
