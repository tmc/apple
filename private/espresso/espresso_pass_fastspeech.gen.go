// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFastspeech] class.
var (
	_EspressoPassFastspeechClass     EspressoPassFastspeechClass
	_EspressoPassFastspeechClassOnce sync.Once
)

func getEspressoPassFastspeechClass() EspressoPassFastspeechClass {
	_EspressoPassFastspeechClassOnce.Do(func() {
		_EspressoPassFastspeechClass = EspressoPassFastspeechClass{class: objc.GetClass("EspressoPass_fastspeech")}
	})
	return _EspressoPassFastspeechClass
}

// GetEspressoPassFastspeechClass returns the class object for EspressoPass_fastspeech.
func GetEspressoPassFastspeechClass() EspressoPassFastspeechClass {
	return getEspressoPassFastspeechClass()
}

type EspressoPassFastspeechClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFastspeechClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFastspeechClass) Alloc() EspressoPassFastspeech {
	rv := objc.SendIfResponds[EspressoPassFastspeech](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFastspeech struct {
	EspressoCustomPass
}

// EspressoPassFastspeechFromID constructs a [EspressoPassFastspeech] from an objc.ID.
func EspressoPassFastspeechFromID(id objc.ID) EspressoPassFastspeech {
	return EspressoPassFastspeech{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fastspeechFromID is an alias for [EspressoPassFastspeechFromID] for cross-framework compatibility.
func EspressoPass_fastspeechFromID(id objc.ID) EspressoPassFastspeech {
	return EspressoPassFastspeechFromID(id)
}

// Ensure EspressoPassFastspeech implements IEspressoPassFastspeech.
var _ IEspressoPassFastspeech = EspressoPassFastspeech{}

// An interface definition for the [EspressoPassFastspeech] class.
type IEspressoPassFastspeech interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFastspeech) Init() EspressoPassFastspeech {
	rv := objc.SendIfResponds[EspressoPassFastspeech](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFastspeech) Autorelease() EspressoPassFastspeech {
	rv := objc.SendIfResponds[EspressoPassFastspeech](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFastspeech creates a new EspressoPassFastspeech instance.
func NewEspressoPassFastspeech() EspressoPassFastspeech {
	class := getEspressoPassFastspeechClass()
	rv := objc.SendIfResponds[EspressoPassFastspeech](objc.ID(class.class), objc.Sel("new"))
	return rv
}
