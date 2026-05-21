// Code generated from Apple documentation for Foundation. DO NOT EDIT.
//go:build ios
// +build ios

package foundation

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
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
// parameter to the corresponding [NSCoder.DecodeCGAffineTransformForKey]
// method to retrieve the data.
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
// parameter to the corresponding [NSCoder.DecodeCGPointForKey] method to
// retrieve the data.
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
// parameter to the corresponding [NSCoder.DecodeCGRectForKey] method to
// retrieve the data.
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
// parameter to the corresponding [NSCoder.DecodeCGSizeForKey] method to
// retrieve the data.
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
// parameter to the corresponding [NSCoder.DecodeCGVectorForKey] method to
// retrieve the data.
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
// # Discussion
//
// When decoding the data from the archive, you pass the value in the key
// parameter to the corresponding [NSCoder.DecodeDirectionalEdgeInsetsForKey]
// method to retrieve the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-7oo2n
func (c NSCoder) EncodeDirectionalEdgeInsetsForKey(insets unsafe.Pointer, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeDirectionalEdgeInsets:forKey:"), insets, objc.String(key))
}

// Encodes edge inset data and associates it with the specified key in the
// coder’s archive.
//
// insets: The edge insets data to encode.
//
// key: The key identifying the data.
//
// # Discussion
//
// When decoding the data from the archive, you pass the value in the `key`
// parameter to the corresponding [NSCoder.DecodeUIEdgeInsetsForKey] method to
// retrieve the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-44zsc
func (c NSCoder) EncodeUIEdgeInsetsForKey(insets unsafe.Pointer, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeUIEdgeInsets:forKey:"), insets, objc.String(key))
}

// Encodes offset data and associates it with the specified key in the
// coder’s archive.
//
// offset: The offset data to encode.
//
// key: The key identifying the data.
//
// # Discussion
//
// When decoding the data from the archive, you pass the value in the `key`
// parameter to the corresponding [NSCoder.DecodeUIOffsetForKey] method to
// retrieve the data.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/encode(_:forKey:)-9d1qy
func (c NSCoder) EncodeUIOffsetForKey(offset unsafe.Pointer, key string) {
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
// using the [NSCoder.EncodeCMTimeMappingForKey] method.
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
// [NSCoder.EncodeCMTimeMappingForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGPoint(forKey:)
func (c NSCoder) DecodeCGPointForKey(key string) NSPoint {
	rv := objc.Send[NSPoint](c.ID, objc.Sel("decodeCGPointForKey:"), objc.String(key))
	return NSPoint(rv)
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
// [NSCoder.EncodeCMTimeMappingForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGRect(forKey:)
func (c NSCoder) DecodeCGRectForKey(key string) NSRect {
	rv := objc.Send[NSRect](c.ID, objc.Sel("decodeCGRectForKey:"), objc.String(key))
	return NSRect(rv)
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
// using the [NSCoder.EncodeCMTimeMappingForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeCGSize(forKey:)
func (c NSCoder) DecodeCGSizeForKey(key string) NSSize {
	rv := objc.Send[NSSize](c.ID, objc.Sel("decodeCGSizeForKey:"), objc.String(key))
	return NSSize(rv)
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
// using the [NSCoder.EncodeCMTimeMappingForKey] method.
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
// encoded using the [NSCoder.EncodeCMTimeMappingForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeDirectionalEdgeInsets(forKey:)
func (c NSCoder) DecodeDirectionalEdgeInsetsForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("decodeDirectionalEdgeInsetsForKey:"), objc.String(key))
	return rv
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
// encoded using the [NSCoder.EncodeCMTimeMappingForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeUIEdgeInsets(forKey:)
func (c NSCoder) DecodeUIEdgeInsetsForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("decodeUIEdgeInsetsForKey:"), objc.String(key))
	return rv
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
// using the [NSCoder.EncodeCMTimeMappingForKey] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoder/decodeUIOffset(forKey:)
func (c NSCoder) DecodeUIOffsetForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("decodeUIOffsetForKey:"), objc.String(key))
	return rv
}
