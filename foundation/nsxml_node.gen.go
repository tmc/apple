// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XMLNode] class.
var (
	_XMLNodeClass     XMLNodeClass
	_XMLNodeClassOnce sync.Once
)

func getXMLNodeClass() XMLNodeClass {
	_XMLNodeClassOnce.Do(func() {
		_XMLNodeClass = XMLNodeClass{class: objc.GetClass("NSXMLNode")}
	})
	return _XMLNodeClass
}

// GetXMLNodeClass returns the class object for NSXMLNode.
func GetXMLNodeClass() XMLNodeClass {
	return getXMLNodeClass()
}

type XMLNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (xc XMLNodeClass) Alloc() XMLNode {
	rv := objc.Send[XMLNode](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}







// The nodes in the abstract, logical tree structure that represents an XML
// document.
//
// # Overview
// 
// Node objects can be of different kinds, corresponding to the following
// markup constructs in an XML document: element, attribute, text, processing
// instruction, namespace, and comment. In addition, a document-node object
// (specifically, an instance of [NSXMLDocument]) represents an XML document
// in its entirety. [NSXMLNode] objects can also represent document type
// declarations as well as declarations in Document Type Definitions (DTDs).
// Class factory methods of [NSXMLNode] enable you to create nodes of each
// kind. Only document, element, and DTD nodes may have child nodes.
// 
// Among the XML family of classes (excluding [NSXMLParser]) the [NSXMLNode]
// class is the base class. Inheriting from it are the classes [NSXMLElement],
// [NSXMLDocument], [NSXMLDTD], and [NSXMLDTDNode]. [NSXMLNode] specifies the
// interface common to all XML node objects and defines common node behavior
// and attributes, for example hierarchy level, node name and value, tree
// traversal, and the ability to emit representative XML markup text.
// 
// # Subclassing Notes
// 
// You can subclass [NSXMLNode] if you want nodes of kinds different from the
// supported ones, You can also create a subclass with more specialized
// attributes or behavior than [NSXMLNode].
// 
// # Methods to Override
// 
// To subclass [NSXMLNode] you need to override the primary initializer,
// [InitWithKindOptions], and the methods listed below. In most cases, you
// need only invoke the superclass implementation, adding any
// subclass-specific code before or after the invocation, as necessary.
// 
// [Table data omitted]
// 
// By default [NSXMLNode] implements the [NSObject] [isEqual(_:)] method to
// perform a deep comparison: two [NSXMLNode] objects are not considered equal
// unless they have the same name, same child nodes, same attributes, and so
// on. The comparison looks at the node and its children, but does not include
// the node’s parent. If you want a different standard of comparison,
// override ``.
// 
// # Special Considerations
// 
// Because of the architecture and data model of NSXML, when it parses and
// processes a source of XML it cannot know about your subclass unless you
// override the [NSXMLDocument] class method [ReplacementClassForClass] to
// return your custom class in place of an NSXML class. If your custom class
// has no direct NSXML counterpart—for example, it is a subclass of
// [NSXMLNode] that represents CDATA sections—then you can walk the tree
// after it has been created and insert the new node where appropriate.
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// # Creating and Initializing Node Objects
//
//   - [XMLNode.InitWithKind]: Returns an [NSXMLNode] instance initialized with the constant indicating node kind.
//   - [XMLNode.InitWithKindOptions]: Returns an [NSXMLNode] instance initialized with the constant indicating node kind and one or more initialization options.
//
// # Managing XML Node Objects
//
//   - [XMLNode.Index]: Returns the index of the receiver identifying its position relative to its sibling nodes.
//   - [XMLNode.SetIndex]
//   - [XMLNode.Kind]: Returns the kind of node the receiver is as a constant of type [XMLNode.Kind](<doc://com.apple.foundation/documentation/Foundation/XMLNode/Kind-swift.enum>).
//   - [XMLNode.Level]: Returns the nesting level of the receiver within the tree hierarchy.
//   - [XMLNode.Name]: Returns the name of the receiver.
//   - [XMLNode.SetName]
//   - [XMLNode.ObjectValue]: Returns the object value of the receiver.
//   - [XMLNode.SetObjectValue]
//   - [XMLNode.StringValue]: Returns the content of the receiver as a string value.
//   - [XMLNode.SetStringValue]
//   - [XMLNode.SetStringValueResolvingEntities]: Sets the content of the receiver as a string value and, optionally, resolves character references, predefined entities, and user-defined entities as declared in the associated DTD.
//   - [XMLNode.URI]: Returns the URI associated with the receiver.
//   - [XMLNode.SetURI]
//
// # Navigating the Tree of Nodes
//
//   - [XMLNode.RootDocument]: Returns the [XMLDocument](<doc://com.apple.foundation/documentation/Foundation/XMLDocument>) object containing the root element and representing the XML document as a whole.
//   - [XMLNode.Parent]: Returns the parent node of the receiver.
//   - [XMLNode.ChildAtIndex]: Returns the child node of the receiver at the specified location.
//   - [XMLNode.ChildCount]: Returns the number of child nodes the receiver has.
//   - [XMLNode.Children]: Returns an immutable array containing the child nodes of the receiver (as [NSXMLNode] objects).
//   - [XMLNode.NextNode]: Returns the next [NSXMLNode] object in document order.
//   - [XMLNode.NextSibling]: Returns the next [NSXMLNode] object that is a sibling node to the receiver.
//   - [XMLNode.PreviousNode]: Returns the previous [NSXMLNode] object in document order.
//   - [XMLNode.PreviousSibling]: Returns the previous [NSXMLNode] object that is a sibling node to the receiver.
//   - [XMLNode.Detach]: Detaches the receiver from its parent node.
//
// # Emitting Node Content
//
//   - [XMLNode.XMLString]: Returns the string representation of the receiver as it would appear in an XML document.
//   - [XMLNode.XMLStringWithOptions]: Returns the string representation of the receiver as it would appear in an XML document, with one or more output options specified.
//   - [XMLNode.CanonicalXMLStringPreservingComments]: Returns a string object encapsulating the receiver’s XML in canonical form.
//   - [XMLNode.Description]
//
// # Executing Queries
//
//   - [XMLNode.NodesForXPathError]: Returns the nodes resulting from executing an XPath query upon the receiver.
//   - [XMLNode.ObjectsForXQueryError]: Returns the objects resulting from executing an XQuery query upon the receiver.
//   - [XMLNode.ObjectsForXQueryConstantsError]: Returns the objects resulting from executing an XQuery query upon the receiver.
//   - [XMLNode.XPath]: Returns the XPath expression identifying the receiver’s location in the document tree.
//
// # Managing Namespaces
//
//   - [XMLNode.LocalName]: Returns the local name of the receiver.
//   - [XMLNode.Prefix]: Returns the prefix of the receiver’s name.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode
type XMLNode struct {
	objectivec.Object
}

// XMLNodeFromID constructs a [XMLNode] from an objc.ID.
//
// The nodes in the abstract, logical tree structure that represents an XML
// document.
func XMLNodeFromID(id objc.ID) XMLNode {
	return NSXMLNode{objectivec.Object{id}}
}

// NSXMLNodeFromID is an alias for [XMLNodeFromID] for cross-framework compatibility.
func NSXMLNodeFromID(id objc.ID) XMLNode { return XMLNodeFromID(id) }
// NOTE: XMLNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [XMLNode] class.
//
// # Creating and Initializing Node Objects
//
//   - [IXMLNode.InitWithKind]: Returns an [NSXMLNode] instance initialized with the constant indicating node kind.
//   - [IXMLNode.InitWithKindOptions]: Returns an [NSXMLNode] instance initialized with the constant indicating node kind and one or more initialization options.
//
// # Managing XML Node Objects
//
//   - [IXMLNode.Index]: Returns the index of the receiver identifying its position relative to its sibling nodes.
//   - [IXMLNode.SetIndex]
//   - [IXMLNode.Kind]: Returns the kind of node the receiver is as a constant of type [XMLNode.Kind](<doc://com.apple.foundation/documentation/Foundation/XMLNode/Kind-swift.enum>).
//   - [IXMLNode.Level]: Returns the nesting level of the receiver within the tree hierarchy.
//   - [IXMLNode.Name]: Returns the name of the receiver.
//   - [IXMLNode.SetName]
//   - [IXMLNode.ObjectValue]: Returns the object value of the receiver.
//   - [IXMLNode.SetObjectValue]
//   - [IXMLNode.StringValue]: Returns the content of the receiver as a string value.
//   - [IXMLNode.SetStringValue]
//   - [IXMLNode.SetStringValueResolvingEntities]: Sets the content of the receiver as a string value and, optionally, resolves character references, predefined entities, and user-defined entities as declared in the associated DTD.
//   - [IXMLNode.URI]: Returns the URI associated with the receiver.
//   - [IXMLNode.SetURI]
//
// # Navigating the Tree of Nodes
//
//   - [IXMLNode.RootDocument]: Returns the [XMLDocument](<doc://com.apple.foundation/documentation/Foundation/XMLDocument>) object containing the root element and representing the XML document as a whole.
//   - [IXMLNode.Parent]: Returns the parent node of the receiver.
//   - [IXMLNode.ChildAtIndex]: Returns the child node of the receiver at the specified location.
//   - [IXMLNode.ChildCount]: Returns the number of child nodes the receiver has.
//   - [IXMLNode.Children]: Returns an immutable array containing the child nodes of the receiver (as [NSXMLNode] objects).
//   - [IXMLNode.NextNode]: Returns the next [NSXMLNode] object in document order.
//   - [IXMLNode.NextSibling]: Returns the next [NSXMLNode] object that is a sibling node to the receiver.
//   - [IXMLNode.PreviousNode]: Returns the previous [NSXMLNode] object in document order.
//   - [IXMLNode.PreviousSibling]: Returns the previous [NSXMLNode] object that is a sibling node to the receiver.
//   - [IXMLNode.Detach]: Detaches the receiver from its parent node.
//
// # Emitting Node Content
//
//   - [IXMLNode.XMLString]: Returns the string representation of the receiver as it would appear in an XML document.
//   - [IXMLNode.XMLStringWithOptions]: Returns the string representation of the receiver as it would appear in an XML document, with one or more output options specified.
//   - [IXMLNode.CanonicalXMLStringPreservingComments]: Returns a string object encapsulating the receiver’s XML in canonical form.
//   - [IXMLNode.Description]
//
// # Executing Queries
//
//   - [IXMLNode.NodesForXPathError]: Returns the nodes resulting from executing an XPath query upon the receiver.
//   - [IXMLNode.ObjectsForXQueryError]: Returns the objects resulting from executing an XQuery query upon the receiver.
//   - [IXMLNode.ObjectsForXQueryConstantsError]: Returns the objects resulting from executing an XQuery query upon the receiver.
//   - [IXMLNode.XPath]: Returns the XPath expression identifying the receiver’s location in the document tree.
//
// # Managing Namespaces
//
//   - [IXMLNode.LocalName]: Returns the local name of the receiver.
//   - [IXMLNode.Prefix]: Returns the prefix of the receiver’s name.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode
type IXMLNode interface {
	objectivec.IObject
	NSCopying

	// Topic: Creating and Initializing Node Objects

	// Returns an [NSXMLNode] instance initialized with the constant indicating node kind.
	InitWithKind(kind NSXMLNodeKind) XMLNode
	// Returns an [NSXMLNode] instance initialized with the constant indicating node kind and one or more initialization options.
	InitWithKindOptions(kind NSXMLNodeKind, options NSXMLNodeOptions) XMLNode

	// Topic: Managing XML Node Objects

	// Returns the index of the receiver identifying its position relative to its sibling nodes.
	Index() int
	SetIndex(value int)
	// Returns the kind of node the receiver is as a constant of type [XMLNode.Kind](<doc://com.apple.foundation/documentation/Foundation/XMLNode/Kind-swift.enum>).
	Kind() NSXMLNodeKind
	// Returns the nesting level of the receiver within the tree hierarchy.
	Level() uint
	// Returns the name of the receiver.
	Name() string
	SetName(value string)
	// Returns the object value of the receiver.
	ObjectValue() objectivec.IObject
	SetObjectValue(value objectivec.IObject)
	// Returns the content of the receiver as a string value.
	StringValue() string
	SetStringValue(value string)
	// Sets the content of the receiver as a string value and, optionally, resolves character references, predefined entities, and user-defined entities as declared in the associated DTD.
	SetStringValueResolvingEntities(string_ string, resolve bool)
	// Returns the URI associated with the receiver.
	URI() string
	SetURI(value string)

	// Topic: Navigating the Tree of Nodes

	// Returns the [XMLDocument](<doc://com.apple.foundation/documentation/Foundation/XMLDocument>) object containing the root element and representing the XML document as a whole.
	RootDocument() INSXMLDocument
	// Returns the parent node of the receiver.
	Parent() INSXMLNode
	// Returns the child node of the receiver at the specified location.
	ChildAtIndex(index uint) INSXMLNode
	// Returns the number of child nodes the receiver has.
	ChildCount() uint
	// Returns an immutable array containing the child nodes of the receiver (as [NSXMLNode] objects).
	Children() []NSXMLNode
	// Returns the next [NSXMLNode] object in document order.
	NextNode() INSXMLNode
	// Returns the next [NSXMLNode] object that is a sibling node to the receiver.
	NextSibling() INSXMLNode
	// Returns the previous [NSXMLNode] object in document order.
	PreviousNode() INSXMLNode
	// Returns the previous [NSXMLNode] object that is a sibling node to the receiver.
	PreviousSibling() INSXMLNode
	// Detaches the receiver from its parent node.
	Detach()

	// Topic: Emitting Node Content

	// Returns the string representation of the receiver as it would appear in an XML document.
	XMLString() string
	// Returns the string representation of the receiver as it would appear in an XML document, with one or more output options specified.
	XMLStringWithOptions(options NSXMLNodeOptions) string
	// Returns a string object encapsulating the receiver’s XML in canonical form.
	CanonicalXMLStringPreservingComments(comments bool) string
	Description() string

	// Topic: Executing Queries

	// Returns the nodes resulting from executing an XPath query upon the receiver.
	NodesForXPathError(xpath string) ([]NSXMLNode, error)
	// Returns the objects resulting from executing an XQuery query upon the receiver.
	ObjectsForXQueryError(xquery string) (INSArray, error)
	// Returns the objects resulting from executing an XQuery query upon the receiver.
	ObjectsForXQueryConstantsError(xquery string, constants INSDictionary) (INSArray, error)
	// Returns the XPath expression identifying the receiver’s location in the document tree.
	XPath() string

	// Topic: Managing Namespaces

	// Returns the local name of the receiver.
	LocalName() string
	// Returns the prefix of the receiver’s name.
	Prefix() string
}





// Init initializes the instance.
func (x XMLNode) Init() XMLNode {
	rv := objc.Send[XMLNode](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XMLNode) Autorelease() XMLNode {
	rv := objc.Send[XMLNode](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLNode creates a new XMLNode instance.
func NewXMLNode() XMLNode {
	class := getXMLNodeClass()
	rv := objc.Send[XMLNode](objc.ID(class.class), objc.Sel("new"))
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
func NewXMLNodeWithKind(kind NSXMLNodeKind) XMLNode {
	instance := getXMLNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:"), kind)
	return XMLNodeFromID(rv)
}


// Returns an [NSXMLNode] instance initialized with the constant indicating
// node kind and one or more initialization options.
//
// kind: An `enum` constant of type [XMLNode.Kind] that indicates the type of node.
// See Constants for a list of valid NSXMLNodeKind constants.
// //
// [XMLNode.Kind]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum
//
// options: One or more constants that specify initialization options; if there are
// multiple constants, bit-OR them together. These options request operations
// on the represented XML related to fidelity (for example, preserving
// entities), quoting style, handling of empty elements, and other things. See
// Constants for a list of valid node-initialization constants.
//
// # Return Value
// 
// An [NSXMLNode] object initialized with the given kind and options, or `nil`
// if the object couldn’t be created. If `kind` is not a valid NSXMLNodeKind
// constant, the method returns an [NSXMLNode] object of kind
// [NSXMLInvalidKind].
//
// # Discussion
// 
// Do not use this initializer for creating instances of [NSXMLDTDNode] for
// attribute-list declarations. Instead, use the [DTDNodeWithXMLString] class
// method of this class or the [InitWithXMLString] method of the
// [NSXMLDTDNode] class.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/init(kind:options:)
func NewXMLNodeWithKindOptions(kind NSXMLNodeKind, options NSXMLNodeOptions) XMLNode {
	instance := getXMLNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:options:"), kind, options)
	return XMLNodeFromID(rv)
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
func (x XMLNode) InitWithKind(kind NSXMLNodeKind) XMLNode {
	rv := objc.Send[XMLNode](x.ID, objc.Sel("initWithKind:"), kind)
	return rv
}

// Returns an [NSXMLNode] instance initialized with the constant indicating
// node kind and one or more initialization options.
//
// kind: An `enum` constant of type [XMLNode.Kind] that indicates the type of node.
// See Constants for a list of valid NSXMLNodeKind constants.
// //
// [XMLNode.Kind]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum
//
// options: One or more constants that specify initialization options; if there are
// multiple constants, bit-OR them together. These options request operations
// on the represented XML related to fidelity (for example, preserving
// entities), quoting style, handling of empty elements, and other things. See
// Constants for a list of valid node-initialization constants.
//
// # Return Value
// 
// An [NSXMLNode] object initialized with the given kind and options, or `nil`
// if the object couldn’t be created. If `kind` is not a valid NSXMLNodeKind
// constant, the method returns an [NSXMLNode] object of kind
// [NSXMLInvalidKind].
//
// # Discussion
// 
// Do not use this initializer for creating instances of [NSXMLDTDNode] for
// attribute-list declarations. Instead, use the [DTDNodeWithXMLString] class
// method of this class or the [InitWithXMLString] method of the
// [NSXMLDTDNode] class.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/init(kind:options:)
func (x XMLNode) InitWithKindOptions(kind NSXMLNodeKind, options NSXMLNodeOptions) XMLNode {
	rv := objc.Send[XMLNode](x.ID, objc.Sel("initWithKind:options:"), kind, options)
	return rv
}

// Sets the content of the receiver as a string value and, optionally,
// resolves character references, predefined entities, and user-defined
// entities as declared in the associated DTD.
//
// string: A string to assign as the value of the receiver.
//
// resolve: [true] to resolve character references, predefined entities, and
// user-defined entities as declared in the associated DTD; [false] otherwise.
// Namespace and processing-instruction nodes have their entities resolved
// even if `resolve` is [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// User-defined entities not declared in the DTD remain in their unresolved
// form. This method can only be invoked on [NSXMLNode] objects that may have
// content, specifically elements, attributes, namespaces, processing
// instructions, text, and DTD-declaration nodes. Setting the string value of
// a node object removes all existing children, including processing
// instructions and comments. Setting the string value of an element -node
// object creates a text node as the sole child.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/setStringValue(_:resolvingEntities:)
func (x XMLNode) SetStringValueResolvingEntities(string_ string, resolve bool) {
	objc.Send[objc.ID](x.ID, objc.Sel("setStringValue:resolvingEntities:"), objc.String(string_), resolve)
}

// Returns the child node of the receiver at the specified location.
//
// index: An integer specifying a node position in the receiver’s array of
// children. If `index` is out of bounds, an exception is raised.
//
// # Return Value
// 
// An NSXMLNode object or `nil` f the receiver has no children.
//
// # Discussion
// 
// The receiver should be an [NSXMLNode] object representing a document,
// element, or document type declaration. The returned node object can
// represent an element, comment, text, or processing instruction.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/child(at:)
func (x XMLNode) ChildAtIndex(index uint) INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("childAtIndex:"), index)
	return NSXMLNodeFromID(rv)
}

// Detaches the receiver from its parent node.
//
// # Discussion
// 
// This method is applicable to [NSXMLNode] objects representing elements,
// text, comments, processing instructions, attributes, and namespaces. Once
// the node object is detached, you can add it as a child node of another
// parent.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/detach()
func (x XMLNode) Detach() {
	objc.Send[objc.ID](x.ID, objc.Sel("detach"))
}

// Returns the string representation of the receiver as it would appear in an
// XML document, with one or more output options specified.
//
// options: One or more `enum` constants identifying an output option; bit-OR multiple
// constants together. See Constants for a list of valid constants for
// specifying output options.
//
// # Discussion
// 
// The returned string includes the string representations of all children.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/xmlString(options:)
func (x XMLNode) XMLStringWithOptions(options NSXMLNodeOptions) string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("XMLStringWithOptions:"), options)
	return NSStringFromID(rv).String()
}

// Returns a string object encapsulating the receiver’s XML in canonical
// form.
//
// comments: [true] to preserve comments, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Be sure to set the input option [nodePreserveWhitespace] for true canonical
// form. The canonical form of an XML document is defined by the World Wide
// Web Consortium at `//www.W3XCUIElementTypeOrg()/TR/xml-c14n`. Generally, if
// two documents with varying physical representations have the same canonical
// form, then they are considered logically equivalent within the given
// application context. The following list summarizes most key aspects of
// canonical form as defined by the W3C recommendation:
// 
// - Encodes the document in UTF-8. - Normalizes line breaks to “#xA” on
// input before parsing. - Normalizes attribute values in the manner of a
// validating processor. - Replaces character and parsed entity references
// with their character content. - Replaces CDATA sections with their
// character content. - Removes the XML declaration and the document type
// declaration (DTD). - Converts empty elements to start-end tag pairs. -
// Normalizes whitespace outside of the document element and within start and
// end tags. - Retains all whitespace characters in content (excluding
// characters removed during line-feed normalization). - Sets attribute value
// delimiters to quotation marks (double quotes). - Replaces special
// characters in attribute values and character content with character
// references. - Removes superfluous namespace declarations from each element.
// - Adds default attributes to each element. - Imposes lexicographic order on
// the namespace declarations and attributes of each element.
//
// [nodePreserveWhitespace]: https://developer.apple.com/documentation/Foundation/XMLNode/Options/nodePreserveWhitespace
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/canonicalXMLStringPreservingComments(_:)
func (x XMLNode) CanonicalXMLStringPreservingComments(comments bool) string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("canonicalXMLStringPreservingComments:"), comments)
	return NSStringFromID(rv).String()
}

