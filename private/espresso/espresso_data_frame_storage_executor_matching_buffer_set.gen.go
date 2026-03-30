// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoDataFrameStorageExecutorMatchingBufferSet] class.
var (
	_EspressoDataFrameStorageExecutorMatchingBufferSetClass     EspressoDataFrameStorageExecutorMatchingBufferSetClass
	_EspressoDataFrameStorageExecutorMatchingBufferSetClassOnce sync.Once
)

func getEspressoDataFrameStorageExecutorMatchingBufferSetClass() EspressoDataFrameStorageExecutorMatchingBufferSetClass {
	_EspressoDataFrameStorageExecutorMatchingBufferSetClassOnce.Do(func() {
		_EspressoDataFrameStorageExecutorMatchingBufferSetClass = EspressoDataFrameStorageExecutorMatchingBufferSetClass{class: objc.GetClass("EspressoDataFrameStorageExecutorMatchingBufferSet")}
	})
	return _EspressoDataFrameStorageExecutorMatchingBufferSetClass
}

// GetEspressoDataFrameStorageExecutorMatchingBufferSetClass returns the class object for EspressoDataFrameStorageExecutorMatchingBufferSet.
func GetEspressoDataFrameStorageExecutorMatchingBufferSetClass() EspressoDataFrameStorageExecutorMatchingBufferSetClass {
	return getEspressoDataFrameStorageExecutorMatchingBufferSetClass()
}

type EspressoDataFrameStorageExecutorMatchingBufferSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoDataFrameStorageExecutorMatchingBufferSetClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDataFrameStorageExecutorMatchingBufferSetClass) Alloc() EspressoDataFrameStorageExecutorMatchingBufferSet {
	rv := objc.Send[EspressoDataFrameStorageExecutorMatchingBufferSet](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.Computed_buffer]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.SetComputed_buffer]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.Computed_pb]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.SetComputed_pb]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.Groundtruth_buffer]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.SetGroundtruth_buffer]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.Index]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.SetIndex]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.Name]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.SetName]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.Reference_buffer]
//   - [EspressoDataFrameStorageExecutorMatchingBufferSet.SetReference_buffer]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutorMatchingBufferSet
type EspressoDataFrameStorageExecutorMatchingBufferSet struct {
	objectivec.Object
}

// EspressoDataFrameStorageExecutorMatchingBufferSetFromID constructs a [EspressoDataFrameStorageExecutorMatchingBufferSet] from an objc.ID.
func EspressoDataFrameStorageExecutorMatchingBufferSetFromID(id objc.ID) EspressoDataFrameStorageExecutorMatchingBufferSet {
	return EspressoDataFrameStorageExecutorMatchingBufferSet{objectivec.Object{ID: id}}
}

// Ensure EspressoDataFrameStorageExecutorMatchingBufferSet implements IEspressoDataFrameStorageExecutorMatchingBufferSet.
var _ IEspressoDataFrameStorageExecutorMatchingBufferSet = EspressoDataFrameStorageExecutorMatchingBufferSet{}

// An interface definition for the [EspressoDataFrameStorageExecutorMatchingBufferSet] class.
//
// # Methods
//
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.Computed_buffer]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.SetComputed_buffer]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.Computed_pb]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.SetComputed_pb]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.Groundtruth_buffer]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.SetGroundtruth_buffer]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.Index]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.SetIndex]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.Name]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.SetName]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.Reference_buffer]
//   - [IEspressoDataFrameStorageExecutorMatchingBufferSet.SetReference_buffer]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutorMatchingBufferSet
type IEspressoDataFrameStorageExecutorMatchingBufferSet interface {
	objectivec.IObject

	// Topic: Methods

	Computed_buffer() objectivec.IObject
	SetComputed_buffer(value objectivec.IObject)
	Computed_pb() corevideo.CVImageBufferRef
	SetComputed_pb(value corevideo.CVImageBufferRef)
	Groundtruth_buffer() objectivec.IObject
	SetGroundtruth_buffer(value objectivec.IObject)
	Index() uint64
	SetIndex(value uint64)
	Name() string
	SetName(value string)
	Reference_buffer() objectivec.IObject
	SetReference_buffer(value objectivec.IObject)
}

// Init initializes the instance.
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Init() EspressoDataFrameStorageExecutorMatchingBufferSet {
	rv := objc.Send[EspressoDataFrameStorageExecutorMatchingBufferSet](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Autorelease() EspressoDataFrameStorageExecutorMatchingBufferSet {
	rv := objc.Send[EspressoDataFrameStorageExecutorMatchingBufferSet](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDataFrameStorageExecutorMatchingBufferSet creates a new EspressoDataFrameStorageExecutorMatchingBufferSet instance.
func NewEspressoDataFrameStorageExecutorMatchingBufferSet() EspressoDataFrameStorageExecutorMatchingBufferSet {
	class := getEspressoDataFrameStorageExecutorMatchingBufferSetClass()
	rv := objc.Send[EspressoDataFrameStorageExecutorMatchingBufferSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutorMatchingBufferSet/computed_buffer
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Computed_buffer() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("computed_buffer"))
	return objectivec.Object{ID: rv}
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetComputed_buffer(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setComputed_buffer:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutorMatchingBufferSet/computed_pb
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Computed_pb() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](e.ID, objc.Sel("computed_pb"))
	return corevideo.CVImageBufferRef(rv)
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetComputed_pb(value corevideo.CVImageBufferRef) {
	objc.Send[struct{}](e.ID, objc.Sel("setComputed_pb:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutorMatchingBufferSet/groundtruth_buffer
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Groundtruth_buffer() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("groundtruth_buffer"))
	return objectivec.Object{ID: rv}
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetGroundtruth_buffer(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setGroundtruth_buffer:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutorMatchingBufferSet/index
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Index() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("index"))
	return rv
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetIndex(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setIndex:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutorMatchingBufferSet/name
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameStorageExecutorMatchingBufferSet/reference_buffer
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Reference_buffer() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("reference_buffer"))
	return objectivec.Object{ID: rv}
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetReference_buffer(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setReference_buffer:"), value)
}
