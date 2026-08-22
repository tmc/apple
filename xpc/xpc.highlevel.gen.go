// Code generated from Apple documentation for XPC. DO NOT EDIT.

package xpc

import (
	"bytes"
	"context"
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

// Xpc_object_t is the Go spelling of the C xpc_object_t handle.
//
// The cross-framework registry attributes xpc_object_t to this package, so
// every other framework resolves the C type to xpc.Xpc_object_t. Until this
// declaration existed that reference dangled, and each consumer silently
// degraded to unsafe.Pointer -- foundation.NSXPCCoder's encode and decode
// methods among them. The degradation left no trace to search for: the
// emitted output contained no occurrence of the missing name, so the only
// evidence was the absence of a symbol nobody knew to look for.
//
// It is an objectivec.Object rather than an opaque pointer because XPC
// objects are Objective-C objects under the bridge. That keeps the handle
// typed for callers, and it keeps a non-Go address out of a slot the Go
// garbage collector would scan.
type Xpc_object_t = objectivec.Object

// Xpc_type_t is the Go spelling of the C xpc_type_t handle, which identifies
// an XPC object's type rather than holding one.
//
// It is a uintptr rather than an objectivec.Object or an unsafe.Pointer
// because the value is an opaque descriptor and not a Go pointer: uintptr is
// the one spelling of the three that the Go garbage collector does not scan.
type Xpc_type_t uintptr

type Dictionary map[string]any

type Endpoint struct {
	raw unsafe.Pointer
}

// UUID is an xpc_uuid_t, the 16 bytes exactly as they travel on the wire.
// It is a distinct type rather than [16]byte so that encoding can tell a
// UUID from a 16-element byte array, which XPC carries as data.
type UUID [16]byte

// Unsupported is an XPC value this package has no Go representation for.
//
// It exists so that decoding is never silently lossy. XPC can carry file
// descriptors and shared memory regions, which are resources with a
// lifetime, and a Go value describing one is not one. Rather than return
// something that looks like ordinary data, decoding yields an Unsupported
// carrying the XPC type name and the description XPC itself prints.
//
// Encoding an Unsupported is an error: it records what arrived, and cannot
// reconstitute it.
type Unsupported struct {
	// Type is the XPC type name, as xpc_type_get_name reports it,
	// for example "fd" or "shmem".
	Type string

	// Description is xpc_copy_description output. It is a debugging aid
	// and nothing more; in particular a descriptor named here may already
	// be closed by the time it is read.
	Description string
}

type PeerRequirement struct {
	raw unsafe.Pointer
}

type RichError struct {
	raw      unsafe.Pointer
	message  string
	canRetry bool
	cause    error
}

func (e RichError) Error() string {
	switch {
	case e.cause != nil:
		return e.cause.Error()
	case e.message != "":
		return e.message
	default:
		return "xpc: unspecified error"
	}
}

func (e RichError) Unwrap() error {
	return e.cause
}

func (e RichError) CanRetry() bool {
	return e.canRetry
}

type Listener struct {
	raw         unsafe.Pointer
	incoming    func(IncomingSessionRequest) IncomingDecision
	createErr   error
	incomingBlk unsafe.Pointer
}

type Session struct {
	raw unsafe.Pointer
	// handlerMu guards incoming and cancellation. The shared block
	// trampolines read them from a native callback thread while the owning
	// goroutine may still be installing them, so a plain field read would be
	// a data race. Handlers are copied out under the lock and invoked
	// without it, since a handler may legitimately call back into the
	// session.
	handlerMu      sync.Mutex
	incoming       MessageHandler
	cancellation   CancellationHandler
	incomingBlk    unsafe.Pointer
	cancellationBl unsafe.Pointer
	// Lifecycle state for the inactive-only setters. active reports whether
	// the native object has been activated; fromHandle reports a session
	// constructed from a raw handle whose lifecycle is unknowable; and
	// requirementSet guards the one-installation rule. All three are advisory:
	// the native object is the authority, and these only exist to refuse API
	// misuse that would otherwise SIGTRAP the process from native code.
	active         bool
	fromHandle     bool
	requirementSet bool
}

type ReceivedMessage struct {
	raw     unsafe.Pointer
	session *Session
	// decoded is set when the message was synthesized from a Dictionary
	// rather than received from XPC; such a message has no raw backing
	// object.
	decoded Dictionary
	// memo caches the one decode of a raw-backed message. It is a pointer so
	// that every copy of the value shares it: ReceivedMessage has value
	// methods and is copied freely, and the whole point is that N calls to
	// Dictionary cost one decode. It is set only by the two constructors that
	// own the raw object for the message's whole life (callSyncDictionary and
	// the SetIncomingMessageHandler block); a nil memo means "decode every
	// call", which is what ReceivedMessageFromHandle keeps, because the
	// caller of that constructor owns the handle and may mutate the
	// dictionary behind it.
	//
	// This carries exactly one invariant: at most one raw decode, hence one
	// xpc_dictionary_apply block registration, per raw-backed message. It
	// carries no reply state and no one-reply flag: one-reply-per-message
	// remains a property of the call graph (see sendReplyDictionary), not of
	// anything stored here.
	memo *messageMemo
}

// messageMemo holds the decoded body of one raw-backed message. The mutex
// guards only that decode; it is not a message state machine.
type messageMemo struct {
	mu   sync.Mutex
	dict Dictionary
}

// Creation flags from <xpc/listener.h> and <xpc/session.h>. Both C enums are
// declared XPC_FLAGS_ENUM(..., uint64_t), so they are passed by value.
const (
	xpcListenerCreateNone            uint64 = 0
	xpcListenerCreateInactive        uint64 = 1 << 0
	xpcListenerCreateForceMach       uint64 = 1 << 1
	xpcListenerCreateForceXPCService uint64 = 1 << 2

	xpcSessionCreateNone           uint64 = 0
	xpcSessionCreateInactive       uint64 = 1 << 0
	xpcSessionCreateMachPrivileged uint64 = 1 << 1
)

type ListenerOptions struct {
	Inactive bool
	// ForceMach and ForceXPCService tell the runtime which kind of listener
	// this is when the name alone does not determine it. A Mach service
	// listener needs ForceMach and a MachServices entry in the job's
	// launchd.plist.
	ForceMach       bool
	ForceXPCService bool
	Requirement     *PeerRequirement
	// TargetQueue is the dispatch queue session events are submitted onto.
	// The zero Queue means no target queue, which libdispatch treats as its
	// default target queue. Dispatch queues are process-lifetime objects
	// (the dispatch package never releases one), which is what makes storing
	// one here sound.
	TargetQueue dispatch.Queue
}

type SessionOptions struct {
	Inactive    bool
	Privileged  bool
	Requirement *PeerRequirement
	// TargetQueue is the dispatch queue session events are submitted onto.
	// The zero Queue means no target queue, which libdispatch treats as its
	// default target queue. Dispatch queues are process-lifetime objects
	// (the dispatch package never releases one), which is what makes storing
	// one here sound.
	TargetQueue dispatch.Queue
}

// flags builds the xpc_listener_create_flags_t value for these options.
func (o ListenerOptions) flags() uint64 {
	f := xpcListenerCreateNone
	if o.Inactive {
		f |= xpcListenerCreateInactive
	}
	if o.ForceMach {
		f |= xpcListenerCreateForceMach
	}
	if o.ForceXPCService {
		f |= xpcListenerCreateForceXPCService
	}
	return f
}

// flags builds the xpc_session_create_flags_t value for these options.
// Privileged applies only to Mach services.
func (o SessionOptions) flags() uint64 {
	f := xpcSessionCreateNone
	if o.Inactive {
		f |= xpcSessionCreateInactive
	}
	if o.Privileged {
		f |= xpcSessionCreateMachPrivileged
	}
	return f
}

type IncomingSessionRequest struct {
	peer     unsafe.Pointer
	listener *Listener
}

type IncomingDecision struct {
	accept   bool
	reason   string
	handler  MessageHandler
	cancel   CancellationHandler
	session  *Session
	listener *Listener
}

func pointerFromHandle(handle uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&handle))
}

// targetQueuePointer converts a dispatch.Queue to the raw target-queue handle
// the create calls take. A zero queue is NULL, which libdispatch treats as
// its default target queue. The conversion goes through pointerFromHandle so
// the bit pattern is preserved; a direct unsafe.Pointer(uintptr) conversion
// is not an accepted construction path.
func targetQueuePointer(queue dispatch.Queue) unsafe.Pointer {
	return pointerFromHandle(queue.Handle())
}

// Handle returns the underlying XPC endpoint handle.
func (e Endpoint) Handle() uintptr {
	return uintptr(e.raw)
}

// EndpointFromHandle wraps a borrowed XPC endpoint handle. The caller retains
// ownership; the wrapper never releases it.
func EndpointFromHandle(handle uintptr) Endpoint {
	return Endpoint{raw: pointerFromHandle(handle)}
}

// Handle returns the underlying XPC peer requirement handle. The handle is
// borrowed: it is valid only until Close and reads as zero afterward.
func (r *PeerRequirement) Handle() uintptr {
	if r == nil || r.raw == nil {
		return 0
	}
	return uintptr(r.raw)
}

// Close releases the requirement's reference. It is safe to call twice: the
// first call releases, and a second call (or a call on an already-closed
// requirement) returns nil. The requirement must not be used after Close.
//
// If xpc_release is unavailable Close cannot release the reference. It clears
// the handle anyway, so Close stays idempotent, and returns the availability
// error; the underlying reference is then leaked and there is no way to
// recover it. This is the only outcome that leaks, and it is reported rather
// than swallowed.
func (r *PeerRequirement) Close() error {
	if r == nil || r.raw == nil {
		return nil
	}
	if err := requireRawSymbols(rawSyms_PeerRequirement_Close...); err != nil {
		r.raw = nil
		return err
	}
	raw_xpc_release(r.raw)
	r.raw = nil
	return nil
}

