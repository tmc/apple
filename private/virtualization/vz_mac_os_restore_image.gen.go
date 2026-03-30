// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacOSRestoreImage] class.
var (
	_VZMacOSRestoreImageClass     VZMacOSRestoreImageClass
	_VZMacOSRestoreImageClassOnce sync.Once
)

func getVZMacOSRestoreImageClass() VZMacOSRestoreImageClass {
	_VZMacOSRestoreImageClassOnce.Do(func() {
		_VZMacOSRestoreImageClass = VZMacOSRestoreImageClass{class: objc.GetClass("VZMacOSRestoreImage")}
	})
	return _VZMacOSRestoreImageClass
}

// GetVZMacOSRestoreImageClass returns the class object for VZMacOSRestoreImage.
func GetVZMacOSRestoreImageClass() VZMacOSRestoreImageClass {
	return getVZMacOSRestoreImageClass()
}

type VZMacOSRestoreImageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacOSRestoreImageClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSRestoreImageClass) Alloc() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacOSRestoreImage._configurations]
//   - [VZMacOSRestoreImage.Supported]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage
type VZMacOSRestoreImage struct {
	objectivec.Object
}

// VZMacOSRestoreImageFromID constructs a [VZMacOSRestoreImage] from an objc.ID.
func VZMacOSRestoreImageFromID(id objc.ID) VZMacOSRestoreImage {
	return VZMacOSRestoreImage{objectivec.Object{ID: id}}
}

// Ensure VZMacOSRestoreImage implements IVZMacOSRestoreImage.
var _ IVZMacOSRestoreImage = VZMacOSRestoreImage{}

// An interface definition for the [VZMacOSRestoreImage] class.
//
// # Methods
//
//   - [IVZMacOSRestoreImage._configurations]
//   - [IVZMacOSRestoreImage.Supported]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage
type IVZMacOSRestoreImage interface {
	objectivec.IObject

	// Topic: Methods

	_configurations() foundation.INSArray
	Supported() bool
}

// Init initializes the instance.
func (m VZMacOSRestoreImage) Init() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacOSRestoreImage) Autorelease() VZMacOSRestoreImage {
	rv := objc.Send[VZMacOSRestoreImage](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSRestoreImage creates a new VZMacOSRestoreImage instance.
func NewVZMacOSRestoreImage() VZMacOSRestoreImage {
	class := getVZMacOSRestoreImageClass()
	rv := objc.Send[VZMacOSRestoreImage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/_fetchAvailableImagesWithCompletionHandler:
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) _fetchAvailableImagesWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_VZMacOSRestoreImageClass.class), objc.Sel("_fetchAvailableImagesWithCompletionHandler:"), _block0)
}

// FetchAvailableImagesWithCompletionHandler is an exported wrapper for the private method _fetchAvailableImagesWithCompletionHandler.
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) FetchAvailableImagesWithCompletionHandler(handler ErrorHandler) {
	_VZMacOSRestoreImageClass._fetchAvailableImagesWithCompletionHandler(handler)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/_fetchLatestSupportedWithOptions:completionHandler:
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) _fetchLatestSupportedWithOptionsCompletionHandler(options objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_VZMacOSRestoreImageClass.class), objc.Sel("_fetchLatestSupportedWithOptions:completionHandler:"), options, _block1)
}

// FetchLatestSupportedWithOptionsCompletionHandler is an exported wrapper for the private method _fetchLatestSupportedWithOptionsCompletionHandler.
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) FetchLatestSupportedWithOptionsCompletionHandler(options objectivec.IObject, handler ErrorHandler) {
	_VZMacOSRestoreImageClass._fetchLatestSupportedWithOptionsCompletionHandler(options, handler)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/_loadCatalogWithOptions:completionHandler:
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) _loadCatalogWithOptionsCompletionHandler(options objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_VZMacOSRestoreImageClass.class), objc.Sel("_loadCatalogWithOptions:completionHandler:"), options, _block1)
}

// LoadCatalogWithOptionsCompletionHandler is an exported wrapper for the private method _loadCatalogWithOptionsCompletionHandler.
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) LoadCatalogWithOptionsCompletionHandler(options objectivec.IObject, handler ErrorHandler) {
	_VZMacOSRestoreImageClass._loadCatalogWithOptionsCompletionHandler(options, handler)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/_loadFileURL:deviceClassParser:completionHandler:
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) _loadFileURLDeviceClassParserCompletionHandler(url foundation.INSURL, parser objectivec.IObject, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_VZMacOSRestoreImageClass.class), objc.Sel("_loadFileURL:deviceClassParser:completionHandler:"), url, parser, _block2)
}

// LoadFileURLDeviceClassParserCompletionHandler is an exported wrapper for the private method _loadFileURLDeviceClassParserCompletionHandler.
func (_VZMacOSRestoreImageClass VZMacOSRestoreImageClass) LoadFileURLDeviceClassParserCompletionHandler(url foundation.INSURL, parser objectivec.IObject, handler ErrorHandler) {
	_VZMacOSRestoreImageClass._loadFileURLDeviceClassParserCompletionHandler(url, parser, handler)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/_configurations
func (m VZMacOSRestoreImage) _configurations() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_configurations"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacOSRestoreImage/supported
func (m VZMacOSRestoreImage) Supported() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supported"))
	return rv
}

// _fetchAvailableImages is a synchronous wrapper around [VZMacOSRestoreImage._fetchAvailableImagesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc VZMacOSRestoreImageClass) _fetchAvailableImages(ctx context.Context) error {
	done := make(chan error, 1)
	mc._fetchAvailableImagesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _fetchLatestSupportedWithOptions is a synchronous wrapper around [VZMacOSRestoreImage._fetchLatestSupportedWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc VZMacOSRestoreImageClass) _fetchLatestSupportedWithOptions(ctx context.Context, options objectivec.IObject) error {
	done := make(chan error, 1)
	mc._fetchLatestSupportedWithOptionsCompletionHandler(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _loadCatalogWithOptions is a synchronous wrapper around [VZMacOSRestoreImage._loadCatalogWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc VZMacOSRestoreImageClass) _loadCatalogWithOptions(ctx context.Context, options objectivec.IObject) error {
	done := make(chan error, 1)
	mc._loadCatalogWithOptionsCompletionHandler(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _loadFileURLDeviceClassParser is a synchronous wrapper around [VZMacOSRestoreImage._loadFileURLDeviceClassParserCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc VZMacOSRestoreImageClass) _loadFileURLDeviceClassParser(ctx context.Context, url foundation.INSURL, parser objectivec.IObject) error {
	done := make(chan error, 1)
	mc._loadFileURLDeviceClassParserCompletionHandler(url, parser, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
