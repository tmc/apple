// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAssetBase] class.
var (
	_TTSAssetBaseClass     TTSAssetBaseClass
	_TTSAssetBaseClassOnce sync.Once
)

func getTTSAssetBaseClass() TTSAssetBaseClass {
	_TTSAssetBaseClassOnce.Do(func() {
		_TTSAssetBaseClass = TTSAssetBaseClass{class: objc.GetClass("TTSAssetBase")}
	})
	return _TTSAssetBaseClass
}

// GetTTSAssetBaseClass returns the class object for TTSAssetBase.
func GetTTSAssetBaseClass() TTSAssetBaseClass {
	return getTTSAssetBaseClass()
}

type TTSAssetBaseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAssetBaseClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAssetBaseClass) Alloc() TTSAssetBase {
	rv := objc.Send[TTSAssetBase](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSAssetBase.BundleIdentifier]
//   - [TTSAssetBase.SetBundleIdentifier]
//   - [TTSAssetBase.CompatibilityVersion]
//   - [TTSAssetBase.SetCompatibilityVersion]
//   - [TTSAssetBase.ContentVersion]
//   - [TTSAssetBase.SetContentVersion]
//   - [TTSAssetBase.EncodeWithCoder]
//   - [TTSAssetBase.MasteredVersion]
//   - [TTSAssetBase.SetMasteredVersion]
//   - [TTSAssetBase.InitWithCoder]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase
type TTSAssetBase struct {
	objectivec.Object
}

// TTSAssetBaseFromID constructs a [TTSAssetBase] from an objc.ID.
func TTSAssetBaseFromID(id objc.ID) TTSAssetBase {
	return TTSAssetBase{objectivec.Object{ID: id}}
}

// Ensure TTSAssetBase implements ITTSAssetBase.
var _ ITTSAssetBase = TTSAssetBase{}

// An interface definition for the [TTSAssetBase] class.
//
// # Methods
//
//   - [ITTSAssetBase.BundleIdentifier]
//   - [ITTSAssetBase.SetBundleIdentifier]
//   - [ITTSAssetBase.CompatibilityVersion]
//   - [ITTSAssetBase.SetCompatibilityVersion]
//   - [ITTSAssetBase.ContentVersion]
//   - [ITTSAssetBase.SetContentVersion]
//   - [ITTSAssetBase.EncodeWithCoder]
//   - [ITTSAssetBase.MasteredVersion]
//   - [ITTSAssetBase.SetMasteredVersion]
//   - [ITTSAssetBase.InitWithCoder]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase
type ITTSAssetBase interface {
	objectivec.IObject

	// Topic: Methods

	BundleIdentifier() string
	SetBundleIdentifier(value string)
	CompatibilityVersion() foundation.NSNumber
	SetCompatibilityVersion(value foundation.NSNumber)
	ContentVersion() foundation.NSNumber
	SetContentVersion(value foundation.NSNumber)
	EncodeWithCoder(coder foundation.INSCoder)
	MasteredVersion() string
	SetMasteredVersion(value string)
	InitWithCoder(coder foundation.INSCoder) TTSAssetBase
}

// Init initializes the instance.
func (t TTSAssetBase) Init() TTSAssetBase {
	rv := objc.Send[TTSAssetBase](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAssetBase) Autorelease() TTSAssetBase {
	rv := objc.Send[TTSAssetBase](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAssetBase creates a new TTSAssetBase instance.
func NewTTSAssetBase() TTSAssetBase {
	class := getTTSAssetBaseClass()
	rv := objc.Send[TTSAssetBase](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase/initWithCoder:
func NewTTSAssetBaseWithCoder(coder objectivec.IObject) TTSAssetBase {
	instance := getTTSAssetBaseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return TTSAssetBaseFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase/encodeWithCoder:
func (t TTSAssetBase) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase/initWithCoder:
func (t TTSAssetBase) InitWithCoder(coder foundation.INSCoder) TTSAssetBase {
	rv := objc.Send[TTSAssetBase](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase/supportsSecureCoding
func (_TTSAssetBaseClass TTSAssetBaseClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_TTSAssetBaseClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase/bundleIdentifier
func (t TTSAssetBase) BundleIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAssetBase) SetBundleIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase/compatibilityVersion
func (t TTSAssetBase) CompatibilityVersion() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("compatibilityVersion"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (t TTSAssetBase) SetCompatibilityVersion(value foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setCompatibilityVersion:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase/contentVersion
func (t TTSAssetBase) ContentVersion() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("contentVersion"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (t TTSAssetBase) SetContentVersion(value foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setContentVersion:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAssetBase/masteredVersion
func (t TTSAssetBase) MasteredVersion() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("masteredVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSAssetBase) SetMasteredVersion(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setMasteredVersion:"), objc.String(value))
}
