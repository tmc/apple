// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSURLSessionWebSocketMessage] class.
var (
	_NSURLSessionWebSocketMessageClass     NSURLSessionWebSocketMessageClass
	_NSURLSessionWebSocketMessageClassOnce sync.Once
)

func getNSURLSessionWebSocketMessageClass() NSURLSessionWebSocketMessageClass {
	_NSURLSessionWebSocketMessageClassOnce.Do(func() {
		_NSURLSessionWebSocketMessageClass = NSURLSessionWebSocketMessageClass{class: objc.GetClass("NSURLSessionWebSocketMessage")}
	})
	return _NSURLSessionWebSocketMessageClass
}

// GetNSURLSessionWebSocketMessageClass returns the class object for NSURLSessionWebSocketMessage.
func GetNSURLSessionWebSocketMessageClass() NSURLSessionWebSocketMessageClass {
	return getNSURLSessionWebSocketMessageClass()
}

type NSURLSessionWebSocketMessageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSURLSessionWebSocketMessageClass) Alloc() NSURLSessionWebSocketMessage {
	rv := objc.Send[NSURLSessionWebSocketMessage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







//
// # Instance Properties
//
//   - [NSURLSessionWebSocketMessage.Data]
//   - [NSURLSessionWebSocketMessage.String]
//   - [NSURLSessionWebSocketMessage.Type]
//
// # Instance Methods
//
//   - [NSURLSessionWebSocketMessage.InitWithData]
//   - [NSURLSessionWebSocketMessage.InitWithString]
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage
type NSURLSessionWebSocketMessage struct {
	objectivec.Object
}

// NSURLSessionWebSocketMessageFromID constructs a [NSURLSessionWebSocketMessage] from an objc.ID.
func NSURLSessionWebSocketMessageFromID(id objc.ID) NSURLSessionWebSocketMessage {
	return NSURLSessionWebSocketMessage{objectivec.Object{id}}
}
// Ensure NSURLSessionWebSocketMessage implements INSURLSessionWebSocketMessage.
var _ INSURLSessionWebSocketMessage = NSURLSessionWebSocketMessage{}





// An interface definition for the [NSURLSessionWebSocketMessage] class.
//
// # Instance Properties
//
//   - [INSURLSessionWebSocketMessage.Data]
//   - [INSURLSessionWebSocketMessage.String]
//   - [INSURLSessionWebSocketMessage.Type]
//
// # Instance Methods
//
//   - [INSURLSessionWebSocketMessage.InitWithData]
//   - [INSURLSessionWebSocketMessage.InitWithString]
//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage
type INSURLSessionWebSocketMessage interface {
	objectivec.IObject

	// Topic: Instance Properties

	Data() INSData
	String() string
	Type() NSURLSessionWebSocketMessageType

	// Topic: Instance Methods

	InitWithData(data INSData) NSURLSessionWebSocketMessage
	InitWithString(string_ string) NSURLSessionWebSocketMessage
}





// Init initializes the instance.
func (u NSURLSessionWebSocketMessage) Init() NSURLSessionWebSocketMessage {
	rv := objc.Send[NSURLSessionWebSocketMessage](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSURLSessionWebSocketMessage) Autorelease() NSURLSessionWebSocketMessage {
	rv := objc.Send[NSURLSessionWebSocketMessage](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSURLSessionWebSocketMessage creates a new NSURLSessionWebSocketMessage instance.
func NewNSURLSessionWebSocketMessage() NSURLSessionWebSocketMessage {
	class := getNSURLSessionWebSocketMessageClass()
	rv := objc.Send[NSURLSessionWebSocketMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/initWithData:
func NewURLSessionWebSocketMessageWithData(data INSData) NSURLSessionWebSocketMessage {
	instance := getNSURLSessionWebSocketMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return NSURLSessionWebSocketMessageFromID(rv)
}


//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/initWithString:
func NewURLSessionWebSocketMessageWithString(string_ string) NSURLSessionWebSocketMessage {
	instance := getNSURLSessionWebSocketMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(string_))
	return NSURLSessionWebSocketMessageFromID(rv)
}







//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/initWithData:
func (u NSURLSessionWebSocketMessage) InitWithData(data INSData) NSURLSessionWebSocketMessage {
	rv := objc.Send[NSURLSessionWebSocketMessage](u.ID, objc.Sel("initWithData:"), data)
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/initWithString:
func (u NSURLSessionWebSocketMessage) InitWithString(string_ string) NSURLSessionWebSocketMessage {
	rv := objc.Send[NSURLSessionWebSocketMessage](u.ID, objc.Sel("initWithString:"), objc.String(string_))
	return rv
}












// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/data
func (u NSURLSessionWebSocketMessage) Data() INSData {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("data"))
	return NSDataFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/string
func (u NSURLSessionWebSocketMessage) String() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("string"))
	return NSStringFromID(rv).String()
}



// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessage/type
func (u NSURLSessionWebSocketMessage) Type() NSURLSessionWebSocketMessageType {
	rv := objc.Send[NSURLSessionWebSocketMessageType](u.ID, objc.Sel("type"))
	return NSURLSessionWebSocketMessageType(rv)
}

















