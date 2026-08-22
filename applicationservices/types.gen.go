// Code generated from Apple documentation for ApplicationServices. DO NOT EDIT.

package applicationservices

import (
	"encoding/binary"
	"math"
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// C struct types

// ATSFlatDataFontNameDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatafontnamedataheader
type ATSFlatDataFontNameDataHeader struct {
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

// NameSpecifierType returns the NameSpecifierType field from the record's packed storage.
func (s *ATSFlatDataFontNameDataHeader) NameSpecifierType() ATSFlatDataFontSpeciferType {
	return ATSFlatDataFontSpeciferType(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNameSpecifierType updates the NameSpecifierType field in the record's packed storage.
func (s *ATSFlatDataFontNameDataHeader) SetNameSpecifierType(v ATSFlatDataFontSpeciferType) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NameSpecifierSize returns the NameSpecifierSize field from the record's packed storage.
func (s *ATSFlatDataFontNameDataHeader) NameSpecifierSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNameSpecifierSize updates the NameSpecifierSize field in the record's packed storage.
func (s *ATSFlatDataFontNameDataHeader) SetNameSpecifierSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// ATSFlatDataFontSpecRawNameData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatafontspecrawnamedata
type ATSFlatDataFontSpecRawNameData struct {
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

// FontNameType returns the FontNameType field from the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) FontNameType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetFontNameType updates the FontNameType field in the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) SetFontNameType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// FontNamePlatform returns the FontNamePlatform field from the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) FontNamePlatform() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetFontNamePlatform updates the FontNamePlatform field in the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) SetFontNamePlatform(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// FontNameScript returns the FontNameScript field from the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) FontNameScript() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetFontNameScript updates the FontNameScript field in the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) SetFontNameScript(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// FontNameLanguage returns the FontNameLanguage field from the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) FontNameLanguage() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetFontNameLanguage updates the FontNameLanguage field in the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) SetFontNameLanguage(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// FontNameLength returns the FontNameLength field from the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) FontNameLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetFontNameLength updates the FontNameLength field in the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameData) SetFontNameLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// ATSFlatDataFontSpecRawNameDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatafontspecrawnamedataheader
type ATSFlatDataFontSpecRawNameDataHeader struct {
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

// NumberOfFlattenedNames returns the NumberOfFlattenedNames field from the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameDataHeader) NumberOfFlattenedNames() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNumberOfFlattenedNames updates the NumberOfFlattenedNames field in the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameDataHeader) SetNumberOfFlattenedNames(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NameDataArray returns the NameDataArray field from the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameDataHeader) NameDataArray() [1]ATSFlatDataFontSpecRawNameData {
	return *(*[1]ATSFlatDataFontSpecRawNameData)(unsafe.Pointer(&s.storage[4]))
}

// SetNameDataArray updates the NameDataArray field in the record's packed storage.
func (s *ATSFlatDataFontSpecRawNameDataHeader) SetNameDataArray(v [1]ATSFlatDataFontSpecRawNameData) {
	*(*[1]ATSFlatDataFontSpecRawNameData)(unsafe.Pointer(&s.storage[4])) = v
}

// ATSFlatDataLayoutControlsDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatalayoutcontrolsdataheader
type ATSFlatDataLayoutControlsDataHeader struct {
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

// NumberOfLayoutControls returns the NumberOfLayoutControls field from the record's packed storage.
func (s *ATSFlatDataLayoutControlsDataHeader) NumberOfLayoutControls() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNumberOfLayoutControls updates the NumberOfLayoutControls field in the record's packed storage.
func (s *ATSFlatDataLayoutControlsDataHeader) SetNumberOfLayoutControls(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// ControlArray returns the ControlArray field from the record's packed storage.
func (s *ATSFlatDataLayoutControlsDataHeader) ControlArray() [1]ATSUAttributeInfo {
	return *(*[1]ATSUAttributeInfo)(unsafe.Pointer(&s.storage[4]))
}

// SetControlArray updates the ControlArray field in the record's packed storage.
func (s *ATSFlatDataLayoutControlsDataHeader) SetControlArray(v [1]ATSUAttributeInfo) {
	*(*[1]ATSUAttributeInfo)(unsafe.Pointer(&s.storage[4])) = v
}

// ATSFlatDataLineInfoData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatalineinfodata
type ATSFlatDataLineInfoData struct {
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

// LineLength returns the LineLength field from the record's packed storage.
func (s *ATSFlatDataLineInfoData) LineLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetLineLength updates the LineLength field in the record's packed storage.
func (s *ATSFlatDataLineInfoData) SetLineLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NumberOfLineControls returns the NumberOfLineControls field from the record's packed storage.
func (s *ATSFlatDataLineInfoData) NumberOfLineControls() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNumberOfLineControls updates the NumberOfLineControls field in the record's packed storage.
func (s *ATSFlatDataLineInfoData) SetNumberOfLineControls(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// ATSFlatDataLineInfoHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatalineinfoheader
type ATSFlatDataLineInfoHeader struct {
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

// NumberOfLines returns the NumberOfLines field from the record's packed storage.
func (s *ATSFlatDataLineInfoHeader) NumberOfLines() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNumberOfLines updates the NumberOfLines field in the record's packed storage.
func (s *ATSFlatDataLineInfoHeader) SetNumberOfLines(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// LineInfoArray returns the LineInfoArray field from the record's packed storage.
func (s *ATSFlatDataLineInfoHeader) LineInfoArray() [1]ATSFlatDataLineInfoData {
	return *(*[1]ATSFlatDataLineInfoData)(unsafe.Pointer(&s.storage[4]))
}

// SetLineInfoArray updates the LineInfoArray field in the record's packed storage.
func (s *ATSFlatDataLineInfoHeader) SetLineInfoArray(v [1]ATSFlatDataLineInfoData) {
	*(*[1]ATSFlatDataLineInfoData)(unsafe.Pointer(&s.storage[4])) = v
}

// ATSFlatDataMainHeaderBlock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatamainheaderblock
type ATSFlatDataMainHeaderBlock struct {
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

// Version returns the Version field from the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) Version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) SetVersion(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// SizeOfDataBlock returns the SizeOfDataBlock field from the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) SizeOfDataBlock() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetSizeOfDataBlock updates the SizeOfDataBlock field in the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) SetSizeOfDataBlock(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// OffsetToTextLayouts returns the OffsetToTextLayouts field from the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) OffsetToTextLayouts() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetOffsetToTextLayouts updates the OffsetToTextLayouts field in the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) SetOffsetToTextLayouts(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// OffsetToStyleRuns returns the OffsetToStyleRuns field from the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) OffsetToStyleRuns() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetOffsetToStyleRuns updates the OffsetToStyleRuns field in the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) SetOffsetToStyleRuns(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// OffsetToStyleList returns the OffsetToStyleList field from the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) OffsetToStyleList() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetOffsetToStyleList updates the OffsetToStyleList field in the record's packed storage.
func (s *ATSFlatDataMainHeaderBlock) SetOffsetToStyleList(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// ATSFlatDataStyleListFeatureData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastylelistfeaturedata
type ATSFlatDataStyleListFeatureData struct {
	TheFeatureType     ATSUFontFeatureType
	TheFeatureSelector ATSUFontFeatureSelector
}

// ATSFlatDataStyleListHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastylelistheader
type ATSFlatDataStyleListHeader struct {
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

// NumberOfStyles returns the NumberOfStyles field from the record's packed storage.
func (s *ATSFlatDataStyleListHeader) NumberOfStyles() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNumberOfStyles updates the NumberOfStyles field in the record's packed storage.
func (s *ATSFlatDataStyleListHeader) SetNumberOfStyles(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// StyleDataArray returns the StyleDataArray field from the record's packed storage.
func (s *ATSFlatDataStyleListHeader) StyleDataArray() [1]ATSFlatDataStyleListStyleDataHeader {
	return *(*[1]ATSFlatDataStyleListStyleDataHeader)(unsafe.Pointer(&s.storage[4]))
}

// SetStyleDataArray updates the StyleDataArray field in the record's packed storage.
func (s *ATSFlatDataStyleListHeader) SetStyleDataArray(v [1]ATSFlatDataStyleListStyleDataHeader) {
	*(*[1]ATSFlatDataStyleListStyleDataHeader)(unsafe.Pointer(&s.storage[4])) = v
}

// ATSFlatDataStyleListStyleDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastyleliststyledataheader
type ATSFlatDataStyleListStyleDataHeader struct {
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

// SizeOfStyleInfo returns the SizeOfStyleInfo field from the record's packed storage.
func (s *ATSFlatDataStyleListStyleDataHeader) SizeOfStyleInfo() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSizeOfStyleInfo updates the SizeOfStyleInfo field in the record's packed storage.
func (s *ATSFlatDataStyleListStyleDataHeader) SetSizeOfStyleInfo(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// NumberOfSetAttributes returns the NumberOfSetAttributes field from the record's packed storage.
func (s *ATSFlatDataStyleListStyleDataHeader) NumberOfSetAttributes() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetNumberOfSetAttributes updates the NumberOfSetAttributes field in the record's packed storage.
func (s *ATSFlatDataStyleListStyleDataHeader) SetNumberOfSetAttributes(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// NumberOfSetFeatures returns the NumberOfSetFeatures field from the record's packed storage.
func (s *ATSFlatDataStyleListStyleDataHeader) NumberOfSetFeatures() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetNumberOfSetFeatures updates the NumberOfSetFeatures field in the record's packed storage.
func (s *ATSFlatDataStyleListStyleDataHeader) SetNumberOfSetFeatures(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// NumberOfSetVariations returns the NumberOfSetVariations field from the record's packed storage.
func (s *ATSFlatDataStyleListStyleDataHeader) NumberOfSetVariations() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetNumberOfSetVariations updates the NumberOfSetVariations field in the record's packed storage.
func (s *ATSFlatDataStyleListStyleDataHeader) SetNumberOfSetVariations(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// ATSFlatDataStyleListVariationData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastylelistvariationdata
type ATSFlatDataStyleListVariationData struct {
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

// TheVariationAxis returns the TheVariationAxis field from the record's packed storage.
func (s *ATSFlatDataStyleListVariationData) TheVariationAxis() ATSUFontVariationAxis {
	return ATSUFontVariationAxis(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetTheVariationAxis updates the TheVariationAxis field in the record's packed storage.
func (s *ATSFlatDataStyleListVariationData) SetTheVariationAxis(v ATSUFontVariationAxis) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// TheVariationValue returns the TheVariationValue field from the record's packed storage.
func (s *ATSFlatDataStyleListVariationData) TheVariationValue() ATSUFontVariationValue {
	return ATSUFontVariationValue(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetTheVariationValue updates the TheVariationValue field in the record's packed storage.
func (s *ATSFlatDataStyleListVariationData) SetTheVariationValue(v ATSUFontVariationValue) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// ATSFlatDataStyleRunDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastylerundataheader
type ATSFlatDataStyleRunDataHeader struct {
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

// NumberOfStyleRuns returns the NumberOfStyleRuns field from the record's packed storage.
func (s *ATSFlatDataStyleRunDataHeader) NumberOfStyleRuns() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNumberOfStyleRuns updates the NumberOfStyleRuns field in the record's packed storage.
func (s *ATSFlatDataStyleRunDataHeader) SetNumberOfStyleRuns(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// StyleRunArray returns the StyleRunArray field from the record's packed storage.
func (s *ATSFlatDataStyleRunDataHeader) StyleRunArray() [1]ATSUStyleRunInfo {
	return *(*[1]ATSUStyleRunInfo)(unsafe.Pointer(&s.storage[4]))
}

// SetStyleRunArray updates the StyleRunArray field in the record's packed storage.
func (s *ATSFlatDataStyleRunDataHeader) SetStyleRunArray(v [1]ATSUStyleRunInfo) {
	*(*[1]ATSUStyleRunInfo)(unsafe.Pointer(&s.storage[4])) = v
}

// ATSFlatDataTextLayoutDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatatextlayoutdataheader
type ATSFlatDataTextLayoutDataHeader struct {
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

// SizeOfLayoutData returns the SizeOfLayoutData field from the record's packed storage.
func (s *ATSFlatDataTextLayoutDataHeader) SizeOfLayoutData() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSizeOfLayoutData updates the SizeOfLayoutData field in the record's packed storage.
func (s *ATSFlatDataTextLayoutDataHeader) SetSizeOfLayoutData(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// TextLayoutLength returns the TextLayoutLength field from the record's packed storage.
func (s *ATSFlatDataTextLayoutDataHeader) TextLayoutLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetTextLayoutLength updates the TextLayoutLength field in the record's packed storage.
func (s *ATSFlatDataTextLayoutDataHeader) SetTextLayoutLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// OffsetToLayoutControls returns the OffsetToLayoutControls field from the record's packed storage.
func (s *ATSFlatDataTextLayoutDataHeader) OffsetToLayoutControls() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetOffsetToLayoutControls updates the OffsetToLayoutControls field in the record's packed storage.
func (s *ATSFlatDataTextLayoutDataHeader) SetOffsetToLayoutControls(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// OffsetToLineInfo returns the OffsetToLineInfo field from the record's packed storage.
func (s *ATSFlatDataTextLayoutDataHeader) OffsetToLineInfo() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetOffsetToLineInfo updates the OffsetToLineInfo field in the record's packed storage.
func (s *ATSFlatDataTextLayoutDataHeader) SetOffsetToLineInfo(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// ATSFlatDataTextLayoutHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatatextlayoutheader
type ATSFlatDataTextLayoutHeader struct {
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

// NumFlattenedTextLayouts returns the NumFlattenedTextLayouts field from the record's packed storage.
func (s *ATSFlatDataTextLayoutHeader) NumFlattenedTextLayouts() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetNumFlattenedTextLayouts updates the NumFlattenedTextLayouts field in the record's packed storage.
func (s *ATSFlatDataTextLayoutHeader) SetNumFlattenedTextLayouts(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// FlattenedTextLayouts returns the FlattenedTextLayouts field from the record's packed storage.
func (s *ATSFlatDataTextLayoutHeader) FlattenedTextLayouts() [1]ATSFlatDataTextLayoutDataHeader {
	return *(*[1]ATSFlatDataTextLayoutDataHeader)(unsafe.Pointer(&s.storage[4]))
}

// SetFlattenedTextLayouts updates the FlattenedTextLayouts field in the record's packed storage.
func (s *ATSFlatDataTextLayoutHeader) SetFlattenedTextLayouts(v [1]ATSFlatDataTextLayoutDataHeader) {
	*(*[1]ATSFlatDataTextLayoutDataHeader)(unsafe.Pointer(&s.storage[4])) = v
}

// ATSFontFilter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsfontfilter
type ATSFontFilter struct {
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
func (s *ATSFontFilter) Version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *ATSFontFilter) SetVersion(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// FilterSelector returns the FilterSelector field from the record's packed storage.
func (s *ATSFontFilter) FilterSelector() ATSFontFilterSelector {
	return *(*ATSFontFilterSelector)(unsafe.Pointer(&s.storage[4]))
}

// SetFilterSelector updates the FilterSelector field in the record's packed storage.
func (s *ATSFontFilter) SetFilterSelector(v ATSFontFilterSelector) {
	*(*ATSFontFilterSelector)(unsafe.Pointer(&s.storage[4])) = v
}

// Filter returns the Filter field from the record's packed storage.
func (s *ATSFontFilter) Filter() [4]uint16 {
	return *(*[4]uint16)(unsafe.Pointer(&s.storage[8]))
}

// SetFilter updates the Filter field in the record's packed storage.
func (s *ATSFontFilter) SetFilter(v [4]uint16) {
	*(*[4]uint16)(unsafe.Pointer(&s.storage[8])) = v
}

// ATSFontMetrics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsfontmetrics
type ATSFontMetrics struct {
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
	storage [116]byte
}

// Version returns the Version field from the record's packed storage.
func (s *ATSFontMetrics) Version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *ATSFontMetrics) SetVersion(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Ascent returns the Ascent field from the record's packed storage.
func (s *ATSFontMetrics) Ascent() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetAscent updates the Ascent field in the record's packed storage.
func (s *ATSFontMetrics) SetAscent(v float64) {
	binary.NativeEndian.PutUint64(s.storage[4:12], math.Float64bits(v))
}

// Descent returns the Descent field from the record's packed storage.
func (s *ATSFontMetrics) Descent() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetDescent updates the Descent field in the record's packed storage.
func (s *ATSFontMetrics) SetDescent(v float64) {
	binary.NativeEndian.PutUint64(s.storage[12:20], math.Float64bits(v))
}

// Leading returns the Leading field from the record's packed storage.
func (s *ATSFontMetrics) Leading() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[20:28]))
}

// SetLeading updates the Leading field in the record's packed storage.
func (s *ATSFontMetrics) SetLeading(v float64) {
	binary.NativeEndian.PutUint64(s.storage[20:28], math.Float64bits(v))
}

// AvgAdvanceWidth returns the AvgAdvanceWidth field from the record's packed storage.
func (s *ATSFontMetrics) AvgAdvanceWidth() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[28:36]))
}

// SetAvgAdvanceWidth updates the AvgAdvanceWidth field in the record's packed storage.
func (s *ATSFontMetrics) SetAvgAdvanceWidth(v float64) {
	binary.NativeEndian.PutUint64(s.storage[28:36], math.Float64bits(v))
}

// MaxAdvanceWidth returns the MaxAdvanceWidth field from the record's packed storage.
func (s *ATSFontMetrics) MaxAdvanceWidth() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[36:44]))
}

// SetMaxAdvanceWidth updates the MaxAdvanceWidth field in the record's packed storage.
func (s *ATSFontMetrics) SetMaxAdvanceWidth(v float64) {
	binary.NativeEndian.PutUint64(s.storage[36:44], math.Float64bits(v))
}

// MinLeftSideBearing returns the MinLeftSideBearing field from the record's packed storage.
func (s *ATSFontMetrics) MinLeftSideBearing() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[44:52]))
}

// SetMinLeftSideBearing updates the MinLeftSideBearing field in the record's packed storage.
func (s *ATSFontMetrics) SetMinLeftSideBearing(v float64) {
	binary.NativeEndian.PutUint64(s.storage[44:52], math.Float64bits(v))
}

// MinRightSideBearing returns the MinRightSideBearing field from the record's packed storage.
func (s *ATSFontMetrics) MinRightSideBearing() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[52:60]))
}

// SetMinRightSideBearing updates the MinRightSideBearing field in the record's packed storage.
func (s *ATSFontMetrics) SetMinRightSideBearing(v float64) {
	binary.NativeEndian.PutUint64(s.storage[52:60], math.Float64bits(v))
}

// StemWidth returns the StemWidth field from the record's packed storage.
func (s *ATSFontMetrics) StemWidth() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[60:68]))
}

// SetStemWidth updates the StemWidth field in the record's packed storage.
func (s *ATSFontMetrics) SetStemWidth(v float64) {
	binary.NativeEndian.PutUint64(s.storage[60:68], math.Float64bits(v))
}

// StemHeight returns the StemHeight field from the record's packed storage.
func (s *ATSFontMetrics) StemHeight() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[68:76]))
}

// SetStemHeight updates the StemHeight field in the record's packed storage.
func (s *ATSFontMetrics) SetStemHeight(v float64) {
	binary.NativeEndian.PutUint64(s.storage[68:76], math.Float64bits(v))
}

// CapHeight returns the CapHeight field from the record's packed storage.
func (s *ATSFontMetrics) CapHeight() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[76:84]))
}

// SetCapHeight updates the CapHeight field in the record's packed storage.
func (s *ATSFontMetrics) SetCapHeight(v float64) {
	binary.NativeEndian.PutUint64(s.storage[76:84], math.Float64bits(v))
}

// XHeight returns the XHeight field from the record's packed storage.
func (s *ATSFontMetrics) XHeight() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[84:92]))
}

// SetXHeight updates the XHeight field in the record's packed storage.
func (s *ATSFontMetrics) SetXHeight(v float64) {
	binary.NativeEndian.PutUint64(s.storage[84:92], math.Float64bits(v))
}

// ItalicAngle returns the ItalicAngle field from the record's packed storage.
func (s *ATSFontMetrics) ItalicAngle() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[92:100]))
}

// SetItalicAngle updates the ItalicAngle field in the record's packed storage.
func (s *ATSFontMetrics) SetItalicAngle(v float64) {
	binary.NativeEndian.PutUint64(s.storage[92:100], math.Float64bits(v))
}

// UnderlinePosition returns the UnderlinePosition field from the record's packed storage.
func (s *ATSFontMetrics) UnderlinePosition() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[100:108]))
}

// SetUnderlinePosition updates the UnderlinePosition field in the record's packed storage.
func (s *ATSFontMetrics) SetUnderlinePosition(v float64) {
	binary.NativeEndian.PutUint64(s.storage[100:108], math.Float64bits(v))
}

// UnderlineThickness returns the UnderlineThickness field from the record's packed storage.
func (s *ATSFontMetrics) UnderlineThickness() float64 {
	return math.Float64frombits(binary.NativeEndian.Uint64(s.storage[108:116]))
}

// SetUnderlineThickness updates the UnderlineThickness field in the record's packed storage.
func (s *ATSFontMetrics) SetUnderlineThickness(v float64) {
	binary.NativeEndian.PutUint64(s.storage[108:116], math.Float64bits(v))
}

// ATSFontQuerySourceContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsfontquerysourcecontext
type ATSFontQuerySourceContext struct {
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

// Version returns the Version field from the record's packed storage.
func (s *ATSFontQuerySourceContext) Version() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *ATSFontQuerySourceContext) SetVersion(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// RefCon returns the RefCon field from the record's packed storage.
func (s *ATSFontQuerySourceContext) RefCon() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetRefCon updates the RefCon field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ATSFontQuerySourceContext) SetRefCon(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// Retain returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *ATSFontQuerySourceContext) Retain() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[12:20]))
}

// SetRetain stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ATSFontQuerySourceContext) SetRetain(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[12:20], uint64(v))
}

