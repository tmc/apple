// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSTransactionPerMenuBarData] class.
var (
	_SLSTransactionPerMenuBarDataClass     SLSTransactionPerMenuBarDataClass
	_SLSTransactionPerMenuBarDataClassOnce sync.Once
)

func getSLSTransactionPerMenuBarDataClass() SLSTransactionPerMenuBarDataClass {
	_SLSTransactionPerMenuBarDataClassOnce.Do(func() {
		_SLSTransactionPerMenuBarDataClass = SLSTransactionPerMenuBarDataClass{class: objc.GetClass("SLSTransactionPerMenuBarData")}
	})
	return _SLSTransactionPerMenuBarDataClass
}

// GetSLSTransactionPerMenuBarDataClass returns the class object for SLSTransactionPerMenuBarData.
func GetSLSTransactionPerMenuBarDataClass() SLSTransactionPerMenuBarDataClass {
	return getSLSTransactionPerMenuBarDataClass()
}

type SLSTransactionPerMenuBarDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSTransactionPerMenuBarDataClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSTransactionPerMenuBarDataClass) Alloc() SLSTransactionPerMenuBarData {
	rv := objc.Send[SLSTransactionPerMenuBarData](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSTransactionPerMenuBarData
type SLSTransactionPerMenuBarData struct {
	objectivec.Object
}

// SLSTransactionPerMenuBarDataFromID constructs a [SLSTransactionPerMenuBarData] from an objc.ID.
func SLSTransactionPerMenuBarDataFromID(id objc.ID) SLSTransactionPerMenuBarData {
	return SLSTransactionPerMenuBarData{objectivec.Object{ID: id}}
}

// Ensure SLSTransactionPerMenuBarData implements ISLSTransactionPerMenuBarData.
var _ ISLSTransactionPerMenuBarData = SLSTransactionPerMenuBarData{}

// An interface definition for the [SLSTransactionPerMenuBarData] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSTransactionPerMenuBarData
type ISLSTransactionPerMenuBarData interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SLSTransactionPerMenuBarData) Init() SLSTransactionPerMenuBarData {
	rv := objc.Send[SLSTransactionPerMenuBarData](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSTransactionPerMenuBarData) Autorelease() SLSTransactionPerMenuBarData {
	rv := objc.Send[SLSTransactionPerMenuBarData](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSTransactionPerMenuBarData creates a new SLSTransactionPerMenuBarData instance.
func NewSLSTransactionPerMenuBarData() SLSTransactionPerMenuBarData {
	class := getSLSTransactionPerMenuBarDataClass()
	rv := objc.Send[SLSTransactionPerMenuBarData](objc.ID(class.class), objc.Sel("new"))
	return rv
}