// PeerRequirementFromHandle is the deliberate exception to the borrowed-handle
// rule: it retains a nonzero handle and returns an owned reference that the
// caller must Close, exactly as if it had been built by one of the
// constructors. A zero handle returns nil. Supplying a handle the caller does
// not own, or one already released, is undefined behaviour at the C level.
func PeerRequirementFromHandle(handle uintptr) *PeerRequirement {
	if handle == 0 {
		return nil
	}
	if err := requireRawSymbols(rawSyms_PeerRequirementFromHandle...); err != nil {
		return nil
	}
	raw := pointerFromHandle(handle)
	raw_xpc_retain(raw)
	return &PeerRequirement{raw: raw}
}

// NewSameTeamRequirement returns a requirement that the peer be signed with
// the same team identifier as the current process. The signing identifier is
// not constrained further.
func NewSameTeamRequirement() (*PeerRequirement, error) {
	return newRequirement(func(richErr *unsafe.Pointer) unsafe.Pointer {
		return raw_xpc_peer_requirement_create_team_identity(nil, richErr)
	}, "xpc_peer_requirement_create_team_identity")
}

// NewSameTeamSignedAsRequirement returns a requirement that the peer be signed
// with the same team identifier as the current process and with the given
// signing identifier. An empty signingIdentifier is rejected: the header's
// meaningful NULL case is expressed by NewSameTeamRequirement, and an empty C
// string is not NULL.
func NewSameTeamSignedAsRequirement(signingIdentifier string) (*PeerRequirement, error) {
	buf, err := nulTerminatedSigningIdentifier(signingIdentifier)
	if err != nil {
		return nil, err
	}
	return newRequirement(func(richErr *unsafe.Pointer) unsafe.Pointer {
		return raw_xpc_peer_requirement_create_team_identity(&buf[0], richErr)
	}, "xpc_peer_requirement_create_team_identity")
}

// NewPlatformBinaryRequirement returns a requirement that the peer be a
// platform binary. The signing identifier is not constrained further.
func NewPlatformBinaryRequirement() (*PeerRequirement, error) {
	return newRequirement(func(richErr *unsafe.Pointer) unsafe.Pointer {
		return raw_xpc_peer_requirement_create_platform_identity(nil, richErr)
	}, "xpc_peer_requirement_create_platform_identity")
}

// NewPlatformBinarySignedAsRequirement returns a requirement that the peer be
// a platform binary signed with the given signing identifier. An empty
// signingIdentifier is rejected: the header's meaningful NULL case is
// expressed by NewPlatformBinaryRequirement, and an empty C string is not
// NULL.
func NewPlatformBinarySignedAsRequirement(signingIdentifier string) (*PeerRequirement, error) {
	buf, err := nulTerminatedSigningIdentifier(signingIdentifier)
	if err != nil {
		return nil, err
	}
	return newRequirement(func(richErr *unsafe.Pointer) unsafe.Pointer {
		return raw_xpc_peer_requirement_create_platform_identity(&buf[0], richErr)
	}, "xpc_peer_requirement_create_platform_identity")
}

// NewEntitlementExistsRequirement returns a requirement that the peer has the
// given entitlement. An empty entitlement is rejected.
func NewEntitlementExistsRequirement(entitlement string) (*PeerRequirement, error) {
	if entitlement == "" {
		return nil, errors.New("xpc: entitlement must not be empty")
	}
	return newRequirement(func(richErr *unsafe.Pointer) unsafe.Pointer {
		return raw_xpc_peer_requirement_create_entitlement_exists(entitlement, richErr)
	}, "xpc_peer_requirement_create_entitlement_exists")
}

// NewEntitlementMatchesRequirement returns a requirement that the peer has the
// given entitlement with a matching value. The value must be a bool, a
// string, or a signed integer that int64 represents exactly; Apple accepts
// exactly XPC_TYPE_BOOL, XPC_TYPE_STRING, and XPC_TYPE_INT64 for this call,
// so the general codec's wider surface is deliberately not used. Anything
// else is rejected with an error naming the value's type.
func NewEntitlementMatchesRequirement(entitlement string, value any) (*PeerRequirement, error) {
	if entitlement == "" {
		return nil, errors.New("xpc: entitlement must not be empty")
	}
	obj, err := entitlementValueToRawObject(value)
	if err != nil {
		return nil, err
	}
	defer releaseRaw(obj)
	return newRequirement(func(richErr *unsafe.Pointer) unsafe.Pointer {
		return raw_xpc_peer_requirement_create_entitlement_matches_value(entitlement, obj, richErr)
	}, "xpc_peer_requirement_create_entitlement_matches_value")
}

// NewLightweightCodeRequirement returns a requirement that the peer satisfy
// the given lightweight code requirement dictionary. A nil lwcr is rejected.
// The temporary object built from the dictionary is released once the
// constructor returns.
func NewLightweightCodeRequirement(lwcr Dictionary) (*PeerRequirement, error) {
	if lwcr == nil {
		return nil, errors.New("xpc: lwcr must not be nil")
	}
	obj, err := dictionaryToRawObject(lwcr)
	if err != nil {
		return nil, err
	}
	defer releaseRaw(obj)
	return newRequirement(func(richErr *unsafe.Pointer) unsafe.Pointer {
		return raw_xpc_peer_requirement_create_lwcr(obj, richErr)
	}, "xpc_peer_requirement_create_lwcr")
}

// newRequirement calls create and wraps its result in the owned
// PeerRequirement form, honouring the raw layer's error_out convention: a nil
// result with a non-nil rich error reports that error, and a nil result
// without one reports a generic failure. Symbol lookup happens after argument
// validation in every caller, so an invalid argument on a runtime that lacks
// the symbol reports the argument error.
func newRequirement(create func(richErr *unsafe.Pointer) unsafe.Pointer, symbols ...string) (*PeerRequirement, error) {
	if err := requireRawSymbols(symbols...); err != nil {
		return nil, err
	}
	var richErr unsafe.Pointer
	raw := create(&richErr)
	if raw == nil {
		if richErr != nil {
			return nil, richErrorFromRaw(richErr)
		}
		return nil, errors.New("xpc: failed to create peer requirement")
	}
	return &PeerRequirement{raw: raw}, nil
}

// nulTerminatedSigningIdentifier validates a signing identifier and returns
// it NUL-terminated. The caller keeps the slice alive for the duration of the
// native call; the constructors pass its first byte, which represents C NULL
// when nil and a real string otherwise.
func nulTerminatedSigningIdentifier(signingIdentifier string) ([]byte, error) {
	if signingIdentifier == "" {
		return nil, errors.New("xpc: signingIdentifier must not be empty")
	}
	return append([]byte(signingIdentifier), 0), nil
}

// entitlementValueToRawObject converts a value for
// xpc_peer_requirement_create_entitlement_matches_value. Apple accepts
// exactly XPC_TYPE_BOOL, XPC_TYPE_STRING, and XPC_TYPE_INT64 for this call.
// Reflection is used so named bool, string, and signed-integer types are
// treated like their underlying kinds; everything else is rejected with an
// error naming the value's Go type.
func entitlementValueToRawObject(value any) (unsafe.Pointer, error) {
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Bool:
		return raw_xpc_bool_create(rv.Bool()), nil
	case reflect.String:
		return raw_xpc_string_create(rv.String()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return raw_xpc_int64_create(rv.Int()), nil
	}
	return nil, fmt.Errorf("xpc: entitlement value of type %T is not supported; want bool, string, or a signed integer", value)
}

// Handle returns the underlying XPC rich error handle.
func (e RichError) Handle() uintptr {
	return uintptr(e.raw)
}

// RichErrorFromHandle wraps a borrowed XPC rich error handle. The caller
// retains ownership; the wrapper never releases it.
func RichErrorFromHandle(handle uintptr) RichError {
	return RichError{raw: pointerFromHandle(handle)}
}

// Handle returns the underlying XPC listener handle.
func (l *Listener) Handle() uintptr {
	if l == nil {
		return 0
	}
	return uintptr(l.raw)
}

// ListenerFromHandle wraps a borrowed XPC listener handle. The caller retains
// ownership; the wrapper never releases it.
func ListenerFromHandle(handle uintptr) *Listener {
	return &Listener{raw: pointerFromHandle(handle)}
}

// Handle returns the underlying XPC session handle.
func (s *Session) Handle() uintptr {
	if s == nil {
		return 0
	}
	return uintptr(s.raw)
}

// SessionFromHandle wraps a borrowed XPC session handle. The caller retains
// ownership; the wrapper never releases it. A session obtained this way has
// an unknowable lifecycle, so the inactive-only setters refuse it.
func SessionFromHandle(handle uintptr) *Session {
	return &Session{raw: pointerFromHandle(handle), fromHandle: true}
}

// Handle returns the underlying XPC message handle.
func (m ReceivedMessage) Handle() uintptr {
	return uintptr(m.raw)
}

// ReceivedMessageFromHandle wraps a borrowed XPC message handle. The caller
// retains ownership; the wrapper never releases it.
func ReceivedMessageFromHandle(handle uintptr) ReceivedMessage {
	return ReceivedMessage{raw: pointerFromHandle(handle)}
}

// Handle returns the underlying XPC incoming-session peer handle.
func (r IncomingSessionRequest) Handle() uintptr {
	return uintptr(r.peer)
}

// IncomingSessionRequestFromHandle wraps a borrowed XPC incoming-session peer
// handle. The caller retains ownership; the wrapper never releases it.
func IncomingSessionRequestFromHandle(handle uintptr) IncomingSessionRequest {
	return IncomingSessionRequest{peer: pointerFromHandle(handle)}
}

