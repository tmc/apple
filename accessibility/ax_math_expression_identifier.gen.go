// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AXMathExpressionIdentifier] class.
var (
	_AXMathExpressionIdentifierClass     AXMathExpressionIdentifierClass
	_AXMathExpressionIdentifierClassOnce sync.Once
)

func getAXMathExpressionIdentifierClass() AXMathExpressionIdentifierClass {
	_AXMathExpressionIdentifierClassOnce.Do(func() {
		_AXMathExpressionIdentifierClass = AXMathExpressionIdentifierClass{class: objc.GetClass("AXMathExpressionIdentifier")}
	})
	return _AXMathExpressionIdentifierClass
}

// GetAXMathExpressionIdentifierClass returns the class object for AXMathExpressionIdentifier.
func GetAXMathExpressionIdentifierClass() AXMathExpressionIdentifierClass {
	return getAXMathExpressionIdentifierClass()
}

type AXMathExpressionIdentifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionIdentifierClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionIdentifierClass) Alloc() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionIdentifier.InitWithContent]
//
// # Instance Properties
//
//   - [AXMathExpressionIdentifier.Content]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier
type AXMathExpressionIdentifier struct {
	AXMathExpression
}

// AXMathExpressionIdentifierFromID constructs a [AXMathExpressionIdentifier] from an objc.ID.
func AXMathExpressionIdentifierFromID(id objc.ID) AXMathExpressionIdentifier {
	return AXMathExpressionIdentifier{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionIdentifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionIdentifier] class.
//
// # Initializers
//
//   - [IAXMathExpressionIdentifier.InitWithContent]
//
// # Instance Properties
//
//   - [IAXMathExpressionIdentifier.Content]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier
type IAXMathExpressionIdentifier interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithContent(content string) AXMathExpressionIdentifier

	// Topic: Instance Properties

	Content() string
}

// Init initializes the instance.
func (a AXMathExpressionIdentifier) Init() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionIdentifier) Autorelease() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionIdentifier creates a new AXMathExpressionIdentifier instance.
func NewAXMathExpressionIdentifier() AXMathExpressionIdentifier {
	class := getAXMathExpressionIdentifierClass()
	rv := objc.Send[AXMathExpressionIdentifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier/init(content:)
func NewAXMathExpressionIdentifierWithContent(content string) AXMathExpressionIdentifier {
	instance := getAXMathExpressionIdentifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContent:"), objc.String(content))
	return AXMathExpressionIdentifierFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier/init(content:)
func (a AXMathExpressionIdentifier) InitWithContent(content string) AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](a.ID, objc.Sel("initWithContent:"), objc.String(content))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier/content
func (a AXMathExpressionIdentifier) Content() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("content"))
	return foundation.NSStringFromID(rv).String()
}
