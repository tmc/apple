// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XRGPUAGXShaderTimelineSignposts] class.
var (
	_XRGPUAGXShaderTimelineSignpostsClass     XRGPUAGXShaderTimelineSignpostsClass
	_XRGPUAGXShaderTimelineSignpostsClassOnce sync.Once
)

func getXRGPUAGXShaderTimelineSignpostsClass() XRGPUAGXShaderTimelineSignpostsClass {
	_XRGPUAGXShaderTimelineSignpostsClassOnce.Do(func() {
		_XRGPUAGXShaderTimelineSignpostsClass = XRGPUAGXShaderTimelineSignpostsClass{class: objc.GetClass("XRGPUAGXShaderTimelineSignposts")}
	})
	return _XRGPUAGXShaderTimelineSignpostsClass
}

// GetXRGPUAGXShaderTimelineSignpostsClass returns the class object for XRGPUAGXShaderTimelineSignposts.
func GetXRGPUAGXShaderTimelineSignpostsClass() XRGPUAGXShaderTimelineSignpostsClass {
	return getXRGPUAGXShaderTimelineSignpostsClass()
}

type XRGPUAGXShaderTimelineSignpostsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (xc XRGPUAGXShaderTimelineSignpostsClass) Class() objc.Class {
	return xc.class
}

