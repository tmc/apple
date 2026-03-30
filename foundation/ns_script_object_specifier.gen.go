// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScriptObjectSpecifier] class.
var (
	_NSScriptObjectSpecifierClass     NSScriptObjectSpecifierClass
	_NSScriptObjectSpecifierClassOnce sync.Once
)

func getNSScriptObjectSpecifierClass() NSScriptObjectSpecifierClass {
	_NSScriptObjectSpecifierClassOnce.Do(func() {
		_NSScriptObjectSpecifierClass = NSScriptObjectSpecifierClass{class: objc.GetClass("NSScriptObjectSpecifier")}
	})
	return _NSScriptObjectSpecifierClass
}

// GetNSScriptObjectSpecifierClass returns the class object for NSScriptObjectSpecifier.
func GetNSScriptObjectSpecifierClass() NSScriptObjectSpecifierClass {
	return getNSScriptObjectSpecifierClass()
}

type NSScriptObjectSpecifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSScriptObjectSpecifierClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScriptObjectSpecifierClass) Alloc() NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class used to represent natural language expressions.
//
// # Overview
//
// [NSScriptObjectSpecifier] is the abstract superclass for classes that
// instantiate objects called “object specifiers.” An object specifier
// represents an AppleScript reference form, which is a natural-language
// expression such as `words 10 through 20` or `front document` or `words
// whose color is red`.
//
// The scripting system maps these words or phrases to attributes and
// relationships of scriptable objects. A reference form rarely occurs in
// isolation; usually a script statement consists of a series of reference
// forms preceded by a command and typically connected to each other by `of`,
// such as:
//
// The expression `words whose color is blue of paragraph 10 of front
// document` specifies a location in the application’s AppleScript object
// model—the objects the application makes available to scripters. The
// classes of objects in the object model often closely match the classes of
// actual objects in the application, but they are not required to. An object
// specifier locates objects in the running application that correspond to the
// specified object model objects.
//
// Your application typically creates object specifiers when it implements the
// `objectSpecifier` method for its scriptable classes. That method is defined
// by the NSScriptObjectSpecifiers protocol.
//
// It is unlikely that you would ever need to create your own subclass of
// [NSScriptObjectSpecifier]; the set of valid AppleScript reference forms is
// determined by Apple Computer and object specifier classes are already
// implemented for this set. If for some reason you do need to create a
// subclass, you must override the primitive method
// [NSScriptObjectSpecifier.IndicesOfObjectsByEvaluatingWithContainerCount] to return indices to the
// elements within the container whose values are matched with the child
// specifier’s key. In addition, you probably need to declare any special
// instance variables and implement an initializer that invokes super’s
// designated initializer,
// [NSScriptObjectSpecifier.InitWithContainerClassDescriptionContainerSpecifierKey], and initializes
// these variables.
//
// For a comprehensive treatment of object specifiers, including sample code,
// see [Object Specifiers] in [Cocoa Scripting Guide].
//
// # Initializing an object specifier
//
//   - [NSScriptObjectSpecifier.InitWithContainerClassDescriptionContainerSpecifierKey]: Returns an [NSScriptObjectSpecifier] object initialized with the given attributes.
//   - [NSScriptObjectSpecifier.InitWithContainerSpecifierKey]: Returns an [NSScriptObjectSpecifier] object initialized with a given container specifier  and key.
//
// # Evaluating an object specifier
//
//   - [NSScriptObjectSpecifier.IndicesOfObjectsByEvaluatingWithContainerCount]: This primitive method must be overridden by subclasses to return a pointer to an array of indices identifying objects in the key of a given container that are identified by the receiver of the message.
//   - [NSScriptObjectSpecifier.ObjectsByEvaluatingSpecifier]: Returns the actual object represented by the nested series of object specifiers.
//   - [NSScriptObjectSpecifier.ObjectsByEvaluatingWithContainers]: Returns the actual object or objects specified by the receiver as evaluated in the context of given container object.
//
// # Getting, testing, and setting containers
//
//   - [NSScriptObjectSpecifier.ContainerClassDescription]: Sets the class description of the receiver’s container specifier to a given specifier.
//   - [NSScriptObjectSpecifier.SetContainerClassDescription]
//   - [NSScriptObjectSpecifier.ContainerIsObjectBeingTested]: Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object.
//   - [NSScriptObjectSpecifier.SetContainerIsObjectBeingTested]
//   - [NSScriptObjectSpecifier.ContainerIsRangeContainerObject]: Sets whether the receiver’s container is to be the container for a range specifier or a top-level object.
//   - [NSScriptObjectSpecifier.SetContainerIsRangeContainerObject]
//   - [NSScriptObjectSpecifier.ContainerSpecifier]: Sets the container specifier of the receiver.
//   - [NSScriptObjectSpecifier.SetContainerSpecifier]
//
// # Getting and setting child references
//
//   - [NSScriptObjectSpecifier.ChildSpecifier]: Sets the receiver’s child reference.
//   - [NSScriptObjectSpecifier.SetChildSpecifier]
//
// # Getting and setting object keys
//
//   - [NSScriptObjectSpecifier.Key]: Sets the key of the receiver.
//   - [NSScriptObjectSpecifier.SetKey]
//   - [NSScriptObjectSpecifier.KeyClassDescription]: Returns the class description of the objects specified by the receiver.
//
// # Getting evaluation errors
//
//   - [NSScriptObjectSpecifier.EvaluationErrorSpecifier]: Returns the object specifier in which an evaluation error occurred.
//   - [NSScriptObjectSpecifier.EvaluationErrorNumber]: Sets the value of the evaluation error.
//   - [NSScriptObjectSpecifier.SetEvaluationErrorNumber]
//
// # Getting a descriptor for the object specifier
//
//   - [NSScriptObjectSpecifier.Descriptor]: Returns an Apple event descriptor that represents the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier
//
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [Object Specifiers]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_object_specifiers/SAppsObjectSpecifiers.html#//apple_ref/doc/uid/TP40002164-CH3
type NSScriptObjectSpecifier struct {
	objectivec.Object
}

