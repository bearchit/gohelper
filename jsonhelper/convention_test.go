package jsonhelper_test

import (
	"encoding/json"
	"github.com/bearchit/gohelper/jsonhelper"
	"testing"
	"time"
)

var model = struct {
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
}{
	Title:       "Example Title",
	Description: "whatever",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
	IsActive:    true,
}

func TestNewSnakeCaseMarshaller(t *testing.T) {
	encoded, _ := json.MarshalIndent(jsonhelper.NewSnakeCaseMarshaller(model), "", "  ")
	t.Log(string(encoded))
}

func TestNewLowerCamelCaseMarshaller(t *testing.T) {
	encoded, _ := json.MarshalIndent(jsonhelper.NewLowerCamelCaseMarshaller(model), "", "  ")
	t.Log(string(encoded))
}
