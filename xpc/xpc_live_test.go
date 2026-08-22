// Copyright 2026 The apple Authors.

//go:build xpclive

// This file drives a real XPC round trip against a service launchd starts, so
// it needs a logged-in GUI session and it mutates ~/Library/LaunchAgents. A
// build tag rather than an env-var skip keeps it out of the default build
// entirely: `go test ./xpc/` never compiles it, so it cannot be tripped into
// bootstrapping a job by accident.
//
//	go test -tags xpclive ./xpc/
package xpc_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/tmc/apple/xpc"
)

// call and result mirror the service's protocol. The Go field names differ
// from the wire keys so that a codec which ignores xpc tags decodes zeros.
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

// serviceName is the Mach service the whole file talks to. It is unique to
// this process so concurrent runs and stale jobs cannot collide.
var serviceName = fmt.Sprintf("dev.tmc.apple.xpc.live.%d", os.Getpid())

// entServiceName is served by a listener with a peer requirement installed
// via ListenerOptions, and tqServiceName by a listener whose accept handler
// asserts it runs on the configured target queue.
var entServiceName = serviceName + ".ent"
var tqServiceName = serviceName + ".tq"

// testEntitlement is embedded ad-hoc into the helper binaries by codesign.
// It is a development entitlement that ad-hoc signing may carry, and the
// service's peer requirement checks it, so a signed peer is admitted and an
// unsigned one is dropped.
const testEntitlement = "com.apple.security.get-task-allow"

