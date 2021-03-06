// Code generated by protoc-gen-validate
// source: envoy/config/trace/v2/trace.proto
// DO NOT EDIT!!!

package v2

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on Tracing with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Tracing) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHttp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TracingValidationError{
				Field:  "Http",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// TracingValidationError is the validation error returned by Tracing.Validate
// if the designated constraints aren't met.
type TracingValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TracingValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracing.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TracingValidationError{}

// Validate checks the field values on LightstepConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LightstepConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCollectorCluster()) < 1 {
		return LightstepConfigValidationError{
			Field:  "CollectorCluster",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetAccessTokenFile()) < 1 {
		return LightstepConfigValidationError{
			Field:  "AccessTokenFile",
			Reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// LightstepConfigValidationError is the validation error returned by
// LightstepConfig.Validate if the designated constraints aren't met.
type LightstepConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e LightstepConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLightstepConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = LightstepConfigValidationError{}

// Validate checks the field values on ZipkinConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ZipkinConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCollectorCluster()) < 1 {
		return ZipkinConfigValidationError{
			Field:  "CollectorCluster",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetCollectorEndpoint()) < 1 {
		return ZipkinConfigValidationError{
			Field:  "CollectorEndpoint",
			Reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for TraceId_128Bit

	return nil
}

// ZipkinConfigValidationError is the validation error returned by
// ZipkinConfig.Validate if the designated constraints aren't met.
type ZipkinConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ZipkinConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sZipkinConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ZipkinConfigValidationError{}

// Validate checks the field values on DynamicOtConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DynamicOtConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetLibrary()) < 1 {
		return DynamicOtConfigValidationError{
			Field:  "Library",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DynamicOtConfigValidationError{
				Field:  "Config",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// DynamicOtConfigValidationError is the validation error returned by
// DynamicOtConfig.Validate if the designated constraints aren't met.
type DynamicOtConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e DynamicOtConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamicOtConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = DynamicOtConfigValidationError{}

// Validate checks the field values on TraceServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TraceServiceConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetGrpcService() == nil {
		return TraceServiceConfigValidationError{
			Field:  "GrpcService",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TraceServiceConfigValidationError{
				Field:  "GrpcService",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// TraceServiceConfigValidationError is the validation error returned by
// TraceServiceConfig.Validate if the designated constraints aren't met.
type TraceServiceConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TraceServiceConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTraceServiceConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TraceServiceConfigValidationError{}

// Validate checks the field values on Tracing_Http with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Tracing_Http) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return Tracing_HttpValidationError{
			Field:  "Name",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Tracing_HttpValidationError{
				Field:  "Config",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// Tracing_HttpValidationError is the validation error returned by
// Tracing_Http.Validate if the designated constraints aren't met.
type Tracing_HttpValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Tracing_HttpValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracing_Http.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Tracing_HttpValidationError{}
