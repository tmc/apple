// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPresentationIntent] class.
var (
	_NSPresentationIntentClass     NSPresentationIntentClass
	_NSPresentationIntentClassOnce sync.Once
)

func getNSPresentationIntentClass() NSPresentationIntentClass {
	_NSPresentationIntentClassOnce.Do(func() {
		_NSPresentationIntentClass = NSPresentationIntentClass{class: objc.GetClass("NSPresentationIntent")}
	})
	return _NSPresentationIntentClass
}

// GetNSPresentationIntentClass returns the class object for NSPresentationIntent.
func GetNSPresentationIntentClass() NSPresentationIntentClass {
	return getNSPresentationIntentClass()
}

type NSPresentationIntentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPresentationIntentClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPresentationIntentClass) Alloc() NSPresentationIntent {
	rv := objc.Send[NSPresentationIntent](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A type that contains the Markdown formatting for blocks of text, like
// paragraphs, lists, code blocks, and parts of tables.
//
// # Overview
// 
// An [NSPresentationIntent] object stores the Markdown semantics for a range
// of characters in an attributed string. When parsing Markdown into an
// attributed string, the system sets the value of the
// [presentationIntentAttributeName] attribute to an instance of this class.
// When displaying your string in system views, the system applies a default
// visual style to match the corresponding information in this type. To
// replace the system’s default formatting, remove these attributes from
// your attributed string and apply the formatting you want.
//
// [presentationIntentAttributeName]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/presentationIntentAttributeName
//
// # Getting the intent identity
//
//   - [NSPresentationIntent.Identity]: A unique identifier for the intent in the document.
//   - [NSPresentationIntent.IntentKind]: The type of the intent.
//   - [NSPresentationIntent.ParentIntent]: The parent of the current intent.
//   - [NSPresentationIntent.IsEquivalentToPresentationIntent]: Returns a Boolean value that indicates whether the current intent is equivalent to the specified intent.
//
// # Getting header information
//
//   - [NSPresentationIntent.HeaderLevel]: The level of a header section.
//
// # Getting list information
//
//   - [NSPresentationIntent.Ordinal]: The number for an item in an ordered list.
//   - [NSPresentationIntent.IndentationLevel]: The indentation level of the intent.
//
// # Getting table information
//
//   - [NSPresentationIntent.Row]: The row number to which this cell belongs.
//   - [NSPresentationIntent.Column]: The column number to which the cell belongs.
//   - [NSPresentationIntent.ColumnCount]: The number of columns in a table.
//   - [NSPresentationIntent.ColumnAlignments]: The alignments for the columns in a table.
//
// # Getting code information
//
//   - [NSPresentationIntent.LanguageHint]: The language associated with the code listing.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent
type NSPresentationIntent struct {
	objectivec.Object
}

// NSPresentationIntentFromID constructs a [NSPresentationIntent] from an objc.ID.
//
// A type that contains the Markdown formatting for blocks of text, like
// paragraphs, lists, code blocks, and parts of tables.
func NSPresentationIntentFromID(id objc.ID) NSPresentationIntent {
	return NSPresentationIntent{objectivec.Object{ID: id}}
}
// NOTE: NSPresentationIntent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPresentationIntent] class.
//
// # Getting the intent identity
//
//   - [INSPresentationIntent.Identity]: A unique identifier for the intent in the document.
//   - [INSPresentationIntent.IntentKind]: The type of the intent.
//   - [INSPresentationIntent.ParentIntent]: The parent of the current intent.
//   - [INSPresentationIntent.IsEquivalentToPresentationIntent]: Returns a Boolean value that indicates whether the current intent is equivalent to the specified intent.
//
// # Getting header information
//
//   - [INSPresentationIntent.HeaderLevel]: The level of a header section.
//
// # Getting list information
//
//   - [INSPresentationIntent.Ordinal]: The number for an item in an ordered list.
//   - [INSPresentationIntent.IndentationLevel]: The indentation level of the intent.
//
// # Getting table information
//
//   - [INSPresentationIntent.Row]: The row number to which this cell belongs.
//   - [INSPresentationIntent.Column]: The column number to which the cell belongs.
//   - [INSPresentationIntent.ColumnCount]: The number of columns in a table.
//   - [INSPresentationIntent.ColumnAlignments]: The alignments for the columns in a table.
//
// # Getting code information
//
//   - [INSPresentationIntent.LanguageHint]: The language associated with the code listing.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent
type INSPresentationIntent interface {
	objectivec.IObject
	NSCopying
	NSSecureCoding

	// Topic: Getting the intent identity

	// A unique identifier for the intent in the document.
	Identity() int
	// The type of the intent.
	IntentKind() NSPresentationIntentKind
	// The parent of the current intent.
	ParentIntent() INSPresentationIntent
	// Returns a Boolean value that indicates whether the current intent is equivalent to the specified intent.
	IsEquivalentToPresentationIntent(other INSPresentationIntent) bool

	// Topic: Getting header information

	// The level of a header section.
	HeaderLevel() int

	// Topic: Getting list information

	// The number for an item in an ordered list.
	Ordinal() int
	// The indentation level of the intent.
	IndentationLevel() int

	// Topic: Getting table information

	// The row number to which this cell belongs.
	Row() int
	// The column number to which the cell belongs.
	Column() int
	// The number of columns in a table.
	ColumnCount() int
	// The alignments for the columns in a table.
	ColumnAlignments() []NSNumber

	// Topic: Getting code information

	// The language associated with the code listing.
	LanguageHint() string

	InitWithCoder(coder INSCoder) NSPresentationIntent
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}

// Init initializes the instance.
func (p NSPresentationIntent) Init() NSPresentationIntent {
	rv := objc.Send[NSPresentationIntent](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPresentationIntent) Autorelease() NSPresentationIntent {
	rv := objc.Send[NSPresentationIntent](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPresentationIntent creates a new NSPresentationIntent instance.
func NewNSPresentationIntent() NSPresentationIntent {
	class := getNSPresentationIntentClass()
	rv := objc.Send[NSPresentationIntent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a Boolean value that indicates whether the current intent is
// equivalent to the specified intent.
//
// other: The other intent to use in the comparison.
//
// # Return Value
// 
// [true] if the current intent is equivalent to the specified intent, or
// [false] if it isn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Two intents are equivalent if their attributes match. This method doesn’t
// consider the [Identity] property of the intents when determining their
// equivalence.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/isEquivalentToPresentationIntent:
func (p NSPresentationIntent) IsEquivalentToPresentationIntent(other INSPresentationIntent) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isEquivalentToPresentationIntent:"), other)
	return rv
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (p NSPresentationIntent) InitWithCoder(coder INSCoder) NSPresentationIntent {
	rv := objc.Send[NSPresentationIntent](p.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (p NSPresentationIntent) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a paragraph intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// parent: The parent intent of the paragraph.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindParagraph].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/paragraphIntentWithIdentity:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) ParagraphIntentWithIdentityNestedInsideIntent(identity int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("paragraphIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return NSPresentationIntentFromID(rv)
}
// Creates a header intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// level: The level for the header section. Specify `1` or greater for this
// parameter. Don’t specify `0`.
//
// parent: The parent intent of the header.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindHeader].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/headerIntentWithIdentity:level:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) HeaderIntentWithIdentityLevelNestedInsideIntent(identity int, level int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("headerIntentWithIdentity:level:nestedInsideIntent:"), identity, level, parent)
	return NSPresentationIntentFromID(rv)
}
// Creates an ordered-list intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// parent: The parent intent of the ordered list.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindOrderedList].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/orderedListIntentWithIdentity:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) OrderedListIntentWithIdentityNestedInsideIntent(identity int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("orderedListIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return NSPresentationIntentFromID(rv)
}
// Creates an unordered-list intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// parent: The parent intent of the list.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindUnorderedList].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/unorderedListIntentWithIdentity:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) UnorderedListIntentWithIdentityNestedInsideIntent(identity int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("unorderedListIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return NSPresentationIntentFromID(rv)
}
// Creates an item for an ordered list with the provided information.
//
// identity: The unique identifier for the intent.
//
// ordinal: The ordinal number to display in front of the list item.
//
// parent: The parent intent of the item.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindListItem].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/listItemIntentWithIdentity:ordinal:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) ListItemIntentWithIdentityOrdinalNestedInsideIntent(identity int, ordinal int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("listItemIntentWithIdentity:ordinal:nestedInsideIntent:"), identity, ordinal, parent)
	return NSPresentationIntentFromID(rv)
}
// Creates an code-block intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// languageHint: The programming language for the code listing.
//
// parent: The parent intent of the code block.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindCodeBlock].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/codeBlockIntentWithIdentity:languageHint:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) CodeBlockIntentWithIdentityLanguageHintNestedInsideIntent(identity int, languageHint string, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("codeBlockIntentWithIdentity:languageHint:nestedInsideIntent:"), identity, objc.String(languageHint), parent)
	return NSPresentationIntentFromID(rv)
}
// Creates a block-quote intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// parent: The parent intent of the block quote.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindBlockQuote].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/blockQuoteIntentWithIdentity:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) BlockQuoteIntentWithIdentityNestedInsideIntent(identity int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("blockQuoteIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return NSPresentationIntentFromID(rv)
}
// Creates a thematic break intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// parent: The parent intent of the horizontal rule.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindThematicBreak].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/thematicBreakIntentWithIdentity:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) ThematicBreakIntentWithIdentityNestedInsideIntent(identity int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("thematicBreakIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return NSPresentationIntentFromID(rv)
}
// Creates a table intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// columnCount: The number of columns in the table.
//
// alignments: The text alignments for each column. For each [NSNumber] in the array, set
// the value to a value from the [NSPresentationIntentTableColumnAlignment]
// enumerated type.
// //
// [NSPresentationIntentTableColumnAlignment]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment
//
// parent: The parent intent of the table.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindTable].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/tableIntentWithIdentity:columnCount:alignments:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) TableIntentWithIdentityColumnCountAlignmentsNestedInsideIntent(identity int, columnCount int, alignments []NSNumber, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("tableIntentWithIdentity:columnCount:alignments:nestedInsideIntent:"), identity, columnCount, objectivec.IObjectSliceToNSArray(alignments), parent)
	return NSPresentationIntentFromID(rv)
}
// Creates a table header intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// parent: The parent intent of the table header row.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindTableHeaderRow].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/tableHeaderRowIntentWithIdentity:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) TableHeaderRowIntentWithIdentityNestedInsideIntent(identity int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("tableHeaderRowIntentWithIdentity:nestedInsideIntent:"), identity, parent)
	return NSPresentationIntentFromID(rv)
}
// Creates a table row intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// row: The row number.
//
// parent: The parent intent of row.
//
// # Return Value
// 
// A new intent with the kind set to [PresentationIntentKindTableRow].
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/tableRowIntentWithIdentity:row:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) TableRowIntentWithIdentityRowNestedInsideIntent(identity int, row int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("tableRowIntentWithIdentity:row:nestedInsideIntent:"), identity, row, parent)
	return NSPresentationIntentFromID(rv)
}
// Creates a table cell intent with the provided information.
//
// identity: The unique identifier for the intent.
//
// column: The column number for the cell.
//
// parent: The parent intent of the cell.
//
// # Return Value
// 
// A new intent with the kind set to [NSPresentationIntentKindTableCell].
//
// [NSPresentationIntentKindTableCell]: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind/NSPresentationIntentKindTableCell
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/tableCellIntentWithIdentity:column:nestedInsideIntent:
func (_NSPresentationIntentClass NSPresentationIntentClass) TableCellIntentWithIdentityColumnNestedInsideIntent(identity int, column int, parent INSPresentationIntent) NSPresentationIntent {
	rv := objc.Send[objc.ID](objc.ID(_NSPresentationIntentClass.class), objc.Sel("tableCellIntentWithIdentity:column:nestedInsideIntent:"), identity, column, parent)
	return NSPresentationIntentFromID(rv)
}

