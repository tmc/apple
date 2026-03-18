// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLArchitecture] class.
var (
	_MTLArchitectureClass     MTLArchitectureClass
	_MTLArchitectureClassOnce sync.Once
)

func getMTLArchitectureClass() MTLArchitectureClass {
	_MTLArchitectureClassOnce.Do(func() {
		_MTLArchitectureClass = MTLArchitectureClass{class: objc.GetClass("MTLArchitecture")}
	})
	return _MTLArchitectureClass
}

// GetMTLArchitectureClass returns the class object for MTLArchitecture.
func GetMTLArchitectureClass() MTLArchitectureClass {
	return getMTLArchitectureClass()
}

type MTLArchitectureClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLArchitectureClass) Alloc() MTLArchitecture {
	rv := objc.Send[MTLArchitecture](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that contains the architectural details of a GPU device.
//
// # Inspecting a GPU device’s architecture details
//
//   - [MTLArchitecture.Name]: The name of a GPU device’s architecture.
//
// See: https://developer.apple.com/documentation/Metal/MTLArchitecture
type MTLArchitecture struct {
	objectivec.Object
}

// MTLArchitectureFromID constructs a [MTLArchitecture] from an objc.ID.
//
// A class that contains the architectural details of a GPU device.
func MTLArchitectureFromID(id objc.ID) MTLArchitecture {
	return MTLArchitecture{objectivec.Object{ID: id}}
}
// NOTE: MTLArchitecture adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLArchitecture] class.
//
// # Inspecting a GPU device’s architecture details
//
//   - [IMTLArchitecture.Name]: The name of a GPU device’s architecture.
//
// See: https://developer.apple.com/documentation/Metal/MTLArchitecture
type IMTLArchitecture interface {
	objectivec.IObject

	// Topic: Inspecting a GPU device’s architecture details

	// The name of a GPU device’s architecture.
	Name() string

	// The architectural details of the GPU device.
	Architecture() IMTLArchitecture
	SetArchitecture(value IMTLArchitecture)
	// A Boolean value that indicates whether a GPU device doesn’t have a connection to a display.
	IsHeadless() bool
	SetIsHeadless(value bool)
	// A Boolean value that indicates whether the GPU lowers its performance to conserve energy.
	IsLowPower() bool
	SetIsLowPower(value bool)
	// A Boolean value that indicates whether the GPU is removable.
	IsRemovable() bool
	SetIsRemovable(value bool)
	// The physical location of the GPU relative to the system.
	Location() MTLDeviceLocation
	SetLocation(value MTLDeviceLocation)
	// A specific GPU position based on its general location.
	LocationNumber() int
	SetLocationNumber(value int)
	// The total number of GPUs in the peer group, if applicable.
	PeerCount() uint32
	SetPeerCount(value uint32)
	// The peer group ID the GPU belongs to, if applicable.
	PeerGroupID() uint64
	SetPeerGroupID(value uint64)
	// The unique identifier for a GPU in a peer group.
	PeerIndex() uint32
	SetPeerIndex(value uint32)
	// The GPU device’s registry identifier.
	RegistryID() uint64
	SetRegistryID(value uint64)
}

// Init initializes the instance.
func (a MTLArchitecture) Init() MTLArchitecture {
	rv := objc.Send[MTLArchitecture](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLArchitecture) Autorelease() MTLArchitecture {
	rv := objc.Send[MTLArchitecture](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLArchitecture creates a new MTLArchitecture instance.
func NewMTLArchitecture() MTLArchitecture {
	class := getMTLArchitectureClass()
	rv := objc.Send[MTLArchitecture](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The name of a GPU device’s architecture.
//
// # Discussion
// 
// The property’s value is equivalent to the output from the `metal-arch`
// command line tool on the same system.
// 
// Apps can use this property’s value to make decisions at runtime. For
// example, an app could retrieve a GPU-specific file from its developer’s
// content delivery network (CDN), such as a shader library or binary archive.
// See [Shader libraries] and [Shader library and archive creation] for more
// information.
//
// [Shader libraries]: https://developer.apple.com/documentation/Metal/shader-libraries
// [Shader library and archive creation]: https://developer.apple.com/documentation/Metal/shader-library-and-archive-creation
//
// See: https://developer.apple.com/documentation/Metal/MTLArchitecture/name
func (a MTLArchitecture) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The architectural details of the GPU device.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/architecture
func (a MTLArchitecture) Architecture() IMTLArchitecture {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("architecture"))
	return MTLArchitectureFromID(objc.ID(rv))
}
func (a MTLArchitecture) SetArchitecture(value IMTLArchitecture) {
	objc.Send[struct{}](a.ID, objc.Sel("setArchitecture:"), value)
}

// A Boolean value that indicates whether a GPU device doesn’t have a
// connection to a display.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/isheadless
func (a MTLArchitecture) IsHeadless() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("headless"))
	return rv
}
func (a MTLArchitecture) SetIsHeadless(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setHeadless:"), value)
}

// A Boolean value that indicates whether the GPU lowers its performance to
// conserve energy.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/islowpower
func (a MTLArchitecture) IsLowPower() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("lowPower"))
	return rv
}
func (a MTLArchitecture) SetIsLowPower(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setLowPower:"), value)
}

// A Boolean value that indicates whether the GPU is removable.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/isremovable
func (a MTLArchitecture) IsRemovable() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("removable"))
	return rv
}
func (a MTLArchitecture) SetIsRemovable(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setRemovable:"), value)
}

// The physical location of the GPU relative to the system.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/location
func (a MTLArchitecture) Location() MTLDeviceLocation {
	rv := objc.Send[MTLDeviceLocation](a.ID, objc.Sel("location"))
	return MTLDeviceLocation(rv)
}
func (a MTLArchitecture) SetLocation(value MTLDeviceLocation) {
	objc.Send[struct{}](a.ID, objc.Sel("setLocation:"), value)
}

// A specific GPU position based on its general location.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/locationnumber
func (a MTLArchitecture) LocationNumber() int {
	rv := objc.Send[int](a.ID, objc.Sel("locationNumber"))
	return rv
}
func (a MTLArchitecture) SetLocationNumber(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setLocationNumber:"), value)
}

// The total number of GPUs in the peer group, if applicable.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/peercount
func (a MTLArchitecture) PeerCount() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("peerCount"))
	return rv
}
func (a MTLArchitecture) SetPeerCount(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPeerCount:"), value)
}

// The peer group ID the GPU belongs to, if applicable.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/peergroupid
func (a MTLArchitecture) PeerGroupID() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("peerGroupID"))
	return rv
}
func (a MTLArchitecture) SetPeerGroupID(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setPeerGroupID:"), value)
}

// The unique identifier for a GPU in a peer group.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/peerindex
func (a MTLArchitecture) PeerIndex() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("peerIndex"))
	return rv
}
func (a MTLArchitecture) SetPeerIndex(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPeerIndex:"), value)
}

// The GPU device’s registry identifier.
//
// See: https://developer.apple.com/documentation/metal/mtldevice/registryid
func (a MTLArchitecture) RegistryID() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("registryID"))
	return rv
}
func (a MTLArchitecture) SetRegistryID(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setRegistryID:"), value)
}

