package xpc

import "testing"

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
	"xpcCancelTrampoline|handler(...)":        "caller-supplied CancellationHandler; not package code",
	"xpcIncomingTrampoline|handler(...)":      "caller-supplied MessageHandler; not package code",
	"xpcIncomingTrampoline|err.Error(...)":    "error interface method; RichError.Error reaches no raw_ call",
	"xpcIncomingTrampoline|encErr.Error(...)": "error interface method on encodeMessage's error; reaches no raw_ call",
	"(RichError).Error|e.cause.Error(...)":    "wrapped error's Error method; not package code",
	"newListener|incoming(...)":               "caller-supplied incoming handler; not package code",
	"newRequirement|create(...)":              "the create closure is supplied by each New*Requirement constructor and its raw call is recorded at the constructor",
	"targetQueuePointer|queue.Handle(...)":    "dispatch.Queue.Handle lives in another package and cannot reach an xpc raw_ call",
	"decodeJSONPayload|dec.Decode(...)":       "encoding/json's Decoder.Decode; it decodes bytes into an any and reaches no xpc raw_ call",
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
}

// closeRequirement releases a requirement built by a canary row, ignoring the
// nil case a poisoned run produces.
func closeRequirement(t *testing.T, r *PeerRequirement) {
	t.Helper()
	if r != nil {
		_ = r.Close()
	}
}

// canaryUnmeasured records entry points whose guard is not proven to fire, and
// why. These are UNMEASURED, not clean: the guard exists in the source and is
// derived from the emitted symbol set, but no test has forced it to return
// non-nil. TestRawReachEveryEntryPointHasACanary fails if a newly emitted
// entry point appears in neither table.
var canaryUnmeasured = map[string]string{
	"(*Listener).Activate":                 "needs a live listener; the guard sits behind l.raw != nil and the negative control would activate a real listener",
	"(*Listener).Cancel":                   "needs a live listener; returns nothing, so a poisoned run has no observable error either",
	"(*Session).Activate":                  "needs a live session",
	"(*Session).Cancel":                    "nonreporting: Cancel returns nothing, and the negative control would cancel a real session",
	"(*Session).Notify":                    "needs a live session",
	"(*Session).NotifyDictionary":          "needs a live session",
	"(*Session).Call":                      "needs a live session; the negative control would block on a real peer",
	"(*Session).CallDictionary":            "needs a live session; see above",
	"(*Session).SetCancellationHandler":    "needs a live session",
	"(*Session).SetIncomingMessageHandler": "needs a live session",
	"(*Session).SetPeerRequirement":        "needs a live inactive session",
	"(*Session).SetTargetQueue":            "needs a live inactive session",
	"(ReceivedMessage).Decode":             "decodeMessage reaches raw calls only through (ReceivedMessage).Dictionary, whose guard is itself UNMEASURED for the same reason: it needs a received message from a live peer",
	"(ReceivedMessage).Dictionary":         "needs a received message from a live peer",
	"(ReceivedMessage).SenderSatisfies":    "needs a received message from a live peer",
	"DialMachService":                      "touches launchd; belongs to the xpclive suite, not the default one",
	"DialXPCService":                       "touches launchd; belongs to the xpclive suite, not the default one",
	"NewAnonymousListener":                 "the negative control would create and activate a real listener",
	"NewServiceListener":                   "the negative control would register a real service listener",
	"PeerRequirementFromHandle":            "nonreporting: returns nil on guard failure, and the negative control would retain a fabricated handle",
}
