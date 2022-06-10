// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package application

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseClientTypeFromString Parse ClientType from string
func ParseClientTypeFromString(str string) (ClientType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := ClientType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown ClientType: %s", str)
	}

	return ClientType(v), nil
}

// Equal type compare
func (t ClientType) Equal(target ClientType) bool {
	return t == target
}

// IsIn todo
func (t ClientType) IsIn(targets ...ClientType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t ClientType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *ClientType) UnmarshalJSON(b []byte) error {
	ins, err := ParseClientTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}