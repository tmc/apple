// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AXMathExpressionProvider protocol.
//
// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionProvider
type AXMathExpressionProvider interface {
	objectivec.IObject

	// AccessibilityMathExpression protocol.
	//
	// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionProvider/accessibilityMathExpression()
	AccessibilityMathExpression() IAXMathExpression
}

// AXMathExpressionProviderObject wraps an existing Objective-C object that conforms to the AXMathExpressionProvider protocol.
type AXMathExpressionProviderObject struct {
	objectivec.Object
}

func (o AXMathExpressionProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// AXMathExpressionProviderObjectFromID constructs a [AXMathExpressionProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AXMathExpressionProviderObjectFromID(id objc.ID) AXMathExpressionProviderObject {
	return AXMathExpressionProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Accessibility/AXMathExpressionProvider/accessibilityMathExpression()
func (o AXMathExpressionProviderObject) AccessibilityMathExpression() IAXMathExpression {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMathExpression"))
	return AXMathExpressionFromID(rv)
}
