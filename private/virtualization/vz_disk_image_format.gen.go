// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDiskImageFormat] class.
var (
	_VZDiskImageFormatClass     VZDiskImageFormatClass
	_VZDiskImageFormatClassOnce sync.Once
)

func getVZDiskImageFormatClass() VZDiskImageFormatClass {
	_VZDiskImageFormatClassOnce.Do(func() {
		_VZDiskImageFormatClass = VZDiskImageFormatClass{class: objc.GetClass("_VZDiskImageFormat")}
	})
	return _VZDiskImageFormatClass
}

// GetVZDiskImageFormatClass returns the class object for _VZDiskImageFormat.
func GetVZDiskImageFormatClass() VZDiskImageFormatClass {
	return getVZDiskImageFormatClass()
}

type VZDiskImageFormatClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDiskImageFormatClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDiskImageFormatClass) Alloc() VZDiskImageFormat {
	rv := objc.Send[VZDiskImageFormat](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZDiskImageFormat.CanCreateDynamicDiskImages]
//   - [VZDiskImageFormat.CanCreateFixedDiskImages]
//   - [VZDiskImageFormat.CreateDynamicDiskImageWithURLSuggestedSizeCompletionHandler]
//   - [VZDiskImageFormat.CreateFixedDiskImageWithURLSuggestedSizeCompletionHandler]
//   - [VZDiskImageFormat.Identifier]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageFormat
type VZDiskImageFormat struct {
	objectivec.Object
}

// VZDiskImageFormatFromID constructs a [VZDiskImageFormat] from an objc.ID.
func VZDiskImageFormatFromID(id objc.ID) VZDiskImageFormat {
	return VZDiskImageFormat{objectivec.Object{ID: id}}
}

// Ensure VZDiskImageFormat implements IVZDiskImageFormat.
var _ IVZDiskImageFormat = VZDiskImageFormat{}

// An interface definition for the [VZDiskImageFormat] class.
//
// # Methods
//
//   - [IVZDiskImageFormat.CanCreateDynamicDiskImages]
//   - [IVZDiskImageFormat.CanCreateFixedDiskImages]
//   - [IVZDiskImageFormat.CreateDynamicDiskImageWithURLSuggestedSizeCompletionHandler]
//   - [IVZDiskImageFormat.CreateFixedDiskImageWithURLSuggestedSizeCompletionHandler]
//   - [IVZDiskImageFormat.Identifier]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageFormat
type IVZDiskImageFormat interface {
	objectivec.IObject

	// Topic: Methods

	CanCreateDynamicDiskImages() bool
	CanCreateFixedDiskImages() bool
	CreateDynamicDiskImageWithURLSuggestedSizeCompletionHandler(url foundation.INSURL, size uint64, handler ErrorHandler) objectivec.IObject
	CreateFixedDiskImageWithURLSuggestedSizeCompletionHandler(url foundation.INSURL, size uint64, handler ErrorHandler) objectivec.IObject
	Identifier() string
}

// Init initializes the instance.
func (v VZDiskImageFormat) Init() VZDiskImageFormat {
	rv := objc.Send[VZDiskImageFormat](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDiskImageFormat) Autorelease() VZDiskImageFormat {
	rv := objc.Send[VZDiskImageFormat](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDiskImageFormat creates a new VZDiskImageFormat instance.
func NewVZDiskImageFormat() VZDiskImageFormat {
	class := getVZDiskImageFormatClass()
	rv := objc.Send[VZDiskImageFormat](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageFormat/createDynamicDiskImageWithURL:suggestedSize:completionHandler:
func (v VZDiskImageFormat) CreateDynamicDiskImageWithURLSuggestedSizeCompletionHandler(url foundation.INSURL, size uint64, handler ErrorHandler) objectivec.IObject {
	_block2, _ := NewErrorBlock(handler)
	rv := objc.Send[objc.ID](v.ID, objc.Sel("createDynamicDiskImageWithURL:suggestedSize:completionHandler:"), url, size, _block2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageFormat/createFixedDiskImageWithURL:suggestedSize:completionHandler:
func (v VZDiskImageFormat) CreateFixedDiskImageWithURLSuggestedSizeCompletionHandler(url foundation.INSURL, size uint64, handler ErrorHandler) objectivec.IObject {
	_block2, _ := NewErrorBlock(handler)
	rv := objc.Send[objc.ID](v.ID, objc.Sel("createFixedDiskImageWithURL:suggestedSize:completionHandler:"), url, size, _block2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageFormat/rawDiskImageFormat
func (_VZDiskImageFormatClass VZDiskImageFormatClass) RawDiskImageFormat() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZDiskImageFormatClass.class), objc.Sel("rawDiskImageFormat"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageFormat/supportedFormats
func (_VZDiskImageFormatClass VZDiskImageFormatClass) SupportedFormats() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZDiskImageFormatClass.class), objc.Sel("supportedFormats"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageFormat/canCreateDynamicDiskImages
func (v VZDiskImageFormat) CanCreateDynamicDiskImages() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canCreateDynamicDiskImages"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageFormat/canCreateFixedDiskImages
func (v VZDiskImageFormat) CanCreateFixedDiskImages() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canCreateFixedDiskImages"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDiskImageFormat/identifier
func (v VZDiskImageFormat) Identifier() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// CreateDynamicDiskImageWithURLSuggestedSize is a synchronous wrapper around [VZDiskImageFormat.CreateDynamicDiskImageWithURLSuggestedSizeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZDiskImageFormat) CreateDynamicDiskImageWithURLSuggestedSize(ctx context.Context, url foundation.INSURL, size uint64) error {
	done := make(chan error, 1)
	v.CreateDynamicDiskImageWithURLSuggestedSizeCompletionHandler(url, size, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CreateFixedDiskImageWithURLSuggestedSize is a synchronous wrapper around [VZDiskImageFormat.CreateFixedDiskImageWithURLSuggestedSizeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZDiskImageFormat) CreateFixedDiskImageWithURLSuggestedSize(ctx context.Context, url foundation.INSURL, size uint64) error {
	done := make(chan error, 1)
	v.CreateFixedDiskImageWithURLSuggestedSizeCompletionHandler(url, size, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
