// Code generated from Apple documentation for ApplicationServices. DO NOT EDIT.

package applicationservices

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/kernel"
)

// C struct types

// ATSFSSpec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsfsspec
type ATSFSSpec struct {
	ParID   unsafe.Pointer
	VRefNum int16
	Name    string
}

// ATSFlatDataFontNameDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatafontnamedataheader
type ATSFlatDataFontNameDataHeader struct {
	NameSpecifierType uint32
	NameSpecifierSize uint32
}

// ATSFlatDataFontSpecRawNameData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatafontspecrawnamedata
type ATSFlatDataFontSpecRawNameData struct {
	FontNameType     uint16
	FontNamePlatform uint16
	FontNameScript   uint16
	FontNameLanguage uint16
	FontNameLength   uint32
}

// ATSFlatDataFontSpecRawNameDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatafontspecrawnamedataheader
type ATSFlatDataFontSpecRawNameDataHeader struct {
	NumberOfFlattenedNames uint32
	NameDataArray          ATSFlatDataFontSpecRawNameData
}

// ATSFlatDataLayoutControlsDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatalayoutcontrolsdataheader
type ATSFlatDataLayoutControlsDataHeader struct {
	NumberOfLayoutControls uint32
	ControlArray           ATSUAttributeInfo
}

// ATSFlatDataLineInfoData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatalineinfodata
type ATSFlatDataLineInfoData struct {
	LineLength           uint32
	NumberOfLineControls uint32
}

// ATSFlatDataLineInfoHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatalineinfoheader
type ATSFlatDataLineInfoHeader struct {
	NumberOfLines uint32
	LineInfoArray ATSFlatDataLineInfoData
}

// ATSFlatDataMainHeaderBlock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatamainheaderblock
type ATSFlatDataMainHeaderBlock struct {
	Version             uint32
	SizeOfDataBlock     uint32
	OffsetToTextLayouts uint32
	OffsetToStyleRuns   uint32
	OffsetToStyleList   uint32
}

// ATSFlatDataStyleListFeatureData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastylelistfeaturedata
type ATSFlatDataStyleListFeatureData struct {
	TheFeatureType     uint16
	TheFeatureSelector uint16
}

// ATSFlatDataStyleListHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastylelistheader
type ATSFlatDataStyleListHeader struct {
	NumberOfStyles uint32
	StyleDataArray ATSFlatDataStyleListStyleDataHeader
}

// ATSFlatDataStyleListStyleDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastyleliststyledataheader
type ATSFlatDataStyleListStyleDataHeader struct {
	SizeOfStyleInfo       uint32
	NumberOfSetAttributes uint32
	NumberOfSetFeatures   uint32
	NumberOfSetVariations uint32
}

// ATSFlatDataStyleListVariationData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastylelistvariationdata
type ATSFlatDataStyleListVariationData struct {
	TheVariationAxis  uint32
	TheVariationValue int32
}

// ATSFlatDataStyleRunDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatastylerundataheader
type ATSFlatDataStyleRunDataHeader struct {
	NumberOfStyleRuns uint32
	StyleRunArray     ATSUStyleRunInfo
}

// ATSFlatDataTextLayoutDataHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatatextlayoutdataheader
type ATSFlatDataTextLayoutDataHeader struct {
	SizeOfLayoutData       uint32
	TextLayoutLength       uint32
	OffsetToLayoutControls uint32
	OffsetToLineInfo       uint32
}

// ATSFlatDataTextLayoutHeader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsflatdatatextlayoutheader
type ATSFlatDataTextLayoutHeader struct {
	NumFlattenedTextLayouts uint32
	FlattenedTextLayouts    ATSFlatDataTextLayoutDataHeader
}

// ATSFontFilter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsfontfilter
type ATSFontFilter struct {
	Version        uint32
	FilterSelector ATSFontFilterSelector
	Filter         unsafe.Pointer
}

