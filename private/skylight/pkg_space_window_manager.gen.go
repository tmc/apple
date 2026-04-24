// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGSpaceWindowManager] class.
var (
	_PKGSpaceWindowManagerClass     PKGSpaceWindowManagerClass
	_PKGSpaceWindowManagerClassOnce sync.Once
)

func getPKGSpaceWindowManagerClass() PKGSpaceWindowManagerClass {
	_PKGSpaceWindowManagerClassOnce.Do(func() {
		_PKGSpaceWindowManagerClass = PKGSpaceWindowManagerClass{class: objc.GetClass("PKGSpaceWindowManager")}
	})
	return _PKGSpaceWindowManagerClass
}

// GetPKGSpaceWindowManagerClass returns the class object for PKGSpaceWindowManager.
func GetPKGSpaceWindowManagerClass() PKGSpaceWindowManagerClass {
	return getPKGSpaceWindowManagerClass()
}

type PKGSpaceWindowManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGSpaceWindowManagerClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGSpaceWindowManagerClass) Alloc() PKGSpaceWindowManager {
	rv := objc.Send[PKGSpaceWindowManager](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/PKGSpaceWindowManager
type PKGSpaceWindowManager struct {
	objectivec.Object
}

// PKGSpaceWindowManagerFromID constructs a [PKGSpaceWindowManager] from an objc.ID.
func PKGSpaceWindowManagerFromID(id objc.ID) PKGSpaceWindowManager {
	return PKGSpaceWindowManager{objectivec.Object{ID: id}}
}

// Ensure PKGSpaceWindowManager implements IPKGSpaceWindowManager.
var _ IPKGSpaceWindowManager = PKGSpaceWindowManager{}

// An interface definition for the [PKGSpaceWindowManager] class.
//
// See: https://developer.apple.com/documentation/SkyLight/PKGSpaceWindowManager
type IPKGSpaceWindowManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (g PKGSpaceWindowManager) Init() PKGSpaceWindowManager {
	rv := objc.Send[PKGSpaceWindowManager](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g PKGSpaceWindowManager) Autorelease() PKGSpaceWindowManager {
	rv := objc.Send[PKGSpaceWindowManager](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGSpaceWindowManager creates a new PKGSpaceWindowManager instance.
func NewPKGSpaceWindowManager() PKGSpaceWindowManager {
	class := getPKGSpaceWindowManagerClass()
	rv := objc.Send[PKGSpaceWindowManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}
