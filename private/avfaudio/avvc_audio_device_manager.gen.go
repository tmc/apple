// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCAudioDeviceManager] class.
var (
	_AVVCAudioDeviceManagerClass     AVVCAudioDeviceManagerClass
	_AVVCAudioDeviceManagerClassOnce sync.Once
)

func getAVVCAudioDeviceManagerClass() AVVCAudioDeviceManagerClass {
	_AVVCAudioDeviceManagerClassOnce.Do(func() {
		_AVVCAudioDeviceManagerClass = AVVCAudioDeviceManagerClass{class: objc.GetClass("AVVCAudioDeviceManager")}
	})
	return _AVVCAudioDeviceManagerClass
}

// GetAVVCAudioDeviceManagerClass returns the class object for AVVCAudioDeviceManager.
func GetAVVCAudioDeviceManagerClass() AVVCAudioDeviceManagerClass {
	return getAVVCAudioDeviceManagerClass()
}

type AVVCAudioDeviceManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCAudioDeviceManagerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCAudioDeviceManagerClass) Alloc() AVVCAudioDeviceManager {
	rv := objc.Send[AVVCAudioDeviceManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCAudioDeviceManager.InputDeviceID]
//   - [AVVCAudioDeviceManager.OutputDeviceID]
//   - [AVVCAudioDeviceManager.SelectDeviceForActivationMode]
//   - [AVVCAudioDeviceManager.InitWithActivationContextWithError]
//   - [AVVCAudioDeviceManager.InitWithError]
type AVVCAudioDeviceManager struct {
	objectivec.Object
}

// AVVCAudioDeviceManagerFromID constructs a [AVVCAudioDeviceManager] from an objc.ID.
func AVVCAudioDeviceManagerFromID(id objc.ID) AVVCAudioDeviceManager {
	return AVVCAudioDeviceManager{objectivec.Object{ID: id}}
}

// Ensure AVVCAudioDeviceManager implements IAVVCAudioDeviceManager.
var _ IAVVCAudioDeviceManager = AVVCAudioDeviceManager{}

// An interface definition for the [AVVCAudioDeviceManager] class.
//
// # Methods
//
//   - [IAVVCAudioDeviceManager.InputDeviceID]
//   - [IAVVCAudioDeviceManager.OutputDeviceID]
//   - [IAVVCAudioDeviceManager.SelectDeviceForActivationMode]
//   - [IAVVCAudioDeviceManager.InitWithActivationContextWithError]
//   - [IAVVCAudioDeviceManager.InitWithError]
type IAVVCAudioDeviceManager interface {
	objectivec.IObject

	// Topic: Methods

	InputDeviceID() uint32
	OutputDeviceID() uint32
	SelectDeviceForActivationMode(mode objectivec.IObject)
	InitWithActivationContextWithError(context objectivec.IObject) (AVVCAudioDeviceManager, error)
	InitWithError() (AVVCAudioDeviceManager, error)
}

// Init initializes the instance.
func (a AVVCAudioDeviceManager) Init() AVVCAudioDeviceManager {
	rv := objc.Send[AVVCAudioDeviceManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCAudioDeviceManager) Autorelease() AVVCAudioDeviceManager {
	rv := objc.Send[AVVCAudioDeviceManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCAudioDeviceManager creates a new AVVCAudioDeviceManager instance.
func NewAVVCAudioDeviceManager() AVVCAudioDeviceManager {
	class := getAVVCAudioDeviceManagerClass()
	rv := objc.Send[AVVCAudioDeviceManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVCAudioDeviceManagerWithActivationContextWithError(context objectivec.IObject) (AVVCAudioDeviceManager, error) {
	var errorPtr objc.ID
	instance := getAVVCAudioDeviceManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithActivationContext:withError:"), context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVVCAudioDeviceManager{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVVCAudioDeviceManagerFromID(rv), nil
}

func NewVCAudioDeviceManagerWithError() (AVVCAudioDeviceManager, error) {
	var errorPtr objc.ID
	instance := getAVVCAudioDeviceManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVVCAudioDeviceManager{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVVCAudioDeviceManagerFromID(rv), nil
}

func (a AVVCAudioDeviceManager) SelectDeviceForActivationMode(mode objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("selectDeviceForActivationMode:"), mode)
}
func (a AVVCAudioDeviceManager) InitWithActivationContextWithError(context objectivec.IObject) (AVVCAudioDeviceManager, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithActivationContext:withError:"), context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVVCAudioDeviceManager{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVVCAudioDeviceManagerFromID(rv), nil

}
func (a AVVCAudioDeviceManager) InitWithError() (AVVCAudioDeviceManager, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVVCAudioDeviceManager{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVVCAudioDeviceManagerFromID(rv), nil

}

func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) BuildAOPAggregateDevice() {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("BuildAOPAggregateDevice"))
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) DestroyAOPAggregateDevice() {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("DestroyAOPAggregateDevice"))
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetAOPDeviceAggregateID() uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetAOPDeviceAggregateID"))
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetAOPDeviceAggregateUID() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetAOPDeviceAggregateUID"))
	return objectivec.Object{ID: rv}
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetAOPDeviceID() uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetAOPDeviceID"))
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetAudioDeviceBuiltInMicrophone() uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetAudioDeviceBuiltInMicrophone"))
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetAudioDeviceWithWiredHeadset(headset bool) uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetAudioDeviceWithWiredHeadset:"), headset)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetBuiltInAudioSpeakerDevice() uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetBuiltInAudioSpeakerDevice"))
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetDefaultAudioDevice(device bool) uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetDefaultAudioDevice:"), device)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetDevicesForActivationModeOutRecordDeviceOutPlaybackDevice(mode int64, device *uint32, device2 *uint32) {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetDevicesForActivationMode:outRecordDevice:outPlaybackDevice:"), mode, device, device2)
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetHALDeviceForBuiltInDeviceOutHALDeviceIDOutHALDeviceUID(device objectivec.IObject, id *uint32, uid []objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetHALDeviceForBuiltInDevice:outHALDeviceID:outHALDeviceUID:"), device, id, objectivec.IObjectSliceToNSArray(uid))
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetHALDeviceIdentifiersForDeviceUIDOutHALDeviceUIDOutPluginDeviceUUID(uid objectivec.IObject, uid2 []objectivec.IObject, uuid []objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetHALDeviceIdentifiersForDeviceUID:outHALDeviceUID:outPluginDeviceUUID:"), uid, objectivec.IObjectSliceToNSArray(uid2), objectivec.IObjectSliceToNSArray(uuid))
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetNonAppleBluetoothSpeakerDevice() uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetNonAppleBluetoothSpeakerDevice"))
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetNumberOfInputChannels() uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetNumberOfInputChannels"))
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetSiriSupportedExternalDevice() uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("GetSiriSupportedExternalDevice"))
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsAOPDevicePresent() bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("IsAOPDevicePresent"))
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsAppleDisplayDevice(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("IsAppleDisplayDevice:"), device)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsDeviceBuiltIn(in objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("IsDeviceBuiltIn:"), in)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsDeviceRunningSomewhere(somewhere objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("IsDeviceRunningSomewhere:"), somewhere)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsDeviceRunningSomewhereElse(else_ objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("IsDeviceRunningSomewhereElse:"), else_)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsDoAPSupportedDevice(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("IsDoAPSupportedDevice:"), device)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsHALDevice(hALDevice objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("IsHALDevice:"), hALDevice)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsSiriSupportedAppleDisplay(display objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("IsSiriSupportedAppleDisplay:"), display)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsSiriSupportedExternalDevice(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("IsSiriSupportedExternalDevice:"), device)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) UpdateAOPDeviceID() {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("UpdateAOPDeviceID"))
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) AddReporterIDToDevice(id int64, device objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("addReporterID:toDevice:"), id, device)
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) FindHALDeviceUIDFromUUID(uuid objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("findHALDeviceUIDFromUUID:"), uuid)
	return objectivec.Object{ID: rv}
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GainOnDevice(device uint32) float32 {
	rv := objc.Send[float32](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("gainOnDevice:"), device)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetAuditTokenFromProcessObjectID(id uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("getAuditTokenFromProcessObjectID:"), id)
	return objectivec.Object{ID: rv}
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetBundleIDFromProcessObjectID(id uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("getBundleIDFromProcessObjectID:"), id)
	return objectivec.Object{ID: rv}
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetDeviceName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("getDeviceName:"), name)
	return objectivec.Object{ID: rv}
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetPIDFromProcessObjectID(id uint32) int {
	rv := objc.Send[int](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("getPIDFromProcessObjectID:"), id)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetRecordingAuditTokenList() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("getRecordingAuditTokenList"))
	return objectivec.Object{ID: rv}
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetRecordingBundleIDList() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("getRecordingBundleIDList"))
	return objectivec.Object{ID: rv}
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetRecordingPIDList() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("getRecordingPIDList"))
	return objectivec.Object{ID: rv}
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) GetUUIDFromBTHALDevice(bTHALDevice objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("getUUIDFromBTHALDevice:"), bTHALDevice)
	return objectivec.Object{ID: rv}
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsDeviceAlive(alive uint32) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("isDeviceAlive:"), alive)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsDeviceMutedIsInputDevice(muted uint32, device bool) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("isDeviceMuted:isInputDevice:"), muted, device)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IsProcessDeviceInputRunning(running uint32) bool {
	rv := objc.Send[bool](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("isProcessDeviceInputRunning:"), running)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) IterateOverProcessObjectList(list VoidHandler) {
	_block0, _ := NewVoidBlock(list)
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("iterateOverProcessObjectList:"), _block0)
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) RemoveReporterIDFromDevice(id int64, device objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("removeReporterID:fromDevice:"), id, device)
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) SetBufferSizeOnDevice(size int, device objectivec.IObject) int {
	rv := objc.Send[int](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("setBufferSize:onDevice:"), size, device)
	return rv
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) SetClientDescriptionKindOnDevice(kind uint32, device uint32) {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("setClientDescriptionKind:onDevice:"), kind, device)
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) SetGainOnDevice(gain float32, device uint32) {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("setGain:onDevice:"), gain, device)
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) SetMuteOnDeviceIsInputDevice(mute bool, device uint32, device2 bool) {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("setMute:onDevice:isInputDevice:"), mute, device, device2)
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) SetReporterIDsOnDevice(iDs unsafe.Pointer, device objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("setReporterIDs:onDevice:"), iDs, device)
}
func (_AVVCAudioDeviceManagerClass AVVCAudioDeviceManagerClass) SuppressRecordingIndicator(indicator uint32) {
	objc.Send[objc.ID](objc.ID(_AVVCAudioDeviceManagerClass.class), objc.Sel("suppressRecordingIndicator:"), indicator)
}

func (a AVVCAudioDeviceManager) InputDeviceID() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("inputDeviceID"))
	return rv
}
func (a AVVCAudioDeviceManager) OutputDeviceID() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("outputDeviceID"))
	return rv
}

// IterateOverProcessObjectListSync is a synchronous wrapper around [AVVCAudioDeviceManager.IterateOverProcessObjectList].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AVVCAudioDeviceManagerClass) IterateOverProcessObjectListSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	ac.IterateOverProcessObjectList(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