// MessageHandler answers one incoming message. The returned value is encoded
// with the package wire format and sent as the reply. A non-nil error replies
// with Dictionary{"error": err.Error()}. Returning (nil, nil) declines to
// reply at all; to reply with an empty body, return Dictionary{}.
//
// Exactly one reply is sent per invocation, by the runtime, after the handler
// returns. There is no way to reply twice, and no way to reply late.
type MessageHandler func(ReceivedMessage) (any, error)
type CancellationHandler func(RichError)

// Marshaler and Unmarshaler are the only points at which the wire format is
// caller-controlled. The format itself is fixed; see encodeMessage.

type Marshaler interface {
	MarshalXPC() (Dictionary, error)
}

type Unmarshaler interface {
	UnmarshalXPC(Dictionary) error
}

var (
	errXPCUnavailable = errors.New("xpc: framework unavailable")
	errNoReply        = errors.New("xpc: no reply")
)

// decodeJSONPayload decodes JSON produced by a json.Marshaler into the Go
// values the wire encoder understands.
//
// json.Unmarshal into an any turns every number into a float64, which silently
// loses integers past 2^53: 9007199254740993 arrived as 9007199254740992. The
// ordinary encoding path carries an int64 exactly, including math.MinInt64, so
// the package was inconsistent about a property it had already decided it cares
// about (see TestCodecDecodeWidensIntegers). json.Decoder.UseNumber keeps the
// literal, and the conversion below picks the narrowest wire type that holds it
// exactly: int64, then uint64 for the range above it, and float64 only for
// numbers that are not integers at all.
func decodeJSONPayload(b []byte) (any, error) {
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	var payload any
	if err := dec.Decode(&payload); err != nil {
		return nil, fmt.Errorf("xpc: unmarshal marshaled json: %w", err)
	}
	return jsonNumbersToWire(payload)
}

// jsonNumbersToWire replaces every json.Number in v with an int64, uint64 or
// float64. Containers are rebuilt rather than mutated in place, since the
// caller's marshaler may hand back a value it still holds.
func jsonNumbersToWire(v any) (any, error) {
	switch x := v.(type) {
	case json.Number:
		if i, err := x.Int64(); err == nil {
			return i, nil
		}
		if u, err := strconv.ParseUint(x.String(), 10, 64); err == nil {
			return u, nil
		}
		f, err := x.Float64()
		if err != nil {
			return nil, fmt.Errorf("xpc: json number %s is not representable on the wire: %w", x, err)
		}
		return f, nil
	case map[string]any:
		out := make(map[string]any, len(x))
		for k, elem := range x {
			conv, err := jsonNumbersToWire(elem)
			if err != nil {
				return nil, err
			}
			out[k] = conv
		}
		return out, nil
	case []any:
		out := make([]any, len(x))
		for i, elem := range x {
			conv, err := jsonNumbersToWire(elem)
			if err != nil {
				return nil, err
			}
			out[i] = conv
		}
		return out, nil
	default:
		return v, nil
	}
}

// encodeMessage converts v to the XPC dictionary sent on the wire.
//
// The encoding is fixed: it is the package's wire format, not a policy the
// caller can replace. Any change to the cases below is a wire change and must
// be accompanied by a change to the golden files in testdata/codec; see
// TestCodecEncodeGoldens.
func encodeMessage(v any) (Dictionary, error) {
	if v == nil {
		return Dictionary{}, nil
	}
	if m, ok := v.(Marshaler); ok {
		return m.MarshalXPC()
	}
	if m, ok := v.(encoding.TextMarshaler); ok {
		b, err := m.MarshalText()
		if err != nil {
			return nil, fmt.Errorf("xpc: marshal text: %w", err)
		}
		return Dictionary{"value": string(b)}, nil
	}
	if m, ok := v.(encoding.BinaryMarshaler); ok {
		b, err := m.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("xpc: marshal binary: %w", err)
		}
		return Dictionary{"value": append([]byte(nil), b...)}, nil
	}
	if m, ok := v.(json.Marshaler); ok {
		b, err := m.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("xpc: marshal json: %w", err)
		}
		payload, err := decodeJSONPayload(b)
		if err != nil {
			return nil, err
		}
		return encodeMessage(payload)
	}
	switch x := v.(type) {
	case Dictionary:
		return cloneDictionary(x), nil
	case map[string]any:
		return cloneDictionary(Dictionary(x)), nil
	}

	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return Dictionary{}, nil
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Struct:
		return structToDictionary(rv), nil
	case reflect.Map:
		if rv.Type().Key().Kind() != reflect.String {
			return nil, fmt.Errorf("xpc: map key type must be string, got %s", rv.Type().Key())
		}
		d := Dictionary{}
		iter := rv.MapRange()
		for iter.Next() {
			d[iter.Key().String()] = iter.Value().Interface()
		}
		return d, nil
	case reflect.Slice, reflect.Array:
		n := rv.Len()
		items := make([]any, 0, n)
		for i := 0; i < n; i++ {
			items = append(items, rv.Index(i).Interface())
		}
		return Dictionary{"items": items}, nil
	default:
		return Dictionary{"value": v}, nil
	}
}

// decodeMessage fills dst from the message's dictionary, inverting
// encodeMessage. See encodeMessage for the stability rule.
func decodeMessage(msg ReceivedMessage, dst any) error {
	if dst == nil {
		return errors.New("xpc: decode destination is nil")
	}
	dict := msg.Dictionary()
	if u, ok := dst.(Unmarshaler); ok {
		return u.UnmarshalXPC(dict)
	}
	switch p := dst.(type) {
	case *Dictionary:
		*p = cloneDictionary(dict)
		return nil
	case *map[string]any:
		out := make(map[string]any, len(dict))
		for k, v := range dict {
			out[k] = v
		}
		*p = out
		return nil
	}
	payload := any(dict)
	if len(dict) == 1 {
		if value, ok := dict["value"]; ok {
			payload = value
		}
	}
	if u, ok := dst.(encoding.TextUnmarshaler); ok {
		switch v := payload.(type) {
		case string:
			return u.UnmarshalText([]byte(v))
		case []byte:
			return u.UnmarshalText(v)
		default:
			return fmt.Errorf("xpc: text decode payload type %T", payload)
		}
	}
	if u, ok := dst.(encoding.BinaryUnmarshaler); ok {
		switch v := payload.(type) {
		case []byte:
			return u.UnmarshalBinary(v)
		case string:
			return u.UnmarshalBinary([]byte(v))
		default:
			return fmt.Errorf("xpc: binary decode payload type %T", payload)
		}
	}
	if u, ok := dst.(json.Unmarshaler); ok {
		b, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("xpc: marshal json payload: %w", err)
		}
		if err := u.UnmarshalJSON(b); err != nil {
			return fmt.Errorf("xpc: json unmarshal payload: %w", err)
		}
		return nil
	}
	// A struct destination is filled field by field so that xpc tags are
	// honoured, matching how structToDictionary produced the message.
	if rv := reflect.ValueOf(dst); rv.Kind() == reflect.Pointer && !rv.IsNil() {
		if elem := rv.Elem(); elem.Kind() == reflect.Struct {
			if d, ok := payload.(Dictionary); ok {
				if err := dictionaryToStruct(d, elem); err != nil {
					return fmt.Errorf("xpc: decode payload: %w", err)
				}
				return nil
			}
		}
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("xpc: marshal decode payload: %w", err)
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return fmt.Errorf("xpc: decode payload: %w", err)
	}
	return nil
}

func NewServiceListener(service string, opts ListenerOptions, incoming func(IncomingSessionRequest) IncomingDecision) (*Listener, error) {
	if service == "" {
		return nil, errors.New("xpc: empty service name")
	}
	return newListener(service, opts, incoming, true)
}

func NewAnonymousListener(opts ListenerOptions, incoming func(IncomingSessionRequest) IncomingDecision) *Listener {
	l, err := newListener("", opts, incoming, false)
	if err != nil {
		return &Listener{
			incoming:  incoming,
			createErr: err,
		}
	}
	return l
}

func newListener(service string, opts ListenerOptions, incoming func(IncomingSessionRequest) IncomingDecision, requiresService bool) (*Listener, error) {
	if frameworkHandle == 0 {
		return nil, errXPCUnavailable
	}
	if incoming == nil {
		return nil, errors.New("xpc: incoming session handler is nil")
	}
	if requiresService && service == "" {
		return nil, errors.New("xpc: empty service name")
	}
	if opts.Requirement != nil && opts.Requirement.raw == nil {
		return nil, errors.New("xpc: peer requirement is closed")
	}
	if err := requireRawSymbols(
		"xpc_listener_create",
		"xpc_listener_activate",
		"xpc_listener_cancel",
		"xpc_listener_reject_peer",
		"xpc_session_set_incoming_message_handler",
		"xpc_session_set_cancel_handler",
		"xpc_session_activate",
	); err != nil {
		return nil, err
	}

	l := &Listener{
		incoming: incoming,
	}
	incomingBlock, err := newXPCBlock(func(_ uintptr, peer unsafe.Pointer) {
		req := IncomingSessionRequest{peer: peer, listener: l}
		decision := incoming(req)
		decision.apply(req)
	})
	if err != nil {
		return nil, err
	}
	l.incomingBlk = incomingBlock

	// Installing a peer requirement needs an inactive listener: the setter
	// traps on an active one. When options carry a requirement and the caller
	// did not ask for an inactive object, create dormant, install, and
	// activate below so the returned listener's lifecycle is unchanged.
	flags := opts.flags()
	if opts.Requirement != nil && !opts.Inactive {
		flags |= xpcListenerCreateInactive
	}

	var richErr unsafe.Pointer
	raw := raw_xpc_listener_create(service, targetQueuePointer(opts.TargetQueue), flags, incomingBlock, &richErr)
	if raw == nil {
		if richErr != nil {
			return nil, richErrorFromRaw(richErr)
		}
		return nil, errors.New("xpc: failed to create listener")
	}
	l.raw = raw

	if opts.Requirement != nil {
		if err := requireRawSymbols("xpc_listener_set_peer_requirement"); err != nil {
			releaseRaw(raw)
			return nil, err
		}
		raw_xpc_listener_set_peer_requirement(l.raw, opts.Requirement.raw)
		if !opts.Inactive {
			if err := l.Activate(); err != nil {
				releaseRaw(raw)
				return nil, err
			}
		}
	}
	return l, nil
}

