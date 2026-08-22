// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [ICCameraFolder] class.
var (
	_ICCameraFolderClass     ICCameraFolderClass
	_ICCameraFolderClassOnce sync.Once
)

func getICCameraFolderClass() ICCameraFolderClass {
	_ICCameraFolderClassOnce.Do(func() {
		_ICCameraFolderClass = ICCameraFolderClass{class: objc.GetClass("ICCameraFolder")}
	})
	return _ICCameraFolderClass
}

// GetICCameraFolderClass returns the class object for ICCameraFolder.
func GetICCameraFolderClass() ICCameraFolderClass {
	return getICCameraFolderClass()
}

type ICCameraFolderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICCameraFolderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICCameraFolderClass) Alloc() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a folder on a camera.
//
// # Inspecting a Folder’s Contents
//
//   - [ICCameraFolder.Contents]: A list of items that this folder contains.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFolder
type ICCameraFolder struct {
	ICCameraItem
}

// ICCameraFolderFromID constructs a [ICCameraFolder] from an objc.ID.
//
// An object that represents a folder on a camera.
func ICCameraFolderFromID(id objc.ID) ICCameraFolder {
	return ICCameraFolder{ICCameraItem: ICCameraItemFromID(id)}
}

// NOTE: ICCameraFolder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICCameraFolder] class.
//
// # Inspecting a Folder’s Contents
//
//   - [IICCameraFolder.Contents]: A list of items that this folder contains.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFolder
type IICCameraFolder interface {
	IICCameraItem

	// Topic: Inspecting a Folder’s Contents

	// A list of items that this folder contains.
	Contents() []ICCameraItem
}

// Init initializes the instance.
func (c ICCameraFolder) Init() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c ICCameraFolder) Autorelease() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraFolder creates a new ICCameraFolder instance.
func NewICCameraFolder() ICCameraFolder {
	class := getICCameraFolderClass()
	rv := objc.Send[ICCameraFolder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A list of items that this folder contains.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFolder/contents
func (c ICCameraFolder) Contents() []ICCameraItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("contents"))
	return objc.ConvertSlice(rv, func(id objc.ID) ICCameraItem {
		return ICCameraItemFromID(id)
	})
}
