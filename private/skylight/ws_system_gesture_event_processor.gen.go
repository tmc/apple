// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSSystemGestureEventProcessor] class.
var (
	_WSSystemGestureEventProcessorClass     WSSystemGestureEventProcessorClass
	_WSSystemGestureEventProcessorClassOnce sync.Once
)

func getWSSystemGestureEventProcessorClass() WSSystemGestureEventProcessorClass {
	_WSSystemGestureEventProcessorClassOnce.Do(func() {
		_WSSystemGestureEventProcessorClass = WSSystemGestureEventProcessorClass{class: objc.GetClass("WSSystemGestureEventProcessor")}
	})
	return _WSSystemGestureEventProcessorClass
}

// GetWSSystemGestureEventProcessorClass returns the class object for WSSystemGestureEventProcessor.
func GetWSSystemGestureEventProcessorClass() WSSystemGestureEventProcessorClass {
	return getWSSystemGestureEventProcessorClass()
}

type WSSystemGestureEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSSystemGestureEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSSystemGestureEventProcessorClass) Alloc() WSSystemGestureEventProcessor {
	rv := objc.SendIfResponds[WSSystemGestureEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSSystemGestureEventProcessor.ApplyGesturePayloadsTargetIdentifier]
//   - [WSSystemGestureEventProcessor.ConfigureDirectTouchEventProcessor]
//   - [WSSystemGestureEventProcessor.SystemGestureEventProcessorGlue]
//   - [WSSystemGestureEventProcessor.TargetIdentifierDetached]
//   - [WSSystemGestureEventProcessor.WindowTerminationObserver]
//   - [WSSystemGestureEventProcessor.SetWindowTerminationObserver]
//   - [WSSystemGestureEventProcessor.InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterTouchCancellationHandler]
//   - [WSSystemGestureEventProcessor.InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterWindowTerminationNotifierTouchCancellationHandler]
type WSSystemGestureEventProcessor struct {
	objectivec.Object
}

// WSSystemGestureEventProcessorFromID constructs a [WSSystemGestureEventProcessor] from an objc.ID.
func WSSystemGestureEventProcessorFromID(id objc.ID) WSSystemGestureEventProcessor {
	return WSSystemGestureEventProcessor{objectivec.Object{ID: id}}
}

// Ensure WSSystemGestureEventProcessor implements IWSSystemGestureEventProcessor.
var _ IWSSystemGestureEventProcessor = WSSystemGestureEventProcessor{}

// An interface definition for the [WSSystemGestureEventProcessor] class.
//
// # Methods
//
//   - [IWSSystemGestureEventProcessor.ApplyGesturePayloadsTargetIdentifier]
//   - [IWSSystemGestureEventProcessor.ConfigureDirectTouchEventProcessor]
//   - [IWSSystemGestureEventProcessor.SystemGestureEventProcessorGlue]
//   - [IWSSystemGestureEventProcessor.TargetIdentifierDetached]
//   - [IWSSystemGestureEventProcessor.WindowTerminationObserver]
//   - [IWSSystemGestureEventProcessor.SetWindowTerminationObserver]
//   - [IWSSystemGestureEventProcessor.InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterTouchCancellationHandler]
//   - [IWSSystemGestureEventProcessor.InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterWindowTerminationNotifierTouchCancellationHandler]
type IWSSystemGestureEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	ApplyGesturePayloadsTargetIdentifier(payloads objectivec.IObject, identifier unsafe.Pointer)
	ConfigureDirectTouchEventProcessor(processor objectivec.IObject)
	SystemGestureEventProcessorGlue() unsafe.Pointer
	TargetIdentifierDetached(detached unsafe.Pointer)
	WindowTerminationObserver() unsafe.Pointer
	SetWindowTerminationObserver(value unsafe.Pointer)
	InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterTouchCancellationHandler(collector objectivec.IObject, cache objectivec.IObject, converter objectivec.IObject, handler VoidHandler) WSSystemGestureEventProcessor
	InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterWindowTerminationNotifierTouchCancellationHandler(collector objectivec.IObject, cache objectivec.IObject, converter objectivec.IObject, notifier objectivec.IObject, handler VoidHandler) WSSystemGestureEventProcessor
}

