// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSMachPort] class.
var (
	_NSMachPortClass     NSMachPortClass
	_NSMachPortClassOnce sync.Once
)

func getNSMachPortClass() NSMachPortClass {
	_NSMachPortClassOnce.Do(func() {
		_NSMachPortClass = NSMachPortClass{class: objc.GetClass("NSMachPort")}
	})
	return _NSMachPortClass
}

// GetNSMachPortClass returns the class object for NSMachPort.
func GetNSMachPortClass() NSMachPortClass {
	return getNSMachPortClass()
}

type NSMachPortClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMachPortClass) Alloc() NSMachPort {
	rv := objc.Send[NSMachPort](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A port that can be used as an endpoint for distributed object connections
// (or raw messaging).
//
// # Overview
// 
// [NSMachPort] is a subclass of [NSPort] that wraps a Mach port, the
// fundamental communication port in macOS. [NSMachPort] allows for local (on
// the same machine) communication only. A companion class, [NSSocketPort],
// allows for both local and remote distributed object communication, but may
// be more expensive than [NSMachPort] for the local case.
// 
// To use [NSMachPort] effectively, you should be familiar with Mach ports,
// port access rights, and Mach messages. See the Mach OS documentation for
// more information.
//
// # Creating and Initializing
//
//   - [NSMachPort.InitWithMachPort]: Initializes a newly allocated [NSMachPort] object with a given Mach port.
//   - [NSMachPort.InitWithMachPortOptions]: Initializes a newly allocated [NSMachPort] object with a given Mach port and the specified options.
//
// # Getting the Mach Port
//
//   - [NSMachPort.MachPort]: The Mach port used by the receiver, represented as an integer.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPort
type NSMachPort struct {
	NSPort
}

// NSMachPortFromID constructs a [NSMachPort] from an objc.ID.
//
// A port that can be used as an endpoint for distributed object connections
// (or raw messaging).
func NSMachPortFromID(id objc.ID) NSMachPort {
	return NSMachPort{NSPort: NSPortFromID(id)}
}
// NOTE: NSMachPort adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSMachPort] class.
//
// # Creating and Initializing
//
//   - [INSMachPort.InitWithMachPort]: Initializes a newly allocated [NSMachPort] object with a given Mach port.
//   - [INSMachPort.InitWithMachPortOptions]: Initializes a newly allocated [NSMachPort] object with a given Mach port and the specified options.
//
// # Getting the Mach Port
//
//   - [INSMachPort.MachPort]: The Mach port used by the receiver, represented as an integer.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPort
type INSMachPort interface {
	INSPort

	// Topic: Creating and Initializing

	// Initializes a newly allocated [NSMachPort] object with a given Mach port.
	InitWithMachPort(machPort uint32) NSMachPort
	// Initializes a newly allocated [NSMachPort] object with a given Mach port and the specified options.
	InitWithMachPortOptions(machPort uint32, f NSMachPortOptions) NSMachPort

	// Topic: Getting the Mach Port

	// The Mach port used by the receiver, represented as an integer.
	MachPort() uint32
}





// Init initializes the instance.
func (m NSMachPort) Init() NSMachPort {
	rv := objc.Send[NSMachPort](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMachPort) Autorelease() NSMachPort {
	rv := objc.Send[NSMachPort](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMachPort creates a new NSMachPort instance.
func NewNSMachPort() NSMachPort {
	class := getNSMachPortClass()
	rv := objc.Send[NSMachPort](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMachPortWithCoder(coder INSCoder) NSMachPort {
	instance := getNSMachPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMachPortFromID(rv)
}


// Initializes a newly allocated [NSMachPort] object with a given Mach port.
//
// machPort: The Mach port for the new port. This parameter should originally be of type
// mach_port_t.
//
// # Return Value
// 
// Returns an initialized [NSMachPort] object that uses `machPort` to send or
// receive messages. The returned object might be different than the original
// receiver
//
// # Discussion
// 
// Depending on the access rights for `machPort`, the new port may be able to
// only send messages. If a port with `machPort` already exists, this method
// deallocates the receiver, then retains and returns the existing port.
// 
// This method is the designated initializer for the [NSMachPort] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPort/init(machPort:)
func NewMachPortWithMachPort(machPort uint32) NSMachPort {
	instance := getNSMachPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMachPort:"), machPort)
	return NSMachPortFromID(rv)
}


// Initializes a newly allocated [NSMachPort] object with a given Mach port
// and the specified options.
//
// machPort: The Mach port for the new port. This parameter should originally be of type
// mach_port_t.
//
// f: Specifies options for what to do with the underlying port rights when the
// [NSMachPort] object is invalidated or destroyed. For a list of constants,
// see `Mach Port Rights`.
//
// # Return Value
// 
// Returns an initialized [NSMachPort] object that uses `machPort` to send or
// receive messages. The returned object might be different than the original
// receiver
//
// # Discussion
// 
// Depending on the access rights for `machPort`, the new port may be able to
// only send messages. If a port with `machPort` already exists, this method
// deallocates the receiver, then retains and returns the existing port.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPort/init(machPort:options:)
func NewMachPortWithMachPortOptions(machPort uint32, f NSMachPortOptions) NSMachPort {
	instance := getNSMachPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMachPort:options:"), machPort, f)
	return NSMachPortFromID(rv)
}







// Initializes a newly allocated [NSMachPort] object with a given Mach port.
//
// machPort: The Mach port for the new port. This parameter should originally be of type
// mach_port_t.
//
// # Return Value
// 
// Returns an initialized [NSMachPort] object that uses `machPort` to send or
// receive messages. The returned object might be different than the original
// receiver
//
// # Discussion
// 
// Depending on the access rights for `machPort`, the new port may be able to
// only send messages. If a port with `machPort` already exists, this method
// deallocates the receiver, then retains and returns the existing port.
// 
// This method is the designated initializer for the [NSMachPort] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPort/init(machPort:)
func (m NSMachPort) InitWithMachPort(machPort uint32) NSMachPort {
	rv := objc.Send[NSMachPort](m.ID, objc.Sel("initWithMachPort:"), machPort)
	return rv
}

// Initializes a newly allocated [NSMachPort] object with a given Mach port
// and the specified options.
//
// machPort: The Mach port for the new port. This parameter should originally be of type
// mach_port_t.
//
// f: Specifies options for what to do with the underlying port rights when the
// [NSMachPort] object is invalidated or destroyed. For a list of constants,
// see `Mach Port Rights`.
//
// # Return Value
// 
// Returns an initialized [NSMachPort] object that uses `machPort` to send or
// receive messages. The returned object might be different than the original
// receiver
//
// # Discussion
// 
// Depending on the access rights for `machPort`, the new port may be able to
// only send messages. If a port with `machPort` already exists, this method
// deallocates the receiver, then retains and returns the existing port.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPort/init(machPort:options:)
func (m NSMachPort) InitWithMachPortOptions(machPort uint32, f NSMachPortOptions) NSMachPort {
	rv := objc.Send[NSMachPort](m.ID, objc.Sel("initWithMachPort:options:"), machPort, f)
	return rv
}





// Creates and returns a port object configured with the given Mach port.
//
// machPort: The Mach port for the new port. This parameter should originally be of type
// mach_port_t.
//
// # Return Value
// 
// An [NSMachPort] object that uses `machPort` to send or receive messages.
//
// # Discussion
// 
// Creates the port object if necessary. Depending on the access rights
// associated with `machPort`, the new port object may be usable only for
// sending messages.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPort/port(withMachPort:)
func (_NSMachPortClass NSMachPortClass) PortWithMachPort(machPort uint32) Port {
	rv := objc.Send[objc.ID](objc.ID(_NSMachPortClass.class), objc.Sel("portWithMachPort:"), machPort)
	return NSPortFromID(rv)
}

// Creates and returns a port object configured with the specified options and
// the given Mach port.
//
// machPort: The Mach port for the new port. This parameter should originally be of type
// mach_port_t.
//
// f: Specifies options for what to do with the underlying port rights when the
// [NSMachPort] object is invalidated or destroyed. For a list of constants,
// see `Mach Port Rights`.
//
// # Return Value
// 
// An [NSMachPort] object that uses `machPort` to send or receive messages.
//
// # Discussion
// 
// Creates the port object if necessary. Depending on the access rights
// associated with `machPort`, the new port object may be usable only for
// sending messages.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPort/port(withMachPort:options:)
func (_NSMachPortClass NSMachPortClass) PortWithMachPortOptions(machPort uint32, f NSMachPortOptions) Port {
	rv := objc.Send[objc.ID](objc.ID(_NSMachPortClass.class), objc.Sel("portWithMachPort:options:"), machPort, f)
	return NSPortFromID(rv)
}








// The Mach port used by the receiver, represented as an integer.
//
// # Discussion
// 
// The Mach port used by the receiver. Cast this value to a mach_port_t when
// using it with Mach system calls.
//
// See: https://developer.apple.com/documentation/Foundation/NSMachPort/machPort
func (m NSMachPort) MachPort() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("machPort"))
	return rv
}



























