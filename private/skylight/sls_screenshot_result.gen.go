// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSScreenshotResult] class.
var (
	_SLSScreenshotResultClass     SLSScreenshotResultClass
	_SLSScreenshotResultClassOnce sync.Once
)

func getSLSScreenshotResultClass() SLSScreenshotResultClass {
	_SLSScreenshotResultClassOnce.Do(func() {
		_SLSScreenshotResultClass = SLSScreenshotResultClass{class: objc.GetClass("SLSScreenshotResult")}
	})
	return _SLSScreenshotResultClass
}

// GetSLSScreenshotResultClass returns the class object for SLSScreenshotResult.
func GetSLSScreenshotResultClass() SLSScreenshotResultClass {
	return getSLSScreenshotResultClass()
}

type SLSScreenshotResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSScreenshotResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSScreenshotResultClass) Alloc() SLSScreenshotResult {
	rv := objc.Send[SLSScreenshotResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSScreenshotResult.FrameSurfaceHDR]
//   - [SLSScreenshotResult.FrameSurfaceSDR]
//   - [SLSScreenshotResult.Status]
//   - [SLSScreenshotResult.InitWithStatusFrameSurfaceSDRFrameSurfaceHDR]
type SLSScreenshotResult struct {
	objectivec.Object
}

// SLSScreenshotResultFromID constructs a [SLSScreenshotResult] from an objc.ID.
func SLSScreenshotResultFromID(id objc.ID) SLSScreenshotResult {
	return SLSScreenshotResult{objectivec.Object{ID: id}}
}

// Ensure SLSScreenshotResult implements ISLSScreenshotResult.
var _ ISLSScreenshotResult = SLSScreenshotResult{}

// An interface definition for the [SLSScreenshotResult] class.
//
// # Methods
//
//   - [ISLSScreenshotResult.FrameSurfaceHDR]
//   - [ISLSScreenshotResult.FrameSurfaceSDR]
//   - [ISLSScreenshotResult.Status]
//   - [ISLSScreenshotResult.InitWithStatusFrameSurfaceSDRFrameSurfaceHDR]
type ISLSScreenshotResult interface {
	objectivec.IObject

	// Topic: Methods

	FrameSurfaceHDR() iosurface.IOSurface
	FrameSurfaceSDR() iosurface.IOSurface
	Status() int
	InitWithStatusFrameSurfaceSDRFrameSurfaceHDR(status int, sdr objectivec.IObject, hdr objectivec.IObject) SLSScreenshotResult
}

// Init initializes the instance.
func (s SLSScreenshotResult) Init() SLSScreenshotResult {
	rv := objc.Send[SLSScreenshotResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSScreenshotResult) Autorelease() SLSScreenshotResult {
	rv := objc.Send[SLSScreenshotResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSScreenshotResult creates a new SLSScreenshotResult instance.
func NewSLSScreenshotResult() SLSScreenshotResult {
	class := getSLSScreenshotResultClass()
	rv := objc.Send[SLSScreenshotResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSScreenshotResultWithStatusFrameSurfaceSDRFrameSurfaceHDR(status int, sdr objectivec.IObject, hdr objectivec.IObject) SLSScreenshotResult {
	instance := getSLSScreenshotResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStatus:frameSurfaceSDR:frameSurfaceHDR:"), status, sdr, hdr)
	return SLSScreenshotResultFromID(rv)
}

func (s SLSScreenshotResult) InitWithStatusFrameSurfaceSDRFrameSurfaceHDR(status int, sdr objectivec.IObject, hdr objectivec.IObject) SLSScreenshotResult {
	rv := objc.Send[SLSScreenshotResult](s.ID, objc.Sel("initWithStatus:frameSurfaceSDR:frameSurfaceHDR:"), status, sdr, hdr)
	return rv
}

func (s SLSScreenshotResult) FrameSurfaceHDR() iosurface.IOSurface {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("frameSurfaceHDR"))
	return iosurface.IOSurfaceFromID(objc.ID(rv))
}
func (s SLSScreenshotResult) FrameSurfaceSDR() iosurface.IOSurface {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("frameSurfaceSDR"))
	return iosurface.IOSurfaceFromID(objc.ID(rv))
}
func (s SLSScreenshotResult) Status() int {
	rv := objc.Send[int](s.ID, objc.Sel("status"))
	return rv
}
