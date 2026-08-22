// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[AVAudioSharedBufferToken](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioSharedBufferToken.EncodeWithCoder]
//   - [AVAudioSharedBufferToken.Surface]
//   - [AVAudioSharedBufferToken.SurfaceXPCType]
//   - [AVAudioSharedBufferToken.TaskToken]
//   - [AVAudioSharedBufferToken.TaskTokenXPCType]
//   - [AVAudioSharedBufferToken.InitWithCoder]
//   - [AVAudioSharedBufferToken.InitWithSurfaceTaskToken]
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
type IAVAudioSharedBufferToken interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	Surface() unsafe.Pointer
	SurfaceXPCType() XPCTypeSRef
	TaskToken() uint32
	TaskTokenXPCType() XPCTypeSRef
	InitWithCoder(coder foundation.INSCoder) AVAudioSharedBufferToken
	InitWithSurfaceTaskToken(surface iosurface.IOSurfaceRef, token uint32) AVAudioSharedBufferToken
}

// Init initializes the instance.
func (a AVAudioSharedBufferToken) Init() AVAudioSharedBufferToken {
	rv := objc.SendIfResponds[AVAudioSharedBufferToken](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSharedBufferToken) Autorelease() AVAudioSharedBufferToken {
	rv := objc.SendIfResponds[AVAudioSharedBufferToken](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSharedBufferToken creates a new AVAudioSharedBufferToken instance.
func NewAVAudioSharedBufferToken() AVAudioSharedBufferToken {
	class := getAVAudioSharedBufferTokenClass()
	rv := objc.SendIfResponds[AVAudioSharedBufferToken](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAudioSharedBufferTokenWithCoder(coder objectivec.IObject) AVAudioSharedBufferToken {
	instance := getAVAudioSharedBufferTokenClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AVAudioSharedBufferTokenFromID(rv)
}

func NewAudioSharedBufferTokenWithSurfaceTaskToken(surface iosurface.IOSurfaceRef, token uint32) AVAudioSharedBufferToken {
	instance := getAVAudioSharedBufferTokenClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSurface:taskToken:"), surface, token)
	return AVAudioSharedBufferTokenFromID(rv)
}

func (a AVAudioSharedBufferToken) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (a AVAudioSharedBufferToken) Surface() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("surface"))
	return rv
}
func (a AVAudioSharedBufferToken) SurfaceXPCType() XPCTypeSRef {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("surfaceXPCType"))
	return XPCTypeSRef(rv)
}
func (a AVAudioSharedBufferToken) TaskToken() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("taskToken"))
	return rv
}
func (a AVAudioSharedBufferToken) TaskTokenXPCType() XPCTypeSRef {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("taskTokenXPCType"))
	return XPCTypeSRef(rv)
}
func (a AVAudioSharedBufferToken) InitWithCoder(coder foundation.INSCoder) AVAudioSharedBufferToken {
	rv := objc.SendIfResponds[AVAudioSharedBufferToken](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a AVAudioSharedBufferToken) InitWithSurfaceTaskToken(surface iosurface.IOSurfaceRef, token uint32) AVAudioSharedBufferToken {
	rv := objc.SendIfResponds[AVAudioSharedBufferToken](a.ID, objc.Sel("initWithSurface:taskToken:"), surface, token)
	return rv
}

func (_AVAudioSharedBufferTokenClass AVAudioSharedBufferTokenClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVAudioSharedBufferTokenClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
