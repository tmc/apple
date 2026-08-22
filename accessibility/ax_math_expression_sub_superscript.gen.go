// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXMathExpressionSubSuperscript] class.
var (
	_AXMathExpressionSubSuperscriptClass     AXMathExpressionSubSuperscriptClass
	_AXMathExpressionSubSuperscriptClassOnce sync.Once
)

func getAXMathExpressionSubSuperscriptClass() AXMathExpressionSubSuperscriptClass {
	_AXMathExpressionSubSuperscriptClassOnce.Do(func() {
		_AXMathExpressionSubSuperscriptClass = AXMathExpressionSubSuperscriptClass{class: objc.GetClass("AXMathExpressionSubSuperscript")}
	})
	return _AXMathExpressionSubSuperscriptClass
}

// GetAXMathExpressionSubSuperscriptClass returns the class object for AXMathExpressionSubSuperscript.
func GetAXMathExpressionSubSuperscriptClass() AXMathExpressionSubSuperscriptClass {
	return getAXMathExpressionSubSuperscriptClass()
}

type AXMathExpressionSubSuperscriptClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionSubSuperscriptClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionSubSuperscriptClass) Alloc() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionSubSuperscript.InitWithBaseExpressionSubscriptExpressionsSuperscriptExpressions]
//
// # Instance Properties
//
//   - [AXMathExpressionSubSuperscript.BaseExpression]
//   - [AXMathExpressionSubSuperscript.SubscriptExpressions]
//   - [AXMathExpressionSubSuperscript.SuperscriptExpressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript
type AXMathExpressionSubSuperscript struct {
	AXMathExpression
}

// AXMathExpressionSubSuperscriptFromID constructs a [AXMathExpressionSubSuperscript] from an objc.ID.
func AXMathExpressionSubSuperscriptFromID(id objc.ID) AXMathExpressionSubSuperscript {
	return AXMathExpressionSubSuperscript{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionSubSuperscript adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionSubSuperscript] class.
//
// # Initializers
//
//   - [IAXMathExpressionSubSuperscript.InitWithBaseExpressionSubscriptExpressionsSuperscriptExpressions]
//
// # Instance Properties
//
//   - [IAXMathExpressionSubSuperscript.BaseExpression]
//   - [IAXMathExpressionSubSuperscript.SubscriptExpressions]
//   - [IAXMathExpressionSubSuperscript.SuperscriptExpressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript
type IAXMathExpressionSubSuperscript interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithBaseExpressionSubscriptExpressionsSuperscriptExpressions(baseExpression []AXMathExpression, subscriptExpressions []AXMathExpression, superscriptExpressions []AXMathExpression) AXMathExpressionSubSuperscript

	// Topic: Instance Properties

	BaseExpression() IAXMathExpression
	SubscriptExpressions() []AXMathExpression
	SuperscriptExpressions() []AXMathExpression
}

// Init initializes the instance.
func (a AXMathExpressionSubSuperscript) Init() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionSubSuperscript) Autorelease() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionSubSuperscript creates a new AXMathExpressionSubSuperscript instance.
func NewAXMathExpressionSubSuperscript() AXMathExpressionSubSuperscript {
	class := getAXMathExpressionSubSuperscriptClass()
	rv := objc.Send[AXMathExpressionSubSuperscript](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/init(baseExpression:subscriptExpressions:superscriptExpressions:)
func NewAXMathExpressionSubSuperscriptWithBaseExpressionSubscriptExpressionsSuperscriptExpressions(baseExpression []AXMathExpression, subscriptExpressions []AXMathExpression, superscriptExpressions []AXMathExpression) AXMathExpressionSubSuperscript {
	instance := getAXMathExpressionSubSuperscriptClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBaseExpression:subscriptExpressions:superscriptExpressions:"), objectivec.IObjectSliceToNSArray(baseExpression), objectivec.IObjectSliceToNSArray(subscriptExpressions), objectivec.IObjectSliceToNSArray(superscriptExpressions))
	return AXMathExpressionSubSuperscriptFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/init(baseExpression:subscriptExpressions:superscriptExpressions:)
func (a AXMathExpressionSubSuperscript) InitWithBaseExpressionSubscriptExpressionsSuperscriptExpressions(baseExpression []AXMathExpression, subscriptExpressions []AXMathExpression, superscriptExpressions []AXMathExpression) AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](a.ID, objc.Sel("initWithBaseExpression:subscriptExpressions:superscriptExpressions:"), objectivec.IObjectSliceToNSArray(baseExpression), objectivec.IObjectSliceToNSArray(subscriptExpressions), objectivec.IObjectSliceToNSArray(superscriptExpressions))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/baseExpression
func (a AXMathExpressionSubSuperscript) BaseExpression() IAXMathExpression {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("baseExpression"))
	return AXMathExpressionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/subscriptExpressions
func (a AXMathExpressionSubSuperscript) SubscriptExpressions() []AXMathExpression {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("subscriptExpressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpression {
		return AXMathExpressionFromID(id)
	})
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/superscriptExpressions
func (a AXMathExpressionSubSuperscript) SuperscriptExpressions() []AXMathExpression {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("superscriptExpressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpression {
		return AXMathExpressionFromID(id)
	})
}