func (l *Listener) Activate() error {
	if l == nil {
		return errors.New("xpc: listener is nil")
	}
	if l.createErr != nil {
		return l.createErr
	}
	if l.raw == nil {
		return errors.New("xpc: listener is not initialized")
	}
	if err := requireRawSymbols(rawSyms_Listener_Activate...); err != nil {
		return err
	}
	var richErr unsafe.Pointer
	if !raw_xpc_listener_activate(l.raw, &richErr) {
		if richErr != nil {
			return richErrorFromRaw(richErr)
		}
		return errors.New("xpc: listener activation failed")
	}
	return nil
}

func (l *Listener) Cancel() {
	if l == nil || l.raw == nil {
		return
	}
	raw_xpc_listener_cancel(l.raw)
}

func (r IncomingSessionRequest) Accept(handler MessageHandler, onCancel CancellationHandler) IncomingDecision {
	d, _ := r.AcceptSession(handler, onCancel)
	return d
}

func (r IncomingSessionRequest) AcceptSession(handler MessageHandler, onCancel CancellationHandler) (IncomingDecision, *Session) {
	s := &Session{
		raw: r.peer,
	}
	return IncomingDecision{
		accept:   true,
		handler:  handler,
		cancel:   onCancel,
		session:  s,
		listener: r.listener,
	}, s
}

func (r IncomingSessionRequest) Reject(reason string) IncomingDecision {
	return IncomingDecision{
		accept: false,
		reason: reason,
	}
}

func (d IncomingDecision) apply(req IncomingSessionRequest) {
	if req.peer == nil {
		return
	}
	if !d.accept {
		reason := d.reason
		if reason == "" {
			reason = "rejected"
		}
		raw_xpc_listener_reject_peer(req.peer, reason)
		return
	}
	s := d.session
	if s == nil {
		s = &Session{raw: req.peer}
	}
	if d.handler != nil {
		_ = s.SetIncomingMessageHandler(d.handler)
	}
	if d.cancel != nil {
		_ = s.SetCancellationHandler(d.cancel)
	}
	var richErr unsafe.Pointer
	if raw_xpc_session_activate(req.peer, &richErr) {
		if s != nil {
			s.active = true
		}
	} else if d.cancel != nil {
		d.cancel(richErrorFromRaw(richErr))
	}
}

func DialXPCService(name string, opts SessionOptions) (*Session, error) {
	return dialSession("xpc", name, opts)
}

func DialMachService(name string, opts SessionOptions) (*Session, error) {
	return dialSession("mach", name, opts)
}

func dialSession(kind, name string, opts SessionOptions) (*Session, error) {
	if frameworkHandle == 0 {
		return nil, errXPCUnavailable
	}
	if name == "" {
		return nil, errors.New("xpc: empty service name")
	}
	if opts.Requirement != nil && opts.Requirement.raw == nil {
		return nil, errors.New("xpc: peer requirement is closed")
	}
	if err := requireRawSymbols("xpc_session_create_xpc_service", "xpc_session_create_mach_service"); err != nil {
		return nil, err
	}

	// Installing a peer requirement needs an inactive session: the setter
	// traps on an active one. When options carry a requirement and the caller
	// did not ask for an inactive object, create dormant, install, and
	// activate below so the returned session's lifecycle is unchanged.
	flags := opts.flags()
	if opts.Requirement != nil && !opts.Inactive {
		flags |= xpcSessionCreateInactive
	}

	var richErr unsafe.Pointer
	var raw unsafe.Pointer
	switch kind {
	case "mach":
		raw = raw_xpc_session_create_mach_service(name, targetQueuePointer(opts.TargetQueue), flags, &richErr)
	default:
		raw = raw_xpc_session_create_xpc_service(name, targetQueuePointer(opts.TargetQueue), flags, &richErr)
	}
	if raw == nil {
		if richErr != nil {
			return nil, richErrorFromRaw(richErr)
		}
		return nil, errors.New("xpc: failed to create session")
	}

	s := &Session{
		raw:    raw,
		active: !opts.Inactive && opts.Requirement == nil,
	}
	if opts.Requirement != nil {
		if err := s.SetPeerRequirement(opts.Requirement); err != nil {
			releaseRaw(raw)
			return nil, err
		}
		if !opts.Inactive {
			if err := s.Activate(); err != nil {
				releaseRaw(raw)
				return nil, err
			}
		}
	}
	return s, nil
}

func (s *Session) Activate() error {
	if s == nil {
		return errors.New("xpc: session is nil")
	}
	if s.raw == nil {
		return errors.New("xpc: session is not initialized")
	}
	if err := requireRawSymbols(rawSyms_Session_Activate...); err != nil {
		return err
	}
	var richErr unsafe.Pointer
	if !raw_xpc_session_activate(s.raw, &richErr) {
		if richErr != nil {
			return richErrorFromRaw(richErr)
		}
		return errors.New("xpc: session activation failed")
	}
	s.active = true
	return nil
}

// Cancel cancels the session. Any in-flight message fails and the cancel
// handler runs. Cancel is idempotent and safe on a nil or already-cancelled
// session.
//
// There is no reason string: xpc_session_cancel takes only the session, so a
// reason parameter would be accepted and discarded. Compare
// IncomingSessionRequest.Reject, which does carry a reason to the peer.
func (s *Session) Cancel() {
	if s == nil || s.raw == nil {
		return
	}
	if err := requireRawSymbols(rawSyms_Session_Cancel...); err != nil {
		return
	}
	raw_xpc_session_cancel(s.raw)
}

func (s *Session) SetIncomingMessageHandler(handler MessageHandler) error {
	if s == nil {
		return errors.New("xpc: session is nil")
	}
	if handler == nil {
		return errors.New("xpc: incoming handler is nil")
	}
	if err := requireRawSymbols(rawSyms_Session_SetIncomingMessageHandler...); err != nil {
		return err
	}
	s.handlerMu.Lock()
	s.incoming = handler
	s.handlerMu.Unlock()
	// One block per session, not one per installation: replacing a handler
	// reuses the block, because the block only carries the session's token
	// and the handler is read from the session at invocation time.
	if s.incomingBlk == nil {
		block, err := newXPCSessionBlock(xpcIncomingTrampoline(), s)
		if err != nil {
			return err
		}
		s.incomingBlk = block
	}
	block := s.incomingBlk
	s.incomingBlk = block
	raw_xpc_session_set_incoming_message_handler(s.raw, block)
	return nil
}

func (s *Session) SetCancellationHandler(handler CancellationHandler) error {
	if s == nil {
		return errors.New("xpc: session is nil")
	}
	if handler == nil {
		return errors.New("xpc: cancellation handler is nil")
	}
	if err := requireRawSymbols(rawSyms_Session_SetCancellationHandler...); err != nil {
		return err
	}
	s.handlerMu.Lock()
	s.cancellation = handler
	s.handlerMu.Unlock()
	if s.cancellationBl == nil {
		block, err := newXPCSessionBlock(xpcCancelTrampoline(), s)
		if err != nil {
			return err
		}
		s.cancellationBl = block
	}
	block := s.cancellationBl
	s.cancellationBl = block
	raw_xpc_session_set_cancel_handler(s.raw, block)
	return nil
}

// SetPeerRequirement installs a peer requirement on the session. The
// requirement must be installed while the session is inactive and at most
// once per session; installing on an active session, on a session created
// from a raw handle whose lifecycle is unknowable, or a second time is API
// misuse that traps the process in native code, so each is refused here
// before native code runs. The session does not retain the requirement
// beyond this call: XPC retains it on successful installation, so the caller
// may Close their reference immediately afterward.
func (s *Session) SetPeerRequirement(req *PeerRequirement) error {
	if s == nil {
		return errors.New("xpc: session is nil")
	}
	if s.raw == nil {
		return errors.New("xpc: session is not initialized")
	}
	if req == nil {
		return errors.New("xpc: peer requirement is nil")
	}
	if req.raw == nil {
		return errors.New("xpc: peer requirement is closed")
	}
	if s.fromHandle {
		return errors.New("xpc: cannot set peer requirement on a handle-derived session")
	}
	if s.active {
		return errors.New("xpc: peer requirement can only be set while the session is inactive")
	}
	if s.requirementSet {
		return errors.New("xpc: peer requirement already set")
	}
	if err := requireRawSymbols(rawSyms_Session_SetPeerRequirement...); err != nil {
		return err
	}
	raw_xpc_session_set_peer_requirement(s.raw, req.raw)
	s.requirementSet = true
	return nil
}

// SetTargetQueue replaces the session's target queue. It may be called more
// than once while the session is inactive; a zero queue restores
// libdispatch's default target queue. It refuses an active session and a
// handle-derived session before calling native code, both of which are API
// misuse.
func (s *Session) SetTargetQueue(queue dispatch.Queue) error {
	if s == nil {
		return errors.New("xpc: session is nil")
	}
	if s.raw == nil {
		return errors.New("xpc: session is not initialized")
	}
	if s.fromHandle {
		return errors.New("xpc: cannot set target queue on a handle-derived session")
	}
	if s.active {
		return errors.New("xpc: target queue can only be set while the session is inactive")
	}
	if err := requireRawSymbols(rawSyms_Session_SetTargetQueue...); err != nil {
		return err
	}
	raw_xpc_session_set_target_queue(s.raw, targetQueuePointer(queue))
	return nil
}