// Release returns a raw C code address from the record's packed storage.
// The address is not callable from Go. To call C through it, use purego to
// bind a Go function value to the address.
func (s *ATSFontQuerySourceContext) Release() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[20:28]))
}

// SetRelease stores a raw C code address in the record's packed storage.
// Use purego.NewCallback to obtain the address for a Go callback. The callback
// and its address must outlive every use C can make through this record.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ATSFontQuerySourceContext) SetRelease(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[20:28], uint64(v))
}

// ATSGlyphIdealMetrics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsglyphidealmetrics
type ATSGlyphIdealMetrics struct {
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
	storage [48]byte
}

// Advance returns the Advance field from the record's packed storage.
func (s *ATSGlyphIdealMetrics) Advance() [2]uint64 {
	return *(*[2]uint64)(unsafe.Pointer(&s.storage[0]))
}

// SetAdvance updates the Advance field in the record's packed storage.
func (s *ATSGlyphIdealMetrics) SetAdvance(v [2]uint64) {
	*(*[2]uint64)(unsafe.Pointer(&s.storage[0])) = v
}

// SideBearing returns the SideBearing field from the record's packed storage.
func (s *ATSGlyphIdealMetrics) SideBearing() [2]uint64 {
	return *(*[2]uint64)(unsafe.Pointer(&s.storage[16]))
}

// SetSideBearing updates the SideBearing field in the record's packed storage.
func (s *ATSGlyphIdealMetrics) SetSideBearing(v [2]uint64) {
	*(*[2]uint64)(unsafe.Pointer(&s.storage[16])) = v
}

// OtherSideBearing returns the OtherSideBearing field from the record's packed storage.
func (s *ATSGlyphIdealMetrics) OtherSideBearing() [2]uint64 {
	return *(*[2]uint64)(unsafe.Pointer(&s.storage[32]))
}

// SetOtherSideBearing updates the OtherSideBearing field in the record's packed storage.
func (s *ATSGlyphIdealMetrics) SetOtherSideBearing(v [2]uint64) {
	*(*[2]uint64)(unsafe.Pointer(&s.storage[32])) = v
}

// ATSGlyphScreenMetrics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsglyphscreenmetrics
type ATSGlyphScreenMetrics struct {
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
	storage [72]byte
}

// DeviceAdvance returns the DeviceAdvance field from the record's packed storage.
func (s *ATSGlyphScreenMetrics) DeviceAdvance() [2]uint64 {
	return *(*[2]uint64)(unsafe.Pointer(&s.storage[0]))
}

// SetDeviceAdvance updates the DeviceAdvance field in the record's packed storage.
func (s *ATSGlyphScreenMetrics) SetDeviceAdvance(v [2]uint64) {
	*(*[2]uint64)(unsafe.Pointer(&s.storage[0])) = v
}

// TopLeft returns the TopLeft field from the record's packed storage.
func (s *ATSGlyphScreenMetrics) TopLeft() [2]uint64 {
	return *(*[2]uint64)(unsafe.Pointer(&s.storage[16]))
}

// SetTopLeft updates the TopLeft field in the record's packed storage.
func (s *ATSGlyphScreenMetrics) SetTopLeft(v [2]uint64) {
	*(*[2]uint64)(unsafe.Pointer(&s.storage[16])) = v
}

// Height returns the Height field from the record's packed storage.
func (s *ATSGlyphScreenMetrics) Height() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetHeight updates the Height field in the record's packed storage.
func (s *ATSGlyphScreenMetrics) SetHeight(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Width returns the Width field from the record's packed storage.
func (s *ATSGlyphScreenMetrics) Width() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetWidth updates the Width field in the record's packed storage.
func (s *ATSGlyphScreenMetrics) SetWidth(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// SideBearing returns the SideBearing field from the record's packed storage.
func (s *ATSGlyphScreenMetrics) SideBearing() [2]uint64 {
	return *(*[2]uint64)(unsafe.Pointer(&s.storage[40]))
}

// SetSideBearing updates the SideBearing field in the record's packed storage.
func (s *ATSGlyphScreenMetrics) SetSideBearing(v [2]uint64) {
	*(*[2]uint64)(unsafe.Pointer(&s.storage[40])) = v
}

// OtherSideBearing returns the OtherSideBearing field from the record's packed storage.
func (s *ATSGlyphScreenMetrics) OtherSideBearing() [2]uint64 {
	return *(*[2]uint64)(unsafe.Pointer(&s.storage[56]))
}

// SetOtherSideBearing updates the OtherSideBearing field in the record's packed storage.
func (s *ATSGlyphScreenMetrics) SetOtherSideBearing(v [2]uint64) {
	*(*[2]uint64)(unsafe.Pointer(&s.storage[56])) = v
}

// ATSJustWidthDeltaEntryOverride
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsjustwidthdeltaentryoverride
type ATSJustWidthDeltaEntryOverride struct {
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

// BeforeGrowLimit returns the BeforeGrowLimit field from the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) BeforeGrowLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetBeforeGrowLimit updates the BeforeGrowLimit field in the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) SetBeforeGrowLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// BeforeShrinkLimit returns the BeforeShrinkLimit field from the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) BeforeShrinkLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetBeforeShrinkLimit updates the BeforeShrinkLimit field in the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) SetBeforeShrinkLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// AfterGrowLimit returns the AfterGrowLimit field from the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) AfterGrowLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetAfterGrowLimit updates the AfterGrowLimit field in the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) SetAfterGrowLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// AfterShrinkLimit returns the AfterShrinkLimit field from the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) AfterShrinkLimit() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetAfterShrinkLimit updates the AfterShrinkLimit field in the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) SetAfterShrinkLimit(v int32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// GrowFlags returns the GrowFlags field from the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) GrowFlags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[16:18]))
}

// SetGrowFlags updates the GrowFlags field in the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) SetGrowFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[16:18], uint16(v))
}

// ShrinkFlags returns the ShrinkFlags field from the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) ShrinkFlags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[18:20]))
}