// Alloc allocates memory for a new instance of the class.
func (xc XRGPUAGXShaderTimelineSignpostsClass) Alloc() XRGPUAGXShaderTimelineSignposts {
	rv := objc.SendIfResponds[XRGPUAGXShaderTimelineSignposts](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [XRGPUAGXShaderTimelineSignposts._setupExtractor]
//   - [XRGPUAGXShaderTimelineSignposts.Encode]
//   - [XRGPUAGXShaderTimelineSignposts.EncodeWithCoder]
//   - [XRGPUAGXShaderTimelineSignposts.EnumerateKickIds]
//   - [XRGPUAGXShaderTimelineSignposts.GetShadersCount]
//   - [XRGPUAGXShaderTimelineSignposts.Load]
//   - [XRGPUAGXShaderTimelineSignposts.Start]
//   - [XRGPUAGXShaderTimelineSignposts.Stop]
//   - [XRGPUAGXShaderTimelineSignposts.InitWithCoder]
type XRGPUAGXShaderTimelineSignposts struct {
	objectivec.Object
}

// XRGPUAGXShaderTimelineSignpostsFromID constructs a [XRGPUAGXShaderTimelineSignposts] from an objc.ID.
func XRGPUAGXShaderTimelineSignpostsFromID(id objc.ID) XRGPUAGXShaderTimelineSignposts {
	return XRGPUAGXShaderTimelineSignposts{objectivec.Object{ID: id}}
}

// Ensure XRGPUAGXShaderTimelineSignposts implements IXRGPUAGXShaderTimelineSignposts.
var _ IXRGPUAGXShaderTimelineSignposts = XRGPUAGXShaderTimelineSignposts{}

// An interface definition for the [XRGPUAGXShaderTimelineSignposts] class.
//
// # Methods
//
//   - [IXRGPUAGXShaderTimelineSignposts._setupExtractor]
//   - [IXRGPUAGXShaderTimelineSignposts.Encode]
//   - [IXRGPUAGXShaderTimelineSignposts.EncodeWithCoder]
//   - [IXRGPUAGXShaderTimelineSignposts.EnumerateKickIds]
//   - [IXRGPUAGXShaderTimelineSignposts.GetShadersCount]
//   - [IXRGPUAGXShaderTimelineSignposts.Load]
//   - [IXRGPUAGXShaderTimelineSignposts.Start]
//   - [IXRGPUAGXShaderTimelineSignposts.Stop]
//   - [IXRGPUAGXShaderTimelineSignposts.InitWithCoder]
type IXRGPUAGXShaderTimelineSignposts interface {
	objectivec.IObject

	// Topic: Methods

	_setupExtractor()
	Encode() objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	EnumerateKickIds(ids VoidHandler)
	GetShadersCount(shaders []XRGPUAGXShaderTimelineSignpostProcessRef, count *uint64)
	Load(load objectivec.IObject) bool
	Start() bool
	Stop()
	InitWithCoder(coder foundation.INSCoder) XRGPUAGXShaderTimelineSignposts
}

// Init initializes the instance.
func (x XRGPUAGXShaderTimelineSignposts) Init() XRGPUAGXShaderTimelineSignposts {
	rv := objc.SendIfResponds[XRGPUAGXShaderTimelineSignposts](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XRGPUAGXShaderTimelineSignposts) Autorelease() XRGPUAGXShaderTimelineSignposts {
	rv := objc.SendIfResponds[XRGPUAGXShaderTimelineSignposts](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXRGPUAGXShaderTimelineSignposts creates a new XRGPUAGXShaderTimelineSignposts instance.
func NewXRGPUAGXShaderTimelineSignposts() XRGPUAGXShaderTimelineSignposts {
	class := getXRGPUAGXShaderTimelineSignpostsClass()
	rv := objc.SendIfResponds[XRGPUAGXShaderTimelineSignposts](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewXRGPUAGXShaderTimelineSignpostsWithCoder(coder objectivec.IObject) XRGPUAGXShaderTimelineSignposts {
	instance := getXRGPUAGXShaderTimelineSignpostsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return XRGPUAGXShaderTimelineSignpostsFromID(rv)
}

func (x XRGPUAGXShaderTimelineSignposts) _setupExtractor() {
	objc.SendIfResponds[objc.ID](x.ID, objc.Sel("_setupExtractor"))
}

// SetupExtractor is an exported wrapper for the private method _setupExtractor.
func (x XRGPUAGXShaderTimelineSignposts) SetupExtractor() error {
	if !objc.RespondsToSelector(x.ID, objc.Sel("_setupExtractor")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setupExtractor"}
		return err
	}
	x._setupExtractor()
	return nil
}

// CanSetupExtractor reports whether the receiver responds to the private selector _setupExtractor.
func (x XRGPUAGXShaderTimelineSignposts) CanSetupExtractor() bool {
	return objc.RespondsToSelector(x.ID, objc.Sel("_setupExtractor"))
}
func (x XRGPUAGXShaderTimelineSignposts) Encode() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](x.ID, objc.Sel("encode"))
	return objectivec.Object{ID: rv}
}
func (x XRGPUAGXShaderTimelineSignposts) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](x.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (x XRGPUAGXShaderTimelineSignposts) EnumerateKickIds(ids VoidHandler) {
	_block0, _ := NewVoidBlock(ids)
	objc.SendIfResponds[objc.ID](x.ID, objc.Sel("enumerateKickIds:"), _block0)
}
func (x XRGPUAGXShaderTimelineSignposts) GetShadersCount(shaders []XRGPUAGXShaderTimelineSignpostProcessRef, count *uint64) {
	objc.SendIfResponds[objc.ID](x.ID, objc.Sel("getShaders:count:"), objc.CArray(shaders), unsafe.Pointer(count))
}
func (x XRGPUAGXShaderTimelineSignposts) Load(load objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](x.ID, objc.Sel("load:"), load)
	return rv
}
func (x XRGPUAGXShaderTimelineSignposts) Start() bool {
	rv := objc.SendIfResponds[bool](x.ID, objc.Sel("start"))
	return rv
}
func (x XRGPUAGXShaderTimelineSignposts) Stop() {
	objc.SendIfResponds[objc.ID](x.ID, objc.Sel("stop"))
}
func (x XRGPUAGXShaderTimelineSignposts) InitWithCoder(coder foundation.INSCoder) XRGPUAGXShaderTimelineSignposts {
	rv := objc.SendIfResponds[XRGPUAGXShaderTimelineSignposts](x.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_XRGPUAGXShaderTimelineSignpostsClass XRGPUAGXShaderTimelineSignpostsClass) FromDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_XRGPUAGXShaderTimelineSignpostsClass.class), objc.Sel("fromData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_XRGPUAGXShaderTimelineSignpostsClass XRGPUAGXShaderTimelineSignpostsClass) FromFileError(file objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_XRGPUAGXShaderTimelineSignpostsClass.class), objc.Sel("fromFile:error:"), file, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_XRGPUAGXShaderTimelineSignpostsClass XRGPUAGXShaderTimelineSignpostsClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_XRGPUAGXShaderTimelineSignpostsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// EnumerateKickIdsSync is a synchronous wrapper around [XRGPUAGXShaderTimelineSignposts.EnumerateKickIds].
// It blocks until the completion handler fires or the context is cancelled.
func (x XRGPUAGXShaderTimelineSignposts) EnumerateKickIdsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	x.EnumerateKickIds(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