// Returns the nodes resulting from executing an XPath query upon the
// receiver.
//
// xpath: A string that expresses an XPath query.
//
// # Return Value
// 
// An array of [NSXMLNode] objects that match the query, or an empty array if
// there are no matches.
//
// # Discussion
// 
// The receiver acts as the context item for the query (”.”). If you have
// explicitly added adjacent text nodes as children of an element, you should
// invoke the [NSXMLElement] method
// [NormalizeAdjacentTextNodesPreservingCDATA] (with an argument of [false])
// on the element before applying any XPath queries to it; this method
// coalesces these text nodes. The same precaution applies if you have
// processed a document preserving CDATA sections and these sections are
// adjacent to text nodes.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/nodes(forXPath:)
func (x XMLNode) NodesForXPathError(xpath string) ([]NSXMLNode, error) {
			var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](x.ID, objc.Sel("nodesForXPath:error:"), objc.String(xpath), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) XMLNode {
		return XMLNodeFromID(id)
	}), nil

}

// Returns the objects resulting from executing an XQuery query upon the
// receiver.
//
// xquery: A string that expresses an XQuery query.
//
// # Discussion
// 
// The receiver acts as the context item for the query (”.”). If the
// receiver has been changed after parsing to have multiple adjacent text
// nodes, you should invoke the [NSXMLElement] method
// [NormalizeAdjacentTextNodesPreservingCDATA] (with an argument of [false])
// to coalesce the text nodes before querying .This convenience method invokes
// [ObjectsForXQueryConstantsError] with `nil` for the `constants` dictionary.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/objects(forXQuery:)
func (x XMLNode) ObjectsForXQueryError(xquery string) (INSArray, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("objectsForXQuery:error:"), objc.String(xquery), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSArrayFromID(rv), nil

}

