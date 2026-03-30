// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSpiceAgent] class.
var (
	_VZSpiceAgentClass     VZSpiceAgentClass
	_VZSpiceAgentClassOnce sync.Once
)

func getVZSpiceAgentClass() VZSpiceAgentClass {
	_VZSpiceAgentClassOnce.Do(func() {
		_VZSpiceAgentClass = VZSpiceAgentClass{class: objc.GetClass("VZSpiceAgent")}
	})
	return _VZSpiceAgentClass
}

// GetVZSpiceAgentClass returns the class object for VZSpiceAgent.
func GetVZSpiceAgentClass() VZSpiceAgentClass {
	return getVZSpiceAgentClass()
}

type VZSpiceAgentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSpiceAgentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSpiceAgentClass) Alloc() VZSpiceAgent {
	rv := objc.Send[VZSpiceAgent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSpiceAgent.ConsoleDeviceDidClosePort]
//   - [VZSpiceAgent.ConsoleDeviceDidOpenPort]
//   - [VZSpiceAgent.DebugDescription]
//   - [VZSpiceAgent.Description]
//   - [VZSpiceAgent.Hash]
//   - [VZSpiceAgent.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgent
type VZSpiceAgent struct {
	objectivec.Object
}

// VZSpiceAgentFromID constructs a [VZSpiceAgent] from an objc.ID.
func VZSpiceAgentFromID(id objc.ID) VZSpiceAgent {
	return VZSpiceAgent{objectivec.Object{ID: id}}
}

// Ensure VZSpiceAgent implements IVZSpiceAgent.
var _ IVZSpiceAgent = VZSpiceAgent{}

// An interface definition for the [VZSpiceAgent] class.
//
// # Methods
//
//   - [IVZSpiceAgent.ConsoleDeviceDidClosePort]
//   - [IVZSpiceAgent.ConsoleDeviceDidOpenPort]
//   - [IVZSpiceAgent.DebugDescription]
//   - [IVZSpiceAgent.Description]
//   - [IVZSpiceAgent.Hash]
//   - [IVZSpiceAgent.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgent
type IVZSpiceAgent interface {
	objectivec.IObject

	// Topic: Methods

	ConsoleDeviceDidClosePort(device objectivec.IObject, port objectivec.IObject)
	ConsoleDeviceDidOpenPort(device objectivec.IObject, port objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s VZSpiceAgent) Init() VZSpiceAgent {
	rv := objc.Send[VZSpiceAgent](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSpiceAgent) Autorelease() VZSpiceAgent {
	rv := objc.Send[VZSpiceAgent](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSpiceAgent creates a new VZSpiceAgent instance.
func NewVZSpiceAgent() VZSpiceAgent {
	class := getVZSpiceAgentClass()
	rv := objc.Send[VZSpiceAgent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgent/consoleDevice:didClosePort:
func (s VZSpiceAgent) ConsoleDeviceDidClosePort(device objectivec.IObject, port objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("consoleDevice:didClosePort:"), device, port)
}

// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgent/consoleDevice:didOpenPort:
func (s VZSpiceAgent) ConsoleDeviceDidOpenPort(device objectivec.IObject, port objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("consoleDevice:didOpenPort:"), device, port)
}

// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgent/debugDescription
func (s VZSpiceAgent) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgent/description
func (s VZSpiceAgent) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgent/hash
func (s VZSpiceAgent) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgent/superclass
func (s VZSpiceAgent) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}
