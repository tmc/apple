// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XMLElement] class.
var (
	_XMLElementClass     XMLElementClass
	_XMLElementClassOnce sync.Once
)

func getXMLElementClass() XMLElementClass {
	_XMLElementClassOnce.Do(func() {
		_XMLElementClass = XMLElementClass{class: objc.GetClass("NSXMLElement")}
	})
	return _XMLElementClass
}

// GetXMLElementClass returns the class object for NSXMLElement.
func GetXMLElementClass() XMLElementClass {
	return getXMLElementClass()
}

type XMLElementClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (xc XMLElementClass) Alloc() XMLElement {
	rv := objc.Send[XMLElement](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// The element nodes in an XML tree structure.
//
// # Overview
// 
// An [NSXMLElement] object may have child nodes, specifically comment nodes,
// processing-instruction nodes, text nodes, and other [NSXMLElement] nodes.
// It may also have attribute nodes and namespace nodes associated with it
// (however, namespace and attribute nodes are not considered children). Any
// attempt to add a [NSXMLDocument] node, [NSXMLDTD] node, namespace node, or
// attribute node as a child raises an exception. If you add a child node to
// an [NSXMLElement] object and that child already has a parent,
// [NSXMLElement] raises an exception; the child must be detached or copied
// first.
// 
// # Subclassing Notes
// 
// You can subclass [NSXMLElement] if you want element nodes with more
// specialized attributes or behavior, for example, paragraph and font
// attributes that specify how the string value of the element should appear.
// 
// # Methods to Override
// 
// To subclass [NSXMLElement] you need to override the primary initializer,
// [InitWithNameURI], and the methods listed below. In most cases, you need
// only invoke the superclass implementation, adding any subclass-specific
// code before or after the invocation, as necessary.
// 
// [Table data omitted]
// 
// [NSXMLElement] implements [isEqual(_:)] to perform a deep comparison: two
// [NSXMLDocument] objects are not considered equal unless they have the same
// name, same child nodes, same attributes, and so on. If you want a different
// standard of comparison, override ``.
// 
// # Special Considerations
// 
// Because of the architecture and data model of NSXML, when it parses and
// processes a source of XML it cannot know about your subclass unless you
// override the class method [ReplacementClassForClass] to return your custom
// class in place of an NSXML class. If your custom class has no direct NSXML
// counterpart—for example, it is a subclass of [NSXMLNode] that represents
// CDATA sections—then you can walk the tree after it has been created and
// insert the new node where appropriate.
// 
// Note that you can safely set the root element of the XML document (using
// the [NSXMLDocument] [SetRootElement]method) to be an instance of your
// subclass because this method only checks to see if the added node is of an
// element kind ([NSXMLElementKind]). These precautions do not apply, of
// course, if you are creating an XML tree programmatically.
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// # Initializing NSXMLElement Objects
//
//   - [XMLElement.InitWithName]: Returns an [NSXMLElement] object initialized with the specified name.
//   - [XMLElement.InitWithNameStringValue]: Returns an [NSXMLElement] object initialized with a specified name and a single text-node child containing a specified value.
//   - [XMLElement.InitWithNameURI]: Returns an [NSXMLElement] object initialized with the specified name and URI.
//   - [XMLElement.InitWithXMLStringError]: Returns an [NSXMLElement] object created from a specified string containing XML markup.
//
// # Obtaining Child Elements
//
//   - [XMLElement.ElementsForName]: Returns the child element nodes (as [NSXMLElement] objects) of the receiver that have a specified name.
//   - [XMLElement.ElementsForLocalNameURI]: Returns the child element nodes (as [NSXMLElement] objects) of the receiver that are matched with the specified local name and URI.
//
// # Manipulating Child Elements
//
//   - [XMLElement.AddChild]: Adds a child node at the end of the receiver’s current list of children.
//   - [XMLElement.InsertChildAtIndex]: Inserts a new child node at a specified location in the receiver’s list of child nodes.
//   - [XMLElement.InsertChildrenAtIndex]: Inserts an array of child nodes at a specified location in the receiver’s list of children.
//   - [XMLElement.RemoveChildAtIndex]: Removes the child node of the receiver identified by a given index.
//   - [XMLElement.ReplaceChildAtIndexWithNode]: Replaces a child node at a specified location with another child node.
//   - [XMLElement.NormalizeAdjacentTextNodesPreservingCDATA]: Coalesces adjacent text nodes of the receiver that you have explicitly added, optionally including CDATA sections.
//
// # Handling Attributes
//
//   - [XMLElement.AddAttribute]: Adds an attribute node to the receiver.
//   - [XMLElement.AttributeForName]: Returns the attribute node of the receiver with the specified name.
//   - [XMLElement.AttributeForLocalNameURI]: Returns the attribute node of the receiver that is identified by a local name and URI.
//   - [XMLElement.Attributes]: Sets all attributes of the receiver at once, replacing any existing attribute nodes.
//   - [XMLElement.SetAttributes]
//   - [XMLElement.RemoveAttributeForName]: Removes an attribute node identified by name.
//   - [XMLElement.SetAttributesWithDictionary]: Sets the attributes of the receiver based on the key-value pairs specified in the passed dictionary.
//
// # Handling Namespaces
//
//   - [XMLElement.AddNamespace]: Adds a namespace node to the receiver.
//   - [XMLElement.Namespaces]: Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
//   - [XMLElement.SetNamespaces]
//   - [XMLElement.NamespaceForPrefix]: Returns the namespace node with a specified prefix.
//   - [XMLElement.RemoveNamespaceForPrefix]: Removes a namespace node that is identified by a given prefix.
//   - [XMLElement.ResolveNamespaceForName]: Returns the namespace node with the prefix matching the given qualified name.
//   - [XMLElement.ResolvePrefixForNamespaceURI]: Returns the prefix associated with the specified URI.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement
type XMLElement struct {
	NSXMLNode
}

// XMLElementFromID constructs a [XMLElement] from an objc.ID.
//
// The element nodes in an XML tree structure.
func XMLElementFromID(id objc.ID) XMLElement {
	return NSXMLElement{NSXMLNode: NSXMLNodeFromID(id)}
}

// NSXMLElementFromID is an alias for [XMLElementFromID] for cross-framework compatibility.
func NSXMLElementFromID(id objc.ID) XMLElement { return XMLElementFromID(id) }
// NOTE: XMLElement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [XMLElement] class.
//
// # Initializing NSXMLElement Objects
//
//   - [IXMLElement.InitWithName]: Returns an [NSXMLElement] object initialized with the specified name.
//   - [IXMLElement.InitWithNameStringValue]: Returns an [NSXMLElement] object initialized with a specified name and a single text-node child containing a specified value.
//   - [IXMLElement.InitWithNameURI]: Returns an [NSXMLElement] object initialized with the specified name and URI.
//   - [IXMLElement.InitWithXMLStringError]: Returns an [NSXMLElement] object created from a specified string containing XML markup.
//
// # Obtaining Child Elements
//
//   - [IXMLElement.ElementsForName]: Returns the child element nodes (as [NSXMLElement] objects) of the receiver that have a specified name.
//   - [IXMLElement.ElementsForLocalNameURI]: Returns the child element nodes (as [NSXMLElement] objects) of the receiver that are matched with the specified local name and URI.
//
// # Manipulating Child Elements
//
//   - [IXMLElement.AddChild]: Adds a child node at the end of the receiver’s current list of children.
//   - [IXMLElement.InsertChildAtIndex]: Inserts a new child node at a specified location in the receiver’s list of child nodes.
//   - [IXMLElement.InsertChildrenAtIndex]: Inserts an array of child nodes at a specified location in the receiver’s list of children.
//   - [IXMLElement.RemoveChildAtIndex]: Removes the child node of the receiver identified by a given index.
//   - [IXMLElement.ReplaceChildAtIndexWithNode]: Replaces a child node at a specified location with another child node.
//   - [IXMLElement.NormalizeAdjacentTextNodesPreservingCDATA]: Coalesces adjacent text nodes of the receiver that you have explicitly added, optionally including CDATA sections.
//
// # Handling Attributes
//
//   - [IXMLElement.AddAttribute]: Adds an attribute node to the receiver.
//   - [IXMLElement.AttributeForName]: Returns the attribute node of the receiver with the specified name.
//   - [IXMLElement.AttributeForLocalNameURI]: Returns the attribute node of the receiver that is identified by a local name and URI.
//   - [IXMLElement.Attributes]: Sets all attributes of the receiver at once, replacing any existing attribute nodes.
//   - [IXMLElement.SetAttributes]
//   - [IXMLElement.RemoveAttributeForName]: Removes an attribute node identified by name.
//   - [IXMLElement.SetAttributesWithDictionary]: Sets the attributes of the receiver based on the key-value pairs specified in the passed dictionary.
//
// # Handling Namespaces
//
//   - [IXMLElement.AddNamespace]: Adds a namespace node to the receiver.
//   - [IXMLElement.Namespaces]: Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
//   - [IXMLElement.SetNamespaces]
//   - [IXMLElement.NamespaceForPrefix]: Returns the namespace node with a specified prefix.
//   - [IXMLElement.RemoveNamespaceForPrefix]: Removes a namespace node that is identified by a given prefix.
//   - [IXMLElement.ResolveNamespaceForName]: Returns the namespace node with the prefix matching the given qualified name.
//   - [IXMLElement.ResolvePrefixForNamespaceURI]: Returns the prefix associated with the specified URI.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement
type IXMLElement interface {
	INSXMLNode

	// Topic: Initializing NSXMLElement Objects

	// Returns an [NSXMLElement] object initialized with the specified name.
	InitWithName(name string) XMLElement
	// Returns an [NSXMLElement] object initialized with a specified name and a single text-node child containing a specified value.
	InitWithNameStringValue(name string, string_ string) XMLElement
	// Returns an [NSXMLElement] object initialized with the specified name and URI.
	InitWithNameURI(name string, URI string) XMLElement
	// Returns an [NSXMLElement] object created from a specified string containing XML markup.
	InitWithXMLStringError(string_ string) (XMLElement, error)

	// Topic: Obtaining Child Elements

	// Returns the child element nodes (as [NSXMLElement] objects) of the receiver that have a specified name.
	ElementsForName(name string) []NSXMLElement
	// Returns the child element nodes (as [NSXMLElement] objects) of the receiver that are matched with the specified local name and URI.
	ElementsForLocalNameURI(localName string, URI string) []NSXMLElement

	// Topic: Manipulating Child Elements

	// Adds a child node at the end of the receiver’s current list of children.
	AddChild(child INSXMLNode)
	// Inserts a new child node at a specified location in the receiver’s list of child nodes.
	InsertChildAtIndex(child INSXMLNode, index uint)
	// Inserts an array of child nodes at a specified location in the receiver’s list of children.
	InsertChildrenAtIndex(children []NSXMLNode, index uint)
	// Removes the child node of the receiver identified by a given index.
	RemoveChildAtIndex(index uint)
	// Replaces a child node at a specified location with another child node.
	ReplaceChildAtIndexWithNode(index uint, node INSXMLNode)
	// Coalesces adjacent text nodes of the receiver that you have explicitly added, optionally including CDATA sections.
	NormalizeAdjacentTextNodesPreservingCDATA(preserve bool)

	// Topic: Handling Attributes

	// Adds an attribute node to the receiver.
	AddAttribute(attribute INSXMLNode)
	// Returns the attribute node of the receiver with the specified name.
	AttributeForName(name string) INSXMLNode
	// Returns the attribute node of the receiver that is identified by a local name and URI.
	AttributeForLocalNameURI(localName string, URI string) INSXMLNode
	// Sets all attributes of the receiver at once, replacing any existing attribute nodes.
	Attributes() []NSXMLNode
	SetAttributes(value []NSXMLNode)
	// Removes an attribute node identified by name.
	RemoveAttributeForName(name string)
	// Sets the attributes of the receiver based on the key-value pairs specified in the passed dictionary.
	SetAttributesWithDictionary(attributes INSDictionary)

	// Topic: Handling Namespaces

	// Adds a namespace node to the receiver.
	AddNamespace(aNamespace INSXMLNode)
	// Sets all of the namespace nodes of the receiver at once, replacing any existing namespace nodes.
	Namespaces() []NSXMLNode
	SetNamespaces(value []NSXMLNode)
	// Returns the namespace node with a specified prefix.
	NamespaceForPrefix(name string) INSXMLNode
	// Removes a namespace node that is identified by a given prefix.
	RemoveNamespaceForPrefix(name string)
	// Returns the namespace node with the prefix matching the given qualified name.
	ResolveNamespaceForName(name string) INSXMLNode
	// Returns the prefix associated with the specified URI.
	ResolvePrefixForNamespaceURI(namespaceURI string) string
}

// Init initializes the instance.
func (x XMLElement) Init() XMLElement {
	rv := objc.Send[XMLElement](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XMLElement) Autorelease() XMLElement {
	rv := objc.Send[XMLElement](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLElement creates a new XMLElement instance.
func NewXMLElement() XMLElement {
	class := getXMLElementClass()
	rv := objc.Send[XMLElement](objc.ID(class.class), objc.Sel("new"))
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
func NewXMLElementWithKind(kind NSXMLNodeKind) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:"), kind)
	return XMLElementFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/init(kind:options:)
func NewXMLElementWithKindOptions(kind NSXMLNodeKind, options NSXMLNodeOptions) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:options:"), kind, options)
	return XMLElementFromID(rv)
}

// Returns an [NSXMLElement] object initialized with the specified name.
//
// name: A string specifying the name of the element.
//
// # Return Value
// 
// The initialized [NSXMLElement] object or `nil` if initialization did not
// succeed.
//
// # Discussion
// 
// The XML string representation of this object is ``. This method invokes
// [InitWithNameURI] with the URI parameter set to `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:)
func NewXMLElementWithName(name string) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	return XMLElementFromID(rv)
}

// Returns an [NSXMLElement] object initialized with a specified name and a
// single text-node child containing a specified value.
//
// name: A string specifying the name of the element.
//
// string: The string value of the receiver’s text node.
//
// # Return Value
// 
// The initialized [NSXMLElement] object or `nil` if initialization did not
// succeed.
//
// # Discussion
// 
// The string representation of this object is ```string```.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:stringValue:)
func NewXMLElementWithNameStringValue(name string, string_ string) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:stringValue:"), objc.String(name), objc.String(string_))
	return XMLElementFromID(rv)
}

// Returns an [NSXMLElement] object initialized with the specified name and
// URI.
//
// name: A string that specifies the qualified name of the element.
//
// URI: A string that specifies the namespace URI associated with the element.
//
// # Return Value
// 
// The initialized [NSXMLElement] object or `nil` if initialization did not
// succeed.
//
// # Discussion
// 
// You can look up the namespace prefix for this element node based on its URI
// using [ResolvePrefixForNamespaceURI]. This method is the primary
// initializer for the [NSXMLElement] class.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:uri:)
func NewXMLElementWithNameURI(name string, URI string) XMLElement {
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:URI:"), objc.String(name), objc.String(URI))
	return XMLElementFromID(rv)
}

// Returns an [NSXMLElement] object created from a specified string containing
// XML markup.
//
// string: A string containing XML markup for an element.
//
// # Return Value
// 
// The initialized [NSXMLElement] object or `nil` if initialization did not
// succeed.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/init(xmlString:)
func NewXMLElementWithXMLStringError(string_ string) (XMLElement, error) {
	var errorPtr objc.ID
	instance := getXMLElementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXMLString:error:"), objc.String(string_), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLElementFromID(rv), NSErrorFrom(errorPtr)
	}
	return XMLElementFromID(rv), nil
}

// Returns an [NSXMLElement] object initialized with the specified name.
//
// name: A string specifying the name of the element.
//
// # Return Value
// 
// The initialized [NSXMLElement] object or `nil` if initialization did not
// succeed.
//
// # Discussion
// 
// The XML string representation of this object is ``. This method invokes
// [InitWithNameURI] with the URI parameter set to `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:)
func (x XMLElement) InitWithName(name string) XMLElement {
	rv := objc.Send[XMLElement](x.ID, objc.Sel("initWithName:"), objc.String(name))
	return rv
}

// Returns an [NSXMLElement] object initialized with a specified name and a
// single text-node child containing a specified value.
//
// name: A string specifying the name of the element.
//
// string: The string value of the receiver’s text node.
//
// # Return Value
// 
// The initialized [NSXMLElement] object or `nil` if initialization did not
// succeed.
//
// # Discussion
// 
// The string representation of this object is ```string```.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:stringValue:)
func (x XMLElement) InitWithNameStringValue(name string, string_ string) XMLElement {
	rv := objc.Send[XMLElement](x.ID, objc.Sel("initWithName:stringValue:"), objc.String(name), objc.String(string_))
	return rv
}

// Returns an [NSXMLElement] object initialized with the specified name and
// URI.
//
// name: A string that specifies the qualified name of the element.
//
// URI: A string that specifies the namespace URI associated with the element.
//
// # Return Value
// 
// The initialized [NSXMLElement] object or `nil` if initialization did not
// succeed.
//
// # Discussion
// 
// You can look up the namespace prefix for this element node based on its URI
// using [ResolvePrefixForNamespaceURI]. This method is the primary
// initializer for the [NSXMLElement] class.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/init(name:uri:)
func (x XMLElement) InitWithNameURI(name string, URI string) XMLElement {
	rv := objc.Send[XMLElement](x.ID, objc.Sel("initWithName:URI:"), objc.String(name), objc.String(URI))
	return rv
}

// Returns an [NSXMLElement] object created from a specified string containing
// XML markup.
//
// string: A string containing XML markup for an element.
//
// # Return Value
// 
// The initialized [NSXMLElement] object or `nil` if initialization did not
// succeed.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/init(xmlString:)
func (x XMLElement) InitWithXMLStringError(string_ string) (XMLElement, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("initWithXMLString:error:"), objc.String(string_), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLElement{}, NSErrorFrom(errorPtr)
	}
	return NSXMLElementFromID(rv), nil

}

// Returns the child element nodes (as [NSXMLElement] objects) of the receiver
// that have a specified name.
//
// name: A string specifying the name of the child element nodes to find and return.
// If `name` is a qualified name, then this method invokes
// [ElementsForLocalNameURI] with the URI parameter set to the URI associated
// with the prefix. Otherwise comparison is based on string equality of the
// qualified or non-qualified name.
//
// # Return Value
// 
// An array of of [NSXMLElement] objects or an empty array if no matching
// children can be found.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/elements(forName:)
func (x XMLElement) ElementsForName(name string) []NSXMLElement {
	rv := objc.Send[[]objc.ID](x.ID, objc.Sel("elementsForName:"), objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) NSXMLElement {
		return NSXMLElementFromID(id)
	})
}

