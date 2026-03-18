// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XMLParser] class.
var (
	_XMLParserClass     XMLParserClass
	_XMLParserClassOnce sync.Once
)

func getXMLParserClass() XMLParserClass {
	_XMLParserClassOnce.Do(func() {
		_XMLParserClass = XMLParserClass{class: objc.GetClass("NSXMLParser")}
	})
	return _XMLParserClass
}

// GetXMLParserClass returns the class object for NSXMLParser.
func GetXMLParserClass() XMLParserClass {
	return getXMLParserClass()
}

type XMLParserClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (xc XMLParserClass) Alloc() XMLParser {
	rv := objc.Send[XMLParser](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// An event driven parser of XML documents (including DTD declarations).
//
// # Overview
// 
// An [NSXMLParser] notifies its delegate about the items (elements,
// attributes, CDATA blocks, comments, and so on) that it encounters as it
// processes an XML document. It does not itself do anything with those parsed
// items except report them. It also reports parsing errors. For convenience,
// an [NSXMLParser] object in the following descriptions is sometimes referred
// to as a parser object. Unless used in a callback, the [NSXMLParser] is a
// thread-safe class as long as any given instance is only used in one thread.
//
// # Initializing a Parser Object
//
//   - [XMLParser.InitWithContentsOfURL]: Initializes a parser with the XML content referenced by the given URL.
//   - [XMLParser.InitWithData]: Initializes a parser with the XML contents encapsulated in a given data object.
//   - [XMLParser.InitWithStream]: Initializes a parser with the XML contents from the specified stream and parses it.
//
// # Managing Delegates
//
//   - [XMLParser.Delegate]: A delegate object that receives messages about the parsing process.
//   - [XMLParser.SetDelegate]
//
// # Managing Parser Behavior
//
//   - [XMLParser.ShouldProcessNamespaces]: A Boolean value that determines whether the parser reports the namespaces and qualified names of elements.
//   - [XMLParser.SetShouldProcessNamespaces]
//   - [XMLParser.ShouldReportNamespacePrefixes]: A Boolean value that determines whether the parser reports the prefixes indicating the scope of namespace declarations.
//   - [XMLParser.SetShouldReportNamespacePrefixes]
//   - [XMLParser.ShouldResolveExternalEntities]: A Boolean value that determines whether the parser reports declarations of external entities.
//   - [XMLParser.SetShouldResolveExternalEntities]
//
// # Parsing
//
//   - [XMLParser.Parse]: Starts the event-driven parsing operation.
//   - [XMLParser.AbortParsing]: Stops the parser object.
//   - [XMLParser.ParserError]: An [NSError](<doc://com.apple.foundation/documentation/Foundation/NSError>) object from which you can obtain information about a parsing error.
//
// # Obtaining Parser State
//
//   - [XMLParser.ColumnNumber]: The column number of the XML document being processed by the parser.
//   - [XMLParser.LineNumber]: The line number of the XML document being processed by the parser.
//   - [XMLParser.PublicID]: The public identifier of the external entity referenced in the XML document.
//   - [XMLParser.SystemID]: The system identifier of the external entity referenced in the XML document.
//
// # Instance Properties
//
//   - [XMLParser.AllowedExternalEntityURLs]
//   - [XMLParser.SetAllowedExternalEntityURLs]
//   - [XMLParser.ExternalEntityResolvingPolicy]
//   - [XMLParser.SetExternalEntityResolvingPolicy]
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser
type XMLParser struct {
	objectivec.Object
}

// XMLParserFromID constructs a [XMLParser] from an objc.ID.
//
// An event driven parser of XML documents (including DTD declarations).
func XMLParserFromID(id objc.ID) XMLParser {
	return NSXMLParser{objectivec.Object{ID: id}}
}

// NSXMLParserFromID is an alias for [XMLParserFromID] for cross-framework compatibility.
func NSXMLParserFromID(id objc.ID) XMLParser { return XMLParserFromID(id) }
// NOTE: XMLParser adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [XMLParser] class.
//
// # Initializing a Parser Object
//
//   - [IXMLParser.InitWithContentsOfURL]: Initializes a parser with the XML content referenced by the given URL.
//   - [IXMLParser.InitWithData]: Initializes a parser with the XML contents encapsulated in a given data object.
//   - [IXMLParser.InitWithStream]: Initializes a parser with the XML contents from the specified stream and parses it.
//
// # Managing Delegates
//
//   - [IXMLParser.Delegate]: A delegate object that receives messages about the parsing process.
//   - [IXMLParser.SetDelegate]
//
// # Managing Parser Behavior
//
//   - [IXMLParser.ShouldProcessNamespaces]: A Boolean value that determines whether the parser reports the namespaces and qualified names of elements.
//   - [IXMLParser.SetShouldProcessNamespaces]
//   - [IXMLParser.ShouldReportNamespacePrefixes]: A Boolean value that determines whether the parser reports the prefixes indicating the scope of namespace declarations.
//   - [IXMLParser.SetShouldReportNamespacePrefixes]
//   - [IXMLParser.ShouldResolveExternalEntities]: A Boolean value that determines whether the parser reports declarations of external entities.
//   - [IXMLParser.SetShouldResolveExternalEntities]
//
// # Parsing
//
//   - [IXMLParser.Parse]: Starts the event-driven parsing operation.
//   - [IXMLParser.AbortParsing]: Stops the parser object.
//   - [IXMLParser.ParserError]: An [NSError](<doc://com.apple.foundation/documentation/Foundation/NSError>) object from which you can obtain information about a parsing error.
//
// # Obtaining Parser State
//
//   - [IXMLParser.ColumnNumber]: The column number of the XML document being processed by the parser.
//   - [IXMLParser.LineNumber]: The line number of the XML document being processed by the parser.
//   - [IXMLParser.PublicID]: The public identifier of the external entity referenced in the XML document.
//   - [IXMLParser.SystemID]: The system identifier of the external entity referenced in the XML document.
//
// # Instance Properties
//
//   - [IXMLParser.AllowedExternalEntityURLs]
//   - [IXMLParser.SetAllowedExternalEntityURLs]
//   - [IXMLParser.ExternalEntityResolvingPolicy]
//   - [IXMLParser.SetExternalEntityResolvingPolicy]
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser
type IXMLParser interface {
	objectivec.IObject

	// Topic: Initializing a Parser Object

	// Initializes a parser with the XML content referenced by the given URL.
	InitWithContentsOfURL(url INSURL) XMLParser
	// Initializes a parser with the XML contents encapsulated in a given data object.
	InitWithData(data INSData) XMLParser
	// Initializes a parser with the XML contents from the specified stream and parses it.
	InitWithStream(stream INSInputStream) XMLParser

	// Topic: Managing Delegates

	// A delegate object that receives messages about the parsing process.
	Delegate() NSXMLParserDelegate
	SetDelegate(value NSXMLParserDelegate)

	// Topic: Managing Parser Behavior

	// A Boolean value that determines whether the parser reports the namespaces and qualified names of elements.
	ShouldProcessNamespaces() bool
	SetShouldProcessNamespaces(value bool)
	// A Boolean value that determines whether the parser reports the prefixes indicating the scope of namespace declarations.
	ShouldReportNamespacePrefixes() bool
	SetShouldReportNamespacePrefixes(value bool)
	// A Boolean value that determines whether the parser reports declarations of external entities.
	ShouldResolveExternalEntities() bool
	SetShouldResolveExternalEntities(value bool)

	// Topic: Parsing

	// Starts the event-driven parsing operation.
	Parse() bool
	// Stops the parser object.
	AbortParsing()
	// An [NSError](<doc://com.apple.foundation/documentation/Foundation/NSError>) object from which you can obtain information about a parsing error.
	ParserError() INSError

	// Topic: Obtaining Parser State

	// The column number of the XML document being processed by the parser.
	ColumnNumber() int
	// The line number of the XML document being processed by the parser.
	LineNumber() int
	// The public identifier of the external entity referenced in the XML document.
	PublicID() string
	// The system identifier of the external entity referenced in the XML document.
	SystemID() string

	// Topic: Instance Properties

	AllowedExternalEntityURLs() INSSet
	SetAllowedExternalEntityURLs(value INSSet)
	ExternalEntityResolvingPolicy() NSXMLParserExternalEntityResolvingPolicy
	SetExternalEntityResolvingPolicy(value NSXMLParserExternalEntityResolvingPolicy)
}

// Init initializes the instance.
func (x XMLParser) Init() XMLParser {
	rv := objc.Send[XMLParser](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XMLParser) Autorelease() XMLParser {
	rv := objc.Send[XMLParser](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXMLParser creates a new XMLParser instance.
func NewXMLParser() XMLParser {
	class := getXMLParserClass()
	rv := objc.Send[XMLParser](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a parser with the XML content referenced by the given URL.
//
// url: An [NSURL] object specifying a URL. The URL must be fully qualified and
// refer to a scheme that is supported by the [NSURL] class.
//
// # Return Value
// 
// An initialized [NSXMLParser] object or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/init(contentsOf:)
func NewXMLParserWithContentsOfURL(url INSURL) XMLParser {
	instance := getXMLParserClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return XMLParserFromID(rv)
}

// Initializes a parser with the XML contents encapsulated in a given data
// object.
//
// data: An [NSData] object containing XML markup.
//
// # Return Value
// 
// An initialized [NSXMLParser] object or `nil` if an error occurs.
//
// # Discussion
// 
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/init(data:)
func NewXMLParserWithData(data INSData) XMLParser {
	instance := getXMLParserClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return XMLParserFromID(rv)
}

// Initializes a parser with the XML contents from the specified stream and
// parses it.
//
// stream: The input stream. The content is incrementally loaded from the specified
// stream and parsed. The [NSXMLParser] will open the stream, and
// synchronously read from it without scheduling it.
//
// # Return Value
// 
// An initialized [NSXMLParser] object or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/init(stream:)
func NewXMLParserWithStream(stream INSInputStream) XMLParser {
	instance := getXMLParserClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStream:"), stream)
	return XMLParserFromID(rv)
}

// Initializes a parser with the XML content referenced by the given URL.
//
// url: An [NSURL] object specifying a URL. The URL must be fully qualified and
// refer to a scheme that is supported by the [NSURL] class.
//
// # Return Value
// 
// An initialized [NSXMLParser] object or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/init(contentsOf:)
func (x XMLParser) InitWithContentsOfURL(url INSURL) XMLParser {
	rv := objc.Send[XMLParser](x.ID, objc.Sel("initWithContentsOfURL:"), url)
	return rv
}

// Initializes a parser with the XML contents encapsulated in a given data
// object.
//
// data: An [NSData] object containing XML markup.
//
// # Return Value
// 
// An initialized [NSXMLParser] object or `nil` if an error occurs.
//
// # Discussion
// 
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/init(data:)
func (x XMLParser) InitWithData(data INSData) XMLParser {
	rv := objc.Send[XMLParser](x.ID, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes a parser with the XML contents from the specified stream and
// parses it.
//
// stream: The input stream. The content is incrementally loaded from the specified
// stream and parsed. The [NSXMLParser] will open the stream, and
// synchronously read from it without scheduling it.
//
// # Return Value
// 
// An initialized [NSXMLParser] object or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/init(stream:)
func (x XMLParser) InitWithStream(stream INSInputStream) XMLParser {
	rv := objc.Send[XMLParser](x.ID, objc.Sel("initWithStream:"), stream)
	return rv
}

// Starts the event-driven parsing operation.
//
// # Return Value
// 
// [true] if the parsing operation succeeds; [false] if an error occurs or if
// the parsing operation aborts.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/parse()
func (x XMLParser) Parse() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("parse"))
	return rv
}

// Stops the parser object.
//
// # Discussion
// 
// If you invoke this method, the delegate, if it implements
// [ParserParseErrorOccurred], is informed of the cancelled parsing operation.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/abortParsing()
func (x XMLParser) AbortParsing() {
	objc.Send[objc.ID](x.ID, objc.Sel("abortParsing"))
}

// A delegate object that receives messages about the parsing process.
//
// # Discussion
// 
// For methods to be implemented by the delegate, see [NSXMLParserDelegate].
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/delegate
func (x XMLParser) Delegate() NSXMLParserDelegate {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("delegate"))
	return NSXMLParserDelegateObjectFromID(rv)
}
func (x XMLParser) SetDelegate(value NSXMLParserDelegate) {
	objc.Send[struct{}](x.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that determines whether the parser reports the namespaces
// and qualified names of elements.
//
// # Discussion
// 
// [true] if the parser reports namespace and qualified name, [false]
// otherwise.
// 
// The parser reports element names with the delegate methods
// [ParserDidStartElementNamespaceURIQualifiedNameAttributes] and
// [ParserDidEndElementNamespaceURIQualifiedName].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/shouldProcessNamespaces
func (x XMLParser) ShouldProcessNamespaces() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("shouldProcessNamespaces"))
	return rv
}
func (x XMLParser) SetShouldProcessNamespaces(value bool) {
	objc.Send[struct{}](x.ID, objc.Sel("setShouldProcessNamespaces:"), value)
}

// A Boolean value that determines whether the parser reports the prefixes
// indicating the scope of namespace declarations.
//
// # Discussion
// 
// [true] if the parser reports the scope of namespace declarations, [false]
// otherwise. The default value is [false].
// 
// The parser reports prefixes with the delegate methods
// [ParserDidStartMappingPrefixToURI] and [ParserDidEndMappingPrefix].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/shouldReportNamespacePrefixes
func (x XMLParser) ShouldReportNamespacePrefixes() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("shouldReportNamespacePrefixes"))
	return rv
}
func (x XMLParser) SetShouldReportNamespacePrefixes(value bool) {
	objc.Send[struct{}](x.ID, objc.Sel("setShouldReportNamespacePrefixes:"), value)
}

// A Boolean value that determines whether the parser reports declarations of
// external entities.
//
// # Discussion
// 
// [true] if the parser reports declarations of external entities, [false]
// otherwise. The default value is [false]. If you set this property to
// [true], you may cause other I/O operations, either network-based or
// disk-based, to load the external DTD.
// 
// The parser reports declarations of external entities with the delegate
// method [ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/shouldResolveExternalEntities
func (x XMLParser) ShouldResolveExternalEntities() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("shouldResolveExternalEntities"))
	return rv
}
func (x XMLParser) SetShouldResolveExternalEntities(value bool) {
	objc.Send[struct{}](x.ID, objc.Sel("setShouldResolveExternalEntities:"), value)
}

// An [NSError] object from which you can obtain information about a parsing
// error.
//
// # Discussion
// 
// You may access this property after a parsing operation abnormally
// terminates to determine the cause of error.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/parserError
func (x XMLParser) ParserError() INSError {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("parserError"))
	return NSErrorFromID(objc.ID(rv))
}

// The column number of the XML document being processed by the parser.
//
// # Discussion
// 
// The column refers to the nesting level of the XML elements in the document.
// You may access this property once a parsing operation has begun or after an
// error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/columnNumber
func (x XMLParser) ColumnNumber() int {
	rv := objc.Send[int](x.ID, objc.Sel("columnNumber"))
	return rv
}

// The line number of the XML document being processed by the parser.
//
// # Discussion
// 
// You may access this property once a parsing operation has begun or after an
// error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/lineNumber
func (x XMLParser) LineNumber() int {
	rv := objc.Send[int](x.ID, objc.Sel("lineNumber"))
	return rv
}

// The public identifier of the external entity referenced in the XML
// document.
//
// # Discussion
// 
// You may access this property once a parsing operation has begun or after an
// error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/publicID
func (x XMLParser) PublicID() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("publicID"))
	return NSStringFromID(rv).String()
}

