// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: necode/permissions/resources.proto

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

// Validate checks the field values on SubjectReference with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SubjectReference) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubjectReference with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubjectReferenceMultiError, or nil if none found.
func (m *SubjectReference) ValidateAll() error {
	return m.validate(true)
}

func (m *SubjectReference) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Spec.(type) {
	case *SubjectReference_UserId:
		if v == nil {
			err := SubjectReferenceValidationError{
				field:  "Spec",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetUserId()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SubjectReferenceValidationError{
						field:  "UserId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SubjectReferenceValidationError{
						field:  "UserId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUserId()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubjectReferenceValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *SubjectReference_GroupId:
		if v == nil {
			err := SubjectReferenceValidationError{
				field:  "Spec",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetGroupId()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SubjectReferenceValidationError{
						field:  "GroupId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SubjectReferenceValidationError{
						field:  "GroupId",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGroupId()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubjectReferenceValidationError{
					field:  "GroupId",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return SubjectReferenceMultiError(errors)
	}

	return nil
}

// SubjectReferenceMultiError is an error wrapping multiple validation errors
// returned by SubjectReference.ValidateAll() if the designated constraints
// aren't met.
type SubjectReferenceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectReferenceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectReferenceMultiError) AllErrors() []error { return m }

// SubjectReferenceValidationError is the validation error returned by
// SubjectReference.Validate if the designated constraints aren't met.
type SubjectReferenceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectReferenceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectReferenceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectReferenceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectReferenceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectReferenceValidationError) ErrorName() string { return "SubjectReferenceValidationError" }

// Error satisfies the builtin error interface
func (e SubjectReferenceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubjectReference.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectReferenceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectReferenceValidationError{}
