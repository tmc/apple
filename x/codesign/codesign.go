package codesign

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

// Options configures [EnsureSigned].
type Options struct {
	// Entitlements is the entitlements plist, as bytes. Callers typically embed
	// it with go:embed or inline a literal. Required.
	Entitlements []byte

	// RequireKeys lists the entitlement keys that must all be present for the
	// binary to count as already signed. If any is missing the binary is
	// (re)signed. At least one key is required.
	//
	// Checking every key — not just one — matters when a binary needs several
	// entitlements: a binary signed with a subset would otherwise be treated as
	// done and never gain the rest.
	RequireKeys []string

	// GuardEnv is the environment variable used to detect the post-sign re-exec
	// and prevent an infinite loop. Required; it must be unique to the binary
	// (e.g. "MYAPP_CODESIGN_DONE"). No default is supplied on purpose: a shared
	// default would let one binary's already-signed re-exec be misread by a
	// differently-built sibling during a rollout.
	GuardEnv string

	// VerifyHash, when true, hashes the executable before and after signing and
	// skips the re-exec if codesign left the file unchanged (it was already
	// signed in a way the entitlement check did not detect). When false the
	// binary always re-execs after a successful codesign.
	VerifyHash bool

	// runner runs the codesign command. nil means the real codesign binary; it
	// exists as a test seam so guard and plist handling can be exercised
	// without a signing toolchain.
	runner func(args ...string) ([]byte, error)
}

// EnsureSigned ad-hoc signs the current executable with opts.Entitlements and
// re-execs it once, unless the binary already carries every key in
// opts.RequireKeys or the guard variable opts.GuardEnv is set. On a successful
// sign it replaces the process image via execve, so it does not return on the
// signing path; it returns nil when no signing is needed.
func EnsureSigned(opts Options) error {
	if opts.GuardEnv == "" {
		return fmt.Errorf("codesign: GuardEnv is required")
	}
	if len(opts.RequireKeys) == 0 {
		return fmt.Errorf("codesign: RequireKeys must list at least one entitlement")
	}
	if len(opts.Entitlements) == 0 {
		return fmt.Errorf("codesign: Entitlements is required")
	}

	// A set guard means we are already inside the post-sign re-exec.
	if os.Getenv(opts.GuardEnv) == "1" {
		return nil
	}

	run := opts.runner
	if run == nil {
		run = func(args ...string) ([]byte, error) {
			cmd := exec.Command("codesign", args...)
			var stderr bytes.Buffer
			cmd.Stderr = &stderr
			out, err := cmd.Output()
			if err != nil {
				return out, fmt.Errorf("%s: %w", stderr.String(), err)
			}
			return out, nil
		}
	}

	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("codesign: get executable: %w", err)
	}
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		return fmt.Errorf("codesign: resolve symlink: %w", err)
	}

	if hasAllEntitlements(run, exe, opts.RequireKeys) {
		return nil
	}

	var hashBefore string
	if opts.VerifyHash {
		if hashBefore, err = hashFile(exe); err != nil {
			return fmt.Errorf("codesign: hash binary: %w", err)
		}
	}

	entFile, err := os.CreateTemp("", "entitlements-*.plist")
	if err != nil {
		return fmt.Errorf("codesign: create entitlements file: %w", err)
	}
	defer os.Remove(entFile.Name())
	if _, err := entFile.Write(opts.Entitlements); err != nil {
		entFile.Close()
		return fmt.Errorf("codesign: write entitlements: %w", err)
	}
	if err := entFile.Close(); err != nil {
		return fmt.Errorf("codesign: close entitlements: %w", err)
	}

	if _, err := run("-s", "-", "-f", "--entitlements", entFile.Name(), exe); err != nil {
		return fmt.Errorf("codesign: sign: %w", err)
	}

	if opts.VerifyHash {
		hashAfter, err := hashFile(exe)
		if err != nil {
			return fmt.Errorf("codesign: hash after signing: %w", err)
		}
		if hashBefore == hashAfter {
			// codesign left the binary unchanged; nothing to re-exec into.
			return nil
		}
	}

	env := append(os.Environ(), opts.GuardEnv+"=1")
	return syscall.Exec(exe, os.Args, env)
}

// hasAllEntitlements reports whether exe's entitlements contain every key.
func hasAllEntitlements(run func(args ...string) ([]byte, error), exe string, keys []string) bool {
	out, err := run("-d", "--entitlements", "-", exe)
	if err != nil {
		return false
	}
	for _, key := range keys {
		if !bytes.Contains(out, []byte(key)) {
			return false
		}
	}
	return true
}

func hashFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:]), nil
}
