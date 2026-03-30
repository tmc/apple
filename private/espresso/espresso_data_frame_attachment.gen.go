// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoDataFrameAttachment] class.
var (
	_EspressoDataFrameAttachmentClass     EspressoDataFrameAttachmentClass
	_EspressoDataFrameAttachmentClassOnce sync.Once
)

func getEspressoDataFrameAttachmentClass() EspressoDataFrameAttachmentClass {
	_EspressoDataFrameAttachmentClassOnce.Do(func() {
		_EspressoDataFrameAttachmentClass = EspressoDataFrameAttachmentClass{class: objc.GetClass("EspressoDataFrameAttachment")}
	})
	return _EspressoDataFrameAttachmentClass
}

// GetEspressoDataFrameAttachmentClass returns the class object for EspressoDataFrameAttachment.
func GetEspressoDataFrameAttachmentClass() EspressoDataFrameAttachmentClass {
	return getEspressoDataFrameAttachmentClass()
}

type EspressoDataFrameAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoDataFrameAttachmentClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDataFrameAttachmentClass) Alloc() EspressoDataFrameAttachment {
	rv := objc.Send[EspressoDataFrameAttachment](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoDataFrameAttachment.Disabled]
//   - [EspressoDataFrameAttachment.SetDisabled]
//   - [EspressoDataFrameAttachment.FilePath]
//   - [EspressoDataFrameAttachment.SetFilePath]
//   - [EspressoDataFrameAttachment.LoadFromDictFrameStorage]
//   - [EspressoDataFrameAttachment.Offset]
//   - [EspressoDataFrameAttachment.SetOffset]
//   - [EspressoDataFrameAttachment.RawPointer]
//   - [EspressoDataFrameAttachment.SetRawPointer]
//   - [EspressoDataFrameAttachment.Size]
//   - [EspressoDataFrameAttachment.SetSize]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameAttachment
type EspressoDataFrameAttachment struct {
	objectivec.Object
}

// EspressoDataFrameAttachmentFromID constructs a [EspressoDataFrameAttachment] from an objc.ID.
func EspressoDataFrameAttachmentFromID(id objc.ID) EspressoDataFrameAttachment {
	return EspressoDataFrameAttachment{objectivec.Object{ID: id}}
}

// Ensure EspressoDataFrameAttachment implements IEspressoDataFrameAttachment.
var _ IEspressoDataFrameAttachment = EspressoDataFrameAttachment{}

// An interface definition for the [EspressoDataFrameAttachment] class.
//
// # Methods
//
//   - [IEspressoDataFrameAttachment.Disabled]
//   - [IEspressoDataFrameAttachment.SetDisabled]
//   - [IEspressoDataFrameAttachment.FilePath]
//   - [IEspressoDataFrameAttachment.SetFilePath]
//   - [IEspressoDataFrameAttachment.LoadFromDictFrameStorage]
//   - [IEspressoDataFrameAttachment.Offset]
//   - [IEspressoDataFrameAttachment.SetOffset]
//   - [IEspressoDataFrameAttachment.RawPointer]
//   - [IEspressoDataFrameAttachment.SetRawPointer]
//   - [IEspressoDataFrameAttachment.Size]
//   - [IEspressoDataFrameAttachment.SetSize]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameAttachment
type IEspressoDataFrameAttachment interface {
	objectivec.IObject

	// Topic: Methods

	Disabled() bool
	SetDisabled(value bool)
	FilePath() string
	SetFilePath(value string)
	LoadFromDictFrameStorage(dict objectivec.IObject, storage objectivec.IObject)
	Offset() uint64
	SetOffset(value uint64)
	RawPointer() unsafe.Pointer
	SetRawPointer(value unsafe.Pointer)
	Size() uint64
	SetSize(value uint64)
}

// Init initializes the instance.
func (e EspressoDataFrameAttachment) Init() EspressoDataFrameAttachment {
	rv := objc.Send[EspressoDataFrameAttachment](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDataFrameAttachment) Autorelease() EspressoDataFrameAttachment {
	rv := objc.Send[EspressoDataFrameAttachment](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDataFrameAttachment creates a new EspressoDataFrameAttachment instance.
func NewEspressoDataFrameAttachment() EspressoDataFrameAttachment {
	class := getEspressoDataFrameAttachmentClass()
	rv := objc.Send[EspressoDataFrameAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameAttachment/loadFromDict:frameStorage:
func (e EspressoDataFrameAttachment) LoadFromDictFrameStorage(dict objectivec.IObject, storage objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("loadFromDict:frameStorage:"), dict, storage)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameAttachment/fromDict:frameStorage:
func (_EspressoDataFrameAttachmentClass EspressoDataFrameAttachmentClass) FromDictFrameStorage(dict objectivec.IObject, storage objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_EspressoDataFrameAttachmentClass.class), objc.Sel("fromDict:frameStorage:"), dict, storage)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameAttachment/disabled
func (e EspressoDataFrameAttachment) Disabled() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("disabled"))
	return rv
}
func (e EspressoDataFrameAttachment) SetDisabled(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setDisabled:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameAttachment/filePath
func (e EspressoDataFrameAttachment) FilePath() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("filePath"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoDataFrameAttachment) SetFilePath(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setFilePath:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameAttachment/offset
func (e EspressoDataFrameAttachment) Offset() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("offset"))
	return rv
}
func (e EspressoDataFrameAttachment) SetOffset(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setOffset:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameAttachment/rawPointer
func (e EspressoDataFrameAttachment) RawPointer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("rawPointer"))
	return rv
}
func (e EspressoDataFrameAttachment) SetRawPointer(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setRawPointer:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameAttachment/size
func (e EspressoDataFrameAttachment) Size() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("size"))
	return rv
}
func (e EspressoDataFrameAttachment) SetSize(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setSize:"), value)
}
