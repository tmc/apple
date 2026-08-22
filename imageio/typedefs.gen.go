// Code generated from Apple documentation. DO NOT EDIT.

package imageio

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
)

// CGImageDestinationRef is an opaque type that you use to write image data to a URL, data object, or data consumer.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestination
type CGImageDestinationRef uintptr

// CGImageMetadataRef is an immutable object that contains the XMP metadata associated with an image.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadata
type CGImageMetadataRef uintptr

// CGImageMetadataTagBlock is the block to execute when enumerating the tags of a metadata object.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagBlock
type CGImageMetadataTagBlock = func(path corefoundation.CFStringRef, tag *CGImageMetadataTagRef) bool

// CGImageMetadataTagRef is an immutable type that contains information about a single piece of image metadata.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTag
type CGImageMetadataTagRef uintptr

// CGImageSourceAnimationBlock is the block to execute for each frame of an image animation.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceAnimationBlock
type CGImageSourceAnimationBlock = func(index uint32, image *coregraphics.CGImageRef, stop *bool)

// CGImageSourceRef is an opaque type that you use to read image data from a URL, data object, or data consumer.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSource
type CGImageSourceRef uintptr

// CGMutableImageMetadataRef is an opaque type for adding or modifying image metadata.
//
// See: https://developer.apple.com/documentation/ImageIO/CGMutableImageMetadata
type CGMutableImageMetadataRef uintptr
