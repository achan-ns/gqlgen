// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package out

import (
	"fmt"
	"io"
	"strconv"
)

type MissingInterface interface {
	IsMissingInterface()
}

type MissingUnion interface {
	IsMissingUnion()
}

type MissingInput struct {
	Name *string      `json:"name"`
	Enum *MissingEnum `json:"enum"`
}

type MissingTypeNotNull struct {
	Name     string               `json:"name"`
	Enum     MissingEnum          `json:"enum"`
	Int      MissingInterface     `json:"int"`
	Existing *ExistingType        `json:"existing"`
	Missing2 *MissingTypeNullable `json:"missing2"`
}

func (MissingTypeNotNull) IsMissingInterface()  {}
func (MissingTypeNotNull) IsExistingInterface() {}
func (MissingTypeNotNull) IsMissingUnion()      {}
func (MissingTypeNotNull) IsExistingUnion()     {}

type MissingTypeNullable struct {
	Name     *string             `json:"name"`
	Enum     *MissingEnum        `json:"enum"`
	Int      MissingInterface    `json:"int"`
	Existing *ExistingType       `json:"existing"`
	Missing2 *MissingTypeNotNull `json:"missing2"`
}

func (MissingTypeNullable) IsMissingInterface()  {}
func (MissingTypeNullable) IsExistingInterface() {}
func (MissingTypeNullable) IsMissingUnion()      {}
func (MissingTypeNullable) IsExistingUnion()     {}

type MissingEnum string

const (
	MissingEnumHello   MissingEnum = "Hello"
	MissingEnumGoodbye MissingEnum = "Goodbye"
)

var AllMissingEnum = []MissingEnum{
	MissingEnumHello,
	MissingEnumGoodbye,
}

func (e MissingEnum) IsValid() bool {
	switch e {
	case MissingEnumHello, MissingEnumGoodbye:
		return true
	}
	return false
}

func (e MissingEnum) String() string {
	return string(e)
}

func (e *MissingEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MissingEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MissingEnum", str)
	}
	return nil
}

func (e MissingEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
