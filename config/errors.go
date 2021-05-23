package config

import (
	"fmt"
)

type InvalidConfigError struct {
	field, value string
}

func NewInvalidConfigError(field string, value string) InvalidConfigError {
	return InvalidConfigError{field: field, value: value}
}

func (ice InvalidConfigError) Error() string {
	return fmt.Sprintf("invalid configuration in field %q with value %q", ice.field, ice.value)
}

type ParseConfigError struct {
	err error
}

func NewParseConfigError(err error) ParseConfigError {
	return ParseConfigError{err: err}
}

func (pce ParseConfigError) Error() string {
	return fmt.Sprintf("error parsing parameters: %s", pce.err.Error())
}
