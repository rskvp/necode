// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: necode/permissions/permissions_api.proto

package permissions

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

// Validate checks the field values on CheckPermissionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckPermissionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckPermissionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckPermissionRequestMultiError, or nil if none found.
func (m *CheckPermissionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckPermissionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Permission

	if all {
		switch v := interface{}(m.GetSubjectRef()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckPermissionRequestValidationError{
					field:  "SubjectRef",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckPermissionRequestValidationError{
					field:  "SubjectRef",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubjectRef()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckPermissionRequestValidationError{
				field:  "SubjectRef",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRef()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckPermissionRequestValidationError{
					field:  "Ref",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckPermissionRequestValidationError{
					field:  "Ref",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRef()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckPermissionRequestValidationError{
				field:  "Ref",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CheckPermissionRequestMultiError(errors)
	}

	return nil
}

// CheckPermissionRequestMultiError is an error wrapping multiple validation
// errors returned by CheckPermissionRequest.ValidateAll() if the designated
// constraints aren't met.
type CheckPermissionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckPermissionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckPermissionRequestMultiError) AllErrors() []error { return m }

// CheckPermissionRequestValidationError is the validation error returned by
// CheckPermissionRequest.Validate if the designated constraints aren't met.
type CheckPermissionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckPermissionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckPermissionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckPermissionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckPermissionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckPermissionRequestValidationError) ErrorName() string {
	return "CheckPermissionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CheckPermissionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckPermissionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckPermissionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckPermissionRequestValidationError{}

// Validate checks the field values on CheckPermissionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CheckPermissionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckPermissionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckPermissionResponseMultiError, or nil if none found.
func (m *CheckPermissionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckPermissionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CheckPermissionResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CheckPermissionResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckPermissionResponseValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CheckPermissionResponseMultiError(errors)
	}

	return nil
}

// CheckPermissionResponseMultiError is an error wrapping multiple validation
// errors returned by CheckPermissionResponse.ValidateAll() if the designated
// constraints aren't met.
type CheckPermissionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckPermissionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckPermissionResponseMultiError) AllErrors() []error { return m }

// CheckPermissionResponseValidationError is the validation error returned by
// CheckPermissionResponse.Validate if the designated constraints aren't met.
type CheckPermissionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckPermissionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckPermissionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckPermissionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckPermissionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckPermissionResponseValidationError) ErrorName() string {
	return "CheckPermissionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CheckPermissionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckPermissionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckPermissionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckPermissionResponseValidationError{}