// SetShrinkFlags updates the ShrinkFlags field in the record's packed storage.
func (s *ATSJustWidthDeltaEntryOverride) SetShrinkFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[18:20], uint16(v))
}

// ATSLayoutRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atslayoutrecord
type ATSLayoutRecord struct {
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
	storage [18]byte
}

// GlyphID returns the GlyphID field from the record's packed storage.
func (s *ATSLayoutRecord) GlyphID() ATSGlyphRef {
	return ATSGlyphRef(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetGlyphID updates the GlyphID field in the record's packed storage.
func (s *ATSLayoutRecord) SetGlyphID(v ATSGlyphRef) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *ATSLayoutRecord) Flags() ATSGlyphInfoFlags {
	return ATSGlyphInfoFlags(binary.NativeEndian.Uint32(s.storage[2:6]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *ATSLayoutRecord) SetFlags(v ATSGlyphInfoFlags) {
	binary.NativeEndian.PutUint32(s.storage[2:6], uint32(v))
}

// OriginalOffset returns the OriginalOffset field from the record's packed storage.
func (s *ATSLayoutRecord) OriginalOffset() uint {
	return uint(binary.NativeEndian.Uint64(s.storage[6:14]))
}

// SetOriginalOffset updates the OriginalOffset field in the record's packed storage.
func (s *ATSLayoutRecord) SetOriginalOffset(v uint) {
	binary.NativeEndian.PutUint64(s.storage[6:14], uint64(v))
}

// RealPos returns the RealPos field from the record's packed storage.
func (s *ATSLayoutRecord) RealPos() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[14:18]))
}

// SetRealPos updates the RealPos field in the record's packed storage.
func (s *ATSLayoutRecord) SetRealPos(v int32) {
	binary.NativeEndian.PutUint32(s.storage[14:18], uint32(v))
}

// ATSTrapezoid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atstrapezoid
type ATSTrapezoid struct {
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

// UpperLeft returns the UpperLeft field from the record's packed storage.
func (s *ATSTrapezoid) UpperLeft() [4]uint16 {
	return *(*[4]uint16)(unsafe.Pointer(&s.storage[0]))
}

// SetUpperLeft updates the UpperLeft field in the record's packed storage.
func (s *ATSTrapezoid) SetUpperLeft(v [4]uint16) {
	*(*[4]uint16)(unsafe.Pointer(&s.storage[0])) = v
}

// UpperRight returns the UpperRight field from the record's packed storage.
func (s *ATSTrapezoid) UpperRight() [4]uint16 {
	return *(*[4]uint16)(unsafe.Pointer(&s.storage[8]))
}

// SetUpperRight updates the UpperRight field in the record's packed storage.
func (s *ATSTrapezoid) SetUpperRight(v [4]uint16) {
	*(*[4]uint16)(unsafe.Pointer(&s.storage[8])) = v
}

// LowerRight returns the LowerRight field from the record's packed storage.
func (s *ATSTrapezoid) LowerRight() [4]uint16 {
	return *(*[4]uint16)(unsafe.Pointer(&s.storage[16]))
}

// SetLowerRight updates the LowerRight field in the record's packed storage.
func (s *ATSTrapezoid) SetLowerRight(v [4]uint16) {
	*(*[4]uint16)(unsafe.Pointer(&s.storage[16])) = v
}

// LowerLeft returns the LowerLeft field from the record's packed storage.
func (s *ATSTrapezoid) LowerLeft() [4]uint16 {
	return *(*[4]uint16)(unsafe.Pointer(&s.storage[24]))
}

// SetLowerLeft updates the LowerLeft field in the record's packed storage.
func (s *ATSTrapezoid) SetLowerLeft(v [4]uint16) {
	*(*[4]uint16)(unsafe.Pointer(&s.storage[24])) = v
}

// ATSUAttributeInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuattributeinfo
type ATSUAttributeInfo struct {
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

// FTag returns the FTag field from the record's packed storage.
func (s *ATSUAttributeInfo) FTag() ATSUAttributeTag {
	return ATSUAttributeTag(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetFTag updates the FTag field in the record's packed storage.
func (s *ATSUAttributeInfo) SetFTag(v ATSUAttributeTag) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// FValueSize returns the FValueSize field from the record's packed storage.
func (s *ATSUAttributeInfo) FValueSize() uint {
	return uint(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetFValueSize updates the FValueSize field in the record's packed storage.
func (s *ATSUAttributeInfo) SetFValueSize(v uint) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// ATSUBackgroundData
type ATSUBackgroundData struct {
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

// BackgroundColor returns the BackgroundColor field from the record's packed storage.
func (s *ATSUBackgroundData) BackgroundColor() ATSUBackgroundColor {
	return *(*ATSUBackgroundColor)(unsafe.Pointer(&s.storage[0]))
}

// SetBackgroundColor updates the BackgroundColor field in the record's packed storage.
func (s *ATSUBackgroundData) SetBackgroundColor(v ATSUBackgroundColor) {
	*(*ATSUBackgroundColor)(unsafe.Pointer(&s.storage[0])) = v
}

// BackgroundUPP returns the BackgroundUPP field from the record's packed storage.
func (s *ATSUBackgroundData) BackgroundUPP() [16]byte {
	return *(*[16]byte)(unsafe.Pointer(&s.storage[0]))
}

// SetBackgroundUPP updates the BackgroundUPP field in the record's packed storage.
func (s *ATSUBackgroundData) SetBackgroundUPP(v [16]byte) {
	*(*[16]byte)(unsafe.Pointer(&s.storage[0])) = v
}

// ATSUCaret
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsucaret
type ATSUCaret struct {
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

// FX returns the FX field from the record's packed storage.
func (s *ATSUCaret) FX() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetFX updates the FX field in the record's packed storage.
func (s *ATSUCaret) SetFX(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// FY returns the FY field from the record's packed storage.
func (s *ATSUCaret) FY() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetFY updates the FY field in the record's packed storage.
func (s *ATSUCaret) SetFY(v int32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// FDeltaX returns the FDeltaX field from the record's packed storage.
func (s *ATSUCaret) FDeltaX() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetFDeltaX updates the FDeltaX field in the record's packed storage.
func (s *ATSUCaret) SetFDeltaX(v int32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// FDeltaY returns the FDeltaY field from the record's packed storage.
func (s *ATSUCaret) FDeltaY() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetFDeltaY updates the FDeltaY field in the record's packed storage.
func (s *ATSUCaret) SetFDeltaY(v int32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// ATSUCurvePath
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsucurvepath
type ATSUCurvePath struct {
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

// Vectors returns the Vectors field from the record's packed storage.
func (s *ATSUCurvePath) Vectors() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetVectors updates the Vectors field in the record's packed storage.
func (s *ATSUCurvePath) SetVectors(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// ControlBits returns the ControlBits field from the record's packed storage.
func (s *ATSUCurvePath) ControlBits() [1]uint32 {
	return *(*[1]uint32)(unsafe.Pointer(&s.storage[4]))
}

// SetControlBits updates the ControlBits field in the record's packed storage.
func (s *ATSUCurvePath) SetControlBits(v [1]uint32) {
	*(*[1]uint32)(unsafe.Pointer(&s.storage[4])) = v
}

// Vector returns the Vector field from the record's packed storage.
func (s *ATSUCurvePath) Vector() [16]byte {
	return *(*[16]byte)(unsafe.Pointer(&s.storage[8]))
}

// SetVector updates the Vector field in the record's packed storage.
func (s *ATSUCurvePath) SetVector(v [16]byte) {
	*(*[16]byte)(unsafe.Pointer(&s.storage[8])) = v
}

// ATSUCurvePaths
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsucurvepaths
type ATSUCurvePaths struct {
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

// Contours returns the Contours field from the record's packed storage.
func (s *ATSUCurvePaths) Contours() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetContours updates the Contours field in the record's packed storage.
func (s *ATSUCurvePaths) SetContours(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Contour returns the Contour field from the record's packed storage.
func (s *ATSUCurvePaths) Contour() [1]ATSUCurvePath {
	return *(*[1]ATSUCurvePath)(unsafe.Pointer(&s.storage[4]))
}

// SetContour updates the Contour field in the record's packed storage.
func (s *ATSUCurvePaths) SetContour(v [1]ATSUCurvePath) {
	*(*[1]ATSUCurvePath)(unsafe.Pointer(&s.storage[4])) = v
}

// ATSUGlyphInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuglyphinfo
type ATSUGlyphInfo struct {
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

// GlyphID returns the GlyphID field from the record's packed storage.
func (s *ATSUGlyphInfo) GlyphID() GlyphID {
	return GlyphID(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetGlyphID updates the GlyphID field in the record's packed storage.
func (s *ATSUGlyphInfo) SetGlyphID(v GlyphID) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *ATSUGlyphInfo) Reserved() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *ATSUGlyphInfo) SetReserved(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// LayoutFlags returns the LayoutFlags field from the record's packed storage.
func (s *ATSUGlyphInfo) LayoutFlags() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetLayoutFlags updates the LayoutFlags field in the record's packed storage.
func (s *ATSUGlyphInfo) SetLayoutFlags(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// CharIndex returns the CharIndex field from the record's packed storage.
func (s *ATSUGlyphInfo) CharIndex() uint {
	return uint(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetCharIndex updates the CharIndex field in the record's packed storage.
func (s *ATSUGlyphInfo) SetCharIndex(v uint) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Style returns the Style field from the record's packed storage.
func (s *ATSUGlyphInfo) Style() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetStyle updates the Style field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ATSUGlyphInfo) SetStyle(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// DeltaY returns the DeltaY field from the record's packed storage.
func (s *ATSUGlyphInfo) DeltaY() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetDeltaY updates the DeltaY field in the record's packed storage.
func (s *ATSUGlyphInfo) SetDeltaY(v float32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], math.Float32bits(v))
}

// IdealX returns the IdealX field from the record's packed storage.
func (s *ATSUGlyphInfo) IdealX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetIdealX updates the IdealX field in the record's packed storage.
func (s *ATSUGlyphInfo) SetIdealX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], math.Float32bits(v))
}

// ScreenX returns the ScreenX field from the record's packed storage.
func (s *ATSUGlyphInfo) ScreenX() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[32:34]))
}

// SetScreenX updates the ScreenX field in the record's packed storage.
func (s *ATSUGlyphInfo) SetScreenX(v int16) {
	binary.NativeEndian.PutUint16(s.storage[32:34], uint16(v))
}

// CaretX returns the CaretX field from the record's packed storage.
func (s *ATSUGlyphInfo) CaretX() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[34:36]))
}

// SetCaretX updates the CaretX field in the record's packed storage.
func (s *ATSUGlyphInfo) SetCaretX(v int16) {
	binary.NativeEndian.PutUint16(s.storage[34:36], uint16(v))
}

// ATSUGlyphInfoArray
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuglyphinfoarray
type ATSUGlyphInfoArray struct {
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
	storage [52]byte
}

// Layout returns the Layout field from the record's packed storage.
func (s *ATSUGlyphInfoArray) Layout() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetLayout updates the Layout field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ATSUGlyphInfoArray) SetLayout(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// NumGlyphs returns the NumGlyphs field from the record's packed storage.
func (s *ATSUGlyphInfoArray) NumGlyphs() uint {
	return uint(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetNumGlyphs updates the NumGlyphs field in the record's packed storage.
func (s *ATSUGlyphInfoArray) SetNumGlyphs(v uint) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// Glyphs returns the Glyphs field from the record's packed storage.
func (s *ATSUGlyphInfoArray) Glyphs() [1]ATSUGlyphInfo {
	return *(*[1]ATSUGlyphInfo)(unsafe.Pointer(&s.storage[16]))
}

// SetGlyphs updates the Glyphs field in the record's packed storage.
func (s *ATSUGlyphInfoArray) SetGlyphs(v [1]ATSUGlyphInfo) {
	*(*[1]ATSUGlyphInfo)(unsafe.Pointer(&s.storage[16])) = v
}

// ATSUGlyphSelector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuglyphselector
type ATSUGlyphSelector struct {
	Collection GlyphCollection
	GlyphID    GlyphID
}

// ATSURGBAlphaColor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsurgbalphacolor
type ATSURGBAlphaColor struct {
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

// Red returns the Red field from the record's packed storage.
func (s *ATSURGBAlphaColor) Red() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetRed updates the Red field in the record's packed storage.
func (s *ATSURGBAlphaColor) SetRed(v float32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], math.Float32bits(v))
}

// Green returns the Green field from the record's packed storage.
func (s *ATSURGBAlphaColor) Green() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetGreen updates the Green field in the record's packed storage.
func (s *ATSURGBAlphaColor) SetGreen(v float32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], math.Float32bits(v))
}

// Blue returns the Blue field from the record's packed storage.
func (s *ATSURGBAlphaColor) Blue() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetBlue updates the Blue field in the record's packed storage.
func (s *ATSURGBAlphaColor) SetBlue(v float32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], math.Float32bits(v))
}

// Alpha returns the Alpha field from the record's packed storage.
func (s *ATSURGBAlphaColor) Alpha() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetAlpha updates the Alpha field in the record's packed storage.
func (s *ATSURGBAlphaColor) SetAlpha(v float32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], math.Float32bits(v))
}

// ATSUStyleRunInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsustyleruninfo
type ATSUStyleRunInfo struct {
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

// RunLength returns the RunLength field from the record's packed storage.
func (s *ATSUStyleRunInfo) RunLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetRunLength updates the RunLength field in the record's packed storage.
func (s *ATSUStyleRunInfo) SetRunLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// StyleObjectIndex returns the StyleObjectIndex field from the record's packed storage.
func (s *ATSUStyleRunInfo) StyleObjectIndex() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetStyleObjectIndex updates the StyleObjectIndex field in the record's packed storage.
func (s *ATSUStyleRunInfo) SetStyleObjectIndex(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// ATSUTab
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsutab
type ATSUTab struct {
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
	storage [6]byte
}

// TabPosition returns the TabPosition field from the record's packed storage.
func (s *ATSUTab) TabPosition() ATSUTextMeasurement {
	return ATSUTextMeasurement(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetTabPosition updates the TabPosition field in the record's packed storage.
func (s *ATSUTab) SetTabPosition(v ATSUTextMeasurement) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// TabType returns the TabType field from the record's packed storage.
func (s *ATSUTab) TabType() ATSUTabType {
	return ATSUTabType(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetTabType updates the TabType field in the record's packed storage.
func (s *ATSUTab) SetTabType(v ATSUTabType) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// ATSUUnhighlightData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuunhighlightdata
type ATSUUnhighlightData struct {
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

// DataType returns the DataType field from the record's packed storage.
func (s *ATSUUnhighlightData) DataType() ATSUBackgroundDataType {
	return ATSUBackgroundDataType(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetDataType updates the DataType field in the record's packed storage.
func (s *ATSUUnhighlightData) SetDataType(v ATSUBackgroundDataType) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// UnhighlightData returns the UnhighlightData field from the record's packed storage.
func (s *ATSUUnhighlightData) UnhighlightData() ATSUBackgroundData {
	return *(*ATSUBackgroundData)(unsafe.Pointer(&s.storage[4]))
}

// SetUnhighlightData updates the UnhighlightData field in the record's packed storage.
func (s *ATSUUnhighlightData) SetUnhighlightData(v ATSUBackgroundData) {
	*(*ATSUBackgroundData)(unsafe.Pointer(&s.storage[4])) = v
}

// AppParameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/appparameters
type AppParameters struct {
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

// What returns the What field from the record's packed storage.
func (s *AppParameters) What() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetWhat updates the What field in the record's packed storage.
func (s *AppParameters) SetWhat(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Message returns the Message field from the record's packed storage.
func (s *AppParameters) Message() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[2:6]))
}

// SetMessage updates the Message field in the record's packed storage.
func (s *AppParameters) SetMessage(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[2:6], uint32(v))
}

// When returns the When field from the record's packed storage.
func (s *AppParameters) When() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[6:10]))
}

// SetWhen updates the When field in the record's packed storage.
func (s *AppParameters) SetWhen(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[6:10], uint32(v))
}

// Where returns the Where field from the record's packed storage.
func (s *AppParameters) Where() Point {
	return *(*Point)(unsafe.Pointer(&s.storage[10]))
}

// SetWhere updates the Where field in the record's packed storage.
func (s *AppParameters) SetWhere(v Point) {
	*(*Point)(unsafe.Pointer(&s.storage[10])) = v
}

// Modifiers returns the Modifiers field from the record's packed storage.
func (s *AppParameters) Modifiers() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[14:16]))
}

// SetModifiers updates the Modifiers field in the record's packed storage.
func (s *AppParameters) SetModifiers(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[14:16], uint16(v))
}

// EventRefCon returns the EventRefCon field from the record's packed storage.
func (s *AppParameters) EventRefCon() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetEventRefCon updates the EventRefCon field in the record's packed storage.
func (s *AppParameters) SetEventRefCon(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// MessageLength returns the MessageLength field from the record's packed storage.
func (s *AppParameters) MessageLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetMessageLength updates the MessageLength field in the record's packed storage.
func (s *AppParameters) SetMessageLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// TheMsgEvent returns the TheMsgEvent field from the record's packed storage.
func (s *AppParameters) TheMsgEvent() [2]byte {
	return *(*[2]byte)(unsafe.Pointer(&s.storage[0]))
}

// SetTheMsgEvent updates the TheMsgEvent field in the record's packed storage.
func (s *AppParameters) SetTheMsgEvent(v [2]byte) {
	*(*[2]byte)(unsafe.Pointer(&s.storage[0])) = v
}

// AsscEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/asscentry
type AsscEntry struct {
	FontSize  int16
	FontStyle int16
	FontID    int16
}

// BitMap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/bitmap
type BitMap struct {
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
	storage [18]byte
}

// BaseAddr returns the BaseAddr field from the record's packed storage.
func (s *BitMap) BaseAddr() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetBaseAddr updates the BaseAddr field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *BitMap) SetBaseAddr(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// RowBytes returns the RowBytes field from the record's packed storage.
func (s *BitMap) RowBytes() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetRowBytes updates the RowBytes field in the record's packed storage.
func (s *BitMap) SetRowBytes(v int16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// Bounds returns the Bounds field from the record's packed storage.
func (s *BitMap) Bounds() Rect {
	return *(*Rect)(unsafe.Pointer(&s.storage[10]))
}

// SetBounds updates the Bounds field in the record's packed storage.
func (s *BitMap) SetBounds(v Rect) {
	*(*Rect)(unsafe.Pointer(&s.storage[10])) = v
}

// CM2Profile
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cm2profile
type CM2Profile struct {
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
	storage [146]byte
}

// Header returns the Header field from the record's packed storage.
func (s *CM2Profile) Header() [128]byte {
	return *(*[128]byte)(unsafe.Pointer(&s.storage[0]))
}

// SetHeader updates the Header field in the record's packed storage.
func (s *CM2Profile) SetHeader(v [128]byte) {
	*(*[128]byte)(unsafe.Pointer(&s.storage[0])) = v
}

// TagTable returns the TagTable field from the record's packed storage.
func (s *CM2Profile) TagTable() [16]byte {
	return *(*[16]byte)(unsafe.Pointer(&s.storage[128]))
}

// SetTagTable updates the TagTable field in the record's packed storage.
func (s *CM2Profile) SetTagTable(v [16]byte) {
	*(*[16]byte)(unsafe.Pointer(&s.storage[128])) = v
}

// ElemData returns the ElemData field from the record's packed storage.
func (s *CM2Profile) ElemData() [2]byte {
	return *(*[2]byte)(unsafe.Pointer(&s.storage[144]))
}

// SetElemData updates the ElemData field in the record's packed storage.
func (s *CM2Profile) SetElemData(v [2]byte) {
	*(*[2]byte)(unsafe.Pointer(&s.storage[144])) = v
}

// CMDeviceInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmdeviceinfo
type CMDeviceInfo struct {
	DataVersion      unsafe.Pointer
	DeviceClass      unsafe.Pointer
	DeviceID         unsafe.Pointer
	DeviceScope      unsafe.Pointer
	DeviceState      unsafe.Pointer
	DefaultProfileID unsafe.Pointer
	DeviceName       unsafe.Pointer // See the CFDictionary documentation for a description of the [CFDictionaryRef] data type.
	ProfileCount     unsafe.Pointer
	Reserved         unsafe.Pointer
}

// CMDeviceProfileArray
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmdeviceprofilearray
type CMDeviceProfileArray struct {
	ProfileCount unsafe.Pointer
	Profiles     unsafe.Pointer
}

// CMDeviceScope
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmdevicescope
type CMDeviceScope struct {
	DeviceUser unsafe.Pointer
	DeviceHost unsafe.Pointer
}

// CMMultiFunctLutType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmmultifunctluttype
type CMMultiFunctLutType struct {
	TypeDescriptor unsafe.Pointer
	Reserved       unsafe.Pointer
	InputChannels  unsafe.Pointer
	OutputChannels unsafe.Pointer
	Reserved2      unsafe.Pointer
	OffsetBcurves  unsafe.Pointer
	OffsetMatrix   unsafe.Pointer
	OffsetMcurves  unsafe.Pointer
	OffsetCLUT     unsafe.Pointer
	OffsetAcurves  unsafe.Pointer
	Data           unsafe.Pointer
}

// CMXYZColor - Contains values for a color specified in XYZ color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmxyzcolor
type CMXYZColor struct {
	X CMXYZComponent
	Y CMXYZComponent
	Z CMXYZComponent
}

// CQDProcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cqdprocs
type CQDProcs struct {
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
	storage [160]byte
}

// TextProc returns the TextProc field from the record's packed storage.
func (s *CQDProcs) TextProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetTextProc updates the TextProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetTextProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// LineProc returns the LineProc field from the record's packed storage.
func (s *CQDProcs) LineProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[8:16]))
}

// SetLineProc updates the LineProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetLineProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[8:16], uint64(v))
}

// RectProc returns the RectProc field from the record's packed storage.
func (s *CQDProcs) RectProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetRectProc updates the RectProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetRectProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// RRectProc returns the RRectProc field from the record's packed storage.
func (s *CQDProcs) RRectProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[24:32]))
}

