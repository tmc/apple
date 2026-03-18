// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSymbolEffectOptions] class.
var (
	_NSSymbolEffectOptionsClass     NSSymbolEffectOptionsClass
	_NSSymbolEffectOptionsClassOnce sync.Once
)

func getNSSymbolEffectOptionsClass() NSSymbolEffectOptionsClass {
	_NSSymbolEffectOptionsClassOnce.Do(func() {
		_NSSymbolEffectOptionsClass = NSSymbolEffectOptionsClass{class: objc.GetClass("NSSymbolEffectOptions")}
	})
	return _NSSymbolEffectOptionsClass
}

// GetNSSymbolEffectOptionsClass returns the class object for NSSymbolEffectOptions.
func GetNSSymbolEffectOptionsClass() NSSymbolEffectOptionsClass {
	return getNSSymbolEffectOptionsClass()
}

type NSSymbolEffectOptionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolEffectOptionsClass) Alloc() NSSymbolEffectOptions {
	rv := objc.Send[NSSymbolEffectOptions](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Options that configure how effects apply to symbol-based images.
//
// # Configuring repeating effects
//
//   - [NSSymbolEffectOptions.OptionsWithNonRepeating]: A set of effect options that prefers to not repeat.
//
// # Configuring effect speed
//
//   - [NSSymbolEffectOptions.OptionsWithSpeedWithSpeed]: Creates a set of effect options with a preferred speed multiplier.
//
// # Instance Methods
//
//   - [NSSymbolEffectOptions.OptionsWithRepeatBehaviorWithBehavior]: Return a copy of the options setting a preferred repeat behavior.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptions
type NSSymbolEffectOptions struct {
	objectivec.Object
}

// NSSymbolEffectOptionsFromID constructs a [NSSymbolEffectOptions] from an objc.ID.
//
// Options that configure how effects apply to symbol-based images.
func NSSymbolEffectOptionsFromID(id objc.ID) NSSymbolEffectOptions {
	return NSSymbolEffectOptions{objectivec.Object{ID: id}}
}
// NOTE: NSSymbolEffectOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSymbolEffectOptions] class.
//
// # Configuring repeating effects
//
//   - [INSSymbolEffectOptions.OptionsWithNonRepeating]: A set of effect options that prefers to not repeat.
//
// # Configuring effect speed
//
//   - [INSSymbolEffectOptions.OptionsWithSpeedWithSpeed]: Creates a set of effect options with a preferred speed multiplier.
//
// # Instance Methods
//
//   - [INSSymbolEffectOptions.OptionsWithRepeatBehaviorWithBehavior]: Return a copy of the options setting a preferred repeat behavior.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptions
type INSSymbolEffectOptions interface {
	objectivec.IObject

	// Topic: Configuring repeating effects

	// A set of effect options that prefers to not repeat.
	OptionsWithNonRepeating() INSSymbolEffectOptions

	// Topic: Configuring effect speed

	// Creates a set of effect options with a preferred speed multiplier.
	OptionsWithSpeedWithSpeed(speed float64) INSSymbolEffectOptions

	// Topic: Instance Methods

	// Return a copy of the options setting a preferred repeat behavior.
	OptionsWithRepeatBehaviorWithBehavior(behavior INSSymbolEffectOptionsRepeatBehavior) INSSymbolEffectOptions

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSymbolEffectOptions) Init() NSSymbolEffectOptions {
	rv := objc.Send[NSSymbolEffectOptions](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolEffectOptions) Autorelease() NSSymbolEffectOptions {
	rv := objc.Send[NSSymbolEffectOptions](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolEffectOptions creates a new NSSymbolEffectOptions instance.
func NewNSSymbolEffectOptions() NSSymbolEffectOptions {
	class := getNSSymbolEffectOptionsClass()
	rv := objc.Send[NSSymbolEffectOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A set of effect options that prefers to not repeat.
//
// # Return Value
// 
// A copy of the effect options that prefers to not repeat.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptions/optionsWithNonRepeating-c.method
func (s NSSymbolEffectOptions) OptionsWithNonRepeating() INSSymbolEffectOptions {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("optionsWithNonRepeating"))
	return NSSymbolEffectOptionsFromID(rv)
}

// Creates a set of effect options with a preferred speed multiplier.
//
// speed: The preferred speed multiplier to play the effect with. The default
// multiplier is `1.0`. The function may clamp very large or small values.
//
// # Return Value
// 
// A copy of the effect options with the preferred speed multiplier.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptions/optionsWithSpeed:-c.method
func (s NSSymbolEffectOptions) OptionsWithSpeedWithSpeed(speed float64) INSSymbolEffectOptions {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("optionsWithSpeed:"), speed)
	return NSSymbolEffectOptionsFromID(rv)
}

// Return a copy of the options setting a preferred repeat behavior.
//
// behavior: The preferred behavior when the effect is repeated.
//
// # Return Value
// 
// A new options object with the preferred repeat behavior.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptions/optionsWithRepeatBehavior:-c.method
func (s NSSymbolEffectOptions) OptionsWithRepeatBehaviorWithBehavior(behavior INSSymbolEffectOptionsRepeatBehavior) INSSymbolEffectOptions {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("optionsWithRepeatBehavior:"), behavior)
	return NSSymbolEffectOptionsFromID(rv)
}
func (s NSSymbolEffectOptions) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The default set of effect options.
//
// # Return Value
// 
// A new instance of effect options.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptions/options
func (_NSSymbolEffectOptionsClass NSSymbolEffectOptionsClass) Options() NSSymbolEffectOptions {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolEffectOptionsClass.class), objc.Sel("options"))
	return NSSymbolEffectOptionsFromID(rv)
}

