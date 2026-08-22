// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXMathExpressionRow] class.
var (
	_AXMathExpressionRowClass     AXMathExpressionRowClass
	_AXMathExpressionRowClassOnce sync.Once
)

func getAXMathExpressionRowClass() AXMathExpressionRowClass {
	_AXMathExpressionRowClassOnce.Do(func() {
		_AXMathExpressionRowClass = AXMathExpressionRowClass{class: objc.GetClass("AXMathExpressionRow")}
	})
	return _AXMathExpressionRowClass
}

// GetAXMathExpressionRowClass returns the class object for AXMathExpressionRow.
func GetAXMathExpressionRowClass() AXMathExpressionRowClass {
	return getAXMathExpressionRowClass()
}

type AXMathExpressionRowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionRowClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionRowClass) Alloc() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionRow.InitWithExpressions]
//
// # Instance Properties
//
//   - [AXMathExpressionRow.Expressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow
type AXMathExpressionRow struct {
	AXMathExpression
}

// AXMathExpressionRowFromID constructs a [AXMathExpressionRow] from an objc.ID.
func AXMathExpressionRowFromID(id objc.ID) AXMathExpressionRow {
	return AXMathExpressionRow{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionRow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionRow] class.
//
// # Initializers
//
//   - [IAXMathExpressionRow.InitWithExpressions]
//
// # Instance Properties
//
//   - [IAXMathExpressionRow.Expressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow
type IAXMathExpressionRow interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithExpressions(expressions []AXMathExpression) AXMathExpressionRow

	// Topic: Instance Properties

	Expressions() []AXMathExpression
}

// Init initializes the instance.
func (a AXMathExpressionRow) Init() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionRow) Autorelease() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionRow creates a new AXMathExpressionRow instance.
func NewAXMathExpressionRow() AXMathExpressionRow {
	class := getAXMathExpressionRowClass()
	rv := objc.Send[AXMathExpressionRow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow/init(expressions:)
func NewAXMathExpressionRowWithExpressions(expressions []AXMathExpression) AXMathExpressionRow {
	instance := getAXMathExpressionRowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExpressions:"), objectivec.IObjectSliceToNSArray(expressions))
	return AXMathExpressionRowFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow/init(expressions:)
func (a AXMathExpressionRow) InitWithExpressions(expressions []AXMathExpression) AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](a.ID, objc.Sel("initWithExpressions:"), objectivec.IObjectSliceToNSArray(expressions))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow/expressions
func (a AXMathExpressionRow) Expressions() []AXMathExpression {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("expressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpression {
		return AXMathExpressionFromID(id)
	})
}
