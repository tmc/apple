// Code generated from Apple documentation for coreservices. DO NOT EDIT.

package coreservices

import (
	"fmt"
)

type AbortErr int

const (
	AbortErrValue AbortErr = -27
	BdNamErr AbortErr = -37
	DceExtErr AbortErr = -30
	DirFulErr AbortErr = -33
	DskFulErr AbortErr = -34
	EofErr AbortErr = -39
	FLckdErr AbortErr = -45
	FnOpnErr AbortErr = -38
	FnfErr AbortErr = -43
	GcrOnMFMErr AbortErr = -400
	IIOAbortErr AbortErr = -27
	IoErr AbortErr = -36
	MFulErr AbortErr = -41
	NotOpenErr AbortErr = -28
	NsvErr AbortErr = -35
	PosErr AbortErr = -40
	SlotNumErr AbortErr = -360
	TmfoErr AbortErr = -42
	UnitTblFullErr AbortErr = -29
	WPrErr AbortErr = -44
)

func (e AbortErr) String() string {
	switch e {
	case AbortErrValue:
		return "AbortErrValue"
	case BdNamErr:
		return "BdNamErr"
	case DceExtErr:
		return "DceExtErr"
	case DirFulErr:
		return "DirFulErr"
	case DskFulErr:
		return "DskFulErr"
	case EofErr:
		return "EofErr"
	case FLckdErr:
		return "FLckdErr"
	case FnOpnErr:
		return "FnOpnErr"
	case FnfErr:
		return "FnfErr"
	case GcrOnMFMErr:
		return "GcrOnMFMErr"
	case IoErr:
		return "IoErr"
	case MFulErr:
		return "MFulErr"
	case NotOpenErr:
		return "NotOpenErr"
	case NsvErr:
		return "NsvErr"
	case PosErr:
		return "PosErr"
	case SlotNumErr:
		return "SlotNumErr"
	case TmfoErr:
		return "TmfoErr"
	case UnitTblFullErr:
		return "UnitTblFullErr"
	case WPrErr:
		return "WPrErr"
	default:
		return fmt.Sprintf("AbortErr(%d)", e)
	}
}

type AeBuildSyntax uint

const (
	// AeBuildSyntaxBadData: Bad data was found inside a variable argument list.
	AeBuildSyntaxBadData AeBuildSyntax = 12
	// AeBuildSyntaxBadDesc: An illegal descriptor was specified.
	AeBuildSyntaxBadDesc AeBuildSyntax = 11
	// AeBuildSyntaxBadEOF: An unexpected end of format string was encountered.
	AeBuildSyntaxBadEOF AeBuildSyntax = 2
	// AeBuildSyntaxBadHex: A hex string contained characters other than hexadecimal digits.
	AeBuildSyntaxBadHex AeBuildSyntax = 6
	// AeBuildSyntaxBadNegative: A minus sign “-” was not followed by digits.
	AeBuildSyntaxBadNegative AeBuildSyntax = 4
	// AeBuildSyntaxBadToken: An illegal character was specified.
	AeBuildSyntaxBadToken AeBuildSyntax = 1
	// AeBuildSyntaxCoercedList: Cannot coerce a list.
	AeBuildSyntaxCoercedList AeBuildSyntax = 18
	// AeBuildSyntaxMissingQuote: A string was not terminated by a closing quotation mark.
	AeBuildSyntaxMissingQuote AeBuildSyntax = 5
	// AeBuildSyntaxNoCloseBrace: A comma or closing brace “}” was expected.
	AeBuildSyntaxNoCloseBrace AeBuildSyntax = 15
	// AeBuildSyntaxNoCloseBracket: A comma or closing bracket “]” was expected.
	AeBuildSyntaxNoCloseBracket AeBuildSyntax = 14
	// AeBuildSyntaxNoCloseHex: A hex string was missing a “$” or “»” character.
	AeBuildSyntaxNoCloseHex AeBuildSyntax = 8
	// AeBuildSyntaxNoCloseParen: A data value was missing a closing parenthesis.
	AeBuildSyntaxNoCloseParen AeBuildSyntax = 13
	// AeBuildSyntaxNoCloseString: A string was missing a closing quote.
	AeBuildSyntaxNoCloseString AeBuildSyntax = 10
	// AeBuildSyntaxNoColon: In a descriptor, one of the keywords was not followed by a colon.
	AeBuildSyntaxNoColon AeBuildSyntax = 17
	// AeBuildSyntaxNoEOF: There were unexpected characters beyond the end of the format string.
	AeBuildSyntaxNoEOF AeBuildSyntax = 3
	// AeBuildSyntaxNoErr: No error.
	AeBuildSyntaxNoErr AeBuildSyntax = 0
	// AeBuildSyntaxNoKey: A keyword was missing from a descriptor.
	AeBuildSyntaxNoKey AeBuildSyntax = 16
	// AeBuildSyntaxOddHex: A hex string contained an odd number of digits.
	AeBuildSyntaxOddHex AeBuildSyntax = 7
	// AeBuildSyntaxUncoercedDoubleAt: You must coerce a “@@” substitution.
	AeBuildSyntaxUncoercedDoubleAt AeBuildSyntax = 19
	// AeBuildSyntaxUncoercedHex: A hex string must be coerced to a type.
	AeBuildSyntaxUncoercedHex AeBuildSyntax = 9
)

func (e AeBuildSyntax) String() string {
	switch e {
	case AeBuildSyntaxBadData:
		return "AeBuildSyntaxBadData"
	case AeBuildSyntaxBadDesc:
		return "AeBuildSyntaxBadDesc"
	case AeBuildSyntaxBadEOF:
		return "AeBuildSyntaxBadEOF"
	case AeBuildSyntaxBadHex:
		return "AeBuildSyntaxBadHex"
	case AeBuildSyntaxBadNegative:
		return "AeBuildSyntaxBadNegative"
	case AeBuildSyntaxBadToken:
		return "AeBuildSyntaxBadToken"
	case AeBuildSyntaxCoercedList:
		return "AeBuildSyntaxCoercedList"
	case AeBuildSyntaxMissingQuote:
		return "AeBuildSyntaxMissingQuote"
	case AeBuildSyntaxNoCloseBrace:
		return "AeBuildSyntaxNoCloseBrace"
	case AeBuildSyntaxNoCloseBracket:
		return "AeBuildSyntaxNoCloseBracket"
	case AeBuildSyntaxNoCloseHex:
		return "AeBuildSyntaxNoCloseHex"
	case AeBuildSyntaxNoCloseParen:
		return "AeBuildSyntaxNoCloseParen"
	case AeBuildSyntaxNoCloseString:
		return "AeBuildSyntaxNoCloseString"
	case AeBuildSyntaxNoColon:
		return "AeBuildSyntaxNoColon"
	case AeBuildSyntaxNoEOF:
		return "AeBuildSyntaxNoEOF"
	case AeBuildSyntaxNoErr:
		return "AeBuildSyntaxNoErr"
	case AeBuildSyntaxNoKey:
		return "AeBuildSyntaxNoKey"
	case AeBuildSyntaxOddHex:
		return "AeBuildSyntaxOddHex"
	case AeBuildSyntaxUncoercedDoubleAt:
		return "AeBuildSyntaxUncoercedDoubleAt"
	case AeBuildSyntaxUncoercedHex:
		return "AeBuildSyntaxUncoercedHex"
	default:
		return fmt.Sprintf("AeBuildSyntax(%d)", e)
	}
}

type Afp int

const (
	AfpAccessDenied Afp = -5000
	AfpAlreadyLoggedInErr Afp = -5047
	AfpAlreadyMounted Afp = -5062
	AfpAuthContinue Afp = -5001
	AfpBadDirIDType Afp = -5060
	AfpBadIDErr Afp = -5039
	AfpBadUAM Afp = -5002
	AfpBadVersNum Afp = -5003
	AfpBitmapErr Afp = -5004
	AfpCallNotAllowed Afp = -5048
	AfpCallNotSupported Afp = -5024
	AfpCantMountMoreSrvre Afp = -5061
	AfpCantMove Afp = -5005
	AfpCantRename Afp = -5028
	AfpCatalogChanged Afp = -5037
	AfpContainsSharedErr Afp = -5033
	AfpDenyConflict Afp = -5006
	AfpDiffVolErr Afp = -5036
	AfpDirNotEmpty Afp = -5007
	AfpDirNotFound Afp = -5029
	AfpDiskFull Afp = -5008
	AfpEofError Afp = -5009
	AfpFileBusy Afp = -5010
	AfpFlatVol Afp = -5011
	AfpIDExists Afp = -5035
	AfpIDNotFound Afp = -5034
	AfpIconTypeError Afp = -5030
	AfpInsideSharedErr Afp = -5043
	AfpInsideTrashErr Afp = -5044
	AfpItemNotFound Afp = -5012
	AfpLockErr Afp = -5013
	AfpMiscErr Afp = -5014
	AfpNoMoreLocks Afp = -5015
	AfpNoServer Afp = -5016
	AfpObjectExists Afp = -5017
	AfpObjectLocked Afp = -5032
	AfpObjectNotFound Afp = -5018
	AfpObjectTypeErr Afp = -5025
	AfpParmErr Afp = -5019
	AfpPwdExpiredErr Afp = -5042
	AfpPwdNeedsChangeErr Afp = -5045
	AfpPwdPolicyErr Afp = -5046
	AfpPwdSameErr Afp = -5040
	AfpPwdTooShortErr Afp = -5041
	AfpRangeNotLocked Afp = -5020
	AfpRangeOverlap Afp = -5021
	AfpSameNodeErr Afp = -5063
	AfpSameObjectErr Afp = -5038
	AfpServerGoingDown Afp = -5027
	AfpSessClosed Afp = -5022
	AfpTooManyFilesOpen Afp = -5026
	AfpUserNotAuth Afp = -5023
	AfpVolLocked Afp = -5031
)

func (e Afp) String() string {
	switch e {
	case AfpAccessDenied:
		return "AfpAccessDenied"
	case AfpAlreadyLoggedInErr:
		return "AfpAlreadyLoggedInErr"
	case AfpAlreadyMounted:
		return "AfpAlreadyMounted"
	case AfpAuthContinue:
		return "AfpAuthContinue"
	case AfpBadDirIDType:
		return "AfpBadDirIDType"
	case AfpBadIDErr:
		return "AfpBadIDErr"
	case AfpBadUAM:
		return "AfpBadUAM"
	case AfpBadVersNum:
		return "AfpBadVersNum"
	case AfpBitmapErr:
		return "AfpBitmapErr"
	case AfpCallNotAllowed:
		return "AfpCallNotAllowed"
	case AfpCallNotSupported:
		return "AfpCallNotSupported"
	case AfpCantMountMoreSrvre:
		return "AfpCantMountMoreSrvre"
	case AfpCantMove:
		return "AfpCantMove"
	case AfpCantRename:
		return "AfpCantRename"
	case AfpCatalogChanged:
		return "AfpCatalogChanged"
	case AfpContainsSharedErr:
		return "AfpContainsSharedErr"
	case AfpDenyConflict:
		return "AfpDenyConflict"
	case AfpDiffVolErr:
		return "AfpDiffVolErr"
	case AfpDirNotEmpty:
		return "AfpDirNotEmpty"
	case AfpDirNotFound:
		return "AfpDirNotFound"
	case AfpDiskFull:
		return "AfpDiskFull"
	case AfpEofError:
		return "AfpEofError"
	case AfpFileBusy:
		return "AfpFileBusy"
	case AfpFlatVol:
		return "AfpFlatVol"
	case AfpIDExists:
		return "AfpIDExists"
	case AfpIDNotFound:
		return "AfpIDNotFound"
	case AfpIconTypeError:
		return "AfpIconTypeError"
	case AfpInsideSharedErr:
		return "AfpInsideSharedErr"
	case AfpInsideTrashErr:
		return "AfpInsideTrashErr"
	case AfpItemNotFound:
		return "AfpItemNotFound"
	case AfpLockErr:
		return "AfpLockErr"
	case AfpMiscErr:
		return "AfpMiscErr"
	case AfpNoMoreLocks:
		return "AfpNoMoreLocks"
	case AfpNoServer:
		return "AfpNoServer"
	case AfpObjectExists:
		return "AfpObjectExists"
	case AfpObjectLocked:
		return "AfpObjectLocked"
	case AfpObjectNotFound:
		return "AfpObjectNotFound"
	case AfpObjectTypeErr:
		return "AfpObjectTypeErr"
	case AfpParmErr:
		return "AfpParmErr"
	case AfpPwdExpiredErr:
		return "AfpPwdExpiredErr"
	case AfpPwdNeedsChangeErr:
		return "AfpPwdNeedsChangeErr"
	case AfpPwdPolicyErr:
		return "AfpPwdPolicyErr"
	case AfpPwdSameErr:
		return "AfpPwdSameErr"
	case AfpPwdTooShortErr:
		return "AfpPwdTooShortErr"
	case AfpRangeNotLocked:
		return "AfpRangeNotLocked"
	case AfpRangeOverlap:
		return "AfpRangeOverlap"
	case AfpSameNodeErr:
		return "AfpSameNodeErr"
	case AfpSameObjectErr:
		return "AfpSameObjectErr"
	case AfpServerGoingDown:
		return "AfpServerGoingDown"
	case AfpSessClosed:
		return "AfpSessClosed"
	case AfpTooManyFilesOpen:
		return "AfpTooManyFilesOpen"
	case AfpUserNotAuth:
		return "AfpUserNotAuth"
	case AfpVolLocked:
		return "AfpVolLocked"
	default:
		return fmt.Sprintf("Afp(%d)", e)
	}
}

type Asp int

const (
	AspBadVersNum Asp = -1066
	AspBufTooSmall Asp = -1067
	AspNoAck Asp = -1075
	AspNoMoreSess Asp = -1068
	AspNoServers Asp = -1069
	AspParamErr Asp = -1070
	AspServerBusy Asp = -1071
	AspSessClosed Asp = -1072
	AspSizeErr Asp = -1073
	AspTooMany Asp = -1074
)

func (e Asp) String() string {
	switch e {
	case AspBadVersNum:
		return "AspBadVersNum"
	case AspBufTooSmall:
		return "AspBufTooSmall"
	case AspNoAck:
		return "AspNoAck"
	case AspNoMoreSess:
		return "AspNoMoreSess"
	case AspNoServers:
		return "AspNoServers"
	case AspParamErr:
		return "AspParamErr"
	case AspServerBusy:
		return "AspServerBusy"
	case AspSessClosed:
		return "AspSessClosed"
	case AspSizeErr:
		return "AspSizeErr"
	case AspTooMany:
		return "AspTooMany"
	default:
		return fmt.Sprintf("Asp(%d)", e)
	}
}

type BadComponent uint

const (
	// BadComponentInstance: Invalid component  passed to Component Manager.
	BadComponentInstance BadComponent = 0
	// BadComponentSelector: Component does not support the specified request code.
	BadComponentSelector BadComponent = 0
)

func (e BadComponent) String() string {
	switch e {
	case BadComponentInstance:
		return "BadComponentInstance"
	default:
		return fmt.Sprintf("BadComponent(%d)", e)
	}
}

type BadDragRefErr int

const (
	BadDragFlavorErr BadDragRefErr = -1852
	BadDragItemErr BadDragRefErr = -1851
	BadDragRefErrValue BadDragRefErr = -1850
	BadImageErr BadDragRefErr = -1861
	BadImageRgnErr BadDragRefErr = -1860
	CantGetFlavorErr BadDragRefErr = -1854
	DragNotAcceptedErr BadDragRefErr = -1857
	DuplicateFlavorErr BadDragRefErr = -1853
	DuplicateHandlerErr BadDragRefErr = -1855
	HandlerNotFoundErr BadDragRefErr = -1856
	NoSuitableDisplaysErr BadDragRefErr = -1859
	NonDragOriginatorErr BadDragRefErr = -1862
	UnsupportedForPlatformErr BadDragRefErr = -1858
)

func (e BadDragRefErr) String() string {
	switch e {
	case BadDragFlavorErr:
		return "BadDragFlavorErr"
	case BadDragItemErr:
		return "BadDragItemErr"
	case BadDragRefErrValue:
		return "BadDragRefErrValue"
	case BadImageErr:
		return "BadImageErr"
	case BadImageRgnErr:
		return "BadImageRgnErr"
	case CantGetFlavorErr:
		return "CantGetFlavorErr"
	case DragNotAcceptedErr:
		return "DragNotAcceptedErr"
	case DuplicateFlavorErr:
		return "DuplicateFlavorErr"
	case DuplicateHandlerErr:
		return "DuplicateHandlerErr"
	case HandlerNotFoundErr:
		return "HandlerNotFoundErr"
	case NoSuitableDisplaysErr:
		return "NoSuitableDisplaysErr"
	case NonDragOriginatorErr:
		return "NonDragOriginatorErr"
	case UnsupportedForPlatformErr:
		return "UnsupportedForPlatformErr"
	default:
		return fmt.Sprintf("BadDragRefErr(%d)", e)
	}
}

type BadFolderDescErr int

const (
	BadFolderDescErrValue BadFolderDescErr = -4270
	BadRoutingSizeErr BadFolderDescErr = -4276
	DuplicateFolderDescErr BadFolderDescErr = -4271
	DuplicateRoutingErr BadFolderDescErr = -4274
	InvalidFolderTypeErr BadFolderDescErr = -4273
	NoMoreFolderDescErr BadFolderDescErr = -4272
	RoutingNotFoundErr BadFolderDescErr = -4275
)

func (e BadFolderDescErr) String() string {
	switch e {
	case BadFolderDescErrValue:
		return "BadFolderDescErrValue"
	case BadRoutingSizeErr:
		return "BadRoutingSizeErr"
	case DuplicateFolderDescErr:
		return "DuplicateFolderDescErr"
	case DuplicateRoutingErr:
		return "DuplicateRoutingErr"
	case InvalidFolderTypeErr:
		return "InvalidFolderTypeErr"
	case NoMoreFolderDescErr:
		return "NoMoreFolderDescErr"
	case RoutingNotFoundErr:
		return "RoutingNotFoundErr"
	default:
		return fmt.Sprintf("BadFolderDescErr(%d)", e)
	}
}

type Buf2SmallErr int

const (
	AtpBadRsp Buf2SmallErr = -3107
	AtpLenErr Buf2SmallErr = -3106
	Buf2SmallErrValue Buf2SmallErr = -3101
	CkSumErr Buf2SmallErr = -3103
	ExtractErr Buf2SmallErr = -3104
	NoMPPErr Buf2SmallErr = -3102
	ReadQErr Buf2SmallErr = -3105
	RecNotFnd Buf2SmallErr = -3108
	SktClosedErr Buf2SmallErr = -3109
)

func (e Buf2SmallErr) String() string {
	switch e {
	case AtpBadRsp:
		return "AtpBadRsp"
	case AtpLenErr:
		return "AtpLenErr"
	case Buf2SmallErrValue:
		return "Buf2SmallErrValue"
	case CkSumErr:
		return "CkSumErr"
	case ExtractErr:
		return "ExtractErr"
	case NoMPPErr:
		return "NoMPPErr"
	case ReadQErr:
		return "ReadQErr"
	case RecNotFnd:
		return "RecNotFnd"
	case SktClosedErr:
		return "SktClosedErr"
	default:
		return fmt.Sprintf("Buf2SmallErr(%d)", e)
	}
}

type CDEFNFndConstants int

const (
	CDEFNFnd CDEFNFndConstants = 88
	DsBadStartupDisk CDEFNFndConstants = 42
	DsHD20Installed CDEFNFndConstants = -12
	DsNotThe1 CDEFNFndConstants = 31
	DsSystemFileErr CDEFNFndConstants = 43
	ExUserBreak CDEFNFndConstants = -492
	FsDSIntErr CDEFNFndConstants = -127
	HMenuFindErr CDEFNFndConstants = -127
	MBarNFnd CDEFNFndConstants = -126
	StrUserBreak CDEFNFndConstants = -491
	UserBreak CDEFNFndConstants = -490
)

func (e CDEFNFndConstants) String() string {
	switch e {
	case CDEFNFnd:
		return "CDEFNFnd"
	case DsBadStartupDisk:
		return "DsBadStartupDisk"
	case DsHD20Installed:
		return "DsHD20Installed"
	case DsNotThe1:
		return "DsNotThe1"
	case DsSystemFileErr:
		return "DsSystemFileErr"
	case ExUserBreak:
		return "ExUserBreak"
	case FsDSIntErr:
		return "FsDSIntErr"
	case MBarNFnd:
		return "MBarNFnd"
	case StrUserBreak:
		return "StrUserBreak"
	case UserBreak:
		return "UserBreak"
	default:
		return fmt.Sprintf("CDEFNFndConstants(%d)", e)
	}
}

type CMatchErr int

const (
	CDepthErr CMatchErr = -157
	CDevErr CMatchErr = -155
	CMatchErrValue CMatchErr = -150
	CNoMemErr CMatchErr = -152
	CProtectErr CMatchErr = -154
	CRangeErr CMatchErr = -153
	CResErr CMatchErr = -156
	CTempMemErr CMatchErr = -151
	CantLoadPickMethodErr CMatchErr = -11003
	ColorsRequestedErr CMatchErr = -11004
	PictInfoIDErr CMatchErr = -11001
	PictInfoVerbErr CMatchErr = -11002
	PictInfoVersionErr CMatchErr = -11000
	PictureDataErr CMatchErr = -11005
	RgnTooBigErr CMatchErr = -500
	UpdPixMemErr CMatchErr = -125
)

func (e CMatchErr) String() string {
	switch e {
	case CDepthErr:
		return "CDepthErr"
	case CDevErr:
		return "CDevErr"
	case CMatchErrValue:
		return "CMatchErrValue"
	case CNoMemErr:
		return "CNoMemErr"
	case CProtectErr:
		return "CProtectErr"
	case CRangeErr:
		return "CRangeErr"
	case CResErr:
		return "CResErr"
	case CTempMemErr:
		return "CTempMemErr"
	case CantLoadPickMethodErr:
		return "CantLoadPickMethodErr"
	case ColorsRequestedErr:
		return "ColorsRequestedErr"
	case PictInfoIDErr:
		return "PictInfoIDErr"
	case PictInfoVerbErr:
		return "PictInfoVerbErr"
	case PictInfoVersionErr:
		return "PictInfoVersionErr"
	case PictureDataErr:
		return "PictureDataErr"
	case RgnTooBigErr:
		return "RgnTooBigErr"
	case UpdPixMemErr:
		return "UpdPixMemErr"
	default:
		return fmt.Sprintf("CMatchErr(%d)", e)
	}
}

type CParagraph uint

const (
	// CPICT: A PICT format figure.
	CPICT CParagraph = 'P'<<24 | 'I'<<16 | 'C'<<8 | 'T' // 'PICT'
	// CParagraphValue: A paragraph of text.
	CParagraphValue CParagraph = 'c'<<24 | 'p'<<16 | 'a'<<8 | 'r' // 'cpar'
	CPixel CParagraph = 'c'<<24 | 'p'<<16 | 'x'<<8 | 'l' // 'cpxl'
	CPixelMap CParagraph = 'c'<<24 | 'p'<<16 | 'i'<<8 | 'x' // 'cpix'
	CPolygon CParagraph = 'c'<<24 | 'p'<<16 | 'g'<<8 | 'n' // 'cpgn'
	// CProperty: A property of any object class.
	CProperty CParagraph = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
	CQDPoint CParagraph = 'Q'<<24 | 'D'<<16 | 'p'<<8 | 't' // 'QDpt'
	CQDRectangle CParagraph = 'q'<<24 | 'd'<<16 | 'r'<<8 | 't' // 'qdrt'
	// CRGBColor: An RGB color value.
	CRGBColor CParagraph = 'c'<<24 | 'R'<<16 | 'G'<<8 | 'B' // 'cRGB'
	CRectangle CParagraph = 'c'<<24 | 'r'<<16 | 'e'<<8 | 'c' // 'crec'
	CRotation CParagraph = 't'<<24 | 'r'<<16 | 'o'<<8 | 't' // 'trot'
	CRoundedRectangle CParagraph = 'c'<<24 | 'r'<<16 | 'r'<<8 | 'c' // 'crrc'
	CRow CParagraph = 'c'<<24 | 'r'<<16 | 'o'<<8 | 'w' // 'crow'
	CSelection CParagraph = 'c'<<24 | 's'<<16 | 'e'<<8 | 'l' // 'csel'
	CShortInteger CParagraph = 's'<<24 | 'h'<<16 | 'o'<<8 | 'r' // 'shor'
	CTable CParagraph = 'c'<<24 | 't'<<16 | 'b'<<8 | 'l' // 'ctbl'
	CText CParagraph = 'c'<<24 | 't'<<16 | 'x'<<8 | 't' // 'ctxt'
	CTextFlow CParagraph = 'c'<<24 | 'f'<<16 | 'l'<<8 | 'o' // 'cflo'
	CTextStyles CParagraph = 't'<<24 | 's'<<16 | 't'<<8 | 'y' // 'tsty'
	CType CParagraph = 't'<<24 | 'y'<<16 | 'p'<<8 | 'e' // 'type'
)

func (e CParagraph) String() string {
	switch e {
	case CPICT:
		return "CPICT"
	case CParagraphValue:
		return "CParagraphValue"
	case CPixel:
		return "CPixel"
	case CPixelMap:
		return "CPixelMap"
	case CPolygon:
		return "CPolygon"
	case CProperty:
		return "CProperty"
	case CQDPoint:
		return "CQDPoint"
	case CQDRectangle:
		return "CQDRectangle"
	case CRGBColor:
		return "CRGBColor"
	case CRectangle:
		return "CRectangle"
	case CRotation:
		return "CRotation"
	case CRoundedRectangle:
		return "CRoundedRectangle"
	case CRow:
		return "CRow"
	case CSelection:
		return "CSelection"
	case CShortInteger:
		return "CShortInteger"
	case CTable:
		return "CTable"
	case CText:
		return "CText"
	case CTextFlow:
		return "CTextFlow"
	case CTextStyles:
		return "CTextStyles"
	case CType:
		return "CType"
	default:
		return fmt.Sprintf("CParagraph(%d)", e)
	}
}

type CVersion uint

const (
	CVersionValue CVersion = 'v'<<24 | 'e'<<16 | 'r'<<8 | 's' // 'vers'
	CWindow CVersion = 'c'<<24 | 'w'<<16 | 'i'<<8 | 'n' // 'cwin'
	CWord CVersion = 'c'<<24 | 'w'<<16 | 'o'<<8 | 'r' // 'cwor'
	EnumArrows CVersion = 'a'<<24 | 'r'<<16 | 'r'<<8 | 'o' // 'arro'
	EnumJustification CVersion = 'j'<<24 | 'u'<<16 | 's'<<8 | 't' // 'just'
	EnumKeyForm CVersion = 'k'<<24 | 'f'<<16 | 'r'<<8 | 'm' // 'kfrm'
	EnumPosition CVersion = 'p'<<24 | 'o'<<16 | 's'<<8 | 'i' // 'posi'
	EnumProtection CVersion = 'p'<<24 | 'r'<<16 | 't'<<8 | 'n' // 'prtn'
	EnumQuality CVersion = 'q'<<24 | 'u'<<16 | 'a'<<8 | 'l' // 'qual'
	EnumSaveOptions CVersion = 's'<<24 | 'a'<<16 | 'v'<<8 | 'o' // 'savo'
	EnumStyle CVersion = 's'<<24 | 't'<<16 | 'y'<<8 | 'l' // 'styl'
	EnumTransferMode CVersion = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
	KAEAbout CVersion = 'a'<<24 | 'b'<<16 | 'o'<<8 | 'u' // 'abou'
	KAEAfter CVersion = 'a'<<24 | 'f'<<16 | 't'<<8 | 'e' // 'afte'
	KAEAliasSelection CVersion = 's'<<24 | 'a'<<16 | 'l'<<8 | 'i' // 'sali'
	KAEAllCaps CVersion = 'a'<<24 | 'l'<<16 | 'c'<<8 | 'p' // 'alcp'
	KAEArrowAtEnd CVersion = 'a'<<24 | 'r'<<16 | 'e'<<8 | 'n' // 'aren'
	KAEArrowAtStart CVersion = 'a'<<24 | 'r'<<16 | 's'<<8 | 't' // 'arst'
	KAEArrowBothEnds CVersion = 'a'<<24 | 'r'<<16 | 'b'<<8 | 'o' // 'arbo'
)

func (e CVersion) String() string {
	switch e {
	case CVersionValue:
		return "CVersionValue"
	case CWindow:
		return "CWindow"
	case CWord:
		return "CWord"
	case EnumArrows:
		return "EnumArrows"
	case EnumJustification:
		return "EnumJustification"
	case EnumKeyForm:
		return "EnumKeyForm"
	case EnumPosition:
		return "EnumPosition"
	case EnumProtection:
		return "EnumProtection"
	case EnumQuality:
		return "EnumQuality"
	case EnumSaveOptions:
		return "EnumSaveOptions"
	case EnumStyle:
		return "EnumStyle"
	case EnumTransferMode:
		return "EnumTransferMode"
	case KAEAbout:
		return "KAEAbout"
	case KAEAfter:
		return "KAEAfter"
	case KAEAliasSelection:
		return "KAEAliasSelection"
	case KAEAllCaps:
		return "KAEAllCaps"
	case KAEArrowAtEnd:
		return "KAEArrowAtEnd"
	case KAEArrowAtStart:
		return "KAEArrowAtStart"
	case KAEArrowBothEnds:
		return "KAEArrowBothEnds"
	default:
		return fmt.Sprintf("CVersion(%d)", e)
	}
}

type CannotFindAtomErr int

const (
	AAPNotCreatedErr CannotFindAtomErr = -2120
	AAPNotFoundErr CannotFindAtomErr = -2121
	ASDBadForkErr CannotFindAtomErr = -2123
	ASDBadHeaderErr CannotFindAtomErr = -2122
	ASDEntryNotFoundErr CannotFindAtomErr = -2124
	AtomIndexInvalidErr CannotFindAtomErr = -2104
	AtomsNotOfSameTypeErr CannotFindAtomErr = -2103
	CannotBeLeafAtomErr CannotFindAtomErr = -2109
	CannotFindAtomErrValue CannotFindAtomErr = -2101
	DuplicateAtomTypeAndIDErr CannotFindAtomErr = -2105
	EmptyPathErr CannotFindAtomErr = -2111
	FileOffsetTooBigErr CannotFindAtomErr = -2125
	InvalidAtomContainerErr CannotFindAtomErr = -2107
	InvalidAtomErr CannotFindAtomErr = -2106
	InvalidAtomTypeErr CannotFindAtomErr = -2108
	NoPathMappingErr CannotFindAtomErr = -2112
	NotAllowedToSaveMovieErr CannotFindAtomErr = -2126
	NotEnoughDataErr CannotFindAtomErr = -2149
	NotLeafAtomErr CannotFindAtomErr = -2102
	PathNotVerifiedErr CannotFindAtomErr = -2113
	PathTooLongErr CannotFindAtomErr = -2110
	QfcbNotCreatedErr CannotFindAtomErr = -2119
	QfcbNotFoundErr CannotFindAtomErr = -2118
	QtActionNotHandledErr CannotFindAtomErr = -2157
	QtNetworkAlreadyAllocatedErr CannotFindAtomErr = -2127
	QtXMLApplicationErr CannotFindAtomErr = -2159
	QtXMLParseErr CannotFindAtomErr = -2158
	UnknownFormatErr CannotFindAtomErr = -2114
	UrlDataHFTPBadNameListErr CannotFindAtomErr = -2144
	UrlDataHFTPBadPasswordErr CannotFindAtomErr = -2136
	UrlDataHFTPBadUserErr CannotFindAtomErr = -2135
	UrlDataHFTPDataConnectionErr CannotFindAtomErr = -2138
	UrlDataHFTPFilenameErr CannotFindAtomErr = -2142
	UrlDataHFTPNeedPasswordErr CannotFindAtomErr = -2145
	UrlDataHFTPNoDirectoryErr CannotFindAtomErr = -2139
	UrlDataHFTPNoNetDriverErr CannotFindAtomErr = -2143
	UrlDataHFTPNoPasswordErr CannotFindAtomErr = -2146
	UrlDataHFTPPermissionsErr CannotFindAtomErr = -2141
	UrlDataHFTPProtocolErr CannotFindAtomErr = -2133
	UrlDataHFTPQuotaErr CannotFindAtomErr = -2140
	UrlDataHFTPServerDisconnectedErr CannotFindAtomErr = -2147
	UrlDataHFTPServerErr CannotFindAtomErr = -2137
	UrlDataHFTPShutdownErr CannotFindAtomErr = -2134
	UrlDataHFTPURLErr CannotFindAtomErr = -2148
	UrlDataHHTTPNoNetDriverErr CannotFindAtomErr = -2130
	UrlDataHHTTPProtocolErr CannotFindAtomErr = -2129
	UrlDataHHTTPRedirectErr CannotFindAtomErr = -2132
	UrlDataHHTTPURLErr CannotFindAtomErr = -2131
	WackBadFileErr CannotFindAtomErr = -2115
	WackBadMetaDataErr CannotFindAtomErr = -2117
	WackForkNotFoundErr CannotFindAtomErr = -2116
)

func (e CannotFindAtomErr) String() string {
	switch e {
	case AAPNotCreatedErr:
		return "AAPNotCreatedErr"
	case AAPNotFoundErr:
		return "AAPNotFoundErr"
	case ASDBadForkErr:
		return "ASDBadForkErr"
	case ASDBadHeaderErr:
		return "ASDBadHeaderErr"
	case ASDEntryNotFoundErr:
		return "ASDEntryNotFoundErr"
	case AtomIndexInvalidErr:
		return "AtomIndexInvalidErr"
	case AtomsNotOfSameTypeErr:
		return "AtomsNotOfSameTypeErr"
	case CannotBeLeafAtomErr:
		return "CannotBeLeafAtomErr"
	case CannotFindAtomErrValue:
		return "CannotFindAtomErrValue"
	case DuplicateAtomTypeAndIDErr:
		return "DuplicateAtomTypeAndIDErr"
	case EmptyPathErr:
		return "EmptyPathErr"
	case FileOffsetTooBigErr:
		return "FileOffsetTooBigErr"
	case InvalidAtomContainerErr:
		return "InvalidAtomContainerErr"
	case InvalidAtomErr:
		return "InvalidAtomErr"
	case InvalidAtomTypeErr:
		return "InvalidAtomTypeErr"
	case NoPathMappingErr:
		return "NoPathMappingErr"
	case NotAllowedToSaveMovieErr:
		return "NotAllowedToSaveMovieErr"
	case NotEnoughDataErr:
		return "NotEnoughDataErr"
	case NotLeafAtomErr:
		return "NotLeafAtomErr"
	case PathNotVerifiedErr:
		return "PathNotVerifiedErr"
	case PathTooLongErr:
		return "PathTooLongErr"
	case QfcbNotCreatedErr:
		return "QfcbNotCreatedErr"
	case QfcbNotFoundErr:
		return "QfcbNotFoundErr"
	case QtActionNotHandledErr:
		return "QtActionNotHandledErr"
	case QtNetworkAlreadyAllocatedErr:
		return "QtNetworkAlreadyAllocatedErr"
	case QtXMLApplicationErr:
		return "QtXMLApplicationErr"
	case QtXMLParseErr:
		return "QtXMLParseErr"
	case UnknownFormatErr:
		return "UnknownFormatErr"
	case UrlDataHFTPBadNameListErr:
		return "UrlDataHFTPBadNameListErr"
	case UrlDataHFTPBadPasswordErr:
		return "UrlDataHFTPBadPasswordErr"
	case UrlDataHFTPBadUserErr:
		return "UrlDataHFTPBadUserErr"
	case UrlDataHFTPDataConnectionErr:
		return "UrlDataHFTPDataConnectionErr"
	case UrlDataHFTPFilenameErr:
		return "UrlDataHFTPFilenameErr"
	case UrlDataHFTPNeedPasswordErr:
		return "UrlDataHFTPNeedPasswordErr"
	case UrlDataHFTPNoDirectoryErr:
		return "UrlDataHFTPNoDirectoryErr"
	case UrlDataHFTPNoNetDriverErr:
		return "UrlDataHFTPNoNetDriverErr"
	case UrlDataHFTPNoPasswordErr:
		return "UrlDataHFTPNoPasswordErr"
	case UrlDataHFTPPermissionsErr:
		return "UrlDataHFTPPermissionsErr"
	case UrlDataHFTPProtocolErr:
		return "UrlDataHFTPProtocolErr"
	case UrlDataHFTPQuotaErr:
		return "UrlDataHFTPQuotaErr"
	case UrlDataHFTPServerDisconnectedErr:
		return "UrlDataHFTPServerDisconnectedErr"
	case UrlDataHFTPServerErr:
		return "UrlDataHFTPServerErr"
	case UrlDataHFTPShutdownErr:
		return "UrlDataHFTPShutdownErr"
	case UrlDataHFTPURLErr:
		return "UrlDataHFTPURLErr"
	case UrlDataHHTTPNoNetDriverErr:
		return "UrlDataHHTTPNoNetDriverErr"
	case UrlDataHHTTPProtocolErr:
		return "UrlDataHHTTPProtocolErr"
	case UrlDataHHTTPRedirectErr:
		return "UrlDataHHTTPRedirectErr"
	case UrlDataHHTTPURLErr:
		return "UrlDataHHTTPURLErr"
	case WackBadFileErr:
		return "WackBadFileErr"
	case WackBadMetaDataErr:
		return "WackBadMetaDataErr"
	case WackForkNotFoundErr:
		return "WackForkNotFoundErr"
	default:
		return fmt.Sprintf("CannotFindAtomErr(%d)", e)
	}
}

type Cfrag int

const (
	CfragAbortClosureErr Cfrag = -2830
	CfragArchitectureErr Cfrag = -2823
	CfragCFMInternalErr Cfrag = -2819
	CfragCFMStartupErr Cfrag = -2818
	CfragCFragRsrcErr Cfrag = -2857
	CfragClosureIDErr Cfrag = -2829
	CfragConnectionIDErr Cfrag = -2801
	CfragContainerIDErr Cfrag = -2828
	CfragContextIDErr Cfrag = -2800
	CfragDupRegistrationErr Cfrag = -2805
	CfragExecFileRefErr Cfrag = -2854
	CfragFileSizeErr Cfrag = -2825
	CfragFirstErrCode Cfrag = -2800
	CfragFirstReservedCode Cfrag = -2897
	CfragFragmentCorruptErr Cfrag = -2820
	CfragFragmentFormatErr Cfrag = -2806
	CfragFragmentUsageErr Cfrag = -2824
	CfragImportTooNewErr Cfrag = -2814
	CfragImportTooOldErr Cfrag = -2813
	CfragInitAtBootErr Cfrag = -2816
	CfragInitFunctionErr Cfrag = -2821
	CfragInitLoopErr Cfrag = -2815
	CfragInitOrderErr Cfrag = -2812
	CfragLastErrCode Cfrag = -2899
	CfragLibConnErr Cfrag = -2817
	CfragMapFileErr Cfrag = -2851
	CfragNoApplicationErr Cfrag = -2822
	CfragNoClientMemErr Cfrag = -2810
	CfragNoIDsErr Cfrag = -2811
	CfragNoLibraryErr Cfrag = -2804
	CfragNoPositionErr Cfrag = -2808
	CfragNoPrivateMemErr Cfrag = -2809
	CfragNoRegistrationErr Cfrag = -2827
	CfragNoSectionErr Cfrag = -2803
	CfragNoSymbolErr Cfrag = -2802
	CfragNotClosureErr Cfrag = -2826
	CfragOutputLengthErr Cfrag = -2831
	CfragReservedCode_1 Cfrag = -2899
	CfragReservedCode_2 Cfrag = -2898
	CfragReservedCode_3 Cfrag = -2897
	CfragRsrcForkErr Cfrag = -2856
	CfragStdFolderErr Cfrag = -2855
	CfragUnresolvedErr Cfrag = -2807
)

func (e Cfrag) String() string {
	switch e {
	case CfragAbortClosureErr:
		return "CfragAbortClosureErr"
	case CfragArchitectureErr:
		return "CfragArchitectureErr"
	case CfragCFMInternalErr:
		return "CfragCFMInternalErr"
	case CfragCFMStartupErr:
		return "CfragCFMStartupErr"
	case CfragCFragRsrcErr:
		return "CfragCFragRsrcErr"
	case CfragClosureIDErr:
		return "CfragClosureIDErr"
	case CfragConnectionIDErr:
		return "CfragConnectionIDErr"
	case CfragContainerIDErr:
		return "CfragContainerIDErr"
	case CfragContextIDErr:
		return "CfragContextIDErr"
	case CfragDupRegistrationErr:
		return "CfragDupRegistrationErr"
	case CfragExecFileRefErr:
		return "CfragExecFileRefErr"
	case CfragFileSizeErr:
		return "CfragFileSizeErr"
	case CfragFirstReservedCode:
		return "CfragFirstReservedCode"
	case CfragFragmentCorruptErr:
		return "CfragFragmentCorruptErr"
	case CfragFragmentFormatErr:
		return "CfragFragmentFormatErr"
	case CfragFragmentUsageErr:
		return "CfragFragmentUsageErr"
	case CfragImportTooNewErr:
		return "CfragImportTooNewErr"
	case CfragImportTooOldErr:
		return "CfragImportTooOldErr"
	case CfragInitAtBootErr:
		return "CfragInitAtBootErr"
	case CfragInitFunctionErr:
		return "CfragInitFunctionErr"
	case CfragInitLoopErr:
		return "CfragInitLoopErr"
	case CfragInitOrderErr:
		return "CfragInitOrderErr"
	case CfragLastErrCode:
		return "CfragLastErrCode"
	case CfragLibConnErr:
		return "CfragLibConnErr"
	case CfragMapFileErr:
		return "CfragMapFileErr"
	case CfragNoApplicationErr:
		return "CfragNoApplicationErr"
	case CfragNoClientMemErr:
		return "CfragNoClientMemErr"
	case CfragNoIDsErr:
		return "CfragNoIDsErr"
	case CfragNoLibraryErr:
		return "CfragNoLibraryErr"
	case CfragNoPositionErr:
		return "CfragNoPositionErr"
	case CfragNoPrivateMemErr:
		return "CfragNoPrivateMemErr"
	case CfragNoRegistrationErr:
		return "CfragNoRegistrationErr"
	case CfragNoSectionErr:
		return "CfragNoSectionErr"
	case CfragNoSymbolErr:
		return "CfragNoSymbolErr"
	case CfragNotClosureErr:
		return "CfragNotClosureErr"
	case CfragOutputLengthErr:
		return "CfragOutputLengthErr"
	case CfragReservedCode_2:
		return "CfragReservedCode_2"
	case CfragRsrcForkErr:
		return "CfragRsrcForkErr"
	case CfragStdFolderErr:
		return "CfragStdFolderErr"
	case CfragUnresolvedErr:
		return "CfragUnresolvedErr"
	default:
		return fmt.Sprintf("Cfrag(%d)", e)
	}
}

type Cm int

const (
	// CmCantConcatenateError: Profiles cannot be concatenated
	CmCantConcatenateError Cm = -178
	// CmCantCopyModifiedV1Profile: It is illegal to copy version 1.0 profiles that have been modified
	CmCantCopyModifiedV1Profile Cm = -4215
	// CmCantDeleteElement: Cannot delete the specified profile element
	CmCantDeleteElement Cm = -4202
	// CmCantDeleteProfile: Responder error
	CmCantDeleteProfile Cm = -180
	// CmCantGamutCheckError: Gamut checking not supported by this color world—that is, the color world does not contain a gamut table because it was built with gamut checking turned off
	CmCantGamutCheckError Cm = -4217
	// CmCantXYZ: CMM does not handle XYZ color space
	CmCantXYZ Cm = -179
	// CmElementTagNotFound: The tag you specified is not in the specified profile
	CmElementTagNotFound Cm = -4200
	// CmErrIncompatibleProfile: Unspecified profile error
	CmErrIncompatibleProfile Cm = -4208
	// CmFatalProfileErr: Returned from File Manager while updating a profile file in response to [CMUpdateProfile]; profile content may be corrupted
	CmFatalProfileErr Cm = -4203
	// CmIndexRangeErr: Tag index out of range
	CmIndexRangeErr Cm = -4201
	// CmInvalidColorSpace: Profile color space does not match bitmap type
	CmInvalidColorSpace Cm = -4209
	// CmInvalidDstMap: Destination pix/bit map was invalid
	CmInvalidDstMap Cm = -4211
	// CmInvalidProfile: Profile reference is invalid or refers to an inappropriate profile
	CmInvalidProfile Cm = -4204
	// CmInvalidProfileComment: Bad profile comment during `drawpicture`
	CmInvalidProfileComment Cm = -4213
	// CmInvalidProfileLocation: Operation not supported for this profile location
	CmInvalidProfileLocation Cm = -4205
	// CmInvalidSearch: Bad search handle
	CmInvalidSearch Cm = -4206
	// CmInvalidSrcMap: Source pixel map or bitmap was invalid
	CmInvalidSrcMap Cm = -4210
	// CmMethodError: An error occurred during the CMM arbitration process that determines the CMM to use
	CmMethodError Cm = -171
	// CmMethodNotFound: CMM not present
	CmMethodNotFound Cm = -175
	// CmNamedColorNotFound: The specified named color was not found in the specified profile
	CmNamedColorNotFound Cm = -4216
	// CmNoCurrentProfile: Responder error
	CmNoCurrentProfile Cm = -182
	// CmNoGDevicesError: Begin matching or end matching—no graphics devices available
	CmNoGDevicesError Cm = -4212
	// CmProfileError: There is something wrong with the content of the profile
	CmProfileError Cm = -170
	// CmProfileNotFound: Responder error
	CmProfileNotFound Cm = -176
	// CmProfilesIdentical: Profiles are the same
	CmProfilesIdentical Cm = -177
	// CmRangeOverFlow: One or more output color value overflows in color conversion; all input color values will be converted and the overflow will be clipped
	CmRangeOverFlow Cm = -4214
	// CmSearchError: Internal error occurred during profile search
	CmSearchError Cm = -4207
	// CmUnsupportedDataType: Responder error
	CmUnsupportedDataType Cm = -181
)

func (e Cm) String() string {
	switch e {
	case CmCantConcatenateError:
		return "CmCantConcatenateError"
	case CmCantCopyModifiedV1Profile:
		return "CmCantCopyModifiedV1Profile"
	case CmCantDeleteElement:
		return "CmCantDeleteElement"
	case CmCantDeleteProfile:
		return "CmCantDeleteProfile"
	case CmCantGamutCheckError:
		return "CmCantGamutCheckError"
	case CmCantXYZ:
		return "CmCantXYZ"
	case CmElementTagNotFound:
		return "CmElementTagNotFound"
	case CmErrIncompatibleProfile:
		return "CmErrIncompatibleProfile"
	case CmFatalProfileErr:
		return "CmFatalProfileErr"
	case CmIndexRangeErr:
		return "CmIndexRangeErr"
	case CmInvalidColorSpace:
		return "CmInvalidColorSpace"
	case CmInvalidDstMap:
		return "CmInvalidDstMap"
	case CmInvalidProfile:
		return "CmInvalidProfile"
	case CmInvalidProfileComment:
		return "CmInvalidProfileComment"
	case CmInvalidProfileLocation:
		return "CmInvalidProfileLocation"
	case CmInvalidSearch:
		return "CmInvalidSearch"
	case CmInvalidSrcMap:
		return "CmInvalidSrcMap"
	case CmMethodError:
		return "CmMethodError"
	case CmMethodNotFound:
		return "CmMethodNotFound"
	case CmNamedColorNotFound:
		return "CmNamedColorNotFound"
	case CmNoCurrentProfile:
		return "CmNoCurrentProfile"
	case CmNoGDevicesError:
		return "CmNoGDevicesError"
	case CmProfileError:
		return "CmProfileError"
	case CmProfileNotFound:
		return "CmProfileNotFound"
	case CmProfilesIdentical:
		return "CmProfilesIdentical"
	case CmRangeOverFlow:
		return "CmRangeOverFlow"
	case CmSearchError:
		return "CmSearchError"
	case CmUnsupportedDataType:
		return "CmUnsupportedDataType"
	default:
		return fmt.Sprintf("Cm(%d)", e)
	}
}

type CodecErr int

const (
	BadCodecCharacterizationErr CodecErr = -8993
	CodecAbortErr CodecErr = -8967
	CodecBadDataErr CodecErr = -8969
	CodecCantQueueErr CodecErr = -8975
	CodecCantWhenErr CodecErr = -8974
	CodecConditionErr CodecErr = -8972
	CodecDataVersErr CodecErr = -8970
	CodecDisabledErr CodecErr = -8978
	CodecDroppedFrameErr CodecErr = -8989
	CodecErrValue CodecErr = -8960
	CodecExtensionNotFoundErr CodecErr = -8971
	CodecImageBufErr CodecErr = -8965
	CodecNeedAccessKeyErr CodecErr = -8987
	CodecNeedToFlushChainErr CodecErr = -8979
	CodecNoMemoryPleaseWaitErr CodecErr = -8977
	CodecNothingToBlitErr CodecErr = -8976
	CodecOffscreenFailedErr CodecErr = -8988
	CodecOffscreenFailedPleaseRetryErr CodecErr = -8992
	CodecOpenErr CodecErr = -8973
	CodecParameterDialogConfirm CodecErr = -8986
	CodecScreenBufErr CodecErr = -8964
	CodecSizeErr CodecErr = -8963
	CodecSpoolErr CodecErr = -8966
	CodecUnimpErr CodecErr = -8962
	CodecWouldOffscreenErr CodecErr = -8968
	DirectXObjectAlreadyExists CodecErr = -8990
	LockPortBitsBadPortErr CodecErr = -8984
	LockPortBitsBadSurfaceErr CodecErr = -8980
	LockPortBitsSurfaceLostErr CodecErr = -8985
	LockPortBitsWindowClippedErr CodecErr = -8983
	LockPortBitsWindowMovedErr CodecErr = -8981
	LockPortBitsWindowResizedErr CodecErr = -8982
	LockPortBitsWrongGDeviceErr CodecErr = -8991
	NoCodecErr CodecErr = -8961
	NoThumbnailFoundErr CodecErr = -8994
	ScTypeNotFoundErr CodecErr = -8971
)

func (e CodecErr) String() string {
	switch e {
	case BadCodecCharacterizationErr:
		return "BadCodecCharacterizationErr"
	case CodecAbortErr:
		return "CodecAbortErr"
	case CodecBadDataErr:
		return "CodecBadDataErr"
	case CodecCantQueueErr:
		return "CodecCantQueueErr"
	case CodecCantWhenErr:
		return "CodecCantWhenErr"
	case CodecConditionErr:
		return "CodecConditionErr"
	case CodecDataVersErr:
		return "CodecDataVersErr"
	case CodecDisabledErr:
		return "CodecDisabledErr"
	case CodecDroppedFrameErr:
		return "CodecDroppedFrameErr"
	case CodecErrValue:
		return "CodecErrValue"
	case CodecExtensionNotFoundErr:
		return "CodecExtensionNotFoundErr"
	case CodecImageBufErr:
		return "CodecImageBufErr"
	case CodecNeedAccessKeyErr:
		return "CodecNeedAccessKeyErr"
	case CodecNeedToFlushChainErr:
		return "CodecNeedToFlushChainErr"
	case CodecNoMemoryPleaseWaitErr:
		return "CodecNoMemoryPleaseWaitErr"
	case CodecNothingToBlitErr:
		return "CodecNothingToBlitErr"
	case CodecOffscreenFailedErr:
		return "CodecOffscreenFailedErr"
	case CodecOffscreenFailedPleaseRetryErr:
		return "CodecOffscreenFailedPleaseRetryErr"
	case CodecOpenErr:
		return "CodecOpenErr"
	case CodecParameterDialogConfirm:
		return "CodecParameterDialogConfirm"
	case CodecScreenBufErr:
		return "CodecScreenBufErr"
	case CodecSizeErr:
		return "CodecSizeErr"
	case CodecSpoolErr:
		return "CodecSpoolErr"
	case CodecUnimpErr:
		return "CodecUnimpErr"
	case CodecWouldOffscreenErr:
		return "CodecWouldOffscreenErr"
	case DirectXObjectAlreadyExists:
		return "DirectXObjectAlreadyExists"
	case LockPortBitsBadPortErr:
		return "LockPortBitsBadPortErr"
	case LockPortBitsBadSurfaceErr:
		return "LockPortBitsBadSurfaceErr"
	case LockPortBitsSurfaceLostErr:
		return "LockPortBitsSurfaceLostErr"
	case LockPortBitsWindowClippedErr:
		return "LockPortBitsWindowClippedErr"
	case LockPortBitsWindowMovedErr:
		return "LockPortBitsWindowMovedErr"
	case LockPortBitsWindowResizedErr:
		return "LockPortBitsWindowResizedErr"
	case LockPortBitsWrongGDeviceErr:
		return "LockPortBitsWrongGDeviceErr"
	case NoCodecErr:
		return "NoCodecErr"
	case NoThumbnailFoundErr:
		return "NoThumbnailFoundErr"
	default:
		return fmt.Sprintf("CodecErr(%d)", e)
	}
}

type Collection int

const (
	CollectionIndexRangeErr Collection = -5752
	CollectionItemLockedErr Collection = -5750
	CollectionItemNotFoundErr Collection = -5751
	CollectionVersionErr Collection = -5753
)

func (e Collection) String() string {
	switch e {
	case CollectionIndexRangeErr:
		return "CollectionIndexRangeErr"
	case CollectionItemLockedErr:
		return "CollectionItemLockedErr"
	case CollectionItemNotFoundErr:
		return "CollectionItemNotFoundErr"
	case CollectionVersionErr:
		return "CollectionVersionErr"
	default:
		return fmt.Sprintf("Collection(%d)", e)
	}
}

type ComponentDllLoadErr int

const (
	// ComponentDllEntryNotFoundErr: Windows error returned when a component is loading.
	ComponentDllEntryNotFoundErr ComponentDllLoadErr = -2092
	// ComponentDllLoadErrValue: Windows error returned when a component is loading.
	ComponentDllLoadErrValue ComponentDllLoadErr = -2091
	ComponentNotThreadSafeErr ComponentDllLoadErr = -2098
	// QtmlDllEntryNotFoundErr: Windows error returned when the QuickTime Media Layer is loading.
	QtmlDllEntryNotFoundErr ComponentDllLoadErr = -2094
	// QtmlDllLoadErr: Windows error returned when the QuickTime Media Layer is loading.
	QtmlDllLoadErr ComponentDllLoadErr = -2093
	QtmlUninitialized ComponentDllLoadErr = -2095
	UnsupportedOSErr ComponentDllLoadErr = -2096
	UnsupportedProcessorErr ComponentDllLoadErr = -2097
)

func (e ComponentDllLoadErr) String() string {
	switch e {
	case ComponentDllEntryNotFoundErr:
		return "ComponentDllEntryNotFoundErr"
	case ComponentDllLoadErrValue:
		return "ComponentDllLoadErrValue"
	case ComponentNotThreadSafeErr:
		return "ComponentNotThreadSafeErr"
	case QtmlDllEntryNotFoundErr:
		return "QtmlDllEntryNotFoundErr"
	case QtmlDllLoadErr:
		return "QtmlDllLoadErr"
	case QtmlUninitialized:
		return "QtmlUninitialized"
	case UnsupportedOSErr:
		return "UnsupportedOSErr"
	case UnsupportedProcessorErr:
		return "UnsupportedProcessorErr"
	default:
		return fmt.Sprintf("ComponentDllLoadErr(%d)", e)
	}
}

const CoreFoundationUnknownErr int = -4960

type CouldNotResolveDataRef int

const (
	AuxiliaryExportDataUnavailable CouldNotResolveDataRef = -2058
	BadComponentType CouldNotResolveDataRef = -2005
	BadDataRefIndex CouldNotResolveDataRef = -2050
	BadEditIndex CouldNotResolveDataRef = -2033
	BadEditList CouldNotResolveDataRef = -2017
	BadImageDescription CouldNotResolveDataRef = -2001
	BadPublicMovieAtom CouldNotResolveDataRef = -2002
	BadTrackIndex CouldNotResolveDataRef = -2028
	// CantCreateSingleForkFile: The file to be created already exists.
	CantCreateSingleForkFile CouldNotResolveDataRef = -2022
	CantEnableTrack CouldNotResolveDataRef = -2035
	CantFindHandler CouldNotResolveDataRef = -2003
	CantOpenHandler CouldNotResolveDataRef = -2004
	CantPutPublicMovieAtom CouldNotResolveDataRef = -2016
	CouldNotResolveDataRefValue CouldNotResolveDataRef = -2000
	CouldNotUseAnExistingSample CouldNotResolveDataRef = -2052
	DataAlreadyClosed CouldNotResolveDataRef = -2045
	DataAlreadyOpenForWrite CouldNotResolveDataRef = -2044
	DataNoDataRef CouldNotResolveDataRef = -2047
	DataNotOpenForRead CouldNotResolveDataRef = -2042
	DataNotOpenForWrite CouldNotResolveDataRef = -2043
	EndOfDataReached CouldNotResolveDataRef = -2046
	FeatureUnsupported CouldNotResolveDataRef = -2053
	GWorldsNotSameDepthAndSizeErr CouldNotResolveDataRef = -2066
	InternalQuickTimeError CouldNotResolveDataRef = -2034
	InvalidChunkCache CouldNotResolveDataRef = -2040
	InvalidChunkNum CouldNotResolveDataRef = -2038
	InvalidDataRef CouldNotResolveDataRef = -2012
	InvalidDataRefContainer CouldNotResolveDataRef = -2049
	InvalidDuration CouldNotResolveDataRef = -2014
	InvalidEditState CouldNotResolveDataRef = -2023
	InvalidHandler CouldNotResolveDataRef = -2013
	InvalidImageIndexErr CouldNotResolveDataRef = -2068
	InvalidMedia CouldNotResolveDataRef = -2008
	InvalidMovie CouldNotResolveDataRef = -2010
	InvalidRect CouldNotResolveDataRef = -2036
	InvalidSampleDescIndex CouldNotResolveDataRef = -2039
	InvalidSampleDescription CouldNotResolveDataRef = -2041
	InvalidSampleNum CouldNotResolveDataRef = -2037
	InvalidSampleTable CouldNotResolveDataRef = -2011
	InvalidSpriteIDErr CouldNotResolveDataRef = -2069
	InvalidSpriteIndexErr CouldNotResolveDataRef = -2067
	InvalidSpritePropertyErr CouldNotResolveDataRef = -2065
	InvalidSpriteWorldPropertyErr CouldNotResolveDataRef = -2064
	InvalidTime CouldNotResolveDataRef = -2015
	InvalidTrack CouldNotResolveDataRef = -2009
	MaxSizeToGrowTooSmall CouldNotResolveDataRef = -2027
	MediaTypesDontMatch CouldNotResolveDataRef = -2018
	MissingRequiredParameterErr CouldNotResolveDataRef = -2063
	MovieTextNotFoundErr CouldNotResolveDataRef = -2062
	MovieToolboxUninitialized CouldNotResolveDataRef = -2020
	NoDataHandler CouldNotResolveDataRef = -2007
	NoDefaultDataRef CouldNotResolveDataRef = -2051
	NoMediaHandler CouldNotResolveDataRef = -2006
	NoMovieFound CouldNotResolveDataRef = -2048
	// NoRecordOfApp: A replica of the movieToolboxUninitialized error.
	NoRecordOfApp CouldNotResolveDataRef = -2020
	NoSoundTrackInMovieErr CouldNotResolveDataRef = -2055
	NoSourceTreeFoundErr CouldNotResolveDataRef = -2060
	NoVideoTrackInMovieErr CouldNotResolveDataRef = -2054
	NonMatchingEditState CouldNotResolveDataRef = -2024
	ProgressProcAborted CouldNotResolveDataRef = -2019
	SamplesAlreadyInMediaErr CouldNotResolveDataRef = -2059
	SoundSupportNotAvailableErr CouldNotResolveDataRef = -2056
	SourceNotFoundErr CouldNotResolveDataRef = -2061
	StaleEditState CouldNotResolveDataRef = -2025
	TimeNotInMedia CouldNotResolveDataRef = -2032
	TimeNotInTrack CouldNotResolveDataRef = -2031
	TrackIDNotFound CouldNotResolveDataRef = -2029
	TrackNotInMovie CouldNotResolveDataRef = -2030
	UnsupportedAuxiliaryImportData CouldNotResolveDataRef = -2057
	UserDataItemNotFound CouldNotResolveDataRef = -2026
	WfFileNotFound CouldNotResolveDataRef = -2021
)

func (e CouldNotResolveDataRef) String() string {
	switch e {
	case AuxiliaryExportDataUnavailable:
		return "AuxiliaryExportDataUnavailable"
	case BadComponentType:
		return "BadComponentType"
	case BadDataRefIndex:
		return "BadDataRefIndex"
	case BadEditIndex:
		return "BadEditIndex"
	case BadEditList:
		return "BadEditList"
	case BadImageDescription:
		return "BadImageDescription"
	case BadPublicMovieAtom:
		return "BadPublicMovieAtom"
	case BadTrackIndex:
		return "BadTrackIndex"
	case CantCreateSingleForkFile:
		return "CantCreateSingleForkFile"
	case CantEnableTrack:
		return "CantEnableTrack"
	case CantFindHandler:
		return "CantFindHandler"
	case CantOpenHandler:
		return "CantOpenHandler"
	case CantPutPublicMovieAtom:
		return "CantPutPublicMovieAtom"
	case CouldNotResolveDataRefValue:
		return "CouldNotResolveDataRefValue"
	case CouldNotUseAnExistingSample:
		return "CouldNotUseAnExistingSample"
	case DataAlreadyClosed:
		return "DataAlreadyClosed"
	case DataAlreadyOpenForWrite:
		return "DataAlreadyOpenForWrite"
	case DataNoDataRef:
		return "DataNoDataRef"
	case DataNotOpenForRead:
		return "DataNotOpenForRead"
	case DataNotOpenForWrite:
		return "DataNotOpenForWrite"
	case EndOfDataReached:
		return "EndOfDataReached"
	case FeatureUnsupported:
		return "FeatureUnsupported"
	case GWorldsNotSameDepthAndSizeErr:
		return "GWorldsNotSameDepthAndSizeErr"
	case InternalQuickTimeError:
		return "InternalQuickTimeError"
	case InvalidChunkCache:
		return "InvalidChunkCache"
	case InvalidChunkNum:
		return "InvalidChunkNum"
	case InvalidDataRef:
		return "InvalidDataRef"
	case InvalidDataRefContainer:
		return "InvalidDataRefContainer"
	case InvalidDuration:
		return "InvalidDuration"
	case InvalidEditState:
		return "InvalidEditState"
	case InvalidHandler:
		return "InvalidHandler"
	case InvalidImageIndexErr:
		return "InvalidImageIndexErr"
	case InvalidMedia:
		return "InvalidMedia"
	case InvalidMovie:
		return "InvalidMovie"
	case InvalidRect:
		return "InvalidRect"
	case InvalidSampleDescIndex:
		return "InvalidSampleDescIndex"
	case InvalidSampleDescription:
		return "InvalidSampleDescription"
	case InvalidSampleNum:
		return "InvalidSampleNum"
	case InvalidSampleTable:
		return "InvalidSampleTable"
	case InvalidSpriteIDErr:
		return "InvalidSpriteIDErr"
	case InvalidSpriteIndexErr:
		return "InvalidSpriteIndexErr"
	case InvalidSpritePropertyErr:
		return "InvalidSpritePropertyErr"
	case InvalidSpriteWorldPropertyErr:
		return "InvalidSpriteWorldPropertyErr"
	case InvalidTime:
		return "InvalidTime"
	case InvalidTrack:
		return "InvalidTrack"
	case MaxSizeToGrowTooSmall:
		return "MaxSizeToGrowTooSmall"
	case MediaTypesDontMatch:
		return "MediaTypesDontMatch"
	case MissingRequiredParameterErr:
		return "MissingRequiredParameterErr"
	case MovieTextNotFoundErr:
		return "MovieTextNotFoundErr"
	case MovieToolboxUninitialized:
		return "MovieToolboxUninitialized"
	case NoDataHandler:
		return "NoDataHandler"
	case NoDefaultDataRef:
		return "NoDefaultDataRef"
	case NoMediaHandler:
		return "NoMediaHandler"
	case NoMovieFound:
		return "NoMovieFound"
	case NoSoundTrackInMovieErr:
		return "NoSoundTrackInMovieErr"
	case NoSourceTreeFoundErr:
		return "NoSourceTreeFoundErr"
	case NoVideoTrackInMovieErr:
		return "NoVideoTrackInMovieErr"
	case NonMatchingEditState:
		return "NonMatchingEditState"
	case ProgressProcAborted:
		return "ProgressProcAborted"
	case SamplesAlreadyInMediaErr:
		return "SamplesAlreadyInMediaErr"
	case SoundSupportNotAvailableErr:
		return "SoundSupportNotAvailableErr"
	case SourceNotFoundErr:
		return "SourceNotFoundErr"
	case StaleEditState:
		return "StaleEditState"
	case TimeNotInMedia:
		return "TimeNotInMedia"
	case TimeNotInTrack:
		return "TimeNotInTrack"
	case TrackIDNotFound:
		return "TrackIDNotFound"
	case TrackNotInMovie:
		return "TrackNotInMovie"
	case UnsupportedAuxiliaryImportData:
		return "UnsupportedAuxiliaryImportData"
	case UserDataItemNotFound:
		return "UserDataItemNotFound"
	case WfFileNotFound:
		return "WfFileNotFound"
	default:
		return fmt.Sprintf("CouldNotResolveDataRef(%d)", e)
	}
}

type Dcm int

const (
	DcmBadDataSizeErr Dcm = -7117
	DcmBadDictionaryErr Dcm = -7102
	DcmBadFeatureErr Dcm = -7124
	DcmBadFieldInfoErr Dcm = -7111
	DcmBadFieldTypeErr Dcm = -7112
	DcmBadFindMethodErr Dcm = -7118
	DcmBadKeyErr Dcm = -7115
	DcmBadPropertyErr Dcm = -7119
	DcmBlockFullErr Dcm = -7107
	DcmBufferOverflowErr Dcm = -7127
	DcmDictionaryBusyErr Dcm = -7105
	DcmDictionaryNotOpenErr Dcm = -7104
	DcmDupRecordErr Dcm = -7109
	DcmIterationCompleteErr Dcm = -7126
	DcmNecessaryFieldErr Dcm = -7110
	DcmNoAccessMethodErr Dcm = -7122
	DcmNoFieldErr Dcm = -7113
	DcmNoRecordErr Dcm = -7108
	DcmNotDictionaryErr Dcm = -7101
	DcmParamErr Dcm = -7100
	DcmPermissionErr Dcm = -7103
	DcmProtectedErr Dcm = -7121
	DcmTooManyKeyErr Dcm = -7116
)

func (e Dcm) String() string {
	switch e {
	case DcmBadDataSizeErr:
		return "DcmBadDataSizeErr"
	case DcmBadDictionaryErr:
		return "DcmBadDictionaryErr"
	case DcmBadFeatureErr:
		return "DcmBadFeatureErr"
	case DcmBadFieldInfoErr:
		return "DcmBadFieldInfoErr"
	case DcmBadFieldTypeErr:
		return "DcmBadFieldTypeErr"
	case DcmBadFindMethodErr:
		return "DcmBadFindMethodErr"
	case DcmBadKeyErr:
		return "DcmBadKeyErr"
	case DcmBadPropertyErr:
		return "DcmBadPropertyErr"
	case DcmBlockFullErr:
		return "DcmBlockFullErr"
	case DcmBufferOverflowErr:
		return "DcmBufferOverflowErr"
	case DcmDictionaryBusyErr:
		return "DcmDictionaryBusyErr"
	case DcmDictionaryNotOpenErr:
		return "DcmDictionaryNotOpenErr"
	case DcmDupRecordErr:
		return "DcmDupRecordErr"
	case DcmIterationCompleteErr:
		return "DcmIterationCompleteErr"
	case DcmNecessaryFieldErr:
		return "DcmNecessaryFieldErr"
	case DcmNoAccessMethodErr:
		return "DcmNoAccessMethodErr"
	case DcmNoFieldErr:
		return "DcmNoFieldErr"
	case DcmNoRecordErr:
		return "DcmNoRecordErr"
	case DcmNotDictionaryErr:
		return "DcmNotDictionaryErr"
	case DcmParamErr:
		return "DcmParamErr"
	case DcmPermissionErr:
		return "DcmPermissionErr"
	case DcmProtectedErr:
		return "DcmProtectedErr"
	case DcmTooManyKeyErr:
		return "DcmTooManyKeyErr"
	default:
		return fmt.Sprintf("Dcm(%d)", e)
	}
}

type DdpSktErr int

const (
	DdpLenErr DdpSktErr = -92
	DdpSktErrValue DdpSktErr = -91
	ExcessCollsns DdpSktErr = -95
	LapProtErr DdpSktErr = -94
	NoBridgeErr DdpSktErr = -93
	PortInUse DdpSktErr = -97
	PortNotCf DdpSktErr = -98
	PortNotPwr DdpSktErr = -96
)

func (e DdpSktErr) String() string {
	switch e {
	case DdpLenErr:
		return "DdpLenErr"
	case DdpSktErrValue:
		return "DdpSktErrValue"
	case ExcessCollsns:
		return "ExcessCollsns"
	case LapProtErr:
		return "LapProtErr"
	case NoBridgeErr:
		return "NoBridgeErr"
	case PortInUse:
		return "PortInUse"
	case PortNotCf:
		return "PortNotCf"
	case PortNotPwr:
		return "PortNotPwr"
	default:
		return fmt.Sprintf("DdpSktErr(%d)", e)
	}
}

type Debugging int

const (
	DebuggingDuplicateOptionErr Debugging = -13882
	DebuggingDuplicateSignatureErr Debugging = -13881
	DebuggingExecutionContextErr Debugging = -13880
	DebuggingInvalidNameErr Debugging = -13885
	DebuggingInvalidOptionErr Debugging = -13884
	DebuggingInvalidSignatureErr Debugging = -13883
	DebuggingNoCallbackErr Debugging = -13886
	DebuggingNoMatchErr Debugging = -13887
)

func (e Debugging) String() string {
	switch e {
	case DebuggingDuplicateOptionErr:
		return "DebuggingDuplicateOptionErr"
	case DebuggingDuplicateSignatureErr:
		return "DebuggingDuplicateSignatureErr"
	case DebuggingExecutionContextErr:
		return "DebuggingExecutionContextErr"
	case DebuggingInvalidNameErr:
		return "DebuggingInvalidNameErr"
	case DebuggingInvalidOptionErr:
		return "DebuggingInvalidOptionErr"
	case DebuggingInvalidSignatureErr:
		return "DebuggingInvalidSignatureErr"
	case DebuggingNoCallbackErr:
		return "DebuggingNoCallbackErr"
	case DebuggingNoMatchErr:
		return "DebuggingNoMatchErr"
	default:
		return fmt.Sprintf("Debugging(%d)", e)
	}
}

const DialogNoTimeoutErr int = -5640

type DigiUnimpErr int

const (
	// BadCallOrderErr: A status call was made before being set up first.
	BadCallOrderErr DigiUnimpErr = -2209
	// BadDepthErr: Can't digitize into the requested pixel depth.
	BadDepthErr DigiUnimpErr = -2207
	// DigiUnimpErrValue: Digitizer feature is unimplemented.
	DigiUnimpErrValue DigiUnimpErr = -2201
	// MatrixErr: Bad matrix; the digitizer did nothing.
	MatrixErr DigiUnimpErr = -2203
	// NoDMAErr: Can't do DMA digitizing; that is, can't go to the requested destination.
	NoDMAErr DigiUnimpErr = -2208
	// NoMoreKeyColorsErr: All the key indexes are in use.
	NoMoreKeyColorsErr DigiUnimpErr = -2205
	// NotExactMatrixErr: Warning of a bad matrix; the digitizer did its best.
	NotExactMatrixErr DigiUnimpErr = -2204
	// NotExactSizeErr: Can't digitize to the exact size requested.
	NotExactSizeErr DigiUnimpErr = -2206
	// QtParamErr: Bad input parameter (out of range, for example).
	QtParamErr DigiUnimpErr = -2202
)

func (e DigiUnimpErr) String() string {
	switch e {
	case BadCallOrderErr:
		return "BadCallOrderErr"
	case BadDepthErr:
		return "BadDepthErr"
	case DigiUnimpErrValue:
		return "DigiUnimpErrValue"
	case MatrixErr:
		return "MatrixErr"
	case NoDMAErr:
		return "NoDMAErr"
	case NoMoreKeyColorsErr:
		return "NoMoreKeyColorsErr"
	case NotExactMatrixErr:
		return "NotExactMatrixErr"
	case NotExactSizeErr:
		return "NotExactSizeErr"
	case QtParamErr:
		return "QtParamErr"
	default:
		return fmt.Sprintf("DigiUnimpErr(%d)", e)
	}
}

type Ds uint

const (
	DsAddressErr Ds = 2
	DsBadLibrary Ds = 1010
	DsBusError Ds = 1
	DsChkErr Ds = 5
	DsCoreErr Ds = 12
	DsFPErr Ds = 16
	DsIOCoreErr Ds = 14
	DsIllInstErr Ds = 3
	DsIrqErr Ds = 13
	DsLineAErr Ds = 9
	DsLineFErr Ds = 10
	DsLoadErr Ds = 15
	DsMiscErr Ds = 11
	DsMixedModeFailure Ds = 1011
	DsNoPackErr Ds = 17
	DsNoPk1 Ds = 18
	DsNoPk2 Ds = 19
	DsOvflowErr Ds = 6
	DsPrivErr Ds = 7
	DsTraceErr Ds = 8
	DsZeroDivErr Ds = 4
)

func (e Ds) String() string {
	switch e {
	case DsAddressErr:
		return "DsAddressErr"
	case DsBadLibrary:
		return "DsBadLibrary"
	case DsBusError:
		return "DsBusError"
	case DsChkErr:
		return "DsChkErr"
	case DsCoreErr:
		return "DsCoreErr"
	case DsFPErr:
		return "DsFPErr"
	case DsIOCoreErr:
		return "DsIOCoreErr"
	case DsIllInstErr:
		return "DsIllInstErr"
	case DsIrqErr:
		return "DsIrqErr"
	case DsLineAErr:
		return "DsLineAErr"
	case DsLineFErr:
		return "DsLineFErr"
	case DsLoadErr:
		return "DsLoadErr"
	case DsMiscErr:
		return "DsMiscErr"
	case DsMixedModeFailure:
		return "DsMixedModeFailure"
	case DsNoPackErr:
		return "DsNoPackErr"
	case DsNoPk1:
		return "DsNoPk1"
	case DsNoPk2:
		return "DsNoPk2"
	case DsOvflowErr:
		return "DsOvflowErr"
	case DsPrivErr:
		return "DsPrivErr"
	case DsTraceErr:
		return "DsTraceErr"
	case DsZeroDivErr:
		return "DsZeroDivErr"
	default:
		return fmt.Sprintf("Ds(%d)", e)
	}
}

type DsNoExtsMacsBug int

const (
	DsDisassemblerInstalled DsNoExtsMacsBug = -11
	DsExtensionsDisabled DsNoExtsMacsBug = -13
	DsGreeting DsNoExtsMacsBug = 40
	DsMacsBugInstalled DsNoExtsMacsBug = -10
	DsNoExtsDisassembler DsNoExtsMacsBug = -2
	DsNoExtsMacsBugValue DsNoExtsMacsBug = -1
	DsSysErr DsNoExtsMacsBug = 32767
	WDEFNFnd DsNoExtsMacsBug = 87
)

func (e DsNoExtsMacsBug) String() string {
	switch e {
	case DsDisassemblerInstalled:
		return "DsDisassemblerInstalled"
	case DsExtensionsDisabled:
		return "DsExtensionsDisabled"
	case DsGreeting:
		return "DsGreeting"
	case DsMacsBugInstalled:
		return "DsMacsBugInstalled"
	case DsNoExtsDisassembler:
		return "DsNoExtsDisassembler"
	case DsNoExtsMacsBugValue:
		return "DsNoExtsMacsBugValue"
	case DsSysErr:
		return "DsSysErr"
	case WDEFNFnd:
		return "WDEFNFnd"
	default:
		return fmt.Sprintf("DsNoExtsMacsBug(%d)", e)
	}
}

type DsNoFPU uint

const (
	Ds32BitMode DsNoFPU = 103
	DsBadPatch DsNoFPU = 99
	DsBufPtrTooLow DsNoFPU = 106
	DsCantHoldSystemHeap DsNoFPU = 114
	DsDirtyDisk DsNoFPU = 20004
	DsForcedQuit DsNoFPU = 20002
	DsGibblyMovedToDisabledFolder DsNoFPU = 117
	DsLostConnectionToNetworkDisk DsNoFPU = 121
	DsMBATAPISysError DsNoFPU = 29203
	DsMBATASysError DsNoFPU = 29202
	DsMBExternFlpySysError DsNoFPU = 29204
	DsMBFlpySysError DsNoFPU = 29201
	DsMBSysError DsNoFPU = 29200
	DsMacOSROMVersionTooOld DsNoFPU = 120
	DsMustUseFCBAccessors DsNoFPU = 119
	DsNeedToWriteBootBlocks DsNoFPU = 104
	DsNoFPUValue DsNoFPU = 90
	DsNoPatch DsNoFPU = 98
	DsNotEnoughRAMToBoot DsNoFPU = 105
	DsOldSystem DsNoFPU = 102
	DsPCCardATASysError DsNoFPU = 29205
	DsParityErr DsNoFPU = 101
	DsRAMDiskTooBig DsNoFPU = 122
	DsReinsert DsNoFPU = 30
	DsRemoveDisk DsNoFPU = 20003
	DsSCSIWarn DsNoFPU = 20010
	DsShutDownOrRestart DsNoFPU = 20000
	DsShutDownOrResume DsNoFPU = 20109
	DsSwitchOffOrRestart DsNoFPU = 20001
	DsSystemRequiresPowerPC DsNoFPU = 116
	DsUnBootableSystem DsNoFPU = 118
	DsVMBadBackingStore DsNoFPU = 113
	DsVMDeferredFuncTableFull DsNoFPU = 112
	DsWriteToSupervisorStackGuardPage DsNoFPU = 128
	ShutDownAlert DsNoFPU = 42
)

func (e DsNoFPU) String() string {
	switch e {
	case Ds32BitMode:
		return "Ds32BitMode"
	case DsBadPatch:
		return "DsBadPatch"
	case DsBufPtrTooLow:
		return "DsBufPtrTooLow"
	case DsCantHoldSystemHeap:
		return "DsCantHoldSystemHeap"
	case DsDirtyDisk:
		return "DsDirtyDisk"
	case DsForcedQuit:
		return "DsForcedQuit"
	case DsGibblyMovedToDisabledFolder:
		return "DsGibblyMovedToDisabledFolder"
	case DsLostConnectionToNetworkDisk:
		return "DsLostConnectionToNetworkDisk"
	case DsMBATAPISysError:
		return "DsMBATAPISysError"
	case DsMBATASysError:
		return "DsMBATASysError"
	case DsMBExternFlpySysError:
		return "DsMBExternFlpySysError"
	case DsMBFlpySysError:
		return "DsMBFlpySysError"
	case DsMBSysError:
		return "DsMBSysError"
	case DsMacOSROMVersionTooOld:
		return "DsMacOSROMVersionTooOld"
	case DsMustUseFCBAccessors:
		return "DsMustUseFCBAccessors"
	case DsNeedToWriteBootBlocks:
		return "DsNeedToWriteBootBlocks"
	case DsNoFPUValue:
		return "DsNoFPUValue"
	case DsNoPatch:
		return "DsNoPatch"
	case DsNotEnoughRAMToBoot:
		return "DsNotEnoughRAMToBoot"
	case DsOldSystem:
		return "DsOldSystem"
	case DsPCCardATASysError:
		return "DsPCCardATASysError"
	case DsParityErr:
		return "DsParityErr"
	case DsRAMDiskTooBig:
		return "DsRAMDiskTooBig"
	case DsReinsert:
		return "DsReinsert"
	case DsRemoveDisk:
		return "DsRemoveDisk"
	case DsSCSIWarn:
		return "DsSCSIWarn"
	case DsShutDownOrRestart:
		return "DsShutDownOrRestart"
	case DsShutDownOrResume:
		return "DsShutDownOrResume"
	case DsSwitchOffOrRestart:
		return "DsSwitchOffOrRestart"
	case DsSystemRequiresPowerPC:
		return "DsSystemRequiresPowerPC"
	case DsUnBootableSystem:
		return "DsUnBootableSystem"
	case DsVMBadBackingStore:
		return "DsVMBadBackingStore"
	case DsVMDeferredFuncTableFull:
		return "DsVMDeferredFuncTableFull"
	case DsWriteToSupervisorStackGuardPage:
		return "DsWriteToSupervisorStackGuardPage"
	case ShutDownAlert:
		return "ShutDownAlert"
	default:
		return fmt.Sprintf("DsNoFPU(%d)", e)
	}
}

type DsNoPk3 uint

const (
	DsBadLaunch DsNoPk3 = 26
	DsBadPatchHeader DsNoPk3 = 83
	DsBadSANEOpcode DsNoPk3 = 81
	DsBadSlotInt DsNoPk3 = 51
	DsCDEFNotFound DsNoPk3 = 88
	DsFSErr DsNoPk3 = 27
	DsFinderErr DsNoPk3 = 41
	DsHMenuFindErr DsNoPk3 = 86
	DsMBarNFnd DsNoPk3 = 85
	DsMDEFNotFound DsNoPk3 = 89
	DsMemFullErr DsNoPk3 = 25
	DsNoPk3Value DsNoPk3 = 20
	DsNoPk4 DsNoPk3 = 21
	DsNoPk5 DsNoPk3 = 22
	DsNoPk6 DsNoPk3 = 23
	DsNoPk7 DsNoPk3 = 24
	DsStknHeap DsNoPk3 = 28
	DsWDEFNotFound DsNoPk3 = 87
	MenuPrgErr DsNoPk3 = 84
	NegZcbFreeErr DsNoPk3 = 33
)

func (e DsNoPk3) String() string {
	switch e {
	case DsBadLaunch:
		return "DsBadLaunch"
	case DsBadPatchHeader:
		return "DsBadPatchHeader"
	case DsBadSANEOpcode:
		return "DsBadSANEOpcode"
	case DsBadSlotInt:
		return "DsBadSlotInt"
	case DsCDEFNotFound:
		return "DsCDEFNotFound"
	case DsFSErr:
		return "DsFSErr"
	case DsFinderErr:
		return "DsFinderErr"
	case DsHMenuFindErr:
		return "DsHMenuFindErr"
	case DsMBarNFnd:
		return "DsMBarNFnd"
	case DsMDEFNotFound:
		return "DsMDEFNotFound"
	case DsMemFullErr:
		return "DsMemFullErr"
	case DsNoPk3Value:
		return "DsNoPk3Value"
	case DsNoPk4:
		return "DsNoPk4"
	case DsNoPk5:
		return "DsNoPk5"
	case DsNoPk6:
		return "DsNoPk6"
	case DsNoPk7:
		return "DsNoPk7"
	case DsStknHeap:
		return "DsStknHeap"
	case DsWDEFNotFound:
		return "DsWDEFNotFound"
	case MenuPrgErr:
		return "MenuPrgErr"
	case NegZcbFreeErr:
		return "NegZcbFreeErr"
	default:
		return fmt.Sprintf("DsNoPk3(%d)", e)
	}
}

type ELenErr int

const (
	ELenErrValue ELenErr = -92
	EMultiErr ELenErr = -91
)

func (e ELenErr) String() string {
	switch e {
	case ELenErrValue:
		return "ELenErrValue"
	case EMultiErr:
		return "EMultiErr"
	default:
		return fmt.Sprintf("ELenErr(%d)", e)
	}
}

type EditionMgrInitErr int

const (
	BadEditionFileErr EditionMgrInitErr = -453
	BadSectionErr EditionMgrInitErr = -451
	BadSubPartErr EditionMgrInitErr = -454
	ContainerAlreadyOpenWrn EditionMgrInitErr = -462
	ContainerNotFoundWrn EditionMgrInitErr = -461
	EditionMgrInitErrValue EditionMgrInitErr = -450
	MultiplePublisherWrn EditionMgrInitErr = -460
	NotRegisteredSectionErr EditionMgrInitErr = -452
	NotThePublisherWrn EditionMgrInitErr = -463
)

func (e EditionMgrInitErr) String() string {
	switch e {
	case BadEditionFileErr:
		return "BadEditionFileErr"
	case BadSectionErr:
		return "BadSectionErr"
	case BadSubPartErr:
		return "BadSubPartErr"
	case ContainerAlreadyOpenWrn:
		return "ContainerAlreadyOpenWrn"
	case ContainerNotFoundWrn:
		return "ContainerNotFoundWrn"
	case EditionMgrInitErrValue:
		return "EditionMgrInitErrValue"
	case MultiplePublisherWrn:
		return "MultiplePublisherWrn"
	case NotRegisteredSectionErr:
		return "NotRegisteredSectionErr"
	case NotThePublisherWrn:
		return "NotThePublisherWrn"
	default:
		return fmt.Sprintf("EditionMgrInitErr(%d)", e)
	}
}

type Err int

const (
	ErrAborted Err = -1279
	ErrAlreadyInImagingMode Err = -5243
	ErrAttention Err = -1276
	ErrCannotUndo Err = -5253
	ErrCppGeneral Err = -32000
	ErrCppLastSystemDefinedError Err = -32020
	ErrCppLastUserDefinedError Err = -32049
	ErrCppbad_alloc Err = -32001
	ErrCppbad_cast Err = -32002
	ErrCppbad_exception Err = -32003
	ErrCppbad_typeid Err = -32004
	ErrCppdomain_error Err = -32006
	ErrCppinvalid_argument Err = -32007
	ErrCppios_base_failure Err = -32014
	ErrCpplength_error Err = -32008
	ErrCpplogic_error Err = -32005
	ErrCppout_of_range Err = -32009
	ErrCppoverflow_error Err = -32011
	ErrCpprange_error Err = -32012
	ErrCppruntime_error Err = -32010
	ErrCppunderflow_error Err = -32013
	ErrDSPQueueSize Err = -1274
	ErrEmptyScrap Err = -5249
	ErrEndOfBody Err = -1813
	ErrEndOfDocument Err = -1812
	ErrEngineNotFound Err = -5244
	ErrFwdReset Err = -1275
	ErrInvalidRange Err = -5246
	ErrIteratorReachedEnd Err = -5245
	ErrMarginWilllNotFit Err = -5241
	ErrNoHiliteText Err = -5248
	ErrNonContiuousAttribute Err = -5252
	ErrNotInImagingMode Err = -5242
	ErrOffsetInvalid Err = -1800
	ErrOffsetIsOutsideOfView Err = -1801
	ErrOffsetNotOnElementBounday Err = -5247
	ErrOpenDenied Err = -1273
	ErrOpening Err = -1277
	ErrReadOnlyText Err = -5250
	ErrRefNum Err = -1280
	ErrState Err = -1278
	ErrTopOfBody Err = -1811
	ErrTopOfDocument Err = -1810
	ErrUnknownAttributeTag Err = -5240
	ErrUnknownElement Err = -5251
)

func (e Err) String() string {
	switch e {
	case ErrAborted:
		return "ErrAborted"
	case ErrAlreadyInImagingMode:
		return "ErrAlreadyInImagingMode"
	case ErrAttention:
		return "ErrAttention"
	case ErrCannotUndo:
		return "ErrCannotUndo"
	case ErrCppGeneral:
		return "ErrCppGeneral"
	case ErrCppLastSystemDefinedError:
		return "ErrCppLastSystemDefinedError"
	case ErrCppLastUserDefinedError:
		return "ErrCppLastUserDefinedError"
	case ErrCppbad_alloc:
		return "ErrCppbad_alloc"
	case ErrCppbad_cast:
		return "ErrCppbad_cast"
	case ErrCppbad_exception:
		return "ErrCppbad_exception"
	case ErrCppbad_typeid:
		return "ErrCppbad_typeid"
	case ErrCppdomain_error:
		return "ErrCppdomain_error"
	case ErrCppinvalid_argument:
		return "ErrCppinvalid_argument"
	case ErrCppios_base_failure:
		return "ErrCppios_base_failure"
	case ErrCpplength_error:
		return "ErrCpplength_error"
	case ErrCpplogic_error:
		return "ErrCpplogic_error"
	case ErrCppout_of_range:
		return "ErrCppout_of_range"
	case ErrCppoverflow_error:
		return "ErrCppoverflow_error"
	case ErrCpprange_error:
		return "ErrCpprange_error"
	case ErrCppruntime_error:
		return "ErrCppruntime_error"
	case ErrCppunderflow_error:
		return "ErrCppunderflow_error"
	case ErrDSPQueueSize:
		return "ErrDSPQueueSize"
	case ErrEmptyScrap:
		return "ErrEmptyScrap"
	case ErrEndOfBody:
		return "ErrEndOfBody"
	case ErrEndOfDocument:
		return "ErrEndOfDocument"
	case ErrEngineNotFound:
		return "ErrEngineNotFound"
	case ErrFwdReset:
		return "ErrFwdReset"
	case ErrInvalidRange:
		return "ErrInvalidRange"
	case ErrIteratorReachedEnd:
		return "ErrIteratorReachedEnd"
	case ErrMarginWilllNotFit:
		return "ErrMarginWilllNotFit"
	case ErrNoHiliteText:
		return "ErrNoHiliteText"
	case ErrNonContiuousAttribute:
		return "ErrNonContiuousAttribute"
	case ErrNotInImagingMode:
		return "ErrNotInImagingMode"
	case ErrOffsetInvalid:
		return "ErrOffsetInvalid"
	case ErrOffsetIsOutsideOfView:
		return "ErrOffsetIsOutsideOfView"
	case ErrOffsetNotOnElementBounday:
		return "ErrOffsetNotOnElementBounday"
	case ErrOpenDenied:
		return "ErrOpenDenied"
	case ErrOpening:
		return "ErrOpening"
	case ErrReadOnlyText:
		return "ErrReadOnlyText"
	case ErrRefNum:
		return "ErrRefNum"
	case ErrState:
		return "ErrState"
	case ErrTopOfBody:
		return "ErrTopOfBody"
	case ErrTopOfDocument:
		return "ErrTopOfDocument"
	case ErrUnknownAttributeTag:
		return "ErrUnknownAttributeTag"
	case ErrUnknownElement:
		return "ErrUnknownElement"
	default:
		return fmt.Sprintf("Err(%d)", e)
	}
}

type ErrAE int

const (
	// ErrAEAccessorNotFound: There is no object accessor function forthe specified object class and container type
	ErrAEAccessorNotFound ErrAE = -1723
	// ErrAEBadKeyForm: Invalid key form.
	ErrAEBadKeyForm ErrAE = -10002
	// ErrAEBadListItem: Operation involving a list item failed
	ErrAEBadListItem ErrAE = -1705
	// ErrAEBadTestKey: The descriptor in a test key is neithera comparison descriptor nor a logical descriptor
	ErrAEBadTestKey ErrAE = -1726
	// ErrAEBufferTooSmall: Buffer for [AEFlattenDesc] toosmall
	ErrAEBufferTooSmall ErrAE = -1741
	// ErrAEBuildSyntaxError: [AEBuildDesc] andrelated functions detected a syntax error
	ErrAEBuildSyntaxError ErrAE = -1740
	// ErrAECantHandleClass: The Apple event handler can’t handle objectsof this class.
	ErrAECantHandleClass ErrAE = -10010
	// ErrAECantPutThatThere: In make new, duplicate, etc.
	ErrAECantPutThatThere ErrAE = -10024
	// ErrAECantSupplyType: Can’t supply the requested descriptortype for the data.
	ErrAECantSupplyType ErrAE = -10009
	// ErrAECantUndo: Can’t undo the previous Apple event oruser action.
	ErrAECantUndo ErrAE = -10015
	// ErrAECoercionFail: Data could not be coerced to the requesteddescriptor type
	ErrAECoercionFail ErrAE = -1700
	// ErrAECorruptData: Data in an Apple event could not be read
	ErrAECorruptData ErrAE = -1702
	// ErrAEDescIsNull: Attempt to perform an invalid operationon a null descriptor
	ErrAEDescIsNull ErrAE = -1739
	// ErrAEDescNotFound: Descriptor was not found
	ErrAEDescNotFound ErrAE = -1701
	// ErrAEDuplicateHandler: Attempt to install handler in table foridentical class and ID (1.1 or greater)
	ErrAEDuplicateHandler ErrAE = -1736
	// ErrAEEmptyListContainer: The container for an Apple event objectis specified by an empty list
	ErrAEEmptyListContainer ErrAE = -1730
	// ErrAEEventFailed: Apple event handler failed.
	ErrAEEventFailed ErrAE = -10000
	// ErrAEEventFiltered: Event has been filtered and should not bepropagated (1.1 or greater)
	ErrAEEventFiltered ErrAE = -1735
	// ErrAEEventNotHandled: Event wasn’t handled by an Apple eventhandler
	ErrAEEventNotHandled ErrAE = -1708
	ErrAEEventNotPermitted ErrAE = -1743
	// ErrAEHandlerNotFound: No handler found for an Apple event
	ErrAEHandlerNotFound ErrAE = -1717
	// ErrAEIllegalIndex: Not a valid list index
	ErrAEIllegalIndex ErrAE = -1719
	// ErrAEImpossibleRange: The range is not valid because it is impossiblefor a range to include the first and last objects that were specified;an example is a range in which the offset of the first object is greaterthan the offset of the last object
	ErrAEImpossibleRange ErrAE = -1720
	// ErrAEInTransaction: Couldn’t handle this command because itwasn’t part of the current transaction.
	ErrAEInTransaction ErrAE = -10011
	// ErrAEIndexTooLarge: The index of the event is too large to bevalid.
	ErrAEIndexTooLarge ErrAE = -10007
	ErrAELocalOnly ErrAE = -10016
	// ErrAENegativeCount: An object-counting function returned a negativeresult
	ErrAENegativeCount ErrAE = -1729
	// ErrAENewerVersion: Need a newer version of the Apple EventManager
	ErrAENewerVersion ErrAE = -1706
	// ErrAENoSuchLogical: The logical operator in a logical descriptoris not `kAEAND`, `kAEOR`,or `kAENOT`
	ErrAENoSuchLogical ErrAE = -1725
	// ErrAENoSuchObject: Runtime resolution of an object failed.
	ErrAENoSuchObject ErrAE = -1728
	// ErrAENoSuchTransaction: The transaction to which this command belongedisn’t a valid transaction.
	ErrAENoSuchTransaction ErrAE = -10012
	// ErrAENoUserInteraction: No user interaction allowed
	ErrAENoUserInteraction ErrAE = -1713
	// ErrAENoUserSelection: There is no user selection.
	ErrAENoUserSelection ErrAE = -10013
	// ErrAENotAEDesc: Not a valid descriptor
	ErrAENotAEDesc ErrAE = -1704
	// ErrAENotASingleObject: Handler only handles single objects.
	ErrAENotASingleObject ErrAE = -10014
	// ErrAENotASpecialFunction: Wrong keyword for a special function
	ErrAENotASpecialFunction ErrAE = -1714
	// ErrAENotAnElement: The specified object is a property, notan element.
	ErrAENotAnElement ErrAE = -10008
	// ErrAENotAnEnumMember: Enumerated value in [SetData] is notallowed for this property
	ErrAENotAnEnumMember ErrAE = -10023
	ErrAENotAnObjSpec ErrAE = -1727
	// ErrAENotAppleEvent: The event is not in AppleEvent format.
	ErrAENotAppleEvent ErrAE = -1707
	// ErrAENotModifiable: Can't set <object or data> to <object or data>.
	ErrAENotModifiable ErrAE = -10003
	// ErrAEParamMissed: A required parameter was not accessed.
	ErrAEParamMissed ErrAE = -1715
	// ErrAEPrivilegeError: A privilege violation occurred.
	ErrAEPrivilegeError ErrAE = -10004
	// ErrAEPropertiesClash: Illegal combination of properties settingsfor SetData, make new, or duplicate
	ErrAEPropertiesClash ErrAE = -10025
	// ErrAEReadDenied: The read operation was not allowed.
	ErrAEReadDenied ErrAE = -10005
	// ErrAEReceiveEscapeCurrent: Break out of lowest level only of [AEReceive] (1.1or greater)
	ErrAEReceiveEscapeCurrent ErrAE = -1734
	// ErrAEReceiveTerminate: Break out of all levels of [AEReceive] tothe topmost (1.1 or greater)
	ErrAEReceiveTerminate ErrAE = -1733
	// ErrAERecordingIsAlreadyOn: Recording is already on
	ErrAERecordingIsAlreadyOn ErrAE = -1732
	// ErrAEReplyNotArrived: Reply has not yet arrived
	ErrAEReplyNotArrived ErrAE = -1718
	// ErrAEReplyNotValid: [AEResetTimer] was passed an invalid reply
	ErrAEReplyNotValid ErrAE = -1709
	// ErrAEStreamAlreadyConverted: Attempt to convert a stream that has alreadybeen converted
	ErrAEStreamAlreadyConverted ErrAE = -1738
	// ErrAEStreamBadNesting: Nesting violation while streaming
	ErrAEStreamBadNesting ErrAE = -1737
	ErrAETargetAddressNotPermitted ErrAE = -1742
	// ErrAETimeout: Apple event timed out
	ErrAETimeout ErrAE = -1712
	// ErrAETypeError: A descriptor type mismatch occurred.
	ErrAETypeError ErrAE = -10001
	// ErrAEUnknownAddressType: Unknown Apple event address type
	ErrAEUnknownAddressType ErrAE = -1716
	// ErrAEUnknownObjectType: The object type isn’t recognized
	ErrAEUnknownObjectType ErrAE = -1731
	// ErrAEUnknownSendMode: Invalid sending mode was passed
	ErrAEUnknownSendMode ErrAE = -1710
	// ErrAEWaitCanceled: User canceled out of wait loop for replyor receipt
	ErrAEWaitCanceled ErrAE = -1711
	// ErrAEWriteDenied: Can't set <object or data> to <object or data>.
	ErrAEWriteDenied ErrAE = -10006
	// ErrAEWrongDataType: Wrong descriptor type
	ErrAEWrongDataType ErrAE = -1703
	// ErrAEWrongNumberArgs: The number of operands provided for the `kAENOT` logicaloperator is not 1
	ErrAEWrongNumberArgs ErrAE = -1721
)

func (e ErrAE) String() string {
	switch e {
	case ErrAEAccessorNotFound:
		return "ErrAEAccessorNotFound"
	case ErrAEBadKeyForm:
		return "ErrAEBadKeyForm"
	case ErrAEBadListItem:
		return "ErrAEBadListItem"
	case ErrAEBadTestKey:
		return "ErrAEBadTestKey"
	case ErrAEBufferTooSmall:
		return "ErrAEBufferTooSmall"
	case ErrAEBuildSyntaxError:
		return "ErrAEBuildSyntaxError"
	case ErrAECantHandleClass:
		return "ErrAECantHandleClass"
	case ErrAECantPutThatThere:
		return "ErrAECantPutThatThere"
	case ErrAECantSupplyType:
		return "ErrAECantSupplyType"
	case ErrAECantUndo:
		return "ErrAECantUndo"
	case ErrAECoercionFail:
		return "ErrAECoercionFail"
	case ErrAECorruptData:
		return "ErrAECorruptData"
	case ErrAEDescIsNull:
		return "ErrAEDescIsNull"
	case ErrAEDescNotFound:
		return "ErrAEDescNotFound"
	case ErrAEDuplicateHandler:
		return "ErrAEDuplicateHandler"
	case ErrAEEmptyListContainer:
		return "ErrAEEmptyListContainer"
	case ErrAEEventFailed:
		return "ErrAEEventFailed"
	case ErrAEEventFiltered:
		return "ErrAEEventFiltered"
	case ErrAEEventNotHandled:
		return "ErrAEEventNotHandled"
	case ErrAEEventNotPermitted:
		return "ErrAEEventNotPermitted"
	case ErrAEHandlerNotFound:
		return "ErrAEHandlerNotFound"
	case ErrAEIllegalIndex:
		return "ErrAEIllegalIndex"
	case ErrAEImpossibleRange:
		return "ErrAEImpossibleRange"
	case ErrAEInTransaction:
		return "ErrAEInTransaction"
	case ErrAEIndexTooLarge:
		return "ErrAEIndexTooLarge"
	case ErrAELocalOnly:
		return "ErrAELocalOnly"
	case ErrAENegativeCount:
		return "ErrAENegativeCount"
	case ErrAENewerVersion:
		return "ErrAENewerVersion"
	case ErrAENoSuchLogical:
		return "ErrAENoSuchLogical"
	case ErrAENoSuchObject:
		return "ErrAENoSuchObject"
	case ErrAENoSuchTransaction:
		return "ErrAENoSuchTransaction"
	case ErrAENoUserInteraction:
		return "ErrAENoUserInteraction"
	case ErrAENoUserSelection:
		return "ErrAENoUserSelection"
	case ErrAENotAEDesc:
		return "ErrAENotAEDesc"
	case ErrAENotASingleObject:
		return "ErrAENotASingleObject"
	case ErrAENotASpecialFunction:
		return "ErrAENotASpecialFunction"
	case ErrAENotAnElement:
		return "ErrAENotAnElement"
	case ErrAENotAnEnumMember:
		return "ErrAENotAnEnumMember"
	case ErrAENotAnObjSpec:
		return "ErrAENotAnObjSpec"
	case ErrAENotAppleEvent:
		return "ErrAENotAppleEvent"
	case ErrAENotModifiable:
		return "ErrAENotModifiable"
	case ErrAEParamMissed:
		return "ErrAEParamMissed"
	case ErrAEPrivilegeError:
		return "ErrAEPrivilegeError"
	case ErrAEPropertiesClash:
		return "ErrAEPropertiesClash"
	case ErrAEReadDenied:
		return "ErrAEReadDenied"
	case ErrAEReceiveEscapeCurrent:
		return "ErrAEReceiveEscapeCurrent"
	case ErrAEReceiveTerminate:
		return "ErrAEReceiveTerminate"
	case ErrAERecordingIsAlreadyOn:
		return "ErrAERecordingIsAlreadyOn"
	case ErrAEReplyNotArrived:
		return "ErrAEReplyNotArrived"
	case ErrAEReplyNotValid:
		return "ErrAEReplyNotValid"
	case ErrAEStreamAlreadyConverted:
		return "ErrAEStreamAlreadyConverted"
	case ErrAEStreamBadNesting:
		return "ErrAEStreamBadNesting"
	case ErrAETargetAddressNotPermitted:
		return "ErrAETargetAddressNotPermitted"
	case ErrAETimeout:
		return "ErrAETimeout"
	case ErrAETypeError:
		return "ErrAETypeError"
	case ErrAEUnknownAddressType:
		return "ErrAEUnknownAddressType"
	case ErrAEUnknownObjectType:
		return "ErrAEUnknownObjectType"
	case ErrAEUnknownSendMode:
		return "ErrAEUnknownSendMode"
	case ErrAEWaitCanceled:
		return "ErrAEWaitCanceled"
	case ErrAEWriteDenied:
		return "ErrAEWriteDenied"
	case ErrAEWrongDataType:
		return "ErrAEWrongDataType"
	case ErrAEWrongNumberArgs:
		return "ErrAEWrongNumberArgs"
	default:
		return fmt.Sprintf("ErrAE(%d)", e)
	}
}

type ErrAS int

const (
	// ErrASCantCompareMoreThan32k: Can’t perform operation on text longerthan 32K bytes.
	ErrASCantCompareMoreThan32k ErrAS = -2721
	// ErrASCantConsiderAndIgnore: Can’t both consider and ignore <attribute>.
	ErrASCantConsiderAndIgnore ErrAS = -2720
	// ErrASIllegalFormalParameter: <name> is illegal as a formal parameter.
	ErrASIllegalFormalParameter ErrAS = -2761
	ErrASInconsistentNames ErrAS = -2780
	// ErrASNoResultReturned: No result was returned for some argumentof this expression.
	ErrASNoResultReturned ErrAS = -2763
	// ErrASParameterNotForEvent: <name> is not a parameter name for the event <event>.
	ErrASParameterNotForEvent ErrAS = -2762
	// ErrASTerminologyNestingTooDeep: Tell statements are nested too deeply.
	ErrASTerminologyNestingTooDeep ErrAS = -2760
)

func (e ErrAS) String() string {
	switch e {
	case ErrASCantCompareMoreThan32k:
		return "ErrASCantCompareMoreThan32k"
	case ErrASCantConsiderAndIgnore:
		return "ErrASCantConsiderAndIgnore"
	case ErrASIllegalFormalParameter:
		return "ErrASIllegalFormalParameter"
	case ErrASInconsistentNames:
		return "ErrASInconsistentNames"
	case ErrASNoResultReturned:
		return "ErrASNoResultReturned"
	case ErrASParameterNotForEvent:
		return "ErrASParameterNotForEvent"
	case ErrASTerminologyNestingTooDeep:
		return "ErrASTerminologyNestingTooDeep"
	default:
		return fmt.Sprintf("ErrAS(%d)", e)
	}
}

type ErrCoreEndianData int

const (
	ErrCoreEndianDataDoesNotMatchFormat ErrCoreEndianData = -4942
	ErrCoreEndianDataTooLongForFormat ErrCoreEndianData = -4941
	ErrCoreEndianDataTooShortForFormat ErrCoreEndianData = -4940
)

func (e ErrCoreEndianData) String() string {
	switch e {
	case ErrCoreEndianDataDoesNotMatchFormat:
		return "ErrCoreEndianDataDoesNotMatchFormat"
	case ErrCoreEndianDataTooLongForFormat:
		return "ErrCoreEndianDataTooLongForFormat"
	case ErrCoreEndianDataTooShortForFormat:
		return "ErrCoreEndianDataTooShortForFormat"
	default:
		return fmt.Sprintf("ErrCoreEndianData(%d)", e)
	}
}

type ErrIA int

const (
	ErrIAAllocationErr ErrIA = -5381
	ErrIABufferTooSmall ErrIA = -5384
	ErrIACanceled ErrIA = -5385
	ErrIAEndOfTextRun ErrIA = -5388
	ErrIAInvalidDocument ErrIA = -5386
	ErrIANoErr ErrIA = 0
	ErrIANoMoreItems ErrIA = -5383
	ErrIAParamErr ErrIA = -5382
	ErrIATextExtractionErr ErrIA = -5387
	ErrIAUnknownErr ErrIA = -5380
)

func (e ErrIA) String() string {
	switch e {
	case ErrIAAllocationErr:
		return "ErrIAAllocationErr"
	case ErrIABufferTooSmall:
		return "ErrIABufferTooSmall"
	case ErrIACanceled:
		return "ErrIACanceled"
	case ErrIAEndOfTextRun:
		return "ErrIAEndOfTextRun"
	case ErrIAInvalidDocument:
		return "ErrIAInvalidDocument"
	case ErrIANoErr:
		return "ErrIANoErr"
	case ErrIANoMoreItems:
		return "ErrIANoMoreItems"
	case ErrIAParamErr:
		return "ErrIAParamErr"
	case ErrIATextExtractionErr:
		return "ErrIATextExtractionErr"
	case ErrIAUnknownErr:
		return "ErrIAUnknownErr"
	default:
		return fmt.Sprintf("ErrIA(%d)", e)
	}
}

type ErrInvalidWindowPtr int

const (
	ErrCorruptWindowDescription ErrInvalidWindowPtr = -5606
	ErrFloatingWindowsNotInitialized ErrInvalidWindowPtr = -5609
	ErrInvalidWindowProperty ErrInvalidWindowPtr = -5603
	ErrInvalidWindowPtrValue ErrInvalidWindowPtr = -5600
	ErrInvalidWindowRef ErrInvalidWindowPtr = -5600
	ErrUnrecognizedWindowClass ErrInvalidWindowPtr = -5605
	ErrUnsupportedWindowAttributesForClass ErrInvalidWindowPtr = -5601
	ErrUserWantsToDragWindow ErrInvalidWindowPtr = -5607
	ErrWindowDoesNotFitOnscreen ErrInvalidWindowPtr = -5611
	ErrWindowDoesNotHaveProxy ErrInvalidWindowPtr = -5602
	ErrWindowDoesntSupportFocus ErrInvalidWindowPtr = -30583
	ErrWindowNotFound ErrInvalidWindowPtr = -5610
	ErrWindowPropertyNotFound ErrInvalidWindowPtr = -5604
	ErrWindowRegionCodeInvalid ErrInvalidWindowPtr = -30593
	ErrWindowsAlreadyInitialized ErrInvalidWindowPtr = -5608
	WindowAppModalStateAlreadyExistsErr ErrInvalidWindowPtr = -5617
	WindowAttributeImmutableErr ErrInvalidWindowPtr = -5612
	WindowAttributesConflictErr ErrInvalidWindowPtr = -5613
	WindowGroupInvalidErr ErrInvalidWindowPtr = -5616
	WindowManagerInternalErr ErrInvalidWindowPtr = -5614
	WindowNoAppModalStateErr ErrInvalidWindowPtr = -5618
	WindowWrongStateErr ErrInvalidWindowPtr = -5615
)

func (e ErrInvalidWindowPtr) String() string {
	switch e {
	case ErrCorruptWindowDescription:
		return "ErrCorruptWindowDescription"
	case ErrFloatingWindowsNotInitialized:
		return "ErrFloatingWindowsNotInitialized"
	case ErrInvalidWindowProperty:
		return "ErrInvalidWindowProperty"
	case ErrInvalidWindowPtrValue:
		return "ErrInvalidWindowPtrValue"
	case ErrUnrecognizedWindowClass:
		return "ErrUnrecognizedWindowClass"
	case ErrUnsupportedWindowAttributesForClass:
		return "ErrUnsupportedWindowAttributesForClass"
	case ErrUserWantsToDragWindow:
		return "ErrUserWantsToDragWindow"
	case ErrWindowDoesNotFitOnscreen:
		return "ErrWindowDoesNotFitOnscreen"
	case ErrWindowDoesNotHaveProxy:
		return "ErrWindowDoesNotHaveProxy"
	case ErrWindowDoesntSupportFocus:
		return "ErrWindowDoesntSupportFocus"
	case ErrWindowNotFound:
		return "ErrWindowNotFound"
	case ErrWindowPropertyNotFound:
		return "ErrWindowPropertyNotFound"
	case ErrWindowRegionCodeInvalid:
		return "ErrWindowRegionCodeInvalid"
	case ErrWindowsAlreadyInitialized:
		return "ErrWindowsAlreadyInitialized"
	case WindowAppModalStateAlreadyExistsErr:
		return "WindowAppModalStateAlreadyExistsErr"
	case WindowAttributeImmutableErr:
		return "WindowAttributeImmutableErr"
	case WindowAttributesConflictErr:
		return "WindowAttributesConflictErr"
	case WindowGroupInvalidErr:
		return "WindowGroupInvalidErr"
	case WindowManagerInternalErr:
		return "WindowManagerInternalErr"
	case WindowNoAppModalStateErr:
		return "WindowNoAppModalStateErr"
	case WindowWrongStateErr:
		return "WindowWrongStateErr"
	default:
		return fmt.Sprintf("ErrInvalidWindowPtr(%d)", e)
	}
}

type ErrKC int

const (
	ErrKCAuthFailed ErrKC = -25293
	ErrKCBufferTooSmall ErrKC = -25301
	ErrKCCreateChainFailed ErrKC = -25318
	ErrKCDataNotAvailable ErrKC = -25316
	ErrKCDataNotModifiable ErrKC = -25317
	ErrKCDataTooLarge ErrKC = -25302
	ErrKCDuplicateCallback ErrKC = -25297
	ErrKCDuplicateItem ErrKC = -25299
	ErrKCDuplicateKeychain ErrKC = -25296
	ErrKCInteractionNotAllowed ErrKC = -25308
	ErrKCInteractionRequired ErrKC = -25315
	ErrKCInvalidCallback ErrKC = -25298
	ErrKCInvalidItemRef ErrKC = -25304
	ErrKCInvalidKeychain ErrKC = -25295
	ErrKCInvalidSearchRef ErrKC = -25305
	ErrKCItemNotFound ErrKC = -25300
	ErrKCKeySizeNotAllowed ErrKC = -25311
	ErrKCNoCertificateModule ErrKC = -25313
	ErrKCNoDefaultKeychain ErrKC = -25307
	ErrKCNoPolicyModule ErrKC = -25314
	ErrKCNoStorageModule ErrKC = -25312
	ErrKCNoSuchAttr ErrKC = -25303
	ErrKCNoSuchClass ErrKC = -25306
	ErrKCNoSuchKeychain ErrKC = -25294
	ErrKCNotAvailable ErrKC = -25291
	ErrKCReadOnly ErrKC = -25292
	ErrKCReadOnlyAttr ErrKC = -25309
	ErrKCWrongKCVersion ErrKC = -25310
)

func (e ErrKC) String() string {
	switch e {
	case ErrKCAuthFailed:
		return "ErrKCAuthFailed"
	case ErrKCBufferTooSmall:
		return "ErrKCBufferTooSmall"
	case ErrKCCreateChainFailed:
		return "ErrKCCreateChainFailed"
	case ErrKCDataNotAvailable:
		return "ErrKCDataNotAvailable"
	case ErrKCDataNotModifiable:
		return "ErrKCDataNotModifiable"
	case ErrKCDataTooLarge:
		return "ErrKCDataTooLarge"
	case ErrKCDuplicateCallback:
		return "ErrKCDuplicateCallback"
	case ErrKCDuplicateItem:
		return "ErrKCDuplicateItem"
	case ErrKCDuplicateKeychain:
		return "ErrKCDuplicateKeychain"
	case ErrKCInteractionNotAllowed:
		return "ErrKCInteractionNotAllowed"
	case ErrKCInteractionRequired:
		return "ErrKCInteractionRequired"
	case ErrKCInvalidCallback:
		return "ErrKCInvalidCallback"
	case ErrKCInvalidItemRef:
		return "ErrKCInvalidItemRef"
	case ErrKCInvalidKeychain:
		return "ErrKCInvalidKeychain"
	case ErrKCInvalidSearchRef:
		return "ErrKCInvalidSearchRef"
	case ErrKCItemNotFound:
		return "ErrKCItemNotFound"
	case ErrKCKeySizeNotAllowed:
		return "ErrKCKeySizeNotAllowed"
	case ErrKCNoCertificateModule:
		return "ErrKCNoCertificateModule"
	case ErrKCNoDefaultKeychain:
		return "ErrKCNoDefaultKeychain"
	case ErrKCNoPolicyModule:
		return "ErrKCNoPolicyModule"
	case ErrKCNoStorageModule:
		return "ErrKCNoStorageModule"
	case ErrKCNoSuchAttr:
		return "ErrKCNoSuchAttr"
	case ErrKCNoSuchClass:
		return "ErrKCNoSuchClass"
	case ErrKCNoSuchKeychain:
		return "ErrKCNoSuchKeychain"
	case ErrKCNotAvailable:
		return "ErrKCNotAvailable"
	case ErrKCReadOnly:
		return "ErrKCReadOnly"
	case ErrKCReadOnlyAttr:
		return "ErrKCReadOnlyAttr"
	case ErrKCWrongKCVersion:
		return "ErrKCWrongKCVersion"
	default:
		return fmt.Sprintf("ErrKC(%d)", e)
	}
}

type ErrMessageNotSupported int

const (
	ControlHandleInvalidErr ErrMessageNotSupported = -30599
	ControlInvalidDataVersionErr ErrMessageNotSupported = -30597
	ControlPropertyInvalid ErrMessageNotSupported = -5603
	ControlPropertyNotFoundErr ErrMessageNotSupported = -5604
	ErrCantEmbedIntoSelf ErrMessageNotSupported = -30594
	ErrCantEmbedRoot ErrMessageNotSupported = -30595
	ErrControlDoesntSupportFocus ErrMessageNotSupported = -30582
	ErrControlHiddenOrDisabled ErrMessageNotSupported = -30592
	ErrControlIsNotEmbedder ErrMessageNotSupported = -30590
	ErrControlsAlreadyExist ErrMessageNotSupported = -30589
	ErrCouldntSetFocus ErrMessageNotSupported = -30585
	ErrDataNotSupported ErrMessageNotSupported = -30581
	ErrDataSizeMismatch ErrMessageNotSupported = -30591
	ErrInvalidPartCode ErrMessageNotSupported = -30588
	ErrItemNotControl ErrMessageNotSupported = -30596
	ErrMessageNotSupportedValue ErrMessageNotSupported = -30580
	ErrNoRootControl ErrMessageNotSupported = -30586
	ErrRootAlreadyExists ErrMessageNotSupported = -30587
	ErrUnknownControl ErrMessageNotSupported = -30584
)

func (e ErrMessageNotSupported) String() string {
	switch e {
	case ControlHandleInvalidErr:
		return "ControlHandleInvalidErr"
	case ControlInvalidDataVersionErr:
		return "ControlInvalidDataVersionErr"
	case ControlPropertyInvalid:
		return "ControlPropertyInvalid"
	case ControlPropertyNotFoundErr:
		return "ControlPropertyNotFoundErr"
	case ErrCantEmbedIntoSelf:
		return "ErrCantEmbedIntoSelf"
	case ErrCantEmbedRoot:
		return "ErrCantEmbedRoot"
	case ErrControlDoesntSupportFocus:
		return "ErrControlDoesntSupportFocus"
	case ErrControlHiddenOrDisabled:
		return "ErrControlHiddenOrDisabled"
	case ErrControlIsNotEmbedder:
		return "ErrControlIsNotEmbedder"
	case ErrControlsAlreadyExist:
		return "ErrControlsAlreadyExist"
	case ErrCouldntSetFocus:
		return "ErrCouldntSetFocus"
	case ErrDataNotSupported:
		return "ErrDataNotSupported"
	case ErrDataSizeMismatch:
		return "ErrDataSizeMismatch"
	case ErrInvalidPartCode:
		return "ErrInvalidPartCode"
	case ErrItemNotControl:
		return "ErrItemNotControl"
	case ErrMessageNotSupportedValue:
		return "ErrMessageNotSupportedValue"
	case ErrNoRootControl:
		return "ErrNoRootControl"
	case ErrRootAlreadyExists:
		return "ErrRootAlreadyExists"
	case ErrUnknownControl:
		return "ErrUnknownControl"
	default:
		return fmt.Sprintf("ErrMessageNotSupported(%d)", e)
	}
}

type ErrOSA int

const (
	ErrOSAAppNotHighLevelEventAware ErrOSA = -2704
	ErrOSABadSelector ErrOSA = -1754
	ErrOSABadStorageType ErrOSA = -1752
	ErrOSACantAccess ErrOSA = -1728
	ErrOSACantAssign ErrOSA = -10006
	ErrOSACantCoerce ErrOSA = -1700
	ErrOSACantCreate ErrOSA = -2710
	ErrOSACantGetTerminology ErrOSA = -2709
	ErrOSACantLaunch ErrOSA = -2703
	ErrOSACantOpenComponent ErrOSA = -1762
	ErrOSACantStorePointers ErrOSA = -1763
	ErrOSAComponentMismatch ErrOSA = -1761
	ErrOSACorruptData ErrOSA = -1702
	ErrOSACorruptTerminology ErrOSA = -2705
	ErrOSADataBlockTooLarge ErrOSA = -2708
	ErrOSADataFormatObsolete ErrOSA = -1758
	ErrOSADataFormatTooNew ErrOSA = -1759
	ErrOSADivideByZero ErrOSA = -2701
	ErrOSAGeneralError ErrOSA = -2700
	ErrOSAInternalTableOverflow ErrOSA = -2707
	ErrOSAInvalidID ErrOSA = -1751
	ErrOSANoSuchDialect ErrOSA = -1757
	ErrOSANumericOverflow ErrOSA = -2702
	ErrOSARecordingIsAlreadyOn ErrOSA = -1732
	ErrOSAScriptError ErrOSA = -1753
	ErrOSASourceNotAvailable ErrOSA = -1756
	ErrOSAStackOverflow ErrOSA = -2706
	ErrOSASystemError ErrOSA = -1750
)

func (e ErrOSA) String() string {
	switch e {
	case ErrOSAAppNotHighLevelEventAware:
		return "ErrOSAAppNotHighLevelEventAware"
	case ErrOSABadSelector:
		return "ErrOSABadSelector"
	case ErrOSABadStorageType:
		return "ErrOSABadStorageType"
	case ErrOSACantAccess:
		return "ErrOSACantAccess"
	case ErrOSACantAssign:
		return "ErrOSACantAssign"
	case ErrOSACantCoerce:
		return "ErrOSACantCoerce"
	case ErrOSACantCreate:
		return "ErrOSACantCreate"
	case ErrOSACantGetTerminology:
		return "ErrOSACantGetTerminology"
	case ErrOSACantLaunch:
		return "ErrOSACantLaunch"
	case ErrOSACantOpenComponent:
		return "ErrOSACantOpenComponent"
	case ErrOSACantStorePointers:
		return "ErrOSACantStorePointers"
	case ErrOSAComponentMismatch:
		return "ErrOSAComponentMismatch"
	case ErrOSACorruptData:
		return "ErrOSACorruptData"
	case ErrOSACorruptTerminology:
		return "ErrOSACorruptTerminology"
	case ErrOSADataBlockTooLarge:
		return "ErrOSADataBlockTooLarge"
	case ErrOSADataFormatObsolete:
		return "ErrOSADataFormatObsolete"
	case ErrOSADataFormatTooNew:
		return "ErrOSADataFormatTooNew"
	case ErrOSADivideByZero:
		return "ErrOSADivideByZero"
	case ErrOSAGeneralError:
		return "ErrOSAGeneralError"
	case ErrOSAInternalTableOverflow:
		return "ErrOSAInternalTableOverflow"
	case ErrOSAInvalidID:
		return "ErrOSAInvalidID"
	case ErrOSANoSuchDialect:
		return "ErrOSANoSuchDialect"
	case ErrOSANumericOverflow:
		return "ErrOSANumericOverflow"
	case ErrOSARecordingIsAlreadyOn:
		return "ErrOSARecordingIsAlreadyOn"
	case ErrOSAScriptError:
		return "ErrOSAScriptError"
	case ErrOSASourceNotAvailable:
		return "ErrOSASourceNotAvailable"
	case ErrOSAStackOverflow:
		return "ErrOSAStackOverflow"
	case ErrOSASystemError:
		return "ErrOSASystemError"
	default:
		return fmt.Sprintf("ErrOSA(%d)", e)
	}
}

type ErrOSATypeError int

const (
	ErrOSATypeErrorValue ErrOSATypeError = -1703
	OSAControlFlowError ErrOSATypeError = -2755
	OSADuplicateHandler ErrOSATypeError = -2752
	OSADuplicateParameter ErrOSATypeError = -2750
	OSADuplicateProperty ErrOSATypeError = -2751
	OSAIllegalAccess ErrOSATypeError = -1723
	OSAIllegalAssign ErrOSATypeError = -10003
	OSAIllegalIndex ErrOSATypeError = -1719
	OSAIllegalRange ErrOSATypeError = -1720
	OSAInconsistentDeclarations ErrOSATypeError = -2754
	OSAMessageNotUnderstood ErrOSATypeError = -1708
	OSAMissingParameter ErrOSATypeError = -1701
	OSAParameterMismatch ErrOSATypeError = -1721
	OSASyntaxError ErrOSATypeError = -2740
	OSASyntaxTypeError ErrOSATypeError = -2741
	OSATokenTooLong ErrOSATypeError = -2742
	OSAUndefinedHandler ErrOSATypeError = -1717
	OSAUndefinedVariable ErrOSATypeError = -2753
)

func (e ErrOSATypeError) String() string {
	switch e {
	case ErrOSATypeErrorValue:
		return "ErrOSATypeErrorValue"
	case OSAControlFlowError:
		return "OSAControlFlowError"
	case OSADuplicateHandler:
		return "OSADuplicateHandler"
	case OSADuplicateParameter:
		return "OSADuplicateParameter"
	case OSADuplicateProperty:
		return "OSADuplicateProperty"
	case OSAIllegalAccess:
		return "OSAIllegalAccess"
	case OSAIllegalAssign:
		return "OSAIllegalAssign"
	case OSAIllegalIndex:
		return "OSAIllegalIndex"
	case OSAIllegalRange:
		return "OSAIllegalRange"
	case OSAInconsistentDeclarations:
		return "OSAInconsistentDeclarations"
	case OSAMessageNotUnderstood:
		return "OSAMessageNotUnderstood"
	case OSAMissingParameter:
		return "OSAMissingParameter"
	case OSAParameterMismatch:
		return "OSAParameterMismatch"
	case OSASyntaxError:
		return "OSASyntaxError"
	case OSASyntaxTypeError:
		return "OSASyntaxTypeError"
	case OSATokenTooLong:
		return "OSATokenTooLong"
	case OSAUndefinedHandler:
		return "OSAUndefinedHandler"
	case OSAUndefinedVariable:
		return "OSAUndefinedVariable"
	default:
		return fmt.Sprintf("ErrOSATypeError(%d)", e)
	}
}

const ErrTaskNotFound int = -10780

const EvtNotEnb uint = 1

type FidNotFound int

const (
	BadFCBErr FidNotFound = -1327
	BadFidErr FidNotFound = -1307
	CatChangedErr FidNotFound = -1304
	DesktopDamagedErr FidNotFound = -1305
	DiffVolErr FidNotFound = -1303
	EnvBadVers FidNotFound = -5501
	EnvNotPresent FidNotFound = -5500
	EnvVersTooBig FidNotFound = -5502
	ErrFSAttributeNotFound FidNotFound = -1427
	ErrFSBadAllocFlags FidNotFound = -1413
	ErrFSBadBuffer FidNotFound = -1403
	ErrFSBadFSRef FidNotFound = -1401
	ErrFSBadForkName FidNotFound = -1402
	ErrFSBadForkRef FidNotFound = -1404
	ErrFSBadInfoBitmap FidNotFound = -1405
	ErrFSBadItemCount FidNotFound = -1418
	ErrFSBadIteratorFlags FidNotFound = -1422
	ErrFSBadPosMode FidNotFound = -1412
	ErrFSBadSearchParams FidNotFound = -1419
	ErrFSForkExists FidNotFound = -1421
	ErrFSForkNotFound FidNotFound = -1409
	ErrFSIteratorNotFound FidNotFound = -1423
	ErrFSIteratorNotSupported FidNotFound = -1424
	ErrFSMissingCatInfo FidNotFound = -1406
	ErrFSMissingName FidNotFound = -1411
	ErrFSNameTooLong FidNotFound = -1410
	ErrFSNoMoreItems FidNotFound = -1417
	ErrFSNotAFolder FidNotFound = -1407
	ErrFSNotEnoughSpaceForOperation FidNotFound = -1429
	ErrFSOperationNotSupported FidNotFound = -1426
	ErrFSPropertyNotValid FidNotFound = -1428
	ErrFSQuotaExceeded FidNotFound = -1425
	ErrFSRefsDifferent FidNotFound = -1420
	ErrFSUnknownCall FidNotFound = -1400
	FidExists FidNotFound = -1301
	FidNotFoundValue FidNotFound = -1300
	FileBoundsErr FidNotFound = -1309
	FirstDskErr FidNotFound = -84
	FontDecError FidNotFound = -64
	FontNotDeclared FidNotFound = -65
	FontNotOutlineErr FidNotFound = -32615
	FontSubErr FidNotFound = -66
	FsDataTooBigErr FidNotFound = -1310
	LastDskErr FidNotFound = -64
	NoDriveErr FidNotFound = -64
	NoNybErr FidNotFound = -66
	NotAFileErr FidNotFound = -1302
	NotARemountErr FidNotFound = -1308
	OffLinErr FidNotFound = -65
	SameFileErr FidNotFound = -1306
	VolVMBusyErr FidNotFound = -1311
)

func (e FidNotFound) String() string {
	switch e {
	case BadFCBErr:
		return "BadFCBErr"
	case BadFidErr:
		return "BadFidErr"
	case CatChangedErr:
		return "CatChangedErr"
	case DesktopDamagedErr:
		return "DesktopDamagedErr"
	case DiffVolErr:
		return "DiffVolErr"
	case EnvBadVers:
		return "EnvBadVers"
	case EnvNotPresent:
		return "EnvNotPresent"
	case EnvVersTooBig:
		return "EnvVersTooBig"
	case ErrFSAttributeNotFound:
		return "ErrFSAttributeNotFound"
	case ErrFSBadAllocFlags:
		return "ErrFSBadAllocFlags"
	case ErrFSBadBuffer:
		return "ErrFSBadBuffer"
	case ErrFSBadFSRef:
		return "ErrFSBadFSRef"
	case ErrFSBadForkName:
		return "ErrFSBadForkName"
	case ErrFSBadForkRef:
		return "ErrFSBadForkRef"
	case ErrFSBadInfoBitmap:
		return "ErrFSBadInfoBitmap"
	case ErrFSBadItemCount:
		return "ErrFSBadItemCount"
	case ErrFSBadIteratorFlags:
		return "ErrFSBadIteratorFlags"
	case ErrFSBadPosMode:
		return "ErrFSBadPosMode"
	case ErrFSBadSearchParams:
		return "ErrFSBadSearchParams"
	case ErrFSForkExists:
		return "ErrFSForkExists"
	case ErrFSForkNotFound:
		return "ErrFSForkNotFound"
	case ErrFSIteratorNotFound:
		return "ErrFSIteratorNotFound"
	case ErrFSIteratorNotSupported:
		return "ErrFSIteratorNotSupported"
	case ErrFSMissingCatInfo:
		return "ErrFSMissingCatInfo"
	case ErrFSMissingName:
		return "ErrFSMissingName"
	case ErrFSNameTooLong:
		return "ErrFSNameTooLong"
	case ErrFSNoMoreItems:
		return "ErrFSNoMoreItems"
	case ErrFSNotAFolder:
		return "ErrFSNotAFolder"
	case ErrFSNotEnoughSpaceForOperation:
		return "ErrFSNotEnoughSpaceForOperation"
	case ErrFSOperationNotSupported:
		return "ErrFSOperationNotSupported"
	case ErrFSPropertyNotValid:
		return "ErrFSPropertyNotValid"
	case ErrFSQuotaExceeded:
		return "ErrFSQuotaExceeded"
	case ErrFSRefsDifferent:
		return "ErrFSRefsDifferent"
	case ErrFSUnknownCall:
		return "ErrFSUnknownCall"
	case FidExists:
		return "FidExists"
	case FidNotFoundValue:
		return "FidNotFoundValue"
	case FileBoundsErr:
		return "FileBoundsErr"
	case FirstDskErr:
		return "FirstDskErr"
	case FontDecError:
		return "FontDecError"
	case FontNotDeclared:
		return "FontNotDeclared"
	case FontNotOutlineErr:
		return "FontNotOutlineErr"
	case FontSubErr:
		return "FontSubErr"
	case FsDataTooBigErr:
		return "FsDataTooBigErr"
	case NotAFileErr:
		return "NotAFileErr"
	case NotARemountErr:
		return "NotARemountErr"
	case SameFileErr:
		return "SameFileErr"
	case VolVMBusyErr:
		return "VolVMBusyErr"
	default:
		return fmt.Sprintf("FidNotFound(%d)", e)
	}
}

type FirstPickerError int

const (
	BadProfileError FirstPickerError = -4008
	CantCreatePickerWindow FirstPickerError = -4004
	CantLoadPackage FirstPickerError = -4005
	CantLoadPicker FirstPickerError = -4003
	ColorSyncNotInstalled FirstPickerError = -4007
	FirstPickerErrorValue FirstPickerError = -4000
	InvalidPickerType FirstPickerError = -4000
	NoHelpForItem FirstPickerError = -4009
	PickerCantLive FirstPickerError = -4006
	PickerResourceError FirstPickerError = -4002
	RequiredFlagsDontMatch FirstPickerError = -4001
)

func (e FirstPickerError) String() string {
	switch e {
	case BadProfileError:
		return "BadProfileError"
	case CantCreatePickerWindow:
		return "CantCreatePickerWindow"
	case CantLoadPackage:
		return "CantLoadPackage"
	case CantLoadPicker:
		return "CantLoadPicker"
	case ColorSyncNotInstalled:
		return "ColorSyncNotInstalled"
	case FirstPickerErrorValue:
		return "FirstPickerErrorValue"
	case NoHelpForItem:
		return "NoHelpForItem"
	case PickerCantLive:
		return "PickerCantLive"
	case PickerResourceError:
		return "PickerResourceError"
	case RequiredFlagsDontMatch:
		return "RequiredFlagsDontMatch"
	default:
		return fmt.Sprintf("FirstPickerError(%d)", e)
	}
}

type Form uint

const (
	// FormAbsolutePosition: # Discussion
	FormAbsolutePosition Form = 'i'<<24 | 'n'<<16 | 'd'<<8 | 'x' // 'indx'
	// FormName: Specifies the Apple event object by name.
	FormName Form = 'n'<<24 | 'a'<<16 | 'm'<<8 | 'e' // 'name'
	// FormPropertyID: Specifies the property ID for an element’s property.
	FormPropertyID Form = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
	// FormRange: # Discussion
	FormRange Form = 'r'<<24 | 'a'<<16 | 'n'<<8 | 'g' // 'rang'
	// FormRelativePosition: # Discussion
	FormRelativePosition Form = 'r'<<24 | 'e'<<16 | 'l'<<8 | 'e' // 'rele'
	// FormTest: # Discussion
	FormTest Form = 't'<<24 | 'e'<<16 | 's'<<8 | 't' // 'test'
	// FormUniqueID: Specifies a value that uniquely identifies an object within its container or across an application.
	FormUniqueID Form = 'I'<<24 | 'D'<<16 | ' '<<8 | ' ' // 'ID  '
)

func (e Form) String() string {
	switch e {
	case FormAbsolutePosition:
		return "FormAbsolutePosition"
	case FormName:
		return "FormName"
	case FormPropertyID:
		return "FormPropertyID"
	case FormRange:
		return "FormRange"
	case FormRelativePosition:
		return "FormRelativePosition"
	case FormTest:
		return "FormTest"
	case FormUniqueID:
		return "FormUniqueID"
	default:
		return fmt.Sprintf("Form(%d)", e)
	}
}

type Fsm int

const (
	FsmBadFFSNameErr Fsm = -433
	FsmBadFSDLenErr Fsm = -434
	FsmBadFSDVersionErr Fsm = -436
	FsmBusyFFSErr Fsm = -432
	FsmDuplicateFSIDErr Fsm = -435
	FsmFFSNotFoundErr Fsm = -431
	FsmNoAlternateStackErr Fsm = -437
	FsmUnknownFSMMessageErr Fsm = -438
)

func (e Fsm) String() string {
	switch e {
	case FsmBadFFSNameErr:
		return "FsmBadFFSNameErr"
	case FsmBadFSDLenErr:
		return "FsmBadFSDLenErr"
	case FsmBadFSDVersionErr:
		return "FsmBadFSDVersionErr"
	case FsmBusyFFSErr:
		return "FsmBusyFFSErr"
	case FsmDuplicateFSIDErr:
		return "FsmDuplicateFSIDErr"
	case FsmFFSNotFoundErr:
		return "FsmFFSNotFoundErr"
	case FsmNoAlternateStackErr:
		return "FsmNoAlternateStackErr"
	case FsmUnknownFSMMessageErr:
		return "FsmUnknownFSMMessageErr"
	default:
		return fmt.Sprintf("Fsm(%d)", e)
	}
}

type GenericDocumentIconResource int

const (
	AppleMenuFolderIconResource GenericDocumentIconResource = -3982
	DesktopIconResource GenericDocumentIconResource = -3992
	FloppyIconResource GenericDocumentIconResource = -3998
	GenericApplicationIconResource GenericDocumentIconResource = -3996
	GenericCDROMIconResource GenericDocumentIconResource = -3987
	GenericDeskAccessoryIconResource GenericDocumentIconResource = -3991
	GenericDocumentIconResourceValue GenericDocumentIconResource = -4000
	GenericEditionFileIconResource GenericDocumentIconResource = -3989
	GenericExtensionIconResource GenericDocumentIconResource = -16415
	GenericFileServerIconResource GenericDocumentIconResource = -3972
	GenericFolderIconResource GenericDocumentIconResource = -3999
	GenericHardDiskIconResource GenericDocumentIconResource = -3995
	GenericMoverObjectIconResource GenericDocumentIconResource = -3969
	GenericPreferencesIconResource GenericDocumentIconResource = -3971
	GenericQueryDocumentIconResource GenericDocumentIconResource = -16506
	GenericRAMDiskIconResource GenericDocumentIconResource = -3988
	GenericStationeryIconResource GenericDocumentIconResource = -3985
	GenericSuitcaseIconResource GenericDocumentIconResource = -3970
	OpenFolderIconResource GenericDocumentIconResource = -3997
	PrivateFolderIconResource GenericDocumentIconResource = -3994
	SystemFolderIconResource GenericDocumentIconResource = -3983
	TrashIconResource GenericDocumentIconResource = -3993
)

func (e GenericDocumentIconResource) String() string {
	switch e {
	case AppleMenuFolderIconResource:
		return "AppleMenuFolderIconResource"
	case DesktopIconResource:
		return "DesktopIconResource"
	case FloppyIconResource:
		return "FloppyIconResource"
	case GenericApplicationIconResource:
		return "GenericApplicationIconResource"
	case GenericCDROMIconResource:
		return "GenericCDROMIconResource"
	case GenericDeskAccessoryIconResource:
		return "GenericDeskAccessoryIconResource"
	case GenericDocumentIconResourceValue:
		return "GenericDocumentIconResourceValue"
	case GenericEditionFileIconResource:
		return "GenericEditionFileIconResource"
	case GenericExtensionIconResource:
		return "GenericExtensionIconResource"
	case GenericFileServerIconResource:
		return "GenericFileServerIconResource"
	case GenericFolderIconResource:
		return "GenericFolderIconResource"
	case GenericHardDiskIconResource:
		return "GenericHardDiskIconResource"
	case GenericMoverObjectIconResource:
		return "GenericMoverObjectIconResource"
	case GenericPreferencesIconResource:
		return "GenericPreferencesIconResource"
	case GenericQueryDocumentIconResource:
		return "GenericQueryDocumentIconResource"
	case GenericRAMDiskIconResource:
		return "GenericRAMDiskIconResource"
	case GenericStationeryIconResource:
		return "GenericStationeryIconResource"
	case GenericSuitcaseIconResource:
		return "GenericSuitcaseIconResource"
	case OpenFolderIconResource:
		return "OpenFolderIconResource"
	case PrivateFolderIconResource:
		return "PrivateFolderIconResource"
	case SystemFolderIconResource:
		return "SystemFolderIconResource"
	case TrashIconResource:
		return "TrashIconResource"
	default:
		return fmt.Sprintf("GenericDocumentIconResource(%d)", e)
	}
}

const (
	// GestaltDupSelectorErr: Specifies you tried to add an entry that already existed.
	GestaltDupSelectorErr int = -5552
	// GestaltLocationErr: Specifies the gestalt function ptr was not in the system heap.
	GestaltLocationErr int = -5553
	// GestaltUndefSelectorErr: Specifies an undefined selector was passed to the Gestalt Manager.
	GestaltUndefSelectorErr int = -5551
	// GestaltUnknownErr: Specifies an unknown error.
	GestaltUnknownErr int = -5550
)

type Hm int

const (
	HmBalloonAborted Hm = -853
	HmCloseViewActive Hm = -863
	HmHelpDisabled Hm = -850
	HmHelpManagerNotInited Hm = -855
	HmNoBalloonUp Hm = -862
	HmOperationUnsupported Hm = -861
	HmSameAsLastBalloon Hm = -854
	HmSkippedBalloon Hm = -857
	HmUnknownHelpType Hm = -859
	HmWrongVersion Hm = -858
)

func (e Hm) String() string {
	switch e {
	case HmBalloonAborted:
		return "HmBalloonAborted"
	case HmCloseViewActive:
		return "HmCloseViewActive"
	case HmHelpDisabled:
		return "HmHelpDisabled"
	case HmHelpManagerNotInited:
		return "HmHelpManagerNotInited"
	case HmNoBalloonUp:
		return "HmNoBalloonUp"
	case HmOperationUnsupported:
		return "HmOperationUnsupported"
	case HmSameAsLastBalloon:
		return "HmSameAsLastBalloon"
	case HmSkippedBalloon:
		return "HmSkippedBalloon"
	case HmUnknownHelpType:
		return "HmUnknownHelpType"
	case HmWrongVersion:
		return "HmWrongVersion"
	default:
		return fmt.Sprintf("Hm(%d)", e)
	}
}

type Hr int

const (
	HrHTMLRenderingLibNotInstalledErr Hr = -5360
	HrMiscellaneousExceptionErr Hr = -5361
	HrURLNotHandledErr Hr = -5363
	HrUnableToResizeHandleErr Hr = -5362
)

func (e Hr) String() string {
	switch e {
	case HrHTMLRenderingLibNotInstalledErr:
		return "HrHTMLRenderingLibNotInstalledErr"
	case HrMiscellaneousExceptionErr:
		return "HrMiscellaneousExceptionErr"
	case HrURLNotHandledErr:
		return "HrURLNotHandledErr"
	case HrUnableToResizeHandleErr:
		return "HrUnableToResizeHandleErr"
	default:
		return fmt.Sprintf("Hr(%d)", e)
	}
}

type IMemFullErr int

const (
	IIOAbort IMemFullErr = -27
	IMemFullErrValue IMemFullErr = -108
)

func (e IMemFullErr) String() string {
	switch e {
	case IIOAbort:
		return "IIOAbort"
	case IMemFullErrValue:
		return "IMemFullErrValue"
	default:
		return fmt.Sprintf("IMemFullErr(%d)", e)
	}
}

type InternalComponentErr int

const (
	CantReceiveFromSynthesizerOSErr InternalComponentErr = -2073
	CantSendToSynthesizerOSErr InternalComponentErr = -2072
	IllegalChannelOSErr InternalComponentErr = -2076
	IllegalControllerOSErr InternalComponentErr = -2080
	IllegalInstrumentOSErr InternalComponentErr = -2079
	IllegalKnobOSErr InternalComponentErr = -2077
	IllegalKnobValueOSErr InternalComponentErr = -2078
	IllegalNoteChannelOSErr InternalComponentErr = -2084
	IllegalPartOSErr InternalComponentErr = -2075
	IllegalVoiceAllocationOSErr InternalComponentErr = -2074
	InternalComponentErrValue InternalComponentErr = -2070
	MidiManagerAbsentOSErr InternalComponentErr = -2081
	NoExportProcAvailableErr InternalComponentErr = -2089
	NotImplementedMusicOSErr InternalComponentErr = -2071
	NoteChannelNotAllocatedOSErr InternalComponentErr = -2085
	SynthesizerNotRespondingOSErr InternalComponentErr = -2082
	SynthesizerOSErr InternalComponentErr = -2083
	TuneParseOSErr InternalComponentErr = -2087
	TunePlayerFullOSErr InternalComponentErr = -2086
	VideoOutputInUseErr InternalComponentErr = -2090
)

func (e InternalComponentErr) String() string {
	switch e {
	case CantReceiveFromSynthesizerOSErr:
		return "CantReceiveFromSynthesizerOSErr"
	case CantSendToSynthesizerOSErr:
		return "CantSendToSynthesizerOSErr"
	case IllegalChannelOSErr:
		return "IllegalChannelOSErr"
	case IllegalControllerOSErr:
		return "IllegalControllerOSErr"
	case IllegalInstrumentOSErr:
		return "IllegalInstrumentOSErr"
	case IllegalKnobOSErr:
		return "IllegalKnobOSErr"
	case IllegalKnobValueOSErr:
		return "IllegalKnobValueOSErr"
	case IllegalNoteChannelOSErr:
		return "IllegalNoteChannelOSErr"
	case IllegalPartOSErr:
		return "IllegalPartOSErr"
	case IllegalVoiceAllocationOSErr:
		return "IllegalVoiceAllocationOSErr"
	case InternalComponentErrValue:
		return "InternalComponentErrValue"
	case MidiManagerAbsentOSErr:
		return "MidiManagerAbsentOSErr"
	case NoExportProcAvailableErr:
		return "NoExportProcAvailableErr"
	case NotImplementedMusicOSErr:
		return "NotImplementedMusicOSErr"
	case NoteChannelNotAllocatedOSErr:
		return "NoteChannelNotAllocatedOSErr"
	case SynthesizerNotRespondingOSErr:
		return "SynthesizerNotRespondingOSErr"
	case SynthesizerOSErr:
		return "SynthesizerOSErr"
	case TuneParseOSErr:
		return "TuneParseOSErr"
	case TunePlayerFullOSErr:
		return "TunePlayerFullOSErr"
	case VideoOutputInUseErr:
		return "VideoOutputInUseErr"
	default:
		return fmt.Sprintf("InternalComponentErr(%d)", e)
	}
}

type InternalScrapErr int

const (
	BadScrapRefErr InternalScrapErr = -4990
	DuplicateScrapFlavorErr InternalScrapErr = -4989
	IllegalScrapFlavorFlagsErr InternalScrapErr = -4997
	IllegalScrapFlavorSizeErr InternalScrapErr = -4999
	IllegalScrapFlavorTypeErr InternalScrapErr = -4998
	InternalScrapErrValue InternalScrapErr = -4988
	NeedClearScrapErr InternalScrapErr = -100
	NilScrapFlavorDataErr InternalScrapErr = -4994
	NoScrapPromiseKeeperErr InternalScrapErr = -4993
	ProcessStateIncorrectErr InternalScrapErr = -4991
	ScrapFlavorFlagsMismatchErr InternalScrapErr = -4995
	ScrapFlavorNotFoundErr InternalScrapErr = -102
	ScrapFlavorSizeMismatchErr InternalScrapErr = -4996
	ScrapPromiseNotKeptErr InternalScrapErr = -4992
)

func (e InternalScrapErr) String() string {
	switch e {
	case BadScrapRefErr:
		return "BadScrapRefErr"
	case DuplicateScrapFlavorErr:
		return "DuplicateScrapFlavorErr"
	case IllegalScrapFlavorFlagsErr:
		return "IllegalScrapFlavorFlagsErr"
	case IllegalScrapFlavorSizeErr:
		return "IllegalScrapFlavorSizeErr"
	case IllegalScrapFlavorTypeErr:
		return "IllegalScrapFlavorTypeErr"
	case InternalScrapErrValue:
		return "InternalScrapErrValue"
	case NeedClearScrapErr:
		return "NeedClearScrapErr"
	case NilScrapFlavorDataErr:
		return "NilScrapFlavorDataErr"
	case NoScrapPromiseKeeperErr:
		return "NoScrapPromiseKeeperErr"
	case ProcessStateIncorrectErr:
		return "ProcessStateIncorrectErr"
	case ScrapFlavorFlagsMismatchErr:
		return "ScrapFlavorFlagsMismatchErr"
	case ScrapFlavorNotFoundErr:
		return "ScrapFlavorNotFoundErr"
	case ScrapFlavorSizeMismatchErr:
		return "ScrapFlavorSizeMismatchErr"
	case ScrapPromiseNotKeptErr:
		return "ScrapPromiseNotKeptErr"
	default:
		return fmt.Sprintf("InternalScrapErr(%d)", e)
	}
}

type InvalidComponentID int

const (
	ComponentDontRegister InvalidComponentID = -3003
	// ComponentNotCaptured: This component has not been captured.
	ComponentNotCaptured InvalidComponentID = -3002
	// InvalidComponentIDValue: Invalid component ID.
	InvalidComponentIDValue InvalidComponentID = -3000
	RetryComponentRegistrationErr InvalidComponentID = -3005
	UnresolvedComponentDLLErr InvalidComponentID = -3004
	// ValidInstancesExist: This component has open connections.
	ValidInstancesExist InvalidComponentID = -3001
)

func (e InvalidComponentID) String() string {
	switch e {
	case ComponentDontRegister:
		return "ComponentDontRegister"
	case ComponentNotCaptured:
		return "ComponentNotCaptured"
	case InvalidComponentIDValue:
		return "InvalidComponentIDValue"
	case RetryComponentRegistrationErr:
		return "RetryComponentRegistrationErr"
	case UnresolvedComponentDLLErr:
		return "UnresolvedComponentDLLErr"
	case ValidInstancesExist:
		return "ValidInstancesExist"
	default:
		return fmt.Sprintf("InvalidComponentID(%d)", e)
	}
}

type InvalidIconRefErr int

const (
	InvalidIconRefErrValue InvalidIconRefErr = -2580
	NoIconDataAvailableErr InvalidIconRefErr = -2582
	NoSuchIconErr InvalidIconRefErr = -2581
)

func (e InvalidIconRefErr) String() string {
	switch e {
	case InvalidIconRefErrValue:
		return "InvalidIconRefErrValue"
	case NoIconDataAvailableErr:
		return "NoIconDataAvailableErr"
	case NoSuchIconErr:
		return "NoSuchIconErr"
	default:
		return fmt.Sprintf("InvalidIconRefErr(%d)", e)
	}
}

type InvalidTranslationPathErr int

const (
	BadTranslationSpecErr InvalidTranslationPathErr = -3031
	CouldNotParseSourceFileErr InvalidTranslationPathErr = -3026
	InvalidTranslationPathErrValue InvalidTranslationPathErr = -3025
	NoPrefAppErr InvalidTranslationPathErr = -3032
	NoTranslationPathErr InvalidTranslationPathErr = -3030
)

func (e InvalidTranslationPathErr) String() string {
	switch e {
	case BadTranslationSpecErr:
		return "BadTranslationSpecErr"
	case CouldNotParseSourceFileErr:
		return "CouldNotParseSourceFileErr"
	case InvalidTranslationPathErrValue:
		return "InvalidTranslationPathErrValue"
	case NoPrefAppErr:
		return "NoPrefAppErr"
	case NoTranslationPathErr:
		return "NoTranslationPathErr"
	default:
		return fmt.Sprintf("InvalidTranslationPathErr(%d)", e)
	}
}

type KAE uint

const (
	KAEActivate KAE = 'a'<<24 | 'c'<<16 | 't'<<8 | 'v' // 'actv'
	// KAEAlwaysInteract: # Discussion
	KAEAlwaysInteract KAE = 48
	// KAEAnswer: Event that is a reply Apple event.
	KAEAnswer KAE = 'a'<<24 | 'n'<<16 | 's'<<8 | 'r' // 'ansr'
	KAEApplicationClass KAE = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'l' // 'appl'
	// KAEApplicationDied: Event sent by the Process Manager to an application that launched another application when the launched application quits or terminates.
	KAEApplicationDied KAE = 'o'<<24 | 'b'<<16 | 'i'<<8 | 't' // 'obit'
	KAEAsk KAE = 'a'<<24 | 's'<<16 | 'k'<<8 | ' ' // 'ask '
	KAEAutoDown KAE = 'a'<<24 | 'u'<<16 | 't'<<8 | 'o' // 'auto'
	KAEBefore KAE = 'b'<<24 | 'e'<<16 | 'f'<<8 | 'o' // 'befo'
	KAEBeginTransaction KAE = 'b'<<24 | 'e'<<16 | 'g'<<8 | 'i' // 'begi'
	KAEBeginning KAE = 'b'<<24 | 'g'<<16 | 'n'<<8 | 'g' // 'bgng'
	// KAEBeginsWith: The value of `operand1` begins with the value of `operand2` (for example, the string `"operand"` begins with the string `"opera"`).
	KAEBeginsWith KAE = 'b'<<24 | 'g'<<16 | 'w'<<8 | 't' // 'bgwt'
	KAEBold KAE = 'b'<<24 | 'o'<<16 | 'l'<<8 | 'd' // 'bold'
	// KAECanInteract: # Discussion
	KAECanInteract KAE = 32
	// KAECanSwitchLayer: # Discussion
	KAECanSwitchLayer KAE = 64
	KAECaseSensEquals KAE = 'c'<<24 | 's'<<16 | 'e'<<8 | 'q' // 'cseq'
	KAECentered KAE = 'c'<<24 | 'e'<<16 | 'n'<<8 | 't' // 'cent'
	KAEChangeView KAE = 'v'<<24 | 'i'<<16 | 'e'<<8 | 'w' // 'view'
	KAEClone KAE = 'c'<<24 | 'l'<<16 | 'o'<<8 | 'n' // 'clon'
	KAEClose KAE = 'c'<<24 | 'l'<<16 | 'o'<<8 | 's' // 'clos'
	KAECommandClass KAE = 'c'<<24 | 'm'<<16 | 'n'<<8 | 'd' // 'cmnd'
	KAECondensed KAE = 'c'<<24 | 'o'<<16 | 'n'<<8 | 'd' // 'cond'
	// KAEContains: The value of `operand1` contains the value of `operand2 `(for example, the string `"operand"` contains the string `"era"`).
	KAEContains KAE = 'c'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'cont'
	KAECopy KAE = 'c'<<24 | 'o'<<16 | 'p'<<8 | 'y' // 'copy'
	// KAECoreSuite: An Apple event in the Standard Suite.
	KAECoreSuite KAE = 'c'<<24 | 'o'<<16 | 'r'<<8 | 'e' // 'core'
	KAECountElements KAE = 'c'<<24 | 'n'<<16 | 't'<<8 | 'e' // 'cnte'
	KAECreateElement KAE = 'c'<<24 | 'r'<<16 | 'e'<<8 | 'l' // 'crel'
	KAECreatePublisher KAE = 'c'<<24 | 'p'<<16 | 'u'<<8 | 'b' // 'cpub'
	KAECut KAE = 'c'<<24 | 'u'<<16 | 't'<<8 | ' ' // 'cut '
	// KAEDataArray: Array items consist of data of the same size and same type, and are aligned on word boundaries.
	KAEDataArray KAE = 0
	KAEDeactivate KAE = 'd'<<24 | 'a'<<16 | 'c'<<8 | 't' // 'dact'
	KAEDelete KAE = 'd'<<24 | 'e'<<16 | 'l'<<8 | 'o' // 'delo'
	// KAEDescArray: Array items consist of descriptors of different descriptor types with data of variable size.
	KAEDescArray KAE = 3
	// KAEDirectCall: The source of the Apple event is a direct call that bypassed the PPC Toolbox.
	KAEDirectCall KAE = 1
	KAEDiskEvent KAE = 'd'<<24 | 'i'<<16 | 's'<<8 | 'k' // 'disk'
	KAEDoNotAutomaticallyAddAnnotationsToEvent KAE = 65536
	KAEDoObjectsExist KAE = 'd'<<24 | 'o'<<16 | 'e'<<8 | 'x' // 'doex'
	KAEDoScript KAE = 'd'<<24 | 'o'<<16 | 's'<<8 | 'c' // 'dosc'
	// KAEDontExecute: # Discussion
	KAEDontExecute KAE = 8192
	// KAEDontReconnect: # Discussion
	KAEDontReconnect KAE = 128
	// KAEDontRecord: # Discussion
	KAEDontRecord KAE = 4096
	KAEDown KAE = 'd'<<24 | 'o'<<16 | 'w'<<8 | 'n' // 'down'
	KAEDrag KAE = 'd'<<24 | 'r'<<16 | 'a'<<8 | 'g' // 'drag'
	KAEDuplicateSelection KAE = 's'<<24 | 'd'<<16 | 'u'<<8 | 'p' // 'sdup'
	KAEEditGraphic KAE = 'e'<<24 | 'd'<<16 | 'i'<<8 | 't' // 'edit'
	KAEEmptyTrash KAE = 'e'<<24 | 'm'<<16 | 'p'<<8 | 't' // 'empt'
	KAEEnd KAE = 'e'<<24 | 'n'<<16 | 'd'<<8 | ' ' // 'end '
	KAEEndTransaction KAE = 'e'<<24 | 'n'<<16 | 'd'<<8 | 't' // 'endt'
	// KAEEndsWith: The value of `operand1` ends with the value of `operand2` (for example, the string `"operand"` ends with the string `"and"`).
	KAEEndsWith KAE = 'e'<<24 | 'n'<<16 | 'd'<<8 | 's' // 'ends'
	// KAEEquals: The value of `operand1` is equal to the value of `operand2`
	KAEEquals KAE = '='<<24 | ' '<<16 | ' '<<8 | ' ' // '=   '
	KAEExpanded KAE = 'p'<<24 | 'e'<<16 | 'x'<<8 | 'p' // 'pexp'
	KAEFast KAE = 'f'<<24 | 'a'<<16 | 's'<<8 | 't' // 'fast'
	// KAEFinderEvents: An event that the Finder accepts.
	KAEFinderEvents KAE = 'F'<<24 | 'N'<<16 | 'D'<<8 | 'R' // 'FNDR'
	KAEFormulaProtect KAE = 'f'<<24 | 'p'<<16 | 'r'<<8 | 'o' // 'fpro'
	KAEFullyJustified KAE = 'f'<<24 | 'u'<<16 | 'l'<<8 | 'l' // 'full'
	KAEGetClassInfo KAE = 'q'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'qobj'
	KAEGetData KAE = 'g'<<24 | 'e'<<16 | 't'<<8 | 'd' // 'getd'
	KAEGetDataSize KAE = 'd'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'dsiz'
	KAEGetEventInfo KAE = 'g'<<24 | 't'<<16 | 'e'<<8 | 'i' // 'gtei'
	KAEGetInfoSelection KAE = 's'<<24 | 'i'<<16 | 'n'<<8 | 'f' // 'sinf'
	KAEGetPrivilegeSelection KAE = 's'<<24 | 'p'<<16 | 'r'<<8 | 'v' // 'sprv'
	KAEGetSuiteInfo KAE = 'g'<<24 | 't'<<16 | 's'<<8 | 'i' // 'gtsi'
	// KAEGreaterThan: The value of `operand1` is greater than the value of `operand2`.
	KAEGreaterThan KAE = '>'<<24 | ' '<<16 | ' '<<8 | ' ' // '>   '
	// KAEGreaterThanEquals: The value of `operand1` is greater than or equal to the value of `operand2`.
	KAEGreaterThanEquals KAE = '>'<<24 | '='<<16 | ' '<<8 | ' ' // '>=  '
	KAEGrow KAE = 'g'<<24 | 'r'<<16 | 'o'<<8 | 'w' // 'grow'
	// KAEHTTPProxyHostAttr: A value of type `typeChar` or `typeUTF8Text`.
	KAEHTTPProxyHostAttr KAE = 'x'<<24 | 'h'<<16 | 't'<<8 | 'h' // 'xhth'
	// KAEHTTPProxyPortAttr: A value of type `typeSInt32`.
	KAEHTTPProxyPortAttr KAE = 'x'<<24 | 'h'<<16 | 't'<<8 | 'p' // 'xhtp'
	KAEHandleSimpleRanges KAE = 32
	KAEHiQuality KAE = 'h'<<24 | 'i'<<16 | 'q'<<8 | 'u' // 'hiqu'
	KAEHidden KAE = 'h'<<24 | 'i'<<16 | 'd'<<8 | 'n' // 'hidn'
	KAEHighLevel KAE = 'h'<<24 | 'i'<<16 | 'g'<<8 | 'h' // 'high'
	// KAEHighPriority: The Apple Event Manager posts the event at the beginning of the event queue of the server process.
	KAEHighPriority KAE = 1
	// KAEIDoMarking: The application provides marking callback functions.
	KAEIDoMarking KAE = 4
	// KAEIDoMinimum: The application does not handle whose tests or provide marking callbacks.
	KAEIDoMinimum KAE = 0
	// KAEIDoWhose: The application supports whose tests (supports key form `formWhose`).
	KAEIDoWhose KAE = 1
	KAEISWebStarSuite KAE = 1465341885
	KAEImageGraphic KAE = 'i'<<24 | 'm'<<16 | 'g'<<8 | 'r' // 'imgr'
	KAEInfo KAE = 11
	KAEInternetSuite KAE = 'g'<<24 | 'u'<<16 | 'r'<<8 | 'l' // 'gurl'
	KAEIsUniform KAE = 'i'<<24 | 's'<<16 | 'u'<<8 | 'n' // 'isun'
	KAEItalic KAE = 'i'<<24 | 't'<<16 | 'a'<<8 | 'l' // 'ital'
	KAEKeyClass KAE = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'c' // 'keyc'
	// KAEKeyDescArray: Array items consist of keyword-specified descriptors with different keywords, different descriptor types, and data of variable size.
	KAEKeyDescArray KAE = 4
	KAEKeyDown KAE = 'k'<<24 | 'd'<<16 | 'w'<<8 | 'n' // 'kdwn'
	KAELeftJustified KAE = 'l'<<24 | 'e'<<16 | 'f'<<8 | 't' // 'left'
	KAELessThan KAE = '<'<<24 | ' '<<16 | ' '<<8 | ' ' // '<   '
	// KAELessThanEquals: The value of `operand1` is less than or equal to the value of `operand2`.
	KAELessThanEquals KAE = '<'<<24 | '='<<16 | ' '<<8 | ' ' // '<=  '
	// KAELocalProcess: The source application is another process on the same computer as the target application.
	KAELocalProcess KAE = 3
	KAELogOut KAE = 'l'<<24 | 'o'<<16 | 'g'<<8 | 'o' // 'logo'
	KAELowercase KAE = 'l'<<24 | 'o'<<16 | 'w'<<8 | 'c' // 'lowc'
	KAEMain KAE = 0
	KAEMakeObjectsVisible KAE = 'm'<<24 | 'v'<<16 | 'i'<<8 | 's' // 'mvis'
	KAEMenuClass KAE = 'm'<<24 | 'e'<<16 | 'n'<<8 | 'u' // 'menu'
	KAEMenuSelect KAE = 'm'<<24 | 'h'<<16 | 'i'<<8 | 't' // 'mhit'
	KAEMiscStandards KAE = 'm'<<24 | 'i'<<16 | 's'<<8 | 'c' // 'misc'
	KAEModifiable KAE = 'm'<<24 | 'o'<<16 | 'd'<<8 | 'f' // 'modf'
	KAEMouseClass KAE = 'm'<<24 | 'o'<<16 | 'u'<<8 | 's' // 'mous'
	KAEMouseDown KAE = 'm'<<24 | 'd'<<16 | 'w'<<8 | 'n' // 'mdwn'
	KAEMouseDownInBack KAE = 'm'<<24 | 'd'<<16 | 'b'<<8 | 'k' // 'mdbk'
	KAEMove KAE = 'm'<<24 | 'o'<<16 | 'v'<<8 | 'e' // 'move'
	KAEMoved KAE = 'm'<<24 | 'o'<<16 | 'v'<<8 | 'e' // 'move'
	KAENavigationKey KAE = 'n'<<24 | 'a'<<16 | 'v'<<8 | 'e' // 'nave'
	// KAENeverInteract: # Discussion
	KAENeverInteract KAE = 16
	KAENo KAE = 'n'<<24 | 'o'<<16 | ' '<<8 | ' ' // 'no  '
	KAENoArrow KAE = 'a'<<24 | 'r'<<16 | 'n'<<8 | 'o' // 'arno'
	// KAENoReply: The reply preference—your application does not want a reply Apple event.
	KAENoReply KAE = 1
	KAENonmodifiable KAE = 'n'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'nmod'
	// KAENormalPriority: The Apple Event Manager posts the event at the end of the event queue of the server process and the server processes the Apple event as soon as it has the opportunity.
	KAENormalPriority KAE = 0
	// KAENotifyRecording: # Discussion
	KAENotifyRecording KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'r' // 'recr'
	// KAENotifyStartRecording: An event that notifies an application that recording has been turned on.
	KAENotifyStartRecording KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | '1' // 'rec1'
	// KAENotifyStopRecording: An event that notifies an application that recording has been turned off.
	KAENotifyStopRecording KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | '0' // 'rec0'
	KAENullEvent KAE = 'n'<<24 | 'u'<<16 | 'l'<<8 | 'l' // 'null'
	KAEOSAXSizeResource KAE = 'o'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'osiz'
	KAEOpen KAE = 'o'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'odoc'
	// KAEOpenApplication: Event that launches an application.
	KAEOpenApplication KAE = 'o'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'oapp'
	// KAEOpenContents: # Discussion
	KAEOpenContents KAE = 'o'<<24 | 'c'<<16 | 'o'<<8 | 'n' // 'ocon'
	// KAEOpenDocuments: # Discussion
	KAEOpenDocuments KAE = 'o'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'odoc'
	KAEOpenSelection KAE = 's'<<24 | 'o'<<16 | 'p'<<8 | 'e' // 'sope'
	KAEOutline KAE = 'o'<<24 | 'u'<<16 | 't'<<8 | 'l' // 'outl'
	// KAEPackedArray: Array items consist of data of the same size and same type, and are packed without regard for word boundaries.
	KAEPackedArray KAE = 1
	KAEPageSetup KAE = 'p'<<24 | 'g'<<16 | 's'<<8 | 'u' // 'pgsu'
	KAEPassSubDescs KAE = 8
	KAEPaste KAE = 'p'<<24 | 'a'<<16 | 's'<<8 | 't' // 'past'
	KAEPlain KAE = 'p'<<24 | 'l'<<16 | 'a'<<8 | 'n' // 'plan'
	KAEPrint KAE = 'p'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'pdoc'
	// KAEPrintDocuments: Event that provides an application with a list of documents to print.
	KAEPrintDocuments KAE = 'p'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'pdoc'
	KAEPrintSelection KAE = 's'<<24 | 'p'<<16 | 'r'<<8 | 'i' // 'spri'
	KAEPrintWindow KAE = 'p'<<24 | 'w'<<16 | 'i'<<8 | 'n' // 'pwin'
	// KAEProcessNonReplyEvents: Allow processing of non-reply Apple events while awaiting a synchronous Apple event reply (you specified `kAEWaitReply` for the reply preference).
	KAEProcessNonReplyEvents KAE = 32768
	KAEPromise KAE = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'm' // 'prom'
	KAEPutAwaySelection KAE = 's'<<24 | 'p'<<16 | 'u'<<8 | 't' // 'sput'
	KAEQDAdMax KAE = 'a'<<24 | 'd'<<16 | 'm'<<8 | 'x' // 'admx'
	KAEQDAdMin KAE = 'a'<<24 | 'd'<<16 | 'm'<<8 | 'n' // 'admn'
	KAEQDAddOver KAE = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'o' // 'addo'
	KAEQDAddPin KAE = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'p' // 'addp'
	KAEQDBic KAE = 'b'<<24 | 'i'<<16 | 'c'<<8 | ' ' // 'bic '
	KAEQDBlend KAE = 'b'<<24 | 'l'<<16 | 'n'<<8 | 'd' // 'blnd'
	KAEQDCopy KAE = 'c'<<24 | 'p'<<16 | 'y'<<8 | ' ' // 'cpy '
	KAEQDNotBic KAE = 'n'<<24 | 'b'<<16 | 'i'<<8 | 'c' // 'nbic'
	KAEQDNotCopy KAE = 'n'<<24 | 'c'<<16 | 'p'<<8 | 'y' // 'ncpy'
	KAEQDNotOr KAE = 'n'<<24 | 't'<<16 | 'o'<<8 | 'r' // 'ntor'
	KAEQDNotXor KAE = 'n'<<24 | 'x'<<16 | 'o'<<8 | 'r' // 'nxor'
	KAEQDOr KAE = 'o'<<24 | 'r'<<16 | ' '<<8 | ' ' // 'or  '
	KAEQDSubOver KAE = 's'<<24 | 'u'<<16 | 'b'<<8 | 'o' // 'subo'
	KAEQDSubPin KAE = 's'<<24 | 'u'<<16 | 'b'<<8 | 'p' // 'subp'
	KAEQDSupplementalSuite KAE = 'q'<<24 | 'd'<<16 | 's'<<8 | 'p' // 'qdsp'
	KAEQDXor KAE = 'x'<<24 | 'o'<<16 | 'r'<<8 | ' ' // 'xor '
	// KAEQueueReply: The reply preference—your application wants a reply Apple event.
	KAEQueueReply KAE = 2
	KAEQuickdrawSuite KAE = 'q'<<24 | 'd'<<16 | 'r'<<8 | 'w' // 'qdrw'
	KAEQuitAll KAE = 'q'<<24 | 'u'<<16 | 'i'<<8 | 'a' // 'quia'
	// KAEQuitApplication: Event that causes the application to quit.
	KAEQuitApplication KAE = 'q'<<24 | 'u'<<16 | 'i'<<8 | 't' // 'quit'
	KAERawKey KAE = 'r'<<24 | 'k'<<16 | 'e'<<8 | 'y' // 'rkey'
	KAEReallyLogOut KAE = 'r'<<24 | 'l'<<16 | 'g'<<8 | 'o' // 'rlgo'
	KAERedo KAE = 'r'<<24 | 'e'<<16 | 'd'<<8 | 'o' // 'redo'
	KAERegular KAE = 'r'<<24 | 'e'<<16 | 'g'<<8 | 'l' // 'regl'
	// KAERemoteProcess: The source application is a process on a remote computer on the network.
	KAERemoteProcess KAE = 4
	// KAEReopenApplication: Event that reopens an application.
	KAEReopenApplication KAE = 'r'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'rapp'
	KAEReplace KAE = 'r'<<24 | 'p'<<16 | 'l'<<8 | 'c' // 'rplc'
	KAERequiredSuite KAE = 'r'<<24 | 'e'<<16 | 'q'<<8 | 'd' // 'reqd'
	KAEResized KAE = 'r'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'rsiz'
	KAEResolveNestedLists KAE = 16
	KAERestart KAE = 'r'<<24 | 'e'<<16 | 's'<<8 | 't' // 'rest'
	KAEResume KAE = 'r'<<24 | 's'<<16 | 'm'<<8 | 'e' // 'rsme'
	KAERevealSelection KAE = 's'<<24 | 'r'<<16 | 'e'<<8 | 'v' // 'srev'
	KAERevert KAE = 'r'<<24 | 'v'<<16 | 'r'<<8 | 't' // 'rvrt'
	KAERightJustified KAE = 'r'<<24 | 'g'<<16 | 'h'<<8 | 't' // 'rght'
	// KAESameProcess: The source of the Apple event is the same application that received the event (the target application and the source application are the same).
	KAESameProcess KAE = 2
	KAESave KAE = 's'<<24 | 'a'<<16 | 'v'<<8 | 'e' // 'save'
	KAEScrapEvent KAE = 's'<<24 | 'c'<<16 | 'r'<<8 | 'p' // 'scrp'
	KAEScriptingSizeResource KAE = 's'<<24 | 'c'<<16 | 's'<<8 | 'z' // 'scsz'
	KAESelect KAE = 's'<<24 | 'l'<<16 | 'c'<<8 | 't' // 'slct'
	KAESetData KAE = 's'<<24 | 'e'<<16 | 't'<<8 | 'd' // 'setd'
	KAESetPosition KAE = 'p'<<24 | 'o'<<16 | 's'<<8 | 'n' // 'posn'
	KAEShadow KAE = 's'<<24 | 'h'<<16 | 'a'<<8 | 'd' // 'shad'
	KAESharing KAE = 13
	KAEShowClipboard KAE = 's'<<24 | 'h'<<16 | 'c'<<8 | 'l' // 'shcl'
	// KAEShowPreferences: # Discussion
	KAEShowPreferences KAE = 'p'<<24 | 'r'<<16 | 'e'<<8 | 'f' // 'pref'
	KAEShowRestartDialog KAE = 'r'<<24 | 'r'<<16 | 's'<<8 | 't' // 'rrst'
	KAEShowShutdownDialog KAE = 'r'<<24 | 's'<<16 | 'd'<<8 | 'n' // 'rsdn'
	KAEShutDown KAE = 's'<<24 | 'h'<<16 | 'u'<<8 | 't' // 'shut'
	KAESleep KAE = 's'<<24 | 'l'<<16 | 'e'<<8 | 'p' // 'slep'
	KAESmallCaps KAE = 's'<<24 | 'm'<<16 | 'c'<<8 | 'p' // 'smcp'
	KAESocks4Protocol KAE = 4
	KAESocks5Protocol KAE = 5
	KAESocksHostAttr KAE = 'x'<<24 | 's'<<16 | 'h'<<8 | 's' // 'xshs'
	KAESocksPasswordAttr KAE = 'x'<<24 | 's'<<16 | 'h'<<8 | 'w' // 'xshw'
	KAESocksPortAttr KAE = 'x'<<24 | 's'<<16 | 'h'<<8 | 'p' // 'xshp'
	KAESocksProxyAttr KAE = 'x'<<24 | 's'<<16 | 'o'<<8 | 'k' // 'xsok'
	KAESocksUserAttr KAE = 'x'<<24 | 's'<<16 | 'h'<<8 | 'u' // 'xshu'
	KAESpecialClassProperties KAE = 'c'<<24 | '@'<<16 | '#'<<8 | '!' // 'c@#!'
	// KAEStartRecording: # Discussion
	KAEStartRecording KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'a' // 'reca'
	// KAEStopRecording: # Discussion
	KAEStopRecording KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'c' // 'recc'
	KAEStoppedMoving KAE = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
	KAEStrikethrough KAE = 's'<<24 | 't'<<16 | 'r'<<8 | 'k' // 'strk'
	KAESubscript KAE = 's'<<24 | 'b'<<16 | 's'<<8 | 'c' // 'sbsc'
	KAESuperscript KAE = 's'<<24 | 'p'<<16 | 's'<<8 | 'c' // 'spsc'
	KAESuspend KAE = 's'<<24 | 'u'<<16 | 's'<<8 | 'p' // 'susp'
	KAETableSuite KAE = 't'<<24 | 'b'<<16 | 'l'<<8 | 's' // 'tbls'
	KAETerminologyExtension KAE = 'a'<<24 | 'e'<<16 | 't'<<8 | 'e' // 'aete'
	KAETextSuite KAE = 'T'<<24 | 'E'<<16 | 'X'<<8 | 'T' // 'TEXT'
	KAETransactionTerminated KAE = 't'<<24 | 't'<<16 | 'r'<<8 | 'm' // 'ttrm'
	KAEUnderline KAE = 'u'<<24 | 'n'<<16 | 'd'<<8 | 'l' // 'undl'
	KAEUndo KAE = 'u'<<24 | 'n'<<16 | 'd'<<8 | 'o' // 'undo'
	// KAEUnknownSource: The source of the Apple event is unknown.
	KAEUnknownSource KAE = 0
	KAEUp KAE = 'u'<<24 | 'p'<<16 | ' '<<8 | ' ' // 'up  '
	KAEUpdate KAE = 'u'<<24 | 'p'<<16 | 'd'<<8 | 't' // 'updt'
	// KAEUseHTTPProxyAttr: A value of type `typeBoolean`.
	KAEUseHTTPProxyAttr KAE = 'x'<<24 | 'u'<<16 | 'p'<<8 | 'r' // 'xupr'
	KAEUseRelativeIterators KAE = 64
	KAEUseSocksAttr KAE = 'x'<<24 | 's'<<16 | 'c'<<8 | 's' // 'xscs'
	KAEUserTerminology KAE = 'a'<<24 | 'e'<<16 | 'u'<<8 | 't' // 'aeut'
	KAEVirtualKey KAE = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'c' // 'keyc'
	// KAEWaitReply: # Discussion
	KAEWaitReply KAE = 3
	KAEWakeUpEvent KAE = 'w'<<24 | 'a'<<16 | 'k'<<8 | 'e' // 'wake'
	// KAEWantReceipt: # Discussion
	KAEWantReceipt KAE = 512
	KAEWholeWordEquals KAE = 'w'<<24 | 'w'<<16 | 'e'<<8 | 'q' // 'wweq'
	KAEWindowClass KAE = 'w'<<24 | 'i'<<16 | 'n'<<8 | 'd' // 'wind'
	KAEYes KAE = 'y'<<24 | 'e'<<16 | 's'<<8 | ' ' // 'yes '
	KAEZoom KAE = 'z'<<24 | 'o'<<16 | 'o'<<8 | 'm' // 'zoom'
)

func (e KAE) String() string {
	switch e {
	case KAEActivate:
		return "KAEActivate"
	case KAEAlwaysInteract:
		return "KAEAlwaysInteract"
	case KAEAnswer:
		return "KAEAnswer"
	case KAEApplicationClass:
		return "KAEApplicationClass"
	case KAEApplicationDied:
		return "KAEApplicationDied"
	case KAEAsk:
		return "KAEAsk"
	case KAEAutoDown:
		return "KAEAutoDown"
	case KAEBefore:
		return "KAEBefore"
	case KAEBeginTransaction:
		return "KAEBeginTransaction"
	case KAEBeginning:
		return "KAEBeginning"
	case KAEBeginsWith:
		return "KAEBeginsWith"
	case KAEBold:
		return "KAEBold"
	case KAECanInteract:
		return "KAECanInteract"
	case KAECanSwitchLayer:
		return "KAECanSwitchLayer"
	case KAECaseSensEquals:
		return "KAECaseSensEquals"
	case KAECentered:
		return "KAECentered"
	case KAEChangeView:
		return "KAEChangeView"
	case KAEClone:
		return "KAEClone"
	case KAEClose:
		return "KAEClose"
	case KAECommandClass:
		return "KAECommandClass"
	case KAECondensed:
		return "KAECondensed"
	case KAEContains:
		return "KAEContains"
	case KAECopy:
		return "KAECopy"
	case KAECoreSuite:
		return "KAECoreSuite"
	case KAECountElements:
		return "KAECountElements"
	case KAECreateElement:
		return "KAECreateElement"
	case KAECreatePublisher:
		return "KAECreatePublisher"
	case KAECut:
		return "KAECut"
	case KAEDataArray:
		return "KAEDataArray"
	case KAEDeactivate:
		return "KAEDeactivate"
	case KAEDelete:
		return "KAEDelete"
	case KAEDescArray:
		return "KAEDescArray"
	case KAEDirectCall:
		return "KAEDirectCall"
	case KAEDiskEvent:
		return "KAEDiskEvent"
	case KAEDoNotAutomaticallyAddAnnotationsToEvent:
		return "KAEDoNotAutomaticallyAddAnnotationsToEvent"
	case KAEDoObjectsExist:
		return "KAEDoObjectsExist"
	case KAEDoScript:
		return "KAEDoScript"
	case KAEDontExecute:
		return "KAEDontExecute"
	case KAEDontReconnect:
		return "KAEDontReconnect"
	case KAEDontRecord:
		return "KAEDontRecord"
	case KAEDown:
		return "KAEDown"
	case KAEDrag:
		return "KAEDrag"
	case KAEDuplicateSelection:
		return "KAEDuplicateSelection"
	case KAEEditGraphic:
		return "KAEEditGraphic"
	case KAEEmptyTrash:
		return "KAEEmptyTrash"
	case KAEEnd:
		return "KAEEnd"
	case KAEEndTransaction:
		return "KAEEndTransaction"
	case KAEEndsWith:
		return "KAEEndsWith"
	case KAEEquals:
		return "KAEEquals"
	case KAEExpanded:
		return "KAEExpanded"
	case KAEFast:
		return "KAEFast"
	case KAEFinderEvents:
		return "KAEFinderEvents"
	case KAEFormulaProtect:
		return "KAEFormulaProtect"
	case KAEFullyJustified:
		return "KAEFullyJustified"
	case KAEGetClassInfo:
		return "KAEGetClassInfo"
	case KAEGetData:
		return "KAEGetData"
	case KAEGetDataSize:
		return "KAEGetDataSize"
	case KAEGetEventInfo:
		return "KAEGetEventInfo"
	case KAEGetInfoSelection:
		return "KAEGetInfoSelection"
	case KAEGetPrivilegeSelection:
		return "KAEGetPrivilegeSelection"
	case KAEGetSuiteInfo:
		return "KAEGetSuiteInfo"
	case KAEGreaterThan:
		return "KAEGreaterThan"
	case KAEGreaterThanEquals:
		return "KAEGreaterThanEquals"
	case KAEGrow:
		return "KAEGrow"
	case KAEHTTPProxyHostAttr:
		return "KAEHTTPProxyHostAttr"
	case KAEHTTPProxyPortAttr:
		return "KAEHTTPProxyPortAttr"
	case KAEHiQuality:
		return "KAEHiQuality"
	case KAEHidden:
		return "KAEHidden"
	case KAEHighLevel:
		return "KAEHighLevel"
	case KAEIDoMarking:
		return "KAEIDoMarking"
	case KAEISWebStarSuite:
		return "KAEISWebStarSuite"
	case KAEImageGraphic:
		return "KAEImageGraphic"
	case KAEInfo:
		return "KAEInfo"
	case KAEInternetSuite:
		return "KAEInternetSuite"
	case KAEIsUniform:
		return "KAEIsUniform"
	case KAEItalic:
		return "KAEItalic"
	case KAEKeyClass:
		return "KAEKeyClass"
	case KAEKeyDown:
		return "KAEKeyDown"
	case KAELeftJustified:
		return "KAELeftJustified"
	case KAELessThan:
		return "KAELessThan"
	case KAELessThanEquals:
		return "KAELessThanEquals"
	case KAELogOut:
		return "KAELogOut"
	case KAELowercase:
		return "KAELowercase"
	case KAEMakeObjectsVisible:
		return "KAEMakeObjectsVisible"
	case KAEMenuClass:
		return "KAEMenuClass"
	case KAEMenuSelect:
		return "KAEMenuSelect"
	case KAEMiscStandards:
		return "KAEMiscStandards"
	case KAEModifiable:
		return "KAEModifiable"
	case KAEMouseClass:
		return "KAEMouseClass"
	case KAEMouseDown:
		return "KAEMouseDown"
	case KAEMouseDownInBack:
		return "KAEMouseDownInBack"
	case KAEMove:
		return "KAEMove"
	case KAENavigationKey:
		return "KAENavigationKey"
	case KAENeverInteract:
		return "KAENeverInteract"
	case KAENo:
		return "KAENo"
	case KAENoArrow:
		return "KAENoArrow"
	case KAENonmodifiable:
		return "KAENonmodifiable"
	case KAENotifyRecording:
		return "KAENotifyRecording"
	case KAENotifyStartRecording:
		return "KAENotifyStartRecording"
	case KAENotifyStopRecording:
		return "KAENotifyStopRecording"
	case KAENullEvent:
		return "KAENullEvent"
	case KAEOSAXSizeResource:
		return "KAEOSAXSizeResource"
	case KAEOpen:
		return "KAEOpen"
	case KAEOpenApplication:
		return "KAEOpenApplication"
	case KAEOpenContents:
		return "KAEOpenContents"
	case KAEOpenSelection:
		return "KAEOpenSelection"
	case KAEOutline:
		return "KAEOutline"
	case KAEPageSetup:
		return "KAEPageSetup"
	case KAEPassSubDescs:
		return "KAEPassSubDescs"
	case KAEPaste:
		return "KAEPaste"
	case KAEPlain:
		return "KAEPlain"
	case KAEPrint:
		return "KAEPrint"
	case KAEPrintSelection:
		return "KAEPrintSelection"
	case KAEPrintWindow:
		return "KAEPrintWindow"
	case KAEProcessNonReplyEvents:
		return "KAEProcessNonReplyEvents"
	case KAEPromise:
		return "KAEPromise"
	case KAEPutAwaySelection:
		return "KAEPutAwaySelection"
	case KAEQDAdMax:
		return "KAEQDAdMax"
	case KAEQDAdMin:
		return "KAEQDAdMin"
	case KAEQDAddOver:
		return "KAEQDAddOver"
	case KAEQDAddPin:
		return "KAEQDAddPin"
	case KAEQDBic:
		return "KAEQDBic"
	case KAEQDBlend:
		return "KAEQDBlend"
	case KAEQDCopy:
		return "KAEQDCopy"
	case KAEQDNotBic:
		return "KAEQDNotBic"
	case KAEQDNotCopy:
		return "KAEQDNotCopy"
	case KAEQDNotOr:
		return "KAEQDNotOr"
	case KAEQDNotXor:
		return "KAEQDNotXor"
	case KAEQDOr:
		return "KAEQDOr"
	case KAEQDSubOver:
		return "KAEQDSubOver"
	case KAEQDSubPin:
		return "KAEQDSubPin"
	case KAEQDSupplementalSuite:
		return "KAEQDSupplementalSuite"
	case KAEQDXor:
		return "KAEQDXor"
	case KAEQueueReply:
		return "KAEQueueReply"
	case KAEQuickdrawSuite:
		return "KAEQuickdrawSuite"
	case KAEQuitAll:
		return "KAEQuitAll"
	case KAEQuitApplication:
		return "KAEQuitApplication"
	case KAERawKey:
		return "KAERawKey"
	case KAEReallyLogOut:
		return "KAEReallyLogOut"
	case KAERedo:
		return "KAERedo"
	case KAERegular:
		return "KAERegular"
	case KAEReopenApplication:
		return "KAEReopenApplication"
	case KAEReplace:
		return "KAEReplace"
	case KAERequiredSuite:
		return "KAERequiredSuite"
	case KAEResized:
		return "KAEResized"
	case KAERestart:
		return "KAERestart"
	case KAEResume:
		return "KAEResume"
	case KAERevealSelection:
		return "KAERevealSelection"
	case KAERevert:
		return "KAERevert"
	case KAERightJustified:
		return "KAERightJustified"
	case KAESave:
		return "KAESave"
	case KAEScrapEvent:
		return "KAEScrapEvent"
	case KAEScriptingSizeResource:
		return "KAEScriptingSizeResource"
	case KAESelect:
		return "KAESelect"
	case KAESetData:
		return "KAESetData"
	case KAESetPosition:
		return "KAESetPosition"
	case KAEShadow:
		return "KAEShadow"
	case KAESharing:
		return "KAESharing"
	case KAEShowClipboard:
		return "KAEShowClipboard"
	case KAEShowPreferences:
		return "KAEShowPreferences"
	case KAEShowRestartDialog:
		return "KAEShowRestartDialog"
	case KAEShowShutdownDialog:
		return "KAEShowShutdownDialog"
	case KAEShutDown:
		return "KAEShutDown"
	case KAESleep:
		return "KAESleep"
	case KAESmallCaps:
		return "KAESmallCaps"
	case KAESocks5Protocol:
		return "KAESocks5Protocol"
	case KAESocksHostAttr:
		return "KAESocksHostAttr"
	case KAESocksPasswordAttr:
		return "KAESocksPasswordAttr"
	case KAESocksPortAttr:
		return "KAESocksPortAttr"
	case KAESocksProxyAttr:
		return "KAESocksProxyAttr"
	case KAESocksUserAttr:
		return "KAESocksUserAttr"
	case KAESpecialClassProperties:
		return "KAESpecialClassProperties"
	case KAEStartRecording:
		return "KAEStartRecording"
	case KAEStopRecording:
		return "KAEStopRecording"
	case KAEStoppedMoving:
		return "KAEStoppedMoving"
	case KAEStrikethrough:
		return "KAEStrikethrough"
	case KAESubscript:
		return "KAESubscript"
	case KAESuperscript:
		return "KAESuperscript"
	case KAESuspend:
		return "KAESuspend"
	case KAETableSuite:
		return "KAETableSuite"
	case KAETerminologyExtension:
		return "KAETerminologyExtension"
	case KAETextSuite:
		return "KAETextSuite"
	case KAETransactionTerminated:
		return "KAETransactionTerminated"
	case KAEUnderline:
		return "KAEUnderline"
	case KAEUndo:
		return "KAEUndo"
	case KAEUp:
		return "KAEUp"
	case KAEUpdate:
		return "KAEUpdate"
	case KAEUseHTTPProxyAttr:
		return "KAEUseHTTPProxyAttr"
	case KAEUseSocksAttr:
		return "KAEUseSocksAttr"
	case KAEUserTerminology:
		return "KAEUserTerminology"
	case KAEWakeUpEvent:
		return "KAEWakeUpEvent"
	case KAEWantReceipt:
		return "KAEWantReceipt"
	case KAEWholeWordEquals:
		return "KAEWholeWordEquals"
	case KAEWindowClass:
		return "KAEWindowClass"
	case KAEYes:
		return "KAEYes"
	case KAEZoom:
		return "KAEZoom"
	default:
		return fmt.Sprintf("KAE(%d)", e)
	}
}

type KAEAND uint

const (
	// KAEANDValue: Specifies a logical [AND] operation.
	KAEANDValue KAEAND = 'A'<<24 | 'N'<<16 | 'D'<<8 | ' ' // 'AND '
	// KAEAll: Specifies all the elements in the container.
	KAEAll KAEAND = 'a'<<24 | 'l'<<16 | 'l'<<8 | ' ' // 'all '
	// KAEAny: Specifies a single element chosen at random from the container.
	KAEAny KAEAND = 'a'<<24 | 'n'<<16 | 'y'<<8 | ' ' // 'any '
	// KAEFirst: The first element in the specified container.
	KAEFirst KAEAND = 'f'<<24 | 'i'<<16 | 'r'<<8 | 's' // 'firs'
	// KAELast: Specifies the last element in the container.
	KAELast KAEAND = 'l'<<24 | 'a'<<16 | 's'<<8 | 't' // 'last'
	// KAEMiddle: # Discussion
	KAEMiddle KAEAND = 'm'<<24 | 'i'<<16 | 'd'<<8 | 'd' // 'midd'
	// KAENOT: Specifies a logical [NOT] operation.
	KAENOT KAEAND = 'N'<<24 | 'O'<<16 | 'T'<<8 | ' ' // 'NOT '
	// KAENext: Specifies the Apple event object after the container.
	KAENext KAEAND = 'n'<<24 | 'e'<<16 | 'x'<<8 | 't' // 'next'
	// KAEOR: Specifies a logical [OR] operation.
	KAEOR KAEAND = 'O'<<24 | 'R'<<16 | ' '<<8 | ' ' // 'OR  '
	// KAEPrevious: Specifies the Apple event object before the container.
	KAEPrevious KAEAND = 'p'<<24 | 'r'<<16 | 'e'<<8 | 'v' // 'prev'
	// KeyAECompOperator: Specifies a descriptor of `typeType`, whose data consists of one of the constant values described in Key Form and Descriptor Type Object Specifier Constants.
	KeyAECompOperator KAEAND = 'r'<<24 | 'e'<<16 | 'l'<<8 | 'o' // 'relo'
	// KeyAEContainer: Specifies the container for the requested object or objects.
	KeyAEContainer KAEAND = 'f'<<24 | 'r'<<16 | 'o'<<8 | 'm' // 'from'
	// KeyAEDesiredClass: # Discussion
	KeyAEDesiredClass KAEAND = 'w'<<24 | 'a'<<16 | 'n'<<8 | 't' // 'want'
	// KeyAEKeyData: # Discussion
	KeyAEKeyData KAEAND = 's'<<24 | 'e'<<16 | 'l'<<8 | 'd' // 'seld'
	// KeyAEKeyForm: # Discussion
	KeyAEKeyForm KAEAND = 'f'<<24 | 'o'<<16 | 'r'<<8 | 'm' // 'form'
	// KeyAELogicalOperator: Specifies a descriptor of type `typeEnumerated` whose data is one of the logical operators (such as `kAEAND`) defined in Key Form and Descriptor Type Object Specifier Constants.
	KeyAELogicalOperator KAEAND = 'l'<<24 | 'o'<<16 | 'g'<<8 | 'c' // 'logc'
	// KeyAELogicalTerms: Specifies a descriptor of type `typeAEList` containing one or more comparison or logical descriptors.
	KeyAELogicalTerms KAEAND = 't'<<24 | 'e'<<16 | 'r'<<8 | 'm' // 'term'
	// KeyAEObject1: # Discussion
	KeyAEObject1 KAEAND = 'o'<<24 | 'b'<<16 | 'j'<<8 | '1' // 'obj1'
	// KeyAEObject2: # Discussion
	KeyAEObject2 KAEAND = 'o'<<24 | 'b'<<16 | 'j'<<8 | '2' // 'obj2'
)

func (e KAEAND) String() string {
	switch e {
	case KAEANDValue:
		return "KAEANDValue"
	case KAEAll:
		return "KAEAll"
	case KAEAny:
		return "KAEAny"
	case KAEFirst:
		return "KAEFirst"
	case KAELast:
		return "KAELast"
	case KAEMiddle:
		return "KAEMiddle"
	case KAENOT:
		return "KAENOT"
	case KAENext:
		return "KAENext"
	case KAEOR:
		return "KAEOR"
	case KAEPrevious:
		return "KAEPrevious"
	case KeyAECompOperator:
		return "KeyAECompOperator"
	case KeyAEContainer:
		return "KeyAEContainer"
	case KeyAEDesiredClass:
		return "KeyAEDesiredClass"
	case KeyAEKeyData:
		return "KeyAEKeyData"
	case KeyAEKeyForm:
		return "KeyAEKeyForm"
	case KeyAELogicalOperator:
		return "KeyAELogicalOperator"
	case KeyAELogicalTerms:
		return "KeyAELogicalTerms"
	case KeyAEObject1:
		return "KeyAEObject1"
	case KeyAEObject2:
		return "KeyAEObject2"
	default:
		return fmt.Sprintf("KAEAND(%d)", e)
	}
}

type KAEDebug uint

const (
	KAEDebugPOSTHeader KAEDebug = 1
	KAEDebugReplyHeader KAEDebug = 2
	KAEDebugXMLDebugAll KAEDebug = 0
	KAEDebugXMLRequest KAEDebug = 4
	KAEDebugXMLResponse KAEDebug = 8
)

func (e KAEDebug) String() string {
	switch e {
	case KAEDebugPOSTHeader:
		return "KAEDebugPOSTHeader"
	case KAEDebugReplyHeader:
		return "KAEDebugReplyHeader"
	case KAEDebugXMLDebugAll:
		return "KAEDebugXMLDebugAll"
	case KAEDebugXMLRequest:
		return "KAEDebugXMLRequest"
	case KAEDebugXMLResponse:
		return "KAEDebugXMLResponse"
	default:
		return fmt.Sprintf("KAEDebug(%d)", e)
	}
}

type KAEDefaultTimeout int

const (
	// KAEDefaultTimeoutValue: The timeout value is determined by the Apple Event Manager.
	KAEDefaultTimeoutValue KAEDefaultTimeout = -1
	// KNoTimeOut: Your application is willing to wait indefinitely.
	KNoTimeOut KAEDefaultTimeout = -2
)

func (e KAEDefaultTimeout) String() string {
	switch e {
	case KAEDefaultTimeoutValue:
		return "KAEDefaultTimeoutValue"
	case KNoTimeOut:
		return "KNoTimeOut"
	default:
		return fmt.Sprintf("KAEDefaultTimeout(%d)", e)
	}
}

type KAEDescListFactor uint

const (
	KAEDescListFactorNone KAEDescListFactor = 0
	KAEDescListFactorType KAEDescListFactor = 4
	KAEDescListFactorTypeAndSize KAEDescListFactor = 8
)

func (e KAEDescListFactor) String() string {
	switch e {
	case KAEDescListFactorNone:
		return "KAEDescListFactorNone"
	case KAEDescListFactorType:
		return "KAEDescListFactorType"
	case KAEDescListFactorTypeAndSize:
		return "KAEDescListFactorTypeAndSize"
	default:
		return fmt.Sprintf("KAEDescListFactor(%d)", e)
	}
}

const KAEHandleArray uint = 2

type KAEIS uint

const (
	KAEISAction KAEIS = 'K'<<24 | 'a'<<16 | 'c'<<8 | 't' // 'Kact'
	KAEISActionPath KAEIS = 'K'<<24 | 'a'<<16 | 'p'<<8 | 't' // 'Kapt'
	KAEISClientAddress KAEIS = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'addr'
	KAEISClientIP KAEIS = 'K'<<24 | 'c'<<16 | 'i'<<8 | 'p' // 'Kcip'
	KAEISContentType KAEIS = 'c'<<24 | 't'<<16 | 'y'<<8 | 'p' // 'ctyp'
	KAEISFromUser KAEIS = 'f'<<24 | 'r'<<16 | 'm'<<8 | 'u' // 'frmu'
	KAEISFullRequest KAEIS = 'K'<<24 | 'f'<<16 | 'r'<<8 | 'q' // 'Kfrq'
	KAEISHTTPSearchArgs KAEIS = 'k'<<24 | 'f'<<16 | 'o'<<8 | 'r' // 'kfor'
	KAEISMethod KAEIS = 'm'<<24 | 'e'<<16 | 't'<<8 | 'h' // 'meth'
	KAEISPassword KAEIS = 'p'<<24 | 'a'<<16 | 's'<<8 | 's' // 'pass'
	KAEISPostArgs KAEIS = 'p'<<24 | 'o'<<16 | 's'<<8 | 't' // 'post'
	KAEISReferrer KAEIS = 'r'<<24 | 'e'<<16 | 'f'<<8 | 'r' // 'refr'
	KAEISScriptName KAEIS = 's'<<24 | 'c'<<16 | 'n'<<8 | 'm' // 'scnm'
	KAEISServerName KAEIS = 's'<<24 | 'v'<<16 | 'n'<<8 | 'm' // 'svnm'
	KAEISServerPort KAEIS = 's'<<24 | 'v'<<16 | 'p'<<8 | 't' // 'svpt'
	KAEISUserAgent KAEIS = 'A'<<24 | 'g'<<16 | 'n'<<8 | 't' // 'Agnt'
	KAEISUserName KAEIS = 'u'<<24 | 's'<<16 | 'e'<<8 | 'r' // 'user'
)

func (e KAEIS) String() string {
	switch e {
	case KAEISAction:
		return "KAEISAction"
	case KAEISActionPath:
		return "KAEISActionPath"
	case KAEISClientAddress:
		return "KAEISClientAddress"
	case KAEISClientIP:
		return "KAEISClientIP"
	case KAEISContentType:
		return "KAEISContentType"
	case KAEISFromUser:
		return "KAEISFromUser"
	case KAEISFullRequest:
		return "KAEISFullRequest"
	case KAEISHTTPSearchArgs:
		return "KAEISHTTPSearchArgs"
	case KAEISMethod:
		return "KAEISMethod"
	case KAEISPassword:
		return "KAEISPassword"
	case KAEISPostArgs:
		return "KAEISPostArgs"
	case KAEISReferrer:
		return "KAEISReferrer"
	case KAEISScriptName:
		return "KAEISScriptName"
	case KAEISServerName:
		return "KAEISServerName"
	case KAEISServerPort:
		return "KAEISServerPort"
	case KAEISUserAgent:
		return "KAEISUserAgent"
	case KAEISUserName:
		return "KAEISUserName"
	default:
		return fmt.Sprintf("KAEIS(%d)", e)
	}
}

type KAEISGetURL uint

const (
	KAEISGetURLValue KAEISGetURL = 'g'<<24 | 'u'<<16 | 'r'<<8 | 'l' // 'gurl'
	KAEISHandleCGI KAEISGetURL = 's'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'sdoc'
)

func (e KAEISGetURL) String() string {
	switch e {
	case KAEISGetURLValue:
		return "KAEISGetURLValue"
	case KAEISHandleCGI:
		return "KAEISHandleCGI"
	default:
		return fmt.Sprintf("KAEISGetURL(%d)", e)
	}
}

type KAEQuit uint

const (
	KAEQuitPreserveState KAEQuit = 's'<<24 | 't'<<16 | 'a'<<8 | 't' // 'stat'
	KAEQuitReason KAEQuit = 'w'<<24 | 'h'<<16 | 'y'<<8 | '?' // 'why?'
)

func (e KAEQuit) String() string {
	switch e {
	case KAEQuitPreserveState:
		return "KAEQuitPreserveState"
	case KAEQuitReason:
		return "KAEQuitReason"
	default:
		return fmt.Sprintf("KAEQuit(%d)", e)
	}
}

type KAEU uint

const (
	KAEUTApostrophe KAEU = 3
	KAEUTChangesState KAEU = 12
	KAEUTDirectParamIsReference KAEU = 9
	KAEUTEnumListIsExclusive KAEU = 10
	KAEUTEnumerated KAEU = 13
	KAEUTEnumsAreTypes KAEU = 11
	KAEUTFeminine KAEU = 2
	KAEUTHasReturningParam KAEU = 31
	KAEUTMasculine KAEU = 1
	KAEUTNotDirectParamIsTarget KAEU = 8
	KAEUTOptional KAEU = 15
	KAEUTParamIsReference KAEU = 9
	KAEUTParamIsTarget KAEU = 8
	KAEUTPlural KAEU = 0
	KAEUTPropertyIsReference KAEU = 9
	KAEUTReadWrite KAEU = 12
	KAEUTReplyIsReference KAEU = 9
	KAEUTTightBindingFunction KAEU = 12
	KAEUTlistOfItems KAEU = 14
)

func (e KAEU) String() string {
	switch e {
	case KAEUTApostrophe:
		return "KAEUTApostrophe"
	case KAEUTChangesState:
		return "KAEUTChangesState"
	case KAEUTDirectParamIsReference:
		return "KAEUTDirectParamIsReference"
	case KAEUTEnumListIsExclusive:
		return "KAEUTEnumListIsExclusive"
	case KAEUTEnumerated:
		return "KAEUTEnumerated"
	case KAEUTEnumsAreTypes:
		return "KAEUTEnumsAreTypes"
	case KAEUTFeminine:
		return "KAEUTFeminine"
	case KAEUTHasReturningParam:
		return "KAEUTHasReturningParam"
	case KAEUTMasculine:
		return "KAEUTMasculine"
	case KAEUTNotDirectParamIsTarget:
		return "KAEUTNotDirectParamIsTarget"
	case KAEUTOptional:
		return "KAEUTOptional"
	case KAEUTPlural:
		return "KAEUTPlural"
	case KAEUTlistOfItems:
		return "KAEUTlistOfItems"
	default:
		return fmt.Sprintf("KAEU(%d)", e)
	}
}

const (
	KAEZoomIn uint = 7
	KAEZoomOut uint = 8
)

type KALM int

const (
	KALMDeferSwitchErr KALM = -30043
	KALMDuplicateModuleErr KALM = -30045
	KALMGroupNotFoundErr KALM = -30048
	KALMInstallationErr KALM = -30044
	KALMInternalErr KALM = -30049
	KALMModuleCommunicationErr KALM = -30046
	KALMNoSuchModuleErr KALM = -30047
	KALMRebootFlagsLevelErr KALM = -30042
)

func (e KALM) String() string {
	switch e {
	case KALMDeferSwitchErr:
		return "KALMDeferSwitchErr"
	case KALMDuplicateModuleErr:
		return "KALMDuplicateModuleErr"
	case KALMGroupNotFoundErr:
		return "KALMGroupNotFoundErr"
	case KALMInstallationErr:
		return "KALMInstallationErr"
	case KALMInternalErr:
		return "KALMInternalErr"
	case KALMModuleCommunicationErr:
		return "KALMModuleCommunicationErr"
	case KALMNoSuchModuleErr:
		return "KALMNoSuchModuleErr"
	case KALMRebootFlagsLevelErr:
		return "KALMRebootFlagsLevelErr"
	default:
		return fmt.Sprintf("KALM(%d)", e)
	}
}

const KALMLocationNotFoundErr int = -30048

type KATSU int

const (
	KATSUBadStreamErr KATSU = -8902
	KATSUBusyObjectErr KATSU = -8809
	KATSUCoordinateOverflowErr KATSU = -8807
	KATSUFontsMatched KATSU = -8793
	KATSUFontsNotMatched KATSU = -8794
	KATSUInvalidAttributeSizeErr KATSU = -8798
	KATSUInvalidAttributeTagErr KATSU = -8799
	KATSUInvalidAttributeValueErr KATSU = -8797
	KATSUInvalidCacheErr KATSU = -8800
	KATSUInvalidCallInsideCallbackErr KATSU = -8904
	KATSUInvalidFontErr KATSU = -8796
	KATSUInvalidFontFallbacksErr KATSU = -8900
	KATSUInvalidStyleErr KATSU = -8791
	KATSUInvalidTextLayoutErr KATSU = -8790
	KATSUInvalidTextRangeErr KATSU = -8792
	KATSULastErr KATSU = -8959
	KATSULineBreakInWord KATSU = -8808
	KATSULowLevelErr KATSU = -8804
	KATSUNoCorrespondingFontErr KATSU = -8795
	KATSUNoFontCmapAvailableErr KATSU = -8805
	KATSUNoFontNameErr KATSU = -8905
	KATSUNoFontScalerAvailableErr KATSU = -8806
	KATSUNoStyleRunsAssignedErr KATSU = -8802
	KATSUNotSetErr KATSU = -8801
	KATSUOutputBufferTooSmallErr KATSU = -8903
	KATSUQuickDrawTextErr KATSU = -8803
	KATSUUnsupportedStreamFormatErr KATSU = -8901
)

func (e KATSU) String() string {
	switch e {
	case KATSUBadStreamErr:
		return "KATSUBadStreamErr"
	case KATSUBusyObjectErr:
		return "KATSUBusyObjectErr"
	case KATSUCoordinateOverflowErr:
		return "KATSUCoordinateOverflowErr"
	case KATSUFontsMatched:
		return "KATSUFontsMatched"
	case KATSUFontsNotMatched:
		return "KATSUFontsNotMatched"
	case KATSUInvalidAttributeSizeErr:
		return "KATSUInvalidAttributeSizeErr"
	case KATSUInvalidAttributeTagErr:
		return "KATSUInvalidAttributeTagErr"
	case KATSUInvalidAttributeValueErr:
		return "KATSUInvalidAttributeValueErr"
	case KATSUInvalidCacheErr:
		return "KATSUInvalidCacheErr"
	case KATSUInvalidCallInsideCallbackErr:
		return "KATSUInvalidCallInsideCallbackErr"
	case KATSUInvalidFontErr:
		return "KATSUInvalidFontErr"
	case KATSUInvalidFontFallbacksErr:
		return "KATSUInvalidFontFallbacksErr"
	case KATSUInvalidStyleErr:
		return "KATSUInvalidStyleErr"
	case KATSUInvalidTextLayoutErr:
		return "KATSUInvalidTextLayoutErr"
	case KATSUInvalidTextRangeErr:
		return "KATSUInvalidTextRangeErr"
	case KATSULastErr:
		return "KATSULastErr"
	case KATSULineBreakInWord:
		return "KATSULineBreakInWord"
	case KATSULowLevelErr:
		return "KATSULowLevelErr"
	case KATSUNoCorrespondingFontErr:
		return "KATSUNoCorrespondingFontErr"
	case KATSUNoFontCmapAvailableErr:
		return "KATSUNoFontCmapAvailableErr"
	case KATSUNoFontNameErr:
		return "KATSUNoFontNameErr"
	case KATSUNoFontScalerAvailableErr:
		return "KATSUNoFontScalerAvailableErr"
	case KATSUNoStyleRunsAssignedErr:
		return "KATSUNoStyleRunsAssignedErr"
	case KATSUNotSetErr:
		return "KATSUNotSetErr"
	case KATSUOutputBufferTooSmallErr:
		return "KATSUOutputBufferTooSmallErr"
	case KATSUQuickDrawTextErr:
		return "KATSUQuickDrawTextErr"
	case KATSUUnsupportedStreamFormatErr:
		return "KATSUUnsupportedStreamFormatErr"
	default:
		return fmt.Sprintf("KATSU(%d)", e)
	}
}

type KAlert uint

const (
	KAlertCautionIcon KAlert = 'c'<<24 | 'a'<<16 | 'u'<<8 | 't' // 'caut'
	KAlertNoteIcon KAlert = 'n'<<24 | 'o'<<16 | 't'<<8 | 'e' // 'note'
	KAlertStopIcon KAlert = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
)

func (e KAlert) String() string {
	switch e {
	case KAlertCautionIcon:
		return "KAlertCautionIcon"
	case KAlertNoteIcon:
		return "KAlertNoteIcon"
	case KAlertStopIcon:
		return "KAlertStopIcon"
	default:
		return fmt.Sprintf("KAlert(%d)", e)
	}
}

type KAppearanceFolderIcon uint

const (
	KAppearanceFolderIconValue KAppearanceFolderIcon = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'r' // 'appr'
	KAppleExtrasFolderIcon KAppearanceFolderIcon = 1634040004
	KAppleMenuFolderIcon KAppearanceFolderIcon = 'a'<<24 | 'm'<<16 | 'n'<<8 | 'u' // 'amnu'
	KApplicationSupportFolderIcon KAppearanceFolderIcon = 'a'<<24 | 's'<<16 | 'u'<<8 | 'p' // 'asup'
	KApplicationsFolderIcon KAppearanceFolderIcon = 'a'<<24 | 'p'<<16 | 'p'<<8 | 's' // 'apps'
	KAssistantsFolderIcon KAppearanceFolderIcon = 1634956484
	KColorSyncFolderIcon KAppearanceFolderIcon = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'f' // 'prof'
	KContextualMenuItemsFolderIcon KAppearanceFolderIcon = 'c'<<24 | 'm'<<16 | 'n'<<8 | 'u' // 'cmnu'
	KControlPanelDisabledFolderIcon KAppearanceFolderIcon = 'c'<<24 | 't'<<16 | 'r'<<8 | 'D' // 'ctrD'
	KControlPanelFolderIcon KAppearanceFolderIcon = 'c'<<24 | 't'<<16 | 'r'<<8 | 'l' // 'ctrl'
	KControlStripModulesFolderIcon KAppearanceFolderIcon = 1935963844
	KDocumentsFolderIcon KAppearanceFolderIcon = 'd'<<24 | 'o'<<16 | 'c'<<8 | 's' // 'docs'
	KExtensionsDisabledFolderIcon KAppearanceFolderIcon = 'e'<<24 | 'x'<<16 | 't'<<8 | 'D' // 'extD'
	KExtensionsFolderIcon KAppearanceFolderIcon = 'e'<<24 | 'x'<<16 | 't'<<8 | 'n' // 'extn'
	KFavoritesFolderIcon KAppearanceFolderIcon = 'f'<<24 | 'a'<<16 | 'v'<<8 | 's' // 'favs'
	KFontsFolderIcon KAppearanceFolderIcon = 'f'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'font'
	KHelpFolderIcon KAppearanceFolderIcon = 0
	KInternetFolderIcon KAppearanceFolderIcon = 1768846532
	KInternetPlugInFolderIcon KAppearanceFolderIcon = 0
	KInternetSearchSitesFolderIcon KAppearanceFolderIcon = 'i'<<24 | 's'<<16 | 's'<<8 | 'f' // 'issf'
	KLocalesFolderIcon KAppearanceFolderIcon = 0
	KMacOSReadMeFolderIcon KAppearanceFolderIcon = 1836020420
	KPreferencesFolderIcon KAppearanceFolderIcon = 1886545604
	KPrintMonitorFolderIcon KAppearanceFolderIcon = 'p'<<24 | 'r'<<16 | 'n'<<8 | 't' // 'prnt'
	KPrinterDescriptionFolderIcon KAppearanceFolderIcon = 'p'<<24 | 'p'<<16 | 'd'<<8 | 'f' // 'ppdf'
	KPrinterDriverFolderIcon KAppearanceFolderIcon = 0
	KPublicFolderIcon KAppearanceFolderIcon = 'p'<<24 | 'u'<<16 | 'b'<<8 | 'f' // 'pubf'
	KRecentApplicationsFolderIcon KAppearanceFolderIcon = 'r'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'rapp'
	KRecentDocumentsFolderIcon KAppearanceFolderIcon = 'r'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'rdoc'
	KRecentServersFolderIcon KAppearanceFolderIcon = 'r'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'rsrv'
	KScriptingAdditionsFolderIcon KAppearanceFolderIcon = 0
	KScriptsFolderIcon KAppearanceFolderIcon = 1935897284
	KSharedLibrariesFolderIcon KAppearanceFolderIcon = 0
	KShutdownItemsDisabledFolderIcon KAppearanceFolderIcon = 's'<<24 | 'h'<<16 | 'd'<<8 | 'D' // 'shdD'
	KShutdownItemsFolderIcon KAppearanceFolderIcon = 's'<<24 | 'h'<<16 | 'd'<<8 | 'f' // 'shdf'
	KSpeakableItemsFolder KAppearanceFolderIcon = 's'<<24 | 'p'<<16 | 'k'<<8 | 'i' // 'spki'
	KStartupItemsDisabledFolderIcon KAppearanceFolderIcon = 's'<<24 | 't'<<16 | 'r'<<8 | 'D' // 'strD'
	KStartupItemsFolderIcon KAppearanceFolderIcon = 's'<<24 | 't'<<16 | 'r'<<8 | 't' // 'strt'
	KSystemExtensionDisabledFolderIcon KAppearanceFolderIcon = 'm'<<24 | 'a'<<16 | 'c'<<8 | 'D' // 'macD'
	KSystemFolderIcon KAppearanceFolderIcon = 'm'<<24 | 'a'<<16 | 'c'<<8 | 's' // 'macs'
	KTextEncodingsFolderIcon KAppearanceFolderIcon = 0
	KUsersFolderIcon KAppearanceFolderIcon = 1970500292
	KUtilitiesFolderIcon KAppearanceFolderIcon = 1970563524
	KVoicesFolderIcon KAppearanceFolderIcon = 'f'<<24 | 'v'<<16 | 'o'<<8 | 'c' // 'fvoc'
)

func (e KAppearanceFolderIcon) String() string {
	switch e {
	case KAppearanceFolderIconValue:
		return "KAppearanceFolderIconValue"
	case KAppleExtrasFolderIcon:
		return "KAppleExtrasFolderIcon"
	case KAppleMenuFolderIcon:
		return "KAppleMenuFolderIcon"
	case KApplicationSupportFolderIcon:
		return "KApplicationSupportFolderIcon"
	case KApplicationsFolderIcon:
		return "KApplicationsFolderIcon"
	case KAssistantsFolderIcon:
		return "KAssistantsFolderIcon"
	case KColorSyncFolderIcon:
		return "KColorSyncFolderIcon"
	case KContextualMenuItemsFolderIcon:
		return "KContextualMenuItemsFolderIcon"
	case KControlPanelDisabledFolderIcon:
		return "KControlPanelDisabledFolderIcon"
	case KControlPanelFolderIcon:
		return "KControlPanelFolderIcon"
	case KControlStripModulesFolderIcon:
		return "KControlStripModulesFolderIcon"
	case KDocumentsFolderIcon:
		return "KDocumentsFolderIcon"
	case KExtensionsDisabledFolderIcon:
		return "KExtensionsDisabledFolderIcon"
	case KExtensionsFolderIcon:
		return "KExtensionsFolderIcon"
	case KFavoritesFolderIcon:
		return "KFavoritesFolderIcon"
	case KFontsFolderIcon:
		return "KFontsFolderIcon"
	case KHelpFolderIcon:
		return "KHelpFolderIcon"
	case KInternetFolderIcon:
		return "KInternetFolderIcon"
	case KInternetSearchSitesFolderIcon:
		return "KInternetSearchSitesFolderIcon"
	case KMacOSReadMeFolderIcon:
		return "KMacOSReadMeFolderIcon"
	case KPreferencesFolderIcon:
		return "KPreferencesFolderIcon"
	case KPrintMonitorFolderIcon:
		return "KPrintMonitorFolderIcon"
	case KPrinterDescriptionFolderIcon:
		return "KPrinterDescriptionFolderIcon"
	case KPublicFolderIcon:
		return "KPublicFolderIcon"
	case KRecentApplicationsFolderIcon:
		return "KRecentApplicationsFolderIcon"
	case KRecentDocumentsFolderIcon:
		return "KRecentDocumentsFolderIcon"
	case KRecentServersFolderIcon:
		return "KRecentServersFolderIcon"
	case KScriptsFolderIcon:
		return "KScriptsFolderIcon"
	case KShutdownItemsDisabledFolderIcon:
		return "KShutdownItemsDisabledFolderIcon"
	case KShutdownItemsFolderIcon:
		return "KShutdownItemsFolderIcon"
	case KSpeakableItemsFolder:
		return "KSpeakableItemsFolder"
	case KStartupItemsDisabledFolderIcon:
		return "KStartupItemsDisabledFolderIcon"
	case KStartupItemsFolderIcon:
		return "KStartupItemsFolderIcon"
	case KSystemExtensionDisabledFolderIcon:
		return "KSystemExtensionDisabledFolderIcon"
	case KSystemFolderIcon:
		return "KSystemFolderIcon"
	case KUsersFolderIcon:
		return "KUsersFolderIcon"
	case KUtilitiesFolderIcon:
		return "KUtilitiesFolderIcon"
	case KVoicesFolderIcon:
		return "KVoicesFolderIcon"
	default:
		return fmt.Sprintf("KAppearanceFolderIcon(%d)", e)
	}
}

type KAppleLogoIcon uint

const (
	KAppleLogoIconValue KAppleLogoIcon = 'c'<<24 | 'a'<<16 | 'p'<<8 | 'l' // 'capl'
	KAppleMenuIcon KAppleLogoIcon = 's'<<24 | 'a'<<16 | 'p'<<8 | 'l' // 'sapl'
	KBackwardArrowIcon KAppleLogoIcon = 'b'<<24 | 'a'<<16 | 'r'<<8 | 'o' // 'baro'
	KBurningIcon KAppleLogoIcon = 'b'<<24 | 'u'<<16 | 'r'<<8 | 'n' // 'burn'
	KConnectToIcon KAppleLogoIcon = 'c'<<24 | 'n'<<16 | 'c'<<8 | 't' // 'cnct'
	KDeleteAliasIcon KAppleLogoIcon = 'd'<<24 | 'a'<<16 | 'l'<<8 | 'i' // 'dali'
	KEjectMediaIcon KAppleLogoIcon = 'e'<<24 | 'j'<<16 | 'e'<<8 | 'c' // 'ejec'
	KFavoriteItemsIcon KAppleLogoIcon = 'f'<<24 | 'a'<<16 | 'v'<<8 | 'r' // 'favr'
	KForwardArrowIcon KAppleLogoIcon = 'f'<<24 | 'a'<<16 | 'r'<<8 | 'o' // 'faro'
	KGenericWindowIcon KAppleLogoIcon = 'g'<<24 | 'w'<<16 | 'i'<<8 | 'n' // 'gwin'
	KGridIcon KAppleLogoIcon = 'g'<<24 | 'r'<<16 | 'i'<<8 | 'd' // 'grid'
	KHelpIcon KAppleLogoIcon = 'h'<<24 | 'e'<<16 | 'l'<<8 | 'p' // 'help'
	KKeepArrangedIcon KAppleLogoIcon = 'a'<<24 | 'r'<<16 | 'n'<<8 | 'g' // 'arng'
	KLockedIcon KAppleLogoIcon = 'l'<<24 | 'o'<<16 | 'c'<<8 | 'k' // 'lock'
	KNoFilesIcon KAppleLogoIcon = 'n'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'nfil'
	KNoFolderIcon KAppleLogoIcon = 'n'<<24 | 'f'<<16 | 'l'<<8 | 'd' // 'nfld'
	KNoWriteIcon KAppleLogoIcon = 'n'<<24 | 'w'<<16 | 'r'<<8 | 't' // 'nwrt'
	KProtectedApplicationFolderIcon KAppleLogoIcon = 'p'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'papp'
	KProtectedSystemFolderIcon KAppleLogoIcon = 'p'<<24 | 's'<<16 | 'y'<<8 | 's' // 'psys'
	KQuestionMarkIcon KAppleLogoIcon = 'q'<<24 | 'u'<<16 | 'e'<<8 | 's' // 'ques'
	KRecentItemsIcon KAppleLogoIcon = 'r'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'rcnt'
	KRightContainerArrowIcon KAppleLogoIcon = 'r'<<24 | 'c'<<16 | 'a'<<8 | 'r' // 'rcar'
	KShortcutIcon KAppleLogoIcon = 's'<<24 | 'h'<<16 | 'r'<<8 | 't' // 'shrt'
	KSortAscendingIcon KAppleLogoIcon = 'a'<<24 | 's'<<16 | 'n'<<8 | 'd' // 'asnd'
	KSortDescendingIcon KAppleLogoIcon = 'd'<<24 | 's'<<16 | 'n'<<8 | 'd' // 'dsnd'
	KUnlockedIcon KAppleLogoIcon = 'u'<<24 | 'l'<<16 | 'c'<<8 | 'k' // 'ulck'
)

func (e KAppleLogoIcon) String() string {
	switch e {
	case KAppleLogoIconValue:
		return "KAppleLogoIconValue"
	case KAppleMenuIcon:
		return "KAppleMenuIcon"
	case KBackwardArrowIcon:
		return "KBackwardArrowIcon"
	case KBurningIcon:
		return "KBurningIcon"
	case KConnectToIcon:
		return "KConnectToIcon"
	case KDeleteAliasIcon:
		return "KDeleteAliasIcon"
	case KEjectMediaIcon:
		return "KEjectMediaIcon"
	case KFavoriteItemsIcon:
		return "KFavoriteItemsIcon"
	case KForwardArrowIcon:
		return "KForwardArrowIcon"
	case KGenericWindowIcon:
		return "KGenericWindowIcon"
	case KGridIcon:
		return "KGridIcon"
	case KHelpIcon:
		return "KHelpIcon"
	case KKeepArrangedIcon:
		return "KKeepArrangedIcon"
	case KLockedIcon:
		return "KLockedIcon"
	case KNoFilesIcon:
		return "KNoFilesIcon"
	case KNoFolderIcon:
		return "KNoFolderIcon"
	case KNoWriteIcon:
		return "KNoWriteIcon"
	case KProtectedApplicationFolderIcon:
		return "KProtectedApplicationFolderIcon"
	case KProtectedSystemFolderIcon:
		return "KProtectedSystemFolderIcon"
	case KQuestionMarkIcon:
		return "KQuestionMarkIcon"
	case KRecentItemsIcon:
		return "KRecentItemsIcon"
	case KRightContainerArrowIcon:
		return "KRightContainerArrowIcon"
	case KShortcutIcon:
		return "KShortcutIcon"
	case KSortAscendingIcon:
		return "KSortAscendingIcon"
	case KSortDescendingIcon:
		return "KSortDescendingIcon"
	case KUnlockedIcon:
		return "KUnlockedIcon"
	default:
		return fmt.Sprintf("KAppleLogoIcon(%d)", e)
	}
}

type KAppleScriptBadgeIcon uint

const (
	KAlertCautionBadgeIcon KAppleScriptBadgeIcon = 'c'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'cbdg'
	KAliasBadgeIcon KAppleScriptBadgeIcon = 'a'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'abdg'
	KAppleScriptBadgeIconValue KAppleScriptBadgeIcon = 's'<<24 | 'c'<<16 | 'r'<<8 | 'p' // 'scrp'
	KLockedBadgeIcon KAppleScriptBadgeIcon = 'l'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'lbdg'
	KMountedBadgeIcon KAppleScriptBadgeIcon = 'm'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'mbdg'
	KSharedBadgeIcon KAppleScriptBadgeIcon = 's'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'sbdg'
)

func (e KAppleScriptBadgeIcon) String() string {
	switch e {
	case KAlertCautionBadgeIcon:
		return "KAlertCautionBadgeIcon"
	case KAliasBadgeIcon:
		return "KAliasBadgeIcon"
	case KAppleScriptBadgeIconValue:
		return "KAppleScriptBadgeIconValue"
	case KLockedBadgeIcon:
		return "KLockedBadgeIcon"
	case KMountedBadgeIcon:
		return "KMountedBadgeIcon"
	case KSharedBadgeIcon:
		return "KSharedBadgeIcon"
	default:
		return fmt.Sprintf("KAppleScriptBadgeIcon(%d)", e)
	}
}

type KAppleTalkIcon uint

const (
	KAFPServerIcon KAppleTalkIcon = 'a'<<24 | 'f'<<16 | 'p'<<8 | 's' // 'afps'
	KAppleTalkIconValue KAppleTalkIcon = 'a'<<24 | 't'<<16 | 'l'<<8 | 'k' // 'atlk'
	KAppleTalkZoneIcon KAppleTalkIcon = 'a'<<24 | 't'<<16 | 'z'<<8 | 'n' // 'atzn'
	KFTPServerIcon KAppleTalkIcon = 'f'<<24 | 't'<<16 | 'p'<<8 | 's' // 'ftps'
	KGenericNetworkIcon KAppleTalkIcon = 'g'<<24 | 'n'<<16 | 'e'<<8 | 't' // 'gnet'
	KHTTPServerIcon KAppleTalkIcon = 'h'<<24 | 't'<<16 | 'p'<<8 | 's' // 'htps'
	KIPFileServerIcon KAppleTalkIcon = 'i'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'isrv'
)

func (e KAppleTalkIcon) String() string {
	switch e {
	case KAFPServerIcon:
		return "KAFPServerIcon"
	case KAppleTalkIconValue:
		return "KAppleTalkIconValue"
	case KAppleTalkZoneIcon:
		return "KAppleTalkZoneIcon"
	case KFTPServerIcon:
		return "KFTPServerIcon"
	case KGenericNetworkIcon:
		return "KGenericNetworkIcon"
	case KHTTPServerIcon:
		return "KHTTPServerIcon"
	case KIPFileServerIcon:
		return "KIPFileServerIcon"
	default:
		return fmt.Sprintf("KAppleTalkIcon(%d)", e)
	}
}

type KAutoGenerateReturnID int

const (
	// KAnyTransactionID: # Discussion
	KAnyTransactionID KAutoGenerateReturnID = 0
	// KAutoGenerateReturnIDValue: If you pass this value for the `returnID` parameter of the AECreateAppleEvent(_:_:_:_:_:_:) function, the Apple Event Manager assigns to the created Apple event a return ID that is unique to the current session.
	KAutoGenerateReturnIDValue KAutoGenerateReturnID = -1
)

func (e KAutoGenerateReturnID) String() string {
	switch e {
	case KAnyTransactionID:
		return "KAnyTransactionID"
	case KAutoGenerateReturnIDValue:
		return "KAutoGenerateReturnIDValue"
	default:
		return fmt.Sprintf("KAutoGenerateReturnID(%d)", e)
	}
}

type KBadAdapterErr int

const (
	K16BitCardErr KBadAdapterErr = -9084
	KAlreadySavedStateErr KBadAdapterErr = -9105
	KAttemptDupCardEntryErr KBadAdapterErr = -9106
	KBadAdapterErrValue KBadAdapterErr = -9050
	KBadArgLengthErr KBadAdapterErr = -9063
	KBadArgsErr KBadAdapterErr = -9064
	KBadAttributeErr KBadAdapterErr = -9051
	KBadBaseErr KBadAdapterErr = -9052
	KBadCISErr KBadAdapterErr = -9066
	KBadCustomIFIDErr KBadAdapterErr = -9093
	KBadDeviceErr KBadAdapterErr = -9083
	KBadEDCErr KBadAdapterErr = -9053
	KBadHandleErr KBadAdapterErr = -9065
	KBadIRQErr KBadAdapterErr = -9054
	KBadLinkErr KBadAdapterErr = -9082
	KBadOffsetErr KBadAdapterErr = -9055
	KBadPageErr KBadAdapterErr = -9056
	KBadSizeErr KBadAdapterErr = -9057
	KBadSocketErr KBadAdapterErr = -9058
	KBadSpeedErr KBadAdapterErr = -9067
	KBadTupleDataErr KBadAdapterErr = -9092
	KBadTypeErr KBadAdapterErr = -9059
	KBadVccErr KBadAdapterErr = -9060
	KBadVppErr KBadAdapterErr = -9061
	KBadWindowErr KBadAdapterErr = -9062
	KBusyErr KBadAdapterErr = -9074
	KCantConfigureCardErr KBadAdapterErr = -9087
	KCardBusCardErr KBadAdapterErr = -9085
	KCardPowerOffErr KBadAdapterErr = -9107
	KClientRequestDenied KBadAdapterErr = -9102
	KConfigurationLockedErr KBadAdapterErr = -9076
	KGeneralFailureErr KBadAdapterErr = -9070
	KInUseErr KBadAdapterErr = -9077
	KInvalidCSClientErr KBadAdapterErr = -9091
	KInvalidDeviceNumber KBadAdapterErr = -9089
	KInvalidRegEntryErr KBadAdapterErr = -9081
	KNoCardBusCISErr KBadAdapterErr = -9109
	KNoCardEnablersFoundErr KBadAdapterErr = -9099
	KNoCardErr KBadAdapterErr = -9071
	KNoCardSevicesSocketsErr KBadAdapterErr = -9080
	KNoClientTableErr KBadAdapterErr = -9097
	KNoCompatibleNameErr KBadAdapterErr = -9101
	KNoEnablerForCardErr KBadAdapterErr = -9100
	KNoIOWindowRequestedErr KBadAdapterErr = -9094
	KNoMoreInterruptSlotsErr KBadAdapterErr = -9096
	KNoMoreItemsErr KBadAdapterErr = -9078
	KNoMoreTimerClientsErr KBadAdapterErr = -9095
	KNotReadyErr KBadAdapterErr = -9103
	KNotZVCapableErr KBadAdapterErr = -9108
	KOutOfResourceErr KBadAdapterErr = -9079
	KPassCallToChainErr KBadAdapterErr = -9086
	KPostCardEventErr KBadAdapterErr = -9088
	KReadFailureErr KBadAdapterErr = -9068
	KTooManyIOWindowsErr KBadAdapterErr = -9104
	KUnsupportedCardErr KBadAdapterErr = -9098
	KUnsupportedFunctionErr KBadAdapterErr = -9072
	KUnsupportedModeErr KBadAdapterErr = -9073
	KUnsupportedVsErr KBadAdapterErr = -9090
	KWriteFailureErr KBadAdapterErr = -9069
	KWriteProtectedErr KBadAdapterErr = -9075
)

func (e KBadAdapterErr) String() string {
	switch e {
	case K16BitCardErr:
		return "K16BitCardErr"
	case KAlreadySavedStateErr:
		return "KAlreadySavedStateErr"
	case KAttemptDupCardEntryErr:
		return "KAttemptDupCardEntryErr"
	case KBadAdapterErrValue:
		return "KBadAdapterErrValue"
	case KBadArgLengthErr:
		return "KBadArgLengthErr"
	case KBadArgsErr:
		return "KBadArgsErr"
	case KBadAttributeErr:
		return "KBadAttributeErr"
	case KBadBaseErr:
		return "KBadBaseErr"
	case KBadCISErr:
		return "KBadCISErr"
	case KBadCustomIFIDErr:
		return "KBadCustomIFIDErr"
	case KBadDeviceErr:
		return "KBadDeviceErr"
	case KBadEDCErr:
		return "KBadEDCErr"
	case KBadHandleErr:
		return "KBadHandleErr"
	case KBadIRQErr:
		return "KBadIRQErr"
	case KBadLinkErr:
		return "KBadLinkErr"
	case KBadOffsetErr:
		return "KBadOffsetErr"
	case KBadPageErr:
		return "KBadPageErr"
	case KBadSizeErr:
		return "KBadSizeErr"
	case KBadSocketErr:
		return "KBadSocketErr"
	case KBadSpeedErr:
		return "KBadSpeedErr"
	case KBadTupleDataErr:
		return "KBadTupleDataErr"
	case KBadTypeErr:
		return "KBadTypeErr"
	case KBadVccErr:
		return "KBadVccErr"
	case KBadVppErr:
		return "KBadVppErr"
	case KBadWindowErr:
		return "KBadWindowErr"
	case KBusyErr:
		return "KBusyErr"
	case KCantConfigureCardErr:
		return "KCantConfigureCardErr"
	case KCardBusCardErr:
		return "KCardBusCardErr"
	case KCardPowerOffErr:
		return "KCardPowerOffErr"
	case KClientRequestDenied:
		return "KClientRequestDenied"
	case KConfigurationLockedErr:
		return "KConfigurationLockedErr"
	case KGeneralFailureErr:
		return "KGeneralFailureErr"
	case KInUseErr:
		return "KInUseErr"
	case KInvalidCSClientErr:
		return "KInvalidCSClientErr"
	case KInvalidDeviceNumber:
		return "KInvalidDeviceNumber"
	case KInvalidRegEntryErr:
		return "KInvalidRegEntryErr"
	case KNoCardBusCISErr:
		return "KNoCardBusCISErr"
	case KNoCardEnablersFoundErr:
		return "KNoCardEnablersFoundErr"
	case KNoCardErr:
		return "KNoCardErr"
	case KNoCardSevicesSocketsErr:
		return "KNoCardSevicesSocketsErr"
	case KNoClientTableErr:
		return "KNoClientTableErr"
	case KNoCompatibleNameErr:
		return "KNoCompatibleNameErr"
	case KNoEnablerForCardErr:
		return "KNoEnablerForCardErr"
	case KNoIOWindowRequestedErr:
		return "KNoIOWindowRequestedErr"
	case KNoMoreInterruptSlotsErr:
		return "KNoMoreInterruptSlotsErr"
	case KNoMoreItemsErr:
		return "KNoMoreItemsErr"
	case KNoMoreTimerClientsErr:
		return "KNoMoreTimerClientsErr"
	case KNotReadyErr:
		return "KNotReadyErr"
	case KNotZVCapableErr:
		return "KNotZVCapableErr"
	case KOutOfResourceErr:
		return "KOutOfResourceErr"
	case KPassCallToChainErr:
		return "KPassCallToChainErr"
	case KPostCardEventErr:
		return "KPostCardEventErr"
	case KReadFailureErr:
		return "KReadFailureErr"
	case KTooManyIOWindowsErr:
		return "KTooManyIOWindowsErr"
	case KUnsupportedCardErr:
		return "KUnsupportedCardErr"
	case KUnsupportedFunctionErr:
		return "KUnsupportedFunctionErr"
	case KUnsupportedModeErr:
		return "KUnsupportedModeErr"
	case KUnsupportedVsErr:
		return "KUnsupportedVsErr"
	case KWriteFailureErr:
		return "KWriteFailureErr"
	case KWriteProtectedErr:
		return "KWriteProtectedErr"
	default:
		return fmt.Sprintf("KBadAdapterErr(%d)", e)
	}
}

const KCSDiskSpaceRecoveryOptionNoUI uint = 1

type KCSIdentity int

const (
	KCSIdentityAuthorityNotAccessibleErr KCSIdentity = -2
	KCSIdentityDeletedErr KCSIdentity = -4
	KCSIdentityDuplicateFullNameErr KCSIdentity = -6
	KCSIdentityDuplicatePosixNameErr KCSIdentity = -8
	KCSIdentityInvalidFullNameErr KCSIdentity = -5
	KCSIdentityInvalidPosixNameErr KCSIdentity = -7
	KCSIdentityPermissionErr KCSIdentity = -3
	KCSIdentityUnknownAuthorityErr KCSIdentity = -1
)

func (e KCSIdentity) String() string {
	switch e {
	case KCSIdentityAuthorityNotAccessibleErr:
		return "KCSIdentityAuthorityNotAccessibleErr"
	case KCSIdentityDeletedErr:
		return "KCSIdentityDeletedErr"
	case KCSIdentityDuplicateFullNameErr:
		return "KCSIdentityDuplicateFullNameErr"
	case KCSIdentityDuplicatePosixNameErr:
		return "KCSIdentityDuplicatePosixNameErr"
	case KCSIdentityInvalidFullNameErr:
		return "KCSIdentityInvalidFullNameErr"
	case KCSIdentityInvalidPosixNameErr:
		return "KCSIdentityInvalidPosixNameErr"
	case KCSIdentityPermissionErr:
		return "KCSIdentityPermissionErr"
	case KCSIdentityUnknownAuthorityErr:
		return "KCSIdentityUnknownAuthorityErr"
	default:
		return fmt.Sprintf("KCSIdentity(%d)", e)
	}
}

type KCSIdentityClass uint

const (
	KCSIdentityClassGroup KCSIdentityClass = 2
	KCSIdentityClassUser KCSIdentityClass = 1
)

func (e KCSIdentityClass) String() string {
	switch e {
	case KCSIdentityClassGroup:
		return "KCSIdentityClassGroup"
	case KCSIdentityClassUser:
		return "KCSIdentityClassUser"
	default:
		return fmt.Sprintf("KCSIdentityClass(%d)", e)
	}
}

const KCSIdentityCommitCompleted uint = 1

type KCSIdentityFlag uint

const (
	KCSIdentityFlagHidden KCSIdentityFlag = 1
	KCSIdentityFlagNone KCSIdentityFlag = 0
)

func (e KCSIdentityFlag) String() string {
	switch e {
	case KCSIdentityFlagHidden:
		return "KCSIdentityFlagHidden"
	case KCSIdentityFlagNone:
		return "KCSIdentityFlagNone"
	default:
		return fmt.Sprintf("KCSIdentityFlag(%d)", e)
	}
}

type KCSIdentityQuery uint

const (
	KCSIdentityQueryGenerateUpdateEvents KCSIdentityQuery = 1
	KCSIdentityQueryIncludeHiddenIdentities KCSIdentityQuery = 2
)

func (e KCSIdentityQuery) String() string {
	switch e {
	case KCSIdentityQueryGenerateUpdateEvents:
		return "KCSIdentityQueryGenerateUpdateEvents"
	case KCSIdentityQueryIncludeHiddenIdentities:
		return "KCSIdentityQueryIncludeHiddenIdentities"
	default:
		return fmt.Sprintf("KCSIdentityQuery(%d)", e)
	}
}

type KCSIdentityQueryEvent uint

const (
	KCSIdentityQueryEventErrorOccurred KCSIdentityQueryEvent = 5
	KCSIdentityQueryEventResultsAdded KCSIdentityQueryEvent = 2
	KCSIdentityQueryEventResultsChanged KCSIdentityQueryEvent = 3
	KCSIdentityQueryEventResultsRemoved KCSIdentityQueryEvent = 4
	KCSIdentityQueryEventSearchPhaseFinished KCSIdentityQueryEvent = 1
)

func (e KCSIdentityQueryEvent) String() string {
	switch e {
	case KCSIdentityQueryEventErrorOccurred:
		return "KCSIdentityQueryEventErrorOccurred"
	case KCSIdentityQueryEventResultsAdded:
		return "KCSIdentityQueryEventResultsAdded"
	case KCSIdentityQueryEventResultsChanged:
		return "KCSIdentityQueryEventResultsChanged"
	case KCSIdentityQueryEventResultsRemoved:
		return "KCSIdentityQueryEventResultsRemoved"
	case KCSIdentityQueryEventSearchPhaseFinished:
		return "KCSIdentityQueryEventSearchPhaseFinished"
	default:
		return fmt.Sprintf("KCSIdentityQueryEvent(%d)", e)
	}
}

type KCSIdentityQueryString uint

const (
	KCSIdentityQueryStringBeginsWith KCSIdentityQueryString = 2
	KCSIdentityQueryStringEquals KCSIdentityQueryString = 1
)

func (e KCSIdentityQueryString) String() string {
	switch e {
	case KCSIdentityQueryStringBeginsWith:
		return "KCSIdentityQueryStringBeginsWith"
	case KCSIdentityQueryStringEquals:
		return "KCSIdentityQueryStringEquals"
	default:
		return fmt.Sprintf("KCSIdentityQueryString(%d)", e)
	}
}

type KClipboardIcon uint

const (
	KClipboardIconValue KClipboardIcon = 'C'<<24 | 'L'<<16 | 'I'<<8 | 'P' // 'CLIP'
	KClippingPictureTypeIcon KClipboardIcon = 'c'<<24 | 'l'<<16 | 'p'<<8 | 'p' // 'clpp'
	KClippingSoundTypeIcon KClipboardIcon = 'c'<<24 | 'l'<<16 | 'p'<<8 | 's' // 'clps'
	KClippingTextTypeIcon KClipboardIcon = 'c'<<24 | 'l'<<16 | 'p'<<8 | 't' // 'clpt'
	KClippingUnknownTypeIcon KClipboardIcon = 'c'<<24 | 'l'<<16 | 'p'<<8 | 'u' // 'clpu'
	KComputerIcon KClipboardIcon = 'r'<<24 | 'o'<<16 | 'o'<<8 | 't' // 'root'
	KDesktopIcon KClipboardIcon = 'd'<<24 | 'e'<<16 | 's'<<8 | 'k' // 'desk'
	KFinderIcon KClipboardIcon = 'F'<<24 | 'N'<<16 | 'D'<<8 | 'R' // 'FNDR'
	KFontSuitcaseIcon KClipboardIcon = 'F'<<24 | 'F'<<16 | 'I'<<8 | 'L' // 'FFIL'
	KFullTrashIcon KClipboardIcon = 'f'<<24 | 't'<<16 | 'r'<<8 | 'h' // 'ftrh'
	KGenericApplicationIcon KClipboardIcon = 'A'<<24 | 'P'<<16 | 'P'<<8 | 'L' // 'APPL'
	KGenericCDROMIcon KClipboardIcon = 'c'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'cddr'
	KGenericComponentIcon KClipboardIcon = 't'<<24 | 'h'<<16 | 'n'<<8 | 'g' // 'thng'
	KGenericControlPanelIcon KClipboardIcon = 'A'<<24 | 'P'<<16 | 'P'<<8 | 'C' // 'APPC'
	KGenericControlStripModuleIcon KClipboardIcon = 's'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'sdev'
	KGenericDeskAccessoryIcon KClipboardIcon = 'A'<<24 | 'P'<<16 | 'P'<<8 | 'D' // 'APPD'
	KGenericDocumentIcon KClipboardIcon = 'd'<<24 | 'o'<<16 | 'c'<<8 | 'u' // 'docu'
	KGenericEditionFileIcon KClipboardIcon = 'e'<<24 | 'd'<<16 | 't'<<8 | 'f' // 'edtf'
	KGenericExtensionIcon KClipboardIcon = 'I'<<24 | 'N'<<16 | 'I'<<8 | 'T' // 'INIT'
	KGenericFileServerIcon KClipboardIcon = 's'<<24 | 'r'<<16 | 'v'<<8 | 'r' // 'srvr'
	KGenericFloppyIcon KClipboardIcon = 'f'<<24 | 'l'<<16 | 'p'<<8 | 'y' // 'flpy'
	KGenericFontIcon KClipboardIcon = 'f'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'ffil'
	KGenericFontScalerIcon KClipboardIcon = 's'<<24 | 'c'<<16 | 'l'<<8 | 'r' // 'sclr'
	KGenericHardDiskIcon KClipboardIcon = 'h'<<24 | 'd'<<16 | 's'<<8 | 'k' // 'hdsk'
	KGenericIDiskIcon KClipboardIcon = 'i'<<24 | 'd'<<16 | 's'<<8 | 'k' // 'idsk'
	KGenericMoverObjectIcon KClipboardIcon = 'm'<<24 | 'o'<<16 | 'v'<<8 | 'r' // 'movr'
	KGenericPCCardIcon KClipboardIcon = 'p'<<24 | 'c'<<16 | 'm'<<8 | 'c' // 'pcmc'
	KGenericPreferencesIcon KClipboardIcon = 'p'<<24 | 'r'<<16 | 'e'<<8 | 'f' // 'pref'
	KGenericQueryDocumentIcon KClipboardIcon = 'q'<<24 | 'e'<<16 | 'r'<<8 | 'y' // 'qery'
	KGenericRAMDiskIcon KClipboardIcon = 'r'<<24 | 'a'<<16 | 'm'<<8 | 'd' // 'ramd'
	KGenericRemovableMediaIcon KClipboardIcon = 'r'<<24 | 'm'<<16 | 'o'<<8 | 'v' // 'rmov'
	KGenericSharedLibaryIcon KClipboardIcon = 's'<<24 | 'h'<<16 | 'l'<<8 | 'b' // 'shlb'
	KGenericStationeryIcon KClipboardIcon = 's'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'sdoc'
	KGenericSuitcaseIcon KClipboardIcon = 's'<<24 | 'u'<<16 | 'i'<<8 | 't' // 'suit'
	KGenericURLIcon KClipboardIcon = 'g'<<24 | 'u'<<16 | 'r'<<8 | 'l' // 'gurl'
	KGenericWORMIcon KClipboardIcon = 'w'<<24 | 'o'<<16 | 'r'<<8 | 'm' // 'worm'
	KInternationResourcesIcon KClipboardIcon = 'i'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'ifil'
	KInternationalResourcesIcon KClipboardIcon = 'i'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'ifil'
	KKeyboardLayoutIcon KClipboardIcon = 'k'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'kfil'
	KSoundFileIcon KClipboardIcon = 's'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'sfil'
	KSystemSuitcaseIcon KClipboardIcon = 'z'<<24 | 's'<<16 | 'y'<<8 | 's' // 'zsys'
	KTrashIcon KClipboardIcon = 't'<<24 | 'r'<<16 | 's'<<8 | 'h' // 'trsh'
	KTrueTypeFlatFontIcon KClipboardIcon = 's'<<24 | 'f'<<16 | 'n'<<8 | 't' // 'sfnt'
	KTrueTypeFontIcon KClipboardIcon = 't'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'tfil'
	KTrueTypeMultiFlatFontIcon KClipboardIcon = 't'<<24 | 't'<<16 | 'c'<<8 | 'f' // 'ttcf'
	KUnknownFSObjectIcon KClipboardIcon = 'u'<<24 | 'n'<<16 | 'f'<<8 | 's' // 'unfs'
	KUserIDiskIcon KClipboardIcon = 'u'<<24 | 'd'<<16 | 's'<<8 | 'k' // 'udsk'
)

func (e KClipboardIcon) String() string {
	switch e {
	case KClipboardIconValue:
		return "KClipboardIconValue"
	case KClippingPictureTypeIcon:
		return "KClippingPictureTypeIcon"
	case KClippingSoundTypeIcon:
		return "KClippingSoundTypeIcon"
	case KClippingTextTypeIcon:
		return "KClippingTextTypeIcon"
	case KClippingUnknownTypeIcon:
		return "KClippingUnknownTypeIcon"
	case KComputerIcon:
		return "KComputerIcon"
	case KDesktopIcon:
		return "KDesktopIcon"
	case KFinderIcon:
		return "KFinderIcon"
	case KFontSuitcaseIcon:
		return "KFontSuitcaseIcon"
	case KFullTrashIcon:
		return "KFullTrashIcon"
	case KGenericApplicationIcon:
		return "KGenericApplicationIcon"
	case KGenericCDROMIcon:
		return "KGenericCDROMIcon"
	case KGenericComponentIcon:
		return "KGenericComponentIcon"
	case KGenericControlPanelIcon:
		return "KGenericControlPanelIcon"
	case KGenericControlStripModuleIcon:
		return "KGenericControlStripModuleIcon"
	case KGenericDeskAccessoryIcon:
		return "KGenericDeskAccessoryIcon"
	case KGenericDocumentIcon:
		return "KGenericDocumentIcon"
	case KGenericEditionFileIcon:
		return "KGenericEditionFileIcon"
	case KGenericExtensionIcon:
		return "KGenericExtensionIcon"
	case KGenericFileServerIcon:
		return "KGenericFileServerIcon"
	case KGenericFloppyIcon:
		return "KGenericFloppyIcon"
	case KGenericFontIcon:
		return "KGenericFontIcon"
	case KGenericFontScalerIcon:
		return "KGenericFontScalerIcon"
	case KGenericHardDiskIcon:
		return "KGenericHardDiskIcon"
	case KGenericIDiskIcon:
		return "KGenericIDiskIcon"
	case KGenericMoverObjectIcon:
		return "KGenericMoverObjectIcon"
	case KGenericPCCardIcon:
		return "KGenericPCCardIcon"
	case KGenericPreferencesIcon:
		return "KGenericPreferencesIcon"
	case KGenericQueryDocumentIcon:
		return "KGenericQueryDocumentIcon"
	case KGenericRAMDiskIcon:
		return "KGenericRAMDiskIcon"
	case KGenericRemovableMediaIcon:
		return "KGenericRemovableMediaIcon"
	case KGenericSharedLibaryIcon:
		return "KGenericSharedLibaryIcon"
	case KGenericStationeryIcon:
		return "KGenericStationeryIcon"
	case KGenericSuitcaseIcon:
		return "KGenericSuitcaseIcon"
	case KGenericURLIcon:
		return "KGenericURLIcon"
	case KGenericWORMIcon:
		return "KGenericWORMIcon"
	case KInternationResourcesIcon:
		return "KInternationResourcesIcon"
	case KKeyboardLayoutIcon:
		return "KKeyboardLayoutIcon"
	case KSoundFileIcon:
		return "KSoundFileIcon"
	case KSystemSuitcaseIcon:
		return "KSystemSuitcaseIcon"
	case KTrashIcon:
		return "KTrashIcon"
	case KTrueTypeFlatFontIcon:
		return "KTrueTypeFlatFontIcon"
	case KTrueTypeFontIcon:
		return "KTrueTypeFontIcon"
	case KTrueTypeMultiFlatFontIcon:
		return "KTrueTypeMultiFlatFontIcon"
	case KUnknownFSObjectIcon:
		return "KUnknownFSObjectIcon"
	case KUserIDiskIcon:
		return "KUserIDiskIcon"
	default:
		return fmt.Sprintf("KClipboardIcon(%d)", e)
	}
}

type KCollate int

const (
	KCollateAttributesNotFoundErr KCollate = -29500
	KCollateBufferTooSmall KCollate = -29506
	KCollateInvalidChar KCollate = -29505
	KCollateInvalidCollationRef KCollate = -29507
	KCollateInvalidOptions KCollate = -29501
	KCollateMissingUnicodeTableErr KCollate = -29502
	KCollatePatternNotFoundErr KCollate = -29504
	KCollateUnicodeConvertFailedErr KCollate = -29503
)

func (e KCollate) String() string {
	switch e {
	case KCollateAttributesNotFoundErr:
		return "KCollateAttributesNotFoundErr"
	case KCollateBufferTooSmall:
		return "KCollateBufferTooSmall"
	case KCollateInvalidChar:
		return "KCollateInvalidChar"
	case KCollateInvalidCollationRef:
		return "KCollateInvalidCollationRef"
	case KCollateInvalidOptions:
		return "KCollateInvalidOptions"
	case KCollateMissingUnicodeTableErr:
		return "KCollateMissingUnicodeTableErr"
	case KCollatePatternNotFoundErr:
		return "KCollatePatternNotFoundErr"
	case KCollateUnicodeConvertFailedErr:
		return "KCollateUnicodeConvertFailedErr"
	default:
		return fmt.Sprintf("KCollate(%d)", e)
	}
}

const KCoreEventClass uint = 'a'<<24 | 'e'<<16 | 'v'<<8 | 't' // 'aevt'

type KDMGenErr int

const (
	KDMCantBlock KDMGenErr = -6224
	KDMDisplayAlreadyInstalledErr KDMGenErr = -6230
	KDMDisplayNotFoundErr KDMGenErr = -6229
	KDMDriverNotDisplayMgrAwareErr KDMGenErr = -6228
	KDMFoundErr KDMGenErr = -6232
	KDMGenErrValue KDMGenErr = -6220
	KDMMainDisplayCannotMoveErr KDMGenErr = -6231
	KDMMirroringBlocked KDMGenErr = -6223
	KDMMirroringNotOn KDMGenErr = -6225
	KDMMirroringOnAlready KDMGenErr = -6221
	KDMNoDeviceTableclothErr KDMGenErr = -6231
	KDMNotFoundErr KDMGenErr = -6229
	KDMSWNotInitializedErr KDMGenErr = -6227
	KDMWrongNumberOfDisplays KDMGenErr = -6222
	KSysSWTooOld KDMGenErr = -6226
)

func (e KDMGenErr) String() string {
	switch e {
	case KDMCantBlock:
		return "KDMCantBlock"
	case KDMDisplayAlreadyInstalledErr:
		return "KDMDisplayAlreadyInstalledErr"
	case KDMDisplayNotFoundErr:
		return "KDMDisplayNotFoundErr"
	case KDMDriverNotDisplayMgrAwareErr:
		return "KDMDriverNotDisplayMgrAwareErr"
	case KDMFoundErr:
		return "KDMFoundErr"
	case KDMGenErrValue:
		return "KDMGenErrValue"
	case KDMMainDisplayCannotMoveErr:
		return "KDMMainDisplayCannotMoveErr"
	case KDMMirroringBlocked:
		return "KDMMirroringBlocked"
	case KDMMirroringNotOn:
		return "KDMMirroringNotOn"
	case KDMMirroringOnAlready:
		return "KDMMirroringOnAlready"
	case KDMSWNotInitializedErr:
		return "KDMSWNotInitializedErr"
	case KDMWrongNumberOfDisplays:
		return "KDMWrongNumberOfDisplays"
	case KSysSWTooOld:
		return "KSysSWTooOld"
	default:
		return fmt.Sprintf("KDMGenErr(%d)", e)
	}
}

type KDSp int

const (
	KDSpConfirmSwitchWarning KDSp = -30448
	KDSpContextAlreadyReservedErr KDSp = -30444
	KDSpContextNotFoundErr KDSp = -30446
	KDSpContextNotReservedErr KDSp = -30445
	KDSpFrameRateNotReadyErr KDSp = -30447
	KDSpInternalErr KDSp = -30449
	KDSpInvalidAttributesErr KDSp = -30443
	KDSpInvalidContextErr KDSp = -30442
	KDSpNotInitializedErr KDSp = -30440
	KDSpStereoContextErr KDSp = -30450
	KDSpSystemSWTooOldErr KDSp = -30441
)

func (e KDSp) String() string {
	switch e {
	case KDSpConfirmSwitchWarning:
		return "KDSpConfirmSwitchWarning"
	case KDSpContextAlreadyReservedErr:
		return "KDSpContextAlreadyReservedErr"
	case KDSpContextNotFoundErr:
		return "KDSpContextNotFoundErr"
	case KDSpContextNotReservedErr:
		return "KDSpContextNotReservedErr"
	case KDSpFrameRateNotReadyErr:
		return "KDSpFrameRateNotReadyErr"
	case KDSpInternalErr:
		return "KDSpInternalErr"
	case KDSpInvalidAttributesErr:
		return "KDSpInvalidAttributesErr"
	case KDSpInvalidContextErr:
		return "KDSpInvalidContextErr"
	case KDSpNotInitializedErr:
		return "KDSpNotInitializedErr"
	case KDSpStereoContextErr:
		return "KDSpStereoContextErr"
	case KDSpSystemSWTooOldErr:
		return "KDSpSystemSWTooOldErr"
	default:
		return fmt.Sprintf("KDSp(%d)", e)
	}
}

type KDTP int

const (
	KDTPAbortJobErr KDTP = 128
	KDTPHoldJobErr KDTP = -4200
	KDTPStopQueueErr KDTP = -4201
	KDTPTryAgainErr KDTP = -4202
)

func (e KDTP) String() string {
	switch e {
	case KDTPAbortJobErr:
		return "KDTPAbortJobErr"
	case KDTPHoldJobErr:
		return "KDTPHoldJobErr"
	case KDTPStopQueueErr:
		return "KDTPStopQueueErr"
	case KDTPTryAgainErr:
		return "KDTPTryAgainErr"
	default:
		return fmt.Sprintf("KDTP(%d)", e)
	}
}

type KDesktopIconResource int

const (
	KDesktopIconResourceValue KDesktopIconResource = -3992
	KGenericFileServerIconResource KDesktopIconResource = -3972
	KGenericHardDiskIconResource KDesktopIconResource = -3995
	KGenericMoverObjectIconResource KDesktopIconResource = -3969
	KGenericSuitcaseIconResource KDesktopIconResource = -3970
	KOpenFolderIconResource KDesktopIconResource = -3997
)

func (e KDesktopIconResource) String() string {
	switch e {
	case KDesktopIconResourceValue:
		return "KDesktopIconResourceValue"
	case KGenericFileServerIconResource:
		return "KGenericFileServerIconResource"
	case KGenericHardDiskIconResource:
		return "KGenericHardDiskIconResource"
	case KGenericMoverObjectIconResource:
		return "KGenericMoverObjectIconResource"
	case KGenericSuitcaseIconResource:
		return "KGenericSuitcaseIconResource"
	case KOpenFolderIconResource:
		return "KOpenFolderIconResource"
	default:
		return fmt.Sprintf("KDesktopIconResource(%d)", e)
	}
}

type KFB int

const (
	KFBCaccessCanceled KFB = -30521
	KFBCaccessorStoreFailed KFB = -30515
	KFBCaddDocFailed KFB = -30516
	KFBCallocFailed KFB = -30502
	KFBCanalysisNotAvailable KFB = -30527
	KFBCbadIndexFile KFB = -30505
	KFBCbadIndexFileVersion KFB = -30528
	KFBCbadParam KFB = -30503
	KFBCbadSearchSession KFB = -30531
	KFBCcommitFailed KFB = -30509
	KFBCcompactionFailed KFB = -30506
	KFBCdeletionFailed KFB = -30510
	KFBCfileNotIndexed KFB = -30504
	KFBCflushFailed KFB = -30517
	KFBCillegalSessionChange KFB = -30526
	KFBCindexCreationFailed KFB = -30514
	KFBCindexDiskIOFailed KFB = -30530
	KFBCindexFileDestroyed KFB = -30522
	KFBCindexNotAvailable KFB = -30523
	KFBCindexNotFound KFB = -30518
	KFBCindexingCanceled KFB = -30520
	KFBCindexingFailed KFB = -30508
	KFBCmergingFailed KFB = -30513
	KFBCmoveFailed KFB = -30511
	KFBCnoIndexesFound KFB = -30501
	KFBCnoSearchSession KFB = -30519
	KFBCnoSuchHit KFB = -30532
	KFBCsearchFailed KFB = -30524
	KFBCsomeFilesNotIndexed KFB = -30525
	KFBCsummarizationCanceled KFB = -30529
	KFBCtokenizationFailed KFB = -30512
	KFBCvTwinExceptionErr KFB = -30500
	KFBCvalidationFailed KFB = -30507
)

func (e KFB) String() string {
	switch e {
	case KFBCaccessCanceled:
		return "KFBCaccessCanceled"
	case KFBCaccessorStoreFailed:
		return "KFBCaccessorStoreFailed"
	case KFBCaddDocFailed:
		return "KFBCaddDocFailed"
	case KFBCallocFailed:
		return "KFBCallocFailed"
	case KFBCanalysisNotAvailable:
		return "KFBCanalysisNotAvailable"
	case KFBCbadIndexFile:
		return "KFBCbadIndexFile"
	case KFBCbadIndexFileVersion:
		return "KFBCbadIndexFileVersion"
	case KFBCbadParam:
		return "KFBCbadParam"
	case KFBCbadSearchSession:
		return "KFBCbadSearchSession"
	case KFBCcommitFailed:
		return "KFBCcommitFailed"
	case KFBCcompactionFailed:
		return "KFBCcompactionFailed"
	case KFBCdeletionFailed:
		return "KFBCdeletionFailed"
	case KFBCfileNotIndexed:
		return "KFBCfileNotIndexed"
	case KFBCflushFailed:
		return "KFBCflushFailed"
	case KFBCillegalSessionChange:
		return "KFBCillegalSessionChange"
	case KFBCindexCreationFailed:
		return "KFBCindexCreationFailed"
	case KFBCindexDiskIOFailed:
		return "KFBCindexDiskIOFailed"
	case KFBCindexFileDestroyed:
		return "KFBCindexFileDestroyed"
	case KFBCindexNotAvailable:
		return "KFBCindexNotAvailable"
	case KFBCindexNotFound:
		return "KFBCindexNotFound"
	case KFBCindexingCanceled:
		return "KFBCindexingCanceled"
	case KFBCindexingFailed:
		return "KFBCindexingFailed"
	case KFBCmergingFailed:
		return "KFBCmergingFailed"
	case KFBCmoveFailed:
		return "KFBCmoveFailed"
	case KFBCnoIndexesFound:
		return "KFBCnoIndexesFound"
	case KFBCnoSearchSession:
		return "KFBCnoSearchSession"
	case KFBCnoSuchHit:
		return "KFBCnoSuchHit"
	case KFBCsearchFailed:
		return "KFBCsearchFailed"
	case KFBCsomeFilesNotIndexed:
		return "KFBCsomeFilesNotIndexed"
	case KFBCsummarizationCanceled:
		return "KFBCsummarizationCanceled"
	case KFBCtokenizationFailed:
		return "KFBCtokenizationFailed"
	case KFBCvTwinExceptionErr:
		return "KFBCvTwinExceptionErr"
	case KFBCvalidationFailed:
		return "KFBCvalidationFailed"
	default:
		return fmt.Sprintf("KFB(%d)", e)
	}
}

type KFM int

const (
	KFMFontContainerAccessErr KFM = -985
	KFMFontTableAccessErr KFM = -984
	KFMInvalidFontErr KFM = -982
	KFMInvalidFontFamilyErr KFM = -981
	KFMIterationCompleted KFM = -980
	KFMIterationScopeModifiedErr KFM = -983
)

func (e KFM) String() string {
	switch e {
	case KFMFontContainerAccessErr:
		return "KFMFontContainerAccessErr"
	case KFMFontTableAccessErr:
		return "KFMFontTableAccessErr"
	case KFMInvalidFontErr:
		return "KFMInvalidFontErr"
	case KFMInvalidFontFamilyErr:
		return "KFMInvalidFontFamilyErr"
	case KFMIterationCompleted:
		return "KFMIterationCompleted"
	case KFMIterationScopeModifiedErr:
		return "KFMIterationScopeModifiedErr"
	default:
		return fmt.Sprintf("KFM(%d)", e)
	}
}

type KFNS int

const (
	KFNSBadFlattenedSizeErr KFNS = -29587
	KFNSBadProfileVersionErr KFNS = -29583
	KFNSBadReferenceVersionErr KFNS = -29581
	KFNSDuplicateReferenceErr KFNS = -29584
	KFNSInsufficientDataErr KFNS = -29586
	KFNSInvalidProfileErr KFNS = -29582
	KFNSInvalidReferenceErr KFNS = -29580
	KFNSMismatchErr KFNS = -29585
	KFNSNameNotFoundErr KFNS = -29589
)

func (e KFNS) String() string {
	switch e {
	case KFNSBadFlattenedSizeErr:
		return "KFNSBadFlattenedSizeErr"
	case KFNSBadProfileVersionErr:
		return "KFNSBadProfileVersionErr"
	case KFNSBadReferenceVersionErr:
		return "KFNSBadReferenceVersionErr"
	case KFNSDuplicateReferenceErr:
		return "KFNSDuplicateReferenceErr"
	case KFNSInsufficientDataErr:
		return "KFNSInsufficientDataErr"
	case KFNSInvalidProfileErr:
		return "KFNSInvalidProfileErr"
	case KFNSInvalidReferenceErr:
		return "KFNSInvalidReferenceErr"
	case KFNSMismatchErr:
		return "KFNSMismatchErr"
	case KFNSNameNotFoundErr:
		return "KFNSNameNotFoundErr"
	default:
		return fmt.Sprintf("KFNS(%d)", e)
	}
}

type KFSEventStreamCreateFlag uint

const (
	// KFSEventStreamCreateFlagFileEvents: # Discussion
	KFSEventStreamCreateFlagFileEvents KFSEventStreamCreateFlag = 16
	// KFSEventStreamCreateFlagIgnoreSelf: # Discussion
	KFSEventStreamCreateFlagIgnoreSelf KFSEventStreamCreateFlag = 8
	KFSEventStreamCreateFlagMarkSelf KFSEventStreamCreateFlag = 32
	// KFSEventStreamCreateFlagNoDefer: # Discussion
	KFSEventStreamCreateFlagNoDefer KFSEventStreamCreateFlag = 2
	// KFSEventStreamCreateFlagNone: # Discussion
	KFSEventStreamCreateFlagNone KFSEventStreamCreateFlag = 0
	// KFSEventStreamCreateFlagUseCFTypes: # Discussion
	KFSEventStreamCreateFlagUseCFTypes KFSEventStreamCreateFlag = 1
	// KFSEventStreamCreateFlagWatchRoot: # Discussion
	KFSEventStreamCreateFlagWatchRoot KFSEventStreamCreateFlag = 4
)

func (e KFSEventStreamCreateFlag) String() string {
	switch e {
	case KFSEventStreamCreateFlagFileEvents:
		return "KFSEventStreamCreateFlagFileEvents"
	case KFSEventStreamCreateFlagIgnoreSelf:
		return "KFSEventStreamCreateFlagIgnoreSelf"
	case KFSEventStreamCreateFlagMarkSelf:
		return "KFSEventStreamCreateFlagMarkSelf"
	case KFSEventStreamCreateFlagNoDefer:
		return "KFSEventStreamCreateFlagNoDefer"
	case KFSEventStreamCreateFlagNone:
		return "KFSEventStreamCreateFlagNone"
	case KFSEventStreamCreateFlagUseCFTypes:
		return "KFSEventStreamCreateFlagUseCFTypes"
	case KFSEventStreamCreateFlagWatchRoot:
		return "KFSEventStreamCreateFlagWatchRoot"
	default:
		return fmt.Sprintf("KFSEventStreamCreateFlag(%d)", e)
	}
}

type KFSEventStreamEventFlag uint

const (
	// KFSEventStreamEventFlagEventIdsWrapped: # Discussion
	KFSEventStreamEventFlagEventIdsWrapped KFSEventStreamEventFlag = 8
	// KFSEventStreamEventFlagHistoryDone: # Discussion
	KFSEventStreamEventFlagHistoryDone KFSEventStreamEventFlag = 16
	KFSEventStreamEventFlagItemChangeOwner KFSEventStreamEventFlag = 16384
	KFSEventStreamEventFlagItemCreated KFSEventStreamEventFlag = 256
	KFSEventStreamEventFlagItemFinderInfoMod KFSEventStreamEventFlag = 8192
	KFSEventStreamEventFlagItemInodeMetaMod KFSEventStreamEventFlag = 1024
	KFSEventStreamEventFlagItemIsDir KFSEventStreamEventFlag = 131072
	KFSEventStreamEventFlagItemIsFile KFSEventStreamEventFlag = 65536
	KFSEventStreamEventFlagItemIsHardlink KFSEventStreamEventFlag = 1048576
	KFSEventStreamEventFlagItemIsLastHardlink KFSEventStreamEventFlag = 2097152
	KFSEventStreamEventFlagItemIsSymlink KFSEventStreamEventFlag = 262144
	KFSEventStreamEventFlagItemModified KFSEventStreamEventFlag = 4096
	KFSEventStreamEventFlagItemRemoved KFSEventStreamEventFlag = 512
	KFSEventStreamEventFlagItemRenamed KFSEventStreamEventFlag = 2048
	KFSEventStreamEventFlagItemXattrMod KFSEventStreamEventFlag = 32768
	KFSEventStreamEventFlagKernelDropped KFSEventStreamEventFlag = 4
	// KFSEventStreamEventFlagMount: # Discussion
	KFSEventStreamEventFlagMount KFSEventStreamEventFlag = 64
	// KFSEventStreamEventFlagMustScanSubDirs: # Discussion
	KFSEventStreamEventFlagMustScanSubDirs KFSEventStreamEventFlag = 1
	// KFSEventStreamEventFlagNone: # Discussion
	KFSEventStreamEventFlagNone KFSEventStreamEventFlag = 0
	KFSEventStreamEventFlagOwnEvent KFSEventStreamEventFlag = 524288
	// KFSEventStreamEventFlagRootChanged: # Discussion
	KFSEventStreamEventFlagRootChanged KFSEventStreamEventFlag = 32
	// KFSEventStreamEventFlagUnmount: # Discussion
	KFSEventStreamEventFlagUnmount KFSEventStreamEventFlag = 128
	// KFSEventStreamEventFlagUserDropped: # Discussion
	KFSEventStreamEventFlagUserDropped KFSEventStreamEventFlag = 2
)

func (e KFSEventStreamEventFlag) String() string {
	switch e {
	case KFSEventStreamEventFlagEventIdsWrapped:
		return "KFSEventStreamEventFlagEventIdsWrapped"
	case KFSEventStreamEventFlagHistoryDone:
		return "KFSEventStreamEventFlagHistoryDone"
	case KFSEventStreamEventFlagItemChangeOwner:
		return "KFSEventStreamEventFlagItemChangeOwner"
	case KFSEventStreamEventFlagItemCreated:
		return "KFSEventStreamEventFlagItemCreated"
	case KFSEventStreamEventFlagItemFinderInfoMod:
		return "KFSEventStreamEventFlagItemFinderInfoMod"
	case KFSEventStreamEventFlagItemInodeMetaMod:
		return "KFSEventStreamEventFlagItemInodeMetaMod"
	case KFSEventStreamEventFlagItemIsDir:
		return "KFSEventStreamEventFlagItemIsDir"
	case KFSEventStreamEventFlagItemIsFile:
		return "KFSEventStreamEventFlagItemIsFile"
	case KFSEventStreamEventFlagItemIsHardlink:
		return "KFSEventStreamEventFlagItemIsHardlink"
	case KFSEventStreamEventFlagItemIsLastHardlink:
		return "KFSEventStreamEventFlagItemIsLastHardlink"
	case KFSEventStreamEventFlagItemIsSymlink:
		return "KFSEventStreamEventFlagItemIsSymlink"
	case KFSEventStreamEventFlagItemModified:
		return "KFSEventStreamEventFlagItemModified"
	case KFSEventStreamEventFlagItemRemoved:
		return "KFSEventStreamEventFlagItemRemoved"
	case KFSEventStreamEventFlagItemRenamed:
		return "KFSEventStreamEventFlagItemRenamed"
	case KFSEventStreamEventFlagItemXattrMod:
		return "KFSEventStreamEventFlagItemXattrMod"
	case KFSEventStreamEventFlagKernelDropped:
		return "KFSEventStreamEventFlagKernelDropped"
	case KFSEventStreamEventFlagMount:
		return "KFSEventStreamEventFlagMount"
	case KFSEventStreamEventFlagMustScanSubDirs:
		return "KFSEventStreamEventFlagMustScanSubDirs"
	case KFSEventStreamEventFlagNone:
		return "KFSEventStreamEventFlagNone"
	case KFSEventStreamEventFlagOwnEvent:
		return "KFSEventStreamEventFlagOwnEvent"
	case KFSEventStreamEventFlagRootChanged:
		return "KFSEventStreamEventFlagRootChanged"
	case KFSEventStreamEventFlagUnmount:
		return "KFSEventStreamEventFlagUnmount"
	case KFSEventStreamEventFlagUserDropped:
		return "KFSEventStreamEventFlagUserDropped"
	default:
		return fmt.Sprintf("KFSEventStreamEventFlag(%d)", e)
	}
}

const KFSEventStreamEventIdSinceNow uint = 0

type KGenericDocumentIconResource int

const (
	KFloppyIconResource KGenericDocumentIconResource = -3998
	KGenericApplicationIconResource KGenericDocumentIconResource = -3996
	KGenericCDROMIconResource KGenericDocumentIconResource = -3987
	KGenericDeskAccessoryIconResource KGenericDocumentIconResource = -3991
	KGenericDocumentIconResourceValue KGenericDocumentIconResource = -4000
	KGenericEditionFileIconResource KGenericDocumentIconResource = -3989
	KGenericFolderIconResource KGenericDocumentIconResource = -3999
	KGenericRAMDiskIconResource KGenericDocumentIconResource = -3988
	KGenericStationeryIconResource KGenericDocumentIconResource = -3985
	KPrivateFolderIconResource KGenericDocumentIconResource = -3994
	KTrashIconResource KGenericDocumentIconResource = -3993
)

func (e KGenericDocumentIconResource) String() string {
	switch e {
	case KFloppyIconResource:
		return "KFloppyIconResource"
	case KGenericApplicationIconResource:
		return "KGenericApplicationIconResource"
	case KGenericCDROMIconResource:
		return "KGenericCDROMIconResource"
	case KGenericDeskAccessoryIconResource:
		return "KGenericDeskAccessoryIconResource"
	case KGenericDocumentIconResourceValue:
		return "KGenericDocumentIconResourceValue"
	case KGenericEditionFileIconResource:
		return "KGenericEditionFileIconResource"
	case KGenericFolderIconResource:
		return "KGenericFolderIconResource"
	case KGenericRAMDiskIconResource:
		return "KGenericRAMDiskIconResource"
	case KGenericStationeryIconResource:
		return "KGenericStationeryIconResource"
	case KPrivateFolderIconResource:
		return "KPrivateFolderIconResource"
	case KTrashIconResource:
		return "KTrashIconResource"
	default:
		return fmt.Sprintf("KGenericDocumentIconResource(%d)", e)
	}
}

type KGenericFolderIcon uint

const (
	KDropFolderIcon KGenericFolderIcon = 'd'<<24 | 'b'<<16 | 'o'<<8 | 'x' // 'dbox'
	KGenericFolderIconValue KGenericFolderIcon = 'f'<<24 | 'l'<<16 | 'd'<<8 | 'r' // 'fldr'
	KMountedFolderIcon KGenericFolderIcon = 'm'<<24 | 'n'<<16 | 't'<<8 | 'd' // 'mntd'
	KOpenFolderIcon KGenericFolderIcon = 'o'<<24 | 'f'<<16 | 'l'<<8 | 'd' // 'ofld'
	KOwnedFolderIcon KGenericFolderIcon = 'o'<<24 | 'w'<<16 | 'n'<<8 | 'd' // 'ownd'
	KPrivateFolderIcon KGenericFolderIcon = 'p'<<24 | 'r'<<16 | 'v'<<8 | 'f' // 'prvf'
	KSharedFolderIcon KGenericFolderIcon = 's'<<24 | 'h'<<16 | 'f'<<8 | 'l' // 'shfl'
)

func (e KGenericFolderIcon) String() string {
	switch e {
	case KDropFolderIcon:
		return "KDropFolderIcon"
	case KGenericFolderIconValue:
		return "KGenericFolderIconValue"
	case KMountedFolderIcon:
		return "KMountedFolderIcon"
	case KOpenFolderIcon:
		return "KOpenFolderIcon"
	case KOwnedFolderIcon:
		return "KOwnedFolderIcon"
	case KPrivateFolderIcon:
		return "KPrivateFolderIcon"
	case KSharedFolderIcon:
		return "KSharedFolderIcon"
	default:
		return fmt.Sprintf("KGenericFolderIcon(%d)", e)
	}
}

type KGenericPreferencesIconResource int

const (
	KAppleMenuFolderIconResource KGenericPreferencesIconResource = -3982
	KGenericExtensionIconResource KGenericPreferencesIconResource = -16415
	KGenericPreferencesIconResourceValue KGenericPreferencesIconResource = -3971
	KGenericQueryDocumentIconResource KGenericPreferencesIconResource = -16506
	KHelpIconResource KGenericPreferencesIconResource = -20271
	KSystemFolderIconResource KGenericPreferencesIconResource = -3983
)

func (e KGenericPreferencesIconResource) String() string {
	switch e {
	case KAppleMenuFolderIconResource:
		return "KAppleMenuFolderIconResource"
	case KGenericExtensionIconResource:
		return "KGenericExtensionIconResource"
	case KGenericPreferencesIconResourceValue:
		return "KGenericPreferencesIconResourceValue"
	case KGenericQueryDocumentIconResource:
		return "KGenericQueryDocumentIconResource"
	case KHelpIconResource:
		return "KHelpIconResource"
	case KSystemFolderIconResource:
		return "KSystemFolderIconResource"
	default:
		return fmt.Sprintf("KGenericPreferencesIconResource(%d)", e)
	}
}

type KHID int

const (
	KHIDBadLogPhysValuesErr KHID = -13942
	KHIDBadLogicalMaximumErr KHID = -13933
	KHIDBadLogicalMinimumErr KHID = -13934
	KHIDBadParameterErr KHID = -13938
	KHIDBaseError KHID = -13950
	KHIDBufferTooSmallErr KHID = -13948
	KHIDDeviceNotReady KHID = -13910
	KHIDEndOfDescriptorErr KHID = -13936
	KHIDIncompatibleReportErr KHID = -13943
	KHIDInvalidPreparsedDataErr KHID = -13944
	KHIDInvalidRangePageErr KHID = -13923
	KHIDInvalidReportLengthErr KHID = -13940
	KHIDInvalidReportTypeErr KHID = -13941
	KHIDInvertedLogicalRangeErr KHID = -13932
	KHIDInvertedPhysicalRangeErr KHID = -13931
	KHIDInvertedUsageRangeErr KHID = -13929
	KHIDNotEnoughMemoryErr KHID = -13937
	KHIDNotValueArrayErr KHID = -13945
	KHIDNullPointerErr KHID = -13939
	KHIDNullStateErr KHID = -13949
	KHIDReportCountZeroErr KHID = -13925
	KHIDReportIDZeroErr KHID = -13924
	KHIDReportSizeZeroErr KHID = -13926
	KHIDSuccess KHID = 0
	KHIDUnmatchedDesignatorRangeErr KHID = -13927
	KHIDUnmatchedStringRangeErr KHID = -13928
	KHIDUnmatchedUsageRangeErr KHID = -13930
	KHIDUsageNotFoundErr KHID = -13946
	KHIDUsagePageZeroErr KHID = -13935
	KHIDValueOutOfRangeErr KHID = -13947
	KHIDVersionIncompatibleErr KHID = -13909
)

func (e KHID) String() string {
	switch e {
	case KHIDBadLogPhysValuesErr:
		return "KHIDBadLogPhysValuesErr"
	case KHIDBadLogicalMaximumErr:
		return "KHIDBadLogicalMaximumErr"
	case KHIDBadLogicalMinimumErr:
		return "KHIDBadLogicalMinimumErr"
	case KHIDBadParameterErr:
		return "KHIDBadParameterErr"
	case KHIDBaseError:
		return "KHIDBaseError"
	case KHIDBufferTooSmallErr:
		return "KHIDBufferTooSmallErr"
	case KHIDDeviceNotReady:
		return "KHIDDeviceNotReady"
	case KHIDEndOfDescriptorErr:
		return "KHIDEndOfDescriptorErr"
	case KHIDIncompatibleReportErr:
		return "KHIDIncompatibleReportErr"
	case KHIDInvalidPreparsedDataErr:
		return "KHIDInvalidPreparsedDataErr"
	case KHIDInvalidRangePageErr:
		return "KHIDInvalidRangePageErr"
	case KHIDInvalidReportLengthErr:
		return "KHIDInvalidReportLengthErr"
	case KHIDInvalidReportTypeErr:
		return "KHIDInvalidReportTypeErr"
	case KHIDInvertedLogicalRangeErr:
		return "KHIDInvertedLogicalRangeErr"
	case KHIDInvertedPhysicalRangeErr:
		return "KHIDInvertedPhysicalRangeErr"
	case KHIDInvertedUsageRangeErr:
		return "KHIDInvertedUsageRangeErr"
	case KHIDNotEnoughMemoryErr:
		return "KHIDNotEnoughMemoryErr"
	case KHIDNotValueArrayErr:
		return "KHIDNotValueArrayErr"
	case KHIDNullPointerErr:
		return "KHIDNullPointerErr"
	case KHIDNullStateErr:
		return "KHIDNullStateErr"
	case KHIDReportCountZeroErr:
		return "KHIDReportCountZeroErr"
	case KHIDReportIDZeroErr:
		return "KHIDReportIDZeroErr"
	case KHIDReportSizeZeroErr:
		return "KHIDReportSizeZeroErr"
	case KHIDSuccess:
		return "KHIDSuccess"
	case KHIDUnmatchedDesignatorRangeErr:
		return "KHIDUnmatchedDesignatorRangeErr"
	case KHIDUnmatchedStringRangeErr:
		return "KHIDUnmatchedStringRangeErr"
	case KHIDUnmatchedUsageRangeErr:
		return "KHIDUnmatchedUsageRangeErr"
	case KHIDUsageNotFoundErr:
		return "KHIDUsageNotFoundErr"
	case KHIDUsagePageZeroErr:
		return "KHIDUsagePageZeroErr"
	case KHIDValueOutOfRangeErr:
		return "KHIDValueOutOfRangeErr"
	case KHIDVersionIncompatibleErr:
		return "KHIDVersionIncompatibleErr"
	default:
		return fmt.Sprintf("KHID(%d)", e)
	}
}

type KISp int

const (
	KISpBufferToSmallErr KISp = -30422
	KISpDeviceActiveErr KISp = -30428
	KISpDeviceInactiveErr KISp = -30426
	KISpElementInListErr KISp = -30423
	KISpElementNotInListErr KISp = -30424
	KISpInternalErr KISp = -30420
	KISpListBusyErr KISp = -30429
	KISpSystemActiveErr KISp = -30427
	KISpSystemInactiveErr KISp = -30425
	KISpSystemListErr KISp = -30421
)

func (e KISp) String() string {
	switch e {
	case KISpBufferToSmallErr:
		return "KISpBufferToSmallErr"
	case KISpDeviceActiveErr:
		return "KISpDeviceActiveErr"
	case KISpDeviceInactiveErr:
		return "KISpDeviceInactiveErr"
	case KISpElementInListErr:
		return "KISpElementInListErr"
	case KISpElementNotInListErr:
		return "KISpElementNotInListErr"
	case KISpInternalErr:
		return "KISpInternalErr"
	case KISpListBusyErr:
		return "KISpListBusyErr"
	case KISpSystemActiveErr:
		return "KISpSystemActiveErr"
	case KISpSystemInactiveErr:
		return "KISpSystemInactiveErr"
	case KISpSystemListErr:
		return "KISpSystemListErr"
	default:
		return fmt.Sprintf("KISp(%d)", e)
	}
}

type KIconServices uint

const (
	KIconServicesNoBadgeFlag KIconServices = 1
	KIconServicesNormalUsageFlag KIconServices = 0
	KIconServicesUpdateIfNeededFlag KIconServices = 2
)

func (e KIconServices) String() string {
	switch e {
	case KIconServicesNoBadgeFlag:
		return "KIconServicesNoBadgeFlag"
	case KIconServicesNormalUsageFlag:
		return "KIconServicesNormalUsageFlag"
	case KIconServicesUpdateIfNeededFlag:
		return "KIconServicesUpdateIfNeededFlag"
	default:
		return fmt.Sprintf("KIconServices(%d)", e)
	}
}

const KIconServicesCatalogInfoMask uint = 0

type KInternetLocation uint

const (
	KInternetLocationAppleShareIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'a'<<8 | 'f' // 'ilaf'
	KInternetLocationAppleTalkZoneIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'a'<<8 | 't' // 'ilat'
	KInternetLocationFTPIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'f'<<8 | 't' // 'ilft'
	KInternetLocationFileIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'f'<<8 | 'i' // 'ilfi'
	KInternetLocationGenericIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'g'<<8 | 'e' // 'ilge'
	KInternetLocationHTTPIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'h'<<8 | 't' // 'ilht'
	KInternetLocationMailIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'm'<<8 | 'a' // 'ilma'
	KInternetLocationNSLNeighborhoodIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'n'<<8 | 's' // 'ilns'
	KInternetLocationNewsIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'n'<<8 | 'w' // 'ilnw'
)

func (e KInternetLocation) String() string {
	switch e {
	case KInternetLocationAppleShareIcon:
		return "KInternetLocationAppleShareIcon"
	case KInternetLocationAppleTalkZoneIcon:
		return "KInternetLocationAppleTalkZoneIcon"
	case KInternetLocationFTPIcon:
		return "KInternetLocationFTPIcon"
	case KInternetLocationFileIcon:
		return "KInternetLocationFileIcon"
	case KInternetLocationGenericIcon:
		return "KInternetLocationGenericIcon"
	case KInternetLocationHTTPIcon:
		return "KInternetLocationHTTPIcon"
	case KInternetLocationMailIcon:
		return "KInternetLocationMailIcon"
	case KInternetLocationNSLNeighborhoodIcon:
		return "KInternetLocationNSLNeighborhoodIcon"
	case KInternetLocationNewsIcon:
		return "KInternetLocationNewsIcon"
	default:
		return fmt.Sprintf("KInternetLocation(%d)", e)
	}
}

type KLS int

const (
	// KLSAppDoesNotClaimTypeErr: Not currently used.
	KLSAppDoesNotClaimTypeErr KLS = -10820
	// KLSAppDoesNotSupportSchemeWarning: Not currently used.
	KLSAppDoesNotSupportSchemeWarning KLS = -10821
	// KLSAppInTrashErr: The app can’t run because it’s inside a Trash folder.
	KLSAppInTrashErr KLS = -10660
	// KLSApplicationNotFoundErr: No app in the Launch Services database matches the input criteria.
	KLSApplicationNotFoundErr KLS = -10814
	// KLSCannotSetInfoErr: The system can’t hide the filename extension.
	KLSCannotSetInfoErr KLS = -10823
	// KLSDataErr: Improper data structure.
	KLSDataErr KLS = -10817
	// KLSDataTooOldErr: Not currently used.
	KLSDataTooOldErr KLS = -10816
	// KLSDataUnavailableErr: Data of the desired type is not available (for example, there is no kind string).
	KLSDataUnavailableErr KLS = -10813
	// KLSIncompatibleSystemVersionErr: The requested app can’t run on the current macOS version.
	KLSIncompatibleSystemVersionErr KLS = -10825
	// KLSLaunchInProgressErr: A launch of the app is already in progress.
	KLSLaunchInProgressErr KLS = -10818
	// KLSMultipleSessionsNotSupportedErr: The requested app can’t run simultaneously in two different user sessions.
	KLSMultipleSessionsNotSupportedErr KLS = -10829
	// KLSNoClassicEnvironmentErr: Not currently used.
	KLSNoClassicEnvironmentErr KLS = -10828
	// KLSNoExecutableErr: The executable file is missing or has an unusable format.
	KLSNoExecutableErr KLS = -10827
	// KLSNoLaunchPermissionErr: The user doesn’t have permission to launch the app on a managed network.
	KLSNoLaunchPermissionErr KLS = -10826
	// KLSNoRegistrationInfoErr: Not currently used.
	KLSNoRegistrationInfoErr KLS = -10824
	// KLSNotAnApplicationErr: The item for registration is not an app.
	KLSNotAnApplicationErr KLS = -10811
	// KLSNotInitializedErr: Not currently used.
	KLSNotInitializedErr KLS = -10812
	// KLSNotRegisteredErr: Not currently used.
	KLSNotRegisteredErr KLS = -10819
	// KLSServerCommunicationErr: There is a problem communicating with the server process that maintains the Launch Services database.
	KLSServerCommunicationErr KLS = -10822
	// KLSUnknownErr: An unknown error has occurred.
	KLSUnknownErr KLS = -10810
	// KLSUnknownTypeErr: Not currently used.
	KLSUnknownTypeErr KLS = -10815
)

func (e KLS) String() string {
	switch e {
	case KLSAppDoesNotClaimTypeErr:
		return "KLSAppDoesNotClaimTypeErr"
	case KLSAppDoesNotSupportSchemeWarning:
		return "KLSAppDoesNotSupportSchemeWarning"
	case KLSAppInTrashErr:
		return "KLSAppInTrashErr"
	case KLSApplicationNotFoundErr:
		return "KLSApplicationNotFoundErr"
	case KLSCannotSetInfoErr:
		return "KLSCannotSetInfoErr"
	case KLSDataErr:
		return "KLSDataErr"
	case KLSDataTooOldErr:
		return "KLSDataTooOldErr"
	case KLSDataUnavailableErr:
		return "KLSDataUnavailableErr"
	case KLSIncompatibleSystemVersionErr:
		return "KLSIncompatibleSystemVersionErr"
	case KLSLaunchInProgressErr:
		return "KLSLaunchInProgressErr"
	case KLSMultipleSessionsNotSupportedErr:
		return "KLSMultipleSessionsNotSupportedErr"
	case KLSNoClassicEnvironmentErr:
		return "KLSNoClassicEnvironmentErr"
	case KLSNoExecutableErr:
		return "KLSNoExecutableErr"
	case KLSNoLaunchPermissionErr:
		return "KLSNoLaunchPermissionErr"
	case KLSNoRegistrationInfoErr:
		return "KLSNoRegistrationInfoErr"
	case KLSNotAnApplicationErr:
		return "KLSNotAnApplicationErr"
	case KLSNotInitializedErr:
		return "KLSNotInitializedErr"
	case KLSNotRegisteredErr:
		return "KLSNotRegisteredErr"
	case KLSServerCommunicationErr:
		return "KLSServerCommunicationErr"
	case KLSUnknownErr:
		return "KLSUnknownErr"
	case KLSUnknownTypeErr:
		return "KLSUnknownTypeErr"
	default:
		return fmt.Sprintf("KLS(%d)", e)
	}
}

type KLSLaunch uint

const (
	// Deprecated.
	KLSLaunchHasUntrustedContents KLSLaunch = 4194304
	// Deprecated.
	KLSLaunchInClassic KLSLaunch = 262144
	// Deprecated.
	KLSLaunchInhibitBGOnly KLSLaunch = 128
	// Deprecated.
	KLSLaunchNoParams KLSLaunch = 2048
	// Deprecated.
	KLSLaunchStartClassic KLSLaunch = 131072
)

func (e KLSLaunch) String() string {
	switch e {
	case KLSLaunchHasUntrustedContents:
		return "KLSLaunchHasUntrustedContents"
	case KLSLaunchInClassic:
		return "KLSLaunchInClassic"
	case KLSLaunchInhibitBGOnly:
		return "KLSLaunchInhibitBGOnly"
	case KLSLaunchNoParams:
		return "KLSLaunchNoParams"
	case KLSLaunchStartClassic:
		return "KLSLaunchStartClassic"
	default:
		return fmt.Sprintf("KLSLaunch(%d)", e)
	}
}

type KLSSharedFileList uint

const (
	KLSSharedFileListDoNotMountVolumes KLSSharedFileList = 2
	KLSSharedFileListNoUserInteraction KLSSharedFileList = 1
)

func (e KLSSharedFileList) String() string {
	switch e {
	case KLSSharedFileListDoNotMountVolumes:
		return "KLSSharedFileListDoNotMountVolumes"
	case KLSSharedFileListNoUserInteraction:
		return "KLSSharedFileListNoUserInteraction"
	default:
		return fmt.Sprintf("KLSSharedFileList(%d)", e)
	}
}

type KLSUnknown uint

const (
	// KLSUnknownCreator: The value to supply as the creator signatureif no file creator information is available.
	KLSUnknownCreator KLSUnknown = 0
	// KLSUnknownType: The value to supply as the file type (for example,to the [LSGetApplicationForInfo] function)if no file type information is available.
	KLSUnknownType KLSUnknown = 0
)

func (e KLSUnknown) String() string {
	switch e {
	case KLSUnknownCreator:
		return "KLSUnknownCreator"
	default:
		return fmt.Sprintf("KLSUnknown(%d)", e)
	}
}

type KLocales int

const (
	KLocalesBufferTooSmallErr KLocales = -30001
	KLocalesDefaultDisplayStatus KLocales = -30029
	KLocalesTableFormatErr KLocales = -30002
)

func (e KLocales) String() string {
	switch e {
	case KLocalesBufferTooSmallErr:
		return "KLocalesBufferTooSmallErr"
	case KLocalesDefaultDisplayStatus:
		return "KLocalesDefaultDisplayStatus"
	case KLocalesTableFormatErr:
		return "KLocalesTableFormatErr"
	default:
		return fmt.Sprintf("KLocales(%d)", e)
	}
}

type KMP int

const (
	KMPBlueBlockingErr KMP = -29293
	// KMPDeletedErr: The desired notification the function was waiting upon was deleted.
	KMPDeletedErr KMP = -29295
	// KMPInsufficientResourcesErr: Could not complete task due to unavailable Multiprocessing Services resources.
	KMPInsufficientResourcesErr KMP = -29298
	// KMPInvalidIDErr: Invalid ID value.
	KMPInvalidIDErr KMP = -29299
	KMPIterationEndErr KMP = -29275
	KMPPrivilegedErr KMP = -29276
	KMPProcessCreatedErr KMP = -29288
	KMPProcessTerminatedErr KMP = -29289
	KMPTaskAbortedErr KMP = -29297
	// KMPTaskBlockedErr: The desired task is blocked.
	KMPTaskBlockedErr KMP = -29291
	KMPTaskCreatedErr KMP = -29290
	// KMPTaskStoppedErr: The desired task is stopped.
	KMPTaskStoppedErr KMP = -29292
	// KMPTimeoutErr: The designated timeout interval passed before the function could take action.
	KMPTimeoutErr KMP = -29296
)

func (e KMP) String() string {
	switch e {
	case KMPBlueBlockingErr:
		return "KMPBlueBlockingErr"
	case KMPDeletedErr:
		return "KMPDeletedErr"
	case KMPInsufficientResourcesErr:
		return "KMPInsufficientResourcesErr"
	case KMPInvalidIDErr:
		return "KMPInvalidIDErr"
	case KMPIterationEndErr:
		return "KMPIterationEndErr"
	case KMPPrivilegedErr:
		return "KMPPrivilegedErr"
	case KMPProcessCreatedErr:
		return "KMPProcessCreatedErr"
	case KMPProcessTerminatedErr:
		return "KMPProcessTerminatedErr"
	case KMPTaskAbortedErr:
		return "KMPTaskAbortedErr"
	case KMPTaskBlockedErr:
		return "KMPTaskBlockedErr"
	case KMPTaskCreatedErr:
		return "KMPTaskCreatedErr"
	case KMPTaskStoppedErr:
		return "KMPTaskStoppedErr"
	case KMPTimeoutErr:
		return "KMPTimeoutErr"
	default:
		return fmt.Sprintf("KMP(%d)", e)
	}
}

const KMPNanokernelNeedsMemoryErr int = -29294

type KModem int

const (
	KModemOutOfMemory KModem = -14000
	KModemPreferencesMissing KModem = -14001
	KModemScriptMissing KModem = -14002
)

func (e KModem) String() string {
	switch e {
	case KModemOutOfMemory:
		return "KModemOutOfMemory"
	case KModemPreferencesMissing:
		return "KModemPreferencesMissing"
	case KModemScriptMissing:
		return "KModemScriptMissing"
	default:
		return fmt.Sprintf("KModem(%d)", e)
	}
}

type KNS int

const (
	KNSL68kContextNotSupported KNS = -4170
	KNSLBadClientInfoPtr KNS = -4185
	KNSLBadDataTypeErr KNS = -4193
	KNSLBadNetConnection KNS = -4192
	KNSLBadProtocolTypeErr KNS = -4183
	KNSLBadReferenceErr KNS = -4195
	KNSLBadServiceTypeErr KNS = -4194
	KNSLBadURLSyntax KNS = -4172
	KNSLBufferTooSmallForData KNS = -4187
	KNSLCannotContinueLookup KNS = -4186
	KNSLErrNullPtrError KNS = -4176
	KNSLInitializationFailed KNS = -4200
	KNSLInsufficientOTVer KNS = -4197
	KNSLInsufficientSysVer KNS = -4198
	KNSLInvalidPluginSpec KNS = -4190
	KNSLNoCarbonLib KNS = -4173
	KNSLNoContextAvailable KNS = -4188
	KNSLNoElementsInList KNS = -4196
	KNSLNoPluginsForSearch KNS = -4179
	KNSLNoPluginsFound KNS = -4181
	KNSLNoSupportForService KNS = -4191
	KNSLNotImplementedYet KNS = -4175
	KNSLNotInitialized KNS = -4199
	KNSLNullListPtr KNS = -4184
	KNSLNullNeighborhoodPtr KNS = -4178
	KNSLPluginLoadFailed KNS = -4182
	KNSLRequestBufferAlreadyInList KNS = -4189
	KNSLSchedulerError KNS = -4171
	KNSLSearchAlreadyInProgress KNS = -4180
	KNSLSomePluginsFailedToLoad KNS = -4177
	KNSLUILibraryNotAvailable KNS = -4174
)

func (e KNS) String() string {
	switch e {
	case KNSL68kContextNotSupported:
		return "KNSL68kContextNotSupported"
	case KNSLBadClientInfoPtr:
		return "KNSLBadClientInfoPtr"
	case KNSLBadDataTypeErr:
		return "KNSLBadDataTypeErr"
	case KNSLBadNetConnection:
		return "KNSLBadNetConnection"
	case KNSLBadProtocolTypeErr:
		return "KNSLBadProtocolTypeErr"
	case KNSLBadReferenceErr:
		return "KNSLBadReferenceErr"
	case KNSLBadServiceTypeErr:
		return "KNSLBadServiceTypeErr"
	case KNSLBadURLSyntax:
		return "KNSLBadURLSyntax"
	case KNSLBufferTooSmallForData:
		return "KNSLBufferTooSmallForData"
	case KNSLCannotContinueLookup:
		return "KNSLCannotContinueLookup"
	case KNSLErrNullPtrError:
		return "KNSLErrNullPtrError"
	case KNSLInitializationFailed:
		return "KNSLInitializationFailed"
	case KNSLInsufficientOTVer:
		return "KNSLInsufficientOTVer"
	case KNSLInsufficientSysVer:
		return "KNSLInsufficientSysVer"
	case KNSLInvalidPluginSpec:
		return "KNSLInvalidPluginSpec"
	case KNSLNoCarbonLib:
		return "KNSLNoCarbonLib"
	case KNSLNoContextAvailable:
		return "KNSLNoContextAvailable"
	case KNSLNoElementsInList:
		return "KNSLNoElementsInList"
	case KNSLNoPluginsForSearch:
		return "KNSLNoPluginsForSearch"
	case KNSLNoPluginsFound:
		return "KNSLNoPluginsFound"
	case KNSLNoSupportForService:
		return "KNSLNoSupportForService"
	case KNSLNotImplementedYet:
		return "KNSLNotImplementedYet"
	case KNSLNotInitialized:
		return "KNSLNotInitialized"
	case KNSLNullListPtr:
		return "KNSLNullListPtr"
	case KNSLNullNeighborhoodPtr:
		return "KNSLNullNeighborhoodPtr"
	case KNSLPluginLoadFailed:
		return "KNSLPluginLoadFailed"
	case KNSLRequestBufferAlreadyInList:
		return "KNSLRequestBufferAlreadyInList"
	case KNSLSchedulerError:
		return "KNSLSchedulerError"
	case KNSLSearchAlreadyInProgress:
		return "KNSLSearchAlreadyInProgress"
	case KNSLSomePluginsFailedToLoad:
		return "KNSLSomePluginsFailedToLoad"
	case KNSLUILibraryNotAvailable:
		return "KNSLUILibraryNotAvailable"
	default:
		return fmt.Sprintf("KNS(%d)", e)
	}
}

type KNSp int

const (
	KNSpAddPlayerFailedErr KNSp = -30389
	KNSpAddressInUseErr KNSp = -30380
	KNSpAlreadyAdvertisingErr KNSp = -30374
	KNSpAlreadyInitializedErr KNSp = -30361
	KNSpCantBlockErr KNSp = -30398
	KNSpConnectFailedErr KNSp = -30395
	KNSpCreateGroupFailedErr KNSp = -30388
	KNSpFeatureNotImplementedErr KNSp = -30381
	KNSpFreeQExhaustedErr KNSp = -30378
	KNSpGameTerminatedErr KNSp = -30394
	KNSpHostFailedErr KNSp = -30365
	KNSpInitializationFailedErr KNSp = -30360
	KNSpInvalidAddressErr KNSp = -30377
	KNSpInvalidDefinitionErr KNSp = -30390
	KNSpInvalidGameRefErr KNSp = -30367
	KNSpInvalidGroupIDErr KNSp = -30384
	KNSpInvalidParameterErr KNSp = -30369
	KNSpInvalidPlayerIDErr KNSp = -30383
	KNSpInvalidProtocolListErr KNSp = -30392
	KNSpInvalidProtocolRefErr KNSp = -30391
	KNSpJoinFailedErr KNSp = -30399
	KNSpMemAllocationErr KNSp = -30373
	KNSpMessageTooBigErr KNSp = -30397
	KNSpNameRequiredErr KNSp = -30382
	KNSpNoGroupsErr KNSp = -30386
	KNSpNoHostVolunteersErr KNSp = -30387
	KNSpNoPlayersErr KNSp = -30385
	KNSpNotAdvertisingErr KNSp = -30376
	KNSpOTNotPresentErr KNSp = -30370
	KNSpOTVersionTooOldErr KNSp = -30371
	KNSpPipeFullErr KNSp = -30364
	KNSpProtocolNotAvailableErr KNSp = -30366
	KNSpRemovePlayerFailedErr KNSp = -30379
	KNSpSendFailedErr KNSp = -30396
	KNSpTimeoutErr KNSp = -30393
	KNSpTopologyNotSupportedErr KNSp = -30362
)

func (e KNSp) String() string {
	switch e {
	case KNSpAddPlayerFailedErr:
		return "KNSpAddPlayerFailedErr"
	case KNSpAddressInUseErr:
		return "KNSpAddressInUseErr"
	case KNSpAlreadyAdvertisingErr:
		return "KNSpAlreadyAdvertisingErr"
	case KNSpAlreadyInitializedErr:
		return "KNSpAlreadyInitializedErr"
	case KNSpCantBlockErr:
		return "KNSpCantBlockErr"
	case KNSpConnectFailedErr:
		return "KNSpConnectFailedErr"
	case KNSpCreateGroupFailedErr:
		return "KNSpCreateGroupFailedErr"
	case KNSpFeatureNotImplementedErr:
		return "KNSpFeatureNotImplementedErr"
	case KNSpFreeQExhaustedErr:
		return "KNSpFreeQExhaustedErr"
	case KNSpGameTerminatedErr:
		return "KNSpGameTerminatedErr"
	case KNSpHostFailedErr:
		return "KNSpHostFailedErr"
	case KNSpInitializationFailedErr:
		return "KNSpInitializationFailedErr"
	case KNSpInvalidAddressErr:
		return "KNSpInvalidAddressErr"
	case KNSpInvalidDefinitionErr:
		return "KNSpInvalidDefinitionErr"
	case KNSpInvalidGameRefErr:
		return "KNSpInvalidGameRefErr"
	case KNSpInvalidGroupIDErr:
		return "KNSpInvalidGroupIDErr"
	case KNSpInvalidParameterErr:
		return "KNSpInvalidParameterErr"
	case KNSpInvalidPlayerIDErr:
		return "KNSpInvalidPlayerIDErr"
	case KNSpInvalidProtocolListErr:
		return "KNSpInvalidProtocolListErr"
	case KNSpInvalidProtocolRefErr:
		return "KNSpInvalidProtocolRefErr"
	case KNSpJoinFailedErr:
		return "KNSpJoinFailedErr"
	case KNSpMemAllocationErr:
		return "KNSpMemAllocationErr"
	case KNSpMessageTooBigErr:
		return "KNSpMessageTooBigErr"
	case KNSpNameRequiredErr:
		return "KNSpNameRequiredErr"
	case KNSpNoGroupsErr:
		return "KNSpNoGroupsErr"
	case KNSpNoHostVolunteersErr:
		return "KNSpNoHostVolunteersErr"
	case KNSpNoPlayersErr:
		return "KNSpNoPlayersErr"
	case KNSpNotAdvertisingErr:
		return "KNSpNotAdvertisingErr"
	case KNSpOTNotPresentErr:
		return "KNSpOTNotPresentErr"
	case KNSpOTVersionTooOldErr:
		return "KNSpOTVersionTooOldErr"
	case KNSpPipeFullErr:
		return "KNSpPipeFullErr"
	case KNSpProtocolNotAvailableErr:
		return "KNSpProtocolNotAvailableErr"
	case KNSpRemovePlayerFailedErr:
		return "KNSpRemovePlayerFailedErr"
	case KNSpSendFailedErr:
		return "KNSpSendFailedErr"
	case KNSpTimeoutErr:
		return "KNSpTimeoutErr"
	case KNSpTopologyNotSupportedErr:
		return "KNSpTopologyNotSupportedErr"
	default:
		return fmt.Sprintf("KNSp(%d)", e)
	}
}

type KNav int

const (
	KNavCustomControlMessageFailedErr KNav = -5697
	KNavInvalidCustomControlMessageErr KNav = -5698
	KNavInvalidSystemConfigErr KNav = -5696
	KNavMissingKindStringErr KNav = -5699
	KNavWrongDialogClassErr KNav = -5695
	KNavWrongDialogStateErr KNav = -5694
)

func (e KNav) String() string {
	switch e {
	case KNavCustomControlMessageFailedErr:
		return "KNavCustomControlMessageFailedErr"
	case KNavInvalidCustomControlMessageErr:
		return "KNavInvalidCustomControlMessageErr"
	case KNavInvalidSystemConfigErr:
		return "KNavInvalidSystemConfigErr"
	case KNavMissingKindStringErr:
		return "KNavMissingKindStringErr"
	case KNavWrongDialogClassErr:
		return "KNavWrongDialogClassErr"
	case KNavWrongDialogStateErr:
		return "KNavWrongDialogStateErr"
	default:
		return fmt.Sprintf("KNav(%d)", e)
	}
}

type KOTNoError int

const (
	KEACCESErr KOTNoError = -3212
	KEADDRINUSEErr KOTNoError = -3247
	KEADDRNOTAVAILErr KOTNoError = -3248
	KEAGAINErr KOTNoError = -3210
	KEALREADYErr KOTNoError = -3236
	KEBADFErr KOTNoError = -3208
	KEBADMSGErr KOTNoError = -3272
	KEBUSYErr KOTNoError = -3215
	KECANCELErr KOTNoError = -3273
	KECONNABORTEDErr KOTNoError = -3252
	KECONNREFUSEDErr KOTNoError = -3260
	KECONNRESETErr KOTNoError = -3253
	KEDEADLKErr KOTNoError = -3234
	KEDESTADDRREQErr KOTNoError = -3238
	KEEXISTErr KOTNoError = -3216
	KEFAULTErr KOTNoError = -3213
	KEHOSTDOWNErr KOTNoError = -3263
	KEHOSTUNREACHErr KOTNoError = -3264
	KEINPROGRESSErr KOTNoError = -3276
	KEINTRErr KOTNoError = -3203
	KEINVALErr KOTNoError = -3221
	KEIOErr KOTNoError = -3204
	KEISCONNErr KOTNoError = -3255
	KEMSGSIZEErr KOTNoError = -3239
	KENETDOWNErr KOTNoError = -3249
	KENETRESETErr KOTNoError = -3251
	KENETUNREACHErr KOTNoError = -3250
	KENOBUFSErr KOTNoError = -3254
	KENODATAErr KOTNoError = -3275
	KENODEVErr KOTNoError = -3218
	KENOENTErr KOTNoError = -3201
	KENOMEMErr KOTNoError = -3211
	KENOMSGErr KOTNoError = -3278
	KENOPROTOOPTErr KOTNoError = -3241
	KENORSRCErr KOTNoError = -3202
	KENOSRErr KOTNoError = -3271
	KENOSTRErr KOTNoError = -3274
	KENOTCONNErr KOTNoError = -3256
	KENOTSOCKErr KOTNoError = -3237
	KENOTTYErr KOTNoError = -3224
	KENXIOErr KOTNoError = -3205
	KEOPNOTSUPPErr KOTNoError = -3244
	KEPERMErr KOTNoError = -3200
	KEPIPEErr KOTNoError = -3231
	KEPROTOErr KOTNoError = -3269
	KEPROTONOSUPPORTErr KOTNoError = -3242
	KEPROTOTYPEErr KOTNoError = -3240
	KERANGEErr KOTNoError = -3233
	KESHUTDOWNErr KOTNoError = -3257
	KESOCKTNOSUPPORTErr KOTNoError = -3243
	KESRCHErr KOTNoError = -3277
	KETIMEDOUTErr KOTNoError = -3259
	KETIMEErr KOTNoError = -3270
	KETOOMANYREFSErr KOTNoError = -3258
	KEWOULDBLOCKErr KOTNoError = -3234
	KOTAccessErr KOTNoError = -3152
	KOTAddressBusyErr KOTNoError = -3172
	KOTBadAddressErr KOTNoError = -3150
	KOTBadConfigurationErr KOTNoError = -3282
	KOTBadDataErr KOTNoError = -3159
	KOTBadFlagErr KOTNoError = -3165
	KOTBadNameErr KOTNoError = -3170
	KOTBadOptionErr KOTNoError = -3151
	KOTBadQLenErr KOTNoError = -3171
	KOTBadReferenceErr KOTNoError = -3153
	KOTBadSequenceErr KOTNoError = -3156
	KOTBadSyncErr KOTNoError = -3179
	KOTBufferOverflowErr KOTNoError = -3160
	KOTCanceledErr KOTNoError = -3180
	KOTClientNotInittedErr KOTNoError = -3279
	KOTConfigurationChangedErr KOTNoError = -3283
	KOTDuplicateFoundErr KOTNoError = -3216
	KOTFlowErr KOTNoError = -3161
	KOTIndOutErr KOTNoError = -3173
	KOTLookErr KOTNoError = -3158
	KOTNoAddressErr KOTNoError = -3154
	KOTNoDataErr KOTNoError = -3162
	KOTNoDisconnectErr KOTNoError = -3163
	KOTNoErrorValue KOTNoError = 0
	KOTNoReleaseErr KOTNoError = -3166
	KOTNoStructureTypeErr KOTNoError = -3169
	KOTNoUDErrErr KOTNoError = -3164
	KOTNotFoundErr KOTNoError = -3201
	KOTNotSupportedErr KOTNoError = -3167
	KOTOutOfMemoryErr KOTNoError = -3211
	KOTOutStateErr KOTNoError = -3155
	KOTPortHasDiedErr KOTNoError = -3280
	KOTPortLostConnection KOTNoError = -3285
	KOTPortWasEjectedErr KOTNoError = -3281
	KOTProtocolErr KOTNoError = -3178
	KOTProviderMismatchErr KOTNoError = -3174
	KOTQFullErr KOTNoError = -3177
	KOTResAddressErr KOTNoError = -3176
	KOTResQLenErr KOTNoError = -3175
	KOTStateChangeErr KOTNoError = -3168
	KOTSysErrorErr KOTNoError = -3157
	KOTUserRequestedErr KOTNoError = -3284
)

func (e KOTNoError) String() string {
	switch e {
	case KEACCESErr:
		return "KEACCESErr"
	case KEADDRINUSEErr:
		return "KEADDRINUSEErr"
	case KEADDRNOTAVAILErr:
		return "KEADDRNOTAVAILErr"
	case KEAGAINErr:
		return "KEAGAINErr"
	case KEALREADYErr:
		return "KEALREADYErr"
	case KEBADFErr:
		return "KEBADFErr"
	case KEBADMSGErr:
		return "KEBADMSGErr"
	case KEBUSYErr:
		return "KEBUSYErr"
	case KECANCELErr:
		return "KECANCELErr"
	case KECONNABORTEDErr:
		return "KECONNABORTEDErr"
	case KECONNREFUSEDErr:
		return "KECONNREFUSEDErr"
	case KECONNRESETErr:
		return "KECONNRESETErr"
	case KEDEADLKErr:
		return "KEDEADLKErr"
	case KEDESTADDRREQErr:
		return "KEDESTADDRREQErr"
	case KEEXISTErr:
		return "KEEXISTErr"
	case KEFAULTErr:
		return "KEFAULTErr"
	case KEHOSTDOWNErr:
		return "KEHOSTDOWNErr"
	case KEHOSTUNREACHErr:
		return "KEHOSTUNREACHErr"
	case KEINPROGRESSErr:
		return "KEINPROGRESSErr"
	case KEINTRErr:
		return "KEINTRErr"
	case KEINVALErr:
		return "KEINVALErr"
	case KEIOErr:
		return "KEIOErr"
	case KEISCONNErr:
		return "KEISCONNErr"
	case KEMSGSIZEErr:
		return "KEMSGSIZEErr"
	case KENETDOWNErr:
		return "KENETDOWNErr"
	case KENETRESETErr:
		return "KENETRESETErr"
	case KENETUNREACHErr:
		return "KENETUNREACHErr"
	case KENOBUFSErr:
		return "KENOBUFSErr"
	case KENODATAErr:
		return "KENODATAErr"
	case KENODEVErr:
		return "KENODEVErr"
	case KENOENTErr:
		return "KENOENTErr"
	case KENOMEMErr:
		return "KENOMEMErr"
	case KENOMSGErr:
		return "KENOMSGErr"
	case KENOPROTOOPTErr:
		return "KENOPROTOOPTErr"
	case KENORSRCErr:
		return "KENORSRCErr"
	case KENOSRErr:
		return "KENOSRErr"
	case KENOSTRErr:
		return "KENOSTRErr"
	case KENOTCONNErr:
		return "KENOTCONNErr"
	case KENOTSOCKErr:
		return "KENOTSOCKErr"
	case KENOTTYErr:
		return "KENOTTYErr"
	case KENXIOErr:
		return "KENXIOErr"
	case KEOPNOTSUPPErr:
		return "KEOPNOTSUPPErr"
	case KEPERMErr:
		return "KEPERMErr"
	case KEPIPEErr:
		return "KEPIPEErr"
	case KEPROTOErr:
		return "KEPROTOErr"
	case KEPROTONOSUPPORTErr:
		return "KEPROTONOSUPPORTErr"
	case KEPROTOTYPEErr:
		return "KEPROTOTYPEErr"
	case KERANGEErr:
		return "KERANGEErr"
	case KESHUTDOWNErr:
		return "KESHUTDOWNErr"
	case KESOCKTNOSUPPORTErr:
		return "KESOCKTNOSUPPORTErr"
	case KESRCHErr:
		return "KESRCHErr"
	case KETIMEDOUTErr:
		return "KETIMEDOUTErr"
	case KETIMEErr:
		return "KETIMEErr"
	case KETOOMANYREFSErr:
		return "KETOOMANYREFSErr"
	case KOTAccessErr:
		return "KOTAccessErr"
	case KOTAddressBusyErr:
		return "KOTAddressBusyErr"
	case KOTBadAddressErr:
		return "KOTBadAddressErr"
	case KOTBadConfigurationErr:
		return "KOTBadConfigurationErr"
	case KOTBadDataErr:
		return "KOTBadDataErr"
	case KOTBadFlagErr:
		return "KOTBadFlagErr"
	case KOTBadNameErr:
		return "KOTBadNameErr"
	case KOTBadOptionErr:
		return "KOTBadOptionErr"
	case KOTBadQLenErr:
		return "KOTBadQLenErr"
	case KOTBadReferenceErr:
		return "KOTBadReferenceErr"
	case KOTBadSequenceErr:
		return "KOTBadSequenceErr"
	case KOTBadSyncErr:
		return "KOTBadSyncErr"
	case KOTBufferOverflowErr:
		return "KOTBufferOverflowErr"
	case KOTCanceledErr:
		return "KOTCanceledErr"
	case KOTClientNotInittedErr:
		return "KOTClientNotInittedErr"
	case KOTConfigurationChangedErr:
		return "KOTConfigurationChangedErr"
	case KOTFlowErr:
		return "KOTFlowErr"
	case KOTIndOutErr:
		return "KOTIndOutErr"
	case KOTLookErr:
		return "KOTLookErr"
	case KOTNoAddressErr:
		return "KOTNoAddressErr"
	case KOTNoDataErr:
		return "KOTNoDataErr"
	case KOTNoDisconnectErr:
		return "KOTNoDisconnectErr"
	case KOTNoErrorValue:
		return "KOTNoErrorValue"
	case KOTNoReleaseErr:
		return "KOTNoReleaseErr"
	case KOTNoStructureTypeErr:
		return "KOTNoStructureTypeErr"
	case KOTNoUDErrErr:
		return "KOTNoUDErrErr"
	case KOTNotSupportedErr:
		return "KOTNotSupportedErr"
	case KOTOutStateErr:
		return "KOTOutStateErr"
	case KOTPortHasDiedErr:
		return "KOTPortHasDiedErr"
	case KOTPortLostConnection:
		return "KOTPortLostConnection"
	case KOTPortWasEjectedErr:
		return "KOTPortWasEjectedErr"
	case KOTProtocolErr:
		return "KOTProtocolErr"
	case KOTProviderMismatchErr:
		return "KOTProviderMismatchErr"
	case KOTQFullErr:
		return "KOTQFullErr"
	case KOTResAddressErr:
		return "KOTResAddressErr"
	case KOTResQLenErr:
		return "KOTResQLenErr"
	case KOTStateChangeErr:
		return "KOTStateChangeErr"
	case KOTSysErrorErr:
		return "KOTSysErrorErr"
	case KOTUserRequestedErr:
		return "KOTUserRequestedErr"
	default:
		return fmt.Sprintf("KOTNoError(%d)", e)
	}
}

type KPOSIXError int

const (
	KPOSIXErrorBase KPOSIXError = 100000
	KPOSIXErrorE2BIG KPOSIXError = 100007
	KPOSIXErrorEACCES KPOSIXError = 100013
	KPOSIXErrorEADDRINUSE KPOSIXError = 100048
	KPOSIXErrorEADDRNOTAVAIL KPOSIXError = 100049
	KPOSIXErrorEAFNOSUPPORT KPOSIXError = 100047
	KPOSIXErrorEAGAIN KPOSIXError = 100035
	KPOSIXErrorEALREADY KPOSIXError = 100037
	KPOSIXErrorEAUTH KPOSIXError = 100080
	KPOSIXErrorEBADARCH KPOSIXError = 100086
	KPOSIXErrorEBADEXEC KPOSIXError = 100085
	KPOSIXErrorEBADF KPOSIXError = 100009
	KPOSIXErrorEBADMACHO KPOSIXError = 100088
	KPOSIXErrorEBADMSG KPOSIXError = 100094
	KPOSIXErrorEBADRPC KPOSIXError = 100072
	KPOSIXErrorEBUSY KPOSIXError = 100016
	KPOSIXErrorECANCELED KPOSIXError = 100089
	KPOSIXErrorECHILD KPOSIXError = 100010
	KPOSIXErrorECONNABORTED KPOSIXError = 100053
	KPOSIXErrorECONNREFUSED KPOSIXError = 100061
	KPOSIXErrorECONNRESET KPOSIXError = 100054
	KPOSIXErrorEDEADLK KPOSIXError = 100011
	KPOSIXErrorEDESTADDRREQ KPOSIXError = 100039
	KPOSIXErrorEDEVERR KPOSIXError = 100083
	KPOSIXErrorEDOM KPOSIXError = 100033
	KPOSIXErrorEDQUOT KPOSIXError = 100069
	KPOSIXErrorEEXIST KPOSIXError = 100017
	KPOSIXErrorEFAULT KPOSIXError = 100014
	KPOSIXErrorEFBIG KPOSIXError = 100027
	KPOSIXErrorEFTYPE KPOSIXError = 100079
	KPOSIXErrorEHOSTDOWN KPOSIXError = 100064
	KPOSIXErrorEHOSTUNREACH KPOSIXError = 100065
	KPOSIXErrorEIDRM KPOSIXError = 100090
	KPOSIXErrorEILSEQ KPOSIXError = 100092
	KPOSIXErrorEINPROGRESS KPOSIXError = 100036
	KPOSIXErrorEINTR KPOSIXError = 100004
	KPOSIXErrorEINVAL KPOSIXError = 100022
	KPOSIXErrorEIO KPOSIXError = 100005
	KPOSIXErrorEISCONN KPOSIXError = 100056
	KPOSIXErrorEISDIR KPOSIXError = 100021
	KPOSIXErrorELOOP KPOSIXError = 100062
	KPOSIXErrorEMFILE KPOSIXError = 100024
	KPOSIXErrorEMLINK KPOSIXError = 100031
	KPOSIXErrorEMSGSIZE KPOSIXError = 100040
	KPOSIXErrorEMULTIHOP KPOSIXError = 100095
	KPOSIXErrorENAMETOOLONG KPOSIXError = 100063
	KPOSIXErrorENEEDAUTH KPOSIXError = 100081
	KPOSIXErrorENETDOWN KPOSIXError = 100050
	KPOSIXErrorENETRESET KPOSIXError = 100052
	KPOSIXErrorENETUNREACH KPOSIXError = 100051
	KPOSIXErrorENFILE KPOSIXError = 100023
	KPOSIXErrorENOATTR KPOSIXError = 100093
	KPOSIXErrorENOBUFS KPOSIXError = 100055
	KPOSIXErrorENODATA KPOSIXError = 100096
	KPOSIXErrorENODEV KPOSIXError = 100019
	KPOSIXErrorENOENT KPOSIXError = 100002
	KPOSIXErrorENOEXEC KPOSIXError = 100008
	KPOSIXErrorENOLCK KPOSIXError = 100077
	KPOSIXErrorENOLINK KPOSIXError = 100097
	KPOSIXErrorENOMEM KPOSIXError = 100012
	KPOSIXErrorENOMSG KPOSIXError = 100091
	KPOSIXErrorENOPROTOOPT KPOSIXError = 100042
	KPOSIXErrorENOSPC KPOSIXError = 100028
	KPOSIXErrorENOSR KPOSIXError = 100098
	KPOSIXErrorENOSTR KPOSIXError = 100099
	KPOSIXErrorENOSYS KPOSIXError = 100078
	KPOSIXErrorENOTBLK KPOSIXError = 100015
	KPOSIXErrorENOTCONN KPOSIXError = 100057
	KPOSIXErrorENOTDIR KPOSIXError = 100020
	KPOSIXErrorENOTEMPTY KPOSIXError = 100066
	KPOSIXErrorENOTSOCK KPOSIXError = 100038
	KPOSIXErrorENOTSUP KPOSIXError = 100045
	KPOSIXErrorENOTTY KPOSIXError = 100025
	KPOSIXErrorENXIO KPOSIXError = 100006
	KPOSIXErrorEOPNOTSUPP KPOSIXError = 100102
	KPOSIXErrorEOVERFLOW KPOSIXError = 100084
	KPOSIXErrorEPERM KPOSIXError = 100001
	KPOSIXErrorEPFNOSUPPORT KPOSIXError = 100046
	KPOSIXErrorEPIPE KPOSIXError = 100032
	KPOSIXErrorEPROCLIM KPOSIXError = 100067
	KPOSIXErrorEPROCUNAVAIL KPOSIXError = 100076
	KPOSIXErrorEPROGMISMATCH KPOSIXError = 100075
	KPOSIXErrorEPROGUNAVAIL KPOSIXError = 100074
	KPOSIXErrorEPROTO KPOSIXError = 100100
	KPOSIXErrorEPROTONOSUPPORT KPOSIXError = 100043
	KPOSIXErrorEPROTOTYPE KPOSIXError = 100041
	KPOSIXErrorEPWROFF KPOSIXError = 100082
	KPOSIXErrorERANGE KPOSIXError = 100034
	KPOSIXErrorEREMOTE KPOSIXError = 100071
	KPOSIXErrorEROFS KPOSIXError = 100030
	KPOSIXErrorERPCMISMATCH KPOSIXError = 100073
	KPOSIXErrorESHLIBVERS KPOSIXError = 100087
	KPOSIXErrorESHUTDOWN KPOSIXError = 100058
	KPOSIXErrorESOCKTNOSUPPORT KPOSIXError = 100044
	KPOSIXErrorESPIPE KPOSIXError = 100029
	KPOSIXErrorESRCH KPOSIXError = 100003
	KPOSIXErrorESTALE KPOSIXError = 100070
	KPOSIXErrorETIME KPOSIXError = 100101
	KPOSIXErrorETIMEDOUT KPOSIXError = 100060
	KPOSIXErrorETOOMANYREFS KPOSIXError = 100059
	KPOSIXErrorETXTBSY KPOSIXError = 100026
	KPOSIXErrorEUSERS KPOSIXError = 100068
	KPOSIXErrorEXDEV KPOSIXError = 100018
)

func (e KPOSIXError) String() string {
	switch e {
	case KPOSIXErrorBase:
		return "KPOSIXErrorBase"
	case KPOSIXErrorE2BIG:
		return "KPOSIXErrorE2BIG"
	case KPOSIXErrorEACCES:
		return "KPOSIXErrorEACCES"
	case KPOSIXErrorEADDRINUSE:
		return "KPOSIXErrorEADDRINUSE"
	case KPOSIXErrorEADDRNOTAVAIL:
		return "KPOSIXErrorEADDRNOTAVAIL"
	case KPOSIXErrorEAFNOSUPPORT:
		return "KPOSIXErrorEAFNOSUPPORT"
	case KPOSIXErrorEAGAIN:
		return "KPOSIXErrorEAGAIN"
	case KPOSIXErrorEALREADY:
		return "KPOSIXErrorEALREADY"
	case KPOSIXErrorEAUTH:
		return "KPOSIXErrorEAUTH"
	case KPOSIXErrorEBADARCH:
		return "KPOSIXErrorEBADARCH"
	case KPOSIXErrorEBADEXEC:
		return "KPOSIXErrorEBADEXEC"
	case KPOSIXErrorEBADF:
		return "KPOSIXErrorEBADF"
	case KPOSIXErrorEBADMACHO:
		return "KPOSIXErrorEBADMACHO"
	case KPOSIXErrorEBADMSG:
		return "KPOSIXErrorEBADMSG"
	case KPOSIXErrorEBADRPC:
		return "KPOSIXErrorEBADRPC"
	case KPOSIXErrorEBUSY:
		return "KPOSIXErrorEBUSY"
	case KPOSIXErrorECANCELED:
		return "KPOSIXErrorECANCELED"
	case KPOSIXErrorECHILD:
		return "KPOSIXErrorECHILD"
	case KPOSIXErrorECONNABORTED:
		return "KPOSIXErrorECONNABORTED"
	case KPOSIXErrorECONNREFUSED:
		return "KPOSIXErrorECONNREFUSED"
	case KPOSIXErrorECONNRESET:
		return "KPOSIXErrorECONNRESET"
	case KPOSIXErrorEDEADLK:
		return "KPOSIXErrorEDEADLK"
	case KPOSIXErrorEDESTADDRREQ:
		return "KPOSIXErrorEDESTADDRREQ"
	case KPOSIXErrorEDEVERR:
		return "KPOSIXErrorEDEVERR"
	case KPOSIXErrorEDOM:
		return "KPOSIXErrorEDOM"
	case KPOSIXErrorEDQUOT:
		return "KPOSIXErrorEDQUOT"
	case KPOSIXErrorEEXIST:
		return "KPOSIXErrorEEXIST"
	case KPOSIXErrorEFAULT:
		return "KPOSIXErrorEFAULT"
	case KPOSIXErrorEFBIG:
		return "KPOSIXErrorEFBIG"
	case KPOSIXErrorEFTYPE:
		return "KPOSIXErrorEFTYPE"
	case KPOSIXErrorEHOSTDOWN:
		return "KPOSIXErrorEHOSTDOWN"
	case KPOSIXErrorEHOSTUNREACH:
		return "KPOSIXErrorEHOSTUNREACH"
	case KPOSIXErrorEIDRM:
		return "KPOSIXErrorEIDRM"
	case KPOSIXErrorEILSEQ:
		return "KPOSIXErrorEILSEQ"
	case KPOSIXErrorEINPROGRESS:
		return "KPOSIXErrorEINPROGRESS"
	case KPOSIXErrorEINTR:
		return "KPOSIXErrorEINTR"
	case KPOSIXErrorEINVAL:
		return "KPOSIXErrorEINVAL"
	case KPOSIXErrorEIO:
		return "KPOSIXErrorEIO"
	case KPOSIXErrorEISCONN:
		return "KPOSIXErrorEISCONN"
	case KPOSIXErrorEISDIR:
		return "KPOSIXErrorEISDIR"
	case KPOSIXErrorELOOP:
		return "KPOSIXErrorELOOP"
	case KPOSIXErrorEMFILE:
		return "KPOSIXErrorEMFILE"
	case KPOSIXErrorEMLINK:
		return "KPOSIXErrorEMLINK"
	case KPOSIXErrorEMSGSIZE:
		return "KPOSIXErrorEMSGSIZE"
	case KPOSIXErrorEMULTIHOP:
		return "KPOSIXErrorEMULTIHOP"
	case KPOSIXErrorENAMETOOLONG:
		return "KPOSIXErrorENAMETOOLONG"
	case KPOSIXErrorENEEDAUTH:
		return "KPOSIXErrorENEEDAUTH"
	case KPOSIXErrorENETDOWN:
		return "KPOSIXErrorENETDOWN"
	case KPOSIXErrorENETRESET:
		return "KPOSIXErrorENETRESET"
	case KPOSIXErrorENETUNREACH:
		return "KPOSIXErrorENETUNREACH"
	case KPOSIXErrorENFILE:
		return "KPOSIXErrorENFILE"
	case KPOSIXErrorENOATTR:
		return "KPOSIXErrorENOATTR"
	case KPOSIXErrorENOBUFS:
		return "KPOSIXErrorENOBUFS"
	case KPOSIXErrorENODATA:
		return "KPOSIXErrorENODATA"
	case KPOSIXErrorENODEV:
		return "KPOSIXErrorENODEV"
	case KPOSIXErrorENOENT:
		return "KPOSIXErrorENOENT"
	case KPOSIXErrorENOEXEC:
		return "KPOSIXErrorENOEXEC"
	case KPOSIXErrorENOLCK:
		return "KPOSIXErrorENOLCK"
	case KPOSIXErrorENOLINK:
		return "KPOSIXErrorENOLINK"
	case KPOSIXErrorENOMEM:
		return "KPOSIXErrorENOMEM"
	case KPOSIXErrorENOMSG:
		return "KPOSIXErrorENOMSG"
	case KPOSIXErrorENOPROTOOPT:
		return "KPOSIXErrorENOPROTOOPT"
	case KPOSIXErrorENOSPC:
		return "KPOSIXErrorENOSPC"
	case KPOSIXErrorENOSR:
		return "KPOSIXErrorENOSR"
	case KPOSIXErrorENOSTR:
		return "KPOSIXErrorENOSTR"
	case KPOSIXErrorENOSYS:
		return "KPOSIXErrorENOSYS"
	case KPOSIXErrorENOTBLK:
		return "KPOSIXErrorENOTBLK"
	case KPOSIXErrorENOTCONN:
		return "KPOSIXErrorENOTCONN"
	case KPOSIXErrorENOTDIR:
		return "KPOSIXErrorENOTDIR"
	case KPOSIXErrorENOTEMPTY:
		return "KPOSIXErrorENOTEMPTY"
	case KPOSIXErrorENOTSOCK:
		return "KPOSIXErrorENOTSOCK"
	case KPOSIXErrorENOTSUP:
		return "KPOSIXErrorENOTSUP"
	case KPOSIXErrorENOTTY:
		return "KPOSIXErrorENOTTY"
	case KPOSIXErrorENXIO:
		return "KPOSIXErrorENXIO"
	case KPOSIXErrorEOPNOTSUPP:
		return "KPOSIXErrorEOPNOTSUPP"
	case KPOSIXErrorEOVERFLOW:
		return "KPOSIXErrorEOVERFLOW"
	case KPOSIXErrorEPERM:
		return "KPOSIXErrorEPERM"
	case KPOSIXErrorEPFNOSUPPORT:
		return "KPOSIXErrorEPFNOSUPPORT"
	case KPOSIXErrorEPIPE:
		return "KPOSIXErrorEPIPE"
	case KPOSIXErrorEPROCLIM:
		return "KPOSIXErrorEPROCLIM"
	case KPOSIXErrorEPROCUNAVAIL:
		return "KPOSIXErrorEPROCUNAVAIL"
	case KPOSIXErrorEPROGMISMATCH:
		return "KPOSIXErrorEPROGMISMATCH"
	case KPOSIXErrorEPROGUNAVAIL:
		return "KPOSIXErrorEPROGUNAVAIL"
	case KPOSIXErrorEPROTO:
		return "KPOSIXErrorEPROTO"
	case KPOSIXErrorEPROTONOSUPPORT:
		return "KPOSIXErrorEPROTONOSUPPORT"
	case KPOSIXErrorEPROTOTYPE:
		return "KPOSIXErrorEPROTOTYPE"
	case KPOSIXErrorEPWROFF:
		return "KPOSIXErrorEPWROFF"
	case KPOSIXErrorERANGE:
		return "KPOSIXErrorERANGE"
	case KPOSIXErrorEREMOTE:
		return "KPOSIXErrorEREMOTE"
	case KPOSIXErrorEROFS:
		return "KPOSIXErrorEROFS"
	case KPOSIXErrorERPCMISMATCH:
		return "KPOSIXErrorERPCMISMATCH"
	case KPOSIXErrorESHLIBVERS:
		return "KPOSIXErrorESHLIBVERS"
	case KPOSIXErrorESHUTDOWN:
		return "KPOSIXErrorESHUTDOWN"
	case KPOSIXErrorESOCKTNOSUPPORT:
		return "KPOSIXErrorESOCKTNOSUPPORT"
	case KPOSIXErrorESPIPE:
		return "KPOSIXErrorESPIPE"
	case KPOSIXErrorESRCH:
		return "KPOSIXErrorESRCH"
	case KPOSIXErrorESTALE:
		return "KPOSIXErrorESTALE"
	case KPOSIXErrorETIME:
		return "KPOSIXErrorETIME"
	case KPOSIXErrorETIMEDOUT:
		return "KPOSIXErrorETIMEDOUT"
	case KPOSIXErrorETOOMANYREFS:
		return "KPOSIXErrorETOOMANYREFS"
	case KPOSIXErrorETXTBSY:
		return "KPOSIXErrorETXTBSY"
	case KPOSIXErrorEUSERS:
		return "KPOSIXErrorEUSERS"
	case KPOSIXErrorEXDEV:
		return "KPOSIXErrorEXDEV"
	default:
		return fmt.Sprintf("KPOSIXError(%d)", e)
	}
}

type KPowerHandlerExistsForDeviceErr int

const (
	KBridgeSoftwareRunningCantSleep KPowerHandlerExistsForDeviceErr = -13038
	KCantReportProcessorTemperatureErr KPowerHandlerExistsForDeviceErr = -13013
	KNoSuchPowerSource KPowerHandlerExistsForDeviceErr = -13020
	KPowerHandlerExistsForDeviceErrValue KPowerHandlerExistsForDeviceErr = -13006
	KPowerHandlerNotFoundForDeviceErr KPowerHandlerExistsForDeviceErr = -13007
	KPowerHandlerNotFoundForProcErr KPowerHandlerExistsForDeviceErr = -13008
	KPowerMgtMessageNotHandled KPowerHandlerExistsForDeviceErr = -13009
	KPowerMgtRequestDenied KPowerHandlerExistsForDeviceErr = -13010
	KProcessorTempRoutineRequiresMPLib2 KPowerHandlerExistsForDeviceErr = -13014
)

func (e KPowerHandlerExistsForDeviceErr) String() string {
	switch e {
	case KBridgeSoftwareRunningCantSleep:
		return "KBridgeSoftwareRunningCantSleep"
	case KCantReportProcessorTemperatureErr:
		return "KCantReportProcessorTemperatureErr"
	case KNoSuchPowerSource:
		return "KNoSuchPowerSource"
	case KPowerHandlerExistsForDeviceErrValue:
		return "KPowerHandlerExistsForDeviceErrValue"
	case KPowerHandlerNotFoundForDeviceErr:
		return "KPowerHandlerNotFoundForDeviceErr"
	case KPowerHandlerNotFoundForProcErr:
		return "KPowerHandlerNotFoundForProcErr"
	case KPowerMgtMessageNotHandled:
		return "KPowerMgtMessageNotHandled"
	case KPowerMgtRequestDenied:
		return "KPowerMgtRequestDenied"
	case KProcessorTempRoutineRequiresMPLib2:
		return "KProcessorTempRoutineRequiresMPLib2"
	default:
		return fmt.Sprintf("KPowerHandlerExistsForDeviceErr(%d)", e)
	}
}

type KQD int

const (
	KQDCorruptPICTDataErr KQD = -3954
	KQDCursorAlreadyRegistered KQD = -3952
	KQDCursorNotRegistered KQD = -3953
	KQDNoColorHWCursorSupport KQD = -3951
	KQDNoPalette KQD = -3950
)

func (e KQD) String() string {
	switch e {
	case KQDCorruptPICTDataErr:
		return "KQDCorruptPICTDataErr"
	case KQDCursorAlreadyRegistered:
		return "KQDCursorAlreadyRegistered"
	case KQDCursorNotRegistered:
		return "KQDCursorNotRegistered"
	case KQDNoColorHWCursorSupport:
		return "KQDNoColorHWCursorSupport"
	case KQDNoPalette:
		return "KQDNoPalette"
	default:
		return fmt.Sprintf("KQD(%d)", e)
	}
}

const KQTSSUnknownErr int = -6150

type KRA int

const (
	KRAATalkInactive KRA = -7134
	KRACallBackFailed KRA = -7138
	KRAConfigurationDBInitErr KRA = -7127
	KRAConnectionCanceled KRA = -7109
	KRADuplicateIPAddr KRA = -7137
	KRAExtAuthenticationFailed KRA = -7135
	KRAIncompatiblePrefs KRA = -7107
	KRAInitOpenTransportFailed KRA = -7122
	KRAInstallationDamaged KRA = -7113
	KRAInternalError KRA = -7112
	KRAInvalidParameter KRA = -7100
	KRAInvalidPassword KRA = -7111
	KRAInvalidPort KRA = -7101
	KRAInvalidPortState KRA = -7116
	KRAInvalidSerialProtocol KRA = -7117
	KRAMissingResources KRA = -7106
	KRANCPRejectedbyPeer KRA = -7136
	KRANotConnected KRA = -7108
	KRANotEnabled KRA = -7139
	KRANotPrimaryInterface KRA = -7126
	KRANotSupported KRA = -7105
	KRAOutOfMemory KRA = -7104
	KRAPPPAuthenticationFailed KRA = -7129
	KRAPPPNegotiationFailed KRA = -7130
	KRAPPPPeerDisconnected KRA = -7132
	KRAPPPProtocolRejected KRA = -7128
	KRAPPPUserDisconnected KRA = -7131
	KRAPeerNotResponding KRA = -7133
	KRAPortBusy KRA = -7114
	KRAPortSetupFailed KRA = -7103
	KRARemoteAccessNotReady KRA = -7123
	KRAStartupFailed KRA = -7102
	KRATCPIPInactive KRA = -7124
	KRATCPIPNotConfigured KRA = -7125
	KRAUnknownPortState KRA = -7115
	KRAUnknownUser KRA = -7110
	KRAUserInteractionRequired KRA = -7121
	KRAUserLoginDisabled KRA = -7118
	KRAUserPwdChangeRequired KRA = -7119
	KRAUserPwdEntryRequired KRA = -7120
)

func (e KRA) String() string {
	switch e {
	case KRAATalkInactive:
		return "KRAATalkInactive"
	case KRACallBackFailed:
		return "KRACallBackFailed"
	case KRAConfigurationDBInitErr:
		return "KRAConfigurationDBInitErr"
	case KRAConnectionCanceled:
		return "KRAConnectionCanceled"
	case KRADuplicateIPAddr:
		return "KRADuplicateIPAddr"
	case KRAExtAuthenticationFailed:
		return "KRAExtAuthenticationFailed"
	case KRAIncompatiblePrefs:
		return "KRAIncompatiblePrefs"
	case KRAInitOpenTransportFailed:
		return "KRAInitOpenTransportFailed"
	case KRAInstallationDamaged:
		return "KRAInstallationDamaged"
	case KRAInternalError:
		return "KRAInternalError"
	case KRAInvalidParameter:
		return "KRAInvalidParameter"
	case KRAInvalidPassword:
		return "KRAInvalidPassword"
	case KRAInvalidPort:
		return "KRAInvalidPort"
	case KRAInvalidPortState:
		return "KRAInvalidPortState"
	case KRAInvalidSerialProtocol:
		return "KRAInvalidSerialProtocol"
	case KRAMissingResources:
		return "KRAMissingResources"
	case KRANCPRejectedbyPeer:
		return "KRANCPRejectedbyPeer"
	case KRANotConnected:
		return "KRANotConnected"
	case KRANotEnabled:
		return "KRANotEnabled"
	case KRANotPrimaryInterface:
		return "KRANotPrimaryInterface"
	case KRANotSupported:
		return "KRANotSupported"
	case KRAOutOfMemory:
		return "KRAOutOfMemory"
	case KRAPPPAuthenticationFailed:
		return "KRAPPPAuthenticationFailed"
	case KRAPPPNegotiationFailed:
		return "KRAPPPNegotiationFailed"
	case KRAPPPPeerDisconnected:
		return "KRAPPPPeerDisconnected"
	case KRAPPPProtocolRejected:
		return "KRAPPPProtocolRejected"
	case KRAPPPUserDisconnected:
		return "KRAPPPUserDisconnected"
	case KRAPeerNotResponding:
		return "KRAPeerNotResponding"
	case KRAPortBusy:
		return "KRAPortBusy"
	case KRAPortSetupFailed:
		return "KRAPortSetupFailed"
	case KRARemoteAccessNotReady:
		return "KRARemoteAccessNotReady"
	case KRAStartupFailed:
		return "KRAStartupFailed"
	case KRATCPIPInactive:
		return "KRATCPIPInactive"
	case KRATCPIPNotConfigured:
		return "KRATCPIPNotConfigured"
	case KRAUnknownPortState:
		return "KRAUnknownPortState"
	case KRAUnknownUser:
		return "KRAUnknownUser"
	case KRAUserInteractionRequired:
		return "KRAUserInteractionRequired"
	case KRAUserLoginDisabled:
		return "KRAUserLoginDisabled"
	case KRAUserPwdChangeRequired:
		return "KRAUserPwdChangeRequired"
	case KRAUserPwdEntryRequired:
		return "KRAUserPwdEntryRequired"
	default:
		return fmt.Sprintf("KRA(%d)", e)
	}
}

type KSKSearchOption uint

const (
	// KSKSearchOptionDefault: # Discussion
	KSKSearchOptionDefault KSKSearchOption = 0
	// KSKSearchOptionFindSimilar: This option alters query behavior so that Search Kit returns references to documents that are similar to an example text string.
	KSKSearchOptionFindSimilar KSKSearchOption = 4
	// KSKSearchOptionNoRelevanceScores: This option saves time during a search by suppressing the computation of relevance scores.
	KSKSearchOptionNoRelevanceScores KSKSearchOption = 1
	// KSKSearchOptionSpaceMeansOR: This option alters query behavior so that spaces are interpreted as Boolean [OR] operators.
	KSKSearchOptionSpaceMeansOR KSKSearchOption = 2
)

func (e KSKSearchOption) String() string {
	switch e {
	case KSKSearchOptionDefault:
		return "KSKSearchOptionDefault"
	case KSKSearchOptionFindSimilar:
		return "KSKSearchOptionFindSimilar"
	case KSKSearchOptionNoRelevanceScores:
		return "KSKSearchOptionNoRelevanceScores"
	case KSKSearchOptionSpaceMeansOR:
		return "KSKSearchOptionSpaceMeansOR"
	default:
		return fmt.Sprintf("KSKSearchOption(%d)", e)
	}
}

type KSSp int

const (
	KSSpCantInstallErr KSSp = -30342
	KSSpInternalErr KSSp = -30340
	KSSpParallelUpVectorErr KSSp = -30343
	KSSpScaleToZeroErr KSSp = -30344
	KSSpVersionErr KSSp = -30341
)

func (e KSSp) String() string {
	switch e {
	case KSSpCantInstallErr:
		return "KSSpCantInstallErr"
	case KSSpInternalErr:
		return "KSSpInternalErr"
	case KSSpParallelUpVectorErr:
		return "KSSpParallelUpVectorErr"
	case KSSpScaleToZeroErr:
		return "KSSpScaleToZeroErr"
	case KSSpVersionErr:
		return "KSSpVersionErr"
	default:
		return fmt.Sprintf("KSSp(%d)", e)
	}
}

type KSharingPrivs uint

const (
	KSharingPrivsNotApplicableIcon KSharingPrivs = 's'<<24 | 'h'<<16 | 'n'<<8 | 'a' // 'shna'
	KSharingPrivsReadOnlyIcon KSharingPrivs = 's'<<24 | 'h'<<16 | 'r'<<8 | 'o' // 'shro'
	KSharingPrivsReadWriteIcon KSharingPrivs = 's'<<24 | 'h'<<16 | 'r'<<8 | 'w' // 'shrw'
	KSharingPrivsUnknownIcon KSharingPrivs = 's'<<24 | 'h'<<16 | 'u'<<8 | 'k' // 'shuk'
	KSharingPrivsWritableIcon KSharingPrivs = 'w'<<24 | 'r'<<16 | 'i'<<8 | 't' // 'writ'
)

func (e KSharingPrivs) String() string {
	switch e {
	case KSharingPrivsNotApplicableIcon:
		return "KSharingPrivsNotApplicableIcon"
	case KSharingPrivsReadOnlyIcon:
		return "KSharingPrivsReadOnlyIcon"
	case KSharingPrivsReadWriteIcon:
		return "KSharingPrivsReadWriteIcon"
	case KSharingPrivsUnknownIcon:
		return "KSharingPrivsUnknownIcon"
	case KSharingPrivsWritableIcon:
		return "KSharingPrivsWritableIcon"
	default:
		return fmt.Sprintf("KSharingPrivs(%d)", e)
	}
}

type KStartupFolderIconResource int

const (
	KControlPanelFolderIconResource KStartupFolderIconResource = -3976
	KDropFolderIconResource KStartupFolderIconResource = -3979
	KExtensionsFolderIconResource KStartupFolderIconResource = -3973
	KFontsFolderIconResource KStartupFolderIconResource = -3968
	KFullTrashIconResource KStartupFolderIconResource = -3984
	KMountedFolderIconResource KStartupFolderIconResource = -3977
	KOwnedFolderIconResource KStartupFolderIconResource = -3980
	KPreferencesFolderIconResource KStartupFolderIconResource = -3974
	KPrintMonitorFolderIconResource KStartupFolderIconResource = -3975
	KSharedFolderIconResource KStartupFolderIconResource = -3978
	KStartupFolderIconResourceValue KStartupFolderIconResource = -3981
)

func (e KStartupFolderIconResource) String() string {
	switch e {
	case KControlPanelFolderIconResource:
		return "KControlPanelFolderIconResource"
	case KDropFolderIconResource:
		return "KDropFolderIconResource"
	case KExtensionsFolderIconResource:
		return "KExtensionsFolderIconResource"
	case KFontsFolderIconResource:
		return "KFontsFolderIconResource"
	case KFullTrashIconResource:
		return "KFullTrashIconResource"
	case KMountedFolderIconResource:
		return "KMountedFolderIconResource"
	case KOwnedFolderIconResource:
		return "KOwnedFolderIconResource"
	case KPreferencesFolderIconResource:
		return "KPreferencesFolderIconResource"
	case KPrintMonitorFolderIconResource:
		return "KPrintMonitorFolderIconResource"
	case KSharedFolderIconResource:
		return "KSharedFolderIconResource"
	case KStartupFolderIconResourceValue:
		return "KStartupFolderIconResourceValue"
	default:
		return fmt.Sprintf("KStartupFolderIconResource(%d)", e)
	}
}

const KSystemIconsCreator uint = 'm'<<24 | 'a'<<16 | 'c'<<8 | 's' // 'macs'

type KTXN int

const (
	KTXNATSUIIsNotInstalledErr KTXN = -22016
	KTXNAlreadyInitializedErr KTXN = -22012
	KTXNAttributeTagInvalidForRunErr KTXN = -22009
	KTXNBadDefaultFileTypeWarning KTXN = -22005
	KTXNCannotAddFrameErr KTXN = -22001
	KTXNCannotSetAutoIndentErr KTXN = -22006
	KTXNCannotTurnTSMOffWhenUsingUnicodeErr KTXN = -22013
	KTXNCopyNotAllowedInEchoModeErr KTXN = -22014
	KTXNDataTypeNotAllowedErr KTXN = -22015
	KTXNEndIterationErr KTXN = -22000
	KTXNIllegalToCrossDataBoundariesErr KTXN = -22003
	KTXNInvalidFrameIDErr KTXN = -22002
	KTXNInvalidRunIndex KTXN = -22011
	KTXNNoMatchErr KTXN = -22008
	KTXNOutsideOfFrameErr KTXN = -22018
	KTXNOutsideOfLineErr KTXN = -22017
	KTXNRunIndexOutofBoundsErr KTXN = -22007
	KTXNSomeOrAllTagsInvalidForRunErr KTXN = -22010
	KTXNUserCanceledOperationErr KTXN = -22004
)

func (e KTXN) String() string {
	switch e {
	case KTXNATSUIIsNotInstalledErr:
		return "KTXNATSUIIsNotInstalledErr"
	case KTXNAlreadyInitializedErr:
		return "KTXNAlreadyInitializedErr"
	case KTXNAttributeTagInvalidForRunErr:
		return "KTXNAttributeTagInvalidForRunErr"
	case KTXNBadDefaultFileTypeWarning:
		return "KTXNBadDefaultFileTypeWarning"
	case KTXNCannotAddFrameErr:
		return "KTXNCannotAddFrameErr"
	case KTXNCannotSetAutoIndentErr:
		return "KTXNCannotSetAutoIndentErr"
	case KTXNCannotTurnTSMOffWhenUsingUnicodeErr:
		return "KTXNCannotTurnTSMOffWhenUsingUnicodeErr"
	case KTXNCopyNotAllowedInEchoModeErr:
		return "KTXNCopyNotAllowedInEchoModeErr"
	case KTXNDataTypeNotAllowedErr:
		return "KTXNDataTypeNotAllowedErr"
	case KTXNEndIterationErr:
		return "KTXNEndIterationErr"
	case KTXNIllegalToCrossDataBoundariesErr:
		return "KTXNIllegalToCrossDataBoundariesErr"
	case KTXNInvalidFrameIDErr:
		return "KTXNInvalidFrameIDErr"
	case KTXNInvalidRunIndex:
		return "KTXNInvalidRunIndex"
	case KTXNNoMatchErr:
		return "KTXNNoMatchErr"
	case KTXNOutsideOfFrameErr:
		return "KTXNOutsideOfFrameErr"
	case KTXNOutsideOfLineErr:
		return "KTXNOutsideOfLineErr"
	case KTXNRunIndexOutofBoundsErr:
		return "KTXNRunIndexOutofBoundsErr"
	case KTXNSomeOrAllTagsInvalidForRunErr:
		return "KTXNSomeOrAllTagsInvalidForRunErr"
	case KTXNUserCanceledOperationErr:
		return "KTXNUserCanceledOperationErr"
	default:
		return fmt.Sprintf("KTXN(%d)", e)
	}
}

type KTextUnsupportedEncodingErr int

const (
	// KTECArrayFullErr: # Discussion
	KTECArrayFullErr KTextUnsupportedEncodingErr = -8751
	KTECBadTextRunErr KTextUnsupportedEncodingErr = -8752
	// KTECBufferBelowMinimumSizeErr: The output text buffer is too small to accommodatethe result of processing of the first input text element.
	KTECBufferBelowMinimumSizeErr KTextUnsupportedEncodingErr = -8750
	// KTECCorruptConverterErr: The converter object is invalid.
	KTECCorruptConverterErr KTextUnsupportedEncodingErr = -8748
	// KTECDirectionErr: An error, such as a direction stack overflow,occurred in directionality processing.
	KTECDirectionErr KTextUnsupportedEncodingErr = -8756
	// KTECGlobalsUnavailableErr: Global variables have already been deallocated,premature termination.
	KTECGlobalsUnavailableErr KTextUnsupportedEncodingErr = -8770
	// KTECIncompleteElementErr: The input text ends with a text elementthat might be incomplete, or contains a text element that is too longfor the internal buffers.
	KTECIncompleteElementErr KTextUnsupportedEncodingErr = -8755
	// KTECItemUnavailableErr: An item (for example, a name) is not availablefor the specified region (and encoding, if relevant).
	KTECItemUnavailableErr KTextUnsupportedEncodingErr = -8771
	// KTECMissingTableErr: The specified encoding is partially supported,but a specific table required for this function is missing.
	KTECMissingTableErr KTextUnsupportedEncodingErr = -8745
	// KTECNeedFlushStatus: The application disposed of a converterobject by calling TECDisposeConverter, but there is still text containedin internal buffers.
	KTECNeedFlushStatus KTextUnsupportedEncodingErr = -8784
	// KTECNoConversionPathErr: The converter supports both the source andtarget encodings, but cannot convert between them either directlyor indirectly.
	KTECNoConversionPathErr KTextUnsupportedEncodingErr = -8749
	// KTECOutputBufferFullStatus: The converter successfully converted partof the input text, but the output buffer was not large enough toaccommodate the entire input text after conversion.
	KTECOutputBufferFullStatus KTextUnsupportedEncodingErr = -8785
	// KTECPartialCharErr: The input text ends in the middle of a multibytecharacter and conversion stopped.
	KTECPartialCharErr KTextUnsupportedEncodingErr = -8753
	// KTECTableChecksumErr: A specific table required for this functionhas a checksum error, indicating that it has become corrupted.
	KTECTableChecksumErr KTextUnsupportedEncodingErr = -8746
	// KTECTableFormatErr: The table format is either invalid or itcannot be handled by the current version of the code.
	KTECTableFormatErr KTextUnsupportedEncodingErr = -8747
	// KTECUnmappableElementErr: An input text element cannot be mapped tothe specified output encoding(s) using the specified options.
	KTECUnmappableElementErr KTextUnsupportedEncodingErr = -8754
	// KTECUsedFallbacksStatus: The function has completely converted theinput string to the specified target using one or more fallbacks.For the Unicode Converter, this status code can only occur if the `kUnicodeUseFallbacksBit`control flag is set.
	KTECUsedFallbacksStatus KTextUnsupportedEncodingErr = -8783
	// KTextMalformedInputErr: The text input contains a sequence thatis not legal in the specified encoding, such as a DBCS high byte followedby an invalid low byte (0x8120 in Shift-JIS).
	KTextMalformedInputErr KTextUnsupportedEncodingErr = -8739
	// KTextUndefinedElementErr: The text input contains a code point thatis undefined in the specified encoding.
	KTextUndefinedElementErr KTextUnsupportedEncodingErr = -8740
	// KTextUnsupportedEncodingErrValue: The encoding or mapping is not supportedfor this function by the current set of tables or plug-ins.
	KTextUnsupportedEncodingErrValue KTextUnsupportedEncodingErr = -8738
	UnicodeBufErr KTextUnsupportedEncodingErr = -8764
	UnicodeCharErr KTextUnsupportedEncodingErr = -8763
	UnicodeChecksumErr KTextUnsupportedEncodingErr = -8769
	UnicodeContextualErr KTextUnsupportedEncodingErr = -8758
	UnicodeDirectionErr KTextUnsupportedEncodingErr = -8759
	UnicodeElementErr KTextUnsupportedEncodingErr = -8762
	UnicodeFallbacksErr KTextUnsupportedEncodingErr = -8766
	UnicodeNoTableErr KTextUnsupportedEncodingErr = -8768
	UnicodeNotFoundErr KTextUnsupportedEncodingErr = -8761
	UnicodePartConvertErr KTextUnsupportedEncodingErr = -8765
	UnicodeTableFormatErr KTextUnsupportedEncodingErr = -8760
	UnicodeTextEncodingDataErr KTextUnsupportedEncodingErr = -8757
	UnicodeVariantErr KTextUnsupportedEncodingErr = -8767
)

func (e KTextUnsupportedEncodingErr) String() string {
	switch e {
	case KTECArrayFullErr:
		return "KTECArrayFullErr"
	case KTECBadTextRunErr:
		return "KTECBadTextRunErr"
	case KTECBufferBelowMinimumSizeErr:
		return "KTECBufferBelowMinimumSizeErr"
	case KTECCorruptConverterErr:
		return "KTECCorruptConverterErr"
	case KTECDirectionErr:
		return "KTECDirectionErr"
	case KTECGlobalsUnavailableErr:
		return "KTECGlobalsUnavailableErr"
	case KTECIncompleteElementErr:
		return "KTECIncompleteElementErr"
	case KTECItemUnavailableErr:
		return "KTECItemUnavailableErr"
	case KTECMissingTableErr:
		return "KTECMissingTableErr"
	case KTECNeedFlushStatus:
		return "KTECNeedFlushStatus"
	case KTECNoConversionPathErr:
		return "KTECNoConversionPathErr"
	case KTECOutputBufferFullStatus:
		return "KTECOutputBufferFullStatus"
	case KTECPartialCharErr:
		return "KTECPartialCharErr"
	case KTECTableChecksumErr:
		return "KTECTableChecksumErr"
	case KTECTableFormatErr:
		return "KTECTableFormatErr"
	case KTECUnmappableElementErr:
		return "KTECUnmappableElementErr"
	case KTECUsedFallbacksStatus:
		return "KTECUsedFallbacksStatus"
	case KTextMalformedInputErr:
		return "KTextMalformedInputErr"
	case KTextUndefinedElementErr:
		return "KTextUndefinedElementErr"
	case KTextUnsupportedEncodingErrValue:
		return "KTextUnsupportedEncodingErrValue"
	case UnicodeBufErr:
		return "UnicodeBufErr"
	case UnicodeCharErr:
		return "UnicodeCharErr"
	case UnicodeChecksumErr:
		return "UnicodeChecksumErr"
	case UnicodeContextualErr:
		return "UnicodeContextualErr"
	case UnicodeDirectionErr:
		return "UnicodeDirectionErr"
	case UnicodeElementErr:
		return "UnicodeElementErr"
	case UnicodeFallbacksErr:
		return "UnicodeFallbacksErr"
	case UnicodeNoTableErr:
		return "UnicodeNoTableErr"
	case UnicodeNotFoundErr:
		return "UnicodeNotFoundErr"
	case UnicodePartConvertErr:
		return "UnicodePartConvertErr"
	case UnicodeTableFormatErr:
		return "UnicodeTableFormatErr"
	case UnicodeTextEncodingDataErr:
		return "UnicodeTextEncodingDataErr"
	case UnicodeVariantErr:
		return "UnicodeVariantErr"
	default:
		return fmt.Sprintf("KTextUnsupportedEncodingErr(%d)", e)
	}
}

type KToolbar uint

const (
	KToolbarAdvancedIcon KToolbar = 't'<<24 | 'b'<<16 | 'a'<<8 | 'v' // 'tbav'
	KToolbarApplicationsFolderIcon KToolbar = 't'<<24 | 'A'<<16 | 'p'<<8 | 's' // 'tAps'
	KToolbarCustomizeIcon KToolbar = 't'<<24 | 'c'<<16 | 'u'<<8 | 's' // 'tcus'
	KToolbarDeleteIcon KToolbar = 't'<<24 | 'd'<<16 | 'e'<<8 | 'l' // 'tdel'
	KToolbarDesktopFolderIcon KToolbar = 't'<<24 | 'D'<<16 | 's'<<8 | 'k' // 'tDsk'
	KToolbarDocumentsFolderIcon KToolbar = 't'<<24 | 'D'<<16 | 'o'<<8 | 'c' // 'tDoc'
	KToolbarDownloadsFolderIcon KToolbar = 't'<<24 | 'D'<<16 | 'w'<<8 | 'n' // 'tDwn'
	KToolbarFavoritesIcon KToolbar = 't'<<24 | 'f'<<16 | 'a'<<8 | 'v' // 'tfav'
	KToolbarHomeIcon KToolbar = 't'<<24 | 'h'<<16 | 'o'<<8 | 'm' // 'thom'
	KToolbarInfoIcon KToolbar = 't'<<24 | 'b'<<16 | 'i'<<8 | 'n' // 'tbin'
	KToolbarLabelsIcon KToolbar = 't'<<24 | 'b'<<16 | 'l'<<8 | 'b' // 'tblb'
	KToolbarLibraryFolderIcon KToolbar = 't'<<24 | 'L'<<16 | 'i'<<8 | 'b' // 'tLib'
	KToolbarMovieFolderIcon KToolbar = 't'<<24 | 'M'<<16 | 'o'<<8 | 'v' // 'tMov'
	KToolbarMusicFolderIcon KToolbar = 't'<<24 | 'M'<<16 | 'u'<<8 | 's' // 'tMus'
	KToolbarPicturesFolderIcon KToolbar = 't'<<24 | 'P'<<16 | 'i'<<8 | 'c' // 'tPic'
	KToolbarPublicFolderIcon KToolbar = 't'<<24 | 'P'<<16 | 'u'<<8 | 'b' // 'tPub'
	KToolbarSitesFolderIcon KToolbar = 't'<<24 | 'S'<<16 | 't'<<8 | 's' // 'tSts'
	KToolbarUtilitiesFolderIcon KToolbar = 't'<<24 | 'U'<<16 | 't'<<8 | 'l' // 'tUtl'
)

func (e KToolbar) String() string {
	switch e {
	case KToolbarAdvancedIcon:
		return "KToolbarAdvancedIcon"
	case KToolbarApplicationsFolderIcon:
		return "KToolbarApplicationsFolderIcon"
	case KToolbarCustomizeIcon:
		return "KToolbarCustomizeIcon"
	case KToolbarDeleteIcon:
		return "KToolbarDeleteIcon"
	case KToolbarDesktopFolderIcon:
		return "KToolbarDesktopFolderIcon"
	case KToolbarDocumentsFolderIcon:
		return "KToolbarDocumentsFolderIcon"
	case KToolbarDownloadsFolderIcon:
		return "KToolbarDownloadsFolderIcon"
	case KToolbarFavoritesIcon:
		return "KToolbarFavoritesIcon"
	case KToolbarHomeIcon:
		return "KToolbarHomeIcon"
	case KToolbarInfoIcon:
		return "KToolbarInfoIcon"
	case KToolbarLabelsIcon:
		return "KToolbarLabelsIcon"
	case KToolbarLibraryFolderIcon:
		return "KToolbarLibraryFolderIcon"
	case KToolbarMovieFolderIcon:
		return "KToolbarMovieFolderIcon"
	case KToolbarMusicFolderIcon:
		return "KToolbarMusicFolderIcon"
	case KToolbarPicturesFolderIcon:
		return "KToolbarPicturesFolderIcon"
	case KToolbarPublicFolderIcon:
		return "KToolbarPublicFolderIcon"
	case KToolbarSitesFolderIcon:
		return "KToolbarSitesFolderIcon"
	case KToolbarUtilitiesFolderIcon:
		return "KToolbarUtilitiesFolderIcon"
	default:
		return fmt.Sprintf("KToolbar(%d)", e)
	}
}

type KUC int

const (
	KUCOutputBufferTooSmall KUC = -25340
	KUCTextBreakLocatorMissingType KUC = -25341
	KUCTokenNotFound KUC = -25346
	KUCTokenizerIterationFinished KUC = -25344
	KUCTokenizerUnknownLang KUC = -25345
)

func (e KUC) String() string {
	switch e {
	case KUCOutputBufferTooSmall:
		return "KUCOutputBufferTooSmall"
	case KUCTextBreakLocatorMissingType:
		return "KUCTextBreakLocatorMissingType"
	case KUCTokenNotFound:
		return "KUCTokenNotFound"
	case KUCTokenizerIterationFinished:
		return "KUCTokenizerIterationFinished"
	case KUCTokenizerUnknownLang:
		return "KUCTokenizerUnknownLang"
	default:
		return fmt.Sprintf("KUC(%d)", e)
	}
}

type KUCCollate uint

const (
	// KUCCollateCaseInsensitiveMask: If the corresponding bit is set, then uppercase and titlecase characters are treated as equivalent to the corresponding lowercase characters.
	KUCCollateCaseInsensitiveMask KUCCollate = 8
	// KUCCollateComposeInsensitiveMask: # Discussion
	KUCCollateComposeInsensitiveMask KUCCollate = 2
	// KUCCollateDiacritInsensitiveMask: If the corresponding bit is set, then characters with diacritics are treated as equivalent to the corresponding characters without diacritics.
	KUCCollateDiacritInsensitiveMask KUCCollate = 16
	// KUCCollateDigitsAsNumberMask: # Discussion
	KUCCollateDigitsAsNumberMask KUCCollate = 131072
	// KUCCollateDigitsOverrideMask: # Discussion
	KUCCollateDigitsOverrideMask KUCCollate = 65536
	// KUCCollatePunctuationSignificantMask: # Discussion
	KUCCollatePunctuationSignificantMask KUCCollate = 32768
	// KUCCollateWidthInsensitiveMask: # Discussion
	KUCCollateWidthInsensitiveMask KUCCollate = 4
)

func (e KUCCollate) String() string {
	switch e {
	case KUCCollateCaseInsensitiveMask:
		return "KUCCollateCaseInsensitiveMask"
	case KUCCollateComposeInsensitiveMask:
		return "KUCCollateComposeInsensitiveMask"
	case KUCCollateDiacritInsensitiveMask:
		return "KUCCollateDiacritInsensitiveMask"
	case KUCCollateDigitsAsNumberMask:
		return "KUCCollateDigitsAsNumberMask"
	case KUCCollateDigitsOverrideMask:
		return "KUCCollateDigitsOverrideMask"
	case KUCCollatePunctuationSignificantMask:
		return "KUCCollatePunctuationSignificantMask"
	case KUCCollateWidthInsensitiveMask:
		return "KUCCollateWidthInsensitiveMask"
	default:
		return fmt.Sprintf("KUCCollate(%d)", e)
	}
}

const KUCCollateStandardOptions uint = 2

type KUCCollateType uint

const (
	// KUCCollateTypeShiftBits: # Discussion
	KUCCollateTypeShiftBits KUCCollateType = 24
	// KUCCollateTypeSourceMask: You can use this mask, in conjunction with the `kUCCollateTypeShiftBits` constant, to obtain a value identifying a fixed ordering scheme.
	KUCCollateTypeSourceMask KUCCollateType = 255
)

func (e KUCCollateType) String() string {
	switch e {
	case KUCCollateTypeShiftBits:
		return "KUCCollateTypeShiftBits"
	case KUCCollateTypeSourceMask:
		return "KUCCollateTypeSourceMask"
	default:
		return fmt.Sprintf("KUCCollateType(%d)", e)
	}
}

const KUCCollateTypeHFSExtended uint = 1

const KUCCollateTypeMask uint = 0

type KUCKey uint

const (
	// KUCKeyLayoutFeatureInfoFormat: The format of a structure of type UCKeyLayoutFeatureInfo.
	KUCKeyLayoutFeatureInfoFormat KUCKey = 8193
	// KUCKeyLayoutHeaderFormat: The format of a structure of type UCKeyboardLayout.
	KUCKeyLayoutHeaderFormat KUCKey = 4098
	// KUCKeyModifiersToTableNumFormat: The format of a structure of type UCKeyModifiersToTableNum.
	KUCKeyModifiersToTableNumFormat KUCKey = 12289
	// KUCKeySequenceDataIndexFormat: The format of a structure of type UCKeySequenceDataIndex.
	KUCKeySequenceDataIndexFormat KUCKey = 28673
	// KUCKeyStateRecordsIndexFormat: The format of a structure of type UCKeyStateRecordsIndex.
	KUCKeyStateRecordsIndexFormat KUCKey = 20481
	// KUCKeyStateTerminatorsFormat: The format of a structure of type UCKeyStateTerminators.
	KUCKeyStateTerminatorsFormat KUCKey = 24577
	// KUCKeyToCharTableIndexFormat: The format of a structure of type UCKeyToCharTableIndex.
	KUCKeyToCharTableIndexFormat KUCKey = 16385
)

func (e KUCKey) String() string {
	switch e {
	case KUCKeyLayoutFeatureInfoFormat:
		return "KUCKeyLayoutFeatureInfoFormat"
	case KUCKeyLayoutHeaderFormat:
		return "KUCKeyLayoutHeaderFormat"
	case KUCKeyModifiersToTableNumFormat:
		return "KUCKeyModifiersToTableNumFormat"
	case KUCKeySequenceDataIndexFormat:
		return "KUCKeySequenceDataIndexFormat"
	case KUCKeyStateRecordsIndexFormat:
		return "KUCKeyStateRecordsIndexFormat"
	case KUCKeyStateTerminatorsFormat:
		return "KUCKeyStateTerminatorsFormat"
	case KUCKeyToCharTableIndexFormat:
		return "KUCKeyToCharTableIndexFormat"
	default:
		return fmt.Sprintf("KUCKey(%d)", e)
	}
}

type KUCKeyAction uint

const (
	// KUCKeyActionAutoKey: The user has the key in an “auto-key” pressed state that is, the user is holding down the key for an extended period of time and is thereby generating multiple key strokes from the single key.
	KUCKeyActionAutoKey KUCKeyAction = 2
	// KUCKeyActionDisplay: The user is requesting information for key display, as in the Key Caps application.
	KUCKeyActionDisplay KUCKeyAction = 3
	// KUCKeyActionDown: The user is pressing the key.
	KUCKeyActionDown KUCKeyAction = 0
	// KUCKeyActionUp: The user is releasing the key.
	KUCKeyActionUp KUCKeyAction = 1
)

func (e KUCKeyAction) String() string {
	switch e {
	case KUCKeyActionAutoKey:
		return "KUCKeyActionAutoKey"
	case KUCKeyActionDisplay:
		return "KUCKeyActionDisplay"
	case KUCKeyActionDown:
		return "KUCKeyActionDown"
	case KUCKeyActionUp:
		return "KUCKeyActionUp"
	default:
		return fmt.Sprintf("KUCKeyAction(%d)", e)
	}
}

type KUCKeyOutput uint

const (
	// KUCKeyOutputGetIndexMask: You can use this mask to test the bits (0–13) in a [UCKeyOutput] value that provide the actual index to another structure.
	KUCKeyOutputGetIndexMask KUCKeyOutput = 16383
	// KUCKeyOutputSequenceIndexMask: If the bit specified by this mask is set, the [UCKeyOutput] value contains an index into a structure of type UCKeySequenceDataIndex.
	KUCKeyOutputSequenceIndexMask KUCKeyOutput = 32768
	// KUCKeyOutputStateIndexMask: If the bit specified by this mask is set, the UCKeyStateRecordsIndex [UCKeyOutput] value contains an index into a structure of type UCKeyStateRecordsIndex.
	KUCKeyOutputStateIndexMask KUCKeyOutput = 16384
	// KUCKeyOutputTestForIndexMask: # Discussion
	KUCKeyOutputTestForIndexMask KUCKeyOutput = 49152
)

func (e KUCKeyOutput) String() string {
	switch e {
	case KUCKeyOutputGetIndexMask:
		return "KUCKeyOutputGetIndexMask"
	case KUCKeyOutputSequenceIndexMask:
		return "KUCKeyOutputSequenceIndexMask"
	case KUCKeyOutputStateIndexMask:
		return "KUCKeyOutputStateIndexMask"
	case KUCKeyOutputTestForIndexMask:
		return "KUCKeyOutputTestForIndexMask"
	default:
		return fmt.Sprintf("KUCKeyOutput(%d)", e)
	}
}

type KUCKeyStateEntry uint

const (
	// KUCKeyStateEntryRangeFormat: Specifies that the entry format is that of a structure of type UCKeyStateEntryRange.
	KUCKeyStateEntryRangeFormat KUCKeyStateEntry = 2
	// KUCKeyStateEntryTerminalFormat: Specifies that the entry format is that of a structure of type UCKeyStateEntryTerminal.
	KUCKeyStateEntryTerminalFormat KUCKeyStateEntry = 1
)

func (e KUCKeyStateEntry) String() string {
	switch e {
	case KUCKeyStateEntryRangeFormat:
		return "KUCKeyStateEntryRangeFormat"
	case KUCKeyStateEntryTerminalFormat:
		return "KUCKeyStateEntryTerminalFormat"
	default:
		return fmt.Sprintf("KUCKeyStateEntry(%d)", e)
	}
}

const KUCKeyTranslateNoDeadKeysBit uint = 0

const KUCKeyTranslateNoDeadKeysMask uint = 1

type KUCTS int

const (
	KUCTSNoKeysAddedToObjectErr KUCTS = -25342
	KUCTSSearchListErr KUCTS = -25343
)

func (e KUCTS) String() string {
	switch e {
	case KUCTSNoKeysAddedToObjectErr:
		return "KUCTSNoKeysAddedToObjectErr"
	case KUCTSSearchListErr:
		return "KUCTSSearchListErr"
	default:
		return fmt.Sprintf("KUCTS(%d)", e)
	}
}

type KUCTSDirection uint

const (
	KUCTSDirectionNext KUCTSDirection = 0
	KUCTSDirectionPrevious KUCTSDirection = 1
)

func (e KUCTSDirection) String() string {
	switch e {
	case KUCTSDirectionNext:
		return "KUCTSDirectionNext"
	case KUCTSDirectionPrevious:
		return "KUCTSDirectionPrevious"
	default:
		return fmt.Sprintf("KUCTSDirection(%d)", e)
	}
}

type KUCTSOptions uint

const (
	KUCTSOptionsDataIsOrderedMask KUCTSOptions = 2
	KUCTSOptionsNoneMask KUCTSOptions = 0
	KUCTSOptionsReleaseStringMask KUCTSOptions = 1
)

func (e KUCTSOptions) String() string {
	switch e {
	case KUCTSOptionsDataIsOrderedMask:
		return "KUCTSOptionsDataIsOrderedMask"
	case KUCTSOptionsNoneMask:
		return "KUCTSOptionsNoneMask"
	case KUCTSOptionsReleaseStringMask:
		return "KUCTSOptionsReleaseStringMask"
	default:
		return fmt.Sprintf("KUCTSOptions(%d)", e)
	}
}

type KUCTextBreak uint

const (
	// KUCTextBreakCharMask: If the bit specified by this mask is set, boundaries of characters may be located (with surrogate pairs treated as a single character).
	KUCTextBreakCharMask KUCTextBreak = 1
	// KUCTextBreakClusterMask: # Discussion
	KUCTextBreakClusterMask KUCTextBreak = 4
	// KUCTextBreakGoBackwardsMask: # Discussion
	KUCTextBreakGoBackwardsMask KUCTextBreak = 2
	// KUCTextBreakIterateMask: # Discussion
	KUCTextBreakIterateMask KUCTextBreak = 4
	// KUCTextBreakLeadingEdgeMask: # Discussion
	KUCTextBreakLeadingEdgeMask KUCTextBreak = 1
	// KUCTextBreakLineMask: If the bit specified by this mask is set, potential line breaks may be located.
	KUCTextBreakLineMask KUCTextBreak = 64
	KUCTextBreakParagraphMask KUCTextBreak = 256
	// KUCTextBreakWordMask: If the bit specified by this mask is set, boundaries of words may be located.
	KUCTextBreakWordMask KUCTextBreak = 16
)

func (e KUCTextBreak) String() string {
	switch e {
	case KUCTextBreakCharMask:
		return "KUCTextBreakCharMask"
	case KUCTextBreakClusterMask:
		return "KUCTextBreakClusterMask"
	case KUCTextBreakGoBackwardsMask:
		return "KUCTextBreakGoBackwardsMask"
	case KUCTextBreakLineMask:
		return "KUCTextBreakLineMask"
	case KUCTextBreakParagraphMask:
		return "KUCTextBreakParagraphMask"
	case KUCTextBreakWordMask:
		return "KUCTextBreakWordMask"
	default:
		return fmt.Sprintf("KUCTextBreak(%d)", e)
	}
}

const KUCTypeSelectMaxListSize uint = 0

type KUR int

const (
	KURL68kNotSupportedError KUR = -30788
	KURLAccessNotAvailableError KUR = -30787
	KURLAuthenticationError KUR = -30776
	KURLDestinationExistsError KUR = -30772
	KURLExtensionFailureError KUR = -30785
	KURLFileEmptyError KUR = -30783
	KURLInvalidCallError KUR = -30781
	KURLInvalidConfigurationError KUR = -30786
	KURLInvalidURLError KUR = -30773
	KURLInvalidURLReferenceError KUR = -30770
	KURLProgressAlreadyDisplayedError KUR = -30771
	KURLPropertyBufferTooSmallError KUR = -30779
	KURLPropertyNotYetKnownError KUR = -30777
	KURLServerBusyError KUR = -30775
	KURLUnknownPropertyError KUR = -30778
	KURLUnsettablePropertyError KUR = -30780
	KURLUnsupportedSchemeError KUR = -30774
)

func (e KUR) String() string {
	switch e {
	case KURL68kNotSupportedError:
		return "KURL68kNotSupportedError"
	case KURLAccessNotAvailableError:
		return "KURLAccessNotAvailableError"
	case KURLAuthenticationError:
		return "KURLAuthenticationError"
	case KURLDestinationExistsError:
		return "KURLDestinationExistsError"
	case KURLExtensionFailureError:
		return "KURLExtensionFailureError"
	case KURLFileEmptyError:
		return "KURLFileEmptyError"
	case KURLInvalidCallError:
		return "KURLInvalidCallError"
	case KURLInvalidConfigurationError:
		return "KURLInvalidConfigurationError"
	case KURLInvalidURLError:
		return "KURLInvalidURLError"
	case KURLInvalidURLReferenceError:
		return "KURLInvalidURLReferenceError"
	case KURLProgressAlreadyDisplayedError:
		return "KURLProgressAlreadyDisplayedError"
	case KURLPropertyBufferTooSmallError:
		return "KURLPropertyBufferTooSmallError"
	case KURLPropertyNotYetKnownError:
		return "KURLPropertyNotYetKnownError"
	case KURLServerBusyError:
		return "KURLServerBusyError"
	case KURLUnknownPropertyError:
		return "KURLUnknownPropertyError"
	case KURLUnsettablePropertyError:
		return "KURLUnsettablePropertyError"
	case KURLUnsupportedSchemeError:
		return "KURLUnsupportedSchemeError"
	default:
		return fmt.Sprintf("KUR(%d)", e)
	}
}

type KUSB int

const (
	KUSBAbortedError KUSB = -6982
	KUSBAlreadyOpenErr KUSB = -6991
	KUSBBadDispatchTable KUSB = -6950
	KUSBBitstufErr KUSB = -6914
	KUSBBufOvrRunErr KUSB = -6904
	KUSBBufUnderRunErr KUSB = -6903
	KUSBCRCErr KUSB = -6915
	KUSBCompletionError KUSB = -6984
	KUSBDataToggleErr KUSB = -6913
	KUSBDeviceBusy KUSB = -6977
	KUSBDeviceDisconnected KUSB = -6972
	KUSBDeviceErr KUSB = -6989
	KUSBDeviceNotSuspended KUSB = -6973
	KUSBDevicePowerProblem KUSB = -6976
	KUSBDeviceSuspended KUSB = -6974
	KUSBEndpointStallErr KUSB = -6912
	KUSBFlagsError KUSB = -6983
	KUSBIncorrectTypeErr KUSB = -6995
	KUSBInternalErr KUSB = -6999
	KUSBInvalidBuffer KUSB = -6975
	KUSBLinkErr KUSB = -6916
	KUSBNoBandwidthError KUSB = -6981
	KUSBNoDelay KUSB = 0
	KUSBNoDeviceErr KUSB = -6990
	KUSBNoErr KUSB = 0
	KUSBNoTran KUSB = 0
	KUSBNotFound KUSB = -6987
	KUSBNotHandled KUSB = -6987
	KUSBNotRespondingErr KUSB = -6911
	KUSBNotSent1Err KUSB = -6902
	KUSBNotSent2Err KUSB = -6901
	KUSBOutOfMemoryErr KUSB = -6988
	KUSBOverRunErr KUSB = -6908
	KUSBPBLengthError KUSB = -6985
	KUSBPBVersionError KUSB = -6986
	KUSBPIDCheckErr KUSB = -6910
	KUSBPending KUSB = 1
	KUSBPipeIdleError KUSB = -6980
	KUSBPipeStalledError KUSB = -6979
	KUSBPortDisabled KUSB = -6969
	KUSBQueueAborted KUSB = -6970
	KUSBQueueFull KUSB = -6948
	KUSBRes1Err KUSB = -6906
	KUSBRes2Err KUSB = -6905
	KUSBRqErr KUSB = -6994
	KUSBTimedOut KUSB = -6971
	KUSBTooManyPipesErr KUSB = -6996
	KUSBTooManyTransactionsErr KUSB = -6992
	KUSBUnderRunErr KUSB = -6907
	KUSBUnknownDeviceErr KUSB = -6998
	KUSBUnknownInterfaceErr KUSB = -6978
	KUSBUnknownNotification KUSB = -6949
	KUSBUnknownPipeErr KUSB = -6997
	KUSBUnknownRequestErr KUSB = -6993
	KUSBWrongPIDErr KUSB = -6909
)

func (e KUSB) String() string {
	switch e {
	case KUSBAbortedError:
		return "KUSBAbortedError"
	case KUSBAlreadyOpenErr:
		return "KUSBAlreadyOpenErr"
	case KUSBBadDispatchTable:
		return "KUSBBadDispatchTable"
	case KUSBBitstufErr:
		return "KUSBBitstufErr"
	case KUSBBufOvrRunErr:
		return "KUSBBufOvrRunErr"
	case KUSBBufUnderRunErr:
		return "KUSBBufUnderRunErr"
	case KUSBCRCErr:
		return "KUSBCRCErr"
	case KUSBCompletionError:
		return "KUSBCompletionError"
	case KUSBDataToggleErr:
		return "KUSBDataToggleErr"
	case KUSBDeviceBusy:
		return "KUSBDeviceBusy"
	case KUSBDeviceDisconnected:
		return "KUSBDeviceDisconnected"
	case KUSBDeviceErr:
		return "KUSBDeviceErr"
	case KUSBDeviceNotSuspended:
		return "KUSBDeviceNotSuspended"
	case KUSBDevicePowerProblem:
		return "KUSBDevicePowerProblem"
	case KUSBDeviceSuspended:
		return "KUSBDeviceSuspended"
	case KUSBEndpointStallErr:
		return "KUSBEndpointStallErr"
	case KUSBFlagsError:
		return "KUSBFlagsError"
	case KUSBIncorrectTypeErr:
		return "KUSBIncorrectTypeErr"
	case KUSBInternalErr:
		return "KUSBInternalErr"
	case KUSBInvalidBuffer:
		return "KUSBInvalidBuffer"
	case KUSBLinkErr:
		return "KUSBLinkErr"
	case KUSBNoBandwidthError:
		return "KUSBNoBandwidthError"
	case KUSBNoDelay:
		return "KUSBNoDelay"
	case KUSBNoDeviceErr:
		return "KUSBNoDeviceErr"
	case KUSBNotFound:
		return "KUSBNotFound"
	case KUSBNotRespondingErr:
		return "KUSBNotRespondingErr"
	case KUSBNotSent1Err:
		return "KUSBNotSent1Err"
	case KUSBNotSent2Err:
		return "KUSBNotSent2Err"
	case KUSBOutOfMemoryErr:
		return "KUSBOutOfMemoryErr"
	case KUSBOverRunErr:
		return "KUSBOverRunErr"
	case KUSBPBLengthError:
		return "KUSBPBLengthError"
	case KUSBPBVersionError:
		return "KUSBPBVersionError"
	case KUSBPIDCheckErr:
		return "KUSBPIDCheckErr"
	case KUSBPending:
		return "KUSBPending"
	case KUSBPipeIdleError:
		return "KUSBPipeIdleError"
	case KUSBPipeStalledError:
		return "KUSBPipeStalledError"
	case KUSBPortDisabled:
		return "KUSBPortDisabled"
	case KUSBQueueAborted:
		return "KUSBQueueAborted"
	case KUSBQueueFull:
		return "KUSBQueueFull"
	case KUSBRes1Err:
		return "KUSBRes1Err"
	case KUSBRes2Err:
		return "KUSBRes2Err"
	case KUSBRqErr:
		return "KUSBRqErr"
	case KUSBTimedOut:
		return "KUSBTimedOut"
	case KUSBTooManyPipesErr:
		return "KUSBTooManyPipesErr"
	case KUSBTooManyTransactionsErr:
		return "KUSBTooManyTransactionsErr"
	case KUSBUnderRunErr:
		return "KUSBUnderRunErr"
	case KUSBUnknownDeviceErr:
		return "KUSBUnknownDeviceErr"
	case KUSBUnknownInterfaceErr:
		return "KUSBUnknownInterfaceErr"
	case KUSBUnknownNotification:
		return "KUSBUnknownNotification"
	case KUSBUnknownPipeErr:
		return "KUSBUnknownPipeErr"
	case KUSBUnknownRequestErr:
		return "KUSBUnknownRequestErr"
	case KUSBWrongPIDErr:
		return "KUSBWrongPIDErr"
	default:
		return fmt.Sprintf("KUSB(%d)", e)
	}
}

type KUSBInternal int

const (
	KUSBInternalReserved1 KUSBInternal = -6960
	KUSBInternalReserved10 KUSBInternal = -6951
	KUSBInternalReserved2 KUSBInternal = -6959
	KUSBInternalReserved3 KUSBInternal = -6958
	KUSBInternalReserved4 KUSBInternal = -6957
	KUSBInternalReserved5 KUSBInternal = -6956
	KUSBInternalReserved6 KUSBInternal = -6955
	KUSBInternalReserved7 KUSBInternal = -6954
	KUSBInternalReserved8 KUSBInternal = -6953
	KUSBInternalReserved9 KUSBInternal = -6952
)

func (e KUSBInternal) String() string {
	switch e {
	case KUSBInternalReserved1:
		return "KUSBInternalReserved1"
	case KUSBInternalReserved10:
		return "KUSBInternalReserved10"
	case KUSBInternalReserved2:
		return "KUSBInternalReserved2"
	case KUSBInternalReserved3:
		return "KUSBInternalReserved3"
	case KUSBInternalReserved4:
		return "KUSBInternalReserved4"
	case KUSBInternalReserved5:
		return "KUSBInternalReserved5"
	case KUSBInternalReserved6:
		return "KUSBInternalReserved6"
	case KUSBInternalReserved7:
		return "KUSBInternalReserved7"
	case KUSBInternalReserved8:
		return "KUSBInternalReserved8"
	case KUSBInternalReserved9:
		return "KUSBInternalReserved9"
	default:
		return fmt.Sprintf("KUSBInternal(%d)", e)
	}
}

type KUTCUnderflowErr int

const (
	KIllegalClockValueErr KUTCUnderflowErr = -8852
	KUTCOverflowErr KUTCUnderflowErr = -8851
	KUTCUnderflowErrValue KUTCUnderflowErr = -8850
)

func (e KUTCUnderflowErr) String() string {
	switch e {
	case KIllegalClockValueErr:
		return "KIllegalClockValueErr"
	case KUTCOverflowErr:
		return "KUTCOverflowErr"
	case KUTCUnderflowErrValue:
		return "KUTCUnderflowErrValue"
	default:
		return fmt.Sprintf("KUTCUnderflowErr(%d)", e)
	}
}

const KUnicodeCollationClass uint = 'u'<<24 | 'c'<<16 | 'o'<<8 | 'l' // 'ucol'

const KUnicodeTextBreakClass uint = 'u'<<24 | 'b'<<16 | 'r'<<8 | 'k' // 'ubrk'

type KUserFolderIcon uint

const (
	KGroupIcon KUserFolderIcon = 'g'<<24 | 'r'<<16 | 'u'<<8 | 'p' // 'grup'
	KGuestUserIcon KUserFolderIcon = 'g'<<24 | 'u'<<16 | 's'<<8 | 'r' // 'gusr'
	KOwnerIcon KUserFolderIcon = 's'<<24 | 'u'<<16 | 's'<<8 | 'r' // 'susr'
	KUserFolderIconValue KUserFolderIcon = 'u'<<24 | 'f'<<16 | 'l'<<8 | 'd' // 'ufld'
	KUserIcon KUserFolderIcon = 'u'<<24 | 's'<<16 | 'e'<<8 | 'r' // 'user'
	KWorkgroupFolderIcon KUserFolderIcon = 'w'<<24 | 'f'<<16 | 'l'<<8 | 'd' // 'wfld'
)

func (e KUserFolderIcon) String() string {
	switch e {
	case KGroupIcon:
		return "KGroupIcon"
	case KGuestUserIcon:
		return "KGuestUserIcon"
	case KOwnerIcon:
		return "KOwnerIcon"
	case KUserFolderIconValue:
		return "KUserFolderIconValue"
	case KUserIcon:
		return "KUserIcon"
	case KWorkgroupFolderIcon:
		return "KWorkgroupFolderIcon"
	default:
		return fmt.Sprintf("KUserFolderIcon(%d)", e)
	}
}

type Kernel int

const (
	KernelAlreadyFreeErr Kernel = -2421
	KernelAsyncReceiveLimitErr Kernel = -2414
	KernelAsyncSendLimitErr Kernel = -2413
	KernelAttributeErr Kernel = -2412
	KernelCanceledErr Kernel = -2402
	KernelDeletePermissionErr Kernel = -2410
	KernelExceptionErr Kernel = -2418
	KernelExecutePermissionErr Kernel = -2409
	KernelExecutionLevelErr Kernel = -2411
	KernelIDErr Kernel = -2419
	KernelInUseErr Kernel = -2416
	KernelIncompleteErr Kernel = -2401
	KernelObjectExistsErr Kernel = -2406
	KernelOptionsErr Kernel = -2403
	KernelPrivilegeErr Kernel = -2404
	KernelReadPermissionErr Kernel = -2408
	KernelReturnValueErr Kernel = -2422
	KernelTerminatedErr Kernel = -2417
	KernelTimeoutErr Kernel = -2415
	KernelUnrecoverableErr Kernel = -2499
	KernelUnsupportedErr Kernel = -2405
	KernelWritePermissionErr Kernel = -2407
)

func (e Kernel) String() string {
	switch e {
	case KernelAlreadyFreeErr:
		return "KernelAlreadyFreeErr"
	case KernelAsyncReceiveLimitErr:
		return "KernelAsyncReceiveLimitErr"
	case KernelAsyncSendLimitErr:
		return "KernelAsyncSendLimitErr"
	case KernelAttributeErr:
		return "KernelAttributeErr"
	case KernelCanceledErr:
		return "KernelCanceledErr"
	case KernelDeletePermissionErr:
		return "KernelDeletePermissionErr"
	case KernelExceptionErr:
		return "KernelExceptionErr"
	case KernelExecutePermissionErr:
		return "KernelExecutePermissionErr"
	case KernelExecutionLevelErr:
		return "KernelExecutionLevelErr"
	case KernelIDErr:
		return "KernelIDErr"
	case KernelInUseErr:
		return "KernelInUseErr"
	case KernelIncompleteErr:
		return "KernelIncompleteErr"
	case KernelObjectExistsErr:
		return "KernelObjectExistsErr"
	case KernelOptionsErr:
		return "KernelOptionsErr"
	case KernelPrivilegeErr:
		return "KernelPrivilegeErr"
	case KernelReadPermissionErr:
		return "KernelReadPermissionErr"
	case KernelReturnValueErr:
		return "KernelReturnValueErr"
	case KernelTerminatedErr:
		return "KernelTerminatedErr"
	case KernelTimeoutErr:
		return "KernelTimeoutErr"
	case KernelUnrecoverableErr:
		return "KernelUnrecoverableErr"
	case KernelUnsupportedErr:
		return "KernelUnsupportedErr"
	case KernelWritePermissionErr:
		return "KernelWritePermissionErr"
	default:
		return fmt.Sprintf("Kernel(%d)", e)
	}
}

type Key uint

const (
	// KeyAEAdjustMarksProc: Mark-adjusting function.
	KeyAEAdjustMarksProc Key = 'a'<<24 | 'd'<<16 | 'j'<<8 | 'm' // 'adjm'
	// KeyAECompareProc: Object-comparison function.
	KeyAECompareProc Key = 'c'<<24 | 'm'<<16 | 'p'<<8 | 'r' // 'cmpr'
	// KeyAECountProc: Object-counting function.
	KeyAECountProc Key = 'c'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'cont'
	// KeyAEGetErrDescProc: Get error descriptor callback function.
	KeyAEGetErrDescProc Key = 'i'<<24 | 'n'<<16 | 'd'<<8 | 'c' // 'indc'
	// KeyAEMarkProc: Object-marking function.
	KeyAEMarkProc Key = 'm'<<24 | 'a'<<16 | 'r'<<8 | 'k' // 'mark'
	// KeyAEMarkTokenProc: Mark token function.
	KeyAEMarkTokenProc Key = 'm'<<24 | 'k'<<16 | 'i'<<8 | 'd' // 'mkid'
	// KeyAERangeStart: Specifies the first Apple event object in a desired range.
	KeyAERangeStart Key = 's'<<24 | 't'<<16 | 'a'<<8 | 'r' // 'star'
	// KeyAERangeStop: Specifies the last Apple event object in the desired range.
	KeyAERangeStop Key = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
	// KeyAERecorderCount: Used with the `keyword` parameter of the AEManagerInfo(_:_:) function.
	KeyAERecorderCount Key = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'r' // 'recr'
	// KeyAEVersion: Used with the `keyword` parameter of the AEManagerInfo(_:_:) function.
	KeyAEVersion Key = 'v'<<24 | 'e'<<16 | 'r'<<8 | 's' // 'vers'
	KeyAcceptTimeoutAttr Key = 'a'<<24 | 'c'<<16 | 't'<<8 | 'm' // 'actm'
	KeyActualSenderAuditToken Key = 'a'<<24 | 'c'<<16 | 'a'<<8 | 't' // 'acat'
	// KeyAddressAttr: Address of a target or client application.
	KeyAddressAttr Key = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'addr'
	KeyCloseAllWindows Key = 'c'<<24 | 'a'<<16 | 'w'<<8 | ' ' // 'caw '
	// KeyDirectObject: Direct parameter.
	KeyDirectObject Key = '-'<<24 | '-'<<16 | '-'<<8 | '-' // '----'
	// KeyDisposeTokenProc: Token disposal function.
	KeyDisposeTokenProc Key = 'x'<<24 | 't'<<16 | 'o'<<8 | 'k' // 'xtok'
	KeyDriveNumber Key = 'd'<<24 | 'r'<<16 | 'v'<<8 | '#' // 'drv#'
	KeyErrorCode Key = 'e'<<24 | 'r'<<16 | 'r'<<8 | '#' // 'err#'
	// KeyErrorNumber: Error number.
	KeyErrorNumber Key = 'e'<<24 | 'r'<<16 | 'r'<<8 | 'n' // 'errn'
	// KeyErrorString: Error string.
	KeyErrorString Key = 'e'<<24 | 'r'<<16 | 'r'<<8 | 's' // 'errs'
	// KeyEventClassAttr: Event class of an Apple event.
	KeyEventClassAttr Key = 'e'<<24 | 'v'<<16 | 'c'<<8 | 'l' // 'evcl'
	// KeyEventIDAttr: Event ID of an Apple event.
	KeyEventIDAttr Key = 'e'<<24 | 'v'<<16 | 'i'<<8 | 'd' // 'evid'
	// KeyEventSourceAttr: Nature of the source application.
	KeyEventSourceAttr Key = 'e'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'esrc'
	KeyHighLevelClass Key = 'h'<<24 | 'c'<<16 | 'l'<<8 | 's' // 'hcls'
	KeyHighLevelID Key = 'h'<<24 | 'i'<<16 | 'd'<<8 | ' ' // 'hid '
	// KeyInteractLevelAttr: Settings for when to allow the Apple Event Manager to bring a server application to the foreground, if necessary, to interact with the user.
	KeyInteractLevelAttr Key = 'i'<<24 | 'n'<<16 | 't'<<8 | 'e' // 'inte'
	KeyKey Key = 'k'<<24 | 'e'<<16 | 'y'<<8 | ' ' // 'key '
	KeyKeyCode Key = 'c'<<24 | 'o'<<16 | 'd'<<8 | 'e' // 'code'
	KeyKeyboard Key = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'b' // 'keyb'
	KeyLocalWhere Key = 'l'<<24 | 'w'<<16 | 'h'<<8 | 'r' // 'lwhr'
	KeyMenuID Key = 'm'<<24 | 'i'<<16 | 'd'<<8 | ' ' // 'mid '
	KeyMenuItem Key = 'm'<<24 | 'i'<<16 | 't'<<8 | 'm' // 'mitm'
	KeyMiscellaneous Key = 'f'<<24 | 'm'<<16 | 's'<<8 | 'c' // 'fmsc'
	// KeyMissedKeywordAttr: # Discussion
	KeyMissedKeywordAttr Key = 'm'<<24 | 'i'<<16 | 's'<<8 | 's' // 'miss'
	KeyModifiers Key = 'm'<<24 | 'o'<<16 | 'd'<<8 | 's' // 'mods'
	KeyNewBounds Key = 'n'<<24 | 'b'<<16 | 'n'<<8 | 'd' // 'nbnd'
	// KeyOptionalKeywordAttr: List of keywords for parameters of an Apple event that should be treated as optional by the target application.
	KeyOptionalKeywordAttr Key = 'o'<<24 | 'p'<<16 | 't'<<8 | 'k' // 'optk'
	// KeyOriginalAddressAttr: Address of original source of Apple event if the event has been forwarded (available only in version 1.01 or later versions of the Apple Event Manager).
	KeyOriginalAddressAttr Key = 'f'<<24 | 'r'<<16 | 'o'<<8 | 'm' // 'from'
	KeyOriginalBounds Key = 'o'<<24 | 'b'<<16 | 'n'<<8 | 'd' // 'obnd'
	// KeyPreDispatch: A predispatch handler (an Apple event handler that the Apple Event Manager calls immediately before it dispatches an Apple event).
	KeyPreDispatch Key = 'p'<<24 | 'h'<<16 | 'a'<<8 | 'c' // 'phac'
	// KeyProcessSerialNumber: Process serial number.
	KeyProcessSerialNumber Key = 'p'<<24 | 's'<<16 | 'n'<<8 | ' ' // 'psn '
	// KeyReplyRequestedAttr: A Boolean value indicating whether the Apple event expects to be replied to.
	KeyReplyRequestedAttr Key = 'r'<<24 | 'e'<<16 | 'p'<<8 | 'q' // 'repq'
	// KeyReturnIDAttr: Return ID for a reply Apple event.
	KeyReturnIDAttr Key = 'r'<<24 | 't'<<16 | 'i'<<8 | 'd' // 'rtid'
	// KeySelectProc: You pass this value in the `functionClass` parameter of the AEManagerInfo(_:_:) function to disable the Object Support Library.
	KeySelectProc Key = 's'<<24 | 'e'<<16 | 'l'<<8 | 'h' // 'selh'
	KeySelection Key = 'f'<<24 | 's'<<16 | 'e'<<8 | 'l' // 'fsel'
	KeySenderApplescriptEntitlementsAttr Key = 'e'<<24 | 'n'<<16 | 't'<<8 | 'l' // 'entl'
	KeySenderApplicationIdentifierEntitlementAttr Key = 'a'<<24 | 'i'<<16 | 'e'<<8 | 'a' // 'aiea'
	KeySenderApplicationSandboxed Key = 's'<<24 | 's'<<16 | 's'<<8 | 'b' // 'sssb'
	KeySenderAuditTokenAttr Key = 't'<<24 | 'o'<<16 | 'k'<<8 | 'n' // 'tokn'
	KeySenderEGIDAttr Key = 's'<<24 | 'g'<<16 | 'i'<<8 | 'd' // 'sgid'
	KeySenderEUIDAttr Key = 's'<<24 | 'e'<<16 | 'i'<<8 | 'd' // 'seid'
	KeySenderGIDAttr Key = 'g'<<24 | 'i'<<16 | 'd'<<8 | 's' // 'gids'
	KeySenderPIDAttr Key = 's'<<24 | 'p'<<16 | 'i'<<8 | 'd' // 'spid'
	KeySenderUIDAttr Key = 'u'<<24 | 'i'<<16 | 'd'<<8 | 's' // 'uids'
	// KeyTimeoutAttr: Length of time, in ticks, that the client will wait for a reply or a result from the server.
	KeyTimeoutAttr Key = 't'<<24 | 'i'<<16 | 'm'<<8 | 'o' // 'timo'
	// KeyTransactionIDAttr: Transaction ID identifying a series of Apple events that are part of one transaction.
	KeyTransactionIDAttr Key = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
	KeyWhen Key = 'w'<<24 | 'h'<<16 | 'e'<<8 | 'n' // 'when'
	KeyWhere Key = 'w'<<24 | 'h'<<16 | 'e'<<8 | 'r' // 'wher'
	KeyWindow Key = 'k'<<24 | 'w'<<16 | 'n'<<8 | 'd' // 'kwnd'
)

func (e Key) String() string {
	switch e {
	case KeyAEAdjustMarksProc:
		return "KeyAEAdjustMarksProc"
	case KeyAECompareProc:
		return "KeyAECompareProc"
	case KeyAECountProc:
		return "KeyAECountProc"
	case KeyAEGetErrDescProc:
		return "KeyAEGetErrDescProc"
	case KeyAEMarkProc:
		return "KeyAEMarkProc"
	case KeyAEMarkTokenProc:
		return "KeyAEMarkTokenProc"
	case KeyAERangeStart:
		return "KeyAERangeStart"
	case KeyAERangeStop:
		return "KeyAERangeStop"
	case KeyAERecorderCount:
		return "KeyAERecorderCount"
	case KeyAEVersion:
		return "KeyAEVersion"
	case KeyAcceptTimeoutAttr:
		return "KeyAcceptTimeoutAttr"
	case KeyActualSenderAuditToken:
		return "KeyActualSenderAuditToken"
	case KeyAddressAttr:
		return "KeyAddressAttr"
	case KeyCloseAllWindows:
		return "KeyCloseAllWindows"
	case KeyDirectObject:
		return "KeyDirectObject"
	case KeyDisposeTokenProc:
		return "KeyDisposeTokenProc"
	case KeyDriveNumber:
		return "KeyDriveNumber"
	case KeyErrorCode:
		return "KeyErrorCode"
	case KeyErrorNumber:
		return "KeyErrorNumber"
	case KeyErrorString:
		return "KeyErrorString"
	case KeyEventClassAttr:
		return "KeyEventClassAttr"
	case KeyEventIDAttr:
		return "KeyEventIDAttr"
	case KeyEventSourceAttr:
		return "KeyEventSourceAttr"
	case KeyHighLevelClass:
		return "KeyHighLevelClass"
	case KeyHighLevelID:
		return "KeyHighLevelID"
	case KeyInteractLevelAttr:
		return "KeyInteractLevelAttr"
	case KeyKey:
		return "KeyKey"
	case KeyKeyCode:
		return "KeyKeyCode"
	case KeyKeyboard:
		return "KeyKeyboard"
	case KeyLocalWhere:
		return "KeyLocalWhere"
	case KeyMenuID:
		return "KeyMenuID"
	case KeyMenuItem:
		return "KeyMenuItem"
	case KeyMiscellaneous:
		return "KeyMiscellaneous"
	case KeyMissedKeywordAttr:
		return "KeyMissedKeywordAttr"
	case KeyModifiers:
		return "KeyModifiers"
	case KeyNewBounds:
		return "KeyNewBounds"
	case KeyOptionalKeywordAttr:
		return "KeyOptionalKeywordAttr"
	case KeyOriginalAddressAttr:
		return "KeyOriginalAddressAttr"
	case KeyOriginalBounds:
		return "KeyOriginalBounds"
	case KeyPreDispatch:
		return "KeyPreDispatch"
	case KeyProcessSerialNumber:
		return "KeyProcessSerialNumber"
	case KeyReplyRequestedAttr:
		return "KeyReplyRequestedAttr"
	case KeyReturnIDAttr:
		return "KeyReturnIDAttr"
	case KeySelectProc:
		return "KeySelectProc"
	case KeySelection:
		return "KeySelection"
	case KeySenderApplescriptEntitlementsAttr:
		return "KeySenderApplescriptEntitlementsAttr"
	case KeySenderApplicationIdentifierEntitlementAttr:
		return "KeySenderApplicationIdentifierEntitlementAttr"
	case KeySenderApplicationSandboxed:
		return "KeySenderApplicationSandboxed"
	case KeySenderAuditTokenAttr:
		return "KeySenderAuditTokenAttr"
	case KeySenderEGIDAttr:
		return "KeySenderEGIDAttr"
	case KeySenderEUIDAttr:
		return "KeySenderEUIDAttr"
	case KeySenderGIDAttr:
		return "KeySenderGIDAttr"
	case KeySenderPIDAttr:
		return "KeySenderPIDAttr"
	case KeySenderUIDAttr:
		return "KeySenderUIDAttr"
	case KeyTimeoutAttr:
		return "KeyTimeoutAttr"
	case KeyTransactionIDAttr:
		return "KeyTransactionIDAttr"
	case KeyWhen:
		return "KeyWhen"
	case KeyWhere:
		return "KeyWhere"
	case KeyWindow:
		return "KeyWindow"
	default:
		return fmt.Sprintf("Key(%d)", e)
	}
}

type KeyAE uint

const (
	KeyAEAngle KeyAE = 'k'<<24 | 'a'<<16 | 'n'<<8 | 'g' // 'kang'
	KeyAEArcAngle KeyAE = 'p'<<24 | 'a'<<16 | 'r'<<8 | 'c' // 'parc'
	KeyAEBaseAddr KeyAE = 'b'<<24 | 'a'<<16 | 'd'<<8 | 'd' // 'badd'
	KeyAEBestType KeyAE = 'p'<<24 | 'b'<<16 | 's'<<8 | 't' // 'pbst'
	KeyAEBgndColor KeyAE = 'k'<<24 | 'b'<<16 | 'c'<<8 | 'l' // 'kbcl'
	KeyAEBgndPattern KeyAE = 'k'<<24 | 'b'<<16 | 'p'<<8 | 't' // 'kbpt'
	KeyAEBounds KeyAE = 'p'<<24 | 'b'<<16 | 'n'<<8 | 'd' // 'pbnd'
	KeyAECellList KeyAE = 'k'<<24 | 'c'<<16 | 'l'<<8 | 't' // 'kclt'
	KeyAEClassID KeyAE = 'c'<<24 | 'l'<<16 | 'I'<<8 | 'D' // 'clID'
	KeyAEClauseOffsets KeyAE = 'c'<<24 | 'l'<<16 | 'a'<<8 | 'u' // 'clau'
	KeyAEColor KeyAE = 'c'<<24 | 'o'<<16 | 'l'<<8 | 'r' // 'colr'
	KeyAEColorTable KeyAE = 'c'<<24 | 'l'<<16 | 't'<<8 | 'b' // 'cltb'
	KeyAECurveHeight KeyAE = 'k'<<24 | 'c'<<16 | 'h'<<8 | 'd' // 'kchd'
	KeyAECurveWidth KeyAE = 'k'<<24 | 'c'<<16 | 'w'<<8 | 'd' // 'kcwd'
	KeyAEDashStyle KeyAE = 'p'<<24 | 'd'<<16 | 's'<<8 | 't' // 'pdst'
	KeyAEData KeyAE = 'd'<<24 | 'a'<<16 | 't'<<8 | 'a' // 'data'
	KeyAEDefaultType KeyAE = 'd'<<24 | 'e'<<16 | 'f'<<8 | 't' // 'deft'
	KeyAEDefinitionRect KeyAE = 'p'<<24 | 'd'<<16 | 'r'<<8 | 't' // 'pdrt'
	KeyAEDescType KeyAE = 'd'<<24 | 's'<<16 | 't'<<8 | 'p' // 'dstp'
	KeyAEDestination KeyAE = 'd'<<24 | 'e'<<16 | 's'<<8 | 't' // 'dest'
	KeyAEDoAntiAlias KeyAE = 'a'<<24 | 'n'<<16 | 't'<<8 | 'a' // 'anta'
	KeyAEDoDithered KeyAE = 'g'<<24 | 'd'<<16 | 'i'<<8 | 't' // 'gdit'
	KeyAEDoRotate KeyAE = 'k'<<24 | 'd'<<16 | 'r'<<8 | 't' // 'kdrt'
	KeyAEDoScale KeyAE = 'k'<<24 | 's'<<16 | 'c'<<8 | 'a' // 'ksca'
	KeyAEDoTranslate KeyAE = 'k'<<24 | 't'<<16 | 'r'<<8 | 'a' // 'ktra'
	KeyAEDragging KeyAE = 'b'<<24 | 'o'<<16 | 'o'<<8 | 'l' // 'bool'
	KeyAEEditionFileLoc KeyAE = 'e'<<24 | 'l'<<16 | 'o'<<8 | 'c' // 'eloc'
	KeyAEElements KeyAE = 'e'<<24 | 'l'<<16 | 'm'<<8 | 's' // 'elms'
	KeyAEEndPoint KeyAE = 'p'<<24 | 'e'<<16 | 'n'<<8 | 'd' // 'pend'
	KeyAEEventClass KeyAE = 'e'<<24 | 'v'<<16 | 'c'<<8 | 'l' // 'evcl'
	KeyAEEventID KeyAE = 'e'<<24 | 'v'<<16 | 't'<<8 | 'i' // 'evti'
	KeyAEFile KeyAE = 'k'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'kfil'
	KeyAEFileType KeyAE = 'f'<<24 | 'l'<<16 | 't'<<8 | 'p' // 'fltp'
	KeyAEFillColor KeyAE = 'f'<<24 | 'l'<<16 | 'c'<<8 | 'l' // 'flcl'
	KeyAEFillPattern KeyAE = 'f'<<24 | 'l'<<16 | 'p'<<8 | 't' // 'flpt'
	KeyAEFlipHorizontal KeyAE = 'k'<<24 | 'f'<<16 | 'h'<<8 | 'o' // 'kfho'
	KeyAEFlipVertical KeyAE = 'k'<<24 | 'f'<<16 | 'v'<<8 | 't' // 'kfvt'
	KeyAEFont KeyAE = 'f'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'font'
	KeyAEFormula KeyAE = 'p'<<24 | 'f'<<16 | 'o'<<8 | 'r' // 'pfor'
	KeyAEGraphicObjects KeyAE = 'g'<<24 | 'o'<<16 | 'b'<<8 | 's' // 'gobs'
	KeyAEHiliteRange KeyAE = 'h'<<24 | 'r'<<16 | 'n'<<8 | 'g' // 'hrng'
	KeyAEID KeyAE = 'I'<<24 | 'D'<<16 | ' '<<8 | ' ' // 'ID  '
	KeyAEImageQuality KeyAE = 'g'<<24 | 'q'<<16 | 'u'<<8 | 'a' // 'gqua'
	KeyAEInsertHere KeyAE = 'i'<<24 | 'n'<<16 | 's'<<8 | 'h' // 'insh'
	KeyAEKeyForms KeyAE = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'f' // 'keyf'
	KeyAEKeyword KeyAE = 'k'<<24 | 'y'<<16 | 'w'<<8 | 'd' // 'kywd'
	KeyAELeftSide KeyAE = 'k'<<24 | 'l'<<16 | 'e'<<8 | 'f' // 'klef'
	KeyAELevel KeyAE = 'l'<<24 | 'e'<<16 | 'v'<<8 | 'l' // 'levl'
	KeyAELineArrow KeyAE = 'a'<<24 | 'r'<<16 | 'r'<<8 | 'o' // 'arro'
	KeyAEName KeyAE = 'p'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'pnam'
	KeyAENewElementLoc KeyAE = 'p'<<24 | 'n'<<16 | 'e'<<8 | 'l' // 'pnel'
	KeyAEObject KeyAE = 'k'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'kobj'
	KeyAEObjectClass KeyAE = 'k'<<24 | 'o'<<16 | 'c'<<8 | 'l' // 'kocl'
	KeyAEOffStyles KeyAE = 'o'<<24 | 'f'<<16 | 's'<<8 | 't' // 'ofst'
	KeyAEOffset KeyAE = 'o'<<24 | 'f'<<16 | 's'<<8 | 't' // 'ofst'
	KeyAEOnStyles KeyAE = 'o'<<24 | 'n'<<16 | 's'<<8 | 't' // 'onst'
	KeyAEPMTable KeyAE = 'k'<<24 | 'p'<<16 | 'm'<<8 | 't' // 'kpmt'
	KeyAEParamFlags KeyAE = 'p'<<24 | 'm'<<16 | 'f'<<8 | 'g' // 'pmfg'
	KeyAEParameters KeyAE = 'p'<<24 | 'r'<<16 | 'm'<<8 | 's' // 'prms'
	KeyAEPenColor KeyAE = 'p'<<24 | 'p'<<16 | 'c'<<8 | 'l' // 'ppcl'
	KeyAEPenPattern KeyAE = 'p'<<24 | 'p'<<16 | 'p'<<8 | 'a' // 'pppa'
	KeyAEPenWidth KeyAE = 'p'<<24 | 'p'<<16 | 'w'<<8 | 'd' // 'ppwd'
	KeyAEPinRange KeyAE = 'p'<<24 | 'n'<<16 | 'r'<<8 | 'g' // 'pnrg'
	KeyAEPixMapMinus KeyAE = 'k'<<24 | 'p'<<16 | 'm'<<8 | 'm' // 'kpmm'
	KeyAEPixelDepth KeyAE = 'p'<<24 | 'd'<<16 | 'p'<<8 | 't' // 'pdpt'
	KeyAEPoint KeyAE = 'g'<<24 | 'p'<<16 | 'o'<<8 | 's' // 'gpos'
	KeyAEPointList KeyAE = 'p'<<24 | 't'<<16 | 'l'<<8 | 't' // 'ptlt'
	KeyAEPointSize KeyAE = 'p'<<24 | 't'<<16 | 's'<<8 | 'z' // 'ptsz'
	KeyAEPosition KeyAE = 'k'<<24 | 'p'<<16 | 'o'<<8 | 's' // 'kpos'
	KeyAEPropData KeyAE = 'p'<<24 | 'r'<<16 | 'd'<<8 | 't' // 'prdt'
	KeyAEPropFlags KeyAE = 'p'<<24 | 'r'<<16 | 'f'<<8 | 'g' // 'prfg'
	KeyAEPropID KeyAE = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
	KeyAEProperties KeyAE = 'q'<<24 | 'p'<<16 | 'r'<<8 | 'o' // 'qpro'
	KeyAEProperty KeyAE = 'k'<<24 | 'p'<<16 | 'r'<<8 | 'p' // 'kprp'
	KeyAEProtection KeyAE = 'p'<<24 | 'p'<<16 | 'r'<<8 | 'o' // 'ppro'
	KeyAERegionClass KeyAE = 'r'<<24 | 'g'<<16 | 'n'<<8 | 'c' // 'rgnc'
	KeyAERenderAs KeyAE = 'k'<<24 | 'r'<<16 | 'e'<<8 | 'n' // 'kren'
	KeyAERequestedType KeyAE = 'r'<<24 | 't'<<16 | 'y'<<8 | 'p' // 'rtyp'
	KeyAEResult KeyAE = '-'<<24 | '-'<<16 | '-'<<8 | '-' // '----'
	KeyAEResultInfo KeyAE = 'r'<<24 | 's'<<16 | 'i'<<8 | 'n' // 'rsin'
	KeyAERotPoint KeyAE = 'k'<<24 | 'r'<<16 | 't'<<8 | 'p' // 'krtp'
	KeyAERotation KeyAE = 'p'<<24 | 'r'<<16 | 'o'<<8 | 't' // 'prot'
	KeyAERowList KeyAE = 'k'<<24 | 'r'<<16 | 'l'<<8 | 's' // 'krls'
	KeyAESaveOptions KeyAE = 's'<<24 | 'a'<<16 | 'v'<<8 | 'o' // 'savo'
	KeyAEScale KeyAE = 'p'<<24 | 's'<<16 | 'c'<<8 | 'l' // 'pscl'
	KeyAEScriptTag KeyAE = 'p'<<24 | 's'<<16 | 'c'<<8 | 't' // 'psct'
	// KeyAESearchText: # Discussion
	KeyAESearchText KeyAE = 's'<<24 | 't'<<16 | 'x'<<8 | 't' // 'stxt'
	KeyAEShowWhere KeyAE = 's'<<24 | 'h'<<16 | 'o'<<8 | 'w' // 'show'
	KeyAEStartAngle KeyAE = 'p'<<24 | 'a'<<16 | 'n'<<8 | 'g' // 'pang'
	KeyAEStartPoint KeyAE = 'p'<<24 | 's'<<16 | 't'<<8 | 'p' // 'pstp'
	KeyAEStyles KeyAE = 'k'<<24 | 's'<<16 | 't'<<8 | 'y' // 'ksty'
	KeyAESuiteID KeyAE = 's'<<24 | 'u'<<16 | 'i'<<8 | 't' // 'suit'
	KeyAEText KeyAE = 'k'<<24 | 't'<<16 | 'x'<<8 | 't' // 'ktxt'
	KeyAETextColor KeyAE = 'p'<<24 | 't'<<16 | 'x'<<8 | 'c' // 'ptxc'
	KeyAETextFont KeyAE = 'p'<<24 | 't'<<16 | 'x'<<8 | 'f' // 'ptxf'
	KeyAETextLineAscent KeyAE = 'k'<<24 | 't'<<16 | 'a'<<8 | 's' // 'ktas'
	KeyAETextLineHeight KeyAE = 'k'<<24 | 't'<<16 | 'l'<<8 | 'h' // 'ktlh'
	KeyAETextPointSize KeyAE = 'p'<<24 | 't'<<16 | 'p'<<8 | 's' // 'ptps'
	KeyAETextStyles KeyAE = 't'<<24 | 'x'<<16 | 's'<<8 | 't' // 'txst'
	KeyAETheText KeyAE = 't'<<24 | 'h'<<16 | 't'<<8 | 'x' // 'thtx'
	KeyAETransferMode KeyAE = 'p'<<24 | 'p'<<16 | 't'<<8 | 'm' // 'pptm'
	KeyAETranslation KeyAE = 'p'<<24 | 't'<<16 | 'r'<<8 | 's' // 'ptrs'
	KeyAETryAsStructGraf KeyAE = 't'<<24 | 'o'<<16 | 'o'<<8 | 'g' // 'toog'
	KeyAEUniformStyles KeyAE = 'u'<<24 | 's'<<16 | 't'<<8 | 'l' // 'ustl'
	KeyAEUpdateOn KeyAE = 'p'<<24 | 'u'<<16 | 'p'<<8 | 'd' // 'pupd'
	KeyAEUserTerm KeyAE = 'u'<<24 | 't'<<16 | 'r'<<8 | 'm' // 'utrm'
	KeyAEWindow KeyAE = 'w'<<24 | 'n'<<16 | 'd'<<8 | 'w' // 'wndw'
	KeyAEWritingCode KeyAE = 'w'<<24 | 'r'<<16 | 'c'<<8 | 'd' // 'wrcd'
)

func (e KeyAE) String() string {
	switch e {
	case KeyAEAngle:
		return "KeyAEAngle"
	case KeyAEArcAngle:
		return "KeyAEArcAngle"
	case KeyAEBaseAddr:
		return "KeyAEBaseAddr"
	case KeyAEBestType:
		return "KeyAEBestType"
	case KeyAEBgndColor:
		return "KeyAEBgndColor"
	case KeyAEBgndPattern:
		return "KeyAEBgndPattern"
	case KeyAEBounds:
		return "KeyAEBounds"
	case KeyAECellList:
		return "KeyAECellList"
	case KeyAEClassID:
		return "KeyAEClassID"
	case KeyAEClauseOffsets:
		return "KeyAEClauseOffsets"
	case KeyAEColor:
		return "KeyAEColor"
	case KeyAEColorTable:
		return "KeyAEColorTable"
	case KeyAECurveHeight:
		return "KeyAECurveHeight"
	case KeyAECurveWidth:
		return "KeyAECurveWidth"
	case KeyAEDashStyle:
		return "KeyAEDashStyle"
	case KeyAEData:
		return "KeyAEData"
	case KeyAEDefaultType:
		return "KeyAEDefaultType"
	case KeyAEDefinitionRect:
		return "KeyAEDefinitionRect"
	case KeyAEDescType:
		return "KeyAEDescType"
	case KeyAEDestination:
		return "KeyAEDestination"
	case KeyAEDoAntiAlias:
		return "KeyAEDoAntiAlias"
	case KeyAEDoDithered:
		return "KeyAEDoDithered"
	case KeyAEDoRotate:
		return "KeyAEDoRotate"
	case KeyAEDoScale:
		return "KeyAEDoScale"
	case KeyAEDoTranslate:
		return "KeyAEDoTranslate"
	case KeyAEDragging:
		return "KeyAEDragging"
	case KeyAEEditionFileLoc:
		return "KeyAEEditionFileLoc"
	case KeyAEElements:
		return "KeyAEElements"
	case KeyAEEndPoint:
		return "KeyAEEndPoint"
	case KeyAEEventClass:
		return "KeyAEEventClass"
	case KeyAEEventID:
		return "KeyAEEventID"
	case KeyAEFile:
		return "KeyAEFile"
	case KeyAEFileType:
		return "KeyAEFileType"
	case KeyAEFillColor:
		return "KeyAEFillColor"
	case KeyAEFillPattern:
		return "KeyAEFillPattern"
	case KeyAEFlipHorizontal:
		return "KeyAEFlipHorizontal"
	case KeyAEFlipVertical:
		return "KeyAEFlipVertical"
	case KeyAEFont:
		return "KeyAEFont"
	case KeyAEFormula:
		return "KeyAEFormula"
	case KeyAEGraphicObjects:
		return "KeyAEGraphicObjects"
	case KeyAEHiliteRange:
		return "KeyAEHiliteRange"
	case KeyAEID:
		return "KeyAEID"
	case KeyAEImageQuality:
		return "KeyAEImageQuality"
	case KeyAEInsertHere:
		return "KeyAEInsertHere"
	case KeyAEKeyForms:
		return "KeyAEKeyForms"
	case KeyAEKeyword:
		return "KeyAEKeyword"
	case KeyAELeftSide:
		return "KeyAELeftSide"
	case KeyAELevel:
		return "KeyAELevel"
	case KeyAELineArrow:
		return "KeyAELineArrow"
	case KeyAEName:
		return "KeyAEName"
	case KeyAENewElementLoc:
		return "KeyAENewElementLoc"
	case KeyAEObject:
		return "KeyAEObject"
	case KeyAEObjectClass:
		return "KeyAEObjectClass"
	case KeyAEOffStyles:
		return "KeyAEOffStyles"
	case KeyAEOnStyles:
		return "KeyAEOnStyles"
	case KeyAEPMTable:
		return "KeyAEPMTable"
	case KeyAEParamFlags:
		return "KeyAEParamFlags"
	case KeyAEParameters:
		return "KeyAEParameters"
	case KeyAEPenColor:
		return "KeyAEPenColor"
	case KeyAEPenPattern:
		return "KeyAEPenPattern"
	case KeyAEPenWidth:
		return "KeyAEPenWidth"
	case KeyAEPinRange:
		return "KeyAEPinRange"
	case KeyAEPixMapMinus:
		return "KeyAEPixMapMinus"
	case KeyAEPixelDepth:
		return "KeyAEPixelDepth"
	case KeyAEPoint:
		return "KeyAEPoint"
	case KeyAEPointList:
		return "KeyAEPointList"
	case KeyAEPointSize:
		return "KeyAEPointSize"
	case KeyAEPosition:
		return "KeyAEPosition"
	case KeyAEPropData:
		return "KeyAEPropData"
	case KeyAEPropFlags:
		return "KeyAEPropFlags"
	case KeyAEPropID:
		return "KeyAEPropID"
	case KeyAEProperties:
		return "KeyAEProperties"
	case KeyAEProperty:
		return "KeyAEProperty"
	case KeyAEProtection:
		return "KeyAEProtection"
	case KeyAERegionClass:
		return "KeyAERegionClass"
	case KeyAERenderAs:
		return "KeyAERenderAs"
	case KeyAERequestedType:
		return "KeyAERequestedType"
	case KeyAEResult:
		return "KeyAEResult"
	case KeyAEResultInfo:
		return "KeyAEResultInfo"
	case KeyAERotPoint:
		return "KeyAERotPoint"
	case KeyAERotation:
		return "KeyAERotation"
	case KeyAERowList:
		return "KeyAERowList"
	case KeyAESaveOptions:
		return "KeyAESaveOptions"
	case KeyAEScale:
		return "KeyAEScale"
	case KeyAEScriptTag:
		return "KeyAEScriptTag"
	case KeyAESearchText:
		return "KeyAESearchText"
	case KeyAEShowWhere:
		return "KeyAEShowWhere"
	case KeyAEStartAngle:
		return "KeyAEStartAngle"
	case KeyAEStartPoint:
		return "KeyAEStartPoint"
	case KeyAEStyles:
		return "KeyAEStyles"
	case KeyAESuiteID:
		return "KeyAESuiteID"
	case KeyAEText:
		return "KeyAEText"
	case KeyAETextColor:
		return "KeyAETextColor"
	case KeyAETextFont:
		return "KeyAETextFont"
	case KeyAETextLineAscent:
		return "KeyAETextLineAscent"
	case KeyAETextLineHeight:
		return "KeyAETextLineHeight"
	case KeyAETextPointSize:
		return "KeyAETextPointSize"
	case KeyAETextStyles:
		return "KeyAETextStyles"
	case KeyAETheText:
		return "KeyAETheText"
	case KeyAETransferMode:
		return "KeyAETransferMode"
	case KeyAETranslation:
		return "KeyAETranslation"
	case KeyAETryAsStructGraf:
		return "KeyAETryAsStructGraf"
	case KeyAEUniformStyles:
		return "KeyAEUniformStyles"
	case KeyAEUpdateOn:
		return "KeyAEUpdateOn"
	case KeyAEUserTerm:
		return "KeyAEUserTerm"
	case KeyAEWindow:
		return "KeyAEWindow"
	case KeyAEWritingCode:
		return "KeyAEWritingCode"
	default:
		return fmt.Sprintf("KeyAE(%d)", e)
	}
}

type KeyAELaunchedAs uint

const (
	// KeyAELaunchedAsLogInItem: If present in a `kAEOpenApplication` event, the receiving application was launched as a login item and should only perform actions suitable to that environment—for example, it probably shouldn't open an untitled document.
	KeyAELaunchedAsLogInItem KeyAELaunchedAs = 'l'<<24 | 'g'<<16 | 'i'<<8 | 't' // 'lgit'
	// KeyAELaunchedAsServiceItem: # Discussion
	KeyAELaunchedAsServiceItem KeyAELaunchedAs = 's'<<24 | 'v'<<16 | 'i'<<8 | 't' // 'svit'
)

func (e KeyAELaunchedAs) String() string {
	switch e {
	case KeyAELaunchedAsLogInItem:
		return "KeyAELaunchedAsLogInItem"
	case KeyAELaunchedAsServiceItem:
		return "KeyAELaunchedAsServiceItem"
	default:
		return fmt.Sprintf("KeyAELaunchedAs(%d)", e)
	}
}

const KeyReplyPortAttr uint = 'r'<<24 | 'e'<<16 | 'p'<<8 | 'p' // 'repp'

type KeySOAP uint

const (
	KeySOAPSMDNamespace KeySOAP = 's'<<24 | 's'<<16 | 'n'<<8 | 's' // 'ssns'
	KeySOAPSMDNamespaceURI KeySOAP = 's'<<24 | 's'<<16 | 'n'<<8 | 'u' // 'ssnu'
	KeySOAPSMDType KeySOAP = 's'<<24 | 's'<<16 | 't'<<8 | 'p' // 'sstp'
	KeySOAPStructureMetaData KeySOAP = '/'<<24 | 's'<<16 | 'm'<<8 | 'd' // '/smd'
)

func (e KeySOAP) String() string {
	switch e {
	case KeySOAPSMDNamespace:
		return "KeySOAPSMDNamespace"
	case KeySOAPSMDNamespaceURI:
		return "KeySOAPSMDNamespaceURI"
	case KeySOAPSMDType:
		return "KeySOAPSMDType"
	case KeySOAPStructureMetaData:
		return "KeySOAPStructureMetaData"
	default:
		return fmt.Sprintf("KeySOAP(%d)", e)
	}
}

type KeyUserNameAttr uint

const (
	KAERPCClass KeyUserNameAttr = 'r'<<24 | 'p'<<16 | 'c'<<8 | ' ' // 'rpc '
	KAESOAPScheme KeyUserNameAttr = 'S'<<24 | 'O'<<16 | 'A'<<8 | 'P' // 'SOAP'
	KAESharedScriptHandler KeyUserNameAttr = 'w'<<24 | 's'<<16 | 'c'<<8 | 'p' // 'wscp'
	KAEXMLRPCScheme KeyUserNameAttr = 'R'<<24 | 'P'<<16 | 'C'<<8 | '2' // 'RPC2'
	KeyAEPOSTHeaderData KeyUserNameAttr = 'p'<<24 | 'h'<<16 | 'e'<<8 | 'd' // 'phed'
	KeyAEReplyHeaderData KeyUserNameAttr = 'r'<<24 | 'h'<<16 | 'e'<<8 | 'd' // 'rhed'
	KeyAEXMLReplyData KeyUserNameAttr = 'x'<<24 | 'r'<<16 | 'e'<<8 | 'p' // 'xrep'
	KeyAEXMLRequestData KeyUserNameAttr = 'x'<<24 | 'r'<<16 | 'e'<<8 | 'q' // 'xreq'
	KeyAdditionalHTTPHeaders KeyUserNameAttr = 'a'<<24 | 'h'<<16 | 'e'<<8 | 'd' // 'ahed'
	KeyDisableAuthenticationAttr KeyUserNameAttr = 'a'<<24 | 'u'<<16 | 't'<<8 | 'h' // 'auth'
	KeyRPCMethodName KeyUserNameAttr = 'm'<<24 | 'e'<<16 | 't'<<8 | 'h' // 'meth'
	KeyRPCMethodParam KeyUserNameAttr = 'p'<<24 | 'a'<<16 | 'r'<<8 | 'm' // 'parm'
	KeyRPCMethodParamOrder KeyUserNameAttr = '/'<<24 | 'o'<<16 | 'r'<<8 | 'd' // '/ord'
	KeySOAPAction KeyUserNameAttr = 's'<<24 | 'a'<<16 | 'c'<<8 | 't' // 'sact'
	KeySOAPMethodNameSpace KeyUserNameAttr = 'm'<<24 | 's'<<16 | 'p'<<8 | 'c' // 'mspc'
	KeySOAPMethodNameSpaceURI KeyUserNameAttr = 'm'<<24 | 's'<<16 | 'p'<<8 | 'u' // 'mspu'
	KeySOAPSchemaVersion KeyUserNameAttr = 's'<<24 | 's'<<16 | 'c'<<8 | 'h' // 'ssch'
	KeyUserNameAttrValue KeyUserNameAttr = 'u'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'unam'
	KeyUserPasswordAttr KeyUserNameAttr = 'p'<<24 | 'a'<<16 | 's'<<8 | 's' // 'pass'
	KeyXMLDebuggingAttr KeyUserNameAttr = 'x'<<24 | 'd'<<16 | 'b'<<8 | 'g' // 'xdbg'
)

func (e KeyUserNameAttr) String() string {
	switch e {
	case KAERPCClass:
		return "KAERPCClass"
	case KAESOAPScheme:
		return "KAESOAPScheme"
	case KAESharedScriptHandler:
		return "KAESharedScriptHandler"
	case KAEXMLRPCScheme:
		return "KAEXMLRPCScheme"
	case KeyAEPOSTHeaderData:
		return "KeyAEPOSTHeaderData"
	case KeyAEReplyHeaderData:
		return "KeyAEReplyHeaderData"
	case KeyAEXMLReplyData:
		return "KeyAEXMLReplyData"
	case KeyAEXMLRequestData:
		return "KeyAEXMLRequestData"
	case KeyAdditionalHTTPHeaders:
		return "KeyAdditionalHTTPHeaders"
	case KeyDisableAuthenticationAttr:
		return "KeyDisableAuthenticationAttr"
	case KeyRPCMethodName:
		return "KeyRPCMethodName"
	case KeyRPCMethodParam:
		return "KeyRPCMethodParam"
	case KeyRPCMethodParamOrder:
		return "KeyRPCMethodParamOrder"
	case KeySOAPAction:
		return "KeySOAPAction"
	case KeySOAPMethodNameSpace:
		return "KeySOAPMethodNameSpace"
	case KeySOAPMethodNameSpaceURI:
		return "KeySOAPMethodNameSpaceURI"
	case KeySOAPSchemaVersion:
		return "KeySOAPSchemaVersion"
	case KeyUserNameAttrValue:
		return "KeyUserNameAttrValue"
	case KeyUserPasswordAttr:
		return "KeyUserPasswordAttr"
	case KeyXMLDebuggingAttr:
		return "KeyXMLDebuggingAttr"
	default:
		return fmt.Sprintf("KeyUserNameAttr(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/lsacceptanceflags
type LSAcceptanceFlags uint32

const (
	// AcceptAllowLoginUI: Requests that the user interface to log in be presented.
	AcceptAllowLoginUI LSAcceptanceFlags = 0
	// AcceptDefault: Requests the default behavior that does not require the user interface to log in be presented.
	AcceptDefault LSAcceptanceFlags = 0
)

func (e LSAcceptanceFlags) String() string {
	switch e {
	case AcceptAllowLoginUI:
		return "AcceptAllowLoginUI"
	default:
		return fmt.Sprintf("LSAcceptanceFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/lshandleroptions
type LSHandlerOptions uint32

const (
	// Deprecated.
	IgnoreCreator LSHandlerOptions = 0
)

func (e LSHandlerOptions) String() string {
	switch e {
	case IgnoreCreator:
		return "IgnoreCreator"
	default:
		return fmt.Sprintf("LSHandlerOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/lsiteminfoflags
type LSItemInfoFlags uint32

const (
	// Deprecated.
	AppIsScriptable LSItemInfoFlags = 0
	// Deprecated.
	AppPrefersClassic LSItemInfoFlags = 0
	// Deprecated.
	AppPrefersNative LSItemInfoFlags = 0
	// Deprecated.
	ExtensionIsHidden LSItemInfoFlags = 0
	// Deprecated.
	IsAliasFile LSItemInfoFlags = 0
	// Deprecated.
	IsApplication LSItemInfoFlags = 0
	// Deprecated.
	IsClassicApp LSItemInfoFlags = 0
	// Deprecated.
	IsContainer LSItemInfoFlags = 0
	// Deprecated.
	IsInvisible LSItemInfoFlags = 0
	// Deprecated.
	IsNativeApp LSItemInfoFlags = 0
	// Deprecated.
	IsPackage LSItemInfoFlags = 0
	// Deprecated.
	IsPlainFile LSItemInfoFlags = 0
	// Deprecated.
	IsSymlink LSItemInfoFlags = 0
	// Deprecated.
	IsVolume LSItemInfoFlags = 0
)

func (e LSItemInfoFlags) String() string {
	switch e {
	case AppIsScriptable:
		return "AppIsScriptable"
	default:
		return fmt.Sprintf("LSItemInfoFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/lslaunchflags
type LSLaunchFlags uint32

const (
	// AndDisplayErrors: Requests that launch and open failures be displayed in the UI.
	AndDisplayErrors LSLaunchFlags = 0
	// AndHide: Requests that the application be hidden as soon as it completes its launch sequence.
	AndHide LSLaunchFlags = 0
	// AndHideOthers: Requests that other applications be hidden as soon as the opened application completes its launch sequence.
	AndHideOthers LSLaunchFlags = 0
	// AndPrint: Requests that documents opened in the application be printed.
	AndPrint LSLaunchFlags = 0
	// Async: Requests that the application be launched asynchronously.
	Async LSLaunchFlags = 0
	// Defaults: Requests launching in the default manner (as if the only flags set were `kLSLaunchNoParams`, `kLSLaunchAsync`, and `kLSLaunchStartClassic`).
	Defaults LSLaunchFlags = 0
	// DontAddToRecents: Requests that the application or documents not be added to the Finder’s Recent Items menu.
	DontAddToRecents LSLaunchFlags = 0
	// DontSwitch: Requests that the application be launched without being brought to the foreground.
	DontSwitch LSLaunchFlags = 0
	// NewInstance: Requests that a new instance of the application be started, even if one is already running.
	NewInstance LSLaunchFlags = 0
)

func (e LSLaunchFlags) String() string {
	switch e {
	case AndDisplayErrors:
		return "AndDisplayErrors"
	default:
		return fmt.Sprintf("LSLaunchFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/lsrequestedinfo
type LSRequestedInfo uint32

const (
	// Deprecated.
	RequestAllFlags LSRequestedInfo = 0
	// Deprecated.
	RequestAllInfo LSRequestedInfo = 0
	// Deprecated.
	RequestAppTypeFlags LSRequestedInfo = 0
	// Deprecated.
	RequestBasicFlagsOnly LSRequestedInfo = 0
	// Deprecated.
	RequestExtension LSRequestedInfo = 0
	// Deprecated.
	RequestExtensionFlagsOnly LSRequestedInfo = 0
	// Deprecated.
	RequestIconAndKind LSRequestedInfo = 0
	// Deprecated.
	RequestTypeCreator LSRequestedInfo = 0
)

func (e LSRequestedInfo) String() string {
	switch e {
	case RequestAllFlags:
		return "RequestAllFlags"
	default:
		return fmt.Sprintf("LSRequestedInfo(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/lsrolesmask
type LSRolesMask uint32

const (
	// All: Accepts any role with respect to the item.
	All LSRolesMask = 0
	// Editor: Requests the role [Editor] (theapplication can read, present, manipulate, and save the item).
	Editor LSRolesMask = 0
	// None: Requests the role [None] (theapplication cannot open the item, but provides an icon and a kindstring for it).
	None LSRolesMask = 0
	// Shell: Requests the role [Shell] (theapplication can execute the item).
	Shell LSRolesMask = 0
	// Viewer: Requests the role [Viewer] (theapplication can read and present the item, but cannot manipulateor save it).
	Viewer LSRolesMask = 0
)

func (e LSRolesMask) String() string {
	switch e {
	case All:
		return "All"
	default:
		return fmt.Sprintf("LSRolesMask(%d)", e)
	}
}

type La int

const (
	LaDictionaryNotOpenedErr La = -6992
	LaDictionaryTooManyErr La = -6994
	LaDictionaryUnknownErr La = -6993
	LaEngineNotFoundErr La = -7000
	LaEnvironmentBusyErr La = -6985
	LaEnvironmentExistErr La = -6987
	LaEnvironmentNotFoundErr La = -6986
	LaFailAnalysisErr La = -6990
	LaInvalidPathErr La = -6988
	LaNoMoreMorphemeErr La = -6989
	LaPropertyErr La = -6999
	LaPropertyIsReadOnlyErr La = -6997
	LaPropertyNotFoundErr La = -6998
	LaPropertyUnknownErr La = -6996
	LaPropertyValueErr La = -6995
	LaTextOverFlowErr La = -6991
	LaTooSmallBufferErr La = -6984
)

func (e La) String() string {
	switch e {
	case LaDictionaryNotOpenedErr:
		return "LaDictionaryNotOpenedErr"
	case LaDictionaryTooManyErr:
		return "LaDictionaryTooManyErr"
	case LaDictionaryUnknownErr:
		return "LaDictionaryUnknownErr"
	case LaEngineNotFoundErr:
		return "LaEngineNotFoundErr"
	case LaEnvironmentBusyErr:
		return "LaEnvironmentBusyErr"
	case LaEnvironmentExistErr:
		return "LaEnvironmentExistErr"
	case LaEnvironmentNotFoundErr:
		return "LaEnvironmentNotFoundErr"
	case LaFailAnalysisErr:
		return "LaFailAnalysisErr"
	case LaInvalidPathErr:
		return "LaInvalidPathErr"
	case LaNoMoreMorphemeErr:
		return "LaNoMoreMorphemeErr"
	case LaPropertyErr:
		return "LaPropertyErr"
	case LaPropertyIsReadOnlyErr:
		return "LaPropertyIsReadOnlyErr"
	case LaPropertyNotFoundErr:
		return "LaPropertyNotFoundErr"
	case LaPropertyUnknownErr:
		return "LaPropertyUnknownErr"
	case LaPropertyValueErr:
		return "LaPropertyValueErr"
	case LaTextOverFlowErr:
		return "LaTextOverFlowErr"
	case LaTooSmallBufferErr:
		return "LaTooSmallBufferErr"
	default:
		return fmt.Sprintf("La(%d)", e)
	}
}

type MemROZWarn int

const (
	MemAZErr MemROZWarn = -113
	MemAdrErr MemROZWarn = -110
	MemBCErr MemROZWarn = -115
	MemFullErr MemROZWarn = -108
	MemLockedErr MemROZWarn = -117
	MemPCErr MemROZWarn = -114
	MemPurErr MemROZWarn = -112
	MemROZErr MemROZWarn = -99
	MemROZError MemROZWarn = -99
	MemROZWarnValue MemROZWarn = -99
	MemSCErr MemROZWarn = -116
	MemWZErr MemROZWarn = -111
	NilHandleErr MemROZWarn = -109
)

func (e MemROZWarn) String() string {
	switch e {
	case MemAZErr:
		return "MemAZErr"
	case MemAdrErr:
		return "MemAdrErr"
	case MemBCErr:
		return "MemBCErr"
	case MemFullErr:
		return "MemFullErr"
	case MemLockedErr:
		return "MemLockedErr"
	case MemPCErr:
		return "MemPCErr"
	case MemPurErr:
		return "MemPurErr"
	case MemROZErr:
		return "MemROZErr"
	case MemSCErr:
		return "MemSCErr"
	case MemWZErr:
		return "MemWZErr"
	case NilHandleErr:
		return "NilHandleErr"
	default:
		return fmt.Sprintf("MemROZWarn(%d)", e)
	}
}

type Menu int

const (
	MenuInvalidErr Menu = -5623
	MenuItemNotFoundErr Menu = -5622
	MenuNotFoundErr Menu = -5620
	MenuPropertyInvalid Menu = -5603
	MenuPropertyInvalidErr Menu = -5603
	MenuPropertyNotFoundErr Menu = -5604
	MenuUsesSystemDefErr Menu = -5621
)

func (e Menu) String() string {
	switch e {
	case MenuInvalidErr:
		return "MenuInvalidErr"
	case MenuItemNotFoundErr:
		return "MenuItemNotFoundErr"
	case MenuNotFoundErr:
		return "MenuNotFoundErr"
	case MenuPropertyInvalid:
		return "MenuPropertyInvalid"
	case MenuPropertyNotFoundErr:
		return "MenuPropertyNotFoundErr"
	case MenuUsesSystemDefErr:
		return "MenuUsesSystemDefErr"
	default:
		return fmt.Sprintf("Menu(%d)", e)
	}
}

type Midi int

const (
	MidiDupIDErr Midi = -260
	MidiInvalidCmdErr Midi = -261
	MidiNameLenErr Midi = -259
	MidiNoClientErr Midi = -250
	MidiNoConErr Midi = -257
	MidiNoPortErr Midi = -251
	MidiTooManyConsErr Midi = -253
	MidiTooManyPortsErr Midi = -252
	MidiVConnectErr Midi = -254
	MidiVConnectMade Midi = -255
	MidiVConnectRmvd Midi = -256
	MidiWriteErr Midi = -258
)

func (e Midi) String() string {
	switch e {
	case MidiDupIDErr:
		return "MidiDupIDErr"
	case MidiInvalidCmdErr:
		return "MidiInvalidCmdErr"
	case MidiNameLenErr:
		return "MidiNameLenErr"
	case MidiNoClientErr:
		return "MidiNoClientErr"
	case MidiNoConErr:
		return "MidiNoConErr"
	case MidiNoPortErr:
		return "MidiNoPortErr"
	case MidiTooManyConsErr:
		return "MidiTooManyConsErr"
	case MidiTooManyPortsErr:
		return "MidiTooManyPortsErr"
	case MidiVConnectErr:
		return "MidiVConnectErr"
	case MidiVConnectMade:
		return "MidiVConnectMade"
	case MidiVConnectRmvd:
		return "MidiVConnectRmvd"
	case MidiWriteErr:
		return "MidiWriteErr"
	default:
		return fmt.Sprintf("Midi(%d)", e)
	}
}

const MmInternalError int = -2526

type Nbp int

const (
	NbpBuffOvr Nbp = -1024
	NbpConfDiff Nbp = -1026
	NbpDuplicate Nbp = -1027
	NbpNISErr Nbp = -1029
	NbpNoConfirm Nbp = -1025
	NbpNotFound Nbp = -1028
)

func (e Nbp) String() string {
	switch e {
	case NbpBuffOvr:
		return "NbpBuffOvr"
	case NbpConfDiff:
		return "NbpConfDiff"
	case NbpDuplicate:
		return "NbpDuplicate"
	case NbpNISErr:
		return "NbpNISErr"
	case NbpNoConfirm:
		return "NbpNoConfirm"
	case NbpNotFound:
		return "NbpNotFound"
	default:
		return fmt.Sprintf("Nbp(%d)", e)
	}
}

const NmTypErr int = -299

type No int

const (
	NoScrapErr No = -100
	NoTypeErr No = -102
)

func (e No) String() string {
	switch e {
	case NoScrapErr:
		return "NoScrapErr"
	case NoTypeErr:
		return "NoTypeErr"
	default:
		return fmt.Sprintf("No(%d)", e)
	}
}

type NoAdrMkErr int

const (
	BadBtSlpErr NoAdrMkErr = -70
	BadCksmErr NoAdrMkErr = -69
	BadDBtSlp NoAdrMkErr = -73
	BadDCksum NoAdrMkErr = -72
	BreakRecd NoAdrMkErr = -90
	CantStepErr NoAdrMkErr = -75
	ClkRdErr NoAdrMkErr = -85
	ClkWrErr NoAdrMkErr = -86
	DataVerErr NoAdrMkErr = -68
	Fmt1Err NoAdrMkErr = -82
	Fmt2Err NoAdrMkErr = -83
	InitIWMErr NoAdrMkErr = -77
	NoAdrMkErrValue NoAdrMkErr = -67
	NoDtaMkErr NoAdrMkErr = -71
	PrInitErr NoAdrMkErr = -88
	PrWrErr NoAdrMkErr = -87
	RcvrErr NoAdrMkErr = -89
	SectNFErr NoAdrMkErr = -81
	SeekErr NoAdrMkErr = -80
	SpdAdjErr NoAdrMkErr = -79
	Tk0BadErr NoAdrMkErr = -76
	TwoSideErr NoAdrMkErr = -78
	VerErr NoAdrMkErr = -84
	WrUnderrun NoAdrMkErr = -74
)

func (e NoAdrMkErr) String() string {
	switch e {
	case BadBtSlpErr:
		return "BadBtSlpErr"
	case BadCksmErr:
		return "BadCksmErr"
	case BadDBtSlp:
		return "BadDBtSlp"
	case BadDCksum:
		return "BadDCksum"
	case BreakRecd:
		return "BreakRecd"
	case CantStepErr:
		return "CantStepErr"
	case ClkRdErr:
		return "ClkRdErr"
	case ClkWrErr:
		return "ClkWrErr"
	case DataVerErr:
		return "DataVerErr"
	case Fmt1Err:
		return "Fmt1Err"
	case Fmt2Err:
		return "Fmt2Err"
	case InitIWMErr:
		return "InitIWMErr"
	case NoAdrMkErrValue:
		return "NoAdrMkErrValue"
	case NoDtaMkErr:
		return "NoDtaMkErr"
	case PrInitErr:
		return "PrInitErr"
	case PrWrErr:
		return "PrWrErr"
	case RcvrErr:
		return "RcvrErr"
	case SectNFErr:
		return "SectNFErr"
	case SeekErr:
		return "SeekErr"
	case SpdAdjErr:
		return "SpdAdjErr"
	case Tk0BadErr:
		return "Tk0BadErr"
	case TwoSideErr:
		return "TwoSideErr"
	case VerErr:
		return "VerErr"
	case WrUnderrun:
		return "WrUnderrun"
	default:
		return fmt.Sprintf("NoAdrMkErr(%d)", e)
	}
}

type NoDeviceForChannel int

const (
	BadControllerHeight NoDeviceForChannel = -9994
	BadSGChannel NoDeviceForChannel = -9406
	CannotMoveAttachedController NoDeviceForChannel = -9999
	CannotSetWidthOfAttachedController NoDeviceForChannel = -9997
	CantDoThatInCurrentMode NoDeviceForChannel = -9402
	ControllerBoundsNotExact NoDeviceForChannel = -9996
	ControllerHasFixedHeight NoDeviceForChannel = -9998
	CouldntGetRequiredComponent NoDeviceForChannel = -9405
	DeviceCantMeetRequest NoDeviceForChannel = -9408
	EditingNotAllowed NoDeviceForChannel = -9995
	GrabTimeComplete NoDeviceForChannel = -9401
	NoDeviceForChannelValue NoDeviceForChannel = -9400
	NotEnoughDiskSpaceToGrab NoDeviceForChannel = -9404
	NotEnoughMemoryToGrab NoDeviceForChannel = -9403
	SeqGrabInfoNotAvailable NoDeviceForChannel = -9407
)

func (e NoDeviceForChannel) String() string {
	switch e {
	case BadControllerHeight:
		return "BadControllerHeight"
	case BadSGChannel:
		return "BadSGChannel"
	case CannotMoveAttachedController:
		return "CannotMoveAttachedController"
	case CannotSetWidthOfAttachedController:
		return "CannotSetWidthOfAttachedController"
	case CantDoThatInCurrentMode:
		return "CantDoThatInCurrentMode"
	case ControllerBoundsNotExact:
		return "ControllerBoundsNotExact"
	case ControllerHasFixedHeight:
		return "ControllerHasFixedHeight"
	case CouldntGetRequiredComponent:
		return "CouldntGetRequiredComponent"
	case DeviceCantMeetRequest:
		return "DeviceCantMeetRequest"
	case EditingNotAllowed:
		return "EditingNotAllowed"
	case GrabTimeComplete:
		return "GrabTimeComplete"
	case NoDeviceForChannelValue:
		return "NoDeviceForChannelValue"
	case NotEnoughDiskSpaceToGrab:
		return "NotEnoughDiskSpaceToGrab"
	case NotEnoughMemoryToGrab:
		return "NotEnoughMemoryToGrab"
	case SeqGrabInfoNotAvailable:
		return "SeqGrabInfoNotAvailable"
	default:
		return fmt.Sprintf("NoDeviceForChannel(%d)", e)
	}
}

type NoHardware int

const (
	BadChannel NoHardware = -205
	BadFileFormat NoHardware = -208
	BadFormat NoHardware = -206
	BuffersTooSmall NoHardware = -210
	ChannelBusy NoHardware = -209
	ChannelNotBusy NoHardware = -211
	NoHardwareValue NoHardware = -200
	NoMoreRealTime NoHardware = -212
	NotEnoughBufferSpace NoHardware = -207
	NotEnoughHardware NoHardware = -201
	QueueFull NoHardware = -203
	ResProblem NoHardware = -204
	SiBadDeviceName NoHardware = -228
	SiBadRefNum NoHardware = -229
	SiBadSoundInDevice NoHardware = -221
	SiDeviceBusyErr NoHardware = -227
	SiHardDriveTooSlow NoHardware = -224
	SiInputDeviceErr NoHardware = -230
	SiInvalidCompression NoHardware = -223
	SiInvalidSampleRate NoHardware = -225
	SiInvalidSampleSize NoHardware = -226
	SiNoBufferSpecified NoHardware = -222
	SiNoSoundInHardware NoHardware = -220
	SiUnknownInfoType NoHardware = -231
	SiUnknownQuality NoHardware = -232
	SiVBRCompressionNotSupported NoHardware = -213
)

func (e NoHardware) String() string {
	switch e {
	case BadChannel:
		return "BadChannel"
	case BadFileFormat:
		return "BadFileFormat"
	case BadFormat:
		return "BadFormat"
	case BuffersTooSmall:
		return "BuffersTooSmall"
	case ChannelBusy:
		return "ChannelBusy"
	case ChannelNotBusy:
		return "ChannelNotBusy"
	case NoHardwareValue:
		return "NoHardwareValue"
	case NoMoreRealTime:
		return "NoMoreRealTime"
	case NotEnoughBufferSpace:
		return "NotEnoughBufferSpace"
	case NotEnoughHardware:
		return "NotEnoughHardware"
	case QueueFull:
		return "QueueFull"
	case ResProblem:
		return "ResProblem"
	case SiBadDeviceName:
		return "SiBadDeviceName"
	case SiBadRefNum:
		return "SiBadRefNum"
	case SiBadSoundInDevice:
		return "SiBadSoundInDevice"
	case SiDeviceBusyErr:
		return "SiDeviceBusyErr"
	case SiHardDriveTooSlow:
		return "SiHardDriveTooSlow"
	case SiInputDeviceErr:
		return "SiInputDeviceErr"
	case SiInvalidCompression:
		return "SiInvalidCompression"
	case SiInvalidSampleRate:
		return "SiInvalidSampleRate"
	case SiInvalidSampleSize:
		return "SiInvalidSampleSize"
	case SiNoBufferSpecified:
		return "SiNoBufferSpecified"
	case SiNoSoundInHardware:
		return "SiNoSoundInHardware"
	case SiUnknownInfoType:
		return "SiUnknownInfoType"
	case SiUnknownQuality:
		return "SiUnknownQuality"
	case SiVBRCompressionNotSupported:
		return "SiVBRCompressionNotSupported"
	default:
		return fmt.Sprintf("NoHardware(%d)", e)
	}
}

const NoMaskFoundErr int = -1000

type NoSynthFound int

const (
	// BadDictFormat: Pronunciation dictionary format error
	BadDictFormat NoSynthFound = -246
	// BadInputText: Raw phoneme text contains invalid characters
	BadInputText NoSynthFound = -247
	// BufTooSmall: Output buffer is too small to hold result
	BufTooSmall NoSynthFound = -243
	// IncompatibleVoice: Specified voice cannot be used with synthesizer
	IncompatibleVoice NoSynthFound = -245
	// NoSynthFoundValue: Could not find the specified speech synthesizer
	NoSynthFoundValue NoSynthFound = -240
	// SynthNotReady: Speech synthesizer is still busy speaking
	SynthNotReady NoSynthFound = -242
	// SynthOpenFailed: Could not open another speech synthesizerchannel
	SynthOpenFailed NoSynthFound = -241
	// VoiceNotFound: Voice resource not found
	VoiceNotFound NoSynthFound = -244
)

func (e NoSynthFound) String() string {
	switch e {
	case BadDictFormat:
		return "BadDictFormat"
	case BadInputText:
		return "BadInputText"
	case BufTooSmall:
		return "BufTooSmall"
	case IncompatibleVoice:
		return "IncompatibleVoice"
	case NoSynthFoundValue:
		return "NoSynthFoundValue"
	case SynthNotReady:
		return "SynthNotReady"
	case SynthOpenFailed:
		return "SynthOpenFailed"
	case VoiceNotFound:
		return "VoiceNotFound"
	default:
		return fmt.Sprintf("NoSynthFound(%d)", e)
	}
}

type NotAQTVRMovieErr int

const (
	CallNotSupportedByNodeErr NotAQTVRMovieErr = -30542
	ConstraintReachedErr NotAQTVRMovieErr = -30541
	InvalidHotSpotIDErr NotAQTVRMovieErr = -30551
	InvalidNodeFormatErr NotAQTVRMovieErr = -30550
	InvalidNodeIDErr NotAQTVRMovieErr = -30544
	InvalidViewStateErr NotAQTVRMovieErr = -30545
	LimitReachedErr NotAQTVRMovieErr = -30549
	NoMemoryNodeFailedInitialize NotAQTVRMovieErr = -30552
	NotAQTVRMovieErrValue NotAQTVRMovieErr = -30540
	PropertyNotSupportedByNodeErr NotAQTVRMovieErr = -30547
	QtvrLibraryLoadErr NotAQTVRMovieErr = -30554
	QtvrUninitialized NotAQTVRMovieErr = -30555
	SelectorNotSupportedByNodeErr NotAQTVRMovieErr = -30543
	SettingNotSupportedByNodeErr NotAQTVRMovieErr = -30548
	StreamingNodeNotReadyErr NotAQTVRMovieErr = -30553
	TimeNotInViewErr NotAQTVRMovieErr = -30546
)

func (e NotAQTVRMovieErr) String() string {
	switch e {
	case CallNotSupportedByNodeErr:
		return "CallNotSupportedByNodeErr"
	case ConstraintReachedErr:
		return "ConstraintReachedErr"
	case InvalidHotSpotIDErr:
		return "InvalidHotSpotIDErr"
	case InvalidNodeFormatErr:
		return "InvalidNodeFormatErr"
	case InvalidNodeIDErr:
		return "InvalidNodeIDErr"
	case InvalidViewStateErr:
		return "InvalidViewStateErr"
	case LimitReachedErr:
		return "LimitReachedErr"
	case NoMemoryNodeFailedInitialize:
		return "NoMemoryNodeFailedInitialize"
	case NotAQTVRMovieErrValue:
		return "NotAQTVRMovieErrValue"
	case PropertyNotSupportedByNodeErr:
		return "PropertyNotSupportedByNodeErr"
	case QtvrLibraryLoadErr:
		return "QtvrLibraryLoadErr"
	case QtvrUninitialized:
		return "QtvrUninitialized"
	case SelectorNotSupportedByNodeErr:
		return "SelectorNotSupportedByNodeErr"
	case SettingNotSupportedByNodeErr:
		return "SettingNotSupportedByNodeErr"
	case StreamingNodeNotReadyErr:
		return "StreamingNodeNotReadyErr"
	case TimeNotInViewErr:
		return "TimeNotInViewErr"
	default:
		return fmt.Sprintf("NotAQTVRMovieErr(%d)", e)
	}
}

type NotBTree int

const (
	BtDupRecErr NotBTree = -414
	BtKeyAttrErr NotBTree = -417
	BtKeyLenErr NotBTree = -416
	BtNoSpace NotBTree = -413
	BtRecNotFnd NotBTree = -415
	InvalidIndexErr NotBTree = -20002
	NotBTreeValue NotBTree = -410
	RecordDataTooBigErr NotBTree = -20001
	UnknownInsertModeErr NotBTree = -20000
)

func (e NotBTree) String() string {
	switch e {
	case BtDupRecErr:
		return "BtDupRecErr"
	case BtKeyAttrErr:
		return "BtKeyAttrErr"
	case BtKeyLenErr:
		return "BtKeyLenErr"
	case BtNoSpace:
		return "BtNoSpace"
	case BtRecNotFnd:
		return "BtRecNotFnd"
	case InvalidIndexErr:
		return "InvalidIndexErr"
	case NotBTreeValue:
		return "NotBTreeValue"
	case RecordDataTooBigErr:
		return "RecordDataTooBigErr"
	case UnknownInsertModeErr:
		return "UnknownInsertModeErr"
	default:
		return fmt.Sprintf("NotBTree(%d)", e)
	}
}

type NotEnoughMemoryErr int

const (
	CannotDeferErr NotEnoughMemoryErr = -625
	CannotMakeContiguousErr NotEnoughMemoryErr = -622
	InterruptsMaskedErr NotEnoughMemoryErr = -624
	NoMMUErr NotEnoughMemoryErr = -626
	NotEnoughMemoryErrValue NotEnoughMemoryErr = -620
	NotHeldErr NotEnoughMemoryErr = -621
	NotLockedErr NotEnoughMemoryErr = -623
)

func (e NotEnoughMemoryErr) String() string {
	switch e {
	case CannotDeferErr:
		return "CannotDeferErr"
	case CannotMakeContiguousErr:
		return "CannotMakeContiguousErr"
	case InterruptsMaskedErr:
		return "InterruptsMaskedErr"
	case NoMMUErr:
		return "NoMMUErr"
	case NotEnoughMemoryErrValue:
		return "NotEnoughMemoryErrValue"
	case NotHeldErr:
		return "NotHeldErr"
	case NotLockedErr:
		return "NotLockedErr"
	default:
		return fmt.Sprintf("NotEnoughMemoryErr(%d)", e)
	}
}

type NotInitErr int

const (
	AuthFailErr NotInitErr = -927
	BadLocNameErr NotInitErr = -931
	BadPortNameErr NotInitErr = -919
	BadReqErr NotInitErr = -909
	BadServiceMethodErr NotInitErr = -930
	// DestPortErr: Server hasn’t set `'SIZE'` resource toindicate awareness of high-level events, or else is not present
	DestPortErr NotInitErr = -906
	GuestNotAllowedErr NotInitErr = -932
	LocalOnlyErr NotInitErr = -905
	NameTypeErr NotInitErr = -902
	NetworkErr NotInitErr = -925
	NoDefaultUserErr NotInitErr = -922
	NoGlobalsErr NotInitErr = -904
	NoInformErr NotInitErr = -926
	NoMachineNameErr NotInitErr = -913
	// NoPortErr: Client hasn’t set `'SIZE'` resource toindicate awareness of high-level events
	NoPortErr NotInitErr = -903
	NoResponseErr NotInitErr = -915
	NoSessionErr NotInitErr = -908
	NoToolboxNameErr NotInitErr = -914
	NoUserNameErr NotInitErr = -911
	NoUserRecErr NotInitErr = -928
	NoUserRefErr NotInitErr = -924
	NotInitErrValue NotInitErr = -900
	NotLoggedInErr NotInitErr = -923
	PortClosedErr NotInitErr = -916
	PortNameExistsErr NotInitErr = -910
	// SessClosedErr: The `kAEDontReconnect` flagin the `sendMode` parameterwas set and the server quit, then restarted
	SessClosedErr NotInitErr = -917
	SessTableErr NotInitErr = -907
	UserRejectErr NotInitErr = -912
)

func (e NotInitErr) String() string {
	switch e {
	case AuthFailErr:
		return "AuthFailErr"
	case BadLocNameErr:
		return "BadLocNameErr"
	case BadPortNameErr:
		return "BadPortNameErr"
	case BadReqErr:
		return "BadReqErr"
	case BadServiceMethodErr:
		return "BadServiceMethodErr"
	case DestPortErr:
		return "DestPortErr"
	case GuestNotAllowedErr:
		return "GuestNotAllowedErr"
	case LocalOnlyErr:
		return "LocalOnlyErr"
	case NameTypeErr:
		return "NameTypeErr"
	case NetworkErr:
		return "NetworkErr"
	case NoDefaultUserErr:
		return "NoDefaultUserErr"
	case NoGlobalsErr:
		return "NoGlobalsErr"
	case NoInformErr:
		return "NoInformErr"
	case NoMachineNameErr:
		return "NoMachineNameErr"
	case NoPortErr:
		return "NoPortErr"
	case NoResponseErr:
		return "NoResponseErr"
	case NoSessionErr:
		return "NoSessionErr"
	case NoToolboxNameErr:
		return "NoToolboxNameErr"
	case NoUserNameErr:
		return "NoUserNameErr"
	case NoUserRecErr:
		return "NoUserRecErr"
	case NoUserRefErr:
		return "NoUserRefErr"
	case NotInitErrValue:
		return "NotInitErrValue"
	case NotLoggedInErr:
		return "NotLoggedInErr"
	case PortClosedErr:
		return "PortClosedErr"
	case PortNameExistsErr:
		return "PortNameExistsErr"
	case SessClosedErr:
		return "SessClosedErr"
	case SessTableErr:
		return "SessTableErr"
	case UserRejectErr:
		return "UserRejectErr"
	default:
		return fmt.Sprintf("NotInitErr(%d)", e)
	}
}

type Nr int

const (
	NrCallNotSupported Nr = -2557
	NrDataTruncatedErr Nr = -2543
	NrExitedIteratorScope Nr = -2555
	NrInvalidEntryIterationOp Nr = -2552
	NrInvalidNodeErr Nr = -2538
	NrIterationDone Nr = -2554
	NrLockedErr Nr = -2536
	NrNameErr Nr = -2541
	NrNotCreatedErr Nr = -2540
	NrNotEnoughMemoryErr Nr = -2537
	NrNotFoundErr Nr = -2539
	NrNotModifiedErr Nr = -2547
	NrNotSlotDeviceErr Nr = -2542
	NrOverrunErr Nr = -2548
	NrPathBufferTooSmall Nr = -2551
	NrPathNotFound Nr = -2550
	NrPowerErr Nr = -2544
	NrPowerSwitchAbortErr Nr = -2545
	NrPropertyAlreadyExists Nr = -2553
	NrResultCodeBase Nr = -2549
	NrTransactionAborted Nr = -2556
	NrTypeMismatchErr Nr = -2546
)

func (e Nr) String() string {
	switch e {
	case NrCallNotSupported:
		return "NrCallNotSupported"
	case NrDataTruncatedErr:
		return "NrDataTruncatedErr"
	case NrExitedIteratorScope:
		return "NrExitedIteratorScope"
	case NrInvalidEntryIterationOp:
		return "NrInvalidEntryIterationOp"
	case NrInvalidNodeErr:
		return "NrInvalidNodeErr"
	case NrIterationDone:
		return "NrIterationDone"
	case NrLockedErr:
		return "NrLockedErr"
	case NrNameErr:
		return "NrNameErr"
	case NrNotCreatedErr:
		return "NrNotCreatedErr"
	case NrNotEnoughMemoryErr:
		return "NrNotEnoughMemoryErr"
	case NrNotFoundErr:
		return "NrNotFoundErr"
	case NrNotModifiedErr:
		return "NrNotModifiedErr"
	case NrNotSlotDeviceErr:
		return "NrNotSlotDeviceErr"
	case NrOverrunErr:
		return "NrOverrunErr"
	case NrPathBufferTooSmall:
		return "NrPathBufferTooSmall"
	case NrPathNotFound:
		return "NrPathNotFound"
	case NrPowerErr:
		return "NrPowerErr"
	case NrPowerSwitchAbortErr:
		return "NrPowerSwitchAbortErr"
	case NrPropertyAlreadyExists:
		return "NrPropertyAlreadyExists"
	case NrResultCodeBase:
		return "NrResultCodeBase"
	case NrTransactionAborted:
		return "NrTransactionAborted"
	case NrTypeMismatchErr:
		return "NrTypeMismatchErr"
	default:
		return fmt.Sprintf("Nr(%d)", e)
	}
}

type Number int

const (
	NumberFormattingBadCurrencyPositionErr Number = -5211
	NumberFormattingBadFormatErr Number = -5207
	NumberFormattingBadNumberFormattingObjectErr Number = -5202
	NumberFormattingBadOptionsErr Number = -5208
	NumberFormattingBadTokenErr Number = -5209
	NumberFormattingDelimiterMissingErr Number = -5205
	NumberFormattingEmptyFormatErr Number = -5206
	NumberFormattingLiteralMissingErr Number = -5204
	NumberFormattingNotADigitErr Number = -5212
	NumberFormattingNotANumberErr Number = -5200
	NumberFormattingOverflowInDestinationErr Number = -5201
	NumberFormattingSpuriousCharErr Number = -5203
	NumberFormattingUnOrderedCurrencyRangeErr Number = -5210
	NumberFormattingUnOrdredCurrencyRangeErr Number = -5210
	NumberFortmattingNotADigitErr Number = -5212
)

func (e Number) String() string {
	switch e {
	case NumberFormattingBadCurrencyPositionErr:
		return "NumberFormattingBadCurrencyPositionErr"
	case NumberFormattingBadFormatErr:
		return "NumberFormattingBadFormatErr"
	case NumberFormattingBadNumberFormattingObjectErr:
		return "NumberFormattingBadNumberFormattingObjectErr"
	case NumberFormattingBadOptionsErr:
		return "NumberFormattingBadOptionsErr"
	case NumberFormattingBadTokenErr:
		return "NumberFormattingBadTokenErr"
	case NumberFormattingDelimiterMissingErr:
		return "NumberFormattingDelimiterMissingErr"
	case NumberFormattingEmptyFormatErr:
		return "NumberFormattingEmptyFormatErr"
	case NumberFormattingLiteralMissingErr:
		return "NumberFormattingLiteralMissingErr"
	case NumberFormattingNotADigitErr:
		return "NumberFormattingNotADigitErr"
	case NumberFormattingNotANumberErr:
		return "NumberFormattingNotANumberErr"
	case NumberFormattingOverflowInDestinationErr:
		return "NumberFormattingOverflowInDestinationErr"
	case NumberFormattingSpuriousCharErr:
		return "NumberFormattingSpuriousCharErr"
	case NumberFormattingUnOrderedCurrencyRangeErr:
		return "NumberFormattingUnOrderedCurrencyRangeErr"
	default:
		return fmt.Sprintf("Number(%d)", e)
	}
}

type ParamErr int

const (
	BadUnitErr ParamErr = -21
	ClosErr ParamErr = -24
	ControlErr ParamErr = -17
	CorErr ParamErr = -3
	DInstErr ParamErr = -26
	DRemovErr ParamErr = -25
	NoHardwareErr ParamErr = -200
	NotEnoughHardwareErr ParamErr = -201
	OpenErr ParamErr = -23
	ParamErrValue ParamErr = -50
	QErr ParamErr = -1
	ReadErr ParamErr = -19
	SeNoDB ParamErr = -8
	SlpTypeErr ParamErr = -5
	StatusErr ParamErr = -18
	UnimpErr ParamErr = -4
	UnitEmptyErr ParamErr = -22
	UserCanceledErr ParamErr = -128
	VTypErr ParamErr = -2
	WritErr ParamErr = -20
)

func (e ParamErr) String() string {
	switch e {
	case BadUnitErr:
		return "BadUnitErr"
	case ClosErr:
		return "ClosErr"
	case ControlErr:
		return "ControlErr"
	case CorErr:
		return "CorErr"
	case DInstErr:
		return "DInstErr"
	case DRemovErr:
		return "DRemovErr"
	case NoHardwareErr:
		return "NoHardwareErr"
	case NotEnoughHardwareErr:
		return "NotEnoughHardwareErr"
	case OpenErr:
		return "OpenErr"
	case ParamErrValue:
		return "ParamErrValue"
	case QErr:
		return "QErr"
	case ReadErr:
		return "ReadErr"
	case SeNoDB:
		return "SeNoDB"
	case SlpTypeErr:
		return "SlpTypeErr"
	case StatusErr:
		return "StatusErr"
	case UnimpErr:
		return "UnimpErr"
	case UnitEmptyErr:
		return "UnitEmptyErr"
	case UserCanceledErr:
		return "UserCanceledErr"
	case VTypErr:
		return "VTypErr"
	case WritErr:
		return "WritErr"
	default:
		return fmt.Sprintf("ParamErr(%d)", e)
	}
}

type Pm int

const (
	PmBusyErr Pm = -13000
	PmRecvEndErr Pm = -13005
	PmRecvStartErr Pm = -13004
	PmReplyTOErr Pm = -13001
	PmSendEndErr Pm = -13003
	PmSendStartErr Pm = -13002
)

func (e Pm) String() string {
	switch e {
	case PmBusyErr:
		return "PmBusyErr"
	case PmRecvEndErr:
		return "PmRecvEndErr"
	case PmRecvStartErr:
		return "PmRecvStartErr"
	case PmReplyTOErr:
		return "PmReplyTOErr"
	case PmSendEndErr:
		return "PmSendEndErr"
	case PmSendStartErr:
		return "PmSendStartErr"
	default:
		return fmt.Sprintf("Pm(%d)", e)
	}
}

const PrinterStatusOpCodeNotSupportedErr int = -25280

type ProcNotFound int

const (
	AppIsDaemon ProcNotFound = -606
	AppMemFullErr ProcNotFound = -605
	AppModeErr ProcNotFound = -602
	BufferIsSmall ProcNotFound = -607
	ConnectionInvalid ProcNotFound = -609
	HardwareConfigErr ProcNotFound = -604
	MemFragErr ProcNotFound = -601
	NoOutstandingHLE ProcNotFound = -608
	NoUserInteractionAllowed ProcNotFound = -610
	ProcNotFoundValue ProcNotFound = -600
	ProtocolErr ProcNotFound = -603
)

func (e ProcNotFound) String() string {
	switch e {
	case AppIsDaemon:
		return "AppIsDaemon"
	case AppMemFullErr:
		return "AppMemFullErr"
	case AppModeErr:
		return "AppModeErr"
	case BufferIsSmall:
		return "BufferIsSmall"
	case ConnectionInvalid:
		return "ConnectionInvalid"
	case HardwareConfigErr:
		return "HardwareConfigErr"
	case MemFragErr:
		return "MemFragErr"
	case NoOutstandingHLE:
		return "NoOutstandingHLE"
	case NoUserInteractionAllowed:
		return "NoUserInteractionAllowed"
	case ProcNotFoundValue:
		return "ProcNotFoundValue"
	case ProtocolErr:
		return "ProtocolErr"
	default:
		return fmt.Sprintf("ProcNotFound(%d)", e)
	}
}

type Qts int

const (
	QtsAddressBusyErr Qts = -5421
	QtsBadDataErr Qts = -5402
	QtsBadSelectorErr Qts = -5400
	QtsBadStateErr Qts = -5401
	QtsConnectionFailedErr Qts = -5420
	QtsTimeoutErr Qts = -5408
	QtsTooMuchDataErr Qts = -5406
	QtsUnknownValueErr Qts = -5407
	QtsUnsupportedDataTypeErr Qts = -5403
	QtsUnsupportedFeatureErr Qts = -5405
	QtsUnsupportedRateErr Qts = -5404
)

func (e Qts) String() string {
	switch e {
	case QtsAddressBusyErr:
		return "QtsAddressBusyErr"
	case QtsBadDataErr:
		return "QtsBadDataErr"
	case QtsBadSelectorErr:
		return "QtsBadSelectorErr"
	case QtsBadStateErr:
		return "QtsBadStateErr"
	case QtsConnectionFailedErr:
		return "QtsConnectionFailedErr"
	case QtsTimeoutErr:
		return "QtsTimeoutErr"
	case QtsTooMuchDataErr:
		return "QtsTooMuchDataErr"
	case QtsUnknownValueErr:
		return "QtsUnknownValueErr"
	case QtsUnsupportedDataTypeErr:
		return "QtsUnsupportedDataTypeErr"
	case QtsUnsupportedFeatureErr:
		return "QtsUnsupportedFeatureErr"
	case QtsUnsupportedRateErr:
		return "QtsUnsupportedRateErr"
	default:
		return fmt.Sprintf("Qts(%d)", e)
	}
}

type RcDB int

const (
	RcDBAsyncNotSupp RcDB = -809
	RcDBBadAsyncPB RcDB = -810
	RcDBBadDDEV RcDB = -808
	RcDBBadSessID RcDB = -806
	RcDBBadSessNum RcDB = -807
	RcDBBadType RcDB = -803
	RcDBBreak RcDB = -804
	RcDBError RcDB = -802
	RcDBExec RcDB = -805
	RcDBNoHandler RcDB = -811
	RcDBNull RcDB = -800
	RcDBPackNotInited RcDB = -813
	RcDBValue RcDB = -801
	RcDBWrongVersion RcDB = -812
)

func (e RcDB) String() string {
	switch e {
	case RcDBAsyncNotSupp:
		return "RcDBAsyncNotSupp"
	case RcDBBadAsyncPB:
		return "RcDBBadAsyncPB"
	case RcDBBadDDEV:
		return "RcDBBadDDEV"
	case RcDBBadSessID:
		return "RcDBBadSessID"
	case RcDBBadSessNum:
		return "RcDBBadSessNum"
	case RcDBBadType:
		return "RcDBBadType"
	case RcDBBreak:
		return "RcDBBreak"
	case RcDBError:
		return "RcDBError"
	case RcDBExec:
		return "RcDBExec"
	case RcDBNoHandler:
		return "RcDBNoHandler"
	case RcDBNull:
		return "RcDBNull"
	case RcDBPackNotInited:
		return "RcDBPackNotInited"
	case RcDBValue:
		return "RcDBValue"
	case RcDBWrongVersion:
		return "RcDBWrongVersion"
	default:
		return fmt.Sprintf("RcDB(%d)", e)
	}
}

type ReqFailed int

const (
	BadATPSkt ReqFailed = -1099
	BadBuffNum ReqFailed = -1100
	CbNotFound ReqFailed = -1102
	NoDataArea ReqFailed = -1104
	NoRelErr ReqFailed = -1101
	NoSendResp ReqFailed = -1103
	ReqAborted ReqFailed = -1105
	ReqFailedValue ReqFailed = -1096
	TooManyReqs ReqFailed = -1097
	TooManySkts ReqFailed = -1098
)

func (e ReqFailed) String() string {
	switch e {
	case BadATPSkt:
		return "BadATPSkt"
	case BadBuffNum:
		return "BadBuffNum"
	case CbNotFound:
		return "CbNotFound"
	case NoDataArea:
		return "NoDataArea"
	case NoRelErr:
		return "NoRelErr"
	case NoSendResp:
		return "NoSendResp"
	case ReqAborted:
		return "ReqAborted"
	case ReqFailedValue:
		return "ReqFailedValue"
	case TooManyReqs:
		return "TooManyReqs"
	case TooManySkts:
		return "TooManySkts"
	default:
		return fmt.Sprintf("ReqFailed(%d)", e)
	}
}

type ResourceInMemory int

const (
	AddRefFailed ResourceInMemory = -195
	AddResFailed ResourceInMemory = -194
	BadExtResource ResourceInMemory = -185
	CantDecompress ResourceInMemory = -186
	InputOutOfBounds ResourceInMemory = -190
	InsufficientStackErr ResourceInMemory = -149
	MapReadErr ResourceInMemory = -199
	NoMemForPictPlaybackErr ResourceInMemory = -145
	NsStackErr ResourceInMemory = -149
	PixMapTooDeepErr ResourceInMemory = -148
	ResAttrErr ResourceInMemory = -198
	ResFNotFound ResourceInMemory = -193
	ResNotFound ResourceInMemory = -192
	ResourceInMemoryValue ResourceInMemory = -188
	RgnOverflowErr ResourceInMemory = -147
	RgnTooBigError ResourceInMemory = -147
	RmvRefFailed ResourceInMemory = -197
	RmvResFailed ResourceInMemory = -196
	WritingPastEnd ResourceInMemory = -189
)

func (e ResourceInMemory) String() string {
	switch e {
	case AddRefFailed:
		return "AddRefFailed"
	case AddResFailed:
		return "AddResFailed"
	case BadExtResource:
		return "BadExtResource"
	case CantDecompress:
		return "CantDecompress"
	case InputOutOfBounds:
		return "InputOutOfBounds"
	case InsufficientStackErr:
		return "InsufficientStackErr"
	case MapReadErr:
		return "MapReadErr"
	case NoMemForPictPlaybackErr:
		return "NoMemForPictPlaybackErr"
	case PixMapTooDeepErr:
		return "PixMapTooDeepErr"
	case ResAttrErr:
		return "ResAttrErr"
	case ResFNotFound:
		return "ResFNotFound"
	case ResNotFound:
		return "ResNotFound"
	case ResourceInMemoryValue:
		return "ResourceInMemoryValue"
	case RgnOverflowErr:
		return "RgnOverflowErr"
	case RmvRefFailed:
		return "RmvRefFailed"
	case RmvResFailed:
		return "RmvResFailed"
	case WritingPastEnd:
		return "WritingPastEnd"
	default:
		return fmt.Sprintf("ResourceInMemory(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/skdocumentindexstate
type SKDocumentIndexState uint32

const (
	// KSKDocumentStateAddPending: Specifies that the document is not in the index but will be added after the index is flushed or closed.
	KSKDocumentStateAddPending SKDocumentIndexState = 2
	// KSKDocumentStateDeletePending: Specifies that the document is in the index but will be deleted after the index is flushed or closed.
	KSKDocumentStateDeletePending SKDocumentIndexState = 3
	// KSKDocumentStateIndexed: Specifies that the document is indexed.
	KSKDocumentStateIndexed SKDocumentIndexState = 1
	// KSKDocumentStateNotIndexed: Specifies that the document is not indexed.
	KSKDocumentStateNotIndexed SKDocumentIndexState = 0
)

func (e SKDocumentIndexState) String() string {
	switch e {
	case KSKDocumentStateAddPending:
		return "KSKDocumentStateAddPending"
	case KSKDocumentStateDeletePending:
		return "KSKDocumentStateDeletePending"
	case KSKDocumentStateIndexed:
		return "KSKDocumentStateIndexed"
	case KSKDocumentStateNotIndexed:
		return "KSKDocumentStateNotIndexed"
	default:
		return fmt.Sprintf("SKDocumentIndexState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/skindextype
type SKIndexType uint32

const (
	// KSKIndexInverted: Specifies an inverted index, mapping terms to documents.
	KSKIndexInverted SKIndexType = 1
	// KSKIndexInvertedVector: Specifies an index type with all the capabilities of an inverted and a vector index.
	KSKIndexInvertedVector SKIndexType = 3
	// KSKIndexUnknown: Specifies an unknown index type.
	KSKIndexUnknown SKIndexType = 0
	// KSKIndexVector: Specifies a vector index, mapping documents to terms.
	KSKIndexVector SKIndexType = 2
)

func (e SKIndexType) String() string {
	switch e {
	case KSKIndexInverted:
		return "KSKIndexInverted"
	case KSKIndexInvertedVector:
		return "KSKIndexInvertedVector"
	case KSKIndexUnknown:
		return "KSKIndexUnknown"
	case KSKIndexVector:
		return "KSKIndexVector"
	default:
		return fmt.Sprintf("SKIndexType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/sksearchtype
type SKSearchType uint32

const (
	// KSKSearchBooleanRanked: Deprecated.
	KSKSearchBooleanRanked SKSearchType = 1
	// KSKSearchPrefixRanked: Deprecated.
	KSKSearchPrefixRanked SKSearchType = 3
	// KSKSearchRanked: Deprecated.
	KSKSearchRanked SKSearchType = 0
	// KSKSearchRequiredRanked: Deprecated.
	KSKSearchRequiredRanked SKSearchType = 2
)

func (e SKSearchType) String() string {
	switch e {
	case KSKSearchBooleanRanked:
		return "KSKSearchBooleanRanked"
	case KSKSearchPrefixRanked:
		return "KSKSearchPrefixRanked"
	case KSKSearchRanked:
		return "KSKSearchRanked"
	case KSKSearchRequiredRanked:
		return "KSKSearchRequiredRanked"
	default:
		return fmt.Sprintf("SKSearchType(%d)", e)
	}
}

type SiInitSDTblErr int

const (
	SdmInitErr SiInitSDTblErr = 11
	SdmJTInitErr SiInitSDTblErr = 10
	SdmPRAMInitErr SiInitSDTblErr = 13
	SdmPriInitErr SiInitSDTblErr = 14
	SdmSRTInitErr SiInitSDTblErr = 12
	SiInitSDTblErrValue SiInitSDTblErr = 1
	SiInitSPTblErr SiInitSDTblErr = 3
	SiInitVBLQsErr SiInitSDTblErr = 2
)

func (e SiInitSDTblErr) String() string {
	switch e {
	case SdmInitErr:
		return "SdmInitErr"
	case SdmJTInitErr:
		return "SdmJTInitErr"
	case SdmPRAMInitErr:
		return "SdmPRAMInitErr"
	case SdmPriInitErr:
		return "SdmPriInitErr"
	case SdmSRTInitErr:
		return "SdmSRTInitErr"
	case SiInitSDTblErrValue:
		return "SiInitSDTblErrValue"
	case SiInitSPTblErr:
		return "SiInitSPTblErr"
	case SiInitVBLQsErr:
		return "SiInitVBLQsErr"
	default:
		return fmt.Sprintf("SiInitSDTblErr(%d)", e)
	}
}

type Sm int

const (
	SmBadsPtrErr Sm = -346
	SmBlkMoveErr Sm = -340
	SmByteLanesErr Sm = -347
	SmCPUErr Sm = -334
	SmCRCFail Sm = -301
	SmCkStatusErr Sm = -341
	SmDisDrvrNamErr Sm = -343
	SmDisabledSlot Sm = -305
	SmEmptySlot Sm = -300
	SmFormatErr Sm = -302
	SmGetDrvrNamErr Sm = -342
	SmNewPErr Sm = -339
	SmNilsBlockErr Sm = -336
	SmNoDir Sm = -304
	SmNoGoodOpens Sm = -349
	SmNoMoresRsrcs Sm = -344
	SmNosInfoArray Sm = -306
	SmOffsetErr Sm = -348
	SmPRAMInitErr Sm = -292
	SmPriInitErr Sm = -293
	SmRecNotFnd Sm = -351
	SmRevisionErr Sm = -303
	SmSDMInitErr Sm = -290
	SmSRTInitErr Sm = -291
	SmSRTOvrFlErr Sm = -350
	SmSelOOBErr Sm = -338
	SmSlotOOBErr Sm = -337
	SmsGetDrvrErr Sm = -345
	SmsPointerNil Sm = -335
)

func (e Sm) String() string {
	switch e {
	case SmBadsPtrErr:
		return "SmBadsPtrErr"
	case SmBlkMoveErr:
		return "SmBlkMoveErr"
	case SmByteLanesErr:
		return "SmByteLanesErr"
	case SmCPUErr:
		return "SmCPUErr"
	case SmCRCFail:
		return "SmCRCFail"
	case SmCkStatusErr:
		return "SmCkStatusErr"
	case SmDisDrvrNamErr:
		return "SmDisDrvrNamErr"
	case SmDisabledSlot:
		return "SmDisabledSlot"
	case SmEmptySlot:
		return "SmEmptySlot"
	case SmFormatErr:
		return "SmFormatErr"
	case SmGetDrvrNamErr:
		return "SmGetDrvrNamErr"
	case SmNewPErr:
		return "SmNewPErr"
	case SmNilsBlockErr:
		return "SmNilsBlockErr"
	case SmNoDir:
		return "SmNoDir"
	case SmNoGoodOpens:
		return "SmNoGoodOpens"
	case SmNoMoresRsrcs:
		return "SmNoMoresRsrcs"
	case SmNosInfoArray:
		return "SmNosInfoArray"
	case SmOffsetErr:
		return "SmOffsetErr"
	case SmPRAMInitErr:
		return "SmPRAMInitErr"
	case SmPriInitErr:
		return "SmPriInitErr"
	case SmRecNotFnd:
		return "SmRecNotFnd"
	case SmRevisionErr:
		return "SmRevisionErr"
	case SmSDMInitErr:
		return "SmSDMInitErr"
	case SmSRTInitErr:
		return "SmSRTInitErr"
	case SmSRTOvrFlErr:
		return "SmSRTOvrFlErr"
	case SmSelOOBErr:
		return "SmSelOOBErr"
	case SmSlotOOBErr:
		return "SmSlotOOBErr"
	case SmsGetDrvrErr:
		return "SmsGetDrvrErr"
	case SmsPointerNil:
		return "SmsPointerNil"
	default:
		return fmt.Sprintf("Sm(%d)", e)
	}
}

type SmResrvErr int

const (
	SmBLFieldBad SmResrvErr = -309
	SmBadBoardId SmResrvErr = -319
	SmBadRefId SmResrvErr = -330
	SmBadsList SmResrvErr = -331
	SmBusErrTO SmResrvErr = -320
	SmCodeRevErr SmResrvErr = -333
	SmDisposePErr SmResrvErr = -312
	SmFHBlkDispErr SmResrvErr = -311
	SmFHBlockRdErr SmResrvErr = -310
	SmGetPRErr SmResrvErr = -314
	SmInitStatVErr SmResrvErr = -316
	SmInitTblVErr SmResrvErr = -317
	SmNoBoardId SmResrvErr = -315
	SmNoBoardSRsrc SmResrvErr = -313
	SmNoJmpTbl SmResrvErr = -318
	SmReservedErr SmResrvErr = -332
	SmReservedSlot SmResrvErr = -318
	SmResrvErrValue SmResrvErr = -307
	SmUnExBusErr SmResrvErr = -308
	SvDisabled SmResrvErr = -32640
	SvTempDisable SmResrvErr = -32768
)

func (e SmResrvErr) String() string {
	switch e {
	case SmBLFieldBad:
		return "SmBLFieldBad"
	case SmBadBoardId:
		return "SmBadBoardId"
	case SmBadRefId:
		return "SmBadRefId"
	case SmBadsList:
		return "SmBadsList"
	case SmBusErrTO:
		return "SmBusErrTO"
	case SmCodeRevErr:
		return "SmCodeRevErr"
	case SmDisposePErr:
		return "SmDisposePErr"
	case SmFHBlkDispErr:
		return "SmFHBlkDispErr"
	case SmFHBlockRdErr:
		return "SmFHBlockRdErr"
	case SmGetPRErr:
		return "SmGetPRErr"
	case SmInitStatVErr:
		return "SmInitStatVErr"
	case SmInitTblVErr:
		return "SmInitTblVErr"
	case SmNoBoardId:
		return "SmNoBoardId"
	case SmNoBoardSRsrc:
		return "SmNoBoardSRsrc"
	case SmNoJmpTbl:
		return "SmNoJmpTbl"
	case SmReservedErr:
		return "SmReservedErr"
	case SmResrvErrValue:
		return "SmResrvErrValue"
	case SmUnExBusErr:
		return "SmUnExBusErr"
	case SvDisabled:
		return "SvDisabled"
	case SvTempDisable:
		return "SvTempDisable"
	default:
		return fmt.Sprintf("SmResrvErr(%d)", e)
	}
}

type StartupFolderIconResource int

const (
	ControlPanelFolderIconResource StartupFolderIconResource = -3976
	DropFolderIconResource StartupFolderIconResource = -3979
	ExtensionsFolderIconResource StartupFolderIconResource = -3973
	FontsFolderIconResource StartupFolderIconResource = -3968
	FullTrashIconResource StartupFolderIconResource = -3984
	MountedFolderIconResource StartupFolderIconResource = -3977
	OwnedFolderIconResource StartupFolderIconResource = -3980
	PreferencesFolderIconResource StartupFolderIconResource = -3974
	PrintMonitorFolderIconResource StartupFolderIconResource = -3975
	SharedFolderIconResource StartupFolderIconResource = -3978
	StartupFolderIconResourceValue StartupFolderIconResource = -3981
)

func (e StartupFolderIconResource) String() string {
	switch e {
	case ControlPanelFolderIconResource:
		return "ControlPanelFolderIconResource"
	case DropFolderIconResource:
		return "DropFolderIconResource"
	case ExtensionsFolderIconResource:
		return "ExtensionsFolderIconResource"
	case FontsFolderIconResource:
		return "FontsFolderIconResource"
	case FullTrashIconResource:
		return "FullTrashIconResource"
	case MountedFolderIconResource:
		return "MountedFolderIconResource"
	case OwnedFolderIconResource:
		return "OwnedFolderIconResource"
	case PreferencesFolderIconResource:
		return "PreferencesFolderIconResource"
	case PrintMonitorFolderIconResource:
		return "PrintMonitorFolderIconResource"
	case SharedFolderIconResource:
		return "SharedFolderIconResource"
	case StartupFolderIconResourceValue:
		return "StartupFolderIconResourceValue"
	default:
		return fmt.Sprintf("StartupFolderIconResource(%d)", e)
	}
}

type TeScrapSizeErr int

const (
	DriverHardwareGoneErr TeScrapSizeErr = -503
	HwParamErr TeScrapSizeErr = -502
	TeScrapSizeErrValue TeScrapSizeErr = -501
)

func (e TeScrapSizeErr) String() string {
	switch e {
	case DriverHardwareGoneErr:
		return "DriverHardwareGoneErr"
	case HwParamErr:
		return "HwParamErr"
	case TeScrapSizeErrValue:
		return "TeScrapSizeErrValue"
	default:
		return fmt.Sprintf("TeScrapSizeErr(%d)", e)
	}
}

type Tel int

const (
	TelAPattNotSupp Tel = -10016
	TelAlreadyOpen Tel = -10070
	TelAutoAnsNotOn Tel = -10112
	TelBadAPattErr Tel = -10015
	TelBadBearerType Tel = -10058
	TelBadCAErr Tel = -10003
	TelBadCodeResource Tel = -10108
	TelBadDNDType Tel = -10023
	TelBadDNErr Tel = -10002
	TelBadDNType Tel = -10050
	TelBadDisplayMode Tel = -10062
	TelBadFeatureID Tel = -10053
	TelBadFunction Tel = -10091
	TelBadFwdType Tel = -10054
	TelBadHTypeErr Tel = -10010
	TelBadHandErr Tel = -10004
	TelBadIndex Tel = -10017
	TelBadIntExt Tel = -10021
	TelBadIntercomID Tel = -10052
	TelBadLevelErr Tel = -10012
	TelBadPageID Tel = -10051
	TelBadParkID Tel = -10056
	TelBadPickupGroupID Tel = -10055
	TelBadProcErr Tel = -10005
	TelBadProcID Tel = -10110
	TelBadRate Tel = -10059
	TelBadSWErr Tel = -10114
	TelBadSampleRate Tel = -10115
	TelBadSelect Tel = -10057
	TelBadStateErr Tel = -10019
	TelBadTermErr Tel = -10001
	TelBadVTypeErr Tel = -10013
	TelCANotAcceptable Tel = -10080
	TelCANotDeflectable Tel = -10082
	TelCANotRejectable Tel = -10081
	TelCAUnavail Tel = -10006
	TelCBErr Tel = -10046
	TelConfErr Tel = -10042
	TelConfLimitErr Tel = -10040
	TelConfLimitExceeded Tel = -10047
	TelConfNoLimit Tel = -10041
	TelConfRej Tel = -10043
	TelDNDTypeNotSupp Tel = -10024
	TelDNTypeNotSupp Tel = -10060
	TelDetAlreadyOn Tel = -10113
	TelDeviceNotFound Tel = -10109
	TelDisplayModeNotSupp Tel = -10063
	TelFeatActive Tel = -10032
	TelFeatNotAvail Tel = -10031
	TelFeatNotSub Tel = -10030
	TelFeatNotSupp Tel = -10033
	TelFwdTypeNotSupp Tel = -10061
	TelGenericError Tel = -1
	TelHTypeNotSupp Tel = -10011
	TelIndexNotSupp Tel = -10018
	TelInitFailed Tel = -10107
	TelIntExtNotSupp Tel = -10022
	TelNoCallbackRef Tel = -10064
	TelNoCommFolder Tel = -10106
	TelNoErr Tel = 0
	TelNoMemErr Tel = -10007
	TelNoOpenErr Tel = -10008
	TelNoSuchTool Tel = -10102
	TelNoTools Tel = 8
	TelNotEnoughdspBW Tel = -10116
	TelPBErr Tel = -10090
	TelStateNotSupp Tel = -10020
	TelStillNeeded Tel = -10071
	TelTermNotOpen Tel = -10072
	TelTransferErr Tel = -10044
	TelTransferRej Tel = -10045
	TelUnknownErr Tel = -10103
	TelVTypeNotSupp Tel = -10014
	TelValidateFailed Tel = -10111
)

func (e Tel) String() string {
	switch e {
	case TelAPattNotSupp:
		return "TelAPattNotSupp"
	case TelAlreadyOpen:
		return "TelAlreadyOpen"
	case TelAutoAnsNotOn:
		return "TelAutoAnsNotOn"
	case TelBadAPattErr:
		return "TelBadAPattErr"
	case TelBadBearerType:
		return "TelBadBearerType"
	case TelBadCAErr:
		return "TelBadCAErr"
	case TelBadCodeResource:
		return "TelBadCodeResource"
	case TelBadDNDType:
		return "TelBadDNDType"
	case TelBadDNErr:
		return "TelBadDNErr"
	case TelBadDNType:
		return "TelBadDNType"
	case TelBadDisplayMode:
		return "TelBadDisplayMode"
	case TelBadFeatureID:
		return "TelBadFeatureID"
	case TelBadFunction:
		return "TelBadFunction"
	case TelBadFwdType:
		return "TelBadFwdType"
	case TelBadHTypeErr:
		return "TelBadHTypeErr"
	case TelBadHandErr:
		return "TelBadHandErr"
	case TelBadIndex:
		return "TelBadIndex"
	case TelBadIntExt:
		return "TelBadIntExt"
	case TelBadIntercomID:
		return "TelBadIntercomID"
	case TelBadLevelErr:
		return "TelBadLevelErr"
	case TelBadPageID:
		return "TelBadPageID"
	case TelBadParkID:
		return "TelBadParkID"
	case TelBadPickupGroupID:
		return "TelBadPickupGroupID"
	case TelBadProcErr:
		return "TelBadProcErr"
	case TelBadProcID:
		return "TelBadProcID"
	case TelBadRate:
		return "TelBadRate"
	case TelBadSWErr:
		return "TelBadSWErr"
	case TelBadSampleRate:
		return "TelBadSampleRate"
	case TelBadSelect:
		return "TelBadSelect"
	case TelBadStateErr:
		return "TelBadStateErr"
	case TelBadTermErr:
		return "TelBadTermErr"
	case TelBadVTypeErr:
		return "TelBadVTypeErr"
	case TelCANotAcceptable:
		return "TelCANotAcceptable"
	case TelCANotDeflectable:
		return "TelCANotDeflectable"
	case TelCANotRejectable:
		return "TelCANotRejectable"
	case TelCAUnavail:
		return "TelCAUnavail"
	case TelCBErr:
		return "TelCBErr"
	case TelConfErr:
		return "TelConfErr"
	case TelConfLimitErr:
		return "TelConfLimitErr"
	case TelConfLimitExceeded:
		return "TelConfLimitExceeded"
	case TelConfNoLimit:
		return "TelConfNoLimit"
	case TelConfRej:
		return "TelConfRej"
	case TelDNDTypeNotSupp:
		return "TelDNDTypeNotSupp"
	case TelDNTypeNotSupp:
		return "TelDNTypeNotSupp"
	case TelDetAlreadyOn:
		return "TelDetAlreadyOn"
	case TelDeviceNotFound:
		return "TelDeviceNotFound"
	case TelDisplayModeNotSupp:
		return "TelDisplayModeNotSupp"
	case TelFeatActive:
		return "TelFeatActive"
	case TelFeatNotAvail:
		return "TelFeatNotAvail"
	case TelFeatNotSub:
		return "TelFeatNotSub"
	case TelFeatNotSupp:
		return "TelFeatNotSupp"
	case TelFwdTypeNotSupp:
		return "TelFwdTypeNotSupp"
	case TelGenericError:
		return "TelGenericError"
	case TelHTypeNotSupp:
		return "TelHTypeNotSupp"
	case TelIndexNotSupp:
		return "TelIndexNotSupp"
	case TelInitFailed:
		return "TelInitFailed"
	case TelIntExtNotSupp:
		return "TelIntExtNotSupp"
	case TelNoCallbackRef:
		return "TelNoCallbackRef"
	case TelNoCommFolder:
		return "TelNoCommFolder"
	case TelNoErr:
		return "TelNoErr"
	case TelNoMemErr:
		return "TelNoMemErr"
	case TelNoOpenErr:
		return "TelNoOpenErr"
	case TelNoSuchTool:
		return "TelNoSuchTool"
	case TelNoTools:
		return "TelNoTools"
	case TelNotEnoughdspBW:
		return "TelNotEnoughdspBW"
	case TelPBErr:
		return "TelPBErr"
	case TelStateNotSupp:
		return "TelStateNotSupp"
	case TelStillNeeded:
		return "TelStillNeeded"
	case TelTermNotOpen:
		return "TelTermNotOpen"
	case TelTransferErr:
		return "TelTransferErr"
	case TelTransferRej:
		return "TelTransferRej"
	case TelUnknownErr:
		return "TelUnknownErr"
	case TelVTypeNotSupp:
		return "TelVTypeNotSupp"
	case TelValidateFailed:
		return "TelValidateFailed"
	default:
		return fmt.Sprintf("Tel(%d)", e)
	}
}

type TextParser int

const (
	TextParserBadParamErr TextParser = -5220
	TextParserBadParserObjectErr TextParser = -5223
	TextParserBadTextEncodingErr TextParser = -5227
	TextParserBadTextLanguageErr TextParser = -5226
	TextParserBadTokenValueErr TextParser = -5222
	TextParserNoMoreTextErr TextParser = -5225
	TextParserNoMoreTokensErr TextParser = -5229
	TextParserNoSuchTokenFoundErr TextParser = -5228
	TextParserObjectNotFoundErr TextParser = -5221
	TextParserParamErr TextParser = -5224
)

func (e TextParser) String() string {
	switch e {
	case TextParserBadParamErr:
		return "TextParserBadParamErr"
	case TextParserBadParserObjectErr:
		return "TextParserBadParserObjectErr"
	case TextParserBadTextEncodingErr:
		return "TextParserBadTextEncodingErr"
	case TextParserBadTextLanguageErr:
		return "TextParserBadTextLanguageErr"
	case TextParserBadTokenValueErr:
		return "TextParserBadTokenValueErr"
	case TextParserNoMoreTextErr:
		return "TextParserNoMoreTextErr"
	case TextParserNoMoreTokensErr:
		return "TextParserNoMoreTokensErr"
	case TextParserNoSuchTokenFoundErr:
		return "TextParserNoSuchTokenFoundErr"
	case TextParserObjectNotFoundErr:
		return "TextParserObjectNotFoundErr"
	case TextParserParamErr:
		return "TextParserParamErr"
	default:
		return fmt.Sprintf("TextParser(%d)", e)
	}
}

type Theme int

const (
	ThemeBadCursorIndexErr Theme = -30565
	ThemeBadTextColorErr Theme = -30563
	ThemeHasNoAccentsErr Theme = -30564
	ThemeInvalidBrushErr Theme = -30560
	ThemeMonitorDepthNotSupportedErr Theme = -30567
	ThemeNoAppropriateBrushErr Theme = -30568
	ThemeProcessNotRegisteredErr Theme = -30562
	ThemeProcessRegisteredErr Theme = -30561
	ThemeScriptFontNotFoundErr Theme = -30566
)

func (e Theme) String() string {
	switch e {
	case ThemeBadCursorIndexErr:
		return "ThemeBadCursorIndexErr"
	case ThemeBadTextColorErr:
		return "ThemeBadTextColorErr"
	case ThemeHasNoAccentsErr:
		return "ThemeHasNoAccentsErr"
	case ThemeInvalidBrushErr:
		return "ThemeInvalidBrushErr"
	case ThemeMonitorDepthNotSupportedErr:
		return "ThemeMonitorDepthNotSupportedErr"
	case ThemeNoAppropriateBrushErr:
		return "ThemeNoAppropriateBrushErr"
	case ThemeProcessNotRegisteredErr:
		return "ThemeProcessNotRegisteredErr"
	case ThemeProcessRegisteredErr:
		return "ThemeProcessRegisteredErr"
	case ThemeScriptFontNotFoundErr:
		return "ThemeScriptFontNotFoundErr"
	default:
		return fmt.Sprintf("Theme(%d)", e)
	}
}

type Thread int

const (
	ThreadNotFoundErr Thread = -618
	ThreadProtocolErr Thread = -619
	ThreadTooManyReqsErr Thread = -617
)

func (e Thread) String() string {
	switch e {
	case ThreadNotFoundErr:
		return "ThreadNotFoundErr"
	case ThreadProtocolErr:
		return "ThreadProtocolErr"
	case ThreadTooManyReqsErr:
		return "ThreadTooManyReqsErr"
	default:
		return fmt.Sprintf("Thread(%d)", e)
	}
}

const ThreadBadAppContextErr int = -616

type Tsm int

const (
	TsmAlreadyRegisteredErr Tsm = -2503
	TsmCantChangeForcedClassStateErr Tsm = -2530
	TsmCantOpenComponentErr Tsm = -2509
	TsmComponentAlreadyOpenErr Tsm = -2515
	TsmComponentNoErr Tsm = 0
	TsmComponentPropertyNotFoundErr Tsm = -2532
	TsmComponentPropertyUnsupportedErr Tsm = -2531
	TsmDefaultIsNotInputMethodErr Tsm = -2524
	TsmDocNotActiveErr Tsm = -2507
	TsmDocPropertyBufferTooSmallErr Tsm = -2529
	TsmDocPropertyNotFoundErr Tsm = -2528
	TsmDocumentOpenErr Tsm = -2511
	TsmInputMethodIsOldErr Tsm = -2516
	TsmInputMethodNotFoundErr Tsm = -2501
	TsmInputModeChangeFailedErr Tsm = -2533
	TsmInvalidContext Tsm = -2520
	TsmInvalidDocIDErr Tsm = -2505
	TsmNeverRegisteredErr Tsm = -2504
	TsmNoHandler Tsm = -2521
	TsmNoMoreTokens Tsm = -2522
	TsmNoOpenTSErr Tsm = -2508
	TsmNoStem Tsm = -2523
	TsmNotAnAppErr Tsm = -2502
	TsmScriptHasNoIMErr Tsm = -2517
	TsmTSHasNoMenuErr Tsm = -2513
	TsmTSMDocBusyErr Tsm = -2506
	TsmTSNotOpenErr Tsm = -2514
	TsmTextServiceNotFoundErr Tsm = -2510
	TsmUnknownErr Tsm = -2519
	TsmUnsupScriptLanguageErr Tsm = -2500
	TsmUnsupportedTypeErr Tsm = -2518
	TsmUseInputWindowErr Tsm = -2512
)

func (e Tsm) String() string {
	switch e {
	case TsmAlreadyRegisteredErr:
		return "TsmAlreadyRegisteredErr"
	case TsmCantChangeForcedClassStateErr:
		return "TsmCantChangeForcedClassStateErr"
	case TsmCantOpenComponentErr:
		return "TsmCantOpenComponentErr"
	case TsmComponentAlreadyOpenErr:
		return "TsmComponentAlreadyOpenErr"
	case TsmComponentNoErr:
		return "TsmComponentNoErr"
	case TsmComponentPropertyNotFoundErr:
		return "TsmComponentPropertyNotFoundErr"
	case TsmComponentPropertyUnsupportedErr:
		return "TsmComponentPropertyUnsupportedErr"
	case TsmDefaultIsNotInputMethodErr:
		return "TsmDefaultIsNotInputMethodErr"
	case TsmDocNotActiveErr:
		return "TsmDocNotActiveErr"
	case TsmDocPropertyBufferTooSmallErr:
		return "TsmDocPropertyBufferTooSmallErr"
	case TsmDocPropertyNotFoundErr:
		return "TsmDocPropertyNotFoundErr"
	case TsmDocumentOpenErr:
		return "TsmDocumentOpenErr"
	case TsmInputMethodIsOldErr:
		return "TsmInputMethodIsOldErr"
	case TsmInputMethodNotFoundErr:
		return "TsmInputMethodNotFoundErr"
	case TsmInputModeChangeFailedErr:
		return "TsmInputModeChangeFailedErr"
	case TsmInvalidContext:
		return "TsmInvalidContext"
	case TsmInvalidDocIDErr:
		return "TsmInvalidDocIDErr"
	case TsmNeverRegisteredErr:
		return "TsmNeverRegisteredErr"
	case TsmNoHandler:
		return "TsmNoHandler"
	case TsmNoMoreTokens:
		return "TsmNoMoreTokens"
	case TsmNoOpenTSErr:
		return "TsmNoOpenTSErr"
	case TsmNoStem:
		return "TsmNoStem"
	case TsmNotAnAppErr:
		return "TsmNotAnAppErr"
	case TsmScriptHasNoIMErr:
		return "TsmScriptHasNoIMErr"
	case TsmTSHasNoMenuErr:
		return "TsmTSHasNoMenuErr"
	case TsmTSMDocBusyErr:
		return "TsmTSMDocBusyErr"
	case TsmTSNotOpenErr:
		return "TsmTSNotOpenErr"
	case TsmTextServiceNotFoundErr:
		return "TsmTextServiceNotFoundErr"
	case TsmUnknownErr:
		return "TsmUnknownErr"
	case TsmUnsupScriptLanguageErr:
		return "TsmUnsupScriptLanguageErr"
	case TsmUnsupportedTypeErr:
		return "TsmUnsupportedTypeErr"
	case TsmUseInputWindowErr:
		return "TsmUseInputWindowErr"
	default:
		return fmt.Sprintf("Tsm(%d)", e)
	}
}

type Type uint

const (
	// Type128BitFloatingPoint: 128-bit floating point value.
	Type128BitFloatingPoint Type = 'l'<<24 | 'd'<<16 | 'b'<<8 | 'l' // 'ldbl'
	// TypeAEList: List of descriptors.
	TypeAEList Type = 'l'<<24 | 'i'<<16 | 's'<<8 | 't' // 'list'
	// TypeAERecord: List of keyword-specified descriptors.
	TypeAERecord Type = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'o' // 'reco'
	// TypeAbsoluteOrdinal: Specifies a descriptor whose data consists of one of the constants `kAEFirst`, `kAEMiddle`, `kAELast`, `kAEAny`, or `kAEAll`, which are described in AEDisposeToken(_:).
	TypeAbsoluteOrdinal Type = 'a'<<24 | 'b'<<16 | 's'<<8 | 'o' // 'abso'
	// TypeAlias: Alias.
	TypeAlias Type = 'a'<<24 | 'l'<<16 | 'i'<<8 | 's' // 'alis'
	// TypeAppParameters: Process Manager launch parameters.
	TypeAppParameters Type = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'a' // 'appa'
	// TypeApplSignature: Application signature.
	TypeApplSignature Type = 's'<<24 | 'i'<<16 | 'g'<<8 | 'n' // 'sign'
	// TypeAppleEvent: Apple event.
	TypeAppleEvent Type = 'a'<<24 | 'e'<<16 | 'v'<<8 | 't' // 'aevt'
	// TypeApplicationURL: For specifying an application by URL.
	TypeApplicationURL Type = 'a'<<24 | 'p'<<16 | 'r'<<8 | 'l' // 'aprl'
	TypeBookmarkData Type = 'b'<<24 | 'm'<<16 | 'r'<<8 | 'k' // 'bmrk'
	// TypeBoolean: Boolean value—single byte with value 0 or 1.
	TypeBoolean Type = 'b'<<24 | 'o'<<16 | 'o'<<8 | 'l' // 'bool'
	TypeCFAbsoluteTime Type = 'c'<<24 | 'f'<<16 | 'a'<<8 | 't' // 'cfat'
	// TypeCString: C string—Mac OS Roman characters followed by a NULL byte.
	TypeCString Type = 'c'<<24 | 's'<<16 | 't'<<8 | 'r' // 'cstr'
	// TypeChar: # Discussion
	TypeChar Type = 'T'<<24 | 'E'<<16 | 'X'<<8 | 'T' // 'TEXT'
	// TypeCompDescriptor: Specifies a comparison descriptor.
	TypeCompDescriptor Type = 'c'<<24 | 'm'<<16 | 'p'<<8 | 'd' // 'cmpd'
	// TypeCurrentContainer: Specifies a container for an element that demarcates one boundary in a range.
	TypeCurrentContainer Type = 'c'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'ccnt'
	// TypeDecimalStruct: Decimal.
	TypeDecimalStruct Type = 'd'<<24 | 'e'<<16 | 'c'<<8 | 'm' // 'decm'
	// TypeEncodedString: Styled Unicode text.
	TypeEncodedString Type = 'e'<<24 | 'n'<<16 | 'c'<<8 | 's' // 'encs'
	// TypeEnumerated: Enumerated data.
	TypeEnumerated Type = 'e'<<24 | 'n'<<16 | 'u'<<8 | 'm' // 'enum'
	TypeEventRecord Type = 'e'<<24 | 'v'<<16 | 'r'<<8 | 'c' // 'evrc'
	// TypeFSRef: File system reference.
	TypeFSRef Type = 'f'<<24 | 's'<<16 | 'r'<<8 | 'f' // 'fsrf'
	// TypeFalse: [FALSE] Boolean value.
	TypeFalse Type = 'f'<<24 | 'a'<<16 | 'l'<<8 | 's' // 'fals'
	// TypeFileURL: A file URL.
	TypeFileURL Type = 'f'<<24 | 'u'<<16 | 'r'<<8 | 'l' // 'furl'
	TypeFinderWindow Type = 'f'<<24 | 'w'<<16 | 'i'<<8 | 'n' // 'fwin'
	TypeFixed Type = 'f'<<24 | 'i'<<16 | 'x'<<8 | 'd' // 'fixd'
	TypeFixedPoint Type = 'f'<<24 | 'p'<<16 | 'n'<<8 | 't' // 'fpnt'
	TypeFixedRectangle Type = 'f'<<24 | 'r'<<16 | 'c'<<8 | 't' // 'frct'
	TypeGraphicLine Type = 'g'<<24 | 'l'<<16 | 'i'<<8 | 'n' // 'glin'
	TypeGraphicText Type = 'c'<<24 | 'g'<<16 | 't'<<8 | 'x' // 'cgtx'
	TypeGroupedGraphic Type = 'c'<<24 | 'p'<<16 | 'i'<<8 | 'c' // 'cpic'
	// TypeIEEE32BitFloatingPoint: 32-bit floating point value.
	TypeIEEE32BitFloatingPoint Type = 's'<<24 | 'i'<<16 | 'n'<<8 | 'g' // 'sing'
	// TypeIEEE64BitFloatingPoint: 64-bit floating point value.
	TypeIEEE64BitFloatingPoint Type = 'd'<<24 | 'o'<<16 | 'u'<<8 | 'b' // 'doub'
	TypeISO8601DateTime Type = 'i'<<24 | 's'<<16 | 'o'<<8 | 't' // 'isot'
	// TypeIndexDescriptor: Specifies a descriptor whose data indicates an indexed position within a range of values.
	TypeIndexDescriptor Type = 'i'<<24 | 'n'<<16 | 'd'<<8 | 'e' // 'inde'
	TypeInsertionLoc Type = 'i'<<24 | 'n'<<16 | 's'<<8 | 'l' // 'insl'
	// TypeIntlText: For important information, see the Version Notes section of the typeUnicodeText enum.
	TypeIntlText Type = 'i'<<24 | 't'<<16 | 'x'<<8 | 't' // 'itxt'
	TypeIntlWritingCode Type = 'i'<<24 | 'n'<<16 | 't'<<8 | 'l' // 'intl'
	// TypeKeyword: Apple event keyword.
	TypeKeyword Type = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'w' // 'keyw'
	// TypeLogicalDescriptor: Specifies a logical descriptor.
	TypeLogicalDescriptor Type = 'l'<<24 | 'o'<<16 | 'g'<<8 | 'i' // 'logi'
	TypeLongDateTime Type = 'l'<<24 | 'd'<<16 | 't'<<8 | ' ' // 'ldt '
	TypeLongFixed Type = 'l'<<24 | 'f'<<16 | 'x'<<8 | 'd' // 'lfxd'
	TypeLongFixedPoint Type = 'l'<<24 | 'f'<<16 | 'p'<<8 | 't' // 'lfpt'
	TypeLongFixedRectangle Type = 'l'<<24 | 'f'<<16 | 'r'<<8 | 'c' // 'lfrc'
	TypeLongPoint Type = 'l'<<24 | 'p'<<16 | 'n'<<8 | 't' // 'lpnt'
	TypeLongRectangle Type = 'l'<<24 | 'r'<<16 | 'c'<<8 | 't' // 'lrct'
	TypeMachineLoc Type = 'm'<<24 | 'L'<<16 | 'o'<<8 | 'c' // 'mLoc'
	// TypeNull: A null data storage pointer.
	TypeNull Type = 'n'<<24 | 'u'<<16 | 'l'<<8 | 'l' // 'null'
	// TypeOSLTokenList: Specifies a descriptor whose data consists of a list of tokens.
	TypeOSLTokenList Type = 'o'<<24 | 's'<<16 | 't'<<8 | 'l' // 'ostl'
	// TypeObjectBeingExamined: # Discussion
	TypeObjectBeingExamined Type = 'e'<<24 | 'x'<<16 | 'm'<<8 | 'n' // 'exmn'
	// TypeObjectSpecifier: Specifies a descriptor used with the `keyAEContainer` keyword in a keyword-specified descriptor.
	TypeObjectSpecifier Type = 'o'<<24 | 'b'<<16 | 'j'<<8 | ' ' // 'obj '
	TypeOval Type = 'c'<<24 | 'o'<<16 | 'v'<<8 | 'l' // 'covl'
	// TypePString: Pascal string—unsigned length byte followed by Mac OS Roman characters.
	TypePString Type = 'p'<<24 | 's'<<16 | 't'<<8 | 'r' // 'pstr'
	TypeParamInfo Type = 'p'<<24 | 'm'<<16 | 'i'<<8 | 'n' // 'pmin'
	TypePict Type = 'P'<<24 | 'I'<<16 | 'C'<<8 | 'T' // 'PICT'
	TypePixMapMinus Type = 't'<<24 | 'p'<<16 | 'm'<<8 | 'm' // 'tpmm'
	TypePixelMap Type = 'c'<<24 | 'p'<<16 | 'i'<<8 | 'x' // 'cpix'
	TypePolygon Type = 'c'<<24 | 'p'<<16 | 'g'<<8 | 'n' // 'cpgn'
	// TypeProcessSerialNumber: A process serial number.
	TypeProcessSerialNumber Type = 'p'<<24 | 's'<<16 | 'n'<<8 | ' ' // 'psn '
	TypePropInfo Type = 'p'<<24 | 'i'<<16 | 'n'<<8 | 'f' // 'pinf'
	// TypeProperty: Apple event object property.
	TypeProperty Type = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
	TypePtr Type = 'p'<<24 | 't'<<16 | 'r'<<8 | ' ' // 'ptr '
	TypeQDPoint Type = 'Q'<<24 | 'D'<<16 | 'p'<<8 | 't' // 'QDpt'
	TypeQDRectangle Type = 'q'<<24 | 'd'<<16 | 'r'<<8 | 't' // 'qdrt'
	TypeQDRegion Type = 'Q'<<24 | 'r'<<16 | 'g'<<8 | 'n' // 'Qrgn'
	TypeRGB16 Type = 't'<<24 | 'r'<<16 | '1'<<8 | '6' // 'tr16'
	TypeRGB96 Type = 't'<<24 | 'r'<<16 | '9'<<8 | '6' // 'tr96'
	TypeRGBColor Type = 'c'<<24 | 'R'<<16 | 'G'<<8 | 'B' // 'cRGB'
	// TypeRangeDescriptor: # Discussion
	TypeRangeDescriptor Type = 'r'<<24 | 'a'<<16 | 'n'<<8 | 'g' // 'rang'
	TypeRectangle Type = 'c'<<24 | 'r'<<16 | 'e'<<8 | 'c' // 'crec'
	// TypeRelativeDescriptor: Specifies a descriptor whose data consists of one of the constants `kAENext` or `kAEPrevious`, which are described in AEDisposeToken(_:).
	TypeRelativeDescriptor Type = 'r'<<24 | 'e'<<16 | 'l'<<8 | ' ' // 'rel '
	TypeRotation Type = 't'<<24 | 'r'<<16 | 'o'<<8 | 't' // 'trot'
	TypeRoundedRectangle Type = 'c'<<24 | 'r'<<16 | 'r'<<8 | 'c' // 'crrc'
	TypeRow Type = 'c'<<24 | 'r'<<16 | 'o'<<8 | 'w' // 'crow'
	// TypeSInt16: 16-bit signed integer.
	TypeSInt16 Type = 's'<<24 | 'h'<<16 | 'o'<<8 | 'r' // 'shor'
	// TypeSInt32: 32-bit signed integer.
	TypeSInt32 Type = 'l'<<24 | 'o'<<16 | 'n'<<8 | 'g' // 'long'
	// TypeSInt64: 64-bit signed integer.
	TypeSInt64 Type = 'c'<<24 | 'o'<<16 | 'm'<<8 | 'p' // 'comp'
	TypeScrapStyles Type = 's'<<24 | 't'<<16 | 'y'<<8 | 'l' // 'styl'
	TypeScript Type = 's'<<24 | 'c'<<16 | 'p'<<8 | 't' // 'scpt'
	// TypeSectionH: Handle to a section record.
	TypeSectionH Type = 's'<<24 | 'e'<<16 | 'c'<<8 | 't' // 'sect'
	// TypeStyledText: # Discussion
	TypeStyledText Type = 'S'<<24 | 'T'<<16 | 'X'<<8 | 'T' // 'STXT'
	// TypeStyledUnicodeText: Styled Unicode text.
	TypeStyledUnicodeText Type = 's'<<24 | 'u'<<16 | 't'<<8 | 'x' // 'sutx'
	TypeSuiteInfo Type = 's'<<24 | 'u'<<16 | 'i'<<8 | 'n' // 'suin'
	TypeTable Type = 'c'<<24 | 't'<<16 | 'b'<<8 | 'l' // 'ctbl'
	TypeTextStyles Type = 't'<<24 | 's'<<16 | 't'<<8 | 'y' // 'tsty'
	// TypeToken: Specifies a descriptor whose data storage pointer refers to a structure of type AEDisposeToken(_:).
	TypeToken Type = 't'<<24 | 'o'<<16 | 'k'<<8 | 'e' // 'toke'
	// TypeTrue: [TRUE] Boolean value.
	TypeTrue Type = 't'<<24 | 'r'<<16 | 'u'<<8 | 'e' // 'true'
	// TypeType: Four-character code for event class or event ID
	TypeType Type = 't'<<24 | 'y'<<16 | 'p'<<8 | 'e' // 'type'
	// TypeUInt16: 16-bit unsigned integer.
	TypeUInt16 Type = 'u'<<24 | 's'<<16 | 'h'<<8 | 'r' // 'ushr'
	// TypeUInt32: 32-bit unsigned integer.
	TypeUInt32 Type = 'm'<<24 | 'a'<<16 | 'g'<<8 | 'n' // 'magn'
	// TypeUInt64: 64-bit unsigned integer.
	TypeUInt64 Type = 'u'<<24 | 'c'<<16 | 'o'<<8 | 'm' // 'ucom'
	// TypeUnicodeText: Unicode text.
	TypeUnicodeText Type = 'u'<<24 | 't'<<16 | 'x'<<8 | 't' // 'utxt'
	// TypeWildCard: Matches any type.
	TypeWildCard Type = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
)

func (e Type) String() string {
	switch e {
	case Type128BitFloatingPoint:
		return "Type128BitFloatingPoint"
	case TypeAEList:
		return "TypeAEList"
	case TypeAERecord:
		return "TypeAERecord"
	case TypeAbsoluteOrdinal:
		return "TypeAbsoluteOrdinal"
	case TypeAlias:
		return "TypeAlias"
	case TypeAppParameters:
		return "TypeAppParameters"
	case TypeApplSignature:
		return "TypeApplSignature"
	case TypeAppleEvent:
		return "TypeAppleEvent"
	case TypeApplicationURL:
		return "TypeApplicationURL"
	case TypeBookmarkData:
		return "TypeBookmarkData"
	case TypeBoolean:
		return "TypeBoolean"
	case TypeCFAbsoluteTime:
		return "TypeCFAbsoluteTime"
	case TypeCString:
		return "TypeCString"
	case TypeChar:
		return "TypeChar"
	case TypeCompDescriptor:
		return "TypeCompDescriptor"
	case TypeCurrentContainer:
		return "TypeCurrentContainer"
	case TypeDecimalStruct:
		return "TypeDecimalStruct"
	case TypeEncodedString:
		return "TypeEncodedString"
	case TypeEnumerated:
		return "TypeEnumerated"
	case TypeEventRecord:
		return "TypeEventRecord"
	case TypeFSRef:
		return "TypeFSRef"
	case TypeFalse:
		return "TypeFalse"
	case TypeFileURL:
		return "TypeFileURL"
	case TypeFinderWindow:
		return "TypeFinderWindow"
	case TypeFixed:
		return "TypeFixed"
	case TypeFixedPoint:
		return "TypeFixedPoint"
	case TypeFixedRectangle:
		return "TypeFixedRectangle"
	case TypeGraphicLine:
		return "TypeGraphicLine"
	case TypeGraphicText:
		return "TypeGraphicText"
	case TypeGroupedGraphic:
		return "TypeGroupedGraphic"
	case TypeIEEE32BitFloatingPoint:
		return "TypeIEEE32BitFloatingPoint"
	case TypeIEEE64BitFloatingPoint:
		return "TypeIEEE64BitFloatingPoint"
	case TypeISO8601DateTime:
		return "TypeISO8601DateTime"
	case TypeIndexDescriptor:
		return "TypeIndexDescriptor"
	case TypeInsertionLoc:
		return "TypeInsertionLoc"
	case TypeIntlText:
		return "TypeIntlText"
	case TypeIntlWritingCode:
		return "TypeIntlWritingCode"
	case TypeKeyword:
		return "TypeKeyword"
	case TypeLogicalDescriptor:
		return "TypeLogicalDescriptor"
	case TypeLongDateTime:
		return "TypeLongDateTime"
	case TypeLongFixed:
		return "TypeLongFixed"
	case TypeLongFixedPoint:
		return "TypeLongFixedPoint"
	case TypeLongFixedRectangle:
		return "TypeLongFixedRectangle"
	case TypeLongPoint:
		return "TypeLongPoint"
	case TypeLongRectangle:
		return "TypeLongRectangle"
	case TypeMachineLoc:
		return "TypeMachineLoc"
	case TypeNull:
		return "TypeNull"
	case TypeOSLTokenList:
		return "TypeOSLTokenList"
	case TypeObjectBeingExamined:
		return "TypeObjectBeingExamined"
	case TypeObjectSpecifier:
		return "TypeObjectSpecifier"
	case TypeOval:
		return "TypeOval"
	case TypePString:
		return "TypePString"
	case TypeParamInfo:
		return "TypeParamInfo"
	case TypePict:
		return "TypePict"
	case TypePixMapMinus:
		return "TypePixMapMinus"
	case TypePixelMap:
		return "TypePixelMap"
	case TypePolygon:
		return "TypePolygon"
	case TypeProcessSerialNumber:
		return "TypeProcessSerialNumber"
	case TypePropInfo:
		return "TypePropInfo"
	case TypeProperty:
		return "TypeProperty"
	case TypePtr:
		return "TypePtr"
	case TypeQDPoint:
		return "TypeQDPoint"
	case TypeQDRectangle:
		return "TypeQDRectangle"
	case TypeQDRegion:
		return "TypeQDRegion"
	case TypeRGB16:
		return "TypeRGB16"
	case TypeRGB96:
		return "TypeRGB96"
	case TypeRGBColor:
		return "TypeRGBColor"
	case TypeRangeDescriptor:
		return "TypeRangeDescriptor"
	case TypeRectangle:
		return "TypeRectangle"
	case TypeRelativeDescriptor:
		return "TypeRelativeDescriptor"
	case TypeRotation:
		return "TypeRotation"
	case TypeRoundedRectangle:
		return "TypeRoundedRectangle"
	case TypeRow:
		return "TypeRow"
	case TypeSInt16:
		return "TypeSInt16"
	case TypeSInt32:
		return "TypeSInt32"
	case TypeSInt64:
		return "TypeSInt64"
	case TypeScrapStyles:
		return "TypeScrapStyles"
	case TypeScript:
		return "TypeScript"
	case TypeSectionH:
		return "TypeSectionH"
	case TypeStyledText:
		return "TypeStyledText"
	case TypeStyledUnicodeText:
		return "TypeStyledUnicodeText"
	case TypeSuiteInfo:
		return "TypeSuiteInfo"
	case TypeTable:
		return "TypeTable"
	case TypeTextStyles:
		return "TypeTextStyles"
	case TypeToken:
		return "TypeToken"
	case TypeTrue:
		return "TypeTrue"
	case TypeType:
		return "TypeType"
	case TypeUInt16:
		return "TypeUInt16"
	case TypeUInt32:
		return "TypeUInt32"
	case TypeUInt64:
		return "TypeUInt64"
	case TypeUnicodeText:
		return "TypeUnicodeText"
	case TypeWildCard:
		return "TypeWildCard"
	default:
		return fmt.Sprintf("Type(%d)", e)
	}
}

const TypeAuditToken uint = 't'<<24 | 'o'<<16 | 'k'<<8 | 'n' // 'tokn'

type TypeCF uint

const (
	TypeCFArrayRef TypeCF = 'c'<<24 | 'f'<<16 | 'a'<<8 | 'r' // 'cfar'
	TypeCFAttributedStringRef TypeCF = 'c'<<24 | 'f'<<16 | 'a'<<8 | 's' // 'cfas'
	TypeCFBooleanRef TypeCF = 'c'<<24 | 'f'<<16 | 't'<<8 | 'f' // 'cftf'
	TypeCFDictionaryRef TypeCF = 'c'<<24 | 'f'<<16 | 'd'<<8 | 'c' // 'cfdc'
	TypeCFMutableArrayRef TypeCF = 'c'<<24 | 'f'<<16 | 'm'<<8 | 'a' // 'cfma'
	TypeCFMutableAttributedStringRef TypeCF = 'c'<<24 | 'f'<<16 | 'a'<<8 | 'a' // 'cfaa'
	TypeCFMutableDictionaryRef TypeCF = 'c'<<24 | 'f'<<16 | 'm'<<8 | 'd' // 'cfmd'
	TypeCFMutableStringRef TypeCF = 'c'<<24 | 'f'<<16 | 'm'<<8 | 's' // 'cfms'
	TypeCFNumberRef TypeCF = 'c'<<24 | 'f'<<16 | 'n'<<8 | 'b' // 'cfnb'
	TypeCFStringRef TypeCF = 'c'<<24 | 'f'<<16 | 's'<<8 | 't' // 'cfst'
	TypeCFTypeRef TypeCF = 'c'<<24 | 'f'<<16 | 't'<<8 | 'y' // 'cfty'
)

func (e TypeCF) String() string {
	switch e {
	case TypeCFArrayRef:
		return "TypeCFArrayRef"
	case TypeCFAttributedStringRef:
		return "TypeCFAttributedStringRef"
	case TypeCFBooleanRef:
		return "TypeCFBooleanRef"
	case TypeCFDictionaryRef:
		return "TypeCFDictionaryRef"
	case TypeCFMutableArrayRef:
		return "TypeCFMutableArrayRef"
	case TypeCFMutableAttributedStringRef:
		return "TypeCFMutableAttributedStringRef"
	case TypeCFMutableDictionaryRef:
		return "TypeCFMutableDictionaryRef"
	case TypeCFMutableStringRef:
		return "TypeCFMutableStringRef"
	case TypeCFNumberRef:
		return "TypeCFNumberRef"
	case TypeCFStringRef:
		return "TypeCFStringRef"
	case TypeCFTypeRef:
		return "TypeCFTypeRef"
	default:
		return fmt.Sprintf("TypeCF(%d)", e)
	}
}

type TypeWhoseDescriptor uint

const (
	// FormWhose: # Discussion
	FormWhose TypeWhoseDescriptor = 'w'<<24 | 'h'<<16 | 'o'<<8 | 's' // 'whos'
	KeyAEIndex TypeWhoseDescriptor = 'k'<<24 | 'i'<<16 | 'd'<<8 | 'x' // 'kidx'
	KeyAETest TypeWhoseDescriptor = 'k'<<24 | 't'<<16 | 's'<<8 | 't' // 'ktst'
	KeyAEWhoseRangeStart TypeWhoseDescriptor = 'w'<<24 | 's'<<16 | 't'<<8 | 'r' // 'wstr'
	KeyAEWhoseRangeStop TypeWhoseDescriptor = 'w'<<24 | 's'<<16 | 't'<<8 | 'p' // 'wstp'
	TypeWhoseDescriptorValue TypeWhoseDescriptor = 'w'<<24 | 'h'<<16 | 'o'<<8 | 's' // 'whos'
	TypeWhoseRange TypeWhoseDescriptor = 'w'<<24 | 'r'<<16 | 'n'<<8 | 'g' // 'wrng'
)

func (e TypeWhoseDescriptor) String() string {
	switch e {
	case FormWhose:
		return "FormWhose"
	case KeyAEIndex:
		return "KeyAEIndex"
	case KeyAETest:
		return "KeyAETest"
	case KeyAEWhoseRangeStart:
		return "KeyAEWhoseRangeStart"
	case KeyAEWhoseRangeStop:
		return "KeyAEWhoseRangeStop"
	case TypeWhoseRange:
		return "TypeWhoseRange"
	default:
		return fmt.Sprintf("TypeWhoseDescriptor(%d)", e)
	}
}

type VLckdErr int

const (
	BadMDBErr VLckdErr = -60
	BadMovErr VLckdErr = -122
	DirNFErr VLckdErr = -120
	DupFNErr VLckdErr = -48
	ExtFSErr VLckdErr = -58
	FBsyErr VLckdErr = -47
	FsRnErr VLckdErr = -59
	GfpErr VLckdErr = -52
	NoMacDskErr VLckdErr = -57
	NsDrvErr VLckdErr = -56
	OpWrErr VLckdErr = -49
	PermErr VLckdErr = -54
	RfNumErr VLckdErr = -51
	TmwdoErr VLckdErr = -121
	VLckdErrValue VLckdErr = -46
	VolGoneErr VLckdErr = -124
	VolOffLinErr VLckdErr = -53
	VolOnLinErr VLckdErr = -55
	WrPermErr VLckdErr = -61
	WrgVolTypErr VLckdErr = -123
)

func (e VLckdErr) String() string {
	switch e {
	case BadMDBErr:
		return "BadMDBErr"
	case BadMovErr:
		return "BadMovErr"
	case DirNFErr:
		return "DirNFErr"
	case DupFNErr:
		return "DupFNErr"
	case ExtFSErr:
		return "ExtFSErr"
	case FBsyErr:
		return "FBsyErr"
	case FsRnErr:
		return "FsRnErr"
	case GfpErr:
		return "GfpErr"
	case NoMacDskErr:
		return "NoMacDskErr"
	case NsDrvErr:
		return "NsDrvErr"
	case OpWrErr:
		return "OpWrErr"
	case PermErr:
		return "PermErr"
	case RfNumErr:
		return "RfNumErr"
	case TmwdoErr:
		return "TmwdoErr"
	case VLckdErrValue:
		return "VLckdErrValue"
	case VolGoneErr:
		return "VolGoneErr"
	case VolOffLinErr:
		return "VolOffLinErr"
	case VolOnLinErr:
		return "VolOnLinErr"
	case WrPermErr:
		return "WrPermErr"
	case WrgVolTypErr:
		return "WrgVolTypErr"
	default:
		return fmt.Sprintf("VLckdErr(%d)", e)
	}
}

type Vm int

const (
	VmAddressNotInFileViewErr Vm = -647
	VmBadDriver Vm = -632
	VmBusyBackingFileErr Vm = -642
	VmFileViewAccessErr Vm = -645
	VmInvalidBackingFileIDErr Vm = -640
	VmInvalidFileViewIDErr Vm = -644
	VmInvalidOwningProcessErr Vm = -648
	VmKernelMMUInitErr Vm = -629
	VmMappingPrivilegesErr Vm = -641
	VmMemLckdErr Vm = -631
	VmMorePhysicalThanVirtualErr Vm = -628
	VmNoMoreBackingFilesErr Vm = -643
	VmNoMoreFileViewsErr Vm = -646
	VmNoVectorErr Vm = -633
	VmOffErr Vm = -630
)

func (e Vm) String() string {
	switch e {
	case VmAddressNotInFileViewErr:
		return "VmAddressNotInFileViewErr"
	case VmBadDriver:
		return "VmBadDriver"
	case VmBusyBackingFileErr:
		return "VmBusyBackingFileErr"
	case VmFileViewAccessErr:
		return "VmFileViewAccessErr"
	case VmInvalidBackingFileIDErr:
		return "VmInvalidBackingFileIDErr"
	case VmInvalidFileViewIDErr:
		return "VmInvalidFileViewIDErr"
	case VmInvalidOwningProcessErr:
		return "VmInvalidOwningProcessErr"
	case VmKernelMMUInitErr:
		return "VmKernelMMUInitErr"
	case VmMappingPrivilegesErr:
		return "VmMappingPrivilegesErr"
	case VmMemLckdErr:
		return "VmMemLckdErr"
	case VmMorePhysicalThanVirtualErr:
		return "VmMorePhysicalThanVirtualErr"
	case VmNoMoreBackingFilesErr:
		return "VmNoMoreBackingFilesErr"
	case VmNoMoreFileViewsErr:
		return "VmNoMoreFileViewsErr"
	case VmNoVectorErr:
		return "VmNoVectorErr"
	case VmOffErr:
		return "VmOffErr"
	default:
		return fmt.Sprintf("Vm(%d)", e)
	}
}

type WrongApplicationPlatform int

const (
	AppVersionTooOld WrongApplicationPlatform = -876
	NotAppropriateForClassic WrongApplicationPlatform = -877
	WrongApplicationPlatformValue WrongApplicationPlatform = -875
)

func (e WrongApplicationPlatform) String() string {
	switch e {
	case AppVersionTooOld:
		return "AppVersionTooOld"
	case NotAppropriateForClassic:
		return "NotAppropriateForClassic"
	case WrongApplicationPlatformValue:
		return "WrongApplicationPlatformValue"
	default:
		return fmt.Sprintf("WrongApplicationPlatform(%d)", e)
	}
}

