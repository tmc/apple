// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAttributedStringMarkdownSourcePosition] class.
var (
	_NSAttributedStringMarkdownSourcePositionClass     NSAttributedStringMarkdownSourcePositionClass
	_NSAttributedStringMarkdownSourcePositionClassOnce sync.Once
)

func getNSAttributedStringMarkdownSourcePositionClass() NSAttributedStringMarkdownSourcePositionClass {
	_NSAttributedStringMarkdownSourcePositionClassOnce.Do(func() {
		_NSAttributedStringMarkdownSourcePositionClass = NSAttributedStringMarkdownSourcePositionClass{class: objc.GetClass("NSAttributedStringMarkdownSourcePosition")}
	})
	return _NSAttributedStringMarkdownSourcePositionClass
}

// GetNSAttributedStringMarkdownSourcePositionClass returns the class object for NSAttributedStringMarkdownSourcePosition.
func GetNSAttributedStringMarkdownSourcePositionClass() NSAttributedStringMarkdownSourcePositionClass {
	return getNSAttributedStringMarkdownSourcePositionClass()
}

type NSAttributedStringMarkdownSourcePositionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAttributedStringMarkdownSourcePositionClass) Alloc() NSAttributedStringMarkdownSourcePosition {
	rv := objc.Send[NSAttributedStringMarkdownSourcePosition](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// The position of attributed string text in its original Markdown source
// string.
//
// # Creating an Attributed String Markdown Source Position Instance
//
//   - [NSAttributedStringMarkdownSourcePosition.InitWithStartLineStartColumnEndLineEndColumn]: Creates a Markdown source position instance from its start and end line and column.
//
// # Getting Markdown Source Position Properties
//
//   - [NSAttributedStringMarkdownSourcePosition.StartLine]: The line where the text begins in the Markdown source.
//   - [NSAttributedStringMarkdownSourcePosition.StartColumn]: The column where the text begins in the Markdown source.
//   - [NSAttributedStringMarkdownSourcePosition.EndLine]: The line where the text ends in the Markdown source.
//   - [NSAttributedStringMarkdownSourcePosition.EndColumn]: The column where the text ends in the Markdown source.
//
// # Getting a Range from a Markdown Source Position Attribute
//
//   - [NSAttributedStringMarkdownSourcePosition.RangeInString]: Returns a range indicating the source portion within a Markdown string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition
type NSAttributedStringMarkdownSourcePosition struct {
	objectivec.Object
}

// NSAttributedStringMarkdownSourcePositionFromID constructs a [NSAttributedStringMarkdownSourcePosition] from an objc.ID.
//
// The position of attributed string text in its original Markdown source
// string.
func NSAttributedStringMarkdownSourcePositionFromID(id objc.ID) NSAttributedStringMarkdownSourcePosition {
	return NSAttributedStringMarkdownSourcePosition{objectivec.Object{id}}
}
// NOTE: NSAttributedStringMarkdownSourcePosition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSAttributedStringMarkdownSourcePosition] class.
//
// # Creating an Attributed String Markdown Source Position Instance
//
//   - [INSAttributedStringMarkdownSourcePosition.InitWithStartLineStartColumnEndLineEndColumn]: Creates a Markdown source position instance from its start and end line and column.
//
// # Getting Markdown Source Position Properties
//
//   - [INSAttributedStringMarkdownSourcePosition.StartLine]: The line where the text begins in the Markdown source.
//   - [INSAttributedStringMarkdownSourcePosition.StartColumn]: The column where the text begins in the Markdown source.
//   - [INSAttributedStringMarkdownSourcePosition.EndLine]: The line where the text ends in the Markdown source.
//   - [INSAttributedStringMarkdownSourcePosition.EndColumn]: The column where the text ends in the Markdown source.
//
// # Getting a Range from a Markdown Source Position Attribute
//
//   - [INSAttributedStringMarkdownSourcePosition.RangeInString]: Returns a range indicating the source portion within a Markdown string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition
type INSAttributedStringMarkdownSourcePosition interface {
	objectivec.IObject
	NSCopying

	// Topic: Creating an Attributed String Markdown Source Position Instance

	// Creates a Markdown source position instance from its start and end line and column.
	InitWithStartLineStartColumnEndLineEndColumn(startLine int, startColumn int, endLine int, endColumn int) NSAttributedStringMarkdownSourcePosition

	// Topic: Getting Markdown Source Position Properties

	// The line where the text begins in the Markdown source.
	StartLine() int
	// The column where the text begins in the Markdown source.
	StartColumn() int
	// The line where the text ends in the Markdown source.
	EndLine() int
	// The column where the text ends in the Markdown source.
	EndColumn() int

	// Topic: Getting a Range from a Markdown Source Position Attribute

	// Returns a range indicating the source portion within a Markdown string.
	RangeInString(string_ string) NSRange

	InitWithCoder(coder INSCoder) NSAttributedStringMarkdownSourcePosition
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}





// Init initializes the instance.
func (a NSAttributedStringMarkdownSourcePosition) Init() NSAttributedStringMarkdownSourcePosition {
	rv := objc.Send[NSAttributedStringMarkdownSourcePosition](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAttributedStringMarkdownSourcePosition) Autorelease() NSAttributedStringMarkdownSourcePosition {
	rv := objc.Send[NSAttributedStringMarkdownSourcePosition](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAttributedStringMarkdownSourcePosition creates a new NSAttributedStringMarkdownSourcePosition instance.
func NewNSAttributedStringMarkdownSourcePosition() NSAttributedStringMarkdownSourcePosition {
	class := getNSAttributedStringMarkdownSourcePositionClass()
	rv := objc.Send[NSAttributedStringMarkdownSourcePosition](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a Markdown source position instance from its start and end line and
// column.
//
// startLine: The line number where text begins in the Markdown source. Specify a 1-based
// number. For example, the number for the first row is 1, for the second row
// is 2, and so on.
//
// startColumn: The column number where text begins in the Markdown source. Specify a
// 1-based number. For example, the number for the first column is 1, for the
// second column is 2, and so on. Columns represent UTF-8 indices; for
// multi-byte characters, the column indicates the first byte.
//
// endLine: The line number where the Markdown source ends. Specify a 1-based number.
//
// endColumn: The column number where the Markdown source ends. Specify a 1-based number.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/initWithStartLine:startColumn:endLine:endColumn:
func NewAttributedStringMarkdownSourcePositionWithStartLineStartColumnEndLineEndColumn(startLine int, startColumn int, endLine int, endColumn int) NSAttributedStringMarkdownSourcePosition {
	instance := getNSAttributedStringMarkdownSourcePositionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStartLine:startColumn:endLine:endColumn:"), startLine, startColumn, endLine, endColumn)
	return NSAttributedStringMarkdownSourcePositionFromID(rv)
}







// Creates a Markdown source position instance from its start and end line and
// column.
//
// startLine: The line number where text begins in the Markdown source. Specify a 1-based
// number. For example, the number for the first row is 1, for the second row
// is 2, and so on.
//
// startColumn: The column number where text begins in the Markdown source. Specify a
// 1-based number. For example, the number for the first column is 1, for the
// second column is 2, and so on. Columns represent UTF-8 indices; for
// multi-byte characters, the column indicates the first byte.
//
// endLine: The line number where the Markdown source ends. Specify a 1-based number.
//
// endColumn: The column number where the Markdown source ends. Specify a 1-based number.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/initWithStartLine:startColumn:endLine:endColumn:
func (a NSAttributedStringMarkdownSourcePosition) InitWithStartLineStartColumnEndLineEndColumn(startLine int, startColumn int, endLine int, endColumn int) NSAttributedStringMarkdownSourcePosition {
	rv := objc.Send[NSAttributedStringMarkdownSourcePosition](a.ID, objc.Sel("initWithStartLine:startColumn:endLine:endColumn:"), startLine, startColumn, endLine, endColumn)
	return rv
}

// Returns a range indicating the source portion within a Markdown string.
//
// string: The Markdown source string that this source position object refers to.
//
// # Return Value
// 
// A range that represents the source portion within a source Markdown string.
//
// # Discussion
// 
// Use this method to access the marked-up region of `string` with an NSRange,
// rather than making manual calculations based on row and column values.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/rangeInString:
func (a NSAttributedStringMarkdownSourcePosition) RangeInString(string_ string) NSRange {
	rv := objc.Send[NSRange](a.ID, objc.Sel("rangeInString:"), objc.String(string_))
	return NSRange(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (a NSAttributedStringMarkdownSourcePosition) InitWithCoder(coder INSCoder) NSAttributedStringMarkdownSourcePosition {
	rv := objc.Send[NSAttributedStringMarkdownSourcePosition](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (a NSAttributedStringMarkdownSourcePosition) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The line where the text begins in the Markdown source.
//
// # Discussion
// 
// This property uses `1`-based counting.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/startLine
func (a NSAttributedStringMarkdownSourcePosition) StartLine() int {
	rv := objc.Send[int](a.ID, objc.Sel("startLine"))
	return rv
}



// The column where the text begins in the Markdown source.
//
// # Discussion
// 
// This property uses `1`-based counting. Columns represent UTF-8 indices; for
// multi-byte characters, the column indicates the first byte.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/startColumn
func (a NSAttributedStringMarkdownSourcePosition) StartColumn() int {
	rv := objc.Send[int](a.ID, objc.Sel("startColumn"))
	return rv
}



// The line where the text ends in the Markdown source.
//
// # Discussion
// 
// This property uses `1`-based counting.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/endLine
func (a NSAttributedStringMarkdownSourcePosition) EndLine() int {
	rv := objc.Send[int](a.ID, objc.Sel("endLine"))
	return rv
}



// The column where the text ends in the Markdown source.
//
// # Discussion
// 
// This property uses `1`-based counting. Columns represent UTF-8 indices; for
// multi-byte characters, the column indicates the first byte.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition/endColumn
func (a NSAttributedStringMarkdownSourcePosition) EndColumn() int {
	rv := objc.Send[int](a.ID, objc.Sel("endColumn"))
	return rv
}








			// Protocol methods for NSCopying
			














