//go:build darwin

// Package anetest coordinates tests that use the system ANE compiler.
package anetest

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

// Run serializes m with ANE tests in other package test binaries.
func Run(m *testing.M) {
	f, err := lock()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ane test lock: %v\n", err)
		os.Exit(1)
	}

	code := m.Run()
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_UN); err != nil && code == 0 {
		fmt.Fprintf(os.Stderr, "ane test unlock: %v\n", err)
		code = 1
	}
	if err := f.Close(); err != nil && code == 0 {
		fmt.Fprintf(os.Stderr, "ane test lock close: %v\n", err)
		code = 1
	}
	os.Exit(code)
}

func lock() (*os.File, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("find home directory: %w", err)
	}
	dir := filepath.Join(home, "tmp")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return nil, fmt.Errorf("create lock directory: %w", err)
	}
	f, err := os.OpenFile(filepath.Join(dir, "tmc-apple-ane-tests.lock"), os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("open lock file: %w", err)
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		f.Close()
		return nil, fmt.Errorf("lock compiler: %w", err)
	}
	return f, nil
}
