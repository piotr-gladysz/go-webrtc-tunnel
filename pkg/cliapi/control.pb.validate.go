// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: control.proto

package cliapi

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

// Validate checks the field values on ConnectRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ConnectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConnectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConnectRequestMultiError,
// or nil if none found.
func (m *ConnectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConnectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	if len(errors) > 0 {
		return ConnectRequestMultiError(errors)
	}

	return nil
}

// ConnectRequestMultiError is an error wrapping multiple validation errors
// returned by ConnectRequest.ValidateAll() if the designated constraints
// aren't met.
type ConnectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConnectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConnectRequestMultiError) AllErrors() []error { return m }

// ConnectRequestValidationError is the validation error returned by
// ConnectRequest.Validate if the designated constraints aren't met.
type ConnectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectRequestValidationError) ErrorName() string { return "ConnectRequestValidationError" }

// Error satisfies the builtin error interface
func (e ConnectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectRequestValidationError{}

// Validate checks the field values on GetStatusResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetStatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStatusResponseMultiError, or nil if none found.
func (m *GetStatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for RetryCount

	// no validation rules for Uptime

	if len(errors) > 0 {
		return GetStatusResponseMultiError(errors)
	}

	return nil
}

// GetStatusResponseMultiError is an error wrapping multiple validation errors
// returned by GetStatusResponse.ValidateAll() if the designated constraints
// aren't met.
type GetStatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStatusResponseMultiError) AllErrors() []error { return m }

// GetStatusResponseValidationError is the validation error returned by
// GetStatusResponse.Validate if the designated constraints aren't met.
type GetStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStatusResponseValidationError) ErrorName() string {
	return "GetStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStatusResponseValidationError{}
