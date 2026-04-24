// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGCoreUITransaction] class.
var (
	_PKGCoreUITransactionClass     PKGCoreUITransactionClass
	_PKGCoreUITransactionClassOnce sync.Once
)

func getPKGCoreUITransactionClass() PKGCoreUITransactionClass {
	_PKGCoreUITransactionClassOnce.Do(func() {
		_PKGCoreUITransactionClass = PKGCoreUITransactionClass{class: objc.GetClass("PKGCoreUITransaction")}
	})
	return _PKGCoreUITransactionClass
}

// GetPKGCoreUITransactionClass returns the class object for PKGCoreUITransaction.
func GetPKGCoreUITransactionClass() PKGCoreUITransactionClass {
	return getPKGCoreUITransactionClass()
}

type PKGCoreUITransactionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGCoreUITransactionClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGCoreUITransactionClass) Alloc() PKGCoreUITransaction {
	rv := objc.Send[PKGCoreUITransaction](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGCoreUITransaction._layerUpdateKeyForOptions]
//   - [PKGCoreUITransaction._scheduleRendererWorkMainThreadWork]
//   - [PKGCoreUITransaction.Commit]
//   - [PKGCoreUITransaction.UpdateLayerKeyRendererWork]
//   - [PKGCoreUITransaction.InitWithThemeUseAX]
//
// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUITransaction
type PKGCoreUITransaction struct {
	objectivec.Object
}

// PKGCoreUITransactionFromID constructs a [PKGCoreUITransaction] from an objc.ID.
func PKGCoreUITransactionFromID(id objc.ID) PKGCoreUITransaction {
	return PKGCoreUITransaction{objectivec.Object{ID: id}}
}

// Ensure PKGCoreUITransaction implements IPKGCoreUITransaction.
var _ IPKGCoreUITransaction = PKGCoreUITransaction{}

// An interface definition for the [PKGCoreUITransaction] class.
//
// # Methods
//
//   - [IPKGCoreUITransaction._layerUpdateKeyForOptions]
//   - [IPKGCoreUITransaction._scheduleRendererWorkMainThreadWork]
//   - [IPKGCoreUITransaction.Commit]
//   - [IPKGCoreUITransaction.UpdateLayerKeyRendererWork]
//   - [IPKGCoreUITransaction.InitWithThemeUseAX]
//
// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUITransaction
type IPKGCoreUITransaction interface {
	objectivec.IObject

	// Topic: Methods

	_layerUpdateKeyForOptions(options objectivec.IObject) objectivec.IObject
	_scheduleRendererWorkMainThreadWork(work VoidHandler, work2 VoidHandler)
	Commit()
	UpdateLayerKeyRendererWork(layer objectivec.IObject, key objectivec.IObject, work VoidHandler)
	InitWithThemeUseAX(theme uint32, ax bool) PKGCoreUITransaction
}

// Init initializes the instance.
func (g PKGCoreUITransaction) Init() PKGCoreUITransaction {
	rv := objc.Send[PKGCoreUITransaction](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g PKGCoreUITransaction) Autorelease() PKGCoreUITransaction {
	rv := objc.Send[PKGCoreUITransaction](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGCoreUITransaction creates a new PKGCoreUITransaction instance.
func NewPKGCoreUITransaction() PKGCoreUITransaction {
	class := getPKGCoreUITransactionClass()
	rv := objc.Send[PKGCoreUITransaction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUITransaction/initWithTheme:useAX:
func NewGCoreUITransactionWithThemeUseAX(theme uint32, ax bool) PKGCoreUITransaction {
	instance := getPKGCoreUITransactionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTheme:useAX:"), theme, ax)
	return PKGCoreUITransactionFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUITransaction/_layerUpdateKeyForOptions:
func (g PKGCoreUITransaction) _layerUpdateKeyForOptions(options objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_layerUpdateKeyForOptions:"), options)
	return objectivec.Object{ID: rv}
}

// LayerUpdateKeyForOptions is an exported wrapper for the private method _layerUpdateKeyForOptions.
func (g PKGCoreUITransaction) LayerUpdateKeyForOptions(options objectivec.IObject) objectivec.IObject {
	return g._layerUpdateKeyForOptions(options)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUITransaction/_scheduleRendererWork:mainThreadWork:
func (g PKGCoreUITransaction) _scheduleRendererWorkMainThreadWork(work VoidHandler, work2 VoidHandler) {
	_block0, _ := NewVoidBlock(work)
	_block1, _ := NewVoidBlock(work2)
	objc.Send[objc.ID](g.ID, objc.Sel("_scheduleRendererWork:mainThreadWork:"), _block0, _block1)
}

// ScheduleRendererWorkMainThreadWork is an exported wrapper for the private method _scheduleRendererWorkMainThreadWork.
func (g PKGCoreUITransaction) ScheduleRendererWorkMainThreadWork(work VoidHandler, work2 VoidHandler) {
	g._scheduleRendererWorkMainThreadWork(work, work2)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUITransaction/commit
func (g PKGCoreUITransaction) Commit() {
	objc.Send[objc.ID](g.ID, objc.Sel("commit"))
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUITransaction/updateLayer:key:rendererWork:
func (g PKGCoreUITransaction) UpdateLayerKeyRendererWork(layer objectivec.IObject, key objectivec.IObject, work VoidHandler) {
	_block2, _ := NewVoidBlock(work)
	objc.Send[objc.ID](g.ID, objc.Sel("updateLayer:key:rendererWork:"), layer, key, _block2)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUITransaction/initWithTheme:useAX:
func (g PKGCoreUITransaction) InitWithThemeUseAX(theme uint32, ax bool) PKGCoreUITransaction {
	rv := objc.Send[PKGCoreUITransaction](g.ID, objc.Sel("initWithTheme:useAX:"), theme, ax)
	return rv
}

// _scheduleRendererWorkMainThreadWorkSync is a synchronous wrapper around [PKGCoreUITransaction._scheduleRendererWorkMainThreadWork].
// It blocks until the completion handler fires or the context is cancelled.
func (g PKGCoreUITransaction) _scheduleRendererWorkMainThreadWorkSync(ctx context.Context, work VoidHandler) error {
	done := make(chan struct{}, 1)
	g._scheduleRendererWorkMainThreadWork(work, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// UpdateLayerKeyRendererWorkSync is a synchronous wrapper around [PKGCoreUITransaction.UpdateLayerKeyRendererWork].
// It blocks until the completion handler fires or the context is cancelled.
func (g PKGCoreUITransaction) UpdateLayerKeyRendererWorkSync(ctx context.Context, layer objectivec.IObject, key objectivec.IObject) error {
	done := make(chan struct{}, 1)
	g.UpdateLayerKeyRendererWork(layer, key, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
