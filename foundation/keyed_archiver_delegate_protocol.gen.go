// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// The optional methods implemented by the delegate of a keyed archiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiverDelegate
type NSKeyedArchiverDelegate interface {
	objectivec.IObject
}

// NSKeyedArchiverDelegateObject wraps an existing Objective-C object that conforms to the NSKeyedArchiverDelegate protocol.
type NSKeyedArchiverDelegateObject struct {
	objectivec.Object
}
func (o NSKeyedArchiverDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSKeyedArchiverDelegateObjectFromID constructs a [NSKeyedArchiverDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSKeyedArchiverDelegateObjectFromID(id objc.ID) NSKeyedArchiverDelegateObject {
	return NSKeyedArchiverDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Informs the delegate that a given object has been encoded.
//
// archiver: The archiver that sent the message.
//
// object: The object that has been encoded. `object` may be `nil`.
//
// # Discussion
// 
// The delegate might restore some state it had modified previously, or use
// this opportunity to keep track of the objects that are encoded.
// 
// This method is not called for conditional objects until they are actually
// encoded (if ever).
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiverDelegate/archiver(_:didEncode:)
func (o NSKeyedArchiverDelegateObject) ArchiverDidEncodeObject(archiver INSKeyedArchiver, object objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("archiver:didEncodeObject:"), archiver, object)
	}
// Notifies the delegate that encoding has finished.
//
// archiver: The archiver that sent the message.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiverDelegate/archiverDidFinish(_:)
func (o NSKeyedArchiverDelegateObject) ArchiverDidFinish(archiver INSKeyedArchiver) {
	objc.Send[struct{}](o.ID, objc.Sel("archiverDidFinish:"), archiver)
	}
// Informs the delegate that `object` is about to be encoded.
//
// archiver: The archiver that sent the message.
//
// object: The object that is about to be encoded. This value is never `nil`.
//
// # Return Value
// 
// Either `object` or a different object to be encoded in its stead. The
// delegate can also modify the coder state. If the delegate returns `nil`,
// `nil` is encoded.
//
// # Discussion
// 
// This method is called after the original object may have replaced itself
// with [replacementObject(for:)]:.
// 
// This method is called whether or not the object is being encoded
// conditionally.
// 
// This method is not called for an object once a replacement mapping has been
// set up for that object (either explicitly, or because the object has
// previously been encoded). This method is also not called when `nil` is
// about to be encoded.
//
// [replacementObject(for:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replacementObject(for:)-60vwc
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiverDelegate/archiver(_:willEncode:)
func (o NSKeyedArchiverDelegateObject) ArchiverWillEncodeObject(archiver INSKeyedArchiver, object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("archiver:willEncodeObject:"), archiver, object)
	return objectivec.Object{ID: rv}
	}
// Notifies the delegate that encoding is about to finish.
//
// archiver: The archiver that sent the message.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiverDelegate/archiverWillFinish(_:)
func (o NSKeyedArchiverDelegateObject) ArchiverWillFinish(archiver INSKeyedArchiver) {
	objc.Send[struct{}](o.ID, objc.Sel("archiverWillFinish:"), archiver)
	}
// Informs the delegate that one given object is being substituted for another
// given object.
//
// archiver: The archiver that sent the message.
//
// object: The object being replaced in the archive.
//
// newObject: The object replacing `object` in the archive.
//
// # Discussion
// 
// This method is called even when the delegate itself is doing, or has done,
// the substitution. The delegate may use this method if it is keeping track
// of the encoded or decoded objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyedArchiverDelegate/archiver(_:willReplace:with:)
func (o NSKeyedArchiverDelegateObject) ArchiverWillReplaceObjectWithObject(archiver INSKeyedArchiver, object objectivec.IObject, newObject objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("archiver:willReplaceObject:withObject:"), archiver, object, newObject)
	}

// NSKeyedArchiverDelegateConfig holds optional typed callbacks for [NSKeyedArchiverDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate
type NSKeyedArchiverDelegateConfig struct {

	// Encoding Data and Objects
	// ArchiverDidFinish — Notifies the delegate that encoding has finished.
	ArchiverDidFinish func(archiver NSKeyedArchiver)
	// ArchiverWillFinish — Notifies the delegate that encoding is about to finish.
	ArchiverWillFinish func(archiver NSKeyedArchiver)
}

// NewNSKeyedArchiverDelegate creates an Objective-C object implementing the [NSKeyedArchiverDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSKeyedArchiverDelegateObject] satisfies the [NSKeyedArchiverDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nskeyedarchiverdelegate
func NewNSKeyedArchiverDelegate(config NSKeyedArchiverDelegateConfig) NSKeyedArchiverDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSKeyedArchiverDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ArchiverDidFinish != nil {
		fn := config.ArchiverDidFinish
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("archiverDidFinish:"),
			Fn: func(self objc.ID, _cmd objc.SEL, archiverID objc.ID) {
				archiver := NSKeyedArchiverFromID(archiverID)
				fn(archiver)
			},
		})
	}

	if config.ArchiverWillFinish != nil {
		fn := config.ArchiverWillFinish
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("archiverWillFinish:"),
			Fn: func(self objc.ID, _cmd objc.SEL, archiverID objc.ID) {
				archiver := NSKeyedArchiverFromID(archiverID)
				fn(archiver)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSKeyedArchiverDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSKeyedArchiverDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSKeyedArchiverDelegateObjectFromID(instance)
}

