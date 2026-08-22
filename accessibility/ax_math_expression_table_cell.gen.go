// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXMathExpressionTableCell] class.
var (
	_AXMathExpressionTableCellClass     AXMathExpressionTableCellClass
	_AXMathExpressionTableCellClassOnce sync.Once
)

func getAXMathExpressionTableCellClass() AXMathExpressionTableCellClass {
	_AXMathExpressionTableCellClassOnce.Do(func() {
		_AXMathExpressionTableCellClass = AXMathExpressionTableCellClass{class: objc.GetClass("AXMathExpressionTableCell")}
	})
	return _AXMathExpressionTableCellClass
}

// GetAXMathExpressionTableCellClass returns the class object for AXMathExpressionTableCell.
func GetAXMathExpressionTableCellClass() AXMathExpressionTableCellClass {
	return getAXMathExpressionTableCellClass()
}

type AXMathExpressionTableCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionTableCellClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionTableCellClass) Alloc() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionTableCell.InitWithExpressions]
//
// # Instance Properties
//
//   - [AXMathExpressionTableCell.Expressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell
type AXMathExpressionTableCell struct {
	AXMathExpression
}

// AXMathExpressionTableCellFromID constructs a [AXMathExpressionTableCell] from an objc.ID.
func AXMathExpressionTableCellFromID(id objc.ID) AXMathExpressionTableCell {
	return AXMathExpressionTableCell{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionTableCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionTableCell] class.
//
// # Initializers
//
//   - [IAXMathExpressionTableCell.InitWithExpressions]
//
// # Instance Properties
//
//   - [IAXMathExpressionTableCell.Expressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell
type IAXMathExpressionTableCell interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithExpressions(expressions []AXMathExpression) AXMathExpressionTableCell

	// Topic: Instance Properties

	Expressions() []AXMathExpression
}

// Init initializes the instance.
func (a AXMathExpressionTableCell) Init() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionTableCell) Autorelease() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionTableCell creates a new AXMathExpressionTableCell instance.
func NewAXMathExpressionTableCell() AXMathExpressionTableCell {
	class := getAXMathExpressionTableCellClass()
	rv := objc.Send[AXMathExpressionTableCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell/init(expressions:)
func NewAXMathExpressionTableCellWithExpressions(expressions []AXMathExpression) AXMathExpressionTableCell {
	instance := getAXMathExpressionTableCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExpressions:"), objectivec.IObjectSliceToNSArray(expressions))
	return AXMathExpressionTableCellFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell/init(expressions:)
func (a AXMathExpressionTableCell) InitWithExpressions(expressions []AXMathExpression) AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](a.ID, objc.Sel("initWithExpressions:"), objectivec.IObjectSliceToNSArray(expressions))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell/expressions
func (a AXMathExpressionTableCell) Expressions() []AXMathExpression {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("expressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpression {
		return AXMathExpressionFromID(id)
	})
}
