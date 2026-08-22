// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AXMathExpressionText] class.
var (
	_AXMathExpressionTextClass     AXMathExpressionTextClass
	_AXMathExpressionTextClassOnce sync.Once
)

func getAXMathExpressionTextClass() AXMathExpressionTextClass {
	_AXMathExpressionTextClassOnce.Do(func() {
		_AXMathExpressionTextClass = AXMathExpressionTextClass{class: objc.GetClass("AXMathExpressionText")}
	})
	return _AXMathExpressionTextClass
}

// GetAXMathExpressionTextClass returns the class object for AXMathExpressionText.
func GetAXMathExpressionTextClass() AXMathExpressionTextClass {
	return getAXMathExpressionTextClass()
}

type AXMathExpressionTextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXMathExpressionTextClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXMathExpressionTextClass) Alloc() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [AXMathExpressionText.InitWithContent]
//
// # Instance Properties
//
//   - [AXMathExpressionText.Content]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText
type AXMathExpressionText struct {
	AXMathExpression
}

// AXMathExpressionTextFromID constructs a [AXMathExpressionText] from an objc.ID.
func AXMathExpressionTextFromID(id objc.ID) AXMathExpressionText {
	return AXMathExpressionText{AXMathExpression: AXMathExpressionFromID(id)}
}

// NOTE: AXMathExpressionText adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXMathExpressionText] class.
//
// # Initializers
//
//   - [IAXMathExpressionText.InitWithContent]
//
// # Instance Properties
//
//   - [IAXMathExpressionText.Content]
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText
type IAXMathExpressionText interface {
	IAXMathExpression

	// Topic: Initializers

	InitWithContent(content string) AXMathExpressionText

	// Topic: Instance Properties

	Content() string
}

// Init initializes the instance.
func (a AXMathExpressionText) Init() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXMathExpressionText) Autorelease() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionText creates a new AXMathExpressionText instance.
func NewAXMathExpressionText() AXMathExpressionText {
	class := getAXMathExpressionTextClass()
	rv := objc.Send[AXMathExpressionText](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText/init(content:)
func NewAXMathExpressionTextWithContent(content string) AXMathExpressionText {
	instance := getAXMathExpressionTextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContent:"), objc.String(content))
	return AXMathExpressionTextFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText/init(content:)
func (a AXMathExpressionText) InitWithContent(content string) AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](a.ID, objc.Sel("initWithContent:"), objc.String(content))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText/content
func (a AXMathExpressionText) Content() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("content"))
	return foundation.NSStringFromID(rv).String()
}