// TestMain builds the helper, installs it as a LaunchAgent, and boots the job
// out again on the way out. One job serves every test: bootstrapping is the
// slow and the destructive part, so it happens once.
func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Fprintln(os.Stderr, "xpclive:", err)
		os.Exit(1)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	dir, err := os.MkdirTemp("", "xpclive")
	if err != nil {
		return 0, err
	}
	defer os.RemoveAll(dir)
	bin := filepath.Join(dir, "xpclivesvc")
	logFile := filepath.Join(dir, "service.log")
	// XPCLIVE_COVERDIR keeps the service's raw coverage data after the run,
	// for merging with the client-side profile. Without it the data is
	// summarised on the way out and discarded.
	covDir := os.Getenv("XPCLIVE_COVERDIR")
	if covDir == "" {
		covDir = filepath.Join(dir, "cov")
	}
	if err := os.MkdirAll(covDir, 0o755); err != nil {
		return 0, err
	}

	// -cover on the helper is what makes the listener side measurable: it is
	// a separate process, so `go test -cover` cannot see it. The helper exits
	// normally on SIGTERM so the counters reach GOCOVERDIR, and the helper's
	// own package must appear in -coverpkg: the runtime hook that writes the
	// counters is installed by main's instrumentation, so leaving it out
	// silently emits nothing.
	build := exec.Command("go", "build", "-cover",
		"-coverpkg=github.com/tmc/apple/xpc,github.com/tmc/apple/xpc/testdata/xpclive/service",
		"-o", bin,
		"github.com/tmc/apple/xpc/testdata/xpclive/service")
	if out, berr := build.CombinedOutput(); berr != nil {
		return 0, fmt.Errorf("build service: %v\n%s", berr, out)
	}

	// The peer-requirement fixture signs the helper with a test entitlement
	// and drives requirements against it. Signing is ad-hoc: the entitlement
	// chosen is one an ad-hoc signature may carry, so no keychain identity
	// or provisioning profile is needed. A failing codesign here means the
	// fixture cannot prove anything about peer enforcement, so it is fatal
	// rather than skipped.
	entPlist := filepath.Join(dir, "entitlements.plist")
	if err := os.WriteFile(entPlist, []byte(entitlementsPlist(testEntitlement)), 0o644); err != nil {
		return 0, err
	}
	codesign := func(out string, args ...string) error {
		all := append([]string{}, args...)
		all = append(all, out)
		if o, cerr := exec.Command("codesign", all...).CombinedOutput(); cerr != nil {
			return fmt.Errorf("codesign %s: %v: %s", out, cerr, o)
		}
		return nil
	}
	if err := codesign(bin, "-s", "-", "--entitlements", entPlist, "--force"); err != nil {
		return 0, err
	}

	// The client helper is built twice: one copy signed with the entitlement
	// (which must be admitted by the requirement listener) and one unsigned
	// (which must be dropped). They are separate processes because the
	// listener checks the peer's signature, and a session from the same
	// process would carry the test binary's signature.
	clientBin := filepath.Join(dir, "xpcclient")
	clientUnsigned := filepath.Join(dir, "xpcclient-unsigned")
	clientBuild := exec.Command("go", "build", "-o", clientBin,
		"github.com/tmc/apple/xpc/testdata/xpclive/client")
	if out, berr := clientBuild.CombinedOutput(); berr != nil {
		return 0, fmt.Errorf("build client: %v\n%s", berr, out)
	}
	if err := codesign(clientBin, "-s", "-", "--entitlements", entPlist, "--force"); err != nil {
		return 0, err
	}
	if err := os.WriteFile(clientUnsigned, mustReadFile(clientBin), 0o755); err != nil {
		return 0, err
	}
	// The copy inherits the signed binary's signature, so strip it: an
	// ad-hoc signature is removed by re-signing with no identity and no
	// entitlements, which is what an unsigned copy must be.
	if err := codesign(clientUnsigned, "-s", "-", "--force"); err != nil {
		return 0, err
	}
	// Package-level for the tests.
	clientPaths = [2]string{clientBin, clientUnsigned}

	home, err := os.UserHomeDir()
	if err != nil {
		return 0, err
	}
	plist := filepath.Join(home, "Library", "LaunchAgents", serviceName+".plist")
	if err := os.MkdirAll(filepath.Dir(plist), 0o755); err != nil {
		return 0, err
	}
	if err := os.WriteFile(plist, []byte(agentPlist(serviceName, entServiceName, tqServiceName, bin, logFile, covDir)), 0o644); err != nil {
		return 0, err
	}
	domain := fmt.Sprintf("gui/%d", os.Getuid())

	// Boot out before removing the plist: a plist removed while the job is
	// loaded leaves the job registered until the next login. A bootout can
	// race a launchd respawn of the on-demand job, so kill whatever is left
	// and boot out again; without the second pass a leftover job holds the
	// mach name for the next run.
	defer func() {
		launchctl("bootout", domain+"/"+serviceName)
		launchctl("kill", "SIGKILL", domain+"/"+serviceName)
		launchctl("bootout", domain+"/"+serviceName)
		if rerr := os.Remove(plist); rerr != nil && !os.IsNotExist(rerr) && err == nil {
			err = rerr
		}
		if b, rerr := os.ReadFile(logFile); rerr == nil && len(b) > 0 {
			fmt.Fprintf(os.Stderr, "service log:\n%s", b)
		}
		// bootout only asks the service to stop; its counters land a moment
		// after launchctl returns.
		for deadline := time.Now().Add(10 * time.Second); time.Now().Before(deadline); {
			if names, _ := filepath.Glob(filepath.Join(covDir, "covcounters.*")); len(names) > 0 {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		out, cerr := exec.Command("go", "tool", "covdata", "percent", "-i="+covDir).CombinedOutput()
		fmt.Fprintf(os.Stderr, "service-side coverage (err=%v):\n%s", cerr, out)
	}()

	if out, berr := exec.Command("launchctl", "bootstrap", domain, plist).CombinedOutput(); berr != nil {
		return 0, fmt.Errorf("launchctl bootstrap: %v: %s", berr, out)
	}
	return m.Run(), nil
}

func launchctl(args ...string) {
	if out, err := exec.Command("launchctl", args...).CombinedOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "launchctl %v: %v: %s\n", args, err, out)
	}
}

// clientPaths holds the signed and unsigned client helpers, set by run.
var clientPaths [2]string

func mustReadFile(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}

// entitlementsPlist is the ad-hoc signing entitlement the helper binaries
// carry. get-task-allow is a development entitlement an ad-hoc signature may
// embed (most entitlements are rejected at exec for ad-hoc identities), and
// it is readable by the XPC peer-requirement check.
func entitlementsPlist(entitlement string) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>` + entitlement + `</key><true/>
</dict>
</plist>
`
}

// agentPlist is the launchd job. The names must appear in MachServices or the
// service cannot create its listeners. launchd gives the job no terminal, so
// its output goes to a file the test prints on the way out.
func agentPlist(name, entName, tqName, program, logFile, covDir string) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key><string>` + name + `</string>
	<key>ProgramArguments</key><array><string>` + program + `</string></array>
	<key>EnvironmentVariables</key><dict>
		<key>XPCLIVE_SERVICE</key><string>` + name + `</string>
		<key>XPCLIVE_ENT_SERVICE</key><string>` + entName + `</string>
		<key>XPCLIVE_TQ_SERVICE</key><string>` + tqName + `</string>
		<key>XPCLIVE_ENTITLEMENT</key><string>` + testEntitlement + `</string>
		<key>GOCOVERDIR</key><string>` + covDir + `</string>
	</dict>
	<key>MachServices</key><dict>
		<key>` + name + `</key><true/>
		<key>` + entName + `</key><true/>
		<key>` + tqName + `</key><true/>
	</dict>
	<key>StandardOutPath</key><string>` + logFile + `</string>
	<key>StandardErrorPath</key><string>` + logFile + `</string>
</dict>
</plist>
`
}

