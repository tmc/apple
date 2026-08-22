// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXMathExpressionTableRow] class.
var (
	_AXMathExpressionTableRowClass     AXMathExpressionTableRowClass
	_AXMathExpressionTableRowClassOnce sync.Once
)

func getAXMathExpressionTableRowClass() AXMathExpressionTableRowClass {
	_AXMathExpressionTableRowClassOnce.Do(func() {
		_AXMathExpressionTableRowClass = AXMathExpressionTableRowClass{class: objc.GetClass("AXMathExpressionTableRow")}
	})
	return _AXMathExpressionTableRowClass
}

// GetAXMathExpressionTableRowClass returns the class object for AXMathExpressionTableRow.
func GetAXMathExpressionTableRowClass() AXMathExpressionTableRowClass {
	return getAXMathExpressionTableRowClass()
}

type AXMathExpressionTableRowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionTableRowClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionTableRowClass) Alloc() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionTableRow.InitWithExpressions]
//
// # Instance Properties
//
//   - [AXMathExpressionTableRow.Expressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow
type AXMathExpressionTableRow struct {
	AXMathExpression
}

// AXMathExpressionTableRowFromID constructs a [AXMathExpressionTableRow] from an objc.ID.
func AXMathExpressionTableRowFromID(id objc.ID) AXMathExpressionTableRow {
	return AXMathExpressionTableRow{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionTableRow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionTableRow] class.
//
// # Initializers
//
//   - [IAXMathExpressionTableRow.InitWithExpressions]
//
// # Instance Properties
//
//   - [IAXMathExpressionTableRow.Expressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow
type IAXMathExpressionTableRow interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithExpressions(expressions []AXMathExpression) AXMathExpressionTableRow

	// Topic: Instance Properties

	Expressions() []AXMathExpression
}

// Init initializes the instance.
func (a AXMathExpressionTableRow) Init() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionTableRow) Autorelease() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionTableRow creates a new AXMathExpressionTableRow instance.
func NewAXMathExpressionTableRow() AXMathExpressionTableRow {
	class := getAXMathExpressionTableRowClass()
	rv := objc.Send[AXMathExpressionTableRow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow/init(expressions:)
func NewAXMathExpressionTableRowWithExpressions(expressions []AXMathExpression) AXMathExpressionTableRow {
	instance := getAXMathExpressionTableRowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExpressions:"), objectivec.IObjectSliceToNSArray(expressions))
	return AXMathExpressionTableRowFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow/init(expressions:)
func (a AXMathExpressionTableRow) InitWithExpressions(expressions []AXMathExpression) AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](a.ID, objc.Sel("initWithExpressions:"), objectivec.IObjectSliceToNSArray(expressions))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow/expressions
func (a AXMathExpressionTableRow) Expressions() []AXMathExpression {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("expressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpression {
		return AXMathExpressionFromID(id)
	})
}
