// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: necode/query/message.proto

package query

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

	enums "github.com/rskvp/necode/proto/gen/necode/enums"
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

	_ = enums.QueryResultType(0)
)

// Validate checks the field values on WorkflowQuery with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WorkflowQuery) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkflowQuery with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WorkflowQueryMultiError, or
// nil if none found.
func (m *WorkflowQuery) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkflowQuery) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QueryType

	if all {
		switch v := interface{}(m.GetQueryArgs()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WorkflowQueryValidationError{
					field:  "QueryArgs",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WorkflowQueryValidationError{
					field:  "QueryArgs",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQueryArgs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowQueryValidationError{
				field:  "QueryArgs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHeader()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WorkflowQueryValidationError{
					field:  "Header",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WorkflowQueryValidationError{
					field:  "Header",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHeader()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowQueryValidationError{
				field:  "Header",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return WorkflowQueryMultiError(errors)
	}

	return nil
}

// WorkflowQueryMultiError is an error wrapping multiple validation errors
// returned by WorkflowQuery.ValidateAll() if the designated constraints
// aren't met.
type WorkflowQueryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkflowQueryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkflowQueryMultiError) AllErrors() []error { return m }

// WorkflowQueryValidationError is the validation error returned by
// WorkflowQuery.Validate if the designated constraints aren't met.
type WorkflowQueryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowQueryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowQueryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowQueryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowQueryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowQueryValidationError) ErrorName() string { return "WorkflowQueryValidationError" }

// Error satisfies the builtin error interface
func (e WorkflowQueryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowQuery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowQueryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowQueryValidationError{}

// Validate checks the field values on WorkflowQueryResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WorkflowQueryResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkflowQueryResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WorkflowQueryResultMultiError, or nil if none found.
func (m *WorkflowQueryResult) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkflowQueryResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ResultType

	if all {
		switch v := interface{}(m.GetAnswer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WorkflowQueryResultValidationError{
					field:  "Answer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WorkflowQueryResultValidationError{
					field:  "Answer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAnswer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowQueryResultValidationError{
				field:  "Answer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ErrorMessage

	if len(errors) > 0 {
		return WorkflowQueryResultMultiError(errors)
	}

	return nil
}

// WorkflowQueryResultMultiError is an error wrapping multiple validation
// errors returned by WorkflowQueryResult.ValidateAll() if the designated
// constraints aren't met.
type WorkflowQueryResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkflowQueryResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkflowQueryResultMultiError) AllErrors() []error { return m }

// WorkflowQueryResultValidationError is the validation error returned by
// WorkflowQueryResult.Validate if the designated constraints aren't met.
type WorkflowQueryResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowQueryResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowQueryResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowQueryResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowQueryResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowQueryResultValidationError) ErrorName() string {
	return "WorkflowQueryResultValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowQueryResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowQueryResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowQueryResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowQueryResultValidationError{}

// Validate checks the field values on QueryRejected with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QueryRejected) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryRejected with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QueryRejectedMultiError, or
// nil if none found.
func (m *QueryRejected) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryRejected) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return QueryRejectedMultiError(errors)
	}

	return nil
}

// QueryRejectedMultiError is an error wrapping multiple validation errors
// returned by QueryRejected.ValidateAll() if the designated constraints
// aren't met.
type QueryRejectedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryRejectedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryRejectedMultiError) AllErrors() []error { return m }

// QueryRejectedValidationError is the validation error returned by
// QueryRejected.Validate if the designated constraints aren't met.
type QueryRejectedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryRejectedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryRejectedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryRejectedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryRejectedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryRejectedValidationError) ErrorName() string { return "QueryRejectedValidationError" }

// Error satisfies the builtin error interface
func (e QueryRejectedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryRejected.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryRejectedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryRejectedValidationError{}