// ATSFontMetrics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsfontmetrics
type ATSFontMetrics struct {
	Version             uint32
	Ascent              float64
	Descent             float64
	Leading             float64
	AvgAdvanceWidth     float64
	MaxAdvanceWidth     float64
	MinLeftSideBearing  float64
	MinRightSideBearing float64
	StemWidth           float64
	StemHeight          float64
	CapHeight           float64
	XHeight             float64
	ItalicAngle         float64
	UnderlinePosition   float64
	UnderlineThickness  float64
}

// ATSFontQuerySourceContext
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsfontquerysourcecontext
type ATSFontQuerySourceContext struct {
	Version uint32
	RefCon  unsafe.Pointer
	Retain  corefoundation.CFAllocatorRetainCallBack
	Release corefoundation.CFAllocatorReleaseCallBack
}

// ATSGlyphIdealMetrics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsglyphidealmetrics
type ATSGlyphIdealMetrics struct {
	Advance          unsafe.Pointer
	SideBearing      unsafe.Pointer
	OtherSideBearing unsafe.Pointer
}

// ATSGlyphScreenMetrics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsglyphscreenmetrics
type ATSGlyphScreenMetrics struct {
	DeviceAdvance    unsafe.Pointer
	TopLeft          unsafe.Pointer
	Height           uint32
	Width            uint32
	SideBearing      unsafe.Pointer
	OtherSideBearing unsafe.Pointer
}

// ATSJustWidthDeltaEntryOverride
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsjustwidthdeltaentryoverride
type ATSJustWidthDeltaEntryOverride struct {
	BeforeGrowLimit   int32
	BeforeShrinkLimit int32
	AfterGrowLimit    int32
	AfterShrinkLimit  int32
	GrowFlags         uint32
	ShrinkFlags       uint32
}

// ATSLayoutRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atslayoutrecord
type ATSLayoutRecord struct {
	GlyphID        ATSGlyphRef
	Flags          uint32
	OriginalOffset int
	RealPos        int32
}

// ATSTrapezoid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atstrapezoid
type ATSTrapezoid struct {
	UpperLeft  unsafe.Pointer
	UpperRight unsafe.Pointer
	LowerRight unsafe.Pointer
	LowerLeft  unsafe.Pointer
}

// ATSUAttributeInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuattributeinfo
type ATSUAttributeInfo struct {
	FTag       uint32
	FValueSize int
}

// ATSUCaret
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsucaret
type ATSUCaret struct {
	FX      int32
	FY      int32
	FDeltaX int32
	FDeltaY int32
}

// ATSUCurvePath
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsucurvepath
type ATSUCurvePath struct {
	Vectors     uint32
	Vector      unsafe.Pointer
	ControlBits uint32
}

// ATSUCurvePaths
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsucurvepaths
type ATSUCurvePaths struct {
	Contours uint32
	Contour  ATSUCurvePath
}

// ATSUGlyphInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuglyphinfo
type ATSUGlyphInfo struct {
	GlyphID     uint16
	Reserved    uint16
	LayoutFlags uint32
	CharIndex   uint32
	Style       unsafe.Pointer
	DeltaY      float32
	IdealX      float32
	ScreenX     unsafe.Pointer
	CaretX      unsafe.Pointer
}

// ATSUGlyphInfoArray
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuglyphinfoarray
type ATSUGlyphInfoArray struct {
	Layout    unsafe.Pointer
	NumGlyphs int
	Glyphs    ATSUGlyphInfo
}

// ATSUGlyphSelector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuglyphselector
type ATSUGlyphSelector struct {
	Collection uint16
	GlyphID    uint16
}

// ATSULayoutOperationOverrideSpecifier
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsulayoutoperationoverridespecifier
type ATSULayoutOperationOverrideSpecifier struct {
	OperationSelector uint32
	OverrideUPP       unsafe.Pointer
}

// ATSURGBAlphaColor
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsurgbalphacolor
type ATSURGBAlphaColor struct {
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}

// ATSUStyleRunInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsustyleruninfo
type ATSUStyleRunInfo struct {
	RunLength        uint32
	StyleObjectIndex uint32
}

// ATSUTab
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsutab
type ATSUTab struct {
	TabPosition int32
	TabType     uint16
}

