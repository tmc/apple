
// Code generated from Apple documentation for Symbols. DO NOT EDIT.

// Package symbols provides Go bindings for the Symbols framework.
//
// Apply universal animations to symbol-based images.
//
// The Symbols framework provides access to symbol effects you can use to animate [SF Symbols](<https://developer.apple.com/design/human-interface-guidelines/sf-symbols>) in your AppKit, UIKit, and SwiftUI apps. These animations exhibit different behaviors:
//
// # Symbol effects
//
//   - appear: An animation that makes the layers of a symbol-based image appear separately or as a whole.
//   - bounce: An animation that applies a transitory scaling effect, or bounce, to the layers in a symbol-based image separately or as a whole.
//   - disappear: An animation that makes the layers of a symbol-based image disappear separately or as a whole.
//   - pulse: An animation that fades the opacity of some or all layers in a symbol-based image.
//   - scale: An animation that scales the layers in a symbol-based image separately or as a whole.
//   - variableColor: An animation that replaces the opacity of variable layers in a symbol-based image in a repeatable sequence.
//
// # Symbol content transitions
//
//   - replace: An animation that replaces the layers of one symbol-based image with those of another.
//   - automatic: A transition that applies the default animation to a symbol-based image in a context-sensitive manner.
//
// # Symbol effect types
//
//   - AppearSymbolEffect: A type that makes the layers of a symbol-based image appear separately or as a whole.
//   - AutomaticSymbolEffect: A type that applies the default animation to a symbol-based image in a context-sensitive manner.
//   - BounceSymbolEffect: A type that applies a transitory scaling effect, or bounce, to the layers in a symbol-based image separately or as a whole.
//   - DisappearSymbolEffect: A type that makes the layers of a symbol-based image disappear separately or as a whole.
//   - PulseSymbolEffect: A type that fades the opacity of some or all layers in a symbol-based image.
//   - ReplaceSymbolEffect: A type that replaces the layers of one symbol-based image with those of another.
//   - ScaleSymbolEffect: A type that scales the layers in a symbol-based image separately or as a whole.
//   - VariableColorSymbolEffect: A type that replaces the opacity of variable layers in a symbol-based image in a repeatable sequence.
//   - BreatheSymbolEffect: A symbol effect that applies the Breathe animation to symbol images.
//   - RotateSymbolEffect: A symbol effect that applies the Rotate animation to symbol images.
//   - WiggleSymbolEffect: A symbol effect that applies the Wiggle animation to symbol images.
//
// # Symbol effect options
//
//   - SymbolEffectOptions: Options that configure how effects apply to symbol-based images.
//
// # Symbol effect protocols
//
//   - SymbolEffect: A presentation effect that you apply to a symbol-based image.
//   - DiscreteSymbolEffect: An effect that performs a transient animation.
//   - IndefiniteSymbolEffect: An animation that continually affects a symbol until it’s disabled or removed.
//   - ContentTransitionSymbolEffect: An effect that animates between symbols or different configurations of the same symbol.
//   - TransitionSymbolEffect: An effect that animates a symbol in or out.
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
//
// [Symbols Documentation]: https://developer.apple.com/documentation/Symbols
package symbols

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Symbols.framework/Symbols"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Symbols: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

