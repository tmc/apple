// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/metalperformanceshaders"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MyMPSCNNConvolutionGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MyMPSCNNConvolutionGradientState struct {
	metalperformanceshaders.MPSCNNConvolutionGradientState
}

// MyMPSCNNConvolutionGradientStateFromID constructs a [MyMPSCNNConvolutionGradientState] from an objc.ID.
func MyMPSCNNConvolutionGradientStateFromID(id objc.ID) MyMPSCNNConvolutionGradientState {
	return MyMPSCNNConvolutionGradientState{MPSCNNConvolutionGradientState: metalperformanceshaders.MPSCNNConvolutionGradientStateFromID(id)}
}

// Ensure MyMPSCNNConvolutionGradientState implements IMyMPSCNNConvolutionGradientState.
var _ IMyMPSCNNConvolutionGradientState = MyMPSCNNConvolutionGradientState{}

// An interface definition for the [MyMPSCNNConvolutionGradientState] class.
type IMyMPSCNNConvolutionGradientState interface {
	metalperformanceshaders.IMPSCNNConvolutionGradientState
}

// Init initializes the instance.
func (m MyMPSCNNConvolutionGradientState) Init() MyMPSCNNConvolutionGradientState {
	rv := objc.SendIfResponds[MyMPSCNNConvolutionGradientState](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MyMPSCNNConvolutionGradientState) Autorelease() MyMPSCNNConvolutionGradientState {
	rv := objc.SendIfResponds[MyMPSCNNConvolutionGradientState](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMyMPSCNNConvolutionGradientState creates a new MyMPSCNNConvolutionGradientState instance.
func NewMyMPSCNNConvolutionGradientState() MyMPSCNNConvolutionGradientState {
	class := getMyMPSCNNConvolutionGradientStateClass()
	rv := objc.SendIfResponds[MyMPSCNNConvolutionGradientState](objc.ID(class.class), objc.Sel("new"))
	return rv
}
