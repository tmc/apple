// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NNLayerInfo] class.
var (
	_NNLayerInfoClass     NNLayerInfoClass
	_NNLayerInfoClassOnce sync.Once
)

func getNNLayerInfoClass() NNLayerInfoClass {
	_NNLayerInfoClassOnce.Do(func() {
		_NNLayerInfoClass = NNLayerInfoClass{class: objc.GetClass("_NNLayerInfo")}
	})
	return _NNLayerInfoClass
}

// GetNNLayerInfoClass returns the class object for _NNLayerInfo.
func GetNNLayerInfoClass() NNLayerInfoClass {
	return getNNLayerInfoClass()
}

type NNLayerInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NNLayerInfoClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NNLayerInfoClass) Alloc() NNLayerInfo {
	rv := objc.Send[NNLayerInfo](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [NNLayerInfo.Bidirectional]
//   - [NNLayerInfo.ConcatenatedInputNames]
//   - [NNLayerInfo.Type]
//   - [NNLayerInfo.InitWithTypeConcatenatedInputNamesBidirectional]
// See: https://developer.apple.com/documentation/CoreML/_NNLayerInfo
type NNLayerInfo struct {
	objectivec.Object
}

// NNLayerInfoFromID constructs a [NNLayerInfo] from an objc.ID.
func NNLayerInfoFromID(id objc.ID) NNLayerInfo {
	return NNLayerInfo{objectivec.Object{ID: id}}
}
// Ensure NNLayerInfo implements INNLayerInfo.
var _ INNLayerInfo = NNLayerInfo{}

// An interface definition for the [NNLayerInfo] class.
//
// # Methods
//
//   - [INNLayerInfo.Bidirectional]
//   - [INNLayerInfo.ConcatenatedInputNames]
//   - [INNLayerInfo.Type]
//   - [INNLayerInfo.InitWithTypeConcatenatedInputNamesBidirectional]
//
// See: https://developer.apple.com/documentation/CoreML/_NNLayerInfo
type INNLayerInfo interface {
	objectivec.IObject

	// Topic: Methods

	Bidirectional() bool
	ConcatenatedInputNames() string
	Type() string
	InitWithTypeConcatenatedInputNamesBidirectional(type_ objectivec.IObject, names objectivec.IObject, bidirectional bool) NNLayerInfo
}

// Init initializes the instance.
func (n NNLayerInfo) Init() NNLayerInfo {
	rv := objc.Send[NNLayerInfo](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NNLayerInfo) Autorelease() NNLayerInfo {
	rv := objc.Send[NNLayerInfo](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNNLayerInfo creates a new NNLayerInfo instance.
func NewNNLayerInfo() NNLayerInfo {
	class := getNNLayerInfoClass()
	rv := objc.Send[NNLayerInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/_NNLayerInfo/initWithType:concatenatedInputNames:bidirectional:
func NewNNLayerInfoWithTypeConcatenatedInputNamesBidirectional(type_ objectivec.IObject, names objectivec.IObject, bidirectional bool) NNLayerInfo {
	instance := getNNLayerInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:concatenatedInputNames:bidirectional:"), type_, names, bidirectional)
	return NNLayerInfoFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/_NNLayerInfo/initWithType:concatenatedInputNames:bidirectional:
func (n NNLayerInfo) InitWithTypeConcatenatedInputNamesBidirectional(type_ objectivec.IObject, names objectivec.IObject, bidirectional bool) NNLayerInfo {
	rv := objc.Send[NNLayerInfo](n.ID, objc.Sel("initWithType:concatenatedInputNames:bidirectional:"), type_, names, bidirectional)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_NNLayerInfo/bidirectional
func (n NNLayerInfo) Bidirectional() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("bidirectional"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/_NNLayerInfo/concatenatedInputNames
func (n NNLayerInfo) ConcatenatedInputNames() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("concatenatedInputNames"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/_NNLayerInfo/type
func (n NNLayerInfo) Type() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}