// NSScriptObjectSpecifierFromID constructs a [NSScriptObjectSpecifier] from an objc.ID.
//
// An abstract class used to represent natural language expressions.
func NSScriptObjectSpecifierFromID(id objc.ID) NSScriptObjectSpecifier {
	return NSScriptObjectSpecifier{objectivec.Object{ID: id}}
}

// NOTE: NSScriptObjectSpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScriptObjectSpecifier] class.
//
// # Initializing an object specifier
//
//   - [INSScriptObjectSpecifier.InitWithContainerClassDescriptionContainerSpecifierKey]: Returns an [NSScriptObjectSpecifier] object initialized with the given attributes.
//   - [INSScriptObjectSpecifier.InitWithContainerSpecifierKey]: Returns an [NSScriptObjectSpecifier] object initialized with a given container specifier  and key.
//
// # Evaluating an object specifier
//
//   - [INSScriptObjectSpecifier.IndicesOfObjectsByEvaluatingWithContainerCount]: This primitive method must be overridden by subclasses to return a pointer to an array of indices identifying objects in the key of a given container that are identified by the receiver of the message.
//   - [INSScriptObjectSpecifier.ObjectsByEvaluatingSpecifier]: Returns the actual object represented by the nested series of object specifiers.
//   - [INSScriptObjectSpecifier.ObjectsByEvaluatingWithContainers]: Returns the actual object or objects specified by the receiver as evaluated in the context of given container object.
//
// # Getting, testing, and setting containers
//
//   - [INSScriptObjectSpecifier.ContainerClassDescription]: Sets the class description of the receiver’s container specifier to a given specifier.
//   - [INSScriptObjectSpecifier.SetContainerClassDescription]
//   - [INSScriptObjectSpecifier.ContainerIsObjectBeingTested]: Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object.
//   - [INSScriptObjectSpecifier.SetContainerIsObjectBeingTested]
//   - [INSScriptObjectSpecifier.ContainerIsRangeContainerObject]: Sets whether the receiver’s container is to be the container for a range specifier or a top-level object.
//   - [INSScriptObjectSpecifier.SetContainerIsRangeContainerObject]
//   - [INSScriptObjectSpecifier.ContainerSpecifier]: Sets the container specifier of the receiver.
//   - [INSScriptObjectSpecifier.SetContainerSpecifier]
//
// # Getting and setting child references
//
//   - [INSScriptObjectSpecifier.ChildSpecifier]: Sets the receiver’s child reference.
//   - [INSScriptObjectSpecifier.SetChildSpecifier]
//
// # Getting and setting object keys
//
//   - [INSScriptObjectSpecifier.Key]: Sets the key of the receiver.
//   - [INSScriptObjectSpecifier.SetKey]
//   - [INSScriptObjectSpecifier.KeyClassDescription]: Returns the class description of the objects specified by the receiver.
//
// # Getting evaluation errors
//
//   - [INSScriptObjectSpecifier.EvaluationErrorSpecifier]: Returns the object specifier in which an evaluation error occurred.
//   - [INSScriptObjectSpecifier.EvaluationErrorNumber]: Sets the value of the evaluation error.
//   - [INSScriptObjectSpecifier.SetEvaluationErrorNumber]
//
// # Getting a descriptor for the object specifier
//
//   - [INSScriptObjectSpecifier.Descriptor]: Returns an Apple event descriptor that represents the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier
type INSScriptObjectSpecifier interface {
	objectivec.IObject
	NSCoding

	// Topic: Initializing an object specifier

	// Returns an [NSScriptObjectSpecifier] object initialized with the given attributes.
	InitWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSScriptObjectSpecifier
	// Returns an [NSScriptObjectSpecifier] object initialized with a given container specifier  and key.
	InitWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSScriptObjectSpecifier

	// Topic: Evaluating an object specifier

	// This primitive method must be overridden by subclasses to return a pointer to an array of indices identifying objects in the key of a given container that are identified by the receiver of the message.
	IndicesOfObjectsByEvaluatingWithContainerCount(container objectivec.IObject, count unsafe.Pointer) unsafe.Pointer
	// Returns the actual object represented by the nested series of object specifiers.
	ObjectsByEvaluatingSpecifier() objectivec.IObject
	// Returns the actual object or objects specified by the receiver as evaluated in the context of given container object.
	ObjectsByEvaluatingWithContainers(containers objectivec.IObject) objectivec.IObject

	// Topic: Getting, testing, and setting containers

	// Sets the class description of the receiver’s container specifier to a given specifier.
	ContainerClassDescription() INSScriptClassDescription
	SetContainerClassDescription(value INSScriptClassDescription)
	// Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object.
	ContainerIsObjectBeingTested() bool
	SetContainerIsObjectBeingTested(value bool)
	// Sets whether the receiver’s container is to be the container for a range specifier or a top-level object.
	ContainerIsRangeContainerObject() bool
	SetContainerIsRangeContainerObject(value bool)
	// Sets the container specifier of the receiver.
	ContainerSpecifier() INSScriptObjectSpecifier
	SetContainerSpecifier(value INSScriptObjectSpecifier)

	// Topic: Getting and setting child references

	// Sets the receiver’s child reference.
	ChildSpecifier() INSScriptObjectSpecifier
	SetChildSpecifier(value INSScriptObjectSpecifier)

	// Topic: Getting and setting object keys

	// Sets the key of the receiver.
	Key() string
	SetKey(value string)
	// Returns the class description of the objects specified by the receiver.
	KeyClassDescription() INSScriptClassDescription

	// Topic: Getting evaluation errors

	// Returns the object specifier in which an evaluation error occurred.
	EvaluationErrorSpecifier() INSScriptObjectSpecifier
	// Sets the value of the evaluation error.
	EvaluationErrorNumber() int
	SetEvaluationErrorNumber(value int)

	// Topic: Getting a descriptor for the object specifier

	// Returns an Apple event descriptor that represents the receiver.
	Descriptor() INSAppleEventDescriptor
}

// Init initializes the instance.
func (s NSScriptObjectSpecifier) Init() NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScriptObjectSpecifier) Autorelease() NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScriptObjectSpecifier creates a new NSScriptObjectSpecifier instance.
func NewNSScriptObjectSpecifier() NSScriptObjectSpecifier {
	class := getNSScriptObjectSpecifierClass()
	rv := objc.Send[NSScriptObjectSpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(coder:)
func NewScriptObjectSpecifierWithCoder(inCoder INSCoder) NSScriptObjectSpecifier {
	instance := getNSScriptObjectSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSScriptObjectSpecifierFromID(rv)
}

// Returns an [NSScriptObjectSpecifier] object initialized with the given
// attributes.
//
// # Return Value
//
// An [NSScriptObjectSpecifier] object initialized with container specifier
// `specifier`, key `key`, and the class description of the object specifier
// `classDescription`, derived from the value of the specifier’s key.
//
// # Discussion
//
// You should never pass `nil` for the value of `classDescription`. The
// receiver’s child reference is set to `nil`.
//
// This is the designated initializer for [NSScriptObjectSpecifier].
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)
func NewScriptObjectSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSScriptObjectSpecifier {
	instance := getNSScriptObjectSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return NSScriptObjectSpecifierFromID(rv)
}

// Returns an [NSScriptObjectSpecifier] object initialized with a given
// container specifier and key.
//
// # Return Value
//
// An [NSScriptObjectSpecifier] object initialized with container specifier
// `specifier` and key `key`.
//
// # Discussion
//
// The class description of the container is set automatically.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(containerSpecifier:key:)
func NewScriptObjectSpecifierWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSScriptObjectSpecifier {
	instance := getNSScriptObjectSpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return NSScriptObjectSpecifierFromID(rv)
}

// Returns a new object specifier for an Apple event descriptor.
//
// descriptor: An Apple event descriptor. The descriptor must have the type
// `typeObjectSpecifier`.
//
// # Return Value
//
// An object specifier, or `nil` if an error occurs.
//
// # Discussion
//
// If “ is invoked and fails during the execution of a script command,
// information about the error that caused the failure is recorded in
// `[NSScriptCommand currentCommand]`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(descriptor:)
func NewScriptObjectSpecifierWithDescriptor(descriptor INSAppleEventDescriptor) NSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](objc.ID(getNSScriptObjectSpecifierClass().class), objc.Sel("objectSpecifierWithDescriptor:"), descriptor)
	return NSScriptObjectSpecifierFromID(rv)
}

// Returns an [NSScriptObjectSpecifier] object initialized with the given
// attributes.
//
// # Return Value
//
// An [NSScriptObjectSpecifier] object initialized with container specifier
// `specifier`, key `key`, and the class description of the object specifier
// `classDescription`, derived from the value of the specifier’s key.
//
// # Discussion
//
// You should never pass `nil` for the value of `classDescription`. The
// receiver’s child reference is set to `nil`.
//
// This is the designated initializer for [NSScriptObjectSpecifier].
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)
func (s NSScriptObjectSpecifier) InitWithContainerClassDescriptionContainerSpecifierKey(classDesc INSScriptClassDescription, container INSScriptObjectSpecifier, property string) NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](s.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	return rv
}

// Returns an [NSScriptObjectSpecifier] object initialized with a given
// container specifier and key.
//
// # Return Value
//
// An [NSScriptObjectSpecifier] object initialized with container specifier
// `specifier` and key `key`.
//
// # Discussion
//
// The class description of the container is set automatically.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(containerSpecifier:key:)
func (s NSScriptObjectSpecifier) InitWithContainerSpecifierKey(container INSScriptObjectSpecifier, property string) NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](s.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	return rv
}

// This primitive method must be overridden by subclasses to return a pointer
// to an array of indices identifying objects in the key of a given container
// that are identified by the receiver of the message.
//
// # Discussion
//
// This primitive method must be overridden by subclasses to return a pointer
// to an array of indices identifying objects in the key of the container
// `aContainer` that are identified by the receiver of the message. The method
// uses key-value coding to obtain values based on the receiver’s key. It
// returns the number of such matching objects by indirection in `numRefs`. It
// returns `nil` directly and –1 via `numRefs` if all objects in the
// container (or the sole object) match the value of the receiver’s key.
// This method is invoked by [ObjectsByEvaluatingWithContainers]. The default
// implementation returns `nil` directly and –1 indirectly via `numRefs`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/indicesOfObjectsByEvaluating(withContainer:count:)
func (s NSScriptObjectSpecifier) IndicesOfObjectsByEvaluatingWithContainerCount(container objectivec.IObject, count unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("indicesOfObjectsByEvaluatingWithContainer:count:"), container, count)
	return rv
}

// Returns the actual object or objects specified by the receiver as evaluated
// in the context of given container object.
//
// # Return Value
//
// The actual object or objects specified by the receiver as evaluated in the
// context of its container object or objects (`containers`).
//
// # Discussion
//
// Invokes [IndicesOfObjectsByEvaluatingWithContainerCount] on `self` to get
// an array of pointers to indices of elements in `containers` that have
// values paired with the message receiver’s key. This method then uses
// key-value coding to obtain the object or objects associated with the key;
// it returns these objects or `nil` if there are no matching values in
// containers. If there are multiple matching values, they are returned in an
// [NSArray]; if matching values are `nil`, [NSNull] objects are substituted.
// If `containers` is an [NSArray], the method recursively evaluates each
// element in the array and returns an [NSArray] with evaluated objects
// (including [NSNulls]) in their corresponding slots.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/objectsByEvaluating(withContainers:)
func (s NSScriptObjectSpecifier) ObjectsByEvaluatingWithContainers(containers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectsByEvaluatingWithContainers:"), containers)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(coder:)
func (s NSScriptObjectSpecifier) InitWithCoder(inCoder INSCoder) NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](s.ID, objc.Sel("initWithCoder:"), inCoder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (s NSScriptObjectSpecifier) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the actual object represented by the nested series of object
// specifiers.
//
// # Return Value
//
// The actual object represented by the nested series of object specifiers.
//
// # Discussion
//
// Recursively obtains the next container in a nested series of object
// specifiers until it reaches the top-level container specifier (which is
// either an [NSWhoseSpecifier] or the application object), after which it
// begins evaluating each object specifier
// ([ObjectsByEvaluatingWithContainers]) going in the opposite direction
// (top-level to innermost) as it unwinds from the stack. Returns the actual
// object represented by the nested series of object specifiers. Returns `nil`
// if a container specifier could not be evaluated or if no top-level
// container specifier could be found. Thus `nil` can be a valid value or can
// indicate an error; you can use [EvaluationErrorNumber] to determine if and
// which error occurred and [EvaluationErrorSpecifier] to find the container
// specifier responsible for the error. In the normal course of command
// processing, this method is invoked by an [NSScriptCommand] object’s
// [EvaluatedArguments] and [EvaluatedReceivers] methods, which take as
// message receiver the innermost object specifier.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/objectsByEvaluatingSpecifier
func (s NSScriptObjectSpecifier) ObjectsByEvaluatingSpecifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectsByEvaluatingSpecifier"))
	return objectivec.Object{ID: rv}
}

// Sets the class description of the receiver’s container specifier to a
// given specifier.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/containerClassDescription
func (s NSScriptObjectSpecifier) ContainerClassDescription() INSScriptClassDescription {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("containerClassDescription"))
	return NSScriptClassDescriptionFromID(objc.ID(rv))
}
func (s NSScriptObjectSpecifier) SetContainerClassDescription(value INSScriptClassDescription) {
	objc.Send[struct{}](s.ID, objc.Sel("setContainerClassDescription:"), value)
}

// Sets whether the receiver’s container should be an object involved in a
// filter reference or the top-level object.
//
// # Discussion
//
// If the receiver’s container specifier is `nil` and `flag` is true, sets
// the receiver’s container to be an object involved in a filter reference
// (for example, `whose color is blue`). If the receiver’s container
// specifier is `nil` and `flag` is false, sets the receiver’s container to
// be the top-level object.
//
// If `flag` is true [ContainerIsRangeContainerObject] should not also be
// invoked with an argument of true.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/containerIsObjectBeingTested
func (s NSScriptObjectSpecifier) ContainerIsObjectBeingTested() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("containerIsObjectBeingTested"))
	return rv
}
func (s NSScriptObjectSpecifier) SetContainerIsObjectBeingTested(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setContainerIsObjectBeingTested:"), value)
}

// Sets whether the receiver’s container is to be the container for a range
// specifier or a top-level object.
//
// # Discussion
//
// If the receiver’s container specifier is `nil` and `flag` is true, sets
// the receiver’s container to be the container for a range specifier. If
// the receiver’s container specifier is `nil` and `flag` is false, sets the
// receiver’s container to be the top-level object.
//
// If `flag` is true, [ContainerIsObjectBeingTested] should not also be
// invoked with an argument of true.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/containerIsRangeContainerObject
func (s NSScriptObjectSpecifier) ContainerIsRangeContainerObject() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("containerIsRangeContainerObject"))
	return rv
}
func (s NSScriptObjectSpecifier) SetContainerIsRangeContainerObject(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setContainerIsRangeContainerObject:"), value)
}

// Sets the container specifier of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/container
func (s NSScriptObjectSpecifier) ContainerSpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("containerSpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}
func (s NSScriptObjectSpecifier) SetContainerSpecifier(value INSScriptObjectSpecifier) {
	objc.Send[struct{}](s.ID, objc.Sel("setContainerSpecifier:"), value)
}

// Sets the receiver’s child reference.
//
// # Discussion
//
// Do not invoke this method directly; it is automatically invoked by
// [ContainerSpecifier].
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/child
func (s NSScriptObjectSpecifier) ChildSpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("childSpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}
func (s NSScriptObjectSpecifier) SetChildSpecifier(value INSScriptObjectSpecifier) {
	objc.Send[struct{}](s.ID, objc.Sel("setChildSpecifier:"), value)
}

// Sets the key of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/key
func (s NSScriptObjectSpecifier) Key() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("key"))
	return NSStringFromID(rv).String()
}
func (s NSScriptObjectSpecifier) SetKey(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setKey:"), objc.String(value))
}

// Returns the class description of the objects specified by the receiver.
//
// # Return Value
//
// The class description of the objects specified by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/keyClassDescription
func (s NSScriptObjectSpecifier) KeyClassDescription() INSScriptClassDescription {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("keyClassDescription"))
	return NSScriptClassDescriptionFromID(objc.ID(rv))
}

// Returns the object specifier in which an evaluation error occurred.
//
// # Return Value
//
// The object specifier in which an evaluation error occurred.
//
// # Discussion
//
// The object specifier failing to evaluate could be the receiver or any
// container specifier “above” the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/evaluationError
func (s NSScriptObjectSpecifier) EvaluationErrorSpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("evaluationErrorSpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}

// Sets the value of the evaluation error.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/evaluationErrorNumber
func (s NSScriptObjectSpecifier) EvaluationErrorNumber() int {
	rv := objc.Send[int](s.ID, objc.Sel("evaluationErrorNumber"))
	return rv
}
func (s NSScriptObjectSpecifier) SetEvaluationErrorNumber(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setEvaluationErrorNumber:"), value)
}

// Returns an Apple event descriptor that represents the receiver.
//
// # Return Value
//
// An Apple event descriptor of type `typeObjectSpecifier`.
//
// # Discussion
//
// If the receiver was created with [ObjectSpecifierWithDescriptor], the
// passed-in descriptor is returned. Otherwise, a new descriptor is created
// and returned, autoreleased.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/descriptor
func (s NSScriptObjectSpecifier) Descriptor() INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("descriptor"))
	return NSAppleEventDescriptorFromID(objc.ID(rv))
}

// Protocol methods for NSCoding
