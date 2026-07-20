// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

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
type IEspressoDataFrameStorageExecutorMatchingBufferSet interface {
	objectivec.IObject

	// Topic: Methods

	Computed_buffer() unsafe.Pointer
	SetComputed_buffer(value unsafe.Pointer)
	Computed_pb() corevideo.CVImageBufferRef
	SetComputed_pb(value corevideo.CVImageBufferRef)
	Groundtruth_buffer() unsafe.Pointer
	SetGroundtruth_buffer(value unsafe.Pointer)
	Index() uint64
	SetIndex(value uint64)
	Name() string
	SetName(value string)
	Reference_buffer() unsafe.Pointer
	SetReference_buffer(value unsafe.Pointer)
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

func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Computed_buffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("computed_buffer"))
	return rv
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetComputed_buffer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setComputed_buffer:"), value)
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Computed_pb() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](e.ID, objc.Sel("computed_pb"))
	return corevideo.CVImageBufferRef(rv)
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetComputed_pb(value corevideo.CVImageBufferRef) {
	objc.Send[struct{}](e.ID, objc.Sel("setComputed_pb:"), value)
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Groundtruth_buffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("groundtruth_buffer"))
	return rv
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetGroundtruth_buffer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setGroundtruth_buffer:"), value)
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Index() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("index"))
	return rv
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetIndex(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setIndex:"), value)
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setName:"), objc.String(value))
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) Reference_buffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("reference_buffer"))
	return rv
}
func (e EspressoDataFrameStorageExecutorMatchingBufferSet) SetReference_buffer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setReference_buffer:"), value)
}
