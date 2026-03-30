// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MyMPSCNNConvolutionGradientState] class.
var (
	_MyMPSCNNConvolutionGradientStateClass     MyMPSCNNConvolutionGradientStateClass
	_MyMPSCNNConvolutionGradientStateClassOnce sync.Once
)

func getMyMPSCNNConvolutionGradientStateClass() MyMPSCNNConvolutionGradientStateClass {
	_MyMPSCNNConvolutionGradientStateClassOnce.Do(func() {
		_MyMPSCNNConvolutionGradientStateClass = MyMPSCNNConvolutionGradientStateClass{class: objc.GetClass("MyMPSCNNConvolutionGradientState")}
	})
	return _MyMPSCNNConvolutionGradientStateClass
}

// GetMyMPSCNNConvolutionGradientStateClass returns the class object for MyMPSCNNConvolutionGradientState.
func GetMyMPSCNNConvolutionGradientStateClass() MyMPSCNNConvolutionGradientStateClass {
	return getMyMPSCNNConvolutionGradientStateClass()
}

type MyMPSCNNConvolutionGradientStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MyMPSCNNConvolutionGradientStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MyMPSCNNConvolutionGradientStateClass) Alloc() MyMPSCNNConvolutionGradientState {
	rv := objc.Send[MyMPSCNNConvolutionGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MyMPSCNNConvolutionGradientState.SourceHeight]
//   - [MyMPSCNNConvolutionGradientState.SourceWidth]
//
// See: https://developer.apple.com/documentation/Espresso/MyMPSCNNConvolutionGradientState
type MyMPSCNNConvolutionGradientState struct {
	objectivec.Object
}

// MyMPSCNNConvolutionGradientStateFromID constructs a [MyMPSCNNConvolutionGradientState] from an objc.ID.
func MyMPSCNNConvolutionGradientStateFromID(id objc.ID) MyMPSCNNConvolutionGradientState {
	return MyMPSCNNConvolutionGradientState{objectivec.Object{ID: id}}
}

// NOTE: MyMPSCNNConvolutionGradientState struct embeds objectivec.Object (parent type unavailable) but
// IMyMPSCNNConvolutionGradientState embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MyMPSCNNConvolutionGradientState] class.
//
// # Methods
//
//   - [IMyMPSCNNConvolutionGradientState.SourceHeight]
//   - [IMyMPSCNNConvolutionGradientState.SourceWidth]
//
// See: https://developer.apple.com/documentation/Espresso/MyMPSCNNConvolutionGradientState
type IMyMPSCNNConvolutionGradientState interface {
	IMPSCNNConvolutionGradientState

	// Topic: Methods

	SourceHeight() uint64
	SourceWidth() uint64
}

// Init initializes the instance.
func (m MyMPSCNNConvolutionGradientState) Init() MyMPSCNNConvolutionGradientState {
	rv := objc.Send[MyMPSCNNConvolutionGradientState](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MyMPSCNNConvolutionGradientState) Autorelease() MyMPSCNNConvolutionGradientState {
	rv := objc.Send[MyMPSCNNConvolutionGradientState](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMyMPSCNNConvolutionGradientState creates a new MyMPSCNNConvolutionGradientState instance.
func NewMyMPSCNNConvolutionGradientState() MyMPSCNNConvolutionGradientState {
	class := getMyMPSCNNConvolutionGradientStateClass()
	rv := objc.Send[MyMPSCNNConvolutionGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/MyMPSCNNConvolutionGradientState/sourceHeight
func (m MyMPSCNNConvolutionGradientState) SourceHeight() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("sourceHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/MyMPSCNNConvolutionGradientState/sourceWidth
func (m MyMPSCNNConvolutionGradientState) SourceWidth() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("sourceWidth"))
	return rv
}
