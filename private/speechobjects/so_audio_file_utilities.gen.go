// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOAudioFileUtilities] class.
var (
	_SOAudioFileUtilitiesClass     SOAudioFileUtilitiesClass
	_SOAudioFileUtilitiesClassOnce sync.Once
)

func getSOAudioFileUtilitiesClass() SOAudioFileUtilitiesClass {
	_SOAudioFileUtilitiesClassOnce.Do(func() {
		_SOAudioFileUtilitiesClass = SOAudioFileUtilitiesClass{class: objc.GetClass("SOAudioFileUtilities")}
	})
	return _SOAudioFileUtilitiesClass
}

// GetSOAudioFileUtilitiesClass returns the class object for SOAudioFileUtilities.
func GetSOAudioFileUtilitiesClass() SOAudioFileUtilitiesClass {
	return getSOAudioFileUtilitiesClass()
}

type SOAudioFileUtilitiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOAudioFileUtilitiesClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOAudioFileUtilitiesClass) Alloc() SOAudioFileUtilities {
	rv := objc.Send[SOAudioFileUtilities](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOAudioFileUtilities
type SOAudioFileUtilities struct {
	objectivec.Object
}

// SOAudioFileUtilitiesFromID constructs a [SOAudioFileUtilities] from an objc.ID.
func SOAudioFileUtilitiesFromID(id objc.ID) SOAudioFileUtilities {
	return SOAudioFileUtilities{objectivec.Object{ID: id}}
}

// Ensure SOAudioFileUtilities implements ISOAudioFileUtilities.
var _ ISOAudioFileUtilities = SOAudioFileUtilities{}

// An interface definition for the [SOAudioFileUtilities] class.
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOAudioFileUtilities
type ISOAudioFileUtilities interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SOAudioFileUtilities) Init() SOAudioFileUtilities {
	rv := objc.Send[SOAudioFileUtilities](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOAudioFileUtilities) Autorelease() SOAudioFileUtilities {
	rv := objc.Send[SOAudioFileUtilities](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOAudioFileUtilities creates a new SOAudioFileUtilities instance.
func NewSOAudioFileUtilities() SOAudioFileUtilities {
	class := getSOAudioFileUtilitiesClass()
	rv := objc.Send[SOAudioFileUtilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOAudioFileUtilities/sampleDataFromContentsOfFile:streamDescription:
func (_SOAudioFileUtilitiesClass SOAudioFileUtilitiesClass) SampleDataFromContentsOfFileStreamDescription(file objectivec.IObject, description unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOAudioFileUtilitiesClass.class), objc.Sel("sampleDataFromContentsOfFile:streamDescription:"), file, description)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOAudioFileUtilities/writeSampleData:toFile:dataStreamDescription:fileStreamDescription:
func (_SOAudioFileUtilitiesClass SOAudioFileUtilitiesClass) WriteSampleDataToFileDataStreamDescriptionFileStreamDescription(data objectivec.IObject, file objectivec.IObject, description unsafe.Pointer, description2 unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(_SOAudioFileUtilitiesClass.class), objc.Sel("writeSampleData:toFile:dataStreamDescription:fileStreamDescription:"), data, file, description, description2)
	return rv
}
