// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The source code that logged a debug message.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogDebugLocation
type MTLFunctionLogDebugLocation interface {
	objectivec.IObject

	// The name of the shader function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogDebugLocation/functionName
	FunctionName() string

	// The URL of the file that contains the shader function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogDebugLocation/url
	URL() foundation.INSURL

	// The line that the log message appears on.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogDebugLocation/line
	Line() uint

	// The column where the log message appears.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogDebugLocation/column
	Column() uint
}



// MTLFunctionLogDebugLocationObject wraps an existing Objective-C object that conforms to the MTLFunctionLogDebugLocation protocol.
type MTLFunctionLogDebugLocationObject struct {
	objectivec.Object
}
func (o MTLFunctionLogDebugLocationObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTLFunctionLogDebugLocationObjectFromID constructs a [MTLFunctionLogDebugLocationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLFunctionLogDebugLocationObjectFromID(id objc.ID) MTLFunctionLogDebugLocationObject {
	return MTLFunctionLogDebugLocationObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// The name of the shader function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogDebugLocation/functionName

func (o MTLFunctionLogDebugLocationObject) FunctionName() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
	}

// The URL of the file that contains the shader function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogDebugLocation/url

func (o MTLFunctionLogDebugLocationObject) URL() foundation.INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(rv)
	}

// The line that the log message appears on.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogDebugLocation/line

func (o MTLFunctionLogDebugLocationObject) Line() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("line"))
	return rv
	}

// The column where the log message appears.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogDebugLocation/column

func (o MTLFunctionLogDebugLocationObject) Column() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("column"))
	return rv
	}















