// Code generated from Apple documentation. DO NOT EDIT.

package applicationservices

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coreservices"
	"github.com/tmc/apple/kernel"
)

// See: https://developer.apple.com/documentation/applicationservices/atscubicclosepathprocptr
type ATSCubicClosePathProcPtr = func(unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atscubicclosepathupp
type ATSCubicClosePathUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atscubiccurvetoprocptr
type ATSCubicCurveToProcPtr = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atscubiccurvetoupp
type ATSCubicCurveToUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atscubiclinetoprocptr
type ATSCubicLineToProcPtr = func(unsafe.Pointer, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atscubiclinetoupp
type ATSCubicLineToUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atscubicmovetoprocptr
type ATSCubicMoveToProcPtr = func(unsafe.Pointer, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atscubicmovetoupp
type ATSCubicMoveToUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atscurvetype
type ATSCurveType = uint16

// See: https://developer.apple.com/documentation/applicationservices/atsflatdatafontspecifertype
type ATSFlatDataFontSpeciferType = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsfontapplierfunction
type ATSFontApplierFunction = func(ATSFontRef, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atsfontautoactivationsetting
type ATSFontAutoActivationSetting = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsfontcontainerref
type ATSFontContainerRef = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsfontcontext
type ATSFontContext = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsfontfamilyapplierfunction
type ATSFontFamilyApplierFunction = func(ATSFontFamilyRef, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atsfontfamilyiterator
type ATSFontFamilyIterator = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsfontfamilyref
type ATSFontFamilyRef = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsfontformat
type ATSFontFormat = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsfontiterator
type ATSFontIterator = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsfontnotificationinforef
type ATSFontNotificationInfoRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/atsfontnotificationref
type ATSFontNotificationRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/atsfontquerycallback
type ATSFontQueryCallback = func(ATSFontQueryMessageID, corefoundation.CFPropertyListRef, unsafe.Pointer) corefoundation.CFPropertyListRef

// See: https://developer.apple.com/documentation/CoreText/ATSFontRef
type ATSFontRef = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsfontsize
type ATSFontSize = float64

// See: https://developer.apple.com/documentation/applicationservices/atsgeneration
type ATSGeneration = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsglyphinfoflags
type ATSGlyphInfoFlags = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsglyphref
type ATSGlyphRef = uint16

// See: https://developer.apple.com/documentation/applicationservices/atsjustprioritywidthdeltaoverrides
// ATSJustPriorityWidthDeltaOverrides is opaque storage with the size and alignment C gives ATSJustPriorityWidthDeltaOverrides:
// 80 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 80 into.
type ATSJustPriorityWidthDeltaOverrides [40]uint16

// See: https://developer.apple.com/documentation/applicationservices/atslinelayoutoptions
type ATSLineLayoutOptions = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsnotificationcallback
type ATSNotificationCallback = func(ATSFontNotificationInfoRef, unsafe.Pointer)

// See: https://developer.apple.com/documentation/applicationservices/atsoptionflags
type ATSOptionFlags = uint32

// See: https://developer.apple.com/documentation/applicationservices/atspoint
// ATSPoint is opaque storage with the size and alignment C gives ATSPoint:
// 16 bytes. C declares a record here, not a handle, so a
// pointer-width rendering would hand the framework eight bytes to write
// 16 into.
type ATSPoint [2]uint64

// See: https://developer.apple.com/documentation/applicationservices/atsquadraticclosepathprocptr
type ATSQuadraticClosePathProcPtr = func(unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atsquadraticclosepathupp
type ATSQuadraticClosePathUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsquadraticcurveprocptr
type ATSQuadraticCurveProcPtr = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atsquadraticcurveupp
type ATSQuadraticCurveUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsquadraticlineprocptr
type ATSQuadraticLineProcPtr = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atsquadraticlineupp
type ATSQuadraticLineUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsquadraticnewpathprocptr
type ATSQuadraticNewPathProcPtr = func(unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/atsquadraticnewpathupp
type ATSQuadraticNewPathUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsstylerenderingoptions
type ATSStyleRenderingOptions = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsuattributetag
type ATSUAttributeTag = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsuattributevalueptr
type ATSUAttributeValuePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsubackgroundcolor
type ATSUBackgroundColor = ATSURGBAlphaColor

// See: https://developer.apple.com/documentation/applicationservices/atsubackgrounddatatype
type ATSUBackgroundDataType = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsucursormovementtype
type ATSUCursorMovementType = uint16

// See: https://developer.apple.com/documentation/applicationservices/atsudirectdataselector
type ATSUDirectDataSelector = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsuflattenstylerunoptions
type ATSUFlattenStyleRunOptions = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsuflatteneddatastreamformat
type ATSUFlattenedDataStreamFormat = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsufontfallbackmethod
type ATSUFontFallbackMethod = uint16

// See: https://developer.apple.com/documentation/applicationservices/atsufontfallbacks
type ATSUFontFallbacks = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsufontfeatureselector
type ATSUFontFeatureSelector = uint16

// See: https://developer.apple.com/documentation/applicationservices/atsufontfeaturetype
type ATSUFontFeatureType = uint16

// See: https://developer.apple.com/documentation/applicationservices/atsufontid
type ATSUFontID = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsufontvariationaxis
type ATSUFontVariationAxis = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsufontvariationvalue
type ATSUFontVariationValue = int32

// See: https://developer.apple.com/documentation/applicationservices/atsuhighlightmethod
type ATSUHighlightMethod = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsulayoutoperationcallbackstatus
type ATSULayoutOperationCallbackStatus = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsulayoutoperationselector
type ATSULayoutOperationSelector = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsulinetruncation
type ATSULineTruncation = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsustyle
type ATSUStyle = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsustylecomparison
type ATSUStyleComparison = uint16

// See: https://developer.apple.com/documentation/applicationservices/atsustylelinecounttype
type ATSUStyleLineCountType = uint16

// See: https://developer.apple.com/documentation/applicationservices/atsustylesettingref
type ATSUStyleSettingRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/atsutabtype
type ATSUTabType = uint16

// See: https://developer.apple.com/documentation/applicationservices/atsutextlayout
type ATSUTextLayout = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/atsutextmeasurement
type ATSUTextMeasurement = int32

// See: https://developer.apple.com/documentation/applicationservices/atsuunflattenstylerunoptions
type ATSUUnFlattenStyleRunOptions = uint32

// See: https://developer.apple.com/documentation/applicationservices/atsuverticalcharactertype
type ATSUVerticalCharacterType = uint16

// See: https://developer.apple.com/documentation/applicationservices/axobservercallback
type AXObserverCallback = func(observer AXObserverRef, element AXUIElementRef, notification corefoundation.CFStringRef, refcon unsafe.Pointer)

// See: https://developer.apple.com/documentation/applicationservices/axobservercallbackwithinfo
type AXObserverCallbackWithInfo = func(observer AXObserverRef, element AXUIElementRef, notification corefoundation.CFStringRef, info corefoundation.CFDictionaryRef, refcon unsafe.Pointer)

// See: https://developer.apple.com/documentation/applicationservices/axobserverref
type AXObserverRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/axtextmarkerrangeref
type AXTextMarkerRangeRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/axtextmarkerref
type AXTextMarkerRef uintptr

// AXUIElementRef is a structure used to refer to an accessibility object.
//
// See: https://developer.apple.com/documentation/applicationservices/axuielementref
type AXUIElementRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/axvalueref
type AXValueRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/appparametersptr
type AppParametersPtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/bitmaphandle
type BitMapHandle = *BitMapPtr

// See: https://developer.apple.com/documentation/applicationservices/bitmapptr
type BitMapPtr = *BitMap

// See: https://developer.apple.com/documentation/applicationservices/cgrafport
type CGrafPort = GrafPort

// See: https://developer.apple.com/documentation/applicationservices/cm2profileptr
type CM2ProfilePtr = *CM2Profile

// CMDeviceClass is define constants to represent a variety of input and output devices.
//
// See: https://developer.apple.com/documentation/applicationservices/cmdeviceclass
type CMDeviceClass = uint32

// CMFlattenProcPtr is defines a pointer to a data transfer callback function that transfers profile data from the format for embedded profiles to disk file format or vice versa.
//
// See: https://developer.apple.com/documentation/applicationservices/cmflattenprocptr
type CMFlattenProcPtr = func(command int32, size unsafe.Pointer, data unsafe.Pointer, refCon unsafe.Pointer) int16

// CMFlattenUPP is defines a universal procedure pointer to a data-flattening callback.
//
// See: https://developer.apple.com/documentation/applicationservices/cmflattenupp
type CMFlattenUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/cmmultifunctluta2btype
type CMMultiFunctLutA2BType = CMMultiFunctLutType

// See: https://developer.apple.com/documentation/applicationservices/cmxyzcomponent
type CMXYZComponent = uint16

// See: https://developer.apple.com/documentation/applicationservices/cqdprocsptr
type CQDProcsPtr = *CQDProcs

// See: https://developer.apple.com/documentation/applicationservices/cspecarray
type CSpecArray = ColorSpec

// See: https://developer.apple.com/documentation/applicationservices/colorcomplementprocptr
type ColorComplementProcPtr = func(unsafe.Pointer) bool

// See: https://developer.apple.com/documentation/applicationservices/colorcomplementupp
type ColorComplementUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/colorsearchprocptr
type ColorSearchProcPtr = func(unsafe.Pointer, unsafe.Pointer) bool

// See: https://developer.apple.com/documentation/applicationservices/colorsearchupp
type ColorSearchUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/colorspecptr
type ColorSpecPtr = *ColorSpec

// See: https://developer.apple.com/documentation/applicationservices/constatsuattributevalueptr
type ConstATSUAttributeValuePtr = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/dragconstraint
type DragConstraint = uint16

// See: https://developer.apple.com/documentation/applicationservices/draggrayrgnprocptr
type DragGrayRgnProcPtr = func()

// See: https://developer.apple.com/documentation/applicationservices/draggrayrgnupp
type DragGrayRgnUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/fmfilterselector
type FMFilterSelector = uint32

// See: https://developer.apple.com/documentation/applicationservices/fmfont
type FMFont = uint32

// See: https://developer.apple.com/documentation/applicationservices/fmfontfamily
type FMFontFamily = int16

// See: https://developer.apple.com/documentation/applicationservices/fmfontsize
type FMFontSize = int16

// See: https://developer.apple.com/documentation/applicationservices/fmfontstyle
type FMFontStyle = int16

// See: https://developer.apple.com/documentation/applicationservices/fmgeneration
type FMGeneration = uint32

// See: https://developer.apple.com/documentation/applicationservices/fontrechdl
type FontRecHdl = *FontRecPtr

// See: https://developer.apple.com/documentation/applicationservices/fontrecptr
type FontRecPtr = *FontRec

// See: https://developer.apple.com/documentation/applicationservices/glyphcollection
type GlyphCollection = uint16

// See: https://developer.apple.com/documentation/applicationservices/glyphid
type GlyphID = uint16

// See: https://developer.apple.com/documentation/applicationservices/grafverb
type GrafVerb = int8

// See: https://developer.apple.com/documentation/applicationservices/himutableshaperef
type HIMutableShapeRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/hishapeenumerateprocptr
type HIShapeEnumerateProcPtr = func(int32, HIShapeRef, unsafe.Pointer, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/hishaperef
type HIShapeRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/icappspechandle
type ICAppSpecHandle = *ICAppSpecPtr

// See: https://developer.apple.com/documentation/applicationservices/icappspeclisthandle
type ICAppSpecListHandle = *ICAppSpecListPtr

// See: https://developer.apple.com/documentation/applicationservices/icappspeclistptr
type ICAppSpecListPtr = *ICAppSpecList

// See: https://developer.apple.com/documentation/applicationservices/icappspecptr
type ICAppSpecPtr = *ICAppSpec

// See: https://developer.apple.com/documentation/applicationservices/icchartablehandle
type ICCharTableHandle = *ICCharTablePtr

// See: https://developer.apple.com/documentation/applicationservices/icchartableptr
type ICCharTablePtr = *ICCharTable

// See: https://developer.apple.com/documentation/applicationservices/icfilespechandle
type ICFileSpecHandle = *ICFileSpecPtr

// See: https://developer.apple.com/documentation/applicationservices/icfilespecptr
type ICFileSpecPtr = *ICFileSpec

// See: https://developer.apple.com/documentation/applicationservices/icfixedlength
type ICFixedLength = int32

// See: https://developer.apple.com/documentation/applicationservices/icfontrecordhandle
type ICFontRecordHandle = *ICFontRecordPtr

// See: https://developer.apple.com/documentation/applicationservices/icfontrecordptr
type ICFontRecordPtr = *ICFontRecord

// See: https://developer.apple.com/documentation/applicationservices/icmapentryflags
type ICMapEntryFlags = int32

// See: https://developer.apple.com/documentation/applicationservices/icmapentryhandle
type ICMapEntryHandle = *ICMapEntryPtr

// See: https://developer.apple.com/documentation/applicationservices/icmapentryptr
type ICMapEntryPtr = *ICMapEntry

// See: https://developer.apple.com/documentation/applicationservices/icprofileidptr
type ICProfileIDPtr = *ICProfileID

// See: https://developer.apple.com/documentation/applicationservices/icserviceentryflags
type ICServiceEntryFlags = int16

// See: https://developer.apple.com/documentation/applicationservices/icserviceentryhandle
type ICServiceEntryHandle = *ICServiceEntryPtr

// See: https://developer.apple.com/documentation/applicationservices/icserviceentryptr
type ICServiceEntryPtr = *ICServiceEntry

// See: https://developer.apple.com/documentation/applicationservices/icserviceshandle
type ICServicesHandle = *ICServicesPtr

// See: https://developer.apple.com/documentation/applicationservices/icservicesptr
type ICServicesPtr = *ICServices

// See: https://developer.apple.com/documentation/applicationservices/iconactionprocptr
type IconActionProcPtr = func(uint32, unsafe.Pointer, unsafe.Pointer) int16

// See: https://developer.apple.com/documentation/applicationservices/iconactionupp
type IconActionUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/iconalignmenttype
type IconAlignmentType = int16

// See: https://developer.apple.com/documentation/applicationservices/icongetterprocptr
type IconGetterProcPtr = func(uint32, unsafe.Pointer) kernel.Handle

// See: https://developer.apple.com/documentation/applicationservices/icongetterupp
type IconGetterUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/iconselectorvalue
type IconSelectorValue = uint32

// See: https://developer.apple.com/documentation/applicationservices/icontransformtype
type IconTransformType = int16

// See: https://developer.apple.com/documentation/applicationservices/launchflags
type LaunchFlags = uint16

// See: https://developer.apple.com/documentation/applicationservices/pmbordertype
type PMBorderType = uint16

// See: https://developer.apple.com/documentation/applicationservices/pmcolorspacemodel
type PMColorSpaceModel = uint32

// PMDestinationType is constants that specify a destination for a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/pmdestinationtype
type PMDestinationType = uint16

// PMDuplexMode is constants that specify duplex mode settings.
//
// See: https://developer.apple.com/documentation/applicationservices/pmduplexmode
type PMDuplexMode = uint32

// See: https://developer.apple.com/documentation/applicationservices/pmlayoutdirection
type PMLayoutDirection = uint16

// PMObject is the base type for all the opaque types used in Core Printing.
//
// See: https://developer.apple.com/documentation/applicationservices/pmobject
type PMObject = unsafe.Pointer

// PMOrientation is constants that specify page orientation.
//
// See: https://developer.apple.com/documentation/applicationservices/pmorientation
type PMOrientation = uint16

// PMPPDDomain is constants that specify the domains for PostScript printer description (PPD) files.
//
// See: https://developer.apple.com/documentation/applicationservices/pmppddomain
type PMPPDDomain = uint16

// PMPageFormat is an opaque type that stores the settings in the Page Setup dialog.
//
// See: https://developer.apple.com/documentation/applicationservices/pmpageformat
type PMPageFormat = uintptr

// PMPaper is an opaque type that stores information about the paper used in a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/pmpaper
type PMPaper = uintptr

// PMPaperMargins is a data structure that specifies the unprintable area of a paper object.
//
// See: https://developer.apple.com/documentation/applicationservices/pmpapermargins
type PMPaperMargins = PMRect

// See: https://developer.apple.com/documentation/applicationservices/pmpapertype
type PMPaperType = uint32

// PMPreset is an opaque type that stores information about a named preset available for a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/pmpreset
type PMPreset = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/pmprintdialogoptionflags
type PMPrintDialogOptionFlags = uint32

// PMPrintSession is an opaque type that stores information about a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/pmprintsession
type PMPrintSession = unsafe.Pointer

// PMPrintSettings is an opaque type that stores the settings in the Print dialog.
//
// See: https://developer.apple.com/documentation/applicationservices/pmprintsettings
type PMPrintSettings = uintptr

// PMPrinter is an opaque type that represents a printer.
//
// See: https://developer.apple.com/documentation/applicationservices/pmprinter
type PMPrinter = uintptr

// PMPrinterState is constants that specify the current state of a print queue.
//
// See: https://developer.apple.com/documentation/applicationservices/pmprinterstate
type PMPrinterState = uint16

// PMQualityMode is constants that specify standard options for print quality.
//
// See: https://developer.apple.com/documentation/applicationservices/pmqualitymode
type PMQualityMode = uint32

// See: https://developer.apple.com/documentation/applicationservices/pmscalingalignment
type PMScalingAlignment = uint16

// PMServer is an opaque type that identifies a local or remote print server.
//
// See: https://developer.apple.com/documentation/applicationservices/pmserver
type PMServer = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/pasteboarditemid
type PasteboardItemID = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/pasteboardpromisekeeperprocptr
type PasteboardPromiseKeeperProcPtr = func(PasteboardRef, PasteboardItemID, corefoundation.CFStringRef, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/pathandle
type PatHandle = *PatPtr

// See: https://developer.apple.com/documentation/applicationservices/patptr
type PatPtr = *Pattern

// See: https://developer.apple.com/documentation/applicationservices/pixpathandle
type PixPatHandle = *PixPatPtr

// See: https://developer.apple.com/documentation/applicationservices/pixpatptr
type PixPatPtr = *PixPat

// See: https://developer.apple.com/documentation/applicationservices/ploticonrefflags
type PlotIconRefFlags = uint32

// See: https://developer.apple.com/documentation/applicationservices/polyhandle
type PolyHandle = *PolyPtr

// See: https://developer.apple.com/documentation/applicationservices/polyptr
type PolyPtr = *MacPolygon

// See: https://developer.apple.com/documentation/applicationservices/polygon
type Polygon = MacPolygon

// See: https://developer.apple.com/documentation/applicationservices/printerstatusopcode
type PrinterStatusOpcode = int32

// See: https://developer.apple.com/documentation/applicationservices/processapplicationtransformstate
type ProcessApplicationTransformState = uint32

// See: https://developer.apple.com/documentation/applicationservices/processinfoextendedrecptr
type ProcessInfoExtendedRecPtr = *ProcessInfoExtendedRec

// See: https://developer.apple.com/documentation/applicationservices/processinforecptr
type ProcessInfoRecPtr = *ProcessInfoRec

// See: https://developer.apple.com/documentation/applicationservices/qdarcprocptr
type QDArcProcPtr = func(GrafVerb, unsafe.Pointer, int16, int16)

// See: https://developer.apple.com/documentation/applicationservices/qdarcupp
type QDArcUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdbitsprocptr
type QDBitsProcPtr = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int16, unsafe.Pointer)

// See: https://developer.apple.com/documentation/applicationservices/qdbitsupp
type QDBitsUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdcommentprocptr
type QDCommentProcPtr = func(int16, int16, kernel.Handle)

// See: https://developer.apple.com/documentation/applicationservices/qdcommentupp
type QDCommentUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qderr
type QDErr = int16

// See: https://developer.apple.com/documentation/applicationservices/qdgetpicprocptr
type QDGetPicProcPtr = func(unsafe.Pointer, int16)

// See: https://developer.apple.com/documentation/applicationservices/qdgetpicupp
type QDGetPicUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdjshieldcursorprocptr
type QDJShieldCursorProcPtr = func(int16, int16, int16, int16)

// See: https://developer.apple.com/documentation/applicationservices/qdlineprocptr
type QDLineProcPtr = func(Point)

// See: https://developer.apple.com/documentation/applicationservices/qdlineupp
type QDLineUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdopcodeprocptr
type QDOpcodeProcPtr = func(unsafe.Pointer, unsafe.Pointer, uint16, int16)

// See: https://developer.apple.com/documentation/applicationservices/qdopcodeupp
type QDOpcodeUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdovalprocptr
type QDOvalProcPtr = func(GrafVerb, unsafe.Pointer)

// See: https://developer.apple.com/documentation/applicationservices/qdovalupp
type QDOvalUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdpolyprocptr
type QDPolyProcPtr = func(GrafVerb, PolyHandle)

// See: https://developer.apple.com/documentation/applicationservices/qdpolyupp
type QDPolyUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdprinterstatusprocptr
type QDPrinterStatusProcPtr = func(PrinterStatusOpcode, unsafe.Pointer, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/qdprinterstatusupp
type QDPrinterStatusUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdputpicprocptr
type QDPutPicProcPtr = func(unsafe.Pointer, int16)

// See: https://developer.apple.com/documentation/applicationservices/qdputpicupp
type QDPutPicUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdrrectprocptr
type QDRRectProcPtr = func(GrafVerb, unsafe.Pointer, int16, int16)

// See: https://developer.apple.com/documentation/applicationservices/qdrrectupp
type QDRRectUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdrectprocptr
type QDRectProcPtr = func(GrafVerb, unsafe.Pointer)

// See: https://developer.apple.com/documentation/applicationservices/qdrectupp
type QDRectUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdregionparsedirection
type QDRegionParseDirection = int32

// See: https://developer.apple.com/documentation/applicationservices/qdrgnprocptr
type QDRgnProcPtr = func(GrafVerb, unsafe.Pointer)

// See: https://developer.apple.com/documentation/applicationservices/qdrgnupp
type QDRgnUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdstdglyphsprocptr
type QDStdGlyphsProcPtr = func(unsafe.Pointer, int32) int32

// See: https://developer.apple.com/documentation/applicationservices/qdstdglyphsupp
type QDStdGlyphsUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdtextprocptr
type QDTextProcPtr = func(int16, unsafe.Pointer, Point, Point)

// See: https://developer.apple.com/documentation/applicationservices/qdtextupp
type QDTextUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/qdtxmeasprocptr
type QDTxMeasProcPtr = func(int16, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int16

// See: https://developer.apple.com/documentation/applicationservices/qdtxmeasupp
type QDTxMeasUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/redrawbackgroundprocptr
type RedrawBackgroundProcPtr = func(unsafe.Pointer, coreservices.UniCharArrayOffset, int32, unsafe.Pointer, int32) bool

// See: https://developer.apple.com/documentation/applicationservices/redrawbackgroundupp
type RedrawBackgroundUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/regiontorectsprocptr
type RegionToRectsProcPtr = func(uint16, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/applicationservices/regiontorectsupp
type RegionToRectsUPP = unsafe.Pointer

// See: https://developer.apple.com/documentation/applicationservices/sizeresourcerechandle
type SizeResourceRecHandle = *SizeResourceRecPtr

// See: https://developer.apple.com/documentation/applicationservices/sizeresourcerecptr
type SizeResourceRecPtr = *SizeResourceRec

// SpeechChannel is defines a pointer to a speech channel record.
//
// See: https://developer.apple.com/documentation/applicationservices/speechchannel
type SpeechChannel = *SpeechChannelRecord

// SpeechErrorCFProcPtr is defines a pointer to an error callback function that handles syntax errors within commands embedded in a [CFString] object being processed by the Speech Synthesis Manager.
//
// See: https://developer.apple.com/documentation/applicationservices/speecherrorcfprocptr
type SpeechErrorCFProcPtr = func(chan_ *SpeechChannelRecord, refCon uintptr, theError corefoundation.CFErrorRef)

// SpeechWordCFProcPtr is defines a pointer to a Core Foundation-based word callback function that is called by the Speech Synthesis Manager before it pronounces a word.
//
// See: https://developer.apple.com/documentation/applicationservices/speechwordcfprocptr
type SpeechWordCFProcPtr = func(chan_ *SpeechChannelRecord, refCon uintptr, aString corefoundation.CFStringRef, wordRange corefoundation.CFRange)

// See: https://developer.apple.com/documentation/applicationservices/translationflags
type TranslationFlags = uint32

// See: https://developer.apple.com/documentation/applicationservices/translationref
type TranslationRef uintptr

// See: https://developer.apple.com/documentation/applicationservices/trunccode
type TruncCode = int16

// UAZoomChangeFocusType is defines the Universal Access zoom change focus type.
//
// See: https://developer.apple.com/documentation/applicationservices/uazoomchangefocustype
type UAZoomChangeFocusType = uint32

// VDGamRecPtr is represents a type used by the Video Components API.
//
// See: https://developer.apple.com/documentation/applicationservices/vdgamrecptr
type VDGamRecPtr = *VDGammaRecord

// See: https://developer.apple.com/documentation/applicationservices/voicespecptr
type VoiceSpecPtr = *VoiceSpec
