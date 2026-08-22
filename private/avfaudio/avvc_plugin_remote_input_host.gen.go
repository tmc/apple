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
	rv := objc.SendIfResponds[AVVCPluginRemoteInputHost](objc.ID(ac.class), objc.Sel("alloc"))
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
func (a AVVCPluginRemoteInputHost) Init() AVVCPluginRemoteInputHost {
	rv := objc.SendIfResponds[AVVCPluginRemoteInputHost](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCPluginRemoteInputHost) Autorelease() AVVCPluginRemoteInputHost {
	rv := objc.SendIfResponds[AVVCPluginRemoteInputHost](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCPluginRemoteInputHost creates a new AVVCPluginRemoteInputHost instance.
func NewAVVCPluginRemoteInputHost() AVVCPluginRemoteInputHost {
	class := getAVVCPluginRemoteInputHostClass()
	rv := objc.SendIfResponds[AVVCPluginRemoteInputHost](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVVCPluginRemoteInputHost) AllBundles(bundles []objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("allBundles:"), objectivec.IObjectSliceToNSArray(bundles))
	return objectivec.Object{ID: rv}
}
func (a AVVCPluginRemoteInputHost) FindDeviceWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("findDeviceWithIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
func (a AVVCPluginRemoteInputHost) FindFirstBluetoothDevice() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("findFirstBluetoothDevice"))
	return objectivec.Object{ID: rv}
}
func (a AVVCPluginRemoteInputHost) InputPluginDidPublishDevice(plugin objectivec.IObject, device objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("inputPlugin:didPublishDevice:"), plugin, device)
}
func (a AVVCPluginRemoteInputHost) InputPluginDidUnpublishDevice(plugin objectivec.IObject, device objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("inputPlugin:didUnpublishDevice:"), plugin, device)
}
func (a AVVCPluginRemoteInputHost) InvalidatePlugins() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("invalidatePlugins"))
}
func (a AVVCPluginRemoteInputHost) MockPluginEndpoint() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("mockPluginEndpoint"))
	return objectivec.Object{ID: rv}
}
func (a AVVCPluginRemoteInputHost) SetParentVoiceController(controller objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("setParentVoiceController:"), controller)
}

func (a AVVCPluginRemoteInputHost) MMotherController() IAVVoiceController {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("mMotherController"))
	return AVVoiceControllerFromID(objc.ID(rv))
}
func (a AVVCPluginRemoteInputHost) SetMMotherController(value IAVVoiceController) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setMMotherController:"), value)
}
