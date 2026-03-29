// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CATapDescription] class.
var (
	_CATapDescriptionClass     CATapDescriptionClass
	_CATapDescriptionClassOnce sync.Once
)

func getCATapDescriptionClass() CATapDescriptionClass {
	_CATapDescriptionClassOnce.Do(func() {
		_CATapDescriptionClass = CATapDescriptionClass{class: objc.GetClass("CATapDescription")}
	})
	return _CATapDescriptionClass
}

// GetCATapDescriptionClass returns the class object for CATapDescription.
func GetCATapDescriptionClass() CATapDescriptionClass {
	return getCATapDescriptionClass()
}

type CATapDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CATapDescriptionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CATapDescriptionClass) Alloc() CATapDescription {
	rv := objc.Send[CATapDescription](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

//
// # Overview
// 
// This class describes a tap object that contains an input stream. The input
// stream is a mix of all of the specified processes output audio.
//
// # Instance Properties
//
//   - [CATapDescription.BundleIDs]
//   - [CATapDescription.SetBundleIDs]
//   - [CATapDescription.DeviceUID]
//   - [CATapDescription.SetDeviceUID]
//   - [CATapDescription.Exclusive]
//   - [CATapDescription.SetExclusive]
//   - [CATapDescription.Mixdown]
//   - [CATapDescription.SetMixdown]
//   - [CATapDescription.Mono]
//   - [CATapDescription.SetMono]
//   - [CATapDescription.PrivateTap]
//   - [CATapDescription.SetPrivateTap]
//   - [CATapDescription.ProcessRestoreEnabled]
//   - [CATapDescription.SetProcessRestoreEnabled]
//   - [CATapDescription.MuteBehavior]
//   - [CATapDescription.SetMuteBehavior]
//   - [CATapDescription.Name]
//   - [CATapDescription.SetName]
//   - [CATapDescription.UUID]
//   - [CATapDescription.SetUUID]
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription
type CATapDescription struct {
	objectivec.Object
}

// CATapDescriptionFromID constructs a [CATapDescription] from an objc.ID.
func CATapDescriptionFromID(id objc.ID) CATapDescription {
	return CATapDescription{objectivec.Object{ID: id}}
}
// NOTE: CATapDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CATapDescription] class.
//
// # Instance Properties
//
//   - [ICATapDescription.BundleIDs]
//   - [ICATapDescription.SetBundleIDs]
//   - [ICATapDescription.DeviceUID]
//   - [ICATapDescription.SetDeviceUID]
//   - [ICATapDescription.Exclusive]
//   - [ICATapDescription.SetExclusive]
//   - [ICATapDescription.Mixdown]
//   - [ICATapDescription.SetMixdown]
//   - [ICATapDescription.Mono]
//   - [ICATapDescription.SetMono]
//   - [ICATapDescription.PrivateTap]
//   - [ICATapDescription.SetPrivateTap]
//   - [ICATapDescription.ProcessRestoreEnabled]
//   - [ICATapDescription.SetProcessRestoreEnabled]
//   - [ICATapDescription.MuteBehavior]
//   - [ICATapDescription.SetMuteBehavior]
//   - [ICATapDescription.Name]
//   - [ICATapDescription.SetName]
//   - [ICATapDescription.UUID]
//   - [ICATapDescription.SetUUID]
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription
type ICATapDescription interface {
	objectivec.IObject

	// Topic: Instance Properties

	BundleIDs() []string
	SetBundleIDs(value []string)
	DeviceUID() string
	SetDeviceUID(value string)
	Exclusive() bool
	SetExclusive(value bool)
	Mixdown() bool
	SetMixdown(value bool)
	Mono() bool
	SetMono(value bool)
	PrivateTap() bool
	SetPrivateTap(value bool)
	ProcessRestoreEnabled() bool
	SetProcessRestoreEnabled(value bool)
	MuteBehavior() CATapMuteBehavior
	SetMuteBehavior(value CATapMuteBehavior)
	Name() string
	SetName(value string)
	UUID() foundation.NSUUID
	SetUUID(value foundation.NSUUID)

	Processes() []foundation.NSNumber
	SetProcesses(value []foundation.NSNumber)
	Stream() foundation.NSNumber
	SetStream(value foundation.NSNumber)
	InitExcludingProcessesAndDeviceUIDWithStream(processesObjectIDsToExcludeFromTap []foundation.NSNumber, deviceUID string, stream int) CATapDescription
	InitMonoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap []foundation.NSNumber) CATapDescription
	InitMonoMixdownOfProcesses(processesObjectIDsToIncludeInTap []foundation.NSNumber) CATapDescription
	InitStereoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap []foundation.NSNumber) CATapDescription
	InitStereoMixdownOfProcesses(processesObjectIDsToIncludeInTap []foundation.NSNumber) CATapDescription
	InitWithProcessesAndDeviceUIDWithStream(processesObjectIDsToIncludeInTap []foundation.NSNumber, deviceUID string, stream int) CATapDescription
}