// SetRRectProc updates the RRectProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetRRectProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[24:32], uint64(v))
}

// OvalProc returns the OvalProc field from the record's packed storage.
func (s *CQDProcs) OvalProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetOvalProc updates the OvalProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetOvalProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// ArcProc returns the ArcProc field from the record's packed storage.
func (s *CQDProcs) ArcProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[40:48]))
}

// SetArcProc updates the ArcProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetArcProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[40:48], uint64(v))
}

// PolyProc returns the PolyProc field from the record's packed storage.
func (s *CQDProcs) PolyProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[48:56]))
}

// SetPolyProc updates the PolyProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetPolyProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[48:56], uint64(v))
}

// RgnProc returns the RgnProc field from the record's packed storage.
func (s *CQDProcs) RgnProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[56:64]))
}

// SetRgnProc updates the RgnProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetRgnProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[56:64], uint64(v))
}

// BitsProc returns the BitsProc field from the record's packed storage.
func (s *CQDProcs) BitsProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetBitsProc updates the BitsProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetBitsProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// CommentProc returns the CommentProc field from the record's packed storage.
func (s *CQDProcs) CommentProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[72:80]))
}

// SetCommentProc updates the CommentProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetCommentProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[72:80], uint64(v))
}

// TxMeasProc returns the TxMeasProc field from the record's packed storage.
func (s *CQDProcs) TxMeasProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[80:88]))
}

// SetTxMeasProc updates the TxMeasProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetTxMeasProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[80:88], uint64(v))
}

// GetPicProc returns the GetPicProc field from the record's packed storage.
func (s *CQDProcs) GetPicProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[88:96]))
}

// SetGetPicProc updates the GetPicProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetGetPicProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[88:96], uint64(v))
}

// PutPicProc returns the PutPicProc field from the record's packed storage.
func (s *CQDProcs) PutPicProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[96:104]))
}

// SetPutPicProc updates the PutPicProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetPutPicProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[96:104], uint64(v))
}

// OpcodeProc returns the OpcodeProc field from the record's packed storage.
func (s *CQDProcs) OpcodeProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[104:112]))
}

// SetOpcodeProc updates the OpcodeProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetOpcodeProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[104:112], uint64(v))
}

// NewProc1 returns the NewProc1 field from the record's packed storage.
func (s *CQDProcs) NewProc1() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[112:120]))
}

// SetNewProc1 updates the NewProc1 field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetNewProc1(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[112:120], uint64(v))
}

// GlyphsProc returns the GlyphsProc field from the record's packed storage.
func (s *CQDProcs) GlyphsProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[120:128]))
}

// SetGlyphsProc updates the GlyphsProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetGlyphsProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[120:128], uint64(v))
}

// PrinterStatusProc returns the PrinterStatusProc field from the record's packed storage.
func (s *CQDProcs) PrinterStatusProc() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[128:136]))
}

// SetPrinterStatusProc updates the PrinterStatusProc field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetPrinterStatusProc(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[128:136], uint64(v))
}

// NewProc4 returns the NewProc4 field from the record's packed storage.
func (s *CQDProcs) NewProc4() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[136:144]))
}

// SetNewProc4 updates the NewProc4 field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetNewProc4(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[136:144], uint64(v))
}

// NewProc5 returns the NewProc5 field from the record's packed storage.
func (s *CQDProcs) NewProc5() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[144:152]))
}

// SetNewProc5 updates the NewProc5 field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetNewProc5(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[144:152], uint64(v))
}

// NewProc6 returns the NewProc6 field from the record's packed storage.
func (s *CQDProcs) NewProc6() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[152:160]))
}

// SetNewProc6 updates the NewProc6 field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *CQDProcs) SetNewProc6(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[152:160], uint64(v))
}

// ColorSpec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/colorspec
type ColorSpec struct {
	Value int16
	Rgb   RGBColor
}

// ColorTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/colortable
type ColorTable struct {
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

// CtSeed returns the CtSeed field from the record's packed storage.
func (s *ColorTable) CtSeed() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetCtSeed updates the CtSeed field in the record's packed storage.
func (s *ColorTable) SetCtSeed(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// CtFlags returns the CtFlags field from the record's packed storage.
func (s *ColorTable) CtFlags() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetCtFlags updates the CtFlags field in the record's packed storage.
func (s *ColorTable) SetCtFlags(v int16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// CtSize returns the CtSize field from the record's packed storage.
func (s *ColorTable) CtSize() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetCtSize updates the CtSize field in the record's packed storage.
func (s *ColorTable) SetCtSize(v int16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// CtTable returns the CtTable field from the record's packed storage.
func (s *ColorTable) CtTable() [1]ColorSpec {
	return *(*[1]ColorSpec)(unsafe.Pointer(&s.storage[8]))
}

// SetCtTable updates the CtTable field in the record's packed storage.
func (s *ColorTable) SetCtTable(v [1]ColorSpec) {
	*(*[1]ColorSpec)(unsafe.Pointer(&s.storage[8])) = v
}

// DelimiterInfo - Defines a delimiter information structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/delimiterinfo
type DelimiterInfo struct {
	StartDelimiter [2]uint8 // The start delimiter for an embedded command. By default, the start delimiter is “`[[`”.
	EndDelimiter   [2]uint8 // The end delimiter for an embedded command. By default, the end delimiter is “`]]`”.

}

// FMFontFamilyInstance
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfontfamilyinstance
type FMFontFamilyInstance struct {
	FontFamily FMFontFamily
	FontStyle  FMFontStyle
}

// FMFontFamilyInstanceIterator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfontfamilyinstanceiterator
type FMFontFamilyInstanceIterator struct {
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
	storage [64]byte
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *FMFontFamilyInstanceIterator) Reserved() [16]uint32 {
	return *(*[16]uint32)(unsafe.Pointer(&s.storage[0]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *FMFontFamilyInstanceIterator) SetReserved(v [16]uint32) {
	*(*[16]uint32)(unsafe.Pointer(&s.storage[0])) = v
}

// FMFontFamilyIterator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfontfamilyiterator
type FMFontFamilyIterator struct {
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
	storage [64]byte
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *FMFontFamilyIterator) Reserved() [16]uint32 {
	return *(*[16]uint32)(unsafe.Pointer(&s.storage[0]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *FMFontFamilyIterator) SetReserved(v [16]uint32) {
	*(*[16]uint32)(unsafe.Pointer(&s.storage[0])) = v
}

// FMFontIterator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfontiterator
type FMFontIterator struct {
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
	storage [64]byte
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *FMFontIterator) Reserved() [16]uint32 {
	return *(*[16]uint32)(unsafe.Pointer(&s.storage[0]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *FMFontIterator) SetReserved(v [16]uint32) {
	*(*[16]uint32)(unsafe.Pointer(&s.storage[0])) = v
}

// FMInput
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fminput
type FMInput struct {
	Family   int16
	Size     int16
	Face     uint8
	NeedBits bool
	Device   int16
	Numer    Point
	Denom    Point
}

// FamRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/famrec
type FamRec struct {
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
	storage [52]byte
}

// FfFlags returns the FfFlags field from the record's packed storage.
func (s *FamRec) FfFlags() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetFfFlags updates the FfFlags field in the record's packed storage.
func (s *FamRec) SetFfFlags(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// FfFamID returns the FfFamID field from the record's packed storage.
func (s *FamRec) FfFamID() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetFfFamID updates the FfFamID field in the record's packed storage.
func (s *FamRec) SetFfFamID(v int16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// FfFirstChar returns the FfFirstChar field from the record's packed storage.
func (s *FamRec) FfFirstChar() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetFfFirstChar updates the FfFirstChar field in the record's packed storage.
func (s *FamRec) SetFfFirstChar(v int16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// FfLastChar returns the FfLastChar field from the record's packed storage.
func (s *FamRec) FfLastChar() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetFfLastChar updates the FfLastChar field in the record's packed storage.
func (s *FamRec) SetFfLastChar(v int16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// FfAscent returns the FfAscent field from the record's packed storage.
func (s *FamRec) FfAscent() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetFfAscent updates the FfAscent field in the record's packed storage.
func (s *FamRec) SetFfAscent(v int16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// FfDescent returns the FfDescent field from the record's packed storage.
func (s *FamRec) FfDescent() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetFfDescent updates the FfDescent field in the record's packed storage.
func (s *FamRec) SetFfDescent(v int16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// FfLeading returns the FfLeading field from the record's packed storage.
func (s *FamRec) FfLeading() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[12:14]))
}

// SetFfLeading updates the FfLeading field in the record's packed storage.
func (s *FamRec) SetFfLeading(v int16) {
	binary.NativeEndian.PutUint16(s.storage[12:14], uint16(v))
}

// FfWidMax returns the FfWidMax field from the record's packed storage.
func (s *FamRec) FfWidMax() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[14:16]))
}

// SetFfWidMax updates the FfWidMax field in the record's packed storage.
func (s *FamRec) SetFfWidMax(v int16) {
	binary.NativeEndian.PutUint16(s.storage[14:16], uint16(v))
}

// FfWTabOff returns the FfWTabOff field from the record's packed storage.
func (s *FamRec) FfWTabOff() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetFfWTabOff updates the FfWTabOff field in the record's packed storage.
func (s *FamRec) SetFfWTabOff(v int32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// FfKernOff returns the FfKernOff field from the record's packed storage.
func (s *FamRec) FfKernOff() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetFfKernOff updates the FfKernOff field in the record's packed storage.
func (s *FamRec) SetFfKernOff(v int32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// FfStylOff returns the FfStylOff field from the record's packed storage.
func (s *FamRec) FfStylOff() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetFfStylOff updates the FfStylOff field in the record's packed storage.
func (s *FamRec) SetFfStylOff(v int32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// FfProperty returns the FfProperty field from the record's packed storage.
func (s *FamRec) FfProperty() [9]int16 {
	return *(*[9]int16)(unsafe.Pointer(&s.storage[28]))
}

// SetFfProperty updates the FfProperty field in the record's packed storage.
func (s *FamRec) SetFfProperty(v [9]int16) {
	*(*[9]int16)(unsafe.Pointer(&s.storage[28])) = v
}

// FfIntl returns the FfIntl field from the record's packed storage.
func (s *FamRec) FfIntl() [2]int16 {
	return *(*[2]int16)(unsafe.Pointer(&s.storage[46]))
}

// SetFfIntl updates the FfIntl field in the record's packed storage.
func (s *FamRec) SetFfIntl(v [2]int16) {
	*(*[2]int16)(unsafe.Pointer(&s.storage[46])) = v
}

// FfVersion returns the FfVersion field from the record's packed storage.
func (s *FamRec) FfVersion() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[50:52]))
}

// SetFfVersion updates the FfVersion field in the record's packed storage.
func (s *FamRec) SetFfVersion(v int16) {
	binary.NativeEndian.PutUint16(s.storage[50:52], uint16(v))
}

// FontAssoc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fontassoc
type FontAssoc struct {
	NumAssoc int16
}

// FontInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fontinfo
type FontInfo struct {
	Ascent  int16
	Descent int16
	WidMax  int16
	Leading int16
}

// FontRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fontrec
type FontRec struct {
	FontType    int16
	FirstChar   int16
	LastChar    int16
	WidMax      int16
	KernMax     int16
	NDescent    int16
	FRectWidth  int16
	FRectHeight int16
	OwTLoc      uint16
	Ascent      int16
	Descent     int16
	Leading     int16
	RowWords    int16
}

// GDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/gdevice
type GDevice struct {
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
	storage [94]byte
}

// GdRefNum returns the GdRefNum field from the record's packed storage.
func (s *GDevice) GdRefNum() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetGdRefNum updates the GdRefNum field in the record's packed storage.
func (s *GDevice) SetGdRefNum(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// GdID returns the GdID field from the record's packed storage.
func (s *GDevice) GdID() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetGdID updates the GdID field in the record's packed storage.
func (s *GDevice) SetGdID(v int16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// GdType returns the GdType field from the record's packed storage.
func (s *GDevice) GdType() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetGdType updates the GdType field in the record's packed storage.
func (s *GDevice) SetGdType(v int16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// GdITable returns the GdITable field from the record's packed storage.
func (s *GDevice) GdITable() kernel.Pointer {
	return *(*kernel.Pointer)(unsafe.Pointer(&s.storage[6]))
}

// SetGdITable updates the GdITable field in the record's packed storage.
func (s *GDevice) SetGdITable(v kernel.Pointer) {
	*(*kernel.Pointer)(unsafe.Pointer(&s.storage[6])) = v
}

// GdResPref returns the GdResPref field from the record's packed storage.
func (s *GDevice) GdResPref() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[14:16]))
}

// SetGdResPref updates the GdResPref field in the record's packed storage.
func (s *GDevice) SetGdResPref(v int16) {
	binary.NativeEndian.PutUint16(s.storage[14:16], uint16(v))
}

// GdSearchProc returns the GdSearchProc field from the record's packed storage.
func (s *GDevice) GdSearchProc() kernel.Pointer {
	return *(*kernel.Pointer)(unsafe.Pointer(&s.storage[16]))
}

// SetGdSearchProc updates the GdSearchProc field in the record's packed storage.
func (s *GDevice) SetGdSearchProc(v kernel.Pointer) {
	*(*kernel.Pointer)(unsafe.Pointer(&s.storage[16])) = v
}

// GdCompProc returns the GdCompProc field from the record's packed storage.
func (s *GDevice) GdCompProc() kernel.Pointer {
	return *(*kernel.Pointer)(unsafe.Pointer(&s.storage[24]))
}

// SetGdCompProc updates the GdCompProc field in the record's packed storage.
func (s *GDevice) SetGdCompProc(v kernel.Pointer) {
	*(*kernel.Pointer)(unsafe.Pointer(&s.storage[24])) = v
}

// GdFlags returns the GdFlags field from the record's packed storage.
func (s *GDevice) GdFlags() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[32:34]))
}

// SetGdFlags updates the GdFlags field in the record's packed storage.
func (s *GDevice) SetGdFlags(v int16) {
	binary.NativeEndian.PutUint16(s.storage[32:34], uint16(v))
}

// GdPMap returns the GdPMap field from the record's packed storage.
func (s *GDevice) GdPMap() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[34:42]))
}

// SetGdPMap updates the GdPMap field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *GDevice) SetGdPMap(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[34:42], uint64(v))
}

// GdRefCon returns the GdRefCon field from the record's packed storage.
func (s *GDevice) GdRefCon() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[42:46]))
}

// SetGdRefCon updates the GdRefCon field in the record's packed storage.
func (s *GDevice) SetGdRefCon(v int32) {
	binary.NativeEndian.PutUint32(s.storage[42:46], uint32(v))
}

// GdNextGD returns the GdNextGD field from the record's packed storage.
func (s *GDevice) GdNextGD() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[46:54]))
}

// SetGdNextGD updates the GdNextGD field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *GDevice) SetGdNextGD(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[46:54], uint64(v))
}

// GdRect returns the GdRect field from the record's packed storage.
func (s *GDevice) GdRect() Rect {
	return *(*Rect)(unsafe.Pointer(&s.storage[54]))
}

// SetGdRect updates the GdRect field in the record's packed storage.
func (s *GDevice) SetGdRect(v Rect) {
	*(*Rect)(unsafe.Pointer(&s.storage[54])) = v
}

// GdMode returns the GdMode field from the record's packed storage.
func (s *GDevice) GdMode() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[62:66]))
}

// SetGdMode updates the GdMode field in the record's packed storage.
func (s *GDevice) SetGdMode(v int32) {
	binary.NativeEndian.PutUint32(s.storage[62:66], uint32(v))
}

// GdCCBytes returns the GdCCBytes field from the record's packed storage.
func (s *GDevice) GdCCBytes() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[66:68]))
}

// SetGdCCBytes updates the GdCCBytes field in the record's packed storage.
func (s *GDevice) SetGdCCBytes(v int16) {
	binary.NativeEndian.PutUint16(s.storage[66:68], uint16(v))
}

// GdCCDepth returns the GdCCDepth field from the record's packed storage.
func (s *GDevice) GdCCDepth() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[68:70]))
}

// SetGdCCDepth updates the GdCCDepth field in the record's packed storage.
func (s *GDevice) SetGdCCDepth(v int16) {
	binary.NativeEndian.PutUint16(s.storage[68:70], uint16(v))
}

// GdCCXData returns the GdCCXData field from the record's packed storage.
func (s *GDevice) GdCCXData() kernel.Pointer {
	return *(*kernel.Pointer)(unsafe.Pointer(&s.storage[70]))
}

// SetGdCCXData updates the GdCCXData field in the record's packed storage.
func (s *GDevice) SetGdCCXData(v kernel.Pointer) {
	*(*kernel.Pointer)(unsafe.Pointer(&s.storage[70])) = v
}

// GdCCXMask returns the GdCCXMask field from the record's packed storage.
func (s *GDevice) GdCCXMask() kernel.Pointer {
	return *(*kernel.Pointer)(unsafe.Pointer(&s.storage[78]))
}

// SetGdCCXMask updates the GdCCXMask field in the record's packed storage.
func (s *GDevice) SetGdCCXMask(v kernel.Pointer) {
	*(*kernel.Pointer)(unsafe.Pointer(&s.storage[78])) = v
}

// GdExt returns the GdExt field from the record's packed storage.
func (s *GDevice) GdExt() kernel.Pointer {
	return *(*kernel.Pointer)(unsafe.Pointer(&s.storage[86]))
}

// SetGdExt updates the GdExt field in the record's packed storage.
func (s *GDevice) SetGdExt(v kernel.Pointer) {
	*(*kernel.Pointer)(unsafe.Pointer(&s.storage[86])) = v
}

// GrafPort
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/grafport
type GrafPort struct {
	Whatever [87]int16
}

// ICAppSpec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icappspec
type ICAppSpec struct {
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
	storage [68]byte
}

