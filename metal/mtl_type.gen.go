// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLType] class.
var (
	_MTLTypeClass     MTLTypeClass
	_MTLTypeClassOnce sync.Once
)

func getMTLTypeClass() MTLTypeClass {
	_MTLTypeClassOnce.Do(func() {
		_MTLTypeClass = MTLTypeClass{class: objc.GetClass("MTLType")}
	})
	return _MTLTypeClass
}

// GetMTLTypeClass returns the class object for MTLType.
func GetMTLTypeClass() MTLTypeClass {
	return getMTLTypeClass()
}

type MTLTypeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTypeClass) Alloc() MTLType {
	rv := objc.Send[MTLType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a data type.
//
// # Identifying the data type
//
//   - [MTLType.DataType]: The data type of the function argument.
//
// See: https://developer.apple.com/documentation/Metal/MTLType
type MTLType struct {
	objectivec.Object
}

// MTLTypeFromID constructs a [MTLType] from an objc.ID.
//
// A description of a data type.
func MTLTypeFromID(id objc.ID) MTLType {
	return MTLType{objectivec.Object{ID: id}}
}
// NOTE: MTLType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLType] class.
//
// # Identifying the data type
//
//   - [IMTLType.DataType]: The data type of the function argument.
//
// See: https://developer.apple.com/documentation/Metal/MTLType
type IMTLType interface {
	objectivec.IObject

	// Topic: Identifying the data type

	// The data type of the function argument.
	DataType() MTLDataType
}

// Init initializes the instance.
func (t MTLType) Init() MTLType {
	rv := objc.Send[MTLType](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLType) Autorelease() MTLType {
	rv := objc.Send[MTLType](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLType creates a new MTLType instance.
func NewMTLType() MTLType {
	class := getMTLTypeClass()
	rv := objc.Send[MTLType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The data type of the function argument.
//
// See: https://developer.apple.com/documentation/Metal/MTLType/dataType
func (t MTLType) DataType() MTLDataType {
	rv := objc.Send[MTLDataType](t.ID, objc.Sel("dataType"))
	return MTLDataType(rv)
}

