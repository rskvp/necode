// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: necode/auth/registry/resources.proto

package registry

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ProviderInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProviderInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProviderInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProviderInfoMultiError, or
// nil if none found.
func (m *ProviderInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *ProviderInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProviderInfoValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProviderInfoValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProviderInfoValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ProviderType

	// no validation rules for Address

	// no validation rules for Description

	if len(errors) > 0 {
		return ProviderInfoMultiError(errors)
	}

	return nil
}

// ProviderInfoMultiError is an error wrapping multiple validation errors
// returned by ProviderInfo.ValidateAll() if the designated constraints aren't met.
type ProviderInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProviderInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProviderInfoMultiError) AllErrors() []error { return m }

// ProviderInfoValidationError is the validation error returned by
// ProviderInfo.Validate if the designated constraints aren't met.
type ProviderInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProviderInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProviderInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProviderInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProviderInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProviderInfoValidationError) ErrorName() string { return "ProviderInfoValidationError" }

// Error satisfies the builtin error interface
func (e ProviderInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProviderInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProviderInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProviderInfoValidationError{}
