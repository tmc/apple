// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Manager] class.
var (
	_ManagerClass     ManagerClass
	_ManagerClassOnce sync.Once
)

func getManagerClass() ManagerClass {
	_ManagerClassOnce.Do(func() {
		_ManagerClass = ManagerClass{class: objc.GetClass("_TtCO12TextToSpeech14TTSAudioEffect7Manager")}
	})
	return _ManagerClass
}

// GetManagerClass returns the class object for _TtCO12TextToSpeech14TTSAudioEffect7Manager.
func GetManagerClass() ManagerClass {
	return getManagerClass()
}

type ManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc ManagerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc ManagerClass) Alloc() Manager {
	rv := objc.Send[Manager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type Manager struct {
	objectivec.Object
}

// ManagerFromID constructs a [Manager] from an objc.ID.
func ManagerFromID(id objc.ID) Manager {
	return Manager{objectivec.Object{ID: id}}
}

// NOTE: Manager struct embeds objectivec.Object (parent type unavailable) but
// IManager embeds the parent interface; skip compile-time assertion.

// An interface definition for the [Manager] class.
type IManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m Manager) Init() Manager {
	rv := objc.Send[Manager](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m Manager) Autorelease() Manager {
	rv := objc.Send[Manager](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewManager creates a new Manager instance.
func NewManager() Manager {
	class := getManagerClass()
	rv := objc.Send[Manager](objc.ID(class.class), objc.Sel("new"))
	return rv
}
