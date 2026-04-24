// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGCoreUIWork] class.
var (
	_PKGCoreUIWorkClass     PKGCoreUIWorkClass
	_PKGCoreUIWorkClassOnce sync.Once
)

func getPKGCoreUIWorkClass() PKGCoreUIWorkClass {
	_PKGCoreUIWorkClassOnce.Do(func() {
		_PKGCoreUIWorkClass = PKGCoreUIWorkClass{class: objc.GetClass("PKGCoreUIWork")}
	})
	return _PKGCoreUIWorkClass
}

// GetPKGCoreUIWorkClass returns the class object for PKGCoreUIWork.
func GetPKGCoreUIWorkClass() PKGCoreUIWorkClass {
	return getPKGCoreUIWorkClass()
}

type PKGCoreUIWorkClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGCoreUIWorkClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGCoreUIWorkClass) Alloc() PKGCoreUIWork {
	rv := objc.Send[PKGCoreUIWork](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGCoreUIWork.SetMainThreadWork]
//   - [PKGCoreUIWork.SetRendererWork]
//
// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIWork
type PKGCoreUIWork struct {
	objectivec.Object
}

// PKGCoreUIWorkFromID constructs a [PKGCoreUIWork] from an objc.ID.
func PKGCoreUIWorkFromID(id objc.ID) PKGCoreUIWork {
	return PKGCoreUIWork{objectivec.Object{ID: id}}
}

// Ensure PKGCoreUIWork implements IPKGCoreUIWork.
var _ IPKGCoreUIWork = PKGCoreUIWork{}

// An interface definition for the [PKGCoreUIWork] class.
//
// # Methods
//
//   - [IPKGCoreUIWork.SetMainThreadWork]
//   - [IPKGCoreUIWork.SetRendererWork]
//
// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIWork
type IPKGCoreUIWork interface {
	objectivec.IObject

	// Topic: Methods

	SetMainThreadWork(work VoidHandler)
	SetRendererWork(work VoidHandler)
}

// Init initializes the instance.
func (g PKGCoreUIWork) Init() PKGCoreUIWork {
	rv := objc.Send[PKGCoreUIWork](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g PKGCoreUIWork) Autorelease() PKGCoreUIWork {
	rv := objc.Send[PKGCoreUIWork](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGCoreUIWork creates a new PKGCoreUIWork instance.
func NewPKGCoreUIWork() PKGCoreUIWork {
	class := getPKGCoreUIWorkClass()
	rv := objc.Send[PKGCoreUIWork](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIWork/setMainThreadWork:
func (g PKGCoreUIWork) SetMainThreadWork(work VoidHandler) {
	_block0, _ := NewVoidBlock(work)
	objc.Send[objc.ID](g.ID, objc.Sel("setMainThreadWork:"), _block0)
}

// See: https://developer.apple.com/documentation/SkyLight/PKGCoreUIWork/setRendererWork:
func (g PKGCoreUIWork) SetRendererWork(work VoidHandler) {
	_block0, _ := NewVoidBlock(work)
	objc.Send[objc.ID](g.ID, objc.Sel("setRendererWork:"), _block0)
}

// SetMainThreadWorkSync is a synchronous wrapper around [PKGCoreUIWork.SetMainThreadWork].
// It blocks until the completion handler fires or the context is cancelled.
func (g PKGCoreUIWork) SetMainThreadWorkSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.SetMainThreadWork(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetRendererWorkSync is a synchronous wrapper around [PKGCoreUIWork.SetRendererWork].
// It blocks until the completion handler fires or the context is cancelled.
func (g PKGCoreUIWork) SetRendererWorkSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.SetRendererWork(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
