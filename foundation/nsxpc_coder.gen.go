// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSXPCCoder] class.
var (
	_NSXPCCoderClass     NSXPCCoderClass
	_NSXPCCoderClassOnce sync.Once
)

func getNSXPCCoderClass() NSXPCCoderClass {
	_NSXPCCoderClassOnce.Do(func() {
		_NSXPCCoderClass = NSXPCCoderClass{class: objc.GetClass("NSXPCCoder")}
	})
	return _NSXPCCoderClass
}

// GetNSXPCCoderClass returns the class object for NSXPCCoder.
func GetNSXPCCoderClass() NSXPCCoderClass {
	return getNSXPCCoderClass()
}

type NSXPCCoderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSXPCCoderClass) Alloc() NSXPCCoder {
	rv := objc.Send[NSXPCCoder](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A coder that encodes and decodes objects that your app sends over an XPC
// connection.
//
// # Overview
// 
// If you want to perform custom encoding or decoding of [Codable] objects
// that your app sends over an [NSXPCConnection], use [isKind(of:)] to
// determine if the coder provided to your object is a kind of [NSXPCCoder].
//
// [Codable]: https://developer.apple.com/documentation/Swift/Codable
// [isKind(of:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isKind(of:)
//
// # Inspecting the Coder
//
//   - [NSXPCCoder.Connection]: The connection currently performing encoding or decoding.
//   - [NSXPCCoder.UserInfo]: An optional user information object associated with the coder.
//   - [NSXPCCoder.SetUserInfo]
//
// # Encoding and Decoding
//
//   - [NSXPCCoder.EncodeXPCObjectForKey]: Encodes an object to send over an XPC connection.
//   - [NSXPCCoder.DecodeXPCObjectOfTypeForKey]: Decodes an object and validates that its type matches the type a service provides over XPC.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCCoder
type NSXPCCoder struct {
	NSCoder
}

// NSXPCCoderFromID constructs a [NSXPCCoder] from an objc.ID.
//
// A coder that encodes and decodes objects that your app sends over an XPC
// connection.
func NSXPCCoderFromID(id objc.ID) NSXPCCoder {
	return NSXPCCoder{NSCoder: NSCoderFromID(id)}
}
// NOTE: NSXPCCoder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSXPCCoder] class.
//
// # Inspecting the Coder
//
//   - [INSXPCCoder.Connection]: The connection currently performing encoding or decoding.
//   - [INSXPCCoder.UserInfo]: An optional user information object associated with the coder.
//   - [INSXPCCoder.SetUserInfo]
//
// # Encoding and Decoding
//
//   - [INSXPCCoder.EncodeXPCObjectForKey]: Encodes an object to send over an XPC connection.
//   - [INSXPCCoder.DecodeXPCObjectOfTypeForKey]: Decodes an object and validates that its type matches the type a service provides over XPC.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCCoder
type INSXPCCoder interface {
	INSCoder

	// Topic: Inspecting the Coder

	// The connection currently performing encoding or decoding.
	Connection() INSXPCConnection
	// An optional user information object associated with the coder.
	UserInfo() objectivec.Object
	SetUserInfo(value objectivec.Object)

	// Topic: Encoding and Decoding

	// Encodes an object to send over an XPC connection.
	EncodeXPCObjectForKey(xpcObject objectivec.Object, key string)
	// Decodes an object and validates that its type matches the type a service provides over XPC.
	DecodeXPCObjectOfTypeForKey(type_ unsafe.Pointer, key string) objectivec.Object
}





// Init initializes the instance.
func (x NSXPCCoder) Init() NSXPCCoder {
	rv := objc.Send[NSXPCCoder](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x NSXPCCoder) Autorelease() NSXPCCoder {
	rv := objc.Send[NSXPCCoder](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSXPCCoder creates a new NSXPCCoder instance.
func NewNSXPCCoder() NSXPCCoder {
	class := getNSXPCCoderClass()
	rv := objc.Send[NSXPCCoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Encodes an object to send over an XPC connection.
//
// xpcObject: An object that XPC can encode.
//
// key: A string that your app uses to reference the encoded object.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCCoder/encodeXPCObject(_:forKey:)
func (x NSXPCCoder) EncodeXPCObjectForKey(xpcObject objectivec.Object, key string) {
	objc.Send[objc.ID](x.ID, objc.Sel("encodeXPCObject:forKey:"), xpcObject, objc.String(key))
}

// Decodes an object and validates that its type matches the type a service
// provides over XPC.
//
// type: An opaque pointer to an encoded XPC object.
//
// key: A string that your app uses to reference the decoded object.
//
// # Return Value
// 
// An object that XPC can encode.
//
// # Discussion
// 
// The [DecodeXPCObjectOfTypeForKey] method validates that the type of the
// decoded object matches the type of the encoded object. If they don’t
// match, the [NSXPCCoder] throws an exception in support of [NSSecureCoding].
// 
// Be sure to check the result against [Null] if you call an [XPC] function
// because calling an [XPC] function on a [Null] object results in a crash.
//
// [XPC]: https://developer.apple.com/documentation/XPC
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCCoder/decodeXPCObject(ofType:forKey:)
func (x NSXPCCoder) DecodeXPCObjectOfTypeForKey(type_ unsafe.Pointer, key string) objectivec.Object {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("decodeXPCObjectOfType:forKey:"), type_, objc.String(key))
	return objectivec.ObjectFromID(rv)
}












// The connection currently performing encoding or decoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCCoder/connection
func (x NSXPCCoder) Connection() INSXPCConnection {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("connection"))
	return NSXPCConnectionFromID(objc.ID(rv))
}



// An optional user information object associated with the coder.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCCoder/userInfo
func (x NSXPCCoder) UserInfo() objectivec.Object {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("userInfo"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (x NSXPCCoder) SetUserInfo(value objectivec.Object) {
	objc.Send[struct{}](x.ID, objc.Sel("setUserInfo:"), value)
}

























