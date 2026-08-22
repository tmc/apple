// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXMathExpressionFenced] class.
var (
	_AXMathExpressionFencedClass     AXMathExpressionFencedClass
	_AXMathExpressionFencedClassOnce sync.Once
)

func getAXMathExpressionFencedClass() AXMathExpressionFencedClass {
	_AXMathExpressionFencedClassOnce.Do(func() {
		_AXMathExpressionFencedClass = AXMathExpressionFencedClass{class: objc.GetClass("AXMathExpressionFenced")}
	})
	return _AXMathExpressionFencedClass
}

// GetAXMathExpressionFencedClass returns the class object for AXMathExpressionFenced.
func GetAXMathExpressionFencedClass() AXMathExpressionFencedClass {
	return getAXMathExpressionFencedClass()
}

type AXMathExpressionFencedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionFencedClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionFencedClass) Alloc() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionFenced.InitWithExpressionsOpenStringCloseString]
//
// # Instance Properties
//
//   - [AXMathExpressionFenced.CloseString]
//   - [AXMathExpressionFenced.Expressions]
//   - [AXMathExpressionFenced.OpenString]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced
type AXMathExpressionFenced struct {
	AXMathExpression
}

// AXMathExpressionFencedFromID constructs a [AXMathExpressionFenced] from an objc.ID.
func AXMathExpressionFencedFromID(id objc.ID) AXMathExpressionFenced {
	return AXMathExpressionFenced{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionFenced adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionFenced] class.
//
// # Initializers
//
//   - [IAXMathExpressionFenced.InitWithExpressionsOpenStringCloseString]
//
// # Instance Properties
//
//   - [IAXMathExpressionFenced.CloseString]
//   - [IAXMathExpressionFenced.Expressions]
//   - [IAXMathExpressionFenced.OpenString]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced
type IAXMathExpressionFenced interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithExpressionsOpenStringCloseString(expressions []AXMathExpression, openString string, closeString string) AXMathExpressionFenced

	// Topic: Instance Properties

	CloseString() string
	Expressions() []AXMathExpression
	OpenString() string
}

// Init initializes the instance.
func (a AXMathExpressionFenced) Init() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionFenced) Autorelease() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionFenced creates a new AXMathExpressionFenced instance.
func NewAXMathExpressionFenced() AXMathExpressionFenced {
	class := getAXMathExpressionFencedClass()
	rv := objc.Send[AXMathExpressionFenced](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/init(expressions:open:close:)
func NewAXMathExpressionFencedWithExpressionsOpenStringCloseString(expressions []AXMathExpression, openString string, closeString string) AXMathExpressionFenced {
	instance := getAXMathExpressionFencedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExpressions:openString:closeString:"), objectivec.IObjectSliceToNSArray(expressions), objc.String(openString), objc.String(closeString))
	return AXMathExpressionFencedFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/init(expressions:open:close:)
func (a AXMathExpressionFenced) InitWithExpressionsOpenStringCloseString(expressions []AXMathExpression, openString string, closeString string) AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](a.ID, objc.Sel("initWithExpressions:openString:closeString:"), objectivec.IObjectSliceToNSArray(expressions), objc.String(openString), objc.String(closeString))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/closeString
func (a AXMathExpressionFenced) CloseString() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("closeString"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/expressions
func (a AXMathExpressionFenced) Expressions() []AXMathExpression {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("expressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpression {
		return AXMathExpressionFromID(id)
	})
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/openString
func (a AXMathExpressionFenced) OpenString() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("openString"))
	return foundation.NSStringFromID(rv).String()
}
