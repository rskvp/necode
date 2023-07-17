// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: necode/ocm/provider/provider_api.proto

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

// Validate checks the field values on IsProviderAllowedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IsProviderAllowedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IsProviderAllowedRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IsProviderAllowedRequestMultiError, or nil if none found.
func (m *IsProviderAllowedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IsProviderAllowedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IsProviderAllowedRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IsProviderAllowedRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IsProviderAllowedRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetProvider()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IsProviderAllowedRequestValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IsProviderAllowedRequestValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProvider()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IsProviderAllowedRequestValidationError{
				field:  "Provider",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IsProviderAllowedRequestMultiError(errors)
	}

	return nil
}

// IsProviderAllowedRequestMultiError is an error wrapping multiple validation
// errors returned by IsProviderAllowedRequest.ValidateAll() if the designated
// constraints aren't met.
type IsProviderAllowedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IsProviderAllowedRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IsProviderAllowedRequestMultiError) AllErrors() []error { return m }

// IsProviderAllowedRequestValidationError is the validation error returned by
// IsProviderAllowedRequest.Validate if the designated constraints aren't met.
type IsProviderAllowedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsProviderAllowedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsProviderAllowedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsProviderAllowedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsProviderAllowedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsProviderAllowedRequestValidationError) ErrorName() string {
	return "IsProviderAllowedRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IsProviderAllowedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsProviderAllowedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsProviderAllowedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsProviderAllowedRequestValidationError{}

// Validate checks the field values on IsProviderAllowedResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IsProviderAllowedResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IsProviderAllowedResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IsProviderAllowedResponseMultiError, or nil if none found.
func (m *IsProviderAllowedResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IsProviderAllowedResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IsProviderAllowedResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IsProviderAllowedResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IsProviderAllowedResponseValidationError{
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
				errors = append(errors, IsProviderAllowedResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IsProviderAllowedResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IsProviderAllowedResponseValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IsProviderAllowedResponseMultiError(errors)
	}

	return nil
}

// IsProviderAllowedResponseMultiError is an error wrapping multiple validation
// errors returned by IsProviderAllowedResponse.ValidateAll() if the
// designated constraints aren't met.
type IsProviderAllowedResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IsProviderAllowedResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IsProviderAllowedResponseMultiError) AllErrors() []error { return m }

// IsProviderAllowedResponseValidationError is the validation error returned by
// IsProviderAllowedResponse.Validate if the designated constraints aren't met.
type IsProviderAllowedResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsProviderAllowedResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsProviderAllowedResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsProviderAllowedResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsProviderAllowedResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsProviderAllowedResponseValidationError) ErrorName() string {
	return "IsProviderAllowedResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IsProviderAllowedResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsProviderAllowedResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsProviderAllowedResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsProviderAllowedResponseValidationError{}

// Validate checks the field values on GetInfoByDomainRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetInfoByDomainRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetInfoByDomainRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetInfoByDomainRequestMultiError, or nil if none found.
func (m *GetInfoByDomainRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetInfoByDomainRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetInfoByDomainRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetInfoByDomainRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetInfoByDomainRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Domain

	if len(errors) > 0 {
		return GetInfoByDomainRequestMultiError(errors)
	}

	return nil
}

// GetInfoByDomainRequestMultiError is an error wrapping multiple validation
// errors returned by GetInfoByDomainRequest.ValidateAll() if the designated
// constraints aren't met.
type GetInfoByDomainRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetInfoByDomainRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetInfoByDomainRequestMultiError) AllErrors() []error { return m }

