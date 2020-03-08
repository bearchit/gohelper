package httphelper

import (
	"net/http"
	"strings"
	"time"

	"github.com/monoculum/formam"
)

type FormDecoder struct {
	*formam.Decoder
}

func NewDecoder(decopt formam.DecoderOptions, options ...func(*FormDecoder)) *FormDecoder {
	d := &FormDecoder{Decoder: formam.NewDecoder(&decopt)}
	for _, opt := range options {
		opt(d)
	}

	return d
}

func WithTimeParser(layout string) func(*FormDecoder) {
	return func(decoder *FormDecoder) {
		decoder.RegisterCustomType(func(vals []string) (interface{}, error) {
			return time.Parse(layout, vals[0])
		}, []interface{}{time.Time{}}, nil)
	}
}

func WithCSV(delimiter string) func(*FormDecoder) {
	return func(decoder *FormDecoder) {
		decoder.RegisterCustomType(func(vals []string) (interface{}, error) {
			tokens := strings.Split(vals[0], delimiter)
			strs := make([]string, 0)
			for _, t := range tokens {
				strs = append(strs, t)
			}
			return strs, nil
		}, []interface{}{[]string{}}, nil)
	}
}

func DefaultDecoder() *FormDecoder {
	return NewDecoder(formam.DecoderOptions{TagName: "form", IgnoreUnknownKeys: true},
		WithTimeParser(time.RFC3339),
		//WithCSV(","),
	)
}

func DecodeForm(r *http.Request, v interface{}) error {
	return DecodeFormWith(DefaultDecoder(), r, v)
}

func DecodeFormWith(decoder *FormDecoder, r *http.Request, v interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	return decoder.Decode(r.Form, v)
}
