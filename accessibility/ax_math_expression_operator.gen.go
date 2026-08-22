// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AXMathExpressionOperator] class.
var (
	_AXMathExpressionOperatorClass     AXMathExpressionOperatorClass
	_AXMathExpressionOperatorClassOnce sync.Once
)

func getAXMathExpressionOperatorClass() AXMathExpressionOperatorClass {
	_AXMathExpressionOperatorClassOnce.Do(func() {
		_AXMathExpressionOperatorClass = AXMathExpressionOperatorClass{class: objc.GetClass("AXMathExpressionOperator")}
	})
	return _AXMathExpressionOperatorClass
}

// GetAXMathExpressionOperatorClass returns the class object for AXMathExpressionOperator.
func GetAXMathExpressionOperatorClass() AXMathExpressionOperatorClass {
	return getAXMathExpressionOperatorClass()
}

type AXMathExpressionOperatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionOperatorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionOperatorClass) Alloc() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionOperator.InitWithContent]
//
// # Instance Properties
//
//   - [AXMathExpressionOperator.Content]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator
type AXMathExpressionOperator struct {
	AXMathExpression
}

// AXMathExpressionOperatorFromID constructs a [AXMathExpressionOperator] from an objc.ID.
func AXMathExpressionOperatorFromID(id objc.ID) AXMathExpressionOperator {
	return AXMathExpressionOperator{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionOperator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionOperator] class.
//
// # Initializers
//
//   - [IAXMathExpressionOperator.InitWithContent]
//
// # Instance Properties
//
//   - [IAXMathExpressionOperator.Content]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator
type IAXMathExpressionOperator interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithContent(content string) AXMathExpressionOperator

	// Topic: Instance Properties

	Content() string
}

// Init initializes the instance.
func (a AXMathExpressionOperator) Init() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionOperator) Autorelease() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionOperator creates a new AXMathExpressionOperator instance.
func NewAXMathExpressionOperator() AXMathExpressionOperator {
	class := getAXMathExpressionOperatorClass()
	rv := objc.Send[AXMathExpressionOperator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator/init(content:)
func NewAXMathExpressionOperatorWithContent(content string) AXMathExpressionOperator {
	instance := getAXMathExpressionOperatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContent:"), objc.String(content))
	return AXMathExpressionOperatorFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator/init(content:)
func (a AXMathExpressionOperator) InitWithContent(content string) AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](a.ID, objc.Sel("initWithContent:"), objc.String(content))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator/content
func (a AXMathExpressionOperator) Content() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("content"))
	return foundation.NSStringFromID(rv).String()
}
