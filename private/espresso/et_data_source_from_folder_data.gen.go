// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETDataSourceFromFolderData] class.
var (
	_ETDataSourceFromFolderDataClass     ETDataSourceFromFolderDataClass
	_ETDataSourceFromFolderDataClassOnce sync.Once
)

func getETDataSourceFromFolderDataClass() ETDataSourceFromFolderDataClass {
	_ETDataSourceFromFolderDataClassOnce.Do(func() {
		_ETDataSourceFromFolderDataClass = ETDataSourceFromFolderDataClass{class: objc.GetClass("ETDataSourceFromFolderData")}
	})
	return _ETDataSourceFromFolderDataClass
}

// GetETDataSourceFromFolderDataClass returns the class object for ETDataSourceFromFolderData.
func GetETDataSourceFromFolderDataClass() ETDataSourceFromFolderDataClass {
	return getETDataSourceFromFolderDataClass()
}

type ETDataSourceFromFolderDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETDataSourceFromFolderDataClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETDataSourceFromFolderDataClass) Alloc() ETDataSourceFromFolderData {
	rv := objc.Send[ETDataSourceFromFolderData](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETDataSourceFromFolderData.BalanceClassesForTraining]
//   - [ETDataSourceFromFolderData.SetBalanceClassesForTraining]
//   - [ETDataSourceFromFolderData.BufferWithPath]
//   - [ETDataSourceFromFolderData.ClassNames]
//   - [ETDataSourceFromFolderData.SetClassNames]
//   - [ETDataSourceFromFolderData.DataPointAtIndex]
//   - [ETDataSourceFromFolderData.FolderToImages]
//   - [ETDataSourceFromFolderData.SetFolderToImages]
//   - [ETDataSourceFromFolderData.ImageFileNames]
//   - [ETDataSourceFromFolderData.SetImageFileNames]
//   - [ETDataSourceFromFolderData.ImagesDir]
//   - [ETDataSourceFromFolderData.SetImagesDir]
//   - [ETDataSourceFromFolderData.NumberOfClasses]
//   - [ETDataSourceFromFolderData.SetNumberOfClasses]
//   - [ETDataSourceFromFolderData.NumberOfDataPoints]
//   - [ETDataSourceFromFolderData.PathToClassIndex]
//   - [ETDataSourceFromFolderData.SetPathToClassIndex]
//   - [ETDataSourceFromFolderData.InitWithFolderBalanceClassesForTraining]
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData
type ETDataSourceFromFolderData struct {
	objectivec.Object
}

// ETDataSourceFromFolderDataFromID constructs a [ETDataSourceFromFolderData] from an objc.ID.
func ETDataSourceFromFolderDataFromID(id objc.ID) ETDataSourceFromFolderData {
	return ETDataSourceFromFolderData{objectivec.Object{ID: id}}
}
// Ensure ETDataSourceFromFolderData implements IETDataSourceFromFolderData.
var _ IETDataSourceFromFolderData = ETDataSourceFromFolderData{}

// An interface definition for the [ETDataSourceFromFolderData] class.
//
// # Methods
//
//   - [IETDataSourceFromFolderData.BalanceClassesForTraining]
//   - [IETDataSourceFromFolderData.SetBalanceClassesForTraining]
//   - [IETDataSourceFromFolderData.BufferWithPath]
//   - [IETDataSourceFromFolderData.ClassNames]
//   - [IETDataSourceFromFolderData.SetClassNames]
//   - [IETDataSourceFromFolderData.DataPointAtIndex]
//   - [IETDataSourceFromFolderData.FolderToImages]
//   - [IETDataSourceFromFolderData.SetFolderToImages]
//   - [IETDataSourceFromFolderData.ImageFileNames]
//   - [IETDataSourceFromFolderData.SetImageFileNames]
//   - [IETDataSourceFromFolderData.ImagesDir]
//   - [IETDataSourceFromFolderData.SetImagesDir]
//   - [IETDataSourceFromFolderData.NumberOfClasses]
//   - [IETDataSourceFromFolderData.SetNumberOfClasses]
//   - [IETDataSourceFromFolderData.NumberOfDataPoints]
//   - [IETDataSourceFromFolderData.PathToClassIndex]
//   - [IETDataSourceFromFolderData.SetPathToClassIndex]
//   - [IETDataSourceFromFolderData.InitWithFolderBalanceClassesForTraining]
//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData
type IETDataSourceFromFolderData interface {
	objectivec.IObject

	// Topic: Methods

	BalanceClassesForTraining() bool
	SetBalanceClassesForTraining(value bool)
	BufferWithPath(path objectivec.IObject) unsafe.Pointer
	ClassNames() foundation.INSArray
	SetClassNames(value foundation.INSArray)
	DataPointAtIndex(index int) objectivec.IObject
	FolderToImages() foundation.INSDictionary
	SetFolderToImages(value foundation.INSDictionary)
	ImageFileNames() foundation.INSArray
	SetImageFileNames(value foundation.INSArray)
	ImagesDir() string
	SetImagesDir(value string)
	NumberOfClasses() int
	SetNumberOfClasses(value int)
	NumberOfDataPoints() int
	PathToClassIndex() foundation.INSDictionary
	SetPathToClassIndex(value foundation.INSDictionary)
	InitWithFolderBalanceClassesForTraining(folder objectivec.IObject, training bool) ETDataSourceFromFolderData
}

// Init initializes the instance.
func (e ETDataSourceFromFolderData) Init() ETDataSourceFromFolderData {
	rv := objc.Send[ETDataSourceFromFolderData](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETDataSourceFromFolderData) Autorelease() ETDataSourceFromFolderData {
	rv := objc.Send[ETDataSourceFromFolderData](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETDataSourceFromFolderData creates a new ETDataSourceFromFolderData instance.
func NewETDataSourceFromFolderData() ETDataSourceFromFolderData {
	class := getETDataSourceFromFolderDataClass()
	rv := objc.Send[ETDataSourceFromFolderData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/initWithFolder:balanceClassesForTraining:
func NewETDataSourceFromFolderDataWithFolderBalanceClassesForTraining(folder objectivec.IObject, training bool) ETDataSourceFromFolderData {
	instance := getETDataSourceFromFolderDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFolder:balanceClassesForTraining:"), folder, training)
	return ETDataSourceFromFolderDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/bufferWithPath:
func (e ETDataSourceFromFolderData) BufferWithPath(path objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("bufferWithPath:"), path)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/dataPointAtIndex:
func (e ETDataSourceFromFolderData) DataPointAtIndex(index int) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("dataPointAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/numberOfDataPoints
func (e ETDataSourceFromFolderData) NumberOfDataPoints() int {
	rv := objc.Send[int](e.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/initWithFolder:balanceClassesForTraining:
func (e ETDataSourceFromFolderData) InitWithFolderBalanceClassesForTraining(folder objectivec.IObject, training bool) ETDataSourceFromFolderData {
	rv := objc.Send[ETDataSourceFromFolderData](e.ID, objc.Sel("initWithFolder:balanceClassesForTraining:"), folder, training)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/balanceClassesForTraining
func (e ETDataSourceFromFolderData) BalanceClassesForTraining() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("balanceClassesForTraining"))
	return rv
}
func (e ETDataSourceFromFolderData) SetBalanceClassesForTraining(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setBalanceClassesForTraining:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/classNames
func (e ETDataSourceFromFolderData) ClassNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("classNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETDataSourceFromFolderData) SetClassNames(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setClassNames:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/folderToImages
func (e ETDataSourceFromFolderData) FolderToImages() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("folderToImages"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e ETDataSourceFromFolderData) SetFolderToImages(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setFolderToImages:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/imageFileNames
func (e ETDataSourceFromFolderData) ImageFileNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("imageFileNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETDataSourceFromFolderData) SetImageFileNames(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setImageFileNames:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/imagesDir
func (e ETDataSourceFromFolderData) ImagesDir() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("imagesDir"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETDataSourceFromFolderData) SetImagesDir(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setImagesDir:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/numberOfClasses
func (e ETDataSourceFromFolderData) NumberOfClasses() int {
	rv := objc.Send[int](e.ID, objc.Sel("numberOfClasses"))
	return rv
}
func (e ETDataSourceFromFolderData) SetNumberOfClasses(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setNumberOfClasses:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETDataSourceFromFolderData/pathToClassIndex
func (e ETDataSourceFromFolderData) PathToClassIndex() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("pathToClassIndex"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e ETDataSourceFromFolderData) SetPathToClassIndex(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setPathToClassIndex:"), value)
}

