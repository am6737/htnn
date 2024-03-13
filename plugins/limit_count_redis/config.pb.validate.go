// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: plugins/limit_count_redis/config.proto

package limit_count_redis

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

	v1 "mosn.io/htnn/api/v1"
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

	_ = v1.StatusCode(0)
)

// Validate checks the field values on Rule with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Rule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Rule with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RuleMultiError, or nil if none found.
func (m *Rule) ValidateAll() error {
	return m.validate(true)
}

func (m *Rule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTimeWindow() == nil {
		err := RuleValidationError{
			field:  "TimeWindow",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetTimeWindow(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = RuleValidationError{
				field:  "TimeWindow",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gte := time.Duration(1*time.Second + 0*time.Nanosecond)

			if dur < gte {
				err := RuleValidationError{
					field:  "TimeWindow",
					reason: "value must be greater than or equal to 1s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if m.GetCount() < 1 {
		err := RuleValidationError{
			field:  "Count",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Key

	if len(errors) > 0 {
		return RuleMultiError(errors)
	}

	return nil
}

// RuleMultiError is an error wrapping multiple validation errors returned by
// Rule.ValidateAll() if the designated constraints aren't met.
type RuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RuleMultiError) AllErrors() []error { return m }

// RuleValidationError is the validation error returned by Rule.Validate if the
// designated constraints aren't met.
type RuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuleValidationError) ErrorName() string { return "RuleValidationError" }

// Error satisfies the builtin error interface
func (e RuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuleValidationError{}

// Validate checks the field values on Cluster with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Cluster) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cluster with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ClusterMultiError, or nil if none found.
func (m *Cluster) ValidateAll() error {
	return m.validate(true)
}

func (m *Cluster) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetAddresses()) < 1 {
		err := ClusterValidationError{
			field:  "Addresses",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ClusterMultiError(errors)
	}

	return nil
}

// ClusterMultiError is an error wrapping multiple validation errors returned
// by Cluster.ValidateAll() if the designated constraints aren't met.
type ClusterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClusterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClusterMultiError) AllErrors() []error { return m }

// ClusterValidationError is the validation error returned by Cluster.Validate
// if the designated constraints aren't met.
type ClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterValidationError) ErrorName() string { return "ClusterValidationError" }

// Error satisfies the builtin error interface
func (e ClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterValidationError{}

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConfigMultiError, or nil if none found.
func (m *Config) ValidateAll() error {
	return m.validate(true)
}

func (m *Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := len(m.GetRules()); l < 1 || l > 8 {
		err := ConfigValidationError{
			field:  "Rules",
			reason: "value must contain between 1 and 8 items, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for FailureModeDeny

	// no validation rules for EnableLimitQuotaHeaders

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Tls

	// no validation rules for TlsSkipVerify

	// no validation rules for StatusOnError

	// no validation rules for RateLimitedStatus

	oneofSourcePresent := false
	switch v := m.Source.(type) {
	case *Config_Address:
		if v == nil {
			err := ConfigValidationError{
				field:  "Source",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofSourcePresent = true
		// no validation rules for Address
	case *Config_Cluster:
		if v == nil {
			err := ConfigValidationError{
				field:  "Source",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofSourcePresent = true

		if all {
			switch v := interface{}(m.GetCluster()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  "Cluster",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  "Cluster",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCluster()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofSourcePresent {
		err := ConfigValidationError{
			field:  "Source",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ConfigMultiError(errors)
	}

	return nil
}

// ConfigMultiError is an error wrapping multiple validation errors returned by
// Config.ValidateAll() if the designated constraints aren't met.
type ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigMultiError) AllErrors() []error { return m }

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}
