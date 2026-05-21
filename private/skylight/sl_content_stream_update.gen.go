// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLContentStreamUpdate] class.
var (
	_SLContentStreamUpdateClass     SLContentStreamUpdateClass
	_SLContentStreamUpdateClassOnce sync.Once
)

func getSLContentStreamUpdateClass() SLContentStreamUpdateClass {
	_SLContentStreamUpdateClassOnce.Do(func() {
		_SLContentStreamUpdateClass = SLContentStreamUpdateClass{class: objc.GetClass("SLContentStreamUpdate")}
	})
	return _SLContentStreamUpdateClass
}

// GetSLContentStreamUpdateClass returns the class object for SLContentStreamUpdate.
func GetSLContentStreamUpdateClass() SLContentStreamUpdateClass {
	return getSLContentStreamUpdateClass()
}

type SLContentStreamUpdateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLContentStreamUpdateClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLContentStreamUpdateClass) Alloc() SLContentStreamUpdate {
	rv := objc.Send[SLContentStreamUpdate](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLContentStreamUpdate.BoundingRect]
//   - [SLContentStreamUpdate.ContentRect]
//   - [SLContentStreamUpdate.ContentScale]
//   - [SLContentStreamUpdate.CopyRects]
//   - [SLContentStreamUpdate.CornerRadius]
//   - [SLContentStreamUpdate.DirtyRects]
//   - [SLContentStreamUpdate.DisplayResolution]
//   - [SLContentStreamUpdate.DisplayTime]
//   - [SLContentStreamUpdate.DropCount]
//   - [SLContentStreamUpdate.FrameSurface]
//   - [SLContentStreamUpdate.MetaData]
//   - [SLContentStreamUpdate.SetMetaData]
//   - [SLContentStreamUpdate.ScreenRect]
//   - [SLContentStreamUpdate.Status]
//   - [SLContentStreamUpdate.UpdateRef]
//   - [SLContentStreamUpdate.SetUpdateRef]
//   - [SLContentStreamUpdate.InitWithStatusDisplayTimeFrameSurfaceUpdateRef]
type SLContentStreamUpdate struct {
	objectivec.Object
}

// SLContentStreamUpdateFromID constructs a [SLContentStreamUpdate] from an objc.ID.
func SLContentStreamUpdateFromID(id objc.ID) SLContentStreamUpdate {
	return SLContentStreamUpdate{objectivec.Object{ID: id}}
}

// Ensure SLContentStreamUpdate implements ISLContentStreamUpdate.
var _ ISLContentStreamUpdate = SLContentStreamUpdate{}

// An interface definition for the [SLContentStreamUpdate] class.
//
// # Methods
//
//   - [ISLContentStreamUpdate.BoundingRect]
//   - [ISLContentStreamUpdate.ContentRect]
//   - [ISLContentStreamUpdate.ContentScale]
//   - [ISLContentStreamUpdate.CopyRects]
//   - [ISLContentStreamUpdate.CornerRadius]
//   - [ISLContentStreamUpdate.DirtyRects]
//   - [ISLContentStreamUpdate.DisplayResolution]
//   - [ISLContentStreamUpdate.DisplayTime]
//   - [ISLContentStreamUpdate.DropCount]
//   - [ISLContentStreamUpdate.FrameSurface]
//   - [ISLContentStreamUpdate.MetaData]
//   - [ISLContentStreamUpdate.SetMetaData]
//   - [ISLContentStreamUpdate.ScreenRect]
//   - [ISLContentStreamUpdate.Status]
//   - [ISLContentStreamUpdate.UpdateRef]
//   - [ISLContentStreamUpdate.SetUpdateRef]
//   - [ISLContentStreamUpdate.InitWithStatusDisplayTimeFrameSurfaceUpdateRef]
type ISLContentStreamUpdate interface {
	objectivec.IObject

	// Topic: Methods

	BoundingRect() corefoundation.CGRect
	ContentRect() corefoundation.CGRect
	ContentScale() float64
	CopyRects(rects int) objectivec.IObject
	CornerRadius() float64
	DirtyRects() foundation.INSArray
	DisplayResolution() float64
	DisplayTime() uint64
	DropCount() uint64
	FrameSurface() iosurface.IOSurface
	MetaData() foundation.INSDictionary
	SetMetaData(value foundation.INSDictionary)
	ScreenRect() corefoundation.CGRect
	Status() int
	UpdateRef() coregraphics.CGDisplayStreamUpdateRef
	SetUpdateRef(value coregraphics.CGDisplayStreamUpdateRef)
	InitWithStatusDisplayTimeFrameSurfaceUpdateRef(status int, time uint64, surface objectivec.IObject, ref coregraphics.CGDisplayStreamUpdateRef) SLContentStreamUpdate
}

// Init initializes the instance.
func (s SLContentStreamUpdate) Init() SLContentStreamUpdate {
	rv := objc.Send[SLContentStreamUpdate](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLContentStreamUpdate) Autorelease() SLContentStreamUpdate {
	rv := objc.Send[SLContentStreamUpdate](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLContentStreamUpdate creates a new SLContentStreamUpdate instance.
func NewSLContentStreamUpdate() SLContentStreamUpdate {
	class := getSLContentStreamUpdateClass()
	rv := objc.Send[SLContentStreamUpdate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLContentStreamUpdateWithStatusDisplayTimeFrameSurfaceUpdateRef(status int, time uint64, surface objectivec.IObject, ref coregraphics.CGDisplayStreamUpdateRef) SLContentStreamUpdate {
	instance := getSLContentStreamUpdateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStatus:displayTime:frameSurface:updateRef:"), status, time, surface, ref)
	return SLContentStreamUpdateFromID(rv)
}

func (s SLContentStreamUpdate) CopyRects(rects int) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("copyRects:"), rects)
	return objectivec.Object{ID: rv}
}
func (s SLContentStreamUpdate) InitWithStatusDisplayTimeFrameSurfaceUpdateRef(status int, time uint64, surface objectivec.IObject, ref coregraphics.CGDisplayStreamUpdateRef) SLContentStreamUpdate {
	rv := objc.Send[SLContentStreamUpdate](s.ID, objc.Sel("initWithStatus:displayTime:frameSurface:updateRef:"), status, time, surface, ref)
	return rv
}

func (s SLContentStreamUpdate) BoundingRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("boundingRect"))
	return corefoundation.CGRect(rv)
}
func (s SLContentStreamUpdate) ContentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("contentRect"))
	return corefoundation.CGRect(rv)
}
func (s SLContentStreamUpdate) ContentScale() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("contentScale"))
	return rv
}
func (s SLContentStreamUpdate) CornerRadius() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("cornerRadius"))
	return rv
}
func (s SLContentStreamUpdate) DirtyRects() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dirtyRects"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SLContentStreamUpdate) DisplayResolution() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("displayResolution"))
	return rv
}
func (s SLContentStreamUpdate) DisplayTime() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("displayTime"))
	return rv
}
func (s SLContentStreamUpdate) DropCount() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("dropCount"))
	return rv
}
func (s SLContentStreamUpdate) FrameSurface() iosurface.IOSurface {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("frameSurface"))
	return iosurface.IOSurfaceFromID(objc.ID(rv))
}
func (s SLContentStreamUpdate) MetaData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("metaData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SLContentStreamUpdate) SetMetaData(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setMetaData:"), value)
}
func (s SLContentStreamUpdate) ScreenRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("screenRect"))
	return corefoundation.CGRect(rv)
}
func (s SLContentStreamUpdate) Status() int {
	rv := objc.Send[int](s.ID, objc.Sel("status"))
	return rv
}
func (s SLContentStreamUpdate) UpdateRef() coregraphics.CGDisplayStreamUpdateRef {
	rv := objc.Send[coregraphics.CGDisplayStreamUpdateRef](s.ID, objc.Sel("updateRef"))
	return coregraphics.CGDisplayStreamUpdateRef(rv)
}
func (s SLContentStreamUpdate) SetUpdateRef(value coregraphics.CGDisplayStreamUpdateRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setUpdateRef:"), value)
}
