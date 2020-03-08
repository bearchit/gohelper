package jsonhelper_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/bearchit/gohelper/jsonhelper"
)

var model = struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
}{
	ID:          "1234",
	Title:       "Example Title",
	Description: "whatever",
	CreatedAt:   time.Time{},
	UpdatedAt:   time.Time{},
	IsActive:    true,
}

func TestNewSnakeCaseMarshaller(t *testing.T) {
	encoded, err := json.MarshalIndent(jsonhelper.NewSnakeCaseMarshaller(model), "", "  ")
	require.NoError(t, err)
	assert.Equal(t, `{
  "id": "1234",
  "title": "Example Title",
  "description": "whatever",
  "created_at": "0001-01-01T00:00:00Z",
  "updated_at": "0001-01-01T00:00:00Z",
  "is_active": true
}`,
		string(encoded),
	)
}

func TestNewLowerCamelCaseMarshaller(t *testing.T) {
	encoded, err := json.MarshalIndent(jsonhelper.NewLowerCamelCaseMarshaller(model), "", "  ")
	require.NoError(t, err)
	assert.Equal(t, `{
  "id": "1234",
  "title": "Example Title",
  "description": "whatever",
  "createdAt": "0001-01-01T00:00:00Z",
  "updatedAt": "0001-01-01T00:00:00Z",
  "isActive": true
}`,
		string(encoded),
	)
}
