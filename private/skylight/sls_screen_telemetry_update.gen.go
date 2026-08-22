// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSScreenTelemetryUpdate] class.
var (
	_SLSScreenTelemetryUpdateClass     SLSScreenTelemetryUpdateClass
	_SLSScreenTelemetryUpdateClassOnce sync.Once
)

func getSLSScreenTelemetryUpdateClass() SLSScreenTelemetryUpdateClass {
	_SLSScreenTelemetryUpdateClassOnce.Do(func() {
		_SLSScreenTelemetryUpdateClass = SLSScreenTelemetryUpdateClass{class: objc.GetClass("SLSScreenTelemetryUpdate")}
	})
	return _SLSScreenTelemetryUpdateClass
}

// GetSLSScreenTelemetryUpdateClass returns the class object for SLSScreenTelemetryUpdate.
func GetSLSScreenTelemetryUpdateClass() SLSScreenTelemetryUpdateClass {
	return getSLSScreenTelemetryUpdateClass()
}

type SLSScreenTelemetryUpdateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSScreenTelemetryUpdateClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSScreenTelemetryUpdateClass) Alloc() SLSScreenTelemetryUpdate {
	rv := objc.SendIfResponds[SLSScreenTelemetryUpdate](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSScreenTelemetryUpdate.Action]
//   - [SLSScreenTelemetryUpdate.Connection]
//   - [SLSScreenTelemetryUpdate.Error]
//   - [SLSScreenTelemetryUpdate.Snapshot]
//   - [SLSScreenTelemetryUpdate.InitWithActionConnectionErrorAndSnapshot]
type SLSScreenTelemetryUpdate struct {
	objectivec.Object
}

// SLSScreenTelemetryUpdateFromID constructs a [SLSScreenTelemetryUpdate] from an objc.ID.
func SLSScreenTelemetryUpdateFromID(id objc.ID) SLSScreenTelemetryUpdate {
	return SLSScreenTelemetryUpdate{objectivec.Object{ID: id}}
}

// Ensure SLSScreenTelemetryUpdate implements ISLSScreenTelemetryUpdate.
var _ ISLSScreenTelemetryUpdate = SLSScreenTelemetryUpdate{}

// An interface definition for the [SLSScreenTelemetryUpdate] class.
//
// # Methods
//
//   - [ISLSScreenTelemetryUpdate.Action]
//   - [ISLSScreenTelemetryUpdate.Connection]
//   - [ISLSScreenTelemetryUpdate.Error]
//   - [ISLSScreenTelemetryUpdate.Snapshot]
//   - [ISLSScreenTelemetryUpdate.InitWithActionConnectionErrorAndSnapshot]
type ISLSScreenTelemetryUpdate interface {
	objectivec.IObject

	// Topic: Methods

	Action() uint32
	Connection() ISLScreenTelemetryConnection
	Error() foundation.NSError
	Snapshot() ISLSScreenTelemetryResultsSnapshotDataWrapper
	InitWithActionConnectionErrorAndSnapshot(action uint32, connection objectivec.IObject, error_ objectivec.IObject, snapshot objectivec.IObject) SLSScreenTelemetryUpdate
}

// Init initializes the instance.
func (s SLSScreenTelemetryUpdate) Init() SLSScreenTelemetryUpdate {
	rv := objc.SendIfResponds[SLSScreenTelemetryUpdate](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSScreenTelemetryUpdate) Autorelease() SLSScreenTelemetryUpdate {
	rv := objc.SendIfResponds[SLSScreenTelemetryUpdate](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSScreenTelemetryUpdate creates a new SLSScreenTelemetryUpdate instance.
func NewSLSScreenTelemetryUpdate() SLSScreenTelemetryUpdate {
	class := getSLSScreenTelemetryUpdateClass()
	rv := objc.SendIfResponds[SLSScreenTelemetryUpdate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSScreenTelemetryUpdateWithActionConnectionErrorAndSnapshot(action uint32, connection objectivec.IObject, error_ objectivec.IObject, snapshot objectivec.IObject) SLSScreenTelemetryUpdate {
	instance := getSLSScreenTelemetryUpdateClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAction:connection:error:andSnapshot:"), action, connection, error_, snapshot)
	return SLSScreenTelemetryUpdateFromID(rv)
}

func (s SLSScreenTelemetryUpdate) InitWithActionConnectionErrorAndSnapshot(action uint32, connection objectivec.IObject, error_ objectivec.IObject, snapshot objectivec.IObject) SLSScreenTelemetryUpdate {
	rv := objc.SendIfResponds[SLSScreenTelemetryUpdate](s.ID, objc.Sel("initWithAction:connection:error:andSnapshot:"), action, connection, error_, snapshot)
	return rv
}

func (_SLSScreenTelemetryUpdateClass SLSScreenTelemetryUpdateClass) UpdateWithActionConnectionErrorAndSnapshot(action uint32, connection objectivec.IObject, error_ objectivec.IObject, snapshot objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLSScreenTelemetryUpdateClass.class), objc.Sel("updateWithAction:connection:error:andSnapshot:"), action, connection, error_, snapshot)
	return objectivec.Object{ID: rv}
}

func (s SLSScreenTelemetryUpdate) Action() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("action"))
	return rv
}
func (s SLSScreenTelemetryUpdate) Connection() ISLScreenTelemetryConnection {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("connection"))
	return SLScreenTelemetryConnectionFromID(objc.ID(rv))
}
func (s SLSScreenTelemetryUpdate) Error() foundation.NSError {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (s SLSScreenTelemetryUpdate) Snapshot() ISLSScreenTelemetryResultsSnapshotDataWrapper {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("snapshot"))
	return SLSScreenTelemetryResultsSnapshotDataWrapperFromID(objc.ID(rv))
}
