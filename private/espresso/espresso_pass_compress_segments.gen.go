// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_compress_segments] class.
var (
	_EspressoPass_compress_segmentsClass     EspressoPass_compress_segmentsClass
	_EspressoPass_compress_segmentsClassOnce sync.Once
)

func getEspressoPass_compress_segmentsClass() EspressoPass_compress_segmentsClass {
	_EspressoPass_compress_segmentsClassOnce.Do(func() {
		_EspressoPass_compress_segmentsClass = EspressoPass_compress_segmentsClass{class: objc.GetClass("EspressoPass_compress_segments")}
	})
	return _EspressoPass_compress_segmentsClass
}

// GetEspressoPass_compress_segmentsClass returns the class object for EspressoPass_compress_segments.
func GetEspressoPass_compress_segmentsClass() EspressoPass_compress_segmentsClass {
	return getEspressoPass_compress_segmentsClass()
}

type EspressoPass_compress_segmentsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_compress_segmentsClass) Alloc() EspressoPass_compress_segments {
	rv := objc.Send[EspressoPass_compress_segments](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_compress_segments
type EspressoPass_compress_segments struct {
	EspressoCustomPass
}

// EspressoPass_compress_segmentsFromID constructs a [EspressoPass_compress_segments] from an objc.ID.
func EspressoPass_compress_segmentsFromID(id objc.ID) EspressoPass_compress_segments {
	return EspressoPass_compress_segments{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_compress_segments implements IEspressoPass_compress_segments.
var _ IEspressoPass_compress_segments = EspressoPass_compress_segments{}

// An interface definition for the [EspressoPass_compress_segments] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_compress_segments
type IEspressoPass_compress_segments interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_compress_segments) Init() EspressoPass_compress_segments {
	rv := objc.Send[EspressoPass_compress_segments](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_compress_segments) Autorelease() EspressoPass_compress_segments {
	rv := objc.Send[EspressoPass_compress_segments](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_compress_segments creates a new EspressoPass_compress_segments instance.
func NewEspressoPass_compress_segments() EspressoPass_compress_segments {
	class := getEspressoPass_compress_segmentsClass()
	rv := objc.Send[EspressoPass_compress_segments](objc.ID(class.class), objc.Sel("new"))
	return rv
}

