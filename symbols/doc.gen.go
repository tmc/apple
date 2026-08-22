// Code generated from Apple documentation for Symbols. DO NOT EDIT.

// Package symbols provides Go bindings for the Symbols framework.
//
// Apply universal animations to symbol-based images.
//
// The Symbols framework provides access to symbol effects you can use to
// animate SF Symbols in your AppKit, UIKit, and SwiftUI apps. These
// animations exhibit different behaviors:
//
// # Symbol effects
//
//   - [NSSymbolAppearEffect]: A type that makes the layers of a symbol-based image appear separately or as a whole.
//   - [NSSymbolBounceEffect]: A type that applies a transitory scaling effect, or bounce, to the layers in a symbol-based image separately or as a whole.
//   - [NSSymbolDisappearEffect]: A type that makes the layers of a symbol-based image disappear separately or as a whole.
//   - [NSSymbolPulseEffect]: A type that fades the opacity of some or all layers in a symbol-based image.
//   - [NSSymbolScaleEffect]: A type that scales the layers in a symbol-based image separately or as a whole.
//   - [NSSymbolVariableColorEffect]: A type that replaces the opacity of variable layers in a symbol-based image in a repeatable sequence.
//   - [NSSymbolBreatheEffect]: A symbol effect that applies the Breathe animation to symbol images.
//   - [NSSymbolRotateEffect]: A symbol effect that applies the Rotate animation to symbol images.
//   - [NSSymbolWiggleEffect]: A symbol effect that applies the Wiggle animation to symbol images.
//
// # Symbol content transitions
//
//   - [NSSymbolReplaceContentTransition]: A type that replaces the layers of one symbol-based image with those of another.
//   - [NSSymbolAutomaticContentTransition]: A type that applies the default animation to a symbol-based image in a context-sensitive manner.
//   - [NSSymbolMagicReplaceContentTransition]: A symbol effect applies the MagicReplace animation to symbol images.
//
// # Symbol effect options
//
//   - [NSSymbolEffectOptions]: Options that configure how effects apply to symbol-based images.
//   - [NSSymbolEffectOptionsRepeatBehavior]: The behavior of repetition to use when a symbol effect is animating.
//
// # Symbol effect classes
//
//   - [NSSymbolEffect]: An abstract base class for effects that you can apply to a symbol-based image.
//   - [NSSymbolContentTransition]: An abstract base class for transitions you can apply to symbol-based images.
//
// # Classes
//
//   - [NSSymbolDrawOffEffect]: A symbol effect that applies the DrawOff animation to symbol images.
//   - [NSSymbolDrawOnEffect]: A symbol effect that applies the DrawOn animation to symbol images.
//
// # Key Types
//
//   - [NSSymbolWiggleEffect] - A symbol effect that applies the Wiggle animation to symbol images.
//   - [NSSymbolEffectOptions] - Options that configure how effects apply to symbol-based images.
//   - [NSSymbolReplaceContentTransition] - A type that replaces the layers of one symbol-based image with those of another.
//   - [NSSymbolVariableColorEffect] - A type that replaces the opacity of variable layers in a symbol-based image in a repeatable sequence.
//   - [NSSymbolDrawOffEffect] - A symbol effect that applies the DrawOff animation to symbol images.
//   - [NSSymbolAppearEffect] - A type that makes the layers of a symbol-based image appear separately or as a whole.
//   - [NSSymbolBounceEffect] - A type that applies a transitory scaling effect, or bounce, to the layers in a symbol-based image separately or as a whole.
//   - [NSSymbolBreatheEffect] - A symbol effect that applies the Breathe animation to symbol images.
//   - [NSSymbolDisappearEffect] - A type that makes the layers of a symbol-based image disappear separately or as a whole.
//   - [NSSymbolEffectOptionsRepeatBehavior] - The behavior of repetition to use when a symbol effect is animating.
package symbols

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Symbols library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Symbols.framework/Symbols",
	"/usr/lib/libSymbols.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: Symbols: failed to load framework from any known path\n")
	}
}
