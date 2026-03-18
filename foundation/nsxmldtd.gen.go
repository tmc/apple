// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XMLDTD] class.
var (
	_XMLDTDClass     XMLDTDClass
	_XMLDTDClassOnce sync.Once
)

func getXMLDTDClass() XMLDTDClass {
	_XMLDTDClassOnce.Do(func() {
		_XMLDTDClass = XMLDTDClass{class: objc.GetClass("NSXMLDTD")}
	})
	return _XMLDTDClass
}

// GetXMLDTDClass returns the class object for NSXMLDTD.
func GetXMLDTDClass() XMLDTDClass {
	return getXMLDTDClass()
}

type XMLDTDClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (xc XMLDTDClass) Alloc() XMLDTD {
	rv := objc.Send[XMLDTD](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a Document Type Definition.
//
// # Overview
// 
// An instance of the [NSXMLDTD] class is held as a property of an
// [NSXMLDocument] instance, accessed through the [NSXMLDocument] property
// [DTD].
// 
// In the data model, an [NSXMLDTD] object is conceptually similar to
// namespace and attribute nodes: it is not considered to be a child of the
// [NSXMLDocument] object although it is closely associated with it. It is at
// the “root” of a shallow tree consisting primarily of nodes representing
// DTD declarations. Acceptable child nodes are instances of the
// [NSXMLDTDNode] class as well as [NSXMLNode] objects representing comment
// nodes and processing-instruction nodes.
// 
// You create an [NSXMLDTD] object in one of three ways:
// 
// - By processing an XML document with its own internal (in-line) DTD - By
// process a standalone (external) DTD - Programmatically
// 
// Once an [NSXMLDTD] instance is in place, you can add, remove, and change
// the [NSXMLDTDNode] objects representing various DTD declarations. When you
// write the document out as XML, the new or modified internal DTD is included
// (assuming you set the DTD in the [NSXMLDocument] instance). You may also
// programmatically create an external DTD and write that out to its own file.
//
// # Initializing an NSXMLDTD Object
//
//   - [XMLDTD.InitWithContentsOfURLOptionsError]: Initializes and returns an [NSXMLDTD] object created from the DTD declarations in a URL-referenced source.
//   - [XMLDTD.InitWithDataOptionsError]: Initializes and returns an [NSXMLDTD] object created from the DTD declarations encapsulated in an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object
//
// # Managing DTD Identifiers
//
//   - [XMLDTD.PublicID]: Returns the receiver’s public identifier.
//   - [XMLDTD.SetPublicID]
//   - [XMLDTD.SystemID]: Returns the receiver’s system identifier.
//   - [XMLDTD.SetSystemID]
//
// # Manipulating Child Nodes
//
//   - [XMLDTD.AddChild]: Adds a child node to the end of the list of existing children.
//   - [XMLDTD.InsertChildAtIndex]: Inserts a child node in the receiver’s list of children at a specific location in the list.
//   - [XMLDTD.InsertChildrenAtIndex]: Inserts an array of child nodes at a specified location in the receiver’s list of children.
//   - [XMLDTD.RemoveChildAtIndex]: Removes the child node at a particular location in the receiver’s list of children.
//   - [XMLDTD.ReplaceChildAtIndexWithNode]: Replaces a child at a particular index with another child.
//
// # Getting DTD Nodes by Name
//
//   - [XMLDTD.ElementDeclarationForName]: Returns the DTD node representing an element declaration for a specified element.
//   - [XMLDTD.AttributeDeclarationForNameElementName]: Returns the DTD node representing an attribute-list declaration for a given attribute and its element.
//   - [XMLDTD.EntityDeclarationForName]: Returns the DTD node representing the entity declaration for a specified entity.
//   - [XMLDTD.NotationDeclarationForName]: Returns the DTD node representing the notation declaration identified by the specified notation name.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD
type XMLDTD struct {
	NSXMLNode
}

// XMLDTDFromID constructs a [XMLDTD] from an objc.ID.
//
// A representation of a Document Type Definition.
func XMLDTDFromID(id objc.ID) XMLDTD {
	return NSXMLDTD{NSXMLNode: NSXMLNodeFromID(id)}
}

// NSXMLDTDFromID is an alias for [XMLDTDFromID] for cross-framework compatibility.
func NSXMLDTDFromID(id objc.ID) XMLDTD { return XMLDTDFromID(id) }
// NOTE: XMLDTD adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [XMLDTD] class.
//
// # Initializing an NSXMLDTD Object
//
//   - [IXMLDTD.InitWithContentsOfURLOptionsError]: Initializes and returns an [NSXMLDTD] object created from the DTD declarations in a URL-referenced source.
//   - [IXMLDTD.InitWithDataOptionsError]: Initializes and returns an [NSXMLDTD] object created from the DTD declarations encapsulated in an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object
//
// # Managing DTD Identifiers
//
//   - [IXMLDTD.PublicID]: Returns the receiver’s public identifier.
//   - [IXMLDTD.SetPublicID]
//   - [IXMLDTD.SystemID]: Returns the receiver’s system identifier.
//   - [IXMLDTD.SetSystemID]
//
// # Manipulating Child Nodes
//
//   - [IXMLDTD.AddChild]: Adds a child node to the end of the list of existing children.
//   - [IXMLDTD.InsertChildAtIndex]: Inserts a child node in the receiver’s list of children at a specific location in the list.
//   - [IXMLDTD.InsertChildrenAtIndex]: Inserts an array of child nodes at a specified location in the receiver’s list of children.
//   - [IXMLDTD.RemoveChildAtIndex]: Removes the child node at a particular location in the receiver’s list of children.
//   - [IXMLDTD.ReplaceChildAtIndexWithNode]: Replaces a child at a particular index with another child.
//
// # Getting DTD Nodes by Name
//
//   - [IXMLDTD.ElementDeclarationForName]: Returns the DTD node representing an element declaration for a specified element.
//   - [IXMLDTD.AttributeDeclarationForNameElementName]: Returns the DTD node representing an attribute-list declaration for a given attribute and its element.
//   - [IXMLDTD.EntityDeclarationForName]: Returns the DTD node representing the entity declaration for a specified entity.
//   - [IXMLDTD.NotationDeclarationForName]: Returns the DTD node representing the notation declaration identified by the specified notation name.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD
type IXMLDTD interface {
	INSXMLNode

	// Topic: Initializing an NSXMLDTD Object

	// Initializes and returns an [NSXMLDTD] object created from the DTD declarations in a URL-referenced source.
	InitWithContentsOfURLOptionsError(url INSURL, mask NSXMLNodeOptions) (XMLDTD, error)
	// Initializes and returns an [NSXMLDTD] object created from the DTD declarations encapsulated in an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object
	InitWithDataOptionsError(data INSData, mask NSXMLNodeOptions) (XMLDTD, error)

	// Topic: Managing DTD Identifiers

	// Returns the receiver’s public identifier.
	PublicID() string
	SetPublicID(value string)
	// Returns the receiver’s system identifier.
	SystemID() string
	SetSystemID(value string)

	// Topic: Manipulating Child Nodes

	// Adds a child node to the end of the list of existing children.
	AddChild(child INSXMLNode)
	// Inserts a child node in the receiver’s list of children at a specific location in the list.
	InsertChildAtIndex(child INSXMLNode, index uint)
	// Inserts an array of child nodes at a specified location in the receiver’s list of children.
	InsertChildrenAtIndex(children []NSXMLNode, index uint)
	// Removes the child node at a particular location in the receiver’s list of children.
	RemoveChildAtIndex(index uint)
	// Replaces a child at a particular index with another child.
	ReplaceChildAtIndexWithNode(index uint, node INSXMLNode)

	// Topic: Getting DTD Nodes by Name

	// Returns the DTD node representing an element declaration for a specified element.
	ElementDeclarationForName(name string) INSXMLDTDNode
	// Returns the DTD node representing an attribute-list declaration for a given attribute and its element.
	AttributeDeclarationForNameElementName(name string, elementName string) INSXMLDTDNode
	// Returns the DTD node representing the entity declaration for a specified entity.
	EntityDeclarationForName(name string) INSXMLDTDNode
	// Returns the DTD node representing the notation declaration identified by the specified notation name.
	NotationDeclarationForName(name string) INSXMLDTDNode

	// Returns an 
	Dtd() INSXMLDTD
	SetDtd(value INSXMLDTD)
}

// Init initializes the instance.
func (x XMLDTD) Init() XMLDTD {
	rv := objc.Send[XMLDTD](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XMLDTD) Autorelease() XMLDTD {
	rv := objc.Send[XMLDTD](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLDTD creates a new XMLDTD instance.
func NewXMLDTD() XMLDTD {
	class := getXMLDTDClass()
	rv := objc.Send[XMLDTD](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns an [NSXMLDTD] object created from the DTD
// declarations in a URL-referenced source.
//
// url: An [NSURL] object identifying a URL source.
//
// mask: A bit mask specifying input options; bit-OR multiple options. The current
// valid options are [NSXMLNodePreserveWhitespace] and
// [NSXMLNodePreserveEntities]; these constants are described in the
// “Constants” section of the [NSXMLNode] reference.
//
// # Return Value
// 
// An initialized [NSXMLDTD] object or `nil` if initialization fails because
// of parsing errors or other reasons.
//
// # Discussion
// 
// You use this method to create a stand-alone DTD which you can thereafter
// query and use for validation. You can associate the DTD created through
// this message with a document by setting the [DTD] property on an
// [NSXMLDocument] object.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/init(contentsOf:options:)
func NewXMLDTDWithContentsOfURLOptionsError(url INSURL, mask NSXMLNodeOptions) (XMLDTD, error) {
	var errorPtr objc.ID
	instance := getXMLDTDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDTDFromID(rv), NSErrorFrom(errorPtr)
	}
	return XMLDTDFromID(rv), nil
}

// Initializes and returns an [NSXMLDTD] object created from the DTD
// declarations encapsulated in an [NSData] object
//
// data: A data object containing DTD declarations.
//
// mask: A bit mask specifying input options; bit-OR multiple options. The current
// valid options are [NSXMLNodePreserveWhitespace] and
// [NSXMLNodePreserveEntities]; these constants are described in the
// “Constants” section of the [NSXMLNode] reference.
//
// # Return Value
// 
// An initialized [NSXMLDTD] object or `nil` if initialization fails because
// of parsing errors or other reasons.
//
// # Discussion
// 
// This method is the designated initializer for the [NSXMLDTD] class. You use
// this method to create a stand-alone DTD which you can thereafter query and
// use for validation. You can associate the DTD created through this message
// with a document by setting the [DTD] property on an [NSXMLDocument] object.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/init(data:options:)
func NewXMLDTDWithDataOptionsError(data INSData, mask NSXMLNodeOptions) (XMLDTD, error) {
	var errorPtr objc.ID
	instance := getXMLDTDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:error:"), data, mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDTDFromID(rv), NSErrorFrom(errorPtr)
	}
	return XMLDTDFromID(rv), nil
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
func NewXMLDTDWithKind(kind NSXMLNodeKind) XMLDTD {
	instance := getXMLDTDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:"), kind)
	return XMLDTDFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSXMLDTD/initWithKind:options:
func NewXMLDTDWithKindOptions(kind NSXMLNodeKind, options NSXMLNodeOptions) XMLDTD {
	instance := getXMLDTDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:options:"), kind, options)
	return XMLDTDFromID(rv)
}

// Initializes and returns an [NSXMLDTD] object created from the DTD
// declarations in a URL-referenced source.
//
// url: An [NSURL] object identifying a URL source.
//
// mask: A bit mask specifying input options; bit-OR multiple options. The current
// valid options are [NSXMLNodePreserveWhitespace] and
// [NSXMLNodePreserveEntities]; these constants are described in the
// “Constants” section of the [NSXMLNode] reference.
//
// # Return Value
// 
// An initialized [NSXMLDTD] object or `nil` if initialization fails because
// of parsing errors or other reasons.
//
// # Discussion
// 
// You use this method to create a stand-alone DTD which you can thereafter
// query and use for validation. You can associate the DTD created through
// this message with a document by setting the [DTD] property on an
// [NSXMLDocument] object.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/init(contentsOf:options:)
func (x XMLDTD) InitWithContentsOfURLOptionsError(url INSURL, mask NSXMLNodeOptions) (XMLDTD, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDTD{}, NSErrorFrom(errorPtr)
	}
	return NSXMLDTDFromID(rv), nil

}

// Initializes and returns an [NSXMLDTD] object created from the DTD
// declarations encapsulated in an [NSData] object
//
// data: A data object containing DTD declarations.
//
// mask: A bit mask specifying input options; bit-OR multiple options. The current
// valid options are [NSXMLNodePreserveWhitespace] and
// [NSXMLNodePreserveEntities]; these constants are described in the
// “Constants” section of the [NSXMLNode] reference.
//
// # Return Value
// 
// An initialized [NSXMLDTD] object or `nil` if initialization fails because
// of parsing errors or other reasons.
//
// # Discussion
// 
// This method is the designated initializer for the [NSXMLDTD] class. You use
// this method to create a stand-alone DTD which you can thereafter query and
// use for validation. You can associate the DTD created through this message
// with a document by setting the [DTD] property on an [NSXMLDocument] object.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/init(data:options:)
func (x XMLDTD) InitWithDataOptionsError(data INSData, mask NSXMLNodeOptions) (XMLDTD, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("initWithData:options:error:"), data, mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDTD{}, NSErrorFrom(errorPtr)
	}
	return NSXMLDTDFromID(rv), nil

}

// Adds a child node to the end of the list of existing children.
//
// child: The node object to add to the existing children.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/addChild(_:)
func (x XMLDTD) AddChild(child INSXMLNode) {
	objc.Send[objc.ID](x.ID, objc.Sel("addChild:"), child)
}

// Inserts a child node in the receiver’s list of children at a specific
// location in the list.
//
// child: An XML-node object that represents the child to insert.
//
// index: An integer identifying the location in the receiver’s list of children to
// insert `child`. The indices of subsequent children in the list are
// incremented by one.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/insertChild(_:at:)
func (x XMLDTD) InsertChildAtIndex(child INSXMLNode, index uint) {
	objc.Send[objc.ID](x.ID, objc.Sel("insertChild:atIndex:"), child, index)
}

// Inserts an array of child nodes at a specified location in the receiver’s
// list of children.
//
// children: An array of [NSXMLNode] objects to insert as children of the receiver.
//
// index: An integer identifying the location in the list of current children to make
// the insertion. The indices of subsequent children in the list are
// incremented by the number of inserted children.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/insertChildren(_:at:)
func (x XMLDTD) InsertChildrenAtIndex(children []NSXMLNode, index uint) {
	objc.Send[objc.ID](x.ID, objc.Sel("insertChildren:atIndex:"), objectivec.IObjectSliceToNSArray(children), index)
}

// Removes the child node at a particular location in the receiver’s list of
// children.
//
// index: An integer identifying the child node to remove. The indices of subsequent
// children in the list are decremented by one.
//
// # Discussion
// 
// The removed child node is released.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/removeChild(at:)
func (x XMLDTD) RemoveChildAtIndex(index uint) {
	objc.Send[objc.ID](x.ID, objc.Sel("removeChildAtIndex:"), index)
}

// Replaces a child at a particular index with another child.
//
// index: An integer identifying the position of a node in the receiver’s list of
// child nodes.
//
// node: An [NSXMLNode] object to replace the object at `index`.
//
// # Discussion
// 
// The replaced child node is released.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/replaceChild(at:with:)
func (x XMLDTD) ReplaceChildAtIndexWithNode(index uint, node INSXMLNode) {
	objc.Send[objc.ID](x.ID, objc.Sel("replaceChildAtIndex:withNode:"), index, node)
}

// Returns the DTD node representing an element declaration for a specified
// element.
//
// name: A string that is the name of an element.
//
// # Return Value
// 
// An autoreleased [NSXMLDTDNode] object, or `nil` if there is no match.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/elementDeclaration(forName:)
func (x XMLDTD) ElementDeclarationForName(name string) INSXMLDTDNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("elementDeclarationForName:"), objc.String(name))
	return NSXMLDTDNodeFromID(rv)
}

// Returns the DTD node representing an attribute-list declaration for a given
// attribute and its element.
//
// name: A string object identifying the name of an attribute.
//
// elementName: A string object identifying the name of an element.
//
// # Return Value
// 
// An autoreleased [NSXMLDTDNode] object, or `nil` if there is no matching
// attribute-list declaration.
//
// # Discussion
// 
// For example, in the attribute-list declaration:
// 
// “idnum” would correspond to `attrName` and “person” would
// correspond to `elementName`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/attributeDeclaration(forName:elementName:)
func (x XMLDTD) AttributeDeclarationForNameElementName(name string, elementName string) INSXMLDTDNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("attributeDeclarationForName:elementName:"), objc.String(name), objc.String(elementName))
	return NSXMLDTDNodeFromID(rv)
}

// Returns the DTD node representing the entity declaration for a specified
// entity.
//
// name: A string that is the name of an entity.
//
// # Return Value
// 
// An autoreleased [NSXMLDTDNode] object, or `nil` if there is no match.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/entityDeclaration(forName:)
func (x XMLDTD) EntityDeclarationForName(name string) INSXMLDTDNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("entityDeclarationForName:"), objc.String(name))
	return NSXMLDTDNodeFromID(rv)
}

