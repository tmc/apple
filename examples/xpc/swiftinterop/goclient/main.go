// Command goclient is the Go half of direction A: an xpc.Session talking to a
// Mach service served by Swift's XPCListener.
//
// The interesting output is -op typezoo, which asks the Swift service to send
// one value of every XPC type and prints what the Go codec turned each into.
// Types the Go codec has no case for arrive as their copy_description string.
//
// See ../README.md for how it is run.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/tmc/apple/xpc"
)

func main() {
	name := flag.String("service", "", "Mach service name to dial")
	op := flag.String("op", "describe", "op to send")
	flag.Parse()
	if *name == "" {
		fmt.Fprintln(os.Stderr, "goclient: -service is required")
		os.Exit(2)
	}
	if err := call(*name, *op); err != nil {
		log.Fatal(err)
	}
}

func call(name, op string) error {
	// Inactive: dialing an absent Mach service with an active session fails
	// at create time, which would hide the real error behind a dial failure.
	session, err := xpc.DialMachService(name, xpc.SessionOptions{Inactive: true})
	if err != nil {
		return fmt.Errorf("dial %s: %w", name, err)
	}
	defer session.Cancel()
	if err := session.Activate(); err != nil {
		return fmt.Errorf("activate: %w", err)
	}

	if op == "endpointrelay" {
		// Ask the Swift service for its own listener endpoint, then hand it
		// straight back. Go cannot build a session from one -- there is no
		// bound equivalent of Swift's XPCSession(endpoint:) -- but it can
		// carry one, and the Swift side proves whether it arrived intact.
		got, err := session.CallDictionary(context.Background(), xpc.Dictionary{"op": "typezoo:endpoint"})
		if err != nil {
			return fmt.Errorf("fetch endpoint: %w", err)
		}
		ep, ok := got.Dictionary()["endpoint"].(xpc.Endpoint)
		if !ok {
			return fmt.Errorf("endpoint key decoded as %T, not xpc.Endpoint", got.Dictionary()["endpoint"])
		}
		fmt.Printf("Go holds an xpc.Endpoint (handle=%#x) and has no API to dial it\n", ep.Handle())
		back, err := session.CallDictionary(context.Background(), xpc.Dictionary{"op": "echoendpoint", "endpoint": ep})
		if err != nil {
			return fmt.Errorf("relay endpoint: %w", err)
		}
		if report, ok := back.Dictionary()["report"].(string); ok {
			fmt.Print(report)
			return nil
		}
		fmt.Print(describeDictionary(back.Dictionary()))
		return nil
	}

	msg := xpc.Dictionary{"op": op}
	if op == "describe" {
		// Everything the Go codec can encode, for the Swift side to report on.
		msg["bool"] = true
		msg["int64"] = int64(-42)
		msg["uint64"] = uint64(1) << 63
		msg["double"] = 3.5
		msg["string"] = "héllo"
		msg["data"] = []byte{0xde, 0xad, 0xbe, 0xef}
		msg["null"] = nil
		msg["array"] = []any{int64(1), "two", false}
		msg["dict"] = xpc.Dictionary{"nested": int64(7)}
	}

	if op == "silent" {
		// The peer sends no reply. Ask for one anyway, with a bounded wait:
		// a Call on context.Background would block forever, which proves
		// nothing about whether the reply is merely late. The deadline is
		// caller-side only: libxpc has no per-message timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		got, err := session.CallDictionary(ctx, msg)
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			fmt.Println("silent op: timedOut=true reply callback never fired")
		case err != nil:
			fmt.Printf("silent op: timedOut=false reply callback fired with error: %v\n", err)
		default:
			fmt.Printf("silent op: timedOut=false reply callback fired: %s\n", strings.TrimSpace(describeDictionary(got.Dictionary())))
		}
		return nil
	}

	got, err := session.CallDictionary(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("send %q: %w", op, err)
	}
	reply := got.Dictionary()
	if report, ok := reply["report"].(string); ok && len(reply) == 1 {
		fmt.Print(report)
		return nil
	}
	fmt.Print(describeDictionary(reply))
	return nil
}

func describeDictionary(d xpc.Dictionary) string {
	keys := make([]string, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for _, k := range keys {
		fmt.Fprintf(&b, "%s\t%s\n", k, describeValue(d[k]))
	}
	return b.String()
}

func describeValue(v any) string {
	switch t := v.(type) {
	case nil:
		return "nil\t<null>"
	case []byte:
		return fmt.Sprintf("[]byte\t%d bytes %x", len(t), t)
	case string:
		return fmt.Sprintf("string\t%q", t)
	case xpc.Dictionary:
		return fmt.Sprintf("xpc.Dictionary\t%s", strings.ReplaceAll(strings.TrimSpace(describeDictionary(t)), "\n", " | "))
	case []any:
		parts := make([]string, len(t))
		for i, e := range t {
			parts[i] = strings.ReplaceAll(describeValue(e), "\t", " ")
		}
		return fmt.Sprintf("[]any\t[%s]", strings.Join(parts, ", "))
	case xpc.Endpoint:
		return fmt.Sprintf("xpc.Endpoint\thandle=%#x", t.Handle())
	default:
		return fmt.Sprintf("%T\t%v", v, v)
	}
}