// dial opens a session to the live service. The session is cancelled when the
// test ends.
func dial(t *testing.T, name string) *xpc.Session {
	t.Helper()
	s, err := xpc.DialMachService(name, xpc.SessionOptions{})
	if err != nil {
		t.Fatalf("dial %s: %v", name, err)
	}
	t.Cleanup(func() { s.Cancel() })
	return s
}

func TestLiveRoundTrip(t *testing.T) {
	session := dial(t, serviceName)

	const big = int64(1)<<53 + 1 // past float64's exact-integer range

	tests := []struct {
		name string
		send call
		want result
	}{
		{"add", call{Op: "add", First: 23, Second: 19}, result{Sum: 42}},
		{"negative", call{Op: "add", First: -5, Second: 2}, result{Sum: -3}},
		{"int64 past 2^53", call{Op: "echo", First: big}, result{Sum: big}},
		{"int64 min", call{Op: "echo", First: -1 << 63}, result{Sum: -1 << 63}},
		{"string field", call{Op: "echo", First: 7, Text: "héllo"}, result{Sum: 7, Echo: "héllo"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			msg, err := session.Call(context.Background(), test.send)
			if err != nil {
				t.Fatalf("Call(%+v): %v", test.send, err)
			}
			var got result
			if err := msg.Decode(&got); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if got != test.want {
				t.Fatalf("Call(%+v) = %+v, want %+v", test.send, got, test.want)
			}
		})
	}
}

// TestLiveWireKeys checks the dictionary that actually crossed the wire, not
// just the struct that came back, so a codec that agreed with itself on the
// wrong keys would still fail.
func TestLiveWireKeys(t *testing.T) {
	session := dial(t, serviceName)

	msg, err := session.CallDictionary(context.Background(), xpc.Dictionary{
		"op":           "add",
		"firstNumber":  int64(100),
		"secondNumber": int64(5),
	})
	if err != nil {
		t.Fatalf("CallDictionary: %v", err)
	}
	dict := msg.Dictionary()
	sum, ok := dict["sumValue"]
	if !ok {
		t.Fatalf("reply has no sumValue key: %v", dict)
	}
	if sum != int64(105) {
		t.Fatalf("sumValue = %v (%T), want int64(105)", sum, sum)
	}
}

// TestLiveErrorReply covers the handler-returns-error path, which replies with
// {"error": ...} rather than failing the send.
func TestLiveErrorReply(t *testing.T) {
	session := dial(t, serviceName)

	msg, err := session.Call(context.Background(), call{Op: "fail", First: 9})
	if err != nil {
		t.Fatalf("Call: %v", err)
	}
	got, _ := msg.Dictionary()["error"].(string)
	if !strings.Contains(got, `refused op "fail" with 9`) {
		t.Fatalf("error reply = %q, want the handler's message", got)
	}
}

