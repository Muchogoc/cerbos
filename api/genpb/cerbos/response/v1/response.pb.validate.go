// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cerbos/response/v1/response.proto

package responsev1

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

	"google.golang.org/protobuf/types/known/anypb"

	effectv1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
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

	_ = effectv1.Effect(0)
)

// Validate checks the field values on CheckResourceSetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckResourceSetResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestId

	for key, val := range m.GetResourceInstances() {
		_ = val

		// no validation rules for ResourceInstances[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResourceSetResponseValidationError{
					field:  fmt.Sprintf("ResourceInstances[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResourceSetResponseValidationError{
				field:  "Meta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CheckResourceSetResponseValidationError is the validation error returned by
// CheckResourceSetResponse.Validate if the designated constraints aren't met.
type CheckResourceSetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponseValidationError) ErrorName() string {
	return "CheckResourceSetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponseValidationError{}

// Validate checks the field values on CheckResourceBatchResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckResourceBatchResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestId

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResourceBatchResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CheckResourceBatchResponseValidationError is the validation error returned
// by CheckResourceBatchResponse.Validate if the designated constraints aren't met.
type CheckResourceBatchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceBatchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceBatchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceBatchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceBatchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceBatchResponseValidationError) ErrorName() string {
	return "CheckResourceBatchResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceBatchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceBatchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceBatchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceBatchResponseValidationError{}

// Validate checks the field values on PlaygroundFailure with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PlaygroundFailure) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetErrors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundFailureValidationError{
					field:  fmt.Sprintf("Errors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PlaygroundFailureValidationError is the validation error returned by
// PlaygroundFailure.Validate if the designated constraints aren't met.
type PlaygroundFailureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundFailureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundFailureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundFailureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundFailureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundFailureValidationError) ErrorName() string {
	return "PlaygroundFailureValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundFailureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundFailure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundFailureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundFailureValidationError{}

// Validate checks the field values on PlaygroundValidateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PlaygroundValidateResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PlaygroundId

	switch m.Outcome.(type) {

	case *PlaygroundValidateResponse_Failure:

		if v, ok := interface{}(m.GetFailure()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundValidateResponseValidationError{
					field:  "Failure",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PlaygroundValidateResponse_Success:

		if v, ok := interface{}(m.GetSuccess()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundValidateResponseValidationError{
					field:  "Success",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PlaygroundValidateResponseValidationError is the validation error returned
// by PlaygroundValidateResponse.Validate if the designated constraints aren't met.
type PlaygroundValidateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundValidateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundValidateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundValidateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundValidateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundValidateResponseValidationError) ErrorName() string {
	return "PlaygroundValidateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundValidateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundValidateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundValidateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundValidateResponseValidationError{}

// Validate checks the field values on PlaygroundEvaluateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PlaygroundEvaluateResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PlaygroundId

	switch m.Outcome.(type) {

	case *PlaygroundEvaluateResponse_Failure:

		if v, ok := interface{}(m.GetFailure()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundEvaluateResponseValidationError{
					field:  "Failure",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PlaygroundEvaluateResponse_Success:

		if v, ok := interface{}(m.GetSuccess()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundEvaluateResponseValidationError{
					field:  "Success",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PlaygroundEvaluateResponseValidationError is the validation error returned
// by PlaygroundEvaluateResponse.Validate if the designated constraints aren't met.
type PlaygroundEvaluateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundEvaluateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundEvaluateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundEvaluateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundEvaluateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundEvaluateResponseValidationError) ErrorName() string {
	return "PlaygroundEvaluateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundEvaluateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundEvaluateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundEvaluateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundEvaluateResponseValidationError{}

// Validate checks the field values on AddOrUpdatePolicyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddOrUpdatePolicyResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSuccess()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddOrUpdatePolicyResponseValidationError{
				field:  "Success",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddOrUpdatePolicyResponseValidationError is the validation error returned by
// AddOrUpdatePolicyResponse.Validate if the designated constraints aren't met.
type AddOrUpdatePolicyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddOrUpdatePolicyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddOrUpdatePolicyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddOrUpdatePolicyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddOrUpdatePolicyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddOrUpdatePolicyResponseValidationError) ErrorName() string {
	return "AddOrUpdatePolicyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddOrUpdatePolicyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddOrUpdatePolicyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddOrUpdatePolicyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddOrUpdatePolicyResponseValidationError{}

// Validate checks the field values on ListAuditLogEntriesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListAuditLogEntriesResponse) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Entry.(type) {

	case *ListAuditLogEntriesResponse_AccessLogEntry:

		if v, ok := interface{}(m.GetAccessLogEntry()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAuditLogEntriesResponseValidationError{
					field:  "AccessLogEntry",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ListAuditLogEntriesResponse_DecisionLogEntry:

		if v, ok := interface{}(m.GetDecisionLogEntry()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAuditLogEntriesResponseValidationError{
					field:  "DecisionLogEntry",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListAuditLogEntriesResponseValidationError is the validation error returned
// by ListAuditLogEntriesResponse.Validate if the designated constraints
// aren't met.
type ListAuditLogEntriesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAuditLogEntriesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAuditLogEntriesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAuditLogEntriesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAuditLogEntriesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAuditLogEntriesResponseValidationError) ErrorName() string {
	return "ListAuditLogEntriesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAuditLogEntriesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAuditLogEntriesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAuditLogEntriesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAuditLogEntriesResponseValidationError{}

// Validate checks the field values on ServerInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ServerInfoResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Version

	// no validation rules for Commit

	// no validation rules for BuildDate

	return nil
}

// ServerInfoResponseValidationError is the validation error returned by
// ServerInfoResponse.Validate if the designated constraints aren't met.
type ServerInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerInfoResponseValidationError) ErrorName() string {
	return "ServerInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ServerInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerInfoResponseValidationError{}

// Validate checks the field values on CheckResourceSetResponse_ActionEffectMap
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *CheckResourceSetResponse_ActionEffectMap) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Actions

	return nil
}

// CheckResourceSetResponse_ActionEffectMapValidationError is the validation
// error returned by CheckResourceSetResponse_ActionEffectMap.Validate if the
// designated constraints aren't met.
type CheckResourceSetResponse_ActionEffectMapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponse_ActionEffectMapValidationError) ErrorName() string {
	return "CheckResourceSetResponse_ActionEffectMapValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponse_ActionEffectMapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse_ActionEffectMap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponse_ActionEffectMapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponse_ActionEffectMapValidationError{}

// Validate checks the field values on CheckResourceSetResponse_Meta with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CheckResourceSetResponse_Meta) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetResourceInstances() {
		_ = val

		// no validation rules for ResourceInstances[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResourceSetResponse_MetaValidationError{
					field:  fmt.Sprintf("ResourceInstances[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CheckResourceSetResponse_MetaValidationError is the validation error
// returned by CheckResourceSetResponse_Meta.Validate if the designated
// constraints aren't met.
type CheckResourceSetResponse_MetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponse_MetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponse_MetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponse_MetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponse_MetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponse_MetaValidationError) ErrorName() string {
	return "CheckResourceSetResponse_MetaValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponse_MetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse_Meta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponse_MetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponse_MetaValidationError{}

// Validate checks the field values on CheckResourceSetResponse_Meta_EffectMeta
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *CheckResourceSetResponse_Meta_EffectMeta) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MatchedPolicy

	return nil
}

// CheckResourceSetResponse_Meta_EffectMetaValidationError is the validation
// error returned by CheckResourceSetResponse_Meta_EffectMeta.Validate if the
// designated constraints aren't met.
type CheckResourceSetResponse_Meta_EffectMetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) ErrorName() string {
	return "CheckResourceSetResponse_Meta_EffectMetaValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponse_Meta_EffectMetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse_Meta_EffectMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponse_Meta_EffectMetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponse_Meta_EffectMetaValidationError{}

// Validate checks the field values on CheckResourceSetResponse_Meta_ActionMeta
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *CheckResourceSetResponse_Meta_ActionMeta) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetActions() {
		_ = val

		// no validation rules for Actions[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResourceSetResponse_Meta_ActionMetaValidationError{
					field:  fmt.Sprintf("Actions[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CheckResourceSetResponse_Meta_ActionMetaValidationError is the validation
// error returned by CheckResourceSetResponse_Meta_ActionMeta.Validate if the
// designated constraints aren't met.
type CheckResourceSetResponse_Meta_ActionMetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) ErrorName() string {
	return "CheckResourceSetResponse_Meta_ActionMetaValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceSetResponse_Meta_ActionMetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceSetResponse_Meta_ActionMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceSetResponse_Meta_ActionMetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceSetResponse_Meta_ActionMetaValidationError{}

// Validate checks the field values on
// CheckResourceBatchResponse_ActionEffectMap with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CheckResourceBatchResponse_ActionEffectMap) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResourceId

	// no validation rules for Actions

	return nil
}

// CheckResourceBatchResponse_ActionEffectMapValidationError is the validation
// error returned by CheckResourceBatchResponse_ActionEffectMap.Validate if
// the designated constraints aren't met.
type CheckResourceBatchResponse_ActionEffectMapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) ErrorName() string {
	return "CheckResourceBatchResponse_ActionEffectMapValidationError"
}

// Error satisfies the builtin error interface
func (e CheckResourceBatchResponse_ActionEffectMapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResourceBatchResponse_ActionEffectMap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResourceBatchResponse_ActionEffectMapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResourceBatchResponse_ActionEffectMapValidationError{}

// Validate checks the field values on PlaygroundFailure_Error with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PlaygroundFailure_Error) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for File

	// no validation rules for Error

	return nil
}

// PlaygroundFailure_ErrorValidationError is the validation error returned by
// PlaygroundFailure_Error.Validate if the designated constraints aren't met.
type PlaygroundFailure_ErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundFailure_ErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundFailure_ErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundFailure_ErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundFailure_ErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundFailure_ErrorValidationError) ErrorName() string {
	return "PlaygroundFailure_ErrorValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundFailure_ErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundFailure_Error.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundFailure_ErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundFailure_ErrorValidationError{}

// Validate checks the field values on PlaygroundEvaluateResponse_EvalResult
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *PlaygroundEvaluateResponse_EvalResult) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Action

	// no validation rules for Effect

	// no validation rules for Policy

	return nil
}

// PlaygroundEvaluateResponse_EvalResultValidationError is the validation error
// returned by PlaygroundEvaluateResponse_EvalResult.Validate if the
// designated constraints aren't met.
type PlaygroundEvaluateResponse_EvalResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundEvaluateResponse_EvalResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundEvaluateResponse_EvalResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundEvaluateResponse_EvalResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundEvaluateResponse_EvalResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundEvaluateResponse_EvalResultValidationError) ErrorName() string {
	return "PlaygroundEvaluateResponse_EvalResultValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundEvaluateResponse_EvalResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundEvaluateResponse_EvalResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundEvaluateResponse_EvalResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundEvaluateResponse_EvalResultValidationError{}

// Validate checks the field values on
// PlaygroundEvaluateResponse_EvalResultList with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PlaygroundEvaluateResponse_EvalResultList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaygroundEvaluateResponse_EvalResultListValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PlaygroundEvaluateResponse_EvalResultListValidationError is the validation
// error returned by PlaygroundEvaluateResponse_EvalResultList.Validate if the
// designated constraints aren't met.
type PlaygroundEvaluateResponse_EvalResultListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaygroundEvaluateResponse_EvalResultListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaygroundEvaluateResponse_EvalResultListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaygroundEvaluateResponse_EvalResultListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaygroundEvaluateResponse_EvalResultListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaygroundEvaluateResponse_EvalResultListValidationError) ErrorName() string {
	return "PlaygroundEvaluateResponse_EvalResultListValidationError"
}

// Error satisfies the builtin error interface
func (e PlaygroundEvaluateResponse_EvalResultListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaygroundEvaluateResponse_EvalResultList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaygroundEvaluateResponse_EvalResultListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaygroundEvaluateResponse_EvalResultListValidationError{}