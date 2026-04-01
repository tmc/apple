// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSXPCInterface] class.
var (
	_NSXPCInterfaceClass     NSXPCInterfaceClass
	_NSXPCInterfaceClassOnce sync.Once
)

func getNSXPCInterfaceClass() NSXPCInterfaceClass {
	_NSXPCInterfaceClassOnce.Do(func() {
		_NSXPCInterfaceClass = NSXPCInterfaceClass{class: objc.GetClass("NSXPCInterface")}
	})
	return _NSXPCInterfaceClass
}

// GetNSXPCInterfaceClass returns the class object for NSXPCInterface.
func GetNSXPCInterfaceClass() NSXPCInterfaceClass {
	return getNSXPCInterfaceClass()
}

type NSXPCInterfaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSXPCInterfaceClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSXPCInterfaceClass) Alloc() NSXPCInterface {
	rv := objc.Send[NSXPCInterface](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An interface that may be sent to an exported object or remote object proxy.
//
// # Overview
//
// This object holds all information about the interface of an exported object
// or remote object proxy. It describes what messages are allowed, what kinds
// of objects are allowed as arguments, what the signature of any reply blocks
// are, and information about additional proxy objects.
//
// # Instance Properties
//
//   - [NSXPCInterface.Protocol]: The Objective-C protocol that this interface is based on.
//   - [NSXPCInterface.SetProtocol]
//
// # Instance Methods
//
//   - [NSXPCInterface.ClassesForSelectorArgumentIndexOfReply]: Returns the current list of allowed classes that can appear within the specified collection object argument to the specified method.
//   - [NSXPCInterface.InterfaceForSelectorArgumentIndexOfReply]: Returns the interface previously set for the specified selector and parameter.
//   - [NSXPCInterface.SetClassesForSelectorArgumentIndexOfReply]: Sets the classes that can appear within the (numerically) specified collection object argument to the specified method.
//   - [NSXPCInterface.SetInterfaceForSelectorArgumentIndexOfReply]: Configures a specific parameter of a method to be sent as a proxy object instead of copied.
//   - [NSXPCInterface.SetXPCTypeForSelectorArgumentIndexOfReply]
//   - [NSXPCInterface.XPCTypeForSelectorArgumentIndexOfReply]
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface
type NSXPCInterface struct {
	objectivec.Object
}

// NSXPCInterfaceFromID constructs a [NSXPCInterface] from an objc.ID.
//
// An interface that may be sent to an exported object or remote object proxy.
func NSXPCInterfaceFromID(id objc.ID) NSXPCInterface {
	return NSXPCInterface{objectivec.Object{ID: id}}
}

// NOTE: NSXPCInterface adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSXPCInterface] class.
//
// # Instance Properties
//
//   - [INSXPCInterface.Protocol]: The Objective-C protocol that this interface is based on.
//   - [INSXPCInterface.SetProtocol]
//
// # Instance Methods
//
//   - [INSXPCInterface.ClassesForSelectorArgumentIndexOfReply]: Returns the current list of allowed classes that can appear within the specified collection object argument to the specified method.
//   - [INSXPCInterface.InterfaceForSelectorArgumentIndexOfReply]: Returns the interface previously set for the specified selector and parameter.
//   - [INSXPCInterface.SetClassesForSelectorArgumentIndexOfReply]: Sets the classes that can appear within the (numerically) specified collection object argument to the specified method.
//   - [INSXPCInterface.SetInterfaceForSelectorArgumentIndexOfReply]: Configures a specific parameter of a method to be sent as a proxy object instead of copied.
//   - [INSXPCInterface.SetXPCTypeForSelectorArgumentIndexOfReply]
//   - [INSXPCInterface.XPCTypeForSelectorArgumentIndexOfReply]
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface
type INSXPCInterface interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The Objective-C protocol that this interface is based on.
	Protocol() *objectivec.Protocol
	SetProtocol(value *objectivec.Protocol)

	// Topic: Instance Methods

	// Returns the current list of allowed classes that can appear within the specified collection object argument to the specified method.
	ClassesForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) INSSet
	// Returns the interface previously set for the specified selector and parameter.
	InterfaceForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) INSXPCInterface
	// Sets the classes that can appear within the (numerically) specified collection object argument to the specified method.
	SetClassesForSelectorArgumentIndexOfReply(classes INSSet, sel objc.SEL, arg uint, ofReply bool)
	// Configures a specific parameter of a method to be sent as a proxy object instead of copied.
	SetInterfaceForSelectorArgumentIndexOfReply(ifc INSXPCInterface, sel objc.SEL, arg uint, ofReply bool)
	SetXPCTypeForSelectorArgumentIndexOfReply(type_ unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool)
	XPCTypeForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) unsafe.Pointer
}

