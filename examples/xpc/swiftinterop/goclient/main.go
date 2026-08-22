// Command goclient is the Go half of direction A: an xpc.Session talking to a
// Mach service served by Swift's XPCListener.
//
// The interesting output is -op typezoo, which asks the Swift service to send
// one value of every XPC type and prints what the Go codec turned each into.
// Descriptors and shared pages are read through rather than printed, because a
// resource that decodes is not yet a resource that works. Anything the codec
// does not model would arrive as an xpc.Unsupported; the zoo sends no such
// value, so that path is not exercised here.
//
// See ../README.md for how it is run.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
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
		peer, err := xpc.NewConnectionFromEndpoint(ep, func(xpc.Dictionary) {})
		if err != nil {
			return fmt.Errorf("dial endpoint: %w", err)
		}
		defer peer.Cancel()
		if err := peer.Activate(); err != nil {
			return fmt.Errorf("activate endpoint connection: %w", err)
		}
		through, err := peer.CallDictionary(context.Background(), xpc.Dictionary{"op": "ping"})
		if err != nil {
			return fmt.Errorf("call endpoint: %w", err)
		}
		if got, _ := through["pong"].(string); got != "reached through a relayed endpoint" {
			return fmt.Errorf("endpoint reply = %#v, want pong", through)
		}
		fmt.Printf("Go dialed xpc.Endpoint (handle=%#x) and received a reply\n", ep.Handle())
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
		msg["date"] = time.Unix(1699999999, 0)
		msg["uuid"] = xpc.UUID{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}

		// The two resource types. Go originates both here, which is the
		// direction that could not be attempted at all until the codec grew
		// cases for them: the Swift side reports whether what arrived is a
		// usable descriptor and a mapped page, or merely something shaped
		// like one.
		if fd, err := goFileDescriptor(); err != nil {
			log.Printf("cannot build a descriptor to send: %v", err)
		} else {
			defer fd.Close()
			msg["fd"] = fd
		}
		if shm, mapping, err := goSharedMemory(); err != nil {
			log.Printf("cannot build a shared region to send: %v", err)
		} else {
			defer shm.Close()
			defer mapping.Close()
			msg["shmem"] = shm
		}
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
	case *xpc.FileDescriptor:
		return describeFD(t)
	case *xpc.SharedMemory:
		return describeShmem(t)
	case xpc.Endpoint:
		return fmt.Sprintf("xpc.Endpoint\thandle=%#x", t.Handle())
	default:
		return fmt.Sprintf("%T\t%v", v, v)
	}
}

// describeFD proves the descriptor is usable rather than merely present: it
// duplicates it and reads what is on the other end. A description string
// cannot be read from.
func describeFD(f *xpc.FileDescriptor) string {
	fd, err := f.Dup()
	if err != nil {
		return fmt.Sprintf("*xpc.FileDescriptor\tDup failed: %v", err)
	}
	// os.NewFile takes ownership of fd, so Close here closes the dup and
	// nothing else; the boxed descriptor is unaffected.
	file := os.NewFile(uintptr(fd), "xpc-fd")
	defer file.Close()
	b, err := io.ReadAll(file)
	if err != nil {
		return fmt.Sprintf("*xpc.FileDescriptor\tdup=%d read failed: %v", fd, err)
	}
	return fmt.Sprintf("*xpc.FileDescriptor\tdup=%d contents=%q", fd, string(b))
}

// describeShmem maps the region and reads its prefix, for the same reason.
func describeShmem(s *xpc.SharedMemory) string {
	m, err := s.Map()
	if err != nil {
		return fmt.Sprintf("*xpc.SharedMemory\tMap failed: %v", err)
	}
	defer m.Close()
	prefix := m.Bytes
	if len(prefix) > 16 {
		prefix = prefix[:16]
	}
	return fmt.Sprintf("*xpc.SharedMemory\t%d bytes mapped, prefix=%q", len(m.Bytes), string(prefix))
}

// goFileDescriptor boxes a descriptor open on a file with known contents, so
// the receiver can prove it got something it can read.
func goFileDescriptor() (*xpc.FileDescriptor, error) {
	f, err := os.CreateTemp("", "goclient-fd-probe")
	if err != nil {
		return nil, err
	}
	// xpc_fd_create duplicates, so this process's own descriptor can go as
	// soon as it is boxed.
	defer os.Remove(f.Name())
	defer f.Close()
	if _, err := f.WriteString("GO FD PAYLOAD\n"); err != nil {
		return nil, err
	}
	if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}
	return xpc.NewFileDescriptor(int(f.Fd()))
}

// goSharedMemory allocates a page, writes a recognisable prefix and boxes it.
// The mapping is the sender's own view and is returned so the caller can
// unmap it after the message has gone.
func goSharedMemory() (*xpc.SharedMemory, *xpc.Mapping, error) {
	shm, m, err := xpc.NewSharedMemory(os.Getpagesize())
	if err != nil {
		return nil, nil, err
	}
	copy(m.Bytes, "GO SHMEM PAYLOAD")
	return shm, m, nil
}
