// Command addclient calls the addservice XPC service.
//
// It is the Go equivalent of the client sketch in Apple's low-level XPC
// service template: create a session against the service, send a dictionary,
// read the reply, cancel the session.
//
// The session only resolves if the named service is reachable from this
// process, which for an XPC service means running as the host application that
// contains the bundle. examples/xpc/build-bundle.sh installs this binary as
// that application's executable:
//
//	./examples/xpc/build-bundle.sh ~/tmp/AddDemo.app
//	~/tmp/AddDemo.app/Contents/MacOS/AddDemo -first 23 -second 19
//
// Run from a plain shell it fails at Dial, which is the correct result rather
// than a hang.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/tmc/apple/xpc"
)

const serviceName = "dev.tmc.sample-xpc-service-ll"

type request struct {
	First  int64 `xpc:"firstNumber"`
	Second int64 `xpc:"secondNumber"`
}

type reply struct {
	Result int64 `xpc:"result"`
}

func main() {
	first := flag.Int64("first", 23, "first addend")
	second := flag.Int64("second", 19, "second addend")
	raw := flag.Bool("raw", false, "send an xpc.Dictionary instead of a struct")
	flag.Parse()

	session, err := xpc.DialXPCService(serviceName, xpc.SessionOptions{})
	if err != nil {
		log.Fatalf("dial %s: %v", serviceName, err)
	}
	defer session.Cancel()

	// A session is live as soon as it is created. Activate is only for
	// sessions created with SessionOptions{Inactive: true}, which exists so a
	// handler can be installed before any message arrives.

	if *raw {
		// The dictionary form, which is what the C template writes by hand
		// with xpc_dictionary_set_int64.
		msg := xpc.Dictionary{"firstNumber": *first, "secondNumber": *second}
		got, err := session.CallDictionary(context.Background(), msg)
		if err != nil {
			log.Fatalf("send: %v", err)
		}
		fmt.Println(got.Dictionary()["result"])
		return
	}

	got, err := session.Call(context.Background(), request{First: *first, Second: *second})
	if err != nil {
		log.Fatalf("send: %v", err)
	}
	var out reply
	if err := got.Decode(&out); err != nil {
		log.Fatalf("decode reply: %v", err)
	}
	fmt.Println(out.Result)
}
