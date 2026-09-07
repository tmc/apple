// Command amfidont relaxes one nominated code-signing verdict inside the
// running amfid daemon, for one binary, once, and then exits.
//
// It attaches to /usr/libexec/amfid with x/machdebug, breaks at the return of
// -[AMFIPathValidator_macos validateWithError:], and — only when the validated
// binary's signing identifier matches the -identity allowlist — overwrites the
// validator's result ivars so the verdict reads Apple-signed and unrestricted.
// It patches exactly that one verdict, detaches, and exits. There is no
// resident daemon, no launchd job, and nothing is written to disk.
//
// This is a research harness. amfid honors com.apple.private.virtualization
// only on binaries whose signature chains to Apple; that gate is what this tool
// relaxes so an experiment that needs a restricted VM device can run under a
// locally-signed binary.
//
// It functions only on a host whose owner has already disabled Debugging
// Restrictions from Recovery ("csrutil enable --without debug"). It cannot be
// used to weaken a machine that has not already been deliberately weakened by
// its owner in a physical-presence flow. amfid is on the critical path for
// every process launch on the system: a crash while a thread is parked at a
// breakpoint, or an exit with a BRK still written, wedges every new launch.
// The one-shot shape and the mandatory watchdog bound that window.
//
// Usage:
//
//	sudo amfidont -identity <signing-identifier> -- ./somebinary [args...]
//
// If anything goes wrong, recover amfid with:
//
//	sudo launchctl kickstart -k system/com.apple.amfid
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/tmc/apple/x/machdebug"
)

// recoveryCommand restores amfid if this tool leaves it wedged. It is printed
// on every error path.
const recoveryCommand = "sudo launchctl kickstart -k system/com.apple.amfid"

// watchdogInterval bounds how long amfid may stay attached waiting for the
// nominated identity. If no matching validation arrives, the watchdog detaches
// and the tool exits. It is mandatory, not optional: amfid must never stay
// under our breakpoints indefinitely.
const watchdogInterval = 60 * time.Second

func main() {
	log.SetFlags(0)
	log.SetPrefix("amfidont: ")
	if err := run(); err != nil {
		log.Print(err)
		fmt.Fprintf(os.Stderr, "amfidont: if amfid is wedged, recover with:\n\t%s\n", recoveryCommand)
		os.Exit(1)
	}
}

func run() error {
	identity := flag.String("identity", "", "signing identifier of the one binary whose verdict to relax (required)")
	flag.Usage = usage
	flag.Parse()
	if *identity == "" {
		usage()
		return errors.New("missing required -identity")
	}
	child := flag.Args()
	if len(child) == 0 {
		usage()
		return errors.New("missing command to run after --")
	}

	if err := checkPreconditions(); err != nil {
		return err
	}

	// Resolve and pre-fault everything before attaching. A debugger that must
	// dlopen a library or bind a symbol while its target is stopped can
	// deadlock against that target — and amfid is exactly the target that
	// would have to validate any dylib we tried to load.
	av, err := resolveAMFI()
	if err != nil {
		return err
	}
	log.Print(av.summary())

	if s := av.bootargState(); s.resolved {
		log.Printf("preflight: amfi_get_out_of_my_way=%v (bootarg state %#x)", s.getOutOfMyWay, s.raw)
	}

	pid, err := findAmfid()
	if err != nil {
		return err
	}
	log.Printf("target: amfid pid %d", pid)

	// The validator IMPs live in the dyld shared cache, whose slide is chosen
	// per boot. Our resolved IMP is slid into our address space; it is only a
	// valid breakpoint address in amfid if amfid carries the same slide. Verify
	// that before writing a BRK into a system daemon.
	ourSlide, err := selfSharedCacheSlide()
	if err != nil {
		return fmt.Errorf("read our shared-cache slide: %w", err)
	}
	amfidSlide, err := taskSharedCacheSlide(pid)
	if err != nil {
		return fmt.Errorf("read amfid shared-cache slide: %w", err)
	}
	targetIMP, err := translateIMP(av.imp, ourSlide, amfidSlide)
	if err != nil {
		return err
	}
	av.imp = targetIMP
	log.Printf("shared-cache slide %#x matches; validateWithError: IMP at %#x", ourSlide, av.imp)

	p, err := machdebug.Attach(pid)
	if err != nil {
		return fmt.Errorf("attach amfid: %w", err)
	}
	// Detach must run on the normal path, on error, and on signal: a process
	// that exits with a BRK still written and no handler on the port hangs the
	// target forever.
	var detached atomic.Bool
	detach := func() {
		if detached.CompareAndSwap(false, true) {
			if err := p.Detach(); err != nil {
				log.Printf("detach: %v", err)
			}
		}
	}
	defer detach()

	var timedOut atomic.Bool
	watchdog := time.AfterFunc(watchdogInterval, func() {
		timedOut.Store(true)
		log.Printf("watchdog: no matching validation within %s; detaching", watchdogInterval)
		detach()
	})
	defer watchdog.Stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		log.Print("signal: detaching")
		detach()
	}()

	// Start the nominated binary. Its launch is what makes amfid validate it,
	// which is the verdict we are here to catch.
	cmd := exec.Command(child[0], child[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start %q: %w", child[0], err)
	}
	log.Printf("launched %q pid %d", child[0], cmd.Process.Pid)

	patched, err := hookLoop(p, av, *identity)
	if err != nil {
		if timedOut.Load() {
			return fmt.Errorf("watchdog fired before identity %q was validated", *identity)
		}
		return err
	}
	if !patched {
		return fmt.Errorf("detached without patching identity %q", *identity)
	}
	log.Printf("patched verdict for %q; detaching. %q keeps running (pid %d)", *identity, child[0], cmd.Process.Pid)
	detach()
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `usage: amfidont -identity <signing-identifier> -- command [args...]

amfidont relaxes one code-signing verdict inside the running amfid daemon, for
the one binary whose signing identifier equals -identity, once, then exits. It
launches command, waits for amfid to validate a binary with that identifier,
overwrites that single verdict to read Apple-signed and unrestricted, detaches,
and exits. Nothing is installed and nothing on disk is modified.

It works ONLY on a host whose owner has already disabled Debugging Restrictions
from Recovery (csrutil enable --without debug) and is running as root. Run it on
a spare install, never on your primary volume: a wedge here costs a reboot.

If amfid is left wedged, recover with:
	%s

`, recoveryCommand)
	flag.PrintDefaults()
}