// Returns the child element nodes (as [NSXMLElement] objects) of the receiver
// that are matched with the specified local name and URI.
//
// localName: A string specifying a local name of an element.
//
// URI: A string specifying a URI associated with an element.
//
// # Return Value
// 
// An array of [NSXMLElement] objects or an empty array if no matching
// children could be found.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/elements(forLocalName:uri:)
func (x XMLElement) ElementsForLocalNameURI(localName string, URI string) []NSXMLElement {
	rv := objc.Send[[]objc.ID](x.ID, objc.Sel("elementsForLocalName:URI:"), objc.String(localName), objc.String(URI))
	return objc.ConvertSlice(rv, func(id objc.ID) NSXMLElement {
		return NSXMLElementFromID(id)
	})
}

// Adds a child node at the end of the receiver’s current list of children.
//
// child: An XML node object to add to the receiver’s children.
//
// # Discussion
// 
// The new node has an index value that is one greater than the last of the
// current children.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/addChild(_:)
func (x XMLElement) AddChild(child INSXMLNode) {
	objc.Send[objc.ID](x.ID, objc.Sel("addChild:"), child)
}

// Inserts a new child node at a specified location in the receiver’s list
// of child nodes.
//
// child: An XML node object to be inserted as a child of the receiver.
//
// index: An integer identifying a position in the receiver’s list of children. An
// exception is raised if `index` is out of bounds.
//
// # Discussion
// 
// Insertion of the node increments the indexes of sibling nodes after it.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/insertChild(_:at:)
func (x XMLElement) InsertChildAtIndex(child INSXMLNode, index uint) {
	objc.Send[objc.ID](x.ID, objc.Sel("insertChild:atIndex:"), child, index)
}

