// Command xpcclient is the client side of the xpc package's live
// peer-requirement tests.
//
// The live suite builds this binary twice and codesigns the copies
// differently: one signed with a test entitlement (which satisfies the
// service's peer requirement) and one without (which must be dropped by the
// listener). Launching a separate, signed process is the only way to prove
// listener-side enforcement: the listener checks the peer's signature, and
// the peer is this process.
//
// Modes:
//
//	-roundtrip -service NAME    dial, send an add, print "REPLIED <sum>", exit 0
//	-expect-reject -service NAME
//	                            dial, send an add. If a reply arrives, print
//	                            "REPLIED" and exit 1 (the requirement was
//	                            ignored). If the send fails, print the error and
//	                            exit 0 (enforcement observed). If the peer drops
//	                            the request, the send blocks and the parent
//	                            kills this process after a bounded window,
//	                            which it also treats as the named rejection.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tmc/apple/xpc"
)

type call struct {
	Op     string `xpc:"op"`
	First  int64  `xpc:"firstNumber"`
	Second int64  `xpc:"secondNumber"`
}

type result struct {
	Sum int64 `xpc:"sumValue"`
}

func main() {
	mode := flag.String("mode", "roundtrip", "roundtrip or expect-reject")
	service := flag.String("service", "", "mach service name")
	flag.Parse()
	if *service == "" {
		log.Fatal("missing -service")
	}
	log.SetPrefix("xpcclient: ")
	log.SetFlags(0)

	session, err := xpc.DialMachService(*service, xpc.SessionOptions{})
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer session.Cancel()

	// SENT marks the last point before the reply wait, so the parent can
	// distinguish "dropped and hung" (no further output, killed by timeout)
	// from "died before sending" (no SENT at all).
	fmt.Println("SENT")

	msg, err := session.Call(context.Background(), call{Op: "add", First: 21, Second: 21})
	if err != nil {
		if *mode == "expect-reject" {
			// The peer failed us, which is the enforcement we expected.
			fmt.Printf("REJECTED %v\n", err)
			return
		}
		log.Fatalf("send: %v", err)
	}
	var got result
	if err := msg.Decode(&got); err != nil {
		if *mode == "expect-reject" {
			fmt.Printf("REJECTED decode: %v\n", err)
			return
		}
		log.Fatalf("decode: %v", err)
	}
	fmt.Printf("REPLIED %d\n", got.Sum)
	if *mode == "expect-reject" {
		// A reply means the listener accepted a peer that should have been
		// dropped.
		log.Fatalf("got a reply %d from a listener that should have rejected us", got.Sum)
	}
	if got.Sum != 42 {
		log.Fatalf("sum = %d, want 42", got.Sum)
	}
	os.Exit(0)
}
