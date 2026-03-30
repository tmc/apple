// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The interface an XML parser uses to inform its delegate about the content of the parsed document.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate
type NSXMLParserDelegate interface {
	objectivec.IObject
}

// NSXMLParserDelegateObject wraps an existing Objective-C object that conforms to the NSXMLParserDelegate protocol.
type NSXMLParserDelegateObject struct {
	objectivec.Object
}

func (o NSXMLParserDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSXMLParserDelegateObjectFromID constructs a [NSXMLParserDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSXMLParserDelegateObjectFromID(id objc.ID) NSXMLParserDelegateObject {
	return NSXMLParserDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sent by the parser object to the delegate when it begins parsing a
// document.
//
// parser: A parser object.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parserDidStartDocument(_:)
func (o NSXMLParserDelegateObject) ParserDidStartDocument(parser INSXMLParser) {
	objc.Send[struct{}](o.ID, objc.Sel("parserDidStartDocument:"), parser)
}

// Sent by the parser object to the delegate when it has successfully
// completed parsing.
//
// parser: A parser object.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parserDidEndDocument(_:)
func (o NSXMLParserDelegateObject) ParserDidEndDocument(parser INSXMLParser) {
	objc.Send[struct{}](o.ID, objc.Sel("parserDidEndDocument:"), parser)
}

// Sent by a parser object to its delegate when it encounters a start tag for
// a given element.
//
// parser: A parser object.
//
// elementName: A string that is the name of an element (in its start tag).
//
// namespaceURI: If namespace processing is turned on, contains the URI for the current
// namespace as a string object.
//
// qName: If namespace processing is turned on, contains the qualified name for the
// current namespace as a string object.
//
// attributeDict: A dictionary that contains any attributes associated with the element. Keys
// are the names of attributes, and values are attribute values.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:didStartElement:namespaceURI:qualifiedName:attributes:)
func (o NSXMLParserDelegateObject) ParserDidStartElementNamespaceURIQualifiedNameAttributes(parser INSXMLParser, elementName string, namespaceURI string, qName string, attributeDict INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:didStartElement:namespaceURI:qualifiedName:attributes:"), parser, objc.String(elementName), objc.String(namespaceURI), objc.String(qName), attributeDict)
}

// Sent by a parser object to its delegate when it encounters an end tag for a
// specific element.
//
// parser: A parser object.
//
// elementName: A string that is the name of an element (in its end tag).
//
// namespaceURI: If namespace processing is turned on, contains the URI for the current
// namespace as a string object.
//
// qName: If namespace processing is turned on, contains the qualified name for the
// current namespace as a string object.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:didEndElement:namespaceURI:qualifiedName:)
func (o NSXMLParserDelegateObject) ParserDidEndElementNamespaceURIQualifiedName(parser INSXMLParser, elementName string, namespaceURI string, qName string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:didEndElement:namespaceURI:qualifiedName:"), parser, objc.String(elementName), objc.String(namespaceURI), objc.String(qName))
}

// Sent by a parser object to its delegate the first time it encounters a
// given namespace prefix, which is mapped to a URI.
//
// parser: A parser object.
//
// prefix: A string that is a namespace prefix.
//
// namespaceURI: A string that specifies a namespace URI.
//
// # Discussion
//
// The parser object sends this message only when namespace-prefix reporting
// is turned on through the [ShouldReportNamespacePrefixes] method.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:didStartMappingPrefix:toURI:)
func (o NSXMLParserDelegateObject) ParserDidStartMappingPrefixToURI(parser INSXMLParser, prefix string, namespaceURI string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:didStartMappingPrefix:toURI:"), parser, objc.String(prefix), objc.String(namespaceURI))
}

// Sent by a parser object to its delegate when a given namespace prefix goes
// out of scope.
//
// parser: A parser object.
//
// prefix: A string that is a namespace prefix.
//
// # Discussion
//
// The parser sends this message only when namespace-prefix reporting is
// turned on through the [ShouldReportNamespacePrefixes] method.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:didEndMappingPrefix:)
func (o NSXMLParserDelegateObject) ParserDidEndMappingPrefix(parser INSXMLParser, prefix string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:didEndMappingPrefix:"), parser, objc.String(prefix))
}

// Sent by a parser object to its delegate when it encounters a given external
// entity with a specific system ID.
//
// parser: A parser object.
//
// name: A string that specifies the external name of an entity.
//
// systemID: A string that specifies the system ID for the external entity.
//
// # Return Value
//
// An [NSData] object that contains the resolution of the given external
// entity.
//
// # Discussion
//
// The delegate can resolve the external entity (for example, locating and
// reading an externally declared DTD) and provide the result to the parser
// object as an [NSData] object.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:resolveExternalEntityName:systemID:)
func (o NSXMLParserDelegateObject) ParserResolveExternalEntityNameSystemID(parser INSXMLParser, name string, systemID string) INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("parser:resolveExternalEntityName:systemID:"), parser, objc.String(name), objc.String(systemID))
	return NSDataFromID(rv)
}

// Sent by a parser object to its delegate when it encounters a fatal error.
//
// parser: A parser object.
//
// parseError: An [NSError] object describing the parsing error that occurred.
//
// # Discussion
//
// When this method is invoked, parsing is stopped. For further information
// about the error, you can query `parseError` or you can send the `parser` a
// [ParserError] message. You can also send the parser [LineNumber] and
// [ColumnNumber] messages to further isolate where the error occurred.
// Typically you implement this method to display information about the error
// to the user.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:parseErrorOccurred:)
func (o NSXMLParserDelegateObject) ParserParseErrorOccurred(parser INSXMLParser, parseError INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:parseErrorOccurred:"), parser, parseError)
}

// Sent by a parser object to its delegate when it encounters a fatal
// validation error. [NSXMLParser] currently does not invoke this method and
// does not perform validation.
//
// parser: A parser object.
//
// validationError: An [NSError] object describing the validation error that occurred.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:validationErrorOccurred:)
func (o NSXMLParserDelegateObject) ParserValidationErrorOccurred(parser INSXMLParser, validationError INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:validationErrorOccurred:"), parser, validationError)
}

// Sent by a parser object to provide its delegate with a string representing
// all or part of the characters of the current element.
//
// parser: A parser object.
//
// string: A string representing the complete or partial textual content of the
// current element.
//
// # Discussion
//
// The parser object may send the delegate several [ParserFoundCharacters]
// messages to report the characters of an element. Because `string` may be
// only part of the total character content for the current element, you
// should append it to the current accumulation of characters until the
// element changes.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundCharacters:)
func (o NSXMLParserDelegateObject) ParserFoundCharacters(parser INSXMLParser, string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundCharacters:"), parser, objc.String(string_))
}

// Reported by a parser object to provide its delegate with a string
// representing all or part of the ignorable whitespace characters of the
// current element.
//
// parser: A parser object.
//
// whitespaceString: A string representing all or part of the ignorable whitespace characters of
// the current element.
//
// # Discussion
//
// All the whitespace characters of the element (including carriage returns,
// tabs, and new-line characters) may not be provided through an individual
// invocation of this method. The parser may send the delegate several
// [ParserFoundIgnorableWhitespace] messages to report the whitespace
// characters of an element. You should append the characters in each
// invocation to the current accumulation of characters.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundIgnorableWhitespace:)
func (o NSXMLParserDelegateObject) ParserFoundIgnorableWhitespace(parser INSXMLParser, whitespaceString string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundIgnorableWhitespace:"), parser, objc.String(whitespaceString))
}

// Sent by a parser object to its delegate when it encounters a processing
// instruction.
//
// parser: A parser object.
//
// target: A string representing the target of a processing instruction.
//
// data: A string representing the data for a processing instruction.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundProcessingInstructionWithTarget:data:)
func (o NSXMLParserDelegateObject) ParserFoundProcessingInstructionWithTargetData(parser INSXMLParser, target string, data string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundProcessingInstructionWithTarget:data:"), parser, objc.String(target), objc.String(data))
}

// Sent by a parser object to its delegate when it encounters a comment in the
// XML.
//
// parser: An [NSXMLParser] object parsing XML.
//
// comment: A string that is a the content of a comment in the XML.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundComment:)
func (o NSXMLParserDelegateObject) ParserFoundComment(parser INSXMLParser, comment string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundComment:"), parser, objc.String(comment))
}

// Sent by a parser object to its delegate when it encounters a CDATA block.
//
// parser: An [NSXMLParser] object parsing XML.
//
// CDATABlock: A data object containing a block of CDATA.
//
// # Discussion
//
// Through this method the parser object passes the contents of the block to
// its delegate in an [NSData] object. The CDATA block is character data that
// is ignored by the parser. The encoding of the character data is UTF-8. To
// convert the data object to a string object, use the [NSString] method
// [InitWithDataEncoding].
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundCDATA:)
func (o NSXMLParserDelegateObject) ParserFoundCDATA(parser INSXMLParser, CDATABlock INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundCDATA:"), parser, CDATABlock)
}