// NotifyDictionary sends msg without waiting for a reply. It returns when the
// message has been handed to the transport; delivery is not confirmed.
func (s *Session) NotifyDictionary(msg Dictionary) error {
	if s == nil {
		return errors.New("xpc: session is nil")
	}
	if err := requireRawSymbols(rawSyms_Session_NotifyDictionary...); err != nil {
		return err
	}
	rawMsg, err := dictionaryToRawObject(msg)
	if err != nil {
		return err
	}
	defer releaseRaw(rawMsg)
	richErr := raw_xpc_session_send_message(s.raw, rawMsg)
	if richErr != nil {
		return richErrorFromRaw(richErr)
	}
	return nil
}

// Notify encodes msg with the package wire format and sends it without
// waiting for a reply. Delivery is not confirmed.
func (s *Session) Notify(msg any) error {
	dict, err := encodeMessage(msg)
	if err != nil {
		return err
	}
	return s.NotifyDictionary(dict)
}

// callAsyncDictionary is the asynchronous C send. It is unexported: its only
// caller is CallDictionary, on the cancellable-context path. Its raw symbols
// are covered by that caller's guard, which is derived transitively.
func (s *Session) callAsyncDictionary(msg Dictionary, reply func(Dictionary, error)) error {
	if s == nil {
		return errors.New("xpc: session is nil")
	}
	if reply == nil {
		return errors.New("xpc: reply callback is nil")
	}
	rawMsg, err := dictionaryToRawObject(msg)
	if err != nil {
		return err
	}
	block, err := newXPCReplyBlock(reply)
	if err != nil {
		releaseRaw(rawMsg)
		return err
	}
	raw_xpc_session_send_message_with_reply_async(s.raw, rawMsg, unsafe.Pointer(block))
	releaseRaw(rawMsg)
	// The literal must outlive this frame; xpcReplyState holds it until the
	// trampoline runs, and this keeps it alive across the native send.
	runtime.KeepAlive(block)
	return nil
}

// CallDictionary sends msg and waits for the peer's reply.
//
// libxpc has no per-message timeout and no per-message cancellation: none of
// the xpc_session_send_message* entry points takes a deadline, and
// xpc_session_cancel tears down the whole session rather than one call. ctx is
// therefore honoured on the caller's side only. If ctx ends first,
// CallDictionary returns ctx.Err(), but the peer keeps running the request and
// the pending reply context stays alive in libxpc until the reply arrives or
// the session is cancelled: a cancelled call costs one leaked reply slot.
//
// If ctx cannot end (ctx.Done() == nil, as for context.Background),
// CallDictionary uses the cheaper blocking C send.
func (s *Session) CallDictionary(ctx context.Context, msg Dictionary) (ReceivedMessage, error) {
	if s == nil {
		return ReceivedMessage{}, errors.New("xpc: session is nil")
	}
	if err := ctx.Err(); err != nil {
		return ReceivedMessage{}, err
	}
	if ctx.Done() == nil {
		return s.callSyncDictionary(msg)
	}
	if err := requireRawSymbols(rawSyms_Session_CallDictionary...); err != nil {
		return ReceivedMessage{}, err
	}
	type outcome struct {
		dict Dictionary
		err  error
	}
	// Buffered so the native callback thread never blocks writing into a
	// call that has already given up on ctx.
	done := make(chan outcome, 1)
	if err := s.callAsyncDictionary(msg, func(d Dictionary, err error) {
		done <- outcome{d, err}
	}); err != nil {
		return ReceivedMessage{}, err
	}
	select {
	case o := <-done:
		if o.err != nil {
			return ReceivedMessage{}, o.err
		}
		return ReceivedMessage{session: s, decoded: cloneDictionary(o.dict)}, nil
	case <-ctx.Done():
		return ReceivedMessage{}, ctx.Err()
	}
}

// callSyncDictionary is the blocking C send, used when ctx cannot end.
func (s *Session) callSyncDictionary(msg Dictionary) (ReceivedMessage, error) {
	if s == nil {
		return ReceivedMessage{}, errors.New("xpc: session is nil")
	}
	if err := requireRawSymbols(rawSyms_Session_CallDictionary...); err != nil {
		return ReceivedMessage{}, err
	}
	rawMsg, err := dictionaryToRawObject(msg)
	if err != nil {
		return ReceivedMessage{}, err
	}
	defer releaseRaw(rawMsg)

	var richErr unsafe.Pointer
	reply := raw_xpc_session_send_message_with_reply_sync(s.raw, rawMsg, &richErr)
	if richErr != nil {
		return ReceivedMessage{}, richErrorFromRaw(richErr)
	}
	if reply == nil {
		return ReceivedMessage{}, errNoReply
	}
	return ReceivedMessage{
		raw:     reply,
		session: s,
		memo:    &messageMemo{},
	}, nil
}

// Call encodes msg with the package wire format, sends it, and waits for the
// peer's reply. See CallDictionary for what ctx does and does not do.
func (s *Session) Call(ctx context.Context, msg any) (ReceivedMessage, error) {
	dict, err := encodeMessage(msg)
	if err != nil {
		return ReceivedMessage{}, err
	}
	return s.CallDictionary(ctx, dict)
}

// Dictionary returns the message body. The returned map is a shallow copy, so
// the caller may modify its top-level keys.
//
// A message delivered by this package (a reply from CallDictionary, or the
// argument to a MessageHandler) decodes its raw object at most once, however
// many times Dictionary is called: each decode costs a permanently registered
// callback block, so repeated calls would otherwise consume a bounded process
// resource. A message wrapped with ReceivedMessageFromHandle has no such cache
// and decodes on every call, because its handle belongs to the caller.
func (m ReceivedMessage) Dictionary() Dictionary {
	if m.decoded != nil {
		return cloneDictionary(m.decoded)
	}
	if m.raw == nil {
		return Dictionary{}
	}
	if m.memo == nil {
		d, err := rawObjectToDictionary(m.raw)
		if err != nil || d == nil {
			return Dictionary{}
		}
		return d
	}
	m.memo.mu.Lock()
	defer m.memo.mu.Unlock()
	if m.memo.dict == nil {
		d, err := rawObjectToDictionary(m.raw)
		if err != nil || d == nil {
			d = Dictionary{}
		}
		m.memo.dict = d
	}
	return cloneDictionary(m.memo.dict)
}

func (m ReceivedMessage) Decode(dst any) error {
	return decodeMessage(m, dst)
}

// sendReplyDictionary sends the one reply for a received message. It is
// unexported and has exactly one caller, the block installed by
// SetIncomingMessageHandler, which is what makes one-reply-per-message a
// property of the control-flow graph rather than of a runtime flag. Its raw
// symbols are covered by that caller's guard, derived transitively.
func (m ReceivedMessage) sendReplyDictionary(reply Dictionary) error {
	if m.session == nil || m.session.raw == nil {
		return errors.New("xpc: cannot reply without an active session")
	}
	if m.raw == nil {
		return errors.New("xpc: reply target is missing")
	}

	replyObj := raw_xpc_dictionary_create_reply(m.raw)
	if replyObj == nil {
		return errors.New("xpc: failed to create reply dictionary")
	}
	defer releaseRaw(replyObj)

	if err := writeDictionaryToRaw(replyObj, reply); err != nil {
		return err
	}
	richErr := raw_xpc_session_send_message(m.session.raw, replyObj)
	if richErr != nil {
		return richErrorFromRaw(richErr)
	}
	return nil
}

func (m ReceivedMessage) SenderSatisfies(req *PeerRequirement) bool {
	if req == nil || req.raw == nil || m.raw == nil {
		return false
	}
	if err := requireRawSymbols(rawSyms_ReceivedMessage_SenderSatisfies...); err != nil {
		return false
	}
	var richErr unsafe.Pointer
	return raw_xpc_peer_requirement_match_received_message(req.raw, m.raw, &richErr)
}

func requireRawSymbols(symbols ...string) error {
	if frameworkHandle == 0 {
		return errXPCUnavailable
	}
	for _, symbol := range symbols {
		if err := rawSymbolError(symbol); err != nil {
			return fmt.Errorf("xpc: symbol %s unavailable: %w", symbol, err)
		}
	}
	return nil
}

func richErrorFromRaw(raw unsafe.Pointer) RichError {
	if raw == nil {
		return RichError{}
	}
	msg := goString(raw_xpc_rich_error_copy_description(raw))
	canRetry := raw_xpc_rich_error_can_retry(raw)
	return RichError{
		raw:      raw,
		message:  msg,
		canRetry: canRetry,
	}
}

func cloneDictionary(in Dictionary) Dictionary {
	if in == nil {
		return Dictionary{}
	}
	out := make(Dictionary, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

// fieldKey is the dictionary key for a struct field: an xpc tag wins, then a
// json tag, then the field name. Both directions of the codec use this, so a
// value encoded from a struct decodes back into the same fields.
func fieldKey(field reflect.StructField) string {
	if tag := field.Tag.Get("xpc"); tag != "" && tag != "-" {
		return tag
	}
	if tag := field.Tag.Get("json"); tag != "" && tag != "-" {
		if comma := indexComma(tag); comma > 0 {
			return tag[:comma]
		} else if comma < 0 {
			return tag
		}
	}
	return field.Name
}

func structToDictionary(v reflect.Value) Dictionary {
	out := Dictionary{}
	typ := v.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			continue
		}
		out[fieldKey(field)] = v.Field(i).Interface()
	}
	return out
}

// dictionaryToStruct fills a struct from a dictionary using fieldKey. Decoding
// used to fall through to encoding/json, which does not know about xpc tags,
// so every xpc-tagged field silently kept its zero value.
func dictionaryToStruct(dict Dictionary, v reflect.Value) error {
	typ := v.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			continue
		}
		raw, ok := dict[fieldKey(field)]
		if !ok || raw == nil {
			continue
		}
		if err := assignValue(v.Field(i), raw); err != nil {
			return fmt.Errorf("field %s: %w", field.Name, err)
		}
	}
	return nil
}

// assignValue stores one decoded XPC value into a struct field. XPC reports
// every integer as int64 or uint64 and every float as float64, so widths are
// converted here rather than rejected.
func assignValue(dst reflect.Value, raw any) error {
	rv := reflect.ValueOf(raw)
	if rv.Type().AssignableTo(dst.Type()) {
		dst.Set(rv)
		return nil
	}
	if nested, ok := raw.(Dictionary); ok && dst.Kind() == reflect.Struct {
		return dictionaryToStruct(nested, dst)
	}
	if rv.Type().ConvertibleTo(dst.Type()) {
		switch dst.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64, reflect.String, reflect.Bool:
			dst.Set(rv.Convert(dst.Type()))
			return nil
		}
	}
	// Anything else (slices, maps, nested pointers) goes through json, which
	// handles the shapes reflect conversion does not.
	b, err := json.Marshal(raw)
	if err != nil {
		return fmt.Errorf("marshal %T: %w", raw, err)
	}
	if err := json.Unmarshal(b, dst.Addr().Interface()); err != nil {
		return fmt.Errorf("unmarshal into %s: %w", dst.Type(), err)
	}
	return nil
}

func indexComma(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			return i
		}
	}
	return -1
}

func dictionaryToRawObject(dict Dictionary) (unsafe.Pointer, error) {
	if err := requireRawSymbols("xpc_dictionary_create_empty", "xpc_dictionary_set_value"); err != nil {
		return nil, err
	}
	raw := raw_xpc_dictionary_create_empty()
	if raw == nil {
		return nil, errors.New("xpc: failed to allocate dictionary")
	}
	if err := writeDictionaryToRaw(raw, dict); err != nil {
		releaseRaw(raw)
		return nil, err
	}
	return raw, nil
}

func writeDictionaryToRaw(dst unsafe.Pointer, dict Dictionary) error {
	for key, value := range dict {
		if err := writeRawDictionaryValue(dst, key, value); err != nil {
			return err
		}
	}
	return nil
}

func writeRawDictionaryValue(dst unsafe.Pointer, key string, value any) error {
	switch v := value.(type) {
	case nil:
		nullObj := raw_xpc_null_create()
		raw_xpc_dictionary_set_value(dst, key, nullObj)
		return nil
	case bool:
		raw_xpc_dictionary_set_bool(dst, key, v)
		return nil
	case int:
		raw_xpc_dictionary_set_int64(dst, key, int64(v))
		return nil
	case int8:
		raw_xpc_dictionary_set_int64(dst, key, int64(v))
		return nil
	case int16:
		raw_xpc_dictionary_set_int64(dst, key, int64(v))
		return nil
	case int32:
		raw_xpc_dictionary_set_int64(dst, key, int64(v))
		return nil
	case int64:
		raw_xpc_dictionary_set_int64(dst, key, v)
		return nil
	case uint:
		raw_xpc_dictionary_set_uint64(dst, key, uint64(v))
		return nil
	case uint8:
		raw_xpc_dictionary_set_uint64(dst, key, uint64(v))
		return nil
	case uint16:
		raw_xpc_dictionary_set_uint64(dst, key, uint64(v))
		return nil
	case uint32:
		raw_xpc_dictionary_set_uint64(dst, key, uint64(v))
		return nil
	case uint64:
		raw_xpc_dictionary_set_uint64(dst, key, v)
		return nil
	case float32:
		raw_xpc_dictionary_set_double(dst, key, float64(v))
		return nil
	case float64:
		raw_xpc_dictionary_set_double(dst, key, v)
		return nil
	case string:
		raw_xpc_dictionary_set_string(dst, key, v)
		return nil
	case []byte:
		if len(v) == 0 {
			raw_xpc_dictionary_set_data(dst, key, nil, 0)
			return nil
		}
		raw_xpc_dictionary_set_data(dst, key, unsafe.Pointer(&v[0]), uintptr(len(v)))
		return nil
	case Dictionary:
		child, err := dictionaryToRawObject(v)
		if err != nil {
			return err
		}
		defer releaseRaw(child)
		raw_xpc_dictionary_set_value(dst, key, child)
		return nil
	case map[string]any:
		child, err := dictionaryToRawObject(Dictionary(v))
		if err != nil {
			return err
		}
		defer releaseRaw(child)
		raw_xpc_dictionary_set_value(dst, key, child)
		return nil
	case []any:
		arr, err := sliceToRawArray(v)
		if err != nil {
			return err
		}
		defer releaseRaw(arr)
		raw_xpc_dictionary_set_value(dst, key, arr)
		return nil
	case Endpoint:
		// A zero Endpoint used to write nothing at all and report success,
		// so the key was simply missing from the message the peer received
		// and the sender was told it had been sent.
		if v.raw == nil {
			return fmt.Errorf("xpc: zero Endpoint for key %q", key)
		}
		raw_xpc_dictionary_set_value(dst, key, v.raw)
		return nil
	case time.Time:
		date := raw_xpc_date_create(v.UnixNano())
		if date == nil {
			return fmt.Errorf("xpc: cannot create date for key %q", key)
		}
		defer releaseRaw(date)
		raw_xpc_dictionary_set_value(dst, key, date)
		return nil
	case UUID:
		uuid := raw_xpc_uuid_create(v)
		if uuid == nil {
			return fmt.Errorf("xpc: cannot create uuid for key %q", key)
		}
		defer releaseRaw(uuid)
		raw_xpc_dictionary_set_value(dst, key, uuid)
		return nil
	case Unsupported:
		// Unsupported records what arrived; it cannot reconstitute it. Say
		// so rather than sending the description as if it were the value.
		return fmt.Errorf("xpc: cannot encode Unsupported value of type %s for key %q", v.Type, key)
	default:
		return fmt.Errorf("xpc: unsupported dictionary value type %T for key %q", value, key)
	}
}

func sliceToRawArray(values []any) (unsafe.Pointer, error) {
	if err := requireRawSymbols("xpc_array_create_empty", "xpc_array_append_value"); err != nil {
		return nil, err
	}
	raw := raw_xpc_array_create_empty()
	if raw == nil {
		return nil, errors.New("xpc: failed to allocate array")
	}
	for _, value := range values {
		obj, err := scalarToRawObject(value)
		if err != nil {
			releaseRaw(raw)
			return nil, err
		}
		// scalarToRawObject returns a reference this loop owns, and
		// xpc_array_append_value takes its own, so the loop's has to go
		// back. Without this every element of every array leaked one.
		raw_xpc_array_append_value(raw, obj)
		releaseRaw(obj)
	}
	return raw, nil
}

func scalarToRawObject(value any) (unsafe.Pointer, error) {
	switch v := value.(type) {
	case nil:
		return raw_xpc_null_create(), nil
	case bool:
		return raw_xpc_bool_create(v), nil
	case int:
		return raw_xpc_int64_create(int64(v)), nil
	case int8:
		return raw_xpc_int64_create(int64(v)), nil
	case int16:
		return raw_xpc_int64_create(int64(v)), nil
	case int32:
		return raw_xpc_int64_create(int64(v)), nil
	case int64:
		return raw_xpc_int64_create(v), nil
	case uint:
		return raw_xpc_uint64_create(uint64(v)), nil
	case uint8:
		return raw_xpc_uint64_create(uint64(v)), nil
	case uint16:
		return raw_xpc_uint64_create(uint64(v)), nil
	case uint32:
		return raw_xpc_uint64_create(uint64(v)), nil
	case uint64:
		return raw_xpc_uint64_create(v), nil
	case float32:
		return raw_xpc_double_create(float64(v)), nil
	case float64:
		return raw_xpc_double_create(v), nil
	case string:
		return raw_xpc_string_create(v), nil
	case []byte:
		if len(v) == 0 {
			return raw_xpc_data_create(nil, 0), nil
		}
		return raw_xpc_data_create(unsafe.Pointer(&v[0]), uintptr(len(v))), nil
	case Dictionary:
		return dictionaryToRawObject(v)
	case map[string]any:
		return dictionaryToRawObject(Dictionary(v))
	case time.Time:
		return raw_xpc_date_create(v.UnixNano()), nil
	case UUID:
		return raw_xpc_uuid_create(v), nil
	case Endpoint:
		// Arrays decode through rawObjectToValue, which yields Endpoint,
		// so an array that arrived carrying one has to be able to go back.
		if v.raw == nil {
			return nil, errors.New("xpc: zero Endpoint")
		}
		raw_xpc_retain(v.raw)
		return v.raw, nil
	case Unsupported:
		return nil, fmt.Errorf("xpc: cannot encode Unsupported value of type %s", v.Type)
	default:
		return nil, fmt.Errorf("xpc: unsupported value type %T", value)
	}
}

func rawObjectToDictionary(raw unsafe.Pointer) (Dictionary, error) {
	if raw == nil {
		return Dictionary{}, nil
	}
	if err := requireRawSymbols("xpc_dictionary_apply"); err != nil {
		return nil, err
	}
	out := Dictionary{}
	blk, token, err := newXPCApplierBlock(xpcDictApplier(), out)
	if err != nil {
		return nil, err
	}
	defer xpcApplierRelease(token)
	raw_xpc_dictionary_apply(raw, unsafe.Pointer(blk))
	runtime.KeepAlive(blk)
	return out, nil
}

func rawArrayToSlice(raw unsafe.Pointer) []any {
	if raw == nil {
		return nil
	}
	if err := requireRawSymbols("xpc_array_apply"); err != nil {
		return nil
	}
	var out []any
	blk, token, err := newXPCApplierBlock(xpcArrayApplier(), &out)
	if err != nil {
		return nil
	}
	defer xpcApplierRelease(token)
	raw_xpc_array_apply(raw, unsafe.Pointer(blk))
	runtime.KeepAlive(blk)
	return out
}