// A unique identifier for the intent in the document.
//
// # Discussion
// 
// Use the value in this property to disambiguate attributes that apply to
// contiguous text. For example, you might use it to differentiate between two
// headers in a row with the same level.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/identity
func (p NSPresentationIntent) Identity() int {
	rv := objc.Send[int](p.ID, objc.Sel("identity"))
	return rv
}
// The type of the intent.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/intentKind
func (p NSPresentationIntent) IntentKind() NSPresentationIntentKind {
	rv := objc.Send[NSPresentationIntentKind](p.ID, objc.Sel("intentKind"))
	return NSPresentationIntentKind(rv)
}
// The parent of the current intent.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/parentIntent
func (p NSPresentationIntent) ParentIntent() INSPresentationIntent {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("parentIntent"))
	return NSPresentationIntentFromID(objc.ID(rv))
}
// The level of a header section.
//
// # Discussion
// 
// This value corresponds to the number of hash marks (`#`) associated with
// the header. If the intent is not a header, the value of this property is
// `0`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/headerLevel
func (p NSPresentationIntent) HeaderLevel() int {
	rv := objc.Send[int](p.ID, objc.Sel("headerLevel"))
	return rv
}
// The number for an item in an ordered list.
//
// # Discussion
// 
// If the intent is not a list, the value of this property is `0`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/ordinal
func (p NSPresentationIntent) Ordinal() int {
	rv := objc.Send[int](p.ID, objc.Sel("ordinal"))
	return rv
}
// The indentation level of the intent.
//
// # Discussion
// 
// The initial list has an indentation level of `0`. Each time you nest a new
// list, the indentation level for new list increases by `1`. All elements
// within the same list have the same indentation level.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/indentationLevel
func (p NSPresentationIntent) IndentationLevel() int {
	rv := objc.Send[int](p.ID, objc.Sel("indentationLevel"))
	return rv
}
// The row number to which this cell belongs.
//
// # Discussion
// 
// The value of this property is `0`-based, with the first row at `0`, the
// second row at `1`, and so on. If The intent is not a cell, this value is
// `0`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/row
func (p NSPresentationIntent) Row() int {
	rv := objc.Send[int](p.ID, objc.Sel("row"))
	return rv
}
// The column number to which the cell belongs.
//
// # Discussion
// 
// The value of this property is `0`-based, with the first column at `0`, the
// second column at `1`, and so on. Header rows are always at row `0`, with
// subsequent rows starting at `1`. If The intent is not a cell, this value is
// `0`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/column
func (p NSPresentationIntent) Column() int {
	rv := objc.Send[int](p.ID, objc.Sel("column"))
	return rv
}
// The number of columns in a table.
//
// # Discussion
// 
// If the intent is not a table, the value of this property is `0`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/columnCount
func (p NSPresentationIntent) ColumnCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("columnCount"))
	return rv
}
// The alignments for the columns in a table.
//
// # Discussion
// 
// If the intent is not a table, the value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/columnAlignments
func (p NSPresentationIntent) ColumnAlignments() []NSNumber {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("columnAlignments"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSNumber {
		return NSNumberFromID(id)
	})
}
// The language associated with the code listing.
//
// # Discussion
// 
// If the intent is not a code block, the value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntent/languageHint
func (p NSPresentationIntent) LanguageHint() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("languageHint"))
	return NSStringFromID(rv).String()
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

