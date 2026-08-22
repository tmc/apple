// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalResource] class.
var (
	_IOGPUMetalResourceClass     IOGPUMetalResourceClass
	_IOGPUMetalResourceClassOnce sync.Once
)

func getIOGPUMetalResourceClass() IOGPUMetalResourceClass {
	_IOGPUMetalResourceClassOnce.Do(func() {
		_IOGPUMetalResourceClass = IOGPUMetalResourceClass{class: objc.GetClass("IOGPUMetalResource")}
	})
	return _IOGPUMetalResourceClass
}

// GetIOGPUMetalResourceClass returns the class object for IOGPUMetalResource.
func GetIOGPUMetalResourceClass() IOGPUMetalResourceClass {
	return getIOGPUMetalResourceClass()
}

type IOGPUMetalResourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalResourceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalResourceClass) Alloc() IOGPUMetalResource {
	rv := objc.SendIfResponds[IOGPUMetalResource](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalResource._setLabel]
//   - [IOGPUMetalResource.AllocationID]
//   - [IOGPUMetalResource.AnnotateResource]
//   - [IOGPUMetalResource.AttachedResourceInfoTraceEmitter]
//   - [IOGPUMetalResource.SetAttachedResourceInfoTraceEmitter]
//   - [IOGPUMetalResource.CopyAnnotationDictionaryObj_key_nameObj_dict]
//   - [IOGPUMetalResource.CopyAnnotations]
//   - [IOGPUMetalResource.DoesAliasAllResourcesCount]
//   - [IOGPUMetalResource.DoesAliasAnyResourcesCount]
//   - [IOGPUMetalResource.DoesAliasResource]
//   - [IOGPUMetalResource.EmitResourceInfoTraceEvent]
//   - [IOGPUMetalResource.GpuAddress]
//   - [IOGPUMetalResource.IsComplete]
//   - [IOGPUMetalResource.IsPurgeable]
//   - [IOGPUMetalResource.IsWriteComplete]
//   - [IOGPUMetalResource.MetadataVirtualAddress]
//   - [IOGPUMetalResource.ProtectionOptions]
//   - [IOGPUMetalResource.ReleaseStrongDevice]
//   - [IOGPUMetalResource.ResourceID]
//   - [IOGPUMetalResource.ResourceRef]
//   - [IOGPUMetalResource.ResourceSize]
//   - [IOGPUMetalResource.ResponsibleProcess]
//   - [IOGPUMetalResource.SetResponsibleProcess]
//   - [IOGPUMetalResource.RetainedLabel]
//   - [IOGPUMetalResource.UnfilteredResourceOptions]
//   - [IOGPUMetalResource.VirtualAddress]
//   - [IOGPUMetalResource.WaitUntilComplete]
//   - [IOGPUMetalResource.WeakDevice]
//   - [IOGPUMetalResource.InitMemorylessDescriptor]
//   - [IOGPUMetalResource.InitStandinWithDevice]
//   - [IOGPUMetalResource.InitWithDeviceOptionsArgsArgsSize]
//   - [IOGPUMetalResource.InitWithDeviceRemoteStorageResourceOptionsArgsArgsSize]
//   - [IOGPUMetalResource.InitWithResource]
//   - [IOGPUMetalResource.DebugDescription]
//   - [IOGPUMetalResource.Description]
//   - [IOGPUMetalResource.Hash]
//   - [IOGPUMetalResource.Superclass]
type IOGPUMetalResource struct {
	metal.MTLResourceObject
}

// IOGPUMetalResourceFromID constructs a [IOGPUMetalResource] from an objc.ID.
func IOGPUMetalResourceFromID(id objc.ID) IOGPUMetalResource {
	return IOGPUMetalResource{MTLResourceObject: metal.MTLResourceObjectFromID(id)}
}

// Ensure IOGPUMetalResource implements IIOGPUMetalResource.
var _ IIOGPUMetalResource = IOGPUMetalResource{}

// An interface definition for the [IOGPUMetalResource] class.
//
// # Methods
//
//   - [IIOGPUMetalResource._setLabel]
//   - [IIOGPUMetalResource.AllocationID]
//   - [IIOGPUMetalResource.AnnotateResource]
//   - [IIOGPUMetalResource.AttachedResourceInfoTraceEmitter]
//   - [IIOGPUMetalResource.SetAttachedResourceInfoTraceEmitter]
//   - [IIOGPUMetalResource.CopyAnnotationDictionaryObj_key_nameObj_dict]
//   - [IIOGPUMetalResource.CopyAnnotations]
//   - [IIOGPUMetalResource.DoesAliasAllResourcesCount]
//   - [IIOGPUMetalResource.DoesAliasAnyResourcesCount]
//   - [IIOGPUMetalResource.DoesAliasResource]
//   - [IIOGPUMetalResource.EmitResourceInfoTraceEvent]
//   - [IIOGPUMetalResource.GpuAddress]
//   - [IIOGPUMetalResource.IsComplete]
//   - [IIOGPUMetalResource.IsPurgeable]
//   - [IIOGPUMetalResource.IsWriteComplete]
//   - [IIOGPUMetalResource.MetadataVirtualAddress]
//   - [IIOGPUMetalResource.ProtectionOptions]
//   - [IIOGPUMetalResource.ReleaseStrongDevice]
//   - [IIOGPUMetalResource.ResourceID]
//   - [IIOGPUMetalResource.ResourceRef]
//   - [IIOGPUMetalResource.ResourceSize]
//   - [IIOGPUMetalResource.ResponsibleProcess]
//   - [IIOGPUMetalResource.SetResponsibleProcess]
//   - [IIOGPUMetalResource.RetainedLabel]
//   - [IIOGPUMetalResource.UnfilteredResourceOptions]
//   - [IIOGPUMetalResource.VirtualAddress]
//   - [IIOGPUMetalResource.WaitUntilComplete]
//   - [IIOGPUMetalResource.WeakDevice]
//   - [IIOGPUMetalResource.InitMemorylessDescriptor]
//   - [IIOGPUMetalResource.InitStandinWithDevice]
//   - [IIOGPUMetalResource.InitWithDeviceOptionsArgsArgsSize]
//   - [IIOGPUMetalResource.InitWithDeviceRemoteStorageResourceOptionsArgsArgsSize]
//   - [IIOGPUMetalResource.InitWithResource]
//   - [IIOGPUMetalResource.DebugDescription]
//   - [IIOGPUMetalResource.Description]
//   - [IIOGPUMetalResource.Hash]
//   - [IIOGPUMetalResource.Superclass]
type IIOGPUMetalResource interface {
	metal.MTLResource

	// Topic: Methods

	_setLabel(label objectivec.IObject)
	AllocationID() uint64
	AnnotateResource(resource corefoundation.CFDictionaryRef)
	AttachedResourceInfoTraceEmitter() objectivec.IObject
	SetAttachedResourceInfoTraceEmitter(value objectivec.IObject)
	CopyAnnotationDictionaryObj_key_nameObj_dict(dictionary uint64, obj_key_name corefoundation.CFStringRef, obj_dict corefoundation.CFDictionaryRef) corefoundation.CFDictionaryRef
	CopyAnnotations() corefoundation.CFArrayRef
	DoesAliasAllResourcesCount(resources []objectivec.IObject, count uint64) bool
	DoesAliasAnyResourcesCount(resources []objectivec.IObject, count uint64) bool
	DoesAliasResource(resource objectivec.IObject) bool
	EmitResourceInfoTraceEvent()
	GpuAddress() uint64
	IsComplete() bool
	IsPurgeable() bool
	IsWriteComplete() bool
	MetadataVirtualAddress() unsafe.Pointer
	ProtectionOptions() uint64
	ReleaseStrongDevice()
	ResourceID() uint32
	ResourceRef() unsafe.Pointer
	ResourceSize() uint64
	ResponsibleProcess() int
	SetResponsibleProcess(value int)
	RetainedLabel() objectivec.IObject
	UnfilteredResourceOptions() uint64
	VirtualAddress() unsafe.Pointer
	WaitUntilComplete()
	WeakDevice() IIOGPUMetalDevice
	InitMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalResource
	InitStandinWithDevice(device objectivec.IObject) IOGPUMetalResource
	InitWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalResource
	InitWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalResource
	InitWithResource(resource objectivec.IObject) IOGPUMetalResource
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (i IOGPUMetalResource) Init() IOGPUMetalResource {
	rv := objc.SendIfResponds[IOGPUMetalResource](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalResource) Autorelease() IOGPUMetalResource {
	rv := objc.SendIfResponds[IOGPUMetalResource](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalResource creates a new IOGPUMetalResource instance.
func NewIOGPUMetalResource() IOGPUMetalResource {
	class := getIOGPUMetalResourceClass()
	rv := objc.SendIfResponds[IOGPUMetalResource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalResourceMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalResource {
	instance := getIOGPUMetalResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return IOGPUMetalResourceFromID(rv)
}

func NewGPUMetalResourceStandinWithDevice(device objectivec.IObject) IOGPUMetalResource {
	instance := getIOGPUMetalResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initStandinWithDevice:"), device)
	return IOGPUMetalResourceFromID(rv)
}

func NewGPUMetalResourceWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalResource {
	instance := getIOGPUMetalResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return IOGPUMetalResourceFromID(rv)
}

func NewGPUMetalResourceWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalResource {
	instance := getIOGPUMetalResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return IOGPUMetalResourceFromID(rv)
}

func NewGPUMetalResourceWithResource(resource objectivec.IObject) IOGPUMetalResource {
	instance := getIOGPUMetalResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return IOGPUMetalResourceFromID(rv)
}

func (i IOGPUMetalResource) _setLabel(label objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_setLabel:"), label)
}
func (i IOGPUMetalResource) AnnotateResource(resource corefoundation.CFDictionaryRef) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("annotateResource:"), resource)
}
func (i IOGPUMetalResource) CopyAnnotationDictionaryObj_key_nameObj_dict(dictionary uint64, obj_key_name corefoundation.CFStringRef, obj_dict corefoundation.CFDictionaryRef) corefoundation.CFDictionaryRef {
	rv := objc.SendIfResponds[corefoundation.CFDictionaryRef](i.ID, objc.Sel("copyAnnotationDictionary:obj_key_name:obj_dict:"), dictionary, obj_key_name, obj_dict)
	return corefoundation.CFDictionaryRef(rv)
}
func (i IOGPUMetalResource) CopyAnnotations() corefoundation.CFArrayRef {
	rv := objc.SendIfResponds[corefoundation.CFArrayRef](i.ID, objc.Sel("copyAnnotations"))
	return corefoundation.CFArrayRef(rv)
}
func (i IOGPUMetalResource) DoesAliasAllResourcesCount(resources []objectivec.IObject, count uint64) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("doesAliasAllResources:count:"), objc.CArray(resources), count)
	return rv
}
func (i IOGPUMetalResource) DoesAliasAnyResourcesCount(resources []objectivec.IObject, count uint64) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("doesAliasAnyResources:count:"), objc.CArray(resources), count)
	return rv
}
func (i IOGPUMetalResource) DoesAliasResource(resource objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("doesAliasResource:"), resource)
	return rv
}
func (i IOGPUMetalResource) EmitResourceInfoTraceEvent() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("emitResourceInfoTraceEvent"))
}
func (i IOGPUMetalResource) IsComplete() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isComplete"))
	return rv
}
func (i IOGPUMetalResource) IsPurgeable() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isPurgeable"))
	return rv
}
func (i IOGPUMetalResource) IsWriteComplete() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isWriteComplete"))
	return rv
}
func (i IOGPUMetalResource) ReleaseStrongDevice() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("releaseStrongDevice"))
}
func (i IOGPUMetalResource) RetainedLabel() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("retainedLabel"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalResource) WaitUntilComplete() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("waitUntilComplete"))
}
func (i IOGPUMetalResource) InitMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalResource {
	rv := objc.SendIfResponds[IOGPUMetalResource](i.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return rv
}
func (i IOGPUMetalResource) InitStandinWithDevice(device objectivec.IObject) IOGPUMetalResource {
	rv := objc.SendIfResponds[IOGPUMetalResource](i.ID, objc.Sel("initStandinWithDevice:"), device)
	return rv
}
func (i IOGPUMetalResource) InitWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalResource {
	rv := objc.SendIfResponds[IOGPUMetalResource](i.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return rv
}
func (i IOGPUMetalResource) InitWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalResource {
	rv := objc.SendIfResponds[IOGPUMetalResource](i.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return rv
}
func (i IOGPUMetalResource) InitWithResource(resource objectivec.IObject) IOGPUMetalResource {
	rv := objc.SendIfResponds[IOGPUMetalResource](i.ID, objc.Sel("initWithResource:"), resource)
	return rv
}

func (i IOGPUMetalResource) AllocationID() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("allocationID"))
	return rv
}
func (i IOGPUMetalResource) AttachedResourceInfoTraceEmitter() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("attachedResourceInfoTraceEmitter"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalResource) SetAttachedResourceInfoTraceEmitter(value objectivec.IObject) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setAttachedResourceInfoTraceEmitter:"), value)
}
func (i IOGPUMetalResource) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalResource) Description() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalResource) GpuAddress() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("gpuAddress"))
	return rv
}
func (i IOGPUMetalResource) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("hash"))
	return rv
}
func (i IOGPUMetalResource) MetadataVirtualAddress() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("metadataVirtualAddress"))
	return rv
}
func (i IOGPUMetalResource) ProtectionOptions() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("protectionOptions"))
	return rv
}
func (i IOGPUMetalResource) ResourceID() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("resourceID"))
	return rv
}
func (i IOGPUMetalResource) ResourceRef() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("resourceRef"))
	return rv
}
func (i IOGPUMetalResource) ResourceSize() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("resourceSize"))
	return rv
}
func (i IOGPUMetalResource) ResponsibleProcess() int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("responsibleProcess"))
	return rv
}
func (i IOGPUMetalResource) SetResponsibleProcess(value int) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setResponsibleProcess:"), value)
}
func (i IOGPUMetalResource) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](i.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (i IOGPUMetalResource) UnfilteredResourceOptions() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("unfilteredResourceOptions"))
	return rv
}
func (i IOGPUMetalResource) VirtualAddress() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("virtualAddress"))
	return rv
}
func (i IOGPUMetalResource) WeakDevice() IIOGPUMetalDevice {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("weakDevice"))
	return IOGPUMetalDeviceFromID(objc.ID(rv))
}
