// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAutoreleasePool] class.
var (
	_NSAutoreleasePoolClass     NSAutoreleasePoolClass
	_NSAutoreleasePoolClassOnce sync.Once
)

func getNSAutoreleasePoolClass() NSAutoreleasePoolClass {
	_NSAutoreleasePoolClassOnce.Do(func() {
		_NSAutoreleasePoolClass = NSAutoreleasePoolClass{class: objc.GetClass("NSAutoreleasePool")}
	})
	return _NSAutoreleasePoolClass
}

// GetNSAutoreleasePoolClass returns the class object for NSAutoreleasePool.
func GetNSAutoreleasePoolClass() NSAutoreleasePoolClass {
	return getNSAutoreleasePoolClass()
}

type NSAutoreleasePoolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAutoreleasePoolClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAutoreleasePoolClass) Alloc() NSAutoreleasePool {
	rv := objc.Send[NSAutoreleasePool](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that supports Cocoa’s reference-counted memory management
// system.
//
// # Overview
// 
// An autorelease pool stores objects that are sent a `release` message when
// the pool itself is drained.
// 
// In a reference-counted environment (as opposed to one which uses garbage
// collection), an [NSAutoreleasePool] object contains objects that have
// received an [autorelease] message and when drained it sends a [release]
// message to each of those objects. Thus, sending [autorelease] instead of
// [release] to an object extends the lifetime of that object at least until
// the pool itself is drained (it may be longer if the object is subsequently
// retained). An object can be put into the same pool several times, in which
// case it receives a [release] message for each time it was put into the
// pool.
// 
// In a reference counted environment, Cocoa expects there to be an
// autorelease pool always available. If a pool is not available, autoreleased
// objects do not get released and you leak memory. In this situation, your
// program will typically log suitable warning messages.
// 
// The Application Kit creates an autorelease pool on the main thread at the
// beginning of every cycle of the event loop, and drains it at the end,
// thereby releasing any autoreleased objects generated while processing an
// event. If you use the Application Kit, you therefore typically don’t have
// to create your own pools. If your application creates a lot of temporary
// autoreleased objects within the event loop, however, it may be beneficial
// to create “local” autorelease pools to help to minimize the peak memory
// footprint.
// 
// You create an [NSAutoreleasePool] object with the usual `alloc` and `init`
// messages and dispose of it with [NSAutoreleasePool.Drain] (or `release`—to understand the
// difference, see [NSAutoreleasePool]). Since you cannot retain an
// autorelease pool (or autorelease it—see `retain` and `autorelease`),
// draining a pool ultimately has the effect of deallocating it. You should
// always drain an autorelease pool in the same context (invocation of a
// method or function, or body of a loop) that it was created. See [Using
// Autorelease Pool Blocks] for more details.
// 
// Each thread (including the main thread) maintains its own stack of
// [NSAutoreleasePool] objects (see [NSAutoreleasePool]). As new pools are
// created, they get added to the top of the stack. When pools are
// deallocated, they are removed from the stack. Autoreleased objects are
// placed into the top autorelease pool for the current thread. When a thread
// terminates, it automatically drains all of the autorelease pools associated
// with itself.
// 
// # Threads
// 
// If you are making Cocoa calls outside of the Application Kit’s main
// thread—for example if you create a Foundation-only application or if you
// detach a thread—you need to create your own autorelease pool.
// 
// If your application or thread is long-lived and potentially generates a lot
// of autoreleased objects, you should periodically drain and create
// autorelease pools (like the Application Kit does on the main thread);
// otherwise, autoreleased objects accumulate and your memory footprint grows.
// If, however, your detached thread does not make Cocoa calls, you do not
// need to create an autorelease pool.
// 
// # Garbage Collection
// 
// In a garbage-collected environment, there is no need for autorelease pools.
// You may, however, write a framework that is designed to work in both a
// garbage-collected and reference-counted environment. In this case, you can
// use autorelease pools to hint to the collector that collection may be
// appropriate. In a garbage-collected environment, sending a [NSAutoreleasePool.Drain] message
// to a pool triggers garbage collection if necessary; `release`, however, is
// a no-op. In a reference-counted environment, [NSAutoreleasePool.Drain] has the same effect as
// `release`. Typically, therefore, you should use [NSAutoreleasePool.Drain] instead of
// `release`.
//
// [Using Autorelease Pool Blocks]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/Articles/mmAutoreleasePools.html#//apple_ref/doc/uid/20000047
// [autorelease]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/autorelease
// [release]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/release
//
// # Managing a Pool
//
//   - [NSAutoreleasePool.Drain]: In a reference-counted environment, releases and pops the receiver; in a garbage-collected environment, triggers garbage collection if the memory allocated since the last collection is greater than the current threshold.
//
// See: https://developer.apple.com/documentation/Foundation/NSAutoreleasePool
type NSAutoreleasePool struct {
	objectivec.Object
}

// NSAutoreleasePoolFromID constructs a [NSAutoreleasePool] from an objc.ID.
//
// An object that supports Cocoa’s reference-counted memory management
// system.
func NSAutoreleasePoolFromID(id objc.ID) NSAutoreleasePool {
	return NSAutoreleasePool{objectivec.Object{ID: id}}
}
// Ensure NSAutoreleasePool implements INSAutoreleasePool.
var _ INSAutoreleasePool = NSAutoreleasePool{}

// An interface definition for the [NSAutoreleasePool] class.
//
// # Managing a Pool
//
//   - [INSAutoreleasePool.Drain]: In a reference-counted environment, releases and pops the receiver; in a garbage-collected environment, triggers garbage collection if the memory allocated since the last collection is greater than the current threshold.
//
// See: https://developer.apple.com/documentation/Foundation/NSAutoreleasePool
type INSAutoreleasePool interface {
	objectivec.IObject

	// Topic: Managing a Pool

	// In a reference-counted environment, releases and pops the receiver; in a garbage-collected environment, triggers garbage collection if the memory allocated since the last collection is greater than the current threshold.
	Drain()
}

// Init initializes the instance.
func (a NSAutoreleasePool) Init() NSAutoreleasePool {
	rv := objc.Send[NSAutoreleasePool](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAutoreleasePool) Autorelease() NSAutoreleasePool {
	rv := objc.Send[NSAutoreleasePool](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAutoreleasePool creates a new NSAutoreleasePool instance.
func NewNSAutoreleasePool() NSAutoreleasePool {
	class := getNSAutoreleasePoolClass()
	rv := objc.Send[NSAutoreleasePool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// In a reference-counted environment, releases and pops the receiver; in a
// garbage-collected environment, triggers garbage collection if the memory
// allocated since the last collection is greater than the current threshold.
//
// # Discussion
// 
// In a reference-counted environment, this method behaves the same as
// [release]. Since an autorelease pool cannot be retained (see
// [NSAutoreleasePool]), this therefore causes the receiver to be deallocated.
// When an autorelease pool is deallocated, it sends a [release] message to
// all its autoreleased objects. If an object is added several times to the
// same pool, when the pool is deallocated it receives a [release] message for
// each time it was added.
// 
// # Special Considerations
// 
// In a garbage-collected environment, `release` is a no-op, so unless you do
// not want to give the collector a hint it is important to use [Drain] in any
// code that may be compiled for a garbage-collected environment.
//
// [release]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/release
//
// See: https://developer.apple.com/documentation/Foundation/NSAutoreleasePool/drain
func (a NSAutoreleasePool) Drain() {
	objc.Send[objc.ID](a.ID, objc.Sel("drain"))
}

// Adds a given object to the active autorelease pool in the current thread.
//
// anObject: The object to add to the active autorelease pool in the current thread.
//
// # Discussion
// 
// The same object may be added several times to the active pool and, when the
// pool is deallocated, it will receive a [release] message for each time it
// was added.
// 
// Normally you don’t invoke this method directly—you send [autorelease]
// to `object` instead.
//
// [autorelease]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/autorelease
// [release]: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/release
//
// See: https://developer.apple.com/documentation/Foundation/NSAutoreleasePool/addObject:-c.type.method
func (_NSAutoreleasePoolClass NSAutoreleasePoolClass) AddObjectWithAnObject(anObject objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_NSAutoreleasePoolClass.class), objc.Sel("addObject:"), anObject)
}
// Displays the state of the current thread’s autorelease pool stack to
// `stderr`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSAutoreleasePool/showPools
func (_NSAutoreleasePoolClass NSAutoreleasePoolClass) ShowPools() {
	objc.Send[objc.ID](objc.ID(_NSAutoreleasePoolClass.class), objc.Sel("showPools"))
}

