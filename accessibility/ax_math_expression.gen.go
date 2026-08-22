// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXMathExpression] class.
var (
	_AXMathExpressionClass     AXMathExpressionClass
	_AXMathExpressionClassOnce sync.Once
)

func getAXMathExpressionClass() AXMathExpressionClass {
	_AXMathExpressionClassOnce.Do(func() {
		_AXMathExpressionClass = AXMathExpressionClass{class: objc.GetClass("AXMathExpression")}
	})
	return _AXMathExpressionClass
}

// GetAXMathExpressionClass returns the class object for AXMathExpression.
func GetAXMathExpressionClass() AXMathExpressionClass {
	return getAXMathExpressionClass()
}

type AXMathExpressionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionClass) Alloc() AXMathExpression {
	rv := objc.Send[AXMathExpression](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpression
type AXMathExpression struct {
	objectivec.Object
}

// AXMathExpressionFromID constructs a [AXMathExpression] from an objc.ID.
func AXMathExpressionFromID(id objc.ID) AXMathExpression {
	return AXMathExpression{objectivec.Object{ID: id}}
}

// NOTE: AXMathExpression adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpression] class.
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpression
type IAXMathExpression interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a AXMathExpression) Init() AXMathExpression {
	rv := objc.Send[AXMathExpression](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpression) Autorelease() AXMathExpression {
	rv := objc.Send[AXMathExpression](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpression creates a new AXMathExpression instance.
func NewAXMathExpression() AXMathExpression {
	class := getAXMathExpressionClass()
	rv := objc.Send[AXMathExpression](objc.ID(class.class), objc.Sel("new"))
	return rv
}
