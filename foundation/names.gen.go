// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Names] class.
var (
	_NamesClass     NamesClass
	_NamesClassOnce sync.Once
)

func getNamesClass() NamesClass {
	_NamesClassOnce.Do(func() {
		_NamesClass = NamesClass{class: objc.GetClass("names")}
	})
	return _NamesClass
}

// GetNamesClass returns the class object for names.
func GetNamesClass() NamesClass {
	return getNamesClass()
}

type NamesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NamesClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NamesClass) Alloc() Names {
	rv := objc.Send[Names](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSHost/names-c.ivar
type Names struct {
	objectivec.Object
}

// NamesFromID constructs a [Names] from an objc.ID.
func NamesFromID(id objc.ID) Names {
	return Names{objectivec.Object{ID: id}}
}

// Ensure Names implements INames.
var _ INames = Names{}

// An interface definition for the [Names] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSHost/names-c.ivar
type INames interface {
	objectivec.IObject
}

// Init initializes the instance.
func (n Names) Init() Names {
	rv := objc.Send[Names](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n Names) Autorelease() Names {
	rv := objc.Send[Names](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNames creates a new Names instance.
func NewNames() Names {
	class := getNamesClass()
	rv := objc.Send[Names](objc.ID(class.class), objc.Sel("new"))
	return rv
}