// Inserts an array of child nodes at a specified location in the receiver’s
// list of children.
//
// children: An array of XML node objects to add as children of the receiver.
//
// index: An integer identifying a position in the receiver’s list of children. An
// exception is raised if `index` is out of bounds.
//
// # Discussion
// 
// Insertion of the node increases the indexes of sibling nodes after it by
// the count of `children`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/insertChildren(_:at:)
func (x XMLElement) InsertChildrenAtIndex(children []NSXMLNode, index uint) {
	objc.Send[objc.ID](x.ID, objc.Sel("insertChildren:atIndex:"), objectivec.IObjectSliceToNSArray(children), index)
}

// Removes the child node of the receiver identified by a given index.
//
// index: An integer identifying the node in the receiver’s list of children to
// remove. An exception is raised if `index` is out of bounds.
//
// # Discussion
// 
// The XML node object is released upon removal. The indices of subsequent
// children are decremented by one.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/removeChild(at:)
func (x XMLElement) RemoveChildAtIndex(index uint) {
	objc.Send[objc.ID](x.ID, objc.Sel("removeChildAtIndex:"), index)
}

// Replaces a child node at a specified location with another child node.
//
// index: An integer identifying a position in the receiver’s list of children. An
// exception is raised if `index` is out of bounds.
//
// node: An XML node object that will replace the current child.
//
// # Discussion
// 
// The replaced XML node object is released upon removal.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/replaceChild(at:with:)
func (x XMLElement) ReplaceChildAtIndexWithNode(index uint, node INSXMLNode) {
	objc.Send[objc.ID](x.ID, objc.Sel("replaceChildAtIndex:withNode:"), index, node)
}

