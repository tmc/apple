// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

import (
	"encoding/binary"
	"unsafe"
)

// C struct types

// ALMXGlyphEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ALMXGlyphEntry
type ALMXGlyphEntry struct {
	GlyphIndexOffset  int16
	HorizontalAdvance int16
	XOffsetToHOrigin  int16
	VerticalAdvance   int16
	YOffsetToVOrigin  int16
}

// ALMXHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ALMXHeader
type ALMXHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [30]byte
}

// Version returns the Version field from the record's packed storage.
func (s *ALMXHeader) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *ALMXHeader) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *ALMXHeader) Flags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *ALMXHeader) SetFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// NMasters returns the NMasters field from the record's packed storage.
func (s *ALMXHeader) NMasters() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetNMasters updates the NMasters field in the record's packed storage.
func (s *ALMXHeader) SetNMasters(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// FirstGlyph returns the FirstGlyph field from the record's packed storage.
func (s *ALMXHeader) FirstGlyph() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetFirstGlyph updates the FirstGlyph field in the record's packed storage.
func (s *ALMXHeader) SetFirstGlyph(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// LastGlyph returns the LastGlyph field from the record's packed storage.
func (s *ALMXHeader) LastGlyph() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetLastGlyph updates the LastGlyph field in the record's packed storage.
func (s *ALMXHeader) SetLastGlyph(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// Lookup returns the Lookup field from the record's packed storage.
func (s *ALMXHeader) Lookup() SFNTLookupTable {
	return *(*SFNTLookupTable)(unsafe.Pointer(&s.storage[12]))
}

// SetLookup updates the Lookup field in the record's packed storage.
func (s *ALMXHeader) SetLookup(v SFNTLookupTable) {
	*(*SFNTLookupTable)(unsafe.Pointer(&s.storage[12])) = v
}

// AnchorPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/AnchorPoint
type AnchorPoint struct {
	X int16
	Y int16
}

// AnchorPointTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/AnchorPointTable
type AnchorPointTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [8]byte
}

// NPoints returns the NPoints field from the record's packed storage.
func (s *AnchorPointTable) NPoints() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNPoints updates the NPoints field in the record's packed storage.
func (s *AnchorPointTable) SetNPoints(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Points returns the Points field from the record's packed storage.
func (s *AnchorPointTable) Points() [1]AnchorPoint {
	return *(*[1]AnchorPoint)(unsafe.Pointer(&s.storage[4]))
}

// SetPoints updates the Points field in the record's packed storage.
func (s *AnchorPointTable) SetPoints(v [1]AnchorPoint) {
	*(*[1]AnchorPoint)(unsafe.Pointer(&s.storage[4])) = v
}

// AnkrTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/AnkrTable
type AnkrTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [12]byte
}

// Version returns the Version field from the record's packed storage.
func (s *AnkrTable) Version() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *AnkrTable) SetVersion(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *AnkrTable) Flags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *AnkrTable) SetFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// LookupTableOffset returns the LookupTableOffset field from the record's packed storage.
func (s *AnkrTable) LookupTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetLookupTableOffset updates the LookupTableOffset field in the record's packed storage.
func (s *AnkrTable) SetLookupTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// AnchorPointTableOffset returns the AnchorPointTableOffset field from the record's packed storage.
func (s *AnkrTable) AnchorPointTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetAnchorPointTableOffset updates the AnchorPointTableOffset field in the record's packed storage.
func (s *AnkrTable) SetAnchorPointTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// BslnFormat0Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat0Part
type BslnFormat0Part struct {
	Deltas [32]int16
}

// BslnFormat1Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat1Part
type BslnFormat1Part struct {
	Deltas      [32]int16
	MappingData SFNTLookupTable
}

// BslnFormat2Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat2Part
type BslnFormat2Part struct {
	StdGlyph  uint16
	CtlPoints [32]int16
}

// BslnFormat3Part
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnFormat3Part
type BslnFormat3Part struct {
	StdGlyph    uint16
	CtlPoints   [32]int16
	MappingData SFNTLookupTable
}

// BslnFormatUnion
type BslnFormatUnion struct {
	Fmt0Part BslnFormat0Part
	Fmt1Part BslnFormat1Part
	Fmt2Part BslnFormat2Part
	Fmt3Part BslnFormat3Part
}

// BslnTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/BslnTable
type BslnTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [92]byte
}

