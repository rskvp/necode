// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: necode/auth/registry/registry_api.proto

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

// Validate checks the field values on GetAuthProvidersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAuthProvidersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAuthProvidersRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAuthProvidersRequestMultiError, or nil if none found.
func (m *GetAuthProvidersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAuthProvidersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetAuthProvidersRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetAuthProvidersRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAuthProvidersRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Type

	if len(errors) > 0 {
		return GetAuthProvidersRequestMultiError(errors)
	}

	return nil
}

// GetAuthProvidersRequestMultiError is an error wrapping multiple validation
// errors returned by GetAuthProvidersRequest.ValidateAll() if the designated
// constraints aren't met.
type GetAuthProvidersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAuthProvidersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAuthProvidersRequestMultiError) AllErrors() []error { return m }

// GetAuthProvidersRequestValidationError is the validation error returned by
// GetAuthProvidersRequest.Validate if the designated constraints aren't met.
type GetAuthProvidersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAuthProvidersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAuthProvidersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAuthProvidersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAuthProvidersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAuthProvidersRequestValidationError) ErrorName() string {
	return "GetAuthProvidersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAuthProvidersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAuthProvidersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAuthProvidersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAuthProvidersRequestValidationError{}

// Validate checks the field values on GetAuthProvidersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAuthProvidersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAuthProvidersResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAuthProvidersResponseMultiError, or nil if none found.
func (m *GetAuthProvidersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAuthProvidersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetAuthProvidersResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetAuthProvidersResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAuthProvidersResponseValidationError{
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
				errors = append(errors, GetAuthProvidersResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetAuthProvidersResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAuthProvidersResponseValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetProviders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetAuthProvidersResponseValidationError{
						field:  fmt.Sprintf("Providers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetAuthProvidersResponseValidationError{
						field:  fmt.Sprintf("Providers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetAuthProvidersResponseValidationError{
					field:  fmt.Sprintf("Providers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetAuthProvidersResponseMultiError(errors)
	}

	return nil
}

// GetAuthProvidersResponseMultiError is an error wrapping multiple validation
// errors returned by GetAuthProvidersResponse.ValidateAll() if the designated
// constraints aren't met.
type GetAuthProvidersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAuthProvidersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAuthProvidersResponseMultiError) AllErrors() []error { return m }

// GetAuthProvidersResponseValidationError is the validation error returned by
// GetAuthProvidersResponse.Validate if the designated constraints aren't met.
type GetAuthProvidersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAuthProvidersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAuthProvidersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAuthProvidersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAuthProvidersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAuthProvidersResponseValidationError) ErrorName() string {
	return "GetAuthProvidersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAuthProvidersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAuthProvidersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAuthProvidersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAuthProvidersResponseValidationError{}

// Validate checks the field values on ListAuthProvidersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAuthProvidersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAuthProvidersRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAuthProvidersRequestMultiError, or nil if none found.
func (m *ListAuthProvidersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAuthProvidersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListAuthProvidersRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListAuthProvidersRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListAuthProvidersRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListAuthProvidersRequestMultiError(errors)
	}

	return nil
}

// ListAuthProvidersRequestMultiError is an error wrapping multiple validation
// errors returned by ListAuthProvidersRequest.ValidateAll() if the designated
// constraints aren't met.
type ListAuthProvidersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAuthProvidersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAuthProvidersRequestMultiError) AllErrors() []error { return m }

// ListAuthProvidersRequestValidationError is the validation error returned by
// ListAuthProvidersRequest.Validate if the designated constraints aren't met.
type ListAuthProvidersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAuthProvidersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAuthProvidersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAuthProvidersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAuthProvidersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAuthProvidersRequestValidationError) ErrorName() string {
	return "ListAuthProvidersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAuthProvidersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAuthProvidersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAuthProvidersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAuthProvidersRequestValidationError{}

// Validate checks the field values on ListAuthProvidersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAuthProvidersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAuthProvidersResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAuthProvidersResponseMultiError, or nil if none found.
func (m *ListAuthProvidersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAuthProvidersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListAuthProvidersResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListAuthProvidersResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListAuthProvidersResponseValidationError{
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
				errors = append(errors, ListAuthProvidersResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListAuthProvidersResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListAuthProvidersResponseValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetProviders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAuthProvidersResponseValidationError{
						field:  fmt.Sprintf("Providers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAuthProvidersResponseValidationError{
						field:  fmt.Sprintf("Providers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAuthProvidersResponseValidationError{
					field:  fmt.Sprintf("Providers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAuthProvidersResponseMultiError(errors)
	}

	return nil
}

// ListAuthProvidersResponseMultiError is an error wrapping multiple validation
// errors returned by ListAuthProvidersResponse.ValidateAll() if the
// designated constraints aren't met.
type ListAuthProvidersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAuthProvidersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAuthProvidersResponseMultiError) AllErrors() []error { return m }

// ListAuthProvidersResponseValidationError is the validation error returned by
// ListAuthProvidersResponse.Validate if the designated constraints aren't met.
type ListAuthProvidersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAuthProvidersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAuthProvidersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAuthProvidersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAuthProvidersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAuthProvidersResponseValidationError) ErrorName() string {
	return "ListAuthProvidersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAuthProvidersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAuthProvidersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAuthProvidersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAuthProvidersResponseValidationError{}
