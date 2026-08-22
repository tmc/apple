package xpc

import (
	"os"
	"syscall"
	"testing"

	"github.com/tmc/apple/dispatch"
)

// This file is hand-maintained. It is deliberately NOT generated: it is where
// a human writes down why a gap in the raw-symbol reachability analysis is
// acceptable. The generator emits the gaps; a person accepts them here.
//
// Three tables live here:
//
//   - rawReachAllowed: call edges the generator could not follow.
//   - canaries: entry points whose availability guard is proven to fire.
//   - canaryUnmeasured: entry points whose guard is NOT proven, with a reason.
//
// UNMEASURED is a third column, distinct from proven and from unguarded.

// rawReachAllowed is keyed "Entry|Expr". The value is the justification.
// TestRawReachUnresolvedIsAllowlisted fails when the generator reports an edge
// that is not listed here, so a newly unfollowable call cannot slip in.
var rawReachAllowed = map[string]string{
	// The reply callback's edge moved out of callAsyncDictionary and into the
	// shared reply trampoline when reply blocks stopped registering a callback
	// per message. It is the same callback with the same justification, at its
	// new home.
	"xpcReplyTrampoline|reply(...)": "the reply callback is supplied by CallDictionary, which only writes to a channel; no caller can substitute one, because callAsyncDictionary is unexported",
	// The handler bodies moved out of Set{Incoming,Cancellation}Handler and
	// into the two shared block trampolines when session blocks stopped being
	// registered per session; the edges are the same four, at their new home.
	"xpcCancelTrampoline|handler(...)":           "caller-supplied CancellationHandler; not package code",
	"xpcIncomingTrampoline|handler(...)":         "caller-supplied MessageHandler; not package code",
	"xpcIncomingTrampoline|err.Error(...)":       "error interface method; RichError.Error reaches no raw_ call",
	"xpcIncomingTrampoline|encErr.Error(...)":    "error interface method on encodeMessage's error; reaches no raw_ call",
	"(RichError).Error|e.cause.Error(...)":       "wrapped error's Error method; not package code",
	"newListener|incoming(...)":                  "caller-supplied incoming handler; not package code",
	"newRequirement|create(...)":                 "the create closure is supplied by each New*Requirement constructor and its raw call is recorded at the constructor",
	"targetQueuePointer|queue.Handle(...)":       "dispatch.Queue.Handle lives in another package and cannot reach an xpc raw_ call",
	"SetEventStreamHandler|handler(...)":         "caller-supplied EventHandler; package code only decodes the inbound dictionary before invoking it",
	"newConnection|handler(...)":                 "caller-supplied ConnectionHandler; package code only decodes the inbound dictionary before invoking it",
	"RegisterActivity|handler(...)":              "caller-supplied ActivityHandler; package code retains the inbound activity before invoking it",
	"encodeMessage|iter.Key(...).String(...)":    "reflect.Value.String is a standard-library method and cannot reach an xpc raw_ call",
	"entitlementValueToRawObject|rv.String(...)": "reflect.Value.String is a standard-library method and cannot reach an xpc raw_ call",
	"jsonNumbersToWire|x.String(...)":            "json.Number.String is a standard-library method and cannot reach an xpc raw_ call",
	"decodeJSONPayload|dec.Decode(...)":          "encoding/json's Decoder.Decode; it decodes bytes into an any and reaches no xpc raw_ call",
	// mmap and munmap are bound from libxpc's dependency chain rather than
	// generated, so they carry no raw_ prefix and the analysis cannot follow
	// the call. Neither reaches an xpc raw_ call: both are the C library
	// entry points of the same name.
	"mapShared|libcfn_mmap(...)":      "libc mmap bound through the framework handle; it is not an xpc symbol and reaches no raw_ call",
	"munmapRegion|libcfn_munmap(...)": "libc munmap bound through the framework handle; it is not an xpc symbol and reaches no raw_ call",
	"freeMemory|libcfn_free(...)":     "libc free bound through the framework handle; it is not an xpc symbol and reaches no raw_ call",
	// Mapping.Close is a method on a concrete type the analysis did not
	// resolve at this site. It reaches munmapRegion and nothing else.
	"NewSharedMemory|m.Close(...)": "Mapping.Close unmaps the region through munmapRegion; it reaches no xpc raw_ call",
}

// canaryRow is one proven-firing availability guard.
type canaryRow struct {
	entry  string
	poison string
	call   func(*testing.T) error
}

// canaries proves the guards fire. Every row's call path must be safe to run
// unpoisoned on a machine with XPC present: the negative control invokes it
// for real. That is why only the PeerRequirement constructors (and Close on an
// object one of them built) appear here — every other entry point needs a live
// session, listener or received message, and running its negative control with
// a fabricated pointer would call native XPC on a bogus address.
var canaries = []canaryRow{
	{"NewAnonymousConnection", "xpc_connection_create", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if c != nil {
			c.Cancel()
		}
		return err
	}},
	{"(*Connection).Activate", "xpc_connection_activate", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if err != nil {
			return err
		}
		defer c.Cancel()
		return c.Activate()
	}},
	{"(*Connection).Cancel", "xpc_connection_cancel", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if err != nil {
			return err
		}
		return c.Cancel()
	}},
	{"(*Connection).Endpoint", "xpc_endpoint_create", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if err != nil {
			return err
		}
		defer c.Cancel()
		if err := c.Activate(); err != nil {
			return err
		}
		_, err = c.Endpoint()
		return err
	}},
	{"(*Connection).SetTargetQueue", "xpc_connection_set_target_queue", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if err != nil {
			return err
		}
		defer c.Cancel()
		return c.SetTargetQueue(dispatch.QueueCreate("xpc-canary"))
	}},
	{"(*Connection).SetPeerRequirement", "xpc_connection_set_peer_requirement", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if err != nil {
			return err
		}
		defer c.Cancel()
		r, err := NewSameTeamRequirement()
		if err != nil {
			return err
		}
		defer r.Close()
		return c.SetPeerRequirement(r)
	}},
	{"(*Connection).SetPeerCodeSigningRequirement", "xpc_connection_set_peer_code_signing_requirement", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if err != nil {
			return err
		}
		defer c.Cancel()
		return c.SetPeerCodeSigningRequirement("identifier \"com.example\"")
	}},
	{"(*Connection).SetPeerEntitlementExistsRequirement", "xpc_connection_set_peer_entitlement_exists_requirement", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if err != nil {
			return err
		}
		defer c.Cancel()
		return c.SetPeerEntitlementExistsRequirement("com.example.entitlement")
	}},
	{"(*Connection).SetPeerPlatformIdentityRequirement", "xpc_connection_set_peer_platform_identity_requirement", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if err != nil {
			return err
		}
		defer c.Cancel()
		return c.SetPeerPlatformIdentityRequirement("com.example")
	}},
	{"(*Connection).SetPeerTeamIdentityRequirement", "xpc_connection_set_peer_team_identity_requirement", func(t *testing.T) error {
		c, err := connectionCanary(t)
		if err != nil {
			return err
		}
		defer c.Cancel()
		return c.SetPeerTeamIdentityRequirement("com.example")
	}},
	{"NewSameTeamRequirement", "xpc_peer_requirement_create_team_identity", func(t *testing.T) error {
		r, err := NewSameTeamRequirement()
		closeRequirement(t, r)
		return err
	}},
	{"NewSameTeamSignedAsRequirement", "xpc_peer_requirement_create_team_identity", func(t *testing.T) error {
		r, err := NewSameTeamSignedAsRequirement("com.example.signing")
		closeRequirement(t, r)
		return err
	}},
	{"NewPlatformBinaryRequirement", "xpc_peer_requirement_create_platform_identity", func(t *testing.T) error {
		r, err := NewPlatformBinaryRequirement()
		closeRequirement(t, r)
		return err
	}},
	{"NewPlatformBinarySignedAsRequirement", "xpc_peer_requirement_create_platform_identity", func(t *testing.T) error {
		r, err := NewPlatformBinarySignedAsRequirement("com.example.signing")
		closeRequirement(t, r)
		return err
	}},
	{"NewEntitlementExistsRequirement", "xpc_peer_requirement_create_entitlement_exists", func(t *testing.T) error {
		r, err := NewEntitlementExistsRequirement("com.example.entitlement")
		closeRequirement(t, r)
		return err
	}},
	{"NewEntitlementMatchesRequirement", "xpc_peer_requirement_create_entitlement_matches_value", func(t *testing.T) error {
		r, err := NewEntitlementMatchesRequirement("com.example.entitlement", true)
		closeRequirement(t, r)
		return err
	}},
	{"NewLightweightCodeRequirement", "xpc_peer_requirement_create_lwcr", func(t *testing.T) error {
		r, err := NewLightweightCodeRequirement(Dictionary{"$or": true})
		closeRequirement(t, r)
		return err
	}},
	{"(*PeerRequirement).Close", "xpc_release", func(t *testing.T) error {
		r, err := NewSameTeamRequirement()
		if err != nil || r == nil {
			t.Skipf("cannot build a requirement to close: %v", err)
		}
		return r.Close()
	}},

	// The resource types are the second family that can be proven. Unlike a
	// session or a listener, a boxed descriptor and a shared region can be
	// built from nothing but a temp file and an mmap, so the negative
	// control runs the real call path rather than a fabricated pointer.
	{"NewFileDescriptor", "xpc_fd_create", func(t *testing.T) error {
		f, err := NewFileDescriptor(canaryFD(t))
		closeFileDescriptor(f)
		return err
	}},
	{"(*FileDescriptor).Dup", "xpc_fd_dup", func(t *testing.T) error {
		f, err := NewFileDescriptor(canaryFD(t))
		if err != nil || f == nil {
			t.Skipf("cannot box a descriptor to dup: %v", err)
		}
		defer closeFileDescriptor(f)
		fd, err := f.Dup()
		if err == nil {
			syscall.Close(fd)
		}
		return err
	}},
	{"(*FileDescriptor).Close", "xpc_release", func(t *testing.T) error {
		f, err := NewFileDescriptor(canaryFD(t))
		if err != nil || f == nil {
			t.Skipf("cannot box a descriptor to close: %v", err)
		}
		return f.Close()
	}},
	{"NewSharedMemory", "xpc_shmem_create", func(t *testing.T) error {
		s, m, err := NewSharedMemory(os.Getpagesize())
		closeSharedMemory(s, m)
		return err
	}},
	{"(*SharedMemory).Map", "xpc_shmem_map", func(t *testing.T) error {
		s, m, err := NewSharedMemory(os.Getpagesize())
		if err != nil || s == nil {
			t.Skipf("cannot build a region to map: %v", err)
		}
		defer closeSharedMemory(s, m)
		got, err := s.Map()
		if got != nil {
			_ = got.Close()
		}
		return err
	}},
	{"(*SharedMemory).Close", "xpc_release", func(t *testing.T) error {
		s, m, err := NewSharedMemory(os.Getpagesize())
		if err != nil || s == nil {
			t.Skipf("cannot build a region to close: %v", err)
		}
		if m != nil {
			_ = m.Close()
		}
		return s.Close()
	}},
	{"CopyValue", "xpc_copy", func(t *testing.T) error {
		_, err := CopyValue("canary")
		return err
	}},
	{"Equal", "xpc_equal", func(t *testing.T) error {
		_, err := Equal("canary", "canary")
		return err
	}},
	{"Hash", "xpc_hash", func(t *testing.T) error {
		_, err := Hash("canary")
		return err
	}},
	{"CurrentDate", "xpc_date_create_from_current", func(t *testing.T) error {
		_, err := CurrentDate()
		return err
	}},
	{"Transaction", "xpc_transaction_begin", func(t *testing.T) error {
		end, err := Transaction()
		if end != nil {
			end()
		}
		return err
	}},
}

// canaryFD returns a descriptor that is certainly valid. It is a temp file
// rather than a standard descriptor so that a poisoned run cannot be
// confused by a test harness that has redirected them.
func canaryFD(t *testing.T) int {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "xpc-canary")
	if err != nil {
		t.Skipf("cannot create a file to box: %v", err)
	}
	t.Cleanup(func() { f.Close() })
	return int(f.Fd())
}

func closeFileDescriptor(f *FileDescriptor) {
	if f != nil {
		_ = f.Close()
	}
}

func closeSharedMemory(s *SharedMemory, m *Mapping) {
	if m != nil {
		_ = m.Close()
	}
	if s != nil {
		_ = s.Close()
	}
}

// closeRequirement releases a requirement built by a canary row, ignoring the
// nil case a poisoned run produces.
func closeRequirement(t *testing.T, r *PeerRequirement) {
	t.Helper()
	if r != nil {
		_ = r.Close()
	}
}

func connectionCanary(t *testing.T) (*Connection, error) {
	t.Helper()
	return NewAnonymousConnection(dispatch.QueueCreate("xpc-canary"), func(Dictionary) {})
}

// canaryUnmeasured records entry points whose guard is not proven to fire, and
// why. These are UNMEASURED, not clean: the guard exists in the source and is
// derived from the emitted symbol set, but no test has forced it to return
// non-nil. TestRawReachEveryEntryPointHasACanary fails if a newly emitted
// entry point appears in neither table.
var canaryUnmeasured = map[string]string{
	"(*Listener).Activate":                              "needs a live listener; the guard sits behind l.raw != nil and the negative control would activate a real listener",
	"(*Listener).Cancel":                                "needs a live listener; returns nothing, so a poisoned run has no observable error either",
	"(*Session).Activate":                               "needs a live session",
	"(*Session).Cancel":                                 "nonreporting: Cancel returns nothing, and the negative control would cancel a real session",
	"(*Session).Notify":                                 "needs a live session",
	"(*Session).NotifyDictionary":                       "needs a live session",
	"(*Session).Call":                                   "needs a live session; the negative control would block on a real peer",
	"(*Session).CallDictionary":                         "needs a live session; see above",
	"(*Session).SetCancellationHandler":                 "needs a live session",
	"(*Session).SetIncomingMessageHandler":              "needs a live session",
	"(*Session).SetPeerRequirement":                     "needs a live inactive session",
	"(*Session).SetTargetQueue":                         "needs a live inactive session",
	"(ReceivedMessage).Decode":                          "decodeMessage reaches raw calls only through (ReceivedMessage).Dictionary, whose guard is itself UNMEASURED for the same reason: it needs a received message from a live peer",
	"(ReceivedMessage).Dictionary":                      "needs a received message from a live peer",
	"(ReceivedMessage).SenderSatisfies":                 "needs a received message from a live peer",
	"DialMachService":                                   "touches launchd; belongs to the xpclive suite, not the default one",
	"DialXPCService":                                    "touches launchd; belongs to the xpclive suite, not the default one",
	"NewAnonymousListener":                              "the negative control would create and activate a real listener",
	"NewServiceListener":                                "the negative control would register a real service listener",
	"PeerRequirementFromHandle":                         "nonreporting: returns nil on guard failure, and the negative control would retain a fabricated handle",
	"(*Listener).SetPeerCodeSigningRequirement":         "needs a live inactive listener; the negative control would create a real listener",
	"(*Listener).String":                                "needs a live listener; String intentionally hides availability errors",
	"(*Session).SetPeerCodeSigningRequirement":          "needs a live inactive session",
	"(*Session).String":                                 "needs a live session; String intentionally hides availability errors",
	"SetEventStreamHandler":                             "process-global and cannot be safely repeated by a negative control",
	"ActivateSocket":                                    "touches launchd and can be called only once per socket name; belongs to the xpclive suite",
	"NewAnonymousConnection":                            "creates an anonymous listener and needs an active endpoint-transfer rig",
	"NewConnectionFromEndpoint":                         "needs a live endpoint from another process",
	"(*Connection).Activate":                            "needs a live connection",
	"(*Connection).Cancel":                              "nonreporting and needs a live connection",
	"(*Connection).Endpoint":                            "needs an active anonymous listener connection",
	"(*Connection).CallDictionary":                      "needs a live peer and may block",
	"(*Connection).SendDictionary":                      "needs a live peer",
	"(*Connection).Suspend":                             "needs a live connection",
	"(*Connection).Resume":                              "needs a live connection",
	"(*Connection).PID":                                 "peer credentials require a connected peer",
	"(*Connection).EUID":                                "peer credentials require a connected peer",
	"(*Connection).EGID":                                "peer credentials require a connected peer",
	"(*Connection).AuditSessionID":                      "peer credentials require a connected peer",
	"(*Connection).Name":                                "needs a live connection",
	"(*Connection).InvalidationReason":                  "needs a live invalidated connection",
	"(*Connection).SetTargetQueue":                      "needs a live inactive connection",
	"(*Connection).SetPeerRequirement":                  "needs a live inactive connection",
	"(*Connection).SetPeerCodeSigningRequirement":       "needs a live inactive connection",
	"NewMachServiceConnection":                          "touches launchd and needs a registered Mach service",
	"(*Connection).SetPeerEntitlementExistsRequirement": "needs a live inactive connection",
	"(*Connection).SetPeerPlatformIdentityRequirement":  "needs a live inactive connection",
	"(*Connection).SetPeerTeamIdentityRequirement":      "needs a live inactive connection",
	"RegisterActivity":                                  "process-global registration needs a unique system activity environment",
	"UnregisterActivity":                                "process-global registration needs a matching registration",
	"(*Activity).Close":                                 "needs an activity delivered by the scheduler",
	"(*Activity).Criteria":                              "needs an activity delivered by the scheduler",
	"(*Activity).ShouldDefer":                           "needs an activity delivered by the scheduler",
}
