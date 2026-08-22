// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MEHEVCDependencyInfo] class.
var (
	_MEHEVCDependencyInfoClass     MEHEVCDependencyInfoClass
	_MEHEVCDependencyInfoClassOnce sync.Once
)

func getMEHEVCDependencyInfoClass() MEHEVCDependencyInfoClass {
	_MEHEVCDependencyInfoClassOnce.Do(func() {
		_MEHEVCDependencyInfoClass = MEHEVCDependencyInfoClass{class: objc.GetClass("MEHEVCDependencyInfo")}
	})
	return _MEHEVCDependencyInfoClass
}

// GetMEHEVCDependencyInfoClass returns the class object for MEHEVCDependencyInfo.
func GetMEHEVCDependencyInfoClass() MEHEVCDependencyInfoClass {
	return getMEHEVCDependencyInfoClass()
}

type MEHEVCDependencyInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MEHEVCDependencyInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MEHEVCDependencyInfoClass) Alloc() MEHEVCDependencyInfo {
	rv := objc.Send[MEHEVCDependencyInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about the HEVC dependency attributes of
// a sample.
//
// # Inspecting the HEVC dependency attributes of a sample
//
//   - [MEHEVCDependencyInfo.TemporalSubLayerAccess]: A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
//   - [MEHEVCDependencyInfo.SetTemporalSubLayerAccess]
//   - [MEHEVCDependencyInfo.StepwiseTemporalSubLayerAccess]: A Boolean value that indicates if the sample has an HEVC stepwise temporal sublayer access (STSA) picture.
//   - [MEHEVCDependencyInfo.SetStepwiseTemporalSubLayerAccess]
//   - [MEHEVCDependencyInfo.SyncSampleNALUnitType]: The NAL unit type for HEVC sync sample groups.
//   - [MEHEVCDependencyInfo.SetSyncSampleNALUnitType]
//   - [MEHEVCDependencyInfo.TemporalLevel]: The HEVC temporal level, if available.
//   - [MEHEVCDependencyInfo.SetTemporalLevel]
//   - [MEHEVCDependencyInfo.ProfileSpace]: The HEVC profile space, if available.
//   - [MEHEVCDependencyInfo.SetProfileSpace]
//   - [MEHEVCDependencyInfo.TierFlag]: The HEVC tier level flag, if available.
//   - [MEHEVCDependencyInfo.SetTierFlag]
//   - [MEHEVCDependencyInfo.ProfileIndex]: The HEVC profile index, if available.
//   - [MEHEVCDependencyInfo.SetProfileIndex]
//   - [MEHEVCDependencyInfo.ProfileCompatibilityFlags]: The HEVC profile compatibility flags (4 bytes), if available.
//   - [MEHEVCDependencyInfo.SetProfileCompatibilityFlags]
//   - [MEHEVCDependencyInfo.ConstraintIndicatorFlags]: The HEVC constraint indicator flags (6 bytes), if available.
//   - [MEHEVCDependencyInfo.SetConstraintIndicatorFlags]
//   - [MEHEVCDependencyInfo.LevelIndex]: The HEVC level index, if available.
//   - [MEHEVCDependencyInfo.SetLevelIndex]
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo
type MEHEVCDependencyInfo struct {
	objectivec.Object
}

// MEHEVCDependencyInfoFromID constructs a [MEHEVCDependencyInfo] from an objc.ID.
//
// An object that provides information about the HEVC dependency attributes of
// a sample.
func MEHEVCDependencyInfoFromID(id objc.ID) MEHEVCDependencyInfo {
	return MEHEVCDependencyInfo{objectivec.Object{ID: id}}
}

// NOTE: MEHEVCDependencyInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MEHEVCDependencyInfo] class.
//
// # Inspecting the HEVC dependency attributes of a sample
//
//   - [IMEHEVCDependencyInfo.TemporalSubLayerAccess]: A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
//   - [IMEHEVCDependencyInfo.SetTemporalSubLayerAccess]
//   - [IMEHEVCDependencyInfo.StepwiseTemporalSubLayerAccess]: A Boolean value that indicates if the sample has an HEVC stepwise temporal sublayer access (STSA) picture.
//   - [IMEHEVCDependencyInfo.SetStepwiseTemporalSubLayerAccess]
//   - [IMEHEVCDependencyInfo.SyncSampleNALUnitType]: The NAL unit type for HEVC sync sample groups.
//   - [IMEHEVCDependencyInfo.SetSyncSampleNALUnitType]
//   - [IMEHEVCDependencyInfo.TemporalLevel]: The HEVC temporal level, if available.
//   - [IMEHEVCDependencyInfo.SetTemporalLevel]
//   - [IMEHEVCDependencyInfo.ProfileSpace]: The HEVC profile space, if available.
//   - [IMEHEVCDependencyInfo.SetProfileSpace]
//   - [IMEHEVCDependencyInfo.TierFlag]: The HEVC tier level flag, if available.
//   - [IMEHEVCDependencyInfo.SetTierFlag]
//   - [IMEHEVCDependencyInfo.ProfileIndex]: The HEVC profile index, if available.
//   - [IMEHEVCDependencyInfo.SetProfileIndex]
//   - [IMEHEVCDependencyInfo.ProfileCompatibilityFlags]: The HEVC profile compatibility flags (4 bytes), if available.
//   - [IMEHEVCDependencyInfo.SetProfileCompatibilityFlags]
//   - [IMEHEVCDependencyInfo.ConstraintIndicatorFlags]: The HEVC constraint indicator flags (6 bytes), if available.
//   - [IMEHEVCDependencyInfo.SetConstraintIndicatorFlags]
//   - [IMEHEVCDependencyInfo.LevelIndex]: The HEVC level index, if available.
//   - [IMEHEVCDependencyInfo.SetLevelIndex]
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo
type IMEHEVCDependencyInfo interface {
	objectivec.IObject

	// Topic: Inspecting the HEVC dependency attributes of a sample

	// A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
	TemporalSubLayerAccess() bool
	SetTemporalSubLayerAccess(value bool)
	// A Boolean value that indicates if the sample has an HEVC stepwise temporal sublayer access (STSA) picture.
	StepwiseTemporalSubLayerAccess() bool
	SetStepwiseTemporalSubLayerAccess(value bool)
	// The NAL unit type for HEVC sync sample groups.
	SyncSampleNALUnitType() int16
	SetSyncSampleNALUnitType(value int16)
	// The HEVC temporal level, if available.
	TemporalLevel() int16
	SetTemporalLevel(value int16)
	// The HEVC profile space, if available.
	ProfileSpace() int16
	SetProfileSpace(value int16)
	// The HEVC tier level flag, if available.
	TierFlag() int16
	SetTierFlag(value int16)
	// The HEVC profile index, if available.
	ProfileIndex() int16
	SetProfileIndex(value int16)
	// The HEVC profile compatibility flags (4 bytes), if available.
	ProfileCompatibilityFlags() foundation.NSData
	SetProfileCompatibilityFlags(value foundation.NSData)
	// The HEVC constraint indicator flags (6 bytes), if available.
	ConstraintIndicatorFlags() foundation.NSData
	SetConstraintIndicatorFlags(value foundation.NSData)
	// The HEVC level index, if available.
	LevelIndex() int16
	SetLevelIndex(value int16)
}

// Init initializes the instance.
func (m MEHEVCDependencyInfo) Init() MEHEVCDependencyInfo {
	rv := objc.Send[MEHEVCDependencyInfo](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MEHEVCDependencyInfo) Autorelease() MEHEVCDependencyInfo {
	rv := objc.Send[MEHEVCDependencyInfo](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEHEVCDependencyInfo creates a new MEHEVCDependencyInfo instance.
func NewMEHEVCDependencyInfo() MEHEVCDependencyInfo {
	class := getMEHEVCDependencyInfoClass()
	rv := objc.Send[MEHEVCDependencyInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates if the sample has an HEVC temporal sublayer
// access (TSA) picture.
//
// # Discussion
//
// This value maps to the [kCMSampleAttachmentKey_HEVCTemporalSubLayerAccess]
// sample buffer attachment.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/hasTemporalSubLayerAccess
//
// [kCMSampleAttachmentKey_HEVCTemporalSubLayerAccess]: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_HEVCTemporalSubLayerAccess
func (m MEHEVCDependencyInfo) TemporalSubLayerAccess() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasTemporalSubLayerAccess"))
	return rv
}
func (m MEHEVCDependencyInfo) SetTemporalSubLayerAccess(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setTemporalSubLayerAccess:"), value)
}

// A Boolean value that indicates if the sample has an HEVC stepwise temporal
// sublayer access (STSA) picture.
//
// # Discussion
//
// This value maps to the
// [kCMSampleAttachmentKey_HEVCStepwiseTemporalSubLayerAccess] sample buffer
// attachment.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/hasStepwiseTemporalSubLayerAccess
//
// [kCMSampleAttachmentKey_HEVCStepwiseTemporalSubLayerAccess]: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_HEVCStepwiseTemporalSubLayerAccess
func (m MEHEVCDependencyInfo) StepwiseTemporalSubLayerAccess() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasStepwiseTemporalSubLayerAccess"))
	return rv
}
func (m MEHEVCDependencyInfo) SetStepwiseTemporalSubLayerAccess(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setStepwiseTemporalSubLayerAccess:"), value)
}

// The NAL unit type for HEVC sync sample groups.
//
// # Discussion
//
// This value maps to the [kCMSampleAttachmentKey_HEVCSyncSampleNALUnitType]
// sample buffer attachment.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/syncSampleNALUnitType
//
// [kCMSampleAttachmentKey_HEVCSyncSampleNALUnitType]: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_HEVCSyncSampleNALUnitType
func (m MEHEVCDependencyInfo) SyncSampleNALUnitType() int16 {
	rv := objc.Send[int16](m.ID, objc.Sel("syncSampleNALUnitType"))
	return rv
}
func (m MEHEVCDependencyInfo) SetSyncSampleNALUnitType(value int16) {
	objc.Send[struct{}](m.ID, objc.Sel("setSyncSampleNALUnitType:"), value)
}

// The HEVC temporal level, if available.
//
// # Discussion
//
// This value maps to the [kCMHEVCTemporalLevelInfoKey_TemporalLevel] sample
// buffer attachment, and is `-1` if this information isn’t available.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/temporalLevel
//
// [kCMHEVCTemporalLevelInfoKey_TemporalLevel]: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_TemporalLevel
func (m MEHEVCDependencyInfo) TemporalLevel() int16 {
	rv := objc.Send[int16](m.ID, objc.Sel("temporalLevel"))
	return rv
}
func (m MEHEVCDependencyInfo) SetTemporalLevel(value int16) {
	objc.Send[struct{}](m.ID, objc.Sel("setTemporalLevel:"), value)
}

// The HEVC profile space, if available.
//
// # Discussion
//
// This value maps to the [kCMHEVCTemporalLevelInfoKey_ProfileSpace] sample
// buffer attachment, and is `-1` if this information isn’t available.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileSpace
//
// [kCMHEVCTemporalLevelInfoKey_ProfileSpace]: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_ProfileSpace
func (m MEHEVCDependencyInfo) ProfileSpace() int16 {
	rv := objc.Send[int16](m.ID, objc.Sel("profileSpace"))
	return rv
}
func (m MEHEVCDependencyInfo) SetProfileSpace(value int16) {
	objc.Send[struct{}](m.ID, objc.Sel("setProfileSpace:"), value)
}

// The HEVC tier level flag, if available.
//
// # Discussion
//
// This value maps to the [kCMHEVCTemporalLevelInfoKey_TierFlag] sample buffer
// attachment, and is `-1` if this information isn’t available.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/tierFlag
//
// [kCMHEVCTemporalLevelInfoKey_TierFlag]: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_TierFlag
func (m MEHEVCDependencyInfo) TierFlag() int16 {
	rv := objc.Send[int16](m.ID, objc.Sel("tierFlag"))
	return rv
}
func (m MEHEVCDependencyInfo) SetTierFlag(value int16) {
	objc.Send[struct{}](m.ID, objc.Sel("setTierFlag:"), value)
}

// The HEVC profile index, if available.
//
// # Discussion
//
// This value maps to the [kCMHEVCTemporalLevelInfoKey_ProfileIndex] sample
// buffer attachment, and is `-1` if this information isn’t available.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileIndex
//
// [kCMHEVCTemporalLevelInfoKey_ProfileIndex]: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_ProfileIndex
func (m MEHEVCDependencyInfo) ProfileIndex() int16 {
	rv := objc.Send[int16](m.ID, objc.Sel("profileIndex"))
	return rv
}
func (m MEHEVCDependencyInfo) SetProfileIndex(value int16) {
	objc.Send[struct{}](m.ID, objc.Sel("setProfileIndex:"), value)
}

// The HEVC profile compatibility flags (4 bytes), if available.
//
// # Discussion
//
// This value maps to the
// [kCMHEVCTemporalLevelInfoKey_ProfileCompatibilityFlags] sample buffer
// attachment, and is `nil` if this information isn’t available.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileCompatibilityFlags
//
// [kCMHEVCTemporalLevelInfoKey_ProfileCompatibilityFlags]: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_ProfileCompatibilityFlags
func (m MEHEVCDependencyInfo) ProfileCompatibilityFlags() foundation.NSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("profileCompatibilityFlags"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m MEHEVCDependencyInfo) SetProfileCompatibilityFlags(value foundation.NSData) {
	objc.Send[struct{}](m.ID, objc.Sel("setProfileCompatibilityFlags:"), value)
}

// The HEVC constraint indicator flags (6 bytes), if available.
//
// # Discussion
//
// This value maps to the
// [kCMHEVCTemporalLevelInfoKey_ConstraintIndicatorFlags] sample buffer
// attachment, and is `nil` if this information isn’t available.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/constraintIndicatorFlags
//
// [kCMHEVCTemporalLevelInfoKey_ConstraintIndicatorFlags]: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_ConstraintIndicatorFlags
func (m MEHEVCDependencyInfo) ConstraintIndicatorFlags() foundation.NSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("constraintIndicatorFlags"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (m MEHEVCDependencyInfo) SetConstraintIndicatorFlags(value foundation.NSData) {
	objc.Send[struct{}](m.ID, objc.Sel("setConstraintIndicatorFlags:"), value)
}

// The HEVC level index, if available.
//
// # Discussion
//
// This value maps to the [kCMHEVCTemporalLevelInfoKey_LevelIndex] sample
// buffer attachment, and is `-1` if this information isn’t available.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/levelIndex
//
// [kCMHEVCTemporalLevelInfoKey_LevelIndex]: https://developer.apple.com/documentation/CoreMedia/kCMHEVCTemporalLevelInfoKey_LevelIndex
func (m MEHEVCDependencyInfo) LevelIndex() int16 {
	rv := objc.Send[int16](m.ID, objc.Sel("levelIndex"))
	return rv
}
func (m MEHEVCDependencyInfo) SetLevelIndex(value int16) {
	objc.Send[struct{}](m.ID, objc.Sel("setLevelIndex:"), value)
}