// Coalesces adjacent text nodes of the receiver that you have explicitly
// added, optionally including CDATA sections.
//
// preserve: [true] if CDATA sections are left alone as text nodes, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// A text node with a value of an empty string is removed. When you process an
// input source of XML, adjacent text nodes are automatically normalized. You
// should invoke this method (with `preserve` as [false]) before using the
// [NSXMLNode] methods [ObjectsForXQueryConstantsError] or
// [NodesForXPathError].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/normalizeAdjacentTextNodesPreservingCDATA(_:)
func (x XMLElement) NormalizeAdjacentTextNodesPreservingCDATA(preserve bool) {
	objc.Send[objc.ID](x.ID, objc.Sel("normalizeAdjacentTextNodesPreservingCDATA:"), preserve)
}

// Adds an attribute node to the receiver.
//
// attribute: An XML node object representing an attribute. If the receiver already has
// an attribute with the same name, `anAttribute` replaces the old attribute.
//
// # Discussion
// 
// The order of multiple attributes is preserved if the
// [NSXMLPreserveAttributeOrder] option is specified when the element is
// created.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/addAttribute(_:)
func (x XMLElement) AddAttribute(attribute INSXMLNode) {
	objc.Send[objc.ID](x.ID, objc.Sel("addAttribute:"), attribute)
}

