// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMutableVideoCompositionInstruction] class.
var (
	_AVMutableVideoCompositionInstructionClass     AVMutableVideoCompositionInstructionClass
	_AVMutableVideoCompositionInstructionClassOnce sync.Once
)

func getAVMutableVideoCompositionInstructionClass() AVMutableVideoCompositionInstructionClass {
	_AVMutableVideoCompositionInstructionClassOnce.Do(func() {
		_AVMutableVideoCompositionInstructionClass = AVMutableVideoCompositionInstructionClass{class: objc.GetClass("AVMutableVideoCompositionInstruction")}
	})
	return _AVMutableVideoCompositionInstructionClass
}

// GetAVMutableVideoCompositionInstructionClass returns the class object for AVMutableVideoCompositionInstruction.
func GetAVMutableVideoCompositionInstructionClass() AVMutableVideoCompositionInstructionClass {
	return getAVMutableVideoCompositionInstructionClass()
}

type AVMutableVideoCompositionInstructionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableVideoCompositionInstructionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableVideoCompositionInstructionClass) Alloc() AVMutableVideoCompositionInstruction {
	rv := objc.Send[AVMutableVideoCompositionInstruction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable video composition instruction subclass.
//
// # Overview
//
// An [AVVideoComposition] object maintains an array of [AVMutableVideoCompositionInstruction.Instructions] to
// perform its composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction
type AVMutableVideoCompositionInstruction struct {
	AVVideoCompositionInstruction
}

// AVMutableVideoCompositionInstructionFromID constructs a [AVMutableVideoCompositionInstruction] from an objc.ID.
//
// A mutable video composition instruction subclass.
func AVMutableVideoCompositionInstructionFromID(id objc.ID) AVMutableVideoCompositionInstruction {
	return AVMutableVideoCompositionInstruction{AVVideoCompositionInstruction: AVVideoCompositionInstructionFromID(id)}
}

// NOTE: AVMutableVideoCompositionInstruction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableVideoCompositionInstruction] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction
type IAVMutableVideoCompositionInstruction interface {
	IAVVideoCompositionInstruction
}

// Init initializes the instance.
func (m AVMutableVideoCompositionInstruction) Init() AVMutableVideoCompositionInstruction {
	rv := objc.Send[AVMutableVideoCompositionInstruction](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableVideoCompositionInstruction) Autorelease() AVMutableVideoCompositionInstruction {
	rv := objc.Send[AVMutableVideoCompositionInstruction](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableVideoCompositionInstruction creates a new AVMutableVideoCompositionInstruction instance.
func NewAVMutableVideoCompositionInstruction() AVMutableVideoCompositionInstruction {
	class := getAVMutableVideoCompositionInstructionClass()
	rv := objc.Send[AVMutableVideoCompositionInstruction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a new mutable video composition instruction.
//
// # Return Value
//
// A new mutable video composition instruction.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoCompositionInstruction/videoCompositionInstruction
func (_AVMutableVideoCompositionInstructionClass AVMutableVideoCompositionInstructionClass) VideoCompositionInstruction() AVMutableVideoCompositionInstruction {
	rv := objc.Send[objc.ID](objc.ID(_AVMutableVideoCompositionInstructionClass.class), objc.Sel("videoCompositionInstruction"))
	return AVMutableVideoCompositionInstructionFromID(rv)
}