// TestLiveNoReply covers the nil, nil handler result. The handler sends no
// reply, but the client does not simply wait forever: XPC releases the unused
// reply context and delivers "Underlying connection interrupted" to the
// pending reply handler. The session survives, so this is a dropped reply and
// not a dropped peer.
func TestLiveNoReply(t *testing.T) {
	session := dial(t, serviceName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	msg, err := session.Call(ctx, call{Op: "silent"})
	if err == nil {
		t.Fatalf("silent op produced a reply body %v, want no reply", msg.Dictionary())
	}
	// XPC actively reports the dropped reply context; it does not hang. If
	// that ever changes, our own deadline would fire and this test must fail
	// rather than pass on a timeout that proves nothing about the peer.
	if errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("silent op hung until our deadline (%v); the interrupted-connection evidence no longer holds", err)
	}
	if d := msg.Dictionary(); len(d) != 0 {
		t.Fatalf("failed reply carries a body %v, want empty", d)
	}
	t.Logf("silent op: %v", err)

	// The session does not survive. Releasing the unused reply context
	// interrupts the connection, so the client's session is cancelled and
	// every later send on it fails. A fresh session still reaches the same
	// service, so this costs the connection and not the peer.
	if _, err := session.Call(context.Background(), call{Op: "add", First: 1, Second: 2}); err == nil {
		t.Fatal("session still usable after an unanswered message; update this test and the docs")
	} else {
		t.Logf("send after silent op: %v", err)
	}

	fresh := dial(t, serviceName)
	msg, err = fresh.Call(context.Background(), call{Op: "add", First: 1, Second: 2})
	if err != nil {
		t.Fatalf("Call on a fresh session: %v", err)
	}
	var got result
	if err := msg.Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.Sum != 3 {
		t.Fatalf("Sum = %d, want 3", got.Sum)
	}
}

