// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [XMLDTDNode] class.
var (
	_XMLDTDNodeClass     XMLDTDNodeClass
	_XMLDTDNodeClassOnce sync.Once
)

func getXMLDTDNodeClass() XMLDTDNodeClass {
	_XMLDTDNodeClassOnce.Do(func() {
		_XMLDTDNodeClass = XMLDTDNodeClass{class: objc.GetClass("NSXMLDTDNode")}
	})
	return _XMLDTDNodeClass
}

// GetXMLDTDNodeClass returns the class object for NSXMLDTDNode.
func GetXMLDTDNodeClass() XMLDTDNodeClass {
	return getXMLDTDNodeClass()
}

type XMLDTDNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (xc XMLDTDNodeClass) Alloc() XMLDTDNode {
	rv := objc.Send[XMLDTDNode](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// A representation of element, attribute-list, entity, and notation
// declarations in a Document Type Definition.
//
// # Overview
// 
// [NSXMLDTDNode] objects are the sole children of a [NSXMLDTD] object
// (possibly along with comment nodes and processing-instruction nodes). They
// themselves cannot have any children.
// 
// [NSXMLDTDNode] objects can be of four kinds—element, attribute-list,
// entity, or notation declaration—and can also be of a subkind, as
// specified by a [XMLDTDNode.DTDKind] constant. For example, a DTD
// entity-declaration node could represent an unparsed entity declaration
// ([XMLEntityUnparsedKind]) rather than a parameter entity declaration
// ([XMLEntityParameterKind]). You can use a DTD node’s subkind to help
// determine how to handle the value of the node.
// 
// You can create an [NSXMLDTDNode] object with the [InitWithXMLString]
// method, the [NSXMLNode] class method [DTDNodeWithXMLString], or with the
// [NSXMLNode] initializer [InitWithKindOptions] (in the latter method
// supplying the appropriate [XMLNode.Kind] constant).
// 
// Setting the object value or string value of an [NSXMLDTDNode] objects
// affects different parts of different kinds of declaration. See the related
// programming topic for more information.
//
// [XMLDTDNode.DTDKind]: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum
// [XMLNode.Kind]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum
//
// # Initializing an NSXMLDTDNode Object
//
//   - [XMLDTDNode.InitWithXMLString]: Returns an [NSXMLDTDNode] object initialized with the DTD declaration in a given string.
//
// # Managing the DTD Node Kind
//
//   - [XMLDTDNode.DTDKind]: Returns the receiver’s DTD kind.
//   - [XMLDTDNode.SetDTDKind]
//
// # Managing DTD Identifiers
//
//   - [XMLDTDNode.External]
//   - [XMLDTDNode.NotationName]: Returns the name of the notation associated with the receiver.
//   - [XMLDTDNode.SetNotationName]
//   - [XMLDTDNode.PublicID]: Returns the public identifier associated with the receiver.
//   - [XMLDTDNode.SetPublicID]
//   - [XMLDTDNode.SystemID]: Returns the system identifier associated with the receiver.
//   - [XMLDTDNode.SetSystemID]
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode
type XMLDTDNode struct {
	NSXMLNode
}

// XMLDTDNodeFromID constructs a [XMLDTDNode] from an objc.ID.
//
// A representation of element, attribute-list, entity, and notation
// declarations in a Document Type Definition.
func XMLDTDNodeFromID(id objc.ID) XMLDTDNode {
	return NSXMLDTDNode{NSXMLNode: NSXMLNodeFromID(id)}
}

// NSXMLDTDNodeFromID is an alias for [XMLDTDNodeFromID] for cross-framework compatibility.
func NSXMLDTDNodeFromID(id objc.ID) XMLDTDNode { return XMLDTDNodeFromID(id) }
// NOTE: XMLDTDNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [XMLDTDNode] class.
//
// # Initializing an NSXMLDTDNode Object
//
//   - [IXMLDTDNode.InitWithXMLString]: Returns an [NSXMLDTDNode] object initialized with the DTD declaration in a given string.
//
// # Managing the DTD Node Kind
//
//   - [IXMLDTDNode.DTDKind]: Returns the receiver’s DTD kind.
//   - [IXMLDTDNode.SetDTDKind]
//
// # Managing DTD Identifiers
//
//   - [IXMLDTDNode.External]
//   - [IXMLDTDNode.NotationName]: Returns the name of the notation associated with the receiver.
//   - [IXMLDTDNode.SetNotationName]
//   - [IXMLDTDNode.PublicID]: Returns the public identifier associated with the receiver.
//   - [IXMLDTDNode.SetPublicID]
//   - [IXMLDTDNode.SystemID]: Returns the system identifier associated with the receiver.
//   - [IXMLDTDNode.SetSystemID]
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode
type IXMLDTDNode interface {
	INSXMLNode

	// Topic: Initializing an NSXMLDTDNode Object

	// Returns an [NSXMLDTDNode] object initialized with the DTD declaration in a given string.
	InitWithXMLString(string_ string) XMLDTDNode

	// Topic: Managing the DTD Node Kind

	// Returns the receiver’s DTD kind.
	DTDKind() NSXMLDTDNodeKind
	SetDTDKind(value NSXMLDTDNodeKind)

	// Topic: Managing DTD Identifiers

	External() bool
	// Returns the name of the notation associated with the receiver.
	NotationName() string
	SetNotationName(value string)
	// Returns the public identifier associated with the receiver.
	PublicID() string
	SetPublicID(value string)
	// Returns the system identifier associated with the receiver.
	SystemID() string
	SetSystemID(value string)
}

// Init initializes the instance.
func (x XMLDTDNode) Init() XMLDTDNode {
	rv := objc.Send[XMLDTDNode](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XMLDTDNode) Autorelease() XMLDTDNode {
	rv := objc.Send[XMLDTDNode](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLDTDNode creates a new XMLDTDNode instance.
func NewXMLDTDNode() XMLDTDNode {
	class := getXMLDTDNodeClass()
	rv := objc.Send[XMLDTDNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an [NSXMLNode] instance initialized with the constant indicating
// node kind.
//
// kind: An `enum` constant of type [XMLNode.Kind] that indicates the type of node.
// See Constants for a list of valid NSXMLNodeKind constants.
// //
// [XMLNode.Kind]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum
//
// # Return Value
// 
// An [NSXMLNode] object initialized with kind or `nil` if the object
// couldn’t be created. If `kind` is not a valid NSXMLNodeKind constant, the
// method returns an [NSXMLNode] object of kind [NSXMLInvalidKind].
//
// # Discussion
// 
// This method invokes [InitWithKindOptions] with the `options` parameter set
// to [NSXMLNodeOptionsNone].
// 
// Do not use this initializer for creating instances of [NSXMLDTDNode] for
// attribute-list declarations. Instead, use the [DTDNodeWithXMLString] class
// method of this class or the [InitWithXMLString] method of the
// [NSXMLDTDNode] class.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/init(kind:)
func NewXMLDTDNodeWithKind(kind NSXMLNodeKind) XMLDTDNode {
	instance := getXMLDTDNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:"), kind)
	return XMLDTDNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode/init(kind:options:)
func NewXMLDTDNodeWithKindOptions(kind NSXMLNodeKind, options NSXMLNodeOptions) XMLDTDNode {
	instance := getXMLDTDNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:options:"), kind, options)
	return XMLDTDNodeFromID(rv)
}

// Returns an [NSXMLDTDNode] object initialized with the DTD declaration in a
// given string.
//
// string: The DTD declaration.
//
// # Return Value
// 
// An [NSXMLDTDNode] object initialized with the DTD declaration in `string`.
// Returns `nil` if initialization did not succeed, as might occur if the
// passed-in declaration is malformed.
//
// # Discussion
// 
// The node kind (NSXMLNode) assigned to the returned object—element,
// attribute, entity, or notation declaration— is based on the full XML
// string that is parsed. To assign a subkind, set the [DTDKind] property.
// 
// You may also use the [DTDNodeWithXMLString] or [InitWithKind] methods to
// create [NSXMLDTDNode] instances. However, you cannot use the latter method
// to create [NSXMLDTDNode] instances for attribute-list declarations.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode/init(xmlString:)
func NewXMLDTDNodeWithXMLString(string_ string) XMLDTDNode {
	instance := getXMLDTDNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXMLString:"), objc.String(string_))
	return XMLDTDNodeFromID(rv)
}

// Returns an [NSXMLDTDNode] object initialized with the DTD declaration in a
// given string.
//
// string: The DTD declaration.
//
// # Return Value
// 
// An [NSXMLDTDNode] object initialized with the DTD declaration in `string`.
// Returns `nil` if initialization did not succeed, as might occur if the
// passed-in declaration is malformed.
//
// # Discussion
// 
// The node kind (NSXMLNode) assigned to the returned object—element,
// attribute, entity, or notation declaration— is based on the full XML
// string that is parsed. To assign a subkind, set the [DTDKind] property.
// 
// You may also use the [DTDNodeWithXMLString] or [InitWithKind] methods to
// create [NSXMLDTDNode] instances. However, you cannot use the latter method
// to create [NSXMLDTDNode] instances for attribute-list declarations.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode/init(xmlString:)
func (x XMLDTDNode) InitWithXMLString(string_ string) XMLDTDNode {
	rv := objc.Send[XMLDTDNode](x.ID, objc.Sel("initWithXMLString:"), objc.String(string_))
	return rv
}

// Returns the receiver’s DTD kind.
//
// # Return Value
// 
// The receiver’s DTD kind. See Constants for a list of valid
// NSXMLDTDNodeKind constants.
// 
// # Discussion
// 
// The DTD kind is distinct from a [NSXMLDTDNode] object’s node kind
// (returned by the [NSXMLNode] [Kind] method).
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode/dtdKind-swift.property
func (x XMLDTDNode) DTDKind() NSXMLDTDNodeKind {
	rv := objc.Send[NSXMLDTDNodeKind](x.ID, objc.Sel("DTDKind"))
	return NSXMLDTDNodeKind(rv)
}
func (x XMLDTDNode) SetDTDKind(value NSXMLDTDNodeKind) {
	objc.Send[struct{}](x.ID, objc.Sel("setDTDKind:"), value)
}

//
// # Discussion
// 
// True if the system id is set. Valid for entities and notations.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode/isExternal
func (x XMLDTDNode) External() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("isExternal"))
	return rv
}

// Returns the name of the notation associated with the receiver.
//
// # Return Value
// 
// The name of the notation associated with the receiver.
// 
// # Discussion
// 
// Notations are applicable to unparsed external entities, processing
// instructions, and some attribute values.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode/notationName
func (x XMLDTDNode) NotationName() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("notationName"))
	return NSStringFromID(rv).String()
}
func (x XMLDTDNode) SetNotationName(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setNotationName:"), objc.String(value))
}

// Returns the public identifier associated with the receiver.
//
// # Return Value
// 
// The public identifier associated with the receiver.
// 
// # Discussion
// 
// The public ID is applicable to entities and notations.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode/publicID
func (x XMLDTDNode) PublicID() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("publicID"))
	return NSStringFromID(rv).String()
}
func (x XMLDTDNode) SetPublicID(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setPublicID:"), objc.String(value))
}

// Returns the system identifier associated with the receiver.
//
// # Return Value
// 
// The system identifier associated with the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode/systemID
func (x XMLDTDNode) SystemID() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("systemID"))
	return NSStringFromID(rv).String()
}
func (x XMLDTDNode) SetSystemID(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setSystemID:"), objc.String(value))
}

