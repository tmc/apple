// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DYGPUDerivedEncoderCounterInfo] class.
var (
	_DYGPUDerivedEncoderCounterInfoClass     DYGPUDerivedEncoderCounterInfoClass
	_DYGPUDerivedEncoderCounterInfoClassOnce sync.Once
)

func getDYGPUDerivedEncoderCounterInfoClass() DYGPUDerivedEncoderCounterInfoClass {
	_DYGPUDerivedEncoderCounterInfoClassOnce.Do(func() {
		_DYGPUDerivedEncoderCounterInfoClass = DYGPUDerivedEncoderCounterInfoClass{class: objc.GetClass("DYGPUDerivedEncoderCounterInfo")}
	})
	return _DYGPUDerivedEncoderCounterInfoClass
}

// GetDYGPUDerivedEncoderCounterInfoClass returns the class object for DYGPUDerivedEncoderCounterInfo.
func GetDYGPUDerivedEncoderCounterInfoClass() DYGPUDerivedEncoderCounterInfoClass {
	return getDYGPUDerivedEncoderCounterInfoClass()
}

type DYGPUDerivedEncoderCounterInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DYGPUDerivedEncoderCounterInfoClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DYGPUDerivedEncoderCounterInfoClass) Alloc() DYGPUDerivedEncoderCounterInfo {
	rv := objc.Send[DYGPUDerivedEncoderCounterInfo](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedData]
//   - [DYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedDataAtIndexWithBlock]
//   - [DYGPUDerivedEncoderCounterInfo.DerivedCounterNames]
//   - [DYGPUDerivedEncoderCounterInfo.SetDerivedCounterNames]
//   - [DYGPUDerivedEncoderCounterInfo.DerivedCounters]
//   - [DYGPUDerivedEncoderCounterInfo.SetDerivedCounters]
//   - [DYGPUDerivedEncoderCounterInfo.EncodeWithCoder]
//   - [DYGPUDerivedEncoderCounterInfo.EncoderInfos]
//   - [DYGPUDerivedEncoderCounterInfo.SetEncoderInfos]
//   - [DYGPUDerivedEncoderCounterInfo.InitWithCoder]
type DYGPUDerivedEncoderCounterInfo struct {
	objectivec.Object
}

// DYGPUDerivedEncoderCounterInfoFromID constructs a [DYGPUDerivedEncoderCounterInfo] from an objc.ID.
func DYGPUDerivedEncoderCounterInfoFromID(id objc.ID) DYGPUDerivedEncoderCounterInfo {
	return DYGPUDerivedEncoderCounterInfo{objectivec.Object{ID: id}}
}

// Ensure DYGPUDerivedEncoderCounterInfo implements IDYGPUDerivedEncoderCounterInfo.
var _ IDYGPUDerivedEncoderCounterInfo = DYGPUDerivedEncoderCounterInfo{}

// An interface definition for the [DYGPUDerivedEncoderCounterInfo] class.
//
// # Methods
//
//   - [IDYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedData]
//   - [IDYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedDataAtIndexWithBlock]
//   - [IDYGPUDerivedEncoderCounterInfo.DerivedCounterNames]
//   - [IDYGPUDerivedEncoderCounterInfo.SetDerivedCounterNames]
//   - [IDYGPUDerivedEncoderCounterInfo.DerivedCounters]
//   - [IDYGPUDerivedEncoderCounterInfo.SetDerivedCounters]
//   - [IDYGPUDerivedEncoderCounterInfo.EncodeWithCoder]
//   - [IDYGPUDerivedEncoderCounterInfo.EncoderInfos]
//   - [IDYGPUDerivedEncoderCounterInfo.SetEncoderInfos]
//   - [IDYGPUDerivedEncoderCounterInfo.InitWithCoder]
type IDYGPUDerivedEncoderCounterInfo interface {
	objectivec.IObject

	// Topic: Methods

	_enumerateEncoderDerivedData(data VoidHandler)
	_enumerateEncoderDerivedDataAtIndexWithBlock(index uint32, block VoidHandler)
	DerivedCounterNames() foundation.INSArray
	SetDerivedCounterNames(value foundation.INSArray)
	DerivedCounters() foundation.NSData
	SetDerivedCounters(value foundation.NSData)
	EncodeWithCoder(coder foundation.INSCoder)
	EncoderInfos() foundation.NSData
	SetEncoderInfos(value foundation.NSData)
	InitWithCoder(coder foundation.INSCoder) DYGPUDerivedEncoderCounterInfo
}

// Init initializes the instance.
func (d DYGPUDerivedEncoderCounterInfo) Init() DYGPUDerivedEncoderCounterInfo {
	rv := objc.Send[DYGPUDerivedEncoderCounterInfo](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DYGPUDerivedEncoderCounterInfo) Autorelease() DYGPUDerivedEncoderCounterInfo {
	rv := objc.Send[DYGPUDerivedEncoderCounterInfo](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDYGPUDerivedEncoderCounterInfo creates a new DYGPUDerivedEncoderCounterInfo instance.
func NewDYGPUDerivedEncoderCounterInfo() DYGPUDerivedEncoderCounterInfo {
	class := getDYGPUDerivedEncoderCounterInfoClass()
	rv := objc.Send[DYGPUDerivedEncoderCounterInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDYGPUDerivedEncoderCounterInfoWithCoder(coder objectivec.IObject) DYGPUDerivedEncoderCounterInfo {
	instance := getDYGPUDerivedEncoderCounterInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DYGPUDerivedEncoderCounterInfoFromID(rv)
}

func (d DYGPUDerivedEncoderCounterInfo) _enumerateEncoderDerivedData(data VoidHandler) {
	_block0, _ := NewVoidBlock(data)
	objc.Send[objc.ID](d.ID, objc.Sel("_enumerateEncoderDerivedData:"), _block0)
}

// EnumerateEncoderDerivedData is an exported wrapper for the private method _enumerateEncoderDerivedData.
func (d DYGPUDerivedEncoderCounterInfo) EnumerateEncoderDerivedData(data VoidHandler) error {
	if !objc.RespondsToSelector(d.ID, objc.Sel("_enumerateEncoderDerivedData:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_enumerateEncoderDerivedData:"}
		return err
	}
	d._enumerateEncoderDerivedData(data)
	return nil
}

// CanEnumerateEncoderDerivedData reports whether the receiver responds to the private selector _enumerateEncoderDerivedData:.
func (d DYGPUDerivedEncoderCounterInfo) CanEnumerateEncoderDerivedData() bool {
	return objc.RespondsToSelector(d.ID, objc.Sel("_enumerateEncoderDerivedData:"))
}
func (d DYGPUDerivedEncoderCounterInfo) _enumerateEncoderDerivedDataAtIndexWithBlock(index uint32, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](d.ID, objc.Sel("_enumerateEncoderDerivedDataAtIndex:withBlock:"), index, _block1)
}

// EnumerateEncoderDerivedDataAtIndexWithBlock is an exported wrapper for the private method _enumerateEncoderDerivedDataAtIndexWithBlock.
func (d DYGPUDerivedEncoderCounterInfo) EnumerateEncoderDerivedDataAtIndexWithBlock(index uint32, block VoidHandler) error {
	if !objc.RespondsToSelector(d.ID, objc.Sel("_enumerateEncoderDerivedDataAtIndex:withBlock:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_enumerateEncoderDerivedDataAtIndex:withBlock:"}
		return err
	}
	d._enumerateEncoderDerivedDataAtIndexWithBlock(index, block)
	return nil
}

// CanEnumerateEncoderDerivedDataAtIndexWithBlock reports whether the receiver responds to the private selector _enumerateEncoderDerivedDataAtIndex:withBlock:.
func (d DYGPUDerivedEncoderCounterInfo) CanEnumerateEncoderDerivedDataAtIndexWithBlock() bool {
	return objc.RespondsToSelector(d.ID, objc.Sel("_enumerateEncoderDerivedDataAtIndex:withBlock:"))
}
func (d DYGPUDerivedEncoderCounterInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (d DYGPUDerivedEncoderCounterInfo) InitWithCoder(coder foundation.INSCoder) DYGPUDerivedEncoderCounterInfo {
	rv := objc.Send[DYGPUDerivedEncoderCounterInfo](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_DYGPUDerivedEncoderCounterInfoClass DYGPUDerivedEncoderCounterInfoClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DYGPUDerivedEncoderCounterInfoClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (d DYGPUDerivedEncoderCounterInfo) DerivedCounterNames() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("derivedCounterNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (d DYGPUDerivedEncoderCounterInfo) SetDerivedCounterNames(value foundation.INSArray) {
	objc.Send[struct{}](d.ID, objc.Sel("setDerivedCounterNames:"), value)
}
func (d DYGPUDerivedEncoderCounterInfo) DerivedCounters() foundation.NSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("derivedCounters"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUDerivedEncoderCounterInfo) SetDerivedCounters(value foundation.NSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setDerivedCounters:"), value)
}
func (d DYGPUDerivedEncoderCounterInfo) EncoderInfos() foundation.NSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("encoderInfos"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DYGPUDerivedEncoderCounterInfo) SetEncoderInfos(value foundation.NSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setEncoderInfos:"), value)
}

// _enumerateEncoderDerivedDataSync is a synchronous wrapper around [DYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedData].
// It blocks until the completion handler fires or the context is cancelled.
func (d DYGPUDerivedEncoderCounterInfo) _enumerateEncoderDerivedDataSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d._enumerateEncoderDerivedData(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _enumerateEncoderDerivedDataAtIndexWithBlockSync is a synchronous wrapper around [DYGPUDerivedEncoderCounterInfo._enumerateEncoderDerivedDataAtIndexWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d DYGPUDerivedEncoderCounterInfo) _enumerateEncoderDerivedDataAtIndexWithBlockSync(ctx context.Context, index uint32) error {
	done := make(chan struct{}, 1)
	d._enumerateEncoderDerivedDataAtIndexWithBlock(index, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