// ATSUUnhighlightData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/atsuunhighlightdata
type ATSUUnhighlightData struct {
	DataType        uint32
	UnhighlightData unsafe.Pointer
}

// AppParameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/appparameters
type AppParameters struct {
	EventRefCon   uint32
	MessageLength uint32
	TheMsgEvent   unsafe.Pointer
	What          uint16
	Message       uint32
	When          uint32
	Where         Point
	Modifiers     uint16
}

// AsscEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/asscentry
type AsscEntry struct {
	FontSize  unsafe.Pointer
	FontStyle unsafe.Pointer
	FontID    unsafe.Pointer
}

// BitMap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/bitmap
type BitMap struct {
	BaseAddr kernel.Ptr
	RowBytes unsafe.Pointer
	Bounds   Rect
}

// CM2Profile
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cm2profile
type CM2Profile struct {
	Header   unsafe.Pointer
	TagTable unsafe.Pointer
	ElemData unsafe.Pointer
}

// CMDeviceInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmdeviceinfo
type CMDeviceInfo struct {
	DataVersion      uint32
	DeviceClass      uint32
	DeviceID         uint32
	DeviceScope      CMDeviceScope
	DeviceState      uint32
	DefaultProfileID uint32
	DeviceName       corefoundation.CFDictionaryRef // See the CFDictionary documentation for a description of the [CFDictionaryRef] data type.
	ProfileCount     uint32
	Reserved         uint32
}

// CMDeviceProfileArray
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmdeviceprofilearray
type CMDeviceProfileArray struct {
	ProfileCount uint32
	Profiles     unsafe.Pointer
}

// CMDeviceScope
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmdevicescope
type CMDeviceScope struct {
	DeviceUser corefoundation.CFStringRef
	DeviceHost corefoundation.CFStringRef
}

// CMMultiFunctLutType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmmultifunctluttype
type CMMultiFunctLutType struct {
	TypeDescriptor uint32
	Reserved       uint32
	InputChannels  uint8
	OutputChannels uint8
	Reserved2      uint16
	OffsetBcurves  uint32
	OffsetMatrix   uint32
	OffsetMcurves  uint32
	OffsetCLUT     uint32
	OffsetAcurves  uint32
	Data           uint8
}

// CMXYZColor - Contains values for a color specified in XYZ color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cmxyzcolor
type CMXYZColor struct {
	X uint16
	Y uint16
	Z uint16
}

// CQDProcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/cqdprocs
type CQDProcs struct {
	TextProc          unsafe.Pointer
	LineProc          unsafe.Pointer
	RectProc          unsafe.Pointer
	RRectProc         unsafe.Pointer
	OvalProc          unsafe.Pointer
	ArcProc           unsafe.Pointer
	PolyProc          unsafe.Pointer
	RgnProc           unsafe.Pointer
	BitsProc          unsafe.Pointer
	CommentProc       unsafe.Pointer
	TxMeasProc        unsafe.Pointer
	GetPicProc        unsafe.Pointer
	PutPicProc        unsafe.Pointer
	OpcodeProc        unsafe.Pointer
	NewProc1          unsafe.Pointer
	GlyphsProc        unsafe.Pointer
	PrinterStatusProc unsafe.Pointer
	NewProc4          unsafe.Pointer
	NewProc5          unsafe.Pointer
	NewProc6          unsafe.Pointer
}

// ColorSpec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/colorspec
type ColorSpec struct {
	Value unsafe.Pointer
	Rgb   RGBColor
}

// ColorTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/colortable
type ColorTable struct {
	CtSeed  unsafe.Pointer
	CtFlags unsafe.Pointer
	CtSize  unsafe.Pointer
	CtTable unsafe.Pointer
}

// DelimiterInfo - Defines a delimiter information structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/delimiterinfo
type DelimiterInfo struct {
	StartDelimiter uint8 // The start delimiter for an embedded command. By default, the start delimiter is “`[[`”.
	EndDelimiter   uint8 // The end delimiter for an embedded command. By default, the end delimiter is “`]]`”.

}

