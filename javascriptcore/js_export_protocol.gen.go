// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol for exporting Objective-C objects to JavaScript.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSExport
type JSExport interface {
	objectivec.IObject
}

// JSExportObject wraps an existing Objective-C object that conforms to the JSExport protocol.
type JSExportObject struct {
	objectivec.Object
}

func (o JSExportObject) BaseObject() objectivec.Object {
	return o.Object
}

// JSExportObjectFromID constructs a [JSExportObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func JSExportObjectFromID(id objc.ID) JSExportObject {
	return JSExportObject{
		Object: objectivec.ObjectFromID(id),
	}
}
