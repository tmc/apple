// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXMathExpressionTable] class.
var (
	_AXMathExpressionTableClass     AXMathExpressionTableClass
	_AXMathExpressionTableClassOnce sync.Once
)

func getAXMathExpressionTableClass() AXMathExpressionTableClass {
	_AXMathExpressionTableClassOnce.Do(func() {
		_AXMathExpressionTableClass = AXMathExpressionTableClass{class: objc.GetClass("AXMathExpressionTable")}
	})
	return _AXMathExpressionTableClass
}

// GetAXMathExpressionTableClass returns the class object for AXMathExpressionTable.
func GetAXMathExpressionTableClass() AXMathExpressionTableClass {
	return getAXMathExpressionTableClass()
}

type AXMathExpressionTableClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionTableClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionTableClass) Alloc() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionTable.InitWithExpressions]
//
// # Instance Properties
//
//   - [AXMathExpressionTable.Expressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable
type AXMathExpressionTable struct {
	AXMathExpression
}

// AXMathExpressionTableFromID constructs a [AXMathExpressionTable] from an objc.ID.
func AXMathExpressionTableFromID(id objc.ID) AXMathExpressionTable {
	return AXMathExpressionTable{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionTable adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionTable] class.
//
// # Initializers
//
//   - [IAXMathExpressionTable.InitWithExpressions]
//
// # Instance Properties
//
//   - [IAXMathExpressionTable.Expressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable
type IAXMathExpressionTable interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithExpressions(expressions []AXMathExpression) AXMathExpressionTable

	// Topic: Instance Properties

	Expressions() []AXMathExpression
}

// Init initializes the instance.
func (a AXMathExpressionTable) Init() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionTable) Autorelease() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionTable creates a new AXMathExpressionTable instance.
func NewAXMathExpressionTable() AXMathExpressionTable {
	class := getAXMathExpressionTableClass()
	rv := objc.Send[AXMathExpressionTable](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable/init(expressions:)
func NewAXMathExpressionTableWithExpressions(expressions []AXMathExpression) AXMathExpressionTable {
	instance := getAXMathExpressionTableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExpressions:"), objectivec.IObjectSliceToNSArray(expressions))
	return AXMathExpressionTableFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable/init(expressions:)
func (a AXMathExpressionTable) InitWithExpressions(expressions []AXMathExpression) AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](a.ID, objc.Sel("initWithExpressions:"), objectivec.IObjectSliceToNSArray(expressions))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable/expressions
func (a AXMathExpressionTable) Expressions() []AXMathExpression {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("expressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpression {
		return AXMathExpressionFromID(id)
	})
}
