package jsonhelper

import (
	"encoding/json"
	"regexp"

	"github.com/iancoleman/strcase"
)

type CaseReplacer func(string) string

type caseMarshaller struct {
	Replacer CaseReplacer
	Value    interface{}
}

var (
	keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
)

func NewSnakeCaseMarshaller(v interface{}) *caseMarshaller {
	return &caseMarshaller{
		Replacer: strcase.ToSnake,
		Value:    v,
	}
}

func NewLowerCamelCaseMarshaller(v interface{}) *caseMarshaller {
	return &caseMarshaller{
		Replacer: strcase.ToLowerCamel,
		Value:    v,
	}
}

func (c caseMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)
	if err != nil {
		return nil, err
	}

	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matcher := regexp.MustCompile(`(\w+)`)
			replaced := []byte(c.Replacer(string(matcher.Find(match))))
			return matcher.ReplaceAll(
				match,
				replaced,
			)
		},
	)

	return converted, err
}