// Returns the attribute node of the receiver with the specified name.
//
// name: A string specifying the name of an attribute.
//
// # Return Value
// 
// An XML node object representing a matching attribute or `nil` if no such
// node was found.
//
// # Discussion
// 
// If `name` is a qualified name, then this method invokes
// [AttributeForLocalNameURI] with the URI parameter set to the URI associated
// with the prefix. Otherwise comparison is based on string equality of the
// qualified or non-qualified name.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/attribute(forName:)
func (x XMLElement) AttributeForName(name string) INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("attributeForName:"), objc.String(name))
	return NSXMLNodeFromID(rv)
}

// Returns the attribute node of the receiver that is identified by a local
// name and URI.
//
// localName: A string specifying the local name of an attribute.
//
// URI: A sting identifying the URI associated with an attribute.
//
// # Return Value
// 
// An XML node object representing a matching attribute or `nil` if no such
// node was found.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/attribute(forLocalName:uri:)
func (x XMLElement) AttributeForLocalNameURI(localName string, URI string) INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("attributeForLocalName:URI:"), objc.String(localName), objc.String(URI))
	return NSXMLNodeFromID(rv)
}

// Removes an attribute node identified by name.
//
// name: A string specifying the name of an attribute.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/removeAttribute(forName:)
func (x XMLElement) RemoveAttributeForName(name string) {
	objc.Send[objc.ID](x.ID, objc.Sel("removeAttributeForName:"), objc.String(name))
}

