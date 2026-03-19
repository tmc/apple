// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSUserInterfaceCompressionOptions] class.
var (
	_NSUserInterfaceCompressionOptionsClass     NSUserInterfaceCompressionOptionsClass
	_NSUserInterfaceCompressionOptionsClassOnce sync.Once
)

func getNSUserInterfaceCompressionOptionsClass() NSUserInterfaceCompressionOptionsClass {
	_NSUserInterfaceCompressionOptionsClassOnce.Do(func() {
		_NSUserInterfaceCompressionOptionsClass = NSUserInterfaceCompressionOptionsClass{class: objc.GetClass("NSUserInterfaceCompressionOptions")}
	})
	return _NSUserInterfaceCompressionOptionsClass
}

// GetNSUserInterfaceCompressionOptionsClass returns the class object for NSUserInterfaceCompressionOptions.
func GetNSUserInterfaceCompressionOptionsClass() NSUserInterfaceCompressionOptionsClass {
	return getNSUserInterfaceCompressionOptionsClass()
}

type NSUserInterfaceCompressionOptionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUserInterfaceCompressionOptionsClass) Alloc() NSUserInterfaceCompressionOptions {
	rv := objc.Send[NSUserInterfaceCompressionOptions](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies how user interface elements resize themselves when
// space is constrained.
//
// # Overview
// 
// An instance of [NSUserInterfaceCompressionOptions] contains zero or more
// options. Because a compression options object behaves like a set, you can
// use common operations like intersection, union and subtraction to interact
// with instances and their members.
// 
// You can access system-defined options through the class methods detailed in
// Creating standard options, or you can create your own custom options with
// the [NSUserInterfaceCompressionOptions.InitWithIdentifier] initializer.
// 
// To compare two different compression options objects, use the methods
// described in the Comparing compression options section.
//
// # Creating a compression option
//
//   - [NSUserInterfaceCompressionOptions.InitWithCompressionOptions]: Creates an option object that represents the union of the supplied options.
//   - [NSUserInterfaceCompressionOptions.InitWithIdentifier]: Creates an option object with the given identifier string.
//   - [NSUserInterfaceCompressionOptions.InitWithCoder]: Creates an option object from data in an unarchiver.
//
// # Comparing compression options
//
//   - [NSUserInterfaceCompressionOptions.Empty]: A Boolean value that denotes whether the option is empty.
//   - [NSUserInterfaceCompressionOptions.ContainsOptions]: Determines whether the supplied compression options are all present in the current instance.
//   - [NSUserInterfaceCompressionOptions.IntersectsOptions]: Determines whether the supplied compression options intersect with the current instance’s options.
//
// # Combining compression options
//
//   - [NSUserInterfaceCompressionOptions.OptionsByAddingOptions]: Creates a new compression options object representing the union with the provided options.
//   - [NSUserInterfaceCompressionOptions.OptionsByRemovingOptions]: Creates a new compression options object with the supplied options removed.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions
type NSUserInterfaceCompressionOptions struct {
	objectivec.Object
}

// NSUserInterfaceCompressionOptionsFromID constructs a [NSUserInterfaceCompressionOptions] from an objc.ID.
//
// An object that specifies how user interface elements resize themselves when
// space is constrained.
func NSUserInterfaceCompressionOptionsFromID(id objc.ID) NSUserInterfaceCompressionOptions {
	return NSUserInterfaceCompressionOptions{objectivec.Object{ID: id}}
}
// NOTE: NSUserInterfaceCompressionOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSUserInterfaceCompressionOptions] class.
//
// # Creating a compression option
//
//   - [INSUserInterfaceCompressionOptions.InitWithCompressionOptions]: Creates an option object that represents the union of the supplied options.
//   - [INSUserInterfaceCompressionOptions.InitWithIdentifier]: Creates an option object with the given identifier string.
//   - [INSUserInterfaceCompressionOptions.InitWithCoder]: Creates an option object from data in an unarchiver.
//
// # Comparing compression options
//
//   - [INSUserInterfaceCompressionOptions.Empty]: A Boolean value that denotes whether the option is empty.
//   - [INSUserInterfaceCompressionOptions.ContainsOptions]: Determines whether the supplied compression options are all present in the current instance.
//   - [INSUserInterfaceCompressionOptions.IntersectsOptions]: Determines whether the supplied compression options intersect with the current instance’s options.
//
// # Combining compression options
//
//   - [INSUserInterfaceCompressionOptions.OptionsByAddingOptions]: Creates a new compression options object representing the union with the provided options.
//   - [INSUserInterfaceCompressionOptions.OptionsByRemovingOptions]: Creates a new compression options object with the supplied options removed.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions
type INSUserInterfaceCompressionOptions interface {
	objectivec.IObject

	// Topic: Creating a compression option

	// Creates an option object that represents the union of the supplied options.
	InitWithCompressionOptions(options foundation.INSSet) NSUserInterfaceCompressionOptions
	// Creates an option object with the given identifier string.
	InitWithIdentifier(identifier string) NSUserInterfaceCompressionOptions
	// Creates an option object from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSUserInterfaceCompressionOptions

	// Topic: Comparing compression options

	// A Boolean value that denotes whether the option is empty.
	Empty() bool
	// Determines whether the supplied compression options are all present in the current instance.
	ContainsOptions(options INSUserInterfaceCompressionOptions) bool
	// Determines whether the supplied compression options intersect with the current instance’s options.
	IntersectsOptions(options INSUserInterfaceCompressionOptions) bool

	// Topic: Combining compression options

	// Creates a new compression options object representing the union with the provided options.
	OptionsByAddingOptions(options INSUserInterfaceCompressionOptions) INSUserInterfaceCompressionOptions
	// Creates a new compression options object with the supplied options removed.
	OptionsByRemovingOptions(options INSUserInterfaceCompressionOptions) INSUserInterfaceCompressionOptions

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u NSUserInterfaceCompressionOptions) Init() NSUserInterfaceCompressionOptions {
	rv := objc.Send[NSUserInterfaceCompressionOptions](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUserInterfaceCompressionOptions) Autorelease() NSUserInterfaceCompressionOptions {
	rv := objc.Send[NSUserInterfaceCompressionOptions](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUserInterfaceCompressionOptions creates a new NSUserInterfaceCompressionOptions instance.
func NewNSUserInterfaceCompressionOptions() NSUserInterfaceCompressionOptions {
	class := getNSUserInterfaceCompressionOptionsClass()
	rv := objc.Send[NSUserInterfaceCompressionOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an option object from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(coder:)
func NewUserInterfaceCompressionOptionsWithCoder(coder foundation.INSCoder) NSUserInterfaceCompressionOptions {
	instance := getNSUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSUserInterfaceCompressionOptionsFromID(rv)
}

// Creates an option object that represents the union of the supplied options.
//
// options: A set of [NSUserInterfaceCompressionOptions] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(options:)
func NewUserInterfaceCompressionOptionsWithCompressionOptions(options foundation.INSSet) NSUserInterfaceCompressionOptions {
	instance := getNSUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompressionOptions:"), options)
	return NSUserInterfaceCompressionOptionsFromID(rv)
}

// Creates an option object with the given identifier string.
//
// # Discussion
// 
// Use this initializer to create custom compression options.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(identifier:)
func NewUserInterfaceCompressionOptionsWithIdentifier(identifier string) NSUserInterfaceCompressionOptions {
	instance := getNSUserInterfaceCompressionOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	return NSUserInterfaceCompressionOptionsFromID(rv)
}

// Creates an option object that represents the union of the supplied options.
//
// options: A set of [NSUserInterfaceCompressionOptions] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(options:)
func (u NSUserInterfaceCompressionOptions) InitWithCompressionOptions(options foundation.INSSet) NSUserInterfaceCompressionOptions {
	rv := objc.Send[NSUserInterfaceCompressionOptions](u.ID, objc.Sel("initWithCompressionOptions:"), options)
	return rv
}
// Creates an option object with the given identifier string.
//
// # Discussion
// 
// Use this initializer to create custom compression options.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(identifier:)
func (u NSUserInterfaceCompressionOptions) InitWithIdentifier(identifier string) NSUserInterfaceCompressionOptions {
	rv := objc.Send[NSUserInterfaceCompressionOptions](u.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	return rv
}
// Creates an option object from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/init(coder:)
func (u NSUserInterfaceCompressionOptions) InitWithCoder(coder foundation.INSCoder) NSUserInterfaceCompressionOptions {
	rv := objc.Send[NSUserInterfaceCompressionOptions](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Determines whether the supplied compression options are all present in the
// current instance.
//
// options: A compression options object to compare with the current instance.
//
// # Return Value
// 
// Returns [true] if all the supplied options are present in the instance’s
// options, or [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/contains(_:)
func (u NSUserInterfaceCompressionOptions) ContainsOptions(options INSUserInterfaceCompressionOptions) bool {
	rv := objc.Send[bool](u.ID, objc.Sel("containsOptions:"), options)
	return rv
}
// Determines whether the supplied compression options intersect with the
// current instance’s options.
//
// options: A compression options object to compare with the current instance.
//
// # Return Value
// 
// Returns [true] if at least one of the supplied options is present in the
// instance’s options, or [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/intersects(_:)
func (u NSUserInterfaceCompressionOptions) IntersectsOptions(options INSUserInterfaceCompressionOptions) bool {
	rv := objc.Send[bool](u.ID, objc.Sel("intersectsOptions:"), options)
	return rv
}
// Creates a new compression options object representing the union with the
// provided options.
//
// options: A set of compression options to add to the current object.
//
// # Return Value
// 
// A new [NSCompressibleUserInterfaceOptions] object which represents the
// union with the supplied compression options.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/union(_:)
func (u NSUserInterfaceCompressionOptions) OptionsByAddingOptions(options INSUserInterfaceCompressionOptions) INSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("optionsByAddingOptions:"), options)
	return NSUserInterfaceCompressionOptionsFromID(rv)
}
// Creates a new compression options object with the supplied options removed.
//
// options: A set of compression options to remove from the current object.
//
// # Return Value
// 
// A new [NSCompressibleUserInterfaceOptions] object with the supplied options
// removed.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/subtracting(_:)
func (u NSUserInterfaceCompressionOptions) OptionsByRemovingOptions(options INSUserInterfaceCompressionOptions) INSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("optionsByRemovingOptions:"), options)
	return NSUserInterfaceCompressionOptionsFromID(rv)
}
func (u NSUserInterfaceCompressionOptions) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that denotes whether the option is empty.
//
// # Discussion
// 
// Returns [true] if the option is equivalent to the empty set, or [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/isEmpty
func (u NSUserInterfaceCompressionOptions) Empty() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isEmpty"))
	return rv
}

// An option specifying that views should hide their images.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideImages
func (_NSUserInterfaceCompressionOptionsClass NSUserInterfaceCompressionOptionsClass) HideImagesOption() NSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](objc.ID(_NSUserInterfaceCompressionOptionsClass.class), objc.Sel("hideImagesOption"))
	return NSUserInterfaceCompressionOptionsFromID(objc.ID(rv))
}
// An option specifying that views should hide their text.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/hideText
func (_NSUserInterfaceCompressionOptionsClass NSUserInterfaceCompressionOptionsClass) HideTextOption() NSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](objc.ID(_NSUserInterfaceCompressionOptionsClass.class), objc.Sel("hideTextOption"))
	return NSUserInterfaceCompressionOptionsFromID(objc.ID(rv))
}
// An option specifying that views should reduce their internal metrics.
//
// # Discussion
// 
// Use this compression option to reduce the padding in system controls.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/reduceMetrics
func (_NSUserInterfaceCompressionOptionsClass NSUserInterfaceCompressionOptionsClass) ReduceMetricsOption() NSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](objc.ID(_NSUserInterfaceCompressionOptionsClass.class), objc.Sel("reduceMetricsOption"))
	return NSUserInterfaceCompressionOptionsFromID(objc.ID(rv))
}
// An option specifying that views should no longer maintain equal width
// constraints.
//
// # Discussion
// 
// This option is handled by the system, and no action is required by the
// views.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/breakEqualWidths
func (_NSUserInterfaceCompressionOptionsClass NSUserInterfaceCompressionOptionsClass) BreakEqualWidthsOption() NSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](objc.ID(_NSUserInterfaceCompressionOptionsClass.class), objc.Sel("breakEqualWidthsOption"))
	return NSUserInterfaceCompressionOptionsFromID(objc.ID(rv))
}
// An option that represents the union of all standard compression options.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceCompressionOptions/standardOptions
func (_NSUserInterfaceCompressionOptionsClass NSUserInterfaceCompressionOptionsClass) StandardOptions() NSUserInterfaceCompressionOptions {
	rv := objc.Send[objc.ID](objc.ID(_NSUserInterfaceCompressionOptionsClass.class), objc.Sel("standardOptions"))
	return NSUserInterfaceCompressionOptionsFromID(objc.ID(rv))
}

