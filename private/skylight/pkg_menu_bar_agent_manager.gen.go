// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGMenuBarAgentManager] class.
var (
	_PKGMenuBarAgentManagerClass     PKGMenuBarAgentManagerClass
	_PKGMenuBarAgentManagerClassOnce sync.Once
)

func getPKGMenuBarAgentManagerClass() PKGMenuBarAgentManagerClass {
	_PKGMenuBarAgentManagerClassOnce.Do(func() {
		_PKGMenuBarAgentManagerClass = PKGMenuBarAgentManagerClass{class: objc.GetClass("PKGMenuBarAgentManager")}
	})
	return _PKGMenuBarAgentManagerClass
}

// GetPKGMenuBarAgentManagerClass returns the class object for PKGMenuBarAgentManager.
func GetPKGMenuBarAgentManagerClass() PKGMenuBarAgentManagerClass {
	return getPKGMenuBarAgentManagerClass()
}

type PKGMenuBarAgentManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGMenuBarAgentManagerClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGMenuBarAgentManagerClass) Alloc() PKGMenuBarAgentManager {
	rv := objc.SendIfResponds[PKGMenuBarAgentManager](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGMenuBarAgentManager.InitialUpdateData]
//   - [PKGMenuBarAgentManager.RebuildMetrics]
type PKGMenuBarAgentManager struct {
	objectivec.Object
}

// PKGMenuBarAgentManagerFromID constructs a [PKGMenuBarAgentManager] from an objc.ID.
func PKGMenuBarAgentManagerFromID(id objc.ID) PKGMenuBarAgentManager {
	return PKGMenuBarAgentManager{objectivec.Object{ID: id}}
}

// Ensure PKGMenuBarAgentManager implements IPKGMenuBarAgentManager.
var _ IPKGMenuBarAgentManager = PKGMenuBarAgentManager{}

// An interface definition for the [PKGMenuBarAgentManager] class.
//
// # Methods
//
//   - [IPKGMenuBarAgentManager.InitialUpdateData]
//   - [IPKGMenuBarAgentManager.RebuildMetrics]
type IPKGMenuBarAgentManager interface {
	objectivec.IObject

	// Topic: Methods

	InitialUpdateData() foundation.NSData
	RebuildMetrics()
}

// Init initializes the instance.
func (p PKGMenuBarAgentManager) Init() PKGMenuBarAgentManager {
	rv := objc.SendIfResponds[PKGMenuBarAgentManager](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGMenuBarAgentManager) Autorelease() PKGMenuBarAgentManager {
	rv := objc.SendIfResponds[PKGMenuBarAgentManager](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGMenuBarAgentManager creates a new PKGMenuBarAgentManager instance.
func NewPKGMenuBarAgentManager() PKGMenuBarAgentManager {
	class := getPKGMenuBarAgentManagerClass()
	rv := objc.SendIfResponds[PKGMenuBarAgentManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p PKGMenuBarAgentManager) RebuildMetrics() {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("rebuildMetrics"))
}

func (p PKGMenuBarAgentManager) InitialUpdateData() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](p.ID, objc.Sel("initialUpdateData"))
	return foundation.NSData(rv)
}
