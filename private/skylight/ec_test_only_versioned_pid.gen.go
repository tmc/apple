// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ECTestOnlyVersionedPID] class.
var (
	_ECTestOnlyVersionedPIDClass     ECTestOnlyVersionedPIDClass
	_ECTestOnlyVersionedPIDClassOnce sync.Once
)

func getECTestOnlyVersionedPIDClass() ECTestOnlyVersionedPIDClass {
	_ECTestOnlyVersionedPIDClassOnce.Do(func() {
		_ECTestOnlyVersionedPIDClass = ECTestOnlyVersionedPIDClass{class: objc.GetClass("ECTestOnlyVersionedPID")}
	})
	return _ECTestOnlyVersionedPIDClass
}

// GetECTestOnlyVersionedPIDClass returns the class object for ECTestOnlyVersionedPID.
func GetECTestOnlyVersionedPIDClass() ECTestOnlyVersionedPIDClass {
	return getECTestOnlyVersionedPIDClass()
}

type ECTestOnlyVersionedPIDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ECTestOnlyVersionedPIDClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ECTestOnlyVersionedPIDClass) Alloc() ECTestOnlyVersionedPID {
	rv := objc.Send[ECTestOnlyVersionedPID](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ECTestOnlyVersionedPID.Pid]
//   - [ECTestOnlyVersionedPID.Version]
//
// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyVersionedPID
type ECTestOnlyVersionedPID struct {
	objectivec.Object
}

// ECTestOnlyVersionedPIDFromID constructs a [ECTestOnlyVersionedPID] from an objc.ID.
func ECTestOnlyVersionedPIDFromID(id objc.ID) ECTestOnlyVersionedPID {
	return ECTestOnlyVersionedPID{objectivec.Object{ID: id}}
}

// Ensure ECTestOnlyVersionedPID implements IECTestOnlyVersionedPID.
var _ IECTestOnlyVersionedPID = ECTestOnlyVersionedPID{}

// An interface definition for the [ECTestOnlyVersionedPID] class.
//
// # Methods
//
//   - [IECTestOnlyVersionedPID.Pid]
//   - [IECTestOnlyVersionedPID.Version]
//
// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyVersionedPID
type IECTestOnlyVersionedPID interface {
	objectivec.IObject

	// Topic: Methods

	Pid() int
	Version() uint32
}

// Init initializes the instance.
func (e ECTestOnlyVersionedPID) Init() ECTestOnlyVersionedPID {
	rv := objc.Send[ECTestOnlyVersionedPID](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ECTestOnlyVersionedPID) Autorelease() ECTestOnlyVersionedPID {
	rv := objc.Send[ECTestOnlyVersionedPID](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewECTestOnlyVersionedPID creates a new ECTestOnlyVersionedPID instance.
func NewECTestOnlyVersionedPID() ECTestOnlyVersionedPID {
	class := getECTestOnlyVersionedPIDClass()
	rv := objc.Send[ECTestOnlyVersionedPID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyVersionedPID/pid
func (e ECTestOnlyVersionedPID) Pid() int {
	rv := objc.Send[int](e.ID, objc.Sel("pid"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyVersionedPID/version
func (e ECTestOnlyVersionedPID) Version() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("version"))
	return rv
}