// Init initializes the instance.
func (x NSXPCInterface) Init() NSXPCInterface {
	rv := objc.Send[NSXPCInterface](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x NSXPCInterface) Autorelease() NSXPCInterface {
	rv := objc.Send[NSXPCInterface](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSXPCInterface creates a new NSXPCInterface instance.
func NewNSXPCInterface() NSXPCInterface {
	class := getNSXPCInterfaceClass()
	rv := objc.Send[NSXPCInterface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an NSXPCInterface instance for a given protocol.
//
// # Discussion
//
// Most interfaces do not need any further configuration. Interfaces with
// collection classes or additional proxy objects should be configured using
// the other methods in this class.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface/init(with:)
func NewXPCInterfaceWithProtocol(protocol_ *objectivec.Protocol) NSXPCInterface {
	rv := objc.Send[objc.ID](objc.ID(getNSXPCInterfaceClass().class), objc.Sel("interfaceWithProtocol:"), protocol_)
	return NSXPCInterfaceFromID(rv)
}

// Returns the current list of allowed classes that can appear within the
// specified collection object argument to the specified method.
//
// sel: Specifies which method in the protocol you want information about.
//
// arg: Specifies the position (starting at index 0) of the parameter for which you
// want to obtain the current set of allowed classes. This may be either the
// position of a parameter in the method itself or the position in its reply
// block.
//
// ofReply: Pass true if `arg` is an index into the parameters of the reply block, or
// false if it is an index into the parameters of the method itself.
//
// # Discussion
//
// See [SetClassesForSelectorArgumentIndexOfReply] for more explanation.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface/classes(for:argumentIndex:ofReply:)
func (x NSXPCInterface) ClassesForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) INSSet {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("classesForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return NSSetFromID(rv)
}

// Returns the interface previously set for the specified selector and
// parameter.
//
// sel: Specifies which method in the protocol you want information about.
//
// arg: Specifies the position (starting at index 0) of the parameter for which you
// want to obtain the current interface. This may be either the position of a
// parameter in the method itself or the position in its reply block.
//
// ofReply: Pass true if `arg` is an index into the parameters of the reply block, or
// false if it is an index into the parameters of the method itself.
//
// # Discussion
//
// See [SetInterfaceForSelectorArgumentIndexOfReply] for more explanation.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface/forSelector(_:argumentIndex:ofReply:)
func (x NSXPCInterface) InterfaceForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) INSXPCInterface {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("interfaceForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return NSXPCInterfaceFromID(rv)
}

// Sets the classes that can appear within the (numerically) specified
// collection object argument to the specified method.
//
// classes: An [NSSet] containing Class objects—for example, `[MyObject class]`.
//
// sel: Specifies which method in the protocol is being configured.
//
// arg: Specifies the position (starting at index 0) of the parameter for which you
// are allowing classes. This may be either the position of a parameter in the
// method itself or the position in its reply block.
//
// ofReply: Pass true if `arg` is an index into the parameters of the reply block, or
// false if it is an index into the parameters of the method itself.
//
// # Discussion
//
// If an argument to a method in your protocol is a collection class (for
// example, NSArray or NSDictionary), then you must explicitly specify the set
// of expected classes that may appear within that collection.
//
// If the expected classes are all property list types, calling this method is
// optional; property list types are allowed by default inside collection
// objects. You may, however, call this method to further restrict the set of
// allowed classes.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface/setClasses(_:for:argumentIndex:ofReply:)
func (x NSXPCInterface) SetClassesForSelectorArgumentIndexOfReply(classes INSSet, sel objc.SEL, arg uint, ofReply bool) {
	objc.Send[objc.ID](x.ID, objc.Sel("setClasses:forSelector:argumentIndex:ofReply:"), classes, sel, arg, ofReply)
}

// Configures a specific parameter of a method to be sent as a proxy object
// instead of copied.
//
// ifc: The [NSXPCInterface] object that describes the protocol for the proxy
// object. The interface is configured the same way as the interface for an
// exported object or remote object proxy.
//
// sel: Specifies which method in the protocol is being configured.
//
// arg: Specifies the position (starting at index 0) of the parameter for which you
// are configuring a proxy object. This may be either the position of a
// parameter in the method itself or the position in its reply block. This
// argument must be an object.
//
// ofReply: Pass true if `arg` is an index into the parameters of the reply block, or
// false if it is an index into the parameters of the method itself.
//
// # Discussion
//
// If an argument to a method in your protocol should be sent as a proxy
// object instead of by copy, then configure the interface for that protocol
// with a new interface for a specific argument. An example of an object that
// should be a proxy instead of being copied is a view object.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface/setInterface(_:for:argumentIndex:ofReply:)
func (x NSXPCInterface) SetInterfaceForSelectorArgumentIndexOfReply(ifc INSXPCInterface, sel objc.SEL, arg uint, ofReply bool) {
	objc.Send[objc.ID](x.ID, objc.Sel("setInterface:forSelector:argumentIndex:ofReply:"), ifc, sel, arg, ofReply)
}

// type is a [xpc.xpc_type_t].
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface/setXPCType(_:for:argumentIndex:ofReply:)
func (x NSXPCInterface) SetXPCTypeForSelectorArgumentIndexOfReply(type_ unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool) {
	objc.Send[objc.ID](x.ID, objc.Sel("setXPCType:forSelector:argumentIndex:ofReply:"), type_, sel, arg, ofReply)
}

// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface/xpcType(for:argumentIndex:ofReply:)
func (x NSXPCInterface) XPCTypeForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x.ID, objc.Sel("XPCTypeForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return rv
}

// The Objective-C protocol that this interface is based on.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCInterface/protocol
func (x NSXPCInterface) Protocol() *objectivec.Protocol {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("protocol"))
	if rv == 0 {
		return nil
	}
	val := objectivec.ProtocolFromID(objc.ID(rv))
	return &val
}
func (x NSXPCInterface) SetProtocol(value *objectivec.Protocol) {
	if value == nil {
		objc.Send[struct{}](x.ID, objc.Sel("setProtocol:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](x.ID, objc.Sel("setProtocol:"), value)
}