// Sent by a parser object to its delegate when it encounters a declaration of
// an attribute that is associated with a specific element.
//
// parser: An [NSXMLParser] object parsing XML.
//
// attributeName: A string that is the name of an attribute.
//
// elementName: A string that is the name of an element that has the attribute
// `attributeName`.
//
// type: A string, such as “ENTITY”, “NOTATION”, or “ID”, that indicates
// the type of the attribute.
//
// defaultValue: A string that specifies the default value of the attribute.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundAttributeDeclarationWithName:forElement:type:defaultValue:)
func (o NSXMLParserDelegateObject) ParserFoundAttributeDeclarationWithNameForElementTypeDefaultValue(parser INSXMLParser, attributeName string, elementName string, type_ string, defaultValue string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundAttributeDeclarationWithName:forElement:type:defaultValue:"), parser, objc.String(attributeName), objc.String(elementName), objc.String(type_), objc.String(defaultValue))
}

// Sent by a parser object to its delegate when it encounters a declaration of
// an element with a given model.
//
// parser: An [NSXMLParser] object parsing XML.
//
// elementName: A string that is the name of an element.
//
// model: A string that specifies a model for `elementName`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundElementDeclarationWithName:model:)
func (o NSXMLParserDelegateObject) ParserFoundElementDeclarationWithNameModel(parser INSXMLParser, elementName string, model string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundElementDeclarationWithName:model:"), parser, objc.String(elementName), objc.String(model))
}