// FMFilter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfilter
type FMFilter struct {
	Selector int32
	Filter   unsafe.Pointer
	Format   uint32
}

// FMFontDirectoryFilter
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfontdirectoryfilter
type FMFontDirectoryFilter struct {
	FontFolderDomain unsafe.Pointer
	Reserved         uint32
}

// FMFontFamilyInstance
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfontfamilyinstance
type FMFontFamilyInstance struct {
	FontFamily int16
	FontStyle  int16
}

// FMFontFamilyInstanceIterator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfontfamilyinstanceiterator
type FMFontFamilyInstanceIterator struct {
	Reserved uint32
}

// FMFontFamilyIterator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfontfamilyiterator
type FMFontFamilyIterator struct {
	Reserved uint32
}

// FMFontIterator
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fmfontiterator
type FMFontIterator struct {
	Reserved uint32
}

// FMInput
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fminput
type FMInput struct {
	Family   unsafe.Pointer
	Size     unsafe.Pointer
	Face     unsafe.Pointer
	NeedBits unsafe.Pointer
	Device   unsafe.Pointer
	Numer    Point
	Denom    Point
}

// FamRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/famrec
type FamRec struct {
	FfFlags     unsafe.Pointer
	FfFamID     unsafe.Pointer
	FfFirstChar unsafe.Pointer
	FfLastChar  unsafe.Pointer
	FfAscent    unsafe.Pointer
	FfDescent   unsafe.Pointer
	FfLeading   unsafe.Pointer
	FfWidMax    unsafe.Pointer
	FfWTabOff   unsafe.Pointer
	FfKernOff   unsafe.Pointer
	FfStylOff   unsafe.Pointer
	FfVersion   unsafe.Pointer
	FfIntl      unsafe.Pointer
	FfProperty  unsafe.Pointer
}

// FontAssoc
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fontassoc
type FontAssoc struct {
	NumAssoc unsafe.Pointer
}

// FontInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fontinfo
type FontInfo struct {
	Ascent  unsafe.Pointer
	Descent unsafe.Pointer
	WidMax  unsafe.Pointer
	Leading unsafe.Pointer
}

// FontRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/fontrec
type FontRec struct {
	FontType    unsafe.Pointer
	FirstChar   unsafe.Pointer
	LastChar    unsafe.Pointer
	WidMax      unsafe.Pointer
	KernMax     unsafe.Pointer
	NDescent    unsafe.Pointer
	FRectWidth  unsafe.Pointer
	FRectHeight unsafe.Pointer
	OwTLoc      uint16
	Ascent      unsafe.Pointer
	Descent     unsafe.Pointer
	Leading     unsafe.Pointer
	RowWords    unsafe.Pointer
}

// GDevice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/gdevice
type GDevice struct {
	GdRefNum     unsafe.Pointer
	GdID         unsafe.Pointer
	GdType       unsafe.Pointer
	GdITable     unsafe.Pointer
	GdResPref    unsafe.Pointer
	GdSearchProc unsafe.Pointer
	GdCompProc   unsafe.Pointer
	GdFlags      unsafe.Pointer
	GdPMap       unsafe.Pointer
	GdRefCon     unsafe.Pointer
	GdNextGD     unsafe.Pointer
	GdRect       Rect
	GdMode       unsafe.Pointer
	GdCCBytes    unsafe.Pointer
	GdCCDepth    unsafe.Pointer
	GdCCXData    unsafe.Pointer
	GdCCXMask    unsafe.Pointer
	GdExt        unsafe.Pointer
}

// GrafPort
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/grafport
type GrafPort struct {
	Whatever unsafe.Pointer
}

// ICAppSpec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icappspec
type ICAppSpec struct {
	FCreator uint32
	Name     unsafe.Pointer
}

// ICAppSpecList
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icappspeclist
type ICAppSpecList struct {
	NumberOfItems unsafe.Pointer
	AppSpecs      ICAppSpec
}

// ICCharTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icchartable
type ICCharTable struct {
	NetToMac uint8
	MacToNet uint8
}

