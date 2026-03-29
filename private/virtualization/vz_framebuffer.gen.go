// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZFramebuffer] class.
var (
	_VZFramebufferClass     VZFramebufferClass
	_VZFramebufferClassOnce sync.Once
)

func getVZFramebufferClass() VZFramebufferClass {
	_VZFramebufferClassOnce.Do(func() {
		_VZFramebufferClass = VZFramebufferClass{class: objc.GetClass("_VZFramebuffer")}
	})
	return _VZFramebufferClass
}

// GetVZFramebufferClass returns the class object for _VZFramebuffer.
func GetVZFramebufferClass() VZFramebufferClass {
	return getVZFramebufferClass()
}

type VZFramebufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZFramebufferClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZFramebufferClass) Alloc() VZFramebuffer {
	rv := objc.Send[VZFramebuffer](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZFramebuffer._takeScreenshotWithCompletionHandlerImageConversionBlock]
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebuffer
type VZFramebuffer struct {
	objectivec.Object
}

// VZFramebufferFromID constructs a [VZFramebuffer] from an objc.ID.
func VZFramebufferFromID(id objc.ID) VZFramebuffer {
	return VZFramebuffer{objectivec.Object{ID: id}}
}
// Ensure VZFramebuffer implements IVZFramebuffer.
var _ IVZFramebuffer = VZFramebuffer{}

// An interface definition for the [VZFramebuffer] class.
//
// # Methods
//
//   - [IVZFramebuffer._takeScreenshotWithCompletionHandlerImageConversionBlock]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebuffer
type IVZFramebuffer interface {
	objectivec.IObject

	// Topic: Methods

	_takeScreenshotWithCompletionHandlerImageConversionBlock(handler VoidHandler, block VoidHandler)
}

// Init initializes the instance.
func (v VZFramebuffer) Init() VZFramebuffer {
	rv := objc.Send[VZFramebuffer](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZFramebuffer) Autorelease() VZFramebuffer {
	rv := objc.Send[VZFramebuffer](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFramebuffer creates a new VZFramebuffer instance.
func NewVZFramebuffer() VZFramebuffer {
	class := getVZFramebufferClass()
	rv := objc.Send[VZFramebuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebuffer/_takeScreenshotWithCompletionHandler:imageConversionBlock:
func (v VZFramebuffer) _takeScreenshotWithCompletionHandlerImageConversionBlock(handler VoidHandler, block VoidHandler) {
_block0, _ := NewVoidBlock(handler)
	_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("_takeScreenshotWithCompletionHandler:imageConversionBlock:"), _block0, _block1)
}

// TakeScreenshotWithCompletionHandlerImageConversionBlock is an exported wrapper for the private method _takeScreenshotWithCompletionHandlerImageConversionBlock.
func (v VZFramebuffer) TakeScreenshotWithCompletionHandlerImageConversionBlock(handler VoidHandler, block VoidHandler) {
	v._takeScreenshotWithCompletionHandlerImageConversionBlock(handler, block)
}

// _takeScreenshotWithCompletionHandlerImageConversionBlockSync is a synchronous wrapper around [VZFramebuffer._takeScreenshotWithCompletionHandlerImageConversionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZFramebuffer) _takeScreenshotWithCompletionHandlerImageConversionBlockSync(ctx context.Context, handler VoidHandler) error {
	done := make(chan struct{}, 1)
	v._takeScreenshotWithCompletionHandlerImageConversionBlock(handler, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

