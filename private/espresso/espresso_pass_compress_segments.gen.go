// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassCompressSegments] class.
var (
	_EspressoPassCompressSegmentsClass     EspressoPassCompressSegmentsClass
	_EspressoPassCompressSegmentsClassOnce sync.Once
)

func getEspressoPassCompressSegmentsClass() EspressoPassCompressSegmentsClass {
	_EspressoPassCompressSegmentsClassOnce.Do(func() {
		_EspressoPassCompressSegmentsClass = EspressoPassCompressSegmentsClass{class: objc.GetClass("EspressoPass_compress_segments")}
	})
	return _EspressoPassCompressSegmentsClass
}

// GetEspressoPassCompressSegmentsClass returns the class object for EspressoPass_compress_segments.
func GetEspressoPassCompressSegmentsClass() EspressoPassCompressSegmentsClass {
	return getEspressoPassCompressSegmentsClass()
}

type EspressoPassCompressSegmentsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassCompressSegmentsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassCompressSegmentsClass) Alloc() EspressoPassCompressSegments {
	rv := objc.SendIfResponds[EspressoPassCompressSegments](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassCompressSegments struct {
	EspressoCustomPass
}

// EspressoPassCompressSegmentsFromID constructs a [EspressoPassCompressSegments] from an objc.ID.
func EspressoPassCompressSegmentsFromID(id objc.ID) EspressoPassCompressSegments {
	return EspressoPassCompressSegments{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_compress_segmentsFromID is an alias for [EspressoPassCompressSegmentsFromID] for cross-framework compatibility.
func EspressoPass_compress_segmentsFromID(id objc.ID) EspressoPassCompressSegments {
	return EspressoPassCompressSegmentsFromID(id)
}

// Ensure EspressoPassCompressSegments implements IEspressoPassCompressSegments.
var _ IEspressoPassCompressSegments = EspressoPassCompressSegments{}

// An interface definition for the [EspressoPassCompressSegments] class.
type IEspressoPassCompressSegments interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassCompressSegments) Init() EspressoPassCompressSegments {
	rv := objc.SendIfResponds[EspressoPassCompressSegments](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassCompressSegments) Autorelease() EspressoPassCompressSegments {
	rv := objc.SendIfResponds[EspressoPassCompressSegments](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassCompressSegments creates a new EspressoPassCompressSegments instance.
func NewEspressoPassCompressSegments() EspressoPassCompressSegments {
	class := getEspressoPassCompressSegmentsClass()
	rv := objc.SendIfResponds[EspressoPassCompressSegments](objc.ID(class.class), objc.Sel("new"))
	return rv
}
