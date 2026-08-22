// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MEFileInfo] class.
var (
	_MEFileInfoClass     MEFileInfoClass
	_MEFileInfoClassOnce sync.Once
)

func getMEFileInfoClass() MEFileInfoClass {
	_MEFileInfoClassOnce.Do(func() {
		_MEFileInfoClass = MEFileInfoClass{class: objc.GetClass("MEFileInfo")}
	})
	return _MEFileInfoClass
}

// GetMEFileInfoClass returns the class object for MEFileInfo.
func GetMEFileInfoClass() MEFileInfoClass {
	return getMEFileInfoClass()
}

type MEFileInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MEFileInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MEFileInfoClass) Alloc() MEFileInfo {
	rv := objc.Send[MEFileInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains file properties from the media asset.
//
// # Inspecting file properties
//
//   - [MEFileInfo.Duration]: The duration of the media asset, if available.
//   - [MEFileInfo.SetDuration]
//   - [MEFileInfo.FragmentsStatus]: Indicates if the media asset contains fragments or is extendable by fragments.
//   - [MEFileInfo.SetFragmentsStatus]
//
// # Instance Properties
//
//   - [MEFileInfo.SidecarFileName]
//   - [MEFileInfo.SetSidecarFileName]
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFileInfo
type MEFileInfo struct {
	objectivec.Object
}

// MEFileInfoFromID constructs a [MEFileInfo] from an objc.ID.
//
// An object that contains file properties from the media asset.
func MEFileInfoFromID(id objc.ID) MEFileInfo {
	return MEFileInfo{objectivec.Object{ID: id}}
}

// NOTE: MEFileInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MEFileInfo] class.
//
// # Inspecting file properties
//
//   - [IMEFileInfo.Duration]: The duration of the media asset, if available.
//   - [IMEFileInfo.SetDuration]
//   - [IMEFileInfo.FragmentsStatus]: Indicates if the media asset contains fragments or is extendable by fragments.
//   - [IMEFileInfo.SetFragmentsStatus]
//
// # Instance Properties
//
//   - [IMEFileInfo.SidecarFileName]
//   - [IMEFileInfo.SetSidecarFileName]
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFileInfo
type IMEFileInfo interface {
	objectivec.IObject

	// Topic: Inspecting file properties

	// The duration of the media asset, if available.
	Duration() coremedia.CMTime
	SetDuration(value coremedia.CMTime)
	// Indicates if the media asset contains fragments or is extendable by fragments.
	FragmentsStatus() MEFileInfoFragmentsStatus
	SetFragmentsStatus(value MEFileInfoFragmentsStatus)

	// Topic: Instance Properties

	SidecarFileName() string
	SetSidecarFileName(value string)
}

// Init initializes the instance.
func (m MEFileInfo) Init() MEFileInfo {
	rv := objc.Send[MEFileInfo](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MEFileInfo) Autorelease() MEFileInfo {
	rv := objc.Send[MEFileInfo](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEFileInfo creates a new MEFileInfo instance.
func NewMEFileInfo() MEFileInfo {
	class := getMEFileInfoClass()
	rv := objc.Send[MEFileInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The duration of the media asset, if available.
//
// # Discussion
//
// This value is [invalid] if the duration isn’t available.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/duration
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (m MEFileInfo) Duration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](m.ID, objc.Sel("duration"))
	return coremedia.CMTime(rv)
}
func (m MEFileInfo) SetDuration(value coremedia.CMTime) {
	objc.Send[struct{}](m.ID, objc.Sel("setDuration:"), value)
}

// Indicates if the media asset contains fragments or is extendable by
// fragments.
//
// # Discussion
//
// The default value is [MEFileInfo.FragmentsStatus.couldNotContainFragments].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/fragmentsStatus-swift.property
//
// [MEFileInfo.FragmentsStatus.couldNotContainFragments]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/FragmentsStatus-swift.enum/couldNotContainFragments
func (m MEFileInfo) FragmentsStatus() MEFileInfoFragmentsStatus {
	rv := objc.Send[MEFileInfoFragmentsStatus](m.ID, objc.Sel("fragmentsStatus"))
	return MEFileInfoFragmentsStatus(rv)
}
func (m MEFileInfo) SetFragmentsStatus(value MEFileInfoFragmentsStatus) {
	objc.Send[struct{}](m.ID, objc.Sel("setFragmentsStatus:"), value)
}

// # Discussion
//
// The sidecar filename used by the MediaExtension.
//
// Represents a new or existing sidecar file located in the same directory as
// the primary media file. The filename should include the file extension, and
// should not contain the file path, or contain any slashes. The file
// extension should be supported by the format reader, and present in the
// EXAppExtensionAttributes and UTExportedTypeDeclarations dictionaries in the
// MediaExtension format reader Info.plist.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/sidecarFileName
func (m MEFileInfo) SidecarFileName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sidecarFileName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MEFileInfo) SetSidecarFileName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setSidecarFileName:"), objc.String(value))
}
