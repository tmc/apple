// Code generated from Apple documentation. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
)

// CIContextOption is an enum string type that your code can use to select different options when creating a Core Image context.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContextOption
type CIContextOption = string

// CIDynamicRangeOption is an enum string type that your code can use to select different System Tone Mapping modes.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDynamicRangeOption
type CIDynamicRangeOption = string

// CIFormat is pixel data formats for image input, output, and processing.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFormat
type CIFormat = int

// See: https://developer.apple.com/documentation/CoreImage/CIImageAutoAdjustmentOption
type CIImageAutoAdjustmentOption = string

// See: https://developer.apple.com/documentation/CoreImage/CIImageOption
type CIImageOption = string

// See: https://developer.apple.com/documentation/CoreImage/CIImageRepresentationOption
type CIImageRepresentationOption = string

// CIKernelROICallback is the signature for a block that computes the region of interest (ROI) for a given area of destination image pixels. Core Image calls this block when applying the kernel. You specify this block when using the [apply(extent:roiCallback:arguments:)] method.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKernelROICallback
type CIKernelROICallback = func(int, corefoundation.CGRect) corefoundation.CGRect

// See: https://developer.apple.com/documentation/CoreImage/CIRAWDecoderVersion
type CIRAWDecoderVersion = string

