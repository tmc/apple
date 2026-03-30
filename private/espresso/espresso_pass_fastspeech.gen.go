// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fastspeech] class.
var (
	_EspressoPass_fastspeechClass     EspressoPass_fastspeechClass
	_EspressoPass_fastspeechClassOnce sync.Once
)

func getEspressoPass_fastspeechClass() EspressoPass_fastspeechClass {
	_EspressoPass_fastspeechClassOnce.Do(func() {
		_EspressoPass_fastspeechClass = EspressoPass_fastspeechClass{class: objc.GetClass("EspressoPass_fastspeech")}
	})
	return _EspressoPass_fastspeechClass
}

// GetEspressoPass_fastspeechClass returns the class object for EspressoPass_fastspeech.
func GetEspressoPass_fastspeechClass() EspressoPass_fastspeechClass {
	return getEspressoPass_fastspeechClass()
}

type EspressoPass_fastspeechClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPass_fastspeechClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fastspeechClass) Alloc() EspressoPass_fastspeech {
	rv := objc.Send[EspressoPass_fastspeech](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fastspeech
type EspressoPass_fastspeech struct {
	EspressoCustomPass
}

// EspressoPass_fastspeechFromID constructs a [EspressoPass_fastspeech] from an objc.ID.
func EspressoPass_fastspeechFromID(id objc.ID) EspressoPass_fastspeech {
	return EspressoPass_fastspeech{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// Ensure EspressoPass_fastspeech implements IEspressoPass_fastspeech.
var _ IEspressoPass_fastspeech = EspressoPass_fastspeech{}

// An interface definition for the [EspressoPass_fastspeech] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fastspeech
type IEspressoPass_fastspeech interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fastspeech) Init() EspressoPass_fastspeech {
	rv := objc.Send[EspressoPass_fastspeech](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fastspeech) Autorelease() EspressoPass_fastspeech {
	rv := objc.Send[EspressoPass_fastspeech](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fastspeech creates a new EspressoPass_fastspeech instance.
func NewEspressoPass_fastspeech() EspressoPass_fastspeech {
	class := getEspressoPass_fastspeechClass()
	rv := objc.Send[EspressoPass_fastspeech](objc.ID(class.class), objc.Sel("new"))
	return rv
}
