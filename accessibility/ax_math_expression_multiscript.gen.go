// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXMathExpressionMultiscript] class.
var (
	_AXMathExpressionMultiscriptClass     AXMathExpressionMultiscriptClass
	_AXMathExpressionMultiscriptClassOnce sync.Once
)

func getAXMathExpressionMultiscriptClass() AXMathExpressionMultiscriptClass {
	_AXMathExpressionMultiscriptClassOnce.Do(func() {
		_AXMathExpressionMultiscriptClass = AXMathExpressionMultiscriptClass{class: objc.GetClass("AXMathExpressionMultiscript")}
	})
	return _AXMathExpressionMultiscriptClass
}

// GetAXMathExpressionMultiscriptClass returns the class object for AXMathExpressionMultiscript.
func GetAXMathExpressionMultiscriptClass() AXMathExpressionMultiscriptClass {
	return getAXMathExpressionMultiscriptClass()
}

type AXMathExpressionMultiscriptClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionMultiscriptClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionMultiscriptClass) Alloc() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionMultiscript.InitWithBaseExpressionPrescriptExpressionsPostscriptExpressions]
//
// # Instance Properties
//
//   - [AXMathExpressionMultiscript.BaseExpression]
//   - [AXMathExpressionMultiscript.PostscriptExpressions]
//   - [AXMathExpressionMultiscript.PrescriptExpressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript
type AXMathExpressionMultiscript struct {
	AXMathExpression
}

// AXMathExpressionMultiscriptFromID constructs a [AXMathExpressionMultiscript] from an objc.ID.
func AXMathExpressionMultiscriptFromID(id objc.ID) AXMathExpressionMultiscript {
	return AXMathExpressionMultiscript{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionMultiscript adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionMultiscript] class.
//
// # Initializers
//
//   - [IAXMathExpressionMultiscript.InitWithBaseExpressionPrescriptExpressionsPostscriptExpressions]
//
// # Instance Properties
//
//   - [IAXMathExpressionMultiscript.BaseExpression]
//   - [IAXMathExpressionMultiscript.PostscriptExpressions]
//   - [IAXMathExpressionMultiscript.PrescriptExpressions]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript
type IAXMathExpressionMultiscript interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithBaseExpressionPrescriptExpressionsPostscriptExpressions(baseExpression IAXMathExpression, prescriptExpressions []AXMathExpressionSubSuperscript, postscriptExpressions []AXMathExpressionSubSuperscript) AXMathExpressionMultiscript

	// Topic: Instance Properties

	BaseExpression() IAXMathExpression
	PostscriptExpressions() []AXMathExpressionSubSuperscript
	PrescriptExpressions() []AXMathExpressionSubSuperscript
}

// Init initializes the instance.
func (a AXMathExpressionMultiscript) Init() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionMultiscript) Autorelease() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionMultiscript creates a new AXMathExpressionMultiscript instance.
func NewAXMathExpressionMultiscript() AXMathExpressionMultiscript {
	class := getAXMathExpressionMultiscriptClass()
	rv := objc.Send[AXMathExpressionMultiscript](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/init(baseExpression:prescriptExpressions:postscriptExpressions:)
func NewAXMathExpressionMultiscriptWithBaseExpressionPrescriptExpressionsPostscriptExpressions(baseExpression IAXMathExpression, prescriptExpressions []AXMathExpressionSubSuperscript, postscriptExpressions []AXMathExpressionSubSuperscript) AXMathExpressionMultiscript {
	instance := getAXMathExpressionMultiscriptClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBaseExpression:prescriptExpressions:postscriptExpressions:"), baseExpression, objectivec.IObjectSliceToNSArray(prescriptExpressions), objectivec.IObjectSliceToNSArray(postscriptExpressions))
	return AXMathExpressionMultiscriptFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/init(baseExpression:prescriptExpressions:postscriptExpressions:)
func (a AXMathExpressionMultiscript) InitWithBaseExpressionPrescriptExpressionsPostscriptExpressions(baseExpression IAXMathExpression, prescriptExpressions []AXMathExpressionSubSuperscript, postscriptExpressions []AXMathExpressionSubSuperscript) AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](a.ID, objc.Sel("initWithBaseExpression:prescriptExpressions:postscriptExpressions:"), baseExpression, objectivec.IObjectSliceToNSArray(prescriptExpressions), objectivec.IObjectSliceToNSArray(postscriptExpressions))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/baseExpression
func (a AXMathExpressionMultiscript) BaseExpression() IAXMathExpression {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("baseExpression"))
	return AXMathExpressionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/postscriptExpressions
func (a AXMathExpressionMultiscript) PostscriptExpressions() []AXMathExpressionSubSuperscript {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("postscriptExpressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpressionSubSuperscript {
		return AXMathExpressionSubSuperscriptFromID(id)
	})
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/prescriptExpressions
func (a AXMathExpressionMultiscript) PrescriptExpressions() []AXMathExpressionSubSuperscript {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("prescriptExpressions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AXMathExpressionSubSuperscript {
		return AXMathExpressionSubSuperscriptFromID(id)
	})
}
