// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: necode/app/provider/resources.proto

package provider

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

// Validate checks the field values on OpenInAppURL with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OpenInAppURL) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OpenInAppURL with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OpenInAppURLMultiError, or
// nil if none found.
func (m *OpenInAppURL) ValidateAll() error {
	return m.validate(true)
}

func (m *OpenInAppURL) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppUrl

	// no validation rules for Method

	// no validation rules for FormParameters

	// no validation rules for Headers

	if len(errors) > 0 {
		return OpenInAppURLMultiError(errors)
	}

	return nil
}

// OpenInAppURLMultiError is an error wrapping multiple validation errors
// returned by OpenInAppURL.ValidateAll() if the designated constraints aren't met.
type OpenInAppURLMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OpenInAppURLMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OpenInAppURLMultiError) AllErrors() []error { return m }

// OpenInAppURLValidationError is the validation error returned by
// OpenInAppURL.Validate if the designated constraints aren't met.
type OpenInAppURLValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenInAppURLValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenInAppURLValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenInAppURLValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenInAppURLValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenInAppURLValidationError) ErrorName() string { return "OpenInAppURLValidationError" }

// Error satisfies the builtin error interface
func (e OpenInAppURLValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpenInAppURL.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenInAppURLValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenInAppURLValidationError{}
