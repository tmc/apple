// Copyright 2026 The tmc/apple Authors. All rights reserved.

package xpc

import (
	"reflect"
	"testing"
)

// Cross-process endpoint transfer is UNMEASURED, and it is blocked before the
// rig can be built. This file records where, and fails when that stops being
// true.
//
// The rig would be: one process creates an anonymous listener connection,
// exports it with xpc_endpoint_create, hands the endpoint to a second process
// over a mach service, and the second process calls
// xpc_connection_create_from_endpoint and completes a round trip. Every symbol
// it needs is bound. The first step is not expressible.
//
// An anonymous listener is xpc_connection_create(NULL, queue). The generated
// binding types name as a Go string, and purego converts a Go string to a
// non-NULL C string, so NULL cannot be passed. Measured against libxpc's own
// description of the resulting object (a standalone purego probe, macOS 26.6):
//
//	NULL name  -> { name = (anonymous), listener = true,  pid = <self> }
//	""         -> { name = ,            listener = false, pid = 0 }
//
// The empty string does not degrade to an anonymous listener; it makes a client
// connection to a service named "". And xpc_endpoint_create on that non-listener
// returns a non-NULL <endpoint> rather than an error — the undefined behaviour
// xpc/endpoint.h:12-14 documents produces a plausible object. A rig built on the
// current binding would therefore hand the second process an endpoint that
// cannot work, and nothing in the transfer would report it. That is why this is
// recorded as blocked rather than attempted: a green rig would have been the
// wrong answer.
//
// Unblocking it is a binding change, not a test change: name must become a
// pointer type that can be nil. Two paths avoid it entirely, both larger:
//
//   - xpc_connection_create_mach_service(name, q, XPC_CONNECTION_MACH_SERVICE_LISTENER)
//     is a listener connection with a real name and needs no NULL, but the name
//     must be registered with launchd by a second LaunchAgent, and the package
//     has no classic-connection server loop.
//   - xpc_listener_create_endpoint is dlsym-present and header-absent. It is out
//     of bounds: a guessed C signature compiles and calls the real symbol with
//     argument registers never loaded, and this project is 3-for-6 wrong on
//     those.
//
// Incidental finding from the same probe, recorded because it is a live process
// abort with no documentation in this package: xpc_release on a connection that
// was never activated traps (SIGTRAP inside libxpc, unrecoverable from Go).
// xpc_connection_cancel on the same object is fine, and cancel-then-release
// without release is clean — the trap is attributable to the release, measured
// by removing it and getting exit 0.
func TestEndpointTransferIsBlockedByConnectionCreateSignature(t *testing.T) {
	fn := reflect.TypeOf(rawfn_xpc_connection_create)
	if fn == nil || fn.Kind() != reflect.Func {
		t.Fatalf("rawfn_xpc_connection_create is not a func value (%v): this gate is measuring the wrong thing", fn)
	}
	if fn.NumIn() == 0 {
		t.Fatalf("rawfn_xpc_connection_create takes no arguments: the binding is not the one this gate describes")
	}
	name := fn.In(0)
	if name.Kind() != reflect.String {
		t.Errorf("rawfn_xpc_connection_create's name parameter is %v, no longer a string: NULL may now be expressible, "+
			"so build the two-process endpoint-transfer rig and delete this test", name)
		return
	}
	t.Logf("UNMEASURED: cross-process endpoint transfer. rawfn_xpc_connection_create's name is a %v, "+
		"which cannot be NULL, so an anonymous listener cannot be created and no endpoint worth transferring exists", name)
}