// Version returns the Version field from the record's packed storage.
func (s *BslnTable) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *BslnTable) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Format returns the Format field from the record's packed storage.
func (s *BslnTable) Format() BslnTableFormat {
	return BslnTableFormat(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFormat updates the Format field in the record's packed storage.
func (s *BslnTable) SetFormat(v BslnTableFormat) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// DefaultBaseline returns the DefaultBaseline field from the record's packed storage.
func (s *BslnTable) DefaultBaseline() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetDefaultBaseline updates the DefaultBaseline field in the record's packed storage.
func (s *BslnTable) SetDefaultBaseline(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// Parts returns the Parts field from the record's packed storage.
func (s *BslnTable) Parts() BslnFormatUnion {
	return *(*BslnFormatUnion)(unsafe.Pointer(&s.storage[8]))
}

// SetParts updates the Parts field in the record's packed storage.
func (s *BslnTable) SetParts(v BslnFormatUnion) {
	*(*BslnFormatUnion)(unsafe.Pointer(&s.storage[8])) = v
}

// CTParagraphStyleSetting - This structure is used to alter the paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyleSetting
type CTParagraphStyleSetting struct {
	Spec      CTParagraphStyleSpecifier // The specifier of the setting. See [CTParagraphStyleSpecifier](<https://developer.apple.com/documentation/CoreText/CTParagraphStyleSpecifier>) for possible values.
	ValueSize uintptr                   // The size of the value pointed to by the `value` field. This value must match the size of the value required by the [CTParagraphStyleSpecifier] set in the `spec` field.
	Value     unsafe.Pointer            // A reference to the value of the setting specified by the `spec` field. The value must be in the proper range for the `spec` value and at least as large as the size specified in `valueSize`.

}

// CTRunDelegateCallbacks - A structure holding pointers to callbacks implemented by the run delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegateCallbacks
type CTRunDelegateCallbacks struct {
	Version    int                             // The version number of the callbacks being passed in as a parameter to [CTRunDelegateCreate(_:_:)](<https://developer.apple.com/documentation/CoreText/CTRunDelegateCreate(_:_:)>). The initial version is [kCTRunDelegateVersion1](<https://developer.apple.com/documentation/CoreText/kCTRunDelegateVersion1>).
	Dealloc    CTRunDelegateDeallocateCallback // The callback invoked when the retain count of a CTRunDelegate reaches 0 and the CTRunDelegate is deallocated. This callback may be [NULL].
	GetAscent  CTRunDelegateGetAscentCallback  // The callback invoked to request the run delegate to determine and return the typographic ascent of glyphs in the run. This callback may be [NULL], which is equivalent to a `getAscent` callback that always returns 0.
	GetDescent CTRunDelegateGetDescentCallback // The callback invoked to request the run delegate to determine and return the typographic descent of glyphs in the run. This callback may be [NULL], which is equivalent to a `getDescent` callback that always returns 0.
	GetWidth   CTRunDelegateGetWidthCallback   // The callback invoked to request the run delegate to determine and return the typographic width of glyphs in the run. This callback may be [NULL], which is equivalent to a `getWidth` callback that always returns 0.

}

// FontVariation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/FontVariation
type FontVariation struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [8]byte
}

// Name returns the Name field from the record's packed storage.
func (s *FontVariation) Name() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetName updates the Name field in the record's packed storage.
func (s *FontVariation) SetName(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Value returns the Value field from the record's packed storage.
func (s *FontVariation) Value() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetValue updates the Value field in the record's packed storage.
func (s *FontVariation) SetValue(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// JustDirectionTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustDirectionTable
type JustDirectionTable struct {
	JustClass          uint16
	WidthDeltaClusters uint16
	Postcomp           uint16
	Lookup             SFNTLookupTable
}

// JustPCAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCAction
type JustPCAction struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [16]byte
}

// ActionCount returns the ActionCount field from the record's packed storage.
func (s *JustPCAction) ActionCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetActionCount updates the ActionCount field in the record's packed storage.
func (s *JustPCAction) SetActionCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Actions returns the Actions field from the record's packed storage.
func (s *JustPCAction) Actions() [1]JustPCActionSubrecord {
	return *(*[1]JustPCActionSubrecord)(unsafe.Pointer(&s.storage[4]))
}

// SetActions updates the Actions field in the record's packed storage.
func (s *JustPCAction) SetActions(v [1]JustPCActionSubrecord) {
	*(*[1]JustPCActionSubrecord)(unsafe.Pointer(&s.storage[4])) = v
}

// JustPCActionSubrecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCActionSubrecord
type JustPCActionSubrecord struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [12]byte
}

// TheClass returns the TheClass field from the record's packed storage.
func (s *JustPCActionSubrecord) TheClass() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetTheClass updates the TheClass field in the record's packed storage.
func (s *JustPCActionSubrecord) SetTheClass(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// TheType returns the TheType field from the record's packed storage.
func (s *JustPCActionSubrecord) TheType() JustPCActionType {
	return JustPCActionType(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetTheType updates the TheType field in the record's packed storage.
func (s *JustPCActionSubrecord) SetTheType(v JustPCActionType) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Length returns the Length field from the record's packed storage.
func (s *JustPCActionSubrecord) Length() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *JustPCActionSubrecord) SetLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Data returns the Data field from the record's packed storage.
func (s *JustPCActionSubrecord) Data() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetData updates the Data field in the record's packed storage.
func (s *JustPCActionSubrecord) SetData(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// JustPCConditionalAddAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCConditionalAddAction
type JustPCConditionalAddAction struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [8]byte
}

// SubstThreshold returns the SubstThreshold field from the record's packed storage.
func (s *JustPCConditionalAddAction) SubstThreshold() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSubstThreshold updates the SubstThreshold field in the record's packed storage.
func (s *JustPCConditionalAddAction) SetSubstThreshold(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// AddGlyph returns the AddGlyph field from the record's packed storage.
func (s *JustPCConditionalAddAction) AddGlyph() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetAddGlyph updates the AddGlyph field in the record's packed storage.
func (s *JustPCConditionalAddAction) SetAddGlyph(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// SubstGlyph returns the SubstGlyph field from the record's packed storage.
func (s *JustPCConditionalAddAction) SubstGlyph() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetSubstGlyph updates the SubstGlyph field in the record's packed storage.
func (s *JustPCConditionalAddAction) SetSubstGlyph(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// JustPCDecompositionAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCDecompositionAction
type JustPCDecompositionAction struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [14]byte
}

// LowerLimit returns the LowerLimit field from the record's packed storage.
func (s *JustPCDecompositionAction) LowerLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetLowerLimit updates the LowerLimit field in the record's packed storage.
func (s *JustPCDecompositionAction) SetLowerLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// UpperLimit returns the UpperLimit field from the record's packed storage.
func (s *JustPCDecompositionAction) UpperLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetUpperLimit updates the UpperLimit field in the record's packed storage.
func (s *JustPCDecompositionAction) SetUpperLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Order returns the Order field from the record's packed storage.
func (s *JustPCDecompositionAction) Order() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetOrder updates the Order field in the record's packed storage.
func (s *JustPCDecompositionAction) SetOrder(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// Count returns the Count field from the record's packed storage.
func (s *JustPCDecompositionAction) Count() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetCount updates the Count field in the record's packed storage.
func (s *JustPCDecompositionAction) SetCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// Glyphs returns the Glyphs field from the record's packed storage.
func (s *JustPCDecompositionAction) Glyphs() [1]uint16 {
	return *(*[1]uint16)(unsafe.Pointer(&s.storage[12]))
}

// SetGlyphs updates the Glyphs field in the record's packed storage.
func (s *JustPCDecompositionAction) SetGlyphs(v [1]uint16) {
	*(*[1]uint16)(unsafe.Pointer(&s.storage[12])) = v
}

// JustPCDuctilityAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCDuctilityAction
type JustPCDuctilityAction struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [16]byte
}

// DuctilityAxis returns the DuctilityAxis field from the record's packed storage.
func (s *JustPCDuctilityAction) DuctilityAxis() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetDuctilityAxis updates the DuctilityAxis field in the record's packed storage.
func (s *JustPCDuctilityAction) SetDuctilityAxis(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MinimumLimit returns the MinimumLimit field from the record's packed storage.
func (s *JustPCDuctilityAction) MinimumLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetMinimumLimit updates the MinimumLimit field in the record's packed storage.
func (s *JustPCDuctilityAction) SetMinimumLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// NoStretchValue returns the NoStretchValue field from the record's packed storage.
func (s *JustPCDuctilityAction) NoStretchValue() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetNoStretchValue updates the NoStretchValue field in the record's packed storage.
func (s *JustPCDuctilityAction) SetNoStretchValue(v int32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// MaximumLimit returns the MaximumLimit field from the record's packed storage.
func (s *JustPCDuctilityAction) MaximumLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetMaximumLimit updates the MaximumLimit field in the record's packed storage.
func (s *JustPCDuctilityAction) SetMaximumLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// JustPCGlyphRepeatAddAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPCGlyphRepeatAddAction
type JustPCGlyphRepeatAddAction struct {
	Flags uint16
	Glyph uint16
}

// JustPostcompTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustPostcompTable
type JustPostcompTable struct {
	LookupTable SFNTLookupTable
}

// JustTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustTable
type JustTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [10]byte
}

// Version returns the Version field from the record's packed storage.
func (s *JustTable) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *JustTable) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Format returns the Format field from the record's packed storage.
func (s *JustTable) Format() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFormat updates the Format field in the record's packed storage.
func (s *JustTable) SetFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// HorizHeaderOffset returns the HorizHeaderOffset field from the record's packed storage.
func (s *JustTable) HorizHeaderOffset() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetHorizHeaderOffset updates the HorizHeaderOffset field in the record's packed storage.
func (s *JustTable) SetHorizHeaderOffset(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// VertHeaderOffset returns the VertHeaderOffset field from the record's packed storage.
func (s *JustTable) VertHeaderOffset() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetVertHeaderOffset updates the VertHeaderOffset field in the record's packed storage.
func (s *JustTable) SetVertHeaderOffset(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// JustWidthDeltaEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustWidthDeltaEntry
type JustWidthDeltaEntry struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [24]byte
}

// JustClass returns the JustClass field from the record's packed storage.
func (s *JustWidthDeltaEntry) JustClass() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetJustClass updates the JustClass field in the record's packed storage.
func (s *JustWidthDeltaEntry) SetJustClass(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// BeforeGrowLimit returns the BeforeGrowLimit field from the record's packed storage.
func (s *JustWidthDeltaEntry) BeforeGrowLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetBeforeGrowLimit updates the BeforeGrowLimit field in the record's packed storage.
func (s *JustWidthDeltaEntry) SetBeforeGrowLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// BeforeShrinkLimit returns the BeforeShrinkLimit field from the record's packed storage.
func (s *JustWidthDeltaEntry) BeforeShrinkLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetBeforeShrinkLimit updates the BeforeShrinkLimit field in the record's packed storage.
func (s *JustWidthDeltaEntry) SetBeforeShrinkLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// AfterGrowLimit returns the AfterGrowLimit field from the record's packed storage.
func (s *JustWidthDeltaEntry) AfterGrowLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetAfterGrowLimit updates the AfterGrowLimit field in the record's packed storage.
func (s *JustWidthDeltaEntry) SetAfterGrowLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// AfterShrinkLimit returns the AfterShrinkLimit field from the record's packed storage.
func (s *JustWidthDeltaEntry) AfterShrinkLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetAfterShrinkLimit updates the AfterShrinkLimit field in the record's packed storage.
func (s *JustWidthDeltaEntry) SetAfterShrinkLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// GrowFlags returns the GrowFlags field from the record's packed storage.
func (s *JustWidthDeltaEntry) GrowFlags() JustificationFlags {
	return JustificationFlags(binary.NativeEndian.Uint16(s.storage[20:22]))
}

// SetGrowFlags updates the GrowFlags field in the record's packed storage.
func (s *JustWidthDeltaEntry) SetGrowFlags(v JustificationFlags) {
	binary.NativeEndian.PutUint16(s.storage[20:22], uint16(v))
}

// ShrinkFlags returns the ShrinkFlags field from the record's packed storage.
func (s *JustWidthDeltaEntry) ShrinkFlags() JustificationFlags {
	return JustificationFlags(binary.NativeEndian.Uint16(s.storage[22:24]))
}

// SetShrinkFlags updates the ShrinkFlags field in the record's packed storage.
func (s *JustWidthDeltaEntry) SetShrinkFlags(v JustificationFlags) {
	binary.NativeEndian.PutUint16(s.storage[22:24], uint16(v))
}

// JustWidthDeltaGroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/JustWidthDeltaGroup
type JustWidthDeltaGroup struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [28]byte
}

// Count returns the Count field from the record's packed storage.
func (s *JustWidthDeltaGroup) Count() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetCount updates the Count field in the record's packed storage.
func (s *JustWidthDeltaGroup) SetCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Entries returns the Entries field from the record's packed storage.
func (s *JustWidthDeltaGroup) Entries() [1]JustWidthDeltaEntry {
	return *(*[1]JustWidthDeltaEntry)(unsafe.Pointer(&s.storage[4]))
}

// SetEntries updates the Entries field in the record's packed storage.
func (s *JustWidthDeltaGroup) SetEntries(v [1]JustWidthDeltaEntry) {
	*(*[1]JustWidthDeltaEntry)(unsafe.Pointer(&s.storage[4])) = v
}

// KernFormatSpecificHeader
type KernFormatSpecificHeader struct {
	OrderedList KernOrderedListHeader
	StateTable  KernStateHeader
	SimpleArray KernSimpleArrayHeader
	IndexArray  KernIndexArrayHeader
}

// KernIndexArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernIndexArrayHeader
type KernIndexArrayHeader struct {
	GlyphCount      uint16
	KernValueCount  uint8
	LeftClassCount  uint8
	RightClassCount uint8
	Flags           uint8
	KernValue       [1]int16
	LeftClass       [1]uint8
	RightClass      [1]uint8
	KernIndex       [1]uint8
}

// KernKerningPair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernKerningPair
type KernKerningPair struct {
	Left  uint16
	Right uint16
}

// KernOffsetTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOffsetTable
type KernOffsetTable struct {
	FirstGlyph  uint16
	NGlyphs     uint16
	OffsetTable [1]KernArrayOffset
}

// KernOrderedListEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOrderedListEntry
type KernOrderedListEntry struct {
	Pair  KernKerningPair
	Value KernKerningValue
}

// KernOrderedListHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernOrderedListHeader
type KernOrderedListHeader struct {
	NPairs        uint16
	SearchRange   uint16
	EntrySelector uint16
	RangeShift    uint16
	Table         [1]uint16
}

// KernSimpleArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernSimpleArrayHeader
type KernSimpleArrayHeader struct {
	RowWidth         uint16
	LeftOffsetTable  uint16
	RightOffsetTable uint16
	TheArray         KernArrayOffset
	FirstTable       [1]uint16
}

// KernStateEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernStateEntry
type KernStateEntry struct {
	NewState uint16
	Flags    uint16
}

// KernStateHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernStateHeader
type KernStateHeader struct {
	Header     STHeader
	ValueTable uint16
	FirstTable [1]uint8
}

// KernSubtableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernSubtableHeader
type KernSubtableHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [20]byte
}

// Length returns the Length field from the record's packed storage.
func (s *KernSubtableHeader) Length() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *KernSubtableHeader) SetLength(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// StInfo returns the StInfo field from the record's packed storage.
func (s *KernSubtableHeader) StInfo() KernSubtableInfo {
	return KernSubtableInfo(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetStInfo updates the StInfo field in the record's packed storage.
func (s *KernSubtableHeader) SetStInfo(v KernSubtableInfo) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// TupleIndex returns the TupleIndex field from the record's packed storage.
func (s *KernSubtableHeader) TupleIndex() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetTupleIndex updates the TupleIndex field in the record's packed storage.
func (s *KernSubtableHeader) SetTupleIndex(v int16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// FsHeader returns the FsHeader field from the record's packed storage.
func (s *KernSubtableHeader) FsHeader() KernFormatSpecificHeader {
	return *(*KernFormatSpecificHeader)(unsafe.Pointer(&s.storage[8]))
}

// SetFsHeader updates the FsHeader field in the record's packed storage.
func (s *KernSubtableHeader) SetFsHeader(v KernFormatSpecificHeader) {
	*(*KernFormatSpecificHeader)(unsafe.Pointer(&s.storage[8])) = v
}

// KernTableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernTableHeader
type KernTableHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [10]byte
}

// Version returns the Version field from the record's packed storage.
func (s *KernTableHeader) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *KernTableHeader) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NTables returns the NTables field from the record's packed storage.
func (s *KernTableHeader) NTables() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNTables updates the NTables field in the record's packed storage.
func (s *KernTableHeader) SetNTables(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// FirstSubtable returns the FirstSubtable field from the record's packed storage.
func (s *KernTableHeader) FirstSubtable() [1]uint16 {
	return *(*[1]uint16)(unsafe.Pointer(&s.storage[8]))
}

// SetFirstSubtable updates the FirstSubtable field in the record's packed storage.
func (s *KernTableHeader) SetFirstSubtable(v [1]uint16) {
	*(*[1]uint16)(unsafe.Pointer(&s.storage[8])) = v
}

// KernVersion0Header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernVersion0Header
type KernVersion0Header struct {
	Version       uint16
	NTables       uint16
	FirstSubtable [1]uint16
}

// KernVersion0SubtableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KernVersion0SubtableHeader
type KernVersion0SubtableHeader struct {
	Version  uint16
	Length   uint16
	StInfo   KernSubtableInfo
	FsHeader KernFormatSpecificHeader
}

// KerxAnchorPointAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxAnchorPointAction
type KerxAnchorPointAction struct {
	MarkAnchorPoint uint16
	CurrAnchorPoint uint16
}

// KerxControlPointAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxControlPointAction
type KerxControlPointAction struct {
	MarkControlPoint uint16
	CurrControlPoint uint16
}

// KerxControlPointEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxControlPointEntry
type KerxControlPointEntry struct {
	NewState    uint16
	Flags       uint16
	ActionIndex uint16
}

// KerxControlPointHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxControlPointHeader
type KerxControlPointHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [22]byte
}

// Header returns the Header field from the record's packed storage.
func (s *KerxControlPointHeader) Header() STXHeader {
	return *(*STXHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetHeader updates the Header field in the record's packed storage.
func (s *KerxControlPointHeader) SetHeader(v STXHeader) {
	*(*STXHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// Flags returns the Flags field from the record's packed storage.
func (s *KerxControlPointHeader) Flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *KerxControlPointHeader) SetFlags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// FirstTable returns the FirstTable field from the record's packed storage.
func (s *KerxControlPointHeader) FirstTable() [1]uint8 {
	return *(*[1]uint8)(unsafe.Pointer(&s.storage[20]))
}

// SetFirstTable updates the FirstTable field in the record's packed storage.
func (s *KerxControlPointHeader) SetFirstTable(v [1]uint8) {
	*(*[1]uint8)(unsafe.Pointer(&s.storage[20])) = v
}

// KerxCoordinateAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxCoordinateAction
type KerxCoordinateAction struct {
	MarkX uint16
	MarkY uint16
	CurrX uint16
	CurrY uint16
}

// KerxFormatSpecificHeader
type KerxFormatSpecificHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [24]byte
}

// OrderedList returns the OrderedList field from the record's packed storage.
func (s *KerxFormatSpecificHeader) OrderedList() KerxOrderedListHeader {
	return *(*KerxOrderedListHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetOrderedList updates the OrderedList field in the record's packed storage.
func (s *KerxFormatSpecificHeader) SetOrderedList(v KerxOrderedListHeader) {
	*(*KerxOrderedListHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// StateTable returns the StateTable field from the record's packed storage.
func (s *KerxFormatSpecificHeader) StateTable() KerxStateHeader {
	return *(*KerxStateHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetStateTable updates the StateTable field in the record's packed storage.
func (s *KerxFormatSpecificHeader) SetStateTable(v KerxStateHeader) {
	*(*KerxStateHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// SimpleArray returns the SimpleArray field from the record's packed storage.
func (s *KerxFormatSpecificHeader) SimpleArray() KerxSimpleArrayHeader {
	return *(*KerxSimpleArrayHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetSimpleArray updates the SimpleArray field in the record's packed storage.
func (s *KerxFormatSpecificHeader) SetSimpleArray(v KerxSimpleArrayHeader) {
	*(*KerxSimpleArrayHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// IndexArray returns the IndexArray field from the record's packed storage.
func (s *KerxFormatSpecificHeader) IndexArray() KerxIndexArrayHeader {
	return *(*KerxIndexArrayHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetIndexArray updates the IndexArray field in the record's packed storage.
func (s *KerxFormatSpecificHeader) SetIndexArray(v KerxIndexArrayHeader) {
	*(*KerxIndexArrayHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// ControlPoint returns the ControlPoint field from the record's packed storage.
func (s *KerxFormatSpecificHeader) ControlPoint() KerxControlPointHeader {
	return *(*KerxControlPointHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetControlPoint updates the ControlPoint field in the record's packed storage.
func (s *KerxFormatSpecificHeader) SetControlPoint(v KerxControlPointHeader) {
	*(*KerxControlPointHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// KerxIndexArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxIndexArrayHeader
type KerxIndexArrayHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [24]byte
}

// Flags returns the Flags field from the record's packed storage.
func (s *KerxIndexArrayHeader) Flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *KerxIndexArrayHeader) SetFlags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// RowCount returns the RowCount field from the record's packed storage.
func (s *KerxIndexArrayHeader) RowCount() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetRowCount updates the RowCount field in the record's packed storage.
func (s *KerxIndexArrayHeader) SetRowCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// ColumnCount returns the ColumnCount field from the record's packed storage.
func (s *KerxIndexArrayHeader) ColumnCount() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetColumnCount updates the ColumnCount field in the record's packed storage.
func (s *KerxIndexArrayHeader) SetColumnCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// RowIndexTableOffset returns the RowIndexTableOffset field from the record's packed storage.
func (s *KerxIndexArrayHeader) RowIndexTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetRowIndexTableOffset updates the RowIndexTableOffset field in the record's packed storage.
func (s *KerxIndexArrayHeader) SetRowIndexTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// ColumnIndexTableOffset returns the ColumnIndexTableOffset field from the record's packed storage.
func (s *KerxIndexArrayHeader) ColumnIndexTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetColumnIndexTableOffset updates the ColumnIndexTableOffset field in the record's packed storage.
func (s *KerxIndexArrayHeader) SetColumnIndexTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// KerningArrayOffset returns the KerningArrayOffset field from the record's packed storage.
func (s *KerxIndexArrayHeader) KerningArrayOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetKerningArrayOffset updates the KerningArrayOffset field in the record's packed storage.
func (s *KerxIndexArrayHeader) SetKerningArrayOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// KerningVectorOffset returns the KerningVectorOffset field from the record's packed storage.
func (s *KerxIndexArrayHeader) KerningVectorOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetKerningVectorOffset updates the KerningVectorOffset field in the record's packed storage.
func (s *KerxIndexArrayHeader) SetKerningVectorOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// KerxKerningPair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxKerningPair
type KerxKerningPair struct {
	Left  uint16
	Right uint16
}

// KerxOrderedListEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxOrderedListEntry
type KerxOrderedListEntry struct {
	Pair  KerxKerningPair
	Value KernKerningValue
}

// KerxOrderedListHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxOrderedListHeader
type KerxOrderedListHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [20]byte
}

// NPairs returns the NPairs field from the record's packed storage.
func (s *KerxOrderedListHeader) NPairs() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNPairs updates the NPairs field in the record's packed storage.
func (s *KerxOrderedListHeader) SetNPairs(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// SearchRange returns the SearchRange field from the record's packed storage.
func (s *KerxOrderedListHeader) SearchRange() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetSearchRange updates the SearchRange field in the record's packed storage.
func (s *KerxOrderedListHeader) SetSearchRange(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// EntrySelector returns the EntrySelector field from the record's packed storage.
func (s *KerxOrderedListHeader) EntrySelector() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetEntrySelector updates the EntrySelector field in the record's packed storage.
func (s *KerxOrderedListHeader) SetEntrySelector(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// RangeShift returns the RangeShift field from the record's packed storage.
func (s *KerxOrderedListHeader) RangeShift() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetRangeShift updates the RangeShift field in the record's packed storage.
func (s *KerxOrderedListHeader) SetRangeShift(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Table returns the Table field from the record's packed storage.
func (s *KerxOrderedListHeader) Table() [1]uint32 {
	return *(*[1]uint32)(unsafe.Pointer(&s.storage[16]))
}

// SetTable updates the Table field in the record's packed storage.
func (s *KerxOrderedListHeader) SetTable(v [1]uint32) {
	*(*[1]uint32)(unsafe.Pointer(&s.storage[16])) = v
}

// KerxSimpleArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxSimpleArrayHeader
type KerxSimpleArrayHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [20]byte
}

// RowWidth returns the RowWidth field from the record's packed storage.
func (s *KerxSimpleArrayHeader) RowWidth() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetRowWidth updates the RowWidth field in the record's packed storage.
func (s *KerxSimpleArrayHeader) SetRowWidth(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// LeftOffsetTable returns the LeftOffsetTable field from the record's packed storage.
func (s *KerxSimpleArrayHeader) LeftOffsetTable() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetLeftOffsetTable updates the LeftOffsetTable field in the record's packed storage.
func (s *KerxSimpleArrayHeader) SetLeftOffsetTable(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// RightOffsetTable returns the RightOffsetTable field from the record's packed storage.
func (s *KerxSimpleArrayHeader) RightOffsetTable() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetRightOffsetTable updates the RightOffsetTable field in the record's packed storage.
func (s *KerxSimpleArrayHeader) SetRightOffsetTable(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// TheArray returns the TheArray field from the record's packed storage.
func (s *KerxSimpleArrayHeader) TheArray() KerxArrayOffset {
	return KerxArrayOffset(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetTheArray updates the TheArray field in the record's packed storage.
func (s *KerxSimpleArrayHeader) SetTheArray(v KerxArrayOffset) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// FirstTable returns the FirstTable field from the record's packed storage.
func (s *KerxSimpleArrayHeader) FirstTable() [1]uint32 {
	return *(*[1]uint32)(unsafe.Pointer(&s.storage[16]))
}

// SetFirstTable updates the FirstTable field in the record's packed storage.
func (s *KerxSimpleArrayHeader) SetFirstTable(v [1]uint32) {
	*(*[1]uint32)(unsafe.Pointer(&s.storage[16])) = v
}

// KerxStateEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxStateEntry
type KerxStateEntry struct {
	NewState   uint16
	Flags      uint16
	ValueIndex uint16
}

// KerxStateHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxStateHeader
type KerxStateHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [22]byte
}

// Header returns the Header field from the record's packed storage.
func (s *KerxStateHeader) Header() STXHeader {
	return *(*STXHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetHeader updates the Header field in the record's packed storage.
func (s *KerxStateHeader) SetHeader(v STXHeader) {
	*(*STXHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// ValueTable returns the ValueTable field from the record's packed storage.
func (s *KerxStateHeader) ValueTable() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetValueTable updates the ValueTable field in the record's packed storage.
func (s *KerxStateHeader) SetValueTable(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// FirstTable returns the FirstTable field from the record's packed storage.
func (s *KerxStateHeader) FirstTable() [1]uint8 {
	return *(*[1]uint8)(unsafe.Pointer(&s.storage[20]))
}

// SetFirstTable updates the FirstTable field in the record's packed storage.
func (s *KerxStateHeader) SetFirstTable(v [1]uint8) {
	*(*[1]uint8)(unsafe.Pointer(&s.storage[20])) = v
}

// KerxSubtableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxSubtableHeader
type KerxSubtableHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [36]byte
}

// Length returns the Length field from the record's packed storage.
func (s *KerxSubtableHeader) Length() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *KerxSubtableHeader) SetLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// StInfo returns the StInfo field from the record's packed storage.
func (s *KerxSubtableHeader) StInfo() KerxSubtableCoverage {
	return KerxSubtableCoverage(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetStInfo updates the StInfo field in the record's packed storage.
func (s *KerxSubtableHeader) SetStInfo(v KerxSubtableCoverage) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// TupleCount returns the TupleCount field from the record's packed storage.
func (s *KerxSubtableHeader) TupleCount() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetTupleCount updates the TupleCount field in the record's packed storage.
func (s *KerxSubtableHeader) SetTupleCount(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// FsHeader returns the FsHeader field from the record's packed storage.
func (s *KerxSubtableHeader) FsHeader() KerxFormatSpecificHeader {
	return *(*KerxFormatSpecificHeader)(unsafe.Pointer(&s.storage[12]))
}

// SetFsHeader updates the FsHeader field in the record's packed storage.
func (s *KerxSubtableHeader) SetFsHeader(v KerxFormatSpecificHeader) {
	*(*KerxFormatSpecificHeader)(unsafe.Pointer(&s.storage[12])) = v
}

// KerxTableHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/KerxTableHeader
type KerxTableHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [12]byte
}

// Version returns the Version field from the record's packed storage.
func (s *KerxTableHeader) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *KerxTableHeader) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NTables returns the NTables field from the record's packed storage.
func (s *KerxTableHeader) NTables() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNTables updates the NTables field in the record's packed storage.
func (s *KerxTableHeader) SetNTables(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// FirstSubtable returns the FirstSubtable field from the record's packed storage.
func (s *KerxTableHeader) FirstSubtable() [1]uint32 {
	return *(*[1]uint32)(unsafe.Pointer(&s.storage[8]))
}

// SetFirstSubtable updates the FirstSubtable field in the record's packed storage.
func (s *KerxTableHeader) SetFirstSubtable(v [1]uint32) {
	*(*[1]uint32)(unsafe.Pointer(&s.storage[8])) = v
}

// LcarCaretClassEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LcarCaretClassEntry
type LcarCaretClassEntry struct {
	Count    uint16
	Partials [1]uint16
}

// LcarCaretTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LcarCaretTable
type LcarCaretTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [24]byte
}

// Version returns the Version field from the record's packed storage.
func (s *LcarCaretTable) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *LcarCaretTable) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Format returns the Format field from the record's packed storage.
func (s *LcarCaretTable) Format() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFormat updates the Format field in the record's packed storage.
func (s *LcarCaretTable) SetFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// Lookup returns the Lookup field from the record's packed storage.
func (s *LcarCaretTable) Lookup() SFNTLookupTable {
	return *(*SFNTLookupTable)(unsafe.Pointer(&s.storage[6]))
}

// SetLookup updates the Lookup field in the record's packed storage.
func (s *LcarCaretTable) SetLookup(v SFNTLookupTable) {
	*(*SFNTLookupTable)(unsafe.Pointer(&s.storage[6])) = v
}

// LtagStringRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LtagStringRange
type LtagStringRange struct {
	Offset uint16
	Length uint16
}

// LtagTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/LtagTable
type LtagTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [16]byte
}

// Version returns the Version field from the record's packed storage.
func (s *LtagTable) Version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *LtagTable) SetVersion(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *LtagTable) Flags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *LtagTable) SetFlags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// NumTags returns the NumTags field from the record's packed storage.
func (s *LtagTable) NumTags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetNumTags updates the NumTags field in the record's packed storage.
func (s *LtagTable) SetNumTags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// TagRange returns the TagRange field from the record's packed storage.
func (s *LtagTable) TagRange() [1]LtagStringRange {
	return *(*[1]LtagStringRange)(unsafe.Pointer(&s.storage[12]))
}

// SetTagRange updates the TagRange field in the record's packed storage.
func (s *LtagTable) SetTagRange(v [1]LtagStringRange) {
	*(*[1]LtagStringRange)(unsafe.Pointer(&s.storage[12])) = v
}

// MortChain
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortChain
type MortChain struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [24]byte
}

// DefaultFlags returns the DefaultFlags field from the record's packed storage.
func (s *MortChain) DefaultFlags() MortSubtableMaskFlags {
	return MortSubtableMaskFlags(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetDefaultFlags updates the DefaultFlags field in the record's packed storage.
func (s *MortChain) SetDefaultFlags(v MortSubtableMaskFlags) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Length returns the Length field from the record's packed storage.
func (s *MortChain) Length() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *MortChain) SetLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// NFeatures returns the NFeatures field from the record's packed storage.
func (s *MortChain) NFeatures() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetNFeatures updates the NFeatures field in the record's packed storage.
func (s *MortChain) SetNFeatures(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// NSubtables returns the NSubtables field from the record's packed storage.
func (s *MortChain) NSubtables() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetNSubtables updates the NSubtables field in the record's packed storage.
func (s *MortChain) SetNSubtables(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// FeatureEntries returns the FeatureEntries field from the record's packed storage.
func (s *MortChain) FeatureEntries() [1]MortFeatureEntry {
	return *(*[1]MortFeatureEntry)(unsafe.Pointer(&s.storage[12]))
}

// SetFeatureEntries updates the FeatureEntries field in the record's packed storage.
func (s *MortChain) SetFeatureEntries(v [1]MortFeatureEntry) {
	*(*[1]MortFeatureEntry)(unsafe.Pointer(&s.storage[12])) = v
}

// MortContextualSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortContextualSubtable
type MortContextualSubtable struct {
	Header                  STHeader
	SubstitutionTableOffset uint16
}

// MortFeatureEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortFeatureEntry
type MortFeatureEntry struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [12]byte
}

// FeatureType returns the FeatureType field from the record's packed storage.
func (s *MortFeatureEntry) FeatureType() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetFeatureType updates the FeatureType field in the record's packed storage.
func (s *MortFeatureEntry) SetFeatureType(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// FeatureSelector returns the FeatureSelector field from the record's packed storage.
func (s *MortFeatureEntry) FeatureSelector() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetFeatureSelector updates the FeatureSelector field in the record's packed storage.
func (s *MortFeatureEntry) SetFeatureSelector(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// EnableFlags returns the EnableFlags field from the record's packed storage.
func (s *MortFeatureEntry) EnableFlags() MortSubtableMaskFlags {
	return MortSubtableMaskFlags(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetEnableFlags updates the EnableFlags field in the record's packed storage.
func (s *MortFeatureEntry) SetEnableFlags(v MortSubtableMaskFlags) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// DisableFlags returns the DisableFlags field from the record's packed storage.
func (s *MortFeatureEntry) DisableFlags() MortSubtableMaskFlags {
	return MortSubtableMaskFlags(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetDisableFlags updates the DisableFlags field in the record's packed storage.
func (s *MortFeatureEntry) SetDisableFlags(v MortSubtableMaskFlags) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// MortInsertionSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortInsertionSubtable
type MortInsertionSubtable struct {
	Header STHeader
}

// MortLigatureSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortLigatureSubtable
type MortLigatureSubtable struct {
	Header                    STHeader
	LigatureActionTableOffset uint16
	ComponentTableOffset      uint16
	LigatureTableOffset       uint16
}

// MortRearrangementSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortRearrangementSubtable
type MortRearrangementSubtable struct {
	Header STHeader
}

// MortSpecificSubtable
type MortSpecificSubtable struct {
	Rearrangement MortRearrangementSubtable
	Contextual    MortContextualSubtable
	Ligature      MortLigatureSubtable
	Swash         MortSwashSubtable
	Insertion     MortInsertionSubtable
}

// MortSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortSubtable
type MortSubtable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [26]byte
}

// Length returns the Length field from the record's packed storage.
func (s *MortSubtable) Length() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *MortSubtable) SetLength(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Coverage returns the Coverage field from the record's packed storage.
func (s *MortSubtable) Coverage() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetCoverage updates the Coverage field in the record's packed storage.
func (s *MortSubtable) SetCoverage(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *MortSubtable) Flags() MortSubtableMaskFlags {
	return MortSubtableMaskFlags(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *MortSubtable) SetFlags(v MortSubtableMaskFlags) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// U returns the U field from the record's packed storage.
func (s *MortSubtable) U() MortSpecificSubtable {
	return *(*MortSpecificSubtable)(unsafe.Pointer(&s.storage[8]))
}

// SetU updates the U field in the record's packed storage.
func (s *MortSubtable) SetU(v MortSpecificSubtable) {
	*(*MortSpecificSubtable)(unsafe.Pointer(&s.storage[8])) = v
}

// MortSwashSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortSwashSubtable
type MortSwashSubtable struct {
	Lookup SFNTLookupTable
}

// MortTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MortTable
type MortTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [32]byte
}

// Version returns the Version field from the record's packed storage.
func (s *MortTable) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *MortTable) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NChains returns the NChains field from the record's packed storage.
func (s *MortTable) NChains() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNChains updates the NChains field in the record's packed storage.
func (s *MortTable) SetNChains(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Chains returns the Chains field from the record's packed storage.
func (s *MortTable) Chains() [1]MortChain {
	return *(*[1]MortChain)(unsafe.Pointer(&s.storage[8]))
}

// SetChains updates the Chains field in the record's packed storage.
func (s *MortTable) SetChains(v [1]MortChain) {
	*(*[1]MortChain)(unsafe.Pointer(&s.storage[8])) = v
}

// MorxChain
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxChain
type MorxChain struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [28]byte
}

// DefaultFlags returns the DefaultFlags field from the record's packed storage.
func (s *MorxChain) DefaultFlags() MortSubtableMaskFlags {
	return MortSubtableMaskFlags(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetDefaultFlags updates the DefaultFlags field in the record's packed storage.
func (s *MorxChain) SetDefaultFlags(v MortSubtableMaskFlags) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Length returns the Length field from the record's packed storage.
func (s *MorxChain) Length() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *MorxChain) SetLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// NFeatures returns the NFeatures field from the record's packed storage.
func (s *MorxChain) NFeatures() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetNFeatures updates the NFeatures field in the record's packed storage.
func (s *MorxChain) SetNFeatures(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// NSubtables returns the NSubtables field from the record's packed storage.
func (s *MorxChain) NSubtables() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetNSubtables updates the NSubtables field in the record's packed storage.
func (s *MorxChain) SetNSubtables(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// FeatureEntries returns the FeatureEntries field from the record's packed storage.
func (s *MorxChain) FeatureEntries() [1]MortFeatureEntry {
	return *(*[1]MortFeatureEntry)(unsafe.Pointer(&s.storage[16]))
}

// SetFeatureEntries updates the FeatureEntries field in the record's packed storage.
func (s *MorxChain) SetFeatureEntries(v [1]MortFeatureEntry) {
	*(*[1]MortFeatureEntry)(unsafe.Pointer(&s.storage[16])) = v
}

// MorxContextualSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxContextualSubtable
type MorxContextualSubtable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [20]byte
}

// Header returns the Header field from the record's packed storage.
func (s *MorxContextualSubtable) Header() STXHeader {
	return *(*STXHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetHeader updates the Header field in the record's packed storage.
func (s *MorxContextualSubtable) SetHeader(v STXHeader) {
	*(*STXHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// SubstitutionTableOffset returns the SubstitutionTableOffset field from the record's packed storage.
func (s *MorxContextualSubtable) SubstitutionTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetSubstitutionTableOffset updates the SubstitutionTableOffset field in the record's packed storage.
func (s *MorxContextualSubtable) SetSubstitutionTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// MorxInsertionSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxInsertionSubtable
type MorxInsertionSubtable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [20]byte
}

// Header returns the Header field from the record's packed storage.
func (s *MorxInsertionSubtable) Header() STXHeader {
	return *(*STXHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetHeader updates the Header field in the record's packed storage.
func (s *MorxInsertionSubtable) SetHeader(v STXHeader) {
	*(*STXHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// InsertionGlyphTableOffset returns the InsertionGlyphTableOffset field from the record's packed storage.
func (s *MorxInsertionSubtable) InsertionGlyphTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetInsertionGlyphTableOffset updates the InsertionGlyphTableOffset field in the record's packed storage.
func (s *MorxInsertionSubtable) SetInsertionGlyphTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// MorxLigatureSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxLigatureSubtable
type MorxLigatureSubtable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [28]byte
}

// Header returns the Header field from the record's packed storage.
func (s *MorxLigatureSubtable) Header() STXHeader {
	return *(*STXHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetHeader updates the Header field in the record's packed storage.
func (s *MorxLigatureSubtable) SetHeader(v STXHeader) {
	*(*STXHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// LigatureActionTableOffset returns the LigatureActionTableOffset field from the record's packed storage.
func (s *MorxLigatureSubtable) LigatureActionTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetLigatureActionTableOffset updates the LigatureActionTableOffset field in the record's packed storage.
func (s *MorxLigatureSubtable) SetLigatureActionTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// ComponentTableOffset returns the ComponentTableOffset field from the record's packed storage.
func (s *MorxLigatureSubtable) ComponentTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetComponentTableOffset updates the ComponentTableOffset field in the record's packed storage.
func (s *MorxLigatureSubtable) SetComponentTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// LigatureTableOffset returns the LigatureTableOffset field from the record's packed storage.
func (s *MorxLigatureSubtable) LigatureTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetLigatureTableOffset updates the LigatureTableOffset field in the record's packed storage.
func (s *MorxLigatureSubtable) SetLigatureTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// MorxRearrangementSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxRearrangementSubtable
type MorxRearrangementSubtable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [16]byte
}

// Header returns the Header field from the record's packed storage.
func (s *MorxRearrangementSubtable) Header() STXHeader {
	return *(*STXHeader)(unsafe.Pointer(&s.storage[0]))
}

// SetHeader updates the Header field in the record's packed storage.
func (s *MorxRearrangementSubtable) SetHeader(v STXHeader) {
	*(*STXHeader)(unsafe.Pointer(&s.storage[0])) = v
}

// MorxSpecificSubtable
type MorxSpecificSubtable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [28]byte
}

// Rearrangement returns the Rearrangement field from the record's packed storage.
func (s *MorxSpecificSubtable) Rearrangement() MorxRearrangementSubtable {
	return *(*MorxRearrangementSubtable)(unsafe.Pointer(&s.storage[0]))
}

// SetRearrangement updates the Rearrangement field in the record's packed storage.
func (s *MorxSpecificSubtable) SetRearrangement(v MorxRearrangementSubtable) {
	*(*MorxRearrangementSubtable)(unsafe.Pointer(&s.storage[0])) = v
}

// Contextual returns the Contextual field from the record's packed storage.
func (s *MorxSpecificSubtable) Contextual() MorxContextualSubtable {
	return *(*MorxContextualSubtable)(unsafe.Pointer(&s.storage[0]))
}

// SetContextual updates the Contextual field in the record's packed storage.
func (s *MorxSpecificSubtable) SetContextual(v MorxContextualSubtable) {
	*(*MorxContextualSubtable)(unsafe.Pointer(&s.storage[0])) = v
}

// Ligature returns the Ligature field from the record's packed storage.
func (s *MorxSpecificSubtable) Ligature() MorxLigatureSubtable {
	return *(*MorxLigatureSubtable)(unsafe.Pointer(&s.storage[0]))
}

// SetLigature updates the Ligature field in the record's packed storage.
func (s *MorxSpecificSubtable) SetLigature(v MorxLigatureSubtable) {
	*(*MorxLigatureSubtable)(unsafe.Pointer(&s.storage[0])) = v
}

// Swash returns the Swash field from the record's packed storage.
func (s *MorxSpecificSubtable) Swash() MortSwashSubtable {
	return *(*MortSwashSubtable)(unsafe.Pointer(&s.storage[0]))
}

// SetSwash updates the Swash field in the record's packed storage.
func (s *MorxSpecificSubtable) SetSwash(v MortSwashSubtable) {
	*(*MortSwashSubtable)(unsafe.Pointer(&s.storage[0])) = v
}

// Insertion returns the Insertion field from the record's packed storage.
func (s *MorxSpecificSubtable) Insertion() MorxInsertionSubtable {
	return *(*MorxInsertionSubtable)(unsafe.Pointer(&s.storage[0]))
}

// SetInsertion updates the Insertion field in the record's packed storage.
func (s *MorxSpecificSubtable) SetInsertion(v MorxInsertionSubtable) {
	*(*MorxInsertionSubtable)(unsafe.Pointer(&s.storage[0])) = v
}

// MorxSubtable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxSubtable
type MorxSubtable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [40]byte
}

// Length returns the Length field from the record's packed storage.
func (s *MorxSubtable) Length() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *MorxSubtable) SetLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Coverage returns the Coverage field from the record's packed storage.
func (s *MorxSubtable) Coverage() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetCoverage updates the Coverage field in the record's packed storage.
func (s *MorxSubtable) SetCoverage(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *MorxSubtable) Flags() MortSubtableMaskFlags {
	return MortSubtableMaskFlags(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *MorxSubtable) SetFlags(v MortSubtableMaskFlags) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// U returns the U field from the record's packed storage.
func (s *MorxSubtable) U() MorxSpecificSubtable {
	return *(*MorxSpecificSubtable)(unsafe.Pointer(&s.storage[12]))
}

// SetU updates the U field in the record's packed storage.
func (s *MorxSubtable) SetU(v MorxSpecificSubtable) {
	*(*MorxSpecificSubtable)(unsafe.Pointer(&s.storage[12])) = v
}

// MorxTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/MorxTable
type MorxTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [36]byte
}

// Version returns the Version field from the record's packed storage.
func (s *MorxTable) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *MorxTable) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NChains returns the NChains field from the record's packed storage.
func (s *MorxTable) NChains() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNChains updates the NChains field in the record's packed storage.
func (s *MorxTable) SetNChains(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Chains returns the Chains field from the record's packed storage.
func (s *MorxTable) Chains() [1]MorxChain {
	return *(*[1]MorxChain)(unsafe.Pointer(&s.storage[8]))
}

// SetChains updates the Chains field in the record's packed storage.
func (s *MorxTable) SetChains(v [1]MorxChain) {
	*(*[1]MorxChain)(unsafe.Pointer(&s.storage[8])) = v
}

// OpbdSideValues
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/OpbdSideValues
type OpbdSideValues struct {
	LeftSideShift   int16
	TopSideShift    int16
	RightSideShift  int16
	BottomSideShift int16
}

// OpbdTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/OpbdTable
type OpbdTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [24]byte
}

// Version returns the Version field from the record's packed storage.
func (s *OpbdTable) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *OpbdTable) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Format returns the Format field from the record's packed storage.
func (s *OpbdTable) Format() OpbdTableFormat {
	return OpbdTableFormat(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFormat updates the Format field in the record's packed storage.
func (s *OpbdTable) SetFormat(v OpbdTableFormat) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// LookupTable returns the LookupTable field from the record's packed storage.
func (s *OpbdTable) LookupTable() SFNTLookupTable {
	return *(*SFNTLookupTable)(unsafe.Pointer(&s.storage[6]))
}

// SetLookupTable updates the LookupTable field in the record's packed storage.
func (s *OpbdTable) SetLookupTable(v SFNTLookupTable) {
	*(*SFNTLookupTable)(unsafe.Pointer(&s.storage[6])) = v
}

// PropLookupSegment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropLookupSegment
type PropLookupSegment struct {
	LastGlyph  uint16
	FirstGlyph uint16
	Value      uint16
}

// PropLookupSingle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropLookupSingle
type PropLookupSingle struct {
	Glyph uint16
	Props PropCharProperties
}

// PropTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/PropTable
type PropTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [26]byte
}

// Version returns the Version field from the record's packed storage.
func (s *PropTable) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *PropTable) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Format returns the Format field from the record's packed storage.
func (s *PropTable) Format() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFormat updates the Format field in the record's packed storage.
func (s *PropTable) SetFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// DefaultProps returns the DefaultProps field from the record's packed storage.
func (s *PropTable) DefaultProps() PropCharProperties {
	return PropCharProperties(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetDefaultProps updates the DefaultProps field in the record's packed storage.
func (s *PropTable) SetDefaultProps(v PropCharProperties) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// Lookup returns the Lookup field from the record's packed storage.
func (s *PropTable) Lookup() SFNTLookupTable {
	return *(*SFNTLookupTable)(unsafe.Pointer(&s.storage[8]))
}

// SetLookup updates the Lookup field in the record's packed storage.
func (s *PropTable) SetLookup(v SFNTLookupTable) {
	*(*SFNTLookupTable)(unsafe.Pointer(&s.storage[8])) = v
}

// ROTAGlyphEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ROTAGlyphEntry
type ROTAGlyphEntry struct {
	GlyphIndexOffset int16
	HBaselineOffset  int16
	VBaselineOffset  int16
}

// ROTAHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/ROTAHeader
type ROTAHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [30]byte
}

// Version returns the Version field from the record's packed storage.
func (s *ROTAHeader) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *ROTAHeader) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *ROTAHeader) Flags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *ROTAHeader) SetFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// NMasters returns the NMasters field from the record's packed storage.
func (s *ROTAHeader) NMasters() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetNMasters updates the NMasters field in the record's packed storage.
func (s *ROTAHeader) SetNMasters(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// FirstGlyph returns the FirstGlyph field from the record's packed storage.
func (s *ROTAHeader) FirstGlyph() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetFirstGlyph updates the FirstGlyph field in the record's packed storage.
func (s *ROTAHeader) SetFirstGlyph(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// LastGlyph returns the LastGlyph field from the record's packed storage.
func (s *ROTAHeader) LastGlyph() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetLastGlyph updates the LastGlyph field in the record's packed storage.
func (s *ROTAHeader) SetLastGlyph(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// Lookup returns the Lookup field from the record's packed storage.
func (s *ROTAHeader) Lookup() SFNTLookupTable {
	return *(*SFNTLookupTable)(unsafe.Pointer(&s.storage[12]))
}

// SetLookup updates the Lookup field in the record's packed storage.
func (s *ROTAHeader) SetLookup(v SFNTLookupTable) {
	*(*SFNTLookupTable)(unsafe.Pointer(&s.storage[12])) = v
}

// SFNTLookupArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupArrayHeader
type SFNTLookupArrayHeader struct {
	LookupValues [1]SFNTLookupValue
}

// SFNTLookupBinarySearchHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupBinarySearchHeader
type SFNTLookupBinarySearchHeader struct {
	UnitSize      uint16
	NUnits        uint16
	SearchRange   uint16
	EntrySelector uint16
	RangeShift    uint16
}

// SFNTLookupFormatSpecificHeader
type SFNTLookupFormatSpecificHeader struct {
	TheArray     SFNTLookupArrayHeader
	Segment      SFNTLookupSegmentHeader
	Single       SFNTLookupSingleHeader
	TrimmedArray SFNTLookupTrimmedArrayHeader
	Vector       SFNTLookupVectorHeader
}

// SFNTLookupSegment
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSegment
type SFNTLookupSegment struct {
	LastGlyph  uint16
	FirstGlyph uint16
	Value      [1]uint16
}

// SFNTLookupSegmentHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSegmentHeader
type SFNTLookupSegmentHeader struct {
	BinSearch SFNTLookupBinarySearchHeader
	Segments  [1]SFNTLookupSegment
}

// SFNTLookupSingle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSingle
type SFNTLookupSingle struct {
	Glyph uint16
	Value [1]uint16
}

// SFNTLookupSingleHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupSingleHeader
type SFNTLookupSingleHeader struct {
	BinSearch SFNTLookupBinarySearchHeader
	Entries   [1]SFNTLookupSingle
}

// SFNTLookupTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupTable
type SFNTLookupTable struct {
	Format   SFNTLookupTableFormat
	FsHeader SFNTLookupFormatSpecificHeader
}

// SFNTLookupTrimmedArrayHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupTrimmedArrayHeader
type SFNTLookupTrimmedArrayHeader struct {
	FirstGlyph uint16
	Count      uint16
	ValueArray [1]SFNTLookupValue
}

// SFNTLookupVectorHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/SFNTLookupVectorHeader
type SFNTLookupVectorHeader struct {
	ValueSize  uint16
	FirstGlyph uint16
	Count      uint16
	Values     [1]uint8
}

// STClassTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STClassTable
type STClassTable struct {
	FirstGlyph uint16
	NGlyphs    uint16
	Classes    [1]STClass
}

// STEntryOne
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryOne
type STEntryOne struct {
	NewState uint16
	Flags    uint16
	Offset1  uint16
}

// STEntryTwo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryTwo
type STEntryTwo struct {
	NewState uint16
	Flags    uint16
	Offset1  uint16
	Offset2  uint16
}

// STEntryZero
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STEntryZero
type STEntryZero struct {
	NewState uint16
	Flags    uint16
}

// STHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STHeader
type STHeader struct {
	Filler           uint8
	NClasses         STClass
	ClassTableOffset uint16
	StateArrayOffset uint16
	EntryTableOffset uint16
}

// STXEntryOne
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryOne
type STXEntryOne struct {
	NewState STXStateIndex
	Flags    uint16
	Index1   uint16
}

// STXEntryTwo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryTwo
type STXEntryTwo struct {
	NewState STXStateIndex
	Flags    uint16
	Index1   uint16
	Index2   uint16
}

// STXEntryZero
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXEntryZero
type STXEntryZero struct {
	NewState STXStateIndex
	Flags    uint16
}

// STXHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/STXHeader
type STXHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [16]byte
}

// NClasses returns the NClasses field from the record's packed storage.
func (s *STXHeader) NClasses() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNClasses updates the NClasses field in the record's packed storage.
func (s *STXHeader) SetNClasses(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// ClassTableOffset returns the ClassTableOffset field from the record's packed storage.
func (s *STXHeader) ClassTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetClassTableOffset updates the ClassTableOffset field in the record's packed storage.
func (s *STXHeader) SetClassTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// StateArrayOffset returns the StateArrayOffset field from the record's packed storage.
func (s *STXHeader) StateArrayOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetStateArrayOffset updates the StateArrayOffset field in the record's packed storage.
func (s *STXHeader) SetStateArrayOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// EntryTableOffset returns the EntryTableOffset field from the record's packed storage.
func (s *STXHeader) EntryTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetEntryTableOffset updates the EntryTableOffset field in the record's packed storage.
func (s *STXHeader) SetEntryTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// TrakTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakTable
type TrakTable struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [10]byte
}

// Version returns the Version field from the record's packed storage.
func (s *TrakTable) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *TrakTable) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Format returns the Format field from the record's packed storage.
func (s *TrakTable) Format() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFormat updates the Format field in the record's packed storage.
func (s *TrakTable) SetFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// HorizOffset returns the HorizOffset field from the record's packed storage.
func (s *TrakTable) HorizOffset() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetHorizOffset updates the HorizOffset field in the record's packed storage.
func (s *TrakTable) SetHorizOffset(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// VertOffset returns the VertOffset field from the record's packed storage.
func (s *TrakTable) VertOffset() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetVertOffset updates the VertOffset field in the record's packed storage.
func (s *TrakTable) SetVertOffset(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// TrakTableData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakTableData
type TrakTableData struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [16]byte
}

// NTracks returns the NTracks field from the record's packed storage.
func (s *TrakTableData) NTracks() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetNTracks updates the NTracks field in the record's packed storage.
func (s *TrakTableData) SetNTracks(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// NSizes returns the NSizes field from the record's packed storage.
func (s *TrakTableData) NSizes() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetNSizes updates the NSizes field in the record's packed storage.
func (s *TrakTableData) SetNSizes(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// SizeTableOffset returns the SizeTableOffset field from the record's packed storage.
func (s *TrakTableData) SizeTableOffset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetSizeTableOffset updates the SizeTableOffset field in the record's packed storage.
func (s *TrakTableData) SetSizeTableOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// TrakTable returns the TrakTable field from the record's packed storage.
func (s *TrakTableData) TrakTable() [1]TrakTableEntry {
	return *(*[1]TrakTableEntry)(unsafe.Pointer(&s.storage[8]))
}

// SetTrakTable updates the TrakTable field in the record's packed storage.
func (s *TrakTableData) SetTrakTable(v [1]TrakTableEntry) {
	*(*[1]TrakTableEntry)(unsafe.Pointer(&s.storage[8])) = v
}

// TrakTableEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/TrakTableEntry
type TrakTableEntry struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [8]byte
}

// Track returns the Track field from the record's packed storage.
func (s *TrakTableEntry) Track() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetTrack updates the Track field in the record's packed storage.
func (s *TrakTableEntry) SetTrack(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NameTableIndex returns the NameTableIndex field from the record's packed storage.
func (s *TrakTableEntry) NameTableIndex() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetNameTableIndex updates the NameTableIndex field in the record's packed storage.
func (s *TrakTableEntry) SetNameTableIndex(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// SizesOffset returns the SizesOffset field from the record's packed storage.
func (s *TrakTableEntry) SizesOffset() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetSizesOffset updates the SizesOffset field in the record's packed storage.
func (s *TrakTableEntry) SetSizesOffset(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// SfntCMapEncoding
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapEncoding
type SfntCMapEncoding struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [8]byte
}

// PlatformID returns the PlatformID field from the record's packed storage.
func (s *SfntCMapEncoding) PlatformID() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetPlatformID updates the PlatformID field in the record's packed storage.
func (s *SfntCMapEncoding) SetPlatformID(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// ScriptID returns the ScriptID field from the record's packed storage.
func (s *SfntCMapEncoding) ScriptID() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetScriptID updates the ScriptID field in the record's packed storage.
func (s *SfntCMapEncoding) SetScriptID(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Offset returns the Offset field from the record's packed storage.
func (s *SfntCMapEncoding) Offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetOffset updates the Offset field in the record's packed storage.
func (s *SfntCMapEncoding) SetOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// SfntCMapExtendedSubHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapExtendedSubHeader
type SfntCMapExtendedSubHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [12]byte
}

// Format returns the Format field from the record's packed storage.
func (s *SfntCMapExtendedSubHeader) Format() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetFormat updates the Format field in the record's packed storage.
func (s *SfntCMapExtendedSubHeader) SetFormat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *SfntCMapExtendedSubHeader) Reserved() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *SfntCMapExtendedSubHeader) SetReserved(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Length returns the Length field from the record's packed storage.
func (s *SfntCMapExtendedSubHeader) Length() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *SfntCMapExtendedSubHeader) SetLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Language returns the Language field from the record's packed storage.
func (s *SfntCMapExtendedSubHeader) Language() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetLanguage updates the Language field in the record's packed storage.
func (s *SfntCMapExtendedSubHeader) SetLanguage(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// SfntCMapHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapHeader
type SfntCMapHeader struct {
	Version   uint16
	NumTables uint16
	Encoding  [1]SfntCMapEncoding
}

// SfntCMapSubHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntCMapSubHeader
type SfntCMapSubHeader struct {
	Format     uint16
	Length     uint16
	LanguageID uint16
}

// SfntDescriptorHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntDescriptorHeader
type SfntDescriptorHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [16]byte
}

// Version returns the Version field from the record's packed storage.
func (s *SfntDescriptorHeader) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *SfntDescriptorHeader) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// DescriptorCount returns the DescriptorCount field from the record's packed storage.
func (s *SfntDescriptorHeader) DescriptorCount() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetDescriptorCount updates the DescriptorCount field in the record's packed storage.
func (s *SfntDescriptorHeader) SetDescriptorCount(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Descriptor returns the Descriptor field from the record's packed storage.
func (s *SfntDescriptorHeader) Descriptor() [1]SfntFontDescriptor {
	return *(*[1]SfntFontDescriptor)(unsafe.Pointer(&s.storage[8]))
}

// SetDescriptor updates the Descriptor field in the record's packed storage.
func (s *SfntDescriptorHeader) SetDescriptor(v [1]SfntFontDescriptor) {
	*(*[1]SfntFontDescriptor)(unsafe.Pointer(&s.storage[8])) = v
}

// SfntDirectory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntDirectory
type SfntDirectory struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [28]byte
}

// Format returns the Format field from the record's packed storage.
func (s *SfntDirectory) Format() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetFormat updates the Format field in the record's packed storage.
func (s *SfntDirectory) SetFormat(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NumOffsets returns the NumOffsets field from the record's packed storage.
func (s *SfntDirectory) NumOffsets() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetNumOffsets updates the NumOffsets field in the record's packed storage.
func (s *SfntDirectory) SetNumOffsets(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// SearchRange returns the SearchRange field from the record's packed storage.
func (s *SfntDirectory) SearchRange() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetSearchRange updates the SearchRange field in the record's packed storage.
func (s *SfntDirectory) SetSearchRange(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// EntrySelector returns the EntrySelector field from the record's packed storage.
func (s *SfntDirectory) EntrySelector() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetEntrySelector updates the EntrySelector field in the record's packed storage.
func (s *SfntDirectory) SetEntrySelector(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// RangeShift returns the RangeShift field from the record's packed storage.
func (s *SfntDirectory) RangeShift() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetRangeShift updates the RangeShift field in the record's packed storage.
func (s *SfntDirectory) SetRangeShift(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// Table returns the Table field from the record's packed storage.
func (s *SfntDirectory) Table() [1]SfntDirectoryEntry {
	return *(*[1]SfntDirectoryEntry)(unsafe.Pointer(&s.storage[12]))
}

// SetTable updates the Table field in the record's packed storage.
func (s *SfntDirectory) SetTable(v [1]SfntDirectoryEntry) {
	*(*[1]SfntDirectoryEntry)(unsafe.Pointer(&s.storage[12])) = v
}

// SfntDirectoryEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntDirectoryEntry
type SfntDirectoryEntry struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [16]byte
}

// TableTag returns the TableTag field from the record's packed storage.
func (s *SfntDirectoryEntry) TableTag() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetTableTag updates the TableTag field in the record's packed storage.
func (s *SfntDirectoryEntry) SetTableTag(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// CheckSum returns the CheckSum field from the record's packed storage.
func (s *SfntDirectoryEntry) CheckSum() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetCheckSum updates the CheckSum field in the record's packed storage.
func (s *SfntDirectoryEntry) SetCheckSum(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// Offset returns the Offset field from the record's packed storage.
func (s *SfntDirectoryEntry) Offset() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetOffset updates the Offset field in the record's packed storage.
func (s *SfntDirectoryEntry) SetOffset(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Length returns the Length field from the record's packed storage.
func (s *SfntDirectoryEntry) Length() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *SfntDirectoryEntry) SetLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// SfntFeatureHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFeatureHeader
type SfntFeatureHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [32]byte
}

// Version returns the Version field from the record's packed storage.
func (s *SfntFeatureHeader) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *SfntFeatureHeader) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// FeatureNameCount returns the FeatureNameCount field from the record's packed storage.
func (s *SfntFeatureHeader) FeatureNameCount() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFeatureNameCount updates the FeatureNameCount field in the record's packed storage.
func (s *SfntFeatureHeader) SetFeatureNameCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// FeatureSetCount returns the FeatureSetCount field from the record's packed storage.
func (s *SfntFeatureHeader) FeatureSetCount() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetFeatureSetCount updates the FeatureSetCount field in the record's packed storage.
func (s *SfntFeatureHeader) SetFeatureSetCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *SfntFeatureHeader) Reserved() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *SfntFeatureHeader) SetReserved(v int32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// Names returns the Names field from the record's packed storage.
func (s *SfntFeatureHeader) Names() [1]SfntFeatureName {
	return *(*[1]SfntFeatureName)(unsafe.Pointer(&s.storage[12]))
}

// SetNames updates the Names field in the record's packed storage.
func (s *SfntFeatureHeader) SetNames(v [1]SfntFeatureName) {
	*(*[1]SfntFeatureName)(unsafe.Pointer(&s.storage[12])) = v
}

// Settings returns the Settings field from the record's packed storage.
func (s *SfntFeatureHeader) Settings() [1]SfntFontFeatureSetting {
	return *(*[1]SfntFontFeatureSetting)(unsafe.Pointer(&s.storage[24]))
}

// SetSettings updates the Settings field in the record's packed storage.
func (s *SfntFeatureHeader) SetSettings(v [1]SfntFontFeatureSetting) {
	*(*[1]SfntFontFeatureSetting)(unsafe.Pointer(&s.storage[24])) = v
}

// Runs returns the Runs field from the record's packed storage.
func (s *SfntFeatureHeader) Runs() [1]SfntFontRunFeature {
	return *(*[1]SfntFontRunFeature)(unsafe.Pointer(&s.storage[28]))
}

// SetRuns updates the Runs field in the record's packed storage.
func (s *SfntFeatureHeader) SetRuns(v [1]SfntFontRunFeature) {
	*(*[1]SfntFontRunFeature)(unsafe.Pointer(&s.storage[28])) = v
}

// SfntFeatureName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFeatureName
type SfntFeatureName struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [12]byte
}

// FeatureType returns the FeatureType field from the record's packed storage.
func (s *SfntFeatureName) FeatureType() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetFeatureType updates the FeatureType field in the record's packed storage.
func (s *SfntFeatureName) SetFeatureType(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// SettingCount returns the SettingCount field from the record's packed storage.
func (s *SfntFeatureName) SettingCount() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetSettingCount updates the SettingCount field in the record's packed storage.
func (s *SfntFeatureName) SetSettingCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// OffsetToSettings returns the OffsetToSettings field from the record's packed storage.
func (s *SfntFeatureName) OffsetToSettings() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetOffsetToSettings updates the OffsetToSettings field in the record's packed storage.
func (s *SfntFeatureName) SetOffsetToSettings(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// FeatureFlags returns the FeatureFlags field from the record's packed storage.
func (s *SfntFeatureName) FeatureFlags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetFeatureFlags updates the FeatureFlags field in the record's packed storage.
func (s *SfntFeatureName) SetFeatureFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// NameID returns the NameID field from the record's packed storage.
func (s *SfntFeatureName) NameID() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetNameID updates the NameID field in the record's packed storage.
func (s *SfntFeatureName) SetNameID(v int16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// SfntFontDescriptor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFontDescriptor
type SfntFontDescriptor struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [8]byte
}

// Name returns the Name field from the record's packed storage.
func (s *SfntFontDescriptor) Name() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetName updates the Name field in the record's packed storage.
func (s *SfntFontDescriptor) SetName(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Value returns the Value field from the record's packed storage.
func (s *SfntFontDescriptor) Value() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetValue updates the Value field in the record's packed storage.
func (s *SfntFontDescriptor) SetValue(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// SfntFontFeatureSetting
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFontFeatureSetting
type SfntFontFeatureSetting struct {
	Setting uint16
	NameID  int16
}

// SfntFontRunFeature
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntFontRunFeature
type SfntFontRunFeature struct {
	FeatureType uint16
	Setting     uint16
}

// SfntInstance
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntInstance
type SfntInstance struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [8]byte
}

// NameID returns the NameID field from the record's packed storage.
func (s *SfntInstance) NameID() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetNameID updates the NameID field in the record's packed storage.
func (s *SfntInstance) SetNameID(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *SfntInstance) Flags() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *SfntInstance) SetFlags(v int16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// Coord returns the Coord field from the record's packed storage.
func (s *SfntInstance) Coord() [1]int32 {
	return *(*[1]int32)(unsafe.Pointer(&s.storage[4]))
}

// SetCoord updates the Coord field in the record's packed storage.
func (s *SfntInstance) SetCoord(v [1]int32) {
	*(*[1]int32)(unsafe.Pointer(&s.storage[4])) = v
}

// SfntNameHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntNameHeader
type SfntNameHeader struct {
	Format       uint16
	Count        uint16
	StringOffset uint16
	Rec          [1]SfntNameRecord
}

// SfntNameRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntNameRecord
type SfntNameRecord struct {
	PlatformID uint16
	ScriptID   uint16
	LanguageID uint16
	NameID     uint16
	Length     uint16
	Offset     uint16
}

// SfntVariationAxis
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntVariationAxis
type SfntVariationAxis struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [20]byte
}

// AxisTag returns the AxisTag field from the record's packed storage.
func (s *SfntVariationAxis) AxisTag() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetAxisTag updates the AxisTag field in the record's packed storage.
func (s *SfntVariationAxis) SetAxisTag(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// MinValue returns the MinValue field from the record's packed storage.
func (s *SfntVariationAxis) MinValue() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetMinValue updates the MinValue field in the record's packed storage.
func (s *SfntVariationAxis) SetMinValue(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// DefaultValue returns the DefaultValue field from the record's packed storage.
func (s *SfntVariationAxis) DefaultValue() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetDefaultValue updates the DefaultValue field in the record's packed storage.
func (s *SfntVariationAxis) SetDefaultValue(v int32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// MaxValue returns the MaxValue field from the record's packed storage.
func (s *SfntVariationAxis) MaxValue() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetMaxValue updates the MaxValue field in the record's packed storage.
func (s *SfntVariationAxis) SetMaxValue(v int32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *SfntVariationAxis) Flags() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[16:18]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *SfntVariationAxis) SetFlags(v int16) {
	binary.NativeEndian.PutUint16(s.storage[16:18], uint16(v))
}

// NameID returns the NameID field from the record's packed storage.
func (s *SfntVariationAxis) NameID() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[18:20]))
}

// SetNameID updates the NameID field in the record's packed storage.
func (s *SfntVariationAxis) SetNameID(v int16) {
	binary.NativeEndian.PutUint16(s.storage[18:20], uint16(v))
}

// SfntVariationHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreText/sfntVariationHeader
type SfntVariationHeader struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.

	// A storage array alone has alignment 1. This zero-length array carries
	// the alignment C measured without contributing any size, so an
	// embedding record places this one where C does.
	_       [0]uint16
	storage [44]byte
}

// Version returns the Version field from the record's packed storage.
func (s *SfntVariationHeader) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *SfntVariationHeader) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// OffsetToData returns the OffsetToData field from the record's packed storage.
func (s *SfntVariationHeader) OffsetToData() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetOffsetToData updates the OffsetToData field in the record's packed storage.
func (s *SfntVariationHeader) SetOffsetToData(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// CountSizePairs returns the CountSizePairs field from the record's packed storage.
func (s *SfntVariationHeader) CountSizePairs() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetCountSizePairs updates the CountSizePairs field in the record's packed storage.
func (s *SfntVariationHeader) SetCountSizePairs(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// AxisCount returns the AxisCount field from the record's packed storage.
func (s *SfntVariationHeader) AxisCount() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetAxisCount updates the AxisCount field in the record's packed storage.
func (s *SfntVariationHeader) SetAxisCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// AxisSize returns the AxisSize field from the record's packed storage.
func (s *SfntVariationHeader) AxisSize() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetAxisSize updates the AxisSize field in the record's packed storage.
func (s *SfntVariationHeader) SetAxisSize(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// InstanceCount returns the InstanceCount field from the record's packed storage.
func (s *SfntVariationHeader) InstanceCount() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[12:14]))
}

// SetInstanceCount updates the InstanceCount field in the record's packed storage.
func (s *SfntVariationHeader) SetInstanceCount(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[12:14], uint16(v))
}

// InstanceSize returns the InstanceSize field from the record's packed storage.
func (s *SfntVariationHeader) InstanceSize() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[14:16]))
}

// SetInstanceSize updates the InstanceSize field in the record's packed storage.
func (s *SfntVariationHeader) SetInstanceSize(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[14:16], uint16(v))
}

// Axis returns the Axis field from the record's packed storage.
func (s *SfntVariationHeader) Axis() [1]SfntVariationAxis {
	return *(*[1]SfntVariationAxis)(unsafe.Pointer(&s.storage[16]))
}

// SetAxis updates the Axis field in the record's packed storage.
func (s *SfntVariationHeader) SetAxis(v [1]SfntVariationAxis) {
	*(*[1]SfntVariationAxis)(unsafe.Pointer(&s.storage[16])) = v
}

// Instance returns the Instance field from the record's packed storage.
func (s *SfntVariationHeader) Instance() [1]SfntInstance {
	return *(*[1]SfntInstance)(unsafe.Pointer(&s.storage[36]))
}

// SetInstance updates the Instance field in the record's packed storage.
func (s *SfntVariationHeader) SetInstance(v [1]SfntInstance) {
	*(*[1]SfntInstance)(unsafe.Pointer(&s.storage[36])) = v
}
