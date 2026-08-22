// Command addservice is an XPC service that adds two numbers.
//
// It is the Go equivalent of Apple's low-level XPC service template, the one
// whose main.c calls xpc_listener_create and replies from an incoming message
// handler. The service name and the message keys match that template, so this
// binary can be dropped into the same .xpc bundle:
//
//	dev.tmc.sample-xpc-service-ll
//	{"firstNumber": int64, "secondNumber": int64} -> {"result": int64}
//
// An XPC service is launched by launchd on demand, not run from a shell. It
// lives in a .xpc bundle inside Contents/XPCServices of a host application,
// and only a process inside that application can reach it. examples/xpc/
// build-bundle.sh assembles both sides:
//
//	./examples/xpc/build-bundle.sh ~/tmp/AddDemo.app
//	~/tmp/AddDemo.app/Contents/MacOS/AddDemo -first 23 -second 19
package main

import (
	"log"

	"github.com/tmc/apple/xpc"
)

const serviceName = "dev.tmc.sample-xpc-service-ll"

// request and reply carry the wire keys as struct tags. The codec reads the
// xpc tag, so the Go field names do not have to match the protocol.
type request struct {
	First  int64 `xpc:"firstNumber"`
	Second int64 `xpc:"secondNumber"`
}

type reply struct {
	Result int64 `xpc:"result"`
}

func main() {
	// Inactive defers the listener's first delivery until Activate, so accept
	// is installed before any peer arrives. Without it the listener is already
	// active and the Activate below is API misuse: XPC traps the process.
	opts := xpc.ListenerOptions{Inactive: true}
	listener, err := xpc.NewServiceListener(serviceName, opts, accept)
	if err != nil {
		log.Fatalf("create listener: %v", err)
	}
	if err := listener.Activate(); err != nil {
		log.Fatalf("activate listener: %v", err)
	}
	log.Printf("listening on %s", serviceName)

	// Activate does not block. The C template calls dispatch_main here;
	// incoming sessions arrive on XPC's own queues, so main only has to stay
	// alive and keep the listener reachable.
	select {}
}

// accept runs once per incoming peer. Returning Accept installs the message
// handler; returning Reject refuses the peer with a reason.
func accept(req xpc.IncomingSessionRequest) xpc.IncomingDecision {
	return req.Accept(handle, func(err xpc.RichError) {
		log.Printf("session cancelled: %v", err)
	})
}

// handle answers one message. Returning a value replies with it; returning an
// error replies with {"error": ...}; returning nil, nil sends no reply.
func handle(msg xpc.ReceivedMessage) (any, error) {
	var req request
	if err := msg.Decode(&req); err != nil {
		return nil, err
	}
	return reply{Result: req.First + req.Second}, nil
}
