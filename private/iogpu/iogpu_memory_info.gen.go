// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMemoryInfo] class.
var (
	_IOGPUMemoryInfoClass     IOGPUMemoryInfoClass
	_IOGPUMemoryInfoClassOnce sync.Once
)

func getIOGPUMemoryInfoClass() IOGPUMemoryInfoClass {
	_IOGPUMemoryInfoClassOnce.Do(func() {
		_IOGPUMemoryInfoClass = IOGPUMemoryInfoClass{class: objc.GetClass("IOGPUMemoryInfo")}
	})
	return _IOGPUMemoryInfoClass
}

// GetIOGPUMemoryInfoClass returns the class object for IOGPUMemoryInfo.
func GetIOGPUMemoryInfoClass() IOGPUMemoryInfoClass {
	return getIOGPUMemoryInfoClass()
}

type IOGPUMemoryInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMemoryInfoClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMemoryInfoClass) Alloc() IOGPUMemoryInfo {
	rv := objc.SendIfResponds[IOGPUMemoryInfo](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMemoryInfo.AddDataSource]
//   - [IOGPUMemoryInfo.AddResourceToList]
//   - [IOGPUMemoryInfo.AnnotationList]
//   - [IOGPUMemoryInfo.GetAnnotationListAndEmitResourceInfo]
//   - [IOGPUMemoryInfo.Lock]
//   - [IOGPUMemoryInfo.RemoveDataSource]
//   - [IOGPUMemoryInfo.RemoveResourceFromList]
//   - [IOGPUMemoryInfo.Shutdown]
//   - [IOGPUMemoryInfo.Unlock]
type IOGPUMemoryInfo struct {
	objectivec.Object
}

// IOGPUMemoryInfoFromID constructs a [IOGPUMemoryInfo] from an objc.ID.
func IOGPUMemoryInfoFromID(id objc.ID) IOGPUMemoryInfo {
	return IOGPUMemoryInfo{objectivec.Object{ID: id}}
}

// Ensure IOGPUMemoryInfo implements IIOGPUMemoryInfo.
var _ IIOGPUMemoryInfo = IOGPUMemoryInfo{}

// An interface definition for the [IOGPUMemoryInfo] class.
//
// # Methods
//
//   - [IIOGPUMemoryInfo.AddDataSource]
//   - [IIOGPUMemoryInfo.AddResourceToList]
//   - [IIOGPUMemoryInfo.AnnotationList]
//   - [IIOGPUMemoryInfo.GetAnnotationListAndEmitResourceInfo]
//   - [IIOGPUMemoryInfo.Lock]
//   - [IIOGPUMemoryInfo.RemoveDataSource]
//   - [IIOGPUMemoryInfo.RemoveResourceFromList]
//   - [IIOGPUMemoryInfo.Shutdown]
//   - [IIOGPUMemoryInfo.Unlock]
type IIOGPUMemoryInfo interface {
	objectivec.IObject

	// Topic: Methods

	AddDataSource(source VoidHandler) unsafe.Pointer
	AddResourceToList(list objectivec.IObject)
	AnnotationList() corefoundation.CFArrayRef
	GetAnnotationListAndEmitResourceInfo() corefoundation.CFArrayRef
	Lock()
	RemoveDataSource(source unsafe.Pointer)
	RemoveResourceFromList(list objectivec.IObject)
	Shutdown()
	Unlock()
}

// Init initializes the instance.
func (i IOGPUMemoryInfo) Init() IOGPUMemoryInfo {
	rv := objc.SendIfResponds[IOGPUMemoryInfo](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMemoryInfo) Autorelease() IOGPUMemoryInfo {
	rv := objc.SendIfResponds[IOGPUMemoryInfo](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMemoryInfo creates a new IOGPUMemoryInfo instance.
func NewIOGPUMemoryInfo() IOGPUMemoryInfo {
	class := getIOGPUMemoryInfoClass()
	rv := objc.SendIfResponds[IOGPUMemoryInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (i IOGPUMemoryInfo) AddDataSource(source VoidHandler) unsafe.Pointer {
	_block0, _ := NewVoidBlock(source)
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("addDataSource:"), _block0)
	return unsafe.Pointer(rv)
}
func (i IOGPUMemoryInfo) AddResourceToList(list objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("addResourceToList:"), list)
}
func (i IOGPUMemoryInfo) AnnotationList() corefoundation.CFArrayRef {
	rv := objc.SendIfResponds[corefoundation.CFArrayRef](i.ID, objc.Sel("annotationList"))
	return corefoundation.CFArrayRef(rv)
}
func (i IOGPUMemoryInfo) GetAnnotationListAndEmitResourceInfo() corefoundation.CFArrayRef {
	rv := objc.SendIfResponds[corefoundation.CFArrayRef](i.ID, objc.Sel("getAnnotationListAndEmitResourceInfo"))
	return corefoundation.CFArrayRef(rv)
}
func (i IOGPUMemoryInfo) Lock() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("lock"))
}
func (i IOGPUMemoryInfo) RemoveDataSource(source unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeDataSource:"), source)
}
func (i IOGPUMemoryInfo) RemoveResourceFromList(list objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("removeResourceFromList:"), list)
}
func (i IOGPUMemoryInfo) Shutdown() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("shutdown"))
}
func (i IOGPUMemoryInfo) Unlock() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("unlock"))
}

// AddDataSourceSync is a synchronous wrapper around [IOGPUMemoryInfo.AddDataSource].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMemoryInfo) AddDataSourceSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	i.AddDataSource(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
