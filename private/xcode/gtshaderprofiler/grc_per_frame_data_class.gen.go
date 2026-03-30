// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GRCPerFrameDataClass] class.
var (
	_GRCPerFrameDataClassClass     GRCPerFrameDataClassClass
	_GRCPerFrameDataClassClassOnce sync.Once
)

func getGRCPerFrameDataClassClass() GRCPerFrameDataClassClass {
	_GRCPerFrameDataClassClassOnce.Do(func() {
		_GRCPerFrameDataClassClass = GRCPerFrameDataClassClass{class: objc.GetClass("GRCPerFrameDataClass")}
	})
	return _GRCPerFrameDataClassClass
}

// GetGRCPerFrameDataClassClass returns the class object for GRCPerFrameDataClass.
func GetGRCPerFrameDataClassClass() GRCPerFrameDataClassClass {
	return getGRCPerFrameDataClassClass()
}

type GRCPerFrameDataClassClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GRCPerFrameDataClassClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GRCPerFrameDataClassClass) Alloc() GRCPerFrameDataClass {
	rv := objc.Send[GRCPerFrameDataClass](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GRCPerFrameDataClass
type GRCPerFrameDataClass struct {
	objectivec.Object
}

// GRCPerFrameDataClassFromID constructs a [GRCPerFrameDataClass] from an objc.ID.
func GRCPerFrameDataClassFromID(id objc.ID) GRCPerFrameDataClass {
	return GRCPerFrameDataClass{objectivec.Object{ID: id}}
}

// Ensure GRCPerFrameDataClass implements IGRCPerFrameDataClass.
var _ IGRCPerFrameDataClass = GRCPerFrameDataClass{}

// An interface definition for the [GRCPerFrameDataClass] class.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GRCPerFrameDataClass
type IGRCPerFrameDataClass interface {
	objectivec.IObject
}

// Init initializes the instance.
func (g GRCPerFrameDataClass) Init() GRCPerFrameDataClass {
	rv := objc.Send[GRCPerFrameDataClass](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GRCPerFrameDataClass) Autorelease() GRCPerFrameDataClass {
	rv := objc.Send[GRCPerFrameDataClass](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGRCPerFrameDataClass creates a new GRCPerFrameDataClass instance.
func NewGRCPerFrameDataClass() GRCPerFrameDataClass {
	class := getGRCPerFrameDataClassClass()
	rv := objc.Send[GRCPerFrameDataClass](objc.ID(class.class), objc.Sel("new"))
	return rv
}
