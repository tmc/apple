// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioFile] class.
var (
	_AVAudioFileClass     AVAudioFileClass
	_AVAudioFileClassOnce sync.Once
)

func getAVAudioFileClass() AVAudioFileClass {
	_AVAudioFileClassOnce.Do(func() {
		_AVAudioFileClass = AVAudioFileClass{class: objc.GetClass("AVAudioFile")}
	})
	return _AVAudioFileClass
}

// GetAVAudioFileClass returns the class object for AVAudioFile.
func GetAVAudioFileClass() AVAudioFileClass {
	return getAVAudioFileClass()
}

type AVAudioFileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioFileClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioFileClass) Alloc() AVAudioFile {
	rv := objc.Send[AVAudioFile](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioFile.Url]
//   - [AVAudioFile.InitSecondaryReaderFormatError]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile
type AVAudioFile struct {
	objectivec.Object
}

// AVAudioFileFromID constructs a [AVAudioFile] from an objc.ID.
func AVAudioFileFromID(id objc.ID) AVAudioFile {
	return AVAudioFile{objectivec.Object{ID: id}}
}

// Ensure AVAudioFile implements IAVAudioFile.
var _ IAVAudioFile = AVAudioFile{}

// An interface definition for the [AVAudioFile] class.
//
// # Methods
//
//   - [IAVAudioFile.Url]
//   - [IAVAudioFile.InitSecondaryReaderFormatError]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile
type IAVAudioFile interface {
	objectivec.IObject

	// Topic: Methods

	Url() foundation.INSURL
	InitSecondaryReaderFormatError(reader objectivec.IObject, format objectivec.IObject) (AVAudioFile, error)
}

// Init initializes the instance.
func (a AVAudioFile) Init() AVAudioFile {
	rv := objc.Send[AVAudioFile](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioFile) Autorelease() AVAudioFile {
	rv := objc.Send[AVAudioFile](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioFile creates a new AVAudioFile instance.
func NewAVAudioFile() AVAudioFile {
	class := getAVAudioFileClass()
	rv := objc.Send[AVAudioFile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/initSecondaryReader:format:error:
func NewAudioFileSecondaryReaderFormatError(reader objectivec.IObject, format objectivec.IObject) (AVAudioFile, error) {
	var errorPtr objc.ID
	instance := getAVAudioFileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initSecondaryReader:format:error:"), reader, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/initSecondaryReader:format:error:
func (a AVAudioFile) InitSecondaryReaderFormatError(reader objectivec.IObject, format objectivec.IObject) (AVAudioFile, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initSecondaryReader:format:error:"), reader, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/url
func (a AVAudioFile) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
