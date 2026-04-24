// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSpiceAgentCore] class.
var (
	_VZSpiceAgentCoreClass     VZSpiceAgentCoreClass
	_VZSpiceAgentCoreClassOnce sync.Once
)

func getVZSpiceAgentCoreClass() VZSpiceAgentCoreClass {
	_VZSpiceAgentCoreClassOnce.Do(func() {
		_VZSpiceAgentCoreClass = VZSpiceAgentCoreClass{class: objc.GetClass("_VZSpiceAgentCore")}
	})
	return _VZSpiceAgentCoreClass
}

// GetVZSpiceAgentCoreClass returns the class object for _VZSpiceAgentCore.
func GetVZSpiceAgentCoreClass() VZSpiceAgentCoreClass {
	return getVZSpiceAgentCoreClass()
}

type VZSpiceAgentCoreClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSpiceAgentCoreClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSpiceAgentCoreClass) Alloc() VZSpiceAgentCore {
	rv := objc.Send[VZSpiceAgentCore](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSpiceAgentCore.DidClosePort]
//   - [VZSpiceAgentCore.DidOpenPort]
//   - [VZSpiceAgentCore.IsValid]
//   - [VZSpiceAgentCore.PasteboardItemProvideDataForType]
//   - [VZSpiceAgentCore.Pause]
//   - [VZSpiceAgentCore.Resume]
//   - [VZSpiceAgentCore.Stop]
//   - [VZSpiceAgentCore.InitWithPasteboardQueueCapabilitiesInputOutput]
//   - [VZSpiceAgentCore.DebugDescription]
//   - [VZSpiceAgentCore.Description]
//   - [VZSpiceAgentCore.Hash]
//   - [VZSpiceAgentCore.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore
type VZSpiceAgentCore struct {
	objectivec.Object
}

// VZSpiceAgentCoreFromID constructs a [VZSpiceAgentCore] from an objc.ID.
func VZSpiceAgentCoreFromID(id objc.ID) VZSpiceAgentCore {
	return VZSpiceAgentCore{objectivec.Object{ID: id}}
}

// Ensure VZSpiceAgentCore implements IVZSpiceAgentCore.
var _ IVZSpiceAgentCore = VZSpiceAgentCore{}

// An interface definition for the [VZSpiceAgentCore] class.
//
// # Methods
//
//   - [IVZSpiceAgentCore.DidClosePort]
//   - [IVZSpiceAgentCore.DidOpenPort]
//   - [IVZSpiceAgentCore.IsValid]
//   - [IVZSpiceAgentCore.PasteboardItemProvideDataForType]
//   - [IVZSpiceAgentCore.Pause]
//   - [IVZSpiceAgentCore.Resume]
//   - [IVZSpiceAgentCore.Stop]
//   - [IVZSpiceAgentCore.InitWithPasteboardQueueCapabilitiesInputOutput]
//   - [IVZSpiceAgentCore.DebugDescription]
//   - [IVZSpiceAgentCore.Description]
//   - [IVZSpiceAgentCore.Hash]
//   - [IVZSpiceAgentCore.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore
type IVZSpiceAgentCore interface {
	objectivec.IObject

	// Topic: Methods

	DidClosePort()
	DidOpenPort()
	IsValid() bool
	PasteboardItemProvideDataForType(pasteboard objectivec.IObject, item objectivec.IObject, type_ objectivec.IObject)
	Pause()
	Resume()
	Stop()
	InitWithPasteboardQueueCapabilitiesInputOutput(pasteboard objectivec.IObject, queue DispatchQueue, capabilities objectivec.IObject, input unsafe.Pointer, output unsafe.Pointer) VZSpiceAgentCore
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZSpiceAgentCore) Init() VZSpiceAgentCore {
	rv := objc.Send[VZSpiceAgentCore](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSpiceAgentCore) Autorelease() VZSpiceAgentCore {
	rv := objc.Send[VZSpiceAgentCore](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSpiceAgentCore creates a new VZSpiceAgentCore instance.
func NewVZSpiceAgentCore() VZSpiceAgentCore {
	class := getVZSpiceAgentCoreClass()
	rv := objc.Send[VZSpiceAgentCore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/initWithPasteboard:queue:capabilities:input:output:
func NewVZSpiceAgentCoreWithPasteboardQueueCapabilitiesInputOutput(pasteboard objectivec.IObject, queue DispatchQueue, capabilities objectivec.IObject, input unsafe.Pointer, output unsafe.Pointer) VZSpiceAgentCore {
	instance := getVZSpiceAgentCoreClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboard:queue:capabilities:input:output:"), pasteboard, queue, capabilities, input, output)
	return VZSpiceAgentCoreFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/didClosePort
func (v VZSpiceAgentCore) DidClosePort() {
	objc.Send[objc.ID](v.ID, objc.Sel("didClosePort"))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/didOpenPort
func (v VZSpiceAgentCore) DidOpenPort() {
	objc.Send[objc.ID](v.ID, objc.Sel("didOpenPort"))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/isValid
func (v VZSpiceAgentCore) IsValid() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isValid"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/pasteboard:item:provideDataForType:
func (v VZSpiceAgentCore) PasteboardItemProvideDataForType(pasteboard objectivec.IObject, item objectivec.IObject, type_ objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("pasteboard:item:provideDataForType:"), pasteboard, item, type_)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/pause
func (v VZSpiceAgentCore) Pause() {
	objc.Send[objc.ID](v.ID, objc.Sel("pause"))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/resume
func (v VZSpiceAgentCore) Resume() {
	objc.Send[objc.ID](v.ID, objc.Sel("resume"))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/stop
func (v VZSpiceAgentCore) Stop() {
	objc.Send[objc.ID](v.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/initWithPasteboard:queue:capabilities:input:output:
func (v VZSpiceAgentCore) InitWithPasteboardQueueCapabilitiesInputOutput(pasteboard objectivec.IObject, queue DispatchQueue, capabilities objectivec.IObject, input unsafe.Pointer, output unsafe.Pointer) VZSpiceAgentCore {
	rv := objc.Send[VZSpiceAgentCore](v.ID, objc.Sel("initWithPasteboard:queue:capabilities:input:output:"), pasteboard, queue, capabilities, input, output)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/debugDescription
func (v VZSpiceAgentCore) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/description
func (v VZSpiceAgentCore) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/hash
func (v VZSpiceAgentCore) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSpiceAgentCore/superclass
func (v VZSpiceAgentCore) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