// Returns the objects resulting from executing an XQuery query upon the
// receiver.
//
// xquery: A string that expresses an XQuery query.
//
// constants: A dictionary containing externally declared constants where the name of
// each constant variable is a key.
//
// # Discussion
// 
// The receiver acts as the context item for the query (”.”). If the
// receiver has been changed after parsing to have multiple adjacent text
// nodes, you should invoke the [NSXMLElement] method
// [NormalizeAdjacentTextNodesPreservingCDATA] (with an argument of [false])
// to coalesce the text nodes before querying.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/objects(forXQuery:constants:)
func (x XMLNode) ObjectsForXQueryConstantsError(xquery string, constants INSDictionary) (INSArray, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("objectsForXQuery:constants:error:"), objc.String(xquery), constants, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return NSArrayFromID(rv), nil

}





// Returns an empty document node.
//
// # Return Value
// 
// An empty document node—that is, an [NSXMLDocument] instance without a
// root element or XML-declaration information (version, encoding, standalone
// flag). Returns `nil` if the object couldn’t be created.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/document()
func (_XMLNodeClass XMLNodeClass) Document() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("document"))
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLDocument] object initialized with a given root element.
//
// element: An [NSXMLElement] object representing an element.
//
// # Return Value
// 
// An [NSXMLDocument] object initialized with the root element `element` or
// `nil` if the object couldn’t be created.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/document(withRootElement:)
func (_XMLNodeClass XMLNodeClass) DocumentWithRootElement(element INSXMLElement) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("documentWithRootElement:"), element)
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLElement] object with a given tag identifier, or name
//
// name: A string that is the name (or tag identifier) of an element.
//
// # Return Value
// 
// An [NSXMLElement] object or `nil` if the object couldn’t be created.
//
// # Discussion
// 
// The equivalent XML markup is ``.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:)
func (_XMLNodeClass XMLNodeClass) ElementWithName(name string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("elementWithName:"), objc.String(name))
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLElement] object with the given tag (name), attributes, and
// children.
//
// name: A string that is the name (tag identifier) of the element.
//
// children: An array of [NSXMLElement] objects or [NSXMLNode] objects of kinds
// [XMLNode.Kind.element], [XMLNode.Kind.processingInstruction],
// [XMLNode.Kind.comment], and [XMLNode.Kind.text]. Specify `nil` if there are
// no children to add to this node object.
// //
// [XMLNode.Kind.comment]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/comment
// [XMLNode.Kind.element]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/element
// [XMLNode.Kind.processingInstruction]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/processingInstruction
// [XMLNode.Kind.text]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/text
//
// attributes: An array of [NSXMLNode] objects of kind [XMLNode.Kind.attribute]. Specify
// `nil` if there are no attributes to add to this node object.
// //
// [XMLNode.Kind.attribute]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/attribute
//
// # Return Value
// 
// An [NSXMLElement] object or `nil` if the object couldn’t be created.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:children:attributes:)
func (_XMLNodeClass XMLNodeClass) ElementWithNameChildrenAttributes(name string, children []NSXMLNode, attributes []NSXMLNode) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("elementWithName:children:attributes:"), objc.String(name), objectivec.IObjectSliceToNSArray(children), objectivec.IObjectSliceToNSArray(attributes))
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLElement] object with a single text-node child containing
// the specified text.
//
// name: A string that is the name (tag identifier) of the element.
//
// string: A string that is the content of the receiver’s text node.
//
// # Return Value
// 
// An [NSXMLElement] object with a single text-node child—an [NSXMLNode]
// object of kind [XMLNode.Kind.text]—containing the text specified in
// `string`. Returns `nil` if the object couldn’t be created.
//
// [XMLNode.Kind.text]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/text
//
// # Discussion
// 
// The equivalent XML markup is ```string```.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:stringValue:)
func (_XMLNodeClass XMLNodeClass) ElementWithNameStringValue(name string, string_ string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("elementWithName:stringValue:"), objc.String(name), objc.String(string_))
	return objectivec.Object{ID: rv}
}