// ICFileSpec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icfilespec
type ICFileSpec struct {
	VolName         unsafe.Pointer
	VolCreationDate unsafe.Pointer
	Fss             unsafe.Pointer
	Alias           unsafe.Pointer
}

// ICFontRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icfontrecord
type ICFontRecord struct {
	Size unsafe.Pointer
	Face unsafe.Pointer
	Pad  unsafe.Pointer
	Font unsafe.Pointer
}

// ICMapEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icmapentry
type ICMapEntry struct {
	TotalLength    unsafe.Pointer
	FixedLength    int32
	Version        unsafe.Pointer
	FileType       uint32
	FileCreator    uint32
	PostCreator    uint32
	Flags          int16
	Extension      unsafe.Pointer
	CreatorAppName unsafe.Pointer
	PostAppName    unsafe.Pointer
	MIMEType       unsafe.Pointer
	EntryName      unsafe.Pointer
}

// ICServiceEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icserviceentry
type ICServiceEntry struct {
	Name  unsafe.Pointer
	Port  unsafe.Pointer
	Flags int16
}

// ICServices
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/icservices
type ICServices struct {
	Count    unsafe.Pointer
	Services ICServiceEntry
}

// KernEntry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/kernentry
type KernEntry struct {
	KernStyle  unsafe.Pointer
	KernLength unsafe.Pointer
}

// KernPair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/kernpair
type KernPair struct {
	KernFirst  unsafe.Pointer
	KernSecond unsafe.Pointer
	KernWidth  unsafe.Pointer
}

// KernTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/kerntable
type KernTable struct {
	NumKerns unsafe.Pointer
}

// LaunchParamBlockRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/launchparamblockrec
type LaunchParamBlockRec struct {
	Reserved1           uint32
	Reserved2           uint16
	LaunchBlockID       uint16
	LaunchEPBLength     uint32
	LaunchFileFlags     uint16
	LaunchControlFlags  uint16
	LaunchAppRef        unsafe.Pointer
	LaunchProcessSN     *ProcessSerialNumber
	LaunchPreferredSize uint32
	LaunchMinimumSize   uint32
	LaunchAvailableSize uint32
	LaunchAppParameters unsafe.Pointer
}

// MacPolygon
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/macpolygon
type MacPolygon struct {
	PolySize   unsafe.Pointer
	PolyBBox   Rect
	PolyPoints Point
}

// NameTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/nametable
type NameTable struct {
	StringCount  unsafe.Pointer
	BaseFontName unsafe.Pointer
}

// OpenCPicParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/opencpicparams
type OpenCPicParams struct {
	SrcRect   Rect           // The optimal bounding rectangle for the resolution indicated by the `hRes` and `vRes` fields. To display a picture at a resolution other than that specified in the `hRes` and `vRes` fields, your application should compute an appropriate destination rectangle by scaling the image’s width and height by the destination resolution divided by the source resolution.
	HRes      int32          // The best horizontal resolution for the picture. A value of 0x0048000 specifies a horizontal resolution of 72 dpi.
	VRes      int32          // The best vertical resolution for the picture. A value of 0x0048000 specifies a vertical resolution of 72 dpi.
	Version   unsafe.Pointer // Always set this field to -2.
	Reserved1 unsafe.Pointer // Reserved; set to 0.
	Reserved2 unsafe.Pointer // Reserved; set to 0.

}

// PMLanguageInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/pmlanguageinfo
type PMLanguageInfo struct {
	Level   unsafe.Pointer // Specifies the level of the imaging language used by the printer driver.
	Version unsafe.Pointer // Specifies the version of the imaging language.
	Release unsafe.Pointer // Specifies the release of the imaging language.

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
	Pat uint8
}

// PhonemeDescriptor - Defines a phoneme descriptor structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/phonemedescriptor
type PhonemeDescriptor struct {
	PhonemeCount unsafe.Pointer // The number of phonemes that the current synthesizer defines. Typically, this will correspond to the number of phonemes in the language supported by the synthesizer.
	ThePhonemes  PhonemeInfo    // An array of phoneme information structures.

}

