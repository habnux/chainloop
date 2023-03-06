// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: controlplane/v1/cas_credentials.proto

package v1

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

// Validate checks the field values on CASCredentialsServiceGetRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CASCredentialsServiceGetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASCredentialsServiceGetRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CASCredentialsServiceGetRequestMultiError, or nil if none found.
func (m *CASCredentialsServiceGetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CASCredentialsServiceGetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _CASCredentialsServiceGetRequest_Role_InLookup[m.GetRole()]; !ok {
		err := CASCredentialsServiceGetRequestValidationError{
			field:  "Role",
			reason: "value must be in list [1 2]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CASCredentialsServiceGetRequestMultiError(errors)
	}

	return nil
}

// CASCredentialsServiceGetRequestMultiError is an error wrapping multiple
// validation errors returned by CASCredentialsServiceGetRequest.ValidateAll()
// if the designated constraints aren't met.
type CASCredentialsServiceGetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASCredentialsServiceGetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASCredentialsServiceGetRequestMultiError) AllErrors() []error { return m }

// CASCredentialsServiceGetRequestValidationError is the validation error
// returned by CASCredentialsServiceGetRequest.Validate if the designated
// constraints aren't met.
type CASCredentialsServiceGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASCredentialsServiceGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASCredentialsServiceGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASCredentialsServiceGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASCredentialsServiceGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASCredentialsServiceGetRequestValidationError) ErrorName() string {
	return "CASCredentialsServiceGetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CASCredentialsServiceGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASCredentialsServiceGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASCredentialsServiceGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASCredentialsServiceGetRequestValidationError{}

var _CASCredentialsServiceGetRequest_Role_InLookup = map[CASCredentialsServiceGetRequest_Role]struct{}{
	1: {},
	2: {},
}

// Validate checks the field values on CASCredentialsServiceGetResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CASCredentialsServiceGetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CASCredentialsServiceGetResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CASCredentialsServiceGetResponseMultiError, or nil if none found.
func (m *CASCredentialsServiceGetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CASCredentialsServiceGetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CASCredentialsServiceGetResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CASCredentialsServiceGetResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CASCredentialsServiceGetResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CASCredentialsServiceGetResponseMultiError(errors)
	}

	return nil
}

// CASCredentialsServiceGetResponseMultiError is an error wrapping multiple
// validation errors returned by
// CASCredentialsServiceGetResponse.ValidateAll() if the designated
// constraints aren't met.
type CASCredentialsServiceGetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASCredentialsServiceGetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASCredentialsServiceGetResponseMultiError) AllErrors() []error { return m }

// CASCredentialsServiceGetResponseValidationError is the validation error
// returned by CASCredentialsServiceGetResponse.Validate if the designated
// constraints aren't met.
type CASCredentialsServiceGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASCredentialsServiceGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASCredentialsServiceGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASCredentialsServiceGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASCredentialsServiceGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASCredentialsServiceGetResponseValidationError) ErrorName() string {
	return "CASCredentialsServiceGetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CASCredentialsServiceGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASCredentialsServiceGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASCredentialsServiceGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASCredentialsServiceGetResponseValidationError{}

// Validate checks the field values on CASCredentialsServiceGetResponse_Result
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CASCredentialsServiceGetResponse_Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CASCredentialsServiceGetResponse_Result with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// CASCredentialsServiceGetResponse_ResultMultiError, or nil if none found.
func (m *CASCredentialsServiceGetResponse_Result) ValidateAll() error {
	return m.validate(true)
}

func (m *CASCredentialsServiceGetResponse_Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return CASCredentialsServiceGetResponse_ResultMultiError(errors)
	}

	return nil
}

// CASCredentialsServiceGetResponse_ResultMultiError is an error wrapping
// multiple validation errors returned by
// CASCredentialsServiceGetResponse_Result.ValidateAll() if the designated
// constraints aren't met.
type CASCredentialsServiceGetResponse_ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CASCredentialsServiceGetResponse_ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CASCredentialsServiceGetResponse_ResultMultiError) AllErrors() []error { return m }

// CASCredentialsServiceGetResponse_ResultValidationError is the validation
// error returned by CASCredentialsServiceGetResponse_Result.Validate if the
// designated constraints aren't met.
type CASCredentialsServiceGetResponse_ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CASCredentialsServiceGetResponse_ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CASCredentialsServiceGetResponse_ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CASCredentialsServiceGetResponse_ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CASCredentialsServiceGetResponse_ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CASCredentialsServiceGetResponse_ResultValidationError) ErrorName() string {
	return "CASCredentialsServiceGetResponse_ResultValidationError"
}

// Error satisfies the builtin error interface
func (e CASCredentialsServiceGetResponse_ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCASCredentialsServiceGetResponse_Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CASCredentialsServiceGetResponse_ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CASCredentialsServiceGetResponse_ResultValidationError{}