// Sets the attributes of the receiver based on the key-value pairs specified
// in the passed dictionary.
//
// attributes: A dictionary of key-value pairs where the attribute name is the key and the
// object value of the attribute is the dictionary value.
//
// # Discussion
// 
// The method uses these names and object values to create [NSXMLNode] objects
// of kind [XMLNode.Kind.attribute]. Existing attributes are removed.
//
// [XMLNode.Kind.attribute]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/attribute
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/setAttributesWith(_:)
func (x XMLElement) SetAttributesWithDictionary(attributes INSDictionary) {
	objc.Send[objc.ID](x.ID, objc.Sel("setAttributesWithDictionary:"), attributes)
}

// Adds a namespace node to the receiver.
//
// aNamespace: An XML node object of kind [XMLNode.Kind.namespace]. If the receiver
// already has a namespace with the same name, `aNamespace` is not added.
// //
// [XMLNode.Kind.namespace]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/namespace
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/addNamespace(_:)
func (x XMLElement) AddNamespace(aNamespace INSXMLNode) {
	objc.Send[objc.ID](x.ID, objc.Sel("addNamespace:"), aNamespace)
}

// Returns the namespace node with a specified prefix.
//
// name: A string specifying a namespace prefix.
//
// # Return Value
// 
// An [NSXMLNode] object of kind [XMLNode.Kind.namespace] or `nil` if there is
// no namespace node with that prefix.
//
// [XMLNode.Kind.namespace]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/namespace
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/namespace(forPrefix:)
func (x XMLElement) NamespaceForPrefix(name string) INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("namespaceForPrefix:"), objc.String(name))
	return NSXMLNodeFromID(rv)
}

