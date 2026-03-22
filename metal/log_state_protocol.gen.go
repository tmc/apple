// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A container for shader log messages.
//
// See: https://developer.apple.com/documentation/Metal/MTLLogState
type MTLLogState interface {
	objectivec.IObject

	// Adds a log handler to customize the presentation of shader logging.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLogState/addLogHandler(_:)
	AddLogHandler(block VoidHandler)
}

// MTLLogStateObject wraps an existing Objective-C object that conforms to the MTLLogState protocol.
type MTLLogStateObject struct {
	objectivec.Object
}
func (o MTLLogStateObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLLogStateObjectFromID constructs a [MTLLogStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLLogStateObjectFromID(id objc.ID) MTLLogStateObject {
	return MTLLogStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Adds a log handler to customize the presentation of shader logging.
//
// # Discussion
// 
// In absence of any log handlers, all messages goes through the unified log
// system and are available through the console.app, log tool, or Xcode.
// 
// Use this method to add your own shader logging presentation or filter
// messages using subsystem, category and levels. For more details on how to
// configure your logging, see [Generating Log Messages from Your Code].
//
// [Generating Log Messages from Your Code]: https://developer.apple.com/documentation/os/generating-log-messages-from-your-code#Create-a-Log-Object-to-Organize-Messages
//
// See: https://developer.apple.com/documentation/Metal/MTLLogState/addLogHandler(_:)
func (o MTLLogStateObject) AddLogHandler(block VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("addLogHandler:"), block)
	}

