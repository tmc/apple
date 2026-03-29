// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XMLDocument] class.
var (
	_XMLDocumentClass     XMLDocumentClass
	_XMLDocumentClassOnce sync.Once
)

func getXMLDocumentClass() XMLDocumentClass {
	_XMLDocumentClassOnce.Do(func() {
		_XMLDocumentClass = XMLDocumentClass{class: objc.GetClass("NSXMLDocument")}
	})
	return _XMLDocumentClass
}

// GetXMLDocumentClass returns the class object for NSXMLDocument.
func GetXMLDocumentClass() XMLDocumentClass {
	return getXMLDocumentClass()
}

type XMLDocumentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (xc XMLDocumentClass) Class() objc.Class {
	return xc.class
}

// Alloc allocates memory for a new instance of the class.
func (xc XMLDocumentClass) Alloc() XMLDocument {
	rv := objc.Send[XMLDocument](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// An XML document as internalized into a logical tree structure.
//
// # Overview
// 
// An [NSXMLDocument] object can have multiple child nodes but only one
// element, the root element. Any other node must be a [NSXMLNode] object
// representing a comment or a processing instruction. If you attempt to add
// any other kind of child node to an [NSXMLDocument] object, such as an
// attribute, namespace, another document object, or an element other than the
// root, [NSXMLDocument] raises an exception. If you add a valid child node
// and that object already has a parent, [NSXMLDocument] raises an exception.
// An [NSXMLDocument] object may also have document-global attributes, such as
// XML version, character encoding, referenced DTD, and MIME type.
// 
// The initializers of the [NSXMLDocument] class read an external source of
// XML, whether it be a local file or remote website, parse it, and process it
// into the tree representation. You can also construct an [NSXMLDocument]
// programmatically. There are accessor methods for getting and setting
// document attributes, methods for transforming documents using XSLT, a
// method for dynamically validating a document, and methods for printing out
// the content of an [NSXMLDocument] as XML, XHTML, HTML, or plain text.
// 
// The [NSXMLDocument] class is thread-safe as long as any given instance is
// used only in one thread.
// 
// # Subclassing Notes
// 
// # Methods to Override
// 
// To subclass [NSXMLDocument] you need to override the primary initializer,
// [InitWithDataOptionsError], and the methods listed below. In most cases,
// you need only invoke the superclass implementation, adding any
// subclass-specific code before or after the invocation, as necessary.
// 
// - [RootElement] - [SetChildren] - [RemoveChildAtIndex] -
// [InsertChildAtIndex] - [CharacterEncoding] - [CharacterEncoding] -
// [DocumentContentKind] - [DocumentContentKind] - [DTD] - [MIMEType] -
// [Standalone] - [Version] - [Version]
// 
// By default [NSXMLDocument] implements the [NSObject] [isEqual(_:)] method
// to perform a deep comparison: two [NSXMLDocument] objects are not
// considered equal unless they have the same name, same child nodes, same
// attributes, and so on. The comparison does not consider the parent node
// (and hence the node’s location). If you want a different standard of
// comparison, override ``.
// 
// # Special Considerations
// 
// Because of the architecture and data model of NSXML, when it parses and
// processes a source of XML it cannot know about your subclass unless you
// override the class method [ReplacementClassForClass] to return your custom
// class in place of an [NSXML] class. If your custom class has no direct
// [NSXML] counterpart—for example, it is a subclass of [NSXMLNode] that
// represents CDATA sections—then you can walk the tree after it has been
// created and insert the new node where appropriate.
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// # Initializing NSXMLDocument Objects
//
//   - [XMLDocument.InitWithContentsOfURLOptionsError]: Initializes and returns an NSXMLDocument object created from the XML or HTML contents of a URL-referenced source
//   - [XMLDocument.InitWithDataOptionsError]: Initializes and returns an [NSXMLDocument] object created from an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object.
//   - [XMLDocument.InitWithRootElement]: Returns an [NSXMLDocument] object initialized with a single child, the root element.
//   - [XMLDocument.InitWithXMLStringOptionsError]: Initializes and returns an [NSXMLDocument] object created from a string containing XML markup text.
//
// # Managing Document Attributes
//
//   - [XMLDocument.CharacterEncoding]: Sets the character encoding of the receiver to `encoding`,
//   - [XMLDocument.SetCharacterEncoding]
//   - [XMLDocument.DocumentContentKind]: Sets the kind of output content for the receiver.
//   - [XMLDocument.SetDocumentContentKind]
//   - [XMLDocument.DTD]: Returns an [XMLDTD](<doc://com.apple.foundation/documentation/Foundation/XMLDTD>) object representing the internal DTD associated with the receiver.
//   - [XMLDocument.SetDTD]
//   - [XMLDocument.Standalone]: Sets a Boolean value that specifies whether the receiver represents a standalone XML document.
//   - [XMLDocument.SetStandalone]
//   - [XMLDocument.MIMEType]: Returns the MIME type for the receiver.
//   - [XMLDocument.SetMIMEType]
//   - [XMLDocument.Version]: Sets the version of the receiver’s XML.
//   - [XMLDocument.SetVersion]
//
// # Managing the Root Element
//
//   - [XMLDocument.RootElement]: Returns the root element of the receiver.
//   - [XMLDocument.SetRootElement]: Set the root element of the receiver.
//
// # Adding and Removing Child Nodes
//
//   - [XMLDocument.AddChild]: Adds a child node after the last of the receiver’s existing children.
//   - [XMLDocument.InsertChildAtIndex]: Inserts a node object at specified position in the receiver’s array of children.
//   - [XMLDocument.InsertChildrenAtIndex]: Inserts an array of children at a specified position in the receiver’s array of children.
//   - [XMLDocument.RemoveChildAtIndex]: Removes the child node of the receiver located at a specified position in its array of children.
//   - [XMLDocument.ReplaceChildAtIndexWithNode]: Replaces the child node of the receiver located at a specified position in its array of children with another node.
//
// # Transforming a Document Using XSLT
//
//   - [XMLDocument.ObjectByApplyingXSLTArgumentsError]: Applies the XSLT pattern rules and templates (specified as a data object) to the receiver and returns a document object containing transformed XML or HTML markup.
//   - [XMLDocument.ObjectByApplyingXSLTStringArgumentsError]: Applies the XSLT pattern rules and templates (specified as a string) to the receiver and returns a document object containing transformed XML or HTML markup.
//   - [XMLDocument.ObjectByApplyingXSLTAtURLArgumentsError]: Applies the XSLT pattern rules and templates located at a specified URL to the receiver and returns a document object containing transformed XML markup or an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object containing plain text, RTF text, and so on.
//
// # Writing a Document as XML Data
//
//   - [XMLDocument.XMLData]: Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//   - [XMLDocument.XMLDataWithOptions]: Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//
// # Validating a Document
//
//   - [XMLDocument.ValidateAndReturnError]: Validates the document against the governing schema and returns whether the document conforms to the schema.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument
type XMLDocument struct {
	NSXMLNode
}

// XMLDocumentFromID constructs a [XMLDocument] from an objc.ID.
//
// An XML document as internalized into a logical tree structure.
func XMLDocumentFromID(id objc.ID) XMLDocument {
	return NSXMLDocument{NSXMLNode: NSXMLNodeFromID(id)}
}

// NSXMLDocumentFromID is an alias for [XMLDocumentFromID] for cross-framework compatibility.
func NSXMLDocumentFromID(id objc.ID) XMLDocument { return XMLDocumentFromID(id) }
// NOTE: XMLDocument adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [XMLDocument] class.
//
// # Initializing NSXMLDocument Objects
//
//   - [IXMLDocument.InitWithContentsOfURLOptionsError]: Initializes and returns an NSXMLDocument object created from the XML or HTML contents of a URL-referenced source
//   - [IXMLDocument.InitWithDataOptionsError]: Initializes and returns an [NSXMLDocument] object created from an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object.
//   - [IXMLDocument.InitWithRootElement]: Returns an [NSXMLDocument] object initialized with a single child, the root element.
//   - [IXMLDocument.InitWithXMLStringOptionsError]: Initializes and returns an [NSXMLDocument] object created from a string containing XML markup text.
//
// # Managing Document Attributes
//
//   - [IXMLDocument.CharacterEncoding]: Sets the character encoding of the receiver to `encoding`,
//   - [IXMLDocument.SetCharacterEncoding]
//   - [IXMLDocument.DocumentContentKind]: Sets the kind of output content for the receiver.
//   - [IXMLDocument.SetDocumentContentKind]
//   - [IXMLDocument.DTD]: Returns an [XMLDTD](<doc://com.apple.foundation/documentation/Foundation/XMLDTD>) object representing the internal DTD associated with the receiver.
//   - [IXMLDocument.SetDTD]
//   - [IXMLDocument.Standalone]: Sets a Boolean value that specifies whether the receiver represents a standalone XML document.
//   - [IXMLDocument.SetStandalone]
//   - [IXMLDocument.MIMEType]: Returns the MIME type for the receiver.
//   - [IXMLDocument.SetMIMEType]
//   - [IXMLDocument.Version]: Sets the version of the receiver’s XML.
//   - [IXMLDocument.SetVersion]
//
// # Managing the Root Element
//
//   - [IXMLDocument.RootElement]: Returns the root element of the receiver.
//   - [IXMLDocument.SetRootElement]: Set the root element of the receiver.
//
// # Adding and Removing Child Nodes
//
//   - [IXMLDocument.AddChild]: Adds a child node after the last of the receiver’s existing children.
//   - [IXMLDocument.InsertChildAtIndex]: Inserts a node object at specified position in the receiver’s array of children.
//   - [IXMLDocument.InsertChildrenAtIndex]: Inserts an array of children at a specified position in the receiver’s array of children.
//   - [IXMLDocument.RemoveChildAtIndex]: Removes the child node of the receiver located at a specified position in its array of children.
//   - [IXMLDocument.ReplaceChildAtIndexWithNode]: Replaces the child node of the receiver located at a specified position in its array of children with another node.
//
// # Transforming a Document Using XSLT
//
//   - [IXMLDocument.ObjectByApplyingXSLTArgumentsError]: Applies the XSLT pattern rules and templates (specified as a data object) to the receiver and returns a document object containing transformed XML or HTML markup.
//   - [IXMLDocument.ObjectByApplyingXSLTStringArgumentsError]: Applies the XSLT pattern rules and templates (specified as a string) to the receiver and returns a document object containing transformed XML or HTML markup.
//   - [IXMLDocument.ObjectByApplyingXSLTAtURLArgumentsError]: Applies the XSLT pattern rules and templates located at a specified URL to the receiver and returns a document object containing transformed XML markup or an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object containing plain text, RTF text, and so on.
//
// # Writing a Document as XML Data
//
//   - [IXMLDocument.XMLData]: Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//   - [IXMLDocument.XMLDataWithOptions]: Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
//
// # Validating a Document
//
//   - [IXMLDocument.ValidateAndReturnError]: Validates the document against the governing schema and returns whether the document conforms to the schema.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument
type IXMLDocument interface {
	INSXMLNode

	// Topic: Initializing NSXMLDocument Objects

	// Initializes and returns an NSXMLDocument object created from the XML or HTML contents of a URL-referenced source
	InitWithContentsOfURLOptionsError(url INSURL, mask NSXMLNodeOptions) (XMLDocument, error)
	// Initializes and returns an [NSXMLDocument] object created from an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object.
	InitWithDataOptionsError(data INSData, mask NSXMLNodeOptions) (XMLDocument, error)
	// Returns an [NSXMLDocument] object initialized with a single child, the root element.
	InitWithRootElement(element INSXMLElement) XMLDocument
	// Initializes and returns an [NSXMLDocument] object created from a string containing XML markup text.
	InitWithXMLStringOptionsError(string_ string, mask NSXMLNodeOptions) (XMLDocument, error)

	// Topic: Managing Document Attributes

	// Sets the character encoding of the receiver to `encoding`,
	CharacterEncoding() string
	SetCharacterEncoding(value string)
	// Sets the kind of output content for the receiver.
	DocumentContentKind() NSXMLDocumentContentKind
	SetDocumentContentKind(value NSXMLDocumentContentKind)
	// Returns an [XMLDTD](<doc://com.apple.foundation/documentation/Foundation/XMLDTD>) object representing the internal DTD associated with the receiver.
	DTD() INSXMLDTD
	SetDTD(value INSXMLDTD)
	// Sets a Boolean value that specifies whether the receiver represents a standalone XML document.
	Standalone() bool
	SetStandalone(value bool)
	// Returns the MIME type for the receiver.
	MIMEType() string
	SetMIMEType(value string)
	// Sets the version of the receiver’s XML.
	Version() string
	SetVersion(value string)

	// Topic: Managing the Root Element

	// Returns the root element of the receiver.
	RootElement() INSXMLElement
	// Set the root element of the receiver.
	SetRootElement(root INSXMLElement)

	// Topic: Adding and Removing Child Nodes

	// Adds a child node after the last of the receiver’s existing children.
	AddChild(child INSXMLNode)
	// Inserts a node object at specified position in the receiver’s array of children.
	InsertChildAtIndex(child INSXMLNode, index uint)
	// Inserts an array of children at a specified position in the receiver’s array of children.
	InsertChildrenAtIndex(children []NSXMLNode, index uint)
	// Removes the child node of the receiver located at a specified position in its array of children.
	RemoveChildAtIndex(index uint)
	// Replaces the child node of the receiver located at a specified position in its array of children with another node.
	ReplaceChildAtIndexWithNode(index uint, node INSXMLNode)

	// Topic: Transforming a Document Using XSLT

	// Applies the XSLT pattern rules and templates (specified as a data object) to the receiver and returns a document object containing transformed XML or HTML markup.
	ObjectByApplyingXSLTArgumentsError(xslt INSData, arguments INSDictionary) (objectivec.IObject, error)
	// Applies the XSLT pattern rules and templates (specified as a string) to the receiver and returns a document object containing transformed XML or HTML markup.
	ObjectByApplyingXSLTStringArgumentsError(xslt string, arguments INSDictionary) (objectivec.IObject, error)
	// Applies the XSLT pattern rules and templates located at a specified URL to the receiver and returns a document object containing transformed XML markup or an [NSData](<doc://com.apple.foundation/documentation/Foundation/NSData>) object containing plain text, RTF text, and so on.
	ObjectByApplyingXSLTAtURLArgumentsError(xsltURL INSURL, argument INSDictionary) (objectivec.IObject, error)

	// Topic: Writing a Document as XML Data

	// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
	XMLData() INSData
	// Returns the XML string representation of the receiver—that is, the entire document—encapsulated in a data object.
	XMLDataWithOptions(options NSXMLNodeOptions) INSData

	// Topic: Validating a Document

	// Validates the document against the governing schema and returns whether the document conforms to the schema.
	ValidateAndReturnError() (bool, error)
}

// Init initializes the instance.
func (x XMLDocument) Init() XMLDocument {
	rv := objc.Send[XMLDocument](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XMLDocument) Autorelease() XMLDocument {
	rv := objc.Send[XMLDocument](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLDocument creates a new XMLDocument instance.
func NewXMLDocument() XMLDocument {
	class := getXMLDocumentClass()
	rv := objc.Send[XMLDocument](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns an NSXMLDocument object created from the XML or
// HTML contents of a URL-referenced source
//
// url: An [NSURL] object specifying a URL source.
//
// mask: A bit mask for input options. You can specify multiple options by
// bit-OR’ing them. See Constants for a list of valid input options.
//
// # Return Value
// 
// An initialized [NSXMLDocument] object, or `nil` if initialization fails
// because of parsing errors or other reasons.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/init(contentsOf:options:)
func NewXMLDocumentWithContentsOfURLOptionsError(url INSURL, mask NSXMLNodeOptions) (XMLDocument, error) {
	var errorPtr objc.ID
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDocument{}, NSErrorFrom(errorPtr)
	}
	return XMLDocumentFromID(rv), nil
}

// Initializes and returns an [NSXMLDocument] object created from an [NSData]
// object.
//
// data: A data object with XML content.
//
// mask: A bit mask for input options. You can specify multiple options by
// bit-OR’ing them. See Constants for a list of valid input options.
//
// # Return Value
// 
// An initialized [NSXMLDocument] object, or `nil` if initialization fails
// because of parsing errors or other reasons.
//
// # Discussion
// 
// This method is the designated initializer for the [NSXMLDocument] class.
// 
// If you specify [NSXMLDocumentTidyXML] as one of the options, NSXMLDocument
// performs several clean-up operations on the document XML (such as removing
// leading tabs). It does respect the `space="preserve"` attribute when it
// attempts to tidy the XML.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/init(data:options:)
func NewXMLDocumentWithDataOptionsError(data INSData, mask NSXMLNodeOptions) (XMLDocument, error) {
	var errorPtr objc.ID
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:error:"), data, mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDocument{}, NSErrorFrom(errorPtr)
	}
	return XMLDocumentFromID(rv), nil
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
func NewXMLDocumentWithKind(kind NSXMLNodeKind) XMLDocument {
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:"), kind)
	return XMLDocumentFromID(rv)
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
func NewXMLDocumentWithKindOptions(kind NSXMLNodeKind, options NSXMLNodeOptions) XMLDocument {
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKind:options:"), kind, options)
	return XMLDocumentFromID(rv)
}

// Returns an [NSXMLDocument] object initialized with a single child, the root
// element.
//
// element: An [NSXMLElement] object representing an XML element.
//
// # Return Value
// 
// An initialized [NSXMLDocument] object, or `nil` if initialization fails for
// any reason.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/init(rootElement:)
func NewXMLDocumentWithRootElement(element INSXMLElement) XMLDocument {
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRootElement:"), element)
	return XMLDocumentFromID(rv)
}

// Initializes and returns an [NSXMLDocument] object created from a string
// containing XML markup text.
//
// string: A string object containing XML markup text.
//
// mask: A bit mask for input options. You can specify multiple options by
// bit-OR’ing them. See Constants for a list of valid input options.
//
// # Return Value
// 
// An initialized [NSXMLDocument] object, or `nil` if initialization fails
// because of parsing errors or other reasons.
//
// # Discussion
// 
// The encoding of the document is set to UTF-8.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/init(xmlString:options:)
func NewXMLDocumentWithXMLStringOptionsError(string_ string, mask NSXMLNodeOptions) (XMLDocument, error) {
	var errorPtr objc.ID
	instance := getXMLDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXMLString:options:error:"), objc.String(string_), mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDocument{}, NSErrorFrom(errorPtr)
	}
	return XMLDocumentFromID(rv), nil
}

// Initializes and returns an NSXMLDocument object created from the XML or
// HTML contents of a URL-referenced source
//
// url: An [NSURL] object specifying a URL source.
//
// mask: A bit mask for input options. You can specify multiple options by
// bit-OR’ing them. See Constants for a list of valid input options.
//
// # Return Value
// 
// An initialized [NSXMLDocument] object, or `nil` if initialization fails
// because of parsing errors or other reasons.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/init(contentsOf:options:)
func (x XMLDocument) InitWithContentsOfURLOptionsError(url INSURL, mask NSXMLNodeOptions) (XMLDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("initWithContentsOfURL:options:error:"), url, mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDocument{}, NSErrorFrom(errorPtr)
	}
	return NSXMLDocumentFromID(rv), nil

}
// Initializes and returns an [NSXMLDocument] object created from an [NSData]
// object.
//
// data: A data object with XML content.
//
// mask: A bit mask for input options. You can specify multiple options by
// bit-OR’ing them. See Constants for a list of valid input options.
//
// # Return Value
// 
// An initialized [NSXMLDocument] object, or `nil` if initialization fails
// because of parsing errors or other reasons.
//
// # Discussion
// 
// This method is the designated initializer for the [NSXMLDocument] class.
// 
// If you specify [NSXMLDocumentTidyXML] as one of the options, NSXMLDocument
// performs several clean-up operations on the document XML (such as removing
// leading tabs). It does respect the `space="preserve"` attribute when it
// attempts to tidy the XML.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/init(data:options:)
func (x XMLDocument) InitWithDataOptionsError(data INSData, mask NSXMLNodeOptions) (XMLDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("initWithData:options:error:"), data, mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDocument{}, NSErrorFrom(errorPtr)
	}
	return NSXMLDocumentFromID(rv), nil

}
// Returns an [NSXMLDocument] object initialized with a single child, the root
// element.
//
// element: An [NSXMLElement] object representing an XML element.
//
// # Return Value
// 
// An initialized [NSXMLDocument] object, or `nil` if initialization fails for
// any reason.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/init(rootElement:)
func (x XMLDocument) InitWithRootElement(element INSXMLElement) XMLDocument {
	rv := objc.Send[XMLDocument](x.ID, objc.Sel("initWithRootElement:"), element)
	return rv
}
// Initializes and returns an [NSXMLDocument] object created from a string
// containing XML markup text.
//
// string: A string object containing XML markup text.
//
// mask: A bit mask for input options. You can specify multiple options by
// bit-OR’ing them. See Constants for a list of valid input options.
//
// # Return Value
// 
// An initialized [NSXMLDocument] object, or `nil` if initialization fails
// because of parsing errors or other reasons.
//
// # Discussion
// 
// The encoding of the document is set to UTF-8.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/init(xmlString:options:)
func (x XMLDocument) InitWithXMLStringOptionsError(string_ string, mask NSXMLNodeOptions) (XMLDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("initWithXMLString:options:error:"), objc.String(string_), mask, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return XMLDocument{}, NSErrorFrom(errorPtr)
	}
	return NSXMLDocumentFromID(rv), nil

}
// Returns the root element of the receiver.
//
// # Return Value
// 
// The root element of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/rootElement()
func (x XMLDocument) RootElement() INSXMLElement {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("rootElement"))
	return NSXMLElementFromID(rv)
}
// Set the root element of the receiver.
//
// root: An [NSXMLNode] object that is to be the root element.
//
// # Discussion
// 
// As a side effect, this method removes all other children, including
// [NSXMLNode] objects representing comments and processing-instructions.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/setRootElement(_:)
func (x XMLDocument) SetRootElement(root INSXMLElement) {
	objc.Send[objc.ID](x.ID, objc.Sel("setRootElement:"), root)
}
// Adds a child node after the last of the receiver’s existing children.
//
// child: The [NSXMLNode] object to be added.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/addChild(_:)
func (x XMLDocument) AddChild(child INSXMLNode) {
	objc.Send[objc.ID](x.ID, objc.Sel("addChild:"), child)
}
// Inserts a node object at specified position in the receiver’s array of
// children.
//
// child: The [NSXMLNode] object to be inserted. The added node must be an
// [NSXMLNode] object representing a comment, processing instruction, or the
// root element.
//
// index: An integer specifying the index of the children array to insert `child`.
// The indexes of children after the new child are incremented. If `index` is
// less than zero or greater than the number of children, an out-of-bounds
// exception is raised.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/insertChild(_:at:)
func (x XMLDocument) InsertChildAtIndex(child INSXMLNode, index uint) {
	objc.Send[objc.ID](x.ID, objc.Sel("insertChild:atIndex:"), child, index)
}
// Inserts an array of children at a specified position in the receiver’s
// array of children.
//
// children: An array of [NSXMLNode] objects representing comments, processing
// instructions, or the root element.
//
// index: An integer identifying the location in the receiver’s children array for
// insertion. The indexes of children after the new child are increased by
// `[children count`]. If `index` is less than zero or greater than the number
// of children, an out-of-bounds exception is raised.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/insertChildren(_:at:)
func (x XMLDocument) InsertChildrenAtIndex(children []NSXMLNode, index uint) {
	objc.Send[objc.ID](x.ID, objc.Sel("insertChildren:atIndex:"), objectivec.IObjectSliceToNSArray(children), index)
}
// Removes the child node of the receiver located at a specified position in
// its array of children.
//
// index: An integer identifying the position of an child in the receiver’s array.
// If `index` is less than zero or greater than the number of children minus
// one, an out-of-bounds exception is raised.
//
// # Discussion
// 
// Subsequent children have their indexes decreased by one. The removed
// [NSXMLNode] object is autoreleased.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/removeChild(at:)
func (x XMLDocument) RemoveChildAtIndex(index uint) {
	objc.Send[objc.ID](x.ID, objc.Sel("removeChildAtIndex:"), index)
}
// Replaces the child node of the receiver located at a specified position in
// its array of children with another node.
//
// index: An integer identifying a position in the receiver’s array of children. If
// `index` is less than zero or greater than the number of children minus one,
// an out-of-bounds exception is raised.
//
// node: An [NSXMLNode] object to replace the one at `index`; it must represent a
// comment, a processing instruction, or the root element.
//
// # Discussion
// 
// The removed [NSXMLNode] object is autoreleased.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/replaceChild(at:with:)
func (x XMLDocument) ReplaceChildAtIndexWithNode(index uint, node INSXMLNode) {
	objc.Send[objc.ID](x.ID, objc.Sel("replaceChildAtIndex:withNode:"), index, node)
}
// Applies the XSLT pattern rules and templates (specified as a data object)
// to the receiver and returns a document object containing transformed XML or
// HTML markup.
//
// xslt: A data object containing the XSLT pattern rules and templates.
//
// arguments: A dictionary containing [NSString] key-value pairs that are passed as
// runtime parameters to the XSLT processor. Pass in `nil` if you have no
// parameters to pass.
//
// # Return Value
// 
// Depending on intended output, the method returns an [NSXMLDocument] object
// an [NSData] data containing transformed XML or HTML markup. If the message
// is supposed to create plain text or RTF, then an [NSData] object is
// returned, otherwise an XML document object. The method returns `nil` if
// XSLT processing did not succeed.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/object(byApplyingXSLT:arguments:)
func (x XMLDocument) ObjectByApplyingXSLTArgumentsError(xslt INSData, arguments INSDictionary) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("objectByApplyingXSLT:arguments:error:"), xslt, arguments, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// Applies the XSLT pattern rules and templates (specified as a string) to the
// receiver and returns a document object containing transformed XML or HTML
// markup.
//
// xslt: A string object containing the XSLT pattern rules and templates.
//
// arguments: A dictionary containing [NSString] key-value pairs that are passed as
// runtime parameters to the XSLT processor. Pass in `nil` if you have no
// parameters to pass.
//
// # Return Value
// 
// Depending on intended output, the method returns an [NSXMLDocument] object
// an [NSData] data containing transformed XML or HTML markup. If the message
// is supposed to create plain text or RTF, then an [NSData] object is
// returned, otherwise an XML document object. The method returns `nil` if
// XSLT processing did not succeed.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/object(byApplyingXSLTString:arguments:)
func (x XMLDocument) ObjectByApplyingXSLTStringArgumentsError(xslt string, arguments INSDictionary) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("objectByApplyingXSLTString:arguments:error:"), objc.String(xslt), arguments, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// Applies the XSLT pattern rules and templates located at a specified URL to
// the receiver and returns a document object containing transformed XML
// markup or an [NSData] object containing plain text, RTF text, and so on.
//
// xsltURL: An [NSURL] object specifying a valid URL.
//
// argument: A dictionary containing [NSString] key-value pairs that are passed as
// runtime parameters to the XSLT processor. Pass in `nil` if you have no
// parameters to pass.
//
// # Return Value
// 
// Depending on intended output, the returns an [NSXMLDocument] object an
// [NSData] data containing transformed XML or HTML markup. If the message is
// supposed to create plain text or RTF, then an [NSData] object is returned,
// otherwise an XML document object. The method returns `nil` if XSLT
// processing did not succeed.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/objectByApplyingXSLT(at:arguments:)
func (x XMLDocument) ObjectByApplyingXSLTAtURLArgumentsError(xsltURL INSURL, argument INSDictionary) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](x.ID, objc.Sel("objectByApplyingXSLTAtURL:arguments:error:"), xsltURL, argument, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// Returns the XML string representation of the receiver—that is, the entire
// document—encapsulated in a data object.
//
// options: One or more options (bit-OR’d if multiple) to affect the output of the
// document; see Constants for the valid output options.
//
// # Discussion
// 
// The encoding used is based on the value returned from [CharacterEncoding].
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/xmlData(options:)
func (x XMLDocument) XMLDataWithOptions(options NSXMLNodeOptions) INSData {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("XMLDataWithOptions:"), options)
	return NSDataFromID(rv)
}
// Validates the document against the governing schema and returns whether the
// document conforms to the schema.
//
// # Discussion
// 
// The constants indicating the kind of validation errors are emitted by the
// underlying parser; see `NSXMLParser.H()` for most of these constants. If
// the schema is defined with a DTD, this method uses the [NSXMLDTD] object
// set for the receiver for validation. If the schema is based on XML Schema,
// the method uses the URL specified as the value of the `schemaLocation`
// attribute of the root element.
// 
// You can validate an XML document when it is first processed by specifying
// the [NSXMLDocumentValidate] option when you initialize an [NSXMLDocument]
// object with the [InitWithContentsOfURLOptionsError],
// [InitWithDataOptionsError], or [InitWithXMLStringOptionsError] methods.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/validate()
func (x XMLDocument) ValidateAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](x.ID, objc.Sel("validateAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Overridden by subclasses to substitute a custom class for an NSXML class
// that the parser uses to create node instances.
//
// cls: A [Class] object identifying an NSXML class that is to be replaced by your
// custom class.
//
// # Return Value
// 
// The substituted class.
//
// # Discussion
// 
// For example, if you have a custom subclass of [NSXMLElement] that you want
// to be used in place of [NSXMLElement], you would make the following
// override:
// 
// This method is invoked before a document is parsed. The substituted class
// must be a subclass of [NSXMLNode], [NSXMLDocument], [NSXMLElement],
// [NSXMLDTD], or [NSXMLDTDNode].
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/replacementClass(for:)
func (_XMLDocumentClass XMLDocumentClass) ReplacementClassForClass(cls objc.Class) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_XMLDocumentClass.class), objc.Sel("replacementClassForClass:"), cls)
	return rv
}

