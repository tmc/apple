// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSWMBridgedWindowChildInfo] class.
var (
	_SLSWMBridgedWindowChildInfoClass     SLSWMBridgedWindowChildInfoClass
	_SLSWMBridgedWindowChildInfoClassOnce sync.Once
)

func getSLSWMBridgedWindowChildInfoClass() SLSWMBridgedWindowChildInfoClass {
	_SLSWMBridgedWindowChildInfoClassOnce.Do(func() {
		_SLSWMBridgedWindowChildInfoClass = SLSWMBridgedWindowChildInfoClass{class: objc.GetClass("SLSWMBridgedWindowChildInfo")}
	})
	return _SLSWMBridgedWindowChildInfoClass
}

// GetSLSWMBridgedWindowChildInfoClass returns the class object for SLSWMBridgedWindowChildInfo.
func GetSLSWMBridgedWindowChildInfoClass() SLSWMBridgedWindowChildInfoClass {
	return getSLSWMBridgedWindowChildInfoClass()
}

type SLSWMBridgedWindowChildInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSWMBridgedWindowChildInfoClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSWMBridgedWindowChildInfoClass) Alloc() SLSWMBridgedWindowChildInfo {
	rv := objc.Send[SLSWMBridgedWindowChildInfo](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSWMBridgedWindowChildInfo.Op]
//   - [SLSWMBridgedWindowChildInfo.SetOp]
//   - [SLSWMBridgedWindowChildInfo.Window]
//   - [SLSWMBridgedWindowChildInfo.SetWindow]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindowChildInfo
type SLSWMBridgedWindowChildInfo struct {
	objectivec.Object
}

// SLSWMBridgedWindowChildInfoFromID constructs a [SLSWMBridgedWindowChildInfo] from an objc.ID.
func SLSWMBridgedWindowChildInfoFromID(id objc.ID) SLSWMBridgedWindowChildInfo {
	return SLSWMBridgedWindowChildInfo{objectivec.Object{ID: id}}
}

// Ensure SLSWMBridgedWindowChildInfo implements ISLSWMBridgedWindowChildInfo.
var _ ISLSWMBridgedWindowChildInfo = SLSWMBridgedWindowChildInfo{}

// An interface definition for the [SLSWMBridgedWindowChildInfo] class.
//
// # Methods
//
//   - [ISLSWMBridgedWindowChildInfo.Op]
//   - [ISLSWMBridgedWindowChildInfo.SetOp]
//   - [ISLSWMBridgedWindowChildInfo.Window]
//   - [ISLSWMBridgedWindowChildInfo.SetWindow]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindowChildInfo
type ISLSWMBridgedWindowChildInfo interface {
	objectivec.IObject

	// Topic: Methods

	Op() int
	SetOp(value int)
	Window() ISLSWMBridgedWindow
	SetWindow(value ISLSWMBridgedWindow)
}

// Init initializes the instance.
func (s SLSWMBridgedWindowChildInfo) Init() SLSWMBridgedWindowChildInfo {
	rv := objc.Send[SLSWMBridgedWindowChildInfo](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSWMBridgedWindowChildInfo) Autorelease() SLSWMBridgedWindowChildInfo {
	rv := objc.Send[SLSWMBridgedWindowChildInfo](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSWMBridgedWindowChildInfo creates a new SLSWMBridgedWindowChildInfo instance.
func NewSLSWMBridgedWindowChildInfo() SLSWMBridgedWindowChildInfo {
	class := getSLSWMBridgedWindowChildInfoClass()
	rv := objc.Send[SLSWMBridgedWindowChildInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindowChildInfo/op
func (s SLSWMBridgedWindowChildInfo) Op() int {
	rv := objc.Send[int](s.ID, objc.Sel("op"))
	return rv
}
func (s SLSWMBridgedWindowChildInfo) SetOp(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setOp:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindowChildInfo/window
func (s SLSWMBridgedWindowChildInfo) Window() ISLSWMBridgedWindow {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("window"))
	return SLSWMBridgedWindowFromID(objc.ID(rv))
}
func (s SLSWMBridgedWindowChildInfo) SetWindow(value ISLSWMBridgedWindow) {
	objc.Send[struct{}](s.ID, objc.Sel("setWindow:"), value)
}