// Init initializes the instance.
func (t CATapDescription) Init() CATapDescription {
	rv := objc.Send[CATapDescription](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t CATapDescription) Autorelease() CATapDescription {
	rv := objc.Send[CATapDescription](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewCATapDescription creates a new CATapDescription instance.
func NewCATapDescription() CATapDescription {
	class := getCATapDescriptionClass()
	rv := objc.Send[CATapDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// processesObjectIDsToExcludeFromTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to exclude from the tap. All other processes that output
// audio will be included in the tap.
//
// deviceUID: The device UID of the output device whose audio will be captured
//
// stream: NSInteger that represents the index of the stream on the device whose audio
// will be captured. The format of the tap will match the format of this
// stream.
//
// # Discussion
// 
// Mix all process audio streams destined for the selected device stream
// except the given processes
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initExcludingProcesses:andDeviceUID:withStream:
func NewTapDescriptionExcludingProcessesAndDeviceUIDWithStream(processesObjectIDsToExcludeFromTap []foundation.NSNumber, deviceUID string, stream int) CATapDescription {
	instance := getCATapDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initExcludingProcesses:andDeviceUID:withStream:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToExcludeFromTap), objc.String(deviceUID), stream)
	return CATapDescriptionFromID(rv)
}

//
// processesObjectIDsToExcludeFromTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to exclude from the tap. All other processes that output
// audio will be included in the tap.
//
// # Discussion
// 
// Mix all processes to a mono stream except the given processes
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initMonoGlobalTapButExcludeProcesses:
func NewTapDescriptionMonoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap []foundation.NSNumber) CATapDescription {
	instance := getCATapDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initMonoGlobalTapButExcludeProcesses:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToExcludeFromTap))
	return CATapDescriptionFromID(rv)
}

//
// processesObjectIDsToIncludeInTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to include in the tap
//
// # Discussion
// 
// Mix all given process audio streams audio to mono.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initMonoMixdownOfProcesses:
func NewTapDescriptionMonoMixdownOfProcesses(processesObjectIDsToIncludeInTap []foundation.NSNumber) CATapDescription {
	instance := getCATapDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initMonoMixdownOfProcesses:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToIncludeInTap))
	return CATapDescriptionFromID(rv)
}

//
// processesObjectIDsToExcludeFromTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to exclude from the tap. All other processes that output
// audio will be included in the tap.
//
// # Discussion
// 
// Mix all processes to a stereo stream except the given processes. Mono
// sources will be duplicated in both right and left channels.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initStereoGlobalTapButExcludeProcesses:
func NewTapDescriptionStereoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap []foundation.NSNumber) CATapDescription {
	instance := getCATapDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initStereoGlobalTapButExcludeProcesses:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToExcludeFromTap))
	return CATapDescriptionFromID(rv)
}

//
// processesObjectIDsToIncludeInTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to include in the tap
//
// # Discussion
// 
// Mix all given process audio streams down to stereo. Mono sources will be
// duplicated in both right and left channels.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initStereoMixdownOfProcesses:
func NewTapDescriptionStereoMixdownOfProcesses(processesObjectIDsToIncludeInTap []foundation.NSNumber) CATapDescription {
	instance := getCATapDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initStereoMixdownOfProcesses:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToIncludeInTap))
	return CATapDescriptionFromID(rv)
}

