// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SODownloadDisplayManager] class.
var (
	_SODownloadDisplayManagerClass     SODownloadDisplayManagerClass
	_SODownloadDisplayManagerClassOnce sync.Once
)

func getSODownloadDisplayManagerClass() SODownloadDisplayManagerClass {
	_SODownloadDisplayManagerClassOnce.Do(func() {
		_SODownloadDisplayManagerClass = SODownloadDisplayManagerClass{class: objc.GetClass("SODownloadDisplayManager")}
	})
	return _SODownloadDisplayManagerClass
}

// GetSODownloadDisplayManagerClass returns the class object for SODownloadDisplayManager.
func GetSODownloadDisplayManagerClass() SODownloadDisplayManagerClass {
	return getSODownloadDisplayManagerClass()
}

type SODownloadDisplayManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SODownloadDisplayManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SODownloadDisplayManagerClass) Alloc() SODownloadDisplayManager {
	rv := objc.Send[SODownloadDisplayManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SODownloadDisplayManager._appendAveragingTimeInterval]
//   - [SODownloadDisplayManager._averagedTimeInterval]
//   - [SODownloadDisplayManager.Reset]
//   - [SODownloadDisplayManager.TimeRemainingForActiveInstallationsWithTagPrefix]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SODownloadDisplayManager
type SODownloadDisplayManager struct {
	objectivec.Object
}

// SODownloadDisplayManagerFromID constructs a [SODownloadDisplayManager] from an objc.ID.
func SODownloadDisplayManagerFromID(id objc.ID) SODownloadDisplayManager {
	return SODownloadDisplayManager{objectivec.Object{ID: id}}
}

// Ensure SODownloadDisplayManager implements ISODownloadDisplayManager.
var _ ISODownloadDisplayManager = SODownloadDisplayManager{}

// An interface definition for the [SODownloadDisplayManager] class.
//
// # Methods
//
//   - [ISODownloadDisplayManager._appendAveragingTimeInterval]
//   - [ISODownloadDisplayManager._averagedTimeInterval]
//   - [ISODownloadDisplayManager.Reset]
//   - [ISODownloadDisplayManager.TimeRemainingForActiveInstallationsWithTagPrefix]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SODownloadDisplayManager
type ISODownloadDisplayManager interface {
	objectivec.IObject

	// Topic: Methods

	_appendAveragingTimeInterval(interval float64)
	_averagedTimeInterval() float64
	Reset()
	TimeRemainingForActiveInstallationsWithTagPrefix(installations objectivec.IObject, prefix objectivec.IObject) float64
}

// Init initializes the instance.
func (s SODownloadDisplayManager) Init() SODownloadDisplayManager {
	rv := objc.Send[SODownloadDisplayManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SODownloadDisplayManager) Autorelease() SODownloadDisplayManager {
	rv := objc.Send[SODownloadDisplayManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSODownloadDisplayManager creates a new SODownloadDisplayManager instance.
func NewSODownloadDisplayManager() SODownloadDisplayManager {
	class := getSODownloadDisplayManagerClass()
	rv := objc.Send[SODownloadDisplayManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SODownloadDisplayManager/_appendAveragingTimeInterval:
func (s SODownloadDisplayManager) _appendAveragingTimeInterval(interval float64) {
	objc.Send[objc.ID](s.ID, objc.Sel("_appendAveragingTimeInterval:"), interval)
}

// AppendAveragingTimeInterval is an exported wrapper for the private method _appendAveragingTimeInterval.
func (s SODownloadDisplayManager) AppendAveragingTimeInterval(interval float64) {
	s._appendAveragingTimeInterval(interval)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SODownloadDisplayManager/_averagedTimeInterval
func (s SODownloadDisplayManager) _averagedTimeInterval() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("_averagedTimeInterval"))
	return rv
}

// AveragedTimeInterval is an exported wrapper for the private method _averagedTimeInterval.
func (s SODownloadDisplayManager) AveragedTimeInterval() float64 {
	return s._averagedTimeInterval()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SODownloadDisplayManager/reset
func (s SODownloadDisplayManager) Reset() {
	objc.Send[objc.ID](s.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/SpeechObjects/SODownloadDisplayManager/timeRemainingForActiveInstallations:withTagPrefix:
func (s SODownloadDisplayManager) TimeRemainingForActiveInstallationsWithTagPrefix(installations objectivec.IObject, prefix objectivec.IObject) float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("timeRemainingForActiveInstallations:withTagPrefix:"), installations, prefix)
	return rv
}
