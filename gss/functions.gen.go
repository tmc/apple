// Code generated from Apple documentation for GSS. DO NOT EDIT.

package gss

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("GSS: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("GSS: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("GSS: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("GSS: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _gSSCreateCredentialFromUUID func(uuid corefoundation.CFUUIDRef) unsafe.Pointer
var _gSSCreateCredentialFromUUIDErr error

func tryGSSCreateCredentialFromUUID(uuid corefoundation.CFUUIDRef) (unsafe.Pointer, error) {
	if _gSSCreateCredentialFromUUID == nil {
		return nil, symbolCallError("GSSCreateCredentialFromUUID", "10.9", _gSSCreateCredentialFromUUIDErr)
	}
	return _gSSCreateCredentialFromUUID(uuid), nil
}

// GSSCreateCredentialFromUUID creates a credential from a universally unique identifier.
//
// See: https://developer.apple.com/documentation/GSS/GSSCreateCredentialFromUUID(_:)
func GSSCreateCredentialFromUUID(uuid corefoundation.CFUUIDRef) unsafe.Pointer {
	result, callErr := tryGSSCreateCredentialFromUUID(uuid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gSSCreateError func(mech unsafe.Pointer, major_status OM_uint32, minor_status OM_uint32) corefoundation.CFErrorRef
var _gSSCreateErrorErr error

func tryGSSCreateError(mech unsafe.Pointer, major_status OM_uint32, minor_status OM_uint32) (corefoundation.CFErrorRef, error) {
	if _gSSCreateError == nil {
		return *new(corefoundation.CFErrorRef), symbolCallError("GSSCreateError", "10.10", _gSSCreateErrorErr)
	}
	return _gSSCreateError(mech, major_status, minor_status), nil
}

// GSSCreateError returns an error object based on GSS-API major and minor status codes.
//
// See: https://developer.apple.com/documentation/GSS/GSSCreateError(_:_:_:)
func GSSCreateError(mech unsafe.Pointer, major_status OM_uint32, minor_status OM_uint32) corefoundation.CFErrorRef {
	result, callErr := tryGSSCreateError(mech, major_status, minor_status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gSSCreateName func(name corefoundation.CFTypeRef, name_type unsafe.Pointer, err *corefoundation.CFErrorRef) unsafe.Pointer
var _gSSCreateNameErr error

func tryGSSCreateName(name corefoundation.CFTypeRef, name_type unsafe.Pointer, err *corefoundation.CFErrorRef) (unsafe.Pointer, error) {
	if _gSSCreateName == nil {
		return nil, symbolCallError("GSSCreateName", "10.9", _gSSCreateNameErr)
	}
	return _gSSCreateName(name, name_type, err), nil
}

// GSSCreateName returns a GSS name given a buffer and a type.
//
// See: https://developer.apple.com/documentation/GSS/GSSCreateName(_:_:_:)
func GSSCreateName(name corefoundation.CFTypeRef, name_type unsafe.Pointer, err *corefoundation.CFErrorRef) unsafe.Pointer {
	result, callErr := tryGSSCreateName(name, name_type, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gSSCredentialCopyName func(cred unsafe.Pointer) unsafe.Pointer
var _gSSCredentialCopyNameErr error

func tryGSSCredentialCopyName(cred unsafe.Pointer) (unsafe.Pointer, error) {
	if _gSSCredentialCopyName == nil {
		return nil, symbolCallError("GSSCredentialCopyName", "10.9", _gSSCredentialCopyNameErr)
	}
	return _gSSCredentialCopyName(cred), nil
}

// GSSCredentialCopyName returns the name describing the credential.
//
// See: https://developer.apple.com/documentation/GSS/GSSCredentialCopyName(_:)
func GSSCredentialCopyName(cred unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryGSSCredentialCopyName(cred)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gSSCredentialCopyUUID func(credential unsafe.Pointer) corefoundation.CFUUID
var _gSSCredentialCopyUUIDErr error

func tryGSSCredentialCopyUUID(credential unsafe.Pointer) (corefoundation.CFUUID, error) {
	if _gSSCredentialCopyUUID == nil {
		return *new(corefoundation.CFUUID), symbolCallError("GSSCredentialCopyUUID", "10.9", _gSSCredentialCopyUUIDErr)
	}
	return _gSSCredentialCopyUUID(credential), nil
}

// GSSCredentialCopyUUID returns a copy of the universally unique identifier corresponding to a GSS credential.
//
// See: https://developer.apple.com/documentation/GSS/GSSCredentialCopyUUID(_:)
func GSSCredentialCopyUUID(credential unsafe.Pointer) corefoundation.CFUUID {
	result, callErr := tryGSSCredentialCopyUUID(credential)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gSSCredentialGetLifetime func(cred unsafe.Pointer) OM_uint32
var _gSSCredentialGetLifetimeErr error

func tryGSSCredentialGetLifetime(cred unsafe.Pointer) (OM_uint32, error) {
	if _gSSCredentialGetLifetime == nil {
		return *new(OM_uint32), symbolCallError("GSSCredentialGetLifetime", "10.9", _gSSCredentialGetLifetimeErr)
	}
	return _gSSCredentialGetLifetime(cred), nil
}

// GSSCredentialGetLifetime returns the remaining time in seconds before the credential expires.
//
// See: https://developer.apple.com/documentation/GSS/GSSCredentialGetLifetime(_:)
func GSSCredentialGetLifetime(cred unsafe.Pointer) OM_uint32 {
	result, callErr := tryGSSCredentialGetLifetime(cred)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gSSNameCreateDisplayString func(name unsafe.Pointer) corefoundation.CFStringRef
var _gSSNameCreateDisplayStringErr error

func tryGSSNameCreateDisplayString(name unsafe.Pointer) (corefoundation.CFStringRef, error) {
	if _gSSNameCreateDisplayString == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("GSSNameCreateDisplayString", "10.9", _gSSNameCreateDisplayStringErr)
	}
	return _gSSNameCreateDisplayString(name), nil
}

// GSSNameCreateDisplayString returns a string suitable for displaying to the user from a GSS name.
//
// See: https://developer.apple.com/documentation/GSS/GSSNameCreateDisplayString(_:)
func GSSNameCreateDisplayString(name unsafe.Pointer) corefoundation.CFStringRef {
	result, callErr := tryGSSNameCreateDisplayString(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_aapl_change_password func(name unsafe.Pointer, mech unsafe.Pointer, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) OM_uint32
var _gss_aapl_change_passwordErr error

func tryGss_aapl_change_password(name unsafe.Pointer, mech unsafe.Pointer, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (OM_uint32, error) {
	if _gss_aapl_change_password == nil {
		return *new(OM_uint32), symbolCallError("gss_aapl_change_password", "10.9", _gss_aapl_change_passwordErr)
	}
	return _gss_aapl_change_password(name, mech, attributes, err), nil
}

// Gss_aapl_change_password changes the password associated with a name.
//
// See: https://developer.apple.com/documentation/GSS/gss_aapl_change_password(_:_:_:_:)
func Gss_aapl_change_password(name unsafe.Pointer, mech unsafe.Pointer, attributes corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) OM_uint32 {
	result, callErr := tryGss_aapl_change_password(name, mech, attributes, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_aapl_initial_cred func(desired_name unsafe.Pointer, desired_mech unsafe.Pointer, attributes corefoundation.CFDictionaryRef, output_cred_handle unsafe.Pointer, err *corefoundation.CFErrorRef) OM_uint32
var _gss_aapl_initial_credErr error

func tryGss_aapl_initial_cred(desired_name unsafe.Pointer, desired_mech unsafe.Pointer, attributes corefoundation.CFDictionaryRef, output_cred_handle unsafe.Pointer, err *corefoundation.CFErrorRef) (OM_uint32, error) {
	if _gss_aapl_initial_cred == nil {
		return *new(OM_uint32), symbolCallError("gss_aapl_initial_cred", "10.7", _gss_aapl_initial_credErr)
	}
	return _gss_aapl_initial_cred(desired_name, desired_mech, attributes, output_cred_handle, err), nil
}

// Gss_aapl_initial_cred acquires a new credential using a password or certificate.
//
// See: https://developer.apple.com/documentation/GSS/gss_aapl_initial_cred(_:_:_:_:_:)
func Gss_aapl_initial_cred(desired_name unsafe.Pointer, desired_mech unsafe.Pointer, attributes corefoundation.CFDictionaryRef, output_cred_handle unsafe.Pointer, err *corefoundation.CFErrorRef) OM_uint32 {
	result, callErr := tryGss_aapl_initial_cred(desired_name, desired_mech, attributes, output_cred_handle, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_accept_sec_context func(minor_status *OM_uint32, context_handle unsafe.Pointer, acceptor_cred_handle unsafe.Pointer, input_token unsafe.Pointer, input_chan_bindings Gss_channel_bindings_t, src_name unsafe.Pointer, mech_type unsafe.Pointer, output_token unsafe.Pointer, ret_flags *OM_uint32, time_rec *OM_uint32, delegated_cred_handle unsafe.Pointer) OM_uint32
var _gss_accept_sec_contextErr error

func tryGss_accept_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, acceptor_cred_handle unsafe.Pointer, input_token unsafe.Pointer, input_chan_bindings Gss_channel_bindings_t, src_name unsafe.Pointer, mech_type unsafe.Pointer, output_token unsafe.Pointer, ret_flags *OM_uint32, time_rec *OM_uint32, delegated_cred_handle unsafe.Pointer) (OM_uint32, error) {
	if _gss_accept_sec_context == nil {
		return *new(OM_uint32), symbolCallError("gss_accept_sec_context", "10.7", _gss_accept_sec_contextErr)
	}
	return _gss_accept_sec_context(minor_status, context_handle, acceptor_cred_handle, input_token, input_chan_bindings, src_name, mech_type, output_token, ret_flags, time_rec, delegated_cred_handle), nil
}

// Gss_accept_sec_context accepts a security context initiated by a peer.
//
// See: https://developer.apple.com/documentation/GSS/gss_accept_sec_context(_:_:_:_:_:_:_:_:_:_:_:)
func Gss_accept_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, acceptor_cred_handle unsafe.Pointer, input_token unsafe.Pointer, input_chan_bindings Gss_channel_bindings_t, src_name unsafe.Pointer, mech_type unsafe.Pointer, output_token unsafe.Pointer, ret_flags *OM_uint32, time_rec *OM_uint32, delegated_cred_handle unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_accept_sec_context(minor_status, context_handle, acceptor_cred_handle, input_token, input_chan_bindings, src_name, mech_type, output_token, ret_flags, time_rec, delegated_cred_handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_acquire_cred func(minor_status *OM_uint32, desired_name unsafe.Pointer, time_req OM_uint32, desired_mechs unsafe.Pointer, cred_usage Gss_cred_usage_t, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, time_rec *OM_uint32) OM_uint32
var _gss_acquire_credErr error

func tryGss_acquire_cred(minor_status *OM_uint32, desired_name unsafe.Pointer, time_req OM_uint32, desired_mechs unsafe.Pointer, cred_usage Gss_cred_usage_t, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, time_rec *OM_uint32) (OM_uint32, error) {
	if _gss_acquire_cred == nil {
		return *new(OM_uint32), symbolCallError("gss_acquire_cred", "10.7", _gss_acquire_credErr)
	}
	return _gss_acquire_cred(minor_status, desired_name, time_req, desired_mechs, cred_usage, output_cred_handle, actual_mechs, time_rec), nil
}

// Gss_acquire_cred acquires a credential for use in establishing a security context.
//
// See: https://developer.apple.com/documentation/GSS/gss_acquire_cred(_:_:_:_:_:_:_:_:)
func Gss_acquire_cred(minor_status *OM_uint32, desired_name unsafe.Pointer, time_req OM_uint32, desired_mechs unsafe.Pointer, cred_usage Gss_cred_usage_t, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, time_rec *OM_uint32) OM_uint32 {
	result, callErr := tryGss_acquire_cred(minor_status, desired_name, time_req, desired_mechs, cred_usage, output_cred_handle, actual_mechs, time_rec)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_acquire_cred_with_password func(minor_status *OM_uint32, desired_name unsafe.Pointer, password unsafe.Pointer, time_req OM_uint32, desired_mechs unsafe.Pointer, cred_usage Gss_cred_usage_t, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, time_rec *OM_uint32) OM_uint32
var _gss_acquire_cred_with_passwordErr error

func tryGss_acquire_cred_with_password(minor_status *OM_uint32, desired_name unsafe.Pointer, password unsafe.Pointer, time_req OM_uint32, desired_mechs unsafe.Pointer, cred_usage Gss_cred_usage_t, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, time_rec *OM_uint32) (OM_uint32, error) {
	if _gss_acquire_cred_with_password == nil {
		return *new(OM_uint32), symbolCallError("gss_acquire_cred_with_password", "10.7", _gss_acquire_cred_with_passwordErr)
	}
	return _gss_acquire_cred_with_password(minor_status, desired_name, password, time_req, desired_mechs, cred_usage, output_cred_handle, actual_mechs, time_rec), nil
}

// Gss_acquire_cred_with_password acquires a credential for use in establishing a security context using a password.
//
// See: https://developer.apple.com/documentation/GSS/gss_acquire_cred_with_password(_:_:_:_:_:_:_:_:_:)
func Gss_acquire_cred_with_password(minor_status *OM_uint32, desired_name unsafe.Pointer, password unsafe.Pointer, time_req OM_uint32, desired_mechs unsafe.Pointer, cred_usage Gss_cred_usage_t, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, time_rec *OM_uint32) OM_uint32 {
	result, callErr := tryGss_acquire_cred_with_password(minor_status, desired_name, password, time_req, desired_mechs, cred_usage, output_cred_handle, actual_mechs, time_rec)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_add_buffer_set_member func(minor_status *OM_uint32, member_buffer unsafe.Pointer, buffer_set unsafe.Pointer) OM_uint32
var _gss_add_buffer_set_memberErr error

func tryGss_add_buffer_set_member(minor_status *OM_uint32, member_buffer unsafe.Pointer, buffer_set unsafe.Pointer) (OM_uint32, error) {
	if _gss_add_buffer_set_member == nil {
		return *new(OM_uint32), symbolCallError("gss_add_buffer_set_member", "10.7", _gss_add_buffer_set_memberErr)
	}
	return _gss_add_buffer_set_member(minor_status, member_buffer, buffer_set), nil
}

// Gss_add_buffer_set_member copies the contents of a buffer into a buffer set.
//
// See: https://developer.apple.com/documentation/GSS/gss_add_buffer_set_member(_:_:_:)
func Gss_add_buffer_set_member(minor_status *OM_uint32, member_buffer unsafe.Pointer, buffer_set unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_add_buffer_set_member(minor_status, member_buffer, buffer_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_add_cred func(minor_status *OM_uint32, input_cred_handle unsafe.Pointer, desired_name unsafe.Pointer, desired_mech unsafe.Pointer, cred_usage Gss_cred_usage_t, initiator_time_req OM_uint32, acceptor_time_req OM_uint32, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, initiator_time_rec *OM_uint32, acceptor_time_rec *OM_uint32) OM_uint32
var _gss_add_credErr error

func tryGss_add_cred(minor_status *OM_uint32, input_cred_handle unsafe.Pointer, desired_name unsafe.Pointer, desired_mech unsafe.Pointer, cred_usage Gss_cred_usage_t, initiator_time_req OM_uint32, acceptor_time_req OM_uint32, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, initiator_time_rec *OM_uint32, acceptor_time_rec *OM_uint32) (OM_uint32, error) {
	if _gss_add_cred == nil {
		return *new(OM_uint32), symbolCallError("gss_add_cred", "10.7", _gss_add_credErr)
	}
	return _gss_add_cred(minor_status, input_cred_handle, desired_name, desired_mech, cred_usage, initiator_time_req, acceptor_time_req, output_cred_handle, actual_mechs, initiator_time_rec, acceptor_time_rec), nil
}

// Gss_add_cred adds a new credential element to an existing credential.
//
// See: https://developer.apple.com/documentation/GSS/gss_add_cred(_:_:_:_:_:_:_:_:_:_:_:)
func Gss_add_cred(minor_status *OM_uint32, input_cred_handle unsafe.Pointer, desired_name unsafe.Pointer, desired_mech unsafe.Pointer, cred_usage Gss_cred_usage_t, initiator_time_req OM_uint32, acceptor_time_req OM_uint32, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, initiator_time_rec *OM_uint32, acceptor_time_rec *OM_uint32) OM_uint32 {
	result, callErr := tryGss_add_cred(minor_status, input_cred_handle, desired_name, desired_mech, cred_usage, initiator_time_req, acceptor_time_req, output_cred_handle, actual_mechs, initiator_time_rec, acceptor_time_rec)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_add_oid_set_member func(minor_status *OM_uint32, member_oid unsafe.Pointer, oid_set unsafe.Pointer) OM_uint32
var _gss_add_oid_set_memberErr error

func tryGss_add_oid_set_member(minor_status *OM_uint32, member_oid unsafe.Pointer, oid_set unsafe.Pointer) (OM_uint32, error) {
	if _gss_add_oid_set_member == nil {
		return *new(OM_uint32), symbolCallError("gss_add_oid_set_member", "10.7", _gss_add_oid_set_memberErr)
	}
	return _gss_add_oid_set_member(minor_status, member_oid, oid_set), nil
}

// Gss_add_oid_set_member adds an object identifier into an OID set.
//
// See: https://developer.apple.com/documentation/GSS/gss_add_oid_set_member(_:_:_:)
func Gss_add_oid_set_member(minor_status *OM_uint32, member_oid unsafe.Pointer, oid_set unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_add_oid_set_member(minor_status, member_oid, oid_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_canonicalize_name func(minor_status *OM_uint32, input_name unsafe.Pointer, mech_type unsafe.Pointer, output_name unsafe.Pointer) OM_uint32
var _gss_canonicalize_nameErr error

func tryGss_canonicalize_name(minor_status *OM_uint32, input_name unsafe.Pointer, mech_type unsafe.Pointer, output_name unsafe.Pointer) (OM_uint32, error) {
	if _gss_canonicalize_name == nil {
		return *new(OM_uint32), symbolCallError("gss_canonicalize_name", "10.7", _gss_canonicalize_nameErr)
	}
	return _gss_canonicalize_name(minor_status, input_name, mech_type, output_name), nil
}

// Gss_canonicalize_name converts an internal name into a mechanism name.
//
// See: https://developer.apple.com/documentation/GSS/gss_canonicalize_name(_:_:_:_:)
func Gss_canonicalize_name(minor_status *OM_uint32, input_name unsafe.Pointer, mech_type unsafe.Pointer, output_name unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_canonicalize_name(minor_status, input_name, mech_type, output_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_compare_name func(minor_status *OM_uint32, name1_arg unsafe.Pointer, name2_arg unsafe.Pointer, name_equal *int32) OM_uint32
var _gss_compare_nameErr error

func tryGss_compare_name(minor_status *OM_uint32, name1_arg unsafe.Pointer, name2_arg unsafe.Pointer, name_equal *int32) (OM_uint32, error) {
	if _gss_compare_name == nil {
		return *new(OM_uint32), symbolCallError("gss_compare_name", "10.7", _gss_compare_nameErr)
	}
	return _gss_compare_name(minor_status, name1_arg, name2_arg, name_equal), nil
}

// Gss_compare_name returns a flag that indicates if two names in internal name format refer to the same entity.
//
// See: https://developer.apple.com/documentation/GSS/gss_compare_name(_:_:_:_:)
func Gss_compare_name(minor_status *OM_uint32, name1_arg unsafe.Pointer, name2_arg unsafe.Pointer, name_equal *int32) OM_uint32 {
	result, callErr := tryGss_compare_name(minor_status, name1_arg, name2_arg, name_equal)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_context_time func(minor_status *OM_uint32, context_handle unsafe.Pointer, time_rec *OM_uint32) OM_uint32
var _gss_context_timeErr error

func tryGss_context_time(minor_status *OM_uint32, context_handle unsafe.Pointer, time_rec *OM_uint32) (OM_uint32, error) {
	if _gss_context_time == nil {
		return *new(OM_uint32), symbolCallError("gss_context_time", "10.7", _gss_context_timeErr)
	}
	return _gss_context_time(minor_status, context_handle, time_rec), nil
}

// Gss_context_time returns the amount of time remaining before a context expires.
//
// See: https://developer.apple.com/documentation/GSS/gss_context_time(_:_:_:)
func Gss_context_time(minor_status *OM_uint32, context_handle unsafe.Pointer, time_rec *OM_uint32) OM_uint32 {
	result, callErr := tryGss_context_time(minor_status, context_handle, time_rec)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_create_empty_buffer_set func(minor_status *OM_uint32, buffer_set unsafe.Pointer) OM_uint32
var _gss_create_empty_buffer_setErr error

func tryGss_create_empty_buffer_set(minor_status *OM_uint32, buffer_set unsafe.Pointer) (OM_uint32, error) {
	if _gss_create_empty_buffer_set == nil {
		return *new(OM_uint32), symbolCallError("gss_create_empty_buffer_set", "10.7", _gss_create_empty_buffer_setErr)
	}
	return _gss_create_empty_buffer_set(minor_status, buffer_set), nil
}

// Gss_create_empty_buffer_set allocates an empty buffer set descriptor that you use to manage an array of buffers.
//
// See: https://developer.apple.com/documentation/GSS/gss_create_empty_buffer_set(_:_:)
func Gss_create_empty_buffer_set(minor_status *OM_uint32, buffer_set unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_create_empty_buffer_set(minor_status, buffer_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_create_empty_oid_set func(minor_status *OM_uint32, oid_set unsafe.Pointer) OM_uint32
var _gss_create_empty_oid_setErr error

func tryGss_create_empty_oid_set(minor_status *OM_uint32, oid_set unsafe.Pointer) (OM_uint32, error) {
	if _gss_create_empty_oid_set == nil {
		return *new(OM_uint32), symbolCallError("gss_create_empty_oid_set", "10.7", _gss_create_empty_oid_setErr)
	}
	return _gss_create_empty_oid_set(minor_status, oid_set), nil
}

// Gss_create_empty_oid_set allocates a new, empty set to hold object identifiers.
//
// See: https://developer.apple.com/documentation/GSS/gss_create_empty_oid_set(_:_:)
func Gss_create_empty_oid_set(minor_status *OM_uint32, oid_set unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_create_empty_oid_set(minor_status, oid_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_decapsulate_token func(input_token unsafe.Pointer, oid unsafe.Pointer, output_token unsafe.Pointer) OM_uint32
var _gss_decapsulate_tokenErr error

func tryGss_decapsulate_token(input_token unsafe.Pointer, oid unsafe.Pointer, output_token unsafe.Pointer) (OM_uint32, error) {
	if _gss_decapsulate_token == nil {
		return *new(OM_uint32), symbolCallError("gss_decapsulate_token", "10.7", _gss_decapsulate_tokenErr)
	}
	return _gss_decapsulate_token(input_token, oid, output_token), nil
}

// Gss_decapsulate_token returns a token encapsulated in a buffer.
//
// See: https://developer.apple.com/documentation/GSS/gss_decapsulate_token(_:_:_:)
func Gss_decapsulate_token(input_token unsafe.Pointer, oid unsafe.Pointer, output_token unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_decapsulate_token(input_token, oid, output_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_delete_sec_context func(minor_status *OM_uint32, context_handle unsafe.Pointer, output_token unsafe.Pointer) OM_uint32
var _gss_delete_sec_contextErr error

func tryGss_delete_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, output_token unsafe.Pointer) (OM_uint32, error) {
	if _gss_delete_sec_context == nil {
		return *new(OM_uint32), symbolCallError("gss_delete_sec_context", "10.7", _gss_delete_sec_contextErr)
	}
	return _gss_delete_sec_context(minor_status, context_handle, output_token), nil
}

// Gss_delete_sec_context deletes a security context.
//
// See: https://developer.apple.com/documentation/GSS/gss_delete_sec_context(_:_:_:)
func Gss_delete_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, output_token unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_delete_sec_context(minor_status, context_handle, output_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_destroy_cred func(min_stat *OM_uint32, cred_handle unsafe.Pointer) OM_uint32
var _gss_destroy_credErr error

func tryGss_destroy_cred(min_stat *OM_uint32, cred_handle unsafe.Pointer) (OM_uint32, error) {
	if _gss_destroy_cred == nil {
		return *new(OM_uint32), symbolCallError("gss_destroy_cred", "10.7", _gss_destroy_credErr)
	}
	return _gss_destroy_cred(min_stat, cred_handle), nil
}

// Gss_destroy_cred purges a credential from memory.
//
// See: https://developer.apple.com/documentation/GSS/gss_destroy_cred(_:_:)
func Gss_destroy_cred(min_stat *OM_uint32, cred_handle unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_destroy_cred(min_stat, cred_handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_display_mech_attr func(minor_status *OM_uint32, mech_attr unsafe.Pointer, name unsafe.Pointer, short_desc unsafe.Pointer, long_desc unsafe.Pointer) OM_uint32
var _gss_display_mech_attrErr error

func tryGss_display_mech_attr(minor_status *OM_uint32, mech_attr unsafe.Pointer, name unsafe.Pointer, short_desc unsafe.Pointer, long_desc unsafe.Pointer) (OM_uint32, error) {
	if _gss_display_mech_attr == nil {
		return *new(OM_uint32), symbolCallError("gss_display_mech_attr", "10.7", _gss_display_mech_attrErr)
	}
	return _gss_display_mech_attr(minor_status, mech_attr, name, short_desc, long_desc), nil
}

// Gss_display_mech_attr returns a human-readable name and description of a mechanism attribute.
//
// See: https://developer.apple.com/documentation/GSS/gss_display_mech_attr(_:_:_:_:_:)
func Gss_display_mech_attr(minor_status *OM_uint32, mech_attr unsafe.Pointer, name unsafe.Pointer, short_desc unsafe.Pointer, long_desc unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_display_mech_attr(minor_status, mech_attr, name, short_desc, long_desc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_display_name func(minor_status *OM_uint32, input_name unsafe.Pointer, output_name_buffer unsafe.Pointer, output_name_type unsafe.Pointer) OM_uint32
var _gss_display_nameErr error

func tryGss_display_name(minor_status *OM_uint32, input_name unsafe.Pointer, output_name_buffer unsafe.Pointer, output_name_type unsafe.Pointer) (OM_uint32, error) {
	if _gss_display_name == nil {
		return *new(OM_uint32), symbolCallError("gss_display_name", "10.7", _gss_display_nameErr)
	}
	return _gss_display_name(minor_status, input_name, output_name_buffer, output_name_type), nil
}

// Gss_display_name converts a name in the internal format to an octet string and the associated name type.
//
// See: https://developer.apple.com/documentation/GSS/gss_display_name(_:_:_:_:)
func Gss_display_name(minor_status *OM_uint32, input_name unsafe.Pointer, output_name_buffer unsafe.Pointer, output_name_type unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_display_name(minor_status, input_name, output_name_buffer, output_name_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_display_status func(minor_status *OM_uint32, status_value OM_uint32, status_type int32, mech_type unsafe.Pointer, message_content *OM_uint32, status_string unsafe.Pointer) OM_uint32
var _gss_display_statusErr error

func tryGss_display_status(minor_status *OM_uint32, status_value OM_uint32, status_type int32, mech_type unsafe.Pointer, message_content *OM_uint32, status_string unsafe.Pointer) (OM_uint32, error) {
	if _gss_display_status == nil {
		return *new(OM_uint32), symbolCallError("gss_display_status", "10.7", _gss_display_statusErr)
	}
	return _gss_display_status(minor_status, status_value, status_type, mech_type, message_content, status_string), nil
}

// Gss_display_status returns a human readable string for a status code.
//
// See: https://developer.apple.com/documentation/GSS/gss_display_status(_:_:_:_:_:_:)
func Gss_display_status(minor_status *OM_uint32, status_value OM_uint32, status_type int32, mech_type unsafe.Pointer, message_content *OM_uint32, status_string unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_display_status(minor_status, status_value, status_type, mech_type, message_content, status_string)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_duplicate_name func(minor_status *OM_uint32, src_name unsafe.Pointer, dest_name unsafe.Pointer) OM_uint32
var _gss_duplicate_nameErr error

func tryGss_duplicate_name(minor_status *OM_uint32, src_name unsafe.Pointer, dest_name unsafe.Pointer) (OM_uint32, error) {
	if _gss_duplicate_name == nil {
		return *new(OM_uint32), symbolCallError("gss_duplicate_name", "10.7", _gss_duplicate_nameErr)
	}
	return _gss_duplicate_name(minor_status, src_name, dest_name), nil
}

// Gss_duplicate_name returns a copy of an internal name.
//
// See: https://developer.apple.com/documentation/GSS/gss_duplicate_name(_:_:_:)
func Gss_duplicate_name(minor_status *OM_uint32, src_name unsafe.Pointer, dest_name unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_duplicate_name(minor_status, src_name, dest_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_encapsulate_token func(input_token unsafe.Pointer, oid unsafe.Pointer, output_token unsafe.Pointer) OM_uint32
var _gss_encapsulate_tokenErr error

func tryGss_encapsulate_token(input_token unsafe.Pointer, oid unsafe.Pointer, output_token unsafe.Pointer) (OM_uint32, error) {
	if _gss_encapsulate_token == nil {
		return *new(OM_uint32), symbolCallError("gss_encapsulate_token", "10.7", _gss_encapsulate_tokenErr)
	}
	return _gss_encapsulate_token(input_token, oid, output_token), nil
}

// Gss_encapsulate_token returns a buffer encapsulating the given token.
//
// See: https://developer.apple.com/documentation/GSS/gss_encapsulate_token(_:_:_:)
func Gss_encapsulate_token(input_token unsafe.Pointer, oid unsafe.Pointer, output_token unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_encapsulate_token(input_token, oid, output_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_export_cred func(minor_status *OM_uint32, cred_handle unsafe.Pointer, token unsafe.Pointer) OM_uint32
var _gss_export_credErr error

func tryGss_export_cred(minor_status *OM_uint32, cred_handle unsafe.Pointer, token unsafe.Pointer) (OM_uint32, error) {
	if _gss_export_cred == nil {
		return *new(OM_uint32), symbolCallError("gss_export_cred", "10.7", _gss_export_credErr)
	}
	return _gss_export_cred(minor_status, cred_handle, token), nil
}

// Gss_export_cred exports a credential to a token.
//
// See: https://developer.apple.com/documentation/GSS/gss_export_cred(_:_:_:)
func Gss_export_cred(minor_status *OM_uint32, cred_handle unsafe.Pointer, token unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_export_cred(minor_status, cred_handle, token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_export_name func(minor_status *OM_uint32, input_name unsafe.Pointer, exported_name unsafe.Pointer) OM_uint32
var _gss_export_nameErr error

func tryGss_export_name(minor_status *OM_uint32, input_name unsafe.Pointer, exported_name unsafe.Pointer) (OM_uint32, error) {
	if _gss_export_name == nil {
		return *new(OM_uint32), symbolCallError("gss_export_name", "10.7", _gss_export_nameErr)
	}
	return _gss_export_name(minor_status, input_name, exported_name), nil
}

// Gss_export_name returns a mechanism name in contiguous octet format.
//
// See: https://developer.apple.com/documentation/GSS/gss_export_name(_:_:_:)
func Gss_export_name(minor_status *OM_uint32, input_name unsafe.Pointer, exported_name unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_export_name(minor_status, input_name, exported_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_export_sec_context func(minor_status *OM_uint32, context_handle unsafe.Pointer, interprocess_token unsafe.Pointer) OM_uint32
var _gss_export_sec_contextErr error

func tryGss_export_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, interprocess_token unsafe.Pointer) (OM_uint32, error) {
	if _gss_export_sec_context == nil {
		return *new(OM_uint32), symbolCallError("gss_export_sec_context", "10.7", _gss_export_sec_contextErr)
	}
	return _gss_export_sec_context(minor_status, context_handle, interprocess_token), nil
}

// Gss_export_sec_context transfers a security context to another process.
//
// See: https://developer.apple.com/documentation/GSS/gss_export_sec_context(_:_:_:)
func Gss_export_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, interprocess_token unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_export_sec_context(minor_status, context_handle, interprocess_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_get_mic func(minor_status *OM_uint32, context_handle unsafe.Pointer, qop_req Gss_qop_t, message_buffer unsafe.Pointer, message_token unsafe.Pointer) OM_uint32
var _gss_get_micErr error

func tryGss_get_mic(minor_status *OM_uint32, context_handle unsafe.Pointer, qop_req Gss_qop_t, message_buffer unsafe.Pointer, message_token unsafe.Pointer) (OM_uint32, error) {
	if _gss_get_mic == nil {
		return *new(OM_uint32), symbolCallError("gss_get_mic", "10.7", _gss_get_micErr)
	}
	return _gss_get_mic(minor_status, context_handle, qop_req, message_buffer, message_token), nil
}

// Gss_get_mic returns a token that contains the MIC for a message.
//
// See: https://developer.apple.com/documentation/GSS/gss_get_mic(_:_:_:_:_:)
func Gss_get_mic(minor_status *OM_uint32, context_handle unsafe.Pointer, qop_req Gss_qop_t, message_buffer unsafe.Pointer, message_token unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_get_mic(minor_status, context_handle, qop_req, message_buffer, message_token)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_import_cred func(minor_status *OM_uint32, token unsafe.Pointer, cred_handle unsafe.Pointer) OM_uint32
var _gss_import_credErr error

func tryGss_import_cred(minor_status *OM_uint32, token unsafe.Pointer, cred_handle unsafe.Pointer) (OM_uint32, error) {
	if _gss_import_cred == nil {
		return *new(OM_uint32), symbolCallError("gss_import_cred", "10.7", _gss_import_credErr)
	}
	return _gss_import_cred(minor_status, token, cred_handle), nil
}

// Gss_import_cred imports a credential from a token.
//
// See: https://developer.apple.com/documentation/GSS/gss_import_cred(_:_:_:)
func Gss_import_cred(minor_status *OM_uint32, token unsafe.Pointer, cred_handle unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_import_cred(minor_status, token, cred_handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_import_name func(minor_status *OM_uint32, input_name_buffer unsafe.Pointer, input_name_type unsafe.Pointer, output_name unsafe.Pointer) OM_uint32
var _gss_import_nameErr error

func tryGss_import_name(minor_status *OM_uint32, input_name_buffer unsafe.Pointer, input_name_type unsafe.Pointer, output_name unsafe.Pointer) (OM_uint32, error) {
	if _gss_import_name == nil {
		return *new(OM_uint32), symbolCallError("gss_import_name", "10.7", _gss_import_nameErr)
	}
	return _gss_import_name(minor_status, input_name_buffer, input_name_type, output_name), nil
}

// Gss_import_name converts a name in contiguous octet format to the internal name format.
//
// See: https://developer.apple.com/documentation/GSS/gss_import_name(_:_:_:_:)
func Gss_import_name(minor_status *OM_uint32, input_name_buffer unsafe.Pointer, input_name_type unsafe.Pointer, output_name unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_import_name(minor_status, input_name_buffer, input_name_type, output_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_import_sec_context func(minor_status *OM_uint32, interprocess_token unsafe.Pointer, context_handle unsafe.Pointer) OM_uint32
var _gss_import_sec_contextErr error

func tryGss_import_sec_context(minor_status *OM_uint32, interprocess_token unsafe.Pointer, context_handle unsafe.Pointer) (OM_uint32, error) {
	if _gss_import_sec_context == nil {
		return *new(OM_uint32), symbolCallError("gss_import_sec_context", "10.7", _gss_import_sec_contextErr)
	}
	return _gss_import_sec_context(minor_status, interprocess_token, context_handle), nil
}

// Gss_import_sec_context imports a security context from another process.
//
// See: https://developer.apple.com/documentation/GSS/gss_import_sec_context(_:_:_:)
func Gss_import_sec_context(minor_status *OM_uint32, interprocess_token unsafe.Pointer, context_handle unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_import_sec_context(minor_status, interprocess_token, context_handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_indicate_mechs func(minor_status *OM_uint32, mech_set unsafe.Pointer) OM_uint32
var _gss_indicate_mechsErr error

func tryGss_indicate_mechs(minor_status *OM_uint32, mech_set unsafe.Pointer) (OM_uint32, error) {
	if _gss_indicate_mechs == nil {
		return *new(OM_uint32), symbolCallError("gss_indicate_mechs", "10.7", _gss_indicate_mechsErr)
	}
	return _gss_indicate_mechs(minor_status, mech_set), nil
}

// Gss_indicate_mechs returns the list of supported underlying security mechanisms.
//
// See: https://developer.apple.com/documentation/GSS/gss_indicate_mechs(_:_:)
func Gss_indicate_mechs(minor_status *OM_uint32, mech_set unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_indicate_mechs(minor_status, mech_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_indicate_mechs_by_attrs func(minor_status *OM_uint32, desired_mech_attrs unsafe.Pointer, except_mech_attrs unsafe.Pointer, critical_mech_attrs unsafe.Pointer, mechs unsafe.Pointer) OM_uint32
var _gss_indicate_mechs_by_attrsErr error

func tryGss_indicate_mechs_by_attrs(minor_status *OM_uint32, desired_mech_attrs unsafe.Pointer, except_mech_attrs unsafe.Pointer, critical_mech_attrs unsafe.Pointer, mechs unsafe.Pointer) (OM_uint32, error) {
	if _gss_indicate_mechs_by_attrs == nil {
		return *new(OM_uint32), symbolCallError("gss_indicate_mechs_by_attrs", "10.10", _gss_indicate_mechs_by_attrsErr)
	}
	return _gss_indicate_mechs_by_attrs(minor_status, desired_mech_attrs, except_mech_attrs, critical_mech_attrs, mechs), nil
}

// Gss_indicate_mechs_by_attrs returns the set of mechanisms that fulfill the given criteria.
//
// See: https://developer.apple.com/documentation/GSS/gss_indicate_mechs_by_attrs(_:_:_:_:_:)
func Gss_indicate_mechs_by_attrs(minor_status *OM_uint32, desired_mech_attrs unsafe.Pointer, except_mech_attrs unsafe.Pointer, critical_mech_attrs unsafe.Pointer, mechs unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_indicate_mechs_by_attrs(minor_status, desired_mech_attrs, except_mech_attrs, critical_mech_attrs, mechs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_init_sec_context func(minor_status *OM_uint32, initiator_cred_handle unsafe.Pointer, context_handle unsafe.Pointer, target_name unsafe.Pointer, input_mech_type unsafe.Pointer, req_flags OM_uint32, time_req OM_uint32, input_chan_bindings Gss_channel_bindings_t, input_token unsafe.Pointer, actual_mech_type unsafe.Pointer, output_token unsafe.Pointer, ret_flags *OM_uint32, time_rec *OM_uint32) OM_uint32
var _gss_init_sec_contextErr error

func tryGss_init_sec_context(minor_status *OM_uint32, initiator_cred_handle unsafe.Pointer, context_handle unsafe.Pointer, target_name unsafe.Pointer, input_mech_type unsafe.Pointer, req_flags OM_uint32, time_req OM_uint32, input_chan_bindings Gss_channel_bindings_t, input_token unsafe.Pointer, actual_mech_type unsafe.Pointer, output_token unsafe.Pointer, ret_flags *OM_uint32, time_rec *OM_uint32) (OM_uint32, error) {
	if _gss_init_sec_context == nil {
		return *new(OM_uint32), symbolCallError("gss_init_sec_context", "10.7", _gss_init_sec_contextErr)
	}
	return _gss_init_sec_context(minor_status, initiator_cred_handle, context_handle, target_name, input_mech_type, req_flags, time_req, input_chan_bindings, input_token, actual_mech_type, output_token, ret_flags, time_rec), nil
}

// Gss_init_sec_context initiates a security context with a peer.
//
// See: https://developer.apple.com/documentation/GSS/gss_init_sec_context(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func Gss_init_sec_context(minor_status *OM_uint32, initiator_cred_handle unsafe.Pointer, context_handle unsafe.Pointer, target_name unsafe.Pointer, input_mech_type unsafe.Pointer, req_flags OM_uint32, time_req OM_uint32, input_chan_bindings Gss_channel_bindings_t, input_token unsafe.Pointer, actual_mech_type unsafe.Pointer, output_token unsafe.Pointer, ret_flags *OM_uint32, time_rec *OM_uint32) OM_uint32 {
	result, callErr := tryGss_init_sec_context(minor_status, initiator_cred_handle, context_handle, target_name, input_mech_type, req_flags, time_req, input_chan_bindings, input_token, actual_mech_type, output_token, ret_flags, time_rec)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_attrs_for_mech func(minor_status *OM_uint32, mech unsafe.Pointer, mech_attr unsafe.Pointer, known_mech_attrs unsafe.Pointer) OM_uint32
var _gss_inquire_attrs_for_mechErr error

func tryGss_inquire_attrs_for_mech(minor_status *OM_uint32, mech unsafe.Pointer, mech_attr unsafe.Pointer, known_mech_attrs unsafe.Pointer) (OM_uint32, error) {
	if _gss_inquire_attrs_for_mech == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_attrs_for_mech", "10.7", _gss_inquire_attrs_for_mechErr)
	}
	return _gss_inquire_attrs_for_mech(minor_status, mech, mech_attr, known_mech_attrs), nil
}

// Gss_inquire_attrs_for_mech returns the supported attributes for one or all mechanisms.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_attrs_for_mech(_:_:_:_:)
func Gss_inquire_attrs_for_mech(minor_status *OM_uint32, mech unsafe.Pointer, mech_attr unsafe.Pointer, known_mech_attrs unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_inquire_attrs_for_mech(minor_status, mech, mech_attr, known_mech_attrs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_context func(minor_status *OM_uint32, context_handle unsafe.Pointer, src_name unsafe.Pointer, targ_name unsafe.Pointer, lifetime_rec *OM_uint32, mech_type unsafe.Pointer, ctx_flags *OM_uint32, locally_initiated *int32, xopen *int32) OM_uint32
var _gss_inquire_contextErr error

func tryGss_inquire_context(minor_status *OM_uint32, context_handle unsafe.Pointer, src_name unsafe.Pointer, targ_name unsafe.Pointer, lifetime_rec *OM_uint32, mech_type unsafe.Pointer, ctx_flags *OM_uint32, locally_initiated *int32, xopen *int32) (OM_uint32, error) {
	if _gss_inquire_context == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_context", "10.7", _gss_inquire_contextErr)
	}
	return _gss_inquire_context(minor_status, context_handle, src_name, targ_name, lifetime_rec, mech_type, ctx_flags, locally_initiated, xopen), nil
}

// Gss_inquire_context returns information about a security context.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_context(_:_:_:_:_:_:_:_:_:)
func Gss_inquire_context(minor_status *OM_uint32, context_handle unsafe.Pointer, src_name unsafe.Pointer, targ_name unsafe.Pointer, lifetime_rec *OM_uint32, mech_type unsafe.Pointer, ctx_flags *OM_uint32, locally_initiated *int32, xopen *int32) OM_uint32 {
	result, callErr := tryGss_inquire_context(minor_status, context_handle, src_name, targ_name, lifetime_rec, mech_type, ctx_flags, locally_initiated, xopen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_cred func(minor_status *OM_uint32, cred_handle unsafe.Pointer, name_ret unsafe.Pointer, lifetime *OM_uint32, cred_usage *Gss_cred_usage_t, mechanisms unsafe.Pointer) OM_uint32
var _gss_inquire_credErr error

func tryGss_inquire_cred(minor_status *OM_uint32, cred_handle unsafe.Pointer, name_ret unsafe.Pointer, lifetime *OM_uint32, cred_usage *Gss_cred_usage_t, mechanisms unsafe.Pointer) (OM_uint32, error) {
	if _gss_inquire_cred == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_cred", "10.7", _gss_inquire_credErr)
	}
	return _gss_inquire_cred(minor_status, cred_handle, name_ret, lifetime, cred_usage, mechanisms), nil
}

// Gss_inquire_cred obtains information about a credential.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_cred(_:_:_:_:_:_:)
func Gss_inquire_cred(minor_status *OM_uint32, cred_handle unsafe.Pointer, name_ret unsafe.Pointer, lifetime *OM_uint32, cred_usage *Gss_cred_usage_t, mechanisms unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_inquire_cred(minor_status, cred_handle, name_ret, lifetime, cred_usage, mechanisms)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_cred_by_mech func(minor_status *OM_uint32, cred_handle unsafe.Pointer, mech_type unsafe.Pointer, cred_name unsafe.Pointer, initiator_lifetime *OM_uint32, acceptor_lifetime *OM_uint32, cred_usage *Gss_cred_usage_t) OM_uint32
var _gss_inquire_cred_by_mechErr error

func tryGss_inquire_cred_by_mech(minor_status *OM_uint32, cred_handle unsafe.Pointer, mech_type unsafe.Pointer, cred_name unsafe.Pointer, initiator_lifetime *OM_uint32, acceptor_lifetime *OM_uint32, cred_usage *Gss_cred_usage_t) (OM_uint32, error) {
	if _gss_inquire_cred_by_mech == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_cred_by_mech", "10.7", _gss_inquire_cred_by_mechErr)
	}
	return _gss_inquire_cred_by_mech(minor_status, cred_handle, mech_type, cred_name, initiator_lifetime, acceptor_lifetime, cred_usage), nil
}

// Gss_inquire_cred_by_mech obtains per-mechanism information about a credential.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_cred_by_mech(_:_:_:_:_:_:_:)
func Gss_inquire_cred_by_mech(minor_status *OM_uint32, cred_handle unsafe.Pointer, mech_type unsafe.Pointer, cred_name unsafe.Pointer, initiator_lifetime *OM_uint32, acceptor_lifetime *OM_uint32, cred_usage *Gss_cred_usage_t) OM_uint32 {
	result, callErr := tryGss_inquire_cred_by_mech(minor_status, cred_handle, mech_type, cred_name, initiator_lifetime, acceptor_lifetime, cred_usage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_cred_by_oid func(minor_status *OM_uint32, cred_handle unsafe.Pointer, desired_object unsafe.Pointer, data_set unsafe.Pointer) OM_uint32
var _gss_inquire_cred_by_oidErr error

func tryGss_inquire_cred_by_oid(minor_status *OM_uint32, cred_handle unsafe.Pointer, desired_object unsafe.Pointer, data_set unsafe.Pointer) (OM_uint32, error) {
	if _gss_inquire_cred_by_oid == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_cred_by_oid", "10.7", _gss_inquire_cred_by_oidErr)
	}
	return _gss_inquire_cred_by_oid(minor_status, cred_handle, desired_object, data_set), nil
}

// Gss_inquire_cred_by_oid inquires about a particular characteristic of a credential.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_cred_by_oid(_:_:_:_:)
func Gss_inquire_cred_by_oid(minor_status *OM_uint32, cred_handle unsafe.Pointer, desired_object unsafe.Pointer, data_set unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_inquire_cred_by_oid(minor_status, cred_handle, desired_object, data_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_mech_for_saslname func(minor_status *OM_uint32, sasl_mech_name unsafe.Pointer, mech_type unsafe.Pointer) OM_uint32
var _gss_inquire_mech_for_saslnameErr error

func tryGss_inquire_mech_for_saslname(minor_status *OM_uint32, sasl_mech_name unsafe.Pointer, mech_type unsafe.Pointer) (OM_uint32, error) {
	if _gss_inquire_mech_for_saslname == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_mech_for_saslname", "10.10", _gss_inquire_mech_for_saslnameErr)
	}
	return _gss_inquire_mech_for_saslname(minor_status, sasl_mech_name, mech_type), nil
}

// Gss_inquire_mech_for_saslname returns the GSS-API mechanism identifier for a given Simple Authentication and Security Layer (SASL) protocol name.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_mech_for_saslname(_:_:_:)
func Gss_inquire_mech_for_saslname(minor_status *OM_uint32, sasl_mech_name unsafe.Pointer, mech_type unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_inquire_mech_for_saslname(minor_status, sasl_mech_name, mech_type)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_mechs_for_name func(minor_status *OM_uint32, input_name unsafe.Pointer, mech_types unsafe.Pointer) OM_uint32
var _gss_inquire_mechs_for_nameErr error

func tryGss_inquire_mechs_for_name(minor_status *OM_uint32, input_name unsafe.Pointer, mech_types unsafe.Pointer) (OM_uint32, error) {
	if _gss_inquire_mechs_for_name == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_mechs_for_name", "10.7", _gss_inquire_mechs_for_nameErr)
	}
	return _gss_inquire_mechs_for_name(minor_status, input_name, mech_types), nil
}

// Gss_inquire_mechs_for_name returns a list of mechanisms that support a particular name type.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_mechs_for_name(_:_:_:)
func Gss_inquire_mechs_for_name(minor_status *OM_uint32, input_name unsafe.Pointer, mech_types unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_inquire_mechs_for_name(minor_status, input_name, mech_types)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_name func(minor_status *OM_uint32, input_name unsafe.Pointer, name_is_MN *int32, MN_mech unsafe.Pointer, attrs unsafe.Pointer) OM_uint32
var _gss_inquire_nameErr error

func tryGss_inquire_name(minor_status *OM_uint32, input_name unsafe.Pointer, name_is_MN *int32, MN_mech unsafe.Pointer, attrs unsafe.Pointer) (OM_uint32, error) {
	if _gss_inquire_name == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_name", "10.7", _gss_inquire_nameErr)
	}
	return _gss_inquire_name(minor_status, input_name, name_is_MN, MN_mech, attrs), nil
}

// Gss_inquire_name returns information about a name.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_name(_:_:_:_:_:)
func Gss_inquire_name(minor_status *OM_uint32, input_name unsafe.Pointer, name_is_MN *int32, MN_mech unsafe.Pointer, attrs unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_inquire_name(minor_status, input_name, name_is_MN, MN_mech, attrs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_names_for_mech func(minor_status *OM_uint32, mechanism unsafe.Pointer, name_types unsafe.Pointer) OM_uint32
var _gss_inquire_names_for_mechErr error

func tryGss_inquire_names_for_mech(minor_status *OM_uint32, mechanism unsafe.Pointer, name_types unsafe.Pointer) (OM_uint32, error) {
	if _gss_inquire_names_for_mech == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_names_for_mech", "10.7", _gss_inquire_names_for_mechErr)
	}
	return _gss_inquire_names_for_mech(minor_status, mechanism, name_types), nil
}

// Gss_inquire_names_for_mech returns a list of name types that a given mechanism supports.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_names_for_mech(_:_:_:)
func Gss_inquire_names_for_mech(minor_status *OM_uint32, mechanism unsafe.Pointer, name_types unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_inquire_names_for_mech(minor_status, mechanism, name_types)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_saslname_for_mech func(minor_status *OM_uint32, desired_mech unsafe.Pointer, sasl_mech_name unsafe.Pointer, mech_name unsafe.Pointer, mech_description unsafe.Pointer) OM_uint32
var _gss_inquire_saslname_for_mechErr error

func tryGss_inquire_saslname_for_mech(minor_status *OM_uint32, desired_mech unsafe.Pointer, sasl_mech_name unsafe.Pointer, mech_name unsafe.Pointer, mech_description unsafe.Pointer) (OM_uint32, error) {
	if _gss_inquire_saslname_for_mech == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_saslname_for_mech", "10.10", _gss_inquire_saslname_for_mechErr)
	}
	return _gss_inquire_saslname_for_mech(minor_status, desired_mech, sasl_mech_name, mech_name, mech_description), nil
}

// Gss_inquire_saslname_for_mech returns the Simple Authentication and Security Layer (SASL) protocol name for a given GSS-API mechanism.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_saslname_for_mech(_:_:_:_:_:)
func Gss_inquire_saslname_for_mech(minor_status *OM_uint32, desired_mech unsafe.Pointer, sasl_mech_name unsafe.Pointer, mech_name unsafe.Pointer, mech_description unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_inquire_saslname_for_mech(minor_status, desired_mech, sasl_mech_name, mech_name, mech_description)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_inquire_sec_context_by_oid func(minor_status *OM_uint32, context_handle unsafe.Pointer, desired_object unsafe.Pointer, data_set unsafe.Pointer) OM_uint32
var _gss_inquire_sec_context_by_oidErr error

func tryGss_inquire_sec_context_by_oid(minor_status *OM_uint32, context_handle unsafe.Pointer, desired_object unsafe.Pointer, data_set unsafe.Pointer) (OM_uint32, error) {
	if _gss_inquire_sec_context_by_oid == nil {
		return *new(OM_uint32), symbolCallError("gss_inquire_sec_context_by_oid", "10.7", _gss_inquire_sec_context_by_oidErr)
	}
	return _gss_inquire_sec_context_by_oid(minor_status, context_handle, desired_object, data_set), nil
}

// Gss_inquire_sec_context_by_oid returns information about a particular part of a context.
//
// See: https://developer.apple.com/documentation/GSS/gss_inquire_sec_context_by_oid(_:_:_:_:)
func Gss_inquire_sec_context_by_oid(minor_status *OM_uint32, context_handle unsafe.Pointer, desired_object unsafe.Pointer, data_set unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_inquire_sec_context_by_oid(minor_status, context_handle, desired_object, data_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_iter_creds func(min_stat *OM_uint32, flags OM_uint32, mech unsafe.Pointer) OM_uint32
var _gss_iter_credsErr error

func tryGss_iter_creds(min_stat *OM_uint32, flags OM_uint32, mech unsafe.Pointer) (OM_uint32, error) {
	if _gss_iter_creds == nil {
		return *new(OM_uint32), symbolCallError("gss_iter_creds", "10.7", _gss_iter_credsErr)
	}
	return _gss_iter_creds(min_stat, flags, mech), nil
}

// Gss_iter_creds iterates over all credentials.
//
// See: https://developer.apple.com/documentation/GSS/gss_iter_creds(_:_:_:_:)
func Gss_iter_creds(min_stat *OM_uint32, flags OM_uint32, mech unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_iter_creds(min_stat, flags, mech)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_iter_creds_f func(min_stat *OM_uint32, flags OM_uint32, mech unsafe.Pointer, userctx unsafe.Pointer, useriter func(unsafe.Pointer, Gss_OID, Gss_cred_id_t)) OM_uint32
var _gss_iter_creds_fErr error

func tryGss_iter_creds_f(min_stat *OM_uint32, flags OM_uint32, mech unsafe.Pointer, userctx unsafe.Pointer, useriter func(unsafe.Pointer, Gss_OID, Gss_cred_id_t)) (OM_uint32, error) {
	if _gss_iter_creds_f == nil {
		return *new(OM_uint32), symbolCallError("gss_iter_creds_f", "10.7", _gss_iter_creds_fErr)
	}
	return _gss_iter_creds_f(min_stat, flags, mech, userctx, useriter), nil
}

// Gss_iter_creds_f iterates over all credentials with a user context.
//
// See: https://developer.apple.com/documentation/GSS/gss_iter_creds_f(_:_:_:_:_:)
func Gss_iter_creds_f(min_stat *OM_uint32, flags OM_uint32, mech unsafe.Pointer, userctx unsafe.Pointer, useriter func(unsafe.Pointer, Gss_OID, Gss_cred_id_t)) OM_uint32 {
	result, callErr := tryGss_iter_creds_f(min_stat, flags, mech, userctx, useriter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_krb5_ccache_name func(minor_status *OM_uint32, name string, out_name string) OM_uint32
var _gss_krb5_ccache_nameErr error

func tryGss_krb5_ccache_name(minor_status *OM_uint32, name string, out_name string) (OM_uint32, error) {
	if _gss_krb5_ccache_name == nil {
		return *new(OM_uint32), symbolCallError("gss_krb5_ccache_name", "10.7", _gss_krb5_ccache_nameErr)
	}
	return _gss_krb5_ccache_name(minor_status, name, out_name), nil
}

// Gss_krb5_ccache_name sets the internal Kerberos 5 credential cache name.
//
// See: https://developer.apple.com/documentation/GSS/gss_krb5_ccache_name(_:_:_:)
func Gss_krb5_ccache_name(minor_status *OM_uint32, name string, out_name string) OM_uint32 {
	result, callErr := tryGss_krb5_ccache_name(minor_status, name, out_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_krb5_export_lucid_sec_context func(minor_status *OM_uint32, context_handle unsafe.Pointer, version OM_uint32, rctx unsafe.Pointer) OM_uint32
var _gss_krb5_export_lucid_sec_contextErr error

func tryGss_krb5_export_lucid_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, version OM_uint32, rctx unsafe.Pointer) (OM_uint32, error) {
	if _gss_krb5_export_lucid_sec_context == nil {
		return *new(OM_uint32), symbolCallError("gss_krb5_export_lucid_sec_context", "10.7", _gss_krb5_export_lucid_sec_contextErr)
	}
	return _gss_krb5_export_lucid_sec_context(minor_status, context_handle, version, rctx), nil
}

// Gss_krb5_export_lucid_sec_context returns a non-opaque version of the internal context information.
//
// See: https://developer.apple.com/documentation/GSS/gss_krb5_export_lucid_sec_context(_:_:_:_:)
func Gss_krb5_export_lucid_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, version OM_uint32, rctx unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_krb5_export_lucid_sec_context(minor_status, context_handle, version, rctx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_krb5_free_lucid_sec_context func(minor_status *OM_uint32, c unsafe.Pointer) OM_uint32
var _gss_krb5_free_lucid_sec_contextErr error

func tryGss_krb5_free_lucid_sec_context(minor_status *OM_uint32, c unsafe.Pointer) (OM_uint32, error) {
	if _gss_krb5_free_lucid_sec_context == nil {
		return *new(OM_uint32), symbolCallError("gss_krb5_free_lucid_sec_context", "10.7", _gss_krb5_free_lucid_sec_contextErr)
	}
	return _gss_krb5_free_lucid_sec_context(minor_status, c), nil
}

// Gss_krb5_free_lucid_sec_context frees allocated storage associated with an exported context.
//
// See: https://developer.apple.com/documentation/GSS/gss_krb5_free_lucid_sec_context(_:_:)
func Gss_krb5_free_lucid_sec_context(minor_status *OM_uint32, c unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_krb5_free_lucid_sec_context(minor_status, c)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_krb5_set_allowable_enctypes func(minor_status *OM_uint32, cred unsafe.Pointer, num_enctypes OM_uint32, enctypes *int32) OM_uint32
var _gss_krb5_set_allowable_enctypesErr error

func tryGss_krb5_set_allowable_enctypes(minor_status *OM_uint32, cred unsafe.Pointer, num_enctypes OM_uint32, enctypes *int32) (OM_uint32, error) {
	if _gss_krb5_set_allowable_enctypes == nil {
		return *new(OM_uint32), symbolCallError("gss_krb5_set_allowable_enctypes", "10.7", _gss_krb5_set_allowable_enctypesErr)
	}
	return _gss_krb5_set_allowable_enctypes(minor_status, cred, num_enctypes, enctypes), nil
}

// Gss_krb5_set_allowable_enctypes limits the keys that can be exported to the specified types.
//
// See: https://developer.apple.com/documentation/GSS/gss_krb5_set_allowable_enctypes(_:_:_:_:)
func Gss_krb5_set_allowable_enctypes(minor_status *OM_uint32, cred unsafe.Pointer, num_enctypes OM_uint32, enctypes *int32) OM_uint32 {
	result, callErr := tryGss_krb5_set_allowable_enctypes(minor_status, cred, num_enctypes, enctypes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_oid_equal func(a unsafe.Pointer, b unsafe.Pointer) int32
var _gss_oid_equalErr error

func tryGss_oid_equal(a unsafe.Pointer, b unsafe.Pointer) (int32, error) {
	if _gss_oid_equal == nil {
		return 0, symbolCallError("gss_oid_equal", "10.7", _gss_oid_equalErr)
	}
	return _gss_oid_equal(a, b), nil
}

// Gss_oid_equal returns a flag that indicates whether two object identifiers are the same.
//
// See: https://developer.apple.com/documentation/GSS/gss_oid_equal(_:_:)
func Gss_oid_equal(a unsafe.Pointer, b unsafe.Pointer) int32 {
	result, callErr := tryGss_oid_equal(a, b)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_oid_to_str func(minor_status *OM_uint32, oid unsafe.Pointer, oid_str unsafe.Pointer) OM_uint32
var _gss_oid_to_strErr error

func tryGss_oid_to_str(minor_status *OM_uint32, oid unsafe.Pointer, oid_str unsafe.Pointer) (OM_uint32, error) {
	if _gss_oid_to_str == nil {
		return *new(OM_uint32), symbolCallError("gss_oid_to_str", "10.7", _gss_oid_to_strErr)
	}
	return _gss_oid_to_str(minor_status, oid, oid_str), nil
}

// Gss_oid_to_str converts an OID object to a human-readable string.
//
// See: https://developer.apple.com/documentation/GSS/gss_oid_to_str(_:_:_:)
func Gss_oid_to_str(minor_status *OM_uint32, oid unsafe.Pointer, oid_str unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_oid_to_str(minor_status, oid, oid_str)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_process_context_token func(minor_status *OM_uint32, context_handle unsafe.Pointer, token_buffer unsafe.Pointer) OM_uint32
var _gss_process_context_tokenErr error

func tryGss_process_context_token(minor_status *OM_uint32, context_handle unsafe.Pointer, token_buffer unsafe.Pointer) (OM_uint32, error) {
	if _gss_process_context_token == nil {
		return *new(OM_uint32), symbolCallError("gss_process_context_token", "10.7", _gss_process_context_tokenErr)
	}
	return _gss_process_context_token(minor_status, context_handle, token_buffer), nil
}

// Gss_process_context_token processes a token from a peer asynchronously.
//
// See: https://developer.apple.com/documentation/GSS/gss_process_context_token(_:_:_:)
func Gss_process_context_token(minor_status *OM_uint32, context_handle unsafe.Pointer, token_buffer unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_process_context_token(minor_status, context_handle, token_buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_pseudo_random func(minor_status *OM_uint32, context unsafe.Pointer, prf_key int32, prf_in unsafe.Pointer, desired_output_len int, prf_out unsafe.Pointer) OM_uint32
var _gss_pseudo_randomErr error

func tryGss_pseudo_random(minor_status *OM_uint32, context unsafe.Pointer, prf_key int32, prf_in unsafe.Pointer, desired_output_len int, prf_out unsafe.Pointer) (OM_uint32, error) {
	if _gss_pseudo_random == nil {
		return *new(OM_uint32), symbolCallError("gss_pseudo_random", "10.7", _gss_pseudo_randomErr)
	}
	return _gss_pseudo_random(minor_status, context, prf_key, prf_in, desired_output_len, prf_out), nil
}

// Gss_pseudo_random returns a pseudo-random byte stream for keying.
//
// See: https://developer.apple.com/documentation/GSS/gss_pseudo_random(_:_:_:_:_:_:)
func Gss_pseudo_random(minor_status *OM_uint32, context unsafe.Pointer, prf_key int32, prf_in unsafe.Pointer, desired_output_len int, prf_out unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_pseudo_random(minor_status, context, prf_key, prf_in, desired_output_len, prf_out)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_release_buffer func(minor_status *OM_uint32, buffer unsafe.Pointer) OM_uint32
var _gss_release_bufferErr error

func tryGss_release_buffer(minor_status *OM_uint32, buffer unsafe.Pointer) (OM_uint32, error) {
	if _gss_release_buffer == nil {
		return *new(OM_uint32), symbolCallError("gss_release_buffer", "10.7", _gss_release_bufferErr)
	}
	return _gss_release_buffer(minor_status, buffer), nil
}

// Gss_release_buffer frees the memory associated with a single buffer descriptor.
//
// See: https://developer.apple.com/documentation/GSS/gss_release_buffer(_:_:)
func Gss_release_buffer(minor_status *OM_uint32, buffer unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_release_buffer(minor_status, buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_release_buffer_set func(minor_status *OM_uint32, buffer_set unsafe.Pointer) OM_uint32
var _gss_release_buffer_setErr error

func tryGss_release_buffer_set(minor_status *OM_uint32, buffer_set unsafe.Pointer) (OM_uint32, error) {
	if _gss_release_buffer_set == nil {
		return *new(OM_uint32), symbolCallError("gss_release_buffer_set", "10.7", _gss_release_buffer_setErr)
	}
	return _gss_release_buffer_set(minor_status, buffer_set), nil
}

// Gss_release_buffer_set frees the memory associated with a buffer set descriptor and all the buffers it contains.
//
// See: https://developer.apple.com/documentation/GSS/gss_release_buffer_set(_:_:)
func Gss_release_buffer_set(minor_status *OM_uint32, buffer_set unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_release_buffer_set(minor_status, buffer_set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_release_cred func(minor_status *OM_uint32, cred_handle unsafe.Pointer) OM_uint32
var _gss_release_credErr error

func tryGss_release_cred(minor_status *OM_uint32, cred_handle unsafe.Pointer) (OM_uint32, error) {
	if _gss_release_cred == nil {
		return *new(OM_uint32), symbolCallError("gss_release_cred", "10.7", _gss_release_credErr)
	}
	return _gss_release_cred(minor_status, cred_handle), nil
}

// Gss_release_cred releases the memory of a credential.
//
// See: https://developer.apple.com/documentation/GSS/gss_release_cred(_:_:)
func Gss_release_cred(minor_status *OM_uint32, cred_handle unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_release_cred(minor_status, cred_handle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_release_name func(minor_status *OM_uint32, input_name unsafe.Pointer) OM_uint32
var _gss_release_nameErr error

func tryGss_release_name(minor_status *OM_uint32, input_name unsafe.Pointer) (OM_uint32, error) {
	if _gss_release_name == nil {
		return *new(OM_uint32), symbolCallError("gss_release_name", "10.7", _gss_release_nameErr)
	}
	return _gss_release_name(minor_status, input_name), nil
}

// Gss_release_name frees the resources associated with a name object.
//
// See: https://developer.apple.com/documentation/GSS/gss_release_name(_:_:)
func Gss_release_name(minor_status *OM_uint32, input_name unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_release_name(minor_status, input_name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_release_oid_set func(minor_status *OM_uint32, set unsafe.Pointer) OM_uint32
var _gss_release_oid_setErr error

func tryGss_release_oid_set(minor_status *OM_uint32, set unsafe.Pointer) (OM_uint32, error) {
	if _gss_release_oid_set == nil {
		return *new(OM_uint32), symbolCallError("gss_release_oid_set", "10.7", _gss_release_oid_setErr)
	}
	return _gss_release_oid_set(minor_status, set), nil
}

// Gss_release_oid_set releases the memory associated with an OID set.
//
// See: https://developer.apple.com/documentation/GSS/gss_release_oid_set(_:_:)
func Gss_release_oid_set(minor_status *OM_uint32, set unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_release_oid_set(minor_status, set)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_set_cred_option func(minor_status *OM_uint32, cred_handle unsafe.Pointer, object unsafe.Pointer, value unsafe.Pointer) OM_uint32
var _gss_set_cred_optionErr error

func tryGss_set_cred_option(minor_status *OM_uint32, cred_handle unsafe.Pointer, object unsafe.Pointer, value unsafe.Pointer) (OM_uint32, error) {
	if _gss_set_cred_option == nil {
		return *new(OM_uint32), symbolCallError("gss_set_cred_option", "10.7", _gss_set_cred_optionErr)
	}
	return _gss_set_cred_option(minor_status, cred_handle, object, value), nil
}

// Gss_set_cred_option changes a credential option.
//
// See: https://developer.apple.com/documentation/GSS/gss_set_cred_option(_:_:_:_:)
func Gss_set_cred_option(minor_status *OM_uint32, cred_handle unsafe.Pointer, object unsafe.Pointer, value unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_set_cred_option(minor_status, cred_handle, object, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_set_sec_context_option func(minor_status *OM_uint32, context_handle unsafe.Pointer, object unsafe.Pointer, value unsafe.Pointer) OM_uint32
var _gss_set_sec_context_optionErr error

func tryGss_set_sec_context_option(minor_status *OM_uint32, context_handle unsafe.Pointer, object unsafe.Pointer, value unsafe.Pointer) (OM_uint32, error) {
	if _gss_set_sec_context_option == nil {
		return *new(OM_uint32), symbolCallError("gss_set_sec_context_option", "10.7", _gss_set_sec_context_optionErr)
	}
	return _gss_set_sec_context_option(minor_status, context_handle, object, value), nil
}

// Gss_set_sec_context_option sets an option on a context.
//
// See: https://developer.apple.com/documentation/GSS/gss_set_sec_context_option(_:_:_:_:)
func Gss_set_sec_context_option(minor_status *OM_uint32, context_handle unsafe.Pointer, object unsafe.Pointer, value unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_set_sec_context_option(minor_status, context_handle, object, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_test_oid_set_member func(minor_status *OM_uint32, member unsafe.Pointer, set unsafe.Pointer, present *int32) OM_uint32
var _gss_test_oid_set_memberErr error

func tryGss_test_oid_set_member(minor_status *OM_uint32, member unsafe.Pointer, set unsafe.Pointer, present *int32) (OM_uint32, error) {
	if _gss_test_oid_set_member == nil {
		return *new(OM_uint32), symbolCallError("gss_test_oid_set_member", "10.7", _gss_test_oid_set_memberErr)
	}
	return _gss_test_oid_set_member(minor_status, member, set, present), nil
}

// Gss_test_oid_set_member returns a flag that indicates if an OID is present in an OID set.
//
// See: https://developer.apple.com/documentation/GSS/gss_test_oid_set_member(_:_:_:_:)
func Gss_test_oid_set_member(minor_status *OM_uint32, member unsafe.Pointer, set unsafe.Pointer, present *int32) OM_uint32 {
	result, callErr := tryGss_test_oid_set_member(minor_status, member, set, present)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_unwrap func(minor_status *OM_uint32, context_handle unsafe.Pointer, input_message_buffer unsafe.Pointer, output_message_buffer unsafe.Pointer, conf_state *int32, qop_state *Gss_qop_t) OM_uint32
var _gss_unwrapErr error

func tryGss_unwrap(minor_status *OM_uint32, context_handle unsafe.Pointer, input_message_buffer unsafe.Pointer, output_message_buffer unsafe.Pointer, conf_state *int32, qop_state *Gss_qop_t) (OM_uint32, error) {
	if _gss_unwrap == nil {
		return *new(OM_uint32), symbolCallError("gss_unwrap", "10.7", _gss_unwrapErr)
	}
	return _gss_unwrap(minor_status, context_handle, input_message_buffer, output_message_buffer, conf_state, qop_state), nil
}

// Gss_unwrap returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC.
//
// See: https://developer.apple.com/documentation/GSS/gss_unwrap(_:_:_:_:_:_:)
func Gss_unwrap(minor_status *OM_uint32, context_handle unsafe.Pointer, input_message_buffer unsafe.Pointer, output_message_buffer unsafe.Pointer, conf_state *int32, qop_state *Gss_qop_t) OM_uint32 {
	result, callErr := tryGss_unwrap(minor_status, context_handle, input_message_buffer, output_message_buffer, conf_state, qop_state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_userok func(name unsafe.Pointer, user string) int32
var _gss_userokErr error

func tryGss_userok(name unsafe.Pointer, user string) (int32, error) {
	if _gss_userok == nil {
		return 0, symbolCallError("gss_userok", "10.9", _gss_userokErr)
	}
	return _gss_userok(name, user), nil
}

// Gss_userok returns a flag that indicates if a given user is authorized.
//
// See: https://developer.apple.com/documentation/GSS/gss_userok(_:_:)
func Gss_userok(name unsafe.Pointer, user string) int32 {
	result, callErr := tryGss_userok(name, user)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_verify_mic func(minor_status *OM_uint32, context_handle unsafe.Pointer, message_buffer unsafe.Pointer, token_buffer unsafe.Pointer, qop_state *Gss_qop_t) OM_uint32
var _gss_verify_micErr error

func tryGss_verify_mic(minor_status *OM_uint32, context_handle unsafe.Pointer, message_buffer unsafe.Pointer, token_buffer unsafe.Pointer, qop_state *Gss_qop_t) (OM_uint32, error) {
	if _gss_verify_mic == nil {
		return *new(OM_uint32), symbolCallError("gss_verify_mic", "10.7", _gss_verify_micErr)
	}
	return _gss_verify_mic(minor_status, context_handle, message_buffer, token_buffer, qop_state), nil
}

// Gss_verify_mic returns an indication of whether the integrity of a message is intact, given its MIC token.
//
// See: https://developer.apple.com/documentation/GSS/gss_verify_mic(_:_:_:_:_:)
func Gss_verify_mic(minor_status *OM_uint32, context_handle unsafe.Pointer, message_buffer unsafe.Pointer, token_buffer unsafe.Pointer, qop_state *Gss_qop_t) OM_uint32 {
	result, callErr := tryGss_verify_mic(minor_status, context_handle, message_buffer, token_buffer, qop_state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_wrap func(minor_status *OM_uint32, context_handle unsafe.Pointer, conf_req_flag int32, qop_req Gss_qop_t, input_message_buffer unsafe.Pointer, conf_state *int32, output_message_buffer unsafe.Pointer) OM_uint32
var _gss_wrapErr error

func tryGss_wrap(minor_status *OM_uint32, context_handle unsafe.Pointer, conf_req_flag int32, qop_req Gss_qop_t, input_message_buffer unsafe.Pointer, conf_state *int32, output_message_buffer unsafe.Pointer) (OM_uint32, error) {
	if _gss_wrap == nil {
		return *new(OM_uint32), symbolCallError("gss_wrap", "10.7", _gss_wrapErr)
	}
	return _gss_wrap(minor_status, context_handle, conf_req_flag, qop_req, input_message_buffer, conf_state, output_message_buffer), nil
}

// Gss_wrap returns a secure message created by calculating and attaching a MIC to the input message, and then optionally encrypting it.
//
// See: https://developer.apple.com/documentation/GSS/gss_wrap(_:_:_:_:_:_:_:)
func Gss_wrap(minor_status *OM_uint32, context_handle unsafe.Pointer, conf_req_flag int32, qop_req Gss_qop_t, input_message_buffer unsafe.Pointer, conf_state *int32, output_message_buffer unsafe.Pointer) OM_uint32 {
	result, callErr := tryGss_wrap(minor_status, context_handle, conf_req_flag, qop_req, input_message_buffer, conf_state, output_message_buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gss_wrap_size_limit func(minor_status *OM_uint32, context_handle unsafe.Pointer, conf_req_flag int32, qop_req Gss_qop_t, req_output_size OM_uint32, max_input_size *OM_uint32) OM_uint32
var _gss_wrap_size_limitErr error

func tryGss_wrap_size_limit(minor_status *OM_uint32, context_handle unsafe.Pointer, conf_req_flag int32, qop_req Gss_qop_t, req_output_size OM_uint32, max_input_size *OM_uint32) (OM_uint32, error) {
	if _gss_wrap_size_limit == nil {
		return *new(OM_uint32), symbolCallError("gss_wrap_size_limit", "10.7", _gss_wrap_size_limitErr)
	}
	return _gss_wrap_size_limit(minor_status, context_handle, conf_req_flag, qop_req, req_output_size, max_input_size), nil
}

// Gss_wrap_size_limit returns the largest allowable wrap size for a given set of constraints.
//
// See: https://developer.apple.com/documentation/GSS/gss_wrap_size_limit(_:_:_:_:_:_:)
func Gss_wrap_size_limit(minor_status *OM_uint32, context_handle unsafe.Pointer, conf_req_flag int32, qop_req Gss_qop_t, req_output_size OM_uint32, max_input_size *OM_uint32) OM_uint32 {
	result, callErr := tryGss_wrap_size_limit(minor_status, context_handle, conf_req_flag, qop_req, req_output_size, max_input_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gsskrb5_extract_authz_data_from_sec_context func(minor_status *OM_uint32, context_handle unsafe.Pointer, ad_type int32, ad_data unsafe.Pointer) OM_uint32
var _gsskrb5_extract_authz_data_from_sec_contextErr error

func tryGsskrb5_extract_authz_data_from_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, ad_type int32, ad_data unsafe.Pointer) (OM_uint32, error) {
	if _gsskrb5_extract_authz_data_from_sec_context == nil {
		return *new(OM_uint32), symbolCallError("gsskrb5_extract_authz_data_from_sec_context", "10.7", _gsskrb5_extract_authz_data_from_sec_contextErr)
	}
	return _gsskrb5_extract_authz_data_from_sec_context(minor_status, context_handle, ad_type, ad_data), nil
}

// Gsskrb5_extract_authz_data_from_sec_context extracts Kerberos authorization data stored within the context.
//
// See: https://developer.apple.com/documentation/GSS/gsskrb5_extract_authz_data_from_sec_context(_:_:_:_:)
func Gsskrb5_extract_authz_data_from_sec_context(minor_status *OM_uint32, context_handle unsafe.Pointer, ad_type int32, ad_data unsafe.Pointer) OM_uint32 {
	result, callErr := tryGsskrb5_extract_authz_data_from_sec_context(minor_status, context_handle, ad_type, ad_data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _gsskrb5_register_acceptor_identity func(identity string) OM_uint32
var _gsskrb5_register_acceptor_identityErr error

func tryGsskrb5_register_acceptor_identity(identity string) (OM_uint32, error) {
	if _gsskrb5_register_acceptor_identity == nil {
		return *new(OM_uint32), symbolCallError("gsskrb5_register_acceptor_identity", "10.7", _gsskrb5_register_acceptor_identityErr)
	}
	return _gsskrb5_register_acceptor_identity(identity), nil
}

// Gsskrb5_register_acceptor_identity sets the Kerberos 5 file-based key that the acceptor will use.
//
// See: https://developer.apple.com/documentation/GSS/gsskrb5_register_acceptor_identity(_:)
func Gsskrb5_register_acceptor_identity(identity string) OM_uint32 {
	result, callErr := tryGsskrb5_register_acceptor_identity(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _krb5_gss_register_acceptor_identity func(identity string) OM_uint32
var _krb5_gss_register_acceptor_identityErr error

func tryKrb5_gss_register_acceptor_identity(identity string) (OM_uint32, error) {
	if _krb5_gss_register_acceptor_identity == nil {
		return *new(OM_uint32), symbolCallError("krb5_gss_register_acceptor_identity", "10.7", _krb5_gss_register_acceptor_identityErr)
	}
	return _krb5_gss_register_acceptor_identity(identity), nil
}

// Krb5_gss_register_acceptor_identity sets the Kerberos 5 file-based key that the acceptor will use.
//
// See: https://developer.apple.com/documentation/GSS/krb5_gss_register_acceptor_identity(_:)
func Krb5_gss_register_acceptor_identity(identity string) OM_uint32 {
	result, callErr := tryKrb5_gss_register_acceptor_identity(identity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_gSSCreateCredentialFromUUID, &_gSSCreateCredentialFromUUIDErr, frameworkHandle, "GSSCreateCredentialFromUUID", "10.9")
	registerFunc(&_gSSCreateError, &_gSSCreateErrorErr, frameworkHandle, "GSSCreateError", "10.10")
	registerFunc(&_gSSCreateName, &_gSSCreateNameErr, frameworkHandle, "GSSCreateName", "10.9")
	registerFunc(&_gSSCredentialCopyName, &_gSSCredentialCopyNameErr, frameworkHandle, "GSSCredentialCopyName", "10.9")
	registerFunc(&_gSSCredentialCopyUUID, &_gSSCredentialCopyUUIDErr, frameworkHandle, "GSSCredentialCopyUUID", "10.9")
	registerFunc(&_gSSCredentialGetLifetime, &_gSSCredentialGetLifetimeErr, frameworkHandle, "GSSCredentialGetLifetime", "10.9")
	registerFunc(&_gSSNameCreateDisplayString, &_gSSNameCreateDisplayStringErr, frameworkHandle, "GSSNameCreateDisplayString", "10.9")
	registerFunc(&_gss_aapl_change_password, &_gss_aapl_change_passwordErr, frameworkHandle, "gss_aapl_change_password", "10.9")
	registerFunc(&_gss_aapl_initial_cred, &_gss_aapl_initial_credErr, frameworkHandle, "gss_aapl_initial_cred", "10.7")
	registerFunc(&_gss_accept_sec_context, &_gss_accept_sec_contextErr, frameworkHandle, "gss_accept_sec_context", "10.7")
	registerFunc(&_gss_acquire_cred, &_gss_acquire_credErr, frameworkHandle, "gss_acquire_cred", "10.7")
	registerFunc(&_gss_acquire_cred_with_password, &_gss_acquire_cred_with_passwordErr, frameworkHandle, "gss_acquire_cred_with_password", "10.7")
	registerFunc(&_gss_add_buffer_set_member, &_gss_add_buffer_set_memberErr, frameworkHandle, "gss_add_buffer_set_member", "10.7")
	registerFunc(&_gss_add_cred, &_gss_add_credErr, frameworkHandle, "gss_add_cred", "10.7")
	registerFunc(&_gss_add_oid_set_member, &_gss_add_oid_set_memberErr, frameworkHandle, "gss_add_oid_set_member", "10.7")
	registerFunc(&_gss_canonicalize_name, &_gss_canonicalize_nameErr, frameworkHandle, "gss_canonicalize_name", "10.7")
	registerFunc(&_gss_compare_name, &_gss_compare_nameErr, frameworkHandle, "gss_compare_name", "10.7")
	registerFunc(&_gss_context_time, &_gss_context_timeErr, frameworkHandle, "gss_context_time", "10.7")
	registerFunc(&_gss_create_empty_buffer_set, &_gss_create_empty_buffer_setErr, frameworkHandle, "gss_create_empty_buffer_set", "10.7")
	registerFunc(&_gss_create_empty_oid_set, &_gss_create_empty_oid_setErr, frameworkHandle, "gss_create_empty_oid_set", "10.7")
	registerFunc(&_gss_decapsulate_token, &_gss_decapsulate_tokenErr, frameworkHandle, "gss_decapsulate_token", "10.7")
	registerFunc(&_gss_delete_sec_context, &_gss_delete_sec_contextErr, frameworkHandle, "gss_delete_sec_context", "10.7")
	registerFunc(&_gss_destroy_cred, &_gss_destroy_credErr, frameworkHandle, "gss_destroy_cred", "10.7")
	registerFunc(&_gss_display_mech_attr, &_gss_display_mech_attrErr, frameworkHandle, "gss_display_mech_attr", "10.7")
	registerFunc(&_gss_display_name, &_gss_display_nameErr, frameworkHandle, "gss_display_name", "10.7")
	registerFunc(&_gss_display_status, &_gss_display_statusErr, frameworkHandle, "gss_display_status", "10.7")
	registerFunc(&_gss_duplicate_name, &_gss_duplicate_nameErr, frameworkHandle, "gss_duplicate_name", "10.7")
	registerFunc(&_gss_encapsulate_token, &_gss_encapsulate_tokenErr, frameworkHandle, "gss_encapsulate_token", "10.7")
	registerFunc(&_gss_export_cred, &_gss_export_credErr, frameworkHandle, "gss_export_cred", "10.7")
	registerFunc(&_gss_export_name, &_gss_export_nameErr, frameworkHandle, "gss_export_name", "10.7")
	registerFunc(&_gss_export_sec_context, &_gss_export_sec_contextErr, frameworkHandle, "gss_export_sec_context", "10.7")
	registerFunc(&_gss_get_mic, &_gss_get_micErr, frameworkHandle, "gss_get_mic", "10.7")
	registerFunc(&_gss_import_cred, &_gss_import_credErr, frameworkHandle, "gss_import_cred", "10.7")
	registerFunc(&_gss_import_name, &_gss_import_nameErr, frameworkHandle, "gss_import_name", "10.7")
	registerFunc(&_gss_import_sec_context, &_gss_import_sec_contextErr, frameworkHandle, "gss_import_sec_context", "10.7")
	registerFunc(&_gss_indicate_mechs, &_gss_indicate_mechsErr, frameworkHandle, "gss_indicate_mechs", "10.7")
	registerFunc(&_gss_indicate_mechs_by_attrs, &_gss_indicate_mechs_by_attrsErr, frameworkHandle, "gss_indicate_mechs_by_attrs", "10.10")
	registerFunc(&_gss_init_sec_context, &_gss_init_sec_contextErr, frameworkHandle, "gss_init_sec_context", "10.7")
	registerFunc(&_gss_inquire_attrs_for_mech, &_gss_inquire_attrs_for_mechErr, frameworkHandle, "gss_inquire_attrs_for_mech", "10.7")
	registerFunc(&_gss_inquire_context, &_gss_inquire_contextErr, frameworkHandle, "gss_inquire_context", "10.7")
	registerFunc(&_gss_inquire_cred, &_gss_inquire_credErr, frameworkHandle, "gss_inquire_cred", "10.7")
	registerFunc(&_gss_inquire_cred_by_mech, &_gss_inquire_cred_by_mechErr, frameworkHandle, "gss_inquire_cred_by_mech", "10.7")
	registerFunc(&_gss_inquire_cred_by_oid, &_gss_inquire_cred_by_oidErr, frameworkHandle, "gss_inquire_cred_by_oid", "10.7")
	registerFunc(&_gss_inquire_mech_for_saslname, &_gss_inquire_mech_for_saslnameErr, frameworkHandle, "gss_inquire_mech_for_saslname", "10.10")
	registerFunc(&_gss_inquire_mechs_for_name, &_gss_inquire_mechs_for_nameErr, frameworkHandle, "gss_inquire_mechs_for_name", "10.7")
	registerFunc(&_gss_inquire_name, &_gss_inquire_nameErr, frameworkHandle, "gss_inquire_name", "10.7")
	registerFunc(&_gss_inquire_names_for_mech, &_gss_inquire_names_for_mechErr, frameworkHandle, "gss_inquire_names_for_mech", "10.7")
	registerFunc(&_gss_inquire_saslname_for_mech, &_gss_inquire_saslname_for_mechErr, frameworkHandle, "gss_inquire_saslname_for_mech", "10.10")
	registerFunc(&_gss_inquire_sec_context_by_oid, &_gss_inquire_sec_context_by_oidErr, frameworkHandle, "gss_inquire_sec_context_by_oid", "10.7")
	registerFunc(&_gss_iter_creds, &_gss_iter_credsErr, frameworkHandle, "gss_iter_creds", "10.7")
	registerFunc(&_gss_iter_creds_f, &_gss_iter_creds_fErr, frameworkHandle, "gss_iter_creds_f", "10.7")
	registerFunc(&_gss_krb5_ccache_name, &_gss_krb5_ccache_nameErr, frameworkHandle, "gss_krb5_ccache_name", "10.7")
	registerFunc(&_gss_krb5_export_lucid_sec_context, &_gss_krb5_export_lucid_sec_contextErr, frameworkHandle, "gss_krb5_export_lucid_sec_context", "10.7")
	registerFunc(&_gss_krb5_free_lucid_sec_context, &_gss_krb5_free_lucid_sec_contextErr, frameworkHandle, "gss_krb5_free_lucid_sec_context", "10.7")
	registerFunc(&_gss_krb5_set_allowable_enctypes, &_gss_krb5_set_allowable_enctypesErr, frameworkHandle, "gss_krb5_set_allowable_enctypes", "10.7")
	registerFunc(&_gss_oid_equal, &_gss_oid_equalErr, frameworkHandle, "gss_oid_equal", "10.7")
	registerFunc(&_gss_oid_to_str, &_gss_oid_to_strErr, frameworkHandle, "gss_oid_to_str", "10.7")
	registerFunc(&_gss_process_context_token, &_gss_process_context_tokenErr, frameworkHandle, "gss_process_context_token", "10.7")
	registerFunc(&_gss_pseudo_random, &_gss_pseudo_randomErr, frameworkHandle, "gss_pseudo_random", "10.7")
	registerFunc(&_gss_release_buffer, &_gss_release_bufferErr, frameworkHandle, "gss_release_buffer", "10.7")
	registerFunc(&_gss_release_buffer_set, &_gss_release_buffer_setErr, frameworkHandle, "gss_release_buffer_set", "10.7")
	registerFunc(&_gss_release_cred, &_gss_release_credErr, frameworkHandle, "gss_release_cred", "10.7")
	registerFunc(&_gss_release_name, &_gss_release_nameErr, frameworkHandle, "gss_release_name", "10.7")
	registerFunc(&_gss_release_oid_set, &_gss_release_oid_setErr, frameworkHandle, "gss_release_oid_set", "10.7")
	registerFunc(&_gss_set_cred_option, &_gss_set_cred_optionErr, frameworkHandle, "gss_set_cred_option", "10.7")
	registerFunc(&_gss_set_sec_context_option, &_gss_set_sec_context_optionErr, frameworkHandle, "gss_set_sec_context_option", "10.7")
	registerFunc(&_gss_test_oid_set_member, &_gss_test_oid_set_memberErr, frameworkHandle, "gss_test_oid_set_member", "10.7")
	registerFunc(&_gss_unwrap, &_gss_unwrapErr, frameworkHandle, "gss_unwrap", "10.7")
	registerFunc(&_gss_userok, &_gss_userokErr, frameworkHandle, "gss_userok", "10.9")
	registerFunc(&_gss_verify_mic, &_gss_verify_micErr, frameworkHandle, "gss_verify_mic", "10.7")
	registerFunc(&_gss_wrap, &_gss_wrapErr, frameworkHandle, "gss_wrap", "10.7")
	registerFunc(&_gss_wrap_size_limit, &_gss_wrap_size_limitErr, frameworkHandle, "gss_wrap_size_limit", "10.7")
	registerFunc(&_gsskrb5_extract_authz_data_from_sec_context, &_gsskrb5_extract_authz_data_from_sec_contextErr, frameworkHandle, "gsskrb5_extract_authz_data_from_sec_context", "10.7")
	registerFunc(&_gsskrb5_register_acceptor_identity, &_gsskrb5_register_acceptor_identityErr, frameworkHandle, "gsskrb5_register_acceptor_identity", "10.7")
	registerFunc(&_krb5_gss_register_acceptor_identity, &_krb5_gss_register_acceptor_identityErr, frameworkHandle, "krb5_gss_register_acceptor_identity", "10.7")
}