// Init initializes the instance.
func (w WSSystemGestureEventProcessor) Init() WSSystemGestureEventProcessor {
	rv := objc.SendIfResponds[WSSystemGestureEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSSystemGestureEventProcessor) Autorelease() WSSystemGestureEventProcessor {
	rv := objc.SendIfResponds[WSSystemGestureEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSSystemGestureEventProcessor creates a new WSSystemGestureEventProcessor instance.
func NewWSSystemGestureEventProcessor() WSSystemGestureEventProcessor {
	class := getWSSystemGestureEventProcessorClass()
	rv := objc.SendIfResponds[WSSystemGestureEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSSystemGestureEventProcessor) ApplyGesturePayloadsTargetIdentifier(payloads objectivec.IObject, identifier unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("applyGesturePayloads:targetIdentifier:"), payloads, identifier)
}
func (w WSSystemGestureEventProcessor) ConfigureDirectTouchEventProcessor(processor objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("configureDirectTouchEventProcessor:"), processor)
}
func (w WSSystemGestureEventProcessor) TargetIdentifierDetached(detached unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("targetIdentifierDetached:"), detached)
}
func (w WSSystemGestureEventProcessor) InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterTouchCancellationHandler(collector objectivec.IObject, cache objectivec.IObject, converter objectivec.IObject, handler VoidHandler) WSSystemGestureEventProcessor {
	_block3, _ := NewVoidBlock(handler)
	rv := objc.SendIfResponds[WSSystemGestureEventProcessor](w.ID, objc.Sel("initWithGestureEventCollector:hidEventSenderCache:coordinateSpaceConverter:touchCancellationHandler:"), collector, cache, converter, _block3)
	return rv
}
func (w WSSystemGestureEventProcessor) InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterWindowTerminationNotifierTouchCancellationHandler(collector objectivec.IObject, cache objectivec.IObject, converter objectivec.IObject, notifier objectivec.IObject, handler VoidHandler) WSSystemGestureEventProcessor {
	_block4, _ := NewVoidBlock(handler)
	rv := objc.SendIfResponds[WSSystemGestureEventProcessor](w.ID, objc.Sel("initWithGestureEventCollector:hidEventSenderCache:coordinateSpaceConverter:windowTerminationNotifier:touchCancellationHandler:"), collector, cache, converter, notifier, _block4)
	return rv
}

func (w WSSystemGestureEventProcessor) SystemGestureEventProcessorGlue() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("systemGestureEventProcessorGlue"))
	return rv
}
func (w WSSystemGestureEventProcessor) WindowTerminationObserver() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("windowTerminationObserver"))
	return rv
}
func (w WSSystemGestureEventProcessor) SetWindowTerminationObserver(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setWindowTerminationObserver:"), value)
}

// InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterTouchCancellationHandlerSync is a synchronous wrapper around [WSSystemGestureEventProcessor.InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterTouchCancellationHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WSSystemGestureEventProcessor) InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterTouchCancellationHandlerSync(ctx context.Context, collector objectivec.IObject, cache objectivec.IObject, converter objectivec.IObject) error {
	done := make(chan struct{}, 1)
	w.InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterTouchCancellationHandler(collector, cache, converter, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterWindowTerminationNotifierTouchCancellationHandlerSync is a synchronous wrapper around [WSSystemGestureEventProcessor.InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterWindowTerminationNotifierTouchCancellationHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w WSSystemGestureEventProcessor) InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterWindowTerminationNotifierTouchCancellationHandlerSync(ctx context.Context, collector objectivec.IObject, cache objectivec.IObject, converter objectivec.IObject, notifier objectivec.IObject) error {
	done := make(chan struct{}, 1)
	w.InitWithGestureEventCollectorHidEventSenderCacheCoordinateSpaceConverterWindowTerminationNotifierTouchCancellationHandler(collector, cache, converter, notifier, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
