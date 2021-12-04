package dataparser

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Field struct {
	name string
	kind string
}

var osGetenv = os.Getenv

func ParseFields(input string) []Field {
	fieldDefs := strings.Fields(input)
	fields := make([]Field, len(fieldDefs))

	for i, fieldDef := range fieldDefs {
		fields[i] = ParseField((fieldDef))
	}

	return fields
}

func ParseField(fieldDef string) Field {
	segments := strings.Split(fieldDef, ":")

	name := segments[0]
	var kind string

	if len(segments) > 1 {
		kind = segments[1]
	} else {
		kind = "string"
	}

	return Field{name, kind}
}

func FetchData(fields []Field) map[string]interface{} {
	dataMap := make(map[string]interface{})

	for _, field := range fields {
		dataMap[field.name] = FetchValue(field)
	}

	return dataMap
}

func FetchValue(field Field) interface{} {
	envValue := osGetenv(strings.ToUpper(field.name))
	var value interface{}

	switch field.kind {
	case "boolean":
		value, _ = strconv.ParseBool(envValue)
	case "list":
		value, _ = csv.NewReader(strings.NewReader(envValue)).Read()
	default:
		value = envValue
	}

	return value
}
