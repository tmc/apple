// Code generated from Apple documentation for CoreImage. DO NOT EDIT.
//go:build ios
// +build ios

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
)

// Returns the maximum size allowed for any image rendered into the context.
//
// # Discussion
//
// Some contexts limit the maximum size of an image that can be rendered into
// them. For example, the maximum size might reflect a limitation in the
// underlying graphics hardware.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/inputImageMaximumSize()
func (c CIContext) InputImageMaximumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("inputImageMaximumSize"))
	return corefoundation.CGSize(rv)
}

// Returns the maximum size allowed for any image created by the context.
//
// # Discussion
//
// Some contexts limit the maximum size of an image that can be created by
// them. For example, the maximum size might reflect a limitation in the
// underlying graphics hardware.
//
// See: https://developer.apple.com/documentation/CoreImage/CIContext/outputImageMaximumSize()
func (c CIContext) OutputImageMaximumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("outputImageMaximumSize"))
	return corefoundation.CGSize(rv)
}