// Returns an element whose fully qualified name is specified.
//
// name: A string that is the name (or tag identifier) of an element.
//
// URI: A URI (Universal Resource Identifier) that qualifies `name`.
//
// # Return Value
// 
// An [NSXMLElement] object or `nil` if the object cannot be created.
//
// # Discussion
// 
// The equivalent XML markup is ````.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/element(withName:uri:)
func (_XMLNodeClass XMLNodeClass) ElementWithNameURI(name string, URI string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("elementWithName:URI:"), objc.String(name), objc.String(URI))
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLNode] object representing an attribute node with a given
// name and string.
//
// name: A string that is the name of an attribute.
//
// stringValue: A string that is the value of an attribute.
//
// # Return Value
// 
// An [NSXMLNode] object of kind [XMLNode.Kind.attribute] or `nil` if the
// object couldn’t be created.
//
// [XMLNode.Kind.attribute]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/attribute
//
// # Discussion
// 
// For example, in the attribute “id=`12345’”, “id” is the attribute
// name and “12345” is the attribute value.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/attribute(withName:stringValue:)
func (_XMLNodeClass XMLNodeClass) AttributeWithNameStringValue(name string, stringValue string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("attributeWithName:stringValue:"), objc.String(name), objc.String(stringValue))
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLNode] object representing an attribute node with a given
// qualified name and string.
//
// name: A string that is the name of an attribute.
//
// URI: A URI (Universal Resource Identifier) that qualifies `name`.
//
// stringValue: A string that is the value of the attribute.
//
// # Return Value
// 
// An [NSXMLNode] object of kind [XMLNode.Kind.attribute] or `nil` if the
// object couldn’t be created.
//
// [XMLNode.Kind.attribute]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/attribute
//
// # Discussion
// 
// For example, in the attribute “bst:id=`12345’”, “bst” is the name
// qualifier (derived from the URI), “id” is the attribute name, and
// “12345” is the attribute value.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/attribute(withName:uri:stringValue:)
func (_XMLNodeClass XMLNodeClass) AttributeWithNameURIStringValue(name string, URI string, stringValue string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("attributeWithName:URI:stringValue:"), objc.String(name), objc.String(URI), objc.String(stringValue))
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLNode] object representing a text node with specified
// content.
//
// stringValue: A string that is the textual content of the node.
//
// # Return Value
// 
// An [NSXMLNode] object of kind [XMLNode.Kind.text] initialized with the
// textual `value` or `nil` if the object couldn’t be created.
//
// [XMLNode.Kind.text]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/text
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/text(withStringValue:)
func (_XMLNodeClass XMLNodeClass) TextWithStringValue(stringValue string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("textWithStringValue:"), objc.String(stringValue))
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLNode] object representing a comment node containing given
// text.
//
// stringValue: A string specifying the text of the comment. You may specify `nil` or an
// empty string (see Return Value).
//
// # Return Value
// 
// An [NSXMLNode] object representing an comment node ([XMLNode.Kind.comment])
// containing the text `stringValue` or `nil` if the object couldn’t be
// created. If `stringValue` is `nil` or an empty string, a content-less
// comment node is returned (``).
//
// [XMLNode.Kind.comment]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/comment
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/comment(withStringValue:)
func (_XMLNodeClass XMLNodeClass) CommentWithStringValue(stringValue string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("commentWithStringValue:"), objc.String(stringValue))
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLNode] object representing a namespace with a specified
// name and URI.
//
// name: A string that is the name of the namespace. Specify an empty string for
// `name` to get the default namespace.
//
// stringValue: A string that identifies the URI associated with the namespace.
//
// # Return Value
// 
// An [NSXMLNode] object of kind [XMLNode.Kind.namespace] or `nil` if the
// object couldn’t be created.
//
// [XMLNode.Kind.namespace]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/namespace
//
// # Discussion
// 
// The equivalent namespace declaration in XML markup is
// ```name``="``value``"`.
// 
// # Special Considerations
// 
// Applications linked on macOS 10.6 or later will throw an exception if the
// `name` parameter is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/namespace(withName:stringValue:)
func (_XMLNodeClass XMLNodeClass) NamespaceWithNameStringValue(name string, stringValue string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("namespaceWithName:stringValue:"), objc.String(name), objc.String(stringValue))
	return objectivec.Object{ID: rv}
}

