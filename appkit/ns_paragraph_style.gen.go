// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSParagraphStyle] class.
var (
	_NSParagraphStyleClass     NSParagraphStyleClass
	_NSParagraphStyleClassOnce sync.Once
)

func getNSParagraphStyleClass() NSParagraphStyleClass {
	_NSParagraphStyleClassOnce.Do(func() {
		_NSParagraphStyleClass = NSParagraphStyleClass{class: objc.GetClass("NSParagraphStyle")}
	})
	return _NSParagraphStyleClass
}

// GetNSParagraphStyleClass returns the class object for NSParagraphStyle.
func GetNSParagraphStyleClass() NSParagraphStyleClass {
	return getNSParagraphStyleClass()
}

type NSParagraphStyleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSParagraphStyleClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSParagraphStyleClass) Alloc() NSParagraphStyle {
	rv := objc.Send[NSParagraphStyle](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The paragraph or ruler attributes for an attributed string.
//
// # Overview
// 
// An [NSParagraphStyle] object stores formatting information for a paragraph
// of text. The formatting information includes the amount of space between
// lines, indentations for lines of text, line heights, tab-stop positions,
// and more. Apply paragraph styles to the text of an attributed string by
// adding the [ParagraphStyle] attribute and setting its value to an instance
// of this class. The text-rendering system uses the paragraph style
// information in an attributed string to lay out and render the text.
// 
// The [NSParagraphStyle] class manages an immutable set of style information,
// but you can create an [NSMutableParagraphStyle] when you want to modify the
// style information before applying it to your text.
//
// # Accessing style information
//
//   - [NSParagraphStyle.Alignment]: The text alignment of the paragraph.
//   - [NSParagraphStyle.FirstLineHeadIndent]: The indentation of the first line of the paragraph.
//   - [NSParagraphStyle.HeadIndent]: The indentation of the paragraph’s lines other than the first.
//   - [NSParagraphStyle.TailIndent]: The trailing indentation of the paragraph.
//   - [NSParagraphStyle.LineHeightMultiple]: The line height multiple.
//   - [NSParagraphStyle.MaximumLineHeight]: The paragraph’s maximum line height.
//   - [NSParagraphStyle.MinimumLineHeight]: The paragraph’s minimum line height.
//   - [NSParagraphStyle.LineSpacing]: The distance in points between the bottom of one line fragment and the top of the next.
//   - [NSParagraphStyle.ParagraphSpacing]: Distance between the bottom of this paragraph and top of next.
//   - [NSParagraphStyle.ParagraphSpacingBefore]: The distance between the paragraph’s top and the beginning of its text content.
//
// # Accessing tab information
//
//   - [NSParagraphStyle.TabStops]: The text tab objects that represent the paragraph’s tab stops.
//   - [NSParagraphStyle.DefaultTabInterval]: The documentwide default tab interval.
//
// # Getting text block and list information
//
//   - [NSParagraphStyle.TextBlocks]: The text blocks that contain the paragraph.
//   - [NSParagraphStyle.TextLists]: The text lists that contain the paragraph.
//
// # Getting line-break information
//
//   - [NSParagraphStyle.LineBreakMode]: The mode for breaking lines in the paragraph that don’t fit within a container.
//   - [NSParagraphStyle.LineBreakStrategy]: The strategy for breaking lines while laying out paragraphs.
//   - [NSParagraphStyle.HyphenationFactor]: The paragraph’s threshold for hyphenation.
//   - [NSParagraphStyle.UsesDefaultHyphenation]: A Boolean value that indicates whether the paragraph style uses the system hyphenation settings.
//   - [NSParagraphStyle.TighteningFactorForTruncation]: The threshold for using tightening as an alternative to truncation.
//   - [NSParagraphStyle.AllowsDefaultTighteningForTruncation]: A Boolean value that indicates whether the system tightens character spacing before truncating text.
//
// # Getting the html header level
//
//   - [NSParagraphStyle.HeaderLevel]: The paragraph’s header level for HTML generation.
//
// # Determining writing direction
//
//   - [NSParagraphStyle.BaseWritingDirection]: The base writing direction for the paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle
type NSParagraphStyle struct {
	objectivec.Object
}

// NSParagraphStyleFromID constructs a [NSParagraphStyle] from an objc.ID.
//
// The paragraph or ruler attributes for an attributed string.
func NSParagraphStyleFromID(id objc.ID) NSParagraphStyle {
	return NSParagraphStyle{objectivec.Object{ID: id}}
}
// NOTE: NSParagraphStyle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSParagraphStyle] class.
//
// # Accessing style information
//
//   - [INSParagraphStyle.Alignment]: The text alignment of the paragraph.
//   - [INSParagraphStyle.FirstLineHeadIndent]: The indentation of the first line of the paragraph.
//   - [INSParagraphStyle.HeadIndent]: The indentation of the paragraph’s lines other than the first.
//   - [INSParagraphStyle.TailIndent]: The trailing indentation of the paragraph.
//   - [INSParagraphStyle.LineHeightMultiple]: The line height multiple.
//   - [INSParagraphStyle.MaximumLineHeight]: The paragraph’s maximum line height.
//   - [INSParagraphStyle.MinimumLineHeight]: The paragraph’s minimum line height.
//   - [INSParagraphStyle.LineSpacing]: The distance in points between the bottom of one line fragment and the top of the next.
//   - [INSParagraphStyle.ParagraphSpacing]: Distance between the bottom of this paragraph and top of next.
//   - [INSParagraphStyle.ParagraphSpacingBefore]: The distance between the paragraph’s top and the beginning of its text content.
//
// # Accessing tab information
//
//   - [INSParagraphStyle.TabStops]: The text tab objects that represent the paragraph’s tab stops.
//   - [INSParagraphStyle.DefaultTabInterval]: The documentwide default tab interval.
//
// # Getting text block and list information
//
//   - [INSParagraphStyle.TextBlocks]: The text blocks that contain the paragraph.
//   - [INSParagraphStyle.TextLists]: The text lists that contain the paragraph.
//
// # Getting line-break information
//
//   - [INSParagraphStyle.LineBreakMode]: The mode for breaking lines in the paragraph that don’t fit within a container.
//   - [INSParagraphStyle.LineBreakStrategy]: The strategy for breaking lines while laying out paragraphs.
//   - [INSParagraphStyle.HyphenationFactor]: The paragraph’s threshold for hyphenation.
//   - [INSParagraphStyle.UsesDefaultHyphenation]: A Boolean value that indicates whether the paragraph style uses the system hyphenation settings.
//   - [INSParagraphStyle.TighteningFactorForTruncation]: The threshold for using tightening as an alternative to truncation.
//   - [INSParagraphStyle.AllowsDefaultTighteningForTruncation]: A Boolean value that indicates whether the system tightens character spacing before truncating text.
//
// # Getting the html header level
//
//   - [INSParagraphStyle.HeaderLevel]: The paragraph’s header level for HTML generation.
//
// # Determining writing direction
//
//   - [INSParagraphStyle.BaseWritingDirection]: The base writing direction for the paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle
type INSParagraphStyle interface {
	objectivec.IObject

	// Topic: Accessing style information

	// The text alignment of the paragraph.
	Alignment() NSTextAlignment
	// The indentation of the first line of the paragraph.
	FirstLineHeadIndent() float64
	// The indentation of the paragraph’s lines other than the first.
	HeadIndent() float64
	// The trailing indentation of the paragraph.
	TailIndent() float64
	// The line height multiple.
	LineHeightMultiple() float64
	// The paragraph’s maximum line height.
	MaximumLineHeight() float64
	// The paragraph’s minimum line height.
	MinimumLineHeight() float64
	// The distance in points between the bottom of one line fragment and the top of the next.
	LineSpacing() float64
	// Distance between the bottom of this paragraph and top of next.
	ParagraphSpacing() float64
	// The distance between the paragraph’s top and the beginning of its text content.
	ParagraphSpacingBefore() float64

	// Topic: Accessing tab information

	// The text tab objects that represent the paragraph’s tab stops.
	TabStops() []NSTextTab
	// The documentwide default tab interval.
	DefaultTabInterval() float64

	// Topic: Getting text block and list information

	// The text blocks that contain the paragraph.
	TextBlocks() []NSTextBlock
	// The text lists that contain the paragraph.
	TextLists() []NSTextList

	// Topic: Getting line-break information

	// The mode for breaking lines in the paragraph that don’t fit within a container.
	LineBreakMode() NSLineBreakMode
	// The strategy for breaking lines while laying out paragraphs.
	LineBreakStrategy() NSLineBreakStrategy
	// The paragraph’s threshold for hyphenation.
	HyphenationFactor() float32
	// A Boolean value that indicates whether the paragraph style uses the system hyphenation settings.
	UsesDefaultHyphenation() bool
	// The threshold for using tightening as an alternative to truncation.
	TighteningFactorForTruncation() float32
	// A Boolean value that indicates whether the system tightens character spacing before truncating text.
	AllowsDefaultTighteningForTruncation() bool

	// Topic: Getting the html header level

	// The paragraph’s header level for HTML generation.
	HeaderLevel() int

	// Topic: Determining writing direction

	// The base writing direction for the paragraph.
	BaseWritingDirection() NSWritingDirection

	// The paragraph style of the text.
	ParagraphStyle() foundation.NSString
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NSParagraphStyle) Init() NSParagraphStyle {
	rv := objc.Send[NSParagraphStyle](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSParagraphStyle) Autorelease() NSParagraphStyle {
	rv := objc.Send[NSParagraphStyle](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSParagraphStyle creates a new NSParagraphStyle instance.
func NewNSParagraphStyle() NSParagraphStyle {
	class := getNSParagraphStyleClass()
	rv := objc.Send[NSParagraphStyle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p NSParagraphStyle) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the default writing direction for the specified language.
//
// languageName: The language specified in ISO language region format. Can be `nil` to
// return a default writing direction derived from the user’s defaults
// database.
//
// # Return Value
// 
// The default writing direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/defaultWritingDirection(forLanguage:)
func (_NSParagraphStyleClass NSParagraphStyleClass) DefaultWritingDirectionForLanguage(languageName string) NSWritingDirection {
	rv := objc.Send[NSWritingDirection](objc.ID(_NSParagraphStyleClass.class), objc.Sel("defaultWritingDirectionForLanguage:"), objc.String(languageName))
	return NSWritingDirection(rv)
}

// The text alignment of the paragraph.
//
// # Discussion
// 
// The framework displays natural text alignment as left or right alignment
// depending on the line sweep direction of the first script contained in the
// paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/alignment
func (p NSParagraphStyle) Alignment() NSTextAlignment {
	rv := objc.Send[NSTextAlignment](p.ID, objc.Sel("alignment"))
	return NSTextAlignment(rv)
}
// The indentation of the first line of the paragraph.
//
// # Discussion
// 
// This property contains the distance (in points) from the leading margin of
// a text container to the beginning of the paragraph’s first line. This
// value is always nonnegative.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/firstLineHeadIndent
func (p NSParagraphStyle) FirstLineHeadIndent() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("firstLineHeadIndent"))
	return rv
}
// The indentation of the paragraph’s lines other than the first.
//
// # Discussion
// 
// This property contains the distance (in points) from the leading margin of
// a text container to the beginning of lines other than the first. This value
// is always nonnegative.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/headIndent
func (p NSParagraphStyle) HeadIndent() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("headIndent"))
	return rv
}
// The trailing indentation of the paragraph.
//
// # Discussion
// 
// If positive, this value is the distance from the leading margin (for
// example, the left margin in left-to-right text). If `0` or negative, it’s
// the distance from the trailing margin.
// 
// For example, a paragraph style designed to fit exactly in a two-inch wide
// container has a head indent of `0.0` and a tail indent of `0.0`. One
// designed to fit with a quarter-inch margin has a head indent of `0.25` and
// a tail indent of `–0.25`.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/tailIndent
func (p NSParagraphStyle) TailIndent() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("tailIndent"))
	return rv
}
// The line height multiple.
//
// # Discussion
// 
// The framework multiplies the natural line height of the receiver by this
// factor (if positive), and constrains the resulting value by the minimum and
// maximum line height. The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineHeightMultiple
func (p NSParagraphStyle) LineHeightMultiple() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("lineHeightMultiple"))
	return rv
}
// The paragraph’s maximum line height.
//
// # Discussion
// 
// This property contains the maximum height in points that any line in the
// receiver occupies, regardless of the font size or size of any attached
// graphic. This value is always nonnegative. The default value is `0`.
// 
// Glyphs and graphics exceeding this height overlaps neighboring lines;
// however, a maximum height of `0` implies no line height limit. Although
// this limit applies to the line itself, line spacing adds extra space
// between adjacent lines.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/maximumLineHeight
func (p NSParagraphStyle) MaximumLineHeight() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("maximumLineHeight"))
	return rv
}
// The paragraph’s minimum line height.
//
// # Discussion
// 
// This property contains the minimum height in points that any line in the
// receiver occupies, regardless of the font size or size of any attached
// graphic. This value is always nonnegative.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/minimumLineHeight
func (p NSParagraphStyle) MinimumLineHeight() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("minimumLineHeight"))
	return rv
}
// The distance in points between the bottom of one line fragment and the top
// of the next.
//
// # Discussion
// 
// This value is always nonnegative. The layout manager uses this value in the
// line fragment height.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineSpacing
func (p NSParagraphStyle) LineSpacing() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("lineSpacing"))
	return rv
}
// Distance between the bottom of this paragraph and top of next.
//
// # Discussion
// 
// This property contains the space (measured in points) between paragraphs.
// This value is always nonnegative. The framework determines the space
// between paragraphs by adding the previous paragraph’s `paragraphSpacing`
// and the current paragraph’s [ParagraphSpacingBefore].
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/paragraphSpacing
func (p NSParagraphStyle) ParagraphSpacing() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("paragraphSpacing"))
	return rv
}
// The distance between the paragraph’s top and the beginning of its text
// content.
//
// # Discussion
// 
// This property contains the space (measured in points) between the current
// and previous paragraphs. The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/paragraphSpacingBefore
func (p NSParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("paragraphSpacingBefore"))
	return rv
}
// The text tab objects that represent the paragraph’s tab stops.
//
// # Discussion
// 
// The [NSTextTab] objects, sorted by location, define the tab stops for the
// paragraph style. The default value is an array of 12 left-aligned tabs at
// 28-point intervals.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/tabStops
func (p NSParagraphStyle) TabStops() []NSTextTab {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("tabStops"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextTab {
		return NSTextTabFromID(id)
	})
}
// The documentwide default tab interval.
//
// # Discussion
// 
// This property represents the default tab interval in points. Tabs after the
// last specified in [TabStops] are placed at integer multiples of this
// distance (if positive). Default value is 0.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/defaultTabInterval
func (p NSParagraphStyle) DefaultTabInterval() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("defaultTabInterval"))
	return rv
}
// The text blocks that contain the paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/textBlocks
func (p NSParagraphStyle) TextBlocks() []NSTextBlock {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("textBlocks"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextBlock {
		return NSTextBlockFromID(id)
	})
}
// The text lists that contain the paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/textLists
func (p NSParagraphStyle) TextLists() []NSTextList {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("textLists"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextList {
		return NSTextListFromID(id)
	})
}
// The mode for breaking lines in the paragraph that don’t fit within a
// container.
//
// # Discussion
// 
// This property controls how the text system lays out lines that don’t fit
// in its container, such as by truncating with an ellipsis (…) or clipping
// the text. This is different from [NSParagraphStyle.LineBreakStrategy],
// which controls where the system places line breaks in a paragraph.
//
// [NSParagraphStyle.LineBreakStrategy]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineBreakMode
func (p NSParagraphStyle) LineBreakMode() NSLineBreakMode {
	rv := objc.Send[NSLineBreakMode](p.ID, objc.Sel("lineBreakMode"))
	return NSLineBreakMode(rv)
}
// The strategy for breaking lines while laying out paragraphs.
//
// # Discussion
// 
// Line-break strategies are collections of options the system uses to
// determine where to break lines in a paragraph. This is different from
// [LineBreakMode], which controls how to lay out lines of text that don’t
// fit in a container. The system ignores this property if the paragraph
// style’s [LineBreakMode] property specifies a mode that doesn’t support
// multiple lines, such as [NSLineBreakMode.byClipping].
// 
// The default value is [NSLineBreakStrategyNone].
//
// [NSLineBreakMode.byClipping]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byClipping
// [NSLineBreakStrategyNone]: https://developer.apple.com/documentation/AppKit/NSLineBreakStrategy/NSLineBreakStrategyNone
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/lineBreakStrategy-swift.property
func (p NSParagraphStyle) LineBreakStrategy() NSLineBreakStrategy {
	rv := objc.Send[NSLineBreakStrategy](p.ID, objc.Sel("lineBreakStrategy"))
	return NSLineBreakStrategy(rv)
}
// The paragraph’s threshold for hyphenation.
//
// # Discussion
// 
// The system attempts hyphenation when the ratio of the text width (as broken
// without hyphenation) to the width of the line fragment is less than the
// hyphenation factor. When the paragraph’s hyphenation factor is `0.0`, the
// system uses the layout manager’s hyphenation factor instead. The system
// disables hyphenation when both are `0.0`. This property detects the
// user-selected language by examining the first item in [preferredLanguages].
//
// [preferredLanguages]: https://developer.apple.com/documentation/Foundation/NSLocale/preferredLanguages
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/hyphenationFactor
func (p NSParagraphStyle) HyphenationFactor() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("hyphenationFactor"))
	return rv
}
// A Boolean value that indicates whether the paragraph style uses the system
// hyphenation settings.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/usesDefaultHyphenation
func (p NSParagraphStyle) UsesDefaultHyphenation() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}
// The threshold for using tightening as an alternative to truncation.
//
// # Discussion
// 
// When the line break mode specifies truncation, the text system attempts to
// tighten character spacing as an alternative to truncation. Provided that
// the ratio of the text width to the line fragment width doesn’t exceed
// `1.0` + the system sets the tightening factor to this property. Otherwise,
// the system truncates the text at a location determined by the line break
// mode.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/tighteningFactorForTruncation
func (p NSParagraphStyle) TighteningFactorForTruncation() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("tighteningFactorForTruncation"))
	return rv
}
// A Boolean value that indicates whether the system tightens character
// spacing before truncating text.
//
// # Discussion
// 
// When this property is [true], the system tries to reduce the space between
// characters before truncating characters. The system performs this
// tightening in cases where the text wouldn’t otherwise fit in the
// available space. The maximum amount of tightening performed by the system
// is dependent on the font, line width, and other factors.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/allowsDefaultTighteningForTruncation
func (p NSParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}
// The paragraph’s header level for HTML generation.
//
// # Discussion
// 
// If the paragraph isn’t a header, the value is `0`. If the paragraph is a
// header, the value ranges from `1` to `6`, depending on the header’s
// level.
// 
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/headerLevel
func (p NSParagraphStyle) HeaderLevel() int {
	rv := objc.Send[int](p.ID, objc.Sel("headerLevel"))
	return rv
}
// The base writing direction for the paragraph.
//
// # Discussion
// 
// If you the value of this property is [NSWritingDirection.natural], the
// receiver resolves the writing direction to either
// [WritingDirectionLeftToRight] or [WritingDirectionRightToLeft], depending
// on the direction for the user’s language preference setting.
//
// [NSWritingDirection.natural]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/natural
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/baseWritingDirection
func (p NSParagraphStyle) BaseWritingDirection() NSWritingDirection {
	rv := objc.Send[NSWritingDirection](p.ID, objc.Sel("baseWritingDirection"))
	return NSWritingDirection(rv)
}
// The paragraph style of the text.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/paragraphStyle
func (p NSParagraphStyle) ParagraphStyle() foundation.NSString {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("paragraphStyle"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The default paragraph style.
//
// # Discussion
// 
// The default paragraph style has the following default values:
// 
// [Table data omitted]
// 
// See individual method descriptions for explanations of each subattribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/default
func (_NSParagraphStyleClass NSParagraphStyleClass) DefaultParagraphStyle() NSParagraphStyle {
	rv := objc.Send[objc.ID](objc.ID(_NSParagraphStyleClass.class), objc.Sel("defaultParagraphStyle"))
	return NSParagraphStyleFromID(objc.ID(rv))
}