// Sent by a parser object to its delegate when it encounters an external
// entity declaration.
//
// parser: An [NSXMLParser] object parsing XML.
//
// name: A string that is the name of an entity.
//
// publicID: A string that specifies the public ID associated with `entityName`.
//
// systemID: A string that specifies the system ID associated with `entityName`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundExternalEntityDeclarationWithName:publicID:systemID:)
func (o NSXMLParserDelegateObject) ParserFoundExternalEntityDeclarationWithNamePublicIDSystemID(parser INSXMLParser, name string, publicID string, systemID string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundExternalEntityDeclarationWithName:publicID:systemID:"), parser, objc.String(name), objc.String(publicID), objc.String(systemID))
}

// Sent by a parser object to the delegate when it encounters an internal
// entity declaration.
//
// parser: An [NSXMLParser] object parsing XML.
//
// name: A string that is the declared name of an internal entity.
//
// value: A string that is the value of entity `name`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundInternalEntityDeclarationWithName:value:)
func (o NSXMLParserDelegateObject) ParserFoundInternalEntityDeclarationWithNameValue(parser INSXMLParser, name string, value string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundInternalEntityDeclarationWithName:value:"), parser, objc.String(name), objc.String(value))
}

// Sent by a parser object to its delegate when it encounters an unparsed
// entity declaration.
//
// parser: An [NSXMLParser] object parsing XML.
//
// name: A string that is the name of the unparsed entity in the declaration.
//
// publicID: A string specifying the public ID associated with the entity `name`.
//
// systemID: A string specifying the system ID associated with the entity `name`.
//
// notationName: A string specifying a notation of the declaration of entity `name`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundUnparsedEntityDeclarationWithName:publicID:systemID:notationName:)
func (o NSXMLParserDelegateObject) ParserFoundUnparsedEntityDeclarationWithNamePublicIDSystemIDNotationName(parser INSXMLParser, name string, publicID string, systemID string, notationName string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundUnparsedEntityDeclarationWithName:publicID:systemID:notationName:"), parser, objc.String(name), objc.String(publicID), objc.String(systemID), objc.String(notationName))
}

