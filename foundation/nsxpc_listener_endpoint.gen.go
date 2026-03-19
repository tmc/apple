// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSXPCListenerEndpoint] class.
var (
	_NSXPCListenerEndpointClass     NSXPCListenerEndpointClass
	_NSXPCListenerEndpointClassOnce sync.Once
)

func getNSXPCListenerEndpointClass() NSXPCListenerEndpointClass {
	_NSXPCListenerEndpointClassOnce.Do(func() {
		_NSXPCListenerEndpointClass = NSXPCListenerEndpointClass{class: objc.GetClass("NSXPCListenerEndpoint")}
	})
	return _NSXPCListenerEndpointClass
}

// GetNSXPCListenerEndpointClass returns the class object for NSXPCListenerEndpoint.
func GetNSXPCListenerEndpointClass() NSXPCListenerEndpointClass {
	return getNSXPCListenerEndpointClass()
}

type NSXPCListenerEndpointClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSXPCListenerEndpointClass) Alloc() NSXPCListenerEndpoint {
	rv := objc.Send[NSXPCListenerEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that names a specific XPC listener.
//
// # Overview
// 
// An instance of [NSXPCListenerEndpoint] may be retrieved from an
// [NSXPCListener] instance and sent over existing [NSXPCConnection]s. A
// process may then use the endpoint to create a new [NSXPCConnection] to the
// original [NSXPCListener].
// 
// This pattern is useful if you have a service which multiplexes work to
// other services. The service can act as an intermediate helper. The
// requesting application does not need to know specifically which service it
// is connecting to, just that it implements a known [NSXPCInterface].
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListenerEndpoint
type NSXPCListenerEndpoint struct {
	objectivec.Object
}

// NSXPCListenerEndpointFromID constructs a [NSXPCListenerEndpoint] from an objc.ID.
//
// An object that names a specific XPC listener.
func NSXPCListenerEndpointFromID(id objc.ID) NSXPCListenerEndpoint {
	return NSXPCListenerEndpoint{objectivec.Object{ID: id}}
}
// NOTE: NSXPCListenerEndpoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSXPCListenerEndpoint] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListenerEndpoint
type INSXPCListenerEndpoint interface {
	objectivec.IObject
	NSCoding
	NSSecureCoding
}

// Init initializes the instance.
func (x NSXPCListenerEndpoint) Init() NSXPCListenerEndpoint {
	rv := objc.Send[NSXPCListenerEndpoint](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x NSXPCListenerEndpoint) Autorelease() NSXPCListenerEndpoint {
	rv := objc.Send[NSXPCListenerEndpoint](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSXPCListenerEndpoint creates a new NSXPCListenerEndpoint instance.
func NewNSXPCListenerEndpoint() NSXPCListenerEndpoint {
	class := getNSXPCListenerEndpointClass()
	rv := objc.Send[NSXPCListenerEndpoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewXPCListenerEndpointWithCoder(coder INSCoder) NSXPCListenerEndpoint {
	instance := getNSXPCListenerEndpointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSXPCListenerEndpointFromID(rv)
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (x NSXPCListenerEndpoint) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](x.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (x NSXPCListenerEndpoint) InitWithCoder(coder INSCoder) NSXPCListenerEndpoint {
	rv := objc.Send[NSXPCListenerEndpoint](x.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

			// Protocol methods for NSSecureCoding
			

