// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AXMathExpressionFraction] class.
var (
	_AXMathExpressionFractionClass     AXMathExpressionFractionClass
	_AXMathExpressionFractionClassOnce sync.Once
)

func getAXMathExpressionFractionClass() AXMathExpressionFractionClass {
	_AXMathExpressionFractionClassOnce.Do(func() {
		_AXMathExpressionFractionClass = AXMathExpressionFractionClass{class: objc.GetClass("AXMathExpressionFraction")}
	})
	return _AXMathExpressionFractionClass
}

// GetAXMathExpressionFractionClass returns the class object for AXMathExpressionFraction.
func GetAXMathExpressionFractionClass() AXMathExpressionFractionClass {
	return getAXMathExpressionFractionClass()
}

type AXMathExpressionFractionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionFractionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionFractionClass) Alloc() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionFraction.InitWithNumeratorExpressionDenimonatorExpression]
//
// # Instance Properties
//
//   - [AXMathExpressionFraction.DenimonatorExpression]
//   - [AXMathExpressionFraction.NumeratorExpression]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction
type AXMathExpressionFraction struct {
	AXMathExpression
}

// AXMathExpressionFractionFromID constructs a [AXMathExpressionFraction] from an objc.ID.
func AXMathExpressionFractionFromID(id objc.ID) AXMathExpressionFraction {
	return AXMathExpressionFraction{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionFraction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionFraction] class.
//
// # Initializers
//
//   - [IAXMathExpressionFraction.InitWithNumeratorExpressionDenimonatorExpression]
//
// # Instance Properties
//
//   - [IAXMathExpressionFraction.DenimonatorExpression]
//   - [IAXMathExpressionFraction.NumeratorExpression]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction
type IAXMathExpressionFraction interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithNumeratorExpressionDenimonatorExpression(numeratorExpression IAXMathExpression, denimonatorExpression IAXMathExpression) AXMathExpressionFraction

	// Topic: Instance Properties

	DenimonatorExpression() IAXMathExpression
	NumeratorExpression() IAXMathExpression
}

// Init initializes the instance.
func (a AXMathExpressionFraction) Init() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionFraction) Autorelease() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionFraction creates a new AXMathExpressionFraction instance.
func NewAXMathExpressionFraction() AXMathExpressionFraction {
	class := getAXMathExpressionFractionClass()
	rv := objc.Send[AXMathExpressionFraction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction/init(numeratorExpression:denimonatorExpression:)
func NewAXMathExpressionFractionWithNumeratorExpressionDenimonatorExpression(numeratorExpression IAXMathExpression, denimonatorExpression IAXMathExpression) AXMathExpressionFraction {
	instance := getAXMathExpressionFractionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNumeratorExpression:denimonatorExpression:"), numeratorExpression, denimonatorExpression)
	return AXMathExpressionFractionFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction/init(numeratorExpression:denimonatorExpression:)
func (a AXMathExpressionFraction) InitWithNumeratorExpressionDenimonatorExpression(numeratorExpression IAXMathExpression, denimonatorExpression IAXMathExpression) AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](a.ID, objc.Sel("initWithNumeratorExpression:denimonatorExpression:"), numeratorExpression, denimonatorExpression)
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction/denimonatorExpression
func (a AXMathExpressionFraction) DenimonatorExpression() IAXMathExpression {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("denimonatorExpression"))
	return AXMathExpressionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction/numeratorExpression
func (a AXMathExpressionFraction) NumeratorExpression() IAXMathExpression {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("numeratorExpression"))
	return AXMathExpressionFromID(objc.ID(rv))
}
