// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXBrailleTranslationResult] class.
var (
	_AXBrailleTranslationResultClass     AXBrailleTranslationResultClass
	_AXBrailleTranslationResultClassOnce sync.Once
)

func getAXBrailleTranslationResultClass() AXBrailleTranslationResultClass {
	_AXBrailleTranslationResultClassOnce.Do(func() {
		_AXBrailleTranslationResultClass = AXBrailleTranslationResultClass{class: objc.GetClass("AXBrailleTranslationResult")}
	})
	return _AXBrailleTranslationResultClass
}

// GetAXBrailleTranslationResultClass returns the class object for AXBrailleTranslationResult.
func GetAXBrailleTranslationResultClass() AXBrailleTranslationResultClass {
	return getAXBrailleTranslationResultClass()
}

type AXBrailleTranslationResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXBrailleTranslationResultClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXBrailleTranslationResultClass) Alloc() AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// The result of translation or back-translation.
//
// # Initializers
//
//   - [AXBrailleTranslationResult.InitWithCoder]
//
// # Instance Properties
//
//   - [AXBrailleTranslationResult.ResultString]: The resulting string after translation or back-translation.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult
type AXBrailleTranslationResult struct {
	objectivec.Object
}

// AXBrailleTranslationResultFromID constructs a [AXBrailleTranslationResult] from an objc.ID.
//
// The result of translation or back-translation.
func AXBrailleTranslationResultFromID(id objc.ID) AXBrailleTranslationResult {
	return AXBrailleTranslationResult{objectivec.Object{ID: id}}
}

// NOTE: AXBrailleTranslationResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXBrailleTranslationResult] class.
//
// # Initializers
//
//   - [IAXBrailleTranslationResult.InitWithCoder]
//
// # Instance Properties
//
//   - [IAXBrailleTranslationResult.ResultString]: The resulting string after translation or back-translation.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult
type IAXBrailleTranslationResult interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) AXBrailleTranslationResult

	// Topic: Instance Properties

	// The resulting string after translation or back-translation.
	ResultString() string

	// An array of integers that has the same length as the resultString. locationMap[i]-th character in the input string corresponds to resultString[i].
	LocationMap() []foundation.NSNumber
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a AXBrailleTranslationResult) Init() AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXBrailleTranslationResult) Autorelease() AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXBrailleTranslationResult creates a new AXBrailleTranslationResult instance.
func NewAXBrailleTranslationResult() AXBrailleTranslationResult {
	class := getAXBrailleTranslationResultClass()
	rv := objc.Send[AXBrailleTranslationResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult/init(coder:)
func NewAXBrailleTranslationResultWithCoder(coder foundation.INSCoder) AXBrailleTranslationResult {
	instance := getAXBrailleTranslationResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AXBrailleTranslationResultFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult/init(coder:)
func (a AXBrailleTranslationResult) InitWithCoder(coder foundation.INSCoder) AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a AXBrailleTranslationResult) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The resulting string after translation or back-translation.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult/resultString
func (a AXBrailleTranslationResult) ResultString() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("resultString"))
	return foundation.NSStringFromID(rv).String()
}

// An array of integers that has the same length as the resultString.
// locationMap[i]-th character in the input string corresponds to
// resultString[i].
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult/locationMap
func (a AXBrailleTranslationResult) LocationMap() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("locationMap"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
