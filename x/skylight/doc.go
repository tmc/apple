// Package skylight provides high-level, idiomatic Go bindings for macOS WindowServer (SkyLight) features.
//
// This package wraps the low-level private SkyLight SPIs in github.com/tmc/apple/private/skylight,
// providing typed and memory-safe abstractions for Space management, off-space window detection,
// focus switching without window raising, and process ownership queries.
//
// # Core Concepts
//
//   - Space: represents a macOS display space (virtual desktop).
//   - Window: represents a WindowServer window ID.
//   - ProcessSerialNumber: uniquely identifies a process in the WindowServer.
//
// # Usage Overview
//
// Space detection:
//
//	activeSpace, err := skylight.ActiveSpace()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	spaces, err := skylight.SpacesForWindow(windowID)
//
// Focusing without raising windows:
//
//	err := skylight.FocusWithoutRaise(targetPID)
//
// Querying window owners:
//
//	pid, err := skylight.WindowOwnerPID(windowID)
//
// # Thread Safety & Initialization
//
// WindowServer connection handles (CGSConnectionID) are lazily resolved upon first call using
// sync.Once. Package methods may be safely invoked concurrently from any goroutine.
package skylight
