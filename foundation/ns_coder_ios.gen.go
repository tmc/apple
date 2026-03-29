// Code generated from Apple documentation for Foundation. DO NOT EDIT.
// +build ios

package foundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/corefoundation"
"github.com/tmc/apple/objectivec"
)

// Encodes an affine transform and associates it with the specified key in the
// receiver’s archive.
//
// transform: The transform information to encode.
//
// key: The key identifying the data.
//
// # Discussion
// 
// When decoding the data from the archive, you pass the value in the `key`
// parameter to the corresponding [DecodeCGAffineTransformForKey] method to
// retrieve the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-29jyx
func (c NSCoder) EncodeCGAffineTransformForKey(transform corefoundation.CGAffineTransform, key string) {
objc.Send[objc.ID](c.ID, objc.Sel("encodeCGAffineTransform:forKey:"), transform, objc.String(key))
}
// Encodes a point and associates it with the specified key in the
// receiver’s archive.
//
// point: The point to encode.
//
// key: The key identifying the data.
//
// # Discussion
// 
// When decoding the data from the archive, you pass the value in the `key`
// parameter to the corresponding [DecodeCGPointForKey] method to retrieve the
// data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7z9kc
func (c NSCoder) EncodeCGPointForKey(point corefoundation.CGPoint, key string) {
objc.Send[objc.ID](c.ID, objc.Sel("encodeCGPoint:forKey:"), point, objc.String(key))
}
// Encodes a rectangle and associates it with the specified key in the
// receiver’s archive.
//
// rect: The rectangle to encode.
//
// key: The key identifying the data.
//
// # Discussion
// 
// When decoding the data from the archive, you pass the value in the `key`
// parameter to the corresponding [DecodeCGRectForKey] method to retrieve the
// data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-10qhm
func (c NSCoder) EncodeCGRectForKey(rect corefoundation.CGRect, key string) {
objc.Send[objc.ID](c.ID, objc.Sel("encodeCGRect:forKey:"), rect, objc.String(key))
}
// Encodes size information and associates it with the specified key in the
// coder’s archive.
//
// size: The size information to encode.
//
// key: The key identifying the data.
//
// # Discussion
// 
// When decoding the data from the archive, you pass the value in the `key`
// parameter to the corresponding [DecodeCGSizeForKey] method to retrieve the
// data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-6wq3n
func (c NSCoder) EncodeCGSizeForKey(size corefoundation.CGSize, key string) {
objc.Send[objc.ID](c.ID, objc.Sel("encodeCGSize:forKey:"), size, objc.String(key))
}
// Encodes vector data and associates it with the specified key in the
// coder’s archive.
//
// vector: The vector data to encode.
//
// key: The key identifying the data.
//
// # Discussion
// 
// When decoding the data from the archive, you pass the value in the `key`
// parameter to the corresponding [DecodeCGVectorForKey] method to retrieve
// the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-26fxa
func (c NSCoder) EncodeCGVectorForKey(vector corefoundation.CGVector, key string) {
objc.Send[objc.ID](c.ID, objc.Sel("encodeCGVector:forKey:"), vector, objc.String(key))
}
// Encodes directional edge inset data and associates it with the specified
// key in the coder’s archive.
//
// insets: The edge insets data to encode.
//
// key: The key identifying the data.
//
// insets is a [appkit.NSDirectionalEdgeInsets].
//
// # Discussion
// 
// When decoding the data from the archive, you pass the value in the key
// parameter to the corresponding [DecodeDirectionalEdgeInsetsForKey] method
// to retrieve the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7oo2n
// insets is a [appkit.NSDirectionalEdgeInsets].
func (c NSCoder) EncodeDirectionalEdgeInsetsForKey(insets objectivec.IObject, key string) {
objc.Send[objc.ID](c.ID, objc.Sel("encodeDirectionalEdgeInsets:forKey:"), insets, objc.String(key))
}
// Encodes edge inset data and associates it with the specified key in the
// coder’s archive.
//
// insets: The edge insets data to encode.
//
// key: The key identifying the data.
//
// insets is a [uikit.UIEdgeInsets].
//
// # Discussion
// 
// When decoding the data from the archive, you pass the value in the `key`
// parameter to the corresponding [DecodeUIEdgeInsetsForKey] method to
// retrieve the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-44zsc
// insets is a [uikit.UIEdgeInsets].
func (c NSCoder) EncodeUIEdgeInsetsForKey(insets objectivec.IObject, key string) {
objc.Send[objc.ID](c.ID, objc.Sel("encodeUIEdgeInsets:forKey:"), insets, objc.String(key))
}
// Encodes offset data and associates it with the specified key in the
// coder’s archive.
//
// offset: The offset data to encode.
//
// key: The key identifying the data.
//
// offset is a [uikit.UIOffset].
//
// # Discussion
// 
// When decoding the data from the archive, you pass the value in the `key`
// parameter to the corresponding [DecodeUIOffsetForKey] method to retrieve
// the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9d1qy
// offset is a [uikit.UIOffset].
func (c NSCoder) EncodeUIOffsetForKey(offset objectivec.IObject, key string) {
objc.Send[objc.ID](c.ID, objc.Sel("encodeUIOffset:forKey:"), offset, objc.String(key))
}
// Decodes and returns the Core Graphics affine transform structure associated
// with the specified key in the coder’s archive.
//
// key: The key that identifies the affine transform.
//
// # Return Value
// 
// The affine transform.
//
// # Discussion
// 
// Use this method to decode size information that was previously encoded
// using the [EncodeCGAffineTransformForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGAffineTransform(forKey:)
func (c NSCoder) DecodeCGAffineTransformForKey(key string) corefoundation.CGAffineTransform {
rv := objc.Send[corefoundation.CGAffineTransform](c.ID, objc.Sel("decodeCGAffineTransformForKey:"), objc.String(key))
return corefoundation.CGAffineTransform(rv)
}
// Decodes and returns the Core Graphics point structure associated with the
// specified key in the coder’s archive.
//
// key: The key that identifies the point.
//
// # Return Value
// 
// The [CGPoint] structure.
//
// # Discussion
// 
// Use this method to decode a point that was previously encoded using the
// [EncodeCGPointForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGPoint(forKey:)
func (c NSCoder) DecodeCGPointForKey(key string) corefoundation.CGPoint {
rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("decodeCGPointForKey:"), objc.String(key))
return corefoundation.CGPoint(rv)
}
// Decodes and returns the Core Graphics rectangle structure associated with
// the specified key in the coder’s archive.
//
// key: The key that identifies the rectangle.
//
// # Return Value
// 
// The [CGRect] structure.
//
// # Discussion
// 
// Use this method to decode a rectangle that was previously encoded using the
// [EncodeCGRectForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGRect(forKey:)
func (c NSCoder) DecodeCGRectForKey(key string) corefoundation.CGRect {
rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("decodeCGRectForKey:"), objc.String(key))
return corefoundation.CGRect(rv)
}
// Decodes and returns the Core Graphics size structure associated with the
// specified key in the coder’s archive.
//
// key: The key that identifies the size information.
//
// # Return Value
// 
// The [CGSize] structure.
//
// # Discussion
// 
// Use this method to decode size information that was previously encoded
// using the [EncodeCGSizeForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGSize(forKey:)
func (c NSCoder) DecodeCGSizeForKey(key string) corefoundation.CGSize {
rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("decodeCGSizeForKey:"), objc.String(key))
return corefoundation.CGSize(rv)
}
// Decodes and returns the Core Graphics vector data associated with the
// specified key in the coder’s archive.
//
// key: The key that identifies the vector.
//
// # Return Value
// 
// The vector data.
//
// # Discussion
// 
// Use this method to decode vector information that was previously encoded
// using the [EncodeCGVectorForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGVector(forKey:)
func (c NSCoder) DecodeCGVectorForKey(key string) corefoundation.CGVector {
rv := objc.Send[corefoundation.CGVector](c.ID, objc.Sel("decodeCGVectorForKey:"), objc.String(key))
return corefoundation.CGVector(rv)
}
// Decodes and returns the UIKit directional edge insets structure associated
// with the specified key in the coder’s archive.
//
// key: The key that identifies the edge insets.
//
// # Discussion
// 
// Use this method to decode edge inset information that was previously
// encoded using the [EncodeDirectionalEdgeInsetsForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDirectionalEdgeInsets(forKey:)
func (c NSCoder) DecodeDirectionalEdgeInsetsForKey(key string) objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeDirectionalEdgeInsetsForKey:"), objc.String(key))
return objectivec.Object{ID: rv}
}
// Decodes and returns the UIKit edge insets structure associated with the
// specified key in the coder’s archive.
//
// key: The key that identifies the edge insets.
//
// # Return Value
// 
// The edge insets data.
//
// # Discussion
// 
// Use this method to decode edge inset information that was previously
// encoded using the [EncodeUIEdgeInsetsForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeUIEdgeInsets(forKey:)
func (c NSCoder) DecodeUIEdgeInsetsForKey(key string) objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeUIEdgeInsetsForKey:"), objc.String(key))
return objectivec.Object{ID: rv}
}
// Decodes and returns the UIKit offset structure associated with the
// specified key in the coder’s archive.
//
// key: The key that identifies the offset.
//
// # Return Value
// 
// The offset data.
//
// # Discussion
// 
// Use this method to decode offset information that was previously encoded
// using the [EncodeUIOffsetForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeUIOffset(forKey:)
func (c NSCoder) DecodeUIOffsetForKey(key string) objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("decodeUIOffsetForKey:"), objc.String(key))
return objectivec.Object{ID: rv}
}