// Sent by a parser object to its delegate when it encounters a notation
// declaration.
//
// parser: An [NSXMLParser] object parsing XML.
//
// name: A string that is the name of the notation.
//
// publicID: A string specifying the public ID associated with the notation `name`.
//
// systemID: A string specifying the system ID associated with the notation `name`.
//
// See: https://developer.apple.com/documentation/Foundation/XMLParserDelegate/parser(_:foundNotationDeclarationWithName:publicID:systemID:)
func (o NSXMLParserDelegateObject) ParserFoundNotationDeclarationWithNamePublicIDSystemID(parser INSXMLParser, name string, publicID string, systemID string) {
	objc.Send[struct{}](o.ID, objc.Sel("parser:foundNotationDeclarationWithName:publicID:systemID:"), parser, objc.String(name), objc.String(publicID), objc.String(systemID))
}

// NSXMLParserDelegateConfig holds optional typed callbacks for [NSXMLParserDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate
type NSXMLParserDelegateConfig struct {

	// Handling XML
	// ParserDidStartDocument — Sent by the parser object to the delegate when it begins parsing a document.
	ParserDidStartDocument func(parser NSXMLParser)
	// ParserDidEndDocument — Sent by the parser object to the delegate when it has successfully completed parsing.
	ParserDidEndDocument func(parser NSXMLParser)
	// ParserParseErrorOccurred — Sent by a parser object to its delegate when it encounters a fatal error.
	ParserParseErrorOccurred func(parser NSXMLParser, parseError objectivec.Object)
	// ParserValidationErrorOccurred — Sent by a parser object to its delegate when it encounters a fatal validation error. [NSXMLParser] currently does not invoke this method and does not perform validation.
	ParserValidationErrorOccurred func(parser NSXMLParser, validationError objectivec.Object)
	// ParserFoundCDATA — Sent by a parser object to its delegate when it encounters a CDATA block.
	ParserFoundCDATA func(parser NSXMLParser, CDATABlock objectivec.Object)
}

// NewNSXMLParserDelegate creates an Objective-C object implementing the [NSXMLParserDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSXMLParserDelegateObject] satisfies the [NSXMLParserDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsxmlparserdelegate
func NewNSXMLParserDelegate(config NSXMLParserDelegateConfig) NSXMLParserDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSXMLParserDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ParserDidStartDocument != nil {
		fn := config.ParserDidStartDocument
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("parserDidStartDocument:"),
			Fn: func(self objc.ID, _cmd objc.SEL, parserID objc.ID) {
				parser := NSXMLParserFromID(parserID)
				fn(parser)
			},
		})
	}

	if config.ParserDidEndDocument != nil {
		fn := config.ParserDidEndDocument
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("parserDidEndDocument:"),
			Fn: func(self objc.ID, _cmd objc.SEL, parserID objc.ID) {
				parser := NSXMLParserFromID(parserID)
				fn(parser)
			},
		})
	}

	if config.ParserParseErrorOccurred != nil {
		fn := config.ParserParseErrorOccurred
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("parser:parseErrorOccurred:"),
			Fn: func(self objc.ID, _cmd objc.SEL, parserID objc.ID, parseErrorID objc.ID) {
				parser := NSXMLParserFromID(parserID)
				parseError := objectivec.ObjectFromID(parseErrorID)
				fn(parser, parseError)
			},
		})
	}

	if config.ParserValidationErrorOccurred != nil {
		fn := config.ParserValidationErrorOccurred
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("parser:validationErrorOccurred:"),
			Fn: func(self objc.ID, _cmd objc.SEL, parserID objc.ID, validationErrorID objc.ID) {
				parser := NSXMLParserFromID(parserID)
				validationError := objectivec.ObjectFromID(validationErrorID)
				fn(parser, validationError)
			},
		})
	}

	if config.ParserFoundCDATA != nil {
		fn := config.ParserFoundCDATA
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("parser:foundCDATA:"),
			Fn: func(self objc.ID, _cmd objc.SEL, parserID objc.ID, CDATABlockID objc.ID) {
				parser := NSXMLParserFromID(parserID)
				CDATABlock := objectivec.ObjectFromID(CDATABlockID)
				fn(parser, CDATABlock)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSXMLParserDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSXMLParserDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSXMLParserDelegateObjectFromID(instance)
}
