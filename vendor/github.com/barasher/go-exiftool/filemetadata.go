package exiftool

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	defaultString = ""
	defaultFloat  = float64(0)
	defaultInt    = int64(0)
)

// ErrKeyNotFound is a sentinel error used when a queried key does not exist
var ErrKeyNotFound = errors.New("key not found")

// FileMetadata is a structure that represents an exiftool extraction. File contains the
// filename that had to be extracted. If anything went wrong, Err will not be nil. Fields
// stores extracted fields.
type FileMetadata struct {
	File   string
	Fields map[string]interface{}
	Err    error
}

// GetString returns a field value as string and an error if one occurred.
// KeyNotFoundError will be returned if the key can't be found
func (fm FileMetadata) GetString(k string) (string, error) {
	v, found := fm.Fields[k]
	if !found || v == nil {
		return defaultString, ErrKeyNotFound
	}

	return toString(v), nil
}

func toString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case int64:
		return strconv.FormatInt(v, 10)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// GetFloat returns a field value as float64 and an error if one occurred.
// KeyNotFoundError will be returned if the key can't be found.
func (fm FileMetadata) GetFloat(k string) (float64, error) {
	v, found := fm.Fields[k]
	if !found || v == nil {
		return defaultFloat, ErrKeyNotFound
	}

	switch v := v.(type) {
	case string:
		return toFloatFallback(v)
	case float64:
		return v, nil
	case int64:
		return float64(v), nil
	default:
		str := fmt.Sprintf("%v", v)
		return toFloatFallback(str)
	}
}

func toFloatFallback(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, -1)
	if err != nil {
		return defaultFloat, fmt.Errorf("float64 parsing error (%v): %w", str, err)
	}

	return f, nil
}

// GetInt returns a field value as int64 and an error if one occurred.
// KeyNotFoundError will be returned if the key can't be found, ParseError if
// a parsing error occurs.
func (fm FileMetadata) GetInt(k string) (int64, error) {
	v, found := fm.Fields[k]
	if !found || v == nil {
		return defaultInt, ErrKeyNotFound
	}

	switch v := v.(type) {
	case string:
		return toIntFallback(v)
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	default:
		str := fmt.Sprintf("%v", v)
		return toIntFallback(str)
	}
}

func toIntFallback(str string) (int64, error) {
	f, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultInt, fmt.Errorf("int64 parsing error (%v): %w", str, err)
	}

	return f, nil
}

// GetStrings returns a field value as []string and an error if one occurred.
// KeyNotFoundError will be returned if the key can't be found.
func (fm FileMetadata) GetStrings(k string) ([]string, error) {
	v, found := fm.Fields[k]
	if !found || v == nil {
		return []string{}, ErrKeyNotFound
	}

	switch v := v.(type) {
	case []interface{}:
		is := v
		res := make([]string, len(is))

		for i, v2 := range is {
			res[i] = toString(v2)
		}

		return res, nil
	default:
		return []string{toString(v)}, nil
	}
}

func (fm FileMetadata) set(k string, v interface{}) {
	fm.Fields[k] = v
}

// SetString sets a string value for a specific field
func (fm FileMetadata) SetString(k string, v string) {
	fm.set(k, v)
}

// SetInt sets a int value for a specific field
func (fm FileMetadata) SetInt(k string, v int64) {
	fm.set(k, v)
}

// SetFloat sets a float value for a specific field
func (fm FileMetadata) SetFloat(k string, v float64) {
	fm.set(k, v)
}

// SetStrings sets a []String value for a specific field
func (fm FileMetadata) SetStrings(k string, v []string) {
	t := make([]interface{}, len(v))
	for i, c := range v {
		t[i] = c
	}
	fm.set(k, t)
}

// Clear removes value for a specific metadata field
func (fm FileMetadata) Clear(k string) {
	fm.set(k, nil)
}

// ClearAll removes all medatadata
func (fm FileMetadata) ClearAll() {
	for k, _ := range fm.Fields {
		fm.set(k, nil)
	}
}

// EmptyFileMetadata creates an empty FileMetadata struct
func EmptyFileMetadata() FileMetadata {
	return FileMetadata{
		Fields: make(map[string]interface{}),
	}
}
