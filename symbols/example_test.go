//go:build darwin

package symbols_test

import (
	"fmt"

	"github.com/tmc/apple/symbols"
)

func ExampleNSSymbolEffectOptions() {
	options := symbols.GetNSSymbolEffectOptionsClass().Options()
	nonRepeating := options.OptionsWithNonRepeating()
	speed := options.OptionsWithSpeedWithSpeed(2.0)

	fmt.Printf("Options class consistent: %t\n", symbols.GetNSSymbolEffectOptionsClass().Class() == symbols.GetNSSymbolEffectOptionsClass().Class())
	fmt.Printf("Configured options non-nil: %t\n", nonRepeating != nil && speed != nil)

	// Output:
	// Options class consistent: true
	// Configured options non-nil: true
}

func ExampleNSSymbolBounceEffect() {
	bounce := symbols.GetNSSymbolBounceEffectClass().Effect()
	byLayer := bounce.EffectWithByLayer()
	whole := bounce.EffectWithWholeSymbol()

	fmt.Printf("Bounce effect class consistent: %t\n", symbols.GetNSSymbolBounceEffectClass().Class() == symbols.GetNSSymbolBounceEffectClass().Class())
	fmt.Printf("Effect variants non-nil: %t\n", byLayer != nil && whole != nil)

	// Output:
	// Bounce effect class consistent: true
	// Effect variants non-nil: true
}
