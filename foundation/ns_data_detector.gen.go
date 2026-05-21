// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSDataDetector] class.
var (
	_NSDataDetectorClass     NSDataDetectorClass
	_NSDataDetectorClassOnce sync.Once
)

func getNSDataDetectorClass() NSDataDetectorClass {
	_NSDataDetectorClassOnce.Do(func() {
		_NSDataDetectorClass = NSDataDetectorClass{class: objc.GetClass("NSDataDetector")}
	})
	return _NSDataDetectorClass
}

// GetNSDataDetectorClass returns the class object for NSDataDetector.
func GetNSDataDetectorClass() NSDataDetectorClass {
	return getNSDataDetectorClass()
}

type NSDataDetectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSDataDetectorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDataDetectorClass) Alloc() NSDataDetector {
	rv := objc.Send[NSDataDetector](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A specialized regular expression object that matches natural language text
// for predefined data patterns.
//
// # Overview
//
// Find dates, addresses, links, phone numbers, and transit information in
// natural language text with [NSDataDetector].
//
// [NSDataDetector] returns the results of matching content in
// [NSTextCheckingResult] objects. The [NSTextCheckingResult] objects that
// [NSDataDetector] returns are different from those that
// [NSRegularExpression] returns. The results are one of the data detector’s
// types and contain the corresponding properties. For example, results of
// type [NSTextCheckingTypeDate] have a [NSTextCheckingResult.Date],
// [NSTextCheckingResult.TimeZone], and [NSTextCheckingResult.Duration]; and
// results of type [NSTextCheckingTypeLink] have a [NSTextCheckingResult.URL].
//
// # Examples
//
// The following shows several graduated examples of using the
// [NSDataDetector] class.
//
// This code fragment creates a data detector that finds URL links and phone
// numbers. If an error occurs, it returns in `error`.
//
// After creating the data detector instance, you can determine the number of
// matches within a range of a string using the [NSRegularExpression] method
// [NSRegularExpression.NumberOfMatchesInStringOptionsRange].
//
// If you’re interested only in the overall range of the first match, the
// [NSRegularExpression.NumberOfMatchesInStringOptionsRange] method provides
// it. However, with data detectors, this is less likely than with regular
// expressions because clients are usually interested in additional
// information as well.
//
// The additional information available depends on the type of the result. For
// results of type [NSTextCheckingTypeLink], it’s the [URL] property
// that’s significant. For results of type [NSTextCheckingTypePhoneNumber] ,
// it’s the `phoneNumber` property instead.
//
// The [NSRegularExpression.MatchesInStringOptionsRange] method is similar to
// [NSRegularExpression.FirstMatchInStringOptionsRange], except that it
// returns all matches rather than only the first. The following code fragment
// finds all the matches for links and phone numbers in a string:
//
// The [NSRegularExpression] block object enumerator is the most general and
// flexible of the matching methods. It allows you to iterate through matches
// in a string, performing arbitrary actions on each as specified by the code
// in the block, and to stop partway through if desired. In the following code
// fragment, the iteration stops after finding a certain number of matches:
//
// # Creating data detector instances
//
//   - [NSDataDetector.InitWithTypesError]: Initializes and returns a data detector instance.
//
// # Getting the checking types
//
//   - [NSDataDetector.CheckingTypes]: Returns the checking types for the data detector.
//
// See: https://developer.apple.com/documentation/Foundation/NSDataDetector
type NSDataDetector struct {
	NSRegularExpression
}

// NSDataDetectorFromID constructs a [NSDataDetector] from an objc.ID.
//
// A specialized regular expression object that matches natural language text
// for predefined data patterns.
func NSDataDetectorFromID(id objc.ID) NSDataDetector {
	return NSDataDetector{NSRegularExpression: NSRegularExpressionFromID(id)}
}

// NOTE: NSDataDetector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDataDetector] class.
//
// # Creating data detector instances
//
//   - [INSDataDetector.InitWithTypesError]: Initializes and returns a data detector instance.
//
// # Getting the checking types
//
//   - [INSDataDetector.CheckingTypes]: Returns the checking types for the data detector.
//
// See: https://developer.apple.com/documentation/Foundation/NSDataDetector
type INSDataDetector interface {
	INSRegularExpression

	// Topic: Creating data detector instances

	// Initializes and returns a data detector instance.
	InitWithTypesError(checkingTypes NSTextCheckingTypes) (NSDataDetector, error)

	// Topic: Getting the checking types

	// Returns the checking types for the data detector.
	CheckingTypes() NSTextCheckingTypes
}

// Init initializes the instance.
func (d NSDataDetector) Init() NSDataDetector {
	rv := objc.Send[NSDataDetector](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDataDetector) Autorelease() NSDataDetector {
	rv := objc.Send[NSDataDetector](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDataDetector creates a new NSDataDetector instance.
func NewNSDataDetector() NSDataDetector {
	class := getNSDataDetectorClass()
	rv := objc.Send[NSDataDetector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/init(coder:)
func NewDataDetectorWithCoder(coder INSCoder) NSDataDetector {
	instance := getNSDataDetectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDataDetectorFromID(rv)
}

// Returns an initialized NSRegularExpression instance with the specified
// regular expression pattern and options.
//
// pattern: The regular expression pattern to compile.
//
// options: The regular expression options that are applied to the expression during
// matching. See [NSRegularExpression.Options] for possible values.
//
// # Return Value
//
// An instance of [NSRegularExpression] for the specified regular expression
// and options.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/init(pattern:options:)
//
// [NSRegularExpression.Options]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
func NewDataDetectorWithPatternOptionsError(pattern string, options NSRegularExpressionOptions) (NSDataDetector, error) {
	var errorPtr objc.ID
	instance := getNSDataDetectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPattern:options:error:"), objc.String(pattern), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDataDetector{}, NSErrorFrom(errorPtr)
	}
	return NSDataDetectorFromID(rv), nil
}

// Initializes and returns a data detector instance.
//
// checkingTypes: The checking types. The supported checking types are a subset of the types
// [NSTextCheckingResult.CheckingType]. Those constants can be combined using
// the C-bitwise OR operator.
//
// # Return Value
//
// Returns the newly initialized data detector. If an error was encountered
// returns `nil`, and `error` contains the error.
//
// # Discussion
//
// Currently, the supported data detectors `checkingTypes` are:
// [NSTextCheckingTypeDate], [NSTextCheckingTypeAddress],
// [NSTextCheckingTypeLink], [NSTextCheckingTypePhoneNumber], and
// [NSTextCheckingTypeTransitInformation].
//
// See: https://developer.apple.com/documentation/Foundation/NSDataDetector/init(types:)
//
// [NSTextCheckingResult.CheckingType]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
func NewDataDetectorWithTypesError(checkingTypes NSTextCheckingTypes) (NSDataDetector, error) {
	var errorPtr objc.ID
	instance := getNSDataDetectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTypes:error:"), checkingTypes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDataDetector{}, NSErrorFrom(errorPtr)
	}
	return NSDataDetectorFromID(rv), nil
}

// Initializes and returns a data detector instance.
//
// checkingTypes: The checking types. The supported checking types are a subset of the types
// [NSTextCheckingResult.CheckingType]. Those constants can be combined using
// the C-bitwise OR operator.
//
// # Return Value
//
// Returns the newly initialized data detector. If an error was encountered
// returns `nil`, and `error` contains the error.
//
// # Discussion
//
// Currently, the supported data detectors `checkingTypes` are:
// [NSTextCheckingTypeDate], [NSTextCheckingTypeAddress],
// [NSTextCheckingTypeLink], [NSTextCheckingTypePhoneNumber], and
// [NSTextCheckingTypeTransitInformation].
//
// See: https://developer.apple.com/documentation/Foundation/NSDataDetector/init(types:)
//
// [NSTextCheckingResult.CheckingType]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
func (d NSDataDetector) InitWithTypesError(checkingTypes NSTextCheckingTypes) (NSDataDetector, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithTypes:error:"), checkingTypes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDataDetector{}, NSErrorFrom(errorPtr)
	}
	return NSDataDetectorFromID(rv), nil

}

// Creates and returns a new data detector instance.
//
// checkingTypes: The checking types. The supported checking types are a subset of the types
// specified in [NSTextCheckingResult.CheckingType]. Those constants can be
// combined using the C-bitwise OR operator.
//
// error: An out parameter that if an error occurs during initialization contains the
// encountered error.
//
// # Return Value
//
// Returns the newly initialized data detector. If an error was encountered
// returns `nil`, and `error` contains the error.
//
// # Discussion
//
// Currently, the supported data detectors `checkingTypes` are:
// [NSTextCheckingTypeDate], [NSTextCheckingTypeAddress],
// [NSTextCheckingTypeLink], [NSTextCheckingTypePhoneNumber], and
// [NSTextCheckingTypeTransitInformation].
//
// See: https://developer.apple.com/documentation/Foundation/NSDataDetector/dataDetectorWithTypes:error:
//
// [NSTextCheckingResult.CheckingType]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
func (_NSDataDetectorClass NSDataDetectorClass) DataDetectorWithTypesError(checkingTypes NSTextCheckingTypes) (NSDataDetector, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSDataDetectorClass.class), objc.Sel("dataDetectorWithTypes:error:"), checkingTypes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDataDetector{}, NSErrorFrom(errorPtr)
	}
	return NSDataDetectorFromID(rv), nil

}

// Returns the checking types for the data detector.
//
// # Discussion
//
// The supported subset of checking types are specified in
// [NSTextCheckingResult.CheckingType]. Those constants can be combined using
// the C-bitwise OR operator.
//
// Currently, the supported data detectors `checkingTypes` are:
// [NSTextCheckingTypeDate], [NSTextCheckingTypeAddress],
// [NSTextCheckingTypeLink], [NSTextCheckingTypePhoneNumber], and
// [NSTextCheckingTypeTransitInformation].
//
// See: https://developer.apple.com/documentation/Foundation/NSDataDetector/checkingTypes
//
// [NSTextCheckingResult.CheckingType]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
func (d NSDataDetector) CheckingTypes() NSTextCheckingTypes {
	rv := objc.Send[NSTextCheckingTypes](d.ID, objc.Sel("checkingTypes"))
	return NSTextCheckingTypes(rv)
}
