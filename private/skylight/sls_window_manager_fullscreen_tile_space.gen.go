// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSWindowManagerFullscreenTileSpace] class.
var (
	_SLSWindowManagerFullscreenTileSpaceClass     SLSWindowManagerFullscreenTileSpaceClass
	_SLSWindowManagerFullscreenTileSpaceClassOnce sync.Once
)

func getSLSWindowManagerFullscreenTileSpaceClass() SLSWindowManagerFullscreenTileSpaceClass {
	_SLSWindowManagerFullscreenTileSpaceClassOnce.Do(func() {
		_SLSWindowManagerFullscreenTileSpaceClass = SLSWindowManagerFullscreenTileSpaceClass{class: objc.GetClass("SLSWindowManagerFullscreenTileSpace")}
	})
	return _SLSWindowManagerFullscreenTileSpaceClass
}

// GetSLSWindowManagerFullscreenTileSpaceClass returns the class object for SLSWindowManagerFullscreenTileSpace.
func GetSLSWindowManagerFullscreenTileSpaceClass() SLSWindowManagerFullscreenTileSpaceClass {
	return getSLSWindowManagerFullscreenTileSpaceClass()
}

type SLSWindowManagerFullscreenTileSpaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSWindowManagerFullscreenTileSpaceClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSWindowManagerFullscreenTileSpaceClass) Alloc() SLSWindowManagerFullscreenTileSpace {
	rv := objc.SendIfResponds[SLSWindowManagerFullscreenTileSpace](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSWindowManagerFullscreenTileSpace.Rect]
//   - [SLSWindowManagerFullscreenTileSpace.SetRect]
//   - [SLSWindowManagerFullscreenTileSpace.SubspaceID]
//   - [SLSWindowManagerFullscreenTileSpace.SetSubspaceID]
//   - [SLSWindowManagerFullscreenTileSpace.WindowID]
//   - [SLSWindowManagerFullscreenTileSpace.SetWindowID]
type SLSWindowManagerFullscreenTileSpace struct {
	objectivec.Object
}

// SLSWindowManagerFullscreenTileSpaceFromID constructs a [SLSWindowManagerFullscreenTileSpace] from an objc.ID.
func SLSWindowManagerFullscreenTileSpaceFromID(id objc.ID) SLSWindowManagerFullscreenTileSpace {
	return SLSWindowManagerFullscreenTileSpace{objectivec.Object{ID: id}}
}

// Ensure SLSWindowManagerFullscreenTileSpace implements ISLSWindowManagerFullscreenTileSpace.
var _ ISLSWindowManagerFullscreenTileSpace = SLSWindowManagerFullscreenTileSpace{}

// An interface definition for the [SLSWindowManagerFullscreenTileSpace] class.
//
// # Methods
//
//   - [ISLSWindowManagerFullscreenTileSpace.Rect]
//   - [ISLSWindowManagerFullscreenTileSpace.SetRect]
//   - [ISLSWindowManagerFullscreenTileSpace.SubspaceID]
//   - [ISLSWindowManagerFullscreenTileSpace.SetSubspaceID]
//   - [ISLSWindowManagerFullscreenTileSpace.WindowID]
//   - [ISLSWindowManagerFullscreenTileSpace.SetWindowID]
type ISLSWindowManagerFullscreenTileSpace interface {
	objectivec.IObject

	// Topic: Methods

	Rect() corefoundation.CGRect
	SetRect(value corefoundation.CGRect)
	SubspaceID() uint64
	SetSubspaceID(value uint64)
	WindowID() uint32
	SetWindowID(value uint32)
}

// Init initializes the instance.
func (s SLSWindowManagerFullscreenTileSpace) Init() SLSWindowManagerFullscreenTileSpace {
	rv := objc.SendIfResponds[SLSWindowManagerFullscreenTileSpace](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSWindowManagerFullscreenTileSpace) Autorelease() SLSWindowManagerFullscreenTileSpace {
	rv := objc.SendIfResponds[SLSWindowManagerFullscreenTileSpace](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSWindowManagerFullscreenTileSpace creates a new SLSWindowManagerFullscreenTileSpace instance.
func NewSLSWindowManagerFullscreenTileSpace() SLSWindowManagerFullscreenTileSpace {
	class := getSLSWindowManagerFullscreenTileSpaceClass()
	rv := objc.SendIfResponds[SLSWindowManagerFullscreenTileSpace](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s SLSWindowManagerFullscreenTileSpace) Rect() corefoundation.CGRect {
	rv := objc.SendIfResponds[corefoundation.CGRect](s.ID, objc.Sel("rect"))
	return corefoundation.CGRect(rv)
}
func (s SLSWindowManagerFullscreenTileSpace) SetRect(value corefoundation.CGRect) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setRect:"), value)
}
func (s SLSWindowManagerFullscreenTileSpace) SubspaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("subspaceID"))
	return rv
}
func (s SLSWindowManagerFullscreenTileSpace) SetSubspaceID(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setSubspaceID:"), value)
}
func (s SLSWindowManagerFullscreenTileSpace) WindowID() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
func (s SLSWindowManagerFullscreenTileSpace) SetWindowID(value uint32) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setWindowID:"), value)
}