// Removes a namespace node that is identified by a given prefix.
//
// name: A string that is the prefix for a namespace.
//
// # Discussion
// 
// The removed XML node object is removed.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/removeNamespace(forPrefix:)
func (x XMLElement) RemoveNamespaceForPrefix(name string) {
	objc.Send[objc.ID](x.ID, objc.Sel("removeNamespaceForPrefix:"), objc.String(name))
}

// Returns the namespace node with the prefix matching the given qualified
// name.
//
// name: A string that is the qualified name for a namespace (a qualified name is
// prefix plus local name).
//
// # Return Value
// 
// An [NSXMLNode] object of kind [XMLNode.Kind.namespace] or `nil` if there is
// no matching namespace node.
//
// [XMLNode.Kind.namespace]: https://developer.apple.com/documentation/Foundation/XMLNode/Kind-swift.enum/namespace
//
// # Discussion
// 
// The method looks in the entire namespace chain for the prefix.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/resolveNamespace(forName:)
func (x XMLElement) ResolveNamespaceForName(name string) INSXMLNode {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("resolveNamespaceForName:"), objc.String(name))
	return NSXMLNodeFromID(rv)
}

// Returns the prefix associated with the specified URI.
//
// namespaceURI: A string identifying the URI associated with the namespace.
//
// # Return Value
// 
// A string that is the matching prefix or `nil` if it finds no matching
// prefix.
//
// # Discussion
// 
// The method looks in the entire namespace chain for the URI.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/resolvePrefix(forNamespaceURI:)
func (x XMLElement) ResolvePrefixForNamespaceURI(namespaceURI string) string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("resolvePrefixForNamespaceURI:"), objc.String(namespaceURI))
	return NSStringFromID(rv).String()
}

// Sets all attributes of the receiver at once, replacing any existing
// attribute nodes.
//
// # Discussion
// 
// To set attributes in an element node using an [NSDictionary] object as the
// input parameter, see [SetAttributesWithDictionary].
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/attributes
func (x XMLElement) Attributes() []NSXMLNode {
	rv := objc.Send[[]objc.ID](x.ID, objc.Sel("attributes"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSXMLNode {
		return NSXMLNodeFromID(id)
	})
}
func (x XMLElement) SetAttributes(value []NSXMLNode) {
	objc.Send[struct{}](x.ID, objc.Sel("setAttributes:"), objectivec.IObjectSliceToNSArray(value))
}

// Sets all of the namespace nodes of the receiver at once, replacing any
// existing namespace nodes.
//
// See: https://developer.apple.com/documentation/Foundation/XMLElement/namespaces
func (x XMLElement) Namespaces() []NSXMLNode {
	rv := objc.Send[[]objc.ID](x.ID, objc.Sel("namespaces"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSXMLNode {
		return NSXMLNodeFromID(id)
	})
}
func (x XMLElement) SetNamespaces(value []NSXMLNode) {
	objc.Send[struct{}](x.ID, objc.Sel("setNamespaces:"), objectivec.IObjectSliceToNSArray(value))
}

