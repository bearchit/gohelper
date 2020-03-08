package jsonhelper

import (
	"bytes"
	"encoding/json"
	"regexp"
	"unicode"
	"unicode/utf8"
)

type caseMarshaller struct {
	Replacer CaseReplacer
	Value    interface{}
}

var (
	keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
)

func NewSnakeCaseMarshaller(v interface{}) *caseMarshaller {
	return &caseMarshaller{
		Replacer: toSnakeCase,
		Value:    v,
	}
}

func NewLowerCamelCaseMarshaller(v interface{}) *caseMarshaller {
	return &caseMarshaller{
		Replacer: toLowerCamelCase,
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
		c.Replacer,
	)

	return converted, err
}

type CaseReplacer func([]byte) []byte

func toSnakeCase(match []byte) []byte {
	matcher := regexp.MustCompile(`(\w)([A-Z])`)
	return bytes.ToLower(matcher.ReplaceAll(
		match,
		[]byte(`${1}_${2}`),
	))
}

func toLowerCamelCase(match []byte) []byte {
	if len(match) > 2 {
		r, width := utf8.DecodeRune(match[1:])
		r = unicode.ToLower(r)
		utf8.EncodeRune(match[1:width+1], r)
	}
	return match
}
