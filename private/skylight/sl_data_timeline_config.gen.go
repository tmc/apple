// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLDataTimelineConfig] class.
var (
	_SLDataTimelineConfigClass     SLDataTimelineConfigClass
	_SLDataTimelineConfigClassOnce sync.Once
)

func getSLDataTimelineConfigClass() SLDataTimelineConfigClass {
	_SLDataTimelineConfigClassOnce.Do(func() {
		_SLDataTimelineConfigClass = SLDataTimelineConfigClass{class: objc.GetClass("SLDataTimelineConfig")}
	})
	return _SLDataTimelineConfigClass
}

// GetSLDataTimelineConfigClass returns the class object for SLDataTimelineConfig.
func GetSLDataTimelineConfigClass() SLDataTimelineConfigClass {
	return getSLDataTimelineConfigClass()
}

type SLDataTimelineConfigClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLDataTimelineConfigClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLDataTimelineConfigClass) Alloc() SLDataTimelineConfig {
	rv := objc.Send[SLDataTimelineConfig](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLDataTimelineConfig.AddInfoOption]
//   - [SLDataTimelineConfig.ConnectionQueue]
//   - [SLDataTimelineConfig.CreateCancellableMachRecvSourceWithQueueCancelActionError]
//   - [SLDataTimelineConfig.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler]
//   - [SLDataTimelineConfig.CreateXPCObject]
//   - [SLDataTimelineConfig.EstablishConnectionWithResultBlock]
//   - [SLDataTimelineConfig.InfoOptions]
//   - [SLDataTimelineConfig.Name]
//   - [SLDataTimelineConfig.ReportIntervals]
//   - [SLDataTimelineConfig.RequestReportIntervalValueForKey]
//   - [SLDataTimelineConfig.RequestSampleIntervalValueForKey]
//   - [SLDataTimelineConfig.SampleIntervals]
//   - [SLDataTimelineConfig.SetTargetQueue]
//   - [SLDataTimelineConfig.UpdateBlock]
//   - [SLDataTimelineConfig.UpdateBlockQueue]
//   - [SLDataTimelineConfig.SetUpdateBlockQueue]
//   - [SLDataTimelineConfig.InitWithNameAndUpdateBlock]
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig
type SLDataTimelineConfig struct {
	objectivec.Object
}

// SLDataTimelineConfigFromID constructs a [SLDataTimelineConfig] from an objc.ID.
func SLDataTimelineConfigFromID(id objc.ID) SLDataTimelineConfig {
	return SLDataTimelineConfig{objectivec.Object{ID: id}}
}

// Ensure SLDataTimelineConfig implements ISLDataTimelineConfig.
var _ ISLDataTimelineConfig = SLDataTimelineConfig{}

// An interface definition for the [SLDataTimelineConfig] class.
//
// # Methods
//
//   - [ISLDataTimelineConfig.AddInfoOption]
//   - [ISLDataTimelineConfig.ConnectionQueue]
//   - [ISLDataTimelineConfig.CreateCancellableMachRecvSourceWithQueueCancelActionError]
//   - [ISLDataTimelineConfig.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler]
//   - [ISLDataTimelineConfig.CreateXPCObject]
//   - [ISLDataTimelineConfig.EstablishConnectionWithResultBlock]
//   - [ISLDataTimelineConfig.InfoOptions]
//   - [ISLDataTimelineConfig.Name]
//   - [ISLDataTimelineConfig.ReportIntervals]
//   - [ISLDataTimelineConfig.RequestReportIntervalValueForKey]
//   - [ISLDataTimelineConfig.RequestSampleIntervalValueForKey]
//   - [ISLDataTimelineConfig.SampleIntervals]
//   - [ISLDataTimelineConfig.SetTargetQueue]
//   - [ISLDataTimelineConfig.UpdateBlock]
//   - [ISLDataTimelineConfig.UpdateBlockQueue]
//   - [ISLDataTimelineConfig.SetUpdateBlockQueue]
//   - [ISLDataTimelineConfig.InitWithNameAndUpdateBlock]
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig
type ISLDataTimelineConfig interface {
	objectivec.IObject

	// Topic: Methods

	AddInfoOption(option objectivec.IObject)
	ConnectionQueue() objectivec.Object
	CreateCancellableMachRecvSourceWithQueueCancelActionError(queue objectivec.IObject, action func()) (objectivec.IObject, error)
	CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler(queue objectivec.IObject, handler VoidHandler, handler2 VoidHandler) objectivec.IObject
	CreateXPCObject() objectivec.IObject
	EstablishConnectionWithResultBlock(block VoidHandler)
	InfoOptions() foundation.INSSet
	Name() string
	ReportIntervals() foundation.INSDictionary
	RequestReportIntervalValueForKey(value uint16, key objectivec.IObject)
	RequestSampleIntervalValueForKey(value uint16, key objectivec.IObject)
	SampleIntervals() foundation.INSDictionary
	SetTargetQueue(queue objectivec.IObject)
	UpdateBlock() VoidHandler
	UpdateBlockQueue() objectivec.Object
	SetUpdateBlockQueue(value objectivec.Object)
	InitWithNameAndUpdateBlock(name objectivec.IObject, block VoidHandler) SLDataTimelineConfig
}

// Init initializes the instance.
func (s SLDataTimelineConfig) Init() SLDataTimelineConfig {
	rv := objc.Send[SLDataTimelineConfig](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLDataTimelineConfig) Autorelease() SLDataTimelineConfig {
	rv := objc.Send[SLDataTimelineConfig](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLDataTimelineConfig creates a new SLDataTimelineConfig instance.
func NewSLDataTimelineConfig() SLDataTimelineConfig {
	class := getSLDataTimelineConfigClass()
	rv := objc.Send[SLDataTimelineConfig](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/addInfoOption:
func (s SLDataTimelineConfig) AddInfoOption(option objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("addInfoOption:"), option)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/createCancellableMachRecvSourceWithQueue:cancelAction:error:
func (s SLDataTimelineConfig) CreateCancellableMachRecvSourceWithQueueCancelActionError(queue objectivec.IObject, action func()) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createCancellableMachRecvSourceWithQueue:cancelAction:error:"), queue, action, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/createNoSenderRecvPairWithQueue:errorHandler:eventHandler:
func (s SLDataTimelineConfig) CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler(queue objectivec.IObject, handler VoidHandler, handler2 VoidHandler) objectivec.IObject {
	_block1, _ := NewVoidBlock(handler)
	_block2, _ := NewVoidBlock(handler2)
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createNoSenderRecvPairWithQueue:errorHandler:eventHandler:"), queue, _block1, _block2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/createXPCObject
func (s SLDataTimelineConfig) CreateXPCObject() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createXPCObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/establishConnectionWithResultBlock:
func (s SLDataTimelineConfig) EstablishConnectionWithResultBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("establishConnectionWithResultBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/requestReportIntervalValue:forKey:
func (s SLDataTimelineConfig) RequestReportIntervalValueForKey(value uint16, key objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("requestReportIntervalValue:forKey:"), value, key)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/requestSampleIntervalValue:forKey:
func (s SLDataTimelineConfig) RequestSampleIntervalValueForKey(value uint16, key objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("requestSampleIntervalValue:forKey:"), value, key)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/setTargetQueue:
func (s SLDataTimelineConfig) SetTargetQueue(queue objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setTargetQueue:"), queue)
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/initWithName:andUpdateBlock:
func (s SLDataTimelineConfig) InitWithNameAndUpdateBlock(name objectivec.IObject, block VoidHandler) SLDataTimelineConfig {
	_block1, _ := NewVoidBlock(block)
	rv := objc.Send[SLDataTimelineConfig](s.ID, objc.Sel("initWithName:andUpdateBlock:"), name, _block1)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/configWithName:andUpdateBlock:
func (_SLDataTimelineConfigClass SLDataTimelineConfigClass) ConfigWithNameAndUpdateBlock(name objectivec.IObject, block VoidHandler) objectivec.IObject {
	_block1, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](objc.ID(_SLDataTimelineConfigClass.class), objc.Sel("configWithName:andUpdateBlock:"), name, _block1)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/connectionQueue
func (s SLDataTimelineConfig) ConnectionQueue() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("connectionQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/infoOptions
func (s SLDataTimelineConfig) InfoOptions() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("infoOptions"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/name
func (s SLDataTimelineConfig) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/reportIntervals
func (s SLDataTimelineConfig) ReportIntervals() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("reportIntervals"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/sampleIntervals
func (s SLDataTimelineConfig) SampleIntervals() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sampleIntervals"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/updateBlock
func (s SLDataTimelineConfig) UpdateBlock() VoidHandler {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("updateBlock"))
	_ = rv
	return nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineConfig/updateBlockQueue
func (s SLDataTimelineConfig) UpdateBlockQueue() objectivec.Object {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("updateBlockQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (s SLDataTimelineConfig) SetUpdateBlockQueue(value objectivec.Object) {
	objc.Send[struct{}](s.ID, objc.Sel("setUpdateBlockQueue:"), value)
}

// CreateNoSenderRecvPairWithQueueErrorHandlerEventHandlerSync is a synchronous wrapper around [SLDataTimelineConfig.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLDataTimelineConfig) CreateNoSenderRecvPairWithQueueErrorHandlerEventHandlerSync(ctx context.Context, queue objectivec.IObject, handler VoidHandler) error {
	done := make(chan struct{}, 1)
	s.CreateNoSenderRecvPairWithQueueErrorHandlerEventHandler(queue, handler, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EstablishConnectionWithResultBlockSync is a synchronous wrapper around [SLDataTimelineConfig.EstablishConnectionWithResultBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLDataTimelineConfig) EstablishConnectionWithResultBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.EstablishConnectionWithResultBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithNameAndUpdateBlockSync is a synchronous wrapper around [SLDataTimelineConfig.InitWithNameAndUpdateBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLDataTimelineConfig) InitWithNameAndUpdateBlockSync(ctx context.Context, name objectivec.IObject) error {
	done := make(chan struct{}, 1)
	s.InitWithNameAndUpdateBlock(name, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ConfigWithNameAndUpdateBlockSync is a synchronous wrapper around [SLDataTimelineConfig.ConfigWithNameAndUpdateBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (sc SLDataTimelineConfigClass) ConfigWithNameAndUpdateBlockSync(ctx context.Context, name objectivec.IObject) error {
	done := make(chan struct{}, 1)
	sc.ConfigWithNameAndUpdateBlock(name, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
