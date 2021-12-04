package dataparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Default type
func TestParseField(t *testing.T) {
	assert := assert.New(t)

	field := ParseField("name")
	assert.Equal(field.name, "name")
	assert.Equal(field.kind, "string")
}

func TestParseFieldWithKind(t *testing.T) {
	field := ParseField("enabled:boolean")

	assert.Equal(t, field.kind, "boolean")
}

func TestFetchValue(t *testing.T) {
	field := Field{"user", "string"}
	value := FetchValue(field)

	assert.Equal(t, "michael", value)
}

func TestFetchBooleanValue(t *testing.T) {
	oldOsGetenv := osGetenv
	defer func() { osGetenv = oldOsGetenv }()
	mockOsGetenv := func(_ string) string {
		return "true"
	}
	osGetenv = mockOsGetenv

	field := Field{"enabled", "boolean"}
	value := FetchValue(field)

	assert.Equal(t, true, value)
}

func TestFetchListValue(t *testing.T) {
	oldOsGetenv := osGetenv
	defer func() { osGetenv = oldOsGetenv }()
	mockOsGetenv := func(_ string) string {
		return "red,blue,green"
	}
	osGetenv = mockOsGetenv

	field := Field{"colors", "list"}
	value := FetchValue(field)

	assert.Len(t, value, 3)
}