// Returns a [NSXMLDTDNode] object representing the DTD declaration for an
// element, attribute, entity, or notation based on a given string.
//
// string: A string that is a DTD declaration. The receiver parses this string to
// determine the kind of DTD node to create.
//
// # Return Value
// 
// An [NSXMLDTDNode] object representing the DTD declaration or `nil` if the
// object couldn’t be created.
//
// # Discussion
// 
// For example, if `string` is the following:
// 
// [NSXMLNode] is able to assign the created node object a kind of
// [XMLNode.Kind.entityDeclaration] by parsing “ENTITY”.
// 
// Note that if an attribute-list declaration (`` )has multiple attributes
// [NSXMLNode] only creates an [NSXMLDTDNode] object for the last attribute in
// the declaration.
//
// [XMLNode.Kind.entityDeclaration]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/entityDeclaration
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/dtdNode(withXMLString:)
func (_XMLNodeClass XMLNodeClass) DTDNodeWithXMLString(string_ string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("DTDNodeWithXMLString:"), objc.String(string_))
	return objectivec.Object{ID: rv}
}

// Returns an [NSXMLNode] object representing one of the predefined namespaces
// with the specified prefix.
//
// name: A string specifying a prefix for a predefined namespace, for example
// “xml”, “xs”, or “xsi”.
//
// # Return Value
// 
// An [NSXMLNode] object of kind [XMLNode.Kind.namespace] or `nil` if the
// object couldn’t be created. If something other than a
// predefined-namespace prefix is specified, the method returns `nil`.
//
// [XMLNode.Kind.namespace]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/namespace
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/predefinedNamespace(forPrefix:)
func (_XMLNodeClass XMLNodeClass) PredefinedNamespaceForPrefix(name string) XMLNode {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("predefinedNamespaceForPrefix:"), objc.String(name))
	return NSXMLNodeFromID(rv)
}