// TestLiveAsyncCall covers the asynchronous reply path
// (xpc_session_send_message_with_reply_async). Call only takes that path when
// ctx can end, so this test must pass a cancellable context: with
// context.Background it would silently retest the synchronous path instead.
func TestLiveAsyncCall(t *testing.T) {
	session := dial(t, serviceName)

	type outcome struct {
		got result
		err error
	}
	done := make(chan outcome, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func() {
		msg, err := session.Call(ctx, call{Op: "add", First: 2, Second: 40})
		if err != nil {
			done <- outcome{err: err}
			return
		}
		var got result
		done <- outcome{got: got, err: msg.Decode(&got)}
	}()
	select {
	case o := <-done:
		if o.err != nil {
			t.Fatalf("reply: %v", o.err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("no reply within 10s")
	}
}

// TestLiveInactiveActivate covers SessionOptions{Inactive: true}: the session
// is created dormant, configured, and only then activated. Activating an
// already-active session is API misuse that traps the process, so a regression
// in the flags plumbing shows up here as a dead test binary.
func TestLiveInactiveActivate(t *testing.T) {

	session, err := xpc.DialMachService(serviceName, xpc.SessionOptions{Inactive: true})
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(func() { session.Cancel() })

	if err := session.SetIncomingMessageHandler(func(xpc.ReceivedMessage) (any, error) {
		return nil, nil
	}); err != nil {
		t.Fatalf("SetIncomingMessageHandler: %v", err)
	}
	if err := session.SetCancellationHandler(func(xpc.RichError) {}); err != nil {
		t.Fatalf("SetCancellationHandler: %v", err)
	}
	if err := session.Activate(); err != nil {
		t.Fatalf("Activate: %v", err)
	}

	msg, err := session.Call(context.Background(), call{Op: "add", First: 20, Second: 22})
	if err != nil {
		t.Fatalf("Call: %v", err)
	}
	var got result
	if err := msg.Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.Sum != 42 {
		t.Fatalf("Sum = %d, want 42", got.Sum)
	}
}

// TestLivePeerGone has the service cancel the peer session after replying, and
// checks the client observes it rather than hanging.
func TestLivePeerGone(t *testing.T) {

	session, err := xpc.DialMachService(serviceName, xpc.SessionOptions{Inactive: true})
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(func() { session.Cancel() })

	cancelled := make(chan xpc.RichError, 1)
	if err := session.SetCancellationHandler(func(e xpc.RichError) {
		select {
		case cancelled <- e:
		default:
		}
	}); err != nil {
		t.Fatalf("SetCancellationHandler: %v", err)
	}
	if err := session.Activate(); err != nil {
		t.Fatalf("Activate: %v", err)
	}

	if _, err := session.Call(context.Background(), call{Op: "bye", First: 1}); err != nil {
		t.Fatalf("Call(bye): %v", err)
	}

	// The peer is gone. Sends must fail rather than block; retry briefly
	// because the cancellation is asynchronous.
	deadline := time.Now().Add(10 * time.Second)
	for {
		_, err := session.Call(context.Background(), call{Op: "add", First: 1, Second: 1})
		if err != nil {
			t.Logf("send after peer cancel: %v", err)
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("sends still succeed 10s after the peer cancelled")
		}
		time.Sleep(100 * time.Millisecond)
	}
	select {
	case e := <-cancelled:
		t.Logf("cancellation handler: %v", e)
	case <-time.After(5 * time.Second):
		t.Error("cancellation handler never ran")
	}
}

// TestLiveDialUnknownService covers the RichError path: a name nothing vends
// must fail with the runtime's own description, not a bare sentinel.
func TestLiveDialUnknownService(t *testing.T) {
	name := fmt.Sprintf("dev.tmc.apple.xpc.absent.%d", os.Getpid())

	session, err := xpc.DialMachService(name, xpc.SessionOptions{})
	if err == nil {
		t.Cleanup(func() { session.Cancel() })
		_, err = session.Call(context.Background(), call{Op: "add", First: 1, Second: 1})
		if err == nil {
			t.Fatalf("send to %s succeeded", name)
		}
	}
	var rich xpc.RichError
	if !errors.As(err, &rich) {
		t.Fatalf("error %v (%T) is not a RichError", err, err)
	}
	if rich.Error() == "" || rich.Error() == "xpc: unspecified error" {
		t.Fatalf("RichError carries no description: %q", rich.Error())
	}
	t.Logf("dial %s: %v", name, rich)
}

// testEntitlementRequirement builds a requirement naming the test entitlement
// that the signed helpers carry. The live suite is macOS-26-only: the
// requirement constructors are unavailable on earlier runtimes, which fails
// here rather than skipping, because the suite's whole point is proving peer
// enforcement.
func testEntitlementRequirement(t *testing.T) *xpc.PeerRequirement {
	t.Helper()
	req, err := xpc.NewEntitlementExistsRequirement(testEntitlement)
	if err != nil {
		t.Fatalf("NewEntitlementExistsRequirement(%q): %v", testEntitlement, err)
	}
	t.Cleanup(func() { req.Close() })
	return req
}

// absentRequirement names an entitlement nothing in the fixture carries.
func absentRequirement(t *testing.T) *xpc.PeerRequirement {
	t.Helper()
	req, err := xpc.NewEntitlementExistsRequirement("com.tmc.apple.xpc.live.absent")
	if err != nil {
		t.Fatalf("NewEntitlementExistsRequirement(absent): %v", err)
	}
	t.Cleanup(func() { req.Close() })
	return req
}

// dialWithRequirement dials the live service with a session-side requirement
// via SessionOptions, exercising the package's force-inactive install and
// auto-activate path. The returned session is active.
func dialWithRequirement(t *testing.T, req *xpc.PeerRequirement) *xpc.Session {
	t.Helper()
	s, err := xpc.DialMachService(serviceName, xpc.SessionOptions{Requirement: req})
	if err != nil {
		t.Fatalf("dial with requirement: %v", err)
	}
	t.Cleanup(func() { s.Cancel() })
	return s
}

// sendRoundTrip performs one add call and checks the reply, mirroring the
// wire keys of the other live tests.
func sendRoundTrip(t *testing.T, session *xpc.Session) {
	t.Helper()
	msg, err := session.Call(context.Background(), call{Op: "add", First: 21, Second: 21})
	if err != nil {
		t.Fatalf("Call: %v", err)
	}
	var got result
	if err := msg.Decode(&got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.Sum != 42 {
		t.Fatalf("Sum = %d, want 42", got.Sum)
	}
}

// sendMustFail retries briefly until the requirement enforcement lands: the
// rejection is delivered through the session, and the first send may race the
// cancellation.
func sendMustFail(t *testing.T, session *xpc.Session, want string) {
	t.Helper()
	deadline := time.Now().Add(10 * time.Second)
	for {
		_, err := session.Call(context.Background(), call{Op: "add", First: 1, Second: 1})
		if err != nil {
			if !strings.Contains(err.Error(), want) {
				t.Fatalf("send error %q does not mention %q", err, want)
			}
			t.Logf("requirement rejection observed: %v", err)
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("sends still succeed 10s after a rejecting requirement was installed")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// TestLiveSessionRequirementAdmits covers SessionOptions.Requirement naming
// the helper's entitlement: the peer satisfies it and the round trip works.
func TestLiveSessionRequirementAdmits(t *testing.T) {
	session := dialWithRequirement(t, testEntitlementRequirement(t))
	sendRoundTrip(t, session)
}

// TestLiveSessionRequirementRejects covers SessionOptions.Requirement naming
// an entitlement the helper lacks: every message is refused with the runtime's
// code-signing error.
func TestLiveSessionRequirementRejects(t *testing.T) {
	session := dialWithRequirement(t, absentRequirement(t))
	sendMustFail(t, session, "forbidden")
}

// TestLiveSetPeerRequirementAdmits covers the SetPeerRequirement method on an
// explicitly inactive session: install, activate, round trip.
func TestLiveSetPeerRequirementAdmits(t *testing.T) {
	session, err := xpc.DialMachService(serviceName, xpc.SessionOptions{Inactive: true})
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(func() { session.Cancel() })
	if err := session.SetPeerRequirement(testEntitlementRequirement(t)); err != nil {
		t.Fatalf("SetPeerRequirement: %v", err)
	}
	if err := session.Activate(); err != nil {
		t.Fatalf("Activate: %v", err)
	}
	sendRoundTrip(t, session)
}

// TestLiveSetPeerRequirementRejects covers SetPeerRequirement with an absent
// entitlement: the session is refused at send time.
func TestLiveSetPeerRequirementRejects(t *testing.T) {
	session, err := xpc.DialMachService(serviceName, xpc.SessionOptions{Inactive: true})
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(func() { session.Cancel() })
	if err := session.SetPeerRequirement(absentRequirement(t)); err != nil {
		t.Fatalf("SetPeerRequirement: %v", err)
	}
	if err := session.Activate(); err != nil {
		t.Fatalf("Activate: %v", err)
	}
	sendMustFail(t, session, "forbidden")
}

// runClient launches a client helper process and reports whether it did what
// the mode promises. For expect-reject the drop can manifest either as an
// error the client prints (exit 0) or as a hang; a hang is detected by the
// bounded deadline and the client's SENT marker, and is the named rejection
// result rather than a failure.
func runClient(t *testing.T, bin, service, mode string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, "-mode", mode, "-service", service)
	out, err := cmd.CombinedOutput()
	logged := string(out)
	switch mode {
	case "roundtrip":
		if err != nil {
			t.Fatalf("client %s: %v\n%s", service, err, logged)
		}
		if !strings.Contains(logged, "REPLIED 42") {
			t.Fatalf("client %s did not report a round trip:\n%s", service, logged)
		}
	case "expect-reject":
		if ctx.Err() == context.DeadlineExceeded {
			// Killed after the bounded window: pass only if the client got as
			// far as sending (the drop manifested as a hang).
			if !strings.Contains(logged, "SENT") {
				t.Fatalf("client %s was killed before sending:\n%s", service, logged)
			}
			t.Logf("client %s hung after SENT, treated as the named rejection", service)
			return
		}
		if err != nil {
			t.Fatalf("client %s exited %v:\n%s", service, err, logged)
		}
		if strings.Contains(logged, "REPLIED") {
			t.Fatalf("client %s got a reply that should have been dropped:\n%s", service, logged)
		}
		t.Logf("client %s observed the rejection: %s", service, strings.TrimSpace(logged))
	}
}

// TestLiveListenerRequirementAdmits drives a separately launched, signed
// client helper against the entitlement listener. This is the only test that
// can prove the listener enforces its requirement at all: a client in this
// test process would carry the test binary's signature, and listener code
// that ignored its requirement would pass every session-side test while every
// peer was admitted.
func TestLiveListenerRequirementAdmits(t *testing.T) {
	runClient(t, clientPaths[0], entServiceName, "roundtrip")
}

// TestLiveListenerRequirementRejects drives the unsigned client helper against
// the same listener: a peer without the entitlement must be dropped, and the
// client must not complete a round trip.
func TestLiveListenerRequirementRejects(t *testing.T) {
	runClient(t, clientPaths[1], entServiceName, "expect-reject")
}

// TestLiveListenerRequirementInactiveLifecycle is the explicit gate for the
// listener created with a requirement and without Inactive in the options: if
// the package created it active and then called the setter, the whole helper
// would have trapped at startup and every live test above would already have
// failed. This test just proves the helper is still alive and serving.
func TestLiveListenerRequirementInactiveLifecycle(t *testing.T) {
	session := dial(t, serviceName)
	sendRoundTrip(t, session)
}

// TestLiveTargetQueuePlacement drives a round trip through the listener whose
// accept handler asserts, inside the helper process, that it runs on the
// configured dispatch queue. If TargetQueue is ignored the assertion aborts
// the helper and this send fails instead of hanging.
func TestLiveTargetQueuePlacement(t *testing.T) {
	runClient(t, clientPaths[0], tqServiceName, "roundtrip")
}
