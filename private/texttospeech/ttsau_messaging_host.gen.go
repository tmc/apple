// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAUMessagingHost] class.
var (
	_TTSAUMessagingHostClass     TTSAUMessagingHostClass
	_TTSAUMessagingHostClassOnce sync.Once
)

func getTTSAUMessagingHostClass() TTSAUMessagingHostClass {
	_TTSAUMessagingHostClassOnce.Do(func() {
		_TTSAUMessagingHostClass = TTSAUMessagingHostClass{class: objc.GetClass("TTSAUMessagingHost")}
	})
	return _TTSAUMessagingHostClass
}

// GetTTSAUMessagingHostClass returns the class object for TTSAUMessagingHost.
func GetTTSAUMessagingHostClass() TTSAUMessagingHostClass {
	return getTTSAUMessagingHostClass()
}

type TTSAUMessagingHostClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAUMessagingHostClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAUMessagingHostClass) Alloc() TTSAUMessagingHost {
	rv := objc.Send[TTSAUMessagingHost](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSAUMessagingHost._loadProtocolMethods]
//   - [TTSAUMessagingHost.AllowedClasses]
//   - [TTSAUMessagingHost.SetAllowedClasses]
//   - [TTSAUMessagingHost.Channel]
//   - [TTSAUMessagingHost.SetChannel]
//   - [TTSAUMessagingHost.Methods]
//   - [TTSAUMessagingHost.SetMethods]
//   - [TTSAUMessagingHost.InitWithMessageChannel]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingHost
type TTSAUMessagingHost struct {
	objectivec.Object
}

// TTSAUMessagingHostFromID constructs a [TTSAUMessagingHost] from an objc.ID.
func TTSAUMessagingHostFromID(id objc.ID) TTSAUMessagingHost {
	return TTSAUMessagingHost{objectivec.Object{ID: id}}
}
// Ensure TTSAUMessagingHost implements ITTSAUMessagingHost.
var _ ITTSAUMessagingHost = TTSAUMessagingHost{}

// An interface definition for the [TTSAUMessagingHost] class.
//
// # Methods
//
//   - [ITTSAUMessagingHost._loadProtocolMethods]
//   - [ITTSAUMessagingHost.AllowedClasses]
//   - [ITTSAUMessagingHost.SetAllowedClasses]
//   - [ITTSAUMessagingHost.Channel]
//   - [ITTSAUMessagingHost.SetChannel]
//   - [ITTSAUMessagingHost.Methods]
//   - [ITTSAUMessagingHost.SetMethods]
//   - [ITTSAUMessagingHost.InitWithMessageChannel]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingHost
type ITTSAUMessagingHost interface {
	objectivec.IObject

	// Topic: Methods

	_loadProtocolMethods()
	AllowedClasses() foundation.INSSet
	SetAllowedClasses(value foundation.INSSet)
	Channel() objectivec.IObject
	SetChannel(value objectivec.IObject)
	Methods() foundation.INSDictionary
	SetMethods(value foundation.INSDictionary)
	InitWithMessageChannel(channel objectivec.IObject) TTSAUMessagingHost
}

// Init initializes the instance.
func (t TTSAUMessagingHost) Init() TTSAUMessagingHost {
	rv := objc.Send[TTSAUMessagingHost](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAUMessagingHost) Autorelease() TTSAUMessagingHost {
	rv := objc.Send[TTSAUMessagingHost](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAUMessagingHost creates a new TTSAUMessagingHost instance.
func NewTTSAUMessagingHost() TTSAUMessagingHost {
	class := getTTSAUMessagingHostClass()
	rv := objc.Send[TTSAUMessagingHost](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingHost/initWithMessageChannel:
func NewTTSAUMessagingHostWithMessageChannel(channel objectivec.IObject) TTSAUMessagingHost {
	instance := getTTSAUMessagingHostClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMessageChannel:"), channel)
	return TTSAUMessagingHostFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingHost/_loadProtocolMethods
func (t TTSAUMessagingHost) _loadProtocolMethods() {
	objc.Send[objc.ID](t.ID, objc.Sel("_loadProtocolMethods"))
}

// LoadProtocolMethods is an exported wrapper for the private method _loadProtocolMethods.
func (t TTSAUMessagingHost) LoadProtocolMethods() {
	t._loadProtocolMethods()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingHost/initWithMessageChannel:
func (t TTSAUMessagingHost) InitWithMessageChannel(channel objectivec.IObject) TTSAUMessagingHost {
	rv := objc.Send[TTSAUMessagingHost](t.ID, objc.Sel("initWithMessageChannel:"), channel)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingHost/_validSelectorsForProtocol
func (_TTSAUMessagingHostClass TTSAUMessagingHostClass) _validSelectorsForProtocol() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSAUMessagingHostClass.class), objc.Sel("_validSelectorsForProtocol"))
	return objectivec.Object{ID: rv}
}

// ValidSelectorsForProtocol is an exported wrapper for the private method _validSelectorsForProtocol.
func (_TTSAUMessagingHostClass TTSAUMessagingHostClass) ValidSelectorsForProtocol() objectivec.IObject {
	return _TTSAUMessagingHostClass._validSelectorsForProtocol()
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingHost/allowedClasses
func (t TTSAUMessagingHost) AllowedClasses() foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("allowedClasses"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (t TTSAUMessagingHost) SetAllowedClasses(value foundation.INSSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowedClasses:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingHost/channel
func (t TTSAUMessagingHost) Channel() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("channel"))
	return objectivec.Object{ID: rv}
}
func (t TTSAUMessagingHost) SetChannel(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setChannel:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingHost/methods
func (t TTSAUMessagingHost) Methods() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("methods"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSAUMessagingHost) SetMethods(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setMethods:"), value)
}

