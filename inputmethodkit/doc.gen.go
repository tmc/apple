// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

// Package inputmethodkit provides Go bindings for the InputMethodKit framework.
//
// Develop input methods and manage communication with client applications,
// candidates windows, and input method modes.
//
// The Input Method Kit, introduced in OS X v10.5, provides a streamlined
// programming interface that lets you develop input methods with far less
// code than older Mac programming interfaces. It is fully integrated with the
// Text Services Manager. The Input Method Kit allows 32-bit applications to
// work with 64-bit applications.
//
// # Classes
//
//   - [IMKCandidates]: The class presents candidates to users and notifies the appropriate object when the user selects a candidate. ([IMKCandidatePanelType], [IMKCandidatesLocationHint])
//   - [IMKInputController]: The class provides a base class for custom input controller classes.
//   - [IMKServer]: The class manages client connections to your input method.
//
// # Protocols
//
//   - [IMKMouseHandling]: The protocol defines methods that your input method can implement to handle mouse events.
//   - [IMKStateSetting]: The protocol defines methods for setting or accessing values that indicate the state of an input method.
//
// # Key Types
//
//   - [IMKCandidates] - The [IMKCandidates] class presents candidates to users and notifies the appropriate IMKInputController object when the user selects a candidate.
//   - [IMKInputController] - The [IMKInputController] class provides a base class for custom input controller classes.
//   - [IMKServer] - The [IMKServer] class manages client connections to your input method.
package inputmethodkit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the InputMethodKit library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/InputMethodKit.framework/InputMethodKit",
	"/usr/lib/libInputMethodKit.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: InputMethodKit: failed to load framework from any known path\n")
	}
}