// Sets the character encoding of the receiver to `encoding`,
//
// # Discussion
// 
// Typically the encoding is specified in the XML declaration of a document
// that is processed, but it can be set at any time. If the specified encoding
// does not match the actual encoding, parsing of the document might fail.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/characterEncoding
func (x XMLDocument) CharacterEncoding() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("characterEncoding"))
	return NSStringFromID(rv).String()
}
func (x XMLDocument) SetCharacterEncoding(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setCharacterEncoding:"), objc.String(value))
}
// Sets the kind of output content for the receiver.
//
// # Discussion
// 
// Most of the differences among document-content kind have to do with the
// handling of content-less tags such as ``.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/documentContentKind
func (x XMLDocument) DocumentContentKind() NSXMLDocumentContentKind {
	rv := objc.Send[NSXMLDocumentContentKind](x.ID, objc.Sel("documentContentKind"))
	return NSXMLDocumentContentKind(rv)
}
func (x XMLDocument) SetDocumentContentKind(value NSXMLDocumentContentKind) {
	objc.Send[struct{}](x.ID, objc.Sel("setDocumentContentKind:"), value)
}
// Returns an [NSXMLDTD] object representing the internal DTD associated with
// the receiver.
//
// # Return Value
// 
// An [NSXMLDTD] object representing the internal DTD associated with the
// receiver or `nil` if no DTD has been associated.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/dtd
func (x XMLDocument) DTD() INSXMLDTD {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("DTD"))
	return NSXMLDTDFromID(objc.ID(rv))
}
func (x XMLDocument) SetDTD(value INSXMLDTD) {
	objc.Send[struct{}](x.ID, objc.Sel("setDTD:"), value)
}
// Sets a Boolean value that specifies whether the receiver represents a
// standalone XML document.
//
// # Discussion
// 
// A standalone document does not have an external DTD associated with it.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/isStandalone
func (x XMLDocument) Standalone() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("isStandalone"))
	return rv
}
func (x XMLDocument) SetStandalone(value bool) {
	objc.Send[struct{}](x.ID, objc.Sel("setStandalone:"), value)
}
// Returns the MIME type for the receiver.
//
// # Return Value
// 
// The MIME type for the receiver (for example, “text/xml”).
// 
// # Discussion
// 
// MIME types are assigned by IANA (see
// [http://www.iana.org/assignments/media-types/index.html]).
//
// [http://www.iana.org/assignments/media-types/index.html]: http://www.iana.org/assignments/media-types/index.html
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/mimeType
func (x XMLDocument) MIMEType() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("MIMEType"))
	return NSStringFromID(rv).String()
}
func (x XMLDocument) SetMIMEType(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setMIMEType:"), objc.String(value))
}
// Sets the version of the receiver’s XML.
//
// # Discussion
// 
// Currently, the version should be either “1.0 “or “1.1”.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/version
func (x XMLDocument) Version() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("version"))
	return NSStringFromID(rv).String()
}
func (x XMLDocument) SetVersion(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setVersion:"), objc.String(value))
}
// Returns the XML string representation of the receiver—that is, the entire
// document—encapsulated in a data object.
//
// # Discussion
// 
// This method invokes [XMLDataWithOptions] with an option of
// [NSXMLNodeOptionsNone]. The encoding used is based on the value returned
// from [CharacterEncoding] or UTF-8 if no valid encoding is returned by that
// method.
//
// See: https://developer.apple.com/documentation/Foundation/XMLDocument/xmlData
func (x XMLDocument) XMLData() INSData {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("XMLData"))
	return NSDataFromID(objc.ID(rv))
}

