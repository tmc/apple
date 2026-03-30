// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Padding] class.
var (
	_PaddingClass     PaddingClass
	_PaddingClassOnce sync.Once
)

func getPaddingClass() PaddingClass {
	_PaddingClassOnce.Do(func() {
		_PaddingClass = PaddingClass{class: objc.GetClass("padding")}
	})
	return _PaddingClass
}

// GetPaddingClass returns the class object for padding.
func GetPaddingClass() PaddingClass {
	return getPaddingClass()
}

type PaddingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PaddingClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PaddingClass) Alloc() Padding {
	rv := objc.Send[Padding](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSXMLDocument/padding
type Padding struct {
	objectivec.Object
}

// PaddingFromID constructs a [Padding] from an objc.ID.
func PaddingFromID(id objc.ID) Padding {
	return Padding{objectivec.Object{ID: id}}
}

// Ensure Padding implements IPadding.
var _ IPadding = Padding{}

// An interface definition for the [Padding] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSXMLDocument/padding
type IPadding interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p Padding) Init() Padding {
	rv := objc.Send[Padding](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p Padding) Autorelease() Padding {
	rv := objc.Send[Padding](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPadding creates a new Padding instance.
func NewPadding() Padding {
	class := getPaddingClass()
	rv := objc.Send[Padding](objc.ID(class.class), objc.Sel("new"))
	return rv
}