// The system identifier of the external entity referenced in the XML
// document.
//
// # Discussion
// 
// You may access this property once a parsing operation has begun or after an
// error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParser/systemID
func (x XMLParser) SystemID() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("systemID"))
	return NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Foundation/XMLParser/allowedExternalEntityURLs
func (x XMLParser) AllowedExternalEntityURLs() INSSet {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("allowedExternalEntityURLs"))
	return NSSetFromID(objc.ID(rv))
}
func (x XMLParser) SetAllowedExternalEntityURLs(value INSSet) {
	objc.Send[struct{}](x.ID, objc.Sel("setAllowedExternalEntityURLs:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/XMLParser/externalEntityResolvingPolicy-swift.property
func (x XMLParser) ExternalEntityResolvingPolicy() NSXMLParserExternalEntityResolvingPolicy {
	rv := objc.Send[NSXMLParserExternalEntityResolvingPolicy](x.ID, objc.Sel("externalEntityResolvingPolicy"))
	return NSXMLParserExternalEntityResolvingPolicy(rv)
}
func (x XMLParser) SetExternalEntityResolvingPolicy(value NSXMLParserExternalEntityResolvingPolicy) {
	objc.Send[struct{}](x.ID, objc.Sel("setExternalEntityResolvingPolicy:"), value)
}

