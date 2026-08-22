// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TwoNetsStyleTransfer] class.
var (
	_TwoNetsStyleTransferClass     TwoNetsStyleTransferClass
	_TwoNetsStyleTransferClassOnce sync.Once
)

func getTwoNetsStyleTransferClass() TwoNetsStyleTransferClass {
	_TwoNetsStyleTransferClassOnce.Do(func() {
		_TwoNetsStyleTransferClass = TwoNetsStyleTransferClass{class: objc.GetClass("TwoNetsStyleTransfer")}
	})
	return _TwoNetsStyleTransferClass
}

// GetTwoNetsStyleTransferClass returns the class object for TwoNetsStyleTransfer.
func GetTwoNetsStyleTransferClass() TwoNetsStyleTransferClass {
	return getTwoNetsStyleTransferClass()
}

type TwoNetsStyleTransferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TwoNetsStyleTransferClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TwoNetsStyleTransferClass) Alloc() TwoNetsStyleTransfer {
	rv := objc.SendIfResponds[TwoNetsStyleTransfer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TwoNetsStyleTransfer.ExecuteSyncWithImageSmallImage]
//   - [TwoNetsStyleTransfer.HeightBig]
//   - [TwoNetsStyleTransfer.HeightSmall]
//   - [TwoNetsStyleTransfer.LoadOutputName]
//   - [TwoNetsStyleTransfer.WidthBig]
//   - [TwoNetsStyleTransfer.WidthSmall]
type TwoNetsStyleTransfer struct {
	objectivec.Object
}

// TwoNetsStyleTransferFromID constructs a [TwoNetsStyleTransfer] from an objc.ID.
func TwoNetsStyleTransferFromID(id objc.ID) TwoNetsStyleTransfer {
	return TwoNetsStyleTransfer{objectivec.Object{ID: id}}
}

// Ensure TwoNetsStyleTransfer implements ITwoNetsStyleTransfer.
var _ ITwoNetsStyleTransfer = TwoNetsStyleTransfer{}

// An interface definition for the [TwoNetsStyleTransfer] class.
//
// # Methods
//
//   - [ITwoNetsStyleTransfer.ExecuteSyncWithImageSmallImage]
//   - [ITwoNetsStyleTransfer.HeightBig]
//   - [ITwoNetsStyleTransfer.HeightSmall]
//   - [ITwoNetsStyleTransfer.LoadOutputName]
//   - [ITwoNetsStyleTransfer.WidthBig]
//   - [ITwoNetsStyleTransfer.WidthSmall]
type ITwoNetsStyleTransfer interface {
	objectivec.IObject

	// Topic: Methods

	ExecuteSyncWithImageSmallImage(image corevideo.CVImageBufferRef, image2 corevideo.CVImageBufferRef) iosurface.IOSurfaceRef
	HeightBig() int
	HeightSmall() int
	LoadOutputName(load objectivec.IObject, name objectivec.IObject) int
	WidthBig() int
	WidthSmall() int
}

// Init initializes the instance.
func (t TwoNetsStyleTransfer) Init() TwoNetsStyleTransfer {
	rv := objc.SendIfResponds[TwoNetsStyleTransfer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TwoNetsStyleTransfer) Autorelease() TwoNetsStyleTransfer {
	rv := objc.SendIfResponds[TwoNetsStyleTransfer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTwoNetsStyleTransfer creates a new TwoNetsStyleTransfer instance.
func NewTwoNetsStyleTransfer() TwoNetsStyleTransfer {
	class := getTwoNetsStyleTransferClass()
	rv := objc.SendIfResponds[TwoNetsStyleTransfer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TwoNetsStyleTransfer) ExecuteSyncWithImageSmallImage(image corevideo.CVImageBufferRef, image2 corevideo.CVImageBufferRef) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](t.ID, objc.Sel("executeSyncWithImage:smallImage:"), image, image2)
	return iosurface.IOSurfaceRef(rv)
}
func (t TwoNetsStyleTransfer) HeightBig() int {
	rv := objc.SendIfResponds[int](t.ID, objc.Sel("heightBig"))
	return rv
}
func (t TwoNetsStyleTransfer) HeightSmall() int {
	rv := objc.SendIfResponds[int](t.ID, objc.Sel("heightSmall"))
	return rv
}
func (t TwoNetsStyleTransfer) LoadOutputName(load objectivec.IObject, name objectivec.IObject) int {
	rv := objc.SendIfResponds[int](t.ID, objc.Sel("load:outputName:"), load, name)
	return rv
}
func (t TwoNetsStyleTransfer) WidthBig() int {
	rv := objc.SendIfResponds[int](t.ID, objc.Sel("widthBig"))
	return rv
}
func (t TwoNetsStyleTransfer) WidthSmall() int {
	rv := objc.SendIfResponds[int](t.ID, objc.Sel("widthSmall"))
	return rv
}

func (_TwoNetsStyleTransferClass TwoNetsStyleTransferClass) SupportsANE() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_TwoNetsStyleTransferClass.class), objc.Sel("supportsANE"))
	return rv
}
