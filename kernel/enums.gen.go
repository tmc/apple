// Code generated from Apple documentation for kernel. DO NOT EDIT.

package kernel

import (
	"fmt"
)

const AddPacketShift uint = 0

type AtaDeviceType uint

const (
	KATADeviceType AtaDeviceType = 0
	KATAPIDeviceType AtaDeviceType = 0
	KUnknownATADeviceType AtaDeviceType = 0
)

func (e AtaDeviceType) String() string {
	switch e {
	case KATADeviceType:
		return "KATADeviceType"
	default:
		return fmt.Sprintf("AtaDeviceType(%d)", e)
	}
}

type AtaSocketType uint

const (
	KInternalATASocket AtaSocketType = 0
	KInternalSATA AtaSocketType = 0
	KInternalSATA2 AtaSocketType = 0
	KMediaBaySocket AtaSocketType = 0
	KPCCardSocket AtaSocketType = 0
	KSATA2Bay AtaSocketType = 0
	KSATABay AtaSocketType = 0
	KUnknownSocket AtaSocketType = 0
)

func (e AtaSocketType) String() string {
	switch e {
	case KInternalATASocket:
		return "KInternalATASocket"
	default:
		return fmt.Sprintf("AtaSocketType(%d)", e)
	}
}

type BATABadBlock uint

const (
	MATAAddressNotFound BATABadBlock = 0
)

func (e BATABadBlock) String() string {
	switch e {
	case MATAAddressNotFound:
		return "MATAAddressNotFound"
	default:
		return fmt.Sprintf("BATABadBlock(%d)", e)
	}
}

type BATABusy uint

const (
	MATADataRequest BATABusy = 0
)

func (e BATABusy) String() string {
	switch e {
	case MATADataRequest:
		return "MATADataRequest"
	default:
		return fmt.Sprintf("BATABusy(%d)", e)
	}
}

const BATADCROne uint = 0

const BATAPIuseDMA uint = 0

type BTreeKeyLimits uint

const (
	KMaxKeyLength BTreeKeyLimits = 0
)

func (e BTreeKeyLimits) String() string {
	switch e {
	case KMaxKeyLength:
		return "KMaxKeyLength"
	default:
		return fmt.Sprintf("BTreeKeyLimits(%d)", e)
	}
}

type BacktraceFlagsT uint

const (
	BTF_KERN_INTERRUPTED BacktraceFlagsT = 0
)

func (e BacktraceFlagsT) String() string {
	switch e {
	case BTF_KERN_INTERRUPTED:
		return "BTF_KERN_INTERRUPTED"
	default:
		return fmt.Sprintf("BacktraceFlagsT(%d)", e)
	}
}

type BacktraceInfoT uint

const (
	BTI_64_BIT BacktraceInfoT = 0
)

func (e BacktraceInfoT) String() string {
	switch e {
	case BTI_64_BIT:
		return "BTI_64_BIT"
	default:
		return fmt.Sprintf("BacktraceInfoT(%d)", e)
	}
}

type BacktracePackT uint

const (
	BTP_KERN_OFFSET_32 BacktracePackT = 0
)

func (e BacktracePackT) String() string {
	switch e {
	case BTP_KERN_OFFSET_32:
		return "BTP_KERN_OFFSET_32"
	default:
		return fmt.Sprintf("BacktracePackT(%d)", e)
	}
}

type BidirectionalUndefined uint

const (
	BIDIRECTIONAL_ECHO_CANCELING_SPEAKERPHONE BidirectionalUndefined = 0
)

func (e BidirectionalUndefined) String() string {
	switch e {
	case BIDIRECTIONAL_ECHO_CANCELING_SPEAKERPHONE:
		return "BIDIRECTIONAL_ECHO_CANCELING_SPEAKERPHONE"
	default:
		return fmt.Sprintf("BidirectionalUndefined(%d)", e)
	}
}

type BpfModeDisabled uint

const (
	// BPF_MODE_OUTPUT: # Discussion
	BPF_MODE_OUTPUT BpfModeDisabled = 0
)

func (e BpfModeDisabled) String() string {
	switch e {
	case BPF_MODE_OUTPUT:
		return "BPF_MODE_OUTPUT"
	default:
		return fmt.Sprintf("BpfModeDisabled(%d)", e)
	}
}

type BpfTapDisable uint

const (
	BPF_TAP_DISABLE BpfTapDisable = 0
)

func (e BpfTapDisable) String() string {
	switch e {
	case BPF_TAP_DISABLE:
		return "BPF_TAP_DISABLE"
	default:
		return fmt.Sprintf("BpfTapDisable(%d)", e)
	}
}

type CacheTypeT uint

const (
	L1I CacheTypeT = 0
	L3U CacheTypeT = 0
)

func (e CacheTypeT) String() string {
	switch e {
	case L1I:
		return "L1I"
	default:
		return fmt.Sprintf("CacheTypeT(%d)", e)
	}
}

type ClusterTypeT uint

const (
	CLUSTER_TYPE_E ClusterTypeT = 0
	MAX_CPU_TYPES ClusterTypeT = 0
)

func (e ClusterTypeT) String() string {
	switch e {
	case CLUSTER_TYPE_E:
		return "CLUSTER_TYPE_E"
	default:
		return fmt.Sprintf("ClusterTypeT(%d)", e)
	}
}

type ClutType uint

const (
	RGBDirect ClutType = 0
)

func (e ClutType) String() string {
	switch e {
	case RGBDirect:
		return "RGBDirect"
	default:
		return fmt.Sprintf("ClutType(%d)", e)
	}
}

type CoalitionFlags uint

const (
	KCoalitionPrivileged CoalitionFlags = 0
)

func (e CoalitionFlags) String() string {
	switch e {
	case KCoalitionPrivileged:
		return "KCoalitionPrivileged"
	default:
		return fmt.Sprintf("CoalitionFlags(%d)", e)
	}
}

type CpuidRegisterT uint

const (
	Eax CpuidRegisterT = 0
)

func (e CpuidRegisterT) String() string {
	switch e {
	case Eax:
		return "Eax"
	default:
		return fmt.Sprintf("CpuidRegisterT(%d)", e)
	}
}

type CryptexAuthTypeT uint

const (
	CRYPTEX_AUTH_MAX CryptexAuthTypeT = 0
)

func (e CryptexAuthTypeT) String() string {
	switch e {
	case CRYPTEX_AUTH_MAX:
		return "CRYPTEX_AUTH_MAX"
	default:
		return fmt.Sprintf("CryptexAuthTypeT(%d)", e)
	}
}

type CsLaunchTypeT uint

const (
	CS_LAUNCH_TYPE_SYSDIAGNOSE CsLaunchTypeT = 0
)

func (e CsLaunchTypeT) String() string {
	switch e {
	case CS_LAUNCH_TYPE_SYSDIAGNOSE:
		return "CS_LAUNCH_TYPE_SYSDIAGNOSE"
	default:
		return fmt.Sprintf("CsLaunchTypeT(%d)", e)
	}
}

type CsLinkageApplicationInvalid uint

const (
	CS_LINKAGE_APPLICATION_INVALID CsLinkageApplicationInvalid = 0
)

func (e CsLinkageApplicationInvalid) String() string {
	switch e {
	case CS_LINKAGE_APPLICATION_INVALID:
		return "CS_LINKAGE_APPLICATION_INVALID"
	default:
		return fmt.Sprintf("CsLinkageApplicationInvalid(%d)", e)
	}
}

type CsLinkageApplicationRosettaAot uint

const (
	CS_LINKAGE_APPLICATION_XOJIT_PREVIEWS CsLinkageApplicationRosettaAot = 0
)

func (e CsLinkageApplicationRosettaAot) String() string {
	switch e {
	case CS_LINKAGE_APPLICATION_XOJIT_PREVIEWS:
		return "CS_LINKAGE_APPLICATION_XOJIT_PREVIEWS"
	default:
		return fmt.Sprintf("CsLinkageApplicationRosettaAot(%d)", e)
	}
}

type CscGetMode uint

const (
	CscGetDetailedTiming CscGetMode = 0
)

func (e CscGetMode) String() string {
	switch e {
	case CscGetDetailedTiming:
		return "CscGetDetailedTiming"
	default:
		return fmt.Sprintf("CscGetMode(%d)", e)
	}
}

type CscReset uint

const (
	CscGrayPage CscReset = 0
)

func (e CscReset) String() string {
	switch e {
	case CscGrayPage:
		return "CscGrayPage"
	default:
		return fmt.Sprintf("CscReset(%d)", e)
	}
}

type CsmagicRequirement uint

const (
	CS_CDHASH_LEN CsmagicRequirement = 0
)

func (e CsmagicRequirement) String() string {
	switch e {
	case CS_CDHASH_LEN:
		return "CS_CDHASH_LEN"
	default:
		return fmt.Sprintf("CsmagicRequirement(%d)", e)
	}
}

type DevelopStage uint

const (
	AlphaStage DevelopStage = 0
)

func (e DevelopStage) String() string {
	switch e {
	case AlphaStage:
		return "AlphaStage"
	default:
		return fmt.Sprintf("DevelopStage(%d)", e)
	}
}

type DirCloneAuthorizerOp uint

const (
	OP_AUTHORIZE DirCloneAuthorizerOp = 0
)

func (e DirCloneAuthorizerOp) String() string {
	switch e {
	case OP_AUTHORIZE:
		return "OP_AUTHORIZE"
	default:
		return fmt.Sprintf("DirCloneAuthorizerOp(%d)", e)
	}
}

const Dot3ChipSetAMD7990 uint = 0

type Dot3ChipSetDigitalDC21040 uint

const (
	Dot3ChipSetDigitalDC21140A Dot3ChipSetDigitalDC21040 = 0
)

func (e Dot3ChipSetDigitalDC21040) String() string {
	switch e {
	case Dot3ChipSetDigitalDC21140A:
		return "Dot3ChipSetDigitalDC21140A"
	default:
		return fmt.Sprintf("Dot3ChipSetDigitalDC21040(%d)", e)
	}
}

const Dot3ChipSetFujitsu86950 uint = 0

type Dot3ChipSetIntel82586 uint

const (
	Dot3ChipSetIntel82557 Dot3ChipSetIntel82586 = 0
)

func (e Dot3ChipSetIntel82586) String() string {
	switch e {
	case Dot3ChipSetIntel82557:
		return "Dot3ChipSetIntel82557"
	default:
		return fmt.Sprintf("Dot3ChipSetIntel82586(%d)", e)
	}
}

const Dot3ChipSetNational8390 uint = 0

const Dot3ChipSetWesternDigital83C690 uint = 0

type Dot3Vendors uint

const (
	Dot3VendorAMD Dot3Vendors = 0
)

func (e Dot3Vendors) String() string {
	switch e {
	case Dot3VendorAMD:
		return "Dot3VendorAMD"
	default:
		return fmt.Sprintf("Dot3Vendors(%d)", e)
	}
}

type DurationMicrosecond uint

const (
	DurationMinute DurationMicrosecond = 0
)

func (e DurationMicrosecond) String() string {
	switch e {
	case DurationMinute:
		return "DurationMinute"
	default:
		return fmt.Sprintf("DurationMicrosecond(%d)", e)
	}
}

type DyldChainedImport uint

const (
	DYLD_CHAINED_IMPORT_ADDEND64 DyldChainedImport = 0
)

func (e DyldChainedImport) String() string {
	switch e {
	case DYLD_CHAINED_IMPORT_ADDEND64:
		return "DYLD_CHAINED_IMPORT_ADDEND64"
	default:
		return fmt.Sprintf("DyldChainedImport(%d)", e)
	}
}

type DyldChainedPtrArm64e uint

const (
	DYLD_CHAINED_PTR_32 DyldChainedPtrArm64e = 0
)

func (e DyldChainedPtrArm64e) String() string {
	switch e {
	case DYLD_CHAINED_PTR_32:
		return "DYLD_CHAINED_PTR_32"
	default:
		return fmt.Sprintf("DyldChainedPtrArm64e(%d)", e)
	}
}

type DyldChainedPtrStartNone uint

const (
	DYLD_CHAINED_PTR_START_MULTI DyldChainedPtrStartNone = 0
)

func (e DyldChainedPtrStartNone) String() string {
	switch e {
	case DYLD_CHAINED_PTR_START_MULTI:
		return "DYLD_CHAINED_PTR_START_MULTI"
	default:
		return fmt.Sprintf("DyldChainedPtrStartNone(%d)", e)
	}
}

type DyldSharedCacheFlags uint

const (
	KSharedCacheAOT DyldSharedCacheFlags = 0
)

func (e DyldSharedCacheFlags) String() string {
	switch e {
	case KSharedCacheAOT:
		return "KSharedCacheAOT"
	default:
		return fmt.Sprintf("DyldSharedCacheFlags(%d)", e)
	}
}

type EIOAccelSurfaceLockBits uint

const (
	KIOAccelSurfaceLockInAccel EIOAccelSurfaceLockBits = 0
)

func (e EIOAccelSurfaceLockBits) String() string {
	switch e {
	case KIOAccelSurfaceLockInAccel:
		return "KIOAccelSurfaceLockInAccel"
	default:
		return fmt.Sprintf("EIOAccelSurfaceLockBits(%d)", e)
	}
}

type EIOAccelSurfaceMemoryTypes uint

const (
	KIOAccelNumSurfaceMemoryTypes EIOAccelSurfaceMemoryTypes = 0
)

func (e EIOAccelSurfaceMemoryTypes) String() string {
	switch e {
	case KIOAccelNumSurfaceMemoryTypes:
		return "KIOAccelNumSurfaceMemoryTypes"
	default:
		return fmt.Sprintf("EIOAccelSurfaceMemoryTypes(%d)", e)
	}
}

type EIOAccelSurfaceModeBits uint

const (
	KIOAccelSurfaceModeBeamSync EIOAccelSurfaceModeBits = 0
)

func (e EIOAccelSurfaceModeBits) String() string {
	switch e {
	case KIOAccelSurfaceModeBeamSync:
		return "KIOAccelSurfaceModeBeamSync"
	default:
		return fmt.Sprintf("EIOAccelSurfaceModeBits(%d)", e)
	}
}

type EIOAccelSurfaceScaleBits uint

const (
	KIOAccelSurfaceBeamSyncSwaps EIOAccelSurfaceScaleBits = 0
)

func (e EIOAccelSurfaceScaleBits) String() string {
	switch e {
	case KIOAccelSurfaceBeamSyncSwaps:
		return "KIOAccelSurfaceBeamSyncSwaps"
	default:
		return fmt.Sprintf("EIOAccelSurfaceScaleBits(%d)", e)
	}
}

type EIOAccelSurfaceShapeBits uint

const (
	KIOAccelSurfaceShapeAssemblyBit EIOAccelSurfaceShapeBits = 0
)

func (e EIOAccelSurfaceShapeBits) String() string {
	switch e {
	case KIOAccelSurfaceShapeAssemblyBit:
		return "KIOAccelSurfaceShapeAssemblyBit"
	default:
		return fmt.Sprintf("EIOAccelSurfaceShapeBits(%d)", e)
	}
}

type EIOAccelSurfaceStateBits uint

const (
	KIOAccelSurfaceStateIdleBit EIOAccelSurfaceStateBits = 0
)

func (e EIOAccelSurfaceStateBits) String() string {
	switch e {
	case KIOAccelSurfaceStateIdleBit:
		return "KIOAccelSurfaceStateIdleBit"
	default:
		return fmt.Sprintf("EIOAccelSurfaceStateBits(%d)", e)
	}
}

const ENoteExitReparentedDeprecated uint = 0

const ENoteReapDeprecated uint = 0

type EXBrightMessageType uint

const (
	KLoadALSSCalibration EXBrightMessageType = 0
)

func (e EXBrightMessageType) String() string {
	switch e {
	case KLoadALSSCalibration:
		return "KLoadALSSCalibration"
	default:
		return fmt.Sprintf("EXBrightMessageType(%d)", e)
	}
}

type EXDisplayPipeIndicator uint

const (
	INDICATOR_CAM EXDisplayPipeIndicator = 0
)

func (e EXDisplayPipeIndicator) String() string {
	switch e {
	case INDICATOR_CAM:
		return "INDICATOR_CAM"
	default:
		return fmt.Sprintf("EXDisplayPipeIndicator(%d)", e)
	}
}

type Ecc uint

const (
	ECC_DB_CORRUPTED Ecc = 0
	ECC_DB_ONLY Ecc = 0
	ECC_IS_CORRECTABLE Ecc = 0
	ECC_IS_TEST_ERROR Ecc = 0
	ECC_NONE Ecc = 0
	ECC_REMOVE_ADDR Ecc = 0
)

func (e Ecc) String() string {
	switch e {
	case ECC_DB_CORRUPTED:
		return "ECC_DB_CORRUPTED"
	default:
		return fmt.Sprintf("Ecc(%d)", e)
	}
}

type Efi uint

const (
	EfiACPIMemoryNVS Efi = 0
	EfiACPIReclaimMemory Efi = 0
	EfiBootServicesCode Efi = 0
	EfiBootServicesData Efi = 0
	EfiConventionalMemory Efi = 0
	EfiLoaderCode Efi = 0
	EfiLoaderData Efi = 0
	EfiMaxMemoryType Efi = 0
	EfiMemoryMappedIO Efi = 0
	EfiMemoryMappedIOPortSpace Efi = 0
	EfiPalCode Efi = 0
	EfiReservedMemoryType Efi = 0
	EfiRuntimeServicesCode Efi = 0
	EfiRuntimeServicesData Efi = 0
	EfiUnusableMemory Efi = 0
)

func (e Efi) String() string {
	switch e {
	case EfiACPIMemoryNVS:
		return "EfiACPIMemoryNVS"
	default:
		return fmt.Sprintf("Efi(%d)", e)
	}
}

type EfiResetType uint

const (
	EfiResetCold EfiResetType = 0
)

func (e EfiResetType) String() string {
	switch e {
	case EfiResetCold:
		return "EfiResetCold"
	default:
		return fmt.Sprintf("EfiResetType(%d)", e)
	}
}

type EmbeddedUndefined uint

const (
	EMBEDDED_PHONOGRAPH EmbeddedUndefined = 0
)

func (e EmbeddedUndefined) String() string {
	switch e {
	case EMBEDDED_PHONOGRAPH:
		return "EMBEDDED_PHONOGRAPH"
	default:
		return fmt.Sprintf("EmbeddedUndefined(%d)", e)
	}
}

type EphPanicFlagsT uint

const (
	EMBEDDED_PANIC_HEADER_FLAG_BUTTON_RESET_PANIC EphPanicFlagsT = 0
)

func (e EphPanicFlagsT) String() string {
	switch e {
	case EMBEDDED_PANIC_HEADER_FLAG_BUTTON_RESET_PANIC:
		return "EMBEDDED_PANIC_HEADER_FLAG_BUTTON_RESET_PANIC"
	default:
		return fmt.Sprintf("EphPanicFlagsT(%d)", e)
	}
}

type ExCbActionT uint

const (
	EXCB_ACTION_NONE ExCbActionT = 0
)

func (e ExCbActionT) String() string {
	switch e {
	case EXCB_ACTION_NONE:
		return "EXCB_ACTION_NONE"
	default:
		return fmt.Sprintf("ExCbActionT(%d)", e)
	}
}

type ExCbClassT uint

const (
	EXCB_CLASS_ILLEGAL_INSTR_SET ExCbClassT = 0
)

func (e ExCbClassT) String() string {
	switch e {
	case EXCB_CLASS_ILLEGAL_INSTR_SET:
		return "EXCB_CLASS_ILLEGAL_INSTR_SET"
	default:
		return fmt.Sprintf("ExCbClassT(%d)", e)
	}
}

type ExclaveAddressspaceFlags uint

const (
	KExclaveAddressSpaceHaveSlide ExclaveAddressspaceFlags = 0
)

func (e ExclaveAddressspaceFlags) String() string {
	switch e {
	case KExclaveAddressSpaceHaveSlide:
		return "KExclaveAddressSpaceHaveSlide"
	default:
		return fmt.Sprintf("ExclaveAddressspaceFlags(%d)", e)
	}
}

type ExclaveScresultFlags uint

const (
	KExclaveScresultHaveIPCStack ExclaveScresultFlags = 0
)

func (e ExclaveScresultFlags) String() string {
	switch e {
	case KExclaveScresultHaveIPCStack:
		return "KExclaveScresultHaveIPCStack"
	default:
		return fmt.Sprintf("ExclaveScresultFlags(%d)", e)
	}
}

type ExtPaniclogCreateOptionsT uint

const (
	EXT_PANICLOG_OPTIONS_ADD_SEPARATE_KEY ExtPaniclogCreateOptionsT = 0
)

func (e ExtPaniclogCreateOptionsT) String() string {
	switch e {
	case EXT_PANICLOG_OPTIONS_ADD_SEPARATE_KEY:
		return "EXT_PANICLOG_OPTIONS_ADD_SEPARATE_KEY"
	default:
		return fmt.Sprintf("ExtPaniclogCreateOptionsT(%d)", e)
	}
}

type ExtPaniclogFlagsT uint

const (
	EXT_PANICLOG_FLAGS_ADD_SEPARATE_KEY ExtPaniclogFlagsT = 0
)

func (e ExtPaniclogFlagsT) String() string {
	switch e {
	case EXT_PANICLOG_FLAGS_ADD_SEPARATE_KEY:
		return "EXT_PANICLOG_FLAGS_ADD_SEPARATE_KEY"
	default:
		return fmt.Sprintf("ExtPaniclogFlagsT(%d)", e)
	}
}

type ExternalUndefined uint

const (
	EXTERNAL_1394_DA_STREAM ExternalUndefined = 0
)

func (e ExternalUndefined) String() string {
	switch e {
	case EXTERNAL_1394_DA_STREAM:
		return "EXTERNAL_1394_DA_STREAM"
	default:
		return fmt.Sprintf("ExternalUndefined(%d)", e)
	}
}

type FftForward uint

const (
	FFT_FORWARD FftForward = 0
)

func (e FftForward) String() string {
	switch e {
	case FFT_FORWARD:
		return "FFT_FORWARD"
	default:
		return fmt.Sprintf("FftForward(%d)", e)
	}
}

type FftRadix2 uint

const (
	FFT_RADIX2 FftRadix2 = 0
)

func (e FftRadix2) String() string {
	switch e {
	case FFT_RADIX2:
		return "FFT_RADIX2"
	default:
		return fmt.Sprintf("FftRadix2(%d)", e)
	}
}

type GenericSnapshotFlags uint

const (
	KKernel64_p GenericSnapshotFlags = 0
	KUser64_p GenericSnapshotFlags = 0
)

func (e GenericSnapshotFlags) String() string {
	switch e {
	case KKernel64_p:
		return "KKernel64_p"
	default:
		return fmt.Sprintf("GenericSnapshotFlags(%d)", e)
	}
}

type GraftdmgTypeT uint

const (
	GRAFTDMG_CRYPTEX_BOOT GraftdmgTypeT = 0
)

func (e GraftdmgTypeT) String() string {
	switch e {
	case GRAFTDMG_CRYPTEX_BOOT:
		return "GRAFTDMG_CRYPTEX_BOOT"
	default:
		return fmt.Sprintf("GraftdmgTypeT(%d)", e)
	}
}

type GssdMechtype uint

const (
	GSSD_NO_MECH GssdMechtype = 0
)

func (e GssdMechtype) String() string {
	switch e {
	case GSSD_NO_MECH:
		return "GSSD_NO_MECH"
	default:
		return fmt.Sprintf("GssdMechtype(%d)", e)
	}
}

type GssdNametype uint

const (
	GSSD_KRB5_REFERRAL GssdNametype = 0
)

func (e GssdNametype) String() string {
	switch e {
	case GSSD_KRB5_REFERRAL:
		return "GSSD_KRB5_REFERRAL"
	default:
		return fmt.Sprintf("GssdNametype(%d)", e)
	}
}

type HvTrapTypeT uint

const (
	HV_THREAD_TRAP HvTrapTypeT = 0
)

func (e HvTrapTypeT) String() string {
	switch e {
	case HV_THREAD_TRAP:
		return "HV_THREAD_TRAP"
	default:
		return fmt.Sprintf("HvTrapTypeT(%d)", e)
	}
}

type HvVolatileStateT uint

const (
	HV_DEBUG_STATE HvVolatileStateT = 0
)

func (e HvVolatileStateT) String() string {
	switch e {
	case HV_DEBUG_STATE:
		return "HV_DEBUG_STATE"
	default:
		return fmt.Sprintf("HvVolatileStateT(%d)", e)
	}
}

type HvgHcallCodeT uint

const (
	HVG_HCALL_GET_BOOTSESSIONUUID HvgHcallCodeT = 0
)

func (e HvgHcallCodeT) String() string {
	switch e {
	case HVG_HCALL_GET_BOOTSESSIONUUID:
		return "HVG_HCALL_GET_BOOTSESSIONUUID"
	default:
		return fmt.Sprintf("HvgHcallCodeT(%d)", e)
	}
}

type HvgHcallDumpOptionT uint

const (
	HVG_HCALL_DUMP_OPTION_REGULAR HvgHcallDumpOptionT = 0
)

func (e HvgHcallDumpOptionT) String() string {
	switch e {
	case HVG_HCALL_DUMP_OPTION_REGULAR:
		return "HVG_HCALL_DUMP_OPTION_REGULAR"
	default:
		return fmt.Sprintf("HvgHcallDumpOptionT(%d)", e)
	}
}

type HvgHcallReturnT uint

const (
	HVG_HCALL_ACCESS_DENIED HvgHcallReturnT = 0
)

func (e HvgHcallReturnT) String() string {
	switch e {
	case HVG_HCALL_ACCESS_DENIED:
		return "HVG_HCALL_ACCESS_DENIED"
	default:
		return fmt.Sprintf("HvgHcallReturnT(%d)", e)
	}
}

type IOCSRKeyType uint

const (
	KCSRDirectoryKeyType IOCSRKeyType = 0
	KCSRImmediateKeyType IOCSRKeyType = 0
	KCSROffsetKeyType IOCSRKeyType = 0
	KInvalidCSRROMEntryType IOCSRKeyType = 0
)

func (e IOCSRKeyType) String() string {
	switch e {
	case KCSRDirectoryKeyType:
		return "KCSRDirectoryKeyType"
	default:
		return fmt.Sprintf("IOCSRKeyType(%d)", e)
	}
}

type IOConfigKeyType uint

const (
	KConfigDirectoryKeyType IOConfigKeyType = 0
	KConfigImmediateKeyType IOConfigKeyType = 0
	KConfigLeafKeyType IOConfigKeyType = 0
	KInvalidConfigROMEntryType IOConfigKeyType = 0
)

func (e IOConfigKeyType) String() string {
	switch e {
	case KConfigDirectoryKeyType:
		return "KConfigDirectoryKeyType"
	default:
		return fmt.Sprintf("IOConfigKeyType(%d)", e)
	}
}

type IODebuggerLockState uint

const (
	// KIODebuggerLockTaken: # Discussion
	KIODebuggerLockTaken IODebuggerLockState = 0
)

func (e IODebuggerLockState) String() string {
	switch e {
	case KIODebuggerLockTaken:
		return "KIODebuggerLockTaken"
	default:
		return fmt.Sprintf("IODebuggerLockState(%d)", e)
	}
}

type IOEthernetControllerAVBState uint

const (
	KIOEthernetControllerAVBStateAVBEnabled IOEthernetControllerAVBState = 0
)

func (e IOEthernetControllerAVBState) String() string {
	switch e {
	case KIOEthernetControllerAVBStateAVBEnabled:
		return "KIOEthernetControllerAVBStateAVBEnabled"
	default:
		return fmt.Sprintf("IOEthernetControllerAVBState(%d)", e)
	}
}

type IOEthernetControllerAVBTimeSyncSupport uint

const (
	KIOEthernetControllerAVBTimeSyncSupportHardware IOEthernetControllerAVBTimeSyncSupport = 0
)

func (e IOEthernetControllerAVBTimeSyncSupport) String() string {
	switch e {
	case KIOEthernetControllerAVBTimeSyncSupportHardware:
		return "KIOEthernetControllerAVBTimeSyncSupportHardware"
	default:
		return fmt.Sprintf("IOEthernetControllerAVBTimeSyncSupport(%d)", e)
	}
}

type IOFWAVCPlug uint

const (
	IOFWAVCPlugAsynchInputType IOFWAVCPlug = 0
	IOFWAVCPlugAsynchOutputType IOFWAVCPlug = 0
	IOFWAVCPlugExternalInputType IOFWAVCPlug = 0
	IOFWAVCPlugExternalOutputType IOFWAVCPlug = 0
	IOFWAVCPlugIsochInputType IOFWAVCPlug = 0
	IOFWAVCPlugIsochOutputType IOFWAVCPlug = 0
	IOFWAVCPlugSubunitDestType IOFWAVCPlug = 0
	IOFWAVCPlugSubunitSourceType IOFWAVCPlug = 0
)

func (e IOFWAVCPlug) String() string {
	switch e {
	case IOFWAVCPlugAsynchInputType:
		return "IOFWAVCPlugAsynchInputType"
	default:
		return fmt.Sprintf("IOFWAVCPlug(%d)", e)
	}
}

type IOHIDKeyboardEventOptions uint

const (
	KIOHIDKeyboardEventOptionsNoKeyRepeat IOHIDKeyboardEventOptions = 0
)

func (e IOHIDKeyboardEventOptions) String() string {
	switch e {
	case KIOHIDKeyboardEventOptionsNoKeyRepeat:
		return "KIOHIDKeyboardEventOptionsNoKeyRepeat"
	default:
		return fmt.Sprintf("IOHIDKeyboardEventOptions(%d)", e)
	}
}

type IOHIDKind uint

const (
	KHIKeyboardDevice IOHIDKind = 0
)

func (e IOHIDKind) String() string {
	switch e {
	case KHIKeyboardDevice:
		return "KHIKeyboardDevice"
	default:
		return fmt.Sprintf("IOHIDKind(%d)", e)
	}
}

type IOHIDPointerEventOptions uint

const (
	KIOHIDPointerEventOptionsNoAcceleration IOHIDPointerEventOptions = 0
)

func (e IOHIDPointerEventOptions) String() string {
	switch e {
	case KIOHIDPointerEventOptionsNoAcceleration:
		return "KIOHIDPointerEventOptionsNoAcceleration"
	default:
		return fmt.Sprintf("IOHIDPointerEventOptions(%d)", e)
	}
}

type IOHIDScrollEventOptions uint

const (
	KIOHIDScrollEventOptionsNoAcceleration IOHIDScrollEventOptions = 0
)

func (e IOHIDScrollEventOptions) String() string {
	switch e {
	case KIOHIDScrollEventOptionsNoAcceleration:
		return "KIOHIDScrollEventOptionsNoAcceleration"
	default:
		return fmt.Sprintf("IOHIDScrollEventOptions(%d)", e)
	}
}

type IOLockState uint

const (
	KIOLockStateLocked IOLockState = 0
)

func (e IOLockState) String() string {
	switch e {
	case KIOLockStateLocked:
		return "KIOLockStateLocked"
	default:
		return fmt.Sprintf("IOLockState(%d)", e)
	}
}

type IOPCIResetType uint

const (
	KIOPCIResetHot IOPCIResetType = 0
)

func (e IOPCIResetType) String() string {
	switch e {
	case KIOPCIResetHot:
		return "KIOPCIResetHot"
	default:
		return fmt.Sprintf("IOPCIResetType(%d)", e)
	}
}

type IOPMNextHigherState uint

const (
	IOPMHighestState IOPMNextHigherState = 0
)

func (e IOPMNextHigherState) String() string {
	switch e {
	case IOPMHighestState:
		return "IOPMHighestState"
	default:
		return fmt.Sprintf("IOPMNextHigherState(%d)", e)
	}
}

type IOUSBHostCICapabilitiesMessage uint

const (
	IOUSBHostCICapabilitiesMessageControlPortCount IOUSBHostCICapabilitiesMessage = 0
	IOUSBHostCICapabilitiesMessageData0ConnectionLatency IOUSBHostCICapabilitiesMessage = 0
)

func (e IOUSBHostCICapabilitiesMessage) String() string {
	switch e {
	case IOUSBHostCICapabilitiesMessageControlPortCount:
		return "IOUSBHostCICapabilitiesMessageControlPortCount"
	default:
		return fmt.Sprintf("IOUSBHostCICapabilitiesMessage(%d)", e)
	}
}

type IOUSBHostCICommandMessageControl uint

const (
	IOUSBHostCICommandMessageControlStatus IOUSBHostCICommandMessageControl = 0
	IOUSBHostCICommandMessageControlStatusPhase IOUSBHostCICommandMessageControl = 0
)

func (e IOUSBHostCICommandMessageControl) String() string {
	switch e {
	case IOUSBHostCICommandMessageControlStatus:
		return "IOUSBHostCICommandMessageControlStatus"
	default:
		return fmt.Sprintf("IOUSBHostCICommandMessageControl(%d)", e)
	}
}

type IOUSBHostCIControllerState uint

const (
	IOUSBHostCIControllerStateActive IOUSBHostCIControllerState = 0
	IOUSBHostCIControllerStateOff IOUSBHostCIControllerState = 0
)

func (e IOUSBHostCIControllerState) String() string {
	switch e {
	case IOUSBHostCIControllerStateActive:
		return "IOUSBHostCIControllerStateActive"
	default:
		return fmt.Sprintf("IOUSBHostCIControllerState(%d)", e)
	}
}

type IOUSBHostCIDeviceCreateCommandData0RootPortConstants uint

const (
	IOUSBHostCIDeviceCreateCommandData0RootPort IOUSBHostCIDeviceCreateCommandData0RootPortConstants = 0
)

func (e IOUSBHostCIDeviceCreateCommandData0RootPortConstants) String() string {
	switch e {
	case IOUSBHostCIDeviceCreateCommandData0RootPort:
		return "IOUSBHostCIDeviceCreateCommandData0RootPort"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceCreateCommandData0RootPortConstants(%d)", e)
	}
}

type IOUSBHostCIDeviceCreateCommandData1DeviceAddressConstants uint

const (
	IOUSBHostCIDeviceCreateCommandData1DeviceAddress IOUSBHostCIDeviceCreateCommandData1DeviceAddressConstants = 0
)

func (e IOUSBHostCIDeviceCreateCommandData1DeviceAddressConstants) String() string {
	switch e {
	case IOUSBHostCIDeviceCreateCommandData1DeviceAddress:
		return "IOUSBHostCIDeviceCreateCommandData1DeviceAddress"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceCreateCommandData1DeviceAddressConstants(%d)", e)
	}
}

type IOUSBHostCIDeviceSpeed uint

const (
	IOUSBHostCIDeviceSpeedFull IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedHigh IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedLow IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedNone IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedSuper IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedSuperPlus IOUSBHostCIDeviceSpeed = 0
)

func (e IOUSBHostCIDeviceSpeed) String() string {
	switch e {
	case IOUSBHostCIDeviceSpeedFull:
		return "IOUSBHostCIDeviceSpeedFull"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceSpeed(%d)", e)
	}
}

type IOUSBHostCIDeviceState uint

const (
	IOUSBHostCIDeviceStateActive IOUSBHostCIDeviceState = 0
	IOUSBHostCIDeviceStateDestroyed IOUSBHostCIDeviceState = 0
	IOUSBHostCIDeviceStatePaused IOUSBHostCIDeviceState = 0
)

func (e IOUSBHostCIDeviceState) String() string {
	switch e {
	case IOUSBHostCIDeviceStateActive:
		return "IOUSBHostCIDeviceStateActive"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceState(%d)", e)
	}
}

type IOUSBHostCIDeviceUpdateCommandData1DescriptorAddressConstants uint

const (
	IOUSBHostCIDeviceUpdateCommandData1DescriptorAddress IOUSBHostCIDeviceUpdateCommandData1DescriptorAddressConstants = 0
)

func (e IOUSBHostCIDeviceUpdateCommandData1DescriptorAddressConstants) String() string {
	switch e {
	case IOUSBHostCIDeviceUpdateCommandData1DescriptorAddress:
		return "IOUSBHostCIDeviceUpdateCommandData1DescriptorAddress"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceUpdateCommandData1DescriptorAddressConstants(%d)", e)
	}
}

type IOUSBHostCIDoorbellDevice uint

const (
	IOUSBHostCIDoorbellDeviceAddress IOUSBHostCIDoorbellDevice = 0
	IOUSBHostCIDoorbellDeviceAddressPhase IOUSBHostCIDoorbellDevice = 0
)

func (e IOUSBHostCIDoorbellDevice) String() string {
	switch e {
	case IOUSBHostCIDoorbellDeviceAddress:
		return "IOUSBHostCIDoorbellDeviceAddress"
	default:
		return fmt.Sprintf("IOUSBHostCIDoorbellDevice(%d)", e)
	}
}

type IOUSBHostCIEndpointCreateCommandData1DescriptorConstants uint

const (
	IOUSBHostCIEndpointCreateCommandData1Descriptor IOUSBHostCIEndpointCreateCommandData1DescriptorConstants = 0
)

func (e IOUSBHostCIEndpointCreateCommandData1DescriptorConstants) String() string {
	switch e {
	case IOUSBHostCIEndpointCreateCommandData1Descriptor:
		return "IOUSBHostCIEndpointCreateCommandData1Descriptor"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointCreateCommandData1DescriptorConstants(%d)", e)
	}
}

type IOUSBHostCIEndpointResetCommandData1ClearStateConstants uint

const (
	IOUSBHostCIEndpointResetCommandData1ClearState IOUSBHostCIEndpointResetCommandData1ClearStateConstants = 0
)

func (e IOUSBHostCIEndpointResetCommandData1ClearStateConstants) String() string {
	switch e {
	case IOUSBHostCIEndpointResetCommandData1ClearState:
		return "IOUSBHostCIEndpointResetCommandData1ClearState"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointResetCommandData1ClearStateConstants(%d)", e)
	}
}

type IOUSBHostCIEndpointSetNextTransferCommandData1AddressConstants uint

const (
	IOUSBHostCIEndpointSetNextTransferCommandData1Address IOUSBHostCIEndpointSetNextTransferCommandData1AddressConstants = 0
)

func (e IOUSBHostCIEndpointSetNextTransferCommandData1AddressConstants) String() string {
	switch e {
	case IOUSBHostCIEndpointSetNextTransferCommandData1Address:
		return "IOUSBHostCIEndpointSetNextTransferCommandData1Address"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointSetNextTransferCommandData1AddressConstants(%d)", e)
	}
}

type IOUSBHostCIEndpointState uint

const (
	IOUSBHostCIEndpointStateActive IOUSBHostCIEndpointState = 0
	IOUSBHostCIEndpointStateDestroyed IOUSBHostCIEndpointState = 0
)

func (e IOUSBHostCIEndpointState) String() string {
	switch e {
	case IOUSBHostCIEndpointStateActive:
		return "IOUSBHostCIEndpointStateActive"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointState(%d)", e)
	}
}

type IOUSBHostCIEndpointUpdateCommandData1 uint

const (
	IOUSBHostCIEndpointUpdateCommandData1Descriptor IOUSBHostCIEndpointUpdateCommandData1 = 0
	IOUSBHostCIEndpointUpdateCommandData1DescriptorPhase IOUSBHostCIEndpointUpdateCommandData1 = 0
)

func (e IOUSBHostCIEndpointUpdateCommandData1) String() string {
	switch e {
	case IOUSBHostCIEndpointUpdateCommandData1Descriptor:
		return "IOUSBHostCIEndpointUpdateCommandData1Descriptor"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointUpdateCommandData1(%d)", e)
	}
}

type IOUSBHostCIExceptionType uint

const (
	IOUSBHostCIExceptionTypeCommandFailure IOUSBHostCIExceptionType = 0
	IOUSBHostCIExceptionTypeFrameUpdateError IOUSBHostCIExceptionType = 0
)

func (e IOUSBHostCIExceptionType) String() string {
	switch e {
	case IOUSBHostCIExceptionTypeCommandFailure:
		return "IOUSBHostCIExceptionTypeCommandFailure"
	default:
		return fmt.Sprintf("IOUSBHostCIExceptionType(%d)", e)
	}
}

type IOUSBHostCIIsochronousTransferControl uint

const (
	IOUSBHostCIIsochronousTransferControlASAP IOUSBHostCIIsochronousTransferControl = 0
	IOUSBHostCIIsochronousTransferControlFrameNumber IOUSBHostCIIsochronousTransferControl = 0
	IOUSBHostCIIsochronousTransferControlFrameNumberPhase IOUSBHostCIIsochronousTransferControl = 0
)

func (e IOUSBHostCIIsochronousTransferControl) String() string {
	switch e {
	case IOUSBHostCIIsochronousTransferControlASAP:
		return "IOUSBHostCIIsochronousTransferControlASAP"
	default:
		return fmt.Sprintf("IOUSBHostCIIsochronousTransferControl(%d)", e)
	}
}

type IOUSBHostCIIsochronousTransferData0LengthConstants uint

const (
	IOUSBHostCIIsochronousTransferData0Length IOUSBHostCIIsochronousTransferData0LengthConstants = 0
)

func (e IOUSBHostCIIsochronousTransferData0LengthConstants) String() string {
	switch e {
	case IOUSBHostCIIsochronousTransferData0Length:
		return "IOUSBHostCIIsochronousTransferData0Length"
	default:
		return fmt.Sprintf("IOUSBHostCIIsochronousTransferData0LengthConstants(%d)", e)
	}
}

type IOUSBHostCIIsochronousTransferData1BufferConstants uint

const (
	IOUSBHostCIIsochronousTransferData1Buffer IOUSBHostCIIsochronousTransferData1BufferConstants = 0
)

func (e IOUSBHostCIIsochronousTransferData1BufferConstants) String() string {
	switch e {
	case IOUSBHostCIIsochronousTransferData1Buffer:
		return "IOUSBHostCIIsochronousTransferData1Buffer"
	default:
		return fmt.Sprintf("IOUSBHostCIIsochronousTransferData1BufferConstants(%d)", e)
	}
}

type IOUSBHostCILinkData1TransferStructureAddressConstants uint

const (
	IOUSBHostCILinkData1TransferStructureAddress IOUSBHostCILinkData1TransferStructureAddressConstants = 0
)

func (e IOUSBHostCILinkData1TransferStructureAddressConstants) String() string {
	switch e {
	case IOUSBHostCILinkData1TransferStructureAddress:
		return "IOUSBHostCILinkData1TransferStructureAddress"
	default:
		return fmt.Sprintf("IOUSBHostCILinkData1TransferStructureAddressConstants(%d)", e)
	}
}

type IOUSBHostCILinkState uint

const (
	IOUSBHostCILinkStateCompliance IOUSBHostCILinkState = 0
)

func (e IOUSBHostCILinkState) String() string {
	switch e {
	case IOUSBHostCILinkStateCompliance:
		return "IOUSBHostCILinkStateCompliance"
	default:
		return fmt.Sprintf("IOUSBHostCILinkState(%d)", e)
	}
}

type IOUSBHostCIMessageControlType uint

const (
	IOUSBHostCIMessageControlNoResponse IOUSBHostCIMessageControlType = 0
)

func (e IOUSBHostCIMessageControlType) String() string {
	switch e {
	case IOUSBHostCIMessageControlNoResponse:
		return "IOUSBHostCIMessageControlNoResponse"
	default:
		return fmt.Sprintf("IOUSBHostCIMessageControlType(%d)", e)
	}
}

type IOUSBHostCIMessageStatus int

const (
	IOUSBHostCIMessageStatusBadArgument IOUSBHostCIMessageStatus = 0
	IOUSBHostCIMessageStatusEndpointStopped IOUSBHostCIMessageStatus = 0
	IOUSBHostCIMessageStatusError IOUSBHostCIMessageStatus = 0
	IOUSBHostCIMessageStatusMissedServiceError IOUSBHostCIMessageStatus = 0
	IOUSBHostCIMessageStatusOffline IOUSBHostCIMessageStatus = 0
)

func (e IOUSBHostCIMessageStatus) String() string {
	switch e {
	case IOUSBHostCIMessageStatusBadArgument:
		return "IOUSBHostCIMessageStatusBadArgument"
	default:
		return fmt.Sprintf("IOUSBHostCIMessageStatus(%d)", e)
	}
}

type IOUSBHostCIMessageType uint

const (
	IOUSBHostCIMessageTypeControllerFrameNumber IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypeDeviceUpdate IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypeEndpointReset IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypeNormalTransfer IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypePortPowerOn IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypePortSuspend IOUSBHostCIMessageType = 0
)

func (e IOUSBHostCIMessageType) String() string {
	switch e {
	case IOUSBHostCIMessageTypeControllerFrameNumber:
		return "IOUSBHostCIMessageTypeControllerFrameNumber"
	default:
		return fmt.Sprintf("IOUSBHostCIMessageType(%d)", e)
	}
}

type IOUSBHostCINormalTransferData0LengthConstants uint

const (
	IOUSBHostCINormalTransferData0Length IOUSBHostCINormalTransferData0LengthConstants = 0
)

func (e IOUSBHostCINormalTransferData0LengthConstants) String() string {
	switch e {
	case IOUSBHostCINormalTransferData0Length:
		return "IOUSBHostCINormalTransferData0Length"
	default:
		return fmt.Sprintf("IOUSBHostCINormalTransferData0LengthConstants(%d)", e)
	}
}

type IOUSBHostCINormalTransferData1BufferConstants uint

const (
	IOUSBHostCINormalTransferData1Buffer IOUSBHostCINormalTransferData1BufferConstants = 0
)

func (e IOUSBHostCINormalTransferData1BufferConstants) String() string {
	switch e {
	case IOUSBHostCINormalTransferData1Buffer:
		return "IOUSBHostCINormalTransferData1Buffer"
	default:
		return fmt.Sprintf("IOUSBHostCINormalTransferData1BufferConstants(%d)", e)
	}
}

type IOUSBHostCIPortCapabilitiesMessageControl uint

const (
	IOUSBHostCIPortCapabilitiesMessageControlInternalConnector IOUSBHostCIPortCapabilitiesMessageControl = 0
	IOUSBHostCIPortCapabilitiesMessageControlPortNumber IOUSBHostCIPortCapabilitiesMessageControl = 0
)

func (e IOUSBHostCIPortCapabilitiesMessageControl) String() string {
	switch e {
	case IOUSBHostCIPortCapabilitiesMessageControlInternalConnector:
		return "IOUSBHostCIPortCapabilitiesMessageControlInternalConnector"
	default:
		return fmt.Sprintf("IOUSBHostCIPortCapabilitiesMessageControl(%d)", e)
	}
}

type IOUSBHostCIPortEventMessageData0PortNumberConstants uint

const (
	IOUSBHostCIPortEventMessageData0PortNumber IOUSBHostCIPortEventMessageData0PortNumberConstants = 0
)

func (e IOUSBHostCIPortEventMessageData0PortNumberConstants) String() string {
	switch e {
	case IOUSBHostCIPortEventMessageData0PortNumber:
		return "IOUSBHostCIPortEventMessageData0PortNumber"
	default:
		return fmt.Sprintf("IOUSBHostCIPortEventMessageData0PortNumberConstants(%d)", e)
	}
}

type IOUSBHostCIPortState uint

const (
	IOUSBHostCIPortStateActive IOUSBHostCIPortState = 0
	IOUSBHostCIPortStateOff IOUSBHostCIPortState = 0
	IOUSBHostCIPortStatePowered IOUSBHostCIPortState = 0
)

func (e IOUSBHostCIPortState) String() string {
	switch e {
	case IOUSBHostCIPortStateActive:
		return "IOUSBHostCIPortStateActive"
	default:
		return fmt.Sprintf("IOUSBHostCIPortState(%d)", e)
	}
}

type IOUSBHostCIPortStatus int

const (
	IOUSBHostCIPortStatusConnectChange IOUSBHostCIPortStatus = 0
	IOUSBHostCIPortStatusLinkState IOUSBHostCIPortStatus = 0
)

func (e IOUSBHostCIPortStatus) String() string {
	switch e {
	case IOUSBHostCIPortStatusConnectChange:
		return "IOUSBHostCIPortStatusConnectChange"
	default:
		return fmt.Sprintf("IOUSBHostCIPortStatus(%d)", e)
	}
}

type IOUSBHostCIPortStatusCommandData1 int

const (
	IOUSBHostCIPortStatusCommandData1ChangeMask IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1Speed IOUSBHostCIPortStatusCommandData1 = 0
)

func (e IOUSBHostCIPortStatusCommandData1) String() string {
	switch e {
	case IOUSBHostCIPortStatusCommandData1ChangeMask:
		return "IOUSBHostCIPortStatusCommandData1ChangeMask"
	default:
		return fmt.Sprintf("IOUSBHostCIPortStatusCommandData1(%d)", e)
	}
}

type IOUSBHostCISetupTransfer uint

const (
	IOUSBHostCISetupTransferData1bRequest IOUSBHostCISetupTransfer = 0
	IOUSBHostCISetupTransferData1wValuePhase IOUSBHostCISetupTransfer = 0
)

func (e IOUSBHostCISetupTransfer) String() string {
	switch e {
	case IOUSBHostCISetupTransferData1bRequest:
		return "IOUSBHostCISetupTransferData1bRequest"
	default:
		return fmt.Sprintf("IOUSBHostCISetupTransfer(%d)", e)
	}
}

type IOUSBHostCITransferCompletionMessageControl uint

const (
	IOUSBHostCITransferCompletionMessageControlDeviceAddress IOUSBHostCITransferCompletionMessageControl = 0
	IOUSBHostCITransferCompletionMessageControlEndpointAddress IOUSBHostCITransferCompletionMessageControl = 0
)

func (e IOUSBHostCITransferCompletionMessageControl) String() string {
	switch e {
	case IOUSBHostCITransferCompletionMessageControlDeviceAddress:
		return "IOUSBHostCITransferCompletionMessageControlDeviceAddress"
	default:
		return fmt.Sprintf("IOUSBHostCITransferCompletionMessageControl(%d)", e)
	}
}

type IOUSBHostCITransferCompletionMessageData0TransferLengthConstants uint

const (
	IOUSBHostCITransferCompletionMessageData0TransferLength IOUSBHostCITransferCompletionMessageData0TransferLengthConstants = 0
)

func (e IOUSBHostCITransferCompletionMessageData0TransferLengthConstants) String() string {
	switch e {
	case IOUSBHostCITransferCompletionMessageData0TransferLength:
		return "IOUSBHostCITransferCompletionMessageData0TransferLength"
	default:
		return fmt.Sprintf("IOUSBHostCITransferCompletionMessageData0TransferLengthConstants(%d)", e)
	}
}

type IOUSBHostCITransferCompletionMessageData1TransferStructureConstants uint

const (
	IOUSBHostCITransferCompletionMessageData1TransferStructure IOUSBHostCITransferCompletionMessageData1TransferStructureConstants = 0
)

func (e IOUSBHostCITransferCompletionMessageData1TransferStructureConstants) String() string {
	switch e {
	case IOUSBHostCITransferCompletionMessageData1TransferStructure:
		return "IOUSBHostCITransferCompletionMessageData1TransferStructure"
	default:
		return fmt.Sprintf("IOUSBHostCITransferCompletionMessageData1TransferStructureConstants(%d)", e)
	}
}

type IOUSBHostCIUserClientVersion uint

const (
	IOUSBHostCIUserClientVersion100 IOUSBHostCIUserClientVersion = 0
)

func (e IOUSBHostCIUserClientVersion) String() string {
	switch e {
	case IOUSBHostCIUserClientVersion100:
		return "IOUSBHostCIUserClientVersion100"
	default:
		return fmt.Sprintf("IOUSBHostCIUserClientVersion(%d)", e)
	}
}

type IfInterfaceAdvisoryNotificationTypeCellular uint

const (
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_BANDWIDTH_LIMITATION_EVENT IfInterfaceAdvisoryNotificationTypeCellular = 0
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_DEFAULT IfInterfaceAdvisoryNotificationTypeCellular = 0
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_DISCONTINUOUS_RECEPTION_EVENT IfInterfaceAdvisoryNotificationTypeCellular = 0
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_MEASUREMENT_UPDATE IfInterfaceAdvisoryNotificationTypeCellular = 0
)

func (e IfInterfaceAdvisoryNotificationTypeCellular) String() string {
	switch e {
	case IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_BANDWIDTH_LIMITATION_EVENT:
		return "IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_BANDWIDTH_LIMITATION_EVENT"
	default:
		return fmt.Sprintf("IfInterfaceAdvisoryNotificationTypeCellular(%d)", e)
	}
}

type IfNetemModelT uint

const (
	IF_NETEM_MODEL_NULL IfNetemModelT = 0
)

func (e IfNetemModelT) String() string {
	switch e {
	case IF_NETEM_MODEL_NULL:
		return "IF_NETEM_MODEL_NULL"
	default:
		return fmt.Sprintf("IfNetemModelT(%d)", e)
	}
}

type IfnetCsumIp uint

const (
	IFNET_HW_TIMESTAMP IfnetCsumIp = 0
)

func (e IfnetCsumIp) String() string {
	switch e {
	case IFNET_HW_TIMESTAMP:
		return "IFNET_HW_TIMESTAMP"
	default:
		return fmt.Sprintf("IfnetCsumIp(%d)", e)
	}
}

type IfnetFamilyAny uint

const (
	// IFNET_FAMILY_BOND: # Discussion
	IFNET_FAMILY_BOND IfnetFamilyAny = 0
)

func (e IfnetFamilyAny) String() string {
	switch e {
	case IFNET_FAMILY_BOND:
		return "IFNET_FAMILY_BOND"
	default:
		return fmt.Sprintf("IfnetFamilyAny(%d)", e)
	}
}

type IfnetInterfaceAdvisoryDirection uint

const (
	IF_INTERFACE_ADVISORY_DIRECTION_RX IfnetInterfaceAdvisoryDirection = 0
)

func (e IfnetInterfaceAdvisoryDirection) String() string {
	switch e {
	case IF_INTERFACE_ADVISORY_DIRECTION_RX:
		return "IF_INTERFACE_ADVISORY_DIRECTION_RX"
	default:
		return fmt.Sprintf("IfnetInterfaceAdvisoryDirection(%d)", e)
	}
}

type IfnetInterfaceAdvisoryInterfaceType uint

const (
	IF_INTERFACE_ADVISORY_INTERFACE_TYPE_CELL IfnetInterfaceAdvisoryInterfaceType = 0
)

func (e IfnetInterfaceAdvisoryInterfaceType) String() string {
	switch e {
	case IF_INTERFACE_ADVISORY_INTERFACE_TYPE_CELL:
		return "IF_INTERFACE_ADVISORY_INTERFACE_TYPE_CELL"
	default:
		return fmt.Sprintf("IfnetInterfaceAdvisoryInterfaceType(%d)", e)
	}
}

type IfnetInterfaceAdvisoryNotificationTypeWifi uint

const (
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_WIFI_UNDEFINED IfnetInterfaceAdvisoryNotificationTypeWifi = 0
)

func (e IfnetInterfaceAdvisoryNotificationTypeWifi) String() string {
	switch e {
	case IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_WIFI_UNDEFINED:
		return "IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_WIFI_UNDEFINED"
	default:
		return fmt.Sprintf("IfnetInterfaceAdvisoryNotificationTypeWifi(%d)", e)
	}
}

type IfnetLqmThreshOff uint

const (
	IFNET_LQM_THRESH_MINIMALLY_VIABLE IfnetLqmThreshOff = 0
)

func (e IfnetLqmThreshOff) String() string {
	switch e {
	case IFNET_LQM_THRESH_MINIMALLY_VIABLE:
		return "IFNET_LQM_THRESH_MINIMALLY_VIABLE"
	default:
		return fmt.Sprintf("IfnetLqmThreshOff(%d)", e)
	}
}

type IfnetNpmThreshUnknown uint

const (
	IFNET_NPM_THRESH_FAR IfnetNpmThreshUnknown = 0
)

func (e IfnetNpmThreshUnknown) String() string {
	switch e {
	case IFNET_NPM_THRESH_FAR:
		return "IFNET_NPM_THRESH_FAR"
	default:
		return fmt.Sprintf("IfnetNpmThreshUnknown(%d)", e)
	}
}

type IfnetRssiUnknown uint

const (
	IFNET_RSSI_UNKNOWN IfnetRssiUnknown = 0
)

func (e IfnetRssiUnknown) String() string {
	switch e {
	case IFNET_RSSI_UNKNOWN:
		return "IFNET_RSSI_UNKNOWN"
	default:
		return fmt.Sprintf("IfnetRssiUnknown(%d)", e)
	}
}

type IfnetSchedModelNormal uint

const (
	IFNET_SCHED_MODEL_DRIVER_MANAGED IfnetSchedModelNormal = 0
)

func (e IfnetSchedModelNormal) String() string {
	switch e {
	case IFNET_SCHED_MODEL_DRIVER_MANAGED:
		return "IFNET_SCHED_MODEL_DRIVER_MANAGED"
	default:
		return fmt.Sprintf("IfnetSchedModelNormal(%d)", e)
	}
}

type IfnetThrottleOff uint

const (
	IFNET_THROTTLE_OFF IfnetThrottleOff = 0
)

func (e IfnetThrottleOff) String() string {
	switch e {
	case IFNET_THROTTLE_OFF:
		return "IFNET_THROTTLE_OFF"
	default:
		return fmt.Sprintf("IfnetThrottleOff(%d)", e)
	}
}

type IfnetWakeOnMagicPacket uint

const (
	// IFNET_WAKE_ON_MAGIC_PACKET: # Discussion
	IFNET_WAKE_ON_MAGIC_PACKET IfnetWakeOnMagicPacket = 0
)

func (e IfnetWakeOnMagicPacket) String() string {
	switch e {
	case IFNET_WAKE_ON_MAGIC_PACKET:
		return "IFNET_WAKE_ON_MAGIC_PACKET"
	default:
		return fmt.Sprintf("IfnetWakeOnMagicPacket(%d)", e)
	}
}

type In6Clat46EvhdlrCodeT uint

const (
	IN6_CLAT46_EVENT_V4_FLOW In6Clat46EvhdlrCodeT = 0
)

func (e In6Clat46EvhdlrCodeT) String() string {
	switch e {
	case IN6_CLAT46_EVENT_V4_FLOW:
		return "IN6_CLAT46_EVENT_V4_FLOW"
	default:
		return fmt.Sprintf("In6Clat46EvhdlrCodeT(%d)", e)
	}
}

type InUseErr int

const (
	AddressRangeErr InUseErr = 0
)

func (e InUseErr) String() string {
	switch e {
	case AddressRangeErr:
		return "AddressRangeErr"
	default:
		return fmt.Sprintf("InUseErr(%d)", e)
	}
}

type Input uint

const (
	INPUT_DESKTOP_MICROPHONE Input = 0
	INPUT_OMNIDIRECTIONAL_MICROPHONE Input = 0
)

func (e Input) String() string {
	switch e {
	case INPUT_DESKTOP_MICROPHONE:
		return "INPUT_DESKTOP_MICROPHONE"
	default:
		return fmt.Sprintf("Input(%d)", e)
	}
}

type Ioaudioenginetraps uint

const (
	KIOAudioEngineTrapPerformClientIO Ioaudioenginetraps = 0
)

func (e Ioaudioenginetraps) String() string {
	switch e {
	case KIOAudioEngineTrapPerformClientIO:
		return "KIOAudioEngineTrapPerformClientIO"
	default:
		return fmt.Sprintf("Ioaudioenginetraps(%d)", e)
	}
}

type Iopm uint

const (
	IOPMDeviceUsable Iopm = 0
	IOPMSoftSleep Iopm = 0
)

func (e Iopm) String() string {
	switch e {
	case IOPMDeviceUsable:
		return "IOPMDeviceUsable"
	default:
		return fmt.Sprintf("Iopm(%d)", e)
	}
}

type IpsecDscpMappingT uint

const (
	IPSEC_DSCP_MAPPING_COPY IpsecDscpMappingT = 0
)

func (e IpsecDscpMappingT) String() string {
	switch e {
	case IPSEC_DSCP_MAPPING_COPY:
		return "IPSEC_DSCP_MAPPING_COPY"
	default:
		return fmt.Sprintf("IpsecDscpMappingT(%d)", e)
	}
}

type KATA uint

const (
	KATADevice0DeviceID KATA = 0
	KATADevice1DeviceID KATA = 0
	KATAEjectRequest KATA = 0
	KATAFnBusReset KATA = 0
	KATAFnExecIO KATA = 0
	KATAFnQFlush KATA = 0
	KATAFnRegAccess KATA = 0
	KATAInvalidDeviceID KATA = 0
	KATALogicalSectorAlignmentMask KATA = 0
	KATAMultipleLogicalSectorsBit KATA = 0
	KATAMultipleLogicalSectorsMask KATA = 0
	KATANoOp KATA = 0
	KATANullEvent KATA = 0
	KATAOfflineEvent KATA = 0
	KATAOfflineRequest KATA = 0
	KATAOnlineEvent KATA = 0
	KATAPIFnExecIO KATA = 0
	KATAPIResetEvent KATA = 0
	KATAPhysicalLogicalEnabledBit1 KATA = 0
	KATARemovedEvent KATA = 0
	KATAReservedEvent KATA = 0
	KATAResetEvent KATA = 0
)

func (e KATA) String() string {
	switch e {
	case KATADevice0DeviceID:
		return "KATADevice0DeviceID"
	default:
		return fmt.Sprintf("KATA(%d)", e)
	}
}

const KATADefaultSectorSize uint = 0

type KATAEnable uint

const (
	KATAEnableAPM KATAEnable = 0
	KATAEnableWriteCache KATAEnable = 0
)

func (e KATAEnable) String() string {
	switch e {
	case KATAEnableAPM:
		return "KATAEnableAPM"
	default:
		return fmt.Sprintf("KATAEnable(%d)", e)
	}
}

type KATAEnableUltraDMAModeMask uint

const (
	KATAEnableMultiWordDMAModeMask KATAEnableUltraDMAModeMask = 0
)

func (e KATAEnableUltraDMAModeMask) String() string {
	switch e {
	case KATAEnableMultiWordDMAModeMask:
		return "KATAEnableMultiWordDMAModeMask"
	default:
		return fmt.Sprintf("KATAEnableUltraDMAModeMask(%d)", e)
	}
}

type KATAErrUnknownType int

const (
	KATADMAErr KATAErrUnknownType = 0
)

func (e KATAErrUnknownType) String() string {
	switch e {
	case KATADMAErr:
		return "KATADMAErr"
	default:
		return fmt.Sprintf("KATAErrUnknownType(%d)", e)
	}
}

const KATAForceUnitAccessFeatureBit uint = 0

const KATAForceUnitAccessFeatureMask uint = 0

type KATAIdentify uint

const (
	KATAIdentifyAdvancedPIOModes KATAIdentify = 0
	KATAIdentifyCommandExtension1 KATAIdentify = 0
	KATAIdentifyCommandExtension2 KATAIdentify = 0
	KATAIdentifyCommandSetSupported KATAIdentify = 0
	KATAIdentifyCommandSetSupported2 KATAIdentify = 0
	KATAIdentifyCommandsDefault KATAIdentify = 0
	KATAIdentifyCommandsEnabled KATAIdentify = 0
	KATAIdentifyConfiguration KATAIdentify = 0
	KATAIdentifyCurrentCapacity KATAIdentify = 0
	KATAIdentifyCurrentCylinders KATAIdentify = 0
	KATAIdentifyCurrentHeads KATAIdentify = 0
	KATAIdentifyCurrentMultipleSectors KATAIdentify = 0
	KATAIdentifyCurrentSectors KATAIdentify = 0
	KATAIdentifyDriveCapabilities KATAIdentify = 0
	KATAIdentifyDriveCapabilitiesExtended KATAIdentify = 0
	KATAIdentifyExtendedInfoSupport KATAIdentify = 0
	KATAIdentifyFirmwareRevision KATAIdentify = 0
	KATAIdentifyIntegrity KATAIdentify = 0
	KATAIdentifyLBACapacity KATAIdentify = 0
	KATAIdentifyLogicalCylinderCount KATAIdentify = 0
	KATAIdentifyLogicalHeadCount KATAIdentify = 0
	KATAIdentifyLogicalSectorAlignment KATAIdentify = 0
	KATAIdentifyMajorVersion KATAIdentify = 0
	KATAIdentifyMinMultiWordDMATime KATAIdentify = 0
	KATAIdentifyMinPIOTime KATAIdentify = 0
	KATAIdentifyMinPIOTimeWithIORDY KATAIdentify = 0
	KATAIdentifyMinorVersion KATAIdentify = 0
	KATAIdentifyModelNumber KATAIdentify = 0
	KATAIdentifyMultiWordDMA KATAIdentify = 0
	KATAIdentifyMultipleSectorCount KATAIdentify = 0
	KATAIdentifyPIOTiming KATAIdentify = 0
	KATAIdentifyPhysicalLogicalSectorSize KATAIdentify = 0
	KATAIdentifyQueueDepth KATAIdentify = 0
	KATAIdentifyRecommendedMultiWordDMATime KATAIdentify = 0
	KATAIdentifySectorsPerTrack KATAIdentify = 0
	KATAIdentifySerialNumber KATAIdentify = 0
	KATAIdentifySingleWordDMA KATAIdentify = 0
	KATAIdentifyUltraDMASupported KATAIdentify = 0
	KATAIdentifyWordsPerLogicalSector1 KATAIdentify = 0
	KATAIdentifyWordsPerLogicalSector2 KATAIdentify = 0
)

func (e KATAIdentify) String() string {
	switch e {
	case KATAIdentifyAdvancedPIOModes:
		return "KATAIdentifyAdvancedPIOModes"
	default:
		return fmt.Sprintf("KATAIdentify(%d)", e)
	}
}

type KATAOperationType uint

const (
	KATAOperationTypeFlushCache KATAOperationType = 0
	KATAOperationTypeSMART KATAOperationType = 0
)

func (e KATAOperationType) String() string {
	switch e {
	case KATAOperationTypeFlushCache:
		return "KATAOperationTypeFlushCache"
	default:
		return fmt.Sprintf("KATAOperationType(%d)", e)
	}
}

type KATAPI uint

const (
	KATAPIDRQFast KATAPI = 0
	KATAPIDRQSlow KATAPI = 0
	KATAPIIRQPacket KATAPI = 0
	KATAPIUnknown KATAPI = 0
)

func (e KATAPI) String() string {
	switch e {
	case KATAPIDRQFast:
		return "KATAPIDRQFast"
	default:
		return fmt.Sprintf("KATAPI(%d)", e)
	}
}

type KATASupports uint

const (
	KATASupportsCompactFlashMask KATASupports = 0
	KATASupportsFlushCacheMask KATASupports = 0
)

func (e KATASupports) String() string {
	switch e {
	case KATASupportsCompactFlashMask:
		return "KATASupportsCompactFlashMask"
	default:
		return fmt.Sprintf("KATASupports(%d)", e)
	}
}

type KATASupportsCompactFlashBit uint

const (
	KATASupports48BitAddressingBit KATASupportsCompactFlashBit = 0
)

func (e KATASupportsCompactFlashBit) String() string {
	switch e {
	case KATASupports48BitAddressingBit:
		return "KATASupports48BitAddressingBit"
	default:
		return fmt.Sprintf("KATASupportsCompactFlashBit(%d)", e)
	}
}

type KATASupportsSMARTBit uint

const (
	KATASupportsPowerManagementBit KATASupportsSMARTBit = 0
)

func (e KATASupportsSMARTBit) String() string {
	switch e {
	case KATASupportsPowerManagementBit:
		return "KATASupportsPowerManagementBit"
	default:
		return fmt.Sprintf("KATASupportsSMARTBit(%d)", e)
	}
}

type KATASupportsSMARTMask uint

const (
	KATASupportsPowerManagementMask KATASupportsSMARTMask = 0
)

func (e KATASupportsSMARTMask) String() string {
	switch e {
	case KATASupportsPowerManagementMask:
		return "KATASupportsPowerManagementMask"
	default:
		return fmt.Sprintf("KATASupportsSMARTMask(%d)", e)
	}
}

type KATATimeout10Seconds uint

const (
	KATADefaultTimeout KATATimeout10Seconds = 0
)

func (e KATATimeout10Seconds) String() string {
	switch e {
	case KATADefaultTimeout:
		return "KATADefaultTimeout"
	default:
		return fmt.Sprintf("KATATimeout10Seconds(%d)", e)
	}
}

const KATAWriteCacheEnabledBit uint = 0

const KATAWriteCacheEnabledMask uint = 0

type KATAZeroRetries uint

const (
	KATADefaultRetries KATAZeroRetries = 0
)

func (e KATAZeroRetries) String() string {
	switch e {
	case KATADefaultRetries:
		return "KATADefaultRetries"
	default:
		return fmt.Sprintf("KATAZeroRetries(%d)", e)
	}
}

type KATAcmdWORetry uint

const (
	KATAcmdFlushCacheExtended KATAcmdWORetry = 0
)

func (e KATAcmdWORetry) String() string {
	switch e {
	case KATAcmdFlushCacheExtended:
		return "KATAcmdFlushCacheExtended"
	default:
		return fmt.Sprintf("KATAcmdWORetry(%d)", e)
	}
}

type KAVC uint

const (
	KAVCAcceptedStatus KAVC = 0
	KAVCAddress KAVC = 0
	KAVCAudio KAVC = 0
	KAVCCameraStorage KAVC = 0
	KAVCChangedStatus KAVC = 0
	KAVCCommandResponse KAVC = 0
	KAVCConnectOpcode KAVC = 0
	KAVCConnectionsOpcode KAVC = 0
	KAVCControlCommand KAVC = 0
	KAVCDisconnectOpcode KAVC = 0
	KAVCDiskRecorder KAVC = 0
	KAVCGeneralInquiryCommand KAVC = 0
	KAVCImplementedStatus KAVC = 0
	KAVCInTransitionStatus KAVC = 0
	KAVCInputPlugSignalFormatOpcode KAVC = 0
	KAVCInputSignalModeOpcode KAVC = 0
	KAVCInterimStatus KAVC = 0
	KAVCNotImplementedStatus KAVC = 0
	KAVCNotifyCommand KAVC = 0
	KAVCNumSubUnitTypes KAVC = 0
	KAVCOpcode KAVC = 0
	KAVCOperand0 KAVC = 0
	KAVCOperand1 KAVC = 0
	KAVCOperand2 KAVC = 0
	KAVCOperand3 KAVC = 0
	KAVCOperand4 KAVC = 0
	KAVCOperand5 KAVC = 0
	KAVCOperand6 KAVC = 0
	KAVCOperand7 KAVC = 0
	KAVCOperand8 KAVC = 0
	KAVCOutputPlugSignalFormatOpcode KAVC = 0
	KAVCOutputSignalModeOpcode KAVC = 0
	KAVCPlugInfoOpcode KAVC = 0
	KAVCPowerOpcode KAVC = 0
	KAVCPrinter KAVC = 0
	KAVCRejectedStatus KAVC = 0
	KAVCSignalModeDVCPro525_60 KAVC = 0
	KAVCSignalModeDVCPro625_50 KAVC = 0
	KAVCSignalModeDummyOperand KAVC = 0
	KAVCSignalModeHD1125_60 KAVC = 0
	KAVCSignalModeHD1250_50 KAVC = 0
	KAVCSignalModeMask_50 KAVC = 0
	KAVCSignalModeMask_DVCPro25 KAVC = 0
	KAVCSignalModeMask_SDL KAVC = 0
	KAVCSignalModeMask_STYPE KAVC = 0
	KAVCSignalModeSD525_60 KAVC = 0
	KAVCSignalModeSD625_50 KAVC = 0
	KAVCSignalModeSDL525_60 KAVC = 0
	KAVCSignalModeSDL625_50 KAVC = 0
	KAVCSignalSourceOpcode KAVC = 0
	KAVCSpecificInquiryCommand KAVC = 0
	KAVCStatusInquiryCommand KAVC = 0
	KAVCSubunitInfoOpcode KAVC = 0
	KAVCTapeRecorder KAVC = 0
	KAVCTuner KAVC = 0
	KAVCUnitInfoOpcode KAVC = 0
	KAVCVendorDependentOpcode KAVC = 0
	KAVCVendorUnique KAVC = 0
	KAVCVideoCamera KAVC = 0
	KAVCVideoMonitor KAVC = 0
)

func (e KAVC) String() string {
	switch e {
	case KAVCAcceptedStatus:
		return "KAVCAcceptedStatus"
	default:
		return fmt.Sprintf("KAVC(%d)", e)
	}
}

type KAVCAsyncCommandState uint

const (
	KAVCAsyncCommandStateBusReset KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateCanceled KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateOutOfMemory KAVCAsyncCommandState = 0
	KAVCAsyncCommandStatePendingRequest KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateReceivedFinalResponse KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateReceivedInterimResponse KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateRequestFailed KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateRequestSent KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateTimeOutBeforeResponse KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateWaitingForResponse KAVCAsyncCommandState = 0
)

func (e KAVCAsyncCommandState) String() string {
	switch e {
	case KAVCAsyncCommandStateBusReset:
		return "KAVCAsyncCommandStateBusReset"
	default:
		return fmt.Sprintf("KAVCAsyncCommandState(%d)", e)
	}
}

type KAVPower uint

const (
	KAVPowerOff KAVPower = 0
	KAVPowerOn KAVPower = 0
)

func (e KAVPower) String() string {
	switch e {
	case KAVPowerOff:
		return "KAVPowerOff"
	default:
		return fmt.Sprintf("KAVPower(%d)", e)
	}
}

type KAllModesValid uint

const (
	KAllModesValidValue KAllModesValid = 0
	KFastCheckForDDC KAllModesValid = 0
)

func (e KAllModesValid) String() string {
	switch e {
	case KAllModesValidValue:
		return "KAllModesValidValue"
	default:
		return fmt.Sprintf("KAllModesValid(%d)", e)
	}
}

type KAppleBasebandBridgeUserRequest uint

const (
	KAppleBasebandBridgeUserRequest_BasebandHostWake KAppleBasebandBridgeUserRequest = 0
	KAppleBasebandBridgeUserRequest_BasebandInReset KAppleBasebandBridgeUserRequest = 0
	KAppleBasebandBridgeUserRequest_BridgeReady KAppleBasebandBridgeUserRequest = 0
	KAppleBasebandBridgeUserRequest_MAX KAppleBasebandBridgeUserRequest = 0
)

func (e KAppleBasebandBridgeUserRequest) String() string {
	switch e {
	case KAppleBasebandBridgeUserRequest_BasebandHostWake:
		return "KAppleBasebandBridgeUserRequest_BasebandHostWake"
	default:
		return fmt.Sprintf("KAppleBasebandBridgeUserRequest(%d)", e)
	}
}

const KAppleVendorID uint = 0

const KBDFeaturesReadBit uint = 0

const KBDFeaturesReadMask uint = 0

type KBDMediaType uint

const (
	KBDMediaTypeMin KBDMediaType = 0
	KBDMediaTypeR KBDMediaType = 0
)

func (e KBDMediaType) String() string {
	switch e {
	case KBDMediaTypeMin:
		return "KBDMediaTypeMin"
	default:
		return fmt.Sprintf("KBDMediaType(%d)", e)
	}
}

const KBTBadCloseMask uint = 0

type KBTLeafNode uint

const (
	KBTHeaderNode KBTLeafNode = 0
)

func (e KBTLeafNode) String() string {
	switch e {
	case KBTHeaderNode:
		return "KBTHeaderNode"
	default:
		return fmt.Sprintf("KBTLeafNode(%d)", e)
	}
}

type KBluetooth uint

const (
	KBluetoothACLLogicalChannelLMP KBluetooth = 0
	KBluetoothL2CAPMaxPacketSize KBluetooth = 0
)

func (e KBluetooth) String() string {
	switch e {
	case KBluetoothACLLogicalChannelLMP:
		return "KBluetoothACLLogicalChannelLMP"
	default:
		return fmt.Sprintf("KBluetooth(%d)", e)
	}
}

type KBluetoothAirMode uint

const (
	KBluetoothAirModeALawLog KBluetoothAirMode = 0
	KBluetoothAirModeULawLog KBluetoothAirMode = 0
)

func (e KBluetoothAirMode) String() string {
	switch e {
	case KBluetoothAirModeALawLog:
		return "KBluetoothAirModeALawLog"
	default:
		return fmt.Sprintf("KBluetoothAirMode(%d)", e)
	}
}

const KBluetoothConnectionHandleNone uint = 0

type KBluetoothDeviceClassMajorMiscellaneous uint

const (
	KBluetoothDeviceClassMajorAudio KBluetoothDeviceClassMajorMiscellaneous = 0
)

func (e KBluetoothDeviceClassMajorMiscellaneous) String() string {
	switch e {
	case KBluetoothDeviceClassMajorAudio:
		return "KBluetoothDeviceClassMajorAudio"
	default:
		return fmt.Sprintf("KBluetoothDeviceClassMajorMiscellaneous(%d)", e)
	}
}

type KBluetoothDeviceClassMinor uint

const (
	KBluetoothDeviceClassMinorHealthGlucoseMeter KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2CardReader KBluetoothDeviceClassMinor = 0
)

func (e KBluetoothDeviceClassMinor) String() string {
	switch e {
	case KBluetoothDeviceClassMinorHealthGlucoseMeter:
		return "KBluetoothDeviceClassMinorHealthGlucoseMeter"
	default:
		return fmt.Sprintf("KBluetoothDeviceClassMinor(%d)", e)
	}
}

const KBluetoothDeviceNameMaxLength uint = 0

type KBluetoothDontAllowRoleSwitch uint

const (
	KBluetoothAllowRoleSwitch KBluetoothDontAllowRoleSwitch = 0
)

func (e KBluetoothDontAllowRoleSwitch) String() string {
	switch e {
	case KBluetoothAllowRoleSwitch:
		return "KBluetoothAllowRoleSwitch"
	default:
		return fmt.Sprintf("KBluetoothDontAllowRoleSwitch(%d)", e)
	}
}

type KBluetoothEncryptionEnable uint

const (
	KBluetoothEncryptionEnableOff KBluetoothEncryptionEnable = 0
	KBluetoothEncryptionEnableOn KBluetoothEncryptionEnable = 0
)

func (e KBluetoothEncryptionEnable) String() string {
	switch e {
	case KBluetoothEncryptionEnableOff:
		return "KBluetoothEncryptionEnableOff"
	default:
		return fmt.Sprintf("KBluetoothEncryptionEnable(%d)", e)
	}
}

type KBluetoothGAPAppearanceGeneric uint

const (
	KBluetoothGAPAppearanceGenericBarcodeScanner KBluetoothGAPAppearanceGeneric = 0
	KBluetoothGAPAppearanceGenericHumanInterfaceDevice KBluetoothGAPAppearanceGeneric = 0
)

func (e KBluetoothGAPAppearanceGeneric) String() string {
	switch e {
	case KBluetoothGAPAppearanceGenericBarcodeScanner:
		return "KBluetoothGAPAppearanceGenericBarcodeScanner"
	default:
		return fmt.Sprintf("KBluetoothGAPAppearanceGeneric(%d)", e)
	}
}

const KBluetoothGeneralInquiryAccessCodeIndex uint = 0

type KBluetoothHCICommand uint

const (
	KBluetoothHCICommandLEWriteSuggestedDefaultDataLength KBluetoothHCICommand = 0
	KBluetoothHCICommandReadLocalAMPASSOC KBluetoothHCICommand = 0
)

func (e KBluetoothHCICommand) String() string {
	switch e {
	case KBluetoothHCICommandLEWriteSuggestedDefaultDataLength:
		return "KBluetoothHCICommandLEWriteSuggestedDefaultDataLength"
	default:
		return fmt.Sprintf("KBluetoothHCICommand(%d)", e)
	}
}

const KBluetoothHCICommandPacketHeaderSize uint = 0

const KBluetoothHCIErroneousDataReportingDisabled int = 0

type KBluetoothHCIError int

const (
	KBluetoothHCIErrorACLConnectionAlreadyExists KBluetoothHCIError = 0
	KBluetoothHCIErrorConnectionTerminatedByLocalHost KBluetoothHCIError = 0
)

func (e KBluetoothHCIError) String() string {
	switch e {
	case KBluetoothHCIErrorACLConnectionAlreadyExists:
		return "KBluetoothHCIErrorACLConnectionAlreadyExists"
	default:
		return fmt.Sprintf("KBluetoothHCIError(%d)", e)
	}
}

const KBluetoothHCIErrorPowerIsOFF int = 0

type KBluetoothHCIEventMask uint

const (
	KBluetoothHCIEventMaskConnectionPacketTypeChanged KBluetoothHCIEventMask = 0
	KBluetoothHCIEventMaskFlushOccurred KBluetoothHCIEventMask = 0
)

func (e KBluetoothHCIEventMask) String() string {
	switch e {
	case KBluetoothHCIEventMaskConnectionPacketTypeChanged:
		return "KBluetoothHCIEventMaskConnectionPacketTypeChanged"
	default:
		return fmt.Sprintf("KBluetoothHCIEventMask(%d)", e)
	}
}

type KBluetoothHCILoopbackMode uint

const (
	KBluetoothHCILoopbackModeLocal KBluetoothHCILoopbackMode = 0
	KBluetoothHCILoopbackModeRemote KBluetoothHCILoopbackMode = 0
)

func (e KBluetoothHCILoopbackMode) String() string {
	switch e {
	case KBluetoothHCILoopbackModeLocal:
		return "KBluetoothHCILoopbackModeLocal"
	default:
		return fmt.Sprintf("KBluetoothHCILoopbackMode(%d)", e)
	}
}

type KBluetoothHCISubEventLE uint

const (
	KBluetoothHCISubEventLEDataLengthChange KBluetoothHCISubEventLE = 0
	KBluetoothHCISubEventLEScanRequestReceived KBluetoothHCISubEventLE = 0
)

func (e KBluetoothHCISubEventLE) String() string {
	switch e {
	case KBluetoothHCISubEventLEDataLengthChange:
		return "KBluetoothHCISubEventLEDataLengthChange"
	default:
		return fmt.Sprintf("KBluetoothHCISubEventLE(%d)", e)
	}
}

const KBluetoothHCITransportUSBClassCode uint = 0

const KBluetoothKeyFlagSemiPermanent uint = 0

type KBluetoothKeyType uint

const (
	KBluetoothKeyTypeAuthenticatedCombination KBluetoothKeyType = 0
	KBluetoothKeyTypeCombination KBluetoothKeyType = 0
)

func (e KBluetoothKeyType) String() string {
	switch e {
	case KBluetoothKeyTypeAuthenticatedCombination:
		return "KBluetoothKeyTypeAuthenticatedCombination"
	default:
		return fmt.Sprintf("KBluetoothKeyType(%d)", e)
	}
}

type KBluetoothL2CAP uint

const (
	KBluetoothL2CAPMTUDefault KBluetoothL2CAP = 0
	KBluetoothL2CAPQoSDelayVariationDefault KBluetoothL2CAP = 0
)

func (e KBluetoothL2CAP) String() string {
	switch e {
	case KBluetoothL2CAPMTUDefault:
		return "KBluetoothL2CAPMTUDefault"
	default:
		return fmt.Sprintf("KBluetoothL2CAP(%d)", e)
	}
}

type KBluetoothL2CAPChannel uint

const (
	KBluetoothL2CAPChannelLESignalling KBluetoothL2CAPChannel = 0
	KBluetoothL2CAPChannelMagnet KBluetoothL2CAPChannel = 0
)

func (e KBluetoothL2CAPChannel) String() string {
	switch e {
	case KBluetoothL2CAPChannelLESignalling:
		return "KBluetoothL2CAPChannelLESignalling"
	default:
		return fmt.Sprintf("KBluetoothL2CAPChannel(%d)", e)
	}
}

type KBluetoothL2CAPConfigurationOption uint

const (
	KBluetoothL2CAPConfigurationOptionFlushTimeoutLength KBluetoothL2CAPConfigurationOption = 0
	KBluetoothL2CAPConfigurationOptionRetransmissionAndFlowControlLength KBluetoothL2CAPConfigurationOption = 0
)

func (e KBluetoothL2CAPConfigurationOption) String() string {
	switch e {
	case KBluetoothL2CAPConfigurationOptionFlushTimeoutLength:
		return "KBluetoothL2CAPConfigurationOptionFlushTimeoutLength"
	default:
		return fmt.Sprintf("KBluetoothL2CAPConfigurationOption(%d)", e)
	}
}

type KBluetoothL2CAPFlushTimeout uint

const (
	KBluetoothL2CAPFlushTimeoutForever KBluetoothL2CAPFlushTimeout = 0
	KBluetoothL2CAPFlushTimeoutUseExisting KBluetoothL2CAPFlushTimeout = 0
)

func (e KBluetoothL2CAPFlushTimeout) String() string {
	switch e {
	case KBluetoothL2CAPFlushTimeoutForever:
		return "KBluetoothL2CAPFlushTimeoutForever"
	default:
		return fmt.Sprintf("KBluetoothL2CAPFlushTimeout(%d)", e)
	}
}

const KBluetoothL2CAPInfoTypeMaxConnectionlessMTUSize uint = 0

type KBluetoothL2CAPPSM uint

const (
	KBluetoothL2CAPPSMAVCTP_Browsing KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMHIDControl KBluetoothL2CAPPSM = 0
)

func (e KBluetoothL2CAPPSM) String() string {
	switch e {
	case KBluetoothL2CAPPSMAVCTP_Browsing:
		return "KBluetoothL2CAPPSMAVCTP_Browsing"
	default:
		return fmt.Sprintf("KBluetoothL2CAPPSM(%d)", e)
	}
}

const KBluetoothL2CAPPacketHeaderSize uint = 0

type KBluetoothL2CAPTCICommandL2CA uint

const (
	KBluetoothL2CAPTCICommandL2CA_ConfigReq KBluetoothL2CAPTCICommandL2CA = 0
	KBluetoothL2CAPTCICommandL2CA_ConnectReq KBluetoothL2CAPTCICommandL2CA = 0
)

func (e KBluetoothL2CAPTCICommandL2CA) String() string {
	switch e {
	case KBluetoothL2CAPTCICommandL2CA_ConfigReq:
		return "KBluetoothL2CAPTCICommandL2CA_ConfigReq"
	default:
		return fmt.Sprintf("KBluetoothL2CAPTCICommandL2CA(%d)", e)
	}
}

type KBluetoothL2CAPTCIEventIDL2CA uint

const (
	KBluetoothL2CAPTCIEventIDL2CA_ConfigInd KBluetoothL2CAPTCIEventIDL2CA = 0
	KBluetoothL2CAPTCIEventIDL2CA_TimeOutInd KBluetoothL2CAPTCIEventIDL2CA = 0
)

func (e KBluetoothL2CAPTCIEventIDL2CA) String() string {
	switch e {
	case KBluetoothL2CAPTCIEventIDL2CA_ConfigInd:
		return "KBluetoothL2CAPTCIEventIDL2CA_ConfigInd"
	default:
		return fmt.Sprintf("KBluetoothL2CAPTCIEventIDL2CA(%d)", e)
	}
}

type KBluetoothLEMaxTXTimeMin uint

const (
	KBluetoothLEMaxTXOctetsDefault KBluetoothLEMaxTXTimeMin = 0
)

func (e KBluetoothLEMaxTXTimeMin) String() string {
	switch e {
	case KBluetoothLEMaxTXOctetsDefault:
		return "KBluetoothLEMaxTXOctetsDefault"
	default:
		return fmt.Sprintf("KBluetoothLEMaxTXTimeMin(%d)", e)
	}
}

type KBluetoothLESecurityManagerNoBonding uint

const (
	KBluetoothLESecurityManagerBonding KBluetoothLESecurityManagerNoBonding = 0
)

func (e KBluetoothLESecurityManagerNoBonding) String() string {
	switch e {
	case KBluetoothLESecurityManagerBonding:
		return "KBluetoothLESecurityManagerBonding"
	default:
		return fmt.Sprintf("KBluetoothLESecurityManagerNoBonding(%d)", e)
	}
}

type KBluetoothLETXTimeMin uint

const (
	KBluetoothLETXOctetsDefault KBluetoothLETXTimeMin = 0
)

func (e KBluetoothLETXTimeMin) String() string {
	switch e {
	case KBluetoothLETXOctetsDefault:
		return "KBluetoothLETXOctetsDefault"
	default:
		return fmt.Sprintf("KBluetoothLETXTimeMin(%d)", e)
	}
}

type KBluetoothPacketType uint

const (
	KBluetoothPacketTypeDH1 KBluetoothPacketType = 0
	KBluetoothPacketTypeHV2 KBluetoothPacketType = 0
)

func (e KBluetoothPacketType) String() string {
	switch e {
	case KBluetoothPacketTypeDH1:
		return "KBluetoothPacketTypeDH1"
	default:
		return fmt.Sprintf("KBluetoothPacketType(%d)", e)
	}
}

const KBluetoothPageScanModeMandatory uint = 0

const KBluetoothPageScanPeriodModeP0 uint = 0

type KBluetoothPageScanRepetitionMode uint

const (
	KBluetoothPageScanRepetitionModeR1 KBluetoothPageScanRepetitionMode = 0
	KBluetoothPageScanRepetitionModeR2 KBluetoothPageScanRepetitionMode = 0
)

func (e KBluetoothPageScanRepetitionMode) String() string {
	switch e {
	case KBluetoothPageScanRepetitionModeR1:
		return "KBluetoothPageScanRepetitionModeR1"
	default:
		return fmt.Sprintf("KBluetoothPageScanRepetitionMode(%d)", e)
	}
}

type KBluetoothRole uint

const (
	KBluetoothRoleBecomeCentral KBluetoothRole = 0
	KBluetoothRoleRemainPeripheral KBluetoothRole = 0
	// Deprecated.
	KBluetoothRoleBecomeMaster KBluetoothRole = 0
	// Deprecated.
	KBluetoothRoleRemainSlave KBluetoothRole = 0
)

func (e KBluetoothRole) String() string {
	switch e {
	case KBluetoothRoleBecomeCentral:
		return "KBluetoothRoleBecomeCentral"
	default:
		return fmt.Sprintf("KBluetoothRole(%d)", e)
	}
}

type KBluetoothSDPAttributeDeviceIdentifier uint

const (
	KBluetoothSDPAttributeDeviceIdentifierClientExecutableURL KBluetoothSDPAttributeDeviceIdentifier = 0
	KBluetoothSDPAttributeDeviceIdentifierDocumentationURL KBluetoothSDPAttributeDeviceIdentifier = 0
)

func (e KBluetoothSDPAttributeDeviceIdentifier) String() string {
	switch e {
	case KBluetoothSDPAttributeDeviceIdentifierClientExecutableURL:
		return "KBluetoothSDPAttributeDeviceIdentifierClientExecutableURL"
	default:
		return fmt.Sprintf("KBluetoothSDPAttributeDeviceIdentifier(%d)", e)
	}
}

type KBluetoothSDPDataElementTypeNil uint

const (
	KBluetoothSDPDataElementTypeBoolean KBluetoothSDPDataElementTypeNil = 0
)

func (e KBluetoothSDPDataElementTypeNil) String() string {
	switch e {
	case KBluetoothSDPDataElementTypeBoolean:
		return "KBluetoothSDPDataElementTypeBoolean"
	default:
		return fmt.Sprintf("KBluetoothSDPDataElementTypeNil(%d)", e)
	}
}

type KBluetoothSDPErrorCodeReserved int

const (
	KBluetoothSDPErrorCodeReservedEnd KBluetoothSDPErrorCodeReserved = 0
	KBluetoothSDPErrorCodeReservedStart KBluetoothSDPErrorCodeReserved = 0
)

func (e KBluetoothSDPErrorCodeReserved) String() string {
	switch e {
	case KBluetoothSDPErrorCodeReservedEnd:
		return "KBluetoothSDPErrorCodeReservedEnd"
	default:
		return fmt.Sprintf("KBluetoothSDPErrorCodeReserved(%d)", e)
	}
}

type KBluetoothSDPPDUID uint

const (
	KBluetoothSDPPDUIDReserved KBluetoothSDPPDUID = 0
	KBluetoothSDPPDUIDServiceSearchAttributeResponse KBluetoothSDPPDUID = 0
)

func (e KBluetoothSDPPDUID) String() string {
	switch e {
	case KBluetoothSDPPDUIDReserved:
		return "KBluetoothSDPPDUIDReserved"
	default:
		return fmt.Sprintf("KBluetoothSDPPDUID(%d)", e)
	}
}

type KBluetoothSDPProtocolParameter uint

const (
	KBluetoothSDPProtocolParameterBNEPSupportedNetworkPacketTypeList KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterBNEPVersion KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterL2CAPPSM KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterRFCOMMChannel KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterTCPPort KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterUDPPort KBluetoothSDPProtocolParameter = 0
)

func (e KBluetoothSDPProtocolParameter) String() string {
	switch e {
	case KBluetoothSDPProtocolParameterBNEPSupportedNetworkPacketTypeList:
		return "KBluetoothSDPProtocolParameterBNEPSupportedNetworkPacketTypeList"
	default:
		return fmt.Sprintf("KBluetoothSDPProtocolParameter(%d)", e)
	}
}

type KBluetoothSDPUUID16Base uint

const (
	KBluetoothSDPUUID16AVCTP KBluetoothSDPUUID16Base = 0
)

func (e KBluetoothSDPUUID16Base) String() string {
	switch e {
	case KBluetoothSDPUUID16AVCTP:
		return "KBluetoothSDPUUID16AVCTP"
	default:
		return fmt.Sprintf("KBluetoothSDPUUID16Base(%d)", e)
	}
}

type KBluetoothSDPUUID16ServiceClass uint

const (
	KBluetoothSDPUUID16ServiceClassAVRemoteControl KBluetoothSDPUUID16ServiceClass = 0
	KBluetoothSDPUUID16ServiceClassDialupNetworking KBluetoothSDPUUID16ServiceClass = 0
	KBluetoothSDPUUID16ServiceClassGN KBluetoothSDPUUID16ServiceClass = 0
)

func (e KBluetoothSDPUUID16ServiceClass) String() string {
	switch e {
	case KBluetoothSDPUUID16ServiceClassAVRemoteControl:
		return "KBluetoothSDPUUID16ServiceClassAVRemoteControl"
	default:
		return fmt.Sprintf("KBluetoothSDPUUID16ServiceClass(%d)", e)
	}
}

type KBluetoothServiceClassMajor uint

const (
	KBluetoothServiceClassMajorAny KBluetoothServiceClassMajor = 0
	KBluetoothServiceClassMajorCapturing KBluetoothServiceClassMajor = 0
)

func (e KBluetoothServiceClassMajor) String() string {
	switch e {
	case KBluetoothServiceClassMajorAny:
		return "KBluetoothServiceClassMajorAny"
	default:
		return fmt.Sprintf("KBluetoothServiceClassMajor(%d)", e)
	}
}

type KBluetoothSynchronousConnectionPacketTypeNone uint

const (
	KBluetoothSynchronousConnectionPacketType3EV5Omit KBluetoothSynchronousConnectionPacketTypeNone = 0
)

func (e KBluetoothSynchronousConnectionPacketTypeNone) String() string {
	switch e {
	case KBluetoothSynchronousConnectionPacketType3EV5Omit:
		return "KBluetoothSynchronousConnectionPacketType3EV5Omit"
	default:
		return fmt.Sprintf("KBluetoothSynchronousConnectionPacketTypeNone(%d)", e)
	}
}

type KBluetoothVoiceSettingAirCodingFormatMask uint

const (
	KBluetoothVoiceSettingAirCodingFormatALaw KBluetoothVoiceSettingAirCodingFormatMask = 0
)

func (e KBluetoothVoiceSettingAirCodingFormatMask) String() string {
	switch e {
	case KBluetoothVoiceSettingAirCodingFormatALaw:
		return "KBluetoothVoiceSettingAirCodingFormatALaw"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingAirCodingFormatMask(%d)", e)
	}
}

type KBluetoothVoiceSettingInputCodingMask uint

const (
	KBluetoothVoiceSettingInputCodingALawInputCoding KBluetoothVoiceSettingInputCodingMask = 0
)

func (e KBluetoothVoiceSettingInputCodingMask) String() string {
	switch e {
	case KBluetoothVoiceSettingInputCodingALawInputCoding:
		return "KBluetoothVoiceSettingInputCodingALawInputCoding"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingInputCodingMask(%d)", e)
	}
}

type KBluetoothVoiceSettingInputDataFormat uint

const (
	KBluetoothVoiceSettingInputDataFormatMask KBluetoothVoiceSettingInputDataFormat = 0
	KBluetoothVoiceSettingInputDataFormatSignMagnitude KBluetoothVoiceSettingInputDataFormat = 0
)

func (e KBluetoothVoiceSettingInputDataFormat) String() string {
	switch e {
	case KBluetoothVoiceSettingInputDataFormatMask:
		return "KBluetoothVoiceSettingInputDataFormatMask"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingInputDataFormat(%d)", e)
	}
}

type KBluetoothVoiceSettingInputSampleSizeMask uint

const (
	KBluetoothVoiceSettingInputSampleSize16Bit KBluetoothVoiceSettingInputSampleSizeMask = 0
)

func (e KBluetoothVoiceSettingInputSampleSizeMask) String() string {
	switch e {
	case KBluetoothVoiceSettingInputSampleSize16Bit:
		return "KBluetoothVoiceSettingInputSampleSize16Bit"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingInputSampleSizeMask(%d)", e)
	}
}

const KBluetoothVoiceSettingPCMBitPositionMask uint = 0

const KBootDriverTypeInvalid uint = 0

type KBootROMTypeOldWorld uint

const (
	KBootROMTypeNewWorld KBootROMTypeOldWorld = 0
)

func (e KBootROMTypeOldWorld) String() string {
	switch e {
	case KBootROMTypeNewWorld:
		return "KBootROMTypeNewWorld"
	default:
		return fmt.Sprintf("KBootROMTypeOldWorld(%d)", e)
	}
}

const KC0DataMaxStringLen uint = 0

type KCDFeatures uint

const (
	KCDFeaturesAnalogAudioBit KCDFeatures = 0
	KCDFeaturesAnalogAudioMask KCDFeatures = 0
	KCDFeaturesBUFWriteBit KCDFeatures = 0
	KCDFeaturesTestWriteMask KCDFeatures = 0
)

func (e KCDFeatures) String() string {
	switch e {
	case KCDFeaturesAnalogAudioBit:
		return "KCDFeaturesAnalogAudioBit"
	default:
		return fmt.Sprintf("KCDFeatures(%d)", e)
	}
}

type KCDMediaType uint

const (
	KCDMediaTypeMax KCDMediaType = 0
	KCDMediaTypeUnknown KCDMediaType = 0
)

func (e KCDMediaType) String() string {
	switch e {
	case KCDMediaTypeMax:
		return "KCDMediaTypeMax"
	default:
		return fmt.Sprintf("KCDMediaType(%d)", e)
	}
}

type KCDSectorArea uint

const (
	KCDSectorAreaAuxiliary KCDSectorArea = 0
	KCDSectorAreaErrorFlags KCDSectorArea = 0
	KCDSectorAreaHeader KCDSectorArea = 0
	KCDSectorAreaSubChannel KCDSectorArea = 0
	KCDSectorAreaSubChannelQ KCDSectorArea = 0
	KCDSectorAreaSubHeader KCDSectorArea = 0
	KCDSectorAreaSync KCDSectorArea = 0
	KCDSectorAreaUser KCDSectorArea = 0
)

func (e KCDSectorArea) String() string {
	switch e {
	case KCDSectorAreaAuxiliary:
		return "KCDSectorAreaAuxiliary"
	default:
		return fmt.Sprintf("KCDSectorArea(%d)", e)
	}
}

type KCDSectorSize uint

const (
	KCDSectorSizeCDDA KCDSectorSize = 0
	KCDSectorSizeMode1 KCDSectorSize = 0
	KCDSectorSizeMode2 KCDSectorSize = 0
	KCDSectorSizeMode2Form1 KCDSectorSize = 0
	KCDSectorSizeMode2Form2 KCDSectorSize = 0
	KCDSectorSizeWhole KCDSectorSize = 0
)

func (e KCDSectorSize) String() string {
	switch e {
	case KCDSectorSizeCDDA:
		return "KCDSectorSizeCDDA"
	default:
		return fmt.Sprintf("KCDSectorSize(%d)", e)
	}
}

type KCDSectorType uint

const (
	KCDSectorTypeCDDA KCDSectorType = 0
	KCDSectorTypeCount KCDSectorType = 0
	KCDSectorTypeMode1 KCDSectorType = 0
	KCDSectorTypeMode2 KCDSectorType = 0
	KCDSectorTypeMode2Form1 KCDSectorType = 0
	KCDSectorTypeMode2Form2 KCDSectorType = 0
	KCDSectorTypeUnknown KCDSectorType = 0
)

func (e KCDSectorType) String() string {
	switch e {
	case KCDSectorTypeCDDA:
		return "KCDSectorTypeCDDA"
	default:
		return fmt.Sprintf("KCDSectorType(%d)", e)
	}
}

type KCDTOCFormatTOC uint

const (
	KCDTOCFormatATIP KCDTOCFormatTOC = 0
)

func (e KCDTOCFormatTOC) String() string {
	switch e {
	case KCDTOCFormatATIP:
		return "KCDTOCFormatATIP"
	default:
		return fmt.Sprintf("KCDTOCFormatTOC(%d)", e)
	}
}

type KCDTrackInfoAddressType uint

const (
	KCDTrackInfoAddressTypeLBA KCDTrackInfoAddressType = 0
	KCDTrackInfoAddressTypeTrackNumber KCDTrackInfoAddressType = 0
)

func (e KCDTrackInfoAddressType) String() string {
	switch e {
	case KCDTrackInfoAddressTypeLBA:
		return "KCDTrackInfoAddressTypeLBA"
	default:
		return fmt.Sprintf("KCDTrackInfoAddressType(%d)", e)
	}
}

type KCFormat uint

const (
	KCFormatDynamic KCFormat = 0
	KCFormatFileset KCFormat = 0
)

func (e KCFormat) String() string {
	switch e {
	case KCFormatDynamic:
		return "KCFormatDynamic"
	default:
		return fmt.Sprintf("KCFormat(%d)", e)
	}
}

type KCKind uint

const (
	KCKindAuxiliary KCKind = 0
	KCKindNone KCKind = 0
)

func (e KCKind) String() string {
	switch e {
	case KCKindAuxiliary:
		return "KCKindAuxiliary"
	default:
		return fmt.Sprintf("KCKind(%d)", e)
	}
}

type KCSR uint

const (
	KCSRArgumentHiAddress KCSR = 0
	KCSRClockInfo2Address KCSR = 0
)

func (e KCSR) String() string {
	switch e {
	case KCSRArgumentHiAddress:
		return "KCSRArgumentHiAddress"
	default:
		return fmt.Sprintf("KCSR(%d)", e)
	}
}

type KCSRState uint

const (
	KCSRStateOff KCSRState = 0
	KCSRStateStateTesting KCSRState = 0
)

func (e KCSRState) String() string {
	switch e {
	case KCSRStateOff:
		return "KCSRStateOff"
	default:
		return fmt.Sprintf("KCSRState(%d)", e)
	}
}

type KChipSetType uint

const (
	KChipSetTypeCore2001 KChipSetType = 0
	KChipSetTypePowerExpress KChipSetType = 0
)

func (e KChipSetType) String() string {
	switch e {
	case KChipSetTypeCore2001:
		return "KChipSetTypeCore2001"
	default:
		return fmt.Sprintf("KChipSetType(%d)", e)
	}
}

type KClamshellStateBit uint

const (
	KClamshellSleepBit KClamshellStateBit = 0
)

func (e KClamshellStateBit) String() string {
	switch e {
	case KClamshellSleepBit:
		return "KClamshellSleepBit"
	default:
		return fmt.Sprintf("KClamshellStateBit(%d)", e)
	}
}

const KClearDeviceFeature uint = 0

const KClearHubFeature uint = 0

type KConfig uint

const (
	KConfigBusDependentInfoKey KConfig = 0
	KConfigBusInfoBlockLength KConfig = 0
	KConfigEntryKeyTypePhase KConfig = 0
	KConfigNodeUniqueIdKey KConfig = 0
)

func (e KConfig) String() string {
	switch e {
	case KConfigBusDependentInfoKey:
		return "KConfigBusDependentInfoKey"
	default:
		return fmt.Sprintf("KConfig(%d)", e)
	}
}

type KConfigSBP2 uint

const (
	KConfigSBP2LUN KConfigSBP2 = 0
	KConfigSBP2Revision KConfigSBP2 = 0
)

func (e KConfigSBP2) String() string {
	switch e {
	case KConfigSBP2LUN:
		return "KConfigSBP2LUN"
	default:
		return fmt.Sprintf("KConfigSBP2(%d)", e)
	}
}

type KConfigUnit uint

const (
	KConfigUnitSWVersIIDC100 KConfigUnit = 0
	KConfigUnitSpecAppleA27 KConfigUnit = 0
)

func (e KConfigUnit) String() string {
	switch e {
	case KConfigUnitSWVersIIDC100:
		return "KConfigUnitSWVersIIDC100"
	default:
		return fmt.Sprintf("KConfigUnit(%d)", e)
	}
}

type KConnection uint

const (
	KConnectionBlueGammaScale KConnection = 0
	KConnectionControllerColorDepth KConnection = 0
)

func (e KConnection) String() string {
	switch e {
	case KConnectionBlueGammaScale:
		return "KConnectionBlueGammaScale"
	default:
		return fmt.Sprintf("KConnection(%d)", e)
	}
}

type KConvolved uint

const (
	KConvolvedValue KConvolved = 0
	KConvolvedMask KConvolved = 0
)

func (e KConvolved) String() string {
	switch e {
	case KConvolvedValue:
		return "KConvolvedValue"
	default:
		return fmt.Sprintf("KConvolved(%d)", e)
	}
}

type KCoprocessor uint

const (
	KCoprocessorVersion1 KCoprocessor = 0
	KCoprocessorVersion2 KCoprocessor = 0
)

func (e KCoprocessor) String() string {
	switch e {
	case KCoprocessorVersion1:
		return "KCoprocessorVersion1"
	default:
		return fmt.Sprintf("KCoprocessor(%d)", e)
	}
}

type KDCLInvalidOp uint

const (
	KDCLCallProcOp KDCLInvalidOp = 0
)

func (e KDCLInvalidOp) String() string {
	switch e {
	case KDCLCallProcOp:
		return "KDCLCallProcOp"
	default:
		return fmt.Sprintf("KDCLInvalidOp(%d)", e)
	}
}

const KDDCBlockSize uint = 0

const KDDCBlockTypeEDID uint = 0

const KDDCForceReadBit uint = 0

type KDMSMode uint

const (
	KDMSModeFree KDMSMode = 0
	KDMSModeNotReady KDMSMode = 0
	KDMSModeReady KDMSMode = 0
)

func (e KDMSMode) String() string {
	switch e {
	case KDMSModeFree:
		return "KDMSModeFree"
	default:
		return fmt.Sprintf("KDMSMode(%d)", e)
	}
}

type KDPMSSyncOn uint

const (
	KDPMSSyncOff KDPMSSyncOn = 0
)

func (e KDPMSSyncOn) String() string {
	switch e {
	case KDPMSSyncOff:
		return "KDPMSSyncOff"
	default:
		return fmt.Sprintf("KDPMSSyncOn(%d)", e)
	}
}

const KDTMaxEntryNameLength uint = 0

const KDTMaxPropertyNameLength uint = 0

const KDTPathNameSeparator uint = 0

type KDVDBookTypeHDR uint

const (
	KDVDBookTypeHDRAM KDVDBookTypeHDR = 0
	KDVDBookTypeHDROM KDVDBookTypeHDR = 0
)

func (e KDVDBookTypeHDR) String() string {
	switch e {
	case KDVDBookTypeHDRAM:
		return "KDVDBookTypeHDRAM"
	default:
		return fmt.Sprintf("KDVDBookTypeHDR(%d)", e)
	}
}

type KDVDCPRMRegion1 uint

const (
	KDVDCPRMRegion2 KDVDCPRMRegion1 = 0
)

func (e KDVDCPRMRegion1) String() string {
	switch e {
	case KDVDCPRMRegion2:
		return "KDVDCPRMRegion2"
	default:
		return fmt.Sprintf("KDVDCPRMRegion1(%d)", e)
	}
}

type KDVDFeatures uint

const (
	KDVDFeaturesBUFWriteBit KDVDFeatures = 0
	KDVDFeaturesBUFWriteMask KDVDFeatures = 0
	KDVDFeaturesHDRBit KDVDFeatures = 0
	KDVDFeaturesHDRMask KDVDFeatures = 0
)

func (e KDVDFeatures) String() string {
	switch e {
	case KDVDFeaturesBUFWriteBit:
		return "KDVDFeaturesBUFWriteBit"
	default:
		return fmt.Sprintf("KDVDFeatures(%d)", e)
	}
}

type KDVDKeyClass uint

const (
	KDVDKeyClassCSS_CPPM_CPRM KDVDKeyClass = 0
	KDVDKeyClassRSSA KDVDKeyClass = 0
)

func (e KDVDKeyClass) String() string {
	switch e {
	case KDVDKeyClassCSS_CPPM_CPRM:
		return "KDVDKeyClassCSS_CPPM_CPRM"
	default:
		return fmt.Sprintf("KDVDKeyClass(%d)", e)
	}
}

type KDVDKeyFormatAGID uint

const (
	KDVDKeyFormatAGID_CPRM KDVDKeyFormatAGID = 0
	KDVDKeyFormatAGID_Invalidate KDVDKeyFormatAGID = 0
)

func (e KDVDKeyFormatAGID) String() string {
	switch e {
	case KDVDKeyFormatAGID_CPRM:
		return "KDVDKeyFormatAGID_CPRM"
	default:
		return fmt.Sprintf("KDVDKeyFormatAGID(%d)", e)
	}
}

type KDVDMediaType uint

const (
	KDVDMediaTypeMax KDVDMediaType = 0
	KDVDMediaTypeRW KDVDMediaType = 0
)

func (e KDVDMediaType) String() string {
	switch e {
	case KDVDMediaTypeMax:
		return "KDVDMediaTypeMax"
	default:
		return fmt.Sprintf("KDVDMediaType(%d)", e)
	}
}

type KDVDRZoneInfoAddressType uint

const (
	KDVDRZoneInfoAddressTypeBorderNumber KDVDRZoneInfoAddressType = 0
	KDVDRZoneInfoAddressTypeLBA KDVDRZoneInfoAddressType = 0
)

func (e KDVDRZoneInfoAddressType) String() string {
	switch e {
	case KDVDRZoneInfoAddressTypeBorderNumber:
		return "KDVDRZoneInfoAddressTypeBorderNumber"
	default:
		return fmt.Sprintf("KDVDRZoneInfoAddressType(%d)", e)
	}
}

const KDVDRegionalPlaybackControlSchemePhase1 uint = 0

type KDVDStructureFormatPhysicalFormatInfo uint

const (
	KDVDStructureFormatCopyrightInfo KDVDStructureFormatPhysicalFormatInfo = 0
)

func (e KDVDStructureFormatPhysicalFormatInfo) String() string {
	switch e {
	case KDVDStructureFormatCopyrightInfo:
		return "KDVDStructureFormatCopyrightInfo"
	default:
		return fmt.Sprintf("KDVDStructureFormatPhysicalFormatInfo(%d)", e)
	}
}

type KDVIPowerSwitchFeature uint

const (
	KDVIPowerSwitchActiveMask KDVIPowerSwitchFeature = 0
)

func (e KDVIPowerSwitchFeature) String() string {
	switch e {
	case KDVIPowerSwitchActiveMask:
		return "KDVIPowerSwitchActiveMask"
	default:
		return fmt.Sprintf("KDVIPowerSwitchFeature(%d)", e)
	}
}

const KDataRotate0 uint = 0

type KDebugTypeNone uint

const (
	KDebugTypeDisplay KDebugTypeNone = 0
)

func (e KDebugTypeNone) String() string {
	switch e {
	case KDebugTypeDisplay:
		return "KDebugTypeDisplay"
	default:
		return fmt.Sprintf("KDebugTypeNone(%d)", e)
	}
}

type KDeclROMtables uint

const (
	KDeclROMtablesValue KDeclROMtables = 0
	KDetailedTimingFormat KDeclROMtables = 0
)

func (e KDeclROMtables) String() string {
	switch e {
	case KDeclROMtablesValue:
		return "KDeclROMtablesValue"
	default:
		return fmt.Sprintf("KDeclROMtables(%d)", e)
	}
}

type KDepth uint

const (
	KDepthMode1 KDepth = 0
	KDepthMode4 KDepth = 0
)

func (e KDepth) String() string {
	switch e {
	case KDepthMode1:
		return "KDepthMode1"
	default:
		return fmt.Sprintf("KDepth(%d)", e)
	}
}

const KDepthDependent uint = 0

type KDigitalSignalBit uint

const (
	KAnalogSetupExpectedBit KDigitalSignalBit = 0
	KDigitalSignalBitValue KDigitalSignalBit = 0
)

func (e KDigitalSignalBit) String() string {
	switch e {
	case KAnalogSetupExpectedBit:
		return "KAnalogSetupExpectedBit"
	default:
		return fmt.Sprintf("KDigitalSignalBit(%d)", e)
	}
}

type KDisable uint

const (
	KDisableHorizontalSyncBit KDisable = 0
	KDisableVerticalSyncBit KDisable = 0
)

func (e KDisable) String() string {
	switch e {
	case KDisableHorizontalSyncBit:
		return "KDisableHorizontalSyncBit"
	default:
		return fmt.Sprintf("KDisable(%d)", e)
	}
}

const KDisabledInterruptState uint = 0

type KDiscStatus int

const (
	KDiscStatusComplete KDiscStatus = 0
	KDiscStatusEmpty KDiscStatus = 0
	KDiscStatusErasableMask KDiscStatus = 0
	KDiscStatusIncomplete KDiscStatus = 0
	KDiscStatusMask KDiscStatus = 0
	KDiscStatusOther KDiscStatus = 0
)

func (e KDiscStatus) String() string {
	switch e {
	case KDiscStatusComplete:
		return "KDiscStatusComplete"
	default:
		return fmt.Sprintf("KDiscStatus(%d)", e)
	}
}

type KDisplayMode uint

const (
	KDisplayModeAcceleratorBackedFlag KDisplayMode = 0
	KDisplayModeAlwaysShowFlag KDisplayMode = 0
)

func (e KDisplayMode) String() string {
	switch e {
	case KDisplayModeAcceleratorBackedFlag:
		return "KDisplayModeAcceleratorBackedFlag"
	default:
		return fmt.Sprintf("KDisplayMode(%d)", e)
	}
}

type KDisplayModeIDCurrent uint

const (
	KDisplayModeIDBootProgrammable KDisplayModeIDCurrent = 0
)

func (e KDisplayModeIDCurrent) String() string {
	switch e {
	case KDisplayModeIDBootProgrammable:
		return "KDisplayModeIDBootProgrammable"
	default:
		return fmt.Sprintf("KDisplayModeIDCurrent(%d)", e)
	}
}

type KDisplayModeValidFlag uint

const (
	KDisplayModeDefaultFlag KDisplayModeValidFlag = 0
)

func (e KDisplayModeValidFlag) String() string {
	switch e {
	case KDisplayModeDefaultFlag:
		return "KDisplayModeDefaultFlag"
	default:
		return fmt.Sprintf("KDisplayModeValidFlag(%d)", e)
	}
}

type KDisplaySubPixel uint

const (
	KDisplaySubPixelConfigurationDelta KDisplaySubPixel = 0
	KDisplaySubPixelShapeRound KDisplaySubPixel = 0
)

func (e KDisplaySubPixel) String() string {
	switch e {
	case KDisplaySubPixelConfigurationDelta:
		return "KDisplaySubPixelConfigurationDelta"
	default:
		return fmt.Sprintf("KDisplaySubPixel(%d)", e)
	}
}

type KDisplayVendorIDUnknown uint

const (
	KDisplayProductIDGeneric KDisplayVendorIDUnknown = 0
)

func (e KDisplayVendorIDUnknown) String() string {
	switch e {
	case KDisplayProductIDGeneric:
		return "KDisplayProductIDGeneric"
	default:
		return fmt.Sprintf("KDisplayVendorIDUnknown(%d)", e)
	}
}

type KDriver uint

const (
	KDriverIsForVirtualDevice KDriver = 0
	KDriverSupportDMSuspendAndResume KDriver = 0
)

func (e KDriver) String() string {
	switch e {
	case KDriverIsForVirtualDevice:
		return "KDriverIsForVirtualDevice"
	default:
		return fmt.Sprintf("KDriver(%d)", e)
	}
}

type KESC uint

const (
	KESCOnePortraitMono KESC = 0
	KESCSixStandard KESC = 0
)

func (e KESC) String() string {
	switch e {
	case KESCOnePortraitMono:
		return "KESCOnePortraitMono"
	default:
		return fmt.Sprintf("KESC(%d)", e)
	}
}

type KEfi uint

const (
	KEfiACPIReclaimMemory KEfi = 0
	KEfiBootServicesCode KEfi = 0
)

func (e KEfi) String() string {
	switch e {
	case KEfiACPIReclaimMemory:
		return "KEfiACPIReclaimMemory"
	default:
		return fmt.Sprintf("KEfi(%d)", e)
	}
}

type KEndpointDirection uint

const (
	KEndpointDirectionIn KEndpointDirection = 0
	KEndpointDirectionOut KEndpointDirection = 0
)

func (e KEndpointDirection) String() string {
	switch e {
	case KEndpointDirectionIn:
		return "KEndpointDirectionIn"
	default:
		return fmt.Sprintf("KEndpointDirection(%d)", e)
	}
}

type KEndpointType uint

const (
	KEndpointTypeBulk KEndpointType = 0
	KEndpointTypeIsochronous KEndpointType = 0
)

func (e KEndpointType) String() string {
	switch e {
	case KEndpointTypeBulk:
		return "KEndpointTypeBulk"
	default:
		return fmt.Sprintf("KEndpointType(%d)", e)
	}
}

const KError int = 0

type KExclave uint

const (
	KExclaveRPCActive KExclave = 0
	KExclaveSchedulerRequest KExclave = 0
)

func (e KExclave) String() string {
	switch e {
	case KExclaveRPCActive:
		return "KExclaveRPCActive"
	default:
		return fmt.Sprintf("KExclave(%d)", e)
	}
}

type KExclaveIpcStackEntryHave uint

const (
	KExclaveIpcStackEntryHaveInvocationID KExclaveIpcStackEntryHave = 0
	KExclaveIpcStackEntryHaveStack KExclaveIpcStackEntryHave = 0
)

func (e KExclaveIpcStackEntryHave) String() string {
	switch e {
	case KExclaveIpcStackEntryHaveInvocationID:
		return "KExclaveIpcStackEntryHaveInvocationID"
	default:
		return fmt.Sprintf("KExclaveIpcStackEntryHave(%d)", e)
	}
}

type KExclaveTextLayoutLoadAddresses uint

const (
	KExclaveTextLayoutLoadAddressesSynthetic KExclaveTextLayoutLoadAddresses = 0
	KExclaveTextLayoutLoadAddressesUnslid KExclaveTextLayoutLoadAddresses = 0
)

func (e KExclaveTextLayoutLoadAddresses) String() string {
	switch e {
	case KExclaveTextLayoutLoadAddressesSynthetic:
		return "KExclaveTextLayoutLoadAddressesSynthetic"
	default:
		return fmt.Sprintf("KExclaveTextLayoutLoadAddresses(%d)", e)
	}
}

type KFBDisplayUsablePowerState uint

const (
	KFBDisplayPowerStateMask KFBDisplayUsablePowerState = 0
)

func (e KFBDisplayUsablePowerState) String() string {
	switch e {
	case KFBDisplayPowerStateMask:
		return "KFBDisplayPowerStateMask"
	default:
		return fmt.Sprintf("KFBDisplayUsablePowerState(%d)", e)
	}
}

type KFFT uint

const (
	KFFTRadix2 KFFT = 0
	KFFTRadix3 KFFT = 0
	// KFFTRadix5: Specifies a radix of 5.
	KFFTRadix5 KFFT = 0
)

func (e KFFT) String() string {
	switch e {
	case KFFTRadix2:
		return "KFFTRadix2"
	default:
		return fmt.Sprintf("KFFT(%d)", e)
	}
}

type KFFTDirection uint

const (
	KFFTDirection_Forward KFFTDirection = 0
	// KFFTDirection_Inverse: Specifies an inverse transform.
	KFFTDirection_Inverse KFFTDirection = 0
)

func (e KFFTDirection) String() string {
	switch e {
	case KFFTDirection_Forward:
		return "KFFTDirection_Forward"
	default:
		return fmt.Sprintf("KFFTDirection(%d)", e)
	}
}

type KFW uint

const (
	KFWAddressBusID KFW = 0
	KFWAllowMultiMode KFW = 0
	KFWAlwaysMultiMode KFW = 0
	KFWBroadcastNodeID KFW = 0
	KFWDCLInvalidNotification KFW = 0
	KFWDCLModifyNotification KFW = 0
	KFWDCLUpdateNotification KFW = 0
	KFWDefaultIsochResourceFlags KFW = 0
	KFWNeverMultiMode KFW = 0
	KFWNuDCLModifyJumpNotification KFW = 0
	KFWNuDCLModifyNotification KFW = 0
	KFWNuDCLUpdateNotification KFW = 0
	KFWSpeed100MBit KFW = 0
	KFWSpeed200MBit KFW = 0
	KFWSpeed400MBit KFW = 0
	KFWSpeed800MBit KFW = 0
	KFWSpeedInvalid KFW = 0
	KFWSpeedMaximum KFW = 0
	KFWSpeedReserved KFW = 0
	KFWSpeedReserved1 KFW = 0
	KFWSpeedUnknownMask KFW = 0
	KFWSuggestMultiMode KFW = 0
)

func (e KFW) String() string {
	switch e {
	case KFWAddressBusID:
		return "KFWAddressBusID"
	default:
		return fmt.Sprintf("KFW(%d)", e)
	}
}

type KFWAVCAsync uint

const (
	KFWAVCAsyncPlug1 KFWAVCAsync = 0
	KFWAVCAsyncPlug11 KFWAVCAsync = 0
)

func (e KFWAVCAsync) String() string {
	switch e {
	case KFWAVCAsyncPlug1:
		return "KFWAVCAsyncPlug1"
	default:
		return fmt.Sprintf("KFWAVCAsync(%d)", e)
	}
}

type KFWAVCConsumerMode uint

const (
	KFWAVCConsumerMode_JUNK KFWAVCConsumerMode = 0
	KFWAVCConsumerMode_MORE KFWAVCConsumerMode = 0
)

func (e KFWAVCConsumerMode) String() string {
	switch e {
	case KFWAVCConsumerMode_JUNK:
		return "KFWAVCConsumerMode_JUNK"
	default:
		return fmt.Sprintf("KFWAVCConsumerMode(%d)", e)
	}
}

type KFWAVCProducerMode uint

const (
	KFWAVCProducerMode_SEND KFWAVCProducerMode = 0
	KFWAVCProducerMode_TOSS KFWAVCProducerMode = 0
)

func (e KFWAVCProducerMode) String() string {
	switch e {
	case KFWAVCProducerMode_SEND:
		return "KFWAVCProducerMode_SEND"
	default:
		return fmt.Sprintf("KFWAVCProducerMode(%d)", e)
	}
}

type KFWAVCState uint

const (
	KFWAVCStateBusResumed KFWAVCState = 0
	KFWAVCStateDeviceRemoved KFWAVCState = 0
)

func (e KFWAVCState) String() string {
	switch e {
	case KFWAVCStateBusResumed:
		return "KFWAVCStateBusResumed"
	default:
		return fmt.Sprintf("KFWAVCState(%d)", e)
	}
}

type KFWAckTimeout uint

const (
	KFWAckBusyA KFWAckTimeout = 0
)

func (e KFWAckTimeout) String() string {
	switch e {
	case KFWAckBusyA:
		return "KFWAckBusyA"
	default:
		return fmt.Sprintf("KFWAckTimeout(%d)", e)
	}
}

type KFWAsynch uint

const (
	KFWAsynchAckSent KFWAsynch = 0
	KFWAsynchPriority KFWAsynch = 0
)

func (e KFWAsynch) String() string {
	switch e {
	case KFWAsynchAckSent:
		return "KFWAsynchAckSent"
	default:
		return fmt.Sprintf("KFWAsynch(%d)", e)
	}
}

type KFWBIB uint

const (
	KFWBIBIrmc KFWBIB = 0
	KFWBIBMaxROMPhase KFWBIB = 0
)

func (e KFWBIB) String() string {
	switch e {
	case KFWBIBIrmc:
		return "KFWBIBIrmc"
	default:
		return fmt.Sprintf("KFWBIB(%d)", e)
	}
}

const KFWBusManagerArbitrationTimeoutDuration uint = 0

type KFWCSRState uint

const (
	KFWCSRStateCMstr KFWCSRState = 0
	KFWCSRStateGone KFWCSRState = 0
	KFWCSRStateLinkOff KFWCSRState = 0
)

func (e KFWCSRState) String() string {
	switch e {
	case KFWCSRStateCMstr:
		return "KFWCSRStateCMstr"
	default:
		return fmt.Sprintf("KFWCSRState(%d)", e)
	}
}

const KFWConfigurationPacketID uint = 0

type KFWDCL uint

const (
	KFWDCLCycleEvent KFWDCL = 0
	KFWDCLSyBitsEvent KFWDCL = 0
)

func (e KFWDCL) String() string {
	switch e {
	case KFWDCLCycleEvent:
		return "KFWDCLCycleEvent"
	default:
		return fmt.Sprintf("KFWDCL(%d)", e)
	}
}

type KFWDCLOp uint

const (
	KFWDCLOpDynamicFlag KFWDCLOp = 0
	KFWDCLOpFlagMask KFWDCLOp = 0
)

func (e KFWDCLOp) String() string {
	switch e {
	case KFWDCLOpDynamicFlag:
		return "KFWDCLOpDynamicFlag"
	default:
		return fmt.Sprintf("KFWDCLOp(%d)", e)
	}
}

const KFWDebugIgnoreNodeNone uint = 0

type KFWExtendedTCodeMaskSwap uint

const (
	KFWExtendedTCodeBoundedAdd KFWExtendedTCodeMaskSwap = 0
)

func (e KFWExtendedTCodeMaskSwap) String() string {
	switch e {
	case KFWExtendedTCodeBoundedAdd:
		return "KFWExtendedTCodeBoundedAdd"
	default:
		return fmt.Sprintf("KFWExtendedTCodeMaskSwap(%d)", e)
	}
}

type KFWIsoch uint

const (
	KFWIsochBigEndianUpdates KFWIsoch = 0
	KFWIsochChanNum KFWIsoch = 0
	KFWIsochEnableRobustness KFWIsoch = 0
	KFWIsochPortDefaultOptions KFWIsoch = 0
	KFWIsochPortUseSeparateKernelThread KFWIsoch = 0
	KFWIsochRequireLastContext KFWIsoch = 0
	KFWIsochSy KFWIsoch = 0
)

func (e KFWIsoch) String() string {
	switch e {
	case KFWIsochBigEndianUpdates:
		return "KFWIsochBigEndianUpdates"
	default:
		return fmt.Sprintf("KFWIsoch(%d)", e)
	}
}

const KFWIsochChannelDefaultFlags uint = 0

type KFWIsochChannelUnknownCondition uint

const (
	KFWIsochChannelChannelNotAvailable KFWIsochChannelUnknownCondition = 0
)

func (e KFWIsochChannelUnknownCondition) String() string {
	switch e {
	case KFWIsochChannelChannelNotAvailable:
		return "KFWIsochChannelChannelNotAvailable"
	default:
		return fmt.Sprintf("KFWIsochChannelUnknownCondition(%d)", e)
	}
}

const KFWMaxBusses uint = 0

const KFWPacketTCode uint = 0

type KFWPhyConfigurationR uint

const (
	KFWPhyConfigurationGapCnt KFWPhyConfigurationR = 0
)

func (e KFWPhyConfigurationR) String() string {
	switch e {
	case KFWPhyConfigurationGapCnt:
		return "KFWPhyConfigurationGapCnt"
	default:
		return fmt.Sprintf("KFWPhyConfigurationR(%d)", e)
	}
}

type KFWPhyPacket uint

const (
	KFWPhyPacketID KFWPhyPacket = 0
	KFWPhyPacketIDPhase KFWPhyPacket = 0
	KFWPhyPacketPhyID KFWPhyPacket = 0
	KFWPhyPacketPhyIDPhase KFWPhyPacket = 0
)

func (e KFWPhyPacket) String() string {
	switch e {
	case KFWPhyPacketID:
		return "KFWPhyPacketID"
	default:
		return fmt.Sprintf("KFWPhyPacket(%d)", e)
	}
}

type KFWResponse uint

const (
	KFWResponseAddressError KFWResponse = 0
	KFWResponseTypeError KFWResponse = 0
)

func (e KFWResponse) String() string {
	switch e {
	case KFWResponseAddressError:
		return "KFWResponseAddressError"
	default:
		return fmt.Sprintf("KFWResponse(%d)", e)
	}
}

type KFWSBP2Command uint

const (
	KFWSBP2CommandCheckGeneration KFWSBP2Command = 0
	KFWSBP2CommandCompleteNotify KFWSBP2Command = 0
	KFWSBP2CommandDummyORB KFWSBP2Command = 0
	KFWSBP2CommandFixedSize KFWSBP2Command = 0
	KFWSBP2CommandImmediate KFWSBP2Command = 0
	KFWSBP2CommandNormalORB KFWSBP2Command = 0
	KFWSBP2CommandReservedORB KFWSBP2Command = 0
	KFWSBP2CommandTransferDataFromTarget KFWSBP2Command = 0
	KFWSBP2CommandVendorORB KFWSBP2Command = 0
	KFWSBP2CommandVirtualORBs KFWSBP2Command = 0
)

func (e KFWSBP2Command) String() string {
	switch e {
	case KFWSBP2CommandCheckGeneration:
		return "KFWSBP2CommandCheckGeneration"
	default:
		return fmt.Sprintf("KFWSBP2Command(%d)", e)
	}
}

const KFWSBP2ConstraintForceDoubleBuffer uint = 0

const KFWSBP2DontSynchronizeMgmtAgent uint = 0

const KFWSBP2MaxPageClusterSize uint = 0

type KFWSBP2NormalCommandStatus int

const (
	KFWSBP2NormalCommandReset KFWSBP2NormalCommandStatus = 0
)

func (e KFWSBP2NormalCommandStatus) String() string {
	switch e {
	case KFWSBP2NormalCommandReset:
		return "KFWSBP2NormalCommandReset"
	default:
		return fmt.Sprintf("KFWSBP2NormalCommandStatus(%d)", e)
	}
}

type KFWSBP2QueryLogins uint

const (
	KFWSBP2AbortTask KFWSBP2QueryLogins = 0
)

func (e KFWSBP2QueryLogins) String() string {
	switch e {
	case KFWSBP2AbortTask:
		return "KFWSBP2AbortTask"
	default:
		return fmt.Sprintf("KFWSBP2QueryLogins(%d)", e)
	}
}

type KFWSelfIDN uint

const (
	KFWSelfIDNPb KFWSelfIDN = 0
	KFWSelfIDNPc KFWSelfIDN = 0
)

func (e KFWSelfIDN) String() string {
	switch e {
	case KFWSelfIDNPb:
		return "KFWSelfIDNPb"
	default:
		return fmt.Sprintf("KFWSelfIDN(%d)", e)
	}
}

type KFWTCode uint

const (
	KFWTCodeCycleStart KFWTCode = 0
	KFWTCodeIsochronousBlock KFWTCode = 0
)

func (e KFWTCode) String() string {
	switch e {
	case KFWTCodeCycleStart:
		return "KFWTCodeCycleStart"
	default:
		return fmt.Sprintf("KFWTCode(%d)", e)
	}
}

type KFirstDepthMode uint

const (
	KFifthDepthMode KFirstDepthMode = 0
	KSecondDepthMode KFirstDepthMode = 0
)

func (e KFirstDepthMode) String() string {
	switch e {
	case KFifthDepthMode:
		return "KFifthDepthMode"
	default:
		return fmt.Sprintf("KFirstDepthMode(%d)", e)
	}
}

type KFirstIOKitNotificationType uint

const (
	KFirstIOKitNotificationTypeValue KFirstIOKitNotificationType = 0
	KLastIOKitNotificationType KFirstIOKitNotificationType = 0
)

func (e KFirstIOKitNotificationType) String() string {
	switch e {
	case KFirstIOKitNotificationTypeValue:
		return "KFirstIOKitNotificationTypeValue"
	default:
		return fmt.Sprintf("KFirstIOKitNotificationType(%d)", e)
	}
}

type KFixedDeviceBit uint

const (
	KDMABit KFixedDeviceBit = 0
	KFixedDeviceBitValue KFixedDeviceBit = 0
)

func (e KFixedDeviceBit) String() string {
	switch e {
	case KDMABit:
		return "KDMABit"
	default:
		return fmt.Sprintf("KFixedDeviceBit(%d)", e)
	}
}

type KFramebuffer uint

const (
	KFramebufferDisableAltivecAccess KFramebuffer = 0
	KFramebufferSupportsGammaCorrection KFramebuffer = 0
)

func (e KFramebuffer) String() string {
	switch e {
	case KFramebufferDisableAltivecAccess:
		return "KFramebufferDisableAltivecAccess"
	default:
		return fmt.Sprintf("KFramebuffer(%d)", e)
	}
}

const KGammaTableIDFindFirst uint = 0

type KGetConnectionCount uint

const (
	KActivateConnection KGetConnectionCount = 0
)

func (e KGetConnectionCount) String() string {
	switch e {
	case KActivateConnection:
		return "KActivateConnection"
	default:
		return fmt.Sprintf("KGetConnectionCount(%d)", e)
	}
}

type KHFS uint

const (
	KHFSAutoCandidateMask KHFS = 0
	KHFSBinaryCompare KHFS = 0
	KHFSCaseFolding KHFS = 0
	KHFSCatalogKeyMaximumLength KHFS = 0
	KHFSExtentDensity KHFS = 0
	KHFSFileRecord KHFS = 0
	KHFSFileThreadRecord KHFS = 0
	KHFSFolderRecord KHFS = 0
	KHFSFolderThreadRecord KHFS = 0
	KHFSHasAttributesMask KHFS = 0
	KHFSPlusAttrMinNodeSize KHFS = 0
	KHFSPlusExtentDensity KHFS = 0
	KHFSPlusFileRecord KHFS = 0
	KHFSPlusFileThreadRecord KHFS = 0
	KHFSPlusFolderRecord KHFS = 0
	KHFSPlusFolderThreadRecord KHFS = 0
	KHFSUnusedNodeFixMask KHFS = 0
	KHFSVolumeJournaledMask KHFS = 0
)

func (e KHFS) String() string {
	switch e {
	case KHFSAutoCandidateMask:
		return "KHFSAutoCandidateMask"
	default:
		return fmt.Sprintf("KHFS(%d)", e)
	}
}

type KHFSMax uint

const (
	KHFSMaxFileNameChars KHFSMax = 0
	KHFSMaxVolumeNameChars KHFSMax = 0
)

func (e KHFSMax) String() string {
	switch e {
	case KHFSMaxFileNameChars:
		return "KHFSMaxFileNameChars"
	default:
		return fmt.Sprintf("KHFSMax(%d)", e)
	}
}

const KHFSMaxAttrNameLen uint = 0

type KHFSPlusAttr uint

const (
	KHFSPlusAttrExtents KHFSPlusAttr = 0
	KHFSPlusAttrForkData KHFSPlusAttr = 0
)

func (e KHFSPlusAttr) String() string {
	switch e {
	case KHFSPlusAttrExtents:
		return "KHFSPlusAttrExtents"
	default:
		return fmt.Sprintf("KHFSPlusAttr(%d)", e)
	}
}

type KHFSRootParentID uint

const (
	KHFSAllocationFileID KHFSRootParentID = 0
)

func (e KHFSRootParentID) String() string {
	switch e {
	case KHFSAllocationFileID:
		return "KHFSAllocationFileID"
	default:
		return fmt.Sprintf("KHFSRootParentID(%d)", e)
	}
}

type KHFSSigWord uint

const (
	KFSKMountVersion KHFSSigWord = 0
	KHFSSigWordValue KHFSSigWord = 0
)

func (e KHFSSigWord) String() string {
	switch e {
	case KFSKMountVersion:
		return "KFSKMountVersion"
	default:
		return fmt.Sprintf("KHFSSigWord(%d)", e)
	}
}

const KHFSUnusedNodesFixDate uint = 0

type KHID uint

const (
	KHIDBootProtocolValue KHID = 0
	KHIDReportProtocolValue KHID = 0
)

func (e KHID) String() string {
	switch e {
	case KHIDBootProtocolValue:
		return "KHIDBootProtocolValue"
	default:
		return fmt.Sprintf("KHID(%d)", e)
	}
}

type KHIDDispatchOption uint

const (
	KHIDDispatchOptionDeliveryNotificationForce KHIDDispatchOption = 0
	KHIDDispatchOptionPhaseCanceled KHIDDispatchOption = 0
)

func (e KHIDDispatchOption) String() string {
	switch e {
	case KHIDDispatchOptionDeliveryNotificationForce:
		return "KHIDDispatchOptionDeliveryNotificationForce"
	default:
		return fmt.Sprintf("KHIDDispatchOption(%d)", e)
	}
}

const KHIDDispatchOptionKeyboardNoRepeat uint = 0

type KHIDDispatchOptionPointerNoAcceleration uint

const (
	KHIDDispatchOptionPointerAbsolutToRelative KHIDDispatchOptionPointerNoAcceleration = 0
)

func (e KHIDDispatchOptionPointerNoAcceleration) String() string {
	switch e {
	case KHIDDispatchOptionPointerAbsolutToRelative:
		return "KHIDDispatchOptionPointerAbsolutToRelative"
	default:
		return fmt.Sprintf("KHIDDispatchOptionPointerNoAcceleration(%d)", e)
	}
}

type KHIDPage uint

const (
	KHIDPage_Arcade KHIDPage = 0
	KHIDPage_Digitizer KHIDPage = 0
)

func (e KHIDPage) String() string {
	switch e {
	case KHIDPage_Arcade:
		return "KHIDPage_Arcade"
	default:
		return fmt.Sprintf("KHIDPage(%d)", e)
	}
}

type KHIDRqGetReport uint

const (
	KHIDRqGetIdle KHIDRqGetReport = 0
)

func (e KHIDRqGetReport) String() string {
	switch e {
	case KHIDRqGetIdle:
		return "KHIDRqGetIdle"
	default:
		return fmt.Sprintf("KHIDRqGetReport(%d)", e)
	}
}

type KHIDRt uint

const (
	KHIDRtFeatureReport KHIDRt = 0
	KHIDRtInputReport KHIDRt = 0
)

func (e KHIDRt) String() string {
	switch e {
	case KHIDRtFeatureReport:
		return "KHIDRtFeatureReport"
	default:
		return fmt.Sprintf("KHIDRt(%d)", e)
	}
}

type KHIDUsage uint

const (
	KHIDUsage_Keyboard3 KHIDUsage = 0
	KHIDUsage_KeyboardV KHIDUsage = 0
)

func (e KHIDUsage) String() string {
	switch e {
	case KHIDUsage_Keyboard3:
		return "KHIDUsage_Keyboard3"
	default:
		return fmt.Sprintf("KHIDUsage(%d)", e)
	}
}

type KHardLinkFileType uint

const (
	KHFSPlusCreator KHardLinkFileType = 0
	KHardLinkFileTypeValue KHardLinkFileType = 0
)

func (e KHardLinkFileType) String() string {
	switch e {
	case KHFSPlusCreator:
		return "KHFSPlusCreator"
	default:
		return fmt.Sprintf("KHardLinkFileType(%d)", e)
	}
}

const KHardwareCursorDescriptorMajorVersion uint = 0

type KHardwareCursorInfo uint

const (
	KHardwareCursorInfoMajorVersion KHardwareCursorInfo = 0
	KHardwareCursorInfoMinorVersion KHardwareCursorInfo = 0
)

func (e KHardwareCursorInfo) String() string {
	switch e {
	case KHardwareCursorInfoMajorVersion:
		return "KHardwareCursorInfoMajorVersion"
	default:
		return fmt.Sprintf("KHardwareCursorInfo(%d)", e)
	}
}

const KHubLocalPowerStatus int = 0

type KHubPortIndicatorAutomatic uint

const (
	KHubPortIndicatorAmber KHubPortIndicatorAutomatic = 0
)

func (e KHubPortIndicatorAutomatic) String() string {
	switch e {
	case KHubPortIndicatorAmber:
		return "KHubPortIndicatorAmber"
	default:
		return fmt.Sprintf("KHubPortIndicatorAutomatic(%d)", e)
	}
}

type KHubSuperSpeedProtocol uint

const (
	KUSBInterfaceAssociationDescriptorProtocol KHubSuperSpeedProtocol = 0
)

func (e KHubSuperSpeedProtocol) String() string {
	switch e {
	case KUSBInterfaceAssociationDescriptorProtocol:
		return "KUSBInterfaceAssociationDescriptorProtocol"
	default:
		return fmt.Sprintf("KHubSuperSpeedProtocol(%d)", e)
	}
}

type KHubSupportsGangPower uint

const (
	KHubPortSetPowerOff KHubSupportsGangPower = 0
)

func (e KHubSupportsGangPower) String() string {
	switch e {
	case KHubPortSetPowerOff:
		return "KHubPortSetPowerOff"
	default:
		return fmt.Sprintf("KHubSupportsGangPower(%d)", e)
	}
}

type KI uint

const (
	KIO128RGBAFloatPixelFormat KI = 0
	KIO16BE555PixelFormat KI = 0
	KIO16BE565PixelFormat KI = 0
	KIO24BGRPixelFormat KI = 0
	KIO2IndexedPixelFormat KI = 0
	KIO64BGRAPixelFormat KI = 0
)

func (e KI) String() string {
	switch e {
	case KIO128RGBAFloatPixelFormat:
		return "KIO128RGBAFloatPixelFormat"
	default:
		return fmt.Sprintf("KI(%d)", e)
	}
}

type KINQUIRY uint

const (
	// KINQUIRY_MaximumDataSize: # Discussion
	KINQUIRY_MaximumDataSize KINQUIRY = 0
	// KINQUIRY_Page00_PageCode: # Discussion
	KINQUIRY_Page00_PageCode KINQUIRY = 0
	// KINQUIRY_Page80_PageCode: # Discussion
	KINQUIRY_Page80_PageCode KINQUIRY = 0
	// KINQUIRY_Page83_PageCode: # Discussion
	KINQUIRY_Page83_PageCode KINQUIRY = 0
	// KINQUIRY_Page89_PageCode: # Discussion
	KINQUIRY_Page89_PageCode KINQUIRY = 0
	KINQUIRY_PageB0_PageCode KINQUIRY = 0
	KINQUIRY_PageB1_PageCode KINQUIRY = 0
	KINQUIRY_PageB2_PageCode KINQUIRY = 0
	KINQUIRY_PageC0_PageCode KINQUIRY = 0
	KINQUIRY_PageC1_PageCode KINQUIRY = 0
	// KINQUIRY_StandardDataHeaderSize: # Discussion
	KINQUIRY_StandardDataHeaderSize KINQUIRY = 0
)

func (e KINQUIRY) String() string {
	switch e {
	case KINQUIRY_MaximumDataSize:
		return "KINQUIRY_MaximumDataSize"
	default:
		return fmt.Sprintf("KINQUIRY(%d)", e)
	}
}

type KIO uint

const (
	KIOAnalogSetupExpected KIO = 0
	KIOBufferDescriptorMemoryFlags KIO = 0
	KIOCLUTPixels KIO = 0
	KIOCapturedAttribute KIO = 0
	KIODisplayAssertionConnectType KIO = 0
	KIOFBDisplayStateValue KIO = 0
	KIOFBServerConnectType KIO = 0
	KIOFBSharedConnectType KIO = 0
	KIOGDiagnoseConnectType KIO = 0
	KIOGDiagnoseGTraceType KIO = 0
	KIOInterlacedCEATiming KIO = 0
	KIOKextSpinDump KIO = 0
	KIOLogDTree KIO = 0
	KIOMemoryPageable KIO = 0
	KIOMonoDirectPixels KIO = 0
)

func (e KIO) String() string {
	switch e {
	case KIOAnalogSetupExpected:
		return "KIOAnalogSetupExpected"
	default:
		return fmt.Sprintf("KIO(%d)", e)
	}
}

type KIOACPIAddressSpaceIDSystem uint

const (
	KIOACPIAddressSpaceIDSystemIO KIOACPIAddressSpaceIDSystem = 0
	KIOACPIAddressSpaceIDSystemMemory KIOACPIAddressSpaceIDSystem = 0
)

func (e KIOACPIAddressSpaceIDSystem) String() string {
	switch e {
	case KIOACPIAddressSpaceIDSystemIO:
		return "KIOACPIAddressSpaceIDSystemIO"
	default:
		return fmt.Sprintf("KIOACPIAddressSpaceIDSystem(%d)", e)
	}
}

const KIOACPIAddressSpaceOpRead uint = 0

type KIOACPIDevicePowerStateD0 uint

const (
	KIOACPIDevicePowerStateCount KIOACPIDevicePowerStateD0 = 0
)

func (e KIOACPIDevicePowerStateD0) String() string {
	switch e {
	case KIOACPIDevicePowerStateCount:
		return "KIOACPIDevicePowerStateCount"
	default:
		return fmt.Sprintf("KIOACPIDevicePowerStateD0(%d)", e)
	}
}

type KIOACPIFixedEvent uint

const (
	KIOACPIFixedEventPMTimer KIOACPIFixedEvent = 0
	KIOACPIFixedEventPowerButton KIOACPIFixedEvent = 0
)

func (e KIOACPIFixedEvent) String() string {
	switch e {
	case KIOACPIFixedEventPMTimer:
		return "KIOACPIFixedEventPMTimer"
	default:
		return fmt.Sprintf("KIOACPIFixedEvent(%d)", e)
	}
}

type KIOACPIMemoryRange uint

const (
	KIOACPIBusNumberRange KIOACPIMemoryRange = 0
)

func (e KIOACPIMemoryRange) String() string {
	switch e {
	case KIOACPIBusNumberRange:
		return "KIOACPIBusNumberRange"
	default:
		return fmt.Sprintf("KIOACPIMemoryRange(%d)", e)
	}
}

type KIOAG uint

const (
	KIOAGP4GbAddressing KIOAG = 0
	KIOAGP4xDataRate KIOAG = 0
)

func (e KIOAG) String() string {
	switch e {
	case KIOAGP4GbAddressing:
		return "KIOAGP4GbAddressing"
	default:
		return fmt.Sprintf("KIOAG(%d)", e)
	}
}

type KIOAGP uint

const (
	KIOAGPIdle KIOAGP = 0
	KIOAGPInvalidGARTEntry KIOAGP = 0
)

func (e KIOAGP) String() string {
	switch e {
	case KIOAGPIdle:
		return "KIOAGPIdle"
	default:
		return fmt.Sprintf("KIOAGP(%d)", e)
	}
}

const KIOAGPDefaultStatus int = 0

type KIOAGPDisable uint

const (
	KIOAGPDisableAGPWrites KIOAGPDisable = 0
	KIOAGPDisableFeature6 KIOAGPDisable = 0
)

func (e KIOAGPDisable) String() string {
	switch e {
	case KIOAGPDisableAGPWrites:
		return "KIOAGPDisableAGPWrites"
	default:
		return fmt.Sprintf("KIOAGPDisable(%d)", e)
	}
}

const KIOAGPGartInvalidate uint = 0

type KIOAGPStateEnabled uint

const (
	KIOAGPStateEnablePending KIOAGPStateEnabled = 0
)

func (e KIOAGPStateEnabled) String() string {
	switch e {
	case KIOAGPStateEnablePending:
		return "KIOAGPStateEnablePending"
	default:
		return fmt.Sprintf("KIOAGPStateEnabled(%d)", e)
	}
}

type KIOATAFeaturePowerManagement uint

const (
	KIOATAFeature48BitLBA KIOATAFeaturePowerManagement = 0
)

func (e KIOATAFeaturePowerManagement) String() string {
	switch e {
	case KIOATAFeature48BitLBA:
		return "KIOATAFeature48BitLBA"
	default:
		return fmt.Sprintf("KIOATAFeaturePowerManagement(%d)", e)
	}
}

type KIOATAMaxPerformance uint

const (
	KIOATADefaultPerformance KIOATAMaxPerformance = 0
)

func (e KIOATAMaxPerformance) String() string {
	switch e {
	case KIOATADefaultPerformance:
		return "KIOATADefaultPerformance"
	default:
		return fmt.Sprintf("KIOATAMaxPerformance(%d)", e)
	}
}

type KIOATAMaximumBlockCount8Bit uint

const (
	KIOATAMaxBlocksPerXfer KIOATAMaximumBlockCount8Bit = 0
)

func (e KIOATAMaximumBlockCount8Bit) String() string {
	switch e {
	case KIOATAMaxBlocksPerXfer:
		return "KIOATAMaxBlocksPerXfer"
	default:
		return fmt.Sprintf("KIOATAMaximumBlockCount8Bit(%d)", e)
	}
}

type KIOATASectorCount8Bit uint

const (
	KIOATASectorCount16Bit KIOATASectorCount8Bit = 0
)

func (e KIOATASectorCount8Bit) String() string {
	switch e {
	case KIOATASectorCount16Bit:
		return "KIOATASectorCount16Bit"
	default:
		return fmt.Sprintf("KIOATASectorCount8Bit(%d)", e)
	}
}

type KIOAccel uint

const (
	KIOAccelNumClientTypes KIOAccel = 0
	KIOAccelNumSurfaceMethods KIOAccel = 0
	KIOAccelSurface2ClientType KIOAccel = 0
	KIOAccelSurfaceClientType KIOAccel = 0
	KIOAccelSurfaceControl KIOAccel = 0
	KIOAccelSurfaceFlush KIOAccel = 0
	KIOAccelSurfaceRead KIOAccel = 0
	KIOAccelSurfaceSetShapeBacking KIOAccel = 0
	KIOAccelSurfaceWriteLockOptions KIOAccel = 0
)

func (e KIOAccel) String() string {
	switch e {
	case KIOAccelNumClientTypes:
		return "KIOAccelNumClientTypes"
	default:
		return fmt.Sprintf("KIOAccel(%d)", e)
	}
}

const KIOAccelPrivateID uint = 0

const KIOAccelVolatileSurface uint = 0

type KIOAsyncReservedIndex uint

const (
	KIOAsyncCalloutCount KIOAsyncReservedIndex = 0
)

func (e KIOAsyncReservedIndex) String() string {
	switch e {
	case KIOAsyncCalloutCount:
		return "KIOAsyncCalloutCount"
	default:
		return fmt.Sprintf("KIOAsyncReservedIndex(%d)", e)
	}
}

type KIOAudio uint

const (
	KIOAudioBuiltInSystemClockDomain KIOAudio = 0
	KIOAudioInputPortSubTypeCD KIOAudio = 0
	KIOAudioNewClockDomain KIOAudio = 0
	KIOAudioOutputPortSubTypeInternalSpeaker KIOAudio = 0
)

func (e KIOAudio) String() string {
	switch e {
	case KIOAudioBuiltInSystemClockDomain:
		return "KIOAudioBuiltInSystemClockDomain"
	default:
		return fmt.Sprintf("KIOAudio(%d)", e)
	}
}

type KIOAudioClockSelectorTypeInternal uint

const (
	KIOAudioClockSelectorTypeADAT9Pin KIOAudioClockSelectorTypeInternal = 0
)

func (e KIOAudioClockSelectorTypeInternal) String() string {
	switch e {
	case KIOAudioClockSelectorTypeADAT9Pin:
		return "KIOAudioClockSelectorTypeADAT9Pin"
	default:
		return fmt.Sprintf("KIOAudioClockSelectorTypeInternal(%d)", e)
	}
}

type KIOAudioControlChannelID uint

const (
	KIOAudioControlChannelIDAll KIOAudioControlChannelID = 0
	KIOAudioControlChannelIDDefaultLeftRear KIOAudioControlChannelID = 0
)

func (e KIOAudioControlChannelID) String() string {
	switch e {
	case KIOAudioControlChannelIDAll:
		return "KIOAudioControlChannelIDAll"
	default:
		return fmt.Sprintf("KIOAudioControlChannelID(%d)", e)
	}
}

type KIOAudioControlTypeLevel uint

const (
	KIOAudioControlTypeToggle KIOAudioControlTypeLevel = 0
)

func (e KIOAudioControlTypeLevel) String() string {
	switch e {
	case KIOAudioControlTypeToggle:
		return "KIOAudioControlTypeToggle"
	default:
		return fmt.Sprintf("KIOAudioControlTypeLevel(%d)", e)
	}
}

type KIOAudioControlUsage uint

const (
	KIOAudioControlUsageCoreAudioProperty KIOAudioControlUsage = 0
	KIOAudioControlUsagePassThru KIOAudioControlUsage = 0
)

func (e KIOAudioControlUsage) String() string {
	switch e {
	case KIOAudioControlUsageCoreAudioProperty:
		return "KIOAudioControlUsageCoreAudioProperty"
	default:
		return fmt.Sprintf("KIOAudioControlUsage(%d)", e)
	}
}

type KIOAudioDevice uint

const (
	// KIOAudioDeviceIdle: # Discussion
	KIOAudioDeviceIdle KIOAudioDevice = 0
	// KIOAudioDeviceSleep: # Discussion
	KIOAudioDeviceSleep KIOAudioDevice = 0
)

func (e KIOAudioDevice) String() string {
	switch e {
	case KIOAudioDeviceIdle:
		return "KIOAudioDeviceIdle"
	default:
		return fmt.Sprintf("KIOAudioDevice(%d)", e)
	}
}

type KIOAudioDeviceCanBeDefaultNothing uint

const (
	KIOAudioDeviceCanBeDefaultInput KIOAudioDeviceCanBeDefaultNothing = 0
)

func (e KIOAudioDeviceCanBeDefaultNothing) String() string {
	switch e {
	case KIOAudioDeviceCanBeDefaultInput:
		return "KIOAudioDeviceCanBeDefaultInput"
	default:
		return fmt.Sprintf("KIOAudioDeviceCanBeDefaultNothing(%d)", e)
	}
}

type KIOAudioDeviceTransportType uint

const (
	KIOAudioDeviceTransportTypeAVB KIOAudioDeviceTransportType = 0
	KIOAudioDeviceTransportTypeBluetooth KIOAudioDeviceTransportType = 0
)

func (e KIOAudioDeviceTransportType) String() string {
	switch e {
	case KIOAudioDeviceTransportTypeAVB:
		return "KIOAudioDeviceTransportTypeAVB"
	default:
		return fmt.Sprintf("KIOAudioDeviceTransportType(%d)", e)
	}
}

type KIOAudioEngine uint

const (
	KIOAudioEngineAllNotifications KIOAudioEngine = 0
	KIOAudioEngineChangeNotification KIOAudioEngine = 0
	KIOAudioEnginePausedNotification KIOAudioEngine = 0
	KIOAudioEngineResumedNotification KIOAudioEngine = 0
	KIOAudioEngineStartedNotification KIOAudioEngine = 0
	KIOAudioEngineStoppedNotification KIOAudioEngine = 0
	KIOAudioEngineStreamFormatChangeNotification KIOAudioEngine = 0
)

func (e KIOAudioEngine) String() string {
	switch e {
	case KIOAudioEngineAllNotifications:
		return "KIOAudioEngineAllNotifications"
	default:
		return fmt.Sprintf("KIOAudioEngine(%d)", e)
	}
}

const KIOAudioLevelControlNegativeInfinity uint = 0

type KIOAudioLevelControlSubTypeVolume uint

const (
	KIOAudioLevelControlSubTypeLFEVolume KIOAudioLevelControlSubTypeVolume = 0
)

func (e KIOAudioLevelControlSubTypeVolume) String() string {
	switch e {
	case KIOAudioLevelControlSubTypeLFEVolume:
		return "KIOAudioLevelControlSubTypeLFEVolume"
	default:
		return fmt.Sprintf("KIOAudioLevelControlSubTypeVolume(%d)", e)
	}
}

type KIOAudioPortType uint

const (
	KIOAudioPortTypeInput KIOAudioPortType = 0
	KIOAudioPortTypeMixer KIOAudioPortType = 0
)

func (e KIOAudioPortType) String() string {
	switch e {
	case KIOAudioPortTypeInput:
		return "KIOAudioPortTypeInput"
	default:
		return fmt.Sprintf("KIOAudioPortType(%d)", e)
	}
}

type KIOAudioSMPTETime uint

const (
	KIOAudioSMPTETimeType2398 KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType24 KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType25 KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType2997 KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType2997Drop KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType30 KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType30Drop KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType50 KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType5994 KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType5994Drop KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType60 KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType60Drop KIOAudioSMPTETime = 0
)

func (e KIOAudioSMPTETime) String() string {
	switch e {
	case KIOAudioSMPTETimeType2398:
		return "KIOAudioSMPTETimeType2398"
	default:
		return fmt.Sprintf("KIOAudioSMPTETime(%d)", e)
	}
}

type KIOAudioSMPTETimeValid uint

const (
	KIOAudioSMPTETimeRunning KIOAudioSMPTETimeValid = 0
)

func (e KIOAudioSMPTETimeValid) String() string {
	switch e {
	case KIOAudioSMPTETimeRunning:
		return "KIOAudioSMPTETimeRunning"
	default:
		return fmt.Sprintf("KIOAudioSMPTETimeValid(%d)", e)
	}
}

type KIOAudioSelectorControlSelectionValueNone uint

const (
	KIOAudioSelectorControlSelectionValueCD KIOAudioSelectorControlSelectionValueNone = 0
)

func (e KIOAudioSelectorControlSelectionValueNone) String() string {
	switch e {
	case KIOAudioSelectorControlSelectionValueCD:
		return "KIOAudioSelectorControlSelectionValueCD"
	default:
		return fmt.Sprintf("KIOAudioSelectorControlSelectionValueNone(%d)", e)
	}
}

type KIOAudioStreamAlignmentLowByte uint

const (
	KIOAudioStreamAlignmentHighByte KIOAudioStreamAlignmentLowByte = 0
)

func (e KIOAudioStreamAlignmentLowByte) String() string {
	switch e {
	case KIOAudioStreamAlignmentHighByte:
		return "KIOAudioStreamAlignmentHighByte"
	default:
		return fmt.Sprintf("KIOAudioStreamAlignmentLowByte(%d)", e)
	}
}

type KIOAudioStreamByteOrder uint

const (
	KIOAudioStreamByteOrderBigEndian KIOAudioStreamByteOrder = 0
	KIOAudioStreamByteOrderLittleEndian KIOAudioStreamByteOrder = 0
)

func (e KIOAudioStreamByteOrder) String() string {
	switch e {
	case KIOAudioStreamByteOrderBigEndian:
		return "KIOAudioStreamByteOrderBigEndian"
	default:
		return fmt.Sprintf("KIOAudioStreamByteOrder(%d)", e)
	}
}

type KIOAudioStreamNumericRepresentationSignedInt uint

const (
	KIOAudioStreamNumericRepresentationIEEE754Float KIOAudioStreamNumericRepresentationSignedInt = 0
)

func (e KIOAudioStreamNumericRepresentationSignedInt) String() string {
	switch e {
	case KIOAudioStreamNumericRepresentationIEEE754Float:
		return "KIOAudioStreamNumericRepresentationIEEE754Float"
	default:
		return fmt.Sprintf("KIOAudioStreamNumericRepresentationSignedInt(%d)", e)
	}
}

type KIOAudioStreamSampleFormatLinearPCM uint

const (
	KIOAudioStreamSampleFormat1937AC3 KIOAudioStreamSampleFormatLinearPCM = 0
)

func (e KIOAudioStreamSampleFormatLinearPCM) String() string {
	switch e {
	case KIOAudioStreamSampleFormat1937AC3:
		return "KIOAudioStreamSampleFormat1937AC3"
	default:
		return fmt.Sprintf("KIOAudioStreamSampleFormatLinearPCM(%d)", e)
	}
}

const KIOAudioTimeStampSampleHostTimeValid uint = 0

type KIOAudioTimeStampSampleTimeValid uint

const (
	KIOAudioTimeStampHostTimeValid KIOAudioTimeStampSampleTimeValid = 0
)

func (e KIOAudioTimeStampSampleTimeValid) String() string {
	switch e {
	case KIOAudioTimeStampHostTimeValid:
		return "KIOAudioTimeStampHostTimeValid"
	default:
		return fmt.Sprintf("KIOAudioTimeStampSampleTimeValid(%d)", e)
	}
}

type KIOBatteryInstalled uint

const (
	KIOBatteryCharge KIOBatteryInstalled = 0
)

func (e KIOBatteryInstalled) String() string {
	switch e {
	case KIOBatteryCharge:
		return "KIOBatteryCharge"
	default:
		return fmt.Sprintf("KIOBatteryInstalled(%d)", e)
	}
}

type KIOBitsPerColor uint

const (
	KIOBitsPerColorComponent10 KIOBitsPerColor = 0
	KIOBitsPerColorComponent12 KIOBitsPerColor = 0
)

func (e KIOBitsPerColor) String() string {
	switch e {
	case KIOBitsPerColorComponent10:
		return "KIOBitsPerColorComponent10"
	default:
		return fmt.Sprintf("KIOBitsPerColor(%d)", e)
	}
}

type KIOBlit uint

const (
	KIOBlitAllOptions KIOBlit = 0
	KIOBlitBeamSyncSpin KIOBlit = 0
	KIOBlitFlushWithSwap KIOBlit = 0
	KIOBlitWaitGlobal KIOBlit = 0
)

func (e KIOBlit) String() string {
	switch e {
	case KIOBlitAllOptions:
		return "KIOBlitAllOptions"
	default:
		return fmt.Sprintf("KIOBlit(%d)", e)
	}
}

const KIOBlitColorSpaceTypes uint = 0

const KIOBlitFramebufferDestination uint = 0

type KIOBlitHasCGSSurface uint

const (
	KIOBlitBeamSyncSwaps KIOBlitHasCGSSurface = 0
)

func (e KIOBlitHasCGSSurface) String() string {
	switch e {
	case KIOBlitBeamSyncSwaps:
		return "KIOBlitBeamSyncSwaps"
	default:
		return fmt.Sprintf("KIOBlitHasCGSSurface(%d)", e)
	}
}

const KIOBlitMemoryRequiresHostFlush uint = 0

type KIOBlitSourceDefault uint

const (
	KIOBlitSourceCGSSurface KIOBlitSourceDefault = 0
)

func (e KIOBlitSourceDefault) String() string {
	switch e {
	case KIOBlitSourceCGSSurface:
		return "KIOBlitSourceCGSSurface"
	default:
		return fmt.Sprintf("KIOBlitSourceDefault(%d)", e)
	}
}

type KIOBlitSynchronizeWaitBeamExit uint

const (
	KIOBlitSynchronizeFlushHostWrites KIOBlitSynchronizeWaitBeamExit = 0
)

func (e KIOBlitSynchronizeWaitBeamExit) String() string {
	switch e {
	case KIOBlitSynchronizeFlushHostWrites:
		return "KIOBlitSynchronizeFlushHostWrites"
	default:
		return fmt.Sprintf("KIOBlitSynchronizeWaitBeamExit(%d)", e)
	}
}

type KIOBlitType uint

const (
	KIOBlitTypeSourceKeyColorNotEqual KIOBlitType = 0
	KIOBlitTypeVerbMask KIOBlitType = 0
)

func (e KIOBlitType) String() string {
	switch e {
	case KIOBlitTypeSourceKeyColorNotEqual:
		return "KIOBlitTypeSourceKeyColorNotEqual"
	default:
		return fmt.Sprintf("KIOBlitType(%d)", e)
	}
}

const KIOBlitUnlockWithSwap uint = 0

type KIOCatalog uint

const (
	// KIOCatalogAddDrivers: # Discussion
	KIOCatalogAddDrivers KIOCatalog = 0
	// KIOCatalogModuleTerminate: # Discussion
	KIOCatalogModuleTerminate KIOCatalog = 0
	KIOCatalogRemoveKernelLinker__Removed KIOCatalog = 0
	// KIOCatalogServiceTerminate: # Discussion
	KIOCatalogServiceTerminate KIOCatalog = 0
)

func (e KIOCatalog) String() string {
	switch e {
	case KIOCatalogAddDrivers:
		return "KIOCatalogAddDrivers"
	default:
		return fmt.Sprintf("KIOCatalog(%d)", e)
	}
}

type KIOCatalogGet uint

const (
	KIOCatalogGetCacheMissList KIOCatalogGet = 0
	KIOCatalogGetModuleDemandList KIOCatalogGet = 0
)

func (e KIOCatalogGet) String() string {
	switch e {
	case KIOCatalogGetCacheMissList:
		return "KIOCatalogGetCacheMissList"
	default:
		return fmt.Sprintf("KIOCatalogGet(%d)", e)
	}
}

const KIOCatalogResetDefault uint = 0

type KIOColorimetryNotSupported uint

const (
	KIOColorimetryAdobeRGB KIOColorimetryNotSupported = 0
)

func (e KIOColorimetryNotSupported) String() string {
	switch e {
	case KIOColorimetryAdobeRGB:
		return "KIOColorimetryAdobeRGB"
	default:
		return fmt.Sprintf("KIOColorimetryNotSupported(%d)", e)
	}
}

const KIOConnectMethodVarOutputSize uint = 0

const KIOConnectionBuiltIn uint = 0

const KIODDCBlockTypeEDID uint = 0

const KIODDCForceRead uint = 0

type KIODDCLow uint

const (
	KIODDCHigh KIODDCLow = 0
)

func (e KIODDCLow) String() string {
	switch e {
	case KIODDCHigh:
		return "KIODDCHigh"
	default:
		return fmt.Sprintf("KIODDCLow(%d)", e)
	}
}

type KIODKEnable uint

const (
	KIODKDisableCDHashChecking KIODKEnable = 0
)

func (e KIODKEnable) String() string {
	switch e {
	case KIODKDisableCDHashChecking:
		return "KIODKDisableCDHashChecking"
	default:
		return fmt.Sprintf("KIODKEnable(%d)", e)
	}
}

const KIODMACommandCompleteDMANoOptions uint = 0

const KIODMACommandCreateNoOptions uint = 0

const KIODMACommandPerformOperationOptionRead uint = 0

const KIODMACommandPrepareForDMANoOptions uint = 0

const KIODMACommandSpecificationNoOptions uint = 0

type KIODMAMap uint

const (
	KIODMAMapDeviceMemory KIODMAMap = 0
	KIODMAMapIdentityMap KIODMAMap = 0
)

func (e KIODMAMap) String() string {
	switch e {
	case KIODMAMapDeviceMemory:
		return "KIODMAMapDeviceMemory"
	default:
		return fmt.Sprintf("KIODMAMap(%d)", e)
	}
}

type KIODMAMapOption uint

const (
	KIODMAMapOptionBypassed KIODMAMapOption = 0
	KIODMAMapOptionIterateOnly KIODMAMapOption = 0
)

func (e KIODMAMapOption) String() string {
	switch e {
	case KIODMAMapOptionBypassed:
		return "KIODMAMapOptionBypassed"
	default:
		return fmt.Sprintf("KIODMAMapOption(%d)", e)
	}
}

type KIODPEvent uint

const (
	KIODPEventAutomatedTestRequest KIODPEvent = 0
	KIODPEventContentProtection KIODPEvent = 0
)

func (e KIODPEvent) String() string {
	switch e {
	case KIODPEventAutomatedTestRequest:
		return "KIODPEventAutomatedTestRequest"
	default:
		return fmt.Sprintf("KIODPEvent(%d)", e)
	}
}

const KIODTInterruptShared uint = 0

type KIODTRecursive uint

const (
	KIODTExclusive KIODTRecursive = 0
)

func (e KIODTRecursive) String() string {
	switch e {
	case KIODTExclusive:
		return "KIODTExclusive"
	default:
		return fmt.Sprintf("KIODTRecursive(%d)", e)
	}
}

type KIODefaultCache uint

const (
	KIOCopybackCache KIODefaultCache = 0
)

func (e KIODefaultCache) String() string {
	switch e {
	case KIOCopybackCache:
		return "KIOCopybackCache"
	default:
		return fmt.Sprintf("KIODefaultCache(%d)", e)
	}
}

const KIODefaultMemoryType uint = 0

const KIODefaultProbeScore uint = 0

const KIODetailedTimingValid uint = 0

type KIODirection uint

const (
	KIODirectionCompleteWithDataValid KIODirection = 0
	KIODirectionCompleteWithError KIODirection = 0
	KIODirectionIn KIODirection = 0
	KIODirectionInOut KIODirection = 0
	KIODirectionNone KIODirection = 0
	KIODirectionOut KIODirection = 0
	KIODirectionOutIn KIODirection = 0
	KIODirectionPrepareNoFault KIODirection = 0
	KIODirectionPrepareNonCoherent KIODirection = 0
	KIODirectionPrepareReserved1 KIODirection = 0
	KIODirectionPrepareToPhys32 KIODirection = 0
)

func (e KIODirection) String() string {
	switch e {
	case KIODirectionCompleteWithDataValid:
		return "KIODirectionCompleteWithDataValid"
	default:
		return fmt.Sprintf("KIODirection(%d)", e)
	}
}

type KIODispatchQueueReentrant uint

const (
	KIODispatchQueueMethodsNotSynchronized KIODispatchQueueReentrant = 0
)

func (e KIODispatchQueueReentrant) String() string {
	switch e {
	case KIODispatchQueueMethodsNotSynchronized:
		return "KIODispatchQueueMethodsNotSynchronized"
	default:
		return fmt.Sprintf("KIODispatchQueueReentrant(%d)", e)
	}
}

const KIODispatchQueueSleepReuseEvent uint = 0

const KIODispatchQueueWakeupAll uint = 0

type KIODisplay uint

const (
	KIODisplayMaxPowerState KIODisplay = 0
	KIODisplayNumPowerStates KIODisplay = 0
)

func (e KIODisplay) String() string {
	switch e {
	case KIODisplayMaxPowerState:
		return "KIODisplayMaxPowerState"
	default:
		return fmt.Sprintf("KIODisplay(%d)", e)
	}
}

const KIODisplayColorMode uint = 0

type KIODisplayDitherDisable uint

const (
	KIODisplayDitherAll KIODisplayDitherDisable = 0
)

func (e KIODisplayDitherDisable) String() string {
	switch e {
	case KIODisplayDitherAll:
		return "KIODisplayDitherAll"
	default:
		return fmt.Sprintf("KIODisplayDitherDisable(%d)", e)
	}
}

type KIODisplayModeID uint

const (
	KIODisplayModeIDBootProgrammable KIODisplayModeID = 0
	KIODisplayModeIDReservedBase KIODisplayModeID = 0
)

func (e KIODisplayModeID) String() string {
	switch e {
	case KIODisplayModeIDBootProgrammable:
		return "KIODisplayModeIDBootProgrammable"
	default:
		return fmt.Sprintf("KIODisplayModeID(%d)", e)
	}
}

const KIODisplayNeedsCEAUnderscan uint = 0

type KIODisplayPowerStateOff uint

const (
	KIODisplayPowerStateMinUsable KIODisplayPowerStateOff = 0
)

func (e KIODisplayPowerStateOff) String() string {
	switch e {
	case KIODisplayPowerStateMinUsable:
		return "KIODisplayPowerStateMinUsable"
	default:
		return fmt.Sprintf("KIODisplayPowerStateOff(%d)", e)
	}
}

type KIODisplayRGBColorComponent uint

const (
	KIODisplayRGBColorComponentBits10 KIODisplayRGBColorComponent = 0
	KIODisplayRGBColorComponentBits14 KIODisplayRGBColorComponent = 0
)

func (e KIODisplayRGBColorComponent) String() string {
	switch e {
	case KIODisplayRGBColorComponentBits10:
		return "KIODisplayRGBColorComponentBits10"
	default:
		return fmt.Sprintf("KIODisplayRGBColorComponent(%d)", e)
	}
}

type KIODynamicRangeNotSupported uint

const (
	KIODynamicRangeDolbyNormalMode KIODynamicRangeNotSupported = 0
)

func (e KIODynamicRangeNotSupported) String() string {
	switch e {
	case KIODynamicRangeDolbyNormalMode:
		return "KIODynamicRangeDolbyNormalMode"
	default:
		return fmt.Sprintf("KIODynamicRangeNotSupported(%d)", e)
	}
}

type KIOEnetPromiscuousModeOff uint

const (
	KIOEnetMulticastModeFilter KIOEnetPromiscuousModeOff = 0
)

func (e KIOEnetPromiscuousModeOff) String() string {
	switch e {
	case KIOEnetMulticastModeFilter:
		return "KIOEnetMulticastModeFilter"
	default:
		return fmt.Sprintf("KIOEnetPromiscuousModeOff(%d)", e)
	}
}

type KIOEthernetControllerAVBStateEvent uint

const (
	KIOEthernetControllerAVBStateEventDisable KIOEthernetControllerAVBStateEvent = 0
	KIOEthernetControllerAVBStateEventEnable KIOEthernetControllerAVBStateEvent = 0
)

func (e KIOEthernetControllerAVBStateEvent) String() string {
	switch e {
	case KIOEthernetControllerAVBStateEventDisable:
		return "KIOEthernetControllerAVBStateEventDisable"
	default:
		return fmt.Sprintf("KIOEthernetControllerAVBStateEvent(%d)", e)
	}
}

const KIOEthernetWakeOnMagicPacket uint = 0

type KIOEventLinkAssociate uint

const (
	KIOEventLinkAssociateCurrentThread KIOEventLinkAssociate = 0
	KIOEventLinkAssociateOnWait KIOEventLinkAssociate = 0
)

func (e KIOEventLinkAssociate) String() string {
	switch e {
	case KIOEventLinkAssociateCurrentThread:
		return "KIOEventLinkAssociateCurrentThread"
	default:
		return fmt.Sprintf("KIOEventLinkAssociate(%d)", e)
	}
}

const KIOEventLinkClockMachAbsoluteTime uint = 0

const KIOEventLinkMaxNameLength uint = 0

const KIOExternalMethodArgumentsCurrentVersion uint = 0

const KIOExternalMethodScalarInputCountMax uint = 0

type KIOFB uint

const (
	KIOFBChangedInterruptType KIOFB = 0
	// KIOFBCurrentShmemVersion: # Discussion
	KIOFBCurrentShmemVersion KIOFB = 0
	KIOFBMainCursorIndex KIOFB = 0
	KIOFBNumCursorIndex KIOFB = 0
	KIOFBOnlineInterruptType KIOFB = 0
	KIOFBShmemVersionMask KIOFB = 0
)

func (e KIOFB) String() string {
	switch e {
	case KIOFBChangedInterruptType:
		return "KIOFBChangedInterruptType"
	default:
		return fmt.Sprintf("KIOFB(%d)", e)
	}
}

type KIOFBAVSignalType uint

const (
	KIOFBAVSignalTypeDP KIOFBAVSignalType = 0
	KIOFBAVSignalTypeDVI KIOFBAVSignalType = 0
	KIOFBAVSignalTypeHDMI KIOFBAVSignalType = 0
	KIOFBAVSignalTypeUnknown KIOFBAVSignalType = 0
	KIOFBAVSignalTypeVGA KIOFBAVSignalType = 0
)

func (e KIOFBAVSignalType) String() string {
	switch e {
	case KIOFBAVSignalTypeDP:
		return "KIOFBAVSignalTypeDP"
	default:
		return fmt.Sprintf("KIOFBAVSignalType(%d)", e)
	}
}

type KIOFBBitRateRBR uint

const (
	KIOFBBitRateHBR KIOFBBitRateRBR = 0
)

func (e KIOFBBitRateRBR) String() string {
	switch e {
	case KIOFBBitRateHBR:
		return "KIOFBBitRateHBR"
	default:
		return fmt.Sprintf("KIOFBBitRateRBR(%d)", e)
	}
}

type KIOFBCursorImageNew uint

const (
	KIOFBCursorHWCapable KIOFBCursorImageNew = 0
)

func (e KIOFBCursorImageNew) String() string {
	switch e {
	case KIOFBCursorHWCapable:
		return "KIOFBCursorHWCapable"
	default:
		return fmt.Sprintf("KIOFBCursorImageNew(%d)", e)
	}
}

type KIOFBDisplayState uint

const (
	KIOFBDisplayState_AlreadyActive KIOFBDisplayState = 0
	KIOFBDisplayState_PipelineBlack KIOFBDisplayState = 0
)

func (e KIOFBDisplayState) String() string {
	switch e {
	case KIOFBDisplayState_AlreadyActive:
		return "KIOFBDisplayState_AlreadyActive"
	default:
		return fmt.Sprintf("KIOFBDisplayState(%d)", e)
	}
}

const KIOFBHardwareCursorActive uint = 0

type KIOFBLinkDownspread uint

const (
	KIOFBLinkDownspreadMax KIOFBLinkDownspread = 0
	KIOFBLinkDownspreadNone KIOFBLinkDownspread = 0
)

func (e KIOFBLinkDownspread) String() string {
	switch e {
	case KIOFBLinkDownspreadMax:
		return "KIOFBLinkDownspreadMax"
	default:
		return fmt.Sprintf("KIOFBLinkDownspread(%d)", e)
	}
}

type KIOFBLinkPreEmphasis uint

const (
	KIOFBLinkPreEmphasisLevel0 KIOFBLinkPreEmphasis = 0
	KIOFBLinkPreEmphasisLevel1 KIOFBLinkPreEmphasis = 0
)

func (e KIOFBLinkPreEmphasis) String() string {
	switch e {
	case KIOFBLinkPreEmphasisLevel0:
		return "KIOFBLinkPreEmphasisLevel0"
	default:
		return fmt.Sprintf("KIOFBLinkPreEmphasis(%d)", e)
	}
}

type KIOFBLinkScramblerNormal uint

const (
	KIOFBLinkScramblerAlternate KIOFBLinkScramblerNormal = 0
)

func (e KIOFBLinkScramblerNormal) String() string {
	switch e {
	case KIOFBLinkScramblerAlternate:
		return "KIOFBLinkScramblerAlternate"
	default:
		return fmt.Sprintf("KIOFBLinkScramblerNormal(%d)", e)
	}
}

type KIOFBLinkVoltage uint

const (
	KIOFBLinkVoltageLevel0 KIOFBLinkVoltage = 0
	KIOFBLinkVoltageLevel1 KIOFBLinkVoltage = 0
	KIOFBLinkVoltageLevel2 KIOFBLinkVoltage = 0
	KIOFBLinkVoltageLevel3 KIOFBLinkVoltage = 0
)

func (e KIOFBLinkVoltage) String() string {
	switch e {
	case KIOFBLinkVoltageLevel0:
		return "KIOFBLinkVoltageLevel0"
	default:
		return fmt.Sprintf("KIOFBLinkVoltage(%d)", e)
	}
}

type KIOFBNS uint

const (
	KIOFBNS_DisplayStateShift KIOFBNS = 0
	KIOFBNS_Doze KIOFBNS = 0
)

func (e KIOFBNS) String() string {
	switch e {
	case KIOFBNS_DisplayStateShift:
		return "KIOFBNS_DisplayStateShift"
	default:
		return fmt.Sprintf("KIOFBNS(%d)", e)
	}
}

type KIOFBNotify uint

const (
	KIOFBNotifyDidSleep KIOFBNotify = 0
	KIOFBNotifyHDACodecDidPowerOff KIOFBNotify = 0
)

func (e KIOFBNotify) String() string {
	switch e {
	case KIOFBNotifyDidSleep:
		return "KIOFBNotifyDidSleep"
	default:
		return fmt.Sprintf("KIOFBNotify(%d)", e)
	}
}

type KIOFBNotifyEvent uint

const (
	KIOFBNotifyEvent_All KIOFBNotifyEvent = 0
	KIOFBNotifyEvent_ChangeSpeed KIOFBNotifyEvent = 0
)

func (e KIOFBNotifyEvent) String() string {
	switch e {
	case KIOFBNotifyEvent_All:
		return "KIOFBNotifyEvent_All"
	default:
		return fmt.Sprintf("KIOFBNotifyEvent(%d)", e)
	}
}

type KIOFBNotifyGroupID uint

const (
	KIOFBNotifyGroupID_AppleGraphicsDevicePolicy KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_Count KIOFBNotifyGroupID = 0
)

func (e KIOFBNotifyGroupID) String() string {
	switch e {
	case KIOFBNotifyGroupID_AppleGraphicsDevicePolicy:
		return "KIOFBNotifyGroupID_AppleGraphicsDevicePolicy"
	default:
		return fmt.Sprintf("KIOFBNotifyGroupID(%d)", e)
	}
}

const KIOFBSystemAperture uint = 0

const KIOFBUserRequestProbe uint = 0

type KIOFWAVCProtocolUserClient uint

const (
	KIOFWAVCProtocolUserClientAVCRequestNotHandled KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientAddSubunit KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientAllocateInputPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientAllocateOutputPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientConnectTargetPlugs KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientDisconnectTargetPlugs KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientFreeInputPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientFreeOutputPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientGetSubunitPlugSignalFormat KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientGetTargetPlugConnection KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientInstallAVCCommandHandler KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientNumAsyncCommands KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientNumCommands KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientPublishAVCUnitDirectory KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientReadInputMasterPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientReadInputPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientReadOutputMasterPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientReadOutputPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientSendAVCResponse KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientSetAVCRequestCallback KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientSetSubunitPlugSignalFormat KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientUpdateInputMasterPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientUpdateInputPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientUpdateOutputMasterPlug KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientUpdateOutputPlug KIOFWAVCProtocolUserClient = 0
)

func (e KIOFWAVCProtocolUserClient) String() string {
	switch e {
	case KIOFWAVCProtocolUserClientAVCRequestNotHandled:
		return "KIOFWAVCProtocolUserClientAVCRequestNotHandled"
	default:
		return fmt.Sprintf("KIOFWAVCProtocolUserClient(%d)", e)
	}
}

type KIOFWAVCSubunitPlugMsg uint

const (
	KIOFWAVCSubunitPlugMsgConnected KIOFWAVCSubunitPlugMsg = 0
	KIOFWAVCSubunitPlugMsgConnectedPlugModified KIOFWAVCSubunitPlugMsg = 0
	KIOFWAVCSubunitPlugMsgDisconnected KIOFWAVCSubunitPlugMsg = 0
	KIOFWAVCSubunitPlugMsgSignalFormatModified KIOFWAVCSubunitPlugMsg = 0
)

func (e KIOFWAVCSubunitPlugMsg) String() string {
	switch e {
	case KIOFWAVCSubunitPlugMsgConnected:
		return "KIOFWAVCSubunitPlugMsgConnected"
	default:
		return fmt.Sprintf("KIOFWAVCSubunitPlugMsg(%d)", e)
	}
}

type KIOFWAVCUserClient uint

const (
	KIOFWAVCUserClientAVCCommand KIOFWAVCUserClient = 0
	KIOFWAVCUserClientAVCCommandInGen KIOFWAVCUserClient = 0
	KIOFWAVCUserClientBreakP2PInputConnection KIOFWAVCUserClient = 0
	KIOFWAVCUserClientBreakP2POutputConnection KIOFWAVCUserClient = 0
	KIOFWAVCUserClientCancelAsyncAVCCommand KIOFWAVCUserClient = 0
	KIOFWAVCUserClientClose KIOFWAVCUserClient = 0
	KIOFWAVCUserClientCreateAsyncAVCCommand KIOFWAVCUserClient = 0
	KIOFWAVCUserClientGetSessionRef KIOFWAVCUserClient = 0
	KIOFWAVCUserClientInstallAsyncAVCCommandCallback KIOFWAVCUserClient = 0
	KIOFWAVCUserClientMakeP2PInputConnection KIOFWAVCUserClient = 0
	KIOFWAVCUserClientMakeP2POutputConnection KIOFWAVCUserClient = 0
	KIOFWAVCUserClientNumAsyncCommands KIOFWAVCUserClient = 0
	KIOFWAVCUserClientNumCommands KIOFWAVCUserClient = 0
	KIOFWAVCUserClientOpen KIOFWAVCUserClient = 0
	KIOFWAVCUserClientOpenWithSessionRef KIOFWAVCUserClient = 0
	KIOFWAVCUserClientReinitAsyncAVCCommand KIOFWAVCUserClient = 0
	KIOFWAVCUserClientReleaseAsyncAVCCommand KIOFWAVCUserClient = 0
	KIOFWAVCUserClientSubmitAsyncAVCCommand KIOFWAVCUserClient = 0
	KIOFWAVCUserClientUpdateAVCCommandTimeout KIOFWAVCUserClient = 0
)

func (e KIOFWAVCUserClient) String() string {
	switch e {
	case KIOFWAVCUserClientAVCCommand:
		return "KIOFWAVCUserClientAVCCommand"
	default:
		return fmt.Sprintf("KIOFWAVCUserClient(%d)", e)
	}
}

type KIOFWDisablePhysicalAccess uint

const (
	// KIOFWDisableAllPhysicalAccess: # Discussion
	KIOFWDisableAllPhysicalAccess KIOFWDisablePhysicalAccess = 0
)

func (e KIOFWDisablePhysicalAccess) String() string {
	switch e {
	case KIOFWDisableAllPhysicalAccess:
		return "KIOFWDisableAllPhysicalAccess"
	default:
		return fmt.Sprintf("KIOFWDisablePhysicalAccess(%d)", e)
	}
}

type KIOFWPhysicalAccess uint

const (
	KIOFWPhysicalAccessDisabled KIOFWPhysicalAccess = 0
	KIOFWPhysicalAccessDisabledForGeneration KIOFWPhysicalAccess = 0
	KIOFWPhysicalAccessEnabled KIOFWPhysicalAccess = 0
)

func (e KIOFWPhysicalAccess) String() string {
	switch e {
	case KIOFWPhysicalAccessDisabled:
		return "KIOFWPhysicalAccessDisabled"
	default:
		return fmt.Sprintf("KIOFWPhysicalAccess(%d)", e)
	}
}

type KIOFWRead uint

const (
	KIOFWReadBlockRequest KIOFWRead = 0
	KIOFWReadFlagsNone KIOFWRead = 0
	KIOFWReadPingTime KIOFWRead = 0
)

func (e KIOFWRead) String() string {
	switch e {
	case KIOFWReadBlockRequest:
		return "KIOFWReadBlockRequest"
	default:
		return fmt.Sprintf("KIOFWRead(%d)", e)
	}
}

type KIOFWSBP2 uint

const (
	KIOFWSBP2DontUsePTPacketLimit KIOFWSBP2 = 0
	KIOFWSBP2FailsOnAckBusy KIOFWSBP2 = 0
)

func (e KIOFWSBP2) String() string {
	switch e {
	case KIOFWSBP2DontUsePTPacketLimit:
		return "KIOFWSBP2DontUsePTPacketLimit"
	default:
		return fmt.Sprintf("KIOFWSBP2(%d)", e)
	}
}

type KIOFWSBP2UserClient uint

const (
	KIOFWSBP2UserClientClose KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientCreateLogin KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientCreateMgmtORB KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientCreateORB KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientEnableUnsolicitedStatus KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientGetLoginID KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientGetMaxCommandBlockSize KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientGetSessionRef KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientLSIWorkaroundSetCommandBuffersAsRanges KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBLSIWorkaroundSyncBuffersForInput KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBLSIWorkaroundSyncBuffersForOutput KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBSetCommandFunction KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBSetManageeLogin KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBSetManageeORB KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBSetResponseBuffer KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientNumCommands KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientOpen KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientOpenWithSessionRef KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientReleaseCommandBuffers KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientReleaseLogin KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientReleaseMgmtORB KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientReleaseORB KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientRingDoorbell KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetBusyTimeoutRegisterValue KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandBlock KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandBuffersAsRanges KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandFlags KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandGeneration KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandTimeout KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetFetchAgentWriteCompletion KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetLoginCallback KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetLoginFlags KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetLogoutCallback KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetMaxORBPayloadSize KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetMaxPayloadSize KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetMessageCallback KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetMgmtORBCallback KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetORBRefCon KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetPassword KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetReconnectTime KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetStatusNotify KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetToDummy KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetUnsolicitedStatusNotify KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitFetchAgentReset KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitLogin KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitLogout KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitMgmtORB KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitORB KIOFWSBP2UserClient = 0
)

func (e KIOFWSBP2UserClient) String() string {
	switch e {
	case KIOFWSBP2UserClientClose:
		return "KIOFWSBP2UserClientClose"
	default:
		return fmt.Sprintf("KIOFWSBP2UserClient(%d)", e)
	}
}

type KIOFWSecurityMode uint

const (
	KIOFWSecurityModeNormal KIOFWSecurityMode = 0
	KIOFWSecurityModeSecure KIOFWSecurityMode = 0
	KIOFWSecurityModeSecurePermanent KIOFWSecurityMode = 0
)

func (e KIOFWSecurityMode) String() string {
	switch e {
	case KIOFWSecurityModeNormal:
		return "KIOFWSecurityModeNormal"
	default:
		return fmt.Sprintf("KIOFWSecurityMode(%d)", e)
	}
}

type KIOFWWrite uint

const (
	KIOFWWriteBlockRequest KIOFWWrite = 0
	KIOFWWriteFastRetryOnBusy KIOFWWrite = 0
	KIOFWWriteFlagsDeferredNotify KIOFWWrite = 0
	KIOFWWriteFlagsNone KIOFWWrite = 0
)

func (e KIOFWWrite) String() string {
	switch e {
	case KIOFWWriteBlockRequest:
		return "KIOFWWriteBlockRequest"
	default:
		return fmt.Sprintf("KIOFWWrite(%d)", e)
	}
}

type KIOHID uint

const (
	KIOHIDActivityDisplayOn KIOHID = 0
	KIOHIDActivityUserIdle KIOHID = 0
	KIOHIDCapsLockState KIOHID = 0
	KIOHIDNumLockState KIOHID = 0
)

func (e KIOHID) String() string {
	switch e {
	case KIOHIDActivityDisplayOn:
		return "KIOHIDActivityDisplayOn"
	default:
		return fmt.Sprintf("KIOHID(%d)", e)
	}
}

type KIOHIDAccelerationAlgorithmTypeTable uint

const (
	KIOHIDAccelerationAlgorithmTypeDefault KIOHIDAccelerationAlgorithmTypeTable = 0
)

func (e KIOHIDAccelerationAlgorithmTypeTable) String() string {
	switch e {
	case KIOHIDAccelerationAlgorithmTypeDefault:
		return "KIOHIDAccelerationAlgorithmTypeDefault"
	default:
		return fmt.Sprintf("KIOHIDAccelerationAlgorithmTypeTable(%d)", e)
	}
}

type KIOHIDButtonMode uint

const (
	KIOHIDButtonMode_BothLeftClicks KIOHIDButtonMode = 0
	KIOHIDButtonMode_EnableRightClick KIOHIDButtonMode = 0
	KIOHIDButtonMode_ReverseLeftRightClicks KIOHIDButtonMode = 0
)

func (e KIOHIDButtonMode) String() string {
	switch e {
	case KIOHIDButtonMode_BothLeftClicks:
		return "KIOHIDButtonMode_BothLeftClicks"
	default:
		return fmt.Sprintf("KIOHIDButtonMode(%d)", e)
	}
}

type KIOHIDElementFlagsConstantMask uint

const (
	KIOHIDElementFlagsBufferedByteMask KIOHIDElementFlagsConstantMask = 0
)

func (e KIOHIDElementFlagsConstantMask) String() string {
	switch e {
	case KIOHIDElementFlagsBufferedByteMask:
		return "KIOHIDElementFlagsBufferedByteMask"
	default:
		return fmt.Sprintf("KIOHIDElementFlagsConstantMask(%d)", e)
	}
}

const KIOHIDEventNotification uint = 0

const KIOHIDEventQueueTypeKernel uint = 0

const KIOHIDGlobalMemory uint = 0

type KIOHIDKeyboardPhysicalLayout uint

const (
	KIOHIDKeyboardPhysicalLayoutType101 KIOHIDKeyboardPhysicalLayout = 0
	KIOHIDKeyboardPhysicalLayoutType106 KIOHIDKeyboardPhysicalLayout = 0
)

func (e KIOHIDKeyboardPhysicalLayout) String() string {
	switch e {
	case KIOHIDKeyboardPhysicalLayoutType101:
		return "KIOHIDKeyboardPhysicalLayoutType101"
	default:
		return fmt.Sprintf("KIOHIDKeyboardPhysicalLayout(%d)", e)
	}
}

const KIOHIDOpenedByEventSystem uint = 0

type KIOHIDOptionsType uint

const (
	KIOHIDOptionsTypeMaskPrivate KIOHIDOptionsType = 0
	// KIOHIDOptionsTypeNone: # Discussion
	KIOHIDOptionsTypeNone KIOHIDOptionsType = 0
)

func (e KIOHIDOptionsType) String() string {
	switch e {
	case KIOHIDOptionsTypeMaskPrivate:
		return "KIOHIDOptionsTypeMaskPrivate"
	default:
		return fmt.Sprintf("KIOHIDOptionsType(%d)", e)
	}
}

type KIOHIDQueueOptionsTypeNone uint

const (
	// KIOHIDQueueOptionsTypeEnqueueAll: # Discussion
	KIOHIDQueueOptionsTypeEnqueueAll KIOHIDQueueOptionsTypeNone = 0
)

func (e KIOHIDQueueOptionsTypeNone) String() string {
	switch e {
	case KIOHIDQueueOptionsTypeEnqueueAll:
		return "KIOHIDQueueOptionsTypeEnqueueAll"
	default:
		return fmt.Sprintf("KIOHIDQueueOptionsTypeNone(%d)", e)
	}
}

type KIOHIDServerConnectType uint

const (
	KIOHIDEventSystemConnectType KIOHIDServerConnectType = 0
)

func (e KIOHIDServerConnectType) String() string {
	switch e {
	case KIOHIDEventSystemConnectType:
		return "KIOHIDEventSystemConnectType"
	default:
		return fmt.Sprintf("KIOHIDServerConnectType(%d)", e)
	}
}

type KIOHIDStandardType uint

const (
	// KIOHIDStandardTypeJIS: # Discussion
	KIOHIDStandardTypeJIS KIOHIDStandardType = 0
	KIOHIDStandardTypeUnspecified KIOHIDStandardType = 0
)

func (e KIOHIDStandardType) String() string {
	switch e {
	case KIOHIDStandardTypeJIS:
		return "KIOHIDStandardTypeJIS"
	default:
		return fmt.Sprintf("KIOHIDStandardType(%d)", e)
	}
}

type KIOHIDValueOptionsFlag uint

const (
	KIOHIDValueOptionsFlagPrevious KIOHIDValueOptionsFlag = 0
	KIOHIDValueOptionsFlagRelativeSimple KIOHIDValueOptionsFlag = 0
)

func (e KIOHIDValueOptionsFlag) String() string {
	switch e {
	case KIOHIDValueOptionsFlagPrevious:
		return "KIOHIDValueOptionsFlagPrevious"
	default:
		return fmt.Sprintf("KIOHIDValueOptionsFlag(%d)", e)
	}
}

type KIOHIDValueScaleType uint

const (
	// KIOHIDValueScaleTypeCalibrated: # Discussion
	KIOHIDValueScaleTypeCalibrated KIOHIDValueScaleType = 0
	KIOHIDValueScaleTypeExponent KIOHIDValueScaleType = 0
)

func (e KIOHIDValueScaleType) String() string {
	switch e {
	case KIOHIDValueScaleTypeCalibrated:
		return "KIOHIDValueScaleTypeCalibrated"
	default:
		return fmt.Sprintf("KIOHIDValueScaleType(%d)", e)
	}
}

type KIOHSyncDisable uint

const (
	KIOCSyncDisable KIOHSyncDisable = 0
)

func (e KIOHSyncDisable) String() string {
	switch e {
	case KIOCSyncDisable:
		return "KIOCSyncDisable"
	default:
		return fmt.Sprintf("KIOHSyncDisable(%d)", e)
	}
}

const KIOHibernatePreviewActive uint = 0

const KIOInterruptDispatchSourceTypeEdge uint = 0

type KIOInterruptSourceIndexMask uint

const (
	KIOInterruptSourceAbsoluteTime KIOInterruptSourceIndexMask = 0
)

func (e KIOInterruptSourceIndexMask) String() string {
	switch e {
	case KIOInterruptSourceAbsoluteTime:
		return "KIOInterruptSourceAbsoluteTime"
	default:
		return fmt.Sprintf("KIOInterruptSourceIndexMask(%d)", e)
	}
}

const KIOKitDebugUserOptions uint = 0

const KIOKitDiagnosticsClientType uint = 0

type KIOMap uint

const (
	KIOMapAnywhere KIOMap = 0
	KIOMapPostedWrite KIOMap = 0
)

func (e KIOMap) String() string {
	switch e {
	case KIOMapAnywhere:
		return "KIOMapAnywhere"
	default:
		return fmt.Sprintf("KIOMap(%d)", e)
	}
}

const KIOMapperUncached uint = 0

type KIOMaxBus uint

const (
	KIOMaxBusStall10usec KIOMaxBus = 0
	KIOMaxBusStall20usec KIOMaxBus = 0
)

func (e KIOMaxBus) String() string {
	switch e {
	case KIOMaxBusStall10usec:
		return "KIOMaxBusStall10usec"
	default:
		return fmt.Sprintf("KIOMaxBus(%d)", e)
	}
}

const KIOMaxPixelBits uint = 0

const KIOMediaAttributeEjectableMask uint = 0

type KIOMediaStateOffline uint

const (
	// KIOMediaStateBusy: # Discussion
	KIOMediaStateBusy KIOMediaStateOffline = 0
)

func (e KIOMediaStateOffline) String() string {
	switch e {
	case KIOMediaStateBusy:
		return "KIOMediaStateBusy"
	default:
		return fmt.Sprintf("KIOMediaStateOffline(%d)", e)
	}
}

type KIOMediumEthernet uint

const (
	KIOMediumEthernetAuto KIOMediumEthernet = 0
	KIOMediumEthernetHomePNA1 KIOMediumEthernet = 0
)

func (e KIOMediumEthernet) String() string {
	switch e {
	case KIOMediumEthernetAuto:
		return "KIOMediumEthernetAuto"
	default:
		return fmt.Sprintf("KIOMediumEthernet(%d)", e)
	}
}

type KIOMediumIEE uint

const (
	KIOMediumIEEE80211 KIOMediumIEE = 0
	KIOMediumIEEE80211Manual KIOMediumIEE = 0
)

func (e KIOMediumIEE) String() string {
	switch e {
	case KIOMediumIEEE80211:
		return "KIOMediumIEEE80211"
	default:
		return fmt.Sprintf("KIOMediumIEE(%d)", e)
	}
}

type KIOMediumOption uint

const (
	KIOMediumOptionEEE KIOMediumOption = 0
	KIOMediumOptionFlowControl KIOMediumOption = 0
)

func (e KIOMediumOption) String() string {
	switch e {
	case KIOMediumOptionEEE:
		return "KIOMediumOptionEEE"
	default:
		return fmt.Sprintf("KIOMediumOption(%d)", e)
	}
}

type KIOMemoryDirection uint

const (
	KIOMemoryDirectionIn KIOMemoryDirection = 0
	KIOMemoryDirectionNone KIOMemoryDirection = 0
)

func (e KIOMemoryDirection) String() string {
	switch e {
	case KIOMemoryDirectionIn:
		return "KIOMemoryDirectionIn"
	default:
		return fmt.Sprintf("KIOMemoryDirection(%d)", e)
	}
}

type KIOMemoryDirectionMask uint

const (
	KIOMemoryUseReserve KIOMemoryDirectionMask = 0
)

func (e KIOMemoryDirectionMask) String() string {
	switch e {
	case KIOMemoryUseReserve:
		return "KIOMemoryUseReserve"
	default:
		return fmt.Sprintf("KIOMemoryDirectionMask(%d)", e)
	}
}

type KIOMemoryIncoherentIOFlush uint

const (
	KIOMemoryClearEncrypted KIOMemoryIncoherentIOFlush = 0
)

func (e KIOMemoryIncoherentIOFlush) String() string {
	switch e {
	case KIOMemoryClearEncrypted:
		return "KIOMemoryClearEncrypted"
	default:
		return fmt.Sprintf("KIOMemoryIncoherentIOFlush(%d)", e)
	}
}

const KIOMemoryLedgerFlagNoFootprint uint = 0

const KIOMemoryLedgerTagDefault uint = 0

type KIOMemoryMap uint

const (
	KIOMemoryMapCacheModeDefault KIOMemoryMap = 0
	KIOMemoryMapGuardedSmall KIOMemoryMap = 0
)

func (e KIOMemoryMap) String() string {
	switch e {
	case KIOMemoryMapCacheModeDefault:
		return "KIOMemoryMapCacheModeDefault"
	default:
		return fmt.Sprintf("KIOMemoryMap(%d)", e)
	}
}

type KIOMemoryPurgeable uint

const (
	KIOMemoryPurgeableFaultOnAccess KIOMemoryPurgeable = 0
	KIOMemoryPurgeableVolatileGroup6 KIOMemoryPurgeable = 0
)

func (e KIOMemoryPurgeable) String() string {
	switch e {
	case KIOMemoryPurgeableFaultOnAccess:
		return "KIOMemoryPurgeableFaultOnAccess"
	default:
		return fmt.Sprintf("KIOMemoryPurgeable(%d)", e)
	}
}

const KIOMirrorDefault uint = 0

type KIOMirrorIsPrimary uint

const (
	KIOMirrorHWClipped KIOMirrorIsPrimary = 0
)

func (e KIOMirrorIsPrimary) String() string {
	switch e {
	case KIOMirrorHWClipped:
		return "KIOMirrorHWClipped"
	default:
		return fmt.Sprintf("KIOMirrorIsPrimary(%d)", e)
	}
}

type KIONDRV uint

const (
	KIONDRVCloseCommand KIONDRV = 0
	KIONDRVFinalizeCommand KIONDRV = 0
)

func (e KIONDRV) String() string {
	switch e {
	case KIONDRVCloseCommand:
		return "KIONDRVCloseCommand"
	default:
		return fmt.Sprintf("KIONDRV(%d)", e)
	}
}

type KIONDRVSynchronousIOCommandKind uint

const (
	KIONDRVAsynchronousIOCommandKind KIONDRVSynchronousIOCommandKind = 0
)

func (e KIONDRVSynchronousIOCommandKind) String() string {
	switch e {
	case KIONDRVAsynchronousIOCommandKind:
		return "KIONDRVAsynchronousIOCommandKind"
	default:
		return fmt.Sprintf("KIONDRVSynchronousIOCommandKind(%d)", e)
	}
}

type KIONVRAMOperation uint

const (
	KIONVRAMOperationDelete KIONVRAMOperation = 0
	KIONVRAMOperationObliterate KIONVRAMOperation = 0
	KIONVRAMOperationRead KIONVRAMOperation = 0
	KIONVRAMOperationReset KIONVRAMOperation = 0
	KIONVRAMOperationWrite KIONVRAMOperation = 0
)

func (e KIONVRAMOperation) String() string {
	switch e {
	case KIONVRAMOperationDelete:
		return "KIONVRAMOperationDelete"
	default:
		return fmt.Sprintf("KIONVRAMOperation(%d)", e)
	}
}

type KIONetworkDataAccessType uint

const (
	KIONetworkDataAccessTypeMask KIONetworkDataAccessType = 0
	// KIONetworkDataAccessTypeRead: # Discussion
	KIONetworkDataAccessTypeRead KIONetworkDataAccessType = 0
)

func (e KIONetworkDataAccessType) String() string {
	switch e {
	case KIONetworkDataAccessTypeMask:
		return "KIONetworkDataAccessTypeMask"
	default:
		return fmt.Sprintf("KIONetworkDataAccessType(%d)", e)
	}
}

type KIONetworkDataBufferTypeInternal uint

const (
	// KIONetworkDataBufferTypeExternal: # Discussion
	KIONetworkDataBufferTypeExternal KIONetworkDataBufferTypeInternal = 0
)

func (e KIONetworkDataBufferTypeInternal) String() string {
	switch e {
	case KIONetworkDataBufferTypeExternal:
		return "KIONetworkDataBufferTypeExternal"
	default:
		return fmt.Sprintf("KIONetworkDataBufferTypeInternal(%d)", e)
	}
}

const KIONetworkEventTypeDLIL uint = 0

type KIONetworkFeature uint

const (
	KIONetworkFeatureHWTimeStamp KIONetworkFeature = 0
	KIONetworkFeatureSWTimeStamp KIONetworkFeature = 0
)

func (e KIONetworkFeature) String() string {
	switch e {
	case KIONetworkFeatureHWTimeStamp:
		return "KIONetworkFeatureHWTimeStamp"
	default:
		return fmt.Sprintf("KIONetworkFeature(%d)", e)
	}
}

type KIONetworkInterfaceRegisteredState uint

const (
	// KIONetworkInterfaceDisabledState: # Discussion
	KIONetworkInterfaceDisabledState KIONetworkInterfaceRegisteredState = 0
)

func (e KIONetworkInterfaceRegisteredState) String() string {
	switch e {
	case KIONetworkInterfaceDisabledState:
		return "KIONetworkInterfaceDisabledState"
	default:
		return fmt.Sprintf("KIONetworkInterfaceRegisteredState(%d)", e)
	}
}

type KIONetworkLink uint

const (
	KIONetworkLinkActive KIONetworkLink = 0
	KIONetworkLinkNoNetworkChange KIONetworkLink = 0
)

func (e KIONetworkLink) String() string {
	switch e {
	case KIONetworkLinkActive:
		return "KIONetworkLinkActive"
	default:
		return fmt.Sprintf("KIONetworkLink(%d)", e)
	}
}

const KIOOutputCommandMask uint = 0

type KIOOutputStatus int

const (
	KIOOutputStatusAccepted KIOOutputStatus = 0
	// KIOOutputStatusDropped: # Discussion
	KIOOutputStatusDropped KIOOutputStatus = 0
	// KIOOutputStatusMask: # Discussion
	KIOOutputStatusMask KIOOutputStatus = 0
	// KIOOutputStatusRetry: # Discussion
	KIOOutputStatusRetry KIOOutputStatus = 0
)

func (e KIOOutputStatus) String() string {
	switch e {
	case KIOOutputStatusAccepted:
		return "KIOOutputStatusAccepted"
	default:
		return fmt.Sprintf("KIOOutputStatus(%d)", e)
	}
}

type KIOPC uint

const (
	KIOPCI32BitMemorySpace KIOPC = 0
	KIOPCIConfigSpace KIOPC = 0
)

func (e KIOPC) String() string {
	switch e {
	case KIOPCI32BitMemorySpace:
		return "KIOPCI32BitMemorySpace"
	default:
		return fmt.Sprintf("KIOPC(%d)", e)
	}
}

type KIOPCI uint

const (
	KIOPCIAGP8Capability KIOPCI = 0
	KIOPCIExpressAccessControlServicesCapability KIOPCI = 0
	KIOPCIExpressErrorReportingCapability KIOPCI = 0
)

func (e KIOPCI) String() string {
	switch e {
	case KIOPCIAGP8Capability:
		return "KIOPCIAGP8Capability"
	default:
		return fmt.Sprintf("KIOPCI(%d)", e)
	}
}

type KIOPCICapabilityIDOffset uint

const (
	KIOPCINextCapabilityOffset KIOPCICapabilityIDOffset = 0
)

func (e KIOPCICapabilityIDOffset) String() string {
	switch e {
	case KIOPCINextCapabilityOffset:
		return "KIOPCINextCapabilityOffset"
	default:
		return fmt.Sprintf("KIOPCICapabilityIDOffset(%d)", e)
	}
}

type KIOPCICommand uint

const (
	KIOPCICommandAddressStepping KIOPCICommand = 0
	KIOPCICommandFastBack2Back KIOPCICommand = 0
)

func (e KIOPCICommand) String() string {
	switch e {
	case KIOPCICommandAddressStepping:
		return "KIOPCICommandAddressStepping"
	default:
		return fmt.Sprintf("KIOPCICommand(%d)", e)
	}
}

type KIOPCIConfig uint

const (
	KIOPCIConfigBaseAddress0 KIOPCIConfig = 0
	KIOPCIConfigExpansionROMBase KIOPCIConfig = 0
)

func (e KIOPCIConfig) String() string {
	switch e {
	case KIOPCIConfigBaseAddress0:
		return "KIOPCIConfigBaseAddress0"
	default:
		return fmt.Sprintf("KIOPCIConfig(%d)", e)
	}
}

type KIOPCIConfigAGPStatusOffset int

const (
	KIOPCIConfigAGPCommandOffset KIOPCIConfigAGPStatusOffset = 0
)

func (e KIOPCIConfigAGPStatusOffset) String() string {
	switch e {
	case KIOPCIConfigAGPCommandOffset:
		return "KIOPCIConfigAGPCommandOffset"
	default:
		return fmt.Sprintf("KIOPCIConfigAGPStatusOffset(%d)", e)
	}
}

const KIOPCIConfigShadowPermanent uint = 0

type KIOPCICorrectableErrorBit int

const (
	KIOPCICorrectableErrorBitBadDLLP KIOPCICorrectableErrorBit = 0
	KIOPCICorrectableErrorBitReceiver KIOPCICorrectableErrorBit = 0
)

func (e KIOPCICorrectableErrorBit) String() string {
	switch e {
	case KIOPCICorrectableErrorBitBadDLLP:
		return "KIOPCICorrectableErrorBitBadDLLP"
	default:
		return fmt.Sprintf("KIOPCICorrectableErrorBit(%d)", e)
	}
}

type KIOPCIDevicePowerStateCount uint

const (
	KIOPCIDeviceDozeState KIOPCIDevicePowerStateCount = 0
)

func (e KIOPCIDevicePowerStateCount) String() string {
	switch e {
	case KIOPCIDeviceDozeState:
		return "KIOPCIDeviceDozeState"
	default:
		return fmt.Sprintf("KIOPCIDevicePowerStateCount(%d)", e)
	}
}

type KIOPCIDeviceResetOption uint

const (
	KIOPCIDeviceResetOptionNone KIOPCIDeviceResetOption = 0
	KIOPCIDeviceResetOptionTerminate KIOPCIDeviceResetOption = 0
)

func (e KIOPCIDeviceResetOption) String() string {
	switch e {
	case KIOPCIDeviceResetOptionNone:
		return "KIOPCIDeviceResetOptionNone"
	default:
		return fmt.Sprintf("KIOPCIDeviceResetOption(%d)", e)
	}
}

type KIOPCIDeviceResetType uint

const (
	KIOPCIDeviceResetTypeFunctionReset KIOPCIDeviceResetType = 0
	KIOPCIDeviceResetTypeHotReset KIOPCIDeviceResetType = 0
	KIOPCIDeviceResetTypeWarmReset KIOPCIDeviceResetType = 0
	KIOPCIDeviceResetTypeWarmResetDisable KIOPCIDeviceResetType = 0
	KIOPCIDeviceResetTypeWarmResetEnable KIOPCIDeviceResetType = 0
)

func (e KIOPCIDeviceResetType) String() string {
	switch e {
	case KIOPCIDeviceResetTypeFunctionReset:
		return "KIOPCIDeviceResetTypeFunctionReset"
	default:
		return fmt.Sprintf("KIOPCIDeviceResetType(%d)", e)
	}
}

type KIOPCIEvent uint

const (
	KIOPCIEventLinkEnableChange KIOPCIEvent = 0
	KIOPCIEventNonFatalError KIOPCIEvent = 0
)

func (e KIOPCIEvent) String() string {
	switch e {
	case KIOPCIEventLinkEnableChange:
		return "KIOPCIEventLinkEnableChange"
	default:
		return fmt.Sprintf("KIOPCIEvent(%d)", e)
	}
}

const KIOPCILatencySnooped uint = 0

type KIOPCILinkSpeed uint

const (
	KIOPCILinkSpeed_16_GTs KIOPCILinkSpeed = 0
	KIOPCILinkSpeed_2_5_GTs KIOPCILinkSpeed = 0
	KIOPCILinkSpeed_32_GTs KIOPCILinkSpeed = 0
	KIOPCILinkSpeed_5_GTs KIOPCILinkSpeed = 0
	KIOPCILinkSpeed_8_GTs KIOPCILinkSpeed = 0
)

func (e KIOPCILinkSpeed) String() string {
	switch e {
	case KIOPCILinkSpeed_16_GTs:
		return "KIOPCILinkSpeed_16_GTs"
	default:
		return fmt.Sprintf("KIOPCILinkSpeed(%d)", e)
	}
}

const KIOPCIProbeOptionDone uint = 0

type KIOPCIStatus int

const (
	KIOPCIStatusCapabilities KIOPCIStatus = 0
	KIOPCIStatusDevSel1 KIOPCIStatus = 0
)

func (e KIOPCIStatus) String() string {
	switch e {
	case KIOPCIStatusCapabilities:
		return "KIOPCIStatusCapabilities"
	default:
		return fmt.Sprintf("KIOPCIStatus(%d)", e)
	}
}

type KIOPCIUncorrectableErrorBit int

const (
	KIOPCIUncorrectableErrorBitCompletionTimeout KIOPCIUncorrectableErrorBit = 0
	KIOPCIUncorrectableErrorBitUnsupportedRequest KIOPCIUncorrectableErrorBit = 0
)

func (e KIOPCIUncorrectableErrorBit) String() string {
	switch e {
	case KIOPCIUncorrectableErrorBitCompletionTimeout:
		return "KIOPCIUncorrectableErrorBitCompletionTimeout"
	default:
		return fmt.Sprintf("KIOPCIUncorrectableErrorBit(%d)", e)
	}
}

type KIOPM uint

const (
	KIOPMACInstalled KIOPM = 0
	KIOPMAllowSleep KIOPM = 0
	KIOPMAuxPowerOn KIOPM = 0
	KIOPMChildClamp KIOPM = 0
	KIOPMClamshellOpened KIOPM = 0
	KIOPMConfigRetained KIOPM = 0
	KIOPMFairValue KIOPM = 0
	KIOPMPreventSystemSleep KIOPM = 0
	KIOPMRawLowBattery KIOPM = 0
	// KIOPMRestartCapability: # Discussion
	KIOPMRestartCapability KIOPM = 0
	// KIOPMSleepCapability: # Discussion
	KIOPMSleepCapability KIOPM = 0
	KIOPMUndefinedValue KIOPM = 0
)

func (e KIOPM) String() string {
	switch e {
	case KIOPMACInstalled:
		return "KIOPMACInstalled"
	default:
		return fmt.Sprintf("KIOPM(%d)", e)
	}
}

const KIOPMBroadcastAggressiveness uint = 0

type KIOPMDriverAssertion uint

const (
	KIOPMDriverAssertionBluetoothHIDDevicePairedBit KIOPMDriverAssertion = 0
	KIOPMDriverAssertionCPUBit KIOPMDriverAssertion = 0
)

func (e KIOPMDriverAssertion) String() string {
	switch e {
	case KIOPMDriverAssertionBluetoothHIDDevicePairedBit:
		return "KIOPMDriverAssertionBluetoothHIDDevicePairedBit"
	default:
		return fmt.Sprintf("KIOPMDriverAssertion(%d)", e)
	}
}

type KIOPMInternalPower uint

const (
	KIOPMExternalPower KIOPMInternalPower = 0
)

func (e KIOPMInternalPower) String() string {
	switch e {
	case KIOPMExternalPower:
		return "KIOPMExternalPower"
	default:
		return fmt.Sprintf("KIOPMInternalPower(%d)", e)
	}
}

type KIOPMMaxPowerStates uint

const (
	IOPMMaxPowerStates KIOPMMaxPowerStates = 0
)

func (e KIOPMMaxPowerStates) String() string {
	switch e {
	case IOPMMaxPowerStates:
		return "IOPMMaxPowerStates"
	default:
		return fmt.Sprintf("KIOPMMaxPowerStates(%d)", e)
	}
}

type KIOPMNextHigherState uint

const (
	KIOPMHighestState KIOPMNextHigherState = 0
)

func (e KIOPMNextHigherState) String() string {
	switch e {
	case KIOPMHighestState:
		return "KIOPMHighestState"
	default:
		return fmt.Sprintf("KIOPMNextHigherState(%d)", e)
	}
}

type KIOPMNoErr int

const (
	IOPMAckImplied KIOPMNoErr = 0
	KIOPMBadSpecification KIOPMNoErr = 0
)

func (e KIOPMNoErr) String() string {
	switch e {
	case IOPMAckImplied:
		return "IOPMAckImplied"
	default:
		return fmt.Sprintf("KIOPMNoErr(%d)", e)
	}
}

const KIOPMPSLocationLeft uint = 0

const KIOPMPowerStateVersion1 uint = 0

const KIOPMSupportedOnAC uint = 0

type KIOPMSystemCapability uint

const (
	KIOPMSystemCapabilityAudio KIOPMSystemCapability = 0
	KIOPMSystemCapabilityCPU KIOPMSystemCapability = 0
	KIOPMSystemCapabilityGraphics KIOPMSystemCapability = 0
	KIOPMSystemCapabilityNetwork KIOPMSystemCapability = 0
)

func (e KIOPMSystemCapability) String() string {
	switch e {
	case KIOPMSystemCapabilityAudio:
		return "KIOPMSystemCapabilityAudio"
	default:
		return fmt.Sprintf("KIOPMSystemCapability(%d)", e)
	}
}

type KIOPMSystemCapabilityWillChange uint

const (
	// KIOPMSystemCapabilityDidChange: # Discussion
	KIOPMSystemCapabilityDidChange KIOPMSystemCapabilityWillChange = 0
)

func (e KIOPMSystemCapabilityWillChange) String() string {
	switch e {
	case KIOPMSystemCapabilityDidChange:
		return "KIOPMSystemCapabilityDidChange"
	default:
		return fmt.Sprintf("KIOPMSystemCapabilityWillChange(%d)", e)
	}
}

type KIOPMThermalLevel uint

const (
	KIOPMThermalLevelCritical KIOPMThermalLevel = 0
	KIOPMThermalLevelNormal KIOPMThermalLevel = 0
)

func (e KIOPMThermalLevel) String() string {
	switch e {
	case KIOPMThermalLevelCritical:
		return "KIOPMThermalLevelCritical"
	default:
		return fmt.Sprintf("KIOPMThermalLevel(%d)", e)
	}
}

const KIOPMUnknown uint = 0

type KIOPSAdapterErrorFlagNoErrors int

const (
	KIOPSAdapterErrorFlagDeviceNeedsToBeRepositioned KIOPSAdapterErrorFlagNoErrors = 0
)

func (e KIOPSAdapterErrorFlagNoErrors) String() string {
	switch e {
	case KIOPSAdapterErrorFlagDeviceNeedsToBeRepositioned:
		return "KIOPSAdapterErrorFlagDeviceNeedsToBeRepositioned"
	default:
		return fmt.Sprintf("KIOPSAdapterErrorFlagNoErrors(%d)", e)
	}
}

type KIOPSFamilyCode uint

const (
	KIOPSFamilyCodeAC KIOPSFamilyCode = 0
	KIOPSFamilyCodeDisconnected KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal2 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal3 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal4 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal5 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal6 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal7 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal8 KIOPSFamilyCode = 0
	KIOPSFamilyCodeFirewire KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBAdapter KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBCBrick KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBCPD KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBCTypeC KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBChargingPort KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBChargingPortDedicated KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBChargingPortDownstream KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBDevice KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBHost KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBHostSuspended KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBUnknown KIOPSFamilyCode = 0
	KIOPSFamilyCodeUnsupported KIOPSFamilyCode = 0
	KIOPSFamilyCodeUnsupportedRegion KIOPSFamilyCode = 0
)

func (e KIOPSFamilyCode) String() string {
	switch e {
	case KIOPSFamilyCodeAC:
		return "KIOPSFamilyCodeAC"
	default:
		return fmt.Sprintf("KIOPSFamilyCode(%d)", e)
	}
}

const KIOPacketBufferAlign1 uint = 0

type KIOPacketFilter uint

const (
	// KIOPacketFilterBroadcast: # Discussion
	KIOPacketFilterBroadcast KIOPacketFilter = 0
	// KIOPacketFilterMulticast: # Discussion
	KIOPacketFilterMulticast KIOPacketFilter = 0
)

func (e KIOPacketFilter) String() string {
	switch e {
	case KIOPacketFilterBroadcast:
		return "KIOPacketFilterBroadcast"
	default:
		return fmt.Sprintf("KIOPacketFilter(%d)", e)
	}
}

type KIOPixelEncoding uint

const (
	KIOPixelEncodingNotSupported KIOPixelEncoding = 0
	KIOPixelEncodingYCbCr422 KIOPixelEncoding = 0
)

func (e KIOPixelEncoding) String() string {
	switch e {
	case KIOPixelEncodingNotSupported:
		return "KIOPixelEncodingNotSupported"
	default:
		return fmt.Sprintf("KIOPixelEncoding(%d)", e)
	}
}

type KIOPreparationIDUnprepared uint

const (
	KIOPreparationIDAlwaysPrepared KIOPreparationIDUnprepared = 0
)

func (e KIOPreparationIDUnprepared) String() string {
	switch e {
	case KIOPreparationIDAlwaysPrepared:
		return "KIOPreparationIDAlwaysPrepared"
	default:
		return fmt.Sprintf("KIOPreparationIDUnprepared(%d)", e)
	}
}

type KIORPCMessage uint

const (
	KIORPCMessageError KIORPCMessage = 0
	KIORPCMessageLocalHost KIORPCMessage = 0
)

func (e KIORPCMessage) String() string {
	switch e {
	case KIORPCMessageError:
		return "KIORPCMessageError"
	default:
		return fmt.Sprintf("KIORPCMessage(%d)", e)
	}
}

const KIORPCMessageIDKernel uint = 0

const KIORPCVersion190615 uint = 0

type KIORangeBitsPerColor uint

const (
	KIORangeBitsPerColorComponent10 KIORangeBitsPerColor = 0
	KIORangeBitsPerColorComponent6 KIORangeBitsPerColor = 0
)

func (e KIORangeBitsPerColor) String() string {
	switch e {
	case KIORangeBitsPerColorComponent10:
		return "KIORangeBitsPerColorComponent10"
	default:
		return fmt.Sprintf("KIORangeBitsPerColor(%d)", e)
	}
}

type KIORangeColorimetry uint

const (
	KIORangeColorimetryBT2100 KIORangeColorimetry = 0
	KIORangeColorimetryNativeRGB KIORangeColorimetry = 0
)

func (e KIORangeColorimetry) String() string {
	switch e {
	case KIORangeColorimetryBT2100:
		return "KIORangeColorimetryBT2100"
	default:
		return fmt.Sprintf("KIORangeColorimetry(%d)", e)
	}
}

type KIORangeDynamicRangeDolby uint

const (
	KIORangeDynamicRangeDolbyNormalMode KIORangeDynamicRangeDolby = 0
	KIORangeDynamicRangeDolbyTunnelMode KIORangeDynamicRangeDolby = 0
)

func (e KIORangeDynamicRangeDolby) String() string {
	switch e {
	case KIORangeDynamicRangeDolbyNormalMode:
		return "KIORangeDynamicRangeDolbyNormalMode"
	default:
		return fmt.Sprintf("KIORangeDynamicRangeDolby(%d)", e)
	}
}

type KIORangePixelEncoding uint

const (
	KIORangePixelEncodingNotSupported KIORangePixelEncoding = 0
	KIORangePixelEncodingRGB444 KIORangePixelEncoding = 0
)

func (e KIORangePixelEncoding) String() string {
	switch e {
	case KIORangePixelEncodingNotSupported:
		return "KIORangePixelEncodingNotSupported"
	default:
		return fmt.Sprintf("KIORangePixelEncoding(%d)", e)
	}
}

type KIORangeSupports uint

const (
	KIORangeSupportsCompositeSync KIORangeSupports = 0
	KIORangeSupportsVSyncSerration KIORangeSupports = 0
)

func (e KIORangeSupports) String() string {
	switch e {
	case KIORangeSupportsCompositeSync:
		return "KIORangeSupportsCompositeSync"
	default:
		return fmt.Sprintf("KIORangeSupports(%d)", e)
	}
}

const KIORangeSupportsInterlacedCEATiming uint = 0

type KIORangeSupportsSignal uint

const (
	KIORangeSupportsSignal_0700_0000 KIORangeSupportsSignal = 0
	KIORangeSupportsSignal_1000_0400 KIORangeSupportsSignal = 0
)

func (e KIORangeSupportsSignal) String() string {
	switch e {
	case KIORangeSupportsSignal_0700_0000:
		return "KIORangeSupportsSignal_0700_0000"
	default:
		return fmt.Sprintf("KIORangeSupportsSignal(%d)", e)
	}
}

type KIORegistryIterateRecursively uint

const (
	KIORegistryIterateParents KIORegistryIterateRecursively = 0
)

func (e KIORegistryIterateRecursively) String() string {
	switch e {
	case KIORegistryIterateParents:
		return "KIORegistryIterateParents"
	default:
		return fmt.Sprintf("KIORegistryIterateRecursively(%d)", e)
	}
}

const KIOReportCopyChannelData uint = 0

type KIOReportEnable uint

const (
	KIOReportDisable KIOReportEnable = 0
)

func (e KIOReportEnable) String() string {
	switch e {
	case KIOReportDisable:
		return "KIOReportDisable"
	default:
		return fmt.Sprintf("KIOReportEnable(%d)", e)
	}
}

type KIOReportInvalidFormat uint

const (
	KIOReportFormatHistogram KIOReportInvalidFormat = 0
)

func (e KIOReportInvalidFormat) String() string {
	switch e {
	case KIOReportFormatHistogram:
		return "KIOReportFormatHistogram"
	default:
		return fmt.Sprintf("KIOReportInvalidFormat(%d)", e)
	}
}

type KIOReportQuantity uint

const (
	KIOReportQuantityCapacitance KIOReportQuantity = 0
	KIOReportQuantityPacketCount KIOReportQuantity = 0
)

func (e KIOReportQuantity) String() string {
	switch e {
	case KIOReportQuantityCapacitance:
		return "KIOReportQuantityCapacitance"
	default:
		return fmt.Sprintf("KIOReportQuantity(%d)", e)
	}
}

type KIOReturnOutputSuccess uint

const (
	// KIOReturnOutputDropped: # Discussion
	KIOReturnOutputDropped KIOReturnOutputSuccess = 0
)

func (e KIOReturnOutputSuccess) String() string {
	switch e {
	case KIOReturnOutputDropped:
		return "KIOReturnOutputDropped"
	default:
		return fmt.Sprintf("KIOReturnOutputSuccess(%d)", e)
	}
}

type KIOScaleCan uint

const (
	KIOScaleCanBorderInsetOnly KIOScaleCan = 0
	KIOScaleCanDownSamplePixels KIOScaleCan = 0
)

func (e KIOScaleCan) String() string {
	switch e {
	case KIOScaleCanBorderInsetOnly:
		return "KIOScaleCanBorderInsetOnly"
	default:
		return fmt.Sprintf("KIOScaleCan(%d)", e)
	}
}

type KIOScaleInvert uint

const (
	KIOScaleInvertX KIOScaleInvert = 0
	KIOScaleInvertY KIOScaleInvert = 0
)

func (e KIOScaleInvert) String() string {
	switch e {
	case KIOScaleInvertX:
		return "KIOScaleInvertX"
	default:
		return fmt.Sprintf("KIOScaleInvert(%d)", e)
	}
}

type KIOService uint

const (
	KIOServiceAsynchronous KIOService = 0
	KIOServiceDextRequirePowerForMatching KIOService = 0
)

func (e KIOService) String() string {
	switch e {
	case KIOServiceAsynchronous:
		return "KIOServiceAsynchronous"
	default:
		return fmt.Sprintf("KIOService(%d)", e)
	}
}

const KIOServiceFamilyCloseOptions uint = 0

const KIOServiceHaltStatePowerOff uint = 0

type KIOServiceInactiveState uint

const (
	KIOServiceFirstMatchState KIOServiceInactiveState = 0
)

func (e KIOServiceInactiveState) String() string {
	switch e {
	case KIOServiceFirstMatchState:
		return "KIOServiceFirstMatchState"
	default:
		return fmt.Sprintf("KIOServiceInactiveState(%d)", e)
	}
}

type KIOServiceNotificationType uint

const (
	KIOServiceNotificationTypeLast KIOServiceNotificationType = 0
	KIOServiceNotificationTypeNone KIOServiceNotificationType = 0
)

func (e KIOServiceNotificationType) String() string {
	switch e {
	case KIOServiceNotificationTypeLast:
		return "KIOServiceNotificationTypeLast"
	default:
		return fmt.Sprintf("KIOServiceNotificationType(%d)", e)
	}
}

type KIOServicePowerCapabilityOff uint

const (
	KIOServicePowerCapabilityLow KIOServicePowerCapabilityOff = 0
)

func (e KIOServicePowerCapabilityOff) String() string {
	switch e {
	case KIOServicePowerCapabilityLow:
		return "KIOServicePowerCapabilityLow"
	default:
		return fmt.Sprintf("KIOServicePowerCapabilityOff(%d)", e)
	}
}

const KIOServiceSearchPropertyParents uint = 0

const KIOServiceSeize uint = 0

type KIOStorageAccess uint

const (
	// KIOStorageAccessExclusiveLock: # Discussion
	KIOStorageAccessExclusiveLock KIOStorageAccess = 0
	KIOStorageAccessInvalid KIOStorageAccess = 0
)

func (e KIOStorageAccess) String() string {
	switch e {
	case KIOStorageAccessExclusiveLock:
		return "KIOStorageAccessExclusiveLock"
	default:
		return fmt.Sprintf("KIOStorageAccess(%d)", e)
	}
}

type KIOStorageOption uint

const (
	// KIOStorageOptionForceUnitAccess: # Discussion
	KIOStorageOptionForceUnitAccess KIOStorageOption = 0
	// KIOStorageOptionIsStatic: # Discussion
	KIOStorageOptionIsStatic KIOStorageOption = 0
)

func (e KIOStorageOption) String() string {
	switch e {
	case KIOStorageOptionForceUnitAccess:
		return "KIOStorageOptionForceUnitAccess"
	default:
		return fmt.Sprintf("KIOStorageOption(%d)", e)
	}
}

type KIOStoragePriorityHigh uint

const (
	KIOStoragePriorityBackground KIOStoragePriorityHigh = 0
)

func (e KIOStoragePriorityHigh) String() string {
	switch e {
	case KIOStoragePriorityBackground:
		return "KIOStoragePriorityBackground"
	default:
		return fmt.Sprintf("KIOStoragePriorityHigh(%d)", e)
	}
}

type KIOStorageProvisionTypeMapped uint

const (
	KIOStorageProvisionTypeAnchored KIOStorageProvisionTypeMapped = 0
)

func (e KIOStorageProvisionTypeMapped) String() string {
	switch e {
	case KIOStorageProvisionTypeAnchored:
		return "KIOStorageProvisionTypeAnchored"
	default:
		return fmt.Sprintf("KIOStorageProvisionTypeMapped(%d)", e)
	}
}

type KIOStorageSynchronizeOptionNone uint

const (
	KIOStorageSynchronizeOptionBarrier KIOStorageSynchronizeOptionNone = 0
)

func (e KIOStorageSynchronizeOptionNone) String() string {
	switch e {
	case KIOStorageSynchronizeOptionBarrier:
		return "KIOStorageSynchronizeOptionBarrier"
	default:
		return fmt.Sprintf("KIOStorageSynchronizeOptionNone(%d)", e)
	}
}

type KIOStorageUnmapOption uint

const (
	KIOStorageUnmapOptionNone KIOStorageUnmapOption = 0
	KIOStorageUnmapOptionWhole KIOStorageUnmapOption = 0
)

func (e KIOStorageUnmapOption) String() string {
	switch e {
	case KIOStorageUnmapOptionNone:
		return "KIOStorageUnmapOptionNone"
	default:
		return fmt.Sprintf("KIOStorageUnmapOption(%d)", e)
	}
}

type KIOStreamEnqueueInputTrap uint

const (
	KIOStreamEnqueueInputSyncTrap KIOStreamEnqueueInputTrap = 0
)

func (e KIOStreamEnqueueInputTrap) String() string {
	switch e {
	case KIOStreamEnqueueInputSyncTrap:
		return "KIOStreamEnqueueInputSyncTrap"
	default:
		return fmt.Sprintf("KIOStreamEnqueueInputTrap(%d)", e)
	}
}

type KIOStreamMemoryTypeOutputQueue uint

const (
	KIOStreamBufferIDMask KIOStreamMemoryTypeOutputQueue = 0
)

func (e KIOStreamMemoryTypeOutputQueue) String() string {
	switch e {
	case KIOStreamBufferIDMask:
		return "KIOStreamBufferIDMask"
	default:
		return fmt.Sprintf("KIOStreamMemoryTypeOutputQueue(%d)", e)
	}
}

type KIOStreamMethod uint

const (
	KIOStreamMethodSetMode KIOStreamMethod = 0
	KIOStreamMethodStop KIOStreamMethod = 0
)

func (e KIOStreamMethod) String() string {
	switch e {
	case KIOStreamMethodSetMode:
		return "KIOStreamMethodSetMode"
	default:
		return fmt.Sprintf("KIOStreamMethod(%d)", e)
	}
}

type KIOStreamMode uint

const (
	KIOStreamModeInput KIOStreamMode = 0
	KIOStreamModeInputOutput KIOStreamMode = 0
	KIOStreamModeOutput KIOStreamMode = 0
)

func (e KIOStreamMode) String() string {
	switch e {
	case KIOStreamModeInput:
		return "KIOStreamModeInput"
	default:
		return fmt.Sprintf("KIOStreamMode(%d)", e)
	}
}

type KIOStreamOptionOpenExclusive uint

const (
	KIOStreamOptionOpenShared KIOStreamOptionOpenExclusive = 0
)

func (e KIOStreamOptionOpenExclusive) String() string {
	switch e {
	case KIOStreamOptionOpenShared:
		return "KIOStreamOptionOpenShared"
	default:
		return fmt.Sprintf("KIOStreamOptionOpenExclusive(%d)", e)
	}
}

type KIOStreamPortTypeOutput uint

const (
	KIOStreamPortTypeInput KIOStreamPortTypeOutput = 0
)

func (e KIOStreamPortTypeOutput) String() string {
	switch e {
	case KIOStreamPortTypeInput:
		return "KIOStreamPortTypeInput"
	default:
		return fmt.Sprintf("KIOStreamPortTypeOutput(%d)", e)
	}
}

const KIOSyncPositivePolarity uint = 0

type KIOSystemStateSleepDescriptionHibernateStateInactive uint

const (
	KIOSystemStateSleepDescriptionHibernateStateHibernating KIOSystemStateSleepDescriptionHibernateStateInactive = 0
)

func (e KIOSystemStateSleepDescriptionHibernateStateInactive) String() string {
	switch e {
	case KIOSystemStateSleepDescriptionHibernateStateHibernating:
		return "KIOSystemStateSleepDescriptionHibernateStateHibernating"
	default:
		return fmt.Sprintf("KIOSystemStateSleepDescriptionHibernateStateInactive(%d)", e)
	}
}

type KIOTimeOptionsWithLeeway uint

const (
	KIOTimeOptionsContinuous KIOTimeOptionsWithLeeway = 0
)

func (e KIOTimeOptionsWithLeeway) String() string {
	switch e {
	case KIOTimeOptionsContinuous:
		return "KIOTimeOptionsContinuous"
	default:
		return fmt.Sprintf("KIOTimeOptionsWithLeeway(%d)", e)
	}
}

type KIOTimerEventSourceOptions uint

const (
	KIOTimerEventSourceOptionsAllowReenter KIOTimerEventSourceOptions = 0
	KIOTimerEventSourceOptionsDefault KIOTimerEventSourceOptions = 0
)

func (e KIOTimerEventSourceOptions) String() string {
	switch e {
	case KIOTimerEventSourceOptionsAllowReenter:
		return "KIOTimerEventSourceOptionsAllowReenter"
	default:
		return fmt.Sprintf("KIOTimerEventSourceOptions(%d)", e)
	}
}

type KIOTimingID uint

const (
	KIOTimingIDAppleNTSC_FFconv KIOTimingID = 0
	KIOTimingIDVESA_848x480_60hz KIOTimingID = 0
)

func (e KIOTimingID) String() string {
	switch e {
	case KIOTimingIDAppleNTSC_FFconv:
		return "KIOTimingIDAppleNTSC_FFconv"
	default:
		return fmt.Sprintf("KIOTimingID(%d)", e)
	}
}

type KIOTimingRangeV2 uint

const (
	KIOTimingRangeV1 KIOTimingRangeV2 = 0
)

func (e KIOTimingRangeV2) String() string {
	switch e {
	case KIOTimingRangeV1:
		return "KIOTimingRangeV1"
	default:
		return fmt.Sprintf("KIOTimingRangeV2(%d)", e)
	}
}

type KIOTrace uint

const (
	KIOTraceCommandGates KIOTrace = 0
	KIOTraceTimers KIOTrace = 0
)

func (e KIOTrace) String() string {
	switch e {
	case KIOTraceCommandGates:
		return "KIOTraceCommandGates"
	default:
		return fmt.Sprintf("KIOTrace(%d)", e)
	}
}

const KIOTrackingCallSiteBTs uint = 0

const KIOTrackingExcludeNames uint = 0

type KIOTrackingGetTracking uint

const (
	KIOTrackingGetMappings KIOTrackingGetTracking = 0
)

func (e KIOTrackingGetTracking) String() string {
	switch e {
	case KIOTrackingGetMappings:
		return "KIOTrackingGetMappings"
	default:
		return fmt.Sprintf("KIOTrackingGetTracking(%d)", e)
	}
}

type KIOTrackingLeakScanStart uint

const (
	KIOTrackingLeakScanEnd KIOTrackingLeakScanStart = 0
)

func (e KIOTrackingLeakScanStart) String() string {
	switch e {
	case KIOTrackingLeakScanEnd:
		return "KIOTrackingLeakScanEnd"
	default:
		return fmt.Sprintf("KIOTrackingLeakScanStart(%d)", e)
	}
}

type KIOUC uint

const (
	KIOUCScalarIStructO KIOUC = 0
	KIOUCStructIStructO KIOUC = 0
)

func (e KIOUC) String() string {
	switch e {
	case KIOUCScalarIStructO:
		return "KIOUCScalarIStructO"
	default:
		return fmt.Sprintf("KIOUC(%d)", e)
	}
}

const KIOUCVariableStructureSize uint = 0

type KIOUSB uint

const (
	// KIOUSBDecriptorTypeHID: The type for a human interaction device descriptor.
	KIOUSBDecriptorTypeHID KIOUSB = 0
	// KIOUSBDecriptorTypeReport: The type for a report descriptor.
	KIOUSBDecriptorTypeReport KIOUSB = 0
	KIOUSBDescriptorTypeBOS KIOUSB = 0
	KIOUSBDescriptorTypeConfiguration KIOUSB = 0
	KIOUSBDescriptorTypeDebug KIOUSB = 0
	KIOUSBDescriptorTypeDevice KIOUSB = 0
	KIOUSBDescriptorTypeDeviceCapability KIOUSB = 0
	KIOUSBDescriptorTypeDeviceQualifier KIOUSB = 0
	KIOUSBDescriptorTypeEndpoint KIOUSB = 0
	KIOUSBDescriptorTypeHub KIOUSB = 0
	KIOUSBDescriptorTypeInterface KIOUSB = 0
	KIOUSBDescriptorTypeInterfaceAssociation KIOUSB = 0
	KIOUSBDescriptorTypeInterfacePower KIOUSB = 0
	KIOUSBDescriptorTypeOTG KIOUSB = 0
	KIOUSBDescriptorTypeOtherSpeedConfiguration KIOUSB = 0
	// KIOUSBDescriptorTypePhysical: The type for a physical descriptor.
	KIOUSBDescriptorTypePhysical KIOUSB = 0
	KIOUSBDescriptorTypeString KIOUSB = 0
	KIOUSBDescriptorTypeSuperSpeedHub KIOUSB = 0
	KIOUSBDescriptorTypeSuperSpeedPlusIsochronousEndpointCompanion KIOUSB = 0
	KIOUSBDescriptorTypeSuperSpeedUSBEndpointCompanion KIOUSB = 0
)

func (e KIOUSB) String() string {
	switch e {
	case KIOUSBDecriptorTypeHID:
		return "KIOUSBDecriptorTypeHID"
	default:
		return fmt.Sprintf("KIOUSB(%d)", e)
	}
}

type KIOUSB20BusCurrent uint

const (
	// KIOUSB20BusCurrentDefault: The default bus current.
	KIOUSB20BusCurrentDefault KIOUSB20BusCurrent = 0
	// KIOUSB20BusCurrentMaxPowerUnits: The maximum current power units.
	KIOUSB20BusCurrentMaxPowerUnits KIOUSB20BusCurrent = 0
	// KIOUSB20BusCurrentMinimum: The minimum current.
	KIOUSB20BusCurrentMinimum KIOUSB20BusCurrent = 0
)

func (e KIOUSB20BusCurrent) String() string {
	switch e {
	case KIOUSB20BusCurrentDefault:
		return "KIOUSB20BusCurrentDefault"
	default:
		return fmt.Sprintf("KIOUSB20BusCurrent(%d)", e)
	}
}

type KIOUSB30 uint

const (
	// KIOUSB30LinKStateU2RxDetectDelay: The USB 2.0 receive detect delay.
	KIOUSB30LinKStateU2RxDetectDelay KIOUSB30 = 0
	// KIOUSB30LinkStateHotResetActiveTimeout: The hot reset active timeout.
	KIOUSB30LinkStateHotResetActiveTimeout KIOUSB30 = 0
	KIOUSB30LinkStateHotResetDeadline KIOUSB30 = 0
	KIOUSB30LinkStateHotResetExitTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateLoopbackExitTimeout: The loopback exit timeout.
	KIOUSB30LinkStateLoopbackExitTimeout KIOUSB30 = 0
	// KIOUSB30LinkStatePollingActiveTimeout: The polling active timeout.
	KIOUSB30LinkStatePollingActiveTimeout KIOUSB30 = 0
	// KIOUSB30LinkStatePollingConfigurationTimeout: The polling configuration timeout.
	KIOUSB30LinkStatePollingConfigurationTimeout KIOUSB30 = 0
	// KIOUSB30LinkStatePollingDeadline: The polling deadline.
	KIOUSB30LinkStatePollingDeadline KIOUSB30 = 0
	// KIOUSB30LinkStatePollingIdleTimeout: The polling idle timeout.
	KIOUSB30LinkStatePollingIdleTimeout KIOUSB30 = 0
	// KIOUSB30LinkStatePollingLFPSTimeout: The polling LFPS timeout.
	KIOUSB30LinkStatePollingLFPSTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateRecoveryActiveTimeout: The recovery active timeout.
	KIOUSB30LinkStateRecoveryActiveTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateRecoveryConfigurationTimeout: The recovery configuration timeout.
	KIOUSB30LinkStateRecoveryConfigurationTimeout KIOUSB30 = 0
	KIOUSB30LinkStateRecoveryDeadline KIOUSB30 = 0
	// KIOUSB30LinkStateRecoveryIdleTimeout: The recovery idle timeout.
	KIOUSB30LinkStateRecoveryIdleTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateRxDetectQuietTimeout: The receive detect quiet timeout.
	KIOUSB30LinkStateRxDetectQuietTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateSSInactiveQuietTimeout: The SuperSpeed inactive quiet timeout.
	KIOUSB30LinkStateSSInactiveQuietTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateSSResumeDeadline: The SuperSpeed resume deadline.
	KIOUSB30LinkStateSSResumeDeadline KIOUSB30 = 0
	// KIOUSB30LinkStateU0LTimeout: The U0L timeout.
	KIOUSB30LinkStateU0LTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateU0RecoveryTimeout: The U0 recovery timeout.
	KIOUSB30LinkStateU0RecoveryTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateU1NoLFPSResponseTimeout: The U1 no LFPS response timeout.
	KIOUSB30LinkStateU1NoLFPSResponseTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateU1PingTimeout: The U1 ping timeout.
	KIOUSB30LinkStateU1PingTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateU2NoLFPSResponseTimeout: The U2 no LFPS response timeout.
	KIOUSB30LinkStateU2NoLFPSResponseTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateU3NoLFPSResponseTimeout: The U3 no LFPS response timeout.
	KIOUSB30LinkStateU3NoLFPSResponseTimeout KIOUSB30 = 0
	// KIOUSB30LinkStateU3RxDetectDelay: The U3 receive detect delay.
	KIOUSB30LinkStateU3RxDetectDelay KIOUSB30 = 0
	// KIOUSB30LinkStateU3WakeupRetryDelay: The U3 wakeup retry delay.
	KIOUSB30LinkStateU3WakeupRetryDelay KIOUSB30 = 0
)

func (e KIOUSB30) String() string {
	switch e {
	case KIOUSB30LinKStateU2RxDetectDelay:
		return "KIOUSB30LinKStateU2RxDetectDelay"
	default:
		return fmt.Sprintf("KIOUSB30(%d)", e)
	}
}

type KIOUSB30BusCurrent uint

const (
	// KIOUSB30BusCurrentDefault: The default bus current.
	KIOUSB30BusCurrentDefault KIOUSB30BusCurrent = 0
	// KIOUSB30BusCurrentMaxPowerUnits: The maximum current power units.
	KIOUSB30BusCurrentMaxPowerUnits KIOUSB30BusCurrent = 0
	// KIOUSB30BusCurrentMinimum: The minimum current.
	KIOUSB30BusCurrentMinimum KIOUSB30BusCurrent = 0
)

func (e KIOUSB30BusCurrent) String() string {
	switch e {
	case KIOUSB30BusCurrentDefault:
		return "KIOUSB30BusCurrentDefault"
	default:
		return fmt.Sprintf("KIOUSB30BusCurrent(%d)", e)
	}
}

type KIOUSB30DeviceNotificationType uint

const (
	// KIOUSB30DeviceNotificationTypeBusIntervalAdjustment: The bus interval adjustment.
	KIOUSB30DeviceNotificationTypeBusIntervalAdjustment KIOUSB30DeviceNotificationType = 0
	// KIOUSB30DeviceNotificationTypeFunctionWake: The function wake flag.
	KIOUSB30DeviceNotificationTypeFunctionWake KIOUSB30DeviceNotificationType = 0
	// KIOUSB30DeviceNotificationTypeHostRoleRequest: The host role request.
	KIOUSB30DeviceNotificationTypeHostRoleRequest KIOUSB30DeviceNotificationType = 0
	// KIOUSB30DeviceNotificationTypeLatencyTolerance: The latency tolerance.
	KIOUSB30DeviceNotificationTypeLatencyTolerance KIOUSB30DeviceNotificationType = 0
	// KIOUSB30DeviceNotificationTypeSublinkSpeed: The sublink speed.
	KIOUSB30DeviceNotificationTypeSublinkSpeed KIOUSB30DeviceNotificationType = 0
)

func (e KIOUSB30DeviceNotificationType) String() string {
	switch e {
	case KIOUSB30DeviceNotificationTypeBusIntervalAdjustment:
		return "KIOUSB30DeviceNotificationTypeBusIntervalAdjustment"
	default:
		return fmt.Sprintf("KIOUSB30DeviceNotificationType(%d)", e)
	}
}

type KIOUSB30HubExtStatus int

const (
	// KIOUSB30HubExtStatusRxLaneCount: The receive lane count.
	KIOUSB30HubExtStatusRxLaneCount KIOUSB30HubExtStatus = 0
	// KIOUSB30HubExtStatusRxLaneCountPhase: The receive lane count phase.
	KIOUSB30HubExtStatusRxLaneCountPhase KIOUSB30HubExtStatus = 0
	// KIOUSB30HubExtStatusRxSublinkSpeedID: The receive sublink speed identifier.
	KIOUSB30HubExtStatusRxSublinkSpeedID KIOUSB30HubExtStatus = 0
	// KIOUSB30HubExtStatusRxSublinkSpeedIDPhase: The receive sublink speed identifier phase.
	KIOUSB30HubExtStatusRxSublinkSpeedIDPhase KIOUSB30HubExtStatus = 0
	// KIOUSB30HubExtStatusTxLaneCount: The transmit lane count.
	KIOUSB30HubExtStatusTxLaneCount KIOUSB30HubExtStatus = 0
	// KIOUSB30HubExtStatusTxLaneCountPhase: The transmit lane count phase.
	KIOUSB30HubExtStatusTxLaneCountPhase KIOUSB30HubExtStatus = 0
	// KIOUSB30HubExtStatusTxSublinkSpeedID: The transmit sublink speed identifier.
	KIOUSB30HubExtStatusTxSublinkSpeedID KIOUSB30HubExtStatus = 0
	// KIOUSB30HubExtStatusTxSublinkSpeedIDPhase: The transmit sublink speed identifier phase.
	KIOUSB30HubExtStatusTxSublinkSpeedIDPhase KIOUSB30HubExtStatus = 0
)

func (e KIOUSB30HubExtStatus) String() string {
	switch e {
	case KIOUSB30HubExtStatusRxLaneCount:
		return "KIOUSB30HubExtStatusRxLaneCount"
	default:
		return fmt.Sprintf("KIOUSB30HubExtStatus(%d)", e)
	}
}

type KIOUSB30HubPortStatusCode int

const (
	// KIOUSB30HubPortStatusCodeCount: The port status code count.
	KIOUSB30HubPortStatusCodeCount KIOUSB30HubPortStatusCode = 0
	// KIOUSB30HubPortStatusCodeExt: The port status code extension.
	KIOUSB30HubPortStatusCodeExt KIOUSB30HubPortStatusCode = 0
	// KIOUSB30HubPortStatusCodePD: The port status code.
	KIOUSB30HubPortStatusCodePD KIOUSB30HubPortStatusCode = 0
	// KIOUSB30HubPortStatusCodeStandard: The hub standard status code.
	KIOUSB30HubPortStatusCodeStandard KIOUSB30HubPortStatusCode = 0
)

func (e KIOUSB30HubPortStatusCode) String() string {
	switch e {
	case KIOUSB30HubPortStatusCodeCount:
		return "KIOUSB30HubPortStatusCodeCount"
	default:
		return fmt.Sprintf("KIOUSB30HubPortStatusCode(%d)", e)
	}
}

type KIOUSB30Reset uint

const (
	// KIOUSB30ResetMaximumTimeout: The maximum reset timeout.
	KIOUSB30ResetMaximumTimeout KIOUSB30Reset = 0
	// KIOUSB30ResetMaximumWithMarginTimeout: The maximum reset timeout with margin.
	KIOUSB30ResetMaximumWithMarginTimeout KIOUSB30Reset = 0
	// KIOUSB30ResetMinimumTimeout: The minimum reset timeout.
	KIOUSB30ResetMinimumTimeout KIOUSB30Reset = 0
	// KIOUSB30ResetTypicalTimeout: The typical reset timeout.
	KIOUSB30ResetTypicalTimeout KIOUSB30Reset = 0
)

func (e KIOUSB30Reset) String() string {
	switch e {
	case KIOUSB30ResetMaximumTimeout:
		return "KIOUSB30ResetMaximumTimeout"
	default:
		return fmt.Sprintf("KIOUSB30Reset(%d)", e)
	}
}

type KIOUSB30TimingParameterBELT uint

const (
	// KIOUSB30TimingParameterBELTDefaultNs: The BELT default in nanoseconds.
	KIOUSB30TimingParameterBELTDefaultNs KIOUSB30TimingParameterBELT = 0
	// KIOUSB30TimingParameterBELTMinNs: The BELT minimum in nanoseconds.
	KIOUSB30TimingParameterBELTMinNs KIOUSB30TimingParameterBELT = 0
)

func (e KIOUSB30TimingParameterBELT) String() string {
	switch e {
	case KIOUSB30TimingParameterBELTDefaultNs:
		return "KIOUSB30TimingParameterBELTDefaultNs"
	default:
		return fmt.Sprintf("KIOUSB30TimingParameterBELT(%d)", e)
	}
}

const KIOUSBAnyClass uint = 0

const KIOUSBAppleVendorID uint = 0

type KIOUSBConfigurationDescriptorAttribute uint

const (
	KIOUSBConfigurationDescriptorAttributeRemoteWakeCapable KIOUSBConfigurationDescriptorAttribute = 0
	KIOUSBConfigurationDescriptorAttributeSelfPowered KIOUSBConfigurationDescriptorAttribute = 0
)

func (e KIOUSBConfigurationDescriptorAttribute) String() string {
	switch e {
	case KIOUSBConfigurationDescriptorAttributeRemoteWakeCapable:
		return "KIOUSBConfigurationDescriptorAttributeRemoteWakeCapable"
	default:
		return fmt.Sprintf("KIOUSBConfigurationDescriptorAttribute(%d)", e)
	}
}

type KIOUSBDescriptor uint

const (
	KIOUSBDescriptorHeaderSize KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeBOS KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeBillboardDeviceMaximum KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeBillboardDeviceMinimum KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeConfiguration KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeContainerIDCapability KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeDevice KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeDeviceCapability KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeDeviceQualifier KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeEndpoint KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeHubMaximum KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeHubMinimum KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeInterface KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeInterfaceAssociation KIOUSBDescriptor = 0
	KIOUSBDescriptorSizePlatformCapability KIOUSBDescriptor = 0
	KIOUSBDescriptorSizePlatformECIDCapability KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeStringMaximum KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeStringMinimum KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeSuperSpeedHub KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeSuperSpeedPlusIsochronousEndpointCompanion KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeSuperSpeedUSBDeviceCapability KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeSuperSpeedUSBEndpointCompanion KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeUSB20ExtensionCapability KIOUSBDescriptor = 0
)

func (e KIOUSBDescriptor) String() string {
	switch e {
	case KIOUSBDescriptorHeaderSize:
		return "KIOUSBDescriptorHeaderSize"
	default:
		return fmt.Sprintf("KIOUSBDescriptor(%d)", e)
	}
}

type KIOUSBDeviceCapabilityDescriptorType uint

const (
	KIOUSBDeviceCapabilityDescriptorLengthMin KIOUSBDeviceCapabilityDescriptorType = 0
)

func (e KIOUSBDeviceCapabilityDescriptorType) String() string {
	switch e {
	case KIOUSBDeviceCapabilityDescriptorLengthMin:
		return "KIOUSBDeviceCapabilityDescriptorLengthMin"
	default:
		return fmt.Sprintf("KIOUSBDeviceCapabilityDescriptorType(%d)", e)
	}
}

type KIOUSBDeviceCapabilityType uint

const (
	KIOUSBDeviceCapabilityTypeBatteryInfo KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeBillboard KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeBillboardAltMode KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeContainerID KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePdConsumerPort KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePdProviderPort KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePlatform KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePowerDelivery KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePrecisionMeasurement KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeSuperSpeed KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeSuperSpeedPlus KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeUSB20Extension KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeWireless KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeWirelessExt KIOUSBDeviceCapabilityType = 0
)

func (e KIOUSBDeviceCapabilityType) String() string {
	switch e {
	case KIOUSBDeviceCapabilityTypeBatteryInfo:
		return "KIOUSBDeviceCapabilityTypeBatteryInfo"
	default:
		return fmt.Sprintf("KIOUSBDeviceCapabilityType(%d)", e)
	}
}

type KIOUSBDeviceFeatureSelector uint

const (
	KIOUSBDeviceFeatureSelectorLTMEnable KIOUSBDeviceFeatureSelector = 0
	KIOUSBDeviceFeatureSelectorU1Enable KIOUSBDeviceFeatureSelector = 0
)

func (e KIOUSBDeviceFeatureSelector) String() string {
	switch e {
	case KIOUSBDeviceFeatureSelectorLTMEnable:
		return "KIOUSBDeviceFeatureSelectorLTMEnable"
	default:
		return fmt.Sprintf("KIOUSBDeviceFeatureSelector(%d)", e)
	}
}

type KIOUSBDeviceRequest uint

const (
	// KIOUSBDeviceRequestDirectionIn: The incoming direction.
	KIOUSBDeviceRequestDirectionIn KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestDirectionMask: The direction mask.
	KIOUSBDeviceRequestDirectionMask KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestDirectionOut: The outgoing direction.
	KIOUSBDeviceRequestDirectionOut KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestDirectionPhase: The direction phase.
	KIOUSBDeviceRequestDirectionPhase KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestRecipientDevice: The recipient device.
	KIOUSBDeviceRequestRecipientDevice KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestRecipientEndpoint: The recipient endpoint.
	KIOUSBDeviceRequestRecipientEndpoint KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestRecipientInterface: The recipient interface.
	KIOUSBDeviceRequestRecipientInterface KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestRecipientMask: The recipient mask.
	KIOUSBDeviceRequestRecipientMask KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestRecipientOther: The other recipient.
	KIOUSBDeviceRequestRecipientOther KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestRecipientPhase: The recipient phase.
	KIOUSBDeviceRequestRecipientPhase KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestSize: The request size.
	KIOUSBDeviceRequestSize KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestTypeClass: The type class.
	KIOUSBDeviceRequestTypeClass KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestTypeMask: The type mask.
	KIOUSBDeviceRequestTypeMask KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestTypePhase: The type phase.
	KIOUSBDeviceRequestTypePhase KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestTypeStandard: The type standard.
	KIOUSBDeviceRequestTypeStandard KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestTypeVendor: The type vendor.
	KIOUSBDeviceRequestTypeVendor KIOUSBDeviceRequest = 0
)

func (e KIOUSBDeviceRequest) String() string {
	switch e {
	case KIOUSBDeviceRequestDirectionIn:
		return "KIOUSBDeviceRequestDirectionIn"
	default:
		return fmt.Sprintf("KIOUSBDeviceRequest(%d)", e)
	}
}

type KIOUSBDeviceRequestDirectionValue uint

const (
	// KIOUSBDeviceRequestDirectionValueIn: The input direction value.
	KIOUSBDeviceRequestDirectionValueIn KIOUSBDeviceRequestDirectionValue = 0
	// KIOUSBDeviceRequestDirectionValueOut: The output direction value.
	KIOUSBDeviceRequestDirectionValueOut KIOUSBDeviceRequestDirectionValue = 0
)

func (e KIOUSBDeviceRequestDirectionValue) String() string {
	switch e {
	case KIOUSBDeviceRequestDirectionValueIn:
		return "KIOUSBDeviceRequestDirectionValueIn"
	default:
		return fmt.Sprintf("KIOUSBDeviceRequestDirectionValue(%d)", e)
	}
}

type KIOUSBDeviceRequestGet uint

const (
	KIOUSBDeviceRequestGetConfiguration KIOUSBDeviceRequestGet = 0
	KIOUSBDeviceRequestGetStatus KIOUSBDeviceRequestGet = 0
)

func (e KIOUSBDeviceRequestGet) String() string {
	switch e {
	case KIOUSBDeviceRequestGetConfiguration:
		return "KIOUSBDeviceRequestGetConfiguration"
	default:
		return fmt.Sprintf("KIOUSBDeviceRequestGet(%d)", e)
	}
}

type KIOUSBDeviceRequestRecipientValue uint

const (
	// KIOUSBDeviceRequestRecipientValueDevice: The recipient device value.
	KIOUSBDeviceRequestRecipientValueDevice KIOUSBDeviceRequestRecipientValue = 0
	// KIOUSBDeviceRequestRecipientValueEndpoint: The recipient endpoint value.
	KIOUSBDeviceRequestRecipientValueEndpoint KIOUSBDeviceRequestRecipientValue = 0
	// KIOUSBDeviceRequestRecipientValueInterface: The recipient interface value.
	KIOUSBDeviceRequestRecipientValueInterface KIOUSBDeviceRequestRecipientValue = 0
	// KIOUSBDeviceRequestRecipientValueOther: The other recipient value.
	KIOUSBDeviceRequestRecipientValueOther KIOUSBDeviceRequestRecipientValue = 0
)

func (e KIOUSBDeviceRequestRecipientValue) String() string {
	switch e {
	case KIOUSBDeviceRequestRecipientValueDevice:
		return "KIOUSBDeviceRequestRecipientValueDevice"
	default:
		return fmt.Sprintf("KIOUSBDeviceRequestRecipientValue(%d)", e)
	}
}

type KIOUSBDeviceRequestTypeValue uint

const (
	// KIOUSBDeviceRequestTypeValueClass: The class type value.
	KIOUSBDeviceRequestTypeValueClass KIOUSBDeviceRequestTypeValue = 0
	// KIOUSBDeviceRequestTypeValueStandard: The standard type value.
	KIOUSBDeviceRequestTypeValueStandard KIOUSBDeviceRequestTypeValue = 0
	// KIOUSBDeviceRequestTypeValueVendor: The vendor type value.
	KIOUSBDeviceRequestTypeValueVendor KIOUSBDeviceRequestTypeValue = 0
)

func (e KIOUSBDeviceRequestTypeValue) String() string {
	switch e {
	case KIOUSBDeviceRequestTypeValueClass:
		return "KIOUSBDeviceRequestTypeValueClass"
	default:
		return fmt.Sprintf("KIOUSBDeviceRequestTypeValue(%d)", e)
	}
}

type KIOUSBDeviceStatusSelfPowered int

const (
	IOUSBEndpointStatusHalt KIOUSBDeviceStatusSelfPowered = 0
	KIOUSBDeviceStatusRemoteWakeEnable KIOUSBDeviceStatusSelfPowered = 0
)

func (e KIOUSBDeviceStatusSelfPowered) String() string {
	switch e {
	case IOUSBEndpointStatusHalt:
		return "IOUSBEndpointStatusHalt"
	default:
		return fmt.Sprintf("KIOUSBDeviceStatusSelfPowered(%d)", e)
	}
}

type KIOUSBEndpointDescriptor uint

const (
	KIOUSBEndpointDescriptorNumber KIOUSBEndpointDescriptor = 0
	KIOUSBEndpointDescriptorSynchronizationTypePhase KIOUSBEndpointDescriptor = 0
)

func (e KIOUSBEndpointDescriptor) String() string {
	switch e {
	case KIOUSBEndpointDescriptorNumber:
		return "KIOUSBEndpointDescriptorNumber"
	default:
		return fmt.Sprintf("KIOUSBEndpointDescriptor(%d)", e)
	}
}

type KIOUSBEndpointDirection uint

const (
	KIOUSBEndpointDirectionIn KIOUSBEndpointDirection = 0
	KIOUSBEndpointDirectionOut KIOUSBEndpointDirection = 0
	KIOUSBEndpointDirectionUnknown KIOUSBEndpointDirection = 0
)

func (e KIOUSBEndpointDirection) String() string {
	switch e {
	case KIOUSBEndpointDirectionIn:
		return "KIOUSBEndpointDirectionIn"
	default:
		return fmt.Sprintf("KIOUSBEndpointDirection(%d)", e)
	}
}

type KIOUSBEndpointSynchronizationType uint

const (
	// KIOUSBEndpointSynchronizationTypeAdaptive: The adaptive type.
	KIOUSBEndpointSynchronizationTypeAdaptive KIOUSBEndpointSynchronizationType = 0
	// KIOUSBEndpointSynchronizationTypeAsynchronous: The asynchronous type.
	KIOUSBEndpointSynchronizationTypeAsynchronous KIOUSBEndpointSynchronizationType = 0
	// KIOUSBEndpointSynchronizationTypeNone: No synchronization type.
	KIOUSBEndpointSynchronizationTypeNone KIOUSBEndpointSynchronizationType = 0
	// KIOUSBEndpointSynchronizationTypeSynchronous: The synchronous type.
	KIOUSBEndpointSynchronizationTypeSynchronous KIOUSBEndpointSynchronizationType = 0
)

func (e KIOUSBEndpointSynchronizationType) String() string {
	switch e {
	case KIOUSBEndpointSynchronizationTypeAdaptive:
		return "KIOUSBEndpointSynchronizationTypeAdaptive"
	default:
		return fmt.Sprintf("KIOUSBEndpointSynchronizationType(%d)", e)
	}
}

type KIOUSBEndpointType uint

const (
	KIOUSBEndpointTypeBulk KIOUSBEndpointType = 0
	KIOUSBEndpointTypeControl KIOUSBEndpointType = 0
	KIOUSBEndpointTypeInterrupt KIOUSBEndpointType = 0
	KIOUSBEndpointTypeIsochronous KIOUSBEndpointType = 0
)

func (e KIOUSBEndpointType) String() string {
	switch e {
	case KIOUSBEndpointTypeBulk:
		return "KIOUSBEndpointTypeBulk"
	default:
		return fmt.Sprintf("KIOUSBEndpointType(%d)", e)
	}
}

type KIOUSBEndpointUsageTypeIsoc uint

const (
	// KIOUSBEndpointUsageTypeIsocData: The isochronous data type.
	KIOUSBEndpointUsageTypeIsocData KIOUSBEndpointUsageTypeIsoc = 0
	// KIOUSBEndpointUsageTypeIsocFeedback: The isochronous feedback type.
	KIOUSBEndpointUsageTypeIsocFeedback KIOUSBEndpointUsageTypeIsoc = 0
	// KIOUSBEndpointUsageTypeIsocImplicit: The isochronous implicit type.
	KIOUSBEndpointUsageTypeIsocImplicit KIOUSBEndpointUsageTypeIsoc = 0
)

func (e KIOUSBEndpointUsageTypeIsoc) String() string {
	switch e {
	case KIOUSBEndpointUsageTypeIsocData:
		return "KIOUSBEndpointUsageTypeIsocData"
	default:
		return fmt.Sprintf("KIOUSBEndpointUsageTypeIsoc(%d)", e)
	}
}

const KIOUSBFindInterfaceDontCare uint = 0

type KIOUSBGetEndpointDescriptor uint

const (
	// KIOUSBGetEndpointDescriptorCurrentPolicy: The descriptor controlling the current endpoint policy.
	KIOUSBGetEndpointDescriptorCurrentPolicy KIOUSBGetEndpointDescriptor = 0
	// KIOUSBGetEndpointDescriptorOriginal: The original descriptor that the system uses to create the pipe.
	KIOUSBGetEndpointDescriptorOriginal KIOUSBGetEndpointDescriptor = 0
)

func (e KIOUSBGetEndpointDescriptor) String() string {
	switch e {
	case KIOUSBGetEndpointDescriptorCurrentPolicy:
		return "KIOUSBGetEndpointDescriptorCurrentPolicy"
	default:
		return fmt.Sprintf("KIOUSBGetEndpointDescriptor(%d)", e)
	}
}

type KIOUSBHostConnectionSpeed uint

const (
	// KIOUSBHostConnectionSpeedCount: The speed count.
	KIOUSBHostConnectionSpeedCount KIOUSBHostConnectionSpeed = 0
	// KIOUSBHostConnectionSpeedFull: The full speed value.
	KIOUSBHostConnectionSpeedFull KIOUSBHostConnectionSpeed = 0
	// KIOUSBHostConnectionSpeedHigh: The high speed value.
	KIOUSBHostConnectionSpeedHigh KIOUSBHostConnectionSpeed = 0
	// KIOUSBHostConnectionSpeedLow: The low speed value.
	KIOUSBHostConnectionSpeedLow KIOUSBHostConnectionSpeed = 0
	// KIOUSBHostConnectionSpeedNone: The no speed value.
	KIOUSBHostConnectionSpeedNone KIOUSBHostConnectionSpeed = 0
	// KIOUSBHostConnectionSpeedSuper: The SuperSpeed value.
	KIOUSBHostConnectionSpeedSuper KIOUSBHostConnectionSpeed = 0
	// KIOUSBHostConnectionSpeedSuperPlus: The SuperSpeedPlus value.
	KIOUSBHostConnectionSpeedSuperPlus KIOUSBHostConnectionSpeed = 0
	// KIOUSBHostConnectionSpeedSuperPlusBy2: The SuperSpeedPlus By 2 value.
	KIOUSBHostConnectionSpeedSuperPlusBy2 KIOUSBHostConnectionSpeed = 0
)

func (e KIOUSBHostConnectionSpeed) String() string {
	switch e {
	case KIOUSBHostConnectionSpeedCount:
		return "KIOUSBHostConnectionSpeedCount"
	default:
		return fmt.Sprintf("KIOUSBHostConnectionSpeed(%d)", e)
	}
}

const KIOUSBHostPipeBundlingMax uint = 0

type KIOUSBHostPortStatus int

const (
	// KIOUSBHostPortStatusConnectedSpeedFull: The full connected speed.
	KIOUSBHostPortStatusConnectedSpeedFull KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusConnectedSpeedHigh: The high connected speed.
	KIOUSBHostPortStatusConnectedSpeedHigh KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusConnectedSpeedLow: The low connected speed.
	KIOUSBHostPortStatusConnectedSpeedLow KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusConnectedSpeedMask: The connected speed mask.
	KIOUSBHostPortStatusConnectedSpeedMask KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusConnectedSpeedNone: No connected speed.
	KIOUSBHostPortStatusConnectedSpeedNone KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusConnectedSpeedPhase: The connected speed phase.
	KIOUSBHostPortStatusConnectedSpeedPhase KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusConnectedSpeedSuper: The SuperSpeed connection value.
	KIOUSBHostPortStatusConnectedSpeedSuper KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusConnectedSpeedSuperPlus: The SuperSpeedPlus connected value.
	KIOUSBHostPortStatusConnectedSpeedSuperPlus KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusConnectedSpeedSuperPlusBy2: The SuperSpeedPlus By 2 connected value.
	KIOUSBHostPortStatusConnectedSpeedSuperPlusBy2 KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusEnabled: The port status enabled value.
	KIOUSBHostPortStatusEnabled KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusOvercurrent: The port status overcurrent.
	KIOUSBHostPortStatusOvercurrent KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusPortTypeAccessory: The accessory port type.
	KIOUSBHostPortStatusPortTypeAccessory KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusPortTypeCaptive: The captive port type.
	KIOUSBHostPortStatusPortTypeCaptive KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusPortTypeInternal: The internal port type.
	KIOUSBHostPortStatusPortTypeInternal KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusPortTypeMask: The port type mask.
	KIOUSBHostPortStatusPortTypeMask KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusPortTypePhase: The port type phase.
	KIOUSBHostPortStatusPortTypePhase KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusPortTypeReserved: The reserved port type.
	KIOUSBHostPortStatusPortTypeReserved KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusPortTypeStandard: The standard port type.
	KIOUSBHostPortStatusPortTypeStandard KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusResetting: The resetting port status.
	KIOUSBHostPortStatusResetting KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusSuspended: The suspended port status.
	KIOUSBHostPortStatusSuspended KIOUSBHostPortStatus = 0
	// KIOUSBHostPortStatusTestMode: The test mode status.
	KIOUSBHostPortStatusTestMode KIOUSBHostPortStatus = 0
)

func (e KIOUSBHostPortStatus) String() string {
	switch e {
	case KIOUSBHostPortStatusConnectedSpeedFull:
		return "KIOUSBHostPortStatusConnectedSpeedFull"
	default:
		return fmt.Sprintf("KIOUSBHostPortStatus(%d)", e)
	}
}

type KIOUSBHostPortType uint

const (
	// KIOUSBHostPortTypeAccessory: The accessory type.
	KIOUSBHostPortTypeAccessory KIOUSBHostPortType = 0
	// KIOUSBHostPortTypeCaptive: The captive type.
	KIOUSBHostPortTypeCaptive KIOUSBHostPortType = 0
	// KIOUSBHostPortTypeCount: The type count.
	KIOUSBHostPortTypeCount KIOUSBHostPortType = 0
	// KIOUSBHostPortTypeExpressCard: The express card type.
	KIOUSBHostPortTypeExpressCard KIOUSBHostPortType = 0
	// KIOUSBHostPortTypeInternal: The internal type.
	KIOUSBHostPortTypeInternal KIOUSBHostPortType = 0
	// KIOUSBHostPortTypeStandard: The standard type.
	KIOUSBHostPortTypeStandard KIOUSBHostPortType = 0
)

func (e KIOUSBHostPortType) String() string {
	switch e {
	case KIOUSBHostPortTypeAccessory:
		return "KIOUSBHostPortTypeAccessory"
	default:
		return fmt.Sprintf("KIOUSBHostPortType(%d)", e)
	}
}

type KIOUSBHubPort2PortExitLatencyNs uint

const (
	KIOUSBHubDelayNs KIOUSBHubPort2PortExitLatencyNs = 0
)

func (e KIOUSBHubPort2PortExitLatencyNs) String() string {
	switch e {
	case KIOUSBHubDelayNs:
		return "KIOUSBHubDelayNs"
	default:
		return fmt.Sprintf("KIOUSBHubPort2PortExitLatencyNs(%d)", e)
	}
}

const KIOUSBInterfaceOpenAlt uint = 0

const KIOUSBInterfaceSuspendLowPower uint = 0

const KIOUSBPingResponseTimeNs uint = 0

type KIOUSBSuperSpeedDeviceCapability uint

const (
	KIOUSBSuperSpeedDeviceCapabilityHighSpeed KIOUSBSuperSpeedDeviceCapability = 0
	KIOUSBSuperSpeedDeviceCapabilitySupportHighSpeed KIOUSBSuperSpeedDeviceCapability = 0
)

func (e KIOUSBSuperSpeedDeviceCapability) String() string {
	switch e {
	case KIOUSBSuperSpeedDeviceCapabilityHighSpeed:
		return "KIOUSBSuperSpeedDeviceCapabilityHighSpeed"
	default:
		return fmt.Sprintf("KIOUSBSuperSpeedDeviceCapability(%d)", e)
	}
}

type KIOUSBSuperSpeedEndpointCompanionDescriptor uint

const (
	KIOUSBSuperSpeedEndpointCompanionDescriptorBulkMaxStreams KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorIsocReservedPhase KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
)

func (e KIOUSBSuperSpeedEndpointCompanionDescriptor) String() string {
	switch e {
	case KIOUSBSuperSpeedEndpointCompanionDescriptorBulkMaxStreams:
		return "KIOUSBSuperSpeedEndpointCompanionDescriptorBulkMaxStreams"
	default:
		return fmt.Sprintf("KIOUSBSuperSpeedEndpointCompanionDescriptor(%d)", e)
	}
}

type KIOUSBSuperSpeedHubCharacteristics uint

const (
	KIOUSBSuperSpeedHubCharacteristicsCompoundDevice KIOUSBSuperSpeedHubCharacteristics = 0
	KIOUSBSuperSpeedHubCharacteristicsOverCurrentGlobal KIOUSBSuperSpeedHubCharacteristics = 0
)

func (e KIOUSBSuperSpeedHubCharacteristics) String() string {
	switch e {
	case KIOUSBSuperSpeedHubCharacteristicsCompoundDevice:
		return "KIOUSBSuperSpeedHubCharacteristicsCompoundDevice"
	default:
		return fmt.Sprintf("KIOUSBSuperSpeedHubCharacteristics(%d)", e)
	}
}

type KIOUSBSuperSpeedPlusDeviceCapabilitySublink uint

const (
	KIOUSBSuperSpeedPlusDeviceCapabilitySublinkDirectionPhase KIOUSBSuperSpeedPlusDeviceCapabilitySublink = 0
	KIOUSBSuperSpeedPlusDeviceCapabilitySublinkSymmetric KIOUSBSuperSpeedPlusDeviceCapabilitySublink = 0
)

func (e KIOUSBSuperSpeedPlusDeviceCapabilitySublink) String() string {
	switch e {
	case KIOUSBSuperSpeedPlusDeviceCapabilitySublinkDirectionPhase:
		return "KIOUSBSuperSpeedPlusDeviceCapabilitySublinkDirectionPhase"
	default:
		return fmt.Sprintf("KIOUSBSuperSpeedPlusDeviceCapabilitySublink(%d)", e)
	}
}

type KIOUSBUSB20ExtensionCapabilityBES uint

const (
	KIOUSBUSB20ExtensionCapabilityBESL KIOUSBUSB20ExtensionCapabilityBES = 0
	KIOUSBUSB20ExtensionCapabilityBESLValid KIOUSBUSB20ExtensionCapabilityBES = 0
)

func (e KIOUSBUSB20ExtensionCapabilityBES) String() string {
	switch e {
	case KIOUSBUSB20ExtensionCapabilityBESL:
		return "KIOUSBUSB20ExtensionCapabilityBESL"
	default:
		return fmt.Sprintf("KIOUSBUSB20ExtensionCapabilityBES(%d)", e)
	}
}

type KIOUSBVendorIDAppleComputer uint

const (
	KIOUSBVendorIDApple KIOUSBVendorIDAppleComputer = 0
)

func (e KIOUSBVendorIDAppleComputer) String() string {
	switch e {
	case KIOUSBVendorIDApple:
		return "KIOUSBVendorIDApple"
	default:
		return fmt.Sprintf("KIOUSBVendorIDAppleComputer(%d)", e)
	}
}

const KIOUserClientAsyncArgumentsCountMax uint = 0

const KIOUserClientAsyncReferenceCountMax uint = 0

const KIOUserClientMemoryReadOnly uint = 0

const KIOUserClientMethodArgumentsCurrentVersion uint = 0

const KIOUserClientScalarArrayCountMax uint = 0

const KIOUserClientVariableStructureSize uint = 0

const KIOUserNotifyMaxMessageSize uint = 0

const KIOUserNotifyOptionCanDrop uint = 0

type KIOVideoBooleanControlClassID uint

const (
	// KIOVideoBooleanControlClassIDDirection: # Discussion
	KIOVideoBooleanControlClassIDDirection KIOVideoBooleanControlClassID = 0
	// KIOVideoBooleanControlClassIDJack: # Discussion
	KIOVideoBooleanControlClassIDJack KIOVideoBooleanControlClassID = 0
)

func (e KIOVideoBooleanControlClassID) String() string {
	switch e {
	case KIOVideoBooleanControlClassIDDirection:
		return "KIOVideoBooleanControlClassIDDirection"
	default:
		return fmt.Sprintf("KIOVideoBooleanControlClassID(%d)", e)
	}
}

type KIOVideoControl uint

const (
	// KIOVideoControlElementMaster: # Discussion
	KIOVideoControlElementMaster KIOVideoControl = 0
	// KIOVideoControlScopeInput: # Discussion
	KIOVideoControlScopeInput KIOVideoControl = 0
)

func (e KIOVideoControl) String() string {
	switch e {
	case KIOVideoControlElementMaster:
		return "KIOVideoControlElementMaster"
	default:
		return fmt.Sprintf("KIOVideoControl(%d)", e)
	}
}

type KIOVideoControlBaseClassID uint

const (
	// KIOVideoControlBaseClassIDBoolean: # Discussion
	KIOVideoControlBaseClassIDBoolean KIOVideoControlBaseClassID = 0
	// KIOVideoControlBaseClassIDFeature: # Discussion
	KIOVideoControlBaseClassIDFeature KIOVideoControlBaseClassID = 0
)

func (e KIOVideoControlBaseClassID) String() string {
	switch e {
	case KIOVideoControlBaseClassIDBoolean:
		return "KIOVideoControlBaseClassIDBoolean"
	default:
		return fmt.Sprintf("KIOVideoControlBaseClassID(%d)", e)
	}
}

type KIOVideoDeviceMethod uint

const (
	KIOVideoDeviceMethodClose KIOVideoDeviceMethod = 0
	KIOVideoDeviceMethodOpen KIOVideoDeviceMethod = 0
)

func (e KIOVideoDeviceMethod) String() string {
	switch e {
	case KIOVideoDeviceMethodClose:
		return "KIOVideoDeviceMethodClose"
	default:
		return fmt.Sprintf("KIOVideoDeviceMethod(%d)", e)
	}
}

type KIOVideoDevicePortTypeNotification uint

const (
	KIOVideoDevicePortTypeInput KIOVideoDevicePortTypeNotification = 0
)

func (e KIOVideoDevicePortTypeNotification) String() string {
	switch e {
	case KIOVideoDevicePortTypeInput:
		return "KIOVideoDevicePortTypeInput"
	default:
		return fmt.Sprintf("KIOVideoDevicePortTypeNotification(%d)", e)
	}
}

type KIOVideoFeatureControlClassID uint

const (
	// KIOVideoFeatureControlClassIDBacklightCompensation: # Discussion
	KIOVideoFeatureControlClassIDBacklightCompensation KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDZoom: # Discussion
	KIOVideoFeatureControlClassIDZoom KIOVideoFeatureControlClassID = 0
)

func (e KIOVideoFeatureControlClassID) String() string {
	switch e {
	case KIOVideoFeatureControlClassIDBacklightCompensation:
		return "KIOVideoFeatureControlClassIDBacklightCompensation"
	default:
		return fmt.Sprintf("KIOVideoFeatureControlClassID(%d)", e)
	}
}

type KIOVideoSelectorControlClassIDDataSource uint

const (
	// KIOVideoSelectorControlClassIDDataDestination: # Discussion
	KIOVideoSelectorControlClassIDDataDestination KIOVideoSelectorControlClassIDDataSource = 0
)

func (e KIOVideoSelectorControlClassIDDataSource) String() string {
	switch e {
	case KIOVideoSelectorControlClassIDDataDestination:
		return "KIOVideoSelectorControlClassIDDataDestination"
	default:
		return fmt.Sprintf("KIOVideoSelectorControlClassIDDataSource(%d)", e)
	}
}

type KIOWSAA uint

const (
	KIOWSAA_Accelerated KIOWSAA = 0
	KIOWSAA_Hibernate KIOWSAA = 0
)

func (e KIOWSAA) String() string {
	switch e {
	case KIOWSAA_Accelerated:
		return "KIOWSAA_Accelerated"
	default:
		return fmt.Sprintf("KIOWSAA(%d)", e)
	}
}

const KIOWorkGroupMaxNameLength uint = 0

const KInflowForciblyEnabledBit uint = 0

const KInitialDriverDescriptor uint = 0

type KIsochronousTransactionOptions uint

const (
	KIsochronousTransactionOptionsNone KIsochronousTransactionOptions = 0
	KIsochronousTransactionOptionsWrap KIsochronousTransactionOptions = 0
)

func (e KIsochronousTransactionOptions) String() string {
	switch e {
	case KIsochronousTransactionOptionsNone:
		return "KIsochronousTransactionOptionsNone"
	default:
		return fmt.Sprintf("KIsochronousTransactionOptions(%d)", e)
	}
}

const KJIJournalInFSMask uint = 0

type KKUNCDefaultResponse uint

const (
	KKUNCAlternateResponse KKUNCDefaultResponse = 0
)

func (e KKUNCDefaultResponse) String() string {
	switch e {
	case KKUNCAlternateResponse:
		return "KKUNCAlternateResponse"
	default:
		return fmt.Sprintf("KKUNCDefaultResponse(%d)", e)
	}
}

type KKeyMaskCtrl uint

const (
	KKeyMaskAlt KKeyMaskCtrl = 0
)

func (e KKeyMaskCtrl) String() string {
	switch e {
	case KKeyMaskAlt:
		return "KKeyMaskAlt"
	default:
		return fmt.Sprintf("KKeyMaskCtrl(%d)", e)
	}
}

type KLinkState uint

const (
	KLinkStateL2 KLinkState = 0
	KLinkStateU0 KLinkState = 0
)

func (e KLinkState) String() string {
	switch e {
	case KLinkStateL2:
		return "KLinkStateL2"
	default:
		return fmt.Sprintf("KLinkState(%d)", e)
	}
}

type KMMC uint

const (
	KMMCNumPowerStates KMMC = 0
	KMMCPowerStateActive KMMC = 0
	KMMCPowerStateIdle KMMC = 0
	KMMCPowerStateSleep KMMC = 0
	KMMCPowerStateStandby KMMC = 0
	KMMCPowerStateSystemSleep KMMC = 0
)

func (e KMMC) String() string {
	switch e {
	case KMMCNumPowerStates:
		return "KMMCNumPowerStates"
	default:
		return fmt.Sprintf("KMMC(%d)", e)
	}
}

const KMachineTypeUnknown uint = 0

const KMaximumNumberOfInquiryAccessCodes uint = 0

type KMediaStateUnlocked uint

const (
	KMediaStateLocked KMediaStateUnlocked = 0
)

func (e KMediaStateUnlocked) String() string {
	switch e {
	case KMediaStateLocked:
		return "KMediaStateLocked"
	default:
		return fmt.Sprintf("KMediaStateUnlocked(%d)", e)
	}
}

const KMessageDeterminingMediaPresence uint = 0

const KMessageTrayStateChangeRequestAccepted uint = 0

type KMirror uint

const (
	KMirrorAreMirroredMask KMirror = 0
	KMirrorVAlignCenterMirrorMask KMirror = 0
)

func (e KMirror) String() string {
	switch e {
	case KMirrorAreMirroredMask:
		return "KMirrorAreMirroredMask"
	default:
		return fmt.Sprintf("KMirror(%d)", e)
	}
}

type KMirrorSameDepthOnlyMirrorMask uint

const (
	KMirrorCommonGammaMask KMirrorSameDepthOnlyMirrorMask = 0
)

func (e KMirrorSameDepthOnlyMirrorMask) String() string {
	switch e {
	case KMirrorCommonGammaMask:
		return "KMirrorCommonGammaMask"
	default:
		return fmt.Sprintf("KMirrorSameDepthOnlyMirrorMask(%d)", e)
	}
}

type KMode uint

const (
	KModeInterlaced KMode = 0
	KModeValidateAgainstDisplay KMode = 0
)

func (e KMode) String() string {
	switch e {
	case KModeInterlaced:
		return "KModeInterlaced"
	default:
		return fmt.Sprintf("KMode(%d)", e)
	}
}

type KModePageControl uint

const (
	KModePageControlChangeableValues KModePageControl = 0
	KModePageControlSavedValues KModePageControl = 0
)

func (e KModePageControl) String() string {
	switch e {
	case KModePageControlChangeableValues:
		return "KModePageControlChangeableValues"
	default:
		return fmt.Sprintf("KModePageControl(%d)", e)
	}
}

type KModeSenseSBCDeviceSpecific uint

const (
	// KModeSenseSBCDeviceSpecific_DPOFUABit: # Discussion
	KModeSenseSBCDeviceSpecific_DPOFUABit KModeSenseSBCDeviceSpecific = 0
	// KModeSenseSBCDeviceSpecific_WriteProtectBit: # Discussion
	KModeSenseSBCDeviceSpecific_WriteProtectBit KModeSenseSBCDeviceSpecific = 0
)

func (e KModeSenseSBCDeviceSpecific) String() string {
	switch e {
	case KModeSenseSBCDeviceSpecific_DPOFUABit:
		return "KModeSenseSBCDeviceSpecific_DPOFUABit"
	default:
		return fmt.Sprintf("KModeSenseSBCDeviceSpecific(%d)", e)
	}
}

const KNVRAMProperty uint = 0

type KNanosecondScale uint

const (
	// KMicrosecondScale: # Discussion
	KMicrosecondScale KNanosecondScale = 0
	// KTickScale: # Discussion
	KTickScale KNanosecondScale = 0
)

func (e KNanosecondScale) String() string {
	switch e {
	case KMicrosecondScale:
		return "KMicrosecondScale"
	default:
		return fmt.Sprintf("KNanosecondScale(%d)", e)
	}
}

type KNdrvTypeIs uint

const (
	KNdrvTypeIsBlockStorage KNdrvTypeIs = 0
	KNdrvTypeIsBusBridge KNdrvTypeIs = 0
)

func (e KNdrvTypeIs) String() string {
	switch e {
	case KNdrvTypeIsBlockStorage:
		return "KNdrvTypeIsBlockStorage"
	default:
		return fmt.Sprintf("KNdrvTypeIs(%d)", e)
	}
}

const KNilOptions uint = 0

type KNuDCL uint

const (
	KNuDCLDynamic KNuDCL = 0
	KNuDCLUpdateBeforeCallback KNuDCL = 0
)

func (e KNuDCL) String() string {
	switch e {
	case KNuDCLDynamic:
		return "KNuDCLDynamic"
	default:
		return fmt.Sprintf("KNuDCL(%d)", e)
	}
}

const KNumPortBytes uint = 0

type KOFVariablePerm uint

const (
	KOFVariablePermKernelOnly KOFVariablePerm = 0
	KOFVariablePermUserRead KOFVariablePerm = 0
)

func (e KOFVariablePerm) String() string {
	switch e {
	case KOFVariablePermKernelOnly:
		return "KOFVariablePermKernelOnly"
	default:
		return fmt.Sprintf("KOFVariablePerm(%d)", e)
	}
}

type KOFVariableType uint

const (
	KOFVariableTypeBoolean KOFVariableType = 0
	KOFVariableTypeData KOFVariableType = 0
	KOFVariableTypeNumber KOFVariableType = 0
	KOFVariableTypeString KOFVariableType = 0
)

func (e KOFVariableType) String() string {
	switch e {
	case KOFVariableTypeBoolean:
		return "KOFVariableTypeBoolean"
	default:
		return fmt.Sprintf("KOFVariableType(%d)", e)
	}
}

const KOSAsyncRef64Count uint = 0

const KOSAsyncRefCount uint = 0

const KOSClassCanRemote uint = 0

type KOSNotificationMessageID uint

const (
	KMaxAsyncArgs KOSNotificationMessageID = 0
	KOSNotificationMessageIDValue KOSNotificationMessageID = 0
)

func (e KOSNotificationMessageID) String() string {
	switch e {
	case KMaxAsyncArgs:
		return "KMaxAsyncArgs"
	default:
		return fmt.Sprintf("KOSNotificationMessageID(%d)", e)
	}
}

const KOSStringNoCopy uint = 0

type KOneSecondTimeoutInMS uint

const (
	KFortyFiveSecondTimeoutInMS KOneSecondTimeoutInMS = 0
)

func (e KOneSecondTimeoutInMS) String() string {
	switch e {
	case KFortyFiveSecondTimeoutInMS:
		return "KFortyFiveSecondTimeoutInMS"
	default:
		return fmt.Sprintf("KOneSecondTimeoutInMS(%d)", e)
	}
}

type KOrConnections uint

const (
	KAndConnections KOrConnections = 0
)

func (e KOrConnections) String() string {
	switch e {
	case KAndConnections:
		return "KAndConnections"
	default:
		return fmt.Sprintf("KOrConnections(%d)", e)
	}
}

type KPCI2PCIPrimaryBus uint

const (
	KPCI2PCIBridgeControl KPCI2PCIPrimaryBus = 0
)

func (e KPCI2PCIPrimaryBus) String() string {
	switch e {
	case KPCI2PCIBridgeControl:
		return "KPCI2PCIBridgeControl"
	default:
		return fmt.Sprintf("KPCI2PCIPrimaryBus(%d)", e)
	}
}

type KPCIPMC uint

const (
	KPCIPMCD2Support KPCIPMC = 0
	KPCIPMCPMESupportFromD1 KPCIPMC = 0
)

func (e KPCIPMC) String() string {
	switch e {
	case KPCIPMCD2Support:
		return "KPCIPMCD2Support"
	default:
		return fmt.Sprintf("KPCIPMC(%d)", e)
	}
}

type KPCIPMCS uint

const (
	KPCIPMCSDefaultEnableBits KPCIPMCS = 0
	KPCIPMCSPowerStateMask KPCIPMCS = 0
)

func (e KPCIPMCS) String() string {
	switch e {
	case KPCIPMCSDefaultEnableBits:
		return "KPCIPMCSDefaultEnableBits"
	default:
		return fmt.Sprintf("KPCIPMCS(%d)", e)
	}
}

const KPEHaltCPU uint = 0

type KPEOptionKey uint

const (
	KPECommandKey KPEOptionKey = 0
)

func (e KPEOptionKey) String() string {
	switch e {
	case KPECommandKey:
		return "KPECommandKey"
	default:
		return fmt.Sprintf("KPEOptionKey(%d)", e)
	}
}

const KPEReadTOD uint = 0

type KPEScaleFactorUnknown uint

const (
	KPEScaleFactor1x KPEScaleFactorUnknown = 0
)

func (e KPEScaleFactorUnknown) String() string {
	switch e {
	case KPEScaleFactor1x:
		return "KPEScaleFactor1x"
	default:
		return fmt.Sprintf("KPEScaleFactorUnknown(%d)", e)
	}
}

type KPEWaitForInput uint

const (
	KPERawInput KPEWaitForInput = 0
)

func (e KPEWaitForInput) String() string {
	switch e {
	case KPERawInput:
		return "KPERawInput"
	default:
		return fmt.Sprintf("KPEWaitForInput(%d)", e)
	}
}

type KPM uint

const (
	KPMGeneralAggressiveness KPM = 0
	KPMMinutesToSleep KPM = 0
)

func (e KPM) String() string {
	switch e {
	case KPMGeneralAggressiveness:
		return "KPMGeneralAggressiveness"
	default:
		return fmt.Sprintf("KPM(%d)", e)
	}
}

type KPerPortSwitchingBit uint

const (
	KCompoundDeviceBit KPerPortSwitchingBit = 0
	KHubPortIndicatorBit KPerPortSwitchingBit = 0
)

func (e KPerPortSwitchingBit) String() string {
	switch e {
	case KCompoundDeviceBit:
		return "KCompoundDeviceBit"
	default:
		return fmt.Sprintf("KPerPortSwitchingBit(%d)", e)
	}
}

type KPeripheralDeviceTypeNoMatch uint

const (
	KDefaultProbeRanking KPeripheralDeviceTypeNoMatch = 0
	KThirdOrderRanking KPeripheralDeviceTypeNoMatch = 0
)

func (e KPeripheralDeviceTypeNoMatch) String() string {
	switch e {
	case KDefaultProbeRanking:
		return "KDefaultProbeRanking"
	default:
		return fmt.Sprintf("KPeripheralDeviceTypeNoMatch(%d)", e)
	}
}

type KPowerState uint

const (
	KPowerStateNeedsRefresh KPowerState = 0
	KPowerStateSupportsReducedPower2BitMask KPowerState = 0
)

func (e KPowerState) String() string {
	switch e {
	case KPowerStateNeedsRefresh:
		return "KPowerStateNeedsRefresh"
	default:
		return fmt.Sprintf("KPowerState(%d)", e)
	}
}

type KPowerStateReduced uint

const (
	KPowerStateReducedPower1 KPowerStateReduced = 0
	KPowerStateReducedPower3 KPowerStateReduced = 0
)

func (e KPowerStateReduced) String() string {
	switch e {
	case KPowerStateReducedPower1:
		return "KPowerStateReducedPower1"
	default:
		return fmt.Sprintf("KPowerStateReduced(%d)", e)
	}
}

type KPrdRootHub uint

const (
	KPrdRootHubApple KPrdRootHub = 0
	KPrdRootHubAppleE KPrdRootHub = 0
)

func (e KPrdRootHub) String() string {
	switch e {
	case KPrdRootHubApple:
		return "KPrdRootHubApple"
	default:
		return fmt.Sprintf("KPrdRootHub(%d)", e)
	}
}

type KRBC uint

const (
	KRBCNumPowerStates KRBC = 0
	KRBCPowerStateSystemSleep KRBC = 0
)

func (e KRBC) String() string {
	switch e {
	case KRBCNumPowerStates:
		return "KRBCNumPowerStates"
	default:
		return fmt.Sprintf("KRBC(%d)", e)
	}
}

type KRBCCmd uint

const (
	KRBCCmd_FORMAT_UNIT KRBCCmd = 0
	KRBCCmd_READ_CAPACITY KRBCCmd = 0
	KRBCCmd_START_STOP_UNIT KRBCCmd = 0
)

func (e KRBCCmd) String() string {
	switch e {
	case KRBCCmd_FORMAT_UNIT:
		return "KRBCCmd_FORMAT_UNIT"
	default:
		return fmt.Sprintf("KRBCCmd(%d)", e)
	}
}

type KRSC uint

const (
	KRSCFour KRSC = 0
	KRSCZero KRSC = 0
)

func (e KRSC) String() string {
	switch e {
	case KRSCFour:
		return "KRSCFour"
	default:
		return fmt.Sprintf("KRSC(%d)", e)
	}
}

type KRangeSupportsCompositeSync uint

const (
	KRangeSupportsCompositeSyncBit KRangeSupportsCompositeSync = 0
	KRangeSupportsCompositeSyncMask KRangeSupportsCompositeSync = 0
)

func (e KRangeSupportsCompositeSync) String() string {
	switch e {
	case KRangeSupportsCompositeSyncBit:
		return "KRangeSupportsCompositeSyncBit"
	default:
		return fmt.Sprintf("KRangeSupportsCompositeSync(%d)", e)
	}
}

type KReg uint

const (
	KRegModifierMask KReg = 0
	KRegNoModifiers KReg = 0
)

func (e KReg) String() string {
	switch e {
	case KRegModifierMask:
		return "KRegModifierMask"
	default:
		return fmt.Sprintf("KReg(%d)", e)
	}
}

const KRegCStrMaxEntryNameLength uint = 0

type KRegIter uint

const (
	KRegIterChildren KRegIter = 0
	KRegIterRoot KRegIter = 0
)

func (e KRegIter) String() string {
	switch e {
	case KRegIterChildren:
		return "KRegIterChildren"
	default:
		return fmt.Sprintf("KRegIter(%d)", e)
	}
}

const KRegMaxPropertyNameLength uint = 0

const KRegMaximumPropertyNameLength uint = 0

type KRegPathNameSeparator uint

const (
	KRegEntryNameTerminator KRegPathNameSeparator = 0
)

func (e KRegPathNameSeparator) String() string {
	switch e {
	case KRegEntryNameTerminator:
		return "KRegEntryNameTerminator"
	default:
		return fmt.Sprintf("KRegPathNameSeparator(%d)", e)
	}
}

type KRegPropertyValueIsSavedToNVRAM uint

const (
	KRegPropertyValueIsSavedToDisk KRegPropertyValueIsSavedToNVRAM = 0
)

func (e KRegPropertyValueIsSavedToNVRAM) String() string {
	switch e {
	case KRegPropertyValueIsSavedToDisk:
		return "KRegPropertyValueIsSavedToDisk"
	default:
		return fmt.Sprintf("KRegPropertyValueIsSavedToNVRAM(%d)", e)
	}
}

type KRequestDirection uint

const (
	KRequestDirectionIn KRequestDirection = 0
	KRequestDirectionOut KRequestDirection = 0
)

func (e KRequestDirection) String() string {
	switch e {
	case KRequestDirectionIn:
		return "KRequestDirectionIn"
	default:
		return fmt.Sprintf("KRequestDirection(%d)", e)
	}
}

type KRequestRecipient uint

const (
	KRequestRecipientDevice KRequestRecipient = 0
	KRequestRecipientEndpoint KRequestRecipient = 0
)

func (e KRequestRecipient) String() string {
	switch e {
	case KRequestRecipientDevice:
		return "KRequestRecipientDevice"
	default:
		return fmt.Sprintf("KRequestRecipient(%d)", e)
	}
}

const KResolutionHasMultipleDepthSizes uint = 0

type KRootDomainSleepNotSupported uint

const (
	KFrameBufferDeepSleepSupported KRootDomainSleepNotSupported = 0
)

func (e KRootDomainSleepNotSupported) String() string {
	switch e {
	case KFrameBufferDeepSleepSupported:
		return "KFrameBufferDeepSleepSupported"
	default:
		return fmt.Sprintf("KRootDomainSleepNotSupported(%d)", e)
	}
}

type KSBC uint

const (
	KSBCNumPowerStates KSBC = 0
	KSBCPowerStateStandby KSBC = 0
)

func (e KSBC) String() string {
	switch e {
	case KSBCNumPowerStates:
		return "KSBCNumPowerStates"
	default:
		return fmt.Sprintf("KSBC(%d)", e)
	}
}

type KSBCCmd uint

const (
	KSBCCmd_CHANGE_DEFINITION KSBCCmd = 0
	KSBCCmd_SEARCH_DATA_EQUAL_10 KSBCCmd = 0
)

func (e KSBCCmd) String() string {
	switch e {
	case KSBCCmd_CHANGE_DEFINITION:
		return "KSBCCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KSBCCmd(%d)", e)
	}
}

type KSBCModePage uint

const (
	// KSBCModePageCachingCode: # Discussion
	KSBCModePageCachingCode KSBCModePage = 0
	// KSBCModePageFormatDeviceCode: # Discussion
	KSBCModePageFormatDeviceCode KSBCModePage = 0
)

func (e KSBCModePage) String() string {
	switch e {
	case KSBCModePageCachingCode:
		return "KSBCModePageCachingCode"
	default:
		return fmt.Sprintf("KSBCModePage(%d)", e)
	}
}

type KSBCModePageCaching uint

const (
	// KSBCModePageCaching_RCD_Mask: # Discussion
	KSBCModePageCaching_RCD_Mask KSBCModePageCaching = 0
	// KSBCModePageCaching_SIZE_Mask: # Discussion
	KSBCModePageCaching_SIZE_Mask KSBCModePageCaching = 0
)

func (e KSBCModePageCaching) String() string {
	switch e {
	case KSBCModePageCaching_RCD_Mask:
		return "KSBCModePageCaching_RCD_Mask"
	default:
		return fmt.Sprintf("KSBCModePageCaching(%d)", e)
	}
}

type KSBCModePageFlexibleDisk uint

const (
	// KSBCModePageFlexibleDisk_MO_Bit: # Discussion
	KSBCModePageFlexibleDisk_MO_Bit KSBCModePageFlexibleDisk = 0
	// KSBCModePageFlexibleDisk_TRDY_Bit: # Discussion
	KSBCModePageFlexibleDisk_TRDY_Bit KSBCModePageFlexibleDisk = 0
)

func (e KSBCModePageFlexibleDisk) String() string {
	switch e {
	case KSBCModePageFlexibleDisk_MO_Bit:
		return "KSBCModePageFlexibleDisk_MO_Bit"
	default:
		return fmt.Sprintf("KSBCModePageFlexibleDisk(%d)", e)
	}
}

type KSBCWOCmd uint

const (
	KSBCWOCmd_CHANGE_DEFINITION KSBCWOCmd = 0
	KSBCWOCmd_LOG_SELECT KSBCWOCmd = 0
)

func (e KSBCWOCmd) String() string {
	switch e {
	case KSBCWOCmd_CHANGE_DEFINITION:
		return "KSBCWOCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KSBCWOCmd(%d)", e)
	}
}

type KSCSICDBSize uint

const (
	// KSCSICDBSize_10Byte: # Discussion
	KSCSICDBSize_10Byte KSCSICDBSize = 0
	// KSCSICDBSize_Maximum: # Discussion
	KSCSICDBSize_Maximum KSCSICDBSize = 0
)

func (e KSCSICDBSize) String() string {
	switch e {
	case KSCSICDBSize_10Byte:
		return "KSCSICDBSize_10Byte"
	default:
		return fmt.Sprintf("KSCSICDBSize(%d)", e)
	}
}

type KSCSICmd uint

const (
	KSCSICmdVariableLengthCDB KSCSICmd = 0
	KSCSICmd_GET_EVENT_STATUS_NOTIFICATION KSCSICmd = 0
)

func (e KSCSICmd) String() string {
	switch e {
	case KSCSICmdVariableLengthCDB:
		return "KSCSICmdVariableLengthCDB"
	default:
		return fmt.Sprintf("KSCSICmd(%d)", e)
	}
}

const KSCSIControllerNotificationBusReset uint = 0

type KSCSIDataTransfer uint

const (
	// KSCSIDataTransfer_FromInitiatorToTarget: # Discussion
	KSCSIDataTransfer_FromInitiatorToTarget KSCSIDataTransfer = 0
	// KSCSIDataTransfer_NoDataTransfer: # Discussion
	KSCSIDataTransfer_NoDataTransfer KSCSIDataTransfer = 0
)

func (e KSCSIDataTransfer) String() string {
	switch e {
	case KSCSIDataTransfer_FromInitiatorToTarget:
		return "KSCSIDataTransfer_FromInitiatorToTarget"
	default:
		return fmt.Sprintf("KSCSIDataTransfer(%d)", e)
	}
}

type KSCSIParallelMessage uint

const (
	KSCSIParallelMessage_ABORT_TASK KSCSIParallelMessage = 0
	KSCSIParallelMessage_ABORT_TASK_SET KSCSIParallelMessage = 0
	KSCSIParallelMessage_ACA KSCSIParallelMessage = 0
	KSCSIParallelMessage_CLEAR_ACA KSCSIParallelMessage = 0
	KSCSIParallelMessage_CLEAR_TASK_SET KSCSIParallelMessage = 0
	KSCSIParallelMessage_DISCONNECT KSCSIParallelMessage = 0
	KSCSIParallelMessage_EXTENDED_MESSAGE KSCSIParallelMessage = 0
	KSCSIParallelMessage_HEAD_OF_QUEUE KSCSIParallelMessage = 0
	KSCSIParallelMessage_IDENTIFY KSCSIParallelMessage = 0
	KSCSIParallelMessage_IGNORE_WIDE_RESIDUE KSCSIParallelMessage = 0
	KSCSIParallelMessage_INITIATOR_DETECTED_ERROR KSCSIParallelMessage = 0
	KSCSIParallelMessage_LINKED_COMMAND_COMPLETE KSCSIParallelMessage = 0
	KSCSIParallelMessage_LOGICAL_UNIT_RESET KSCSIParallelMessage = 0
	KSCSIParallelMessage_MESSAGE_PARITY_ERROR KSCSIParallelMessage = 0
	KSCSIParallelMessage_MESSAGE_REJECT KSCSIParallelMessage = 0
	KSCSIParallelMessage_MODIFY_DATA_POINTER KSCSIParallelMessage = 0
	KSCSIParallelMessage_NO_OPERATION KSCSIParallelMessage = 0
	KSCSIParallelMessage_ORDERED KSCSIParallelMessage = 0
	KSCSIParallelMessage_PARALLEL_PROTOCOL_REQUEST KSCSIParallelMessage = 0
	KSCSIParallelMessage_QAS_REQUEST KSCSIParallelMessage = 0
	KSCSIParallelMessage_RESTORE_POINTERS KSCSIParallelMessage = 0
	KSCSIParallelMessage_SAVE_DATA_POINTER KSCSIParallelMessage = 0
	KSCSIParallelMessage_SIMPLE KSCSIParallelMessage = 0
	KSCSIParallelMessage_SYNCHONOUS_DATA_TRANSFER_REQUEST KSCSIParallelMessage = 0
	KSCSIParallelMessage_TARGET_RESET KSCSIParallelMessage = 0
	KSCSIParallelMessage_TASK_COMPLETE KSCSIParallelMessage = 0
	KSCSIParallelMessage_WIDE_DATA_TRANSFER_REQUEST KSCSIParallelMessage = 0
)

func (e KSCSIParallelMessage) String() string {
	switch e {
	case KSCSIParallelMessage_ABORT_TASK:
		return "KSCSIParallelMessage_ABORT_TASK"
	default:
		return fmt.Sprintf("KSCSIParallelMessage(%d)", e)
	}
}

const KSCSIParallelTaskControllerIDQueueHead uint = 0

type KSCSIProtocolFeature uint

const (
	// KSCSIProtocolFeature_CPUInDiskMode: # Discussion
	KSCSIProtocolFeature_CPUInDiskMode KSCSIProtocolFeature = 0
	// KSCSIProtocolFeature_MaximumReadBlockTransferCount: # Discussion
	KSCSIProtocolFeature_MaximumReadBlockTransferCount KSCSIProtocolFeature = 0
)

func (e KSCSIProtocolFeature) String() string {
	switch e {
	case KSCSIProtocolFeature_CPUInDiskMode:
		return "KSCSIProtocolFeature_CPUInDiskMode"
	default:
		return fmt.Sprintf("KSCSIProtocolFeature(%d)", e)
	}
}

type KSCSIProtocolIdentifier uint

const (
	// KSCSIProtocolIdentifier_ADT: # Discussion
	KSCSIProtocolIdentifier_ADT KSCSIProtocolIdentifier = 0
	// KSCSIProtocolIdentifier_ATAPI: # Discussion
	KSCSIProtocolIdentifier_ATAPI KSCSIProtocolIdentifier = 0
	// KSCSIProtocolIdentifier_FibreChannel: # Discussion
	KSCSIProtocolIdentifier_FibreChannel KSCSIProtocolIdentifier = 0
	// KSCSIProtocolIdentifier_FireWire: # Discussion
	KSCSIProtocolIdentifier_FireWire KSCSIProtocolIdentifier = 0
	// KSCSIProtocolIdentifier_None: # Discussion
	KSCSIProtocolIdentifier_None KSCSIProtocolIdentifier = 0
	// KSCSIProtocolIdentifier_ParallelSCSI: # Discussion
	KSCSIProtocolIdentifier_ParallelSCSI KSCSIProtocolIdentifier = 0
	// KSCSIProtocolIdentifier_RDMA: # Discussion
	KSCSIProtocolIdentifier_RDMA KSCSIProtocolIdentifier = 0
	// KSCSIProtocolIdentifier_SAS: # Discussion
	KSCSIProtocolIdentifier_SAS KSCSIProtocolIdentifier = 0
	// KSCSIProtocolIdentifier_SSA: # Discussion
	KSCSIProtocolIdentifier_SSA KSCSIProtocolIdentifier = 0
	// KSCSIProtocolIdentifier_iSCSI: # Discussion
	KSCSIProtocolIdentifier_iSCSI KSCSIProtocolIdentifier = 0
)

func (e KSCSIProtocolIdentifier) String() string {
	switch e {
	case KSCSIProtocolIdentifier_ADT:
		return "KSCSIProtocolIdentifier_ADT"
	default:
		return fmt.Sprintf("KSCSIProtocolIdentifier(%d)", e)
	}
}

type KSCSIProtocolLayerPowerStateOff uint

const (
	KSCSIProtocolLayerNumDefaultStates KSCSIProtocolLayerPowerStateOff = 0
)

func (e KSCSIProtocolLayerPowerStateOff) String() string {
	switch e {
	case KSCSIProtocolLayerNumDefaultStates:
		return "KSCSIProtocolLayerNumDefaultStates"
	default:
		return fmt.Sprintf("KSCSIProtocolLayerPowerStateOff(%d)", e)
	}
}

const KSCSIProtocolPowerStateOff uint = 0

type KSCSIServiceResponse uint

const (
	// KSCSIServiceResponse_FUNCTION_COMPLETE: # Discussion
	KSCSIServiceResponse_FUNCTION_COMPLETE KSCSIServiceResponse = 0
	// KSCSIServiceResponse_FUNCTION_REJECTED: # Discussion
	KSCSIServiceResponse_FUNCTION_REJECTED KSCSIServiceResponse = 0
	// KSCSIServiceResponse_LINK_COMMAND_COMPLETE: # Discussion
	KSCSIServiceResponse_LINK_COMMAND_COMPLETE KSCSIServiceResponse = 0
	// KSCSIServiceResponse_Request_In_Process: # Discussion
	KSCSIServiceResponse_Request_In_Process KSCSIServiceResponse = 0
	// KSCSIServiceResponse_SERVICE_DELIVERY_OR_TARGET_FAILURE: # Discussion
	KSCSIServiceResponse_SERVICE_DELIVERY_OR_TARGET_FAILURE KSCSIServiceResponse = 0
	// KSCSIServiceResponse_TASK_COMPLETE: # Discussion
	KSCSIServiceResponse_TASK_COMPLETE KSCSIServiceResponse = 0
)

func (e KSCSIServiceResponse) String() string {
	switch e {
	case KSCSIServiceResponse_FUNCTION_COMPLETE:
		return "KSCSIServiceResponse_FUNCTION_COMPLETE"
	default:
		return fmt.Sprintf("KSCSIServiceResponse(%d)", e)
	}
}

type KSCSITask uint

const (
	// KSCSITask_ACA: # Discussion
	KSCSITask_ACA KSCSITask = 0
	// KSCSITask_HEAD_OF_QUEUE: # Discussion
	KSCSITask_HEAD_OF_QUEUE KSCSITask = 0
	// KSCSITask_ORDERED: # Discussion
	KSCSITask_ORDERED KSCSITask = 0
	// KSCSITask_SIMPLE: # Discussion
	KSCSITask_SIMPLE KSCSITask = 0
)

func (e KSCSITask) String() string {
	switch e {
	case KSCSITask_ACA:
		return "KSCSITask_ACA"
	default:
		return fmt.Sprintf("KSCSITask(%d)", e)
	}
}

type KSCSITaskMode uint

const (
	KSCSITaskMode_Autosense KSCSITaskMode = 0
	KSCSITaskMode_CommandExecution KSCSITaskMode = 0
)

func (e KSCSITaskMode) String() string {
	switch e {
	case KSCSITaskMode_Autosense:
		return "KSCSITaskMode_Autosense"
	default:
		return fmt.Sprintf("KSCSITaskMode(%d)", e)
	}
}

type KSCSITaskState uint

const (
	// KSCSITaskState_BLOCKED: # Discussion
	KSCSITaskState_BLOCKED KSCSITaskState = 0
	// KSCSITaskState_DORMANT: # Discussion
	KSCSITaskState_DORMANT KSCSITaskState = 0
	// KSCSITaskState_ENABLED: # Discussion
	KSCSITaskState_ENABLED KSCSITaskState = 0
	// KSCSITaskState_ENDED: # Discussion
	KSCSITaskState_ENDED KSCSITaskState = 0
	// KSCSITaskState_NEW_TASK: # Discussion
	KSCSITaskState_NEW_TASK KSCSITaskState = 0
)

func (e KSCSITaskState) String() string {
	switch e {
	case KSCSITaskState_BLOCKED:
		return "KSCSITaskState_BLOCKED"
	default:
		return fmt.Sprintf("KSCSITaskState(%d)", e)
	}
}

type KSCSITaskStatus int

const (
	// KSCSITaskStatus_ACA_ACTIVE: # Discussion
	KSCSITaskStatus_ACA_ACTIVE KSCSITaskStatus = 0
	// KSCSITaskStatus_BUSY: # Discussion
	KSCSITaskStatus_BUSY KSCSITaskStatus = 0
	// KSCSITaskStatus_CHECK_CONDITION: # Discussion
	KSCSITaskStatus_CHECK_CONDITION KSCSITaskStatus = 0
	// KSCSITaskStatus_CONDITION_MET: # Discussion
	KSCSITaskStatus_CONDITION_MET KSCSITaskStatus = 0
	// KSCSITaskStatus_DeliveryFailure: # Discussion
	KSCSITaskStatus_DeliveryFailure KSCSITaskStatus = 0
	// KSCSITaskStatus_DeviceNotPresent: # Discussion
	KSCSITaskStatus_DeviceNotPresent KSCSITaskStatus = 0
	// KSCSITaskStatus_DeviceNotResponding: # Discussion
	KSCSITaskStatus_DeviceNotResponding KSCSITaskStatus = 0
	// KSCSITaskStatus_GOOD: # Discussion
	KSCSITaskStatus_GOOD KSCSITaskStatus = 0
	// KSCSITaskStatus_INTERMEDIATE: # Discussion
	KSCSITaskStatus_INTERMEDIATE KSCSITaskStatus = 0
	// KSCSITaskStatus_INTERMEDIATE_CONDITION_MET: # Discussion
	KSCSITaskStatus_INTERMEDIATE_CONDITION_MET KSCSITaskStatus = 0
	// KSCSITaskStatus_No_Status: # Discussion
	KSCSITaskStatus_No_Status KSCSITaskStatus = 0
	// KSCSITaskStatus_ProtocolTimeoutOccurred: # Discussion
	KSCSITaskStatus_ProtocolTimeoutOccurred KSCSITaskStatus = 0
	// KSCSITaskStatus_RESERVATION_CONFLICT: # Discussion
	KSCSITaskStatus_RESERVATION_CONFLICT KSCSITaskStatus = 0
	// KSCSITaskStatus_TASK_SET_FULL: # Discussion
	KSCSITaskStatus_TASK_SET_FULL KSCSITaskStatus = 0
	// KSCSITaskStatus_TaskTimeoutOccurred: # Discussion
	KSCSITaskStatus_TaskTimeoutOccurred KSCSITaskStatus = 0
)

func (e KSCSITaskStatus) String() string {
	switch e {
	case KSCSITaskStatus_ACA_ACTIVE:
		return "KSCSITaskStatus_ACA_ACTIVE"
	default:
		return fmt.Sprintf("KSCSITaskStatus(%d)", e)
	}
}

const KSCSIUntaggedTaskIdentifier uint = 0

type KSMCCmd uint

const (
	KSMCCmd_MODE_SELECT_6 KSMCCmd = 0
	KSMCCmd_RELEASE_ELEMENT_10 KSMCCmd = 0
)

func (e KSMCCmd) String() string {
	switch e {
	case KSMCCmd_MODE_SELECT_6:
		return "KSMCCmd_MODE_SELECT_6"
	default:
		return fmt.Sprintf("KSMCCmd(%d)", e)
	}
}

type KSOFTRESET uint

const (
	KID_DRIVE KSOFTRESET = 0
	KPACKET KSOFTRESET = 0
	KSOFTRESETValue KSOFTRESET = 0
)

func (e KSOFTRESET) String() string {
	switch e {
	case KID_DRIVE:
		return "KID_DRIVE"
	default:
		return fmt.Sprintf("KSOFTRESET(%d)", e)
	}
}

type KSPCCmd uint

const (
	KSPCCmd_REPORT_LUNS KSPCCmd = 0
	KSPCCmd_SEND_DIAGNOSTICS KSPCCmd = 0
)

func (e KSPCCmd) String() string {
	switch e {
	case KSPCCmd_REPORT_LUNS:
		return "KSPCCmd_REPORT_LUNS"
	default:
		return fmt.Sprintf("KSPCCmd(%d)", e)
	}
}

type KSPCModePagePowerConditionCode uint

const (
	// KSPCModePageAllPagesCode: # Discussion
	KSPCModePageAllPagesCode KSPCModePagePowerConditionCode = 0
)

func (e KSPCModePagePowerConditionCode) String() string {
	switch e {
	case KSPCModePageAllPagesCode:
		return "KSPCModePageAllPagesCode"
	default:
		return fmt.Sprintf("KSPCModePagePowerConditionCode(%d)", e)
	}
}

type KSSCSeqCmd uint

const (
	KSSCSeqCmd_CHANGE_DEFINITION KSSCSeqCmd = 0
	KSSCSeqCmd_RECOVER_BUFFERED_DATA KSSCSeqCmd = 0
)

func (e KSSCSeqCmd) String() string {
	switch e {
	case KSSCSeqCmd_CHANGE_DEFINITION:
		return "KSSCSeqCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KSSCSeqCmd(%d)", e)
	}
}

type KSSHubPortLinkStateU0 uint

const (
	KSSHubPortLinkStateU3 KSSHubPortLinkStateU0 = 0
)

func (e KSSHubPortLinkStateU0) String() string {
	switch e {
	case KSSHubPortLinkStateU3:
		return "KSSHubPortLinkStateU3"
	default:
		return fmt.Sprintf("KSSHubPortLinkStateU0(%d)", e)
	}
}

type KSSHubPortStatusConnectionBit int

const (
	KHubPortBeingReset KSSHubPortStatusConnectionBit = 0
	KSSHubPortStatusOverCurrentBit KSSHubPortStatusConnectionBit = 0
)

func (e KSSHubPortStatusConnectionBit) String() string {
	switch e {
	case KHubPortBeingReset:
		return "KHubPortBeingReset"
	default:
		return fmt.Sprintf("KSSHubPortStatusConnectionBit(%d)", e)
	}
}

type KScale uint

const (
	KScaleCanSupportInsetMask KScale = 0
	KScaleStretchOnlyMask KScale = 0
)

func (e KScale) String() string {
	switch e {
	case KScaleCanSupportInsetMask:
		return "KScaleCanSupportInsetMask"
	default:
		return fmt.Sprintf("KScale(%d)", e)
	}
}

type KScaleStretchToFitMask uint

const (
	KScaleInvertXMask KScaleStretchToFitMask = 0
)

func (e KScaleStretchToFitMask) String() string {
	switch e {
	case KScaleInvertXMask:
		return "KScaleInvertXMask"
	default:
		return fmt.Sprintf("KScaleStretchToFitMask(%d)", e)
	}
}

type KSecondsInAMinute uint

const (
	K5Minutes KSecondsInAMinute = 0
	KSecondsInAMinuteValue KSecondsInAMinute = 0
)

func (e KSecondsInAMinute) String() string {
	switch e {
	case K5Minutes:
		return "K5Minutes"
	default:
		return fmt.Sprintf("KSecondsInAMinute(%d)", e)
	}
}

const KSecondsPerHour uint = 0

type KSelfIDPacketSize uint

const (
	KMaxSelfIDs KSelfIDPacketSize = 0
)

func (e KSelfIDPacketSize) String() string {
	switch e {
	case KMaxSelfIDs:
		return "KMaxSelfIDs"
	default:
		return fmt.Sprintf("KSelfIDPacketSize(%d)", e)
	}
}

const KSenseDefaultSize uint = 0

type KServiceCategory uint

const (
	KServiceCategoryADB KServiceCategory = 0
	KServiceCategoryBlockStorage KServiceCategory = 0
	KServiceCategoryDFM KServiceCategory = 0
	KServiceCategoryDisplay KServiceCategory = 0
	KServiceCategoryFileManager KServiceCategory = 0
	KServiceCategoryGeneric KServiceCategory = 0
	KServiceCategoryIDE KServiceCategory = 0
	KServiceCategoryKeyboard KServiceCategory = 0
	KServiceCategoryMotherBoard KServiceCategory = 0
	KServiceCategoryNVRAM KServiceCategory = 0
	KServiceCategoryNdrvDriver KServiceCategory = 0
	KServiceCategoryOpenTransport KServiceCategory = 0
	KServiceCategoryPCI KServiceCategory = 0
	KServiceCategoryPointing KServiceCategory = 0
	KServiceCategoryPowerMgt KServiceCategory = 0
	KServiceCategoryRTC KServiceCategory = 0
	KServiceCategoryScsiSIM KServiceCategory = 0
	KServiceCategorySound KServiceCategory = 0
)

func (e KServiceCategory) String() string {
	switch e {
	case KServiceCategoryADB:
		return "KServiceCategoryADB"
	default:
		return fmt.Sprintf("KServiceCategory(%d)", e)
	}
}

const KSetCLUTByValue uint = 0

const KSetClutAtSetEntries uint = 0

type KSetHubDepth uint

const (
	KGetPortErrorCount KSetHubDepth = 0
	KSetHubDepthValue KSetHubDepth = 0
)

func (e KSetHubDepth) String() string {
	switch e {
	case KGetPortErrorCount:
		return "KGetPortErrorCount"
	default:
		return fmt.Sprintf("KSetHubDepth(%d)", e)
	}
}

type KSizeOfATA uint

const (
	KSizeOfATAModelString KSizeOfATA = 0
	KSizeOfATARevisionString KSizeOfATA = 0
)

func (e KSizeOfATA) String() string {
	switch e {
	case KSizeOfATAModelString:
		return "KSizeOfATAModelString"
	default:
		return fmt.Sprintf("KSizeOfATA(%d)", e)
	}
}

type KSymLinkFileType uint

const (
	KSymLinkCreator KSymLinkFileType = 0
)

func (e KSymLinkFileType) String() string {
	switch e {
	case KSymLinkCreator:
		return "KSymLinkCreator"
	default:
		return fmt.Sprintf("KSymLinkFileType(%d)", e)
	}
}

type KSyncAnalog uint

const (
	KSyncAnalogBipolarSRGBSyncMask KSyncAnalog = 0
	KSyncAnalogCompositeSerrateMask KSyncAnalog = 0
)

func (e KSyncAnalog) String() string {
	switch e {
	case KSyncAnalogBipolarSRGBSyncMask:
		return "KSyncAnalogBipolarSRGBSyncMask"
	default:
		return fmt.Sprintf("KSyncAnalog(%d)", e)
	}
}

const KSyncPositivePolarityBit uint = 0

type KTheDescriptionSignature uint

const (
	KDriverDescriptionSignature KTheDescriptionSignature = 0
)

func (e KTheDescriptionSignature) String() string {
	switch e {
	case KDriverDescriptionSignature:
		return "KDriverDescriptionSignature"
	default:
		return fmt.Sprintf("KTheDescriptionSignature(%d)", e)
	}
}

type KThreadGroup uint

const (
	KThreadGroupApplication KThreadGroup = 0
	KThreadGroupBestEffort KThreadGroup = 0
	KThreadGroupCritical KThreadGroup = 0
	KThreadGroupEfficient KThreadGroup = 0
	KThreadGroupManaged KThreadGroup = 0
	KThreadGroupStrictTimers KThreadGroup = 0
	KThreadGroupUIApp KThreadGroup = 0
	KThreadGroupUIApplication KThreadGroup = 0
)

func (e KThreadGroup) String() string {
	switch e {
	case KThreadGroupApplication:
		return "KThreadGroupApplication"
	default:
		return fmt.Sprintf("KThreadGroup(%d)", e)
	}
}

type KThreadWait uint

const (
	KThreadWaitKernelMutex KThreadWait = 0
	KThreadWaitNone KThreadWait = 0
)

func (e KThreadWait) String() string {
	switch e {
	case KThreadWaitKernelMutex:
		return "KThreadWaitKernelMutex"
	default:
		return fmt.Sprintf("KThreadWait(%d)", e)
	}
}

type KTimingChangeRestrictedErr int

const (
	KTimingChangeRestrictedErrValue KTimingChangeRestrictedErr = 0
	KVideoBufferSizeErr KTimingChangeRestrictedErr = 0
)

func (e KTimingChangeRestrictedErr) String() string {
	switch e {
	case KTimingChangeRestrictedErrValue:
		return "KTimingChangeRestrictedErrValue"
	default:
		return fmt.Sprintf("KTimingChangeRestrictedErr(%d)", e)
	}
}

type KTransparentEncoding uint

const (
	KInvertingEncoding KTransparentEncoding = 0
)

func (e KTransparentEncoding) String() string {
	switch e {
	case KInvertingEncoding:
		return "KInvertingEncoding"
	default:
		return fmt.Sprintf("KTransparentEncoding(%d)", e)
	}
}

type KTransparentEncodingShift uint

const (
	KInvertingEncodedPixel KTransparentEncodingShift = 0
	KTransparentEncodedPixel KTransparentEncodingShift = 0
)

func (e KTransparentEncodingShift) String() string {
	switch e {
	case KInvertingEncodedPixel:
		return "KInvertingEncodedPixel"
	default:
		return fmt.Sprintf("KTransparentEncodingShift(%d)", e)
	}
}

type KUS uint

const (
	KUSB20ExtensionLPMSupported KUS = 0
	KUSB3HUBDesc KUS = 0
	KUSB3HubDescriptorType KUS = 0
	KUSBDeviceCapabilityValue KUS = 0
	KUSBHubDescriptorType KUS = 0
	KUSBSuperSpeedSupportsFS KUS = 0
)

func (e KUS) String() string {
	switch e {
	case KUSB20ExtensionLPMSupported:
		return "KUSB20ExtensionLPMSupported"
	default:
		return fmt.Sprintf("KUS(%d)", e)
	}
}

type KUSB uint

const (
	KUSBAdaptiveIsocSyncType KUSB = 0
	// KUSBAddExtraResetTimeBit: Request extra time after reset.
	KUSBAddExtraResetTimeBit KUSB = 0
	// KUSBAddExtraResetTimeMask: The mask to request extra time after reset.
	KUSBAddExtraResetTimeMask KUSB = 0
	KUSBAllStreams KUSB = 0
	KUSBAnyType KUSB = 0
	KUSBBulk KUSB = 0
	KUSBCommEthernetNetworkingSubClass KUSB = 0
	KUSBCommunicationControlInterfaceClass KUSB = 0
	KUSBCompositeSubClass KUSB = 0
	KUSBDevice KUSB = 0
	KUSBDeviceCountExceededNotificationType KUSB = 0
	KUSBDeviceIDMask KUSB = 0
	KUSBDeviceIDShift KUSB = 0
	KUSBDeviceMask KUSB = 0
	KUSBEndPtShift KUSB = 0
	KUSBEndpointCountExceededNotificationType KUSB = 0
	KUSBEndpointTransferTypeUCMask KUSB = 0
	KUSBEndpointbmAttributesUsageTypeShift KUSB = 0
	// KUSBFullSpeedMicrosecondsInFrame: # Discussion
	KUSBFullSpeedMicrosecondsInFrame KUSB = 0
	// KUSBHighSpeedMicrosecondsInFrame: # Discussion
	KUSBHighSpeedMicrosecondsInFrame KUSB = 0
	KUSBInterfaceIDMask KUSB = 0
	KUSBInterfaceIDShift KUSB = 0
	KUSBMaxDevice KUSB = 0
	KUSBMaxDevices KUSB = 0
	KUSBMaxInterfaces KUSB = 0
	KUSBMaxPipes KUSB = 0
	KUSBMaxStream KUSB = 0
	KUSBNoPipeIdx KUSB = 0
	KUSBNoStream KUSB = 0
	KUSBOther KUSB = 0
	KUSBPRimeStream KUSB = 0
	KUSBPhysicalInterfaceClass KUSB = 0
	KUSBPipeIDMask KUSB = 0
	// KUSBReEnumerateCaptureDeviceBit: The bit to capture the device.
	KUSBReEnumerateCaptureDeviceBit KUSB = 0
	// KUSBReEnumerateCaptureDeviceMask: The mask to capture the device.
	KUSBReEnumerateCaptureDeviceMask KUSB = 0
	// KUSBReEnumerateReleaseDeviceBit: The bit to release the device.
	KUSBReEnumerateReleaseDeviceBit KUSB = 0
	// KUSBReEnumerateReleaseDeviceMask: The mask to release the device.
	KUSBReEnumerateReleaseDeviceMask KUSB = 0
	KUSBStream0 KUSB = 0
	KUSBStreamIDAllStreamsMask KUSB = 0
	KUSBStreamIDMask KUSB = 0
	KUSBTooManyDevicesAddress KUSB = 0
	KUSBUCRequestWithoutUSBNotificationMask KUSB = 0
	KUSB_EPDesc_bmAttributes_SyncType_Shift KUSB = 0
	KUSB_HSFSEPDesc_wMaxPacketSize_Mult_Shift KUSB = 0
)

func (e KUSB) String() string {
	switch e {
	case KUSBAdaptiveIsocSyncType:
		return "KUSBAdaptiveIsocSyncType"
	default:
		return fmt.Sprintf("KUSB(%d)", e)
	}
}

const KUSB100mAAvailable uint = 0

type KUSB3LPMMaxU1SEL uint

const (
	KUSB2LPMMaxL1Timeout KUSB3LPMMaxU1SEL = 0
)

func (e KUSB3LPMMaxU1SEL) String() string {
	switch e {
	case KUSB2LPMMaxL1Timeout:
		return "KUSB2LPMMaxL1Timeout"
	default:
		return fmt.Sprintf("KUSB3LPMMaxU1SEL(%d)", e)
	}
}

const KUSBBillboardAdditinalInfoNoPower uint = 0

type KUSBBillboardUnspecifiedError int

const (
	KUSBBillboardAltModeConfigSuccess KUSBBillboardUnspecifiedError = 0
)

func (e KUSBBillboardUnspecifiedError) String() string {
	switch e {
	case KUSBBillboardAltModeConfigSuccess:
		return "KUSBBillboardAltModeConfigSuccess"
	default:
		return fmt.Sprintf("KUSBBillboardUnspecifiedError(%d)", e)
	}
}

type KUSBBillboardV uint

const (
	KUSBBillboardVConn1P5Watt KUSBBillboardV = 0
	KUSBBillboardVConn6Watt KUSBBillboardV = 0
)

func (e KUSBBillboardV) String() string {
	switch e {
	case KUSBBillboardVConn1P5Watt:
		return "KUSBBillboardVConn1P5Watt"
	default:
		return fmt.Sprintf("KUSBBillboardV(%d)", e)
	}
}

const KUSBCapsLockKey uint = 0

type KUSBCompositeClass uint

const (
	KUSBApplicationSpecificClass KUSBCompositeClass = 0
)

func (e KUSBCompositeClass) String() string {
	switch e {
	case KUSBApplicationSpecificClass:
		return "KUSBApplicationSpecificClass"
	default:
		return fmt.Sprintf("KUSBCompositeClass(%d)", e)
	}
}

const KUSBDFUAttributesMask uint = 0

type KUSBDefaultControlNoDataTimeoutMS uint

const (
	KUSBDefaultControlCompletionTimeoutMS KUSBDefaultControlNoDataTimeoutMS = 0
)

func (e KUSBDefaultControlNoDataTimeoutMS) String() string {
	switch e {
	case KUSBDefaultControlCompletionTimeoutMS:
		return "KUSBDefaultControlCompletionTimeoutMS"
	default:
		return fmt.Sprintf("KUSBDefaultControlNoDataTimeoutMS(%d)", e)
	}
}

type KUSBDeviceCapability uint

const (
	KUSBDeviceCapabilityContainerID KUSBDeviceCapability = 0
	KUSBDeviceCapabilitySuperSpeedUSB KUSBDeviceCapability = 0
)

func (e KUSBDeviceCapability) String() string {
	switch e {
	case KUSBDeviceCapabilityContainerID:
		return "KUSBDeviceCapabilityContainerID"
	default:
		return fmt.Sprintf("KUSBDeviceCapability(%d)", e)
	}
}

type KUSBDeviceLPMStatus int

const (
	KUSBDeviceLPMStatusDefault KUSBDeviceLPMStatus = 0
	KUSBDeviceLPMStatusDisabled KUSBDeviceLPMStatus = 0
)

func (e KUSBDeviceLPMStatus) String() string {
	switch e {
	case KUSBDeviceLPMStatusDefault:
		return "KUSBDeviceLPMStatusDefault"
	default:
		return fmt.Sprintf("KUSBDeviceLPMStatus(%d)", e)
	}
}

type KUSBDeviceSpeed uint

const (
	// KUSBDeviceSpeedHigh: # Discussion
	KUSBDeviceSpeedHigh KUSBDeviceSpeed = 0
	KUSBDeviceSpeedSuperPlus KUSBDeviceSpeed = 0
)

func (e KUSBDeviceSpeed) String() string {
	switch e {
	case KUSBDeviceSpeedHigh:
		return "KUSBDeviceSpeedHigh"
	default:
		return fmt.Sprintf("KUSBDeviceSpeed(%d)", e)
	}
}

const KUSBDisplayClass uint = 0

const KUSBEndpointPropertiesVersion3 uint = 0

type KUSBFeature uint

const (
	KUSBFeatureEndpointStall KUSBFeature = 0
	KUSBFeatureTestMode KUSBFeature = 0
)

func (e KUSBFeature) String() string {
	switch e {
	case KUSBFeatureEndpointStall:
		return "KUSBFeatureEndpointStall"
	default:
		return fmt.Sprintf("KUSBFeature(%d)", e)
	}
}

const KUSBFunctionRemoteWakeCapableBit uint = 0

const KUSBHSHubCommandAddHub uint = 0

type KUSBHost uint

const (
	KUSBHostClassRequestCompletionTimeout KUSBHost = 0
	KUSBHostDefaultControlCompletionTimeoutMS KUSBHost = 0
	KUSBHostMaxCountFullSpeedIsochronous KUSBHost = 0
	KUSBHostVendorIDAppleComputer KUSBHost = 0
)

func (e KUSBHost) String() string {
	switch e {
	case KUSBHostClassRequestCompletionTimeout:
		return "KUSBHostClassRequestCompletionTimeout"
	default:
		return fmt.Sprintf("KUSBHost(%d)", e)
	}
}

type KUSBHostConnectionSpeed uint

const (
	// KUSBHostConnectionSpeedCount: The speed count.
	KUSBHostConnectionSpeedCount KUSBHostConnectionSpeed = 0
	// KUSBHostConnectionSpeedFull: The full speed connection.
	KUSBHostConnectionSpeedFull KUSBHostConnectionSpeed = 0
	// KUSBHostConnectionSpeedHigh: The high speed connection.
	KUSBHostConnectionSpeedHigh KUSBHostConnectionSpeed = 0
	// KUSBHostConnectionSpeedLow: The low speed connection.
	KUSBHostConnectionSpeedLow KUSBHostConnectionSpeed = 0
	// KUSBHostConnectionSpeedSuper: The SuperSpeed connection.
	KUSBHostConnectionSpeedSuper KUSBHostConnectionSpeed = 0
	// KUSBHostConnectionSpeedSuperPlus: The SuperSpeedPlus connection.
	KUSBHostConnectionSpeedSuperPlus KUSBHostConnectionSpeed = 0
	// KUSBHostConnectionSpeedSuperPlusBy2: The SuperSpeedPlus By 2 connection.
	KUSBHostConnectionSpeedSuperPlusBy2 KUSBHostConnectionSpeed = 0
)

func (e KUSBHostConnectionSpeed) String() string {
	switch e {
	case KUSBHostConnectionSpeedCount:
		return "KUSBHostConnectionSpeedCount"
	default:
		return fmt.Sprintf("KUSBHostConnectionSpeed(%d)", e)
	}
}

type KUSBHostConnectorType uint

const (
	KUSB3TypeMicroABConnector KUSBHostConnectorType = 0
	KUSBHostConnectorTypeMiniAB KUSBHostConnectorType = 0
	KUSBHostConnectorTypeUSBTypeC KUSBHostConnectorType = 0
)

func (e KUSBHostConnectorType) String() string {
	switch e {
	case KUSB3TypeMicroABConnector:
		return "KUSB3TypeMicroABConnector"
	default:
		return fmt.Sprintf("KUSBHostConnectorType(%d)", e)
	}
}

type KUSBHostOpenOption uint

const (
	KUSBHostOpenOptionSelectAlternateSetting KUSBHostOpenOption = 0
	KUSBHostOpenOptionUserClientSession KUSBHostOpenOption = 0
)

func (e KUSBHostOpenOption) String() string {
	switch e {
	case KUSBHostOpenOptionSelectAlternateSetting:
		return "KUSBHostOpenOptionSelectAlternateSetting"
	default:
		return fmt.Sprintf("KUSBHostOpenOption(%d)", e)
	}
}

type KUSBHostPortConnectionSpeed uint

const (
	// KUSBHostPortConnectionSpeedCount: The speed count.
	KUSBHostPortConnectionSpeedCount KUSBHostPortConnectionSpeed = 0
	// KUSBHostPortConnectionSpeedFull: The full speed connection.
	KUSBHostPortConnectionSpeedFull KUSBHostPortConnectionSpeed = 0
	// KUSBHostPortConnectionSpeedHigh: The high speed connection.
	KUSBHostPortConnectionSpeedHigh KUSBHostPortConnectionSpeed = 0
	// KUSBHostPortConnectionSpeedLow: The low speed connection.
	KUSBHostPortConnectionSpeedLow KUSBHostPortConnectionSpeed = 0
	// KUSBHostPortConnectionSpeedNone: No connection speed.
	KUSBHostPortConnectionSpeedNone KUSBHostPortConnectionSpeed = 0
	// KUSBHostPortConnectionSpeedSuper: The SuperSpeed connection.
	KUSBHostPortConnectionSpeedSuper KUSBHostPortConnectionSpeed = 0
	// KUSBHostPortConnectionSpeedSuperPlus: The SuperSpeedPlus connection.
	KUSBHostPortConnectionSpeedSuperPlus KUSBHostPortConnectionSpeed = 0
	// KUSBHostPortConnectionSpeedSuperPlusBy2: The SuperSpeedPlus By 2 connection.
	KUSBHostPortConnectionSpeedSuperPlusBy2 KUSBHostPortConnectionSpeed = 0
)

func (e KUSBHostPortConnectionSpeed) String() string {
	switch e {
	case KUSBHostPortConnectionSpeedCount:
		return "KUSBHostPortConnectionSpeedCount"
	default:
		return fmt.Sprintf("KUSBHostPortConnectionSpeed(%d)", e)
	}
}

type KUSBHostPortStatus int

const (
	// KUSBHostPortStatusConnectedSpeedFull: The full connected speed.
	KUSBHostPortStatusConnectedSpeedFull KUSBHostPortStatus = 0
	// KUSBHostPortStatusConnectedSpeedHigh: The high connected speed.
	KUSBHostPortStatusConnectedSpeedHigh KUSBHostPortStatus = 0
	// KUSBHostPortStatusConnectedSpeedLow: The low connected speed.
	KUSBHostPortStatusConnectedSpeedLow KUSBHostPortStatus = 0
	// KUSBHostPortStatusConnectedSpeedMask: The connected speed mask.
	KUSBHostPortStatusConnectedSpeedMask KUSBHostPortStatus = 0
	// KUSBHostPortStatusConnectedSpeedNone: No connected speed.
	KUSBHostPortStatusConnectedSpeedNone KUSBHostPortStatus = 0
	// KUSBHostPortStatusConnectedSpeedPhase: The connected speed phase.
	KUSBHostPortStatusConnectedSpeedPhase KUSBHostPortStatus = 0
	// KUSBHostPortStatusConnectedSpeedSuper: The SuperSpeed connected speed.
	KUSBHostPortStatusConnectedSpeedSuper KUSBHostPortStatus = 0
	// KUSBHostPortStatusConnectedSpeedSuperPlus: The SuperSpeedPlus connected speed.
	KUSBHostPortStatusConnectedSpeedSuperPlus KUSBHostPortStatus = 0
	// KUSBHostPortStatusConnectedSpeedSuperPlusBy2: The SuperSpeedPlus By 2 connected speed.
	KUSBHostPortStatusConnectedSpeedSuperPlusBy2 KUSBHostPortStatus = 0
	// KUSBHostPortStatusEnabled: The port status enabled value.
	KUSBHostPortStatusEnabled KUSBHostPortStatus = 0
	// KUSBHostPortStatusOvercurrent: The port status overcurrent.
	KUSBHostPortStatusOvercurrent KUSBHostPortStatus = 0
	// KUSBHostPortStatusPortTypeAccessory: The accessory port type.
	KUSBHostPortStatusPortTypeAccessory KUSBHostPortStatus = 0
	// KUSBHostPortStatusPortTypeCaptive: The captive port type.
	KUSBHostPortStatusPortTypeCaptive KUSBHostPortStatus = 0
	// KUSBHostPortStatusPortTypeInternal: The internal port type.
	KUSBHostPortStatusPortTypeInternal KUSBHostPortStatus = 0
	// KUSBHostPortStatusPortTypeMask: The port type mask.
	KUSBHostPortStatusPortTypeMask KUSBHostPortStatus = 0
	// KUSBHostPortStatusPortTypePhase: The port type phase.
	KUSBHostPortStatusPortTypePhase KUSBHostPortStatus = 0
	// KUSBHostPortStatusPortTypeReserved: The reserved port type.
	KUSBHostPortStatusPortTypeReserved KUSBHostPortStatus = 0
	// KUSBHostPortStatusPortTypeStandard: The standard port type.
	KUSBHostPortStatusPortTypeStandard KUSBHostPortStatus = 0
	// KUSBHostPortStatusResetting: The resetting port status.
	KUSBHostPortStatusResetting KUSBHostPortStatus = 0
	// KUSBHostPortStatusSuspended: The suspended port status.
	KUSBHostPortStatusSuspended KUSBHostPortStatus = 0
	// KUSBHostPortStatusTestMode: The test mode port status.
	KUSBHostPortStatusTestMode KUSBHostPortStatus = 0
)

func (e KUSBHostPortStatus) String() string {
	switch e {
	case KUSBHostPortStatusConnectedSpeedFull:
		return "KUSBHostPortStatusConnectedSpeedFull"
	default:
		return fmt.Sprintf("KUSBHostPortStatus(%d)", e)
	}
}

type KUSBHostPortType uint

const (
	// KUSBHostPortTypeAccessory: The accessory port type.
	KUSBHostPortTypeAccessory KUSBHostPortType = 0
	// KUSBHostPortTypeCaptive: The captive port type.
	KUSBHostPortTypeCaptive KUSBHostPortType = 0
	// KUSBHostPortTypeCount: The port type count.
	KUSBHostPortTypeCount KUSBHostPortType = 0
	// KUSBHostPortTypeExpressCard: The express card port type.
	KUSBHostPortTypeExpressCard KUSBHostPortType = 0
	// KUSBHostPortTypeInternal: The internal port type.
	KUSBHostPortTypeInternal KUSBHostPortType = 0
	// KUSBHostPortTypeStandard: The standard port type.
	KUSBHostPortTypeStandard KUSBHostPortType = 0
)

func (e KUSBHostPortType) String() string {
	switch e {
	case KUSBHostPortTypeAccessory:
		return "KUSBHostPortTypeAccessory"
	default:
		return fmt.Sprintf("KUSBHostPortType(%d)", e)
	}
}

type KUSBHub uint

const (
	KUSBHubLocalPowerChangeFeature KUSBHub = 0
	KUSBHubPortConnectionFeature KUSBHub = 0
)

func (e KUSBHub) String() string {
	switch e {
	case KUSBHubLocalPowerChangeFeature:
		return "KUSBHubLocalPowerChangeFeature"
	default:
		return fmt.Sprintf("KUSBHub(%d)", e)
	}
}

type KUSBHubRq uint

const (
	// KUSBHubRqClearFeature: A request to clear a feature.
	KUSBHubRqClearFeature KUSBHubRq = 0
	// KUSBHubRqGetConfig: A request to get a configuration.
	KUSBHubRqGetConfig KUSBHubRq = 0
	// KUSBHubRqGetDescriptor: A request to obtain a descriptor.
	KUSBHubRqGetDescriptor KUSBHubRq = 0
	// KUSBHubRqGetInterface: A request to get an interface.
	KUSBHubRqGetInterface KUSBHubRq = 0
	// KUSBHubRqGetState: A request to get state.
	KUSBHubRqGetState KUSBHubRq = 0
	// KUSBHubRqGetStatus: A request to get status.
	KUSBHubRqGetStatus KUSBHubRq = 0
	// KUSBHubRqReserved2: A reserved request.
	KUSBHubRqReserved2 KUSBHubRq = 0
	// KUSBHubRqSetAddress: A request to set an address.
	KUSBHubRqSetAddress KUSBHubRq = 0
	// KUSBHubRqSetConfig: A request to set a configuration.
	KUSBHubRqSetConfig KUSBHubRq = 0
	// KUSBHubRqSetDescriptor: A request to set a descriptor.
	KUSBHubRqSetDescriptor KUSBHubRq = 0
	// KUSBHubRqSetFeature: A request to set a feature.
	KUSBHubRqSetFeature KUSBHubRq = 0
	// KUSBHubRqSetInterface: A request to set an interface.
	KUSBHubRqSetInterface KUSBHubRq = 0
)

func (e KUSBHubRq) String() string {
	switch e {
	case KUSBHubRqClearFeature:
		return "KUSBHubRqClearFeature"
	default:
		return fmt.Sprintf("KUSBHubRq(%d)", e)
	}
}

type KUSBHubRqSetHubDepth uint

const (
	KUSBHubRqGetPortErrorCount KUSBHubRqSetHubDepth = 0
)

func (e KUSBHubRqSetHubDepth) String() string {
	switch e {
	case KUSBHubRqGetPortErrorCount:
		return "KUSBHubRqGetPortErrorCount"
	default:
		return fmt.Sprintf("KUSBHubRqSetHubDepth(%d)", e)
	}
}

type KUSBInformation uint

const (
	// KUSBInformationDeviceIsAttachedToEnclosure: The hub port that the USB device connects to has a USB connector on the enclosure.
	KUSBInformationDeviceIsAttachedToEnclosure KUSBInformation = 0
	// KUSBInformationDeviceIsAttachedToEnclosureMask: The mask indicating that the USB device has attached to an enclosure.
	KUSBInformationDeviceIsAttachedToEnclosureMask KUSBInformation = 0
	// KUSBInformationDeviceIsAttachedToRootHubBit: The USB device has directly attached to the root hub.
	KUSBInformationDeviceIsAttachedToRootHubBit KUSBInformation = 0
	// KUSBInformationDeviceIsAttachedToRootHubMask: The mask indicating that the USB device has attached to a root hub.
	KUSBInformationDeviceIsAttachedToRootHubMask KUSBInformation = 0
	// KUSBInformationDeviceIsCaptiveBit: The USB device has directly attached to its hub and isn’t removable.
	KUSBInformationDeviceIsCaptiveBit KUSBInformation = 0
	// KUSBInformationDeviceIsCaptiveMask: The mask indicating that the USB device is captive.
	KUSBInformationDeviceIsCaptiveMask KUSBInformation = 0
	// KUSBInformationDeviceIsConnectedBit: The system has connected the USB device to its hub.
	KUSBInformationDeviceIsConnectedBit KUSBInformation = 0
	// KUSBInformationDeviceIsConnectedMask: The mask indicating that the USB device has connected.
	KUSBInformationDeviceIsConnectedMask KUSBInformation = 0
	// KUSBInformationDeviceIsEnabledBit: The system has enabled the hub port that the USB device attaches to.
	KUSBInformationDeviceIsEnabledBit KUSBInformation = 0
	// KUSBInformationDeviceIsEnabledMask: The mask indicating that the USB device is in an enabled state.
	KUSBInformationDeviceIsEnabledMask KUSBInformation = 0
	// KUSBInformationDeviceIsInResetBit: The system is resetting the hub port that the USB attaches to.
	KUSBInformationDeviceIsInResetBit KUSBInformation = 0
	// KUSBInformationDeviceIsInResetMask: The mask indicating that the USB device is in a reset state.
	KUSBInformationDeviceIsInResetMask KUSBInformation = 0
	// KUSBInformationDeviceIsInternalBit: The USB device is internal to the enclosure.
	KUSBInformationDeviceIsInternalBit KUSBInformation = 0
	// KUSBInformationDeviceIsInternalMask: The mask indicating that the USB device is internal.
	KUSBInformationDeviceIsInternalMask KUSBInformation = 0
	// KUSBInformationDeviceIsOnThunderboltBit: The bit indicating that the USB device is on a Thunderbolt port.
	KUSBInformationDeviceIsOnThunderboltBit KUSBInformation = 0
	// KUSBInformationDeviceIsOnThunderboltMask: The mask indicating that the USB device is on a Thunderbolt port.
	KUSBInformationDeviceIsOnThunderboltMask KUSBInformation = 0
	// KUSBInformationDeviceIsRemote: This device attaches to the controller through a remote connection.
	KUSBInformationDeviceIsRemote KUSBInformation = 0
	// KUSBInformationDeviceIsRemoteMask: The mask indicating that the USB device is remote.
	KUSBInformationDeviceIsRemoteMask KUSBInformation = 0
	// KUSBInformationDeviceIsRootHub: The device is the root hub simulation.
	KUSBInformationDeviceIsRootHub KUSBInformation = 0
	// KUSBInformationDeviceIsRootHubMask: The mask indicating that the USB device is a root hub.
	KUSBInformationDeviceIsRootHubMask KUSBInformation = 0
	// KUSBInformationDeviceIsSuspendedBit: The system has suspended the hub port that the USB device attaches to.
	KUSBInformationDeviceIsSuspendedBit KUSBInformation = 0
	// KUSBInformationDeviceIsSuspendedMask: The mask indicating that the USB device is in a suspended state.
	KUSBInformationDeviceIsSuspendedMask KUSBInformation = 0
	// KUSBInformationDeviceOvercurrentBit: The USB device generated an overcurrent.
	KUSBInformationDeviceOvercurrentBit KUSBInformation = 0
	// KUSBInformationDeviceOvercurrentMask: The mask indicating that the USB device is in overcurrent.
	KUSBInformationDeviceOvercurrentMask KUSBInformation = 0
	// KUSBInformationDevicePortIsInTestModeBit: The hub port that the USB device attaches to is in test mode.
	KUSBInformationDevicePortIsInTestModeBit KUSBInformation = 0
	// KUSBInformationDevicePortIsInTestModeMask: The mask indicating that the USB device is in test mode.
	KUSBInformationDevicePortIsInTestModeMask KUSBInformation = 0
	// KUSBInformationRootHubIsBuiltInBit: The bit indicating that the root hub is built-in.
	KUSBInformationRootHubIsBuiltInBit KUSBInformation = 0
	// KUSBInformationRootHubIsBuiltInMask: The mask indicating that the root hub is built-in.
	KUSBInformationRootHubIsBuiltInMask KUSBInformation = 0
	// KUSBInformationRootHubisBuiltIn: The USB root hub is built-in.
	KUSBInformationRootHubisBuiltIn KUSBInformation = 0
	// KUSBInformationRootHubisBuiltInMask: The mask indicating that the root hub is built-in.
	KUSBInformationRootHubisBuiltInMask KUSBInformation = 0
)

func (e KUSBInformation) String() string {
	switch e {
	case KUSBInformationDeviceIsAttachedToEnclosure:
		return "KUSBInformationDeviceIsAttachedToEnclosure"
	default:
		return fmt.Sprintf("KUSBInformation(%d)", e)
	}
}

type KUSBLowLatency uint

const (
	// KUSBLowLatencyFrameListBuffer: A low-latency frame list buffer.
	KUSBLowLatencyFrameListBuffer KUSBLowLatency = 0
	// KUSBLowLatencyReadBuffer: A low-latency read buffer.
	KUSBLowLatencyReadBuffer KUSBLowLatency = 0
	// KUSBLowLatencyWriteBuffer: A low-latency write buffer.
	KUSBLowLatencyWriteBuffer KUSBLowLatency = 0
)

func (e KUSBLowLatency) String() string {
	switch e {
	case KUSBLowLatencyFrameListBuffer:
		return "KUSBLowLatencyFrameListBuffer"
	default:
		return fmt.Sprintf("KUSBLowLatency(%d)", e)
	}
}

const KUSBLowLatencyIsochTransferKey uint = 0

const KUSBMaxFSIsocEndpointReqCount uint = 0

type KUSBNotification uint

const (
	// KUSBNotificationPostForcedResume: The system sends a notification after a resume, which occurs after a forced suspend.
	KUSBNotificationPostForcedResume KUSBNotification = 0
	// KUSBNotificationPostForcedSuspend: The system sends a notification after a forced suspend completes.
	KUSBNotificationPostForcedSuspend KUSBNotification = 0
	// KUSBNotificationPreForcedResume: The system sends a notification before a resume, which occurs after a forced suspend.
	KUSBNotificationPreForcedResume KUSBNotification = 0
	// KUSBNotificationPreForcedSuspend: The system sends a notification prior to a forced suspend.
	KUSBNotificationPreForcedSuspend KUSBNotification = 0
)

func (e KUSBNotification) String() string {
	switch e {
	case KUSBNotificationPostForcedResume:
		return "KUSBNotificationPostForcedResume"
	default:
		return fmt.Sprintf("KUSBNotification(%d)", e)
	}
}

type KUSBNotificationPreForcedSuspendBit uint

const (
	KUSBNotificationPostForcedResumeBit KUSBNotificationPreForcedSuspendBit = 0
)

func (e KUSBNotificationPreForcedSuspendBit) String() string {
	switch e {
	case KUSBNotificationPostForcedResumeBit:
		return "KUSBNotificationPostForcedResumeBit"
	default:
		return fmt.Sprintf("KUSBNotificationPreForcedSuspendBit(%d)", e)
	}
}

type KUSBOut uint

const (
	KUSBAnyDirn KUSBOut = 0
)

func (e KUSBOut) String() string {
	switch e {
	case KUSBAnyDirn:
		return "KUSBAnyDirn"
	default:
		return fmt.Sprintf("KUSBOut(%d)", e)
	}
}

type KUSBPort uint

const (
	KUSBPortConnectable KUSBPort = 0
	KUSBPortNotConnectable KUSBPort = 0
)

func (e KUSBPort) String() string {
	switch e {
	case KUSBPortConnectable:
		return "KUSBPortConnectable"
	default:
		return fmt.Sprintf("KUSBPort(%d)", e)
	}
}

type KUSBPower uint

const (
	// KUSBPowerDuringSleep: The system uses the power during sleep.
	KUSBPowerDuringSleep KUSBPower = 0
	// KUSBPowerDuringWake: The system uses the power while it’s awake.
	KUSBPowerDuringWake KUSBPower = 0
	// KUSBPowerDuringWakeRevocable: The system requests power reallocation as revocable.
	KUSBPowerDuringWakeRevocable KUSBPower = 0
	// KUSBPowerDuringWakeUSB3: The system requests extra power allocation.
	KUSBPowerDuringWakeUSB3 KUSBPower = 0
	// KUSBPowerRequestSleepReallocate: The system requests power reallocation upon sleeping.
	KUSBPowerRequestSleepReallocate KUSBPower = 0
	// KUSBPowerRequestSleepRelease: The system requests a release of power upon sleeping.
	KUSBPowerRequestSleepRelease KUSBPower = 0
	// KUSBPowerRequestWakeReallocate: The system requests power reallocation upon waking.
	KUSBPowerRequestWakeReallocate KUSBPower = 0
	// KUSBPowerRequestWakeRelease: The system requests a release of power upon waking.
	KUSBPowerRequestWakeRelease KUSBPower = 0
)

func (e KUSBPower) String() string {
	switch e {
	case KUSBPowerDuringSleep:
		return "KUSBPowerDuringSleep"
	default:
		return fmt.Sprintf("KUSBPower(%d)", e)
	}
}

const KUSBRel10 uint = 0

type KUSBRqDirnShift uint

const (
	KUSBRqDirnMask KUSBRqDirnShift = 0
)

func (e KUSBRqDirnShift) String() string {
	switch e {
	case KUSBRqDirnMask:
		return "KUSBRqDirnMask"
	default:
		return fmt.Sprintf("KUSBRqDirnShift(%d)", e)
	}
}

type KUSBRqGetStatus int

const (
	KUSBRqClearFeature KUSBRqGetStatus = 0
)

func (e KUSBRqGetStatus) String() string {
	switch e {
	case KUSBRqClearFeature:
		return "KUSBRqClearFeature"
	default:
		return fmt.Sprintf("KUSBRqGetStatus(%d)", e)
	}
}

type KUSBStandard uint

const (
	KUSBClass KUSBStandard = 0
)

func (e KUSBStandard) String() string {
	switch e {
	case KUSBClass:
		return "KUSBClass"
	default:
		return fmt.Sprintf("KUSBStandard(%d)", e)
	}
}

type KUSPrintingClassGetDeviceID uint

const (
	KUSPrintingClassGePortStatus KUSPrintingClassGetDeviceID = 0
)

func (e KUSPrintingClassGetDeviceID) String() string {
	switch e {
	case KUSPrintingClassGePortStatus:
		return "KUSPrintingClassGePortStatus"
	default:
		return fmt.Sprintf("KUSPrintingClassGetDeviceID(%d)", e)
	}
}

type KUnknownConnect uint

const (
	KFixedModeCRTConnect KUnknownConnect = 0
	KPanelFSTNConnect KUnknownConnect = 0
)

func (e KUnknownConnect) String() string {
	switch e {
	case KFixedModeCRTConnect:
		return "KFixedModeCRTConnect"
	default:
		return fmt.Sprintf("KUnknownConnect(%d)", e)
	}
}

type KVBLInterruptServiceType uint

const (
	KFBOnlineInterruptServiceType KVBLInterruptServiceType = 0
	KFrameInterruptServiceType KVBLInterruptServiceType = 0
)

func (e KVBLInterruptServiceType) String() string {
	switch e {
	case KFBOnlineInterruptServiceType:
		return "KFBOnlineInterruptServiceType"
	default:
		return fmt.Sprintf("KVBLInterruptServiceType(%d)", e)
	}
}

type KVCDarkReboot uint

const (
	KVCAcquireImmediate KVCDarkReboot = 0
)

func (e KVCDarkReboot) String() string {
	switch e {
	case KVCAcquireImmediate:
		return "KVCAcquireImmediate"
	default:
		return fmt.Sprintf("KVCDarkReboot(%d)", e)
	}
}

const KVSLClamshellStateGestaltType uint = 0

const KVariableLengthArray uint = 0

type KVideo uint

const (
	KVideoCombinedI2CTypeMask KVideo = 0
	KVideoNoTransactionType KVideo = 0
)

func (e KVideo) String() string {
	switch e {
	case KVideoCombinedI2CTypeMask:
		return "KVideoCombinedI2CTypeMask"
	default:
		return fmt.Sprintf("KVideo(%d)", e)
	}
}

type KVideoBusType uint

const (
	KVideoBusTypeDisplayPort KVideoBusType = 0
	KVideoBusTypeInvalid KVideoBusType = 0
)

func (e KVideoBusType) String() string {
	switch e {
	case KVideoBusTypeDisplayPort:
		return "KVideoBusTypeDisplayPort"
	default:
		return fmt.Sprintf("KVideoBusType(%d)", e)
	}
}

const KVideoDefaultBus uint = 0

type KVideoUsageAddrSubAddr uint

const (
	KVideoUsageAddrSubAddrBit KVideoUsageAddrSubAddr = 0
	KVideoUsageAddrSubAddrMask KVideoUsageAddrSubAddr = 0
)

func (e KVideoUsageAddrSubAddr) String() string {
	switch e {
	case KVideoUsageAddrSubAddrBit:
		return "KVideoUsageAddrSubAddrBit"
	default:
		return fmt.Sprintf("KVideoUsageAddrSubAddr(%d)", e)
	}
}

type KXHCISSRootHubAddress uint

const (
	KSuperSpeedBusBitMask KXHCISSRootHubAddress = 0
	KXHCISSRootHubAddressValue KXHCISSRootHubAddress = 0
)

func (e KXHCISSRootHubAddress) String() string {
	switch e {
	case KSuperSpeedBusBitMask:
		return "KSuperSpeedBusBitMask"
	default:
		return fmt.Sprintf("KXHCISSRootHubAddress(%d)", e)
	}
}

type Kanalogsignallevel07000300 uint

const (
	KAnalogSignalLevel_0700_0000 Kanalogsignallevel07000300 = 0
)

func (e Kanalogsignallevel07000300) String() string {
	switch e {
	case KAnalogSignalLevel_0700_0000:
		return "KAnalogSignalLevel_0700_0000"
	default:
		return fmt.Sprintf("Kanalogsignallevel07000300(%d)", e)
	}
}

type KcdCompressionTypeT uint

const (
	KCDCT_NONE KcdCompressionTypeT = 0
)

func (e KcdCompressionTypeT) String() string {
	switch e {
	case KCDCT_NONE:
		return "KCDCT_NONE"
	default:
		return fmt.Sprintf("KcdCompressionTypeT(%d)", e)
	}
}

type KcdataSubtypeTypes uint

const (
	KC_ST_CHAR KcdataSubtypeTypes = 0
)

func (e KcdataSubtypeTypes) String() string {
	switch e {
	case KC_ST_CHAR:
		return "KC_ST_CHAR"
	default:
		return fmt.Sprintf("KcdataSubtypeTypes(%d)", e)
	}
}

type KdCallback uint

const (
	KD_CALLBACK_KDEBUG_DISABLED KdCallback = 0
	KD_CALLBACK_KDEBUG_ENABLED KdCallback = 0
	KD_CALLBACK_SYNC_FLUSH KdCallback = 0
)

func (e KdCallback) String() string {
	switch e {
	case KD_CALLBACK_KDEBUG_DISABLED:
		return "KD_CALLBACK_KDEBUG_DISABLED"
	default:
		return fmt.Sprintf("KdCallback(%d)", e)
	}
}

type Kdbg uint

const (
	KDBG_BUFINIT Kdbg = 0
	KDBG_MATCH_DISABLE Kdbg = 0
	KDBG_NOWRAP Kdbg = 0
	KDBG_WRAPPED Kdbg = 0
)

func (e Kdbg) String() string {
	switch e {
	case KDBG_BUFINIT:
		return "KDBG_BUFINIT"
	default:
		return fmt.Sprintf("Kdbg(%d)", e)
	}
}

type KdebugCoprocFlagsT uint

const (
	KDCP_CONTINUOUS_TIME KdebugCoprocFlagsT = 0
)

func (e KdebugCoprocFlagsT) String() string {
	switch e {
	case KDCP_CONTINUOUS_TIME:
		return "KDCP_CONTINUOUS_TIME"
	default:
		return fmt.Sprintf("KdebugCoprocFlagsT(%d)", e)
	}
}

type KdpEventE uint

const (
	KDP_EVENT_ENTER KdpEventE = 0
	KDP_EVENT_EXIT KdpEventE = 0
)

func (e KdpEventE) String() string {
	switch e {
	case KDP_EVENT_ENTER:
		return "KDP_EVENT_ENTER"
	default:
		return fmt.Sprintf("KdpEventE(%d)", e)
	}
}

type Kdtest uint

const (
	KDTEST_ABSOLUTE_TIMESTAMP Kdtest = 0
	KDTEST_CONTINUOUS_TIMESTAMP Kdtest = 0
)

func (e Kdtest) String() string {
	switch e {
	case KDTEST_ABSOLUTE_TIMESTAMP:
		return "KDTEST_ABSOLUTE_TIMESTAMP"
	default:
		return fmt.Sprintf("Kdtest(%d)", e)
	}
}

type Kextint9Nmiintsource uint

const (
	KExtInt9_NMIIntSource Kextint9Nmiintsource = 0
)

func (e Kextint9Nmiintsource) String() string {
	switch e {
	case KExtInt9_NMIIntSource:
		return "KExtInt9_NMIIntSource"
	default:
		return fmt.Sprintf("Kextint9Nmiintsource(%d)", e)
	}
}

type Kf uint

const (
	KF_COMPRSV_OVRD Kf = 0
	KF_DISABLE_FP_POPC_ON_PGFLT Kf = 0
	KF_DISABLE_PROCREF_TRACKING_OVRD Kf = 0
	KF_DISABLE_PROD_TRC_VALIDATION Kf = 0
	KF_INTERRUPT_MASKED_DEBUG_OVRD Kf = 0
	KF_INTERRUPT_MASKED_DEBUG_STACKSHOT_OVRD Kf = 0
	KF_IOTRACE_OVRD Kf = 0
	KF_IO_TIMEOUT_OVRD Kf = 0
	KF_MACH_ASSERT_OVRD Kf = 0
	KF_MADVISE_FREE_DEBUG_OVRD Kf = 0
	KF_MATV_OVRD Kf = 0
	KF_SERIAL_OVRD Kf = 0
)

func (e Kf) String() string {
	switch e {
	case KF_COMPRSV_OVRD:
		return "KF_COMPRSV_OVRD"
	default:
		return fmt.Sprintf("Kf(%d)", e)
	}
}

type KguardExc uint

const (
	KGUARD_EXC_DEALLOC_GAP KguardExc = 0
	KGUARD_EXC_DESTROY KguardExc = 0
	KGUARD_EXC_EXCEPTION_BEHAVIOR_ENFORCE KguardExc = 0
	KGUARD_EXC_IMMOVABLE KguardExc = 0
	KGUARD_EXC_IMMOVABLE_NON_FATAL KguardExc = 0
	KGUARD_EXC_INCORRECT_GUARD KguardExc = 0
	KGUARD_EXC_INVALID_ARGUMENT KguardExc = 0
	KGUARD_EXC_INVALID_NAME KguardExc = 0
	KGUARD_EXC_INVALID_OPTIONS KguardExc = 0
	KGUARD_EXC_INVALID_RIGHT KguardExc = 0
	KGUARD_EXC_INVALID_VALUE KguardExc = 0
	KGUARD_EXC_KERN_FAILURE KguardExc = 0
	KGUARD_EXC_KERN_NO_SPACE KguardExc = 0
	KGUARD_EXC_KERN_RESOURCE KguardExc = 0
	KGUARD_EXC_MOD_REFS KguardExc = 0
	KGUARD_EXC_MOD_REFS_NON_FATAL KguardExc = 0
	KGUARD_EXC_MSG_FILTERED KguardExc = 0
	KGUARD_EXC_PROVISIONAL_REPLY_PORT KguardExc = 0
	KGUARD_EXC_RCV_GUARDED_DESC KguardExc = 0
	KGUARD_EXC_RCV_INVALID_NAME KguardExc = 0
	KGUARD_EXC_RECLAIM_COPYIO_FAILURE KguardExc = 0
	KGUARD_EXC_REQUIRE_REPLY_PORT_SEMANTICS KguardExc = 0
	KGUARD_EXC_RIGHT_EXISTS KguardExc = 0
	KGUARD_EXC_SEND_INVALID_REPLY KguardExc = 0
	KGUARD_EXC_SEND_INVALID_RIGHT KguardExc = 0
	KGUARD_EXC_SEND_INVALID_VOUCHER KguardExc = 0
	KGUARD_EXC_SERVICE_PORT_VIOLATION_FATAL KguardExc = 0
	KGUARD_EXC_SERVICE_PORT_VIOLATION_NON_FATAL KguardExc = 0
	KGUARD_EXC_SET_CONTEXT KguardExc = 0
	KGUARD_EXC_STRICT_REPLY KguardExc = 0
	KGUARD_EXC_THREAD_SET_STATE KguardExc = 0
	KGUARD_EXC_UNGUARDED KguardExc = 0
)

func (e KguardExc) String() string {
	switch e {
	case KGUARD_EXC_DEALLOC_GAP:
		return "KGUARD_EXC_DEALLOC_GAP"
	default:
		return fmt.Sprintf("KguardExc(%d)", e)
	}
}

type KhidusageAdAlphanumericdisplay uint

const (
	KHIDUsage_AD_ASCIICharacterSet KhidusageAdAlphanumericdisplay = 0
)

func (e KhidusageAdAlphanumericdisplay) String() string {
	switch e {
	case KHIDUsage_AD_ASCIICharacterSet:
		return "KHIDUsage_AD_ASCIICharacterSet"
	default:
		return fmt.Sprintf("KhidusageAdAlphanumericdisplay(%d)", e)
	}
}

type KhidusageBcs uint

const (
	KHIDUsage_BCS_2DControlReport KhidusageBcs = 0
	KHIDUsage_BCS_CheckDisablePrice KhidusageBcs = 0
)

func (e KhidusageBcs) String() string {
	switch e {
	case KHIDUsage_BCS_2DControlReport:
		return "KHIDUsage_BCS_2DControlReport"
	default:
		return fmt.Sprintf("KhidusageBcs(%d)", e)
	}
}

type KhidusageBd uint

const (
	KHIDUsage_BD_6DotBrailleCell KhidusageBd = 0
	KHIDUsage_BD_BrailleRightControls KhidusageBd = 0
)

func (e KhidusageBd) String() string {
	switch e {
	case KHIDUsage_BD_6DotBrailleCell:
		return "KHIDUsage_BD_6DotBrailleCell"
	default:
		return fmt.Sprintf("KhidusageBd(%d)", e)
	}
}

type KhidusageBs uint

const (
	KHIDUsage_BS_AbsoluteStateOfCharge KhidusageBs = 0
	KHIDUsage_BS_ConnectionToSMBus KhidusageBs = 0
)

func (e KhidusageBs) String() string {
	switch e {
	case KHIDUsage_BS_AbsoluteStateOfCharge:
		return "KHIDUsage_BS_AbsoluteStateOfCharge"
	default:
		return fmt.Sprintf("KhidusageBs(%d)", e)
	}
}

type KhidusageButton uint

const (
	KHIDUsage_Button_109 KhidusageButton = 0
	KHIDUsage_Button_123 KhidusageButton = 0
)

func (e KhidusageButton) String() string {
	switch e {
	case KHIDUsage_Button_109:
		return "KHIDUsage_Button_109"
	default:
		return fmt.Sprintf("KhidusageButton(%d)", e)
	}
}

type KhidusageCc uint

const (
	KHIDUsage_CC_Autofocus KhidusageCc = 0
	KHIDUsage_CC_Shutter KhidusageCc = 0
	KHIDUsage_CC_Undefined KhidusageCc = 0
)

func (e KhidusageCc) String() string {
	switch e {
	case KHIDUsage_CC_Autofocus:
		return "KHIDUsage_CC_Autofocus"
	default:
		return fmt.Sprintf("KhidusageCc(%d)", e)
	}
}

type KhidusageCsmr uint

const (
	KHIDUsage_Csmr_ACSelectTable KhidusageCsmr = 0
	KHIDUsage_Csmr_Repeat KhidusageCsmr = 0
)

func (e KhidusageCsmr) String() string {
	switch e {
	case KHIDUsage_Csmr_ACSelectTable:
		return "KHIDUsage_Csmr_ACSelectTable"
	default:
		return fmt.Sprintf("KhidusageCsmr(%d)", e)
	}
}

type KhidusageDig uint

const (
	KHIDUsage_Dig_Altitude KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterEncodingUTF16BE KhidusageDig = 0
)

func (e KhidusageDig) String() string {
	switch e {
	case KHIDUsage_Dig_Altitude:
		return "KHIDUsage_Dig_Altitude"
	default:
		return fmt.Sprintf("KhidusageDig(%d)", e)
	}
}

type KhidusageGame3dgamecontroller uint

const (
	KHIDUsage_Game_3DGameController KhidusageGame3dgamecontroller = 0
)

func (e KhidusageGame3dgamecontroller) String() string {
	switch e {
	case KHIDUsage_Game_3DGameController:
		return "KHIDUsage_Game_3DGameController"
	default:
		return fmt.Sprintf("KhidusageGame3dgamecontroller(%d)", e)
	}
}

type KhidusageGd uint

const (
	KHIDUsage_GD_ByteCount KhidusageGd = 0
	KHIDUsage_GD_SystemMenu KhidusageGd = 0
)

func (e KhidusageGd) String() string {
	switch e {
	case KHIDUsage_GD_ByteCount:
		return "KHIDUsage_GD_ByteCount"
	default:
		return fmt.Sprintf("KhidusageGd(%d)", e)
	}
}

type KhidusageGendevcontrolsBackgroundcontrols uint

const (
	KHIDUsage_GenDevControls_BackgroundControls KhidusageGendevcontrolsBackgroundcontrols = 0
)

func (e KhidusageGendevcontrolsBackgroundcontrols) String() string {
	switch e {
	case KHIDUsage_GenDevControls_BackgroundControls:
		return "KHIDUsage_GenDevControls_BackgroundControls"
	default:
		return fmt.Sprintf("KhidusageGendevcontrolsBackgroundcontrols(%d)", e)
	}
}

type KhidusageLed uint

const (
	KHIDUsage_LED_Compose KhidusageLed = 0
	KHIDUsage_LED_Coverage KhidusageLed = 0
)

func (e KhidusageLed) String() string {
	switch e {
	case KHIDUsage_LED_Compose:
		return "KHIDUsage_LED_Compose"
	default:
		return fmt.Sprintf("KhidusageLed(%d)", e)
	}
}

type KhidusageMsrUndefined uint

const (
	KHIDUsage_MSR_DeviceReadOnly KhidusageMsrUndefined = 0
)

func (e KhidusageMsrUndefined) String() string {
	switch e {
	case KHIDUsage_MSR_DeviceReadOnly:
		return "KHIDUsage_MSR_DeviceReadOnly"
	default:
		return fmt.Sprintf("KhidusageMsrUndefined(%d)", e)
	}
}

type KhidusageOrdInstance1 uint

const (
	KHIDUsage_Ord_Instance1 KhidusageOrdInstance1 = 0
)

func (e KhidusageOrdInstance1) String() string {
	switch e {
	case KHIDUsage_Ord_Instance1:
		return "KHIDUsage_Ord_Instance1"
	default:
		return fmt.Sprintf("KhidusageOrdInstance1(%d)", e)
	}
}

type KhidusagePd uint

const (
	KHIDUsage_PD_PowerSummary KhidusagePd = 0
	KHIDUsage_PD_ShutdownRequested KhidusagePd = 0
)

func (e KhidusagePd) String() string {
	switch e {
	case KHIDUsage_PD_PowerSummary:
		return "KHIDUsage_PD_PowerSummary"
	default:
		return fmt.Sprintf("KhidusagePd(%d)", e)
	}
}

type KhidusagePid uint

const (
	KHIDUsage_PID_ActuatorsEnabled KhidusagePid = 0
	KHIDUsage_PID_PositiveSaturation KhidusagePid = 0
)

func (e KhidusagePid) String() string {
	switch e {
	case KHIDUsage_PID_ActuatorsEnabled:
		return "KHIDUsage_PID_ActuatorsEnabled"
	default:
		return fmt.Sprintf("KhidusagePid(%d)", e)
	}
}

type KhidusageSim uint

const (
	KHIDUsage_Sim_Accelerator KhidusageSim = 0
	KHIDUsage_Sim_CyclicControl KhidusageSim = 0
	KHIDUsage_Sim_HelicopterSimulationDevice KhidusageSim = 0
)

func (e KhidusageSim) String() string {
	switch e {
	case KHIDUsage_Sim_Accelerator:
		return "KHIDUsage_Sim_Accelerator"
	default:
		return fmt.Sprintf("KhidusageSim(%d)", e)
	}
}

type KhidusageSnsr uint

const (
	KHIDUsage_Snsr_Data_Motion_AngularVelocityXAxis KhidusageSnsr = 0
	KHIDUsage_Snsr_Event_SensorEvent_RangeMaxReached KhidusageSnsr = 0
)

func (e KhidusageSnsr) String() string {
	switch e {
	case KHIDUsage_Snsr_Data_Motion_AngularVelocityXAxis:
		return "KHIDUsage_Snsr_Data_Motion_AngularVelocityXAxis"
	default:
		return fmt.Sprintf("KhidusageSnsr(%d)", e)
	}
}

type KhidusageSprt uint

const (
	KHIDUsage_Sprt_11Iron KhidusageSprt = 0
	KHIDUsage_Sprt_PowerWedge KhidusageSprt = 0
)

func (e KhidusageSprt) String() string {
	switch e {
	case KHIDUsage_Sprt_11Iron:
		return "KHIDUsage_Sprt_11Iron"
	default:
		return fmt.Sprintf("KhidusageSprt(%d)", e)
	}
}

type KhidusageTfon uint

const (
	KHIDUsage_Tfon_LineBusyTone KhidusageTfon = 0
	KHIDUsage_Tfon_Redial KhidusageTfon = 0
)

func (e KhidusageTfon) String() string {
	switch e {
	case KHIDUsage_Tfon_LineBusyTone:
		return "KHIDUsage_Tfon_LineBusyTone"
	default:
		return fmt.Sprintf("KhidusageTfon(%d)", e)
	}
}

type KhidusageUndefined uint

const (
	KHIDUsage_Undefined KhidusageUndefined = 0
)

func (e KhidusageUndefined) String() string {
	switch e {
	case KHIDUsage_Undefined:
		return "KHIDUsage_Undefined"
	default:
		return fmt.Sprintf("KhidusageUndefined(%d)", e)
	}
}

type KhidusageVr uint

const (
	KHIDUsage_VR_AnimatronicDevice KhidusageVr = 0
	KHIDUsage_VR_HandTracker KhidusageVr = 0
)

func (e KhidusageVr) String() string {
	switch e {
	case KHIDUsage_VR_AnimatronicDevice:
		return "KHIDUsage_VR_AnimatronicDevice"
	default:
		return fmt.Sprintf("KhidusageVr(%d)", e)
	}
}

type KhidusageWd uint

const (
	KHIDUsage_WD_ScaleAtrributeReport KhidusageWd = 0
	KHIDUsage_WD_WeightUnitTaels KhidusageWd = 0
)

func (e KhidusageWd) String() string {
	switch e {
	case KHIDUsage_WD_ScaleAtrributeReport:
		return "KHIDUsage_WD_ScaleAtrributeReport"
	default:
		return fmt.Sprintf("KhidusageWd(%d)", e)
	}
}

type KhvIonNone uint

const (
	KHV_ION_ANY_SIZE KhvIonNone = 0
)

func (e KhvIonNone) String() string {
	switch e {
	case KHV_ION_ANY_SIZE:
		return "KHV_ION_ANY_SIZE"
	default:
		return fmt.Sprintf("KhvIonNone(%d)", e)
	}
}

type KinquiryByte3 uint

const (
	// KINQUIRY_Byte3_AERC_Bit: # Discussion
	KINQUIRY_Byte3_AERC_Bit KinquiryByte3 = 0
	// KINQUIRY_Byte3_NORMACA_Bit: # Discussion
	KINQUIRY_Byte3_NORMACA_Bit KinquiryByte3 = 0
)

func (e KinquiryByte3) String() string {
	switch e {
	case KINQUIRY_Byte3_AERC_Bit:
		return "KINQUIRY_Byte3_AERC_Bit"
	default:
		return fmt.Sprintf("KinquiryByte3(%d)", e)
	}
}

type KinquiryByte5 uint

const (
	// KINQUIRY_Byte5_3PC_Bit: # Discussion
	KINQUIRY_Byte5_3PC_Bit KinquiryByte5 = 0
	// KINQUIRY_Byte5_PROTECT_Mask: # Discussion
	KINQUIRY_Byte5_PROTECT_Mask KinquiryByte5 = 0
)

func (e KinquiryByte5) String() string {
	switch e {
	case KINQUIRY_Byte5_3PC_Bit:
		return "KINQUIRY_Byte5_3PC_Bit"
	default:
		return fmt.Sprintf("KinquiryByte5(%d)", e)
	}
}

type KinquiryByte56 uint

const (
	// KINQUIRY_Byte56_CLOCKING_ONLY_DT: # Discussion
	KINQUIRY_Byte56_CLOCKING_ONLY_DT KinquiryByte56 = 0
	KINQUIRY_Byte56_Offset KinquiryByte56 = 0
)

func (e KinquiryByte56) String() string {
	switch e {
	case KINQUIRY_Byte56_CLOCKING_ONLY_DT:
		return "KINQUIRY_Byte56_CLOCKING_ONLY_DT"
	default:
		return fmt.Sprintf("KinquiryByte56(%d)", e)
	}
}

type KinquiryByte6 uint

const (
	// KINQUIRY_Byte6_BQUE_Bit: # Discussion
	KINQUIRY_Byte6_BQUE_Bit KinquiryByte6 = 0
	// KINQUIRY_Byte6_MCHNGR_Mask: # Discussion
	KINQUIRY_Byte6_MCHNGR_Mask KinquiryByte6 = 0
)

func (e KinquiryByte6) String() string {
	switch e {
	case KINQUIRY_Byte6_BQUE_Bit:
		return "KINQUIRY_Byte6_BQUE_Bit"
	default:
		return fmt.Sprintf("KinquiryByte6(%d)", e)
	}
}

type KinquiryByte7 uint

const (
	// KINQUIRY_Byte7_CMDQUE_Bit: # Discussion
	KINQUIRY_Byte7_CMDQUE_Bit KinquiryByte7 = 0
	// KINQUIRY_Byte7_SYNC_Mask: # Discussion
	KINQUIRY_Byte7_SYNC_Mask KinquiryByte7 = 0
)

func (e KinquiryByte7) String() string {
	switch e {
	case KINQUIRY_Byte7_CMDQUE_Bit:
		return "KINQUIRY_Byte7_CMDQUE_Bit"
	default:
		return fmt.Sprintf("KinquiryByte7(%d)", e)
	}
}

type KinquiryPage83 uint

const (
	// KINQUIRY_Page83_IdentifierTypeIEEE_EUI64: # Discussion
	KINQUIRY_Page83_IdentifierTypeIEEE_EUI64 KinquiryPage83 = 0
	// KINQUIRY_Page83_IdentifierTypeLogicalUnitGroup: # Discussion
	KINQUIRY_Page83_IdentifierTypeLogicalUnitGroup KinquiryPage83 = 0
	// KINQUIRY_Page83_IdentifierTypeMD5LogicalUnitIdentifier: # Discussion
	KINQUIRY_Page83_IdentifierTypeMD5LogicalUnitIdentifier KinquiryPage83 = 0
	// KINQUIRY_Page83_IdentifierTypeMask: # Discussion
	KINQUIRY_Page83_IdentifierTypeMask KinquiryPage83 = 0
	// KINQUIRY_Page83_IdentifierTypeNAAIdentifier: # Discussion
	KINQUIRY_Page83_IdentifierTypeNAAIdentifier KinquiryPage83 = 0
	// KINQUIRY_Page83_IdentifierTypeRelativePortIdentifier: # Discussion
	KINQUIRY_Page83_IdentifierTypeRelativePortIdentifier KinquiryPage83 = 0
	// KINQUIRY_Page83_IdentifierTypeSCSINameString: # Discussion
	KINQUIRY_Page83_IdentifierTypeSCSINameString KinquiryPage83 = 0
	// KINQUIRY_Page83_IdentifierTypeTargetPortGroup: # Discussion
	KINQUIRY_Page83_IdentifierTypeTargetPortGroup KinquiryPage83 = 0
	// KINQUIRY_Page83_IdentifierTypeVendorID: # Discussion
	KINQUIRY_Page83_IdentifierTypeVendorID KinquiryPage83 = 0
	// KINQUIRY_Page83_IdentifierTypeVendorSpecific: # Discussion
	KINQUIRY_Page83_IdentifierTypeVendorSpecific KinquiryPage83 = 0
	// KINQUIRY_Page83_ProtocolIdentifierValidBit: # Discussion
	KINQUIRY_Page83_ProtocolIdentifierValidBit KinquiryPage83 = 0
	// KINQUIRY_Page83_ProtocolIdentifierValidMask: # Discussion
	KINQUIRY_Page83_ProtocolIdentifierValidMask KinquiryPage83 = 0
)

func (e KinquiryPage83) String() string {
	switch e {
	case KINQUIRY_Page83_IdentifierTypeIEEE_EUI64:
		return "KINQUIRY_Page83_IdentifierTypeIEEE_EUI64"
	default:
		return fmt.Sprintf("KinquiryPage83(%d)", e)
	}
}

type KinquiryPage83Association uint

const (
	// KINQUIRY_Page83_AssociationDevice: # Discussion
	KINQUIRY_Page83_AssociationDevice KinquiryPage83Association = 0
	KINQUIRY_Page83_AssociationShift KinquiryPage83Association = 0
)

func (e KinquiryPage83Association) String() string {
	switch e {
	case KINQUIRY_Page83_AssociationDevice:
		return "KINQUIRY_Page83_AssociationDevice"
	default:
		return fmt.Sprintf("KinquiryPage83Association(%d)", e)
	}
}

type KinquiryPage83Codeset uint

const (
	KINQUIRY_Page83_CodeSetMask KinquiryPage83Codeset = 0
	KINQUIRY_Page83_CodeSetReserved KinquiryPage83Codeset = 0
)

func (e KinquiryPage83Codeset) String() string {
	switch e {
	case KINQUIRY_Page83_CodeSetMask:
		return "KINQUIRY_Page83_CodeSetMask"
	default:
		return fmt.Sprintf("KinquiryPage83Codeset(%d)", e)
	}
}

type KinquiryPageb1PageLength uint

const (
	KINQUIRY_PageB1_Page_Length KinquiryPageb1PageLength = 0
)

func (e KinquiryPageb1PageLength) String() string {
	switch e {
	case KINQUIRY_PageB1_Page_Length:
		return "KINQUIRY_PageB1_Page_Length"
	default:
		return fmt.Sprintf("KinquiryPageb1PageLength(%d)", e)
	}
}

type KinquiryPagec0FeaturesHassepLun uint

const (
	KINQUIRY_PageC0_Features_HasSEP_LUN KinquiryPagec0FeaturesHassepLun = 0
)

func (e KinquiryPagec0FeaturesHassepLun) String() string {
	switch e {
	case KINQUIRY_PageC0_Features_HasSEP_LUN:
		return "KINQUIRY_PageC0_Features_HasSEP_LUN"
	default:
		return fmt.Sprintf("KinquiryPagec0FeaturesHassepLun(%d)", e)
	}
}

type KinquiryPeripheralQualifierConnected uint

const (
	// KINQUIRY_PERIPHERAL_QUALIFIER_Connected: # Discussion
	KINQUIRY_PERIPHERAL_QUALIFIER_Connected KinquiryPeripheralQualifierConnected = 0
)

func (e KinquiryPeripheralQualifierConnected) String() string {
	switch e {
	case KINQUIRY_PERIPHERAL_QUALIFIER_Connected:
		return "KINQUIRY_PERIPHERAL_QUALIFIER_Connected"
	default:
		return fmt.Sprintf("KinquiryPeripheralQualifierConnected(%d)", e)
	}
}

type KinquiryPeripheralRmb uint

const (
	// KINQUIRY_PERIPHERAL_RMB_BitMask: # Discussion
	KINQUIRY_PERIPHERAL_RMB_BitMask KinquiryPeripheralRmb = 0
	// KINQUIRY_PERIPHERAL_RMB_MediumFixed: # Discussion
	KINQUIRY_PERIPHERAL_RMB_MediumFixed KinquiryPeripheralRmb = 0
)

func (e KinquiryPeripheralRmb) String() string {
	switch e {
	case KINQUIRY_PERIPHERAL_RMB_BitMask:
		return "KINQUIRY_PERIPHERAL_RMB_BitMask"
	default:
		return fmt.Sprintf("KinquiryPeripheralRmb(%d)", e)
	}
}

type KinquiryPeripheralType uint

const (
	// KINQUIRY_PERIPHERAL_TYPE_CDROM_MMCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_CDROM_MMCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_ProcessorSPCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_ProcessorSPCDevice KinquiryPeripheralType = 0
)

func (e KinquiryPeripheralType) String() string {
	switch e {
	case KINQUIRY_PERIPHERAL_TYPE_CDROM_MMCDevice:
		return "KINQUIRY_PERIPHERAL_TYPE_CDROM_MMCDevice"
	default:
		return fmt.Sprintf("KinquiryPeripheralType(%d)", e)
	}
}

type KinquiryVendorIdentificationLength uint

const (
	// KINQUIRY_PRODUCT_IDENTIFICATION_Length: # Discussion
	KINQUIRY_PRODUCT_IDENTIFICATION_Length KinquiryVendorIdentificationLength = 0
)

func (e KinquiryVendorIdentificationLength) String() string {
	switch e {
	case KINQUIRY_PRODUCT_IDENTIFICATION_Length:
		return "KINQUIRY_PRODUCT_IDENTIFICATION_Length"
	default:
		return fmt.Sprintf("KinquiryVendorIdentificationLength(%d)", e)
	}
}

type Kioanalogsignallevel07000300 uint

const (
	KIOAnalogSignalLevel_0700_0000 Kioanalogsignallevel07000300 = 0
)

func (e Kioanalogsignallevel07000300) String() string {
	switch e {
	case KIOAnalogSignalLevel_0700_0000:
		return "KIOAnalogSignalLevel_0700_0000"
	default:
		return fmt.Sprintf("Kioanalogsignallevel07000300(%d)", e)
	}
}

type KioaudiochannellabelUnknown uint

const (
	KIOAudioChannelLabel_Ambisonic_W KioaudiochannellabelUnknown = 0
)

func (e KioaudiochannellabelUnknown) String() string {
	switch e {
	case KIOAudioChannelLabel_Ambisonic_W:
		return "KIOAudioChannelLabel_Ambisonic_W"
	default:
		return fmt.Sprintf("KioaudiochannellabelUnknown(%d)", e)
	}
}

type KiofbhdcplimitAllowall uint

const (
	KIOFBHDCPLimit_AllowAll KiofbhdcplimitAllowall = 0
)

func (e KiofbhdcplimitAllowall) String() string {
	switch e {
	case KIOFBHDCPLimit_AllowAll:
		return "KIOFBHDCPLimit_AllowAll"
	default:
		return fmt.Sprintf("KiofbhdcplimitAllowall(%d)", e)
	}
}

type KiofbnotifypriorityMin uint

const (
	KIOFBNotifyPriority_Max KiofbnotifypriorityMin = 0
)

func (e KiofbnotifypriorityMin) String() string {
	switch e {
	case KIOFBNotifyPriority_Max:
		return "KIOFBNotifyPriority_Max"
	default:
		return fmt.Sprintf("KiofbnotifypriorityMin(%d)", e)
	}
}

type KiofwspecidAapl uint

const (
	KIOFWSWVers_KPF KiofwspecidAapl = 0
)

func (e KiofwspecidAapl) String() string {
	switch e {
	case KIOFWSWVers_KPF:
		return "KIOFWSWVers_KPF"
	default:
		return fmt.Sprintf("KiofwspecidAapl(%d)", e)
	}
}

type KiovideodevicenotificationidControl uint

const (
	// KIOVideoDeviceNotificationID_ControlRangeChanged: # Discussion
	KIOVideoDeviceNotificationID_ControlRangeChanged KiovideodevicenotificationidControl = 0
	// KIOVideoDeviceNotificationID_ControlValueChanged: # Discussion
	KIOVideoDeviceNotificationID_ControlValueChanged KiovideodevicenotificationidControl = 0
)

func (e KiovideodevicenotificationidControl) String() string {
	switch e {
	case KIOVideoDeviceNotificationID_ControlRangeChanged:
		return "KIOVideoDeviceNotificationID_ControlRangeChanged"
	default:
		return fmt.Sprintf("KiovideodevicenotificationidControl(%d)", e)
	}
}

type KmmccmdBlank uint

const (
	KMMCCmd_BLANK KmmccmdBlank = 0
)

func (e KmmccmdBlank) String() string {
	switch e {
	case KMMCCmd_BLANK:
		return "KMMCCmd_BLANK"
	default:
		return fmt.Sprintf("KmmccmdBlank(%d)", e)
	}
}

type KmodepageformatPsBit uint

const (
	// KModePageFormat_PS_Bit: # Discussion
	KModePageFormat_PS_Bit KmodepageformatPsBit = 0
)

func (e KmodepageformatPsBit) String() string {
	switch e {
	case KModePageFormat_PS_Bit:
		return "KModePageFormat_PS_Bit"
	default:
		return fmt.Sprintf("KmodepageformatPsBit(%d)", e)
	}
}

type Kmodesenseparameterheader10Longlbabit uint

const (
	// KModeSenseParameterHeader10_LongLBABit: # Discussion
	KModeSenseParameterHeader10_LongLBABit Kmodesenseparameterheader10Longlbabit = 0
)

func (e Kmodesenseparameterheader10Longlbabit) String() string {
	switch e {
	case KModeSenseParameterHeader10_LongLBABit:
		return "KModeSenseParameterHeader10_LongLBABit"
	default:
		return fmt.Sprintf("Kmodesenseparameterheader10Longlbabit(%d)", e)
	}
}

type KmodifierDidperformmodifiy uint

const (
	KModifier_DidKeyUp KmodifierDidperformmodifiy = 0
)

func (e KmodifierDidperformmodifiy) String() string {
	switch e {
	case KModifier_DidKeyUp:
		return "KModifier_DidKeyUp"
	default:
		return fmt.Sprintf("KmodifierDidperformmodifiy(%d)", e)
	}
}

type KperfKpcFlagsT uint

const (
	KPC_CAPTURED_PC KperfKpcFlagsT = 0
)

func (e KperfKpcFlagsT) String() string {
	switch e {
	case KPC_CAPTURED_PC:
		return "KPC_CAPTURED_PC"
	default:
		return fmt.Sprintf("KperfKpcFlagsT(%d)", e)
	}
}

type Krangesupportssignal07000000 uint

const (
	KRangeSupportsSignal_0700_0000_Bit Krangesupportssignal07000000 = 0
	KRangeSupportsSignal_0700_0000_Mask Krangesupportssignal07000000 = 0
)

func (e Krangesupportssignal07000000) String() string {
	switch e {
	case KRangeSupportsSignal_0700_0000_Bit:
		return "KRangeSupportsSignal_0700_0000_Bit"
	default:
		return fmt.Sprintf("Krangesupportssignal07000000(%d)", e)
	}
}

type KreadCapacityProtEnabled uint

const (
	// KREAD_CAPACITY_PROT_Disabled: # Discussion
	KREAD_CAPACITY_PROT_Disabled KreadCapacityProtEnabled = 0
)

func (e KreadCapacityProtEnabled) String() string {
	switch e {
	case KREAD_CAPACITY_PROT_Disabled:
		return "KREAD_CAPACITY_PROT_Disabled"
	default:
		return fmt.Sprintf("KreadCapacityProtEnabled(%d)", e)
	}
}

type KreadCapacityRto uint

const (
	// KREAD_CAPACITY_RTO_Disabled: # Discussion
	KREAD_CAPACITY_RTO_Disabled KreadCapacityRto = 0
	// KREAD_CAPACITY_RTO_Mask: # Discussion
	KREAD_CAPACITY_RTO_Mask KreadCapacityRto = 0
)

func (e KreadCapacityRto) String() string {
	switch e {
	case KREAD_CAPACITY_RTO_Disabled:
		return "KREAD_CAPACITY_RTO_Disabled"
	default:
		return fmt.Sprintf("KreadCapacityRto(%d)", e)
	}
}

type KreportCapacityDatasize uint

const (
	// KREPORT_CAPACITY_16_DataSize: # Discussion
	KREPORT_CAPACITY_16_DataSize KreportCapacityDatasize = 0
)

func (e KreportCapacityDatasize) String() string {
	switch e {
	case KREPORT_CAPACITY_16_DataSize:
		return "KREPORT_CAPACITY_16_DataSize"
	default:
		return fmt.Sprintf("KreportCapacityDatasize(%d)", e)
	}
}

type KreportLunsAddress uint

const (
	// KREPORT_LUNS_ADDRESS_DEVICE_TYPE_SPECIFIC: # Discussion
	KREPORT_LUNS_ADDRESS_DEVICE_TYPE_SPECIFIC KreportLunsAddress = 0
	// KREPORT_LUNS_ADDRESS_METHOD_LOGICAL_UNIT: # Discussion
	KREPORT_LUNS_ADDRESS_METHOD_LOGICAL_UNIT KreportLunsAddress = 0
)

func (e KreportLunsAddress) String() string {
	switch e {
	case KREPORT_LUNS_ADDRESS_DEVICE_TYPE_SPECIFIC:
		return "KREPORT_LUNS_ADDRESS_DEVICE_TYPE_SPECIFIC"
	default:
		return fmt.Sprintf("KreportLunsAddress(%d)", e)
	}
}

type KsbcmodepagecachingDemandWriteMask uint

const (
	// KSBCModePageCaching_DEMAND_READ_Mask: # Discussion
	KSBCModePageCaching_DEMAND_READ_Mask KsbcmodepagecachingDemandWriteMask = 0
)

func (e KsbcmodepagecachingDemandWriteMask) String() string {
	switch e {
	case KSBCModePageCaching_DEMAND_READ_Mask:
		return "KSBCModePageCaching_DEMAND_READ_Mask"
	default:
		return fmt.Sprintf("KsbcmodepagecachingDemandWriteMask(%d)", e)
	}
}

type KsbcmodepagecachingVs1Bit uint

const (
	// KSBCModePageCaching_DRA_Bit: # Discussion
	KSBCModePageCaching_DRA_Bit KsbcmodepagecachingVs1Bit = 0
)

func (e KsbcmodepagecachingVs1Bit) String() string {
	switch e {
	case KSBCModePageCaching_DRA_Bit:
		return "KSBCModePageCaching_DRA_Bit"
	default:
		return fmt.Sprintf("KsbcmodepagecachingVs1Bit(%d)", e)
	}
}

type KsbcmodepageflexiblediskPin1Mask uint

const (
	// KSBCModePageFlexibleDisk_PIN_1_Mask: # Discussion
	KSBCModePageFlexibleDisk_PIN_1_Mask KsbcmodepageflexiblediskPin1Mask = 0
)

func (e KsbcmodepageflexiblediskPin1Mask) String() string {
	switch e {
	case KSBCModePageFlexibleDisk_PIN_1_Mask:
		return "KSBCModePageFlexibleDisk_PIN_1_Mask"
	default:
		return fmt.Sprintf("KsbcmodepageflexiblediskPin1Mask(%d)", e)
	}
}

type KsbcmodepageflexiblediskPin2Mask uint

const (
	// KSBCModePageFlexibleDisk_PIN_2_Mask: # Discussion
	KSBCModePageFlexibleDisk_PIN_2_Mask KsbcmodepageflexiblediskPin2Mask = 0
)

func (e KsbcmodepageflexiblediskPin2Mask) String() string {
	switch e {
	case KSBCModePageFlexibleDisk_PIN_2_Mask:
		return "KSBCModePageFlexibleDisk_PIN_2_Mask"
	default:
		return fmt.Sprintf("KsbcmodepageflexiblediskPin2Mask(%d)", e)
	}
}

type KsbcmodepageflexiblediskSpcMask uint

const (
	// KSBCModePageFlexibleDisk_SPC_Mask: # Discussion
	KSBCModePageFlexibleDisk_SPC_Mask KsbcmodepageflexiblediskSpcMask = 0
)

func (e KsbcmodepageflexiblediskSpcMask) String() string {
	switch e {
	case KSBCModePageFlexibleDisk_SPC_Mask:
		return "KSBCModePageFlexibleDisk_SPC_Mask"
	default:
		return fmt.Sprintf("KsbcmodepageflexiblediskSpcMask(%d)", e)
	}
}

type KsbcmodepagerigiddiskgeometryRplMask uint

const (
	// KSBCModePageRigidDiskGeometry_RPL_Mask: # Discussion
	KSBCModePageRigidDiskGeometry_RPL_Mask KsbcmodepagerigiddiskgeometryRplMask = 0
)

func (e KsbcmodepagerigiddiskgeometryRplMask) String() string {
	switch e {
	case KSBCModePageRigidDiskGeometry_RPL_Mask:
		return "KSBCModePageRigidDiskGeometry_RPL_Mask"
	default:
		return fmt.Sprintf("KsbcmodepagerigiddiskgeometryRplMask(%d)", e)
	}
}

type KscccmdMaintenanceIn uint

const (
	KSCCCmd_MAINTENANCE_IN KscccmdMaintenanceIn = 0
)

func (e KscccmdMaintenanceIn) String() string {
	switch e {
	case KSCCCmd_MAINTENANCE_IN:
		return "KSCCCmd_MAINTENANCE_IN"
	default:
		return fmt.Sprintf("KscccmdMaintenanceIn(%d)", e)
	}
}

type KscsiportNotificationstatuschange uint

const (
	KSCSIPort_NotificationStatusChange KscsiportNotificationstatuschange = 0
)

func (e KscsiportNotificationstatuschange) String() string {
	switch e {
	case KSCSIPort_NotificationStatusChange:
		return "KSCSIPort_NotificationStatusChange"
	default:
		return fmt.Sprintf("KscsiportNotificationstatuschange(%d)", e)
	}
}

type KscsiportStatusonline int

const (
	// KSCSIPort_StatusFailure: # Discussion
	KSCSIPort_StatusFailure KscsiportStatusonline = 0
)

func (e KscsiportStatusonline) String() string {
	switch e {
	case KSCSIPort_StatusFailure:
		return "KSCSIPort_StatusFailure"
	default:
		return fmt.Sprintf("KscsiportStatusonline(%d)", e)
	}
}

type KscsiprotocolnotificationDeviceremoved uint

const (
	// KSCSIProtocolNotification_DeviceRemoved: # Discussion
	KSCSIProtocolNotification_DeviceRemoved KscsiprotocolnotificationDeviceremoved = 0
)

func (e KscsiprotocolnotificationDeviceremoved) String() string {
	switch e {
	case KSCSIProtocolNotification_DeviceRemoved:
		return "KSCSIProtocolNotification_DeviceRemoved"
	default:
		return fmt.Sprintf("KscsiprotocolnotificationDeviceremoved(%d)", e)
	}
}

type KscsiserviceactionChangeAliases uint

const (
	KSCSIServiceAction_CHANGE_ALIASES KscsiserviceactionChangeAliases = 0
)

func (e KscsiserviceactionChangeAliases) String() string {
	switch e {
	case KSCSIServiceAction_CHANGE_ALIASES:
		return "KSCSIServiceAction_CHANGE_ALIASES"
	default:
		return fmt.Sprintf("KscsiserviceactionChangeAliases(%d)", e)
	}
}

type KscsiserviceactionGetLbaStatus int

const (
	KSCSIServiceAction_GET_LBA_STATUS KscsiserviceactionGetLbaStatus = 0
)

func (e KscsiserviceactionGetLbaStatus) String() string {
	switch e {
	case KSCSIServiceAction_GET_LBA_STATUS:
		return "KSCSIServiceAction_GET_LBA_STATUS"
	default:
		return fmt.Sprintf("KscsiserviceactionGetLbaStatus(%d)", e)
	}
}

type KscsiserviceactionRead32 uint

const (
	KSCSIServiceAction_READ_32 KscsiserviceactionRead32 = 0
)

func (e KscsiserviceactionRead32) String() string {
	switch e {
	case KSCSIServiceAction_READ_32:
		return "KSCSIServiceAction_READ_32"
	default:
		return fmt.Sprintf("KscsiserviceactionRead32(%d)", e)
	}
}

type KscsiserviceactionReportAliases uint

const (
	KSCSIServiceAction_REPORT_ALIASES KscsiserviceactionReportAliases = 0
)

func (e KscsiserviceactionReportAliases) String() string {
	switch e {
	case KSCSIServiceAction_REPORT_ALIASES:
		return "KSCSIServiceAction_REPORT_ALIASES"
	default:
		return fmt.Sprintf("KscsiserviceactionReportAliases(%d)", e)
	}
}

type KscsiserviceactionWriteLong16 uint

const (
	KSCSIServiceAction_WRITE_LONG_16 KscsiserviceactionWriteLong16 = 0
)

func (e KscsiserviceactionWriteLong16) String() string {
	switch e {
	case KSCSIServiceAction_WRITE_LONG_16:
		return "KSCSIServiceAction_WRITE_LONG_16"
	default:
		return fmt.Sprintf("KscsiserviceactionWriteLong16(%d)", e)
	}
}

type KscsiservicesnotificationSuspend uint

const (
	KSCSIServicesNotification_Resume KscsiservicesnotificationSuspend = 0
)

func (e KscsiservicesnotificationSuspend) String() string {
	switch e {
	case KSCSIServicesNotification_Resume:
		return "KSCSIServicesNotification_Resume"
	default:
		return fmt.Sprintf("KscsiservicesnotificationSuspend(%d)", e)
	}
}

type KsenseDataValid uint

const (
	// KSENSE_NOT_DATA_VALID: # Discussion
	KSENSE_NOT_DATA_VALID KsenseDataValid = 0
)

func (e KsenseDataValid) String() string {
	switch e {
	case KSENSE_NOT_DATA_VALID:
		return "KSENSE_NOT_DATA_VALID"
	default:
		return fmt.Sprintf("KsenseDataValid(%d)", e)
	}
}

type KsenseEomSet uint

const (
	// KSENSE_EOM_Mask: # Discussion
	KSENSE_EOM_Mask KsenseEomSet = 0
)

func (e KsenseEomSet) String() string {
	switch e {
	case KSENSE_EOM_Mask:
		return "KSENSE_EOM_Mask"
	default:
		return fmt.Sprintf("KsenseEomSet(%d)", e)
	}
}

type KsenseFilemarkSet uint

const (
	// KSENSE_FILEMARK_Mask: # Discussion
	KSENSE_FILEMARK_Mask KsenseFilemarkSet = 0
)

func (e KsenseFilemarkSet) String() string {
	switch e {
	case KSENSE_FILEMARK_Mask:
		return "KSENSE_FILEMARK_Mask"
	default:
		return fmt.Sprintf("KsenseFilemarkSet(%d)", e)
	}
}

type KsenseIliSet uint

const (
	// KSENSE_ILI_Mask: # Discussion
	KSENSE_ILI_Mask KsenseIliSet = 0
)

func (e KsenseIliSet) String() string {
	switch e {
	case KSENSE_ILI_Mask:
		return "KSENSE_ILI_Mask"
	default:
		return fmt.Sprintf("KsenseIliSet(%d)", e)
	}
}

type KsenseKey uint

const (
	// KSENSE_KEY_ABORTED_COMMAND: # Discussion
	KSENSE_KEY_ABORTED_COMMAND KsenseKey = 0
	// KSENSE_KEY_BLANK_CHECK: # Discussion
	KSENSE_KEY_BLANK_CHECK KsenseKey = 0
	// KSENSE_KEY_COPY_ABORTED: # Discussion
	KSENSE_KEY_COPY_ABORTED KsenseKey = 0
	// KSENSE_KEY_DATA_PROTECT: # Discussion
	KSENSE_KEY_DATA_PROTECT KsenseKey = 0
	// KSENSE_KEY_HARDWARE_ERROR: # Discussion
	KSENSE_KEY_HARDWARE_ERROR KsenseKey = 0
	// KSENSE_KEY_ILLEGAL_REQUEST: # Discussion
	KSENSE_KEY_ILLEGAL_REQUEST KsenseKey = 0
	// KSENSE_KEY_MEDIUM_ERROR: # Discussion
	KSENSE_KEY_MEDIUM_ERROR KsenseKey = 0
	// KSENSE_KEY_MISCOMPARE: # Discussion
	KSENSE_KEY_MISCOMPARE KsenseKey = 0
	// KSENSE_KEY_Mask: # Discussion
	KSENSE_KEY_Mask KsenseKey = 0
	// KSENSE_KEY_NOT_READY: # Discussion
	KSENSE_KEY_NOT_READY KsenseKey = 0
	// KSENSE_KEY_NO_SENSE: # Discussion
	KSENSE_KEY_NO_SENSE KsenseKey = 0
	// KSENSE_KEY_RECOVERED_ERROR: # Discussion
	KSENSE_KEY_RECOVERED_ERROR KsenseKey = 0
	// KSENSE_KEY_UNIT_ATTENTION: # Discussion
	KSENSE_KEY_UNIT_ATTENTION KsenseKey = 0
	// KSENSE_KEY_VENDOR_SPECIFIC: # Discussion
	KSENSE_KEY_VENDOR_SPECIFIC KsenseKey = 0
	// KSENSE_KEY_VOLUME_OVERFLOW: # Discussion
	KSENSE_KEY_VOLUME_OVERFLOW KsenseKey = 0
)

func (e KsenseKey) String() string {
	switch e {
	case KSENSE_KEY_ABORTED_COMMAND:
		return "KSENSE_KEY_ABORTED_COMMAND"
	default:
		return fmt.Sprintf("KsenseKey(%d)", e)
	}
}

type KsenseResponseCodeCurrentErrors int

const (
	// KSENSE_RESPONSE_CODE_Current_Errors: # Discussion
	KSENSE_RESPONSE_CODE_Current_Errors KsenseResponseCodeCurrentErrors = 0
)

func (e KsenseResponseCodeCurrentErrors) String() string {
	switch e {
	case KSENSE_RESPONSE_CODE_Current_Errors:
		return "KSENSE_RESPONSE_CODE_Current_Errors"
	default:
		return fmt.Sprintf("KsenseResponseCodeCurrentErrors(%d)", e)
	}
}

type KsescmdModeSelect6 uint

const (
	KSESCmd_MODE_SELECT_6 KsescmdModeSelect6 = 0
)

func (e KsescmdModeSelect6) String() string {
	switch e {
	case KSESCmd_MODE_SELECT_6:
		return "KSESCmd_MODE_SELECT_6"
	default:
		return fmt.Sprintf("KsescmdModeSelect6(%d)", e)
	}
}

type KsgccmdC uint

const (
	KSGCCmd_CHANGE_DEFINITION KsgccmdC = 0
	KSGCCmd_COPY_AND_VERIFY KsgccmdC = 0
)

func (e KsgccmdC) String() string {
	switch e {
	case KSGCCmd_CHANGE_DEFINITION:
		return "KSGCCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KsgccmdC(%d)", e)
	}
}

type KspcproccmdChangeDefinition uint

const (
	KSPCProcCmd_CHANGE_DEFINITION KspcproccmdChangeDefinition = 0
)

func (e KspcproccmdChangeDefinition) String() string {
	switch e {
	case KSPCProcCmd_CHANGE_DEFINITION:
		return "KSPCProcCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KspcproccmdChangeDefinition(%d)", e)
	}
}

type KsscprintercmdChangeDefinition uint

const (
	KSSCPrinterCmd_CHANGE_DEFINITION KsscprintercmdChangeDefinition = 0
)

func (e KsscprintercmdChangeDefinition) String() string {
	switch e {
	case KSSCPrinterCmd_CHANGE_DEFINITION:
		return "KSSCPrinterCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KsscprintercmdChangeDefinition(%d)", e)
	}
}

type KstateDisabledFlag uint

const (
	KState_Disabled_Flag KstateDisabledFlag = 0
)

func (e KstateDisabledFlag) String() string {
	switch e {
	case KState_Disabled_Flag:
		return "KState_Disabled_Flag"
	default:
		return fmt.Sprintf("KstateDisabledFlag(%d)", e)
	}
}

type KusbSscompdescBulkMaxstreamsMask uint

const (
	KUSB_SSCompDesc_Bulk_MaxStreams_Mask KusbSscompdescBulkMaxstreamsMask = 0
)

func (e KusbSscompdescBulkMaxstreamsMask) String() string {
	switch e {
	case KUSB_SSCompDesc_Bulk_MaxStreams_Mask:
		return "KUSB_SSCompDesc_Bulk_MaxStreams_Mask"
	default:
		return fmt.Sprintf("KusbSscompdescBulkMaxstreamsMask(%d)", e)
	}
}

type KusbspeedMask uint

const (
	KUSBAddress_Mask KusbspeedMask = 0
)

func (e KusbspeedMask) String() string {
	switch e {
	case KUSBAddress_Mask:
		return "KUSBAddress_Mask"
	default:
		return fmt.Sprintf("KusbspeedMask(%d)", e)
	}
}

type LatencyQosTier uint

const (
	LATENCY_QOS_TIER_0 LatencyQosTier = 0
	LATENCY_QOS_TIER_1 LatencyQosTier = 0
	LATENCY_QOS_TIER_2 LatencyQosTier = 0
	LATENCY_QOS_TIER_3 LatencyQosTier = 0
	LATENCY_QOS_TIER_4 LatencyQosTier = 0
	LATENCY_QOS_TIER_5 LatencyQosTier = 0
	LATENCY_QOS_TIER_UNSPECIFIED LatencyQosTier = 0
)

func (e LatencyQosTier) String() string {
	switch e {
	case LATENCY_QOS_TIER_0:
		return "LATENCY_QOS_TIER_0"
	default:
		return fmt.Sprintf("LatencyQosTier(%d)", e)
	}
}

type LckSleepActionT uint

const (
	LCK_SLEEP_DEFAULT LckSleepActionT = 0
)

func (e LckSleepActionT) String() string {
	switch e {
	case LCK_SLEEP_DEFAULT:
		return "LCK_SLEEP_DEFAULT"
	default:
		return fmt.Sprintf("LckSleepActionT(%d)", e)
	}
}

type LckWakeActionT uint

const (
	LCK_WAKE_DEFAULT LckWakeActionT = 0
)

func (e LckWakeActionT) String() string {
	switch e {
	case LCK_WAKE_DEFAULT:
		return "LCK_WAKE_DEFAULT"
	default:
		return fmt.Sprintf("LckWakeActionT(%d)", e)
	}
}

type Libsptm uint

const (
	LIBSPTM_INVALID_ARG Libsptm = 0
	LIBSPTM_TYPE_MISMATCH Libsptm = 0
)

func (e Libsptm) String() string {
	switch e {
	case LIBSPTM_INVALID_ARG:
		return "LIBSPTM_INVALID_ARG"
	default:
		return fmt.Sprintf("Libsptm(%d)", e)
	}
}

type LibsptmCpuStateT uint

const (
	CPUSTATE_PANIC_SPIN LibsptmCpuStateT = 0
)

func (e LibsptmCpuStateT) String() string {
	switch e {
	case CPUSTATE_PANIC_SPIN:
		return "CPUSTATE_PANIC_SPIN"
	default:
		return fmt.Sprintf("LibsptmCpuStateT(%d)", e)
	}
}

type MATA uint

const (
	MATAAltSDevCValid MATA = 0
	MATACylinderHiValid MATA = 0
	MATACylinderLoValid MATA = 0
	MATADataValid MATA = 0
	MATAErrFeaturesValid MATA = 0
	MATAFlag48BitLBA MATA = 0
	MATAFlagByteSwap MATA = 0
	MATAFlagDMAQueued MATA = 0
	MATAFlagIORead MATA = 0
	MATAFlagIOWrite MATA = 0
	MATAFlagImmediate MATA = 0
	MATAFlagOverlapped MATA = 0
	MATAFlagProtocolATAPI MATA = 0
	MATAFlagQuiesce MATA = 0
	MATAFlagTFAccess MATA = 0
	MATAFlagTFAccessResult MATA = 0
	MATAFlagUseConfigSpeed MATA = 0
	MATAFlagUseDMA MATA = 0
	MATAFlagUseNoIRQ MATA = 0
	MATASDHValid MATA = 0
	MATASectorCntValid MATA = 0
	MATASectorNumValid MATA = 0
	MATAStatusCmdValid MATA = 0
)

func (e MATA) String() string {
	switch e {
	case MATAAltSDevCValid:
		return "MATAAltSDevCValid"
	default:
		return fmt.Sprintf("MATA(%d)", e)
	}
}

type MATAHeadNumber uint

const (
	MATADriveSelect MATAHeadNumber = 0
)

func (e MATAHeadNumber) String() string {
	switch e {
	case MATADriveSelect:
		return "MATADriveSelect"
	default:
		return fmt.Sprintf("MATAHeadNumber(%d)", e)
	}
}

type MBaseOffset uint

const (
	EightBitMode MBaseOffset = 0
	MCmpCount MBaseOffset = 0
)

func (e MBaseOffset) String() string {
	switch e {
	case EightBitMode:
		return "EightBitMode"
	default:
		return fmt.Sprintf("MBaseOffset(%d)", e)
	}
}

type MachAssertTypeT uint

const (
	MACH_ASSERT_3P MachAssertTypeT = 0
)

func (e MachAssertTypeT) String() string {
	switch e {
	case MACH_ASSERT_3P:
		return "MACH_ASSERT_3P"
	default:
		return fmt.Sprintf("MachAssertTypeT(%d)", e)
	}
}

type MachVmRangeFlagsT uint

const (
	MACH_VM_RANGE_NONE MachVmRangeFlagsT = 0
)

func (e MachVmRangeFlagsT) String() string {
	switch e {
	case MACH_VM_RANGE_NONE:
		return "MACH_VM_RANGE_NONE"
	default:
		return fmt.Sprintf("MachVmRangeFlagsT(%d)", e)
	}
}

type MachVmRangeFlavorT uint

const (
	MACH_VM_RANGE_FLAVOR_INVALID MachVmRangeFlavorT = 0
)

func (e MachVmRangeFlavorT) String() string {
	switch e {
	case MACH_VM_RANGE_FLAVOR_INVALID:
		return "MACH_VM_RANGE_FLAVOR_INVALID"
	default:
		return fmt.Sprintf("MachVmRangeFlavorT(%d)", e)
	}
}

type MachVmRangeTagT uint

const (
	MACH_VM_RANGE_DATA MachVmRangeTagT = 0
)

func (e MachVmRangeTagT) String() string {
	switch e {
	case MACH_VM_RANGE_DATA:
		return "MACH_VM_RANGE_DATA"
	default:
		return fmt.Sprintf("MachVmRangeTagT(%d)", e)
	}
}

type MacosPanicHeaderFlagCore uint

const (
	MACOS_PANIC_HEADER_FLAG_COREDUMP_FAILED MacosPanicHeaderFlagCore = 0
	MACOS_PANIC_HEADER_FLAG_COREFILE_UNLINKED MacosPanicHeaderFlagCore = 0
)

func (e MacosPanicHeaderFlagCore) String() string {
	switch e {
	case MACOS_PANIC_HEADER_FLAG_COREDUMP_FAILED:
		return "MACOS_PANIC_HEADER_FLAG_COREDUMP_FAILED"
	default:
		return fmt.Sprintf("MacosPanicHeaderFlagCore(%d)", e)
	}
}

type MbufCsum uint

const (
	// MBUF_CSUM_DID_DATA: # Discussion
	MBUF_CSUM_DID_DATA MbufCsum = 0
	// MBUF_CSUM_PSEUDO_HDR: # Discussion
	MBUF_CSUM_PSEUDO_HDR MbufCsum = 0
)

func (e MbufCsum) String() string {
	switch e {
	case MBUF_CSUM_DID_DATA:
		return "MBUF_CSUM_DID_DATA"
	default:
		return fmt.Sprintf("MbufCsum(%d)", e)
	}
}

type MbufCsumReqIp uint

const (
	// MBUF_CSUM_REQ_IP: # Discussion
	MBUF_CSUM_REQ_IP MbufCsumReqIp = 0
)

func (e MbufCsumReqIp) String() string {
	switch e {
	case MBUF_CSUM_REQ_IP:
		return "MBUF_CSUM_REQ_IP"
	default:
		return fmt.Sprintf("MbufCsumReqIp(%d)", e)
	}
}

type MbufExt uint

const (
	// MBUF_BCAST: # Discussion
	MBUF_BCAST MbufExt = 0
)

func (e MbufExt) String() string {
	switch e {
	case MBUF_BCAST:
		return "MBUF_BCAST"
	default:
		return fmt.Sprintf("MbufExt(%d)", e)
	}
}

type MbufTc uint

const (
	// MBUF_TC_BE: # Discussion
	MBUF_TC_BE MbufTc = 0
	// MBUF_TC_VO: # Discussion
	MBUF_TC_VO MbufTc = 0
)

func (e MbufTc) String() string {
	switch e {
	case MBUF_TC_BE:
		return "MBUF_TC_BE"
	default:
		return fmt.Sprintf("MbufTc(%d)", e)
	}
}

type MbufTsoIpv4 uint

const (
	MBUF_TSO_IPV4 MbufTsoIpv4 = 0
)

func (e MbufTsoIpv4) String() string {
	switch e {
	case MBUF_TSO_IPV4:
		return "MBUF_TSO_IPV4"
	default:
		return fmt.Sprintf("MbufTsoIpv4(%d)", e)
	}
}

type MbufType uint

const (
	MBUF_TYPE_ATABLE MbufType = 0
	MBUF_TYPE_SOCKET MbufType = 0
)

func (e MbufType) String() string {
	switch e {
	case MBUF_TYPE_ATABLE:
		return "MBUF_TYPE_ATABLE"
	default:
		return fmt.Sprintf("MbufType(%d)", e)
	}
}

type MbufWaitok uint

const (
	// MBUF_DONTWAIT: # Discussion
	MBUF_DONTWAIT MbufWaitok = 0
)

func (e MbufWaitok) String() string {
	switch e {
	case MBUF_DONTWAIT:
		return "MBUF_DONTWAIT"
	default:
		return fmt.Sprintf("MbufWaitok(%d)", e)
	}
}

type MccFlagsT uint

const (
	MCC_IS_MULTI_BIT MccFlagsT = 0
)

func (e MccFlagsT) String() string {
	switch e {
	case MCC_IS_MULTI_BIT:
		return "MCC_IS_MULTI_BIT"
	default:
		return fmt.Sprintf("MccFlagsT(%d)", e)
	}
}

type MicroSnapshotFlags uint

const (
	KInterruptRecord MicroSnapshotFlags = 0
)

func (e MicroSnapshotFlags) String() string {
	switch e {
	case KInterruptRecord:
		return "KInterruptRecord"
	default:
		return fmt.Sprintf("MicroSnapshotFlags(%d)", e)
	}
}

type MicrostackshotFlagsT uint

const (
	STACKSHOT_GET_KERNEL_MICROSTACKSHOT MicrostackshotFlagsT = 0
)

func (e MicrostackshotFlagsT) String() string {
	switch e {
	case STACKSHOT_GET_KERNEL_MICROSTACKSHOT:
		return "STACKSHOT_GET_KERNEL_MICROSTACKSHOT"
	default:
		return fmt.Sprintf("MicrostackshotFlagsT(%d)", e)
	}
}

type Mtf uint

const (
	MTF_HTABLE Mtf = 0
	MTF_RIGHTS Mtf = 0
)

func (e Mtf) String() string {
	switch e {
	case MTF_HTABLE:
		return "MTF_HTABLE"
	default:
		return fmt.Sprintf("Mtf(%d)", e)
	}
}

type NVRAMPartitionType uint

const (
	KIONVRAMPartitionCommon NVRAMPartitionType = 0
)

func (e NVRAMPartitionType) String() string {
	switch e {
	case KIONVRAMPartitionCommon:
		return "KIONVRAMPartitionCommon"
	default:
		return fmt.Sprintf("NVRAMPartitionType(%d)", e)
	}
}

type NfsSupportedKerberosEtypes uint

const (
	NFS_AES128_CTS_HMAC_SHA1_96 NfsSupportedKerberosEtypes = 0
)

func (e NfsSupportedKerberosEtypes) String() string {
	switch e {
	case NFS_AES128_CTS_HMAC_SHA1_96:
		return "NFS_AES128_CTS_HMAC_SHA1_96"
	default:
		return fmt.Sprintf("NfsSupportedKerberosEtypes(%d)", e)
	}
}

type Nfstype uint

const (
	NFATTRDIR Nfstype = 0
)

func (e Nfstype) String() string {
	switch e {
	case NFATTRDIR:
		return "NFATTRDIR"
	default:
		return fmt.Sprintf("Nfstype(%d)", e)
	}
}

const NoErr int = 0

type NrLockedErr int

const (
	GestaltUndefSelectorErr NrLockedErr = 0
	NrLockedErrValue NrLockedErr = 0
)

func (e NrLockedErr) String() string {
	switch e {
	case GestaltUndefSelectorErr:
		return "GestaltUndefSelectorErr"
	default:
		return fmt.Sprintf("NrLockedErr(%d)", e)
	}
}

type Nx uint

const (
	NX_BigEndian Nx = 0
	NX_LeftButton Nx = 0
	NX_LittleEndian Nx = 0
	NX_OneButton Nx = 0
	NX_RightButton Nx = 0
	NX_UnknownByteOrder Nx = 0
)

func (e Nx) String() string {
	switch e {
	case NX_BigEndian:
		return "NX_BigEndian"
	default:
		return fmt.Sprintf("Nx(%d)", e)
	}
}

type Os uint

const (
	OSBigEndian Os = 0
	OSLittleEndian Os = 0
)

func (e Os) String() string {
	switch e {
	case OSBigEndian:
		return "OSBigEndian"
	default:
		return fmt.Sprintf("Os(%d)", e)
	}
}

type OsLogCoprocRegT uint

const (
	Os_log_coproc_register_harvest_fs_ftab OsLogCoprocRegT = 0
)

func (e OsLogCoprocRegT) String() string {
	switch e {
	case Os_log_coproc_register_harvest_fs_ftab:
		return "Os_log_coproc_register_harvest_fs_ftab"
	default:
		return fmt.Sprintf("OsLogCoprocRegT(%d)", e)
	}
}

type OsLogType uint

const (
	// OS_LOG_TYPE_DEBUG: Debug-level messages are only captured in memory when debug logging is enabled through a configuration change.
	OS_LOG_TYPE_DEBUG OsLogType = 0
	// OS_LOG_TYPE_ERROR: Error-level messages are always saved in the data store.
	OS_LOG_TYPE_ERROR OsLogType = 0
)

func (e OsLogType) String() string {
	switch e {
	case OS_LOG_TYPE_DEBUG:
		return "OS_LOG_TYPE_DEBUG"
	default:
		return fmt.Sprintf("OsLogType(%d)", e)
	}
}

type Output uint

const (
	OUTPUT_COMMUNICATION_SPEAKER Output = 0
	OUTPUT_UNDEFINED Output = 0
)

func (e Output) String() string {
	switch e {
	case OUTPUT_COMMUNICATION_SPEAKER:
		return "OUTPUT_COMMUNICATION_SPEAKER"
	default:
		return fmt.Sprintf("Output(%d)", e)
	}
}

type OutputNull uint

const (
	INPUT_NULL OutputNull = 0
	OUTPUT_NULL OutputNull = 0
)

func (e OutputNull) String() string {
	switch e {
	case INPUT_NULL:
		return "INPUT_NULL"
	default:
		return fmt.Sprintf("OutputNull(%d)", e)
	}
}

type P uint

const (
	P_ALL P = 0
	P_PID P = 0
)

func (e P) String() string {
	switch e {
	case P_ALL:
		return "P_ALL"
	default:
		return fmt.Sprintf("P(%d)", e)
	}
}

type ParamErr int

const (
	OpenErr ParamErr = 0
	UnitEmptyErr ParamErr = 0
)

func (e ParamErr) String() string {
	switch e {
	case OpenErr:
		return "OpenErr"
	default:
		return fmt.Sprintf("ParamErr(%d)", e)
	}
}

type PriorityQueueEntry uint

const (
	PRIORITY_QUEUE_ENTRY_NONE PriorityQueueEntry = 0
	PRIORITY_QUEUE_ENTRY_PREEMPTED PriorityQueueEntry = 0
)

func (e PriorityQueueEntry) String() string {
	switch e {
	case PRIORITY_QUEUE_ENTRY_NONE:
		return "PRIORITY_QUEUE_ENTRY_NONE"
	default:
		return fmt.Sprintf("PriorityQueueEntry(%d)", e)
	}
}

type Processor uint

const (
	PROCESSOR_GENERAL Processor = 0
	PROCESSOR_UNDEFINED Processor = 0
)

func (e Processor) String() string {
	switch e {
	case PROCESSOR_GENERAL:
		return "PROCESSOR_GENERAL"
	default:
		return fmt.Sprintf("Processor(%d)", e)
	}
}

type SDPAttributeIdentifierCodes uint

const (
	KBluetoothSDPAttributeIdentifierWAPStackType SDPAttributeIdentifierCodes = 0
)

func (e SDPAttributeIdentifierCodes) String() string {
	switch e {
	case KBluetoothSDPAttributeIdentifierWAPStackType:
		return "KBluetoothSDPAttributeIdentifierWAPStackType"
	default:
		return fmt.Sprintf("SDPAttributeIdentifierCodes(%d)", e)
	}
}

type SfltGlobal uint

const (
	// SFLT_EXTENDED: # Discussion
	SFLT_EXTENDED SfltGlobal = 0
)

func (e SfltGlobal) String() string {
	switch e {
	case SFLT_EXTENDED:
		return "SFLT_EXTENDED"
	default:
		return fmt.Sprintf("SfltGlobal(%d)", e)
	}
}

type SixteenBitMode uint

const (
	SecondVidMode SixteenBitMode = 0
)

func (e SixteenBitMode) String() string {
	switch e {
	case SecondVidMode:
		return "SecondVidMode"
	default:
		return fmt.Sprintf("SixteenBitMode(%d)", e)
	}
}

type SoTrackerAction uint

const (
	SO_TRACKER_ACTION_ADD SoTrackerAction = 0
	SO_TRACKER_ACTION_DUMP_ALL SoTrackerAction = 0
	SO_TRACKER_ACTION_DUMP_BY_APP SoTrackerAction = 0
	SO_TRACKER_ACTION_DUMP_MAX SoTrackerAction = 0
	SO_TRACKER_ACTION_INVALID SoTrackerAction = 0
)

func (e SoTrackerAction) String() string {
	switch e {
	case SO_TRACKER_ACTION_ADD:
		return "SO_TRACKER_ACTION_ADD"
	default:
		return fmt.Sprintf("SoTrackerAction(%d)", e)
	}
}

type SoTrackerAttributeA uint

const (
	SO_TRACKER_ATTRIBUTE_ADDRESS SoTrackerAttributeA = 0
	SO_TRACKER_ATTRIBUTE_ADDRESS_FAMILY SoTrackerAttributeA = 0
	SO_TRACKER_ATTRIBUTE_APP_UUID SoTrackerAttributeA = 0
)

func (e SoTrackerAttributeA) String() string {
	switch e {
	case SO_TRACKER_ATTRIBUTE_ADDRESS:
		return "SO_TRACKER_ATTRIBUTE_ADDRESS"
	default:
		return fmt.Sprintf("SoTrackerAttributeA(%d)", e)
	}
}

type SockDataFiltFlagOob uint

const (
	// Sock_data_filt_flag_oob: # Discussion
	Sock_data_filt_flag_oob SockDataFiltFlagOob = 0
)

func (e SockDataFiltFlagOob) String() string {
	switch e {
	case Sock_data_filt_flag_oob:
		return "Sock_data_filt_flag_oob"
	default:
		return fmt.Sprintf("SockDataFiltFlagOob(%d)", e)
	}
}

type SockEvtConnecting uint

const (
	// Sock_evt_bound: # Discussion
	Sock_evt_bound SockEvtConnecting = 0
)

func (e SockEvtConnecting) String() string {
	switch e {
	case Sock_evt_bound:
		return "Sock_evt_bound"
	default:
		return fmt.Sprintf("SockEvtConnecting(%d)", e)
	}
}

type SockoptGet uint

const (
	Sockopt_get SockoptGet = 0
)

func (e SockoptGet) String() string {
	switch e {
	case Sockopt_get:
		return "Sockopt_get"
	default:
		return fmt.Sprintf("SockoptGet(%d)", e)
	}
}

type Sptm uint

const (
	SPTM_N_VECTOR_TYPES Sptm = 0
	SPTM_VECTOR_FIQ Sptm = 0
)

func (e Sptm) String() string {
	switch e {
	case SPTM_N_VECTOR_TYPES:
		return "SPTM_N_VECTOR_TYPES"
	default:
		return fmt.Sprintf("Sptm(%d)", e)
	}
}

type SptmNTraces uint

const (
	SPTM_N_TRACES SptmNTraces = 0
)

func (e SptmNTraces) String() string {
	switch e {
	case SPTM_N_TRACES:
		return "SPTM_N_TRACES"
	default:
		return fmt.Sprintf("SptmNTraces(%d)", e)
	}
}

type SptmRefcount uint

const (
	SPTM_REFCOUNT_IOMMU SptmRefcount = 0
	SPTM_REFCOUNT_NONE SptmRefcount = 0
)

func (e SptmRefcount) String() string {
	switch e {
	case SPTM_REFCOUNT_IOMMU:
		return "SPTM_REFCOUNT_IOMMU"
	default:
		return fmt.Sprintf("SptmRefcount(%d)", e)
	}
}

type Stackshot uint

const (
	STACKSHOT_DO_COMPRESS Stackshot = 0
	STACKSHOT_PAGE_TABLES Stackshot = 0
)

func (e Stackshot) String() string {
	switch e {
	case STACKSHOT_DO_COMPRESS:
		return "STACKSHOT_DO_COMPRESS"
	default:
		return fmt.Sprintf("Stackshot(%d)", e)
	}
}

type TDeviceRequestType uint

const (
	KRequestTypeClass TDeviceRequestType = 0
)

func (e TDeviceRequestType) String() string {
	switch e {
	case KRequestTypeClass:
		return "KRequestTypeClass"
	default:
		return fmt.Sprintf("TDeviceRequestType(%d)", e)
	}
}

type TEndpointSynchronizationType uint

const (
	KEndpointSynchronizationTypeAdaptive TEndpointSynchronizationType = 0
)

func (e TEndpointSynchronizationType) String() string {
	switch e {
	case KEndpointSynchronizationTypeAdaptive:
		return "KEndpointSynchronizationTypeAdaptive"
	default:
		return fmt.Sprintf("TEndpointSynchronizationType(%d)", e)
	}
}

type TEndpointUsageType uint

const (
	KEndpointUsageTypeInterruptNotification TEndpointUsageType = 0
)

func (e TEndpointUsageType) String() string {
	switch e {
	case KEndpointUsageTypeInterruptNotification:
		return "KEndpointUsageTypeInterruptNotification"
	default:
		return fmt.Sprintf("TEndpointUsageType(%d)", e)
	}
}

type TIOUSBBusVoltage uint

const (
	// KIOUSBBusVoltageDefault: The default bus voltage.
	KIOUSBBusVoltageDefault TIOUSBBusVoltage = 0
)

func (e TIOUSBBusVoltage) String() string {
	switch e {
	case KIOUSBBusVoltageDefault:
		return "KIOUSBBusVoltageDefault"
	default:
		return fmt.Sprintf("TIOUSBBusVoltage(%d)", e)
	}
}

type TIOUSBLanguageID uint

const (
	KIOUSBLanguageIDEnglishUS TIOUSBLanguageID = 0
)

func (e TIOUSBLanguageID) String() string {
	switch e {
	case KIOUSBLanguageIDEnglishUS:
		return "KIOUSBLanguageIDEnglishUS"
	default:
		return fmt.Sprintf("TIOUSBLanguageID(%d)", e)
	}
}

type TIsochronousTransferOptions uint

const (
	KIsochronousTransferOptionsNone TIsochronousTransferOptions = 0
)

func (e TIsochronousTransferOptions) String() string {
	switch e {
	case KIsochronousTransferOptionsNone:
		return "KIsochronousTransferOptionsNone"
	default:
		return fmt.Sprintf("TIsochronousTransferOptions(%d)", e)
	}
}

type TUSBCTypeCableType uint

const (
	KUSBTypeCCableTypeNone TUSBCTypeCableType = 0
)

func (e TUSBCTypeCableType) String() string {
	switch e {
	case KUSBTypeCCableTypeNone:
		return "KUSBTypeCCableTypeNone"
	default:
		return fmt.Sprintf("TUSBCTypeCableType(%d)", e)
	}
}

type TUSBHostPortConnectable uint

const (
	KUSBHostPortConnectable TUSBHostPortConnectable = 0
)

func (e TUSBHostPortConnectable) String() string {
	switch e {
	case KUSBHostPortConnectable:
		return "KUSBHostPortConnectable"
	default:
		return fmt.Sprintf("TUSBHostPortConnectable(%d)", e)
	}
}

type TUSBHostPowerSourceType uint

const (
	KUSBHostPowerSourceTypeHardware TUSBHostPowerSourceType = 0
)

func (e TUSBHostPowerSourceType) String() string {
	switch e {
	case KUSBHostPowerSourceTypeHardware:
		return "KUSBHostPowerSourceTypeHardware"
	default:
		return fmt.Sprintf("TUSBHostPowerSourceType(%d)", e)
	}
}

type TUSBLPMExitLatency uint

const (
	KBestEffortServiceLatency TUSBLPMExitLatency = 0
)

func (e TUSBLPMExitLatency) String() string {
	switch e {
	case KBestEffortServiceLatency:
		return "KBestEffortServiceLatency"
	default:
		return fmt.Sprintf("TUSBLPMExitLatency(%d)", e)
	}
}

type Task uint

const (
	TASK_BACKGROUND_APPLICATION Task = 0
	TASK_GRAPHICS_SERVER Task = 0
)

func (e Task) String() string {
	switch e {
	case TASK_BACKGROUND_APPLICATION:
		return "TASK_BACKGROUND_APPLICATION"
	default:
		return fmt.Sprintf("Task(%d)", e)
	}
}

type TaskInspectFlavor uint

const (
	TASK_INSPECT_BASIC_COUNTS TaskInspectFlavor = 0
)

func (e TaskInspectFlavor) String() string {
	switch e {
	case TASK_INSPECT_BASIC_COUNTS:
		return "TASK_INSPECT_BASIC_COUNTS"
	default:
		return fmt.Sprintf("TaskInspectFlavor(%d)", e)
	}
}

type TaskSnapshotFlags uint

const (
	KFrozen TaskSnapshotFlags = 0
	KPidSuspended TaskSnapshotFlags = 0
	KTaskAllowIdleExit TaskSnapshotFlags = 0
	KTaskDarwinBG TaskSnapshotFlags = 0
	KTaskDyldCompactInfoFaultedIn TaskSnapshotFlags = 0
	KTaskDyldCompactInfoMissing TaskSnapshotFlags = 0
	KTaskDyldCompactInfoNone TaskSnapshotFlags = 0
	KTaskDyldCompactInfoTooBig TaskSnapshotFlags = 0
	KTaskDyldCompactInfoTriedFault TaskSnapshotFlags = 0
	KTaskExtDarwinBG TaskSnapshotFlags = 0
	KTaskIsBoosted TaskSnapshotFlags = 0
	KTaskIsDirty TaskSnapshotFlags = 0
	KTaskIsDirtyTracked TaskSnapshotFlags = 0
	KTaskIsForeground TaskSnapshotFlags = 0
	KTaskIsImpDonor TaskSnapshotFlags = 0
	KTaskIsLiveImpDonor TaskSnapshotFlags = 0
	KTaskIsSuppressed TaskSnapshotFlags = 0
	KTaskIsTimerThrottled TaskSnapshotFlags = 0
	KTaskIsTranslated TaskSnapshotFlags = 0
	KTaskRsrcFlagged TaskSnapshotFlags = 0
	KTaskSharedRegionInfoUnavailable TaskSnapshotFlags = 0
	KTaskSharedRegionNone TaskSnapshotFlags = 0
	KTaskSharedRegionOther TaskSnapshotFlags = 0
	KTaskSharedRegionSystem TaskSnapshotFlags = 0
	KTaskTALEngaged TaskSnapshotFlags = 0
	KTaskUUIDInfoFaultedIn TaskSnapshotFlags = 0
	KTaskUUIDInfoMissing TaskSnapshotFlags = 0
	KTaskUUIDInfoTriedFault TaskSnapshotFlags = 0
	KTaskVisNonvisible TaskSnapshotFlags = 0
	KTaskVisVisible TaskSnapshotFlags = 0
	KTaskWqExceededActiveConstrainedThreadLimit TaskSnapshotFlags = 0
	KTaskWqExceededConstrainedThreadLimit TaskSnapshotFlags = 0
	KTaskWqExceededCooperativeThreadLimit TaskSnapshotFlags = 0
	KTaskWqExceededTotalThreadLimit TaskSnapshotFlags = 0
	KTaskWqFlagsAvailable TaskSnapshotFlags = 0
	KTerminatedSnapshot TaskSnapshotFlags = 0
)

func (e TaskSnapshotFlags) String() string {
	switch e {
	case KFrozen:
		return "KFrozen"
	default:
		return fmt.Sprintf("TaskSnapshotFlags(%d)", e)
	}
}

type TaskTransitionType uint

const (
	KTaskIsTerminated TaskTransitionType = 0
)

func (e TaskTransitionType) String() string {
	switch e {
	case KTaskIsTerminated:
		return "KTaskIsTerminated"
	default:
		return fmt.Sprintf("TaskTransitionType(%d)", e)
	}
}

type TcpConnectionClientAccurateEcn uint

const (
	Tcp_connection_client_accurate_ecn_feature_enabled TcpConnectionClientAccurateEcn = 0
	Tcp_connection_client_accurate_ecn_negotiation_success_ect_mangling_detected TcpConnectionClientAccurateEcn = 0
)

func (e TcpConnectionClientAccurateEcn) String() string {
	switch e {
	case Tcp_connection_client_accurate_ecn_feature_enabled:
		return "Tcp_connection_client_accurate_ecn_feature_enabled"
	default:
		return fmt.Sprintf("TcpConnectionClientAccurateEcn(%d)", e)
	}
}

type TcpConnectionServerAccurateEcn uint

const (
	Tcp_connection_server_accurate_ecn_negotiation_success_ect_bleaching_detected TcpConnectionServerAccurateEcn = 0
	Tcp_connection_server_accurate_ecn_requested TcpConnectionServerAccurateEcn = 0
)

func (e TcpConnectionServerAccurateEcn) String() string {
	switch e {
	case Tcp_connection_server_accurate_ecn_negotiation_success_ect_bleaching_detected:
		return "Tcp_connection_server_accurate_ecn_negotiation_success_ect_bleaching_detected"
	default:
		return fmt.Sprintf("TcpConnectionServerAccurateEcn(%d)", e)
	}
}

type TelemetryNoticeT uint

const (
	TELEMETRY_NOTICE_BASE TelemetryNoticeT = 0
)

func (e TelemetryNoticeT) String() string {
	switch e {
	case TELEMETRY_NOTICE_BASE:
		return "TELEMETRY_NOTICE_BASE"
	default:
		return fmt.Sprintf("TelemetryNoticeT(%d)", e)
	}
}

type TelemetryPmi uint

const (
	TELEMETRY_PMI_CYCLES TelemetryPmi = 0
)

func (e TelemetryPmi) String() string {
	switch e {
	case TELEMETRY_PMI_CYCLES:
		return "TELEMETRY_PMI_CYCLES"
	default:
		return fmt.Sprintf("TelemetryPmi(%d)", e)
	}
}

type Telephony uint

const (
	TELEPHONY_DOWN_LINE_PHONE Telephony = 0
	TELEPHONY_TELEPHONE Telephony = 0
)

func (e Telephony) String() string {
	switch e {
	case TELEPHONY_DOWN_LINE_PHONE:
		return "TELEPHONY_DOWN_LINE_PHONE"
	default:
		return fmt.Sprintf("Telephony(%d)", e)
	}
}

type TerminationState uint

const (
	KNeedsTermination TerminationState = 0
	KNotTerminated TerminationState = 0
	KTerminated TerminationState = 0
)

func (e TerminationState) String() string {
	switch e {
	case KNeedsTermination:
		return "KNeedsTermination"
	default:
		return fmt.Sprintf("TerminationState(%d)", e)
	}
}

type ThreadCallOptionsOnce uint

const (
	THREAD_CALL_OPTIONS_ONCE ThreadCallOptionsOnce = 0
)

func (e ThreadCallOptionsOnce) String() string {
	switch e {
	case THREAD_CALL_OPTIONS_ONCE:
		return "THREAD_CALL_OPTIONS_ONCE"
	default:
		return fmt.Sprintf("ThreadCallOptionsOnce(%d)", e)
	}
}

type ThreadCallPriority uint

const (
	// THREAD_CALL_PRIORITY_HIGH: # Discussion
	THREAD_CALL_PRIORITY_HIGH ThreadCallPriority = 0
	// THREAD_CALL_PRIORITY_KERNEL: # Discussion
	THREAD_CALL_PRIORITY_KERNEL ThreadCallPriority = 0
	THREAD_CALL_PRIORITY_KERNEL_HIGH ThreadCallPriority = 0
	// THREAD_CALL_PRIORITY_LOW: # Discussion
	THREAD_CALL_PRIORITY_LOW ThreadCallPriority = 0
	// THREAD_CALL_PRIORITY_USER: # Discussion
	THREAD_CALL_PRIORITY_USER ThreadCallPriority = 0
)

func (e ThreadCallPriority) String() string {
	switch e {
	case THREAD_CALL_PRIORITY_HIGH:
		return "THREAD_CALL_PRIORITY_HIGH"
	default:
		return fmt.Sprintf("ThreadCallPriority(%d)", e)
	}
}

type ThreadSelfcountsKindT uint

const (
	THSC_CPI_PER_PERF_LEVEL ThreadSelfcountsKindT = 0
)

func (e ThreadSelfcountsKindT) String() string {
	switch e {
	case THSC_CPI_PER_PERF_LEVEL:
		return "THSC_CPI_PER_PERF_LEVEL"
	default:
		return fmt.Sprintf("ThreadSelfcountsKindT(%d)", e)
	}
}

type ThreadSnapshotFlags uint

const (
	KGlobalForcedIdle ThreadSnapshotFlags = 0
	KHasDispatchSerial ThreadSnapshotFlags = 0
	KStacksPCOnly ThreadSnapshotFlags = 0
	KThreadDarwinBG ThreadSnapshotFlags = 0
	KThreadFaultedBT ThreadSnapshotFlags = 0
	KThreadIOPassive ThreadSnapshotFlags = 0
	KThreadIdleWorker ThreadSnapshotFlags = 0
	KThreadMain ThreadSnapshotFlags = 0
	KThreadOnCore ThreadSnapshotFlags = 0
	KThreadSuspended ThreadSnapshotFlags = 0
	KThreadTriedFaultBT ThreadSnapshotFlags = 0
	KThreadTruncKernBT ThreadSnapshotFlags = 0
	KThreadTruncUserAsyncBT ThreadSnapshotFlags = 0
	KThreadTruncUserBT ThreadSnapshotFlags = 0
	KThreadTruncatedBT ThreadSnapshotFlags = 0
)

func (e ThreadSnapshotFlags) String() string {
	switch e {
	case KGlobalForcedIdle:
		return "KGlobalForcedIdle"
	default:
		return fmt.Sprintf("ThreadSnapshotFlags(%d)", e)
	}
}

type ThroughputQosTier uint

const (
	THROUGHPUT_QOS_TIER_0 ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_1 ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_2 ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_3 ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_4 ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_5 ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_UNSPECIFIED ThroughputQosTier = 0
)

func (e ThroughputQosTier) String() string {
	switch e {
	case THROUGHPUT_QOS_TIER_0:
		return "THROUGHPUT_QOS_TIER_0"
	default:
		return fmt.Sprintf("ThroughputQosTier(%d)", e)
	}
}

type Timing uint

const (
	TimingApple12 Timing = 0
	TimingApple19 Timing = 0
)

func (e Timing) String() string {
	switch e {
	case TimingApple12:
		return "TimingApple12"
	default:
		return fmt.Sprintf("Timing(%d)", e)
	}
}

type TimingVESA uint

const (
	TimingVESA_1024x768_85hz TimingVESA = 0
	TimingVESA_1280x960_60hz TimingVESA = 0
)

func (e TimingVESA) String() string {
	switch e {
	case TimingVESA_1024x768_85hz:
		return "TimingVESA_1024x768_85hz"
	default:
		return fmt.Sprintf("TimingVESA(%d)", e)
	}
}

type USBClassSpecificDesc uint

const (
	// KUSBClassSpecificDescriptor: The class-specific descriptor value.
	KUSBClassSpecificDescriptor USBClassSpecificDesc = 0
)

func (e USBClassSpecificDesc) String() string {
	switch e {
	case KUSBClassSpecificDescriptor:
		return "KUSBClassSpecificDescriptor"
	default:
		return fmt.Sprintf("USBClassSpecificDesc(%d)", e)
	}
}

type Uio uint

const (
	UIO_READ Uio = 0
	UIO_SYSSPACE Uio = 0
	UIO_SYSSPACE32 Uio = 0
	UIO_WRITE Uio = 0
)

func (e Uio) String() string {
	switch e {
	case UIO_READ:
		return "UIO_READ"
	default:
		return fmt.Sprintf("Uio(%d)", e)
	}
}

type Uiof uint

const (
	UIOF_SYSSPACE Uiof = 0
	UIOF_USERSPACE32 Uiof = 0
)

func (e Uiof) String() string {
	switch e {
	case UIOF_SYSSPACE:
		return "UIOF_SYSSPACE"
	default:
		return fmt.Sprintf("Uiof(%d)", e)
	}
}

type VdspHa uint

const (
	VDSP_HALF_WINDOW VdspHa = 0
	// VDSP_HANN_DENORM: Specifies a denormalized Hann window.
	VDSP_HANN_DENORM VdspHa = 0
	// VDSP_HANN_NORM: Specifies a normalized Hann window
	VDSP_HANN_NORM VdspHa = 0
)

func (e VdspHa) String() string {
	switch e {
	case VDSP_HALF_WINDOW:
		return "VDSP_HALF_WINDOW"
	default:
		return fmt.Sprintf("VdspHa(%d)", e)
	}
}

type VdspIir uint

const (
	VDSP_IIRMonoLeft VdspIir = 0
	VDSP_IIRStereo VdspIir = 0
)

func (e VdspIir) String() string {
	switch e {
	case VDSP_IIRMonoLeft:
		return "VDSP_IIRMonoLeft"
	default:
		return fmt.Sprintf("VdspIir(%d)", e)
	}
}

type Vfs uint

const (
	VFS_RECOVERY_ROLE Vfs = 0
	VFS_VM_ROLE Vfs = 0
)

func (e Vfs) String() string {
	switch e {
	case VFS_RECOVERY_ROLE:
		return "VFS_RECOVERY_ROLE"
	default:
		return fmt.Sprintf("Vfs(%d)", e)
	}
}

type VfsRenameSeclude uint

const (
	VFS_RENAME_EXCL VfsRenameSeclude = 0
)

func (e VfsRenameSeclude) String() string {
	switch e {
	case VFS_RENAME_EXCL:
		return "VFS_RENAME_EXCL"
	default:
		return fmt.Sprintf("VfsRenameSeclude(%d)", e)
	}
}

type ViolationInvalidIOPaddr uint

const (
	SPTM_N_PUBLIC_VIOLATIONS ViolationInvalidIOPaddr = 0
)

func (e ViolationInvalidIOPaddr) String() string {
	switch e {
	case SPTM_N_PUBLIC_VIOLATIONS:
		return "SPTM_N_PUBLIC_VIOLATIONS"
	default:
		return fmt.Sprintf("ViolationInvalidIOPaddr(%d)", e)
	}
}

type VnodeVerifyContext uint

const (
	VNODE_VERIFY_CONTEXT_ALLOC VnodeVerifyContext = 0
	VNODE_VERIFY_CONTEXT_FREE VnodeVerifyContext = 0
)

func (e VnodeVerifyContext) String() string {
	switch e {
	case VNODE_VERIFY_CONTEXT_ALLOC:
		return "VNODE_VERIFY_CONTEXT_ALLOC"
	default:
		return fmt.Sprintf("VnodeVerifyContext(%d)", e)
	}
}

type Vtagtype uint

const (
	VT_HFS Vtagtype = 0
)

func (e Vtagtype) String() string {
	switch e {
	case VT_HFS:
		return "VT_HFS"
	default:
		return fmt.Sprintf("Vtagtype(%d)", e)
	}
}

type Vtype uint

const (
	VBAD Vtype = 0
	VBLK Vtype = 0
	VCHR Vtype = 0
	VCPLX Vtype = 0
	VDIR Vtype = 0
	VLNK Vtype = 0
	VSOCK Vtype = 0
)

func (e Vtype) String() string {
	switch e {
	case VBAD:
		return "VBAD"
	default:
		return fmt.Sprintf("Vtype(%d)", e)
	}
}

type XdrbufType uint

const (
	XDRBUF_BUFFER XdrbufType = 0
)

func (e XdrbufType) String() string {
	switch e {
	case XDRBUF_BUFFER:
		return "XDRBUF_BUFFER"
	default:
		return fmt.Sprintf("XdrbufType(%d)", e)
	}
}

