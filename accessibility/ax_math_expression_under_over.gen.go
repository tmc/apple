// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AXMathExpressionUnderOver] class.
var (
	_AXMathExpressionUnderOverClass     AXMathExpressionUnderOverClass
	_AXMathExpressionUnderOverClassOnce sync.Once
)

func getAXMathExpressionUnderOverClass() AXMathExpressionUnderOverClass {
	_AXMathExpressionUnderOverClassOnce.Do(func() {
		_AXMathExpressionUnderOverClass = AXMathExpressionUnderOverClass{class: objc.GetClass("AXMathExpressionUnderOver")}
	})
	return _AXMathExpressionUnderOverClass
}

// GetAXMathExpressionUnderOverClass returns the class object for AXMathExpressionUnderOver.
func GetAXMathExpressionUnderOverClass() AXMathExpressionUnderOverClass {
	return getAXMathExpressionUnderOverClass()
}

type AXMathExpressionUnderOverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionUnderOverClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionUnderOverClass) Alloc() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionUnderOver.InitWithBaseExpressionUnderExpressionOverExpression]
//
// # Instance Properties
//
//   - [AXMathExpressionUnderOver.BaseExpression]
//   - [AXMathExpressionUnderOver.OverExpression]
//   - [AXMathExpressionUnderOver.UnderExpression]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver
type AXMathExpressionUnderOver struct {
	AXMathExpression
}

// AXMathExpressionUnderOverFromID constructs a [AXMathExpressionUnderOver] from an objc.ID.
func AXMathExpressionUnderOverFromID(id objc.ID) AXMathExpressionUnderOver {
	return AXMathExpressionUnderOver{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionUnderOver adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionUnderOver] class.
//
// # Initializers
//
//   - [IAXMathExpressionUnderOver.InitWithBaseExpressionUnderExpressionOverExpression]
//
// # Instance Properties
//
//   - [IAXMathExpressionUnderOver.BaseExpression]
//   - [IAXMathExpressionUnderOver.OverExpression]
//   - [IAXMathExpressionUnderOver.UnderExpression]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver
type IAXMathExpressionUnderOver interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithBaseExpressionUnderExpressionOverExpression(baseExpression IAXMathExpression, underExpression IAXMathExpression, overExpression IAXMathExpression) AXMathExpressionUnderOver

	// Topic: Instance Properties

	BaseExpression() IAXMathExpression
	OverExpression() IAXMathExpression
	UnderExpression() IAXMathExpression
}

// Init initializes the instance.
func (a AXMathExpressionUnderOver) Init() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionUnderOver) Autorelease() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionUnderOver creates a new AXMathExpressionUnderOver instance.
func NewAXMathExpressionUnderOver() AXMathExpressionUnderOver {
	class := getAXMathExpressionUnderOverClass()
	rv := objc.Send[AXMathExpressionUnderOver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/init(baseExpression:underExpression:overExpression:)
func NewAXMathExpressionUnderOverWithBaseExpressionUnderExpressionOverExpression(baseExpression IAXMathExpression, underExpression IAXMathExpression, overExpression IAXMathExpression) AXMathExpressionUnderOver {
	instance := getAXMathExpressionUnderOverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBaseExpression:underExpression:overExpression:"), baseExpression, underExpression, overExpression)
	return AXMathExpressionUnderOverFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/init(baseExpression:underExpression:overExpression:)
func (a AXMathExpressionUnderOver) InitWithBaseExpressionUnderExpressionOverExpression(baseExpression IAXMathExpression, underExpression IAXMathExpression, overExpression IAXMathExpression) AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](a.ID, objc.Sel("initWithBaseExpression:underExpression:overExpression:"), baseExpression, underExpression, overExpression)
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/baseExpression
func (a AXMathExpressionUnderOver) BaseExpression() IAXMathExpression {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("baseExpression"))
	return AXMathExpressionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/overExpression
func (a AXMathExpressionUnderOver) OverExpression() IAXMathExpression {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("overExpression"))
	return AXMathExpressionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/underExpression
func (a AXMathExpressionUnderOver) UnderExpression() IAXMathExpression {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("underExpression"))
	return AXMathExpressionFromID(objc.ID(rv))
}