// GetInfoByDomainRequestValidationError is the validation error returned by
// GetInfoByDomainRequest.Validate if the designated constraints aren't met.
type GetInfoByDomainRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoByDomainRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoByDomainRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoByDomainRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoByDomainRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoByDomainRequestValidationError) ErrorName() string {
	return "GetInfoByDomainRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetInfoByDomainRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoByDomainRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoByDomainRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoByDomainRequestValidationError{}

// Validate checks the field values on GetInfoByDomainResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetInfoByDomainResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetInfoByDomainResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetInfoByDomainResponseMultiError, or nil if none found.
func (m *GetInfoByDomainResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetInfoByDomainResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetInfoByDomainResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetInfoByDomainResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetInfoByDomainResponseValidationError{
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
				errors = append(errors, GetInfoByDomainResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetInfoByDomainResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetInfoByDomainResponseValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetProviderInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetInfoByDomainResponseValidationError{
					field:  "ProviderInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetInfoByDomainResponseValidationError{
					field:  "ProviderInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProviderInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetInfoByDomainResponseValidationError{
				field:  "ProviderInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetInfoByDomainResponseMultiError(errors)
	}

	return nil
}

// GetInfoByDomainResponseMultiError is an error wrapping multiple validation
// errors returned by GetInfoByDomainResponse.ValidateAll() if the designated
// constraints aren't met.
type GetInfoByDomainResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetInfoByDomainResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetInfoByDomainResponseMultiError) AllErrors() []error { return m }

// GetInfoByDomainResponseValidationError is the validation error returned by
// GetInfoByDomainResponse.Validate if the designated constraints aren't met.
type GetInfoByDomainResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoByDomainResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoByDomainResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoByDomainResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoByDomainResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoByDomainResponseValidationError) ErrorName() string {
	return "GetInfoByDomainResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetInfoByDomainResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoByDomainResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoByDomainResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoByDomainResponseValidationError{}

// Validate checks the field values on ListAllProvidersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAllProvidersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAllProvidersRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAllProvidersRequestMultiError, or nil if none found.
func (m *ListAllProvidersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAllProvidersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListAllProvidersRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListAllProvidersRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListAllProvidersRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListAllProvidersRequestMultiError(errors)
	}

	return nil
}

// ListAllProvidersRequestMultiError is an error wrapping multiple validation
// errors returned by ListAllProvidersRequest.ValidateAll() if the designated
// constraints aren't met.
type ListAllProvidersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAllProvidersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAllProvidersRequestMultiError) AllErrors() []error { return m }

// ListAllProvidersRequestValidationError is the validation error returned by
// ListAllProvidersRequest.Validate if the designated constraints aren't met.
type ListAllProvidersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAllProvidersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAllProvidersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAllProvidersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAllProvidersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAllProvidersRequestValidationError) ErrorName() string {
	return "ListAllProvidersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAllProvidersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAllProvidersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAllProvidersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAllProvidersRequestValidationError{}

// Validate checks the field values on ListAllProvidersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAllProvidersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAllProvidersResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAllProvidersResponseMultiError, or nil if none found.
func (m *ListAllProvidersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAllProvidersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListAllProvidersResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListAllProvidersResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListAllProvidersResponseValidationError{
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
				errors = append(errors, ListAllProvidersResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListAllProvidersResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListAllProvidersResponseValidationError{
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
					errors = append(errors, ListAllProvidersResponseValidationError{
						field:  fmt.Sprintf("Providers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAllProvidersResponseValidationError{
						field:  fmt.Sprintf("Providers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAllProvidersResponseValidationError{
					field:  fmt.Sprintf("Providers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAllProvidersResponseMultiError(errors)
	}

	return nil
}

// ListAllProvidersResponseMultiError is an error wrapping multiple validation
// errors returned by ListAllProvidersResponse.ValidateAll() if the designated
// constraints aren't met.
type ListAllProvidersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAllProvidersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAllProvidersResponseMultiError) AllErrors() []error { return m }

// ListAllProvidersResponseValidationError is the validation error returned by
// ListAllProvidersResponse.Validate if the designated constraints aren't met.
type ListAllProvidersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAllProvidersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAllProvidersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAllProvidersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAllProvidersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAllProvidersResponseValidationError) ErrorName() string {
	return "ListAllProvidersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAllProvidersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAllProvidersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAllProvidersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAllProvidersResponseValidationError{}