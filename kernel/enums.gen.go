// Code generated from Apple documentation for kernel. DO NOT EDIT.

package kernel

import (
	"fmt"
)

type AddPacket uint

const (
	AddPacketShift AddPacket = 0
)

func (e AddPacket) String() string {
	switch e {
	case AddPacketShift:
		return "AddPacketShift"
	default:
		return fmt.Sprintf("AddPacket(%d)", e)
	}
}

type AddressRange uint

const (
	AddressRangeErr AddressRange = 0
)

func (e AddressRange) String() string {
	switch e {
	case AddressRangeErr:
		return "AddressRangeErr"
	default:
		return fmt.Sprintf("AddressRange(%d)", e)
	}
}

type Alpha uint

const (
	AlphaStage Alpha = 0
)

func (e Alpha) String() string {
	switch e {
	case AlphaStage:
		return "AlphaStage"
	default:
		return fmt.Sprintf("Alpha(%d)", e)
	}
}

type AtaDeviceType uint

const (
	KATADeviceType        AtaDeviceType = 0
	KATAPIDeviceType      AtaDeviceType = 0
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
	KInternalSATA      AtaSocketType = 0
	KInternalSATA2     AtaSocketType = 0
	KMediaBaySocket    AtaSocketType = 0
	KPCCardSocket      AtaSocketType = 0
	KSATA2Bay          AtaSocketType = 0
	KSATABay           AtaSocketType = 0
	KUnknownSocket     AtaSocketType = 0
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
	BATAAddressNotFound BATABadBlock = 0
	BATABadBlockValue   BATABadBlock = 0
	BATACommandAborted  BATABadBlock = 0
	BATAIDNotFound      BATABadBlock = 0
	BATAMediaChangeReq  BATABadBlock = 0
	BATAMediaChanged    BATABadBlock = 0
	BATATrack0NotFound  BATABadBlock = 0
	BATAUncorrectable   BATABadBlock = 0
	MATAAddressNotFound BATABadBlock = 0
	MATABadBlock        BATABadBlock = 0
	MATACommandAborted  BATABadBlock = 0
	MATAIDNotFound      BATABadBlock = 0
	MATAMediaChangeReq  BATABadBlock = 0
	MATAMediaChanged    BATABadBlock = 0
	MATATrack0NotFound  BATABadBlock = 0
	MATAUncorrectable   BATABadBlock = 0
)

func (e BATABadBlock) String() string {
	switch e {
	case BATAAddressNotFound:
		return "BATAAddressNotFound"
	default:
		return fmt.Sprintf("BATABadBlock(%d)", e)
	}
}

type BATADCROne uint

const (
	BATADCROneValue   BATADCROne = 0
	BATADCRReset      BATADCROne = 0
	BATADCRnIntEnable BATADCROne = 0
	MATADCROne        BATADCROne = 0
	MATADCRReset      BATADCROne = 0
	MATADCRnIntEnable BATADCROne = 0
)

func (e BATADCROne) String() string {
	switch e {
	case BATADCROneValue:
		return "BATADCROneValue"
	default:
		return fmt.Sprintf("BATADCROne(%d)", e)
	}
}

type BATAPIuseDM uint

const (
	BATAPIuseDMA BATAPIuseDM = 0
)

func (e BATAPIuseDM) String() string {
	switch e {
	case BATAPIuseDMA:
		return "BATAPIuseDMA"
	default:
		return fmt.Sprintf("BATAPIuseDM(%d)", e)
	}
}

type BidirectionalEchoCanceling uint

const (
	BIDIRECTIONAL_ECHO_CANCELING_SPEAKERPHONE BidirectionalEchoCanceling = 0
)

func (e BidirectionalEchoCanceling) String() string {
	switch e {
	case BIDIRECTIONAL_ECHO_CANCELING_SPEAKERPHONE:
		return "BIDIRECTIONAL_ECHO_CANCELING_SPEAKERPHONE"
	default:
		return fmt.Sprintf("BidirectionalEchoCanceling(%d)", e)
	}
}

type BpfMode uint

const (
	// BPF_MODE_OUTPUT: # Discussion
	BPF_MODE_OUTPUT BpfMode = 0
)

func (e BpfMode) String() string {
	switch e {
	case BPF_MODE_OUTPUT:
		return "BPF_MODE_OUTPUT"
	default:
		return fmt.Sprintf("BpfMode(%d)", e)
	}
}

type BpfTap uint

const (
	BPF_TAP_DISABLE BpfTap = 0
)

func (e BpfTap) String() string {
	switch e {
	case BPF_TAP_DISABLE:
		return "BPF_TAP_DISABLE"
	default:
		return fmt.Sprintf("BpfTap(%d)", e)
	}
}

type Btf uint

const (
	BTF_KERN_INTERRUPTED Btf = 0
	BTF_NONE             Btf = 0
)

func (e Btf) String() string {
	switch e {
	case BTF_KERN_INTERRUPTED:
		return "BTF_KERN_INTERRUPTED"
	default:
		return fmt.Sprintf("Btf(%d)", e)
	}
}

type Bti64 uint

const (
	BTI_64_BIT Bti64 = 0
)

func (e Bti64) String() string {
	switch e {
	case BTI_64_BIT:
		return "BTI_64_BIT"
	default:
		return fmt.Sprintf("Bti64(%d)", e)
	}
}

type BtpKernOffset uint

const (
	BTP_KERN_OFFSET_32 BtpKernOffset = 0
)

func (e BtpKernOffset) String() string {
	switch e {
	case BTP_KERN_OFFSET_32:
		return "BTP_KERN_OFFSET_32"
	default:
		return fmt.Sprintf("BtpKernOffset(%d)", e)
	}
}

type CacheTypeT uint

const (
	L1D        CacheTypeT = 0
	L1I        CacheTypeT = 0
	L2U        CacheTypeT = 0
	L3U        CacheTypeT = 0
	LCACHE_MAX CacheTypeT = 0
	Lnone      CacheTypeT = 0
)

func (e CacheTypeT) String() string {
	switch e {
	case L1D:
		return "L1D"
	default:
		return fmt.Sprintf("CacheTypeT(%d)", e)
	}
}

type ClusterTypeT uint

const (
	CLUSTER_TYPE_E ClusterTypeT = 0
	MAX_CPU_TYPES  ClusterTypeT = 0
)

func (e ClusterTypeT) String() string {
	switch e {
	case CLUSTER_TYPE_E:
		return "CLUSTER_TYPE_E"
	default:
		return fmt.Sprintf("ClusterTypeT(%d)", e)
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

type CpustatePanic uint

const (
	CPUSTATE_PANIC_SPIN CpustatePanic = 0
)

func (e CpustatePanic) String() string {
	switch e {
	case CPUSTATE_PANIC_SPIN:
		return "CPUSTATE_PANIC_SPIN"
	default:
		return fmt.Sprintf("CpustatePanic(%d)", e)
	}
}

type Crypte uint

const (
	CRYPTEX1_AUTH_ENV_GENERIC              Crypte = 0
	CRYPTEX1_AUTH_ENV_GENERIC_SUPPLEMENTAL Crypte = 0
	CRYPTEX_AUTH_MAX                       Crypte = 0
	CRYPTEX_AUTH_MOBILE_ASSET              Crypte = 0
	CRYPTEX_AUTH_PDI_NONCE                 Crypte = 0
)

func (e Crypte) String() string {
	switch e {
	case CRYPTEX1_AUTH_ENV_GENERIC:
		return "CRYPTEX1_AUTH_ENV_GENERIC"
	default:
		return fmt.Sprintf("Crypte(%d)", e)
	}
}

type CsCdhash uint

const (
	CS_CDHASH_LEN CsCdhash = 0
)

func (e CsCdhash) String() string {
	switch e {
	case CS_CDHASH_LEN:
		return "CS_CDHASH_LEN"
	default:
		return fmt.Sprintf("CsCdhash(%d)", e)
	}
}

type CsLaunchType uint

const (
	CS_LAUNCH_TYPE_APPLICATION    CsLaunchType = 0
	CS_LAUNCH_TYPE_NONE           CsLaunchType = 0
	CS_LAUNCH_TYPE_SYSDIAGNOSE    CsLaunchType = 0
	CS_LAUNCH_TYPE_SYSTEM_SERVICE CsLaunchType = 0
)

func (e CsLaunchType) String() string {
	switch e {
	case CS_LAUNCH_TYPE_APPLICATION:
		return "CS_LAUNCH_TYPE_APPLICATION"
	default:
		return fmt.Sprintf("CsLaunchType(%d)", e)
	}
}

type CsLinkageApplication uint

const (
	CS_LINKAGE_APPLICATION_INVALID           CsLinkageApplication = 0
	CS_LINKAGE_APPLICATION_OOPJIT_INVALID    CsLinkageApplication = 0
	CS_LINKAGE_APPLICATION_OOPJIT_MLCOMPILER CsLinkageApplication = 0
	CS_LINKAGE_APPLICATION_OOPJIT_PREVIEWS   CsLinkageApplication = 0
	CS_LINKAGE_APPLICATION_OOPJIT_TOTAL      CsLinkageApplication = 0
	CS_LINKAGE_APPLICATION_ROSETTA_AOT       CsLinkageApplication = 0
	CS_LINKAGE_APPLICATION_XOJIT_PREVIEWS    CsLinkageApplication = 0
)

func (e CsLinkageApplication) String() string {
	switch e {
	case CS_LINKAGE_APPLICATION_INVALID:
		return "CS_LINKAGE_APPLICATION_INVALID"
	default:
		return fmt.Sprintf("CsLinkageApplication(%d)", e)
	}
}

type Csc uint

const (
	CscDirectSetEntries           Csc = 0
	CscDoCommunication            Csc = 0
	CscDrawHardwareCursor         Csc = 0
	CscGetBaseAddr                Csc = 0
	CscGetClutBehavior            Csc = 0
	CscGetCommunicationInfo       Csc = 0
	CscGetConnection              Csc = 0
	CscGetConvolution             Csc = 0
	CscGetCurMode                 Csc = 0
	CscGetDDCBlock                Csc = 0
	CscGetDefaultMode             Csc = 0
	CscGetDetailedTiming          Csc = 0
	CscGetEntries                 Csc = 0
	CscGetFeatureConfiguration    Csc = 0
	CscGetFeatureList             Csc = 0
	CscGetGamma                   Csc = 0
	CscGetGammaInfoList           Csc = 0
	CscGetGray                    Csc = 0
	CscGetHardwareCursorDrawState Csc = 0
	CscGetInterrupt               Csc = 0
	CscGetMirror                  Csc = 0
	CscGetMode                    Csc = 0
	CscGetModeBaseAddress         Csc = 0
	CscGetModeTiming              Csc = 0
	CscGetMultiConnect            Csc = 0
	CscGetNextResolution          Csc = 0
	CscGetPageBase                Csc = 0
	CscGetPageCnt                 Csc = 0
	CscGetPages                   Csc = 0
	CscGetPowerState              Csc = 0
	CscGetPreferredConfiguration  Csc = 0
	CscGetScaler                  Csc = 0
	CscGetScalerInfo              Csc = 0
	CscGetScanProc                Csc = 0
	CscGetSync                    Csc = 0
	CscGetTimingRanges            Csc = 0
	CscGetVideoParameters         Csc = 0
	CscGrayPage                   Csc = 0
	CscGrayScreen                 Csc = 0
	CscKillIO                     Csc = 0
	CscPrivateControlCall         Csc = 0
	CscPrivateStatusCall          Csc = 0
	CscProbeConnection            Csc = 0
	CscReset                      Csc = 0
	CscRetrieveGammaTable         Csc = 0
	CscSavePreferredConfiguration Csc = 0
	CscSetClutBehavior            Csc = 0
	CscSetConvolution             Csc = 0
	CscSetDefaultMode             Csc = 0
	CscSetDetailedTiming          Csc = 0
	CscSetEntries                 Csc = 0
	CscSetFeatureConfiguration    Csc = 0
	CscSetGamma                   Csc = 0
	CscSetGray                    Csc = 0
	CscSetHardwareCursor          Csc = 0
	CscSetInterrupt               Csc = 0
	CscSetMirror                  Csc = 0
	CscSetMode                    Csc = 0
	CscSetMultiConnect            Csc = 0
	CscSetPowerState              Csc = 0
	CscSetScaler                  Csc = 0
	CscSetSync                    Csc = 0
	CscSupportsHardwareCursor     Csc = 0
	CscSwitchMode                 Csc = 0
	CscUnusedCall                 Csc = 0
)

func (e Csc) String() string {
	switch e {
	case CscDirectSetEntries:
		return "CscDirectSetEntries"
	default:
		return fmt.Sprintf("Csc(%d)", e)
	}
}

type Dot3ChipSet uint

const (
	Dot3ChipSetFujitsu86950 Dot3ChipSet = 0
	Dot3ChipSetIntel82557   Dot3ChipSet = 0
	Dot3ChipSetIntel82586   Dot3ChipSet = 0
	Dot3ChipSetIntel82596   Dot3ChipSet = 0
	Dot3ChipSetNational8390 Dot3ChipSet = 0
)

func (e Dot3ChipSet) String() string {
	switch e {
	case Dot3ChipSetFujitsu86950:
		return "Dot3ChipSetFujitsu86950"
	default:
		return fmt.Sprintf("Dot3ChipSet(%d)", e)
	}
}

type Dot3ChipSetAM uint

const (
	Dot3ChipSetAMD7990 Dot3ChipSetAM = 0
)

func (e Dot3ChipSetAM) String() string {
	switch e {
	case Dot3ChipSetAMD7990:
		return "Dot3ChipSetAMD7990"
	default:
		return fmt.Sprintf("Dot3ChipSetAM(%d)", e)
	}
}

type Dot3ChipSetDigitalD uint

const (
	Dot3ChipSetDigitalDC21040  Dot3ChipSetDigitalD = 0
	Dot3ChipSetDigitalDC21041  Dot3ChipSetDigitalD = 0
	Dot3ChipSetDigitalDC21140  Dot3ChipSetDigitalD = 0
	Dot3ChipSetDigitalDC21140A Dot3ChipSetDigitalD = 0
	Dot3ChipSetDigitalDC21142  Dot3ChipSetDigitalD = 0
)

func (e Dot3ChipSetDigitalD) String() string {
	switch e {
	case Dot3ChipSetDigitalDC21040:
		return "Dot3ChipSetDigitalDC21040"
	default:
		return fmt.Sprintf("Dot3ChipSetDigitalD(%d)", e)
	}
}

type Dot3ChipSetWesternDigital83 uint

const (
	Dot3ChipSetWesternDigital83C690 Dot3ChipSetWesternDigital83 = 0
)

func (e Dot3ChipSetWesternDigital83) String() string {
	switch e {
	case Dot3ChipSetWesternDigital83C690:
		return "Dot3ChipSetWesternDigital83C690"
	default:
		return fmt.Sprintf("Dot3ChipSetWesternDigital83(%d)", e)
	}
}

type Dot3VendorAM uint

const (
	Dot3VendorAMD Dot3VendorAM = 0
)

func (e Dot3VendorAM) String() string {
	switch e {
	case Dot3VendorAMD:
		return "Dot3VendorAMD"
	default:
		return fmt.Sprintf("Dot3VendorAM(%d)", e)
	}
}

type Duration uint

const (
	DurationMinute Duration = 0
)

func (e Duration) String() string {
	switch e {
	case DurationMinute:
		return "DurationMinute"
	default:
		return fmt.Sprintf("Duration(%d)", e)
	}
}

type DyldChained uint

const (
	DYLD_CHAINED_IMPORT          DyldChained = 0
	DYLD_CHAINED_IMPORT_ADDEND   DyldChained = 0
	DYLD_CHAINED_IMPORT_ADDEND64 DyldChained = 0
)

func (e DyldChained) String() string {
	switch e {
	case DYLD_CHAINED_IMPORT:
		return "DYLD_CHAINED_IMPORT"
	default:
		return fmt.Sprintf("DyldChained(%d)", e)
	}
}

type DyldChainedPtr uint

const (
	DYLD_CHAINED_PTR_32 DyldChainedPtr = 0
)

func (e DyldChainedPtr) String() string {
	switch e {
	case DYLD_CHAINED_PTR_32:
		return "DYLD_CHAINED_PTR_32"
	default:
		return fmt.Sprintf("DyldChainedPtr(%d)", e)
	}
}

type DyldChainedPtrStart uint

const (
	DYLD_CHAINED_PTR_START_MULTI DyldChainedPtrStart = 0
)

func (e DyldChainedPtrStart) String() string {
	switch e {
	case DYLD_CHAINED_PTR_START_MULTI:
		return "DYLD_CHAINED_PTR_START_MULTI"
	default:
		return fmt.Sprintf("DyldChainedPtrStart(%d)", e)
	}
}

type ENoteExitReparented uint

const (
	// Deprecated.
	ENoteExitReparentedDeprecated ENoteExitReparented = 0
)

func (e ENoteExitReparented) String() string {
	switch e {
	case ENoteExitReparentedDeprecated:
		return "ENoteExitReparentedDeprecated"
	default:
		return fmt.Sprintf("ENoteExitReparented(%d)", e)
	}
}

type ENoteReap uint

const (
	// Deprecated.
	ENoteReapDeprecated ENoteReap = 0
)

func (e ENoteReap) String() string {
	switch e {
	case ENoteReapDeprecated:
		return "ENoteReapDeprecated"
	default:
		return fmt.Sprintf("ENoteReap(%d)", e)
	}
}

const Eax uint = 0

type Ecc uint

const (
	ECC_DB_CORRUPTED   Ecc = 0
	ECC_DB_ONLY        Ecc = 0
	ECC_IS_CORRECTABLE Ecc = 0
	ECC_IS_TEST_ERROR  Ecc = 0
	ECC_NONE           Ecc = 0
	ECC_REMOVE_ADDR    Ecc = 0
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
	EfiACPIMemoryNVS           Efi = 0
	EfiACPIReclaimMemory       Efi = 0
	EfiBootServicesCode        Efi = 0
	EfiBootServicesData        Efi = 0
	EfiConventionalMemory      Efi = 0
	EfiLoaderCode              Efi = 0
	EfiLoaderData              Efi = 0
	EfiMaxMemoryType           Efi = 0
	EfiMemoryMappedIO          Efi = 0
	EfiMemoryMappedIOPortSpace Efi = 0
	EfiPalCode                 Efi = 0
	EfiReservedMemoryType      Efi = 0
	EfiRuntimeServicesCode     Efi = 0
	EfiRuntimeServicesData     Efi = 0
	EfiUnusableMemory          Efi = 0
)

func (e Efi) String() string {
	switch e {
	case EfiACPIMemoryNVS:
		return "EfiACPIMemoryNVS"
	default:
		return fmt.Sprintf("Efi(%d)", e)
	}
}

type EfiReset uint

const (
	EfiResetCold EfiReset = 0
)

func (e EfiReset) String() string {
	switch e {
	case EfiResetCold:
		return "EfiResetCold"
	default:
		return fmt.Sprintf("EfiReset(%d)", e)
	}
}

type Embedded uint

const (
	EMBEDDED_PHONOGRAPH Embedded = 0
)

func (e Embedded) String() string {
	switch e {
	case EMBEDDED_PHONOGRAPH:
		return "EMBEDDED_PHONOGRAPH"
	default:
		return fmt.Sprintf("Embedded(%d)", e)
	}
}

type EmbeddedPanicHeaderFlagButtonReset uint

const (
	EMBEDDED_PANIC_HEADER_FLAG_BUTTON_RESET_PANIC EmbeddedPanicHeaderFlagButtonReset = 0
)

func (e EmbeddedPanicHeaderFlagButtonReset) String() string {
	switch e {
	case EMBEDDED_PANIC_HEADER_FLAG_BUTTON_RESET_PANIC:
		return "EMBEDDED_PANIC_HEADER_FLAG_BUTTON_RESET_PANIC"
	default:
		return fmt.Sprintf("EmbeddedPanicHeaderFlagButtonReset(%d)", e)
	}
}

type ExcbAction uint

const (
	EXCB_ACTION_NONE ExcbAction = 0
)

func (e ExcbAction) String() string {
	switch e {
	case EXCB_ACTION_NONE:
		return "EXCB_ACTION_NONE"
	default:
		return fmt.Sprintf("ExcbAction(%d)", e)
	}
}

type ExcbClassIllegalInstr uint

const (
	EXCB_CLASS_ILLEGAL_INSTR_SET ExcbClassIllegalInstr = 0
)

func (e ExcbClassIllegalInstr) String() string {
	switch e {
	case EXCB_CLASS_ILLEGAL_INSTR_SET:
		return "EXCB_CLASS_ILLEGAL_INSTR_SET"
	default:
		return fmt.Sprintf("ExcbClassIllegalInstr(%d)", e)
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

type ExtPaniclogFlags uint

const (
	EXT_PANICLOG_FLAGS_ADD_SEPARATE_KEY ExtPaniclogFlags = 0
	EXT_PANICLOG_FLAGS_NONE             ExtPaniclogFlags = 0
)

func (e ExtPaniclogFlags) String() string {
	switch e {
	case EXT_PANICLOG_FLAGS_ADD_SEPARATE_KEY:
		return "EXT_PANICLOG_FLAGS_ADD_SEPARATE_KEY"
	default:
		return fmt.Sprintf("ExtPaniclogFlags(%d)", e)
	}
}

type ExtPaniclogOptions uint

const (
	EXT_PANICLOG_OPTIONS_ADD_SEPARATE_KEY ExtPaniclogOptions = 0
	EXT_PANICLOG_OPTIONS_NONE             ExtPaniclogOptions = 0
	EXT_PANICLOG_OPTIONS_WITH_BUFFER      ExtPaniclogOptions = 0
)

func (e ExtPaniclogOptions) String() string {
	switch e {
	case EXT_PANICLOG_OPTIONS_ADD_SEPARATE_KEY:
		return "EXT_PANICLOG_OPTIONS_ADD_SEPARATE_KEY"
	default:
		return fmt.Sprintf("ExtPaniclogOptions(%d)", e)
	}
}

type External1394Da uint

const (
	EXTERNAL_1394_DA_STREAM External1394Da = 0
)

func (e External1394Da) String() string {
	switch e {
	case EXTERNAL_1394_DA_STREAM:
		return "EXTERNAL_1394_DA_STREAM"
	default:
		return fmt.Sprintf("External1394Da(%d)", e)
	}
}

type Fft uint

const (
	FFT_FORWARD Fft = 0
	FFT_INVERSE Fft = 0
	FFT_RADIX2  Fft = 0
)

func (e Fft) String() string {
	switch e {
	case FFT_FORWARD:
		return "FFT_FORWARD"
	default:
		return fmt.Sprintf("Fft(%d)", e)
	}
}

type GenericSnapshotFlags uint

const (
	KKernel64_p GenericSnapshotFlags = 0
	KUser64_p   GenericSnapshotFlags = 0
)

func (e GenericSnapshotFlags) String() string {
	switch e {
	case KKernel64_p:
		return "KKernel64_p"
	default:
		return fmt.Sprintf("GenericSnapshotFlags(%d)", e)
	}
}

type GraftdmgCryptex uint

const (
	GRAFTDMG_CRYPTEX_BOOT GraftdmgCryptex = 0
)

func (e GraftdmgCryptex) String() string {
	switch e {
	case GRAFTDMG_CRYPTEX_BOOT:
		return "GRAFTDMG_CRYPTEX_BOOT"
	default:
		return fmt.Sprintf("GraftdmgCryptex(%d)", e)
	}
}

type GssdKrb5 uint

const (
	GSSD_KRB5_REFERRAL GssdKrb5 = 0
)

func (e GssdKrb5) String() string {
	switch e {
	case GSSD_KRB5_REFERRAL:
		return "GSSD_KRB5_REFERRAL"
	default:
		return fmt.Sprintf("GssdKrb5(%d)", e)
	}
}

type GssdNo uint

const (
	GSSD_NO_MECH GssdNo = 0
)

func (e GssdNo) String() string {
	switch e {
	case GSSD_NO_MECH:
		return "GSSD_NO_MECH"
	default:
		return fmt.Sprintf("GssdNo(%d)", e)
	}
}

type HvDebug uint

const (
	HV_DEBUG_STATE HvDebug = 0
)

func (e HvDebug) String() string {
	switch e {
	case HV_DEBUG_STATE:
		return "HV_DEBUG_STATE"
	default:
		return fmt.Sprintf("HvDebug(%d)", e)
	}
}

type HvT uint

const (
	HV_TASK_TRAP   HvT = 0
	HV_THREAD_TRAP HvT = 0
)

func (e HvT) String() string {
	switch e {
	case HV_TASK_TRAP:
		return "HV_TASK_TRAP"
	default:
		return fmt.Sprintf("HvT(%d)", e)
	}
}

type HvgHcall uint

const (
	HVG_HCALL_ACCESS_DENIED       HvgHcall = 0
	HVG_HCALL_COUNT               HvgHcall = 0
	HVG_HCALL_FEAT_DISABLED       HvgHcall = 0
	HVG_HCALL_GET_BOOTSESSIONUUID HvgHcall = 0
	HVG_HCALL_GET_MABS_OFFSET     HvgHcall = 0
	HVG_HCALL_INVALID_CODE        HvgHcall = 0
	HVG_HCALL_INVALID_PARAMETER   HvgHcall = 0
	HVG_HCALL_IO_FAILED           HvgHcall = 0
	HVG_HCALL_SET_COREDUMP_DATA   HvgHcall = 0
	HVG_HCALL_SUCCESS             HvgHcall = 0
	HVG_HCALL_TRIGGER_DUMP        HvgHcall = 0
	HVG_HCALL_UNSUPPORTED         HvgHcall = 0
	HVG_HCALL_VCPU_KICK           HvgHcall = 0
	HVG_HCALL_VCPU_WFK            HvgHcall = 0
)

func (e HvgHcall) String() string {
	switch e {
	case HVG_HCALL_ACCESS_DENIED:
		return "HVG_HCALL_ACCESS_DENIED"
	default:
		return fmt.Sprintf("HvgHcall(%d)", e)
	}
}

type HvgHcallDumpOption uint

const (
	HVG_HCALL_DUMP_OPTION_REGULAR HvgHcallDumpOption = 0
)

func (e HvgHcallDumpOption) String() string {
	switch e {
	case HVG_HCALL_DUMP_OPTION_REGULAR:
		return "HVG_HCALL_DUMP_OPTION_REGULAR"
	default:
		return fmt.Sprintf("HvgHcallDumpOption(%d)", e)
	}
}

type IOCSRKeyType uint

const (
	KCSRDirectoryKeyType    IOCSRKeyType = 0
	KCSRImmediateKeyType    IOCSRKeyType = 0
	KCSROffsetKeyType       IOCSRKeyType = 0
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
	KConfigDirectoryKeyType    IOConfigKeyType = 0
	KConfigImmediateKeyType    IOConfigKeyType = 0
	KConfigLeafKeyType         IOConfigKeyType = 0
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

type IOFWAVCPlug uint

const (
	IOFWAVCPlugAsynchInputType    IOFWAVCPlug = 0
	IOFWAVCPlugAsynchOutputType   IOFWAVCPlug = 0
	IOFWAVCPlugExternalInputType  IOFWAVCPlug = 0
	IOFWAVCPlugExternalOutputType IOFWAVCPlug = 0
	IOFWAVCPlugIsochInputType     IOFWAVCPlug = 0
	IOFWAVCPlugIsochOutputType    IOFWAVCPlug = 0
	IOFWAVCPlugSubunitDestType    IOFWAVCPlug = 0
	IOFWAVCPlugSubunitSourceType  IOFWAVCPlug = 0
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

type IOPMHighest uint

const (
	IOPMHighestState IOPMHighest = 0
)

func (e IOPMHighest) String() string {
	switch e {
	case IOPMHighestState:
		return "IOPMHighestState"
	default:
		return fmt.Sprintf("IOPMHighest(%d)", e)
	}
}

type IOPMMaxPower uint

const (
	IOPMMaxPowerStates IOPMMaxPower = 0
)

func (e IOPMMaxPower) String() string {
	switch e {
	case IOPMMaxPowerStates:
		return "IOPMMaxPowerStates"
	default:
		return fmt.Sprintf("IOPMMaxPower(%d)", e)
	}
}

type IOUSBHostCICapabilitiesMessage uint

const (
	IOUSBHostCICapabilitiesMessageControlPortCount                  IOUSBHostCICapabilitiesMessage = 0
	IOUSBHostCICapabilitiesMessageControlPortCountPhase             IOUSBHostCICapabilitiesMessage = 0
	IOUSBHostCICapabilitiesMessageData0CommandTimeoutThreshold      IOUSBHostCICapabilitiesMessage = 0
	IOUSBHostCICapabilitiesMessageData0CommandTimeoutThresholdPhase IOUSBHostCICapabilitiesMessage = 0
	IOUSBHostCICapabilitiesMessageData0ConnectionLatency            IOUSBHostCICapabilitiesMessage = 0
	IOUSBHostCICapabilitiesMessageData0ConnectionLatencyPhase       IOUSBHostCICapabilitiesMessage = 0
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
	IOUSBHostCICommandMessageControlStatus      IOUSBHostCICommandMessageControl = 0
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
	IOUSBHostCIControllerStateOff    IOUSBHostCIControllerState = 0
)

func (e IOUSBHostCIControllerState) String() string {
	switch e {
	case IOUSBHostCIControllerStateActive:
		return "IOUSBHostCIControllerStateActive"
	default:
		return fmt.Sprintf("IOUSBHostCIControllerState(%d)", e)
	}
}

type IOUSBHostCIDeviceCreateCommandData0Root uint

const (
	IOUSBHostCIDeviceCreateCommandData0RootPort IOUSBHostCIDeviceCreateCommandData0Root = 0
)

func (e IOUSBHostCIDeviceCreateCommandData0Root) String() string {
	switch e {
	case IOUSBHostCIDeviceCreateCommandData0RootPort:
		return "IOUSBHostCIDeviceCreateCommandData0RootPort"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceCreateCommandData0Root(%d)", e)
	}
}

type IOUSBHostCIDeviceCreateCommandData1Device uint

const (
	IOUSBHostCIDeviceCreateCommandData1DeviceAddress IOUSBHostCIDeviceCreateCommandData1Device = 0
)

func (e IOUSBHostCIDeviceCreateCommandData1Device) String() string {
	switch e {
	case IOUSBHostCIDeviceCreateCommandData1DeviceAddress:
		return "IOUSBHostCIDeviceCreateCommandData1DeviceAddress"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceCreateCommandData1Device(%d)", e)
	}
}

type IOUSBHostCIDeviceSpeed uint

const (
	IOUSBHostCIDeviceSpeedFull      IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedHigh      IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedLow       IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedNone      IOUSBHostCIDeviceSpeed = 0
	IOUSBHostCIDeviceSpeedSuper     IOUSBHostCIDeviceSpeed = 0
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
	IOUSBHostCIDeviceStateActive    IOUSBHostCIDeviceState = 0
	IOUSBHostCIDeviceStateDestroyed IOUSBHostCIDeviceState = 0
	IOUSBHostCIDeviceStatePaused    IOUSBHostCIDeviceState = 0
)

func (e IOUSBHostCIDeviceState) String() string {
	switch e {
	case IOUSBHostCIDeviceStateActive:
		return "IOUSBHostCIDeviceStateActive"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceState(%d)", e)
	}
}

type IOUSBHostCIDeviceUpdateCommandData1Descriptor uint

const (
	IOUSBHostCIDeviceUpdateCommandData1DescriptorAddress IOUSBHostCIDeviceUpdateCommandData1Descriptor = 0
)

func (e IOUSBHostCIDeviceUpdateCommandData1Descriptor) String() string {
	switch e {
	case IOUSBHostCIDeviceUpdateCommandData1DescriptorAddress:
		return "IOUSBHostCIDeviceUpdateCommandData1DescriptorAddress"
	default:
		return fmt.Sprintf("IOUSBHostCIDeviceUpdateCommandData1Descriptor(%d)", e)
	}
}

type IOUSBHostCIDoorbellDevice uint

const (
	IOUSBHostCIDoorbellDeviceAddress      IOUSBHostCIDoorbellDevice = 0
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

type IOUSBHostCIEndpointCreateCommandData1 uint

const (
	IOUSBHostCIEndpointCreateCommandData1Descriptor IOUSBHostCIEndpointCreateCommandData1 = 0
)

func (e IOUSBHostCIEndpointCreateCommandData1) String() string {
	switch e {
	case IOUSBHostCIEndpointCreateCommandData1Descriptor:
		return "IOUSBHostCIEndpointCreateCommandData1Descriptor"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointCreateCommandData1(%d)", e)
	}
}

type IOUSBHostCIEndpointResetCommandData1Clear uint

const (
	IOUSBHostCIEndpointResetCommandData1ClearState IOUSBHostCIEndpointResetCommandData1Clear = 0
)

func (e IOUSBHostCIEndpointResetCommandData1Clear) String() string {
	switch e {
	case IOUSBHostCIEndpointResetCommandData1ClearState:
		return "IOUSBHostCIEndpointResetCommandData1ClearState"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointResetCommandData1Clear(%d)", e)
	}
}

type IOUSBHostCIEndpointSetNextTransferCommandData1 uint

const (
	IOUSBHostCIEndpointSetNextTransferCommandData1Address IOUSBHostCIEndpointSetNextTransferCommandData1 = 0
)

func (e IOUSBHostCIEndpointSetNextTransferCommandData1) String() string {
	switch e {
	case IOUSBHostCIEndpointSetNextTransferCommandData1Address:
		return "IOUSBHostCIEndpointSetNextTransferCommandData1Address"
	default:
		return fmt.Sprintf("IOUSBHostCIEndpointSetNextTransferCommandData1(%d)", e)
	}
}

type IOUSBHostCIEndpointState uint

const (
	IOUSBHostCIEndpointStateActive    IOUSBHostCIEndpointState = 0
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
	IOUSBHostCIEndpointUpdateCommandData1Descriptor      IOUSBHostCIEndpointUpdateCommandData1 = 0
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
	IOUSBHostCIExceptionTypeCommandFailure   IOUSBHostCIExceptionType = 0
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
	IOUSBHostCIIsochronousTransferControlASAP             IOUSBHostCIIsochronousTransferControl = 0
	IOUSBHostCIIsochronousTransferControlFrameNumber      IOUSBHostCIIsochronousTransferControl = 0
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

type IOUSBHostCIIsochronousTransferData0 uint

const (
	IOUSBHostCIIsochronousTransferData0Length      IOUSBHostCIIsochronousTransferData0 = 0
	IOUSBHostCIIsochronousTransferData0LengthPhase IOUSBHostCIIsochronousTransferData0 = 0
)

func (e IOUSBHostCIIsochronousTransferData0) String() string {
	switch e {
	case IOUSBHostCIIsochronousTransferData0Length:
		return "IOUSBHostCIIsochronousTransferData0Length"
	default:
		return fmt.Sprintf("IOUSBHostCIIsochronousTransferData0(%d)", e)
	}
}

type IOUSBHostCIIsochronousTransferData1 uint

const (
	IOUSBHostCIIsochronousTransferData1Buffer IOUSBHostCIIsochronousTransferData1 = 0
)

func (e IOUSBHostCIIsochronousTransferData1) String() string {
	switch e {
	case IOUSBHostCIIsochronousTransferData1Buffer:
		return "IOUSBHostCIIsochronousTransferData1Buffer"
	default:
		return fmt.Sprintf("IOUSBHostCIIsochronousTransferData1(%d)", e)
	}
}

type IOUSBHostCILinkData1TransferStructure uint

const (
	IOUSBHostCILinkData1TransferStructureAddress IOUSBHostCILinkData1TransferStructure = 0
)

func (e IOUSBHostCILinkData1TransferStructure) String() string {
	switch e {
	case IOUSBHostCILinkData1TransferStructureAddress:
		return "IOUSBHostCILinkData1TransferStructureAddress"
	default:
		return fmt.Sprintf("IOUSBHostCILinkData1TransferStructure(%d)", e)
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

type IOUSBHostCIMessageControlNo uint

const (
	IOUSBHostCIMessageControlNoResponse IOUSBHostCIMessageControlNo = 0
)

func (e IOUSBHostCIMessageControlNo) String() string {
	switch e {
	case IOUSBHostCIMessageControlNoResponse:
		return "IOUSBHostCIMessageControlNoResponse"
	default:
		return fmt.Sprintf("IOUSBHostCIMessageControlNo(%d)", e)
	}
}

type IOUSBHostCIMessageStatus int

const (
	IOUSBHostCIMessageStatusBadArgument        IOUSBHostCIMessageStatus = 0
	IOUSBHostCIMessageStatusEndpointStopped    IOUSBHostCIMessageStatus = 0
	IOUSBHostCIMessageStatusError              IOUSBHostCIMessageStatus = 0
	IOUSBHostCIMessageStatusMissedServiceError IOUSBHostCIMessageStatus = 0
	IOUSBHostCIMessageStatusOffline            IOUSBHostCIMessageStatus = 0
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
	IOUSBHostCIMessageTypeDeviceUpdate          IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypeEndpointReset         IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypeNormalTransfer        IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypePortPowerOn           IOUSBHostCIMessageType = 0
	IOUSBHostCIMessageTypePortSuspend           IOUSBHostCIMessageType = 0
)

func (e IOUSBHostCIMessageType) String() string {
	switch e {
	case IOUSBHostCIMessageTypeControllerFrameNumber:
		return "IOUSBHostCIMessageTypeControllerFrameNumber"
	default:
		return fmt.Sprintf("IOUSBHostCIMessageType(%d)", e)
	}
}

type IOUSBHostCINormalTransferData0 uint

const (
	IOUSBHostCINormalTransferData0Length IOUSBHostCINormalTransferData0 = 0
)

func (e IOUSBHostCINormalTransferData0) String() string {
	switch e {
	case IOUSBHostCINormalTransferData0Length:
		return "IOUSBHostCINormalTransferData0Length"
	default:
		return fmt.Sprintf("IOUSBHostCINormalTransferData0(%d)", e)
	}
}

type IOUSBHostCINormalTransferData1 uint

const (
	IOUSBHostCINormalTransferData1Buffer IOUSBHostCINormalTransferData1 = 0
)

func (e IOUSBHostCINormalTransferData1) String() string {
	switch e {
	case IOUSBHostCINormalTransferData1Buffer:
		return "IOUSBHostCINormalTransferData1Buffer"
	default:
		return fmt.Sprintf("IOUSBHostCINormalTransferData1(%d)", e)
	}
}

type IOUSBHostCIPortCapabilitiesMessageControl uint

const (
	IOUSBHostCIPortCapabilitiesMessageControlInternalConnector IOUSBHostCIPortCapabilitiesMessageControl = 0
	IOUSBHostCIPortCapabilitiesMessageControlPortNumber        IOUSBHostCIPortCapabilitiesMessageControl = 0
)

func (e IOUSBHostCIPortCapabilitiesMessageControl) String() string {
	switch e {
	case IOUSBHostCIPortCapabilitiesMessageControlInternalConnector:
		return "IOUSBHostCIPortCapabilitiesMessageControlInternalConnector"
	default:
		return fmt.Sprintf("IOUSBHostCIPortCapabilitiesMessageControl(%d)", e)
	}
}

type IOUSBHostCIPortEventMessageData0Port uint

const (
	IOUSBHostCIPortEventMessageData0PortNumber IOUSBHostCIPortEventMessageData0Port = 0
)

func (e IOUSBHostCIPortEventMessageData0Port) String() string {
	switch e {
	case IOUSBHostCIPortEventMessageData0PortNumber:
		return "IOUSBHostCIPortEventMessageData0PortNumber"
	default:
		return fmt.Sprintf("IOUSBHostCIPortEventMessageData0Port(%d)", e)
	}
}

type IOUSBHostCIPortState uint

const (
	IOUSBHostCIPortStateActive  IOUSBHostCIPortState = 0
	IOUSBHostCIPortStateOff     IOUSBHostCIPortState = 0
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
	IOUSBHostCIPortStatusLinkState     IOUSBHostCIPortStatus = 0
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
	IOUSBHostCIPortStatusCommandData1ChangeMask        IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1ConnectChange     IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1Connected         IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1LinkState         IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1LinkStateChange   IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1LinkStatePhase    IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1Overcurrent       IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1OvercurrentChange IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1Powered           IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1Speed             IOUSBHostCIPortStatusCommandData1 = 0
	IOUSBHostCIPortStatusCommandData1SpeedPhase        IOUSBHostCIPortStatusCommandData1 = 0
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
	IOUSBHostCISetupTransferData1bRequest    IOUSBHostCISetupTransfer = 0
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
	IOUSBHostCITransferCompletionMessageControlDeviceAddress        IOUSBHostCITransferCompletionMessageControl = 0
	IOUSBHostCITransferCompletionMessageControlDeviceAddressPhase   IOUSBHostCITransferCompletionMessageControl = 0
	IOUSBHostCITransferCompletionMessageControlEndpointAddress      IOUSBHostCITransferCompletionMessageControl = 0
	IOUSBHostCITransferCompletionMessageControlEndpointAddressPhase IOUSBHostCITransferCompletionMessageControl = 0
	IOUSBHostCITransferCompletionMessageControlStatus               IOUSBHostCITransferCompletionMessageControl = 0
	IOUSBHostCITransferCompletionMessageControlStatusPhase          IOUSBHostCITransferCompletionMessageControl = 0
)

func (e IOUSBHostCITransferCompletionMessageControl) String() string {
	switch e {
	case IOUSBHostCITransferCompletionMessageControlDeviceAddress:
		return "IOUSBHostCITransferCompletionMessageControlDeviceAddress"
	default:
		return fmt.Sprintf("IOUSBHostCITransferCompletionMessageControl(%d)", e)
	}
}

type IOUSBHostCITransferCompletionMessageData0Transfer uint

const (
	IOUSBHostCITransferCompletionMessageData0TransferLength IOUSBHostCITransferCompletionMessageData0Transfer = 0
)

func (e IOUSBHostCITransferCompletionMessageData0Transfer) String() string {
	switch e {
	case IOUSBHostCITransferCompletionMessageData0TransferLength:
		return "IOUSBHostCITransferCompletionMessageData0TransferLength"
	default:
		return fmt.Sprintf("IOUSBHostCITransferCompletionMessageData0Transfer(%d)", e)
	}
}

type IOUSBHostCITransferCompletionMessageData1Transfer uint

const (
	IOUSBHostCITransferCompletionMessageData1TransferStructure IOUSBHostCITransferCompletionMessageData1Transfer = 0
)

func (e IOUSBHostCITransferCompletionMessageData1Transfer) String() string {
	switch e {
	case IOUSBHostCITransferCompletionMessageData1TransferStructure:
		return "IOUSBHostCITransferCompletionMessageData1TransferStructure"
	default:
		return fmt.Sprintf("IOUSBHostCITransferCompletionMessageData1Transfer(%d)", e)
	}
}

type IOUSBHostCIUserClient uint

const (
	IOUSBHostCIUserClientVersion100 IOUSBHostCIUserClient = 0
)

func (e IOUSBHostCIUserClient) String() string {
	switch e {
	case IOUSBHostCIUserClientVersion100:
		return "IOUSBHostCIUserClientVersion100"
	default:
		return fmt.Sprintf("IOUSBHostCIUserClient(%d)", e)
	}
}

type IfInterfaceAdvisoryDirection uint

const (
	IF_INTERFACE_ADVISORY_DIRECTION_RX IfInterfaceAdvisoryDirection = 0
)

func (e IfInterfaceAdvisoryDirection) String() string {
	switch e {
	case IF_INTERFACE_ADVISORY_DIRECTION_RX:
		return "IF_INTERFACE_ADVISORY_DIRECTION_RX"
	default:
		return fmt.Sprintf("IfInterfaceAdvisoryDirection(%d)", e)
	}
}

type IfInterfaceAdvisoryInterfaceType uint

const (
	IF_INTERFACE_ADVISORY_INTERFACE_TYPE_CELL IfInterfaceAdvisoryInterfaceType = 0
)

func (e IfInterfaceAdvisoryInterfaceType) String() string {
	switch e {
	case IF_INTERFACE_ADVISORY_INTERFACE_TYPE_CELL:
		return "IF_INTERFACE_ADVISORY_INTERFACE_TYPE_CELL"
	default:
		return fmt.Sprintf("IfInterfaceAdvisoryInterfaceType(%d)", e)
	}
}

type IfInterfaceAdvisoryNotificationTypeCellular uint

const (
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_BANDWIDTH_LIMITATION_EVENT    IfInterfaceAdvisoryNotificationTypeCellular = 0
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_DEFAULT                       IfInterfaceAdvisoryNotificationTypeCellular = 0
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_DISCONTINUOUS_RECEPTION_EVENT IfInterfaceAdvisoryNotificationTypeCellular = 0
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_MEASUREMENT_UPDATE            IfInterfaceAdvisoryNotificationTypeCellular = 0
)

func (e IfInterfaceAdvisoryNotificationTypeCellular) String() string {
	switch e {
	case IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_BANDWIDTH_LIMITATION_EVENT:
		return "IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_CELLULAR_BANDWIDTH_LIMITATION_EVENT"
	default:
		return fmt.Sprintf("IfInterfaceAdvisoryNotificationTypeCellular(%d)", e)
	}
}

type IfInterfaceAdvisoryNotificationTypeWifi uint

const (
	IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_WIFI_UNDEFINED IfInterfaceAdvisoryNotificationTypeWifi = 0
)

func (e IfInterfaceAdvisoryNotificationTypeWifi) String() string {
	switch e {
	case IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_WIFI_UNDEFINED:
		return "IF_INTERFACE_ADVISORY_NOTIFICATION_TYPE_WIFI_UNDEFINED"
	default:
		return fmt.Sprintf("IfInterfaceAdvisoryNotificationTypeWifi(%d)", e)
	}
}

type IfNetemModel uint

const (
	IF_NETEM_MODEL_NULL IfNetemModel = 0
)

func (e IfNetemModel) String() string {
	switch e {
	case IF_NETEM_MODEL_NULL:
		return "IF_NETEM_MODEL_NULL"
	default:
		return fmt.Sprintf("IfNetemModel(%d)", e)
	}
}

type IfnetFamily uint

const (
	// IFNET_FAMILY_BOND: # Discussion
	IFNET_FAMILY_BOND IfnetFamily = 0
	// IFNET_FAMILY_CELLULAR: # Discussion
	IFNET_FAMILY_CELLULAR IfnetFamily = 0
	// IFNET_FAMILY_DISC: # Discussion
	IFNET_FAMILY_DISC IfnetFamily = 0
	// IFNET_FAMILY_FAITH: # Discussion
	IFNET_FAMILY_FAITH IfnetFamily = 0
	// IFNET_FAMILY_FIREWIRE: # Discussion
	IFNET_FAMILY_FIREWIRE IfnetFamily = 0
	// IFNET_FAMILY_GIF: # Discussion
	IFNET_FAMILY_GIF   IfnetFamily = 0
	IFNET_FAMILY_IPSEC IfnetFamily = 0
	// IFNET_FAMILY_LOOPBACK: # Discussion
	IFNET_FAMILY_LOOPBACK IfnetFamily = 0
	// IFNET_FAMILY_MDECAP: # Discussion
	IFNET_FAMILY_MDECAP IfnetFamily = 0
	// IFNET_FAMILY_PPP: # Discussion
	IFNET_FAMILY_PPP IfnetFamily = 0
	// IFNET_FAMILY_PVC: # Discussion
	IFNET_FAMILY_PVC IfnetFamily = 0
	// IFNET_FAMILY_SLIP: # Discussion
	IFNET_FAMILY_SLIP IfnetFamily = 0
	// IFNET_FAMILY_STF: # Discussion
	IFNET_FAMILY_STF IfnetFamily = 0
	// IFNET_FAMILY_TUN: # Discussion
	IFNET_FAMILY_TUN       IfnetFamily = 0
	IFNET_FAMILY_UNUSED_16 IfnetFamily = 0
	IFNET_FAMILY_UTUN      IfnetFamily = 0
	// IFNET_FAMILY_VLAN: # Discussion
	IFNET_FAMILY_VLAN IfnetFamily = 0
)

func (e IfnetFamily) String() string {
	switch e {
	case IFNET_FAMILY_BOND:
		return "IFNET_FAMILY_BOND"
	default:
		return fmt.Sprintf("IfnetFamily(%d)", e)
	}
}

type IfnetHw uint

const (
	IFNET_HW_TIMESTAMP IfnetHw = 0
)

func (e IfnetHw) String() string {
	switch e {
	case IFNET_HW_TIMESTAMP:
		return "IFNET_HW_TIMESTAMP"
	default:
		return fmt.Sprintf("IfnetHw(%d)", e)
	}
}

type IfnetLqmThresh uint

const (
	IFNET_LQM_THRESH_ABORT            IfnetLqmThresh = 0
	IFNET_LQM_THRESH_GOOD             IfnetLqmThresh = 0
	IFNET_LQM_THRESH_MINIMALLY_VIABLE IfnetLqmThresh = 0
	IFNET_LQM_THRESH_OFF              IfnetLqmThresh = 0
	IFNET_LQM_THRESH_POOR             IfnetLqmThresh = 0
	IFNET_LQM_THRESH_UNKNOWN          IfnetLqmThresh = 0
)

func (e IfnetLqmThresh) String() string {
	switch e {
	case IFNET_LQM_THRESH_ABORT:
		return "IFNET_LQM_THRESH_ABORT"
	default:
		return fmt.Sprintf("IfnetLqmThresh(%d)", e)
	}
}

type IfnetNpmThresh uint

const (
	IFNET_NPM_THRESH_FAR IfnetNpmThresh = 0
)

func (e IfnetNpmThresh) String() string {
	switch e {
	case IFNET_NPM_THRESH_FAR:
		return "IFNET_NPM_THRESH_FAR"
	default:
		return fmt.Sprintf("IfnetNpmThresh(%d)", e)
	}
}

type IfnetRssi uint

const (
	IFNET_RSSI_UNKNOWN IfnetRssi = 0
)

func (e IfnetRssi) String() string {
	switch e {
	case IFNET_RSSI_UNKNOWN:
		return "IFNET_RSSI_UNKNOWN"
	default:
		return fmt.Sprintf("IfnetRssi(%d)", e)
	}
}

type IfnetSchedModelDriver uint

const (
	IFNET_SCHED_MODEL_DRIVER_MANAGED IfnetSchedModelDriver = 0
)

func (e IfnetSchedModelDriver) String() string {
	switch e {
	case IFNET_SCHED_MODEL_DRIVER_MANAGED:
		return "IFNET_SCHED_MODEL_DRIVER_MANAGED"
	default:
		return fmt.Sprintf("IfnetSchedModelDriver(%d)", e)
	}
}

type IfnetThrottleO uint

const (
	IFNET_THROTTLE_OFF           IfnetThrottleO = 0
	IFNET_THROTTLE_OPPORTUNISTIC IfnetThrottleO = 0
)

func (e IfnetThrottleO) String() string {
	switch e {
	case IFNET_THROTTLE_OFF:
		return "IFNET_THROTTLE_OFF"
	default:
		return fmt.Sprintf("IfnetThrottleO(%d)", e)
	}
}

type IfnetWakeOnMagic uint

const (
	// IFNET_WAKE_ON_MAGIC_PACKET: # Discussion
	IFNET_WAKE_ON_MAGIC_PACKET IfnetWakeOnMagic = 0
)

func (e IfnetWakeOnMagic) String() string {
	switch e {
	case IFNET_WAKE_ON_MAGIC_PACKET:
		return "IFNET_WAKE_ON_MAGIC_PACKET"
	default:
		return fmt.Sprintf("IfnetWakeOnMagic(%d)", e)
	}
}

type In6Clat46Event uint

const (
	IN6_CLAT46_EVENT_V4_FLOW          In6Clat46Event = 0
	IN6_CLAT46_EVENT_V6_ADDR_CONFFAIL In6Clat46Event = 0
)

func (e In6Clat46Event) String() string {
	switch e {
	case IN6_CLAT46_EVENT_V4_FLOW:
		return "IN6_CLAT46_EVENT_V4_FLOW"
	default:
		return fmt.Sprintf("In6Clat46Event(%d)", e)
	}
}

type Indicator uint

const (
	INDICATOR_CAM Indicator = 0
)

func (e Indicator) String() string {
	switch e {
	case INDICATOR_CAM:
		return "INDICATOR_CAM"
	default:
		return fmt.Sprintf("Indicator(%d)", e)
	}
}

type Input uint

const (
	INPUT_DESKTOP_MICROPHONE         Input = 0
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
	IOPMSoftSleep    Iopm = 0
)

func (e Iopm) String() string {
	switch e {
	case IOPMDeviceUsable:
		return "IOPMDeviceUsable"
	default:
		return fmt.Sprintf("Iopm(%d)", e)
	}
}

type IpsecDscpMapping uint

const (
	IPSEC_DSCP_MAPPING_COPY   IpsecDscpMapping = 0
	IPSEC_DSCP_MAPPING_LEGACY IpsecDscpMapping = 0
)

func (e IpsecDscpMapping) String() string {
	switch e {
	case IPSEC_DSCP_MAPPING_COPY:
		return "IPSEC_DSCP_MAPPING_COPY"
	default:
		return fmt.Sprintf("IpsecDscpMapping(%d)", e)
	}
}

type KATA uint

const (
	KATADefaultRetries             KATA = 0
	KATADevice0DeviceID            KATA = 0
	KATADevice1DeviceID            KATA = 0
	KATAEjectRequest               KATA = 0
	KATAFnBusReset                 KATA = 0
	KATAFnExecIO                   KATA = 0
	KATAFnQFlush                   KATA = 0
	KATAFnRegAccess                KATA = 0
	KATAInvalidDeviceID            KATA = 0
	KATALogicalSectorAlignmentMask KATA = 0
	KATAMultipleLogicalSectorsBit  KATA = 0
	KATAMultipleLogicalSectorsMask KATA = 0
	KATANoOp                       KATA = 0
	KATANullEvent                  KATA = 0
	KATAOfflineEvent               KATA = 0
	KATAOfflineRequest             KATA = 0
	KATAOnlineEvent                KATA = 0
	KATAPIFnExecIO                 KATA = 0
	KATAPIResetEvent               KATA = 0
	KATAPhysicalLogicalEnabledBit1 KATA = 0
	KATARemovedEvent               KATA = 0
	KATAReservedEvent              KATA = 0
	KATAResetEvent                 KATA = 0
	KATAZeroRetries                KATA = 0
)

func (e KATA) String() string {
	switch e {
	case KATADefaultRetries:
		return "KATADefaultRetries"
	default:
		return fmt.Sprintf("KATA(%d)", e)
	}
}

type KATADMA uint

const (
	KATADMAErr KATADMA = 0
)

func (e KATADMA) String() string {
	switch e {
	case KATADMAErr:
		return "KATADMAErr"
	default:
		return fmt.Sprintf("KATADMA(%d)", e)
	}
}

type KATADefault uint

const (
	KATADefaultTimeout KATADefault = 0
)

func (e KATADefault) String() string {
	switch e {
	case KATADefaultTimeout:
		return "KATADefaultTimeout"
	default:
		return fmt.Sprintf("KATADefault(%d)", e)
	}
}

type KATADefaultSector uint

const (
	KATADefaultSectorSize KATADefaultSector = 0
)

func (e KATADefaultSector) String() string {
	switch e {
	case KATADefaultSectorSize:
		return "KATADefaultSectorSize"
	default:
		return fmt.Sprintf("KATADefaultSector(%d)", e)
	}
}

type KATAEnable uint

const (
	KATAEnableAPM        KATAEnable = 0
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

type KATAEnableMultiWordDMAMode uint

const (
	KATAEnableMultiWordDMAModeMask KATAEnableMultiWordDMAMode = 0
)

func (e KATAEnableMultiWordDMAMode) String() string {
	switch e {
	case KATAEnableMultiWordDMAModeMask:
		return "KATAEnableMultiWordDMAModeMask"
	default:
		return fmt.Sprintf("KATAEnableMultiWordDMAMode(%d)", e)
	}
}

type KATAForceUnitAccessFeature uint

const (
	KATAForceUnitAccessFeatureBit  KATAForceUnitAccessFeature = 0
	KATAForceUnitAccessFeatureMask KATAForceUnitAccessFeature = 0
)

func (e KATAForceUnitAccessFeature) String() string {
	switch e {
	case KATAForceUnitAccessFeatureBit:
		return "KATAForceUnitAccessFeatureBit"
	default:
		return fmt.Sprintf("KATAForceUnitAccessFeature(%d)", e)
	}
}

type KATAIdentify uint

const (
	KATAIdentifyAdvancedPIOModes            KATAIdentify = 0
	KATAIdentifyCommandExtension1           KATAIdentify = 0
	KATAIdentifyCommandExtension2           KATAIdentify = 0
	KATAIdentifyCommandSetSupported         KATAIdentify = 0
	KATAIdentifyCommandSetSupported2        KATAIdentify = 0
	KATAIdentifyCommandsDefault             KATAIdentify = 0
	KATAIdentifyCommandsEnabled             KATAIdentify = 0
	KATAIdentifyConfiguration               KATAIdentify = 0
	KATAIdentifyCurrentCapacity             KATAIdentify = 0
	KATAIdentifyCurrentCylinders            KATAIdentify = 0
	KATAIdentifyCurrentHeads                KATAIdentify = 0
	KATAIdentifyCurrentMultipleSectors      KATAIdentify = 0
	KATAIdentifyCurrentSectors              KATAIdentify = 0
	KATAIdentifyDriveCapabilities           KATAIdentify = 0
	KATAIdentifyDriveCapabilitiesExtended   KATAIdentify = 0
	KATAIdentifyExtendedInfoSupport         KATAIdentify = 0
	KATAIdentifyFirmwareRevision            KATAIdentify = 0
	KATAIdentifyIntegrity                   KATAIdentify = 0
	KATAIdentifyLBACapacity                 KATAIdentify = 0
	KATAIdentifyLogicalCylinderCount        KATAIdentify = 0
	KATAIdentifyLogicalHeadCount            KATAIdentify = 0
	KATAIdentifyLogicalSectorAlignment      KATAIdentify = 0
	KATAIdentifyMajorVersion                KATAIdentify = 0
	KATAIdentifyMinMultiWordDMATime         KATAIdentify = 0
	KATAIdentifyMinPIOTime                  KATAIdentify = 0
	KATAIdentifyMinPIOTimeWithIORDY         KATAIdentify = 0
	KATAIdentifyMinorVersion                KATAIdentify = 0
	KATAIdentifyModelNumber                 KATAIdentify = 0
	KATAIdentifyMultiWordDMA                KATAIdentify = 0
	KATAIdentifyMultipleSectorCount         KATAIdentify = 0
	KATAIdentifyPIOTiming                   KATAIdentify = 0
	KATAIdentifyPhysicalLogicalSectorSize   KATAIdentify = 0
	KATAIdentifyQueueDepth                  KATAIdentify = 0
	KATAIdentifyRecommendedMultiWordDMATime KATAIdentify = 0
	KATAIdentifySectorsPerTrack             KATAIdentify = 0
	KATAIdentifySerialNumber                KATAIdentify = 0
	KATAIdentifySingleWordDMA               KATAIdentify = 0
	KATAIdentifyUltraDMASupported           KATAIdentify = 0
	KATAIdentifyWordsPerLogicalSector1      KATAIdentify = 0
	KATAIdentifyWordsPerLogicalSector2      KATAIdentify = 0
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
	KATAOperationTypeSMART      KATAOperationType = 0
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
	KATAPIDRQFast   KATAPI = 0
	KATAPIDRQSlow   KATAPI = 0
	KATAPIIRQPacket KATAPI = 0
	KATAPIUnknown   KATAPI = 0
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
	KATASupportsCompactFlashMask    KATASupports = 0
	KATASupportsFlushCacheMask      KATASupports = 0
	KATASupportsPowerManagementMask KATASupports = 0
	KATASupportsSMARTMask           KATASupports = 0
	KATASupportsWriteCacheMask      KATASupports = 0
)

func (e KATASupports) String() string {
	switch e {
	case KATASupportsCompactFlashMask:
		return "KATASupportsCompactFlashMask"
	default:
		return fmt.Sprintf("KATASupports(%d)", e)
	}
}

type KATASupports48BitAddressing uint

const (
	KATASupports48BitAddressingBit KATASupports48BitAddressing = 0
)

func (e KATASupports48BitAddressing) String() string {
	switch e {
	case KATASupports48BitAddressingBit:
		return "KATASupports48BitAddressingBit"
	default:
		return fmt.Sprintf("KATASupports48BitAddressing(%d)", e)
	}
}

type KATASupportsPowerManagement uint

const (
	KATASupportsPowerManagementBit KATASupportsPowerManagement = 0
)

func (e KATASupportsPowerManagement) String() string {
	switch e {
	case KATASupportsPowerManagementBit:
		return "KATASupportsPowerManagementBit"
	default:
		return fmt.Sprintf("KATASupportsPowerManagement(%d)", e)
	}
}

type KATAWriteCacheEnabled uint

const (
	KATAWriteCacheEnabledBit  KATAWriteCacheEnabled = 0
	KATAWriteCacheEnabledMask KATAWriteCacheEnabled = 0
)

func (e KATAWriteCacheEnabled) String() string {
	switch e {
	case KATAWriteCacheEnabledBit:
		return "KATAWriteCacheEnabledBit"
	default:
		return fmt.Sprintf("KATAWriteCacheEnabled(%d)", e)
	}
}

type KATAcmd uint

const (
	KATAcmdCheckPowerMode     KATAcmd = 0
	KATAcmdDiagnostic         KATAcmd = 0
	KATAcmdDoorLock           KATAcmd = 0
	KATAcmdDoorUnlock         KATAcmd = 0
	KATAcmdDriveIdentify      KATAcmd = 0
	KATAcmdFlushCache         KATAcmd = 0
	KATAcmdFlushCacheExtended KATAcmd = 0
	KATAcmdFormatTrack        KATAcmd = 0
	KATAcmdIdle               KATAcmd = 0
	KATAcmdIdleImmed          KATAcmd = 0
	KATAcmdInitDrive          KATAcmd = 0
	KATAcmdMCAcknowledge      KATAcmd = 0
	KATAcmdMediaEject         KATAcmd = 0
	KATAcmdNOP                KATAcmd = 0
	KATAcmdRead               KATAcmd = 0
	KATAcmdReadBuffer         KATAcmd = 0
	KATAcmdReadDMA            KATAcmd = 0
	KATAcmdReadDMAExtended    KATAcmd = 0
	KATAcmdReadExtended       KATAcmd = 0
	KATAcmdReadLong           KATAcmd = 0
	KATAcmdReadMultiple       KATAcmd = 0
	KATAcmdReadVerify         KATAcmd = 0
	KATAcmdRecal              KATAcmd = 0
	KATAcmdSeek               KATAcmd = 0
	KATAcmdSetFeatures        KATAcmd = 0
	KATAcmdSetRWMultiple      KATAcmd = 0
	KATAcmdSleep              KATAcmd = 0
	KATAcmdStandby            KATAcmd = 0
	KATAcmdStandbyImmed       KATAcmd = 0
	KATAcmdWORetry            KATAcmd = 0
	KATAcmdWrite              KATAcmd = 0
	KATAcmdWriteBuffer        KATAcmd = 0
	KATAcmdWriteDMA           KATAcmd = 0
	KATAcmdWriteDMAExtended   KATAcmd = 0
	KATAcmdWriteExtended      KATAcmd = 0
	KATAcmdWriteLong          KATAcmd = 0
	KATAcmdWriteMultiple      KATAcmd = 0
	KATAcmdWriteSame          KATAcmd = 0
	KATAcmdWriteVerify        KATAcmd = 0
)

func (e KATAcmd) String() string {
	switch e {
	case KATAcmdCheckPowerMode:
		return "KATAcmdCheckPowerMode"
	default:
		return fmt.Sprintf("KATAcmd(%d)", e)
	}
}

type KAVC uint

const (
	KAVCAcceptedStatus               KAVC = 0
	KAVCAddress                      KAVC = 0
	KAVCAudio                        KAVC = 0
	KAVCCameraStorage                KAVC = 0
	KAVCChangedStatus                KAVC = 0
	KAVCCommandResponse              KAVC = 0
	KAVCConnectOpcode                KAVC = 0
	KAVCConnectionsOpcode            KAVC = 0
	KAVCControlCommand               KAVC = 0
	KAVCDisconnectOpcode             KAVC = 0
	KAVCDiskRecorder                 KAVC = 0
	KAVCGeneralInquiryCommand        KAVC = 0
	KAVCImplementedStatus            KAVC = 0
	KAVCInTransitionStatus           KAVC = 0
	KAVCInputPlugSignalFormatOpcode  KAVC = 0
	KAVCInputSignalModeOpcode        KAVC = 0
	KAVCInterimStatus                KAVC = 0
	KAVCNotImplementedStatus         KAVC = 0
	KAVCNotifyCommand                KAVC = 0
	KAVCNumSubUnitTypes              KAVC = 0
	KAVCOpcode                       KAVC = 0
	KAVCOperand0                     KAVC = 0
	KAVCOperand1                     KAVC = 0
	KAVCOperand2                     KAVC = 0
	KAVCOperand3                     KAVC = 0
	KAVCOperand4                     KAVC = 0
	KAVCOperand5                     KAVC = 0
	KAVCOperand6                     KAVC = 0
	KAVCOperand7                     KAVC = 0
	KAVCOperand8                     KAVC = 0
	KAVCOutputPlugSignalFormatOpcode KAVC = 0
	KAVCOutputSignalModeOpcode       KAVC = 0
	KAVCPlugInfoOpcode               KAVC = 0
	KAVCPowerOpcode                  KAVC = 0
	KAVCPrinter                      KAVC = 0
	KAVCRejectedStatus               KAVC = 0
	KAVCSignalModeDVCPro525_60       KAVC = 0
	KAVCSignalModeDVCPro625_50       KAVC = 0
	KAVCSignalModeDummyOperand       KAVC = 0
	KAVCSignalModeHD1125_60          KAVC = 0
	KAVCSignalModeHD1250_50          KAVC = 0
	KAVCSignalModeMask_50            KAVC = 0
	KAVCSignalModeMask_DVCPro25      KAVC = 0
	KAVCSignalModeMask_SDL           KAVC = 0
	KAVCSignalModeMask_STYPE         KAVC = 0
	KAVCSignalModeSD525_60           KAVC = 0
	KAVCSignalModeSD625_50           KAVC = 0
	KAVCSignalModeSDL525_60          KAVC = 0
	KAVCSignalModeSDL625_50          KAVC = 0
	KAVCSignalSourceOpcode           KAVC = 0
	KAVCSpecificInquiryCommand       KAVC = 0
	KAVCStatusInquiryCommand         KAVC = 0
	KAVCSubunitInfoOpcode            KAVC = 0
	KAVCTapeRecorder                 KAVC = 0
	KAVCTuner                        KAVC = 0
	KAVCUnitInfoOpcode               KAVC = 0
	KAVCVendorDependentOpcode        KAVC = 0
	KAVCVendorUnique                 KAVC = 0
	KAVCVideoCamera                  KAVC = 0
	KAVCVideoMonitor                 KAVC = 0
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
	KAVCAsyncCommandStateBusReset                KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateCanceled                KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateOutOfMemory             KAVCAsyncCommandState = 0
	KAVCAsyncCommandStatePendingRequest          KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateReceivedFinalResponse   KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateReceivedInterimResponse KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateRequestFailed           KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateRequestSent             KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateTimeOutBeforeResponse   KAVCAsyncCommandState = 0
	KAVCAsyncCommandStateWaitingForResponse      KAVCAsyncCommandState = 0
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
	KAVPowerOn  KAVPower = 0
)

func (e KAVPower) String() string {
	switch e {
	case KAVPowerOff:
		return "KAVPowerOff"
	default:
		return fmt.Sprintf("KAVPower(%d)", e)
	}
}

type KActivate uint

const (
	KActivateConnection KActivate = 0
)

func (e KActivate) String() string {
	switch e {
	case KActivateConnection:
		return "KActivateConnection"
	default:
		return fmt.Sprintf("KActivate(%d)", e)
	}
}

type KAllModesValid uint

const (
	KAllModesValidValue KAllModesValid = 0
	KFastCheckForDDC    KAllModesValid = 0
)

func (e KAllModesValid) String() string {
	switch e {
	case KAllModesValidValue:
		return "KAllModesValidValue"
	default:
		return fmt.Sprintf("KAllModesValid(%d)", e)
	}
}

type KAnd uint

const (
	KAndConnections KAnd = 0
)

func (e KAnd) String() string {
	switch e {
	case KAndConnections:
		return "KAndConnections"
	default:
		return fmt.Sprintf("KAnd(%d)", e)
	}
}

type KAppleBasebandBridgeUserRequest uint

const (
	KAppleBasebandBridgeUserRequest_BasebandHostWake KAppleBasebandBridgeUserRequest = 0
	KAppleBasebandBridgeUserRequest_BasebandInReset  KAppleBasebandBridgeUserRequest = 0
	KAppleBasebandBridgeUserRequest_BridgeReady      KAppleBasebandBridgeUserRequest = 0
	KAppleBasebandBridgeUserRequest_MAX              KAppleBasebandBridgeUserRequest = 0
)

func (e KAppleBasebandBridgeUserRequest) String() string {
	switch e {
	case KAppleBasebandBridgeUserRequest_BasebandHostWake:
		return "KAppleBasebandBridgeUserRequest_BasebandHostWake"
	default:
		return fmt.Sprintf("KAppleBasebandBridgeUserRequest(%d)", e)
	}
}

type KAppleVendorI uint

const (
	KAppleVendorID KAppleVendorI = 0
)

func (e KAppleVendorI) String() string {
	switch e {
	case KAppleVendorID:
		return "KAppleVendorID"
	default:
		return fmt.Sprintf("KAppleVendorI(%d)", e)
	}
}

type KBDFeatures uint

const (
	KBDFeaturesReadBit   KBDFeatures = 0
	KBDFeaturesReadMask  KBDFeatures = 0
	KBDFeaturesWriteBit  KBDFeatures = 0
	KBDFeaturesWriteMask KBDFeatures = 0
)

func (e KBDFeatures) String() string {
	switch e {
	case KBDFeaturesReadBit:
		return "KBDFeaturesReadBit"
	default:
		return fmt.Sprintf("KBDFeatures(%d)", e)
	}
}

type KBDMediaType uint

const (
	KBDMediaTypeMax     KBDMediaType = 0
	KBDMediaTypeMin     KBDMediaType = 0
	KBDMediaTypeR       KBDMediaType = 0
	KBDMediaTypeRE      KBDMediaType = 0
	KBDMediaTypeROM     KBDMediaType = 0
	KBDMediaTypeUnknown KBDMediaType = 0
)

func (e KBDMediaType) String() string {
	switch e {
	case KBDMediaTypeMax:
		return "KBDMediaTypeMax"
	default:
		return fmt.Sprintf("KBDMediaType(%d)", e)
	}
}

type KBTBadClose uint

const (
	KBTBadCloseMask KBTBadClose = 0
)

func (e KBTBadClose) String() string {
	switch e {
	case KBTBadCloseMask:
		return "KBTBadCloseMask"
	default:
		return fmt.Sprintf("KBTBadClose(%d)", e)
	}
}

type KBTHeader uint

const (
	KBTHeaderNode KBTHeader = 0
)

func (e KBTHeader) String() string {
	switch e {
	case KBTHeaderNode:
		return "KBTHeaderNode"
	default:
		return fmt.Sprintf("KBTHeader(%d)", e)
	}
}

type KBestEffortService uint

const (
	KBestEffortServiceLatency KBestEffortService = 0
)

func (e KBestEffortService) String() string {
	switch e {
	case KBestEffortServiceLatency:
		return "KBestEffortServiceLatency"
	default:
		return fmt.Sprintf("KBestEffortService(%d)", e)
	}
}

type KBluetooth uint

const (
	KBluetoothACLLogicalChannelLMP KBluetooth = 0
	KBluetoothAllowRoleSwitch      KBluetooth = 0
	KBluetoothDontAllowRoleSwitch  KBluetooth = 0
	KBluetoothL2CAPMaxPacketSize   KBluetooth = 0
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

type KBluetoothConnectionHandle uint

const (
	KBluetoothConnectionHandleNone KBluetoothConnectionHandle = 0
)

func (e KBluetoothConnectionHandle) String() string {
	switch e {
	case KBluetoothConnectionHandleNone:
		return "KBluetoothConnectionHandleNone"
	default:
		return fmt.Sprintf("KBluetoothConnectionHandle(%d)", e)
	}
}

type KBluetoothDeviceClassMajor uint

const (
	KBluetoothDeviceClassMajorAny            KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorAudio          KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorComputer       KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorEnd            KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorHealth         KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorImaging        KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorLANAccessPoint KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorMiscellaneous  KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorNone           KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorPeripheral     KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorPhone          KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorToy            KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorUnclassified   KBluetoothDeviceClassMajor = 0
	KBluetoothDeviceClassMajorWearable       KBluetoothDeviceClassMajor = 0
)

func (e KBluetoothDeviceClassMajor) String() string {
	switch e {
	case KBluetoothDeviceClassMajorAny:
		return "KBluetoothDeviceClassMajorAny"
	default:
		return fmt.Sprintf("KBluetoothDeviceClassMajor(%d)", e)
	}
}

type KBluetoothDeviceClassMinor uint

const (
	KBluetoothDeviceClassMinorAny                             KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioCamcorder                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioCar                        KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioGamingToy                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioHandsFree                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioHeadphones                 KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioHeadset                    KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioHiFi                       KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioLoudspeaker                KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioMicrophone                 KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioPortable                   KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioReserved1                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioReserved2                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioSetTopBox                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioUnclassified               KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioVCR                        KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioVideoCamera                KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioVideoConferencing          KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioVideoDisplayAndLoudspeaker KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorAudioVideoMonitor               KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorComputerDesktopWorkstation      KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorComputerHandheld                KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorComputerLaptop                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorComputerPalmSized               KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorComputerServer                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorComputerUnclassified            KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorComputerWearable                KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorEnd                             KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorHealthBloodPressureMonitor      KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorHealthDataDisplay               KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorHealthGlucoseMeter              KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorHealthHeartRateMonitor          KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorHealthPulseOximeter             KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorHealthScale                     KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorHealthThermometer               KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorHealthUndefined                 KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorImaging1Camera                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorImaging1Display                 KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorImaging1Printer                 KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorImaging1Scanner                 KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorImaging2Unclassified            KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorNone                            KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral1Combo                KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral1Keyboard             KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral1Pointing             KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2AnyPointing          KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2CardReader           KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2DigitalPen           KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2DigitizerTablet      KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2Gamepad              KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2GesturalInputDevice  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2HandheldScanner      KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2Joystick             KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2RemoteControl        KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2SensingDevice        KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPeripheral2Unclassified         KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPhoneCellular                   KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPhoneCommonISDNAccess           KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPhoneCordless                   KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPhoneSmartPhone                 KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPhoneUnclassified               KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorPhoneWiredModemOrVoiceGateway   KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorToyController                   KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorToyDollActionFigure             KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorToyGame                         KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorToyRobot                        KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorToyVehicle                      KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorWearableGlasses                 KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorWearableHelmet                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorWearableJacket                  KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorWearablePager                   KBluetoothDeviceClassMinor = 0
	KBluetoothDeviceClassMinorWearableWristWatch              KBluetoothDeviceClassMinor = 0
)

func (e KBluetoothDeviceClassMinor) String() string {
	switch e {
	case KBluetoothDeviceClassMinorAny:
		return "KBluetoothDeviceClassMinorAny"
	default:
		return fmt.Sprintf("KBluetoothDeviceClassMinor(%d)", e)
	}
}

type KBluetoothDeviceNameMax uint

const (
	KBluetoothDeviceNameMaxLength KBluetoothDeviceNameMax = 0
)

func (e KBluetoothDeviceNameMax) String() string {
	switch e {
	case KBluetoothDeviceNameMaxLength:
		return "KBluetoothDeviceNameMaxLength"
	default:
		return fmt.Sprintf("KBluetoothDeviceNameMax(%d)", e)
	}
}

type KBluetoothEncryptionEnable uint

const (
	KBluetoothEncryptionEnableBREDRAESCCM KBluetoothEncryptionEnable = 0
	KBluetoothEncryptionEnableBREDRE0     KBluetoothEncryptionEnable = 0
	KBluetoothEncryptionEnableLEAESCCM    KBluetoothEncryptionEnable = 0
	KBluetoothEncryptionEnableOff         KBluetoothEncryptionEnable = 0
	KBluetoothEncryptionEnableOn          KBluetoothEncryptionEnable = 0
)

func (e KBluetoothEncryptionEnable) String() string {
	switch e {
	case KBluetoothEncryptionEnableBREDRAESCCM:
		return "KBluetoothEncryptionEnableBREDRAESCCM"
	default:
		return fmt.Sprintf("KBluetoothEncryptionEnable(%d)", e)
	}
}

type KBluetoothGAPAppearanceGeneric uint

const (
	KBluetoothGAPAppearanceGenericBarcodeScanner       KBluetoothGAPAppearanceGeneric = 0
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

type KBluetoothGeneralInquiryAccessCode uint

const (
	KBluetoothGeneralInquiryAccessCodeIndex KBluetoothGeneralInquiryAccessCode = 0
)

func (e KBluetoothGeneralInquiryAccessCode) String() string {
	switch e {
	case KBluetoothGeneralInquiryAccessCodeIndex:
		return "KBluetoothGeneralInquiryAccessCodeIndex"
	default:
		return fmt.Sprintf("KBluetoothGeneralInquiryAccessCode(%d)", e)
	}
}

type KBluetoothHCI uint

const (
	KBluetoothHCICommandPacketHeaderSize                      KBluetoothHCI = 0
	KBluetoothHCICommandPacketMaxDataSize                     KBluetoothHCI = 0
	KBluetoothHCIDataPacketHeaderSize                         KBluetoothHCI = 0
	KBluetoothHCIDataPacketMaxDataSize                        KBluetoothHCI = 0
	KBluetoothHCIEventAMPReceiverReport                       KBluetoothHCI = 0
	KBluetoothHCIEventAMPStartTest                            KBluetoothHCI = 0
	KBluetoothHCIEventAMPStatusChange                         KBluetoothHCI = 0
	KBluetoothHCIEventAMPTestEnd                              KBluetoothHCI = 0
	KBluetoothHCIEventAuthenticationComplete                  KBluetoothHCI = 0
	KBluetoothHCIEventChangeConnectionLinkKeyComplete         KBluetoothHCI = 0
	KBluetoothHCIEventChannelSelected                         KBluetoothHCI = 0
	KBluetoothHCIEventCommandComplete                         KBluetoothHCI = 0
	KBluetoothHCIEventCommandStatus                           KBluetoothHCI = 0
	KBluetoothHCIEventConnectionComplete                      KBluetoothHCI = 0
	KBluetoothHCIEventConnectionPacketType                    KBluetoothHCI = 0
	KBluetoothHCIEventConnectionRequest                       KBluetoothHCI = 0
	KBluetoothHCIEventDataBufferOverflow                      KBluetoothHCI = 0
	KBluetoothHCIEventDisconnectionComplete                   KBluetoothHCI = 0
	KBluetoothHCIEventDisconnectionLogicalLinkComplete        KBluetoothHCI = 0
	KBluetoothHCIEventDisconnectionPhysicalLinkComplete       KBluetoothHCI = 0
	KBluetoothHCIEventEncryptionChange                        KBluetoothHCI = 0
	KBluetoothHCIEventEncryptionKeyRefreshComplete            KBluetoothHCI = 0
	KBluetoothHCIEventEnhancedFlushComplete                   KBluetoothHCI = 0
	KBluetoothHCIEventExtendedInquiryResult                   KBluetoothHCI = 0
	KBluetoothHCIEventFlowSpecModifyComplete                  KBluetoothHCI = 0
	KBluetoothHCIEventFlowSpecificationComplete               KBluetoothHCI = 0
	KBluetoothHCIEventFlushOccurred                           KBluetoothHCI = 0
	KBluetoothHCIEventHardwareError                           KBluetoothHCI = 0
	KBluetoothHCIEventIOCapabilityRequest                     KBluetoothHCI = 0
	KBluetoothHCIEventIOCapabilityResponse                    KBluetoothHCI = 0
	KBluetoothHCIEventInquiryComplete                         KBluetoothHCI = 0
	KBluetoothHCIEventInquiryResult                           KBluetoothHCI = 0
	KBluetoothHCIEventInquiryResultWithRSSI                   KBluetoothHCI = 0
	KBluetoothHCIEventKeypressNotification                    KBluetoothHCI = 0
	KBluetoothHCIEventLEMetaEvent                             KBluetoothHCI = 0
	KBluetoothHCIEventLinkKeyNotification                     KBluetoothHCI = 0
	KBluetoothHCIEventLinkKeyRequest                          KBluetoothHCI = 0
	KBluetoothHCIEventLinkSupervisionTimeoutChanged           KBluetoothHCI = 0
	KBluetoothHCIEventLogicalLinkComplete                     KBluetoothHCI = 0
	KBluetoothHCIEventLogoTesting                             KBluetoothHCI = 0
	KBluetoothHCIEventLoopbackCommand                         KBluetoothHCI = 0
	KBluetoothHCIEventMasterLinkKeyComplete                   KBluetoothHCI = 0
	KBluetoothHCIEventMaxSlotsChange                          KBluetoothHCI = 0
	KBluetoothHCIEventModeChange                              KBluetoothHCI = 0
	KBluetoothHCIEventNumberOfCompletedDataBlocks             KBluetoothHCI = 0
	KBluetoothHCIEventNumberOfCompletedPackets                KBluetoothHCI = 0
	KBluetoothHCIEventPINCodeRequest                          KBluetoothHCI = 0
	KBluetoothHCIEventPacketHeaderSize                        KBluetoothHCI = 0
	KBluetoothHCIEventPacketMaxDataSize                       KBluetoothHCI = 0
	KBluetoothHCIEventPageScanModeChange                      KBluetoothHCI = 0
	KBluetoothHCIEventPageScanRepetitionModeChange            KBluetoothHCI = 0
	KBluetoothHCIEventPhysicalLinkComplete                    KBluetoothHCI = 0
	KBluetoothHCIEventPhysicalLinkLossEarlyWarning            KBluetoothHCI = 0
	KBluetoothHCIEventPhysicalLinkRecovery                    KBluetoothHCI = 0
	KBluetoothHCIEventQoSSetupComplete                        KBluetoothHCI = 0
	KBluetoothHCIEventQoSViolation                            KBluetoothHCI = 0
	KBluetoothHCIEventReadClockOffsetComplete                 KBluetoothHCI = 0
	KBluetoothHCIEventReadRemoteExtendedFeaturesComplete      KBluetoothHCI = 0
	KBluetoothHCIEventReadRemoteSupportedFeaturesComplete     KBluetoothHCI = 0
	KBluetoothHCIEventReadRemoteVersionInformationComplete    KBluetoothHCI = 0
	KBluetoothHCIEventRemoteHostSupportedFeaturesNotification KBluetoothHCI = 0
	KBluetoothHCIEventRemoteNameRequestComplete               KBluetoothHCI = 0
	KBluetoothHCIEventRemoteOOBDataRequest                    KBluetoothHCI = 0
	KBluetoothHCIEventReturnLinkKeys                          KBluetoothHCI = 0
	KBluetoothHCIEventRoleChange                              KBluetoothHCI = 0
	KBluetoothHCIEventShortRangeModeChangeComplete            KBluetoothHCI = 0
	KBluetoothHCIEventSimplePairingComplete                   KBluetoothHCI = 0
	KBluetoothHCIEventSniffSubrating                          KBluetoothHCI = 0
	KBluetoothHCIEventSynchronousConnectionChanged            KBluetoothHCI = 0
	KBluetoothHCIEventSynchronousConnectionComplete           KBluetoothHCI = 0
	KBluetoothHCIEventUserConfirmationRequest                 KBluetoothHCI = 0
	KBluetoothHCIEventUserPasskeyNotification                 KBluetoothHCI = 0
	KBluetoothHCIEventUserPasskeyRequest                      KBluetoothHCI = 0
	KBluetoothHCIEventVendorSpecific                          KBluetoothHCI = 0
	KBluetoothHCIMaxCommandPacketSize                         KBluetoothHCI = 0
	KBluetoothHCIMaxDataPacketSize                            KBluetoothHCI = 0
	KBluetoothHCIMaxEventPacketSize                           KBluetoothHCI = 0
	KBluetoothHCISubEventLEAdvertisingReport                  KBluetoothHCI = 0
	KBluetoothHCISubEventLEAdvertisingSetTerminated           KBluetoothHCI = 0
	KBluetoothHCISubEventLEChannelSelectionAlgorithm          KBluetoothHCI = 0
	KBluetoothHCISubEventLEConnectionComplete                 KBluetoothHCI = 0
	KBluetoothHCISubEventLEConnectionUpdateComplete           KBluetoothHCI = 0
	KBluetoothHCISubEventLEDataLengthChange                   KBluetoothHCI = 0
	KBluetoothHCISubEventLEDirectAdvertisingReport            KBluetoothHCI = 0
	KBluetoothHCISubEventLEEnhancedConnectionComplete         KBluetoothHCI = 0
	KBluetoothHCISubEventLEExtendedAdvertising                KBluetoothHCI = 0
	KBluetoothHCISubEventLEGenerateDHKeyComplete              KBluetoothHCI = 0
	KBluetoothHCISubEventLELongTermKeyRequest                 KBluetoothHCI = 0
	KBluetoothHCISubEventLEPeriodicAdvertisingReport          KBluetoothHCI = 0
	KBluetoothHCISubEventLEPeriodicAdvertisingSyncEstablished KBluetoothHCI = 0
	KBluetoothHCISubEventLEPeriodicAdvertisingSyncLost        KBluetoothHCI = 0
	KBluetoothHCISubEventLEPhyUpdateComplete                  KBluetoothHCI = 0
	KBluetoothHCISubEventLEReadLocalP256PublicKeyComplete     KBluetoothHCI = 0
	KBluetoothHCISubEventLEReadRemoteUsedFeaturesComplete     KBluetoothHCI = 0
	KBluetoothHCISubEventLERemoteConnectionParameterRequest   KBluetoothHCI = 0
	KBluetoothHCISubEventLEScanRequestReceived                KBluetoothHCI = 0
	KBluetoothHCISubEventLEScanTimeout                        KBluetoothHCI = 0
)

func (e KBluetoothHCI) String() string {
	switch e {
	case KBluetoothHCICommandPacketHeaderSize:
		return "KBluetoothHCICommandPacketHeaderSize"
	default:
		return fmt.Sprintf("KBluetoothHCI(%d)", e)
	}
}

type KBluetoothHCICommand uint

const (
	KBluetoothHCICommandLEWriteSuggestedDefaultDataLength KBluetoothHCICommand = 0
	KBluetoothHCICommandReadLocalAMPASSOC                 KBluetoothHCICommand = 0
)

func (e KBluetoothHCICommand) String() string {
	switch e {
	case KBluetoothHCICommandLEWriteSuggestedDefaultDataLength:
		return "KBluetoothHCICommandLEWriteSuggestedDefaultDataLength"
	default:
		return fmt.Sprintf("KBluetoothHCICommand(%d)", e)
	}
}

type KBluetoothHCIErroneousDataReporting int

const (
	KBluetoothHCIErroneousDataReportingDisabled KBluetoothHCIErroneousDataReporting = 0
)

func (e KBluetoothHCIErroneousDataReporting) String() string {
	switch e {
	case KBluetoothHCIErroneousDataReportingDisabled:
		return "KBluetoothHCIErroneousDataReportingDisabled"
	default:
		return fmt.Sprintf("KBluetoothHCIErroneousDataReporting(%d)", e)
	}
}

type KBluetoothHCIError int

const (
	KBluetoothHCIErrorACLConnectionAlreadyExists                    KBluetoothHCIError = 0
	KBluetoothHCIErrorAuthenticationFailure                         KBluetoothHCIError = 0
	KBluetoothHCIErrorChannelClassificationNotSupported             KBluetoothHCIError = 0
	KBluetoothHCIErrorCoarseClockAdjustmentRejected                 KBluetoothHCIError = 0
	KBluetoothHCIErrorCommandDisallowed                             KBluetoothHCIError = 0
	KBluetoothHCIErrorConnectionFailedToBeEstablished               KBluetoothHCIError = 0
	KBluetoothHCIErrorConnectionRejectedDueToNoSuitableChannelFound KBluetoothHCIError = 0
	KBluetoothHCIErrorConnectionTerminatedByLocalHost               KBluetoothHCIError = 0
	KBluetoothHCIErrorConnectionTerminatedDueToMICFailure           KBluetoothHCIError = 0
	KBluetoothHCIErrorConnectionTimeout                             KBluetoothHCIError = 0
	KBluetoothHCIErrorControllerBusy                                KBluetoothHCIError = 0
	KBluetoothHCIErrorDifferentTransactionCollision                 KBluetoothHCIError = 0
	KBluetoothHCIErrorDirectedAdvertisingTimeout                    KBluetoothHCIError = 0
	KBluetoothHCIErrorEncryptionModeNotAcceptable                   KBluetoothHCIError = 0
	KBluetoothHCIErrorExtendedInquiryResponseTooLarge               KBluetoothHCIError = 0
	KBluetoothHCIErrorHardwareFailure                               KBluetoothHCIError = 0
	KBluetoothHCIErrorHostBusyPairing                               KBluetoothHCIError = 0
	KBluetoothHCIErrorHostRejectedLimitedResources                  KBluetoothHCIError = 0
	KBluetoothHCIErrorHostRejectedRemoteDeviceIsPersonal            KBluetoothHCIError = 0
	KBluetoothHCIErrorHostRejectedSecurityReasons                   KBluetoothHCIError = 0
	KBluetoothHCIErrorHostRejectedUnacceptableDeviceAddress         KBluetoothHCIError = 0
	KBluetoothHCIErrorHostTimeout                                   KBluetoothHCIError = 0
	KBluetoothHCIErrorInstantPassed                                 KBluetoothHCIError = 0
	KBluetoothHCIErrorInsufficientSecurity                          KBluetoothHCIError = 0
	KBluetoothHCIErrorInvalidHCICommandParameters                   KBluetoothHCIError = 0
	KBluetoothHCIErrorInvalidLMPParameters                          KBluetoothHCIError = 0
	KBluetoothHCIErrorKeyMissing                                    KBluetoothHCIError = 0
	KBluetoothHCIErrorLMPErrorTransactionCollision                  KBluetoothHCIError = 0
	KBluetoothHCIErrorLMPPDUNotAllowed                              KBluetoothHCIError = 0
	KBluetoothHCIErrorLMPResponseTimeout                            KBluetoothHCIError = 0
	KBluetoothHCIErrorMACConnectionFailed                           KBluetoothHCIError = 0
	KBluetoothHCIErrorMax                                           KBluetoothHCIError = 0
	KBluetoothHCIErrorMaxNumberOfConnections                        KBluetoothHCIError = 0
	KBluetoothHCIErrorMaxNumberOfSCOConnectionsToADevice            KBluetoothHCIError = 0
	KBluetoothHCIErrorMemoryFull                                    KBluetoothHCIError = 0
	KBluetoothHCIErrorNoConnection                                  KBluetoothHCIError = 0
	KBluetoothHCIErrorOtherEndTerminatedConnectionAboutToPowerOff   KBluetoothHCIError = 0
	KBluetoothHCIErrorOtherEndTerminatedConnectionLowResources      KBluetoothHCIError = 0
	KBluetoothHCIErrorOtherEndTerminatedConnectionUserEnded         KBluetoothHCIError = 0
	KBluetoothHCIErrorPageTimeout                                   KBluetoothHCIError = 0
	KBluetoothHCIErrorPairingNotAllowed                             KBluetoothHCIError = 0
	KBluetoothHCIErrorPairingWithUnitKeyNotSupported                KBluetoothHCIError = 0
	KBluetoothHCIErrorParameterOutOfMandatoryRange                  KBluetoothHCIError = 0
	KBluetoothHCIErrorQoSNotSupported                               KBluetoothHCIError = 0
	KBluetoothHCIErrorQoSRejected                                   KBluetoothHCIError = 0
	KBluetoothHCIErrorQoSUnacceptableParameter                      KBluetoothHCIError = 0
	KBluetoothHCIErrorRepeatedAttempts                              KBluetoothHCIError = 0
	KBluetoothHCIErrorReservedSlotViolation                         KBluetoothHCIError = 0
	KBluetoothHCIErrorRoleChangeNotAllowed                          KBluetoothHCIError = 0
	KBluetoothHCIErrorRoleSwitchFailed                              KBluetoothHCIError = 0
	KBluetoothHCIErrorRoleSwitchPending                             KBluetoothHCIError = 0
	KBluetoothHCIErrorSCOAirModeRejected                            KBluetoothHCIError = 0
	KBluetoothHCIErrorSCOIntervalRejected                           KBluetoothHCIError = 0
	KBluetoothHCIErrorSCOOffsetRejected                             KBluetoothHCIError = 0
	KBluetoothHCIErrorSecureSimplePairingNotSupportedByHost         KBluetoothHCIError = 0
	KBluetoothHCIErrorSuccess                                       KBluetoothHCIError = 0
	KBluetoothHCIErrorUnacceptableConnectionInterval                KBluetoothHCIError = 0
	KBluetoothHCIErrorUnitKeyUsed                                   KBluetoothHCIError = 0
	KBluetoothHCIErrorUnknownHCICommand                             KBluetoothHCIError = 0
	KBluetoothHCIErrorUnknownLMPPDU                                 KBluetoothHCIError = 0
	KBluetoothHCIErrorUnspecifiedError                              KBluetoothHCIError = 0
	KBluetoothHCIErrorUnsupportedFeatureOrParameterValue            KBluetoothHCIError = 0
	KBluetoothHCIErrorUnsupportedLMPParameterValue                  KBluetoothHCIError = 0
	KBluetoothHCIErrorUnsupportedRemoteFeature                      KBluetoothHCIError = 0
)

func (e KBluetoothHCIError) String() string {
	switch e {
	case KBluetoothHCIErrorACLConnectionAlreadyExists:
		return "KBluetoothHCIErrorACLConnectionAlreadyExists"
	default:
		return fmt.Sprintf("KBluetoothHCIError(%d)", e)
	}
}

type KBluetoothHCIErrorPowerIsOF int

const (
	KBluetoothHCIErrorPowerIsOFF KBluetoothHCIErrorPowerIsOF = 0
)

func (e KBluetoothHCIErrorPowerIsOF) String() string {
	switch e {
	case KBluetoothHCIErrorPowerIsOFF:
		return "KBluetoothHCIErrorPowerIsOFF"
	default:
		return fmt.Sprintf("KBluetoothHCIErrorPowerIsOF(%d)", e)
	}
}

type KBluetoothHCIEventMask uint

const (
	KBluetoothHCIEventMaskConnectionPacketTypeChanged KBluetoothHCIEventMask = 0
	KBluetoothHCIEventMaskFlushOccurred               KBluetoothHCIEventMask = 0
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
	KBluetoothHCILoopbackModeLocal  KBluetoothHCILoopbackMode = 0
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

type KBluetoothHCITransportUSBClass uint

const (
	// KBluetoothHCITransportUSBClassCode: # Discussion
	KBluetoothHCITransportUSBClassCode KBluetoothHCITransportUSBClass = 0
)

func (e KBluetoothHCITransportUSBClass) String() string {
	switch e {
	case KBluetoothHCITransportUSBClassCode:
		return "KBluetoothHCITransportUSBClassCode"
	default:
		return fmt.Sprintf("KBluetoothHCITransportUSBClass(%d)", e)
	}
}

type KBluetoothKeyFlag uint

const (
	KBluetoothKeyFlagSemiPermanent KBluetoothKeyFlag = 0
	KBluetoothKeyFlagTemporary     KBluetoothKeyFlag = 0
)

func (e KBluetoothKeyFlag) String() string {
	switch e {
	case KBluetoothKeyFlagSemiPermanent:
		return "KBluetoothKeyFlagSemiPermanent"
	default:
		return fmt.Sprintf("KBluetoothKeyFlag(%d)", e)
	}
}

type KBluetoothKeyType uint

const (
	KBluetoothKeyTypeAuthenticatedCombination       KBluetoothKeyType = 0
	KBluetoothKeyTypeAuthenticatedCombinationP256   KBluetoothKeyType = 0
	KBluetoothKeyTypeChangedCombination             KBluetoothKeyType = 0
	KBluetoothKeyTypeCombination                    KBluetoothKeyType = 0
	KBluetoothKeyTypeDebugCombination               KBluetoothKeyType = 0
	KBluetoothKeyTypeLocalUnit                      KBluetoothKeyType = 0
	KBluetoothKeyTypeRemoteUnit                     KBluetoothKeyType = 0
	KBluetoothKeyTypeUnauthenticatedCombination     KBluetoothKeyType = 0
	KBluetoothKeyTypeUnauthenticatedCombinationP256 KBluetoothKeyType = 0
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
	KBluetoothL2CAPFlushTimeoutDefault       KBluetoothL2CAP = 0
	KBluetoothL2CAPMTUDefault                KBluetoothL2CAP = 0
	KBluetoothL2CAPMTULowEnergyDefault       KBluetoothL2CAP = 0
	KBluetoothL2CAPMTULowEnergyMax           KBluetoothL2CAP = 0
	KBluetoothL2CAPMTUMaximum                KBluetoothL2CAP = 0
	KBluetoothL2CAPMTUMinimum                KBluetoothL2CAP = 0
	KBluetoothL2CAPMTUSIG                    KBluetoothL2CAP = 0
	KBluetoothL2CAPMTUStart                  KBluetoothL2CAP = 0
	KBluetoothL2CAPQoSDelayVariationDefault  KBluetoothL2CAP = 0
	KBluetoothL2CAPQoSFlagsDefault           KBluetoothL2CAP = 0
	KBluetoothL2CAPQoSLatencyDefault         KBluetoothL2CAP = 0
	KBluetoothL2CAPQoSPeakBandwidthDefault   KBluetoothL2CAP = 0
	KBluetoothL2CAPQoSTokenBucketSizeDefault KBluetoothL2CAP = 0
	KBluetoothL2CAPQoSTokenRateDefault       KBluetoothL2CAP = 0
	KBluetoothL2CAPQoSTypeDefault            KBluetoothL2CAP = 0
)

func (e KBluetoothL2CAP) String() string {
	switch e {
	case KBluetoothL2CAPFlushTimeoutDefault:
		return "KBluetoothL2CAPFlushTimeoutDefault"
	default:
		return fmt.Sprintf("KBluetoothL2CAP(%d)", e)
	}
}

type KBluetoothL2CAPChannel uint

const (
	KBluetoothL2CAPChannelLESignalling KBluetoothL2CAPChannel = 0
	KBluetoothL2CAPChannelMagnet       KBluetoothL2CAPChannel = 0
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
	KBluetoothL2CAPConfigurationOptionFlushTimeoutLength                 KBluetoothL2CAPConfigurationOption = 0
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
	KBluetoothL2CAPFlushTimeoutForever     KBluetoothL2CAPFlushTimeout = 0
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

type KBluetoothL2CAPInfoTypeMaxConnectionlessMTU uint

const (
	KBluetoothL2CAPInfoTypeMaxConnectionlessMTUSize KBluetoothL2CAPInfoTypeMaxConnectionlessMTU = 0
)

func (e KBluetoothL2CAPInfoTypeMaxConnectionlessMTU) String() string {
	switch e {
	case KBluetoothL2CAPInfoTypeMaxConnectionlessMTUSize:
		return "KBluetoothL2CAPInfoTypeMaxConnectionlessMTUSize"
	default:
		return fmt.Sprintf("KBluetoothL2CAPInfoTypeMaxConnectionlessMTU(%d)", e)
	}
}

type KBluetoothL2CAPPSM uint

const (
	KBluetoothL2CAPPSMAACP             KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMATT              KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMAVCTP            KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMAVCTP_Browsing   KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMAVDTP            KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMBNEP             KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMD2D              KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMDynamicEnd       KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMDynamicStart     KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMHIDControl       KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMHIDInterrupt     KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMNone             KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMRFCOMM           KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMReservedEnd      KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMReservedStart    KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMSDP              KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMTCS_BIN          KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMTCS_BIN_Cordless KBluetoothL2CAPPSM = 0
	KBluetoothL2CAPPSMUID_C_Plane      KBluetoothL2CAPPSM = 0
)

func (e KBluetoothL2CAPPSM) String() string {
	switch e {
	case KBluetoothL2CAPPSMAACP:
		return "KBluetoothL2CAPPSMAACP"
	default:
		return fmt.Sprintf("KBluetoothL2CAPPSM(%d)", e)
	}
}

type KBluetoothL2CAPPacketHeader uint

const (
	KBluetoothL2CAPPacketHeaderSize KBluetoothL2CAPPacketHeader = 0
)

func (e KBluetoothL2CAPPacketHeader) String() string {
	switch e {
	case KBluetoothL2CAPPacketHeaderSize:
		return "KBluetoothL2CAPPacketHeaderSize"
	default:
		return fmt.Sprintf("KBluetoothL2CAPPacketHeader(%d)", e)
	}
}

type KBluetoothL2CAPTCICommandL2CA uint

const (
	KBluetoothL2CAPTCICommandL2CA_ConfigReq  KBluetoothL2CAPTCICommandL2CA = 0
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

type KBluetoothL2CAPTCIEventID uint

const (
	KBluetoothL2CAPTCIEventIDL2CA_ConfigInd       KBluetoothL2CAPTCIEventID = 0
	KBluetoothL2CAPTCIEventIDL2CA_ConnectInd      KBluetoothL2CAPTCIEventID = 0
	KBluetoothL2CAPTCIEventIDL2CA_DisconnectInd   KBluetoothL2CAPTCIEventID = 0
	KBluetoothL2CAPTCIEventIDL2CA_QoSViolationInd KBluetoothL2CAPTCIEventID = 0
	KBluetoothL2CAPTCIEventIDL2CA_TimeOutInd      KBluetoothL2CAPTCIEventID = 0
	KBluetoothL2CAPTCIEventIDReserved             KBluetoothL2CAPTCIEventID = 0
)

func (e KBluetoothL2CAPTCIEventID) String() string {
	switch e {
	case KBluetoothL2CAPTCIEventIDL2CA_ConfigInd:
		return "KBluetoothL2CAPTCIEventIDL2CA_ConfigInd"
	default:
		return fmt.Sprintf("KBluetoothL2CAPTCIEventID(%d)", e)
	}
}

type KBluetoothLEMaxTXOctets uint

const (
	KBluetoothLEMaxTXOctetsDefault KBluetoothLEMaxTXOctets = 0
)

func (e KBluetoothLEMaxTXOctets) String() string {
	switch e {
	case KBluetoothLEMaxTXOctetsDefault:
		return "KBluetoothLEMaxTXOctetsDefault"
	default:
		return fmt.Sprintf("KBluetoothLEMaxTXOctets(%d)", e)
	}
}

type KBluetoothLESecurityManager uint

const (
	KBluetoothLESecurityManagerBonding KBluetoothLESecurityManager = 0
)

func (e KBluetoothLESecurityManager) String() string {
	switch e {
	case KBluetoothLESecurityManagerBonding:
		return "KBluetoothLESecurityManagerBonding"
	default:
		return fmt.Sprintf("KBluetoothLESecurityManager(%d)", e)
	}
}

type KBluetoothLETXOctets uint

const (
	KBluetoothLETXOctetsDefault KBluetoothLETXOctets = 0
)

func (e KBluetoothLETXOctets) String() string {
	switch e {
	case KBluetoothLETXOctetsDefault:
		return "KBluetoothLETXOctetsDefault"
	default:
		return fmt.Sprintf("KBluetoothLETXOctets(%d)", e)
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

type KBluetoothPageScanMode uint

const (
	KBluetoothPageScanModeMandatory KBluetoothPageScanMode = 0
)

func (e KBluetoothPageScanMode) String() string {
	switch e {
	case KBluetoothPageScanModeMandatory:
		return "KBluetoothPageScanModeMandatory"
	default:
		return fmt.Sprintf("KBluetoothPageScanMode(%d)", e)
	}
}

type KBluetoothPageScanPeriodMode uint

const (
	KBluetoothPageScanPeriodModeP0 KBluetoothPageScanPeriodMode = 0
	KBluetoothPageScanPeriodModeP1 KBluetoothPageScanPeriodMode = 0
	KBluetoothPageScanPeriodModeP2 KBluetoothPageScanPeriodMode = 0
)

func (e KBluetoothPageScanPeriodMode) String() string {
	switch e {
	case KBluetoothPageScanPeriodModeP0:
		return "KBluetoothPageScanPeriodModeP0"
	default:
		return fmt.Sprintf("KBluetoothPageScanPeriodMode(%d)", e)
	}
}

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
	KBluetoothRoleBecomeCentral    KBluetoothRole = 0
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
	KBluetoothSDPAttributeDeviceIdentifierDocumentationURL    KBluetoothSDPAttributeDeviceIdentifier = 0
)

func (e KBluetoothSDPAttributeDeviceIdentifier) String() string {
	switch e {
	case KBluetoothSDPAttributeDeviceIdentifierClientExecutableURL:
		return "KBluetoothSDPAttributeDeviceIdentifierClientExecutableURL"
	default:
		return fmt.Sprintf("KBluetoothSDPAttributeDeviceIdentifier(%d)", e)
	}
}

type KBluetoothSDPDataElementType uint

const (
	KBluetoothSDPDataElementTypeBoolean KBluetoothSDPDataElementType = 0
)

func (e KBluetoothSDPDataElementType) String() string {
	switch e {
	case KBluetoothSDPDataElementTypeBoolean:
		return "KBluetoothSDPDataElementTypeBoolean"
	default:
		return fmt.Sprintf("KBluetoothSDPDataElementType(%d)", e)
	}
}

type KBluetoothSDPErrorCodeReserved int

const (
	KBluetoothSDPErrorCodeReservedEnd   KBluetoothSDPErrorCodeReserved = 0
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
	KBluetoothSDPPDUIDErrorResponse                  KBluetoothSDPPDUID = 0
	KBluetoothSDPPDUIDReserved                       KBluetoothSDPPDUID = 0
	KBluetoothSDPPDUIDServiceAttributeRequest        KBluetoothSDPPDUID = 0
	KBluetoothSDPPDUIDServiceAttributeResponse       KBluetoothSDPPDUID = 0
	KBluetoothSDPPDUIDServiceSearchAttributeRequest  KBluetoothSDPPDUID = 0
	KBluetoothSDPPDUIDServiceSearchAttributeResponse KBluetoothSDPPDUID = 0
	KBluetoothSDPPDUIDServiceSearchRequest           KBluetoothSDPPDUID = 0
	KBluetoothSDPPDUIDServiceSearchResponse          KBluetoothSDPPDUID = 0
)

func (e KBluetoothSDPPDUID) String() string {
	switch e {
	case KBluetoothSDPPDUIDErrorResponse:
		return "KBluetoothSDPPDUIDErrorResponse"
	default:
		return fmt.Sprintf("KBluetoothSDPPDUID(%d)", e)
	}
}

type KBluetoothSDPProtocolParameter uint

const (
	KBluetoothSDPProtocolParameterBNEPSupportedNetworkPacketTypeList KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterBNEPVersion                        KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterL2CAPPSM                           KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterRFCOMMChannel                      KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterTCPPort                            KBluetoothSDPProtocolParameter = 0
	KBluetoothSDPProtocolParameterUDPPort                            KBluetoothSDPProtocolParameter = 0
)

func (e KBluetoothSDPProtocolParameter) String() string {
	switch e {
	case KBluetoothSDPProtocolParameterBNEPSupportedNetworkPacketTypeList:
		return "KBluetoothSDPProtocolParameterBNEPSupportedNetworkPacketTypeList"
	default:
		return fmt.Sprintf("KBluetoothSDPProtocolParameter(%d)", e)
	}
}

type KBluetoothSDPUUID16AVCT uint

const (
	KBluetoothSDPUUID16AVCTP KBluetoothSDPUUID16AVCT = 0
)

func (e KBluetoothSDPUUID16AVCT) String() string {
	switch e {
	case KBluetoothSDPUUID16AVCTP:
		return "KBluetoothSDPUUID16AVCTP"
	default:
		return fmt.Sprintf("KBluetoothSDPUUID16AVCT(%d)", e)
	}
}

type KBluetoothSDPUUID16ServiceClass uint

const (
	KBluetoothSDPUUID16ServiceClassAVRemoteControl  KBluetoothSDPUUID16ServiceClass = 0
	KBluetoothSDPUUID16ServiceClassDialupNetworking KBluetoothSDPUUID16ServiceClass = 0
	KBluetoothSDPUUID16ServiceClassGN               KBluetoothSDPUUID16ServiceClass = 0
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
	KBluetoothServiceClassMajorAny       KBluetoothServiceClassMajor = 0
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

type KBluetoothSynchronousConnectionPacketType3EV5 uint

const (
	KBluetoothSynchronousConnectionPacketType3EV5Omit KBluetoothSynchronousConnectionPacketType3EV5 = 0
)

func (e KBluetoothSynchronousConnectionPacketType3EV5) String() string {
	switch e {
	case KBluetoothSynchronousConnectionPacketType3EV5Omit:
		return "KBluetoothSynchronousConnectionPacketType3EV5Omit"
	default:
		return fmt.Sprintf("KBluetoothSynchronousConnectionPacketType3EV5(%d)", e)
	}
}

type KBluetoothVoiceSettingAirCodingFormatA uint

const (
	KBluetoothVoiceSettingAirCodingFormatALaw KBluetoothVoiceSettingAirCodingFormatA = 0
)

func (e KBluetoothVoiceSettingAirCodingFormatA) String() string {
	switch e {
	case KBluetoothVoiceSettingAirCodingFormatALaw:
		return "KBluetoothVoiceSettingAirCodingFormatALaw"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingAirCodingFormatA(%d)", e)
	}
}

type KBluetoothVoiceSettingInputCoding uint

const (
	KBluetoothVoiceSettingInputCodingALawInputCoding   KBluetoothVoiceSettingInputCoding = 0
	KBluetoothVoiceSettingInputCodingLinearInputCoding KBluetoothVoiceSettingInputCoding = 0
	KBluetoothVoiceSettingInputCodingMask              KBluetoothVoiceSettingInputCoding = 0
	KBluetoothVoiceSettingInputCodingULawInputCoding   KBluetoothVoiceSettingInputCoding = 0
)

func (e KBluetoothVoiceSettingInputCoding) String() string {
	switch e {
	case KBluetoothVoiceSettingInputCodingALawInputCoding:
		return "KBluetoothVoiceSettingInputCodingALawInputCoding"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingInputCoding(%d)", e)
	}
}

type KBluetoothVoiceSettingInputDataFormat uint

const (
	KBluetoothVoiceSettingInputDataFormatMask          KBluetoothVoiceSettingInputDataFormat = 0
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

type KBluetoothVoiceSettingInputSampleSize16 uint

const (
	KBluetoothVoiceSettingInputSampleSize16Bit KBluetoothVoiceSettingInputSampleSize16 = 0
)

func (e KBluetoothVoiceSettingInputSampleSize16) String() string {
	switch e {
	case KBluetoothVoiceSettingInputSampleSize16Bit:
		return "KBluetoothVoiceSettingInputSampleSize16Bit"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingInputSampleSize16(%d)", e)
	}
}

type KBluetoothVoiceSettingPCMBitPosition uint

const (
	KBluetoothVoiceSettingPCMBitPositionMask KBluetoothVoiceSettingPCMBitPosition = 0
)

func (e KBluetoothVoiceSettingPCMBitPosition) String() string {
	switch e {
	case KBluetoothVoiceSettingPCMBitPositionMask:
		return "KBluetoothVoiceSettingPCMBitPositionMask"
	default:
		return fmt.Sprintf("KBluetoothVoiceSettingPCMBitPosition(%d)", e)
	}
}

type KBootDriverType uint

const (
	KBootDriverTypeInvalid KBootDriverType = 0
)

func (e KBootDriverType) String() string {
	switch e {
	case KBootDriverTypeInvalid:
		return "KBootDriverTypeInvalid"
	default:
		return fmt.Sprintf("KBootDriverType(%d)", e)
	}
}

type KBootROMType uint

const (
	KBootROMTypeNewWorld KBootROMType = 0
	KBootROMTypeOldWorld KBootROMType = 0
)

func (e KBootROMType) String() string {
	switch e {
	case KBootROMTypeNewWorld:
		return "KBootROMTypeNewWorld"
	default:
		return fmt.Sprintf("KBootROMType(%d)", e)
	}
}

type KC0DataMaxString uint

const (
	KC0DataMaxStringLen KC0DataMaxString = 0
)

func (e KC0DataMaxString) String() string {
	switch e {
	case KC0DataMaxStringLen:
		return "KC0DataMaxStringLen"
	default:
		return fmt.Sprintf("KC0DataMaxString(%d)", e)
	}
}

type KCDFeatures uint

const (
	KCDFeaturesAnalogAudioBit        KCDFeatures = 0
	KCDFeaturesAnalogAudioMask       KCDFeatures = 0
	KCDFeaturesBUFWriteBit           KCDFeatures = 0
	KCDFeaturesCDDAStreamAccurateBit KCDFeatures = 0
	KCDFeaturesPacketWriteBit        KCDFeatures = 0
	KCDFeaturesRawWriteBit           KCDFeatures = 0
	KCDFeaturesReWriteableBit        KCDFeatures = 0
	KCDFeaturesReadStructuresBit     KCDFeatures = 0
	KCDFeaturesSAOWriteBit           KCDFeatures = 0
	KCDFeaturesTAOWriteBit           KCDFeatures = 0
	KCDFeaturesTestWriteBit          KCDFeatures = 0
	KCDFeaturesTestWriteMask         KCDFeatures = 0
	KCDFeaturesWriteOnceBit          KCDFeatures = 0
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
	KCDMediaTypeMax     KCDMediaType = 0
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
	KCDSectorAreaAuxiliary   KCDSectorArea = 0
	KCDSectorAreaErrorFlags  KCDSectorArea = 0
	KCDSectorAreaHeader      KCDSectorArea = 0
	KCDSectorAreaSubChannel  KCDSectorArea = 0
	KCDSectorAreaSubChannelQ KCDSectorArea = 0
	KCDSectorAreaSubHeader   KCDSectorArea = 0
	KCDSectorAreaSync        KCDSectorArea = 0
	KCDSectorAreaUser        KCDSectorArea = 0
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
	KCDSectorSizeCDDA       KCDSectorSize = 0
	KCDSectorSizeMode1      KCDSectorSize = 0
	KCDSectorSizeMode2      KCDSectorSize = 0
	KCDSectorSizeMode2Form1 KCDSectorSize = 0
	KCDSectorSizeMode2Form2 KCDSectorSize = 0
	KCDSectorSizeWhole      KCDSectorSize = 0
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
	KCDSectorTypeCDDA       KCDSectorType = 0
	KCDSectorTypeCount      KCDSectorType = 0
	KCDSectorTypeMode1      KCDSectorType = 0
	KCDSectorTypeMode2      KCDSectorType = 0
	KCDSectorTypeMode2Form1 KCDSectorType = 0
	KCDSectorTypeMode2Form2 KCDSectorType = 0
	KCDSectorTypeUnknown    KCDSectorType = 0
)

func (e KCDSectorType) String() string {
	switch e {
	case KCDSectorTypeCDDA:
		return "KCDSectorTypeCDDA"
	default:
		return fmt.Sprintf("KCDSectorType(%d)", e)
	}
}

type KCDTOCFormatATI uint

const (
	KCDTOCFormatATIP KCDTOCFormatATI = 0
)

func (e KCDTOCFormatATI) String() string {
	switch e {
	case KCDTOCFormatATIP:
		return "KCDTOCFormatATIP"
	default:
		return fmt.Sprintf("KCDTOCFormatATI(%d)", e)
	}
}

type KCDTrackInfoAddressType uint

const (
	KCDTrackInfoAddressTypeLBA           KCDTrackInfoAddressType = 0
	KCDTrackInfoAddressTypeSessionNumber KCDTrackInfoAddressType = 0
	KCDTrackInfoAddressTypeTrackNumber   KCDTrackInfoAddressType = 0
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
	KCKindNone      KCKind = 0
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
	KCSRStateOff          KCSRState = 0
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
	KChipSetTypeCore2001     KChipSetType = 0
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

type KClamshellSleep uint

const (
	KClamshellSleepBit KClamshellSleep = 0
)

func (e KClamshellSleep) String() string {
	switch e {
	case KClamshellSleepBit:
		return "KClamshellSleepBit"
	default:
		return fmt.Sprintf("KClamshellSleep(%d)", e)
	}
}

type KClearDevice uint

const (
	KClearDeviceFeature KClearDevice = 0
)

func (e KClearDevice) String() string {
	switch e {
	case KClearDeviceFeature:
		return "KClearDeviceFeature"
	default:
		return fmt.Sprintf("KClearDevice(%d)", e)
	}
}

type KClearHub uint

const (
	KClearHubFeature KClearHub = 0
)

func (e KClearHub) String() string {
	switch e {
	case KClearHubFeature:
		return "KClearHubFeature"
	default:
		return fmt.Sprintf("KClearHub(%d)", e)
	}
}

type KConfig uint

const (
	KConfigBusDependentInfoKey KConfig = 0
	KConfigBusInfoBlockLength  KConfig = 0
	KConfigEntryKeyTypePhase   KConfig = 0
	KConfigNodeUniqueIdKey     KConfig = 0
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
	KConfigSBP2LUN      KConfigSBP2 = 0
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
	KConfigUnitSWVersIIDC100     KConfigUnit = 0
	KConfigUnitSWVersIIDC101     KConfigUnit = 0
	KConfigUnitSWVersIIDC102     KConfigUnit = 0
	KConfigUnitSWVersMacintosh10 KConfigUnit = 0
	KConfigUnitSpec1394TA1       KConfigUnit = 0
	KConfigUnitSpecAppleA27      KConfigUnit = 0
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
	KConnectionBlueGammaScale       KConnection = 0
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
	KConvolvedMask  KConvolved = 0
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
	KCoprocessorVersion1    KCoprocessor = 0
	KCoprocessorVersion2    KCoprocessor = 0
	KCoprocessorVersionNone KCoprocessor = 0
)

func (e KCoprocessor) String() string {
	switch e {
	case KCoprocessorVersion1:
		return "KCoprocessorVersion1"
	default:
		return fmt.Sprintf("KCoprocessor(%d)", e)
	}
}

type KDCLCallProc uint

const (
	KDCLCallProcOp KDCLCallProc = 0
)

func (e KDCLCallProc) String() string {
	switch e {
	case KDCLCallProcOp:
		return "KDCLCallProcOp"
	default:
		return fmt.Sprintf("KDCLCallProc(%d)", e)
	}
}

type KDDCBlock uint

const (
	KDDCBlockSize KDDCBlock = 0
)

func (e KDDCBlock) String() string {
	switch e {
	case KDDCBlockSize:
		return "KDDCBlockSize"
	default:
		return fmt.Sprintf("KDDCBlock(%d)", e)
	}
}

type KDDCBlockTypeEDI uint

const (
	KDDCBlockTypeEDID KDDCBlockTypeEDI = 0
)

func (e KDDCBlockTypeEDI) String() string {
	switch e {
	case KDDCBlockTypeEDID:
		return "KDDCBlockTypeEDID"
	default:
		return fmt.Sprintf("KDDCBlockTypeEDI(%d)", e)
	}
}

type KDDCForceRead uint

const (
	KDDCForceReadBit KDDCForceRead = 0
)

func (e KDDCForceRead) String() string {
	switch e {
	case KDDCForceReadBit:
		return "KDDCForceReadBit"
	default:
		return fmt.Sprintf("KDDCForceRead(%d)", e)
	}
}

type KDMSMode uint

const (
	KDMSModeFree     KDMSMode = 0
	KDMSModeNotReady KDMSMode = 0
	KDMSModeReady    KDMSMode = 0
)

func (e KDMSMode) String() string {
	switch e {
	case KDMSModeFree:
		return "KDMSModeFree"
	default:
		return fmt.Sprintf("KDMSMode(%d)", e)
	}
}

type KDPMSSync uint

const (
	KDPMSSyncOff KDPMSSync = 0
)

func (e KDPMSSync) String() string {
	switch e {
	case KDPMSSyncOff:
		return "KDPMSSyncOff"
	default:
		return fmt.Sprintf("KDPMSSync(%d)", e)
	}
}

type KDTMaxEntryName uint

const (
	KDTMaxEntryNameLength KDTMaxEntryName = 0
)

func (e KDTMaxEntryName) String() string {
	switch e {
	case KDTMaxEntryNameLength:
		return "KDTMaxEntryNameLength"
	default:
		return fmt.Sprintf("KDTMaxEntryName(%d)", e)
	}
}

type KDTMaxPropertyName uint

const (
	KDTMaxPropertyNameLength KDTMaxPropertyName = 0
)

func (e KDTMaxPropertyName) String() string {
	switch e {
	case KDTMaxPropertyNameLength:
		return "KDTMaxPropertyNameLength"
	default:
		return fmt.Sprintf("KDTMaxPropertyName(%d)", e)
	}
}

type KDTPathName uint

const (
	KDTPathNameSeparator KDTPathName = 0
)

func (e KDTPathName) String() string {
	switch e {
	case KDTPathNameSeparator:
		return "KDTPathNameSeparator"
	default:
		return fmt.Sprintf("KDTPathName(%d)", e)
	}
}

type KDVDBookType uint

const (
	KDVDBookTypeHDR               KDVDBookType = 0
	KDVDBookTypeHDRAM             KDVDBookType = 0
	KDVDBookTypeHDROM             KDVDBookType = 0
	KDVDBookTypeHDRW              KDVDBookType = 0
	KDVDBookTypePlusR             KDVDBookType = 0
	KDVDBookTypePlusRDoubleLayer  KDVDBookType = 0
	KDVDBookTypePlusRW            KDVDBookType = 0
	KDVDBookTypePlusRWDoubleLayer KDVDBookType = 0
	KDVDBookTypeR                 KDVDBookType = 0
	KDVDBookTypeRAM               KDVDBookType = 0
	KDVDBookTypeROM               KDVDBookType = 0
	KDVDBookTypeRW                KDVDBookType = 0
)

func (e KDVDBookType) String() string {
	switch e {
	case KDVDBookTypeHDR:
		return "KDVDBookTypeHDR"
	default:
		return fmt.Sprintf("KDVDBookType(%d)", e)
	}
}

type KDVDCPRM uint

const (
	KDVDCPRMRegion1 KDVDCPRM = 0
	KDVDCPRMRegion2 KDVDCPRM = 0
	KDVDCPRMRegion3 KDVDCPRM = 0
	KDVDCPRMRegion4 KDVDCPRM = 0
	KDVDCPRMRegion5 KDVDCPRM = 0
	KDVDCPRMRegion6 KDVDCPRM = 0
)

func (e KDVDCPRM) String() string {
	switch e {
	case KDVDCPRMRegion1:
		return "KDVDCPRMRegion1"
	default:
		return fmt.Sprintf("KDVDCPRM(%d)", e)
	}
}

type KDVDFeatures uint

const (
	KDVDFeaturesBUFWriteBit         KDVDFeatures = 0
	KDVDFeaturesBUFWriteMask        KDVDFeatures = 0
	KDVDFeaturesCSSMask             KDVDFeatures = 0
	KDVDFeaturesHDRAMMask           KDVDFeatures = 0
	KDVDFeaturesHDRBit              KDVDFeatures = 0
	KDVDFeaturesHDRMask             KDVDFeatures = 0
	KDVDFeaturesHDRWMask            KDVDFeatures = 0
	KDVDFeaturesHDReadMask          KDVDFeatures = 0
	KDVDFeaturesPlusRMask           KDVDFeatures = 0
	KDVDFeaturesPlusRWMask          KDVDFeatures = 0
	KDVDFeaturesRandomWriteableMask KDVDFeatures = 0
	KDVDFeaturesReWriteableMask     KDVDFeatures = 0
	KDVDFeaturesReadStructuresMask  KDVDFeatures = 0
	KDVDFeaturesTestWriteMask       KDVDFeatures = 0
	KDVDFeaturesWriteOnceMask       KDVDFeatures = 0
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
	KDVDKeyClassRSSA          KDVDKeyClass = 0
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
	KDVDKeyFormatAGID_CPRM       KDVDKeyFormatAGID = 0
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
	KDVDMediaTypeRW  KDVDMediaType = 0
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
	KDVDRZoneInfoAddressTypeLBA          KDVDRZoneInfoAddressType = 0
	KDVDRZoneInfoAddressTypeRZoneNumber  KDVDRZoneInfoAddressType = 0
)

func (e KDVDRZoneInfoAddressType) String() string {
	switch e {
	case KDVDRZoneInfoAddressTypeBorderNumber:
		return "KDVDRZoneInfoAddressTypeBorderNumber"
	default:
		return fmt.Sprintf("KDVDRZoneInfoAddressType(%d)", e)
	}
}

type KDVDRegionalPlaybackControlScheme uint

const (
	KDVDRegionalPlaybackControlSchemePhase1 KDVDRegionalPlaybackControlScheme = 0
)

func (e KDVDRegionalPlaybackControlScheme) String() string {
	switch e {
	case KDVDRegionalPlaybackControlSchemePhase1:
		return "KDVDRegionalPlaybackControlSchemePhase1"
	default:
		return fmt.Sprintf("KDVDRegionalPlaybackControlScheme(%d)", e)
	}
}

type KDVDStructureFormatCopyright uint

const (
	KDVDStructureFormatCopyrightInfo KDVDStructureFormatCopyright = 0
)

func (e KDVDStructureFormatCopyright) String() string {
	switch e {
	case KDVDStructureFormatCopyrightInfo:
		return "KDVDStructureFormatCopyrightInfo"
	default:
		return fmt.Sprintf("KDVDStructureFormatCopyright(%d)", e)
	}
}

type KDVIPowerSwitch uint

const (
	KDVIPowerSwitchActiveMask  KDVIPowerSwitch = 0
	KDVIPowerSwitchFeature     KDVIPowerSwitch = 0
	KDVIPowerSwitchSupportMask KDVIPowerSwitch = 0
)

func (e KDVIPowerSwitch) String() string {
	switch e {
	case KDVIPowerSwitchActiveMask:
		return "KDVIPowerSwitchActiveMask"
	default:
		return fmt.Sprintf("KDVIPowerSwitch(%d)", e)
	}
}

type KData uint

const (
	KDataRotate0 KData = 0
)

func (e KData) String() string {
	switch e {
	case KDataRotate0:
		return "KDataRotate0"
	default:
		return fmt.Sprintf("KData(%d)", e)
	}
}

type KDebugType uint

const (
	KDebugTypeDisplay KDebugType = 0
	KDebugTypeNone    KDebugType = 0
	KDebugTypeSerial  KDebugType = 0
)

func (e KDebugType) String() string {
	switch e {
	case KDebugTypeDisplay:
		return "KDebugTypeDisplay"
	default:
		return fmt.Sprintf("KDebugType(%d)", e)
	}
}

type KDeclROMtables uint

const (
	KDeclROMtablesValue   KDeclROMtables = 0
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
	KDepthDependent KDepth = 0
	KDepthMode1     KDepth = 0
	KDepthMode4     KDepth = 0
)

func (e KDepth) String() string {
	switch e {
	case KDepthDependent:
		return "KDepthDependent"
	default:
		return fmt.Sprintf("KDepth(%d)", e)
	}
}

type KDigitalSignalBit uint

const (
	KAnalogSetupExpectedBit KDigitalSignalBit = 0
	KDigitalSignalBitValue  KDigitalSignalBit = 0
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
	KDisableVerticalSyncBit   KDisable = 0
)

func (e KDisable) String() string {
	switch e {
	case KDisableHorizontalSyncBit:
		return "KDisableHorizontalSyncBit"
	default:
		return fmt.Sprintf("KDisable(%d)", e)
	}
}

type KDisabledInterruptState uint

const (
	KDisabledInterruptStateValue KDisabledInterruptState = 0
	KEnabledInterruptState       KDisabledInterruptState = 0
)

func (e KDisabledInterruptState) String() string {
	switch e {
	case KDisabledInterruptStateValue:
		return "KDisabledInterruptStateValue"
	default:
		return fmt.Sprintf("KDisabledInterruptState(%d)", e)
	}
}

type KDiscStatus int

const (
	KDiscStatusComplete     KDiscStatus = 0
	KDiscStatusEmpty        KDiscStatus = 0
	KDiscStatusErasableMask KDiscStatus = 0
	KDiscStatusIncomplete   KDiscStatus = 0
	KDiscStatusMask         KDiscStatus = 0
	KDiscStatusOther        KDiscStatus = 0
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
	KDisplayModeAlwaysShowFlag        KDisplayMode = 0
	KDisplayModeDefaultFlag           KDisplayMode = 0
	KDisplayModeSafeFlag              KDisplayMode = 0
	KDisplayModeValidFlag             KDisplayMode = 0
)

func (e KDisplayMode) String() string {
	switch e {
	case KDisplayModeAcceleratorBackedFlag:
		return "KDisplayModeAcceleratorBackedFlag"
	default:
		return fmt.Sprintf("KDisplayMode(%d)", e)
	}
}

type KDisplayModeIDBoot uint

const (
	KDisplayModeIDBootProgrammable KDisplayModeIDBoot = 0
)

func (e KDisplayModeIDBoot) String() string {
	switch e {
	case KDisplayModeIDBootProgrammable:
		return "KDisplayModeIDBootProgrammable"
	default:
		return fmt.Sprintf("KDisplayModeIDBoot(%d)", e)
	}
}

type KDisplayProductID uint

const (
	KDisplayProductIDGeneric KDisplayProductID = 0
)

func (e KDisplayProductID) String() string {
	switch e {
	case KDisplayProductIDGeneric:
		return "KDisplayProductIDGeneric"
	default:
		return fmt.Sprintf("KDisplayProductID(%d)", e)
	}
}

type KDisplaySubPixel uint

const (
	KDisplaySubPixelConfigurationDelta KDisplaySubPixel = 0
	KDisplaySubPixelShapeRound         KDisplaySubPixel = 0
)

func (e KDisplaySubPixel) String() string {
	switch e {
	case KDisplaySubPixelConfigurationDelta:
		return "KDisplaySubPixelConfigurationDelta"
	default:
		return fmt.Sprintf("KDisplaySubPixel(%d)", e)
	}
}

type KDriver uint

const (
	KDriverIsConcurrent              KDriver = 0
	KDriverIsForVirtualDevice        KDriver = 0
	KDriverIsLoadedAtBoot            KDriver = 0
	KDriverIsLoadedUponDiscovery     KDriver = 0
	KDriverIsOpenedUponLoad          KDriver = 0
	KDriverIsUnderExpertControl      KDriver = 0
	KDriverQueuesIOPB                KDriver = 0
	KDriverSupportDMSuspendAndResume KDriver = 0
)

func (e KDriver) String() string {
	switch e {
	case KDriverIsConcurrent:
		return "KDriverIsConcurrent"
	default:
		return fmt.Sprintf("KDriver(%d)", e)
	}
}

type KESC uint

const (
	KESCOnePortraitMono KESC = 0
	KESCSixStandard     KESC = 0
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
	KEfiBootServicesCode  KEfi = 0
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
	KEndpointDirectionIn      KEndpointDirection = 0
	KEndpointDirectionOut     KEndpointDirection = 0
	KEndpointDirectionUnknown KEndpointDirection = 0
)

func (e KEndpointDirection) String() string {
	switch e {
	case KEndpointDirectionIn:
		return "KEndpointDirectionIn"
	default:
		return fmt.Sprintf("KEndpointDirection(%d)", e)
	}
}

type KEndpointSynchronizationType uint

const (
	KEndpointSynchronizationTypeAdaptive     KEndpointSynchronizationType = 0
	KEndpointSynchronizationTypeAsynchronous KEndpointSynchronizationType = 0
	KEndpointSynchronizationTypeNone         KEndpointSynchronizationType = 0
	KEndpointSynchronizationTypeSynchronous  KEndpointSynchronizationType = 0
)

func (e KEndpointSynchronizationType) String() string {
	switch e {
	case KEndpointSynchronizationTypeAdaptive:
		return "KEndpointSynchronizationTypeAdaptive"
	default:
		return fmt.Sprintf("KEndpointSynchronizationType(%d)", e)
	}
}

type KEndpointType uint

const (
	KEndpointTypeBulk        KEndpointType = 0
	KEndpointTypeControl     KEndpointType = 0
	KEndpointTypeInterrupt   KEndpointType = 0
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

type KEndpointUsageTypeInterrupt uint

const (
	KEndpointUsageTypeInterruptNotification KEndpointUsageTypeInterrupt = 0
)

func (e KEndpointUsageTypeInterrupt) String() string {
	switch e {
	case KEndpointUsageTypeInterruptNotification:
		return "KEndpointUsageTypeInterruptNotification"
	default:
		return fmt.Sprintf("KEndpointUsageTypeInterrupt(%d)", e)
	}
}

type KError int

const (
	KErrorValue    KError = 0
	KIterationDone KError = 0
	KSuccess       KError = 0
)

func (e KError) String() string {
	switch e {
	case KErrorValue:
		return "KErrorValue"
	default:
		return fmt.Sprintf("KError(%d)", e)
	}
}

type KExclave uint

const (
	KExclaveRPCActive        KExclave = 0
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
	KExclaveIpcStackEntryHaveStack        KExclaveIpcStackEntryHave = 0
)

func (e KExclaveIpcStackEntryHave) String() string {
	switch e {
	case KExclaveIpcStackEntryHaveInvocationID:
		return "KExclaveIpcStackEntryHaveInvocationID"
	default:
		return fmt.Sprintf("KExclaveIpcStackEntryHave(%d)", e)
	}
}

type KExclaveScresultHaveIPC uint

const (
	KExclaveScresultHaveIPCStack KExclaveScresultHaveIPC = 0
)

func (e KExclaveScresultHaveIPC) String() string {
	switch e {
	case KExclaveScresultHaveIPCStack:
		return "KExclaveScresultHaveIPCStack"
	default:
		return fmt.Sprintf("KExclaveScresultHaveIPC(%d)", e)
	}
}

type KExclaveTextLayoutLoadAddresses uint

const (
	KExclaveTextLayoutLoadAddressesSynthetic KExclaveTextLayoutLoadAddresses = 0
	KExclaveTextLayoutLoadAddressesUnslid    KExclaveTextLayoutLoadAddresses = 0
)

func (e KExclaveTextLayoutLoadAddresses) String() string {
	switch e {
	case KExclaveTextLayoutLoadAddressesSynthetic:
		return "KExclaveTextLayoutLoadAddressesSynthetic"
	default:
		return fmt.Sprintf("KExclaveTextLayoutLoadAddresses(%d)", e)
	}
}

type KFBDisplayPowerState uint

const (
	KFBDisplayPowerStateMask KFBDisplayPowerState = 0
)

func (e KFBDisplayPowerState) String() string {
	switch e {
	case KFBDisplayPowerStateMask:
		return "KFBDisplayPowerStateMask"
	default:
		return fmt.Sprintf("KFBDisplayPowerState(%d)", e)
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
	KFWAddressBusID                KFW = 0
	KFWAllowMultiMode              KFW = 0
	KFWAlwaysMultiMode             KFW = 0
	KFWBroadcastNodeID             KFW = 0
	KFWDCLInvalidNotification      KFW = 0
	KFWDCLModifyNotification       KFW = 0
	KFWDCLUpdateNotification       KFW = 0
	KFWDefaultIsochResourceFlags   KFW = 0
	KFWNeverMultiMode              KFW = 0
	KFWNuDCLModifyJumpNotification KFW = 0
	KFWNuDCLModifyNotification     KFW = 0
	KFWNuDCLUpdateNotification     KFW = 0
	KFWSpeed100MBit                KFW = 0
	KFWSpeed200MBit                KFW = 0
	KFWSpeed400MBit                KFW = 0
	KFWSpeed800MBit                KFW = 0
	KFWSpeedInvalid                KFW = 0
	KFWSpeedMaximum                KFW = 0
	KFWSpeedReserved               KFW = 0
	KFWSpeedReserved1              KFW = 0
	KFWSpeedUnknownMask            KFW = 0
	KFWSuggestMultiMode            KFW = 0
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
	KFWAVCAsyncPlug1  KFWAVCAsync = 0
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
	KFWAVCStateBusResumed    KFWAVCState = 0
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

type KFWAck uint

const (
	KFWAckBusyA     KFWAck = 0
	KFWAckBusyB     KFWAck = 0
	KFWAckBusyX     KFWAck = 0
	KFWAckComplete  KFWAck = 0
	KFWAckDataError KFWAck = 0
	KFWAckPending   KFWAck = 0
	KFWAckTimeout   KFWAck = 0
	KFWAckTypeError KFWAck = 0
)

func (e KFWAck) String() string {
	switch e {
	case KFWAckBusyA:
		return "KFWAckBusyA"
	default:
		return fmt.Sprintf("KFWAck(%d)", e)
	}
}

type KFWAsynchSpd uint

const (
	KFWAsynchAckSent                    KFWAsynchSpd = 0
	KFWAsynchAckSentPhase               KFWAsynchSpd = 0
	KFWAsynchDataLength                 KFWAsynchSpd = 0
	KFWAsynchDataLengthPhase            KFWAsynchSpd = 0
	KFWAsynchDestinationID              KFWAsynchSpd = 0
	KFWAsynchDestinationIDPhase         KFWAsynchSpd = 0
	KFWAsynchDestinationOffsetHigh      KFWAsynchSpd = 0
	KFWAsynchDestinationOffsetHighPhase KFWAsynchSpd = 0
	KFWAsynchDestinationOffsetLow       KFWAsynchSpd = 0
	KFWAsynchDestinationOffsetLowPhase  KFWAsynchSpd = 0
	KFWAsynchExtendedTCode              KFWAsynchSpd = 0
	KFWAsynchExtendedTCodePhase         KFWAsynchSpd = 0
	KFWAsynchNew                        KFWAsynchSpd = 0
	KFWAsynchPriority                   KFWAsynchSpd = 0
	KFWAsynchPriorityPhase              KFWAsynchSpd = 0
	KFWAsynchRCode                      KFWAsynchSpd = 0
	KFWAsynchRCodePhase                 KFWAsynchSpd = 0
	KFWAsynchRetryA                     KFWAsynchSpd = 0
	KFWAsynchRt                         KFWAsynchSpd = 0
	KFWAsynchRtPhase                    KFWAsynchSpd = 0
	KFWAsynchSourceID                   KFWAsynchSpd = 0
	KFWAsynchSourceIDPhase              KFWAsynchSpd = 0
	KFWAsynchSpdValue                   KFWAsynchSpd = 0
	KFWAsynchSpdPhase                   KFWAsynchSpd = 0
	KFWAsynchTLabel                     KFWAsynchSpd = 0
	KFWAsynchTLabelPhase                KFWAsynchSpd = 0
	KFWAsynchTTotal                     KFWAsynchSpd = 0
	KTIAsycnhRetryB                     KFWAsynchSpd = 0
)

func (e KFWAsynchSpd) String() string {
	switch e {
	case KFWAsynchAckSent:
		return "KFWAsynchAckSent"
	default:
		return fmt.Sprintf("KFWAsynchSpd(%d)", e)
	}
}

type KFWBIB uint

const (
	KFWBIBIrmc        KFWBIB = 0
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

type KFWBusManagerArbitrationTimeout uint

const (
	KFWBusManagerArbitrationTimeoutDuration KFWBusManagerArbitrationTimeout = 0
)

func (e KFWBusManagerArbitrationTimeout) String() string {
	switch e {
	case KFWBusManagerArbitrationTimeoutDuration:
		return "KFWBusManagerArbitrationTimeoutDuration"
	default:
		return fmt.Sprintf("KFWBusManagerArbitrationTimeout(%d)", e)
	}
}

type KFWCSRState uint

const (
	KFWCSRStateCMstr   KFWCSRState = 0
	KFWCSRStateGone    KFWCSRState = 0
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

type KFWConfigurationPacketI uint

const (
	KFWConfigurationPacketID KFWConfigurationPacketI = 0
)

func (e KFWConfigurationPacketI) String() string {
	switch e {
	case KFWConfigurationPacketID:
		return "KFWConfigurationPacketID"
	default:
		return fmt.Sprintf("KFWConfigurationPacketI(%d)", e)
	}
}

type KFWDCL uint

const (
	KFWDCLCycleEvent     KFWDCL = 0
	KFWDCLImmediateEvent KFWDCL = 0
	KFWDCLSyBitsEvent    KFWDCL = 0
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
	KFWDCLOpDynamicFlag       KFWDCLOp = 0
	KFWDCLOpFlagMask          KFWDCLOp = 0
	KFWDCLOpFlagPhase         KFWDCLOp = 0
	KFWDCLOpVendorDefinedFlag KFWDCLOp = 0
)

func (e KFWDCLOp) String() string {
	switch e {
	case KFWDCLOpDynamicFlag:
		return "KFWDCLOpDynamicFlag"
	default:
		return fmt.Sprintf("KFWDCLOp(%d)", e)
	}
}

type KFWDebugIgnoreNode uint

const (
	KFWDebugIgnoreNodeNone KFWDebugIgnoreNode = 0
)

func (e KFWDebugIgnoreNode) String() string {
	switch e {
	case KFWDebugIgnoreNodeNone:
		return "KFWDebugIgnoreNodeNone"
	default:
		return fmt.Sprintf("KFWDebugIgnoreNode(%d)", e)
	}
}

type KFWExtendedTCodeBounded uint

const (
	KFWExtendedTCodeBoundedAdd KFWExtendedTCodeBounded = 0
)

func (e KFWExtendedTCodeBounded) String() string {
	switch e {
	case KFWExtendedTCodeBoundedAdd:
		return "KFWExtendedTCodeBoundedAdd"
	default:
		return fmt.Sprintf("KFWExtendedTCodeBounded(%d)", e)
	}
}

type KFWIsoch uint

const (
	KFWIsochBigEndianUpdates            KFWIsoch = 0
	KFWIsochChanNum                     KFWIsoch = 0
	KFWIsochChanNumPhase                KFWIsoch = 0
	KFWIsochDataLength                  KFWIsoch = 0
	KFWIsochDataLengthPhase             KFWIsoch = 0
	KFWIsochEnableRobustness            KFWIsoch = 0
	KFWIsochPortDefaultOptions          KFWIsoch = 0
	KFWIsochPortUseSeparateKernelThread KFWIsoch = 0
	KFWIsochRequireLastContext          KFWIsoch = 0
	KFWIsochSy                          KFWIsoch = 0
	KFWIsochSyPhase                     KFWIsoch = 0
	KFWIsochTCode                       KFWIsoch = 0
	KFWIsochTCodePhase                  KFWIsoch = 0
	KFWIsochTag                         KFWIsoch = 0
	KFWIsochTagPhase                    KFWIsoch = 0
)

func (e KFWIsoch) String() string {
	switch e {
	case KFWIsochBigEndianUpdates:
		return "KFWIsochBigEndianUpdates"
	default:
		return fmt.Sprintf("KFWIsoch(%d)", e)
	}
}

type KFWIsochChannel uint

const (
	KFWIsochChannelDefaultFlags      KFWIsochChannel = 0
	KFWIsochChannelDoNotResumeOnWake KFWIsochChannel = 0
)

func (e KFWIsochChannel) String() string {
	switch e {
	case KFWIsochChannelDefaultFlags:
		return "KFWIsochChannelDefaultFlags"
	default:
		return fmt.Sprintf("KFWIsochChannel(%d)", e)
	}
}

type KFWIsochChannelChannelNot uint

const (
	KFWIsochChannelChannelNotAvailable KFWIsochChannelChannelNot = 0
)

func (e KFWIsochChannelChannelNot) String() string {
	switch e {
	case KFWIsochChannelChannelNotAvailable:
		return "KFWIsochChannelChannelNotAvailable"
	default:
		return fmt.Sprintf("KFWIsochChannelChannelNot(%d)", e)
	}
}

type KFWMax uint

const (
	KFWMaxBusses KFWMax = 0
)

func (e KFWMax) String() string {
	switch e {
	case KFWMaxBusses:
		return "KFWMaxBusses"
	default:
		return fmt.Sprintf("KFWMax(%d)", e)
	}
}

type KFWPacketT uint

const (
	KFWPacketTCode      KFWPacketT = 0
	KFWPacketTCodePhase KFWPacketT = 0
)

func (e KFWPacketT) String() string {
	switch e {
	case KFWPacketTCode:
		return "KFWPacketTCode"
	default:
		return fmt.Sprintf("KFWPacketT(%d)", e)
	}
}

type KFWPhyConfigurationGap uint

const (
	KFWPhyConfigurationGapCnt KFWPhyConfigurationGap = 0
)

func (e KFWPhyConfigurationGap) String() string {
	switch e {
	case KFWPhyConfigurationGapCnt:
		return "KFWPhyConfigurationGapCnt"
	default:
		return fmt.Sprintf("KFWPhyConfigurationGap(%d)", e)
	}
}

type KFWPhyPacket uint

const (
	KFWPhyPacketID         KFWPhyPacket = 0
	KFWPhyPacketIDPhase    KFWPhyPacket = 0
	KFWPhyPacketPhyID      KFWPhyPacket = 0
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
	KFWResponseTypeError    KFWResponse = 0
)

func (e KFWResponse) String() string {
	switch e {
	case KFWResponseAddressError:
		return "KFWResponseAddressError"
	default:
		return fmt.Sprintf("KFWResponse(%d)", e)
	}
}

type KFWSBP2 uint

const (
	KFWSBP2NormalCommandReset   KFWSBP2 = 0
	KFWSBP2NormalCommandStatus  KFWSBP2 = 0
	KFWSBP2NormalCommandTimeout KFWSBP2 = 0
	KFWSBP2UnsolicitedStatus    KFWSBP2 = 0
)

func (e KFWSBP2) String() string {
	switch e {
	case KFWSBP2NormalCommandReset:
		return "KFWSBP2NormalCommandReset"
	default:
		return fmt.Sprintf("KFWSBP2(%d)", e)
	}
}

type KFWSBP2Abort uint

const (
	KFWSBP2AbortTask KFWSBP2Abort = 0
)

func (e KFWSBP2Abort) String() string {
	switch e {
	case KFWSBP2AbortTask:
		return "KFWSBP2AbortTask"
	default:
		return fmt.Sprintf("KFWSBP2Abort(%d)", e)
	}
}

type KFWSBP2Command uint

const (
	KFWSBP2CommandCheckGeneration        KFWSBP2Command = 0
	KFWSBP2CommandCompleteNotify         KFWSBP2Command = 0
	KFWSBP2CommandDummyORB               KFWSBP2Command = 0
	KFWSBP2CommandFixedSize              KFWSBP2Command = 0
	KFWSBP2CommandImmediate              KFWSBP2Command = 0
	KFWSBP2CommandNormalORB              KFWSBP2Command = 0
	KFWSBP2CommandReservedORB            KFWSBP2Command = 0
	KFWSBP2CommandTransferDataFromTarget KFWSBP2Command = 0
	KFWSBP2CommandVendorORB              KFWSBP2Command = 0
	KFWSBP2CommandVirtualORBs            KFWSBP2Command = 0
)

func (e KFWSBP2Command) String() string {
	switch e {
	case KFWSBP2CommandCheckGeneration:
		return "KFWSBP2CommandCheckGeneration"
	default:
		return fmt.Sprintf("KFWSBP2Command(%d)", e)
	}
}

type KFWSBP2ConstraintForceDouble uint

const (
	KFWSBP2ConstraintForceDoubleBuffer KFWSBP2ConstraintForceDouble = 0
)

func (e KFWSBP2ConstraintForceDouble) String() string {
	switch e {
	case KFWSBP2ConstraintForceDoubleBuffer:
		return "KFWSBP2ConstraintForceDoubleBuffer"
	default:
		return fmt.Sprintf("KFWSBP2ConstraintForceDouble(%d)", e)
	}
}

type KFWSBP2DontSynchronizeMgmt uint

const (
	KFWSBP2DontSynchronizeMgmtAgent KFWSBP2DontSynchronizeMgmt = 0
)

func (e KFWSBP2DontSynchronizeMgmt) String() string {
	switch e {
	case KFWSBP2DontSynchronizeMgmtAgent:
		return "KFWSBP2DontSynchronizeMgmtAgent"
	default:
		return fmt.Sprintf("KFWSBP2DontSynchronizeMgmt(%d)", e)
	}
}

type KFWSBP2MaxPageCluster uint

const (
	KFWSBP2MaxPageClusterSize KFWSBP2MaxPageCluster = 0
)

func (e KFWSBP2MaxPageCluster) String() string {
	switch e {
	case KFWSBP2MaxPageClusterSize:
		return "KFWSBP2MaxPageClusterSize"
	default:
		return fmt.Sprintf("KFWSBP2MaxPageCluster(%d)", e)
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
	KFWTCodeCycleStart          KFWTCode = 0
	KFWTCodeIsochronousBlock    KFWTCode = 0
	KFWTCodeLock                KFWTCode = 0
	KFWTCodeLockResponse        KFWTCode = 0
	KFWTCodePHYPacket           KFWTCode = 0
	KFWTCodeReadBlock           KFWTCode = 0
	KFWTCodeReadBlockResponse   KFWTCode = 0
	KFWTCodeReadQuadlet         KFWTCode = 0
	KFWTCodeReadQuadletResponse KFWTCode = 0
	KFWTCodeWriteBlock          KFWTCode = 0
	KFWTCodeWriteQuadlet        KFWTCode = 0
	KFWTCodeWriteResponse       KFWTCode = 0
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
	KFifthDepthMode      KFirstDepthMode = 0
	KFirstDepthModeValue KFirstDepthMode = 0
	KFourthDepthMode     KFirstDepthMode = 0
	KSecondDepthMode     KFirstDepthMode = 0
	KSixthDepthMode      KFirstDepthMode = 0
	KThirdDepthMode      KFirstDepthMode = 0
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
	KLastIOKitNotificationType       KFirstIOKitNotificationType = 0
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
	KDMABit              KFixedDeviceBit = 0
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

type KFortyFiveSecondTimeoutInM uint

const (
	KFortyFiveSecondTimeoutInMS KFortyFiveSecondTimeoutInM = 0
)

func (e KFortyFiveSecondTimeoutInM) String() string {
	switch e {
	case KFortyFiveSecondTimeoutInMS:
		return "KFortyFiveSecondTimeoutInMS"
	default:
		return fmt.Sprintf("KFortyFiveSecondTimeoutInM(%d)", e)
	}
}

type KFramebuffer uint

const (
	KFramebufferDisableAltivecAccess    KFramebuffer = 0
	KFramebufferSupportsCopybackCache   KFramebuffer = 0
	KFramebufferSupportsGammaCorrection KFramebuffer = 0
	KFramebufferSupportsWritethruCache  KFramebuffer = 0
)

func (e KFramebuffer) String() string {
	switch e {
	case KFramebufferDisableAltivecAccess:
		return "KFramebufferDisableAltivecAccess"
	default:
		return fmt.Sprintf("KFramebuffer(%d)", e)
	}
}

type KGammaTableIDFind uint

const (
	KGammaTableIDFindFirst KGammaTableIDFind = 0
)

func (e KGammaTableIDFind) String() string {
	switch e {
	case KGammaTableIDFindFirst:
		return "KGammaTableIDFindFirst"
	default:
		return fmt.Sprintf("KGammaTableIDFind(%d)", e)
	}
}

type KHFS uint

const (
	KHFSAllocationFileID        KHFS = 0
	KHFSAttributeDataFileID     KHFS = 0
	KHFSAttributesFileID        KHFS = 0
	KHFSAutoCandidateMask       KHFS = 0
	KHFSBadBlockFileID          KHFS = 0
	KHFSBinaryCompare           KHFS = 0
	KHFSBogusExtentFileID       KHFS = 0
	KHFSCaseFolding             KHFS = 0
	KHFSCatalogFileID           KHFS = 0
	KHFSCatalogKeyMaximumLength KHFS = 0
	KHFSExtentDensity           KHFS = 0
	KHFSExtentsFileID           KHFS = 0
	KHFSFileRecord              KHFS = 0
	KHFSFileThreadRecord        KHFS = 0
	KHFSFirstUserCatalogNodeID  KHFS = 0
	KHFSFolderRecord            KHFS = 0
	KHFSFolderThreadRecord      KHFS = 0
	KHFSHasAttributesMask       KHFS = 0
	KHFSPlusAttrMinNodeSize     KHFS = 0
	KHFSPlusExtentDensity       KHFS = 0
	KHFSPlusFileRecord          KHFS = 0
	KHFSPlusFileThreadRecord    KHFS = 0
	KHFSPlusFolderRecord        KHFS = 0
	KHFSPlusFolderThreadRecord  KHFS = 0
	KHFSRepairCatalogFileID     KHFS = 0
	KHFSRootFolderID            KHFS = 0
	KHFSRootParentID            KHFS = 0
	KHFSStartupFileID           KHFS = 0
	KHFSUnusedNodeFixMask       KHFS = 0
	KHFSVolumeJournaledMask     KHFS = 0
)

func (e KHFS) String() string {
	switch e {
	case KHFSAllocationFileID:
		return "KHFSAllocationFileID"
	default:
		return fmt.Sprintf("KHFS(%d)", e)
	}
}

type KHFSMax uint

const (
	KHFSMaxFileNameChars   KHFSMax = 0
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

type KHFSMaxAttrName uint

const (
	KHFSMaxAttrNameLen KHFSMaxAttrName = 0
)

func (e KHFSMaxAttrName) String() string {
	switch e {
	case KHFSMaxAttrNameLen:
		return "KHFSMaxAttrNameLen"
	default:
		return fmt.Sprintf("KHFSMaxAttrName(%d)", e)
	}
}

type KHFSPlusAttr uint

const (
	KHFSPlusAttrExtents  KHFSPlusAttr = 0
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

type KHFSUnusedNodesFix uint

const (
	KHFSUnusedNodesFixDate         KHFSUnusedNodesFix = 0
	KHFSUnusedNodesFixExpandedDate KHFSUnusedNodesFix = 0
)

func (e KHFSUnusedNodesFix) String() string {
	switch e {
	case KHFSUnusedNodesFixDate:
		return "KHFSUnusedNodesFixDate"
	default:
		return fmt.Sprintf("KHFSUnusedNodesFix(%d)", e)
	}
}

type KHI uint

const (
	KHIKeyboardDevice         KHI = 0
	KHIRelativePointingDevice KHI = 0
	KHIUnknownDevice          KHI = 0
)

func (e KHI) String() string {
	switch e {
	case KHIKeyboardDevice:
		return "KHIKeyboardDevice"
	default:
		return fmt.Sprintf("KHI(%d)", e)
	}
}

type KHID uint

const (
	KHIDBootProtocolValue   KHID = 0
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
	KHIDDispatchOptionPhaseCanceled             KHIDDispatchOption = 0
)

func (e KHIDDispatchOption) String() string {
	switch e {
	case KHIDDispatchOptionDeliveryNotificationForce:
		return "KHIDDispatchOptionDeliveryNotificationForce"
	default:
		return fmt.Sprintf("KHIDDispatchOption(%d)", e)
	}
}

type KHIDDispatchOptionKeyboardNo uint

const (
	KHIDDispatchOptionKeyboardNoRepeat KHIDDispatchOptionKeyboardNo = 0
)

func (e KHIDDispatchOptionKeyboardNo) String() string {
	switch e {
	case KHIDDispatchOptionKeyboardNoRepeat:
		return "KHIDDispatchOptionKeyboardNoRepeat"
	default:
		return fmt.Sprintf("KHIDDispatchOptionKeyboardNo(%d)", e)
	}
}

type KHIDDispatchOptionPointer uint

const (
	KHIDDispatchOptionPointerAbsolutToRelative KHIDDispatchOptionPointer = 0
	KHIDDispatchOptionPointerAffixToScreen     KHIDDispatchOptionPointer = 0
	KHIDDispatchOptionPointerDisplayIntegrated KHIDDispatchOptionPointer = 0
	KHIDDispatchOptionPointerNoAcceleration    KHIDDispatchOptionPointer = 0
)

func (e KHIDDispatchOptionPointer) String() string {
	switch e {
	case KHIDDispatchOptionPointerAbsolutToRelative:
		return "KHIDDispatchOptionPointerAbsolutToRelative"
	default:
		return fmt.Sprintf("KHIDDispatchOptionPointer(%d)", e)
	}
}

type KHIDPage uint

const (
	KHIDPage_Arcade    KHIDPage = 0
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

type KHIDRq uint

const (
	KHIDRqGetIdle     KHIDRq = 0
	KHIDRqGetProtocol KHIDRq = 0
	KHIDRqGetReport   KHIDRq = 0
	KHIDRqSetIdle     KHIDRq = 0
	KHIDRqSetProtocol KHIDRq = 0
	KHIDRqSetReport   KHIDRq = 0
)

func (e KHIDRq) String() string {
	switch e {
	case KHIDRqGetIdle:
		return "KHIDRqGetIdle"
	default:
		return fmt.Sprintf("KHIDRq(%d)", e)
	}
}

type KHIDRt uint

const (
	KHIDRtFeatureReport KHIDRt = 0
	KHIDRtInputReport   KHIDRt = 0
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
	KHIDUsage_Undefined KHIDUsage = 0
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
	KHFSPlusCreator        KHardLinkFileType = 0
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

type KHardwareCursorDescriptorMajor uint

const (
	KHardwareCursorDescriptorMajorVersion KHardwareCursorDescriptorMajor = 0
)

func (e KHardwareCursorDescriptorMajor) String() string {
	switch e {
	case KHardwareCursorDescriptorMajorVersion:
		return "KHardwareCursorDescriptorMajorVersion"
	default:
		return fmt.Sprintf("KHardwareCursorDescriptorMajor(%d)", e)
	}
}

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

type KHub uint

const (
	KHubPortSetPowerOff             KHub = 0
	KHubPortSetPowerOn              KHub = 0
	KHubSupportsGangPower           KHub = 0
	KHubSupportsIndividualPortPower KHub = 0
)

func (e KHub) String() string {
	switch e {
	case KHubPortSetPowerOff:
		return "KHubPortSetPowerOff"
	default:
		return fmt.Sprintf("KHub(%d)", e)
	}
}

type KHubLocalPower uint

const (
	KHubLocalPowerStatus KHubLocalPower = 0
)

func (e KHubLocalPower) String() string {
	switch e {
	case KHubLocalPowerStatus:
		return "KHubLocalPowerStatus"
	default:
		return fmt.Sprintf("KHubLocalPower(%d)", e)
	}
}

type KHubPortIndicator uint

const (
	KHubPortIndicatorAmber KHubPortIndicator = 0
)

func (e KHubPortIndicator) String() string {
	switch e {
	case KHubPortIndicatorAmber:
		return "KHubPortIndicatorAmber"
	default:
		return fmt.Sprintf("KHubPortIndicator(%d)", e)
	}
}

type KHubSuperSpeedProtocol uint

const (
	KHIDKeyboardInterfaceProtocol              KHubSuperSpeedProtocol = 0
	KHIDMouseInterfaceProtocol                 KHubSuperSpeedProtocol = 0
	KHIDNoInterfaceProtocol                    KHubSuperSpeedProtocol = 0
	KHubSuperSpeedProtocolValue                KHubSuperSpeedProtocol = 0
	KMSCProtocolBulkOnly                       KHubSuperSpeedProtocol = 0
	KMSCProtocolControlBulk                    KHubSuperSpeedProtocol = 0
	KMSCProtocolControlBulkInterrupt           KHubSuperSpeedProtocol = 0
	KMSCProtocolUSBAttachedSCSI                KHubSuperSpeedProtocol = 0
	KUSB2ComplianceDeviceProtocol              KHubSuperSpeedProtocol = 0
	KUSBBluetoothProgrammingInterfaceProtocol  KHubSuperSpeedProtocol = 0
	KUSBInterfaceAssociationDescriptorProtocol KHubSuperSpeedProtocol = 0
	KUSBVendorSpecificProtocol                 KHubSuperSpeedProtocol = 0
)

func (e KHubSuperSpeedProtocol) String() string {
	switch e {
	case KHIDKeyboardInterfaceProtocol:
		return "KHIDKeyboardInterfaceProtocol"
	default:
		return fmt.Sprintf("KHubSuperSpeedProtocol(%d)", e)
	}
}

type KI uint

const (
	KIO128RGBAFloatPixelFormat KI = 0
	KIO16BE555PixelFormat      KI = 0
	KIO16BE565PixelFormat      KI = 0
	KIO1IndexedGrayPixelFormat KI = 0
	KIO1MonochromePixelFormat  KI = 0
	KIO24BGRPixelFormat        KI = 0
	KIO24RGBPixelFormat        KI = 0
	KIO2IndexedGrayPixelFormat KI = 0
	KIO2IndexedPixelFormat     KI = 0
	KIO32ARGBPixelFormat       KI = 0
	KIO4IndexedGrayPixelFormat KI = 0
	KIO4IndexedPixelFormat     KI = 0
	KIO64BGRAPixelFormat       KI = 0
	KIO8IndexedGrayPixelFormat KI = 0
	KIO8IndexedPixelFormat     KI = 0
	KIOMemoryLedgerTagDefault  KI = 0
	KIOMemoryLedgerTagGraphics KI = 0
	KIOMemoryLedgerTagMedia    KI = 0
	KIOMemoryLedgerTagNeural   KI = 0
	KIOmemoryLedgerTagNetwork  KI = 0
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
	// KINQUIRY_PRODUCT_IDENTIFICATION_Length: # Discussion
	KINQUIRY_PRODUCT_IDENTIFICATION_Length KINQUIRY = 0
	// KINQUIRY_PRODUCT_REVISION_LEVEL_Length: # Discussion
	KINQUIRY_PRODUCT_REVISION_LEVEL_Length KINQUIRY = 0
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
	// KINQUIRY_VENDOR_IDENTIFICATION_Length: # Discussion
	KINQUIRY_VENDOR_IDENTIFICATION_Length KINQUIRY = 0
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
	KIOAnalogSetupExpected          KIO = 0
	KIOBufferDescriptorMemoryFlags  KIO = 0
	KIOCLUTPixels                   KIO = 0
	KIOCSyncDisable                 KIO = 0
	KIOCapturedAttribute            KIO = 0
	KIOColorimetryAdobeRGB          KIO = 0
	KIOColorimetryBT2020            KIO = 0
	KIOColorimetryBT2100            KIO = 0
	KIOColorimetryBT601             KIO = 0
	KIOColorimetryBT709             KIO = 0
	KIOColorimetryDCIP3             KIO = 0
	KIOColorimetryNativeRGB         KIO = 0
	KIOColorimetryNotSupported      KIO = 0
	KIOColorimetryWGRGB             KIO = 0
	KIOColorimetrysRGB              KIO = 0
	KIOColorimetryxvYCC             KIO = 0
	KIOCopybackCache                KIO = 0
	KIOCopybackInnerCache           KIO = 0
	KIODefaultCache                 KIO = 0
	KIODetailedTimingValid          KIO = 0
	KIODisplayAssertionConnectType  KIO = 0
	KIOFBDisplayStateValue          KIO = 0
	KIOFBServerConnectType          KIO = 0
	KIOFBSharedConnectType          KIO = 0
	KIOFixedCLUTPixels              KIO = 0
	KIOGDiagnoseConnectType         KIO = 0
	KIOGDiagnoseGTraceType          KIO = 0
	KIOHSyncDisable                 KIO = 0
	KIOInhibitCache                 KIO = 0
	KIOInterlacedCEATiming          KIO = 0
	KIOKextSpinDump                 KIO = 0
	KIOLogDTree                     KIO = 0
	KIOMemoryPageable               KIO = 0
	KIOMonoDirectPixels             KIO = 0
	KIOMonoInverseDirectPixels      KIO = 0
	KIONoSeparateSyncControl        KIO = 0
	KIOPostedCombinedReordered      KIO = 0
	KIOPostedReordered              KIO = 0
	KIOPostedWrite                  KIO = 0
	KIORGBDirectPixels              KIO = 0
	KIORGBSignedDirectPixels        KIO = 0
	KIORGBSignedFloatingPointPixels KIO = 0
	KIORealTimeCache                KIO = 0
	KIOScalingInfoValid             KIO = 0
	KIOSyncOnBlue                   KIO = 0
	KIOSyncOnGreen                  KIO = 0
	KIOSyncOnRed                    KIO = 0
	KIOTriStateSyncs                KIO = 0
	KIOVSyncDisable                 KIO = 0
	KIOWriteCombineCache            KIO = 0
	KIOWriteThruCache               KIO = 0
)

func (e KIO) String() string {
	switch e {
	case KIOAnalogSetupExpected:
		return "KIOAnalogSetupExpected"
	default:
		return fmt.Sprintf("KIO(%d)", e)
	}
}

type KIOACPI uint

const (
	KIOACPIBusNumberRange KIOACPI = 0
	KIOACPIIORange        KIOACPI = 0
	KIOACPIMemoryRange    KIOACPI = 0
)

func (e KIOACPI) String() string {
	switch e {
	case KIOACPIBusNumberRange:
		return "KIOACPIBusNumberRange"
	default:
		return fmt.Sprintf("KIOACPI(%d)", e)
	}
}

type KIOACPIAddressSpaceIDSystem uint

const (
	KIOACPIAddressSpaceIDSystemIO     KIOACPIAddressSpaceIDSystem = 0
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

type KIOACPIAddressSpaceOp uint

const (
	KIOACPIAddressSpaceOpRead KIOACPIAddressSpaceOp = 0
)

func (e KIOACPIAddressSpaceOp) String() string {
	switch e {
	case KIOACPIAddressSpaceOpRead:
		return "KIOACPIAddressSpaceOpRead"
	default:
		return fmt.Sprintf("KIOACPIAddressSpaceOp(%d)", e)
	}
}

type KIOACPIDevicePowerState uint

const (
	KIOACPIDevicePowerStateCount KIOACPIDevicePowerState = 0
)

func (e KIOACPIDevicePowerState) String() string {
	switch e {
	case KIOACPIDevicePowerStateCount:
		return "KIOACPIDevicePowerStateCount"
	default:
		return fmt.Sprintf("KIOACPIDevicePowerState(%d)", e)
	}
}

type KIOACPIFixedEvent uint

const (
	KIOACPIFixedEventPMTimer     KIOACPIFixedEvent = 0
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

type KIOAG uint

const (
	KIOAGP4GbAddressing KIOAG = 0
	KIOAGP4xDataRate    KIOAG = 0
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
	KIOAGPIdle             KIOAGP = 0
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

type KIOAGPDefault uint

const (
	KIOAGPDefaultStatus KIOAGPDefault = 0
)

func (e KIOAGPDefault) String() string {
	switch e {
	case KIOAGPDefaultStatus:
		return "KIOAGPDefaultStatus"
	default:
		return fmt.Sprintf("KIOAGPDefault(%d)", e)
	}
}

type KIOAGPDisable uint

const (
	KIOAGPDisableAGPWrites KIOAGPDisable = 0
	KIOAGPDisableFeature6  KIOAGPDisable = 0
)

func (e KIOAGPDisable) String() string {
	switch e {
	case KIOAGPDisableAGPWrites:
		return "KIOAGPDisableAGPWrites"
	default:
		return fmt.Sprintf("KIOAGPDisable(%d)", e)
	}
}

type KIOAGPGart uint

const (
	KIOAGPGartInvalidate KIOAGPGart = 0
)

func (e KIOAGPGart) String() string {
	switch e {
	case KIOAGPGartInvalidate:
		return "KIOAGPGartInvalidate"
	default:
		return fmt.Sprintf("KIOAGPGart(%d)", e)
	}
}

type KIOAGPStateEnable uint

const (
	KIOAGPStateEnablePending KIOAGPStateEnable = 0
)

func (e KIOAGPStateEnable) String() string {
	switch e {
	case KIOAGPStateEnablePending:
		return "KIOAGPStateEnablePending"
	default:
		return fmt.Sprintf("KIOAGPStateEnable(%d)", e)
	}
}

type KIOATA uint

const (
	KIOATADefaultPerformance KIOATA = 0
	KIOATAMaxPerformance     KIOATA = 0
	KIOATAMaxPowerSavings    KIOATA = 0
)

func (e KIOATA) String() string {
	switch e {
	case KIOATADefaultPerformance:
		return "KIOATADefaultPerformance"
	default:
		return fmt.Sprintf("KIOATA(%d)", e)
	}
}

type KIOATAFeature48BitLB uint

const (
	KIOATAFeature48BitLBA KIOATAFeature48BitLB = 0
)

func (e KIOATAFeature48BitLB) String() string {
	switch e {
	case KIOATAFeature48BitLBA:
		return "KIOATAFeature48BitLBA"
	default:
		return fmt.Sprintf("KIOATAFeature48BitLB(%d)", e)
	}
}

type KIOATAMaxBlocksPer uint

const (
	KIOATAMaxBlocksPerXfer KIOATAMaxBlocksPer = 0
)

func (e KIOATAMaxBlocksPer) String() string {
	switch e {
	case KIOATAMaxBlocksPerXfer:
		return "KIOATAMaxBlocksPerXfer"
	default:
		return fmt.Sprintf("KIOATAMaxBlocksPer(%d)", e)
	}
}

type KIOATASector uint

const (
	KIOATASectorCount16Bit KIOATASector = 0
	KIOATASectorCount8Bit  KIOATASector = 0
)

func (e KIOATASector) String() string {
	switch e {
	case KIOATASectorCount16Bit:
		return "KIOATASectorCount16Bit"
	default:
		return fmt.Sprintf("KIOATASector(%d)", e)
	}
}

type KIOAccel uint

const (
	KIOAccelNumClientTypes          KIOAccel = 0
	KIOAccelNumSurfaceMethods       KIOAccel = 0
	KIOAccelSurface2ClientType      KIOAccel = 0
	KIOAccelSurfaceClientType       KIOAccel = 0
	KIOAccelSurfaceControl          KIOAccel = 0
	KIOAccelSurfaceFlush            KIOAccel = 0
	KIOAccelSurfaceRead             KIOAccel = 0
	KIOAccelSurfaceSetShapeBacking  KIOAccel = 0
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

type KIOAccelNumSurfaceMemory uint

const (
	KIOAccelNumSurfaceMemoryTypes KIOAccelNumSurfaceMemory = 0
)

func (e KIOAccelNumSurfaceMemory) String() string {
	switch e {
	case KIOAccelNumSurfaceMemoryTypes:
		return "KIOAccelNumSurfaceMemoryTypes"
	default:
		return fmt.Sprintf("KIOAccelNumSurfaceMemory(%d)", e)
	}
}

type KIOAccelPrivateI uint

const (
	KIOAccelPrivateID KIOAccelPrivateI = 0
)

func (e KIOAccelPrivateI) String() string {
	switch e {
	case KIOAccelPrivateID:
		return "KIOAccelPrivateID"
	default:
		return fmt.Sprintf("KIOAccelPrivateI(%d)", e)
	}
}

type KIOAccelSurfaceBeamSync uint

const (
	KIOAccelSurfaceBeamSyncSwaps KIOAccelSurfaceBeamSync = 0
)

func (e KIOAccelSurfaceBeamSync) String() string {
	switch e {
	case KIOAccelSurfaceBeamSyncSwaps:
		return "KIOAccelSurfaceBeamSyncSwaps"
	default:
		return fmt.Sprintf("KIOAccelSurfaceBeamSync(%d)", e)
	}
}

type KIOAccelSurfaceLockIn uint

const (
	KIOAccelSurfaceLockInAccel KIOAccelSurfaceLockIn = 0
)

func (e KIOAccelSurfaceLockIn) String() string {
	switch e {
	case KIOAccelSurfaceLockInAccel:
		return "KIOAccelSurfaceLockInAccel"
	default:
		return fmt.Sprintf("KIOAccelSurfaceLockIn(%d)", e)
	}
}

type KIOAccelSurfaceModeBeam uint

const (
	KIOAccelSurfaceModeBeamSync KIOAccelSurfaceModeBeam = 0
)

func (e KIOAccelSurfaceModeBeam) String() string {
	switch e {
	case KIOAccelSurfaceModeBeamSync:
		return "KIOAccelSurfaceModeBeamSync"
	default:
		return fmt.Sprintf("KIOAccelSurfaceModeBeam(%d)", e)
	}
}

type KIOAccelSurfaceShapeAssembly uint

const (
	KIOAccelSurfaceShapeAssemblyBit KIOAccelSurfaceShapeAssembly = 0
)

func (e KIOAccelSurfaceShapeAssembly) String() string {
	switch e {
	case KIOAccelSurfaceShapeAssemblyBit:
		return "KIOAccelSurfaceShapeAssemblyBit"
	default:
		return fmt.Sprintf("KIOAccelSurfaceShapeAssembly(%d)", e)
	}
}

type KIOAccelSurfaceStateIdle uint

const (
	KIOAccelSurfaceStateIdleBit KIOAccelSurfaceStateIdle = 0
)

func (e KIOAccelSurfaceStateIdle) String() string {
	switch e {
	case KIOAccelSurfaceStateIdleBit:
		return "KIOAccelSurfaceStateIdleBit"
	default:
		return fmt.Sprintf("KIOAccelSurfaceStateIdle(%d)", e)
	}
}

type KIOAccelVolatile uint

const (
	KIOAccelVolatileSurface KIOAccelVolatile = 0
)

func (e KIOAccelVolatile) String() string {
	switch e {
	case KIOAccelVolatileSurface:
		return "KIOAccelVolatileSurface"
	default:
		return fmt.Sprintf("KIOAccelVolatile(%d)", e)
	}
}

type KIOAsyncCallout uint

const (
	KIOAsyncCalloutCount KIOAsyncCallout = 0
)

func (e KIOAsyncCallout) String() string {
	switch e {
	case KIOAsyncCalloutCount:
		return "KIOAsyncCalloutCount"
	default:
		return fmt.Sprintf("KIOAsyncCallout(%d)", e)
	}
}

type KIOAudio uint

const (
	KIOAudioBuiltInSystemClockDomain         KIOAudio = 0
	KIOAudioInputPortSubTypeCD               KIOAudio = 0
	KIOAudioNewClockDomain                   KIOAudio = 0
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

type KIOAudioClockSelectorTypeADAT9 uint

const (
	KIOAudioClockSelectorTypeADAT9Pin KIOAudioClockSelectorTypeADAT9 = 0
)

func (e KIOAudioClockSelectorTypeADAT9) String() string {
	switch e {
	case KIOAudioClockSelectorTypeADAT9Pin:
		return "KIOAudioClockSelectorTypeADAT9Pin"
	default:
		return fmt.Sprintf("KIOAudioClockSelectorTypeADAT9(%d)", e)
	}
}

type KIOAudioControlChannelID uint

const (
	KIOAudioControlChannelIDAll             KIOAudioControlChannelID = 0
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

type KIOAudioControlType uint

const (
	KIOAudioControlTypeToggle KIOAudioControlType = 0
)

func (e KIOAudioControlType) String() string {
	switch e {
	case KIOAudioControlTypeToggle:
		return "KIOAudioControlTypeToggle"
	default:
		return fmt.Sprintf("KIOAudioControlType(%d)", e)
	}
}

type KIOAudioControlUsage uint

const (
	KIOAudioControlUsageCoreAudioProperty KIOAudioControlUsage = 0
	KIOAudioControlUsageInput             KIOAudioControlUsage = 0
	KIOAudioControlUsageOutput            KIOAudioControlUsage = 0
	KIOAudioControlUsagePassThru          KIOAudioControlUsage = 0
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
	// KIOAudioDeviceActive: # Discussion
	KIOAudioDeviceActive KIOAudioDevice = 0
	// KIOAudioDeviceIdle: # Discussion
	KIOAudioDeviceIdle KIOAudioDevice = 0
	// KIOAudioDeviceSleep: # Discussion
	KIOAudioDeviceSleep KIOAudioDevice = 0
)

func (e KIOAudioDevice) String() string {
	switch e {
	case KIOAudioDeviceActive:
		return "KIOAudioDeviceActive"
	default:
		return fmt.Sprintf("KIOAudioDevice(%d)", e)
	}
}

type KIOAudioDeviceCanBeDefault uint

const (
	KIOAudioDeviceCanBeDefaultInput KIOAudioDeviceCanBeDefault = 0
)

func (e KIOAudioDeviceCanBeDefault) String() string {
	switch e {
	case KIOAudioDeviceCanBeDefaultInput:
		return "KIOAudioDeviceCanBeDefaultInput"
	default:
		return fmt.Sprintf("KIOAudioDeviceCanBeDefault(%d)", e)
	}
}

type KIOAudioDeviceTransportType uint

const (
	KIOAudioDeviceTransportTypeAVB       KIOAudioDeviceTransportType = 0
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
	KIOAudioEngineAllNotifications               KIOAudioEngine = 0
	KIOAudioEngineChangeNotification             KIOAudioEngine = 0
	KIOAudioEnginePausedNotification             KIOAudioEngine = 0
	KIOAudioEngineResumedNotification            KIOAudioEngine = 0
	KIOAudioEngineStartedNotification            KIOAudioEngine = 0
	KIOAudioEngineStoppedNotification            KIOAudioEngine = 0
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

type KIOAudioLevelControlNegative uint

const (
	KIOAudioLevelControlNegativeInfinity KIOAudioLevelControlNegative = 0
)

func (e KIOAudioLevelControlNegative) String() string {
	switch e {
	case KIOAudioLevelControlNegativeInfinity:
		return "KIOAudioLevelControlNegativeInfinity"
	default:
		return fmt.Sprintf("KIOAudioLevelControlNegative(%d)", e)
	}
}

type KIOAudioLevelControlSubTypeLFE uint

const (
	KIOAudioLevelControlSubTypeLFEVolume KIOAudioLevelControlSubTypeLFE = 0
)

func (e KIOAudioLevelControlSubTypeLFE) String() string {
	switch e {
	case KIOAudioLevelControlSubTypeLFEVolume:
		return "KIOAudioLevelControlSubTypeLFEVolume"
	default:
		return fmt.Sprintf("KIOAudioLevelControlSubTypeLFE(%d)", e)
	}
}

type KIOAudioPortType uint

const (
	KIOAudioPortTypeInput      KIOAudioPortType = 0
	KIOAudioPortTypeMixer      KIOAudioPortType = 0
	KIOAudioPortTypeOutput     KIOAudioPortType = 0
	KIOAudioPortTypePassThru   KIOAudioPortType = 0
	KIOAudioPortTypeProcessing KIOAudioPortType = 0
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
	KIOAudioSMPTETimeRunning      KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType2398     KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType24       KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType25       KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType2997     KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType2997Drop KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType30       KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType30Drop   KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType50       KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType5994     KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType5994Drop KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType60       KIOAudioSMPTETime = 0
	KIOAudioSMPTETimeType60Drop   KIOAudioSMPTETime = 0
)

func (e KIOAudioSMPTETime) String() string {
	switch e {
	case KIOAudioSMPTETimeRunning:
		return "KIOAudioSMPTETimeRunning"
	default:
		return fmt.Sprintf("KIOAudioSMPTETime(%d)", e)
	}
}

type KIOAudioSelectorControlSelectionValueC uint

const (
	KIOAudioSelectorControlSelectionValueCD KIOAudioSelectorControlSelectionValueC = 0
)

func (e KIOAudioSelectorControlSelectionValueC) String() string {
	switch e {
	case KIOAudioSelectorControlSelectionValueCD:
		return "KIOAudioSelectorControlSelectionValueCD"
	default:
		return fmt.Sprintf("KIOAudioSelectorControlSelectionValueC(%d)", e)
	}
}

type KIOAudioStreamAlignmentHigh uint

const (
	KIOAudioStreamAlignmentHighByte KIOAudioStreamAlignmentHigh = 0
)

func (e KIOAudioStreamAlignmentHigh) String() string {
	switch e {
	case KIOAudioStreamAlignmentHighByte:
		return "KIOAudioStreamAlignmentHighByte"
	default:
		return fmt.Sprintf("KIOAudioStreamAlignmentHigh(%d)", e)
	}
}

type KIOAudioStreamByteOrder uint

const (
	KIOAudioStreamByteOrderBigEndian    KIOAudioStreamByteOrder = 0
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

type KIOAudioStreamNumericRepresentationIEEE754 uint

const (
	KIOAudioStreamNumericRepresentationIEEE754Float KIOAudioStreamNumericRepresentationIEEE754 = 0
)

func (e KIOAudioStreamNumericRepresentationIEEE754) String() string {
	switch e {
	case KIOAudioStreamNumericRepresentationIEEE754Float:
		return "KIOAudioStreamNumericRepresentationIEEE754Float"
	default:
		return fmt.Sprintf("KIOAudioStreamNumericRepresentationIEEE754(%d)", e)
	}
}

type KIOAudioStreamSampleFormat1937A uint

const (
	KIOAudioStreamSampleFormat1937AC3 KIOAudioStreamSampleFormat1937A = 0
)

func (e KIOAudioStreamSampleFormat1937A) String() string {
	switch e {
	case KIOAudioStreamSampleFormat1937AC3:
		return "KIOAudioStreamSampleFormat1937AC3"
	default:
		return fmt.Sprintf("KIOAudioStreamSampleFormat1937A(%d)", e)
	}
}

type KIOAudioTimeStampHostTime uint

const (
	KIOAudioTimeStampHostTimeValid KIOAudioTimeStampHostTime = 0
)

func (e KIOAudioTimeStampHostTime) String() string {
	switch e {
	case KIOAudioTimeStampHostTimeValid:
		return "KIOAudioTimeStampHostTimeValid"
	default:
		return fmt.Sprintf("KIOAudioTimeStampHostTime(%d)", e)
	}
}

type KIOAudioTimeStampSampleHostTime uint

const (
	KIOAudioTimeStampSampleHostTimeValid KIOAudioTimeStampSampleHostTime = 0
)

func (e KIOAudioTimeStampSampleHostTime) String() string {
	switch e {
	case KIOAudioTimeStampSampleHostTimeValid:
		return "KIOAudioTimeStampSampleHostTimeValid"
	default:
		return fmt.Sprintf("KIOAudioTimeStampSampleHostTime(%d)", e)
	}
}

type KIOBattery uint

const (
	KIOBatteryCharge         KIOBattery = 0
	KIOBatteryChargerConnect KIOBattery = 0
	KIOBatteryInstalled      KIOBattery = 0
)

func (e KIOBattery) String() string {
	switch e {
	case KIOBatteryCharge:
		return "KIOBatteryCharge"
	default:
		return fmt.Sprintf("KIOBattery(%d)", e)
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
	KIOBlitAllOptions             KIOBlit = 0
	KIOBlitBeamSyncSpin           KIOBlit = 0
	KIOBlitFlushWithSwap          KIOBlit = 0
	KIOBlitFramebufferDestination KIOBlit = 0
	KIOBlitSurfaceDestination     KIOBlit = 0
	KIOBlitWaitGlobal             KIOBlit = 0
)

func (e KIOBlit) String() string {
	switch e {
	case KIOBlitAllOptions:
		return "KIOBlitAllOptions"
	default:
		return fmt.Sprintf("KIOBlit(%d)", e)
	}
}

type KIOBlitBeamSync uint

const (
	KIOBlitBeamSyncSwaps KIOBlitBeamSync = 0
)

func (e KIOBlitBeamSync) String() string {
	switch e {
	case KIOBlitBeamSyncSwaps:
		return "KIOBlitBeamSyncSwaps"
	default:
		return fmt.Sprintf("KIOBlitBeamSync(%d)", e)
	}
}

type KIOBlitColorSpace uint

const (
	KIOBlitColorSpaceTypes KIOBlitColorSpace = 0
)

func (e KIOBlitColorSpace) String() string {
	switch e {
	case KIOBlitColorSpaceTypes:
		return "KIOBlitColorSpaceTypes"
	default:
		return fmt.Sprintf("KIOBlitColorSpace(%d)", e)
	}
}

type KIOBlitMemoryRequiresHost uint

const (
	KIOBlitMemoryRequiresHostFlush KIOBlitMemoryRequiresHost = 0
)

func (e KIOBlitMemoryRequiresHost) String() string {
	switch e {
	case KIOBlitMemoryRequiresHostFlush:
		return "KIOBlitMemoryRequiresHostFlush"
	default:
		return fmt.Sprintf("KIOBlitMemoryRequiresHost(%d)", e)
	}
}

type KIOBlitSource uint

const (
	KIOBlitSourceCGSSurface  KIOBlitSource = 0
	KIOBlitSourceDefault     KIOBlitSource = 0
	KIOBlitSourceFramebuffer KIOBlitSource = 0
	KIOBlitSourceIsSame      KIOBlitSource = 0
	KIOBlitSourceMemory      KIOBlitSource = 0
	KIOBlitSourceOOLMemory   KIOBlitSource = 0
	KIOBlitSourceOOLPattern  KIOBlitSource = 0
	KIOBlitSourcePattern     KIOBlitSource = 0
	KIOBlitSourceSolid       KIOBlitSource = 0
)

func (e KIOBlitSource) String() string {
	switch e {
	case KIOBlitSourceCGSSurface:
		return "KIOBlitSourceCGSSurface"
	default:
		return fmt.Sprintf("KIOBlitSource(%d)", e)
	}
}

type KIOBlitSynchronizeFlushHost uint

const (
	KIOBlitSynchronizeFlushHostWrites KIOBlitSynchronizeFlushHost = 0
)

func (e KIOBlitSynchronizeFlushHost) String() string {
	switch e {
	case KIOBlitSynchronizeFlushHostWrites:
		return "KIOBlitSynchronizeFlushHostWrites"
	default:
		return fmt.Sprintf("KIOBlitSynchronizeFlushHost(%d)", e)
	}
}

type KIOBlitType uint

const (
	KIOBlitTypeSourceKeyColorNotEqual KIOBlitType = 0
	KIOBlitTypeVerbMask               KIOBlitType = 0
)

func (e KIOBlitType) String() string {
	switch e {
	case KIOBlitTypeSourceKeyColorNotEqual:
		return "KIOBlitTypeSourceKeyColorNotEqual"
	default:
		return fmt.Sprintf("KIOBlitType(%d)", e)
	}
}

type KIOBlitUnlockWith uint

const (
	KIOBlitUnlockWithSwap KIOBlitUnlockWith = 0
)

func (e KIOBlitUnlockWith) String() string {
	switch e {
	case KIOBlitUnlockWithSwap:
		return "KIOBlitUnlockWithSwap"
	default:
		return fmt.Sprintf("KIOBlitUnlockWith(%d)", e)
	}
}

type KIOCatalog uint

const (
	// KIOCatalogAddDrivers: # Discussion
	KIOCatalogAddDrivers KIOCatalog = 0
	// KIOCatalogAddDriversNoMatch: # Discussion
	KIOCatalogAddDriversNoMatch KIOCatalog = 0
	// KIOCatalogKextdActive: # Discussion
	KIOCatalogKextdActive KIOCatalog = 0
	// KIOCatalogKextdFinishedLaunching: # Discussion
	KIOCatalogKextdFinishedLaunching KIOCatalog = 0
	// KIOCatalogModuleTerminate: # Discussion
	KIOCatalogModuleTerminate KIOCatalog = 0
	// KIOCatalogRemoveDrivers: # Discussion
	KIOCatalogRemoveDrivers KIOCatalog = 0
	// KIOCatalogRemoveDriversNoMatch: # Discussion
	KIOCatalogRemoveDriversNoMatch        KIOCatalog = 0
	KIOCatalogRemoveKernelLinker__Removed KIOCatalog = 0
	// KIOCatalogResetDrivers: # Discussion
	KIOCatalogResetDrivers KIOCatalog = 0
	// KIOCatalogResetDriversNoMatch: # Discussion
	KIOCatalogResetDriversNoMatch KIOCatalog = 0
	// KIOCatalogServiceTerminate: # Discussion
	KIOCatalogServiceTerminate       KIOCatalog = 0
	KIOCatalogStartMatching__Removed KIOCatalog = 0
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
	KIOCatalogGetCacheMissList    KIOCatalogGet = 0
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

type KIOCatalogReset uint

const (
	// KIOCatalogResetDefault: # Discussion
	KIOCatalogResetDefault KIOCatalogReset = 0
)

func (e KIOCatalogReset) String() string {
	switch e {
	case KIOCatalogResetDefault:
		return "KIOCatalogResetDefault"
	default:
		return fmt.Sprintf("KIOCatalogReset(%d)", e)
	}
}

type KIOConnectMethodVarOutput uint

const (
	KIOConnectMethodVarOutputSize KIOConnectMethodVarOutput = 0
)

func (e KIOConnectMethodVarOutput) String() string {
	switch e {
	case KIOConnectMethodVarOutputSize:
		return "KIOConnectMethodVarOutputSize"
	default:
		return fmt.Sprintf("KIOConnectMethodVarOutput(%d)", e)
	}
}

type KIOConnectionBuilt uint

const (
	KIOConnectionBuiltIn KIOConnectionBuilt = 0
)

func (e KIOConnectionBuilt) String() string {
	switch e {
	case KIOConnectionBuiltIn:
		return "KIOConnectionBuiltIn"
	default:
		return fmt.Sprintf("KIOConnectionBuilt(%d)", e)
	}
}

type KIODDC uint

const (
	KIODDCHigh KIODDC = 0
)

func (e KIODDC) String() string {
	switch e {
	case KIODDCHigh:
		return "KIODDCHigh"
	default:
		return fmt.Sprintf("KIODDC(%d)", e)
	}
}

type KIODDCBlockTypeEDI uint

const (
	KIODDCBlockTypeEDID KIODDCBlockTypeEDI = 0
)

func (e KIODDCBlockTypeEDI) String() string {
	switch e {
	case KIODDCBlockTypeEDID:
		return "KIODDCBlockTypeEDID"
	default:
		return fmt.Sprintf("KIODDCBlockTypeEDI(%d)", e)
	}
}

type KIODDCForce uint

const (
	KIODDCForceRead KIODDCForce = 0
)

func (e KIODDCForce) String() string {
	switch e {
	case KIODDCForceRead:
		return "KIODDCForceRead"
	default:
		return fmt.Sprintf("KIODDCForce(%d)", e)
	}
}

type KIODK uint

const (
	KIODKDisableCDHashChecking           KIODK = 0
	KIODKDisableCheckInTokenVerification KIODK = 0
	KIODKDisableDextLaunch               KIODK = 0
	KIODKDisableDextTag                  KIODK = 0
	KIODKDisableEntitlementChecking      KIODK = 0
	KIODKDisablePM                       KIODK = 0
	KIODKEnable                          KIODK = 0
	KIODKLogIPC                          KIODK = 0
	KIODKLogMessages                     KIODK = 0
	KIODKLogPM                           KIODK = 0
	KIODKLogSetup                        KIODK = 0
)

func (e KIODK) String() string {
	switch e {
	case KIODKDisableCDHashChecking:
		return "KIODKDisableCDHashChecking"
	default:
		return fmt.Sprintf("KIODK(%d)", e)
	}
}

type KIODMACommandCompleteDMANo uint

const (
	KIODMACommandCompleteDMANoOptions KIODMACommandCompleteDMANo = 0
)

func (e KIODMACommandCompleteDMANo) String() string {
	switch e {
	case KIODMACommandCompleteDMANoOptions:
		return "KIODMACommandCompleteDMANoOptions"
	default:
		return fmt.Sprintf("KIODMACommandCompleteDMANo(%d)", e)
	}
}

type KIODMACommandCreateNo uint

const (
	KIODMACommandCreateNoOptions KIODMACommandCreateNo = 0
)

func (e KIODMACommandCreateNo) String() string {
	switch e {
	case KIODMACommandCreateNoOptions:
		return "KIODMACommandCreateNoOptions"
	default:
		return fmt.Sprintf("KIODMACommandCreateNo(%d)", e)
	}
}

type KIODMACommandPerformOperationOption uint

const (
	KIODMACommandPerformOperationOptionRead  KIODMACommandPerformOperationOption = 0
	KIODMACommandPerformOperationOptionWrite KIODMACommandPerformOperationOption = 0
	KIODMACommandPerformOperationOptionZero  KIODMACommandPerformOperationOption = 0
)

func (e KIODMACommandPerformOperationOption) String() string {
	switch e {
	case KIODMACommandPerformOperationOptionRead:
		return "KIODMACommandPerformOperationOptionRead"
	default:
		return fmt.Sprintf("KIODMACommandPerformOperationOption(%d)", e)
	}
}

type KIODMACommandPrepareForDMANo uint

const (
	KIODMACommandPrepareForDMANoOptions KIODMACommandPrepareForDMANo = 0
)

func (e KIODMACommandPrepareForDMANo) String() string {
	switch e {
	case KIODMACommandPrepareForDMANoOptions:
		return "KIODMACommandPrepareForDMANoOptions"
	default:
		return fmt.Sprintf("KIODMACommandPrepareForDMANo(%d)", e)
	}
}

type KIODMACommandSpecificationNo uint

const (
	KIODMACommandSpecificationNoOptions KIODMACommandSpecificationNo = 0
)

func (e KIODMACommandSpecificationNo) String() string {
	switch e {
	case KIODMACommandSpecificationNoOptions:
		return "KIODMACommandSpecificationNoOptions"
	default:
		return fmt.Sprintf("KIODMACommandSpecificationNo(%d)", e)
	}
}

type KIODMAMap uint

const (
	KIODMAMapDeviceMemory KIODMAMap = 0
	KIODMAMapIdentityMap  KIODMAMap = 0
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
	KIODMAMapOptionBypassed    KIODMAMapOption = 0
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
	KIODPEventContentProtection    KIODPEvent = 0
)

func (e KIODPEvent) String() string {
	switch e {
	case KIODPEventAutomatedTestRequest:
		return "KIODPEventAutomatedTestRequest"
	default:
		return fmt.Sprintf("KIODPEvent(%d)", e)
	}
}

type KIODT uint

const (
	KIODTExclusive KIODT = 0
)

func (e KIODT) String() string {
	switch e {
	case KIODTExclusive:
		return "KIODTExclusive"
	default:
		return fmt.Sprintf("KIODT(%d)", e)
	}
}

type KIODTInterrupt uint

const (
	KIODTInterruptShared KIODTInterrupt = 0
)

func (e KIODTInterrupt) String() string {
	switch e {
	case KIODTInterruptShared:
		return "KIODTInterruptShared"
	default:
		return fmt.Sprintf("KIODTInterrupt(%d)", e)
	}
}

type KIODebuggerLock uint

const (
	// KIODebuggerLockTaken: # Discussion
	KIODebuggerLockTaken KIODebuggerLock = 0
)

func (e KIODebuggerLock) String() string {
	switch e {
	case KIODebuggerLockTaken:
		return "KIODebuggerLockTaken"
	default:
		return fmt.Sprintf("KIODebuggerLock(%d)", e)
	}
}

type KIODefaultMemory uint

const (
	KIODefaultMemoryType KIODefaultMemory = 0
)

func (e KIODefaultMemory) String() string {
	switch e {
	case KIODefaultMemoryType:
		return "KIODefaultMemoryType"
	default:
		return fmt.Sprintf("KIODefaultMemory(%d)", e)
	}
}

type KIODefaultProbe uint

const (
	KIODefaultProbeScore KIODefaultProbe = 0
)

func (e KIODefaultProbe) String() string {
	switch e {
	case KIODefaultProbeScore:
		return "KIODefaultProbeScore"
	default:
		return fmt.Sprintf("KIODefaultProbe(%d)", e)
	}
}

type KIODirection uint

const (
	KIODirectionCompleteWithDataValid KIODirection = 0
	KIODirectionCompleteWithError     KIODirection = 0
	KIODirectionIn                    KIODirection = 0
	KIODirectionInOut                 KIODirection = 0
	KIODirectionNone                  KIODirection = 0
	KIODirectionOut                   KIODirection = 0
	KIODirectionOutIn                 KIODirection = 0
	KIODirectionPrepareNoFault        KIODirection = 0
	KIODirectionPrepareNonCoherent    KIODirection = 0
	KIODirectionPrepareReserved1      KIODirection = 0
	KIODirectionPrepareToPhys32       KIODirection = 0
)

func (e KIODirection) String() string {
	switch e {
	case KIODirectionCompleteWithDataValid:
		return "KIODirectionCompleteWithDataValid"
	default:
		return fmt.Sprintf("KIODirection(%d)", e)
	}
}

type KIODispatchQueueMethodsNot uint

const (
	KIODispatchQueueMethodsNotSynchronized KIODispatchQueueMethodsNot = 0
)

func (e KIODispatchQueueMethodsNot) String() string {
	switch e {
	case KIODispatchQueueMethodsNotSynchronized:
		return "KIODispatchQueueMethodsNotSynchronized"
	default:
		return fmt.Sprintf("KIODispatchQueueMethodsNot(%d)", e)
	}
}

type KIODispatchQueueSleepReuse uint

const (
	KIODispatchQueueSleepReuseEvent KIODispatchQueueSleepReuse = 0
)

func (e KIODispatchQueueSleepReuse) String() string {
	switch e {
	case KIODispatchQueueSleepReuseEvent:
		return "KIODispatchQueueSleepReuseEvent"
	default:
		return fmt.Sprintf("KIODispatchQueueSleepReuse(%d)", e)
	}
}

type KIODispatchQueueWakeup uint

const (
	KIODispatchQueueWakeupAll KIODispatchQueueWakeup = 0
)

func (e KIODispatchQueueWakeup) String() string {
	switch e {
	case KIODispatchQueueWakeupAll:
		return "KIODispatchQueueWakeupAll"
	default:
		return fmt.Sprintf("KIODispatchQueueWakeup(%d)", e)
	}
}

type KIODisplay uint

const (
	KIODisplayMaxPowerState  KIODisplay = 0
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

type KIODisplayColor uint

const (
	KIODisplayColorMode KIODisplayColor = 0
)

func (e KIODisplayColor) String() string {
	switch e {
	case KIODisplayColorMode:
		return "KIODisplayColorMode"
	default:
		return fmt.Sprintf("KIODisplayColor(%d)", e)
	}
}

type KIODisplayDither uint

const (
	KIODisplayDitherAll KIODisplayDither = 0
)

func (e KIODisplayDither) String() string {
	switch e {
	case KIODisplayDitherAll:
		return "KIODisplayDitherAll"
	default:
		return fmt.Sprintf("KIODisplayDither(%d)", e)
	}
}

type KIODisplayModeID uint

const (
	KIODisplayModeIDBootProgrammable KIODisplayModeID = 0
	KIODisplayModeIDReservedBase     KIODisplayModeID = 0
)

func (e KIODisplayModeID) String() string {
	switch e {
	case KIODisplayModeIDBootProgrammable:
		return "KIODisplayModeIDBootProgrammable"
	default:
		return fmt.Sprintf("KIODisplayModeID(%d)", e)
	}
}

type KIODisplayNeedsCEA uint

const (
	KIODisplayNeedsCEAUnderscan KIODisplayNeedsCEA = 0
)

func (e KIODisplayNeedsCEA) String() string {
	switch e {
	case KIODisplayNeedsCEAUnderscan:
		return "KIODisplayNeedsCEAUnderscan"
	default:
		return fmt.Sprintf("KIODisplayNeedsCEA(%d)", e)
	}
}

type KIODisplayPowerStateMin uint

const (
	KIODisplayPowerStateMinUsable KIODisplayPowerStateMin = 0
)

func (e KIODisplayPowerStateMin) String() string {
	switch e {
	case KIODisplayPowerStateMinUsable:
		return "KIODisplayPowerStateMinUsable"
	default:
		return fmt.Sprintf("KIODisplayPowerStateMin(%d)", e)
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

type KIODynamicRangeDolbyNormal uint

const (
	KIODynamicRangeDolbyNormalMode KIODynamicRangeDolbyNormal = 0
)

func (e KIODynamicRangeDolbyNormal) String() string {
	switch e {
	case KIODynamicRangeDolbyNormalMode:
		return "KIODynamicRangeDolbyNormalMode"
	default:
		return fmt.Sprintf("KIODynamicRangeDolbyNormal(%d)", e)
	}
}

type KIOEnetMulticastMode uint

const (
	KIOEnetMulticastModeFilter KIOEnetMulticastMode = 0
)

func (e KIOEnetMulticastMode) String() string {
	switch e {
	case KIOEnetMulticastModeFilter:
		return "KIOEnetMulticastModeFilter"
	default:
		return fmt.Sprintf("KIOEnetMulticastMode(%d)", e)
	}
}

type KIOEthernetControllerAVBState uint

const (
	KIOEthernetControllerAVBStateAVBEnabled      KIOEthernetControllerAVBState = 0
	KIOEthernetControllerAVBStateActivated       KIOEthernetControllerAVBState = 0
	KIOEthernetControllerAVBStateDisabled        KIOEthernetControllerAVBState = 0
	KIOEthernetControllerAVBStateTimeSyncEnabled KIOEthernetControllerAVBState = 0
)

func (e KIOEthernetControllerAVBState) String() string {
	switch e {
	case KIOEthernetControllerAVBStateAVBEnabled:
		return "KIOEthernetControllerAVBStateAVBEnabled"
	default:
		return fmt.Sprintf("KIOEthernetControllerAVBState(%d)", e)
	}
}

type KIOEthernetControllerAVBStateEvent uint

const (
	KIOEthernetControllerAVBStateEventDisable        KIOEthernetControllerAVBStateEvent = 0
	KIOEthernetControllerAVBStateEventEnable         KIOEthernetControllerAVBStateEvent = 0
	KIOEthernetControllerAVBStateEventStartStreaming KIOEthernetControllerAVBStateEvent = 0
	KIOEthernetControllerAVBStateEventStartTimeSync  KIOEthernetControllerAVBStateEvent = 0
	KIOEthernetControllerAVBStateEventStopStreaming  KIOEthernetControllerAVBStateEvent = 0
	KIOEthernetControllerAVBStateEventStopTimeSync   KIOEthernetControllerAVBStateEvent = 0
)

func (e KIOEthernetControllerAVBStateEvent) String() string {
	switch e {
	case KIOEthernetControllerAVBStateEventDisable:
		return "KIOEthernetControllerAVBStateEventDisable"
	default:
		return fmt.Sprintf("KIOEthernetControllerAVBStateEvent(%d)", e)
	}
}

type KIOEthernetControllerAVBTimeSyncSupport uint

const (
	KIOEthernetControllerAVBTimeSyncSupportHardware KIOEthernetControllerAVBTimeSyncSupport = 0
)

func (e KIOEthernetControllerAVBTimeSyncSupport) String() string {
	switch e {
	case KIOEthernetControllerAVBTimeSyncSupportHardware:
		return "KIOEthernetControllerAVBTimeSyncSupportHardware"
	default:
		return fmt.Sprintf("KIOEthernetControllerAVBTimeSyncSupport(%d)", e)
	}
}

type KIOEthernetWakeOnMagic uint

const (
	// KIOEthernetWakeOnMagicPacket: # Discussion
	KIOEthernetWakeOnMagicPacket KIOEthernetWakeOnMagic = 0
)

func (e KIOEthernetWakeOnMagic) String() string {
	switch e {
	case KIOEthernetWakeOnMagicPacket:
		return "KIOEthernetWakeOnMagicPacket"
	default:
		return fmt.Sprintf("KIOEthernetWakeOnMagic(%d)", e)
	}
}

type KIOEventLinkAssociate uint

const (
	KIOEventLinkAssociateCurrentThread KIOEventLinkAssociate = 0
	KIOEventLinkAssociateOnWait        KIOEventLinkAssociate = 0
)

func (e KIOEventLinkAssociate) String() string {
	switch e {
	case KIOEventLinkAssociateCurrentThread:
		return "KIOEventLinkAssociateCurrentThread"
	default:
		return fmt.Sprintf("KIOEventLinkAssociate(%d)", e)
	}
}

type KIOEventLinkClockMachAbsolute uint

const (
	KIOEventLinkClockMachAbsoluteTime KIOEventLinkClockMachAbsolute = 0
)

func (e KIOEventLinkClockMachAbsolute) String() string {
	switch e {
	case KIOEventLinkClockMachAbsoluteTime:
		return "KIOEventLinkClockMachAbsoluteTime"
	default:
		return fmt.Sprintf("KIOEventLinkClockMachAbsolute(%d)", e)
	}
}

type KIOEventLinkMaxName uint

const (
	KIOEventLinkMaxNameLength KIOEventLinkMaxName = 0
)

func (e KIOEventLinkMaxName) String() string {
	switch e {
	case KIOEventLinkMaxNameLength:
		return "KIOEventLinkMaxNameLength"
	default:
		return fmt.Sprintf("KIOEventLinkMaxName(%d)", e)
	}
}

type KIOExternalMethodArgumentsCurrent uint

const (
	KIOExternalMethodArgumentsCurrentVersion KIOExternalMethodArgumentsCurrent = 0
)

func (e KIOExternalMethodArgumentsCurrent) String() string {
	switch e {
	case KIOExternalMethodArgumentsCurrentVersion:
		return "KIOExternalMethodArgumentsCurrentVersion"
	default:
		return fmt.Sprintf("KIOExternalMethodArgumentsCurrent(%d)", e)
	}
}

type KIOExternalMethodScalarInputCount uint

const (
	// KIOExternalMethodScalarInputCountMax: # Discussion
	KIOExternalMethodScalarInputCountMax KIOExternalMethodScalarInputCount = 0
)

func (e KIOExternalMethodScalarInputCount) String() string {
	switch e {
	case KIOExternalMethodScalarInputCountMax:
		return "KIOExternalMethodScalarInputCountMax"
	default:
		return fmt.Sprintf("KIOExternalMethodScalarInputCount(%d)", e)
	}
}

type KIOFB uint

const (
	KIOFBChangedInterruptType KIOFB = 0
	KIOFBConnectInterruptType KIOFB = 0
	// KIOFBCurrentShmemVersion: # Discussion
	KIOFBCurrentShmemVersion                KIOFB = 0
	KIOFBDisplayPortInterruptType           KIOFB = 0
	KIOFBDisplayPortLinkChangeInterruptType KIOFB = 0
	KIOFBFrameInterruptType                 KIOFB = 0
	KIOFBHBLInterruptType                   KIOFB = 0
	KIOFBMCCSInterruptType                  KIOFB = 0
	KIOFBMainCursorIndex                    KIOFB = 0
	KIOFBNumCursorIndex                     KIOFB = 0
	KIOFBOfflineInterruptType               KIOFB = 0
	KIOFBOnlineInterruptType                KIOFB = 0
	KIOFBShmemVersionMask                   KIOFB = 0
	KIOFBVBLInterruptType                   KIOFB = 0
	KIOFBWakeInterruptType                  KIOFB = 0
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
	KIOFBAVSignalTypeDP      KIOFBAVSignalType = 0
	KIOFBAVSignalTypeDVI     KIOFBAVSignalType = 0
	KIOFBAVSignalTypeHDMI    KIOFBAVSignalType = 0
	KIOFBAVSignalTypeUnknown KIOFBAVSignalType = 0
	KIOFBAVSignalTypeVGA     KIOFBAVSignalType = 0
)

func (e KIOFBAVSignalType) String() string {
	switch e {
	case KIOFBAVSignalTypeDP:
		return "KIOFBAVSignalTypeDP"
	default:
		return fmt.Sprintf("KIOFBAVSignalType(%d)", e)
	}
}

type KIOFBBitRateHB uint

const (
	KIOFBBitRateHBR KIOFBBitRateHB = 0
)

func (e KIOFBBitRateHB) String() string {
	switch e {
	case KIOFBBitRateHBR:
		return "KIOFBBitRateHBR"
	default:
		return fmt.Sprintf("KIOFBBitRateHB(%d)", e)
	}
}

type KIOFBCursorHW uint

const (
	KIOFBCursorHWCapable KIOFBCursorHW = 0
)

func (e KIOFBCursorHW) String() string {
	switch e {
	case KIOFBCursorHWCapable:
		return "KIOFBCursorHWCapable"
	default:
		return fmt.Sprintf("KIOFBCursorHW(%d)", e)
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

type KIOFBHDCPLimit uint

const (
	KIOFBHDCPLimit_AllowAll      KIOFBHDCPLimit = 0
	KIOFBHDCPLimit_NoHDCP1x      KIOFBHDCPLimit = 0
	KIOFBHDCPLimit_NoHDCP20Type0 KIOFBHDCPLimit = 0
	KIOFBHDCPLimit_NoHDCP20Type1 KIOFBHDCPLimit = 0
)

func (e KIOFBHDCPLimit) String() string {
	switch e {
	case KIOFBHDCPLimit_AllowAll:
		return "KIOFBHDCPLimit_AllowAll"
	default:
		return fmt.Sprintf("KIOFBHDCPLimit(%d)", e)
	}
}

type KIOFBHardwareCursor uint

const (
	KIOFBHardwareCursorActive KIOFBHardwareCursor = 0
)

func (e KIOFBHardwareCursor) String() string {
	switch e {
	case KIOFBHardwareCursorActive:
		return "KIOFBHardwareCursorActive"
	default:
		return fmt.Sprintf("KIOFBHardwareCursor(%d)", e)
	}
}

type KIOFBLinkDownspread uint

const (
	KIOFBLinkDownspreadMax  KIOFBLinkDownspread = 0
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
	KIOFBLinkPreEmphasisLevel2 KIOFBLinkPreEmphasis = 0
	KIOFBLinkPreEmphasisLevel3 KIOFBLinkPreEmphasis = 0
)

func (e KIOFBLinkPreEmphasis) String() string {
	switch e {
	case KIOFBLinkPreEmphasisLevel0:
		return "KIOFBLinkPreEmphasisLevel0"
	default:
		return fmt.Sprintf("KIOFBLinkPreEmphasis(%d)", e)
	}
}

type KIOFBLinkScrambler uint

const (
	KIOFBLinkScramblerAlternate KIOFBLinkScrambler = 0
	KIOFBLinkScramblerNormal    KIOFBLinkScrambler = 0
)

func (e KIOFBLinkScrambler) String() string {
	switch e {
	case KIOFBLinkScramblerAlternate:
		return "KIOFBLinkScramblerAlternate"
	default:
		return fmt.Sprintf("KIOFBLinkScrambler(%d)", e)
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
	KIOFBNS_Doze              KIOFBNS = 0
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
	KIOFBNotifyDidSleep            KIOFBNotify = 0
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
	KIOFBNotifyEvent_All         KIOFBNotifyEvent = 0
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
	KIOFBNotifyGroupID_AppleGraphicsControl          KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_AppleGraphicsDevicePolicy     KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_AppleGraphicsDisplayPolicy    KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_AppleGraphicsMGPUPowerControl KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_AppleGraphicsMUXControl       KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_AppleGraphicsPowerManagement  KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_AppleHDAController            KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_AppleIOAccelDisplayPipe       KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_AppleMCCSControl              KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_Count                         KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_IODisplay                     KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_Legacy                        KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_ThirdParty                    KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_VendorAMD                     KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_VendorIntel                   KIOFBNotifyGroupID = 0
	KIOFBNotifyGroupID_VendorNVIDIA                  KIOFBNotifyGroupID = 0
)

func (e KIOFBNotifyGroupID) String() string {
	switch e {
	case KIOFBNotifyGroupID_AppleGraphicsControl:
		return "KIOFBNotifyGroupID_AppleGraphicsControl"
	default:
		return fmt.Sprintf("KIOFBNotifyGroupID(%d)", e)
	}
}

type KIOFBNotifyPriority uint

const (
	KIOFBNotifyPriority_Max KIOFBNotifyPriority = 0
)

func (e KIOFBNotifyPriority) String() string {
	switch e {
	case KIOFBNotifyPriority_Max:
		return "KIOFBNotifyPriority_Max"
	default:
		return fmt.Sprintf("KIOFBNotifyPriority(%d)", e)
	}
}

type KIOFBSystem uint

const (
	KIOFBSystemAperture KIOFBSystem = 0
)

func (e KIOFBSystem) String() string {
	switch e {
	case KIOFBSystemAperture:
		return "KIOFBSystemAperture"
	default:
		return fmt.Sprintf("KIOFBSystem(%d)", e)
	}
}

type KIOFBUserRequest uint

const (
	KIOFBUserRequestProbe KIOFBUserRequest = 0
)

func (e KIOFBUserRequest) String() string {
	switch e {
	case KIOFBUserRequestProbe:
		return "KIOFBUserRequestProbe"
	default:
		return fmt.Sprintf("KIOFBUserRequest(%d)", e)
	}
}

type KIOFW uint

const (
	// KIOFWDisableAllPhysicalAccess: # Discussion
	KIOFWDisableAllPhysicalAccess KIOFW = 0
	// KIOFWDisablePhyOnSleep: # Discussion
	KIOFWDisablePhyOnSleep KIOFW = 0
	// KIOFWDisablePhysicalAccess: # Discussion
	KIOFWDisablePhysicalAccess KIOFW = 0
	// KIOFWEnableRetryOnAckD: # Discussion
	KIOFWEnableRetryOnAckD KIOFW = 0
	// KIOFWLimitAsyncPacketSize: # Discussion
	KIOFWLimitAsyncPacketSize KIOFW = 0
	// KIOFWMustBeRoot: # Discussion
	KIOFWMustBeRoot KIOFW = 0
	// KIOFWMustHaveGap63: # Discussion
	KIOFWMustHaveGap63 KIOFW = 0
	// KIOFWMustNotBeRoot: # Discussion
	KIOFWMustNotBeRoot KIOFW = 0
	KIOFWSWVers_KPF    KIOFW = 0
	KIOFWSpecID_AAPL   KIOFW = 0
)

func (e KIOFW) String() string {
	switch e {
	case KIOFWDisableAllPhysicalAccess:
		return "KIOFWDisableAllPhysicalAccess"
	default:
		return fmt.Sprintf("KIOFW(%d)", e)
	}
}

type KIOFWAVCProtocolUserClient uint

const (
	KIOFWAVCProtocolUserClientAVCRequestNotHandled       KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientAddSubunit                 KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientAllocateInputPlug          KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientAllocateOutputPlug         KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientConnectTargetPlugs         KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientDisconnectTargetPlugs      KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientFreeInputPlug              KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientFreeOutputPlug             KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientGetSubunitPlugSignalFormat KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientGetTargetPlugConnection    KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientInstallAVCCommandHandler   KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientNumAsyncCommands           KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientNumCommands                KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientPublishAVCUnitDirectory    KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientReadInputMasterPlug        KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientReadInputPlug              KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientReadOutputMasterPlug       KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientReadOutputPlug             KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientSendAVCResponse            KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientSetAVCRequestCallback      KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientSetSubunitPlugSignalFormat KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientUpdateInputMasterPlug      KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientUpdateInputPlug            KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientUpdateOutputMasterPlug     KIOFWAVCProtocolUserClient = 0
	KIOFWAVCProtocolUserClientUpdateOutputPlug           KIOFWAVCProtocolUserClient = 0
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
	KIOFWAVCSubunitPlugMsgConnected             KIOFWAVCSubunitPlugMsg = 0
	KIOFWAVCSubunitPlugMsgConnectedPlugModified KIOFWAVCSubunitPlugMsg = 0
	KIOFWAVCSubunitPlugMsgDisconnected          KIOFWAVCSubunitPlugMsg = 0
	KIOFWAVCSubunitPlugMsgSignalFormatModified  KIOFWAVCSubunitPlugMsg = 0
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
	KIOFWAVCUserClientAVCCommand                     KIOFWAVCUserClient = 0
	KIOFWAVCUserClientAVCCommandInGen                KIOFWAVCUserClient = 0
	KIOFWAVCUserClientBreakP2PInputConnection        KIOFWAVCUserClient = 0
	KIOFWAVCUserClientBreakP2POutputConnection       KIOFWAVCUserClient = 0
	KIOFWAVCUserClientCancelAsyncAVCCommand          KIOFWAVCUserClient = 0
	KIOFWAVCUserClientClose                          KIOFWAVCUserClient = 0
	KIOFWAVCUserClientCreateAsyncAVCCommand          KIOFWAVCUserClient = 0
	KIOFWAVCUserClientGetSessionRef                  KIOFWAVCUserClient = 0
	KIOFWAVCUserClientInstallAsyncAVCCommandCallback KIOFWAVCUserClient = 0
	KIOFWAVCUserClientMakeP2PInputConnection         KIOFWAVCUserClient = 0
	KIOFWAVCUserClientMakeP2POutputConnection        KIOFWAVCUserClient = 0
	KIOFWAVCUserClientNumAsyncCommands               KIOFWAVCUserClient = 0
	KIOFWAVCUserClientNumCommands                    KIOFWAVCUserClient = 0
	KIOFWAVCUserClientOpen                           KIOFWAVCUserClient = 0
	KIOFWAVCUserClientOpenWithSessionRef             KIOFWAVCUserClient = 0
	KIOFWAVCUserClientReinitAsyncAVCCommand          KIOFWAVCUserClient = 0
	KIOFWAVCUserClientReleaseAsyncAVCCommand         KIOFWAVCUserClient = 0
	KIOFWAVCUserClientSubmitAsyncAVCCommand          KIOFWAVCUserClient = 0
	KIOFWAVCUserClientUpdateAVCCommandTimeout        KIOFWAVCUserClient = 0
)

func (e KIOFWAVCUserClient) String() string {
	switch e {
	case KIOFWAVCUserClientAVCCommand:
		return "KIOFWAVCUserClientAVCCommand"
	default:
		return fmt.Sprintf("KIOFWAVCUserClient(%d)", e)
	}
}

type KIOFWPhysicalAccess uint

const (
	KIOFWPhysicalAccessDisabled              KIOFWPhysicalAccess = 0
	KIOFWPhysicalAccessDisabledForGeneration KIOFWPhysicalAccess = 0
	KIOFWPhysicalAccessEnabled               KIOFWPhysicalAccess = 0
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
	KIOFWReadFlagsNone    KIOFWRead = 0
	KIOFWReadPingTime     KIOFWRead = 0
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
	KIOFWSBP2FailsOnAckBusy       KIOFWSBP2 = 0
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
	KIOFWSBP2UserClientClose                                    KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientCreateLogin                              KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientCreateMgmtORB                            KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientCreateORB                                KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientEnableUnsolicitedStatus                  KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientGetLoginID                               KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientGetMaxCommandBlockSize                   KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientGetSessionRef                            KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientLSIWorkaroundSetCommandBuffersAsRanges   KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBLSIWorkaroundSyncBuffersForInput  KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBLSIWorkaroundSyncBuffersForOutput KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBSetCommandFunction                KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBSetManageeLogin                   KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBSetManageeORB                     KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientMgmtORBSetResponseBuffer                 KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientNumCommands                              KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientOpen                                     KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientOpenWithSessionRef                       KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientReleaseCommandBuffers                    KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientReleaseLogin                             KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientReleaseMgmtORB                           KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientReleaseORB                               KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientRingDoorbell                             KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetBusyTimeoutRegisterValue              KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandBlock                          KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandBuffersAsRanges                KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandFlags                          KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandGeneration                     KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetCommandTimeout                        KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetFetchAgentWriteCompletion             KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetLoginCallback                         KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetLoginFlags                            KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetLogoutCallback                        KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetMaxORBPayloadSize                     KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetMaxPayloadSize                        KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetMessageCallback                       KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetMgmtORBCallback                       KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetORBRefCon                             KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetPassword                              KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetReconnectTime                         KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetStatusNotify                          KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetToDummy                               KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSetUnsolicitedStatusNotify               KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitFetchAgentReset                    KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitLogin                              KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitLogout                             KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitMgmtORB                            KIOFWSBP2UserClient = 0
	KIOFWSBP2UserClientSubmitORB                                KIOFWSBP2UserClient = 0
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
	KIOFWSecurityModeNormal          KIOFWSecurityMode = 0
	KIOFWSecurityModeSecure          KIOFWSecurityMode = 0
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
	KIOFWWriteBlockRequest        KIOFWWrite = 0
	KIOFWWriteFastRetryOnBusy     KIOFWWrite = 0
	KIOFWWriteFlagsDeferredNotify KIOFWWrite = 0
	KIOFWWriteFlagsNone           KIOFWWrite = 0
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
	KIOHIDActivityUserIdle  KIOHID = 0
	KIOHIDCapsLockState     KIOHID = 0
	KIOHIDNumLockState      KIOHID = 0
)

func (e KIOHID) String() string {
	switch e {
	case KIOHIDActivityDisplayOn:
		return "KIOHIDActivityDisplayOn"
	default:
		return fmt.Sprintf("KIOHID(%d)", e)
	}
}

type KIOHIDAccelerationAlgorithmType uint

const (
	KIOHIDAccelerationAlgorithmTypeDefault    KIOHIDAccelerationAlgorithmType = 0
	KIOHIDAccelerationAlgorithmTypeParametric KIOHIDAccelerationAlgorithmType = 0
	KIOHIDAccelerationAlgorithmTypeTable      KIOHIDAccelerationAlgorithmType = 0
)

func (e KIOHIDAccelerationAlgorithmType) String() string {
	switch e {
	case KIOHIDAccelerationAlgorithmTypeDefault:
		return "KIOHIDAccelerationAlgorithmTypeDefault"
	default:
		return fmt.Sprintf("KIOHIDAccelerationAlgorithmType(%d)", e)
	}
}

type KIOHIDButtonMode uint

const (
	KIOHIDButtonMode_BothLeftClicks         KIOHIDButtonMode = 0
	KIOHIDButtonMode_EnableRightClick       KIOHIDButtonMode = 0
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

type KIOHIDElementFlagsBufferedByte uint

const (
	KIOHIDElementFlagsBufferedByteMask KIOHIDElementFlagsBufferedByte = 0
)

func (e KIOHIDElementFlagsBufferedByte) String() string {
	switch e {
	case KIOHIDElementFlagsBufferedByteMask:
		return "KIOHIDElementFlagsBufferedByteMask"
	default:
		return fmt.Sprintf("KIOHIDElementFlagsBufferedByte(%d)", e)
	}
}

type KIOHIDEvent uint

const (
	KIOHIDEventNotification KIOHIDEvent = 0
)

func (e KIOHIDEvent) String() string {
	switch e {
	case KIOHIDEventNotification:
		return "KIOHIDEventNotification"
	default:
		return fmt.Sprintf("KIOHIDEvent(%d)", e)
	}
}

type KIOHIDEventQueueType uint

const (
	KIOHIDEventQueueTypeKernel KIOHIDEventQueueType = 0
)

func (e KIOHIDEventQueueType) String() string {
	switch e {
	case KIOHIDEventQueueTypeKernel:
		return "KIOHIDEventQueueTypeKernel"
	default:
		return fmt.Sprintf("KIOHIDEventQueueType(%d)", e)
	}
}

type KIOHIDEventSystemConnect uint

const (
	KIOHIDEventSystemConnectType KIOHIDEventSystemConnect = 0
)

func (e KIOHIDEventSystemConnect) String() string {
	switch e {
	case KIOHIDEventSystemConnectType:
		return "KIOHIDEventSystemConnectType"
	default:
		return fmt.Sprintf("KIOHIDEventSystemConnect(%d)", e)
	}
}

type KIOHIDGlobal uint

const (
	KIOHIDGlobalMemory KIOHIDGlobal = 0
)

func (e KIOHIDGlobal) String() string {
	switch e {
	case KIOHIDGlobalMemory:
		return "KIOHIDGlobalMemory"
	default:
		return fmt.Sprintf("KIOHIDGlobal(%d)", e)
	}
}

type KIOHIDKeyboardPhysicalLayout uint

const (
	KIOHIDKeyboardPhysicalLayoutType101     KIOHIDKeyboardPhysicalLayout = 0
	KIOHIDKeyboardPhysicalLayoutType102     KIOHIDKeyboardPhysicalLayout = 0
	KIOHIDKeyboardPhysicalLayoutType103     KIOHIDKeyboardPhysicalLayout = 0
	KIOHIDKeyboardPhysicalLayoutType104     KIOHIDKeyboardPhysicalLayout = 0
	KIOHIDKeyboardPhysicalLayoutType106     KIOHIDKeyboardPhysicalLayout = 0
	KIOHIDKeyboardPhysicalLayoutTypeUnknown KIOHIDKeyboardPhysicalLayout = 0
	KIOHIDKeyboardPhysicalLayoutTypeVendor  KIOHIDKeyboardPhysicalLayout = 0
)

func (e KIOHIDKeyboardPhysicalLayout) String() string {
	switch e {
	case KIOHIDKeyboardPhysicalLayoutType101:
		return "KIOHIDKeyboardPhysicalLayoutType101"
	default:
		return fmt.Sprintf("KIOHIDKeyboardPhysicalLayout(%d)", e)
	}
}

type KIOHIDOpenedByEvent uint

const (
	KIOHIDOpenedByEventSystem KIOHIDOpenedByEvent = 0
)

func (e KIOHIDOpenedByEvent) String() string {
	switch e {
	case KIOHIDOpenedByEventSystem:
		return "KIOHIDOpenedByEventSystem"
	default:
		return fmt.Sprintf("KIOHIDOpenedByEvent(%d)", e)
	}
}

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

type KIOHIDQueueOptionsType uint

const (
	// KIOHIDQueueOptionsTypeEnqueueAll: # Discussion
	KIOHIDQueueOptionsTypeEnqueueAll KIOHIDQueueOptionsType = 0
	// KIOHIDQueueOptionsTypeNone: # Discussion
	KIOHIDQueueOptionsTypeNone KIOHIDQueueOptionsType = 0
)

func (e KIOHIDQueueOptionsType) String() string {
	switch e {
	case KIOHIDQueueOptionsTypeEnqueueAll:
		return "KIOHIDQueueOptionsTypeEnqueueAll"
	default:
		return fmt.Sprintf("KIOHIDQueueOptionsType(%d)", e)
	}
}

type KIOHIDStandardType uint

const (
	// KIOHIDStandardTypeJIS: # Discussion
	KIOHIDStandardTypeJIS         KIOHIDStandardType = 0
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
	KIOHIDValueOptionsFlagPrevious       KIOHIDValueOptionsFlag = 0
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
	KIOHIDValueScaleTypeExponent   KIOHIDValueScaleType = 0
)

func (e KIOHIDValueScaleType) String() string {
	switch e {
	case KIOHIDValueScaleTypeCalibrated:
		return "KIOHIDValueScaleTypeCalibrated"
	default:
		return fmt.Sprintf("KIOHIDValueScaleType(%d)", e)
	}
}

type KIOHibernatePreview uint

const (
	KIOHibernatePreviewActive  KIOHibernatePreview = 0
	KIOHibernatePreviewUpdates KIOHibernatePreview = 0
)

func (e KIOHibernatePreview) String() string {
	switch e {
	case KIOHibernatePreviewActive:
		return "KIOHibernatePreviewActive"
	default:
		return fmt.Sprintf("KIOHibernatePreview(%d)", e)
	}
}

type KIOInterruptDispatchSourceType uint

const (
	KIOInterruptDispatchSourceTypeEdge KIOInterruptDispatchSourceType = 0
)

func (e KIOInterruptDispatchSourceType) String() string {
	switch e {
	case KIOInterruptDispatchSourceTypeEdge:
		return "KIOInterruptDispatchSourceTypeEdge"
	default:
		return fmt.Sprintf("KIOInterruptDispatchSourceType(%d)", e)
	}
}

type KIOInterruptSourceAbsolute uint

const (
	KIOInterruptSourceAbsoluteTime KIOInterruptSourceAbsolute = 0
)

func (e KIOInterruptSourceAbsolute) String() string {
	switch e {
	case KIOInterruptSourceAbsoluteTime:
		return "KIOInterruptSourceAbsoluteTime"
	default:
		return fmt.Sprintf("KIOInterruptSourceAbsolute(%d)", e)
	}
}

type KIOKitDebugUser uint

const (
	KIOKitDebugUserOptions KIOKitDebugUser = 0
)

func (e KIOKitDebugUser) String() string {
	switch e {
	case KIOKitDebugUserOptions:
		return "KIOKitDebugUserOptions"
	default:
		return fmt.Sprintf("KIOKitDebugUser(%d)", e)
	}
}

type KIOKitDiagnosticsClient uint

const (
	KIOKitDiagnosticsClientType KIOKitDiagnosticsClient = 0
)

func (e KIOKitDiagnosticsClient) String() string {
	switch e {
	case KIOKitDiagnosticsClientType:
		return "KIOKitDiagnosticsClientType"
	default:
		return fmt.Sprintf("KIOKitDiagnosticsClient(%d)", e)
	}
}

type KIOLockState uint

const (
	KIOLockStateLocked KIOLockState = 0
)

func (e KIOLockState) String() string {
	switch e {
	case KIOLockStateLocked:
		return "KIOLockStateLocked"
	default:
		return fmt.Sprintf("KIOLockState(%d)", e)
	}
}

type KIOMap uint

const (
	KIOMapAnywhere    KIOMap = 0
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

type KIOMapper uint

const (
	KIOMapperUncached KIOMapper = 0
)

func (e KIOMapper) String() string {
	switch e {
	case KIOMapperUncached:
		return "KIOMapperUncached"
	default:
		return fmt.Sprintf("KIOMapper(%d)", e)
	}
}

type KIOMaxBus uint

const (
	KIOMaxBusStall10usec KIOMaxBus = 0
	KIOMaxBusStall20usec KIOMaxBus = 0
	KIOMaxBusStall25usec KIOMaxBus = 0
	KIOMaxBusStall30usec KIOMaxBus = 0
	KIOMaxBusStall40usec KIOMaxBus = 0
	KIOMaxBusStall5usec  KIOMaxBus = 0
	KIOMaxBusStallNone   KIOMaxBus = 0
)

func (e KIOMaxBus) String() string {
	switch e {
	case KIOMaxBusStall10usec:
		return "KIOMaxBusStall10usec"
	default:
		return fmt.Sprintf("KIOMaxBus(%d)", e)
	}
}

type KIOMaxPixel uint

const (
	KIOMaxPixelBits KIOMaxPixel = 0
)

func (e KIOMaxPixel) String() string {
	switch e {
	case KIOMaxPixelBits:
		return "KIOMaxPixelBits"
	default:
		return fmt.Sprintf("KIOMaxPixel(%d)", e)
	}
}

type KIOMediaAttribute uint

const (
	// KIOMediaAttributeEjectableMask: # Discussion
	KIOMediaAttributeEjectableMask KIOMediaAttribute = 0
	// KIOMediaAttributeRemovableMask: # Discussion
	KIOMediaAttributeRemovableMask KIOMediaAttribute = 0
	KIOMediaAttributeReservedMask  KIOMediaAttribute = 0
)

func (e KIOMediaAttribute) String() string {
	switch e {
	case KIOMediaAttributeEjectableMask:
		return "KIOMediaAttributeEjectableMask"
	default:
		return fmt.Sprintf("KIOMediaAttribute(%d)", e)
	}
}

type KIOMediaState uint

const (
	// KIOMediaStateBusy: # Discussion
	KIOMediaStateBusy KIOMediaState = 0
)

func (e KIOMediaState) String() string {
	switch e {
	case KIOMediaStateBusy:
		return "KIOMediaStateBusy"
	default:
		return fmt.Sprintf("KIOMediaState(%d)", e)
	}
}

type KIOMediumEthernet uint

const (
	KIOMediumEthernetAuto     KIOMediumEthernet = 0
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
	KIOMediumIEEE80211            KIOMediumIEE = 0
	KIOMediumIEEE80211Auto        KIOMediumIEE = 0
	KIOMediumIEEE80211DS1         KIOMediumIEE = 0
	KIOMediumIEEE80211DS11        KIOMediumIEE = 0
	KIOMediumIEEE80211DS2         KIOMediumIEE = 0
	KIOMediumIEEE80211DS5         KIOMediumIEE = 0
	KIOMediumIEEE80211FH1         KIOMediumIEE = 0
	KIOMediumIEEE80211FH2         KIOMediumIEE = 0
	KIOMediumIEEE80211Manual      KIOMediumIEE = 0
	KIOMediumIEEE80211None        KIOMediumIEE = 0
	KIOMediumIEEE80211OptionAdhoc KIOMediumIEE = 0
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
	KIOMediumOptionEEE         KIOMediumOption = 0
	KIOMediumOptionFlag0       KIOMediumOption = 0
	KIOMediumOptionFlag1       KIOMediumOption = 0
	KIOMediumOptionFlag2       KIOMediumOption = 0
	KIOMediumOptionFlowControl KIOMediumOption = 0
	KIOMediumOptionFullDuplex  KIOMediumOption = 0
	KIOMediumOptionHalfDuplex  KIOMediumOption = 0
	KIOMediumOptionLoopback    KIOMediumOption = 0
)

func (e KIOMediumOption) String() string {
	switch e {
	case KIOMediumOptionEEE:
		return "KIOMediumOptionEEE"
	default:
		return fmt.Sprintf("KIOMediumOption(%d)", e)
	}
}

type KIOMemory uint

const (
	KIOMemoryAsReference      KIOMemory = 0
	KIOMemoryBufferPageable   KIOMemory = 0
	KIOMemoryClearEncrypt     KIOMemory = 0
	KIOMemoryDirectionMask    KIOMemory = 0
	KIOMemoryHostOnly         KIOMemory = 0
	KIOMemoryMapCopyOnWrite   KIOMemory = 0
	KIOMemoryMapperNone       KIOMemory = 0
	KIOMemoryPersistent       KIOMemory = 0
	KIOMemoryRemote           KIOMemory = 0
	KIOMemoryThreadSafe       KIOMemory = 0
	KIOMemoryTypeMask         KIOMemory = 0
	KIOMemoryTypePersistentMD KIOMemory = 0
	KIOMemoryTypePhysical     KIOMemory = 0
	KIOMemoryTypePhysical64   KIOMemory = 0
	KIOMemoryTypeUIO          KIOMemory = 0
	KIOMemoryTypeUPL          KIOMemory = 0
	KIOMemoryTypeVirtual      KIOMemory = 0
	KIOMemoryTypeVirtual64    KIOMemory = 0
	KIOMemoryUseReserve       KIOMemory = 0
)

func (e KIOMemory) String() string {
	switch e {
	case KIOMemoryAsReference:
		return "KIOMemoryAsReference"
	default:
		return fmt.Sprintf("KIOMemory(%d)", e)
	}
}

type KIOMemoryClear uint

const (
	KIOMemoryClearEncrypted KIOMemoryClear = 0
)

func (e KIOMemoryClear) String() string {
	switch e {
	case KIOMemoryClearEncrypted:
		return "KIOMemoryClearEncrypted"
	default:
		return fmt.Sprintf("KIOMemoryClear(%d)", e)
	}
}

type KIOMemoryDirection uint

const (
	KIOMemoryDirectionIn   KIOMemoryDirection = 0
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

type KIOMemoryLedgerFlagNo uint

const (
	KIOMemoryLedgerFlagNoFootprint KIOMemoryLedgerFlagNo = 0
)

func (e KIOMemoryLedgerFlagNo) String() string {
	switch e {
	case KIOMemoryLedgerFlagNoFootprint:
		return "KIOMemoryLedgerFlagNoFootprint"
	default:
		return fmt.Sprintf("KIOMemoryLedgerFlagNo(%d)", e)
	}
}

type KIOMemoryMap uint

const (
	KIOMemoryMapCacheModeDefault KIOMemoryMap = 0
	KIOMemoryMapGuardedSmall     KIOMemoryMap = 0
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
	KIOMemoryPurgeableFaultOnAccess  KIOMemoryPurgeable = 0
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

type KIOMirror uint

const (
	KIOMirrorDefault KIOMirror = 0
)

func (e KIOMirror) String() string {
	switch e {
	case KIOMirrorDefault:
		return "KIOMirrorDefault"
	default:
		return fmt.Sprintf("KIOMirror(%d)", e)
	}
}

type KIOMirrorHW uint

const (
	KIOMirrorHWClipped KIOMirrorHW = 0
)

func (e KIOMirrorHW) String() string {
	switch e {
	case KIOMirrorHWClipped:
		return "KIOMirrorHWClipped"
	default:
		return fmt.Sprintf("KIOMirrorHW(%d)", e)
	}
}

type KIONDRV uint

const (
	KIONDRVAsynchronousIOCommandKind KIONDRV = 0
	KIONDRVCloseCommand              KIONDRV = 0
	KIONDRVFinalizeCommand           KIONDRV = 0
	KIONDRVImmediateIOCommandKind    KIONDRV = 0
	KIONDRVSynchronousIOCommandKind  KIONDRV = 0
)

func (e KIONDRV) String() string {
	switch e {
	case KIONDRVAsynchronousIOCommandKind:
		return "KIONDRVAsynchronousIOCommandKind"
	default:
		return fmt.Sprintf("KIONDRV(%d)", e)
	}
}

type KIONVRAMOperation uint

const (
	KIONVRAMOperationDelete     KIONVRAMOperation = 0
	KIONVRAMOperationObliterate KIONVRAMOperation = 0
	KIONVRAMOperationRead       KIONVRAMOperation = 0
	KIONVRAMOperationReset      KIONVRAMOperation = 0
	KIONVRAMOperationWrite      KIONVRAMOperation = 0
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

type KIONetworkDataBufferType uint

const (
	// KIONetworkDataBufferTypeExternal: # Discussion
	KIONetworkDataBufferTypeExternal KIONetworkDataBufferType = 0
)

func (e KIONetworkDataBufferType) String() string {
	switch e {
	case KIONetworkDataBufferTypeExternal:
		return "KIONetworkDataBufferTypeExternal"
	default:
		return fmt.Sprintf("KIONetworkDataBufferType(%d)", e)
	}
}

type KIONetworkEvent uint

const (
	KIONetworkEventTypeDLIL                KIONetworkEvent = 0
	KIONetworkEventTypeLinkDown            KIONetworkEvent = 0
	KIONetworkEventTypeLinkSpeedChange     KIONetworkEvent = 0
	KIONetworkEventTypeLinkUp              KIONetworkEvent = 0
	KIONetworkEventWakeOnLANSupportChanged KIONetworkEvent = 0
)

func (e KIONetworkEvent) String() string {
	switch e {
	case KIONetworkEventTypeDLIL:
		return "KIONetworkEventTypeDLIL"
	default:
		return fmt.Sprintf("KIONetworkEvent(%d)", e)
	}
}

type KIONetworkFeature uint

const (
	KIONetworkFeatureHWTimeStamp KIONetworkFeature = 0
	// KIONetworkFeatureHardwareVlan: # Discussion
	KIONetworkFeatureHardwareVlan KIONetworkFeature = 0
	KIONetworkFeatureLRO          KIONetworkFeature = 0
	// KIONetworkFeatureMultiPages: # Discussion
	KIONetworkFeatureMultiPages KIONetworkFeature = 0
	// KIONetworkFeatureNoBSDWait: # Discussion
	KIONetworkFeatureNoBSDWait   KIONetworkFeature = 0
	KIONetworkFeatureSWTimeStamp KIONetworkFeature = 0
	// KIONetworkFeatureSoftwareVlan: # Discussion
	KIONetworkFeatureSoftwareVlan KIONetworkFeature = 0
	// KIONetworkFeatureTSOIPv4: # Discussion
	KIONetworkFeatureTSOIPv4 KIONetworkFeature = 0
	// KIONetworkFeatureTSOIPv6: # Discussion
	KIONetworkFeatureTSOIPv6 KIONetworkFeature = 0
	// KIONetworkFeatureTransmitCompletionStatus: # Discussion
	KIONetworkFeatureTransmitCompletionStatus KIONetworkFeature = 0
)

func (e KIONetworkFeature) String() string {
	switch e {
	case KIONetworkFeatureHWTimeStamp:
		return "KIONetworkFeatureHWTimeStamp"
	default:
		return fmt.Sprintf("KIONetworkFeature(%d)", e)
	}
}

type KIONetworkInterface uint

const (
	// KIONetworkInterfaceDisabledState: # Discussion
	KIONetworkInterfaceDisabledState KIONetworkInterface = 0
	// KIONetworkInterfaceOpenedState: # Discussion
	KIONetworkInterfaceOpenedState KIONetworkInterface = 0
	// KIONetworkInterfaceRegisteredState: # Discussion
	KIONetworkInterfaceRegisteredState KIONetworkInterface = 0
)

func (e KIONetworkInterface) String() string {
	switch e {
	case KIONetworkInterfaceDisabledState:
		return "KIONetworkInterfaceDisabledState"
	default:
		return fmt.Sprintf("KIONetworkInterface(%d)", e)
	}
}

type KIONetworkLink uint

const (
	KIONetworkLinkActive          KIONetworkLink = 0
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

type KIOOutputCommand uint

const (
	// KIOOutputCommandMask: # Discussion
	KIOOutputCommandMask KIOOutputCommand = 0
	// KIOOutputCommandNone: # Discussion
	KIOOutputCommandNone KIOOutputCommand = 0
	// KIOOutputCommandStall: # Discussion
	KIOOutputCommandStall KIOOutputCommand = 0
)

func (e KIOOutputCommand) String() string {
	switch e {
	case KIOOutputCommandMask:
		return "KIOOutputCommandMask"
	default:
		return fmt.Sprintf("KIOOutputCommand(%d)", e)
	}
}

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
	KIOPCI64BitMemorySpace KIOPC = 0
	KIOPCIConfigSpace      KIOPC = 0
	KIOPCIIOSpace          KIOPC = 0
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
	KIOPCIAGP8Capability                             KIOPCI = 0
	KIOPCIAGPCapability                              KIOPCI = 0
	KIOPCICPCIHotswapCapability                      KIOPCI = 0
	KIOPCICPCIResourceControlCapability              KIOPCI = 0
	KIOPCICapabilityIDOffset                         KIOPCI = 0
	KIOPCIDebugPortCapability                        KIOPCI = 0
	KIOPCIExpressAccessControlServicesCapability     KIOPCI = 0
	KIOPCIExpressDeviceSerialNumberCapability        KIOPCI = 0
	KIOPCIExpressErrorReportingCapability            KIOPCI = 0
	KIOPCIExpressL1PMSubstatesCapability             KIOPCI = 0
	KIOPCIExpressLatencyTolerenceReportingCapability KIOPCI = 0
	KIOPCIExpressPowerBudgetCapability               KIOPCI = 0
	KIOPCIExpressVirtualChannelCapability            KIOPCI = 0
	KIOPCIFPBCapability                              KIOPCI = 0
	KIOPCIHotplugCapability                          KIOPCI = 0
	KIOPCILDTCapability                              KIOPCI = 0
	KIOPCIMSICapability                              KIOPCI = 0
	KIOPCIMSIXCapability                             KIOPCI = 0
	KIOPCINextCapabilityOffset                       KIOPCI = 0
	KIOPCIPCIExpressCapability                       KIOPCI = 0
	KIOPCIPCIXCapability                             KIOPCI = 0
	KIOPCIPowerManagementCapability                  KIOPCI = 0
	KIOPCISecureCapability                           KIOPCI = 0
	KIOPCISlotIDCapability                           KIOPCI = 0
	KIOPCIVendorSpecificCapability                   KIOPCI = 0
	KIOPCIVitalProductDataCapability                 KIOPCI = 0
)

func (e KIOPCI) String() string {
	switch e {
	case KIOPCIAGP8Capability:
		return "KIOPCIAGP8Capability"
	default:
		return fmt.Sprintf("KIOPCI(%d)", e)
	}
}

type KIOPCICommand uint

const (
	KIOPCICommandAddressStepping  KIOPCICommand = 0
	KIOPCICommandFastBack2Back    KIOPCICommand = 0
	KIOPCICommandIOSpace          KIOPCICommand = 0
	KIOPCICommandInterruptDisable KIOPCICommand = 0
	KIOPCICommandMemWrInvalidate  KIOPCICommand = 0
	KIOPCICommandMemorySpace      KIOPCICommand = 0
	KIOPCICommandPaletteSnoop     KIOPCICommand = 0
	KIOPCICommandParityError      KIOPCICommand = 0
	KIOPCICommandSERR             KIOPCICommand = 0
	KIOPCICommandSpecialCycles    KIOPCICommand = 0
	// Deprecated.
	KIOPCICommandBusMaster KIOPCICommand = 0
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
	KIOPCIConfigBIST              KIOPCIConfig = 0
	KIOPCIConfigBaseAddress0      KIOPCIConfig = 0
	KIOPCIConfigBaseAddress1      KIOPCIConfig = 0
	KIOPCIConfigBaseAddress2      KIOPCIConfig = 0
	KIOPCIConfigBaseAddress3      KIOPCIConfig = 0
	KIOPCIConfigBaseAddress4      KIOPCIConfig = 0
	KIOPCIConfigBaseAddress5      KIOPCIConfig = 0
	KIOPCIConfigCacheLineSize     KIOPCIConfig = 0
	KIOPCIConfigCapabilitiesPtr   KIOPCIConfig = 0
	KIOPCIConfigCardBusCISPtr     KIOPCIConfig = 0
	KIOPCIConfigClassCode         KIOPCIConfig = 0
	KIOPCIConfigCommand           KIOPCIConfig = 0
	KIOPCIConfigDeviceID          KIOPCIConfig = 0
	KIOPCIConfigExpansionROMBase  KIOPCIConfig = 0
	KIOPCIConfigHeaderType        KIOPCIConfig = 0
	KIOPCIConfigInterruptLine     KIOPCIConfig = 0
	KIOPCIConfigInterruptPin      KIOPCIConfig = 0
	KIOPCIConfigLatencyTimer      KIOPCIConfig = 0
	KIOPCIConfigMaximumLatency    KIOPCIConfig = 0
	KIOPCIConfigMinimumGrant      KIOPCIConfig = 0
	KIOPCIConfigRevisionID        KIOPCIConfig = 0
	KIOPCIConfigStatus            KIOPCIConfig = 0
	KIOPCIConfigSubSystemID       KIOPCIConfig = 0
	KIOPCIConfigSubSystemVendorID KIOPCIConfig = 0
	KIOPCIConfigVendorID          KIOPCIConfig = 0
)

func (e KIOPCIConfig) String() string {
	switch e {
	case KIOPCIConfigBIST:
		return "KIOPCIConfigBIST"
	default:
		return fmt.Sprintf("KIOPCIConfig(%d)", e)
	}
}

type KIOPCIConfigAGP uint

const (
	KIOPCIConfigAGPCommandOffset KIOPCIConfigAGP = 0
	KIOPCIConfigAGPStatusOffset  KIOPCIConfigAGP = 0
)

func (e KIOPCIConfigAGP) String() string {
	switch e {
	case KIOPCIConfigAGPCommandOffset:
		return "KIOPCIConfigAGPCommandOffset"
	default:
		return fmt.Sprintf("KIOPCIConfigAGP(%d)", e)
	}
}

type KIOPCIConfigShadow uint

const (
	KIOPCIConfigShadowPermanent KIOPCIConfigShadow = 0
)

func (e KIOPCIConfigShadow) String() string {
	switch e {
	case KIOPCIConfigShadowPermanent:
		return "KIOPCIConfigShadowPermanent"
	default:
		return fmt.Sprintf("KIOPCIConfigShadow(%d)", e)
	}
}

type KIOPCICorrectableErrorBit int

const (
	KIOPCICorrectableErrorBitAdvisoryNonFatal   KIOPCICorrectableErrorBit = 0
	KIOPCICorrectableErrorBitBadDLLP            KIOPCICorrectableErrorBit = 0
	KIOPCICorrectableErrorBitBadTLP             KIOPCICorrectableErrorBit = 0
	KIOPCICorrectableErrorBitCorrectedInternal  KIOPCICorrectableErrorBit = 0
	KIOPCICorrectableErrorBitHeaderLogOverflow  KIOPCICorrectableErrorBit = 0
	KIOPCICorrectableErrorBitReceiver           KIOPCICorrectableErrorBit = 0
	KIOPCICorrectableErrorBitReplayNumRollover  KIOPCICorrectableErrorBit = 0
	KIOPCICorrectableErrorBitReplayTimerTimeout KIOPCICorrectableErrorBit = 0
)

func (e KIOPCICorrectableErrorBit) String() string {
	switch e {
	case KIOPCICorrectableErrorBitAdvisoryNonFatal:
		return "KIOPCICorrectableErrorBitAdvisoryNonFatal"
	default:
		return fmt.Sprintf("KIOPCICorrectableErrorBit(%d)", e)
	}
}

type KIOPCIDevice uint

const (
	KIOPCIDeviceDozeState       KIOPCIDevice = 0
	KIOPCIDeviceOffState        KIOPCIDevice = 0
	KIOPCIDeviceOnState         KIOPCIDevice = 0
	KIOPCIDevicePausedState     KIOPCIDevice = 0
	KIOPCIDevicePowerStateCount KIOPCIDevice = 0
)

func (e KIOPCIDevice) String() string {
	switch e {
	case KIOPCIDeviceDozeState:
		return "KIOPCIDeviceDozeState"
	default:
		return fmt.Sprintf("KIOPCIDevice(%d)", e)
	}
}

type KIOPCIDeviceResetOption uint

const (
	KIOPCIDeviceResetOptionNone      KIOPCIDeviceResetOption = 0
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
	KIOPCIDeviceResetTypeFunctionReset    KIOPCIDeviceResetType = 0
	KIOPCIDeviceResetTypeHotReset         KIOPCIDeviceResetType = 0
	KIOPCIDeviceResetTypeWarmReset        KIOPCIDeviceResetType = 0
	KIOPCIDeviceResetTypeWarmResetDisable KIOPCIDeviceResetType = 0
	KIOPCIDeviceResetTypeWarmResetEnable  KIOPCIDeviceResetType = 0
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
	KIOPCIEventNonFatalError    KIOPCIEvent = 0
)

func (e KIOPCIEvent) String() string {
	switch e {
	case KIOPCIEventLinkEnableChange:
		return "KIOPCIEventLinkEnableChange"
	default:
		return fmt.Sprintf("KIOPCIEvent(%d)", e)
	}
}

type KIOPCILatency uint

const (
	KIOPCILatencySnooped   KIOPCILatency = 0
	KIOPCILatencyUnsnooped KIOPCILatency = 0
)

func (e KIOPCILatency) String() string {
	switch e {
	case KIOPCILatencySnooped:
		return "KIOPCILatencySnooped"
	default:
		return fmt.Sprintf("KIOPCILatency(%d)", e)
	}
}

type KIOPCILinkSpeed uint

const (
	KIOPCILinkSpeed_16_GTs  KIOPCILinkSpeed = 0
	KIOPCILinkSpeed_2_5_GTs KIOPCILinkSpeed = 0
	KIOPCILinkSpeed_32_GTs  KIOPCILinkSpeed = 0
	KIOPCILinkSpeed_5_GTs   KIOPCILinkSpeed = 0
	KIOPCILinkSpeed_8_GTs   KIOPCILinkSpeed = 0
)

func (e KIOPCILinkSpeed) String() string {
	switch e {
	case KIOPCILinkSpeed_16_GTs:
		return "KIOPCILinkSpeed_16_GTs"
	default:
		return fmt.Sprintf("KIOPCILinkSpeed(%d)", e)
	}
}

type KIOPCIProbeOption uint

const (
	KIOPCIProbeOptionDone KIOPCIProbeOption = 0
)

func (e KIOPCIProbeOption) String() string {
	switch e {
	case KIOPCIProbeOptionDone:
		return "KIOPCIProbeOptionDone"
	default:
		return fmt.Sprintf("KIOPCIProbeOption(%d)", e)
	}
}

type KIOPCIReset uint

const (
	KIOPCIResetHot KIOPCIReset = 0
)

func (e KIOPCIReset) String() string {
	switch e {
	case KIOPCIResetHot:
		return "KIOPCIResetHot"
	default:
		return fmt.Sprintf("KIOPCIReset(%d)", e)
	}
}

type KIOPCIStatus int

const (
	KIOPCIStatusCapabilities KIOPCIStatus = 0
	KIOPCIStatusDevSel1      KIOPCIStatus = 0
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
	KIOPCIUncorrectableErrorBitCompletionTimeout  KIOPCIUncorrectableErrorBit = 0
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
	KIOPMACInstalled          KIOPM = 0
	KIOPMACnoChargeCapability KIOPM = 0
	KIOPMAllowSleep           KIOPM = 0
	KIOPMAuxPowerOn           KIOPM = 0
	KIOPMBatteryAtWarn        KIOPM = 0
	KIOPMBatteryCharging      KIOPM = 0
	KIOPMBatteryDepleted      KIOPM = 0
	KIOPMBatteryInstalled     KIOPM = 0
	KIOPMChildClamp           KIOPM = 0
	KIOPMChildClamp2          KIOPM = 0
	KIOPMClamshellOpened      KIOPM = 0
	KIOPMClamshellStateOnWake KIOPM = 0
	KIOPMClockNormal          KIOPM = 0
	KIOPMClockRunning         KIOPM = 0
	KIOPMClosedClamshell      KIOPM = 0
	KIOPMConfigRetained       KIOPM = 0
	// KIOPMDeviceUsable: # Discussion
	KIOPMDeviceUsable  KIOPM = 0
	KIOPMDoze          KIOPM = 0
	KIOPMFairValue     KIOPM = 0
	KIOPMForceLowSpeed KIOPM = 0
	KIOPMHighestState  KIOPM = 0
	// KIOPMInitialDeviceState: # Discussion
	KIOPMInitialDeviceState KIOPM = 0
	// KIOPMLowPower: # Discussion
	KIOPMLowPower        KIOPM = 0
	KIOPMLowestState     KIOPM = 0
	KIOPMNextHigherState KIOPM = 0
	KIOPMNextLowerState  KIOPM = 0
	KIOPMNotPowerManaged KIOPM = 0
	// KIOPMPowerOn: # Discussion
	KIOPMPowerOn KIOPM = 0
	// KIOPMPreventIdleSleep: # Discussion
	KIOPMPreventIdleSleep   KIOPM = 0
	KIOPMPreventSystemSleep KIOPM = 0
	KIOPMRawLowBattery      KIOPM = 0
	// KIOPMRestart: # Discussion
	KIOPMRestart KIOPM = 0
	// KIOPMRestartCapability: # Discussion
	KIOPMRestartCapability KIOPM = 0
	// KIOPMRootDomainState: # Discussion
	KIOPMRootDomainState KIOPM = 0
	// KIOPMSleep: # Discussion
	KIOPMSleep KIOPM = 0
	// KIOPMSleepCapability: # Discussion
	KIOPMSleepCapability KIOPM = 0
	KIOPMUPSInstalled    KIOPM = 0
	KIOPMUndefinedValue  KIOPM = 0
	KIOPMUnknown         KIOPM = 0
)

func (e KIOPM) String() string {
	switch e {
	case KIOPMACInstalled:
		return "KIOPMACInstalled"
	default:
		return fmt.Sprintf("KIOPM(%d)", e)
	}
}

type KIOPMBroadcast uint

const (
	KIOPMBroadcastAggressiveness KIOPMBroadcast = 0
)

func (e KIOPMBroadcast) String() string {
	switch e {
	case KIOPMBroadcastAggressiveness:
		return "KIOPMBroadcastAggressiveness"
	default:
		return fmt.Sprintf("KIOPMBroadcast(%d)", e)
	}
}

type KIOPMDriverAssertion uint

const (
	KIOPMDriverAssertionBluetoothHIDDevicePairedBit KIOPMDriverAssertion = 0
	KIOPMDriverAssertionCPUBit                      KIOPMDriverAssertion = 0
	KIOPMDriverAssertionExternalMediaMountedBit     KIOPMDriverAssertion = 0
	KIOPMDriverAssertionMagicPacketWakeEnabledBit   KIOPMDriverAssertion = 0
	KIOPMDriverAssertionNetworkKeepAliveActiveBit   KIOPMDriverAssertion = 0
	KIOPMDriverAssertionPreventDisplaySleepBit      KIOPMDriverAssertion = 0
	KIOPMDriverAssertionPreventSystemIdleSleepBit   KIOPMDriverAssertion = 0
	KIOPMDriverAssertionReservedBit5                KIOPMDriverAssertion = 0
	KIOPMDriverAssertionReservedBit7                KIOPMDriverAssertion = 0
	KIOPMDriverAssertionUSBExternalDeviceBit        KIOPMDriverAssertion = 0
)

func (e KIOPMDriverAssertion) String() string {
	switch e {
	case KIOPMDriverAssertionBluetoothHIDDevicePairedBit:
		return "KIOPMDriverAssertionBluetoothHIDDevicePairedBit"
	default:
		return fmt.Sprintf("KIOPMDriverAssertion(%d)", e)
	}
}

type KIOPMExternal uint

const (
	KIOPMExternalPower KIOPMExternal = 0
)

func (e KIOPMExternal) String() string {
	switch e {
	case KIOPMExternalPower:
		return "KIOPMExternalPower"
	default:
		return fmt.Sprintf("KIOPMExternal(%d)", e)
	}
}

type KIOPMNoErr int

const (
	IOPMAckImplied        KIOPMNoErr = 0
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

type KIOPMPSLocation uint

const (
	KIOPMPSLocationLeft KIOPMPSLocation = 0
)

func (e KIOPMPSLocation) String() string {
	switch e {
	case KIOPMPSLocationLeft:
		return "KIOPMPSLocationLeft"
	default:
		return fmt.Sprintf("KIOPMPSLocation(%d)", e)
	}
}

type KIOPMPowerState uint

const (
	KIOPMPowerStateVersion1 KIOPMPowerState = 0
	KIOPMPowerStateVersion2 KIOPMPowerState = 0
)

func (e KIOPMPowerState) String() string {
	switch e {
	case KIOPMPowerStateVersion1:
		return "KIOPMPowerStateVersion1"
	default:
		return fmt.Sprintf("KIOPMPowerState(%d)", e)
	}
}

type KIOPMSupportedOn uint

const (
	KIOPMSupportedOnAC   KIOPMSupportedOn = 0
	KIOPMSupportedOnBatt KIOPMSupportedOn = 0
	KIOPMSupportedOnUPS  KIOPMSupportedOn = 0
)

func (e KIOPMSupportedOn) String() string {
	switch e {
	case KIOPMSupportedOnAC:
		return "KIOPMSupportedOnAC"
	default:
		return fmt.Sprintf("KIOPMSupportedOn(%d)", e)
	}
}

type KIOPMSystemCapability uint

const (
	KIOPMSystemCapabilityAudio    KIOPMSystemCapability = 0
	KIOPMSystemCapabilityCPU      KIOPMSystemCapability = 0
	KIOPMSystemCapabilityGraphics KIOPMSystemCapability = 0
	KIOPMSystemCapabilityNetwork  KIOPMSystemCapability = 0
)

func (e KIOPMSystemCapability) String() string {
	switch e {
	case KIOPMSystemCapabilityAudio:
		return "KIOPMSystemCapabilityAudio"
	default:
		return fmt.Sprintf("KIOPMSystemCapability(%d)", e)
	}
}

type KIOPMSystemCapabilityDid uint

const (
	// KIOPMSystemCapabilityDidChange: # Discussion
	KIOPMSystemCapabilityDidChange KIOPMSystemCapabilityDid = 0
)

func (e KIOPMSystemCapabilityDid) String() string {
	switch e {
	case KIOPMSystemCapabilityDidChange:
		return "KIOPMSystemCapabilityDidChange"
	default:
		return fmt.Sprintf("KIOPMSystemCapabilityDid(%d)", e)
	}
}

type KIOPMThermalLevel uint

const (
	KIOPMThermalLevelCritical KIOPMThermalLevel = 0
	KIOPMThermalLevelDanger   KIOPMThermalLevel = 0
	KIOPMThermalLevelNormal   KIOPMThermalLevel = 0
	KIOPMThermalLevelTrap     KIOPMThermalLevel = 0
	KIOPMThermalLevelUnknown  KIOPMThermalLevel = 0
	KIOPMThermalLevelWarning  KIOPMThermalLevel = 0
)

func (e KIOPMThermalLevel) String() string {
	switch e {
	case KIOPMThermalLevelCritical:
		return "KIOPMThermalLevelCritical"
	default:
		return fmt.Sprintf("KIOPMThermalLevel(%d)", e)
	}
}

type KIOPSAdapterErrorFlagDeviceNeedsToBe int

const (
	KIOPSAdapterErrorFlagDeviceNeedsToBeRepositioned KIOPSAdapterErrorFlagDeviceNeedsToBe = 0
)

func (e KIOPSAdapterErrorFlagDeviceNeedsToBe) String() string {
	switch e {
	case KIOPSAdapterErrorFlagDeviceNeedsToBeRepositioned:
		return "KIOPSAdapterErrorFlagDeviceNeedsToBeRepositioned"
	default:
		return fmt.Sprintf("KIOPSAdapterErrorFlagDeviceNeedsToBe(%d)", e)
	}
}

type KIOPSFamilyCode uint

const (
	KIOPSFamilyCodeAC                        KIOPSFamilyCode = 0
	KIOPSFamilyCodeDisconnected              KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal                  KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal2                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal3                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal4                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal5                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal6                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal7                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeExternal8                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeFirewire                  KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBAdapter                KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBCBrick                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBCPD                    KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBCTypeC                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBChargingPort           KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBChargingPortDedicated  KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBChargingPortDownstream KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBDevice                 KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBHost                   KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBHostSuspended          KIOPSFamilyCode = 0
	KIOPSFamilyCodeUSBUnknown                KIOPSFamilyCode = 0
	KIOPSFamilyCodeUnsupported               KIOPSFamilyCode = 0
	KIOPSFamilyCodeUnsupportedRegion         KIOPSFamilyCode = 0
)

func (e KIOPSFamilyCode) String() string {
	switch e {
	case KIOPSFamilyCodeAC:
		return "KIOPSFamilyCodeAC"
	default:
		return fmt.Sprintf("KIOPSFamilyCode(%d)", e)
	}
}

type KIOPacketBuffer uint

const (
	KIOPacketBufferAlign1 KIOPacketBuffer = 0
)

func (e KIOPacketBuffer) String() string {
	switch e {
	case KIOPacketBufferAlign1:
		return "KIOPacketBufferAlign1"
	default:
		return fmt.Sprintf("KIOPacketBuffer(%d)", e)
	}
}

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
	KIOPixelEncodingRGB444       KIOPixelEncoding = 0
	KIOPixelEncodingYCbCr420     KIOPixelEncoding = 0
	KIOPixelEncodingYCbCr422     KIOPixelEncoding = 0
	KIOPixelEncodingYCbCr444     KIOPixelEncoding = 0
)

func (e KIOPixelEncoding) String() string {
	switch e {
	case KIOPixelEncodingNotSupported:
		return "KIOPixelEncodingNotSupported"
	default:
		return fmt.Sprintf("KIOPixelEncoding(%d)", e)
	}
}

type KIOPreparationID uint

const (
	KIOPreparationIDAlwaysPrepared KIOPreparationID = 0
	KIOPreparationIDUnprepared     KIOPreparationID = 0
	KIOPreparationIDUnsupported    KIOPreparationID = 0
)

func (e KIOPreparationID) String() string {
	switch e {
	case KIOPreparationIDAlwaysPrepared:
		return "KIOPreparationIDAlwaysPrepared"
	default:
		return fmt.Sprintf("KIOPreparationID(%d)", e)
	}
}

type KIORPC uint

const (
	KIORPCVersion190615 KIORPC = 0
)

func (e KIORPC) String() string {
	switch e {
	case KIORPCVersion190615:
		return "KIORPCVersion190615"
	default:
		return fmt.Sprintf("KIORPC(%d)", e)
	}
}

type KIORPCMessage uint

const (
	KIORPCMessageError     KIORPCMessage = 0
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

type KIORPCMessageID uint

const (
	KIORPCMessageIDKernel KIORPCMessageID = 0
)

func (e KIORPCMessageID) String() string {
	switch e {
	case KIORPCMessageIDKernel:
		return "KIORPCMessageIDKernel"
	default:
		return fmt.Sprintf("KIORPCMessageID(%d)", e)
	}
}

type KIORangeBitsPerColor uint

const (
	KIORangeBitsPerColorComponent10 KIORangeBitsPerColor = 0
	KIORangeBitsPerColorComponent6  KIORangeBitsPerColor = 0
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
	KIORangeColorimetryBT2100    KIORangeColorimetry = 0
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

type KIORangeDynamicRange uint

const (
	KIORangeDynamicRangeDolbyNormalMode     KIORangeDynamicRange = 0
	KIORangeDynamicRangeDolbyTunnelMode     KIORangeDynamicRange = 0
	KIORangeDynamicRangeHDR10               KIORangeDynamicRange = 0
	KIORangeDynamicRangeNotSupported        KIORangeDynamicRange = 0
	KIORangeDynamicRangeSDR                 KIORangeDynamicRange = 0
	KIORangeDynamicRangeTraditionalGammaHDR KIORangeDynamicRange = 0
	KIORangeDynamicRangeTraditionalGammaSDR KIORangeDynamicRange = 0
)

func (e KIORangeDynamicRange) String() string {
	switch e {
	case KIORangeDynamicRangeDolbyNormalMode:
		return "KIORangeDynamicRangeDolbyNormalMode"
	default:
		return fmt.Sprintf("KIORangeDynamicRange(%d)", e)
	}
}

type KIORangePixelEncoding uint

const (
	KIORangePixelEncodingNotSupported KIORangePixelEncoding = 0
	KIORangePixelEncodingRGB444       KIORangePixelEncoding = 0
	KIORangePixelEncodingYCbCr420     KIORangePixelEncoding = 0
	KIORangePixelEncodingYCbCr422     KIORangePixelEncoding = 0
	KIORangePixelEncodingYCbCr444     KIORangePixelEncoding = 0
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
	KIORangeSupportsCompositeSync                  KIORangeSupports = 0
	KIORangeSupportsInterlacedCEATiming            KIORangeSupports = 0
	KIORangeSupportsInterlacedCEATimingWithConfirm KIORangeSupports = 0
	KIORangeSupportsMultiAlignedTiming             KIORangeSupports = 0
	KIORangeSupportsSeparateSyncs                  KIORangeSupports = 0
	KIORangeSupportsSyncOnGreen                    KIORangeSupports = 0
	KIORangeSupportsVRR                            KIORangeSupports = 0
	KIORangeSupportsVSyncSerration                 KIORangeSupports = 0
)

func (e KIORangeSupports) String() string {
	switch e {
	case KIORangeSupportsCompositeSync:
		return "KIORangeSupportsCompositeSync"
	default:
		return fmt.Sprintf("KIORangeSupports(%d)", e)
	}
}

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

type KIORegistryIterate uint

const (
	KIORegistryIterateParents     KIORegistryIterate = 0
	KIORegistryIterateRecursively KIORegistryIterate = 0
)

func (e KIORegistryIterate) String() string {
	switch e {
	case KIORegistryIterateParents:
		return "KIORegistryIterateParents"
	default:
		return fmt.Sprintf("KIORegistryIterate(%d)", e)
	}
}

type KIOReport uint

const (
	KIOReportCopyChannelData  KIOReport = 0
	KIOReportDisable          KIOReport = 0
	KIOReportTraceChannelData KIOReport = 0
)

func (e KIOReport) String() string {
	switch e {
	case KIOReportCopyChannelData:
		return "KIOReportCopyChannelData"
	default:
		return fmt.Sprintf("KIOReport(%d)", e)
	}
}

type KIOReportFormat uint

const (
	KIOReportFormatHistogram KIOReportFormat = 0
)

func (e KIOReportFormat) String() string {
	switch e {
	case KIOReportFormatHistogram:
		return "KIOReportFormatHistogram"
	default:
		return fmt.Sprintf("KIOReportFormat(%d)", e)
	}
}

type KIOReportQuantity uint

const (
	KIOReportQuantityCPUInstrs   KIOReportQuantity = 0
	KIOReportQuantityCapacitance KIOReportQuantity = 0
	KIOReportQuantityCurrent     KIOReportQuantity = 0
	KIOReportQuantityData        KIOReportQuantity = 0
	KIOReportQuantityEnergy      KIOReportQuantity = 0
	KIOReportQuantityEventCount  KIOReportQuantity = 0
	KIOReportQuantityFrequency   KIOReportQuantity = 0
	KIOReportQuantityInductance  KIOReportQuantity = 0
	KIOReportQuantityPacketCount KIOReportQuantity = 0
	KIOReportQuantityPower       KIOReportQuantity = 0
	KIOReportQuantityTemperature KIOReportQuantity = 0
	KIOReportQuantityTime        KIOReportQuantity = 0
	KIOReportQuantityUndefined   KIOReportQuantity = 0
	KIOReportQuantityVoltage     KIOReportQuantity = 0
)

func (e KIOReportQuantity) String() string {
	switch e {
	case KIOReportQuantityCPUInstrs:
		return "KIOReportQuantityCPUInstrs"
	default:
		return fmt.Sprintf("KIOReportQuantity(%d)", e)
	}
}

type KIOReturnOutput uint

const (
	// KIOReturnOutputDropped: # Discussion
	KIOReturnOutputDropped KIOReturnOutput = 0
)

func (e KIOReturnOutput) String() string {
	switch e {
	case KIOReturnOutputDropped:
		return "KIOReturnOutputDropped"
	default:
		return fmt.Sprintf("KIOReturnOutput(%d)", e)
	}
}

type KIOScaleCan uint

const (
	KIOScaleCanBorderInsetOnly  KIOScaleCan = 0
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
	KIOServiceAsynchronous                KIOService = 0
	KIOServiceDextRequirePowerForMatching KIOService = 0
	KIOServiceSeize                       KIOService = 0
)

func (e KIOService) String() string {
	switch e {
	case KIOServiceAsynchronous:
		return "KIOServiceAsynchronous"
	default:
		return fmt.Sprintf("KIOService(%d)", e)
	}
}

type KIOServiceFamilyClose uint

const (
	KIOServiceFamilyCloseOptions KIOServiceFamilyClose = 0
)

func (e KIOServiceFamilyClose) String() string {
	switch e {
	case KIOServiceFamilyCloseOptions:
		return "KIOServiceFamilyCloseOptions"
	default:
		return fmt.Sprintf("KIOServiceFamilyClose(%d)", e)
	}
}

type KIOServiceFirstMatch uint

const (
	KIOServiceFirstMatchState KIOServiceFirstMatch = 0
)

func (e KIOServiceFirstMatch) String() string {
	switch e {
	case KIOServiceFirstMatchState:
		return "KIOServiceFirstMatchState"
	default:
		return fmt.Sprintf("KIOServiceFirstMatch(%d)", e)
	}
}

type KIOServiceHaltStatePower uint

const (
	KIOServiceHaltStatePowerOff KIOServiceHaltStatePower = 0
)

func (e KIOServiceHaltStatePower) String() string {
	switch e {
	case KIOServiceHaltStatePowerOff:
		return "KIOServiceHaltStatePowerOff"
	default:
		return fmt.Sprintf("KIOServiceHaltStatePower(%d)", e)
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

type KIOServicePowerCapability uint

const (
	KIOServicePowerCapabilityLow KIOServicePowerCapability = 0
)

func (e KIOServicePowerCapability) String() string {
	switch e {
	case KIOServicePowerCapabilityLow:
		return "KIOServicePowerCapabilityLow"
	default:
		return fmt.Sprintf("KIOServicePowerCapability(%d)", e)
	}
}

type KIOServiceSearchProperty uint

const (
	KIOServiceSearchPropertyParents KIOServiceSearchProperty = 0
)

func (e KIOServiceSearchProperty) String() string {
	switch e {
	case KIOServiceSearchPropertyParents:
		return "KIOServiceSearchPropertyParents"
	default:
		return fmt.Sprintf("KIOServiceSearchProperty(%d)", e)
	}
}

type KIOStorageAccess uint

const (
	// KIOStorageAccessExclusiveLock: # Discussion
	KIOStorageAccessExclusiveLock KIOStorageAccess = 0
	KIOStorageAccessInvalid       KIOStorageAccess = 0
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
	// KIOStorageOptionIsEncrypted: # Discussion
	KIOStorageOptionIsEncrypted KIOStorageOption = 0
	// KIOStorageOptionIsStatic: # Discussion
	KIOStorageOptionIsStatic KIOStorageOption = 0
	KIOStorageOptionNone     KIOStorageOption = 0
	KIOStorageOptionReserved KIOStorageOption = 0
)

func (e KIOStorageOption) String() string {
	switch e {
	case KIOStorageOptionForceUnitAccess:
		return "KIOStorageOptionForceUnitAccess"
	default:
		return fmt.Sprintf("KIOStorageOption(%d)", e)
	}
}

type KIOStoragePriority uint

const (
	KIOStoragePriorityBackground KIOStoragePriority = 0
)

func (e KIOStoragePriority) String() string {
	switch e {
	case KIOStoragePriorityBackground:
		return "KIOStoragePriorityBackground"
	default:
		return fmt.Sprintf("KIOStoragePriority(%d)", e)
	}
}

type KIOStorageProvisionType uint

const (
	KIOStorageProvisionTypeAnchored KIOStorageProvisionType = 0
)

func (e KIOStorageProvisionType) String() string {
	switch e {
	case KIOStorageProvisionTypeAnchored:
		return "KIOStorageProvisionTypeAnchored"
	default:
		return fmt.Sprintf("KIOStorageProvisionType(%d)", e)
	}
}

type KIOStorageSynchronizeOption uint

const (
	KIOStorageSynchronizeOptionBarrier KIOStorageSynchronizeOption = 0
)

func (e KIOStorageSynchronizeOption) String() string {
	switch e {
	case KIOStorageSynchronizeOptionBarrier:
		return "KIOStorageSynchronizeOptionBarrier"
	default:
		return fmt.Sprintf("KIOStorageSynchronizeOption(%d)", e)
	}
}

type KIOStorageUnmapOption uint

const (
	KIOStorageUnmapOptionNone  KIOStorageUnmapOption = 0
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

type KIOStreamBufferID uint

const (
	KIOStreamBufferIDMask KIOStreamBufferID = 0
)

func (e KIOStreamBufferID) String() string {
	switch e {
	case KIOStreamBufferIDMask:
		return "KIOStreamBufferIDMask"
	default:
		return fmt.Sprintf("KIOStreamBufferID(%d)", e)
	}
}

type KIOStreamEnqueueInputSync uint

const (
	KIOStreamEnqueueInputSyncTrap KIOStreamEnqueueInputSync = 0
)

func (e KIOStreamEnqueueInputSync) String() string {
	switch e {
	case KIOStreamEnqueueInputSyncTrap:
		return "KIOStreamEnqueueInputSyncTrap"
	default:
		return fmt.Sprintf("KIOStreamEnqueueInputSync(%d)", e)
	}
}

type KIOStreamMethod uint

const (
	KIOStreamMethodClose          KIOStreamMethod = 0
	KIOStreamMethodGetBufferCount KIOStreamMethod = 0
	KIOStreamMethodGetMode        KIOStreamMethod = 0
	KIOStreamMethodOpen           KIOStreamMethod = 0
	KIOStreamMethodSetMode        KIOStreamMethod = 0
	KIOStreamMethodStart          KIOStreamMethod = 0
	KIOStreamMethodStop           KIOStreamMethod = 0
	KIOStreamMethodSuspend        KIOStreamMethod = 0
)

func (e KIOStreamMethod) String() string {
	switch e {
	case KIOStreamMethodClose:
		return "KIOStreamMethodClose"
	default:
		return fmt.Sprintf("KIOStreamMethod(%d)", e)
	}
}

type KIOStreamMode uint

const (
	KIOStreamModeInput       KIOStreamMode = 0
	KIOStreamModeInputOutput KIOStreamMode = 0
	KIOStreamModeOutput      KIOStreamMode = 0
)

func (e KIOStreamMode) String() string {
	switch e {
	case KIOStreamModeInput:
		return "KIOStreamModeInput"
	default:
		return fmt.Sprintf("KIOStreamMode(%d)", e)
	}
}

type KIOStreamOptionOpen uint

const (
	KIOStreamOptionOpenExclusive KIOStreamOptionOpen = 0
	KIOStreamOptionOpenShared    KIOStreamOptionOpen = 0
)

func (e KIOStreamOptionOpen) String() string {
	switch e {
	case KIOStreamOptionOpenExclusive:
		return "KIOStreamOptionOpenExclusive"
	default:
		return fmt.Sprintf("KIOStreamOptionOpen(%d)", e)
	}
}

type KIOStreamPortType uint

const (
	KIOStreamPortTypeInput KIOStreamPortType = 0
)

func (e KIOStreamPortType) String() string {
	switch e {
	case KIOStreamPortTypeInput:
		return "KIOStreamPortTypeInput"
	default:
		return fmt.Sprintf("KIOStreamPortType(%d)", e)
	}
}

type KIOSyncPositive uint

const (
	KIOSyncPositivePolarity KIOSyncPositive = 0
)

func (e KIOSyncPositive) String() string {
	switch e {
	case KIOSyncPositivePolarity:
		return "KIOSyncPositivePolarity"
	default:
		return fmt.Sprintf("KIOSyncPositive(%d)", e)
	}
}

type KIOSystemStateSleepDescriptionHibernateState uint

const (
	KIOSystemStateSleepDescriptionHibernateStateHibernating KIOSystemStateSleepDescriptionHibernateState = 0
)

func (e KIOSystemStateSleepDescriptionHibernateState) String() string {
	switch e {
	case KIOSystemStateSleepDescriptionHibernateStateHibernating:
		return "KIOSystemStateSleepDescriptionHibernateStateHibernating"
	default:
		return fmt.Sprintf("KIOSystemStateSleepDescriptionHibernateState(%d)", e)
	}
}

type KIOTimeOptions uint

const (
	KIOTimeOptionsContinuous KIOTimeOptions = 0
)

func (e KIOTimeOptions) String() string {
	switch e {
	case KIOTimeOptionsContinuous:
		return "KIOTimeOptionsContinuous"
	default:
		return fmt.Sprintf("KIOTimeOptions(%d)", e)
	}
}

type KIOTimerEventSourceOptions uint

const (
	KIOTimerEventSourceOptionsAllowReenter       KIOTimerEventSourceOptions = 0
	KIOTimerEventSourceOptionsDefault            KIOTimerEventSourceOptions = 0
	KIOTimerEventSourceOptionsPriorityHigh       KIOTimerEventSourceOptions = 0
	KIOTimerEventSourceOptionsPriorityKernel     KIOTimerEventSourceOptions = 0
	KIOTimerEventSourceOptionsPriorityKernelHigh KIOTimerEventSourceOptions = 0
	KIOTimerEventSourceOptionsPriorityLow        KIOTimerEventSourceOptions = 0
	KIOTimerEventSourceOptionsPriorityMask       KIOTimerEventSourceOptions = 0
	KIOTimerEventSourceOptionsPriorityUser       KIOTimerEventSourceOptions = 0
	KIOTimerEventSourceOptionsPriorityWorkLoop   KIOTimerEventSourceOptions = 0
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
	KIOTimingIDAppleNTSC_FFconv  KIOTimingID = 0
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

type KIOTimingRange uint

const (
	KIOTimingRangeV1 KIOTimingRange = 0
	KIOTimingRangeV2 KIOTimingRange = 0
)

func (e KIOTimingRange) String() string {
	switch e {
	case KIOTimingRangeV1:
		return "KIOTimingRangeV1"
	default:
		return fmt.Sprintf("KIOTimingRange(%d)", e)
	}
}

type KIOTrace uint

const (
	KIOTraceCommandGates KIOTrace = 0
	KIOTraceTimers       KIOTrace = 0
)

func (e KIOTrace) String() string {
	switch e {
	case KIOTraceCommandGates:
		return "KIOTraceCommandGates"
	default:
		return fmt.Sprintf("KIOTrace(%d)", e)
	}
}

type KIOTracking uint

const (
	KIOTrackingGetMappings       KIOTracking = 0
	KIOTrackingGetTracking       KIOTracking = 0
	KIOTrackingInvalid           KIOTracking = 0
	KIOTrackingLeaks             KIOTracking = 0
	KIOTrackingResetTracking     KIOTracking = 0
	KIOTrackingSetMinCaptureSize KIOTracking = 0
	KIOTrackingStartCapture      KIOTracking = 0
	KIOTrackingStopCapture       KIOTracking = 0
)

func (e KIOTracking) String() string {
	switch e {
	case KIOTrackingGetMappings:
		return "KIOTrackingGetMappings"
	default:
		return fmt.Sprintf("KIOTracking(%d)", e)
	}
}

type KIOTrackingCallSiteB uint

const (
	KIOTrackingCallSiteBTs KIOTrackingCallSiteB = 0
)

func (e KIOTrackingCallSiteB) String() string {
	switch e {
	case KIOTrackingCallSiteBTs:
		return "KIOTrackingCallSiteBTs"
	default:
		return fmt.Sprintf("KIOTrackingCallSiteB(%d)", e)
	}
}

type KIOTrackingExclude uint

const (
	KIOTrackingExcludeNames KIOTrackingExclude = 0
)

func (e KIOTrackingExclude) String() string {
	switch e {
	case KIOTrackingExcludeNames:
		return "KIOTrackingExcludeNames"
	default:
		return fmt.Sprintf("KIOTrackingExclude(%d)", e)
	}
}

type KIOTrackingLeakScan uint

const (
	KIOTrackingLeakScanEnd KIOTrackingLeakScan = 0
)

func (e KIOTrackingLeakScan) String() string {
	switch e {
	case KIOTrackingLeakScanEnd:
		return "KIOTrackingLeakScanEnd"
	default:
		return fmt.Sprintf("KIOTrackingLeakScan(%d)", e)
	}
}

type KIOUC uint

const (
	KIOUCForegroundOnly KIOUC = 0
	KIOUCScalarIScalarO KIOUC = 0
	KIOUCScalarIStructI KIOUC = 0
	KIOUCScalarIStructO KIOUC = 0
	KIOUCStructIStructO KIOUC = 0
	KIOUCTypeMask       KIOUC = 0
)

func (e KIOUC) String() string {
	switch e {
	case KIOUCForegroundOnly:
		return "KIOUCForegroundOnly"
	default:
		return fmt.Sprintf("KIOUC(%d)", e)
	}
}

type KIOUCVariableStructure uint

const (
	// KIOUCVariableStructureSize: # Discussion
	KIOUCVariableStructureSize KIOUCVariableStructure = 0
)

func (e KIOUCVariableStructure) String() string {
	switch e {
	case KIOUCVariableStructureSize:
		return "KIOUCVariableStructureSize"
	default:
		return fmt.Sprintf("KIOUCVariableStructure(%d)", e)
	}
}

type KIOUSB uint

const (
	// KIOUSBDecriptorTypeHID: The type for a human interaction device descriptor.
	KIOUSBDecriptorTypeHID KIOUSB = 0
	// KIOUSBDecriptorTypeReport: The type for a report descriptor.
	KIOUSBDecriptorTypeReport                   KIOUSB = 0
	KIOUSBDescriptorTypeBOS                     KIOUSB = 0
	KIOUSBDescriptorTypeConfiguration           KIOUSB = 0
	KIOUSBDescriptorTypeDebug                   KIOUSB = 0
	KIOUSBDescriptorTypeDevice                  KIOUSB = 0
	KIOUSBDescriptorTypeDeviceCapability        KIOUSB = 0
	KIOUSBDescriptorTypeDeviceQualifier         KIOUSB = 0
	KIOUSBDescriptorTypeEndpoint                KIOUSB = 0
	KIOUSBDescriptorTypeHub                     KIOUSB = 0
	KIOUSBDescriptorTypeInterface               KIOUSB = 0
	KIOUSBDescriptorTypeInterfaceAssociation    KIOUSB = 0
	KIOUSBDescriptorTypeInterfacePower          KIOUSB = 0
	KIOUSBDescriptorTypeOTG                     KIOUSB = 0
	KIOUSBDescriptorTypeOtherSpeedConfiguration KIOUSB = 0
	// KIOUSBDescriptorTypePhysical: The type for a physical descriptor.
	KIOUSBDescriptorTypePhysical                                   KIOUSB = 0
	KIOUSBDescriptorTypeString                                     KIOUSB = 0
	KIOUSBDescriptorTypeSuperSpeedHub                              KIOUSB = 0
	KIOUSBDescriptorTypeSuperSpeedPlusIsochronousEndpointCompanion KIOUSB = 0
	KIOUSBDescriptorTypeSuperSpeedUSBEndpointCompanion             KIOUSB = 0
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
	KIOUSB30LinkStateHotResetDeadline      KIOUSB30 = 0
	KIOUSB30LinkStateHotResetExitTimeout   KIOUSB30 = 0
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
	KIOUSB30LinkStateRecoveryDeadline             KIOUSB30 = 0
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

type KIOUSBAny uint

const (
	KIOUSBAnyClass KIOUSBAny = 0
)

func (e KIOUSBAny) String() string {
	switch e {
	case KIOUSBAnyClass:
		return "KIOUSBAnyClass"
	default:
		return fmt.Sprintf("KIOUSBAny(%d)", e)
	}
}

type KIOUSBAppleVendorI uint

const (
	KIOUSBAppleVendorID KIOUSBAppleVendorI = 0
)

func (e KIOUSBAppleVendorI) String() string {
	switch e {
	case KIOUSBAppleVendorID:
		return "KIOUSBAppleVendorID"
	default:
		return fmt.Sprintf("KIOUSBAppleVendorI(%d)", e)
	}
}

type KIOUSBBusVoltage uint

const (
	// KIOUSBBusVoltageDefault: The default bus voltage.
	KIOUSBBusVoltageDefault KIOUSBBusVoltage = 0
)

func (e KIOUSBBusVoltage) String() string {
	switch e {
	case KIOUSBBusVoltageDefault:
		return "KIOUSBBusVoltageDefault"
	default:
		return fmt.Sprintf("KIOUSBBusVoltage(%d)", e)
	}
}

type KIOUSBConfigurationDescriptorAttribute uint

const (
	KIOUSBConfigurationDescriptorAttributeRemoteWakeCapable KIOUSBConfigurationDescriptorAttribute = 0
	KIOUSBConfigurationDescriptorAttributeSelfPowered       KIOUSBConfigurationDescriptorAttribute = 0
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
	KIOUSBDescriptorHeaderSize                                     KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeBOS                                        KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeBillboardDeviceMaximum                     KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeBillboardDeviceMinimum                     KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeConfiguration                              KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeContainerIDCapability                      KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeDevice                                     KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeDeviceCapability                           KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeDeviceQualifier                            KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeEndpoint                                   KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeHubMaximum                                 KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeHubMinimum                                 KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeInterface                                  KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeInterfaceAssociation                       KIOUSBDescriptor = 0
	KIOUSBDescriptorSizePlatformCapability                         KIOUSBDescriptor = 0
	KIOUSBDescriptorSizePlatformECIDCapability                     KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeStringMaximum                              KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeStringMinimum                              KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeSuperSpeedHub                              KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeSuperSpeedPlusIsochronousEndpointCompanion KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeSuperSpeedUSBDeviceCapability              KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeSuperSpeedUSBEndpointCompanion             KIOUSBDescriptor = 0
	KIOUSBDescriptorSizeUSB20ExtensionCapability                   KIOUSBDescriptor = 0
)

func (e KIOUSBDescriptor) String() string {
	switch e {
	case KIOUSBDescriptorHeaderSize:
		return "KIOUSBDescriptorHeaderSize"
	default:
		return fmt.Sprintf("KIOUSBDescriptor(%d)", e)
	}
}

type KIOUSBDeviceCapabilityDescriptor uint

const (
	KIOUSBDeviceCapabilityDescriptorLengthMin KIOUSBDeviceCapabilityDescriptor = 0
	KIOUSBDeviceCapabilityDescriptorType      KIOUSBDeviceCapabilityDescriptor = 0
)

func (e KIOUSBDeviceCapabilityDescriptor) String() string {
	switch e {
	case KIOUSBDeviceCapabilityDescriptorLengthMin:
		return "KIOUSBDeviceCapabilityDescriptorLengthMin"
	default:
		return fmt.Sprintf("KIOUSBDeviceCapabilityDescriptor(%d)", e)
	}
}

type KIOUSBDeviceCapabilityType uint

const (
	KIOUSBDeviceCapabilityTypeBatteryInfo          KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeBillboard            KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeBillboardAltMode     KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeContainerID          KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePdConsumerPort       KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePdProviderPort       KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePlatform             KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePowerDelivery        KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypePrecisionMeasurement KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeSuperSpeed           KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeSuperSpeedPlus       KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeUSB20Extension       KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeWireless             KIOUSBDeviceCapabilityType = 0
	KIOUSBDeviceCapabilityTypeWirelessExt          KIOUSBDeviceCapabilityType = 0
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
	KIOUSBDeviceFeatureSelectorU1Enable  KIOUSBDeviceFeatureSelector = 0
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
	KIOUSBDeviceRequestClearFeature KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestDirectionIn: The incoming direction.
	KIOUSBDeviceRequestDirectionIn KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestDirectionMask: The direction mask.
	KIOUSBDeviceRequestDirectionMask KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestDirectionOut: The outgoing direction.
	KIOUSBDeviceRequestDirectionOut KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestDirectionPhase: The direction phase.
	KIOUSBDeviceRequestDirectionPhase   KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestGetConfiguration KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestGetDescriptor    KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestGetInterface     KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestGetState         KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestGetStatus        KIOUSBDeviceRequest = 0
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
	KIOUSBDeviceRequestRecipientPhase      KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestSetAddress          KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestSetConfiguration    KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestSetDescriptor       KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestSetFeature          KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestSetInterface        KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestSetIsochronousDelay KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestSetSel              KIOUSBDeviceRequest = 0
	// KIOUSBDeviceRequestSize: The request size.
	KIOUSBDeviceRequestSize       KIOUSBDeviceRequest = 0
	KIOUSBDeviceRequestSynchFrame KIOUSBDeviceRequest = 0
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
	case KIOUSBDeviceRequestClearFeature:
		return "KIOUSBDeviceRequestClearFeature"
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
	IOUSBEndpointStatusHalt            KIOUSBDeviceStatusSelfPowered = 0
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
	KIOUSBEndpointDescriptorNumber                   KIOUSBEndpointDescriptor = 0
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
	KIOUSBEndpointDirectionIn      KIOUSBEndpointDirection = 0
	KIOUSBEndpointDirectionOut     KIOUSBEndpointDirection = 0
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
	KIOUSBEndpointTypeBulk        KIOUSBEndpointType = 0
	KIOUSBEndpointTypeControl     KIOUSBEndpointType = 0
	KIOUSBEndpointTypeInterrupt   KIOUSBEndpointType = 0
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

type KIOUSBFindInterfaceDont uint

const (
	KIOUSBFindInterfaceDontCare KIOUSBFindInterfaceDont = 0
)

func (e KIOUSBFindInterfaceDont) String() string {
	switch e {
	case KIOUSBFindInterfaceDontCare:
		return "KIOUSBFindInterfaceDontCare"
	default:
		return fmt.Sprintf("KIOUSBFindInterfaceDont(%d)", e)
	}
}

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

type KIOUSBHostPipeBundling uint

const (
	KIOUSBHostPipeBundlingMax KIOUSBHostPipeBundling = 0
)

func (e KIOUSBHostPipeBundling) String() string {
	switch e {
	case KIOUSBHostPipeBundlingMax:
		return "KIOUSBHostPipeBundlingMax"
	default:
		return fmt.Sprintf("KIOUSBHostPipeBundling(%d)", e)
	}
}

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

type KIOUSBHub uint

const (
	KIOUSBHubDelayNs                KIOUSBHub = 0
	KIOUSBHubPort2PortExitLatencyNs KIOUSBHub = 0
)

func (e KIOUSBHub) String() string {
	switch e {
	case KIOUSBHubDelayNs:
		return "KIOUSBHubDelayNs"
	default:
		return fmt.Sprintf("KIOUSBHub(%d)", e)
	}
}

type KIOUSBInterfaceOpen uint

const (
	// KIOUSBInterfaceOpenAlt: # Discussion
	KIOUSBInterfaceOpenAlt KIOUSBInterfaceOpen = 0
)

func (e KIOUSBInterfaceOpen) String() string {
	switch e {
	case KIOUSBInterfaceOpenAlt:
		return "KIOUSBInterfaceOpenAlt"
	default:
		return fmt.Sprintf("KIOUSBInterfaceOpen(%d)", e)
	}
}

type KIOUSBInterfaceSuspendLow uint

const (
	KIOUSBInterfaceSuspendLowPower KIOUSBInterfaceSuspendLow = 0
)

func (e KIOUSBInterfaceSuspendLow) String() string {
	switch e {
	case KIOUSBInterfaceSuspendLowPower:
		return "KIOUSBInterfaceSuspendLowPower"
	default:
		return fmt.Sprintf("KIOUSBInterfaceSuspendLow(%d)", e)
	}
}

type KIOUSBLanguageIDEnglishU uint

const (
	KIOUSBLanguageIDEnglishUS KIOUSBLanguageIDEnglishU = 0
)

func (e KIOUSBLanguageIDEnglishU) String() string {
	switch e {
	case KIOUSBLanguageIDEnglishUS:
		return "KIOUSBLanguageIDEnglishUS"
	default:
		return fmt.Sprintf("KIOUSBLanguageIDEnglishU(%d)", e)
	}
}

type KIOUSBPingResponseTime uint

const (
	KIOUSBPingResponseTimeNs KIOUSBPingResponseTime = 0
)

func (e KIOUSBPingResponseTime) String() string {
	switch e {
	case KIOUSBPingResponseTimeNs:
		return "KIOUSBPingResponseTimeNs"
	default:
		return fmt.Sprintf("KIOUSBPingResponseTime(%d)", e)
	}
}

type KIOUSBSuperSpeedDeviceCapability uint

const (
	KIOUSBSuperSpeedDeviceCapabilityHighSpeed        KIOUSBSuperSpeedDeviceCapability = 0
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
	KIOUSBSuperSpeedEndpointCompanionDescriptorBulkMaxStreams      KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorBulkMaxStreamsPhase KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorBulkReserved        KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorBulkReservedPhase   KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorIsocMult            KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorIsocMultPhase       KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorIsocReserved        KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorIsocReservedPhase   KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorMaxBurst            KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorMaxBurstPhase       KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
	KIOUSBSuperSpeedEndpointCompanionDescriptorSSPIsocCompanion    KIOUSBSuperSpeedEndpointCompanionDescriptor = 0
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
	KIOUSBSuperSpeedHubCharacteristicsCompoundDevice    KIOUSBSuperSpeedHubCharacteristics = 0
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
	KIOUSBSuperSpeedPlusDeviceCapabilitySublinkSymmetric      KIOUSBSuperSpeedPlusDeviceCapabilitySublink = 0
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
	KIOUSBUSB20ExtensionCapabilityBESL      KIOUSBUSB20ExtensionCapabilityBES = 0
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

type KIOUSBVendorID uint

const (
	KIOUSBVendorIDApple KIOUSBVendorID = 0
)

func (e KIOUSBVendorID) String() string {
	switch e {
	case KIOUSBVendorIDApple:
		return "KIOUSBVendorIDApple"
	default:
		return fmt.Sprintf("KIOUSBVendorID(%d)", e)
	}
}

type KIOUserClientAsyncArgumentsCount uint

const (
	KIOUserClientAsyncArgumentsCountMax KIOUserClientAsyncArgumentsCount = 0
)

func (e KIOUserClientAsyncArgumentsCount) String() string {
	switch e {
	case KIOUserClientAsyncArgumentsCountMax:
		return "KIOUserClientAsyncArgumentsCountMax"
	default:
		return fmt.Sprintf("KIOUserClientAsyncArgumentsCount(%d)", e)
	}
}

type KIOUserClientAsyncReferenceCount uint

const (
	KIOUserClientAsyncReferenceCountMax KIOUserClientAsyncReferenceCount = 0
)

func (e KIOUserClientAsyncReferenceCount) String() string {
	switch e {
	case KIOUserClientAsyncReferenceCountMax:
		return "KIOUserClientAsyncReferenceCountMax"
	default:
		return fmt.Sprintf("KIOUserClientAsyncReferenceCount(%d)", e)
	}
}

type KIOUserClientMemoryRead uint

const (
	KIOUserClientMemoryReadOnly KIOUserClientMemoryRead = 0
)

func (e KIOUserClientMemoryRead) String() string {
	switch e {
	case KIOUserClientMemoryReadOnly:
		return "KIOUserClientMemoryReadOnly"
	default:
		return fmt.Sprintf("KIOUserClientMemoryRead(%d)", e)
	}
}

type KIOUserClientMethodArgumentsCurrent uint

const (
	KIOUserClientMethodArgumentsCurrentVersion KIOUserClientMethodArgumentsCurrent = 0
)

func (e KIOUserClientMethodArgumentsCurrent) String() string {
	switch e {
	case KIOUserClientMethodArgumentsCurrentVersion:
		return "KIOUserClientMethodArgumentsCurrentVersion"
	default:
		return fmt.Sprintf("KIOUserClientMethodArgumentsCurrent(%d)", e)
	}
}

type KIOUserClientScalarArrayCount uint

const (
	KIOUserClientScalarArrayCountMax KIOUserClientScalarArrayCount = 0
)

func (e KIOUserClientScalarArrayCount) String() string {
	switch e {
	case KIOUserClientScalarArrayCountMax:
		return "KIOUserClientScalarArrayCountMax"
	default:
		return fmt.Sprintf("KIOUserClientScalarArrayCount(%d)", e)
	}
}

type KIOUserClientVariableStructure uint

const (
	KIOUserClientVariableStructureSize KIOUserClientVariableStructure = 0
)

func (e KIOUserClientVariableStructure) String() string {
	switch e {
	case KIOUserClientVariableStructureSize:
		return "KIOUserClientVariableStructureSize"
	default:
		return fmt.Sprintf("KIOUserClientVariableStructure(%d)", e)
	}
}

type KIOUserNotifyMaxMessage uint

const (
	KIOUserNotifyMaxMessageSize KIOUserNotifyMaxMessage = 0
)

func (e KIOUserNotifyMaxMessage) String() string {
	switch e {
	case KIOUserNotifyMaxMessageSize:
		return "KIOUserNotifyMaxMessageSize"
	default:
		return fmt.Sprintf("KIOUserNotifyMaxMessage(%d)", e)
	}
}

type KIOUserNotifyOptionCan uint

const (
	KIOUserNotifyOptionCanDrop KIOUserNotifyOptionCan = 0
)

func (e KIOUserNotifyOptionCan) String() string {
	switch e {
	case KIOUserNotifyOptionCanDrop:
		return "KIOUserNotifyOptionCanDrop"
	default:
		return fmt.Sprintf("KIOUserNotifyOptionCan(%d)", e)
	}
}

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
	// KIOVideoControlScopeGlobal: # Discussion
	KIOVideoControlScopeGlobal KIOVideoControl = 0
	// KIOVideoControlScopeInput: # Discussion
	KIOVideoControlScopeInput KIOVideoControl = 0
	// KIOVideoControlScopeOutput: # Discussion
	KIOVideoControlScopeOutput KIOVideoControl = 0
	// KIOVideoControlScopePlayThrough: # Discussion
	KIOVideoControlScopePlayThrough KIOVideoControl = 0
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
	KIOVideoDeviceMethodOpen  KIOVideoDeviceMethod = 0
)

func (e KIOVideoDeviceMethod) String() string {
	switch e {
	case KIOVideoDeviceMethodClose:
		return "KIOVideoDeviceMethodClose"
	default:
		return fmt.Sprintf("KIOVideoDeviceMethod(%d)", e)
	}
}

type KIOVideoDevicePortType uint

const (
	KIOVideoDevicePortTypeInput KIOVideoDevicePortType = 0
)

func (e KIOVideoDevicePortType) String() string {
	switch e {
	case KIOVideoDevicePortTypeInput:
		return "KIOVideoDevicePortTypeInput"
	default:
		return fmt.Sprintf("KIOVideoDevicePortType(%d)", e)
	}
}

type KIOVideoFeatureControlClassID uint

const (
	// KIOVideoFeatureControlClassIDBacklightCompensation: # Discussion
	KIOVideoFeatureControlClassIDBacklightCompensation KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDBlackLevel: # Discussion
	KIOVideoFeatureControlClassIDBlackLevel KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDBrightness: # Discussion
	KIOVideoFeatureControlClassIDBrightness KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDContrast: # Discussion
	KIOVideoFeatureControlClassIDContrast KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDExposure: # Discussion
	KIOVideoFeatureControlClassIDExposure KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDFocus: # Discussion
	KIOVideoFeatureControlClassIDFocus KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDGain: # Discussion
	KIOVideoFeatureControlClassIDGain KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDGamma: # Discussion
	KIOVideoFeatureControlClassIDGamma KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDHue: # Discussion
	KIOVideoFeatureControlClassIDHue KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDIris: # Discussion
	KIOVideoFeatureControlClassIDIris KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDOpticalFilter: # Discussion
	KIOVideoFeatureControlClassIDOpticalFilter KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDPan: # Discussion
	KIOVideoFeatureControlClassIDPan KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDPowerLineFrequency: # Discussion
	KIOVideoFeatureControlClassIDPowerLineFrequency KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDSaturation: # Discussion
	KIOVideoFeatureControlClassIDSaturation KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDSharpness: # Discussion
	KIOVideoFeatureControlClassIDSharpness KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDShutter: # Discussion
	KIOVideoFeatureControlClassIDShutter KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDTemperature: # Discussion
	KIOVideoFeatureControlClassIDTemperature KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDTilt: # Discussion
	KIOVideoFeatureControlClassIDTilt KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDWhiteBalanceU: # Discussion
	KIOVideoFeatureControlClassIDWhiteBalanceU KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDWhiteBalanceV: # Discussion
	KIOVideoFeatureControlClassIDWhiteBalanceV KIOVideoFeatureControlClassID = 0
	// KIOVideoFeatureControlClassIDWhiteLevel: # Discussion
	KIOVideoFeatureControlClassIDWhiteLevel KIOVideoFeatureControlClassID = 0
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

type KIOVideoSelectorControlClassIDData uint

const (
	// KIOVideoSelectorControlClassIDDataDestination: # Discussion
	KIOVideoSelectorControlClassIDDataDestination KIOVideoSelectorControlClassIDData = 0
)

func (e KIOVideoSelectorControlClassIDData) String() string {
	switch e {
	case KIOVideoSelectorControlClassIDDataDestination:
		return "KIOVideoSelectorControlClassIDDataDestination"
	default:
		return fmt.Sprintf("KIOVideoSelectorControlClassIDData(%d)", e)
	}
}

type KIOWSAA uint

const (
	KIOWSAA_Accelerated KIOWSAA = 0
	KIOWSAA_Hibernate   KIOWSAA = 0
)

func (e KIOWSAA) String() string {
	switch e {
	case KIOWSAA_Accelerated:
		return "KIOWSAA_Accelerated"
	default:
		return fmt.Sprintf("KIOWSAA(%d)", e)
	}
}

type KIOWorkGroupMaxName uint

const (
	KIOWorkGroupMaxNameLength KIOWorkGroupMaxName = 0
)

func (e KIOWorkGroupMaxName) String() string {
	switch e {
	case KIOWorkGroupMaxNameLength:
		return "KIOWorkGroupMaxNameLength"
	default:
		return fmt.Sprintf("KIOWorkGroupMaxName(%d)", e)
	}
}

type KInflowForciblyEnabled uint

const (
	KInflowForciblyEnabledBit KInflowForciblyEnabled = 0
)

func (e KInflowForciblyEnabled) String() string {
	switch e {
	case KInflowForciblyEnabledBit:
		return "KInflowForciblyEnabledBit"
	default:
		return fmt.Sprintf("KInflowForciblyEnabled(%d)", e)
	}
}

type KInitialDriver uint

const (
	KInitialDriverDescriptor KInitialDriver = 0
)

func (e KInitialDriver) String() string {
	switch e {
	case KInitialDriverDescriptor:
		return "KInitialDriverDescriptor"
	default:
		return fmt.Sprintf("KInitialDriver(%d)", e)
	}
}

type KInterrupt uint

const (
	KInterruptRecord KInterrupt = 0
)

func (e KInterrupt) String() string {
	switch e {
	case KInterruptRecord:
		return "KInterruptRecord"
	default:
		return fmt.Sprintf("KInterrupt(%d)", e)
	}
}

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

type KJIJournalInFS uint

const (
	KJIJournalInFSMask KJIJournalInFS = 0
)

func (e KJIJournalInFS) String() string {
	switch e {
	case KJIJournalInFSMask:
		return "KJIJournalInFSMask"
	default:
		return fmt.Sprintf("KJIJournalInFS(%d)", e)
	}
}

type KKUNCAlternate uint

const (
	KKUNCAlternateResponse KKUNCAlternate = 0
)

func (e KKUNCAlternate) String() string {
	switch e {
	case KKUNCAlternateResponse:
		return "KKUNCAlternateResponse"
	default:
		return fmt.Sprintf("KKUNCAlternate(%d)", e)
	}
}

type KKeyMask uint

const (
	KKeyMaskAlt          KKeyMask = 0
	KKeyMaskComma        KKeyMask = 0
	KKeyMaskCtrl         KKeyMask = 0
	KKeyMaskDelete       KKeyMask = 0
	KKeyMaskEsc          KKeyMask = 0
	KKeyMaskFn           KKeyMask = 0
	KKeyMaskLeftCommand  KKeyMask = 0
	KKeyMaskPeriod       KKeyMask = 0
	KKeyMaskPower        KKeyMask = 0
	KKeyMaskRightCommand KKeyMask = 0
	KKeyMaskShift        KKeyMask = 0
	KKeyMaskSlash        KKeyMask = 0
	KKeyMaskUnknown      KKeyMask = 0
)

func (e KKeyMask) String() string {
	switch e {
	case KKeyMaskAlt:
		return "KKeyMaskAlt"
	default:
		return fmt.Sprintf("KKeyMask(%d)", e)
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

type KLoadALSS uint

const (
	KLoadALSSCalibration KLoadALSS = 0
)

func (e KLoadALSS) String() string {
	switch e {
	case KLoadALSSCalibration:
		return "KLoadALSSCalibration"
	default:
		return fmt.Sprintf("KLoadALSS(%d)", e)
	}
}

type KMMC uint

const (
	KMMCNumPowerStates        KMMC = 0
	KMMCPowerStateActive      KMMC = 0
	KMMCPowerStateIdle        KMMC = 0
	KMMCPowerStateSleep       KMMC = 0
	KMMCPowerStateStandby     KMMC = 0
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

type KMMCCmd uint

const (
	KMMCCmd_BLANK                         KMMCCmd = 0
	KMMCCmd_CHANGE_DEFINITION             KMMCCmd = 0
	KMMCCmd_CLOSE_TRACK_SESSION           KMMCCmd = 0
	KMMCCmd_COMPARE                       KMMCCmd = 0
	KMMCCmd_COPY                          KMMCCmd = 0
	KMMCCmd_COPY_AND_VERIFY               KMMCCmd = 0
	KMMCCmd_ERASE                         KMMCCmd = 0
	KMMCCmd_FORMAT_UNIT                   KMMCCmd = 0
	KMMCCmd_GET_CONFIGURATION             KMMCCmd = 0
	KMMCCmd_GET_EVENT_STATUS_NOTIFICATION KMMCCmd = 0
	KMMCCmd_GET_PERFORMANCE               KMMCCmd = 0
	KMMCCmd_INQUIRY                       KMMCCmd = 0
	KMMCCmd_LOAD_UNLOAD_MEDIUM            KMMCCmd = 0
	KMMCCmd_LOG_SELECT                    KMMCCmd = 0
	KMMCCmd_LOG_SENSE                     KMMCCmd = 0
	KMMCCmd_MECHANISM_STATUS              KMMCCmd = 0
	KMMCCmd_MODE_SELECT_10                KMMCCmd = 0
	KMMCCmd_MODE_SELECT_6                 KMMCCmd = 0
	KMMCCmd_MODE_SENSE_10                 KMMCCmd = 0
	KMMCCmd_MODE_SENSE_6                  KMMCCmd = 0
	KMMCCmd_PAUSE_RESUME                  KMMCCmd = 0
	KMMCCmd_PLAY_AUDIO_10                 KMMCCmd = 0
	KMMCCmd_PLAY_AUDIO_12                 KMMCCmd = 0
	KMMCCmd_PLAY_AUDIO_MSF                KMMCCmd = 0
	KMMCCmd_PLAY_AUDIO_TRACK_INDEX        KMMCCmd = 0
	KMMCCmd_PLAY_CD                       KMMCCmd = 0
	KMMCCmd_PLAY_RELATIVE_10              KMMCCmd = 0
	KMMCCmd_PLAY_RELATIVE_12              KMMCCmd = 0
	KMMCCmd_PREFETCH                      KMMCCmd = 0
	KMMCCmd_PREVENT_ALLOW_MEDIUM_REMOVAL  KMMCCmd = 0
	KMMCCmd_READ_10                       KMMCCmd = 0
	KMMCCmd_READ_12                       KMMCCmd = 0
	KMMCCmd_READ_6                        KMMCCmd = 0
	KMMCCmd_READ_BUFFER                   KMMCCmd = 0
	KMMCCmd_READ_BUFFER_CAPACITY          KMMCCmd = 0
	KMMCCmd_READ_CAPACITY                 KMMCCmd = 0
	KMMCCmd_READ_CD                       KMMCCmd = 0
	KMMCCmd_READ_CD_MSF                   KMMCCmd = 0
	KMMCCmd_READ_DISC_INFORMATION         KMMCCmd = 0
	KMMCCmd_READ_DISC_STRUCTURE           KMMCCmd = 0
	KMMCCmd_READ_DVD_STRUCTURE            KMMCCmd = 0
	KMMCCmd_READ_FORMAT_CAPACITIES        KMMCCmd = 0
	KMMCCmd_READ_HEADER                   KMMCCmd = 0
	KMMCCmd_READ_LONG                     KMMCCmd = 0
	KMMCCmd_READ_MASTER_CUE               KMMCCmd = 0
	KMMCCmd_READ_SUB_CHANNEL              KMMCCmd = 0
	KMMCCmd_READ_TOC_PMA_ATIP             KMMCCmd = 0
	KMMCCmd_READ_TRACK_INFORMATION        KMMCCmd = 0
	KMMCCmd_RECEIVE_DIAGNOSTICS_RESULTS   KMMCCmd = 0
	KMMCCmd_RELEASE_10                    KMMCCmd = 0
	KMMCCmd_RELEASE_6                     KMMCCmd = 0
	KMMCCmd_REPAIR_TRACK                  KMMCCmd = 0
	KMMCCmd_REPORT_KEY                    KMMCCmd = 0
	KMMCCmd_REQUEST_SENSE                 KMMCCmd = 0
	KMMCCmd_RESERVE_10                    KMMCCmd = 0
	KMMCCmd_RESERVE_6                     KMMCCmd = 0
	KMMCCmd_RESERVE_TRACK                 KMMCCmd = 0
	KMMCCmd_SCAN_MMC                      KMMCCmd = 0
	KMMCCmd_SEARCH_DATA_EQUAL_10          KMMCCmd = 0
	KMMCCmd_SEARCH_DATA_EQUAL_12          KMMCCmd = 0
	KMMCCmd_SEARCH_DATA_HIGH_10           KMMCCmd = 0
	KMMCCmd_SEARCH_DATA_HIGH_12           KMMCCmd = 0
	KMMCCmd_SEARCH_DATA_LOW_10            KMMCCmd = 0
	KMMCCmd_SEARCH_DATA_LOW_12            KMMCCmd = 0
	KMMCCmd_SEEK_10                       KMMCCmd = 0
	KMMCCmd_SEEK_6                        KMMCCmd = 0
	KMMCCmd_SEND_CUE_SHEET                KMMCCmd = 0
	KMMCCmd_SEND_DIAGNOSTICS              KMMCCmd = 0
	KMMCCmd_SEND_DVD_STRUCTURE            KMMCCmd = 0
	KMMCCmd_SEND_EVENT                    KMMCCmd = 0
	KMMCCmd_SEND_KEY                      KMMCCmd = 0
	KMMCCmd_SEND_OPC_INFORMATION          KMMCCmd = 0
	KMMCCmd_SET_CD_SPEED                  KMMCCmd = 0
	KMMCCmd_SET_LIMITS_10                 KMMCCmd = 0
	KMMCCmd_SET_LIMITS_12                 KMMCCmd = 0
	KMMCCmd_SET_READ_AHEAD                KMMCCmd = 0
	KMMCCmd_SET_STREAMING                 KMMCCmd = 0
	KMMCCmd_START_STOP_UNIT               KMMCCmd = 0
	KMMCCmd_STOP_PLAY_SCAN                KMMCCmd = 0
	KMMCCmd_SYNCHRONIZE_CACHE             KMMCCmd = 0
	KMMCCmd_TEST_UNIT_READY               KMMCCmd = 0
	KMMCCmd_VERIFY_10                     KMMCCmd = 0
	KMMCCmd_VERIFY_12                     KMMCCmd = 0
	KMMCCmd_WRITE_10                      KMMCCmd = 0
	KMMCCmd_WRITE_12                      KMMCCmd = 0
	KMMCCmd_WRITE_AND_VERIFY_10           KMMCCmd = 0
	KMMCCmd_WRITE_BUFFER                  KMMCCmd = 0
)

func (e KMMCCmd) String() string {
	switch e {
	case KMMCCmd_BLANK:
		return "KMMCCmd_BLANK"
	default:
		return fmt.Sprintf("KMMCCmd(%d)", e)
	}
}

type KMachineType uint

const (
	KMachineTypeUnknown KMachineType = 0
)

func (e KMachineType) String() string {
	switch e {
	case KMachineTypeUnknown:
		return "KMachineTypeUnknown"
	default:
		return fmt.Sprintf("KMachineType(%d)", e)
	}
}

type KMaxKey uint

const (
	KMaxKeyLength KMaxKey = 0
)

func (e KMaxKey) String() string {
	switch e {
	case KMaxKeyLength:
		return "KMaxKeyLength"
	default:
		return fmt.Sprintf("KMaxKey(%d)", e)
	}
}

type KMaximumNumberOfInquiryAccess uint

const (
	KMaximumNumberOfInquiryAccessCodes KMaximumNumberOfInquiryAccess = 0
)

func (e KMaximumNumberOfInquiryAccess) String() string {
	switch e {
	case KMaximumNumberOfInquiryAccessCodes:
		return "KMaximumNumberOfInquiryAccessCodes"
	default:
		return fmt.Sprintf("KMaximumNumberOfInquiryAccess(%d)", e)
	}
}

type KMediaState uint

const (
	KMediaStateLocked   KMediaState = 0
	KMediaStateUnlocked KMediaState = 0
)

func (e KMediaState) String() string {
	switch e {
	case KMediaStateLocked:
		return "KMediaStateLocked"
	default:
		return fmt.Sprintf("KMediaState(%d)", e)
	}
}

type KMessageDeterminingMedia uint

const (
	KMessageDeterminingMediaPresence KMessageDeterminingMedia = 0
)

func (e KMessageDeterminingMedia) String() string {
	switch e {
	case KMessageDeterminingMediaPresence:
		return "KMessageDeterminingMediaPresence"
	default:
		return fmt.Sprintf("KMessageDeterminingMedia(%d)", e)
	}
}

type KMessageTrayStateChangeRequest uint

const (
	KMessageTrayStateChangeRequestAccepted KMessageTrayStateChangeRequest = 0
	KMessageTrayStateChangeRequestRejected KMessageTrayStateChangeRequest = 0
)

func (e KMessageTrayStateChangeRequest) String() string {
	switch e {
	case KMessageTrayStateChangeRequestAccepted:
		return "KMessageTrayStateChangeRequestAccepted"
	default:
		return fmt.Sprintf("KMessageTrayStateChangeRequest(%d)", e)
	}
}

type KMirror uint

const (
	KMirrorAreMirroredMask          KMirror = 0
	KMirrorCanChangePixelFormatMask KMirror = 0
	KMirrorCanChangeTimingMask      KMirror = 0
	KMirrorCanMirrorMask            KMirror = 0
	KMirrorClippedMirrorMask        KMirror = 0
	KMirrorHAlignCenterMirrorMask   KMirror = 0
	KMirrorUnclippedMirrorMask      KMirror = 0
	KMirrorVAlignCenterMirrorMask   KMirror = 0
)

func (e KMirror) String() string {
	switch e {
	case KMirrorAreMirroredMask:
		return "KMirrorAreMirroredMask"
	default:
		return fmt.Sprintf("KMirror(%d)", e)
	}
}

type KMirrorCommonGamma uint

const (
	KMirrorCommonGammaMask KMirrorCommonGamma = 0
)

func (e KMirrorCommonGamma) String() string {
	switch e {
	case KMirrorCommonGammaMask:
		return "KMirrorCommonGammaMask"
	default:
		return fmt.Sprintf("KMirrorCommonGamma(%d)", e)
	}
}

type KMode uint

const (
	KModeBuiltIn                KMode = 0
	KModeDefault                KMode = 0
	KModeInterlaced             KMode = 0
	KModeNotGraphicsQuality     KMode = 0
	KModeNotPreset              KMode = 0
	KModeNotResize              KMode = 0
	KModeRequiresPan            KMode = 0
	KModeSafe                   KMode = 0
	KModeShowNever              KMode = 0
	KModeShowNow                KMode = 0
	KModeSimulscan              KMode = 0
	KModeStretched              KMode = 0
	KModeValid                  KMode = 0
	KModeValidateAgainstDisplay KMode = 0
)

func (e KMode) String() string {
	switch e {
	case KModeBuiltIn:
		return "KModeBuiltIn"
	default:
		return fmt.Sprintf("KMode(%d)", e)
	}
}

type KModePageControl uint

const (
	KModePageControlChangeableValues KModePageControl = 0
	KModePageControlCurrentValues    KModePageControl = 0
	KModePageControlDefaultValues    KModePageControl = 0
	KModePageControlSavedValues      KModePageControl = 0
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

type KModifier uint

const (
	KModifier_DidKeyUp          KModifier = 0
	KModifier_DidPerformModifiy KModifier = 0
	KModifier_Locked            KModifier = 0
)

func (e KModifier) String() string {
	switch e {
	case KModifier_DidKeyUp:
		return "KModifier_DidKeyUp"
	default:
		return fmt.Sprintf("KModifier(%d)", e)
	}
}

type KNVRAM uint

const (
	KNVRAMProperty KNVRAM = 0
)

func (e KNVRAM) String() string {
	switch e {
	case KNVRAMProperty:
		return "KNVRAMProperty"
	default:
		return fmt.Sprintf("KNVRAM(%d)", e)
	}
}

type KNanosecondScale uint

const (
	// KMicrosecondScale: # Discussion
	KMicrosecondScale KNanosecondScale = 0
	// KMillisecondScale: # Discussion
	KMillisecondScale KNanosecondScale = 0
	// KNanosecondScaleValue: # Discussion
	KNanosecondScaleValue KNanosecondScale = 0
	// KSecondScale: # Discussion
	KSecondScale KNanosecondScale = 0
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
	KNdrvTypeIsBusBridge    KNdrvTypeIs = 0
	KNdrvTypeIsGeneric      KNdrvTypeIs = 0
	KNdrvTypeIsNetworking   KNdrvTypeIs = 0
	KNdrvTypeIsParallel     KNdrvTypeIs = 0
	KNdrvTypeIsSerial       KNdrvTypeIs = 0
	KNdrvTypeIsSound        KNdrvTypeIs = 0
	KNdrvTypeIsVideo        KNdrvTypeIs = 0
)

func (e KNdrvTypeIs) String() string {
	switch e {
	case KNdrvTypeIsBlockStorage:
		return "KNdrvTypeIsBlockStorage"
	default:
		return fmt.Sprintf("KNdrvTypeIs(%d)", e)
	}
}

type KNil uint

const (
	KNilOptions KNil = 0
)

func (e KNil) String() string {
	switch e {
	case KNilOptions:
		return "KNilOptions"
	default:
		return fmt.Sprintf("KNil(%d)", e)
	}
}

type KNuDCL uint

const (
	KNuDCLDynamic              KNuDCL = 0
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

type KNumPort uint

const (
	KNumPortBytes KNumPort = 0
)

func (e KNumPort) String() string {
	switch e {
	case KNumPortBytes:
		return "KNumPortBytes"
	default:
		return fmt.Sprintf("KNumPort(%d)", e)
	}
}

type KOFVariablePerm uint

const (
	KOFVariablePermKernelOnly KOFVariablePerm = 0
	KOFVariablePermUserRead   KOFVariablePerm = 0
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
	KOFVariableTypeData    KOFVariableType = 0
	KOFVariableTypeNumber  KOFVariableType = 0
	KOFVariableTypeString  KOFVariableType = 0
)

func (e KOFVariableType) String() string {
	switch e {
	case KOFVariableTypeBoolean:
		return "KOFVariableTypeBoolean"
	default:
		return fmt.Sprintf("KOFVariableType(%d)", e)
	}
}

type KOSAsyncRef uint

const (
	KOSAsyncRefCount KOSAsyncRef = 0
	KOSAsyncRefSize  KOSAsyncRef = 0
)

func (e KOSAsyncRef) String() string {
	switch e {
	case KOSAsyncRefCount:
		return "KOSAsyncRefCount"
	default:
		return fmt.Sprintf("KOSAsyncRef(%d)", e)
	}
}

type KOSAsyncRef64 uint

const (
	KOSAsyncRef64Count KOSAsyncRef64 = 0
	KOSAsyncRef64Size  KOSAsyncRef64 = 0
)

func (e KOSAsyncRef64) String() string {
	switch e {
	case KOSAsyncRef64Count:
		return "KOSAsyncRef64Count"
	default:
		return fmt.Sprintf("KOSAsyncRef64(%d)", e)
	}
}

type KOSClassCan uint

const (
	KOSClassCanRemote KOSClassCan = 0
)

func (e KOSClassCan) String() string {
	switch e {
	case KOSClassCanRemote:
		return "KOSClassCanRemote"
	default:
		return fmt.Sprintf("KOSClassCan(%d)", e)
	}
}

type KOSNotificationMessageID uint

const (
	KMaxAsyncArgs                 KOSNotificationMessageID = 0
	KOSAsyncCompleteMessageID     KOSNotificationMessageID = 0
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

type KOSStringNo uint

const (
	KOSStringNoCopy KOSStringNo = 0
)

func (e KOSStringNo) String() string {
	switch e {
	case KOSStringNoCopy:
		return "KOSStringNoCopy"
	default:
		return fmt.Sprintf("KOSStringNo(%d)", e)
	}
}

type KPCI2PCIBridge uint

const (
	KPCI2PCIBridgeControl KPCI2PCIBridge = 0
)

func (e KPCI2PCIBridge) String() string {
	switch e {
	case KPCI2PCIBridgeControl:
		return "KPCI2PCIBridgeControl"
	default:
		return fmt.Sprintf("KPCI2PCIBridge(%d)", e)
	}
}

type KPCIPMC uint

const (
	KPCIPMCD1Support            KPCIPMC = 0
	KPCIPMCD2Support            KPCIPMC = 0
	KPCIPMCD3Support            KPCIPMC = 0
	KPCIPMCPMESupportFromD0     KPCIPMC = 0
	KPCIPMCPMESupportFromD1     KPCIPMC = 0
	KPCIPMCPMESupportFromD2     KPCIPMC = 0
	KPCIPMCPMESupportFromD3Cold KPCIPMC = 0
	KPCIPMCPMESupportFromD3Hot  KPCIPMC = 0
)

func (e KPCIPMC) String() string {
	switch e {
	case KPCIPMCD1Support:
		return "KPCIPMCD1Support"
	default:
		return fmt.Sprintf("KPCIPMC(%d)", e)
	}
}

type KPCIPMCS uint

const (
	KPCIPMCSDefaultEnableBits KPCIPMCS = 0
	KPCIPMCSPMEDisableInS3    KPCIPMCS = 0
	KPCIPMCSPMEEnable         KPCIPMCS = 0
	KPCIPMCSPMEStatus         KPCIPMCS = 0
	KPCIPMCSPMEWakeReason     KPCIPMCS = 0
	KPCIPMCSPowerStateD0      KPCIPMCS = 0
	KPCIPMCSPowerStateD1      KPCIPMCS = 0
	KPCIPMCSPowerStateD2      KPCIPMCS = 0
	KPCIPMCSPowerStateD3      KPCIPMCS = 0
	KPCIPMCSPowerStateMask    KPCIPMCS = 0
)

func (e KPCIPMCS) String() string {
	switch e {
	case KPCIPMCSDefaultEnableBits:
		return "KPCIPMCSDefaultEnableBits"
	default:
		return fmt.Sprintf("KPCIPMCS(%d)", e)
	}
}

type KPE uint

const (
	KPECommandKey   KPE = 0
	KPEControlKey   KPE = 0
	KPEOptionKey    KPE = 0
	KPERawInput     KPE = 0
	KPEReadTOD      KPE = 0
	KPEShiftKey     KPE = 0
	KPEWaitForInput KPE = 0
	KPEWriteTOD     KPE = 0
)

func (e KPE) String() string {
	switch e {
	case KPECommandKey:
		return "KPECommandKey"
	default:
		return fmt.Sprintf("KPE(%d)", e)
	}
}

type KPEHaltCP uint

const (
	KPEHaltCPU KPEHaltCP = 0
)

func (e KPEHaltCP) String() string {
	switch e {
	case KPEHaltCPU:
		return "KPEHaltCPU"
	default:
		return fmt.Sprintf("KPEHaltCP(%d)", e)
	}
}

type KPEScale uint

const (
	KPEScaleFactor1x      KPEScale = 0
	KPEScaleFactor2x      KPEScale = 0
	KPEScaleFactorUnknown KPEScale = 0
)

func (e KPEScale) String() string {
	switch e {
	case KPEScaleFactor1x:
		return "KPEScaleFactor1x"
	default:
		return fmt.Sprintf("KPEScale(%d)", e)
	}
}

type KPM uint

const (
	KPMGeneralAggressiveness KPM = 0
	KPMMinutesToSleep        KPM = 0
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
	KCompoundDeviceBit   KPerPortSwitchingBit = 0
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
	KThirdOrderRanking   KPeripheralDeviceTypeNoMatch = 0
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
	KPowerStateNeedsRefresh                 KPowerState = 0
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
	KPrdRootHubApple  KPrdRootHub = 0
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
	KRBCNumPowerStates        KRBC = 0
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
	KRBCCmd_FORMAT_UNIT     KRBCCmd = 0
	KRBCCmd_READ_CAPACITY   KRBCCmd = 0
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

type KRangeSupports uint

const (
	KRangeSupportsCompositeSyncBit   KRangeSupports = 0
	KRangeSupportsCompositeSyncMask  KRangeSupports = 0
	KRangeSupportsSeperateSyncsBit   KRangeSupports = 0
	KRangeSupportsSeperateSyncsMask  KRangeSupports = 0
	KRangeSupportsSyncOnGreenBit     KRangeSupports = 0
	KRangeSupportsSyncOnGreenMask    KRangeSupports = 0
	KRangeSupportsVSyncSerrationBit  KRangeSupports = 0
	KRangeSupportsVSyncSerrationMask KRangeSupports = 0
)

func (e KRangeSupports) String() string {
	switch e {
	case KRangeSupportsCompositeSyncBit:
		return "KRangeSupportsCompositeSyncBit"
	default:
		return fmt.Sprintf("KRangeSupports(%d)", e)
	}
}

type KRangeSupportsSignal uint

const (
	KRangeSupportsSignal_0700_0000_Bit  KRangeSupportsSignal = 0
	KRangeSupportsSignal_0700_0000_Mask KRangeSupportsSignal = 0
	KRangeSupportsSignal_0700_0300_Bit  KRangeSupportsSignal = 0
	KRangeSupportsSignal_0700_0300_Mask KRangeSupportsSignal = 0
	KRangeSupportsSignal_0714_0286_Bit  KRangeSupportsSignal = 0
	KRangeSupportsSignal_0714_0286_Mask KRangeSupportsSignal = 0
	KRangeSupportsSignal_1000_0400_Bit  KRangeSupportsSignal = 0
	KRangeSupportsSignal_1000_0400_Mask KRangeSupportsSignal = 0
)

func (e KRangeSupportsSignal) String() string {
	switch e {
	case KRangeSupportsSignal_0700_0000_Bit:
		return "KRangeSupportsSignal_0700_0000_Bit"
	default:
		return fmt.Sprintf("KRangeSupportsSignal(%d)", e)
	}
}

type KReg uint

const (
	KRegModifierMask          KReg = 0
	KRegNameSpaceModifierMask KReg = 0
	KRegNoModifiers           KReg = 0
	KRegUniversalModifierMask KReg = 0
)

func (e KReg) String() string {
	switch e {
	case KRegModifierMask:
		return "KRegModifierMask"
	default:
		return fmt.Sprintf("KReg(%d)", e)
	}
}

type KRegCStrMaxEntryName uint

const (
	KRegCStrMaxEntryNameLength KRegCStrMaxEntryName = 0
)

func (e KRegCStrMaxEntryName) String() string {
	switch e {
	case KRegCStrMaxEntryNameLength:
		return "KRegCStrMaxEntryNameLength"
	default:
		return fmt.Sprintf("KRegCStrMaxEntryName(%d)", e)
	}
}

type KRegEntryName uint

const (
	KRegEntryNameTerminator KRegEntryName = 0
)

func (e KRegEntryName) String() string {
	switch e {
	case KRegEntryNameTerminator:
		return "KRegEntryNameTerminator"
	default:
		return fmt.Sprintf("KRegEntryName(%d)", e)
	}
}

type KRegIter uint

const (
	KRegIterChildren KRegIter = 0
	KRegIterRoot     KRegIter = 0
)

func (e KRegIter) String() string {
	switch e {
	case KRegIterChildren:
		return "KRegIterChildren"
	default:
		return fmt.Sprintf("KRegIter(%d)", e)
	}
}

type KRegMaxPropertyName uint

const (
	KRegMaxPropertyNameLength KRegMaxPropertyName = 0
)

func (e KRegMaxPropertyName) String() string {
	switch e {
	case KRegMaxPropertyNameLength:
		return "KRegMaxPropertyNameLength"
	default:
		return fmt.Sprintf("KRegMaxPropertyName(%d)", e)
	}
}

type KRegMaximumPropertyName uint

const (
	KRegMaximumPropertyNameLength KRegMaximumPropertyName = 0
)

func (e KRegMaximumPropertyName) String() string {
	switch e {
	case KRegMaximumPropertyNameLength:
		return "KRegMaximumPropertyNameLength"
	default:
		return fmt.Sprintf("KRegMaximumPropertyName(%d)", e)
	}
}

type KRegPropertyValueIsSavedTo uint

const (
	KRegPropertyValueIsSavedToDisk  KRegPropertyValueIsSavedTo = 0
	KRegPropertyValueIsSavedToNVRAM KRegPropertyValueIsSavedTo = 0
)

func (e KRegPropertyValueIsSavedTo) String() string {
	switch e {
	case KRegPropertyValueIsSavedToDisk:
		return "KRegPropertyValueIsSavedToDisk"
	default:
		return fmt.Sprintf("KRegPropertyValueIsSavedTo(%d)", e)
	}
}

type KRequestDirection uint

const (
	KRequestDirectionIn  KRequestDirection = 0
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
	KRequestRecipientDevice   KRequestRecipient = 0
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

type KRequestType uint

const (
	KRequestTypeClass KRequestType = 0
)

func (e KRequestType) String() string {
	switch e {
	case KRequestTypeClass:
		return "KRequestTypeClass"
	default:
		return fmt.Sprintf("KRequestType(%d)", e)
	}
}

type KResolutionHasMultipleDepth uint

const (
	KResolutionHasMultipleDepthSizes KResolutionHasMultipleDepth = 0
)

func (e KResolutionHasMultipleDepth) String() string {
	switch e {
	case KResolutionHasMultipleDepthSizes:
		return "KResolutionHasMultipleDepthSizes"
	default:
		return fmt.Sprintf("KResolutionHasMultipleDepth(%d)", e)
	}
}

type KRootDomainSleepNotSupported uint

const (
	KFrameBufferDeepSleepSupported    KRootDomainSleepNotSupported = 0
	KPCICantSleep                     KRootDomainSleepNotSupported = 0
	KRootDomainSleepNotSupportedValue KRootDomainSleepNotSupported = 0
	KRootDomainSleepSupported         KRootDomainSleepNotSupported = 0
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
	KSBCNumPowerStates    KSBC = 0
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
	KSBCCmd_CHANGE_DEFINITION            KSBCCmd = 0
	KSBCCmd_COMPARE                      KSBCCmd = 0
	KSBCCmd_COPY                         KSBCCmd = 0
	KSBCCmd_COPY_AND_VERIFY              KSBCCmd = 0
	KSBCCmd_FORMAT_UNIT                  KSBCCmd = 0
	KSBCCmd_INQUIRY                      KSBCCmd = 0
	KSBCCmd_LOCK_UNLOCK_CACHE            KSBCCmd = 0
	KSBCCmd_LOG_SELECT                   KSBCCmd = 0
	KSBCCmd_LOG_SENSE                    KSBCCmd = 0
	KSBCCmd_MODE_SELECT_10               KSBCCmd = 0
	KSBCCmd_MODE_SELECT_6                KSBCCmd = 0
	KSBCCmd_MODE_SENSE_10                KSBCCmd = 0
	KSBCCmd_MODE_SENSE_6                 KSBCCmd = 0
	KSBCCmd_MOVE_MEDIUM_ATTACHED         KSBCCmd = 0
	KSBCCmd_PERSISTENT_RESERVE_IN        KSBCCmd = 0
	KSBCCmd_PERSISTENT_RESERVE_OUT       KSBCCmd = 0
	KSBCCmd_PREFETCH                     KSBCCmd = 0
	KSBCCmd_PREVENT_ALLOW_MEDIUM_REMOVAL KSBCCmd = 0
	KSBCCmd_READ_10                      KSBCCmd = 0
	KSBCCmd_READ_12                      KSBCCmd = 0
	KSBCCmd_READ_6                       KSBCCmd = 0
	KSBCCmd_READ_BUFFER                  KSBCCmd = 0
	KSBCCmd_READ_CAPACITY                KSBCCmd = 0
	KSBCCmd_READ_DEFECT_DATA_10          KSBCCmd = 0
	KSBCCmd_READ_DEFECT_DATA_12          KSBCCmd = 0
	KSBCCmd_READ_ELEMENT_STATUS_ATTACHED KSBCCmd = 0
	KSBCCmd_READ_GENERATION              KSBCCmd = 0
	KSBCCmd_READ_LONG                    KSBCCmd = 0
	KSBCCmd_READ_UPDATED_BLOCK_10        KSBCCmd = 0
	KSBCCmd_REASSIGN_BLOCKS              KSBCCmd = 0
	KSBCCmd_REBUILD                      KSBCCmd = 0
	KSBCCmd_RECEIVE_DIAGNOSTICS_RESULTS  KSBCCmd = 0
	KSBCCmd_REGENERATE                   KSBCCmd = 0
	KSBCCmd_RELEASE_10                   KSBCCmd = 0
	KSBCCmd_RELEASE_6                    KSBCCmd = 0
	KSBCCmd_REPORT_LUNS                  KSBCCmd = 0
	KSBCCmd_REQUEST_SENSE                KSBCCmd = 0
	KSBCCmd_RESERVE_10                   KSBCCmd = 0
	KSBCCmd_RESERVE_6                    KSBCCmd = 0
	KSBCCmd_REZERO_UNIT                  KSBCCmd = 0
	KSBCCmd_SEARCH_DATA_EQUAL_10         KSBCCmd = 0
	KSBCCmd_SEARCH_DATA_HIGH_10          KSBCCmd = 0
	KSBCCmd_SEARCH_DATA_LOW_10           KSBCCmd = 0
	KSBCCmd_SEEK_10                      KSBCCmd = 0
	KSBCCmd_SEEK_6                       KSBCCmd = 0
	KSBCCmd_SEND_DIAGNOSTICS             KSBCCmd = 0
	KSBCCmd_SET_LIMITS_10                KSBCCmd = 0
	KSBCCmd_SET_LIMITS_12                KSBCCmd = 0
	KSBCCmd_START_STOP_UNIT              KSBCCmd = 0
	KSBCCmd_SYNCHRONIZE_CACHE            KSBCCmd = 0
	KSBCCmd_TEST_UNIT_READY              KSBCCmd = 0
	KSBCCmd_UPDATE_BLOCK                 KSBCCmd = 0
	KSBCCmd_VERIFY_10                    KSBCCmd = 0
	KSBCCmd_WRITE_10                     KSBCCmd = 0
	KSBCCmd_WRITE_12                     KSBCCmd = 0
	KSBCCmd_WRITE_6                      KSBCCmd = 0
	KSBCCmd_WRITE_AND_VERIFY_10          KSBCCmd = 0
	KSBCCmd_WRITE_AND_VERIFY_12          KSBCCmd = 0
	KSBCCmd_WRITE_BUFFER                 KSBCCmd = 0
	KSBCCmd_WRITE_LONG                   KSBCCmd = 0
	KSBCCmd_WRITE_SAME                   KSBCCmd = 0
	KSBCCmd_XDREAD                       KSBCCmd = 0
	KSBCCmd_XDWRITE                      KSBCCmd = 0
	KSBCCmd_XDWRITE_EXTENDED             KSBCCmd = 0
	KSBCCmd_XPWRITE                      KSBCCmd = 0
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
	KSBCWOCmd_LOG_SELECT        KSBCWOCmd = 0
)

func (e KSBCWOCmd) String() string {
	switch e {
	case KSBCWOCmd_CHANGE_DEFINITION:
		return "KSBCWOCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KSBCWOCmd(%d)", e)
	}
}

type KSCCCmd uint

const (
	KSCCCmd_MAINTENANCE_IN         KSCCCmd = 0
	KSCCCmd_MAINTENANCE_OUT        KSCCCmd = 0
	KSCCCmd_MODE_SELECT_10         KSCCCmd = 0
	KSCCCmd_MODE_SELECT_6          KSCCCmd = 0
	KSCCCmd_MODE_SENSE_10          KSCCCmd = 0
	KSCCCmd_MODE_SENSE_6           KSCCCmd = 0
	KSCCCmd_PERSISTENT_RESERVE_IN  KSCCCmd = 0
	KSCCCmd_PERSISTENT_RESERVE_OUT KSCCCmd = 0
	KSCCCmd_PORT_STATUS            KSCCCmd = 0
	KSCCCmd_REDUNDANCY_GROUP_IN    KSCCCmd = 0
	KSCCCmd_REDUNDANCY_GROUP_OUT   KSCCCmd = 0
	KSCCCmd_RELEASE_10             KSCCCmd = 0
	KSCCCmd_RELEASE_6              KSCCCmd = 0
	KSCCCmd_REPORT_LUNS            KSCCCmd = 0
	KSCCCmd_REQUEST_SENSE          KSCCCmd = 0
	KSCCCmd_RESERVE_10             KSCCCmd = 0
	KSCCCmd_RESERVE_6              KSCCCmd = 0
	KSCCCmd_SEND_DIAGNOSTICS       KSCCCmd = 0
	KSCCCmd_SPARE_IN               KSCCCmd = 0
	KSCCCmd_SPARE_OUT              KSCCCmd = 0
)

func (e KSCCCmd) String() string {
	switch e {
	case KSCCCmd_MAINTENANCE_IN:
		return "KSCCCmd_MAINTENANCE_IN"
	default:
		return fmt.Sprintf("KSCCCmd(%d)", e)
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
	KSCSICmdVariableLengthCDB              KSCSICmd = 0
	KSCSICmd_ACCESS_CONTROL_IN             KSCSICmd = 0
	KSCSICmd_ACCESS_CONTROL_OUT            KSCSICmd = 0
	KSCSICmd_BLANK                         KSCSICmd = 0
	KSCSICmd_CHANGE_DEFINITION             KSCSICmd = 0
	KSCSICmd_CLOSE_TRACK_SESSION           KSCSICmd = 0
	KSCSICmd_COMPARE                       KSCSICmd = 0
	KSCSICmd_COPY                          KSCSICmd = 0
	KSCSICmd_COPY_AND_VERIFY               KSCSICmd = 0
	KSCSICmd_ERASE_10                      KSCSICmd = 0
	KSCSICmd_ERASE_12                      KSCSICmd = 0
	KSCSICmd_EXTENDED_COPY                 KSCSICmd = 0
	KSCSICmd_FORMAT_UNIT                   KSCSICmd = 0
	KSCSICmd_GET_CONFIGURATION             KSCSICmd = 0
	KSCSICmd_GET_EVENT_STATUS_NOTIFICATION KSCSICmd = 0
	KSCSICmd_GET_PERFORMANCE               KSCSICmd = 0
	KSCSICmd_INQUIRY                       KSCSICmd = 0
	KSCSICmd_LOAD_UNLOAD_MEDIUM            KSCSICmd = 0
	KSCSICmd_LOCK_UNLOCK_CACHE             KSCSICmd = 0
	KSCSICmd_LOCK_UNLOCK_CACHE_16          KSCSICmd = 0
	KSCSICmd_LOG_SELECT                    KSCSICmd = 0
	KSCSICmd_LOG_SENSE                     KSCSICmd = 0
	KSCSICmd_MAINTENANCE_IN                KSCSICmd = 0
	KSCSICmd_MAINTENANCE_OUT               KSCSICmd = 0
	KSCSICmd_MECHANISM_STATUS              KSCSICmd = 0
	KSCSICmd_MEDIUM_SCAN                   KSCSICmd = 0
	KSCSICmd_MODE_SELECT_10                KSCSICmd = 0
	KSCSICmd_MODE_SELECT_6                 KSCSICmd = 0
	KSCSICmd_MODE_SENSE_10                 KSCSICmd = 0
	KSCSICmd_MODE_SENSE_6                  KSCSICmd = 0
	KSCSICmd_MOVE_MEDIUM_ATTACHED          KSCSICmd = 0
	KSCSICmd_PAUSE_RESUME                  KSCSICmd = 0
	KSCSICmd_PERSISTENT_RESERVE_IN         KSCSICmd = 0
	KSCSICmd_PERSISTENT_RESERVE_OUT        KSCSICmd = 0
	KSCSICmd_PLAY_AUDIO_10                 KSCSICmd = 0
	KSCSICmd_PLAY_AUDIO_12                 KSCSICmd = 0
	KSCSICmd_PLAY_AUDIO_MSF                KSCSICmd = 0
	KSCSICmd_PLAY_AUDIO_TRACK_INDEX        KSCSICmd = 0
	KSCSICmd_PLAY_CD                       KSCSICmd = 0
	KSCSICmd_PLAY_RELATIVE_10              KSCSICmd = 0
	KSCSICmd_PLAY_RELATIVE_12              KSCSICmd = 0
	KSCSICmd_PREFETCH                      KSCSICmd = 0
	KSCSICmd_PREFETCH_16                   KSCSICmd = 0
	KSCSICmd_PREVENT_ALLOW_MEDIUM_REMOVAL  KSCSICmd = 0
	KSCSICmd_READ_10                       KSCSICmd = 0
	KSCSICmd_READ_12                       KSCSICmd = 0
	KSCSICmd_READ_16                       KSCSICmd = 0
	KSCSICmd_READ_6                        KSCSICmd = 0
	KSCSICmd_READ_ATTRIBUTE                KSCSICmd = 0
	KSCSICmd_READ_BUFFER                   KSCSICmd = 0
	KSCSICmd_READ_BUFFER_CAPACITY          KSCSICmd = 0
	KSCSICmd_READ_CAPACITY                 KSCSICmd = 0
	KSCSICmd_READ_CD                       KSCSICmd = 0
	KSCSICmd_READ_CD_MSF                   KSCSICmd = 0
	KSCSICmd_READ_DEFECT_DATA_10           KSCSICmd = 0
	KSCSICmd_READ_DEFECT_DATA_12           KSCSICmd = 0
	KSCSICmd_READ_DISC_INFORMATION         KSCSICmd = 0
	KSCSICmd_READ_DISC_STRUCTURE           KSCSICmd = 0
	KSCSICmd_READ_DVD_STRUCTURE            KSCSICmd = 0
	KSCSICmd_READ_ELEMENT_STATUS_ATTACHED  KSCSICmd = 0
	KSCSICmd_READ_FORMAT_CAPACITIES        KSCSICmd = 0
	KSCSICmd_READ_GENERATION               KSCSICmd = 0
	KSCSICmd_READ_HEADER                   KSCSICmd = 0
	KSCSICmd_READ_LONG                     KSCSICmd = 0
	KSCSICmd_READ_MASTER_CUE               KSCSICmd = 0
	KSCSICmd_READ_SUB_CHANNEL              KSCSICmd = 0
	KSCSICmd_READ_TOC_PMA_ATIP             KSCSICmd = 0
	KSCSICmd_READ_TRACK_INFORMATION        KSCSICmd = 0
	KSCSICmd_READ_UPDATED_BLOCK_10         KSCSICmd = 0
	KSCSICmd_REASSIGN_BLOCKS               KSCSICmd = 0
	KSCSICmd_REBUILD                       KSCSICmd = 0
	KSCSICmd_RECEIVE                       KSCSICmd = 0
	KSCSICmd_RECEIVE_COPY_RESULTS          KSCSICmd = 0
	KSCSICmd_RECEIVE_DIAGNOSTICS_RESULTS   KSCSICmd = 0
	KSCSICmd_REDUNDANCY_GROUP_IN           KSCSICmd = 0
	KSCSICmd_REDUNDANCY_GROUP_OUT          KSCSICmd = 0
	KSCSICmd_REGENERATE                    KSCSICmd = 0
	KSCSICmd_RELEASE_10                    KSCSICmd = 0
	KSCSICmd_RELEASE_6                     KSCSICmd = 0
	KSCSICmd_REPAIR_TRACK                  KSCSICmd = 0
	KSCSICmd_REPORT_DEVICE_IDENTIFIER      KSCSICmd = 0
	KSCSICmd_REPORT_KEY                    KSCSICmd = 0
	KSCSICmd_REPORT_LUNS                   KSCSICmd = 0
	KSCSICmd_REQUEST_SENSE                 KSCSICmd = 0
	KSCSICmd_RESERVE_10                    KSCSICmd = 0
	KSCSICmd_RESERVE_6                     KSCSICmd = 0
	KSCSICmd_RESERVE_TRACK                 KSCSICmd = 0
	KSCSICmd_REZERO_UNIT                   KSCSICmd = 0
	KSCSICmd_SCAN_MMC                      KSCSICmd = 0
	KSCSICmd_SEARCH_DATA_EQUAL_10          KSCSICmd = 0
	KSCSICmd_SEARCH_DATA_EQUAL_12          KSCSICmd = 0
	KSCSICmd_SEARCH_DATA_HIGH_10           KSCSICmd = 0
	KSCSICmd_SEARCH_DATA_HIGH_12           KSCSICmd = 0
	KSCSICmd_SEARCH_DATA_LOW_10            KSCSICmd = 0
	KSCSICmd_SEARCH_DATA_LOW_12            KSCSICmd = 0
	KSCSICmd_SEEK_10                       KSCSICmd = 0
	KSCSICmd_SEEK_6                        KSCSICmd = 0
	KSCSICmd_SEND                          KSCSICmd = 0
	KSCSICmd_SEND_CUE_SHEET                KSCSICmd = 0
	KSCSICmd_SEND_DIAGNOSTICS              KSCSICmd = 0
	KSCSICmd_SEND_DVD_STRUCTURE            KSCSICmd = 0
	KSCSICmd_SEND_EVENT                    KSCSICmd = 0
	KSCSICmd_SEND_KEY                      KSCSICmd = 0
	KSCSICmd_SEND_OPC_INFORMATION          KSCSICmd = 0
	KSCSICmd_SERVICE_ACTION_IN             KSCSICmd = 0
	KSCSICmd_SERVICE_ACTION_OUT            KSCSICmd = 0
	KSCSICmd_SET_CD_SPEED                  KSCSICmd = 0
	KSCSICmd_SET_DEVICE_IDENTIFIER         KSCSICmd = 0
	KSCSICmd_SET_LIMITS_10                 KSCSICmd = 0
	KSCSICmd_SET_LIMITS_12                 KSCSICmd = 0
	KSCSICmd_SET_READ_AHEAD                KSCSICmd = 0
	KSCSICmd_SET_STREAMING                 KSCSICmd = 0
	KSCSICmd_SPARE_IN                      KSCSICmd = 0
	KSCSICmd_SPARE_OUT                     KSCSICmd = 0
	KSCSICmd_START_STOP_UNIT               KSCSICmd = 0
	KSCSICmd_STOP_PLAY_SCAN                KSCSICmd = 0
	KSCSICmd_SYNCHRONIZE_CACHE             KSCSICmd = 0
	KSCSICmd_SYNCHRONIZE_CACHE_16          KSCSICmd = 0
	KSCSICmd_TEST_UNIT_READY               KSCSICmd = 0
	KSCSICmd_UNMAP                         KSCSICmd = 0
	KSCSICmd_UPDATE_BLOCK                  KSCSICmd = 0
	KSCSICmd_VERIFY_10                     KSCSICmd = 0
	KSCSICmd_VERIFY_12                     KSCSICmd = 0
	KSCSICmd_VERIFY_16                     KSCSICmd = 0
	KSCSICmd_VOLUME_SET_IN                 KSCSICmd = 0
	KSCSICmd_VOLUME_SET_OUT                KSCSICmd = 0
	KSCSICmd_VendorSpecific_End            KSCSICmd = 0
	KSCSICmd_VendorSpecific_Start          KSCSICmd = 0
	KSCSICmd_WRITE_10                      KSCSICmd = 0
	KSCSICmd_WRITE_12                      KSCSICmd = 0
	KSCSICmd_WRITE_16                      KSCSICmd = 0
	KSCSICmd_WRITE_6                       KSCSICmd = 0
	KSCSICmd_WRITE_AND_VERIFY_10           KSCSICmd = 0
	KSCSICmd_WRITE_AND_VERIFY_12           KSCSICmd = 0
	KSCSICmd_WRITE_AND_VERIFY_16           KSCSICmd = 0
	KSCSICmd_WRITE_ATTRIBUTE               KSCSICmd = 0
	KSCSICmd_WRITE_BUFFER                  KSCSICmd = 0
	KSCSICmd_WRITE_LONG                    KSCSICmd = 0
	KSCSICmd_WRITE_SAME                    KSCSICmd = 0
	KSCSICmd_WRITE_SAME_16                 KSCSICmd = 0
	KSCSICmd_XDREAD                        KSCSICmd = 0
	KSCSICmd_XDWRITE                       KSCSICmd = 0
	KSCSICmd_XDWRITEREAD_10                KSCSICmd = 0
	KSCSICmd_XDWRITE_EXTENDED              KSCSICmd = 0
	KSCSICmd_XPWRITE                       KSCSICmd = 0
)

func (e KSCSICmd) String() string {
	switch e {
	case KSCSICmdVariableLengthCDB:
		return "KSCSICmdVariableLengthCDB"
	default:
		return fmt.Sprintf("KSCSICmd(%d)", e)
	}
}

type KSCSIControllerNotificationBus uint

const (
	KSCSIControllerNotificationBusReset KSCSIControllerNotificationBus = 0
)

func (e KSCSIControllerNotificationBus) String() string {
	switch e {
	case KSCSIControllerNotificationBusReset:
		return "KSCSIControllerNotificationBusReset"
	default:
		return fmt.Sprintf("KSCSIControllerNotificationBus(%d)", e)
	}
}

type KSCSIDataTransfer uint

const (
	// KSCSIDataTransfer_FromInitiatorToTarget: # Discussion
	KSCSIDataTransfer_FromInitiatorToTarget KSCSIDataTransfer = 0
	// KSCSIDataTransfer_FromTargetToInitiator: # Discussion
	KSCSIDataTransfer_FromTargetToInitiator KSCSIDataTransfer = 0
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
	KSCSIParallelMessage_ABORT_TASK                       KSCSIParallelMessage = 0
	KSCSIParallelMessage_ABORT_TASK_SET                   KSCSIParallelMessage = 0
	KSCSIParallelMessage_ACA                              KSCSIParallelMessage = 0
	KSCSIParallelMessage_CLEAR_ACA                        KSCSIParallelMessage = 0
	KSCSIParallelMessage_CLEAR_TASK_SET                   KSCSIParallelMessage = 0
	KSCSIParallelMessage_DISCONNECT                       KSCSIParallelMessage = 0
	KSCSIParallelMessage_EXTENDED_MESSAGE                 KSCSIParallelMessage = 0
	KSCSIParallelMessage_HEAD_OF_QUEUE                    KSCSIParallelMessage = 0
	KSCSIParallelMessage_IDENTIFY                         KSCSIParallelMessage = 0
	KSCSIParallelMessage_IGNORE_WIDE_RESIDUE              KSCSIParallelMessage = 0
	KSCSIParallelMessage_INITIATOR_DETECTED_ERROR         KSCSIParallelMessage = 0
	KSCSIParallelMessage_LINKED_COMMAND_COMPLETE          KSCSIParallelMessage = 0
	KSCSIParallelMessage_LOGICAL_UNIT_RESET               KSCSIParallelMessage = 0
	KSCSIParallelMessage_MESSAGE_PARITY_ERROR             KSCSIParallelMessage = 0
	KSCSIParallelMessage_MESSAGE_REJECT                   KSCSIParallelMessage = 0
	KSCSIParallelMessage_MODIFY_DATA_POINTER              KSCSIParallelMessage = 0
	KSCSIParallelMessage_NO_OPERATION                     KSCSIParallelMessage = 0
	KSCSIParallelMessage_ORDERED                          KSCSIParallelMessage = 0
	KSCSIParallelMessage_PARALLEL_PROTOCOL_REQUEST        KSCSIParallelMessage = 0
	KSCSIParallelMessage_QAS_REQUEST                      KSCSIParallelMessage = 0
	KSCSIParallelMessage_RESTORE_POINTERS                 KSCSIParallelMessage = 0
	KSCSIParallelMessage_SAVE_DATA_POINTER                KSCSIParallelMessage = 0
	KSCSIParallelMessage_SIMPLE                           KSCSIParallelMessage = 0
	KSCSIParallelMessage_SYNCHONOUS_DATA_TRANSFER_REQUEST KSCSIParallelMessage = 0
	KSCSIParallelMessage_TARGET_RESET                     KSCSIParallelMessage = 0
	KSCSIParallelMessage_TASK_COMPLETE                    KSCSIParallelMessage = 0
	KSCSIParallelMessage_WIDE_DATA_TRANSFER_REQUEST       KSCSIParallelMessage = 0
)

func (e KSCSIParallelMessage) String() string {
	switch e {
	case KSCSIParallelMessage_ABORT_TASK:
		return "KSCSIParallelMessage_ABORT_TASK"
	default:
		return fmt.Sprintf("KSCSIParallelMessage(%d)", e)
	}
}

type KSCSIParallelTaskControllerIDQueue uint

const (
	KSCSIParallelTaskControllerIDQueueHead KSCSIParallelTaskControllerIDQueue = 0
)

func (e KSCSIParallelTaskControllerIDQueue) String() string {
	switch e {
	case KSCSIParallelTaskControllerIDQueueHead:
		return "KSCSIParallelTaskControllerIDQueueHead"
	default:
		return fmt.Sprintf("KSCSIParallelTaskControllerIDQueue(%d)", e)
	}
}

type KSCSIPort uint

const (
	KSCSIPort_NotificationStatusChange KSCSIPort = 0
)

func (e KSCSIPort) String() string {
	switch e {
	case KSCSIPort_NotificationStatusChange:
		return "KSCSIPort_NotificationStatusChange"
	default:
		return fmt.Sprintf("KSCSIPort(%d)", e)
	}
}

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

type KSCSIProtocolLayerNumDefault uint

const (
	KSCSIProtocolLayerNumDefaultStates KSCSIProtocolLayerNumDefault = 0
)

func (e KSCSIProtocolLayerNumDefault) String() string {
	switch e {
	case KSCSIProtocolLayerNumDefaultStates:
		return "KSCSIProtocolLayerNumDefaultStates"
	default:
		return fmt.Sprintf("KSCSIProtocolLayerNumDefault(%d)", e)
	}
}

type KSCSIProtocolNotification uint

const (
	// KSCSIProtocolNotification_DeviceRemoved: # Discussion
	KSCSIProtocolNotification_DeviceRemoved KSCSIProtocolNotification = 0
)

func (e KSCSIProtocolNotification) String() string {
	switch e {
	case KSCSIProtocolNotification_DeviceRemoved:
		return "KSCSIProtocolNotification_DeviceRemoved"
	default:
		return fmt.Sprintf("KSCSIProtocolNotification(%d)", e)
	}
}

type KSCSIProtocolPowerState uint

const (
	// KSCSIProtocolPowerStateOff: # Discussion
	KSCSIProtocolPowerStateOff KSCSIProtocolPowerState = 0
)

func (e KSCSIProtocolPowerState) String() string {
	switch e {
	case KSCSIProtocolPowerStateOff:
		return "KSCSIProtocolPowerStateOff"
	default:
		return fmt.Sprintf("KSCSIProtocolPowerState(%d)", e)
	}
}

type KSCSIServiceAction uint

const (
	KSCSIServiceAction_READ_32             KSCSIServiceAction = 0
	KSCSIServiceAction_VERIFY_32           KSCSIServiceAction = 0
	KSCSIServiceAction_WRITE_32            KSCSIServiceAction = 0
	KSCSIServiceAction_WRITE_AND_VERIFY_32 KSCSIServiceAction = 0
	KSCSIServiceAction_WRITE_SAME_32       KSCSIServiceAction = 0
	KSCSIServiceAction_XDREAD_32           KSCSIServiceAction = 0
	KSCSIServiceAction_XDWRITEREAD_32      KSCSIServiceAction = 0
	KSCSIServiceAction_XDWRITE_32          KSCSIServiceAction = 0
	KSCSIServiceAction_XPWRITE_32          KSCSIServiceAction = 0
)

func (e KSCSIServiceAction) String() string {
	switch e {
	case KSCSIServiceAction_READ_32:
		return "KSCSIServiceAction_READ_32"
	default:
		return fmt.Sprintf("KSCSIServiceAction(%d)", e)
	}
}

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

type KSCSIServicesNotification uint

const (
	KSCSIServicesNotification_Resume KSCSIServicesNotification = 0
)

func (e KSCSIServicesNotification) String() string {
	switch e {
	case KSCSIServicesNotification_Resume:
		return "KSCSIServicesNotification_Resume"
	default:
		return fmt.Sprintf("KSCSIServicesNotification(%d)", e)
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
	KSCSITaskMode_Autosense        KSCSITaskMode = 0
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

type KSCSIUntaggedTask uint

const (
	// KSCSIUntaggedTaskIdentifier: # Discussion
	KSCSIUntaggedTaskIdentifier KSCSIUntaggedTask = 0
)

func (e KSCSIUntaggedTask) String() string {
	switch e {
	case KSCSIUntaggedTaskIdentifier:
		return "KSCSIUntaggedTaskIdentifier"
	default:
		return fmt.Sprintf("KSCSIUntaggedTask(%d)", e)
	}
}

type KSMCCmd uint

const (
	KSMCCmd_MODE_SELECT_6      KSMCCmd = 0
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
	KID_DRIVE       KSOFTRESET = 0
	KPACKET         KSOFTRESET = 0
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
	KSPCCmd_CHANGE_DEFINITION            KSPCCmd = 0
	KSPCCmd_COMPARE                      KSPCCmd = 0
	KSPCCmd_COPY                         KSPCCmd = 0
	KSPCCmd_COPY_AND_VERIFY              KSPCCmd = 0
	KSPCCmd_EXTENDED_COPY                KSPCCmd = 0
	KSPCCmd_INQUIRY                      KSPCCmd = 0
	KSPCCmd_LOG_SELECT                   KSPCCmd = 0
	KSPCCmd_LOG_SENSE                    KSPCCmd = 0
	KSPCCmd_MODE_SELECT_10               KSPCCmd = 0
	KSPCCmd_MODE_SELECT_6                KSPCCmd = 0
	KSPCCmd_MODE_SENSE_10                KSPCCmd = 0
	KSPCCmd_MODE_SENSE_6                 KSPCCmd = 0
	KSPCCmd_MOVE_MEDIUM_ATTACHED         KSPCCmd = 0
	KSPCCmd_PERSISTENT_RESERVE_IN        KSPCCmd = 0
	KSPCCmd_PERSISTENT_RESERVE_OUT       KSPCCmd = 0
	KSPCCmd_PREVENT_ALLOW_MEDIUM_REMOVAL KSPCCmd = 0
	KSPCCmd_READ_BUFFER                  KSPCCmd = 0
	KSPCCmd_READ_ELEMENT_STATUS_ATTACHED KSPCCmd = 0
	KSPCCmd_RECEIVE_COPY_RESULTS         KSPCCmd = 0
	KSPCCmd_RECEIVE_DIAGNOSTICS_RESULTS  KSPCCmd = 0
	KSPCCmd_RELEASE_10                   KSPCCmd = 0
	KSPCCmd_RELEASE_6                    KSPCCmd = 0
	KSPCCmd_REPORT_DEVICE_IDENTIFIER     KSPCCmd = 0
	KSPCCmd_REPORT_LUNS                  KSPCCmd = 0
	KSPCCmd_REQUEST_SENSE                KSPCCmd = 0
	KSPCCmd_RESERVE_10                   KSPCCmd = 0
	KSPCCmd_RESERVE_6                    KSPCCmd = 0
	KSPCCmd_SEND_DIAGNOSTICS             KSPCCmd = 0
	KSPCCmd_SET_DEVICE_IDENTIFIER        KSPCCmd = 0
	KSPCCmd_TEST_UNIT_READY              KSPCCmd = 0
	KSPCCmd_WRITE_BUFFER                 KSPCCmd = 0
)

func (e KSPCCmd) String() string {
	switch e {
	case KSPCCmd_CHANGE_DEFINITION:
		return "KSPCCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KSPCCmd(%d)", e)
	}
}

type KSPCModePage uint

const (
	// KSPCModePageAllPagesCode: # Discussion
	KSPCModePageAllPagesCode KSPCModePage = 0
	// KSPCModePagePowerConditionCode: # Discussion
	KSPCModePagePowerConditionCode KSPCModePage = 0
)

func (e KSPCModePage) String() string {
	switch e {
	case KSPCModePageAllPagesCode:
		return "KSPCModePageAllPagesCode"
	default:
		return fmt.Sprintf("KSPCModePage(%d)", e)
	}
}

type KSSCSeqCmd uint

const (
	KSSCSeqCmd_CHANGE_DEFINITION     KSSCSeqCmd = 0
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

type KSSHubPort uint

const (
	KSSHubPortLinkStateComplianceMode KSSHubPort = 0
	KSSHubPortLinkStateHotReset       KSSHubPort = 0
	KSSHubPortLinkStateLoopBack       KSSHubPort = 0
	KSSHubPortLinkStatePolling        KSSHubPort = 0
	KSSHubPortLinkStateRecovery       KSSHubPort = 0
	KSSHubPortLinkStateRxDetect       KSSHubPort = 0
	KSSHubPortLinkStateSSDisabled     KSSHubPort = 0
	KSSHubPortLinkStateSSInactive     KSSHubPort = 0
	KSSHubPortLinkStateU0             KSSHubPort = 0
	KSSHubPortLinkStateU1             KSSHubPort = 0
	KSSHubPortLinkStateU2             KSSHubPort = 0
	KSSHubPortLinkStateU3             KSSHubPort = 0
	KSSHubPortSpeed5Gbps              KSSHubPort = 0
)

func (e KSSHubPort) String() string {
	switch e {
	case KSSHubPortLinkStateComplianceMode:
		return "KSSHubPortLinkStateComplianceMode"
	default:
		return fmt.Sprintf("KSSHubPort(%d)", e)
	}
}

type KSSHubPortStatusConnectionBit int

const (
	KHubPortBeingReset             KSSHubPortStatusConnectionBit = 0
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
	KScaleCanBorderInsetOnlyMask  KScale = 0
	KScaleCanDownSamplePixelsMask KScale = 0
	KScaleCanRotateMask           KScale = 0
	KScaleCanScaleInterlacedMask  KScale = 0
	KScaleCanSupportInsetMask     KScale = 0
	KScaleCanUpSamplePixelsMask   KScale = 0
	KScaleStretchOnlyMask         KScale = 0
)

func (e KScale) String() string {
	switch e {
	case KScaleCanBorderInsetOnlyMask:
		return "KScaleCanBorderInsetOnlyMask"
	default:
		return fmt.Sprintf("KScale(%d)", e)
	}
}

type KScaleInvertX uint

const (
	KScaleInvertXMask KScaleInvertX = 0
)

func (e KScaleInvertX) String() string {
	switch e {
	case KScaleInvertXMask:
		return "KScaleInvertXMask"
	default:
		return fmt.Sprintf("KScaleInvertX(%d)", e)
	}
}

type KSecondsInAMinute uint

const (
	K5Minutes              KSecondsInAMinute = 0
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

type KSecondsPer uint

const (
	KSecondsPerHour KSecondsPer = 0
)

func (e KSecondsPer) String() string {
	switch e {
	case KSecondsPerHour:
		return "KSecondsPerHour"
	default:
		return fmt.Sprintf("KSecondsPer(%d)", e)
	}
}

type KSelfIDPacketSize uint

const (
	KMaxSelfIDs            KSelfIDPacketSize = 0
	KSelfIDPacketSizeValue KSelfIDPacketSize = 0
)

func (e KSelfIDPacketSize) String() string {
	switch e {
	case KMaxSelfIDs:
		return "KMaxSelfIDs"
	default:
		return fmt.Sprintf("KSelfIDPacketSize(%d)", e)
	}
}

type KSenseDefault uint

const (
	KSenseDefaultSize KSenseDefault = 0
)

func (e KSenseDefault) String() string {
	switch e {
	case KSenseDefaultSize:
		return "KSenseDefaultSize"
	default:
		return fmt.Sprintf("KSenseDefault(%d)", e)
	}
}

type KServiceCategory uint

const (
	KServiceCategoryADB           KServiceCategory = 0
	KServiceCategoryBlockStorage  KServiceCategory = 0
	KServiceCategoryDFM           KServiceCategory = 0
	KServiceCategoryDisplay       KServiceCategory = 0
	KServiceCategoryFileManager   KServiceCategory = 0
	KServiceCategoryGeneric       KServiceCategory = 0
	KServiceCategoryIDE           KServiceCategory = 0
	KServiceCategoryKeyboard      KServiceCategory = 0
	KServiceCategoryMotherBoard   KServiceCategory = 0
	KServiceCategoryNVRAM         KServiceCategory = 0
	KServiceCategoryNdrvDriver    KServiceCategory = 0
	KServiceCategoryOpenTransport KServiceCategory = 0
	KServiceCategoryPCI           KServiceCategory = 0
	KServiceCategoryPointing      KServiceCategory = 0
	KServiceCategoryPowerMgt      KServiceCategory = 0
	KServiceCategoryRTC           KServiceCategory = 0
	KServiceCategoryScsiSIM       KServiceCategory = 0
	KServiceCategorySound         KServiceCategory = 0
)

func (e KServiceCategory) String() string {
	switch e {
	case KServiceCategoryADB:
		return "KServiceCategoryADB"
	default:
		return fmt.Sprintf("KServiceCategory(%d)", e)
	}
}

type KSetCLUT uint

const (
	KSetCLUTByValue       KSetCLUT = 0
	KSetCLUTImmediately   KSetCLUT = 0
	KSetCLUTWithLuminance KSetCLUT = 0
)

func (e KSetCLUT) String() string {
	switch e {
	case KSetCLUTByValue:
		return "KSetCLUTByValue"
	default:
		return fmt.Sprintf("KSetCLUT(%d)", e)
	}
}

type KSetClutAt uint

const (
	KSetClutAtSetEntries KSetClutAt = 0
	KSetClutAtVBL        KSetClutAt = 0
)

func (e KSetClutAt) String() string {
	switch e {
	case KSetClutAtSetEntries:
		return "KSetClutAtSetEntries"
	default:
		return fmt.Sprintf("KSetClutAt(%d)", e)
	}
}

type KSetHubDepth uint

const (
	KGetPortErrorCount KSetHubDepth = 0
	KSetHubDepthValue  KSetHubDepth = 0
)

func (e KSetHubDepth) String() string {
	switch e {
	case KGetPortErrorCount:
		return "KGetPortErrorCount"
	default:
		return fmt.Sprintf("KSetHubDepth(%d)", e)
	}
}

type KSharedCacheAO uint

const (
	KSharedCacheAOT KSharedCacheAO = 0
)

func (e KSharedCacheAO) String() string {
	switch e {
	case KSharedCacheAOT:
		return "KSharedCacheAOT"
	default:
		return fmt.Sprintf("KSharedCacheAO(%d)", e)
	}
}

type KSizeOfATA uint

const (
	KSizeOfATAModelString    KSizeOfATA = 0
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

type KSymLink uint

const (
	KSymLinkCreator KSymLink = 0
)

func (e KSymLink) String() string {
	switch e {
	case KSymLinkCreator:
		return "KSymLinkCreator"
	default:
		return fmt.Sprintf("KSymLink(%d)", e)
	}
}

type KSync uint

const (
	KSyncAnalogBipolarMask              KSync = 0
	KSyncAnalogBipolarSRGBSyncMask      KSync = 0
	KSyncAnalogBipolarSerrateMask       KSync = 0
	KSyncAnalogCompositeMask            KSync = 0
	KSyncAnalogCompositeRGBSyncMask     KSync = 0
	KSyncAnalogCompositeSerrateMask     KSync = 0
	KSyncDigitalCompositeMask           KSync = 0
	KSyncDigitalCompositeMatchHSyncMask KSync = 0
	KSyncDigitalCompositeSerrateMask    KSync = 0
	KSyncDigitalHSyncPositiveMask       KSync = 0
	KSyncDigitalSeperateMask            KSync = 0
	KSyncDigitalVSyncPositiveMask       KSync = 0
	KSyncInterlaceMask                  KSync = 0
)

func (e KSync) String() string {
	switch e {
	case KSyncAnalogBipolarMask:
		return "KSyncAnalogBipolarMask"
	default:
		return fmt.Sprintf("KSync(%d)", e)
	}
}

type KSyncPositivePolarity uint

const (
	KSyncPositivePolarityBit KSyncPositivePolarity = 0
)

func (e KSyncPositivePolarity) String() string {
	switch e {
	case KSyncPositivePolarityBit:
		return "KSyncPositivePolarityBit"
	default:
		return fmt.Sprintf("KSyncPositivePolarity(%d)", e)
	}
}

type KTaskIs uint

const (
	KTaskIsTerminated KTaskIs = 0
)

func (e KTaskIs) String() string {
	switch e {
	case KTaskIsTerminated:
		return "KTaskIsTerminated"
	default:
		return fmt.Sprintf("KTaskIs(%d)", e)
	}
}

type KTheDescriptionSignature uint

const (
	KDriverDescriptionSignature   KTheDescriptionSignature = 0
	KTheDescriptionSignatureValue KTheDescriptionSignature = 0
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
	KThreadGroupApplication   KThreadGroup = 0
	KThreadGroupBestEffort    KThreadGroup = 0
	KThreadGroupCritical      KThreadGroup = 0
	KThreadGroupEfficient     KThreadGroup = 0
	KThreadGroupManaged       KThreadGroup = 0
	KThreadGroupStrictTimers  KThreadGroup = 0
	KThreadGroupUIApp         KThreadGroup = 0
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
	KThreadWaitNone        KThreadWait = 0
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
	KVideoBufferSizeErr             KTimingChangeRestrictedErr = 0
	KVideoCannotMirrorErr           KTimingChangeRestrictedErr = 0
	KVideoI2CBusyErr                KTimingChangeRestrictedErr = 0
	KVideoI2CReplyPendingErr        KTimingChangeRestrictedErr = 0
	KVideoI2CTransactionErr         KTimingChangeRestrictedErr = 0
	KVideoI2CTransactionTypeErr     KTimingChangeRestrictedErr = 0
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
	KInvertingEncoding        KTransparentEncoding = 0
	KTransparentEncodingValue KTransparentEncoding = 0
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
	KInvertingEncodedPixel   KTransparentEncodingShift = 0
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
	KUSB20ExtensionLPMSupported                      KUS = 0
	KUSB3HUBDesc                                     KUS = 0
	KUSB3HubDescriptorType                           KUS = 0
	KUSBAdaptiveIsocSyncType                         KUS = 0
	KUSBAsynchronousIsocSyncType                     KUS = 0
	KUSBDataIsocUsageType                            KUS = 0
	KUSBDeviceCapabilityValue                        KUS = 0
	KUSBEndpointDirectionIn                          KUS = 0
	KUSBEndpointDirectionOut                         KUS = 0
	KUSBEndpointbmAttributesSynchronizationTypeMask  KUS = 0
	KUSBEndpointbmAttributesSynchronizationTypeShift KUS = 0
	KUSBEndpointbmAttributesTransferTypeMask         KUS = 0
	KUSBEndpointbmAttributesUsageTypeMask            KUS = 0
	KUSBEndpointbmAttributesUsageTypeShift           KUS = 0
	KUSBFeedbackIsocUsageType                        KUS = 0
	KUSBHubDescriptorType                            KUS = 0
	KUSBImplicitFeedbackDataIsocUsageType            KUS = 0
	KUSBNoSynchronizationIsocSyncType                KUS = 0
	KUSBNotificationInterruptUsageType               KUS = 0
	KUSBPeriodicInterruptUsageType                   KUS = 0
	KUSBSuperSpeedLTMCapable                         KUS = 0
	KUSBSuperSpeedSupportsFS                         KUS = 0
	KUSBSuperSpeedSupportsHS                         KUS = 0
	KUSBSuperSpeedSupportsLS                         KUS = 0
	KUSBSuperSpeedSupportsSS                         KUS = 0
	KUSBSynchronousIsocSyncType                      KUS = 0
	KUSBbEndpointAddressMask                         KUS = 0
	KUSBbEndpointDirectionBit                        KUS = 0
	KUSBbEndpointDirectionMask                       KUS = 0
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
	// KUSBAddExtraResetTimeBit: Request extra time after reset.
	KUSBAddExtraResetTimeBit KUSB = 0
	// KUSBAddExtraResetTimeMask: The mask to request extra time after reset.
	KUSBAddExtraResetTimeMask                 KUSB = 0
	KUSBAllStreams                            KUSB = 0
	KUSBAnyType                               KUSB = 0
	KUSBBulk                                  KUSB = 0
	KUSBClass                                 KUSB = 0
	KUSBCommEthernetNetworkingSubClass        KUSB = 0
	KUSBCommunicationControlInterfaceClass    KUSB = 0
	KUSBCompositeSubClass                     KUSB = 0
	KUSBDevice                                KUSB = 0
	KUSBDeviceCountExceededNotificationType   KUSB = 0
	KUSBDeviceIDMask                          KUSB = 0
	KUSBDeviceIDShift                         KUSB = 0
	KUSBDeviceMask                            KUSB = 0
	KUSBEndPtShift                            KUSB = 0
	KUSBEndpoint                              KUSB = 0
	KUSBEndpointCountExceededNotificationType KUSB = 0
	KUSBEndpointTransferTypeUCMask            KUSB = 0
	// KUSBFullSpeedMicrosecondsInFrame: # Discussion
	KUSBFullSpeedMicrosecondsInFrame KUSB = 0
	KUSBFunctionRemoteWakeCapableBit KUSB = 0
	KUSBFunctionRemoteWakeEnableBit  KUSB = 0
	KUSBFunctionRemoteWakeupBit      KUSB = 0
	// KUSBHighSpeedMicrosecondsInFrame: # Discussion
	KUSBHighSpeedMicrosecondsInFrame KUSB = 0
	KUSBInterface                    KUSB = 0
	KUSBInterfaceIDMask              KUSB = 0
	KUSBInterfaceIDShift             KUSB = 0
	KUSBLowPowerSuspendStateBit      KUSB = 0
	KUSBMaxDevice                    KUSB = 0
	KUSBMaxDevices                   KUSB = 0
	KUSBMaxInterfaces                KUSB = 0
	KUSBMaxPipes                     KUSB = 0
	KUSBMaxStream                    KUSB = 0
	KUSBNoPipeIdx                    KUSB = 0
	KUSBNoStream                     KUSB = 0
	KUSBOther                        KUSB = 0
	KUSBPRimeStream                  KUSB = 0
	KUSBPhysicalInterfaceClass       KUSB = 0
	KUSBPipeIDMask                   KUSB = 0
	// KUSBReEnumerateCaptureDeviceBit: The bit to capture the device.
	KUSBReEnumerateCaptureDeviceBit KUSB = 0
	// KUSBReEnumerateCaptureDeviceMask: The mask to capture the device.
	KUSBReEnumerateCaptureDeviceMask KUSB = 0
	// KUSBReEnumerateReleaseDeviceBit: The bit to release the device.
	KUSBReEnumerateReleaseDeviceBit KUSB = 0
	// KUSBReEnumerateReleaseDeviceMask: The mask to release the device.
	KUSBReEnumerateReleaseDeviceMask          KUSB = 0
	KUSBRel10                                 KUSB = 0
	KUSBStandard                              KUSB = 0
	KUSBStream0                               KUSB = 0
	KUSBStreamIDAllStreamsMask                KUSB = 0
	KUSBStreamIDMask                          KUSB = 0
	KUSBTooManyDevicesAddress                 KUSB = 0
	KUSBUCRequestWithoutUSBNotificationMask   KUSB = 0
	KUSBVendor                                KUSB = 0
	KUSB_EPDesc_MaxMPS                        KUSB = 0
	KUSB_EPDesc_bmAttributes_SyncType_Mask    KUSB = 0
	KUSB_EPDesc_bmAttributes_SyncType_Shift   KUSB = 0
	KUSB_EPDesc_bmAttributes_TranType_Mask    KUSB = 0
	KUSB_EPDesc_bmAttributes_TranType_Shift   KUSB = 0
	KUSB_EPDesc_bmAttributes_UsageType_Mask   KUSB = 0
	KUSB_EPDesc_bmAttributes_UsageType_Shift  KUSB = 0
	KUSB_EPDesc_wMaxPacketSize_MPS_Mask       KUSB = 0
	KUSB_EPDesc_wMaxPacketSize_MPS_Shift      KUSB = 0
	KUSB_HSFSEPDesc_wMaxPacketSize_Mult_Mask  KUSB = 0
	KUSB_HSFSEPDesc_wMaxPacketSize_Mult_Shift KUSB = 0
)

func (e KUSB) String() string {
	switch e {
	case KUSBAddExtraResetTimeBit:
		return "KUSBAddExtraResetTimeBit"
	default:
		return fmt.Sprintf("KUSB(%d)", e)
	}
}

type KUSB100mA uint

const (
	KUSB100mAAvailable KUSB100mA = 0
)

func (e KUSB100mA) String() string {
	switch e {
	case KUSB100mAAvailable:
		return "KUSB100mAAvailable"
	default:
		return fmt.Sprintf("KUSB100mA(%d)", e)
	}
}

type KUSB2LPMMaxL1 uint

const (
	KUSB2LPMMaxL1Timeout KUSB2LPMMaxL1 = 0
)

func (e KUSB2LPMMaxL1) String() string {
	switch e {
	case KUSB2LPMMaxL1Timeout:
		return "KUSB2LPMMaxL1Timeout"
	default:
		return fmt.Sprintf("KUSB2LPMMaxL1(%d)", e)
	}
}

type KUSBAddress uint

const (
	KUSBAddress_Mask KUSBAddress = 0
)

func (e KUSBAddress) String() string {
	switch e {
	case KUSBAddress_Mask:
		return "KUSBAddress_Mask"
	default:
		return fmt.Sprintf("KUSBAddress(%d)", e)
	}
}

type KUSBAny uint

const (
	KUSBAnyDirn KUSBAny = 0
)

func (e KUSBAny) String() string {
	switch e {
	case KUSBAnyDirn:
		return "KUSBAnyDirn"
	default:
		return fmt.Sprintf("KUSBAny(%d)", e)
	}
}

type KUSBApplicationSpecific uint

const (
	KUSBApplicationSpecificClass KUSBApplicationSpecific = 0
)

func (e KUSBApplicationSpecific) String() string {
	switch e {
	case KUSBApplicationSpecificClass:
		return "KUSBApplicationSpecificClass"
	default:
		return fmt.Sprintf("KUSBApplicationSpecific(%d)", e)
	}
}

type KUSBBillboardAdditinalInfoNo uint

const (
	KUSBBillboardAdditinalInfoNoPower KUSBBillboardAdditinalInfoNo = 0
	KUSBBillboardAdditinalInfoNoUSBPD KUSBBillboardAdditinalInfoNo = 0
)

func (e KUSBBillboardAdditinalInfoNo) String() string {
	switch e {
	case KUSBBillboardAdditinalInfoNoPower:
		return "KUSBBillboardAdditinalInfoNoPower"
	default:
		return fmt.Sprintf("KUSBBillboardAdditinalInfoNo(%d)", e)
	}
}

type KUSBBillboardAltModeConfig uint

const (
	KUSBBillboardAltModeConfigSuccess KUSBBillboardAltModeConfig = 0
)

func (e KUSBBillboardAltModeConfig) String() string {
	switch e {
	case KUSBBillboardAltModeConfigSuccess:
		return "KUSBBillboardAltModeConfigSuccess"
	default:
		return fmt.Sprintf("KUSBBillboardAltModeConfig(%d)", e)
	}
}

type KUSBBillboardV uint

const (
	KUSBBillboardVConn1P5Watt KUSBBillboardV = 0
	KUSBBillboardVConn6Watt   KUSBBillboardV = 0
)

func (e KUSBBillboardV) String() string {
	switch e {
	case KUSBBillboardVConn1P5Watt:
		return "KUSBBillboardVConn1P5Watt"
	default:
		return fmt.Sprintf("KUSBBillboardV(%d)", e)
	}
}

type KUSBCapsLock uint

const (
	KUSBCapsLockKey KUSBCapsLock = 0
)

func (e KUSBCapsLock) String() string {
	switch e {
	case KUSBCapsLockKey:
		return "KUSBCapsLockKey"
	default:
		return fmt.Sprintf("KUSBCapsLock(%d)", e)
	}
}

type KUSBClassSpecific uint

const (
	// KUSBClassSpecificDescriptor: The class-specific descriptor value.
	KUSBClassSpecificDescriptor KUSBClassSpecific = 0
)

func (e KUSBClassSpecific) String() string {
	switch e {
	case KUSBClassSpecificDescriptor:
		return "KUSBClassSpecificDescriptor"
	default:
		return fmt.Sprintf("KUSBClassSpecific(%d)", e)
	}
}

type KUSBDFUAttributes uint

const (
	KUSBDFUAttributesMask KUSBDFUAttributes = 0
)

func (e KUSBDFUAttributes) String() string {
	switch e {
	case KUSBDFUAttributesMask:
		return "KUSBDFUAttributesMask"
	default:
		return fmt.Sprintf("KUSBDFUAttributes(%d)", e)
	}
}

type KUSBDefaultControlCompletionTimeoutM uint

const (
	KUSBDefaultControlCompletionTimeoutMS KUSBDefaultControlCompletionTimeoutM = 0
)

func (e KUSBDefaultControlCompletionTimeoutM) String() string {
	switch e {
	case KUSBDefaultControlCompletionTimeoutMS:
		return "KUSBDefaultControlCompletionTimeoutMS"
	default:
		return fmt.Sprintf("KUSBDefaultControlCompletionTimeoutM(%d)", e)
	}
}

type KUSBDeviceCapability uint

const (
	KUSBDeviceCapabilityContainerID   KUSBDeviceCapability = 0
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
	KUSBDeviceLPMStatusDefault  KUSBDeviceLPMStatus = 0
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
	// KUSBDeviceSpeedFull: # Discussion
	KUSBDeviceSpeedFull KUSBDeviceSpeed = 0
	// KUSBDeviceSpeedHigh: # Discussion
	KUSBDeviceSpeedHigh KUSBDeviceSpeed = 0
	// KUSBDeviceSpeedLow: # Discussion
	KUSBDeviceSpeedLow KUSBDeviceSpeed = 0
	// KUSBDeviceSpeedSuper: # Discussion
	KUSBDeviceSpeedSuper        KUSBDeviceSpeed = 0
	KUSBDeviceSpeedSuperPlus    KUSBDeviceSpeed = 0
	KUSBDeviceSpeedSuperPlusBy2 KUSBDeviceSpeed = 0
)

func (e KUSBDeviceSpeed) String() string {
	switch e {
	case KUSBDeviceSpeedFull:
		return "KUSBDeviceSpeedFull"
	default:
		return fmt.Sprintf("KUSBDeviceSpeed(%d)", e)
	}
}

type KUSBDisplay uint

const (
	KUSBDisplayClass KUSBDisplay = 0
)

func (e KUSBDisplay) String() string {
	switch e {
	case KUSBDisplayClass:
		return "KUSBDisplayClass"
	default:
		return fmt.Sprintf("KUSBDisplay(%d)", e)
	}
}

type KUSBEndpointProperties uint

const (
	// KUSBEndpointPropertiesVersion3: # Discussion
	KUSBEndpointPropertiesVersion3 KUSBEndpointProperties = 0
)

func (e KUSBEndpointProperties) String() string {
	switch e {
	case KUSBEndpointPropertiesVersion3:
		return "KUSBEndpointPropertiesVersion3"
	default:
		return fmt.Sprintf("KUSBEndpointProperties(%d)", e)
	}
}

type KUSBFeature uint

const (
	KUSBFeatureEndpointStall KUSBFeature = 0
	KUSBFeatureTestMode      KUSBFeature = 0
)

func (e KUSBFeature) String() string {
	switch e {
	case KUSBFeatureEndpointStall:
		return "KUSBFeatureEndpointStall"
	default:
		return fmt.Sprintf("KUSBFeature(%d)", e)
	}
}

type KUSBHSHub uint

const (
	KUSBHSHubCommandAddHub         KUSBHSHub = 0
	KUSBHSHubCommandRemoveHub      KUSBHSHub = 0
	KUSBHSHubFlagsMoreInfoMask     KUSBHSHub = 0
	KUSBHSHubFlagsMultiTTMask      KUSBHSHub = 0
	KUSBHSHubFlagsNumPortsMask     KUSBHSHub = 0
	KUSBHSHubFlagsNumPortsShift    KUSBHSHub = 0
	KUSBHSHubFlagsTTThinkTimeMask  KUSBHSHub = 0
	KUSBHSHubFlagsTTThinkTimeShift KUSBHSHub = 0
)

func (e KUSBHSHub) String() string {
	switch e {
	case KUSBHSHubCommandAddHub:
		return "KUSBHSHubCommandAddHub"
	default:
		return fmt.Sprintf("KUSBHSHub(%d)", e)
	}
}

type KUSBHost uint

const (
	KUSBHostClassRequestCompletionTimeout     KUSBHost = 0
	KUSBHostDefaultControlCompletionTimeoutMS KUSBHost = 0
	KUSBHostHubClass                          KUSBHost = 0
	KUSBHostMaxCountFullSpeedIsochronous      KUSBHost = 0
	KUSBHostMaxDevices                        KUSBHost = 0
	KUSBHostMaxPipes                          KUSBHost = 0
	KUSBHostVendorIDAppleComputer             KUSBHost = 0
	KUSBHostVendorSpecificClass               KUSBHost = 0
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
	KUSB3TypeMicroABConnector        KUSBHostConnectorType = 0
	KUSBHostConnectorTypeA           KUSBHostConnectorType = 0
	KUSBHostConnectorTypeExpressCard KUSBHostConnectorType = 0
	KUSBHostConnectorTypeMiniAB      KUSBHostConnectorType = 0
	KUSBHostConnectorTypeProprietary KUSBHostConnectorType = 0
	KUSBHostConnectorTypeUSB3A       KUSBHostConnectorType = 0
	KUSBHostConnectorTypeUSB3B       KUSBHostConnectorType = 0
	KUSBHostConnectorTypeUSB3MicroAB KUSBHostConnectorType = 0
	KUSBHostConnectorTypeUSB3MicroB  KUSBHostConnectorType = 0
	KUSBHostConnectorTypeUSB3PowerB  KUSBHostConnectorType = 0
	KUSBHostConnectorTypeUSBTypeC    KUSBHostConnectorType = 0
	KUSBHostConnectorTypeUnknown     KUSBHostConnectorType = 0
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
	KUSBHostOpenOptionUserClientSession      KUSBHostOpenOption = 0
)

func (e KUSBHostOpenOption) String() string {
	switch e {
	case KUSBHostOpenOptionSelectAlternateSetting:
		return "KUSBHostOpenOptionSelectAlternateSetting"
	default:
		return fmt.Sprintf("KUSBHostOpenOption(%d)", e)
	}
}

type KUSBHostPort uint

const (
	KUSBHostPortConnectable KUSBHostPort = 0
)

func (e KUSBHostPort) String() string {
	switch e {
	case KUSBHostPortConnectable:
		return "KUSBHostPortConnectable"
	default:
		return fmt.Sprintf("KUSBHostPort(%d)", e)
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

type KUSBHostPowerSourceType uint

const (
	KUSBHostPowerSourceTypeHardware KUSBHostPowerSourceType = 0
)

func (e KUSBHostPowerSourceType) String() string {
	switch e {
	case KUSBHostPowerSourceTypeHardware:
		return "KUSBHostPowerSourceTypeHardware"
	default:
		return fmt.Sprintf("KUSBHostPowerSourceType(%d)", e)
	}
}

type KUSBHub uint

const (
	KUSBHubLocalPowerChangeFeature KUSBHub = 0
	KUSBHubPortConnectionFeature   KUSBHub = 0
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
	KUSBHubRqGetInterface      KUSBHubRq = 0
	KUSBHubRqGetPortErrorCount KUSBHubRq = 0
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
	KUSBHubRqSetFeature  KUSBHubRq = 0
	KUSBHubRqSetHubDepth KUSBHubRq = 0
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

type KUSBLowLatencyIsochTransfer uint

const (
	KUSBLowLatencyIsochTransferKey KUSBLowLatencyIsochTransfer = 0
)

func (e KUSBLowLatencyIsochTransfer) String() string {
	switch e {
	case KUSBLowLatencyIsochTransferKey:
		return "KUSBLowLatencyIsochTransferKey"
	default:
		return fmt.Sprintf("KUSBLowLatencyIsochTransfer(%d)", e)
	}
}

type KUSBMaxFSIsocEndpointReq uint

const (
	KUSBMaxFSIsocEndpointReqCount KUSBMaxFSIsocEndpointReq = 0
)

func (e KUSBMaxFSIsocEndpointReq) String() string {
	switch e {
	case KUSBMaxFSIsocEndpointReqCount:
		return "KUSBMaxFSIsocEndpointReqCount"
	default:
		return fmt.Sprintf("KUSBMaxFSIsocEndpointReq(%d)", e)
	}
}

type KUSBNotification uint

const (
	// KUSBNotificationPostForcedResumeValue: The system sends a notification after a resume, which occurs after a forced suspend.
	KUSBNotificationPostForcedResumeValue KUSBNotification = 0
	// KUSBNotificationPostForcedSuspend: The system sends a notification after a forced suspend completes.
	KUSBNotificationPostForcedSuspend KUSBNotification = 0
	// KUSBNotificationPreForcedResume: The system sends a notification before a resume, which occurs after a forced suspend.
	KUSBNotificationPreForcedResume KUSBNotification = 0
	// KUSBNotificationPreForcedSuspend: The system sends a notification prior to a forced suspend.
	KUSBNotificationPreForcedSuspend KUSBNotification = 0
)

func (e KUSBNotification) String() string {
	switch e {
	case KUSBNotificationPostForcedResumeValue:
		return "KUSBNotificationPostForcedResumeValue"
	default:
		return fmt.Sprintf("KUSBNotification(%d)", e)
	}
}

type KUSBNotificationPostForcedResume uint

const (
	KUSBNotificationPostForcedResumeBit KUSBNotificationPostForcedResume = 0
)

func (e KUSBNotificationPostForcedResume) String() string {
	switch e {
	case KUSBNotificationPostForcedResumeBit:
		return "KUSBNotificationPostForcedResumeBit"
	default:
		return fmt.Sprintf("KUSBNotificationPostForcedResume(%d)", e)
	}
}

type KUSBPort uint

const (
	KUSBPortConnectable    KUSBPort = 0
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

type KUSBRq uint

const (
	KUSBRqDirnMask      KUSBRq = 0
	KUSBRqDirnShift     KUSBRq = 0
	KUSBRqRecipientMask KUSBRq = 0
	KUSBRqTypeMask      KUSBRq = 0
	KUSBRqTypeShift     KUSBRq = 0
)

func (e KUSBRq) String() string {
	switch e {
	case KUSBRqDirnMask:
		return "KUSBRqDirnMask"
	default:
		return fmt.Sprintf("KUSBRq(%d)", e)
	}
}

type KUSBRqClear uint

const (
	KUSBRqClearFeature KUSBRqClear = 0
)

func (e KUSBRqClear) String() string {
	switch e {
	case KUSBRqClearFeature:
		return "KUSBRqClearFeature"
	default:
		return fmt.Sprintf("KUSBRqClear(%d)", e)
	}
}

type KUSBTypeCCableType uint

const (
	KUSBTypeCCableTypeNone KUSBTypeCCableType = 0
)

func (e KUSBTypeCCableType) String() string {
	switch e {
	case KUSBTypeCCableTypeNone:
		return "KUSBTypeCCableTypeNone"
	default:
		return fmt.Sprintf("KUSBTypeCCableType(%d)", e)
	}
}

type KUSPrintingClassGePort uint

const (
	KUSPrintingClassGePortStatus KUSPrintingClassGePort = 0
)

func (e KUSPrintingClassGePort) String() string {
	switch e {
	case KUSPrintingClassGePortStatus:
		return "KUSPrintingClassGePortStatus"
	default:
		return fmt.Sprintf("KUSPrintingClassGePort(%d)", e)
	}
}

type KUnknownConnect uint

const (
	KFixedModeCRTConnect KUnknownConnect = 0
	KPanelFSTNConnect    KUnknownConnect = 0
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
	KFrameInterruptServiceType    KVBLInterruptServiceType = 0
)

func (e KVBLInterruptServiceType) String() string {
	switch e {
	case KFBOnlineInterruptServiceType:
		return "KFBOnlineInterruptServiceType"
	default:
		return fmt.Sprintf("KVBLInterruptServiceType(%d)", e)
	}
}

type KVCAcquire uint

const (
	KVCAcquireImmediate KVCAcquire = 0
)

func (e KVCAcquire) String() string {
	switch e {
	case KVCAcquireImmediate:
		return "KVCAcquireImmediate"
	default:
		return fmt.Sprintf("KVCAcquire(%d)", e)
	}
}

type KVSLClamshellStateGestalt uint

const (
	KVSLClamshellStateGestaltType KVSLClamshellStateGestalt = 0
)

func (e KVSLClamshellStateGestalt) String() string {
	switch e {
	case KVSLClamshellStateGestaltType:
		return "KVSLClamshellStateGestaltType"
	default:
		return fmt.Sprintf("KVSLClamshellStateGestalt(%d)", e)
	}
}

type KVariableLength uint

const (
	KVariableLengthArray KVariableLength = 0
)

func (e KVariableLength) String() string {
	switch e {
	case KVariableLengthArray:
		return "KVariableLengthArray"
	default:
		return fmt.Sprintf("KVariableLength(%d)", e)
	}
}

type KVideo uint

const (
	KVideoCombinedI2CTypeMask KVideo = 0
	KVideoNoTransactionType   KVideo = 0
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
	KVideoBusTypeInvalid     KVideoBusType = 0
)

func (e KVideoBusType) String() string {
	switch e {
	case KVideoBusTypeDisplayPort:
		return "KVideoBusTypeDisplayPort"
	default:
		return fmt.Sprintf("KVideoBusType(%d)", e)
	}
}

type KVideoDefault uint

const (
	KVideoDefaultBus KVideoDefault = 0
)

func (e KVideoDefault) String() string {
	switch e {
	case KVideoDefaultBus:
		return "KVideoDefaultBus"
	default:
		return fmt.Sprintf("KVideoDefault(%d)", e)
	}
}

type KVideoUsageAddrSubAddr uint

const (
	KVideoUsageAddrSubAddrBit  KVideoUsageAddrSubAddr = 0
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
	KSuperSpeedBusBitMask      KXHCISSRootHubAddress = 0
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

type Kanalogsignallevel0700 uint

const (
	KAnalogSignalLevel_0700_0000 Kanalogsignallevel0700 = 0
)

func (e Kanalogsignallevel0700) String() string {
	switch e {
	case KAnalogSignalLevel_0700_0000:
		return "KAnalogSignalLevel_0700_0000"
	default:
		return fmt.Sprintf("Kanalogsignallevel0700(%d)", e)
	}
}

type KcSt uint

const (
	KC_ST_CHAR KcSt = 0
)

func (e KcSt) String() string {
	switch e {
	case KC_ST_CHAR:
		return "KC_ST_CHAR"
	default:
		return fmt.Sprintf("KcSt(%d)", e)
	}
}

type Kcdct uint

const (
	KCDCT_NONE Kcdct = 0
)

func (e Kcdct) String() string {
	switch e {
	case KCDCT_NONE:
		return "KCDCT_NONE"
	default:
		return fmt.Sprintf("Kcdct(%d)", e)
	}
}

type KdCallback uint

const (
	KD_CALLBACK_KDEBUG_DISABLED KdCallback = 0
	KD_CALLBACK_KDEBUG_ENABLED  KdCallback = 0
	KD_CALLBACK_SYNC_FLUSH      KdCallback = 0
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
	KDBG_BUFINIT       Kdbg = 0
	KDBG_MATCH_DISABLE Kdbg = 0
	KDBG_NOWRAP        Kdbg = 0
	KDBG_WRAPPED       Kdbg = 0
)

func (e Kdbg) String() string {
	switch e {
	case KDBG_BUFINIT:
		return "KDBG_BUFINIT"
	default:
		return fmt.Sprintf("Kdbg(%d)", e)
	}
}

type KdcpContinuous uint

const (
	KDCP_CONTINUOUS_TIME KdcpContinuous = 0
)

func (e KdcpContinuous) String() string {
	switch e {
	case KDCP_CONTINUOUS_TIME:
		return "KDCP_CONTINUOUS_TIME"
	default:
		return fmt.Sprintf("KdcpContinuous(%d)", e)
	}
}

type KdpEvent uint

const (
	KDP_EVENT_ENTER    KdpEvent = 0
	KDP_EVENT_EXIT     KdpEvent = 0
	KDP_EVENT_PANICLOG KdpEvent = 0
)

func (e KdpEvent) String() string {
	switch e {
	case KDP_EVENT_ENTER:
		return "KDP_EVENT_ENTER"
	default:
		return fmt.Sprintf("KdpEvent(%d)", e)
	}
}

type Kdtest uint

const (
	KDTEST_ABSOLUTE_TIMESTAMP   Kdtest = 0
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
	KNMIIntLevelMask      Kextint9Nmiintsource = 0
	KNMIIntMask           Kextint9Nmiintsource = 0
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
	KF_COMPRSV_OVRD                          Kf = 0
	KF_DISABLE_FP_POPC_ON_PGFLT              Kf = 0
	KF_DISABLE_PROCREF_TRACKING_OVRD         Kf = 0
	KF_DISABLE_PROD_TRC_VALIDATION           Kf = 0
	KF_INTERRUPT_MASKED_DEBUG_OVRD           Kf = 0
	KF_INTERRUPT_MASKED_DEBUG_STACKSHOT_OVRD Kf = 0
	KF_IOTRACE_OVRD                          Kf = 0
	KF_IO_TIMEOUT_OVRD                       Kf = 0
	KF_MACH_ASSERT_OVRD                      Kf = 0
	KF_MADVISE_FREE_DEBUG_OVRD               Kf = 0
	KF_MATV_OVRD                             Kf = 0
	KF_SERIAL_OVRD                           Kf = 0
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
	KGUARD_EXC_DEALLOC_GAP                      KguardExc = 0
	KGUARD_EXC_DESTROY                          KguardExc = 0
	KGUARD_EXC_EXCEPTION_BEHAVIOR_ENFORCE       KguardExc = 0
	KGUARD_EXC_IMMOVABLE                        KguardExc = 0
	KGUARD_EXC_IMMOVABLE_NON_FATAL              KguardExc = 0
	KGUARD_EXC_INCORRECT_GUARD                  KguardExc = 0
	KGUARD_EXC_INVALID_ARGUMENT                 KguardExc = 0
	KGUARD_EXC_INVALID_NAME                     KguardExc = 0
	KGUARD_EXC_INVALID_OPTIONS                  KguardExc = 0
	KGUARD_EXC_INVALID_RIGHT                    KguardExc = 0
	KGUARD_EXC_INVALID_VALUE                    KguardExc = 0
	KGUARD_EXC_KERN_FAILURE                     KguardExc = 0
	KGUARD_EXC_KERN_NO_SPACE                    KguardExc = 0
	KGUARD_EXC_KERN_RESOURCE                    KguardExc = 0
	KGUARD_EXC_MOD_REFS                         KguardExc = 0
	KGUARD_EXC_MOD_REFS_NON_FATAL               KguardExc = 0
	KGUARD_EXC_MSG_FILTERED                     KguardExc = 0
	KGUARD_EXC_PROVISIONAL_REPLY_PORT           KguardExc = 0
	KGUARD_EXC_RCV_GUARDED_DESC                 KguardExc = 0
	KGUARD_EXC_RCV_INVALID_NAME                 KguardExc = 0
	KGUARD_EXC_RECLAIM_COPYIO_FAILURE           KguardExc = 0
	KGUARD_EXC_REQUIRE_REPLY_PORT_SEMANTICS     KguardExc = 0
	KGUARD_EXC_RIGHT_EXISTS                     KguardExc = 0
	KGUARD_EXC_SEND_INVALID_REPLY               KguardExc = 0
	KGUARD_EXC_SEND_INVALID_RIGHT               KguardExc = 0
	KGUARD_EXC_SEND_INVALID_VOUCHER             KguardExc = 0
	KGUARD_EXC_SERVICE_PORT_VIOLATION_FATAL     KguardExc = 0
	KGUARD_EXC_SERVICE_PORT_VIOLATION_NON_FATAL KguardExc = 0
	KGUARD_EXC_SET_CONTEXT                      KguardExc = 0
	KGUARD_EXC_STRICT_REPLY                     KguardExc = 0
	KGUARD_EXC_THREAD_SET_STATE                 KguardExc = 0
	KGUARD_EXC_UNGUARDED                        KguardExc = 0
)

func (e KguardExc) String() string {
	switch e {
	case KGUARD_EXC_DEALLOC_GAP:
		return "KGUARD_EXC_DEALLOC_GAP"
	default:
		return fmt.Sprintf("KguardExc(%d)", e)
	}
}

type KhidusageAd uint

const (
	KHIDUsage_AD_ASCIICharacterSet          KhidusageAd = 0
	KHIDUsage_AD_AlphanumericDisplay        KhidusageAd = 0
	KHIDUsage_AD_CharacterHeight            KhidusageAd = 0
	KHIDUsage_AD_CharacterReport            KhidusageAd = 0
	KHIDUsage_AD_CharacterSpacingHorizontal KhidusageAd = 0
	KHIDUsage_AD_CharacterSpacingVertical   KhidusageAd = 0
	KHIDUsage_AD_CharacterWidth             KhidusageAd = 0
	KHIDUsage_AD_ClearDisplay               KhidusageAd = 0
	KHIDUsage_AD_Column                     KhidusageAd = 0
	KHIDUsage_AD_Columns                    KhidusageAd = 0
	KHIDUsage_AD_CursorBlink                KhidusageAd = 0
	KHIDUsage_AD_CursorEnable               KhidusageAd = 0
	KHIDUsage_AD_CursorMode                 KhidusageAd = 0
	KHIDUsage_AD_CursorPixelPositioning     KhidusageAd = 0
	KHIDUsage_AD_CursorPositionReport       KhidusageAd = 0
	KHIDUsage_AD_DataReadBack               KhidusageAd = 0
	KHIDUsage_AD_DisplayAttributesReport    KhidusageAd = 0
	KHIDUsage_AD_DisplayControlReport       KhidusageAd = 0
	KHIDUsage_AD_DisplayData                KhidusageAd = 0
	KHIDUsage_AD_DisplayEnable              KhidusageAd = 0
	KHIDUsage_AD_DisplayStatus              KhidusageAd = 0
	KHIDUsage_AD_ErrFontdatacannotberead    KhidusageAd = 0
	KHIDUsage_AD_ErrNotaloadablecharacter   KhidusageAd = 0
	KHIDUsage_AD_FontData                   KhidusageAd = 0
	KHIDUsage_AD_FontReadBack               KhidusageAd = 0
	KHIDUsage_AD_FontReport                 KhidusageAd = 0
	KHIDUsage_AD_HorizontalScroll           KhidusageAd = 0
	KHIDUsage_AD_Reserved                   KhidusageAd = 0
	KHIDUsage_AD_Row                        KhidusageAd = 0
	KHIDUsage_AD_Rows                       KhidusageAd = 0
	KHIDUsage_AD_ScreenSaverDelay           KhidusageAd = 0
	KHIDUsage_AD_ScreenSaverEnable          KhidusageAd = 0
	KHIDUsage_AD_StatNotReady               KhidusageAd = 0
	KHIDUsage_AD_StatReady                  KhidusageAd = 0
	KHIDUsage_AD_UnicodeCharacterSet        KhidusageAd = 0
	KHIDUsage_AD_VerticalScroll             KhidusageAd = 0
)

func (e KhidusageAd) String() string {
	switch e {
	case KHIDUsage_AD_ASCIICharacterSet:
		return "KHIDUsage_AD_ASCIICharacterSet"
	default:
		return fmt.Sprintf("KhidusageAd(%d)", e)
	}
}

type KhidusageBcs uint

const (
	KHIDUsage_BCS_2DControlReport                            KhidusageBcs = 0
	KHIDUsage_BCS_ActiveTime                                 KhidusageBcs = 0
	KHIDUsage_BCS_AddEAN2_3LabelDefinition                   KhidusageBcs = 0
	KHIDUsage_BCS_AimDuration                                KhidusageBcs = 0
	KHIDUsage_BCS_AimingLaserPattern                         KhidusageBcs = 0
	KHIDUsage_BCS_Aiming_PointerMide                         KhidusageBcs = 0
	KHIDUsage_BCS_AttributeReport                            KhidusageBcs = 0
	KHIDUsage_BCS_AztecCode                                  KhidusageBcs = 0
	KHIDUsage_BCS_BC412                                      KhidusageBcs = 0
	KHIDUsage_BCS_BadgeReader                                KhidusageBcs = 0
	KHIDUsage_BCS_BarCodePresent                             KhidusageBcs = 0
	KHIDUsage_BCS_BarCodePresentSensor                       KhidusageBcs = 0
	KHIDUsage_BCS_BarCodeScanner                             KhidusageBcs = 0
	KHIDUsage_BCS_BarCodeScannerCradle                       KhidusageBcs = 0
	KHIDUsage_BCS_BarSpaceData                               KhidusageBcs = 0
	KHIDUsage_BCS_BeeperState                                KhidusageBcs = 0
	KHIDUsage_BCS_BooklandEAN                                KhidusageBcs = 0
	KHIDUsage_BCS_ChannelCode                                KhidusageBcs = 0
	KHIDUsage_BCS_Check                                      KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigit                                 KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigitCodabarEnable                    KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigitCode99Enable                     KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigitDisable                          KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigitEnableInterleaved2of5OPCC        KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigitEnableInterleaved2of5USS         KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigitEnableOneMSIPlessey              KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigitEnableStandard2of5OPCC           KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigitEnableStandard2of5USS            KhidusageBcs = 0
	KHIDUsage_BCS_CheckDigitEnableTwoMSIPlessey              KhidusageBcs = 0
	KHIDUsage_BCS_CheckDisablePrice                          KhidusageBcs = 0
	KHIDUsage_BCS_CheckEnable4DigitPrice                     KhidusageBcs = 0
	KHIDUsage_BCS_CheckEnable5DigitPrice                     KhidusageBcs = 0
	KHIDUsage_BCS_CheckEnableEuropean4DigitPrice             KhidusageBcs = 0
	KHIDUsage_BCS_CheckEnableEuropean5DigitPrice             KhidusageBcs = 0
	KHIDUsage_BCS_Class1ALaser                               KhidusageBcs = 0
	KHIDUsage_BCS_Class2Laser                                KhidusageBcs = 0
	KHIDUsage_BCS_ClearAllEAN2_3LabelDefinitions             KhidusageBcs = 0
	KHIDUsage_BCS_Codabar                                    KhidusageBcs = 0
	KHIDUsage_BCS_CodabarControlReport                       KhidusageBcs = 0
	KHIDUsage_BCS_Code128                                    KhidusageBcs = 0
	KHIDUsage_BCS_Code128ControlReport                       KhidusageBcs = 0
	KHIDUsage_BCS_Code16                                     KhidusageBcs = 0
	KHIDUsage_BCS_Code32                                     KhidusageBcs = 0
	KHIDUsage_BCS_Code39                                     KhidusageBcs = 0
	KHIDUsage_BCS_Code39ControlReport                        KhidusageBcs = 0
	KHIDUsage_BCS_Code49                                     KhidusageBcs = 0
	KHIDUsage_BCS_Code93                                     KhidusageBcs = 0
	KHIDUsage_BCS_CodeOne                                    KhidusageBcs = 0
	KHIDUsage_BCS_Colorcode                                  KhidusageBcs = 0
	KHIDUsage_BCS_CommitParametersToNVM                      KhidusageBcs = 0
	KHIDUsage_BCS_ConstantElectronicArticleSurveillance      KhidusageBcs = 0
	KHIDUsage_BCS_ContactScanner                             KhidusageBcs = 0
	KHIDUsage_BCS_ConvertEAN8To13Type                        KhidusageBcs = 0
	KHIDUsage_BCS_ConvertUPCAToEAN_13                        KhidusageBcs = 0
	KHIDUsage_BCS_ConvertUPC_EToA                            KhidusageBcs = 0
	KHIDUsage_BCS_CordlessScannerBase                        KhidusageBcs = 0
	KHIDUsage_BCS_DLMethodCheckForDiscrete                   KhidusageBcs = 0
	KHIDUsage_BCS_DLMethodCheckInRange                       KhidusageBcs = 0
	KHIDUsage_BCS_DLMethodReadAny                            KhidusageBcs = 0
	KHIDUsage_BCS_DataLengthMethod                           KhidusageBcs = 0
	KHIDUsage_BCS_DataMatrix                                 KhidusageBcs = 0
	KHIDUsage_BCS_DataPrefix                                 KhidusageBcs = 0
	KHIDUsage_BCS_DecodeDataContinued                        KhidusageBcs = 0
	KHIDUsage_BCS_DecodedData                                KhidusageBcs = 0
	KHIDUsage_BCS_DisableCheckDigitTransmit                  KhidusageBcs = 0
	KHIDUsage_BCS_DumbBarCodeScanner                         KhidusageBcs = 0
	KHIDUsage_BCS_EAN13FlagDigit1                            KhidusageBcs = 0
	KHIDUsage_BCS_EAN13FlagDigit2                            KhidusageBcs = 0
	KHIDUsage_BCS_EAN13FlagDigit3                            KhidusageBcs = 0
	KHIDUsage_BCS_EAN2_3LabelControlReport                   KhidusageBcs = 0
	KHIDUsage_BCS_EAN8FlagDigit1                             KhidusageBcs = 0
	KHIDUsage_BCS_EAN8FlagDigit2                             KhidusageBcs = 0
	KHIDUsage_BCS_EAN8FlagDigit3                             KhidusageBcs = 0
	KHIDUsage_BCS_EANThreeLabel                              KhidusageBcs = 0
	KHIDUsage_BCS_EANTwoLabel                                KhidusageBcs = 0
	KHIDUsage_BCS_EAN_13                                     KhidusageBcs = 0
	KHIDUsage_BCS_EAN_8                                      KhidusageBcs = 0
	KHIDUsage_BCS_EAN_99_128_Mandatory                       KhidusageBcs = 0
	KHIDUsage_BCS_EAN_99_P5_128_Optional                     KhidusageBcs = 0
	KHIDUsage_BCS_ElectronicArticleSurveillanceNotification  KhidusageBcs = 0
	KHIDUsage_BCS_EnableCheckDigitTransmit                   KhidusageBcs = 0
	KHIDUsage_BCS_ErrorIndication                            KhidusageBcs = 0
	KHIDUsage_BCS_FirstDiscreteLengthToDecode                KhidusageBcs = 0
	KHIDUsage_BCS_FixedBeeper                                KhidusageBcs = 0
	KHIDUsage_BCS_FragmentDecoding                           KhidusageBcs = 0
	KHIDUsage_BCS_FullASCIIConversion                        KhidusageBcs = 0
	KHIDUsage_BCS_GRWTIAfterDecode                           KhidusageBcs = 0
	KHIDUsage_BCS_GRWTIBeep_LampAfterTransmit                KhidusageBcs = 0
	KHIDUsage_BCS_GRWTINoBeep_LampUseAtAll                   KhidusageBcs = 0
	KHIDUsage_BCS_GoodDecodeIndication                       KhidusageBcs = 0
	KHIDUsage_BCS_GoodReadLED                                KhidusageBcs = 0
	KHIDUsage_BCS_GoodReadLampDuration                       KhidusageBcs = 0
	KHIDUsage_BCS_GoodReadLampIntensity                      KhidusageBcs = 0
	KHIDUsage_BCS_GoodReadToneFrequency                      KhidusageBcs = 0
	KHIDUsage_BCS_GoodReadToneLength                         KhidusageBcs = 0
	KHIDUsage_BCS_GoodReadToneVolume                         KhidusageBcs = 0
	KHIDUsage_BCS_GoodReadWhenToWrite                        KhidusageBcs = 0
	KHIDUsage_BCS_HandsFreeScanning                          KhidusageBcs = 0
	KHIDUsage_BCS_HeaterPresent                              KhidusageBcs = 0
	KHIDUsage_BCS_InitiateBarcodeRead                        KhidusageBcs = 0
	KHIDUsage_BCS_Interleaved2of5                            KhidusageBcs = 0
	KHIDUsage_BCS_Interleaved2of5ControlReport               KhidusageBcs = 0
	KHIDUsage_BCS_IntrinsicallySafe                          KhidusageBcs = 0
	KHIDUsage_BCS_ItalianPharmacyCode                        KhidusageBcs = 0
	KHIDUsage_BCS_KlasseEinsLaser                            KhidusageBcs = 0
	KHIDUsage_BCS_LaserOnTime                                KhidusageBcs = 0
	KHIDUsage_BCS_LaserState                                 KhidusageBcs = 0
	KHIDUsage_BCS_LockoutTime                                KhidusageBcs = 0
	KHIDUsage_BCS_LongRangeScanner                           KhidusageBcs = 0
	KHIDUsage_BCS_MSIPlesseyControlReport                    KhidusageBcs = 0
	KHIDUsage_BCS_MSI_Plessey                                KhidusageBcs = 0
	KHIDUsage_BCS_MaxiCode                                   KhidusageBcs = 0
	KHIDUsage_BCS_MaximumLengthToDecode                      KhidusageBcs = 0
	KHIDUsage_BCS_MicroPDF                                   KhidusageBcs = 0
	KHIDUsage_BCS_MinimumLengthToDecode                      KhidusageBcs = 0
	KHIDUsage_BCS_MirrorSpeedControl                         KhidusageBcs = 0
	KHIDUsage_BCS_Misc1DControlReport                        KhidusageBcs = 0
	KHIDUsage_BCS_MotorState                                 KhidusageBcs = 0
	KHIDUsage_BCS_MotorTimeout                               KhidusageBcs = 0
	KHIDUsage_BCS_MultiRangeScanner                          KhidusageBcs = 0
	KHIDUsage_BCS_NoReadMessage                              KhidusageBcs = 0
	KHIDUsage_BCS_NotOnFileIndication                        KhidusageBcs = 0
	KHIDUsage_BCS_NotOnFileVolume                            KhidusageBcs = 0
	KHIDUsage_BCS_PDF_417                                    KhidusageBcs = 0
	KHIDUsage_BCS_ParameterScanning                          KhidusageBcs = 0
	KHIDUsage_BCS_ParametersChanged                          KhidusageBcs = 0
	KHIDUsage_BCS_Periodical                                 KhidusageBcs = 0
	KHIDUsage_BCS_PeriodicalAutoDiscriminatePlus2            KhidusageBcs = 0
	KHIDUsage_BCS_PeriodicalAutoDiscriminatePlus5            KhidusageBcs = 0
	KHIDUsage_BCS_PeriodicalIgnorePlus2                      KhidusageBcs = 0
	KHIDUsage_BCS_PeriodicalIgnorePlus5                      KhidusageBcs = 0
	KHIDUsage_BCS_PeriodicalOnlyDecodeWithPlus2              KhidusageBcs = 0
	KHIDUsage_BCS_PeriodicalOnlyDecodeWithPlus5              KhidusageBcs = 0
	KHIDUsage_BCS_PolarityInvertedBarCode                    KhidusageBcs = 0
	KHIDUsage_BCS_PolarityNormalBarCode                      KhidusageBcs = 0
	KHIDUsage_BCS_PosiCode                                   KhidusageBcs = 0
	KHIDUsage_BCS_PowerOnResetScanner                        KhidusageBcs = 0
	KHIDUsage_BCS_PowerupBeep                                KhidusageBcs = 0
	KHIDUsage_BCS_PrefixAIMI                                 KhidusageBcs = 0
	KHIDUsage_BCS_PrefixNone                                 KhidusageBcs = 0
	KHIDUsage_BCS_PrefixProprietary                          KhidusageBcs = 0
	KHIDUsage_BCS_PreventReadOfBarcodes                      KhidusageBcs = 0
	KHIDUsage_BCS_ProgrammableBeeper                         KhidusageBcs = 0
	KHIDUsage_BCS_ProximitySensor                            KhidusageBcs = 0
	KHIDUsage_BCS_QRCode                                     KhidusageBcs = 0
	KHIDUsage_BCS_RawDataPolarity                            KhidusageBcs = 0
	KHIDUsage_BCS_RawScannedDataReport                       KhidusageBcs = 0
	KHIDUsage_BCS_ScannedDataReport                          KhidusageBcs = 0
	KHIDUsage_BCS_ScannerDataAccuracy                        KhidusageBcs = 0
	KHIDUsage_BCS_ScannerInCradle                            KhidusageBcs = 0
	KHIDUsage_BCS_ScannerInRange                             KhidusageBcs = 0
	KHIDUsage_BCS_ScannerReadConfidence                      KhidusageBcs = 0
	KHIDUsage_BCS_SecondDiscreteLengthToDecode               KhidusageBcs = 0
	KHIDUsage_BCS_SetParameterDefaultValues                  KhidusageBcs = 0
	KHIDUsage_BCS_SettingsReport                             KhidusageBcs = 0
	KHIDUsage_BCS_SoundErrorBeep                             KhidusageBcs = 0
	KHIDUsage_BCS_SoundGoodReadBeep                          KhidusageBcs = 0
	KHIDUsage_BCS_SoundNotOnFileBeep                         KhidusageBcs = 0
	KHIDUsage_BCS_Standard2of5                               KhidusageBcs = 0
	KHIDUsage_BCS_Standard2of5ControlReport                  KhidusageBcs = 0
	KHIDUsage_BCS_Standard2of5IATA                           KhidusageBcs = 0
	KHIDUsage_BCS_StatusReport                               KhidusageBcs = 0
	KHIDUsage_BCS_SuperCode                                  KhidusageBcs = 0
	KHIDUsage_BCS_SymbologyIdentifier1                       KhidusageBcs = 0
	KHIDUsage_BCS_SymbologyIdentifier2                       KhidusageBcs = 0
	KHIDUsage_BCS_SymbologyIdentifier3                       KhidusageBcs = 0
	KHIDUsage_BCS_TransmitCheckDigit                         KhidusageBcs = 0
	KHIDUsage_BCS_TransmitStart_Stop                         KhidusageBcs = 0
	KHIDUsage_BCS_TriOptic                                   KhidusageBcs = 0
	KHIDUsage_BCS_TriggerMode                                KhidusageBcs = 0
	KHIDUsage_BCS_TriggerModeBlinkingLaserOn                 KhidusageBcs = 0
	KHIDUsage_BCS_TriggerModeContinuousLaserOn               KhidusageBcs = 0
	KHIDUsage_BCS_TriggerModeLaserOnWhilePulled              KhidusageBcs = 0
	KHIDUsage_BCS_TriggerModeLaserStaysOnAfterTriggerRelease KhidusageBcs = 0
	KHIDUsage_BCS_TriggerReport                              KhidusageBcs = 0
	KHIDUsage_BCS_TriggerState                               KhidusageBcs = 0
	KHIDUsage_BCS_Triggerless                                KhidusageBcs = 0
	KHIDUsage_BCS_UCC_EAN_128                                KhidusageBcs = 0
	KHIDUsage_BCS_UPC_A                                      KhidusageBcs = 0
	KHIDUsage_BCS_UPC_AWith128Mandatory                      KhidusageBcs = 0
	KHIDUsage_BCS_UPC_AWith128Optical                        KhidusageBcs = 0
	KHIDUsage_BCS_UPC_AWithP5Optional                        KhidusageBcs = 0
	KHIDUsage_BCS_UPC_E                                      KhidusageBcs = 0
	KHIDUsage_BCS_UPC_E1                                     KhidusageBcs = 0
	KHIDUsage_BCS_UPC_EAN                                    KhidusageBcs = 0
	KHIDUsage_BCS_UPC_EANControlReport                       KhidusageBcs = 0
	KHIDUsage_BCS_UPC_EANCouponCode                          KhidusageBcs = 0
	KHIDUsage_BCS_UPC_EANPeriodicals                         KhidusageBcs = 0
	KHIDUsage_BCS_USB_5_SlugCode                             KhidusageBcs = 0
	KHIDUsage_BCS_UltraCode                                  KhidusageBcs = 0
	KHIDUsage_BCS_Undefined                                  KhidusageBcs = 0
	KHIDUsage_BCS_VeriCode                                   KhidusageBcs = 0
	KHIDUsage_BCS_Wand                                       KhidusageBcs = 0
	KHIDUsage_BCS_WaterResistant                             KhidusageBcs = 0
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
	KHIDUsage_BD_6DotBrailleCell           KhidusageBd = 0
	KHIDUsage_BD_8DotBrailleCell           KhidusageBd = 0
	KHIDUsage_BD_BrailleButtons            KhidusageBd = 0
	KHIDUsage_BD_BrailleDPadCenter         KhidusageBd = 0
	KHIDUsage_BD_BrailleDPadDown           KhidusageBd = 0
	KHIDUsage_BD_BrailleDPadLeft           KhidusageBd = 0
	KHIDUsage_BD_BrailleDPadRight          KhidusageBd = 0
	KHIDUsage_BD_BrailleDPadUp             KhidusageBd = 0
	KHIDUsage_BD_BrailleDisplay            KhidusageBd = 0
	KHIDUsage_BD_BrailleFaceControls       KhidusageBd = 0
	KHIDUsage_BD_BrailleJoystickCenter     KhidusageBd = 0
	KHIDUsage_BD_BrailleJoystickDown       KhidusageBd = 0
	KHIDUsage_BD_BrailleJoystickLeft       KhidusageBd = 0
	KHIDUsage_BD_BrailleJoystickRight      KhidusageBd = 0
	KHIDUsage_BD_BrailleJoystickUp         KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardDot1       KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardDot2       KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardDot3       KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardDot4       KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardDot5       KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardDot6       KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardDot7       KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardDot8       KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardLeftSpace  KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardRightSpace KhidusageBd = 0
	KHIDUsage_BD_BrailleKeyboardSpace      KhidusageBd = 0
	KHIDUsage_BD_BrailleLeftControls       KhidusageBd = 0
	KHIDUsage_BD_BraillePanLeft            KhidusageBd = 0
	KHIDUsage_BD_BraillePanRight           KhidusageBd = 0
	KHIDUsage_BD_BrailleRightControls      KhidusageBd = 0
	KHIDUsage_BD_BrailleRockerDown         KhidusageBd = 0
	KHIDUsage_BD_BrailleRockerPress        KhidusageBd = 0
	KHIDUsage_BD_BrailleRockerUp           KhidusageBd = 0
	KHIDUsage_BD_BrailleRow                KhidusageBd = 0
	KHIDUsage_BD_BrailleTopControls        KhidusageBd = 0
	KHIDUsage_BD_NumberOfBrailleCells      KhidusageBd = 0
	KHIDUsage_BD_RouterKey                 KhidusageBd = 0
	KHIDUsage_BD_RouterSet1                KhidusageBd = 0
	KHIDUsage_BD_RouterSet2                KhidusageBd = 0
	KHIDUsage_BD_RouterSet3                KhidusageBd = 0
	KHIDUsage_BD_RowRouterKey              KhidusageBd = 0
	KHIDUsage_BD_ScreenReaderControl       KhidusageBd = 0
	KHIDUsage_BD_ScreenReaderIdentifier    KhidusageBd = 0
	KHIDUsage_BD_Undefined                 KhidusageBd = 0
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
	KHIDUsage_BS_ACPresent                   KhidusageBs = 0
	KHIDUsage_BS_AbsoluteStateOfCharge       KhidusageBs = 0
	KHIDUsage_BS_AlarmInhibited              KhidusageBs = 0
	KHIDUsage_BS_AtRate                      KhidusageBs = 0
	KHIDUsage_BS_AtRateOK                    KhidusageBs = 0
	KHIDUsage_BS_AtRateTimeToEmpty           KhidusageBs = 0
	KHIDUsage_BS_AtRateTimeToFull            KhidusageBs = 0
	KHIDUsage_BS_AverageCurrent              KhidusageBs = 0
	KHIDUsage_BS_AverageTimeToEmpty          KhidusageBs = 0
	KHIDUsage_BS_AverageTimeToFull           KhidusageBs = 0
	KHIDUsage_BS_BattPackModelLevel          KhidusageBs = 0
	KHIDUsage_BS_BatteryInsertion            KhidusageBs = 0
	KHIDUsage_BS_BatteryPresent              KhidusageBs = 0
	KHIDUsage_BS_BatterySupported            KhidusageBs = 0
	KHIDUsage_BS_BelowRemainingCapacityLimit KhidusageBs = 0
	KHIDUsage_BS_BroadcastToCharger          KhidusageBs = 0
	KHIDUsage_BS_CapacityGranularity1        KhidusageBs = 0
	KHIDUsage_BS_CapacityGranularity2        KhidusageBs = 0
	KHIDUsage_BS_CapacityMode                KhidusageBs = 0
	KHIDUsage_BS_ChargeController            KhidusageBs = 0
	KHIDUsage_BS_ChargerConnection           KhidusageBs = 0
	KHIDUsage_BS_ChargerSelectorSupport      KhidusageBs = 0
	KHIDUsage_BS_ChargerSpec                 KhidusageBs = 0
	KHIDUsage_BS_Charging                    KhidusageBs = 0
	KHIDUsage_BS_ChargingIndicator           KhidusageBs = 0
	KHIDUsage_BS_ConditioningFlag            KhidusageBs = 0
	KHIDUsage_BS_ConnectionToSMBus           KhidusageBs = 0
	KHIDUsage_BS_CurrentNotRegulated         KhidusageBs = 0
	KHIDUsage_BS_CurrentOutOfRange           KhidusageBs = 0
	KHIDUsage_BS_CycleCount                  KhidusageBs = 0
	KHIDUsage_BS_DesignCapacity              KhidusageBs = 0
	KHIDUsage_BS_Discharging                 KhidusageBs = 0
	KHIDUsage_BS_EnablePolling               KhidusageBs = 0
	KHIDUsage_BS_FullChargeCapacity          KhidusageBs = 0
	KHIDUsage_BS_FullyCharged                KhidusageBs = 0
	KHIDUsage_BS_FullyDischarged             KhidusageBs = 0
	KHIDUsage_BS_InhibitCharge               KhidusageBs = 0
	KHIDUsage_BS_InternalChargeController    KhidusageBs = 0
	KHIDUsage_BS_Level2                      KhidusageBs = 0
	KHIDUsage_BS_Level3                      KhidusageBs = 0
	KHIDUsage_BS_ManufacturerAccess          KhidusageBs = 0
	KHIDUsage_BS_ManufacturerData            KhidusageBs = 0
	KHIDUsage_BS_ManufacturerDate            KhidusageBs = 0
	KHIDUsage_BS_MasterMode                  KhidusageBs = 0
	KHIDUsage_BS_Maxerror                    KhidusageBs = 0
	KHIDUsage_BS_NeedReplacement             KhidusageBs = 0
	KHIDUsage_BS_OKToUse                     KhidusageBs = 0
	KHIDUsage_BS_OptionalMfgFunction1        KhidusageBs = 0
	KHIDUsage_BS_OptionalMfgFunction2        KhidusageBs = 0
	KHIDUsage_BS_OptionalMfgFunction3        KhidusageBs = 0
	KHIDUsage_BS_OptionalMfgFunction4        KhidusageBs = 0
	KHIDUsage_BS_OptionalMfgFunction5        KhidusageBs = 0
	KHIDUsage_BS_OutputConnection            KhidusageBs = 0
	KHIDUsage_BS_PowerFail                   KhidusageBs = 0
	KHIDUsage_BS_PrimaryBattery              KhidusageBs = 0
	KHIDUsage_BS_PrimaryBatterySupport       KhidusageBs = 0
	KHIDUsage_BS_Rechargable                 KhidusageBs = 0
	KHIDUsage_BS_RelativeStateOfCharge       KhidusageBs = 0
	KHIDUsage_BS_RemainingCapacity           KhidusageBs = 0
	KHIDUsage_BS_RemainingCapacityLimit      KhidusageBs = 0
	KHIDUsage_BS_RemainingTimeLimit          KhidusageBs = 0
	KHIDUsage_BS_RemainingTimeLimitExpired   KhidusageBs = 0
	KHIDUsage_BS_ResetToZero                 KhidusageBs = 0
	KHIDUsage_BS_RunTimeToEmpty              KhidusageBs = 0
	KHIDUsage_BS_SMBAlarmWarning             KhidusageBs = 0
	KHIDUsage_BS_SMBBatteryMode              KhidusageBs = 0
	KHIDUsage_BS_SMBBatteryStatus            KhidusageBs = 0
	KHIDUsage_BS_SMBChargerMode              KhidusageBs = 0
	KHIDUsage_BS_SMBChargerSpecInfo          KhidusageBs = 0
	KHIDUsage_BS_SMBChargerStatus            KhidusageBs = 0
	KHIDUsage_BS_SMBErrorCode                KhidusageBs = 0
	KHIDUsage_BS_SMBSelectorInfo             KhidusageBs = 0
	KHIDUsage_BS_SMBSelectorPresets          KhidusageBs = 0
	KHIDUsage_BS_SMBSelectorState            KhidusageBs = 0
	KHIDUsage_BS_SelectorRevision            KhidusageBs = 0
	KHIDUsage_BS_SerialNumber                KhidusageBs = 0
	KHIDUsage_BS_SpecificationInfo           KhidusageBs = 0
	KHIDUsage_BS_TerminateCharge             KhidusageBs = 0
	KHIDUsage_BS_TerminateDischarge          KhidusageBs = 0
	KHIDUsage_BS_ThermistorCold              KhidusageBs = 0
	KHIDUsage_BS_ThermistorHot               KhidusageBs = 0
	KHIDUsage_BS_ThermistorOverRange         KhidusageBs = 0
	KHIDUsage_BS_ThermistorUnderRange        KhidusageBs = 0
	KHIDUsage_BS_Undefined                   KhidusageBs = 0
	KHIDUsage_BS_Usenext                     KhidusageBs = 0
	KHIDUsage_BS_VoltageNotRegulated         KhidusageBs = 0
	KHIDUsage_BS_VoltageOutOfRange           KhidusageBs = 0
	KHIDUsage_BS_WarningCapacityLimit        KhidusageBs = 0
	KHIDUsage_BS_iDeviceChemistry            KhidusageBs = 0
	KHIDUsage_BS_iDevicename                 KhidusageBs = 0
	KHIDUsage_BS_iManufacturerName           KhidusageBs = 0
	KHIDUsage_BS_iOEMInformation             KhidusageBs = 0
)

func (e KhidusageBs) String() string {
	switch e {
	case KHIDUsage_BS_ACPresent:
		return "KHIDUsage_BS_ACPresent"
	default:
		return fmt.Sprintf("KhidusageBs(%d)", e)
	}
}

type KhidusageButton uint

const (
	KHIDUsage_Button_1     KhidusageButton = 0
	KHIDUsage_Button_100   KhidusageButton = 0
	KHIDUsage_Button_104   KhidusageButton = 0
	KHIDUsage_Button_107   KhidusageButton = 0
	KHIDUsage_Button_109   KhidusageButton = 0
	KHIDUsage_Button_110   KhidusageButton = 0
	KHIDUsage_Button_112   KhidusageButton = 0
	KHIDUsage_Button_118   KhidusageButton = 0
	KHIDUsage_Button_119   KhidusageButton = 0
	KHIDUsage_Button_123   KhidusageButton = 0
	KHIDUsage_Button_124   KhidusageButton = 0
	KHIDUsage_Button_129   KhidusageButton = 0
	KHIDUsage_Button_137   KhidusageButton = 0
	KHIDUsage_Button_138   KhidusageButton = 0
	KHIDUsage_Button_14    KhidusageButton = 0
	KHIDUsage_Button_145   KhidusageButton = 0
	KHIDUsage_Button_149   KhidusageButton = 0
	KHIDUsage_Button_15    KhidusageButton = 0
	KHIDUsage_Button_150   KhidusageButton = 0
	KHIDUsage_Button_153   KhidusageButton = 0
	KHIDUsage_Button_157   KhidusageButton = 0
	KHIDUsage_Button_160   KhidusageButton = 0
	KHIDUsage_Button_170   KhidusageButton = 0
	KHIDUsage_Button_174   KhidusageButton = 0
	KHIDUsage_Button_175   KhidusageButton = 0
	KHIDUsage_Button_179   KhidusageButton = 0
	KHIDUsage_Button_18    KhidusageButton = 0
	KHIDUsage_Button_182   KhidusageButton = 0
	KHIDUsage_Button_186   KhidusageButton = 0
	KHIDUsage_Button_187   KhidusageButton = 0
	KHIDUsage_Button_19    KhidusageButton = 0
	KHIDUsage_Button_192   KhidusageButton = 0
	KHIDUsage_Button_201   KhidusageButton = 0
	KHIDUsage_Button_203   KhidusageButton = 0
	KHIDUsage_Button_211   KhidusageButton = 0
	KHIDUsage_Button_212   KhidusageButton = 0
	KHIDUsage_Button_22    KhidusageButton = 0
	KHIDUsage_Button_220   KhidusageButton = 0
	KHIDUsage_Button_222   KhidusageButton = 0
	KHIDUsage_Button_223   KhidusageButton = 0
	KHIDUsage_Button_226   KhidusageButton = 0
	KHIDUsage_Button_229   KhidusageButton = 0
	KHIDUsage_Button_23    KhidusageButton = 0
	KHIDUsage_Button_233   KhidusageButton = 0
	KHIDUsage_Button_252   KhidusageButton = 0
	KHIDUsage_Button_27    KhidusageButton = 0
	KHIDUsage_Button_30    KhidusageButton = 0
	KHIDUsage_Button_39    KhidusageButton = 0
	KHIDUsage_Button_40    KhidusageButton = 0
	KHIDUsage_Button_43    KhidusageButton = 0
	KHIDUsage_Button_47    KhidusageButton = 0
	KHIDUsage_Button_48    KhidusageButton = 0
	KHIDUsage_Button_49    KhidusageButton = 0
	KHIDUsage_Button_5     KhidusageButton = 0
	KHIDUsage_Button_62    KhidusageButton = 0
	KHIDUsage_Button_63    KhidusageButton = 0
	KHIDUsage_Button_64    KhidusageButton = 0
	KHIDUsage_Button_65535 KhidusageButton = 0
	KHIDUsage_Button_66    KhidusageButton = 0
	KHIDUsage_Button_67    KhidusageButton = 0
	KHIDUsage_Button_7     KhidusageButton = 0
	KHIDUsage_Button_71    KhidusageButton = 0
	KHIDUsage_Button_74    KhidusageButton = 0
	KHIDUsage_Button_84    KhidusageButton = 0
	KHIDUsage_Button_86    KhidusageButton = 0
	KHIDUsage_Button_87    KhidusageButton = 0
	KHIDUsage_Button_89    KhidusageButton = 0
	KHIDUsage_Button_9     KhidusageButton = 0
	KHIDUsage_Button_95    KhidusageButton = 0
)

func (e KhidusageButton) String() string {
	switch e {
	case KHIDUsage_Button_1:
		return "KHIDUsage_Button_1"
	default:
		return fmt.Sprintf("KhidusageButton(%d)", e)
	}
}

type KhidusageCc uint

const (
	KHIDUsage_CC_Autofocus KhidusageCc = 0
	KHIDUsage_CC_Shutter   KhidusageCc = 0
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
	KHIDUsage_Csmr_Repeat        KhidusageCsmr = 0
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
	KHIDUsage_Dig_3DDigitizer                       KhidusageDig = 0
	KHIDUsage_Dig_Altitude                          KhidusageDig = 0
	KHIDUsage_Dig_Armature                          KhidusageDig = 0
	KHIDUsage_Dig_ArticulatedArm                    KhidusageDig = 0
	KHIDUsage_Dig_Azimuth                           KhidusageDig = 0
	KHIDUsage_Dig_BarrelPressure                    KhidusageDig = 0
	KHIDUsage_Dig_BarrelSwitch                      KhidusageDig = 0
	KHIDUsage_Dig_BatteryStrength                   KhidusageDig = 0
	KHIDUsage_Dig_CapacitiveHeatMapDigitizer        KhidusageDig = 0
	KHIDUsage_Dig_CapacitiveHeatMapFrameData        KhidusageDig = 0
	KHIDUsage_Dig_CapacitiveHeatMapProtocolVendorID KhidusageDig = 0
	KHIDUsage_Dig_CapacitiveHeatMapProtocolVersion  KhidusageDig = 0
	KHIDUsage_Dig_ContactCount                      KhidusageDig = 0
	KHIDUsage_Dig_ContactCountMaximum               KhidusageDig = 0
	KHIDUsage_Dig_ContactIdentifier                 KhidusageDig = 0
	KHIDUsage_Dig_CoordinateMeasuringMachine        KhidusageDig = 0
	KHIDUsage_Dig_DataValid                         KhidusageDig = 0
	KHIDUsage_Dig_DeviceConfiguration               KhidusageDig = 0
	KHIDUsage_Dig_DeviceIdentifier                  KhidusageDig = 0
	KHIDUsage_Dig_DeviceMode                        KhidusageDig = 0
	KHIDUsage_Dig_DeviceSettings                    KhidusageDig = 0
	KHIDUsage_Dig_Digitizer                         KhidusageDig = 0
	KHIDUsage_Dig_Eraser                            KhidusageDig = 0
	KHIDUsage_Dig_Finger                            KhidusageDig = 0
	KHIDUsage_Dig_FreeSpaceWand                     KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacter                  KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterData              KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterDataLength        KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterEnable            KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterEncoding          KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterEncodingUTF16BE   KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterEncodingUTF16LE   KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterEncodingUTF32BE   KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterEncodingUTF32LE   KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterEncodingUTF8      KhidusageDig = 0
	KHIDUsage_Dig_GestureCharacterQuality           KhidusageDig = 0
	KHIDUsage_Dig_Height                            KhidusageDig = 0
	KHIDUsage_Dig_InRange                           KhidusageDig = 0
	KHIDUsage_Dig_Invert                            KhidusageDig = 0
	KHIDUsage_Dig_LightPen                          KhidusageDig = 0
	KHIDUsage_Dig_MultiplePointDigitizer            KhidusageDig = 0
	KHIDUsage_Dig_Pen                               KhidusageDig = 0
	KHIDUsage_Dig_ProgramChangeKeys                 KhidusageDig = 0
	KHIDUsage_Dig_Puck                              KhidusageDig = 0
	KHIDUsage_Dig_Quality                           KhidusageDig = 0
	KHIDUsage_Dig_Reserved                          KhidusageDig = 0
	KHIDUsage_Dig_SecondaryTipSwitch                KhidusageDig = 0
	KHIDUsage_Dig_StereoPlotter                     KhidusageDig = 0
	KHIDUsage_Dig_Stylus                            KhidusageDig = 0
	KHIDUsage_Dig_TabletFunctionKeys                KhidusageDig = 0
	KHIDUsage_Dig_TabletPick                        KhidusageDig = 0
	KHIDUsage_Dig_Tap                               KhidusageDig = 0
	KHIDUsage_Dig_TipPressure                       KhidusageDig = 0
	KHIDUsage_Dig_TipSwitch                         KhidusageDig = 0
	KHIDUsage_Dig_Touch                             KhidusageDig = 0
	KHIDUsage_Dig_TouchPad                          KhidusageDig = 0
	KHIDUsage_Dig_TouchScreen                       KhidusageDig = 0
	KHIDUsage_Dig_TouchValid                        KhidusageDig = 0
	KHIDUsage_Dig_TransducerIndex                   KhidusageDig = 0
	KHIDUsage_Dig_Twist                             KhidusageDig = 0
	KHIDUsage_Dig_Untouch                           KhidusageDig = 0
	KHIDUsage_Dig_WhiteBoard                        KhidusageDig = 0
	KHIDUsage_Dig_Width                             KhidusageDig = 0
	KHIDUsage_Dig_XTilt                             KhidusageDig = 0
	KHIDUsage_Dig_YTilt                             KhidusageDig = 0
)

func (e KhidusageDig) String() string {
	switch e {
	case KHIDUsage_Dig_3DDigitizer:
		return "KHIDUsage_Dig_3DDigitizer"
	default:
		return fmt.Sprintf("KhidusageDig(%d)", e)
	}
}

type KhidusageGame uint

const (
	KHIDUsage_Game_3DGameController                 KhidusageGame = 0
	KHIDUsage_Game_Bump                             KhidusageGame = 0
	KHIDUsage_Game_Flipper                          KhidusageGame = 0
	KHIDUsage_Game_GamepadFireOrJump                KhidusageGame = 0
	KHIDUsage_Game_GamepadFormFitting               KhidusageGame = 0
	KHIDUsage_Game_GamepadFormFitting_Compatibility KhidusageGame = 0
	KHIDUsage_Game_GamepadTrigger                   KhidusageGame = 0
	KHIDUsage_Game_Gun                              KhidusageGame = 0
	KHIDUsage_Game_GunAutomatic                     KhidusageGame = 0
	KHIDUsage_Game_GunBolt                          KhidusageGame = 0
	KHIDUsage_Game_GunBurst                         KhidusageGame = 0
	KHIDUsage_Game_GunClip                          KhidusageGame = 0
	KHIDUsage_Game_GunDevice                        KhidusageGame = 0
	KHIDUsage_Game_GunSafety                        KhidusageGame = 0
	KHIDUsage_Game_GunSingleShot                    KhidusageGame = 0
	KHIDUsage_Game_HeightOfPOV                      KhidusageGame = 0
	KHIDUsage_Game_LeanForwardOrBackward            KhidusageGame = 0
	KHIDUsage_Game_LeanRightOrLeft                  KhidusageGame = 0
	KHIDUsage_Game_MoveForwardOrBackward            KhidusageGame = 0
	KHIDUsage_Game_MoveRightOrLeft                  KhidusageGame = 0
	KHIDUsage_Game_MoveUpOrDown                     KhidusageGame = 0
	KHIDUsage_Game_NewGame                          KhidusageGame = 0
	KHIDUsage_Game_PinballDevice                    KhidusageGame = 0
	KHIDUsage_Game_PitchUpOrDown                    KhidusageGame = 0
	KHIDUsage_Game_Player                           KhidusageGame = 0
	KHIDUsage_Game_PointofView                      KhidusageGame = 0
	KHIDUsage_Game_Reserved                         KhidusageGame = 0
	KHIDUsage_Game_RollRightOrLeft                  KhidusageGame = 0
	KHIDUsage_Game_SecondaryFlipper                 KhidusageGame = 0
	KHIDUsage_Game_ShootBall                        KhidusageGame = 0
	KHIDUsage_Game_TurnRightOrLeft                  KhidusageGame = 0
)

func (e KhidusageGame) String() string {
	switch e {
	case KHIDUsage_Game_3DGameController:
		return "KHIDUsage_Game_3DGameController"
	default:
		return fmt.Sprintf("KhidusageGame(%d)", e)
	}
}

type KhidusageGd uint

const (
	KHIDUsage_GD_ByteCount  KhidusageGd = 0
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

type KhidusageGendevcontrols uint

const (
	KHIDUsage_GenDevControls_BackgroundControls KhidusageGendevcontrols = 0
)

func (e KhidusageGendevcontrols) String() string {
	switch e {
	case KHIDUsage_GenDevControls_BackgroundControls:
		return "KHIDUsage_GenDevControls_BackgroundControls"
	default:
		return fmt.Sprintf("KhidusageGendevcontrols(%d)", e)
	}
}

type KhidusageLed uint

const (
	KHIDUsage_LED_BatteryLow              KhidusageLed = 0
	KHIDUsage_LED_BatteryOK               KhidusageLed = 0
	KHIDUsage_LED_BatteryOperation        KhidusageLed = 0
	KHIDUsage_LED_Busy                    KhidusageLed = 0
	KHIDUsage_LED_CAV                     KhidusageLed = 0
	KHIDUsage_LED_CLV                     KhidusageLed = 0
	KHIDUsage_LED_CallPickup              KhidusageLed = 0
	KHIDUsage_LED_CameraOff               KhidusageLed = 0
	KHIDUsage_LED_CameraOn                KhidusageLed = 0
	KHIDUsage_LED_CapsLock                KhidusageLed = 0
	KHIDUsage_LED_Compose                 KhidusageLed = 0
	KHIDUsage_LED_Conference              KhidusageLed = 0
	KHIDUsage_LED_Coverage                KhidusageLed = 0
	KHIDUsage_LED_DataMode                KhidusageLed = 0
	KHIDUsage_LED_DoNotDisturb            KhidusageLed = 0
	KHIDUsage_LED_EqualizerEnable         KhidusageLed = 0
	KHIDUsage_LED_Error                   KhidusageLed = 0
	KHIDUsage_LED_FastBlinkOffTime        KhidusageLed = 0
	KHIDUsage_LED_FastBlinkOnTime         KhidusageLed = 0
	KHIDUsage_LED_FastForward             KhidusageLed = 0
	KHIDUsage_LED_FlashOnTime             KhidusageLed = 0
	KHIDUsage_LED_Forward                 KhidusageLed = 0
	KHIDUsage_LED_GenericIndicator        KhidusageLed = 0
	KHIDUsage_LED_HeadSet                 KhidusageLed = 0
	KHIDUsage_LED_HighCutFilter           KhidusageLed = 0
	KHIDUsage_LED_Hold                    KhidusageLed = 0
	KHIDUsage_LED_IndicatorAmber          KhidusageLed = 0
	KHIDUsage_LED_IndicatorFastBlink      KhidusageLed = 0
	KHIDUsage_LED_IndicatorFlash          KhidusageLed = 0
	KHIDUsage_LED_IndicatorGreen          KhidusageLed = 0
	KHIDUsage_LED_IndicatorOff            KhidusageLed = 0
	KHIDUsage_LED_IndicatorOn             KhidusageLed = 0
	KHIDUsage_LED_IndicatorRed            KhidusageLed = 0
	KHIDUsage_LED_IndicatorSlowBlink      KhidusageLed = 0
	KHIDUsage_LED_Kana                    KhidusageLed = 0
	KHIDUsage_LED_LowCutFilter            KhidusageLed = 0
	KHIDUsage_LED_MessageWaiting          KhidusageLed = 0
	KHIDUsage_LED_Microphone              KhidusageLed = 0
	KHIDUsage_LED_Mute                    KhidusageLed = 0
	KHIDUsage_LED_NightMode               KhidusageLed = 0
	KHIDUsage_LED_NumLock                 KhidusageLed = 0
	KHIDUsage_LED_OffHook                 KhidusageLed = 0
	KHIDUsage_LED_OffLine                 KhidusageLed = 0
	KHIDUsage_LED_OnLine                  KhidusageLed = 0
	KHIDUsage_LED_PaperJam                KhidusageLed = 0
	KHIDUsage_LED_PaperOut                KhidusageLed = 0
	KHIDUsage_LED_Pause                   KhidusageLed = 0
	KHIDUsage_LED_Play                    KhidusageLed = 0
	KHIDUsage_LED_Player1                 KhidusageLed = 0
	KHIDUsage_LED_Player2                 KhidusageLed = 0
	KHIDUsage_LED_Player3                 KhidusageLed = 0
	KHIDUsage_LED_Player4                 KhidusageLed = 0
	KHIDUsage_LED_Player5                 KhidusageLed = 0
	KHIDUsage_LED_Player6                 KhidusageLed = 0
	KHIDUsage_LED_Player7                 KhidusageLed = 0
	KHIDUsage_LED_Player8                 KhidusageLed = 0
	KHIDUsage_LED_PlayerIndicator         KhidusageLed = 0
	KHIDUsage_LED_Power                   KhidusageLed = 0
	KHIDUsage_LED_Ready                   KhidusageLed = 0
	KHIDUsage_LED_Record                  KhidusageLed = 0
	KHIDUsage_LED_RecordingFormatDetect   KhidusageLed = 0
	KHIDUsage_LED_Remote                  KhidusageLed = 0
	KHIDUsage_LED_Repeat                  KhidusageLed = 0
	KHIDUsage_LED_Reserved                KhidusageLed = 0
	KHIDUsage_LED_Reverse                 KhidusageLed = 0
	KHIDUsage_LED_Rewind                  KhidusageLed = 0
	KHIDUsage_LED_Ring                    KhidusageLed = 0
	KHIDUsage_LED_SamplingRateDetect      KhidusageLed = 0
	KHIDUsage_LED_ScrollLock              KhidusageLed = 0
	KHIDUsage_LED_SendCalls               KhidusageLed = 0
	KHIDUsage_LED_Shift                   KhidusageLed = 0
	KHIDUsage_LED_SlowBlinkOffTime        KhidusageLed = 0
	KHIDUsage_LED_SlowBlinkOnTime         KhidusageLed = 0
	KHIDUsage_LED_SoundFieldOn            KhidusageLed = 0
	KHIDUsage_LED_Speaker                 KhidusageLed = 0
	KHIDUsage_LED_Spinning                KhidusageLed = 0
	KHIDUsage_LED_StandBy                 KhidusageLed = 0
	KHIDUsage_LED_Stereo                  KhidusageLed = 0
	KHIDUsage_LED_Stop                    KhidusageLed = 0
	KHIDUsage_LED_SurroundOn              KhidusageLed = 0
	KHIDUsage_LED_SystemSuspend           KhidusageLed = 0
	KHIDUsage_LED_ToneEnable              KhidusageLed = 0
	KHIDUsage_LED_Usage                   KhidusageLed = 0
	KHIDUsage_LED_UsageInUseIndicator     KhidusageLed = 0
	KHIDUsage_LED_UsageIndicatorColor     KhidusageLed = 0
	KHIDUsage_LED_UsageMultiModeIndicator KhidusageLed = 0
)

func (e KhidusageLed) String() string {
	switch e {
	case KHIDUsage_LED_BatteryLow:
		return "KHIDUsage_LED_BatteryLow"
	default:
		return fmt.Sprintf("KhidusageLed(%d)", e)
	}
}

type KhidusageMsr uint

const (
	KHIDUsage_MSR_DeviceReadOnly KhidusageMsr = 0
	KHIDUsage_MSR_Track1Data     KhidusageMsr = 0
	KHIDUsage_MSR_Track1Length   KhidusageMsr = 0
	KHIDUsage_MSR_Track2Data     KhidusageMsr = 0
	KHIDUsage_MSR_Track2Length   KhidusageMsr = 0
	KHIDUsage_MSR_Track3Data     KhidusageMsr = 0
	KHIDUsage_MSR_Track3Length   KhidusageMsr = 0
	KHIDUsage_MSR_TrackData      KhidusageMsr = 0
	KHIDUsage_MSR_TrackJISData   KhidusageMsr = 0
	KHIDUsage_MSR_TrackJISLength KhidusageMsr = 0
	KHIDUsage_MSR_Undefined      KhidusageMsr = 0
)

func (e KhidusageMsr) String() string {
	switch e {
	case KHIDUsage_MSR_DeviceReadOnly:
		return "KHIDUsage_MSR_DeviceReadOnly"
	default:
		return fmt.Sprintf("KhidusageMsr(%d)", e)
	}
}

type KhidusageOrd uint

const (
	KHIDUsage_Ord_Instance1 KhidusageOrd = 0
)

func (e KhidusageOrd) String() string {
	switch e {
	case KHIDUsage_Ord_Instance1:
		return "KHIDUsage_Ord_Instance1"
	default:
		return fmt.Sprintf("KhidusageOrd(%d)", e)
	}
}

type KhidusagePd uint

const (
	KHIDUsage_PD_PowerSummary      KhidusagePd = 0
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
	KHIDUsage_PID_ActuatorsEnabled   KhidusagePid = 0
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
	KHIDUsage_Sim_Accelerator                 KhidusageSim = 0
	KHIDUsage_Sim_Aileron                     KhidusageSim = 0
	KHIDUsage_Sim_AileronTrim                 KhidusageSim = 0
	KHIDUsage_Sim_AirplaneSimulationDevice    KhidusageSim = 0
	KHIDUsage_Sim_AntiTorqueControl           KhidusageSim = 0
	KHIDUsage_Sim_AutomobileSimulationDevice  KhidusageSim = 0
	KHIDUsage_Sim_AutopilotEnable             KhidusageSim = 0
	KHIDUsage_Sim_Ballast                     KhidusageSim = 0
	KHIDUsage_Sim_BarrelElevation             KhidusageSim = 0
	KHIDUsage_Sim_BicycleCrank                KhidusageSim = 0
	KHIDUsage_Sim_BicycleSimulationDevice     KhidusageSim = 0
	KHIDUsage_Sim_Brake                       KhidusageSim = 0
	KHIDUsage_Sim_ChaffRelease                KhidusageSim = 0
	KHIDUsage_Sim_Clutch                      KhidusageSim = 0
	KHIDUsage_Sim_CollectiveControl           KhidusageSim = 0
	KHIDUsage_Sim_CyclicControl               KhidusageSim = 0
	KHIDUsage_Sim_CyclicTrim                  KhidusageSim = 0
	KHIDUsage_Sim_DiveBrake                   KhidusageSim = 0
	KHIDUsage_Sim_DivePlane                   KhidusageSim = 0
	KHIDUsage_Sim_ElectronicCountermeasures   KhidusageSim = 0
	KHIDUsage_Sim_Elevator                    KhidusageSim = 0
	KHIDUsage_Sim_ElevatorTrim                KhidusageSim = 0
	KHIDUsage_Sim_FlareRelease                KhidusageSim = 0
	KHIDUsage_Sim_FlightCommunications        KhidusageSim = 0
	KHIDUsage_Sim_FlightControlStick          KhidusageSim = 0
	KHIDUsage_Sim_FlightSimulationDevice      KhidusageSim = 0
	KHIDUsage_Sim_FlightStick                 KhidusageSim = 0
	KHIDUsage_Sim_FlightYoke                  KhidusageSim = 0
	KHIDUsage_Sim_FrontBrake                  KhidusageSim = 0
	KHIDUsage_Sim_HandleBars                  KhidusageSim = 0
	KHIDUsage_Sim_HelicopterSimulationDevice  KhidusageSim = 0
	KHIDUsage_Sim_LandingGear                 KhidusageSim = 0
	KHIDUsage_Sim_MagicCarpetSimulationDevice KhidusageSim = 0
	KHIDUsage_Sim_MotorcycleSimulationDevice  KhidusageSim = 0
	KHIDUsage_Sim_RearBrake                   KhidusageSim = 0
	KHIDUsage_Sim_Reserved                    KhidusageSim = 0
	KHIDUsage_Sim_Rudder                      KhidusageSim = 0
	KHIDUsage_Sim_SailingSimulationDevice     KhidusageSim = 0
	KHIDUsage_Sim_Shifter                     KhidusageSim = 0
	KHIDUsage_Sim_SpaceshipSimulationDevice   KhidusageSim = 0
	KHIDUsage_Sim_SportsSimulationDevice      KhidusageSim = 0
	KHIDUsage_Sim_Steering                    KhidusageSim = 0
	KHIDUsage_Sim_SubmarineSimulationDevice   KhidusageSim = 0
	KHIDUsage_Sim_TankSimulationDevice        KhidusageSim = 0
	KHIDUsage_Sim_Throttle                    KhidusageSim = 0
	KHIDUsage_Sim_ToeBrake                    KhidusageSim = 0
	KHIDUsage_Sim_TrackControl                KhidusageSim = 0
	KHIDUsage_Sim_Trigger                     KhidusageSim = 0
	KHIDUsage_Sim_TurretDirection             KhidusageSim = 0
	KHIDUsage_Sim_Weapons                     KhidusageSim = 0
	KHIDUsage_Sim_WeaponsArm                  KhidusageSim = 0
	KHIDUsage_Sim_WingFlaps                   KhidusageSim = 0
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
	KHIDUsage_Snsr_Data_Motion_AngularVelocityXAxis  KhidusageSnsr = 0
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
	KHIDUsage_Sprt_11Iron     KhidusageSprt = 0
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
	KHIDUsage_Tfon_Redial       KhidusageTfon = 0
)

func (e KhidusageTfon) String() string {
	switch e {
	case KHIDUsage_Tfon_LineBusyTone:
		return "KHIDUsage_Tfon_LineBusyTone"
	default:
		return fmt.Sprintf("KhidusageTfon(%d)", e)
	}
}

type KhidusageVr uint

const (
	KHIDUsage_VR_AnimatronicDevice  KhidusageVr = 0
	KHIDUsage_VR_Belt               KhidusageVr = 0
	KHIDUsage_VR_BodySuit           KhidusageVr = 0
	KHIDUsage_VR_DisplayEnable      KhidusageVr = 0
	KHIDUsage_VR_Flexor             KhidusageVr = 0
	KHIDUsage_VR_Glove              KhidusageVr = 0
	KHIDUsage_VR_HandTracker        KhidusageVr = 0
	KHIDUsage_VR_HeadMountedDisplay KhidusageVr = 0
	KHIDUsage_VR_HeadTracker        KhidusageVr = 0
	KHIDUsage_VR_Oculometer         KhidusageVr = 0
	KHIDUsage_VR_Reserved           KhidusageVr = 0
	KHIDUsage_VR_StereoEnable       KhidusageVr = 0
	KHIDUsage_VR_Vest               KhidusageVr = 0
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
	KHIDUsage_WD_CalibrationCount               KhidusageWd = 0
	KHIDUsage_WD_DataScaling                    KhidusageWd = 0
	KHIDUsage_WD_DataWeight                     KhidusageWd = 0
	KHIDUsage_WD_EnforcedZeroReturn             KhidusageWd = 0
	KHIDUsage_WD_RezeroCount                    KhidusageWd = 0
	KHIDUsage_WD_ScaleAtrributeReport           KhidusageWd = 0
	KHIDUsage_WD_ScaleControlReport             KhidusageWd = 0
	KHIDUsage_WD_ScaleDataReport                KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassGeneric         KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassIIIEnglish      KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassIIILEnglish     KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassIIILMetric      KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassIIIMetric       KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassIIMetric        KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassIMetric         KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassIMetricCL       KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassIVEnglish       KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleClassIVMetric        KhidusageWd = 0
	KHIDUsage_WD_ScaleScaleDevice               KhidusageWd = 0
	KHIDUsage_WD_ScaleStatisticsReport          KhidusageWd = 0
	KHIDUsage_WD_ScaleStatus                    KhidusageWd = 0
	KHIDUsage_WD_ScaleStatusFault               KhidusageWd = 0
	KHIDUsage_WD_ScaleStatusInMotion            KhidusageWd = 0
	KHIDUsage_WD_ScaleStatusOverWeightLimit     KhidusageWd = 0
	KHIDUsage_WD_ScaleStatusReport              KhidusageWd = 0
	KHIDUsage_WD_ScaleStatusRequiresCalibration KhidusageWd = 0
	KHIDUsage_WD_ScaleStatusRequiresRezeroing   KhidusageWd = 0
	KHIDUsage_WD_ScaleStatusStableAtZero        KhidusageWd = 0
	KHIDUsage_WD_ScaleStatusUnderZero           KhidusageWd = 0
	KHIDUsage_WD_ScaleStatusWeightStable        KhidusageWd = 0
	KHIDUsage_WD_ScaleWeightLimitReport         KhidusageWd = 0
	KHIDUsage_WD_Undefined                      KhidusageWd = 0
	KHIDUsage_WD_WeighingDevice                 KhidusageWd = 0
	KHIDUsage_WD_WeightUnit                     KhidusageWd = 0
	KHIDUsage_WD_WeightUnitAvoirTon             KhidusageWd = 0
	KHIDUsage_WD_WeightUnitCarats               KhidusageWd = 0
	KHIDUsage_WD_WeightUnitGrains               KhidusageWd = 0
	KHIDUsage_WD_WeightUnitGram                 KhidusageWd = 0
	KHIDUsage_WD_WeightUnitKilogram             KhidusageWd = 0
	KHIDUsage_WD_WeightUnitMetricTon            KhidusageWd = 0
	KHIDUsage_WD_WeightUnitMilligram            KhidusageWd = 0
	KHIDUsage_WD_WeightUnitOunce                KhidusageWd = 0
	KHIDUsage_WD_WeightUnitPennyweights         KhidusageWd = 0
	KHIDUsage_WD_WeightUnitPound                KhidusageWd = 0
	KHIDUsage_WD_WeightUnitTaels                KhidusageWd = 0
	KHIDUsage_WD_WeightUnitTroyOunce            KhidusageWd = 0
	KHIDUsage_WD_ZeroScale                      KhidusageWd = 0
)

func (e KhidusageWd) String() string {
	switch e {
	case KHIDUsage_WD_CalibrationCount:
		return "KHIDUsage_WD_CalibrationCount"
	default:
		return fmt.Sprintf("KhidusageWd(%d)", e)
	}
}

type KhvIonAny uint

const (
	KHV_ION_ANY_SIZE KhvIonAny = 0
)

func (e KhvIonAny) String() string {
	switch e {
	case KHV_ION_ANY_SIZE:
		return "KHV_ION_ANY_SIZE"
	default:
		return fmt.Sprintf("KhvIonAny(%d)", e)
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
	KINQUIRY_Byte56_Offset           KinquiryByte56 = 0
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
	// KINQUIRY_Page83_AssociationLogicalUnit: # Discussion
	KINQUIRY_Page83_AssociationLogicalUnit KinquiryPage83Association = 0
	// KINQUIRY_Page83_AssociationMask: # Discussion
	KINQUIRY_Page83_AssociationMask  KinquiryPage83Association = 0
	KINQUIRY_Page83_AssociationShift KinquiryPage83Association = 0
	// KINQUIRY_Page83_AssociationTargetDevice: # Discussion
	KINQUIRY_Page83_AssociationTargetDevice KinquiryPage83Association = 0
	// KINQUIRY_Page83_AssociationTargetPort: # Discussion
	KINQUIRY_Page83_AssociationTargetPort KinquiryPage83Association = 0
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
	KINQUIRY_Page83_CodeSetMask     KinquiryPage83Codeset = 0
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

type KinquiryPageb1Page uint

const (
	KINQUIRY_PageB1_Page_Length KinquiryPageb1Page = 0
)

func (e KinquiryPageb1Page) String() string {
	switch e {
	case KINQUIRY_PageB1_Page_Length:
		return "KINQUIRY_PageB1_Page_Length"
	default:
		return fmt.Sprintf("KinquiryPageb1Page(%d)", e)
	}
}

type KinquiryPagec0FeaturesHassep uint

const (
	KINQUIRY_PageC0_Features_HasSEP_LUN KinquiryPagec0FeaturesHassep = 0
)

func (e KinquiryPagec0FeaturesHassep) String() string {
	switch e {
	case KINQUIRY_PageC0_Features_HasSEP_LUN:
		return "KINQUIRY_PageC0_Features_HasSEP_LUN"
	default:
		return fmt.Sprintf("KinquiryPagec0FeaturesHassep(%d)", e)
	}
}

type KinquiryPeripheralQualifier uint

const (
	// KINQUIRY_PERIPHERAL_QUALIFIER_Connected: # Discussion
	KINQUIRY_PERIPHERAL_QUALIFIER_Connected KinquiryPeripheralQualifier = 0
	// KINQUIRY_PERIPHERAL_QUALIFIER_Mask: # Discussion
	KINQUIRY_PERIPHERAL_QUALIFIER_Mask KinquiryPeripheralQualifier = 0
	// KINQUIRY_PERIPHERAL_QUALIFIER_NotSupported: # Discussion
	KINQUIRY_PERIPHERAL_QUALIFIER_NotSupported KinquiryPeripheralQualifier = 0
	// KINQUIRY_PERIPHERAL_QUALIFIER_SupportedButNotConnected: # Discussion
	KINQUIRY_PERIPHERAL_QUALIFIER_SupportedButNotConnected KinquiryPeripheralQualifier = 0
)

func (e KinquiryPeripheralQualifier) String() string {
	switch e {
	case KINQUIRY_PERIPHERAL_QUALIFIER_Connected:
		return "KINQUIRY_PERIPHERAL_QUALIFIER_Connected"
	default:
		return fmt.Sprintf("KinquiryPeripheralQualifier(%d)", e)
	}
}

type KinquiryPeripheralRmb uint

const (
	// KINQUIRY_PERIPHERAL_RMB_BitMask: # Discussion
	KINQUIRY_PERIPHERAL_RMB_BitMask KinquiryPeripheralRmb = 0
	// KINQUIRY_PERIPHERAL_RMB_MediumFixed: # Discussion
	KINQUIRY_PERIPHERAL_RMB_MediumFixed KinquiryPeripheralRmb = 0
	// KINQUIRY_PERIPHERAL_RMB_MediumRemovable: # Discussion
	KINQUIRY_PERIPHERAL_RMB_MediumRemovable KinquiryPeripheralRmb = 0
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
	// KINQUIRY_PERIPHERAL_TYPE_AutomationDriveInterface: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_AutomationDriveInterface KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_CDROM_MMCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_CDROM_MMCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_CommunicationsSSCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_CommunicationsSSCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_DirectAccessSBCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_DirectAccessSBCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_EnclosureServicesSESDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_EnclosureServicesSESDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_Mask: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_Mask KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_MediumChangerSMCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_MediumChangerSMCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_ObjectBasedStorageDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_ObjectBasedStorageDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_OpticalCardReaderOCRWDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_OpticalCardReaderOCRWDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_OpticalMemorySBCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_OpticalMemorySBCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_PrinterSSCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_PrinterSSCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_ProcessorSPCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_ProcessorSPCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_ScannerSCSI2Device: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_ScannerSCSI2Device KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_SequentialAccessSSCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_SequentialAccessSSCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_SimplifiedDirectAccessRBCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_SimplifiedDirectAccessRBCDevice KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_StorageArrayControllerSCC2Device: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_StorageArrayControllerSCC2Device KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_UnknownOrNoDeviceType: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_UnknownOrNoDeviceType KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_WellKnownLogicalUnit: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_WellKnownLogicalUnit KinquiryPeripheralType = 0
	// KINQUIRY_PERIPHERAL_TYPE_WriteOnceSBCDevice: # Discussion
	KINQUIRY_PERIPHERAL_TYPE_WriteOnceSBCDevice KinquiryPeripheralType = 0
)

func (e KinquiryPeripheralType) String() string {
	switch e {
	case KINQUIRY_PERIPHERAL_TYPE_AutomationDriveInterface:
		return "KINQUIRY_PERIPHERAL_TYPE_AutomationDriveInterface"
	default:
		return fmt.Sprintf("KinquiryPeripheralType(%d)", e)
	}
}

type Kioanalogsignallevel0700 uint

const (
	KIOAnalogSignalLevel_0700_0000 Kioanalogsignallevel0700 = 0
)

func (e Kioanalogsignallevel0700) String() string {
	switch e {
	case KIOAnalogSignalLevel_0700_0000:
		return "KIOAnalogSignalLevel_0700_0000"
	default:
		return fmt.Sprintf("Kioanalogsignallevel0700(%d)", e)
	}
}

type KioaudiochannellabelAmbisonic uint

const (
	KIOAudioChannelLabel_Ambisonic_W KioaudiochannellabelAmbisonic = 0
)

func (e KioaudiochannellabelAmbisonic) String() string {
	switch e {
	case KIOAudioChannelLabel_Ambisonic_W:
		return "KIOAudioChannelLabel_Ambisonic_W"
	default:
		return fmt.Sprintf("KioaudiochannellabelAmbisonic(%d)", e)
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

type KmodepageformatP uint

const (
	// KModePageFormat_PAGE_CODE_Mask: # Discussion
	KModePageFormat_PAGE_CODE_Mask KmodepageformatP = 0
	// KModePageFormat_PS_Bit: # Discussion
	KModePageFormat_PS_Bit KmodepageformatP = 0
	// KModePageFormat_PS_Mask: # Discussion
	KModePageFormat_PS_Mask KmodepageformatP = 0
)

func (e KmodepageformatP) String() string {
	switch e {
	case KModePageFormat_PAGE_CODE_Mask:
		return "KModePageFormat_PAGE_CODE_Mask"
	default:
		return fmt.Sprintf("KmodepageformatP(%d)", e)
	}
}

type Kmodesenseparameterheader10Longlba uint

const (
	// KModeSenseParameterHeader10_LongLBABit: # Discussion
	KModeSenseParameterHeader10_LongLBABit Kmodesenseparameterheader10Longlba = 0
	// KModeSenseParameterHeader10_LongLBAMask: # Discussion
	KModeSenseParameterHeader10_LongLBAMask Kmodesenseparameterheader10Longlba = 0
)

func (e Kmodesenseparameterheader10Longlba) String() string {
	switch e {
	case KModeSenseParameterHeader10_LongLBABit:
		return "KModeSenseParameterHeader10_LongLBABit"
	default:
		return fmt.Sprintf("Kmodesenseparameterheader10Longlba(%d)", e)
	}
}

type Kpc uint

const (
	KPC_CAPTURED_PC     Kpc = 0
	KPC_KERNEL_COUNTING Kpc = 0
	KPC_KERNEL_PC       Kpc = 0
	KPC_USER_COUNTING   Kpc = 0
)

func (e Kpc) String() string {
	switch e {
	case KPC_CAPTURED_PC:
		return "KPC_CAPTURED_PC"
	default:
		return fmt.Sprintf("Kpc(%d)", e)
	}
}

type KreadCapacityProt uint

const (
	// KREAD_CAPACITY_PROT_Disabled: # Discussion
	KREAD_CAPACITY_PROT_Disabled KreadCapacityProt = 0
)

func (e KreadCapacityProt) String() string {
	switch e {
	case KREAD_CAPACITY_PROT_Disabled:
		return "KREAD_CAPACITY_PROT_Disabled"
	default:
		return fmt.Sprintf("KreadCapacityProt(%d)", e)
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

type KreportCapacity16 uint

const (
	// KREPORT_CAPACITY_16_DataSize: # Discussion
	KREPORT_CAPACITY_16_DataSize KreportCapacity16 = 0
)

func (e KreportCapacity16) String() string {
	switch e {
	case KREPORT_CAPACITY_16_DataSize:
		return "KREPORT_CAPACITY_16_DataSize"
	default:
		return fmt.Sprintf("KreportCapacity16(%d)", e)
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

type KsbcmodepagecachingDemand uint

const (
	// KSBCModePageCaching_DEMAND_READ_Mask: # Discussion
	KSBCModePageCaching_DEMAND_READ_Mask KsbcmodepagecachingDemand = 0
	// KSBCModePageCaching_DEMAND_WRITE_Mask: # Discussion
	KSBCModePageCaching_DEMAND_WRITE_Mask KsbcmodepagecachingDemand = 0
)

func (e KsbcmodepagecachingDemand) String() string {
	switch e {
	case KSBCModePageCaching_DEMAND_READ_Mask:
		return "KSBCModePageCaching_DEMAND_READ_Mask"
	default:
		return fmt.Sprintf("KsbcmodepagecachingDemand(%d)", e)
	}
}

type KsbcmodepagecachingDra uint

const (
	// KSBCModePageCaching_DRA_Bit: # Discussion
	KSBCModePageCaching_DRA_Bit KsbcmodepagecachingDra = 0
)

func (e KsbcmodepagecachingDra) String() string {
	switch e {
	case KSBCModePageCaching_DRA_Bit:
		return "KSBCModePageCaching_DRA_Bit"
	default:
		return fmt.Sprintf("KsbcmodepagecachingDra(%d)", e)
	}
}

type KsbcmodepageflexiblediskPin uint

const (
	// KSBCModePageFlexibleDisk_PIN_2_Mask: # Discussion
	KSBCModePageFlexibleDisk_PIN_2_Mask KsbcmodepageflexiblediskPin = 0
	// KSBCModePageFlexibleDisk_PIN_34_Mask: # Discussion
	KSBCModePageFlexibleDisk_PIN_34_Mask KsbcmodepageflexiblediskPin = 0
)

func (e KsbcmodepageflexiblediskPin) String() string {
	switch e {
	case KSBCModePageFlexibleDisk_PIN_2_Mask:
		return "KSBCModePageFlexibleDisk_PIN_2_Mask"
	default:
		return fmt.Sprintf("KsbcmodepageflexiblediskPin(%d)", e)
	}
}

type KsbcmodepageflexiblediskPin1 uint

const (
	// KSBCModePageFlexibleDisk_PIN_1_Mask: # Discussion
	KSBCModePageFlexibleDisk_PIN_1_Mask KsbcmodepageflexiblediskPin1 = 0
)

func (e KsbcmodepageflexiblediskPin1) String() string {
	switch e {
	case KSBCModePageFlexibleDisk_PIN_1_Mask:
		return "KSBCModePageFlexibleDisk_PIN_1_Mask"
	default:
		return fmt.Sprintf("KsbcmodepageflexiblediskPin1(%d)", e)
	}
}

type KsbcmodepageflexiblediskSpc uint

const (
	// KSBCModePageFlexibleDisk_SPC_Mask: # Discussion
	KSBCModePageFlexibleDisk_SPC_Mask KsbcmodepageflexiblediskSpc = 0
)

func (e KsbcmodepageflexiblediskSpc) String() string {
	switch e {
	case KSBCModePageFlexibleDisk_SPC_Mask:
		return "KSBCModePageFlexibleDisk_SPC_Mask"
	default:
		return fmt.Sprintf("KsbcmodepageflexiblediskSpc(%d)", e)
	}
}

type KsbcmodepagerigiddiskgeometryRpl uint

const (
	// KSBCModePageRigidDiskGeometry_RPL_Mask: # Discussion
	KSBCModePageRigidDiskGeometry_RPL_Mask KsbcmodepagerigiddiskgeometryRpl = 0
)

func (e KsbcmodepagerigiddiskgeometryRpl) String() string {
	switch e {
	case KSBCModePageRigidDiskGeometry_RPL_Mask:
		return "KSBCModePageRigidDiskGeometry_RPL_Mask"
	default:
		return fmt.Sprintf("KsbcmodepagerigiddiskgeometryRpl(%d)", e)
	}
}

type KscsiportStatus int

const (
	// KSCSIPort_StatusFailure: # Discussion
	KSCSIPort_StatusFailure KscsiportStatus = 0
	// KSCSIPort_StatusOffline: # Discussion
	KSCSIPort_StatusOffline KscsiportStatus = 0
	// KSCSIPort_StatusOnline: # Discussion
	KSCSIPort_StatusOnline KscsiportStatus = 0
)

func (e KscsiportStatus) String() string {
	switch e {
	case KSCSIPort_StatusFailure:
		return "KSCSIPort_StatusFailure"
	default:
		return fmt.Sprintf("KscsiportStatus(%d)", e)
	}
}

type KscsiserviceactionChange uint

const (
	KSCSIServiceAction_CHANGE_ALIASES KscsiserviceactionChange = 0
)

func (e KscsiserviceactionChange) String() string {
	switch e {
	case KSCSIServiceAction_CHANGE_ALIASES:
		return "KSCSIServiceAction_CHANGE_ALIASES"
	default:
		return fmt.Sprintf("KscsiserviceactionChange(%d)", e)
	}
}

type KscsiserviceactionGetLba uint

const (
	KSCSIServiceAction_GET_LBA_STATUS KscsiserviceactionGetLba = 0
)

func (e KscsiserviceactionGetLba) String() string {
	switch e {
	case KSCSIServiceAction_GET_LBA_STATUS:
		return "KSCSIServiceAction_GET_LBA_STATUS"
	default:
		return fmt.Sprintf("KscsiserviceactionGetLba(%d)", e)
	}
}

type KscsiserviceactionReport uint

const (
	KSCSIServiceAction_REPORT_ALIASES KscsiserviceactionReport = 0
)

func (e KscsiserviceactionReport) String() string {
	switch e {
	case KSCSIServiceAction_REPORT_ALIASES:
		return "KSCSIServiceAction_REPORT_ALIASES"
	default:
		return fmt.Sprintf("KscsiserviceactionReport(%d)", e)
	}
}

type KscsiserviceactionWriteLong uint

const (
	KSCSIServiceAction_WRITE_LONG_16 KscsiserviceactionWriteLong = 0
)

func (e KscsiserviceactionWriteLong) String() string {
	switch e {
	case KSCSIServiceAction_WRITE_LONG_16:
		return "KSCSIServiceAction_WRITE_LONG_16"
	default:
		return fmt.Sprintf("KscsiserviceactionWriteLong(%d)", e)
	}
}

type KsenseEom uint

const (
	// KSENSE_EOM_Mask: # Discussion
	KSENSE_EOM_Mask KsenseEom = 0
)

func (e KsenseEom) String() string {
	switch e {
	case KSENSE_EOM_Mask:
		return "KSENSE_EOM_Mask"
	default:
		return fmt.Sprintf("KsenseEom(%d)", e)
	}
}

type KsenseFilemark uint

const (
	// KSENSE_FILEMARK_Mask: # Discussion
	KSENSE_FILEMARK_Mask KsenseFilemark = 0
)

func (e KsenseFilemark) String() string {
	switch e {
	case KSENSE_FILEMARK_Mask:
		return "KSENSE_FILEMARK_Mask"
	default:
		return fmt.Sprintf("KsenseFilemark(%d)", e)
	}
}

type KsenseIli uint

const (
	// KSENSE_ILI_Mask: # Discussion
	KSENSE_ILI_Mask KsenseIli = 0
)

func (e KsenseIli) String() string {
	switch e {
	case KSENSE_ILI_Mask:
		return "KSENSE_ILI_Mask"
	default:
		return fmt.Sprintf("KsenseIli(%d)", e)
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

type KsenseNotData uint

const (
	// KSENSE_NOT_DATA_VALID: # Discussion
	KSENSE_NOT_DATA_VALID KsenseNotData = 0
)

func (e KsenseNotData) String() string {
	switch e {
	case KSENSE_NOT_DATA_VALID:
		return "KSENSE_NOT_DATA_VALID"
	default:
		return fmt.Sprintf("KsenseNotData(%d)", e)
	}
}

type KsenseResponseCodeCurrent uint

const (
	// KSENSE_RESPONSE_CODE_Current_Errors: # Discussion
	KSENSE_RESPONSE_CODE_Current_Errors KsenseResponseCodeCurrent = 0
)

func (e KsenseResponseCodeCurrent) String() string {
	switch e {
	case KSENSE_RESPONSE_CODE_Current_Errors:
		return "KSENSE_RESPONSE_CODE_Current_Errors"
	default:
		return fmt.Sprintf("KsenseResponseCodeCurrent(%d)", e)
	}
}

type KsescmdModeSelect uint

const (
	KSESCmd_MODE_SELECT_6 KsescmdModeSelect = 0
)

func (e KsescmdModeSelect) String() string {
	switch e {
	case KSESCmd_MODE_SELECT_6:
		return "KSESCmd_MODE_SELECT_6"
	default:
		return fmt.Sprintf("KsescmdModeSelect(%d)", e)
	}
}

type KsgccmdC uint

const (
	KSGCCmd_CHANGE_DEFINITION KsgccmdC = 0
	KSGCCmd_COPY_AND_VERIFY   KsgccmdC = 0
)

func (e KsgccmdC) String() string {
	switch e {
	case KSGCCmd_CHANGE_DEFINITION:
		return "KSGCCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KsgccmdC(%d)", e)
	}
}

type KspcproccmdChange uint

const (
	KSPCProcCmd_CHANGE_DEFINITION KspcproccmdChange = 0
)

func (e KspcproccmdChange) String() string {
	switch e {
	case KSPCProcCmd_CHANGE_DEFINITION:
		return "KSPCProcCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KspcproccmdChange(%d)", e)
	}
}

type KsscprintercmdChange uint

const (
	KSSCPrinterCmd_CHANGE_DEFINITION KsscprintercmdChange = 0
)

func (e KsscprintercmdChange) String() string {
	switch e {
	case KSSCPrinterCmd_CHANGE_DEFINITION:
		return "KSSCPrinterCmd_CHANGE_DEFINITION"
	default:
		return fmt.Sprintf("KsscprintercmdChange(%d)", e)
	}
}

type KstateDisabled uint

const (
	KState_Disabled_Flag KstateDisabled = 0
)

func (e KstateDisabled) String() string {
	switch e {
	case KState_Disabled_Flag:
		return "KState_Disabled_Flag"
	default:
		return fmt.Sprintf("KstateDisabled(%d)", e)
	}
}

type KusbSscompdesc uint

const (
	KUSB_SSCompDesc_Bulk_MaxStreams_Mask  KusbSscompdesc = 0
	KUSB_SSCompDesc_Bulk_MaxStreams_Shift KusbSscompdesc = 0
	KUSB_SSCompDesc_Isoc_Mult_Mask        KusbSscompdesc = 0
	KUSB_SSCompDesc_Isoc_Mult_Shift       KusbSscompdesc = 0
)

func (e KusbSscompdesc) String() string {
	switch e {
	case KUSB_SSCompDesc_Bulk_MaxStreams_Mask:
		return "KUSB_SSCompDesc_Bulk_MaxStreams_Mask"
	default:
		return fmt.Sprintf("KusbSscompdesc(%d)", e)
	}
}

type LatencyQosTier uint

const (
	LATENCY_QOS_TIER_0           LatencyQosTier = 0
	LATENCY_QOS_TIER_1           LatencyQosTier = 0
	LATENCY_QOS_TIER_2           LatencyQosTier = 0
	LATENCY_QOS_TIER_3           LatencyQosTier = 0
	LATENCY_QOS_TIER_4           LatencyQosTier = 0
	LATENCY_QOS_TIER_5           LatencyQosTier = 0
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

type LckSleep uint

const (
	LCK_SLEEP_DEFAULT LckSleep = 0
)

func (e LckSleep) String() string {
	switch e {
	case LCK_SLEEP_DEFAULT:
		return "LCK_SLEEP_DEFAULT"
	default:
		return fmt.Sprintf("LckSleep(%d)", e)
	}
}

type LckWake uint

const (
	LCK_WAKE_DEFAULT LckWake = 0
)

func (e LckWake) String() string {
	switch e {
	case LCK_WAKE_DEFAULT:
		return "LCK_WAKE_DEFAULT"
	default:
		return fmt.Sprintf("LckWake(%d)", e)
	}
}

type Libsptm uint

const (
	LIBSPTM_INVALID_ARG   Libsptm = 0
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

type MATA uint

const (
	MATAAltSDevCValid      MATA = 0
	MATACylinderHiValid    MATA = 0
	MATACylinderLoValid    MATA = 0
	MATADataValid          MATA = 0
	MATAErrFeaturesValid   MATA = 0
	MATAFlag48BitLBA       MATA = 0
	MATAFlagByteSwap       MATA = 0
	MATAFlagDMAQueued      MATA = 0
	MATAFlagIORead         MATA = 0
	MATAFlagIOWrite        MATA = 0
	MATAFlagImmediate      MATA = 0
	MATAFlagOverlapped     MATA = 0
	MATAFlagProtocolATAPI  MATA = 0
	MATAFlagQuiesce        MATA = 0
	MATAFlagTFAccess       MATA = 0
	MATAFlagTFAccessResult MATA = 0
	MATAFlagUseConfigSpeed MATA = 0
	MATAFlagUseDMA         MATA = 0
	MATAFlagUseNoIRQ       MATA = 0
	MATASDHValid           MATA = 0
	MATASectorCntValid     MATA = 0
	MATASectorNumValid     MATA = 0
	MATAStatusCmdValid     MATA = 0
)

func (e MATA) String() string {
	switch e {
	case MATAAltSDevCValid:
		return "MATAAltSDevCValid"
	default:
		return fmt.Sprintf("MATA(%d)", e)
	}
}

type MATAData uint

const (
	MATADataRequest MATAData = 0
)

func (e MATAData) String() string {
	switch e {
	case MATADataRequest:
		return "MATADataRequest"
	default:
		return fmt.Sprintf("MATAData(%d)", e)
	}
}

type MATADrive uint

const (
	MATADriveSelect MATADrive = 0
)

func (e MATADrive) String() string {
	switch e {
	case MATADriveSelect:
		return "MATADriveSelect"
	default:
		return fmt.Sprintf("MATADrive(%d)", e)
	}
}

type MBaseOffset uint

const (
	EightBitMode     MBaseOffset = 0
	FourBitMode      MBaseOffset = 0
	MBaseOffsetValue MBaseOffset = 0
	MBounds          MBaseOffset = 0
	MCmpCount        MBaseOffset = 0
	MCmpSize         MBaseOffset = 0
	MDevType         MBaseOffset = 0
	MHRes            MBaseOffset = 0
	MPageCnt         MBaseOffset = 0
	MPixelSize       MBaseOffset = 0
	MPixelType       MBaseOffset = 0
	MPlaneBytes      MBaseOffset = 0
	MRowBytes        MBaseOffset = 0
	MTable           MBaseOffset = 0
	MVRes            MBaseOffset = 0
	MVersion         MBaseOffset = 0
	MVertRefRate     MBaseOffset = 0
	MVidParams       MBaseOffset = 0
	OneBitMode       MBaseOffset = 0
	TwoBitMode       MBaseOffset = 0
)

func (e MBaseOffset) String() string {
	switch e {
	case EightBitMode:
		return "EightBitMode"
	default:
		return fmt.Sprintf("MBaseOffset(%d)", e)
	}
}

type MachAssert uint

const (
	MACH_ASSERT_3P MachAssert = 0
)

func (e MachAssert) String() string {
	switch e {
	case MACH_ASSERT_3P:
		return "MACH_ASSERT_3P"
	default:
		return fmt.Sprintf("MachAssert(%d)", e)
	}
}

type MachVmRange uint

const (
	MACH_VM_RANGE_DATA MachVmRange = 0
	MACH_VM_RANGE_NONE MachVmRange = 0
)

func (e MachVmRange) String() string {
	switch e {
	case MACH_VM_RANGE_DATA:
		return "MACH_VM_RANGE_DATA"
	default:
		return fmt.Sprintf("MachVmRange(%d)", e)
	}
}

type MachVmRangeFlavor uint

const (
	MACH_VM_RANGE_FLAVOR_INVALID MachVmRangeFlavor = 0
)

func (e MachVmRangeFlavor) String() string {
	switch e {
	case MACH_VM_RANGE_FLAVOR_INVALID:
		return "MACH_VM_RANGE_FLAVOR_INVALID"
	default:
		return fmt.Sprintf("MachVmRangeFlavor(%d)", e)
	}
}

type MacosPanicHeaderFlag uint

const (
	MACOS_PANIC_HEADER_FLAG_COMPANION_PROC_INITIATED_PANIC                      MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_COREDUMP_COMPLETE                                   MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_COREDUMP_FAILED                                     MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_COREFILE_UNLINKED                                   MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_ENCRYPTED_COREDUMP_SKIPPED                          MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_INCOHERENT_PANICLOG                                 MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_INTEGRATED_COPROC_INITIATED_PANIC                   MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_KERNEL_COREDUMP_SKIPPED_EXCLUDE_REGIONS_UNAVAILABLE MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_NESTED_PANIC                                        MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_STACKSHOT_DATA_COMPRESSED                           MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_STACKSHOT_FAILED_COMPRESS                           MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_STACKSHOT_FAILED_DEBUGGERSYNC                       MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_STACKSHOT_FAILED_ERROR                              MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_STACKSHOT_FAILED_INCOMPLETE                         MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_STACKSHOT_FAILED_NESTED                             MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_STACKSHOT_KERNEL_ONLY                               MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_STACKSHOT_SUCCEEDED                                 MacosPanicHeaderFlag = 0
	MACOS_PANIC_HEADER_FLAG_USERSPACE_INITIATED_PANIC                           MacosPanicHeaderFlag = 0
)

func (e MacosPanicHeaderFlag) String() string {
	switch e {
	case MACOS_PANIC_HEADER_FLAG_COMPANION_PROC_INITIATED_PANIC:
		return "MACOS_PANIC_HEADER_FLAG_COMPANION_PROC_INITIATED_PANIC"
	default:
		return fmt.Sprintf("MacosPanicHeaderFlag(%d)", e)
	}
}

type Mbuf uint

const (
	// MBUF_BCAST: # Discussion
	MBUF_BCAST Mbuf = 0
	// MBUF_DONTWAIT: # Discussion
	MBUF_DONTWAIT Mbuf = 0
)

func (e Mbuf) String() string {
	switch e {
	case MBUF_BCAST:
		return "MBUF_BCAST"
	default:
		return fmt.Sprintf("Mbuf(%d)", e)
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

type MbufCsumReq uint

const (
	// MBUF_CSUM_REQ_IP: # Discussion
	MBUF_CSUM_REQ_IP MbufCsumReq = 0
	// MBUF_CSUM_REQ_TCP: # Discussion
	MBUF_CSUM_REQ_TCP MbufCsumReq = 0
	// MBUF_CSUM_REQ_TCPIPV6: # Discussion
	MBUF_CSUM_REQ_TCPIPV6 MbufCsumReq = 0
	// MBUF_CSUM_REQ_UDP: # Discussion
	MBUF_CSUM_REQ_UDP MbufCsumReq = 0
	// MBUF_CSUM_REQ_UDPIPV6: # Discussion
	MBUF_CSUM_REQ_UDPIPV6 MbufCsumReq = 0
)

func (e MbufCsumReq) String() string {
	switch e {
	case MBUF_CSUM_REQ_IP:
		return "MBUF_CSUM_REQ_IP"
	default:
		return fmt.Sprintf("MbufCsumReq(%d)", e)
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

type MbufTso uint

const (
	MBUF_TSO_IPV4 MbufTso = 0
	MBUF_TSO_IPV6 MbufTso = 0
)

func (e MbufTso) String() string {
	switch e {
	case MBUF_TSO_IPV4:
		return "MBUF_TSO_IPV4"
	default:
		return fmt.Sprintf("MbufTso(%d)", e)
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

type MccIsMulti uint

const (
	MCC_IS_MULTI_BIT MccIsMulti = 0
)

func (e MccIsMulti) String() string {
	switch e {
	case MCC_IS_MULTI_BIT:
		return "MCC_IS_MULTI_BIT"
	default:
		return fmt.Sprintf("MccIsMulti(%d)", e)
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

type Nf uint

const (
	NFATTRDIR   Nf = 0
	NFBLK       Nf = 0
	NFCHR       Nf = 0
	NFDIR       Nf = 0
	NFFIFO      Nf = 0
	NFLNK       Nf = 0
	NFNAMEDATTR Nf = 0
	NFNON       Nf = 0
	NFREG       Nf = 0
	NFSOCK      Nf = 0
)

func (e Nf) String() string {
	switch e {
	case NFATTRDIR:
		return "NFATTRDIR"
	default:
		return fmt.Sprintf("Nf(%d)", e)
	}
}

type NfsAes128CtsHmacSha1 uint

const (
	NFS_AES128_CTS_HMAC_SHA1_96 NfsAes128CtsHmacSha1 = 0
)

func (e NfsAes128CtsHmacSha1) String() string {
	switch e {
	case NFS_AES128_CTS_HMAC_SHA1_96:
		return "NFS_AES128_CTS_HMAC_SHA1_96"
	default:
		return fmt.Sprintf("NfsAes128CtsHmacSha1(%d)", e)
	}
}

type No uint

const (
	// NoErr: No error
	NoErr No = 0
)

func (e No) String() string {
	switch e {
	case NoErr:
		return "NoErr"
	default:
		return fmt.Sprintf("No(%d)", e)
	}
}

type NrLockedErr int

const (
	GestaltUndefSelectorErr NrLockedErr = 0
	NrLockedErrValue        NrLockedErr = 0
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
	NX_BigEndian        Nx = 0
	NX_LeftButton       Nx = 0
	NX_LittleEndian     Nx = 0
	NX_OneButton        Nx = 0
	NX_RightButton      Nx = 0
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

type Op uint

const (
	OP_AUTHORIZE Op = 0
)

func (e Op) String() string {
	switch e {
	case OP_AUTHORIZE:
		return "OP_AUTHORIZE"
	default:
		return fmt.Sprintf("Op(%d)", e)
	}
}

type Os uint

const (
	OSBigEndian    Os = 0
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

type OsLogCoprocRegisterHarvestFs uint

const (
	Os_log_coproc_register_harvest_fs_ftab OsLogCoprocRegisterHarvestFs = 0
)

func (e OsLogCoprocRegisterHarvestFs) String() string {
	switch e {
	case Os_log_coproc_register_harvest_fs_ftab:
		return "Os_log_coproc_register_harvest_fs_ftab"
	default:
		return fmt.Sprintf("OsLogCoprocRegisterHarvestFs(%d)", e)
	}
}

type OsLogType uint

const (
	// OS_LOG_TYPE_DEBUG: Debug-level messages are only captured in memory when debug logging is enabled through a configuration change.
	OS_LOG_TYPE_DEBUG OsLogType = 0
	// OS_LOG_TYPE_DEFAULT: Default-level messages are initially stored in memory buffers.
	OS_LOG_TYPE_DEFAULT OsLogType = 0
	// OS_LOG_TYPE_ERROR: Error-level messages are always saved in the data store.
	OS_LOG_TYPE_ERROR OsLogType = 0
	// OS_LOG_TYPE_INFO: Info-level messages are initially stored in memory buffers.
	OS_LOG_TYPE_INFO OsLogType = 0
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
	OUTPUT_UNDEFINED             Output = 0
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
	INPUT_NULL  OutputNull = 0
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
	BadCksmErr           ParamErr = 0
	BadUnitErr           ParamErr = 0
	ClosErr              ParamErr = 0
	ControlErr           ParamErr = 0
	CorErr               ParamErr = 0
	DInstErr             ParamErr = 0
	DRemovErr            ParamErr = 0
	NoHardwareErr        ParamErr = 0
	NotEnoughHardwareErr ParamErr = 0
	OpenErr              ParamErr = 0
	ParamErrValue        ParamErr = 0
	QErr                 ParamErr = 0
	ReadErr              ParamErr = 0
	SeNoDB               ParamErr = 0
	SlpTypeErr           ParamErr = 0
	StatusErr            ParamErr = 0
	UnimpErr             ParamErr = 0
	UnitEmptyErr         ParamErr = 0
	UserCanceledErr      ParamErr = 0
	VTypErr              ParamErr = 0
	WritErr              ParamErr = 0
)

func (e ParamErr) String() string {
	switch e {
	case BadCksmErr:
		return "BadCksmErr"
	default:
		return fmt.Sprintf("ParamErr(%d)", e)
	}
}

type PriorityQueueEntry uint

const (
	PRIORITY_QUEUE_ENTRY_NONE      PriorityQueueEntry = 0
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
	PROCESSOR_GENERAL   Processor = 0
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

type RGB uint

const (
	RGBDirect RGB = 0
)

func (e RGB) String() string {
	switch e {
	case RGBDirect:
		return "RGBDirect"
	default:
		return fmt.Sprintf("RGB(%d)", e)
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

type SecondVid uint

const (
	SecondVidMode SecondVid = 0
)

func (e SecondVid) String() string {
	switch e {
	case SecondVidMode:
		return "SecondVidMode"
	default:
		return fmt.Sprintf("SecondVid(%d)", e)
	}
}

type Sflt uint

const (
	// SFLT_EXTENDED: # Discussion
	SFLT_EXTENDED Sflt = 0
	// SFLT_EXTENDED_REGISTRY: # Discussion
	SFLT_EXTENDED_REGISTRY Sflt = 0
	// SFLT_GLOBAL: # Discussion
	SFLT_GLOBAL Sflt = 0
	// SFLT_PROG: # Discussion
	SFLT_PROG Sflt = 0
)

func (e Sflt) String() string {
	switch e {
	case SFLT_EXTENDED:
		return "SFLT_EXTENDED"
	default:
		return fmt.Sprintf("Sflt(%d)", e)
	}
}

type SoTrackerAction uint

const (
	SO_TRACKER_ACTION_ADD         SoTrackerAction = 0
	SO_TRACKER_ACTION_DUMP_ALL    SoTrackerAction = 0
	SO_TRACKER_ACTION_DUMP_BY_APP SoTrackerAction = 0
	SO_TRACKER_ACTION_DUMP_MAX    SoTrackerAction = 0
	SO_TRACKER_ACTION_INVALID     SoTrackerAction = 0
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
	SO_TRACKER_ATTRIBUTE_ADDRESS        SoTrackerAttributeA = 0
	SO_TRACKER_ATTRIBUTE_ADDRESS_FAMILY SoTrackerAttributeA = 0
	SO_TRACKER_ATTRIBUTE_APP_UUID       SoTrackerAttributeA = 0
)

func (e SoTrackerAttributeA) String() string {
	switch e {
	case SO_TRACKER_ATTRIBUTE_ADDRESS:
		return "SO_TRACKER_ATTRIBUTE_ADDRESS"
	default:
		return fmt.Sprintf("SoTrackerAttributeA(%d)", e)
	}
}

type SockDataFiltFlag uint

const (
	// Sock_data_filt_flag_oob: # Discussion
	Sock_data_filt_flag_oob SockDataFiltFlag = 0
)

func (e SockDataFiltFlag) String() string {
	switch e {
	case Sock_data_filt_flag_oob:
		return "Sock_data_filt_flag_oob"
	default:
		return fmt.Sprintf("SockDataFiltFlag(%d)", e)
	}
}

type SockEvt uint

const (
	// Sock_evt_bound: # Discussion
	Sock_evt_bound SockEvt = 0
	// Sock_evt_cantrecvmore: # Discussion
	Sock_evt_cantrecvmore SockEvt = 0
	// Sock_evt_cantsendmore: # Discussion
	Sock_evt_cantsendmore SockEvt = 0
	// Sock_evt_closing: # Discussion
	Sock_evt_closing SockEvt = 0
	// Sock_evt_connected: # Discussion
	Sock_evt_connected  SockEvt = 0
	Sock_evt_connecting SockEvt = 0
	// Sock_evt_disconnected: # Discussion
	Sock_evt_disconnected  SockEvt = 0
	Sock_evt_disconnecting SockEvt = 0
	// Sock_evt_flush_read: # Discussion
	Sock_evt_flush_read SockEvt = 0
	// Sock_evt_shutdown: # Discussion
	Sock_evt_shutdown SockEvt = 0
)

func (e SockEvt) String() string {
	switch e {
	case Sock_evt_bound:
		return "Sock_evt_bound"
	default:
		return fmt.Sprintf("SockEvt(%d)", e)
	}
}

type Sockopt uint

const (
	Sockopt_get Sockopt = 0
	Sockopt_set Sockopt = 0
)

func (e Sockopt) String() string {
	switch e {
	case Sockopt_get:
		return "Sockopt_get"
	default:
		return fmt.Sprintf("Sockopt(%d)", e)
	}
}

type Sptm uint

const (
	SPTM_N_VECTOR_TYPES Sptm = 0
	SPTM_VECTOR_FIQ     Sptm = 0
)

func (e Sptm) String() string {
	switch e {
	case SPTM_N_VECTOR_TYPES:
		return "SPTM_N_VECTOR_TYPES"
	default:
		return fmt.Sprintf("Sptm(%d)", e)
	}
}

type SptmN uint

const (
	SPTM_N_TRACES SptmN = 0
)

func (e SptmN) String() string {
	switch e {
	case SPTM_N_TRACES:
		return "SPTM_N_TRACES"
	default:
		return fmt.Sprintf("SptmN(%d)", e)
	}
}

type SptmNPublic uint

const (
	SPTM_N_PUBLIC_VIOLATIONS SptmNPublic = 0
)

func (e SptmNPublic) String() string {
	switch e {
	case SPTM_N_PUBLIC_VIOLATIONS:
		return "SPTM_N_PUBLIC_VIOLATIONS"
	default:
		return fmt.Sprintf("SptmNPublic(%d)", e)
	}
}

type SptmRefcount uint

const (
	SPTM_REFCOUNT_IOMMU SptmRefcount = 0
	SPTM_REFCOUNT_NONE  SptmRefcount = 0
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
	STACKSHOT_ACTIVE_KERNEL_THREADS_ONLY       Stackshot = 0
	STACKSHOT_ASID                             Stackshot = 0
	STACKSHOT_COLLECT_DELTA_SNAPSHOT           Stackshot = 0
	STACKSHOT_COLLECT_SHAREDCACHE_LAYOUT       Stackshot = 0
	STACKSHOT_DISABLE_LATENCY_INFO             Stackshot = 0
	STACKSHOT_DO_COMPRESS                      Stackshot = 0
	STACKSHOT_ENABLE_BT_FAULTING               Stackshot = 0
	STACKSHOT_ENABLE_UUID_FAULTING             Stackshot = 0
	STACKSHOT_EXCLAVES                         Stackshot = 0
	STACKSHOT_FROM_PANIC                       Stackshot = 0
	STACKSHOT_GET_BOOT_PROFILE                 Stackshot = 0
	STACKSHOT_GET_DQ                           Stackshot = 0
	STACKSHOT_GET_GLOBAL_MEM_STATS             Stackshot = 0
	STACKSHOT_INCLUDE_DRIVER_THREADS_IN_KERNEL Stackshot = 0
	STACKSHOT_INSTRS_CYCLES                    Stackshot = 0
	STACKSHOT_KCDATA_FORMAT                    Stackshot = 0
	STACKSHOT_NO_IO_STATS                      Stackshot = 0
	STACKSHOT_PAGE_TABLES                      Stackshot = 0
	STACKSHOT_RETRIEVE_EXISTING_BUFFER         Stackshot = 0
	STACKSHOT_SAVE_DYLD_COMPACTINFO            Stackshot = 0
	STACKSHOT_SAVE_IMP_DONATION_PIDS           Stackshot = 0
	STACKSHOT_SAVE_IN_KERNEL_BUFFER            Stackshot = 0
	STACKSHOT_SAVE_JETSAM_COALITIONS           Stackshot = 0
	STACKSHOT_SAVE_KEXT_LOADINFO               Stackshot = 0
	STACKSHOT_SAVE_LOADINFO                    Stackshot = 0
	STACKSHOT_SKIP_EXCLAVES                    Stackshot = 0
	STACKSHOT_THREAD_GROUP                     Stackshot = 0
	STACKSHOT_THREAD_WAITINFO                  Stackshot = 0
	STACKSHOT_TRYLOCK                          Stackshot = 0
)

func (e Stackshot) String() string {
	switch e {
	case STACKSHOT_ACTIVE_KERNEL_THREADS_ONLY:
		return "STACKSHOT_ACTIVE_KERNEL_THREADS_ONLY"
	default:
		return fmt.Sprintf("Stackshot(%d)", e)
	}
}

type StackshotGetKernel uint

const (
	STACKSHOT_GET_KERNEL_MICROSTACKSHOT StackshotGetKernel = 0
)

func (e StackshotGetKernel) String() string {
	switch e {
	case STACKSHOT_GET_KERNEL_MICROSTACKSHOT:
		return "STACKSHOT_GET_KERNEL_MICROSTACKSHOT"
	default:
		return fmt.Sprintf("StackshotGetKernel(%d)", e)
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

type Task uint

const (
	TASK_BACKGROUND_APPLICATION Task = 0
	TASK_CONTROL_APPLICATION    Task = 0
	TASK_DARWINBG_APPLICATION   Task = 0
	TASK_DEFAULT_APPLICATION    Task = 0
	TASK_FOREGROUND_APPLICATION Task = 0
	TASK_GRAPHICS_SERVER        Task = 0
	TASK_NONUI_APPLICATION      Task = 0
	TASK_RENICED                Task = 0
	TASK_THROTTLE_APPLICATION   Task = 0
	TASK_UNSPECIFIED            Task = 0
)

func (e Task) String() string {
	switch e {
	case TASK_BACKGROUND_APPLICATION:
		return "TASK_BACKGROUND_APPLICATION"
	default:
		return fmt.Sprintf("Task(%d)", e)
	}
}

type TaskInspectBasic uint

const (
	TASK_INSPECT_BASIC_COUNTS TaskInspectBasic = 0
)

func (e TaskInspectBasic) String() string {
	switch e {
	case TASK_INSPECT_BASIC_COUNTS:
		return "TASK_INSPECT_BASIC_COUNTS"
	default:
		return fmt.Sprintf("TaskInspectBasic(%d)", e)
	}
}

type TaskSnapshotFlags uint

const (
	KFrozen                                     TaskSnapshotFlags = 0
	KPidSuspended                               TaskSnapshotFlags = 0
	KTaskAllowIdleExit                          TaskSnapshotFlags = 0
	KTaskDarwinBG                               TaskSnapshotFlags = 0
	KTaskDyldCompactInfoFaultedIn               TaskSnapshotFlags = 0
	KTaskDyldCompactInfoMissing                 TaskSnapshotFlags = 0
	KTaskDyldCompactInfoNone                    TaskSnapshotFlags = 0
	KTaskDyldCompactInfoTooBig                  TaskSnapshotFlags = 0
	KTaskDyldCompactInfoTriedFault              TaskSnapshotFlags = 0
	KTaskExtDarwinBG                            TaskSnapshotFlags = 0
	KTaskIsBoosted                              TaskSnapshotFlags = 0
	KTaskIsDirty                                TaskSnapshotFlags = 0
	KTaskIsDirtyTracked                         TaskSnapshotFlags = 0
	KTaskIsForeground                           TaskSnapshotFlags = 0
	KTaskIsImpDonor                             TaskSnapshotFlags = 0
	KTaskIsLiveImpDonor                         TaskSnapshotFlags = 0
	KTaskIsSuppressed                           TaskSnapshotFlags = 0
	KTaskIsTimerThrottled                       TaskSnapshotFlags = 0
	KTaskIsTranslated                           TaskSnapshotFlags = 0
	KTaskRsrcFlagged                            TaskSnapshotFlags = 0
	KTaskSharedRegionInfoUnavailable            TaskSnapshotFlags = 0
	KTaskSharedRegionNone                       TaskSnapshotFlags = 0
	KTaskSharedRegionOther                      TaskSnapshotFlags = 0
	KTaskSharedRegionSystem                     TaskSnapshotFlags = 0
	KTaskTALEngaged                             TaskSnapshotFlags = 0
	KTaskUUIDInfoFaultedIn                      TaskSnapshotFlags = 0
	KTaskUUIDInfoMissing                        TaskSnapshotFlags = 0
	KTaskUUIDInfoTriedFault                     TaskSnapshotFlags = 0
	KTaskVisNonvisible                          TaskSnapshotFlags = 0
	KTaskVisVisible                             TaskSnapshotFlags = 0
	KTaskWqExceededActiveConstrainedThreadLimit TaskSnapshotFlags = 0
	KTaskWqExceededConstrainedThreadLimit       TaskSnapshotFlags = 0
	KTaskWqExceededCooperativeThreadLimit       TaskSnapshotFlags = 0
	KTaskWqExceededTotalThreadLimit             TaskSnapshotFlags = 0
	KTaskWqFlagsAvailable                       TaskSnapshotFlags = 0
	KTerminatedSnapshot                         TaskSnapshotFlags = 0
)

func (e TaskSnapshotFlags) String() string {
	switch e {
	case KFrozen:
		return "KFrozen"
	default:
		return fmt.Sprintf("TaskSnapshotFlags(%d)", e)
	}
}

type TcpConnectionClientAccurateEcn uint

const (
	Tcp_connection_client_accurate_ecn_feature_enabled                           TcpConnectionClientAccurateEcn = 0
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
	Tcp_connection_server_accurate_ecn_requested                                  TcpConnectionServerAccurateEcn = 0
)

func (e TcpConnectionServerAccurateEcn) String() string {
	switch e {
	case Tcp_connection_server_accurate_ecn_negotiation_success_ect_bleaching_detected:
		return "Tcp_connection_server_accurate_ecn_negotiation_success_ect_bleaching_detected"
	default:
		return fmt.Sprintf("TcpConnectionServerAccurateEcn(%d)", e)
	}
}

type TelemetryNotice uint

const (
	TELEMETRY_NOTICE_BASE                  TelemetryNotice = 0
	TELEMETRY_NOTICE_KERNEL_MICROSTACKSHOT TelemetryNotice = 0
)

func (e TelemetryNotice) String() string {
	switch e {
	case TELEMETRY_NOTICE_BASE:
		return "TELEMETRY_NOTICE_BASE"
	default:
		return fmt.Sprintf("TelemetryNotice(%d)", e)
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
	TELEPHONY_TELEPHONE       Telephony = 0
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
	KNotTerminated    TerminationState = 0
	KTerminated       TerminationState = 0
)

func (e TerminationState) String() string {
	switch e {
	case KNeedsTermination:
		return "KNeedsTermination"
	default:
		return fmt.Sprintf("TerminationState(%d)", e)
	}
}

type ThreadCallOptions uint

const (
	THREAD_CALL_OPTIONS_ONCE ThreadCallOptions = 0
)

func (e ThreadCallOptions) String() string {
	switch e {
	case THREAD_CALL_OPTIONS_ONCE:
		return "THREAD_CALL_OPTIONS_ONCE"
	default:
		return fmt.Sprintf("ThreadCallOptions(%d)", e)
	}
}

type ThreadCallPriority uint

const (
	// THREAD_CALL_PRIORITY_HIGH: # Discussion
	THREAD_CALL_PRIORITY_HIGH ThreadCallPriority = 0
	// THREAD_CALL_PRIORITY_KERNEL: # Discussion
	THREAD_CALL_PRIORITY_KERNEL      ThreadCallPriority = 0
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

type ThreadSnapshotFlags uint

const (
	KGlobalForcedIdle       ThreadSnapshotFlags = 0
	KHasDispatchSerial      ThreadSnapshotFlags = 0
	KStacksPCOnly           ThreadSnapshotFlags = 0
	KThreadDarwinBG         ThreadSnapshotFlags = 0
	KThreadFaultedBT        ThreadSnapshotFlags = 0
	KThreadIOPassive        ThreadSnapshotFlags = 0
	KThreadIdleWorker       ThreadSnapshotFlags = 0
	KThreadMain             ThreadSnapshotFlags = 0
	KThreadOnCore           ThreadSnapshotFlags = 0
	KThreadSuspended        ThreadSnapshotFlags = 0
	KThreadTriedFaultBT     ThreadSnapshotFlags = 0
	KThreadTruncKernBT      ThreadSnapshotFlags = 0
	KThreadTruncUserAsyncBT ThreadSnapshotFlags = 0
	KThreadTruncUserBT      ThreadSnapshotFlags = 0
	KThreadTruncatedBT      ThreadSnapshotFlags = 0
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
	THROUGHPUT_QOS_TIER_0           ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_1           ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_2           ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_3           ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_4           ThroughputQosTier = 0
	THROUGHPUT_QOS_TIER_5           ThroughputQosTier = 0
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

type ThscCpiPerPerf uint

const (
	THSC_CPI_PER_PERF_LEVEL ThscCpiPerPerf = 0
)

func (e ThscCpiPerPerf) String() string {
	switch e {
	case THSC_CPI_PER_PERF_LEVEL:
		return "THSC_CPI_PER_PERF_LEVEL"
	default:
		return fmt.Sprintf("ThscCpiPerPerf(%d)", e)
	}
}

type Timing uint

const (
	TimingApple12             Timing = 0
	TimingApple12x            Timing = 0
	TimingApple13             Timing = 0
	TimingApple13x            Timing = 0
	TimingApple15             Timing = 0
	TimingApple15x            Timing = 0
	TimingApple16             Timing = 0
	TimingApple19             Timing = 0
	TimingApple1Ka            Timing = 0
	TimingApple1Kb            Timing = 0
	TimingApple21             Timing = 0
	TimingAppleSVGA           Timing = 0
	TimingAppleVGA            Timing = 0
	TimingSony_1900x1200_74hz Timing = 0
	TimingSony_1900x1200_76hz Timing = 0
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

type Uio uint

const (
	UIO_READ       Uio = 0
	UIO_SYSSPACE   Uio = 0
	UIO_SYSSPACE32 Uio = 0
	UIO_WRITE      Uio = 0
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
	UIOF_SYSSPACE    Uiof = 0
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
	VDSP_IIRMonoLeft  VdspIir = 0
	VDSP_IIRMonoRight VdspIir = 0
	VDSP_IIRStereo    VdspIir = 0
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
	VFS_VM_ROLE       Vfs = 0
)

func (e Vfs) String() string {
	switch e {
	case VFS_RECOVERY_ROLE:
		return "VFS_RECOVERY_ROLE"
	default:
		return fmt.Sprintf("Vfs(%d)", e)
	}
}

type VfsRename uint

const (
	VFS_RENAME_EXCL VfsRename = 0
)

func (e VfsRename) String() string {
	switch e {
	case VFS_RENAME_EXCL:
		return "VFS_RENAME_EXCL"
	default:
		return fmt.Sprintf("VfsRename(%d)", e)
	}
}

type VnodeVerifyContext uint

const (
	VNODE_VERIFY_CONTEXT_ALLOC VnodeVerifyContext = 0
	VNODE_VERIFY_CONTEXT_FREE  VnodeVerifyContext = 0
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
	VBAD  Vtype = 0
	VBLK  Vtype = 0
	VCHR  Vtype = 0
	VCPLX Vtype = 0
	VDIR  Vtype = 0
	VLNK  Vtype = 0
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

const (
	XDRBUF_BUFFER uint = 0
	XDRBUF_NONE   uint = 0
)