// Returns an [NSXMLNode] object representing a processing instruction with a
// specified name and value.
//
// name: A string that is the name of the processing instruction.
//
// stringValue: A string that is the value of the processing instruction.
//
// # Return Value
// 
// An [NSXMLNode] object of kind [XMLNode.Kind.processingInstruction] or `nil`
// if the object couldn’t be created.
//
// [XMLNode.Kind.processingInstruction]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/processingInstruction
//
// # Discussion
// 
// The equivalent XML markup is ``.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/processingInstruction(withName:stringValue:)
func (_XMLNodeClass XMLNodeClass) ProcessingInstructionWithNameStringValue(name string, stringValue string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("processingInstructionWithName:stringValue:"), objc.String(name), objc.String(stringValue))
	return objectivec.Object{ID: rv}
}

// Returns the local name from the specified qualified name.
//
// name: A string that is a qualified name.
//
// # Discussion
// 
// For example, if the qualified name is “bst:title”, this method returns
// “title”.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/localName(forName:)
func (_XMLNodeClass XMLNodeClass) LocalNameForName(name string) string {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("localNameForName:"), objc.String(name))
	return NSStringFromID(rv).String()
}

// Returns the prefix from the specified qualified name.
//
// name: A string that is a qualified name.
//
// # Discussion
// 
// For example, if the qualified name is “bst:title”, this method returns
// “bst”.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/prefix(forName:)
func (_XMLNodeClass XMLNodeClass) PrefixForName(name string) string {
	rv := objc.Send[objc.ID](objc.ID(_XMLNodeClass.class), objc.Sel("prefixForName:"), objc.String(name))
	return NSStringFromID(rv).String()
}








