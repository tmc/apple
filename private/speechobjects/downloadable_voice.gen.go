// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DownloadableVoice] class.
var (
	_DownloadableVoiceClass     DownloadableVoiceClass
	_DownloadableVoiceClassOnce sync.Once
)

func getDownloadableVoiceClass() DownloadableVoiceClass {
	_DownloadableVoiceClassOnce.Do(func() {
		_DownloadableVoiceClass = DownloadableVoiceClass{class: objc.GetClass("DownloadableVoice")}
	})
	return _DownloadableVoiceClass
}

// GetDownloadableVoiceClass returns the class object for DownloadableVoice.
func GetDownloadableVoiceClass() DownloadableVoiceClass {
	return getDownloadableVoiceClass()
}

type DownloadableVoiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DownloadableVoiceClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DownloadableVoiceClass) Alloc() DownloadableVoice {
	rv := objc.Send[DownloadableVoice](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DownloadableVoice.ByteSize]
//   - [DownloadableVoice.CompactSizeBundleIdentifier]
//   - [DownloadableVoice.CompactSizeByteSize]
//   - [DownloadableVoice.CompactSizeTagName]
//   - [DownloadableVoice.CompactSizeVersion]
//   - [DownloadableVoice.DisplayedSize]
//   - [DownloadableVoice.DownloadCompactSize]
//   - [DownloadableVoice.SetDownloadCompactSize]
//   - [DownloadableVoice.TagName]
//   - [DownloadableVoice.Variant]
//   - [DownloadableVoice.VoiceIdentifierToMarkAsPurgeableAfterInstall]
//   - [DownloadableVoice.InitWithVoiceIDProperties]
//   - [DownloadableVoice.Version]
//
// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice
type DownloadableVoice struct {
	SOVoiceObject
}

// DownloadableVoiceFromID constructs a [DownloadableVoice] from an objc.ID.
func DownloadableVoiceFromID(id objc.ID) DownloadableVoice {
	return DownloadableVoice{SOVoiceObject: SOVoiceObjectFromID(id)}
}

// Ensure DownloadableVoice implements IDownloadableVoice.
var _ IDownloadableVoice = DownloadableVoice{}

// An interface definition for the [DownloadableVoice] class.
//
// # Methods
//
//   - [IDownloadableVoice.ByteSize]
//   - [IDownloadableVoice.CompactSizeBundleIdentifier]
//   - [IDownloadableVoice.CompactSizeByteSize]
//   - [IDownloadableVoice.CompactSizeTagName]
//   - [IDownloadableVoice.CompactSizeVersion]
//   - [IDownloadableVoice.DisplayedSize]
//   - [IDownloadableVoice.DownloadCompactSize]
//   - [IDownloadableVoice.SetDownloadCompactSize]
//   - [IDownloadableVoice.TagName]
//   - [IDownloadableVoice.Variant]
//   - [IDownloadableVoice.VoiceIdentifierToMarkAsPurgeableAfterInstall]
//   - [IDownloadableVoice.InitWithVoiceIDProperties]
//   - [IDownloadableVoice.Version]
//
// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice
type IDownloadableVoice interface {
	ISOVoiceObject

	// Topic: Methods

	ByteSize() uint64
	CompactSizeBundleIdentifier() string
	CompactSizeByteSize() uint64
	CompactSizeTagName() string
	CompactSizeVersion() string
	DisplayedSize() objectivec.IObject
	DownloadCompactSize() bool
	SetDownloadCompactSize(value bool)
	TagName() string
	Variant() string
	VoiceIdentifierToMarkAsPurgeableAfterInstall() string
	InitWithVoiceIDProperties(id objectivec.IObject, properties objectivec.IObject) DownloadableVoice
	Version() string
}

// Init initializes the instance.
func (d DownloadableVoice) Init() DownloadableVoice {
	rv := objc.Send[DownloadableVoice](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DownloadableVoice) Autorelease() DownloadableVoice {
	rv := objc.Send[DownloadableVoice](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDownloadableVoice creates a new DownloadableVoice instance.
func NewDownloadableVoice() DownloadableVoice {
	class := getDownloadableVoiceClass()
	rv := objc.Send[DownloadableVoice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/initWithVoiceID:properties:
func NewDownloadableVoiceWithVoiceIDProperties(id objectivec.IObject, properties objectivec.IObject) DownloadableVoice {
	instance := getDownloadableVoiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVoiceID:properties:"), id, properties)
	return DownloadableVoiceFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/initWithVoice:identifier:
func NewDownloadableVoiceWithVoiceIdentifier(voice objectivec.IObject, identifier objectivec.IObject) DownloadableVoice {
	instance := getDownloadableVoiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVoice:identifier:"), voice, identifier)
	return DownloadableVoiceFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/displayedSize
func (d DownloadableVoice) DisplayedSize() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("displayedSize"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/initWithVoiceID:properties:
func (d DownloadableVoice) InitWithVoiceIDProperties(id objectivec.IObject, properties objectivec.IObject) DownloadableVoice {
	rv := objc.Send[DownloadableVoice](d.ID, objc.Sel("initWithVoiceID:properties:"), id, properties)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/byteSize
func (d DownloadableVoice) ByteSize() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("byteSize"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/compactSizeBundleIdentifier
func (d DownloadableVoice) CompactSizeBundleIdentifier() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("compactSizeBundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/compactSizeByteSize
func (d DownloadableVoice) CompactSizeByteSize() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("compactSizeByteSize"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/compactSizeTagName
func (d DownloadableVoice) CompactSizeTagName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("compactSizeTagName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/compactSizeVersion
func (d DownloadableVoice) CompactSizeVersion() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("compactSizeVersion"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/downloadCompactSize
func (d DownloadableVoice) DownloadCompactSize() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("downloadCompactSize"))
	return rv
}
func (d DownloadableVoice) SetDownloadCompactSize(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setDownloadCompactSize:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/tagName
func (d DownloadableVoice) TagName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("tagName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/variant
func (d DownloadableVoice) Variant() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("variant"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/version
func (d DownloadableVoice) Version() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("version"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/DownloadableVoice/voiceIdentifierToMarkAsPurgeableAfterInstall
func (d DownloadableVoice) VoiceIdentifierToMarkAsPurgeableAfterInstall() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("voiceIdentifierToMarkAsPurgeableAfterInstall"))
	return foundation.NSStringFromID(rv).String()
}
