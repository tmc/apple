// Code generated from Apple documentation for coreservices. DO NOT EDIT.

package coreservices

import (
	"fmt"
)

type AbortErr int

const (
	AbortErrValue  AbortErr = -27
	BdNamErr       AbortErr = -37
	DceExtErr      AbortErr = -30
	DirFulErr      AbortErr = -33
	DskFulErr      AbortErr = -34
	EofErr         AbortErr = -39
	FLckdErr       AbortErr = -45
	FnOpnErr       AbortErr = -38
	FnfErr         AbortErr = -43
	GcrOnMFMErr    AbortErr = -400
	IIOAbortErr    AbortErr = -27
	IoErr          AbortErr = -36
	MFulErr        AbortErr = -41
	NotOpenErr     AbortErr = -28
	NsvErr         AbortErr = -35
	PosErr         AbortErr = -40
	SlotNumErr     AbortErr = -360
	TmfoErr        AbortErr = -42
	UnitTblFullErr AbortErr = -29
	WPrErr         AbortErr = -44
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
	AfpAccessDenied       Afp = -5000
	AfpAlreadyLoggedInErr Afp = -5047
	AfpAlreadyMounted     Afp = -5062
	AfpAuthContinue       Afp = -5001
	AfpBadDirIDType       Afp = -5060
	AfpBadIDErr           Afp = -5039
	AfpBadUAM             Afp = -5002
	AfpBadVersNum         Afp = -5003
	AfpBitmapErr          Afp = -5004
	AfpCallNotAllowed     Afp = -5048
	AfpCallNotSupported   Afp = -5024
	AfpCantMountMoreSrvre Afp = -5061
	AfpCantMove           Afp = -5005
	AfpCantRename         Afp = -5028
	AfpCatalogChanged     Afp = -5037
	AfpContainsSharedErr  Afp = -5033
	AfpDenyConflict       Afp = -5006
	AfpDiffVolErr         Afp = -5036
	AfpDirNotEmpty        Afp = -5007
	AfpDirNotFound        Afp = -5029
	AfpDiskFull           Afp = -5008
	AfpEofError           Afp = -5009
	AfpFileBusy           Afp = -5010
	AfpFlatVol            Afp = -5011
	AfpIDExists           Afp = -5035
	AfpIDNotFound         Afp = -5034
	AfpIconTypeError      Afp = -5030
	AfpInsideSharedErr    Afp = -5043
	AfpInsideTrashErr     Afp = -5044
	AfpItemNotFound       Afp = -5012
	AfpLockErr            Afp = -5013
	AfpMiscErr            Afp = -5014
	AfpNoMoreLocks        Afp = -5015
	AfpNoServer           Afp = -5016
	AfpObjectExists       Afp = -5017
	AfpObjectLocked       Afp = -5032
	AfpObjectNotFound     Afp = -5018
	AfpObjectTypeErr      Afp = -5025
	AfpParmErr            Afp = -5019
	AfpPwdExpiredErr      Afp = -5042
	AfpPwdNeedsChangeErr  Afp = -5045
	AfpPwdPolicyErr       Afp = -5046
	AfpPwdSameErr         Afp = -5040
	AfpPwdTooShortErr     Afp = -5041
	AfpRangeNotLocked     Afp = -5020
	AfpRangeOverlap       Afp = -5021
	AfpSameNodeErr        Afp = -5063
	AfpSameObjectErr      Afp = -5038
	AfpServerGoingDown    Afp = -5027
	AfpSessClosed         Afp = -5022
	AfpTooManyFilesOpen   Afp = -5026
	AfpUserNotAuth        Afp = -5023
	AfpVolLocked          Afp = -5031
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

type Aifc uint

const (
	AIFCVersion1 Aifc = 0
)

func (e Aifc) String() string {
	switch e {
	case AIFCVersion1:
		return "AIFCVersion1"
	default:
		return fmt.Sprintf("Aifc(%d)", e)
	}
}

type Aiffid uint

const (
	AIFCID                Aiffid = 'A'<<24 | 'I'<<16 | 'F'<<8 | 'C' // 'AIFC'
	AIFFID                Aiffid = 'A'<<24 | 'I'<<16 | 'F'<<8 | 'F' // 'AIFF'
	AnnotationID          Aiffid = 'A'<<24 | 'N'<<16 | 'N'<<8 | 'O' // 'ANNO'
	ApplicationSpecificID Aiffid = 'A'<<24 | 'P'<<16 | 'P'<<8 | 'L' // 'APPL'
	AudioRecordingID      Aiffid = 'A'<<24 | 'E'<<16 | 'S'<<8 | 'D' // 'AESD'
	AuthorID              Aiffid = 'A'<<24 | 'U'<<16 | 'T'<<8 | 'H' // 'AUTH'
	CommentID             Aiffid = 'C'<<24 | 'O'<<16 | 'M'<<8 | 'T' // 'COMT'
	CommonID              Aiffid = 'C'<<24 | 'O'<<16 | 'M'<<8 | 'M' // 'COMM'
	CopyrightID           Aiffid = '('<<24 | 'c'<<16 | ')'<<8 | ' ' // '(c) '
	FORMID                Aiffid = 'F'<<24 | 'O'<<16 | 'R'<<8 | 'M' // 'FORM'
	FormatVersionID       Aiffid = 'F'<<24 | 'V'<<16 | 'E'<<8 | 'R' // 'FVER'
	InstrumentID          Aiffid = 'I'<<24 | 'N'<<16 | 'S'<<8 | 'T' // 'INST'
	MIDIDataID            Aiffid = 'M'<<24 | 'I'<<16 | 'D'<<8 | 'I' // 'MIDI'
	MarkerID              Aiffid = 'M'<<24 | 'A'<<16 | 'R'<<8 | 'K' // 'MARK'
	NameID                Aiffid = 'N'<<24 | 'A'<<16 | 'M'<<8 | 'E' // 'NAME'
	SoundDataID           Aiffid = 'S'<<24 | 'S'<<16 | 'N'<<8 | 'D' // 'SSND'
)

func (e Aiffid) String() string {
	switch e {
	case AIFCID:
		return "AIFCID"
	case AIFFID:
		return "AIFFID"
	case AnnotationID:
		return "AnnotationID"
	case ApplicationSpecificID:
		return "ApplicationSpecificID"
	case AudioRecordingID:
		return "AudioRecordingID"
	case AuthorID:
		return "AuthorID"
	case CommentID:
		return "CommentID"
	case CommonID:
		return "CommonID"
	case CopyrightID:
		return "CopyrightID"
	case FORMID:
		return "FORMID"
	case FormatVersionID:
		return "FormatVersionID"
	case InstrumentID:
		return "InstrumentID"
	case MIDIDataID:
		return "MIDIDataID"
	case MarkerID:
		return "MarkerID"
	case NameID:
		return "NameID"
	case SoundDataID:
		return "SoundDataID"
	default:
		return fmt.Sprintf("Aiffid(%d)", e)
	}
}

type AppleShareMedia uint

const (
	AppleShareMediaType AppleShareMedia = 'a'<<24 | 'f'<<16 | 'p'<<8 | 'm' // 'afpm'
)

func (e AppleShareMedia) String() string {
	switch e {
	case AppleShareMediaType:
		return "AppleShareMediaType"
	default:
		return fmt.Sprintf("AppleShareMedia(%d)", e)
	}
}

type Asi int

const (
	// Deprecated.
	AsiAliasName Asi = 0
	// Deprecated.
	AsiParentName Asi = 1
	// Deprecated.
	AsiServerName Asi = -2
	// Deprecated.
	AsiVolumeName Asi = -1
	// Deprecated.
	AsiZoneName Asi = -3
)

func (e Asi) String() string {
	switch e {
	case AsiAliasName:
		return "AsiAliasName"
	case AsiParentName:
		return "AsiParentName"
	case AsiServerName:
		return "AsiServerName"
	case AsiVolumeName:
		return "AsiVolumeName"
	case AsiZoneName:
		return "AsiZoneName"
	default:
		return fmt.Sprintf("Asi(%d)", e)
	}
}

type Asp int

const (
	AspBadVersNum  Asp = -1066
	AspBufTooSmall Asp = -1067
	AspNoAck       Asp = -1075
	AspNoMoreSess  Asp = -1068
	AspNoServers   Asp = -1069
	AspParamErr    Asp = -1070
	AspServerBusy  Asp = -1071
	AspSessClosed  Asp = -1072
	AspSizeErr     Asp = -1073
	AspTooMany     Asp = -1074
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

type BHasDirectI uint

const (
	BHasDirectIO BHasDirectI = 1
)

func (e BHasDirectI) String() string {
	switch e {
	case BHasDirectIO:
		return "BHasDirectIO"
	default:
		return fmt.Sprintf("BHasDirectI(%d)", e)
	}
}

type BIsEjectable uint

const (
	BAllowCDiDataHandler          BIsEjectable = 17
	BAncestorModDateChanges       BIsEjectable = 11
	BDoNotDisplay                 BIsEjectable = 24
	BIsAutoMounted                BIsEjectable = 14
	BIsCasePreserving             BIsEjectable = 23
	BIsCaseSensitive              BIsEjectable = 22
	BIsEjectableValue             BIsEjectable = 0
	BIsOnExternalBus              BIsEjectable = 27
	BIsOnInternalBus              BIsEjectable = 21
	BIsRemovable                  BIsEjectable = 25
	BL2PCanMapFileBlocks          BIsEjectable = 9
	BNoRootTimes                  BIsEjectable = 26
	BNoVolumeSizes                BIsEjectable = 20
	BParentModDateChanges         BIsEjectable = 10
	BSupports2TBFiles             BIsEjectable = 4
	BSupportsExclusiveLocks       BIsEjectable = 18
	BSupportsExtendedFileSecurity BIsEjectable = 28
	BSupportsFSCatalogSearch      BIsEjectable = 2
	BSupportsFSExchangeObjects    BIsEjectable = 3
	BSupportsHFSPlusAPIs          BIsEjectable = 1
	BSupportsJournaling           BIsEjectable = 19
	BSupportsLongNames            BIsEjectable = 5
	BSupportsMultiScriptNames     BIsEjectable = 6
	BSupportsNamedForks           BIsEjectable = 7
	BSupportsSubtreeIterators     BIsEjectable = 8
	BSupportsSymbolicLinks        BIsEjectable = 13
)

func (e BIsEjectable) String() string {
	switch e {
	case BAllowCDiDataHandler:
		return "BAllowCDiDataHandler"
	case BAncestorModDateChanges:
		return "BAncestorModDateChanges"
	case BDoNotDisplay:
		return "BDoNotDisplay"
	case BIsAutoMounted:
		return "BIsAutoMounted"
	case BIsCasePreserving:
		return "BIsCasePreserving"
	case BIsCaseSensitive:
		return "BIsCaseSensitive"
	case BIsEjectableValue:
		return "BIsEjectableValue"
	case BIsOnExternalBus:
		return "BIsOnExternalBus"
	case BIsOnInternalBus:
		return "BIsOnInternalBus"
	case BIsRemovable:
		return "BIsRemovable"
	case BL2PCanMapFileBlocks:
		return "BL2PCanMapFileBlocks"
	case BNoRootTimes:
		return "BNoRootTimes"
	case BNoVolumeSizes:
		return "BNoVolumeSizes"
	case BParentModDateChanges:
		return "BParentModDateChanges"
	case BSupports2TBFiles:
		return "BSupports2TBFiles"
	case BSupportsExclusiveLocks:
		return "BSupportsExclusiveLocks"
	case BSupportsExtendedFileSecurity:
		return "BSupportsExtendedFileSecurity"
	case BSupportsFSCatalogSearch:
		return "BSupportsFSCatalogSearch"
	case BSupportsFSExchangeObjects:
		return "BSupportsFSExchangeObjects"
	case BSupportsHFSPlusAPIs:
		return "BSupportsHFSPlusAPIs"
	case BSupportsJournaling:
		return "BSupportsJournaling"
	case BSupportsLongNames:
		return "BSupportsLongNames"
	case BSupportsMultiScriptNames:
		return "BSupportsMultiScriptNames"
	case BSupportsNamedForks:
		return "BSupportsNamedForks"
	case BSupportsSubtreeIterators:
		return "BSupportsSubtreeIterators"
	case BSupportsSymbolicLinks:
		return "BSupportsSymbolicLinks"
	default:
		return fmt.Sprintf("BIsEjectable(%d)", e)
	}
}

type BLimitFCBs uint

const (
	BAccessCntl                  BLimitFCBs = 18
	BHasBTreeMgr                 BLimitFCBs = 5
	BHasBlankAccessPrivileges    BLimitFCBs = 4
	BHasCatSearch                BLimitFCBs = 7
	BHasCopyFile                 BLimitFCBs = 14
	BHasDesktopMgr               BLimitFCBs = 12
	BHasExtFSVol                 BLimitFCBs = 16
	BHasFileIDs                  BLimitFCBs = 6
	BHasFolderLock               BLimitFCBs = 10
	BHasMoveRename               BLimitFCBs = 13
	BHasOpenDeny                 BLimitFCBs = 15
	BHasPersonalAccessPrivileges BLimitFCBs = 9
	BHasShortName                BLimitFCBs = 11
	BHasUserGroupList            BLimitFCBs = 8
	BLimitFCBsValue              BLimitFCBs = 31
	BLocalWList                  BLimitFCBs = 30
	BNoBootBlks                  BLimitFCBs = 19
	BNoDeskItems                 BLimitFCBs = 20
	BNoLclSync                   BLimitFCBs = 27
	BNoMiniFndr                  BLimitFCBs = 29
	BNoSwitchTo                  BLimitFCBs = 25
	BNoSysDir                    BLimitFCBs = 17
	BNoVNEdit                    BLimitFCBs = 28
	BSupportsAsyncRequests       BLimitFCBs = 3
	BSupportsTrashVolumeCache    BLimitFCBs = 2
	BTrshOffLine                 BLimitFCBs = 26
)

func (e BLimitFCBs) String() string {
	switch e {
	case BAccessCntl:
		return "BAccessCntl"
	case BHasBTreeMgr:
		return "BHasBTreeMgr"
	case BHasBlankAccessPrivileges:
		return "BHasBlankAccessPrivileges"
	case BHasCatSearch:
		return "BHasCatSearch"
	case BHasCopyFile:
		return "BHasCopyFile"
	case BHasDesktopMgr:
		return "BHasDesktopMgr"
	case BHasExtFSVol:
		return "BHasExtFSVol"
	case BHasFileIDs:
		return "BHasFileIDs"
	case BHasFolderLock:
		return "BHasFolderLock"
	case BHasMoveRename:
		return "BHasMoveRename"
	case BHasOpenDeny:
		return "BHasOpenDeny"
	case BHasPersonalAccessPrivileges:
		return "BHasPersonalAccessPrivileges"
	case BHasShortName:
		return "BHasShortName"
	case BHasUserGroupList:
		return "BHasUserGroupList"
	case BLimitFCBsValue:
		return "BLimitFCBsValue"
	case BLocalWList:
		return "BLocalWList"
	case BNoBootBlks:
		return "BNoBootBlks"
	case BNoDeskItems:
		return "BNoDeskItems"
	case BNoLclSync:
		return "BNoLclSync"
	case BNoMiniFndr:
		return "BNoMiniFndr"
	case BNoSwitchTo:
		return "BNoSwitchTo"
	case BNoSysDir:
		return "BNoSysDir"
	case BNoVNEdit:
		return "BNoVNEdit"
	case BSupportsAsyncRequests:
		return "BSupportsAsyncRequests"
	case BSupportsTrashVolumeCache:
		return "BSupportsTrashVolumeCache"
	case BTrshOffLine:
		return "BTrshOffLine"
	default:
		return fmt.Sprintf("BLimitFCBs(%d)", e)
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
	BadDragFlavorErr          BadDragRefErr = -1852
	BadDragItemErr            BadDragRefErr = -1851
	BadDragRefErrValue        BadDragRefErr = -1850
	BadImageErr               BadDragRefErr = -1861
	BadImageRgnErr            BadDragRefErr = -1860
	CantGetFlavorErr          BadDragRefErr = -1854
	DragNotAcceptedErr        BadDragRefErr = -1857
	DuplicateFlavorErr        BadDragRefErr = -1853
	DuplicateHandlerErr       BadDragRefErr = -1855
	HandlerNotFoundErr        BadDragRefErr = -1856
	NoSuitableDisplaysErr     BadDragRefErr = -1859
	NonDragOriginatorErr      BadDragRefErr = -1862
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
	BadFolderDescErrValue  BadFolderDescErr = -4270
	BadRoutingSizeErr      BadFolderDescErr = -4276
	DuplicateFolderDescErr BadFolderDescErr = -4271
	DuplicateRoutingErr    BadFolderDescErr = -4274
	InvalidFolderTypeErr   BadFolderDescErr = -4273
	NoMoreFolderDescErr    BadFolderDescErr = -4272
	RoutingNotFoundErr     BadFolderDescErr = -4275
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
	AtpBadRsp         Buf2SmallErr = -3107
	AtpLenErr         Buf2SmallErr = -3106
	Buf2SmallErrValue Buf2SmallErr = -3101
	CkSumErr          Buf2SmallErr = -3103
	ExtractErr        Buf2SmallErr = -3104
	NoMPPErr          Buf2SmallErr = -3102
	ReadQErr          Buf2SmallErr = -3105
	RecNotFnd         Buf2SmallErr = -3108
	SktClosedErr      Buf2SmallErr = -3109
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

type CAEList uint

const (
	CAEListValue    CAEList = 'l'<<24 | 'i'<<16 | 's'<<8 | 't' // 'list'
	CApplication    CAEList = 'c'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'capp'
	CArc            CAEList = 'c'<<24 | 'a'<<16 | 'r'<<8 | 'c' // 'carc'
	CBoolean        CAEList = 'b'<<24 | 'o'<<16 | 'o'<<8 | 'l' // 'bool'
	CCell           CAEList = 'c'<<24 | 'c'<<16 | 'e'<<8 | 'l' // 'ccel'
	CChar           CAEList = 'c'<<24 | 'h'<<16 | 'a'<<8 | ' ' // 'cha '
	CColorTable     CAEList = 'c'<<24 | 'l'<<16 | 'r'<<8 | 't' // 'clrt'
	CColumn         CAEList = 'c'<<24 | 'c'<<16 | 'o'<<8 | 'l' // 'ccol'
	CDocument       CAEList = 'd'<<24 | 'o'<<16 | 'c'<<8 | 'u' // 'docu'
	CDrawingArea    CAEList = 'c'<<24 | 'd'<<16 | 'r'<<8 | 'w' // 'cdrw'
	CEnumeration    CAEList = 'e'<<24 | 'n'<<16 | 'u'<<8 | 'm' // 'enum'
	CFile           CAEList = 'f'<<24 | 'i'<<16 | 'l'<<8 | 'e' // 'file'
	CFixed          CAEList = 'f'<<24 | 'i'<<16 | 'x'<<8 | 'd' // 'fixd'
	CFixedPoint     CAEList = 'f'<<24 | 'p'<<16 | 'n'<<8 | 't' // 'fpnt'
	CFixedRectangle CAEList = 'f'<<24 | 'r'<<16 | 'c'<<8 | 't' // 'frct'
	CGraphicLine    CAEList = 'g'<<24 | 'l'<<16 | 'i'<<8 | 'n' // 'glin'
	CGraphicObject  CAEList = 'c'<<24 | 'g'<<16 | 'o'<<8 | 'b' // 'cgob'
	CGraphicShape   CAEList = 'c'<<24 | 'g'<<16 | 's'<<8 | 'h' // 'cgsh'
	CGraphicText    CAEList = 'c'<<24 | 'g'<<16 | 't'<<8 | 'x' // 'cgtx'
	CGroupedGraphic CAEList = 'c'<<24 | 'p'<<16 | 'i'<<8 | 'c' // 'cpic'
)

func (e CAEList) String() string {
	switch e {
	case CAEListValue:
		return "CAEListValue"
	case CApplication:
		return "CApplication"
	case CArc:
		return "CArc"
	case CBoolean:
		return "CBoolean"
	case CCell:
		return "CCell"
	case CChar:
		return "CChar"
	case CColorTable:
		return "CColorTable"
	case CColumn:
		return "CColumn"
	case CDocument:
		return "CDocument"
	case CDrawingArea:
		return "CDrawingArea"
	case CEnumeration:
		return "CEnumeration"
	case CFile:
		return "CFile"
	case CFixed:
		return "CFixed"
	case CFixedPoint:
		return "CFixedPoint"
	case CFixedRectangle:
		return "CFixedRectangle"
	case CGraphicLine:
		return "CGraphicLine"
	case CGraphicObject:
		return "CGraphicObject"
	case CGraphicShape:
		return "CGraphicShape"
	case CGraphicText:
		return "CGraphicText"
	case CGroupedGraphic:
		return "CGroupedGraphic"
	default:
		return fmt.Sprintf("CAEList(%d)", e)
	}
}

type CDEFNFndConstants int

const (
	CDEFNFnd         CDEFNFndConstants = 88
	DsBadStartupDisk CDEFNFndConstants = 42
	DsHD20Installed  CDEFNFndConstants = -12
	DsNotThe1        CDEFNFndConstants = 31
	DsSystemFileErr  CDEFNFndConstants = 43
	ExUserBreak      CDEFNFndConstants = -492
	FsDSIntErr       CDEFNFndConstants = -127
	HMenuFindErr     CDEFNFndConstants = -127
	MBarNFnd         CDEFNFndConstants = -126
	StrUserBreak     CDEFNFndConstants = -491
	UserBreak        CDEFNFndConstants = -490
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

type CInsertionLoc uint

const (
	CInsertionLocValue  CInsertionLoc = 'i'<<24 | 'n'<<16 | 's'<<8 | 'l' // 'insl'
	CInsertionPoint     CInsertionLoc = 'c'<<24 | 'i'<<16 | 'n'<<8 | 's' // 'cins'
	CIntlText           CInsertionLoc = 'i'<<24 | 't'<<16 | 'x'<<8 | 't' // 'itxt'
	CIntlWritingCode    CInsertionLoc = 'i'<<24 | 'n'<<16 | 't'<<8 | 'l' // 'intl'
	CItem               CInsertionLoc = 'c'<<24 | 'i'<<16 | 't'<<8 | 'm' // 'citm'
	CLine               CInsertionLoc = 'c'<<24 | 'l'<<16 | 'i'<<8 | 'n' // 'clin'
	CLongDateTime       CInsertionLoc = 'l'<<24 | 'd'<<16 | 't'<<8 | ' ' // 'ldt '
	CLongFixed          CInsertionLoc = 'l'<<24 | 'f'<<16 | 'x'<<8 | 'd' // 'lfxd'
	CLongFixedPoint     CInsertionLoc = 'l'<<24 | 'f'<<16 | 'p'<<8 | 't' // 'lfpt'
	CLongFixedRectangle CInsertionLoc = 'l'<<24 | 'f'<<16 | 'r'<<8 | 'c' // 'lfrc'
	CLongInteger        CInsertionLoc = 'l'<<24 | 'o'<<16 | 'n'<<8 | 'g' // 'long'
	CLongPoint          CInsertionLoc = 'l'<<24 | 'p'<<16 | 'n'<<8 | 't' // 'lpnt'
	CLongRectangle      CInsertionLoc = 'l'<<24 | 'r'<<16 | 'c'<<8 | 't' // 'lrct'
	CMachineLoc         CInsertionLoc = 'm'<<24 | 'L'<<16 | 'o'<<8 | 'c' // 'mLoc'
	CMenu               CInsertionLoc = 'c'<<24 | 'm'<<16 | 'n'<<8 | 'u' // 'cmnu'
	CMenuItem           CInsertionLoc = 'c'<<24 | 'm'<<16 | 'e'<<8 | 'n' // 'cmen'
	CObject             CInsertionLoc = 'c'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'cobj'
	CObjectSpecifier    CInsertionLoc = 'o'<<24 | 'b'<<16 | 'j'<<8 | ' ' // 'obj '
	COpenableObject     CInsertionLoc = 'c'<<24 | 'o'<<16 | 'o'<<8 | 'b' // 'coob'
	COval               CInsertionLoc = 'c'<<24 | 'o'<<16 | 'v'<<8 | 'l' // 'covl'
)

func (e CInsertionLoc) String() string {
	switch e {
	case CInsertionLocValue:
		return "CInsertionLocValue"
	case CInsertionPoint:
		return "CInsertionPoint"
	case CIntlText:
		return "CIntlText"
	case CIntlWritingCode:
		return "CIntlWritingCode"
	case CItem:
		return "CItem"
	case CLine:
		return "CLine"
	case CLongDateTime:
		return "CLongDateTime"
	case CLongFixed:
		return "CLongFixed"
	case CLongFixedPoint:
		return "CLongFixedPoint"
	case CLongFixedRectangle:
		return "CLongFixedRectangle"
	case CLongInteger:
		return "CLongInteger"
	case CLongPoint:
		return "CLongPoint"
	case CLongRectangle:
		return "CLongRectangle"
	case CMachineLoc:
		return "CMachineLoc"
	case CMenu:
		return "CMenu"
	case CMenuItem:
		return "CMenuItem"
	case CObject:
		return "CObject"
	case CObjectSpecifier:
		return "CObjectSpecifier"
	case COpenableObject:
		return "COpenableObject"
	case COval:
		return "COval"
	default:
		return fmt.Sprintf("CInsertionLoc(%d)", e)
	}
}

type CKeystroke uint

const (
	CKeystrokeValue CKeystroke = 'k'<<24 | 'p'<<16 | 'r'<<8 | 's' // 'kprs'
	ECapsLockDown   CKeystroke = 'K'<<24 | 'c'<<16 | 'l'<<8 | 'k' // 'Kclk'
	EClearKey       CKeystroke = 1802716928
	ECommandDown    CKeystroke = 'K'<<24 | 'c'<<16 | 'm'<<8 | 'd' // 'Kcmd'
	EControlDown    CKeystroke = 'K'<<24 | 'c'<<16 | 't'<<8 | 'l' // 'Kctl'
	EDeleteKey      CKeystroke = 1802711808
	EDownArrowKey   CKeystroke = 1802730752
	EEndKey         CKeystroke = 1802729216
	EEnterKey       CKeystroke = 1802718208
	EEscapeKey      CKeystroke = 1802712320
	EF10Key         CKeystroke = 1802726656
	EF11Key         CKeystroke = 1802725120
	EF12Key         CKeystroke = 1802727168
	EF13Key         CKeystroke = 1802725632
	EF14Key         CKeystroke = 1802726144
	EF15Key         CKeystroke = 1802727680
	EF1Key          CKeystroke = 1802729984
	EF2Key          CKeystroke = 1802729472
	EF3Key          CKeystroke = 1802724096
	EF4Key          CKeystroke = 1802728960
	EF5Key          CKeystroke = 1802723328
	EF6Key          CKeystroke = 1802723584
	EF7Key          CKeystroke = 1802723840
	EF8Key          CKeystroke = 1802724352
	EF9Key          CKeystroke = 1802724608
	EForwardDelKey  CKeystroke = 1802728704
	EHelpKey        CKeystroke = 1802727936
	EHomeKey        CKeystroke = 1802728192
	EKeyKind        CKeystroke = 'e'<<24 | 'k'<<16 | 's'<<8 | 't' // 'ekst'
	ELeftArrowKey   CKeystroke = 1802730240
	EModifiers      CKeystroke = 'e'<<24 | 'M'<<16 | 'd'<<8 | 's' // 'eMds'
	EOptionDown     CKeystroke = 'K'<<24 | 'o'<<16 | 'p'<<8 | 't' // 'Kopt'
	EPageDownKey    CKeystroke = 1802729728
	EPageUpKey      CKeystroke = 1802728448
	EReturnKey      CKeystroke = 1802707968
	ERightArrowKey  CKeystroke = 1802730496
	EShiftDown      CKeystroke = 'K'<<24 | 's'<<16 | 'f'<<8 | 't' // 'Ksft'
	ETabKey         CKeystroke = 1802711040
	EUpArrowKey     CKeystroke = 1802731008
	PKeyKind        CKeystroke = 'k'<<24 | 'k'<<16 | 'n'<<8 | 'd' // 'kknd'
	PKeystrokeKey   CKeystroke = 'k'<<24 | 'M'<<16 | 's'<<8 | 'g' // 'kMsg'
	PModifiers      CKeystroke = 'k'<<24 | 'M'<<16 | 'o'<<8 | 'd' // 'kMod'
)

func (e CKeystroke) String() string {
	switch e {
	case CKeystrokeValue:
		return "CKeystrokeValue"
	case ECapsLockDown:
		return "ECapsLockDown"
	case EClearKey:
		return "EClearKey"
	case ECommandDown:
		return "ECommandDown"
	case EControlDown:
		return "EControlDown"
	case EDeleteKey:
		return "EDeleteKey"
	case EDownArrowKey:
		return "EDownArrowKey"
	case EEndKey:
		return "EEndKey"
	case EEnterKey:
		return "EEnterKey"
	case EEscapeKey:
		return "EEscapeKey"
	case EF10Key:
		return "EF10Key"
	case EF11Key:
		return "EF11Key"
	case EF12Key:
		return "EF12Key"
	case EF13Key:
		return "EF13Key"
	case EF14Key:
		return "EF14Key"
	case EF15Key:
		return "EF15Key"
	case EF1Key:
		return "EF1Key"
	case EF2Key:
		return "EF2Key"
	case EF3Key:
		return "EF3Key"
	case EF4Key:
		return "EF4Key"
	case EF5Key:
		return "EF5Key"
	case EF6Key:
		return "EF6Key"
	case EF7Key:
		return "EF7Key"
	case EF8Key:
		return "EF8Key"
	case EF9Key:
		return "EF9Key"
	case EForwardDelKey:
		return "EForwardDelKey"
	case EHelpKey:
		return "EHelpKey"
	case EHomeKey:
		return "EHomeKey"
	case EKeyKind:
		return "EKeyKind"
	case ELeftArrowKey:
		return "ELeftArrowKey"
	case EModifiers:
		return "EModifiers"
	case EOptionDown:
		return "EOptionDown"
	case EPageDownKey:
		return "EPageDownKey"
	case EPageUpKey:
		return "EPageUpKey"
	case EReturnKey:
		return "EReturnKey"
	case ERightArrowKey:
		return "ERightArrowKey"
	case EShiftDown:
		return "EShiftDown"
	case ETabKey:
		return "ETabKey"
	case EUpArrowKey:
		return "EUpArrowKey"
	case PKeyKind:
		return "PKeyKind"
	case PKeystrokeKey:
		return "PKeystrokeKey"
	case PModifiers:
		return "PModifiers"
	default:
		return fmt.Sprintf("CKeystroke(%d)", e)
	}
}

type CMatchErr int

const (
	CDepthErr             CMatchErr = -157
	CDevErr               CMatchErr = -155
	CMatchErrValue        CMatchErr = -150
	CNoMemErr             CMatchErr = -152
	CProtectErr           CMatchErr = -154
	CRangeErr             CMatchErr = -153
	CResErr               CMatchErr = -156
	CTempMemErr           CMatchErr = -151
	CantLoadPickMethodErr CMatchErr = -11003
	ColorsRequestedErr    CMatchErr = -11004
	PictInfoIDErr         CMatchErr = -11001
	PictInfoVerbErr       CMatchErr = -11002
	PictInfoVersionErr    CMatchErr = -11000
	PictureDataErr        CMatchErr = -11005
	RgnTooBigErr          CMatchErr = -500
	UpdPixMemErr          CMatchErr = -125
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
	CPixel          CParagraph = 'c'<<24 | 'p'<<16 | 'x'<<8 | 'l' // 'cpxl'
	CPixelMap       CParagraph = 'c'<<24 | 'p'<<16 | 'i'<<8 | 'x' // 'cpix'
	CPolygon        CParagraph = 'c'<<24 | 'p'<<16 | 'g'<<8 | 'n' // 'cpgn'
	// CProperty: A property of any object class.
	CProperty    CParagraph = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
	CQDPoint     CParagraph = 'Q'<<24 | 'D'<<16 | 'p'<<8 | 't' // 'QDpt'
	CQDRectangle CParagraph = 'q'<<24 | 'd'<<16 | 'r'<<8 | 't' // 'qdrt'
	// CRGBColor: An RGB color value.
	CRGBColor         CParagraph = 'c'<<24 | 'R'<<16 | 'G'<<8 | 'B' // 'cRGB'
	CRectangle        CParagraph = 'c'<<24 | 'r'<<16 | 'e'<<8 | 'c' // 'crec'
	CRotation         CParagraph = 't'<<24 | 'r'<<16 | 'o'<<8 | 't' // 'trot'
	CRoundedRectangle CParagraph = 'c'<<24 | 'r'<<16 | 'r'<<8 | 'c' // 'crrc'
	CRow              CParagraph = 'c'<<24 | 'r'<<16 | 'o'<<8 | 'w' // 'crow'
	CSelection        CParagraph = 'c'<<24 | 's'<<16 | 'e'<<8 | 'l' // 'csel'
	CShortInteger     CParagraph = 's'<<24 | 'h'<<16 | 'o'<<8 | 'r' // 'shor'
	CTable            CParagraph = 'c'<<24 | 't'<<16 | 'b'<<8 | 'l' // 'ctbl'
	CText             CParagraph = 'c'<<24 | 't'<<16 | 'x'<<8 | 't' // 'ctxt'
	CTextFlow         CParagraph = 'c'<<24 | 'f'<<16 | 'l'<<8 | 'o' // 'cflo'
	CTextStyles       CParagraph = 't'<<24 | 's'<<16 | 't'<<8 | 'y' // 'tsty'
	CType             CParagraph = 't'<<24 | 'y'<<16 | 'p'<<8 | 'e' // 'type'
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

type CURL uint

const (
	// CFTPItem: Specifies FTP (File Transfer Protocol) protocol.
	CFTPItem CURL = 'f'<<24 | 't'<<16 | 'p'<<8 | ' ' // 'ftp '
	// CHTML: Specifies HTML (HyperText Markup Language) format.
	CHTML CURL = 'h'<<24 | 't'<<16 | 'm'<<8 | 'l' // 'html'
	// CInternetAddress: Specifies an Internet or Intranet address for the TCP/IP protocol.
	CInternetAddress CURL = 'I'<<24 | 'P'<<16 | 'A'<<8 | 'D' // 'IPAD'
	// CURLValue: Specifies a Uniform Resource Locator or Uniform Resource ID (URI).
	CURLValue CURL = 'u'<<24 | 'r'<<16 | 'l'<<8 | ' ' // 'url '
)

func (e CURL) String() string {
	switch e {
	case CFTPItem:
		return "CFTPItem"
	case CHTML:
		return "CHTML"
	case CInternetAddress:
		return "CInternetAddress"
	case CURLValue:
		return "CURLValue"
	default:
		return fmt.Sprintf("CURL(%d)", e)
	}
}

type CVersion uint

const (
	CVersionValue     CVersion = 'v'<<24 | 'e'<<16 | 'r'<<8 | 's' // 'vers'
	CWindow           CVersion = 'c'<<24 | 'w'<<16 | 'i'<<8 | 'n' // 'cwin'
	CWord             CVersion = 'c'<<24 | 'w'<<16 | 'o'<<8 | 'r' // 'cwor'
	EnumArrows        CVersion = 'a'<<24 | 'r'<<16 | 'r'<<8 | 'o' // 'arro'
	EnumJustification CVersion = 'j'<<24 | 'u'<<16 | 's'<<8 | 't' // 'just'
	EnumKeyForm       CVersion = 'k'<<24 | 'f'<<16 | 'r'<<8 | 'm' // 'kfrm'
	EnumPosition      CVersion = 'p'<<24 | 'o'<<16 | 's'<<8 | 'i' // 'posi'
	EnumProtection    CVersion = 'p'<<24 | 'r'<<16 | 't'<<8 | 'n' // 'prtn'
	EnumQuality       CVersion = 'q'<<24 | 'u'<<16 | 'a'<<8 | 'l' // 'qual'
	EnumSaveOptions   CVersion = 's'<<24 | 'a'<<16 | 'v'<<8 | 'o' // 'savo'
	EnumStyle         CVersion = 's'<<24 | 't'<<16 | 'y'<<8 | 'l' // 'styl'
	EnumTransferMode  CVersion = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
	KAEAbout          CVersion = 'a'<<24 | 'b'<<16 | 'o'<<8 | 'u' // 'abou'
	KAEAfter          CVersion = 'a'<<24 | 'f'<<16 | 't'<<8 | 'e' // 'afte'
	KAEAliasSelection CVersion = 's'<<24 | 'a'<<16 | 'l'<<8 | 'i' // 'sali'
	KAEAllCaps        CVersion = 'a'<<24 | 'l'<<16 | 'c'<<8 | 'p' // 'alcp'
	KAEArrowAtEnd     CVersion = 'a'<<24 | 'r'<<16 | 'e'<<8 | 'n' // 'aren'
	KAEArrowAtStart   CVersion = 'a'<<24 | 'r'<<16 | 's'<<8 | 't' // 'arst'
	KAEArrowBothEnds  CVersion = 'a'<<24 | 'r'<<16 | 'b'<<8 | 'o' // 'arbo'
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

type Cal uint

const (
	// Deprecated.
	CalArabicCivil Cal = 1
	// Deprecated.
	CalArabicLunar Cal = 2
	// Deprecated.
	CalCoptic Cal = 5
	// Deprecated.
	CalGregorian Cal = 0
	// Deprecated.
	CalJapanese Cal = 3
	// Deprecated.
	CalJewish Cal = 4
	// Deprecated.
	CalPersian Cal = 6
)

func (e Cal) String() string {
	switch e {
	case CalArabicCivil:
		return "CalArabicCivil"
	case CalArabicLunar:
		return "CalArabicLunar"
	case CalCoptic:
		return "CalCoptic"
	case CalGregorian:
		return "CalGregorian"
	case CalJapanese:
		return "CalJapanese"
	case CalJewish:
		return "CalJewish"
	case CalPersian:
		return "CalPersian"
	default:
		return fmt.Sprintf("Cal(%d)", e)
	}
}

type CannotFindAtomErr int

const (
	AAPNotCreatedErr                 CannotFindAtomErr = -2120
	AAPNotFoundErr                   CannotFindAtomErr = -2121
	ASDBadForkErr                    CannotFindAtomErr = -2123
	ASDBadHeaderErr                  CannotFindAtomErr = -2122
	ASDEntryNotFoundErr              CannotFindAtomErr = -2124
	AtomIndexInvalidErr              CannotFindAtomErr = -2104
	AtomsNotOfSameTypeErr            CannotFindAtomErr = -2103
	CannotBeLeafAtomErr              CannotFindAtomErr = -2109
	CannotFindAtomErrValue           CannotFindAtomErr = -2101
	DuplicateAtomTypeAndIDErr        CannotFindAtomErr = -2105
	EmptyPathErr                     CannotFindAtomErr = -2111
	FileOffsetTooBigErr              CannotFindAtomErr = -2125
	InvalidAtomContainerErr          CannotFindAtomErr = -2107
	InvalidAtomErr                   CannotFindAtomErr = -2106
	InvalidAtomTypeErr               CannotFindAtomErr = -2108
	NoPathMappingErr                 CannotFindAtomErr = -2112
	NotAllowedToSaveMovieErr         CannotFindAtomErr = -2126
	NotEnoughDataErr                 CannotFindAtomErr = -2149
	NotLeafAtomErr                   CannotFindAtomErr = -2102
	PathNotVerifiedErr               CannotFindAtomErr = -2113
	PathTooLongErr                   CannotFindAtomErr = -2110
	QfcbNotCreatedErr                CannotFindAtomErr = -2119
	QfcbNotFoundErr                  CannotFindAtomErr = -2118
	QtActionNotHandledErr            CannotFindAtomErr = -2157
	QtNetworkAlreadyAllocatedErr     CannotFindAtomErr = -2127
	QtXMLApplicationErr              CannotFindAtomErr = -2159
	QtXMLParseErr                    CannotFindAtomErr = -2158
	UnknownFormatErr                 CannotFindAtomErr = -2114
	UrlDataHFTPBadNameListErr        CannotFindAtomErr = -2144
	UrlDataHFTPBadPasswordErr        CannotFindAtomErr = -2136
	UrlDataHFTPBadUserErr            CannotFindAtomErr = -2135
	UrlDataHFTPDataConnectionErr     CannotFindAtomErr = -2138
	UrlDataHFTPFilenameErr           CannotFindAtomErr = -2142
	UrlDataHFTPNeedPasswordErr       CannotFindAtomErr = -2145
	UrlDataHFTPNoDirectoryErr        CannotFindAtomErr = -2139
	UrlDataHFTPNoNetDriverErr        CannotFindAtomErr = -2143
	UrlDataHFTPNoPasswordErr         CannotFindAtomErr = -2146
	UrlDataHFTPPermissionsErr        CannotFindAtomErr = -2141
	UrlDataHFTPProtocolErr           CannotFindAtomErr = -2133
	UrlDataHFTPQuotaErr              CannotFindAtomErr = -2140
	UrlDataHFTPServerDisconnectedErr CannotFindAtomErr = -2147
	UrlDataHFTPServerErr             CannotFindAtomErr = -2137
	UrlDataHFTPShutdownErr           CannotFindAtomErr = -2134
	UrlDataHFTPURLErr                CannotFindAtomErr = -2148
	UrlDataHHTTPNoNetDriverErr       CannotFindAtomErr = -2130
	UrlDataHHTTPProtocolErr          CannotFindAtomErr = -2129
	UrlDataHHTTPRedirectErr          CannotFindAtomErr = -2132
	UrlDataHHTTPURLErr               CannotFindAtomErr = -2131
	WackBadFileErr                   CannotFindAtomErr = -2115
	WackBadMetaDataErr               CannotFindAtomErr = -2117
	WackForkNotFoundErr              CannotFindAtomErr = -2116
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
	CfragAbortClosureErr    Cfrag = -2830
	CfragArchitectureErr    Cfrag = -2823
	CfragCFMInternalErr     Cfrag = -2819
	CfragCFMStartupErr      Cfrag = -2818
	CfragCFragRsrcErr       Cfrag = -2857
	CfragClosureIDErr       Cfrag = -2829
	CfragConnectionIDErr    Cfrag = -2801
	CfragContainerIDErr     Cfrag = -2828
	CfragContextIDErr       Cfrag = -2800
	CfragDupRegistrationErr Cfrag = -2805
	CfragExecFileRefErr     Cfrag = -2854
	CfragFileSizeErr        Cfrag = -2825
	CfragFirstErrCode       Cfrag = -2800
	CfragFirstReservedCode  Cfrag = -2897
	CfragFragmentCorruptErr Cfrag = -2820
	CfragFragmentFormatErr  Cfrag = -2806
	CfragFragmentUsageErr   Cfrag = -2824
	CfragImportTooNewErr    Cfrag = -2814
	CfragImportTooOldErr    Cfrag = -2813
	CfragInitAtBootErr      Cfrag = -2816
	CfragInitFunctionErr    Cfrag = -2821
	CfragInitLoopErr        Cfrag = -2815
	CfragInitOrderErr       Cfrag = -2812
	CfragLastErrCode        Cfrag = -2899
	CfragLibConnErr         Cfrag = -2817
	CfragMapFileErr         Cfrag = -2851
	CfragNoApplicationErr   Cfrag = -2822
	CfragNoClientMemErr     Cfrag = -2810
	CfragNoIDsErr           Cfrag = -2811
	CfragNoLibraryErr       Cfrag = -2804
	CfragNoPositionErr      Cfrag = -2808
	CfragNoPrivateMemErr    Cfrag = -2809
	CfragNoRegistrationErr  Cfrag = -2827
	CfragNoSectionErr       Cfrag = -2803
	CfragNoSymbolErr        Cfrag = -2802
	CfragNotClosureErr      Cfrag = -2826
	CfragOutputLengthErr    Cfrag = -2831
	CfragReservedCode_1     Cfrag = -2899
	CfragReservedCode_2     Cfrag = -2898
	CfragReservedCode_3     Cfrag = -2897
	CfragRsrcForkErr        Cfrag = -2856
	CfragStdFolderErr       Cfrag = -2855
	CfragUnresolvedErr      Cfrag = -2807
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

type Cmp uint

const (
	// Deprecated.
	CmpIsMissing Cmp = 536870912
	// Deprecated.
	CmpThreadSafe Cmp = 268435456
	// Deprecated.
	CmpWantsRegisterMessage Cmp = 2147483648
)

func (e Cmp) String() string {
	switch e {
	case CmpIsMissing:
		return "CmpIsMissing"
	case CmpThreadSafe:
		return "CmpThreadSafe"
	case CmpWantsRegisterMessage:
		return "CmpWantsRegisterMessage"
	default:
		return fmt.Sprintf("Cmp(%d)", e)
	}
}

type CmpAlias uint

const (
	// Deprecated.
	CmpAliasNoFlags CmpAlias = 0
	// Deprecated.
	CmpAliasOnlyThisFile CmpAlias = 1
)

func (e CmpAlias) String() string {
	switch e {
	case CmpAliasNoFlags:
		return "CmpAliasNoFlags"
	case CmpAliasOnlyThisFile:
		return "CmpAliasOnlyThisFile"
	default:
		return fmt.Sprintf("CmpAlias(%d)", e)
	}
}

type CodecErr int

const (
	BadCodecCharacterizationErr        CodecErr = -8993
	CodecAbortErr                      CodecErr = -8967
	CodecBadDataErr                    CodecErr = -8969
	CodecCantQueueErr                  CodecErr = -8975
	CodecCantWhenErr                   CodecErr = -8974
	CodecConditionErr                  CodecErr = -8972
	CodecDataVersErr                   CodecErr = -8970
	CodecDisabledErr                   CodecErr = -8978
	CodecDroppedFrameErr               CodecErr = -8989
	CodecErrValue                      CodecErr = -8960
	CodecExtensionNotFoundErr          CodecErr = -8971
	CodecImageBufErr                   CodecErr = -8965
	CodecNeedAccessKeyErr              CodecErr = -8987
	CodecNeedToFlushChainErr           CodecErr = -8979
	CodecNoMemoryPleaseWaitErr         CodecErr = -8977
	CodecNothingToBlitErr              CodecErr = -8976
	CodecOffscreenFailedErr            CodecErr = -8988
	CodecOffscreenFailedPleaseRetryErr CodecErr = -8992
	CodecOpenErr                       CodecErr = -8973
	CodecParameterDialogConfirm        CodecErr = -8986
	CodecScreenBufErr                  CodecErr = -8964
	CodecSizeErr                       CodecErr = -8963
	CodecSpoolErr                      CodecErr = -8966
	CodecUnimpErr                      CodecErr = -8962
	CodecWouldOffscreenErr             CodecErr = -8968
	DirectXObjectAlreadyExists         CodecErr = -8990
	LockPortBitsBadPortErr             CodecErr = -8984
	LockPortBitsBadSurfaceErr          CodecErr = -8980
	LockPortBitsSurfaceLostErr         CodecErr = -8985
	LockPortBitsWindowClippedErr       CodecErr = -8983
	LockPortBitsWindowMovedErr         CodecErr = -8981
	LockPortBitsWindowResizedErr       CodecErr = -8982
	LockPortBitsWrongGDeviceErr        CodecErr = -8991
	NoCodecErr                         CodecErr = -8961
	NoThumbnailFoundErr                CodecErr = -8994
	ScTypeNotFoundErr                  CodecErr = -8971
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
	CollectionIndexRangeErr   Collection = -5752
	CollectionItemLockedErr   Collection = -5750
	CollectionItemNotFoundErr Collection = -5751
	CollectionVersionErr      Collection = -5753
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

type Component uint

const (
	// Deprecated.
	ComponentAutoVersionIncludeFlags Component = 4
	// Deprecated.
	ComponentDoAutoVersion Component = 1
	// Deprecated.
	ComponentHasMultiplePlatforms Component = 8
	// Deprecated.
	ComponentLoadResident Component = 16
	// Deprecated.
	ComponentWantsUnregister Component = 2
)

func (e Component) String() string {
	switch e {
	case ComponentAutoVersionIncludeFlags:
		return "ComponentAutoVersionIncludeFlags"
	case ComponentDoAutoVersion:
		return "ComponentDoAutoVersion"
	case ComponentHasMultiplePlatforms:
		return "ComponentHasMultiplePlatforms"
	case ComponentLoadResident:
		return "ComponentLoadResident"
	case ComponentWantsUnregister:
		return "ComponentWantsUnregister"
	default:
		return fmt.Sprintf("Component(%d)", e)
	}
}

type ComponentDllLoadErr int

const (
	// ComponentDllEntryNotFoundErr: Windows error returned when a component is loading.
	ComponentDllEntryNotFoundErr ComponentDllLoadErr = -2092
	// ComponentDllLoadErrValue: Windows error returned when a component is loading.
	ComponentDllLoadErrValue  ComponentDllLoadErr = -2091
	ComponentNotThreadSafeErr ComponentDllLoadErr = -2098
	// QtmlDllEntryNotFoundErr: Windows error returned when the QuickTime Media Layer is loading.
	QtmlDllEntryNotFoundErr ComponentDllLoadErr = -2094
	// QtmlDllLoadErr: Windows error returned when the QuickTime Media Layer is loading.
	QtmlDllLoadErr          ComponentDllLoadErr = -2093
	QtmlUninitialized       ComponentDllLoadErr = -2095
	UnsupportedOSErr        ComponentDllLoadErr = -2096
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

type CoreFoundationUnknown int

const (
	CoreFoundationUnknownErr CoreFoundationUnknown = -4960
)

func (e CoreFoundationUnknown) String() string {
	switch e {
	case CoreFoundationUnknownErr:
		return "CoreFoundationUnknownErr"
	default:
		return fmt.Sprintf("CoreFoundationUnknown(%d)", e)
	}
}

type CouldNotResolveDataRef int

const (
	AuxiliaryExportDataUnavailable CouldNotResolveDataRef = -2058
	BadComponentType               CouldNotResolveDataRef = -2005
	BadDataRefIndex                CouldNotResolveDataRef = -2050
	BadEditIndex                   CouldNotResolveDataRef = -2033
	BadEditList                    CouldNotResolveDataRef = -2017
	BadImageDescription            CouldNotResolveDataRef = -2001
	BadPublicMovieAtom             CouldNotResolveDataRef = -2002
	BadTrackIndex                  CouldNotResolveDataRef = -2028
	// CantCreateSingleForkFile: The file to be created already exists.
	CantCreateSingleForkFile      CouldNotResolveDataRef = -2022
	CantEnableTrack               CouldNotResolveDataRef = -2035
	CantFindHandler               CouldNotResolveDataRef = -2003
	CantOpenHandler               CouldNotResolveDataRef = -2004
	CantPutPublicMovieAtom        CouldNotResolveDataRef = -2016
	CouldNotResolveDataRefValue   CouldNotResolveDataRef = -2000
	CouldNotUseAnExistingSample   CouldNotResolveDataRef = -2052
	DataAlreadyClosed             CouldNotResolveDataRef = -2045
	DataAlreadyOpenForWrite       CouldNotResolveDataRef = -2044
	DataNoDataRef                 CouldNotResolveDataRef = -2047
	DataNotOpenForRead            CouldNotResolveDataRef = -2042
	DataNotOpenForWrite           CouldNotResolveDataRef = -2043
	EndOfDataReached              CouldNotResolveDataRef = -2046
	FeatureUnsupported            CouldNotResolveDataRef = -2053
	GWorldsNotSameDepthAndSizeErr CouldNotResolveDataRef = -2066
	InternalQuickTimeError        CouldNotResolveDataRef = -2034
	InvalidChunkCache             CouldNotResolveDataRef = -2040
	InvalidChunkNum               CouldNotResolveDataRef = -2038
	InvalidDataRef                CouldNotResolveDataRef = -2012
	InvalidDataRefContainer       CouldNotResolveDataRef = -2049
	InvalidDuration               CouldNotResolveDataRef = -2014
	InvalidEditState              CouldNotResolveDataRef = -2023
	InvalidHandler                CouldNotResolveDataRef = -2013
	InvalidImageIndexErr          CouldNotResolveDataRef = -2068
	InvalidMedia                  CouldNotResolveDataRef = -2008
	InvalidMovie                  CouldNotResolveDataRef = -2010
	InvalidRect                   CouldNotResolveDataRef = -2036
	InvalidSampleDescIndex        CouldNotResolveDataRef = -2039
	InvalidSampleDescription      CouldNotResolveDataRef = -2041
	InvalidSampleNum              CouldNotResolveDataRef = -2037
	InvalidSampleTable            CouldNotResolveDataRef = -2011
	InvalidSpriteIDErr            CouldNotResolveDataRef = -2069
	InvalidSpriteIndexErr         CouldNotResolveDataRef = -2067
	InvalidSpritePropertyErr      CouldNotResolveDataRef = -2065
	InvalidSpriteWorldPropertyErr CouldNotResolveDataRef = -2064
	InvalidTime                   CouldNotResolveDataRef = -2015
	InvalidTrack                  CouldNotResolveDataRef = -2009
	MaxSizeToGrowTooSmall         CouldNotResolveDataRef = -2027
	MediaTypesDontMatch           CouldNotResolveDataRef = -2018
	MissingRequiredParameterErr   CouldNotResolveDataRef = -2063
	MovieTextNotFoundErr          CouldNotResolveDataRef = -2062
	MovieToolboxUninitialized     CouldNotResolveDataRef = -2020
	NoDataHandler                 CouldNotResolveDataRef = -2007
	NoDefaultDataRef              CouldNotResolveDataRef = -2051
	NoMediaHandler                CouldNotResolveDataRef = -2006
	NoMovieFound                  CouldNotResolveDataRef = -2048
	// NoRecordOfApp: A replica of the movieToolboxUninitialized error.
	NoRecordOfApp                  CouldNotResolveDataRef = -2020
	NoSoundTrackInMovieErr         CouldNotResolveDataRef = -2055
	NoSourceTreeFoundErr           CouldNotResolveDataRef = -2060
	NoVideoTrackInMovieErr         CouldNotResolveDataRef = -2054
	NonMatchingEditState           CouldNotResolveDataRef = -2024
	ProgressProcAborted            CouldNotResolveDataRef = -2019
	SamplesAlreadyInMediaErr       CouldNotResolveDataRef = -2059
	SoundSupportNotAvailableErr    CouldNotResolveDataRef = -2056
	SourceNotFoundErr              CouldNotResolveDataRef = -2061
	StaleEditState                 CouldNotResolveDataRef = -2025
	TimeNotInMedia                 CouldNotResolveDataRef = -2032
	TimeNotInTrack                 CouldNotResolveDataRef = -2031
	TrackIDNotFound                CouldNotResolveDataRef = -2029
	TrackNotInMovie                CouldNotResolveDataRef = -2030
	UnsupportedAuxiliaryImportData CouldNotResolveDataRef = -2057
	UserDataItemNotFound           CouldNotResolveDataRef = -2026
	WfFileNotFound                 CouldNotResolveDataRef = -2021
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

type Curr uint

const (
	CurrLeadingZ  Curr = 128
	CurrNegSym    Curr = 32
	CurrSymLead   Curr = 16
	CurrTrailingZ Curr = 64
)

func (e Curr) String() string {
	switch e {
	case CurrLeadingZ:
		return "CurrLeadingZ"
	case CurrNegSym:
		return "CurrNegSym"
	case CurrSymLead:
		return "CurrSymLead"
	case CurrTrailingZ:
		return "CurrTrailingZ"
	default:
		return fmt.Sprintf("Curr(%d)", e)
	}
}

type Dcm int

const (
	DcmBadDataSizeErr       Dcm = -7117
	DcmBadDictionaryErr     Dcm = -7102
	DcmBadFeatureErr        Dcm = -7124
	DcmBadFieldInfoErr      Dcm = -7111
	DcmBadFieldTypeErr      Dcm = -7112
	DcmBadFindMethodErr     Dcm = -7118
	DcmBadKeyErr            Dcm = -7115
	DcmBadPropertyErr       Dcm = -7119
	DcmBlockFullErr         Dcm = -7107
	DcmBufferOverflowErr    Dcm = -7127
	DcmDictionaryBusyErr    Dcm = -7105
	DcmDictionaryNotOpenErr Dcm = -7104
	DcmDupRecordErr         Dcm = -7109
	DcmIterationCompleteErr Dcm = -7126
	DcmNecessaryFieldErr    Dcm = -7110
	DcmNoAccessMethodErr    Dcm = -7122
	DcmNoFieldErr           Dcm = -7113
	DcmNoRecordErr          Dcm = -7108
	DcmNotDictionaryErr     Dcm = -7101
	DcmParamErr             Dcm = -7100
	DcmPermissionErr        Dcm = -7103
	DcmProtectedErr         Dcm = -7121
	DcmTooManyKeyErr        Dcm = -7116
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
	DdpLenErr      DdpSktErr = -92
	DdpSktErrValue DdpSktErr = -91
	ExcessCollsns  DdpSktErr = -95
	LapProtErr     DdpSktErr = -94
	NoBridgeErr    DdpSktErr = -93
	PortInUse      DdpSktErr = -97
	PortNotCf      DdpSktErr = -98
	PortNotPwr     DdpSktErr = -96
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
	DebuggingDuplicateOptionErr    Debugging = -13882
	DebuggingDuplicateSignatureErr Debugging = -13881
	DebuggingExecutionContextErr   Debugging = -13880
	DebuggingInvalidNameErr        Debugging = -13885
	DebuggingInvalidOptionErr      Debugging = -13884
	DebuggingInvalidSignatureErr   Debugging = -13883
	DebuggingNoCallbackErr         Debugging = -13886
	DebuggingNoMatchErr            Debugging = -13887
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

type DefaultComponent uint

const (
	// Deprecated.
	DefaultComponentAnyFlags DefaultComponent = 1
	// Deprecated.
	DefaultComponentAnyFlagsAnyManufacturer DefaultComponent = 0
	// Deprecated.
	DefaultComponentAnyFlagsAnyManufacturerAnySubType DefaultComponent = 0
	// Deprecated.
	DefaultComponentAnyManufacturer DefaultComponent = 2
	// Deprecated.
	DefaultComponentAnySubType DefaultComponent = 4
	// Deprecated.
	DefaultComponentIdentical DefaultComponent = 0
)

func (e DefaultComponent) String() string {
	switch e {
	case DefaultComponentAnyFlags:
		return "DefaultComponentAnyFlags"
	case DefaultComponentAnyFlagsAnyManufacturer:
		return "DefaultComponentAnyFlagsAnyManufacturer"
	case DefaultComponentAnyManufacturer:
		return "DefaultComponentAnyManufacturer"
	case DefaultComponentAnySubType:
		return "DefaultComponentAnySubType"
	default:
		return fmt.Sprintf("DefaultComponent(%d)", e)
	}
}

type DefaultPhysicalEntry uint

const (
	// Deprecated.
	DefaultPhysicalEntryCount DefaultPhysicalEntry = 8
)

func (e DefaultPhysicalEntry) String() string {
	switch e {
	case DefaultPhysicalEntryCount:
		return "DefaultPhysicalEntryCount"
	default:
		return fmt.Sprintf("DefaultPhysicalEntry(%d)", e)
	}
}

type DelimPad int

const (
	// Deprecated.
	DelimPadValue DelimPad = -2
	// Deprecated.
	TokenCarat DelimPad = 55
	// Deprecated.
	TokenTilda DelimPad = 44
)

func (e DelimPad) String() string {
	switch e {
	case DelimPadValue:
		return "DelimPadValue"
	case TokenCarat:
		return "TokenCarat"
	case TokenTilda:
		return "TokenTilda"
	default:
		return fmt.Sprintf("DelimPad(%d)", e)
	}
}

type DiaeresisUprY uint

const (
	// Deprecated.
	AcuteUprA DiaeresisUprY = 231
	// Deprecated.
	AcuteUprI DiaeresisUprY = 234
	// Deprecated.
	AcuteUprO DiaeresisUprY = 238
	// Deprecated.
	AcuteUprU DiaeresisUprY = 242
	// Deprecated.
	AppleLogo DiaeresisUprY = 240
	// Deprecated.
	BaseDblQuote DiaeresisUprY = 227
	// Deprecated.
	BaseSingQuote DiaeresisUprY = 226
	// Deprecated.
	BreveMark DiaeresisUprY = 249
	// Deprecated.
	Cedilla DiaeresisUprY = 252
	// Deprecated.
	CenteredDot DiaeresisUprY = 225
	// Deprecated.
	Circumflex DiaeresisUprY = 246
	// Deprecated.
	CircumflexUprA DiaeresisUprY = 229
	// Deprecated.
	CircumflexUprE DiaeresisUprY = 230
	// Deprecated.
	CircumflexUprI DiaeresisUprY = 235
	// Deprecated.
	CircumflexUprO DiaeresisUprY = 239
	// Deprecated.
	CircumflexUprU DiaeresisUprY = 243
	// Deprecated.
	DblDagger DiaeresisUprY = 224
	// Deprecated.
	DiaeresisUprE DiaeresisUprY = 232
	// Deprecated.
	DiaeresisUprI DiaeresisUprY = 236
	// Deprecated.
	DiaeresisUprYValue DiaeresisUprY = 217
	// Deprecated.
	DotlessLwrI DiaeresisUprY = 245
	// Deprecated.
	DoubleAcute DiaeresisUprY = 253
	// Deprecated.
	FiLigature DiaeresisUprY = 222
	// Deprecated.
	FlLigature DiaeresisUprY = 223
	// Deprecated.
	Fraction DiaeresisUprY = 218
	// Deprecated.
	GraveUprE DiaeresisUprY = 233
	// Deprecated.
	GraveUprI DiaeresisUprY = 237
	// Deprecated.
	GraveUprO DiaeresisUprY = 241
	// Deprecated.
	GraveUprU DiaeresisUprY = 244
	// Deprecated.
	Hachek DiaeresisUprY = 255
	// Deprecated.
	IntlCurrency DiaeresisUprY = 219
	// Deprecated.
	LeftSingGuillemet DiaeresisUprY = 220
	// Deprecated.
	Macron DiaeresisUprY = 248
	// Deprecated.
	Ogonek DiaeresisUprY = 254
	// Deprecated.
	OverDot DiaeresisUprY = 250
	// Deprecated.
	PerThousand DiaeresisUprY = 228
	// Deprecated.
	RightSingGuillemet DiaeresisUprY = 221
	// Deprecated.
	RingMark DiaeresisUprY = 251
	// Deprecated.
	Tilde DiaeresisUprY = 247
)

func (e DiaeresisUprY) String() string {
	switch e {
	case AcuteUprA:
		return "AcuteUprA"
	case AcuteUprI:
		return "AcuteUprI"
	case AcuteUprO:
		return "AcuteUprO"
	case AcuteUprU:
		return "AcuteUprU"
	case AppleLogo:
		return "AppleLogo"
	case BaseDblQuote:
		return "BaseDblQuote"
	case BaseSingQuote:
		return "BaseSingQuote"
	case BreveMark:
		return "BreveMark"
	case Cedilla:
		return "Cedilla"
	case CenteredDot:
		return "CenteredDot"
	case Circumflex:
		return "Circumflex"
	case CircumflexUprA:
		return "CircumflexUprA"
	case CircumflexUprE:
		return "CircumflexUprE"
	case CircumflexUprI:
		return "CircumflexUprI"
	case CircumflexUprO:
		return "CircumflexUprO"
	case CircumflexUprU:
		return "CircumflexUprU"
	case DblDagger:
		return "DblDagger"
	case DiaeresisUprE:
		return "DiaeresisUprE"
	case DiaeresisUprI:
		return "DiaeresisUprI"
	case DiaeresisUprYValue:
		return "DiaeresisUprYValue"
	case DotlessLwrI:
		return "DotlessLwrI"
	case DoubleAcute:
		return "DoubleAcute"
	case FiLigature:
		return "FiLigature"
	case FlLigature:
		return "FlLigature"
	case Fraction:
		return "Fraction"
	case GraveUprE:
		return "GraveUprE"
	case GraveUprI:
		return "GraveUprI"
	case GraveUprO:
		return "GraveUprO"
	case GraveUprU:
		return "GraveUprU"
	case Hachek:
		return "Hachek"
	case IntlCurrency:
		return "IntlCurrency"
	case LeftSingGuillemet:
		return "LeftSingGuillemet"
	case Macron:
		return "Macron"
	case Ogonek:
		return "Ogonek"
	case OverDot:
		return "OverDot"
	case PerThousand:
		return "PerThousand"
	case RightSingGuillemet:
		return "RightSingGuillemet"
	case RingMark:
		return "RingMark"
	case Tilde:
		return "Tilde"
	default:
		return fmt.Sprintf("DiaeresisUprY(%d)", e)
	}
}

type DialogNoTimeout int

const (
	DialogNoTimeoutErr DialogNoTimeout = -5640
)

func (e DialogNoTimeout) String() string {
	switch e {
	case DialogNoTimeoutErr:
		return "DialogNoTimeoutErr"
	default:
		return fmt.Sprintf("DialogNoTimeout(%d)", e)
	}
}

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
	DsAddressErr       Ds = 2
	DsBadLibrary       Ds = 1010
	DsBusError         Ds = 1
	DsChkErr           Ds = 5
	DsCoreErr          Ds = 12
	DsFPErr            Ds = 16
	DsIOCoreErr        Ds = 14
	DsIllInstErr       Ds = 3
	DsIrqErr           Ds = 13
	DsLineAErr         Ds = 9
	DsLineFErr         Ds = 10
	DsLoadErr          Ds = 15
	DsMiscErr          Ds = 11
	DsMixedModeFailure Ds = 1011
	DsNoPackErr        Ds = 17
	DsNoPk1            Ds = 18
	DsNoPk2            Ds = 19
	DsOvflowErr        Ds = 6
	DsPrivErr          Ds = 7
	DsTraceErr         Ds = 8
	DsZeroDivErr       Ds = 4
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
	DsExtensionsDisabled    DsNoExtsMacsBug = -13
	DsGreeting              DsNoExtsMacsBug = 40
	DsMacsBugInstalled      DsNoExtsMacsBug = -10
	DsNoExtsDisassembler    DsNoExtsMacsBug = -2
	DsNoExtsMacsBugValue    DsNoExtsMacsBug = -1
	DsSysErr                DsNoExtsMacsBug = 32767
	WDEFNFnd                DsNoExtsMacsBug = 87
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
	Ds32BitMode                       DsNoFPU = 103
	DsBadPatch                        DsNoFPU = 99
	DsBufPtrTooLow                    DsNoFPU = 106
	DsCantHoldSystemHeap              DsNoFPU = 114
	DsDirtyDisk                       DsNoFPU = 20004
	DsForcedQuit                      DsNoFPU = 20002
	DsGibblyMovedToDisabledFolder     DsNoFPU = 117
	DsLostConnectionToNetworkDisk     DsNoFPU = 121
	DsMBATAPISysError                 DsNoFPU = 29203
	DsMBATASysError                   DsNoFPU = 29202
	DsMBExternFlpySysError            DsNoFPU = 29204
	DsMBFlpySysError                  DsNoFPU = 29201
	DsMBSysError                      DsNoFPU = 29200
	DsMacOSROMVersionTooOld           DsNoFPU = 120
	DsMustUseFCBAccessors             DsNoFPU = 119
	DsNeedToWriteBootBlocks           DsNoFPU = 104
	DsNoFPUValue                      DsNoFPU = 90
	DsNoPatch                         DsNoFPU = 98
	DsNotEnoughRAMToBoot              DsNoFPU = 105
	DsOldSystem                       DsNoFPU = 102
	DsPCCardATASysError               DsNoFPU = 29205
	DsParityErr                       DsNoFPU = 101
	DsRAMDiskTooBig                   DsNoFPU = 122
	DsReinsert                        DsNoFPU = 30
	DsRemoveDisk                      DsNoFPU = 20003
	DsSCSIWarn                        DsNoFPU = 20010
	DsShutDownOrRestart               DsNoFPU = 20000
	DsShutDownOrResume                DsNoFPU = 20109
	DsSwitchOffOrRestart              DsNoFPU = 20001
	DsSystemRequiresPowerPC           DsNoFPU = 116
	DsUnBootableSystem                DsNoFPU = 118
	DsVMBadBackingStore               DsNoFPU = 113
	DsVMDeferredFuncTableFull         DsNoFPU = 112
	DsWriteToSupervisorStackGuardPage DsNoFPU = 128
	ShutDownAlert                     DsNoFPU = 42
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
	DsBadLaunch      DsNoPk3 = 26
	DsBadPatchHeader DsNoPk3 = 83
	DsBadSANEOpcode  DsNoPk3 = 81
	DsBadSlotInt     DsNoPk3 = 51
	DsCDEFNotFound   DsNoPk3 = 88
	DsFSErr          DsNoPk3 = 27
	DsFinderErr      DsNoPk3 = 41
	DsHMenuFindErr   DsNoPk3 = 86
	DsMBarNFnd       DsNoPk3 = 85
	DsMDEFNotFound   DsNoPk3 = 89
	DsMemFullErr     DsNoPk3 = 25
	DsNoPk3Value     DsNoPk3 = 20
	DsNoPk4          DsNoPk3 = 21
	DsNoPk5          DsNoPk3 = 22
	DsNoPk6          DsNoPk3 = 23
	DsNoPk7          DsNoPk3 = 24
	DsStknHeap       DsNoPk3 = 28
	DsWDEFNotFound   DsNoPk3 = 87
	MenuPrgErr       DsNoPk3 = 84
	NegZcbFreeErr    DsNoPk3 = 33
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

type DummyType uint

const (
	// Deprecated.
	DrvQType DummyType = 3
	// Deprecated.
	DtQType DummyType = 7
	// Deprecated.
	DummyTypeValue DummyType = 0
	// Deprecated.
	EvType DummyType = 4
	// Deprecated.
	FsQType DummyType = 5
	// Deprecated.
	IoQType DummyType = 2
	// Deprecated.
	NmType DummyType = 8
	// Deprecated.
	SIQType DummyType = 6
	// Deprecated.
	VType DummyType = 1
)

func (e DummyType) String() string {
	switch e {
	case DrvQType:
		return "DrvQType"
	case DtQType:
		return "DtQType"
	case DummyTypeValue:
		return "DummyTypeValue"
	case EvType:
		return "EvType"
	case FsQType:
		return "FsQType"
	case IoQType:
		return "IoQType"
	case NmType:
		return "NmType"
	case SIQType:
		return "SIQType"
	case VType:
		return "VType"
	default:
		return fmt.Sprintf("DummyType(%d)", e)
	}
}

type Duration int

const (
	// Deprecated.
	DurationDay Duration = 86400000
	// Deprecated.
	DurationForever Duration = 2147483647
	// Deprecated.
	DurationHour Duration = 3600000
	// Deprecated.
	DurationMicrosecond Duration = -1
	// Deprecated.
	DurationMillisecond Duration = 1
	// Deprecated.
	DurationMinute Duration = 60000
	// Deprecated.
	DurationNoWait Duration = 0
	// Deprecated.
	DurationSecond Duration = 1000
)

func (e Duration) String() string {
	switch e {
	case DurationDay:
		return "DurationDay"
	case DurationForever:
		return "DurationForever"
	case DurationHour:
		return "DurationHour"
	case DurationMicrosecond:
		return "DurationMicrosecond"
	case DurationMillisecond:
		return "DurationMillisecond"
	case DurationMinute:
		return "DurationMinute"
	case DurationNoWait:
		return "DurationNoWait"
	case DurationSecond:
		return "DurationSecond"
	default:
		return fmt.Sprintf("Duration(%d)", e)
	}
}

type ELenErr int

const (
	ELenErrValue ELenErr = -92
	EMultiErr    ELenErr = -91
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

type EScheme uint

const (
	ESchemeValue EScheme = 'e'<<24 | 's'<<16 | 'c'<<8 | 'h' // 'esch'
	EurlAFP      EScheme = 'a'<<24 | 'f'<<16 | 'p'<<8 | ' ' // 'afp '
	EurlAT       EScheme = 'a'<<24 | 't'<<16 | ' '<<8 | ' ' // 'at  '
	EurlEPPC     EScheme = 'e'<<24 | 'p'<<16 | 'p'<<8 | 'c' // 'eppc'
	EurlFTP      EScheme = 'f'<<24 | 't'<<16 | 'p'<<8 | ' ' // 'ftp '
	EurlFile     EScheme = 'f'<<24 | 'i'<<16 | 'l'<<8 | 'e' // 'file'
	EurlGopher   EScheme = 'g'<<24 | 'p'<<16 | 'h'<<8 | 'r' // 'gphr'
	EurlHTTP     EScheme = 'h'<<24 | 't'<<16 | 't'<<8 | 'p' // 'http'
	EurlHTTPS    EScheme = 'h'<<24 | 't'<<16 | 'p'<<8 | 's' // 'htps'
	EurlIMAP     EScheme = 'i'<<24 | 'm'<<16 | 'a'<<8 | 'p' // 'imap'
	EurlLDAP     EScheme = 'u'<<24 | 'l'<<16 | 'd'<<8 | 'p' // 'uldp'
	EurlLaunch   EScheme = 'l'<<24 | 'a'<<16 | 'u'<<8 | 'n' // 'laun'
	EurlMail     EScheme = 'm'<<24 | 'a'<<16 | 'i'<<8 | 'l' // 'mail'
	EurlMailbox  EScheme = 'm'<<24 | 'b'<<16 | 'o'<<8 | 'x' // 'mbox'
	EurlMessage  EScheme = 'm'<<24 | 'e'<<16 | 's'<<8 | 's' // 'mess'
	EurlMulti    EScheme = 'm'<<24 | 'u'<<16 | 'l'<<8 | 't' // 'mult'
	EurlNFS      EScheme = 'u'<<24 | 'n'<<16 | 'f'<<8 | 's' // 'unfs'
	EurlNNTP     EScheme = 'n'<<24 | 'n'<<16 | 't'<<8 | 'p' // 'nntp'
	EurlNews     EScheme = 'n'<<24 | 'e'<<16 | 'w'<<8 | 's' // 'news'
	EurlPOP      EScheme = 'u'<<24 | 'p'<<16 | 'o'<<8 | 'p' // 'upop'
	EurlRTSP     EScheme = 'r'<<24 | 't'<<16 | 's'<<8 | 'p' // 'rtsp'
	EurlSNews    EScheme = 's'<<24 | 'n'<<16 | 'w'<<8 | 's' // 'snws'
	EurlTelnet   EScheme = 't'<<24 | 'l'<<16 | 'n'<<8 | 't' // 'tlnt'
	EurlUnknown  EScheme = 'u'<<24 | 'r'<<16 | 'l'<<8 | '?' // 'url?'
)

func (e EScheme) String() string {
	switch e {
	case ESchemeValue:
		return "ESchemeValue"
	case EurlAFP:
		return "EurlAFP"
	case EurlAT:
		return "EurlAT"
	case EurlEPPC:
		return "EurlEPPC"
	case EurlFTP:
		return "EurlFTP"
	case EurlFile:
		return "EurlFile"
	case EurlGopher:
		return "EurlGopher"
	case EurlHTTP:
		return "EurlHTTP"
	case EurlHTTPS:
		return "EurlHTTPS"
	case EurlIMAP:
		return "EurlIMAP"
	case EurlLDAP:
		return "EurlLDAP"
	case EurlLaunch:
		return "EurlLaunch"
	case EurlMail:
		return "EurlMail"
	case EurlMailbox:
		return "EurlMailbox"
	case EurlMessage:
		return "EurlMessage"
	case EurlMulti:
		return "EurlMulti"
	case EurlNFS:
		return "EurlNFS"
	case EurlNNTP:
		return "EurlNNTP"
	case EurlNews:
		return "EurlNews"
	case EurlPOP:
		return "EurlPOP"
	case EurlRTSP:
		return "EurlRTSP"
	case EurlSNews:
		return "EurlSNews"
	case EurlTelnet:
		return "EurlTelnet"
	case EurlUnknown:
		return "EurlUnknown"
	default:
		return fmt.Sprintf("EScheme(%d)", e)
	}
}

type EWS uint

const (
	// EWSArrayType: Maps to [CFArrayRef].
	EWSArrayType EWS = 0
	// EWSBooleanType: Maps to [CFBooleanRef].
	EWSBooleanType EWS = 0
	// EWSDataType: Maps to [CFDataRef].
	EWSDataType EWS = 0
	// EWSDateType: Maps to [CFDateRef].
	EWSDateType EWS = 0
	// EWSDictionaryType: Maps to [CFDictionaryRef].
	EWSDictionaryType EWS = 0
	// EWSDoubleType: Maps to [CFNumberRef] for long, double, or real numbers.
	EWSDoubleType EWS = 0
	// EWSIntegerType: Maps to [CFNumberRef] for 8, 16, 32 bit integers.
	EWSIntegerType EWS = 0
	// EWSNullType: Maps to [CFNullRef].
	EWSNullType EWS = 0
	// EWSStringType: Maps to [CFStringRef].
	EWSStringType EWS = 0
	// EWSUnknownType: No mapping is known for this type.
	EWSUnknownType EWS = 0
)

func (e EWS) String() string {
	switch e {
	case EWSArrayType:
		return "EWSArrayType"
	default:
		return fmt.Sprintf("EWS(%d)", e)
	}
}

type EditionMgrInitErr int

const (
	BadEditionFileErr       EditionMgrInitErr = -453
	BadSectionErr           EditionMgrInitErr = -451
	BadSubPartErr           EditionMgrInitErr = -454
	ContainerAlreadyOpenWrn EditionMgrInitErr = -462
	ContainerNotFoundWrn    EditionMgrInitErr = -461
	EditionMgrInitErrValue  EditionMgrInitErr = -450
	MultiplePublisherWrn    EditionMgrInitErr = -460
	NotRegisteredSectionErr EditionMgrInitErr = -452
	NotThePublisherWrn      EditionMgrInitErr = -463
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

type EraField uint

const (
	DayField        EraField = 3
	DayOfWeekField  EraField = 7
	DayOfYearField  EraField = 8
	EraFieldValue   EraField = 0
	HourField       EraField = 4
	MinuteField     EraField = 5
	MonthField      EraField = 2
	PmField         EraField = 10
	Res1Field       EraField = 11
	Res2Field       EraField = 12
	Res3Field       EraField = 13
	SecondField     EraField = 6
	WeekOfYearField EraField = 9
	YearField       EraField = 1
)

func (e EraField) String() string {
	switch e {
	case DayField:
		return "DayField"
	case DayOfWeekField:
		return "DayOfWeekField"
	case DayOfYearField:
		return "DayOfYearField"
	case EraFieldValue:
		return "EraFieldValue"
	case HourField:
		return "HourField"
	case MinuteField:
		return "MinuteField"
	case MonthField:
		return "MonthField"
	case PmField:
		return "PmField"
	case Res1Field:
		return "Res1Field"
	case Res2Field:
		return "Res2Field"
	case Res3Field:
		return "Res3Field"
	case SecondField:
		return "SecondField"
	case WeekOfYearField:
		return "WeekOfYearField"
	case YearField:
		return "YearField"
	default:
		return fmt.Sprintf("EraField(%d)", e)
	}
}

type EraMask uint

const (
	DateStdMask    EraMask = 127
	DayMask        EraMask = 8
	DayOfWeekMask  EraMask = 128
	DayOfYearMask  EraMask = 256
	EraMaskValue   EraMask = 1
	HourMask       EraMask = 16
	MinuteMask     EraMask = 32
	MonthMask      EraMask = 4
	PmMask         EraMask = 1024
	SecondMask     EraMask = 64
	WeekOfYearMask EraMask = 512
	YearMask       EraMask = 2
)

func (e EraMask) String() string {
	switch e {
	case DateStdMask:
		return "DateStdMask"
	case DayMask:
		return "DayMask"
	case DayOfWeekMask:
		return "DayOfWeekMask"
	case DayOfYearMask:
		return "DayOfYearMask"
	case EraMaskValue:
		return "EraMaskValue"
	case HourMask:
		return "HourMask"
	case MinuteMask:
		return "MinuteMask"
	case MonthMask:
		return "MonthMask"
	case PmMask:
		return "PmMask"
	case SecondMask:
		return "SecondMask"
	case WeekOfYearMask:
		return "WeekOfYearMask"
	case YearMask:
		return "YearMask"
	default:
		return fmt.Sprintf("EraMask(%d)", e)
	}
}

type Err int

const (
	ErrAborted                   Err = -1279
	ErrAlreadyInImagingMode      Err = -5243
	ErrAttention                 Err = -1276
	ErrCannotUndo                Err = -5253
	ErrCppGeneral                Err = -32000
	ErrCppLastSystemDefinedError Err = -32020
	ErrCppLastUserDefinedError   Err = -32049
	ErrCppbad_alloc              Err = -32001
	ErrCppbad_cast               Err = -32002
	ErrCppbad_exception          Err = -32003
	ErrCppbad_typeid             Err = -32004
	ErrCppdomain_error           Err = -32006
	ErrCppinvalid_argument       Err = -32007
	ErrCppios_base_failure       Err = -32014
	ErrCpplength_error           Err = -32008
	ErrCpplogic_error            Err = -32005
	ErrCppout_of_range           Err = -32009
	ErrCppoverflow_error         Err = -32011
	ErrCpprange_error            Err = -32012
	ErrCppruntime_error          Err = -32010
	ErrCppunderflow_error        Err = -32013
	ErrDSPQueueSize              Err = -1274
	ErrEmptyScrap                Err = -5249
	ErrEndOfBody                 Err = -1813
	ErrEndOfDocument             Err = -1812
	ErrEngineNotFound            Err = -5244
	ErrFwdReset                  Err = -1275
	ErrInvalidRange              Err = -5246
	ErrIteratorReachedEnd        Err = -5245
	ErrMarginWilllNotFit         Err = -5241
	ErrNoHiliteText              Err = -5248
	ErrNonContiuousAttribute     Err = -5252
	ErrNotInImagingMode          Err = -5242
	ErrOffsetInvalid             Err = -1800
	ErrOffsetIsOutsideOfView     Err = -1801
	ErrOffsetNotOnElementBounday Err = -5247
	ErrOpenDenied                Err = -1273
	ErrOpening                   Err = -1277
	ErrReadOnlyText              Err = -5250
	ErrRefNum                    Err = -1280
	ErrState                     Err = -1278
	ErrTopOfBody                 Err = -1811
	ErrTopOfDocument             Err = -1810
	ErrUnknownAttributeTag       Err = -5240
	ErrUnknownElement            Err = -5251
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
	ErrAEEventNotHandled   ErrAE = -1708
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
	ErrAELocalOnly     ErrAE = -10016
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
	ErrAENotAnObjSpec    ErrAE = -1727
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
	ErrAEStreamBadNesting          ErrAE = -1737
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

type ErrAEEventWouldRequireUser int

const (
	ErrAEEventWouldRequireUserConsent ErrAEEventWouldRequireUser = -1744
)

func (e ErrAEEventWouldRequireUser) String() string {
	switch e {
	case ErrAEEventWouldRequireUserConsent:
		return "ErrAEEventWouldRequireUserConsent"
	default:
		return fmt.Sprintf("ErrAEEventWouldRequireUser(%d)", e)
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
	ErrASInconsistentNames      ErrAS = -2780
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
	ErrCoreEndianDataTooLongForFormat   ErrCoreEndianData = -4941
	ErrCoreEndianDataTooShortForFormat  ErrCoreEndianData = -4940
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
	ErrIAAllocationErr     ErrIA = -5381
	ErrIABufferTooSmall    ErrIA = -5384
	ErrIACanceled          ErrIA = -5385
	ErrIAEndOfTextRun      ErrIA = -5388
	ErrIAInvalidDocument   ErrIA = -5386
	ErrIANoErr             ErrIA = 0
	ErrIANoMoreItems       ErrIA = -5383
	ErrIAParamErr          ErrIA = -5382
	ErrIATextExtractionErr ErrIA = -5387
	ErrIAUnknownErr        ErrIA = -5380
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
	ErrCorruptWindowDescription            ErrInvalidWindowPtr = -5606
	ErrFloatingWindowsNotInitialized       ErrInvalidWindowPtr = -5609
	ErrInvalidWindowProperty               ErrInvalidWindowPtr = -5603
	ErrInvalidWindowPtrValue               ErrInvalidWindowPtr = -5600
	ErrInvalidWindowRef                    ErrInvalidWindowPtr = -5600
	ErrUnrecognizedWindowClass             ErrInvalidWindowPtr = -5605
	ErrUnsupportedWindowAttributesForClass ErrInvalidWindowPtr = -5601
	ErrUserWantsToDragWindow               ErrInvalidWindowPtr = -5607
	ErrWindowDoesNotFitOnscreen            ErrInvalidWindowPtr = -5611
	ErrWindowDoesNotHaveProxy              ErrInvalidWindowPtr = -5602
	ErrWindowDoesntSupportFocus            ErrInvalidWindowPtr = -30583
	ErrWindowNotFound                      ErrInvalidWindowPtr = -5610
	ErrWindowPropertyNotFound              ErrInvalidWindowPtr = -5604
	ErrWindowRegionCodeInvalid             ErrInvalidWindowPtr = -30593
	ErrWindowsAlreadyInitialized           ErrInvalidWindowPtr = -5608
	WindowAppModalStateAlreadyExistsErr    ErrInvalidWindowPtr = -5617
	WindowAttributeImmutableErr            ErrInvalidWindowPtr = -5612
	WindowAttributesConflictErr            ErrInvalidWindowPtr = -5613
	WindowGroupInvalidErr                  ErrInvalidWindowPtr = -5616
	WindowManagerInternalErr               ErrInvalidWindowPtr = -5614
	WindowNoAppModalStateErr               ErrInvalidWindowPtr = -5618
	WindowWrongStateErr                    ErrInvalidWindowPtr = -5615
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
	ErrKCAuthFailed            ErrKC = -25293
	ErrKCBufferTooSmall        ErrKC = -25301
	ErrKCCreateChainFailed     ErrKC = -25318
	ErrKCDataNotAvailable      ErrKC = -25316
	ErrKCDataNotModifiable     ErrKC = -25317
	ErrKCDataTooLarge          ErrKC = -25302
	ErrKCDuplicateCallback     ErrKC = -25297
	ErrKCDuplicateItem         ErrKC = -25299
	ErrKCDuplicateKeychain     ErrKC = -25296
	ErrKCInteractionNotAllowed ErrKC = -25308
	ErrKCInteractionRequired   ErrKC = -25315
	ErrKCInvalidCallback       ErrKC = -25298
	ErrKCInvalidItemRef        ErrKC = -25304
	ErrKCInvalidKeychain       ErrKC = -25295
	ErrKCInvalidSearchRef      ErrKC = -25305
	ErrKCItemNotFound          ErrKC = -25300
	ErrKCKeySizeNotAllowed     ErrKC = -25311
	ErrKCNoCertificateModule   ErrKC = -25313
	ErrKCNoDefaultKeychain     ErrKC = -25307
	ErrKCNoPolicyModule        ErrKC = -25314
	ErrKCNoStorageModule       ErrKC = -25312
	ErrKCNoSuchAttr            ErrKC = -25303
	ErrKCNoSuchClass           ErrKC = -25306
	ErrKCNoSuchKeychain        ErrKC = -25294
	ErrKCNotAvailable          ErrKC = -25291
	ErrKCReadOnly              ErrKC = -25292
	ErrKCReadOnlyAttr          ErrKC = -25309
	ErrKCWrongKCVersion        ErrKC = -25310
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
	ControlHandleInvalidErr      ErrMessageNotSupported = -30599
	ControlInvalidDataVersionErr ErrMessageNotSupported = -30597
	ControlPropertyInvalid       ErrMessageNotSupported = -5603
	ControlPropertyNotFoundErr   ErrMessageNotSupported = -5604
	ErrCantEmbedIntoSelf         ErrMessageNotSupported = -30594
	ErrCantEmbedRoot             ErrMessageNotSupported = -30595
	ErrControlDoesntSupportFocus ErrMessageNotSupported = -30582
	ErrControlHiddenOrDisabled   ErrMessageNotSupported = -30592
	ErrControlIsNotEmbedder      ErrMessageNotSupported = -30590
	ErrControlsAlreadyExist      ErrMessageNotSupported = -30589
	ErrCouldntSetFocus           ErrMessageNotSupported = -30585
	ErrDataNotSupported          ErrMessageNotSupported = -30581
	ErrDataSizeMismatch          ErrMessageNotSupported = -30591
	ErrInvalidPartCode           ErrMessageNotSupported = -30588
	ErrItemNotControl            ErrMessageNotSupported = -30596
	ErrMessageNotSupportedValue  ErrMessageNotSupported = -30580
	ErrNoRootControl             ErrMessageNotSupported = -30586
	ErrRootAlreadyExists         ErrMessageNotSupported = -30587
	ErrUnknownControl            ErrMessageNotSupported = -30584
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
	ErrOSABadSelector               ErrOSA = -1754
	ErrOSABadStorageType            ErrOSA = -1752
	ErrOSACantAccess                ErrOSA = -1728
	ErrOSACantAssign                ErrOSA = -10006
	ErrOSACantCoerce                ErrOSA = -1700
	ErrOSACantCreate                ErrOSA = -2710
	ErrOSACantGetTerminology        ErrOSA = -2709
	ErrOSACantLaunch                ErrOSA = -2703
	ErrOSACantOpenComponent         ErrOSA = -1762
	ErrOSACantStorePointers         ErrOSA = -1763
	ErrOSAComponentMismatch         ErrOSA = -1761
	ErrOSACorruptData               ErrOSA = -1702
	ErrOSACorruptTerminology        ErrOSA = -2705
	ErrOSADataBlockTooLarge         ErrOSA = -2708
	ErrOSADataFormatObsolete        ErrOSA = -1758
	ErrOSADataFormatTooNew          ErrOSA = -1759
	ErrOSADivideByZero              ErrOSA = -2701
	ErrOSAGeneralError              ErrOSA = -2700
	ErrOSAInternalTableOverflow     ErrOSA = -2707
	ErrOSAInvalidID                 ErrOSA = -1751
	ErrOSANoSuchDialect             ErrOSA = -1757
	ErrOSANumericOverflow           ErrOSA = -2702
	ErrOSARecordingIsAlreadyOn      ErrOSA = -1732
	ErrOSAScriptError               ErrOSA = -1753
	ErrOSASourceNotAvailable        ErrOSA = -1756
	ErrOSAStackOverflow             ErrOSA = -2706
	ErrOSASystemError               ErrOSA = -1750
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
	ErrOSATypeErrorValue        ErrOSATypeError = -1703
	OSAControlFlowError         ErrOSATypeError = -2755
	OSADuplicateHandler         ErrOSATypeError = -2752
	OSADuplicateParameter       ErrOSATypeError = -2750
	OSADuplicateProperty        ErrOSATypeError = -2751
	OSAIllegalAccess            ErrOSATypeError = -1723
	OSAIllegalAssign            ErrOSATypeError = -10003
	OSAIllegalIndex             ErrOSATypeError = -1719
	OSAIllegalRange             ErrOSATypeError = -1720
	OSAInconsistentDeclarations ErrOSATypeError = -2754
	OSAMessageNotUnderstood     ErrOSATypeError = -1708
	OSAMissingParameter         ErrOSATypeError = -1701
	OSAParameterMismatch        ErrOSATypeError = -1721
	OSASyntaxError              ErrOSATypeError = -2740
	OSASyntaxTypeError          ErrOSATypeError = -2741
	OSATokenTooLong             ErrOSATypeError = -2742
	OSAUndefinedHandler         ErrOSATypeError = -1717
	OSAUndefinedVariable        ErrOSATypeError = -2753
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

type ErrTaskNot int

const (
	ErrTaskNotFound ErrTaskNot = -10780
)

func (e ErrTaskNot) String() string {
	switch e {
	case ErrTaskNotFound:
		return "ErrTaskNotFound"
	default:
		return fmt.Sprintf("ErrTaskNot(%d)", e)
	}
}

type ErrWS int

const (
	// ErrWSInternalError: An internal framework error occured.
	ErrWSInternalError ErrWS = -65793
	// ErrWSParseError: The server response was not valid XML.
	ErrWSParseError ErrWS = -65795
	// ErrWSTimeoutError: The method invocation timed out.
	ErrWSTimeoutError ErrWS = -65796
	// ErrWSTransportError: A network error occured.
	ErrWSTransportError ErrWS = -65794
)

func (e ErrWS) String() string {
	switch e {
	case ErrWSInternalError:
		return "ErrWSInternalError"
	case ErrWSParseError:
		return "ErrWSParseError"
	case ErrWSTimeoutError:
		return "ErrWSTimeoutError"
	case ErrWSTransportError:
		return "ErrWSTransportError"
	default:
		return fmt.Sprintf("ErrWS(%d)", e)
	}
}

type EvtNot uint

const (
	EvtNotEnb EvtNot = 1
)

func (e EvtNot) String() string {
	switch e {
	case EvtNotEnb:
		return "EvtNotEnb"
	default:
		return fmt.Sprintf("EvtNot(%d)", e)
	}
}

type FFormatOK uint

const (
	// Deprecated.
	FBadPartsTable FFormatOK = 10
	// Deprecated.
	FBestGuess FFormatOK = 1
	// Deprecated.
	FEmptyFormatString FFormatOK = 13
	// Deprecated.
	FExtraDecimal FFormatOK = 5
	// Deprecated.
	FExtraExp FFormatOK = 7
	// Deprecated.
	FExtraPercent FFormatOK = 11
	// Deprecated.
	FExtraSeparator FFormatOK = 12
	// Deprecated.
	FFormStrIsNAN FFormatOK = 9
	// Deprecated.
	FFormatOKValue FFormatOK = 0
	// Deprecated.
	FFormatOverflow FFormatOK = 8
	// Deprecated.
	FMissingDelimiter FFormatOK = 4
	// Deprecated.
	FMissingLiteral FFormatOK = 6
	// Deprecated.
	FOutOfSynch FFormatOK = 2
	// Deprecated.
	FSpuriousChars FFormatOK = 3
)

func (e FFormatOK) String() string {
	switch e {
	case FBadPartsTable:
		return "FBadPartsTable"
	case FBestGuess:
		return "FBestGuess"
	case FEmptyFormatString:
		return "FEmptyFormatString"
	case FExtraDecimal:
		return "FExtraDecimal"
	case FExtraExp:
		return "FExtraExp"
	case FExtraPercent:
		return "FExtraPercent"
	case FExtraSeparator:
		return "FExtraSeparator"
	case FFormStrIsNAN:
		return "FFormStrIsNAN"
	case FFormatOKValue:
		return "FFormatOKValue"
	case FFormatOverflow:
		return "FFormatOverflow"
	case FMissingDelimiter:
		return "FMissingDelimiter"
	case FMissingLiteral:
		return "FMissingLiteral"
	case FOutOfSynch:
		return "FOutOfSynch"
	case FSpuriousChars:
		return "FSpuriousChars"
	default:
		return fmt.Sprintf("FFormatOK(%d)", e)
	}
}

type FOnDesk uint

const (
	FHasBundle   FOnDesk = 8192
	FInvisible   FOnDesk = 16384
	FOnDeskValue FOnDesk = 1
)

func (e FOnDesk) String() string {
	switch e {
	case FHasBundle:
		return "FHasBundle"
	case FInvisible:
		return "FInvisible"
	case FOnDeskValue:
		return "FOnDeskValue"
	default:
		return fmt.Sprintf("FOnDesk(%d)", e)
	}
}

type FPositive uint

const (
	// Deprecated.
	FNegative FPositive = 1
	// Deprecated.
	FPositiveValue FPositive = 0
	// Deprecated.
	FZero FPositive = 2
)

func (e FPositive) String() string {
	switch e {
	case FNegative:
		return "FNegative"
	case FPositiveValue:
		return "FPositiveValue"
	case FZero:
		return "FZero"
	default:
		return fmt.Sprintf("FPositive(%d)", e)
	}
}

type FTrash int

const (
	FDesktop    FTrash = -2
	FDisk       FTrash = 0
	FTrashValue FTrash = -3
)

func (e FTrash) String() string {
	switch e {
	case FDesktop:
		return "FDesktop"
	case FDisk:
		return "FDisk"
	case FTrashValue:
		return "FTrashValue"
	default:
		return fmt.Sprintf("FTrash(%d)", e)
	}
}

type FV uint

const (
	// Deprecated.
	FVNumber FV = 0
)

func (e FV) String() string {
	switch e {
	case FVNumber:
		return "FVNumber"
	default:
		return fmt.Sprintf("FV(%d)", e)
	}
}

type False32b uint

const (
	// Deprecated.
	False32bValue False32b = 0
	// Deprecated.
	True32b False32b = 1
)

func (e False32b) String() string {
	switch e {
	case False32bValue:
		return "False32bValue"
	case True32b:
		return "True32b"
	default:
		return fmt.Sprintf("False32b(%d)", e)
	}
}

type FatalDateTime uint

const (
	CantReadUtilities  FatalDateTime = 33280
	DateTimeInvalid    FatalDateTime = 34816
	DateTimeNotFound   FatalDateTime = 33792
	ExtraneousStrings  FatalDateTime = 16
	FatalDateTimeValue FatalDateTime = 32768
	FieldOrderNotIntl  FatalDateTime = 8
	LeftOverChars      FatalDateTime = 2
	LongDateFound      FatalDateTime = 1
	SepNotConsistent   FatalDateTime = 64
	SepNotIntlSep      FatalDateTime = 4
	TokenErr           FatalDateTime = 33024
	TooManySeps        FatalDateTime = 32
)

func (e FatalDateTime) String() string {
	switch e {
	case CantReadUtilities:
		return "CantReadUtilities"
	case DateTimeInvalid:
		return "DateTimeInvalid"
	case DateTimeNotFound:
		return "DateTimeNotFound"
	case ExtraneousStrings:
		return "ExtraneousStrings"
	case FatalDateTimeValue:
		return "FatalDateTimeValue"
	case FieldOrderNotIntl:
		return "FieldOrderNotIntl"
	case LeftOverChars:
		return "LeftOverChars"
	case LongDateFound:
		return "LongDateFound"
	case SepNotConsistent:
		return "SepNotConsistent"
	case SepNotIntlSep:
		return "SepNotIntlSep"
	case TokenErr:
		return "TokenErr"
	case TooManySeps:
		return "TooManySeps"
	default:
		return fmt.Sprintf("FatalDateTime(%d)", e)
	}
}

type FidNotFound int

const (
	BadFCBErr                       FidNotFound = -1327
	BadFidErr                       FidNotFound = -1307
	CatChangedErr                   FidNotFound = -1304
	DesktopDamagedErr               FidNotFound = -1305
	DiffVolErr                      FidNotFound = -1303
	EnvBadVers                      FidNotFound = -5501
	EnvNotPresent                   FidNotFound = -5500
	EnvVersTooBig                   FidNotFound = -5502
	ErrFSAttributeNotFound          FidNotFound = -1427
	ErrFSBadAllocFlags              FidNotFound = -1413
	ErrFSBadBuffer                  FidNotFound = -1403
	ErrFSBadFSRef                   FidNotFound = -1401
	ErrFSBadForkName                FidNotFound = -1402
	ErrFSBadForkRef                 FidNotFound = -1404
	ErrFSBadInfoBitmap              FidNotFound = -1405
	ErrFSBadItemCount               FidNotFound = -1418
	ErrFSBadIteratorFlags           FidNotFound = -1422
	ErrFSBadPosMode                 FidNotFound = -1412
	ErrFSBadSearchParams            FidNotFound = -1419
	ErrFSForkExists                 FidNotFound = -1421
	ErrFSForkNotFound               FidNotFound = -1409
	ErrFSIteratorNotFound           FidNotFound = -1423
	ErrFSIteratorNotSupported       FidNotFound = -1424
	ErrFSMissingCatInfo             FidNotFound = -1406
	ErrFSMissingName                FidNotFound = -1411
	ErrFSNameTooLong                FidNotFound = -1410
	ErrFSNoMoreItems                FidNotFound = -1417
	ErrFSNotAFolder                 FidNotFound = -1407
	ErrFSNotEnoughSpaceForOperation FidNotFound = -1429
	ErrFSOperationNotSupported      FidNotFound = -1426
	ErrFSPropertyNotValid           FidNotFound = -1428
	ErrFSQuotaExceeded              FidNotFound = -1425
	ErrFSRefsDifferent              FidNotFound = -1420
	ErrFSUnknownCall                FidNotFound = -1400
	FidExists                       FidNotFound = -1301
	FidNotFoundValue                FidNotFound = -1300
	FileBoundsErr                   FidNotFound = -1309
	FirstDskErr                     FidNotFound = -84
	FontDecError                    FidNotFound = -64
	FontNotDeclared                 FidNotFound = -65
	FontNotOutlineErr               FidNotFound = -32615
	FontSubErr                      FidNotFound = -66
	FsDataTooBigErr                 FidNotFound = -1310
	LastDskErr                      FidNotFound = -64
	NoDriveErr                      FidNotFound = -64
	NoNybErr                        FidNotFound = -66
	NotAFileErr                     FidNotFound = -1302
	NotARemountErr                  FidNotFound = -1308
	OffLinErr                       FidNotFound = -65
	SameFileErr                     FidNotFound = -1306
	VolVMBusyErr                    FidNotFound = -1311
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
	BadProfileError        FirstPickerError = -4008
	CantCreatePickerWindow FirstPickerError = -4004
	CantLoadPackage        FirstPickerError = -4005
	CantLoadPicker         FirstPickerError = -4003
	ColorSyncNotInstalled  FirstPickerError = -4007
	FirstPickerErrorValue  FirstPickerError = -4000
	InvalidPickerType      FirstPickerError = -4000
	NoHelpForItem          FirstPickerError = -4009
	PickerCantLive         FirstPickerError = -4006
	PickerResourceError    FirstPickerError = -4002
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

type Fs uint

const (
	FsAtMark       Fs = 0
	FsCurPerm      Fs = 0
	FsFromLEOF     Fs = 2
	FsFromMark     Fs = 3
	FsFromStart    Fs = 1
	FsRdAccessPerm Fs = 1
	FsRdDenyPerm   Fs = 16
	FsRdPerm       Fs = 1
	FsRdWrPerm     Fs = 3
	FsRdWrShPerm   Fs = 4
	FsWrAccessPerm Fs = 2
	FsWrDenyPerm   Fs = 32
	FsWrPerm       Fs = 2
)

func (e Fs) String() string {
	switch e {
	case FsAtMark:
		return "FsAtMark"
	case FsFromLEOF:
		return "FsFromLEOF"
	case FsFromMark:
		return "FsFromMark"
	case FsFromStart:
		return "FsFromStart"
	case FsRdDenyPerm:
		return "FsRdDenyPerm"
	case FsRdWrShPerm:
		return "FsRdWrShPerm"
	case FsWrDenyPerm:
		return "FsWrDenyPerm"
	default:
		return fmt.Sprintf("Fs(%d)", e)
	}
}

type FsRt uint

const (
	FsRtDirID FsRt = 2
	FsRtParID FsRt = 1
)

func (e FsRt) String() string {
	switch e {
	case FsRtDirID:
		return "FsRtDirID"
	case FsRtParID:
		return "FsRtParID"
	default:
		return fmt.Sprintf("FsRt(%d)", e)
	}
}

type FsSB uint

const (
	FsSBAccessDate             FsSB = 131072
	FsSBAccessDateBit          FsSB = 17
	FsSBAttributeModDate       FsSB = 65536
	FsSBAttributeModDateBit    FsSB = 16
	FsSBDrBkDat                FsSB = 2048
	FsSBDrBkDatBit             FsSB = 11
	FsSBDrCrDat                FsSB = 512
	FsSBDrCrDatBit             FsSB = 9
	FsSBDrFndrInfo             FsSB = 4096
	FsSBDrFndrInfoBit          FsSB = 12
	FsSBDrMdDat                FsSB = 1024
	FsSBDrMdDatBit             FsSB = 10
	FsSBDrNmFls                FsSB = 16
	FsSBDrNmFlsBit             FsSB = 4
	FsSBDrParID                FsSB = 8192
	FsSBDrParIDBit             FsSB = 13
	FsSBDrUsrWds               FsSB = 8
	FsSBDrUsrWdsBit            FsSB = 3
	FsSBFlAttrib               FsSB = 4
	FsSBFlAttribBit            FsSB = 2
	FsSBFlBkDat                FsSB = 2048
	FsSBFlBkDatBit             FsSB = 11
	FsSBFlCrDat                FsSB = 512
	FsSBFlCrDatBit             FsSB = 9
	FsSBFlFndrInfo             FsSB = 8
	FsSBFlFndrInfoBit          FsSB = 3
	FsSBFlLgLen                FsSB = 32
	FsSBFlLgLenBit             FsSB = 5
	FsSBFlMdDat                FsSB = 1024
	FsSBFlMdDatBit             FsSB = 10
	FsSBFlParID                FsSB = 8192
	FsSBFlParIDBit             FsSB = 13
	FsSBFlPyLen                FsSB = 64
	FsSBFlPyLenBit             FsSB = 6
	FsSBFlRLgLen               FsSB = 128
	FsSBFlRLgLenBit            FsSB = 7
	FsSBFlRPyLen               FsSB = 256
	FsSBFlRPyLenBit            FsSB = 8
	FsSBFlXFndrInfo            FsSB = 4096
	FsSBFlXFndrInfoBit         FsSB = 12
	FsSBFullName               FsSB = 2
	FsSBFullNameBit            FsSB = 1
	FsSBGroupID                FsSB = 4194304
	FsSBGroupIDBit             FsSB = 22
	FsSBNegate                 FsSB = 16384
	FsSBNegateBit              FsSB = 14
	FsSBNodeID                 FsSB = 32768
	FsSBNodeIDBit              FsSB = 15
	FsSBPartialName            FsSB = 1
	FsSBPartialNameBit         FsSB = 0
	FsSBPermissions            FsSB = 262144
	FsSBPermissionsBit         FsSB = 18
	FsSBSkipHiddenItems        FsSB = 1048576
	FsSBSkipHiddenItemsBit     FsSB = 20
	FsSBSkipPackageContents    FsSB = 524288
	FsSBSkipPackageContentsBit FsSB = 19
	FsSBUserID                 FsSB = 2097152
	FsSBUserIDBit              FsSB = 21
)

func (e FsSB) String() string {
	switch e {
	case FsSBAccessDate:
		return "FsSBAccessDate"
	case FsSBAccessDateBit:
		return "FsSBAccessDateBit"
	case FsSBAttributeModDate:
		return "FsSBAttributeModDate"
	case FsSBAttributeModDateBit:
		return "FsSBAttributeModDateBit"
	case FsSBDrBkDat:
		return "FsSBDrBkDat"
	case FsSBDrBkDatBit:
		return "FsSBDrBkDatBit"
	case FsSBDrCrDat:
		return "FsSBDrCrDat"
	case FsSBDrCrDatBit:
		return "FsSBDrCrDatBit"
	case FsSBDrFndrInfo:
		return "FsSBDrFndrInfo"
	case FsSBDrFndrInfoBit:
		return "FsSBDrFndrInfoBit"
	case FsSBDrMdDat:
		return "FsSBDrMdDat"
	case FsSBDrMdDatBit:
		return "FsSBDrMdDatBit"
	case FsSBDrNmFlsBit:
		return "FsSBDrNmFlsBit"
	case FsSBDrParID:
		return "FsSBDrParID"
	case FsSBDrParIDBit:
		return "FsSBDrParIDBit"
	case FsSBDrUsrWds:
		return "FsSBDrUsrWds"
	case FsSBDrUsrWdsBit:
		return "FsSBDrUsrWdsBit"
	case FsSBFlAttribBit:
		return "FsSBFlAttribBit"
	case FsSBFlLgLen:
		return "FsSBFlLgLen"
	case FsSBFlLgLenBit:
		return "FsSBFlLgLenBit"
	case FsSBFlPyLen:
		return "FsSBFlPyLen"
	case FsSBFlPyLenBit:
		return "FsSBFlPyLenBit"
	case FsSBFlRLgLen:
		return "FsSBFlRLgLen"
	case FsSBFlRLgLenBit:
		return "FsSBFlRLgLenBit"
	case FsSBFlRPyLen:
		return "FsSBFlRPyLen"
	case FsSBFullNameBit:
		return "FsSBFullNameBit"
	case FsSBGroupID:
		return "FsSBGroupID"
	case FsSBGroupIDBit:
		return "FsSBGroupIDBit"
	case FsSBNegate:
		return "FsSBNegate"
	case FsSBNegateBit:
		return "FsSBNegateBit"
	case FsSBNodeID:
		return "FsSBNodeID"
	case FsSBNodeIDBit:
		return "FsSBNodeIDBit"
	case FsSBPartialNameBit:
		return "FsSBPartialNameBit"
	case FsSBPermissions:
		return "FsSBPermissions"
	case FsSBPermissionsBit:
		return "FsSBPermissionsBit"
	case FsSBSkipHiddenItems:
		return "FsSBSkipHiddenItems"
	case FsSBSkipHiddenItemsBit:
		return "FsSBSkipHiddenItemsBit"
	case FsSBSkipPackageContents:
		return "FsSBSkipPackageContents"
	case FsSBSkipPackageContentsBit:
		return "FsSBSkipPackageContentsBit"
	case FsSBUserID:
		return "FsSBUserID"
	case FsSBUserIDBit:
		return "FsSBUserIDBit"
	default:
		return fmt.Sprintf("FsSB(%d)", e)
	}
}

type FsUnix uint

const (
	FsUnixPriv FsUnix = 1
)

func (e FsUnix) String() string {
	switch e {
	case FsUnixPriv:
		return "FsUnixPriv"
	default:
		return fmt.Sprintf("FsUnix(%d)", e)
	}
}

type Fsm int

const (
	FsmBadFFSNameErr        Fsm = -433
	FsmBadFSDLenErr         Fsm = -434
	FsmBadFSDVersionErr     Fsm = -436
	FsmBusyFFSErr           Fsm = -432
	FsmDuplicateFSIDErr     Fsm = -435
	FsmFFSNotFoundErr       Fsm = -431
	FsmNoAlternateStackErr  Fsm = -437
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
	AppleMenuFolderIconResource      GenericDocumentIconResource = -3982
	DesktopIconResource              GenericDocumentIconResource = -3992
	FloppyIconResource               GenericDocumentIconResource = -3998
	GenericApplicationIconResource   GenericDocumentIconResource = -3996
	GenericCDROMIconResource         GenericDocumentIconResource = -3987
	GenericDeskAccessoryIconResource GenericDocumentIconResource = -3991
	GenericDocumentIconResourceValue GenericDocumentIconResource = -4000
	GenericEditionFileIconResource   GenericDocumentIconResource = -3989
	GenericExtensionIconResource     GenericDocumentIconResource = -16415
	GenericFileServerIconResource    GenericDocumentIconResource = -3972
	GenericFolderIconResource        GenericDocumentIconResource = -3999
	GenericHardDiskIconResource      GenericDocumentIconResource = -3995
	GenericMoverObjectIconResource   GenericDocumentIconResource = -3969
	GenericPreferencesIconResource   GenericDocumentIconResource = -3971
	GenericQueryDocumentIconResource GenericDocumentIconResource = -16506
	GenericRAMDiskIconResource       GenericDocumentIconResource = -3988
	GenericStationeryIconResource    GenericDocumentIconResource = -3985
	GenericSuitcaseIconResource      GenericDocumentIconResource = -3970
	OpenFolderIconResource           GenericDocumentIconResource = -3997
	PrivateFolderIconResource        GenericDocumentIconResource = -3994
	SystemFolderIconResource         GenericDocumentIconResource = -3983
	TrashIconResource                GenericDocumentIconResource = -3993
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
	// Gestalt16BitAudioSupport: # Discussion
	Gestalt16BitAudioSupport int = 12
	// Gestalt16BitSoundIO: # Discussion
	Gestalt16BitSoundIO    int = 7
	Gestalt20thAnniversary int = 512
	// Gestalt32BitAddressing: If `true`, the operating system is using 32-bit addressing mode.
	Gestalt32BitAddressing int = 0
	// Gestalt32BitCapable: If `true`, Machine is 32-bit capable.
	Gestalt32BitCapable int = 2
	Gestalt32BitQD      int = 512
	Gestalt32BitQD11    int = 513
	Gestalt32BitQD12    int = 544
	Gestalt32BitQD13    int = 560
	// Gestalt32BitSysZone: If `true`, there is a 32-bit compatible system zone.
	Gestalt32BitSysZone int = 1
	Gestalt68000        int = 1
	Gestalt68010        int = 2
	Gestalt68020        int = 3
	Gestalt68030        int = 4
	Gestalt68030MMU     int = 3
	Gestalt68040        int = 5
	Gestalt68040FPU     int = 3
	Gestalt68040MMU     int = 4
	Gestalt68851        int = 2
	Gestalt68881        int = 1
	Gestalt68882        int = 2
	// Gestalt68k: If the [Gestalt] function returns `gestalt68k`, the system is a MC680x0 Macintosh.
	Gestalt68k         int = 1
	Gestalt8BitQD      int = 256
	GestaltADBISOKbdII int = 11
	GestaltADBKbdII    int = 10
	GestaltAMU         int = 1
	// GestaltATSUUpdate1: Indicates that version 1.1 of ATSUI is installed on the user’s system.
	GestaltATSUUpdate1 int = 131072
	// GestaltATSUUpdate2: Indicates that version 1.2 of ATSUI is installed on the user’s system.
	GestaltATSUUpdate2 int = 196608
	// GestaltATSUUpdate3: Indicates that version 2.0 of ATSUI is installed on the user’s system.
	GestaltATSUUpdate3 int = 262144
	// GestaltATSUUpdate4: Indicates that ATSUI for a version of macOS from 10.0.1 through 10.0.4 is installed on the user’s system.
	GestaltATSUUpdate4 int = 327680
	// GestaltATSUUpdate5: Indicates that version 2.3 of ATSUI is installed on the user’s system.
	GestaltATSUUpdate5 int = 393216
	// GestaltATSUUpdate6: Indicates that version 2.4 of ATSUI is installed on the user’s system.
	GestaltATSUUpdate6 int = 458752
	// GestaltATSUUpdate7: Indicates that version 2.5 of ATSUI is installed on the user’s system.
	GestaltATSUUpdate7 int = 524288
	// GestaltATSUVersion: Specifies the version of ATSUI installed on the user’s system.
	GestaltATSUVersion int = 'u'<<24 | 'i'<<16 | 's'<<8 | 'v' // 'uisv'
	GestaltAWS6150_60  int = 75
	GestaltAWS6150_66  int = 100
	GestaltAWS8150_110 int = 40
	GestaltAWS8150_80  int = 65
	GestaltAWS8550     int = 68
	GestaltAWS9150_120 int = 57
	GestaltAWS9150_80  int = 39
	// GestaltAddressingModeAttr: The [Gestalt] selector you pass to determine the addressing mode attributes that are present.
	GestaltAddressingModeAttr int = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'addr'
	// GestaltAdminFeaturesFlagsAttr: The [Gestalt] selector you pass to determine the admin features that are present.
	GestaltAdminFeaturesFlagsAttr int = 'f'<<24 | 'r'<<16 | 'e'<<8 | 'd' // 'fred'
	GestaltAllegroQD              int = 592
	// GestaltAllegroQDText: This is the version of QuickDraw Text used with Mac OS 8.2 and up.
	GestaltAllegroQDText                       int = 256
	GestaltAltivecRegistersSwappedCorrectlyBit int = 20
	// GestaltAntiAliasedTextAvailable: Capable of antialiased text.
	GestaltAntiAliasedTextAvailable int = 2
	GestaltAppleAdjustADBKbd        int = 15
	GestaltAppleAdjustISOKbd        int = 16
	GestaltAppleAdjustKeypad        int = 14
	// GestaltAppleEventsAttr: # Discussion
	GestaltAppleEventsAttr int = 'e'<<24 | 'v'<<16 | 'n'<<8 | 't' // 'evnt'
	// GestaltAppleEventsPresent: A [Gestalt] attribute constant.
	GestaltAppleEventsPresent int = 0
	GestaltAppleGuideIsDebug  int = 30
	GestaltAppleGuidePresent  int = 31
	GestaltArbitorAttr        int = 'a'<<24 | 'r'<<16 | 'b'<<8 | ' ' // 'arb '
	GestaltArm                int = 20
	GestaltAsyncSCSI          int = 0
	GestaltAsyncSCSIINROM     int = 1
	// GestaltBuiltInSoundInput: Set if a built-in sound input device is available.
	GestaltBuiltInSoundInput int = 4
	GestaltCPU601            int = 257
	GestaltCPU603            int = 259
	GestaltCPU603e           int = 262
	GestaltCPU603ev          int = 263
	GestaltCPU604            int = 260
	GestaltCPU604e           int = 265
	GestaltCPU604ev          int = 266
	GestaltCPU68000          int = 0
	GestaltCPU68010          int = 1
	GestaltCPU68020          int = 2
	GestaltCPU68030          int = 3
	GestaltCPU68040          int = 4
	GestaltCPU750            int = 264
	GestaltCPUG4             int = 268
	GestaltCPUG47450         int = 272
	// GestaltCanStartDragInFloatWindow: If the bit specified by this mask is set, the Drag Manager can start a drag in a floating window.
	GestaltCanStartDragInFloatWindow int = 4
	GestaltCanUseCGTextRendering     int = 6
	GestaltCardServicesPresent       int = 0
	GestaltClassic                   int = 1
	GestaltClassicII                 int = 23
	// GestaltColorMatchingAttr: The selector for obtaining version information.
	GestaltColorMatchingAttr int = 'c'<<24 | 'm'<<16 | 't'<<8 | 'a' // 'cmta'
	// GestaltColorMatchingLibLoaded: This constant is provided for backward compatibility only.
	GestaltColorMatchingLibLoaded int = 1
	GestaltCreatesAliasFontRsrc   int = 4
	GestaltCurrentGraphicsVersion int = 66048
	// GestaltDTMgrSupportsFSM: If this bit is set in the `response` parameter, the desktop database supports File System Manager-based foreign file systems.
	GestaltDTMgrSupportsFSM int = 7
	// GestaltDesktopSpeechRecognition: If this bit is set, the Speech Recognition Manager supports the desktop microphone.
	GestaltDesktopSpeechRecognition int = 1
	// GestaltDragMgrAttr: The Gestalt selector passed to determine what features of the Drag Manager are present.
	GestaltDragMgrAttr int = 'd'<<24 | 'r'<<16 | 'a'<<8 | 'g' // 'drag'
	// GestaltDragMgrFloatingWind: If the bit specified by this mask is set, the Drag Manager floating window support functions are available.
	GestaltDragMgrFloatingWind int = 1
	// GestaltDragMgrHasImageSupport: If the bit specified by this mask is set, the Drag Manager image support functions are available.
	GestaltDragMgrHasImageSupport int = 3
	// GestaltDragMgrPresent: If the bit specified by this mask is set, the Drag Manager functions are available.
	GestaltDragMgrPresent int = 0
	// GestaltDupSelectorErr: Specifies you tried to add an entry that already existed.
	GestaltDupSelectorErr int = -5552
	GestaltEMMU1          int = 5
	GestaltExtADBKbd      int = 4
	GestaltExtISOADBKbd   int = 9
	// GestaltExtendedTimeMgr: If this bit is set, the extended Time Manager is present.
	GestaltExtendedTimeMgr              int = 3
	GestaltExtendedWindowAttributes     int = 1
	GestaltExtendedWindowAttributesBit  int = 1
	GestaltExtendedWindowAttributesMask int = 2
	GestaltFBCCurrentVersion            int = 17
	GestaltFBCVersion                   int = 'f'<<24 | 'b'<<16 | 'c'<<8 | 'v' // 'fbcv'
	// GestaltFPUType: A constant that represents the type of floating-point unit currently installed, if any.
	GestaltFPUType                   int = 'f'<<24 | 'p'<<16 | 'u'<<8 | ' ' // 'fpu '
	GestaltFSAllowsConcurrentAsyncIO int = 17
	// GestaltFSAttr: A selector you pass to the [Gestalt] function.
	GestaltFSAttr int = 'f'<<24 | 's'<<16 | ' '<<8 | ' ' // 'fs  '
	// GestaltFSIncompatibleDFA82: If this bit is set in the `response` parameter, VCB and FCB structures are changed; DFA 8.2 is incompatible.
	GestaltFSIncompatibleDFA82 int = 10
	// GestaltFSMDoesDynamicLoad: If this bit is set in the `response` parameter, the File System Manager supports dynamic loading of external file system code resources.
	GestaltFSMDoesDynamicLoad int = 3
	// GestaltFSNoMFSVols: If this bit is set in the `response` parameter, the file system does not support MFS volumes.
	GestaltFSNoMFSVols int = 8
	// GestaltFSSupports2TBVols: If this bit is set in the `response` parameter, the file system supports 2 terabyte volumes.
	GestaltFSSupports2TBVols int = 5
	// GestaltFSSupports4GBVols: If this bit is set in the `response` parameter, the file system supports 4 gigabyte volumes.
	GestaltFSSupports4GBVols        int = 4
	GestaltFSSupportsExclusiveLocks int = 15
	// GestaltFSSupportsHFSPlusVols: If this bit is set in the `response` parameter, the file system supports HFS Plus volumes.
	GestaltFSSupportsHFSPlusVols         int = 9
	GestaltFSSupportsHardLinkDetection   int = 16
	GestaltFSUsesPOSIXPathsForConversion int = 14
	GestaltFileAllocationZeroedBlocksBit int = 4
	// GestaltFindFolderAttr: The selector you pass to the [Gestalt] function to determine the  [FindFolder] function attributes.
	GestaltFindFolderAttr                           int = 'f'<<24 | 'o'<<16 | 'l'<<8 | 'd' // 'fold'
	GestaltFindFolderPresent                        int = 0
	GestaltFinderAttr                               int = 'f'<<24 | 'n'<<16 | 'd'<<8 | 'r' // 'fndr'
	GestaltFinderCallsAEProcess                     int = 2
	GestaltFinderDropEvent                          int = 0
	GestaltFinderFloppyRootComments                 int = 8
	GestaltFinderFullDragManagerSupport             int = 7
	GestaltFinderHasClippings                       int = 6
	GestaltFinderLargeAndNotSavedFlavorsOK          int = 9
	GestaltFinderMagicPlacement                     int = 1
	GestaltFinderSupports4GBVolumes                 int = 4
	GestaltFinderUnderstandsRedirectedDesktopFolder int = 11
	GestaltFinderUsesExtensibleFolderManager        int = 10
	// GestaltFinderUsesSpecialOpenFoldersFile: Specifies that the Finder uses a special file to store the list of open folders.
	GestaltFinderUsesSpecialOpenFoldersFile int = 0
	// GestaltFolderDescSupport: If this bit is set, the extended Folder Manager functionality supporting folder descriptors and routings is available.
	GestaltFolderDescSupport                    int = 1
	GestaltFolderMgrFollowsAliasesWhenResolving int = 2
	GestaltFolderMgrSupportsDomains             int = 4
	GestaltFolderMgrSupportsExtendedCalls       int = 3
	GestaltFolderMgrSupportsFSCalls             int = 5
	// GestaltFontMgrAttr: The [Gestalt] selector you pass to determine which Font Manager attributes are present.
	GestaltFontMgrAttr                int = 'f'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'font'
	GestaltFrontWindowMayBeHiddenBit  int = 8
	GestaltFrontWindowMayBeHiddenMask int = 256
	// GestaltFullExtFSDispatching: # Discussion
	GestaltFullExtFSDispatching int = 0
	GestaltGraphicsVersion      int = 'g'<<24 | 'r'<<16 | 'f'<<8 | 'x' // 'grfx'
	// GestaltHardwareAttr: # Discussion
	GestaltHardwareAttr     int = 'h'<<24 | 'd'<<16 | 'w'<<8 | 'r' // 'hdwr'
	GestaltHasASC           int = 3
	GestaltHasColor         int = 0
	GestaltHasDeepGWorlds   int = 1
	GestaltHasDirectPixMaps int = 2
	GestaltHasEnhancedLtalk int = 30
	// GestaltHasExtendedDiskInit: # Discussion
	GestaltHasExtendedDiskInit int = 6
	GestaltHasFMTuner          int = 7
	// GestaltHasFSSpecCalls: If this bit is set in the `response` parameter, the operating environment provides the file system specification ([FSSpec]) versions of the basic file-manipulation functions, as well as the [FSMakeFSSpec] function.
	GestaltHasFSSpecCalls int = 1
	// GestaltHasFileSystemManager: If this bit is set in the `response` parameter, the File System Manager is present.
	GestaltHasFileSystemManager   int = 2
	GestaltHasFloatingWindows     int = 2
	GestaltHasFloatingWindowsBit  int = 2
	GestaltHasFloatingWindowsMask int = 4
	GestaltHasGPIaToDCDa          int = 0
	GestaltHasGPIaToRTxCa         int = 1
	GestaltHasGPIbToDCDb          int = 2
	GestaltHasGrayishTextOr       int = 3
	// GestaltHasHFSPlusAPIs: # Discussion
	GestaltHasHFSPlusAPIs        int = 12
	GestaltHasHWClosedCaptioning int = 2
	GestaltHasIRRemote           int = 3
	GestaltHasParityCapability   int = 0
	GestaltHasResourceOverrides  int = 1
	// GestaltHasSCC: # Discussion
	GestaltHasSCC int = 4
	// GestaltHasSCSI: The `gestaltHasSCSI` bit means the machine is equipped with a SCSI implementation based on the 53C80 chip, which was introduced in the Macintosh Plus.
	GestaltHasSCSI int = 7
	// GestaltHasSCSI961: This bit is set if the machine has a SCSI implementation based on the 53C96 chip installed on an internal  bus.
	GestaltHasSCSI961 int = 21
	// GestaltHasSCSI962: This bit is set if the machine has a SCSI implementation based on the 53C96 chip installed on an  external bus.
	GestaltHasSCSI962      int = 22
	GestaltHasSerialFader  int = 6
	GestaltHasSoftPowerOff int = 19
	GestaltHasSoundFader   int = 1
	// GestaltHasSoundInputDevice: Set if a sound input device is available.
	GestaltHasSoundInputDevice    int = 5
	GestaltHasStereoDecoder       int = 5
	GestaltHasSystemIRFunction    int = 8
	GestaltHasTVTuner             int = 0
	GestaltHasUniversalROM        int = 24
	GestaltHasVIA1                int = 0
	GestaltHasVIA2                int = 1
	GestaltHasVidDecoderScaler    int = 4
	GestaltHasWindowBuffering     int = 3
	GestaltHasWindowBufferingBit  int = 3
	GestaltHasWindowBufferingMask int = 8
	GestaltHasWindowShadowsBit    int = 6
	GestaltHasWindowShadowsMask   int = 64
	GestaltHasZoomedVideo         int = 11
	// GestaltHelpMgrAttr: The selector you pass to the [Gestalt] function to determine the Help Manager attributes.
	GestaltHelpMgrAttr       int = 'h'<<24 | 'e'<<16 | 'l'<<8 | 'p' // 'help'
	GestaltHelpMgrExtensions int = 1
	GestaltHelpMgrPresent    int = 0
	GestaltHidePortA         int = 3
	GestaltHidePortB         int = 4
	// GestaltHighLevelMatching: This constant is provided for backward compatibility only.
	GestaltHighLevelMatching        int = 0
	GestaltINeedIRPowerOffConfirm   int = 10
	GestaltIPCSupport               int = 7
	GestaltIRDisabled               int = 9
	GestaltInitHeapZerosOutHeapsBit int = 1
	// GestaltIntel: If the [Gestalt] function returns `gestaltIntel`, the system is is an Intel-based Macintosh.
	GestaltIntel             int = 10
	GestaltJapanAdjustADBKbd int = 17
	// GestaltKeyboardType: # Discussion
	GestaltKeyboardType int = 'k'<<24 | 'b'<<16 | 'd'<<8 | ' ' // 'kbd '
	// GestaltLaunchCanReturn: # Discussion
	GestaltLaunchCanReturn int = 1
	// GestaltLaunchControl: If this bit is set, the Process Manager is available.
	GestaltLaunchControl int = 3
	// GestaltLaunchFullFileSpec: # Discussion
	GestaltLaunchFullFileSpec int = 2
	// GestaltLineLevelInput: # Discussion
	GestaltLineLevelInput int = 9
	// GestaltLocationErr: Specifies the gestalt function ptr was not in the system heap.
	GestaltLocationErr    int = -5553
	GestaltMBLegacy       int = 0
	GestaltMBMultipleBays int = 2
	GestaltMBSingleBay    int = 1
	// GestaltMMUType: The selector you pass to the [Gestalt] function to determine the  type of MMU currently installed.
	GestaltMMUType                 int = 'm'<<24 | 'm'<<16 | 'u'<<8 | ' ' // 'mmu '
	GestaltMac512KE                int = 3
	GestaltMacAndPad               int = 2
	GestaltMacCentris610           int = 52
	GestaltMacCentris650           int = 30
	GestaltMacCentris660AV         int = 60
	GestaltMacClassic              int = 17
	GestaltMacColorClassic         int = 49
	GestaltMacII                   int = 6
	GestaltMacIIci                 int = 11
	GestaltMacIIcx                 int = 8
	GestaltMacIIfx                 int = 13
	GestaltMacIIsi                 int = 18
	GestaltMacIIvi                 int = 44
	GestaltMacIIvm                 int = 45
	GestaltMacIIvx                 int = 48
	GestaltMacIIx                  int = 7
	GestaltMacKbd                  int = 1
	GestaltMacLC                   int = 19
	GestaltMacLC475                int = 89
	GestaltMacLC520                int = 56
	GestaltMacLC575                int = 92
	GestaltMacLC580                int = 99
	GestaltMacLCII                 int = 37
	GestaltMacLCIII                int = 27
	GestaltMacOSCompatibilityValue int = 1206
	GestaltMacOSXQD                int = 768
	GestaltMacOSXQDText            int = 512
	GestaltMacPlus                 int = 4
	GestaltMacPlusKbd              int = 3
	GestaltMacQuadra605            int = 94
	GestaltMacQuadra610            int = 53
	GestaltMacQuadra630            int = 98
	GestaltMacQuadra650            int = 36
	GestaltMacQuadra660AV          int = 60
	GestaltMacQuadra700            int = 22
	GestaltMacQuadra800            int = 35
	GestaltMacQuadra840AV          int = 78
	GestaltMacQuadra900            int = 20
	GestaltMacQuadra950            int = 26
	GestaltMacSE                   int = 5
	GestaltMacSE030                int = 9
	GestaltMacTV                   int = 88
	GestaltMacXL                   int = 2
	GestaltMachineType             int = 'm'<<24 | 'a'<<16 | 'c'<<8 | 'h' // 'mach'
	GestaltMediaBay                int = 'm'<<24 | 'b'<<16 | 'e'<<8 | 'h' // 'mbeh'
	// GestaltMiscAttr: The selector you pass to the [Gestalt] function to determine information about miscellaneous pieces of the Operating System or hardware configuration.
	GestaltMiscAttr int = 'm'<<24 | 'i'<<16 | 's'<<8 | 'c' // 'misc'
	// GestaltMixedModeAttr: The Gestalt selector you pass to determine what version of Mixed Mode Manager is present.
	GestaltMixedModeAttr int = 'm'<<24 | 'i'<<16 | 'x'<<8 | 'd' // 'mixd'
	// GestaltMixedModeCFM68K: True if Mixed Mode supports CFM-68K calling conventions
	GestaltMixedModeCFM68K int = 1
	// GestaltMixedModeCFM68KHasState: True if CFM-68K Mixed Mode exports Save/RestoreMixedModeState
	GestaltMixedModeCFM68KHasState int = 3
	// GestaltMixedModeCFM68KHasTrap: True if CFM-68K Mixed Mode implements `_MixedModeDispatch` (versions 1.0.1 and prior did not)
	GestaltMixedModeCFM68KHasTrap int = 2
	// GestaltMixedModePowerPC: True if Mixed Mode supports PowerPC ABI calling conventions
	GestaltMixedModePowerPC int = 0
	// GestaltMultiChannels: # Discussion
	GestaltMultiChannels int = 11
	// GestaltMustUseFCBAccessors: If this bit is set in the `response` parameter, the File Manager no longer supports the low memory globals [FCBSPtr] and [FSFCBLen].
	GestaltMustUseFCBAccessors int = 13
	GestaltNativeCPUfamily     int = 'c'<<24 | 'p'<<16 | 'u'<<8 | 'f' // 'cpuf'
	// GestaltNativeCPUtype: # Discussion
	GestaltNativeCPUtype       int = 'c'<<24 | 'p'<<16 | 'u'<<8 | 't' // 'cput'
	GestaltNativeProcessMgrBit int = 19
	// GestaltNativeTimeMgr: If this bit is set, the native Time Manager is present.
	GestaltNativeTimeMgr                   int = 4
	GestaltNativeType1FontSupport          int = 5
	GestaltNewHandleReturnsZeroedMemoryBit int = 2
	GestaltNewPtrReturnsZeroedMemoryBit    int = 3
	GestaltNoFPU                           int = 0
	GestaltNoMMU                           int = 0
	GestaltNuBusPresent                    int = 1
	GestaltOCETB                           int = 258
	GestaltOCEToolboxVersion               int = 'o'<<24 | 'c'<<16 | 'e'<<8 | 't' // 'ocet'
	// GestaltOFA2available: OFA2 is available.
	GestaltOFA2available int = 3
	// GestaltOSAttr: # Discussion
	GestaltOSAttr             int = 'o'<<24 | 's'<<16 | ' '<<8 | ' ' // 'os  '
	GestaltOSLCompliantFinder int = 3
	// GestaltOSLInSystem: A [Gestalt] attribute constant.
	GestaltOSLInSystem          int = 2
	GestaltOSXFBCCurrentVersion int = 256
	// GestaltOriginalATSUVersion: Indicates that version 1.0 of ATSUI is installed on the user’s system.
	GestaltOriginalATSUVersion int = 65536
	GestaltOriginalQD          int = 0
	// GestaltOriginalQDText: This is the original version of QuickDraw Text, used through Mac OS 8.1.
	GestaltOriginalQDText int = 0
	// GestaltOutlineFonts: If `true`, outline fonts are supported.
	GestaltOutlineFonts          int = 0
	GestaltPCCard                int = 'p'<<24 | 'c'<<16 | 'c'<<8 | 'd' // 'pccd'
	GestaltPCCardFamilyPresent   int = 1
	GestaltPCCardHasPowerControl int = 2
	GestaltPCCardSupportsCardBus int = 3
	// GestaltPMgrCPUIdle: If true the CPU is capable of going into a low–power-consumption state.
	GestaltPMgrCPUIdle int = 1
	// GestaltPMgrDispatchExists: If true, Dispatch is present.
	GestaltPMgrDispatchExists int = 4
	// GestaltPMgrExists: If true, the Power Manager is present.
	GestaltPMgrExists int = 0
	// GestaltPMgrSCC: If true, it is possible to stop the SCC clock, thus effectively turning off the serial ports.
	GestaltPMgrSCC int = 2
	// GestaltPMgrSound: If true, it is possible to turn off power to the sound circuits.
	GestaltPMgrSound                           int = 3
	GestaltPMgrSupportsAVPowerStateAtSleepWake int = 5
	// GestaltPPCDragLibPresent: If the bit specified by this mask is set, the Drag Manager PPC Drag Library functions are available.
	GestaltPPCDragLibPresent      int = 2
	GestaltPPCQuickTimeLibPresent int = 0
	GestaltPS2Keyboard            int = 27
	// GestaltParityAttr: # Discussion
	GestaltParityAttr    int = 'p'<<24 | 'r'<<16 | 't'<<8 | 'y' // 'prty'
	GestaltParityEnabled int = 1
	// GestaltPartialRsrcs: If `true`, partial resources exist.
	GestaltPartialRsrcs int = 0
	GestaltPerforma250  int = 49
	GestaltPerforma450  int = 27
	GestaltPerforma46x  int = 62
	GestaltPerforma47x  int = 89
	GestaltPerforma5300 int = 41
	GestaltPerforma550  int = 80
	GestaltPerforma580  int = 99
	GestaltPerforma600  int = 45
	GestaltPerforma6300 int = 42
	GestaltPerforma6360 int = 58
	GestaltPerforma6400 int = 58
	// GestaltPlayAndRecord: # Discussion
	GestaltPlayAndRecord          int = 6
	GestaltPortADisabled          int = 5
	GestaltPortBDisabled          int = 6
	GestaltPortable               int = 10
	GestaltPortableSlotPresent    int = 4
	GestaltPortableUSBANSIKbd     int = 37
	GestaltPortableUSBISOKbd      int = 38
	GestaltPortableUSBJISKbd      int = 39
	GestaltPowerBook100           int = 24
	GestaltPowerBook140           int = 25
	GestaltPowerBook1400          int = 310
	GestaltPowerBook145           int = 54
	GestaltPowerBook150           int = 115
	GestaltPowerBook160           int = 34
	GestaltPowerBook165           int = 84
	GestaltPowerBook165c          int = 50
	GestaltPowerBook170           int = 21
	GestaltPowerBook180           int = 33
	GestaltPowerBook180c          int = 71
	GestaltPowerBook190           int = 85
	GestaltPowerBook2400          int = 307
	GestaltPowerBook3400          int = 306
	GestaltPowerBook500PPCUpgrade int = 126
	GestaltPowerBook520           int = 72
	GestaltPowerBook520c          int = 72
	GestaltPowerBook5300          int = 128
	GestaltPowerBook540           int = 72
	GestaltPowerBook540c          int = 72
	GestaltPowerBookDuo210        int = 29
	GestaltPowerBookDuo230        int = 32
	GestaltPowerBookDuo2300       int = 124
	GestaltPowerBookDuo250        int = 38
	GestaltPowerBookDuo270c       int = 77
	GestaltPowerBookDuo280        int = 102
	GestaltPowerBookDuo280c       int = 103
	GestaltPowerBookG3            int = 313
	GestaltPowerBookG3Series      int = 312
	GestaltPowerBookG3Series2     int = 314
	GestaltPowerMac4400           int = 515
	GestaltPowerMac4400_160       int = 514
	GestaltPowerMac5200           int = 41
	GestaltPowerMac5260           int = 41
	GestaltPowerMac5400           int = 74
	GestaltPowerMac5500           int = 512
	GestaltPowerMac6100_60        int = 75
	GestaltPowerMac6100_66        int = 100
	GestaltPowerMac6200           int = 42
	GestaltPowerMac6400           int = 58
	GestaltPowerMac6500           int = 513
	GestaltPowerMac7100_66        int = 112
	GestaltPowerMac7100_80        int = 47
	GestaltPowerMac7200           int = 108
	GestaltPowerMac7300           int = 109
	GestaltPowerMac7500           int = 68
	GestaltPowerMac7600           int = 68
	GestaltPowerMac8100_100       int = 55
	GestaltPowerMac8100_110       int = 40
	GestaltPowerMac8100_120       int = 12
	GestaltPowerMac8100_80        int = 65
	GestaltPowerMac8500           int = 69
	GestaltPowerMac8600           int = 69
	GestaltPowerMac9500           int = 67
	GestaltPowerMac9600           int = 67
	GestaltPowerMacCentris610     int = 119
	GestaltPowerMacCentris650     int = 120
	GestaltPowerMacG3             int = 510
	GestaltPowerMacLC475          int = 104
	GestaltPowerMacLC575          int = 105
	GestaltPowerMacLC630          int = 106
	GestaltPowerMacNewWorld       int = 406
	GestaltPowerMacPerforma47x    int = 104
	GestaltPowerMacPerforma57x    int = 105
	GestaltPowerMacPerforma63x    int = 106
	GestaltPowerMacQuadra610      int = 121
	GestaltPowerMacQuadra630      int = 106
	GestaltPowerMacQuadra650      int = 122
	GestaltPowerMacQuadra700      int = 116
	GestaltPowerMacQuadra800      int = 123
	GestaltPowerMacQuadra900      int = 117
	GestaltPowerMacQuadra950      int = 118
	// GestaltPowerMgrAttr: The [Gestalt] selector you pass to determine which Power Manager capabilities are available.
	GestaltPowerMgrAttr int = 'p'<<24 | 'o'<<16 | 'w'<<8 | 'r' // 'powr'
	// GestaltPowerPCValue: If the [Gestalt] function returns `gestaltPowerPC`, the system is a PowerPC Macintosh.
	GestaltPowerPCValue int = 2
	// GestaltPowerPCAware: Old name for `gestaltMixedModePowerPC`
	GestaltPowerPCAware  int = 0
	GestaltProF16ANSIKbd int = 31
	GestaltProF16ISOKbd  int = 32
	GestaltProF16JISKbd  int = 33
	// GestaltProcessorType: The selector you pass to the [Gestalt] function to determine the type of microprocessor currently running.
	GestaltProcessorType     int = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'c' // 'proc'
	GestaltPrtblADBKbd       int = 6
	GestaltPrtblISOKbd       int = 7
	GestaltPwrBk99JISKbd     int = 201
	GestaltPwrBkEKDomKbd     int = 195
	GestaltPwrBkEKISOKbd     int = 196
	GestaltPwrBkEKJISKbd     int = 197
	GestaltPwrBkExtADBKbd    int = 24
	GestaltPwrBkExtISOKbd    int = 20
	GestaltPwrBkExtJISKbd    int = 21
	GestaltPwrBkSubDomKbd    int = 28
	GestaltPwrBkSubISOKbd    int = 29
	GestaltPwrBkSubJISKbd    int = 30
	GestaltPwrBookADBKbd     int = 12
	GestaltPwrBookISOADBKbd  int = 13
	GestaltQDHasLongRowBytes int = 5
	GestaltQDTextFeatures    int = 'q'<<24 | 'd'<<16 | 't'<<8 | 'f' // 'qdtf'
	GestaltQDTextVersion     int = 'q'<<24 | 'd'<<16 | 't'<<8 | 'x' // 'qdtx'
	GestaltQuadra605         int = 94
	GestaltQuadra610         int = 53
	GestaltQuadra630         int = 98
	GestaltQuadra650         int = 36
	GestaltQuadra660AV       int = 60
	GestaltQuadra700         int = 22
	GestaltQuadra800         int = 35
	GestaltQuadra840AV       int = 78
	GestaltQuadra900         int = 20
	GestaltQuadra950         int = 26
	GestaltQuickTimeFeatures int = 'q'<<24 | 't'<<16 | 'r'<<8 | 's' // 'qtrs'
	// GestaltQuickdrawFeatures: The selector you pass to the [Gestalt] function to determine the QuickDraw features.
	GestaltQuickdrawFeatures int = 'q'<<24 | 'd'<<16 | 'r'<<8 | 'w' // 'qdrw'
	// GestaltQuickdrawVersion: The Gestalt selector you pass to determine what version of QuickDraw is present.
	GestaltQuickdrawVersion             int = 'q'<<24 | 'd'<<16 | ' '<<8 | ' ' // 'qd  '
	GestaltRMFakeAppleMenuItemsRolledIn int = 1
	GestaltRMForceSysHeapRolledIn       int = 0
	GestaltRMSupportsFSCalls            int = 4
	GestaltRMTypeIndexOrderingReverse   int = 8
	// GestaltRealTempMemory: If `true`, temporary memory handles are real.
	GestaltRealTempMemory int = 5
	// GestaltResourceMgrAttr: The [Gestalt] selector you pass to determine which Resource Manager attributes are present.
	GestaltResourceMgrAttr          int = 'r'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'rsrc'
	GestaltResourceMgrBugFixesAttrs int = 'r'<<24 | 'm'<<16 | 'b'<<8 | 'g' // 'rmbg'
	// GestaltRevisedTimeMgr: If this bit is set, the revised Time Manager is present.
	GestaltRevisedTimeMgr           int = 2
	GestaltSCSI                     int = 's'<<24 | 'c'<<16 | 's'<<8 | 'i' // 'scsi'
	GestaltSCSIPollSIH              int = 3
	GestaltSCSISlotBoot             int = 2
	GestaltSE30SlotPresent          int = 3
	GestaltSESlotPresent            int = 2
	GestaltSFServer                 int = 256
	GestaltSafeOFAttr               int = 's'<<24 | 'a'<<16 | 'f'<<8 | 'e' // 'safe'
	GestaltSanityCheckResourceFiles int = 2
	// GestaltSbitFontSupport: sbit-only fonts are supported.
	GestaltSbitFontSupport int = 1
	// GestaltScriptingSupport: # Discussion
	GestaltScriptingSupport        int = 1
	GestaltScrollingThrottle       int = 0
	GestaltSerialArbitrationExists int = 0
	// GestaltSerialAttr: The selector you pass to the [Gestalt] function to determine the serial hardware attributes of the machine, such as whether or not the GPIa line is connected and can be used for external clocking.
	GestaltSerialAttr               int = 's'<<24 | 'e'<<16 | 'r'<<8 | ' ' // 'ser '
	GestaltSetDragImageUpdates      int = 5
	GestaltSheetsAreWindowModalBit  int = 7
	GestaltSheetsAreWindowModalMask int = 128
	// GestaltSlotAttr: The selector you pass to the [Gestalt] function to determine the Slot Manager attributes.
	GestaltSlotAttr      int = 's'<<24 | 'l'<<16 | 'o'<<8 | 't' // 'slot'
	GestaltSlotMgrExists int = 0
	// GestaltSndPlayDoubleBuffer: # Discussion
	GestaltSndPlayDoubleBuffer int = 10
	// GestaltSoundAttr: The Gestalt selector which you pass to the [Gestalt] function.
	GestaltSoundAttr int = 's'<<24 | 'n'<<16 | 'd'<<8 | ' ' // 'snd '
	// GestaltSoundIOMgrPresent: Set if the Sound Input Manager is available.
	GestaltSoundIOMgrPresent int = 3
	// GestaltSpecificMatchSupport: This bit is set if the Thread Manager supports the allocation of threads based on an exact match with the requested stack size.
	GestaltSpecificMatchSupport int = 1
	// GestaltSpeechRecognitionAttr: The selector which you pass to the [Gestalt] function to determine the Speech Recognition Manager attributes.
	GestaltSpeechRecognitionAttr int = 's'<<24 | 'r'<<16 | 't'<<8 | 'a' // 'srta'
	GestaltSquareMenuBar         int = 2
	// GestaltStandardTimeMgr: If this bit is set, the original Time Manager is present.
	GestaltStandardTimeMgr int = 1
	GestaltStdADBKbd       int = 5
	GestaltStdISOADBKbd    int = 8
	// GestaltStereoCapability: Set if the built-in sound hardware is able to produce stereo sounds.
	GestaltStereoCapability int = 0
	// GestaltStereoInput: # Discussion
	GestaltStereoInput int = 8
	// GestaltStereoMixing: Set if the built-in sound hardware mixes both left and right channels of stereo sound into a single audio signal for the internal speaker.
	GestaltStereoMixing                          int = 1
	GestaltSupportsApplicationURL                int = 4
	GestaltSupportsFSpResourceFileAlreadyOpenBit int = 3
	GestaltSupportsMirroring                     int = 4
	// GestaltSysArchitecture: The selector you pass to the [Gestalt] function to determine the native system architecture.
	GestaltSysArchitecture    int = 's'<<24 | 'y'<<16 | 's'<<8 | 'a' // 'sysa'
	GestaltSysDebuggerSupport int = 8
	GestaltSysZoneGrowable    int = 0
	// GestaltTE1: The version of TextEdit found in Mac IIci ROM.
	GestaltTE1 int = 1
	// GestaltTE2: The version of TextEdit shipped with 6.0.4 Script Systems on Mac IIci (Script bug fixes for Mac IIci).
	GestaltTE2 int = 2
	// GestaltTE3: The version of TextEdit shipped with 6.0.4 Script Systems (all but Mac IIci).
	GestaltTE3 int = 3
	// GestaltTE4: The version of TextEdit shipped in System 7.0.
	GestaltTE4 int = 4
	// GestaltTE5: [TextWidthHook] is available in TextEdit.
	GestaltTE5    int = 5
	GestaltTVAttr int = 't'<<24 | 'v'<<16 | ' '<<8 | ' ' // 'tv  '
	// GestaltTelephoneSpeechRecognition: If this bit is set, the Speech Recognition Manager supports telephone input.
	GestaltTelephoneSpeechRecognition int = 2
	// GestaltTempMemSupport: If `true`, there is temporary memory support.
	GestaltTempMemSupport int = 4
	// GestaltTempMemTracked: If `true`, temporary memory handles are tracked.
	GestaltTempMemTracked int = 6
	// GestaltTextEditVersion: The [Gestalt] selector you pass to determine what version of TextEdit is present.
	GestaltTextEditVersion   int = 't'<<24 | 'e'<<16 | ' '<<8 | ' ' // 'te  '
	GestaltThirdPartyANSIKbd int = 40
	GestaltThirdPartyISOKbd  int = 41
	GestaltThirdPartyJISKbd  int = 42
	GestaltThreadMgrAttr     int = 't'<<24 | 'h'<<16 | 'd'<<8 | 's' // 'thds'
	// GestaltThreadMgrPresent: This bit is set if the Thread Manager is present.
	GestaltThreadMgrPresent int = 0
	// GestaltThreadsLibraryPresent: This bit is set if the native version of the threads library has been loaded.
	GestaltThreadsLibraryPresent int = 2
	// GestaltTimeMgrVersion: The Gestalt selector you pass to determine what version of the Time Manager is present.
	GestaltTimeMgrVersion   int = 't'<<24 | 'm'<<16 | 'g'<<8 | 'r' // 'tmgr'
	GestaltUSBAndyANSIKbd   int = 204
	GestaltUSBAndyISOKbd    int = 205
	GestaltUSBAndyJISKbd    int = 206
	GestaltUSBCosmoANSIKbd  int = 198
	GestaltUSBCosmoISOKbd   int = 199
	GestaltUSBCosmoJISKbd   int = 200
	GestaltUSBProF16ANSIKbd int = 34
	GestaltUSBProF16ISOKbd  int = 35
	GestaltUSBProF16JISKbd  int = 36
	// GestaltUndefSelectorErr: Specifies an undefined selector was passed to the Gestalt Manager.
	GestaltUndefSelectorErr int = -5551
	// GestaltUnknownErr: Specifies an unknown error.
	GestaltUnknownErr           int = -5550
	GestaltUnknownThirdPartyKbd int = 3
	GestaltVMZerosPagesBit      int = 0
	// GestaltValueImplementedVers: The first version of the Gestalt Manager that implements this selector.
	GestaltValueImplementedVers int = 5
	// GestaltVersion: The selector you pass to the function Gestalt to determine the version of the Gestalt Manager.
	GestaltVersion                         int = 'v'<<24 | 'e'<<16 | 'r'<<8 | 's' // 'vers'
	GestaltWSIICanPrintWithoutPrGeneralBit int = 0
	// GestaltWSIISupport: WSII support is included.
	GestaltWSIISupport          int = 0
	GestaltWindowLiveResizeBit  int = 4
	GestaltWindowLiveResizeMask int = 16
	// GestaltWindowMgrAttr: # Discussion
	GestaltWindowMgrAttr int = 'w'<<24 | 'i'<<16 | 'n'<<8 | 'd' // 'wind'
	// GestaltWindowMgrPresent: If the bit specified by this mask is set, the Window Manager functionality for Appearance Manager 1.1 is available.
	GestaltWindowMgrPresent         int = 1
	GestaltWindowMgrPresentBit      int = 0
	GestaltWindowMgrPresentMask     int = 1
	GestaltWindowMinimizeToDockBit  int = 5
	GestaltWindowMinimizeToDockMask int = 32
	GestaltWorldScriptIIAttr        int = 'w'<<24 | 's'<<16 | 'a'<<8 | 't' // 'wsat'
	GestaltWorldScriptIIVersion     int = 'd'<<24 | 'o'<<16 | 'u'<<8 | 'b' // 'doub'
)

type GestaltAFP uint

const (
	GestaltAFPClient              GestaltAFP = 'a'<<24 | 'f'<<16 | 'p'<<8 | 's' // 'afps'
	GestaltAFPClient3_5           GestaltAFP = 1
	GestaltAFPClient3_6           GestaltAFP = 2
	GestaltAFPClient3_6_1         GestaltAFP = 3
	GestaltAFPClient3_6_2         GestaltAFP = 4
	GestaltAFPClient3_6_3         GestaltAFP = 5
	GestaltAFPClient3_7           GestaltAFP = 6
	GestaltAFPClient3_7_2         GestaltAFP = 7
	GestaltAFPClient3_8           GestaltAFP = 8
	GestaltAFPClient3_8_1         GestaltAFP = 9
	GestaltAFPClient3_8_3         GestaltAFP = 10
	GestaltAFPClient3_8_4         GestaltAFP = 11
	GestaltAFPClientAttributeMask GestaltAFP = 0
	GestaltAFPClientCfgRsrc       GestaltAFP = 16
	GestaltAFPClientMultiReq      GestaltAFP = 31
	GestaltAFPClientSupportsIP    GestaltAFP = 29
	GestaltAFPClientVMUI          GestaltAFP = 30
	GestaltAFPClientVersionMask   GestaltAFP = 65535
)

func (e GestaltAFP) String() string {
	switch e {
	case GestaltAFPClient:
		return "GestaltAFPClient"
	case GestaltAFPClient3_5:
		return "GestaltAFPClient3_5"
	case GestaltAFPClient3_6:
		return "GestaltAFPClient3_6"
	case GestaltAFPClient3_6_1:
		return "GestaltAFPClient3_6_1"
	case GestaltAFPClient3_6_2:
		return "GestaltAFPClient3_6_2"
	case GestaltAFPClient3_6_3:
		return "GestaltAFPClient3_6_3"
	case GestaltAFPClient3_7:
		return "GestaltAFPClient3_7"
	case GestaltAFPClient3_7_2:
		return "GestaltAFPClient3_7_2"
	case GestaltAFPClient3_8:
		return "GestaltAFPClient3_8"
	case GestaltAFPClient3_8_1:
		return "GestaltAFPClient3_8_1"
	case GestaltAFPClient3_8_3:
		return "GestaltAFPClient3_8_3"
	case GestaltAFPClient3_8_4:
		return "GestaltAFPClient3_8_4"
	case GestaltAFPClientAttributeMask:
		return "GestaltAFPClientAttributeMask"
	case GestaltAFPClientCfgRsrc:
		return "GestaltAFPClientCfgRsrc"
	case GestaltAFPClientMultiReq:
		return "GestaltAFPClientMultiReq"
	case GestaltAFPClientSupportsIP:
		return "GestaltAFPClientSupportsIP"
	case GestaltAFPClientVMUI:
		return "GestaltAFPClientVMUI"
	case GestaltAFPClientVersionMask:
		return "GestaltAFPClientVersionMask"
	default:
		return fmt.Sprintf("GestaltAFP(%d)", e)
	}
}

type GestaltALM uint

const (
	GestaltALMAttr               GestaltALM = 't'<<24 | 'r'<<16 | 'i'<<8 | 'p' // 'trip'
	GestaltALMHasCFMSupport      GestaltALM = 2
	GestaltALMHasRescanNotifiers GestaltALM = 3
	GestaltALMHasSFGroup         GestaltALM = 1
	GestaltALMPresent            GestaltALM = 0
	GestaltALMVers               GestaltALM = 'w'<<24 | 'a'<<16 | 'l'<<8 | 'k' // 'walk'
)

func (e GestaltALM) String() string {
	switch e {
	case GestaltALMAttr:
		return "GestaltALMAttr"
	case GestaltALMHasCFMSupport:
		return "GestaltALMHasCFMSupport"
	case GestaltALMHasRescanNotifiers:
		return "GestaltALMHasRescanNotifiers"
	case GestaltALMHasSFGroup:
		return "GestaltALMHasSFGroup"
	case GestaltALMPresent:
		return "GestaltALMPresent"
	case GestaltALMVers:
		return "GestaltALMVers"
	default:
		return fmt.Sprintf("GestaltALM(%d)", e)
	}
}

type GestaltALMHasSF uint

const (
	GestaltALMHasSFLocation GestaltALMHasSF = 1
)

func (e GestaltALMHasSF) String() string {
	switch e {
	case GestaltALMHasSFLocation:
		return "GestaltALMHasSFLocation"
	default:
		return fmt.Sprintf("GestaltALMHasSF(%d)", e)
	}
}

type GestaltATA uint

const (
	GestaltATAAttr    GestaltATA = 'a'<<24 | 't'<<16 | 'a'<<8 | ' ' // 'ata '
	GestaltATAPresent GestaltATA = 0
)

func (e GestaltATA) String() string {
	switch e {
	case GestaltATAAttr:
		return "GestaltATAAttr"
	case GestaltATAPresent:
		return "GestaltATAPresent"
	default:
		return fmt.Sprintf("GestaltATA(%d)", e)
	}
}

type GestaltATSU uint

const (
	// GestaltATSUAscentDescentControlsFeature: If the bit specified by this mask is set, ascent and descent controls (`kATSUDescentTag` and `kATSUAscentTag`) are available.
	GestaltATSUAscentDescentControlsFeature GestaltATSU = 16
	// GestaltATSUBatchBreakLinesFeature: If the bit specified by this mask is set, the [ATSUBatchBreakLines] function is available.
	GestaltATSUBatchBreakLinesFeature GestaltATSU = 16
	// GestaltATSUBiDiCursorPositionFeature: If the bit specified by this mask is set, support for bidirectional cursor positioning is available.
	GestaltATSUBiDiCursorPositionFeature GestaltATSU = 32
	// GestaltATSUByCharacterClusterFeature: If the bit specified by this mask is set, ATSUI cursor movement types are available.
	GestaltATSUByCharacterClusterFeature GestaltATSU = 16
	// GestaltATSUDecimalTabFeature: If the bit specified by this mask is set, your application can set a decimal tab character.
	GestaltATSUDecimalTabFeature GestaltATSU = 32
	// GestaltATSUDirectAccess: If the bit specified by this mask is set, ATSU direct-access functions are available.
	GestaltATSUDirectAccess GestaltATSU = 16
	// GestaltATSUDropShadowStyleFeature: If the bit specified by this mask is set, drop shadow features are available.
	GestaltATSUDropShadowStyleFeature GestaltATSU = 32
	// GestaltATSUFallbacksFeature: If the bit specified by this mask is set, the functions [ATSUSetFontFallbacks] and [ATSUGetFontFallbacks] are available.
	GestaltATSUFallbacksFeature GestaltATSU = 1
	// GestaltATSUFallbacksObjFeatures: If the bit specified by this mask is set, [ATSUFontFallbacks] objects are available.
	GestaltATSUFallbacksObjFeatures GestaltATSU = 8
	// GestaltATSUFeatures: # Discussion
	GestaltATSUFeatures GestaltATSU = 'u'<<24 | 'i'<<16 | 's'<<8 | 'f' // 'uisf'
	// GestaltATSUGlyphBoundsFeature: If the bit specified by this mask is set, the function [ATSUGetGlyphBounds] is available.
	GestaltATSUGlyphBoundsFeature GestaltATSU = 1
	// GestaltATSUHighlightColorControlFeature: If the bit specified by this mask is set, your application can control highlight color.
	GestaltATSUHighlightColorControlFeature GestaltATSU = 32
	// GestaltATSUHighlightInactiveTextFeature: If the bit specified by this mask is set, the highlight inactive text feature is available.
	GestaltATSUHighlightInactiveTextFeature GestaltATSU = 16
	// GestaltATSUIgnoreLeadingFeature: If the bit specified by this mask is set, the line layout option (`kATSIgnoreFontLeadingTag`) to ignore the font leading value is available.
	GestaltATSUIgnoreLeadingFeature GestaltATSU = 8
	// GestaltATSULayoutCacheClearFeature: If the bit specified by this mask is set, the function [ATSUClearLayoutCache] is available.
	GestaltATSULayoutCacheClearFeature GestaltATSU = 1
	// GestaltATSULayoutCreateAndCopyFeature: If the bit specified by this mask is set, the function [ATSUCreateAndCopyTextLayout] is available.
	GestaltATSULayoutCreateAndCopyFeature GestaltATSU = 1
	// GestaltATSULineControlFeature: If the bit specified by this mask is set, the functions [ATSUCopyLineControls], [ATSUSetLineControls], [ATSUGetLineControl], [ATSUGetAllLineControls], and [ATSUClearLineControls] are available.
	GestaltATSULineControlFeature GestaltATSU = 1
	// GestaltATSULowLevelOrigFeatures: If the bit specified by this mask is set, the low-level features introduced in ATSUI version 2.0 are available.
	GestaltATSULowLevelOrigFeatures GestaltATSU = 4
	// GestaltATSUMemoryFeature: If the bit specified by this mask is set, the functions [ATSUCreateMemorySetting], [ATSUSetCurrentMemorySetting], [ATSUGetCurrentMemorySetting], and [ATSUDisposeMemorySetting] are available.
	GestaltATSUMemoryFeature GestaltATSU = 1
	// GestaltATSUNearestCharLineBreakFeature: If the bit specified by this mask is set, the nearest character line break feature is available.
	GestaltATSUNearestCharLineBreakFeature GestaltATSU = 32
	// GestaltATSUPositionToCursorFeature: If the bit specified by this mask is set, the position-to-cursor feature is available.
	GestaltATSUPositionToCursorFeature GestaltATSU = 16
	// GestaltATSUStrikeThroughStyleFeature: If the bit specified by this mask is set, strike through styles are available.
	GestaltATSUStrikeThroughStyleFeature GestaltATSU = 32
	// GestaltATSUTabSupportFeature: If the bit specified by this mask is set, support for tabs is available.
	GestaltATSUTabSupportFeature GestaltATSU = 16
	// GestaltATSUTextLocatorUsageFeature: If the bit specified by this mask is set, the text-break locator attribute is available for both style and text layout objects.
	GestaltATSUTextLocatorUsageFeature GestaltATSU = 2
	// GestaltATSUTrackingFeature: If the bit specified by this mask constant is set, the functions [ATSUCountFontTracking] and [ATSUGetIndFontTracking] are available.
	GestaltATSUTrackingFeature GestaltATSU = 1
	// GestaltATSUUnderlineOptionsStyleFeature: If the bit specified by this mask is set, underline options are available.
	GestaltATSUUnderlineOptionsStyleFeature GestaltATSU = 32
)

func (e GestaltATSU) String() string {
	switch e {
	case GestaltATSUAscentDescentControlsFeature:
		return "GestaltATSUAscentDescentControlsFeature"
	case GestaltATSUBiDiCursorPositionFeature:
		return "GestaltATSUBiDiCursorPositionFeature"
	case GestaltATSUFallbacksFeature:
		return "GestaltATSUFallbacksFeature"
	case GestaltATSUFallbacksObjFeatures:
		return "GestaltATSUFallbacksObjFeatures"
	case GestaltATSUFeatures:
		return "GestaltATSUFeatures"
	case GestaltATSULowLevelOrigFeatures:
		return "GestaltATSULowLevelOrigFeatures"
	case GestaltATSUTextLocatorUsageFeature:
		return "GestaltATSUTextLocatorUsageFeature"
	default:
		return fmt.Sprintf("GestaltATSU(%d)", e)
	}
}

type GestaltATalk uint

const (
	// GestaltATalkVersion: # Discussion
	GestaltATalkVersion GestaltATalk = 'a'<<24 | 't'<<16 | 'k'<<8 | 'v' // 'atkv'
)

func (e GestaltATalk) String() string {
	switch e {
	case GestaltATalkVersion:
		return "GestaltATalkVersion"
	default:
		return fmt.Sprintf("GestaltATalk(%d)", e)
	}
}

type GestaltAUX uint

const (
	// GestaltAUXVersion: The version of A/UX if it is currently executing.
	GestaltAUXVersion GestaltAUX = 'a'<<24 | '/'<<16 | 'u'<<8 | 'x' // 'a/ux'
)

func (e GestaltAUX) String() string {
	switch e {
	case GestaltAUXVersion:
		return "GestaltAUXVersion"
	default:
		return fmt.Sprintf("GestaltAUX(%d)", e)
	}
}

type GestaltAVLTree uint

const (
	GestaltAVLTreeAttr                       GestaltAVLTree = 't'<<24 | 'r'<<16 | 'e'<<8 | 'e' // 'tree'
	GestaltAVLTreePresentBit                 GestaltAVLTree = 0
	GestaltAVLTreeSupportsHandleBasedTreeBit GestaltAVLTree = 1
	GestaltAVLTreeSupportsTreeLockingBit     GestaltAVLTree = 2
)

func (e GestaltAVLTree) String() string {
	switch e {
	case GestaltAVLTreeAttr:
		return "GestaltAVLTreeAttr"
	case GestaltAVLTreePresentBit:
		return "GestaltAVLTreePresentBit"
	case GestaltAVLTreeSupportsHandleBasedTreeBit:
		return "GestaltAVLTreeSupportsHandleBasedTreeBit"
	case GestaltAVLTreeSupportsTreeLockingBit:
		return "GestaltAVLTreeSupportsTreeLockingBit"
	default:
		return fmt.Sprintf("GestaltAVLTree(%d)", e)
	}
}

type GestaltAliasMgr uint

const (
	// GestaltAliasMgrAttr: The selector you pass to the [Gestalt] function to determine the Alias Manager attributes.
	GestaltAliasMgrAttr                             GestaltAliasMgr = 'a'<<24 | 'l'<<16 | 'i'<<8 | 's' // 'alis'
	GestaltAliasMgrFollowsAliasesWhenResolving      GestaltAliasMgr = 4
	GestaltAliasMgrPrefersPath                      GestaltAliasMgr = 7
	GestaltAliasMgrPresent                          GestaltAliasMgr = 0
	GestaltAliasMgrRequiresAccessors                GestaltAliasMgr = 8
	GestaltAliasMgrResolveAliasFileWithMountOptions GestaltAliasMgr = 3
	GestaltAliasMgrSupportsAOCEKeychain             GestaltAliasMgr = 2
	GestaltAliasMgrSupportsExtendedCalls            GestaltAliasMgr = 5
	GestaltAliasMgrSupportsFSCalls                  GestaltAliasMgr = 6
	GestaltAliasMgrSupportsRemoteAppletalk          GestaltAliasMgr = 1
)

func (e GestaltAliasMgr) String() string {
	switch e {
	case GestaltAliasMgrAttr:
		return "GestaltAliasMgrAttr"
	case GestaltAliasMgrFollowsAliasesWhenResolving:
		return "GestaltAliasMgrFollowsAliasesWhenResolving"
	case GestaltAliasMgrPrefersPath:
		return "GestaltAliasMgrPrefersPath"
	case GestaltAliasMgrPresent:
		return "GestaltAliasMgrPresent"
	case GestaltAliasMgrRequiresAccessors:
		return "GestaltAliasMgrRequiresAccessors"
	case GestaltAliasMgrResolveAliasFileWithMountOptions:
		return "GestaltAliasMgrResolveAliasFileWithMountOptions"
	case GestaltAliasMgrSupportsAOCEKeychain:
		return "GestaltAliasMgrSupportsAOCEKeychain"
	case GestaltAliasMgrSupportsExtendedCalls:
		return "GestaltAliasMgrSupportsExtendedCalls"
	case GestaltAliasMgrSupportsFSCalls:
		return "GestaltAliasMgrSupportsFSCalls"
	case GestaltAliasMgrSupportsRemoteAppletalk:
		return "GestaltAliasMgrSupportsRemoteAppletalk"
	default:
		return fmt.Sprintf("GestaltAliasMgr(%d)", e)
	}
}

type GestaltAppearance uint

const (
	// GestaltAppearanceAttr: The [Gestalt] selector passed to determine whether the Appearance Manager is present.
	GestaltAppearanceAttr GestaltAppearance = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'r' // 'appr'
	// GestaltAppearanceCompatMode: # Discussion
	GestaltAppearanceCompatMode GestaltAppearance = 1
	// GestaltAppearanceExists: # Discussion
	GestaltAppearanceExists GestaltAppearance = 0
	// GestaltAppearanceVersion: # Discussion
	GestaltAppearanceVersion GestaltAppearance = 'a'<<24 | 'p'<<16 | 'v'<<8 | 'r' // 'apvr'
)

func (e GestaltAppearance) String() string {
	switch e {
	case GestaltAppearanceAttr:
		return "GestaltAppearanceAttr"
	case GestaltAppearanceCompatMode:
		return "GestaltAppearanceCompatMode"
	case GestaltAppearanceExists:
		return "GestaltAppearanceExists"
	case GestaltAppearanceVersion:
		return "GestaltAppearanceVersion"
	default:
		return fmt.Sprintf("GestaltAppearance(%d)", e)
	}
}

type GestaltAppleScript uint

const (
	// GestaltAppleScriptAttr: # Discussion
	GestaltAppleScriptAttr           GestaltAppleScript = 'a'<<24 | 's'<<16 | 'c'<<8 | 'r' // 'ascr'
	GestaltAppleScriptPowerPCSupport GestaltAppleScript = 1
	// GestaltAppleScriptPresent: A [Gestalt] attribute constant.
	GestaltAppleScriptPresent GestaltAppleScript = 0
	// GestaltAppleScriptVersion: # Discussion
	GestaltAppleScriptVersion GestaltAppleScript = 'a'<<24 | 's'<<16 | 'c'<<8 | 'v' // 'ascv'
)

func (e GestaltAppleScript) String() string {
	switch e {
	case GestaltAppleScriptAttr:
		return "GestaltAppleScriptAttr"
	case GestaltAppleScriptPowerPCSupport:
		return "GestaltAppleScriptPowerPCSupport"
	case GestaltAppleScriptPresent:
		return "GestaltAppleScriptPresent"
	case GestaltAppleScriptVersion:
		return "GestaltAppleScriptVersion"
	default:
		return fmt.Sprintf("GestaltAppleScript(%d)", e)
	}
}

type GestaltAppleTalk uint

const (
	// GestaltAppleTalkVersion: # Discussion
	GestaltAppleTalkVersion GestaltAppleTalk = 'a'<<24 | 't'<<16 | 'l'<<8 | 'k' // 'atlk'
)

func (e GestaltAppleTalk) String() string {
	switch e {
	case GestaltAppleTalkVersion:
		return "GestaltAppleTalkVersion"
	default:
		return fmt.Sprintf("GestaltAppleTalk(%d)", e)
	}
}

type GestaltBusClk uint

const (
	GestaltBusClkSpeed GestaltBusClk = 'b'<<24 | 'c'<<16 | 'l'<<8 | 'k' // 'bclk'
)

func (e GestaltBusClk) String() string {
	switch e {
	case GestaltBusClkSpeed:
		return "GestaltBusClkSpeed"
	default:
		return fmt.Sprintf("GestaltBusClk(%d)", e)
	}
}

type GestaltBusClkSpeedM uint

const (
	GestaltBusClkSpeedMHz GestaltBusClkSpeedM = 'b'<<24 | 'c'<<16 | 'l'<<8 | 'm' // 'bclm'
)

func (e GestaltBusClkSpeedM) String() string {
	switch e {
	case GestaltBusClkSpeedMHz:
		return "GestaltBusClkSpeedMHz"
	default:
		return fmt.Sprintf("GestaltBusClkSpeedM(%d)", e)
	}
}

type GestaltCF uint

const (
	GestaltCFM99Present     GestaltCF = 2
	GestaltCFM99PresentMask GestaltCF = 4
	GestaltCFMAttr          GestaltCF = 'c'<<24 | 'f'<<16 | 'r'<<8 | 'g' // 'cfrg'
	GestaltCFMPresent       GestaltCF = 0
	GestaltCFMPresentMask   GestaltCF = 1
)

func (e GestaltCF) String() string {
	switch e {
	case GestaltCFM99Present:
		return "GestaltCFM99Present"
	case GestaltCFM99PresentMask:
		return "GestaltCFM99PresentMask"
	case GestaltCFMAttr:
		return "GestaltCFMAttr"
	case GestaltCFMPresent:
		return "GestaltCFMPresent"
	case GestaltCFMPresentMask:
		return "GestaltCFMPresentMask"
	default:
		return fmt.Sprintf("GestaltCF(%d)", e)
	}
}

type GestaltCP uint

const (
	GestaltCPU486        GestaltCP = 'i'<<24 | '4'<<16 | '8'<<8 | '6' // 'i486'
	GestaltCPU750FX      GestaltCP = 288
	GestaltCPU970        GestaltCP = 313
	GestaltCPU970FX      GestaltCP = 316
	GestaltCPU970MP      GestaltCP = 324
	GestaltCPUApollo     GestaltCP = 273
	GestaltCPUG47447     GestaltCP = 274
	GestaltCPUPentium    GestaltCP = 'i'<<24 | '5'<<16 | '8'<<8 | '6' // 'i586'
	GestaltCPUPentium4   GestaltCP = 'i'<<24 | '5'<<16 | 'i'<<8 | 'v' // 'i5iv'
	GestaltCPUPentiumII  GestaltCP = 'i'<<24 | '5'<<16 | 'i'<<8 | 'i' // 'i5ii'
	GestaltCPUPentiumPro GestaltCP = 'i'<<24 | '5'<<16 | 'p'<<8 | 'r' // 'i5pr'
	GestaltCPUX86        GestaltCP = 'i'<<24 | 'x'<<16 | 'x'<<8 | 'x' // 'ixxx'
)

func (e GestaltCP) String() string {
	switch e {
	case GestaltCPU486:
		return "GestaltCPU486"
	case GestaltCPU750FX:
		return "GestaltCPU750FX"
	case GestaltCPU970:
		return "GestaltCPU970"
	case GestaltCPU970FX:
		return "GestaltCPU970FX"
	case GestaltCPU970MP:
		return "GestaltCPU970MP"
	case GestaltCPUApollo:
		return "GestaltCPUApollo"
	case GestaltCPUG47447:
		return "GestaltCPUG47447"
	case GestaltCPUPentium:
		return "GestaltCPUPentium"
	case GestaltCPUPentium4:
		return "GestaltCPUPentium4"
	case GestaltCPUPentiumII:
		return "GestaltCPUPentiumII"
	case GestaltCPUPentiumPro:
		return "GestaltCPUPentiumPro"
	case GestaltCPUX86:
		return "GestaltCPUX86"
	default:
		return fmt.Sprintf("GestaltCP(%d)", e)
	}
}

type GestaltCPUAR uint

const (
	GestaltCPUARM64     GestaltCPUAR = 'a'<<24 | 'x'<<16 | '6'<<8 | '4' // 'ax64'
	GestaltCPUARMFamily GestaltCPUAR = 'a'<<24 | 'r'<<16 | 'm'<<8 | ' ' // 'arm '
)

func (e GestaltCPUAR) String() string {
	switch e {
	case GestaltCPUARM64:
		return "GestaltCPUARM64"
	case GestaltCPUARMFamily:
		return "GestaltCPUARMFamily"
	default:
		return fmt.Sprintf("GestaltCPUAR(%d)", e)
	}
}

type GestaltCRM uint

const (
	GestaltCRMAttr          GestaltCRM = 'c'<<24 | 'r'<<16 | 'm'<<8 | ' ' // 'crm '
	GestaltCRMPersistentFix GestaltCRM = 1
	GestaltCRMPresent       GestaltCRM = 0
	GestaltCRMToolRsrcCalls GestaltCRM = 2
)

func (e GestaltCRM) String() string {
	switch e {
	case GestaltCRMAttr:
		return "GestaltCRMAttr"
	case GestaltCRMPersistentFix:
		return "GestaltCRMPersistentFix"
	case GestaltCRMPresent:
		return "GestaltCRMPresent"
	case GestaltCRMToolRsrcCalls:
		return "GestaltCRMToolRsrcCalls"
	default:
		return fmt.Sprintf("GestaltCRM(%d)", e)
	}
}

type GestaltCTB uint

const (
	// GestaltCTBVersion: The version number of the Communications Toolbox (in the low-order word of the return value).
	GestaltCTBVersion GestaltCTB = 'c'<<24 | 't'<<16 | 'b'<<8 | 'v' // 'ctbv'
)

func (e GestaltCTB) String() string {
	switch e {
	case GestaltCTBVersion:
		return "GestaltCTBVersion"
	default:
		return fmt.Sprintf("GestaltCTB(%d)", e)
	}
}

type GestaltCarbon uint

const (
	GestaltCarbonVersion GestaltCarbon = 'c'<<24 | 'b'<<16 | 'o'<<8 | 'n' // 'cbon'
)

func (e GestaltCarbon) String() string {
	switch e {
	case GestaltCarbonVersion:
		return "GestaltCarbonVersion"
	default:
		return fmt.Sprintf("GestaltCarbon(%d)", e)
	}
}

type GestaltCloseView uint

const (
	GestaltCloseViewAttr               GestaltCloseView = 'B'<<24 | 'S'<<16 | 'D'<<8 | 'a' // 'BSDa'
	GestaltCloseViewDisplayMgrFriendly GestaltCloseView = 1
	GestaltCloseViewEnabled            GestaltCloseView = 0
)

func (e GestaltCloseView) String() string {
	switch e {
	case GestaltCloseViewAttr:
		return "GestaltCloseViewAttr"
	case GestaltCloseViewDisplayMgrFriendly:
		return "GestaltCloseViewDisplayMgrFriendly"
	case GestaltCloseViewEnabled:
		return "GestaltCloseViewEnabled"
	default:
		return fmt.Sprintf("GestaltCloseView(%d)", e)
	}
}

type GestaltCollectionMgr uint

const (
	// GestaltCollectionMgrVersion: Collection Manager version.
	GestaltCollectionMgrVersion GestaltCollectionMgr = 'c'<<24 | 'l'<<16 | 't'<<8 | 'n' // 'cltn'
)

func (e GestaltCollectionMgr) String() string {
	switch e {
	case GestaltCollectionMgrVersion:
		return "GestaltCollectionMgrVersion"
	default:
		return fmt.Sprintf("GestaltCollectionMgr(%d)", e)
	}
}

type GestaltColor uint

const (
	// GestaltColorMatchingVersion: The selector for obtaining version information.
	GestaltColorMatchingVersion GestaltColor = 'c'<<24 | 'm'<<16 | 't'<<8 | 'c' // 'cmtc'
	GestaltColorPicker          GestaltColor = 'c'<<24 | 'p'<<16 | 'k'<<8 | 'r' // 'cpkr'
	GestaltColorPickerVersion   GestaltColor = 'c'<<24 | 'p'<<16 | 'k'<<8 | 'r' // 'cpkr'
	// GestaltColorSync10: A [Gestalt] response value of `gestaltColorSync10` indicates version 1.0 of the ColorSync Manager is present.
	GestaltColorSync10 GestaltColor = 256
	// GestaltColorSync104: A [Gestalt] response value of `gestaltColorSync104` indicates version 1.4 of the ColorSync Manager is present.
	GestaltColorSync104 GestaltColor = 260
	// GestaltColorSync105: A [Gestalt] response value of `gestaltColorSync105` indicates version 1.5 of the ColorSync Manager is present.
	GestaltColorSync105 GestaltColor = 261
	// GestaltColorSync11: A [Gestalt] response value of `gestaltColorSync11` indicates version 1.0.3 of the ColorSync Manager is present.
	GestaltColorSync11 GestaltColor = 272
	// GestaltColorSync20: A [Gestalt] response value of `gestaltColorSync20` indicates version 2.0 of the ColorSync Manager is present.
	GestaltColorSync20 GestaltColor = 512
	// GestaltColorSync21: A [Gestalt] response value of `gestaltColorSync21` indicates version 2.1 of the ColorSync Manager is present.
	GestaltColorSync21  GestaltColor = 528
	GestaltColorSync211 GestaltColor = 529
	GestaltColorSync212 GestaltColor = 530
	GestaltColorSync213 GestaltColor = 531
	// GestaltColorSync25: A [Gestalt] response value of `gestaltColorSync25` indicates version 2.5 of the ColorSync Manager is present.
	GestaltColorSync25  GestaltColor = 592
	GestaltColorSync26  GestaltColor = 608
	GestaltColorSync261 GestaltColor = 609
	GestaltColorSync30  GestaltColor = 768
)

func (e GestaltColor) String() string {
	switch e {
	case GestaltColorMatchingVersion:
		return "GestaltColorMatchingVersion"
	case GestaltColorPicker:
		return "GestaltColorPicker"
	case GestaltColorSync10:
		return "GestaltColorSync10"
	case GestaltColorSync104:
		return "GestaltColorSync104"
	case GestaltColorSync105:
		return "GestaltColorSync105"
	case GestaltColorSync11:
		return "GestaltColorSync11"
	case GestaltColorSync20:
		return "GestaltColorSync20"
	case GestaltColorSync21:
		return "GestaltColorSync21"
	case GestaltColorSync211:
		return "GestaltColorSync211"
	case GestaltColorSync212:
		return "GestaltColorSync212"
	case GestaltColorSync213:
		return "GestaltColorSync213"
	case GestaltColorSync25:
		return "GestaltColorSync25"
	case GestaltColorSync26:
		return "GestaltColorSync26"
	case GestaltColorSync261:
		return "GestaltColorSync261"
	case GestaltColorSync30:
		return "GestaltColorSync30"
	default:
		return fmt.Sprintf("GestaltColor(%d)", e)
	}
}

type GestaltComponent uint

const (
	// GestaltComponentMgr: The [Gestalt] selector you pass to determine what version of the Component Manager is present.
	GestaltComponentMgr      GestaltComponent = 'c'<<24 | 'p'<<16 | 'n'<<8 | 't' // 'cpnt'
	GestaltComponentPlatform GestaltComponent = 'c'<<24 | 'o'<<16 | 'p'<<8 | 'l' // 'copl'
)

func (e GestaltComponent) String() string {
	switch e {
	case GestaltComponentMgr:
		return "GestaltComponentMgr"
	case GestaltComponentPlatform:
		return "GestaltComponentPlatform"
	default:
		return fmt.Sprintf("GestaltComponent(%d)", e)
	}
}

type GestaltCompression uint

const (
	GestaltCompressionMgr GestaltCompression = 'i'<<24 | 'c'<<16 | 'm'<<8 | 'p' // 'icmp'
)

func (e GestaltCompression) String() string {
	switch e {
	case GestaltCompressionMgr:
		return "GestaltCompressionMgr"
	default:
		return fmt.Sprintf("GestaltCompression(%d)", e)
	}
}

type GestaltConnMgr uint

const (
	GestaltConnMgrAttr GestaltConnMgr = 'c'<<24 | 'o'<<16 | 'n'<<8 | 'n' // 'conn'
	// GestaltConnMgrCMSearchFix: The `gestaltConnMgrCMSearchFix` bit flag indicates that the fix is present that allows the [CMAddSearch] function to work over the `mAttn` channel.
	GestaltConnMgrCMSearchFix  GestaltConnMgr = 1
	GestaltConnMgrErrorString  GestaltConnMgr = 2
	GestaltConnMgrMultiAsyncIO GestaltConnMgr = 3
	GestaltConnMgrPresent      GestaltConnMgr = 0
)

func (e GestaltConnMgr) String() string {
	switch e {
	case GestaltConnMgrAttr:
		return "GestaltConnMgrAttr"
	case GestaltConnMgrCMSearchFix:
		return "GestaltConnMgrCMSearchFix"
	case GestaltConnMgrErrorString:
		return "GestaltConnMgrErrorString"
	case GestaltConnMgrMultiAsyncIO:
		return "GestaltConnMgrMultiAsyncIO"
	case GestaltConnMgrPresent:
		return "GestaltConnMgrPresent"
	default:
		return fmt.Sprintf("GestaltConnMgr(%d)", e)
	}
}

type GestaltControl uint

const (
	// GestaltControlMgrAttr: # Discussion
	GestaltControlMgrAttr GestaltControl = 'c'<<24 | 'n'<<16 | 't'<<8 | 'l' // 'cntl'
	// GestaltControlMgrPresent: If the bit specified by this mask is set, the Control Manager functionality for Appearance Manager 1.1 is available.
	GestaltControlMgrPresent     GestaltControl = 1
	GestaltControlMgrPresentBit  GestaltControl = 0
	GestaltControlMsgPresentMask GestaltControl = 1
)

func (e GestaltControl) String() string {
	switch e {
	case GestaltControlMgrAttr:
		return "GestaltControlMgrAttr"
	case GestaltControlMgrPresent:
		return "GestaltControlMgrPresent"
	case GestaltControlMgrPresentBit:
		return "GestaltControlMgrPresentBit"
	default:
		return fmt.Sprintf("GestaltControl(%d)", e)
	}
}

type GestaltControlMgr uint

const (
	GestaltControlMgrVersion GestaltControlMgr = 'c'<<24 | 'm'<<16 | 'v'<<8 | 'r' // 'cmvr'
)

func (e GestaltControlMgr) String() string {
	switch e {
	case GestaltControlMgrVersion:
		return "GestaltControlMgrVersion"
	default:
		return fmt.Sprintf("GestaltControlMgr(%d)", e)
	}
}

type GestaltControlStrip uint

const (
	GestaltControlStripAttr         GestaltControlStrip = 's'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'sdev'
	GestaltControlStripExists       GestaltControlStrip = 0
	GestaltControlStripUserFont     GestaltControlStrip = 2
	GestaltControlStripUserHotKey   GestaltControlStrip = 3
	GestaltControlStripVersion      GestaltControlStrip = 'c'<<24 | 's'<<16 | 'v'<<8 | 'r' // 'csvr'
	GestaltControlStripVersionFixed GestaltControlStrip = 1
)

func (e GestaltControlStrip) String() string {
	switch e {
	case GestaltControlStripAttr:
		return "GestaltControlStripAttr"
	case GestaltControlStripExists:
		return "GestaltControlStripExists"
	case GestaltControlStripUserFont:
		return "GestaltControlStripUserFont"
	case GestaltControlStripUserHotKey:
		return "GestaltControlStripUserHotKey"
	case GestaltControlStripVersion:
		return "GestaltControlStripVersion"
	case GestaltControlStripVersionFixed:
		return "GestaltControlStripVersionFixed"
	default:
		return fmt.Sprintf("GestaltControlStrip(%d)", e)
	}
}

type GestaltCountOfCP uint

const (
	GestaltCountOfCPUs GestaltCountOfCP = 'c'<<24 | 'p'<<16 | 'u'<<8 | 's' // 'cpus'
)

func (e GestaltCountOfCP) String() string {
	switch e {
	case GestaltCountOfCPUs:
		return "GestaltCountOfCPUs"
	default:
		return fmt.Sprintf("GestaltCountOfCP(%d)", e)
	}
}

type GestaltDBAccessMgr uint

const (
	GestaltDBAccessMgrAttr    GestaltDBAccessMgr = 'd'<<24 | 'b'<<16 | 'a'<<8 | 'c' // 'dbac'
	GestaltDBAccessMgrPresent GestaltDBAccessMgr = 0
)

func (e GestaltDBAccessMgr) String() string {
	switch e {
	case GestaltDBAccessMgrAttr:
		return "GestaltDBAccessMgrAttr"
	case GestaltDBAccessMgrPresent:
		return "GestaltDBAccessMgrPresent"
	default:
		return fmt.Sprintf("GestaltDBAccessMgr(%d)", e)
	}
}

type GestaltDITLExt uint

const (
	GestaltDITLExtAttr GestaltDITLExt = 'd'<<24 | 'i'<<16 | 't'<<8 | 'l' // 'ditl'
	// GestaltDITLExtPresent: If this flag bit is [TRUE], then the Dialog Manager extensions included in System 7 are available.
	GestaltDITLExtPresent      GestaltDITLExt = 0
	GestaltDITLExtSupportsIctb GestaltDITLExt = 1
)

func (e GestaltDITLExt) String() string {
	switch e {
	case GestaltDITLExtAttr:
		return "GestaltDITLExtAttr"
	case GestaltDITLExtPresent:
		return "GestaltDITLExtPresent"
	case GestaltDITLExtSupportsIctb:
		return "GestaltDITLExtSupportsIctb"
	default:
		return fmt.Sprintf("GestaltDITLExt(%d)", e)
	}
}

type GestaltDTP uint

const (
	GestaltDTPInfo GestaltDTP = 'd'<<24 | 't'<<16 | 'p'<<8 | 'x' // 'dtpx'
)

func (e GestaltDTP) String() string {
	switch e {
	case GestaltDTPInfo:
		return "GestaltDTPInfo"
	default:
		return fmt.Sprintf("GestaltDTP(%d)", e)
	}
}

type GestaltDTPFeatures uint

const (
	GestaltDTPFeaturesValue GestaltDTPFeatures = 'd'<<24 | 't'<<16 | 'p'<<8 | 'f' // 'dtpf'
	KDTPThirdPartySupported GestaltDTPFeatures = 4
)

func (e GestaltDTPFeatures) String() string {
	switch e {
	case GestaltDTPFeaturesValue:
		return "GestaltDTPFeaturesValue"
	case KDTPThirdPartySupported:
		return "KDTPThirdPartySupported"
	default:
		return fmt.Sprintf("GestaltDTPFeatures(%d)", e)
	}
}

type GestaltDesktopPictures uint

const (
	GestaltDesktopPicturesAttr      GestaltDesktopPictures = 'd'<<24 | 'k'<<16 | 'p'<<8 | 'x' // 'dkpx'
	GestaltDesktopPicturesDisplayed GestaltDesktopPictures = 1
	GestaltDesktopPicturesInstalled GestaltDesktopPictures = 0
)

func (e GestaltDesktopPictures) String() string {
	switch e {
	case GestaltDesktopPicturesAttr:
		return "GestaltDesktopPicturesAttr"
	case GestaltDesktopPicturesDisplayed:
		return "GestaltDesktopPicturesDisplayed"
	case GestaltDesktopPicturesInstalled:
		return "GestaltDesktopPicturesInstalled"
	default:
		return fmt.Sprintf("GestaltDesktopPictures(%d)", e)
	}
}

type GestaltDialog uint

const (
	// GestaltDialogMgrAttr: # Discussion
	GestaltDialogMgrAttr             GestaltDialog = 'd'<<24 | 'l'<<16 | 'o'<<8 | 'g' // 'dlog'
	GestaltDialogMgrHasAquaAlertBit  GestaltDialog = 2
	GestaltDialogMgrHasAquaAlertMask GestaltDialog = 4
	// GestaltDialogMgrPresent: If the bit specified by this mask is set, the Dialog Manager functionality for Appearance Manager 1.1 is available.
	GestaltDialogMgrPresent     GestaltDialog = 1
	GestaltDialogMgrPresentBit  GestaltDialog = 0
	GestaltDialogMgrPresentMask GestaltDialog = 1
	GestaltDialogMsgPresentMask GestaltDialog = 1
)

func (e GestaltDialog) String() string {
	switch e {
	case GestaltDialogMgrAttr:
		return "GestaltDialogMgrAttr"
	case GestaltDialogMgrHasAquaAlertBit:
		return "GestaltDialogMgrHasAquaAlertBit"
	case GestaltDialogMgrHasAquaAlertMask:
		return "GestaltDialogMgrHasAquaAlertMask"
	case GestaltDialogMgrPresent:
		return "GestaltDialogMgrPresent"
	case GestaltDialogMgrPresentBit:
		return "GestaltDialogMgrPresentBit"
	default:
		return fmt.Sprintf("GestaltDialog(%d)", e)
	}
}

type GestaltDictionaryMgr uint

const (
	GestaltDictionaryMgrAttr    GestaltDictionaryMgr = 'd'<<24 | 'i'<<16 | 'c'<<8 | 't' // 'dict'
	GestaltDictionaryMgrPresent GestaltDictionaryMgr = 0
)

func (e GestaltDictionaryMgr) String() string {
	switch e {
	case GestaltDictionaryMgrAttr:
		return "GestaltDictionaryMgrAttr"
	case GestaltDictionaryMgrPresent:
		return "GestaltDictionaryMgrPresent"
	default:
		return fmt.Sprintf("GestaltDictionaryMgr(%d)", e)
	}
}

type GestaltDigitalSignature uint

const (
	GestaltDigitalSignatureVersion GestaltDigitalSignature = 'd'<<24 | 's'<<16 | 'i'<<8 | 'g' // 'dsig'
)

func (e GestaltDigitalSignature) String() string {
	switch e {
	case GestaltDigitalSignatureVersion:
		return "GestaltDigitalSignatureVersion"
	default:
		return fmt.Sprintf("GestaltDigitalSignature(%d)", e)
	}
}

type GestaltDiskCache uint

const (
	// GestaltDiskCacheSize: A selector that you pass to the [Gestalt] function.
	GestaltDiskCacheSize GestaltDiskCache = 'd'<<24 | 'c'<<16 | 's'<<8 | 'z' // 'dcsz'
)

func (e GestaltDiskCache) String() string {
	switch e {
	case GestaltDiskCacheSize:
		return "GestaltDiskCacheSize"
	default:
		return fmt.Sprintf("GestaltDiskCache(%d)", e)
	}
}

type GestaltDisplayMgr uint

const (
	// GestaltDisplayMgrAttr: The [Gestalt] selector you pass to determine which Display Manager attributes are present.
	GestaltDisplayMgrAttr GestaltDisplayMgr = 'd'<<24 | 'p'<<16 | 'l'<<8 | 'y' // 'dply'
	// GestaltDisplayMgrCanConfirm: Not yet supported.
	GestaltDisplayMgrCanConfirm GestaltDisplayMgr = 4
	// GestaltDisplayMgrCanSwitchMirrored: If `true`, the Display Manager can switch modes on mirrored displays.
	GestaltDisplayMgrCanSwitchMirrored GestaltDisplayMgr = 2
	// GestaltDisplayMgrColorSyncAware: If `true`, Display Manager supports profiles for displays.
	GestaltDisplayMgrColorSyncAware    GestaltDisplayMgr = 5
	GestaltDisplayMgrGeneratesProfiles GestaltDisplayMgr = 6
	// GestaltDisplayMgrPresent: If `true`, the Display Manager is present.
	GestaltDisplayMgrPresent GestaltDisplayMgr = 0
	// GestaltDisplayMgrSetDepthNotifies: If `true`, and you have registered for notification and you will be notified of depth mode changes.
	GestaltDisplayMgrSetDepthNotifies GestaltDisplayMgr = 3
	GestaltDisplayMgrSleepNotifies    GestaltDisplayMgr = 7
	// GestaltDisplayMgrVers: The [Gestalt] selector you pass to determine what version of the Display Manager is present.
	GestaltDisplayMgrVers GestaltDisplayMgr = 'd'<<24 | 'p'<<16 | 'l'<<8 | 'v' // 'dplv'
)

func (e GestaltDisplayMgr) String() string {
	switch e {
	case GestaltDisplayMgrAttr:
		return "GestaltDisplayMgrAttr"
	case GestaltDisplayMgrCanConfirm:
		return "GestaltDisplayMgrCanConfirm"
	case GestaltDisplayMgrCanSwitchMirrored:
		return "GestaltDisplayMgrCanSwitchMirrored"
	case GestaltDisplayMgrColorSyncAware:
		return "GestaltDisplayMgrColorSyncAware"
	case GestaltDisplayMgrGeneratesProfiles:
		return "GestaltDisplayMgrGeneratesProfiles"
	case GestaltDisplayMgrPresent:
		return "GestaltDisplayMgrPresent"
	case GestaltDisplayMgrSetDepthNotifies:
		return "GestaltDisplayMgrSetDepthNotifies"
	case GestaltDisplayMgrSleepNotifies:
		return "GestaltDisplayMgrSleepNotifies"
	case GestaltDisplayMgrVers:
		return "GestaltDisplayMgrVers"
	default:
		return fmt.Sprintf("GestaltDisplayMgr(%d)", e)
	}
}

type GestaltDrawSprocket uint

const (
	GestaltDrawSprocketVersion GestaltDrawSprocket = 'd'<<24 | 's'<<16 | 'p'<<8 | 'v' // 'dspv'
)

func (e GestaltDrawSprocket) String() string {
	switch e {
	case GestaltDrawSprocketVersion:
		return "GestaltDrawSprocketVersion"
	default:
		return fmt.Sprintf("GestaltDrawSprocket(%d)", e)
	}
}

type GestaltEasyAccess uint

const (
	GestaltEasyAccessAttr   GestaltEasyAccess = 'e'<<24 | 'a'<<16 | 's'<<8 | 'y' // 'easy'
	GestaltEasyAccessLocked GestaltEasyAccess = 3
	GestaltEasyAccessOff    GestaltEasyAccess = 0
	GestaltEasyAccessOn     GestaltEasyAccess = 1
	GestaltEasyAccessSticky GestaltEasyAccess = 2
)

func (e GestaltEasyAccess) String() string {
	switch e {
	case GestaltEasyAccessAttr:
		return "GestaltEasyAccessAttr"
	case GestaltEasyAccessLocked:
		return "GestaltEasyAccessLocked"
	case GestaltEasyAccessOff:
		return "GestaltEasyAccessOff"
	case GestaltEasyAccessOn:
		return "GestaltEasyAccessOn"
	case GestaltEasyAccessSticky:
		return "GestaltEasyAccessSticky"
	default:
		return fmt.Sprintf("GestaltEasyAccess(%d)", e)
	}
}

type GestaltEditionMgr uint

const (
	GestaltEditionMgrAttr             GestaltEditionMgr = 'e'<<24 | 'd'<<16 | 't'<<8 | 'n' // 'edtn'
	GestaltEditionMgrPresent          GestaltEditionMgr = 0
	GestaltEditionMgrTranslationAware GestaltEditionMgr = 1
)

func (e GestaltEditionMgr) String() string {
	switch e {
	case GestaltEditionMgrAttr:
		return "GestaltEditionMgrAttr"
	case GestaltEditionMgrPresent:
		return "GestaltEditionMgrPresent"
	case GestaltEditionMgrTranslationAware:
		return "GestaltEditionMgrTranslationAware"
	default:
		return fmt.Sprintf("GestaltEditionMgr(%d)", e)
	}
}

type GestaltExtToolbox uint

const (
	// GestaltExtToolboxTable: The base address of the second half of the Toolbox trap table if the table is discontiguous.
	GestaltExtToolboxTable GestaltExtToolbox = 'x'<<24 | 't'<<16 | 't'<<8 | 't' // 'xttt'
)

func (e GestaltExtToolbox) String() string {
	switch e {
	case GestaltExtToolboxTable:
		return "GestaltExtToolboxTable"
	default:
		return fmt.Sprintf("GestaltExtToolbox(%d)", e)
	}
}

type GestaltExtensionTable uint

const (
	GestaltExtensionTableVersion GestaltExtensionTable = 'e'<<24 | 't'<<16 | 'b'<<8 | 'l' // 'etbl'
)

func (e GestaltExtensionTable) String() string {
	switch e {
	case GestaltExtensionTableVersion:
		return "GestaltExtensionTableVersion"
	default:
		return fmt.Sprintf("GestaltExtensionTable(%d)", e)
	}
}

type GestaltFB uint

const (
	GestaltFBCIndexingState    GestaltFB = 'f'<<24 | 'b'<<16 | 'c'<<8 | 'i' // 'fbci'
	GestaltFBCindexingCritical GestaltFB = 1
	GestaltFBCindexingSafe     GestaltFB = 0
)

func (e GestaltFB) String() string {
	switch e {
	case GestaltFBCIndexingState:
		return "GestaltFBCIndexingState"
	case GestaltFBCindexingCritical:
		return "GestaltFBCindexingCritical"
	case GestaltFBCindexingSafe:
		return "GestaltFBCindexingSafe"
	default:
		return fmt.Sprintf("GestaltFB(%d)", e)
	}
}

type GestaltFSM uint

const (
	// GestaltFSMVersion: Pass this selector to the [Gestalt] function to determine the version of the HFS External File Systems Manager (FSM).
	GestaltFSMVersion GestaltFSM = 'f'<<24 | 's'<<16 | 'm'<<8 | ' ' // 'fsm '
)

func (e GestaltFSM) String() string {
	switch e {
	case GestaltFSMVersion:
		return "GestaltFSMVersion"
	default:
		return fmt.Sprintf("GestaltFSM(%d)", e)
	}
}

type GestaltFSSupportsDirectI uint

const (
	GestaltFSSupportsDirectIO GestaltFSSupportsDirectI = 11
)

func (e GestaltFSSupportsDirectI) String() string {
	switch e {
	case GestaltFSSupportsDirectIO:
		return "GestaltFSSupportsDirectIO"
	default:
		return fmt.Sprintf("GestaltFSSupportsDirectI(%d)", e)
	}
}

type GestaltFXfrMgr uint

const (
	GestaltFXfrMgrAsync GestaltFXfrMgr = 3
	// GestaltFXfrMgrAttr: The selector you pass to the [Gestalt] function to determine the File Transfer Manager attributes.
	GestaltFXfrMgrAttr        GestaltFXfrMgr = 'f'<<24 | 'x'<<16 | 'f'<<8 | 'r' // 'fxfr'
	GestaltFXfrMgrErrorString GestaltFXfrMgr = 2
	GestaltFXfrMgrMultiFile   GestaltFXfrMgr = 1
	GestaltFXfrMgrPresent     GestaltFXfrMgr = 0
)

func (e GestaltFXfrMgr) String() string {
	switch e {
	case GestaltFXfrMgrAsync:
		return "GestaltFXfrMgrAsync"
	case GestaltFXfrMgrAttr:
		return "GestaltFXfrMgrAttr"
	case GestaltFXfrMgrErrorString:
		return "GestaltFXfrMgrErrorString"
	case GestaltFXfrMgrMultiFile:
		return "GestaltFXfrMgrMultiFile"
	case GestaltFXfrMgrPresent:
		return "GestaltFXfrMgrPresent"
	default:
		return fmt.Sprintf("GestaltFXfrMgr(%d)", e)
	}
}

type GestaltFileMapping uint

const (
	GestaltFileMappingAttr             GestaltFileMapping = 'f'<<24 | 'l'<<16 | 'm'<<8 | 'p' // 'flmp'
	GestaltFileMappingMultipleFilesFix GestaltFileMapping = 1
	GestaltFileMappingPresent          GestaltFileMapping = 0
)

func (e GestaltFileMapping) String() string {
	switch e {
	case GestaltFileMappingAttr:
		return "GestaltFileMappingAttr"
	case GestaltFileMappingMultipleFilesFix:
		return "GestaltFileMappingMultipleFilesFix"
	case GestaltFileMappingPresent:
		return "GestaltFileMappingPresent"
	default:
		return fmt.Sprintf("GestaltFileMapping(%d)", e)
	}
}

type GestaltFindFolderRedirection uint

const (
	GestaltFindFolderRedirectionAttr GestaltFindFolderRedirection = 'f'<<24 | 'o'<<16 | 'l'<<8 | 'e' // 'fole'
)

func (e GestaltFindFolderRedirection) String() string {
	switch e {
	case GestaltFindFolderRedirectionAttr:
		return "GestaltFindFolderRedirectionAttr"
	default:
		return fmt.Sprintf("GestaltFindFolderRedirection(%d)", e)
	}
}

type GestaltFirstSlot uint

const (
	// GestaltFirstSlotNumber: The first physical slot.
	GestaltFirstSlotNumber GestaltFirstSlot = 's'<<24 | 'l'<<16 | 't'<<8 | '1' // 'slt1'
)

func (e GestaltFirstSlot) String() string {
	switch e {
	case GestaltFirstSlotNumber:
		return "GestaltFirstSlotNumber"
	default:
		return fmt.Sprintf("GestaltFirstSlot(%d)", e)
	}
}

type GestaltFloppy uint

const (
	GestaltFloppyAttr            GestaltFloppy = 'f'<<24 | 'l'<<16 | 'p'<<8 | 'y' // 'flpy'
	GestaltFloppyIsMFMOnly       GestaltFloppy = 0
	GestaltFloppyIsManualEject   GestaltFloppy = 1
	GestaltFloppyUsesDiskInPlace GestaltFloppy = 2
)

func (e GestaltFloppy) String() string {
	switch e {
	case GestaltFloppyAttr:
		return "GestaltFloppyAttr"
	case GestaltFloppyIsMFMOnly:
		return "GestaltFloppyIsMFMOnly"
	case GestaltFloppyIsManualEject:
		return "GestaltFloppyIsManualEject"
	case GestaltFloppyUsesDiskInPlace:
		return "GestaltFloppyUsesDiskInPlace"
	default:
		return fmt.Sprintf("GestaltFloppy(%d)", e)
	}
}

type GestaltGX uint

const (
	GestaltGXVersion GestaltGX = 'q'<<24 | 'd'<<16 | 'g'<<8 | 'x' // 'qdgx'
)

func (e GestaltGX) String() string {
	switch e {
	case GestaltGXVersion:
		return "GestaltGXVersion"
	default:
		return fmt.Sprintf("GestaltGX(%d)", e)
	}
}

type GestaltGXPrintingMgr uint

const (
	GestaltGXPrintingMgrVersion GestaltGXPrintingMgr = 'p'<<24 | 'm'<<16 | 'g'<<8 | 'r' // 'pmgr'
)

func (e GestaltGXPrintingMgr) String() string {
	switch e {
	case GestaltGXPrintingMgrVersion:
		return "GestaltGXPrintingMgrVersion"
	default:
		return fmt.Sprintf("GestaltGXPrintingMgr(%d)", e)
	}
}

type GestaltGraphics uint

const (
	GestaltGraphicsAttr        GestaltGraphics = 'g'<<24 | 'f'<<16 | 'x'<<8 | 'a' // 'gfxa'
	GestaltGraphicsIsDebugging GestaltGraphics = 1
	GestaltGraphicsIsLoaded    GestaltGraphics = 2
	GestaltGraphicsIsPowerPC   GestaltGraphics = 4
)

func (e GestaltGraphics) String() string {
	switch e {
	case GestaltGraphicsAttr:
		return "GestaltGraphicsAttr"
	case GestaltGraphicsIsDebugging:
		return "GestaltGraphicsIsDebugging"
	case GestaltGraphicsIsLoaded:
		return "GestaltGraphicsIsLoaded"
	case GestaltGraphicsIsPowerPC:
		return "GestaltGraphicsIsPowerPC"
	default:
		return fmt.Sprintf("GestaltGraphics(%d)", e)
	}
}

type GestaltHardwareVendor uint

const (
	GestaltHardwareVendorApple GestaltHardwareVendor = 'A'<<24 | 'p'<<16 | 'p'<<8 | 'l' // 'Appl'
	GestaltHardwareVendorCode  GestaltHardwareVendor = 'h'<<24 | 'r'<<16 | 'a'<<8 | 'd' // 'hrad'
)

func (e GestaltHardwareVendor) String() string {
	switch e {
	case GestaltHardwareVendorApple:
		return "GestaltHardwareVendorApple"
	case GestaltHardwareVendorCode:
		return "GestaltHardwareVendorCode"
	default:
		return fmt.Sprintf("GestaltHardwareVendor(%d)", e)
	}
}

type GestaltHasSingleWindowMode uint

const (
	GestaltHasSingleWindowModeBit  GestaltHasSingleWindowMode = 8
	GestaltHasSingleWindowModeMask GestaltHasSingleWindowMode = 256
)

func (e GestaltHasSingleWindowMode) String() string {
	switch e {
	case GestaltHasSingleWindowModeBit:
		return "GestaltHasSingleWindowModeBit"
	case GestaltHasSingleWindowModeMask:
		return "GestaltHasSingleWindowModeMask"
	default:
		return fmt.Sprintf("GestaltHasSingleWindowMode(%d)", e)
	}
}

type GestaltIconUtilities uint

const (
	// GestaltIconUtilitiesAttr: # Discussion
	GestaltIconUtilitiesAttr GestaltIconUtilities = 'i'<<24 | 'c'<<16 | 'o'<<8 | 'n' // 'icon'
	// GestaltIconUtilitiesHas32BitIcons: True if 32-bit deep icons are supported.
	GestaltIconUtilitiesHas32BitIcons GestaltIconUtilities = 2
	// GestaltIconUtilitiesHas48PixelIcons: True if 48x48 icons are supported by IconUtilities.
	GestaltIconUtilitiesHas48PixelIcons GestaltIconUtilities = 1
	// GestaltIconUtilitiesHas8BitDeepMasks: True if 8-bit deep masks are supported.
	GestaltIconUtilitiesHas8BitDeepMasks GestaltIconUtilities = 3
	// GestaltIconUtilitiesHasIconServices: True if IconServices is present.
	GestaltIconUtilitiesHasIconServices GestaltIconUtilities = 4
	// GestaltIconUtilitiesPresent: True if icon utilities are present.
	GestaltIconUtilitiesPresent GestaltIconUtilities = 0
)

func (e GestaltIconUtilities) String() string {
	switch e {
	case GestaltIconUtilitiesAttr:
		return "GestaltIconUtilitiesAttr"
	case GestaltIconUtilitiesHas32BitIcons:
		return "GestaltIconUtilitiesHas32BitIcons"
	case GestaltIconUtilitiesHas48PixelIcons:
		return "GestaltIconUtilitiesHas48PixelIcons"
	case GestaltIconUtilitiesHas8BitDeepMasks:
		return "GestaltIconUtilitiesHas8BitDeepMasks"
	case GestaltIconUtilitiesHasIconServices:
		return "GestaltIconUtilitiesHasIconServices"
	case GestaltIconUtilitiesPresent:
		return "GestaltIconUtilitiesPresent"
	default:
		return fmt.Sprintf("GestaltIconUtilities(%d)", e)
	}
}

type GestaltInternal uint

const (
	GestaltInternalDisplay GestaltInternal = 'i'<<24 | 'd'<<16 | 's'<<8 | 'p' // 'idsp'
)

func (e GestaltInternal) String() string {
	switch e {
	case GestaltInternalDisplay:
		return "GestaltInternalDisplay"
	default:
		return fmt.Sprintf("GestaltInternal(%d)", e)
	}
}

type GestaltLogicalPage uint

const (
	// GestaltLogicalPageSize: The logical page size.
	GestaltLogicalPageSize GestaltLogicalPage = 'p'<<24 | 'g'<<16 | 's'<<8 | 'z' // 'pgsz'
)

func (e GestaltLogicalPage) String() string {
	switch e {
	case GestaltLogicalPageSize:
		return "GestaltLogicalPageSize"
	default:
		return fmt.Sprintf("GestaltLogicalPage(%d)", e)
	}
}

type GestaltLogicalRAM uint

const (
	// GestaltLogicalRAMSize: # Discussion
	GestaltLogicalRAMSize GestaltLogicalRAM = 'l'<<24 | 'r'<<16 | 'a'<<8 | 'm' // 'lram'
)

func (e GestaltLogicalRAM) String() string {
	switch e {
	case GestaltLogicalRAMSize:
		return "GestaltLogicalRAMSize"
	default:
		return fmt.Sprintf("GestaltLogicalRAM(%d)", e)
	}
}

type GestaltLowMemory uint

const (
	// GestaltLowMemorySize: The size (in bytes) of the low-memory area.
	GestaltLowMemorySize GestaltLowMemory = 'l'<<24 | 'm'<<16 | 'e'<<8 | 'm' // 'lmem'
)

func (e GestaltLowMemory) String() string {
	switch e {
	case GestaltLowMemorySize:
		return "GestaltLowMemorySize"
	default:
		return fmt.Sprintf("GestaltLowMemory(%d)", e)
	}
}

type GestaltMP uint

const (
	// GestaltMPCallableAPIsAttr: The Gestalt selector passed to determine the availability of preemptive system software functions.
	GestaltMPCallableAPIsAttr GestaltMP = 'm'<<24 | 'p'<<16 | 's'<<8 | 'c' // 'mpsc'
	// GestaltMPDeviceManager: If this bit is set, you can call preemptively safe Device Manager function.
	GestaltMPDeviceManager GestaltMP = 1
	// GestaltMPFileManager: If this bit is set, you can call preemptively safe File Manager functions.
	GestaltMPFileManager GestaltMP = 0
	GestaltMPTrapCalls   GestaltMP = 2
)

func (e GestaltMP) String() string {
	switch e {
	case GestaltMPCallableAPIsAttr:
		return "GestaltMPCallableAPIsAttr"
	case GestaltMPDeviceManager:
		return "GestaltMPDeviceManager"
	case GestaltMPFileManager:
		return "GestaltMPFileManager"
	case GestaltMPTrapCalls:
		return "GestaltMPTrapCalls"
	default:
		return fmt.Sprintf("GestaltMP(%d)", e)
	}
}

type GestaltMacOSCompatibility uint

const (
	GestaltMacOSCompatibilityBoxAttr      GestaltMacOSCompatibility = 'b'<<24 | 'b'<<16 | 'o'<<8 | 'x' // 'bbox'
	GestaltMacOSCompatibilityBoxHasSerial GestaltMacOSCompatibility = 1
	GestaltMacOSCompatibilityBoxPresent   GestaltMacOSCompatibility = 0
	GestaltMacOSCompatibilityBoxless      GestaltMacOSCompatibility = 2
)

func (e GestaltMacOSCompatibility) String() string {
	switch e {
	case GestaltMacOSCompatibilityBoxAttr:
		return "GestaltMacOSCompatibilityBoxAttr"
	case GestaltMacOSCompatibilityBoxHasSerial:
		return "GestaltMacOSCompatibilityBoxHasSerial"
	case GestaltMacOSCompatibilityBoxPresent:
		return "GestaltMacOSCompatibilityBoxPresent"
	case GestaltMacOSCompatibilityBoxless:
		return "GestaltMacOSCompatibilityBoxless"
	default:
		return fmt.Sprintf("GestaltMacOSCompatibility(%d)", e)
	}
}

type GestaltMachine uint

const (
	// GestaltMachineIcon: # Discussion
	GestaltMachineIcon GestaltMachine = 'm'<<24 | 'i'<<16 | 'c'<<8 | 'n' // 'micn'
)

func (e GestaltMachine) String() string {
	switch e {
	case GestaltMachineIcon:
		return "GestaltMachineIcon"
	default:
		return fmt.Sprintf("GestaltMachine(%d)", e)
	}
}

type GestaltMemoryMap uint

const (
	GestaltMemoryMapAttr   GestaltMemoryMap = 'm'<<24 | 'm'<<16 | 'a'<<8 | 'p' // 'mmap'
	GestaltMemoryMapSparse GestaltMemoryMap = 0
)

func (e GestaltMemoryMap) String() string {
	switch e {
	case GestaltMemoryMapAttr:
		return "GestaltMemoryMapAttr"
	case GestaltMemoryMapSparse:
		return "GestaltMemoryMapSparse"
	default:
		return fmt.Sprintf("GestaltMemoryMap(%d)", e)
	}
}

type GestaltMenuMgr uint

const (
	GestaltMenuMgrAquaLayoutBit  GestaltMenuMgr = 1
	GestaltMenuMgrAquaLayoutMask GestaltMenuMgr = 2
	// GestaltMenuMgrAttr: # Discussion
	GestaltMenuMgrAttr                           GestaltMenuMgr = 'm'<<24 | 'e'<<16 | 'n'<<8 | 'u' // 'menu'
	GestaltMenuMgrCGImageMenuTitleBit            GestaltMenuMgr = 6
	GestaltMenuMgrCGImageMenuTitleMask           GestaltMenuMgr = 64
	GestaltMenuMgrMoreThanFiveMenusDeepBit       GestaltMenuMgr = 5
	GestaltMenuMgrMoreThanFiveMenusDeepMask      GestaltMenuMgr = 32
	GestaltMenuMgrMultipleItemsWithCommandIDBit  GestaltMenuMgr = 2
	GestaltMenuMgrMultipleItemsWithCommandIDMask GestaltMenuMgr = 4
	// GestaltMenuMgrPresent: If the bit specified by this mask is set, the Menu Manager functionality for Appearance Manager 1.1 is available.
	GestaltMenuMgrPresent                      GestaltMenuMgr = 1
	GestaltMenuMgrPresentBit                   GestaltMenuMgr = 0
	GestaltMenuMgrPresentMask                  GestaltMenuMgr = 1
	GestaltMenuMgrRetainsIconRefBit            GestaltMenuMgr = 3
	GestaltMenuMgrRetainsIconRefMask           GestaltMenuMgr = 8
	GestaltMenuMgrSendsMenuBoundsToDefProcBit  GestaltMenuMgr = 4
	GestaltMenuMgrSendsMenuBoundsToDefProcMask GestaltMenuMgr = 16
)

func (e GestaltMenuMgr) String() string {
	switch e {
	case GestaltMenuMgrAquaLayoutBit:
		return "GestaltMenuMgrAquaLayoutBit"
	case GestaltMenuMgrAquaLayoutMask:
		return "GestaltMenuMgrAquaLayoutMask"
	case GestaltMenuMgrAttr:
		return "GestaltMenuMgrAttr"
	case GestaltMenuMgrCGImageMenuTitleBit:
		return "GestaltMenuMgrCGImageMenuTitleBit"
	case GestaltMenuMgrCGImageMenuTitleMask:
		return "GestaltMenuMgrCGImageMenuTitleMask"
	case GestaltMenuMgrMoreThanFiveMenusDeepBit:
		return "GestaltMenuMgrMoreThanFiveMenusDeepBit"
	case GestaltMenuMgrMoreThanFiveMenusDeepMask:
		return "GestaltMenuMgrMoreThanFiveMenusDeepMask"
	case GestaltMenuMgrMultipleItemsWithCommandIDMask:
		return "GestaltMenuMgrMultipleItemsWithCommandIDMask"
	case GestaltMenuMgrPresentBit:
		return "GestaltMenuMgrPresentBit"
	case GestaltMenuMgrRetainsIconRefBit:
		return "GestaltMenuMgrRetainsIconRefBit"
	case GestaltMenuMgrRetainsIconRefMask:
		return "GestaltMenuMgrRetainsIconRefMask"
	case GestaltMenuMgrSendsMenuBoundsToDefProcMask:
		return "GestaltMenuMgrSendsMenuBoundsToDefProcMask"
	default:
		return fmt.Sprintf("GestaltMenuMgr(%d)", e)
	}
}

type GestaltMessageMgr uint

const (
	GestaltMessageMgrVersion GestaltMessageMgr = 'm'<<24 | 'e'<<16 | 's'<<8 | 's' // 'mess'
)

func (e GestaltMessageMgr) String() string {
	switch e {
	case GestaltMessageMgrVersion:
		return "GestaltMessageMgrVersion"
	default:
		return fmt.Sprintf("GestaltMessageMgr(%d)", e)
	}
}

type GestaltMixedMode uint

const (
	// GestaltMixedModeVersion: The selector you pass to the [Gestalt] function to determine the  version of Mixed Mode Manager.
	GestaltMixedModeVersion GestaltMixedMode = 'm'<<24 | 'i'<<16 | 'x'<<8 | 'd' // 'mixd'
)

func (e GestaltMixedMode) String() string {
	switch e {
	case GestaltMixedModeVersion:
		return "GestaltMixedModeVersion"
	default:
		return fmt.Sprintf("GestaltMixedMode(%d)", e)
	}
}

type GestaltMultipleUsers uint

const (
	GestaltMultipleUsersState GestaltMultipleUsers = 'm'<<24 | 'f'<<16 | 'd'<<8 | 'r' // 'mfdr'
)

func (e GestaltMultipleUsers) String() string {
	switch e {
	case GestaltMultipleUsersState:
		return "GestaltMultipleUsersState"
	default:
		return fmt.Sprintf("GestaltMultipleUsers(%d)", e)
	}
}

type GestaltNameRegistry uint

const (
	GestaltNameRegistryVersion GestaltNameRegistry = 'n'<<24 | 'r'<<16 | 'e'<<8 | 'g' // 'nreg'
)

func (e GestaltNameRegistry) String() string {
	switch e {
	case GestaltNameRegistryVersion:
		return "GestaltNameRegistryVersion"
	default:
		return fmt.Sprintf("GestaltNameRegistry(%d)", e)
	}
}

type GestaltNotification uint

const (
	// GestaltNotificationMgrAttr: .The Gestalt selector which you pass to the [Gestalt] function to determine Notification Manager attributes.
	GestaltNotificationMgrAttr GestaltNotification = 'n'<<24 | 'm'<<16 | 'g'<<8 | 'r' // 'nmgr'
	// GestaltNotificationPresent: True if the Notification Manager exists.
	GestaltNotificationPresent GestaltNotification = 0
)

func (e GestaltNotification) String() string {
	switch e {
	case GestaltNotificationMgrAttr:
		return "GestaltNotificationMgrAttr"
	case GestaltNotificationPresent:
		return "GestaltNotificationPresent"
	default:
		return fmt.Sprintf("GestaltNotification(%d)", e)
	}
}

type GestaltNuBus uint

const (
	// GestaltNuBusConnectors: A bitmap that describes the NuBus slot connector locations.
	GestaltNuBusConnectors GestaltNuBus = 's'<<24 | 'l'<<16 | 't'<<8 | 'c' // 'sltc'
)

func (e GestaltNuBus) String() string {
	switch e {
	case GestaltNuBusConnectors:
		return "GestaltNuBusConnectors"
	default:
		return fmt.Sprintf("GestaltNuBus(%d)", e)
	}
}

type GestaltNuBusSlot uint

const (
	GestaltNuBusSlotCount GestaltNuBusSlot = 'n'<<24 | 'u'<<16 | 'b'<<8 | 's' // 'nubs'
)

func (e GestaltNuBusSlot) String() string {
	switch e {
	case GestaltNuBusSlotCount:
		return "GestaltNuBusSlotCount"
	default:
		return fmt.Sprintf("GestaltNuBusSlot(%d)", e)
	}
}

type GestaltOCE uint

const (
	GestaltOCESFServerAvailable     GestaltOCE = 4
	GestaltOCETBAvailable           GestaltOCE = 2
	GestaltOCETBNativeGlueAvailable GestaltOCE = 16
	GestaltOCETBPresent             GestaltOCE = 1
	GestaltOCEToolboxAttr           GestaltOCE = 'o'<<24 | 'c'<<16 | 'e'<<8 | 'u' // 'oceu'
)

func (e GestaltOCE) String() string {
	switch e {
	case GestaltOCESFServerAvailable:
		return "GestaltOCESFServerAvailable"
	case GestaltOCETBAvailable:
		return "GestaltOCETBAvailable"
	case GestaltOCETBNativeGlueAvailable:
		return "GestaltOCETBNativeGlueAvailable"
	case GestaltOCETBPresent:
		return "GestaltOCETBPresent"
	case GestaltOCEToolboxAttr:
		return "GestaltOCEToolboxAttr"
	default:
		return fmt.Sprintf("GestaltOCE(%d)", e)
	}
}

type GestaltOS uint

const (
	// GestaltOSTable: The selector you pass to the [Gestalt] function to determine the  base address of the operating system trap dispatch table.
	GestaltOSTable GestaltOS = 'o'<<24 | 's'<<16 | 't'<<8 | 't' // 'ostt'
)

func (e GestaltOS) String() string {
	switch e {
	case GestaltOSTable:
		return "GestaltOSTable"
	default:
		return fmt.Sprintf("GestaltOS(%d)", e)
	}
}

type GestaltOpen uint

const (
	GestaltOpenTptValue                GestaltOpen = 'o'<<24 | 't'<<16 | 'a'<<8 | 'n' // 'otan'
	GestaltOpenTptAppleTalkLoadedBit   GestaltOpen = 3
	GestaltOpenTptAppleTalkLoadedMask  GestaltOpen = 8
	GestaltOpenTptAppleTalkPresentBit  GestaltOpen = 2
	GestaltOpenTptAppleTalkPresentMask GestaltOpen = 4
	GestaltOpenTptIPXSPXLoadedBit      GestaltOpen = 7
	GestaltOpenTptIPXSPXLoadedMask     GestaltOpen = 128
	GestaltOpenTptIPXSPXPresentBit     GestaltOpen = 6
	GestaltOpenTptIPXSPXPresentMask    GestaltOpen = 64
	GestaltOpenTptLoadedBit            GestaltOpen = 1
	GestaltOpenTptLoadedMask           GestaltOpen = 2
	GestaltOpenTptPresentBit           GestaltOpen = 0
	GestaltOpenTptPresentMask          GestaltOpen = 1
	GestaltOpenTptTCPLoadedBit         GestaltOpen = 5
	GestaltOpenTptTCPLoadedMask        GestaltOpen = 32
	GestaltOpenTptTCPPresentBit        GestaltOpen = 4
	GestaltOpenTptTCPPresentMask       GestaltOpen = 16
)

func (e GestaltOpen) String() string {
	switch e {
	case GestaltOpenTptValue:
		return "GestaltOpenTptValue"
	case GestaltOpenTptAppleTalkLoadedBit:
		return "GestaltOpenTptAppleTalkLoadedBit"
	case GestaltOpenTptAppleTalkLoadedMask:
		return "GestaltOpenTptAppleTalkLoadedMask"
	case GestaltOpenTptAppleTalkPresentBit:
		return "GestaltOpenTptAppleTalkPresentBit"
	case GestaltOpenTptAppleTalkPresentMask:
		return "GestaltOpenTptAppleTalkPresentMask"
	case GestaltOpenTptIPXSPXLoadedBit:
		return "GestaltOpenTptIPXSPXLoadedBit"
	case GestaltOpenTptIPXSPXLoadedMask:
		return "GestaltOpenTptIPXSPXLoadedMask"
	case GestaltOpenTptIPXSPXPresentBit:
		return "GestaltOpenTptIPXSPXPresentBit"
	case GestaltOpenTptIPXSPXPresentMask:
		return "GestaltOpenTptIPXSPXPresentMask"
	case GestaltOpenTptLoadedBit:
		return "GestaltOpenTptLoadedBit"
	case GestaltOpenTptPresentBit:
		return "GestaltOpenTptPresentBit"
	case GestaltOpenTptTCPLoadedBit:
		return "GestaltOpenTptTCPLoadedBit"
	case GestaltOpenTptTCPLoadedMask:
		return "GestaltOpenTptTCPLoadedMask"
	case GestaltOpenTptTCPPresentMask:
		return "GestaltOpenTptTCPPresentMask"
	default:
		return fmt.Sprintf("GestaltOpen(%d)", e)
	}
}

type GestaltOpenFirmware uint

const (
	GestaltOpenFirmwareInfo GestaltOpenFirmware = 'o'<<24 | 'p'<<16 | 'f'<<8 | 'w' // 'opfw'
)

func (e GestaltOpenFirmware) String() string {
	switch e {
	case GestaltOpenFirmwareInfo:
		return "GestaltOpenFirmwareInfo"
	default:
		return fmt.Sprintf("GestaltOpenFirmware(%d)", e)
	}
}

type GestaltOpenTpt uint

const (
	GestaltOpenTptARAPPresent            GestaltOpenTpt = 6
	GestaltOpenTptPPPPresent             GestaltOpenTpt = 5
	GestaltOpenTptRemoteAccessValue      GestaltOpenTpt = 'o'<<24 | 't'<<16 | 'r'<<8 | 'a' // 'otra'
	GestaltOpenTptRemoteAccessClientOnly GestaltOpenTpt = 2
	GestaltOpenTptRemoteAccessLoaded     GestaltOpenTpt = 1
	GestaltOpenTptRemoteAccessMPServer   GestaltOpenTpt = 4
	GestaltOpenTptRemoteAccessPServer    GestaltOpenTpt = 3
	GestaltOpenTptRemoteAccessPresent    GestaltOpenTpt = 0
	GestaltOpenTptVersions               GestaltOpenTpt = 'o'<<24 | 't'<<16 | 'v'<<8 | 'r' // 'otvr'
)

func (e GestaltOpenTpt) String() string {
	switch e {
	case GestaltOpenTptARAPPresent:
		return "GestaltOpenTptARAPPresent"
	case GestaltOpenTptPPPPresent:
		return "GestaltOpenTptPPPPresent"
	case GestaltOpenTptRemoteAccessValue:
		return "GestaltOpenTptRemoteAccessValue"
	case GestaltOpenTptRemoteAccessClientOnly:
		return "GestaltOpenTptRemoteAccessClientOnly"
	case GestaltOpenTptRemoteAccessLoaded:
		return "GestaltOpenTptRemoteAccessLoaded"
	case GestaltOpenTptRemoteAccessMPServer:
		return "GestaltOpenTptRemoteAccessMPServer"
	case GestaltOpenTptRemoteAccessPServer:
		return "GestaltOpenTptRemoteAccessPServer"
	case GestaltOpenTptRemoteAccessPresent:
		return "GestaltOpenTptRemoteAccessPresent"
	case GestaltOpenTptVersions:
		return "GestaltOpenTptVersions"
	default:
		return fmt.Sprintf("GestaltOpenTpt(%d)", e)
	}
}

type GestaltOpenTptNetwork uint

const (
	GestaltOpenTptNetworkSetupValue               GestaltOpenTptNetwork = 'o'<<24 | 't'<<16 | 'c'<<8 | 'f' // 'otcf'
	GestaltOpenTptNetworkSetupLegacyExport        GestaltOpenTptNetwork = 1
	GestaltOpenTptNetworkSetupLegacyImport        GestaltOpenTptNetwork = 0
	GestaltOpenTptNetworkSetupSupportsMultihoming GestaltOpenTptNetwork = 2
)

func (e GestaltOpenTptNetwork) String() string {
	switch e {
	case GestaltOpenTptNetworkSetupValue:
		return "GestaltOpenTptNetworkSetupValue"
	case GestaltOpenTptNetworkSetupLegacyExport:
		return "GestaltOpenTptNetworkSetupLegacyExport"
	case GestaltOpenTptNetworkSetupLegacyImport:
		return "GestaltOpenTptNetworkSetupLegacyImport"
	case GestaltOpenTptNetworkSetupSupportsMultihoming:
		return "GestaltOpenTptNetworkSetupSupportsMultihoming"
	default:
		return fmt.Sprintf("GestaltOpenTptNetwork(%d)", e)
	}
}

type GestaltOpenTptNetworkSetup uint

const (
	GestaltOpenTptNetworkSetupVersion GestaltOpenTptNetworkSetup = 'o'<<24 | 't'<<16 | 'c'<<8 | 'v' // 'otcv'
)

func (e GestaltOpenTptNetworkSetup) String() string {
	switch e {
	case GestaltOpenTptNetworkSetupVersion:
		return "GestaltOpenTptNetworkSetupVersion"
	default:
		return fmt.Sprintf("GestaltOpenTptNetworkSetup(%d)", e)
	}
}

type GestaltOpenTptRemoteAccess uint

const (
	GestaltOpenTptRemoteAccessVersion GestaltOpenTptRemoteAccess = 'o'<<24 | 't'<<16 | 'r'<<8 | 'v' // 'otrv'
)

func (e GestaltOpenTptRemoteAccess) String() string {
	switch e {
	case GestaltOpenTptRemoteAccessVersion:
		return "GestaltOpenTptRemoteAccessVersion"
	default:
		return fmt.Sprintf("GestaltOpenTptRemoteAccess(%d)", e)
	}
}

type GestaltPCX uint

const (
	// GestaltPCXAttr: The selector you pass to the [Gestalt] function to determine the PC Exchange attributes.
	GestaltPCXAttr            GestaltPCX = 'p'<<24 | 'c'<<16 | 'x'<<8 | 'g' // 'pcxg'
	GestaltPCXHas8and16BitFAT GestaltPCX = 0
	GestaltPCXHasProDOS       GestaltPCX = 1
	GestaltPCXNewUI           GestaltPCX = 2
	GestaltPCXUseICMapping    GestaltPCX = 3
)

func (e GestaltPCX) String() string {
	switch e {
	case GestaltPCXAttr:
		return "GestaltPCXAttr"
	case GestaltPCXHas8and16BitFAT:
		return "GestaltPCXHas8and16BitFAT"
	case GestaltPCXHasProDOS:
		return "GestaltPCXHasProDOS"
	case GestaltPCXNewUI:
		return "GestaltPCXNewUI"
	case GestaltPCXUseICMapping:
		return "GestaltPCXUseICMapping"
	default:
		return fmt.Sprintf("GestaltPCX(%d)", e)
	}
}

type GestaltPPC uint

const (
	GestaltPPCSupportsIncoming          GestaltPPC = 1
	GestaltPPCSupportsIncomingAppleTalk GestaltPPC = 16
	GestaltPPCSupportsIncomingTCP_IP    GestaltPPC = 32
	GestaltPPCSupportsOutGoing          GestaltPPC = 2
	GestaltPPCSupportsOutgoingAppleTalk GestaltPPC = 256
	GestaltPPCSupportsOutgoingTCP_IP    GestaltPPC = 512
	GestaltPPCSupportsRealTime          GestaltPPC = 4096
	GestaltPPCSupportsTCP_IP            GestaltPPC = 4
	// GestaltPPCToolboxAttr: The selector you pass to the Gestalt function to determine the Program-to-Program Communication (PPC) Toolbox attributes.
	GestaltPPCToolboxAttr    GestaltPPC = 'p'<<24 | 'p'<<16 | 'c'<<8 | ' ' // 'ppc '
	GestaltPPCToolboxPresent GestaltPPC = 0
)

func (e GestaltPPC) String() string {
	switch e {
	case GestaltPPCSupportsIncoming:
		return "GestaltPPCSupportsIncoming"
	case GestaltPPCSupportsIncomingAppleTalk:
		return "GestaltPPCSupportsIncomingAppleTalk"
	case GestaltPPCSupportsIncomingTCP_IP:
		return "GestaltPPCSupportsIncomingTCP_IP"
	case GestaltPPCSupportsOutGoing:
		return "GestaltPPCSupportsOutGoing"
	case GestaltPPCSupportsOutgoingAppleTalk:
		return "GestaltPPCSupportsOutgoingAppleTalk"
	case GestaltPPCSupportsOutgoingTCP_IP:
		return "GestaltPPCSupportsOutgoingTCP_IP"
	case GestaltPPCSupportsRealTime:
		return "GestaltPPCSupportsRealTime"
	case GestaltPPCSupportsTCP_IP:
		return "GestaltPPCSupportsTCP_IP"
	case GestaltPPCToolboxAttr:
		return "GestaltPPCToolboxAttr"
	case GestaltPPCToolboxPresent:
		return "GestaltPPCToolboxPresent"
	default:
		return fmt.Sprintf("GestaltPPC(%d)", e)
	}
}

type GestaltPhysicalRAM uint

const (
	// GestaltPhysicalRAMSize: The selector you pass to the [Gestalt] function to determine the  number of bytes of physical RAM currently installed.
	GestaltPhysicalRAMSize GestaltPhysicalRAM = 'r'<<24 | 'a'<<16 | 'm'<<8 | ' ' // 'ram '
)

func (e GestaltPhysicalRAM) String() string {
	switch e {
	case GestaltPhysicalRAMSize:
		return "GestaltPhysicalRAMSize"
	default:
		return fmt.Sprintf("GestaltPhysicalRAM(%d)", e)
	}
}

type GestaltPhysicalRAMSizeIn uint

const (
	GestaltPhysicalRAMSizeInMegabytes GestaltPhysicalRAMSizeIn = 'r'<<24 | 'a'<<16 | 'm'<<8 | 'm' // 'ramm'
)

func (e GestaltPhysicalRAMSizeIn) String() string {
	switch e {
	case GestaltPhysicalRAMSizeInMegabytes:
		return "GestaltPhysicalRAMSizeInMegabytes"
	default:
		return fmt.Sprintf("GestaltPhysicalRAMSizeIn(%d)", e)
	}
}

type GestaltPopup uint

const (
	// GestaltPopupAttr: The selector you pass to the [Gestalt] function to determine the  attribute of the pop-up control definition.
	GestaltPopupAttr    GestaltPopup = 'p'<<24 | 'o'<<16 | 'p'<<8 | '!' // 'pop!'
	GestaltPopupPresent GestaltPopup = 0
)

func (e GestaltPopup) String() string {
	switch e {
	case GestaltPopupAttr:
		return "GestaltPopupAttr"
	case GestaltPopupPresent:
		return "GestaltPopupPresent"
	default:
		return fmt.Sprintf("GestaltPopup(%d)", e)
	}
}

type GestaltPortable2001 uint

const (
	GestaltPortable2001ANSIKbd GestaltPortable2001 = 202
	GestaltPortable2001ISOKbd  GestaltPortable2001 = 203
	GestaltPortable2001JISKbd  GestaltPortable2001 = 207
)

func (e GestaltPortable2001) String() string {
	switch e {
	case GestaltPortable2001ANSIKbd:
		return "GestaltPortable2001ANSIKbd"
	case GestaltPortable2001ISOKbd:
		return "GestaltPortable2001ISOKbd"
	case GestaltPortable2001JISKbd:
		return "GestaltPortable2001JISKbd"
	default:
		return fmt.Sprintf("GestaltPortable2001(%d)", e)
	}
}

type GestaltPowerMgr uint

const (
	GestaltPowerMgrVers GestaltPowerMgr = 'p'<<24 | 'w'<<16 | 'r'<<8 | 'v' // 'pwrv'
)

func (e GestaltPowerMgr) String() string {
	switch e {
	case GestaltPowerMgrVers:
		return "GestaltPowerMgrVers"
	default:
		return fmt.Sprintf("GestaltPowerMgr(%d)", e)
	}
}

type GestaltPowerPC uint

const (
	GestaltPowerPCASArchitecture            GestaltPowerPC = 8
	GestaltPowerPCHas64BitSupport           GestaltPowerPC = 6
	GestaltPowerPCHasDCBAInstruction        GestaltPowerPC = 3
	GestaltPowerPCHasDCBTStreams            GestaltPowerPC = 7
	GestaltPowerPCHasDataStreams            GestaltPowerPC = 5
	GestaltPowerPCHasGraphicsInstructions   GestaltPowerPC = 0
	GestaltPowerPCHasSTFIWXInstruction      GestaltPowerPC = 1
	GestaltPowerPCHasSquareRootInstructions GestaltPowerPC = 2
	GestaltPowerPCHasVectorInstructions     GestaltPowerPC = 4
	GestaltPowerPCIgnoresDCBST              GestaltPowerPC = 9
	GestaltPowerPCProcessorFeatures         GestaltPowerPC = 'p'<<24 | 'p'<<16 | 'c'<<8 | 'f' // 'ppcf'
)

func (e GestaltPowerPC) String() string {
	switch e {
	case GestaltPowerPCASArchitecture:
		return "GestaltPowerPCASArchitecture"
	case GestaltPowerPCHas64BitSupport:
		return "GestaltPowerPCHas64BitSupport"
	case GestaltPowerPCHasDCBAInstruction:
		return "GestaltPowerPCHasDCBAInstruction"
	case GestaltPowerPCHasDCBTStreams:
		return "GestaltPowerPCHasDCBTStreams"
	case GestaltPowerPCHasDataStreams:
		return "GestaltPowerPCHasDataStreams"
	case GestaltPowerPCHasGraphicsInstructions:
		return "GestaltPowerPCHasGraphicsInstructions"
	case GestaltPowerPCHasSTFIWXInstruction:
		return "GestaltPowerPCHasSTFIWXInstruction"
	case GestaltPowerPCHasSquareRootInstructions:
		return "GestaltPowerPCHasSquareRootInstructions"
	case GestaltPowerPCHasVectorInstructions:
		return "GestaltPowerPCHasVectorInstructions"
	case GestaltPowerPCIgnoresDCBST:
		return "GestaltPowerPCIgnoresDCBST"
	case GestaltPowerPCProcessorFeatures:
		return "GestaltPowerPCProcessorFeatures"
	default:
		return fmt.Sprintf("GestaltPowerPC(%d)", e)
	}
}

type GestaltProcClk uint

const (
	GestaltProcClkSpeed GestaltProcClk = 'p'<<24 | 'c'<<16 | 'l'<<8 | 'k' // 'pclk'
)

func (e GestaltProcClk) String() string {
	switch e {
	case GestaltProcClkSpeed:
		return "GestaltProcClkSpeed"
	default:
		return fmt.Sprintf("GestaltProcClk(%d)", e)
	}
}

type GestaltProcClkSpeedM uint

const (
	GestaltProcClkSpeedMHz GestaltProcClkSpeedM = 'm'<<24 | 'c'<<16 | 'l'<<8 | 'k' // 'mclk'
)

func (e GestaltProcClkSpeedM) String() string {
	switch e {
	case GestaltProcClkSpeedMHz:
		return "GestaltProcClkSpeedMHz"
	default:
		return fmt.Sprintf("GestaltProcClkSpeedM(%d)", e)
	}
}

type GestaltProcessorCacheLine uint

const (
	GestaltProcessorCacheLineSize GestaltProcessorCacheLine = 'c'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'csiz'
)

func (e GestaltProcessorCacheLine) String() string {
	switch e {
	case GestaltProcessorCacheLineSize:
		return "GestaltProcessorCacheLineSize"
	default:
		return fmt.Sprintf("GestaltProcessorCacheLine(%d)", e)
	}
}

type GestaltQD3 uint

const (
	GestaltQD3DValue   GestaltQD3 = 'q'<<24 | 'd'<<16 | '3'<<8 | 'd' // 'qd3d'
	GestaltQD3DPresent GestaltQD3 = 0
)

func (e GestaltQD3) String() string {
	switch e {
	case GestaltQD3DValue:
		return "GestaltQD3DValue"
	case GestaltQD3DPresent:
		return "GestaltQD3DPresent"
	default:
		return fmt.Sprintf("GestaltQD3(%d)", e)
	}
}

type GestaltQD3D uint

const (
	GestaltQD3DVersion       GestaltQD3D = 'q'<<24 | '3'<<16 | 'v'<<8 | ' ' // 'q3v '
	GestaltQD3DViewer        GestaltQD3D = 'q'<<24 | '3'<<16 | 'v'<<8 | 'c' // 'q3vc'
	GestaltQD3DViewerPresent GestaltQD3D = 0
)

func (e GestaltQD3D) String() string {
	switch e {
	case GestaltQD3DVersion:
		return "GestaltQD3DVersion"
	case GestaltQD3DViewer:
		return "GestaltQD3DViewer"
	case GestaltQD3DViewerPresent:
		return "GestaltQD3DViewerPresent"
	default:
		return fmt.Sprintf("GestaltQD3D(%d)", e)
	}
}

type GestaltQTVR uint

const (
	GestaltQTVRCubicPanosPresent    GestaltQTVR = 3
	GestaltQTVRCylinderPanosPresent GestaltQTVR = 2
	GestaltQTVRMgrAttr              GestaltQTVR = 'q'<<24 | 't'<<16 | 'v'<<8 | 'r' // 'qtvr'
	GestaltQTVRMgrPresent           GestaltQTVR = 0
	GestaltQTVRObjMoviesPresent     GestaltQTVR = 1
)

func (e GestaltQTVR) String() string {
	switch e {
	case GestaltQTVRCubicPanosPresent:
		return "GestaltQTVRCubicPanosPresent"
	case GestaltQTVRCylinderPanosPresent:
		return "GestaltQTVRCylinderPanosPresent"
	case GestaltQTVRMgrAttr:
		return "GestaltQTVRMgrAttr"
	case GestaltQTVRMgrPresent:
		return "GestaltQTVRMgrPresent"
	case GestaltQTVRObjMoviesPresent:
		return "GestaltQTVRObjMoviesPresent"
	default:
		return fmt.Sprintf("GestaltQTVR(%d)", e)
	}
}

type GestaltQTVRMgr uint

const (
	GestaltQTVRMgrVers GestaltQTVRMgr = 'q'<<24 | 't'<<16 | 'v'<<8 | 'v' // 'qtvv'
)

func (e GestaltQTVRMgr) String() string {
	switch e {
	case GestaltQTVRMgrVers:
		return "GestaltQTVRMgrVers"
	default:
		return fmt.Sprintf("GestaltQTVRMgr(%d)", e)
	}
}

type GestaltQuick uint

const (
	GestaltQuickTimeValue GestaltQuick = 'q'<<24 | 't'<<16 | 'i'<<8 | 'm' // 'qtim'
	// GestaltQuickTimeVersion: The selector you pass to the [Gestalt] function to determine the QuickTime version.
	GestaltQuickTimeVersion GestaltQuick = 'q'<<24 | 't'<<16 | 'i'<<8 | 'm' // 'qtim'
)

func (e GestaltQuick) String() string {
	switch e {
	case GestaltQuickTimeValue:
		return "GestaltQuickTimeValue"
	default:
		return fmt.Sprintf("GestaltQuick(%d)", e)
	}
}

type GestaltQuickTime uint

const (
	GestaltQuickTimeConferencingValue GestaltQuickTime = 'm'<<24 | 't'<<16 | 'l'<<8 | 'k' // 'mtlk'
)

func (e GestaltQuickTime) String() string {
	switch e {
	case GestaltQuickTimeConferencingValue:
		return "GestaltQuickTimeConferencingValue"
	default:
		return fmt.Sprintf("GestaltQuickTime(%d)", e)
	}
}

type GestaltQuickTimeConferencing uint

const (
	GestaltQuickTimeConferencingInfo GestaltQuickTimeConferencing = 'q'<<24 | 't'<<16 | 'c'<<8 | 'i' // 'qtci'
)

func (e GestaltQuickTimeConferencing) String() string {
	switch e {
	case GestaltQuickTimeConferencingInfo:
		return "GestaltQuickTimeConferencingInfo"
	default:
		return fmt.Sprintf("GestaltQuickTimeConferencing(%d)", e)
	}
}

type GestaltQuickTimeStreaming uint

const (
	GestaltQuickTimeStreamingFeatures GestaltQuickTimeStreaming = 'q'<<24 | 't'<<16 | 's'<<8 | 'f' // 'qtsf'
	GestaltQuickTimeStreamingVersion  GestaltQuickTimeStreaming = 'q'<<24 | 't'<<16 | 's'<<8 | 't' // 'qtst'
)

func (e GestaltQuickTimeStreaming) String() string {
	switch e {
	case GestaltQuickTimeStreamingFeatures:
		return "GestaltQuickTimeStreamingFeatures"
	case GestaltQuickTimeStreamingVersion:
		return "GestaltQuickTimeStreamingVersion"
	default:
		return fmt.Sprintf("GestaltQuickTimeStreaming(%d)", e)
	}
}

type GestaltQuickTimeThreadSafe uint

const (
	GestaltQuickTimeThreadSafeFeaturesAttr   GestaltQuickTimeThreadSafe = 'q'<<24 | 't'<<16 | 't'<<8 | 'h' // 'qtth'
	GestaltQuickTimeThreadSafeGraphicsExport GestaltQuickTimeThreadSafe = 5
	GestaltQuickTimeThreadSafeGraphicsImport GestaltQuickTimeThreadSafe = 4
	GestaltQuickTimeThreadSafeICM            GestaltQuickTimeThreadSafe = 0
	GestaltQuickTimeThreadSafeMovieExport    GestaltQuickTimeThreadSafe = 3
	GestaltQuickTimeThreadSafeMovieImport    GestaltQuickTimeThreadSafe = 2
	GestaltQuickTimeThreadSafeMoviePlayback  GestaltQuickTimeThreadSafe = 6
	GestaltQuickTimeThreadSafeMovieToolbox   GestaltQuickTimeThreadSafe = 1
)

func (e GestaltQuickTimeThreadSafe) String() string {
	switch e {
	case GestaltQuickTimeThreadSafeFeaturesAttr:
		return "GestaltQuickTimeThreadSafeFeaturesAttr"
	case GestaltQuickTimeThreadSafeGraphicsExport:
		return "GestaltQuickTimeThreadSafeGraphicsExport"
	case GestaltQuickTimeThreadSafeGraphicsImport:
		return "GestaltQuickTimeThreadSafeGraphicsImport"
	case GestaltQuickTimeThreadSafeICM:
		return "GestaltQuickTimeThreadSafeICM"
	case GestaltQuickTimeThreadSafeMovieExport:
		return "GestaltQuickTimeThreadSafeMovieExport"
	case GestaltQuickTimeThreadSafeMovieImport:
		return "GestaltQuickTimeThreadSafeMovieImport"
	case GestaltQuickTimeThreadSafeMoviePlayback:
		return "GestaltQuickTimeThreadSafeMoviePlayback"
	case GestaltQuickTimeThreadSafeMovieToolbox:
		return "GestaltQuickTimeThreadSafeMovieToolbox"
	default:
		return fmt.Sprintf("GestaltQuickTimeThreadSafe(%d)", e)
	}
}

type GestaltRBV uint

const (
	GestaltRBVAddr GestaltRBV = 'r'<<24 | 'b'<<16 | 'v'<<8 | ' ' // 'rbv '
)

func (e GestaltRBV) String() string {
	switch e {
	case GestaltRBVAddr:
		return "GestaltRBVAddr"
	default:
		return fmt.Sprintf("GestaltRBV(%d)", e)
	}
}

type GestaltROM uint

const (
	// GestaltROMSize: # Discussion
	GestaltROMSize GestaltROM = 'r'<<24 | 'o'<<16 | 'm'<<8 | ' ' // 'rom '
	// GestaltROMVersion: # Discussion
	GestaltROMVersion GestaltROM = 'r'<<24 | 'o'<<16 | 'm'<<8 | 'v' // 'romv'
)

func (e GestaltROM) String() string {
	switch e {
	case GestaltROMSize:
		return "GestaltROMSize"
	case GestaltROMVersion:
		return "GestaltROMVersion"
	default:
		return fmt.Sprintf("GestaltROM(%d)", e)
	}
}

type GestaltRealtimeMgr uint

const (
	// GestaltRealtimeMgrAttr: The selector you pass to the [Gestalt] function to determine the Realtime Manager attributes.
	GestaltRealtimeMgrAttr GestaltRealtimeMgr = 'r'<<24 | 't'<<16 | 'm'<<8 | 'r' // 'rtmr'
	// GestaltRealtimeMgrPresent: ( description forthcoming )
	GestaltRealtimeMgrPresent GestaltRealtimeMgr = 0
)

func (e GestaltRealtimeMgr) String() string {
	switch e {
	case GestaltRealtimeMgrAttr:
		return "GestaltRealtimeMgrAttr"
	case GestaltRealtimeMgrPresent:
		return "GestaltRealtimeMgrPresent"
	default:
		return fmt.Sprintf("GestaltRealtimeMgr(%d)", e)
	}
}

type GestaltSCCRead uint

const (
	GestaltSCCReadAddr GestaltSCCRead = 's'<<24 | 'c'<<16 | 'c'<<8 | 'r' // 'sccr'
)

func (e GestaltSCCRead) String() string {
	switch e {
	case GestaltSCCReadAddr:
		return "GestaltSCCReadAddr"
	default:
		return fmt.Sprintf("GestaltSCCRead(%d)", e)
	}
}

type GestaltSCCWrite uint

const (
	GestaltSCCWriteAddr GestaltSCCWrite = 's'<<24 | 'c'<<16 | 'c'<<8 | 'w' // 'sccw'
)

func (e GestaltSCCWrite) String() string {
	switch e {
	case GestaltSCCWriteAddr:
		return "GestaltSCCWriteAddr"
	default:
		return fmt.Sprintf("GestaltSCCWrite(%d)", e)
	}
}

type GestaltSDPFind uint

const (
	GestaltSDPFindVersion GestaltSDPFind = 'd'<<24 | 'f'<<16 | 'n'<<8 | 'd' // 'dfnd'
)

func (e GestaltSDPFind) String() string {
	switch e {
	case GestaltSDPFindVersion:
		return "GestaltSDPFindVersion"
	default:
		return fmt.Sprintf("GestaltSDPFind(%d)", e)
	}
}

type GestaltSDPPrompt uint

const (
	GestaltSDPPromptVersion GestaltSDPPrompt = 'p'<<24 | 'r'<<16 | 'p'<<8 | 'v' // 'prpv'
)

func (e GestaltSDPPrompt) String() string {
	switch e {
	case GestaltSDPPromptVersion:
		return "GestaltSDPPromptVersion"
	default:
		return fmt.Sprintf("GestaltSDPPrompt(%d)", e)
	}
}

type GestaltSDPStandardDirectory uint

const (
	GestaltSDPStandardDirectoryVersion GestaltSDPStandardDirectory = 's'<<24 | 'd'<<16 | 'v'<<8 | 'r' // 'sdvr'
)

func (e GestaltSDPStandardDirectory) String() string {
	switch e {
	case GestaltSDPStandardDirectoryVersion:
		return "GestaltSDPStandardDirectoryVersion"
	default:
		return fmt.Sprintf("GestaltSDPStandardDirectory(%d)", e)
	}
}

type GestaltSMPMailer uint

const (
	GestaltSMPMailerVersion GestaltSMPMailer = 'm'<<24 | 'a'<<16 | 'l'<<8 | 'r' // 'malr'
)

func (e GestaltSMPMailer) String() string {
	switch e {
	case GestaltSMPMailerVersion:
		return "GestaltSMPMailerVersion"
	default:
		return fmt.Sprintf("GestaltSMPMailer(%d)", e)
	}
}

type GestaltSMPSPSendLetter uint

const (
	GestaltSMPSPSendLetterVersion GestaltSMPSPSendLetter = 's'<<24 | 'p'<<16 | 's'<<8 | 'l' // 'spsl'
)

func (e GestaltSMPSPSendLetter) String() string {
	switch e {
	case GestaltSMPSPSendLetterVersion:
		return "GestaltSMPSPSendLetterVersion"
	default:
		return fmt.Sprintf("GestaltSMPSPSendLetter(%d)", e)
	}
}

type GestaltScrapMgr uint

const (
	// GestaltScrapMgrAttr: The [Gestalt] selector you pass to determine which Scrap Manager attributes are present.
	GestaltScrapMgrAttr GestaltScrapMgr = 's'<<24 | 'c'<<16 | 'r'<<8 | 'a' // 'scra'
	// GestaltScrapMgrTranslationAware: If `true`, the Scrap Manager supports Translation Manager.
	GestaltScrapMgrTranslationAware GestaltScrapMgr = 0
)

func (e GestaltScrapMgr) String() string {
	switch e {
	case GestaltScrapMgrAttr:
		return "GestaltScrapMgrAttr"
	case GestaltScrapMgrTranslationAware:
		return "GestaltScrapMgrTranslationAware"
	default:
		return fmt.Sprintf("GestaltScrapMgr(%d)", e)
	}
}

type GestaltScreenCapture uint

const (
	GestaltScreenCaptureDir  GestaltScreenCapture = 'p'<<24 | 'i'<<16 | 'c'<<8 | '2' // 'pic2'
	GestaltScreenCaptureMain GestaltScreenCapture = 'p'<<24 | 'i'<<16 | 'c'<<8 | '1' // 'pic1'
)

func (e GestaltScreenCapture) String() string {
	switch e {
	case GestaltScreenCaptureDir:
		return "GestaltScreenCaptureDir"
	case GestaltScreenCaptureMain:
		return "GestaltScreenCaptureMain"
	default:
		return fmt.Sprintf("GestaltScreenCapture(%d)", e)
	}
}

type GestaltScript uint

const (
	// GestaltScriptCount: The selector you pass to the Gestalt function to determine the number of script systems currently active.
	GestaltScriptCount GestaltScript = 's'<<24 | 'c'<<16 | 'r'<<8 | '#' // 'scr#'
)

func (e GestaltScript) String() string {
	switch e {
	case GestaltScriptCount:
		return "GestaltScriptCount"
	default:
		return fmt.Sprintf("GestaltScript(%d)", e)
	}
}

type GestaltScriptMgr uint

const (
	// GestaltScriptMgrVersion: The selector you pass to the [Gestalt] function to determine the version number of the Script Manager (in the low-order word of the return value).
	GestaltScriptMgrVersion GestaltScriptMgr = 's'<<24 | 'c'<<16 | 'r'<<8 | 'i' // 'scri'
)

func (e GestaltScriptMgr) String() string {
	switch e {
	case GestaltScriptMgrVersion:
		return "GestaltScriptMgrVersion"
	default:
		return fmt.Sprintf("GestaltScriptMgr(%d)", e)
	}
}

type GestaltShutdown uint

const (
	GestaltShutdownAttributes            GestaltShutdown = 's'<<24 | 'h'<<16 | 'u'<<8 | 't' // 'shut'
	GestaltShutdownHassdOnBootVolUnmount GestaltShutdown = 0
)

func (e GestaltShutdown) String() string {
	switch e {
	case GestaltShutdownAttributes:
		return "GestaltShutdownAttributes"
	case GestaltShutdownHassdOnBootVolUnmount:
		return "GestaltShutdownHassdOnBootVolUnmount"
	default:
		return fmt.Sprintf("GestaltShutdown(%d)", e)
	}
}

type GestaltSoftwareVendor uint

const (
	GestaltSoftwareVendorApple    GestaltSoftwareVendor = 'A'<<24 | 'p'<<16 | 'p'<<8 | 'l' // 'Appl'
	GestaltSoftwareVendorCode     GestaltSoftwareVendor = 's'<<24 | 'r'<<16 | 'a'<<8 | 'd' // 'srad'
	GestaltSoftwareVendorLicensee GestaltSoftwareVendor = 'L'<<24 | 'c'<<16 | 'n'<<8 | 's' // 'Lcns'
)

func (e GestaltSoftwareVendor) String() string {
	switch e {
	case GestaltSoftwareVendorApple:
		return "GestaltSoftwareVendorApple"
	case GestaltSoftwareVendorCode:
		return "GestaltSoftwareVendorCode"
	case GestaltSoftwareVendorLicensee:
		return "GestaltSoftwareVendorLicensee"
	default:
		return fmt.Sprintf("GestaltSoftwareVendor(%d)", e)
	}
}

type GestaltSpeech uint

const (
	// GestaltSpeechAttr: The selector you pass to the [Gestalt] function to determine the Speech Manager attributes.
	GestaltSpeechAttr       GestaltSpeech = 't'<<24 | 't'<<16 | 's'<<8 | 'c' // 'ttsc'
	GestaltSpeechHasPPCGlue GestaltSpeech = 1
	GestaltSpeechMgrPresent GestaltSpeech = 0
)

func (e GestaltSpeech) String() string {
	switch e {
	case GestaltSpeechAttr:
		return "GestaltSpeechAttr"
	case GestaltSpeechHasPPCGlue:
		return "GestaltSpeechHasPPCGlue"
	case GestaltSpeechMgrPresent:
		return "GestaltSpeechMgrPresent"
	default:
		return fmt.Sprintf("GestaltSpeech(%d)", e)
	}
}

type GestaltSpeechRecognition uint

const (
	GestaltSpeechRecognitionVersion GestaltSpeechRecognition = 's'<<24 | 'r'<<16 | 't'<<8 | 'b' // 'srtb'
)

func (e GestaltSpeechRecognition) String() string {
	switch e {
	case GestaltSpeechRecognitionVersion:
		return "GestaltSpeechRecognitionVersion"
	default:
		return fmt.Sprintf("GestaltSpeechRecognition(%d)", e)
	}
}

type GestaltSplitOS uint

const (
	GestaltSplitOSAttr                                   GestaltSplitOS = 's'<<24 | 'p'<<16 | 'o'<<8 | 's' // 'spos'
	GestaltSplitOSAware                                  GestaltSplitOS = 1
	GestaltSplitOSBootDriveIsNetworkVolume               GestaltSplitOS = 0
	GestaltSplitOSEnablerVolumeIsDifferentFromBootVolume GestaltSplitOS = 2
	GestaltSplitOSMachineNameSetToNetworkNameTemp        GestaltSplitOS = 3
	GestaltSplitOSMachineNameStartupDiskIsNonPersistent  GestaltSplitOS = 5
)

func (e GestaltSplitOS) String() string {
	switch e {
	case GestaltSplitOSAttr:
		return "GestaltSplitOSAttr"
	case GestaltSplitOSAware:
		return "GestaltSplitOSAware"
	case GestaltSplitOSBootDriveIsNetworkVolume:
		return "GestaltSplitOSBootDriveIsNetworkVolume"
	case GestaltSplitOSEnablerVolumeIsDifferentFromBootVolume:
		return "GestaltSplitOSEnablerVolumeIsDifferentFromBootVolume"
	case GestaltSplitOSMachineNameSetToNetworkNameTemp:
		return "GestaltSplitOSMachineNameSetToNetworkNameTemp"
	case GestaltSplitOSMachineNameStartupDiskIsNonPersistent:
		return "GestaltSplitOSMachineNameStartupDiskIsNonPersistent"
	default:
		return fmt.Sprintf("GestaltSplitOS(%d)", e)
	}
}

type GestaltStandard uint

const (
	// GestaltStandardFile58: # Discussion
	GestaltStandardFile58 GestaltStandard = 0
	// GestaltStandardFileAttr: The selector you pass to the Gestalt function to determine the Standard File Package attributes.
	GestaltStandardFileAttr                       GestaltStandard = 's'<<24 | 't'<<16 | 'd'<<8 | 'f' // 'stdf'
	GestaltStandardFileHasColorIcons              GestaltStandard = 2
	GestaltStandardFileHasDynamicVolumeAllocation GestaltStandard = 4
	GestaltStandardFileTranslationAware           GestaltStandard = 1
	GestaltStandardFileUseGenericIcons            GestaltStandard = 3
)

func (e GestaltStandard) String() string {
	switch e {
	case GestaltStandardFile58:
		return "GestaltStandardFile58"
	case GestaltStandardFileAttr:
		return "GestaltStandardFileAttr"
	case GestaltStandardFileHasColorIcons:
		return "GestaltStandardFileHasColorIcons"
	case GestaltStandardFileHasDynamicVolumeAllocation:
		return "GestaltStandardFileHasDynamicVolumeAllocation"
	case GestaltStandardFileTranslationAware:
		return "GestaltStandardFileTranslationAware"
	case GestaltStandardFileUseGenericIcons:
		return "GestaltStandardFileUseGenericIcons"
	default:
		return fmt.Sprintf("GestaltStandard(%d)", e)
	}
}

type GestaltStdNBP uint

const (
	// GestaltStdNBPAttr: The selector you pass to the [Gestalt] function to determine information about the StandardNBP (Name-Binding Protocol) function.
	GestaltStdNBPAttr                 GestaltStdNBP = 'n'<<24 | 'l'<<16 | 'u'<<8 | 'p' // 'nlup'
	GestaltStdNBPPresent              GestaltStdNBP = 0
	GestaltStdNBPSupportsAutoPosition GestaltStdNBP = 1
)

func (e GestaltStdNBP) String() string {
	switch e {
	case GestaltStdNBPAttr:
		return "GestaltStdNBPAttr"
	case GestaltStdNBPPresent:
		return "GestaltStdNBPPresent"
	case GestaltStdNBPSupportsAutoPosition:
		return "GestaltStdNBPSupportsAutoPosition"
	default:
		return fmt.Sprintf("GestaltStdNBP(%d)", e)
	}
}

type GestaltSystem uint

const (
	// Deprecated.
	GestaltSystemVersion GestaltSystem = 's'<<24 | 'y'<<16 | 's'<<8 | 'v' // 'sysv'
	// Deprecated.
	GestaltSystemVersionBugFix GestaltSystem = 's'<<24 | 'y'<<16 | 's'<<8 | '3' // 'sys3'
	// Deprecated.
	GestaltSystemVersionMajor GestaltSystem = 's'<<24 | 'y'<<16 | 's'<<8 | '1' // 'sys1'
	// Deprecated.
	GestaltSystemVersionMinor GestaltSystem = 's'<<24 | 'y'<<16 | 's'<<8 | '2' // 'sys2'
)

func (e GestaltSystem) String() string {
	switch e {
	case GestaltSystemVersion:
		return "GestaltSystemVersion"
	case GestaltSystemVersionBugFix:
		return "GestaltSystemVersionBugFix"
	case GestaltSystemVersionMajor:
		return "GestaltSystemVersionMajor"
	case GestaltSystemVersionMinor:
		return "GestaltSystemVersionMinor"
	default:
		return fmt.Sprintf("GestaltSystem(%d)", e)
	}
}

type GestaltSystemUpdate uint

const (
	GestaltSystemUpdateVersion GestaltSystemUpdate = 's'<<24 | 'y'<<16 | 's'<<8 | 'u' // 'sysu'
)

func (e GestaltSystemUpdate) String() string {
	switch e {
	case GestaltSystemUpdateVersion:
		return "GestaltSystemUpdateVersion"
	default:
		return fmt.Sprintf("GestaltSystemUpdate(%d)", e)
	}
}

type GestaltT uint

const (
	GestaltTE6 GestaltT = 6
)

func (e GestaltT) String() string {
	switch e {
	case GestaltTE6:
		return "GestaltTE6"
	default:
		return fmt.Sprintf("GestaltT(%d)", e)
	}
}

type GestaltTE uint

const (
	// GestaltTEAttr: The [Gestalt] selector you pass to determine whichTextEdit attributes are present.
	GestaltTEAttr GestaltTE = 't'<<24 | 'e'<<16 | 'a'<<8 | 't' // 'teat'
	// GestaltTEHasGetHiliteRgn: If true, TextEdit has [TEGetHiliteRgn].
	GestaltTEHasGetHiliteRgn GestaltTE = 0
	// GestaltTEHasWhiteBackground: If `true`, TextEdit supports overriding the [TERec]' data structure background field to white.
	GestaltTEHasWhiteBackground GestaltTE = 3
	// GestaltTESupportsInlineInput: If `true`, TextEdit does Inline Input.
	GestaltTESupportsInlineInput GestaltTE = 1
	// GestaltTESupportsTextObjects: If `true`, TextEdit does Text Objects.
	GestaltTESupportsTextObjects GestaltTE = 2
)

func (e GestaltTE) String() string {
	switch e {
	case GestaltTEAttr:
		return "GestaltTEAttr"
	case GestaltTEHasGetHiliteRgn:
		return "GestaltTEHasGetHiliteRgn"
	case GestaltTEHasWhiteBackground:
		return "GestaltTEHasWhiteBackground"
	case GestaltTESupportsInlineInput:
		return "GestaltTESupportsInlineInput"
	case GestaltTESupportsTextObjects:
		return "GestaltTESupportsTextObjects"
	default:
		return fmt.Sprintf("GestaltTE(%d)", e)
	}
}

type GestaltTS uint

const (
	GestaltTSMDisplayMgrAwareBit GestaltTS = 0
	GestaltTSMdoesTSMTEBit       GestaltTS = 1
	GestaltTSMgr15               GestaltTS = 336
	GestaltTSMgr20               GestaltTS = 512
	GestaltTSMgr22               GestaltTS = 544
	GestaltTSMgr23               GestaltTS = 560
	GestaltTSMgrAttr             GestaltTS = 't'<<24 | 's'<<16 | 'm'<<8 | 'a' // 'tsma'
	// GestaltTSMgrVersion: The selector you pass to the [Gestalt] function to determine the version of the Text Services Manager.
	GestaltTSMgrVersion GestaltTS = 't'<<24 | 's'<<16 | 'm'<<8 | 'v' // 'tsmv'
)

func (e GestaltTS) String() string {
	switch e {
	case GestaltTSMDisplayMgrAwareBit:
		return "GestaltTSMDisplayMgrAwareBit"
	case GestaltTSMdoesTSMTEBit:
		return "GestaltTSMdoesTSMTEBit"
	case GestaltTSMgr15:
		return "GestaltTSMgr15"
	case GestaltTSMgr20:
		return "GestaltTSMgr20"
	case GestaltTSMgr22:
		return "GestaltTSMgr22"
	case GestaltTSMgr23:
		return "GestaltTSMgr23"
	case GestaltTSMgrAttr:
		return "GestaltTSMgrAttr"
	case GestaltTSMgrVersion:
		return "GestaltTSMgrVersion"
	default:
		return fmt.Sprintf("GestaltTS(%d)", e)
	}
}

type GestaltTSMT uint

const (
	GestaltTSMTE        GestaltTSMT = 0
	GestaltTSMTE1       GestaltTSMT = 256
	GestaltTSMTE15      GestaltTSMT = 336
	GestaltTSMTE152     GestaltTSMT = 338
	GestaltTSMTEAttr    GestaltTSMT = 't'<<24 | 'm'<<16 | 'T'<<8 | 'E' // 'tmTE'
	GestaltTSMTEPresent GestaltTSMT = 0
	GestaltTSMTEVersion GestaltTSMT = 't'<<24 | 'm'<<16 | 'T'<<8 | 'V' // 'tmTV'
)

func (e GestaltTSMT) String() string {
	switch e {
	case GestaltTSMTE:
		return "GestaltTSMTE"
	case GestaltTSMTE1:
		return "GestaltTSMTE1"
	case GestaltTSMTE15:
		return "GestaltTSMTE15"
	case GestaltTSMTE152:
		return "GestaltTSMTE152"
	case GestaltTSMTEAttr:
		return "GestaltTSMTEAttr"
	case GestaltTSMTEVersion:
		return "GestaltTSMTEVersion"
	default:
		return fmt.Sprintf("GestaltTSMT(%d)", e)
	}
}

type GestaltTeleMgr uint

const (
	GestaltTeleMgrAttr             GestaltTeleMgr = 't'<<24 | 'e'<<16 | 'l'<<8 | 'e' // 'tele'
	GestaltTeleMgrAutoAnswer       GestaltTeleMgr = 3
	GestaltTeleMgrIndHandset       GestaltTeleMgr = 4
	GestaltTeleMgrNewTELNewSupport GestaltTeleMgr = 6
	GestaltTeleMgrPowerPCSupport   GestaltTeleMgr = 1
	GestaltTeleMgrPresent          GestaltTeleMgr = 0
	GestaltTeleMgrSilenceDetect    GestaltTeleMgr = 5
	GestaltTeleMgrSoundStreams     GestaltTeleMgr = 2
)

func (e GestaltTeleMgr) String() string {
	switch e {
	case GestaltTeleMgrAttr:
		return "GestaltTeleMgrAttr"
	case GestaltTeleMgrAutoAnswer:
		return "GestaltTeleMgrAutoAnswer"
	case GestaltTeleMgrIndHandset:
		return "GestaltTeleMgrIndHandset"
	case GestaltTeleMgrNewTELNewSupport:
		return "GestaltTeleMgrNewTELNewSupport"
	case GestaltTeleMgrPowerPCSupport:
		return "GestaltTeleMgrPowerPCSupport"
	case GestaltTeleMgrPresent:
		return "GestaltTeleMgrPresent"
	case GestaltTeleMgrSilenceDetect:
		return "GestaltTeleMgrSilenceDetect"
	case GestaltTeleMgrSoundStreams:
		return "GestaltTeleMgrSoundStreams"
	default:
		return fmt.Sprintf("GestaltTeleMgr(%d)", e)
	}
}

type GestaltTermMgr uint

const (
	// GestaltTermMgrAttr: The selector you pass to the [Gestalt] function to determine the Terminal Manager attributes.
	GestaltTermMgrAttr        GestaltTermMgr = 't'<<24 | 'e'<<16 | 'r'<<8 | 'm' // 'term'
	GestaltTermMgrErrorString GestaltTermMgr = 2
	GestaltTermMgrPresent     GestaltTermMgr = 0
)

func (e GestaltTermMgr) String() string {
	switch e {
	case GestaltTermMgrAttr:
		return "GestaltTermMgrAttr"
	case GestaltTermMgrErrorString:
		return "GestaltTermMgrErrorString"
	case GestaltTermMgrPresent:
		return "GestaltTermMgrPresent"
	default:
		return fmt.Sprintf("GestaltTermMgr(%d)", e)
	}
}

type GestaltToolbox uint

const (
	// GestaltToolboxTable: The selector you pass to the [Gestalt] function to determine the base address of the Toolbox trap dispatch table.
	GestaltToolboxTable GestaltToolbox = 't'<<24 | 'b'<<16 | 't'<<8 | 't' // 'tbtt'
)

func (e GestaltToolbox) String() string {
	switch e {
	case GestaltToolboxTable:
		return "GestaltToolboxTable"
	default:
		return fmt.Sprintf("GestaltToolbox(%d)", e)
	}
}

type GestaltTranslation uint

const (
	// GestaltTranslationAttr: The [Gestalt] selector you pass to determine which Translation Manager attributes are present.
	GestaltTranslationAttr GestaltTranslation = 'x'<<24 | 'l'<<16 | 'a'<<8 | 't' // 'xlat'
	// GestaltTranslationGetPathAPIAvail: If true, the functions [GetFileTranslationPath] and [GetPathTranslationDialog] are available.
	GestaltTranslationGetPathAPIAvail GestaltTranslation = 3
	// GestaltTranslationMgrExists: If `true`, the Translation Manager is present.
	GestaltTranslationMgrExists GestaltTranslation = 0
	// GestaltTranslationMgrHintOrder: In earlier versions of the Translation Manager, the scrap hints in the [DoTranslateScrapProcPtr] function were reversed.
	GestaltTranslationMgrHintOrder GestaltTranslation = 1
	// GestaltTranslationPPCAvail: If `true`, the PowerPC Translation Library is available, and you can call the Translation Manager from native PowerPC code.
	GestaltTranslationPPCAvail GestaltTranslation = 2
)

func (e GestaltTranslation) String() string {
	switch e {
	case GestaltTranslationAttr:
		return "GestaltTranslationAttr"
	case GestaltTranslationGetPathAPIAvail:
		return "GestaltTranslationGetPathAPIAvail"
	case GestaltTranslationMgrExists:
		return "GestaltTranslationMgrExists"
	case GestaltTranslationMgrHintOrder:
		return "GestaltTranslationMgrHintOrder"
	case GestaltTranslationPPCAvail:
		return "GestaltTranslationPPCAvail"
	default:
		return fmt.Sprintf("GestaltTranslation(%d)", e)
	}
}

type GestaltUDF uint

const (
	GestaltUDFSupport GestaltUDF = 'k'<<24 | 'u'<<16 | 'd'<<8 | 'f' // 'kudf'
)

func (e GestaltUDF) String() string {
	switch e {
	case GestaltUDFSupport:
		return "GestaltUDFSupport"
	default:
		return fmt.Sprintf("GestaltUDF(%d)", e)
	}
}

type GestaltUSB uint

const (
	GestaltUSBAttr     GestaltUSB = 'u'<<24 | 's'<<16 | 'b'<<8 | ' ' // 'usb '
	GestaltUSBHasIsoch GestaltUSB = 1
	GestaltUSBPresent  GestaltUSB = 0
	GestaltUSBVersion  GestaltUSB = 'u'<<24 | 's'<<16 | 'b'<<8 | 'v' // 'usbv'
)

func (e GestaltUSB) String() string {
	switch e {
	case GestaltUSBAttr:
		return "GestaltUSBAttr"
	case GestaltUSBHasIsoch:
		return "GestaltUSBHasIsoch"
	case GestaltUSBPresent:
		return "GestaltUSBPresent"
	case GestaltUSBVersion:
		return "GestaltUSBVersion"
	default:
		return fmt.Sprintf("GestaltUSB(%d)", e)
	}
}

type GestaltUSBPrinterSharing uint

const (
	GestaltUSBPrinterSharingAttr        GestaltUSBPrinterSharing = 'z'<<24 | 'a'<<16 | 'k'<<8 | ' ' // 'zak '
	GestaltUSBPrinterSharingAttrBooted  GestaltUSBPrinterSharing = 1073741824
	GestaltUSBPrinterSharingAttrMask    GestaltUSBPrinterSharing = 0
	GestaltUSBPrinterSharingAttrRunning GestaltUSBPrinterSharing = 0
	GestaltUSBPrinterSharingVersion     GestaltUSBPrinterSharing = 'z'<<24 | 'a'<<16 | 'k'<<8 | ' ' // 'zak '
	GestaltUSBPrinterSharingVersionMask GestaltUSBPrinterSharing = 65535
)

func (e GestaltUSBPrinterSharing) String() string {
	switch e {
	case GestaltUSBPrinterSharingAttr:
		return "GestaltUSBPrinterSharingAttr"
	case GestaltUSBPrinterSharingAttrBooted:
		return "GestaltUSBPrinterSharingAttrBooted"
	case GestaltUSBPrinterSharingAttrMask:
		return "GestaltUSBPrinterSharingAttrMask"
	case GestaltUSBPrinterSharingVersionMask:
		return "GestaltUSBPrinterSharingVersionMask"
	default:
		return fmt.Sprintf("GestaltUSBPrinterSharing(%d)", e)
	}
}

type GestaltUserVisibleMachine uint

const (
	GestaltUserVisibleMachineName GestaltUserVisibleMachine = 'm'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'mnam'
)

func (e GestaltUserVisibleMachine) String() string {
	switch e {
	case GestaltUserVisibleMachineName:
		return "GestaltUserVisibleMachineName"
	default:
		return fmt.Sprintf("GestaltUserVisibleMachine(%d)", e)
	}
}

type GestaltVIA1 uint

const (
	GestaltVIA1Addr GestaltVIA1 = 'v'<<24 | 'i'<<16 | 'a'<<8 | '1' // 'via1'
)

func (e GestaltVIA1) String() string {
	switch e {
	case GestaltVIA1Addr:
		return "GestaltVIA1Addr"
	default:
		return fmt.Sprintf("GestaltVIA1(%d)", e)
	}
}

type GestaltVIA2 uint

const (
	GestaltVIA2Addr GestaltVIA2 = 'v'<<24 | 'i'<<16 | 'a'<<8 | '2' // 'via2'
)

func (e GestaltVIA2) String() string {
	switch e {
	case GestaltVIA2Addr:
		return "GestaltVIA2Addr"
	default:
		return fmt.Sprintf("GestaltVIA2(%d)", e)
	}
}

type GestaltVM uint

const (
	// GestaltVMAttr: The [Gestalt] selector you pass to determine the virtual memory attributes that are present.
	GestaltVMAttr                   GestaltVM = 'v'<<24 | 'm'<<16 | ' '<<8 | ' ' // 'vm  '
	GestaltVMFilemappingOn          GestaltVM = 3
	GestaltVMHasLockMemoryForOutput GestaltVM = 1
	GestaltVMHasPagingControl       GestaltVM = 4
	// GestaltVMPresent: If `true`, virtual memory is present.
	GestaltVMPresent GestaltVM = 0
)

func (e GestaltVM) String() string {
	switch e {
	case GestaltVMAttr:
		return "GestaltVMAttr"
	case GestaltVMFilemappingOn:
		return "GestaltVMFilemappingOn"
	case GestaltVMHasLockMemoryForOutput:
		return "GestaltVMHasLockMemoryForOutput"
	case GestaltVMHasPagingControl:
		return "GestaltVMHasPagingControl"
	case GestaltVMPresent:
		return "GestaltVMPresent"
	default:
		return fmt.Sprintf("GestaltVM(%d)", e)
	}
}

type GestaltVMBackingStoreFileRef uint

const (
	GestaltVMBackingStoreFileRefNum GestaltVMBackingStoreFileRef = 'v'<<24 | 'm'<<16 | 'b'<<8 | 's' // 'vmbs'
)

func (e GestaltVMBackingStoreFileRef) String() string {
	switch e {
	case GestaltVMBackingStoreFileRefNum:
		return "GestaltVMBackingStoreFileRefNum"
	default:
		return fmt.Sprintf("GestaltVMBackingStoreFileRef(%d)", e)
	}
}

type GestaltVMInfo uint

const (
	GestaltVMInfoNoneType        GestaltVMInfo = 3
	GestaltVMInfoSimpleType      GestaltVMInfo = 2
	GestaltVMInfoSizeStorageType GestaltVMInfo = 0
	GestaltVMInfoSizeType        GestaltVMInfo = 1
	GestaltVMInfoType            GestaltVMInfo = 'v'<<24 | 'm'<<16 | 'i'<<8 | 'n' // 'vmin'
)

func (e GestaltVMInfo) String() string {
	switch e {
	case GestaltVMInfoNoneType:
		return "GestaltVMInfoNoneType"
	case GestaltVMInfoSimpleType:
		return "GestaltVMInfoSimpleType"
	case GestaltVMInfoSizeStorageType:
		return "GestaltVMInfoSizeStorageType"
	case GestaltVMInfoSizeType:
		return "GestaltVMInfoSizeType"
	case GestaltVMInfoType:
		return "GestaltVMInfoType"
	default:
		return fmt.Sprintf("GestaltVMInfo(%d)", e)
	}
}

type GestaltX86 uint

const (
	GestaltX86AdditionalFeatures  GestaltX86 = 'x'<<24 | '8'<<16 | '6'<<8 | 'a' // 'x86a'
	GestaltX86Features            GestaltX86 = 'x'<<24 | '8'<<16 | '6'<<8 | 'f' // 'x86f'
	GestaltX86HasAPIC             GestaltX86 = 9
	GestaltX86HasCID              GestaltX86 = 10
	GestaltX86HasCLFSH            GestaltX86 = 19
	GestaltX86HasCMOV             GestaltX86 = 15
	GestaltX86HasCX16             GestaltX86 = 13
	GestaltX86HasCX8              GestaltX86 = 8
	GestaltX86HasDE               GestaltX86 = 2
	GestaltX86HasDS               GestaltX86 = 21
	GestaltX86HasDSCPL            GestaltX86 = 4
	GestaltX86HasEST              GestaltX86 = 7
	GestaltX86HasFPU              GestaltX86 = 0
	GestaltX86HasFXSR             GestaltX86 = 24
	GestaltX86HasHTT              GestaltX86 = 28
	GestaltX86HasMCA              GestaltX86 = 14
	GestaltX86HasMCE              GestaltX86 = 7
	GestaltX86HasMMX              GestaltX86 = 23
	GestaltX86HasMONITOR          GestaltX86 = 3
	GestaltX86HasMSR              GestaltX86 = 5
	GestaltX86HasMTRR             GestaltX86 = 12
	GestaltX86HasPAE              GestaltX86 = 6
	GestaltX86HasPAT              GestaltX86 = 16
	GestaltX86HasPGE              GestaltX86 = 13
	GestaltX86HasPSE              GestaltX86 = 3
	GestaltX86HasPSE36            GestaltX86 = 17
	GestaltX86HasPSN              GestaltX86 = 18
	GestaltX86HasSEP              GestaltX86 = 11
	GestaltX86HasSMX              GestaltX86 = 6
	GestaltX86HasSS               GestaltX86 = 27
	GestaltX86HasSSE              GestaltX86 = 25
	GestaltX86HasSSE2             GestaltX86 = 26
	GestaltX86HasSSE3             GestaltX86 = 0
	GestaltX86HasSupplementalSSE3 GestaltX86 = 9
	GestaltX86HasTM               GestaltX86 = 29
	GestaltX86HasTM2              GestaltX86 = 8
	GestaltX86HasTSC              GestaltX86 = 4
	GestaltX86HasVME              GestaltX86 = 1
	GestaltX86HasVMX              GestaltX86 = 5
	GestaltX86HasxTPR             GestaltX86 = 14
	GestaltX86ResACPI             GestaltX86 = 22
	GestaltX86Serviced20          GestaltX86 = 20
)

func (e GestaltX86) String() string {
	switch e {
	case GestaltX86AdditionalFeatures:
		return "GestaltX86AdditionalFeatures"
	case GestaltX86Features:
		return "GestaltX86Features"
	case GestaltX86HasAPIC:
		return "GestaltX86HasAPIC"
	case GestaltX86HasCID:
		return "GestaltX86HasCID"
	case GestaltX86HasCLFSH:
		return "GestaltX86HasCLFSH"
	case GestaltX86HasCMOV:
		return "GestaltX86HasCMOV"
	case GestaltX86HasCX16:
		return "GestaltX86HasCX16"
	case GestaltX86HasCX8:
		return "GestaltX86HasCX8"
	case GestaltX86HasDE:
		return "GestaltX86HasDE"
	case GestaltX86HasDS:
		return "GestaltX86HasDS"
	case GestaltX86HasDSCPL:
		return "GestaltX86HasDSCPL"
	case GestaltX86HasEST:
		return "GestaltX86HasEST"
	case GestaltX86HasFPU:
		return "GestaltX86HasFPU"
	case GestaltX86HasFXSR:
		return "GestaltX86HasFXSR"
	case GestaltX86HasHTT:
		return "GestaltX86HasHTT"
	case GestaltX86HasMCA:
		return "GestaltX86HasMCA"
	case GestaltX86HasMMX:
		return "GestaltX86HasMMX"
	case GestaltX86HasMONITOR:
		return "GestaltX86HasMONITOR"
	case GestaltX86HasMSR:
		return "GestaltX86HasMSR"
	case GestaltX86HasMTRR:
		return "GestaltX86HasMTRR"
	case GestaltX86HasPAE:
		return "GestaltX86HasPAE"
	case GestaltX86HasPAT:
		return "GestaltX86HasPAT"
	case GestaltX86HasPSE36:
		return "GestaltX86HasPSE36"
	case GestaltX86HasPSN:
		return "GestaltX86HasPSN"
	case GestaltX86HasSEP:
		return "GestaltX86HasSEP"
	case GestaltX86HasSS:
		return "GestaltX86HasSS"
	case GestaltX86HasSSE:
		return "GestaltX86HasSSE"
	case GestaltX86HasSSE2:
		return "GestaltX86HasSSE2"
	case GestaltX86HasTM:
		return "GestaltX86HasTM"
	case GestaltX86HasVME:
		return "GestaltX86HasVME"
	case GestaltX86ResACPI:
		return "GestaltX86ResACPI"
	case GestaltX86Serviced20:
		return "GestaltX86Serviced20"
	default:
		return fmt.Sprintf("GestaltX86(%d)", e)
	}
}

type Greaterthan uint

const (
	EQUALTO     Greaterthan = 2
	GREATERTHAN Greaterthan = 0
	LESSTHAN    Greaterthan = 1
	UNORDERED   Greaterthan = 3
)

func (e Greaterthan) String() string {
	switch e {
	case EQUALTO:
		return "EQUALTO"
	case GREATERTHAN:
		return "GREATERTHAN"
	case LESSTHAN:
		return "LESSTHAN"
	case UNORDERED:
		return "UNORDERED"
	default:
		return fmt.Sprintf("Greaterthan(%d)", e)
	}
}

type Hm int

const (
	HmBalloonAborted       Hm = -853
	HmCloseViewActive      Hm = -863
	HmHelpDisabled         Hm = -850
	HmHelpManagerNotInited Hm = -855
	HmNoBalloonUp          Hm = -862
	HmOperationUnsupported Hm = -861
	HmSameAsLastBalloon    Hm = -854
	HmSkippedBalloon       Hm = -857
	HmUnknownHelpType      Hm = -859
	HmWrongVersion         Hm = -858
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
	HrMiscellaneousExceptionErr       Hr = -5361
	HrURLNotHandledErr                Hr = -5363
	HrUnableToResizeHandleErr         Hr = -5362
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
	IIOAbort         IMemFullErr = -27
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

type Int uint

const (
	// Deprecated.
	IntArabic Int = 1
	// Deprecated.
	IntEuropean Int = 4
	// Deprecated.
	IntJapanese Int = 3
	// Deprecated.
	IntOutputMask Int = 32768
	// Deprecated.
	IntRoman Int = 2
	// Deprecated.
	IntWestern Int = 0
)

func (e Int) String() string {
	switch e {
	case IntArabic:
		return "IntArabic"
	case IntEuropean:
		return "IntEuropean"
	case IntJapanese:
		return "IntJapanese"
	case IntOutputMask:
		return "IntOutputMask"
	case IntRoman:
		return "IntRoman"
	case IntWestern:
		return "IntWestern"
	default:
		return fmt.Sprintf("Int(%d)", e)
	}
}

type InternalComponentErr int

const (
	CantReceiveFromSynthesizerOSErr InternalComponentErr = -2073
	CantSendToSynthesizerOSErr      InternalComponentErr = -2072
	IllegalChannelOSErr             InternalComponentErr = -2076
	IllegalControllerOSErr          InternalComponentErr = -2080
	IllegalInstrumentOSErr          InternalComponentErr = -2079
	IllegalKnobOSErr                InternalComponentErr = -2077
	IllegalKnobValueOSErr           InternalComponentErr = -2078
	IllegalNoteChannelOSErr         InternalComponentErr = -2084
	IllegalPartOSErr                InternalComponentErr = -2075
	IllegalVoiceAllocationOSErr     InternalComponentErr = -2074
	InternalComponentErrValue       InternalComponentErr = -2070
	MidiManagerAbsentOSErr          InternalComponentErr = -2081
	NoExportProcAvailableErr        InternalComponentErr = -2089
	NotImplementedMusicOSErr        InternalComponentErr = -2071
	NoteChannelNotAllocatedOSErr    InternalComponentErr = -2085
	SynthesizerNotRespondingOSErr   InternalComponentErr = -2082
	SynthesizerOSErr                InternalComponentErr = -2083
	TuneParseOSErr                  InternalComponentErr = -2087
	TunePlayerFullOSErr             InternalComponentErr = -2086
	VideoOutputInUseErr             InternalComponentErr = -2090
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
	BadScrapRefErr              InternalScrapErr = -4990
	DuplicateScrapFlavorErr     InternalScrapErr = -4989
	IllegalScrapFlavorFlagsErr  InternalScrapErr = -4997
	IllegalScrapFlavorSizeErr   InternalScrapErr = -4999
	IllegalScrapFlavorTypeErr   InternalScrapErr = -4998
	InternalScrapErrValue       InternalScrapErr = -4988
	NeedClearScrapErr           InternalScrapErr = -100
	NilScrapFlavorDataErr       InternalScrapErr = -4994
	NoScrapPromiseKeeperErr     InternalScrapErr = -4993
	ProcessStateIncorrectErr    InternalScrapErr = -4991
	ScrapFlavorFlagsMismatchErr InternalScrapErr = -4995
	ScrapFlavorNotFoundErr      InternalScrapErr = -102
	ScrapFlavorSizeMismatchErr  InternalScrapErr = -4996
	ScrapPromiseNotKeptErr      InternalScrapErr = -4992
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
	InvalidComponentIDValue       InvalidComponentID = -3000
	RetryComponentRegistrationErr InvalidComponentID = -3005
	UnresolvedComponentDLLErr     InvalidComponentID = -3004
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
	NoSuchIconErr          InvalidIconRefErr = -2581
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
	BadTranslationSpecErr          InvalidTranslationPathErr = -3031
	CouldNotParseSourceFileErr     InvalidTranslationPathErr = -3026
	InvalidTranslationPathErrValue InvalidTranslationPathErr = -3025
	NoPrefAppErr                   InvalidTranslationPathErr = -3032
	NoTranslationPathErr           InvalidTranslationPathErr = -3030
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

type Itlc uint

const (
	ItlcDualCaret    Itlc = 6
	ItlcShowIcon     Itlc = 7
	ItlcSysDirection Itlc = 15
)

func (e Itlc) String() string {
	switch e {
	case ItlcDualCaret:
		return "ItlcDualCaret"
	case ItlcShowIcon:
		return "ItlcShowIcon"
	case ItlcSysDirection:
		return "ItlcSysDirection"
	default:
		return fmt.Sprintf("Itlc(%d)", e)
	}
}

type ItlcDisableKeyScript uint

const (
	ItlcDisableKeyScriptSyncValue ItlcDisableKeyScript = 3
)

func (e ItlcDisableKeyScript) String() string {
	switch e {
	case ItlcDisableKeyScriptSyncValue:
		return "ItlcDisableKeyScriptSyncValue"
	default:
		return fmt.Sprintf("ItlcDisableKeyScript(%d)", e)
	}
}

type ItlcDisableKeyScriptSync uint

const (
	ItlcDisableKeyScriptSyncMask ItlcDisableKeyScriptSync = 8
)

func (e ItlcDisableKeyScriptSync) String() string {
	switch e {
	case ItlcDisableKeyScriptSyncMask:
		return "ItlcDisableKeyScriptSyncMask"
	default:
		return fmt.Sprintf("ItlcDisableKeyScriptSync(%d)", e)
	}
}

type Iu int

const (
	IuCurrentCurLang Iu = -4
	IuCurrentDefLang Iu = -5
	IuScriptCurLang  Iu = -6
	IuScriptDefLang  Iu = -7
	IuSystemCurLang  Iu = -2
	IuSystemDefLang  Iu = -3
	// Deprecated.
	IuCurrentScript Iu = -2
	// Deprecated.
	IuSystemScript Iu = -1
)

func (e Iu) String() string {
	switch e {
	case IuCurrentCurLang:
		return "IuCurrentCurLang"
	case IuCurrentDefLang:
		return "IuCurrentDefLang"
	case IuScriptCurLang:
		return "IuScriptCurLang"
	case IuScriptDefLang:
		return "IuScriptDefLang"
	case IuSystemCurLang:
		return "IuSystemCurLang"
	case IuSystemDefLang:
		return "IuSystemDefLang"
	case IuSystemScript:
		return "IuSystemScript"
	default:
		return fmt.Sprintf("Iu(%d)", e)
	}
}

type K32BitHeap uint

const (
	// Deprecated.
	K32BitHeapValue K32BitHeap = 1
	// Deprecated.
	KNewDebugHeap K32BitHeap = 4
	// Deprecated.
	KNewStyleHeap K32BitHeap = 2
)

func (e K32BitHeap) String() string {
	switch e {
	case K32BitHeapValue:
		return "K32BitHeapValue"
	case KNewDebugHeap:
		return "KNewDebugHeap"
	case KNewStyleHeap:
		return "KNewStyleHeap"
	default:
		return fmt.Sprintf("K32BitHeap(%d)", e)
	}
}

type K68kInterruptLevelMask uint

const (
	// Deprecated.
	K68kInterruptLevelMaskValue K68kInterruptLevelMask = 7
	// Deprecated.
	KInDeferredTaskMask K68kInterruptLevelMask = 32
	// Deprecated.
	KInNestedInterruptMask K68kInterruptLevelMask = 128
	// Deprecated.
	KInSecondaryIntHandlerMask K68kInterruptLevelMask = 64
	// Deprecated.
	KInVBLTaskMask K68kInterruptLevelMask = 16
)

func (e K68kInterruptLevelMask) String() string {
	switch e {
	case K68kInterruptLevelMaskValue:
		return "K68kInterruptLevelMaskValue"
	case KInDeferredTaskMask:
		return "KInDeferredTaskMask"
	case KInNestedInterruptMask:
		return "KInNestedInterruptMask"
	case KInSecondaryIntHandlerMask:
		return "KInSecondaryIntHandlerMask"
	case KInVBLTaskMask:
		return "KInVBLTaskMask"
	default:
		return fmt.Sprintf("K68kInterruptLevelMask(%d)", e)
	}
}

type KAE uint

const (
	KAEActivate KAE = 'a'<<24 | 'c'<<16 | 't'<<8 | 'v' // 'actv'
	// KAEAlwaysInteract: # Discussion
	KAEAlwaysInteract KAE = 48
	// KAEAnswer: Event that is a reply Apple event.
	KAEAnswer           KAE = 'a'<<24 | 'n'<<16 | 's'<<8 | 'r' // 'ansr'
	KAEApplicationClass KAE = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'l' // 'appl'
	// KAEApplicationDied: Event sent by the Process Manager to an application that launched another application when the launched application quits or terminates.
	KAEApplicationDied  KAE = 'o'<<24 | 'b'<<16 | 'i'<<8 | 't' // 'obit'
	KAEAsk              KAE = 'a'<<24 | 's'<<16 | 'k'<<8 | ' ' // 'ask '
	KAEAutoDown         KAE = 'a'<<24 | 'u'<<16 | 't'<<8 | 'o' // 'auto'
	KAEBefore           KAE = 'b'<<24 | 'e'<<16 | 'f'<<8 | 'o' // 'befo'
	KAEBeginTransaction KAE = 'b'<<24 | 'e'<<16 | 'g'<<8 | 'i' // 'begi'
	KAEBeginning        KAE = 'b'<<24 | 'g'<<16 | 'n'<<8 | 'g' // 'bgng'
	// KAEBeginsWith: The value of `operand1` begins with the value of `operand2` (for example, the string `"operand"` begins with the string `"opera"`).
	KAEBeginsWith KAE = 'b'<<24 | 'g'<<16 | 'w'<<8 | 't' // 'bgwt'
	KAEBold       KAE = 'b'<<24 | 'o'<<16 | 'l'<<8 | 'd' // 'bold'
	// KAECanInteract: # Discussion
	KAECanInteract KAE = 32
	// KAECanSwitchLayer: # Discussion
	KAECanSwitchLayer KAE = 64
	KAECaseSensEquals KAE = 'c'<<24 | 's'<<16 | 'e'<<8 | 'q' // 'cseq'
	KAECentered       KAE = 'c'<<24 | 'e'<<16 | 'n'<<8 | 't' // 'cent'
	KAEChangeView     KAE = 'v'<<24 | 'i'<<16 | 'e'<<8 | 'w' // 'view'
	KAEClone          KAE = 'c'<<24 | 'l'<<16 | 'o'<<8 | 'n' // 'clon'
	KAEClose          KAE = 'c'<<24 | 'l'<<16 | 'o'<<8 | 's' // 'clos'
	KAECommandClass   KAE = 'c'<<24 | 'm'<<16 | 'n'<<8 | 'd' // 'cmnd'
	KAECondensed      KAE = 'c'<<24 | 'o'<<16 | 'n'<<8 | 'd' // 'cond'
	// KAEContains: The value of `operand1` contains the value of `operand2 `(for example, the string `"operand"` contains the string `"era"`).
	KAEContains KAE = 'c'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'cont'
	KAECopy     KAE = 'c'<<24 | 'o'<<16 | 'p'<<8 | 'y' // 'copy'
	// KAECoreSuite: An Apple event in the Standard Suite.
	KAECoreSuite       KAE = 'c'<<24 | 'o'<<16 | 'r'<<8 | 'e' // 'core'
	KAECountElements   KAE = 'c'<<24 | 'n'<<16 | 't'<<8 | 'e' // 'cnte'
	KAECreateElement   KAE = 'c'<<24 | 'r'<<16 | 'e'<<8 | 'l' // 'crel'
	KAECreatePublisher KAE = 'c'<<24 | 'p'<<16 | 'u'<<8 | 'b' // 'cpub'
	KAECut             KAE = 'c'<<24 | 'u'<<16 | 't'<<8 | ' ' // 'cut '
	// KAEDataArray: Array items consist of data of the same size and same type, and are aligned on word boundaries.
	KAEDataArray  KAE = 0
	KAEDeactivate KAE = 'd'<<24 | 'a'<<16 | 'c'<<8 | 't' // 'dact'
	KAEDelete     KAE = 'd'<<24 | 'e'<<16 | 'l'<<8 | 'o' // 'delo'
	// KAEDescArray: Array items consist of descriptors of different descriptor types with data of variable size.
	KAEDescArray KAE = 3
	// KAEDirectCall: The source of the Apple event is a direct call that bypassed the PPC Toolbox.
	KAEDirectCall                              KAE = 1
	KAEDiskEvent                               KAE = 'd'<<24 | 'i'<<16 | 's'<<8 | 'k' // 'disk'
	KAEDoNotAutomaticallyAddAnnotationsToEvent KAE = 65536
	KAEDoObjectsExist                          KAE = 'd'<<24 | 'o'<<16 | 'e'<<8 | 'x' // 'doex'
	KAEDoScript                                KAE = 'd'<<24 | 'o'<<16 | 's'<<8 | 'c' // 'dosc'
	// KAEDontExecute: # Discussion
	KAEDontExecute KAE = 8192
	// KAEDontReconnect: # Discussion
	KAEDontReconnect KAE = 128
	// KAEDontRecord: # Discussion
	KAEDontRecord         KAE = 4096
	KAEDown               KAE = 'd'<<24 | 'o'<<16 | 'w'<<8 | 'n' // 'down'
	KAEDrag               KAE = 'd'<<24 | 'r'<<16 | 'a'<<8 | 'g' // 'drag'
	KAEDuplicateSelection KAE = 's'<<24 | 'd'<<16 | 'u'<<8 | 'p' // 'sdup'
	KAEEditGraphic        KAE = 'e'<<24 | 'd'<<16 | 'i'<<8 | 't' // 'edit'
	KAEEmptyTrash         KAE = 'e'<<24 | 'm'<<16 | 'p'<<8 | 't' // 'empt'
	KAEEnd                KAE = 'e'<<24 | 'n'<<16 | 'd'<<8 | ' ' // 'end '
	KAEEndTransaction     KAE = 'e'<<24 | 'n'<<16 | 'd'<<8 | 't' // 'endt'
	// KAEEndsWith: The value of `operand1` ends with the value of `operand2` (for example, the string `"operand"` ends with the string `"and"`).
	KAEEndsWith KAE = 'e'<<24 | 'n'<<16 | 'd'<<8 | 's' // 'ends'
	// KAEEquals: The value of `operand1` is equal to the value of `operand2`
	KAEEquals   KAE = '='<<24 | ' '<<16 | ' '<<8 | ' ' // '=   '
	KAEExpanded KAE = 'p'<<24 | 'e'<<16 | 'x'<<8 | 'p' // 'pexp'
	KAEFast     KAE = 'f'<<24 | 'a'<<16 | 's'<<8 | 't' // 'fast'
	// KAEFinderEvents: An event that the Finder accepts.
	KAEFinderEvents          KAE = 'F'<<24 | 'N'<<16 | 'D'<<8 | 'R' // 'FNDR'
	KAEFormulaProtect        KAE = 'f'<<24 | 'p'<<16 | 'r'<<8 | 'o' // 'fpro'
	KAEFullyJustified        KAE = 'f'<<24 | 'u'<<16 | 'l'<<8 | 'l' // 'full'
	KAEGetClassInfo          KAE = 'q'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'qobj'
	KAEGetData               KAE = 'g'<<24 | 'e'<<16 | 't'<<8 | 'd' // 'getd'
	KAEGetDataSize           KAE = 'd'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'dsiz'
	KAEGetEventInfo          KAE = 'g'<<24 | 't'<<16 | 'e'<<8 | 'i' // 'gtei'
	KAEGetInfoSelection      KAE = 's'<<24 | 'i'<<16 | 'n'<<8 | 'f' // 'sinf'
	KAEGetPrivilegeSelection KAE = 's'<<24 | 'p'<<16 | 'r'<<8 | 'v' // 'sprv'
	KAEGetSuiteInfo          KAE = 'g'<<24 | 't'<<16 | 's'<<8 | 'i' // 'gtsi'
	// KAEGreaterThan: The value of `operand1` is greater than the value of `operand2`.
	KAEGreaterThan KAE = '>'<<24 | ' '<<16 | ' '<<8 | ' ' // '>   '
	// KAEGreaterThanEquals: The value of `operand1` is greater than or equal to the value of `operand2`.
	KAEGreaterThanEquals KAE = '>'<<24 | '='<<16 | ' '<<8 | ' ' // '>=  '
	KAEGrow              KAE = 'g'<<24 | 'r'<<16 | 'o'<<8 | 'w' // 'grow'
	// KAEHTTPProxyHostAttr: A value of type `typeChar` or `typeUTF8Text`.
	KAEHTTPProxyHostAttr KAE = 'x'<<24 | 'h'<<16 | 't'<<8 | 'h' // 'xhth'
	// KAEHTTPProxyPortAttr: A value of type `typeSInt32`.
	KAEHTTPProxyPortAttr  KAE = 'x'<<24 | 'h'<<16 | 't'<<8 | 'p' // 'xhtp'
	KAEHandleSimpleRanges KAE = 32
	KAEHiQuality          KAE = 'h'<<24 | 'i'<<16 | 'q'<<8 | 'u' // 'hiqu'
	KAEHidden             KAE = 'h'<<24 | 'i'<<16 | 'd'<<8 | 'n' // 'hidn'
	KAEHighLevel          KAE = 'h'<<24 | 'i'<<16 | 'g'<<8 | 'h' // 'high'
	// KAEHighPriority: The Apple Event Manager posts the event at the beginning of the event queue of the server process.
	KAEHighPriority KAE = 1
	// KAEIDoMarking: The application provides marking callback functions.
	KAEIDoMarking KAE = 4
	// KAEIDoMinimum: The application does not handle whose tests or provide marking callbacks.
	KAEIDoMinimum KAE = 0
	// KAEIDoWhose: The application supports whose tests (supports key form `formWhose`).
	KAEIDoWhose       KAE = 1
	KAEISWebStarSuite KAE = 1465341885
	KAEImageGraphic   KAE = 'i'<<24 | 'm'<<16 | 'g'<<8 | 'r' // 'imgr'
	KAEInfo           KAE = 11
	KAEInternetSuite  KAE = 'g'<<24 | 'u'<<16 | 'r'<<8 | 'l' // 'gurl'
	KAEIsUniform      KAE = 'i'<<24 | 's'<<16 | 'u'<<8 | 'n' // 'isun'
	KAEItalic         KAE = 'i'<<24 | 't'<<16 | 'a'<<8 | 'l' // 'ital'
	KAEKeyClass       KAE = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'c' // 'keyc'
	// KAEKeyDescArray: Array items consist of keyword-specified descriptors with different keywords, different descriptor types, and data of variable size.
	KAEKeyDescArray  KAE = 4
	KAEKeyDown       KAE = 'k'<<24 | 'd'<<16 | 'w'<<8 | 'n' // 'kdwn'
	KAELeftJustified KAE = 'l'<<24 | 'e'<<16 | 'f'<<8 | 't' // 'left'
	KAELessThan      KAE = '<'<<24 | ' '<<16 | ' '<<8 | ' ' // '<   '
	// KAELessThanEquals: The value of `operand1` is less than or equal to the value of `operand2`.
	KAELessThanEquals KAE = '<'<<24 | '='<<16 | ' '<<8 | ' ' // '<=  '
	// KAELocalProcess: The source application is another process on the same computer as the target application.
	KAELocalProcess       KAE = 3
	KAELogOut             KAE = 'l'<<24 | 'o'<<16 | 'g'<<8 | 'o' // 'logo'
	KAELowercase          KAE = 'l'<<24 | 'o'<<16 | 'w'<<8 | 'c' // 'lowc'
	KAEMain               KAE = 0
	KAEMakeObjectsVisible KAE = 'm'<<24 | 'v'<<16 | 'i'<<8 | 's' // 'mvis'
	KAEMenuClass          KAE = 'm'<<24 | 'e'<<16 | 'n'<<8 | 'u' // 'menu'
	KAEMenuSelect         KAE = 'm'<<24 | 'h'<<16 | 'i'<<8 | 't' // 'mhit'
	KAEMiscStandards      KAE = 'm'<<24 | 'i'<<16 | 's'<<8 | 'c' // 'misc'
	KAEModifiable         KAE = 'm'<<24 | 'o'<<16 | 'd'<<8 | 'f' // 'modf'
	KAEMouseClass         KAE = 'm'<<24 | 'o'<<16 | 'u'<<8 | 's' // 'mous'
	KAEMouseDown          KAE = 'm'<<24 | 'd'<<16 | 'w'<<8 | 'n' // 'mdwn'
	KAEMouseDownInBack    KAE = 'm'<<24 | 'd'<<16 | 'b'<<8 | 'k' // 'mdbk'
	KAEMove               KAE = 'm'<<24 | 'o'<<16 | 'v'<<8 | 'e' // 'move'
	KAEMoved              KAE = 'm'<<24 | 'o'<<16 | 'v'<<8 | 'e' // 'move'
	KAENavigationKey      KAE = 'n'<<24 | 'a'<<16 | 'v'<<8 | 'e' // 'nave'
	// KAENeverInteract: # Discussion
	KAENeverInteract KAE = 16
	KAENo            KAE = 'n'<<24 | 'o'<<16 | ' '<<8 | ' ' // 'no  '
	KAENoArrow       KAE = 'a'<<24 | 'r'<<16 | 'n'<<8 | 'o' // 'arno'
	// KAENoReply: The reply preference—your application does not want a reply Apple event.
	KAENoReply       KAE = 1
	KAENonmodifiable KAE = 'n'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'nmod'
	// KAENormalPriority: The Apple Event Manager posts the event at the end of the event queue of the server process and the server processes the Apple event as soon as it has the opportunity.
	KAENormalPriority KAE = 0
	// KAENotifyRecording: # Discussion
	KAENotifyRecording KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'r' // 'recr'
	// KAENotifyStartRecording: An event that notifies an application that recording has been turned on.
	KAENotifyStartRecording KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | '1' // 'rec1'
	// KAENotifyStopRecording: An event that notifies an application that recording has been turned off.
	KAENotifyStopRecording KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | '0' // 'rec0'
	KAENullEvent           KAE = 'n'<<24 | 'u'<<16 | 'l'<<8 | 'l' // 'null'
	KAEOSAXSizeResource    KAE = 'o'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'osiz'
	KAEOpen                KAE = 'o'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'odoc'
	// KAEOpenApplication: Event that launches an application.
	KAEOpenApplication KAE = 'o'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'oapp'
	// KAEOpenContents: # Discussion
	KAEOpenContents KAE = 'o'<<24 | 'c'<<16 | 'o'<<8 | 'n' // 'ocon'
	// KAEOpenDocuments: # Discussion
	KAEOpenDocuments KAE = 'o'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'odoc'
	KAEOpenSelection KAE = 's'<<24 | 'o'<<16 | 'p'<<8 | 'e' // 'sope'
	KAEOutline       KAE = 'o'<<24 | 'u'<<16 | 't'<<8 | 'l' // 'outl'
	// KAEPackedArray: Array items consist of data of the same size and same type, and are packed without regard for word boundaries.
	KAEPackedArray  KAE = 1
	KAEPageSetup    KAE = 'p'<<24 | 'g'<<16 | 's'<<8 | 'u' // 'pgsu'
	KAEPassSubDescs KAE = 8
	KAEPaste        KAE = 'p'<<24 | 'a'<<16 | 's'<<8 | 't' // 'past'
	KAEPlain        KAE = 'p'<<24 | 'l'<<16 | 'a'<<8 | 'n' // 'plan'
	KAEPrint        KAE = 'p'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'pdoc'
	// KAEPrintDocuments: Event that provides an application with a list of documents to print.
	KAEPrintDocuments KAE = 'p'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'pdoc'
	KAEPrintSelection KAE = 's'<<24 | 'p'<<16 | 'r'<<8 | 'i' // 'spri'
	KAEPrintWindow    KAE = 'p'<<24 | 'w'<<16 | 'i'<<8 | 'n' // 'pwin'
	// KAEProcessNonReplyEvents: Allow processing of non-reply Apple events while awaiting a synchronous Apple event reply (you specified `kAEWaitReply` for the reply preference).
	KAEProcessNonReplyEvents KAE = 32768
	KAEPromise               KAE = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'm' // 'prom'
	KAEPutAwaySelection      KAE = 's'<<24 | 'p'<<16 | 'u'<<8 | 't' // 'sput'
	KAEQDAdMax               KAE = 'a'<<24 | 'd'<<16 | 'm'<<8 | 'x' // 'admx'
	KAEQDAdMin               KAE = 'a'<<24 | 'd'<<16 | 'm'<<8 | 'n' // 'admn'
	KAEQDAddOver             KAE = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'o' // 'addo'
	KAEQDAddPin              KAE = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'p' // 'addp'
	KAEQDBic                 KAE = 'b'<<24 | 'i'<<16 | 'c'<<8 | ' ' // 'bic '
	KAEQDBlend               KAE = 'b'<<24 | 'l'<<16 | 'n'<<8 | 'd' // 'blnd'
	KAEQDCopy                KAE = 'c'<<24 | 'p'<<16 | 'y'<<8 | ' ' // 'cpy '
	KAEQDNotBic              KAE = 'n'<<24 | 'b'<<16 | 'i'<<8 | 'c' // 'nbic'
	KAEQDNotCopy             KAE = 'n'<<24 | 'c'<<16 | 'p'<<8 | 'y' // 'ncpy'
	KAEQDNotOr               KAE = 'n'<<24 | 't'<<16 | 'o'<<8 | 'r' // 'ntor'
	KAEQDNotXor              KAE = 'n'<<24 | 'x'<<16 | 'o'<<8 | 'r' // 'nxor'
	KAEQDOr                  KAE = 'o'<<24 | 'r'<<16 | ' '<<8 | ' ' // 'or  '
	KAEQDSubOver             KAE = 's'<<24 | 'u'<<16 | 'b'<<8 | 'o' // 'subo'
	KAEQDSubPin              KAE = 's'<<24 | 'u'<<16 | 'b'<<8 | 'p' // 'subp'
	KAEQDSupplementalSuite   KAE = 'q'<<24 | 'd'<<16 | 's'<<8 | 'p' // 'qdsp'
	KAEQDXor                 KAE = 'x'<<24 | 'o'<<16 | 'r'<<8 | ' ' // 'xor '
	// KAEQueueReply: The reply preference—your application wants a reply Apple event.
	KAEQueueReply     KAE = 2
	KAEQuickdrawSuite KAE = 'q'<<24 | 'd'<<16 | 'r'<<8 | 'w' // 'qdrw'
	KAEQuitAll        KAE = 'q'<<24 | 'u'<<16 | 'i'<<8 | 'a' // 'quia'
	// KAEQuitApplication: Event that causes the application to quit.
	KAEQuitApplication KAE = 'q'<<24 | 'u'<<16 | 'i'<<8 | 't' // 'quit'
	KAERawKey          KAE = 'r'<<24 | 'k'<<16 | 'e'<<8 | 'y' // 'rkey'
	KAEReallyLogOut    KAE = 'r'<<24 | 'l'<<16 | 'g'<<8 | 'o' // 'rlgo'
	KAERedo            KAE = 'r'<<24 | 'e'<<16 | 'd'<<8 | 'o' // 'redo'
	KAERegular         KAE = 'r'<<24 | 'e'<<16 | 'g'<<8 | 'l' // 'regl'
	// KAERemoteProcess: The source application is a process on a remote computer on the network.
	KAERemoteProcess KAE = 4
	// KAEReopenApplication: Event that reopens an application.
	KAEReopenApplication  KAE = 'r'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'rapp'
	KAEReplace            KAE = 'r'<<24 | 'p'<<16 | 'l'<<8 | 'c' // 'rplc'
	KAERequiredSuite      KAE = 'r'<<24 | 'e'<<16 | 'q'<<8 | 'd' // 'reqd'
	KAEResized            KAE = 'r'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'rsiz'
	KAEResolveNestedLists KAE = 16
	KAERestart            KAE = 'r'<<24 | 'e'<<16 | 's'<<8 | 't' // 'rest'
	KAEResume             KAE = 'r'<<24 | 's'<<16 | 'm'<<8 | 'e' // 'rsme'
	KAERevealSelection    KAE = 's'<<24 | 'r'<<16 | 'e'<<8 | 'v' // 'srev'
	KAERevert             KAE = 'r'<<24 | 'v'<<16 | 'r'<<8 | 't' // 'rvrt'
	KAERightJustified     KAE = 'r'<<24 | 'g'<<16 | 'h'<<8 | 't' // 'rght'
	// KAESameProcess: The source of the Apple event is the same application that received the event (the target application and the source application are the same).
	KAESameProcess           KAE = 2
	KAESave                  KAE = 's'<<24 | 'a'<<16 | 'v'<<8 | 'e' // 'save'
	KAEScrapEvent            KAE = 's'<<24 | 'c'<<16 | 'r'<<8 | 'p' // 'scrp'
	KAEScriptingSizeResource KAE = 's'<<24 | 'c'<<16 | 's'<<8 | 'z' // 'scsz'
	KAESelect                KAE = 's'<<24 | 'l'<<16 | 'c'<<8 | 't' // 'slct'
	KAESetData               KAE = 's'<<24 | 'e'<<16 | 't'<<8 | 'd' // 'setd'
	KAESetPosition           KAE = 'p'<<24 | 'o'<<16 | 's'<<8 | 'n' // 'posn'
	KAEShadow                KAE = 's'<<24 | 'h'<<16 | 'a'<<8 | 'd' // 'shad'
	KAESharing               KAE = 13
	KAEShowClipboard         KAE = 's'<<24 | 'h'<<16 | 'c'<<8 | 'l' // 'shcl'
	// KAEShowPreferences: # Discussion
	KAEShowPreferences        KAE = 'p'<<24 | 'r'<<16 | 'e'<<8 | 'f' // 'pref'
	KAEShowRestartDialog      KAE = 'r'<<24 | 'r'<<16 | 's'<<8 | 't' // 'rrst'
	KAEShowShutdownDialog     KAE = 'r'<<24 | 's'<<16 | 'd'<<8 | 'n' // 'rsdn'
	KAEShutDown               KAE = 's'<<24 | 'h'<<16 | 'u'<<8 | 't' // 'shut'
	KAESleep                  KAE = 's'<<24 | 'l'<<16 | 'e'<<8 | 'p' // 'slep'
	KAESmallCaps              KAE = 's'<<24 | 'm'<<16 | 'c'<<8 | 'p' // 'smcp'
	KAESocks4Protocol         KAE = 4
	KAESocks5Protocol         KAE = 5
	KAESocksHostAttr          KAE = 'x'<<24 | 's'<<16 | 'h'<<8 | 's' // 'xshs'
	KAESocksPasswordAttr      KAE = 'x'<<24 | 's'<<16 | 'h'<<8 | 'w' // 'xshw'
	KAESocksPortAttr          KAE = 'x'<<24 | 's'<<16 | 'h'<<8 | 'p' // 'xshp'
	KAESocksProxyAttr         KAE = 'x'<<24 | 's'<<16 | 'o'<<8 | 'k' // 'xsok'
	KAESocksUserAttr          KAE = 'x'<<24 | 's'<<16 | 'h'<<8 | 'u' // 'xshu'
	KAESpecialClassProperties KAE = 'c'<<24 | '@'<<16 | '#'<<8 | '!' // 'c@#!'
	// KAEStartRecording: # Discussion
	KAEStartRecording KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'a' // 'reca'
	// KAEStopRecording: # Discussion
	KAEStopRecording         KAE = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'c' // 'recc'
	KAEStoppedMoving         KAE = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
	KAEStrikethrough         KAE = 's'<<24 | 't'<<16 | 'r'<<8 | 'k' // 'strk'
	KAESubscript             KAE = 's'<<24 | 'b'<<16 | 's'<<8 | 'c' // 'sbsc'
	KAESuperscript           KAE = 's'<<24 | 'p'<<16 | 's'<<8 | 'c' // 'spsc'
	KAESuspend               KAE = 's'<<24 | 'u'<<16 | 's'<<8 | 'p' // 'susp'
	KAETableSuite            KAE = 't'<<24 | 'b'<<16 | 'l'<<8 | 's' // 'tbls'
	KAETerminologyExtension  KAE = 'a'<<24 | 'e'<<16 | 't'<<8 | 'e' // 'aete'
	KAETextSuite             KAE = 'T'<<24 | 'E'<<16 | 'X'<<8 | 'T' // 'TEXT'
	KAETransactionTerminated KAE = 't'<<24 | 't'<<16 | 'r'<<8 | 'm' // 'ttrm'
	KAEUnderline             KAE = 'u'<<24 | 'n'<<16 | 'd'<<8 | 'l' // 'undl'
	KAEUndo                  KAE = 'u'<<24 | 'n'<<16 | 'd'<<8 | 'o' // 'undo'
	// KAEUnknownSource: The source of the Apple event is unknown.
	KAEUnknownSource KAE = 0
	KAEUp            KAE = 'u'<<24 | 'p'<<16 | ' '<<8 | ' ' // 'up  '
	KAEUpdate        KAE = 'u'<<24 | 'p'<<16 | 'd'<<8 | 't' // 'updt'
	// KAEUseHTTPProxyAttr: A value of type `typeBoolean`.
	KAEUseHTTPProxyAttr     KAE = 'x'<<24 | 'u'<<16 | 'p'<<8 | 'r' // 'xupr'
	KAEUseRelativeIterators KAE = 64
	KAEUseSocksAttr         KAE = 'x'<<24 | 's'<<16 | 'c'<<8 | 's' // 'xscs'
	KAEUserTerminology      KAE = 'a'<<24 | 'e'<<16 | 'u'<<8 | 't' // 'aeut'
	KAEVirtualKey           KAE = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'c' // 'keyc'
	// KAEWaitReply: # Discussion
	KAEWaitReply   KAE = 3
	KAEWakeUpEvent KAE = 'w'<<24 | 'a'<<16 | 'k'<<8 | 'e' // 'wake'
	// KAEWantReceipt: # Discussion
	KAEWantReceipt     KAE = 512
	KAEWholeWordEquals KAE = 'w'<<24 | 'w'<<16 | 'e'<<8 | 'q' // 'wweq'
	KAEWindowClass     KAE = 'w'<<24 | 'i'<<16 | 'n'<<8 | 'd' // 'wind'
	KAEYes             KAE = 'y'<<24 | 'e'<<16 | 's'<<8 | ' ' // 'yes '
	KAEZoomValue       KAE = 'z'<<24 | 'o'<<16 | 'o'<<8 | 'm' // 'zoom'
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
	case KAEZoomValue:
		return "KAEZoomValue"
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

type KAEApplicationActivation uint

const (
	KAEApplicationActivationExpected KAEApplicationActivation = 'a'<<24 | 'a'<<16 | 'p'<<8 | 'd' // 'aapd'
)

func (e KAEApplicationActivation) String() string {
	switch e {
	case KAEApplicationActivationExpected:
		return "KAEApplicationActivationExpected"
	default:
		return fmt.Sprintf("KAEApplicationActivation(%d)", e)
	}
}

type KAEDebug uint

const (
	KAEDebugPOSTHeader  KAEDebug = 1
	KAEDebugReplyHeader KAEDebug = 2
	KAEDebugXMLDebugAll KAEDebug = 0
	KAEDebugXMLRequest  KAEDebug = 4
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
	KAEDescListFactorNone        KAEDescListFactor = 0
	KAEDescListFactorType        KAEDescListFactor = 4
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

type KAEDoNotPromptForUser uint

const (
	KAEDoNotPromptForUserConsent KAEDoNotPromptForUser = 131072
)

func (e KAEDoNotPromptForUser) String() string {
	switch e {
	case KAEDoNotPromptForUserConsent:
		return "KAEDoNotPromptForUserConsent"
	default:
		return fmt.Sprintf("KAEDoNotPromptForUser(%d)", e)
	}
}

type KAEHandle uint

const (
	// KAEHandleArray: Array items consist of handles to data of the same type and possibly variable size.
	KAEHandleArray KAEHandle = 2
)

func (e KAEHandle) String() string {
	switch e {
	case KAEHandleArray:
		return "KAEHandleArray"
	default:
		return fmt.Sprintf("KAEHandle(%d)", e)
	}
}

type KAEIS uint

const (
	KAEISAction         KAEIS = 'K'<<24 | 'a'<<16 | 'c'<<8 | 't' // 'Kact'
	KAEISActionPath     KAEIS = 'K'<<24 | 'a'<<16 | 'p'<<8 | 't' // 'Kapt'
	KAEISClientAddress  KAEIS = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'addr'
	KAEISClientIP       KAEIS = 'K'<<24 | 'c'<<16 | 'i'<<8 | 'p' // 'Kcip'
	KAEISContentType    KAEIS = 'c'<<24 | 't'<<16 | 'y'<<8 | 'p' // 'ctyp'
	KAEISFromUser       KAEIS = 'f'<<24 | 'r'<<16 | 'm'<<8 | 'u' // 'frmu'
	KAEISFullRequest    KAEIS = 'K'<<24 | 'f'<<16 | 'r'<<8 | 'q' // 'Kfrq'
	KAEISHTTPSearchArgs KAEIS = 'k'<<24 | 'f'<<16 | 'o'<<8 | 'r' // 'kfor'
	KAEISMethod         KAEIS = 'm'<<24 | 'e'<<16 | 't'<<8 | 'h' // 'meth'
	KAEISPassword       KAEIS = 'p'<<24 | 'a'<<16 | 's'<<8 | 's' // 'pass'
	KAEISPostArgs       KAEIS = 'p'<<24 | 'o'<<16 | 's'<<8 | 't' // 'post'
	KAEISReferrer       KAEIS = 'r'<<24 | 'e'<<16 | 'f'<<8 | 'r' // 'refr'
	KAEISScriptName     KAEIS = 's'<<24 | 'c'<<16 | 'n'<<8 | 'm' // 'scnm'
	KAEISServerName     KAEIS = 's'<<24 | 'v'<<16 | 'n'<<8 | 'm' // 'svnm'
	KAEISServerPort     KAEIS = 's'<<24 | 'v'<<16 | 'p'<<8 | 't' // 'svpt'
	KAEISUserAgent      KAEIS = 'A'<<24 | 'g'<<16 | 'n'<<8 | 't' // 'Agnt'
	KAEISUserName       KAEIS = 'u'<<24 | 's'<<16 | 'e'<<8 | 'r' // 'user'
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
	KAEISHandleCGI   KAEISGetURL = 's'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'sdoc'
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
	KAEQuitReason        KAEQuit = 'w'<<24 | 'h'<<16 | 'y'<<8 | '?' // 'why?'
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
	KAEUTApostrophe             KAEU = 3
	KAEUTChangesState           KAEU = 12
	KAEUTDirectParamIsReference KAEU = 9
	KAEUTEnumListIsExclusive    KAEU = 10
	KAEUTEnumerated             KAEU = 13
	KAEUTEnumsAreTypes          KAEU = 11
	KAEUTFeminine               KAEU = 2
	KAEUTHasReturningParam      KAEU = 31
	KAEUTMasculine              KAEU = 1
	KAEUTNotDirectParamIsTarget KAEU = 8
	KAEUTOptional               KAEU = 15
	KAEUTParamIsReference       KAEU = 9
	KAEUTParamIsTarget          KAEU = 8
	KAEUTPlural                 KAEU = 0
	KAEUTPropertyIsReference    KAEU = 9
	KAEUTReadWrite              KAEU = 12
	KAEUTReplyIsReference       KAEU = 9
	KAEUTTightBindingFunction   KAEU = 12
	KAEUTlistOfItems            KAEU = 14
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

type KAEZoom uint

const (
	KAEZoomIn  KAEZoom = 7
	KAEZoomOut KAEZoom = 8
)

func (e KAEZoom) String() string {
	switch e {
	case KAEZoomIn:
		return "KAEZoomIn"
	case KAEZoomOut:
		return "KAEZoomOut"
	default:
		return fmt.Sprintf("KAEZoom(%d)", e)
	}
}

type KAFPExtendedFlagsAlternateAddress uint

const (
	KAFPExtendedFlagsAlternateAddressMask KAFPExtendedFlagsAlternateAddress = 1
)

func (e KAFPExtendedFlagsAlternateAddress) String() string {
	switch e {
	case KAFPExtendedFlagsAlternateAddressMask:
		return "KAFPExtendedFlagsAlternateAddressMask"
	default:
		return fmt.Sprintf("KAFPExtendedFlagsAlternateAddress(%d)", e)
	}
}

type KAFPTagLength uint

const (
	KAFPTagLengthDDP    KAFPTagLength = 6
	KAFPTagLengthIP     KAFPTagLength = 6
	KAFPTagLengthIPPort KAFPTagLength = 8
)

func (e KAFPTagLength) String() string {
	switch e {
	case KAFPTagLengthDDP:
		return "KAFPTagLengthDDP"
	case KAFPTagLengthIPPort:
		return "KAFPTagLengthIPPort"
	default:
		return fmt.Sprintf("KAFPTagLength(%d)", e)
	}
}

type KAFPTagType uint

const (
	KAFPTagTypeDDP    KAFPTagType = 3
	KAFPTagTypeDNS    KAFPTagType = 4
	KAFPTagTypeIP     KAFPTagType = 1
	KAFPTagTypeIPPort KAFPTagType = 2
)

func (e KAFPTagType) String() string {
	switch e {
	case KAFPTagTypeDDP:
		return "KAFPTagTypeDDP"
	case KAFPTagTypeDNS:
		return "KAFPTagTypeDNS"
	case KAFPTagTypeIP:
		return "KAFPTagTypeIP"
	case KAFPTagTypeIPPort:
		return "KAFPTagTypeIPPort"
	default:
		return fmt.Sprintf("KAFPTagType(%d)", e)
	}
}

type KALM int

const (
	KALMDeferSwitchErr         KALM = -30043
	KALMDuplicateModuleErr     KALM = -30045
	KALMGroupNotFoundErr       KALM = -30048
	KALMInstallationErr        KALM = -30044
	KALMInternalErr            KALM = -30049
	KALMModuleCommunicationErr KALM = -30046
	KALMNoSuchModuleErr        KALM = -30047
	KALMRebootFlagsLevelErr    KALM = -30042
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

type KALMLocationNotFound int

const (
	KALMLocationNotFoundErr KALMLocationNotFound = -30048
)

func (e KALMLocationNotFound) String() string {
	switch e {
	case KALMLocationNotFoundErr:
		return "KALMLocationNotFoundErr"
	default:
		return fmt.Sprintf("KALMLocationNotFound(%d)", e)
	}
}

type KARM uint

const (
	// Deprecated.
	KARMMountVol KARM = 1
	// Deprecated.
	KARMMultVols KARM = 8
	// Deprecated.
	KARMNoUI KARM = 2
	// Deprecated.
	KARMSearch KARM = 256
	// Deprecated.
	KARMSearchMore KARM = 512
	// Deprecated.
	KARMSearchRelFirst KARM = 1024
	// Deprecated.
	KARMTryFileIDFirst KARM = 2048
)

func (e KARM) String() string {
	switch e {
	case KARMMountVol:
		return "KARMMountVol"
	case KARMMultVols:
		return "KARMMultVols"
	case KARMNoUI:
		return "KARMNoUI"
	case KARMSearch:
		return "KARMSearch"
	case KARMSearchMore:
		return "KARMSearchMore"
	case KARMSearchRelFirst:
		return "KARMSearchRelFirst"
	case KARMTryFileIDFirst:
		return "KARMTryFileIDFirst"
	default:
		return fmt.Sprintf("KARM(%d)", e)
	}
}

type KATSU int

const (
	KATSUBadStreamErr                 KATSU = -8902
	KATSUBusyObjectErr                KATSU = -8809
	KATSUCoordinateOverflowErr        KATSU = -8807
	KATSUFontsMatched                 KATSU = -8793
	KATSUFontsNotMatched              KATSU = -8794
	KATSUInvalidAttributeSizeErr      KATSU = -8798
	KATSUInvalidAttributeTagErr       KATSU = -8799
	KATSUInvalidAttributeValueErr     KATSU = -8797
	KATSUInvalidCacheErr              KATSU = -8800
	KATSUInvalidCallInsideCallbackErr KATSU = -8904
	KATSUInvalidFontErr               KATSU = -8796
	KATSUInvalidFontFallbacksErr      KATSU = -8900
	KATSUInvalidStyleErr              KATSU = -8791
	KATSUInvalidTextLayoutErr         KATSU = -8790
	KATSUInvalidTextRangeErr          KATSU = -8792
	KATSULastErr                      KATSU = -8959
	KATSULineBreakInWord              KATSU = -8808
	KATSULowLevelErr                  KATSU = -8804
	KATSUNoCorrespondingFontErr       KATSU = -8795
	KATSUNoFontCmapAvailableErr       KATSU = -8805
	KATSUNoFontNameErr                KATSU = -8905
	KATSUNoFontScalerAvailableErr     KATSU = -8806
	KATSUNoStyleRunsAssignedErr       KATSU = -8802
	KATSUNotSetErr                    KATSU = -8801
	KATSUOutputBufferTooSmallErr      KATSU = -8903
	KATSUQuickDrawTextErr             KATSU = -8803
	KATSUUnsupportedStreamFormatErr   KATSU = -8901
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
	KAlertNoteIcon    KAlert = 'n'<<24 | 'o'<<16 | 't'<<8 | 'e' // 'note'
	KAlertStopIcon    KAlert = 's'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'stop'
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

type KAny uint

const (
	KAnyAuthType KAny = 0
	KAnyPort     KAny = 0
	KAnyProtocol KAny = 0
)

func (e KAny) String() string {
	switch e {
	case KAnyAuthType:
		return "KAnyAuthType"
	default:
		return fmt.Sprintf("KAny(%d)", e)
	}
}

type KAnyComponent uint

const (
	// Deprecated.
	KAnyComponentFlagsMask KAnyComponent = 0
	// Deprecated.
	KAnyComponentManufacturer KAnyComponent = 0
	// Deprecated.
	KAnyComponentSubType KAnyComponent = 0
	// Deprecated.
	KAnyComponentType KAnyComponent = 0
)

func (e KAnyComponent) String() string {
	switch e {
	case KAnyComponentFlagsMask:
		return "KAnyComponentFlagsMask"
	default:
		return fmt.Sprintf("KAnyComponent(%d)", e)
	}
}

type KAppearanceFolderIcon uint

const (
	KAppearanceFolderIconValue         KAppearanceFolderIcon = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'r' // 'appr'
	KAppleExtrasFolderIcon             KAppearanceFolderIcon = 1634040004
	KAppleMenuFolderIcon               KAppearanceFolderIcon = 'a'<<24 | 'm'<<16 | 'n'<<8 | 'u' // 'amnu'
	KApplicationSupportFolderIcon      KAppearanceFolderIcon = 'a'<<24 | 's'<<16 | 'u'<<8 | 'p' // 'asup'
	KApplicationsFolderIcon            KAppearanceFolderIcon = 'a'<<24 | 'p'<<16 | 'p'<<8 | 's' // 'apps'
	KAssistantsFolderIcon              KAppearanceFolderIcon = 1634956484
	KColorSyncFolderIcon               KAppearanceFolderIcon = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'f' // 'prof'
	KContextualMenuItemsFolderIcon     KAppearanceFolderIcon = 'c'<<24 | 'm'<<16 | 'n'<<8 | 'u' // 'cmnu'
	KControlPanelDisabledFolderIcon    KAppearanceFolderIcon = 'c'<<24 | 't'<<16 | 'r'<<8 | 'D' // 'ctrD'
	KControlPanelFolderIcon            KAppearanceFolderIcon = 'c'<<24 | 't'<<16 | 'r'<<8 | 'l' // 'ctrl'
	KControlStripModulesFolderIcon     KAppearanceFolderIcon = 1935963844
	KDocumentsFolderIcon               KAppearanceFolderIcon = 'd'<<24 | 'o'<<16 | 'c'<<8 | 's' // 'docs'
	KExtensionsDisabledFolderIcon      KAppearanceFolderIcon = 'e'<<24 | 'x'<<16 | 't'<<8 | 'D' // 'extD'
	KExtensionsFolderIcon              KAppearanceFolderIcon = 'e'<<24 | 'x'<<16 | 't'<<8 | 'n' // 'extn'
	KFavoritesFolderIcon               KAppearanceFolderIcon = 'f'<<24 | 'a'<<16 | 'v'<<8 | 's' // 'favs'
	KFontsFolderIcon                   KAppearanceFolderIcon = 'f'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'font'
	KHelpFolderIcon                    KAppearanceFolderIcon = 0
	KInternetFolderIcon                KAppearanceFolderIcon = 1768846532
	KInternetPlugInFolderIcon          KAppearanceFolderIcon = 0
	KInternetSearchSitesFolderIcon     KAppearanceFolderIcon = 'i'<<24 | 's'<<16 | 's'<<8 | 'f' // 'issf'
	KLocalesFolderIcon                 KAppearanceFolderIcon = 0
	KMacOSReadMeFolderIcon             KAppearanceFolderIcon = 1836020420
	KPreferencesFolderIcon             KAppearanceFolderIcon = 1886545604
	KPrintMonitorFolderIcon            KAppearanceFolderIcon = 'p'<<24 | 'r'<<16 | 'n'<<8 | 't' // 'prnt'
	KPrinterDescriptionFolderIcon      KAppearanceFolderIcon = 'p'<<24 | 'p'<<16 | 'd'<<8 | 'f' // 'ppdf'
	KPrinterDriverFolderIcon           KAppearanceFolderIcon = 0
	KPublicFolderIcon                  KAppearanceFolderIcon = 'p'<<24 | 'u'<<16 | 'b'<<8 | 'f' // 'pubf'
	KRecentApplicationsFolderIcon      KAppearanceFolderIcon = 'r'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'rapp'
	KRecentDocumentsFolderIcon         KAppearanceFolderIcon = 'r'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'rdoc'
	KRecentServersFolderIcon           KAppearanceFolderIcon = 'r'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'rsrv'
	KScriptingAdditionsFolderIcon      KAppearanceFolderIcon = 0
	KScriptsFolderIcon                 KAppearanceFolderIcon = 1935897284
	KSharedLibrariesFolderIcon         KAppearanceFolderIcon = 0
	KShutdownItemsDisabledFolderIcon   KAppearanceFolderIcon = 's'<<24 | 'h'<<16 | 'd'<<8 | 'D' // 'shdD'
	KShutdownItemsFolderIcon           KAppearanceFolderIcon = 's'<<24 | 'h'<<16 | 'd'<<8 | 'f' // 'shdf'
	KSpeakableItemsFolder              KAppearanceFolderIcon = 's'<<24 | 'p'<<16 | 'k'<<8 | 'i' // 'spki'
	KStartupItemsDisabledFolderIcon    KAppearanceFolderIcon = 's'<<24 | 't'<<16 | 'r'<<8 | 'D' // 'strD'
	KStartupItemsFolderIcon            KAppearanceFolderIcon = 's'<<24 | 't'<<16 | 'r'<<8 | 't' // 'strt'
	KSystemExtensionDisabledFolderIcon KAppearanceFolderIcon = 'm'<<24 | 'a'<<16 | 'c'<<8 | 'D' // 'macD'
	KSystemFolderIcon                  KAppearanceFolderIcon = 'm'<<24 | 'a'<<16 | 'c'<<8 | 's' // 'macs'
	KTextEncodingsFolderIcon           KAppearanceFolderIcon = 0
	KUsersFolderIcon                   KAppearanceFolderIcon = 1970500292
	KUtilitiesFolderIcon               KAppearanceFolderIcon = 1970563524
	KVoicesFolderIcon                  KAppearanceFolderIcon = 'f'<<24 | 'v'<<16 | 'o'<<8 | 'c' // 'fvoc'
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
	KAppleLogoIconValue             KAppleLogoIcon = 'c'<<24 | 'a'<<16 | 'p'<<8 | 'l' // 'capl'
	KAppleMenuIcon                  KAppleLogoIcon = 's'<<24 | 'a'<<16 | 'p'<<8 | 'l' // 'sapl'
	KBackwardArrowIcon              KAppleLogoIcon = 'b'<<24 | 'a'<<16 | 'r'<<8 | 'o' // 'baro'
	KBurningIcon                    KAppleLogoIcon = 'b'<<24 | 'u'<<16 | 'r'<<8 | 'n' // 'burn'
	KConnectToIcon                  KAppleLogoIcon = 'c'<<24 | 'n'<<16 | 'c'<<8 | 't' // 'cnct'
	KDeleteAliasIcon                KAppleLogoIcon = 'd'<<24 | 'a'<<16 | 'l'<<8 | 'i' // 'dali'
	KEjectMediaIcon                 KAppleLogoIcon = 'e'<<24 | 'j'<<16 | 'e'<<8 | 'c' // 'ejec'
	KFavoriteItemsIcon              KAppleLogoIcon = 'f'<<24 | 'a'<<16 | 'v'<<8 | 'r' // 'favr'
	KForwardArrowIcon               KAppleLogoIcon = 'f'<<24 | 'a'<<16 | 'r'<<8 | 'o' // 'faro'
	KGenericWindowIcon              KAppleLogoIcon = 'g'<<24 | 'w'<<16 | 'i'<<8 | 'n' // 'gwin'
	KGridIcon                       KAppleLogoIcon = 'g'<<24 | 'r'<<16 | 'i'<<8 | 'd' // 'grid'
	KHelpIcon                       KAppleLogoIcon = 'h'<<24 | 'e'<<16 | 'l'<<8 | 'p' // 'help'
	KKeepArrangedIcon               KAppleLogoIcon = 'a'<<24 | 'r'<<16 | 'n'<<8 | 'g' // 'arng'
	KLockedIcon                     KAppleLogoIcon = 'l'<<24 | 'o'<<16 | 'c'<<8 | 'k' // 'lock'
	KNoFilesIcon                    KAppleLogoIcon = 'n'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'nfil'
	KNoFolderIcon                   KAppleLogoIcon = 'n'<<24 | 'f'<<16 | 'l'<<8 | 'd' // 'nfld'
	KNoWriteIcon                    KAppleLogoIcon = 'n'<<24 | 'w'<<16 | 'r'<<8 | 't' // 'nwrt'
	KProtectedApplicationFolderIcon KAppleLogoIcon = 'p'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'papp'
	KProtectedSystemFolderIcon      KAppleLogoIcon = 'p'<<24 | 's'<<16 | 'y'<<8 | 's' // 'psys'
	KQuestionMarkIcon               KAppleLogoIcon = 'q'<<24 | 'u'<<16 | 'e'<<8 | 's' // 'ques'
	KRecentItemsIcon                KAppleLogoIcon = 'r'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'rcnt'
	KRightContainerArrowIcon        KAppleLogoIcon = 'r'<<24 | 'c'<<16 | 'a'<<8 | 'r' // 'rcar'
	KShortcutIcon                   KAppleLogoIcon = 's'<<24 | 'h'<<16 | 'r'<<8 | 't' // 'shrt'
	KSortAscendingIcon              KAppleLogoIcon = 'a'<<24 | 's'<<16 | 'n'<<8 | 'd' // 'asnd'
	KSortDescendingIcon             KAppleLogoIcon = 'd'<<24 | 's'<<16 | 'n'<<8 | 'd' // 'dsnd'
	KUnlockedIcon                   KAppleLogoIcon = 'u'<<24 | 'l'<<16 | 'c'<<8 | 'k' // 'ulck'
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

type KAppleManufacturer uint

const (
	// Deprecated.
	KAppleManufacturerValue KAppleManufacturer = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'l' // 'appl'
	// Deprecated.
	KComponentAliasResourceType KAppleManufacturer = 't'<<24 | 'h'<<16 | 'g'<<8 | 'a' // 'thga'
	// Deprecated.
	KComponentResourceType KAppleManufacturer = 't'<<24 | 'h'<<16 | 'n'<<8 | 'g' // 'thng'
)

func (e KAppleManufacturer) String() string {
	switch e {
	case KAppleManufacturerValue:
		return "KAppleManufacturerValue"
	case KComponentAliasResourceType:
		return "KComponentAliasResourceType"
	case KComponentResourceType:
		return "KComponentResourceType"
	default:
		return fmt.Sprintf("KAppleManufacturer(%d)", e)
	}
}

type KAppleScriptBadgeIcon uint

const (
	KAlertCautionBadgeIcon     KAppleScriptBadgeIcon = 'c'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'cbdg'
	KAliasBadgeIcon            KAppleScriptBadgeIcon = 'a'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'abdg'
	KAppleScriptBadgeIconValue KAppleScriptBadgeIcon = 's'<<24 | 'c'<<16 | 'r'<<8 | 'p' // 'scrp'
	KLockedBadgeIcon           KAppleScriptBadgeIcon = 'l'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'lbdg'
	KMountedBadgeIcon          KAppleScriptBadgeIcon = 'm'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'mbdg'
	KSharedBadgeIcon           KAppleScriptBadgeIcon = 's'<<24 | 'b'<<16 | 'd'<<8 | 'g' // 'sbdg'
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
	KAFPServerIcon      KAppleTalkIcon = 'a'<<24 | 'f'<<16 | 'p'<<8 | 's' // 'afps'
	KAppleTalkIconValue KAppleTalkIcon = 'a'<<24 | 't'<<16 | 'l'<<8 | 'k' // 'atlk'
	KAppleTalkZoneIcon  KAppleTalkIcon = 'a'<<24 | 't'<<16 | 'z'<<8 | 'n' // 'atzn'
	KFTPServerIcon      KAppleTalkIcon = 'f'<<24 | 't'<<16 | 'p'<<8 | 's' // 'ftps'
	KGenericNetworkIcon KAppleTalkIcon = 'g'<<24 | 'n'<<16 | 'e'<<8 | 't' // 'gnet'
	KHTTPServerIcon     KAppleTalkIcon = 'h'<<24 | 't'<<16 | 'p'<<8 | 's' // 'htps'
	KIPFileServerIcon   KAppleTalkIcon = 'i'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'isrv'
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

type KAsync uint

const (
	KAsyncEjectComplete     KAsync = 6
	KAsyncEjectInProgress   KAsync = 5
	KAsyncMountComplete     KAsync = 2
	KAsyncMountInProgress   KAsync = 1
	KAsyncUnmountComplete   KAsync = 4
	KAsyncUnmountInProgress KAsync = 3
)

func (e KAsync) String() string {
	switch e {
	case KAsyncEjectComplete:
		return "KAsyncEjectComplete"
	case KAsyncEjectInProgress:
		return "KAsyncEjectInProgress"
	case KAsyncMountComplete:
		return "KAsyncMountComplete"
	case KAsyncMountInProgress:
		return "KAsyncMountInProgress"
	case KAsyncUnmountComplete:
		return "KAsyncUnmountComplete"
	case KAsyncUnmountInProgress:
		return "KAsyncUnmountInProgress"
	default:
		return fmt.Sprintf("KAsync(%d)", e)
	}
}

type KAutoGenerateReturnID int

const (
	// KAnyTransactionID: # Discussion
	KAnyTransactionID KAutoGenerateReturnID = 0
	// KAutoGenerateReturnIDValue: If you pass this value for the `returnID` parameter of the AECreateAppleEvent function, the Apple Event Manager assigns to the created Apple event a return ID that is unique to the current session.
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
	K16BitCardErr            KBadAdapterErr = -9084
	KAlreadySavedStateErr    KBadAdapterErr = -9105
	KAttemptDupCardEntryErr  KBadAdapterErr = -9106
	KBadAdapterErrValue      KBadAdapterErr = -9050
	KBadArgLengthErr         KBadAdapterErr = -9063
	KBadArgsErr              KBadAdapterErr = -9064
	KBadAttributeErr         KBadAdapterErr = -9051
	KBadBaseErr              KBadAdapterErr = -9052
	KBadCISErr               KBadAdapterErr = -9066
	KBadCustomIFIDErr        KBadAdapterErr = -9093
	KBadDeviceErr            KBadAdapterErr = -9083
	KBadEDCErr               KBadAdapterErr = -9053
	KBadHandleErr            KBadAdapterErr = -9065
	KBadIRQErr               KBadAdapterErr = -9054
	KBadLinkErr              KBadAdapterErr = -9082
	KBadOffsetErr            KBadAdapterErr = -9055
	KBadPageErr              KBadAdapterErr = -9056
	KBadSizeErr              KBadAdapterErr = -9057
	KBadSocketErr            KBadAdapterErr = -9058
	KBadSpeedErr             KBadAdapterErr = -9067
	KBadTupleDataErr         KBadAdapterErr = -9092
	KBadTypeErr              KBadAdapterErr = -9059
	KBadVccErr               KBadAdapterErr = -9060
	KBadVppErr               KBadAdapterErr = -9061
	KBadWindowErr            KBadAdapterErr = -9062
	KBusyErr                 KBadAdapterErr = -9074
	KCantConfigureCardErr    KBadAdapterErr = -9087
	KCardBusCardErr          KBadAdapterErr = -9085
	KCardPowerOffErr         KBadAdapterErr = -9107
	KClientRequestDenied     KBadAdapterErr = -9102
	KConfigurationLockedErr  KBadAdapterErr = -9076
	KGeneralFailureErr       KBadAdapterErr = -9070
	KInUseErr                KBadAdapterErr = -9077
	KInvalidCSClientErr      KBadAdapterErr = -9091
	KInvalidDeviceNumber     KBadAdapterErr = -9089
	KInvalidRegEntryErr      KBadAdapterErr = -9081
	KNoCardBusCISErr         KBadAdapterErr = -9109
	KNoCardEnablersFoundErr  KBadAdapterErr = -9099
	KNoCardErr               KBadAdapterErr = -9071
	KNoCardSevicesSocketsErr KBadAdapterErr = -9080
	KNoClientTableErr        KBadAdapterErr = -9097
	KNoCompatibleNameErr     KBadAdapterErr = -9101
	KNoEnablerForCardErr     KBadAdapterErr = -9100
	KNoIOWindowRequestedErr  KBadAdapterErr = -9094
	KNoMoreInterruptSlotsErr KBadAdapterErr = -9096
	KNoMoreItemsErr          KBadAdapterErr = -9078
	KNoMoreTimerClientsErr   KBadAdapterErr = -9095
	KNotReadyErr             KBadAdapterErr = -9103
	KNotZVCapableErr         KBadAdapterErr = -9108
	KOutOfResourceErr        KBadAdapterErr = -9079
	KPassCallToChainErr      KBadAdapterErr = -9086
	KPostCardEventErr        KBadAdapterErr = -9088
	KReadFailureErr          KBadAdapterErr = -9068
	KTooManyIOWindowsErr     KBadAdapterErr = -9104
	KUnsupportedCardErr      KBadAdapterErr = -9098
	KUnsupportedFunctionErr  KBadAdapterErr = -9072
	KUnsupportedModeErr      KBadAdapterErr = -9073
	KUnsupportedVsErr        KBadAdapterErr = -9090
	KWriteFailureErr         KBadAdapterErr = -9069
	KWriteProtectedErr       KBadAdapterErr = -9075
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

type KBig5 uint

const (
	// KBig5_BasicVariant: The basic encoding variant.
	KBig5_BasicVariant KBig5 = 0
	KBig5_DOSVariant   KBig5 = 3
	// KBig5_ETenVariant: Adds kana, Cyrillic, radicals, and so forth with high-bytes C6-C8, F9.
	KBig5_ETenVariant KBig5 = 2
	// KBig5_StandardVariant: The standard variant; 0xC6A1-0xC7FC: kana, Cyrillic, enclosed numerics.
	KBig5_StandardVariant KBig5 = 1
)

func (e KBig5) String() string {
	switch e {
	case KBig5_BasicVariant:
		return "KBig5_BasicVariant"
	case KBig5_DOSVariant:
		return "KBig5_DOSVariant"
	case KBig5_ETenVariant:
		return "KBig5_ETenVariant"
	case KBig5_StandardVariant:
		return "KBig5_StandardVariant"
	default:
		return fmt.Sprintf("KBig5(%d)", e)
	}
}

type KBlessedBusError int

const (
	// Deprecated.
	KBlessedBusErrorBait KBlessedBusError = 1760651505
)

func (e KBlessedBusError) String() string {
	switch e {
	case KBlessedBusErrorBait:
		return "KBlessedBusErrorBait"
	default:
		return fmt.Sprintf("KBlessedBusError(%d)", e)
	}
}

type KBlessedFolder uint

const (
	KBlessedFolderValue KBlessedFolder = 'b'<<24 | 'l'<<16 | 's'<<8 | 'f' // 'blsf'
	KRootFolder         KBlessedFolder = 'r'<<24 | 'o'<<16 | 't'<<8 | 'f' // 'rotf'
)

func (e KBlessedFolder) String() string {
	switch e {
	case KBlessedFolderValue:
		return "KBlessedFolderValue"
	case KRootFolder:
		return "KRootFolder"
	default:
		return fmt.Sprintf("KBlessedFolder(%d)", e)
	}
}

type KBy uint

const (
	KByCommentView KBy = 6
	KByDateView    KBy = 3
	KByIconView    KBy = 1
	KByKindView    KBy = 5
	KByLabelView   KBy = 7
	KByNameView    KBy = 2
	KBySizeView    KBy = 4
	KBySmallIcon   KBy = 0
	KByVersionView KBy = 8
)

func (e KBy) String() string {
	switch e {
	case KByCommentView:
		return "KByCommentView"
	case KByDateView:
		return "KByDateView"
	case KByIconView:
		return "KByIconView"
	case KByKindView:
		return "KByKindView"
	case KByLabelView:
		return "KByLabelView"
	case KByNameView:
		return "KByNameView"
	case KBySizeView:
		return "KBySizeView"
	case KBySmallIcon:
		return "KBySmallIcon"
	case KByVersionView:
		return "KByVersionView"
	default:
		return fmt.Sprintf("KBy(%d)", e)
	}
}

type KCSAccept uint

const (
	// Deprecated.
	KCSAcceptAllComponentsMode KCSAccept = 0
	// Deprecated.
	KCSAcceptThreadSafeComponentsOnlyMode KCSAccept = 1
)

func (e KCSAccept) String() string {
	switch e {
	case KCSAcceptAllComponentsMode:
		return "KCSAcceptAllComponentsMode"
	case KCSAcceptThreadSafeComponentsOnlyMode:
		return "KCSAcceptThreadSafeComponentsOnlyMode"
	default:
		return fmt.Sprintf("KCSAccept(%d)", e)
	}
}

type KCSDiskSpaceRecoveryOptionNoU uint

const (
	KCSDiskSpaceRecoveryOptionNoUI KCSDiskSpaceRecoveryOptionNoU = 1
)

func (e KCSDiskSpaceRecoveryOptionNoU) String() string {
	switch e {
	case KCSDiskSpaceRecoveryOptionNoUI:
		return "KCSDiskSpaceRecoveryOptionNoUI"
	default:
		return fmt.Sprintf("KCSDiskSpaceRecoveryOptionNoU(%d)", e)
	}
}

type KCSIdentity int

const (
	KCSIdentityAuthorityNotAccessibleErr KCSIdentity = -2
	KCSIdentityDeletedErr                KCSIdentity = -4
	KCSIdentityDuplicateFullNameErr      KCSIdentity = -6
	KCSIdentityDuplicatePosixNameErr     KCSIdentity = -8
	KCSIdentityInvalidFullNameErr        KCSIdentity = -5
	KCSIdentityInvalidPosixNameErr       KCSIdentity = -7
	KCSIdentityPermissionErr             KCSIdentity = -3
	KCSIdentityUnknownAuthorityErr       KCSIdentity = -1
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
	KCSIdentityClassUser  KCSIdentityClass = 1
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

type KCSIdentityCommit uint

const (
	KCSIdentityCommitCompleted KCSIdentityCommit = 1
)

func (e KCSIdentityCommit) String() string {
	switch e {
	case KCSIdentityCommitCompleted:
		return "KCSIdentityCommitCompleted"
	default:
		return fmt.Sprintf("KCSIdentityCommit(%d)", e)
	}
}

type KCSIdentityFlag uint

const (
	KCSIdentityFlagHidden KCSIdentityFlag = 1
	KCSIdentityFlagNone   KCSIdentityFlag = 0
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
	KCSIdentityQueryGenerateUpdateEvents    KCSIdentityQuery = 1
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
	KCSIdentityQueryEventErrorOccurred       KCSIdentityQueryEvent = 5
	KCSIdentityQueryEventResultsAdded        KCSIdentityQueryEvent = 2
	KCSIdentityQueryEventResultsChanged      KCSIdentityQueryEvent = 3
	KCSIdentityQueryEventResultsRemoved      KCSIdentityQueryEvent = 4
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
	KCSIdentityQueryStringEquals     KCSIdentityQueryString = 1
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

type KCallingConventionWidth uint

const (
	// Deprecated.
	KCallingConventionMask KCallingConventionWidth = 15
	// Deprecated.
	KCallingConventionPhase KCallingConventionWidth = 0
	// Deprecated.
	KCallingConventionWidthValue KCallingConventionWidth = 4
	// Deprecated.
	KDispatchedParameterPhase KCallingConventionWidth = 0
	// Deprecated.
	KDispatchedSelectorSizePhase KCallingConventionWidth = 0
	// Deprecated.
	KDispatchedSelectorSizeWidth KCallingConventionWidth = 2
	// Deprecated.
	KRegisterParameterMask KCallingConventionWidth = 2147481600
	// Deprecated.
	KRegisterParameterPhase KCallingConventionWidth = 0
	// Deprecated.
	KRegisterParameterSizePhase KCallingConventionWidth = 0
	// Deprecated.
	KRegisterParameterSizeWidth KCallingConventionWidth = 2
	// Deprecated.
	KRegisterParameterWhichPhase KCallingConventionWidth = 2
	// Deprecated.
	KRegisterParameterWhichWidth KCallingConventionWidth = 3
	// Deprecated.
	KRegisterParameterWidth KCallingConventionWidth = 5
	// Deprecated.
	KRegisterResultLocationPhase KCallingConventionWidth = 0
	// Deprecated.
	KRegisterResultLocationWidth KCallingConventionWidth = 5
	// Deprecated.
	KResultSizeMask KCallingConventionWidth = 48
	// Deprecated.
	KResultSizePhase KCallingConventionWidth = 4
	// Deprecated.
	KResultSizeWidth KCallingConventionWidth = 2
	// Deprecated.
	KSpecialCaseSelectorMask KCallingConventionWidth = 1008
	// Deprecated.
	KSpecialCaseSelectorPhase KCallingConventionWidth = 4
	// Deprecated.
	KSpecialCaseSelectorWidth KCallingConventionWidth = 6
	// Deprecated.
	KStackParameterMask KCallingConventionWidth = 0
	// Deprecated.
	KStackParameterPhase KCallingConventionWidth = 0
	// Deprecated.
	KStackParameterWidth KCallingConventionWidth = 2
)

func (e KCallingConventionWidth) String() string {
	switch e {
	case KCallingConventionMask:
		return "KCallingConventionMask"
	case KCallingConventionPhase:
		return "KCallingConventionPhase"
	case KCallingConventionWidthValue:
		return "KCallingConventionWidthValue"
	case KDispatchedSelectorSizeWidth:
		return "KDispatchedSelectorSizeWidth"
	case KRegisterParameterMask:
		return "KRegisterParameterMask"
	case KRegisterParameterWhichWidth:
		return "KRegisterParameterWhichWidth"
	case KRegisterParameterWidth:
		return "KRegisterParameterWidth"
	case KResultSizeMask:
		return "KResultSizeMask"
	case KSpecialCaseSelectorMask:
		return "KSpecialCaseSelectorMask"
	case KSpecialCaseSelectorWidth:
		return "KSpecialCaseSelectorWidth"
	default:
		return fmt.Sprintf("KCallingConventionWidth(%d)", e)
	}
}

type KCertSearch uint

const (
	KCertSearchAny               KCertSearch = 0
	KCertSearchDecryptAllowed    KCertSearch = 1
	KCertSearchDecryptDisallowed KCertSearch = 1
	KCertSearchDecryptIgnored    KCertSearch = 0
	KCertSearchDecryptMask       KCertSearch = 0
	KCertSearchEncryptAllowed    KCertSearch = 1
	KCertSearchEncryptDisallowed KCertSearch = 1
	KCertSearchEncryptIgnored    KCertSearch = 0
	KCertSearchEncryptMask       KCertSearch = 0
	KCertSearchPrivKeyRequired   KCertSearch = 1
	KCertSearchShift             KCertSearch = 0
	KCertSearchSigningAllowed    KCertSearch = 1
	KCertSearchSigningDisallowed KCertSearch = 1
	KCertSearchSigningIgnored    KCertSearch = 0
	KCertSearchSigningMask       KCertSearch = 0
	KCertSearchUnwrapAllowed     KCertSearch = 1
	KCertSearchUnwrapDisallowed  KCertSearch = 1
	KCertSearchUnwrapIgnored     KCertSearch = 0
	KCertSearchUnwrapMask        KCertSearch = 0
	KCertSearchVerifyAllowed     KCertSearch = 1
	KCertSearchVerifyDisallowed  KCertSearch = 1
	KCertSearchVerifyIgnored     KCertSearch = 0
	KCertSearchVerifyMask        KCertSearch = 0
	KCertSearchWrapAllowed       KCertSearch = 1
	KCertSearchWrapDisallowed    KCertSearch = 1
	KCertSearchWrapIgnored       KCertSearch = 0
	KCertSearchWrapMask          KCertSearch = 0
)

func (e KCertSearch) String() string {
	switch e {
	case KCertSearchAny:
		return "KCertSearchAny"
	case KCertSearchDecryptAllowed:
		return "KCertSearchDecryptAllowed"
	default:
		return fmt.Sprintf("KCertSearch(%d)", e)
	}
}

type KCertificateKCItemClass uint

const (
	KAppleSharePasswordKCItemClass KCertificateKCItemClass = 'a'<<24 | 's'<<16 | 'h'<<8 | 'p' // 'ashp'
	KCertificateKCItemClassValue   KCertificateKCItemClass = 'c'<<24 | 'e'<<16 | 'r'<<8 | 't' // 'cert'
	KGenericPasswordKCItemClass    KCertificateKCItemClass = 'g'<<24 | 'e'<<16 | 'n'<<8 | 'p' // 'genp'
	KInternetPasswordKCItemClass   KCertificateKCItemClass = 'i'<<24 | 'n'<<16 | 'e'<<8 | 't' // 'inet'
)

func (e KCertificateKCItemClass) String() string {
	switch e {
	case KAppleSharePasswordKCItemClass:
		return "KAppleSharePasswordKCItemClass"
	case KCertificateKCItemClassValue:
		return "KCertificateKCItemClassValue"
	case KGenericPasswordKCItemClass:
		return "KGenericPasswordKCItemClass"
	case KInternetPasswordKCItemClass:
		return "KInternetPasswordKCItemClass"
	default:
		return fmt.Sprintf("KCertificateKCItemClass(%d)", e)
	}
}

type KClassKCItemAttr uint

const (
	KAccountKCItemAttr        KClassKCItemAttr = 'a'<<24 | 'c'<<16 | 'c'<<8 | 't' // 'acct'
	KAddressKCItemAttr        KClassKCItemAttr = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'addr'
	KAuthTypeKCItemAttr       KClassKCItemAttr = 'a'<<24 | 't'<<16 | 'y'<<8 | 'p' // 'atyp'
	KClassKCItemAttrValue     KClassKCItemAttr = 'c'<<24 | 'l'<<16 | 'a'<<8 | 's' // 'clas'
	KCommentKCItemAttr        KClassKCItemAttr = 'i'<<24 | 'c'<<16 | 'm'<<8 | 't' // 'icmt'
	KCommonNameKCItemAttr     KClassKCItemAttr = 'c'<<24 | 'n'<<16 | ' '<<8 | ' ' // 'cn  '
	KCreationDateKCItemAttr   KClassKCItemAttr = 'c'<<24 | 'd'<<16 | 'a'<<8 | 't' // 'cdat'
	KCreatorKCItemAttr        KClassKCItemAttr = 'c'<<24 | 'r'<<16 | 't'<<8 | 'r' // 'crtr'
	KCustomIconKCItemAttr     KClassKCItemAttr = 'c'<<24 | 'u'<<16 | 's'<<8 | 'i' // 'cusi'
	KDecryptKCItemAttr        KClassKCItemAttr = 'd'<<24 | 'e'<<16 | 'c'<<8 | 'r' // 'decr'
	KDescriptionKCItemAttr    KClassKCItemAttr = 'd'<<24 | 'e'<<16 | 's'<<8 | 'c' // 'desc'
	KEMailKCItemAttr          KClassKCItemAttr = 'm'<<24 | 'a'<<16 | 'i'<<8 | 'l' // 'mail'
	KEncryptKCItemAttr        KClassKCItemAttr = 'e'<<24 | 'n'<<16 | 'c'<<8 | 'r' // 'encr'
	KEndDateKCItemAttr        KClassKCItemAttr = 'e'<<24 | 'd'<<16 | 'a'<<8 | 't' // 'edat'
	KGenericKCItemAttr        KClassKCItemAttr = 'g'<<24 | 'e'<<16 | 'n'<<8 | 'a' // 'gena'
	KInvisibleKCItemAttr      KClassKCItemAttr = 'i'<<24 | 'n'<<16 | 'v'<<8 | 'i' // 'invi'
	KIssuerKCItemAttr         KClassKCItemAttr = 'i'<<24 | 's'<<16 | 's'<<8 | 'u' // 'issu'
	KIssuerURLKCItemAttr      KClassKCItemAttr = 'i'<<24 | 'u'<<16 | 'r'<<8 | 'l' // 'iurl'
	KLabelKCItemAttr          KClassKCItemAttr = 'l'<<24 | 'a'<<16 | 'b'<<8 | 'l' // 'labl'
	KModDateKCItemAttr        KClassKCItemAttr = 'm'<<24 | 'd'<<16 | 'a'<<8 | 't' // 'mdat'
	KNegativeKCItemAttr       KClassKCItemAttr = 'n'<<24 | 'e'<<16 | 'g'<<8 | 'a' // 'nega'
	KPathKCItemAttr           KClassKCItemAttr = 'p'<<24 | 'a'<<16 | 't'<<8 | 'h' // 'path'
	KPortKCItemAttr           KClassKCItemAttr = 'p'<<24 | 'o'<<16 | 'r'<<8 | 't' // 'port'
	KProtocolKCItemAttr       KClassKCItemAttr = 'p'<<24 | 't'<<16 | 'c'<<8 | 'l' // 'ptcl'
	KPublicKeyHashKCItemAttr  KClassKCItemAttr = 'h'<<24 | 'p'<<16 | 'k'<<8 | 'y' // 'hpky'
	KScriptCodeKCItemAttr     KClassKCItemAttr = 's'<<24 | 'c'<<16 | 'r'<<8 | 'p' // 'scrp'
	KSecurityDomainKCItemAttr KClassKCItemAttr = 's'<<24 | 'd'<<16 | 'm'<<8 | 'n' // 'sdmn'
	KSerialNumberKCItemAttr   KClassKCItemAttr = 's'<<24 | 'n'<<16 | 'b'<<8 | 'r' // 'snbr'
	KServerKCItemAttr         KClassKCItemAttr = 's'<<24 | 'r'<<16 | 'v'<<8 | 'r' // 'srvr'
	KServiceKCItemAttr        KClassKCItemAttr = 's'<<24 | 'v'<<16 | 'c'<<8 | 'e' // 'svce'
	KSignKCItemAttr           KClassKCItemAttr = 's'<<24 | 'i'<<16 | 'g'<<8 | 'n' // 'sign'
	KSignatureKCItemAttr      KClassKCItemAttr = 's'<<24 | 's'<<16 | 'i'<<8 | 'g' // 'ssig'
	KStartDateKCItemAttr      KClassKCItemAttr = 's'<<24 | 'd'<<16 | 'a'<<8 | 't' // 'sdat'
	KSubjectKCItemAttr        KClassKCItemAttr = 's'<<24 | 'u'<<16 | 'b'<<8 | 'j' // 'subj'
	KTypeKCItemAttr           KClassKCItemAttr = 't'<<24 | 'y'<<16 | 'p'<<8 | 'e' // 'type'
	KUnwrapKCItemAttr         KClassKCItemAttr = 'u'<<24 | 'n'<<16 | 'w'<<8 | 'r' // 'unwr'
	KVerifyKCItemAttr         KClassKCItemAttr = 'v'<<24 | 'e'<<16 | 'r'<<8 | 'i' // 'veri'
	KVolumeKCItemAttr         KClassKCItemAttr = 'v'<<24 | 'l'<<16 | 'm'<<8 | 'e' // 'vlme'
	KWrapKCItemAttr           KClassKCItemAttr = 'w'<<24 | 'r'<<16 | 'a'<<8 | 'p' // 'wrap'
)

func (e KClassKCItemAttr) String() string {
	switch e {
	case KAccountKCItemAttr:
		return "KAccountKCItemAttr"
	case KAddressKCItemAttr:
		return "KAddressKCItemAttr"
	case KAuthTypeKCItemAttr:
		return "KAuthTypeKCItemAttr"
	case KClassKCItemAttrValue:
		return "KClassKCItemAttrValue"
	case KCommentKCItemAttr:
		return "KCommentKCItemAttr"
	case KCommonNameKCItemAttr:
		return "KCommonNameKCItemAttr"
	case KCreationDateKCItemAttr:
		return "KCreationDateKCItemAttr"
	case KCreatorKCItemAttr:
		return "KCreatorKCItemAttr"
	case KCustomIconKCItemAttr:
		return "KCustomIconKCItemAttr"
	case KDecryptKCItemAttr:
		return "KDecryptKCItemAttr"
	case KDescriptionKCItemAttr:
		return "KDescriptionKCItemAttr"
	case KEMailKCItemAttr:
		return "KEMailKCItemAttr"
	case KEncryptKCItemAttr:
		return "KEncryptKCItemAttr"
	case KEndDateKCItemAttr:
		return "KEndDateKCItemAttr"
	case KGenericKCItemAttr:
		return "KGenericKCItemAttr"
	case KInvisibleKCItemAttr:
		return "KInvisibleKCItemAttr"
	case KIssuerKCItemAttr:
		return "KIssuerKCItemAttr"
	case KIssuerURLKCItemAttr:
		return "KIssuerURLKCItemAttr"
	case KLabelKCItemAttr:
		return "KLabelKCItemAttr"
	case KModDateKCItemAttr:
		return "KModDateKCItemAttr"
	case KNegativeKCItemAttr:
		return "KNegativeKCItemAttr"
	case KPathKCItemAttr:
		return "KPathKCItemAttr"
	case KPortKCItemAttr:
		return "KPortKCItemAttr"
	case KProtocolKCItemAttr:
		return "KProtocolKCItemAttr"
	case KPublicKeyHashKCItemAttr:
		return "KPublicKeyHashKCItemAttr"
	case KScriptCodeKCItemAttr:
		return "KScriptCodeKCItemAttr"
	case KSecurityDomainKCItemAttr:
		return "KSecurityDomainKCItemAttr"
	case KSerialNumberKCItemAttr:
		return "KSerialNumberKCItemAttr"
	case KServerKCItemAttr:
		return "KServerKCItemAttr"
	case KServiceKCItemAttr:
		return "KServiceKCItemAttr"
	case KSignKCItemAttr:
		return "KSignKCItemAttr"
	case KSignatureKCItemAttr:
		return "KSignatureKCItemAttr"
	case KStartDateKCItemAttr:
		return "KStartDateKCItemAttr"
	case KSubjectKCItemAttr:
		return "KSubjectKCItemAttr"
	case KTypeKCItemAttr:
		return "KTypeKCItemAttr"
	case KUnwrapKCItemAttr:
		return "KUnwrapKCItemAttr"
	case KVerifyKCItemAttr:
		return "KVerifyKCItemAttr"
	case KVolumeKCItemAttr:
		return "KVolumeKCItemAttr"
	case KWrapKCItemAttr:
		return "KWrapKCItemAttr"
	default:
		return fmt.Sprintf("KClassKCItemAttr(%d)", e)
	}
}

type KClassic int

const (
	// Deprecated.
	KClassicDomain KClassic = -32762
)

func (e KClassic) String() string {
	switch e {
	case KClassicDomain:
		return "KClassicDomain"
	default:
		return fmt.Sprintf("KClassic(%d)", e)
	}
}

type KClassicPreferencesFolder uint

const (
	// Deprecated.
	KClassicPreferencesFolderType KClassicPreferencesFolder = 'c'<<24 | 'p'<<16 | 'r'<<8 | 'f' // 'cprf'
)

func (e KClassicPreferencesFolder) String() string {
	switch e {
	case KClassicPreferencesFolderType:
		return "KClassicPreferencesFolderType"
	default:
		return fmt.Sprintf("KClassicPreferencesFolder(%d)", e)
	}
}

type KClipboardIcon uint

const (
	KClipboardIconValue            KClipboardIcon = 'C'<<24 | 'L'<<16 | 'I'<<8 | 'P' // 'CLIP'
	KClippingPictureTypeIcon       KClipboardIcon = 'c'<<24 | 'l'<<16 | 'p'<<8 | 'p' // 'clpp'
	KClippingSoundTypeIcon         KClipboardIcon = 'c'<<24 | 'l'<<16 | 'p'<<8 | 's' // 'clps'
	KClippingTextTypeIcon          KClipboardIcon = 'c'<<24 | 'l'<<16 | 'p'<<8 | 't' // 'clpt'
	KClippingUnknownTypeIcon       KClipboardIcon = 'c'<<24 | 'l'<<16 | 'p'<<8 | 'u' // 'clpu'
	KComputerIcon                  KClipboardIcon = 'r'<<24 | 'o'<<16 | 'o'<<8 | 't' // 'root'
	KDesktopIcon                   KClipboardIcon = 'd'<<24 | 'e'<<16 | 's'<<8 | 'k' // 'desk'
	KFinderIcon                    KClipboardIcon = 'F'<<24 | 'N'<<16 | 'D'<<8 | 'R' // 'FNDR'
	KFontSuitcaseIcon              KClipboardIcon = 'F'<<24 | 'F'<<16 | 'I'<<8 | 'L' // 'FFIL'
	KFullTrashIcon                 KClipboardIcon = 'f'<<24 | 't'<<16 | 'r'<<8 | 'h' // 'ftrh'
	KGenericApplicationIcon        KClipboardIcon = 'A'<<24 | 'P'<<16 | 'P'<<8 | 'L' // 'APPL'
	KGenericCDROMIcon              KClipboardIcon = 'c'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'cddr'
	KGenericComponentIcon          KClipboardIcon = 't'<<24 | 'h'<<16 | 'n'<<8 | 'g' // 'thng'
	KGenericControlPanelIcon       KClipboardIcon = 'A'<<24 | 'P'<<16 | 'P'<<8 | 'C' // 'APPC'
	KGenericControlStripModuleIcon KClipboardIcon = 's'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'sdev'
	KGenericDeskAccessoryIcon      KClipboardIcon = 'A'<<24 | 'P'<<16 | 'P'<<8 | 'D' // 'APPD'
	KGenericDocumentIcon           KClipboardIcon = 'd'<<24 | 'o'<<16 | 'c'<<8 | 'u' // 'docu'
	KGenericEditionFileIcon        KClipboardIcon = 'e'<<24 | 'd'<<16 | 't'<<8 | 'f' // 'edtf'
	KGenericExtensionIcon          KClipboardIcon = 'I'<<24 | 'N'<<16 | 'I'<<8 | 'T' // 'INIT'
	KGenericFileServerIcon         KClipboardIcon = 's'<<24 | 'r'<<16 | 'v'<<8 | 'r' // 'srvr'
	KGenericFloppyIcon             KClipboardIcon = 'f'<<24 | 'l'<<16 | 'p'<<8 | 'y' // 'flpy'
	KGenericFontIcon               KClipboardIcon = 'f'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'ffil'
	KGenericFontScalerIcon         KClipboardIcon = 's'<<24 | 'c'<<16 | 'l'<<8 | 'r' // 'sclr'
	KGenericHardDiskIcon           KClipboardIcon = 'h'<<24 | 'd'<<16 | 's'<<8 | 'k' // 'hdsk'
	KGenericIDiskIcon              KClipboardIcon = 'i'<<24 | 'd'<<16 | 's'<<8 | 'k' // 'idsk'
	KGenericMoverObjectIcon        KClipboardIcon = 'm'<<24 | 'o'<<16 | 'v'<<8 | 'r' // 'movr'
	KGenericPCCardIcon             KClipboardIcon = 'p'<<24 | 'c'<<16 | 'm'<<8 | 'c' // 'pcmc'
	KGenericPreferencesIcon        KClipboardIcon = 'p'<<24 | 'r'<<16 | 'e'<<8 | 'f' // 'pref'
	KGenericQueryDocumentIcon      KClipboardIcon = 'q'<<24 | 'e'<<16 | 'r'<<8 | 'y' // 'qery'
	KGenericRAMDiskIcon            KClipboardIcon = 'r'<<24 | 'a'<<16 | 'm'<<8 | 'd' // 'ramd'
	KGenericRemovableMediaIcon     KClipboardIcon = 'r'<<24 | 'm'<<16 | 'o'<<8 | 'v' // 'rmov'
	KGenericSharedLibaryIcon       KClipboardIcon = 's'<<24 | 'h'<<16 | 'l'<<8 | 'b' // 'shlb'
	KGenericStationeryIcon         KClipboardIcon = 's'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'sdoc'
	KGenericSuitcaseIcon           KClipboardIcon = 's'<<24 | 'u'<<16 | 'i'<<8 | 't' // 'suit'
	KGenericURLIcon                KClipboardIcon = 'g'<<24 | 'u'<<16 | 'r'<<8 | 'l' // 'gurl'
	KGenericWORMIcon               KClipboardIcon = 'w'<<24 | 'o'<<16 | 'r'<<8 | 'm' // 'worm'
	KInternationResourcesIcon      KClipboardIcon = 'i'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'ifil'
	KInternationalResourcesIcon    KClipboardIcon = 'i'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'ifil'
	KKeyboardLayoutIcon            KClipboardIcon = 'k'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'kfil'
	KSoundFileIcon                 KClipboardIcon = 's'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'sfil'
	KSystemSuitcaseIcon            KClipboardIcon = 'z'<<24 | 's'<<16 | 'y'<<8 | 's' // 'zsys'
	KTrashIcon                     KClipboardIcon = 't'<<24 | 'r'<<16 | 's'<<8 | 'h' // 'trsh'
	KTrueTypeFlatFontIcon          KClipboardIcon = 's'<<24 | 'f'<<16 | 'n'<<8 | 't' // 'sfnt'
	KTrueTypeFontIcon              KClipboardIcon = 't'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'tfil'
	KTrueTypeMultiFlatFontIcon     KClipboardIcon = 't'<<24 | 't'<<16 | 'c'<<8 | 'f' // 'ttcf'
	KUnknownFSObjectIcon           KClipboardIcon = 'u'<<24 | 'n'<<16 | 'f'<<8 | 's' // 'unfs'
	KUserIDiskIcon                 KClipboardIcon = 'u'<<24 | 'd'<<16 | 's'<<8 | 'k' // 'udsk'
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

type KClipping uint

const (
	KClippingCreator     KClipping = 'd'<<24 | 'r'<<16 | 'a'<<8 | 'g' // 'drag'
	KClippingPictureType KClipping = 'c'<<24 | 'l'<<16 | 'p'<<8 | 'p' // 'clpp'
	KClippingSoundType   KClipping = 'c'<<24 | 'l'<<16 | 'p'<<8 | 's' // 'clps'
	KClippingTextType    KClipping = 'c'<<24 | 'l'<<16 | 'p'<<8 | 't' // 'clpt'
	KClippingUnknownType KClipping = 'c'<<24 | 'l'<<16 | 'p'<<8 | 'u' // 'clpu'
)

func (e KClipping) String() string {
	switch e {
	case KClippingCreator:
		return "KClippingCreator"
	case KClippingPictureType:
		return "KClippingPictureType"
	case KClippingSoundType:
		return "KClippingSoundType"
	case KClippingTextType:
		return "KClippingTextType"
	case KClippingUnknownType:
		return "KClippingUnknownType"
	default:
		return fmt.Sprintf("KClipping(%d)", e)
	}
}

type KCollate int

const (
	KCollateAttributesNotFoundErr   KCollate = -29500
	KCollateBufferTooSmall          KCollate = -29506
	KCollateInvalidChar             KCollate = -29505
	KCollateInvalidCollationRef     KCollate = -29507
	KCollateInvalidOptions          KCollate = -29501
	KCollateMissingUnicodeTableErr  KCollate = -29502
	KCollatePatternNotFoundErr      KCollate = -29504
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

type KCollection uint

const (
	// Deprecated.
	KCollectionAllAttributes KCollection = 0
	// Deprecated.
	KCollectionDefaultAttributes KCollection = 1073741824
	// Deprecated.
	KCollectionLockBit KCollection = 31
	// Deprecated.
	KCollectionLockMask KCollection = 2147483648
	// Deprecated.
	KCollectionNoAttributes KCollection = 0
	// Deprecated.
	KCollectionPersistenceBit KCollection = 30
	// Deprecated.
	KCollectionPersistenceMask KCollection = 1073741824
	// Deprecated.
	KCollectionReserved0Bit KCollection = 16
	// Deprecated.
	KCollectionReserved0Mask KCollection = 65536
	// Deprecated.
	KCollectionReserved10Bit KCollection = 26
	// Deprecated.
	KCollectionReserved10Mask KCollection = 67108864
	// Deprecated.
	KCollectionReserved11Bit KCollection = 27
	// Deprecated.
	KCollectionReserved11Mask KCollection = 134217728
	// Deprecated.
	KCollectionReserved12Bit KCollection = 28
	// Deprecated.
	KCollectionReserved12Mask KCollection = 268435456
	// Deprecated.
	KCollectionReserved13Bit KCollection = 29
	// Deprecated.
	KCollectionReserved13Mask KCollection = 536870912
	// Deprecated.
	KCollectionReserved1Bit KCollection = 17
	// Deprecated.
	KCollectionReserved1Mask KCollection = 131072
	// Deprecated.
	KCollectionReserved2Bit KCollection = 18
	// Deprecated.
	KCollectionReserved2Mask KCollection = 262144
	// Deprecated.
	KCollectionReserved3Bit KCollection = 19
	// Deprecated.
	KCollectionReserved3Mask KCollection = 524288
	// Deprecated.
	KCollectionReserved4Bit KCollection = 20
	// Deprecated.
	KCollectionReserved4Mask KCollection = 1048576
	// Deprecated.
	KCollectionReserved5Bit KCollection = 21
	// Deprecated.
	KCollectionReserved5Mask KCollection = 2097152
	// Deprecated.
	KCollectionReserved6Bit KCollection = 22
	// Deprecated.
	KCollectionReserved6Mask KCollection = 4194304
	// Deprecated.
	KCollectionReserved7Bit KCollection = 23
	// Deprecated.
	KCollectionReserved7Mask KCollection = 8388608
	// Deprecated.
	KCollectionReserved8Bit KCollection = 24
	// Deprecated.
	KCollectionReserved8Mask KCollection = 16777216
	// Deprecated.
	KCollectionReserved9Bit KCollection = 25
	// Deprecated.
	KCollectionReserved9Mask KCollection = 33554432
	// Deprecated.
	KCollectionUser0Bit KCollection = 0
	// Deprecated.
	KCollectionUser0Mask KCollection = 1
	// Deprecated.
	KCollectionUser10Bit KCollection = 10
	// Deprecated.
	KCollectionUser10Mask KCollection = 1024
	// Deprecated.
	KCollectionUser11Bit KCollection = 11
	// Deprecated.
	KCollectionUser11Mask KCollection = 2048
	// Deprecated.
	KCollectionUser12Bit KCollection = 12
	// Deprecated.
	KCollectionUser12Mask KCollection = 4096
	// Deprecated.
	KCollectionUser13Bit KCollection = 13
	// Deprecated.
	KCollectionUser13Mask KCollection = 8192
	// Deprecated.
	KCollectionUser14Bit KCollection = 14
	// Deprecated.
	KCollectionUser14Mask KCollection = 16384
	// Deprecated.
	KCollectionUser15Bit KCollection = 15
	// Deprecated.
	KCollectionUser15Mask KCollection = 32768
	// Deprecated.
	KCollectionUser1Bit KCollection = 1
	// Deprecated.
	KCollectionUser1Mask KCollection = 2
	// Deprecated.
	KCollectionUser2Bit KCollection = 2
	// Deprecated.
	KCollectionUser2Mask KCollection = 4
	// Deprecated.
	KCollectionUser3Bit KCollection = 3
	// Deprecated.
	KCollectionUser3Mask KCollection = 8
	// Deprecated.
	KCollectionUser4Bit KCollection = 4
	// Deprecated.
	KCollectionUser4Mask KCollection = 16
	// Deprecated.
	KCollectionUser5Bit KCollection = 5
	// Deprecated.
	KCollectionUser5Mask KCollection = 32
	// Deprecated.
	KCollectionUser6Bit KCollection = 6
	// Deprecated.
	KCollectionUser6Mask KCollection = 64
	// Deprecated.
	KCollectionUser7Bit KCollection = 7
	// Deprecated.
	KCollectionUser7Mask KCollection = 128
	// Deprecated.
	KCollectionUser8Bit KCollection = 8
	// Deprecated.
	KCollectionUser8Mask KCollection = 256
	// Deprecated.
	KCollectionUser9Bit KCollection = 9
	// Deprecated.
	KCollectionUser9Mask KCollection = 512
	// Deprecated.
	KCollectionUserAttributes KCollection = 65535
)

func (e KCollection) String() string {
	switch e {
	case KCollectionAllAttributes:
		return "KCollectionAllAttributes"
	case KCollectionDefaultAttributes:
		return "KCollectionDefaultAttributes"
	case KCollectionLockBit:
		return "KCollectionLockBit"
	case KCollectionLockMask:
		return "KCollectionLockMask"
	case KCollectionPersistenceBit:
		return "KCollectionPersistenceBit"
	case KCollectionReserved0Bit:
		return "KCollectionReserved0Bit"
	case KCollectionReserved0Mask:
		return "KCollectionReserved0Mask"
	case KCollectionReserved10Bit:
		return "KCollectionReserved10Bit"
	case KCollectionReserved10Mask:
		return "KCollectionReserved10Mask"
	case KCollectionReserved11Bit:
		return "KCollectionReserved11Bit"
	case KCollectionReserved11Mask:
		return "KCollectionReserved11Mask"
	case KCollectionReserved12Bit:
		return "KCollectionReserved12Bit"
	case KCollectionReserved12Mask:
		return "KCollectionReserved12Mask"
	case KCollectionReserved13Bit:
		return "KCollectionReserved13Bit"
	case KCollectionReserved13Mask:
		return "KCollectionReserved13Mask"
	case KCollectionReserved1Bit:
		return "KCollectionReserved1Bit"
	case KCollectionReserved1Mask:
		return "KCollectionReserved1Mask"
	case KCollectionReserved2Bit:
		return "KCollectionReserved2Bit"
	case KCollectionReserved2Mask:
		return "KCollectionReserved2Mask"
	case KCollectionReserved3Bit:
		return "KCollectionReserved3Bit"
	case KCollectionReserved3Mask:
		return "KCollectionReserved3Mask"
	case KCollectionReserved4Bit:
		return "KCollectionReserved4Bit"
	case KCollectionReserved4Mask:
		return "KCollectionReserved4Mask"
	case KCollectionReserved5Bit:
		return "KCollectionReserved5Bit"
	case KCollectionReserved5Mask:
		return "KCollectionReserved5Mask"
	case KCollectionReserved6Bit:
		return "KCollectionReserved6Bit"
	case KCollectionReserved6Mask:
		return "KCollectionReserved6Mask"
	case KCollectionReserved7Bit:
		return "KCollectionReserved7Bit"
	case KCollectionReserved7Mask:
		return "KCollectionReserved7Mask"
	case KCollectionReserved8Bit:
		return "KCollectionReserved8Bit"
	case KCollectionReserved8Mask:
		return "KCollectionReserved8Mask"
	case KCollectionReserved9Bit:
		return "KCollectionReserved9Bit"
	case KCollectionReserved9Mask:
		return "KCollectionReserved9Mask"
	case KCollectionUser0Mask:
		return "KCollectionUser0Mask"
	case KCollectionUser10Bit:
		return "KCollectionUser10Bit"
	case KCollectionUser10Mask:
		return "KCollectionUser10Mask"
	case KCollectionUser11Bit:
		return "KCollectionUser11Bit"
	case KCollectionUser11Mask:
		return "KCollectionUser11Mask"
	case KCollectionUser12Bit:
		return "KCollectionUser12Bit"
	case KCollectionUser12Mask:
		return "KCollectionUser12Mask"
	case KCollectionUser13Bit:
		return "KCollectionUser13Bit"
	case KCollectionUser13Mask:
		return "KCollectionUser13Mask"
	case KCollectionUser14Bit:
		return "KCollectionUser14Bit"
	case KCollectionUser14Mask:
		return "KCollectionUser14Mask"
	case KCollectionUser15Bit:
		return "KCollectionUser15Bit"
	case KCollectionUser15Mask:
		return "KCollectionUser15Mask"
	case KCollectionUser1Mask:
		return "KCollectionUser1Mask"
	case KCollectionUser2Mask:
		return "KCollectionUser2Mask"
	case KCollectionUser3Bit:
		return "KCollectionUser3Bit"
	case KCollectionUser3Mask:
		return "KCollectionUser3Mask"
	case KCollectionUser5Bit:
		return "KCollectionUser5Bit"
	case KCollectionUser5Mask:
		return "KCollectionUser5Mask"
	case KCollectionUser6Bit:
		return "KCollectionUser6Bit"
	case KCollectionUser6Mask:
		return "KCollectionUser6Mask"
	case KCollectionUser7Bit:
		return "KCollectionUser7Bit"
	case KCollectionUser7Mask:
		return "KCollectionUser7Mask"
	case KCollectionUser8Mask:
		return "KCollectionUser8Mask"
	case KCollectionUser9Bit:
		return "KCollectionUser9Bit"
	case KCollectionUser9Mask:
		return "KCollectionUser9Mask"
	case KCollectionUserAttributes:
		return "KCollectionUserAttributes"
	default:
		return fmt.Sprintf("KCollection(%d)", e)
	}
}

type KCollectionDontWant uint

const (
	// Deprecated.
	KCollectionDontWantAttributes KCollectionDontWant = 0
	// Deprecated.
	KCollectionDontWantData KCollectionDontWant = 0
	// Deprecated.
	KCollectionDontWantId KCollectionDontWant = 0
	// Deprecated.
	KCollectionDontWantIndex KCollectionDontWant = 0
	// Deprecated.
	KCollectionDontWantSize KCollectionDontWant = 0
	// Deprecated.
	KCollectionDontWantTag KCollectionDontWant = 0
)

func (e KCollectionDontWant) String() string {
	switch e {
	case KCollectionDontWantAttributes:
		return "KCollectionDontWantAttributes"
	default:
		return fmt.Sprintf("KCollectionDontWant(%d)", e)
	}
}

type KColorSyncFolderType uint

const (
	// Deprecated.
	KAppleShareAuthenticationFolderType KColorSyncFolderType = 'a'<<24 | 'u'<<16 | 't'<<8 | 'h' // 'auth'
	// Deprecated.
	KAppleShareSupportFolderType KColorSyncFolderType = 's'<<24 | 'h'<<16 | 'a'<<8 | 'r' // 'shar'
	// Deprecated.
	KAudioAlertSoundsFolderType KColorSyncFolderType = 'a'<<24 | 'l'<<16 | 'r'<<8 | 't' // 'alrt'
	// Deprecated.
	KAudioComponentsFolderType KColorSyncFolderType = 'a'<<24 | 'c'<<16 | 'm'<<8 | 'p' // 'acmp'
	// Deprecated.
	KAudioDigidesignFolderType KColorSyncFolderType = 'a'<<24 | 'd'<<16 | 'i'<<8 | 'g' // 'adig'
	// Deprecated.
	KAudioPlugInsFolderType KColorSyncFolderType = 'a'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'aplg'
	// Deprecated.
	KAudioPresetsFolderType KColorSyncFolderType = 'a'<<24 | 'p'<<16 | 's'<<8 | 't' // 'apst'
	// Deprecated.
	KAudioSoundBanksFolderType KColorSyncFolderType = 'b'<<24 | 'a'<<16 | 'n'<<8 | 'k' // 'bank'
	// Deprecated.
	KAudioSoundsFolderType KColorSyncFolderType = 'a'<<24 | 's'<<16 | 'n'<<8 | 'd' // 'asnd'
	// Deprecated.
	KAudioSupportFolderType KColorSyncFolderType = 'a'<<24 | 'd'<<16 | 'i'<<8 | 'o' // 'adio'
	// Deprecated.
	KAudioVSTFolderType KColorSyncFolderType = 'a'<<24 | 'v'<<16 | 's'<<8 | 't' // 'avst'
	// Deprecated.
	KAutomatorWorkflowsFolderType KColorSyncFolderType = 'f'<<24 | 'l'<<16 | 'o'<<8 | 'w' // 'flow'
	// Deprecated.
	KAutosaveInformationFolderType KColorSyncFolderType = 'a'<<24 | 's'<<16 | 'a'<<8 | 'v' // 'asav'
	// Deprecated.
	KBootTimeStartupItemsFolderType KColorSyncFolderType = 'e'<<24 | 'm'<<16 | 'p'<<8 | 'z' // 'empz'
	// Deprecated.
	KCachedDataFolderType KColorSyncFolderType = 'c'<<24 | 'a'<<16 | 'c'<<8 | 'h' // 'cach'
	// Deprecated.
	KCarbonLibraryFolderType KColorSyncFolderType = 'c'<<24 | 'a'<<16 | 'r'<<8 | 'b' // 'carb'
	// Deprecated.
	KClassicDesktopFolderType KColorSyncFolderType = 's'<<24 | 'd'<<16 | 's'<<8 | 'k' // 'sdsk'
	// Deprecated.
	KColorPickersFolderType KColorSyncFolderType = 'c'<<24 | 'p'<<16 | 'k'<<8 | 'r' // 'cpkr'
	// Deprecated.
	KColorSyncCMMFolderType KColorSyncFolderType = 'c'<<24 | 'c'<<16 | 'm'<<8 | 'm' // 'ccmm'
	// Deprecated.
	KColorSyncFolderTypeValue KColorSyncFolderType = 's'<<24 | 'y'<<16 | 'n'<<8 | 'c' // 'sync'
	// Deprecated.
	KColorSyncScriptingFolderType KColorSyncFolderType = 'c'<<24 | 's'<<16 | 'c'<<8 | 'r' // 'cscr'
	// Deprecated.
	KComponentsFolderType KColorSyncFolderType = 'c'<<24 | 'm'<<16 | 'p'<<8 | 'd' // 'cmpd'
	// Deprecated.
	KCompositionsFolderType KColorSyncFolderType = 'c'<<24 | 'm'<<16 | 'p'<<8 | 's' // 'cmps'
	// Deprecated.
	KCoreServicesFolderType KColorSyncFolderType = 'c'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'csrv'
	// Deprecated.
	KDirectoryServicesFolderType KColorSyncFolderType = 'd'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'dsrv'
	// Deprecated.
	KDirectoryServicesPlugInsFolderType KColorSyncFolderType = 'd'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'dplg'
	// Deprecated.
	KDocumentationFolderType KColorSyncFolderType = 'i'<<24 | 'n'<<16 | 'f'<<8 | 'o' // 'info'
	// Deprecated.
	KDownloadsFolderType KColorSyncFolderType = 'd'<<24 | 'o'<<16 | 'w'<<8 | 'n' // 'down'
	// Deprecated.
	KFileSystemSupportFolderType KColorSyncFolderType = 'f'<<24 | 's'<<16 | 'y'<<8 | 's' // 'fsys'
	// Deprecated.
	KFindByContentIndexesFolderType KColorSyncFolderType = 'f'<<24 | 'b'<<16 | 'c'<<8 | 'x' // 'fbcx'
	// Deprecated.
	KFontCollectionsFolderType KColorSyncFolderType = 'f'<<24 | 'n'<<16 | 'c'<<8 | 'l' // 'fncl'
	// Deprecated.
	KFrameworksFolderType KColorSyncFolderType = 'f'<<24 | 'r'<<16 | 'a'<<8 | 'm' // 'fram'
	// Deprecated.
	KISSDownloadsFolderType KColorSyncFolderType = 'i'<<24 | 's'<<16 | 's'<<8 | 'd' // 'issd'
	// Deprecated.
	KIndexFilesFolderType KColorSyncFolderType = 'i'<<24 | 'n'<<16 | 'd'<<8 | 'x' // 'indx'
	// Deprecated.
	KInputManagersFolderType KColorSyncFolderType = 'i'<<24 | 'n'<<16 | 'p'<<8 | 't' // 'inpt'
	// Deprecated.
	KInputMethodsFolderType KColorSyncFolderType = 'i'<<24 | 'n'<<16 | 'p'<<8 | 'f' // 'inpf'
	// Deprecated.
	KInstallerReceiptsFolderType KColorSyncFolderType = 'r'<<24 | 'c'<<16 | 'p'<<8 | 't' // 'rcpt'
	// Deprecated.
	KKernelExtensionsFolderType KColorSyncFolderType = 'k'<<24 | 'e'<<16 | 'x'<<8 | 't' // 'kext'
	// Deprecated.
	KKeyboardLayoutsFolderType KColorSyncFolderType = 'k'<<24 | 'l'<<16 | 'a'<<8 | 'y' // 'klay'
	// Deprecated.
	KLibraryAssistantsFolderType KColorSyncFolderType = 'a'<<24 | 's'<<16 | 't'<<8 | 'l' // 'astl'
	// Deprecated.
	KMIDIDriversFolderType KColorSyncFolderType = 'm'<<24 | 'i'<<16 | 'd'<<8 | 'i' // 'midi'
	// Deprecated.
	KManagedItemsFolderType KColorSyncFolderType = 'm'<<24 | 'a'<<16 | 'n'<<8 | 'g' // 'mang'
	// Deprecated.
	KPrintersFolderType KColorSyncFolderType = 'i'<<24 | 'm'<<16 | 'p'<<8 | 'r' // 'impr'
	// Deprecated.
	KPrivateFrameworksFolderType KColorSyncFolderType = 'p'<<24 | 'f'<<16 | 'r'<<8 | 'm' // 'pfrm'
	// Deprecated.
	KQuickTimeComponentsFolderType KColorSyncFolderType = 'w'<<24 | 'c'<<16 | 'm'<<8 | 'p' // 'wcmp'
	// Deprecated.
	KSpeechFolderType KColorSyncFolderType = 's'<<24 | 'p'<<16 | 'c'<<8 | 'h' // 'spch'
	// Deprecated.
	KSpotlightImportersFolderType KColorSyncFolderType = 's'<<24 | 'i'<<16 | 'm'<<8 | 'p' // 'simp'
	// Deprecated.
	KSpotlightMetadataCacheFolderType KColorSyncFolderType = 's'<<24 | 'c'<<16 | 'c'<<8 | 'h' // 'scch'
	// Deprecated.
	KSpotlightSavedSearchesFolderType KColorSyncFolderType = 's'<<24 | 'p'<<16 | 'o'<<8 | 't' // 'spot'
	// Deprecated.
	KSystemSoundsFolderType KColorSyncFolderType = 's'<<24 | 's'<<16 | 'n'<<8 | 'd' // 'ssnd'
	// Deprecated.
	KUserSpecificTmpFolderType KColorSyncFolderType = 'u'<<24 | 't'<<16 | 'm'<<8 | 'p' // 'utmp'
	// Deprecated.
	KiMovieFolderType KColorSyncFolderType = 'i'<<24 | 'm'<<16 | 'o'<<8 | 'v' // 'imov'
	// Deprecated.
	KiMoviePlugInsFolderType KColorSyncFolderType = 'i'<<24 | 'm'<<16 | 'p'<<8 | 'i' // 'impi'
	// Deprecated.
	KiMovieSoundEffectsFolderType KColorSyncFolderType = 'i'<<24 | 'm'<<16 | 's'<<8 | 'e' // 'imse'
)

func (e KColorSyncFolderType) String() string {
	switch e {
	case KAppleShareAuthenticationFolderType:
		return "KAppleShareAuthenticationFolderType"
	case KAppleShareSupportFolderType:
		return "KAppleShareSupportFolderType"
	case KAudioAlertSoundsFolderType:
		return "KAudioAlertSoundsFolderType"
	case KAudioComponentsFolderType:
		return "KAudioComponentsFolderType"
	case KAudioDigidesignFolderType:
		return "KAudioDigidesignFolderType"
	case KAudioPlugInsFolderType:
		return "KAudioPlugInsFolderType"
	case KAudioPresetsFolderType:
		return "KAudioPresetsFolderType"
	case KAudioSoundBanksFolderType:
		return "KAudioSoundBanksFolderType"
	case KAudioSoundsFolderType:
		return "KAudioSoundsFolderType"
	case KAudioSupportFolderType:
		return "KAudioSupportFolderType"
	case KAudioVSTFolderType:
		return "KAudioVSTFolderType"
	case KAutomatorWorkflowsFolderType:
		return "KAutomatorWorkflowsFolderType"
	case KAutosaveInformationFolderType:
		return "KAutosaveInformationFolderType"
	case KBootTimeStartupItemsFolderType:
		return "KBootTimeStartupItemsFolderType"
	case KCachedDataFolderType:
		return "KCachedDataFolderType"
	case KCarbonLibraryFolderType:
		return "KCarbonLibraryFolderType"
	case KClassicDesktopFolderType:
		return "KClassicDesktopFolderType"
	case KColorPickersFolderType:
		return "KColorPickersFolderType"
	case KColorSyncCMMFolderType:
		return "KColorSyncCMMFolderType"
	case KColorSyncFolderTypeValue:
		return "KColorSyncFolderTypeValue"
	case KColorSyncScriptingFolderType:
		return "KColorSyncScriptingFolderType"
	case KComponentsFolderType:
		return "KComponentsFolderType"
	case KCompositionsFolderType:
		return "KCompositionsFolderType"
	case KCoreServicesFolderType:
		return "KCoreServicesFolderType"
	case KDirectoryServicesFolderType:
		return "KDirectoryServicesFolderType"
	case KDirectoryServicesPlugInsFolderType:
		return "KDirectoryServicesPlugInsFolderType"
	case KDocumentationFolderType:
		return "KDocumentationFolderType"
	case KDownloadsFolderType:
		return "KDownloadsFolderType"
	case KFileSystemSupportFolderType:
		return "KFileSystemSupportFolderType"
	case KFindByContentIndexesFolderType:
		return "KFindByContentIndexesFolderType"
	case KFontCollectionsFolderType:
		return "KFontCollectionsFolderType"
	case KFrameworksFolderType:
		return "KFrameworksFolderType"
	case KISSDownloadsFolderType:
		return "KISSDownloadsFolderType"
	case KIndexFilesFolderType:
		return "KIndexFilesFolderType"
	case KInputManagersFolderType:
		return "KInputManagersFolderType"
	case KInputMethodsFolderType:
		return "KInputMethodsFolderType"
	case KInstallerReceiptsFolderType:
		return "KInstallerReceiptsFolderType"
	case KKernelExtensionsFolderType:
		return "KKernelExtensionsFolderType"
	case KKeyboardLayoutsFolderType:
		return "KKeyboardLayoutsFolderType"
	case KLibraryAssistantsFolderType:
		return "KLibraryAssistantsFolderType"
	case KMIDIDriversFolderType:
		return "KMIDIDriversFolderType"
	case KManagedItemsFolderType:
		return "KManagedItemsFolderType"
	case KPrintersFolderType:
		return "KPrintersFolderType"
	case KPrivateFrameworksFolderType:
		return "KPrivateFrameworksFolderType"
	case KQuickTimeComponentsFolderType:
		return "KQuickTimeComponentsFolderType"
	case KSpeechFolderType:
		return "KSpeechFolderType"
	case KSpotlightImportersFolderType:
		return "KSpotlightImportersFolderType"
	case KSpotlightMetadataCacheFolderType:
		return "KSpotlightMetadataCacheFolderType"
	case KSpotlightSavedSearchesFolderType:
		return "KSpotlightSavedSearchesFolderType"
	case KSystemSoundsFolderType:
		return "KSystemSoundsFolderType"
	case KUserSpecificTmpFolderType:
		return "KUserSpecificTmpFolderType"
	case KiMovieFolderType:
		return "KiMovieFolderType"
	case KiMoviePlugInsFolderType:
		return "KiMoviePlugInsFolderType"
	case KiMovieSoundEffectsFolderType:
		return "KiMovieSoundEffectsFolderType"
	default:
		return fmt.Sprintf("KColorSyncFolderType(%d)", e)
	}
}

type KColorSyncProfilesFolderType uint

const (
	// Deprecated.
	KApplicationSupportFolderType KColorSyncProfilesFolderType = 'a'<<24 | 's'<<16 | 'u'<<8 | 'p' // 'asup'
	// Deprecated.
	KColorSyncProfilesFolderTypeValue KColorSyncProfilesFolderType = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'f' // 'prof'
	// Deprecated.
	KPrinterDescriptionFolderType KColorSyncProfilesFolderType = 'p'<<24 | 'p'<<16 | 'd'<<8 | 'f' // 'ppdf'
	// Deprecated.
	KPrinterDriverFolderType KColorSyncProfilesFolderType = 3295703652
	// Deprecated.
	KScriptingAdditionsFolderType KColorSyncProfilesFolderType = 3295896434
	// Deprecated.
	KTextEncodingsFolderType KColorSyncProfilesFolderType = 3295962488
)

func (e KColorSyncProfilesFolderType) String() string {
	switch e {
	case KApplicationSupportFolderType:
		return "KApplicationSupportFolderType"
	case KColorSyncProfilesFolderTypeValue:
		return "KColorSyncProfilesFolderTypeValue"
	case KPrinterDescriptionFolderType:
		return "KPrinterDescriptionFolderType"
	case KPrinterDriverFolderType:
		return "KPrinterDriverFolderType"
	case KScriptingAdditionsFolderType:
		return "KScriptingAdditionsFolderType"
	case KTextEncodingsFolderType:
		return "KTextEncodingsFolderType"
	default:
		return fmt.Sprintf("KColorSyncProfilesFolderType(%d)", e)
	}
}

type KComponent int

const (
	// Deprecated.
	KComponentCanDoSelect KComponent = -3
	// Deprecated.
	KComponentCloseSelect KComponent = -2
	// Deprecated.
	KComponentExecuteWiredActionSelect KComponent = -9
	// Deprecated.
	KComponentGetMPWorkFunctionSelect KComponent = -8
	// Deprecated.
	KComponentGetPublicResourceSelect KComponent = -10
	// Deprecated.
	KComponentOpenSelect KComponent = -1
	// Deprecated.
	KComponentRegisterSelect KComponent = -5
	// Deprecated.
	KComponentTargetSelect KComponent = -6
	// Deprecated.
	KComponentUnregisterSelect KComponent = -7
	// Deprecated.
	KComponentVersionSelect KComponent = -4
)

func (e KComponent) String() string {
	switch e {
	case KComponentCanDoSelect:
		return "KComponentCanDoSelect"
	case KComponentCloseSelect:
		return "KComponentCloseSelect"
	case KComponentExecuteWiredActionSelect:
		return "KComponentExecuteWiredActionSelect"
	case KComponentGetMPWorkFunctionSelect:
		return "KComponentGetMPWorkFunctionSelect"
	case KComponentGetPublicResourceSelect:
		return "KComponentGetPublicResourceSelect"
	case KComponentOpenSelect:
		return "KComponentOpenSelect"
	case KComponentRegisterSelect:
		return "KComponentRegisterSelect"
	case KComponentTargetSelect:
		return "KComponentTargetSelect"
	case KComponentUnregisterSelect:
		return "KComponentUnregisterSelect"
	case KComponentVersionSelect:
		return "KComponentVersionSelect"
	default:
		return fmt.Sprintf("KComponent(%d)", e)
	}
}

type KComponentDebug uint

const (
	// Deprecated.
	KComponentDebugOption KComponentDebug = 0
)

func (e KComponentDebug) String() string {
	switch e {
	case KComponentDebugOption:
		return "KComponentDebugOption"
	default:
		return fmt.Sprintf("KComponentDebug(%d)", e)
	}
}

type KConnSuite uint

const (
	CADBAddress       KConnSuite = 'c'<<24 | 'a'<<16 | 'd'<<8 | 'b' // 'cadb'
	CAddressSpec      KConnSuite = 'c'<<24 | 'a'<<16 | 'd'<<8 | 'r' // 'cadr'
	CAppleTalkAddress KConnSuite = 'c'<<24 | 'a'<<16 | 't'<<8 | ' ' // 'cat '
	CBusAddress       KConnSuite = 'c'<<24 | 'b'<<16 | 'u'<<8 | 's' // 'cbus'
	CDevSpec          KConnSuite = 'c'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'cdev'
	CEthernetAddress  KConnSuite = 'c'<<24 | 'e'<<16 | 'n'<<8 | ' ' // 'cen '
	CFireWireAddress  KConnSuite = 'c'<<24 | 'f'<<16 | 'w'<<8 | ' ' // 'cfw '
	CIPAddress        KConnSuite = 'c'<<24 | 'i'<<16 | 'p'<<8 | ' ' // 'cip '
	CLocalTalkAddress KConnSuite = 'c'<<24 | 'l'<<16 | 't'<<8 | ' ' // 'clt '
	CSCSIAddress      KConnSuite = 'c'<<24 | 's'<<16 | 'c'<<8 | 's' // 'cscs'
	CTokenRingAddress KConnSuite = 'c'<<24 | 't'<<16 | 'o'<<8 | 'k' // 'ctok'
	CUSBAddress       KConnSuite = 'c'<<24 | 'u'<<16 | 's'<<8 | 'b' // 'cusb'
	EADB              KConnSuite = 'e'<<24 | 'a'<<16 | 'd'<<8 | 'b' // 'eadb'
	EAddressSpec      KConnSuite = 'e'<<24 | 'a'<<16 | 'd'<<8 | 's' // 'eads'
	EAnalogAudio      KConnSuite = 'e'<<24 | 'p'<<16 | 'a'<<8 | 'u' // 'epau'
	EAppleTalk        KConnSuite = 'e'<<24 | 'p'<<16 | 'a'<<8 | 't' // 'epat'
	EAudioLineIn      KConnSuite = 'e'<<24 | 'c'<<16 | 'a'<<8 | 'i' // 'ecai'
	EAudioLineOut     KConnSuite = 'e'<<24 | 'c'<<16 | 'a'<<8 | 'l' // 'ecal'
	EAudioOut         KConnSuite = 'e'<<24 | 'c'<<16 | 'a'<<8 | 'o' // 'ecao'
	EBus              KConnSuite = 'e'<<24 | 'b'<<16 | 'u'<<8 | 's' // 'ebus'
	ECDROM            KConnSuite = 'e'<<24 | 'c'<<16 | 'd'<<8 | ' ' // 'ecd '
	ECommSlot         KConnSuite = 'e'<<24 | 'c'<<16 | 'c'<<8 | 'm' // 'eccm'
	EConduit          KConnSuite = 'e'<<24 | 'c'<<16 | 'o'<<8 | 'n' // 'econ'
	EDVD              KConnSuite = 'e'<<24 | 'd'<<16 | 'v'<<8 | 'd' // 'edvd'
	EDeviceType       KConnSuite = 'e'<<24 | 'd'<<16 | 'v'<<8 | 't' // 'edvt'
	EDigitalAudio     KConnSuite = 'e'<<24 | 'p'<<16 | 'd'<<8 | 'a' // 'epda'
	EDisplay          KConnSuite = 'e'<<24 | 'd'<<16 | 'd'<<8 | 's' // 'edds'
	EEthernet         KConnSuite = 'e'<<24 | 'c'<<16 | 'e'<<8 | 'n' // 'ecen'
	EFireWire         KConnSuite = 'e'<<24 | 'c'<<16 | 'f'<<8 | 'w' // 'ecfw'
	EFloppy           KConnSuite = 'e'<<24 | 'f'<<16 | 'd'<<8 | ' ' // 'efd '
	EHD               KConnSuite = 'e'<<24 | 'h'<<16 | 'd'<<8 | ' ' // 'ehd '
	EIP               KConnSuite = 'e'<<24 | 'p'<<16 | 'i'<<8 | 'p' // 'epip'
	EIRTalk           KConnSuite = 'e'<<24 | 'p'<<16 | 'i'<<8 | 't' // 'epit'
	EInfrared         KConnSuite = 'e'<<24 | 'c'<<16 | 'i'<<8 | 'r' // 'ecir'
	EIrDA             KConnSuite = 'e'<<24 | 'p'<<16 | 'i'<<8 | 'r' // 'epir'
	EKeyboard         KConnSuite = 'e'<<24 | 'k'<<16 | 'b'<<8 | 'd' // 'ekbd'
	ELCD              KConnSuite = 'e'<<24 | 'd'<<16 | 'l'<<8 | 'c' // 'edlc'
	ELocalTalk        KConnSuite = 'e'<<24 | 'c'<<16 | 'l'<<8 | 't' // 'eclt'
	EMacIP            KConnSuite = 'e'<<24 | 'p'<<16 | 'm'<<8 | 'i' // 'epmi'
	EMacVideo         KConnSuite = 'e'<<24 | 'p'<<16 | 'm'<<8 | 'v' // 'epmv'
	EMicrophone       KConnSuite = 'e'<<24 | 'c'<<16 | 'm'<<8 | 'i' // 'ecmi'
	EModem            KConnSuite = 'e'<<24 | 'd'<<16 | 'm'<<8 | 'm' // 'edmm'
	EModemPort        KConnSuite = 'e'<<24 | 'c'<<16 | 'm'<<8 | 'p' // 'ecmp'
	EModemPrinterPort KConnSuite = 'e'<<24 | 'm'<<16 | 'p'<<8 | 'p' // 'empp'
	EMonitorOut       KConnSuite = 'e'<<24 | 'c'<<16 | 'm'<<8 | 'n' // 'ecmn'
	EMouse            KConnSuite = 'e'<<24 | 'm'<<16 | 'o'<<8 | 'u' // 'emou'
	ENuBus            KConnSuite = 'e'<<24 | 'n'<<16 | 'u'<<8 | 'b' // 'enub'
	ENuBusCard        KConnSuite = 'e'<<24 | 'd'<<16 | 'n'<<8 | 'b' // 'ednb'
	EPCIbus           KConnSuite = 'e'<<24 | 'c'<<16 | 'p'<<8 | 'i' // 'ecpi'
	EPCIcard          KConnSuite = 'e'<<24 | 'd'<<16 | 'p'<<8 | 'i' // 'edpi'
	EPCcard           KConnSuite = 'e'<<24 | 'c'<<16 | 'p'<<8 | 'c' // 'ecpc'
	EPDScard          KConnSuite = 'e'<<24 | 'p'<<16 | 'd'<<8 | 's' // 'epds'
	EPDSslot          KConnSuite = 'e'<<24 | 'c'<<16 | 'p'<<8 | 'd' // 'ecpd'
	EPPP              KConnSuite = 'e'<<24 | 'p'<<16 | 'p'<<8 | 'p' // 'eppp'
	EPointingDevice   KConnSuite = 'e'<<24 | 'd'<<16 | 'p'<<8 | 'd' // 'edpd'
	EPostScript       KConnSuite = 'e'<<24 | 'p'<<16 | 'p'<<8 | 's' // 'epps'
	EPrinter          KConnSuite = 'e'<<24 | 'd'<<16 | 'p'<<8 | 'r' // 'edpr'
	EPrinterPort      KConnSuite = 'e'<<24 | 'c'<<16 | 'p'<<8 | 'p' // 'ecpp'
	EProtocol         KConnSuite = 'e'<<24 | 'p'<<16 | 'r'<<8 | 'o' // 'epro'
	ESCSI             KConnSuite = 'e'<<24 | 'c'<<16 | 's'<<8 | 'c' // 'ecsc'
	ESVGA             KConnSuite = 'e'<<24 | 'p'<<16 | 's'<<8 | 'g' // 'epsg'
	ESerial           KConnSuite = 'e'<<24 | 'p'<<16 | 's'<<8 | 'r' // 'epsr'
	ESpeakers         KConnSuite = 'e'<<24 | 'd'<<16 | 's'<<8 | 'p' // 'edsp'
	EStorageDevice    KConnSuite = 'e'<<24 | 'd'<<16 | 's'<<8 | 't' // 'edst'
	ESvideo           KConnSuite = 'e'<<24 | 'p'<<16 | 's'<<8 | 'v' // 'epsv'
	ETokenRing        KConnSuite = 'e'<<24 | 't'<<16 | 'o'<<8 | 'k' // 'etok'
	ETrackball        KConnSuite = 'e'<<24 | 't'<<16 | 'r'<<8 | 'k' // 'etrk'
	ETrackpad         KConnSuite = 'e'<<24 | 'd'<<16 | 't'<<8 | 'p' // 'edtp'
	EUSB              KConnSuite = 'e'<<24 | 'c'<<16 | 'u'<<8 | 's' // 'ecus'
	EVideoIn          KConnSuite = 'e'<<24 | 'c'<<16 | 'v'<<8 | 'i' // 'ecvi'
	EVideoMonitor     KConnSuite = 'e'<<24 | 'd'<<16 | 'v'<<8 | 'm' // 'edvm'
	EVideoOut         KConnSuite = 'e'<<24 | 'c'<<16 | 'v'<<8 | 'o' // 'ecvo'
	KConnSuiteValue   KConnSuite = 'm'<<24 | 'a'<<16 | 'c'<<8 | 'c' // 'macc'
	PATMachine        KConnSuite = 'p'<<24 | 'a'<<16 | 't'<<8 | 'm' // 'patm'
	PATType           KConnSuite = 'p'<<24 | 'a'<<16 | 't'<<8 | 't' // 'patt'
	PATZone           KConnSuite = 'p'<<24 | 'a'<<16 | 't'<<8 | 'z' // 'patz'
	PConduit          KConnSuite = 'p'<<24 | 'c'<<16 | 'o'<<8 | 'n' // 'pcon'
	PDNS              KConnSuite = 'p'<<24 | 'd'<<16 | 'n'<<8 | 's' // 'pdns'
	PDeviceAddress    KConnSuite = 'p'<<24 | 'd'<<16 | 'v'<<8 | 'a' // 'pdva'
	PDeviceType       KConnSuite = 'p'<<24 | 'd'<<16 | 'v'<<8 | 't' // 'pdvt'
	PDottedDecimal    KConnSuite = 'p'<<24 | 'i'<<16 | 'p'<<8 | 'd' // 'pipd'
	PNetwork          KConnSuite = 'p'<<24 | 'n'<<16 | 'e'<<8 | 't' // 'pnet'
	PNode             KConnSuite = 'p'<<24 | 'n'<<16 | 'o'<<8 | 'd' // 'pnod'
	PPort             KConnSuite = 'p'<<24 | 'p'<<16 | 'o'<<8 | 'r' // 'ppor'
	PProtocol         KConnSuite = 'p'<<24 | 'p'<<16 | 'r'<<8 | 't' // 'pprt'
	PSCSIBus          KConnSuite = 'p'<<24 | 's'<<16 | 'c'<<8 | 'b' // 'pscb'
	PSCSILUN          KConnSuite = 'p'<<24 | 's'<<16 | 'l'<<8 | 'u' // 'pslu'
	PSocket           KConnSuite = 'p'<<24 | 's'<<16 | 'o'<<8 | 'c' // 'psoc'
)

func (e KConnSuite) String() string {
	switch e {
	case CADBAddress:
		return "CADBAddress"
	case CAddressSpec:
		return "CAddressSpec"
	case CAppleTalkAddress:
		return "CAppleTalkAddress"
	case CBusAddress:
		return "CBusAddress"
	case CDevSpec:
		return "CDevSpec"
	case CEthernetAddress:
		return "CEthernetAddress"
	case CFireWireAddress:
		return "CFireWireAddress"
	case CIPAddress:
		return "CIPAddress"
	case CLocalTalkAddress:
		return "CLocalTalkAddress"
	case CSCSIAddress:
		return "CSCSIAddress"
	case CTokenRingAddress:
		return "CTokenRingAddress"
	case CUSBAddress:
		return "CUSBAddress"
	case EADB:
		return "EADB"
	case EAddressSpec:
		return "EAddressSpec"
	case EAnalogAudio:
		return "EAnalogAudio"
	case EAppleTalk:
		return "EAppleTalk"
	case EAudioLineIn:
		return "EAudioLineIn"
	case EAudioLineOut:
		return "EAudioLineOut"
	case EAudioOut:
		return "EAudioOut"
	case EBus:
		return "EBus"
	case ECDROM:
		return "ECDROM"
	case ECommSlot:
		return "ECommSlot"
	case EConduit:
		return "EConduit"
	case EDVD:
		return "EDVD"
	case EDeviceType:
		return "EDeviceType"
	case EDigitalAudio:
		return "EDigitalAudio"
	case EDisplay:
		return "EDisplay"
	case EEthernet:
		return "EEthernet"
	case EFireWire:
		return "EFireWire"
	case EFloppy:
		return "EFloppy"
	case EHD:
		return "EHD"
	case EIP:
		return "EIP"
	case EIRTalk:
		return "EIRTalk"
	case EInfrared:
		return "EInfrared"
	case EIrDA:
		return "EIrDA"
	case EKeyboard:
		return "EKeyboard"
	case ELCD:
		return "ELCD"
	case ELocalTalk:
		return "ELocalTalk"
	case EMacIP:
		return "EMacIP"
	case EMacVideo:
		return "EMacVideo"
	case EMicrophone:
		return "EMicrophone"
	case EModem:
		return "EModem"
	case EModemPort:
		return "EModemPort"
	case EModemPrinterPort:
		return "EModemPrinterPort"
	case EMonitorOut:
		return "EMonitorOut"
	case EMouse:
		return "EMouse"
	case ENuBus:
		return "ENuBus"
	case ENuBusCard:
		return "ENuBusCard"
	case EPCIbus:
		return "EPCIbus"
	case EPCIcard:
		return "EPCIcard"
	case EPCcard:
		return "EPCcard"
	case EPDScard:
		return "EPDScard"
	case EPDSslot:
		return "EPDSslot"
	case EPPP:
		return "EPPP"
	case EPointingDevice:
		return "EPointingDevice"
	case EPostScript:
		return "EPostScript"
	case EPrinter:
		return "EPrinter"
	case EPrinterPort:
		return "EPrinterPort"
	case EProtocol:
		return "EProtocol"
	case ESCSI:
		return "ESCSI"
	case ESVGA:
		return "ESVGA"
	case ESerial:
		return "ESerial"
	case ESpeakers:
		return "ESpeakers"
	case EStorageDevice:
		return "EStorageDevice"
	case ESvideo:
		return "ESvideo"
	case ETokenRing:
		return "ETokenRing"
	case ETrackball:
		return "ETrackball"
	case ETrackpad:
		return "ETrackpad"
	case EUSB:
		return "EUSB"
	case EVideoIn:
		return "EVideoIn"
	case EVideoMonitor:
		return "EVideoMonitor"
	case EVideoOut:
		return "EVideoOut"
	case KConnSuiteValue:
		return "KConnSuiteValue"
	case PATMachine:
		return "PATMachine"
	case PATType:
		return "PATType"
	case PATZone:
		return "PATZone"
	case PConduit:
		return "PConduit"
	case PDNS:
		return "PDNS"
	case PDeviceAddress:
		return "PDeviceAddress"
	case PDeviceType:
		return "PDeviceType"
	case PDottedDecimal:
		return "PDottedDecimal"
	case PNetwork:
		return "PNetwork"
	case PNode:
		return "PNode"
	case PPort:
		return "PPort"
	case PProtocol:
		return "PProtocol"
	case PSCSIBus:
		return "PSCSIBus"
	case PSCSILUN:
		return "PSCSILUN"
	case PSocket:
		return "PSocket"
	default:
		return fmt.Sprintf("KConnSuite(%d)", e)
	}
}

type KContainerFolderAliasType uint

const (
	KAppPackageAliasType           KContainerFolderAliasType = 'f'<<24 | 'a'<<16 | 'p'<<8 | 'a' // 'fapa'
	KApplicationAliasType          KContainerFolderAliasType = 'a'<<24 | 'd'<<16 | 'r'<<8 | 'p' // 'adrp'
	KApplicationCPAliasType        KContainerFolderAliasType = 'a'<<24 | 'c'<<16 | 'd'<<8 | 'p' // 'acdp'
	KApplicationDAAliasType        KContainerFolderAliasType = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'p' // 'addp'
	KContainerAliasType            KContainerFolderAliasType = 'd'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'drop'
	KContainerCDROMAliasType       KContainerFolderAliasType = 'c'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'cddr'
	KContainerFloppyAliasType      KContainerFolderAliasType = 'f'<<24 | 'l'<<16 | 'p'<<8 | 'y' // 'flpy'
	KContainerFolderAliasTypeValue KContainerFolderAliasType = 'f'<<24 | 'd'<<16 | 'r'<<8 | 'p' // 'fdrp'
	KContainerHardDiskAliasType    KContainerFolderAliasType = 'h'<<24 | 'd'<<16 | 's'<<8 | 'k' // 'hdsk'
	KContainerServerAliasType      KContainerFolderAliasType = 's'<<24 | 'r'<<16 | 'v'<<8 | 'r' // 'srvr'
	KContainerTrashAliasType       KContainerFolderAliasType = 't'<<24 | 'r'<<16 | 's'<<8 | 'h' // 'trsh'
	KDesktopPrinterAliasType       KContainerFolderAliasType = 'd'<<24 | 't'<<16 | 'p'<<8 | 'a' // 'dtpa'
	KPackageAliasType              KContainerFolderAliasType = 'f'<<24 | 'p'<<16 | 'k'<<8 | 'a' // 'fpka'
)

func (e KContainerFolderAliasType) String() string {
	switch e {
	case KAppPackageAliasType:
		return "KAppPackageAliasType"
	case KApplicationAliasType:
		return "KApplicationAliasType"
	case KApplicationCPAliasType:
		return "KApplicationCPAliasType"
	case KApplicationDAAliasType:
		return "KApplicationDAAliasType"
	case KContainerAliasType:
		return "KContainerAliasType"
	case KContainerCDROMAliasType:
		return "KContainerCDROMAliasType"
	case KContainerFloppyAliasType:
		return "KContainerFloppyAliasType"
	case KContainerFolderAliasTypeValue:
		return "KContainerFolderAliasTypeValue"
	case KContainerHardDiskAliasType:
		return "KContainerHardDiskAliasType"
	case KContainerServerAliasType:
		return "KContainerServerAliasType"
	case KContainerTrashAliasType:
		return "KContainerTrashAliasType"
	case KDesktopPrinterAliasType:
		return "KDesktopPrinterAliasType"
	case KPackageAliasType:
		return "KPackageAliasType"
	default:
		return fmt.Sprintf("KContainerFolderAliasType(%d)", e)
	}
}

type KCooperativeThread uint

const (
	// Deprecated.
	KCooperativeThreadValue KCooperativeThread = 1
	// Deprecated.
	KPreemptiveThread KCooperativeThread = 2
)

func (e KCooperativeThread) String() string {
	switch e {
	case KCooperativeThreadValue:
		return "KCooperativeThreadValue"
	case KPreemptiveThread:
		return "KPreemptiveThread"
	default:
		return fmt.Sprintf("KCooperativeThread(%d)", e)
	}
}

type KCoreEndian uint

const (
	// KCoreEndianAppleEventManagerDomain: Specifies that the domain is limited to Appleevents.
	KCoreEndianAppleEventManagerDomain KCoreEndian = 'a'<<24 | 'e'<<16 | 'v'<<8 | 't' // 'aevt'
	// KCoreEndianResourceManagerDomain: Specifies that the domain is limited to theresources for a specific application.
	KCoreEndianResourceManagerDomain KCoreEndian = 'r'<<24 | 's'<<16 | 'r'<<8 | 'c' // 'rsrc'
)

func (e KCoreEndian) String() string {
	switch e {
	case KCoreEndianAppleEventManagerDomain:
		return "KCoreEndianAppleEventManagerDomain"
	case KCoreEndianResourceManagerDomain:
		return "KCoreEndianResourceManagerDomain"
	default:
		return fmt.Sprintf("KCoreEndian(%d)", e)
	}
}

type KCoreEvent uint

const (
	// KCoreEventClass: An Apple event sent by the Mac OS; applications that present a graphical interface to the user should be able to any events sent by the Mac OS that apply to the application.
	KCoreEventClass KCoreEvent = 'a'<<24 | 'e'<<16 | 'v'<<8 | 't' // 'aevt'
)

func (e KCoreEvent) String() string {
	switch e {
	case KCoreEventClass:
		return "KCoreEventClass"
	default:
		return fmt.Sprintf("KCoreEvent(%d)", e)
	}
}

type KCreateFolder uint

const (
	// Deprecated.
	KCreateFolderValue KCreateFolder = 1
	// Deprecated.
	KDontCreateFolder KCreateFolder = 0
)

func (e KCreateFolder) String() string {
	switch e {
	case KCreateFolderValue:
		return "KCreateFolderValue"
	case KDontCreateFolder:
		return "KDontCreateFolder"
	default:
		return fmt.Sprintf("KCreateFolder(%d)", e)
	}
}

type KCreateFolderAtBoot uint

const (
	// Deprecated.
	KCreateFolderAtBootValue KCreateFolderAtBoot = 2
	// Deprecated.
	KCreateFolderAtBootBit KCreateFolderAtBoot = 1
	// Deprecated.
	KFolderCreatedAdminPrivs KCreateFolderAtBoot = 16
	// Deprecated.
	KFolderCreatedAdminPrivsBit KCreateFolderAtBoot = 4
	// Deprecated.
	KFolderCreatedInvisible KCreateFolderAtBoot = 4
	// Deprecated.
	KFolderCreatedInvisibleBit KCreateFolderAtBoot = 2
	// Deprecated.
	KFolderCreatedNameLocked KCreateFolderAtBoot = 8
	// Deprecated.
	KFolderCreatedNameLockedBit KCreateFolderAtBoot = 3
)

func (e KCreateFolderAtBoot) String() string {
	switch e {
	case KCreateFolderAtBootValue:
		return "KCreateFolderAtBootValue"
	case KCreateFolderAtBootBit:
		return "KCreateFolderAtBootBit"
	case KFolderCreatedAdminPrivs:
		return "KFolderCreatedAdminPrivs"
	case KFolderCreatedAdminPrivsBit:
		return "KFolderCreatedAdminPrivsBit"
	case KFolderCreatedNameLocked:
		return "KFolderCreatedNameLocked"
	case KFolderCreatedNameLockedBit:
		return "KFolderCreatedNameLockedBit"
	default:
		return fmt.Sprintf("KCreateFolderAtBoot(%d)", e)
	}
}

type KCurrentMixedModeState uint

const (
	// Deprecated.
	KCurrentMixedModeStateRecord KCurrentMixedModeState = 1
)

func (e KCurrentMixedModeState) String() string {
	switch e {
	case KCurrentMixedModeStateRecord:
		return "KCurrentMixedModeStateRecord"
	default:
		return fmt.Sprintf("KCurrentMixedModeState(%d)", e)
	}
}

type KCurrentUserFolder uint

const (
	KCurrentUserFolderLocation KCurrentUserFolder = 'c'<<24 | 'u'<<16 | 's'<<8 | 'f' // 'cusf'
)

func (e KCurrentUserFolder) String() string {
	switch e {
	case KCurrentUserFolderLocation:
		return "KCurrentUserFolderLocation"
	default:
		return fmt.Sprintf("KCurrentUserFolder(%d)", e)
	}
}

type KCustomBadgeResource int

const (
	KCustomBadgeResourceID      KCustomBadgeResource = -16455
	KCustomBadgeResourceType    KCustomBadgeResource = 'b'<<24 | 'a'<<16 | 'd'<<8 | 'g' // 'badg'
	KCustomBadgeResourceVersion KCustomBadgeResource = 0
)

func (e KCustomBadgeResource) String() string {
	switch e {
	case KCustomBadgeResourceID:
		return "KCustomBadgeResourceID"
	case KCustomBadgeResourceType:
		return "KCustomBadgeResourceType"
	case KCustomBadgeResourceVersion:
		return "KCustomBadgeResourceVersion"
	default:
		return fmt.Sprintf("KCustomBadgeResource(%d)", e)
	}
}

type KCustomIcon int

const (
	KCustomIconResource KCustomIcon = -16455
)

func (e KCustomIcon) String() string {
	switch e {
	case KCustomIconResource:
		return "KCustomIconResource"
	default:
		return fmt.Sprintf("KCustomIcon(%d)", e)
	}
}

type KDMGenErr int

const (
	KDMCantBlock                   KDMGenErr = -6224
	KDMDisplayAlreadyInstalledErr  KDMGenErr = -6230
	KDMDisplayNotFoundErr          KDMGenErr = -6229
	KDMDriverNotDisplayMgrAwareErr KDMGenErr = -6228
	KDMFoundErr                    KDMGenErr = -6232
	KDMGenErrValue                 KDMGenErr = -6220
	KDMMainDisplayCannotMoveErr    KDMGenErr = -6231
	KDMMirroringBlocked            KDMGenErr = -6223
	KDMMirroringNotOn              KDMGenErr = -6225
	KDMMirroringOnAlready          KDMGenErr = -6221
	KDMNoDeviceTableclothErr       KDMGenErr = -6231
	KDMNotFoundErr                 KDMGenErr = -6229
	KDMSWNotInitializedErr         KDMGenErr = -6227
	KDMWrongNumberOfDisplays       KDMGenErr = -6222
	KSysSWTooOld                   KDMGenErr = -6226
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

type KDOSJapanese uint

const (
	KDOSJapanesePalmVariant     KDOSJapanese = 1
	KDOSJapaneseStandardVariant KDOSJapanese = 0
)

func (e KDOSJapanese) String() string {
	switch e {
	case KDOSJapanesePalmVariant:
		return "KDOSJapanesePalmVariant"
	case KDOSJapaneseStandardVariant:
		return "KDOSJapaneseStandardVariant"
	default:
		return fmt.Sprintf("KDOSJapanese(%d)", e)
	}
}

type KDSp int

const (
	KDSpConfirmSwitchWarning      KDSp = -30448
	KDSpContextAlreadyReservedErr KDSp = -30444
	KDSpContextNotFoundErr        KDSp = -30446
	KDSpContextNotReservedErr     KDSp = -30445
	KDSpFrameRateNotReadyErr      KDSp = -30447
	KDSpInternalErr               KDSp = -30449
	KDSpInvalidAttributesErr      KDSp = -30443
	KDSpInvalidContextErr         KDSp = -30442
	KDSpNotInitializedErr         KDSp = -30440
	KDSpStereoContextErr          KDSp = -30450
	KDSpSystemSWTooOldErr         KDSp = -30441
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
	KDTPAbortJobErr  KDTP = 128
	KDTPHoldJobErr   KDTP = -4200
	KDTPStopQueueErr KDTP = -4201
	KDTPTryAgainErr  KDTP = -4202
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

type KDesktopFolderType uint

const (
	// Deprecated.
	KApplicationsFolderType KDesktopFolderType = 'a'<<24 | 'p'<<16 | 'p'<<8 | 's' // 'apps'
	// Deprecated.
	KChewableItemsFolderType KDesktopFolderType = 'f'<<24 | 'l'<<16 | 'n'<<8 | 't' // 'flnt'
	// Deprecated.
	KCurrentUserFolderType KDesktopFolderType = 'c'<<24 | 'u'<<16 | 's'<<8 | 'r' // 'cusr'
	// Deprecated.
	KDesktopFolderTypeValue KDesktopFolderType = 'd'<<24 | 'e'<<16 | 's'<<8 | 'k' // 'desk'
	// Deprecated.
	KDomainLibraryFolderType KDesktopFolderType = 'd'<<24 | 'l'<<16 | 'i'<<8 | 'b' // 'dlib'
	// Deprecated.
	KDomainTopLevelFolderType KDesktopFolderType = 'd'<<24 | 't'<<16 | 'o'<<8 | 'p' // 'dtop'
	// Deprecated.
	KFontsFolderType KDesktopFolderType = 'f'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'font'
	// Deprecated.
	KPreferencesFolderType KDesktopFolderType = 'p'<<24 | 'r'<<16 | 'e'<<8 | 'f' // 'pref'
	// Deprecated.
	KSharedUserDataFolderType KDesktopFolderType = 's'<<24 | 'd'<<16 | 'a'<<8 | 't' // 'sdat'
	// Deprecated.
	KSystemPreferencesFolderType KDesktopFolderType = 's'<<24 | 'p'<<16 | 'r'<<8 | 'f' // 'sprf'
	// Deprecated.
	KTemporaryFolderType KDesktopFolderType = 't'<<24 | 'e'<<16 | 'm'<<8 | 'p' // 'temp'
	// Deprecated.
	KTemporaryItemsInCacheDataFolderType KDesktopFolderType = 'v'<<24 | 't'<<16 | 'm'<<8 | 'p' // 'vtmp'
	// Deprecated.
	KTrashFolderType KDesktopFolderType = 't'<<24 | 'r'<<16 | 's'<<8 | 'h' // 'trsh'
	// Deprecated.
	KUsersFolderType KDesktopFolderType = 'u'<<24 | 's'<<16 | 'r'<<8 | 's' // 'usrs'
	// Deprecated.
	KVolumeRootFolderType KDesktopFolderType = 'r'<<24 | 'o'<<16 | 'o'<<8 | 't' // 'root'
	// Deprecated.
	KWhereToEmptyTrashFolderType KDesktopFolderType = 'e'<<24 | 'm'<<16 | 'p'<<8 | 't' // 'empt'
)

func (e KDesktopFolderType) String() string {
	switch e {
	case KApplicationsFolderType:
		return "KApplicationsFolderType"
	case KChewableItemsFolderType:
		return "KChewableItemsFolderType"
	case KCurrentUserFolderType:
		return "KCurrentUserFolderType"
	case KDesktopFolderTypeValue:
		return "KDesktopFolderTypeValue"
	case KDomainLibraryFolderType:
		return "KDomainLibraryFolderType"
	case KDomainTopLevelFolderType:
		return "KDomainTopLevelFolderType"
	case KFontsFolderType:
		return "KFontsFolderType"
	case KPreferencesFolderType:
		return "KPreferencesFolderType"
	case KSharedUserDataFolderType:
		return "KSharedUserDataFolderType"
	case KSystemPreferencesFolderType:
		return "KSystemPreferencesFolderType"
	case KTemporaryFolderType:
		return "KTemporaryFolderType"
	case KTemporaryItemsInCacheDataFolderType:
		return "KTemporaryItemsInCacheDataFolderType"
	case KTrashFolderType:
		return "KTrashFolderType"
	case KUsersFolderType:
		return "KUsersFolderType"
	case KVolumeRootFolderType:
		return "KVolumeRootFolderType"
	case KWhereToEmptyTrashFolderType:
		return "KWhereToEmptyTrashFolderType"
	default:
		return fmt.Sprintf("KDesktopFolderType(%d)", e)
	}
}

type KDesktopIconResource int

const (
	KDesktopIconResourceValue       KDesktopIconResource = -3992
	KGenericFileServerIconResource  KDesktopIconResource = -3972
	KGenericHardDiskIconResource    KDesktopIconResource = -3995
	KGenericMoverObjectIconResource KDesktopIconResource = -3969
	KGenericSuitcaseIconResource    KDesktopIconResource = -3970
	KOpenFolderIconResource         KDesktopIconResource = -3997
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

type KDeveloper uint

const (
	// Deprecated.
	KDeveloperApplicationsFolderType KDeveloper = 'd'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'dapp'
	// Deprecated.
	KDeveloperDocsFolderType KDeveloper = 'd'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'ddoc'
	// Deprecated.
	KDeveloperFolderType KDeveloper = 'd'<<24 | 'e'<<16 | 'v'<<8 | 'f' // 'devf'
	// Deprecated.
	KDeveloperHelpFolderType KDeveloper = 'd'<<24 | 'e'<<16 | 'v'<<8 | 'h' // 'devh'
)

func (e KDeveloper) String() string {
	switch e {
	case KDeveloperApplicationsFolderType:
		return "KDeveloperApplicationsFolderType"
	case KDeveloperDocsFolderType:
		return "KDeveloperDocsFolderType"
	case KDeveloperFolderType:
		return "KDeveloperFolderType"
	case KDeveloperHelpFolderType:
		return "KDeveloperHelpFolderType"
	default:
		return fmt.Sprintf("KDeveloper(%d)", e)
	}
}

type KDictionariesFolderType uint

const (
	KDictionariesFolderTypeValue KDictionariesFolderType = 'd'<<24 | 'i'<<16 | 'c'<<8 | 't' // 'dict'
	KLogsFolderType              KDictionariesFolderType = 'l'<<24 | 'o'<<16 | 'g'<<8 | 's' // 'logs'
	KPreferencePanesFolderType   KDictionariesFolderType = 'p'<<24 | 'p'<<16 | 'a'<<8 | 'n' // 'ppan'
)

func (e KDictionariesFolderType) String() string {
	switch e {
	case KDictionariesFolderTypeValue:
		return "KDictionariesFolderTypeValue"
	case KLogsFolderType:
		return "KLogsFolderType"
	case KPreferencePanesFolderType:
		return "KPreferencePanesFolderType"
	default:
		return fmt.Sprintf("KDictionariesFolderType(%d)", e)
	}
}

type KDocumentsFolderType uint

const (
	// Deprecated.
	KDocumentsFolderTypeValue KDocumentsFolderType = 'd'<<24 | 'o'<<16 | 'c'<<8 | 's' // 'docs'
	// Deprecated.
	KInternetSitesFolderType KDocumentsFolderType = 's'<<24 | 'i'<<16 | 't'<<8 | 'e' // 'site'
	// Deprecated.
	KMovieDocumentsFolderType KDocumentsFolderType = 'm'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'mdoc'
	// Deprecated.
	KMusicDocumentsFolderType KDocumentsFolderType = 3043258211
	// Deprecated.
	KPictureDocumentsFolderType KDocumentsFolderType = 'p'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'pdoc'
	// Deprecated.
	KPublicFolderType KDocumentsFolderType = 'p'<<24 | 'u'<<16 | 'b'<<8 | 'b' // 'pubb'
)

func (e KDocumentsFolderType) String() string {
	switch e {
	case KDocumentsFolderTypeValue:
		return "KDocumentsFolderTypeValue"
	case KInternetSitesFolderType:
		return "KInternetSitesFolderType"
	case KMovieDocumentsFolderType:
		return "KMovieDocumentsFolderType"
	case KMusicDocumentsFolderType:
		return "KMusicDocumentsFolderType"
	case KPictureDocumentsFolderType:
		return "KPictureDocumentsFolderType"
	case KPublicFolderType:
		return "KPublicFolderType"
	default:
		return fmt.Sprintf("KDocumentsFolderType(%d)", e)
	}
}

type KDropBoxFolder uint

const (
	// Deprecated.
	KDropBoxFolderType KDropBoxFolder = 'd'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'drop'
)

func (e KDropBoxFolder) String() string {
	switch e {
	case KDropBoxFolderType:
		return "KDropBoxFolderType"
	default:
		return fmt.Sprintf("KDropBoxFolder(%d)", e)
	}
}

type KDuration int

const (
	// Deprecated.
	KDurationForever KDuration = 2147483647
	// Deprecated.
	KDurationImmediate KDuration = 0
	// Deprecated.
	KDurationMicrosecond KDuration = -1
	// Deprecated.
	KDurationMillisecond KDuration = 1
)

func (e KDuration) String() string {
	switch e {
	case KDurationForever:
		return "KDurationForever"
	case KDurationImmediate:
		return "KDurationImmediate"
	case KDurationMicrosecond:
		return "KDurationMicrosecond"
	case KDurationMillisecond:
		return "KDurationMillisecond"
	default:
		return fmt.Sprintf("KDuration(%d)", e)
	}
}

type KExportedFolderAliasType uint

const (
	KDropFolderAliasType          KExportedFolderAliasType = 'f'<<24 | 'a'<<16 | 'd'<<8 | 'r' // 'fadr'
	KExportedFolderAliasTypeValue KExportedFolderAliasType = 'f'<<24 | 'a'<<16 | 'e'<<8 | 't' // 'faet'
	KMountedFolderAliasType       KExportedFolderAliasType = 'f'<<24 | 'a'<<16 | 'm'<<8 | 'n' // 'famn'
	KSharedFolderAliasType        KExportedFolderAliasType = 'f'<<24 | 'a'<<16 | 's'<<8 | 'h' // 'fash'
)

func (e KExportedFolderAliasType) String() string {
	switch e {
	case KDropFolderAliasType:
		return "KDropFolderAliasType"
	case KExportedFolderAliasTypeValue:
		return "KExportedFolderAliasTypeValue"
	case KMountedFolderAliasType:
		return "KMountedFolderAliasType"
	case KSharedFolderAliasType:
		return "KSharedFolderAliasType"
	default:
		return fmt.Sprintf("KExportedFolderAliasType(%d)", e)
	}
}

type KExtended uint

const (
	KExtendedFlagHasCustomBadge KExtended = 256
	KExtendedFlagHasRoutingInfo KExtended = 4
	KExtendedFlagObjectIsBusy   KExtended = 128
	KExtendedFlagsAreInvalid    KExtended = 32768
)

func (e KExtended) String() string {
	switch e {
	case KExtendedFlagHasCustomBadge:
		return "KExtendedFlagHasCustomBadge"
	case KExtendedFlagHasRoutingInfo:
		return "KExtendedFlagHasRoutingInfo"
	case KExtendedFlagObjectIsBusy:
		return "KExtendedFlagObjectIsBusy"
	case KExtendedFlagsAreInvalid:
		return "KExtendedFlagsAreInvalid"
	default:
		return fmt.Sprintf("KExtended(%d)", e)
	}
}

type KFAServerApp uint

const (
	KDoFolderActionEvent     KFAServerApp = 'f'<<24 | 'o'<<16 | 'l'<<8 | 'a' // 'fola'
	KFAAttachCommand         KFAServerApp = 'a'<<24 | 't'<<16 | 'f'<<8 | 'a' // 'atfa'
	KFAEditCommand           KFAServerApp = 'e'<<24 | 'd'<<16 | 'f'<<8 | 'a' // 'edfa'
	KFAFileParam             KFAServerApp = 'f'<<24 | 'a'<<16 | 'a'<<8 | 'l' // 'faal'
	KFAIndexParam            KFAServerApp = 'i'<<24 | 'n'<<16 | 'd'<<8 | 'x' // 'indx'
	KFARemoveCommand         KFAServerApp = 'r'<<24 | 'm'<<16 | 'f'<<8 | 'a' // 'rmfa'
	KFAServerAppValue        KFAServerApp = 's'<<24 | 's'<<16 | 'r'<<8 | 'v' // 'ssrv'
	KFASuiteCode             KFAServerApp = 'f'<<24 | 'a'<<16 | 'c'<<8 | 'o' // 'faco'
	KFolderActionCode        KFAServerApp = 'a'<<24 | 'c'<<16 | 't'<<8 | 'n' // 'actn'
	KFolderClosedEvent       KFAServerApp = 'f'<<24 | 'c'<<16 | 'l'<<8 | 'o' // 'fclo'
	KFolderItemsAddedEvent   KFAServerApp = 'f'<<24 | 'g'<<16 | 'e'<<8 | 't' // 'fget'
	KFolderItemsRemovedEvent KFAServerApp = 'f'<<24 | 'l'<<16 | 'o'<<8 | 's' // 'flos'
	KFolderOpenedEvent       KFAServerApp = 'f'<<24 | 'o'<<16 | 'p'<<8 | 'n' // 'fopn'
	KFolderWindowMovedEvent  KFAServerApp = 'f'<<24 | 's'<<16 | 'i'<<8 | 'z' // 'fsiz'
	KItemList                KFAServerApp = 'f'<<24 | 'l'<<16 | 's'<<8 | 't' // 'flst'
	KNewSizeParameter        KFAServerApp = 'f'<<24 | 'n'<<16 | 's'<<8 | 'z' // 'fnsz'
)

func (e KFAServerApp) String() string {
	switch e {
	case KDoFolderActionEvent:
		return "KDoFolderActionEvent"
	case KFAAttachCommand:
		return "KFAAttachCommand"
	case KFAEditCommand:
		return "KFAEditCommand"
	case KFAFileParam:
		return "KFAFileParam"
	case KFAIndexParam:
		return "KFAIndexParam"
	case KFARemoveCommand:
		return "KFARemoveCommand"
	case KFAServerAppValue:
		return "KFAServerAppValue"
	case KFASuiteCode:
		return "KFASuiteCode"
	case KFolderActionCode:
		return "KFolderActionCode"
	case KFolderClosedEvent:
		return "KFolderClosedEvent"
	case KFolderItemsAddedEvent:
		return "KFolderItemsAddedEvent"
	case KFolderItemsRemovedEvent:
		return "KFolderItemsRemovedEvent"
	case KFolderOpenedEvent:
		return "KFolderOpenedEvent"
	case KFolderWindowMovedEvent:
		return "KFolderWindowMovedEvent"
	case KItemList:
		return "KItemList"
	case KNewSizeParameter:
		return "KNewSizeParameter"
	default:
		return fmt.Sprintf("KFAServerApp(%d)", e)
	}
}

type KFB int

const (
	KFBCaccessCanceled        KFB = -30521
	KFBCaccessorStoreFailed   KFB = -30515
	KFBCaddDocFailed          KFB = -30516
	KFBCallocFailed           KFB = -30502
	KFBCanalysisNotAvailable  KFB = -30527
	KFBCbadIndexFile          KFB = -30505
	KFBCbadIndexFileVersion   KFB = -30528
	KFBCbadParam              KFB = -30503
	KFBCbadSearchSession      KFB = -30531
	KFBCcommitFailed          KFB = -30509
	KFBCcompactionFailed      KFB = -30506
	KFBCdeletionFailed        KFB = -30510
	KFBCfileNotIndexed        KFB = -30504
	KFBCflushFailed           KFB = -30517
	KFBCillegalSessionChange  KFB = -30526
	KFBCindexCreationFailed   KFB = -30514
	KFBCindexDiskIOFailed     KFB = -30530
	KFBCindexFileDestroyed    KFB = -30522
	KFBCindexNotAvailable     KFB = -30523
	KFBCindexNotFound         KFB = -30518
	KFBCindexingCanceled      KFB = -30520
	KFBCindexingFailed        KFB = -30508
	KFBCmergingFailed         KFB = -30513
	KFBCmoveFailed            KFB = -30511
	KFBCnoIndexesFound        KFB = -30501
	KFBCnoSearchSession       KFB = -30519
	KFBCnoSuchHit             KFB = -30532
	KFBCsearchFailed          KFB = -30524
	KFBCsomeFilesNotIndexed   KFB = -30525
	KFBCsummarizationCanceled KFB = -30529
	KFBCtokenizationFailed    KFB = -30512
	KFBCvTwinExceptionErr     KFB = -30500
	KFBCvalidationFailed      KFB = -30507
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
	KFMFontContainerAccessErr    KFM = -985
	KFMFontTableAccessErr        KFM = -984
	KFMInvalidFontErr            KFM = -982
	KFMInvalidFontFamilyErr      KFM = -981
	KFMIterationCompleted        KFM = -980
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

type KFN uint

const (
	KFNNoImplicitAllSubscription KFN = 1
	KFNNotifyInBackground        KFN = 2
)

func (e KFN) String() string {
	switch e {
	case KFNNoImplicitAllSubscription:
		return "KFNNoImplicitAllSubscription"
	case KFNNotifyInBackground:
		return "KFNNotifyInBackground"
	default:
		return fmt.Sprintf("KFN(%d)", e)
	}
}

type KFNDirectoryModified uint

const (
	KFNDirectoryModifiedMessage KFNDirectoryModified = 1
)

func (e KFNDirectoryModified) String() string {
	switch e {
	case KFNDirectoryModifiedMessage:
		return "KFNDirectoryModifiedMessage"
	default:
		return fmt.Sprintf("KFNDirectoryModified(%d)", e)
	}
}

type KFNS int

const (
	KFNSBadFlattenedSizeErr    KFNS = -29587
	KFNSBadProfileVersionErr   KFNS = -29583
	KFNSBadReferenceVersionErr KFNS = -29581
	KFNSDuplicateReferenceErr  KFNS = -29584
	KFNSInsufficientDataErr    KFNS = -29586
	KFNSInvalidProfileErr      KFNS = -29582
	KFNSInvalidReferenceErr    KFNS = -29580
	KFNSMismatchErr            KFNS = -29585
	KFNSNameNotFoundErr        KFNS = -29589
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

type KFS uint

const (
	KFSAllowConcurrentAsyncIOBit  KFS = 3
	KFSAllowConcurrentAsyncIOMask KFS = 8
	KFSForceReadBit               KFS = 6
	KFSForceReadMask              KFS = 64
	KFSNewLineBit                 KFS = 7
	KFSNewLineCharMask            KFS = 65280
	KFSNewLineMask                KFS = 128
	KFSNoCacheBit                 KFS = 5
	KFSNoCacheMask                KFS = 32
	KFSPleaseCacheBit             KFS = 4
	KFSPleaseCacheMask            KFS = 16
	KFSRdVerifyBit                KFS = 6
	KFSRdVerifyMask               KFS = 64
)

func (e KFS) String() string {
	switch e {
	case KFSAllowConcurrentAsyncIOBit:
		return "KFSAllowConcurrentAsyncIOBit"
	case KFSAllowConcurrentAsyncIOMask:
		return "KFSAllowConcurrentAsyncIOMask"
	case KFSForceReadBit:
		return "KFSForceReadBit"
	case KFSForceReadMask:
		return "KFSForceReadMask"
	case KFSNewLineBit:
		return "KFSNewLineBit"
	case KFSNewLineCharMask:
		return "KFSNewLineCharMask"
	case KFSNewLineMask:
		return "KFSNewLineMask"
	case KFSNoCacheBit:
		return "KFSNoCacheBit"
	case KFSNoCacheMask:
		return "KFSNoCacheMask"
	case KFSPleaseCacheBit:
		return "KFSPleaseCacheBit"
	case KFSPleaseCacheMask:
		return "KFSPleaseCacheMask"
	default:
		return fmt.Sprintf("KFS(%d)", e)
	}
}

type KFSAliasInfo uint

const (
	// Deprecated.
	KFSAliasInfoFSInfo KFSAliasInfo = 32
	// Deprecated.
	KFSAliasInfoFinderInfo KFSAliasInfo = 4
	// Deprecated.
	KFSAliasInfoIDs KFSAliasInfo = 16
	// Deprecated.
	KFSAliasInfoIsDirectory KFSAliasInfo = 8
	// Deprecated.
	KFSAliasInfoNone KFSAliasInfo = 0
	// Deprecated.
	KFSAliasInfoTargetCreateDate KFSAliasInfo = 2
	// Deprecated.
	KFSAliasInfoVolumeCreateDate KFSAliasInfo = 1
	// Deprecated.
	KFSAliasInfoVolumeFlags KFSAliasInfo = 64
)

func (e KFSAliasInfo) String() string {
	switch e {
	case KFSAliasInfoFSInfo:
		return "KFSAliasInfoFSInfo"
	case KFSAliasInfoFinderInfo:
		return "KFSAliasInfoFinderInfo"
	case KFSAliasInfoIDs:
		return "KFSAliasInfoIDs"
	case KFSAliasInfoIsDirectory:
		return "KFSAliasInfoIsDirectory"
	case KFSAliasInfoNone:
		return "KFSAliasInfoNone"
	case KFSAliasInfoTargetCreateDate:
		return "KFSAliasInfoTargetCreateDate"
	case KFSAliasInfoVolumeCreateDate:
		return "KFSAliasInfoVolumeCreateDate"
	case KFSAliasInfoVolumeFlags:
		return "KFSAliasInfoVolumeFlags"
	default:
		return fmt.Sprintf("KFSAliasInfo(%d)", e)
	}
}

type KFSAlloc uint

const (
	KFSAllocAllOrNothingMask KFSAlloc = 1
	KFSAllocContiguousMask   KFSAlloc = 2
	KFSAllocDefaultFlags     KFSAlloc = 0
	KFSAllocNoRoundUpMask    KFSAlloc = 4
	KFSAllocReservedMask     KFSAlloc = 65528
)

func (e KFSAlloc) String() string {
	switch e {
	case KFSAllocAllOrNothingMask:
		return "KFSAllocAllOrNothingMask"
	case KFSAllocContiguousMask:
		return "KFSAllocContiguousMask"
	case KFSAllocDefaultFlags:
		return "KFSAllocDefaultFlags"
	case KFSAllocNoRoundUpMask:
		return "KFSAllocNoRoundUpMask"
	case KFSAllocReservedMask:
		return "KFSAllocReservedMask"
	default:
		return fmt.Sprintf("KFSAlloc(%d)", e)
	}
}

type KFSCatInfo uint

const (
	KFSCatInfoAccessDate        KFSCatInfo = 256
	KFSCatInfoAllDates          KFSCatInfo = 992
	KFSCatInfoAttrMod           KFSCatInfo = 128
	KFSCatInfoBackupDate        KFSCatInfo = 512
	KFSCatInfoContentMod        KFSCatInfo = 64
	KFSCatInfoCreateDate        KFSCatInfo = 32
	KFSCatInfoDataSizes         KFSCatInfo = 16384
	KFSCatInfoFSFileSecurityRef KFSCatInfo = 4194304
	KFSCatInfoFinderInfo        KFSCatInfo = 2048
	KFSCatInfoFinderXInfo       KFSCatInfo = 4096
	KFSCatInfoGettableInfo      KFSCatInfo = 262143
	KFSCatInfoNodeFlags         KFSCatInfo = 2
	KFSCatInfoNodeID            KFSCatInfo = 16
	KFSCatInfoNone              KFSCatInfo = 0
	KFSCatInfoParentDirID       KFSCatInfo = 8
	KFSCatInfoPermissions       KFSCatInfo = 1024
	KFSCatInfoReserved          KFSCatInfo = 0
	KFSCatInfoRsrcSizes         KFSCatInfo = 32768
	KFSCatInfoSetOwnership      KFSCatInfo = 1048576
	KFSCatInfoSettableInfo      KFSCatInfo = 8163
	KFSCatInfoSharingFlags      KFSCatInfo = 65536
	KFSCatInfoTextEncoding      KFSCatInfo = 1
	KFSCatInfoUserAccess        KFSCatInfo = 524288
	KFSCatInfoUserPrivs         KFSCatInfo = 131072
	KFSCatInfoValence           KFSCatInfo = 8192
	KFSCatInfoVolume            KFSCatInfo = 4
)

func (e KFSCatInfo) String() string {
	switch e {
	case KFSCatInfoAccessDate:
		return "KFSCatInfoAccessDate"
	case KFSCatInfoAllDates:
		return "KFSCatInfoAllDates"
	case KFSCatInfoAttrMod:
		return "KFSCatInfoAttrMod"
	case KFSCatInfoBackupDate:
		return "KFSCatInfoBackupDate"
	case KFSCatInfoContentMod:
		return "KFSCatInfoContentMod"
	case KFSCatInfoCreateDate:
		return "KFSCatInfoCreateDate"
	case KFSCatInfoDataSizes:
		return "KFSCatInfoDataSizes"
	case KFSCatInfoFSFileSecurityRef:
		return "KFSCatInfoFSFileSecurityRef"
	case KFSCatInfoFinderInfo:
		return "KFSCatInfoFinderInfo"
	case KFSCatInfoFinderXInfo:
		return "KFSCatInfoFinderXInfo"
	case KFSCatInfoGettableInfo:
		return "KFSCatInfoGettableInfo"
	case KFSCatInfoNodeFlags:
		return "KFSCatInfoNodeFlags"
	case KFSCatInfoNodeID:
		return "KFSCatInfoNodeID"
	case KFSCatInfoNone:
		return "KFSCatInfoNone"
	case KFSCatInfoParentDirID:
		return "KFSCatInfoParentDirID"
	case KFSCatInfoPermissions:
		return "KFSCatInfoPermissions"
	case KFSCatInfoRsrcSizes:
		return "KFSCatInfoRsrcSizes"
	case KFSCatInfoSetOwnership:
		return "KFSCatInfoSetOwnership"
	case KFSCatInfoSettableInfo:
		return "KFSCatInfoSettableInfo"
	case KFSCatInfoSharingFlags:
		return "KFSCatInfoSharingFlags"
	case KFSCatInfoTextEncoding:
		return "KFSCatInfoTextEncoding"
	case KFSCatInfoUserAccess:
		return "KFSCatInfoUserAccess"
	case KFSCatInfoUserPrivs:
		return "KFSCatInfoUserPrivs"
	case KFSCatInfoValence:
		return "KFSCatInfoValence"
	case KFSCatInfoVolume:
		return "KFSCatInfoVolume"
	default:
		return fmt.Sprintf("KFSCatInfo(%d)", e)
	}
}

type KFSEjectVolumeForce uint

const (
	KFSEjectVolumeForceEject KFSEjectVolumeForce = 1
)

func (e KFSEjectVolumeForce) String() string {
	switch e {
	case KFSEjectVolumeForceEject:
		return "KFSEjectVolumeForceEject"
	default:
		return fmt.Sprintf("KFSEjectVolumeForce(%d)", e)
	}
}

type KFSEventStreamCreate uint

const (
	// KFSEventStreamCreateFlagFileEvents: # Discussion
	KFSEventStreamCreateFlagFileEvents  KFSEventStreamCreate = 16
	KFSEventStreamCreateFlagFullHistory KFSEventStreamCreate = 128
	// KFSEventStreamCreateFlagIgnoreSelf: # Discussion
	KFSEventStreamCreateFlagIgnoreSelf KFSEventStreamCreate = 8
	KFSEventStreamCreateFlagMarkSelf   KFSEventStreamCreate = 32
	// KFSEventStreamCreateFlagNoDefer: # Discussion
	KFSEventStreamCreateFlagNoDefer KFSEventStreamCreate = 2
	// KFSEventStreamCreateFlagNone: # Discussion
	KFSEventStreamCreateFlagNone KFSEventStreamCreate = 0
	// KFSEventStreamCreateFlagUseCFTypes: # Discussion
	KFSEventStreamCreateFlagUseCFTypes      KFSEventStreamCreate = 1
	KFSEventStreamCreateFlagUseExtendedData KFSEventStreamCreate = 64
	// KFSEventStreamCreateFlagWatchRoot: # Discussion
	KFSEventStreamCreateFlagWatchRoot KFSEventStreamCreate = 4
	KFSEventStreamCreateWithDocID     KFSEventStreamCreate = 256
)

func (e KFSEventStreamCreate) String() string {
	switch e {
	case KFSEventStreamCreateFlagFileEvents:
		return "KFSEventStreamCreateFlagFileEvents"
	case KFSEventStreamCreateFlagFullHistory:
		return "KFSEventStreamCreateFlagFullHistory"
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
	case KFSEventStreamCreateFlagUseExtendedData:
		return "KFSEventStreamCreateFlagUseExtendedData"
	case KFSEventStreamCreateFlagWatchRoot:
		return "KFSEventStreamCreateFlagWatchRoot"
	case KFSEventStreamCreateWithDocID:
		return "KFSEventStreamCreateWithDocID"
	default:
		return fmt.Sprintf("KFSEventStreamCreate(%d)", e)
	}
}

type KFSEventStreamEventFlag uint

const (
	// KFSEventStreamEventFlagEventIdsWrapped: # Discussion
	KFSEventStreamEventFlagEventIdsWrapped KFSEventStreamEventFlag = 8
	// KFSEventStreamEventFlagHistoryDone: # Discussion
	KFSEventStreamEventFlagHistoryDone        KFSEventStreamEventFlag = 16
	KFSEventStreamEventFlagItemChangeOwner    KFSEventStreamEventFlag = 16384
	KFSEventStreamEventFlagItemCloned         KFSEventStreamEventFlag = 4194304
	KFSEventStreamEventFlagItemCreated        KFSEventStreamEventFlag = 256
	KFSEventStreamEventFlagItemFinderInfoMod  KFSEventStreamEventFlag = 8192
	KFSEventStreamEventFlagItemInodeMetaMod   KFSEventStreamEventFlag = 1024
	KFSEventStreamEventFlagItemIsDir          KFSEventStreamEventFlag = 131072
	KFSEventStreamEventFlagItemIsFile         KFSEventStreamEventFlag = 65536
	KFSEventStreamEventFlagItemIsHardlink     KFSEventStreamEventFlag = 1048576
	KFSEventStreamEventFlagItemIsLastHardlink KFSEventStreamEventFlag = 2097152
	KFSEventStreamEventFlagItemIsSymlink      KFSEventStreamEventFlag = 262144
	KFSEventStreamEventFlagItemModified       KFSEventStreamEventFlag = 4096
	KFSEventStreamEventFlagItemRemoved        KFSEventStreamEventFlag = 512
	KFSEventStreamEventFlagItemRenamed        KFSEventStreamEventFlag = 2048
	KFSEventStreamEventFlagItemXattrMod       KFSEventStreamEventFlag = 32768
	KFSEventStreamEventFlagKernelDropped      KFSEventStreamEventFlag = 4
	// KFSEventStreamEventFlagMount: # Discussion
	KFSEventStreamEventFlagMount KFSEventStreamEventFlag = 64
	// KFSEventStreamEventFlagMustScanSubDirs: # Discussion
	KFSEventStreamEventFlagMustScanSubDirs KFSEventStreamEventFlag = 1
	// KFSEventStreamEventFlagNone: # Discussion
	KFSEventStreamEventFlagNone     KFSEventStreamEventFlag = 0
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
	case KFSEventStreamEventFlagItemCloned:
		return "KFSEventStreamEventFlagItemCloned"
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

type KFSEventStreamEventIdSince uint

const (
	KFSEventStreamEventIdSinceNow KFSEventStreamEventIdSince = 0
)

func (e KFSEventStreamEventIdSince) String() string {
	switch e {
	case KFSEventStreamEventIdSinceNow:
		return "KFSEventStreamEventIdSinceNow"
	default:
		return fmt.Sprintf("KFSEventStreamEventIdSince(%d)", e)
	}
}

type KFSFileOperation uint

const (
	KFSFileOperationDefaultOptions             KFSFileOperation = 0
	KFSFileOperationDoNotMoveAcrossVolumes     KFSFileOperation = 4
	KFSFileOperationOverwrite                  KFSFileOperation = 1
	KFSFileOperationSkipPreflight              KFSFileOperation = 8
	KFSFileOperationSkipSourcePermissionErrors KFSFileOperation = 2
)

func (e KFSFileOperation) String() string {
	switch e {
	case KFSFileOperationDefaultOptions:
		return "KFSFileOperationDefaultOptions"
	case KFSFileOperationDoNotMoveAcrossVolumes:
		return "KFSFileOperationDoNotMoveAcrossVolumes"
	case KFSFileOperationOverwrite:
		return "KFSFileOperationOverwrite"
	case KFSFileOperationSkipPreflight:
		return "KFSFileOperationSkipPreflight"
	case KFSFileOperationSkipSourcePermissionErrors:
		return "KFSFileOperationSkipSourcePermissionErrors"
	default:
		return fmt.Sprintf("KFSFileOperation(%d)", e)
	}
}

type KFSInvalidVolumeRef uint

const (
	KFSInvalidVolumeRefNum KFSInvalidVolumeRef = 0
)

func (e KFSInvalidVolumeRef) String() string {
	switch e {
	case KFSInvalidVolumeRefNum:
		return "KFSInvalidVolumeRefNum"
	default:
		return fmt.Sprintf("KFSInvalidVolumeRef(%d)", e)
	}
}

type KFSIterate uint

const (
	KFSIterateDelete   KFSIterate = 2
	KFSIterateFlat     KFSIterate = 0
	KFSIterateReserved KFSIterate = 0
	KFSIterateSubtree  KFSIterate = 1
)

func (e KFSIterate) String() string {
	switch e {
	case KFSIterateDelete:
		return "KFSIterateDelete"
	case KFSIterateFlat:
		return "KFSIterateFlat"
	case KFSIterateSubtree:
		return "KFSIterateSubtree"
	default:
		return fmt.Sprintf("KFSIterate(%d)", e)
	}
}

type KFSMountServer uint

const (
	KFSMountServerMarkDoNotDisplay     KFSMountServer = 1
	KFSMountServerMountOnMountDir      KFSMountServer = 4
	KFSMountServerSuppressConnectionUI KFSMountServer = 64
)

func (e KFSMountServer) String() string {
	switch e {
	case KFSMountServerMarkDoNotDisplay:
		return "KFSMountServerMarkDoNotDisplay"
	case KFSMountServerMountOnMountDir:
		return "KFSMountServerMountOnMountDir"
	case KFSMountServerSuppressConnectionUI:
		return "KFSMountServerSuppressConnectionUI"
	default:
		return fmt.Sprintf("KFSMountServer(%d)", e)
	}
}

type KFSMountServerMountWithout uint

const (
	KFSMountServerMountWithoutNotification KFSMountServerMountWithout = 2
)

func (e KFSMountServerMountWithout) String() string {
	switch e {
	case KFSMountServerMountWithoutNotification:
		return "KFSMountServerMountWithoutNotification"
	default:
		return fmt.Sprintf("KFSMountServerMountWithout(%d)", e)
	}
}

type KFSNode uint

const (
	KFSNodeCopyProtectBit   KFSNode = 6
	KFSNodeCopyProtectMask  KFSNode = 64
	KFSNodeDataOpenBit      KFSNode = 3
	KFSNodeDataOpenMask     KFSNode = 8
	KFSNodeForkOpenBit      KFSNode = 7
	KFSNodeForkOpenMask     KFSNode = 128
	KFSNodeHardLinkBit      KFSNode = 8
	KFSNodeHardLinkMask     KFSNode = 256
	KFSNodeInSharedBit      KFSNode = 2
	KFSNodeInSharedMask     KFSNode = 4
	KFSNodeIsDirectoryBit   KFSNode = 4
	KFSNodeIsDirectoryMask  KFSNode = 16
	KFSNodeIsMountedBit     KFSNode = 3
	KFSNodeIsMountedMask    KFSNode = 8
	KFSNodeIsSharePointBit  KFSNode = 5
	KFSNodeIsSharePointMask KFSNode = 32
	KFSNodeLockedBit        KFSNode = 0
	KFSNodeLockedMask       KFSNode = 1
	KFSNodeResOpenBit       KFSNode = 2
	KFSNodeResOpenMask      KFSNode = 4
)

func (e KFSNode) String() string {
	switch e {
	case KFSNodeCopyProtectBit:
		return "KFSNodeCopyProtectBit"
	case KFSNodeCopyProtectMask:
		return "KFSNodeCopyProtectMask"
	case KFSNodeDataOpenBit:
		return "KFSNodeDataOpenBit"
	case KFSNodeDataOpenMask:
		return "KFSNodeDataOpenMask"
	case KFSNodeForkOpenBit:
		return "KFSNodeForkOpenBit"
	case KFSNodeForkOpenMask:
		return "KFSNodeForkOpenMask"
	case KFSNodeHardLinkMask:
		return "KFSNodeHardLinkMask"
	case KFSNodeInSharedBit:
		return "KFSNodeInSharedBit"
	case KFSNodeInSharedMask:
		return "KFSNodeInSharedMask"
	case KFSNodeIsDirectoryMask:
		return "KFSNodeIsDirectoryMask"
	case KFSNodeIsSharePointBit:
		return "KFSNodeIsSharePointBit"
	case KFSNodeIsSharePointMask:
		return "KFSNodeIsSharePointMask"
	case KFSNodeLockedBit:
		return "KFSNodeLockedBit"
	case KFSNodeLockedMask:
		return "KFSNodeLockedMask"
	default:
		return fmt.Sprintf("KFSNode(%d)", e)
	}
}

type KFSOperationStage uint

const (
	KFSOperationStageComplete     KFSOperationStage = 3
	KFSOperationStagePreflighting KFSOperationStage = 1
	KFSOperationStageRunning      KFSOperationStage = 2
	KFSOperationStageUndefined    KFSOperationStage = 0
)

func (e KFSOperationStage) String() string {
	switch e {
	case KFSOperationStageComplete:
		return "KFSOperationStageComplete"
	case KFSOperationStagePreflighting:
		return "KFSOperationStagePreflighting"
	case KFSOperationStageRunning:
		return "KFSOperationStageRunning"
	case KFSOperationStageUndefined:
		return "KFSOperationStageUndefined"
	default:
		return fmt.Sprintf("KFSOperationStage(%d)", e)
	}
}

type KFSPathMakeRef uint

const (
	KFSPathMakeRefDefaultOptions         KFSPathMakeRef = 0
	KFSPathMakeRefDoNotFollowLeafSymlink KFSPathMakeRef = 1
)

func (e KFSPathMakeRef) String() string {
	switch e {
	case KFSPathMakeRefDefaultOptions:
		return "KFSPathMakeRefDefaultOptions"
	case KFSPathMakeRefDoNotFollowLeafSymlink:
		return "KFSPathMakeRefDoNotFollowLeafSymlink"
	default:
		return fmt.Sprintf("KFSPathMakeRef(%d)", e)
	}
}

type KFSReplaceObject uint

const (
	KFSReplaceObjectDefaultOptions              KFSReplaceObject = 0
	KFSReplaceObjectDoNotCheckObjectWriteAccess KFSReplaceObject = 16
	KFSReplaceObjectPreservePermissionInfo      KFSReplaceObject = 8
	KFSReplaceObjectReplaceMetadata             KFSReplaceObject = 1
	KFSReplaceObjectReplacePermissionInfo       KFSReplaceObject = 4
	KFSReplaceObjectSaveOriginalAsABackup       KFSReplaceObject = 2
)

func (e KFSReplaceObject) String() string {
	switch e {
	case KFSReplaceObjectDefaultOptions:
		return "KFSReplaceObjectDefaultOptions"
	case KFSReplaceObjectDoNotCheckObjectWriteAccess:
		return "KFSReplaceObjectDoNotCheckObjectWriteAccess"
	case KFSReplaceObjectPreservePermissionInfo:
		return "KFSReplaceObjectPreservePermissionInfo"
	case KFSReplaceObjectReplaceMetadata:
		return "KFSReplaceObjectReplaceMetadata"
	case KFSReplaceObjectReplacePermissionInfo:
		return "KFSReplaceObjectReplacePermissionInfo"
	case KFSReplaceObjectSaveOriginalAsABackup:
		return "KFSReplaceObjectSaveOriginalAsABackup"
	default:
		return fmt.Sprintf("KFSReplaceObject(%d)", e)
	}
}

type KFSUnmountVolumeForce uint

const (
	KFSUnmountVolumeForceUnmount KFSUnmountVolumeForce = 1
)

func (e KFSUnmountVolumeForce) String() string {
	switch e {
	case KFSUnmountVolumeForceUnmount:
		return "KFSUnmountVolumeForceUnmount"
	default:
		return fmt.Sprintf("KFSUnmountVolumeForce(%d)", e)
	}
}

type KFSVolFlag uint

const (
	KFSVolFlagDefaultVolumeBit     KFSVolFlag = 5
	KFSVolFlagDefaultVolumeMask    KFSVolFlag = 32
	KFSVolFlagFilesOpenBit         KFSVolFlag = 6
	KFSVolFlagFilesOpenMask        KFSVolFlag = 64
	KFSVolFlagHardwareLockedBit    KFSVolFlag = 7
	KFSVolFlagHardwareLockedMask   KFSVolFlag = 128
	KFSVolFlagJournalingActiveBit  KFSVolFlag = 14
	KFSVolFlagJournalingActiveMask KFSVolFlag = 16384
	KFSVolFlagSoftwareLockedBit    KFSVolFlag = 15
	KFSVolFlagSoftwareLockedMask   KFSVolFlag = 32768
)

func (e KFSVolFlag) String() string {
	switch e {
	case KFSVolFlagDefaultVolumeBit:
		return "KFSVolFlagDefaultVolumeBit"
	case KFSVolFlagDefaultVolumeMask:
		return "KFSVolFlagDefaultVolumeMask"
	case KFSVolFlagFilesOpenBit:
		return "KFSVolFlagFilesOpenBit"
	case KFSVolFlagFilesOpenMask:
		return "KFSVolFlagFilesOpenMask"
	case KFSVolFlagHardwareLockedBit:
		return "KFSVolFlagHardwareLockedBit"
	case KFSVolFlagHardwareLockedMask:
		return "KFSVolFlagHardwareLockedMask"
	case KFSVolFlagJournalingActiveBit:
		return "KFSVolFlagJournalingActiveBit"
	case KFSVolFlagJournalingActiveMask:
		return "KFSVolFlagJournalingActiveMask"
	case KFSVolFlagSoftwareLockedBit:
		return "KFSVolFlagSoftwareLockedBit"
	case KFSVolFlagSoftwareLockedMask:
		return "KFSVolFlagSoftwareLockedMask"
	default:
		return fmt.Sprintf("KFSVolFlag(%d)", e)
	}
}

type KFSVolInfo uint

const (
	KFSVolInfoBackupDate   KFSVolInfo = 4
	KFSVolInfoBlocks       KFSVolInfo = 128
	KFSVolInfoCheckedDate  KFSVolInfo = 8
	KFSVolInfoCreateDate   KFSVolInfo = 1
	KFSVolInfoDataClump    KFSVolInfo = 1024
	KFSVolInfoDirCount     KFSVolInfo = 32
	KFSVolInfoDriveInfo    KFSVolInfo = 32768
	KFSVolInfoFSInfo       KFSVolInfo = 16384
	KFSVolInfoFileCount    KFSVolInfo = 16
	KFSVolInfoFinderInfo   KFSVolInfo = 4096
	KFSVolInfoFlags        KFSVolInfo = 8192
	KFSVolInfoGettableInfo KFSVolInfo = 65535
	KFSVolInfoModDate      KFSVolInfo = 2
	KFSVolInfoNextAlloc    KFSVolInfo = 256
	KFSVolInfoNextID       KFSVolInfo = 2048
	KFSVolInfoNone         KFSVolInfo = 0
	KFSVolInfoRsrcClump    KFSVolInfo = 512
	KFSVolInfoSettableInfo KFSVolInfo = 12292
	KFSVolInfoSizes        KFSVolInfo = 64
)

func (e KFSVolInfo) String() string {
	switch e {
	case KFSVolInfoBackupDate:
		return "KFSVolInfoBackupDate"
	case KFSVolInfoBlocks:
		return "KFSVolInfoBlocks"
	case KFSVolInfoCheckedDate:
		return "KFSVolInfoCheckedDate"
	case KFSVolInfoCreateDate:
		return "KFSVolInfoCreateDate"
	case KFSVolInfoDataClump:
		return "KFSVolInfoDataClump"
	case KFSVolInfoDirCount:
		return "KFSVolInfoDirCount"
	case KFSVolInfoDriveInfo:
		return "KFSVolInfoDriveInfo"
	case KFSVolInfoFSInfo:
		return "KFSVolInfoFSInfo"
	case KFSVolInfoFileCount:
		return "KFSVolInfoFileCount"
	case KFSVolInfoFinderInfo:
		return "KFSVolInfoFinderInfo"
	case KFSVolInfoFlags:
		return "KFSVolInfoFlags"
	case KFSVolInfoGettableInfo:
		return "KFSVolInfoGettableInfo"
	case KFSVolInfoModDate:
		return "KFSVolInfoModDate"
	case KFSVolInfoNextAlloc:
		return "KFSVolInfoNextAlloc"
	case KFSVolInfoNextID:
		return "KFSVolInfoNextID"
	case KFSVolInfoNone:
		return "KFSVolInfoNone"
	case KFSVolInfoRsrcClump:
		return "KFSVolInfoRsrcClump"
	case KFSVolInfoSettableInfo:
		return "KFSVolInfoSettableInfo"
	case KFSVolInfoSizes:
		return "KFSVolInfoSizes"
	default:
		return fmt.Sprintf("KFSVolInfo(%d)", e)
	}
}

type KFirstMagicBusyFiletype uint

const (
	KFirstMagicBusyFiletypeValue KFirstMagicBusyFiletype = 'b'<<24 | 'z'<<16 | 'y'<<8 | ' ' // 'bzy '
	KLastMagicBusyFiletype       KFirstMagicBusyFiletype = 'b'<<24 | 'z'<<16 | 'y'<<8 | '?' // 'bzy?'
)

func (e KFirstMagicBusyFiletype) String() string {
	switch e {
	case KFirstMagicBusyFiletypeValue:
		return "KFirstMagicBusyFiletypeValue"
	case KLastMagicBusyFiletype:
		return "KLastMagicBusyFiletype"
	default:
		return fmt.Sprintf("KFirstMagicBusyFiletype(%d)", e)
	}
}

type KFolder uint

const (
	// Deprecated.
	KFolderInLocalOrRemoteUserFolder KFolder = 32
	// Deprecated.
	KFolderInRemoteUserFolderIfAvailable KFolder = 128
	// Deprecated.
	KFolderInRemoteUserFolderIfAvailableBit KFolder = 7
	// Deprecated.
	KFolderInUserFolder KFolder = 32
	// Deprecated.
	KFolderInUserFolderBit KFolder = 5
	// Deprecated.
	KFolderManagerFolderInMacOS9FolderIfMacOSXIsInstalledBit KFolder = 10
	// Deprecated.
	KFolderManagerFolderInMacOS9FolderIfMacOSXIsInstalledMask KFolder = 1024
	// Deprecated.
	KFolderManagerNewlyCreatedFolderIsLocalizedBit KFolder = 12
	// Deprecated.
	KFolderManagerNewlyCreatedFolderShouldHaveDotLocalizedCreatedWithinMask KFolder = 4096
	// Deprecated.
	KFolderManagerNotCreatedOnRemoteVolumesBit KFolder = 11
	// Deprecated.
	KFolderManagerNotCreatedOnRemoteVolumesMask KFolder = 2048
	// Deprecated.
	KFolderMustStayOnSameVolume KFolder = 512
	// Deprecated.
	KFolderMustStayOnSameVolumeBit KFolder = 9
	// Deprecated.
	KFolderNeverMatchedInIdentifyFolder KFolder = 256
	// Deprecated.
	KFolderNeverMatchedInIdentifyFolderBit KFolder = 8
	// Deprecated.
	KFolderTrackedByAlias KFolder = 64
	// Deprecated.
	KFolderTrackedByAliasBit KFolder = 6
)

func (e KFolder) String() string {
	switch e {
	case KFolderInLocalOrRemoteUserFolder:
		return "KFolderInLocalOrRemoteUserFolder"
	case KFolderInRemoteUserFolderIfAvailable:
		return "KFolderInRemoteUserFolderIfAvailable"
	case KFolderInRemoteUserFolderIfAvailableBit:
		return "KFolderInRemoteUserFolderIfAvailableBit"
	case KFolderInUserFolderBit:
		return "KFolderInUserFolderBit"
	case KFolderManagerFolderInMacOS9FolderIfMacOSXIsInstalledBit:
		return "KFolderManagerFolderInMacOS9FolderIfMacOSXIsInstalledBit"
	case KFolderManagerFolderInMacOS9FolderIfMacOSXIsInstalledMask:
		return "KFolderManagerFolderInMacOS9FolderIfMacOSXIsInstalledMask"
	case KFolderManagerNewlyCreatedFolderIsLocalizedBit:
		return "KFolderManagerNewlyCreatedFolderIsLocalizedBit"
	case KFolderManagerNewlyCreatedFolderShouldHaveDotLocalizedCreatedWithinMask:
		return "KFolderManagerNewlyCreatedFolderShouldHaveDotLocalizedCreatedWithinMask"
	case KFolderManagerNotCreatedOnRemoteVolumesBit:
		return "KFolderManagerNotCreatedOnRemoteVolumesBit"
	case KFolderManagerNotCreatedOnRemoteVolumesMask:
		return "KFolderManagerNotCreatedOnRemoteVolumesMask"
	case KFolderMustStayOnSameVolume:
		return "KFolderMustStayOnSameVolume"
	case KFolderMustStayOnSameVolumeBit:
		return "KFolderMustStayOnSameVolumeBit"
	case KFolderNeverMatchedInIdentifyFolder:
		return "KFolderNeverMatchedInIdentifyFolder"
	case KFolderNeverMatchedInIdentifyFolderBit:
		return "KFolderNeverMatchedInIdentifyFolderBit"
	case KFolderTrackedByAlias:
		return "KFolderTrackedByAlias"
	case KFolderTrackedByAliasBit:
		return "KFolderTrackedByAliasBit"
	default:
		return fmt.Sprintf("KFolder(%d)", e)
	}
}

type KForkInfoFlags uint

const (
	KForkInfoFlagsFileLockedBit   KForkInfoFlags = 0
	KForkInfoFlagsFileLockedMask  KForkInfoFlags = 0
	KForkInfoFlagsLargeFileBit    KForkInfoFlags = 0
	KForkInfoFlagsLargeFileMask   KForkInfoFlags = 0
	KForkInfoFlagsModifiedBit     KForkInfoFlags = 0
	KForkInfoFlagsModifiedMask    KForkInfoFlags = 0
	KForkInfoFlagsOwnClumpBit     KForkInfoFlags = 0
	KForkInfoFlagsOwnClumpMask    KForkInfoFlags = 0
	KForkInfoFlagsResourceBit     KForkInfoFlags = 0
	KForkInfoFlagsResourceMask    KForkInfoFlags = 0
	KForkInfoFlagsSharedWriteBit  KForkInfoFlags = 0
	KForkInfoFlagsSharedWriteMask KForkInfoFlags = 0
	KForkInfoFlagsWriteBit        KForkInfoFlags = 0
	KForkInfoFlagsWriteLockedBit  KForkInfoFlags = 0
	KForkInfoFlagsWriteLockedMask KForkInfoFlags = 0
	KForkInfoFlagsWriteMask       KForkInfoFlags = 0
)

func (e KForkInfoFlags) String() string {
	switch e {
	case KForkInfoFlagsFileLockedBit:
		return "KForkInfoFlagsFileLockedBit"
	default:
		return fmt.Sprintf("KForkInfoFlags(%d)", e)
	}
}

type KFragment uint

const (
	// Deprecated.
	KFragmentIsPrepared KFragment = 0
	// Deprecated.
	KFragmentNeedsPreparing KFragment = 2
)

func (e KFragment) String() string {
	switch e {
	case KFragmentIsPrepared:
		return "KFragmentIsPrepared"
	case KFragmentNeedsPreparing:
		return "KFragmentNeedsPreparing"
	default:
		return fmt.Sprintf("KFragment(%d)", e)
	}
}

type KGenericDocumentIconResource int

const (
	KFloppyIconResource               KGenericDocumentIconResource = -3998
	KGenericApplicationIconResource   KGenericDocumentIconResource = -3996
	KGenericCDROMIconResource         KGenericDocumentIconResource = -3987
	KGenericDeskAccessoryIconResource KGenericDocumentIconResource = -3991
	KGenericDocumentIconResourceValue KGenericDocumentIconResource = -4000
	KGenericEditionFileIconResource   KGenericDocumentIconResource = -3989
	KGenericFolderIconResource        KGenericDocumentIconResource = -3999
	KGenericRAMDiskIconResource       KGenericDocumentIconResource = -3988
	KGenericStationeryIconResource    KGenericDocumentIconResource = -3985
	KPrivateFolderIconResource        KGenericDocumentIconResource = -3994
	KTrashIconResource                KGenericDocumentIconResource = -3993
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
	KDropFolderIcon         KGenericFolderIcon = 'd'<<24 | 'b'<<16 | 'o'<<8 | 'x' // 'dbox'
	KGenericFolderIconValue KGenericFolderIcon = 'f'<<24 | 'l'<<16 | 'd'<<8 | 'r' // 'fldr'
	KMountedFolderIcon      KGenericFolderIcon = 'm'<<24 | 'n'<<16 | 't'<<8 | 'd' // 'mntd'
	KOpenFolderIcon         KGenericFolderIcon = 'o'<<24 | 'f'<<16 | 'l'<<8 | 'd' // 'ofld'
	KOwnedFolderIcon        KGenericFolderIcon = 'o'<<24 | 'w'<<16 | 'n'<<8 | 'd' // 'ownd'
	KPrivateFolderIcon      KGenericFolderIcon = 'p'<<24 | 'r'<<16 | 'v'<<8 | 'f' // 'prvf'
	KSharedFolderIcon       KGenericFolderIcon = 's'<<24 | 'h'<<16 | 'f'<<8 | 'l' // 'shfl'
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
	KAppleMenuFolderIconResource         KGenericPreferencesIconResource = -3982
	KGenericExtensionIconResource        KGenericPreferencesIconResource = -16415
	KGenericPreferencesIconResourceValue KGenericPreferencesIconResource = -3971
	KGenericQueryDocumentIconResource    KGenericPreferencesIconResource = -16506
	KHelpIconResource                    KGenericPreferencesIconResource = -20271
	KSystemFolderIconResource            KGenericPreferencesIconResource = -3983
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

type KGetDebugOption uint

const (
	// Deprecated.
	KGetDebugOptionValue KGetDebugOption = 1
	// Deprecated.
	KSetDebugOption KGetDebugOption = 2
)

func (e KGetDebugOption) String() string {
	switch e {
	case KGetDebugOptionValue:
		return "KGetDebugOptionValue"
	case KSetDebugOption:
		return "KSetDebugOption"
	default:
		return fmt.Sprintf("KGetDebugOption(%d)", e)
	}
}

type KHID int

const (
	KHIDBadLogPhysValuesErr         KHID = -13942
	KHIDBadLogicalMaximumErr        KHID = -13933
	KHIDBadLogicalMinimumErr        KHID = -13934
	KHIDBadParameterErr             KHID = -13938
	KHIDBaseError                   KHID = -13950
	KHIDBufferTooSmallErr           KHID = -13948
	KHIDDeviceNotReady              KHID = -13910
	KHIDEndOfDescriptorErr          KHID = -13936
	KHIDIncompatibleReportErr       KHID = -13943
	KHIDInvalidPreparsedDataErr     KHID = -13944
	KHIDInvalidRangePageErr         KHID = -13923
	KHIDInvalidReportLengthErr      KHID = -13940
	KHIDInvalidReportTypeErr        KHID = -13941
	KHIDInvertedLogicalRangeErr     KHID = -13932
	KHIDInvertedPhysicalRangeErr    KHID = -13931
	KHIDInvertedUsageRangeErr       KHID = -13929
	KHIDNotEnoughMemoryErr          KHID = -13937
	KHIDNotValueArrayErr            KHID = -13945
	KHIDNullPointerErr              KHID = -13939
	KHIDNullStateErr                KHID = -13949
	KHIDReportCountZeroErr          KHID = -13925
	KHIDReportIDZeroErr             KHID = -13924
	KHIDReportSizeZeroErr           KHID = -13926
	KHIDSuccess                     KHID = 0
	KHIDUnmatchedDesignatorRangeErr KHID = -13927
	KHIDUnmatchedStringRangeErr     KHID = -13928
	KHIDUnmatchedUsageRangeErr      KHID = -13930
	KHIDUsageNotFoundErr            KHID = -13946
	KHIDUsagePageZeroErr            KHID = -13935
	KHIDValueOutOfRangeErr          KHID = -13947
	KHIDVersionIncompatibleErr      KHID = -13909
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

type KHandle uint

const (
	// Deprecated.
	KHandleIsResourceBit KHandle = 5
	// Deprecated.
	KHandleIsResourceMask KHandle = 32
	// Deprecated.
	KHandleLockedBit KHandle = 7
	// Deprecated.
	KHandleLockedMask KHandle = 128
	// Deprecated.
	KHandlePurgeableBit KHandle = 6
	// Deprecated.
	KHandlePurgeableMask KHandle = 64
)

func (e KHandle) String() string {
	switch e {
	case KHandleIsResourceBit:
		return "KHandleIsResourceBit"
	case KHandleIsResourceMask:
		return "KHandleIsResourceMask"
	case KHandleLockedBit:
		return "KHandleLockedBit"
	case KHandleLockedMask:
		return "KHandleLockedMask"
	case KHandlePurgeableBit:
		return "KHandlePurgeableBit"
	case KHandlePurgeableMask:
		return "KHandlePurgeableMask"
	default:
		return fmt.Sprintf("KHandle(%d)", e)
	}
}

type KHuge1BitMask uint

const (
	KHuge1BitMaskValue KHuge1BitMask = 'i'<<24 | 'c'<<16 | 'h'<<8 | '#' // 'ich#'
	KHuge32BitData     KHuge1BitMask = 'i'<<24 | 'h'<<16 | '3'<<8 | '2' // 'ih32'
	KHuge4BitData      KHuge1BitMask = 'i'<<24 | 'c'<<16 | 'h'<<8 | '4' // 'ich4'
	KHuge8BitData      KHuge1BitMask = 'i'<<24 | 'c'<<16 | 'h'<<8 | '8' // 'ich8'
	KHuge8BitMask      KHuge1BitMask = 'h'<<24 | '8'<<16 | 'm'<<8 | 'k' // 'h8mk'
)

func (e KHuge1BitMask) String() string {
	switch e {
	case KHuge1BitMaskValue:
		return "KHuge1BitMaskValue"
	case KHuge32BitData:
		return "KHuge32BitData"
	case KHuge4BitData:
		return "KHuge4BitData"
	case KHuge8BitData:
		return "KHuge8BitData"
	case KHuge8BitMask:
		return "KHuge8BitMask"
	default:
		return fmt.Sprintf("KHuge1BitMask(%d)", e)
	}
}

type KISOLatin1 uint

const (
	KISOLatin1MusicCDVariant  KISOLatin1 = 1
	KISOLatin1StandardVariant KISOLatin1 = 0
)

func (e KISOLatin1) String() string {
	switch e {
	case KISOLatin1MusicCDVariant:
		return "KISOLatin1MusicCDVariant"
	case KISOLatin1StandardVariant:
		return "KISOLatin1StandardVariant"
	default:
		return fmt.Sprintf("KISOLatin1(%d)", e)
	}
}

type KISOLatinArabic uint

const (
	KISOLatinArabicExplicitOrderVariant KISOLatinArabic = 2
	KISOLatinArabicImplicitOrderVariant KISOLatinArabic = 0
	KISOLatinArabicVisualOrderVariant   KISOLatinArabic = 1
)

func (e KISOLatinArabic) String() string {
	switch e {
	case KISOLatinArabicExplicitOrderVariant:
		return "KISOLatinArabicExplicitOrderVariant"
	case KISOLatinArabicImplicitOrderVariant:
		return "KISOLatinArabicImplicitOrderVariant"
	case KISOLatinArabicVisualOrderVariant:
		return "KISOLatinArabicVisualOrderVariant"
	default:
		return fmt.Sprintf("KISOLatinArabic(%d)", e)
	}
}

type KISOLatinHebrew uint

const (
	KISOLatinHebrewExplicitOrderVariant KISOLatinHebrew = 2
	KISOLatinHebrewImplicitOrderVariant KISOLatinHebrew = 0
	KISOLatinHebrewVisualOrderVariant   KISOLatinHebrew = 1
)

func (e KISOLatinHebrew) String() string {
	switch e {
	case KISOLatinHebrewExplicitOrderVariant:
		return "KISOLatinHebrewExplicitOrderVariant"
	case KISOLatinHebrewImplicitOrderVariant:
		return "KISOLatinHebrewImplicitOrderVariant"
	case KISOLatinHebrewVisualOrderVariant:
		return "KISOLatinHebrewVisualOrderVariant"
	default:
		return fmt.Sprintf("KISOLatinHebrew(%d)", e)
	}
}

type KISp int

const (
	KISpBufferToSmallErr    KISp = -30422
	KISpDeviceActiveErr     KISp = -30428
	KISpDeviceInactiveErr   KISp = -30426
	KISpElementInListErr    KISp = -30423
	KISpElementNotInListErr KISp = -30424
	KISpInternalErr         KISp = -30420
	KISpListBusyErr         KISp = -30429
	KISpSystemActiveErr     KISp = -30427
	KISpSystemInactiveErr   KISp = -30425
	KISpSystemListErr       KISp = -30421
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

type KIcon uint

const (
	KIconServices128PixelDataARGB KIcon = 'i'<<24 | 'c'<<16 | '0'<<8 | '7' // 'ic07'
	KIconServices16PixelDataARGB  KIcon = 'i'<<24 | 'c'<<16 | '0'<<8 | '4' // 'ic04'
	KIconServices32PixelDataARGB  KIcon = 'i'<<24 | 'c'<<16 | '0'<<8 | '5' // 'ic05'
	KIconServices48PixelDataARGB  KIcon = 'i'<<24 | 'c'<<16 | '0'<<8 | '6' // 'ic06'
)

func (e KIcon) String() string {
	switch e {
	case KIconServices128PixelDataARGB:
		return "KIconServices128PixelDataARGB"
	case KIconServices16PixelDataARGB:
		return "KIconServices16PixelDataARGB"
	case KIconServices32PixelDataARGB:
		return "KIconServices32PixelDataARGB"
	case KIconServices48PixelDataARGB:
		return "KIconServices48PixelDataARGB"
	default:
		return fmt.Sprintf("KIcon(%d)", e)
	}
}

type KIconFamily uint

const (
	KIconFamilyType KIconFamily = 'i'<<24 | 'c'<<16 | 'n'<<8 | 's' // 'icns'
)

func (e KIconFamily) String() string {
	switch e {
	case KIconFamilyType:
		return "KIconFamilyType"
	default:
		return fmt.Sprintf("KIconFamily(%d)", e)
	}
}

type KIconServices uint

const (
	KIconServicesNoBadgeFlag        KIconServices = 1
	KIconServicesNormalUsageFlag    KIconServices = 0
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

type KIconServices256PixelDataARGB uint

const (
	KIconServices1024PixelDataARGB     KIconServices256PixelDataARGB = 'i'<<24 | 'c'<<16 | '1'<<8 | '0' // 'ic10'
	KIconServices256PixelDataARGBValue KIconServices256PixelDataARGB = 'i'<<24 | 'c'<<16 | '0'<<8 | '8' // 'ic08'
	KIconServices512PixelDataARGB      KIconServices256PixelDataARGB = 'i'<<24 | 'c'<<16 | '0'<<8 | '9' // 'ic09'
	KThumbnail32BitData                KIconServices256PixelDataARGB = 'i'<<24 | 't'<<16 | '3'<<8 | '2' // 'it32'
	KThumbnail8BitMask                 KIconServices256PixelDataARGB = 't'<<24 | '8'<<16 | 'm'<<8 | 'k' // 't8mk'
)

func (e KIconServices256PixelDataARGB) String() string {
	switch e {
	case KIconServices1024PixelDataARGB:
		return "KIconServices1024PixelDataARGB"
	case KIconServices256PixelDataARGBValue:
		return "KIconServices256PixelDataARGBValue"
	case KIconServices512PixelDataARGB:
		return "KIconServices512PixelDataARGB"
	case KThumbnail32BitData:
		return "KThumbnail32BitData"
	case KThumbnail8BitMask:
		return "KThumbnail8BitMask"
	default:
		return fmt.Sprintf("KIconServices256PixelDataARGB(%d)", e)
	}
}

type KIconServicesCatalogInfo uint

const (
	KIconServicesCatalogInfoMask KIconServicesCatalogInfo = 0
)

func (e KIconServicesCatalogInfo) String() string {
	switch e {
	case KIconServicesCatalogInfoMask:
		return "KIconServicesCatalogInfoMask"
	default:
		return fmt.Sprintf("KIconServicesCatalogInfo(%d)", e)
	}
}

type KIdleKCEvent uint

const (
	KAddKCEvent                 KIdleKCEvent = 3
	KDataAccessKCEvent          KIdleKCEvent = 10
	KDefaultChangedKCEvent      KIdleKCEvent = 9
	KDeleteKCEvent              KIdleKCEvent = 4
	KIdleKCEventValue           KIdleKCEvent = 0
	KKeychainListChangedKCEvent KIdleKCEvent = 11
	KLockKCEvent                KIdleKCEvent = 1
	KPasswordChangedKCEvent     KIdleKCEvent = 6
	KSystemKCEvent              KIdleKCEvent = 8
	KUnlockKCEvent              KIdleKCEvent = 2
	KUpdateKCEvent              KIdleKCEvent = 5
)

func (e KIdleKCEvent) String() string {
	switch e {
	case KAddKCEvent:
		return "KAddKCEvent"
	case KDataAccessKCEvent:
		return "KDataAccessKCEvent"
	case KDefaultChangedKCEvent:
		return "KDefaultChangedKCEvent"
	case KDeleteKCEvent:
		return "KDeleteKCEvent"
	case KIdleKCEventValue:
		return "KIdleKCEventValue"
	case KKeychainListChangedKCEvent:
		return "KKeychainListChangedKCEvent"
	case KLockKCEvent:
		return "KLockKCEvent"
	case KPasswordChangedKCEvent:
		return "KPasswordChangedKCEvent"
	case KSystemKCEvent:
		return "KSystemKCEvent"
	case KUnlockKCEvent:
		return "KUnlockKCEvent"
	case KUpdateKCEvent:
		return "KUpdateKCEvent"
	default:
		return fmt.Sprintf("KIdleKCEvent(%d)", e)
	}
}

type KIdleKCEventMask uint

const (
	KAddKCEventMask             KIdleKCEventMask = 8
	KDataAccessKCEventMask      KIdleKCEventMask = 1024
	KDefaultChangedKCEventMask  KIdleKCEventMask = 512
	KDeleteKCEventMask          KIdleKCEventMask = 16
	KEveryKCEventMask           KIdleKCEventMask = 65535
	KIdleKCEventMaskValue       KIdleKCEventMask = 1
	KLockKCEventMask            KIdleKCEventMask = 2
	KPasswordChangedKCEventMask KIdleKCEventMask = 64
	KSystemEventKCEventMask     KIdleKCEventMask = 256
	KUnlockKCEventMask          KIdleKCEventMask = 4
	KUpdateKCEventMask          KIdleKCEventMask = 32
)

func (e KIdleKCEventMask) String() string {
	switch e {
	case KAddKCEventMask:
		return "KAddKCEventMask"
	case KDataAccessKCEventMask:
		return "KDataAccessKCEventMask"
	case KDefaultChangedKCEventMask:
		return "KDefaultChangedKCEventMask"
	case KDeleteKCEventMask:
		return "KDeleteKCEventMask"
	case KEveryKCEventMask:
		return "KEveryKCEventMask"
	case KIdleKCEventMaskValue:
		return "KIdleKCEventMaskValue"
	case KLockKCEventMask:
		return "KLockKCEventMask"
	case KPasswordChangedKCEventMask:
		return "KPasswordChangedKCEventMask"
	case KSystemEventKCEventMask:
		return "KSystemEventKCEventMask"
	case KUnlockKCEventMask:
		return "KUnlockKCEventMask"
	case KUpdateKCEventMask:
		return "KUpdateKCEventMask"
	default:
		return fmt.Sprintf("KIdleKCEventMask(%d)", e)
	}
}

type KInternetLocation uint

const (
	KInternetLocationAFP                 KInternetLocation = 'i'<<24 | 'l'<<16 | 'a'<<8 | 'f' // 'ilaf'
	KInternetLocationAppleShareIcon      KInternetLocation = 'i'<<24 | 'l'<<16 | 'a'<<8 | 'f' // 'ilaf'
	KInternetLocationAppleTalk           KInternetLocation = 'i'<<24 | 'l'<<16 | 'a'<<8 | 't' // 'ilat'
	KInternetLocationAppleTalkZoneIcon   KInternetLocation = 'i'<<24 | 'l'<<16 | 'a'<<8 | 't' // 'ilat'
	KInternetLocationCreator             KInternetLocation = 'd'<<24 | 'r'<<16 | 'a'<<8 | 'g' // 'drag'
	KInternetLocationFTP                 KInternetLocation = 'i'<<24 | 'l'<<16 | 'f'<<8 | 't' // 'ilft'
	KInternetLocationFTPIcon             KInternetLocation = 'i'<<24 | 'l'<<16 | 'f'<<8 | 't' // 'ilft'
	KInternetLocationFile                KInternetLocation = 'i'<<24 | 'l'<<16 | 'f'<<8 | 'i' // 'ilfi'
	KInternetLocationFileIcon            KInternetLocation = 'i'<<24 | 'l'<<16 | 'f'<<8 | 'i' // 'ilfi'
	KInternetLocationGeneric             KInternetLocation = 'i'<<24 | 'l'<<16 | 'g'<<8 | 'e' // 'ilge'
	KInternetLocationGenericIcon         KInternetLocation = 'i'<<24 | 'l'<<16 | 'g'<<8 | 'e' // 'ilge'
	KInternetLocationHTTP                KInternetLocation = 'i'<<24 | 'l'<<16 | 'h'<<8 | 't' // 'ilht'
	KInternetLocationHTTPIcon            KInternetLocation = 'i'<<24 | 'l'<<16 | 'h'<<8 | 't' // 'ilht'
	KInternetLocationMail                KInternetLocation = 'i'<<24 | 'l'<<16 | 'm'<<8 | 'a' // 'ilma'
	KInternetLocationMailIcon            KInternetLocation = 'i'<<24 | 'l'<<16 | 'm'<<8 | 'a' // 'ilma'
	KInternetLocationNNTP                KInternetLocation = 'i'<<24 | 'l'<<16 | 'n'<<8 | 'w' // 'ilnw'
	KInternetLocationNSL                 KInternetLocation = 'i'<<24 | 'l'<<16 | 'n'<<8 | 's' // 'ilns'
	KInternetLocationNSLNeighborhoodIcon KInternetLocation = 'i'<<24 | 'l'<<16 | 'n'<<8 | 's' // 'ilns'
	KInternetLocationNewsIcon            KInternetLocation = 'i'<<24 | 'l'<<16 | 'n'<<8 | 'w' // 'ilnw'
)

func (e KInternetLocation) String() string {
	switch e {
	case KInternetLocationAFP:
		return "KInternetLocationAFP"
	case KInternetLocationAppleTalk:
		return "KInternetLocationAppleTalk"
	case KInternetLocationCreator:
		return "KInternetLocationCreator"
	case KInternetLocationFTP:
		return "KInternetLocationFTP"
	case KInternetLocationFile:
		return "KInternetLocationFile"
	case KInternetLocationGeneric:
		return "KInternetLocationGeneric"
	case KInternetLocationHTTP:
		return "KInternetLocationHTTP"
	case KInternetLocationMail:
		return "KInternetLocationMail"
	case KInternetLocationNNTP:
		return "KInternetLocationNNTP"
	case KInternetLocationNSL:
		return "KInternetLocationNSL"
	default:
		return fmt.Sprintf("KInternetLocation(%d)", e)
	}
}

type KIsOnDesk uint

const (
	KColor         KIsOnDesk = 14
	KHasBeenInited KIsOnDesk = 256
	KHasBundle     KIsOnDesk = 8192
	KHasCustomIcon KIsOnDesk = 1024
	KHasNoINITs    KIsOnDesk = 128
	KIsAlias       KIsOnDesk = 32768
	KIsInvisible   KIsOnDesk = 16384
	KIsOnDeskValue KIsOnDesk = 1
	KIsShared      KIsOnDesk = 64
	KIsStationery  KIsOnDesk = 2048
	KNameLocked    KIsOnDesk = 4096
)

func (e KIsOnDesk) String() string {
	switch e {
	case KColor:
		return "KColor"
	case KHasBeenInited:
		return "KHasBeenInited"
	case KHasBundle:
		return "KHasBundle"
	case KHasCustomIcon:
		return "KHasCustomIcon"
	case KHasNoINITs:
		return "KHasNoINITs"
	case KIsAlias:
		return "KIsAlias"
	case KIsInvisible:
		return "KIsInvisible"
	case KIsOnDeskValue:
		return "KIsOnDeskValue"
	case KIsShared:
		return "KIsShared"
	case KIsStationery:
		return "KIsStationery"
	case KNameLocked:
		return "KNameLocked"
	default:
		return fmt.Sprintf("KIsOnDesk(%d)", e)
	}
}

type KKCAuthType uint

const (
	KKCAuthTypeDPA        KKCAuthType = 'd'<<24 | 'p'<<16 | 'a'<<8 | 'a' // 'dpaa'
	KKCAuthTypeDefault    KKCAuthType = 'd'<<24 | 'f'<<16 | 'l'<<8 | 't' // 'dflt'
	KKCAuthTypeHTTPDigest KKCAuthType = 'h'<<24 | 't'<<16 | 't'<<8 | 'd' // 'httd'
	KKCAuthTypeMSN        KKCAuthType = 'm'<<24 | 's'<<16 | 'n'<<8 | 'a' // 'msna'
	KKCAuthTypeNTLM       KKCAuthType = 'n'<<24 | 't'<<16 | 'l'<<8 | 'm' // 'ntlm'
	KKCAuthTypeRPA        KKCAuthType = 'r'<<24 | 'p'<<16 | 'a'<<8 | 'a' // 'rpaa'
)

func (e KKCAuthType) String() string {
	switch e {
	case KKCAuthTypeDPA:
		return "KKCAuthTypeDPA"
	case KKCAuthTypeDefault:
		return "KKCAuthTypeDefault"
	case KKCAuthTypeHTTPDigest:
		return "KKCAuthTypeHTTPDigest"
	case KKCAuthTypeMSN:
		return "KKCAuthTypeMSN"
	case KKCAuthTypeNTLM:
		return "KKCAuthTypeNTLM"
	case KKCAuthTypeRPA:
		return "KKCAuthTypeRPA"
	default:
		return fmt.Sprintf("KKCAuthType(%d)", e)
	}
}

type KKCProtocolType uint

const (
	KKCProtocolTypeAFP        KKCProtocolType = 'a'<<24 | 'f'<<16 | 'p'<<8 | ' ' // 'afp '
	KKCProtocolTypeAppleTalk  KKCProtocolType = 'a'<<24 | 't'<<16 | 'l'<<8 | 'k' // 'atlk'
	KKCProtocolTypeFTP        KKCProtocolType = 'f'<<24 | 't'<<16 | 'p'<<8 | ' ' // 'ftp '
	KKCProtocolTypeFTPAccount KKCProtocolType = 'f'<<24 | 't'<<16 | 'p'<<8 | 'a' // 'ftpa'
	KKCProtocolTypeHTTP       KKCProtocolType = 'h'<<24 | 't'<<16 | 't'<<8 | 'p' // 'http'
	KKCProtocolTypeIMAP       KKCProtocolType = 'i'<<24 | 'm'<<16 | 'a'<<8 | 'p' // 'imap'
	KKCProtocolTypeIRC        KKCProtocolType = 'i'<<24 | 'r'<<16 | 'c'<<8 | ' ' // 'irc '
	KKCProtocolTypeLDAP       KKCProtocolType = 'l'<<24 | 'd'<<16 | 'a'<<8 | 'p' // 'ldap'
	KKCProtocolTypeNNTP       KKCProtocolType = 'n'<<24 | 'n'<<16 | 't'<<8 | 'p' // 'nntp'
	KKCProtocolTypePOP3       KKCProtocolType = 'p'<<24 | 'o'<<16 | 'p'<<8 | '3' // 'pop3'
	KKCProtocolTypeSMTP       KKCProtocolType = 's'<<24 | 'm'<<16 | 't'<<8 | 'p' // 'smtp'
	KKCProtocolTypeSOCKS      KKCProtocolType = 's'<<24 | 'o'<<16 | 'x'<<8 | ' ' // 'sox '
	KKCProtocolTypeTelnet     KKCProtocolType = 't'<<24 | 'e'<<16 | 'l'<<8 | 'n' // 'teln'
)

func (e KKCProtocolType) String() string {
	switch e {
	case KKCProtocolTypeAFP:
		return "KKCProtocolTypeAFP"
	case KKCProtocolTypeAppleTalk:
		return "KKCProtocolTypeAppleTalk"
	case KKCProtocolTypeFTP:
		return "KKCProtocolTypeFTP"
	case KKCProtocolTypeFTPAccount:
		return "KKCProtocolTypeFTPAccount"
	case KKCProtocolTypeHTTP:
		return "KKCProtocolTypeHTTP"
	case KKCProtocolTypeIMAP:
		return "KKCProtocolTypeIMAP"
	case KKCProtocolTypeIRC:
		return "KKCProtocolTypeIRC"
	case KKCProtocolTypeLDAP:
		return "KKCProtocolTypeLDAP"
	case KKCProtocolTypeNNTP:
		return "KKCProtocolTypeNNTP"
	case KKCProtocolTypePOP3:
		return "KKCProtocolTypePOP3"
	case KKCProtocolTypeSMTP:
		return "KKCProtocolTypeSMTP"
	case KKCProtocolTypeSOCKS:
		return "KKCProtocolTypeSOCKS"
	case KKCProtocolTypeTelnet:
		return "KKCProtocolTypeTelnet"
	default:
		return fmt.Sprintf("KKCProtocolType(%d)", e)
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
	KLSApplicationNotFoundErr  KLS = -10814
	KLSAttributeNotFoundErr    KLS = -10662
	KLSAttributeNotSettableErr KLS = -10663
	// KLSCannotSetInfoErr: The system can’t hide the filename extension.
	KLSCannotSetInfoErr KLS = -10823
	// KLSDataErr: Improper data structure.
	KLSDataErr KLS = -10817
	// KLSDataTooOldErr: Not currently used.
	KLSDataTooOldErr KLS = -10816
	// KLSDataUnavailableErr: Data of the desired type is not available (for example, there is no kind string).
	KLSDataUnavailableErr                KLS = -10813
	KLSExecutableIncorrectFormat         KLS = -10661
	KLSGarbageCollectionUnsupportedErr   KLS = -10666
	KLSIncompatibleApplicationVersionErr KLS = -10664
	// KLSIncompatibleSystemVersionErr: The requested app can’t run on the current macOS version.
	KLSIncompatibleSystemVersionErr KLS = -10825
	// KLSLaunchInProgressErr: A launch of the app is already in progress.
	KLSLaunchInProgressErr KLS = -10818
	KLSMalformedLocErr     KLS = -10400
	// KLSMultipleSessionsNotSupportedErr: The requested app can’t run simultaneously in two different user sessions.
	KLSMultipleSessionsNotSupportedErr KLS = -10829
	KLSNo32BitEnvironmentErr           KLS = -10386
	// KLSNoClassicEnvironmentErr: Not currently used.
	KLSNoClassicEnvironmentErr KLS = -10828
	// KLSNoExecutableErr: The executable file is missing or has an unusable format.
	KLSNoExecutableErr KLS = -10827
	// KLSNoLaunchPermissionErr: The user doesn’t have permission to launch the app on a managed network.
	KLSNoLaunchPermissionErr KLS = -10826
	// KLSNoRegistrationInfoErr: Not currently used.
	KLSNoRegistrationInfoErr   KLS = -10824
	KLSNoRosettaEnvironmentErr KLS = -10665
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
	case KLSAttributeNotFoundErr:
		return "KLSAttributeNotFoundErr"
	case KLSAttributeNotSettableErr:
		return "KLSAttributeNotSettableErr"
	case KLSCannotSetInfoErr:
		return "KLSCannotSetInfoErr"
	case KLSDataErr:
		return "KLSDataErr"
	case KLSDataTooOldErr:
		return "KLSDataTooOldErr"
	case KLSDataUnavailableErr:
		return "KLSDataUnavailableErr"
	case KLSExecutableIncorrectFormat:
		return "KLSExecutableIncorrectFormat"
	case KLSGarbageCollectionUnsupportedErr:
		return "KLSGarbageCollectionUnsupportedErr"
	case KLSIncompatibleApplicationVersionErr:
		return "KLSIncompatibleApplicationVersionErr"
	case KLSIncompatibleSystemVersionErr:
		return "KLSIncompatibleSystemVersionErr"
	case KLSLaunchInProgressErr:
		return "KLSLaunchInProgressErr"
	case KLSMalformedLocErr:
		return "KLSMalformedLocErr"
	case KLSMultipleSessionsNotSupportedErr:
		return "KLSMultipleSessionsNotSupportedErr"
	case KLSNo32BitEnvironmentErr:
		return "KLSNo32BitEnvironmentErr"
	case KLSNoClassicEnvironmentErr:
		return "KLSNoClassicEnvironmentErr"
	case KLSNoExecutableErr:
		return "KLSNoExecutableErr"
	case KLSNoLaunchPermissionErr:
		return "KLSNoLaunchPermissionErr"
	case KLSNoRegistrationInfoErr:
		return "KLSNoRegistrationInfoErr"
	case KLSNoRosettaEnvironmentErr:
		return "KLSNoRosettaEnvironmentErr"
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

type KLarge1BitMask uint

const (
	KLarge1BitMaskValue KLarge1BitMask = 'I'<<24 | 'C'<<16 | 'N'<<8 | '#' // 'ICN#'
	KLarge32BitData     KLarge1BitMask = 'i'<<24 | 'l'<<16 | '3'<<8 | '2' // 'il32'
	KLarge4BitData      KLarge1BitMask = 'i'<<24 | 'c'<<16 | 'l'<<8 | '4' // 'icl4'
	KLarge8BitData      KLarge1BitMask = 'i'<<24 | 'c'<<16 | 'l'<<8 | '8' // 'icl8'
	KLarge8BitMask      KLarge1BitMask = 'l'<<24 | '8'<<16 | 'm'<<8 | 'k' // 'l8mk'
	KMini1BitMask       KLarge1BitMask = 'i'<<24 | 'c'<<16 | 'm'<<8 | '#' // 'icm#'
	KMini4BitData       KLarge1BitMask = 'i'<<24 | 'c'<<16 | 'm'<<8 | '4' // 'icm4'
	KMini8BitData       KLarge1BitMask = 'i'<<24 | 'c'<<16 | 'm'<<8 | '8' // 'icm8'
	KSmall1BitMask      KLarge1BitMask = 'i'<<24 | 'c'<<16 | 's'<<8 | '#' // 'ics#'
	KSmall32BitData     KLarge1BitMask = 'i'<<24 | 's'<<16 | '3'<<8 | '2' // 'is32'
	KSmall4BitData      KLarge1BitMask = 'i'<<24 | 'c'<<16 | 's'<<8 | '4' // 'ics4'
	KSmall8BitData      KLarge1BitMask = 'i'<<24 | 'c'<<16 | 's'<<8 | '8' // 'ics8'
	KSmall8BitMask      KLarge1BitMask = 's'<<24 | '8'<<16 | 'm'<<8 | 'k' // 's8mk'
)

func (e KLarge1BitMask) String() string {
	switch e {
	case KLarge1BitMaskValue:
		return "KLarge1BitMaskValue"
	case KLarge32BitData:
		return "KLarge32BitData"
	case KLarge4BitData:
		return "KLarge4BitData"
	case KLarge8BitData:
		return "KLarge8BitData"
	case KLarge8BitMask:
		return "KLarge8BitMask"
	case KMini1BitMask:
		return "KMini1BitMask"
	case KMini4BitData:
		return "KMini4BitData"
	case KMini8BitData:
		return "KMini8BitData"
	case KSmall1BitMask:
		return "KSmall1BitMask"
	case KSmall32BitData:
		return "KSmall32BitData"
	case KSmall4BitData:
		return "KSmall4BitData"
	case KSmall8BitData:
		return "KSmall8BitData"
	case KSmall8BitMask:
		return "KSmall8BitMask"
	default:
		return fmt.Sprintf("KLarge1BitMask(%d)", e)
	}
}

type KLargeIconSize uint

const (
	KLarge4BitIconSize  KLargeIconSize = 512
	KLarge8BitIconSize  KLargeIconSize = 1024
	KLargeIconSizeValue KLargeIconSize = 256
	KSmall4BitIconSize  KLargeIconSize = 128
	KSmall8BitIconSize  KLargeIconSize = 256
	KSmallIconSize      KLargeIconSize = 64
)

func (e KLargeIconSize) String() string {
	switch e {
	case KLarge4BitIconSize:
		return "KLarge4BitIconSize"
	case KLarge8BitIconSize:
		return "KLarge8BitIconSize"
	case KLargeIconSizeValue:
		return "KLargeIconSizeValue"
	case KSmall4BitIconSize:
		return "KSmall4BitIconSize"
	case KSmallIconSize:
		return "KSmallIconSize"
	default:
		return fmt.Sprintf("KLargeIconSize(%d)", e)
	}
}

type KLastDomain int

const (
	// Deprecated.
	KLastDomainConstant KLastDomain = -32760
)

func (e KLastDomain) String() string {
	switch e {
	case KLastDomainConstant:
		return "KLastDomainConstant"
	default:
		return fmt.Sprintf("KLastDomain(%d)", e)
	}
}

type KLaunchToGetTerminology uint

const (
	KAlwaysSendSubject           KLaunchToGetTerminology = 8192
	KDontFindAppBySignature      KLaunchToGetTerminology = 16384
	KLaunchToGetTerminologyValue KLaunchToGetTerminology = 32768
)

func (e KLaunchToGetTerminology) String() string {
	switch e {
	case KAlwaysSendSubject:
		return "KAlwaysSendSubject"
	case KDontFindAppBySignature:
		return "KDontFindAppBySignature"
	case KLaunchToGetTerminologyValue:
		return "KLaunchToGetTerminologyValue"
	default:
		return fmt.Sprintf("KLaunchToGetTerminology(%d)", e)
	}
}

type KLocale uint

const (
	KLocaleAllPartsMask             KLocale = 63
	KLocaleAndVariantNameMask       KLocale = 3
	KLocaleLanguageMask             KLocale = 1
	KLocaleLanguageVariantMask      KLocale = 2
	KLocaleNameMask                 KLocale = 1
	KLocaleOperationVariantNameMask KLocale = 2
	KLocaleRegionMask               KLocale = 16
	KLocaleRegionVariantMask        KLocale = 32
	KLocaleScriptMask               KLocale = 4
	KLocaleScriptVariantMask        KLocale = 8
)

func (e KLocale) String() string {
	switch e {
	case KLocaleAllPartsMask:
		return "KLocaleAllPartsMask"
	case KLocaleAndVariantNameMask:
		return "KLocaleAndVariantNameMask"
	case KLocaleLanguageMask:
		return "KLocaleLanguageMask"
	case KLocaleLanguageVariantMask:
		return "KLocaleLanguageVariantMask"
	case KLocaleRegionMask:
		return "KLocaleRegionMask"
	case KLocaleRegionVariantMask:
		return "KLocaleRegionVariantMask"
	case KLocaleScriptMask:
		return "KLocaleScriptMask"
	case KLocaleScriptVariantMask:
		return "KLocaleScriptVariantMask"
	default:
		return fmt.Sprintf("KLocale(%d)", e)
	}
}

type KLocales int

const (
	KLocalesBufferTooSmallErr    KLocales = -30001
	KLocalesDefaultDisplayStatus KLocales = -30029
	KLocalesTableFormatErr       KLocales = -30002
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

type KM68kISA uint

const (
	// Deprecated.
	KM68kISAValue KM68kISA = 0
	// Deprecated.
	KPowerPCISA KM68kISA = 1
)

func (e KM68kISA) String() string {
	switch e {
	case KM68kISAValue:
		return "KM68kISAValue"
	case KPowerPCISA:
		return "KPowerPCISA"
	default:
		return fmt.Sprintf("KM68kISA(%d)", e)
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
	KMPInvalidIDErr         KMP = -29299
	KMPIterationEndErr      KMP = -29275
	KMPPrivilegedErr        KMP = -29276
	KMPProcessCreatedErr    KMP = -29288
	KMPProcessTerminatedErr KMP = -29289
	KMPTaskAbortedErr       KMP = -29297
	// KMPTaskBlockedErr: The desired task is blocked.
	KMPTaskBlockedErr KMP = -29291
	KMPTaskCreatedErr KMP = -29290
	// KMPTaskStoppedErr: The desired task is stopped.
	KMPTaskStoppedErr KMP = -29292
	// KMPTimeoutErr: The designated timeout interval passed before the function could take action.
	KMPTimeoutErr KMP = -29296
	// Deprecated.
	KMPAddressSpaceInfoVersion KMP = 1
	// Deprecated.
	KMPAllocate1024ByteAligned KMP = 10
	// Deprecated.
	KMPAllocate16ByteAligned KMP = 4
	// Deprecated.
	KMPAllocate32ByteAligned KMP = 5
	// Deprecated.
	KMPAllocate4096ByteAligned KMP = 12
	// Deprecated.
	KMPAllocate8ByteAligned KMP = 3
	// Deprecated.
	KMPAllocateAltiVecAligned KMP = 4
	// Deprecated.
	KMPAllocateDefaultAligned KMP = 0
	// Deprecated.
	KMPAllocateInterlockAligned KMP = 255
	// Deprecated.
	KMPAllocateMaxAlignment KMP = 16
	// Deprecated.
	KMPAllocateVMPageAligned KMP = 254
	// Deprecated.
	KMPAllocateVMXAligned KMP = 4
	// Deprecated.
	KMPAnyRemoteContext KMP = 0
	// Deprecated.
	KMPAsyncInterruptRemoteContext KMP = 3
	// Deprecated.
	KMPCriticalRegionInfoVersion KMP = 1
	// Deprecated.
	KMPEventInfoVersion KMP = 1
	// Deprecated.
	KMPHighLevelDebugger KMP = 536870912
	// Deprecated.
	KMPInterruptRemoteContext KMP = 2
	// Deprecated.
	KMPLowLevelDebugger KMP = 0
	// Deprecated.
	KMPMidLevelDebugger KMP = 268435456
	// Deprecated.
	KMPNotificationInfoVersion KMP = 1
	// Deprecated.
	KMPOwningProcessRemoteContext KMP = 1
	// Deprecated.
	KMPPreserveTimerIDMask KMP = 1
	// Deprecated.
	KMPQueueInfoVersion KMP = 1
	// Deprecated.
	KMPSemaphoreInfoVersion KMP = 1
	// Deprecated.
	KMPTimeIsDeltaMask KMP = 2
	// Deprecated.
	KMPTimeIsDurationMask KMP = 4
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
	case KMPAddressSpaceInfoVersion:
		return "KMPAddressSpaceInfoVersion"
	case KMPAllocate1024ByteAligned:
		return "KMPAllocate1024ByteAligned"
	case KMPAllocate16ByteAligned:
		return "KMPAllocate16ByteAligned"
	case KMPAllocate32ByteAligned:
		return "KMPAllocate32ByteAligned"
	case KMPAllocate4096ByteAligned:
		return "KMPAllocate4096ByteAligned"
	case KMPAllocate8ByteAligned:
		return "KMPAllocate8ByteAligned"
	case KMPAllocateDefaultAligned:
		return "KMPAllocateDefaultAligned"
	case KMPAllocateInterlockAligned:
		return "KMPAllocateInterlockAligned"
	case KMPAllocateMaxAlignment:
		return "KMPAllocateMaxAlignment"
	case KMPAllocateVMPageAligned:
		return "KMPAllocateVMPageAligned"
	case KMPHighLevelDebugger:
		return "KMPHighLevelDebugger"
	case KMPInterruptRemoteContext:
		return "KMPInterruptRemoteContext"
	case KMPMidLevelDebugger:
		return "KMPMidLevelDebugger"
	default:
		return fmt.Sprintf("KMP(%d)", e)
	}
}

type KMPAllocate uint

const (
	// Deprecated.
	KMPAllocateClearMask KMPAllocate = 1
	// Deprecated.
	KMPAllocateGloballyMask KMPAllocate = 2
	// Deprecated.
	KMPAllocateNoCreateMask KMPAllocate = 32
	// Deprecated.
	KMPAllocateNoGrowthMask KMPAllocate = 16
	// Deprecated.
	KMPAllocateResidentMask KMPAllocate = 4
)

func (e KMPAllocate) String() string {
	switch e {
	case KMPAllocateClearMask:
		return "KMPAllocateClearMask"
	case KMPAllocateGloballyMask:
		return "KMPAllocateGloballyMask"
	case KMPAllocateNoCreateMask:
		return "KMPAllocateNoCreateMask"
	case KMPAllocateNoGrowthMask:
		return "KMPAllocateNoGrowthMask"
	case KMPAllocateResidentMask:
		return "KMPAllocateResidentMask"
	default:
		return fmt.Sprintf("KMPAllocate(%d)", e)
	}
}

type KMPCreateTask uint

const (
	// Deprecated.
	KMPCreateTaskNotDebuggableMask KMPCreateTask = 4
	// Deprecated.
	KMPCreateTaskSuspendedMask KMPCreateTask = 1
	// Deprecated.
	KMPCreateTaskTakesAllExceptionsMask KMPCreateTask = 2
	// Deprecated.
	KMPCreateTaskValidOptionsMask KMPCreateTask = 1
)

func (e KMPCreateTask) String() string {
	switch e {
	case KMPCreateTaskNotDebuggableMask:
		return "KMPCreateTaskNotDebuggableMask"
	case KMPCreateTaskSuspendedMask:
		return "KMPCreateTaskSuspendedMask"
	case KMPCreateTaskTakesAllExceptionsMask:
		return "KMPCreateTaskTakesAllExceptionsMask"
	default:
		return fmt.Sprintf("KMPCreateTask(%d)", e)
	}
}

type KMPMaxAlloc uint

const (
	// Deprecated.
	KMPMaxAllocSize KMPMaxAlloc = 0
)

func (e KMPMaxAlloc) String() string {
	switch e {
	case KMPMaxAllocSize:
		return "KMPMaxAllocSize"
	default:
		return fmt.Sprintf("KMPMaxAlloc(%d)", e)
	}
}

type KMPNanokernelNeedsMemory int

const (
	KMPNanokernelNeedsMemoryErr KMPNanokernelNeedsMemory = -29294
)

func (e KMPNanokernelNeedsMemory) String() string {
	switch e {
	case KMPNanokernelNeedsMemoryErr:
		return "KMPNanokernelNeedsMemoryErr"
	default:
		return fmt.Sprintf("KMPNanokernelNeedsMemory(%d)", e)
	}
}

type KMPNoI uint

const (
	// Deprecated.
	KMPNoID KMPNoI = 0
)

func (e KMPNoI) String() string {
	switch e {
	case KMPNoID:
		return "KMPNoID"
	default:
		return fmt.Sprintf("KMPNoI(%d)", e)
	}
}

type KMPTask uint

const (
	// Deprecated.
	KMPTaskBlocked KMPTask = 0
	// Deprecated.
	KMPTaskPropagate KMPTask = 0
	// Deprecated.
	KMPTaskPropagateMask KMPTask = 1
	// Deprecated.
	KMPTaskReady KMPTask = 1
	// Deprecated.
	KMPTaskResumeBranch KMPTask = 2
	// Deprecated.
	KMPTaskResumeBranchMask KMPTask = 4
	// Deprecated.
	KMPTaskResumeMask KMPTask = 0
	// Deprecated.
	KMPTaskResumeStep KMPTask = 1
	// Deprecated.
	KMPTaskResumeStepMask KMPTask = 2
	// Deprecated.
	KMPTaskRunning KMPTask = 2
	// Deprecated.
	KMPTaskState32BitMemoryException KMPTask = 4
	// Deprecated.
	KMPTaskStateFPU KMPTask = 1
	// Deprecated.
	KMPTaskStateMachine KMPTask = 3
	// Deprecated.
	KMPTaskStateRegisters KMPTask = 0
	// Deprecated.
	KMPTaskStateTaskInfo KMPTask = 5
	// Deprecated.
	KMPTaskStateVectors KMPTask = 2
)

func (e KMPTask) String() string {
	switch e {
	case KMPTaskBlocked:
		return "KMPTaskBlocked"
	case KMPTaskPropagateMask:
		return "KMPTaskPropagateMask"
	case KMPTaskResumeBranch:
		return "KMPTaskResumeBranch"
	case KMPTaskResumeBranchMask:
		return "KMPTaskResumeBranchMask"
	case KMPTaskStateMachine:
		return "KMPTaskStateMachine"
	case KMPTaskStateTaskInfo:
		return "KMPTaskStateTaskInfo"
	default:
		return fmt.Sprintf("KMPTask(%d)", e)
	}
}

type KMPTaskInfo uint

const (
	// Deprecated.
	KMPTaskInfoVersion KMPTaskInfo = 3
)

func (e KMPTaskInfo) String() string {
	switch e {
	case KMPTaskInfoVersion:
		return "KMPTaskInfoVersion"
	default:
		return fmt.Sprintf("KMPTaskInfo(%d)", e)
	}
}

type KMacArabic uint

const (
	// KMacArabicAlBayanVariant: A Mac OS Arabic variant used for the Arabic TrueType font Al Bayan.
	KMacArabicAlBayanVariant KMacArabic = 3
	// KMacArabicStandardVariant: A Mac OS Arabic variant is supported by the Cairo font (the system font for Arabic) and is the encoding supported by the text processing utilities.
	KMacArabicStandardVariant KMacArabic = 0
	// KMacArabicThuluthVariant: A Mac OS Arabic variant used for the Arabic PostScript-only fonts: Thuluth and Thuluth bold.
	KMacArabicThuluthVariant KMacArabic = 2
	// KMacArabicTrueTypeVariant: A Mac OS Arabic variant used for most of the Arabic TrueType fonts: Baghdad, Geeza, Kufi, Nadeem.
	KMacArabicTrueTypeVariant KMacArabic = 1
)

func (e KMacArabic) String() string {
	switch e {
	case KMacArabicAlBayanVariant:
		return "KMacArabicAlBayanVariant"
	case KMacArabicStandardVariant:
		return "KMacArabicStandardVariant"
	case KMacArabicThuluthVariant:
		return "KMacArabicThuluthVariant"
	case KMacArabicTrueTypeVariant:
		return "KMacArabicTrueTypeVariant"
	default:
		return fmt.Sprintf("KMacArabic(%d)", e)
	}
}

type KMacCroatian uint

const (
	// KMacCroatianCurrencySignVariant: In versions of Mac OS earlier than 8.5, 0xDB is the currency sign.
	KMacCroatianCurrencySignVariant KMacCroatian = 1
	// KMacCroatianDefaultVariant: This is a meta value that maps to one of the following constants, depending on version of the Mac OS.
	KMacCroatianDefaultVariant KMacCroatian = 0
	// KMacCroatianEuroSignVariant: In Mac OS version 8.5 and later, 0xDB is the Euro sign.
	KMacCroatianEuroSignVariant KMacCroatian = 2
)

func (e KMacCroatian) String() string {
	switch e {
	case KMacCroatianCurrencySignVariant:
		return "KMacCroatianCurrencySignVariant"
	case KMacCroatianDefaultVariant:
		return "KMacCroatianDefaultVariant"
	case KMacCroatianEuroSignVariant:
		return "KMacCroatianEuroSignVariant"
	default:
		return fmt.Sprintf("KMacCroatian(%d)", e)
	}
}

type KMacCyrillic uint

const (
	// KMacCyrillicCurrSignStdVariant: In Mac OS versions prior to 9.0 (RU, BG), 0xFF = currency sign, 0xA2/0xB6 = CENT / PARTIAL DIFF.
	KMacCyrillicCurrSignStdVariant KMacCyrillic = 1
	// KMacCyrillicCurrSignUkrVariant: In Mac OS version 9.0 and later (UA, LangKit), 0xFF = currency sign, 0xA2/0xB6 = GHE with upturn.
	KMacCyrillicCurrSignUkrVariant KMacCyrillic = 2
	// KMacCyrillicDefaultVariant: This is a meta value that maps to one of the following constants, depending on version of the Mac OS.
	KMacCyrillicDefaultVariant KMacCyrillic = 0
	// KMacCyrillicEuroSignVariant: In Mac OS 9.0 and later, 0xFF is Euro sign, 0xA2/0xB6 = GHE with upturn.
	KMacCyrillicEuroSignVariant KMacCyrillic = 3
)

func (e KMacCyrillic) String() string {
	switch e {
	case KMacCyrillicCurrSignStdVariant:
		return "KMacCyrillicCurrSignStdVariant"
	case KMacCyrillicCurrSignUkrVariant:
		return "KMacCyrillicCurrSignUkrVariant"
	case KMacCyrillicDefaultVariant:
		return "KMacCyrillicDefaultVariant"
	case KMacCyrillicEuroSignVariant:
		return "KMacCyrillicEuroSignVariant"
	default:
		return fmt.Sprintf("KMacCyrillic(%d)", e)
	}
}

type KMacFarsi uint

const (
	// KMacFarsiStandardVariant: This Mac OS Farsi variant is supported by the Tehran font (the system font for Farsi) and is the encoding supported by the text processing utilities.
	KMacFarsiStandardVariant KMacFarsi = 0
	// KMacFarsiTrueTypeVariant: This Mac OS Farsi variant is used for most of the Farsi TrueType fonts: Ashfahan, Amir, Kamran, Mashad, NadeemFarsi.
	KMacFarsiTrueTypeVariant KMacFarsi = 1
)

func (e KMacFarsi) String() string {
	switch e {
	case KMacFarsiStandardVariant:
		return "KMacFarsiStandardVariant"
	case KMacFarsiTrueTypeVariant:
		return "KMacFarsiTrueTypeVariant"
	default:
		return fmt.Sprintf("KMacFarsi(%d)", e)
	}
}

type KMacGreek uint

const (
	KMacGreekDefaultVariant    KMacGreek = 0
	KMacGreekEuroSignVariant   KMacGreek = 2
	KMacGreekNoEuroSignVariant KMacGreek = 1
)

func (e KMacGreek) String() string {
	switch e {
	case KMacGreekDefaultVariant:
		return "KMacGreekDefaultVariant"
	case KMacGreekEuroSignVariant:
		return "KMacGreekEuroSignVariant"
	case KMacGreekNoEuroSignVariant:
		return "KMacGreekNoEuroSignVariant"
	default:
		return fmt.Sprintf("KMacGreek(%d)", e)
	}
}

type KMacHebrew uint

const (
	// KMacHebrewFigureSpaceVariant: The Mac OS Hebrew variant in which 0xD4 represents figure space, not left single quotation mark as in the standard variant.
	KMacHebrewFigureSpaceVariant KMacHebrew = 1
	// KMacHebrewStandardVariant: The standard Mac OS Hebrew variant.
	KMacHebrewStandardVariant KMacHebrew = 0
)

func (e KMacHebrew) String() string {
	switch e {
	case KMacHebrewFigureSpaceVariant:
		return "KMacHebrewFigureSpaceVariant"
	case KMacHebrewStandardVariant:
		return "KMacHebrewStandardVariant"
	default:
		return fmt.Sprintf("KMacHebrew(%d)", e)
	}
}

type KMacIcelandic uint

const (
	// KMacIcelandicStdCurrSignVariant: In Mac OS versions prior to 8.5, 0xDB is the currency sign; 0xBB/0xBC are fem./masc.
	KMacIcelandicStdCurrSignVariant KMacIcelandic = 2
	// KMacIcelandicStdDefaultVariant: This is a meta value that maps to `kMacIcelandicStdCurrSignVariant` or `kMacIcelandicStdEuroSignVariant`, depending on version of the Mac OS.
	KMacIcelandicStdDefaultVariant KMacIcelandic = 0
	// KMacIcelandicStdEuroSignVariant: In Mac OS version 8.5 and later , 0xDB is the Euro sign ; 0xBB/0xBC are fem./masc.
	KMacIcelandicStdEuroSignVariant KMacIcelandic = 4
	// KMacIcelandicTTCurrSignVariant: In Mac OS versions prior to 8.5, 0xDB is the currency sign; 0xBB/0xBC are fi/fl ligatures
	KMacIcelandicTTCurrSignVariant KMacIcelandic = 3
	// KMacIcelandicTTDefaultVariant: This is a meta value that maps to `kMacIcelandicTTCurrSignVariant` or `kMacIcelandicTTEuroSignVariant`, depending on version of the Mac OS.
	KMacIcelandicTTDefaultVariant KMacIcelandic = 1
	// KMacIcelandicTTEuroSignVariant: In Mac OS versions earlier than 8.5, 0xDB is the Euro sign; 0xBB/0xBC are fi/fl ligatures.
	KMacIcelandicTTEuroSignVariant KMacIcelandic = 5
)

func (e KMacIcelandic) String() string {
	switch e {
	case KMacIcelandicStdCurrSignVariant:
		return "KMacIcelandicStdCurrSignVariant"
	case KMacIcelandicStdDefaultVariant:
		return "KMacIcelandicStdDefaultVariant"
	case KMacIcelandicStdEuroSignVariant:
		return "KMacIcelandicStdEuroSignVariant"
	case KMacIcelandicTTCurrSignVariant:
		return "KMacIcelandicTTCurrSignVariant"
	case KMacIcelandicTTDefaultVariant:
		return "KMacIcelandicTTDefaultVariant"
	case KMacIcelandicTTEuroSignVariant:
		return "KMacIcelandicTTEuroSignVariant"
	default:
		return fmt.Sprintf("KMacIcelandic(%d)", e)
	}
}

type KMacJapanese uint

const (
	// KMacJapaneseBasicVariant: An artificial Mac OS Japanese variant without Apple double-byte extensions.
	KMacJapaneseBasicVariant KMacJapanese = 2
	// KMacJapanesePostScriptPrintVariant: The Mac OS Japanese variant for PostScript printing versions of the Sai Mincho and Chu Gothic PostScript fonts.
	KMacJapanesePostScriptPrintVariant KMacJapanese = 4
	// KMacJapanesePostScriptScrnVariant: The Mac OS Japanese variant for the screen bitmap version of the Sai Mincho and Chu Gothic fonts.
	KMacJapanesePostScriptScrnVariant KMacJapanese = 3
	// KMacJapaneseStandardVariant: The standard Mac OS Japanese variant.
	KMacJapaneseStandardVariant KMacJapanese = 0
	// KMacJapaneseStdNoVerticalsVariant: An artificial Mac OS Japanese variant for callers who don’t want to use separately encoded vertical forms (for example, developers using QuickDraw GX).
	KMacJapaneseStdNoVerticalsVariant KMacJapanese = 1
	// KMacJapaneseVertAtKuPlusTenVariant: The Mac OS Japanese variant for the Hon Mincho and Maru Gothic fonts used in the Japanese localized version of System 7.1.
	KMacJapaneseVertAtKuPlusTenVariant KMacJapanese = 5
)

func (e KMacJapanese) String() string {
	switch e {
	case KMacJapaneseBasicVariant:
		return "KMacJapaneseBasicVariant"
	case KMacJapanesePostScriptPrintVariant:
		return "KMacJapanesePostScriptPrintVariant"
	case KMacJapanesePostScriptScrnVariant:
		return "KMacJapanesePostScriptScrnVariant"
	case KMacJapaneseStandardVariant:
		return "KMacJapaneseStandardVariant"
	case KMacJapaneseStdNoVerticalsVariant:
		return "KMacJapaneseStdNoVerticalsVariant"
	case KMacJapaneseVertAtKuPlusTenVariant:
		return "KMacJapaneseVertAtKuPlusTenVariant"
	default:
		return fmt.Sprintf("KMacJapanese(%d)", e)
	}
}

type KMacMemoryMaximumMemoryManagerBlock uint

const (
	// Deprecated.
	KMacMemoryMaximumMemoryManagerBlockSize KMacMemoryMaximumMemoryManagerBlock = 2147483632
)

func (e KMacMemoryMaximumMemoryManagerBlock) String() string {
	switch e {
	case KMacMemoryMaximumMemoryManagerBlockSize:
		return "KMacMemoryMaximumMemoryManagerBlockSize"
	default:
		return fmt.Sprintf("KMacMemoryMaximumMemoryManagerBlock(%d)", e)
	}
}

type KMacRoman uint

const (
	// KMacRomanCurrencySignVariant: In Mac OS versions earlier than 8.5 0xDB is the currency sign; still used for some older fonts even in Mac OS 8.5.
	KMacRomanCurrencySignVariant KMacRoman = 1
	// KMacRomanDefaultVariant: This is a meta value that maps to one of the following constants, depending on version of the Mac OS.
	KMacRomanDefaultVariant KMacRoman = 0
	// KMacRomanEuroSignVariant: In Mac OS version 8.5 and later, 0xDB is the Euro sign.
	KMacRomanEuroSignVariant KMacRoman = 2
)

func (e KMacRoman) String() string {
	switch e {
	case KMacRomanCurrencySignVariant:
		return "KMacRomanCurrencySignVariant"
	case KMacRomanDefaultVariant:
		return "KMacRomanDefaultVariant"
	case KMacRomanEuroSignVariant:
		return "KMacRomanEuroSignVariant"
	default:
		return fmt.Sprintf("KMacRoman(%d)", e)
	}
}

type KMacRomanLatin1 uint

const (
	// KMacRomanLatin1CroatianVariant: Permuted MacCroatian, Euro sign variant.
	KMacRomanLatin1CroatianVariant KMacRomanLatin1 = 8
	// KMacRomanLatin1DefaultVariant: This is a meta value that maps to one of the following constants, depending on version of the Mac OS.
	KMacRomanLatin1DefaultVariant KMacRomanLatin1 = 0
	// KMacRomanLatin1IcelandicVariant: Permuted MacIcelandic, standard Euro sign variant.
	KMacRomanLatin1IcelandicVariant KMacRomanLatin1 = 11
	// KMacRomanLatin1RomanianVariant: Permuted MacRomanian, Euro sign variant.
	KMacRomanLatin1RomanianVariant KMacRomanLatin1 = 14
	// KMacRomanLatin1StandardVariant: Permuted MacRoman, Euro sign variant.
	KMacRomanLatin1StandardVariant KMacRomanLatin1 = 2
	// KMacRomanLatin1TurkishVariant: Permuted MacTurkish.
	KMacRomanLatin1TurkishVariant KMacRomanLatin1 = 6
)

func (e KMacRomanLatin1) String() string {
	switch e {
	case KMacRomanLatin1CroatianVariant:
		return "KMacRomanLatin1CroatianVariant"
	case KMacRomanLatin1DefaultVariant:
		return "KMacRomanLatin1DefaultVariant"
	case KMacRomanLatin1IcelandicVariant:
		return "KMacRomanLatin1IcelandicVariant"
	case KMacRomanLatin1RomanianVariant:
		return "KMacRomanLatin1RomanianVariant"
	case KMacRomanLatin1StandardVariant:
		return "KMacRomanLatin1StandardVariant"
	case KMacRomanLatin1TurkishVariant:
		return "KMacRomanLatin1TurkishVariant"
	default:
		return fmt.Sprintf("KMacRomanLatin1(%d)", e)
	}
}

type KMacRomanStandardVariant uint

const (
	KHebrewFigureSpaceVariant       KMacRomanStandardVariant = 1
	KHebrewStandardVariant          KMacRomanStandardVariant = 0
	KJapaneseBasicVariant           KMacRomanStandardVariant = 2
	KJapanesePostScriptPrintVariant KMacRomanStandardVariant = 4
	KJapanesePostScriptScrnVariant  KMacRomanStandardVariant = 3
	KJapaneseStandardVariant        KMacRomanStandardVariant = 0
	KJapaneseStdNoVerticalsVariant  KMacRomanStandardVariant = 1
	KJapaneseVertAtKuPlusTenVariant KMacRomanStandardVariant = 5
	// KMacIcelandicStandardVariant: The standard Mac OS Icelandic encoding supported by the bitmap versions of Chicago, Geneva, Monaco, and New York in the Icelandic system.
	KMacIcelandicStandardVariant KMacRomanStandardVariant = 0
	// KMacIcelandicTrueTypeVariant: The Mac OS Icelandic variant used for the bitmap versions of Courier, Helvetica, Palatino, and Times in the Icelandic system, and for the TrueType versions of Chicago, Geneva, Monaco, New York, Courier, Helvetica, Palatino, and Times.
	KMacIcelandicTrueTypeVariant KMacRomanStandardVariant = 1
	// KMacRomanStandardVariantValue: The standard variant of Mac OS Roman for Mac OS 8.5 and later; 0xDB is the Euro sign.
	KMacRomanStandardVariantValue KMacRomanStandardVariant = 0
	// KTextEncodingShiftJIS_X0213_00: Shift-JIS format encoding of JIS X0213 planes 1 and 2
	KTextEncodingShiftJIS_X0213_00 KMacRomanStandardVariant = 1576
	// KUnicodeCanonicalCompVariant: This is the normal canonical composition according to Unicode 3.2 rules.
	KUnicodeCanonicalCompVariant KMacRomanStandardVariant = 3
	// KUnicodeCanonicalDecompVariant: # Discussion
	KUnicodeCanonicalDecompVariant KMacRomanStandardVariant = 2
	// KUnicodeMaxDecomposedVariant: Replaced by `kUnicodeCanonicalDecompVariant`.
	KUnicodeMaxDecomposedVariant KMacRomanStandardVariant = 2
	// KUnicodeNoComposedVariant: Replaced by `kUnicodeCanonicalCompVariant`.
	KUnicodeNoComposedVariant KMacRomanStandardVariant = 3
)

func (e KMacRomanStandardVariant) String() string {
	switch e {
	case KHebrewFigureSpaceVariant:
		return "KHebrewFigureSpaceVariant"
	case KHebrewStandardVariant:
		return "KHebrewStandardVariant"
	case KJapaneseBasicVariant:
		return "KJapaneseBasicVariant"
	case KJapanesePostScriptPrintVariant:
		return "KJapanesePostScriptPrintVariant"
	case KJapanesePostScriptScrnVariant:
		return "KJapanesePostScriptScrnVariant"
	case KJapaneseVertAtKuPlusTenVariant:
		return "KJapaneseVertAtKuPlusTenVariant"
	case KTextEncodingShiftJIS_X0213_00:
		return "KTextEncodingShiftJIS_X0213_00"
	default:
		return fmt.Sprintf("KMacRomanStandardVariant(%d)", e)
	}
}

type KMacRomanian uint

const (
	// KMacRomanianCurrencySignVariant: In Mac OS versions earlier than 8.5, 0xDB is the currency sign.
	KMacRomanianCurrencySignVariant KMacRomanian = 1
	// KMacRomanianDefaultVariant: This is a meta value that maps to one of the following constants, depending on version of the Mac OS.
	KMacRomanianDefaultVariant KMacRomanian = 0
	// KMacRomanianEuroSignVariant: In Mac OS version 8.5 and later, 0xDB is the Euro sign.
	KMacRomanianEuroSignVariant KMacRomanian = 2
)

func (e KMacRomanian) String() string {
	switch e {
	case KMacRomanianCurrencySignVariant:
		return "KMacRomanianCurrencySignVariant"
	case KMacRomanianDefaultVariant:
		return "KMacRomanianDefaultVariant"
	case KMacRomanianEuroSignVariant:
		return "KMacRomanianEuroSignVariant"
	default:
		return fmt.Sprintf("KMacRomanian(%d)", e)
	}
}

type KMacVT100 uint

const (
	// KMacVT100CurrencySignVariant: In Mac OS versions earlier than 8.5, 0xDB is the currency sign.
	KMacVT100CurrencySignVariant KMacVT100 = 1
	// KMacVT100DefaultVariant: This is a meta value that maps to one of the following constants, depending on version of the Mac OS.
	KMacVT100DefaultVariant KMacVT100 = 0
	// KMacVT100EuroSignVariant: In Mac OS version 8.5 and later, 0xDB is the Euro sign.
	KMacVT100EuroSignVariant KMacVT100 = 2
)

func (e KMacVT100) String() string {
	switch e {
	case KMacVT100CurrencySignVariant:
		return "KMacVT100CurrencySignVariant"
	case KMacVT100DefaultVariant:
		return "KMacVT100DefaultVariant"
	case KMacVT100EuroSignVariant:
		return "KMacVT100EuroSignVariant"
	default:
		return fmt.Sprintf("KMacVT100(%d)", e)
	}
}

type KMachineNameStrI int

const (
	KMachineNameStrID KMachineNameStrI = -16395
)

func (e KMachineNameStrI) String() string {
	switch e {
	case KMachineNameStrID:
		return "KMachineNameStrID"
	default:
		return fmt.Sprintf("KMachineNameStrI(%d)", e)
	}
}

type KMagicBusyCreation uint

const (
	KMagicBusyCreationDate KMagicBusyCreation = 1329266096
)

func (e KMagicBusyCreation) String() string {
	switch e {
	case KMagicBusyCreationDate:
		return "KMagicBusyCreationDate"
	default:
		return fmt.Sprintf("KMagicBusyCreation(%d)", e)
	}
}

type KMagicTemporaryItemsFolderType uint

const (
	// Deprecated.
	KCurrentUserRemoteFolderLocation KMagicTemporaryItemsFolderType = 'r'<<24 | 'u'<<16 | 's'<<8 | 'f' // 'rusf'
	// Deprecated.
	KCurrentUserRemoteFolderType KMagicTemporaryItemsFolderType = 'r'<<24 | 'u'<<16 | 's'<<8 | 'r' // 'rusr'
	// Deprecated.
	KMagicTemporaryItemsFolderTypeValue KMagicTemporaryItemsFolderType = 'm'<<24 | 't'<<16 | 'm'<<8 | 'p' // 'mtmp'
	// Deprecated.
	KTemporaryItemsInUserDomainFolderType KMagicTemporaryItemsFolderType = 't'<<24 | 'e'<<16 | 'm'<<8 | 'q' // 'temq'
)

func (e KMagicTemporaryItemsFolderType) String() string {
	switch e {
	case KCurrentUserRemoteFolderLocation:
		return "KCurrentUserRemoteFolderLocation"
	case KCurrentUserRemoteFolderType:
		return "KCurrentUserRemoteFolderType"
	case KMagicTemporaryItemsFolderTypeValue:
		return "KMagicTemporaryItemsFolderTypeValue"
	case KTemporaryItemsInUserDomainFolderType:
		return "KTemporaryItemsInUserDomainFolderType"
	default:
		return fmt.Sprintf("KMagicTemporaryItemsFolderType(%d)", e)
	}
}

type KModem int

const (
	KModemOutOfMemory        KModem = -14000
	KModemPreferencesMissing KModem = -14001
	KModemScriptMissing      KModem = -14002
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
	KNSL68kContextNotSupported     KNS = -4170
	KNSLBadClientInfoPtr           KNS = -4185
	KNSLBadDataTypeErr             KNS = -4193
	KNSLBadNetConnection           KNS = -4192
	KNSLBadProtocolTypeErr         KNS = -4183
	KNSLBadReferenceErr            KNS = -4195
	KNSLBadServiceTypeErr          KNS = -4194
	KNSLBadURLSyntax               KNS = -4172
	KNSLBufferTooSmallForData      KNS = -4187
	KNSLCannotContinueLookup       KNS = -4186
	KNSLErrNullPtrError            KNS = -4176
	KNSLInitializationFailed       KNS = -4200
	KNSLInsufficientOTVer          KNS = -4197
	KNSLInsufficientSysVer         KNS = -4198
	KNSLInvalidPluginSpec          KNS = -4190
	KNSLNoCarbonLib                KNS = -4173
	KNSLNoContextAvailable         KNS = -4188
	KNSLNoElementsInList           KNS = -4196
	KNSLNoPluginsForSearch         KNS = -4179
	KNSLNoPluginsFound             KNS = -4181
	KNSLNoSupportForService        KNS = -4191
	KNSLNotImplementedYet          KNS = -4175
	KNSLNotInitialized             KNS = -4199
	KNSLNullListPtr                KNS = -4184
	KNSLNullNeighborhoodPtr        KNS = -4178
	KNSLPluginLoadFailed           KNS = -4182
	KNSLRequestBufferAlreadyInList KNS = -4189
	KNSLSchedulerError             KNS = -4171
	KNSLSearchAlreadyInProgress    KNS = -4180
	KNSLSomePluginsFailedToLoad    KNS = -4177
	KNSLUILibraryNotAvailable      KNS = -4174
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
	KNSpAddPlayerFailedErr       KNSp = -30389
	KNSpAddressInUseErr          KNSp = -30380
	KNSpAlreadyAdvertisingErr    KNSp = -30374
	KNSpAlreadyInitializedErr    KNSp = -30361
	KNSpCantBlockErr             KNSp = -30398
	KNSpConnectFailedErr         KNSp = -30395
	KNSpCreateGroupFailedErr     KNSp = -30388
	KNSpFeatureNotImplementedErr KNSp = -30381
	KNSpFreeQExhaustedErr        KNSp = -30378
	KNSpGameTerminatedErr        KNSp = -30394
	KNSpHostFailedErr            KNSp = -30365
	KNSpInitializationFailedErr  KNSp = -30360
	KNSpInvalidAddressErr        KNSp = -30377
	KNSpInvalidDefinitionErr     KNSp = -30390
	KNSpInvalidGameRefErr        KNSp = -30367
	KNSpInvalidGroupIDErr        KNSp = -30384
	KNSpInvalidParameterErr      KNSp = -30369
	KNSpInvalidPlayerIDErr       KNSp = -30383
	KNSpInvalidProtocolListErr   KNSp = -30392
	KNSpInvalidProtocolRefErr    KNSp = -30391
	KNSpJoinFailedErr            KNSp = -30399
	KNSpMemAllocationErr         KNSp = -30373
	KNSpMessageTooBigErr         KNSp = -30397
	KNSpNameRequiredErr          KNSp = -30382
	KNSpNoGroupsErr              KNSp = -30386
	KNSpNoHostVolunteersErr      KNSp = -30387
	KNSpNoPlayersErr             KNSp = -30385
	KNSpNotAdvertisingErr        KNSp = -30376
	KNSpOTNotPresentErr          KNSp = -30370
	KNSpOTVersionTooOldErr       KNSp = -30371
	KNSpPipeFullErr              KNSp = -30364
	KNSpProtocolNotAvailableErr  KNSp = -30366
	KNSpRemovePlayerFailedErr    KNSp = -30379
	KNSpSendFailedErr            KNSp = -30396
	KNSpTimeoutErr               KNSp = -30393
	KNSpTopologyNotSupportedErr  KNSp = -30362
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
	KNavCustomControlMessageFailedErr  KNav = -5697
	KNavInvalidCustomControlMessageErr KNav = -5698
	KNavInvalidSystemConfigErr         KNav = -5696
	KNavMissingKindStringErr           KNav = -5699
	KNavWrongDialogClassErr            KNav = -5695
	KNavWrongDialogStateErr            KNav = -5694
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

type KNewSuspend uint

const (
	// Deprecated.
	KCreateIfNeeded KNewSuspend = 4
	// Deprecated.
	KExactMatchThread KNewSuspend = 16
	// Deprecated.
	KFPUNotNeeded KNewSuspend = 8
	// Deprecated.
	KNewSuspendValue KNewSuspend = 1
	// Deprecated.
	KUsePremadeThread KNewSuspend = 2
)

func (e KNewSuspend) String() string {
	switch e {
	case KCreateIfNeeded:
		return "KCreateIfNeeded"
	case KExactMatchThread:
		return "KExactMatchThread"
	case KFPUNotNeeded:
		return "KFPUNotNeeded"
	case KNewSuspendValue:
		return "KNewSuspendValue"
	case KUsePremadeThread:
		return "KUsePremadeThread"
	default:
		return fmt.Sprintf("KNewSuspend(%d)", e)
	}
}

type KNextBody uint

const (
	KNextBodyValue KNextBody = 1
	KPreviousBody  KNextBody = 2
)

func (e KNextBody) String() string {
	switch e {
	case KNextBodyValue:
		return "KNextBodyValue"
	case KPreviousBody:
		return "KPreviousBody"
	default:
		return fmt.Sprintf("KNextBody(%d)", e)
	}
}

type KNoByteCode uint

const (
	// Deprecated.
	KFourByteCode KNoByteCode = 3
	// Deprecated.
	KNoByteCodeValue KNoByteCode = 0
	// Deprecated.
	KOneByteCode KNoByteCode = 1
	// Deprecated.
	KTwoByteCode KNoByteCode = 2
)

func (e KNoByteCode) String() string {
	switch e {
	case KFourByteCode:
		return "KFourByteCode"
	case KNoByteCodeValue:
		return "KNoByteCodeValue"
	case KOneByteCode:
		return "KOneByteCode"
	case KTwoByteCode:
		return "KTwoByteCode"
	default:
		return fmt.Sprintf("KNoByteCode(%d)", e)
	}
}

type KNoThreadID uint

const (
	// Deprecated.
	KApplicationThreadID KNoThreadID = 2
	// Deprecated.
	KCurrentThreadID KNoThreadID = 1
	// Deprecated.
	KNoThreadIDValue KNoThreadID = 0
)

func (e KNoThreadID) String() string {
	switch e {
	case KApplicationThreadID:
		return "KApplicationThreadID"
	case KCurrentThreadID:
		return "KCurrentThreadID"
	case KNoThreadIDValue:
		return "KNoThreadIDValue"
	default:
		return fmt.Sprintf("KNoThreadID(%d)", e)
	}
}

type KNoUserAuthentication uint

const (
	KEncryptPassword           KNoUserAuthentication = 3
	KNoUserAuthenticationValue KNoUserAuthentication = 1
	KPassword                  KNoUserAuthentication = 2
	KTwoWayEncryptPassword     KNoUserAuthentication = 6
)

func (e KNoUserAuthentication) String() string {
	switch e {
	case KEncryptPassword:
		return "KEncryptPassword"
	case KNoUserAuthenticationValue:
		return "KNoUserAuthenticationValue"
	case KPassword:
		return "KPassword"
	case KTwoWayEncryptPassword:
		return "KTwoWayEncryptPassword"
	default:
		return fmt.Sprintf("KNoUserAuthentication(%d)", e)
	}
}

type KOSI uint

const (
	KOSIZCodeInSharedLibraries  KOSI = 11
	KOSIZDontOpenResourceFile   KOSI = 15
	KOSIZOpenWithReadPermission KOSI = 13
	KOSIZdontAcceptRemoteEvents KOSI = 14
)

func (e KOSI) String() string {
	switch e {
	case KOSIZCodeInSharedLibraries:
		return "KOSIZCodeInSharedLibraries"
	case KOSIZDontOpenResourceFile:
		return "KOSIZDontOpenResourceFile"
	case KOSIZOpenWithReadPermission:
		return "KOSIZOpenWithReadPermission"
	case KOSIZdontAcceptRemoteEvents:
		return "KOSIZdontAcceptRemoteEvents"
	default:
		return fmt.Sprintf("KOSI(%d)", e)
	}
}

type KOTNoError int

const (
	KEACCESErr                 KOTNoError = -3212
	KEADDRINUSEErr             KOTNoError = -3247
	KEADDRNOTAVAILErr          KOTNoError = -3248
	KEAGAINErr                 KOTNoError = -3210
	KEALREADYErr               KOTNoError = -3236
	KEBADFErr                  KOTNoError = -3208
	KEBADMSGErr                KOTNoError = -3272
	KEBUSYErr                  KOTNoError = -3215
	KECANCELErr                KOTNoError = -3273
	KECONNABORTEDErr           KOTNoError = -3252
	KECONNREFUSEDErr           KOTNoError = -3260
	KECONNRESETErr             KOTNoError = -3253
	KEDEADLKErr                KOTNoError = -3234
	KEDESTADDRREQErr           KOTNoError = -3238
	KEEXISTErr                 KOTNoError = -3216
	KEFAULTErr                 KOTNoError = -3213
	KEHOSTDOWNErr              KOTNoError = -3263
	KEHOSTUNREACHErr           KOTNoError = -3264
	KEINPROGRESSErr            KOTNoError = -3276
	KEINTRErr                  KOTNoError = -3203
	KEINVALErr                 KOTNoError = -3221
	KEIOErr                    KOTNoError = -3204
	KEISCONNErr                KOTNoError = -3255
	KEMSGSIZEErr               KOTNoError = -3239
	KENETDOWNErr               KOTNoError = -3249
	KENETRESETErr              KOTNoError = -3251
	KENETUNREACHErr            KOTNoError = -3250
	KENOBUFSErr                KOTNoError = -3254
	KENODATAErr                KOTNoError = -3275
	KENODEVErr                 KOTNoError = -3218
	KENOENTErr                 KOTNoError = -3201
	KENOMEMErr                 KOTNoError = -3211
	KENOMSGErr                 KOTNoError = -3278
	KENOPROTOOPTErr            KOTNoError = -3241
	KENORSRCErr                KOTNoError = -3202
	KENOSRErr                  KOTNoError = -3271
	KENOSTRErr                 KOTNoError = -3274
	KENOTCONNErr               KOTNoError = -3256
	KENOTSOCKErr               KOTNoError = -3237
	KENOTTYErr                 KOTNoError = -3224
	KENXIOErr                  KOTNoError = -3205
	KEOPNOTSUPPErr             KOTNoError = -3244
	KEPERMErr                  KOTNoError = -3200
	KEPIPEErr                  KOTNoError = -3231
	KEPROTOErr                 KOTNoError = -3269
	KEPROTONOSUPPORTErr        KOTNoError = -3242
	KEPROTOTYPEErr             KOTNoError = -3240
	KERANGEErr                 KOTNoError = -3233
	KESHUTDOWNErr              KOTNoError = -3257
	KESOCKTNOSUPPORTErr        KOTNoError = -3243
	KESRCHErr                  KOTNoError = -3277
	KETIMEDOUTErr              KOTNoError = -3259
	KETIMEErr                  KOTNoError = -3270
	KETOOMANYREFSErr           KOTNoError = -3258
	KEWOULDBLOCKErr            KOTNoError = -3234
	KOTAccessErr               KOTNoError = -3152
	KOTAddressBusyErr          KOTNoError = -3172
	KOTBadAddressErr           KOTNoError = -3150
	KOTBadConfigurationErr     KOTNoError = -3282
	KOTBadDataErr              KOTNoError = -3159
	KOTBadFlagErr              KOTNoError = -3165
	KOTBadNameErr              KOTNoError = -3170
	KOTBadOptionErr            KOTNoError = -3151
	KOTBadQLenErr              KOTNoError = -3171
	KOTBadReferenceErr         KOTNoError = -3153
	KOTBadSequenceErr          KOTNoError = -3156
	KOTBadSyncErr              KOTNoError = -3179
	KOTBufferOverflowErr       KOTNoError = -3160
	KOTCanceledErr             KOTNoError = -3180
	KOTClientNotInittedErr     KOTNoError = -3279
	KOTConfigurationChangedErr KOTNoError = -3283
	KOTDuplicateFoundErr       KOTNoError = -3216
	KOTFlowErr                 KOTNoError = -3161
	KOTIndOutErr               KOTNoError = -3173
	KOTLookErr                 KOTNoError = -3158
	KOTNoAddressErr            KOTNoError = -3154
	KOTNoDataErr               KOTNoError = -3162
	KOTNoDisconnectErr         KOTNoError = -3163
	KOTNoErrorValue            KOTNoError = 0
	KOTNoReleaseErr            KOTNoError = -3166
	KOTNoStructureTypeErr      KOTNoError = -3169
	KOTNoUDErrErr              KOTNoError = -3164
	KOTNotFoundErr             KOTNoError = -3201
	KOTNotSupportedErr         KOTNoError = -3167
	KOTOutOfMemoryErr          KOTNoError = -3211
	KOTOutStateErr             KOTNoError = -3155
	KOTPortHasDiedErr          KOTNoError = -3280
	KOTPortLostConnection      KOTNoError = -3285
	KOTPortWasEjectedErr       KOTNoError = -3281
	KOTProtocolErr             KOTNoError = -3178
	KOTProviderMismatchErr     KOTNoError = -3174
	KOTQFullErr                KOTNoError = -3177
	KOTResAddressErr           KOTNoError = -3176
	KOTResQLenErr              KOTNoError = -3175
	KOTStateChangeErr          KOTNoError = -3168
	KOTSysErrorErr             KOTNoError = -3157
	KOTUserRequestedErr        KOTNoError = -3284
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

type KOld68kRTA uint

const (
	// Deprecated.
	KCFM68kRTA KOld68kRTA = 16
	// Deprecated.
	KOld68kRTAValue KOld68kRTA = 0
	// Deprecated.
	KPowerPCRTA KOld68kRTA = 0
)

func (e KOld68kRTA) String() string {
	switch e {
	case KCFM68kRTA:
		return "KCFM68kRTA"
	case KOld68kRTAValue:
		return "KOld68kRTAValue"
	default:
		return fmt.Sprintf("KOld68kRTA(%d)", e)
	}
}

type KOnSystemDisk int

const (
	// Deprecated.
	KFolderManagerLastDomain KOnSystemDisk = -32760
	// Deprecated.
	KLocalDomain KOnSystemDisk = -32765
	// Deprecated.
	KNetworkDomain KOnSystemDisk = -32764
	// Deprecated.
	KOnAppropriateDisk KOnSystemDisk = -32767
	// Deprecated.
	KOnSystemDiskValue KOnSystemDisk = -32768
	// Deprecated.
	KSystemDomain KOnSystemDisk = -32766
	// Deprecated.
	KUserDomain KOnSystemDisk = -32763
)

func (e KOnSystemDisk) String() string {
	switch e {
	case KFolderManagerLastDomain:
		return "KFolderManagerLastDomain"
	case KLocalDomain:
		return "KLocalDomain"
	case KNetworkDomain:
		return "KNetworkDomain"
	case KOnAppropriateDisk:
		return "KOnAppropriateDisk"
	case KOnSystemDiskValue:
		return "KOnSystemDiskValue"
	case KSystemDomain:
		return "KSystemDomain"
	case KUserDomain:
		return "KUserDomain"
	default:
		return fmt.Sprintf("KOnSystemDisk(%d)", e)
	}
}

type KOpaque uint

const (
	// Deprecated.
	KOpaqueAddressSpaceID KOpaque = 8
	// Deprecated.
	KOpaqueAnyID KOpaque = 0
	// Deprecated.
	KOpaqueAreaID KOpaque = 11
	// Deprecated.
	KOpaqueCoherenceID KOpaque = 10
	// Deprecated.
	KOpaqueConsoleID KOpaque = 13
	// Deprecated.
	KOpaqueCpuID KOpaque = 7
	// Deprecated.
	KOpaqueCriticalRegionID KOpaque = 6
	// Deprecated.
	KOpaqueEventID KOpaque = 9
	// Deprecated.
	KOpaqueNotificationID KOpaque = 12
	// Deprecated.
	KOpaqueProcessID KOpaque = 1
	// Deprecated.
	KOpaqueQueueID KOpaque = 4
	// Deprecated.
	KOpaqueSemaphoreID KOpaque = 5
	// Deprecated.
	KOpaqueTaskID KOpaque = 2
	// Deprecated.
	KOpaqueTimerID KOpaque = 3
)

func (e KOpaque) String() string {
	switch e {
	case KOpaqueAddressSpaceID:
		return "KOpaqueAddressSpaceID"
	case KOpaqueAnyID:
		return "KOpaqueAnyID"
	case KOpaqueAreaID:
		return "KOpaqueAreaID"
	case KOpaqueCoherenceID:
		return "KOpaqueCoherenceID"
	case KOpaqueConsoleID:
		return "KOpaqueConsoleID"
	case KOpaqueCpuID:
		return "KOpaqueCpuID"
	case KOpaqueCriticalRegionID:
		return "KOpaqueCriticalRegionID"
	case KOpaqueEventID:
		return "KOpaqueEventID"
	case KOpaqueNotificationID:
		return "KOpaqueNotificationID"
	case KOpaqueProcessID:
		return "KOpaqueProcessID"
	case KOpaqueQueueID:
		return "KOpaqueQueueID"
	case KOpaqueSemaphoreID:
		return "KOpaqueSemaphoreID"
	case KOpaqueTaskID:
		return "KOpaqueTaskID"
	case KOpaqueTimerID:
		return "KOpaqueTimerID"
	default:
		return fmt.Sprintf("KOpaque(%d)", e)
	}
}

type KOwnerID2Name uint

const (
	KGroupID2Name      KOwnerID2Name = 2
	KGroupName2ID      KOwnerID2Name = 4
	KOwnerID2NameValue KOwnerID2Name = 1
	KOwnerName2ID      KOwnerID2Name = 3
	KReturnNextGroup   KOwnerID2Name = 2
	KReturnNextUG      KOwnerID2Name = 3
	KReturnNextUser    KOwnerID2Name = 1
)

func (e KOwnerID2Name) String() string {
	switch e {
	case KGroupID2Name:
		return "KGroupID2Name"
	case KGroupName2ID:
		return "KGroupName2ID"
	case KOwnerID2NameValue:
		return "KOwnerID2NameValue"
	case KOwnerName2ID:
		return "KOwnerName2ID"
	default:
		return fmt.Sprintf("KOwnerID2Name(%d)", e)
	}
}

type KPEF int

const (
	// Deprecated.
	KPEFAbsoluteExport KPEF = -2
	// Deprecated.
	KPEFCodeSection KPEF = 0
	// Deprecated.
	KPEFCodeSymbol KPEF = 0
	// Deprecated.
	KPEFConstantSection KPEF = 3
	// Deprecated.
	KPEFDataSymbol KPEF = 1
	// Deprecated.
	KPEFDebugSection KPEF = 5
	// Deprecated.
	KPEFExceptionSection KPEF = 7
	// Deprecated.
	KPEFExecDataSection KPEF = 6
	// Deprecated.
	KPEFGlobalShare KPEF = 4
	// Deprecated.
	KPEFGlueSymbol KPEF = 4
	// Deprecated.
	KPEFInitLibBeforeMask KPEF = 128
	// Deprecated.
	KPEFLoaderSection KPEF = 4
	// Deprecated.
	KPEFPackedDataSection KPEF = 2
	// Deprecated.
	KPEFProcessShare KPEF = 1
	// Deprecated.
	KPEFProtectedShare KPEF = 5
	// Deprecated.
	KPEFReexportedImport KPEF = -3
	// Deprecated.
	KPEFTOCSymbol KPEF = 3
	// Deprecated.
	KPEFTVectorSymbol KPEF = 2
	// Deprecated.
	KPEFTag1 KPEF = 'J'<<24 | 'o'<<16 | 'y'<<8 | '!' // 'Joy!'
	// Deprecated.
	KPEFTag2 KPEF = 'p'<<24 | 'e'<<16 | 'f'<<8 | 'f' // 'peff'
	// Deprecated.
	KPEFTracebackSection KPEF = 8
	// Deprecated.
	KPEFUndefinedSymbol KPEF = 15
	// Deprecated.
	KPEFUnpackedDataSection KPEF = 1
	// Deprecated.
	KPEFVersion KPEF = 1
	// Deprecated.
	KPEFWeakImportLibMask KPEF = 64
	// Deprecated.
	KPEFWeakImportSymMask KPEF = 128
)

func (e KPEF) String() string {
	switch e {
	case KPEFAbsoluteExport:
		return "KPEFAbsoluteExport"
	case KPEFCodeSection:
		return "KPEFCodeSection"
	case KPEFConstantSection:
		return "KPEFConstantSection"
	case KPEFDataSymbol:
		return "KPEFDataSymbol"
	case KPEFDebugSection:
		return "KPEFDebugSection"
	case KPEFExceptionSection:
		return "KPEFExceptionSection"
	case KPEFExecDataSection:
		return "KPEFExecDataSection"
	case KPEFGlobalShare:
		return "KPEFGlobalShare"
	case KPEFInitLibBeforeMask:
		return "KPEFInitLibBeforeMask"
	case KPEFPackedDataSection:
		return "KPEFPackedDataSection"
	case KPEFReexportedImport:
		return "KPEFReexportedImport"
	case KPEFTag1:
		return "KPEFTag1"
	case KPEFTag2:
		return "KPEFTag2"
	case KPEFTracebackSection:
		return "KPEFTracebackSection"
	case KPEFUndefinedSymbol:
		return "KPEFUndefinedSymbol"
	case KPEFWeakImportLibMask:
		return "KPEFWeakImportLibMask"
	default:
		return fmt.Sprintf("KPEF(%d)", e)
	}
}

type KPEFExpSym uint

const (
	// Deprecated.
	KPEFExpSymClassShift KPEFExpSym = 24
	// Deprecated.
	KPEFExpSymMaxNameOffset KPEFExpSym = 16777215
	// Deprecated.
	KPEFExpSymNameOffsetMask KPEFExpSym = 16777215
)

func (e KPEFExpSym) String() string {
	switch e {
	case KPEFExpSymClassShift:
		return "KPEFExpSymClassShift"
	case KPEFExpSymMaxNameOffset:
		return "KPEFExpSymMaxNameOffset"
	default:
		return fmt.Sprintf("KPEFExpSym(%d)", e)
	}
}

type KPEFFirstSectionHeader uint

const (
	// Deprecated.
	KPEFFirstSectionHeaderOffset KPEFFirstSectionHeader = 0
)

func (e KPEFFirstSectionHeader) String() string {
	switch e {
	case KPEFFirstSectionHeaderOffset:
		return "KPEFFirstSectionHeaderOffset"
	default:
		return fmt.Sprintf("KPEFFirstSectionHeader(%d)", e)
	}
}

type KPEFHash uint

const (
	// Deprecated.
	KPEFHashLengthShift KPEFHash = 16
	// Deprecated.
	KPEFHashMaxLength KPEFHash = 65535
	// Deprecated.
	KPEFHashValueMask KPEFHash = 65535
)

func (e KPEFHash) String() string {
	switch e {
	case KPEFHashLengthShift:
		return "KPEFHashLengthShift"
	case KPEFHashMaxLength:
		return "KPEFHashMaxLength"
	default:
		return fmt.Sprintf("KPEFHash(%d)", e)
	}
}

type KPEFHashSlot uint

const (
	// Deprecated.
	KPEFHashSlotFirstKeyMask KPEFHashSlot = 262143
	// Deprecated.
	KPEFHashSlotMaxKeyIndex KPEFHashSlot = 262143
	// Deprecated.
	KPEFHashSlotMaxSymbolCount KPEFHashSlot = 16383
	// Deprecated.
	KPEFHashSlotSymCountShift KPEFHashSlot = 18
)

func (e KPEFHashSlot) String() string {
	switch e {
	case KPEFHashSlotFirstKeyMask:
		return "KPEFHashSlotFirstKeyMask"
	case KPEFHashSlotMaxSymbolCount:
		return "KPEFHashSlotMaxSymbolCount"
	case KPEFHashSlotSymCountShift:
		return "KPEFHashSlotSymCountShift"
	default:
		return fmt.Sprintf("KPEFHashSlot(%d)", e)
	}
}

type KPEFImpSym uint

const (
	// Deprecated.
	KPEFImpSymClassShift KPEFImpSym = 24
	// Deprecated.
	KPEFImpSymMaxNameOffset KPEFImpSym = 16777215
	// Deprecated.
	KPEFImpSymNameOffsetMask KPEFImpSym = 16777215
)

func (e KPEFImpSym) String() string {
	switch e {
	case KPEFImpSymClassShift:
		return "KPEFImpSymClassShift"
	case KPEFImpSymMaxNameOffset:
		return "KPEFImpSymMaxNameOffset"
	default:
		return fmt.Sprintf("KPEFImpSym(%d)", e)
	}
}

type KPEFPkData uint

const (
	// Deprecated.
	KPEFPkDataBlock KPEFPkData = 1
	// Deprecated.
	KPEFPkDataCount5Mask KPEFPkData = 31
	// Deprecated.
	KPEFPkDataMaxCount5 KPEFPkData = 31
	// Deprecated.
	KPEFPkDataOpcodeShift KPEFPkData = 5
	// Deprecated.
	KPEFPkDataRepeat KPEFPkData = 2
	// Deprecated.
	KPEFPkDataRepeatBlock KPEFPkData = 3
	// Deprecated.
	KPEFPkDataRepeatZero KPEFPkData = 4
	// Deprecated.
	KPEFPkDataVCountEndMask KPEFPkData = 128
	// Deprecated.
	KPEFPkDataVCountMask KPEFPkData = 127
	// Deprecated.
	KPEFPkDataVCountShift KPEFPkData = 7
	// Deprecated.
	KPEFPkDataZero KPEFPkData = 0
)

func (e KPEFPkData) String() string {
	switch e {
	case KPEFPkDataBlock:
		return "KPEFPkDataBlock"
	case KPEFPkDataCount5Mask:
		return "KPEFPkDataCount5Mask"
	case KPEFPkDataOpcodeShift:
		return "KPEFPkDataOpcodeShift"
	case KPEFPkDataRepeat:
		return "KPEFPkDataRepeat"
	case KPEFPkDataRepeatBlock:
		return "KPEFPkDataRepeatBlock"
	case KPEFPkDataRepeatZero:
		return "KPEFPkDataRepeatZero"
	case KPEFPkDataVCountEndMask:
		return "KPEFPkDataVCountEndMask"
	case KPEFPkDataVCountMask:
		return "KPEFPkDataVCountMask"
	case KPEFPkDataVCountShift:
		return "KPEFPkDataVCountShift"
	case KPEFPkDataZero:
		return "KPEFPkDataZero"
	default:
		return fmt.Sprintf("KPEFPkData(%d)", e)
	}
}

type KPEFReloc uint

const (
	// Deprecated.
	KPEFRelocBySectC KPEFReloc = 32
	// Deprecated.
	KPEFRelocBySectD KPEFReloc = 33
	// Deprecated.
	KPEFRelocBySectDWithSkip KPEFReloc = 0
	// Deprecated.
	KPEFRelocImportRun KPEFReloc = 37
	// Deprecated.
	KPEFRelocIncrPosition KPEFReloc = 64
	// Deprecated.
	KPEFRelocLgByImport KPEFReloc = 82
	// Deprecated.
	KPEFRelocLgRepeat KPEFReloc = 88
	// Deprecated.
	KPEFRelocLgSetOrBySection KPEFReloc = 90
	// Deprecated.
	KPEFRelocSetPosition KPEFReloc = 80
	// Deprecated.
	KPEFRelocSmByImport KPEFReloc = 48
	// Deprecated.
	KPEFRelocSmBySection KPEFReloc = 51
	// Deprecated.
	KPEFRelocSmRepeat KPEFReloc = 72
	// Deprecated.
	KPEFRelocSmSetSectC KPEFReloc = 49
	// Deprecated.
	KPEFRelocSmSetSectD KPEFReloc = 50
	// Deprecated.
	KPEFRelocTVector12 KPEFReloc = 34
	// Deprecated.
	KPEFRelocTVector8 KPEFReloc = 35
	// Deprecated.
	KPEFRelocUndefinedOpcode KPEFReloc = 255
	// Deprecated.
	KPEFRelocVTable8 KPEFReloc = 36
)

func (e KPEFReloc) String() string {
	switch e {
	case KPEFRelocBySectC:
		return "KPEFRelocBySectC"
	case KPEFRelocBySectD:
		return "KPEFRelocBySectD"
	case KPEFRelocBySectDWithSkip:
		return "KPEFRelocBySectDWithSkip"
	case KPEFRelocImportRun:
		return "KPEFRelocImportRun"
	case KPEFRelocIncrPosition:
		return "KPEFRelocIncrPosition"
	case KPEFRelocLgByImport:
		return "KPEFRelocLgByImport"
	case KPEFRelocLgRepeat:
		return "KPEFRelocLgRepeat"
	case KPEFRelocLgSetOrBySection:
		return "KPEFRelocLgSetOrBySection"
	case KPEFRelocSetPosition:
		return "KPEFRelocSetPosition"
	case KPEFRelocSmByImport:
		return "KPEFRelocSmByImport"
	case KPEFRelocSmBySection:
		return "KPEFRelocSmBySection"
	case KPEFRelocSmRepeat:
		return "KPEFRelocSmRepeat"
	case KPEFRelocSmSetSectC:
		return "KPEFRelocSmSetSectC"
	case KPEFRelocSmSetSectD:
		return "KPEFRelocSmSetSectD"
	case KPEFRelocTVector12:
		return "KPEFRelocTVector12"
	case KPEFRelocTVector8:
		return "KPEFRelocTVector8"
	case KPEFRelocUndefinedOpcode:
		return "KPEFRelocUndefinedOpcode"
	case KPEFRelocVTable8:
		return "KPEFRelocVTable8"
	default:
		return fmt.Sprintf("KPEFReloc(%d)", e)
	}
}

type KPEFRelocBasicOpcode uint

const (
	// Deprecated.
	KPEFRelocBasicOpcodeRange KPEFRelocBasicOpcode = 128
)

func (e KPEFRelocBasicOpcode) String() string {
	switch e {
	case KPEFRelocBasicOpcodeRange:
		return "KPEFRelocBasicOpcodeRange"
	default:
		return fmt.Sprintf("KPEFRelocBasicOpcode(%d)", e)
	}
}

type KPEFRelocIncrPositionMax uint

const (
	// Deprecated.
	KPEFRelocIncrPositionMaxOffset KPEFRelocIncrPositionMax = 4096
)

func (e KPEFRelocIncrPositionMax) String() string {
	switch e {
	case KPEFRelocIncrPositionMaxOffset:
		return "KPEFRelocIncrPositionMaxOffset"
	default:
		return fmt.Sprintf("KPEFRelocIncrPositionMax(%d)", e)
	}
}

type KPEFRelocLg uint

const (
	// Deprecated.
	KPEFRelocLgBySectionSubopcode KPEFRelocLg = 0
	// Deprecated.
	KPEFRelocLgSetSectCSubopcode KPEFRelocLg = 1
	// Deprecated.
	KPEFRelocLgSetSectDSubopcode KPEFRelocLg = 2
)

func (e KPEFRelocLg) String() string {
	switch e {
	case KPEFRelocLgBySectionSubopcode:
		return "KPEFRelocLgBySectionSubopcode"
	case KPEFRelocLgSetSectCSubopcode:
		return "KPEFRelocLgSetSectCSubopcode"
	case KPEFRelocLgSetSectDSubopcode:
		return "KPEFRelocLgSetSectDSubopcode"
	default:
		return fmt.Sprintf("KPEFRelocLg(%d)", e)
	}
}

type KPEFRelocLgByImportMax uint

const (
	// Deprecated.
	KPEFRelocLgByImportMaxIndex KPEFRelocLgByImportMax = 67108863
)

func (e KPEFRelocLgByImportMax) String() string {
	switch e {
	case KPEFRelocLgByImportMaxIndex:
		return "KPEFRelocLgByImportMaxIndex"
	default:
		return fmt.Sprintf("KPEFRelocLgByImportMax(%d)", e)
	}
}

type KPEFRelocLgRepeatMax uint

const (
	// Deprecated.
	KPEFRelocLgRepeatMaxChunkCount KPEFRelocLgRepeatMax = 16
	// Deprecated.
	KPEFRelocLgRepeatMaxRepeatCount KPEFRelocLgRepeatMax = 4194303
)

func (e KPEFRelocLgRepeatMax) String() string {
	switch e {
	case KPEFRelocLgRepeatMaxChunkCount:
		return "KPEFRelocLgRepeatMaxChunkCount"
	case KPEFRelocLgRepeatMaxRepeatCount:
		return "KPEFRelocLgRepeatMaxRepeatCount"
	default:
		return fmt.Sprintf("KPEFRelocLgRepeatMax(%d)", e)
	}
}

type KPEFRelocLgSetOrBySectionMax uint

const (
	// Deprecated.
	KPEFRelocLgSetOrBySectionMaxIndex KPEFRelocLgSetOrBySectionMax = 4194303
)

func (e KPEFRelocLgSetOrBySectionMax) String() string {
	switch e {
	case KPEFRelocLgSetOrBySectionMaxIndex:
		return "KPEFRelocLgSetOrBySectionMaxIndex"
	default:
		return fmt.Sprintf("KPEFRelocLgSetOrBySectionMax(%d)", e)
	}
}

type KPEFRelocRunMaxRun uint

const (
	// Deprecated.
	KPEFRelocRunMaxRunLength KPEFRelocRunMaxRun = 512
)

func (e KPEFRelocRunMaxRun) String() string {
	switch e {
	case KPEFRelocRunMaxRunLength:
		return "KPEFRelocRunMaxRunLength"
	default:
		return fmt.Sprintf("KPEFRelocRunMaxRun(%d)", e)
	}
}

type KPEFRelocSetPosMax uint

const (
	// Deprecated.
	KPEFRelocSetPosMaxOffset KPEFRelocSetPosMax = 67108863
)

func (e KPEFRelocSetPosMax) String() string {
	switch e {
	case KPEFRelocSetPosMaxOffset:
		return "KPEFRelocSetPosMaxOffset"
	default:
		return fmt.Sprintf("KPEFRelocSetPosMax(%d)", e)
	}
}

type KPEFRelocSmIndexMax uint

const (
	// Deprecated.
	KPEFRelocSmIndexMaxIndex KPEFRelocSmIndexMax = 511
)

func (e KPEFRelocSmIndexMax) String() string {
	switch e {
	case KPEFRelocSmIndexMaxIndex:
		return "KPEFRelocSmIndexMaxIndex"
	default:
		return fmt.Sprintf("KPEFRelocSmIndexMax(%d)", e)
	}
}

type KPEFRelocSmRepeatMax uint

const (
	// Deprecated.
	KPEFRelocSmRepeatMaxChunkCount KPEFRelocSmRepeatMax = 16
	// Deprecated.
	KPEFRelocSmRepeatMaxRepeatCount KPEFRelocSmRepeatMax = 256
)

func (e KPEFRelocSmRepeatMax) String() string {
	switch e {
	case KPEFRelocSmRepeatMaxChunkCount:
		return "KPEFRelocSmRepeatMaxChunkCount"
	case KPEFRelocSmRepeatMaxRepeatCount:
		return "KPEFRelocSmRepeatMaxRepeatCount"
	default:
		return fmt.Sprintf("KPEFRelocSmRepeatMax(%d)", e)
	}
}

type KPEFRelocWithSkipMax uint

const (
	// Deprecated.
	KPEFRelocWithSkipMaxRelocCount KPEFRelocWithSkipMax = 63
	// Deprecated.
	KPEFRelocWithSkipMaxSkipCount KPEFRelocWithSkipMax = 255
)

func (e KPEFRelocWithSkipMax) String() string {
	switch e {
	case KPEFRelocWithSkipMaxRelocCount:
		return "KPEFRelocWithSkipMaxRelocCount"
	case KPEFRelocWithSkipMaxSkipCount:
		return "KPEFRelocWithSkipMaxSkipCount"
	default:
		return fmt.Sprintf("KPEFRelocWithSkipMax(%d)", e)
	}
}

type KPOSIXError int

const (
	KPOSIXErrorBase            KPOSIXError = 100000
	KPOSIXErrorE2BIG           KPOSIXError = 100007
	KPOSIXErrorEACCES          KPOSIXError = 100013
	KPOSIXErrorEADDRINUSE      KPOSIXError = 100048
	KPOSIXErrorEADDRNOTAVAIL   KPOSIXError = 100049
	KPOSIXErrorEAFNOSUPPORT    KPOSIXError = 100047
	KPOSIXErrorEAGAIN          KPOSIXError = 100035
	KPOSIXErrorEALREADY        KPOSIXError = 100037
	KPOSIXErrorEAUTH           KPOSIXError = 100080
	KPOSIXErrorEBADARCH        KPOSIXError = 100086
	KPOSIXErrorEBADEXEC        KPOSIXError = 100085
	KPOSIXErrorEBADF           KPOSIXError = 100009
	KPOSIXErrorEBADMACHO       KPOSIXError = 100088
	KPOSIXErrorEBADMSG         KPOSIXError = 100094
	KPOSIXErrorEBADRPC         KPOSIXError = 100072
	KPOSIXErrorEBUSY           KPOSIXError = 100016
	KPOSIXErrorECANCELED       KPOSIXError = 100089
	KPOSIXErrorECHILD          KPOSIXError = 100010
	KPOSIXErrorECONNABORTED    KPOSIXError = 100053
	KPOSIXErrorECONNREFUSED    KPOSIXError = 100061
	KPOSIXErrorECONNRESET      KPOSIXError = 100054
	KPOSIXErrorEDEADLK         KPOSIXError = 100011
	KPOSIXErrorEDESTADDRREQ    KPOSIXError = 100039
	KPOSIXErrorEDEVERR         KPOSIXError = 100083
	KPOSIXErrorEDOM            KPOSIXError = 100033
	KPOSIXErrorEDQUOT          KPOSIXError = 100069
	KPOSIXErrorEEXIST          KPOSIXError = 100017
	KPOSIXErrorEFAULT          KPOSIXError = 100014
	KPOSIXErrorEFBIG           KPOSIXError = 100027
	KPOSIXErrorEFTYPE          KPOSIXError = 100079
	KPOSIXErrorEHOSTDOWN       KPOSIXError = 100064
	KPOSIXErrorEHOSTUNREACH    KPOSIXError = 100065
	KPOSIXErrorEIDRM           KPOSIXError = 100090
	KPOSIXErrorEILSEQ          KPOSIXError = 100092
	KPOSIXErrorEINPROGRESS     KPOSIXError = 100036
	KPOSIXErrorEINTR           KPOSIXError = 100004
	KPOSIXErrorEINVAL          KPOSIXError = 100022
	KPOSIXErrorEIO             KPOSIXError = 100005
	KPOSIXErrorEISCONN         KPOSIXError = 100056
	KPOSIXErrorEISDIR          KPOSIXError = 100021
	KPOSIXErrorELOOP           KPOSIXError = 100062
	KPOSIXErrorEMFILE          KPOSIXError = 100024
	KPOSIXErrorEMLINK          KPOSIXError = 100031
	KPOSIXErrorEMSGSIZE        KPOSIXError = 100040
	KPOSIXErrorEMULTIHOP       KPOSIXError = 100095
	KPOSIXErrorENAMETOOLONG    KPOSIXError = 100063
	KPOSIXErrorENEEDAUTH       KPOSIXError = 100081
	KPOSIXErrorENETDOWN        KPOSIXError = 100050
	KPOSIXErrorENETRESET       KPOSIXError = 100052
	KPOSIXErrorENETUNREACH     KPOSIXError = 100051
	KPOSIXErrorENFILE          KPOSIXError = 100023
	KPOSIXErrorENOATTR         KPOSIXError = 100093
	KPOSIXErrorENOBUFS         KPOSIXError = 100055
	KPOSIXErrorENODATA         KPOSIXError = 100096
	KPOSIXErrorENODEV          KPOSIXError = 100019
	KPOSIXErrorENOENT          KPOSIXError = 100002
	KPOSIXErrorENOEXEC         KPOSIXError = 100008
	KPOSIXErrorENOLCK          KPOSIXError = 100077
	KPOSIXErrorENOLINK         KPOSIXError = 100097
	KPOSIXErrorENOMEM          KPOSIXError = 100012
	KPOSIXErrorENOMSG          KPOSIXError = 100091
	KPOSIXErrorENOPROTOOPT     KPOSIXError = 100042
	KPOSIXErrorENOSPC          KPOSIXError = 100028
	KPOSIXErrorENOSR           KPOSIXError = 100098
	KPOSIXErrorENOSTR          KPOSIXError = 100099
	KPOSIXErrorENOSYS          KPOSIXError = 100078
	KPOSIXErrorENOTBLK         KPOSIXError = 100015
	KPOSIXErrorENOTCONN        KPOSIXError = 100057
	KPOSIXErrorENOTDIR         KPOSIXError = 100020
	KPOSIXErrorENOTEMPTY       KPOSIXError = 100066
	KPOSIXErrorENOTSOCK        KPOSIXError = 100038
	KPOSIXErrorENOTSUP         KPOSIXError = 100045
	KPOSIXErrorENOTTY          KPOSIXError = 100025
	KPOSIXErrorENXIO           KPOSIXError = 100006
	KPOSIXErrorEOPNOTSUPP      KPOSIXError = 100102
	KPOSIXErrorEOVERFLOW       KPOSIXError = 100084
	KPOSIXErrorEPERM           KPOSIXError = 100001
	KPOSIXErrorEPFNOSUPPORT    KPOSIXError = 100046
	KPOSIXErrorEPIPE           KPOSIXError = 100032
	KPOSIXErrorEPROCLIM        KPOSIXError = 100067
	KPOSIXErrorEPROCUNAVAIL    KPOSIXError = 100076
	KPOSIXErrorEPROGMISMATCH   KPOSIXError = 100075
	KPOSIXErrorEPROGUNAVAIL    KPOSIXError = 100074
	KPOSIXErrorEPROTO          KPOSIXError = 100100
	KPOSIXErrorEPROTONOSUPPORT KPOSIXError = 100043
	KPOSIXErrorEPROTOTYPE      KPOSIXError = 100041
	KPOSIXErrorEPWROFF         KPOSIXError = 100082
	KPOSIXErrorERANGE          KPOSIXError = 100034
	KPOSIXErrorEREMOTE         KPOSIXError = 100071
	KPOSIXErrorEROFS           KPOSIXError = 100030
	KPOSIXErrorERPCMISMATCH    KPOSIXError = 100073
	KPOSIXErrorESHLIBVERS      KPOSIXError = 100087
	KPOSIXErrorESHUTDOWN       KPOSIXError = 100058
	KPOSIXErrorESOCKTNOSUPPORT KPOSIXError = 100044
	KPOSIXErrorESPIPE          KPOSIXError = 100029
	KPOSIXErrorESRCH           KPOSIXError = 100003
	KPOSIXErrorESTALE          KPOSIXError = 100070
	KPOSIXErrorETIME           KPOSIXError = 100101
	KPOSIXErrorETIMEDOUT       KPOSIXError = 100060
	KPOSIXErrorETOOMANYREFS    KPOSIXError = 100059
	KPOSIXErrorETXTBSY         KPOSIXError = 100026
	KPOSIXErrorEUSERS          KPOSIXError = 100068
	KPOSIXErrorEXDEV           KPOSIXError = 100018
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

type KPageInMemory uint

const (
	// Deprecated.
	KNotPaged KPageInMemory = 2
	// Deprecated.
	KPageInMemoryValue KPageInMemory = 0
	// Deprecated.
	KPageOnDisk KPageInMemory = 1
)

func (e KPageInMemory) String() string {
	switch e {
	case KNotPaged:
		return "KNotPaged"
	case KPageInMemoryValue:
		return "KPageInMemoryValue"
	case KPageOnDisk:
		return "KPageOnDisk"
	default:
		return fmt.Sprintf("KPageInMemory(%d)", e)
	}
}

type KPascalStackBased uint

const (
	// Deprecated.
	KCStackBased KPascalStackBased = 1
	// Deprecated.
	KD0DispatchedCStackBased KPascalStackBased = 9
	// Deprecated.
	KD0DispatchedPascalStackBased KPascalStackBased = 8
	// Deprecated.
	KD1DispatchedPascalStackBased KPascalStackBased = 12
	// Deprecated.
	KPascalStackBasedValue KPascalStackBased = 0
	// Deprecated.
	KRegisterBased KPascalStackBased = 2
	// Deprecated.
	KStackDispatchedPascalStackBased KPascalStackBased = 14
	// Deprecated.
	KThinkCStackBased KPascalStackBased = 5
)

func (e KPascalStackBased) String() string {
	switch e {
	case KCStackBased:
		return "KCStackBased"
	case KD0DispatchedCStackBased:
		return "KD0DispatchedCStackBased"
	case KD0DispatchedPascalStackBased:
		return "KD0DispatchedPascalStackBased"
	case KD1DispatchedPascalStackBased:
		return "KD1DispatchedPascalStackBased"
	case KPascalStackBasedValue:
		return "KPascalStackBasedValue"
	case KRegisterBased:
		return "KRegisterBased"
	case KStackDispatchedPascalStackBased:
		return "KStackDispatchedPascalStackBased"
	case KThinkCStackBased:
		return "KThinkCStackBased"
	default:
		return fmt.Sprintf("KPascalStackBased(%d)", e)
	}
}

type KPassSelector uint

const (
	// Deprecated.
	KDontPassSelector KPassSelector = 8
	// Deprecated.
	KPassSelectorValue KPassSelector = 0
)

func (e KPassSelector) String() string {
	switch e {
	case KDontPassSelector:
		return "KDontPassSelector"
	case KPassSelectorValue:
		return "KPassSelectorValue"
	default:
		return fmt.Sprintf("KPassSelector(%d)", e)
	}
}

type KPolicyKCStopOn uint

const (
	KFirstFailKCStopOn   KPolicyKCStopOn = 3
	KFirstPassKCStopOn   KPolicyKCStopOn = 2
	KNoneKCStopOn        KPolicyKCStopOn = 1
	KPolicyKCStopOnValue KPolicyKCStopOn = 0
)

func (e KPolicyKCStopOn) String() string {
	switch e {
	case KFirstFailKCStopOn:
		return "KFirstFailKCStopOn"
	case KFirstPassKCStopOn:
		return "KFirstPassKCStopOn"
	case KNoneKCStopOn:
		return "KNoneKCStopOn"
	case KPolicyKCStopOnValue:
		return "KPolicyKCStopOnValue"
	default:
		return fmt.Sprintf("KPolicyKCStopOn(%d)", e)
	}
}

type KPowerHandlerExistsForDeviceErr int

const (
	KBridgeSoftwareRunningCantSleep      KPowerHandlerExistsForDeviceErr = -13038
	KCantReportProcessorTemperatureErr   KPowerHandlerExistsForDeviceErr = -13013
	KNoSuchPowerSource                   KPowerHandlerExistsForDeviceErr = -13020
	KPowerHandlerExistsForDeviceErrValue KPowerHandlerExistsForDeviceErr = -13006
	KPowerHandlerNotFoundForDeviceErr    KPowerHandlerExistsForDeviceErr = -13007
	KPowerHandlerNotFoundForProcErr      KPowerHandlerExistsForDeviceErr = -13008
	KPowerMgtMessageNotHandled           KPowerHandlerExistsForDeviceErr = -13009
	KPowerMgtRequestDenied               KPowerHandlerExistsForDeviceErr = -13010
	KProcessorTempRoutineRequiresMPLib2  KPowerHandlerExistsForDeviceErr = -13014
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

type KProcDescriptorIs uint

const (
	// Deprecated.
	KProcDescriptorIsAbsolute KProcDescriptorIs = 0
	// Deprecated.
	KProcDescriptorIsIndex KProcDescriptorIs = 32
	// Deprecated.
	KProcDescriptorIsProcPtr KProcDescriptorIs = 0
	// Deprecated.
	KProcDescriptorIsRelative KProcDescriptorIs = 1
)

func (e KProcDescriptorIs) String() string {
	switch e {
	case KProcDescriptorIsAbsolute:
		return "KProcDescriptorIsAbsolute"
	case KProcDescriptorIsIndex:
		return "KProcDescriptorIsIndex"
	case KProcDescriptorIsRelative:
		return "KProcDescriptorIsRelative"
	default:
		return fmt.Sprintf("KProcDescriptorIs(%d)", e)
	}
}

type KQD int

const (
	KQDCorruptPICTDataErr      KQD = -3954
	KQDCursorAlreadyRegistered KQD = -3952
	KQDCursorNotRegistered     KQD = -3953
	KQDNoColorHWCursorSupport  KQD = -3951
	KQDNoPalette               KQD = -3950
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

type KQTSSUnknown int

const (
	KQTSSUnknownErr KQTSSUnknown = -6150
)

func (e KQTSSUnknown) String() string {
	switch e {
	case KQTSSUnknownErr:
		return "KQTSSUnknownErr"
	default:
		return fmt.Sprintf("KQTSSUnknown(%d)", e)
	}
}

type KQuickLookFolder uint

const (
	// Deprecated.
	KQuickLookFolderType KQuickLookFolder = 'q'<<24 | 'l'<<16 | 'c'<<8 | 'k' // 'qlck'
)

func (e KQuickLookFolder) String() string {
	switch e {
	case KQuickLookFolderType:
		return "KQuickLookFolderType"
	default:
		return fmt.Sprintf("KQuickLookFolder(%d)", e)
	}
}

type KRA int

const (
	KRAATalkInactive           KRA = -7134
	KRACallBackFailed          KRA = -7138
	KRAConfigurationDBInitErr  KRA = -7127
	KRAConnectionCanceled      KRA = -7109
	KRADuplicateIPAddr         KRA = -7137
	KRAExtAuthenticationFailed KRA = -7135
	KRAIncompatiblePrefs       KRA = -7107
	KRAInitOpenTransportFailed KRA = -7122
	KRAInstallationDamaged     KRA = -7113
	KRAInternalError           KRA = -7112
	KRAInvalidParameter        KRA = -7100
	KRAInvalidPassword         KRA = -7111
	KRAInvalidPort             KRA = -7101
	KRAInvalidPortState        KRA = -7116
	KRAInvalidSerialProtocol   KRA = -7117
	KRAMissingResources        KRA = -7106
	KRANCPRejectedbyPeer       KRA = -7136
	KRANotConnected            KRA = -7108
	KRANotEnabled              KRA = -7139
	KRANotPrimaryInterface     KRA = -7126
	KRANotSupported            KRA = -7105
	KRAOutOfMemory             KRA = -7104
	KRAPPPAuthenticationFailed KRA = -7129
	KRAPPPNegotiationFailed    KRA = -7130
	KRAPPPPeerDisconnected     KRA = -7132
	KRAPPPProtocolRejected     KRA = -7128
	KRAPPPUserDisconnected     KRA = -7131
	KRAPeerNotResponding       KRA = -7133
	KRAPortBusy                KRA = -7114
	KRAPortSetupFailed         KRA = -7103
	KRARemoteAccessNotReady    KRA = -7123
	KRAStartupFailed           KRA = -7102
	KRATCPIPInactive           KRA = -7124
	KRATCPIPNotConfigured      KRA = -7125
	KRAUnknownPortState        KRA = -7115
	KRAUnknownUser             KRA = -7110
	KRAUserInteractionRequired KRA = -7121
	KRAUserLoginDisabled       KRA = -7118
	KRAUserPwdChangeRequired   KRA = -7119
	KRAUserPwdEntryRequired    KRA = -7120
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

type KReadExtensionTerms uint

const (
	KReadExtensionTermsMask KReadExtensionTerms = 32768
)

func (e KReadExtensionTerms) String() string {
	switch e {
	case KReadExtensionTermsMask:
		return "KReadExtensionTermsMask"
	default:
		return fmt.Sprintf("KReadExtensionTerms(%d)", e)
	}
}

type KReadyThreadState uint

const (
	// Deprecated.
	KReadyThreadStateValue KReadyThreadState = 0
	// Deprecated.
	KRunningThreadState KReadyThreadState = 2
	// Deprecated.
	KStoppedThreadState KReadyThreadState = 1
)

func (e KReadyThreadState) String() string {
	switch e {
	case KReadyThreadStateValue:
		return "KReadyThreadStateValue"
	case KRunningThreadState:
		return "KRunningThreadState"
	case KStoppedThreadState:
		return "KStoppedThreadState"
	default:
		return fmt.Sprintf("KReadyThreadState(%d)", e)
	}
}

type KRegisterD0 uint

const (
	// Deprecated.
	KCCRegisterCBit KRegisterD0 = 16
	// Deprecated.
	KCCRegisterNBit KRegisterD0 = 19
	// Deprecated.
	KCCRegisterVBit KRegisterD0 = 17
	// Deprecated.
	KCCRegisterXBit KRegisterD0 = 20
	// Deprecated.
	KCCRegisterZBit KRegisterD0 = 18
	// Deprecated.
	KRegisterA0 KRegisterD0 = 4
	// Deprecated.
	KRegisterA1 KRegisterD0 = 5
	// Deprecated.
	KRegisterA2 KRegisterD0 = 6
	// Deprecated.
	KRegisterA3 KRegisterD0 = 7
	// Deprecated.
	KRegisterA4 KRegisterD0 = 12
	// Deprecated.
	KRegisterA5 KRegisterD0 = 13
	// Deprecated.
	KRegisterA6 KRegisterD0 = 14
	// Deprecated.
	KRegisterD0Value KRegisterD0 = 0
	// Deprecated.
	KRegisterD1 KRegisterD0 = 1
	// Deprecated.
	KRegisterD2 KRegisterD0 = 2
	// Deprecated.
	KRegisterD3 KRegisterD0 = 3
	// Deprecated.
	KRegisterD4 KRegisterD0 = 8
	// Deprecated.
	KRegisterD5 KRegisterD0 = 9
	// Deprecated.
	KRegisterD6 KRegisterD0 = 10
	// Deprecated.
	KRegisterD7 KRegisterD0 = 11
)

func (e KRegisterD0) String() string {
	switch e {
	case KCCRegisterCBit:
		return "KCCRegisterCBit"
	case KCCRegisterNBit:
		return "KCCRegisterNBit"
	case KCCRegisterVBit:
		return "KCCRegisterVBit"
	case KCCRegisterXBit:
		return "KCCRegisterXBit"
	case KCCRegisterZBit:
		return "KCCRegisterZBit"
	case KRegisterA0:
		return "KRegisterA0"
	case KRegisterA1:
		return "KRegisterA1"
	case KRegisterA2:
		return "KRegisterA2"
	case KRegisterA3:
		return "KRegisterA3"
	case KRegisterA4:
		return "KRegisterA4"
	case KRegisterA5:
		return "KRegisterA5"
	case KRegisterA6:
		return "KRegisterA6"
	case KRegisterD0Value:
		return "KRegisterD0Value"
	case KRegisterD1:
		return "KRegisterD1"
	case KRegisterD2:
		return "KRegisterD2"
	case KRegisterD3:
		return "KRegisterD3"
	case KRegisterD4:
		return "KRegisterD4"
	case KRegisterD5:
		return "KRegisterD5"
	case KRegisterD6:
		return "KRegisterD6"
	case KRegisterD7:
		return "KRegisterD7"
	default:
		return fmt.Sprintf("KRegisterD0(%d)", e)
	}
}

type KRelativeFolder uint

const (
	KRedirectedRelativeFolder KRelativeFolder = 'r'<<24 | 'r'<<16 | 'e'<<8 | 'l' // 'rrel'
	KRelativeFolderValue      KRelativeFolder = 'r'<<24 | 'e'<<16 | 'l'<<8 | 'f' // 'relf'
	KSpecialFolder            KRelativeFolder = 's'<<24 | 'p'<<16 | 'c'<<8 | 'f' // 'spcf'
)

func (e KRelativeFolder) String() string {
	switch e {
	case KRedirectedRelativeFolder:
		return "KRedirectedRelativeFolder"
	case KRelativeFolderValue:
		return "KRelativeFolderValue"
	case KSpecialFolder:
		return "KSpecialFolder"
	default:
		return fmt.Sprintf("KRelativeFolder(%d)", e)
	}
}

type KResFileNotOpened int

const (
	// Deprecated.
	KResFileNotOpenedValue KResFileNotOpened = -1
	// Deprecated.
	KSystemResFile KResFileNotOpened = 0
)

func (e KResFileNotOpened) String() string {
	switch e {
	case KResFileNotOpenedValue:
		return "KResFileNotOpenedValue"
	case KSystemResFile:
		return "KSystemResFile"
	default:
		return fmt.Sprintf("KResFileNotOpened(%d)", e)
	}
}

type KResolveAlias uint

const (
	// Deprecated.
	KResolveAliasFileNoUI KResolveAlias = 1
	// Deprecated.
	KResolveAliasTryFileIDFirst KResolveAlias = 2
)

func (e KResolveAlias) String() string {
	switch e {
	case KResolveAliasFileNoUI:
		return "KResolveAliasFileNoUI"
	case KResolveAliasTryFileIDFirst:
		return "KResolveAliasTryFileIDFirst"
	default:
		return fmt.Sprintf("KResolveAlias(%d)", e)
	}
}

type KRoutineDescriptor uint

const (
	// Deprecated.
	KRoutineDescriptorVersion KRoutineDescriptor = 7
)

func (e KRoutineDescriptor) String() string {
	switch e {
	case KRoutineDescriptorVersion:
		return "KRoutineDescriptorVersion"
	default:
		return fmt.Sprintf("KRoutineDescriptor(%d)", e)
	}
}

type KRoutineIs uint

const (
	// Deprecated.
	KRoutineIsDispatchedDefaultRoutine KRoutineIs = 16
	// Deprecated.
	KRoutineIsNotDispatchedDefaultRoutine KRoutineIs = 0
)

func (e KRoutineIs) String() string {
	switch e {
	case KRoutineIsDispatchedDefaultRoutine:
		return "KRoutineIsDispatchedDefaultRoutine"
	case KRoutineIsNotDispatchedDefaultRoutine:
		return "KRoutineIsNotDispatchedDefaultRoutine"
	default:
		return fmt.Sprintf("KRoutineIs(%d)", e)
	}
}

type KRoutingResource uint

const (
	KRoutingResourceID   KRoutingResource = 0
	KRoutingResourceType KRoutingResource = 'r'<<24 | 'o'<<16 | 'u'<<8 | 't' // 'rout'
)

func (e KRoutingResource) String() string {
	switch e {
	case KRoutingResourceID:
		return "KRoutingResourceID"
	case KRoutingResourceType:
		return "KRoutingResourceType"
	default:
		return fmt.Sprintf("KRoutingResource(%d)", e)
	}
}

type KRsrcChain uint

const (
	// Deprecated.
	KRsrcChainAboveAllMaps KRsrcChain = 4
	// Deprecated.
	KRsrcChainAboveApplicationMap KRsrcChain = 2
	// Deprecated.
	KRsrcChainBelowApplicationMap KRsrcChain = 1
	// Deprecated.
	KRsrcChainBelowSystemMap KRsrcChain = 0
)

func (e KRsrcChain) String() string {
	switch e {
	case KRsrcChainAboveAllMaps:
		return "KRsrcChainAboveAllMaps"
	case KRsrcChainAboveApplicationMap:
		return "KRsrcChainAboveApplicationMap"
	case KRsrcChainBelowApplicationMap:
		return "KRsrcChainBelowApplicationMap"
	case KRsrcChainBelowSystemMap:
		return "KRsrcChainBelowSystemMap"
	default:
		return fmt.Sprintf("KRsrcChain(%d)", e)
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

type KSOA uint

const (
	KSOAP1999Schema KSOA = 's'<<24 | 's'<<16 | '9'<<8 | '9' // 'ss99'
	KSOAP2001Schema KSOA = 's'<<24 | 's'<<16 | '0'<<8 | '1' // 'ss01'
)

func (e KSOA) String() string {
	switch e {
	case KSOAP1999Schema:
		return "KSOAP1999Schema"
	case KSOAP2001Schema:
		return "KSOAP2001Schema"
	default:
		return fmt.Sprintf("KSOA(%d)", e)
	}
}

type KSSp int

const (
	KSSpCantInstallErr      KSSp = -30342
	KSSpInternalErr         KSSp = -30340
	KSSpParallelUpVectorErr KSSp = -30343
	KSSpScaleToZeroErr      KSSp = -30344
	KSSpVersionErr          KSSp = -30341
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

type KSecOptionReserved uint

const (
	KCertUsageAllAdd           KSecOptionReserved = 2147483392
	KCertUsageDecryptAdd       KSecOptionReserved = 1
	KCertUsageDecryptAskAndAdd KSecOptionReserved = 1
	KCertUsageEncryptAdd       KSecOptionReserved = 1
	KCertUsageEncryptAskAndAdd KSecOptionReserved = 1
	KCertUsageKeyExchAdd       KSecOptionReserved = 1
	KCertUsageKeyExchAskAndAdd KSecOptionReserved = 1
	KCertUsageRootAdd          KSecOptionReserved = 1
	KCertUsageRootAskAndAdd    KSecOptionReserved = 1
	KCertUsageSSLAdd           KSecOptionReserved = 1
	KCertUsageSSLAskAndAdd     KSecOptionReserved = 1
	KCertUsageShift            KSecOptionReserved = 8
	KCertUsageSigningAdd       KSecOptionReserved = 1
	KCertUsageSigningAskAndAdd KSecOptionReserved = 1
	KCertUsageVerifyAdd        KSecOptionReserved = 1
	KCertUsageVerifyAskAndAdd  KSecOptionReserved = 1
	KSecOptionReservedValue    KSecOptionReserved = 255
)

func (e KSecOptionReserved) String() string {
	switch e {
	case KCertUsageAllAdd:
		return "KCertUsageAllAdd"
	case KCertUsageDecryptAdd:
		return "KCertUsageDecryptAdd"
	case KCertUsageShift:
		return "KCertUsageShift"
	case KSecOptionReservedValue:
		return "KSecOptionReservedValue"
	default:
		return fmt.Sprintf("KSecOptionReserved(%d)", e)
	}
}

type KSelectorsAre uint

const (
	// Deprecated.
	KSelectorsAreIndexable KSelectorsAre = 1
	// Deprecated.
	KSelectorsAreNotIndexable KSelectorsAre = 0
)

func (e KSelectorsAre) String() string {
	switch e {
	case KSelectorsAreIndexable:
		return "KSelectorsAreIndexable"
	case KSelectorsAreNotIndexable:
		return "KSelectorsAreNotIndexable"
	default:
		return fmt.Sprintf("KSelectorsAre(%d)", e)
	}
}

type KServicesFolder uint

const (
	// Deprecated.
	KServicesFolderType KServicesFolder = 's'<<24 | 'v'<<16 | 'c'<<8 | 's' // 'svcs'
)

func (e KServicesFolder) String() string {
	switch e {
	case KServicesFolderType:
		return "KServicesFolderType"
	default:
		return fmt.Sprintf("KServicesFolder(%d)", e)
	}
}

type KSharedLibrariesFolderType uint

const (
	// Deprecated.
	KFavoritesFolderType KSharedLibrariesFolderType = 'f'<<24 | 'a'<<16 | 'v'<<8 | 's' // 'favs'
	// Deprecated.
	KFolderActionsFolderType KSharedLibrariesFolderType = 'f'<<24 | 'a'<<16 | 's'<<8 | 'f' // 'fasf'
	// Deprecated.
	KInstallerLogsFolderType KSharedLibrariesFolderType = 'i'<<24 | 'l'<<16 | 'g'<<8 | 'f' // 'ilgf'
	// Deprecated.
	KInternetSearchSitesFolderType KSharedLibrariesFolderType = 'i'<<24 | 's'<<16 | 's'<<8 | 'f' // 'issf'
	// Deprecated.
	KKeychainFolderType KSharedLibrariesFolderType = 'k'<<24 | 'c'<<16 | 'h'<<8 | 'n' // 'kchn'
	// Deprecated.
	KScriptsFolderType KSharedLibrariesFolderType = 1935897284
	// Deprecated.
	KSharedLibrariesFolderTypeValue KSharedLibrariesFolderType = 3295439202
	// Deprecated.
	KSpeakableItemsFolderType KSharedLibrariesFolderType = 's'<<24 | 'p'<<16 | 'k'<<8 | 'i' // 'spki'
	// Deprecated.
	KThemesFolderType KSharedLibrariesFolderType = 't'<<24 | 'h'<<16 | 'm'<<8 | 'e' // 'thme'
	// Deprecated.
	KUtilitiesFolderType KSharedLibrariesFolderType = 1970563524
	// Deprecated.
	KVoicesFolderType KSharedLibrariesFolderType = 'f'<<24 | 'v'<<16 | 'o'<<8 | 'c' // 'fvoc'
)

func (e KSharedLibrariesFolderType) String() string {
	switch e {
	case KFavoritesFolderType:
		return "KFavoritesFolderType"
	case KFolderActionsFolderType:
		return "KFolderActionsFolderType"
	case KInstallerLogsFolderType:
		return "KInstallerLogsFolderType"
	case KInternetSearchSitesFolderType:
		return "KInternetSearchSitesFolderType"
	case KKeychainFolderType:
		return "KKeychainFolderType"
	case KScriptsFolderType:
		return "KScriptsFolderType"
	case KSharedLibrariesFolderTypeValue:
		return "KSharedLibrariesFolderTypeValue"
	case KSpeakableItemsFolderType:
		return "KSpeakableItemsFolderType"
	case KThemesFolderType:
		return "KThemesFolderType"
	case KUtilitiesFolderType:
		return "KUtilitiesFolderType"
	case KVoicesFolderType:
		return "KVoicesFolderType"
	default:
		return fmt.Sprintf("KSharedLibrariesFolderType(%d)", e)
	}
}

type KSharingPrivs uint

const (
	KSharingPrivsNotApplicableIcon KSharingPrivs = 's'<<24 | 'h'<<16 | 'n'<<8 | 'a' // 'shna'
	KSharingPrivsReadOnlyIcon      KSharingPrivs = 's'<<24 | 'h'<<16 | 'r'<<8 | 'o' // 'shro'
	KSharingPrivsReadWriteIcon     KSharingPrivs = 's'<<24 | 'h'<<16 | 'r'<<8 | 'w' // 'shrw'
	KSharingPrivsUnknownIcon       KSharingPrivs = 's'<<24 | 'h'<<16 | 'u'<<8 | 'k' // 'shuk'
	KSharingPrivsWritableIcon      KSharingPrivs = 'w'<<24 | 'r'<<16 | 'i'<<8 | 't' // 'writ'
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

type KShiftJIS uint

const (
	KShiftJIS_BasicVariant   KShiftJIS = 0
	KShiftJIS_DOSVariant     KShiftJIS = 1
	KShiftJIS_MusicCDVariant KShiftJIS = 2
)

func (e KShiftJIS) String() string {
	switch e {
	case KShiftJIS_BasicVariant:
		return "KShiftJIS_BasicVariant"
	case KShiftJIS_DOSVariant:
		return "KShiftJIS_DOSVariant"
	case KShiftJIS_MusicCDVariant:
		return "KShiftJIS_MusicCDVariant"
	default:
		return fmt.Sprintf("KShiftJIS(%d)", e)
	}
}

type KSleepRequest uint

const (
	KDeviceInitiatedWake KSleepRequest = 18
	KDozeDemand          KSleepRequest = 7
	KDozeRequest         KSleepRequest = 9
	KDozeToFullWakeUp    KSleepRequest = 20
	KDozeWakeUp          KSleepRequest = 8
	KEnterIdle           KSleepRequest = 24
	KEnterRun            KSleepRequest = 11
	KEnterStandby        KSleepRequest = 10
	KExitIdle            KSleepRequest = 26
	KGetPowerInfo        KSleepRequest = 21
	KGetPowerLevel       KSleepRequest = 16
	KGetWakeOnNetInfo    KSleepRequest = 22
	KSetPowerLevel       KSleepRequest = 17
	KSleepDemand         KSleepRequest = 2
	KSleepDeny           KSleepRequest = 5
	KSleepNow            KSleepRequest = 6
	KSleepRequestValue   KSleepRequest = 1
	KSleepRevoke         KSleepRequest = 4
	KSleepUnlock         KSleepRequest = 4
	KSleepWakeUp         KSleepRequest = 3
	KStillIdle           KSleepRequest = 25
	KSuspendDemand       KSleepRequest = 13
	KSuspendRequest      KSleepRequest = 12
	KSuspendRevoke       KSleepRequest = 14
	KSuspendWakeToDoze   KSleepRequest = 23
	KSuspendWakeUp       KSleepRequest = 15
	KWakeToDoze          KSleepRequest = 19
)

func (e KSleepRequest) String() string {
	switch e {
	case KDeviceInitiatedWake:
		return "KDeviceInitiatedWake"
	case KDozeDemand:
		return "KDozeDemand"
	case KDozeRequest:
		return "KDozeRequest"
	case KDozeToFullWakeUp:
		return "KDozeToFullWakeUp"
	case KDozeWakeUp:
		return "KDozeWakeUp"
	case KEnterIdle:
		return "KEnterIdle"
	case KEnterRun:
		return "KEnterRun"
	case KEnterStandby:
		return "KEnterStandby"
	case KExitIdle:
		return "KExitIdle"
	case KGetPowerInfo:
		return "KGetPowerInfo"
	case KGetPowerLevel:
		return "KGetPowerLevel"
	case KGetWakeOnNetInfo:
		return "KGetWakeOnNetInfo"
	case KSetPowerLevel:
		return "KSetPowerLevel"
	case KSleepDemand:
		return "KSleepDemand"
	case KSleepDeny:
		return "KSleepDeny"
	case KSleepNow:
		return "KSleepNow"
	case KSleepRequestValue:
		return "KSleepRequestValue"
	case KSleepRevoke:
		return "KSleepRevoke"
	case KSleepWakeUp:
		return "KSleepWakeUp"
	case KStillIdle:
		return "KStillIdle"
	case KSuspendDemand:
		return "KSuspendDemand"
	case KSuspendRequest:
		return "KSuspendRequest"
	case KSuspendRevoke:
		return "KSuspendRevoke"
	case KSuspendWakeToDoze:
		return "KSuspendWakeToDoze"
	case KSuspendWakeUp:
		return "KSuspendWakeUp"
	case KWakeToDoze:
		return "KWakeToDoze"
	default:
		return fmt.Sprintf("KSleepRequest(%d)", e)
	}
}

type KSpecial uint

const (
	// Deprecated.
	KSpecialCaseValue KSpecial = 15
)

func (e KSpecial) String() string {
	switch e {
	case KSpecialCaseValue:
		return "KSpecialCaseValue"
	default:
		return fmt.Sprintf("KSpecial(%d)", e)
	}
}

type KSpecialCase uint

const (
	// Deprecated.
	KSpecialCaseCaretHook KSpecialCase = 0
	// Deprecated.
	KSpecialCaseDrawHook KSpecialCase = 4
	// Deprecated.
	KSpecialCaseEOLHook KSpecialCase = 1
	// Deprecated.
	KSpecialCaseGNEFilterProc KSpecialCase = 11
	// Deprecated.
	KSpecialCaseHighHook KSpecialCase = 0
	// Deprecated.
	KSpecialCaseHitTestHook KSpecialCase = 5
	// Deprecated.
	KSpecialCaseMBarHook KSpecialCase = 12
	// Deprecated.
	KSpecialCaseNWidthHook KSpecialCase = 3
	// Deprecated.
	KSpecialCaseProtocolHandler KSpecialCase = 7
	// Deprecated.
	KSpecialCaseSocketListener KSpecialCase = 8
	// Deprecated.
	KSpecialCaseTEDoText KSpecialCase = 10
	// Deprecated.
	KSpecialCaseTEFindWord KSpecialCase = 6
	// Deprecated.
	KSpecialCaseTERecalc KSpecialCase = 9
	// Deprecated.
	KSpecialCaseTextWidthHook KSpecialCase = 2
	// Deprecated.
	KSpecialCaseWidthHook KSpecialCase = 2
)

func (e KSpecialCase) String() string {
	switch e {
	case KSpecialCaseCaretHook:
		return "KSpecialCaseCaretHook"
	case KSpecialCaseDrawHook:
		return "KSpecialCaseDrawHook"
	case KSpecialCaseEOLHook:
		return "KSpecialCaseEOLHook"
	case KSpecialCaseGNEFilterProc:
		return "KSpecialCaseGNEFilterProc"
	case KSpecialCaseHitTestHook:
		return "KSpecialCaseHitTestHook"
	case KSpecialCaseMBarHook:
		return "KSpecialCaseMBarHook"
	case KSpecialCaseNWidthHook:
		return "KSpecialCaseNWidthHook"
	case KSpecialCaseProtocolHandler:
		return "KSpecialCaseProtocolHandler"
	case KSpecialCaseSocketListener:
		return "KSpecialCaseSocketListener"
	case KSpecialCaseTEDoText:
		return "KSpecialCaseTEDoText"
	case KSpecialCaseTEFindWord:
		return "KSpecialCaseTEFindWord"
	case KSpecialCaseTERecalc:
		return "KSpecialCaseTERecalc"
	case KSpecialCaseTextWidthHook:
		return "KSpecialCaseTextWidthHook"
	default:
		return fmt.Sprintf("KSpecialCase(%d)", e)
	}
}

type KStartupFolderIconResource int

const (
	KControlPanelFolderIconResource KStartupFolderIconResource = -3976
	KDropFolderIconResource         KStartupFolderIconResource = -3979
	KExtensionsFolderIconResource   KStartupFolderIconResource = -3973
	KFontsFolderIconResource        KStartupFolderIconResource = -3968
	KFullTrashIconResource          KStartupFolderIconResource = -3984
	KMountedFolderIconResource      KStartupFolderIconResource = -3977
	KOwnedFolderIconResource        KStartupFolderIconResource = -3980
	KPreferencesFolderIconResource  KStartupFolderIconResource = -3974
	KPrintMonitorFolderIconResource KStartupFolderIconResource = -3975
	KSharedFolderIconResource       KStartupFolderIconResource = -3978
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

type KSystemFolderAliasType uint

const (
	KAppleMenuFolderAliasType        KSystemFolderAliasType = 'f'<<24 | 'a'<<16 | 'a'<<8 | 'm' // 'faam'
	KControlPanelFolderAliasType     KSystemFolderAliasType = 'f'<<24 | 'a'<<16 | 'c'<<8 | 't' // 'fact'
	KExtensionFolderAliasType        KSystemFolderAliasType = 'f'<<24 | 'a'<<16 | 'e'<<8 | 'x' // 'faex'
	KPreferencesFolderAliasType      KSystemFolderAliasType = 'f'<<24 | 'a'<<16 | 'p'<<8 | 'f' // 'fapf'
	KPrintMonitorDocsFolderAliasType KSystemFolderAliasType = 'f'<<24 | 'a'<<16 | 'p'<<8 | 'n' // 'fapn'
	KStartupFolderAliasType          KSystemFolderAliasType = 'f'<<24 | 'a'<<16 | 's'<<8 | 't' // 'fast'
	KSystemFolderAliasTypeValue      KSystemFolderAliasType = 'f'<<24 | 'a'<<16 | 's'<<8 | 'y' // 'fasy'
)

func (e KSystemFolderAliasType) String() string {
	switch e {
	case KAppleMenuFolderAliasType:
		return "KAppleMenuFolderAliasType"
	case KControlPanelFolderAliasType:
		return "KControlPanelFolderAliasType"
	case KExtensionFolderAliasType:
		return "KExtensionFolderAliasType"
	case KPreferencesFolderAliasType:
		return "KPreferencesFolderAliasType"
	case KPrintMonitorDocsFolderAliasType:
		return "KPrintMonitorDocsFolderAliasType"
	case KStartupFolderAliasType:
		return "KStartupFolderAliasType"
	case KSystemFolderAliasTypeValue:
		return "KSystemFolderAliasTypeValue"
	default:
		return fmt.Sprintf("KSystemFolderAliasType(%d)", e)
	}
}

type KSystemFolderType uint

const (
	// Deprecated.
	KALMLocationsFolderType KSystemFolderType = 'f'<<24 | 'a'<<16 | 'l'<<8 | 'l' // 'fall'
	// Deprecated.
	KALMModulesFolderType KSystemFolderType = 'w'<<24 | 'a'<<16 | 'l'<<8 | 'k' // 'walk'
	// Deprecated.
	KALMPreferencesFolderType KSystemFolderType = 't'<<24 | 'r'<<16 | 'i'<<8 | 'p' // 'trip'
	// Deprecated.
	KAppearanceFolderType KSystemFolderType = 'a'<<24 | 'p'<<16 | 'p'<<8 | 'r' // 'appr'
	// Deprecated.
	KAppleExtrasFolderType KSystemFolderType = 1634040004
	// Deprecated.
	KAppleMenuFolderType KSystemFolderType = 'a'<<24 | 'm'<<16 | 'n'<<8 | 'u' // 'amnu'
	// Deprecated.
	KAppleshareAutomountServerAliasesFolderType KSystemFolderType = 1936881348
	// Deprecated.
	KAssistantsFolderType KSystemFolderType = 1634956484
	// Deprecated.
	KContextualMenuItemsFolderType KSystemFolderType = 'c'<<24 | 'm'<<16 | 'n'<<8 | 'u' // 'cmnu'
	// Deprecated.
	KControlPanelDisabledFolderType KSystemFolderType = 'c'<<24 | 't'<<16 | 'r'<<8 | 'D' // 'ctrD'
	// Deprecated.
	KControlPanelFolderType KSystemFolderType = 'c'<<24 | 't'<<16 | 'r'<<8 | 'l' // 'ctrl'
	// Deprecated.
	KControlStripModulesFolderType KSystemFolderType = 's'<<24 | 'd'<<16 | 'e'<<8 | 'v' // 'sdev'
	// Deprecated.
	KDesktopPicturesFolderType KSystemFolderType = 1685352644
	// Deprecated.
	KDisplayExtensionsFolderType KSystemFolderType = 'd'<<24 | 's'<<16 | 'p'<<8 | 'l' // 'dspl'
	// Deprecated.
	KEditorsFolderType KSystemFolderType = 'o'<<24 | 'd'<<16 | 'e'<<8 | 'd' // 'oded'
	// Deprecated.
	KExtensionDisabledFolderType KSystemFolderType = 'e'<<24 | 'x'<<16 | 't'<<8 | 'D' // 'extD'
	// Deprecated.
	KExtensionFolderType KSystemFolderType = 'e'<<24 | 'x'<<16 | 't'<<8 | 'n' // 'extn'
	// Deprecated.
	KFindByContentFolderType KSystemFolderType = 'f'<<24 | 'b'<<16 | 'c'<<8 | 'f' // 'fbcf'
	// Deprecated.
	KFindByContentPluginsFolderType KSystemFolderType = 'f'<<24 | 'b'<<16 | 'c'<<8 | 'p' // 'fbcp'
	// Deprecated.
	KFindSupportFolderType KSystemFolderType = 'f'<<24 | 'n'<<16 | 'd'<<8 | 's' // 'fnds'
	// Deprecated.
	KGenEditorsFolderType KSystemFolderType = 3294979177
	// Deprecated.
	KHelpFolderType KSystemFolderType = 3295177840
	// Deprecated.
	KInternetFolderType KSystemFolderType = 1768846532
	// Deprecated.
	KInternetPlugInFolderType KSystemFolderType = 3295569268
	// Deprecated.
	KLauncherItemsFolderType KSystemFolderType = 'l'<<24 | 'a'<<16 | 'u'<<8 | 'n' // 'laun'
	// Deprecated.
	KLocalesFolderType KSystemFolderType = 3295440739
	// Deprecated.
	KMacOSReadMesFolderType KSystemFolderType = 1836020420
	// Deprecated.
	KModemScriptsFolderType KSystemFolderType = 3295506276
	// Deprecated.
	KMultiprocessingFolderType KSystemFolderType = 'm'<<24 | 'p'<<16 | 'x'<<8 | 'f' // 'mpxf'
	// Deprecated.
	KOpenDocEditorsFolderType KSystemFolderType = 3295634534
	// Deprecated.
	KOpenDocFolderType KSystemFolderType = 'o'<<24 | 'd'<<16 | 'o'<<8 | 'd' // 'odod'
	// Deprecated.
	KOpenDocLibrariesFolderType KSystemFolderType = 'o'<<24 | 'd'<<16 | 'l'<<8 | 'b' // 'odlb'
	// Deprecated.
	KOpenDocShellPlugInsFolderType KSystemFolderType = 'o'<<24 | 'd'<<16 | 's'<<8 | 'p' // 'odsp'
	// Deprecated.
	KPreMacOS91AppleExtrasFolderType KSystemFolderType = 2355460292
	// Deprecated.
	KPreMacOS91ApplicationsFolderType KSystemFolderType = 2356179059
	// Deprecated.
	KPreMacOS91AssistantsFolderType KSystemFolderType = 2356376772
	// Deprecated.
	KPreMacOS91AutomountedServersFolderType KSystemFolderType = 2809296580
	// Deprecated.
	KPreMacOS91InstallerLogsFolderType KSystemFolderType = 2490132326
	// Deprecated.
	KPreMacOS91InternetFolderType KSystemFolderType = 2490266820
	// Deprecated.
	KPreMacOS91MacOSReadMesFolderType KSystemFolderType = 3043979972
	// Deprecated.
	KPreMacOS91StationeryFolderType KSystemFolderType = 3211031412
	// Deprecated.
	KPreMacOS91UtilitiesFolderType KSystemFolderType = 2675206596
	// Deprecated.
	KPrintMonitorDocsFolderType KSystemFolderType = 'p'<<24 | 'r'<<16 | 'n'<<8 | 't' // 'prnt'
	// Deprecated.
	KPrintingPlugInsFolderType KSystemFolderType = 'p'<<24 | 'p'<<16 | 'l'<<8 | 'g' // 'pplg'
	// Deprecated.
	KQuickTimeExtensionsFolderType KSystemFolderType = 'q'<<24 | 't'<<16 | 'e'<<8 | 'x' // 'qtex'
	// Deprecated.
	KRecentApplicationsFolderType KSystemFolderType = 'r'<<24 | 'a'<<16 | 'p'<<8 | 'p' // 'rapp'
	// Deprecated.
	KRecentDocumentsFolderType KSystemFolderType = 'r'<<24 | 'd'<<16 | 'o'<<8 | 'c' // 'rdoc'
	// Deprecated.
	KRecentServersFolderType KSystemFolderType = 'r'<<24 | 's'<<16 | 'v'<<8 | 'r' // 'rsvr'
	// Deprecated.
	KShutdownFolderType KSystemFolderType = 's'<<24 | 'h'<<16 | 'd'<<8 | 'f' // 'shdf'
	// Deprecated.
	KShutdownItemsDisabledFolderType KSystemFolderType = 's'<<24 | 'h'<<16 | 'd'<<8 | 'D' // 'shdD'
	// Deprecated.
	KSoundSetsFolderType KSystemFolderType = 's'<<24 | 'n'<<16 | 'd'<<8 | 's' // 'snds'
	// Deprecated.
	KStartupFolderType KSystemFolderType = 's'<<24 | 't'<<16 | 'r'<<8 | 't' // 'strt'
	// Deprecated.
	KStartupItemsDisabledFolderType KSystemFolderType = 's'<<24 | 't'<<16 | 'r'<<8 | 'D' // 'strD'
	// Deprecated.
	KStationeryFolderType KSystemFolderType = 'o'<<24 | 'd'<<16 | 's'<<8 | 't' // 'odst'
	// Deprecated.
	KSystemControlPanelFolderType KSystemFolderType = 's'<<24 | 'c'<<16 | 't'<<8 | 'l' // 'sctl'
	// Deprecated.
	KSystemDesktopFolderType KSystemFolderType = 's'<<24 | 'd'<<16 | 's'<<8 | 'k' // 'sdsk'
	// Deprecated.
	KSystemExtensionDisabledFolderType KSystemFolderType = 'm'<<24 | 'a'<<16 | 'c'<<8 | 'D' // 'macD'
	// Deprecated.
	KSystemFolderTypeValue KSystemFolderType = 'm'<<24 | 'a'<<16 | 'c'<<8 | 's' // 'macs'
	// Deprecated.
	KSystemTrashFolderType KSystemFolderType = 's'<<24 | 't'<<16 | 'r'<<8 | 's' // 'strs'
	// Deprecated.
	KVolumeSettingsFolderType KSystemFolderType = 'v'<<24 | 's'<<16 | 'f'<<8 | 'd' // 'vsfd'
)

func (e KSystemFolderType) String() string {
	switch e {
	case KALMLocationsFolderType:
		return "KALMLocationsFolderType"
	case KALMModulesFolderType:
		return "KALMModulesFolderType"
	case KALMPreferencesFolderType:
		return "KALMPreferencesFolderType"
	case KAppearanceFolderType:
		return "KAppearanceFolderType"
	case KAppleExtrasFolderType:
		return "KAppleExtrasFolderType"
	case KAppleMenuFolderType:
		return "KAppleMenuFolderType"
	case KAppleshareAutomountServerAliasesFolderType:
		return "KAppleshareAutomountServerAliasesFolderType"
	case KAssistantsFolderType:
		return "KAssistantsFolderType"
	case KContextualMenuItemsFolderType:
		return "KContextualMenuItemsFolderType"
	case KControlPanelDisabledFolderType:
		return "KControlPanelDisabledFolderType"
	case KControlPanelFolderType:
		return "KControlPanelFolderType"
	case KControlStripModulesFolderType:
		return "KControlStripModulesFolderType"
	case KDesktopPicturesFolderType:
		return "KDesktopPicturesFolderType"
	case KDisplayExtensionsFolderType:
		return "KDisplayExtensionsFolderType"
	case KEditorsFolderType:
		return "KEditorsFolderType"
	case KExtensionDisabledFolderType:
		return "KExtensionDisabledFolderType"
	case KExtensionFolderType:
		return "KExtensionFolderType"
	case KFindByContentFolderType:
		return "KFindByContentFolderType"
	case KFindByContentPluginsFolderType:
		return "KFindByContentPluginsFolderType"
	case KFindSupportFolderType:
		return "KFindSupportFolderType"
	case KGenEditorsFolderType:
		return "KGenEditorsFolderType"
	case KHelpFolderType:
		return "KHelpFolderType"
	case KInternetFolderType:
		return "KInternetFolderType"
	case KInternetPlugInFolderType:
		return "KInternetPlugInFolderType"
	case KLauncherItemsFolderType:
		return "KLauncherItemsFolderType"
	case KLocalesFolderType:
		return "KLocalesFolderType"
	case KMacOSReadMesFolderType:
		return "KMacOSReadMesFolderType"
	case KModemScriptsFolderType:
		return "KModemScriptsFolderType"
	case KMultiprocessingFolderType:
		return "KMultiprocessingFolderType"
	case KOpenDocEditorsFolderType:
		return "KOpenDocEditorsFolderType"
	case KOpenDocFolderType:
		return "KOpenDocFolderType"
	case KOpenDocLibrariesFolderType:
		return "KOpenDocLibrariesFolderType"
	case KOpenDocShellPlugInsFolderType:
		return "KOpenDocShellPlugInsFolderType"
	case KPreMacOS91AppleExtrasFolderType:
		return "KPreMacOS91AppleExtrasFolderType"
	case KPreMacOS91ApplicationsFolderType:
		return "KPreMacOS91ApplicationsFolderType"
	case KPreMacOS91AssistantsFolderType:
		return "KPreMacOS91AssistantsFolderType"
	case KPreMacOS91AutomountedServersFolderType:
		return "KPreMacOS91AutomountedServersFolderType"
	case KPreMacOS91InstallerLogsFolderType:
		return "KPreMacOS91InstallerLogsFolderType"
	case KPreMacOS91InternetFolderType:
		return "KPreMacOS91InternetFolderType"
	case KPreMacOS91MacOSReadMesFolderType:
		return "KPreMacOS91MacOSReadMesFolderType"
	case KPreMacOS91StationeryFolderType:
		return "KPreMacOS91StationeryFolderType"
	case KPreMacOS91UtilitiesFolderType:
		return "KPreMacOS91UtilitiesFolderType"
	case KPrintMonitorDocsFolderType:
		return "KPrintMonitorDocsFolderType"
	case KPrintingPlugInsFolderType:
		return "KPrintingPlugInsFolderType"
	case KQuickTimeExtensionsFolderType:
		return "KQuickTimeExtensionsFolderType"
	case KRecentApplicationsFolderType:
		return "KRecentApplicationsFolderType"
	case KRecentDocumentsFolderType:
		return "KRecentDocumentsFolderType"
	case KRecentServersFolderType:
		return "KRecentServersFolderType"
	case KShutdownFolderType:
		return "KShutdownFolderType"
	case KShutdownItemsDisabledFolderType:
		return "KShutdownItemsDisabledFolderType"
	case KSoundSetsFolderType:
		return "KSoundSetsFolderType"
	case KStartupFolderType:
		return "KStartupFolderType"
	case KStartupItemsDisabledFolderType:
		return "KStartupItemsDisabledFolderType"
	case KStationeryFolderType:
		return "KStationeryFolderType"
	case KSystemControlPanelFolderType:
		return "KSystemControlPanelFolderType"
	case KSystemDesktopFolderType:
		return "KSystemDesktopFolderType"
	case KSystemExtensionDisabledFolderType:
		return "KSystemExtensionDisabledFolderType"
	case KSystemFolderTypeValue:
		return "KSystemFolderTypeValue"
	case KSystemTrashFolderType:
		return "KSystemTrashFolderType"
	case KVolumeSettingsFolderType:
		return "KVolumeSettingsFolderType"
	default:
		return fmt.Sprintf("KSystemFolderType(%d)", e)
	}
}

type KSystemIcons uint

const (
	KSystemIconsCreator KSystemIcons = 'm'<<24 | 'a'<<16 | 'c'<<8 | 's' // 'macs'
)

func (e KSystemIcons) String() string {
	switch e {
	case KSystemIconsCreator:
		return "KSystemIconsCreator"
	default:
		return fmt.Sprintf("KSystemIcons(%d)", e)
	}
}

type KTEC int

const (
	KTECAddFallbackInterruptBit  KTEC = 7
	KTECAddFallbackInterruptMask KTEC = 128
	// KTECAddForceASCIIChangesBit: This is set if the new control flag bits `kUnicodeForceASCIIRangeBit` and `kUnicodeNoHalfwidthCharsBit` are supported for use with the functions [ConvertFromTextToUnicode], [ConvertFromUnicodeToText], and so forth.
	KTECAddForceASCIIChangesBit   KTEC = 4
	KTECAddForceASCIIChangesMask  KTEC = 16
	KTECAddTextRunHeuristicsBit   KTEC = 6
	KTECAddTextRunHeuristicsMask  KTEC = 64
	KTECAvailableEncodingsResType KTEC = 'c'<<24 | 'v'<<16 | 'e'<<8 | 'n' // 'cven'
	KTECAvailableSniffersResType  KTEC = 'c'<<24 | 'v'<<16 | 's'<<8 | 'f' // 'cvsf'
	KTECChinesePluginSignature    KTEC = 'p'<<24 | 'z'<<16 | 'h'<<8 | 'o' // 'pzho'
	KTECConversionInfoResType     KTEC = 'c'<<24 | 'v'<<16 | 'i'<<8 | 'f' // 'cvif'
	// KTECFallbackTextLengthFixBit: # Discussion
	KTECFallbackTextLengthFixBit  KTEC = 1
	KTECFallbackTextLengthFixMask KTEC = 2
	KTECInternetNamesResType      KTEC = 'c'<<24 | 'v'<<16 | 'm'<<8 | 'm' // 'cvmm'
	KTECJapanesePluginSignature   KTEC = 'p'<<24 | 'j'<<16 | 'p'<<8 | 'n' // 'pjpn'
	// KTECKeepInfoFixBit: This is set if the Unicode Converter has a bug fix to stop ignoring certain control flags
	KTECKeepInfoFixBit        KTEC = 0
	KTECKeepInfoFixMask       KTEC = 1
	KTECKoreanPluginSignature KTEC = 'p'<<24 | 'k'<<16 | 'o'<<8 | 'r' // 'pkor'
	KTECMailEncodingsResType  KTEC = 'c'<<24 | 'v'<<16 | 'm'<<8 | 'l' // 'cvml'
	// KTECPreferredEncodingFixBit: This is set to indicate that if a preferred encoding is specified for [CreateUnicodeToTextRunInfo] and related functions, they handle it correctly even if it does not match the system script.
	KTECPreferredEncodingFixBit  KTEC = 5
	KTECPreferredEncodingFixMask KTEC = 32
	KTECSignature                KTEC = 'e'<<24 | 'n'<<16 | 'c'<<8 | 'v' // 'encv'
	KTECSubTextEncodingsResType  KTEC = 'c'<<24 | 'v'<<16 | 's'<<8 | 'b' // 'cvsb'
	// KTECTextRunBitClearFixBit: This is set if [ConvertFromUnicodeToTextRun] and [ConvertFromUnicodeToScriptCodeRun] function correctly if the `kUnicodeTextRunBit` is clear.
	KTECTextRunBitClearFixBit  KTEC = 2
	KTECTextRunBitClearFixMask KTEC = 4
	// KTECTextToUnicodeScanFixBit: # Discussion
	KTECTextToUnicodeScanFixBit  KTEC = 3
	KTECTextToUnicodeScanFixMask KTEC = 8
	KTECUnicodePluginSignature   KTEC = 'p'<<24 | 'u'<<16 | 'n'<<8 | 'i' // 'puni'
	KTECWebEncodingsResType      KTEC = 'c'<<24 | 'v'<<16 | 'w'<<8 | 'b' // 'cvwb'
	KTEC_MIBEnumDontCare         KTEC = -1
)

func (e KTEC) String() string {
	switch e {
	case KTECAddFallbackInterruptBit:
		return "KTECAddFallbackInterruptBit"
	case KTECAddFallbackInterruptMask:
		return "KTECAddFallbackInterruptMask"
	case KTECAddForceASCIIChangesBit:
		return "KTECAddForceASCIIChangesBit"
	case KTECAddForceASCIIChangesMask:
		return "KTECAddForceASCIIChangesMask"
	case KTECAddTextRunHeuristicsBit:
		return "KTECAddTextRunHeuristicsBit"
	case KTECAddTextRunHeuristicsMask:
		return "KTECAddTextRunHeuristicsMask"
	case KTECAvailableEncodingsResType:
		return "KTECAvailableEncodingsResType"
	case KTECAvailableSniffersResType:
		return "KTECAvailableSniffersResType"
	case KTECChinesePluginSignature:
		return "KTECChinesePluginSignature"
	case KTECConversionInfoResType:
		return "KTECConversionInfoResType"
	case KTECFallbackTextLengthFixBit:
		return "KTECFallbackTextLengthFixBit"
	case KTECFallbackTextLengthFixMask:
		return "KTECFallbackTextLengthFixMask"
	case KTECInternetNamesResType:
		return "KTECInternetNamesResType"
	case KTECJapanesePluginSignature:
		return "KTECJapanesePluginSignature"
	case KTECKeepInfoFixBit:
		return "KTECKeepInfoFixBit"
	case KTECKoreanPluginSignature:
		return "KTECKoreanPluginSignature"
	case KTECMailEncodingsResType:
		return "KTECMailEncodingsResType"
	case KTECPreferredEncodingFixBit:
		return "KTECPreferredEncodingFixBit"
	case KTECPreferredEncodingFixMask:
		return "KTECPreferredEncodingFixMask"
	case KTECSignature:
		return "KTECSignature"
	case KTECSubTextEncodingsResType:
		return "KTECSubTextEncodingsResType"
	case KTECTextToUnicodeScanFixBit:
		return "KTECTextToUnicodeScanFixBit"
	case KTECTextToUnicodeScanFixMask:
		return "KTECTextToUnicodeScanFixMask"
	case KTECUnicodePluginSignature:
		return "KTECUnicodePluginSignature"
	case KTECWebEncodingsResType:
		return "KTECWebEncodingsResType"
	case KTEC_MIBEnumDontCare:
		return "KTEC_MIBEnumDontCare"
	default:
		return fmt.Sprintf("KTEC(%d)", e)
	}
}

type KTECDisable uint

const (
	KTECDisableFallbacksBit      KTECDisable = 16
	KTECDisableFallbacksMask     KTECDisable = 65536
	KTECDisableLooseMappingsBit  KTECDisable = 17
	KTECDisableLooseMappingsMask KTECDisable = 131072
)

func (e KTECDisable) String() string {
	switch e {
	case KTECDisableFallbacksBit:
		return "KTECDisableFallbacksBit"
	case KTECDisableFallbacksMask:
		return "KTECDisableFallbacksMask"
	case KTECDisableLooseMappingsBit:
		return "KTECDisableLooseMappingsBit"
	case KTECDisableLooseMappingsMask:
		return "KTECDisableLooseMappingsMask"
	default:
		return fmt.Sprintf("KTECDisable(%d)", e)
	}
}

type KTECInfoCurrent uint

const (
	KTECInfoCurrentFormat KTECInfoCurrent = 2
)

func (e KTECInfoCurrent) String() string {
	switch e {
	case KTECInfoCurrentFormat:
		return "KTECInfoCurrentFormat"
	default:
		return fmt.Sprintf("KTECInfoCurrent(%d)", e)
	}
}

type KTECInternetName uint

const (
	KTECInternetNameDefaultUsageMask  KTECInternetName = 0
	KTECInternetNameStrictUsageMask   KTECInternetName = 1
	KTECInternetNameTolerantUsageMask KTECInternetName = 2
)

func (e KTECInternetName) String() string {
	switch e {
	case KTECInternetNameDefaultUsageMask:
		return "KTECInternetNameDefaultUsageMask"
	case KTECInternetNameStrictUsageMask:
		return "KTECInternetNameStrictUsageMask"
	case KTECInternetNameTolerantUsageMask:
		return "KTECInternetNameTolerantUsageMask"
	default:
		return fmt.Sprintf("KTECInternetName(%d)", e)
	}
}

type KTECPlugin uint

const (
	KTECPluginCreator   KTECPlugin = 'e'<<24 | 'n'<<16 | 'c'<<8 | 'v' // 'encv'
	KTECPluginManyToOne KTECPlugin = 'm'<<24 | 't'<<16 | 'o'<<8 | 'o' // 'mtoo'
	KTECPluginOneToMany KTECPlugin = 'o'<<24 | 't'<<16 | 'o'<<8 | 'm' // 'otom'
	KTECPluginOneToOne  KTECPlugin = 'o'<<24 | 't'<<16 | 'o'<<8 | 'o' // 'otoo'
	KTECPluginSniffObj  KTECPlugin = 's'<<24 | 'n'<<16 | 'i'<<8 | 'f' // 'snif'
	KTECPluginType      KTECPlugin = 'e'<<24 | 'c'<<16 | 'p'<<8 | 'g' // 'ecpg'
)

func (e KTECPlugin) String() string {
	switch e {
	case KTECPluginCreator:
		return "KTECPluginCreator"
	case KTECPluginManyToOne:
		return "KTECPluginManyToOne"
	case KTECPluginOneToMany:
		return "KTECPluginOneToMany"
	case KTECPluginOneToOne:
		return "KTECPluginOneToOne"
	case KTECPluginSniffObj:
		return "KTECPluginSniffObj"
	case KTECPluginType:
		return "KTECPluginType"
	default:
		return fmt.Sprintf("KTECPlugin(%d)", e)
	}
}

type KTECPluginDispatchTable uint

const (
	// KTECPluginDispatchTableCurrentVersion: A meta value that specifies the current version.
	KTECPluginDispatchTableCurrentVersion KTECPluginDispatchTable = 65538
	// KTECPluginDispatchTableVersion1: Specifies versions 1.0 through 1.0.3.
	KTECPluginDispatchTableVersion1 KTECPluginDispatchTable = 65536
	// KTECPluginDispatchTableVersion1_1: Specifies version 1.1.
	KTECPluginDispatchTableVersion1_1 KTECPluginDispatchTable = 65537
	// KTECPluginDispatchTableVersion1_2: Specifies version 1.2.
	KTECPluginDispatchTableVersion1_2 KTECPluginDispatchTable = 65538
)

func (e KTECPluginDispatchTable) String() string {
	switch e {
	case KTECPluginDispatchTableCurrentVersion:
		return "KTECPluginDispatchTableCurrentVersion"
	case KTECPluginDispatchTableVersion1:
		return "KTECPluginDispatchTableVersion1"
	case KTECPluginDispatchTableVersion1_1:
		return "KTECPluginDispatchTableVersion1_1"
	default:
		return fmt.Sprintf("KTECPluginDispatchTable(%d)", e)
	}
}

type KTMTask uint

const (
	// Deprecated.
	KTMTaskActive KTMTask = 32768
)

func (e KTMTask) String() string {
	switch e {
	case KTMTaskActive:
		return "KTMTaskActive"
	default:
		return fmt.Sprintf("KTMTask(%d)", e)
	}
}

type KTSM uint

const (
	KTSMInsideOfActiveInputArea KTSM = 3
	KTSMInsideOfBody            KTSM = 2
	KTSMOutsideOfBody           KTSM = 1
)

func (e KTSM) String() string {
	switch e {
	case KTSMInsideOfActiveInputArea:
		return "KTSMInsideOfActiveInputArea"
	case KTSMInsideOfBody:
		return "KTSMInsideOfBody"
	case KTSMOutsideOfBody:
		return "KTSMOutsideOfBody"
	default:
		return fmt.Sprintf("KTSM(%d)", e)
	}
}

type KTSMHilite uint

const (
	// KTSMHiliteBlockFillText: Specifies block fill highlight style.
	KTSMHiliteBlockFillText KTSMHilite = 6
	// KTSMHiliteCaretPosition: Specifies caret position.
	KTSMHiliteCaretPosition KTSMHilite = 1
	// KTSMHiliteConvertedText: Specifies range of converted text.
	KTSMHiliteConvertedText KTSMHilite = 4
	// KTSMHiliteNoHilite: Specifies range of non-highlighted text.
	KTSMHiliteNoHilite KTSMHilite = 9
	// KTSMHiliteOutlineText: Specifies outline highlight style.
	KTSMHiliteOutlineText KTSMHilite = 7
	// KTSMHiliteRawText: Specifies range of raw text.
	KTSMHiliteRawText KTSMHilite = 2
	// KTSMHiliteSelectedConvertedText: Specifies range of selected converted text.
	KTSMHiliteSelectedConvertedText KTSMHilite = 5
	// KTSMHiliteSelectedRawText: Specifies range of selected raw text.
	KTSMHiliteSelectedRawText KTSMHilite = 3
	// KTSMHiliteSelectedText: Specifies selected highlight style.
	KTSMHiliteSelectedText KTSMHilite = 8
)

func (e KTSMHilite) String() string {
	switch e {
	case KTSMHiliteBlockFillText:
		return "KTSMHiliteBlockFillText"
	case KTSMHiliteCaretPosition:
		return "KTSMHiliteCaretPosition"
	case KTSMHiliteConvertedText:
		return "KTSMHiliteConvertedText"
	case KTSMHiliteNoHilite:
		return "KTSMHiliteNoHilite"
	case KTSMHiliteOutlineText:
		return "KTSMHiliteOutlineText"
	case KTSMHiliteRawText:
		return "KTSMHiliteRawText"
	case KTSMHiliteSelectedConvertedText:
		return "KTSMHiliteSelectedConvertedText"
	case KTSMHiliteSelectedRawText:
		return "KTSMHiliteSelectedRawText"
	case KTSMHiliteSelectedText:
		return "KTSMHiliteSelectedText"
	default:
		return fmt.Sprintf("KTSMHilite(%d)", e)
	}
}

type KTXN int

const (
	KTXNATSUIIsNotInstalledErr              KTXN = -22016
	KTXNAlreadyInitializedErr               KTXN = -22012
	KTXNAttributeTagInvalidForRunErr        KTXN = -22009
	KTXNBadDefaultFileTypeWarning           KTXN = -22005
	KTXNCannotAddFrameErr                   KTXN = -22001
	KTXNCannotSetAutoIndentErr              KTXN = -22006
	KTXNCannotTurnTSMOffWhenUsingUnicodeErr KTXN = -22013
	KTXNCopyNotAllowedInEchoModeErr         KTXN = -22014
	KTXNDataTypeNotAllowedErr               KTXN = -22015
	KTXNEndIterationErr                     KTXN = -22000
	KTXNIllegalToCrossDataBoundariesErr     KTXN = -22003
	KTXNInvalidFrameIDErr                   KTXN = -22002
	KTXNInvalidRunIndex                     KTXN = -22011
	KTXNNoMatchErr                          KTXN = -22008
	KTXNOutsideOfFrameErr                   KTXN = -22018
	KTXNOutsideOfLineErr                    KTXN = -22017
	KTXNRunIndexOutofBoundsErr              KTXN = -22007
	KTXNSomeOrAllTagsInvalidForRunErr       KTXN = -22010
	KTXNUserCanceledOperationErr            KTXN = -22004
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

type KText int

const (
	KTextCenter       KText = 1
	KTextFlushDefault KText = 0
	KTextFlushLeft    KText = -2
	KTextFlushRight   KText = -1
	// KTextLanguageDontCare: Indicates that language code is not provided for the derivation.
	KTextLanguageDontCare KText = -128
	// KTextRegionDontCare: The region code is not provided for the derivation.
	KTextRegionDontCare KText = -128
	// KTextScriptDontCare: Indicates that the code is not provided for the derivation.
	KTextScriptDontCare KText = -128
)

func (e KText) String() string {
	switch e {
	case KTextCenter:
		return "KTextCenter"
	case KTextFlushDefault:
		return "KTextFlushDefault"
	case KTextFlushLeft:
		return "KTextFlushLeft"
	case KTextFlushRight:
		return "KTextFlushRight"
	case KTextLanguageDontCare:
		return "KTextLanguageDontCare"
	default:
		return fmt.Sprintf("KText(%d)", e)
	}
}

type KTextEncoding uint

const (
	KTextEncodingANSEL KTextEncoding = 1537
	// KTextEncodingBaseName: Requests the name of the base encoding.
	KTextEncodingBaseName KTextEncoding = 1
	// KTextEncodingBig5: Big-5 encoding.
	KTextEncodingBig5            KTextEncoding = 2563
	KTextEncodingBig5_E          KTextEncoding = 2569
	KTextEncodingBig5_HKSCS_1999 KTextEncoding = 2566
	// KTextEncodingCNS_11643_92_P1: CNS 11643-1992 plane 1.
	KTextEncodingCNS_11643_92_P1 KTextEncoding = 1617
	// KTextEncodingCNS_11643_92_P2: CNS 11643-1992 plane 2.
	KTextEncodingCNS_11643_92_P2 KTextEncoding = 1618
	// KTextEncodingCNS_11643_92_P3: CNS 11643-1992 plane 3 (11643-1986 plane 14).
	KTextEncodingCNS_11643_92_P3 KTextEncoding = 1619
	// KTextEncodingDOSArabic: Code page 864.
	KTextEncodingDOSArabic KTextEncoding = 1049
	// KTextEncodingDOSBalticRim: Code page 775.
	KTextEncodingDOSBalticRim KTextEncoding = 1030
	// KTextEncodingDOSCanadianFrench: Code page 863.
	KTextEncodingDOSCanadianFrench KTextEncoding = 1048
	// KTextEncodingDOSChineseSimplif: Code page 936, also for Windows.
	KTextEncodingDOSChineseSimplif KTextEncoding = 1057
	// KTextEncodingDOSChineseTrad: Code page 950, also for Windows.
	KTextEncodingDOSChineseTrad KTextEncoding = 1059
	// KTextEncodingDOSCyrillic: Code page 855, IBM Cyrillic.
	KTextEncodingDOSCyrillic KTextEncoding = 1043
	// KTextEncodingDOSGreek: Code page 737, formerly 437G.
	KTextEncodingDOSGreek KTextEncoding = 1029
	// KTextEncodingDOSGreek1: Code page 851.
	KTextEncodingDOSGreek1 KTextEncoding = 1041
	// KTextEncodingDOSGreek2: Code page 869, IBM Modern Green.
	KTextEncodingDOSGreek2 KTextEncoding = 1052
	// KTextEncodingDOSHebrew: Code page 862.
	KTextEncodingDOSHebrew KTextEncoding = 1047
	// KTextEncodingDOSIcelandic: Code page 861.
	KTextEncodingDOSIcelandic KTextEncoding = 1046
	// KTextEncodingDOSJapanese: Code page 932, also for Windows
	KTextEncodingDOSJapanese KTextEncoding = 1056
	// KTextEncodingDOSKorean: Code page 949, also for Windows; unified Hangul.
	KTextEncodingDOSKorean KTextEncoding = 1058
	// KTextEncodingDOSLatin1: Code page 860.
	KTextEncodingDOSLatin1 KTextEncoding = 1040
	// KTextEncodingDOSLatin2: Code page 852, Slavic.
	KTextEncodingDOSLatin2 KTextEncoding = 1042
	// KTextEncodingDOSLatinUS: Code page 437.
	KTextEncodingDOSLatinUS KTextEncoding = 1024
	// KTextEncodingDOSNordic: Cde page 865.
	KTextEncodingDOSNordic KTextEncoding = 1050
	// KTextEncodingDOSPortuguese: Code page 860.
	KTextEncodingDOSPortuguese KTextEncoding = 1045
	// KTextEncodingDOSRussian: Code page 866.
	KTextEncodingDOSRussian KTextEncoding = 1051
	// KTextEncodingDOSThai: Code page 874, also for Windows.
	KTextEncodingDOSThai KTextEncoding = 1053
	// KTextEncodingDOSTurkish: Code page 857, IBM Turkish.
	KTextEncodingDOSTurkish KTextEncoding = 1044
	// KTextEncodingFormatName: Requests the name of the encoding format, if available.
	KTextEncodingFormatName KTextEncoding = 3
	// KTextEncodingFullName: Requests the full name of the text encoding.
	KTextEncodingFullName KTextEncoding = 0
	// KTextEncodingGBK_95: Annex to GB13000-93, for Windows 95; EUC-CN extended.
	KTextEncodingGBK_95        KTextEncoding = 1585
	KTextEncodingGB_18030_2000 KTextEncoding = 1586
	KTextEncodingGB_18030_2005 KTextEncoding = 1586
	KTextEncodingGB_2312_80    KTextEncoding = 1584
	// KTextEncodingHZ_GB_2312: See RFC 1842; for Chinese mail and news.
	KTextEncodingHZ_GB_2312 KTextEncoding = 2565
	// KTextEncodingISO10646_1993: This ISO UCS encoding has code points identical to Unicode 1.1.
	KTextEncodingISO10646_1993 KTextEncoding = 257
	KTextEncodingJIS_C6226_78  KTextEncoding = 1572
	// KTextEncodingJIS_X0201_76: JIS Roman and 1-byte katakana (halfwidth).
	KTextEncodingJIS_X0201_76       KTextEncoding = 1568
	KTextEncodingJIS_X0208_83       KTextEncoding = 1569
	KTextEncodingJIS_X0208_90       KTextEncoding = 1570
	KTextEncodingJIS_X0212_90       KTextEncoding = 1571
	KTextEncodingJIS_X0213_MenKuTen KTextEncoding = 1577
	// KTextEncodingKOI8_R: Russian Internet standard.
	KTextEncodingKOI8_R KTextEncoding = 2562
	KTextEncodingKOI8_U KTextEncoding = 2568
	// KTextEncodingKSC_5601_87: This is the same as KSC 5601-92 without Johab annex.
	KTextEncodingKSC_5601_87 KTextEncoding = 1600
	// KTextEncodingKSC_5601_92_Johab: KSC 5601-92 Johab annex.
	KTextEncodingKSC_5601_92_Johab KTextEncoding = 1601
	// KTextEncodingMacRomanLatin1: Mac OS Roman permuted to align with 8859-1.
	KTextEncodingMacRomanLatin1 KTextEncoding = 2564
	// KTextEncodingMultiRun: This is a special value for multiple encoded text, external run information.
	KTextEncodingMultiRun KTextEncoding = 4095
	// KTextEncodingShiftJIS: Plain Shift-JIS.
	KTextEncodingShiftJIS       KTextEncoding = 2561
	KTextEncodingShiftJIS_X0213 KTextEncoding = 1576
	KTextEncodingUS_ASCII       KTextEncoding = 1536
	// KTextEncodingUnicodeDefault: This is a meta value that takes on one of the following values, depending on the system.
	KTextEncodingUnicodeDefault KTextEncoding = 256
	KTextEncodingUnicodeV10_0   KTextEncoding = 276
	KTextEncodingUnicodeV11_0   KTextEncoding = 277
	KTextEncodingUnicodeV12_1   KTextEncoding = 278
	KTextEncodingUnicodeV13_0   KTextEncoding = 279
	KTextEncodingUnicodeV14_0   KTextEncoding = 280
	KTextEncodingUnicodeV15_0   KTextEncoding = 281
	KTextEncodingUnicodeV15_1   KTextEncoding = 282
	// KTextEncodingUnicodeV1_1: This is a Unicode encoding.
	KTextEncodingUnicodeV1_1 KTextEncoding = 257
	// KTextEncodingUnicodeV2_0: This is the new location for Korean Hangul.
	KTextEncodingUnicodeV2_0 KTextEncoding = 259
	// KTextEncodingUnicodeV2_1: For the Text Encoding Converter, Unicode 2.0 is equivalent to Unicode 2.1.
	KTextEncodingUnicodeV2_1 KTextEncoding = 259
	KTextEncodingUnicodeV3_0 KTextEncoding = 260
	// KTextEncodingUnicodeV3_1: Adds characters requiring surrogate pairs in UTF-16
	KTextEncodingUnicodeV3_1 KTextEncoding = 261
	KTextEncodingUnicodeV3_2 KTextEncoding = 262
	KTextEncodingUnicodeV4_0 KTextEncoding = 264
	KTextEncodingUnicodeV5_0 KTextEncoding = 266
	KTextEncodingUnicodeV5_1 KTextEncoding = 267
	KTextEncodingUnicodeV6_0 KTextEncoding = 269
	KTextEncodingUnicodeV6_1 KTextEncoding = 270
	KTextEncodingUnicodeV6_3 KTextEncoding = 272
	KTextEncodingUnicodeV7_0 KTextEncoding = 273
	KTextEncodingUnicodeV8_0 KTextEncoding = 274
	KTextEncodingUnicodeV9_0 KTextEncoding = 275
	KTextEncodingUnknown     KTextEncoding = 65535
	KTextEncodingVISCII      KTextEncoding = 2567
	// KTextEncodingVariantName: Requests the name of the encoding variant, if available.
	KTextEncodingVariantName KTextEncoding = 2
	// KTextEncodingWindowsANSI: Code page 1252 (alternate name).
	KTextEncodingWindowsANSI KTextEncoding = 1280
	// KTextEncodingWindowsArabic: Code page 1256.
	KTextEncodingWindowsArabic KTextEncoding = 1286
	// KTextEncodingWindowsBalticRim: Code page 1257.
	KTextEncodingWindowsBalticRim KTextEncoding = 1287
	// KTextEncodingWindowsCyrillic: Code page 1251, Slavic Cyrillic.
	KTextEncodingWindowsCyrillic KTextEncoding = 1282
	// KTextEncodingWindowsGreek: Code page 1253.
	KTextEncodingWindowsGreek KTextEncoding = 1283
	// KTextEncodingWindowsHebrew: Code page 1255.
	KTextEncodingWindowsHebrew KTextEncoding = 1285
	// KTextEncodingWindowsKoreanJohab: Code page 1361, for Window NT.
	KTextEncodingWindowsKoreanJohab KTextEncoding = 1296
	// KTextEncodingWindowsLatin1: Code page 1252.
	KTextEncodingWindowsLatin1 KTextEncoding = 1280
	// KTextEncodingWindowsLatin2: Code page 1250, Central Europe.
	KTextEncodingWindowsLatin2 KTextEncoding = 1281
	// KTextEncodingWindowsLatin5: Code page 1254, Turkish.
	KTextEncodingWindowsLatin5 KTextEncoding = 1284
	// KTextEncodingWindowsVietnamese: Code page 1258.
	KTextEncodingWindowsVietnamese KTextEncoding = 1288
)

func (e KTextEncoding) String() string {
	switch e {
	case KTextEncodingANSEL:
		return "KTextEncodingANSEL"
	case KTextEncodingBaseName:
		return "KTextEncodingBaseName"
	case KTextEncodingBig5:
		return "KTextEncodingBig5"
	case KTextEncodingBig5_E:
		return "KTextEncodingBig5_E"
	case KTextEncodingBig5_HKSCS_1999:
		return "KTextEncodingBig5_HKSCS_1999"
	case KTextEncodingCNS_11643_92_P1:
		return "KTextEncodingCNS_11643_92_P1"
	case KTextEncodingCNS_11643_92_P2:
		return "KTextEncodingCNS_11643_92_P2"
	case KTextEncodingCNS_11643_92_P3:
		return "KTextEncodingCNS_11643_92_P3"
	case KTextEncodingDOSArabic:
		return "KTextEncodingDOSArabic"
	case KTextEncodingDOSBalticRim:
		return "KTextEncodingDOSBalticRim"
	case KTextEncodingDOSCanadianFrench:
		return "KTextEncodingDOSCanadianFrench"
	case KTextEncodingDOSChineseSimplif:
		return "KTextEncodingDOSChineseSimplif"
	case KTextEncodingDOSChineseTrad:
		return "KTextEncodingDOSChineseTrad"
	case KTextEncodingDOSCyrillic:
		return "KTextEncodingDOSCyrillic"
	case KTextEncodingDOSGreek:
		return "KTextEncodingDOSGreek"
	case KTextEncodingDOSGreek1:
		return "KTextEncodingDOSGreek1"
	case KTextEncodingDOSGreek2:
		return "KTextEncodingDOSGreek2"
	case KTextEncodingDOSHebrew:
		return "KTextEncodingDOSHebrew"
	case KTextEncodingDOSIcelandic:
		return "KTextEncodingDOSIcelandic"
	case KTextEncodingDOSJapanese:
		return "KTextEncodingDOSJapanese"
	case KTextEncodingDOSKorean:
		return "KTextEncodingDOSKorean"
	case KTextEncodingDOSLatin1:
		return "KTextEncodingDOSLatin1"
	case KTextEncodingDOSLatin2:
		return "KTextEncodingDOSLatin2"
	case KTextEncodingDOSLatinUS:
		return "KTextEncodingDOSLatinUS"
	case KTextEncodingDOSNordic:
		return "KTextEncodingDOSNordic"
	case KTextEncodingDOSPortuguese:
		return "KTextEncodingDOSPortuguese"
	case KTextEncodingDOSRussian:
		return "KTextEncodingDOSRussian"
	case KTextEncodingDOSThai:
		return "KTextEncodingDOSThai"
	case KTextEncodingDOSTurkish:
		return "KTextEncodingDOSTurkish"
	case KTextEncodingFormatName:
		return "KTextEncodingFormatName"
	case KTextEncodingFullName:
		return "KTextEncodingFullName"
	case KTextEncodingGBK_95:
		return "KTextEncodingGBK_95"
	case KTextEncodingGB_18030_2000:
		return "KTextEncodingGB_18030_2000"
	case KTextEncodingGB_2312_80:
		return "KTextEncodingGB_2312_80"
	case KTextEncodingHZ_GB_2312:
		return "KTextEncodingHZ_GB_2312"
	case KTextEncodingISO10646_1993:
		return "KTextEncodingISO10646_1993"
	case KTextEncodingJIS_C6226_78:
		return "KTextEncodingJIS_C6226_78"
	case KTextEncodingJIS_X0201_76:
		return "KTextEncodingJIS_X0201_76"
	case KTextEncodingJIS_X0208_83:
		return "KTextEncodingJIS_X0208_83"
	case KTextEncodingJIS_X0208_90:
		return "KTextEncodingJIS_X0208_90"
	case KTextEncodingJIS_X0212_90:
		return "KTextEncodingJIS_X0212_90"
	case KTextEncodingJIS_X0213_MenKuTen:
		return "KTextEncodingJIS_X0213_MenKuTen"
	case KTextEncodingKOI8_R:
		return "KTextEncodingKOI8_R"
	case KTextEncodingKOI8_U:
		return "KTextEncodingKOI8_U"
	case KTextEncodingKSC_5601_87:
		return "KTextEncodingKSC_5601_87"
	case KTextEncodingKSC_5601_92_Johab:
		return "KTextEncodingKSC_5601_92_Johab"
	case KTextEncodingMacRomanLatin1:
		return "KTextEncodingMacRomanLatin1"
	case KTextEncodingMultiRun:
		return "KTextEncodingMultiRun"
	case KTextEncodingShiftJIS:
		return "KTextEncodingShiftJIS"
	case KTextEncodingShiftJIS_X0213:
		return "KTextEncodingShiftJIS_X0213"
	case KTextEncodingUS_ASCII:
		return "KTextEncodingUS_ASCII"
	case KTextEncodingUnicodeDefault:
		return "KTextEncodingUnicodeDefault"
	case KTextEncodingUnicodeV10_0:
		return "KTextEncodingUnicodeV10_0"
	case KTextEncodingUnicodeV11_0:
		return "KTextEncodingUnicodeV11_0"
	case KTextEncodingUnicodeV12_1:
		return "KTextEncodingUnicodeV12_1"
	case KTextEncodingUnicodeV13_0:
		return "KTextEncodingUnicodeV13_0"
	case KTextEncodingUnicodeV14_0:
		return "KTextEncodingUnicodeV14_0"
	case KTextEncodingUnicodeV15_0:
		return "KTextEncodingUnicodeV15_0"
	case KTextEncodingUnicodeV15_1:
		return "KTextEncodingUnicodeV15_1"
	case KTextEncodingUnicodeV2_0:
		return "KTextEncodingUnicodeV2_0"
	case KTextEncodingUnicodeV3_0:
		return "KTextEncodingUnicodeV3_0"
	case KTextEncodingUnicodeV3_1:
		return "KTextEncodingUnicodeV3_1"
	case KTextEncodingUnicodeV3_2:
		return "KTextEncodingUnicodeV3_2"
	case KTextEncodingUnicodeV4_0:
		return "KTextEncodingUnicodeV4_0"
	case KTextEncodingUnicodeV5_0:
		return "KTextEncodingUnicodeV5_0"
	case KTextEncodingUnicodeV5_1:
		return "KTextEncodingUnicodeV5_1"
	case KTextEncodingUnicodeV6_0:
		return "KTextEncodingUnicodeV6_0"
	case KTextEncodingUnicodeV6_1:
		return "KTextEncodingUnicodeV6_1"
	case KTextEncodingUnicodeV6_3:
		return "KTextEncodingUnicodeV6_3"
	case KTextEncodingUnicodeV7_0:
		return "KTextEncodingUnicodeV7_0"
	case KTextEncodingUnicodeV8_0:
		return "KTextEncodingUnicodeV8_0"
	case KTextEncodingUnicodeV9_0:
		return "KTextEncodingUnicodeV9_0"
	case KTextEncodingUnknown:
		return "KTextEncodingUnknown"
	case KTextEncodingVISCII:
		return "KTextEncodingVISCII"
	case KTextEncodingVariantName:
		return "KTextEncodingVariantName"
	case KTextEncodingWindowsANSI:
		return "KTextEncodingWindowsANSI"
	case KTextEncodingWindowsArabic:
		return "KTextEncodingWindowsArabic"
	case KTextEncodingWindowsBalticRim:
		return "KTextEncodingWindowsBalticRim"
	case KTextEncodingWindowsCyrillic:
		return "KTextEncodingWindowsCyrillic"
	case KTextEncodingWindowsGreek:
		return "KTextEncodingWindowsGreek"
	case KTextEncodingWindowsHebrew:
		return "KTextEncodingWindowsHebrew"
	case KTextEncodingWindowsKoreanJohab:
		return "KTextEncodingWindowsKoreanJohab"
	case KTextEncodingWindowsLatin2:
		return "KTextEncodingWindowsLatin2"
	case KTextEncodingWindowsLatin5:
		return "KTextEncodingWindowsLatin5"
	case KTextEncodingWindowsVietnamese:
		return "KTextEncodingWindowsVietnamese"
	default:
		return fmt.Sprintf("KTextEncoding(%d)", e)
	}
}

type KTextEncodingDefault uint

const (
	// KTextEncodingDefaultVariant: The standard default variant for any base encoding.
	KTextEncodingDefaultVariant KTextEncodingDefault = 0
)

func (e KTextEncodingDefault) String() string {
	switch e {
	case KTextEncodingDefaultVariant:
		return "KTextEncodingDefaultVariant"
	default:
		return fmt.Sprintf("KTextEncodingDefault(%d)", e)
	}
}

type KTextEncodingDefaultFormat uint

const (
	// KTextEncodingDefaultFormatValue: The standard default format for any base encoding.
	KTextEncodingDefaultFormatValue KTextEncodingDefaultFormat = 0
	// KUnicode16BitFormat: The 16-bit character encoding format specified by the Unicode standard, equivalent to the UCS-2 format for ISO 10646.
	KUnicode16BitFormat KTextEncodingDefaultFormat = 0
	// KUnicode32BitFormat: The UCS-4 32-bit format defined for ISO 10646.
	KUnicode32BitFormat   KTextEncodingDefaultFormat = 3
	KUnicodeSCSUFormat    KTextEncodingDefaultFormat = 8
	KUnicodeUTF16BEFormat KTextEncodingDefaultFormat = 4
	KUnicodeUTF16Format   KTextEncodingDefaultFormat = 0
	KUnicodeUTF16LEFormat KTextEncodingDefaultFormat = 5
	KUnicodeUTF32BEFormat KTextEncodingDefaultFormat = 6
	KUnicodeUTF32Format   KTextEncodingDefaultFormat = 3
	KUnicodeUTF32LEFormat KTextEncodingDefaultFormat = 7
	// KUnicodeUTF7Format: The Unicode transformation format in which characters encodings are represented by a sequence of 7-bit values.
	KUnicodeUTF7Format KTextEncodingDefaultFormat = 1
	// KUnicodeUTF8Format: The Unicode transformation format in which characters are represented by a sequence of 8-bit values.
	KUnicodeUTF8Format KTextEncodingDefaultFormat = 2
)

func (e KTextEncodingDefaultFormat) String() string {
	switch e {
	case KTextEncodingDefaultFormatValue:
		return "KTextEncodingDefaultFormatValue"
	case KUnicode32BitFormat:
		return "KUnicode32BitFormat"
	case KUnicodeSCSUFormat:
		return "KUnicodeSCSUFormat"
	case KUnicodeUTF16BEFormat:
		return "KUnicodeUTF16BEFormat"
	case KUnicodeUTF16LEFormat:
		return "KUnicodeUTF16LEFormat"
	case KUnicodeUTF32BEFormat:
		return "KUnicodeUTF32BEFormat"
	case KUnicodeUTF32LEFormat:
		return "KUnicodeUTF32LEFormat"
	case KUnicodeUTF7Format:
		return "KUnicodeUTF7Format"
	case KUnicodeUTF8Format:
		return "KUnicodeUTF8Format"
	default:
		return fmt.Sprintf("KTextEncodingDefaultFormat(%d)", e)
	}
}

type KTextEncodingEBCDIC uint

const (
	// KTextEncodingEBCDIC_CP037: Code page 037, extended EBCDIC-US Latin1.
	KTextEncodingEBCDIC_CP037     KTextEncodingEBCDIC = 3074
	KTextEncodingEBCDIC_LatinCore KTextEncodingEBCDIC = 3073
	// KTextEncodingEBCDIC_US: Basic EBCDIC-US encoding.
	KTextEncodingEBCDIC_US KTextEncodingEBCDIC = 3073
)

func (e KTextEncodingEBCDIC) String() string {
	switch e {
	case KTextEncodingEBCDIC_CP037:
		return "KTextEncodingEBCDIC_CP037"
	case KTextEncodingEBCDIC_LatinCore:
		return "KTextEncodingEBCDIC_LatinCore"
	default:
		return fmt.Sprintf("KTextEncodingEBCDIC(%d)", e)
	}
}

type KTextEncodingEUC uint

const (
	// KTextEncodingEUC_CN: ISO 646, GB 2312-80.
	KTextEncodingEUC_CN KTextEncodingEUC = 2352
	// KTextEncodingEUC_JP: ISO 646,1-byte katakana,JIS 208 ,JIS 212.
	KTextEncodingEUC_JP KTextEncodingEUC = 2336
	// KTextEncodingEUC_KR: ISO 646, KS C 5601-1987.
	KTextEncodingEUC_KR KTextEncodingEUC = 2368
	// KTextEncodingEUC_TW: ISO 646, CNS 11643-1992 Planes 1-16.
	KTextEncodingEUC_TW KTextEncodingEUC = 2353
)

func (e KTextEncodingEUC) String() string {
	switch e {
	case KTextEncodingEUC_CN:
		return "KTextEncodingEUC_CN"
	case KTextEncodingEUC_JP:
		return "KTextEncodingEUC_JP"
	case KTextEncodingEUC_KR:
		return "KTextEncodingEUC_KR"
	case KTextEncodingEUC_TW:
		return "KTextEncodingEUC_TW"
	default:
		return fmt.Sprintf("KTextEncodingEUC(%d)", e)
	}
}

type KTextEncodingISO uint

const (
	// KTextEncodingISOLatin1: ISO 8859-1.
	KTextEncodingISOLatin1  KTextEncodingISO = 513
	KTextEncodingISOLatin10 KTextEncodingISO = 528
	// KTextEncodingISOLatin2: ISO 8859-2.
	KTextEncodingISOLatin2 KTextEncodingISO = 514
	// KTextEncodingISOLatin3: ISO 8859-3.
	KTextEncodingISOLatin3 KTextEncodingISO = 515
	// KTextEncodingISOLatin4: ISO 8859-4.
	KTextEncodingISOLatin4 KTextEncodingISO = 516
	// KTextEncodingISOLatin5: ISO 8859-9.
	KTextEncodingISOLatin5 KTextEncodingISO = 521
	// KTextEncodingISOLatin6: ISO 8859-10.
	KTextEncodingISOLatin6 KTextEncodingISO = 522
	// KTextEncodingISOLatin7: ISO 8859-13; Baltic Rim
	KTextEncodingISOLatin7 KTextEncodingISO = 525
	// KTextEncodingISOLatin8: ISO 8859-14; Celtic
	KTextEncodingISOLatin8 KTextEncodingISO = 526
	// KTextEncodingISOLatin9: ISO 8859-15, 8859-1; changed for Euro & CP1252 letters
	KTextEncodingISOLatin9 KTextEncodingISO = 527
	// KTextEncodingISOLatinArabic: ISO 8859-6; equivalent to ASMO 708 and DOS CP 708.
	KTextEncodingISOLatinArabic KTextEncodingISO = 518
	// KTextEncodingISOLatinCyrillic: ISO 8859-5.
	KTextEncodingISOLatinCyrillic KTextEncodingISO = 517
	// KTextEncodingISOLatinGreek: ISO 8859-7.
	KTextEncodingISOLatinGreek KTextEncodingISO = 519
	// KTextEncodingISOLatinHebrew: ISO 8859-8.
	KTextEncodingISOLatinHebrew KTextEncodingISO = 520
)

func (e KTextEncodingISO) String() string {
	switch e {
	case KTextEncodingISOLatin1:
		return "KTextEncodingISOLatin1"
	case KTextEncodingISOLatin10:
		return "KTextEncodingISOLatin10"
	case KTextEncodingISOLatin2:
		return "KTextEncodingISOLatin2"
	case KTextEncodingISOLatin3:
		return "KTextEncodingISOLatin3"
	case KTextEncodingISOLatin4:
		return "KTextEncodingISOLatin4"
	case KTextEncodingISOLatin5:
		return "KTextEncodingISOLatin5"
	case KTextEncodingISOLatin6:
		return "KTextEncodingISOLatin6"
	case KTextEncodingISOLatin7:
		return "KTextEncodingISOLatin7"
	case KTextEncodingISOLatin8:
		return "KTextEncodingISOLatin8"
	case KTextEncodingISOLatin9:
		return "KTextEncodingISOLatin9"
	case KTextEncodingISOLatinArabic:
		return "KTextEncodingISOLatinArabic"
	case KTextEncodingISOLatinCyrillic:
		return "KTextEncodingISOLatinCyrillic"
	case KTextEncodingISOLatinGreek:
		return "KTextEncodingISOLatinGreek"
	case KTextEncodingISOLatinHebrew:
		return "KTextEncodingISOLatinHebrew"
	default:
		return fmt.Sprintf("KTextEncodingISO(%d)", e)
	}
}

type KTextEncodingMac uint

const (
	// KTextEncodingMacArabic: The encoding for Mac OS Arabic.
	KTextEncodingMacArabic KTextEncodingMac = 4
	// KTextEncodingMacArmenian: The encoding for Mac OS Armenian.
	KTextEncodingMacArmenian KTextEncodingMac = 24
	// KTextEncodingMacBengali: The encoding for Mac OS Bengali.
	KTextEncodingMacBengali KTextEncodingMac = 13
	// KTextEncodingMacBurmese: The encoding for Mac OS Burmese.
	KTextEncodingMacBurmese KTextEncodingMac = 19
	// KTextEncodingMacCeltic: This Mac OS encoding uses script code 0, `smRoman`.
	KTextEncodingMacCeltic KTextEncodingMac = 39
	// KTextEncodingMacCentralEurRoman: The encoding for Mac OS Central European Roman.
	KTextEncodingMacCentralEurRoman KTextEncodingMac = 29
	// KTextEncodingMacChineseSimp: The encoding for Mac OS simple Chinese.
	KTextEncodingMacChineseSimp KTextEncodingMac = 25
	// KTextEncodingMacChineseTrad: The encoding for Mac OS traditional Chinese.
	KTextEncodingMacChineseTrad KTextEncodingMac = 2
	// KTextEncodingMacCroatian: This Mac OS encoding uses script code 0, `smRoman`.
	KTextEncodingMacCroatian KTextEncodingMac = 36
	// KTextEncodingMacCyrillic: The encoding for Mac OS Cyrillic.
	KTextEncodingMacCyrillic KTextEncodingMac = 7
	// KTextEncodingMacDevanagari: The encoding for Mac OS Devanagari.
	KTextEncodingMacDevanagari KTextEncodingMac = 9
	// KTextEncodingMacDingbats: This Mac OS encoding uses script code 0, `smRoman`.
	KTextEncodingMacDingbats     KTextEncodingMac = 34
	KTextEncodingMacEastEurRoman KTextEncodingMac = 29
	// KTextEncodingMacEthiopic: The encoding for Mac OS Ethiopic.
	KTextEncodingMacEthiopic KTextEncodingMac = 28
	// KTextEncodingMacExtArabic: The encoding for Mac OS ExtArabic.
	KTextEncodingMacExtArabic KTextEncodingMac = 31
	// KTextEncodingMacFarsi: Uses script code 4, `smArabic`.
	KTextEncodingMacFarsi KTextEncodingMac = 140
	// KTextEncodingMacGaelic: This Mac OS encoding uses script code 0, `smRoman`.
	KTextEncodingMacGaelic KTextEncodingMac = 40
	KTextEncodingMacGeez   KTextEncodingMac = 28
	// KTextEncodingMacGeorgian: The encoding for Mac OS Georgian.
	KTextEncodingMacGeorgian KTextEncodingMac = 23
	// KTextEncodingMacGreek: The encoding for Mac OS Greek.
	KTextEncodingMacGreek KTextEncodingMac = 6
	// KTextEncodingMacGujarati: The encoding for Mac OS Gujurati.
	KTextEncodingMacGujarati KTextEncodingMac = 11
	// KTextEncodingMacGurmukhi: The encoding for Mac OS Gurmukhi.
	KTextEncodingMacGurmukhi KTextEncodingMac = 10
	// KTextEncodingMacHebrew: The encoding for Mac OS Hebrew.
	KTextEncodingMacHebrew KTextEncodingMac = 5
	// KTextEncodingMacIcelandic: This Mac OS encoding uses script code 0, `smRoman`.
	KTextEncodingMacIcelandic KTextEncodingMac = 37
	// KTextEncodingMacInuit: Uses script code 28, `smEthiopic`.]
	KTextEncodingMacInuit KTextEncodingMac = 236
	// KTextEncodingMacJapanese: The encoding for Mac OS Japanese.
	KTextEncodingMacJapanese KTextEncodingMac = 1
	// KTextEncodingMacKannada: The encoding for Mac OS Kannada.
	KTextEncodingMacKannada        KTextEncodingMac = 16
	KTextEncodingMacKeyboardGlyphs KTextEncodingMac = 41
	// KTextEncodingMacKhmer: The encoding for Mac OS Khmer.
	KTextEncodingMacKhmer KTextEncodingMac = 20
	// KTextEncodingMacKorean: The encoding for Mac OS Korean.
	KTextEncodingMacKorean KTextEncodingMac = 3
	// KTextEncodingMacLaotian: The encoding for Mac OS Laotian.
	KTextEncodingMacLaotian KTextEncodingMac = 22
	// KTextEncodingMacMalayalam: The encoding for Mac OS Malayalam.
	KTextEncodingMacMalayalam KTextEncodingMac = 17
	// KTextEncodingMacMongolian: The encoding for Mac OS Mongolian.
	KTextEncodingMacMongolian KTextEncodingMac = 27
	// KTextEncodingMacOriya: The encoding for Mac OS Oriya.
	KTextEncodingMacOriya   KTextEncodingMac = 12
	KTextEncodingMacRSymbol KTextEncodingMac = 8
	// KTextEncodingMacRoman: The encoding for Mac OS Roman.
	KTextEncodingMacRoman KTextEncodingMac = 0
	// KTextEncodingMacRomanian: This Mac OS encoding uses script code 0, `smRoman`.
	KTextEncodingMacRomanian    KTextEncodingMac = 38
	KTextEncodingMacSimpChinese KTextEncodingMac = 25
	// KTextEncodingMacSinhalese: The encoding for Mac OS Sinhalese.
	KTextEncodingMacSinhalese KTextEncodingMac = 18
	// KTextEncodingMacSymbol: This Mac OS encoding uses script code 0, `smRoman`.
	KTextEncodingMacSymbol KTextEncodingMac = 33
	// KTextEncodingMacTamil: The encoding for Mac OS Tamil.
	KTextEncodingMacTamil KTextEncodingMac = 14
	// KTextEncodingMacTelugu: The encoding for Mac OS Telugu.
	KTextEncodingMacTelugu KTextEncodingMac = 15
	// KTextEncodingMacThai: The encoding for Mac OS Thai.
	KTextEncodingMacThai KTextEncodingMac = 21
	// KTextEncodingMacTibetan: The encoding for Mac OS Tibetan.
	KTextEncodingMacTibetan     KTextEncodingMac = 26
	KTextEncodingMacTradChinese KTextEncodingMac = 2
	// KTextEncodingMacTurkish: This Mac OS encoding uses script code 0, `smRoman`.
	KTextEncodingMacTurkish KTextEncodingMac = 35
	// KTextEncodingMacUkrainian: Uses script code 7, `smCyrillic`.]
	KTextEncodingMacUkrainian KTextEncodingMac = 152
	// KTextEncodingMacUnicode: # Discussion
	KTextEncodingMacUnicode  KTextEncodingMac = 126
	KTextEncodingMacUninterp KTextEncodingMac = 32
	// KTextEncodingMacVT100: Uses script code 32, `smUninterp`; VT100/102 font from the common toolbox; Latin-1 characters plus box drawing.]
	KTextEncodingMacVT100 KTextEncodingMac = 252
	// KTextEncodingMacVietnamese: The encoding for Mac OS Vietnamese.
	KTextEncodingMacVietnamese KTextEncodingMac = 30
)

func (e KTextEncodingMac) String() string {
	switch e {
	case KTextEncodingMacArabic:
		return "KTextEncodingMacArabic"
	case KTextEncodingMacArmenian:
		return "KTextEncodingMacArmenian"
	case KTextEncodingMacBengali:
		return "KTextEncodingMacBengali"
	case KTextEncodingMacBurmese:
		return "KTextEncodingMacBurmese"
	case KTextEncodingMacCeltic:
		return "KTextEncodingMacCeltic"
	case KTextEncodingMacCentralEurRoman:
		return "KTextEncodingMacCentralEurRoman"
	case KTextEncodingMacChineseSimp:
		return "KTextEncodingMacChineseSimp"
	case KTextEncodingMacChineseTrad:
		return "KTextEncodingMacChineseTrad"
	case KTextEncodingMacCroatian:
		return "KTextEncodingMacCroatian"
	case KTextEncodingMacCyrillic:
		return "KTextEncodingMacCyrillic"
	case KTextEncodingMacDevanagari:
		return "KTextEncodingMacDevanagari"
	case KTextEncodingMacDingbats:
		return "KTextEncodingMacDingbats"
	case KTextEncodingMacEthiopic:
		return "KTextEncodingMacEthiopic"
	case KTextEncodingMacExtArabic:
		return "KTextEncodingMacExtArabic"
	case KTextEncodingMacFarsi:
		return "KTextEncodingMacFarsi"
	case KTextEncodingMacGaelic:
		return "KTextEncodingMacGaelic"
	case KTextEncodingMacGeorgian:
		return "KTextEncodingMacGeorgian"
	case KTextEncodingMacGreek:
		return "KTextEncodingMacGreek"
	case KTextEncodingMacGujarati:
		return "KTextEncodingMacGujarati"
	case KTextEncodingMacGurmukhi:
		return "KTextEncodingMacGurmukhi"
	case KTextEncodingMacHebrew:
		return "KTextEncodingMacHebrew"
	case KTextEncodingMacIcelandic:
		return "KTextEncodingMacIcelandic"
	case KTextEncodingMacInuit:
		return "KTextEncodingMacInuit"
	case KTextEncodingMacJapanese:
		return "KTextEncodingMacJapanese"
	case KTextEncodingMacKannada:
		return "KTextEncodingMacKannada"
	case KTextEncodingMacKeyboardGlyphs:
		return "KTextEncodingMacKeyboardGlyphs"
	case KTextEncodingMacKhmer:
		return "KTextEncodingMacKhmer"
	case KTextEncodingMacKorean:
		return "KTextEncodingMacKorean"
	case KTextEncodingMacLaotian:
		return "KTextEncodingMacLaotian"
	case KTextEncodingMacMalayalam:
		return "KTextEncodingMacMalayalam"
	case KTextEncodingMacMongolian:
		return "KTextEncodingMacMongolian"
	case KTextEncodingMacOriya:
		return "KTextEncodingMacOriya"
	case KTextEncodingMacRSymbol:
		return "KTextEncodingMacRSymbol"
	case KTextEncodingMacRoman:
		return "KTextEncodingMacRoman"
	case KTextEncodingMacRomanian:
		return "KTextEncodingMacRomanian"
	case KTextEncodingMacSinhalese:
		return "KTextEncodingMacSinhalese"
	case KTextEncodingMacSymbol:
		return "KTextEncodingMacSymbol"
	case KTextEncodingMacTamil:
		return "KTextEncodingMacTamil"
	case KTextEncodingMacTelugu:
		return "KTextEncodingMacTelugu"
	case KTextEncodingMacThai:
		return "KTextEncodingMacThai"
	case KTextEncodingMacTibetan:
		return "KTextEncodingMacTibetan"
	case KTextEncodingMacTurkish:
		return "KTextEncodingMacTurkish"
	case KTextEncodingMacUkrainian:
		return "KTextEncodingMacUkrainian"
	case KTextEncodingMacUnicode:
		return "KTextEncodingMacUnicode"
	case KTextEncodingMacUninterp:
		return "KTextEncodingMacUninterp"
	case KTextEncodingMacVT100:
		return "KTextEncodingMacVT100"
	case KTextEncodingMacVietnamese:
		return "KTextEncodingMacVietnamese"
	default:
		return fmt.Sprintf("KTextEncodingMac(%d)", e)
	}
}

type KTextEncodingMacHF uint

const (
	// KTextEncodingMacHFS: This is a metavalue for a special Mac OS encoding.
	KTextEncodingMacHFS KTextEncodingMacHF = 255
)

func (e KTextEncodingMacHF) String() string {
	switch e {
	case KTextEncodingMacHFS:
		return "KTextEncodingMacHFS"
	default:
		return fmt.Sprintf("KTextEncodingMacHF(%d)", e)
	}
}

type KTextEncodingNextStep uint

const (
	KTextEncodingNextStepJapanese KTextEncodingNextStep = 2818
	KTextEncodingNextStepLatin    KTextEncodingNextStep = 2817
)

func (e KTextEncodingNextStep) String() string {
	switch e {
	case KTextEncodingNextStepJapanese:
		return "KTextEncodingNextStepJapanese"
	case KTextEncodingNextStepLatin:
		return "KTextEncodingNextStepLatin"
	default:
		return fmt.Sprintf("KTextEncodingNextStep(%d)", e)
	}
}

type KTextServiceClass uint

const (
	KGetSelectedText            KTextServiceClass = 'g'<<24 | 't'<<16 | 'x'<<8 | 't' // 'gtxt'
	KOffset2Pos                 KTextServiceClass = 's'<<24 | 't'<<16 | '2'<<8 | 'p' // 'st2p'
	KPos2Offset                 KTextServiceClass = 'p'<<24 | '2'<<16 | 's'<<8 | 't' // 'p2st'
	KShowHideInputWindow        KTextServiceClass = 's'<<24 | 'h'<<16 | 'i'<<8 | 'w' // 'shiw'
	KTextServiceClassValue      KTextServiceClass = 't'<<24 | 's'<<16 | 'v'<<8 | 'c' // 'tsvc'
	KUnicodeNotFromInputMethod  KTextServiceClass = 'u'<<24 | 'n'<<16 | 'i'<<8 | 'm' // 'unim'
	KUpdateActiveInputArea      KTextServiceClass = 'u'<<24 | 'p'<<16 | 'd'<<8 | 't' // 'updt'
	KeyAEBufferSize             KTextServiceClass = 'b'<<24 | 'u'<<16 | 'f'<<8 | 'f' // 'buff'
	KeyAECurrentPoint           KTextServiceClass = 'c'<<24 | 'p'<<16 | 'o'<<8 | 's' // 'cpos'
	KeyAEFixLength              KTextServiceClass = 'f'<<24 | 'i'<<16 | 'x'<<8 | 'l' // 'fixl'
	KeyAEMoveView               KTextServiceClass = 'm'<<24 | 'v'<<16 | 'v'<<8 | 'w' // 'mvvw'
	KeyAENextBody               KTextServiceClass = 'n'<<24 | 'x'<<16 | 'b'<<8 | 'd' // 'nxbd'
	KeyAEServerInstance         KTextServiceClass = 's'<<24 | 'r'<<16 | 'v'<<8 | 'i' // 'srvi'
	KeyAETSMDocumentRefcon      KTextServiceClass = 'r'<<24 | 'e'<<16 | 'f'<<8 | 'c' // 'refc'
	KeyAETSMEventRecord         KTextServiceClass = 't'<<24 | 'e'<<16 | 'v'<<8 | 't' // 'tevt'
	KeyAETSMEventRef            KTextServiceClass = 't'<<24 | 'e'<<16 | 'v'<<8 | 'r' // 'tevr'
	KeyAETSMGlyphInfoArray      KTextServiceClass = 't'<<24 | 'g'<<16 | 'i'<<8 | 'a' // 'tgia'
	KeyAETSMScriptTag           KTextServiceClass = 's'<<24 | 'c'<<16 | 'l'<<8 | 'g' // 'sclg'
	KeyAETSMTextFMFont          KTextServiceClass = 'k'<<24 | 't'<<16 | 'x'<<8 | 'm' // 'ktxm'
	KeyAETSMTextFont            KTextServiceClass = 'k'<<24 | 't'<<16 | 'x'<<8 | 'f' // 'ktxf'
	KeyAETSMTextPointSize       KTextServiceClass = 'k'<<24 | 't'<<16 | 'p'<<8 | 's' // 'ktps'
	KeyAETextServiceEncoding    KTextServiceClass = 't'<<24 | 's'<<16 | 'e'<<8 | 'n' // 'tsen'
	KeyAETextServiceMacEncoding KTextServiceClass = 't'<<24 | 'm'<<16 | 'e'<<8 | 'n' // 'tmen'
	KeyAETheData                KTextServiceClass = 'k'<<24 | 'd'<<16 | 'a'<<8 | 't' // 'kdat'
	KeyAEUpdateRange            KTextServiceClass = 'u'<<24 | 'd'<<16 | 'n'<<8 | 'g' // 'udng'
	TypeComponentInstance       KTextServiceClass = 'c'<<24 | 'm'<<16 | 'p'<<8 | 'i' // 'cmpi'
	TypeEventRef                KTextServiceClass = 'e'<<24 | 'v'<<16 | 'r'<<8 | 'f' // 'evrf'
	TypeGlyphInfoArray          KTextServiceClass = 'g'<<24 | 'l'<<16 | 'i'<<8 | 'a' // 'glia'
	TypeLowLevelEventRecord     KTextServiceClass = 'e'<<24 | 'v'<<16 | 't'<<8 | 'r' // 'evtr'
	TypeOffsetArray             KTextServiceClass = 'o'<<24 | 'f'<<16 | 'a'<<8 | 'y' // 'ofay'
	TypeText                    KTextServiceClass = 'T'<<24 | 'E'<<16 | 'X'<<8 | 'T' // 'TEXT'
	TypeTextRange               KTextServiceClass = 't'<<24 | 'x'<<16 | 'r'<<8 | 'n' // 'txrn'
	TypeTextRangeArray          KTextServiceClass = 't'<<24 | 'r'<<16 | 'a'<<8 | 'y' // 'tray'
)

func (e KTextServiceClass) String() string {
	switch e {
	case KGetSelectedText:
		return "KGetSelectedText"
	case KOffset2Pos:
		return "KOffset2Pos"
	case KPos2Offset:
		return "KPos2Offset"
	case KShowHideInputWindow:
		return "KShowHideInputWindow"
	case KTextServiceClassValue:
		return "KTextServiceClassValue"
	case KUnicodeNotFromInputMethod:
		return "KUnicodeNotFromInputMethod"
	case KUpdateActiveInputArea:
		return "KUpdateActiveInputArea"
	case KeyAEBufferSize:
		return "KeyAEBufferSize"
	case KeyAECurrentPoint:
		return "KeyAECurrentPoint"
	case KeyAEFixLength:
		return "KeyAEFixLength"
	case KeyAEMoveView:
		return "KeyAEMoveView"
	case KeyAENextBody:
		return "KeyAENextBody"
	case KeyAEServerInstance:
		return "KeyAEServerInstance"
	case KeyAETSMDocumentRefcon:
		return "KeyAETSMDocumentRefcon"
	case KeyAETSMEventRecord:
		return "KeyAETSMEventRecord"
	case KeyAETSMEventRef:
		return "KeyAETSMEventRef"
	case KeyAETSMGlyphInfoArray:
		return "KeyAETSMGlyphInfoArray"
	case KeyAETSMScriptTag:
		return "KeyAETSMScriptTag"
	case KeyAETSMTextFMFont:
		return "KeyAETSMTextFMFont"
	case KeyAETSMTextFont:
		return "KeyAETSMTextFont"
	case KeyAETSMTextPointSize:
		return "KeyAETSMTextPointSize"
	case KeyAETextServiceEncoding:
		return "KeyAETextServiceEncoding"
	case KeyAETextServiceMacEncoding:
		return "KeyAETextServiceMacEncoding"
	case KeyAETheData:
		return "KeyAETheData"
	case KeyAEUpdateRange:
		return "KeyAEUpdateRange"
	case TypeComponentInstance:
		return "TypeComponentInstance"
	case TypeEventRef:
		return "TypeEventRef"
	case TypeGlyphInfoArray:
		return "TypeGlyphInfoArray"
	case TypeLowLevelEventRecord:
		return "TypeLowLevelEventRecord"
	case TypeOffsetArray:
		return "TypeOffsetArray"
	case TypeText:
		return "TypeText"
	case TypeTextRange:
		return "TypeTextRange"
	case TypeTextRangeArray:
		return "TypeTextRangeArray"
	default:
		return fmt.Sprintf("KTextServiceClass(%d)", e)
	}
}

type KTextUnsupportedEncodingErr int

const (
	// KTECArrayFullErr: # Discussion
	KTECArrayFullErr  KTextUnsupportedEncodingErr = -8751
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
	UnicodeBufErr                    KTextUnsupportedEncodingErr = -8764
	UnicodeCharErr                   KTextUnsupportedEncodingErr = -8763
	UnicodeChecksumErr               KTextUnsupportedEncodingErr = -8769
	UnicodeContextualErr             KTextUnsupportedEncodingErr = -8758
	UnicodeDirectionErr              KTextUnsupportedEncodingErr = -8759
	UnicodeElementErr                KTextUnsupportedEncodingErr = -8762
	UnicodeFallbacksErr              KTextUnsupportedEncodingErr = -8766
	UnicodeNoTableErr                KTextUnsupportedEncodingErr = -8768
	UnicodeNotFoundErr               KTextUnsupportedEncodingErr = -8761
	UnicodePartConvertErr            KTextUnsupportedEncodingErr = -8765
	UnicodeTableFormatErr            KTextUnsupportedEncodingErr = -8760
	UnicodeTextEncodingDataErr       KTextUnsupportedEncodingErr = -8757
	UnicodeVariantErr                KTextUnsupportedEncodingErr = -8767
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

type KTileIconVariant uint

const (
	KDropIconVariant      KTileIconVariant = 'd'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'drop'
	KOpenDropIconVariant  KTileIconVariant = 'o'<<24 | 'd'<<16 | 'r'<<8 | 'p' // 'odrp'
	KOpenIconVariant      KTileIconVariant = 'o'<<24 | 'p'<<16 | 'e'<<8 | 'n' // 'open'
	KRolloverIconVariant  KTileIconVariant = 'o'<<24 | 'v'<<16 | 'e'<<8 | 'r' // 'over'
	KTileIconVariantValue KTileIconVariant = 't'<<24 | 'i'<<16 | 'l'<<8 | 'e' // 'tile'
)

func (e KTileIconVariant) String() string {
	switch e {
	case KDropIconVariant:
		return "KDropIconVariant"
	case KOpenDropIconVariant:
		return "KOpenDropIconVariant"
	case KOpenIconVariant:
		return "KOpenIconVariant"
	case KRolloverIconVariant:
		return "KRolloverIconVariant"
	case KTileIconVariantValue:
		return "KTileIconVariantValue"
	default:
		return fmt.Sprintf("KTileIconVariant(%d)", e)
	}
}

type KToolbar uint

const (
	KToolbarAdvancedIcon           KToolbar = 't'<<24 | 'b'<<16 | 'a'<<8 | 'v' // 'tbav'
	KToolbarApplicationsFolderIcon KToolbar = 't'<<24 | 'A'<<16 | 'p'<<8 | 's' // 'tAps'
	KToolbarCustomizeIcon          KToolbar = 't'<<24 | 'c'<<16 | 'u'<<8 | 's' // 'tcus'
	KToolbarDeleteIcon             KToolbar = 't'<<24 | 'd'<<16 | 'e'<<8 | 'l' // 'tdel'
	KToolbarDesktopFolderIcon      KToolbar = 't'<<24 | 'D'<<16 | 's'<<8 | 'k' // 'tDsk'
	KToolbarDocumentsFolderIcon    KToolbar = 't'<<24 | 'D'<<16 | 'o'<<8 | 'c' // 'tDoc'
	KToolbarDownloadsFolderIcon    KToolbar = 't'<<24 | 'D'<<16 | 'w'<<8 | 'n' // 'tDwn'
	KToolbarFavoritesIcon          KToolbar = 't'<<24 | 'f'<<16 | 'a'<<8 | 'v' // 'tfav'
	KToolbarHomeIcon               KToolbar = 't'<<24 | 'h'<<16 | 'o'<<8 | 'm' // 'thom'
	KToolbarInfoIcon               KToolbar = 't'<<24 | 'b'<<16 | 'i'<<8 | 'n' // 'tbin'
	KToolbarLabelsIcon             KToolbar = 't'<<24 | 'b'<<16 | 'l'<<8 | 'b' // 'tblb'
	KToolbarLibraryFolderIcon      KToolbar = 't'<<24 | 'L'<<16 | 'i'<<8 | 'b' // 'tLib'
	KToolbarMovieFolderIcon        KToolbar = 't'<<24 | 'M'<<16 | 'o'<<8 | 'v' // 'tMov'
	KToolbarMusicFolderIcon        KToolbar = 't'<<24 | 'M'<<16 | 'u'<<8 | 's' // 'tMus'
	KToolbarPicturesFolderIcon     KToolbar = 't'<<24 | 'P'<<16 | 'i'<<8 | 'c' // 'tPic'
	KToolbarPublicFolderIcon       KToolbar = 't'<<24 | 'P'<<16 | 'u'<<8 | 'b' // 'tPub'
	KToolbarSitesFolderIcon        KToolbar = 't'<<24 | 'S'<<16 | 't'<<8 | 's' // 'tSts'
	KToolbarUtilitiesFolderIcon    KToolbar = 't'<<24 | 'U'<<16 | 't'<<8 | 'l' // 'tUtl'
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
	KUCHighSurrogateRangeEnd       KUC = 56319
	KUCHighSurrogateRangeStart     KUC = 55296
	KUCLowSurrogateRangeEnd        KUC = 57343
	KUCLowSurrogateRangeStart      KUC = 56320
	KUCOutputBufferTooSmall        KUC = -25340
	KUCTextBreakLocatorMissingType KUC = -25341
	KUCTokenNotFound               KUC = -25346
	KUCTokenizerIterationFinished  KUC = -25344
	KUCTokenizerUnknownLang        KUC = -25345
)

func (e KUC) String() string {
	switch e {
	case KUCHighSurrogateRangeEnd:
		return "KUCHighSurrogateRangeEnd"
	case KUCHighSurrogateRangeStart:
		return "KUCHighSurrogateRangeStart"
	case KUCLowSurrogateRangeEnd:
		return "KUCLowSurrogateRangeEnd"
	case KUCLowSurrogateRangeStart:
		return "KUCLowSurrogateRangeStart"
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

type KUCBidiCat uint

const (
	// KUCBidiCatArabicNumber: Weak types: AN Arabic number.
	KUCBidiCatArabicNumber KUCBidiCat = 6
	// KUCBidiCatBlockSeparator: Separators: B  paragraph separator (was block separator).
	KUCBidiCatBlockSeparator KUCBidiCat = 8
	// KUCBidiCatBoundaryNeutral: Unicode 3.0; BN boundary neutral.
	KUCBidiCatBoundaryNeutral KUCBidiCat = 19
	// KUCBidiCatCommonNumberSeparator: Weak types: CS common number separator.
	KUCBidiCatCommonNumberSeparator KUCBidiCat = 7
	// KUCBidiCatEuroNumber: Weak types: EN European number.
	KUCBidiCatEuroNumber KUCBidiCat = 3
	// KUCBidiCatEuroNumberSeparator: Weak types: ES European number separator.
	KUCBidiCatEuroNumberSeparator KUCBidiCat = 4
	// KUCBidiCatEuroNumberTerminator: Weak types: ET European number terminator.
	KUCBidiCatEuroNumberTerminator KUCBidiCat = 5
	KUCBidiCatFirstStrongIsolate   KUCBidiCat = 22
	// KUCBidiCatLeftRight: Strong types: L  left-to-right.
	KUCBidiCatLeftRight KUCBidiCat = 1
	// KUCBidiCatLeftRightEmbedding: Unicode 3.0; LRE eft-to-right embedding.
	KUCBidiCatLeftRightEmbedding KUCBidiCat = 13
	KUCBidiCatLeftRightIsolate   KUCBidiCat = 20
	// KUCBidiCatLeftRightOverride: Unicode 3.0; LRO left-to-right override.
	KUCBidiCatLeftRightOverride KUCBidiCat = 15
	// KUCBidiCatNonSpacingMark: Unicode 3.0; NSM non-spacing mark.
	KUCBidiCatNonSpacingMark KUCBidiCat = 18
	// KUCBidiCatNotApplicable: Unassigned.
	KUCBidiCatNotApplicable KUCBidiCat = 0
	// KUCBidiCatOtherNeutral: Neutrals: ON other neutrals (unassigned codes could use this).
	KUCBidiCatOtherNeutral KUCBidiCat = 11
	// KUCBidiCatPopDirectionalFormat: Unicode 3.0; PDF pop directional Format.
	KUCBidiCatPopDirectionalFormat  KUCBidiCat = 17
	KUCBidiCatPopDirectionalIsolate KUCBidiCat = 23
	// KUCBidiCatRightLeft: Strong types: R  right-to-left.
	KUCBidiCatRightLeft KUCBidiCat = 2
	// KUCBidiCatRightLeftArabic: Unicode 3.0; AL right-to-left Arabic (was Arabic letter).
	KUCBidiCatRightLeftArabic KUCBidiCat = 12
	// KUCBidiCatRightLeftEmbedding: Unicode 3.0; RLE right-to-left embedding.
	KUCBidiCatRightLeftEmbedding KUCBidiCat = 14
	KUCBidiCatRightLeftIsolate   KUCBidiCat = 21
	// KUCBidiCatRightLeftOverride: Unicode 3.0; RLO right-to-left override.
	KUCBidiCatRightLeftOverride KUCBidiCat = 16
	// KUCBidiCatSegmentSeparator: Separators: S  segment separator.
	KUCBidiCatSegmentSeparator KUCBidiCat = 9
	// KUCBidiCatWhitespace: Neutrals: WS whitespace.
	KUCBidiCatWhitespace KUCBidiCat = 10
)

func (e KUCBidiCat) String() string {
	switch e {
	case KUCBidiCatArabicNumber:
		return "KUCBidiCatArabicNumber"
	case KUCBidiCatBlockSeparator:
		return "KUCBidiCatBlockSeparator"
	case KUCBidiCatBoundaryNeutral:
		return "KUCBidiCatBoundaryNeutral"
	case KUCBidiCatCommonNumberSeparator:
		return "KUCBidiCatCommonNumberSeparator"
	case KUCBidiCatEuroNumber:
		return "KUCBidiCatEuroNumber"
	case KUCBidiCatEuroNumberSeparator:
		return "KUCBidiCatEuroNumberSeparator"
	case KUCBidiCatEuroNumberTerminator:
		return "KUCBidiCatEuroNumberTerminator"
	case KUCBidiCatFirstStrongIsolate:
		return "KUCBidiCatFirstStrongIsolate"
	case KUCBidiCatLeftRight:
		return "KUCBidiCatLeftRight"
	case KUCBidiCatLeftRightEmbedding:
		return "KUCBidiCatLeftRightEmbedding"
	case KUCBidiCatLeftRightIsolate:
		return "KUCBidiCatLeftRightIsolate"
	case KUCBidiCatLeftRightOverride:
		return "KUCBidiCatLeftRightOverride"
	case KUCBidiCatNonSpacingMark:
		return "KUCBidiCatNonSpacingMark"
	case KUCBidiCatNotApplicable:
		return "KUCBidiCatNotApplicable"
	case KUCBidiCatOtherNeutral:
		return "KUCBidiCatOtherNeutral"
	case KUCBidiCatPopDirectionalFormat:
		return "KUCBidiCatPopDirectionalFormat"
	case KUCBidiCatPopDirectionalIsolate:
		return "KUCBidiCatPopDirectionalIsolate"
	case KUCBidiCatRightLeft:
		return "KUCBidiCatRightLeft"
	case KUCBidiCatRightLeftArabic:
		return "KUCBidiCatRightLeftArabic"
	case KUCBidiCatRightLeftEmbedding:
		return "KUCBidiCatRightLeftEmbedding"
	case KUCBidiCatRightLeftIsolate:
		return "KUCBidiCatRightLeftIsolate"
	case KUCBidiCatRightLeftOverride:
		return "KUCBidiCatRightLeftOverride"
	case KUCBidiCatSegmentSeparator:
		return "KUCBidiCatSegmentSeparator"
	case KUCBidiCatWhitespace:
		return "KUCBidiCatWhitespace"
	default:
		return fmt.Sprintf("KUCBidiCat(%d)", e)
	}
}

type KUCCharPropType uint

const (
	// KUCCharPropTypeBidiCategory: Requests enumeration value.
	KUCCharPropTypeBidiCategory KUCCharPropType = 3
	// KUCCharPropTypeCombiningClass: Requests numeric value 0 to 255.
	KUCCharPropTypeCombiningClass    KUCCharPropType = 2
	KUCCharPropTypeDecimalDigitValue KUCCharPropType = 4
	// KUCCharPropTypeGenlCategory: Requests enumeration value.
	KUCCharPropTypeGenlCategory KUCCharPropType = 1
)

func (e KUCCharPropType) String() string {
	switch e {
	case KUCCharPropTypeBidiCategory:
		return "KUCCharPropTypeBidiCategory"
	case KUCCharPropTypeCombiningClass:
		return "KUCCharPropTypeCombiningClass"
	case KUCCharPropTypeDecimalDigitValue:
		return "KUCCharPropTypeDecimalDigitValue"
	case KUCCharPropTypeGenlCategory:
		return "KUCCharPropTypeGenlCategory"
	default:
		return fmt.Sprintf("KUCCharPropType(%d)", e)
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

type KUCCollateStandard uint

const (
	// KUCCollateStandardOptions: # Discussion
	KUCCollateStandardOptions KUCCollateStandard = 2
)

func (e KUCCollateStandard) String() string {
	switch e {
	case KUCCollateStandardOptions:
		return "KUCCollateStandardOptions"
	default:
		return fmt.Sprintf("KUCCollateStandard(%d)", e)
	}
}

type KUCCollateType uint

const (
	// KUCCollateTypeMask: You can use this mask to directly test bits 24-31 of a [UCCollateOptions] value.
	KUCCollateTypeMask KUCCollateType = 0
	// KUCCollateTypeShiftBits: # Discussion
	KUCCollateTypeShiftBits KUCCollateType = 24
	// KUCCollateTypeSourceMask: You can use this mask, in conjunction with the `kUCCollateTypeShiftBits` constant, to obtain a value identifying a fixed ordering scheme.
	KUCCollateTypeSourceMask KUCCollateType = 255
)

func (e KUCCollateType) String() string {
	switch e {
	case KUCCollateTypeMask:
		return "KUCCollateTypeMask"
	case KUCCollateTypeShiftBits:
		return "KUCCollateTypeShiftBits"
	case KUCCollateTypeSourceMask:
		return "KUCCollateTypeSourceMask"
	default:
		return fmt.Sprintf("KUCCollateType(%d)", e)
	}
}

type KUCCollateTypeHFS uint

const (
	// KUCCollateTypeHFSExtended: # Discussion
	KUCCollateTypeHFSExtended KUCCollateTypeHFS = 1
)

func (e KUCCollateTypeHFS) String() string {
	switch e {
	case KUCCollateTypeHFSExtended:
		return "KUCCollateTypeHFSExtended"
	default:
		return fmt.Sprintf("KUCCollateTypeHFS(%d)", e)
	}
}

type KUCGenlCat uint

const (
	// KUCGenlCatLetterLowercase: Ll Letter; lowercase.
	KUCGenlCatLetterLowercase KUCGenlCat = 15
	// KUCGenlCatLetterModifier: Lm Letter; modifier.
	KUCGenlCatLetterModifier KUCGenlCat = 17
	// KUCGenlCatLetterOther: Lo Letter; other.
	KUCGenlCatLetterOther KUCGenlCat = 18
	// KUCGenlCatLetterTitlecase: Lt Letter; titlecase.
	KUCGenlCatLetterTitlecase KUCGenlCat = 16
	// KUCGenlCatLetterUppercase: Lu Letter; uppercase.
	KUCGenlCatLetterUppercase KUCGenlCat = 14
	// KUCGenlCatMarkEnclosing: Me mark; enclosing.
	KUCGenlCatMarkEnclosing KUCGenlCat = 7
	// KUCGenlCatMarkNonSpacing: Mn mark; non-spacing.
	KUCGenlCatMarkNonSpacing KUCGenlCat = 5
	// KUCGenlCatMarkSpacingCombining: Mc mark; spacing combining.
	KUCGenlCatMarkSpacingCombining KUCGenlCat = 6
	// KUCGenlCatNumberDecimalDigit: Nd number; decimal digit.
	KUCGenlCatNumberDecimalDigit KUCGenlCat = 8
	// KUCGenlCatNumberLetter: Nl number; letter.
	KUCGenlCatNumberLetter KUCGenlCat = 9
	// KUCGenlCatNumberOther: No number; other.
	KUCGenlCatNumberOther KUCGenlCat = 10
	// KUCGenlCatOtherControl: Cc other; control.
	KUCGenlCatOtherControl KUCGenlCat = 1
	// KUCGenlCatOtherFormat: Cf other; format.
	KUCGenlCatOtherFormat KUCGenlCat = 2
	// KUCGenlCatOtherNotAssigned: Cn other; not assigned.
	KUCGenlCatOtherNotAssigned KUCGenlCat = 0
	// KUCGenlCatOtherPrivateUse: Co other; private use.
	KUCGenlCatOtherPrivateUse KUCGenlCat = 4
	// KUCGenlCatOtherSurrogate: Cs other; surrogate.
	KUCGenlCatOtherSurrogate KUCGenlCat = 3
	// KUCGenlCatPunctClose: Pe punctuation; close.
	KUCGenlCatPunctClose KUCGenlCat = 23
	// KUCGenlCatPunctConnector: Pc punctuation; connector.
	KUCGenlCatPunctConnector KUCGenlCat = 20
	// KUCGenlCatPunctDash: Pd punctuation; dash.
	KUCGenlCatPunctDash KUCGenlCat = 21
	// KUCGenlCatPunctFinalQuote: Pf punctuation; final quote.
	KUCGenlCatPunctFinalQuote KUCGenlCat = 25
	// KUCGenlCatPunctInitialQuote: Pi punctuation; initial quote.
	KUCGenlCatPunctInitialQuote KUCGenlCat = 24
	// KUCGenlCatPunctOpen: Ps punctuation; open.
	KUCGenlCatPunctOpen KUCGenlCat = 22
	// KUCGenlCatPunctOther: Po punctuation; other.
	KUCGenlCatPunctOther KUCGenlCat = 26
	// KUCGenlCatSeparatorLine: Zl separator; Line.
	KUCGenlCatSeparatorLine KUCGenlCat = 12
	// KUCGenlCatSeparatorParagraph: Zp separator; paragraph.
	KUCGenlCatSeparatorParagraph KUCGenlCat = 13
	// KUCGenlCatSeparatorSpace: Zs separator; space.
	KUCGenlCatSeparatorSpace KUCGenlCat = 11
	// KUCGenlCatSymbolCurrency: Sc symbol; currency.
	KUCGenlCatSymbolCurrency KUCGenlCat = 29
	// KUCGenlCatSymbolMath: Sm symbol; math.
	KUCGenlCatSymbolMath KUCGenlCat = 28
	// KUCGenlCatSymbolModifier: Sk symbol; modifier.
	KUCGenlCatSymbolModifier KUCGenlCat = 30
	// KUCGenlCatSymbolOther: So symbol; other.
	KUCGenlCatSymbolOther KUCGenlCat = 31
)

func (e KUCGenlCat) String() string {
	switch e {
	case KUCGenlCatLetterLowercase:
		return "KUCGenlCatLetterLowercase"
	case KUCGenlCatLetterModifier:
		return "KUCGenlCatLetterModifier"
	case KUCGenlCatLetterOther:
		return "KUCGenlCatLetterOther"
	case KUCGenlCatLetterTitlecase:
		return "KUCGenlCatLetterTitlecase"
	case KUCGenlCatLetterUppercase:
		return "KUCGenlCatLetterUppercase"
	case KUCGenlCatMarkEnclosing:
		return "KUCGenlCatMarkEnclosing"
	case KUCGenlCatMarkNonSpacing:
		return "KUCGenlCatMarkNonSpacing"
	case KUCGenlCatMarkSpacingCombining:
		return "KUCGenlCatMarkSpacingCombining"
	case KUCGenlCatNumberDecimalDigit:
		return "KUCGenlCatNumberDecimalDigit"
	case KUCGenlCatNumberLetter:
		return "KUCGenlCatNumberLetter"
	case KUCGenlCatNumberOther:
		return "KUCGenlCatNumberOther"
	case KUCGenlCatOtherControl:
		return "KUCGenlCatOtherControl"
	case KUCGenlCatOtherFormat:
		return "KUCGenlCatOtherFormat"
	case KUCGenlCatOtherNotAssigned:
		return "KUCGenlCatOtherNotAssigned"
	case KUCGenlCatOtherPrivateUse:
		return "KUCGenlCatOtherPrivateUse"
	case KUCGenlCatOtherSurrogate:
		return "KUCGenlCatOtherSurrogate"
	case KUCGenlCatPunctClose:
		return "KUCGenlCatPunctClose"
	case KUCGenlCatPunctConnector:
		return "KUCGenlCatPunctConnector"
	case KUCGenlCatPunctDash:
		return "KUCGenlCatPunctDash"
	case KUCGenlCatPunctFinalQuote:
		return "KUCGenlCatPunctFinalQuote"
	case KUCGenlCatPunctInitialQuote:
		return "KUCGenlCatPunctInitialQuote"
	case KUCGenlCatPunctOpen:
		return "KUCGenlCatPunctOpen"
	case KUCGenlCatPunctOther:
		return "KUCGenlCatPunctOther"
	case KUCGenlCatSeparatorLine:
		return "KUCGenlCatSeparatorLine"
	case KUCGenlCatSeparatorParagraph:
		return "KUCGenlCatSeparatorParagraph"
	case KUCGenlCatSeparatorSpace:
		return "KUCGenlCatSeparatorSpace"
	case KUCGenlCatSymbolCurrency:
		return "KUCGenlCatSymbolCurrency"
	case KUCGenlCatSymbolMath:
		return "KUCGenlCatSymbolMath"
	case KUCGenlCatSymbolModifier:
		return "KUCGenlCatSymbolModifier"
	case KUCGenlCatSymbolOther:
		return "KUCGenlCatSymbolOther"
	default:
		return fmt.Sprintf("KUCGenlCat(%d)", e)
	}
}

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

type KUCKeyTranslateNoDeadKeys uint

const (
	// KUCKeyTranslateNoDeadKeysBit: The bit number of the bit that turns off dead-key processing.
	KUCKeyTranslateNoDeadKeysBit KUCKeyTranslateNoDeadKeys = 0
	// KUCKeyTranslateNoDeadKeysMask: The mask for the bit that turns off dead-key processing.
	KUCKeyTranslateNoDeadKeysMask KUCKeyTranslateNoDeadKeys = 1
)

func (e KUCKeyTranslateNoDeadKeys) String() string {
	switch e {
	case KUCKeyTranslateNoDeadKeysBit:
		return "KUCKeyTranslateNoDeadKeysBit"
	case KUCKeyTranslateNoDeadKeysMask:
		return "KUCKeyTranslateNoDeadKeysMask"
	default:
		return fmt.Sprintf("KUCKeyTranslateNoDeadKeys(%d)", e)
	}
}

type KUCTS int

const (
	KUCTSNoKeysAddedToObjectErr KUCTS = -25342
	KUCTSSearchListErr          KUCTS = -25343
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
	KUCTSDirectionNext     KUCTSDirection = 0
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
	KUCTSOptionsNoneMask          KUCTSOptions = 0
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
	KUCTextBreakLineMask      KUCTextBreak = 64
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

type KUCTypeSelectMaxList uint

const (
	KUCTypeSelectMaxListSize KUCTypeSelectMaxList = 0
)

func (e KUCTypeSelectMaxList) String() string {
	switch e {
	case KUCTypeSelectMaxListSize:
		return "KUCTypeSelectMaxListSize"
	default:
		return fmt.Sprintf("KUCTypeSelectMaxList(%d)", e)
	}
}

type KUR int

const (
	KURL68kNotSupportedError          KUR = -30788
	KURLAccessNotAvailableError       KUR = -30787
	KURLAuthenticationError           KUR = -30776
	KURLDestinationExistsError        KUR = -30772
	KURLExtensionFailureError         KUR = -30785
	KURLFileEmptyError                KUR = -30783
	KURLInvalidCallError              KUR = -30781
	KURLInvalidConfigurationError     KUR = -30786
	KURLInvalidURLError               KUR = -30773
	KURLInvalidURLReferenceError      KUR = -30770
	KURLProgressAlreadyDisplayedError KUR = -30771
	KURLPropertyBufferTooSmallError   KUR = -30779
	KURLPropertyNotYetKnownError      KUR = -30777
	KURLServerBusyError               KUR = -30775
	KURLUnknownPropertyError          KUR = -30778
	KURLUnsettablePropertyError       KUR = -30780
	KURLUnsupportedSchemeError        KUR = -30774
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
	KUSBAbortedError           KUSB = -6982
	KUSBAlreadyOpenErr         KUSB = -6991
	KUSBBadDispatchTable       KUSB = -6950
	KUSBBitstufErr             KUSB = -6914
	KUSBBufOvrRunErr           KUSB = -6904
	KUSBBufUnderRunErr         KUSB = -6903
	KUSBCRCErr                 KUSB = -6915
	KUSBCompletionError        KUSB = -6984
	KUSBDataToggleErr          KUSB = -6913
	KUSBDeviceBusy             KUSB = -6977
	KUSBDeviceDisconnected     KUSB = -6972
	KUSBDeviceErr              KUSB = -6989
	KUSBDeviceNotSuspended     KUSB = -6973
	KUSBDevicePowerProblem     KUSB = -6976
	KUSBDeviceSuspended        KUSB = -6974
	KUSBEndpointStallErr       KUSB = -6912
	KUSBFlagsError             KUSB = -6983
	KUSBIncorrectTypeErr       KUSB = -6995
	KUSBInternalErr            KUSB = -6999
	KUSBInvalidBuffer          KUSB = -6975
	KUSBLinkErr                KUSB = -6916
	KUSBNoBandwidthError       KUSB = -6981
	KUSBNoDelay                KUSB = 0
	KUSBNoDeviceErr            KUSB = -6990
	KUSBNoErr                  KUSB = 0
	KUSBNoTran                 KUSB = 0
	KUSBNotFound               KUSB = -6987
	KUSBNotHandled             KUSB = -6987
	KUSBNotRespondingErr       KUSB = -6911
	KUSBNotSent1Err            KUSB = -6902
	KUSBNotSent2Err            KUSB = -6901
	KUSBOutOfMemoryErr         KUSB = -6988
	KUSBOverRunErr             KUSB = -6908
	KUSBPBLengthError          KUSB = -6985
	KUSBPBVersionError         KUSB = -6986
	KUSBPIDCheckErr            KUSB = -6910
	KUSBPending                KUSB = 1
	KUSBPipeIdleError          KUSB = -6980
	KUSBPipeStalledError       KUSB = -6979
	KUSBPortDisabled           KUSB = -6969
	KUSBQueueAborted           KUSB = -6970
	KUSBQueueFull              KUSB = -6948
	KUSBRes1Err                KUSB = -6906
	KUSBRes2Err                KUSB = -6905
	KUSBRqErr                  KUSB = -6994
	KUSBTimedOut               KUSB = -6971
	KUSBTooManyPipesErr        KUSB = -6996
	KUSBTooManyTransactionsErr KUSB = -6992
	KUSBUnderRunErr            KUSB = -6907
	KUSBUnknownDeviceErr       KUSB = -6998
	KUSBUnknownInterfaceErr    KUSB = -6978
	KUSBUnknownNotification    KUSB = -6949
	KUSBUnknownPipeErr         KUSB = -6997
	KUSBUnknownRequestErr      KUSB = -6993
	KUSBWrongPIDErr            KUSB = -6909
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
	KUSBInternalReserved1  KUSBInternal = -6960
	KUSBInternalReserved10 KUSBInternal = -6951
	KUSBInternalReserved2  KUSBInternal = -6959
	KUSBInternalReserved3  KUSBInternal = -6958
	KUSBInternalReserved4  KUSBInternal = -6957
	KUSBInternalReserved5  KUSBInternal = -6956
	KUSBInternalReserved6  KUSBInternal = -6955
	KUSBInternalReserved7  KUSBInternal = -6954
	KUSBInternalReserved8  KUSBInternal = -6953
	KUSBInternalReserved9  KUSBInternal = -6952
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

type KUTCDefault uint

const (
	// Deprecated.
	KUTCDefaultOptions KUTCDefault = 0
)

func (e KUTCDefault) String() string {
	switch e {
	case KUTCDefaultOptions:
		return "KUTCDefaultOptions"
	default:
		return fmt.Sprintf("KUTCDefault(%d)", e)
	}
}

type KUTCUnderflowErr int

const (
	KIllegalClockValueErr KUTCUnderflowErr = -8852
	KUTCOverflowErr       KUTCUnderflowErr = -8851
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

type KUnicode uint

const (
	KUnicodeByteOrderMark KUnicode = 65279
	// KUnicodeDefaultDirection: Use the default direction.
	KUnicodeDefaultDirection KUnicode = 0
	// KUnicodeDefaultDirectionMask: # Discussion
	KUnicodeDefaultDirectionMask KUnicode = 0
	// KUnicodeDirectionalityBits: Sets directionality.
	KUnicodeDirectionalityBits KUnicode = 2
	// KUnicodeDirectionalityMask: A mask for setting the directionality control flag
	KUnicodeDirectionalityMask KUnicode = 12
	// KUnicodeForceASCIIRangeBit: Sets the force ASCII range control flag.
	KUnicodeForceASCIIRangeBit KUnicode = 9
	// KUnicodeForceASCIIRangeMask: # Discussion
	KUnicodeForceASCIIRangeMask KUnicode = 512
	// KUnicodeHFSPlusCompVariant: # Discussion
	KUnicodeHFSPlusCompVariant KUnicode = 9
	// KUnicodeHFSPlusDecompVariant: Specifies canonical decomposition according to Unicode 3.2 rules, with HFS+ exclusions ("HFS+ decomposition 3.2").
	KUnicodeHFSPlusDecompVariant KUnicode = 8
	// KUnicodeKeepInfoBit: Sets the keep-information control flag.
	KUnicodeKeepInfoBit KUnicode = 1
	// KUnicodeKeepInfoMask: # Discussion
	KUnicodeKeepInfoMask KUnicode = 2
	// KUnicodeKeepSameEncodingBit: Sets the keep-same-encoding control flag.
	KUnicodeKeepSameEncodingBit KUnicode = 8
	// KUnicodeKeepSameEncodingMask: # Discussion
	KUnicodeKeepSameEncodingMask KUnicode = 256
	// KUnicodeLeftToRight: Indicates left to right direction.
	KUnicodeLeftToRight KUnicode = 1
	// KUnicodeLeftToRightMask: # Discussion
	KUnicodeLeftToRightMask KUnicode = 1
	// KUnicodeLooseMappingsBit: Enables use of the loose-mapping portion of a character mapping table.
	KUnicodeLooseMappingsBit KUnicode = 5
	// KUnicodeLooseMappingsMask: # Discussion
	KUnicodeLooseMappingsMask      KUnicode = 32
	KUnicodeMapLineFeedToReturnBit KUnicode = 12
	// KUnicodeMapLineFeedToReturnMask: # Discussion
	KUnicodeMapLineFeedToReturnMask KUnicode = 4096
	KUnicodeNoHalfwidthCharsBit     KUnicode = 10
	// KUnicodeNoHalfwidthCharsMask: Sets the no halfwidth characters control flag.
	KUnicodeNoHalfwidthCharsMask KUnicode = 1024
	// KUnicodeNoSubset: The standard Unicode encoded character set in which the full set of Unicode characters are supported.
	KUnicodeNoSubset           KUnicode = 0
	KUnicodeNormalizationFormC KUnicode = 3
	KUnicodeNormalizationFormD KUnicode = 5
	// KUnicodeNotAChar: Not a Unicode character; may be used as a terminator.
	KUnicodeNotAChar KUnicode = 65535
	// KUnicodeObjectReplacement: A placeholder for a non-text object.
	KUnicodeObjectReplacement KUnicode = 65532
	// KUnicodeReplacementChar: Unicode replacement for an input character that cannot be converted.
	KUnicodeReplacementChar KUnicode = 65533
	// KUnicodeRightToLeft: Indicates right to left direction.
	KUnicodeRightToLeft KUnicode = 2
	// KUnicodeRightToLeftMask: # Discussion
	KUnicodeRightToLeftMask KUnicode = 2
	// KUnicodeStringUnterminatedBit: Sets the string-unterminated control flag.
	KUnicodeStringUnterminatedBit KUnicode = 6
	// KUnicodeStringUnterminatedMask: # Discussion
	KUnicodeStringUnterminatedMask KUnicode = 64
	// KUnicodeSwappedByteOrderMark: Not a Unicode character; byte-swapped version of FEFF.
	KUnicodeSwappedByteOrderMark KUnicode = 65534
	// KUnicodeTextRunBit: Sets the text-run control flag.
	KUnicodeTextRunBit            KUnicode = 7
	KUnicodeTextRunHeuristicsBit  KUnicode = 11
	KUnicodeTextRunHeuristicsMask KUnicode = 2048
	// KUnicodeTextRunMask: # Discussion
	KUnicodeTextRunMask                 KUnicode = 128
	KUnicodeUseExternalEncodingFormBit  KUnicode = 13
	KUnicodeUseExternalEncodingFormMask KUnicode = 8192
	// KUnicodeUseFallbacksBit: Enables use of fallback mappings.
	KUnicodeUseFallbacksBit KUnicode = 0
	// KUnicodeUseFallbacksMask: # Discussion
	KUnicodeUseFallbacksMask KUnicode = 1
	// KUnicodeVerticalFormBit: Sets the vertical form control flag.
	KUnicodeVerticalFormBit KUnicode = 4
	// KUnicodeVerticalFormMask: # Discussion
	KUnicodeVerticalFormMask KUnicode = 16
)

func (e KUnicode) String() string {
	switch e {
	case KUnicodeByteOrderMark:
		return "KUnicodeByteOrderMark"
	case KUnicodeDefaultDirection:
		return "KUnicodeDefaultDirection"
	case KUnicodeDirectionalityBits:
		return "KUnicodeDirectionalityBits"
	case KUnicodeDirectionalityMask:
		return "KUnicodeDirectionalityMask"
	case KUnicodeForceASCIIRangeBit:
		return "KUnicodeForceASCIIRangeBit"
	case KUnicodeForceASCIIRangeMask:
		return "KUnicodeForceASCIIRangeMask"
	case KUnicodeHFSPlusDecompVariant:
		return "KUnicodeHFSPlusDecompVariant"
	case KUnicodeKeepInfoBit:
		return "KUnicodeKeepInfoBit"
	case KUnicodeKeepSameEncodingMask:
		return "KUnicodeKeepSameEncodingMask"
	case KUnicodeLooseMappingsBit:
		return "KUnicodeLooseMappingsBit"
	case KUnicodeLooseMappingsMask:
		return "KUnicodeLooseMappingsMask"
	case KUnicodeMapLineFeedToReturnMask:
		return "KUnicodeMapLineFeedToReturnMask"
	case KUnicodeNoHalfwidthCharsBit:
		return "KUnicodeNoHalfwidthCharsBit"
	case KUnicodeNoHalfwidthCharsMask:
		return "KUnicodeNoHalfwidthCharsMask"
	case KUnicodeNormalizationFormC:
		return "KUnicodeNormalizationFormC"
	case KUnicodeNotAChar:
		return "KUnicodeNotAChar"
	case KUnicodeObjectReplacement:
		return "KUnicodeObjectReplacement"
	case KUnicodeReplacementChar:
		return "KUnicodeReplacementChar"
	case KUnicodeStringUnterminatedBit:
		return "KUnicodeStringUnterminatedBit"
	case KUnicodeStringUnterminatedMask:
		return "KUnicodeStringUnterminatedMask"
	case KUnicodeSwappedByteOrderMark:
		return "KUnicodeSwappedByteOrderMark"
	case KUnicodeTextRunBit:
		return "KUnicodeTextRunBit"
	case KUnicodeTextRunHeuristicsBit:
		return "KUnicodeTextRunHeuristicsBit"
	case KUnicodeTextRunHeuristicsMask:
		return "KUnicodeTextRunHeuristicsMask"
	case KUnicodeTextRunMask:
		return "KUnicodeTextRunMask"
	case KUnicodeUseExternalEncodingFormBit:
		return "KUnicodeUseExternalEncodingFormBit"
	case KUnicodeUseExternalEncodingFormMask:
		return "KUnicodeUseExternalEncodingFormMask"
	case KUnicodeVerticalFormBit:
		return "KUnicodeVerticalFormBit"
	case KUnicodeVerticalFormMask:
		return "KUnicodeVerticalFormMask"
	default:
		return fmt.Sprintf("KUnicode(%d)", e)
	}
}

type KUnicodeCollation uint

const (
	// KUnicodeCollationClass: Identifies collation as a class of operations.
	KUnicodeCollationClass KUnicodeCollation = 'u'<<24 | 'c'<<16 | 'o'<<8 | 'l' // 'ucol'
)

func (e KUnicodeCollation) String() string {
	switch e {
	case KUnicodeCollationClass:
		return "KUnicodeCollationClass"
	default:
		return fmt.Sprintf("KUnicodeCollation(%d)", e)
	}
}

type KUnicodeFallback uint

const (
	// KUnicodeFallbackCustomFirst: Use the custom fallback handler first, then the default one.
	KUnicodeFallbackCustomFirst KUnicodeFallback = 3
	// KUnicodeFallbackCustomOnly: Use the custom fallback handler only.
	KUnicodeFallbackCustomOnly KUnicodeFallback = 1
	// KUnicodeFallbackDefaultFirst: Use the default fallback handler first, then the custom one.
	KUnicodeFallbackDefaultFirst KUnicodeFallback = 2
	// KUnicodeFallbackDefaultOnly: Use the default fallback handler only.
	KUnicodeFallbackDefaultOnly KUnicodeFallback = 0
	// KUnicodeFallbackInterruptSafeMask: Indicate that the caller’s fallback routine doesn’t move memory.
	KUnicodeFallbackInterruptSafeMask KUnicodeFallback = 4
	KUnicodeFallbackSequencingMask    KUnicodeFallback = 3
)

func (e KUnicodeFallback) String() string {
	switch e {
	case KUnicodeFallbackCustomFirst:
		return "KUnicodeFallbackCustomFirst"
	case KUnicodeFallbackCustomOnly:
		return "KUnicodeFallbackCustomOnly"
	case KUnicodeFallbackDefaultFirst:
		return "KUnicodeFallbackDefaultFirst"
	case KUnicodeFallbackDefaultOnly:
		return "KUnicodeFallbackDefaultOnly"
	case KUnicodeFallbackInterruptSafeMask:
		return "KUnicodeFallbackInterruptSafeMask"
	default:
		return fmt.Sprintf("KUnicodeFallback(%d)", e)
	}
}

type KUnicodeFallbackSequencing uint

const (
	KUnicodeFallbackSequencingBits KUnicodeFallbackSequencing = 0
)

func (e KUnicodeFallbackSequencing) String() string {
	switch e {
	case KUnicodeFallbackSequencingBits:
		return "KUnicodeFallbackSequencingBits"
	default:
		return fmt.Sprintf("KUnicodeFallbackSequencing(%d)", e)
	}
}

type KUnicodeMatch uint

const (
	// KUnicodeMatchOtherBaseBit: Excludes mappings that do not match the text encoding base of the `otherEncoding` field of the structure UnicodeMapping.
	KUnicodeMatchOtherBaseBit KUnicodeMatch = 3
	// KUnicodeMatchOtherBaseMask: If set, excludes mappings that do not match the text encoding base of the `otherEncoding` field of the structure UnicodeMapping.
	KUnicodeMatchOtherBaseMask KUnicodeMatch = 8
	// KUnicodeMatchOtherFormatBit: Excludes mappings that do not match the text encoding format of the `otherEncoding` field of the specified Unicode mapping structure.
	KUnicodeMatchOtherFormatBit KUnicodeMatch = 5
	// KUnicodeMatchOtherFormatMask: If set, excludes mappings that do not match the text encoding format of the `otherEncoding` field of the specified Unicode mapping structure.
	KUnicodeMatchOtherFormatMask KUnicodeMatch = 32
	// KUnicodeMatchOtherVariantBit: Excludes mappings that do not match the text encoding variant of the `otherEncoding` field of the specified Unicode mapping structure.
	KUnicodeMatchOtherVariantBit KUnicodeMatch = 4
	// KUnicodeMatchOtherVariantMask: If set, excludes mappings that do not match the text encoding variant of the `otherEncoding` field of the specified Unicode mapping structure.
	KUnicodeMatchOtherVariantMask KUnicodeMatch = 16
	// KUnicodeMatchUnicodeBaseBit: Excludes mappings that do not match the text encoding base of the `unicodeEncoding` field of the structure UnicodeMapping.
	KUnicodeMatchUnicodeBaseBit KUnicodeMatch = 0
	// KUnicodeMatchUnicodeBaseMask: If set, excludes mappings that do not match the text encoding base of the `unicodeEncoding` field of the structure UnicodeMapping.
	KUnicodeMatchUnicodeBaseMask KUnicodeMatch = 1
	// KUnicodeMatchUnicodeFormatBit: Excludes mappings that do not match the text encoding format of the `unicodeEncoding` field of the specified Unicode mapping structure.
	KUnicodeMatchUnicodeFormatBit KUnicodeMatch = 2
	// KUnicodeMatchUnicodeFormatMask: If set, excludes mappings that do not match the text encoding format of the `unicodeEncoding` field of the specified Unicode mapping structure.
	KUnicodeMatchUnicodeFormatMask KUnicodeMatch = 4
	// KUnicodeMatchUnicodeVariantBit: Excludes mappings that do not match the text encoding variant of the `unicodeEncoding` field of the specified Unicode mapping structure.
	KUnicodeMatchUnicodeVariantBit KUnicodeMatch = 1
	// KUnicodeMatchUnicodeVariantMask: If set, excludes mappings that do not match the text encoding variant of the `unicodeEncoding` field of the specified Unicode mapping structure.
	KUnicodeMatchUnicodeVariantMask KUnicodeMatch = 2
)

func (e KUnicodeMatch) String() string {
	switch e {
	case KUnicodeMatchOtherBaseBit:
		return "KUnicodeMatchOtherBaseBit"
	case KUnicodeMatchOtherBaseMask:
		return "KUnicodeMatchOtherBaseMask"
	case KUnicodeMatchOtherFormatBit:
		return "KUnicodeMatchOtherFormatBit"
	case KUnicodeMatchOtherFormatMask:
		return "KUnicodeMatchOtherFormatMask"
	case KUnicodeMatchOtherVariantBit:
		return "KUnicodeMatchOtherVariantBit"
	case KUnicodeMatchOtherVariantMask:
		return "KUnicodeMatchOtherVariantMask"
	case KUnicodeMatchUnicodeBaseBit:
		return "KUnicodeMatchUnicodeBaseBit"
	case KUnicodeMatchUnicodeBaseMask:
		return "KUnicodeMatchUnicodeBaseMask"
	case KUnicodeMatchUnicodeFormatBit:
		return "KUnicodeMatchUnicodeFormatBit"
	default:
		return fmt.Sprintf("KUnicodeMatch(%d)", e)
	}
}

type KUnicodeNo uint

const (
	KUnicodeNoCompatibilityVariant KUnicodeNo = 1
	KUnicodeNoCorporateVariant     KUnicodeNo = 4
)

func (e KUnicodeNo) String() string {
	switch e {
	case KUnicodeNoCompatibilityVariant:
		return "KUnicodeNoCompatibilityVariant"
	case KUnicodeNoCorporateVariant:
		return "KUnicodeNoCorporateVariant"
	default:
		return fmt.Sprintf("KUnicodeNo(%d)", e)
	}
}

type KUnicodeTextBreak uint

const (
	// KUnicodeTextBreakClass: Identifies the class of Unicode utility operations that find text boundaries.
	KUnicodeTextBreakClass KUnicodeTextBreak = 'u'<<24 | 'b'<<16 | 'r'<<8 | 'k' // 'ubrk'
)

func (e KUnicodeTextBreak) String() string {
	switch e {
	case KUnicodeTextBreakClass:
		return "KUnicodeTextBreakClass"
	default:
		return fmt.Sprintf("KUnicodeTextBreak(%d)", e)
	}
}

type KUnicodeUse int

const (
	// KUnicodeUseHFSPlusMapping: Indicates the mapping version used by HFS Plus to convert filenames between Mac OS encodings and Unicode.
	KUnicodeUseHFSPlusMapping KUnicodeUse = 4
	// KUnicodeUseLatestMapping: Instead of explicitly specifying the mapping version of the Unicode mapping table to be used for conversion of a text string, you can use this constant to specify that the latest version be used.
	KUnicodeUseLatestMapping KUnicodeUse = -1
)

func (e KUnicodeUse) String() string {
	switch e {
	case KUnicodeUseHFSPlusMapping:
		return "KUnicodeUseHFSPlusMapping"
	case KUnicodeUseLatestMapping:
		return "KUnicodeUseLatestMapping"
	default:
		return fmt.Sprintf("KUnicodeUse(%d)", e)
	}
}

type KUnknownException uint

const (
	// Deprecated.
	KAccessException KUnknownException = 3
	// Deprecated.
	KDataAlignmentException KUnknownException = 17
	// Deprecated.
	KDataBreakpointException KUnknownException = 11
	// Deprecated.
	KExcludedMemoryException KUnknownException = 5
	// Deprecated.
	KFloatingPointException KUnknownException = 13
	// Deprecated.
	KIllegalInstructionException KUnknownException = 1
	// Deprecated.
	KInstructionBreakpointException KUnknownException = 10
	// Deprecated.
	KIntegerException KUnknownException = 12
	// Deprecated.
	KPrivilegeViolationException KUnknownException = 8
	// Deprecated.
	KReadOnlyMemoryException KUnknownException = 6
	// Deprecated.
	KStackOverflowException KUnknownException = 14
	// Deprecated.
	KTaskCreationException KUnknownException = 16
	// Deprecated.
	KTaskTerminationException KUnknownException = 15
	// Deprecated.
	KTraceException KUnknownException = 9
	// Deprecated.
	KTrapException KUnknownException = 2
	// Deprecated.
	KUnknownExceptionValue KUnknownException = 0
	// Deprecated.
	KUnmappedMemoryException KUnknownException = 4
	// Deprecated.
	KUnresolvablePageFaultException KUnknownException = 7
)

func (e KUnknownException) String() string {
	switch e {
	case KAccessException:
		return "KAccessException"
	case KDataAlignmentException:
		return "KDataAlignmentException"
	case KDataBreakpointException:
		return "KDataBreakpointException"
	case KExcludedMemoryException:
		return "KExcludedMemoryException"
	case KFloatingPointException:
		return "KFloatingPointException"
	case KIllegalInstructionException:
		return "KIllegalInstructionException"
	case KInstructionBreakpointException:
		return "KInstructionBreakpointException"
	case KIntegerException:
		return "KIntegerException"
	case KPrivilegeViolationException:
		return "KPrivilegeViolationException"
	case KReadOnlyMemoryException:
		return "KReadOnlyMemoryException"
	case KStackOverflowException:
		return "KStackOverflowException"
	case KTaskCreationException:
		return "KTaskCreationException"
	case KTaskTerminationException:
		return "KTaskTerminationException"
	case KTraceException:
		return "KTraceException"
	case KTrapException:
		return "KTrapException"
	case KUnknownExceptionValue:
		return "KUnknownExceptionValue"
	case KUnmappedMemoryException:
		return "KUnmappedMemoryException"
	case KUnresolvablePageFaultException:
		return "KUnresolvablePageFaultException"
	default:
		return fmt.Sprintf("KUnknownException(%d)", e)
	}
}

type KUnlockStateKCStatus int

const (
	KRdPermKCStatus           KUnlockStateKCStatus = 2
	KUnlockStateKCStatusValue KUnlockStateKCStatus = 1
	KWrPermKCStatus           KUnlockStateKCStatus = 4
)

func (e KUnlockStateKCStatus) String() string {
	switch e {
	case KRdPermKCStatus:
		return "KRdPermKCStatus"
	case KUnlockStateKCStatusValue:
		return "KUnlockStateKCStatusValue"
	case KWrPermKCStatus:
		return "KWrPermKCStatus"
	default:
		return fmt.Sprintf("KUnlockStateKCStatus(%d)", e)
	}
}

type KUse uint

const (
	// Deprecated.
	KUseCurrentISA KUse = 0
	// Deprecated.
	KUseNativeISA KUse = 4
)

func (e KUse) String() string {
	switch e {
	case KUseCurrentISA:
		return "KUseCurrentISA"
	case KUseNativeISA:
		return "KUseNativeISA"
	default:
		return fmt.Sprintf("KUse(%d)", e)
	}
}

type KUserFolderIcon uint

const (
	KGroupIcon           KUserFolderIcon = 'g'<<24 | 'r'<<16 | 'u'<<8 | 'p' // 'grup'
	KGuestUserIcon       KUserFolderIcon = 'g'<<24 | 'u'<<16 | 's'<<8 | 'r' // 'gusr'
	KOwnerIcon           KUserFolderIcon = 's'<<24 | 'u'<<16 | 's'<<8 | 'r' // 'susr'
	KUserFolderIconValue KUserFolderIcon = 'u'<<24 | 'f'<<16 | 'l'<<8 | 'd' // 'ufld'
	KUserIcon            KUserFolderIcon = 'u'<<24 | 's'<<16 | 'e'<<8 | 'r' // 'user'
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

type KVCBFlags uint

const (
	KVCBFlagsHFSPlusAPIsBit   KVCBFlags = 4
	KVCBFlagsHFSPlusAPIsMask  KVCBFlags = 16
	KVCBFlagsHardwareGoneBit  KVCBFlags = 5
	KVCBFlagsHardwareGoneMask KVCBFlags = 32
	KVCBFlagsIdleFlushBit     KVCBFlags = 3
	KVCBFlagsIdleFlushMask    KVCBFlags = 8
	KVCBFlagsVolumeDirtyBit   KVCBFlags = 15
	KVCBFlagsVolumeDirtyMask  KVCBFlags = 32768
)

func (e KVCBFlags) String() string {
	switch e {
	case KVCBFlagsHFSPlusAPIsBit:
		return "KVCBFlagsHFSPlusAPIsBit"
	case KVCBFlagsHFSPlusAPIsMask:
		return "KVCBFlagsHFSPlusAPIsMask"
	case KVCBFlagsHardwareGoneBit:
		return "KVCBFlagsHardwareGoneBit"
	case KVCBFlagsHardwareGoneMask:
		return "KVCBFlagsHardwareGoneMask"
	case KVCBFlagsIdleFlushBit:
		return "KVCBFlagsIdleFlushBit"
	case KVCBFlagsIdleFlushMask:
		return "KVCBFlagsIdleFlushMask"
	case KVCBFlagsVolumeDirtyBit:
		return "KVCBFlagsVolumeDirtyBit"
	case KVCBFlagsVolumeDirtyMask:
		return "KVCBFlagsVolumeDirtyMask"
	default:
		return fmt.Sprintf("KVCBFlags(%d)", e)
	}
}

type KWidePosOffsetBit uint

const (
	KMaximumBlocksIn4GB    KWidePosOffsetBit = 8388607
	KUseWidePositioning    KWidePosOffsetBit = 256
	KWidePosOffsetBitValue KWidePosOffsetBit = 8
)

func (e KWidePosOffsetBit) String() string {
	switch e {
	case KMaximumBlocksIn4GB:
		return "KMaximumBlocksIn4GB"
	case KUseWidePositioning:
		return "KUseWidePositioning"
	case KWidePosOffsetBitValue:
		return "KWidePosOffsetBitValue"
	default:
		return fmt.Sprintf("KWidePosOffsetBit(%d)", e)
	}
}

type KWidgetsFolderType uint

const (
	KScreenSaversFolderType KWidgetsFolderType = 's'<<24 | 'c'<<16 | 'r'<<8 | 'n' // 'scrn'
	KWidgetsFolderTypeValue KWidgetsFolderType = 'w'<<24 | 'd'<<16 | 'g'<<8 | 't' // 'wdgt'
)

func (e KWidgetsFolderType) String() string {
	switch e {
	case KScreenSaversFolderType:
		return "KScreenSaversFolderType"
	case KWidgetsFolderTypeValue:
		return "KWidgetsFolderTypeValue"
	default:
		return fmt.Sprintf("KWidgetsFolderType(%d)", e)
	}
}

type KWindowsLatin1 uint

const (
	KWindowsLatin1PalmVariant     KWindowsLatin1 = 1
	KWindowsLatin1StandardVariant KWindowsLatin1 = 0
)

func (e KWindowsLatin1) String() string {
	switch e {
	case KWindowsLatin1PalmVariant:
		return "KWindowsLatin1PalmVariant"
	case KWindowsLatin1StandardVariant:
		return "KWindowsLatin1StandardVariant"
	default:
		return fmt.Sprintf("KWindowsLatin1(%d)", e)
	}
}

type KWriteReference uint

const (
	// Deprecated.
	FetchReference KWriteReference = 2
	// Deprecated.
	KFetchReference KWriteReference = 2
	// Deprecated.
	KReadReference KWriteReference = 1
	// Deprecated.
	KWriteReferenceValue KWriteReference = 0
	// Deprecated.
	ReadReference KWriteReference = 1
	// Deprecated.
	WriteReference KWriteReference = 0
)

func (e KWriteReference) String() string {
	switch e {
	case FetchReference:
		return "FetchReference"
	case KReadReference:
		return "KReadReference"
	case KWriteReferenceValue:
		return "KWriteReferenceValue"
	default:
		return fmt.Sprintf("KWriteReference(%d)", e)
	}
}

type KX86IS uint

const (
	// Deprecated.
	KX86ISA KX86IS = 2
)

func (e KX86IS) String() string {
	switch e {
	case KX86ISA:
		return "KX86ISA"
	default:
		return fmt.Sprintf("KX86IS(%d)", e)
	}
}

type KX86RT uint

const (
	// Deprecated.
	KX86RTA KX86RT = 32
)

func (e KX86RT) String() string {
	switch e {
	case KX86RTA:
		return "KX86RTA"
	default:
		return fmt.Sprintf("KX86RT(%d)", e)
	}
}

type KXLibTag1 uint

const (
	// Deprecated.
	KBLibTag2 KXLibTag1 = 'B'<<24 | 'L'<<16 | 'i'<<8 | 'b' // 'BLib'
	// Deprecated.
	KVLibTag2 KXLibTag1 = 'V'<<24 | 'L'<<16 | 'i'<<8 | 'b' // 'VLib'
	// Deprecated.
	KXLibTag1Value KXLibTag1 = 0
	// Deprecated.
	KXLibVersion KXLibTag1 = 1
)

func (e KXLibTag1) String() string {
	switch e {
	case KBLibTag2:
		return "KBLibTag2"
	case KVLibTag2:
		return "KVLibTag2"
	case KXLibTag1Value:
		return "KXLibTag1Value"
	case KXLibVersion:
		return "KXLibVersion"
	default:
		return fmt.Sprintf("KXLibTag1(%d)", e)
	}
}

type Kernel int

const (
	KernelAlreadyFreeErr       Kernel = -2421
	KernelAsyncReceiveLimitErr Kernel = -2414
	KernelAsyncSendLimitErr    Kernel = -2413
	KernelAttributeErr         Kernel = -2412
	KernelCanceledErr          Kernel = -2402
	KernelDeletePermissionErr  Kernel = -2410
	KernelExceptionErr         Kernel = -2418
	KernelExecutePermissionErr Kernel = -2409
	KernelExecutionLevelErr    Kernel = -2411
	KernelIDErr                Kernel = -2419
	KernelInUseErr             Kernel = -2416
	KernelIncompleteErr        Kernel = -2401
	KernelObjectExistsErr      Kernel = -2406
	KernelOptionsErr           Kernel = -2403
	KernelPrivilegeErr         Kernel = -2404
	KernelReadPermissionErr    Kernel = -2408
	KernelReturnValueErr       Kernel = -2422
	KernelTerminatedErr        Kernel = -2417
	KernelTimeoutErr           Kernel = -2415
	KernelUnrecoverableErr     Kernel = -2499
	KernelUnsupportedErr       Kernel = -2405
	KernelWritePermissionErr   Kernel = -2407
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

type KeucCn uint

const (
	KEUC_CN_BasicVariant KeucCn = 0
	KEUC_CN_DOSVariant   KeucCn = 1
)

func (e KeucCn) String() string {
	switch e {
	case KEUC_CN_BasicVariant:
		return "KEUC_CN_BasicVariant"
	case KEUC_CN_DOSVariant:
		return "KEUC_CN_DOSVariant"
	default:
		return fmt.Sprintf("KeucCn(%d)", e)
	}
}

type KeucKr uint

const (
	KEUC_KR_BasicVariant KeucKr = 0
	KEUC_KR_DOSVariant   KeucKr = 1
)

func (e KeucKr) String() string {
	switch e {
	case KEUC_KR_BasicVariant:
		return "KEUC_KR_BasicVariant"
	case KEUC_KR_DOSVariant:
		return "KEUC_KR_DOSVariant"
	default:
		return fmt.Sprintf("KeucKr(%d)", e)
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
	// KeyAERecorderCount: Used with the `keyword` parameter of the AEManagerInfo function.
	KeyAERecorderCount Key = 'r'<<24 | 'e'<<16 | 'c'<<8 | 'r' // 'recr'
	// KeyAEVersion: Used with the `keyword` parameter of the AEManagerInfo function.
	KeyAEVersion              Key = 'v'<<24 | 'e'<<16 | 'r'<<8 | 's' // 'vers'
	KeyAcceptTimeoutAttr      Key = 'a'<<24 | 'c'<<16 | 't'<<8 | 'm' // 'actm'
	KeyActualSenderAuditToken Key = 'a'<<24 | 'c'<<16 | 'a'<<8 | 't' // 'acat'
	// KeyAddressAttr: Address of a target or client application.
	KeyAddressAttr              Key = 'a'<<24 | 'd'<<16 | 'd'<<8 | 'r' // 'addr'
	KeyAppleEventAttributesAttr Key = 'a'<<24 | 't'<<16 | 't'<<8 | 'r' // 'attr'
	KeyCloseAllWindows          Key = 'c'<<24 | 'a'<<16 | 'w'<<8 | ' ' // 'caw '
	// KeyDirectObject: Direct parameter.
	KeyDirectObject Key = '-'<<24 | '-'<<16 | '-'<<8 | '-' // '----'
	// KeyDisposeTokenProc: Token disposal function.
	KeyDisposeTokenProc Key = 'x'<<24 | 't'<<16 | 'o'<<8 | 'k' // 'xtok'
	KeyDriveNumber      Key = 'd'<<24 | 'r'<<16 | 'v'<<8 | '#' // 'drv#'
	KeyErrorCode        Key = 'e'<<24 | 'r'<<16 | 'r'<<8 | '#' // 'err#'
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
	KeyHighLevelClass  Key = 'h'<<24 | 'c'<<16 | 'l'<<8 | 's' // 'hcls'
	KeyHighLevelID     Key = 'h'<<24 | 'i'<<16 | 'd'<<8 | ' ' // 'hid '
	// KeyInteractLevelAttr: Settings for when to allow the Apple Event Manager to bring a server application to the foreground, if necessary, to interact with the user.
	KeyInteractLevelAttr Key = 'i'<<24 | 'n'<<16 | 't'<<8 | 'e' // 'inte'
	KeyKey               Key = 'k'<<24 | 'e'<<16 | 'y'<<8 | ' ' // 'key '
	KeyKeyCode           Key = 'c'<<24 | 'o'<<16 | 'd'<<8 | 'e' // 'code'
	KeyKeyboard          Key = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'b' // 'keyb'
	KeyLocalWhere        Key = 'l'<<24 | 'w'<<16 | 'h'<<8 | 'r' // 'lwhr'
	KeyMenuID            Key = 'm'<<24 | 'i'<<16 | 'd'<<8 | ' ' // 'mid '
	KeyMenuItem          Key = 'm'<<24 | 'i'<<16 | 't'<<8 | 'm' // 'mitm'
	KeyMiscellaneous     Key = 'f'<<24 | 'm'<<16 | 's'<<8 | 'c' // 'fmsc'
	// KeyMissedKeywordAttr: # Discussion
	KeyMissedKeywordAttr Key = 'm'<<24 | 'i'<<16 | 's'<<8 | 's' // 'miss'
	KeyModifiers         Key = 'm'<<24 | 'o'<<16 | 'd'<<8 | 's' // 'mods'
	KeyNewBounds         Key = 'n'<<24 | 'b'<<16 | 'n'<<8 | 'd' // 'nbnd'
	// KeyOptionalKeywordAttr: List of keywords for parameters of an Apple event that should be treated as optional by the target application.
	KeyOptionalKeywordAttr Key = 'o'<<24 | 'p'<<16 | 't'<<8 | 'k' // 'optk'
	// KeyOriginalAddressAttr: Address of original source of Apple event if the event has been forwarded (available only in version 1.01 or later versions of the Apple Event Manager).
	KeyOriginalAddressAttr Key = 'f'<<24 | 'r'<<16 | 'o'<<8 | 'm' // 'from'
	KeyOriginalBounds      Key = 'o'<<24 | 'b'<<16 | 'n'<<8 | 'd' // 'obnd'
	// KeyPreDispatch: A predispatch handler (an Apple event handler that the Apple Event Manager calls immediately before it dispatches an Apple event).
	KeyPreDispatch Key = 'p'<<24 | 'h'<<16 | 'a'<<8 | 'c' // 'phac'
	// KeyProcessSerialNumber: Process serial number.
	KeyProcessSerialNumber Key = 'p'<<24 | 's'<<16 | 'n'<<8 | ' ' // 'psn '
	// KeyReplyRequestedAttr: A Boolean value indicating whether the Apple event expects to be replied to.
	KeyReplyRequestedAttr Key = 'r'<<24 | 'e'<<16 | 'p'<<8 | 'q' // 'repq'
	// KeyReturnIDAttr: Return ID for a reply Apple event.
	KeyReturnIDAttr Key = 'r'<<24 | 't'<<16 | 'i'<<8 | 'd' // 'rtid'
	// KeySelectProc: You pass this value in the `functionClass` parameter of the AEManagerInfo function to disable the Object Support Library.
	KeySelectProc                                 Key = 's'<<24 | 'e'<<16 | 'l'<<8 | 'h' // 'selh'
	KeySelection                                  Key = 'f'<<24 | 's'<<16 | 'e'<<8 | 'l' // 'fsel'
	KeySenderApplescriptEntitlementsAttr          Key = 'e'<<24 | 'n'<<16 | 't'<<8 | 'l' // 'entl'
	KeySenderApplicationIdentifierEntitlementAttr Key = 'a'<<24 | 'i'<<16 | 'e'<<8 | 'a' // 'aiea'
	KeySenderApplicationSandboxed                 Key = 's'<<24 | 's'<<16 | 's'<<8 | 'b' // 'sssb'
	KeySenderAuditTokenAttr                       Key = 't'<<24 | 'o'<<16 | 'k'<<8 | 'n' // 'tokn'
	KeySenderEGIDAttr                             Key = 's'<<24 | 'g'<<16 | 'i'<<8 | 'd' // 'sgid'
	KeySenderEUIDAttr                             Key = 's'<<24 | 'e'<<16 | 'i'<<8 | 'd' // 'seid'
	KeySenderGIDAttr                              Key = 'g'<<24 | 'i'<<16 | 'd'<<8 | 's' // 'gids'
	KeySenderPIDAttr                              Key = 's'<<24 | 'p'<<16 | 'i'<<8 | 'd' // 'spid'
	KeySenderUIDAttr                              Key = 'u'<<24 | 'i'<<16 | 'd'<<8 | 's' // 'uids'
	// KeyTimeoutAttr: Length of time, in ticks, that the client will wait for a reply or a result from the server.
	KeyTimeoutAttr Key = 't'<<24 | 'i'<<16 | 'm'<<8 | 'o' // 'timo'
	// KeyTransactionIDAttr: Transaction ID identifying a series of Apple events that are part of one transaction.
	KeyTransactionIDAttr Key = 't'<<24 | 'r'<<16 | 'a'<<8 | 'n' // 'tran'
	KeyWhen              Key = 'w'<<24 | 'h'<<16 | 'e'<<8 | 'n' // 'when'
	KeyWhere             Key = 'w'<<24 | 'h'<<16 | 'e'<<8 | 'r' // 'wher'
	KeyWindow            Key = 'k'<<24 | 'w'<<16 | 'n'<<8 | 'd' // 'kwnd'
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
	case KeyAppleEventAttributesAttr:
		return "KeyAppleEventAttributesAttr"
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
	KeyAEAngle          KeyAE = 'k'<<24 | 'a'<<16 | 'n'<<8 | 'g' // 'kang'
	KeyAEArcAngle       KeyAE = 'p'<<24 | 'a'<<16 | 'r'<<8 | 'c' // 'parc'
	KeyAEBaseAddr       KeyAE = 'b'<<24 | 'a'<<16 | 'd'<<8 | 'd' // 'badd'
	KeyAEBestType       KeyAE = 'p'<<24 | 'b'<<16 | 's'<<8 | 't' // 'pbst'
	KeyAEBgndColor      KeyAE = 'k'<<24 | 'b'<<16 | 'c'<<8 | 'l' // 'kbcl'
	KeyAEBgndPattern    KeyAE = 'k'<<24 | 'b'<<16 | 'p'<<8 | 't' // 'kbpt'
	KeyAEBounds         KeyAE = 'p'<<24 | 'b'<<16 | 'n'<<8 | 'd' // 'pbnd'
	KeyAECellList       KeyAE = 'k'<<24 | 'c'<<16 | 'l'<<8 | 't' // 'kclt'
	KeyAEClassID        KeyAE = 'c'<<24 | 'l'<<16 | 'I'<<8 | 'D' // 'clID'
	KeyAEClauseOffsets  KeyAE = 'c'<<24 | 'l'<<16 | 'a'<<8 | 'u' // 'clau'
	KeyAEColor          KeyAE = 'c'<<24 | 'o'<<16 | 'l'<<8 | 'r' // 'colr'
	KeyAEColorTable     KeyAE = 'c'<<24 | 'l'<<16 | 't'<<8 | 'b' // 'cltb'
	KeyAECurveHeight    KeyAE = 'k'<<24 | 'c'<<16 | 'h'<<8 | 'd' // 'kchd'
	KeyAECurveWidth     KeyAE = 'k'<<24 | 'c'<<16 | 'w'<<8 | 'd' // 'kcwd'
	KeyAEDashStyle      KeyAE = 'p'<<24 | 'd'<<16 | 's'<<8 | 't' // 'pdst'
	KeyAEData           KeyAE = 'd'<<24 | 'a'<<16 | 't'<<8 | 'a' // 'data'
	KeyAEDefaultType    KeyAE = 'd'<<24 | 'e'<<16 | 'f'<<8 | 't' // 'deft'
	KeyAEDefinitionRect KeyAE = 'p'<<24 | 'd'<<16 | 'r'<<8 | 't' // 'pdrt'
	KeyAEDescType       KeyAE = 'd'<<24 | 's'<<16 | 't'<<8 | 'p' // 'dstp'
	KeyAEDestination    KeyAE = 'd'<<24 | 'e'<<16 | 's'<<8 | 't' // 'dest'
	KeyAEDoAntiAlias    KeyAE = 'a'<<24 | 'n'<<16 | 't'<<8 | 'a' // 'anta'
	KeyAEDoDithered     KeyAE = 'g'<<24 | 'd'<<16 | 'i'<<8 | 't' // 'gdit'
	KeyAEDoRotate       KeyAE = 'k'<<24 | 'd'<<16 | 'r'<<8 | 't' // 'kdrt'
	KeyAEDoScale        KeyAE = 'k'<<24 | 's'<<16 | 'c'<<8 | 'a' // 'ksca'
	KeyAEDoTranslate    KeyAE = 'k'<<24 | 't'<<16 | 'r'<<8 | 'a' // 'ktra'
	KeyAEDragging       KeyAE = 'b'<<24 | 'o'<<16 | 'o'<<8 | 'l' // 'bool'
	KeyAEEditionFileLoc KeyAE = 'e'<<24 | 'l'<<16 | 'o'<<8 | 'c' // 'eloc'
	KeyAEElements       KeyAE = 'e'<<24 | 'l'<<16 | 'm'<<8 | 's' // 'elms'
	KeyAEEndPoint       KeyAE = 'p'<<24 | 'e'<<16 | 'n'<<8 | 'd' // 'pend'
	KeyAEEventClass     KeyAE = 'e'<<24 | 'v'<<16 | 'c'<<8 | 'l' // 'evcl'
	KeyAEEventID        KeyAE = 'e'<<24 | 'v'<<16 | 't'<<8 | 'i' // 'evti'
	KeyAEFile           KeyAE = 'k'<<24 | 'f'<<16 | 'i'<<8 | 'l' // 'kfil'
	KeyAEFileType       KeyAE = 'f'<<24 | 'l'<<16 | 't'<<8 | 'p' // 'fltp'
	KeyAEFillColor      KeyAE = 'f'<<24 | 'l'<<16 | 'c'<<8 | 'l' // 'flcl'
	KeyAEFillPattern    KeyAE = 'f'<<24 | 'l'<<16 | 'p'<<8 | 't' // 'flpt'
	KeyAEFlipHorizontal KeyAE = 'k'<<24 | 'f'<<16 | 'h'<<8 | 'o' // 'kfho'
	KeyAEFlipVertical   KeyAE = 'k'<<24 | 'f'<<16 | 'v'<<8 | 't' // 'kfvt'
	KeyAEFont           KeyAE = 'f'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'font'
	KeyAEFormula        KeyAE = 'p'<<24 | 'f'<<16 | 'o'<<8 | 'r' // 'pfor'
	KeyAEGraphicObjects KeyAE = 'g'<<24 | 'o'<<16 | 'b'<<8 | 's' // 'gobs'
	KeyAEHiliteRange    KeyAE = 'h'<<24 | 'r'<<16 | 'n'<<8 | 'g' // 'hrng'
	KeyAEID             KeyAE = 'I'<<24 | 'D'<<16 | ' '<<8 | ' ' // 'ID  '
	KeyAEImageQuality   KeyAE = 'g'<<24 | 'q'<<16 | 'u'<<8 | 'a' // 'gqua'
	KeyAEInsertHere     KeyAE = 'i'<<24 | 'n'<<16 | 's'<<8 | 'h' // 'insh'
	KeyAEKeyForms       KeyAE = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'f' // 'keyf'
	KeyAEKeyword        KeyAE = 'k'<<24 | 'y'<<16 | 'w'<<8 | 'd' // 'kywd'
	KeyAELeftSide       KeyAE = 'k'<<24 | 'l'<<16 | 'e'<<8 | 'f' // 'klef'
	KeyAELevel          KeyAE = 'l'<<24 | 'e'<<16 | 'v'<<8 | 'l' // 'levl'
	KeyAELineArrow      KeyAE = 'a'<<24 | 'r'<<16 | 'r'<<8 | 'o' // 'arro'
	KeyAEName           KeyAE = 'p'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'pnam'
	KeyAENewElementLoc  KeyAE = 'p'<<24 | 'n'<<16 | 'e'<<8 | 'l' // 'pnel'
	KeyAEObject         KeyAE = 'k'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'kobj'
	KeyAEObjectClass    KeyAE = 'k'<<24 | 'o'<<16 | 'c'<<8 | 'l' // 'kocl'
	KeyAEOffStyles      KeyAE = 'o'<<24 | 'f'<<16 | 's'<<8 | 't' // 'ofst'
	KeyAEOffset         KeyAE = 'o'<<24 | 'f'<<16 | 's'<<8 | 't' // 'ofst'
	KeyAEOnStyles       KeyAE = 'o'<<24 | 'n'<<16 | 's'<<8 | 't' // 'onst'
	KeyAEPMTable        KeyAE = 'k'<<24 | 'p'<<16 | 'm'<<8 | 't' // 'kpmt'
	KeyAEParamFlags     KeyAE = 'p'<<24 | 'm'<<16 | 'f'<<8 | 'g' // 'pmfg'
	KeyAEParameters     KeyAE = 'p'<<24 | 'r'<<16 | 'm'<<8 | 's' // 'prms'
	KeyAEPenColor       KeyAE = 'p'<<24 | 'p'<<16 | 'c'<<8 | 'l' // 'ppcl'
	KeyAEPenPattern     KeyAE = 'p'<<24 | 'p'<<16 | 'p'<<8 | 'a' // 'pppa'
	KeyAEPenWidth       KeyAE = 'p'<<24 | 'p'<<16 | 'w'<<8 | 'd' // 'ppwd'
	KeyAEPinRange       KeyAE = 'p'<<24 | 'n'<<16 | 'r'<<8 | 'g' // 'pnrg'
	KeyAEPixMapMinus    KeyAE = 'k'<<24 | 'p'<<16 | 'm'<<8 | 'm' // 'kpmm'
	KeyAEPixelDepth     KeyAE = 'p'<<24 | 'd'<<16 | 'p'<<8 | 't' // 'pdpt'
	KeyAEPoint          KeyAE = 'g'<<24 | 'p'<<16 | 'o'<<8 | 's' // 'gpos'
	KeyAEPointList      KeyAE = 'p'<<24 | 't'<<16 | 'l'<<8 | 't' // 'ptlt'
	KeyAEPointSize      KeyAE = 'p'<<24 | 't'<<16 | 's'<<8 | 'z' // 'ptsz'
	KeyAEPosition       KeyAE = 'k'<<24 | 'p'<<16 | 'o'<<8 | 's' // 'kpos'
	KeyAEPropData       KeyAE = 'p'<<24 | 'r'<<16 | 'd'<<8 | 't' // 'prdt'
	KeyAEPropFlags      KeyAE = 'p'<<24 | 'r'<<16 | 'f'<<8 | 'g' // 'prfg'
	KeyAEPropID         KeyAE = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
	KeyAEProperties     KeyAE = 'q'<<24 | 'p'<<16 | 'r'<<8 | 'o' // 'qpro'
	KeyAEProperty       KeyAE = 'k'<<24 | 'p'<<16 | 'r'<<8 | 'p' // 'kprp'
	KeyAEProtection     KeyAE = 'p'<<24 | 'p'<<16 | 'r'<<8 | 'o' // 'ppro'
	KeyAERegionClass    KeyAE = 'r'<<24 | 'g'<<16 | 'n'<<8 | 'c' // 'rgnc'
	KeyAERenderAs       KeyAE = 'k'<<24 | 'r'<<16 | 'e'<<8 | 'n' // 'kren'
	KeyAERequestedType  KeyAE = 'r'<<24 | 't'<<16 | 'y'<<8 | 'p' // 'rtyp'
	KeyAEResult         KeyAE = '-'<<24 | '-'<<16 | '-'<<8 | '-' // '----'
	KeyAEResultInfo     KeyAE = 'r'<<24 | 's'<<16 | 'i'<<8 | 'n' // 'rsin'
	KeyAERotPoint       KeyAE = 'k'<<24 | 'r'<<16 | 't'<<8 | 'p' // 'krtp'
	KeyAERotation       KeyAE = 'p'<<24 | 'r'<<16 | 'o'<<8 | 't' // 'prot'
	KeyAERowList        KeyAE = 'k'<<24 | 'r'<<16 | 'l'<<8 | 's' // 'krls'
	KeyAESaveOptions    KeyAE = 's'<<24 | 'a'<<16 | 'v'<<8 | 'o' // 'savo'
	KeyAEScale          KeyAE = 'p'<<24 | 's'<<16 | 'c'<<8 | 'l' // 'pscl'
	KeyAEScriptTag      KeyAE = 'p'<<24 | 's'<<16 | 'c'<<8 | 't' // 'psct'
	// KeyAESearchText: # Discussion
	KeyAESearchText      KeyAE = 's'<<24 | 't'<<16 | 'x'<<8 | 't' // 'stxt'
	KeyAEShowWhere       KeyAE = 's'<<24 | 'h'<<16 | 'o'<<8 | 'w' // 'show'
	KeyAEStartAngle      KeyAE = 'p'<<24 | 'a'<<16 | 'n'<<8 | 'g' // 'pang'
	KeyAEStartPoint      KeyAE = 'p'<<24 | 's'<<16 | 't'<<8 | 'p' // 'pstp'
	KeyAEStyles          KeyAE = 'k'<<24 | 's'<<16 | 't'<<8 | 'y' // 'ksty'
	KeyAESuiteID         KeyAE = 's'<<24 | 'u'<<16 | 'i'<<8 | 't' // 'suit'
	KeyAEText            KeyAE = 'k'<<24 | 't'<<16 | 'x'<<8 | 't' // 'ktxt'
	KeyAETextColor       KeyAE = 'p'<<24 | 't'<<16 | 'x'<<8 | 'c' // 'ptxc'
	KeyAETextFont        KeyAE = 'p'<<24 | 't'<<16 | 'x'<<8 | 'f' // 'ptxf'
	KeyAETextLineAscent  KeyAE = 'k'<<24 | 't'<<16 | 'a'<<8 | 's' // 'ktas'
	KeyAETextLineHeight  KeyAE = 'k'<<24 | 't'<<16 | 'l'<<8 | 'h' // 'ktlh'
	KeyAETextPointSize   KeyAE = 'p'<<24 | 't'<<16 | 'p'<<8 | 's' // 'ptps'
	KeyAETextStyles      KeyAE = 't'<<24 | 'x'<<16 | 's'<<8 | 't' // 'txst'
	KeyAETheText         KeyAE = 't'<<24 | 'h'<<16 | 't'<<8 | 'x' // 'thtx'
	KeyAETransferMode    KeyAE = 'p'<<24 | 'p'<<16 | 't'<<8 | 'm' // 'pptm'
	KeyAETranslation     KeyAE = 'p'<<24 | 't'<<16 | 'r'<<8 | 's' // 'ptrs'
	KeyAETryAsStructGraf KeyAE = 't'<<24 | 'o'<<16 | 'o'<<8 | 'g' // 'toog'
	KeyAEUniformStyles   KeyAE = 'u'<<24 | 's'<<16 | 't'<<8 | 'l' // 'ustl'
	KeyAEUpdateOn        KeyAE = 'p'<<24 | 'u'<<16 | 'p'<<8 | 'd' // 'pupd'
	KeyAEUserTerm        KeyAE = 'u'<<24 | 't'<<16 | 'r'<<8 | 'm' // 'utrm'
	KeyAEWindow          KeyAE = 'w'<<24 | 'n'<<16 | 'd'<<8 | 'w' // 'wndw'
	KeyAEWritingCode     KeyAE = 'w'<<24 | 'r'<<16 | 'c'<<8 | 'd' // 'wrcd'
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

type KeyAERestoreApp uint

const (
	KeyAERestoreAppState KeyAERestoreApp = 'r'<<24 | 's'<<16 | 't'<<8 | 'o' // 'rsto'
)

func (e KeyAERestoreApp) String() string {
	switch e {
	case KeyAERestoreAppState:
		return "KeyAERestoreAppState"
	default:
		return fmt.Sprintf("KeyAERestoreApp(%d)", e)
	}
}

type KeyReplyPort uint

const (
	KeyReplyPortAttr KeyReplyPort = 'r'<<24 | 'e'<<16 | 'p'<<8 | 'p' // 'repp'
)

func (e KeyReplyPort) String() string {
	switch e {
	case KeyReplyPortAttr:
		return "KeyReplyPortAttr"
	default:
		return fmt.Sprintf("KeyReplyPort(%d)", e)
	}
}

type KeySOAP uint

const (
	KeySOAPSMDNamespace      KeySOAP = 's'<<24 | 's'<<16 | 'n'<<8 | 's' // 'ssns'
	KeySOAPSMDNamespaceURI   KeySOAP = 's'<<24 | 's'<<16 | 'n'<<8 | 'u' // 'ssnu'
	KeySOAPSMDType           KeySOAP = 's'<<24 | 's'<<16 | 't'<<8 | 'p' // 'sstp'
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
	KAERPCClass                  KeyUserNameAttr = 'r'<<24 | 'p'<<16 | 'c'<<8 | ' ' // 'rpc '
	KAESOAPScheme                KeyUserNameAttr = 'S'<<24 | 'O'<<16 | 'A'<<8 | 'P' // 'SOAP'
	KAESharedScriptHandler       KeyUserNameAttr = 'w'<<24 | 's'<<16 | 'c'<<8 | 'p' // 'wscp'
	KAEXMLRPCScheme              KeyUserNameAttr = 'R'<<24 | 'P'<<16 | 'C'<<8 | '2' // 'RPC2'
	KeyAEPOSTHeaderData          KeyUserNameAttr = 'p'<<24 | 'h'<<16 | 'e'<<8 | 'd' // 'phed'
	KeyAEReplyHeaderData         KeyUserNameAttr = 'r'<<24 | 'h'<<16 | 'e'<<8 | 'd' // 'rhed'
	KeyAEXMLReplyData            KeyUserNameAttr = 'x'<<24 | 'r'<<16 | 'e'<<8 | 'p' // 'xrep'
	KeyAEXMLRequestData          KeyUserNameAttr = 'x'<<24 | 'r'<<16 | 'e'<<8 | 'q' // 'xreq'
	KeyAdditionalHTTPHeaders     KeyUserNameAttr = 'a'<<24 | 'h'<<16 | 'e'<<8 | 'd' // 'ahed'
	KeyDisableAuthenticationAttr KeyUserNameAttr = 'a'<<24 | 'u'<<16 | 't'<<8 | 'h' // 'auth'
	KeyRPCMethodName             KeyUserNameAttr = 'm'<<24 | 'e'<<16 | 't'<<8 | 'h' // 'meth'
	KeyRPCMethodParam            KeyUserNameAttr = 'p'<<24 | 'a'<<16 | 'r'<<8 | 'm' // 'parm'
	KeyRPCMethodParamOrder       KeyUserNameAttr = '/'<<24 | 'o'<<16 | 'r'<<8 | 'd' // '/ord'
	KeySOAPAction                KeyUserNameAttr = 's'<<24 | 'a'<<16 | 'c'<<8 | 't' // 'sact'
	KeySOAPMethodNameSpace       KeyUserNameAttr = 'm'<<24 | 's'<<16 | 'p'<<8 | 'c' // 'mspc'
	KeySOAPMethodNameSpaceURI    KeyUserNameAttr = 'm'<<24 | 's'<<16 | 'p'<<8 | 'u' // 'mspu'
	KeySOAPSchemaVersion         KeyUserNameAttr = 's'<<24 | 's'<<16 | 'c'<<8 | 'h' // 'ssch'
	KeyUserNameAttrValue         KeyUserNameAttr = 'u'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'unam'
	KeyUserPasswordAttr          KeyUserNameAttr = 'p'<<24 | 'a'<<16 | 's'<<8 | 's' // 'pass'
	KeyXMLDebuggingAttr          KeyUserNameAttr = 'x'<<24 | 'd'<<16 | 'b'<<8 | 'g' // 'xdbg'
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

type KioACAccessOwnerBit uint

const (
	KfullPrivileges               KioACAccessOwnerBit = 458759
	KioACAccessBlankAccessBit     KioACAccessOwnerBit = 28
	KioACAccessBlankAccessMask    KioACAccessOwnerBit = 268435456
	KioACAccessEveryoneReadBit    KioACAccessOwnerBit = 17
	KioACAccessEveryoneReadMask   KioACAccessOwnerBit = 131072
	KioACAccessEveryoneSearchBit  KioACAccessOwnerBit = 16
	KioACAccessEveryoneSearchMask KioACAccessOwnerBit = 65536
	KioACAccessEveryoneWriteBit   KioACAccessOwnerBit = 18
	KioACAccessEveryoneWriteMask  KioACAccessOwnerBit = 262144
	KioACAccessGroupReadBit       KioACAccessOwnerBit = 9
	KioACAccessGroupReadMask      KioACAccessOwnerBit = 512
	KioACAccessGroupSearchBit     KioACAccessOwnerBit = 8
	KioACAccessGroupSearchMask    KioACAccessOwnerBit = 256
	KioACAccessGroupWriteBit      KioACAccessOwnerBit = 10
	KioACAccessGroupWriteMask     KioACAccessOwnerBit = 1024
	KioACAccessOwnerBitValue      KioACAccessOwnerBit = 31
	KioACAccessOwnerMask          KioACAccessOwnerBit = 0
	KioACAccessOwnerReadBit       KioACAccessOwnerBit = 1
	KioACAccessOwnerReadMask      KioACAccessOwnerBit = 2
	KioACAccessOwnerSearchBit     KioACAccessOwnerBit = 0
	KioACAccessOwnerSearchMask    KioACAccessOwnerBit = 1
	KioACAccessOwnerWriteBit      KioACAccessOwnerBit = 2
	KioACAccessOwnerWriteMask     KioACAccessOwnerBit = 4
	KioACAccessUserReadBit        KioACAccessOwnerBit = 25
	KioACAccessUserReadMask       KioACAccessOwnerBit = 33554432
	KioACAccessUserSearchBit      KioACAccessOwnerBit = 24
	KioACAccessUserSearchMask     KioACAccessOwnerBit = 16777216
	KioACAccessUserWriteBit       KioACAccessOwnerBit = 26
	KioACAccessUserWriteMask      KioACAccessOwnerBit = 67108864
	KownerPrivileges              KioACAccessOwnerBit = 7
)

func (e KioACAccessOwnerBit) String() string {
	switch e {
	case KfullPrivileges:
		return "KfullPrivileges"
	case KioACAccessBlankAccessBit:
		return "KioACAccessBlankAccessBit"
	case KioACAccessBlankAccessMask:
		return "KioACAccessBlankAccessMask"
	case KioACAccessEveryoneReadBit:
		return "KioACAccessEveryoneReadBit"
	case KioACAccessEveryoneReadMask:
		return "KioACAccessEveryoneReadMask"
	case KioACAccessEveryoneSearchBit:
		return "KioACAccessEveryoneSearchBit"
	case KioACAccessEveryoneSearchMask:
		return "KioACAccessEveryoneSearchMask"
	case KioACAccessEveryoneWriteBit:
		return "KioACAccessEveryoneWriteBit"
	case KioACAccessEveryoneWriteMask:
		return "KioACAccessEveryoneWriteMask"
	case KioACAccessGroupReadBit:
		return "KioACAccessGroupReadBit"
	case KioACAccessGroupReadMask:
		return "KioACAccessGroupReadMask"
	case KioACAccessGroupSearchBit:
		return "KioACAccessGroupSearchBit"
	case KioACAccessGroupSearchMask:
		return "KioACAccessGroupSearchMask"
	case KioACAccessGroupWriteBit:
		return "KioACAccessGroupWriteBit"
	case KioACAccessGroupWriteMask:
		return "KioACAccessGroupWriteMask"
	case KioACAccessOwnerBitValue:
		return "KioACAccessOwnerBitValue"
	case KioACAccessOwnerMask:
		return "KioACAccessOwnerMask"
	case KioACAccessOwnerReadBit:
		return "KioACAccessOwnerReadBit"
	case KioACAccessOwnerReadMask:
		return "KioACAccessOwnerReadMask"
	case KioACAccessOwnerWriteMask:
		return "KioACAccessOwnerWriteMask"
	case KioACAccessUserReadBit:
		return "KioACAccessUserReadBit"
	case KioACAccessUserReadMask:
		return "KioACAccessUserReadMask"
	case KioACAccessUserSearchBit:
		return "KioACAccessUserSearchBit"
	case KioACAccessUserSearchMask:
		return "KioACAccessUserSearchMask"
	case KioACAccessUserWriteBit:
		return "KioACAccessUserWriteBit"
	case KioACAccessUserWriteMask:
		return "KioACAccessUserWriteMask"
	case KownerPrivileges:
		return "KownerPrivileges"
	default:
		return fmt.Sprintf("KioACAccessOwnerBit(%d)", e)
	}
}

type KioACUser uint

const (
	KioACUserNoMakeChangesBit  KioACUser = 2
	KioACUserNoMakeChangesMask KioACUser = 4
	KioACUserNoSeeFilesBit     KioACUser = 1
	KioACUserNoSeeFilesMask    KioACUser = 2
	KioACUserNoSeeFolderBit    KioACUser = 0
	KioACUserNoSeeFolderMask   KioACUser = 1
	KioACUserNotOwnerBit       KioACUser = 7
	KioACUserNotOwnerMask      KioACUser = 128
)

func (e KioACUser) String() string {
	switch e {
	case KioACUserNoMakeChangesBit:
		return "KioACUserNoMakeChangesBit"
	case KioACUserNoMakeChangesMask:
		return "KioACUserNoMakeChangesMask"
	case KioACUserNoSeeFilesBit:
		return "KioACUserNoSeeFilesBit"
	case KioACUserNoSeeFolderBit:
		return "KioACUserNoSeeFolderBit"
	case KioACUserNotOwnerBit:
		return "KioACUserNotOwnerBit"
	case KioACUserNotOwnerMask:
		return "KioACUserNotOwnerMask"
	default:
		return fmt.Sprintf("KioACUser(%d)", e)
	}
}

type KioFCB uint

const (
	KioFCBFileLockedBit   KioFCB = 13
	KioFCBFileLockedMask  KioFCB = 8192
	KioFCBLargeFileBit    KioFCB = 11
	KioFCBLargeFileMask   KioFCB = 2048
	KioFCBModifiedBit     KioFCB = 15
	KioFCBModifiedMask    KioFCB = 32768
	KioFCBOwnClumpBit     KioFCB = 14
	KioFCBOwnClumpMask    KioFCB = 16384
	KioFCBResourceBit     KioFCB = 9
	KioFCBResourceMask    KioFCB = 512
	KioFCBSharedWriteBit  KioFCB = 12
	KioFCBSharedWriteMask KioFCB = 4096
	KioFCBWriteBit        KioFCB = 8
	KioFCBWriteLockedBit  KioFCB = 10
	KioFCBWriteLockedMask KioFCB = 1024
	KioFCBWriteMask       KioFCB = 256
)

func (e KioFCB) String() string {
	switch e {
	case KioFCBFileLockedBit:
		return "KioFCBFileLockedBit"
	case KioFCBFileLockedMask:
		return "KioFCBFileLockedMask"
	case KioFCBLargeFileBit:
		return "KioFCBLargeFileBit"
	case KioFCBLargeFileMask:
		return "KioFCBLargeFileMask"
	case KioFCBModifiedBit:
		return "KioFCBModifiedBit"
	case KioFCBModifiedMask:
		return "KioFCBModifiedMask"
	case KioFCBOwnClumpBit:
		return "KioFCBOwnClumpBit"
	case KioFCBOwnClumpMask:
		return "KioFCBOwnClumpMask"
	case KioFCBResourceBit:
		return "KioFCBResourceBit"
	case KioFCBResourceMask:
		return "KioFCBResourceMask"
	case KioFCBSharedWriteBit:
		return "KioFCBSharedWriteBit"
	case KioFCBSharedWriteMask:
		return "KioFCBSharedWriteMask"
	case KioFCBWriteBit:
		return "KioFCBWriteBit"
	case KioFCBWriteLockedBit:
		return "KioFCBWriteLockedBit"
	case KioFCBWriteLockedMask:
		return "KioFCBWriteLockedMask"
	case KioFCBWriteMask:
		return "KioFCBWriteMask"
	default:
		return fmt.Sprintf("KioFCB(%d)", e)
	}
}

type KioFlAttribLockedBit uint

const (
	IoDirFlg                  KioFlAttribLockedBit = 4
	IoDirMask                 KioFlAttribLockedBit = 16
	KioFlAttribCopyProtBit    KioFlAttribLockedBit = 6
	KioFlAttribCopyProtMask   KioFlAttribLockedBit = 64
	KioFlAttribDataOpenBit    KioFlAttribLockedBit = 3
	KioFlAttribDataOpenMask   KioFlAttribLockedBit = 8
	KioFlAttribDirBit         KioFlAttribLockedBit = 4
	KioFlAttribDirMask        KioFlAttribLockedBit = 16
	KioFlAttribFileOpenBit    KioFlAttribLockedBit = 7
	KioFlAttribFileOpenMask   KioFlAttribLockedBit = 128
	KioFlAttribInSharedBit    KioFlAttribLockedBit = 2
	KioFlAttribInSharedMask   KioFlAttribLockedBit = 4
	KioFlAttribLockedBitValue KioFlAttribLockedBit = 0
	KioFlAttribLockedMask     KioFlAttribLockedBit = 1
	KioFlAttribMountedBit     KioFlAttribLockedBit = 3
	KioFlAttribMountedMask    KioFlAttribLockedBit = 8
	KioFlAttribResOpenBit     KioFlAttribLockedBit = 2
	KioFlAttribResOpenMask    KioFlAttribLockedBit = 4
	KioFlAttribSharePointBit  KioFlAttribLockedBit = 5
	KioFlAttribSharePointMask KioFlAttribLockedBit = 32
)

func (e KioFlAttribLockedBit) String() string {
	switch e {
	case IoDirFlg:
		return "IoDirFlg"
	case IoDirMask:
		return "IoDirMask"
	case KioFlAttribCopyProtBit:
		return "KioFlAttribCopyProtBit"
	case KioFlAttribCopyProtMask:
		return "KioFlAttribCopyProtMask"
	case KioFlAttribDataOpenBit:
		return "KioFlAttribDataOpenBit"
	case KioFlAttribDataOpenMask:
		return "KioFlAttribDataOpenMask"
	case KioFlAttribFileOpenBit:
		return "KioFlAttribFileOpenBit"
	case KioFlAttribFileOpenMask:
		return "KioFlAttribFileOpenMask"
	case KioFlAttribInSharedBit:
		return "KioFlAttribInSharedBit"
	case KioFlAttribLockedBitValue:
		return "KioFlAttribLockedBitValue"
	case KioFlAttribLockedMask:
		return "KioFlAttribLockedMask"
	case KioFlAttribSharePointBit:
		return "KioFlAttribSharePointBit"
	case KioFlAttribSharePointMask:
		return "KioFlAttribSharePointMask"
	default:
		return fmt.Sprintf("KioFlAttribLockedBit(%d)", e)
	}
}

type KioVAtrb uint

const (
	KioVAtrbDefaultVolumeBit   KioVAtrb = 5
	KioVAtrbDefaultVolumeMask  KioVAtrb = 32
	KioVAtrbFilesOpenBit       KioVAtrb = 6
	KioVAtrbFilesOpenMask      KioVAtrb = 64
	KioVAtrbHardwareLockedBit  KioVAtrb = 7
	KioVAtrbHardwareLockedMask KioVAtrb = 128
	KioVAtrbSoftwareLockedBit  KioVAtrb = 15
	KioVAtrbSoftwareLockedMask KioVAtrb = 32768
)

func (e KioVAtrb) String() string {
	switch e {
	case KioVAtrbDefaultVolumeBit:
		return "KioVAtrbDefaultVolumeBit"
	case KioVAtrbDefaultVolumeMask:
		return "KioVAtrbDefaultVolumeMask"
	case KioVAtrbFilesOpenBit:
		return "KioVAtrbFilesOpenBit"
	case KioVAtrbFilesOpenMask:
		return "KioVAtrbFilesOpenMask"
	case KioVAtrbHardwareLockedBit:
		return "KioVAtrbHardwareLockedBit"
	case KioVAtrbHardwareLockedMask:
		return "KioVAtrbHardwareLockedMask"
	case KioVAtrbSoftwareLockedBit:
		return "KioVAtrbSoftwareLockedBit"
	case KioVAtrbSoftwareLockedMask:
		return "KioVAtrbSoftwareLockedMask"
	default:
		return fmt.Sprintf("KioVAtrb(%d)", e)
	}
}

type Kno uint

const (
	KnoGroup Kno = 0
)

func (e Kno) String() string {
	switch e {
	case KnoGroup:
		return "KnoGroup"
	default:
		return fmt.Sprintf("Kno(%d)", e)
	}
}

type KnoUser uint

const (
	KadministratorUser KnoUser = 1
	KnoUserValue       KnoUser = 0
)

func (e KnoUser) String() string {
	switch e {
	case KadministratorUser:
		return "KadministratorUser"
	case KnoUserValue:
		return "KnoUserValue"
	default:
		return fmt.Sprintf("KnoUser(%d)", e)
	}
}

type Ktextencodingiso2022 uint

const (
	KTextEncodingISO_2022_CN     Ktextencodingiso2022 = 2096
	KTextEncodingISO_2022_CN_EXT Ktextencodingiso2022 = 2097
	// KTextEncodingISO_2022_JP: See RFC 1468.
	KTextEncodingISO_2022_JP Ktextencodingiso2022 = 2080
	// KTextEncodingISO_2022_JP_1: See RFC 2237.
	KTextEncodingISO_2022_JP_1 Ktextencodingiso2022 = 2082
	// KTextEncodingISO_2022_JP_2: See RFC 1554.
	KTextEncodingISO_2022_JP_2 Ktextencodingiso2022 = 2081
	// KTextEncodingISO_2022_JP_3: JIS X0213
	KTextEncodingISO_2022_JP_3 Ktextencodingiso2022 = 2083
	KTextEncodingISO_2022_KR   Ktextencodingiso2022 = 2112
)

func (e Ktextencodingiso2022) String() string {
	switch e {
	case KTextEncodingISO_2022_CN:
		return "KTextEncodingISO_2022_CN"
	case KTextEncodingISO_2022_CN_EXT:
		return "KTextEncodingISO_2022_CN_EXT"
	case KTextEncodingISO_2022_JP:
		return "KTextEncodingISO_2022_JP"
	case KTextEncodingISO_2022_JP_1:
		return "KTextEncodingISO_2022_JP_1"
	case KTextEncodingISO_2022_JP_2:
		return "KTextEncodingISO_2022_JP_2"
	case KTextEncodingISO_2022_JP_3:
		return "KTextEncodingISO_2022_JP_3"
	case KTextEncodingISO_2022_KR:
		return "KTextEncodingISO_2022_KR"
	default:
		return fmt.Sprintf("Ktextencodingiso2022(%d)", e)
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
	LaDictionaryTooManyErr   La = -6994
	LaDictionaryUnknownErr   La = -6993
	LaEngineNotFoundErr      La = -7000
	LaEnvironmentBusyErr     La = -6985
	LaEnvironmentExistErr    La = -6987
	LaEnvironmentNotFoundErr La = -6986
	LaFailAnalysisErr        La = -6990
	LaInvalidPathErr         La = -6988
	LaNoMoreMorphemeErr      La = -6989
	LaPropertyErr            La = -6999
	LaPropertyIsReadOnlyErr  La = -6997
	LaPropertyNotFoundErr    La = -6998
	LaPropertyUnknownErr     La = -6996
	LaPropertyValueErr       La = -6995
	LaTextOverFlowErr        La = -6991
	LaTooSmallBufferErr      La = -6984
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

type Lang uint

const (
	// Deprecated.
	LangAfricaans Lang = 141
	// Deprecated.
	LangAfrikaans Lang = 141
	// Deprecated.
	LangAlbanian Lang = 36
	// Deprecated.
	LangAmharic Lang = 85
	// Deprecated.
	LangArabic Lang = 12
	// Deprecated.
	LangArmenian Lang = 51
	// Deprecated.
	LangAssamese Lang = 68
	// Deprecated.
	LangAymara Lang = 134
	// Deprecated.
	LangAzerbaijanAr Lang = 50
	// Deprecated.
	LangAzerbaijanRoman Lang = 150
	// Deprecated.
	LangAzerbaijani Lang = 49
	// Deprecated.
	LangBasque Lang = 129
	// Deprecated.
	LangBelorussian Lang = 46
	// Deprecated.
	LangBengali Lang = 67
	// Deprecated.
	LangBreton Lang = 142
	// Deprecated.
	LangBulgarian Lang = 44
	// Deprecated.
	LangBurmese Lang = 77
	// Deprecated.
	LangByelorussian Lang = 46
	// Deprecated.
	LangCatalan Lang = 130
	// Deprecated.
	LangChewa Lang = 92
	// Deprecated.
	LangChinese Lang = 19
	// Deprecated.
	LangCroatian Lang = 18
	// Deprecated.
	LangCzech Lang = 38
	// Deprecated.
	LangDanish Lang = 7
	// Deprecated.
	LangDutch Lang = 4
	// Deprecated.
	LangDzongkha Lang = 137
	// Deprecated.
	LangEnglish Lang = 0
	// Deprecated.
	LangEsperanto Lang = 94
	// Deprecated.
	LangEstonian Lang = 27
	// Deprecated.
	LangFaeroese Lang = 30
	// Deprecated.
	LangFaroese Lang = 30
	// Deprecated.
	LangFarsi Lang = 31
	// Deprecated.
	LangFinnish Lang = 13
	// Deprecated.
	LangFlemish Lang = 34
	// Deprecated.
	LangFrench Lang = 1
	// Deprecated.
	LangGalician Lang = 140
	// Deprecated.
	LangGalla Lang = 87
	// Deprecated.
	LangGeorgian Lang = 52
	// Deprecated.
	LangGerman Lang = 2
	// Deprecated.
	LangGreek Lang = 14
	// Deprecated.
	LangGreekAncient Lang = 148
	// Deprecated.
	LangGreekPoly Lang = 148
	// Deprecated.
	LangGreenlandic Lang = 149
	// Deprecated.
	LangGuarani Lang = 133
	// Deprecated.
	LangGujarati Lang = 69
	// Deprecated.
	LangHebrew Lang = 10
	// Deprecated.
	LangHindi Lang = 21
	// Deprecated.
	LangHungarian Lang = 26
	// Deprecated.
	LangIcelandic Lang = 15
	// Deprecated.
	LangIndonesian Lang = 81
	// Deprecated.
	LangInuktitut Lang = 143
	// Deprecated.
	LangIrish Lang = 35
	// Deprecated.
	LangIrishGaelic Lang = 35
	// Deprecated.
	LangIrishGaelicScript Lang = 146
	// Deprecated.
	LangItalian Lang = 3
	// Deprecated.
	LangJapanese Lang = 11
	// Deprecated.
	LangJavaneseRom Lang = 138
	// Deprecated.
	LangKannada Lang = 73
	// Deprecated.
	LangKashmiri Lang = 61
	// Deprecated.
	LangKazakh Lang = 48
	// Deprecated.
	LangKhmer Lang = 78
	// Deprecated.
	LangKinyarwanda Lang = 90
	// Deprecated.
	LangKirghiz Lang = 54
	// Deprecated.
	LangKorean Lang = 23
	// Deprecated.
	LangKurdish Lang = 60
	// Deprecated.
	LangLao Lang = 79
	// Deprecated.
	LangLappish Lang = 29
	// Deprecated.
	LangLapponian Lang = 29
	// Deprecated.
	LangLatin Lang = 131
	// Deprecated.
	LangLatvian Lang = 28
	// Deprecated.
	LangLettish Lang = 28
	// Deprecated.
	LangLithuanian Lang = 24
	// Deprecated.
	LangMacedonian Lang = 43
	// Deprecated.
	LangMalagasy Lang = 93
	// Deprecated.
	LangMalayArabic Lang = 84
	// Deprecated.
	LangMalayRoman Lang = 83
	// Deprecated.
	LangMalayalam Lang = 72
	// Deprecated.
	LangMalta Lang = 16
	// Deprecated.
	LangMaltese Lang = 16
	// Deprecated.
	LangManxGaelic Lang = 145
	// Deprecated.
	LangMarathi Lang = 66
	// Deprecated.
	LangMoldavian Lang = 53
	// Deprecated.
	LangMongolian Lang = 57
	// Deprecated.
	LangMongolianCyr Lang = 58
	// Deprecated.
	LangNepali Lang = 64
	// Deprecated.
	LangNorwegian Lang = 9
	// Deprecated.
	LangNyanja Lang = 92
	// Deprecated.
	LangNynorsk Lang = 151
	// Deprecated.
	LangOriya Lang = 71
	// Deprecated.
	LangOromo Lang = 87
	// Deprecated.
	LangPashto Lang = 59
	// Deprecated.
	LangPersian Lang = 31
	// Deprecated.
	LangPolish Lang = 25
	// Deprecated.
	LangPortugese Lang = 8
	// Deprecated.
	LangPortuguese Lang = 8
	// Deprecated.
	LangPunjabi Lang = 70
	// Deprecated.
	LangQuechua Lang = 132
	// Deprecated.
	LangRomanian Lang = 37
	// Deprecated.
	LangRuanda Lang = 90
	// Deprecated.
	LangRundi Lang = 91
	// Deprecated.
	LangRussian Lang = 32
	// Deprecated.
	LangSaamisk Lang = 29
	// Deprecated.
	LangSami Lang = 29
	// Deprecated.
	LangSanskrit Lang = 65
	// Deprecated.
	LangScottishGaelic Lang = 144
	// Deprecated.
	LangSerbian Lang = 42
	// Deprecated.
	LangSimpChinese Lang = 33
	// Deprecated.
	LangSindhi Lang = 62
	// Deprecated.
	LangSinhalese Lang = 76
	// Deprecated.
	LangSlovak Lang = 39
	// Deprecated.
	LangSlovenian Lang = 40
	// Deprecated.
	LangSomali Lang = 88
	// Deprecated.
	LangSpanish Lang = 6
	// Deprecated.
	LangSundaneseRom Lang = 139
	// Deprecated.
	LangSwahili Lang = 89
	// Deprecated.
	LangSwedish Lang = 5
	// Deprecated.
	LangTagalog Lang = 82
	// Deprecated.
	LangTajiki Lang = 55
	// Deprecated.
	LangTamil Lang = 74
	// Deprecated.
	LangTatar Lang = 135
	// Deprecated.
	LangTelugu Lang = 75
	// Deprecated.
	LangThai Lang = 22
	// Deprecated.
	LangTibetan Lang = 63
	// Deprecated.
	LangTigrinya Lang = 86
	// Deprecated.
	LangTongan Lang = 147
	// Deprecated.
	LangTradChinese Lang = 19
	// Deprecated.
	LangTurkish Lang = 17
	// Deprecated.
	LangTurkmen Lang = 56
	// Deprecated.
	LangUighur Lang = 136
	// Deprecated.
	LangUkrainian Lang = 45
	// Deprecated.
	LangUnspecified Lang = 32767
	// Deprecated.
	LangUrdu Lang = 20
	// Deprecated.
	LangUzbek Lang = 47
	// Deprecated.
	LangVietnamese Lang = 80
	// Deprecated.
	LangWelsh Lang = 128
	// Deprecated.
	LangYiddish Lang = 41
	// Deprecated.
	LangYugoslavian Lang = 18
)

func (e Lang) String() string {
	switch e {
	case LangAfricaans:
		return "LangAfricaans"
	case LangAlbanian:
		return "LangAlbanian"
	case LangAmharic:
		return "LangAmharic"
	case LangArabic:
		return "LangArabic"
	case LangArmenian:
		return "LangArmenian"
	case LangAssamese:
		return "LangAssamese"
	case LangAymara:
		return "LangAymara"
	case LangAzerbaijanAr:
		return "LangAzerbaijanAr"
	case LangAzerbaijanRoman:
		return "LangAzerbaijanRoman"
	case LangAzerbaijani:
		return "LangAzerbaijani"
	case LangBasque:
		return "LangBasque"
	case LangBelorussian:
		return "LangBelorussian"
	case LangBengali:
		return "LangBengali"
	case LangBreton:
		return "LangBreton"
	case LangBulgarian:
		return "LangBulgarian"
	case LangBurmese:
		return "LangBurmese"
	case LangCatalan:
		return "LangCatalan"
	case LangChewa:
		return "LangChewa"
	case LangChinese:
		return "LangChinese"
	case LangCroatian:
		return "LangCroatian"
	case LangCzech:
		return "LangCzech"
	case LangDanish:
		return "LangDanish"
	case LangDutch:
		return "LangDutch"
	case LangDzongkha:
		return "LangDzongkha"
	case LangEnglish:
		return "LangEnglish"
	case LangEsperanto:
		return "LangEsperanto"
	case LangEstonian:
		return "LangEstonian"
	case LangFaeroese:
		return "LangFaeroese"
	case LangFarsi:
		return "LangFarsi"
	case LangFinnish:
		return "LangFinnish"
	case LangFlemish:
		return "LangFlemish"
	case LangFrench:
		return "LangFrench"
	case LangGalician:
		return "LangGalician"
	case LangGalla:
		return "LangGalla"
	case LangGeorgian:
		return "LangGeorgian"
	case LangGerman:
		return "LangGerman"
	case LangGreek:
		return "LangGreek"
	case LangGreekAncient:
		return "LangGreekAncient"
	case LangGreenlandic:
		return "LangGreenlandic"
	case LangGuarani:
		return "LangGuarani"
	case LangGujarati:
		return "LangGujarati"
	case LangHebrew:
		return "LangHebrew"
	case LangHindi:
		return "LangHindi"
	case LangHungarian:
		return "LangHungarian"
	case LangIcelandic:
		return "LangIcelandic"
	case LangIndonesian:
		return "LangIndonesian"
	case LangInuktitut:
		return "LangInuktitut"
	case LangIrish:
		return "LangIrish"
	case LangIrishGaelicScript:
		return "LangIrishGaelicScript"
	case LangItalian:
		return "LangItalian"
	case LangJapanese:
		return "LangJapanese"
	case LangJavaneseRom:
		return "LangJavaneseRom"
	case LangKannada:
		return "LangKannada"
	case LangKashmiri:
		return "LangKashmiri"
	case LangKazakh:
		return "LangKazakh"
	case LangKhmer:
		return "LangKhmer"
	case LangKinyarwanda:
		return "LangKinyarwanda"
	case LangKirghiz:
		return "LangKirghiz"
	case LangKorean:
		return "LangKorean"
	case LangKurdish:
		return "LangKurdish"
	case LangLao:
		return "LangLao"
	case LangLappish:
		return "LangLappish"
	case LangLatin:
		return "LangLatin"
	case LangLatvian:
		return "LangLatvian"
	case LangLithuanian:
		return "LangLithuanian"
	case LangMacedonian:
		return "LangMacedonian"
	case LangMalagasy:
		return "LangMalagasy"
	case LangMalayArabic:
		return "LangMalayArabic"
	case LangMalayRoman:
		return "LangMalayRoman"
	case LangMalayalam:
		return "LangMalayalam"
	case LangMalta:
		return "LangMalta"
	case LangManxGaelic:
		return "LangManxGaelic"
	case LangMarathi:
		return "LangMarathi"
	case LangMoldavian:
		return "LangMoldavian"
	case LangMongolian:
		return "LangMongolian"
	case LangMongolianCyr:
		return "LangMongolianCyr"
	case LangNepali:
		return "LangNepali"
	case LangNorwegian:
		return "LangNorwegian"
	case LangNynorsk:
		return "LangNynorsk"
	case LangOriya:
		return "LangOriya"
	case LangPashto:
		return "LangPashto"
	case LangPolish:
		return "LangPolish"
	case LangPortugese:
		return "LangPortugese"
	case LangPunjabi:
		return "LangPunjabi"
	case LangQuechua:
		return "LangQuechua"
	case LangRomanian:
		return "LangRomanian"
	case LangRundi:
		return "LangRundi"
	case LangRussian:
		return "LangRussian"
	case LangSanskrit:
		return "LangSanskrit"
	case LangScottishGaelic:
		return "LangScottishGaelic"
	case LangSerbian:
		return "LangSerbian"
	case LangSimpChinese:
		return "LangSimpChinese"
	case LangSindhi:
		return "LangSindhi"
	case LangSinhalese:
		return "LangSinhalese"
	case LangSlovak:
		return "LangSlovak"
	case LangSlovenian:
		return "LangSlovenian"
	case LangSomali:
		return "LangSomali"
	case LangSpanish:
		return "LangSpanish"
	case LangSundaneseRom:
		return "LangSundaneseRom"
	case LangSwahili:
		return "LangSwahili"
	case LangSwedish:
		return "LangSwedish"
	case LangTagalog:
		return "LangTagalog"
	case LangTajiki:
		return "LangTajiki"
	case LangTamil:
		return "LangTamil"
	case LangTatar:
		return "LangTatar"
	case LangTelugu:
		return "LangTelugu"
	case LangThai:
		return "LangThai"
	case LangTibetan:
		return "LangTibetan"
	case LangTigrinya:
		return "LangTigrinya"
	case LangTongan:
		return "LangTongan"
	case LangTurkish:
		return "LangTurkish"
	case LangTurkmen:
		return "LangTurkmen"
	case LangUighur:
		return "LangUighur"
	case LangUkrainian:
		return "LangUkrainian"
	case LangUnspecified:
		return "LangUnspecified"
	case LangUrdu:
		return "LangUrdu"
	case LangUzbek:
		return "LangUzbek"
	case LangVietnamese:
		return "LangVietnamese"
	case LangWelsh:
		return "LangWelsh"
	case LangYiddish:
		return "LangYiddish"
	default:
		return fmt.Sprintf("Lang(%d)", e)
	}
}

type Large1BitMask uint

const (
	Large1BitMaskValue Large1BitMask = 'I'<<24 | 'C'<<16 | 'N'<<8 | '#' // 'ICN#'
	Large4BitData      Large1BitMask = 'i'<<24 | 'c'<<16 | 'l'<<8 | '4' // 'icl4'
	Large8BitData      Large1BitMask = 'i'<<24 | 'c'<<16 | 'l'<<8 | '8' // 'icl8'
	Mini1BitMask       Large1BitMask = 'i'<<24 | 'c'<<16 | 'm'<<8 | '#' // 'icm#'
	Mini4BitData       Large1BitMask = 'i'<<24 | 'c'<<16 | 'm'<<8 | '4' // 'icm4'
	Mini8BitData       Large1BitMask = 'i'<<24 | 'c'<<16 | 'm'<<8 | '8' // 'icm8'
	Small1BitMask      Large1BitMask = 'i'<<24 | 'c'<<16 | 's'<<8 | '#' // 'ics#'
	Small4BitData      Large1BitMask = 'i'<<24 | 'c'<<16 | 's'<<8 | '4' // 'ics4'
	Small8BitData      Large1BitMask = 'i'<<24 | 'c'<<16 | 's'<<8 | '8' // 'ics8'
)

func (e Large1BitMask) String() string {
	switch e {
	case Large1BitMaskValue:
		return "Large1BitMaskValue"
	case Large4BitData:
		return "Large4BitData"
	case Large8BitData:
		return "Large8BitData"
	case Mini1BitMask:
		return "Mini1BitMask"
	case Mini4BitData:
		return "Mini4BitData"
	case Mini8BitData:
		return "Mini8BitData"
	case Small1BitMask:
		return "Small1BitMask"
	case Small4BitData:
		return "Small4BitData"
	case Small8BitData:
		return "Small8BitData"
	default:
		return fmt.Sprintf("Large1BitMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/mdlabeldomain
type MDLabelDomain uint32

const (
	KMDLabelLocalDomain MDLabelDomain = 1
	KMDLabelUserDomain  MDLabelDomain = 0
)

func (e MDLabelDomain) String() string {
	switch e {
	case KMDLabelLocalDomain:
		return "KMDLabelLocalDomain"
	case KMDLabelUserDomain:
		return "KMDLabelUserDomain"
	default:
		return fmt.Sprintf("MDLabelDomain(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/mdqueryoptionflags
type MDQueryOptionFlags uint32

const (
	KMDQueryAllowFSTranslation MDQueryOptionFlags = 8
	// KMDQuerySynchronous: Specifies that a query should block during the initial gather phase.
	KMDQuerySynchronous MDQueryOptionFlags = 1
	// KMDQueryWantsUpdates: # Discussion
	KMDQueryWantsUpdates MDQueryOptionFlags = 4
)

func (e MDQueryOptionFlags) String() string {
	switch e {
	case KMDQueryAllowFSTranslation:
		return "KMDQueryAllowFSTranslation"
	case KMDQuerySynchronous:
		return "KMDQuerySynchronous"
	case KMDQueryWantsUpdates:
		return "KMDQueryWantsUpdates"
	default:
		return fmt.Sprintf("MDQueryOptionFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/coreservices/mdquerysortoptionflags
type MDQuerySortOptionFlags uint32

const (
	KMDQueryReverseSortOrderFlag MDQuerySortOptionFlags = 1
)

func (e MDQuerySortOptionFlags) String() string {
	switch e {
	case KMDQueryReverseSortOrderFlag:
		return "KMDQueryReverseSortOrderFlag"
	default:
		return fmt.Sprintf("MDQuerySortOptionFlags(%d)", e)
	}
}

type MPLibrary uint

const (
	// Deprecated.
	MPLibrary_DevelopmentRevision MPLibrary = 1
	// Deprecated.
	MPLibrary_MajorVersion MPLibrary = 2
	// Deprecated.
	MPLibrary_MinorVersion MPLibrary = 3
	// Deprecated.
	MPLibrary_Release MPLibrary = 1
)

func (e MPLibrary) String() string {
	switch e {
	case MPLibrary_DevelopmentRevision:
		return "MPLibrary_DevelopmentRevision"
	case MPLibrary_MajorVersion:
		return "MPLibrary_MajorVersion"
	case MPLibrary_MinorVersion:
		return "MPLibrary_MinorVersion"
	default:
		return fmt.Sprintf("MPLibrary(%d)", e)
	}
}

type Map uint

const (
	// Deprecated.
	MapChanged Map = 32
	// Deprecated.
	MapChangedBit Map = 5
	// Deprecated.
	MapCompact Map = 64
	// Deprecated.
	MapCompactBit Map = 6
	// Deprecated.
	MapReadOnly Map = 128
	// Deprecated.
	MapReadOnlyBit Map = 7
)

func (e Map) String() string {
	switch e {
	case MapChanged:
		return "MapChanged"
	case MapChangedBit:
		return "MapChangedBit"
	case MapCompact:
		return "MapCompact"
	case MapCompactBit:
		return "MapCompactBit"
	case MapReadOnly:
		return "MapReadOnly"
	case MapReadOnlyBit:
		return "MapReadOnlyBit"
	default:
		return fmt.Sprintf("Map(%d)", e)
	}
}

type Max uint

const (
	// Deprecated.
	MaxSize Max = 0
)

func (e Max) String() string {
	switch e {
	case MaxSize:
		return "MaxSize"
	default:
		return fmt.Sprintf("Max(%d)", e)
	}
}

type Mdy uint

const (
	Dmy      Mdy = 1
	Dym      Mdy = 4
	MdyValue Mdy = 0
	Myd      Mdy = 3
	Ydm      Mdy = 5
	Ymd      Mdy = 2
)

func (e Mdy) String() string {
	switch e {
	case Dmy:
		return "Dmy"
	case Dym:
		return "Dym"
	case MdyValue:
		return "MdyValue"
	case Myd:
		return "Myd"
	case Ydm:
		return "Ydm"
	case Ymd:
		return "Ymd"
	default:
		return fmt.Sprintf("Mdy(%d)", e)
	}
}

type MemROZWarn int

const (
	MemAZErr        MemROZWarn = -113
	MemAdrErr       MemROZWarn = -110
	MemBCErr        MemROZWarn = -115
	MemFullErr      MemROZWarn = -108
	MemLockedErr    MemROZWarn = -117
	MemPCErr        MemROZWarn = -114
	MemPurErr       MemROZWarn = -112
	MemROZErr       MemROZWarn = -99
	MemROZError     MemROZWarn = -99
	MemROZWarnValue MemROZWarn = -99
	MemSCErr        MemROZWarn = -116
	MemWZErr        MemROZWarn = -111
	NilHandleErr    MemROZWarn = -109
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
	MenuInvalidErr          Menu = -5623
	MenuItemNotFoundErr     Menu = -5622
	MenuNotFoundErr         Menu = -5620
	MenuPropertyInvalid     Menu = -5603
	MenuPropertyInvalidErr  Menu = -5603
	MenuPropertyNotFoundErr Menu = -5604
	MenuUsesSystemDefErr    Menu = -5621
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
	MidiDupIDErr        Midi = -260
	MidiInvalidCmdErr   Midi = -261
	MidiNameLenErr      Midi = -259
	MidiNoClientErr     Midi = -250
	MidiNoConErr        Midi = -257
	MidiNoPortErr       Midi = -251
	MidiTooManyConsErr  Midi = -253
	MidiTooManyPortsErr Midi = -252
	MidiVConnectErr     Midi = -254
	MidiVConnectMade    Midi = -255
	MidiVConnectRmvd    Midi = -256
	MidiWriteErr        Midi = -258
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

type MinCountry uint

const (
	// Deprecated.
	MaxCountry MinCountry = 108
	// Deprecated.
	MinCountryValue MinCountry = 0
)

func (e MinCountry) String() string {
	switch e {
	case MaxCountry:
		return "MaxCountry"
	case MinCountryValue:
		return "MinCountryValue"
	default:
		return fmt.Sprintf("MinCountry(%d)", e)
	}
}

type MmInternal int

const (
	MmInternalError MmInternal = -2526
)

func (e MmInternal) String() string {
	switch e {
	case MmInternalError:
		return "MmInternalError"
	default:
		return fmt.Sprintf("MmInternal(%d)", e)
	}
}

type MpWorkFlag uint

const (
	// Deprecated.
	MpWorkFlagCopyWorkBlock MpWorkFlag = 4
	// Deprecated.
	MpWorkFlagDoCompletion MpWorkFlag = 2
	// Deprecated.
	MpWorkFlagDoWork MpWorkFlag = 1
	// Deprecated.
	MpWorkFlagDontBlock MpWorkFlag = 8
	// Deprecated.
	MpWorkFlagGetIsRunning MpWorkFlag = 64
	// Deprecated.
	MpWorkFlagGetProcessorCount MpWorkFlag = 16
)

func (e MpWorkFlag) String() string {
	switch e {
	case MpWorkFlagCopyWorkBlock:
		return "MpWorkFlagCopyWorkBlock"
	case MpWorkFlagDoCompletion:
		return "MpWorkFlagDoCompletion"
	case MpWorkFlagDoWork:
		return "MpWorkFlagDoWork"
	case MpWorkFlagDontBlock:
		return "MpWorkFlagDontBlock"
	case MpWorkFlagGetIsRunning:
		return "MpWorkFlagGetIsRunning"
	case MpWorkFlagGetProcessorCount:
		return "MpWorkFlagGetProcessorCount"
	default:
		return fmt.Sprintf("MpWorkFlag(%d)", e)
	}
}

type Nbp int

const (
	NbpBuffOvr   Nbp = -1024
	NbpConfDiff  Nbp = -1026
	NbpDuplicate Nbp = -1027
	NbpNISErr    Nbp = -1029
	NbpNoConfirm Nbp = -1025
	NbpNotFound  Nbp = -1028
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

type NmTyp int

const (
	NmTypErr NmTyp = -299
)

func (e NmTyp) String() string {
	switch e {
	case NmTypErr:
		return "NmTypErr"
	default:
		return fmt.Sprintf("NmTyp(%d)", e)
	}
}

type No int

const (
	NoScrapErr No = -100
	NoTypeErr  No = -102
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
	BadBtSlpErr     NoAdrMkErr = -70
	BadCksmErr      NoAdrMkErr = -69
	BadDBtSlp       NoAdrMkErr = -73
	BadDCksum       NoAdrMkErr = -72
	BreakRecd       NoAdrMkErr = -90
	CantStepErr     NoAdrMkErr = -75
	ClkRdErr        NoAdrMkErr = -85
	ClkWrErr        NoAdrMkErr = -86
	DataVerErr      NoAdrMkErr = -68
	Fmt1Err         NoAdrMkErr = -82
	Fmt2Err         NoAdrMkErr = -83
	InitIWMErr      NoAdrMkErr = -77
	NoAdrMkErrValue NoAdrMkErr = -67
	NoDtaMkErr      NoAdrMkErr = -71
	PrInitErr       NoAdrMkErr = -88
	PrWrErr         NoAdrMkErr = -87
	RcvrErr         NoAdrMkErr = -89
	SectNFErr       NoAdrMkErr = -81
	SeekErr         NoAdrMkErr = -80
	SpdAdjErr       NoAdrMkErr = -79
	Tk0BadErr       NoAdrMkErr = -76
	TwoSideErr      NoAdrMkErr = -78
	VerErr          NoAdrMkErr = -84
	WrUnderrun      NoAdrMkErr = -74
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

type NoCalls uint

const (
	NoCallsValue NoCalls = 1
	NoRequest    NoCalls = 2
	SleepQType   NoCalls = 16
	SlpQType     NoCalls = 16
)

func (e NoCalls) String() string {
	switch e {
	case NoCallsValue:
		return "NoCallsValue"
	case NoRequest:
		return "NoRequest"
	case SleepQType:
		return "SleepQType"
	default:
		return fmt.Sprintf("NoCalls(%d)", e)
	}
}

type NoDeviceForChannel int

const (
	BadControllerHeight                NoDeviceForChannel = -9994
	BadSGChannel                       NoDeviceForChannel = -9406
	CannotMoveAttachedController       NoDeviceForChannel = -9999
	CannotSetWidthOfAttachedController NoDeviceForChannel = -9997
	CantDoThatInCurrentMode            NoDeviceForChannel = -9402
	ControllerBoundsNotExact           NoDeviceForChannel = -9996
	ControllerHasFixedHeight           NoDeviceForChannel = -9998
	CouldntGetRequiredComponent        NoDeviceForChannel = -9405
	DeviceCantMeetRequest              NoDeviceForChannel = -9408
	EditingNotAllowed                  NoDeviceForChannel = -9995
	GrabTimeComplete                   NoDeviceForChannel = -9401
	NoDeviceForChannelValue            NoDeviceForChannel = -9400
	NotEnoughDiskSpaceToGrab           NoDeviceForChannel = -9404
	NotEnoughMemoryToGrab              NoDeviceForChannel = -9403
	SeqGrabInfoNotAvailable            NoDeviceForChannel = -9407
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
	BadChannel                   NoHardware = -205
	BadFileFormat                NoHardware = -208
	BadFormat                    NoHardware = -206
	BuffersTooSmall              NoHardware = -210
	ChannelBusy                  NoHardware = -209
	ChannelNotBusy               NoHardware = -211
	NoHardwareValue              NoHardware = -200
	NoMoreRealTime               NoHardware = -212
	NotEnoughBufferSpace         NoHardware = -207
	NotEnoughHardware            NoHardware = -201
	QueueFull                    NoHardware = -203
	ResProblem                   NoHardware = -204
	SiBadDeviceName              NoHardware = -228
	SiBadRefNum                  NoHardware = -229
	SiBadSoundInDevice           NoHardware = -221
	SiDeviceBusyErr              NoHardware = -227
	SiHardDriveTooSlow           NoHardware = -224
	SiInputDeviceErr             NoHardware = -230
	SiInvalidCompression         NoHardware = -223
	SiInvalidSampleRate          NoHardware = -225
	SiInvalidSampleSize          NoHardware = -226
	SiNoBufferSpecified          NoHardware = -222
	SiNoSoundInHardware          NoHardware = -220
	SiUnknownInfoType            NoHardware = -231
	SiUnknownQuality             NoHardware = -232
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

type NoLoopingConstants uint

const (
	ForwardBackwardLooping NoLoopingConstants = 2
	ForwardLooping         NoLoopingConstants = 1
	NoLooping              NoLoopingConstants = 0
)

func (e NoLoopingConstants) String() string {
	switch e {
	case ForwardBackwardLooping:
		return "ForwardBackwardLooping"
	case ForwardLooping:
		return "ForwardLooping"
	case NoLooping:
		return "NoLooping"
	default:
		return fmt.Sprintf("NoLoopingConstants(%d)", e)
	}
}

type NoMaskFound int

const (
	NoMaskFoundErr NoMaskFound = -1000
)

func (e NoMaskFound) String() string {
	switch e {
	case NoMaskFoundErr:
		return "NoMaskFoundErr"
	default:
		return fmt.Sprintf("NoMaskFound(%d)", e)
	}
}

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

type NoneTypeConstants uint

const (
	ACE2Type  NoneTypeConstants = 'A'<<24 | 'C'<<16 | 'E'<<8 | '2' // 'ACE2'
	ACE8Type  NoneTypeConstants = 'A'<<24 | 'C'<<16 | 'E'<<8 | '8' // 'ACE8'
	MACE3Type NoneTypeConstants = 'M'<<24 | 'A'<<16 | 'C'<<8 | '3' // 'MAC3'
	MACE6Type NoneTypeConstants = 'M'<<24 | 'A'<<16 | 'C'<<8 | '6' // 'MAC6'
	NoneType  NoneTypeConstants = 'N'<<24 | 'O'<<16 | 'N'<<8 | 'E' // 'NONE'
)

func (e NoneTypeConstants) String() string {
	switch e {
	case ACE2Type:
		return "ACE2Type"
	case ACE8Type:
		return "ACE8Type"
	case MACE3Type:
		return "MACE3Type"
	case MACE6Type:
		return "MACE6Type"
	case NoneType:
		return "NoneType"
	default:
		return fmt.Sprintf("NoneTypeConstants(%d)", e)
	}
}

type NotAQTVRMovieErr int

const (
	CallNotSupportedByNodeErr     NotAQTVRMovieErr = -30542
	ConstraintReachedErr          NotAQTVRMovieErr = -30541
	InvalidHotSpotIDErr           NotAQTVRMovieErr = -30551
	InvalidNodeFormatErr          NotAQTVRMovieErr = -30550
	InvalidNodeIDErr              NotAQTVRMovieErr = -30544
	InvalidViewStateErr           NotAQTVRMovieErr = -30545
	LimitReachedErr               NotAQTVRMovieErr = -30549
	NoMemoryNodeFailedInitialize  NotAQTVRMovieErr = -30552
	NotAQTVRMovieErrValue         NotAQTVRMovieErr = -30540
	PropertyNotSupportedByNodeErr NotAQTVRMovieErr = -30547
	QtvrLibraryLoadErr            NotAQTVRMovieErr = -30554
	QtvrUninitialized             NotAQTVRMovieErr = -30555
	SelectorNotSupportedByNodeErr NotAQTVRMovieErr = -30543
	SettingNotSupportedByNodeErr  NotAQTVRMovieErr = -30548
	StreamingNodeNotReadyErr      NotAQTVRMovieErr = -30553
	TimeNotInViewErr              NotAQTVRMovieErr = -30546
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
	BtDupRecErr          NotBTree = -414
	BtKeyAttrErr         NotBTree = -417
	BtKeyLenErr          NotBTree = -416
	BtNoSpace            NotBTree = -413
	BtRecNotFnd          NotBTree = -415
	InvalidIndexErr      NotBTree = -20002
	NotBTreeValue        NotBTree = -410
	RecordDataTooBigErr  NotBTree = -20001
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
	CannotDeferErr          NotEnoughMemoryErr = -625
	CannotMakeContiguousErr NotEnoughMemoryErr = -622
	InterruptsMaskedErr     NotEnoughMemoryErr = -624
	NoMMUErr                NotEnoughMemoryErr = -626
	NotEnoughMemoryErrValue NotEnoughMemoryErr = -620
	NotHeldErr              NotEnoughMemoryErr = -621
	NotLockedErr            NotEnoughMemoryErr = -623
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
	AuthFailErr         NotInitErr = -927
	BadLocNameErr       NotInitErr = -931
	BadPortNameErr      NotInitErr = -919
	BadReqErr           NotInitErr = -909
	BadServiceMethodErr NotInitErr = -930
	// DestPortErr: Server hasn’t set `'SIZE'` resource toindicate awareness of high-level events, or else is not present
	DestPortErr        NotInitErr = -906
	GuestNotAllowedErr NotInitErr = -932
	LocalOnlyErr       NotInitErr = -905
	NameTypeErr        NotInitErr = -902
	NetworkErr         NotInitErr = -925
	NoDefaultUserErr   NotInitErr = -922
	NoGlobalsErr       NotInitErr = -904
	NoInformErr        NotInitErr = -926
	NoMachineNameErr   NotInitErr = -913
	// NoPortErr: Client hasn’t set `'SIZE'` resource toindicate awareness of high-level events
	NoPortErr         NotInitErr = -903
	NoResponseErr     NotInitErr = -915
	NoSessionErr      NotInitErr = -908
	NoToolboxNameErr  NotInitErr = -914
	NoUserNameErr     NotInitErr = -911
	NoUserRecErr      NotInitErr = -928
	NoUserRefErr      NotInitErr = -924
	NotInitErrValue   NotInitErr = -900
	NotLoggedInErr    NotInitErr = -923
	PortClosedErr     NotInitErr = -916
	PortNameExistsErr NotInitErr = -910
	// SessClosedErr: The `kAEDontReconnect` flagin the `sendMode` parameterwas set and the server quit, then restarted
	SessClosedErr NotInitErr = -917
	SessTableErr  NotInitErr = -907
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
	NrCallNotSupported        Nr = -2557
	NrDataTruncatedErr        Nr = -2543
	NrExitedIteratorScope     Nr = -2555
	NrInvalidEntryIterationOp Nr = -2552
	NrInvalidNodeErr          Nr = -2538
	NrIterationDone           Nr = -2554
	NrLockedErr               Nr = -2536
	NrNameErr                 Nr = -2541
	NrNotCreatedErr           Nr = -2540
	NrNotEnoughMemoryErr      Nr = -2537
	NrNotFoundErr             Nr = -2539
	NrNotModifiedErr          Nr = -2547
	NrNotSlotDeviceErr        Nr = -2542
	NrOverrunErr              Nr = -2548
	NrPathBufferTooSmall      Nr = -2551
	NrPathNotFound            Nr = -2550
	NrPowerErr                Nr = -2544
	NrPowerSwitchAbortErr     Nr = -2545
	NrPropertyAlreadyExists   Nr = -2553
	NrResultCodeBase          Nr = -2549
	NrTransactionAborted      Nr = -2556
	NrTypeMismatchErr         Nr = -2546
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
	NumberFormattingBadCurrencyPositionErr       Number = -5211
	NumberFormattingBadFormatErr                 Number = -5207
	NumberFormattingBadNumberFormattingObjectErr Number = -5202
	NumberFormattingBadOptionsErr                Number = -5208
	NumberFormattingBadTokenErr                  Number = -5209
	NumberFormattingDelimiterMissingErr          Number = -5205
	NumberFormattingEmptyFormatErr               Number = -5206
	NumberFormattingLiteralMissingErr            Number = -5204
	NumberFormattingNotADigitErr                 Number = -5212
	NumberFormattingNotANumberErr                Number = -5200
	NumberFormattingOverflowInDestinationErr     Number = -5201
	NumberFormattingSpuriousCharErr              Number = -5203
	NumberFormattingUnOrderedCurrencyRangeErr    Number = -5210
	NumberFormattingUnOrdredCurrencyRangeErr     Number = -5210
	NumberFortmattingNotADigitErr                Number = -5212
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

type OverallActConstants uint

const (
	HDActivity   OverallActConstants = 3
	IdleActivity OverallActConstants = 4
	NetActivity  OverallActConstants = 2
	OverallAct   OverallActConstants = 0
	UsrActivity  OverallActConstants = 1
)

func (e OverallActConstants) String() string {
	switch e {
	case HDActivity:
		return "HDActivity"
	case IdleActivity:
		return "IdleActivity"
	case NetActivity:
		return "NetActivity"
	case OverallAct:
		return "OverallAct"
	case UsrActivity:
		return "UsrActivity"
	default:
		return fmt.Sprintf("OverallActConstants(%d)", e)
	}
}

type PArcAngle uint

const (
	PArcAngleValue     PArcAngle = 'p'<<24 | 'a'<<16 | 'r'<<8 | 'c' // 'parc'
	PBackgroundColor   PArcAngle = 'p'<<24 | 'b'<<16 | 'c'<<8 | 'l' // 'pbcl'
	PBackgroundPattern PArcAngle = 'p'<<24 | 'b'<<16 | 'p'<<8 | 't' // 'pbpt'
	PBestType          PArcAngle = 'p'<<24 | 'b'<<16 | 's'<<8 | 't' // 'pbst'
	PBounds            PArcAngle = 'p'<<24 | 'b'<<16 | 'n'<<8 | 'd' // 'pbnd'
	PClass             PArcAngle = 'p'<<24 | 'c'<<16 | 'l'<<8 | 's' // 'pcls'
	PClipboard         PArcAngle = 'p'<<24 | 'c'<<16 | 'l'<<8 | 'i' // 'pcli'
	PColor             PArcAngle = 'c'<<24 | 'o'<<16 | 'l'<<8 | 'r' // 'colr'
	PColorTable        PArcAngle = 'c'<<24 | 'l'<<16 | 't'<<8 | 'b' // 'cltb'
	PContents          PArcAngle = 'p'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'pcnt'
	PCornerCurveHeight PArcAngle = 'p'<<24 | 'c'<<16 | 'h'<<8 | 'd' // 'pchd'
	PCornerCurveWidth  PArcAngle = 'p'<<24 | 'c'<<16 | 'w'<<8 | 'd' // 'pcwd'
	PDashStyle         PArcAngle = 'p'<<24 | 'd'<<16 | 's'<<8 | 't' // 'pdst'
	PDefaultType       PArcAngle = 'd'<<24 | 'e'<<16 | 'f'<<8 | 't' // 'deft'
	PDefinitionRect    PArcAngle = 'p'<<24 | 'd'<<16 | 'r'<<8 | 't' // 'pdrt'
	PEnabled           PArcAngle = 'e'<<24 | 'n'<<16 | 'b'<<8 | 'l' // 'enbl'
	PEndPoint          PArcAngle = 'p'<<24 | 'e'<<16 | 'n'<<8 | 'd' // 'pend'
	PFillColor         PArcAngle = 'f'<<24 | 'l'<<16 | 'c'<<8 | 'l' // 'flcl'
	PFillPattern       PArcAngle = 'f'<<24 | 'l'<<16 | 'p'<<8 | 't' // 'flpt'
	PFont              PArcAngle = 'f'<<24 | 'o'<<16 | 'n'<<8 | 't' // 'font'
)

func (e PArcAngle) String() string {
	switch e {
	case PArcAngleValue:
		return "PArcAngleValue"
	case PBackgroundColor:
		return "PBackgroundColor"
	case PBackgroundPattern:
		return "PBackgroundPattern"
	case PBestType:
		return "PBestType"
	case PBounds:
		return "PBounds"
	case PClass:
		return "PClass"
	case PClipboard:
		return "PClipboard"
	case PColor:
		return "PColor"
	case PColorTable:
		return "PColorTable"
	case PContents:
		return "PContents"
	case PCornerCurveHeight:
		return "PCornerCurveHeight"
	case PCornerCurveWidth:
		return "PCornerCurveWidth"
	case PDashStyle:
		return "PDashStyle"
	case PDefaultType:
		return "PDefaultType"
	case PDefinitionRect:
		return "PDefinitionRect"
	case PEnabled:
		return "PEnabled"
	case PEndPoint:
		return "PEndPoint"
	case PFillColor:
		return "PFillColor"
	case PFillPattern:
		return "PFillPattern"
	case PFont:
		return "PFont"
	default:
		return fmt.Sprintf("PArcAngle(%d)", e)
	}
}

type PFormula uint

const (
	PFormulaValue    PFormula = 'p'<<24 | 'f'<<16 | 'o'<<8 | 'r' // 'pfor'
	PGraphicObjects  PFormula = 'g'<<24 | 'o'<<16 | 'b'<<8 | 's' // 'gobs'
	PHasCloseBox     PFormula = 'h'<<24 | 'c'<<16 | 'l'<<8 | 'b' // 'hclb'
	PHasTitleBar     PFormula = 'p'<<24 | 't'<<16 | 'i'<<8 | 't' // 'ptit'
	PID              PFormula = 'I'<<24 | 'D'<<16 | ' '<<8 | ' ' // 'ID  '
	PIndex           PFormula = 'p'<<24 | 'i'<<16 | 'd'<<8 | 'x' // 'pidx'
	PInsertionLoc    PFormula = 'p'<<24 | 'i'<<16 | 'n'<<8 | 's' // 'pins'
	PIsFloating      PFormula = 'i'<<24 | 's'<<16 | 'f'<<8 | 'l' // 'isfl'
	PIsFrontProcess  PFormula = 'p'<<24 | 'i'<<16 | 's'<<8 | 'f' // 'pisf'
	PIsModal         PFormula = 'p'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'pmod'
	PIsModified      PFormula = 'i'<<24 | 'm'<<16 | 'o'<<8 | 'd' // 'imod'
	PIsResizable     PFormula = 'p'<<24 | 'r'<<16 | 's'<<8 | 'z' // 'prsz'
	PIsStationeryPad PFormula = 'p'<<24 | 's'<<16 | 'p'<<8 | 'd' // 'pspd'
	PIsZoomable      PFormula = 'i'<<24 | 's'<<16 | 'z'<<8 | 'm' // 'iszm'
	PIsZoomed        PFormula = 'p'<<24 | 'z'<<16 | 'u'<<8 | 'm' // 'pzum'
	PItemNumber      PFormula = 'i'<<24 | 't'<<16 | 'm'<<8 | 'n' // 'itmn'
	PJustification   PFormula = 'p'<<24 | 'j'<<16 | 's'<<8 | 't' // 'pjst'
	PLineArrow       PFormula = 'a'<<24 | 'r'<<16 | 'r'<<8 | 'o' // 'arro'
	PMenuID          PFormula = 'm'<<24 | 'n'<<16 | 'i'<<8 | 'd' // 'mnid'
	PName            PFormula = 'p'<<24 | 'n'<<16 | 'a'<<8 | 'm' // 'pnam'
)

func (e PFormula) String() string {
	switch e {
	case PFormulaValue:
		return "PFormulaValue"
	case PGraphicObjects:
		return "PGraphicObjects"
	case PHasCloseBox:
		return "PHasCloseBox"
	case PHasTitleBar:
		return "PHasTitleBar"
	case PID:
		return "PID"
	case PIndex:
		return "PIndex"
	case PInsertionLoc:
		return "PInsertionLoc"
	case PIsFloating:
		return "PIsFloating"
	case PIsFrontProcess:
		return "PIsFrontProcess"
	case PIsModal:
		return "PIsModal"
	case PIsModified:
		return "PIsModified"
	case PIsResizable:
		return "PIsResizable"
	case PIsStationeryPad:
		return "PIsStationeryPad"
	case PIsZoomable:
		return "PIsZoomable"
	case PIsZoomed:
		return "PIsZoomed"
	case PItemNumber:
		return "PItemNumber"
	case PJustification:
		return "PJustification"
	case PLineArrow:
		return "PLineArrow"
	case PMenuID:
		return "PMenuID"
	case PName:
		return "PName"
	default:
		return fmt.Sprintf("PFormula(%d)", e)
	}
}

type PNewElementLoc uint

const (
	PNewElementLocValue PNewElementLoc = 'p'<<24 | 'n'<<16 | 'e'<<8 | 'l' // 'pnel'
	PPenColor           PNewElementLoc = 'p'<<24 | 'p'<<16 | 'c'<<8 | 'l' // 'ppcl'
	PPenPattern         PNewElementLoc = 'p'<<24 | 'p'<<16 | 'p'<<8 | 'a' // 'pppa'
	PPenWidth           PNewElementLoc = 'p'<<24 | 'p'<<16 | 'w'<<8 | 'd' // 'ppwd'
	PPixelDepth         PNewElementLoc = 'p'<<24 | 'd'<<16 | 'p'<<8 | 't' // 'pdpt'
	PPointList          PNewElementLoc = 'p'<<24 | 't'<<16 | 'l'<<8 | 't' // 'ptlt'
	PPointSize          PNewElementLoc = 'p'<<24 | 't'<<16 | 's'<<8 | 'z' // 'ptsz'
	PProtection         PNewElementLoc = 'p'<<24 | 'p'<<16 | 'r'<<8 | 'o' // 'ppro'
	PRotation           PNewElementLoc = 'p'<<24 | 'r'<<16 | 'o'<<8 | 't' // 'prot'
	PScale              PNewElementLoc = 'p'<<24 | 's'<<16 | 'c'<<8 | 'l' // 'pscl'
	PScript             PNewElementLoc = 's'<<24 | 'c'<<16 | 'p'<<8 | 't' // 'scpt'
	PScriptTag          PNewElementLoc = 'p'<<24 | 's'<<16 | 'c'<<8 | 't' // 'psct'
	PSelected           PNewElementLoc = 's'<<24 | 'e'<<16 | 'l'<<8 | 'c' // 'selc'
	PSelection          PNewElementLoc = 's'<<24 | 'e'<<16 | 'l'<<8 | 'e' // 'sele'
	PStartAngle         PNewElementLoc = 'p'<<24 | 'a'<<16 | 'n'<<8 | 'g' // 'pang'
	PStartPoint         PNewElementLoc = 'p'<<24 | 's'<<16 | 't'<<8 | 'p' // 'pstp'
	PTextColor          PNewElementLoc = 'p'<<24 | 't'<<16 | 'x'<<8 | 'c' // 'ptxc'
	PTextFont           PNewElementLoc = 'p'<<24 | 't'<<16 | 'x'<<8 | 'f' // 'ptxf'
	PTextItemDelimiters PNewElementLoc = 't'<<24 | 'x'<<16 | 'd'<<8 | 'l' // 'txdl'
	PTextPointSize      PNewElementLoc = 'p'<<24 | 't'<<16 | 'p'<<8 | 's' // 'ptps'
)

func (e PNewElementLoc) String() string {
	switch e {
	case PNewElementLocValue:
		return "PNewElementLocValue"
	case PPenColor:
		return "PPenColor"
	case PPenPattern:
		return "PPenPattern"
	case PPenWidth:
		return "PPenWidth"
	case PPixelDepth:
		return "PPixelDepth"
	case PPointList:
		return "PPointList"
	case PPointSize:
		return "PPointSize"
	case PProtection:
		return "PProtection"
	case PRotation:
		return "PRotation"
	case PScale:
		return "PScale"
	case PScript:
		return "PScript"
	case PScriptTag:
		return "PScriptTag"
	case PSelected:
		return "PSelected"
	case PSelection:
		return "PSelection"
	case PStartAngle:
		return "PStartAngle"
	case PStartPoint:
		return "PStartPoint"
	case PTextColor:
		return "PTextColor"
	case PTextFont:
		return "PTextFont"
	case PTextItemDelimiters:
		return "PTextItemDelimiters"
	case PTextPointSize:
		return "PTextPointSize"
	default:
		return fmt.Sprintf("PNewElementLoc(%d)", e)
	}
}

type PScheme uint

const (
	PDNSForm      PScheme = 'p'<<24 | 'D'<<16 | 'N'<<8 | 'S' // 'pDNS'
	PFTPKind      PScheme = 'k'<<24 | 'i'<<16 | 'n'<<8 | 'd' // 'kind'
	PHost         PScheme = 'H'<<24 | 'O'<<16 | 'S'<<8 | 'T' // 'HOST'
	PPath         PScheme = 'F'<<24 | 'T'<<16 | 'P'<<8 | 'c' // 'FTPc'
	PSchemeValue  PScheme = 'p'<<24 | 'u'<<16 | 's'<<8 | 'c' // 'pusc'
	PTextEncoding PScheme = 'p'<<24 | 't'<<16 | 'x'<<8 | 'e' // 'ptxe'
	PURL          PScheme = 'p'<<24 | 'U'<<16 | 'R'<<8 | 'L' // 'pURL'
	PUserName     PScheme = 'R'<<24 | 'A'<<16 | 'u'<<8 | 'n' // 'RAun'
	PUserPassword PScheme = 'R'<<24 | 'A'<<16 | 'p'<<8 | 'w' // 'RApw'
)

func (e PScheme) String() string {
	switch e {
	case PDNSForm:
		return "PDNSForm"
	case PFTPKind:
		return "PFTPKind"
	case PHost:
		return "PHost"
	case PPath:
		return "PPath"
	case PSchemeValue:
		return "PSchemeValue"
	case PTextEncoding:
		return "PTextEncoding"
	case PURL:
		return "PURL"
	case PUserName:
		return "PUserName"
	case PUserPassword:
		return "PUserPassword"
	default:
		return fmt.Sprintf("PScheme(%d)", e)
	}
}

type PTextStyles uint

const (
	PTextStylesValue PTextStyles = 't'<<24 | 'x'<<16 | 's'<<8 | 't' // 'txst'
	PTransferMode    PTextStyles = 'p'<<24 | 'p'<<16 | 't'<<8 | 'm' // 'pptm'
	PTranslation     PTextStyles = 'p'<<24 | 't'<<16 | 'r'<<8 | 's' // 'ptrs'
	PUniformStyles   PTextStyles = 'u'<<24 | 's'<<16 | 't'<<8 | 'l' // 'ustl'
	PUpdateOn        PTextStyles = 'p'<<24 | 'u'<<16 | 'p'<<8 | 'd' // 'pupd'
	PUserSelection   PTextStyles = 'p'<<24 | 'u'<<16 | 's'<<8 | 'l' // 'pusl'
	PVersion         PTextStyles = 'v'<<24 | 'e'<<16 | 'r'<<8 | 's' // 'vers'
	PVisible         PTextStyles = 'p'<<24 | 'v'<<16 | 'i'<<8 | 's' // 'pvis'
)

func (e PTextStyles) String() string {
	switch e {
	case PTextStylesValue:
		return "PTextStylesValue"
	case PTransferMode:
		return "PTransferMode"
	case PTranslation:
		return "PTranslation"
	case PUniformStyles:
		return "PUniformStyles"
	case PUpdateOn:
		return "PUpdateOn"
	case PUserSelection:
		return "PUserSelection"
	case PVersion:
		return "PVersion"
	case PVisible:
		return "PVisible"
	default:
		return fmt.Sprintf("PTextStyles(%d)", e)
	}
}

type ParamErr int

const (
	BadUnitErr           ParamErr = -21
	ClosErr              ParamErr = -24
	ControlErr           ParamErr = -17
	CorErr               ParamErr = -3
	DInstErr             ParamErr = -26
	DRemovErr            ParamErr = -25
	NoHardwareErr        ParamErr = -200
	NotEnoughHardwareErr ParamErr = -201
	OpenErr              ParamErr = -23
	ParamErrValue        ParamErr = -50
	QErr                 ParamErr = -1
	ReadErr              ParamErr = -19
	SeNoDB               ParamErr = -8
	SlpTypeErr           ParamErr = -5
	StatusErr            ParamErr = -18
	UnimpErr             ParamErr = -4
	UnitEmptyErr         ParamErr = -22
	UserCanceledErr      ParamErr = -128
	VTypErr              ParamErr = -2
	WritErr              ParamErr = -20
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

type Platform uint

const (
	// Deprecated.
	Platform68k Platform = 1
	// Deprecated.
	PlatformAIXppc Platform = 1300
	// Deprecated.
	PlatformArm64NativeEntryPoint Platform = 9
	// Deprecated.
	PlatformIA32NativeEntryPoint Platform = 6
	// Deprecated.
	PlatformIRIXmips Platform = 1000
	// Deprecated.
	PlatformInterpreted Platform = 3
	// Deprecated.
	PlatformLinuxintel Platform = 1201
	// Deprecated.
	PlatformLinuxppc Platform = 1200
	// Deprecated.
	PlatformMacOSx86 Platform = 1500
	// Deprecated.
	PlatformNeXT68k Platform = 1403
	// Deprecated.
	PlatformNeXTIntel Platform = 1400
	// Deprecated.
	PlatformNeXTppc Platform = 1401
	// Deprecated.
	PlatformNeXTsparc Platform = 1402
	// Deprecated.
	PlatformPowerPC Platform = 2
	// Deprecated.
	PlatformPowerPC64NativeEntryPoint Platform = 7
	// Deprecated.
	PlatformPowerPCNativeEntryPoint Platform = 5
	// Deprecated.
	PlatformSunOSintel Platform = 1101
	// Deprecated.
	PlatformSunOSsparc Platform = 1100
	// Deprecated.
	PlatformWin32 Platform = 4
	// Deprecated.
	PlatformX86_64NativeEntryPoint Platform = 8
)

func (e Platform) String() string {
	switch e {
	case Platform68k:
		return "Platform68k"
	case PlatformAIXppc:
		return "PlatformAIXppc"
	case PlatformArm64NativeEntryPoint:
		return "PlatformArm64NativeEntryPoint"
	case PlatformIA32NativeEntryPoint:
		return "PlatformIA32NativeEntryPoint"
	case PlatformIRIXmips:
		return "PlatformIRIXmips"
	case PlatformInterpreted:
		return "PlatformInterpreted"
	case PlatformLinuxintel:
		return "PlatformLinuxintel"
	case PlatformLinuxppc:
		return "PlatformLinuxppc"
	case PlatformMacOSx86:
		return "PlatformMacOSx86"
	case PlatformNeXT68k:
		return "PlatformNeXT68k"
	case PlatformNeXTIntel:
		return "PlatformNeXTIntel"
	case PlatformNeXTppc:
		return "PlatformNeXTppc"
	case PlatformNeXTsparc:
		return "PlatformNeXTsparc"
	case PlatformPowerPC:
		return "PlatformPowerPC"
	case PlatformPowerPC64NativeEntryPoint:
		return "PlatformPowerPC64NativeEntryPoint"
	case PlatformPowerPCNativeEntryPoint:
		return "PlatformPowerPCNativeEntryPoint"
	case PlatformSunOSintel:
		return "PlatformSunOSintel"
	case PlatformSunOSsparc:
		return "PlatformSunOSsparc"
	case PlatformWin32:
		return "PlatformWin32"
	case PlatformX86_64NativeEntryPoint:
		return "PlatformX86_64NativeEntryPoint"
	default:
		return fmt.Sprintf("Platform(%d)", e)
	}
}

type PleaseCacheBit uint

const (
	ForceReadBit        PleaseCacheBit = 6
	ForceReadMask       PleaseCacheBit = 64
	NewLineBit          PleaseCacheBit = 7
	NewLineCharMask     PleaseCacheBit = 65280
	NewLineMask         PleaseCacheBit = 128
	NoCacheBit          PleaseCacheBit = 5
	NoCacheMask         PleaseCacheBit = 32
	PleaseCacheBitValue PleaseCacheBit = 4
	PleaseCacheMask     PleaseCacheBit = 16
	RdVerify            PleaseCacheBit = 64
	RdVerifyBit         PleaseCacheBit = 6
	RdVerifyMask        PleaseCacheBit = 64
)

func (e PleaseCacheBit) String() string {
	switch e {
	case ForceReadBit:
		return "ForceReadBit"
	case ForceReadMask:
		return "ForceReadMask"
	case NewLineBit:
		return "NewLineBit"
	case NewLineCharMask:
		return "NewLineCharMask"
	case NewLineMask:
		return "NewLineMask"
	case NoCacheBit:
		return "NoCacheBit"
	case NoCacheMask:
		return "NoCacheMask"
	case PleaseCacheBitValue:
		return "PleaseCacheBitValue"
	case PleaseCacheMask:
		return "PleaseCacheMask"
	default:
		return fmt.Sprintf("PleaseCacheBit(%d)", e)
	}
}

type Pm int

const (
	PmBusyErr      Pm = -13000
	PmRecvEndErr   Pm = -13005
	PmRecvStartErr Pm = -13004
	PmReplyTOErr   Pm = -13001
	PmSendEndErr   Pm = -13003
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

type PrinterStatusOpCodeNotSupported int

const (
	PrinterStatusOpCodeNotSupportedErr PrinterStatusOpCodeNotSupported = -25280
)

func (e PrinterStatusOpCodeNotSupported) String() string {
	switch e {
	case PrinterStatusOpCodeNotSupportedErr:
		return "PrinterStatusOpCodeNotSupportedErr"
	default:
		return fmt.Sprintf("PrinterStatusOpCodeNotSupported(%d)", e)
	}
}

type ProcNotFound int

const (
	AppIsDaemon              ProcNotFound = -606
	AppMemFullErr            ProcNotFound = -605
	AppModeErr               ProcNotFound = -602
	BufferIsSmall            ProcNotFound = -607
	ConnectionInvalid        ProcNotFound = -609
	HardwareConfigErr        ProcNotFound = -604
	MemFragErr               ProcNotFound = -601
	NoOutstandingHLE         ProcNotFound = -608
	NoUserInteractionAllowed ProcNotFound = -610
	ProcNotFoundValue        ProcNotFound = -600
	ProtocolErr              ProcNotFound = -603
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
	QtsAddressBusyErr         Qts = -5421
	QtsBadDataErr             Qts = -5402
	QtsBadSelectorErr         Qts = -5400
	QtsBadStateErr            Qts = -5401
	QtsConnectionFailedErr    Qts = -5420
	QtsTimeoutErr             Qts = -5408
	QtsTooMuchDataErr         Qts = -5406
	QtsUnknownValueErr        Qts = -5407
	QtsUnsupportedDataTypeErr Qts = -5403
	QtsUnsupportedFeatureErr  Qts = -5405
	QtsUnsupportedRateErr     Qts = -5404
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

type RAlias uint

const (
	// Deprecated.
	RAliasType RAlias = 'a'<<24 | 'l'<<16 | 'i'<<8 | 's' // 'alis'
)

func (e RAlias) String() string {
	switch e {
	case RAliasType:
		return "RAliasType"
	default:
		return fmt.Sprintf("RAlias(%d)", e)
	}
}

type RcDB int

const (
	RcDBAsyncNotSupp  RcDB = -809
	RcDBBadAsyncPB    RcDB = -810
	RcDBBadDDEV       RcDB = -808
	RcDBBadSessID     RcDB = -806
	RcDBBadSessNum    RcDB = -807
	RcDBBadType       RcDB = -803
	RcDBBreak         RcDB = -804
	RcDBError         RcDB = -802
	RcDBExec          RcDB = -805
	RcDBNoHandler     RcDB = -811
	RcDBNull          RcDB = -800
	RcDBPackNotInited RcDB = -813
	RcDBValue         RcDB = -801
	RcDBWrongVersion  RcDB = -812
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

const (
	// Deprecated.
	RegisterComponentAfterExisting uint = 4
	// Deprecated.
	RegisterComponentAliasesOnly uint = 8
	// Deprecated.
	RegisterComponentGlobal uint = 1
	// Deprecated.
	RegisterComponentNoDuplicates uint = 2
)

type ReqFailed int

const (
	BadATPSkt      ReqFailed = -1099
	BadBuffNum     ReqFailed = -1100
	CbNotFound     ReqFailed = -1102
	NoDataArea     ReqFailed = -1104
	NoRelErr       ReqFailed = -1101
	NoSendResp     ReqFailed = -1103
	ReqAborted     ReqFailed = -1105
	ReqFailedValue ReqFailed = -1096
	TooManyReqs    ReqFailed = -1097
	TooManySkts    ReqFailed = -1098
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

type Res uint

const (
	// Deprecated.
	ResChanged Res = 2
	// Deprecated.
	ResChangedBit Res = 1
	// Deprecated.
	ResLocked Res = 16
	// Deprecated.
	ResLockedBit Res = 4
	// Deprecated.
	ResPreload Res = 4
	// Deprecated.
	ResPreloadBit Res = 2
	// Deprecated.
	ResProtected Res = 8
	// Deprecated.
	ResProtectedBit Res = 3
	// Deprecated.
	ResPurgeable Res = 32
	// Deprecated.
	ResPurgeableBit Res = 5
	// Deprecated.
	ResSysHeap Res = 64
	// Deprecated.
	ResSysHeapBit Res = 6
	// Deprecated.
	ResSysRefBit Res = 7
)

func (e Res) String() string {
	switch e {
	case ResChanged:
		return "ResChanged"
	case ResChangedBit:
		return "ResChangedBit"
	case ResLocked:
		return "ResLocked"
	case ResLockedBit:
		return "ResLockedBit"
	case ResProtected:
		return "ResProtected"
	case ResProtectedBit:
		return "ResProtectedBit"
	case ResPurgeable:
		return "ResPurgeable"
	case ResPurgeableBit:
		return "ResPurgeableBit"
	case ResSysHeap:
		return "ResSysHeap"
	case ResSysHeapBit:
		return "ResSysHeapBit"
	case ResSysRefBit:
		return "ResSysRefBit"
	default:
		return fmt.Sprintf("Res(%d)", e)
	}
}

type ResourceInMemory int

const (
	AddRefFailed            ResourceInMemory = -195
	AddResFailed            ResourceInMemory = -194
	BadExtResource          ResourceInMemory = -185
	CantDecompress          ResourceInMemory = -186
	InputOutOfBounds        ResourceInMemory = -190
	InsufficientStackErr    ResourceInMemory = -149
	MapReadErr              ResourceInMemory = -199
	NoMemForPictPlaybackErr ResourceInMemory = -145
	NsStackErr              ResourceInMemory = -149
	PixMapTooDeepErr        ResourceInMemory = -148
	ResAttrErr              ResourceInMemory = -198
	ResFNotFound            ResourceInMemory = -193
	ResNotFound             ResourceInMemory = -192
	ResourceInMemoryValue   ResourceInMemory = -188
	RgnOverflowErr          ResourceInMemory = -147
	RgnTooBigError          ResourceInMemory = -147
	RmvRefFailed            ResourceInMemory = -197
	RmvResFailed            ResourceInMemory = -196
	WritingPastEnd          ResourceInMemory = -189
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

type RomanSysFond uint

const (
	// Deprecated.
	RomanAppFond RomanSysFond = 3
	// Deprecated.
	RomanFlags RomanSysFond = 7
	// Deprecated.
	RomanSysFondValue RomanSysFond = 16383
	// Deprecated.
	SmFondEnd RomanSysFond = 49152
	// Deprecated.
	SmFondStart RomanSysFond = 16384
	// Deprecated.
	SmUprHalfCharSet RomanSysFond = 128
)

func (e RomanSysFond) String() string {
	switch e {
	case RomanAppFond:
		return "RomanAppFond"
	case RomanFlags:
		return "RomanFlags"
	case RomanSysFondValue:
		return "RomanSysFondValue"
	case SmFondEnd:
		return "SmFondEnd"
	case SmFondStart:
		return "SmFondStart"
	case SmUprHalfCharSet:
		return "SmUprHalfCharSet"
	default:
		return fmt.Sprintf("RomanSysFond(%d)", e)
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

type ShortDate uint

const (
	AbbrevDate     ShortDate = 2
	LongDate       ShortDate = 1
	ShortDateValue ShortDate = 0
)

func (e ShortDate) String() string {
	switch e {
	case AbbrevDate:
		return "AbbrevDate"
	case LongDate:
		return "LongDate"
	case ShortDateValue:
		return "ShortDateValue"
	default:
		return fmt.Sprintf("ShortDate(%d)", e)
	}
}

type SiInitSDTblErr int

const (
	SdmInitErr          SiInitSDTblErr = 11
	SdmJTInitErr        SiInitSDTblErr = 10
	SdmPRAMInitErr      SiInitSDTblErr = 13
	SdmPriInitErr       SiInitSDTblErr = 14
	SdmSRTInitErr       SiInitSDTblErr = 12
	SiInitSDTblErrValue SiInitSDTblErr = 1
	SiInitSPTblErr      SiInitSDTblErr = 3
	SiInitVBLQsErr      SiInitSDTblErr = 2
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
	SmBadsPtrErr    Sm = -346
	SmBlkMoveErr    Sm = -340
	SmByteLanesErr  Sm = -347
	SmCPUErr        Sm = -334
	SmCRCFail       Sm = -301
	SmCkStatusErr   Sm = -341
	SmDisDrvrNamErr Sm = -343
	SmDisabledSlot  Sm = -305
	SmEmptySlot     Sm = -300
	SmFormatErr     Sm = -302
	SmGetDrvrNamErr Sm = -342
	SmNewPErr       Sm = -339
	SmNilsBlockErr  Sm = -336
	SmNoDir         Sm = -304
	SmNoGoodOpens   Sm = -349
	SmNoMoresRsrcs  Sm = -344
	SmNosInfoArray  Sm = -306
	SmOffsetErr     Sm = -348
	SmPRAMInitErr   Sm = -292
	SmPriInitErr    Sm = -293
	SmRecNotFnd     Sm = -351
	SmRevisionErr   Sm = -303
	SmSDMInitErr    Sm = -290
	SmSRTInitErr    Sm = -291
	SmSRTOvrFlErr   Sm = -350
	SmSelOOBErr     Sm = -338
	SmSlotOOBErr    Sm = -337
	SmsGetDrvrErr   Sm = -345
	SmsPointerNil   Sm = -335
	// Deprecated.
	SmAllScripts Sm = -3
	// Deprecated.
	SmAmharic Sm = 28
	// Deprecated.
	SmArabic Sm = 4
	// Deprecated.
	SmArmenian Sm = 24
	// Deprecated.
	SmBadScript Sm = -2
	// Deprecated.
	SmBadVerb Sm = -1
	// Deprecated.
	SmBengali Sm = 13
	// Deprecated.
	SmBurmese Sm = 19
	// Deprecated.
	SmCentralEuroRoman Sm = 29
	// Deprecated.
	SmChar1byte Sm = 0
	// Deprecated.
	SmChar2byte Sm = 32768
	// Deprecated.
	SmCharFISGreek Sm = 5
	// Deprecated.
	SmCharFISRussian Sm = 6
	// Deprecated.
	SmCharHorizontal Sm = 0
	// Deprecated.
	SmCharLeft Sm = 0
	// Deprecated.
	SmCharLower Sm = 0
	// Deprecated.
	SmCharRight Sm = 8192
	// Deprecated.
	SmCharUpper Sm = 16384
	// Deprecated.
	SmCharVertical Sm = 4096
	// Deprecated.
	SmChinese Sm = 2
	// Deprecated.
	SmCurrentScript Sm = -2
	// Deprecated.
	SmCyrillic Sm = 7
	// Deprecated.
	SmDevanagari Sm = 9
	// Deprecated.
	SmEastEurRoman Sm = 29
	// Deprecated.
	SmEthiopic Sm = 28
	// Deprecated.
	SmExtArabic Sm = 31
	// Deprecated.
	SmFISClassLvl1 Sm = 0
	// Deprecated.
	SmFISClassLvl2 Sm = 256
	// Deprecated.
	SmFISClassUser Sm = 512
	// Deprecated.
	SmFirstByte Sm = -1
	// Deprecated.
	SmGeez Sm = 28
	// Deprecated.
	SmGeorgian Sm = 23
	// Deprecated.
	SmGreek Sm = 6
	// Deprecated.
	SmGujarati Sm = 11
	// Deprecated.
	SmGurmukhi Sm = 10
	// Deprecated.
	SmHebrew Sm = 5
	// Deprecated.
	SmIdeographicLevel1 Sm = 0
	// Deprecated.
	SmIdeographicLevel2 Sm = 256
	// Deprecated.
	SmIdeographicUser Sm = 512
	// Deprecated.
	SmJamoBogJaeum Sm = 256
	// Deprecated.
	SmJamoBogMoeum Sm = 768
	// Deprecated.
	SmJamoJaeum Sm = 0
	// Deprecated.
	SmJamoMoeum Sm = 512
	// Deprecated.
	SmJapanese Sm = 1
	// Deprecated.
	SmKCHRCache Sm = 38
	// Deprecated.
	SmKanaHardOK Sm = 512
	// Deprecated.
	SmKanaSmall Sm = 256
	// Deprecated.
	SmKanaSoftOK Sm = 768
	// Deprecated.
	SmKannada Sm = 16
	// Deprecated.
	SmKeyScript Sm = 22
	// Deprecated.
	SmKhmer Sm = 20
	// Deprecated.
	SmKlingon Sm = 32
	// Deprecated.
	SmKorean Sm = 3
	// Deprecated.
	SmLao Sm = 22
	// Deprecated.
	SmLaotian Sm = 22
	// Deprecated.
	SmLastByte Sm = 1
	// Deprecated.
	SmMalayalam Sm = 17
	// Deprecated.
	SmMiddleByte Sm = 2
	// Deprecated.
	SmMongolian Sm = 27
	// Deprecated.
	SmNotInstalled Sm = 0
	// Deprecated.
	SmOriya Sm = 12
	// Deprecated.
	SmPunctBlank Sm = 768
	// Deprecated.
	SmPunctGraphic Sm = 1280
	// Deprecated.
	SmPunctNormal Sm = 0
	// Deprecated.
	SmPunctNumber Sm = 256
	// Deprecated.
	SmPunctRepeat Sm = 1024
	// Deprecated.
	SmPunctSymbol Sm = 512
	// Deprecated.
	SmRSymbol Sm = 8
	// Deprecated.
	SmRegionCode Sm = 40
	// Deprecated.
	SmRoman Sm = 0
	// Deprecated.
	SmRussian Sm = 7
	// Deprecated.
	SmSimpChinese Sm = 25
	// Deprecated.
	SmSindhi Sm = 31
	// Deprecated.
	SmSingleByte Sm = 0
	// Deprecated.
	SmSinhalese Sm = 18
	// Deprecated.
	SmSlavic Sm = 29
	// Deprecated.
	SmSysScript Sm = 18
	// Deprecated.
	SmSystemScript Sm = -1
	// Deprecated.
	SmTamil Sm = 14
	// Deprecated.
	SmTelugu Sm = 15
	// Deprecated.
	SmThai Sm = 21
	// Deprecated.
	SmTibetan Sm = 26
	// Deprecated.
	SmTradChinese Sm = 2
	// Deprecated.
	SmUninterp Sm = 32
	// Deprecated.
	SmVietnamese Sm = 30
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
	case SmAllScripts:
		return "SmAllScripts"
	case SmAmharic:
		return "SmAmharic"
	case SmArabic:
		return "SmArabic"
	case SmArmenian:
		return "SmArmenian"
	case SmBadScript:
		return "SmBadScript"
	case SmBadVerb:
		return "SmBadVerb"
	case SmBengali:
		return "SmBengali"
	case SmBurmese:
		return "SmBurmese"
	case SmCentralEuroRoman:
		return "SmCentralEuroRoman"
	case SmChar1byte:
		return "SmChar1byte"
	case SmChar2byte:
		return "SmChar2byte"
	case SmCharFISGreek:
		return "SmCharFISGreek"
	case SmCharFISRussian:
		return "SmCharFISRussian"
	case SmCharRight:
		return "SmCharRight"
	case SmCharUpper:
		return "SmCharUpper"
	case SmCharVertical:
		return "SmCharVertical"
	case SmChinese:
		return "SmChinese"
	case SmCyrillic:
		return "SmCyrillic"
	case SmDevanagari:
		return "SmDevanagari"
	case SmExtArabic:
		return "SmExtArabic"
	case SmFISClassLvl2:
		return "SmFISClassLvl2"
	case SmFISClassUser:
		return "SmFISClassUser"
	case SmGeorgian:
		return "SmGeorgian"
	case SmGujarati:
		return "SmGujarati"
	case SmGurmukhi:
		return "SmGurmukhi"
	case SmJamoBogMoeum:
		return "SmJamoBogMoeum"
	case SmJapanese:
		return "SmJapanese"
	case SmKCHRCache:
		return "SmKCHRCache"
	case SmKannada:
		return "SmKannada"
	case SmKeyScript:
		return "SmKeyScript"
	case SmKhmer:
		return "SmKhmer"
	case SmKlingon:
		return "SmKlingon"
	case SmKorean:
		return "SmKorean"
	case SmMalayalam:
		return "SmMalayalam"
	case SmMongolian:
		return "SmMongolian"
	case SmOriya:
		return "SmOriya"
	case SmPunctGraphic:
		return "SmPunctGraphic"
	case SmPunctRepeat:
		return "SmPunctRepeat"
	case SmRSymbol:
		return "SmRSymbol"
	case SmRegionCode:
		return "SmRegionCode"
	case SmSimpChinese:
		return "SmSimpChinese"
	case SmSinhalese:
		return "SmSinhalese"
	case SmTamil:
		return "SmTamil"
	case SmTelugu:
		return "SmTelugu"
	case SmThai:
		return "SmThai"
	case SmTibetan:
		return "SmTibetan"
	case SmVietnamese:
		return "SmVietnamese"
	default:
		return fmt.Sprintf("Sm(%d)", e)
	}
}

type SmChar uint

const (
	// Deprecated.
	SmCharAscii SmChar = 1
	// Deprecated.
	SmCharBidirect SmChar = 8
	// Deprecated.
	SmCharBopomofo SmChar = 14
	// Deprecated.
	SmCharContextualLR SmChar = 9
	// Deprecated.
	SmCharEuro SmChar = 7
	// Deprecated.
	SmCharExtAscii SmChar = 7
	// Deprecated.
	SmCharFISGana SmChar = 3
	// Deprecated.
	SmCharFISIdeo SmChar = 4
	// Deprecated.
	SmCharFISKana SmChar = 2
	// Deprecated.
	SmCharGanaKana SmChar = 15
	// Deprecated.
	SmCharHangul SmChar = 12
	// Deprecated.
	SmCharHiragana SmChar = 3
	// Deprecated.
	SmCharIdeographic SmChar = 4
	// Deprecated.
	SmCharJamo SmChar = 13
	// Deprecated.
	SmCharKatakana SmChar = 2
	// Deprecated.
	SmCharNonContextualLR SmChar = 10
	// Deprecated.
	SmCharPunct SmChar = 0
	// Deprecated.
	SmCharTwoByteGreek SmChar = 5
	// Deprecated.
	SmCharTwoByteRussian SmChar = 6
)

func (e SmChar) String() string {
	switch e {
	case SmCharAscii:
		return "SmCharAscii"
	case SmCharBidirect:
		return "SmCharBidirect"
	case SmCharBopomofo:
		return "SmCharBopomofo"
	case SmCharContextualLR:
		return "SmCharContextualLR"
	case SmCharEuro:
		return "SmCharEuro"
	case SmCharFISGana:
		return "SmCharFISGana"
	case SmCharFISIdeo:
		return "SmCharFISIdeo"
	case SmCharFISKana:
		return "SmCharFISKana"
	case SmCharGanaKana:
		return "SmCharGanaKana"
	case SmCharHangul:
		return "SmCharHangul"
	case SmCharJamo:
		return "SmCharJamo"
	case SmCharNonContextualLR:
		return "SmCharNonContextualLR"
	case SmCharPunct:
		return "SmCharPunct"
	case SmCharTwoByteGreek:
		return "SmCharTwoByteGreek"
	case SmCharTwoByteRussian:
		return "SmCharTwoByteRussian"
	default:
		return fmt.Sprintf("SmChar(%d)", e)
	}
}

type SmKey int

const (
	// Deprecated.
	SmKeyDisableKybdSwitch SmKey = -12
	// Deprecated.
	SmKeyDisableKybds SmKey = -6
	// Deprecated.
	SmKeyEnableKybds SmKey = -7
	// Deprecated.
	SmKeyNextInputMethod SmKey = -10
	// Deprecated.
	SmKeyNextKybd SmKey = -4
	// Deprecated.
	SmKeyNextScript SmKey = -1
	// Deprecated.
	SmKeyRoman SmKey = -17
	// Deprecated.
	SmKeySetDirLeftRight SmKey = -15
	// Deprecated.
	SmKeySetDirRightLeft SmKey = -16
	// Deprecated.
	SmKeySwapInputMethod SmKey = -11
	// Deprecated.
	SmKeySwapKybd SmKey = -5
	// Deprecated.
	SmKeySwapScript SmKey = -3
	// Deprecated.
	SmKeySysScript SmKey = -2
	// Deprecated.
	SmKeyToggleDirection SmKey = -9
	// Deprecated.
	SmKeyToggleInline SmKey = -8
)

func (e SmKey) String() string {
	switch e {
	case SmKeyDisableKybdSwitch:
		return "SmKeyDisableKybdSwitch"
	case SmKeyDisableKybds:
		return "SmKeyDisableKybds"
	case SmKeyEnableKybds:
		return "SmKeyEnableKybds"
	case SmKeyNextInputMethod:
		return "SmKeyNextInputMethod"
	case SmKeyNextKybd:
		return "SmKeyNextKybd"
	case SmKeyNextScript:
		return "SmKeyNextScript"
	case SmKeyRoman:
		return "SmKeyRoman"
	case SmKeySetDirLeftRight:
		return "SmKeySetDirLeftRight"
	case SmKeySetDirRightLeft:
		return "SmKeySetDirRightLeft"
	case SmKeySwapInputMethod:
		return "SmKeySwapInputMethod"
	case SmKeySwapKybd:
		return "SmKeySwapKybd"
	case SmKeySwapScript:
		return "SmKeySwapScript"
	case SmKeySysScript:
		return "SmKeySysScript"
	case SmKeyToggleDirection:
		return "SmKeyToggleDirection"
	case SmKeyToggleInline:
		return "SmKeyToggleInline"
	default:
		return fmt.Sprintf("SmKey(%d)", e)
	}
}

type SmKeyForceKeyScript uint

const (
	// Deprecated.
	SmKeyForceKeyScriptBit SmKeyForceKeyScript = 7
	// Deprecated.
	SmKeyForceKeyScriptMask SmKeyForceKeyScript = 128
)

func (e SmKeyForceKeyScript) String() string {
	switch e {
	case SmKeyForceKeyScriptBit:
		return "SmKeyForceKeyScriptBit"
	case SmKeyForceKeyScriptMask:
		return "SmKeyForceKeyScriptMask"
	default:
		return fmt.Sprintf("SmKeyForceKeyScript(%d)", e)
	}
}

type SmMask uint

const (
	// Deprecated.
	SmMaskAll SmMask = 0
	// Deprecated.
	SmMaskAscii SmMask = 1
	// Deprecated.
	SmMaskAscii1 SmMask = 4
	// Deprecated.
	SmMaskAscii2 SmMask = 8
	// Deprecated.
	SmMaskBopomofo2 SmMask = 1024
	// Deprecated.
	SmMaskGana2 SmMask = 128
	// Deprecated.
	SmMaskHangul2 SmMask = 256
	// Deprecated.
	SmMaskJamo2 SmMask = 512
	// Deprecated.
	SmMaskKana1 SmMask = 16
	// Deprecated.
	SmMaskKana2 SmMask = 32
	// Deprecated.
	SmMaskNative SmMask = 2
)

func (e SmMask) String() string {
	switch e {
	case SmMaskAll:
		return "SmMaskAll"
	case SmMaskAscii:
		return "SmMaskAscii"
	case SmMaskAscii1:
		return "SmMaskAscii1"
	case SmMaskAscii2:
		return "SmMaskAscii2"
	case SmMaskBopomofo2:
		return "SmMaskBopomofo2"
	case SmMaskGana2:
		return "SmMaskGana2"
	case SmMaskHangul2:
		return "SmMaskHangul2"
	case SmMaskJamo2:
		return "SmMaskJamo2"
	case SmMaskKana1:
		return "SmMaskKana1"
	case SmMaskKana2:
		return "SmMaskKana2"
	case SmMaskNative:
		return "SmMaskNative"
	default:
		return fmt.Sprintf("SmMask(%d)", e)
	}
}

type SmResrvErr int

const (
	SmBLFieldBad    SmResrvErr = -309
	SmBadBoardId    SmResrvErr = -319
	SmBadRefId      SmResrvErr = -330
	SmBadsList      SmResrvErr = -331
	SmBusErrTO      SmResrvErr = -320
	SmCodeRevErr    SmResrvErr = -333
	SmDisposePErr   SmResrvErr = -312
	SmFHBlkDispErr  SmResrvErr = -311
	SmFHBlockRdErr  SmResrvErr = -310
	SmGetPRErr      SmResrvErr = -314
	SmInitStatVErr  SmResrvErr = -316
	SmInitTblVErr   SmResrvErr = -317
	SmNoBoardId     SmResrvErr = -315
	SmNoBoardSRsrc  SmResrvErr = -313
	SmNoJmpTbl      SmResrvErr = -318
	SmReservedErr   SmResrvErr = -332
	SmReservedSlot  SmResrvErr = -318
	SmResrvErrValue SmResrvErr = -307
	SmUnExBusErr    SmResrvErr = -308
	SvDisabled      SmResrvErr = -32640
	SvTempDisable   SmResrvErr = -32768
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

type SmTrans uint

const (
	// Deprecated.
	SmTransAscii SmTrans = 0
	// Deprecated.
	SmTransAscii1 SmTrans = 2
	// Deprecated.
	SmTransAscii2 SmTrans = 3
	// Deprecated.
	SmTransBopomofo2 SmTrans = 10
	// Deprecated.
	SmTransCase SmTrans = 254
	// Deprecated.
	SmTransGana2 SmTrans = 7
	// Deprecated.
	SmTransHangul2 SmTrans = 8
	// Deprecated.
	SmTransHangulFormat SmTrans = 2
	// Deprecated.
	SmTransJamo2 SmTrans = 9
	// Deprecated.
	SmTransKana1 SmTrans = 4
	// Deprecated.
	SmTransKana2 SmTrans = 5
	// Deprecated.
	SmTransLower SmTrans = 16384
	// Deprecated.
	SmTransNative SmTrans = 1
	// Deprecated.
	SmTransPreDoubleByting SmTrans = 1
	// Deprecated.
	SmTransPreLowerCasing SmTrans = 2
	// Deprecated.
	SmTransRuleBaseFormat SmTrans = 1
	// Deprecated.
	SmTransSystem SmTrans = 255
	// Deprecated.
	SmTransUpper SmTrans = 32768
)

func (e SmTrans) String() string {
	switch e {
	case SmTransAscii:
		return "SmTransAscii"
	case SmTransAscii1:
		return "SmTransAscii1"
	case SmTransAscii2:
		return "SmTransAscii2"
	case SmTransBopomofo2:
		return "SmTransBopomofo2"
	case SmTransCase:
		return "SmTransCase"
	case SmTransGana2:
		return "SmTransGana2"
	case SmTransHangul2:
		return "SmTransHangul2"
	case SmTransJamo2:
		return "SmTransJamo2"
	case SmTransKana1:
		return "SmTransKana1"
	case SmTransKana2:
		return "SmTransKana2"
	case SmTransLower:
		return "SmTransLower"
	case SmTransNative:
		return "SmTransNative"
	case SmTransSystem:
		return "SmTransSystem"
	case SmTransUpper:
		return "SmTransUpper"
	default:
		return fmt.Sprintf("SmTrans(%d)", e)
	}
}

type SmUnicode uint

const (
	SmUnicodeScript SmUnicode = 126
)

func (e SmUnicode) String() string {
	switch e {
	case SmUnicodeScript:
		return "SmUnicodeScript"
	default:
		return fmt.Sprintf("SmUnicode(%d)", e)
	}
}

type SmWordSelectTable uint

const (
	// Deprecated.
	IuNumberPartsTable SmWordSelectTable = 2
	// Deprecated.
	IuUnTokenTable SmWordSelectTable = 3
	// Deprecated.
	IuWhiteSpaceList SmWordSelectTable = 4
	// Deprecated.
	IuWordSelectTable SmWordSelectTable = 0
	// Deprecated.
	IuWordWrapTable SmWordSelectTable = 1
	// Deprecated.
	SmNumberPartsTable SmWordSelectTable = 2
	// Deprecated.
	SmUnTokenTable SmWordSelectTable = 3
	// Deprecated.
	SmWhiteSpaceList SmWordSelectTable = 4
	// Deprecated.
	SmWordSelectTableValue SmWordSelectTable = 0
	// Deprecated.
	SmWordWrapTable SmWordSelectTable = 1
)

func (e SmWordSelectTable) String() string {
	switch e {
	case IuNumberPartsTable:
		return "IuNumberPartsTable"
	case IuUnTokenTable:
		return "IuUnTokenTable"
	case IuWhiteSpaceList:
		return "IuWhiteSpaceList"
	case IuWordSelectTable:
		return "IuWordSelectTable"
	case IuWordWrapTable:
		return "IuWordWrapTable"
	default:
		return fmt.Sprintf("SmWordSelectTable(%d)", e)
	}
}

type SmallDateBit int

const (
	GenCdevRangeBit   SmallDateBit = 27
	MaxDateField      SmallDateBit = 10
	SmallDateBitValue SmallDateBit = 31
	TogChar12HourBit  SmallDateBit = 30
	TogCharZCycleBit  SmallDateBit = 29
	TogDelta12HourBit SmallDateBit = 28
	ValidDateFields   SmallDateBit = -1
)

func (e SmallDateBit) String() string {
	switch e {
	case GenCdevRangeBit:
		return "GenCdevRangeBit"
	case MaxDateField:
		return "MaxDateField"
	case SmallDateBitValue:
		return "SmallDateBitValue"
	case TogChar12HourBit:
		return "TogChar12HourBit"
	case TogCharZCycleBit:
		return "TogCharZCycleBit"
	case TogDelta12HourBit:
		return "TogDelta12HourBit"
	case ValidDateFields:
		return "ValidDateFields"
	default:
		return fmt.Sprintf("SmallDateBit(%d)", e)
	}
}

type Smc uint

const (
	// Deprecated.
	SmcClassMask Smc = 3840
	// Deprecated.
	SmcDoubleMask Smc = 32768
	// Deprecated.
	SmcOrientationMask Smc = 4096
	// Deprecated.
	SmcReserved Smc = 240
	// Deprecated.
	SmcRightMask Smc = 8192
	// Deprecated.
	SmcTypeMask Smc = 15
	// Deprecated.
	SmcUpperMask Smc = 16384
)

func (e Smc) String() string {
	switch e {
	case SmcClassMask:
		return "SmcClassMask"
	case SmcDoubleMask:
		return "SmcDoubleMask"
	case SmcOrientationMask:
		return "SmcOrientationMask"
	case SmcReserved:
		return "SmcReserved"
	case SmcRightMask:
		return "SmcRightMask"
	case SmcTypeMask:
		return "SmcTypeMask"
	case SmcUpperMask:
		return "SmcUpperMask"
	default:
		return fmt.Sprintf("Smc(%d)", e)
	}
}

type Smf uint

const (
	// Deprecated.
	SmfDisableKeyScriptSyncValue Smf = 27
	// Deprecated.
	SmfDualCaret Smf = 30
	// Deprecated.
	SmfNameTagEnab Smf = 29
	// Deprecated.
	SmfShowIcon Smf = 31
	// Deprecated.
	SmfUseAssocFontInfo Smf = 28
)

func (e Smf) String() string {
	switch e {
	case SmfDisableKeyScriptSyncValue:
		return "SmfDisableKeyScriptSyncValue"
	case SmfDualCaret:
		return "SmfDualCaret"
	case SmfNameTagEnab:
		return "SmfNameTagEnab"
	case SmfShowIcon:
		return "SmfShowIcon"
	case SmfUseAssocFontInfo:
		return "SmfUseAssocFontInfo"
	default:
		return fmt.Sprintf("Smf(%d)", e)
	}
}

type SmfDisableKeyScriptSync uint

const (
	// Deprecated.
	SmfDisableKeyScriptSyncMask SmfDisableKeyScriptSync = 134217728
)

func (e SmfDisableKeyScriptSync) String() string {
	switch e {
	case SmfDisableKeyScriptSyncMask:
		return "SmfDisableKeyScriptSyncMask"
	default:
		return fmt.Sprintf("SmfDisableKeyScriptSync(%d)", e)
	}
}

type Smsf uint

const (
	// Deprecated.
	SmsfAutoInit Smsf = 6
	// Deprecated.
	SmsfB0Digits Smsf = 5
	// Deprecated.
	SmsfContext Smsf = 3
	// Deprecated.
	SmsfForms Smsf = 13
	// Deprecated.
	SmsfIntellCP Smsf = 0
	// Deprecated.
	SmsfLigatures Smsf = 14
	// Deprecated.
	SmsfNatCase Smsf = 2
	// Deprecated.
	SmsfNoForceFont Smsf = 4
	// Deprecated.
	SmsfReverse Smsf = 15
	// Deprecated.
	SmsfSingByte Smsf = 1
	// Deprecated.
	SmsfSynchUnstyledTE Smsf = 8
	// Deprecated.
	SmsfUnivExt Smsf = 7
)

func (e Smsf) String() string {
	switch e {
	case SmsfAutoInit:
		return "SmsfAutoInit"
	case SmsfB0Digits:
		return "SmsfB0Digits"
	case SmsfContext:
		return "SmsfContext"
	case SmsfForms:
		return "SmsfForms"
	case SmsfIntellCP:
		return "SmsfIntellCP"
	case SmsfLigatures:
		return "SmsfLigatures"
	case SmsfNatCase:
		return "SmsfNatCase"
	case SmsfNoForceFont:
		return "SmsfNoForceFont"
	case SmsfReverse:
		return "SmsfReverse"
	case SmsfSingByte:
		return "SmsfSingByte"
	case SmsfSynchUnstyledTE:
		return "SmsfSynchUnstyledTE"
	case SmsfUnivExt:
		return "SmsfUnivExt"
	default:
		return fmt.Sprintf("Smsf(%d)", e)
	}
}

type Sorts int

const (
	// Deprecated.
	SortsAfter Sorts = 1
	// Deprecated.
	SortsBefore Sorts = -1
	// Deprecated.
	SortsEqual Sorts = 0
)

func (e Sorts) String() string {
	switch e {
	case SortsAfter:
		return "SortsAfter"
	case SortsBefore:
		return "SortsBefore"
	case SortsEqual:
		return "SortsEqual"
	default:
		return fmt.Sprintf("Sorts(%d)", e)
	}
}

type StartupFolderIconResource int

const (
	ControlPanelFolderIconResource StartupFolderIconResource = -3976
	DropFolderIconResource         StartupFolderIconResource = -3979
	ExtensionsFolderIconResource   StartupFolderIconResource = -3973
	FontsFolderIconResource        StartupFolderIconResource = -3968
	FullTrashIconResource          StartupFolderIconResource = -3984
	MountedFolderIconResource      StartupFolderIconResource = -3977
	OwnedFolderIconResource        StartupFolderIconResource = -3980
	PreferencesFolderIconResource  StartupFolderIconResource = -3974
	PrintMonitorFolderIconResource StartupFolderIconResource = -3975
	SharedFolderIconResource       StartupFolderIconResource = -3978
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

type SystemCurLang int

const (
	CurrentCurLang     SystemCurLang = -4
	CurrentDefLang     SystemCurLang = -5
	ScriptCurLang      SystemCurLang = -6
	ScriptDefLang      SystemCurLang = -7
	SystemCurLangValue SystemCurLang = -2
	SystemDefLang      SystemCurLang = -3
)

func (e SystemCurLang) String() string {
	switch e {
	case CurrentCurLang:
		return "CurrentCurLang"
	case CurrentDefLang:
		return "CurrentDefLang"
	case ScriptCurLang:
		return "ScriptCurLang"
	case ScriptDefLang:
		return "ScriptDefLang"
	case SystemCurLangValue:
		return "SystemCurLangValue"
	case SystemDefLang:
		return "SystemDefLang"
	default:
		return fmt.Sprintf("SystemCurLang(%d)", e)
	}
}

type TeScrapSizeErr int

const (
	DriverHardwareGoneErr TeScrapSizeErr = -503
	HwParamErr            TeScrapSizeErr = -502
	TeScrapSizeErrValue   TeScrapSizeErr = -501
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
	TelAPattNotSupp       Tel = -10016
	TelAlreadyOpen        Tel = -10070
	TelAutoAnsNotOn       Tel = -10112
	TelBadAPattErr        Tel = -10015
	TelBadBearerType      Tel = -10058
	TelBadCAErr           Tel = -10003
	TelBadCodeResource    Tel = -10108
	TelBadDNDType         Tel = -10023
	TelBadDNErr           Tel = -10002
	TelBadDNType          Tel = -10050
	TelBadDisplayMode     Tel = -10062
	TelBadFeatureID       Tel = -10053
	TelBadFunction        Tel = -10091
	TelBadFwdType         Tel = -10054
	TelBadHTypeErr        Tel = -10010
	TelBadHandErr         Tel = -10004
	TelBadIndex           Tel = -10017
	TelBadIntExt          Tel = -10021
	TelBadIntercomID      Tel = -10052
	TelBadLevelErr        Tel = -10012
	TelBadPageID          Tel = -10051
	TelBadParkID          Tel = -10056
	TelBadPickupGroupID   Tel = -10055
	TelBadProcErr         Tel = -10005
	TelBadProcID          Tel = -10110
	TelBadRate            Tel = -10059
	TelBadSWErr           Tel = -10114
	TelBadSampleRate      Tel = -10115
	TelBadSelect          Tel = -10057
	TelBadStateErr        Tel = -10019
	TelBadTermErr         Tel = -10001
	TelBadVTypeErr        Tel = -10013
	TelCANotAcceptable    Tel = -10080
	TelCANotDeflectable   Tel = -10082
	TelCANotRejectable    Tel = -10081
	TelCAUnavail          Tel = -10006
	TelCBErr              Tel = -10046
	TelConfErr            Tel = -10042
	TelConfLimitErr       Tel = -10040
	TelConfLimitExceeded  Tel = -10047
	TelConfNoLimit        Tel = -10041
	TelConfRej            Tel = -10043
	TelDNDTypeNotSupp     Tel = -10024
	TelDNTypeNotSupp      Tel = -10060
	TelDetAlreadyOn       Tel = -10113
	TelDeviceNotFound     Tel = -10109
	TelDisplayModeNotSupp Tel = -10063
	TelFeatActive         Tel = -10032
	TelFeatNotAvail       Tel = -10031
	TelFeatNotSub         Tel = -10030
	TelFeatNotSupp        Tel = -10033
	TelFwdTypeNotSupp     Tel = -10061
	TelGenericError       Tel = -1
	TelHTypeNotSupp       Tel = -10011
	TelIndexNotSupp       Tel = -10018
	TelInitFailed         Tel = -10107
	TelIntExtNotSupp      Tel = -10022
	TelNoCallbackRef      Tel = -10064
	TelNoCommFolder       Tel = -10106
	TelNoErr              Tel = 0
	TelNoMemErr           Tel = -10007
	TelNoOpenErr          Tel = -10008
	TelNoSuchTool         Tel = -10102
	TelNoTools            Tel = 8
	TelNotEnoughdspBW     Tel = -10116
	TelPBErr              Tel = -10090
	TelStateNotSupp       Tel = -10020
	TelStillNeeded        Tel = -10071
	TelTermNotOpen        Tel = -10072
	TelTransferErr        Tel = -10044
	TelTransferRej        Tel = -10045
	TelUnknownErr         Tel = -10103
	TelVTypeNotSupp       Tel = -10014
	TelValidateFailed     Tel = -10111
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
	TextParserBadParamErr         TextParser = -5220
	TextParserBadParserObjectErr  TextParser = -5223
	TextParserBadTextEncodingErr  TextParser = -5227
	TextParserBadTextLanguageErr  TextParser = -5226
	TextParserBadTokenValueErr    TextParser = -5222
	TextParserNoMoreTextErr       TextParser = -5225
	TextParserNoMoreTokensErr     TextParser = -5229
	TextParserNoSuchTokenFoundErr TextParser = -5228
	TextParserObjectNotFoundErr   TextParser = -5221
	TextParserParamErr            TextParser = -5224
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
	ThemeBadCursorIndexErr           Theme = -30565
	ThemeBadTextColorErr             Theme = -30563
	ThemeHasNoAccentsErr             Theme = -30564
	ThemeInvalidBrushErr             Theme = -30560
	ThemeMonitorDepthNotSupportedErr Theme = -30567
	ThemeNoAppropriateBrushErr       Theme = -30568
	ThemeProcessNotRegisteredErr     Theme = -30562
	ThemeProcessRegisteredErr        Theme = -30561
	ThemeScriptFontNotFoundErr       Theme = -30566
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
	ThreadNotFoundErr    Thread = -618
	ThreadProtocolErr    Thread = -619
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

type ThreadBadAppContext int

const (
	ThreadBadAppContextErr ThreadBadAppContext = -616
)

func (e ThreadBadAppContext) String() string {
	switch e {
	case ThreadBadAppContextErr:
		return "ThreadBadAppContextErr"
	default:
		return fmt.Sprintf("ThreadBadAppContext(%d)", e)
	}
}

type TimeCycle24 uint

const (
	Century          TimeCycle24 = 128
	DayLdingZ        TimeCycle24 = 32
	HrLeadingZ       TimeCycle24 = 128
	LongDay          TimeCycle24 = 0
	LongMonth        TimeCycle24 = 2
	LongWeek         TimeCycle24 = 1
	LongYear         TimeCycle24 = 3
	MinLeadingZ      TimeCycle24 = 64
	MntLdingZ        TimeCycle24 = 64
	SecLeadingZ      TimeCycle24 = 32
	SupDay           TimeCycle24 = 1
	SupMonth         TimeCycle24 = 4
	SupWeek          TimeCycle24 = 2
	SupYear          TimeCycle24 = 8
	TimeCycle12      TimeCycle24 = 255
	TimeCycle24Value TimeCycle24 = 0
	TimeCycleZero    TimeCycle24 = 1
	ZeroCycle        TimeCycle24 = 1
)

func (e TimeCycle24) String() string {
	switch e {
	case Century:
		return "Century"
	case DayLdingZ:
		return "DayLdingZ"
	case LongDay:
		return "LongDay"
	case LongMonth:
		return "LongMonth"
	case LongWeek:
		return "LongWeek"
	case LongYear:
		return "LongYear"
	case MinLeadingZ:
		return "MinLeadingZ"
	case SupMonth:
		return "SupMonth"
	case SupYear:
		return "SupYear"
	case TimeCycle12:
		return "TimeCycle12"
	default:
		return fmt.Sprintf("TimeCycle24(%d)", e)
	}
}

type Toggle uint

const (
	ToggleBadChar    Toggle = 4
	ToggleBadDelta   Toggle = 3
	ToggleBadField   Toggle = 2
	ToggleBadNum     Toggle = 6
	ToggleErr3       Toggle = 7
	ToggleErr4       Toggle = 8
	ToggleErr5       Toggle = 9
	ToggleOK         Toggle = 1
	ToggleOutOfRange Toggle = 7
	ToggleUndefined  Toggle = 0
	ToggleUnknown    Toggle = 5
)

func (e Toggle) String() string {
	switch e {
	case ToggleBadChar:
		return "ToggleBadChar"
	case ToggleBadDelta:
		return "ToggleBadDelta"
	case ToggleBadField:
		return "ToggleBadField"
	case ToggleBadNum:
		return "ToggleBadNum"
	case ToggleErr3:
		return "ToggleErr3"
	case ToggleErr4:
		return "ToggleErr4"
	case ToggleErr5:
		return "ToggleErr5"
	case ToggleOK:
		return "ToggleOK"
	case ToggleUndefined:
		return "ToggleUndefined"
	case ToggleUnknown:
		return "ToggleUnknown"
	default:
		return fmt.Sprintf("Toggle(%d)", e)
	}
}

type TokLeftQuote uint

const (
	CurNumberPartsVersion TokLeftQuote = 1
	TokDecPoint           TokLeftQuote = 14
	TokEMinus             TokLeftQuote = 16
	TokEPlus              TokLeftQuote = 15
	TokEscape             TokLeftQuote = 13
	TokLeadPlacer         TokLeftQuote = 3
	TokLeader             TokLeftQuote = 4
	TokLeftQuoteValue     TokLeftQuote = 1
	TokMaxSymbols         TokLeftQuote = 31
	TokMinusSign          TokLeftQuote = 9
	TokNonLeader          TokLeftQuote = 5
	TokPercent            TokLeftQuote = 7
	TokPlusSign           TokLeftQuote = 8
	TokReserved           TokLeftQuote = 11
	TokRightQuote         TokLeftQuote = 2
	TokSeparator          TokLeftQuote = 12
	TokThousands          TokLeftQuote = 10
	TokZeroLead           TokLeftQuote = 6
)

func (e TokLeftQuote) String() string {
	switch e {
	case CurNumberPartsVersion:
		return "CurNumberPartsVersion"
	case TokDecPoint:
		return "TokDecPoint"
	case TokEMinus:
		return "TokEMinus"
	case TokEPlus:
		return "TokEPlus"
	case TokEscape:
		return "TokEscape"
	case TokLeadPlacer:
		return "TokLeadPlacer"
	case TokLeader:
		return "TokLeader"
	case TokMaxSymbols:
		return "TokMaxSymbols"
	case TokMinusSign:
		return "TokMinusSign"
	case TokNonLeader:
		return "TokNonLeader"
	case TokPercent:
		return "TokPercent"
	case TokPlusSign:
		return "TokPlusSign"
	case TokReserved:
		return "TokReserved"
	case TokRightQuote:
		return "TokRightQuote"
	case TokSeparator:
		return "TokSeparator"
	case TokThousands:
		return "TokThousands"
	case TokZeroLead:
		return "TokZeroLead"
	default:
		return fmt.Sprintf("TokLeftQuote(%d)", e)
	}
}

type Token int

const (
	// Deprecated.
	Token1Quote Token = 52
	// Deprecated.
	Token2Equal Token = 38
	// Deprecated.
	Token2Quote Token = 51
	// Deprecated.
	TokenAlpha Token = 4
	// Deprecated.
	TokenAltNum Token = 11
	// Deprecated.
	TokenAltReal Token = 13
	// Deprecated.
	TokenAmpersand Token = 57
	// Deprecated.
	TokenAsterisk Token = 26
	// Deprecated.
	TokenAtSign Token = 58
	// Deprecated.
	TokenBackSlash Token = 30
	// Deprecated.
	TokenBar Token = 59
	// Deprecated.
	TokenCapPi Token = 66
	// Deprecated.
	TokenCaret Token = 55
	// Deprecated.
	TokenCenterDot Token = 78
	// Deprecated.
	TokenColon Token = 68
	// Deprecated.
	TokenColonEqual Token = 39
	// Deprecated.
	TokenComma Token = 45
	// Deprecated.
	TokenDivide Token = 27
	// Deprecated.
	TokenDollar Token = 70
	// Deprecated.
	TokenEllipsis Token = 77
	// Deprecated.
	TokenEmpty Token = -1
	// Deprecated.
	TokenEqual Token = 33
	// Deprecated.
	TokenEscape Token = 10
	// Deprecated.
	TokenExclam Token = 43
	// Deprecated.
	TokenExclamEqual Token = 42
	// Deprecated.
	TokenFraction Token = 72
	// Deprecated.
	TokenGreat Token = 32
	// Deprecated.
	TokenGreatEqual1 Token = 37
	// Deprecated.
	TokenGreatEqual2 Token = 36
	// Deprecated.
	TokenHash Token = 69
	// Deprecated.
	TokenInfinity Token = 67
	// Deprecated.
	TokenIntegral Token = 64
	// Deprecated.
	TokenIntl Token = 4
	// Deprecated.
	TokenIntlCurrency Token = 73
	// Deprecated.
	TokenLeft1Quote Token = 49
	// Deprecated.
	TokenLeft2Quote Token = 47
	// Deprecated.
	TokenLeftBracket Token = 18
	// Deprecated.
	TokenLeftComment Token = 7
	// Deprecated.
	TokenLeftCurly Token = 20
	// Deprecated.
	TokenLeftEnclose Token = 22
	// Deprecated.
	TokenLeftLit Token = 2
	// Deprecated.
	TokenLeftParen Token = 16
	// Deprecated.
	TokenLeftSingGuillemet Token = 74
	// Deprecated.
	TokenLess Token = 31
	// Deprecated.
	TokenLessEqual1 Token = 35
	// Deprecated.
	TokenLessEqual2 Token = 34
	// Deprecated.
	TokenLessGreat Token = 41
	// Deprecated.
	TokenLiteral Token = 9
	// Deprecated.
	TokenMicro Token = 65
	// Deprecated.
	TokenMinus Token = 25
	// Deprecated.
	TokenNewLine Token = 6
	// Deprecated.
	TokenNil Token = 127
	// Deprecated.
	TokenNoBreakSpace Token = 71
	// Deprecated.
	TokenNotEqual Token = 40
	// Deprecated.
	TokenNumeric Token = 5
	// Deprecated.
	TokenPerThousand Token = 76
	// Deprecated.
	TokenPercent Token = 54
	// Deprecated.
	TokenPeriod Token = 46
	// Deprecated.
	TokenPi Token = 61
	// Deprecated.
	TokenPlus Token = 24
	// Deprecated.
	TokenPlusMinus Token = 28
	// Deprecated.
	TokenQuestion Token = 60
	// Deprecated.
	TokenRealNum Token = 12
	// Deprecated.
	TokenReserve1 Token = 14
	// Deprecated.
	TokenReserve2 Token = 15
	// Deprecated.
	TokenRight1Quote Token = 50
	// Deprecated.
	TokenRight2Quote Token = 48
	// Deprecated.
	TokenRightBracket Token = 19
	// Deprecated.
	TokenRightComment Token = 8
	// Deprecated.
	TokenRightCurly Token = 21
	// Deprecated.
	TokenRightEnclose Token = 23
	// Deprecated.
	TokenRightLit Token = 3
	// Deprecated.
	TokenRightParen Token = 17
	// Deprecated.
	TokenRightSingGuillemet Token = 75
	// Deprecated.
	TokenRoot Token = 62
	// Deprecated.
	TokenSemicolon Token = 53
	// Deprecated.
	TokenSigma Token = 63
	// Deprecated.
	TokenSlash Token = 29
	// Deprecated.
	TokenTilde Token = 44
	// Deprecated.
	TokenUnderline Token = 56
	// Deprecated.
	TokenUnknown Token = 0
	// Deprecated.
	TokenWhite Token = 1
)

func (e Token) String() string {
	switch e {
	case Token1Quote:
		return "Token1Quote"
	case Token2Equal:
		return "Token2Equal"
	case Token2Quote:
		return "Token2Quote"
	case TokenAlpha:
		return "TokenAlpha"
	case TokenAltNum:
		return "TokenAltNum"
	case TokenAltReal:
		return "TokenAltReal"
	case TokenAmpersand:
		return "TokenAmpersand"
	case TokenAsterisk:
		return "TokenAsterisk"
	case TokenAtSign:
		return "TokenAtSign"
	case TokenBackSlash:
		return "TokenBackSlash"
	case TokenBar:
		return "TokenBar"
	case TokenCapPi:
		return "TokenCapPi"
	case TokenCaret:
		return "TokenCaret"
	case TokenCenterDot:
		return "TokenCenterDot"
	case TokenColon:
		return "TokenColon"
	case TokenColonEqual:
		return "TokenColonEqual"
	case TokenComma:
		return "TokenComma"
	case TokenDivide:
		return "TokenDivide"
	case TokenDollar:
		return "TokenDollar"
	case TokenEllipsis:
		return "TokenEllipsis"
	case TokenEmpty:
		return "TokenEmpty"
	case TokenEqual:
		return "TokenEqual"
	case TokenEscape:
		return "TokenEscape"
	case TokenExclam:
		return "TokenExclam"
	case TokenExclamEqual:
		return "TokenExclamEqual"
	case TokenFraction:
		return "TokenFraction"
	case TokenGreat:
		return "TokenGreat"
	case TokenGreatEqual1:
		return "TokenGreatEqual1"
	case TokenGreatEqual2:
		return "TokenGreatEqual2"
	case TokenHash:
		return "TokenHash"
	case TokenInfinity:
		return "TokenInfinity"
	case TokenIntegral:
		return "TokenIntegral"
	case TokenIntlCurrency:
		return "TokenIntlCurrency"
	case TokenLeft1Quote:
		return "TokenLeft1Quote"
	case TokenLeft2Quote:
		return "TokenLeft2Quote"
	case TokenLeftBracket:
		return "TokenLeftBracket"
	case TokenLeftComment:
		return "TokenLeftComment"
	case TokenLeftCurly:
		return "TokenLeftCurly"
	case TokenLeftEnclose:
		return "TokenLeftEnclose"
	case TokenLeftLit:
		return "TokenLeftLit"
	case TokenLeftParen:
		return "TokenLeftParen"
	case TokenLeftSingGuillemet:
		return "TokenLeftSingGuillemet"
	case TokenLess:
		return "TokenLess"
	case TokenLessEqual1:
		return "TokenLessEqual1"
	case TokenLessEqual2:
		return "TokenLessEqual2"
	case TokenLessGreat:
		return "TokenLessGreat"
	case TokenLiteral:
		return "TokenLiteral"
	case TokenMicro:
		return "TokenMicro"
	case TokenMinus:
		return "TokenMinus"
	case TokenNewLine:
		return "TokenNewLine"
	case TokenNil:
		return "TokenNil"
	case TokenNoBreakSpace:
		return "TokenNoBreakSpace"
	case TokenNotEqual:
		return "TokenNotEqual"
	case TokenNumeric:
		return "TokenNumeric"
	case TokenPerThousand:
		return "TokenPerThousand"
	case TokenPercent:
		return "TokenPercent"
	case TokenPeriod:
		return "TokenPeriod"
	case TokenPi:
		return "TokenPi"
	case TokenPlus:
		return "TokenPlus"
	case TokenPlusMinus:
		return "TokenPlusMinus"
	case TokenQuestion:
		return "TokenQuestion"
	case TokenRealNum:
		return "TokenRealNum"
	case TokenReserve1:
		return "TokenReserve1"
	case TokenReserve2:
		return "TokenReserve2"
	case TokenRight1Quote:
		return "TokenRight1Quote"
	case TokenRight2Quote:
		return "TokenRight2Quote"
	case TokenRightBracket:
		return "TokenRightBracket"
	case TokenRightComment:
		return "TokenRightComment"
	case TokenRightCurly:
		return "TokenRightCurly"
	case TokenRightEnclose:
		return "TokenRightEnclose"
	case TokenRightLit:
		return "TokenRightLit"
	case TokenRightParen:
		return "TokenRightParen"
	case TokenRightSingGuillemet:
		return "TokenRightSingGuillemet"
	case TokenRoot:
		return "TokenRoot"
	case TokenSemicolon:
		return "TokenSemicolon"
	case TokenSigma:
		return "TokenSigma"
	case TokenSlash:
		return "TokenSlash"
	case TokenTilde:
		return "TokenTilde"
	case TokenUnderline:
		return "TokenUnderline"
	case TokenUnknown:
		return "TokenUnknown"
	case TokenWhite:
		return "TokenWhite"
	default:
		return fmt.Sprintf("Token(%d)", e)
	}
}

type TokenOK uint

const (
	// Deprecated.
	BadDelim TokenOK = 3
	// Deprecated.
	BadEnding TokenOK = 4
	// Deprecated.
	Crash TokenOK = 5
	// Deprecated.
	StringOverflow TokenOK = 2
	// Deprecated.
	TokenOKValue TokenOK = 0
	// Deprecated.
	TokenOverflow TokenOK = 1
)

func (e TokenOK) String() string {
	switch e {
	case BadDelim:
		return "BadDelim"
	case BadEnding:
		return "BadEnding"
	case Crash:
		return "Crash"
	case StringOverflow:
		return "StringOverflow"
	case TokenOKValue:
		return "TokenOKValue"
	case TokenOverflow:
		return "TokenOverflow"
	default:
		return fmt.Sprintf("TokenOK(%d)", e)
	}
}

type Tsm int

const (
	TsmAlreadyRegisteredErr            Tsm = -2503
	TsmCantChangeForcedClassStateErr   Tsm = -2530
	TsmCantOpenComponentErr            Tsm = -2509
	TsmComponentAlreadyOpenErr         Tsm = -2515
	TsmComponentNoErr                  Tsm = 0
	TsmComponentPropertyNotFoundErr    Tsm = -2532
	TsmComponentPropertyUnsupportedErr Tsm = -2531
	TsmDefaultIsNotInputMethodErr      Tsm = -2524
	TsmDocNotActiveErr                 Tsm = -2507
	TsmDocPropertyBufferTooSmallErr    Tsm = -2529
	TsmDocPropertyNotFoundErr          Tsm = -2528
	TsmDocumentOpenErr                 Tsm = -2511
	TsmInputMethodIsOldErr             Tsm = -2516
	TsmInputMethodNotFoundErr          Tsm = -2501
	TsmInputModeChangeFailedErr        Tsm = -2533
	TsmInvalidContext                  Tsm = -2520
	TsmInvalidDocIDErr                 Tsm = -2505
	TsmNeverRegisteredErr              Tsm = -2504
	TsmNoHandler                       Tsm = -2521
	TsmNoMoreTokens                    Tsm = -2522
	TsmNoOpenTSErr                     Tsm = -2508
	TsmNoStem                          Tsm = -2523
	TsmNotAnAppErr                     Tsm = -2502
	TsmScriptHasNoIMErr                Tsm = -2517
	TsmTSHasNoMenuErr                  Tsm = -2513
	TsmTSMDocBusyErr                   Tsm = -2506
	TsmTSNotOpenErr                    Tsm = -2514
	TsmTextServiceNotFoundErr          Tsm = -2510
	TsmUnknownErr                      Tsm = -2519
	TsmUnsupScriptLanguageErr          Tsm = -2500
	TsmUnsupportedTypeErr              Tsm = -2518
	TsmUseInputWindowErr               Tsm = -2512
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
	TypeAEText   Type = 't'<<24 | 'T'<<16 | 'X'<<8 | 'T' // 'tTXT'
	// TypeAbsoluteOrdinal: Specifies a descriptor whose data consists of one of the constants `kAEFirst`, `kAEMiddle`, `kAELast`, `kAEAny`, or `kAEAll`, which are described in AEDisposeToken.
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
	TypeArc            Type = 'c'<<24 | 'a'<<16 | 'r'<<8 | 'c' // 'carc'
	TypeBest           Type = 'b'<<24 | 'e'<<16 | 's'<<8 | 't' // 'best'
	TypeBookmarkData   Type = 'b'<<24 | 'm'<<16 | 'r'<<8 | 'k' // 'bmrk'
	// TypeBoolean: Boolean value—single byte with value 0 or 1.
	TypeBoolean        Type = 'b'<<24 | 'o'<<16 | 'o'<<8 | 'l' // 'bool'
	TypeCFAbsoluteTime Type = 'c'<<24 | 'f'<<16 | 'a'<<8 | 't' // 'cfat'
	// TypeCString: C string—Mac OS Roman characters followed by a NULL byte.
	TypeCString     Type = 'c'<<24 | 's'<<16 | 't'<<8 | 'r' // 'cstr'
	TypeCell        Type = 'c'<<24 | 'c'<<16 | 'e'<<8 | 'l' // 'ccel'
	TypeCentimeters Type = 'c'<<24 | 'm'<<16 | 't'<<8 | 'r' // 'cmtr'
	// TypeChar: # Discussion
	TypeChar       Type = 'T'<<24 | 'E'<<16 | 'X'<<8 | 'T' // 'TEXT'
	TypeClassInfo  Type = 'g'<<24 | 'c'<<16 | 'l'<<8 | 'i' // 'gcli'
	TypeColorTable Type = 'c'<<24 | 'l'<<16 | 'r'<<8 | 't' // 'clrt'
	TypeColumn     Type = 'c'<<24 | 'c'<<16 | 'o'<<8 | 'l' // 'ccol'
	// TypeCompDescriptor: Specifies a comparison descriptor.
	TypeCompDescriptor  Type = 'c'<<24 | 'm'<<16 | 'p'<<8 | 'd' // 'cmpd'
	TypeCubicCentimeter Type = 'c'<<24 | 'c'<<16 | 'm'<<8 | 't' // 'ccmt'
	TypeCubicFeet       Type = 'c'<<24 | 'f'<<16 | 'e'<<8 | 't' // 'cfet'
	TypeCubicInches     Type = 'c'<<24 | 'u'<<16 | 'i'<<8 | 'n' // 'cuin'
	TypeCubicMeters     Type = 'c'<<24 | 'm'<<16 | 'e'<<8 | 't' // 'cmet'
	TypeCubicYards      Type = 'c'<<24 | 'y'<<16 | 'r'<<8 | 'd' // 'cyrd'
	// TypeCurrentContainer: Specifies a container for an element that demarcates one boundary in a range.
	TypeCurrentContainer Type = 'c'<<24 | 'c'<<16 | 'n'<<8 | 't' // 'ccnt'
	TypeDashStyle        Type = 't'<<24 | 'd'<<16 | 'a'<<8 | 's' // 'tdas'
	TypeData             Type = 't'<<24 | 'd'<<16 | 't'<<8 | 'a' // 'tdta'
	// TypeDecimalStruct: Decimal.
	TypeDecimalStruct Type = 'd'<<24 | 'e'<<16 | 'c'<<8 | 'm' // 'decm'
	TypeDegreesC      Type = 'd'<<24 | 'e'<<16 | 'g'<<8 | 'c' // 'degc'
	TypeDegreesF      Type = 'd'<<24 | 'e'<<16 | 'g'<<8 | 'f' // 'degf'
	TypeDegreesK      Type = 'd'<<24 | 'e'<<16 | 'g'<<8 | 'k' // 'degk'
	TypeDrawingArea   Type = 'c'<<24 | 'd'<<16 | 'r'<<8 | 'w' // 'cdrw'
	TypeEPS           Type = 'E'<<24 | 'P'<<16 | 'S'<<8 | ' ' // 'EPS '
	TypeElemInfo      Type = 'e'<<24 | 'l'<<16 | 'i'<<8 | 'n' // 'elin'
	// TypeEncodedString: Styled Unicode text.
	TypeEncodedString Type = 'e'<<24 | 'n'<<16 | 'c'<<8 | 's' // 'encs'
	// TypeEnumerated: Enumerated data.
	TypeEnumerated  Type = 'e'<<24 | 'n'<<16 | 'u'<<8 | 'm' // 'enum'
	TypeEnumeration Type = 'e'<<24 | 'n'<<16 | 'u'<<8 | 'm' // 'enum'
	TypeEventInfo   Type = 'e'<<24 | 'v'<<16 | 'i'<<8 | 'n' // 'evin'
	TypeEventRecord Type = 'e'<<24 | 'v'<<16 | 'r'<<8 | 'c' // 'evrc'
	// TypeFSRef: File system reference.
	TypeFSRef Type = 'f'<<24 | 's'<<16 | 'r'<<8 | 'f' // 'fsrf'
	// TypeFalse: [FALSE] Boolean value.
	TypeFalse Type = 'f'<<24 | 'a'<<16 | 'l'<<8 | 's' // 'fals'
	TypeFeet  Type = 'f'<<24 | 'e'<<16 | 'e'<<8 | 't' // 'feet'
	// TypeFileURL: A file URL.
	TypeFileURL        Type = 'f'<<24 | 'u'<<16 | 'r'<<8 | 'l' // 'furl'
	TypeFinderWindow   Type = 'f'<<24 | 'w'<<16 | 'i'<<8 | 'n' // 'fwin'
	TypeFixed          Type = 'f'<<24 | 'i'<<16 | 'x'<<8 | 'd' // 'fixd'
	TypeFixedPoint     Type = 'f'<<24 | 'p'<<16 | 'n'<<8 | 't' // 'fpnt'
	TypeFixedRectangle Type = 'f'<<24 | 'r'<<16 | 'c'<<8 | 't' // 'frct'
	TypeGIF            Type = 'G'<<24 | 'I'<<16 | 'F'<<8 | 'f' // 'GIFf'
	TypeGallons        Type = 'g'<<24 | 'a'<<16 | 'l'<<8 | 'n' // 'galn'
	TypeGrams          Type = 'g'<<24 | 'r'<<16 | 'a'<<8 | 'm' // 'gram'
	TypeGraphicLine    Type = 'g'<<24 | 'l'<<16 | 'i'<<8 | 'n' // 'glin'
	TypeGraphicText    Type = 'c'<<24 | 'g'<<16 | 't'<<8 | 'x' // 'cgtx'
	TypeGroupedGraphic Type = 'c'<<24 | 'p'<<16 | 'i'<<8 | 'c' // 'cpic'
	// TypeIEEE32BitFloatingPoint: 32-bit floating point value.
	TypeIEEE32BitFloatingPoint Type = 's'<<24 | 'i'<<16 | 'n'<<8 | 'g' // 'sing'
	// TypeIEEE64BitFloatingPoint: 64-bit floating point value.
	TypeIEEE64BitFloatingPoint Type = 'd'<<24 | 'o'<<16 | 'u'<<8 | 'b' // 'doub'
	TypeISO8601DateTime        Type = 'i'<<24 | 's'<<16 | 'o'<<8 | 't' // 'isot'
	TypeInches                 Type = 'i'<<24 | 'n'<<16 | 'c'<<8 | 'h' // 'inch'
	// TypeIndexDescriptor: Specifies a descriptor whose data indicates an indexed position within a range of values.
	TypeIndexDescriptor Type = 'i'<<24 | 'n'<<16 | 'd'<<8 | 'e' // 'inde'
	TypeInsertionLoc    Type = 'i'<<24 | 'n'<<16 | 's'<<8 | 'l' // 'insl'
	// TypeIntlText: For important information, see the Version Notes section of the typeUnicodeText enum.
	TypeIntlText        Type = 'i'<<24 | 't'<<16 | 'x'<<8 | 't' // 'itxt'
	TypeIntlWritingCode Type = 'i'<<24 | 'n'<<16 | 't'<<8 | 'l' // 'intl'
	TypeJPEG            Type = 'J'<<24 | 'P'<<16 | 'E'<<8 | 'G' // 'JPEG'
	// TypeKernelProcessID: Indicates a descriptor containing a UNIX process ID.
	TypeKernelProcessID Type = 'k'<<24 | 'p'<<16 | 'i'<<8 | 'd' // 'kpid'
	// TypeKeyword: Apple event keyword.
	TypeKeyword    Type = 'k'<<24 | 'e'<<16 | 'y'<<8 | 'w' // 'keyw'
	TypeKilograms  Type = 'k'<<24 | 'g'<<16 | 'r'<<8 | 'm' // 'kgrm'
	TypeKilometers Type = 'k'<<24 | 'm'<<16 | 't'<<8 | 'r' // 'kmtr'
	TypeLiters     Type = 'l'<<24 | 'i'<<16 | 't'<<8 | 'r' // 'litr'
	// TypeLogicalDescriptor: Specifies a logical descriptor.
	TypeLogicalDescriptor  Type = 'l'<<24 | 'o'<<16 | 'g'<<8 | 'i' // 'logi'
	TypeLongDateTime       Type = 'l'<<24 | 'd'<<16 | 't'<<8 | ' ' // 'ldt '
	TypeLongFixed          Type = 'l'<<24 | 'f'<<16 | 'x'<<8 | 'd' // 'lfxd'
	TypeLongFixedPoint     Type = 'l'<<24 | 'f'<<16 | 'p'<<8 | 't' // 'lfpt'
	TypeLongFixedRectangle Type = 'l'<<24 | 'f'<<16 | 'r'<<8 | 'c' // 'lfrc'
	TypeLongPoint          Type = 'l'<<24 | 'p'<<16 | 'n'<<8 | 't' // 'lpnt'
	TypeLongRectangle      Type = 'l'<<24 | 'r'<<16 | 'c'<<8 | 't' // 'lrct'
	// TypeMachPort: Indicates a descriptor that specifies a Mach port.
	TypeMachPort   Type = 'p'<<24 | 'o'<<16 | 'r'<<8 | 't' // 'port'
	TypeMachineLoc Type = 'm'<<24 | 'L'<<16 | 'o'<<8 | 'c' // 'mLoc'
	TypeMeters     Type = 'm'<<24 | 'e'<<16 | 't'<<8 | 'r' // 'metr'
	TypeMiles      Type = 'm'<<24 | 'i'<<16 | 'l'<<8 | 'e' // 'mile'
	// TypeNull: A null data storage pointer.
	TypeNull Type = 'n'<<24 | 'u'<<16 | 'l'<<8 | 'l' // 'null'
	// TypeOSLTokenList: Specifies a descriptor whose data consists of a list of tokens.
	TypeOSLTokenList Type = 'o'<<24 | 's'<<16 | 't'<<8 | 'l' // 'ostl'
	// TypeObjectBeingExamined: # Discussion
	TypeObjectBeingExamined Type = 'e'<<24 | 'x'<<16 | 'm'<<8 | 'n' // 'exmn'
	// TypeObjectSpecifier: Specifies a descriptor used with the `keyAEContainer` keyword in a keyword-specified descriptor.
	TypeObjectSpecifier Type = 'o'<<24 | 'b'<<16 | 'j'<<8 | ' ' // 'obj '
	TypeOunces          Type = 'o'<<24 | 'z'<<16 | 's'<<8 | ' ' // 'ozs '
	TypeOval            Type = 'c'<<24 | 'o'<<16 | 'v'<<8 | 'l' // 'covl'
	// TypePString: Pascal string—unsigned length byte followed by Mac OS Roman characters.
	TypePString     Type = 'p'<<24 | 's'<<16 | 't'<<8 | 'r' // 'pstr'
	TypeParamInfo   Type = 'p'<<24 | 'm'<<16 | 'i'<<8 | 'n' // 'pmin'
	TypePict        Type = 'P'<<24 | 'I'<<16 | 'C'<<8 | 'T' // 'PICT'
	TypePixMapMinus Type = 't'<<24 | 'p'<<16 | 'm'<<8 | 'm' // 'tpmm'
	TypePixelMap    Type = 'c'<<24 | 'p'<<16 | 'i'<<8 | 'x' // 'cpix'
	TypePolygon     Type = 'c'<<24 | 'p'<<16 | 'g'<<8 | 'n' // 'cpgn'
	TypePounds      Type = 'l'<<24 | 'b'<<16 | 's'<<8 | ' ' // 'lbs '
	// TypeProcessSerialNumber: A process serial number.
	TypeProcessSerialNumber Type = 'p'<<24 | 's'<<16 | 'n'<<8 | ' ' // 'psn '
	TypePropInfo            Type = 'p'<<24 | 'i'<<16 | 'n'<<8 | 'f' // 'pinf'
	// TypeProperty: Apple event object property.
	TypeProperty    Type = 'p'<<24 | 'r'<<16 | 'o'<<8 | 'p' // 'prop'
	TypePtr         Type = 'p'<<24 | 't'<<16 | 'r'<<8 | ' ' // 'ptr '
	TypeQDPoint     Type = 'Q'<<24 | 'D'<<16 | 'p'<<8 | 't' // 'QDpt'
	TypeQDRectangle Type = 'q'<<24 | 'd'<<16 | 'r'<<8 | 't' // 'qdrt'
	TypeQDRegion    Type = 'Q'<<24 | 'r'<<16 | 'g'<<8 | 'n' // 'Qrgn'
	TypeQuarts      Type = 'q'<<24 | 'r'<<16 | 't'<<8 | 's' // 'qrts'
	TypeRGB16       Type = 't'<<24 | 'r'<<16 | '1'<<8 | '6' // 'tr16'
	TypeRGB96       Type = 't'<<24 | 'r'<<16 | '9'<<8 | '6' // 'tr96'
	TypeRGBColor    Type = 'c'<<24 | 'R'<<16 | 'G'<<8 | 'B' // 'cRGB'
	// TypeRangeDescriptor: # Discussion
	TypeRangeDescriptor Type = 'r'<<24 | 'a'<<16 | 'n'<<8 | 'g' // 'rang'
	TypeRectangle       Type = 'c'<<24 | 'r'<<16 | 'e'<<8 | 'c' // 'crec'
	// TypeRelativeDescriptor: Specifies a descriptor whose data consists of one of the constants `kAENext` or `kAEPrevious`, which are described in AEDisposeToken.
	TypeRelativeDescriptor Type = 'r'<<24 | 'e'<<16 | 'l'<<8 | ' ' // 'rel '
	TypeRotation           Type = 't'<<24 | 'r'<<16 | 'o'<<8 | 't' // 'trot'
	TypeRoundedRectangle   Type = 'c'<<24 | 'r'<<16 | 'r'<<8 | 'c' // 'crrc'
	TypeRow                Type = 'c'<<24 | 'r'<<16 | 'o'<<8 | 'w' // 'crow'
	// TypeSInt16: 16-bit signed integer.
	TypeSInt16 Type = 's'<<24 | 'h'<<16 | 'o'<<8 | 'r' // 'shor'
	// TypeSInt32: 32-bit signed integer.
	TypeSInt32 Type = 'l'<<24 | 'o'<<16 | 'n'<<8 | 'g' // 'long'
	// TypeSInt64: 64-bit signed integer.
	TypeSInt64      Type = 'c'<<24 | 'o'<<16 | 'm'<<8 | 'p' // 'comp'
	TypeScrapStyles Type = 's'<<24 | 't'<<16 | 'y'<<8 | 'l' // 'styl'
	TypeScript      Type = 's'<<24 | 'c'<<16 | 'p'<<8 | 't' // 'scpt'
	// TypeSectionH: Handle to a section record.
	TypeSectionH         Type = 's'<<24 | 'e'<<16 | 'c'<<8 | 't' // 'sect'
	TypeSquareFeet       Type = 's'<<24 | 'q'<<16 | 'f'<<8 | 't' // 'sqft'
	TypeSquareKilometers Type = 's'<<24 | 'q'<<16 | 'k'<<8 | 'm' // 'sqkm'
	TypeSquareMeters     Type = 's'<<24 | 'q'<<16 | 'r'<<8 | 'm' // 'sqrm'
	TypeSquareMiles      Type = 's'<<24 | 'q'<<16 | 'm'<<8 | 'i' // 'sqmi'
	TypeSquareYards      Type = 's'<<24 | 'q'<<16 | 'y'<<8 | 'd' // 'sqyd'
	// TypeStyledText: # Discussion
	TypeStyledText Type = 'S'<<24 | 'T'<<16 | 'X'<<8 | 'T' // 'STXT'
	// TypeStyledUnicodeText: Styled Unicode text.
	TypeStyledUnicodeText Type = 's'<<24 | 'u'<<16 | 't'<<8 | 'x' // 'sutx'
	TypeSuiteInfo         Type = 's'<<24 | 'u'<<16 | 'i'<<8 | 'n' // 'suin'
	TypeTIFF              Type = 'T'<<24 | 'I'<<16 | 'F'<<8 | 'F' // 'TIFF'
	TypeTable             Type = 'c'<<24 | 't'<<16 | 'b'<<8 | 'l' // 'ctbl'
	TypeTextStyles        Type = 't'<<24 | 's'<<16 | 't'<<8 | 'y' // 'tsty'
	// TypeToken: Specifies a descriptor whose data storage pointer refers to a structure of type AEDisposeToken.
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
	TypeVersion     Type = 'v'<<24 | 'e'<<16 | 'r'<<8 | 's' // 'vers'
	// TypeWildCard: Matches any type.
	TypeWildCard Type = '*'<<24 | '*'<<16 | '*'<<8 | '*' // '****'
	TypeYards    Type = 'y'<<24 | 'a'<<16 | 'r'<<8 | 'd' // 'yard'
)

func (e Type) String() string {
	switch e {
	case Type128BitFloatingPoint:
		return "Type128BitFloatingPoint"
	case TypeAEList:
		return "TypeAEList"
	case TypeAERecord:
		return "TypeAERecord"
	case TypeAEText:
		return "TypeAEText"
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
	case TypeArc:
		return "TypeArc"
	case TypeBest:
		return "TypeBest"
	case TypeBookmarkData:
		return "TypeBookmarkData"
	case TypeBoolean:
		return "TypeBoolean"
	case TypeCFAbsoluteTime:
		return "TypeCFAbsoluteTime"
	case TypeCString:
		return "TypeCString"
	case TypeCell:
		return "TypeCell"
	case TypeCentimeters:
		return "TypeCentimeters"
	case TypeChar:
		return "TypeChar"
	case TypeClassInfo:
		return "TypeClassInfo"
	case TypeColorTable:
		return "TypeColorTable"
	case TypeColumn:
		return "TypeColumn"
	case TypeCompDescriptor:
		return "TypeCompDescriptor"
	case TypeCubicCentimeter:
		return "TypeCubicCentimeter"
	case TypeCubicFeet:
		return "TypeCubicFeet"
	case TypeCubicInches:
		return "TypeCubicInches"
	case TypeCubicMeters:
		return "TypeCubicMeters"
	case TypeCubicYards:
		return "TypeCubicYards"
	case TypeCurrentContainer:
		return "TypeCurrentContainer"
	case TypeDashStyle:
		return "TypeDashStyle"
	case TypeData:
		return "TypeData"
	case TypeDecimalStruct:
		return "TypeDecimalStruct"
	case TypeDegreesC:
		return "TypeDegreesC"
	case TypeDegreesF:
		return "TypeDegreesF"
	case TypeDegreesK:
		return "TypeDegreesK"
	case TypeDrawingArea:
		return "TypeDrawingArea"
	case TypeEPS:
		return "TypeEPS"
	case TypeElemInfo:
		return "TypeElemInfo"
	case TypeEncodedString:
		return "TypeEncodedString"
	case TypeEnumerated:
		return "TypeEnumerated"
	case TypeEventInfo:
		return "TypeEventInfo"
	case TypeEventRecord:
		return "TypeEventRecord"
	case TypeFSRef:
		return "TypeFSRef"
	case TypeFalse:
		return "TypeFalse"
	case TypeFeet:
		return "TypeFeet"
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
	case TypeGIF:
		return "TypeGIF"
	case TypeGallons:
		return "TypeGallons"
	case TypeGrams:
		return "TypeGrams"
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
	case TypeInches:
		return "TypeInches"
	case TypeIndexDescriptor:
		return "TypeIndexDescriptor"
	case TypeInsertionLoc:
		return "TypeInsertionLoc"
	case TypeIntlText:
		return "TypeIntlText"
	case TypeIntlWritingCode:
		return "TypeIntlWritingCode"
	case TypeJPEG:
		return "TypeJPEG"
	case TypeKernelProcessID:
		return "TypeKernelProcessID"
	case TypeKeyword:
		return "TypeKeyword"
	case TypeKilograms:
		return "TypeKilograms"
	case TypeKilometers:
		return "TypeKilometers"
	case TypeLiters:
		return "TypeLiters"
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
	case TypeMachPort:
		return "TypeMachPort"
	case TypeMachineLoc:
		return "TypeMachineLoc"
	case TypeMeters:
		return "TypeMeters"
	case TypeMiles:
		return "TypeMiles"
	case TypeNull:
		return "TypeNull"
	case TypeOSLTokenList:
		return "TypeOSLTokenList"
	case TypeObjectBeingExamined:
		return "TypeObjectBeingExamined"
	case TypeObjectSpecifier:
		return "TypeObjectSpecifier"
	case TypeOunces:
		return "TypeOunces"
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
	case TypePounds:
		return "TypePounds"
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
	case TypeQuarts:
		return "TypeQuarts"
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
	case TypeSquareFeet:
		return "TypeSquareFeet"
	case TypeSquareKilometers:
		return "TypeSquareKilometers"
	case TypeSquareMeters:
		return "TypeSquareMeters"
	case TypeSquareMiles:
		return "TypeSquareMiles"
	case TypeSquareYards:
		return "TypeSquareYards"
	case TypeStyledText:
		return "TypeStyledText"
	case TypeStyledUnicodeText:
		return "TypeStyledUnicodeText"
	case TypeSuiteInfo:
		return "TypeSuiteInfo"
	case TypeTIFF:
		return "TypeTIFF"
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
	case TypeVersion:
		return "TypeVersion"
	case TypeWildCard:
		return "TypeWildCard"
	case TypeYards:
		return "TypeYards"
	default:
		return fmt.Sprintf("Type(%d)", e)
	}
}

type TypeApplicationBundleI uint

const (
	// TypeApplicationBundleID: Indicates a descriptor containing UTF-8 characters that specify the bundle ID of an application.
	TypeApplicationBundleID TypeApplicationBundleI = 'b'<<24 | 'u'<<16 | 'n'<<8 | 'd' // 'bund'
)

func (e TypeApplicationBundleI) String() string {
	switch e {
	case TypeApplicationBundleID:
		return "TypeApplicationBundleID"
	default:
		return fmt.Sprintf("TypeApplicationBundleI(%d)", e)
	}
}

type TypeAudit uint

const (
	TypeAuditToken TypeAudit = 't'<<24 | 'o'<<16 | 'k'<<8 | 'n' // 'tokn'
)

func (e TypeAudit) String() string {
	switch e {
	case TypeAuditToken:
		return "TypeAuditToken"
	default:
		return fmt.Sprintf("TypeAudit(%d)", e)
	}
}

type TypeCF uint

const (
	TypeCFArrayRef                   TypeCF = 'c'<<24 | 'f'<<16 | 'a'<<8 | 'r' // 'cfar'
	TypeCFAttributedStringRef        TypeCF = 'c'<<24 | 'f'<<16 | 'a'<<8 | 's' // 'cfas'
	TypeCFBooleanRef                 TypeCF = 'c'<<24 | 'f'<<16 | 't'<<8 | 'f' // 'cftf'
	TypeCFDictionaryRef              TypeCF = 'c'<<24 | 'f'<<16 | 'd'<<8 | 'c' // 'cfdc'
	TypeCFMutableArrayRef            TypeCF = 'c'<<24 | 'f'<<16 | 'm'<<8 | 'a' // 'cfma'
	TypeCFMutableAttributedStringRef TypeCF = 'c'<<24 | 'f'<<16 | 'a'<<8 | 'a' // 'cfaa'
	TypeCFMutableDictionaryRef       TypeCF = 'c'<<24 | 'f'<<16 | 'm'<<8 | 'd' // 'cfmd'
	TypeCFMutableStringRef           TypeCF = 'c'<<24 | 'f'<<16 | 'm'<<8 | 's' // 'cfms'
	TypeCFNumberRef                  TypeCF = 'c'<<24 | 'f'<<16 | 'n'<<8 | 'b' // 'cfnb'
	TypeCFStringRef                  TypeCF = 'c'<<24 | 'f'<<16 | 's'<<8 | 't' // 'cfst'
	TypeCFTypeRef                    TypeCF = 'c'<<24 | 'f'<<16 | 't'<<8 | 'y' // 'cfty'
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

type TypeHI uint

const (
	TypeHIMenu   TypeHI = 'm'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'mobj'
	TypeHIWindow TypeHI = 'w'<<24 | 'o'<<16 | 'b'<<8 | 'j' // 'wobj'
)

func (e TypeHI) String() string {
	switch e {
	case TypeHIMenu:
		return "TypeHIMenu"
	case TypeHIWindow:
		return "TypeHIWindow"
	default:
		return fmt.Sprintf("TypeHI(%d)", e)
	}
}

type TypeReplyPort uint

const (
	TypeReplyPortAttr TypeReplyPort = 'r'<<24 | 'e'<<16 | 'p'<<8 | 'p' // 'repp'
)

func (e TypeReplyPort) String() string {
	switch e {
	case TypeReplyPortAttr:
		return "TypeReplyPortAttr"
	default:
		return fmt.Sprintf("TypeReplyPort(%d)", e)
	}
}

type TypeUT uint

const (
	// TypeUTF16ExternalRepresentation: # Discussion
	TypeUTF16ExternalRepresentation TypeUT = 'u'<<24 | 't'<<16 | '1'<<8 | '6' // 'ut16'
	// TypeUTF8Text: 8-bit Unicode (UTF-8 encoding).
	TypeUTF8Text TypeUT = 'u'<<24 | 't'<<16 | 'f'<<8 | '8' // 'utf8'
)

func (e TypeUT) String() string {
	switch e {
	case TypeUTF16ExternalRepresentation:
		return "TypeUTF16ExternalRepresentation"
	case TypeUTF8Text:
		return "TypeUTF8Text"
	default:
		return fmt.Sprintf("TypeUT(%d)", e)
	}
}

type TypeWhoseDescriptor uint

const (
	// FormWhose: # Discussion
	FormWhose                TypeWhoseDescriptor = 'w'<<24 | 'h'<<16 | 'o'<<8 | 's' // 'whos'
	KeyAEIndex               TypeWhoseDescriptor = 'k'<<24 | 'i'<<16 | 'd'<<8 | 'x' // 'kidx'
	KeyAETest                TypeWhoseDescriptor = 'k'<<24 | 't'<<16 | 's'<<8 | 't' // 'ktst'
	KeyAEWhoseRangeStart     TypeWhoseDescriptor = 'w'<<24 | 's'<<16 | 't'<<8 | 'r' // 'wstr'
	KeyAEWhoseRangeStop      TypeWhoseDescriptor = 'w'<<24 | 's'<<16 | 't'<<8 | 'p' // 'wstp'
	TypeWhoseDescriptorValue TypeWhoseDescriptor = 'w'<<24 | 'h'<<16 | 'o'<<8 | 's' // 'whos'
	TypeWhoseRange           TypeWhoseDescriptor = 'w'<<24 | 'r'<<16 | 'n'<<8 | 'g' // 'wrng'
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

type Upp uint

const (
	// Deprecated.
	UppCallComponentCanDoProcInfo Upp = 752
	// Deprecated.
	UppCallComponentCloseProcInfo Upp = 1008
	// Deprecated.
	UppCallComponentGetMPWorkFunctionProcInfo Upp = 4080
	// Deprecated.
	UppCallComponentGetPublicResourceProcInfo Upp = 15344
	// Deprecated.
	UppCallComponentOpenProcInfo Upp = 1008
	// Deprecated.
	UppCallComponentRegisterProcInfo Upp = 240
	// Deprecated.
	UppCallComponentTargetProcInfo Upp = 1008
	// Deprecated.
	UppCallComponentUnregisterProcInfo Upp = 240
	// Deprecated.
	UppCallComponentVersionProcInfo Upp = 240
	// Deprecated.
	UppComponentFunctionImplementedProcInfo Upp = 752
	// Deprecated.
	UppComponentSetTargetProcInfo Upp = 1008
	// Deprecated.
	UppGetComponentVersionProcInfo Upp = 240
)

func (e Upp) String() string {
	switch e {
	case UppCallComponentCanDoProcInfo:
		return "UppCallComponentCanDoProcInfo"
	case UppCallComponentCloseProcInfo:
		return "UppCallComponentCloseProcInfo"
	case UppCallComponentGetMPWorkFunctionProcInfo:
		return "UppCallComponentGetMPWorkFunctionProcInfo"
	case UppCallComponentGetPublicResourceProcInfo:
		return "UppCallComponentGetPublicResourceProcInfo"
	case UppCallComponentRegisterProcInfo:
		return "UppCallComponentRegisterProcInfo"
	default:
		return fmt.Sprintf("Upp(%d)", e)
	}
}

type Use uint

const (
	// Deprecated.
	UseATalk Use = 1
	// Deprecated.
	UseAsync Use = 2
	// Deprecated.
	UseExtClk Use = 3
	// Deprecated.
	UseFree Use = 0
	// Deprecated.
	UseMIDI Use = 4
)

func (e Use) String() string {
	switch e {
	case UseATalk:
		return "UseATalk"
	case UseAsync:
		return "UseAsync"
	case UseExtClk:
		return "UseExtClk"
	case UseFree:
		return "UseFree"
	case UseMIDI:
		return "UseMIDI"
	default:
		return fmt.Sprintf("Use(%d)", e)
	}
}

type VLckdErr int

const (
	BadMDBErr     VLckdErr = -60
	BadMovErr     VLckdErr = -122
	DirNFErr      VLckdErr = -120
	DupFNErr      VLckdErr = -48
	ExtFSErr      VLckdErr = -58
	FBsyErr       VLckdErr = -47
	FsRnErr       VLckdErr = -59
	GfpErr        VLckdErr = -52
	NoMacDskErr   VLckdErr = -57
	NsDrvErr      VLckdErr = -56
	OpWrErr       VLckdErr = -49
	PermErr       VLckdErr = -54
	RfNumErr      VLckdErr = -51
	TmwdoErr      VLckdErr = -121
	VLckdErrValue VLckdErr = -46
	VolGoneErr    VLckdErr = -124
	VolOffLinErr  VLckdErr = -53
	VolOnLinErr   VLckdErr = -55
	WrPermErr     VLckdErr = -61
	WrgVolTypErr  VLckdErr = -123
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

type Ver uint

const (
	// Deprecated.
	VerAfrikaans Ver = 102
	// Deprecated.
	VerAlternateGr Ver = 64
	// Deprecated.
	VerArabia Ver = 16
	// Deprecated.
	VerArabic Ver = 16
	// Deprecated.
	VerArmenia Ver = 84
	// Deprecated.
	VerArmenian Ver = 84
	// Deprecated.
	VerAustralia Ver = 15
	// Deprecated.
	VerAustria Ver = 92
	// Deprecated.
	VerAustriaGerman Ver = 92
	// Deprecated.
	VerBelarus Ver = 61
	// Deprecated.
	VerBelgiumLux Ver = 6
	// Deprecated.
	VerBelgiumLuxPoint Ver = 27
	// Deprecated.
	VerBengali Ver = 60
	// Deprecated.
	VerBhutan Ver = 83
	// Deprecated.
	VerBrazil Ver = 71
	// Deprecated.
	VerBreton Ver = 77
	// Deprecated.
	VerBritain Ver = 2
	// Deprecated.
	VerBrittany Ver = 77
	// Deprecated.
	VerBulgaria Ver = 72
	// Deprecated.
	VerByeloRussian Ver = 61
	// Deprecated.
	VerCanadaComma Ver = 28
	// Deprecated.
	VerCanadaPoint Ver = 29
	// Deprecated.
	VerCatalonia Ver = 73
	// Deprecated.
	VerChina Ver = 52
	// Deprecated.
	VerCroatia Ver = 68
	// Deprecated.
	VerCyprus Ver = 23
	// Deprecated.
	VerCzech Ver = 56
	// Deprecated.
	VerDenmark Ver = 9
	// Deprecated.
	VerEastAsiaGeneric Ver = 58
	// Deprecated.
	VerEngCanada Ver = 82
	// Deprecated.
	VerEsperanto Ver = 103
	// Deprecated.
	VerEstonia Ver = 44
	// Deprecated.
	VerFaeroeIsl Ver = 47
	// Deprecated.
	VerFarEastGeneric Ver = 58
	// Deprecated.
	VerFaroeIsl Ver = 47
	// Deprecated.
	VerFinland Ver = 17
	// Deprecated.
	VerFlemish Ver = 6
	// Deprecated.
	VerFlemishPoint Ver = 27
	// Deprecated.
	VerFrBelgium Ver = 98
	// Deprecated.
	VerFrBelgiumLux Ver = 6
	// Deprecated.
	VerFrCanada Ver = 11
	// Deprecated.
	VerFrSwiss Ver = 18
	// Deprecated.
	VerFrance Ver = 1
	// Deprecated.
	VerFrenchUniversal Ver = 91
	// Deprecated.
	VerGenericFE Ver = 58
	// Deprecated.
	VerGeorgia Ver = 85
	// Deprecated.
	VerGeorgian Ver = 85
	// Deprecated.
	VerGermanReformed Ver = 70
	// Deprecated.
	VerGermany Ver = 3
	// Deprecated.
	VerGrSwiss Ver = 19
	// Deprecated.
	VerGreece Ver = 20
	// Deprecated.
	VerGreeceAlt Ver = 64
	// Deprecated.
	VerGreecePoly Ver = 40
	// Deprecated.
	VerGreekAncient Ver = 40
	// Deprecated.
	VerGreenland Ver = 107
	// Deprecated.
	VerGujarati Ver = 94
	// Deprecated.
	VerHungary Ver = 43
	// Deprecated.
	VerIceland Ver = 21
	// Deprecated.
	VerIndia Ver = 33
	// Deprecated.
	VerIndiaHindi Ver = 33
	// Deprecated.
	VerIndiaUrdu Ver = 96
	// Deprecated.
	VerInternational Ver = 37
	// Deprecated.
	VerIran Ver = 48
	// Deprecated.
	VerIreland Ver = 50
	// Deprecated.
	VerIrelandEnglish Ver = 108
	// Deprecated.
	VerIrishGaelicScript Ver = 81
	// Deprecated.
	VerIsrael Ver = 13
	// Deprecated.
	VerItalianSwiss Ver = 36
	// Deprecated.
	VerItaly Ver = 4
	// Deprecated.
	VerJapan Ver = 14
	// Deprecated.
	VerKorea Ver = 51
	// Deprecated.
	VerLapland Ver = 46
	// Deprecated.
	VerLatvia Ver = 45
	// Deprecated.
	VerLithuania Ver = 41
	// Deprecated.
	VerMacedonia Ver = 67
	// Deprecated.
	VerMacedonian Ver = 67
	// Deprecated.
	VerMagyar Ver = 59
	// Deprecated.
	VerMalta Ver = 22
	// Deprecated.
	VerManxGaelic Ver = 76
	// Deprecated.
	VerMarathi Ver = 104
	// Deprecated.
	VerMultilingual Ver = 74
	// Deprecated.
	VerNepal Ver = 106
	// Deprecated.
	VerNetherlands Ver = 5
	// Deprecated.
	VerNetherlandsComma Ver = 26
	// Deprecated.
	VerNorway Ver = 12
	// Deprecated.
	VerNunavut Ver = 78
	// Deprecated.
	VerNynorsk Ver = 101
	// Deprecated.
	VerPakistan Ver = 34
	// Deprecated.
	VerPakistanUrdu Ver = 34
	// Deprecated.
	VerPoland Ver = 42
	// Deprecated.
	VerPortugal Ver = 10
	// Deprecated.
	VerPunjabi Ver = 95
	// Deprecated.
	VerRomania Ver = 39
	// Deprecated.
	VerRumania Ver = 39
	// Deprecated.
	VerRussia Ver = 49
	// Deprecated.
	VerSami Ver = 46
	// Deprecated.
	VerScottishGaelic Ver = 75
	// Deprecated.
	VerScriptGeneric Ver = 55
	// Deprecated.
	VerSerbia Ver = 65
	// Deprecated.
	VerSerbian Ver = 65
	// Deprecated.
	VerSingapore Ver = 100
	// Deprecated.
	VerSlovak Ver = 57
	// Deprecated.
	VerSlovenia Ver = 66
	// Deprecated.
	VerSlovenian Ver = 66
	// Deprecated.
	VerSpLatinAmerica Ver = 86
	// Deprecated.
	VerSpain Ver = 8
	// Deprecated.
	VerSweden Ver = 7
	// Deprecated.
	VerTaiwan Ver = 53
	// Deprecated.
	VerThailand Ver = 54
	// Deprecated.
	VerTibet Ver = 105
	// Deprecated.
	VerTibetan Ver = 105
	// Deprecated.
	VerTonga Ver = 88
	// Deprecated.
	VerTurkey Ver = 24
	// Deprecated.
	VerTurkishModified Ver = 35
	// Deprecated.
	VerUS Ver = 0
	// Deprecated.
	VerUkraine Ver = 62
	// Deprecated.
	VerUkrania Ver = 62
	// Deprecated.
	VerUzbek Ver = 99
	// Deprecated.
	VerVietnam Ver = 97
	// Deprecated.
	VerWales Ver = 79
	// Deprecated.
	VerWelsh Ver = 79
	// Deprecated.
	VerYugoCroatian Ver = 25
	// Deprecated.
	VerYugoslavia Ver = 25
	// Deprecated.
	VervariantDenmark Ver = 32
	// Deprecated.
	VervariantNorway Ver = 31
	// Deprecated.
	VervariantPortugal Ver = 30
)

func (e Ver) String() string {
	switch e {
	case VerAfrikaans:
		return "VerAfrikaans"
	case VerAlternateGr:
		return "VerAlternateGr"
	case VerArabia:
		return "VerArabia"
	case VerArmenia:
		return "VerArmenia"
	case VerAustralia:
		return "VerAustralia"
	case VerAustria:
		return "VerAustria"
	case VerBelarus:
		return "VerBelarus"
	case VerBelgiumLux:
		return "VerBelgiumLux"
	case VerBelgiumLuxPoint:
		return "VerBelgiumLuxPoint"
	case VerBengali:
		return "VerBengali"
	case VerBhutan:
		return "VerBhutan"
	case VerBrazil:
		return "VerBrazil"
	case VerBreton:
		return "VerBreton"
	case VerBritain:
		return "VerBritain"
	case VerBulgaria:
		return "VerBulgaria"
	case VerCanadaComma:
		return "VerCanadaComma"
	case VerCanadaPoint:
		return "VerCanadaPoint"
	case VerCatalonia:
		return "VerCatalonia"
	case VerChina:
		return "VerChina"
	case VerCroatia:
		return "VerCroatia"
	case VerCyprus:
		return "VerCyprus"
	case VerCzech:
		return "VerCzech"
	case VerDenmark:
		return "VerDenmark"
	case VerEastAsiaGeneric:
		return "VerEastAsiaGeneric"
	case VerEngCanada:
		return "VerEngCanada"
	case VerEsperanto:
		return "VerEsperanto"
	case VerEstonia:
		return "VerEstonia"
	case VerFaeroeIsl:
		return "VerFaeroeIsl"
	case VerFinland:
		return "VerFinland"
	case VerFrBelgium:
		return "VerFrBelgium"
	case VerFrCanada:
		return "VerFrCanada"
	case VerFrSwiss:
		return "VerFrSwiss"
	case VerFrance:
		return "VerFrance"
	case VerFrenchUniversal:
		return "VerFrenchUniversal"
	case VerGeorgia:
		return "VerGeorgia"
	case VerGermanReformed:
		return "VerGermanReformed"
	case VerGermany:
		return "VerGermany"
	case VerGrSwiss:
		return "VerGrSwiss"
	case VerGreece:
		return "VerGreece"
	case VerGreecePoly:
		return "VerGreecePoly"
	case VerGreenland:
		return "VerGreenland"
	case VerGujarati:
		return "VerGujarati"
	case VerHungary:
		return "VerHungary"
	case VerIceland:
		return "VerIceland"
	case VerIndia:
		return "VerIndia"
	case VerIndiaUrdu:
		return "VerIndiaUrdu"
	case VerInternational:
		return "VerInternational"
	case VerIran:
		return "VerIran"
	case VerIreland:
		return "VerIreland"
	case VerIrelandEnglish:
		return "VerIrelandEnglish"
	case VerIrishGaelicScript:
		return "VerIrishGaelicScript"
	case VerIsrael:
		return "VerIsrael"
	case VerItalianSwiss:
		return "VerItalianSwiss"
	case VerItaly:
		return "VerItaly"
	case VerJapan:
		return "VerJapan"
	case VerKorea:
		return "VerKorea"
	case VerLapland:
		return "VerLapland"
	case VerLatvia:
		return "VerLatvia"
	case VerLithuania:
		return "VerLithuania"
	case VerMacedonia:
		return "VerMacedonia"
	case VerMagyar:
		return "VerMagyar"
	case VerMalta:
		return "VerMalta"
	case VerManxGaelic:
		return "VerManxGaelic"
	case VerMarathi:
		return "VerMarathi"
	case VerMultilingual:
		return "VerMultilingual"
	case VerNepal:
		return "VerNepal"
	case VerNetherlands:
		return "VerNetherlands"
	case VerNetherlandsComma:
		return "VerNetherlandsComma"
	case VerNorway:
		return "VerNorway"
	case VerNunavut:
		return "VerNunavut"
	case VerNynorsk:
		return "VerNynorsk"
	case VerPakistan:
		return "VerPakistan"
	case VerPoland:
		return "VerPoland"
	case VerPortugal:
		return "VerPortugal"
	case VerPunjabi:
		return "VerPunjabi"
	case VerRomania:
		return "VerRomania"
	case VerRussia:
		return "VerRussia"
	case VerScottishGaelic:
		return "VerScottishGaelic"
	case VerScriptGeneric:
		return "VerScriptGeneric"
	case VerSerbia:
		return "VerSerbia"
	case VerSingapore:
		return "VerSingapore"
	case VerSlovak:
		return "VerSlovak"
	case VerSlovenia:
		return "VerSlovenia"
	case VerSpLatinAmerica:
		return "VerSpLatinAmerica"
	case VerSpain:
		return "VerSpain"
	case VerSweden:
		return "VerSweden"
	case VerTaiwan:
		return "VerTaiwan"
	case VerThailand:
		return "VerThailand"
	case VerTibet:
		return "VerTibet"
	case VerTonga:
		return "VerTonga"
	case VerTurkey:
		return "VerTurkey"
	case VerTurkishModified:
		return "VerTurkishModified"
	case VerUS:
		return "VerUS"
	case VerUkraine:
		return "VerUkraine"
	case VerUzbek:
		return "VerUzbek"
	case VerVietnam:
		return "VerVietnam"
	case VerWales:
		return "VerWales"
	case VerYugoCroatian:
		return "VerYugoCroatian"
	case VervariantDenmark:
		return "VervariantDenmark"
	case VervariantNorway:
		return "VervariantNorway"
	case VervariantPortugal:
		return "VervariantPortugal"
	default:
		return fmt.Sprintf("Ver(%d)", e)
	}
}

type VerUnspecified uint

const (
	KTECResourceID      VerUnspecified = 128
	VerUnspecifiedValue VerUnspecified = 32767
)

func (e VerUnspecified) String() string {
	switch e {
	case KTECResourceID:
		return "KTECResourceID"
	case VerUnspecifiedValue:
		return "VerUnspecifiedValue"
	default:
		return fmt.Sprintf("VerUnspecified(%d)", e)
	}
}

type Vm int

const (
	VmAddressNotInFileViewErr    Vm = -647
	VmBadDriver                  Vm = -632
	VmBusyBackingFileErr         Vm = -642
	VmFileViewAccessErr          Vm = -645
	VmInvalidBackingFileIDErr    Vm = -640
	VmInvalidFileViewIDErr       Vm = -644
	VmInvalidOwningProcessErr    Vm = -648
	VmKernelMMUInitErr           Vm = -629
	VmMappingPrivilegesErr       Vm = -641
	VmMemLckdErr                 Vm = -631
	VmMorePhysicalThanVirtualErr Vm = -628
	VmNoMoreBackingFilesErr      Vm = -643
	VmNoMoreFileViewsErr         Vm = -646
	VmNoVectorErr                Vm = -633
	VmOffErr                     Vm = -630
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

type VolMount uint

const (
	VolMountChangedBit         VolMount = 14
	VolMountChangedMask        VolMount = 16384
	VolMountExtendedFlagsBit   VolMount = 7
	VolMountExtendedFlagsMask  VolMount = 128
	VolMountFSReservedMask     VolMount = 255
	VolMountInteractBit        VolMount = 15
	VolMountInteractMask       VolMount = 32768
	VolMountNoLoginMsgFlagBit  VolMount = 0
	VolMountNoLoginMsgFlagMask VolMount = 1
	VolMountSysReservedMask    VolMount = 65280
)

func (e VolMount) String() string {
	switch e {
	case VolMountChangedBit:
		return "VolMountChangedBit"
	case VolMountChangedMask:
		return "VolMountChangedMask"
	case VolMountExtendedFlagsBit:
		return "VolMountExtendedFlagsBit"
	case VolMountExtendedFlagsMask:
		return "VolMountExtendedFlagsMask"
	case VolMountFSReservedMask:
		return "VolMountFSReservedMask"
	case VolMountInteractBit:
		return "VolMountInteractBit"
	case VolMountInteractMask:
		return "VolMountInteractMask"
	case VolMountNoLoginMsgFlagBit:
		return "VolMountNoLoginMsgFlagBit"
	case VolMountNoLoginMsgFlagMask:
		return "VolMountNoLoginMsgFlagMask"
	case VolMountSysReservedMask:
		return "VolMountSysReservedMask"
	default:
		return fmt.Sprintf("VolMount(%d)", e)
	}
}

type WrongApplicationPlatform int

const (
	AppVersionTooOld              WrongApplicationPlatform = -876
	NotAppropriateForClassic      WrongApplicationPlatform = -877
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
