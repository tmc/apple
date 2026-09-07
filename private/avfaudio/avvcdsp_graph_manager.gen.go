// Code generated from Apple documentation for avfaudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCDSPGraphManager] class.
var (
	_AVVCDSPGraphManagerClass     AVVCDSPGraphManagerClass
	_AVVCDSPGraphManagerClassOnce sync.Once
)

func getAVVCDSPGraphManagerClass() AVVCDSPGraphManagerClass {
	_AVVCDSPGraphManagerClassOnce.Do(func() {
		_AVVCDSPGraphManagerClass = AVVCDSPGraphManagerClass{class: objc.GetClass("AVVCDSPGraphManager")}
	})
	return _AVVCDSPGraphManagerClass
}

// GetAVVCDSPGraphManagerClass returns the class object for AVVCDSPGraphManager.
func GetAVVCDSPGraphManagerClass() AVVCDSPGraphManagerClass {
	return getAVVCDSPGraphManagerClass()
}

type AVVCDSPGraphManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCDSPGraphManagerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCDSPGraphManagerClass) Alloc() AVVCDSPGraphManager {
	rv := objc.SendIfResponds[AVVCDSPGraphManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCDSPGraphManager.DSPGraphAU]
//   - [AVVCDSPGraphManager.SetDSPGraphAU]
//   - [AVVCDSPGraphManager.SetAggressiveECMode]
type AVVCDSPGraphManager struct {
	objectivec.Object
}

// AVVCDSPGraphManagerFromID constructs a [AVVCDSPGraphManager] from an objc.ID.
func AVVCDSPGraphManagerFromID(id objc.ID) AVVCDSPGraphManager {
	return AVVCDSPGraphManager{objectivec.Object{ID: id}}
}

// Ensure AVVCDSPGraphManager implements IAVVCDSPGraphManager.
var _ IAVVCDSPGraphManager = AVVCDSPGraphManager{}

// An interface definition for the [AVVCDSPGraphManager] class.
//
// # Methods
//
//   - [IAVVCDSPGraphManager.DSPGraphAU]
//   - [IAVVCDSPGraphManager.SetDSPGraphAU]
//   - [IAVVCDSPGraphManager.SetAggressiveECMode]
type IAVVCDSPGraphManager interface {
	objectivec.IObject

	// Topic: Methods

	DSPGraphAU() *ComponentInstanceRecord
	SetDSPGraphAU(value *ComponentInstanceRecord)
	SetAggressiveECMode(eCMode bool) int32
}

// Init initializes the instance.
func (a AVVCDSPGraphManager) Init() AVVCDSPGraphManager {
	rv := objc.SendIfResponds[AVVCDSPGraphManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCDSPGraphManager) Autorelease() AVVCDSPGraphManager {
	rv := objc.SendIfResponds[AVVCDSPGraphManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCDSPGraphManager creates a new AVVCDSPGraphManager instance.
func NewAVVCDSPGraphManager() AVVCDSPGraphManager {
	class := getAVVCDSPGraphManagerClass()
	rv := objc.SendIfResponds[AVVCDSPGraphManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVVCDSPGraphManager) SetAggressiveECMode(eCMode bool) int32 {
	rv := objc.SendIfResponds[int32](a.ID, objc.Sel("setAggressiveECMode:"), eCMode)
	return rv
}

func (_AVVCDSPGraphManagerClass AVVCDSPGraphManagerClass) SharedInstance() AVVCDSPGraphManager {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_AVVCDSPGraphManagerClass.class), objc.Sel("sharedInstance"))
	return AVVCDSPGraphManagerFromID(rv)
}

func (a AVVCDSPGraphManager) DSPGraphAU() *ComponentInstanceRecord {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("DSPGraphAU"))
	return (*ComponentInstanceRecord)(rv)
}
func (a AVVCDSPGraphManager) SetDSPGraphAU(value *ComponentInstanceRecord) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setDSPGraphAU:"), value)
}