// Returns the DTD node representing the notation declaration identified by
// the specified notation name.
//
// name: A string that is the name of a notation.
//
// # Return Value
// 
// An autoreleased [NSXMLDTDNode] object, or `nil` if there is no match.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/notationDeclaration(forName:)
func (x XMLDTD) NotationDeclarationForName(name string) INSXMLDTDNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("notationDeclarationForName:"), objc.String(name))
	return NSXMLDTDNodeFromID(rv)
}

// Returns a DTD node representing the predefined entity declaration with the
// specified name.
//
// name: A string identifying a predefined entity declaration.
//
// # Return Value
// 
// An autoreleased [NSXMLDTDNode] object, or `nil` if there is no match for
// `name`.
//
// # Discussion
// 
// The five predefined entity references (or character references) are “”
// (greater-than sign), “&” (ampersand), “"” (quotation mark), and
// “'” (apostrophe).
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/predefinedEntityDeclaration(forName:)
func (_XMLDTDClass XMLDTDClass) PredefinedEntityDeclarationForName(name string) XMLDTDNode {
	rv := objc.Send[objc.ID](objc.ID(_XMLDTDClass.class), objc.Sel("predefinedEntityDeclarationForName:"), objc.String(name))
	return NSXMLDTDNodeFromID(rv)
}

// Returns the receiver’s public identifier.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/publicID
func (x XMLDTD) PublicID() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("publicID"))
	return NSStringFromID(rv).String()
}
func (x XMLDTD) SetPublicID(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setPublicID:"), objc.String(value))
}

// Returns the receiver’s system identifier.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDTD/systemID
func (x XMLDTD) SystemID() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("systemID"))
	return NSStringFromID(rv).String()
}
func (x XMLDTD) SetSystemID(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setSystemID:"), objc.String(value))
}

// Returns an
//
// See: https://developer.apple.com/documentation/foundation/xmldocument/dtd
func (x XMLDTD) Dtd() INSXMLDTD {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("DTD"))
	return NSXMLDTDFromID(objc.ID(rv))
}
func (x XMLDTD) SetDtd(value INSXMLDTD) {
	objc.Send[struct{}](x.ID, objc.Sel("setDTD:"), value)
}