// Returns the index of the receiver identifying its position relative to its
// sibling nodes.
//
// See: https://developer.apple.com/documentation/foundation/xmlnode/index
func (x XMLNode) Index() int {
	rv := objc.Send[int](x.ID, objc.Sel("index"))
	return rv
}
func (x XMLNode) SetIndex(value int) {
	objc.Send[struct{}](x.ID, objc.Sel("setIndex:"), value)
}



// Returns the kind of node the receiver is as a constant of type
// [XMLNode.Kind].
//
// [XMLNode.Kind]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum
//
// # Discussion
// 
// [NSXMLNode] objects can represent documents, elements, attributes,
// namespaces, text, processing instructions, comments, document type
// declarations, and specific declarations within DTDs. See Constants for a
// list of valid NSXMLNodeKind constants
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/kind-swift.property
func (x XMLNode) Kind() NSXMLNodeKind {
	rv := objc.Send[NSXMLNodeKind](x.ID, objc.Sel("kind"))
	return NSXMLNodeKind(rv)
}



// Returns the nesting level of the receiver within the tree hierarchy.
//
// # Return Value
// 
// An integer indicating a nesting level.
// 
// # Discussion
// 
// The root element of a document has a nesting level of one.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/level
func (x XMLNode) Level() uint {
	rv := objc.Send[uint](x.ID, objc.Sel("level"))
	return rv
}



// Returns the name of the receiver.
//
// # Return Value
// 
// Returns a string specifying the name of the receiver. May return `nil` if
// the receiver is not a valid kind of node (see discussion).
// 
// # Discussion
// 
// This method is applicable only to [NSXMLNode] objects representing
// elements, attributes, namespaces, processing instructions, and
// DTD-declaration nodes. If the receiver is not an object of one of these
// kinds, this method returns `nil`. For example, in the following
// construction:
// 
// The returned name for the element is “title”. If the name is associated
// with a namespace, the qualified name is returned. For example, if you
// create an element with local name “foo” and URI “http://bar.com”
// and the namespace “xmlns:baz=‘http://bar.com’” is applied to this
// node, when you invoke this method on the node you get “baz:foo”.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/name
func (x XMLNode) Name() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (x XMLNode) SetName(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setName:"), objc.String(value))
}



// Returns the object value of the receiver.
//
// # Return Value
// 
// The object value of the receiver, which may be the same as the value
// returned by [StringValue]. For nodes without content (for example, document
// nodes), this method returns the string value, or an empty string if there
// is no string value.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/objectValue
func (x XMLNode) ObjectValue() objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("objectValue"))
	return objectivec.Object{ID: rv}
}
func (x XMLNode) SetObjectValue(value objectivec.IObject) {
	objc.Send[struct{}](x.ID, objc.Sel("setObjectValue:"), value)
}



// Returns the content of the receiver as a string value.
//
// # Discussion
// 
// If the receiver is a node object of element kind, the content is that of
// any text-node children. This method recursively visits elements nodes and
// concatenates their text nodes in document order with no intervening spaces.
// If the receiver’s content is set as an object value, this method returns
// the string value representing the object. If the object value is one of the
// standard, built-in ones ([NSNumber], [NSCalendarDate], and so on), the
// string value is in canonical format as defined by the W3C XML Schema Data
// Types specification. If the object value is not represented by one of these
// classes (or if the default value transformer for a class has been
// overridden), the string value is generated by the [NSValueTransformer]
// registered for that object type.
//
// [NSCalendarDate]: https://developer.apple.com/documentation/Foundation/NSCalendarDate
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/stringValue
func (x XMLNode) StringValue() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("stringValue"))
	return NSStringFromID(rv).String()
}
func (x XMLNode) SetStringValue(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setStringValue:"), objc.String(value))
}



// Returns the URI associated with the receiver.
//
// # Discussion
// 
// A node’s URI is derived from its namespace or a document’s URI; for
// documents, the URI comes either from the parsed XML or is explicitly set.
// You cannot change the URI for a particular node other for than a namespace
// or document node.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/uri
func (x XMLNode) URI() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("URI"))
	return NSStringFromID(rv).String()
}
func (x XMLNode) SetURI(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setURI:"), objc.String(value))
}



