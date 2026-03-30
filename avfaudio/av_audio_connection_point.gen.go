// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioConnectionPoint] class.
var (
	_AVAudioConnectionPointClass     AVAudioConnectionPointClass
	_AVAudioConnectionPointClassOnce sync.Once
)

func getAVAudioConnectionPointClass() AVAudioConnectionPointClass {
	_AVAudioConnectionPointClassOnce.Do(func() {
		_AVAudioConnectionPointClass = AVAudioConnectionPointClass{class: objc.GetClass("AVAudioConnectionPoint")}
	})
	return _AVAudioConnectionPointClass
}

// GetAVAudioConnectionPointClass returns the class object for AVAudioConnectionPoint.
func GetAVAudioConnectionPointClass() AVAudioConnectionPointClass {
	return getAVAudioConnectionPointClass()
}

type AVAudioConnectionPointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioConnectionPointClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioConnectionPointClass) Alloc() AVAudioConnectionPoint {
	rv := objc.Send[AVAudioConnectionPoint](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A representation of either a source or destination connection point in the
// audio engine.
//
// # Overview
//
// Instances of this class are immutable.
//
// # Creating a Connection Point
//
//   - [AVAudioConnectionPoint.InitWithNodeBus]: Creates a connection point object.
//
// # Getting Connection Point Properties
//
//   - [AVAudioConnectionPoint.InputConnectionPointForNodeInputBus]: Returns connection information about a node’s input bus.
//   - [AVAudioConnectionPoint.OutputConnectionPointsForNodeOutputBus]: Returns connection information about a node’s output bus.
//   - [AVAudioConnectionPoint.Bus]: The bus on the node in the connection point.
//   - [AVAudioConnectionPoint.Node]: The node in the connection point.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint
type AVAudioConnectionPoint struct {
	objectivec.Object
}

// AVAudioConnectionPointFromID constructs a [AVAudioConnectionPoint] from an objc.ID.
//
// A representation of either a source or destination connection point in the
// audio engine.
func AVAudioConnectionPointFromID(id objc.ID) AVAudioConnectionPoint {
	return AVAudioConnectionPoint{objectivec.Object{ID: id}}
}

// NOTE: AVAudioConnectionPoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioConnectionPoint] class.
//
// # Creating a Connection Point
//
//   - [IAVAudioConnectionPoint.InitWithNodeBus]: Creates a connection point object.
//
// # Getting Connection Point Properties
//
//   - [IAVAudioConnectionPoint.InputConnectionPointForNodeInputBus]: Returns connection information about a node’s input bus.
//   - [IAVAudioConnectionPoint.OutputConnectionPointsForNodeOutputBus]: Returns connection information about a node’s output bus.
//   - [IAVAudioConnectionPoint.Bus]: The bus on the node in the connection point.
//   - [IAVAudioConnectionPoint.Node]: The node in the connection point.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint
type IAVAudioConnectionPoint interface {
	objectivec.IObject

	// Topic: Creating a Connection Point

	// Creates a connection point object.
	InitWithNodeBus(node IAVAudioNode, bus AVAudioNodeBus) AVAudioConnectionPoint

	// Topic: Getting Connection Point Properties

	// Returns connection information about a node’s input bus.
	InputConnectionPointForNodeInputBus(node IAVAudioNode, bus AVAudioNodeBus) IAVAudioConnectionPoint
	// Returns connection information about a node’s output bus.
	OutputConnectionPointsForNodeOutputBus(node IAVAudioNode, bus AVAudioNodeBus) []AVAudioConnectionPoint
	// The bus on the node in the connection point.
	Bus() AVAudioNodeBus
	// The node in the connection point.
	Node() IAVAudioNode
}

// Init initializes the instance.
func (a AVAudioConnectionPoint) Init() AVAudioConnectionPoint {
	rv := objc.Send[AVAudioConnectionPoint](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioConnectionPoint) Autorelease() AVAudioConnectionPoint {
	rv := objc.Send[AVAudioConnectionPoint](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioConnectionPoint creates a new AVAudioConnectionPoint instance.
func NewAVAudioConnectionPoint() AVAudioConnectionPoint {
	class := getAVAudioConnectionPointClass()
	rv := objc.Send[AVAudioConnectionPoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a connection point object.
//
// node: The source or destination node.
//
// bus: The output or input bus on the node.
//
// # Discussion
//
// If the node is `nil`, this method fails and returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint/init(node:bus:)
func NewAudioConnectionPointWithNodeBus(node IAVAudioNode, bus AVAudioNodeBus) AVAudioConnectionPoint {
	instance := getAVAudioConnectionPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNode:bus:"), node, bus)
	return AVAudioConnectionPointFromID(rv)
}

// Creates a connection point object.
//
// node: The source or destination node.
//
// bus: The output or input bus on the node.
//
// # Discussion
//
// If the node is `nil`, this method fails and returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint/init(node:bus:)
func (a AVAudioConnectionPoint) InitWithNodeBus(node IAVAudioNode, bus AVAudioNodeBus) AVAudioConnectionPoint {
	rv := objc.Send[AVAudioConnectionPoint](a.ID, objc.Sel("initWithNode:bus:"), node, bus)
	return rv
}

// Returns connection information about a node’s input bus.
//
// node: The node with the input connection you’re querying.
//
// bus: The node’s input bus for the connection you’re querying.
//
// # Return Value
//
// An [AVAudioConnectionPoint] object with connection information on the
// node’s input bus.
//
// # Discussion
//
// Connections are always one-to-one or one-to-many. This method returns `nil`
// if there’s no connection on the node’s specified input bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/inputConnectionPoint(for:inputBus:)
func (a AVAudioConnectionPoint) InputConnectionPointForNodeInputBus(node IAVAudioNode, bus AVAudioNodeBus) IAVAudioConnectionPoint {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputConnectionPointForNode:inputBus:"), node, bus)
	return AVAudioConnectionPointFromID(rv)
}

// Returns connection information about a node’s output bus.
//
// node: The node with the output connections you’re querying.
//
// bus: The node’s output bus for connections you’re querying.
//
// # Return Value
//
// An array of [AVAudioConnectionPoint] objects with connection information on
// the node’s output bus.
//
// # Discussion
//
// Connections are always one-to-one or one-to-many. This method returns an
// empty array if there are no connections on the node’s specified output
// bus.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/outputConnectionPoints(for:outputBus:)
func (a AVAudioConnectionPoint) OutputConnectionPointsForNodeOutputBus(node IAVAudioNode, bus AVAudioNodeBus) []AVAudioConnectionPoint {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("outputConnectionPointsForNode:outputBus:"), node, bus)
	return objc.ConvertSlice(rv, func(id objc.ID) AVAudioConnectionPoint {
		return AVAudioConnectionPointFromID(id)
	})
}

// The bus on the node in the connection point.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint/bus
func (a AVAudioConnectionPoint) Bus() AVAudioNodeBus {
	rv := objc.Send[AVAudioNodeBus](a.ID, objc.Sel("bus"))
	return AVAudioNodeBus(rv)
}

// The node in the connection point.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint/node
func (a AVAudioConnectionPoint) Node() IAVAudioNode {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("node"))
	return AVAudioNodeFromID(objc.ID(rv))
}