// FCreator returns the FCreator field from the record's packed storage.
func (s *ICAppSpec) FCreator() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetFCreator updates the FCreator field in the record's packed storage.
func (s *ICAppSpec) SetFCreator(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Name returns the Name field from the record's packed storage.
func (s *ICAppSpec) Name() [64]uint8 {
	return *(*[64]uint8)(unsafe.Pointer(&s.storage[4]))
}

// SetName updates the Name field in the record's packed storage.
func (s *ICAppSpec) SetName(v [64]uint8) {
	*(*[64]uint8)(unsafe.Pointer(&s.storage[4])) = v
}

// ICAppSpecList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icappspeclist
type ICAppSpecList struct {
	NumberOfItems int16
	AppSpecs      [1]ICAppSpec
}

// ICCharTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icchartable
type ICCharTable struct {
	NetToMac [256]uint8
	MacToNet [256]uint8
}

// ICFileSpec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icfilespec
type ICFileSpec struct {
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
	storage [112]byte
}

// VolName returns the VolName field from the record's packed storage.
func (s *ICFileSpec) VolName() [32]uint8 {
	return *(*[32]uint8)(unsafe.Pointer(&s.storage[0]))
}

// SetVolName updates the VolName field in the record's packed storage.
func (s *ICFileSpec) SetVolName(v [32]uint8) {
	*(*[32]uint8)(unsafe.Pointer(&s.storage[0])) = v
}

// VolCreationDate returns the VolCreationDate field from the record's packed storage.
func (s *ICFileSpec) VolCreationDate() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetVolCreationDate updates the VolCreationDate field in the record's packed storage.
func (s *ICFileSpec) SetVolCreationDate(v int32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// Fss returns the Fss field from the record's packed storage.
func (s *ICFileSpec) Fss() [70]byte {
	return *(*[70]byte)(unsafe.Pointer(&s.storage[36]))
}

// SetFss updates the Fss field in the record's packed storage.
func (s *ICFileSpec) SetFss(v [70]byte) {
	*(*[70]byte)(unsafe.Pointer(&s.storage[36])) = v
}

// Alias returns the Alias field from the record's packed storage.
func (s *ICFileSpec) Alias() [6]byte {
	return *(*[6]byte)(unsafe.Pointer(&s.storage[106]))
}

// SetAlias updates the Alias field in the record's packed storage.
func (s *ICFileSpec) SetAlias(v [6]byte) {
	*(*[6]byte)(unsafe.Pointer(&s.storage[106])) = v
}

// ICFontRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icfontrecord
type ICFontRecord struct {
	Size int16
	Face uint8
	Pad  int8
	Font [256]uint8
}

// ICMapEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icmapentry
type ICMapEntry struct {
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
	storage [1302]byte
}

// TotalLength returns the TotalLength field from the record's packed storage.
func (s *ICMapEntry) TotalLength() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetTotalLength updates the TotalLength field in the record's packed storage.
func (s *ICMapEntry) SetTotalLength(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// FixedLength returns the FixedLength field from the record's packed storage.
func (s *ICMapEntry) FixedLength() ICFixedLength {
	return ICFixedLength(binary.NativeEndian.Uint32(s.storage[2:4]))
}

// SetFixedLength updates the FixedLength field in the record's packed storage.
func (s *ICMapEntry) SetFixedLength(v ICFixedLength) {
	binary.NativeEndian.PutUint32(s.storage[2:4], uint32(v))
}

// Version returns the Version field from the record's packed storage.
func (s *ICMapEntry) Version() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *ICMapEntry) SetVersion(v int16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// FileType returns the FileType field from the record's packed storage.
func (s *ICMapEntry) FileType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[6:10]))
}

// SetFileType updates the FileType field in the record's packed storage.
func (s *ICMapEntry) SetFileType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[6:10], uint32(v))
}

// FileCreator returns the FileCreator field from the record's packed storage.
func (s *ICMapEntry) FileCreator() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[10:14]))
}

// SetFileCreator updates the FileCreator field in the record's packed storage.
func (s *ICMapEntry) SetFileCreator(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[10:14], uint32(v))
}

// PostCreator returns the PostCreator field from the record's packed storage.
func (s *ICMapEntry) PostCreator() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[14:18]))
}

// SetPostCreator updates the PostCreator field in the record's packed storage.
func (s *ICMapEntry) SetPostCreator(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[14:18], uint32(v))
}

// Flags returns the Flags field from the record's packed storage.
func (s *ICMapEntry) Flags() ICMapEntryFlags {
	return ICMapEntryFlags(binary.NativeEndian.Uint32(s.storage[18:22]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *ICMapEntry) SetFlags(v ICMapEntryFlags) {
	binary.NativeEndian.PutUint32(s.storage[18:22], uint32(v))
}

// Extension returns the Extension field from the record's packed storage.
func (s *ICMapEntry) Extension() [256]uint8 {
	return *(*[256]uint8)(unsafe.Pointer(&s.storage[22]))
}

// SetExtension updates the Extension field in the record's packed storage.
func (s *ICMapEntry) SetExtension(v [256]uint8) {
	*(*[256]uint8)(unsafe.Pointer(&s.storage[22])) = v
}

// CreatorAppName returns the CreatorAppName field from the record's packed storage.
func (s *ICMapEntry) CreatorAppName() [256]uint8 {
	return *(*[256]uint8)(unsafe.Pointer(&s.storage[278]))
}

// SetCreatorAppName updates the CreatorAppName field in the record's packed storage.
func (s *ICMapEntry) SetCreatorAppName(v [256]uint8) {
	*(*[256]uint8)(unsafe.Pointer(&s.storage[278])) = v
}

// PostAppName returns the PostAppName field from the record's packed storage.
func (s *ICMapEntry) PostAppName() [256]uint8 {
	return *(*[256]uint8)(unsafe.Pointer(&s.storage[534]))
}

// SetPostAppName updates the PostAppName field in the record's packed storage.
func (s *ICMapEntry) SetPostAppName(v [256]uint8) {
	*(*[256]uint8)(unsafe.Pointer(&s.storage[534])) = v
}

// MIMEType returns the MIMEType field from the record's packed storage.
func (s *ICMapEntry) MIMEType() [256]uint8 {
	return *(*[256]uint8)(unsafe.Pointer(&s.storage[790]))
}

// SetMIMEType updates the MIMEType field in the record's packed storage.
func (s *ICMapEntry) SetMIMEType(v [256]uint8) {
	*(*[256]uint8)(unsafe.Pointer(&s.storage[790])) = v
}

// EntryName returns the EntryName field from the record's packed storage.
func (s *ICMapEntry) EntryName() [256]uint8 {
	return *(*[256]uint8)(unsafe.Pointer(&s.storage[1046]))
}

// SetEntryName updates the EntryName field in the record's packed storage.
func (s *ICMapEntry) SetEntryName(v [256]uint8) {
	*(*[256]uint8)(unsafe.Pointer(&s.storage[1046])) = v
}

// ICServiceEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icserviceentry
type ICServiceEntry struct {
	Name  [256]uint8
	Port  int16
	Flags ICServiceEntryFlags
}

// ICServices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icservices
type ICServices struct {
	Count    int16
	Services [1]ICServiceEntry
}

// KernEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/kernentry
type KernEntry struct {
	KernStyle  int16
	KernLength int16
}

// KernPair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/kernpair
type KernPair struct {
	KernFirst  int8
	KernSecond int8
	KernWidth  int16
}

// KernTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/kerntable
type KernTable struct {
	NumKerns int16
}

// LaunchParamBlockRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/launchparamblockrec
type LaunchParamBlockRec struct {
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
	storage [52]byte
}

// Reserved1 returns the Reserved1 field from the record's packed storage.
func (s *LaunchParamBlockRec) Reserved1() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetReserved1 updates the Reserved1 field in the record's packed storage.
func (s *LaunchParamBlockRec) SetReserved1(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Reserved2 returns the Reserved2 field from the record's packed storage.
func (s *LaunchParamBlockRec) Reserved2() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetReserved2 updates the Reserved2 field in the record's packed storage.
func (s *LaunchParamBlockRec) SetReserved2(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// LaunchBlockID returns the LaunchBlockID field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchBlockID() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetLaunchBlockID updates the LaunchBlockID field in the record's packed storage.
func (s *LaunchParamBlockRec) SetLaunchBlockID(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// LaunchEPBLength returns the LaunchEPBLength field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchEPBLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetLaunchEPBLength updates the LaunchEPBLength field in the record's packed storage.
func (s *LaunchParamBlockRec) SetLaunchEPBLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// LaunchFileFlags returns the LaunchFileFlags field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchFileFlags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[12:14]))
}

// SetLaunchFileFlags updates the LaunchFileFlags field in the record's packed storage.
func (s *LaunchParamBlockRec) SetLaunchFileFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[12:14], uint16(v))
}

// LaunchControlFlags returns the LaunchControlFlags field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchControlFlags() LaunchFlags {
	return LaunchFlags(binary.NativeEndian.Uint16(s.storage[14:16]))
}

// SetLaunchControlFlags updates the LaunchControlFlags field in the record's packed storage.
func (s *LaunchParamBlockRec) SetLaunchControlFlags(v LaunchFlags) {
	binary.NativeEndian.PutUint16(s.storage[14:16], uint16(v))
}

// LaunchAppRef returns the LaunchAppRef field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchAppRef() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[16:24]))
}

// SetLaunchAppRef updates the LaunchAppRef field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LaunchParamBlockRec) SetLaunchAppRef(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[16:24], uint64(v))
}

// LaunchProcessSN returns the LaunchProcessSN field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchProcessSN() *ProcessSerialNumber {
	return *(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[24]))
}

// SetLaunchProcessSN updates the LaunchProcessSN field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LaunchParamBlockRec) SetLaunchProcessSN(v *ProcessSerialNumber) {
	*(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[24])) = v
}

// LaunchPreferredSize returns the LaunchPreferredSize field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchPreferredSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetLaunchPreferredSize updates the LaunchPreferredSize field in the record's packed storage.
func (s *LaunchParamBlockRec) SetLaunchPreferredSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], uint32(v))
}

// LaunchMinimumSize returns the LaunchMinimumSize field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchMinimumSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetLaunchMinimumSize updates the LaunchMinimumSize field in the record's packed storage.
func (s *LaunchParamBlockRec) SetLaunchMinimumSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[36:40], uint32(v))
}

// LaunchAvailableSize returns the LaunchAvailableSize field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchAvailableSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetLaunchAvailableSize updates the LaunchAvailableSize field in the record's packed storage.
func (s *LaunchParamBlockRec) SetLaunchAvailableSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// LaunchAppParameters returns the LaunchAppParameters field from the record's packed storage.
func (s *LaunchParamBlockRec) LaunchAppParameters() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[44:52]))
}

// SetLaunchAppParameters updates the LaunchAppParameters field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *LaunchParamBlockRec) SetLaunchAppParameters(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[44:52], uint64(v))
}

// MacPolygon
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/macpolygon
type MacPolygon struct {
	PolySize   int16
	PolyBBox   Rect
	PolyPoints [1]Point
}

// NameTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/nametable
type NameTable struct {
	StringCount  int16
	BaseFontName [256]uint8
}

// OpenCPicParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/opencpicparams
type OpenCPicParams struct {
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

// SrcRect returns the SrcRect field from the record's packed storage.
func (s *OpenCPicParams) SrcRect() Rect {
	return *(*Rect)(unsafe.Pointer(&s.storage[0]))
}

// SetSrcRect updates the SrcRect field in the record's packed storage.
func (s *OpenCPicParams) SetSrcRect(v Rect) {
	*(*Rect)(unsafe.Pointer(&s.storage[0])) = v
}

// HRes returns the HRes field from the record's packed storage.
func (s *OpenCPicParams) HRes() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetHRes updates the HRes field in the record's packed storage.
func (s *OpenCPicParams) SetHRes(v int32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// VRes returns the VRes field from the record's packed storage.
func (s *OpenCPicParams) VRes() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetVRes updates the VRes field in the record's packed storage.
func (s *OpenCPicParams) SetVRes(v int32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Version returns the Version field from the record's packed storage.
func (s *OpenCPicParams) Version() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[16:18]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *OpenCPicParams) SetVersion(v int16) {
	binary.NativeEndian.PutUint16(s.storage[16:18], uint16(v))
}

// Reserved1 returns the Reserved1 field from the record's packed storage.
func (s *OpenCPicParams) Reserved1() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[18:20]))
}

// SetReserved1 updates the Reserved1 field in the record's packed storage.
func (s *OpenCPicParams) SetReserved1(v int16) {
	binary.NativeEndian.PutUint16(s.storage[18:20], uint16(v))
}

// Reserved2 returns the Reserved2 field from the record's packed storage.
func (s *OpenCPicParams) Reserved2() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetReserved2 updates the Reserved2 field in the record's packed storage.
func (s *OpenCPicParams) SetReserved2(v int32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// PMLanguageInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/pmlanguageinfo
type PMLanguageInfo struct {
	Level   [33]uint8 // Specifies the level of the imaging language used by the printer driver.
	Version [33]uint8 // Specifies the version of the imaging language.
	Release [33]uint8 // Specifies the release of the imaging language.

}

// PMRect
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/pmrect
type PMRect struct {
	Top    float64 // The vertical coordinate for the upper-left point of the rectangle.
	Left   float64 // The horizontal coordinate for the upper-left point of the rectangle.
	Bottom float64 // The vertical coordinate for the lower-right point of the rectangle.
	Right  float64 // The horizontal coordinate for the lower-right point of the rectangle.

}

// PMResolution
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/pmresolution
type PMResolution struct {
	HRes float64 // The horizontal resolution in dots per inch (dpi).
	VRes float64 // The vertical resolution in dots per inch (dpi).

}

// Pattern
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/pattern
type Pattern struct {
	Pat [8]uint8
}

// PhonemeDescriptor - Defines a phoneme descriptor structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/phonemedescriptor
type PhonemeDescriptor struct {
	PhonemeCount int16          // The number of phonemes that the current synthesizer defines. Typically, this will correspond to the number of phonemes in the language supported by the synthesizer.
	ThePhonemes  [1]PhonemeInfo // An array of phoneme information structures.

}

// PhonemeInfo - Defines a structure that stores information about a phoneme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/phonemeinfo
type PhonemeInfo struct {
	Opcode      int16     // The opcode for the phoneme.
	PhStr       [16]uint8 // The string used to represent the phoneme. The string does not necessarily have a phonetic connection to the phoneme, but might simply be an abstract textual representation of it.
	ExampleStr  [32]uint8 // An example word that illustrates use of the phoneme.
	HiliteStart int16     // The number of characters in the example word that precede the portion of that word representing the phoneme.
	HiliteEnd   int16     // The number of characters between the beginning of the example word and the end of the portion of that word representing the phoneme.

}

// Picture
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/picture
type Picture struct {
	PicSize  int16
	PicFrame Rect
}

// PixMap
//
// PixMap records are normally obtained through a PixMapHandle rather than built with a composite literal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/pixmap
type PixMap struct {
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
	storage [62]byte
}

// BaseAddr returns the BaseAddr field from the record's packed storage.
func (s *PixMap) BaseAddr() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetBaseAddr updates the BaseAddr field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *PixMap) SetBaseAddr(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// RowBytes returns the RowBytes field from the record's packed storage.
func (s *PixMap) RowBytes() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[8:10]))
}