// Returns the [NSXMLDocument] object containing the root element and
// representing the XML document as a whole.
//
// # Discussion
// 
// If the receiver is a standalone node (that is, a node at the head of a
// detached branch of the tree), this method returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/rootDocument
func (x XMLNode) RootDocument() INSXMLDocument {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("rootDocument"))
	return NSXMLDocumentFromID(objc.ID(rv))
}



// Returns the parent node of the receiver.
//
// # Discussion
// 
// Document nodes and standalone nodes (that is, the root of a detached branch
// of a tree) have no parent, and sending this message to them returns `nil`.
// A one-to-one relationship does not always exists between a parent and its
// children; although a namespace or attribute node cannot be a child, it
// still has a parent element.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/parent
func (x XMLNode) Parent() INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("parent"))
	return NSXMLNodeFromID(objc.ID(rv))
}



// Returns the number of child nodes the receiver has.
//
// # Discussion
// 
// This receiver should be an [NSXMLNode] object representing a document,
// element, or document type declaration. For performance reasons, use this
// method instead of getting the count from the array returned by [Children]
// (for example, `[[thisNode children] count`]).
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/childCount
func (x XMLNode) ChildCount() uint {
	rv := objc.Send[uint](x.ID, objc.Sel("childCount"))
	return rv
}



// Returns an immutable array containing the child nodes of the receiver (as
// [NSXMLNode] objects).
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/children
func (x XMLNode) Children() []NSXMLNode {
	rv := objc.Send[[]objc.ID](x.ID, objc.Sel("children"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSXMLNode {
		return NSXMLNodeFromID(id)
	})
}



// Returns the next [NSXMLNode] object in document order.
//
// # Discussion
// 
// You use this method to “walk” forward through the tree structure
// representing an XML document or document section. (Use [PreviousNode] to
// traverse the tree in the opposite direction.) Document order is the natural
// order that XML constructs appear in markup text. If you send this message
// to the last node in the tree, `nil` is returned. [NSXMLNode] bypasses
// namespace and attribute nodes when it traverses a tree in document order.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/next
func (x XMLNode) NextNode() INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("nextNode"))
	return NSXMLNodeFromID(objc.ID(rv))
}



// Returns the next [NSXMLNode] object that is a sibling node to the receiver.
//
// # Discussion
// 
// This object will have an [index] value that is one more than the
// receiver’s. If there are no more subsequent siblings (that is, other
// child nodes of the receiver’s parent) the method returns `nil`.
//
// [index]: https://developer.apple.com/documentation/Foundation/XMLNode/index
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/nextSibling
func (x XMLNode) NextSibling() INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("nextSibling"))
	return NSXMLNodeFromID(objc.ID(rv))
}



// Returns the previous [NSXMLNode] object in document order.
//
// # Discussion
// 
// You use this method to “walk” backward through the tree structure
// representing an XML document or document section. (Use [NextNode] to
// traverse the tree in the opposite direction.) Document order is the natural
// order that XML constructs appear in markup text. If you send this message
// to the first node in the tree (that is, the root element), `nil` is
// returned. [NSXMLNode] bypasses namespace and attribute nodes when it
// traverses a tree in document order.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/previous
func (x XMLNode) PreviousNode() INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("previousNode"))
	return NSXMLNodeFromID(objc.ID(rv))
}



// Returns the previous [NSXMLNode] object that is a sibling node to the
// receiver.
//
// # Discussion
// 
// This object will have an [index] value that is one less than the
// receiver’s. If there are no more previous siblings (that is, other child
// nodes of the receiver’s parent) the method returns `nil`
//
// [index]: https://developer.apple.com/documentation/Foundation/XMLNode/index
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/previousSibling
func (x XMLNode) PreviousSibling() INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("previousSibling"))
	return NSXMLNodeFromID(objc.ID(rv))
}



// Returns the string representation of the receiver as it would appear in an
// XML document.
//
// # Discussion
// 
// The returned string includes the string representations of all children.
// This method invokes [XMLStringWithOptions] with an `options` argument of
// [NSXMLNodeOptionsNone].
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/xmlString
func (x XMLNode) XMLString() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("XMLString"))
	return NSStringFromID(rv).String()
}



//
// # Discussion
// 
// Used for debugging. May give more information than XMLString.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/description
func (x XMLNode) Description() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}



// Returns the XPath expression identifying the receiver’s location in the
// document tree.
//
// # Discussion
// 
// For example, this method might return a string such as
// “foo/bar[2]/baz”. The result of this method can be used directly in the
// [NodesForXPathError] and [ObjectsForXQueryConstantsError] methods.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/xPath
func (x XMLNode) XPath() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("XPath"))
	return NSStringFromID(rv).String()
}



// Returns the local name of the receiver.
//
// # Return Value
// 
// A string containing the local name of the receiver.
// 
// # Discussion
// 
// The local name is the part of a node name that follows a
// namespace-qualifying colon or the full name if there is no colon. For
// example, “chapter” is the local name in the qualified name
// “acme:chapter”.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/localName
func (x XMLNode) LocalName() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("localName"))
	return NSStringFromID(rv).String()
}



// Returns the prefix of the receiver’s name.
//
// # Return Value
// 
// A string containing the receiver’s prefix. This method returns an empty
// string if the receiver’s name is not qualified by a namespace.
// 
// # Discussion
// 
// The prefix is the part of a namespace-qualified name that precedes the
// colon. For example, “acme” is the prefix in the qualified name
// “acme:chapter”.
//
// See: https://developer.apple.com/documentation/Foundation/XMLNode/prefix
func (x XMLNode) Prefix() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("prefix"))
	return NSStringFromID(rv).String()
}













			// Protocol methods for NSCopying
			