// PhonemeInfo - Defines a structure that stores information about a phoneme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/phonemeinfo
type PhonemeInfo struct {
	Opcode      unsafe.Pointer // The opcode for the phoneme.
	PhStr       unsafe.Pointer // The string used to represent the phoneme. The string does not necessarily have a phonetic connection to the phoneme, but might simply be an abstract textual representation of it.
	ExampleStr  unsafe.Pointer // An example word that illustrates use of the phoneme.
	HiliteStart unsafe.Pointer // The number of characters in the example word that precede the portion of that word representing the phoneme.
	HiliteEnd   unsafe.Pointer // The number of characters between the beginning of the example word and the end of the portion of that word representing the phoneme.

}

// Picture
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/picture
type Picture struct {
	PicSize  unsafe.Pointer
	PicFrame Rect
}

// PixMap
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/pixmap
type PixMap struct {
	BaseAddr    kernel.Ptr     // For an onscreen pixel image, a pointer to the first byte of the image. For optimal performance, this should be a multiple of 4. The `baseAddr` field of the [PixMap] record for an offscreen graphics world contains a handle instead of a pointer. Your application should never directly access the `baseAddr` field of the [PixMap] record for an offscreen graphics world.
	RowBytes    unsafe.Pointer // The offset in bytes from one row of the image to the next. The value must be even, less than 0x4000, and for best performance it should be a multiple of 4. The high 2 bits of `rowBytes` are used as flags. If bit 15 =1, the data structure pointed to is a [PixMap] structure; otherwise it is a [BitMap](<doc://com.apple.documentation/documentation/applicationservices/bitmap>) structure.
	Bounds      Rect           // The boundary rectangle, which links the local coordinate system of a graphics port to QuickDraw's global coordinate system and defines the area of the bit image into which QuickDraw can draw. By default, the boundary rectangle is the entire main screen. Do not use the `value` of this field to determine the size of the screen; instead use the `value` of the `gdRect` field of the [GDevice](<doc://com.apple.documentation/documentation/applicationservices/gdevice>) structure for the screen.
	PmVersion   unsafe.Pointer // The version number of Color QuickDraw that created this [PixMap] structure. The value of `pmVersion` is normally 0. If `pmVersion` is 4, Color QuickDraw treats the [PixMap] record's `baseAddr` field as 32-bit clean. All other flags are private. Most applications never need to set this field
	PackType    unsafe.Pointer // The packing algorithm used to compress image data. Color QuickDraw currently supports a `packType` of 0, which means no packing, and values of 1 to 4 for packing direct pixels.
	PackSize    unsafe.Pointer // The size of the packed image in bytes. When the `packType` field contains the `value` 0, this field is always set to 0.
	HRes        int32          // The horizontal resolution of the pixel image in pixels per inch. By default, this value is 0x00480000 (for 72 pixels per inch).
	VRes        int32          // The vertical resolution of the pixel image in pixels per inch. By default, this value is 0x00480000 (for 72 pixels per inch).
	PixelType   unsafe.Pointer // The storage format for a pixel image. Indexed pixels are indicated by a value of 0. Direct pixels are specified by a value of [RGBDirect], or 16. In the [PixMap] record of the [GDevice](<doc://com.apple.documentation/documentation/applicationservices/gdevice>) structure for a direct device, this field is set to [RGBDirect] when the screen depth is set.
	PixelSize   unsafe.Pointer // The number of bits used to represent a pixel. Indexed pixels can have sizes of 1, 2, 4, and 8 bits; direct pixel sizes are 16 and 32 bits.
	CmpCount    unsafe.Pointer // The number of components used to represent a color for a pixel. With indexed pixels, each pixel is a single value representing an index in a color table, and therefore this field contains the value 1; the index is the single component. With direct pixels, each pixel contains three components (one integer each for the intensities of red, green, and blue) so this field contains the value 3.
	CmpSize     unsafe.Pointer // The size in bits of each component for a pixel.
	PixelFormat uint32         // The way the pixels are arranged; see `k4444YpCbCrA8PixelFormat`.
	PmTable     unsafe.Pointer // Color map for this structure.
	PmExt       unsafe.Pointer // [Handle] to a PixMapExtension structure. Set to [NIL] if there is no extension.

}