func rawObjectToValue(raw unsafe.Pointer) any {
	if raw == nil {
		return nil
	}
	typ := raw_xpc_get_type(raw)
	switch {
	case typ == xpcTypeSymbol("xpc_type_null"):
		return nil
	case typ == xpcTypeSymbol("xpc_type_bool"):
		return raw_xpc_bool_get_value(raw)
	case typ == xpcTypeSymbol("xpc_type_int64"):
		return raw_xpc_int64_get_value(raw)
	case typ == xpcTypeSymbol("xpc_type_uint64"):
		return raw_xpc_uint64_get_value(raw)
	case typ == xpcTypeSymbol("xpc_type_double"):
		return raw_xpc_double_get_value(raw)
	case typ == xpcTypeSymbol("xpc_type_string"):
		return goString(raw_xpc_string_get_string_ptr(raw))
	case typ == xpcTypeSymbol("xpc_type_data"):
		return copyRawData(raw)
	case typ == xpcTypeSymbol("xpc_type_dictionary"):
		d, _ := rawObjectToDictionary(raw)
		return d
	case typ == xpcTypeSymbol("xpc_type_array"):
		return rawArrayToSlice(raw)
	case typ == xpcTypeSymbol("xpc_type_endpoint"):
		return Endpoint{raw: raw}
	case typ == xpcTypeSymbol("xpc_type_date"):
		// XPC dates are nanoseconds since the Unix epoch.
		return time.Unix(0, raw_xpc_date_get_value(raw))
	case typ == xpcTypeSymbol("xpc_type_uuid"):
		var u UUID
		if p := raw_xpc_uuid_get_bytes(raw); p != nil {
			copy(u[:], unsafe.Slice(p, len(u)))
		}
		return u
	}
	// Not a type this package models. Decoding must not invent a value that
	// a caller cannot distinguish from real data: returning the description
	// string here made a file descriptor arrive as an ordinary Go string,
	// with no error and nothing to test. Unsupported is that distinction.
	return Unsupported{
		Type:        goString(raw_xpc_type_get_name(typ)),
		Description: goString(raw_xpc_copy_description(raw)),
	}
}

func copyRawData(raw unsafe.Pointer) []byte {
	if raw == nil {
		return nil
	}
	n := raw_xpc_data_get_length(raw)
	if n == 0 {
		return []byte{}
	}
	ptr := raw_xpc_data_get_bytes_ptr(raw)
	if ptr == nil {
		return []byte{}
	}
	src := unsafe.Slice((*byte)(ptr), int(n))
	out := make([]byte, len(src))
	copy(out, src)
	return out
}

var (
	xpcTypeMu    sync.Mutex
	xpcTypeCache = map[string]unsafe.Pointer{}
)

func xpcTypeSymbol(symbol string) unsafe.Pointer {
	xpcTypeMu.Lock()
	defer xpcTypeMu.Unlock()
	if v, ok := xpcTypeCache[symbol]; ok {
		return v
	}
	if frameworkHandle == 0 {
		xpcTypeCache[symbol] = nil
		return nil
	}
	// The headers define each type as a macro over an underscore-prefixed
	// object, for example
	//
	//	#define XPC_TYPE_INT64 (&_xpc_type_int64)
	//
	// so the exported C symbol is "_xpc_type_int64" and the type value is its
	// ADDRESS. Looking up the unprefixed name resolves to nothing, and
	// dereferencing the address yields the first word of the type struct
	// rather than the type. Either mistake makes every comparison in
	// rawObjectToValue fail, so every value decodes through the
	// copy_description fallback as a string.
	sym, err := purego.Dlsym(frameworkHandle, "_"+symbol)
	if err != nil || sym == 0 {
		xpcTypeCache[symbol] = nil
		return nil
	}
	value := pointerFromHandle(sym)
	xpcTypeCache[symbol] = value
	return value
}

func releaseRaw(raw unsafe.Pointer) {
	if raw == nil {
		return
	}
	raw_xpc_release(raw)
}

type xpcBlockDescriptor struct {
	reserved uint64
	size     uint64
}

type xpcBlockLiteral struct {
	isa        uintptr
	flags      int32
	reserved   int32
	invoke     uintptr
	descriptor *xpcBlockDescriptor
}

const xpcBlockIsGlobal = 1 << 28

var (
	xpcBlockKeepaliveMu sync.Mutex
	xpcBlockKeepalive   []any
	xpcNSBlockOnce      sync.Once
	xpcNSBlockClass     uintptr
	xpcNSBlockErr       error
)

// xpcBlockClass resolves _NSConcreteGlobalBlock, the isa every global block
// literal carries.
//
// The handle matters: dlsym's process-wide search is RTLD_DEFAULT, not 0. With
// handle 0 purego reports "invalid handle" for every lookup, so this used to
// resolve to nil and the error was discarded, leaving every block literal with
// a nil isa. libxpc tolerates that today because a global block is never sent
// -copy or -release, but a nil isa is one ObjC message away from a crash and
// the discarded error hid it. The error is now recorded and returned to the
// caller instead.
func xpcBlockClass() (uintptr, error) {
	xpcNSBlockOnce.Do(func() {
		xpcNSBlockClass, xpcNSBlockErr = purego.Dlsym(purego.RTLD_DEFAULT, "_NSConcreteGlobalBlock")
		if xpcNSBlockErr == nil && xpcNSBlockClass == 0 {
			xpcNSBlockErr = errors.New("resolved to nil")
		}
	})
	if xpcNSBlockErr != nil {
		return 0, fmt.Errorf("xpc: resolve _NSConcreteGlobalBlock: %w", xpcNSBlockErr)
	}
	return xpcNSBlockClass, nil
}

// newXPCBlock wraps fn in a global block literal and registers a purego
// callback for it.
//
// Every call consumes one slot in purego's callback table, which is a fixed
// array of maxCB (2000) entries with no unregister API; exhausting it panics,
// and the panic unwinds across a native callback thread, which aborts the
// process. Use this only for blocks whose lifetime is a session's or a
// listener's, never on a per-call or per-message path. For the latter use
// newXPCApplierBlock, which shares one callback per block signature.
func newXPCBlock(fn any) (unsafe.Pointer, error) {
	class, err := xpcBlockClass()
	if err != nil {
		return nil, err
	}
	desc := &xpcBlockDescriptor{
		size: uint64(unsafe.Sizeof(xpcBlockLiteral{})),
	}
	block := &xpcBlockLiteral{
		isa:        class,
		flags:      xpcBlockIsGlobal,
		invoke:     purego.NewCallback(fn),
		descriptor: desc,
	}
	xpcBlockKeepaliveMu.Lock()
	xpcBlockKeepalive = append(xpcBlockKeepalive, fn, block, desc)
	xpcBlockKeepaliveMu.Unlock()
	return unsafe.Pointer(block), nil
}

// xpcApplierBlock is a block literal extended with a token identifying the
// per-call state its invocation needs.
//
// The applier blocks passed to xpc_dictionary_apply and xpc_array_apply are
// stateless dispatchers: the only thing that varies per call is the
// accumulator. A block's first invocation argument is the block's own address,
// so the accumulator can travel in the literal rather than in a fresh Go
// closure, and one purego callback per block signature serves every call. The
// literal itself is an ordinary Go allocation and costs no callback slot.
type xpcApplierBlock struct {
	xpcBlockLiteral
	token uint64
}

var (
	xpcApplierMu    sync.Mutex
	xpcApplierNext  uint64
	xpcApplierState = map[uint64]any{}

	xpcDictApplierOnce  sync.Once
	xpcDictApplierFn    uintptr
	xpcArrayApplierOnce sync.Once
	xpcArrayApplierFn   uintptr
)

// xpcApplierRegister records state and returns its token. The token is keyed by
// a monotonically increasing counter, never reused, and is freed by
// xpcApplierRelease. Both apply functions are synchronous, so every caller
// releases its token before returning and the table's steady-state size is the
// number of decodes in flight.
func xpcApplierRegister(state any) uint64 {
	xpcApplierMu.Lock()
	defer xpcApplierMu.Unlock()
	xpcApplierNext++
	token := xpcApplierNext
	xpcApplierState[token] = state
	return token
}

func xpcApplierLookup(token uint64) any {
	xpcApplierMu.Lock()
	defer xpcApplierMu.Unlock()
	return xpcApplierState[token]
}

func xpcApplierRelease(token uint64) {
	xpcApplierMu.Lock()
	delete(xpcApplierState, token)
	xpcApplierMu.Unlock()
}

// newXPCApplierBlock builds a block literal that dispatches to the shared
// trampoline at invoke and carries state under a fresh token. The caller must
// release the token once the apply call has returned, and must keep the
// returned literal alive across it.
func newXPCApplierBlock(invoke uintptr, state any) (*xpcApplierBlock, uint64, error) {
	class, err := xpcBlockClass()
	if err != nil {
		return nil, 0, err
	}
	token := xpcApplierRegister(state)
	return &xpcApplierBlock{
		xpcBlockLiteral: xpcBlockLiteral{
			isa:   class,
			flags: xpcBlockIsGlobal,
			// The descriptor's size must cover the whole literal,
			// token included, or a copying runtime would truncate it.
			invoke:     invoke,
			descriptor: &xpcBlockDescriptor{size: uint64(unsafe.Sizeof(xpcApplierBlock{}))},
		},
		token: token,
	}, token, nil
}

// xpcDictApplier returns the one callback shared by every
// xpc_dictionary_apply block.
func xpcDictApplier() uintptr {
	xpcDictApplierOnce.Do(func() {
		// The first parameter is typed as the literal rather than as a
		// uintptr on purpose: converting an incoming uintptr back to a
		// pointer is the unsafe.Pointer misuse go vet reports, and the
		// value is a live block address the caller holds, so purego can
		// hand it over as a pointer directly.
		xpcDictApplierFn = purego.NewCallback(func(blk *xpcApplierBlock, key *byte, value unsafe.Pointer) bool {
			out, ok := xpcApplierLookup(blk.token).(Dictionary)
			if !ok {
				// The token was released before the apply
				// returned. Stop rather than write to a dead
				// accumulator.
				return false
			}
			out[goString(key)] = rawObjectToValue(value)
			return true
		})
	})
	return xpcDictApplierFn
}

