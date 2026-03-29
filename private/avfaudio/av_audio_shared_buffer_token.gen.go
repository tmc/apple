// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioSharedBufferToken] class.
var (
	_AVAudioSharedBufferTokenClass     AVAudioSharedBufferTokenClass
	_AVAudioSharedBufferTokenClassOnce sync.Once
)

func getAVAudioSharedBufferTokenClass() AVAudioSharedBufferTokenClass {
	_AVAudioSharedBufferTokenClassOnce.Do(func() {
		_AVAudioSharedBufferTokenClass = AVAudioSharedBufferTokenClass{class: objc.GetClass("AVAudioSharedBufferToken")}
	})
	return _AVAudioSharedBufferTokenClass
}

// GetAVAudioSharedBufferTokenClass returns the class object for AVAudioSharedBufferToken.
func GetAVAudioSharedBufferTokenClass() AVAudioSharedBufferTokenClass {
	return getAVAudioSharedBufferTokenClass()
}

type AVAudioSharedBufferTokenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioSharedBufferTokenClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioSharedBufferTokenClass) Alloc() AVAudioSharedBufferToken {
	rv := objc.Send[AVAudioSharedBufferToken](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioSharedBufferToken.EncodeWithCoder]
//   - [AVAudioSharedBufferToken.Surface]
//   - [AVAudioSharedBufferToken.SurfaceXPCType]
//   - [AVAudioSharedBufferToken.TaskToken]
//   - [AVAudioSharedBufferToken.TaskTokenXPCType]
//   - [AVAudioSharedBufferToken.InitWithCoder]
//   - [AVAudioSharedBufferToken.InitWithSurfaceTaskToken]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken
type AVAudioSharedBufferToken struct {
	objectivec.Object
}

// AVAudioSharedBufferTokenFromID constructs a [AVAudioSharedBufferToken] from an objc.ID.
func AVAudioSharedBufferTokenFromID(id objc.ID) AVAudioSharedBufferToken {
	return AVAudioSharedBufferToken{objectivec.Object{ID: id}}
}
// Ensure AVAudioSharedBufferToken implements IAVAudioSharedBufferToken.
var _ IAVAudioSharedBufferToken = AVAudioSharedBufferToken{}

// An interface definition for the [AVAudioSharedBufferToken] class.
//
// # Methods
//
//   - [IAVAudioSharedBufferToken.EncodeWithCoder]
//   - [IAVAudioSharedBufferToken.Surface]
//   - [IAVAudioSharedBufferToken.SurfaceXPCType]
//   - [IAVAudioSharedBufferToken.TaskToken]
//   - [IAVAudioSharedBufferToken.TaskTokenXPCType]
//   - [IAVAudioSharedBufferToken.InitWithCoder]
//   - [IAVAudioSharedBufferToken.InitWithSurfaceTaskToken]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken
type IAVAudioSharedBufferToken interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	Surface() unsafe.Pointer
	SurfaceXPCType() objectivec.IObject
	TaskToken() objectivec.IObject
	TaskTokenXPCType() objectivec.IObject
	InitWithCoder(coder foundation.INSCoder) AVAudioSharedBufferToken
	InitWithSurfaceTaskToken(surface coregraphics.IOSurfaceRef, token uint32) AVAudioSharedBufferToken
}

// Init initializes the instance.
func (a AVAudioSharedBufferToken) Init() AVAudioSharedBufferToken {
	rv := objc.Send[AVAudioSharedBufferToken](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSharedBufferToken) Autorelease() AVAudioSharedBufferToken {
	rv := objc.Send[AVAudioSharedBufferToken](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSharedBufferToken creates a new AVAudioSharedBufferToken instance.
func NewAVAudioSharedBufferToken() AVAudioSharedBufferToken {
	class := getAVAudioSharedBufferTokenClass()
	rv := objc.Send[AVAudioSharedBufferToken](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/initWithCoder:
func NewAudioSharedBufferTokenWithCoder(coder objectivec.IObject) AVAudioSharedBufferToken {
	instance := getAVAudioSharedBufferTokenClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AVAudioSharedBufferTokenFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/initWithSurface:taskToken:
func NewAudioSharedBufferTokenWithSurfaceTaskToken(surface coregraphics.IOSurfaceRef, token uint32) AVAudioSharedBufferToken {
	instance := getAVAudioSharedBufferTokenClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSurface:taskToken:"), surface, token)
	return AVAudioSharedBufferTokenFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/encodeWithCoder:
func (a AVAudioSharedBufferToken) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/surface
func (a AVAudioSharedBufferToken) Surface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("surface"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/surfaceXPCType
func (a AVAudioSharedBufferToken) SurfaceXPCType() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("surfaceXPCType"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/taskToken
func (a AVAudioSharedBufferToken) TaskToken() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("taskToken"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/taskTokenXPCType
func (a AVAudioSharedBufferToken) TaskTokenXPCType() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("taskTokenXPCType"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/initWithCoder:
func (a AVAudioSharedBufferToken) InitWithCoder(coder foundation.INSCoder) AVAudioSharedBufferToken {
	rv := objc.Send[AVAudioSharedBufferToken](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/initWithSurface:taskToken:
func (a AVAudioSharedBufferToken) InitWithSurfaceTaskToken(surface coregraphics.IOSurfaceRef, token uint32) AVAudioSharedBufferToken {
	rv := objc.Send[AVAudioSharedBufferToken](a.ID, objc.Sel("initWithSurface:taskToken:"), surface, token)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedBufferToken/supportsSecureCoding
func (_AVAudioSharedBufferTokenClass AVAudioSharedBufferTokenClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVAudioSharedBufferTokenClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