// SetRowBytes updates the RowBytes field in the record's packed storage.
func (s *PixMap) SetRowBytes(v int16) {
	binary.NativeEndian.PutUint16(s.storage[8:10], uint16(v))
}

// Bounds returns the Bounds field from the record's packed storage.
func (s *PixMap) Bounds() Rect {
	return *(*Rect)(unsafe.Pointer(&s.storage[10]))
}

// SetBounds updates the Bounds field in the record's packed storage.
func (s *PixMap) SetBounds(v Rect) {
	*(*Rect)(unsafe.Pointer(&s.storage[10])) = v
}

// PmVersion returns the PmVersion field from the record's packed storage.
func (s *PixMap) PmVersion() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[18:20]))
}

// SetPmVersion updates the PmVersion field in the record's packed storage.
func (s *PixMap) SetPmVersion(v int16) {
	binary.NativeEndian.PutUint16(s.storage[18:20], uint16(v))
}

// PackType returns the PackType field from the record's packed storage.
func (s *PixMap) PackType() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[20:22]))
}

// SetPackType updates the PackType field in the record's packed storage.
func (s *PixMap) SetPackType(v int16) {
	binary.NativeEndian.PutUint16(s.storage[20:22], uint16(v))
}

// PackSize returns the PackSize field from the record's packed storage.
func (s *PixMap) PackSize() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[22:26]))
}

// SetPackSize updates the PackSize field in the record's packed storage.
func (s *PixMap) SetPackSize(v int32) {
	binary.NativeEndian.PutUint32(s.storage[22:26], uint32(v))
}

// HRes returns the HRes field from the record's packed storage.
func (s *PixMap) HRes() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[26:30]))
}

// SetHRes updates the HRes field in the record's packed storage.
func (s *PixMap) SetHRes(v int32) {
	binary.NativeEndian.PutUint32(s.storage[26:30], uint32(v))
}

// VRes returns the VRes field from the record's packed storage.
func (s *PixMap) VRes() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[30:34]))
}

// SetVRes updates the VRes field in the record's packed storage.
func (s *PixMap) SetVRes(v int32) {
	binary.NativeEndian.PutUint32(s.storage[30:34], uint32(v))
}

// PixelType returns the PixelType field from the record's packed storage.
func (s *PixMap) PixelType() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[34:36]))
}

// SetPixelType updates the PixelType field in the record's packed storage.
func (s *PixMap) SetPixelType(v int16) {
	binary.NativeEndian.PutUint16(s.storage[34:36], uint16(v))
}

// PixelSize returns the PixelSize field from the record's packed storage.
func (s *PixMap) PixelSize() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[36:38]))
}

// SetPixelSize updates the PixelSize field in the record's packed storage.
func (s *PixMap) SetPixelSize(v int16) {
	binary.NativeEndian.PutUint16(s.storage[36:38], uint16(v))
}

// CmpCount returns the CmpCount field from the record's packed storage.
func (s *PixMap) CmpCount() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[38:40]))
}

// SetCmpCount updates the CmpCount field in the record's packed storage.
func (s *PixMap) SetCmpCount(v int16) {
	binary.NativeEndian.PutUint16(s.storage[38:40], uint16(v))
}

// CmpSize returns the CmpSize field from the record's packed storage.
func (s *PixMap) CmpSize() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[40:42]))
}

// SetCmpSize updates the CmpSize field in the record's packed storage.
func (s *PixMap) SetCmpSize(v int16) {
	binary.NativeEndian.PutUint16(s.storage[40:42], uint16(v))
}

// PixelFormat returns the PixelFormat field from the record's packed storage.
func (s *PixMap) PixelFormat() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[42:46]))
}

// SetPixelFormat updates the PixelFormat field in the record's packed storage.
func (s *PixMap) SetPixelFormat(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[42:46], uint32(v))
}

// PmTable returns the PmTable field from the record's packed storage.
func (s *PixMap) PmTable() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[46:54]))
}

// SetPmTable updates the PmTable field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *PixMap) SetPmTable(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[46:54], uint64(v))
}

// PmExt returns the PmExt field from the record's packed storage.
func (s *PixMap) PmExt() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[54:62]))
}

// SetPmExt updates the PmExt field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *PixMap) SetPmExt(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[54:62], uint64(v))
}

// PixPat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/pixpat
type PixPat struct {
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

// PatType returns the PatType field from the record's packed storage.
func (s *PixPat) PatType() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetPatType updates the PatType field in the record's packed storage.
func (s *PixPat) SetPatType(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// PatMap returns the PatMap field from the record's packed storage.
func (s *PixPat) PatMap() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[2:10]))
}

// SetPatMap updates the PatMap field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *PixPat) SetPatMap(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[2:10], uint64(v))
}

// PatData returns the PatData field from the record's packed storage.
func (s *PixPat) PatData() kernel.Pointer {
	return *(*kernel.Pointer)(unsafe.Pointer(&s.storage[10]))
}

// SetPatData updates the PatData field in the record's packed storage.
func (s *PixPat) SetPatData(v kernel.Pointer) {
	*(*kernel.Pointer)(unsafe.Pointer(&s.storage[10])) = v
}

// PatXData returns the PatXData field from the record's packed storage.
func (s *PixPat) PatXData() kernel.Pointer {
	return *(*kernel.Pointer)(unsafe.Pointer(&s.storage[18]))
}

// SetPatXData updates the PatXData field in the record's packed storage.
func (s *PixPat) SetPatXData(v kernel.Pointer) {
	*(*kernel.Pointer)(unsafe.Pointer(&s.storage[18])) = v
}

// PatXValid returns the PatXValid field from the record's packed storage.
func (s *PixPat) PatXValid() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[26:28]))
}

// SetPatXValid updates the PatXValid field in the record's packed storage.
func (s *PixPat) SetPatXValid(v int16) {
	binary.NativeEndian.PutUint16(s.storage[26:28], uint16(v))
}

// PatXMap returns the PatXMap field from the record's packed storage.
func (s *PixPat) PatXMap() kernel.Pointer {
	return *(*kernel.Pointer)(unsafe.Pointer(&s.storage[28]))
}

// SetPatXMap updates the PatXMap field in the record's packed storage.
func (s *PixPat) SetPatXMap(v kernel.Pointer) {
	*(*kernel.Pointer)(unsafe.Pointer(&s.storage[28])) = v
}

// Pat1Data returns the Pat1Data field from the record's packed storage.
func (s *PixPat) Pat1Data() Pattern {
	return *(*Pattern)(unsafe.Pointer(&s.storage[36]))
}

// SetPat1Data updates the Pat1Data field in the record's packed storage.
func (s *PixPat) SetPat1Data(v Pattern) {
	*(*Pattern)(unsafe.Pointer(&s.storage[36])) = v
}

// ProcessInfoExtendedRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/processinfoextendedrec
type ProcessInfoExtendedRec struct {
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
	storage [80]byte
}

// ProcessInfoLength returns the ProcessInfoLength field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessInfoLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetProcessInfoLength updates the ProcessInfoLength field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessInfoLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// ProcessName returns the ProcessName field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessName() *byte {
	return *(**byte)(unsafe.Pointer(&s.storage[4]))
}

// SetProcessName updates the ProcessName field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoExtendedRec) SetProcessName(v *byte) {
	*(**byte)(unsafe.Pointer(&s.storage[4])) = v
}

// ProcessNumber returns the ProcessNumber field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessNumber() *ProcessSerialNumber {
	return *(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[12]))
}

// SetProcessNumber updates the ProcessNumber field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoExtendedRec) SetProcessNumber(v *ProcessSerialNumber) {
	*(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[12])) = v
}

// ProcessType returns the ProcessType field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetProcessType updates the ProcessType field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// ProcessSignature returns the ProcessSignature field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessSignature() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetProcessSignature updates the ProcessSignature field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessSignature(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// ProcessMode returns the ProcessMode field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessMode() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetProcessMode updates the ProcessMode field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessMode(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// ProcessLocation returns the ProcessLocation field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessLocation() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetProcessLocation updates the ProcessLocation field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoExtendedRec) SetProcessLocation(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// ProcessSize returns the ProcessSize field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetProcessSize updates the ProcessSize field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// ProcessFreeMem returns the ProcessFreeMem field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessFreeMem() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetProcessFreeMem updates the ProcessFreeMem field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessFreeMem(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// ProcessLauncher returns the ProcessLauncher field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessLauncher() *ProcessSerialNumber {
	return *(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[48]))
}

// SetProcessLauncher updates the ProcessLauncher field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoExtendedRec) SetProcessLauncher(v *ProcessSerialNumber) {
	*(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[48])) = v
}

// ProcessLaunchDate returns the ProcessLaunchDate field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessLaunchDate() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetProcessLaunchDate updates the ProcessLaunchDate field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessLaunchDate(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// ProcessActiveTime returns the ProcessActiveTime field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessActiveTime() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetProcessActiveTime updates the ProcessActiveTime field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessActiveTime(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// ProcessAppRef returns the ProcessAppRef field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessAppRef() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetProcessAppRef updates the ProcessAppRef field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoExtendedRec) SetProcessAppRef(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// ProcessTempMemTotal returns the ProcessTempMemTotal field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessTempMemTotal() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[72:76]))
}

// SetProcessTempMemTotal updates the ProcessTempMemTotal field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessTempMemTotal(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[72:76], uint32(v))
}

// ProcessPurgeableTempMemTotal returns the ProcessPurgeableTempMemTotal field from the record's packed storage.
func (s *ProcessInfoExtendedRec) ProcessPurgeableTempMemTotal() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[76:80]))
}

// SetProcessPurgeableTempMemTotal updates the ProcessPurgeableTempMemTotal field in the record's packed storage.
func (s *ProcessInfoExtendedRec) SetProcessPurgeableTempMemTotal(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[76:80], uint32(v))
}

// ProcessInfoRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/processinforec
type ProcessInfoRec struct {
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
	storage [72]byte
}

// ProcessInfoLength returns the ProcessInfoLength field from the record's packed storage.
func (s *ProcessInfoRec) ProcessInfoLength() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetProcessInfoLength updates the ProcessInfoLength field in the record's packed storage.
func (s *ProcessInfoRec) SetProcessInfoLength(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// ProcessName returns the ProcessName field from the record's packed storage.
func (s *ProcessInfoRec) ProcessName() *byte {
	return *(**byte)(unsafe.Pointer(&s.storage[4]))
}

// SetProcessName updates the ProcessName field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoRec) SetProcessName(v *byte) {
	*(**byte)(unsafe.Pointer(&s.storage[4])) = v
}

// ProcessNumber returns the ProcessNumber field from the record's packed storage.
func (s *ProcessInfoRec) ProcessNumber() *ProcessSerialNumber {
	return *(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[12]))
}

// SetProcessNumber updates the ProcessNumber field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoRec) SetProcessNumber(v *ProcessSerialNumber) {
	*(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[12])) = v
}

// ProcessType returns the ProcessType field from the record's packed storage.
func (s *ProcessInfoRec) ProcessType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetProcessType updates the ProcessType field in the record's packed storage.
func (s *ProcessInfoRec) SetProcessType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], uint32(v))
}

// ProcessSignature returns the ProcessSignature field from the record's packed storage.
func (s *ProcessInfoRec) ProcessSignature() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetProcessSignature updates the ProcessSignature field in the record's packed storage.
func (s *ProcessInfoRec) SetProcessSignature(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], uint32(v))
}

// ProcessMode returns the ProcessMode field from the record's packed storage.
func (s *ProcessInfoRec) ProcessMode() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetProcessMode updates the ProcessMode field in the record's packed storage.
func (s *ProcessInfoRec) SetProcessMode(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], uint32(v))
}

// ProcessLocation returns the ProcessLocation field from the record's packed storage.
func (s *ProcessInfoRec) ProcessLocation() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[32:40]))
}

// SetProcessLocation updates the ProcessLocation field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoRec) SetProcessLocation(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[32:40], uint64(v))
}

// ProcessSize returns the ProcessSize field from the record's packed storage.
func (s *ProcessInfoRec) ProcessSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetProcessSize updates the ProcessSize field in the record's packed storage.
func (s *ProcessInfoRec) SetProcessSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[40:44], uint32(v))
}

// ProcessFreeMem returns the ProcessFreeMem field from the record's packed storage.
func (s *ProcessInfoRec) ProcessFreeMem() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetProcessFreeMem updates the ProcessFreeMem field in the record's packed storage.
func (s *ProcessInfoRec) SetProcessFreeMem(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[44:48], uint32(v))
}

// ProcessLauncher returns the ProcessLauncher field from the record's packed storage.
func (s *ProcessInfoRec) ProcessLauncher() *ProcessSerialNumber {
	return *(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[48]))
}

// SetProcessLauncher updates the ProcessLauncher field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoRec) SetProcessLauncher(v *ProcessSerialNumber) {
	*(**ProcessSerialNumber)(unsafe.Pointer(&s.storage[48])) = v
}

// ProcessLaunchDate returns the ProcessLaunchDate field from the record's packed storage.
func (s *ProcessInfoRec) ProcessLaunchDate() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetProcessLaunchDate updates the ProcessLaunchDate field in the record's packed storage.
func (s *ProcessInfoRec) SetProcessLaunchDate(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], uint32(v))
}

// ProcessActiveTime returns the ProcessActiveTime field from the record's packed storage.
func (s *ProcessInfoRec) ProcessActiveTime() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[60:64]))
}

// SetProcessActiveTime updates the ProcessActiveTime field in the record's packed storage.
func (s *ProcessInfoRec) SetProcessActiveTime(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[60:64], uint32(v))
}

// ProcessAppRef returns the ProcessAppRef field from the record's packed storage.
func (s *ProcessInfoRec) ProcessAppRef() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[64:72]))
}

// SetProcessAppRef updates the ProcessAppRef field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *ProcessInfoRec) SetProcessAppRef(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[64:72], uint64(v))
}