// xpcSessionBlock is a block literal carrying a session token. Session
// handlers were the other per-traffic callback registration: a listener
// installs an incoming-message handler and a cancellation handler on every
// accepted peer, so two callback slots were consumed per connection and a
// listener died of callback exhaustion after about a thousand peers regardless
// of how cheap decoding became.
type xpcSessionBlock struct {
	xpcBlockLiteral
	token uint64
}

var (
	xpcSessionMu     sync.Mutex
	xpcSessionNext   uint64
	xpcSessionTokens = map[uint64]*Session{}

	xpcIncomingOnce sync.Once
	xpcIncomingFn   uintptr
	xpcCancelOnce   sync.Once
	xpcCancelFn     uintptr
)

// newXPCSessionBlock builds a block literal that dispatches to the shared
// trampoline at invoke and identifies s by token.
//
// The token is NOT freed. Unlike an applier token, whose apply call is
// synchronous, a session block's lifetime belongs to libxpc: it retains the
// block, and the cancellation handler fires after the session is cancelled, so
// there is no point at which releasing the token is known to be safe. Dropping
// it early would silently stop delivering messages, which is worse than
// holding it. What the token table buys is the kind of growth: a map entry per
// session is ordinary memory that grows without limit, where a purego callback
// slot per session hits a fixed ceiling of 2000 and aborts the process. Each
// session takes at most one incoming token and one cancel token no matter how
// many times its handlers are replaced.
func newXPCSessionBlock(invoke uintptr, s *Session) (unsafe.Pointer, error) {
	class, err := xpcBlockClass()
	if err != nil {
		return nil, err
	}
	xpcSessionMu.Lock()
	xpcSessionNext++
	token := xpcSessionNext
	xpcSessionTokens[token] = s
	xpcSessionMu.Unlock()
	blk := &xpcSessionBlock{
		xpcBlockLiteral: xpcBlockLiteral{
			isa:        class,
			flags:      xpcBlockIsGlobal,
			invoke:     invoke,
			descriptor: &xpcBlockDescriptor{size: uint64(unsafe.Sizeof(xpcSessionBlock{}))},
		},
		token: token,
	}
	// The literal must outlive the native retain, and nothing else refers to
	// it once the caller stores only the unsafe.Pointer, so keep it here too.
	xpcBlockKeepaliveMu.Lock()
	xpcBlockKeepalive = append(xpcBlockKeepalive, blk, blk.descriptor, blk)
	xpcBlockKeepaliveMu.Unlock()
	return unsafe.Pointer(blk), nil
}

func xpcSessionForToken(token uint64) *Session {
	xpcSessionMu.Lock()
	defer xpcSessionMu.Unlock()
	return xpcSessionTokens[token]
}

// xpcIncomingTrampoline returns the one callback shared by every session's
// incoming-message block.
func xpcIncomingTrampoline() uintptr {
	xpcIncomingOnce.Do(func() {
		xpcIncomingFn = purego.NewCallback(func(blk *xpcSessionBlock, message unsafe.Pointer) {
			s := xpcSessionForToken(blk.token)
			if s == nil {
				return
			}
			s.handlerMu.Lock()
			handler := s.incoming
			s.handlerMu.Unlock()
			if handler == nil {
				return
			}
			received := ReceivedMessage{
				raw:     message,
				session: s,
				memo:    &messageMemo{},
			}
			reply, err := handler(received)
			if reply == nil && err == nil {
				// The handler declined to reply. XPC releases the
				// unused reply context, which the peer observes
				// as an interrupted connection.
				return
			}
			if err != nil {
				_ = received.sendReplyDictionary(Dictionary{"error": err.Error()})
				return
			}
			dict, encErr := encodeMessage(reply)
			if encErr != nil {
				_ = received.sendReplyDictionary(Dictionary{"error": encErr.Error()})
				return
			}
			_ = received.sendReplyDictionary(dict)
		})
	})
	return xpcIncomingFn
}

// xpcCancelTrampoline returns the one callback shared by every session's
// cancellation block.
func xpcCancelTrampoline() uintptr {
	xpcCancelOnce.Do(func() {
		xpcCancelFn = purego.NewCallback(func(blk *xpcSessionBlock, richErr unsafe.Pointer) {
			s := xpcSessionForToken(blk.token)
			if s == nil {
				return
			}
			s.handlerMu.Lock()
			handler := s.cancellation
			s.handlerMu.Unlock()
			if handler != nil {
				handler(richErrorFromRaw(richErr))
			}
		})
	})
	return xpcCancelFn
}

// xpcReplyBlock is a block literal carrying a token identifying the Go
// callback for one in-flight reply.
//
// The reply block was the last per-message callback registration. Every
// CallDictionary on the cancellable-context path built its reply handler with
// newXPCBlock, which burns one entry of purego's fixed 2000-entry callback
// table with no way to give it back, so any caller passing a context with a
// deadline or a cancel died at about two thousand calls — and died by panicking
// on a native callback thread, which aborts the process rather than returning
// an error. The literal is now an ordinary Go allocation carrying a token, and
// one shared callback serves every call.
type xpcReplyBlock struct {
	xpcBlockLiteral
	token uint64
}

// xpcReplyEntry holds the Go callback for one in-flight reply and the literal
// libxpc was handed. The literal is kept here rather than in
// xpcBlockKeepalive: the keepalive slice only grows, and a per-message entry
// on it would trade a bounded callback table for unbounded memory. This entry
// is deleted when the reply arrives, so the table's steady-state size is the
// number of calls in flight.
type xpcReplyEntry struct {
	fn  func(Dictionary, error)
	blk *xpcReplyBlock
}

var (
	xpcReplyMu    sync.Mutex
	xpcReplyNext  uint64
	xpcReplyState = map[uint64]*xpcReplyEntry{}

	xpcReplyOnce sync.Once
	xpcReplyFn   uintptr
)

// xpcReplyTake removes and returns the entry for token. A reply block is
// invoked at most once by libxpc — on the reply, or on the rich error that
// stands in for it — so taking the entry is what frees the token, and a second
// invocation finds nothing and does nothing rather than calling a Go closure
// twice.
func xpcReplyTake(token uint64) *xpcReplyEntry {
	xpcReplyMu.Lock()
	defer xpcReplyMu.Unlock()
	e := xpcReplyState[token]
	delete(xpcReplyState, token)
	return e
}

// xpcReplyTrampoline returns the one callback shared by every reply block.
func xpcReplyTrampoline() uintptr {
	xpcReplyOnce.Do(func() {
		xpcReplyFn = purego.NewCallback(func(blk *xpcReplyBlock, message unsafe.Pointer, richErr unsafe.Pointer) {
			e := xpcReplyTake(blk.token)
			if e == nil {
				return
			}
			// The callback is bound to a local before it is called, so the
			// call reads as "reply(...)" rather than "e.fn(...)". That is
			// not cosmetic: the raw-symbol reachability analysis follows a
			// call through an identifier and records it as an unresolved
			// edge to be justified in rawReachAllowed, but a call through a
			// struct field of func type is dropped without a report (see
			// xpcgen/rawreach.go: a selector whose name is not a declared
			// method "cannot reach a raw_ call"). Spelling it this way keeps
			// this edge in the audited population.
			reply := e.fn
			if richErr != nil {
				reply(nil, richErrorFromRaw(richErr))
				return
			}
			d, derr := rawObjectToDictionary(message)
			reply(d, derr)
		})
	})
	return xpcReplyFn
}

// newXPCReplyBlock builds a reply block dispatching to the shared trampoline
// and carrying fn under a fresh token. The token is freed by the trampoline.
//
// If libxpc never invokes the block — a session torn down with a send still
// pending — the entry stays in the map. That is ordinary memory, one map entry
// per abandoned call, and it is the deliberate trade: the resource it replaces
// has a hard ceiling of 2000 and aborts the process when it is reached.
func newXPCReplyBlock(fn func(Dictionary, error)) (*xpcReplyBlock, error) {
	class, err := xpcBlockClass()
	if err != nil {
		return nil, err
	}
	invoke := xpcReplyTrampoline()
	blk := &xpcReplyBlock{
		xpcBlockLiteral: xpcBlockLiteral{
			isa:   class,
			flags: xpcBlockIsGlobal,
			// The descriptor's size must cover the whole literal,
			// token included, or a copying runtime would truncate it.
			invoke:     invoke,
			descriptor: &xpcBlockDescriptor{size: uint64(unsafe.Sizeof(xpcReplyBlock{}))},
		},
	}
	xpcReplyMu.Lock()
	xpcReplyNext++
	blk.token = xpcReplyNext
	xpcReplyState[blk.token] = &xpcReplyEntry{fn: fn, blk: blk}
	xpcReplyMu.Unlock()
	return blk, nil
}

// xpcArrayApplier returns the one callback shared by every xpc_array_apply
// block.
func xpcArrayApplier() uintptr {
	xpcArrayApplierOnce.Do(func() {
		xpcArrayApplierFn = purego.NewCallback(func(blk *xpcApplierBlock, _ uintptr, value unsafe.Pointer) bool {
			out, ok := xpcApplierLookup(blk.token).(*[]any)
			if !ok {
				return false
			}
			*out = append(*out, rawObjectToValue(value))
			return true
		})
	})
	return xpcArrayApplierFn
}

func goString(p *byte) string {
	if p == nil {
		return ""
	}
	var n int
	for x := p; *x != 0; x = (*byte)(unsafe.Add(unsafe.Pointer(x), 1)) {
		n++
	}
	if n == 0 {
		return ""
	}
	return string(unsafe.Slice(p, n))
}