// PixPat
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/pixpat
type PixPat struct {
	PatType   unsafe.Pointer
	PatMap    unsafe.Pointer
	PatData   unsafe.Pointer
	PatXData  unsafe.Pointer
	PatXValid unsafe.Pointer
	PatXMap   unsafe.Pointer
	Pat1Data  Pattern
}

// ProcessInfoExtendedRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/processinfoextendedrec
type ProcessInfoExtendedRec struct {
	ProcessInfoLength            uint32
	ProcessName                  *byte
	ProcessNumber                *ProcessSerialNumber
	ProcessType                  uint32
	ProcessSignature             uint32
	ProcessMode                  uint32
	ProcessLocation              kernel.Ptr
	ProcessSize                  uint32
	ProcessFreeMem               uint32
	ProcessLauncher              *ProcessSerialNumber
	ProcessLaunchDate            uint32
	ProcessActiveTime            uint32
	ProcessAppRef                unsafe.Pointer
	ProcessTempMemTotal          uint32
	ProcessPurgeableTempMemTotal uint32
}

// ProcessInfoRec
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/processinforec
type ProcessInfoRec struct {
	ProcessInfoLength uint32
	ProcessName       *byte
	ProcessNumber     *ProcessSerialNumber
	ProcessType       uint32
	ProcessSignature  uint32
	ProcessMode       uint32
	ProcessLocation   kernel.Ptr
	ProcessSize       uint32
	ProcessFreeMem    uint32
	ProcessLauncher   *ProcessSerialNumber
	ProcessLaunchDate uint32
	ProcessActiveTime uint32
	ProcessAppRef     unsafe.Pointer
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
	Flags             uint16
	PreferredHeapSize uint32
	MinimumHeapSize   uint32
}

// SpeechChannelRecord - Represents a speech channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speechchannelrecord
type SpeechChannelRecord struct {
	Data int
}

// SpeechErrorInfo - Defines a speech error information structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speecherrorinfo
type SpeechErrorInfo struct {
	Count  unsafe.Pointer // The number of errors that have occurred in processing the current text buffer since the last call to the [GetSpeechInfo] function with the `soErrors` selector. Of these errors, you can find information about only the first and last error that occurred.
	Oldest int16          // The error code of the first error that occurred after the previous call to the [GetSpeechInfo] function with the `soErrors` selector.
	OldPos int            // The character position within the text buffer being processed of the first error that occurred after the previous call to the [GetSpeechInfo] function with the `soErrors` selector.
	Newest int16          // The error code of the most recent error.
	NewPos int            // The character position within the text buffer being processed of the most recent error.

}

// SpeechStatusInfo - Defines a speech status information structure, which stores information about the status of a speech channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speechstatusinfo
type SpeechStatusInfo struct {
	OutputBusy     unsafe.Pointer // Whether the speech channel is currently producing speech. A speech channel is considered to be producing speech even at some times when no audio data is being produced through the Macintosh speaker. This occurs, for example, when the Speech Synthesis Manager is processing an input buffer but has not yet initiated speech or when speech output is paused.
	OutputPaused   unsafe.Pointer // Whether speech output in the speech channel has been paused by a call to the [PauseSpeechAt(_:_:)](<doc://com.apple.documentation/documentation/applicationservices/1461174-pausespeechat>) function.
	InputBytesLeft int            // The number of input bytes of the text that the speech channel must still process. When `inputBytesLeft` is 0, the buffer of input text passed to one of the [SpeakText] or [SpeakBuffer] functions may be disposed of. When you call the [SpeakString] function, the Speech Synthesis Manager stores a duplicate of the string to be spoken in an internal buffer; thus, you may delete the original string immediately after calling [SpeakString].
	PhonemeCode    unsafe.Pointer // The opcode for the phoneme that the speech channel is currently processing.

}

