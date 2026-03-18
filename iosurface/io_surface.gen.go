// Code generated from Apple documentation for IOSurface. DO NOT EDIT.

package iosurface

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOSurface] class.
var (
	_IOSurfaceClass     IOSurfaceClass
	_IOSurfaceClassOnce sync.Once
)

func getIOSurfaceClass() IOSurfaceClass {
	_IOSurfaceClassOnce.Do(func() {
		_IOSurfaceClass = IOSurfaceClass{class: objc.GetClass("IOSurface")}
	})
	return _IOSurfaceClass
}

// GetIOSurfaceClass returns the class object for IOSurface.
func GetIOSurfaceClass() IOSurfaceClass {
	return getIOSurfaceClass()
}

type IOSurfaceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOSurfaceClass) Alloc() IOSurface {
	rv := objc.Send[IOSurface](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// Data type representing an IOSurface opaque object.
//
// # Initializers
//
//   - [IOSurface.InitWithProperties]
//
// # Instance Properties
//
//   - [IOSurface.AllocationSize]
//   - [IOSurface.AllowsPixelSizeCasting]
//   - [IOSurface.BaseAddress]
//   - [IOSurface.BytesPerElement]
//   - [IOSurface.BytesPerRow]
//   - [IOSurface.ElementHeight]
//   - [IOSurface.ElementWidth]
//   - [IOSurface.Height]
//   - [IOSurface.InUse]
//   - [IOSurface.LocalUseCount]
//   - [IOSurface.PixelFormat]
//   - [IOSurface.PlaneCount]
//   - [IOSurface.Seed]
//   - [IOSurface.Width]
//   - [IOSurface.SurfaceID]
//
// # Instance Methods
//
//   - [IOSurface.AllAttachments]
//   - [IOSurface.AttachmentForKey]
//   - [IOSurface.BaseAddressOfPlaneAtIndex]
//   - [IOSurface.BytesPerElementOfPlaneAtIndex]
//   - [IOSurface.BytesPerRowOfPlaneAtIndex]
//   - [IOSurface.DecrementUseCount]
//   - [IOSurface.ElementHeightOfPlaneAtIndex]
//   - [IOSurface.ElementWidthOfPlaneAtIndex]
//   - [IOSurface.HeightOfPlaneAtIndex]
//   - [IOSurface.IncrementUseCount]
//   - [IOSurface.LockWithOptionsSeed]
//   - [IOSurface.RemoveAllAttachments]
//   - [IOSurface.RemoveAttachmentForKey]
//   - [IOSurface.SetAllAttachments]
//   - [IOSurface.SetAttachmentForKey]
//   - [IOSurface.SetPurgeableOldState]
//   - [IOSurface.UnlockWithOptionsSeed]
//   - [IOSurface.WidthOfPlaneAtIndex]
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface
type IOSurface struct {
	objectivec.Object
}

// IOSurfaceFromID constructs a [IOSurface] from an objc.ID.
//
// Data type representing an IOSurface opaque object.
func IOSurfaceFromID(id objc.ID) IOSurface {
	return IOSurface{objectivec.Object{ID: id}}
}
// NOTE: IOSurface adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOSurface] class.
//
// # Initializers
//
//   - [IIOSurface.InitWithProperties]
//
// # Instance Properties
//
//   - [IIOSurface.AllocationSize]
//   - [IIOSurface.AllowsPixelSizeCasting]
//   - [IIOSurface.BaseAddress]
//   - [IIOSurface.BytesPerElement]
//   - [IIOSurface.BytesPerRow]
//   - [IIOSurface.ElementHeight]
//   - [IIOSurface.ElementWidth]
//   - [IIOSurface.Height]
//   - [IIOSurface.InUse]
//   - [IIOSurface.LocalUseCount]
//   - [IIOSurface.PixelFormat]
//   - [IIOSurface.PlaneCount]
//   - [IIOSurface.Seed]
//   - [IIOSurface.Width]
//   - [IIOSurface.SurfaceID]
//
// # Instance Methods
//
//   - [IIOSurface.AllAttachments]
//   - [IIOSurface.AttachmentForKey]
//   - [IIOSurface.BaseAddressOfPlaneAtIndex]
//   - [IIOSurface.BytesPerElementOfPlaneAtIndex]
//   - [IIOSurface.BytesPerRowOfPlaneAtIndex]
//   - [IIOSurface.DecrementUseCount]
//   - [IIOSurface.ElementHeightOfPlaneAtIndex]
//   - [IIOSurface.ElementWidthOfPlaneAtIndex]
//   - [IIOSurface.HeightOfPlaneAtIndex]
//   - [IIOSurface.IncrementUseCount]
//   - [IIOSurface.LockWithOptionsSeed]
//   - [IIOSurface.RemoveAllAttachments]
//   - [IIOSurface.RemoveAttachmentForKey]
//   - [IIOSurface.SetAllAttachments]
//   - [IIOSurface.SetAttachmentForKey]
//   - [IIOSurface.SetPurgeableOldState]
//   - [IIOSurface.UnlockWithOptionsSeed]
//   - [IIOSurface.WidthOfPlaneAtIndex]
//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface
type IIOSurface interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithProperties(properties foundation.INSDictionary) IOSurface

	// Topic: Instance Properties

	AllocationSize() int
	AllowsPixelSizeCasting() bool
	BaseAddress() unsafe.Pointer
	BytesPerElement() int
	BytesPerRow() int
	ElementHeight() int
	ElementWidth() int
	Height() int
	InUse() bool
	LocalUseCount() int32
	PixelFormat() uint32
	PlaneCount() uint
	Seed() uint32
	Width() int
	SurfaceID() uint32

	// Topic: Instance Methods

	AllAttachments() foundation.INSDictionary
	AttachmentForKey(key string) objectivec.IObject
	BaseAddressOfPlaneAtIndex(planeIndex uint)
	BytesPerElementOfPlaneAtIndex(planeIndex uint) int
	BytesPerRowOfPlaneAtIndex(planeIndex uint) int
	DecrementUseCount()
	ElementHeightOfPlaneAtIndex(planeIndex uint) int
	ElementWidthOfPlaneAtIndex(planeIndex uint) int
	HeightOfPlaneAtIndex(planeIndex uint) int
	IncrementUseCount()
	LockWithOptionsSeed(options IOSurfaceLockOptions, seed unsafe.Pointer) int32
	RemoveAllAttachments()
	RemoveAttachmentForKey(key string)
	SetAllAttachments(dict foundation.INSDictionary)
	SetAttachmentForKey(anObject objectivec.IObject, key string)
	SetPurgeableOldState(newState IOSurfacePurgeabilityState, oldState *IOSurfacePurgeabilityState) int32
	UnlockWithOptionsSeed(options IOSurfaceLockOptions, seed unsafe.Pointer) int32
	WidthOfPlaneAtIndex(planeIndex uint) int

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s IOSurface) Init() IOSurface {
	rv := objc.Send[IOSurface](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s IOSurface) Autorelease() IOSurface {
	rv := objc.Send[IOSurface](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOSurface creates a new IOSurface instance.
func NewIOSurface() IOSurface {
	class := getIOSurfaceClass()
	rv := objc.Send[IOSurface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/init(properties:)
func NewSurfaceWithProperties(properties foundation.INSDictionary) IOSurface {
	instance := getIOSurfaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProperties:"), properties)
	return IOSurfaceFromID(rv)
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/init(properties:)
func (s IOSurface) InitWithProperties(properties foundation.INSDictionary) IOSurface {
	rv := objc.Send[IOSurface](s.ID, objc.Sel("initWithProperties:"), properties)
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/allAttachments()
func (s IOSurface) AllAttachments() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("allAttachments"))
	return foundation.NSDictionaryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/attachment(forKey:)
func (s IOSurface) AttachmentForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("attachmentForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/baseAddressOfPlane(at:)
func (s IOSurface) BaseAddressOfPlaneAtIndex(planeIndex uint) {
	objc.Send[objc.ID](s.ID, objc.Sel("baseAddressOfPlaneAtIndex:"), planeIndex)
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/bytesPerElementOfPlane(at:)
func (s IOSurface) BytesPerElementOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s.ID, objc.Sel("bytesPerElementOfPlaneAtIndex:"), planeIndex)
	return rv
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/bytesPerRowOfPlane(at:)
func (s IOSurface) BytesPerRowOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s.ID, objc.Sel("bytesPerRowOfPlaneAtIndex:"), planeIndex)
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/decrementUseCount()
func (s IOSurface) DecrementUseCount() {
	objc.Send[objc.ID](s.ID, objc.Sel("decrementUseCount"))
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/elementHeightOfPlane(at:)
func (s IOSurface) ElementHeightOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s.ID, objc.Sel("elementHeightOfPlaneAtIndex:"), planeIndex)
	return rv
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/elementWidthOfPlane(at:)
func (s IOSurface) ElementWidthOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s.ID, objc.Sel("elementWidthOfPlaneAtIndex:"), planeIndex)
	return rv
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/heightOfPlane(at:)
func (s IOSurface) HeightOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s.ID, objc.Sel("heightOfPlaneAtIndex:"), planeIndex)
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/incrementUseCount()
func (s IOSurface) IncrementUseCount() {
	objc.Send[objc.ID](s.ID, objc.Sel("incrementUseCount"))
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/lock(options:seed:)
func (s IOSurface) LockWithOptionsSeed(options IOSurfaceLockOptions, seed unsafe.Pointer) int32 {
	rv := objc.Send[int32](s.ID, objc.Sel("lockWithOptions:seed:"), options, seed)
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/removeAllAttachments()
func (s IOSurface) RemoveAllAttachments() {
	objc.Send[objc.ID](s.ID, objc.Sel("removeAllAttachments"))
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/removeAttachment(forKey:)
func (s IOSurface) RemoveAttachmentForKey(key string) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeAttachmentForKey:"), objc.String(key))
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/setAllAttachments(_:)
func (s IOSurface) SetAllAttachments(dict foundation.INSDictionary) {
	objc.Send[objc.ID](s.ID, objc.Sel("setAllAttachments:"), dict)
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/setAttachment(_:forKey:)
func (s IOSurface) SetAttachmentForKey(anObject objectivec.IObject, key string) {
	objc.Send[objc.ID](s.ID, objc.Sel("setAttachment:forKey:"), anObject, objc.String(key))
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/setPurgeable(_:oldState:)
func (s IOSurface) SetPurgeableOldState(newState IOSurfacePurgeabilityState, oldState *IOSurfacePurgeabilityState) int32 {
	rv := objc.Send[int32](s.ID, objc.Sel("setPurgeable:oldState:"), newState, oldState)
	return rv
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/unlock(options:seed:)
func (s IOSurface) UnlockWithOptionsSeed(options IOSurfaceLockOptions, seed unsafe.Pointer) int32 {
	rv := objc.Send[int32](s.ID, objc.Sel("unlockWithOptions:seed:"), options, seed)
	return rv
}

//
// See: https://developer.apple.com/documentation/IOSurface/IOSurface/widthOfPlane(at:)
func (s IOSurface) WidthOfPlaneAtIndex(planeIndex uint) int {
	rv := objc.Send[int](s.ID, objc.Sel("widthOfPlaneAtIndex:"), planeIndex)
	return rv
}
func (s IOSurface) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/allocationSize
func (s IOSurface) AllocationSize() int {
	rv := objc.Send[int](s.ID, objc.Sel("allocationSize"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/allowsPixelSizeCasting
func (s IOSurface) AllowsPixelSizeCasting() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("allowsPixelSizeCasting"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/baseAddress
func (s IOSurface) BaseAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("baseAddress"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/bytesPerElement
func (s IOSurface) BytesPerElement() int {
	rv := objc.Send[int](s.ID, objc.Sel("bytesPerElement"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/bytesPerRow
func (s IOSurface) BytesPerRow() int {
	rv := objc.Send[int](s.ID, objc.Sel("bytesPerRow"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/elementHeight
func (s IOSurface) ElementHeight() int {
	rv := objc.Send[int](s.ID, objc.Sel("elementHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/elementWidth
func (s IOSurface) ElementWidth() int {
	rv := objc.Send[int](s.ID, objc.Sel("elementWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/height
func (s IOSurface) Height() int {
	rv := objc.Send[int](s.ID, objc.Sel("height"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/isInUse
func (s IOSurface) InUse() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isInUse"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/localUseCount
func (s IOSurface) LocalUseCount() int32 {
	rv := objc.Send[int32](s.ID, objc.Sel("localUseCount"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/pixelFormat
func (s IOSurface) PixelFormat() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("pixelFormat"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/planeCount
func (s IOSurface) PlaneCount() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("planeCount"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/seed
func (s IOSurface) Seed() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("seed"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/width
func (s IOSurface) Width() int {
	rv := objc.Send[int](s.ID, objc.Sel("width"))
	return rv
}

// See: https://developer.apple.com/documentation/IOSurface/IOSurface/surfaceID
func (s IOSurface) SurfaceID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("surfaceID"))
	return rv
}

