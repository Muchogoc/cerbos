// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: test/v1/test.proto

package testv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _test_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CompileTestCase with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CompileTestCase) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MainDef

	for key, val := range m.GetInputDefs() {
		_ = val

		// no validation rules for InputDefs[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CompileTestCaseValidationError{
					field:  fmt.Sprintf("InputDefs[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CompileTestCaseValidationError is the validation error returned by
// CompileTestCase.Validate if the designated constraints aren't met.
type CompileTestCaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompileTestCaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompileTestCaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompileTestCaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompileTestCaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompileTestCaseValidationError) ErrorName() string { return "CompileTestCaseValidationError" }

// Error satisfies the builtin error interface
func (e CompileTestCaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompileTestCase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompileTestCaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompileTestCaseValidationError{}