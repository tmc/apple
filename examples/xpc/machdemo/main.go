// Command machdemo is a runnable XPC round trip over a Mach service.
//
// The addservice and addclient examples mirror Apple's XPC service template,
// which only runs from inside an application bundle. This one uses a Mach
// service instead, so it can be exercised from a shell with nothing but a
// LaunchAgent:
//
//	go build -o ~/tmp/machdemo ./examples/xpc/machdemo
//	~/tmp/machdemo -install        # writes and loads the LaunchAgent
//	~/tmp/machdemo -call 23 19     # prints 42
//	~/tmp/machdemo -uninstall
//
// launchd starts the -serve side on demand when the first message arrives.
// The service name must appear in the agent's MachServices dictionary or the
// listener cannot be created.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/xpc"
)

const serviceName = "dev.tmc.apple.xpc.machdemo"

type request struct {
	First  int64 `xpc:"firstNumber"`
	Second int64 `xpc:"secondNumber"`
}

type reply struct {
	Result int64 `xpc:"result"`
}

func main() {
	serve := flag.Bool("serve", false, "run the service (launchd starts this)")
	call := flag.Bool("call", false, "call the service with two addends")
	install := flag.Bool("install", false, "write and load the LaunchAgent")
	uninstall := flag.Bool("uninstall", false, "unload and remove the LaunchAgent")
	peerEntitlement := flag.String("require-peer-entitlement", "", "require this entitlement on the remote process")
	targetQueue := flag.Bool("target-queue", false, "deliver session or listener events on a private dispatch queue")
	flag.Parse()

	switch {
	case *serve:
		runService(*peerEntitlement, *targetQueue)
	case *install:
		must(installAgent(*peerEntitlement, *targetQueue))
	case *uninstall:
		must(uninstallAgent())
	case *call:
		must(callService(flag.Args(), *peerEntitlement, *targetQueue))
	default:
		flag.Usage()
		os.Exit(2)
	}
}

// runService is the launchd side. ForceMach is required: without it the
// runtime does not treat the name as a Mach service and creation fails.
//
// Inactive matters just as much. A listener created without it is already
// active, and calling Activate on an active listener is API misuse: XPC traps
// the process rather than returning an error.
func runService(peerEntitlement string, useTargetQueue bool) {
	opts := xpc.ListenerOptions{ForceMach: true, Inactive: true}
	if useTargetQueue {
		opts.TargetQueue = dispatch.QueueCreate("xpc.machdemo.listener")
	}
	if peerEntitlement != "" {
		req, err := xpc.NewEntitlementExistsRequirement(peerEntitlement)
		if err != nil {
			log.Fatalf("create peer requirement: %v", err)
		}
		defer req.Close()
		opts.Requirement = req
	}
	listener, err := xpc.NewServiceListener(serviceName, opts, accept)
	if err != nil {
		log.Fatalf("create listener: %v", err)
	}
	if err := listener.Activate(); err != nil {
		log.Fatalf("activate listener: %v", err)
	}
	log.Printf("serving %s", serviceName)
	select {}
}

func accept(req xpc.IncomingSessionRequest) xpc.IncomingDecision {
	return req.Accept(handle, func(err xpc.RichError) {
		log.Printf("session cancelled: %v", err)
	})
}

func handle(msg xpc.ReceivedMessage) (any, error) {
	var req request
	if err := msg.Decode(&req); err != nil {
		return nil, err
	}
	return reply{Result: req.First + req.Second}, nil
}

func callService(args []string, peerEntitlement string, useTargetQueue bool) error {
	if len(args) != 2 {
		return fmt.Errorf("need two addends, got %d", len(args))
	}
	first, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("first addend: %w", err)
	}
	second, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("second addend: %w", err)
	}

	// Lifecycle-sensitive configuration is deliberately shown through the
	// setters. Both setters require an inactive session; Activate makes it
	// usable after configuration.
	configured := peerEntitlement != "" || useTargetQueue
	session, err := xpc.DialMachService(serviceName, xpc.SessionOptions{Inactive: configured})
	if err != nil {
		return fmt.Errorf("dial %s: %w", serviceName, err)
	}
	defer session.Cancel()
	if useTargetQueue {
		queue := dispatch.QueueCreate("xpc.machdemo.session")
		if err := session.SetTargetQueue(queue); err != nil {
			return fmt.Errorf("set target queue: %w", err)
		}
	}
	if peerEntitlement != "" {
		req, err := xpc.NewEntitlementExistsRequirement(peerEntitlement)
		if err != nil {
			return fmt.Errorf("create peer requirement: %w", err)
		}
		defer req.Close()
		if err := session.SetPeerRequirement(req); err != nil {
			return fmt.Errorf("set peer requirement: %w", err)
		}
	}
	if configured {
		if err := session.Activate(); err != nil {
			return fmt.Errorf("activate session: %w", err)
		}
	}

	got, err := session.Call(context.Background(), request{First: first, Second: second})
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	var out reply
	if err := got.Decode(&out); err != nil {
		return fmt.Errorf("decode reply: %w", err)
	}
	fmt.Println(out.Result)
	return nil
}

var agentPlist = template.Must(template.New("plist").Parse(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>{{.Label}}</string>
	<key>ProgramArguments</key>
	<array>
		<string>{{.Program}}</string>
		{{range .Args}}<string>{{.}}</string>
		{{end}}
	</array>
	<key>MachServices</key>
	<dict>
		<key>{{.Label}}</key>
		<true/>
	</dict>
	<key>StandardOutPath</key>
	<string>{{.Log}}</string>
	<key>StandardErrorPath</key>
	<string>{{.Log}}</string>
</dict>
</plist>
`))

// logPath is where launchd sends the service's stdout and stderr. A service
// started by launchd has no terminal, so without this its errors are invisible.
func logPath() string {
	return filepath.Join(os.TempDir(), serviceName+".log")
}

func agentPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library", "LaunchAgents", serviceName+".plist"), nil
}

func installAgent(peerEntitlement string, useTargetQueue bool) error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	path, err := agentPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	args := []string{"-serve"}
	if peerEntitlement != "" {
		args = append(args, "-require-peer-entitlement", peerEntitlement)
	}
	if useTargetQueue {
		args = append(args, "-target-queue")
	}
	err = agentPlist.Execute(f, struct {
		Label, Program, Log string
		Args                []string
	}{serviceName, self, logPath(), args})
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return err
	}
	if err := launchctl("bootstrap", domain(), path); err != nil {
		return err
	}
	fmt.Printf("loaded %s\n", path)
	return nil
}

func uninstallAgent() error {
	path, err := agentPath()
	if err != nil {
		return err
	}
	// A missing job is not an error here; the plist may have been loaded by
	// hand or already unloaded.
	_ = launchctl("bootout", domain()+"/"+serviceName)
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	fmt.Printf("removed %s\n", path)
	return nil
}

func domain() string {
	return fmt.Sprintf("gui/%d", os.Getuid())
}

func launchctl(args ...string) error {
	cmd := exec.Command("launchctl", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("launchctl %v: %v: %s", args, err, out)
	}
	return nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
