// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Name] class.
var (
	_NameClass     NameClass
	_NameClassOnce sync.Once
)

func getNameClass() NameClass {
	_NameClassOnce.Do(func() {
		_NameClass = NameClass{class: objc.GetClass("name")}
	})
	return _NameClass
}

// GetNameClass returns the class object for name.
func GetNameClass() NameClass {
	return getNameClass()
}

type NameClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NameClass) Alloc() Name {
	rv := objc.Send[Name](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSException/name-c.ivar
type Name struct {
	objectivec.Object
}

// NameFromID constructs a [Name] from an objc.ID.
func NameFromID(id objc.ID) Name {
	return Name{objectivec.Object{ID: id}}
}
// Ensure Name implements IName.
var _ IName = Name{}

// An interface definition for the [Name] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/name-c.ivar
type IName interface {
	objectivec.IObject
}

// Init initializes the instance.
func (n Name) Init() Name {
	rv := objc.Send[Name](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n Name) Autorelease() Name {
	rv := objc.Send[Name](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewName creates a new Name instance.
func NewName() Name {
	class := getNameClass()
	rv := objc.Send[Name](objc.ID(class.class), objc.Sel("new"))
	return rv
}

