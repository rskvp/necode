// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: necode/app/provider/provider_api.proto

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

// Validate checks the field values on OpenInAppRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OpenInAppRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OpenInAppRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OpenInAppRequestMultiError, or nil if none found.
func (m *OpenInAppRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OpenInAppRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenInAppRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenInAppRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenInAppRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResourceInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenInAppRequestValidationError{
					field:  "ResourceInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenInAppRequestValidationError{
					field:  "ResourceInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResourceInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenInAppRequestValidationError{
				field:  "ResourceInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ViewMode

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return OpenInAppRequestMultiError(errors)
	}

	return nil
}

// OpenInAppRequestMultiError is an error wrapping multiple validation errors
// returned by OpenInAppRequest.ValidateAll() if the designated constraints
// aren't met.
type OpenInAppRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OpenInAppRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OpenInAppRequestMultiError) AllErrors() []error { return m }

// OpenInAppRequestValidationError is the validation error returned by
// OpenInAppRequest.Validate if the designated constraints aren't met.
type OpenInAppRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenInAppRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenInAppRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenInAppRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenInAppRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenInAppRequestValidationError) ErrorName() string { return "OpenInAppRequestValidationError" }

// Error satisfies the builtin error interface
func (e OpenInAppRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpenInAppRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenInAppRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenInAppRequestValidationError{}

// Validate checks the field values on OpenInAppResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OpenInAppResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OpenInAppResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OpenInAppResponseMultiError, or nil if none found.
func (m *OpenInAppResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OpenInAppResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenInAppResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenInAppResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenInAppResponseValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenInAppResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenInAppResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenInAppResponseValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAppUrl()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenInAppResponseValidationError{
					field:  "AppUrl",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenInAppResponseValidationError{
					field:  "AppUrl",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAppUrl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenInAppResponseValidationError{
				field:  "AppUrl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OpenInAppResponseMultiError(errors)
	}

	return nil
}

// OpenInAppResponseMultiError is an error wrapping multiple validation errors
// returned by OpenInAppResponse.ValidateAll() if the designated constraints
// aren't met.
type OpenInAppResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OpenInAppResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OpenInAppResponseMultiError) AllErrors() []error { return m }

// OpenInAppResponseValidationError is the validation error returned by
// OpenInAppResponse.Validate if the designated constraints aren't met.
type OpenInAppResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenInAppResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenInAppResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenInAppResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenInAppResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenInAppResponseValidationError) ErrorName() string {
	return "OpenInAppResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OpenInAppResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpenInAppResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenInAppResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenInAppResponseValidationError{}
