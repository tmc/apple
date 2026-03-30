// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoDataFrameMappedFile] class.
var (
	_EspressoDataFrameMappedFileClass     EspressoDataFrameMappedFileClass
	_EspressoDataFrameMappedFileClassOnce sync.Once
)

func getEspressoDataFrameMappedFileClass() EspressoDataFrameMappedFileClass {
	_EspressoDataFrameMappedFileClassOnce.Do(func() {
		_EspressoDataFrameMappedFileClass = EspressoDataFrameMappedFileClass{class: objc.GetClass("EspressoDataFrameMappedFile")}
	})
	return _EspressoDataFrameMappedFileClass
}

// GetEspressoDataFrameMappedFileClass returns the class object for EspressoDataFrameMappedFile.
func GetEspressoDataFrameMappedFileClass() EspressoDataFrameMappedFileClass {
	return getEspressoDataFrameMappedFileClass()
}

type EspressoDataFrameMappedFileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoDataFrameMappedFileClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDataFrameMappedFileClass) Alloc() EspressoDataFrameMappedFile {
	rv := objc.Send[EspressoDataFrameMappedFile](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoDataFrameMappedFile.BasePtr]
//   - [EspressoDataFrameMappedFile.SetBasePtr]
//   - [EspressoDataFrameMappedFile.Path]
//   - [EspressoDataFrameMappedFile.SetPath]
//   - [EspressoDataFrameMappedFile.InitWithPath]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameMappedFile
type EspressoDataFrameMappedFile struct {
	objectivec.Object
}

// EspressoDataFrameMappedFileFromID constructs a [EspressoDataFrameMappedFile] from an objc.ID.
func EspressoDataFrameMappedFileFromID(id objc.ID) EspressoDataFrameMappedFile {
	return EspressoDataFrameMappedFile{objectivec.Object{ID: id}}
}

// Ensure EspressoDataFrameMappedFile implements IEspressoDataFrameMappedFile.
var _ IEspressoDataFrameMappedFile = EspressoDataFrameMappedFile{}

// An interface definition for the [EspressoDataFrameMappedFile] class.
//
// # Methods
//
//   - [IEspressoDataFrameMappedFile.BasePtr]
//   - [IEspressoDataFrameMappedFile.SetBasePtr]
//   - [IEspressoDataFrameMappedFile.Path]
//   - [IEspressoDataFrameMappedFile.SetPath]
//   - [IEspressoDataFrameMappedFile.InitWithPath]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameMappedFile
type IEspressoDataFrameMappedFile interface {
	objectivec.IObject

	// Topic: Methods

	BasePtr() string
	SetBasePtr(value string)
	Path() string
	SetPath(value string)
	InitWithPath(path objectivec.IObject) EspressoDataFrameMappedFile
}

// Init initializes the instance.
func (e EspressoDataFrameMappedFile) Init() EspressoDataFrameMappedFile {
	rv := objc.Send[EspressoDataFrameMappedFile](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDataFrameMappedFile) Autorelease() EspressoDataFrameMappedFile {
	rv := objc.Send[EspressoDataFrameMappedFile](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDataFrameMappedFile creates a new EspressoDataFrameMappedFile instance.
func NewEspressoDataFrameMappedFile() EspressoDataFrameMappedFile {
	class := getEspressoDataFrameMappedFileClass()
	rv := objc.Send[EspressoDataFrameMappedFile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameMappedFile/initWithPath:
func NewEspressoDataFrameMappedFileWithPath(path objectivec.IObject) EspressoDataFrameMappedFile {
	instance := getEspressoDataFrameMappedFileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPath:"), path)
	return EspressoDataFrameMappedFileFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameMappedFile/initWithPath:
func (e EspressoDataFrameMappedFile) InitWithPath(path objectivec.IObject) EspressoDataFrameMappedFile {
	rv := objc.Send[EspressoDataFrameMappedFile](e.ID, objc.Sel("initWithPath:"), path)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameMappedFile/basePtr
func (e EspressoDataFrameMappedFile) BasePtr() string {
	rv := objc.Send[*byte](e.ID, objc.Sel("basePtr"))
	return objc.GoString(rv)
}
func (e EspressoDataFrameMappedFile) SetBasePtr(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setBasePtr:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameMappedFile/path
func (e EspressoDataFrameMappedFile) Path() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("path"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoDataFrameMappedFile) SetPath(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setPath:"), objc.String(value))
}
