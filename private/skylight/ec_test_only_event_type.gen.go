// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ECTestOnlyEventType] class.
var (
	_ECTestOnlyEventTypeClass     ECTestOnlyEventTypeClass
	_ECTestOnlyEventTypeClassOnce sync.Once
)

func getECTestOnlyEventTypeClass() ECTestOnlyEventTypeClass {
	_ECTestOnlyEventTypeClassOnce.Do(func() {
		_ECTestOnlyEventTypeClass = ECTestOnlyEventTypeClass{class: objc.GetClass("ECTestOnlyEventType")}
	})
	return _ECTestOnlyEventTypeClass
}

// GetECTestOnlyEventTypeClass returns the class object for ECTestOnlyEventType.
func GetECTestOnlyEventTypeClass() ECTestOnlyEventTypeClass {
	return getECTestOnlyEventTypeClass()
}

type ECTestOnlyEventTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ECTestOnlyEventTypeClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ECTestOnlyEventTypeClass) Alloc() ECTestOnlyEventType {
	rv := objc.Send[ECTestOnlyEventType](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ECTestOnlyEventType.CgSubType]
//   - [ECTestOnlyEventType.CgType]
//   - [ECTestOnlyEventType.HidType]
//   - [ECTestOnlyEventType.IsCGType]
//
// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventType
type ECTestOnlyEventType struct {
	objectivec.Object
}

// ECTestOnlyEventTypeFromID constructs a [ECTestOnlyEventType] from an objc.ID.
func ECTestOnlyEventTypeFromID(id objc.ID) ECTestOnlyEventType {
	return ECTestOnlyEventType{objectivec.Object{ID: id}}
}

// Ensure ECTestOnlyEventType implements IECTestOnlyEventType.
var _ IECTestOnlyEventType = ECTestOnlyEventType{}

// An interface definition for the [ECTestOnlyEventType] class.
//
// # Methods
//
//   - [IECTestOnlyEventType.CgSubType]
//   - [IECTestOnlyEventType.CgType]
//   - [IECTestOnlyEventType.HidType]
//   - [IECTestOnlyEventType.IsCGType]
//
// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventType
type IECTestOnlyEventType interface {
	objectivec.IObject

	// Topic: Methods

	CgSubType() uint64
	CgType() uint32
	HidType() uint32
	IsCGType() bool
}

// Init initializes the instance.
func (e ECTestOnlyEventType) Init() ECTestOnlyEventType {
	rv := objc.Send[ECTestOnlyEventType](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ECTestOnlyEventType) Autorelease() ECTestOnlyEventType {
	rv := objc.Send[ECTestOnlyEventType](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewECTestOnlyEventType creates a new ECTestOnlyEventType instance.
func NewECTestOnlyEventType() ECTestOnlyEventType {
	class := getECTestOnlyEventTypeClass()
	rv := objc.Send[ECTestOnlyEventType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventType/cgSubType
func (e ECTestOnlyEventType) CgSubType() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("cgSubType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventType/cgType
func (e ECTestOnlyEventType) CgType() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("cgType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventType/hidType
func (e ECTestOnlyEventType) HidType() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("hidType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/ECTestOnlyEventType/isCGType
func (e ECTestOnlyEventType) IsCGType() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isCGType"))
	return rv
}