// RGBColor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/rgbcolor
type RGBColor struct {
	Red   uint16 // The magnitude of the red component
	Green uint16
	Blue  uint16
}

// SizeResourceRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/sizeresourcerec
type SizeResourceRec struct {
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

// Flags returns the Flags field from the record's packed storage.
func (s *SizeResourceRec) Flags() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetFlags updates the Flags field in the record's packed storage.
func (s *SizeResourceRec) SetFlags(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// PreferredHeapSize returns the PreferredHeapSize field from the record's packed storage.
func (s *SizeResourceRec) PreferredHeapSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[2:6]))
}

// SetPreferredHeapSize updates the PreferredHeapSize field in the record's packed storage.
func (s *SizeResourceRec) SetPreferredHeapSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[2:6], uint32(v))
}

// MinimumHeapSize returns the MinimumHeapSize field from the record's packed storage.
func (s *SizeResourceRec) MinimumHeapSize() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[6:10]))
}

// SetMinimumHeapSize updates the MinimumHeapSize field in the record's packed storage.
func (s *SizeResourceRec) SetMinimumHeapSize(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[6:10], uint32(v))
}

// SpeechChannelRecord - Represents a speech channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speechchannelrecord
type SpeechChannelRecord struct {
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

// Data returns the Data field from the record's packed storage.
func (s *SpeechChannelRecord) Data() [1]int {
	return *(*[1]int)(unsafe.Pointer(&s.storage[0]))
}

// SetData updates the Data field in the record's packed storage.
func (s *SpeechChannelRecord) SetData(v [1]int) {
	*(*[1]int)(unsafe.Pointer(&s.storage[0])) = v
}

// SpeechErrorInfo - Defines a speech error information structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speecherrorinfo
type SpeechErrorInfo struct {
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

// Count returns the Count field from the record's packed storage.
func (s *SpeechErrorInfo) Count() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetCount updates the Count field in the record's packed storage.
func (s *SpeechErrorInfo) SetCount(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Oldest returns the Oldest field from the record's packed storage.
func (s *SpeechErrorInfo) Oldest() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetOldest updates the Oldest field in the record's packed storage.
func (s *SpeechErrorInfo) SetOldest(v int16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// OldPos returns the OldPos field from the record's packed storage.
func (s *SpeechErrorInfo) OldPos() int {
	return int(binary.NativeEndian.Uint64(s.storage[4:12]))
}

// SetOldPos updates the OldPos field in the record's packed storage.
func (s *SpeechErrorInfo) SetOldPos(v int) {
	binary.NativeEndian.PutUint64(s.storage[4:12], uint64(v))
}

// Newest returns the Newest field from the record's packed storage.
func (s *SpeechErrorInfo) Newest() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[12:14]))
}

// SetNewest updates the Newest field in the record's packed storage.
func (s *SpeechErrorInfo) SetNewest(v int16) {
	binary.NativeEndian.PutUint16(s.storage[12:14], uint16(v))
}

// NewPos returns the NewPos field from the record's packed storage.
func (s *SpeechErrorInfo) NewPos() int {
	return int(binary.NativeEndian.Uint64(s.storage[14:22]))
}

// SetNewPos updates the NewPos field in the record's packed storage.
func (s *SpeechErrorInfo) SetNewPos(v int) {
	binary.NativeEndian.PutUint64(s.storage[14:22], uint64(v))
}

// SpeechStatusInfo - Defines a speech status information structure, which stores information about the status of a speech channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speechstatusinfo
type SpeechStatusInfo struct {
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

// OutputBusy returns the OutputBusy field from the record's packed storage.
func (s *SpeechStatusInfo) OutputBusy() bool {
	return s.storage[0] != 0
}

// SetOutputBusy updates the OutputBusy field in the record's packed storage.
func (s *SpeechStatusInfo) SetOutputBusy(v bool) {
	if v {
		s.storage[0] = 1
	} else {
		s.storage[0] = 0
	}
}

// OutputPaused returns the OutputPaused field from the record's packed storage.
func (s *SpeechStatusInfo) OutputPaused() bool {
	return s.storage[1] != 0
}

// SetOutputPaused updates the OutputPaused field in the record's packed storage.
func (s *SpeechStatusInfo) SetOutputPaused(v bool) {
	if v {
		s.storage[1] = 1
	} else {
		s.storage[1] = 0
	}
}

// InputBytesLeft returns the InputBytesLeft field from the record's packed storage.
func (s *SpeechStatusInfo) InputBytesLeft() int {
	return int(binary.NativeEndian.Uint64(s.storage[2:10]))
}

// SetInputBytesLeft updates the InputBytesLeft field in the record's packed storage.
func (s *SpeechStatusInfo) SetInputBytesLeft(v int) {
	binary.NativeEndian.PutUint64(s.storage[2:10], uint64(v))
}

// PhonemeCode returns the PhonemeCode field from the record's packed storage.
func (s *SpeechStatusInfo) PhonemeCode() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetPhonemeCode updates the PhonemeCode field in the record's packed storage.
func (s *SpeechStatusInfo) SetPhonemeCode(v int16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// SpeechVersionInfo - Defines a speech version information structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speechversioninfo
type SpeechVersionInfo struct {
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

// SynthType returns the SynthType field from the record's packed storage.
func (s *SpeechVersionInfo) SynthType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSynthType updates the SynthType field in the record's packed storage.
func (s *SpeechVersionInfo) SetSynthType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// SynthSubType returns the SynthSubType field from the record's packed storage.
func (s *SpeechVersionInfo) SynthSubType() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetSynthSubType updates the SynthSubType field in the record's packed storage.
func (s *SpeechVersionInfo) SetSynthSubType(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}

// SynthManufacturer returns the SynthManufacturer field from the record's packed storage.
func (s *SpeechVersionInfo) SynthManufacturer() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetSynthManufacturer updates the SynthManufacturer field in the record's packed storage.
func (s *SpeechVersionInfo) SetSynthManufacturer(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], uint32(v))
}

// SynthFlags returns the SynthFlags field from the record's packed storage.
func (s *SpeechVersionInfo) SynthFlags() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetSynthFlags updates the SynthFlags field in the record's packed storage.
func (s *SpeechVersionInfo) SetSynthFlags(v int32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// SynthVersion returns the SynthVersion field from the record's packed storage.
func (s *SpeechVersionInfo) SynthVersion() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetSynthVersion updates the SynthVersion field in the record's packed storage.
func (s *SpeechVersionInfo) SetSynthVersion(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], uint32(v))
}

// SpeechXtndData - Defines a speech extension data structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speechxtnddata
type SpeechXtndData struct {
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
	storage [6]byte
}

// SynthCreator returns the SynthCreator field from the record's packed storage.
func (s *SpeechXtndData) SynthCreator() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetSynthCreator updates the SynthCreator field in the record's packed storage.
func (s *SpeechXtndData) SetSynthCreator(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// SynthData returns the SynthData field from the record's packed storage.
func (s *SpeechXtndData) SynthData() [2]uint8 {
	return *(*[2]uint8)(unsafe.Pointer(&s.storage[4]))
}

// SetSynthData updates the SynthData field in the record's packed storage.
func (s *SpeechXtndData) SetSynthData(v [2]uint8) {
	*(*[2]uint8)(unsafe.Pointer(&s.storage[4])) = v
}

// StyleTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/styletable
type StyleTable struct {
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
	storage [58]byte
}

// FontClass returns the FontClass field from the record's packed storage.
func (s *StyleTable) FontClass() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetFontClass updates the FontClass field in the record's packed storage.
func (s *StyleTable) SetFontClass(v int16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Offset returns the Offset field from the record's packed storage.
func (s *StyleTable) Offset() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[2:6]))
}

// SetOffset updates the Offset field in the record's packed storage.
func (s *StyleTable) SetOffset(v int32) {
	binary.NativeEndian.PutUint32(s.storage[2:6], uint32(v))
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *StyleTable) Reserved() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[6:10]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *StyleTable) SetReserved(v int32) {
	binary.NativeEndian.PutUint32(s.storage[6:10], uint32(v))
}

// Indexes returns the Indexes field from the record's packed storage.
func (s *StyleTable) Indexes() [48]int8 {
	return *(*[48]int8)(unsafe.Pointer(&s.storage[10]))
}

// SetIndexes updates the Indexes field in the record's packed storage.
func (s *StyleTable) SetIndexes(v [48]int8) {
	*(*[48]int8)(unsafe.Pointer(&s.storage[10])) = v
}

// VDGammaRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/vdgammarecord
type VDGammaRecord struct {
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

// CsGTable returns the CsGTable field from the record's packed storage.
func (s *VDGammaRecord) CsGTable() uintptr {
	return uintptr(binary.NativeEndian.Uint64(s.storage[0:8]))
}

// SetCsGTable updates the CsGTable field in the record's packed storage.
//
// v is stored as a raw address. The storage is a byte array, so the garbage
// collector scans it as integers and an address written here keeps nothing
// alive. v must name C-owned memory, or Go memory the caller pins with
// runtime.Pinner for as long as C reads the record.
func (s *VDGammaRecord) SetCsGTable(v uintptr) {
	binary.NativeEndian.PutUint64(s.storage[0:8], uint64(v))
}

// VoiceDescription - Defines a voice description structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/voicedescription
type VoiceDescription struct {
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
	storage [362]byte
}

// Length returns the Length field from the record's packed storage.
func (s *VoiceDescription) Length() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetLength updates the Length field in the record's packed storage.
func (s *VoiceDescription) SetLength(v int32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Voice returns the Voice field from the record's packed storage.
func (s *VoiceDescription) Voice() VoiceSpec {
	return *(*VoiceSpec)(unsafe.Pointer(&s.storage[4]))
}

// SetVoice updates the Voice field in the record's packed storage.
func (s *VoiceDescription) SetVoice(v VoiceSpec) {
	*(*VoiceSpec)(unsafe.Pointer(&s.storage[4])) = v
}

// Version returns the Version field from the record's packed storage.
func (s *VoiceDescription) Version() int32 {
	return int32(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *VoiceDescription) SetVersion(v int32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], uint32(v))
}

// Name returns the Name field from the record's packed storage.
func (s *VoiceDescription) Name() [64]uint8 {
	return *(*[64]uint8)(unsafe.Pointer(&s.storage[16]))
}

// SetName updates the Name field in the record's packed storage.
func (s *VoiceDescription) SetName(v [64]uint8) {
	*(*[64]uint8)(unsafe.Pointer(&s.storage[16])) = v
}

// Comment returns the Comment field from the record's packed storage.
func (s *VoiceDescription) Comment() [256]uint8 {
	return *(*[256]uint8)(unsafe.Pointer(&s.storage[80]))
}

// SetComment updates the Comment field in the record's packed storage.
func (s *VoiceDescription) SetComment(v [256]uint8) {
	*(*[256]uint8)(unsafe.Pointer(&s.storage[80])) = v
}

// Gender returns the Gender field from the record's packed storage.
func (s *VoiceDescription) Gender() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[336:338]))
}

// SetGender updates the Gender field in the record's packed storage.
func (s *VoiceDescription) SetGender(v int16) {
	binary.NativeEndian.PutUint16(s.storage[336:338], uint16(v))
}

// Age returns the Age field from the record's packed storage.
func (s *VoiceDescription) Age() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[338:340]))
}

// SetAge updates the Age field in the record's packed storage.
func (s *VoiceDescription) SetAge(v int16) {
	binary.NativeEndian.PutUint16(s.storage[338:340], uint16(v))
}

// Script returns the Script field from the record's packed storage.
func (s *VoiceDescription) Script() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[340:342]))
}

// SetScript updates the Script field in the record's packed storage.
func (s *VoiceDescription) SetScript(v int16) {
	binary.NativeEndian.PutUint16(s.storage[340:342], uint16(v))
}

// Language returns the Language field from the record's packed storage.
func (s *VoiceDescription) Language() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[342:344]))
}

// SetLanguage updates the Language field in the record's packed storage.
func (s *VoiceDescription) SetLanguage(v int16) {
	binary.NativeEndian.PutUint16(s.storage[342:344], uint16(v))
}

// Region returns the Region field from the record's packed storage.
func (s *VoiceDescription) Region() int16 {
	return int16(binary.NativeEndian.Uint16(s.storage[344:346]))
}

// SetRegion updates the Region field in the record's packed storage.
func (s *VoiceDescription) SetRegion(v int16) {
	binary.NativeEndian.PutUint16(s.storage[344:346], uint16(v))
}

// Reserved returns the Reserved field from the record's packed storage.
func (s *VoiceDescription) Reserved() [4]int32 {
	return *(*[4]int32)(unsafe.Pointer(&s.storage[346]))
}

// SetReserved updates the Reserved field in the record's packed storage.
func (s *VoiceDescription) SetReserved(v [4]int32) {
	*(*[4]int32)(unsafe.Pointer(&s.storage[346])) = v
}

// VoiceFileInfo - Defines a voice file information structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/voicefileinfo
type VoiceFileInfo struct {
	FileSpec [70]byte // A file system specification structure that contains the volume, directory, and name of the file containing the voice. Generally, files containing a single voice are of type `kTextToSpeechVoiceFileType`, and files containing multiple voices are of type `kTextToSpeechVoiceBundleType`.
	ResID    int16    // The resource ID of the voice in the file. Voices are stored in resources of type `kTextToSpeechVoiceType`.

}

// VoiceSpec - Defines a voice specification structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/voicespec
type VoiceSpec struct {
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

// Creator returns the Creator field from the record's packed storage.
func (s *VoiceSpec) Creator() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[0:4]))
}

// SetCreator updates the Creator field in the record's packed storage.
func (s *VoiceSpec) SetCreator(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[0:4], uint32(v))
}

// Id returns the Id field from the record's packed storage.
func (s *VoiceSpec) Id() uint32 {
	return uint32(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetId updates the Id field in the record's packed storage.
func (s *VoiceSpec) SetId(v uint32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], uint32(v))
}
