// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCPluginRemoteInputHost] class.
var (
	_AVVCPluginRemoteInputHostClass     AVVCPluginRemoteInputHostClass
	_AVVCPluginRemoteInputHostClassOnce sync.Once
)

func getAVVCPluginRemoteInputHostClass() AVVCPluginRemoteInputHostClass {
	_AVVCPluginRemoteInputHostClassOnce.Do(func() {
		_AVVCPluginRemoteInputHostClass = AVVCPluginRemoteInputHostClass{class: objc.GetClass("AVVCPluginRemoteInputHost")}
	})
	return _AVVCPluginRemoteInputHostClass
}

// GetAVVCPluginRemoteInputHostClass returns the class object for AVVCPluginRemoteInputHost.
func GetAVVCPluginRemoteInputHostClass() AVVCPluginRemoteInputHostClass {
	return getAVVCPluginRemoteInputHostClass()
}

type AVVCPluginRemoteInputHostClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCPluginRemoteInputHostClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCPluginRemoteInputHostClass) Alloc() AVVCPluginRemoteInputHost {
	rv := objc.Send[AVVCPluginRemoteInputHost](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCPluginRemoteInputHost.AllBundles]
//   - [AVVCPluginRemoteInputHost.FindDeviceWithIdentifier]
//   - [AVVCPluginRemoteInputHost.FindFirstBluetoothDevice]
//   - [AVVCPluginRemoteInputHost.InputPluginDidPublishDevice]
//   - [AVVCPluginRemoteInputHost.InputPluginDidUnpublishDevice]
//   - [AVVCPluginRemoteInputHost.InvalidatePlugins]
//   - [AVVCPluginRemoteInputHost.MMotherController]
//   - [AVVCPluginRemoteInputHost.SetMMotherController]
//   - [AVVCPluginRemoteInputHost.MockPluginEndpoint]
//   - [AVVCPluginRemoteInputHost.SetParentVoiceController]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost
type AVVCPluginRemoteInputHost struct {
	objectivec.Object
}

// AVVCPluginRemoteInputHostFromID constructs a [AVVCPluginRemoteInputHost] from an objc.ID.
func AVVCPluginRemoteInputHostFromID(id objc.ID) AVVCPluginRemoteInputHost {
	return AVVCPluginRemoteInputHost{objectivec.Object{ID: id}}
}

// Ensure AVVCPluginRemoteInputHost implements IAVVCPluginRemoteInputHost.
var _ IAVVCPluginRemoteInputHost = AVVCPluginRemoteInputHost{}

// An interface definition for the [AVVCPluginRemoteInputHost] class.
//
// # Methods
//
//   - [IAVVCPluginRemoteInputHost.AllBundles]
//   - [IAVVCPluginRemoteInputHost.FindDeviceWithIdentifier]
//   - [IAVVCPluginRemoteInputHost.FindFirstBluetoothDevice]
//   - [IAVVCPluginRemoteInputHost.InputPluginDidPublishDevice]
//   - [IAVVCPluginRemoteInputHost.InputPluginDidUnpublishDevice]
//   - [IAVVCPluginRemoteInputHost.InvalidatePlugins]
//   - [IAVVCPluginRemoteInputHost.MMotherController]
//   - [IAVVCPluginRemoteInputHost.SetMMotherController]
//   - [IAVVCPluginRemoteInputHost.MockPluginEndpoint]
//   - [IAVVCPluginRemoteInputHost.SetParentVoiceController]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost
type IAVVCPluginRemoteInputHost interface {
	objectivec.IObject

	// Topic: Methods

	AllBundles(bundles []objectivec.IObject) objectivec.IObject
	FindDeviceWithIdentifier(identifier objectivec.IObject) objectivec.IObject
	FindFirstBluetoothDevice() objectivec.IObject
	InputPluginDidPublishDevice(plugin objectivec.IObject, device objectivec.IObject)
	InputPluginDidUnpublishDevice(plugin objectivec.IObject, device objectivec.IObject)
	InvalidatePlugins()
	MMotherController() IAVVoiceController
	SetMMotherController(value IAVVoiceController)
	MockPluginEndpoint() objectivec.IObject
	SetParentVoiceController(controller objectivec.IObject)
}

// Init initializes the instance.
func (v AVVCPluginRemoteInputHost) Init() AVVCPluginRemoteInputHost {
	rv := objc.Send[AVVCPluginRemoteInputHost](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCPluginRemoteInputHost) Autorelease() AVVCPluginRemoteInputHost {
	rv := objc.Send[AVVCPluginRemoteInputHost](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCPluginRemoteInputHost creates a new AVVCPluginRemoteInputHost instance.
func NewAVVCPluginRemoteInputHost() AVVCPluginRemoteInputHost {
	class := getAVVCPluginRemoteInputHostClass()
	rv := objc.Send[AVVCPluginRemoteInputHost](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost/allBundles:
func (v AVVCPluginRemoteInputHost) AllBundles(bundles []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("allBundles:"), objectivec.IObjectSliceToNSArray(bundles))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost/findDeviceWithIdentifier:
func (v AVVCPluginRemoteInputHost) FindDeviceWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("findDeviceWithIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost/findFirstBluetoothDevice
func (v AVVCPluginRemoteInputHost) FindFirstBluetoothDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("findFirstBluetoothDevice"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost/inputPlugin:didPublishDevice:
func (v AVVCPluginRemoteInputHost) InputPluginDidPublishDevice(plugin objectivec.IObject, device objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("inputPlugin:didPublishDevice:"), plugin, device)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost/inputPlugin:didUnpublishDevice:
func (v AVVCPluginRemoteInputHost) InputPluginDidUnpublishDevice(plugin objectivec.IObject, device objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("inputPlugin:didUnpublishDevice:"), plugin, device)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost/invalidatePlugins
func (v AVVCPluginRemoteInputHost) InvalidatePlugins() {
	objc.Send[objc.ID](v.ID, objc.Sel("invalidatePlugins"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost/mockPluginEndpoint
func (v AVVCPluginRemoteInputHost) MockPluginEndpoint() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("mockPluginEndpoint"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost/setParentVoiceController:
func (v AVVCPluginRemoteInputHost) SetParentVoiceController(controller objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("setParentVoiceController:"), controller)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPluginRemoteInputHost/mMotherController
func (v AVVCPluginRemoteInputHost) MMotherController() IAVVoiceController {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("mMotherController"))
	return AVVoiceControllerFromID(objc.ID(rv))
}
func (v AVVCPluginRemoteInputHost) SetMMotherController(value IAVVoiceController) {
	objc.Send[struct{}](v.ID, objc.Sel("setMMotherController:"), value)
}
