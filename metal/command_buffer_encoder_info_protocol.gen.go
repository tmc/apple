// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A container that provides additional information about a runtime failure a GPU encounters as it runs the commands in a command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferEncoderInfo
type MTLCommandBufferEncoderInfo interface {
	objectivec.IObject

	// The name of the encoder that generates the error information.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferEncoderInfo/label
	Label() string

	// An array of debug signposts that Metal records as the GPU executes the commands of the encoder’s pass.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferEncoderInfo/debugSignposts
	DebugSignposts() []string

	// The execution status of the command encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferEncoderInfo/errorState
	ErrorState() MTLCommandEncoderErrorState
}

// MTLCommandBufferEncoderInfoObject wraps an existing Objective-C object that conforms to the MTLCommandBufferEncoderInfo protocol.
type MTLCommandBufferEncoderInfoObject struct {
	objectivec.Object
}

func (o MTLCommandBufferEncoderInfoObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLCommandBufferEncoderInfoObjectFromID constructs a [MTLCommandBufferEncoderInfoObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLCommandBufferEncoderInfoObjectFromID(id objc.ID) MTLCommandBufferEncoderInfoObject {
	return MTLCommandBufferEncoderInfoObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The name of the encoder that generates the error information.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferEncoderInfo/label
func (o MTLCommandBufferEncoderInfoObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// An array of debug signposts that Metal records as the GPU executes the
// commands of the encoder’s pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferEncoderInfo/debugSignposts
func (o MTLCommandBufferEncoderInfoObject) DebugSignposts() []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("debugSignposts"))
	return objc.ConvertSliceToStrings(rv)
}

// The execution status of the command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferEncoderInfo/errorState
func (o MTLCommandBufferEncoderInfoObject) ErrorState() MTLCommandEncoderErrorState {
	rv := objc.Send[MTLCommandEncoderErrorState](o.ID, objc.Sel("errorState"))
	return rv
}
