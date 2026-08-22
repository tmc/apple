// Command xpclivesvc is the service side of the xpc package's live test.
//
// launchd starts it on demand out of a LaunchAgent that the test writes; the
// test process is the client. It is a separate binary because a listener can
// only be created by a process launchd started as that service.
//
// The service names and the test entitlement come from environment variables
// that the agent's plist sets, so each test run gets its own names and its
// own job. The helper binary is codesigned with the test entitlement by the
// test before the job is bootstrapped, which is what lets the session-side
// tests drive requirements against a real signature.
//
// It serves three listeners:
//
//   - the plain service, which every round-trip test talks to;
//   - an entitlement listener that requires the test entitlement, created
//     with the requirement in the options and without Inactive set, so the
//     package must force the native inactive flag, install the requirement,
//     and activate on the package's side of the call. If that path traps, the
//     whole helper dies at startup and every live test fails;
//   - a target-queue listener whose accept handler asserts it is running on
//     the configured dispatch queue, which traps the helper if the TargetQueue
//     option is ignored.
//
// The protocol is one dictionary in, one dictionary out. Field names differ
// from the wire keys on purpose: a codec that ignores xpc tags decodes zeros.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/xpc"
)

type call struct {
	Op     string `xpc:"op"`
	First  int64  `xpc:"firstNumber"`
	Second int64  `xpc:"secondNumber"`
	Text   string `xpc:"someText"`
}

type result struct {
	Sum  int64  `xpc:"sumValue"`
	Echo string `xpc:"echoText"`
}

func main() {
	service := os.Getenv("XPCLIVE_SERVICE")
	entService := os.Getenv("XPCLIVE_ENT_SERVICE")
	tqService := os.Getenv("XPCLIVE_TQ_SERVICE")
	entitlement := os.Getenv("XPCLIVE_ENTITLEMENT")
	if service == "" || entService == "" || tqService == "" || entitlement == "" {
		log.Fatalf("missing XPCLIVE_SERVICE=%q XPCLIVE_ENT_SERVICE=%q XPCLIVE_TQ_SERVICE=%q XPCLIVE_ENTITLEMENT=%q",
			service, entService, tqService, entitlement)
	}

	// bootout sends SIGTERM. Exiting normally lets a -cover build flush its
	// counters to GOCOVERDIR; the default disposition would not.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sig
		os.Exit(0)
	}()

	// ForceMach is required for a Mach service name. Inactive defers the
	// first delivery until Activate, so the handler is installed before any
	// peer arrives; without it the listener is already active and Activate is
	// API misuse that traps the process.
	plain, err := xpc.NewServiceListener(service, xpc.ListenerOptions{ForceMach: true, Inactive: true}, accept)
	if err != nil {
		log.Fatalf("create plain listener: %v", err)
	}
	if err := plain.Activate(); err != nil {
		log.Fatalf("activate plain listener: %v", err)
	}

	// The entitlement listener is deliberately created WITHOUT Inactive in
	// the options: the package must force the native inactive flag, install
	// the requirement, and activate. Installing on an active object traps,
	// so a regression there kills the helper at startup.
	req, err := xpc.NewEntitlementExistsRequirement(entitlement)
	if err != nil {
		log.Fatalf("create requirement: %v", err)
	}
	defer req.Close()
	ent, err := xpc.NewServiceListener(entService, xpc.ListenerOptions{ForceMach: true, Requirement: req}, accept)
	if err != nil {
		log.Fatalf("create entitlement listener: %v", err)
	}
	_ = ent // already active; nothing to configure

	// The target-queue listener asserts, inside the helper process, that its
	// accept handler really runs on the configured queue. AssertCurrent
	// aborts the process on failure, and the parent treats helper death as
	// the named failing result instead of hanging.
	tq := dispatch.QueueCreate("xpc.live.target." + service)
	tqListener, err := xpc.NewServiceListener(tqService, xpc.ListenerOptions{ForceMach: true, TargetQueue: tq}, func(req xpc.IncomingSessionRequest) xpc.IncomingDecision {
		tq.AssertCurrent()
		return accept(req)
	})
	if err != nil {
		log.Fatalf("create target-queue listener: %v", err)
	}
	_ = tqListener

	log.Printf("serving %s (plain), %s (entitlement), %s (target queue)", service, entService, tqService)
	select {}
}

func accept(req xpc.IncomingSessionRequest) xpc.IncomingDecision {
	// AcceptSession hands back the peer session so the handler can cancel it.
	// The handler does not run until the decision is applied, so filling in
	// session after the call is in time.
	var session *xpc.Session
	decision, s := req.AcceptSession(func(msg xpc.ReceivedMessage) (any, error) {
		return handle(session, msg)
	}, func(err xpc.RichError) {
		log.Printf("session cancelled: %v", err)
	})
	session = s
	return decision
}

// handle answers one message. Returning a value replies with it, returning an
// error replies with {"error": ...}, and returning nil, nil sends no reply.
func handle(session *xpc.Session, msg xpc.ReceivedMessage) (any, error) {
	var req call
	if err := msg.Decode(&req); err != nil {
		return nil, err
	}
	switch req.Op {
	case "mirror":
		// Reply with the dictionary that actually arrived, so the client can
		// see what libxpc carried rather than what the client encoded.
		return msg.Dictionary(), nil
	case "slow":
		// Sleep firstNumber milliseconds before replying, so a client can
		// cancel a call that is still in flight. The reply is still sent: a
		// cancelled Call does not stop the peer.
		time.Sleep(time.Duration(req.First) * time.Millisecond)
		return result{Sum: req.First, Echo: "slow"}, nil
	case "add":
		return result{Sum: req.First + req.Second}, nil
	case "echo":
		return result{Sum: req.First, Echo: req.Text}, nil
	case "fail":
		return nil, fmt.Errorf("service refused op %q with %d", req.Op, req.First)
	case "silent":
		return nil, nil
	case "bye":
		// Cancel after the reply has been sent. The runtime sends the
		// returned value before the handler's block returns, so deferring
		// the cancel to another goroutine orders it after the reply. libxpc
		// offers no completion callback on xpc_session_send_message, so this
		// ordering is arranged, not guaranteed.
		defer func() { go session.Cancel() }()
		return result{Sum: req.First}, nil
	default:
		return nil, fmt.Errorf("unknown op %q", req.Op)
	}
}
