// Command goservice is the Go half of the Swift/Go XPC interoperability
// examples: a Mach-service listener that a Swift XPCSession client talks to.
//
// It answers a handful of ops chosen to probe the edges of the Go binding
// rather than to compute anything:
//
//	describe   report the Go type and value the codec produced for every key
//	typezoo    reply with one value of every type the Go codec can encode
//	silent     return (nil, nil), which sends no reply at all
//	fail       return an error, which replies with {"error": string}
//	errorkey   reply with a legitimate payload that contains an "error" key
//
// See ../README.md for how it is installed and run.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/xpc"
)

func main() {
	name := flag.String("service", "", "Mach service name to serve")
	flag.Parse()
	if *name == "" {
		fmt.Fprintln(os.Stderr, "goservice: -service is required")
		os.Exit(2)
	}
	listener, err := xpc.NewServiceListener(*name, xpc.ListenerOptions{
		ForceMach: true,
		Inactive:  true,
	}, accept)
	if err != nil {
		log.Fatalf("create listener %s: %v", *name, err)
	}
	if err := listener.Activate(); err != nil {
		log.Fatalf("activate listener: %v", err)
	}
	log.Printf("serving %s", *name)
	select {}
}

// accept probes the accept-inactive question. Go's IncomingDecision always
// activates the peer when it is applied, and there is no accept-but-inactive
// form -- but AcceptSession hands back the *Session before the decision is
// returned, and the decision is not applied until then. That is the window
// Swift's accept { (session: XPCSession) in ... } closure makes explicit,
// and the log line below records whether Go's inactive-only setters accept it.
func accept(req xpc.IncomingSessionRequest) xpc.IncomingDecision {
	decision, peer := req.AcceptSession(handle, func(err xpc.RichError) {
		log.Printf("session cancelled: %v", err)
	})
	queue := dispatch.QueueCreate("goservice.peer")
	if err := peer.SetTargetQueue(queue); err != nil {
		log.Printf("configure peer before activation: REFUSED: %v", err)
	} else {
		log.Printf("configure peer before activation: accepted (SetTargetQueue)")
	}
	return decision
}

func handle(msg xpc.ReceivedMessage) (any, error) {
	in := msg.Dictionary()
	op, _ := in["op"].(string)
	log.Printf("op=%q keys=%d", op, len(in))
	switch op {
	case "describe":
		return xpc.Dictionary{"report": Describe(in)}, nil
	case "typezoo":
		return TypeZoo(), nil
	case "silent":
		// The binding's documented "do not reply" form. Swift's handler
		// spells the same thing as returning a nil Output.
		return nil, nil
	case "fail":
		// The binding turns this into the dictionary {"error": "<text>"}.
		// Swift has no such convention: its client sees an ordinary reply.
		return nil, fmt.Errorf("goservice refused op %q on purpose", op)
	case "errorkey":
		// A legitimate payload that happens to carry an "error" key, which
		// is indistinguishable on the wire from the case above.
		return xpc.Dictionary{
			"error":  "this is real data, not a failure",
			"status": "ok",
		}, nil
	default:
		return nil, fmt.Errorf("unknown op %q", op)
	}
}

// Describe renders what the Go codec produced for each key of a received
// dictionary: the Go static type and a readable value. Keys are sorted so the
// output is stable.
func Describe(d xpc.Dictionary) string {
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
		// xpc_type_null and "key absent" are the same thing in Go.
		return "nil\t<null>"
	case []byte:
		return fmt.Sprintf("[]byte\t%d bytes %x", len(t), t)
	case string:
		return fmt.Sprintf("string\t%q", t)
	case xpc.Dictionary:
		return fmt.Sprintf("xpc.Dictionary\t%s", strings.ReplaceAll(strings.TrimSpace(Describe(t)), "\n", " | "))
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

// TypeZoo is one value of every type the Go codec can encode, for the Swift
// side to report on. The date and uuid carry the same values the Swift zoo
// sends, so the two directions can be read against each other. fd and shmem
// are absent because they are resources the codec does not originate; that
// absence is the finding, and a Swift peer's fd arrives here as an
// xpc.Unsupported rather than as something that looks like data.
func TypeZoo() xpc.Dictionary {
	return xpc.Dictionary{
		"bool":   true,
		"int64":  int64(-42),
		"uint64": uint64(1) << 63,
		"double": 3.5,
		"string": "héllo",
		"data":   []byte{0xde, 0xad, 0xbe, 0xef},
		"null":   nil,
		"array":  []any{int64(1), "two", false},
		"dict":   xpc.Dictionary{"nested": int64(7)},
		"date":   time.Unix(1699999999, 0),
		"uuid": xpc.UUID{
			0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
			0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		},
	}
}
