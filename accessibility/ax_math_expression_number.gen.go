// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AXMathExpressionNumber] class.
var (
	_AXMathExpressionNumberClass     AXMathExpressionNumberClass
	_AXMathExpressionNumberClassOnce sync.Once
)

func getAXMathExpressionNumberClass() AXMathExpressionNumberClass {
	_AXMathExpressionNumberClassOnce.Do(func() {
		_AXMathExpressionNumberClass = AXMathExpressionNumberClass{class: objc.GetClass("AXMathExpressionNumber")}
	})
	return _AXMathExpressionNumberClass
}

// GetAXMathExpressionNumberClass returns the class object for AXMathExpressionNumber.
func GetAXMathExpressionNumberClass() AXMathExpressionNumberClass {
	return getAXMathExpressionNumberClass()
}

type AXMathExpressionNumberClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionNumberClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionNumberClass) Alloc() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionNumber.InitWithContent]
//
// # Instance Properties
//
//   - [AXMathExpressionNumber.Content]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber
type AXMathExpressionNumber struct {
	AXMathExpression
}

// AXMathExpressionNumberFromID constructs a [AXMathExpressionNumber] from an objc.ID.
func AXMathExpressionNumberFromID(id objc.ID) AXMathExpressionNumber {
	return AXMathExpressionNumber{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionNumber adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionNumber] class.
//
// # Initializers
//
//   - [IAXMathExpressionNumber.InitWithContent]
//
// # Instance Properties
//
//   - [IAXMathExpressionNumber.Content]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber
type IAXMathExpressionNumber interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithContent(content string) AXMathExpressionNumber

	// Topic: Instance Properties

	Content() string
}

// Init initializes the instance.
func (a AXMathExpressionNumber) Init() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionNumber) Autorelease() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionNumber creates a new AXMathExpressionNumber instance.
func NewAXMathExpressionNumber() AXMathExpressionNumber {
	class := getAXMathExpressionNumberClass()
	rv := objc.Send[AXMathExpressionNumber](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber/init(content:)
func NewAXMathExpressionNumberWithContent(content string) AXMathExpressionNumber {
	instance := getAXMathExpressionNumberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContent:"), objc.String(content))
	return AXMathExpressionNumberFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber/init(content:)
func (a AXMathExpressionNumber) InitWithContent(content string) AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](a.ID, objc.Sel("initWithContent:"), objc.String(content))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber/content
func (a AXMathExpressionNumber) Content() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("content"))
	return foundation.NSStringFromID(rv).String()
}