// SpeechVersionInfo - Defines a speech version information structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speechversioninfo
type SpeechVersionInfo struct {
	SynthType         uint32         // The general type of the synthesizer. For the current version of the Speech Synthesis Manager, this field always contains the value `kTextToSpeechSynthType`, indicating that the synthesizer converts text into speech.
	SynthSubType      uint32         // The specific type of the synthesizer. Currently, no specific types of synthesizer are defined. If you define a new type of synthesizer, you should register the four-character code for your type with Developer Technical Support.
	SynthManufacturer uint32         // A unique identification of a synthesizer engine. If you develop synthesizers, then you should register a different four-character code for each synthesizer you develop with Developer Technical Support. The `creatorID` field of the voice specification structure and the `synthCreator` field of a speech extension data structure should each be set to the value stored in this field for the desired synthesizer.
	SynthFlags        unsafe.Pointer // A set of flags indicating which synthesizer features are activated. Specific constants define the bits in this field whose meanings are defined for all synthesizers.
	SynthVersion      uint32         // The version number of the synthesizer.

}

// SpeechXtndData - Defines a speech extension data structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/speechxtnddata
type SpeechXtndData struct {
	SynthCreator uint32 // The synthesizer’s creator ID, identical to the value stored in the `synthManufacturer` field of a speech version information structure. You should set this field to the appropriate value before calling [GetSpeechInfo] or [SetSpeechInfo].
	SynthData    uint8  // Synthesizer-specific data. The size and format of the data in this field may vary.

}

// StyleTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/styletable
type StyleTable struct {
	FontClass unsafe.Pointer
	Offset    unsafe.Pointer
	Reserved  unsafe.Pointer
	Indexes   unsafe.Pointer
}

// VDGammaRecord
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/vdgammarecord
type VDGammaRecord struct {
	CsGTable kernel.Ptr // A pointer to a gamma table.

}

// VoiceDescription - Defines a voice description structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/voicedescription
type VoiceDescription struct {
	Length   unsafe.Pointer // The size of the voice description structure, in bytes.
	Voice    VoiceSpec      // A voice specification structure that uniquely identifies the voice.
	Version  unsafe.Pointer // The version number of the voice.
	Name     unsafe.Pointer // The name of the voice, preceded by a length byte. Names must be 63 characters or less.
	Comment  unsafe.Pointer // Additional text information about the voice. Some synthesizers use this field to store a phrase that can be spoken.
	Gender   unsafe.Pointer // The gender of the individual represented by the voice. See [Gender Constants](<doc://com.apple.documentation/documentation/applicationservices/speech_synthesis_manager/1552246-gender_constants>).
	Age      unsafe.Pointer // The approximate age in years of the individual represented by the voice.
	Script   unsafe.Pointer // The encoding code of the text that the voice can process.
	Language unsafe.Pointer // A code that indicates the language of voice output.
	Region   unsafe.Pointer // A code that indicates the region represented by the voice.
	Reserved unsafe.Pointer // Reserved. May be used to hold a 32-bit encoding value, if necessary (see the description of the `script` field for more information).

}

// VoiceFileInfo - Defines a voice file information structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/voicefileinfo
type VoiceFileInfo struct {
	FileSpec unsafe.Pointer // A file system specification structure that contains the volume, directory, and name of the file containing the voice. Generally, files containing a single voice are of type `kTextToSpeechVoiceFileType`, and files containing multiple voices are of type `kTextToSpeechVoiceBundleType`.
	ResID    unsafe.Pointer // The resource ID of the voice in the file. Voices are stored in resources of type `kTextToSpeechVoiceType`.

}

// VoiceSpec - Defines a voice specification structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/applicationservices/voicespec
type VoiceSpec struct {
	Creator uint32 // The synthesizer that is required to use the voice. This is equivalent to the value contained in the `synthManufacturer` field of a speech version information structure and that contained in the `synthCreator` field of a speech extension data structure. The set of [OSType] values specified entirely by space characters and lowercase letters is reserved.
	Id      uint32 // The voice ID of the voice for the synthesizer. Every voice on a synthesizer has a unique ID.

}
