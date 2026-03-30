// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnitDSPGraph] class.
var (
	_AVAudioUnitDSPGraphClass     AVAudioUnitDSPGraphClass
	_AVAudioUnitDSPGraphClassOnce sync.Once
)

func getAVAudioUnitDSPGraphClass() AVAudioUnitDSPGraphClass {
	_AVAudioUnitDSPGraphClassOnce.Do(func() {
		_AVAudioUnitDSPGraphClass = AVAudioUnitDSPGraphClass{class: objc.GetClass("AVAudioUnitDSPGraph")}
	})
	return _AVAudioUnitDSPGraphClass
}

// GetAVAudioUnitDSPGraphClass returns the class object for AVAudioUnitDSPGraph.
func GetAVAudioUnitDSPGraphClass() AVAudioUnitDSPGraphClass {
	return getAVAudioUnitDSPGraphClass()
}

type AVAudioUnitDSPGraphClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitDSPGraphClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitDSPGraphClass) Alloc() AVAudioUnitDSPGraph {
	rv := objc.Send[AVAudioUnitDSPGraph](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioUnitDSPGraph.AuProcessingStripURL]
//   - [AVAudioUnitDSPGraph.DspGraphURL]
//   - [AVAudioUnitDSPGraph.LoadAudioDSPManager]
//   - [AVAudioUnitDSPGraph.LoadAudioUnitProcessingStripAtURLError]
//   - [AVAudioUnitDSPGraph.LoadDSPGraphAtURLError]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDSPGraph
type AVAudioUnitDSPGraph struct {
	AVAudioUnit
}

// AVAudioUnitDSPGraphFromID constructs a [AVAudioUnitDSPGraph] from an objc.ID.
func AVAudioUnitDSPGraphFromID(id objc.ID) AVAudioUnitDSPGraph {
	return AVAudioUnitDSPGraph{AVAudioUnit: AVAudioUnitFromID(id)}
}

// Ensure AVAudioUnitDSPGraph implements IAVAudioUnitDSPGraph.
var _ IAVAudioUnitDSPGraph = AVAudioUnitDSPGraph{}

// An interface definition for the [AVAudioUnitDSPGraph] class.
//
// # Methods
//
//   - [IAVAudioUnitDSPGraph.AuProcessingStripURL]
//   - [IAVAudioUnitDSPGraph.DspGraphURL]
//   - [IAVAudioUnitDSPGraph.LoadAudioDSPManager]
//   - [IAVAudioUnitDSPGraph.LoadAudioUnitProcessingStripAtURLError]
//   - [IAVAudioUnitDSPGraph.LoadDSPGraphAtURLError]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDSPGraph
type IAVAudioUnitDSPGraph interface {
	IAVAudioUnit

	// Topic: Methods

	AuProcessingStripURL() foundation.INSURL
	DspGraphURL() foundation.INSURL
	LoadAudioDSPManager() bool
	LoadAudioUnitProcessingStripAtURLError(url foundation.INSURL) (bool, error)
	LoadDSPGraphAtURLError(url foundation.INSURL) (bool, error)
}

// Init initializes the instance.
func (a AVAudioUnitDSPGraph) Init() AVAudioUnitDSPGraph {
	rv := objc.Send[AVAudioUnitDSPGraph](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitDSPGraph) Autorelease() AVAudioUnitDSPGraph {
	rv := objc.Send[AVAudioUnitDSPGraph](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitDSPGraph creates a new AVAudioUnitDSPGraph instance.
func NewAVAudioUnitDSPGraph() AVAudioUnitDSPGraph {
	class := getAVAudioUnitDSPGraphClass()
	rv := objc.Send[AVAudioUnitDSPGraph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioUnitDSPGraphWithImpl(impl unsafe.Pointer) AVAudioUnitDSPGraph {
	instance := getAVAudioUnitDSPGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioUnitDSPGraphFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDSPGraph/loadAudioDSPManager
func (a AVAudioUnitDSPGraph) LoadAudioDSPManager() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("loadAudioDSPManager"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDSPGraph/loadAudioUnitProcessingStripAtURL:error:
func (a AVAudioUnitDSPGraph) LoadAudioUnitProcessingStripAtURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadAudioUnitProcessingStripAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadAudioUnitProcessingStripAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDSPGraph/loadDSPGraphAtURL:error:
func (a AVAudioUnitDSPGraph) LoadDSPGraphAtURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadDSPGraphAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadDSPGraphAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDSPGraph/auProcessingStripURL
func (a AVAudioUnitDSPGraph) AuProcessingStripURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("auProcessingStripURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDSPGraph/dspGraphURL
func (a AVAudioUnitDSPGraph) DspGraphURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dspGraphURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
