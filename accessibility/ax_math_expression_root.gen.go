// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXMathExpressionRoot] class.
var (
	_AXMathExpressionRootClass     AXMathExpressionRootClass
	_AXMathExpressionRootClassOnce sync.Once
)

func getAXMathExpressionRootClass() AXMathExpressionRootClass {
	_AXMathExpressionRootClassOnce.Do(func() {
		_AXMathExpressionRootClass = AXMathExpressionRootClass{class: objc.GetClass("AXMathExpressionRoot")}
	})
	return _AXMathExpressionRootClass
}

// GetAXMathExpressionRootClass returns the class object for AXMathExpressionRoot.
func GetAXMathExpressionRootClass() AXMathExpressionRootClass {
	return getAXMathExpressionRootClass()
}

type AXMathExpressionRootClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionRootClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionRootClass) Alloc() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionRoot.InitWithRadicandExpressionsRootIndexExpression]
//
// # Instance Properties
//
//   - [AXMathExpressionRoot.RadicandExpressions]
//   - [AXMathExpressionRoot.RootIndexExpression]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot
type AXMathExpressionRoot struct {
	AXMathExpression
}

// AXMathExpressionRootFromID constructs a [AXMathExpressionRoot] from an objc.ID.
func AXMathExpressionRootFromID(id objc.ID) AXMathExpressionRoot {
	return AXMathExpressionRoot{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionRoot adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionRoot] class.
//
// # Initializers
//
//   - [IAXMathExpressionRoot.InitWithRadicandExpressionsRootIndexExpression]
//
// # Instance Properties
//
//   - [IAXMathExpressionRoot.RadicandExpressions]
//   - [IAXMathExpressionRoot.RootIndexExpression]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot
type IAXMathExpressionRoot interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithRadicandExpressionsRootIndexExpression(radicandExpressions []AXMathExpression, rootIndexExpression IAXMathExpression) AXMathExpressionRoot

	// Topic: Instance Properties

	RadicandExpressions() []AXMathExpression
	RootIndexExpression() IAXMathExpression
}

// Init initializes the instance.
func (a AXMathExpressionRoot) Init() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionRoot) Autorelease() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionRoot creates a new AXMathExpressionRoot instance.
func NewAXMathExpressionRoot() AXMathExpressionRoot {
	class := getAXMathExpressionRootClass()
	rv := objc.Send[AXMathExpressionRoot](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/init(radicandExpressions:rootIndexExpression:)
func NewAXMathExpressionRootWithRadicandExpressionsRootIndexExpression(radicandExpressions []AXMathExpression, rootIndexExpression IAXMathExpression) AXMathExpressionRoot {
	instance := getAXMathExpressionRootClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRadicandExpressions:rootIndexExpression:"), objectivec.IObjectSliceToNSArray(radicandExpressions), rootIndexExpression)
	return AXMathExpressionRootFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/init(radicandExpressions:rootIndexExpression:)
func (a AXMathExpressionRoot) InitWithRadicandExpressionsRootIndexExpression(radicandExpressions []AXMathExpression, rootIndexExpression IAXMathExpression) AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](a.ID, objc.Sel("initWithRadicandExpressions:rootIndexExpression:"), objectivec.IObjectSliceToNSArray(radicandExpressions), rootIndexExpression)
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/radicandExpressions
func (a AXMathExpressionRoot) RadicandExpressions() []AXMathExpression {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("radicandExpressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpression {
		return AXMathExpressionFromID(id)
	})
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/rootIndexExpression
func (a AXMathExpressionRoot) RootIndexExpression() IAXMathExpression {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("rootIndexExpression"))
	return AXMathExpressionFromID(objc.ID(rv))
}
