// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: necode/admin/group/group_api.proto

package group

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

// Validate checks the field values on CreateGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateGroupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateGroupRequestMultiError, or nil if none found.
func (m *CreateGroupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateGroupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateGroupRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateGroupRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateGroupRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetGroup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateGroupRequestValidationError{
					field:  "Group",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateGroupRequestValidationError{
					field:  "Group",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateGroupRequestValidationError{
				field:  "Group",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateGroupRequestMultiError(errors)
	}

	return nil
}

// CreateGroupRequestMultiError is an error wrapping multiple validation errors
// returned by CreateGroupRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateGroupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateGroupRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateGroupRequestMultiError) AllErrors() []error { return m }

// CreateGroupRequestValidationError is the validation error returned by
// CreateGroupRequest.Validate if the designated constraints aren't met.
type CreateGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateGroupRequestValidationError) ErrorName() string {
	return "CreateGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateGroupRequestValidationError{}

// Validate checks the field values on CreateGroupResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateGroupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateGroupResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateGroupResponseMultiError, or nil if none found.
func (m *CreateGroupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateGroupResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateGroupResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateGroupResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateGroupResponseValidationError{
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
				errors = append(errors, CreateGroupResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateGroupResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateGroupResponseValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetGroup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateGroupResponseValidationError{
					field:  "Group",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateGroupResponseValidationError{
					field:  "Group",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateGroupResponseValidationError{
				field:  "Group",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateGroupResponseMultiError(errors)
	}

	return nil
}

// CreateGroupResponseMultiError is an error wrapping multiple validation
// errors returned by CreateGroupResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateGroupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateGroupResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateGroupResponseMultiError) AllErrors() []error { return m }

// CreateGroupResponseValidationError is the validation error returned by
// CreateGroupResponse.Validate if the designated constraints aren't met.
type CreateGroupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateGroupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateGroupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateGroupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateGroupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateGroupResponseValidationError) ErrorName() string {
	return "CreateGroupResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateGroupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateGroupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateGroupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateGroupResponseValidationError{}

// Validate checks the field values on DeleteGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteGroupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteGroupRequestMultiError, or nil if none found.
func (m *DeleteGroupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteGroupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteGroupRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteGroupRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteGroupRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetGroupId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteGroupRequestValidationError{
					field:  "GroupId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteGroupRequestValidationError{
					field:  "GroupId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGroupId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteGroupRequestValidationError{
				field:  "GroupId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeleteGroupRequestMultiError(errors)
	}

	return nil
}

// DeleteGroupRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteGroupRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteGroupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteGroupRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteGroupRequestMultiError) AllErrors() []error { return m }

// DeleteGroupRequestValidationError is the validation error returned by
// DeleteGroupRequest.Validate if the designated constraints aren't met.
type DeleteGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteGroupRequestValidationError) ErrorName() string {
	return "DeleteGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteGroupRequestValidationError{}

// Validate checks the field values on DeleteGroupResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteGroupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteGroupResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteGroupResponseMultiError, or nil if none found.
func (m *DeleteGroupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteGroupResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteGroupResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteGroupResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteGroupResponseValidationError{
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
				errors = append(errors, DeleteGroupResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteGroupResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteGroupResponseValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeleteGroupResponseMultiError(errors)
	}

	return nil
}

// DeleteGroupResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteGroupResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteGroupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteGroupResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteGroupResponseMultiError) AllErrors() []error { return m }

// DeleteGroupResponseValidationError is the validation error returned by
// DeleteGroupResponse.Validate if the designated constraints aren't met.
type DeleteGroupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteGroupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteGroupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteGroupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteGroupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteGroupResponseValidationError) ErrorName() string {
	return "DeleteGroupResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteGroupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteGroupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteGroupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteGroupResponseValidationError{}

// Validate checks the field values on AddUserToGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddUserToGroupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddUserToGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddUserToGroupRequestMultiError, or nil if none found.
func (m *AddUserToGroupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddUserToGroupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUserId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddUserToGroupRequestValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddUserToGroupRequestValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddUserToGroupRequestValidationError{
				field:  "UserId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetGroupId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddUserToGroupRequestValidationError{
					field:  "GroupId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddUserToGroupRequestValidationError{
					field:  "GroupId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGroupId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddUserToGroupRequestValidationError{
				field:  "GroupId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddUserToGroupRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddUserToGroupRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddUserToGroupRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddUserToGroupRequestMultiError(errors)
	}

	return nil
}

// AddUserToGroupRequestMultiError is an error wrapping multiple validation
// errors returned by AddUserToGroupRequest.ValidateAll() if the designated
// constraints aren't met.
type AddUserToGroupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddUserToGroupRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddUserToGroupRequestMultiError) AllErrors() []error { return m }

// AddUserToGroupRequestValidationError is the validation error returned by
// AddUserToGroupRequest.Validate if the designated constraints aren't met.
type AddUserToGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddUserToGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddUserToGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddUserToGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddUserToGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddUserToGroupRequestValidationError) ErrorName() string {
	return "AddUserToGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddUserToGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddUserToGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddUserToGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddUserToGroupRequestValidationError{}

// Validate checks the field values on AddUserToGroupResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddUserToGroupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddUserToGroupResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddUserToGroupResponseMultiError, or nil if none found.
func (m *AddUserToGroupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddUserToGroupResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddUserToGroupResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddUserToGroupResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddUserToGroupResponseValidationError{
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
				errors = append(errors, AddUserToGroupResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddUserToGroupResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddUserToGroupResponseValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddUserToGroupResponseMultiError(errors)
	}

	return nil
}

// AddUserToGroupResponseMultiError is an error wrapping multiple validation
// errors returned by AddUserToGroupResponse.ValidateAll() if the designated
// constraints aren't met.
type AddUserToGroupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddUserToGroupResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddUserToGroupResponseMultiError) AllErrors() []error { return m }

// AddUserToGroupResponseValidationError is the validation error returned by
// AddUserToGroupResponse.Validate if the designated constraints aren't met.
type AddUserToGroupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddUserToGroupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddUserToGroupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddUserToGroupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddUserToGroupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddUserToGroupResponseValidationError) ErrorName() string {
	return "AddUserToGroupResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddUserToGroupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddUserToGroupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddUserToGroupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddUserToGroupResponseValidationError{}

// Validate checks the field values on RemoveUserFromGroupRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveUserFromGroupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveUserFromGroupRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveUserFromGroupRequestMultiError, or nil if none found.
func (m *RemoveUserFromGroupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveUserFromGroupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUserId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemoveUserFromGroupRequestValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveUserFromGroupRequestValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveUserFromGroupRequestValidationError{
				field:  "UserId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetGroupId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemoveUserFromGroupRequestValidationError{
					field:  "GroupId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveUserFromGroupRequestValidationError{
					field:  "GroupId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGroupId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveUserFromGroupRequestValidationError{
				field:  "GroupId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOpaque()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemoveUserFromGroupRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveUserFromGroupRequestValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveUserFromGroupRequestValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RemoveUserFromGroupRequestMultiError(errors)
	}

	return nil
}

// RemoveUserFromGroupRequestMultiError is an error wrapping multiple
// validation errors returned by RemoveUserFromGroupRequest.ValidateAll() if
// the designated constraints aren't met.
type RemoveUserFromGroupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveUserFromGroupRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveUserFromGroupRequestMultiError) AllErrors() []error { return m }

// RemoveUserFromGroupRequestValidationError is the validation error returned
// by RemoveUserFromGroupRequest.Validate if the designated constraints aren't met.
type RemoveUserFromGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveUserFromGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveUserFromGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveUserFromGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveUserFromGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveUserFromGroupRequestValidationError) ErrorName() string {
	return "RemoveUserFromGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveUserFromGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveUserFromGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveUserFromGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveUserFromGroupRequestValidationError{}

// Validate checks the field values on RemoveUserFromGroupResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveUserFromGroupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveUserFromGroupResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RemoveUserFromGroupResponseMultiError, or nil if none found.
func (m *RemoveUserFromGroupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveUserFromGroupResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RemoveUserFromGroupResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveUserFromGroupResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveUserFromGroupResponseValidationError{
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
				errors = append(errors, RemoveUserFromGroupResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RemoveUserFromGroupResponseValidationError{
					field:  "Opaque",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpaque()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoveUserFromGroupResponseValidationError{
				field:  "Opaque",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RemoveUserFromGroupResponseMultiError(errors)
	}

	return nil
}

// RemoveUserFromGroupResponseMultiError is an error wrapping multiple
// validation errors returned by RemoveUserFromGroupResponse.ValidateAll() if
// the designated constraints aren't met.
type RemoveUserFromGroupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveUserFromGroupResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveUserFromGroupResponseMultiError) AllErrors() []error { return m }

// RemoveUserFromGroupResponseValidationError is the validation error returned
// by RemoveUserFromGroupResponse.Validate if the designated constraints
// aren't met.
type RemoveUserFromGroupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveUserFromGroupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveUserFromGroupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveUserFromGroupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveUserFromGroupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveUserFromGroupResponseValidationError) ErrorName() string {
	return "RemoveUserFromGroupResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveUserFromGroupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveUserFromGroupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveUserFromGroupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveUserFromGroupResponseValidationError{}