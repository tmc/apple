// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// The optional methods implemented by the delegate of a keyed unarchiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiverDelegate
type NSKeyedUnarchiverDelegate interface {
	objectivec.IObject
}

// NSKeyedUnarchiverDelegateObject wraps an existing Objective-C object that conforms to the NSKeyedUnarchiverDelegate protocol.
type NSKeyedUnarchiverDelegateObject struct {
	objectivec.Object
}
func (o NSKeyedUnarchiverDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSKeyedUnarchiverDelegateObjectFromID constructs a [NSKeyedUnarchiverDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSKeyedUnarchiverDelegateObjectFromID(id objc.ID) NSKeyedUnarchiverDelegateObject {
	return NSKeyedUnarchiverDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Informs the delegate that the class with a given name is not available
// during decoding.
//
// unarchiver: An unarchiver for which the receiver is the delegate.
//
// name: The name of the class of an object `unarchiver` is trying to decode.
//
// classNames: An array describing the class hierarchy of the encoded object, where the
// first element is the class name string of the encoded object, the second
// element is the class name of its immediate superclass, and so on.
//
// # Return Value
// 
// The class unarchiver should use in place of the class named `name`.
//
// # Discussion
// 
// The delegate may, for example, load some code to introduce the class to the
// runtime and return the class, or substitute a different class object. If
// the delegate returns `nil`, unarchiving aborts and the method raises an
// [NSInvalidUnarchiveOperationException].
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiverDelegate/unarchiver(_:cannotDecodeObjectOfClassName:originalClasses:)
func (o NSKeyedUnarchiverDelegateObject) UnarchiverCannotDecodeObjectOfClassNameOriginalClasses(unarchiver INSKeyedUnarchiver, name string, classNames []string) objc.Class {
	rv := objc.Send[objc.Class](o.ID, objc.Sel("unarchiver:cannotDecodeObjectOfClassName:originalClasses:"), unarchiver, objc.String(name), objectivec.StringSliceToNSArray(classNames))
	return rv
	}
// Informs the delegate that a given object has been decoded.
//
// unarchiver: An unarchiver for which the receiver is the delegate.
//
// object: The object that has been decoded. `object` may be `nil`.
//
// # Return Value
// 
// The object to use in place of `object`. The delegate can either return
// `object` or return a different object to replace the decoded one. In apps
// using ARC, the delegate should only return `nil` if `object` itself is
// `nil`. In apps not using ARC, the delegate can return `nil` to indicate
// that the decoded value is unchanged—that is, `object` will be decoded.
//
// # Discussion
// 
// This method is called after `object` has been sent [InitWithCoder] and
// [awakeAfter(using:)].
// 
// The delegate may use this method to keep track of the decoded objects.
//
// [awakeAfter(using:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/awakeAfter(using:)
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiverDelegate/unarchiver(_:didDecode:)
func (o NSKeyedUnarchiverDelegateObject) UnarchiverDidDecodeObject(unarchiver INSKeyedUnarchiver, object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("unarchiver:didDecodeObject:"), unarchiver, object)
	return objectivec.Object{ID: rv}
	}
// Informs the delegate that one object is being substituted for another.
//
// unarchiver: An unarchiver for which the receiver is the delegate.
//
// object: An object in the archive.
//
// newObject: The object with which `unarchiver` will replace `object`.
//
// # Discussion
// 
// This method is called even when the delegate itself is doing, or has done,
// the substitution with [UnarchiverDidDecodeObject].
// 
// The delegate may use this method if it is keeping track of the encoded or
// decoded objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiverDelegate/unarchiver(_:willReplace:with:)
func (o NSKeyedUnarchiverDelegateObject) UnarchiverWillReplaceObjectWithObject(unarchiver INSKeyedUnarchiver, object objectivec.IObject, newObject objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("unarchiver:willReplaceObject:withObject:"), unarchiver, object, newObject)
	}
// Notifies the delegate that decoding has finished.
//
// unarchiver: An unarchiver for which the receiver is the delegate.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiverDelegate/unarchiverDidFinish(_:)
func (o NSKeyedUnarchiverDelegateObject) UnarchiverDidFinish(unarchiver INSKeyedUnarchiver) {
	objc.Send[struct{}](o.ID, objc.Sel("unarchiverDidFinish:"), unarchiver)
	}
// Notifies the delegate that decoding is about to finish.
//
// unarchiver: An unarchiver for which the receiver is the delegate.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedUnarchiverDelegate/unarchiverWillFinish(_:)
func (o NSKeyedUnarchiverDelegateObject) UnarchiverWillFinish(unarchiver INSKeyedUnarchiver) {
	objc.Send[struct{}](o.ID, objc.Sel("unarchiverWillFinish:"), unarchiver)
	}

// NSKeyedUnarchiverDelegateConfig holds optional typed callbacks for [NSKeyedUnarchiverDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate
type NSKeyedUnarchiverDelegateConfig struct {

	// Finishing Decoding
	// UnarchiverDidFinish — Notifies the delegate that decoding has finished.
	UnarchiverDidFinish func(unarchiver NSKeyedUnarchiver)
	// UnarchiverWillFinish — Notifies the delegate that decoding is about to finish.
	UnarchiverWillFinish func(unarchiver NSKeyedUnarchiver)
}

// NewNSKeyedUnarchiverDelegate creates an Objective-C object implementing the [NSKeyedUnarchiverDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSKeyedUnarchiverDelegateObject] satisfies the [NSKeyedUnarchiverDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nskeyedunarchiverdelegate
func NewNSKeyedUnarchiverDelegate(config NSKeyedUnarchiverDelegateConfig) NSKeyedUnarchiverDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSKeyedUnarchiverDelegate_%d", n)

	var methods []objc.MethodDef

	if config.UnarchiverDidFinish != nil {
		fn := config.UnarchiverDidFinish
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("unarchiverDidFinish:"),
			Fn: func(self objc.ID, _cmd objc.SEL, unarchiverID objc.ID) {
				unarchiver := NSKeyedUnarchiverFromID(unarchiverID)
				fn(unarchiver)
			},
		})
	}

	if config.UnarchiverWillFinish != nil {
		fn := config.UnarchiverWillFinish
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("unarchiverWillFinish:"),
			Fn: func(self objc.ID, _cmd objc.SEL, unarchiverID objc.ID) {
				unarchiver := NSKeyedUnarchiverFromID(unarchiverID)
				fn(unarchiver)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSKeyedUnarchiverDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSKeyedUnarchiverDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSKeyedUnarchiverDelegateObjectFromID(instance)
}

