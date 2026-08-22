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
	rv := objc.SendIfResponds[DownloadableVoice](objc.ID(dc.class), objc.Sel("alloc"))
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
	rv := objc.SendIfResponds[DownloadableVoice](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DownloadableVoice) Autorelease() DownloadableVoice {
	rv := objc.SendIfResponds[DownloadableVoice](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDownloadableVoice creates a new DownloadableVoice instance.
func NewDownloadableVoice() DownloadableVoice {
	class := getDownloadableVoiceClass()
	rv := objc.SendIfResponds[DownloadableVoice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDownloadableVoiceWithVoiceIDProperties(id objectivec.IObject, properties objectivec.IObject) DownloadableVoice {
	instance := getDownloadableVoiceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithVoiceID:properties:"), id, properties)
	return DownloadableVoiceFromID(rv)
}

func NewDownloadableVoiceWithVoiceIdentifier(voice objectivec.IObject, identifier objectivec.IObject) DownloadableVoice {
	instance := getDownloadableVoiceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithVoice:identifier:"), voice, identifier)
	return DownloadableVoiceFromID(rv)
}

func (d DownloadableVoice) DisplayedSize() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("displayedSize"))
	return objectivec.Object{ID: rv}
}
func (d DownloadableVoice) InitWithVoiceIDProperties(id objectivec.IObject, properties objectivec.IObject) DownloadableVoice {
	rv := objc.SendIfResponds[DownloadableVoice](d.ID, objc.Sel("initWithVoiceID:properties:"), id, properties)
	return rv
}

func (d DownloadableVoice) ByteSize() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("byteSize"))
	return rv
}
func (d DownloadableVoice) CompactSizeBundleIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("compactSizeBundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (d DownloadableVoice) CompactSizeByteSize() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("compactSizeByteSize"))
	return rv
}
func (d DownloadableVoice) CompactSizeTagName() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("compactSizeTagName"))
	return foundation.NSStringFromID(rv).String()
}
func (d DownloadableVoice) CompactSizeVersion() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("compactSizeVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (d DownloadableVoice) DownloadCompactSize() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("downloadCompactSize"))
	return rv
}
func (d DownloadableVoice) SetDownloadCompactSize(value bool) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setDownloadCompactSize:"), value)
}
func (d DownloadableVoice) TagName() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("tagName"))
	return foundation.NSStringFromID(rv).String()
}
func (d DownloadableVoice) Variant() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("variant"))
	return foundation.NSStringFromID(rv).String()
}
func (d DownloadableVoice) Version() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("version"))
	return foundation.NSStringFromID(rv).String()
}
func (d DownloadableVoice) VoiceIdentifierToMarkAsPurgeableAfterInstall() string {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("voiceIdentifierToMarkAsPurgeableAfterInstall"))
	return foundation.NSStringFromID(rv).String()
}
