// Code generated from Apple documentation for Foundation. DO NOT EDIT.
// +build ios

package foundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/corefoundation"
"github.com/tmc/apple/objectivec"
)

// Creates a new value object containing the specified CoreGraphics point
// structure.
//
// point: The value for the new object.
//
// # Return Value
// 
// A new value object that contains the point information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGPoint:)
func (_NSValueClass NSValueClass) ValueWithCGPoint(point corefoundation.CGPoint) NSValue {
rv := objc.Send[objc.ID](objc.ID(_NSValueClass.class), objc.Sel("valueWithCGPoint:"), point)
return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreGraphics vector
// structure.
//
// vector: The value for the new object.
//
// # Return Value
// 
// A new value object that contains the vector information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGVector:)
func (_NSValueClass NSValueClass) ValueWithCGVector(vector corefoundation.CGVector) NSValue {
rv := objc.Send[objc.ID](objc.ID(_NSValueClass.class), objc.Sel("valueWithCGVector:"), vector)
return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreGraphics size
// structure.
//
// size: The value for the new object.
//
// # Return Value
// 
// A new value object that contains the size information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGSize:)
func (_NSValueClass NSValueClass) ValueWithCGSize(size corefoundation.CGSize) NSValue {
rv := objc.Send[objc.ID](objc.ID(_NSValueClass.class), objc.Sel("valueWithCGSize:"), size)
return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreGraphics rectangle
// structure.
//
// rect: The value for the new object.
//
// # Return Value
// 
// A new value object that contains the rectangle information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGRect:)
func (_NSValueClass NSValueClass) ValueWithCGRect(rect corefoundation.CGRect) NSValue {
rv := objc.Send[objc.ID](objc.ID(_NSValueClass.class), objc.Sel("valueWithCGRect:"), rect)
return NSValueFromID(rv)
}

// Creates a new value object containing the specified CoreGraphics affine
// transform structure.
//
// transform: The value for the new object.
//
// # Return Value
// 
// A new value object that contains the affine transform information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(CGAffineTransform:)
func (_NSValueClass NSValueClass) ValueWithCGAffineTransform(transform corefoundation.CGAffineTransform) NSValue {
rv := objc.Send[objc.ID](objc.ID(_NSValueClass.class), objc.Sel("valueWithCGAffineTransform:"), transform)
return NSValueFromID(rv)
}

// Creates a new value object containing the specified UIKit edge insets
// structure.
//
// insets: The value for the new object.
//
// insets is a [uikit.UIEdgeInsets].
//
// # Return Value
// 
// A new value object that contains the edge inset information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(UIEdgeInsets:)
func (_NSValueClass NSValueClass) ValueWithUIEdgeInsets(insets objectivec.IObject) NSValue {
rv := objc.Send[objc.ID](objc.ID(_NSValueClass.class), objc.Sel("valueWithUIEdgeInsets:"), insets)
return NSValueFromID(rv)
}

// Creates a new value object containing the specified UIKit offset structure.
//
// insets: The value for the new object.
//
// insets is a [uikit.UIOffset].
//
// # Return Value
// 
// A new value object that contains the offset information.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(UIOffset:)
func (_NSValueClass NSValueClass) ValueWithUIOffset(insets objectivec.IObject) NSValue {
rv := objc.Send[objc.ID](objc.ID(_NSValueClass.class), objc.Sel("valueWithUIOffset:"), insets)
return NSValueFromID(rv)
}

//
// insets is a [appkit.NSDirectionalEdgeInsets].
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/init(directionalEdgeInsets:)
func (_NSValueClass NSValueClass) ValueWithDirectionalEdgeInsets(insets objectivec.IObject) NSValue {
rv := objc.Send[objc.ID](objc.ID(_NSValueClass.class), objc.Sel("valueWithDirectionalEdgeInsets:"), insets)
return NSValueFromID(rv)
}






// Returns the CoreGraphics point structure representation of the value.
//
// # Return Value
// 
// The CoreGraphics point structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/cgPointValue
func (v NSValue) CGPointValue() corefoundation.CGPoint {
rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("CGPointValue"))
		return corefoundation.CGPoint(rv)
}



// Returns the CoreGraphics vector structure representation of the value.
//
// # Return Value
// 
// The CoreGraphics vector structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/cgVectorValue
func (v NSValue) CGVectorValue() corefoundation.CGVector {
rv := objc.Send[corefoundation.CGVector](v.ID, objc.Sel("CGVectorValue"))
		return corefoundation.CGVector(rv)
}



// Returns the CoreGraphics size structure representation of the value.
//
// # Return Value
// 
// The CoreGraphics size structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/cgSizeValue
func (v NSValue) CGSizeValue() corefoundation.CGSize {
rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("CGSizeValue"))
		return corefoundation.CGSize(rv)
}



// Returns the CoreGraphics rectangle structure representation of the value.
//
// # Return Value
// 
// The CoreGraphics rectangle structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/cgRectValue
func (v NSValue) CGRectValue() corefoundation.CGRect {
rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("CGRectValue"))
		return corefoundation.CGRect(rv)
}



// Returns the CoreGraphics affine transform representation of the value.
//
// # Return Value
// 
// The CoreGraphics affine transform representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/cgAffineTransformValue
func (v NSValue) CGAffineTransformValue() corefoundation.CGAffineTransform {
rv := objc.Send[corefoundation.CGAffineTransform](v.ID, objc.Sel("CGAffineTransformValue"))
		return corefoundation.CGAffineTransform(rv)
}



// Returns the UIKit edge insets structure representation of the value.
//
// # Return Value
// 
// The UIKit edge insets structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/uiEdgeInsetsValue
func (v NSValue) UIEdgeInsetsValue() objectivec.IObject {
rv := objc.Send[objc.ID](v.ID, objc.Sel("UIEdgeInsetsValue"))
return objectivec.Object{ID: rv}
}



// Returns the UIKit offset structure representation of the value.
//
// # Return Value
// 
// The UIKit offset structure representation of the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSValue/uiOffsetValue
func (v NSValue) UIOffsetValue() objectivec.IObject {
rv := objc.Send[objc.ID](v.ID, objc.Sel("UIOffsetValue"))
return objectivec.Object{ID: rv}
}



// See: https://developer.apple.com/documentation/Foundation/NSValue/directionalEdgeInsetsValue
func (v NSValue) DirectionalEdgeInsetsValue() objectivec.IObject {
rv := objc.Send[objc.ID](v.ID, objc.Sel("directionalEdgeInsetsValue"))
return objectivec.Object{ID: rv}
}








