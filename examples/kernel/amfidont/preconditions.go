package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/ebitengine/purego"
)

// csrAllowTaskForPID is CSR_ALLOW_TASK_FOR_PID from <sys/csr.h>. csr_check
// returns 0 when the operation is permitted on this boot and non-zero
// otherwise. task_for_pid on a platform binary needs this bit, which
// "csrutil enable --without debug" is what grants. These values are
// ABI-pinned in xnu and absent from Apple's documentation set, so they are
// written here rather than generated.
const csrAllowTaskForPID = 1 << 2

// checkPreconditions refuses to run unless every gate this tool depends on is
// satisfied, naming the specific gate that failed. It never infers a gate from
// a later failure: each is checked empirically here, before anything is
// attached.
func checkPreconditions() error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("must run as root (euid=%d): task_for_pid on a platform binary needs root", os.Geteuid())
	}

	// Verify Debugging Restrictions are actually off on this boot rather than
	// discovering it from a task_for_pid failure against a stopped daemon.
	var csrCheck func(uint32) int32
	if err := registerFunc(&csrCheck, "csr_check"); err != nil {
		return fmt.Errorf("resolve csr_check: %w", err)
	}
	if kr := csrCheck(csrAllowTaskForPID); kr != 0 {
		return fmt.Errorf("Debugging Restrictions are enabled (csr_check(CSR_ALLOW_TASK_FOR_PID)=%d): boot Recovery and run 'csrutil enable --without debug'", kr)
	}

	warnPrimaryInstall()
	return nil
}

// warnPrimaryInstall prints a loud warning to run on a spare install. A wedge
// of amfid costs a reboot, so this work does not belong on the volume you rely
// on. The boot volume cannot be identified as "the spare" programmatically —
// only its owner knows which install is disposable — so this warns rather than
// refuses, and names the volume so the operator can confirm.
func warnPrimaryInstall() {
	name := bootVolumeName()
	log.Printf("WARNING: running against amfid on boot volume %q.", name)
	log.Print("WARNING: run this ONLY on a spare install. A wedge here costs a reboot; do not do this on a volume you rely on.")
}

// bootVolumeName reports a human-readable name for the root volume,
// best-effort. macOS surfaces the boot volume under /Volumes as a symlink to
// "/", so a matching entry there gives the friendly name ("Macintosh HD");
// otherwise fall back to the backing device from statfs.
func bootVolumeName() string {
	rootDev := deviceOf("/")
	entries, err := os.ReadDir("/Volumes")
	if err == nil {
		for _, e := range entries {
			p := "/Volumes/" + e.Name()
			if target, err := os.Readlink(p); err == nil && target == "/" {
				return e.Name()
			}
			if deviceOf(p) == rootDev {
				return e.Name()
			}
		}
	}
	if rootDev != 0 {
		return fmt.Sprintf("device %d", rootDev)
	}
	return "/"
}

// deviceOf returns the device id backing path, or 0 if it cannot be read.
func deviceOf(path string) uint64 {
	fi, err := os.Stat(path)
	if err != nil {
		return 0
	}
	if st, ok := fi.Sys().(*syscall.Stat_t); ok {
		return uint64(st.Dev)
	}
	return 0
}

// registerFunc binds fptr to the named symbol from the global symbol table.
// It is the single choke point for symbol resolution so every bind happens up
// front, before the target is attached.
func registerFunc(fptr any, name string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("bind %s: %v", name, r)
		}
	}()
	purego.RegisterLibFunc(fptr, purego.RTLD_DEFAULT, name)
	return nil
}