//
// processesObjectIDsToIncludeInTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to exclude from the tap. All other processes that output
// audio will be included in the tap.
//
// deviceUID: The device UID of the output device whose audio will be captured
//
// stream: NSInteger that represents the index of the stream on the device whose audio
// will be captured. The format of the tap will match the format of this
// stream.
//
// # Discussion
// 
// Mix all given process audio streams destined for the selected device stream
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initWithProcesses:andDeviceUID:withStream:
func NewTapDescriptionWithProcessesAndDeviceUIDWithStream(processesObjectIDsToIncludeInTap []foundation.NSNumber, deviceUID string, stream int) CATapDescription {
	instance := getCATapDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcesses:andDeviceUID:withStream:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToIncludeInTap), objc.String(deviceUID), stream)
	return CATapDescriptionFromID(rv)
}

//
// processesObjectIDsToExcludeFromTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to exclude from the tap. All other processes that output
// audio will be included in the tap.
//
// deviceUID: The device UID of the output device whose audio will be captured
//
// stream: NSInteger that represents the index of the stream on the device whose audio
// will be captured. The format of the tap will match the format of this
// stream.
//
// # Discussion
// 
// Mix all process audio streams destined for the selected device stream
// except the given processes
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initExcludingProcesses:andDeviceUID:withStream:
func (t CATapDescription) InitExcludingProcessesAndDeviceUIDWithStream(processesObjectIDsToExcludeFromTap []foundation.NSNumber, deviceUID string, stream int) CATapDescription {
	rv := objc.Send[CATapDescription](t.ID, objc.Sel("initExcludingProcesses:andDeviceUID:withStream:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToExcludeFromTap), objc.String(deviceUID), stream)
	return rv
}
//
// processesObjectIDsToExcludeFromTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to exclude from the tap. All other processes that output
// audio will be included in the tap.
//
// # Discussion
// 
// Mix all processes to a mono stream except the given processes
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initMonoGlobalTapButExcludeProcesses:
func (t CATapDescription) InitMonoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap []foundation.NSNumber) CATapDescription {
	rv := objc.Send[CATapDescription](t.ID, objc.Sel("initMonoGlobalTapButExcludeProcesses:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToExcludeFromTap))
	return rv
}
//
// processesObjectIDsToIncludeInTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to include in the tap
//
// # Discussion
// 
// Mix all given process audio streams audio to mono.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initMonoMixdownOfProcesses:
func (t CATapDescription) InitMonoMixdownOfProcesses(processesObjectIDsToIncludeInTap []foundation.NSNumber) CATapDescription {
	rv := objc.Send[CATapDescription](t.ID, objc.Sel("initMonoMixdownOfProcesses:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToIncludeInTap))
	return rv
}
//
// processesObjectIDsToExcludeFromTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to exclude from the tap. All other processes that output
// audio will be included in the tap.
//
// # Discussion
// 
// Mix all processes to a stereo stream except the given processes. Mono
// sources will be duplicated in both right and left channels.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initStereoGlobalTapButExcludeProcesses:
func (t CATapDescription) InitStereoGlobalTapButExcludeProcesses(processesObjectIDsToExcludeFromTap []foundation.NSNumber) CATapDescription {
	rv := objc.Send[CATapDescription](t.ID, objc.Sel("initStereoGlobalTapButExcludeProcesses:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToExcludeFromTap))
	return rv
}
//
// processesObjectIDsToIncludeInTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to include in the tap
//
// # Discussion
// 
// Mix all given process audio streams down to stereo. Mono sources will be
// duplicated in both right and left channels.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initStereoMixdownOfProcesses:
func (t CATapDescription) InitStereoMixdownOfProcesses(processesObjectIDsToIncludeInTap []foundation.NSNumber) CATapDescription {
	rv := objc.Send[CATapDescription](t.ID, objc.Sel("initStereoMixdownOfProcesses:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToIncludeInTap))
	return rv
}
//
// processesObjectIDsToIncludeInTap: An NSArray of NSNumbers where each NSNumber holds an AudioObjectID of the
// process object to exclude from the tap. All other processes that output
// audio will be included in the tap.
//
// deviceUID: The device UID of the output device whose audio will be captured
//
// stream: NSInteger that represents the index of the stream on the device whose audio
// will be captured. The format of the tap will match the format of this
// stream.
//
// # Discussion
// 
// Mix all given process audio streams destined for the selected device stream
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/initWithProcesses:andDeviceUID:withStream:
func (t CATapDescription) InitWithProcessesAndDeviceUIDWithStream(processesObjectIDsToIncludeInTap []foundation.NSNumber, deviceUID string, stream int) CATapDescription {
	rv := objc.Send[CATapDescription](t.ID, objc.Sel("initWithProcesses:andDeviceUID:withStream:"), objectivec.IObjectSliceToNSArray(processesObjectIDsToIncludeInTap), objc.String(deviceUID), stream)
	return rv
}

//
// # Discussion
// 
// An Array of Strings where each String holds the bundle ID of a process to
// tap or exclude.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/bundleIDs
func (t CATapDescription) BundleIDs() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("bundleIDs"))
	return objc.ConvertSliceToStrings(rv)
}
func (t CATapDescription) SetBundleIDs(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setBundleIDs:"), objectivec.StringSliceToNSArray(value))
}
//
// # Discussion
// 
// An optional deviceUID that will have a value if this tap only taps a
// specific hardware device
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/deviceUID
func (t CATapDescription) DeviceUID() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("deviceUID"))
	return foundation.NSStringFromID(rv).String()
}
func (t CATapDescription) SetDeviceUID(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setDeviceUID:"), objc.String(value))
}
//
// # Discussion
// 
// True if this description should tap all processes except the process listed
// in the ‘processes’ property.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isExclusive
func (t CATapDescription) Exclusive() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isExclusive"))
	return rv
}
func (t CATapDescription) SetExclusive(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setExclusive:"), value)
}
//
// # Discussion
// 
// True if this description is a mono or stereo mix of the tapped device’s
// channels.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMixdown
func (t CATapDescription) Mixdown() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isMixdown"))
	return rv
}
func (t CATapDescription) SetMixdown(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setMixdown:"), value)
}
//
// # Discussion
// 
// True if this description is a mono mixdown of channels.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isMono
func (t CATapDescription) Mono() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isMono"))
	return rv
}
func (t CATapDescription) SetMono(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setMono:"), value)
}
//
// # Discussion
// 
// True if this tap is only visible to the client process that created the
// tap.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isPrivate
func (t CATapDescription) PrivateTap() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isPrivate"))
	return rv
}
func (t CATapDescription) SetPrivateTap(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setPrivate:"), value)
}
//
// # Discussion
// 
// True if this tap should save tapped processes by bundle ID when they exit,
// and restore them to the tap when they start up again.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/isProcessRestoreEnabled
func (t CATapDescription) ProcessRestoreEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isProcessRestoreEnabled"))
	return rv
}
func (t CATapDescription) SetProcessRestoreEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setProcessRestoreEnabled:"), value)
}
//
// # Discussion
// 
// Set the tap’s mute behavior. See CATapMuteBehavior above.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/muteBehavior
func (t CATapDescription) MuteBehavior() CATapMuteBehavior {
	rv := objc.Send[CATapMuteBehavior](t.ID, objc.Sel("isMuted"))
	return CATapMuteBehavior(rv)
}
func (t CATapDescription) SetMuteBehavior(value CATapMuteBehavior) {
	objc.Send[struct{}](t.ID, objc.Sel("setMuteBehavior:"), value)
}
//
// # Discussion
// 
// Human readable name of this tap.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/name
func (t CATapDescription) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (t CATapDescription) SetName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setName:"), objc.String(value))
}
//
// # Discussion
// 
// UID of this tap.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/uuid
func (t CATapDescription) UUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("UUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (t CATapDescription) SetUUID(value foundation.NSUUID) {
	objc.Send[struct{}](t.ID, objc.Sel("setUUID:"), value)
}
//
// # Discussion
// 
// An NSArray of NSNumbers where each NSNumber holds the AudioObjectID of a
// process object to tap or exclude.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/processes-3cdzw
func (t CATapDescription) Processes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("processes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (t CATapDescription) SetProcesses(value []foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setProcesses:"), objectivec.IObjectSliceToNSArray(value))
}
//
// # Discussion
// 
// An optional NSNumber that will have a value if this tap taps a specific
// device stream. The value represents the index of the hardware stream.
//
// See: https://developer.apple.com/documentation/CoreAudio/CATapDescription/stream-u4ff
func (t CATapDescription) Stream() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("stream"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (t CATapDescription) SetStream(value foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setStream:"), value)
}

