// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor

import (
	"fmt"
)

type HVAPICCtrl uint

const (
	HVAPICCtrlDefault   HVAPICCtrl = 0
	HVAPICCtrlEoiIcrTpr HVAPICCtrl = 1
	HVAPICCtrlGuestIdle HVAPICCtrl = 2
	HVAPICCtrlIOAPICEoi HVAPICCtrl = 8
	HVAPICCtrlNoTimer   HVAPICCtrl = 4
)

func (e HVAPICCtrl) String() string {
	switch e {
	case HVAPICCtrlDefault:
		return "HVAPICCtrlDefault"
	case HVAPICCtrlEoiIcrTpr:
		return "HVAPICCtrlEoiIcrTpr"
	case HVAPICCtrlGuestIdle:
		return "HVAPICCtrlGuestIdle"
	case HVAPICCtrlIOAPICEoi:
		return "HVAPICCtrlIOAPICEoi"
	case HVAPICCtrlNoTimer:
		return "HVAPICCtrlNoTimer"
	default:
		return fmt.Sprintf("HVAPICCtrl(%d)", e)
	}
}

type HVAPICIntrTrigger uint

const (
	HVAPICEdgeTrigger     HVAPICIntrTrigger = 0
	HVAPICEdgeTriggerAeoi HVAPICIntrTrigger = 1
	HVAPICLevelTrigger    HVAPICIntrTrigger = 2
)

func (e HVAPICIntrTrigger) String() string {
	switch e {
	case HVAPICEdgeTrigger:
		return "HVAPICEdgeTrigger"
	case HVAPICEdgeTriggerAeoi:
		return "HVAPICEdgeTriggerAeoi"
	case HVAPICLevelTrigger:
		return "HVAPICLevelTrigger"
	default:
		return fmt.Sprintf("HVAPICIntrTrigger(%d)", e)
	}
}

type HVAPICLvtFlavor uint

const (
	HVAPICLvtFlavorTimer HVAPICLvtFlavor = 1
)

func (e HVAPICLvtFlavor) String() string {
	switch e {
	case HVAPICLvtFlavorTimer:
		return "HVAPICLvtFlavorTimer"
	default:
		return fmt.Sprintf("HVAPICLvtFlavor(%d)", e)
	}
}

type HVCacheType uint

const (
	// HVCacheTypeData: The value that describes a cached data value.
	HVCacheTypeData HVCacheType = 0
	// HVCacheTypeInstruction: The value that describes a cached instuction value.
	HVCacheTypeInstruction HVCacheType = 1
)

func (e HVCacheType) String() string {
	switch e {
	case HVCacheTypeData:
		return "HVCacheTypeData"
	case HVCacheTypeInstruction:
		return "HVCacheTypeInstruction"
	default:
		return fmt.Sprintf("HVCacheType(%d)", e)
	}
}

type HVExitReason uint

const (
	// HVExitReasonCanceled: The value that identifies exits requested by exit handler on the host.
	HVExitReasonCanceled HVExitReason = 0
	// HVExitReasonException: The value that identifies traps caused by the guest operations.
	HVExitReasonException HVExitReason = 1
	// HVExitReasonUnknown: The value that identifies unexpected exits.
	HVExitReasonUnknown HVExitReason = 3
	// HVExitReasonVtimerActivated: The value that identifies when the virtual timer enters the pending state.
	HVExitReasonVtimerActivated HVExitReason = 2
)

func (e HVExitReason) String() string {
	switch e {
	case HVExitReasonCanceled:
		return "HVExitReasonCanceled"
	case HVExitReasonException:
		return "HVExitReasonException"
	case HVExitReasonUnknown:
		return "HVExitReasonUnknown"
	case HVExitReasonVtimerActivated:
		return "HVExitReasonVtimerActivated"
	default:
		return fmt.Sprintf("HVExitReason(%d)", e)
	}
}

type HVFeatureReg uint

const (
	// HVFeatureRegClidrEl1: The value that describes Cache Level ID Register, EL1.
	HVFeatureRegClidrEl1 HVFeatureReg = 10
	// HVFeatureRegCtrEl0: The value that describes Cache Type Register, EL0.
	HVFeatureRegCtrEl0 HVFeatureReg = 9
	// HVFeatureRegDczidEl0: The value that describes Data Cache Zero ID Register, EL0.
	HVFeatureRegDczidEl0 HVFeatureReg = 11
	// HVFeatureRegIDAa64dfr0El1: The value that identifies debug feature register 0, EL1 (DFR0_EL1).
	HVFeatureRegIDAa64dfr0El1 HVFeatureReg = 0
	// HVFeatureRegIDAa64dfr1El1: The value that identifies debug feature register 1, EL1 (DFR1_EL1).
	HVFeatureRegIDAa64dfr1El1 HVFeatureReg = 1
	// HVFeatureRegIDAa64isar0El1: The value that identifies instruction set attribute register 0, EL1 (ISAR0_EL1).
	HVFeatureRegIDAa64isar0El1 HVFeatureReg = 2
	// HVFeatureRegIDAa64isar1El1: The value that identifies instruction set attribute register 1, EL1 (ISAR_EL1).
	HVFeatureRegIDAa64isar1El1 HVFeatureReg = 3
	// HVFeatureRegIDAa64mmfr0El1: The value that identifies memory model feature register 0, EL1(MMFR0_EL1).
	HVFeatureRegIDAa64mmfr0El1 HVFeatureReg = 4
	// HVFeatureRegIDAa64mmfr1El1: The value that identifies memory model feature register 1, EL1 (MMFR1_EL1).
	HVFeatureRegIDAa64mmfr1El1 HVFeatureReg = 5
	// HVFeatureRegIDAa64mmfr2El1: The value that identifies memory model feature register 2, EL1 (MMFR2_EL1).
	HVFeatureRegIDAa64mmfr2El1 HVFeatureReg = 6
	// HVFeatureRegIDAa64pfr0El1: The value that identifies processor feature register 0, EL1 (PFR0_EL1).
	HVFeatureRegIDAa64pfr0El1 HVFeatureReg = 7
	// HVFeatureRegIDAa64pfr1El1: The value that identifies processor feature register 1, EL1 (PFR1_EL1).
	HVFeatureRegIDAa64pfr1El1 HVFeatureReg = 8
	// HVFeatureRegIDAa64smfr0El1: The value that describes Scalable Matrix Extension (SME) Feature ID Register 0.
	HVFeatureRegIDAa64smfr0El1 HVFeatureReg = 12
	// HVFeatureRegIDAa64zfr0El1: The value that describes Scalable Vector Extension instruction (SVE) Feature ID register 0.
	HVFeatureRegIDAa64zfr0El1 HVFeatureReg = 13
)

func (e HVFeatureReg) String() string {
	switch e {
	case HVFeatureRegClidrEl1:
		return "HVFeatureRegClidrEl1"
	case HVFeatureRegCtrEl0:
		return "HVFeatureRegCtrEl0"
	case HVFeatureRegDczidEl0:
		return "HVFeatureRegDczidEl0"
	case HVFeatureRegIDAa64dfr0El1:
		return "HVFeatureRegIDAa64dfr0El1"
	case HVFeatureRegIDAa64dfr1El1:
		return "HVFeatureRegIDAa64dfr1El1"
	case HVFeatureRegIDAa64isar0El1:
		return "HVFeatureRegIDAa64isar0El1"
	case HVFeatureRegIDAa64isar1El1:
		return "HVFeatureRegIDAa64isar1El1"
	case HVFeatureRegIDAa64mmfr0El1:
		return "HVFeatureRegIDAa64mmfr0El1"
	case HVFeatureRegIDAa64mmfr1El1:
		return "HVFeatureRegIDAa64mmfr1El1"
	case HVFeatureRegIDAa64mmfr2El1:
		return "HVFeatureRegIDAa64mmfr2El1"
	case HVFeatureRegIDAa64pfr0El1:
		return "HVFeatureRegIDAa64pfr0El1"
	case HVFeatureRegIDAa64pfr1El1:
		return "HVFeatureRegIDAa64pfr1El1"
	case HVFeatureRegIDAa64smfr0El1:
		return "HVFeatureRegIDAa64smfr0El1"
	case HVFeatureRegIDAa64zfr0El1:
		return "HVFeatureRegIDAa64zfr0El1"
	default:
		return fmt.Sprintf("HVFeatureReg(%d)", e)
	}
}

type HVGICDistributorReg uint

const (
	HVGICDistributorRegGICDCtlr          HVGICDistributorReg = 0
	HVGICDistributorRegGICDIcactiver0    HVGICDistributorReg = 0x380
	HVGICDistributorRegGICDIcactiver1    HVGICDistributorReg = 0x384
	HVGICDistributorRegGICDIcactiver10   HVGICDistributorReg = 0x3a8
	HVGICDistributorRegGICDIcactiver11   HVGICDistributorReg = 0x3ac
	HVGICDistributorRegGICDIcactiver12   HVGICDistributorReg = 0x3b0
	HVGICDistributorRegGICDIcactiver13   HVGICDistributorReg = 0x3b4
	HVGICDistributorRegGICDIcactiver14   HVGICDistributorReg = 0x3b8
	HVGICDistributorRegGICDIcactiver15   HVGICDistributorReg = 0x3bc
	HVGICDistributorRegGICDIcactiver16   HVGICDistributorReg = 0x3c0
	HVGICDistributorRegGICDIcactiver17   HVGICDistributorReg = 0x3c4
	HVGICDistributorRegGICDIcactiver18   HVGICDistributorReg = 0x3c8
	HVGICDistributorRegGICDIcactiver19   HVGICDistributorReg = 0x3cc
	HVGICDistributorRegGICDIcactiver2    HVGICDistributorReg = 0x388
	HVGICDistributorRegGICDIcactiver20   HVGICDistributorReg = 0x3d0
	HVGICDistributorRegGICDIcactiver21   HVGICDistributorReg = 0x3d4
	HVGICDistributorRegGICDIcactiver22   HVGICDistributorReg = 0x3d8
	HVGICDistributorRegGICDIcactiver23   HVGICDistributorReg = 0x3dc
	HVGICDistributorRegGICDIcactiver24   HVGICDistributorReg = 0x3e0
	HVGICDistributorRegGICDIcactiver25   HVGICDistributorReg = 0x3e4
	HVGICDistributorRegGICDIcactiver26   HVGICDistributorReg = 0x3e8
	HVGICDistributorRegGICDIcactiver27   HVGICDistributorReg = 0x3ec
	HVGICDistributorRegGICDIcactiver28   HVGICDistributorReg = 0x3f0
	HVGICDistributorRegGICDIcactiver29   HVGICDistributorReg = 0x3f4
	HVGICDistributorRegGICDIcactiver3    HVGICDistributorReg = 0x38c
	HVGICDistributorRegGICDIcactiver30   HVGICDistributorReg = 0x3f8
	HVGICDistributorRegGICDIcactiver31   HVGICDistributorReg = 0x3fc
	HVGICDistributorRegGICDIcactiver4    HVGICDistributorReg = 0x390
	HVGICDistributorRegGICDIcactiver5    HVGICDistributorReg = 0x394
	HVGICDistributorRegGICDIcactiver6    HVGICDistributorReg = 0x398
	HVGICDistributorRegGICDIcactiver7    HVGICDistributorReg = 0x39c
	HVGICDistributorRegGICDIcactiver8    HVGICDistributorReg = 0x3a0
	HVGICDistributorRegGICDIcactiver9    HVGICDistributorReg = 0x3a4
	HVGICDistributorRegGICDIcenabler0    HVGICDistributorReg = 0x180
	HVGICDistributorRegGICDIcenabler1    HVGICDistributorReg = 0x184
	HVGICDistributorRegGICDIcenabler10   HVGICDistributorReg = 0x1a8
	HVGICDistributorRegGICDIcenabler11   HVGICDistributorReg = 0x1ac
	HVGICDistributorRegGICDIcenabler12   HVGICDistributorReg = 0x1b0
	HVGICDistributorRegGICDIcenabler13   HVGICDistributorReg = 0x1b4
	HVGICDistributorRegGICDIcenabler14   HVGICDistributorReg = 0x1b8
	HVGICDistributorRegGICDIcenabler15   HVGICDistributorReg = 0x1bc
	HVGICDistributorRegGICDIcenabler16   HVGICDistributorReg = 0x1c0
	HVGICDistributorRegGICDIcenabler17   HVGICDistributorReg = 0x1c4
	HVGICDistributorRegGICDIcenabler18   HVGICDistributorReg = 0x1c8
	HVGICDistributorRegGICDIcenabler19   HVGICDistributorReg = 0x1cc
	HVGICDistributorRegGICDIcenabler2    HVGICDistributorReg = 0x188
	HVGICDistributorRegGICDIcenabler20   HVGICDistributorReg = 0x1d0
	HVGICDistributorRegGICDIcenabler21   HVGICDistributorReg = 0x1d4
	HVGICDistributorRegGICDIcenabler22   HVGICDistributorReg = 0x1d8
	HVGICDistributorRegGICDIcenabler23   HVGICDistributorReg = 0x1dc
	HVGICDistributorRegGICDIcenabler24   HVGICDistributorReg = 0x1e0
	HVGICDistributorRegGICDIcenabler25   HVGICDistributorReg = 0x1e4
	HVGICDistributorRegGICDIcenabler26   HVGICDistributorReg = 0x1e8
	HVGICDistributorRegGICDIcenabler27   HVGICDistributorReg = 0x1ec
	HVGICDistributorRegGICDIcenabler28   HVGICDistributorReg = 0x1f0
	HVGICDistributorRegGICDIcenabler29   HVGICDistributorReg = 0x1f4
	HVGICDistributorRegGICDIcenabler3    HVGICDistributorReg = 0x18c
	HVGICDistributorRegGICDIcenabler30   HVGICDistributorReg = 0x1f8
	HVGICDistributorRegGICDIcenabler31   HVGICDistributorReg = 0x1fc
	HVGICDistributorRegGICDIcenabler4    HVGICDistributorReg = 0x190
	HVGICDistributorRegGICDIcenabler5    HVGICDistributorReg = 0x194
	HVGICDistributorRegGICDIcenabler6    HVGICDistributorReg = 0x198
	HVGICDistributorRegGICDIcenabler7    HVGICDistributorReg = 0x19c
	HVGICDistributorRegGICDIcenabler8    HVGICDistributorReg = 0x1a0
	HVGICDistributorRegGICDIcenabler9    HVGICDistributorReg = 0x1a4
	HVGICDistributorRegGICDIcfgr0        HVGICDistributorReg = 0xc00
	HVGICDistributorRegGICDIcfgr1        HVGICDistributorReg = 0xc04
	HVGICDistributorRegGICDIcfgr10       HVGICDistributorReg = 0xc28
	HVGICDistributorRegGICDIcfgr11       HVGICDistributorReg = 0xc2c
	HVGICDistributorRegGICDIcfgr12       HVGICDistributorReg = 0xc30
	HVGICDistributorRegGICDIcfgr13       HVGICDistributorReg = 0xc34
	HVGICDistributorRegGICDIcfgr14       HVGICDistributorReg = 0xc38
	HVGICDistributorRegGICDIcfgr15       HVGICDistributorReg = 0xc3c
	HVGICDistributorRegGICDIcfgr16       HVGICDistributorReg = 0xc40
	HVGICDistributorRegGICDIcfgr17       HVGICDistributorReg = 0xc44
	HVGICDistributorRegGICDIcfgr18       HVGICDistributorReg = 0xc48
	HVGICDistributorRegGICDIcfgr19       HVGICDistributorReg = 0xc4c
	HVGICDistributorRegGICDIcfgr2        HVGICDistributorReg = 0xc08
	HVGICDistributorRegGICDIcfgr20       HVGICDistributorReg = 0xc50
	HVGICDistributorRegGICDIcfgr21       HVGICDistributorReg = 0xc54
	HVGICDistributorRegGICDIcfgr22       HVGICDistributorReg = 0xc58
	HVGICDistributorRegGICDIcfgr23       HVGICDistributorReg = 0xc5c
	HVGICDistributorRegGICDIcfgr24       HVGICDistributorReg = 0xc60
	HVGICDistributorRegGICDIcfgr25       HVGICDistributorReg = 0xc64
	HVGICDistributorRegGICDIcfgr26       HVGICDistributorReg = 0xc68
	HVGICDistributorRegGICDIcfgr27       HVGICDistributorReg = 0xc6c
	HVGICDistributorRegGICDIcfgr28       HVGICDistributorReg = 0xc70
	HVGICDistributorRegGICDIcfgr29       HVGICDistributorReg = 0xc74
	HVGICDistributorRegGICDIcfgr3        HVGICDistributorReg = 0xc0c
	HVGICDistributorRegGICDIcfgr30       HVGICDistributorReg = 0xc78
	HVGICDistributorRegGICDIcfgr31       HVGICDistributorReg = 0xc7c
	HVGICDistributorRegGICDIcfgr32       HVGICDistributorReg = 0xc80
	HVGICDistributorRegGICDIcfgr33       HVGICDistributorReg = 0xc84
	HVGICDistributorRegGICDIcfgr34       HVGICDistributorReg = 0xc88
	HVGICDistributorRegGICDIcfgr35       HVGICDistributorReg = 0xc8c
	HVGICDistributorRegGICDIcfgr36       HVGICDistributorReg = 0xc90
	HVGICDistributorRegGICDIcfgr37       HVGICDistributorReg = 0xc94
	HVGICDistributorRegGICDIcfgr38       HVGICDistributorReg = 0xc98
	HVGICDistributorRegGICDIcfgr39       HVGICDistributorReg = 0xc9c
	HVGICDistributorRegGICDIcfgr4        HVGICDistributorReg = 0xc10
	HVGICDistributorRegGICDIcfgr40       HVGICDistributorReg = 0xca0
	HVGICDistributorRegGICDIcfgr41       HVGICDistributorReg = 0xca4
	HVGICDistributorRegGICDIcfgr42       HVGICDistributorReg = 0xca8
	HVGICDistributorRegGICDIcfgr43       HVGICDistributorReg = 0xcac
	HVGICDistributorRegGICDIcfgr44       HVGICDistributorReg = 0xcb0
	HVGICDistributorRegGICDIcfgr45       HVGICDistributorReg = 0xcb4
	HVGICDistributorRegGICDIcfgr46       HVGICDistributorReg = 0xcb8
	HVGICDistributorRegGICDIcfgr47       HVGICDistributorReg = 0xcbc
	HVGICDistributorRegGICDIcfgr48       HVGICDistributorReg = 0xcc0
	HVGICDistributorRegGICDIcfgr49       HVGICDistributorReg = 0xcc4
	HVGICDistributorRegGICDIcfgr5        HVGICDistributorReg = 0xc14
	HVGICDistributorRegGICDIcfgr50       HVGICDistributorReg = 0xcc8
	HVGICDistributorRegGICDIcfgr51       HVGICDistributorReg = 0xccc
	HVGICDistributorRegGICDIcfgr52       HVGICDistributorReg = 0xcd0
	HVGICDistributorRegGICDIcfgr53       HVGICDistributorReg = 0xcd4
	HVGICDistributorRegGICDIcfgr54       HVGICDistributorReg = 0xcd8
	HVGICDistributorRegGICDIcfgr55       HVGICDistributorReg = 0xcdc
	HVGICDistributorRegGICDIcfgr56       HVGICDistributorReg = 0xce0
	HVGICDistributorRegGICDIcfgr57       HVGICDistributorReg = 0xce4
	HVGICDistributorRegGICDIcfgr58       HVGICDistributorReg = 0xce8
	HVGICDistributorRegGICDIcfgr59       HVGICDistributorReg = 0xcec
	HVGICDistributorRegGICDIcfgr6        HVGICDistributorReg = 0xc18
	HVGICDistributorRegGICDIcfgr60       HVGICDistributorReg = 0xcf0
	HVGICDistributorRegGICDIcfgr61       HVGICDistributorReg = 0xcf4
	HVGICDistributorRegGICDIcfgr62       HVGICDistributorReg = 0xcf8
	HVGICDistributorRegGICDIcfgr63       HVGICDistributorReg = 0xcfc
	HVGICDistributorRegGICDIcfgr7        HVGICDistributorReg = 0xc1c
	HVGICDistributorRegGICDIcfgr8        HVGICDistributorReg = 0xc20
	HVGICDistributorRegGICDIcfgr9        HVGICDistributorReg = 0xc24
	HVGICDistributorRegGICDIcpendr0      HVGICDistributorReg = 0x280
	HVGICDistributorRegGICDIcpendr1      HVGICDistributorReg = 0x284
	HVGICDistributorRegGICDIcpendr10     HVGICDistributorReg = 0x2a8
	HVGICDistributorRegGICDIcpendr11     HVGICDistributorReg = 0x2ac
	HVGICDistributorRegGICDIcpendr12     HVGICDistributorReg = 0x2b0
	HVGICDistributorRegGICDIcpendr13     HVGICDistributorReg = 0x2b4
	HVGICDistributorRegGICDIcpendr14     HVGICDistributorReg = 0x2b8
	HVGICDistributorRegGICDIcpendr15     HVGICDistributorReg = 0x2bc
	HVGICDistributorRegGICDIcpendr16     HVGICDistributorReg = 0x2c0
	HVGICDistributorRegGICDIcpendr17     HVGICDistributorReg = 0x2c4
	HVGICDistributorRegGICDIcpendr18     HVGICDistributorReg = 0x2c8
	HVGICDistributorRegGICDIcpendr19     HVGICDistributorReg = 0x2cc
	HVGICDistributorRegGICDIcpendr2      HVGICDistributorReg = 0x288
	HVGICDistributorRegGICDIcpendr20     HVGICDistributorReg = 0x2d0
	HVGICDistributorRegGICDIcpendr21     HVGICDistributorReg = 0x2d4
	HVGICDistributorRegGICDIcpendr22     HVGICDistributorReg = 0x2d8
	HVGICDistributorRegGICDIcpendr23     HVGICDistributorReg = 0x2dc
	HVGICDistributorRegGICDIcpendr24     HVGICDistributorReg = 0x2e0
	HVGICDistributorRegGICDIcpendr25     HVGICDistributorReg = 0x2e4
	HVGICDistributorRegGICDIcpendr26     HVGICDistributorReg = 0x2e8
	HVGICDistributorRegGICDIcpendr27     HVGICDistributorReg = 0x2ec
	HVGICDistributorRegGICDIcpendr28     HVGICDistributorReg = 0x2f0
	HVGICDistributorRegGICDIcpendr29     HVGICDistributorReg = 0x2f4
	HVGICDistributorRegGICDIcpendr3      HVGICDistributorReg = 0x28c
	HVGICDistributorRegGICDIcpendr30     HVGICDistributorReg = 0x2f8
	HVGICDistributorRegGICDIcpendr31     HVGICDistributorReg = 0x2fc
	HVGICDistributorRegGICDIcpendr4      HVGICDistributorReg = 0x290
	HVGICDistributorRegGICDIcpendr5      HVGICDistributorReg = 0x294
	HVGICDistributorRegGICDIcpendr6      HVGICDistributorReg = 0x298
	HVGICDistributorRegGICDIcpendr7      HVGICDistributorReg = 0x29c
	HVGICDistributorRegGICDIcpendr8      HVGICDistributorReg = 0x2a0
	HVGICDistributorRegGICDIcpendr9      HVGICDistributorReg = 0x2a4
	HVGICDistributorRegGICDIgroupr0      HVGICDistributorReg = 0x80
	HVGICDistributorRegGICDIgroupr1      HVGICDistributorReg = 0x84
	HVGICDistributorRegGICDIgroupr10     HVGICDistributorReg = 0xa8
	HVGICDistributorRegGICDIgroupr11     HVGICDistributorReg = 0xac
	HVGICDistributorRegGICDIgroupr12     HVGICDistributorReg = 0xb0
	HVGICDistributorRegGICDIgroupr13     HVGICDistributorReg = 0xb4
	HVGICDistributorRegGICDIgroupr14     HVGICDistributorReg = 0xb8
	HVGICDistributorRegGICDIgroupr15     HVGICDistributorReg = 0xbc
	HVGICDistributorRegGICDIgroupr16     HVGICDistributorReg = 0xc0
	HVGICDistributorRegGICDIgroupr17     HVGICDistributorReg = 0xc4
	HVGICDistributorRegGICDIgroupr18     HVGICDistributorReg = 0xc8
	HVGICDistributorRegGICDIgroupr19     HVGICDistributorReg = 0xcc
	HVGICDistributorRegGICDIgroupr2      HVGICDistributorReg = 0x88
	HVGICDistributorRegGICDIgroupr20     HVGICDistributorReg = 0xd0
	HVGICDistributorRegGICDIgroupr21     HVGICDistributorReg = 0xd4
	HVGICDistributorRegGICDIgroupr22     HVGICDistributorReg = 0xd8
	HVGICDistributorRegGICDIgroupr23     HVGICDistributorReg = 0xdc
	HVGICDistributorRegGICDIgroupr24     HVGICDistributorReg = 0xe0
	HVGICDistributorRegGICDIgroupr25     HVGICDistributorReg = 0xe4
	HVGICDistributorRegGICDIgroupr26     HVGICDistributorReg = 0xe8
	HVGICDistributorRegGICDIgroupr27     HVGICDistributorReg = 0xec
	HVGICDistributorRegGICDIgroupr28     HVGICDistributorReg = 0xf0
	HVGICDistributorRegGICDIgroupr29     HVGICDistributorReg = 0xf4
	HVGICDistributorRegGICDIgroupr3      HVGICDistributorReg = 0x8c
	HVGICDistributorRegGICDIgroupr30     HVGICDistributorReg = 0xf8
	HVGICDistributorRegGICDIgroupr31     HVGICDistributorReg = 0xfc
	HVGICDistributorRegGICDIgroupr4      HVGICDistributorReg = 0x90
	HVGICDistributorRegGICDIgroupr5      HVGICDistributorReg = 0x94
	HVGICDistributorRegGICDIgroupr6      HVGICDistributorReg = 0x98
	HVGICDistributorRegGICDIgroupr7      HVGICDistributorReg = 0x9c
	HVGICDistributorRegGICDIgroupr8      HVGICDistributorReg = 0xa0
	HVGICDistributorRegGICDIgroupr9      HVGICDistributorReg = 0xa4
	HVGICDistributorRegGICDIpriorityr0   HVGICDistributorReg = 0x400
	HVGICDistributorRegGICDIpriorityr1   HVGICDistributorReg = 0x404
	HVGICDistributorRegGICDIpriorityr10  HVGICDistributorReg = 0x428
	HVGICDistributorRegGICDIpriorityr100 HVGICDistributorReg = 0x590
	HVGICDistributorRegGICDIpriorityr101 HVGICDistributorReg = 0x594
	HVGICDistributorRegGICDIpriorityr102 HVGICDistributorReg = 0x598
	HVGICDistributorRegGICDIpriorityr103 HVGICDistributorReg = 0x59c
	HVGICDistributorRegGICDIpriorityr104 HVGICDistributorReg = 0x5a0
	HVGICDistributorRegGICDIpriorityr105 HVGICDistributorReg = 0x5a4
	HVGICDistributorRegGICDIpriorityr106 HVGICDistributorReg = 0x5a8
	HVGICDistributorRegGICDIpriorityr107 HVGICDistributorReg = 0x5ac
	HVGICDistributorRegGICDIpriorityr108 HVGICDistributorReg = 0x5b0
	HVGICDistributorRegGICDIpriorityr109 HVGICDistributorReg = 0x5b4
	HVGICDistributorRegGICDIpriorityr11  HVGICDistributorReg = 0x42c
	HVGICDistributorRegGICDIpriorityr110 HVGICDistributorReg = 0x5b8
	HVGICDistributorRegGICDIpriorityr111 HVGICDistributorReg = 0x5bc
	HVGICDistributorRegGICDIpriorityr112 HVGICDistributorReg = 0x5c0
	HVGICDistributorRegGICDIpriorityr113 HVGICDistributorReg = 0x5c4
	HVGICDistributorRegGICDIpriorityr114 HVGICDistributorReg = 0x5c8
	HVGICDistributorRegGICDIpriorityr115 HVGICDistributorReg = 0x5cc
	HVGICDistributorRegGICDIpriorityr116 HVGICDistributorReg = 0x5d0
	HVGICDistributorRegGICDIpriorityr117 HVGICDistributorReg = 0x5d4
	HVGICDistributorRegGICDIpriorityr118 HVGICDistributorReg = 0x5d8
	HVGICDistributorRegGICDIpriorityr119 HVGICDistributorReg = 0x5dc
	HVGICDistributorRegGICDIpriorityr12  HVGICDistributorReg = 0x430
	HVGICDistributorRegGICDIpriorityr120 HVGICDistributorReg = 0x5e0
	HVGICDistributorRegGICDIpriorityr121 HVGICDistributorReg = 0x5e4
	HVGICDistributorRegGICDIpriorityr122 HVGICDistributorReg = 0x5e8
	HVGICDistributorRegGICDIpriorityr123 HVGICDistributorReg = 0x5ec
	HVGICDistributorRegGICDIpriorityr124 HVGICDistributorReg = 0x5f0
	HVGICDistributorRegGICDIpriorityr125 HVGICDistributorReg = 0x5f4
	HVGICDistributorRegGICDIpriorityr126 HVGICDistributorReg = 0x5f8
	HVGICDistributorRegGICDIpriorityr127 HVGICDistributorReg = 0x5fc
	HVGICDistributorRegGICDIpriorityr128 HVGICDistributorReg = 0x600
	HVGICDistributorRegGICDIpriorityr129 HVGICDistributorReg = 0x604
	HVGICDistributorRegGICDIpriorityr13  HVGICDistributorReg = 0x434
	HVGICDistributorRegGICDIpriorityr130 HVGICDistributorReg = 0x608
	HVGICDistributorRegGICDIpriorityr131 HVGICDistributorReg = 0x60c
	HVGICDistributorRegGICDIpriorityr132 HVGICDistributorReg = 0x610
	HVGICDistributorRegGICDIpriorityr133 HVGICDistributorReg = 0x614
	HVGICDistributorRegGICDIpriorityr134 HVGICDistributorReg = 0x618
	HVGICDistributorRegGICDIpriorityr135 HVGICDistributorReg = 0x61c
	HVGICDistributorRegGICDIpriorityr136 HVGICDistributorReg = 0x620
	HVGICDistributorRegGICDIpriorityr137 HVGICDistributorReg = 0x624
	HVGICDistributorRegGICDIpriorityr138 HVGICDistributorReg = 0x628
	HVGICDistributorRegGICDIpriorityr139 HVGICDistributorReg = 0x62c
	HVGICDistributorRegGICDIpriorityr14  HVGICDistributorReg = 0x438
	HVGICDistributorRegGICDIpriorityr140 HVGICDistributorReg = 0x630
	HVGICDistributorRegGICDIpriorityr141 HVGICDistributorReg = 0x634
	HVGICDistributorRegGICDIpriorityr142 HVGICDistributorReg = 0x638
	HVGICDistributorRegGICDIpriorityr143 HVGICDistributorReg = 0x63c
	HVGICDistributorRegGICDIpriorityr144 HVGICDistributorReg = 0x640
	HVGICDistributorRegGICDIpriorityr145 HVGICDistributorReg = 0x644
	HVGICDistributorRegGICDIpriorityr146 HVGICDistributorReg = 0x648
	HVGICDistributorRegGICDIpriorityr147 HVGICDistributorReg = 0x64c
	HVGICDistributorRegGICDIpriorityr148 HVGICDistributorReg = 0x650
	HVGICDistributorRegGICDIpriorityr149 HVGICDistributorReg = 0x654
	HVGICDistributorRegGICDIpriorityr15  HVGICDistributorReg = 0x43c
	HVGICDistributorRegGICDIpriorityr150 HVGICDistributorReg = 0x658
	HVGICDistributorRegGICDIpriorityr151 HVGICDistributorReg = 0x65c
	HVGICDistributorRegGICDIpriorityr152 HVGICDistributorReg = 0x660
	HVGICDistributorRegGICDIpriorityr153 HVGICDistributorReg = 0x664
	HVGICDistributorRegGICDIpriorityr154 HVGICDistributorReg = 0x668
	HVGICDistributorRegGICDIpriorityr155 HVGICDistributorReg = 0x66c
	HVGICDistributorRegGICDIpriorityr156 HVGICDistributorReg = 0x670
	HVGICDistributorRegGICDIpriorityr157 HVGICDistributorReg = 0x674
	HVGICDistributorRegGICDIpriorityr158 HVGICDistributorReg = 0x678
	HVGICDistributorRegGICDIpriorityr159 HVGICDistributorReg = 0x67c
	HVGICDistributorRegGICDIpriorityr16  HVGICDistributorReg = 0x440
	HVGICDistributorRegGICDIpriorityr160 HVGICDistributorReg = 0x680
	HVGICDistributorRegGICDIpriorityr161 HVGICDistributorReg = 0x684
	HVGICDistributorRegGICDIpriorityr162 HVGICDistributorReg = 0x688
	HVGICDistributorRegGICDIpriorityr163 HVGICDistributorReg = 0x68c
	HVGICDistributorRegGICDIpriorityr164 HVGICDistributorReg = 0x690
	HVGICDistributorRegGICDIpriorityr165 HVGICDistributorReg = 0x694
	HVGICDistributorRegGICDIpriorityr166 HVGICDistributorReg = 0x698
	HVGICDistributorRegGICDIpriorityr167 HVGICDistributorReg = 0x69c
	HVGICDistributorRegGICDIpriorityr168 HVGICDistributorReg = 0x6a0
	HVGICDistributorRegGICDIpriorityr169 HVGICDistributorReg = 0x6a4
	HVGICDistributorRegGICDIpriorityr17  HVGICDistributorReg = 0x444
	HVGICDistributorRegGICDIpriorityr170 HVGICDistributorReg = 0x6a8
	HVGICDistributorRegGICDIpriorityr171 HVGICDistributorReg = 0x6ac
	HVGICDistributorRegGICDIpriorityr172 HVGICDistributorReg = 0x6b0
	HVGICDistributorRegGICDIpriorityr173 HVGICDistributorReg = 0x6b4
	HVGICDistributorRegGICDIpriorityr174 HVGICDistributorReg = 0x6b8
	HVGICDistributorRegGICDIpriorityr175 HVGICDistributorReg = 0x6bc
	HVGICDistributorRegGICDIpriorityr176 HVGICDistributorReg = 0x6c0
	HVGICDistributorRegGICDIpriorityr177 HVGICDistributorReg = 0x6c4
	HVGICDistributorRegGICDIpriorityr178 HVGICDistributorReg = 0x6c8
	HVGICDistributorRegGICDIpriorityr179 HVGICDistributorReg = 0x6cc
	HVGICDistributorRegGICDIpriorityr18  HVGICDistributorReg = 0x448
	HVGICDistributorRegGICDIpriorityr180 HVGICDistributorReg = 0x6d0
	HVGICDistributorRegGICDIpriorityr181 HVGICDistributorReg = 0x6d4
	HVGICDistributorRegGICDIpriorityr182 HVGICDistributorReg = 0x6d8
	HVGICDistributorRegGICDIpriorityr183 HVGICDistributorReg = 0x6dc
	HVGICDistributorRegGICDIpriorityr184 HVGICDistributorReg = 0x6e0
	HVGICDistributorRegGICDIpriorityr185 HVGICDistributorReg = 0x6e4
	HVGICDistributorRegGICDIpriorityr186 HVGICDistributorReg = 0x6e8
	HVGICDistributorRegGICDIpriorityr187 HVGICDistributorReg = 0x6ec
	HVGICDistributorRegGICDIpriorityr188 HVGICDistributorReg = 0x6f0
	HVGICDistributorRegGICDIpriorityr189 HVGICDistributorReg = 0x6f4
	HVGICDistributorRegGICDIpriorityr19  HVGICDistributorReg = 0x44c
	HVGICDistributorRegGICDIpriorityr190 HVGICDistributorReg = 0x6f8
	HVGICDistributorRegGICDIpriorityr191 HVGICDistributorReg = 0x6fc
	HVGICDistributorRegGICDIpriorityr192 HVGICDistributorReg = 0x700
	HVGICDistributorRegGICDIpriorityr193 HVGICDistributorReg = 0x704
	HVGICDistributorRegGICDIpriorityr194 HVGICDistributorReg = 0x708
	HVGICDistributorRegGICDIpriorityr195 HVGICDistributorReg = 0x70c
	HVGICDistributorRegGICDIpriorityr196 HVGICDistributorReg = 0x710
	HVGICDistributorRegGICDIpriorityr197 HVGICDistributorReg = 0x714
	HVGICDistributorRegGICDIpriorityr198 HVGICDistributorReg = 0x718
	HVGICDistributorRegGICDIpriorityr199 HVGICDistributorReg = 0x71c
	HVGICDistributorRegGICDIpriorityr2   HVGICDistributorReg = 0x408
	HVGICDistributorRegGICDIpriorityr20  HVGICDistributorReg = 0x450
	HVGICDistributorRegGICDIpriorityr200 HVGICDistributorReg = 0x720
	HVGICDistributorRegGICDIpriorityr201 HVGICDistributorReg = 0x724
	HVGICDistributorRegGICDIpriorityr202 HVGICDistributorReg = 0x728
	HVGICDistributorRegGICDIpriorityr203 HVGICDistributorReg = 0x72c
	HVGICDistributorRegGICDIpriorityr204 HVGICDistributorReg = 0x730
	HVGICDistributorRegGICDIpriorityr205 HVGICDistributorReg = 0x734
	HVGICDistributorRegGICDIpriorityr206 HVGICDistributorReg = 0x738
	HVGICDistributorRegGICDIpriorityr207 HVGICDistributorReg = 0x73c
	HVGICDistributorRegGICDIpriorityr208 HVGICDistributorReg = 0x740
	HVGICDistributorRegGICDIpriorityr209 HVGICDistributorReg = 0x744
	HVGICDistributorRegGICDIpriorityr21  HVGICDistributorReg = 0x454
	HVGICDistributorRegGICDIpriorityr210 HVGICDistributorReg = 0x748
	HVGICDistributorRegGICDIpriorityr211 HVGICDistributorReg = 0x74c
	HVGICDistributorRegGICDIpriorityr212 HVGICDistributorReg = 0x750
	HVGICDistributorRegGICDIpriorityr213 HVGICDistributorReg = 0x754
	HVGICDistributorRegGICDIpriorityr214 HVGICDistributorReg = 0x758
	HVGICDistributorRegGICDIpriorityr215 HVGICDistributorReg = 0x75c
	HVGICDistributorRegGICDIpriorityr216 HVGICDistributorReg = 0x760
	HVGICDistributorRegGICDIpriorityr217 HVGICDistributorReg = 0x764
	HVGICDistributorRegGICDIpriorityr218 HVGICDistributorReg = 0x768
	HVGICDistributorRegGICDIpriorityr219 HVGICDistributorReg = 0x76c
	HVGICDistributorRegGICDIpriorityr22  HVGICDistributorReg = 0x458
	HVGICDistributorRegGICDIpriorityr220 HVGICDistributorReg = 0x770
	HVGICDistributorRegGICDIpriorityr221 HVGICDistributorReg = 0x774
	HVGICDistributorRegGICDIpriorityr222 HVGICDistributorReg = 0x778
	HVGICDistributorRegGICDIpriorityr223 HVGICDistributorReg = 0x77c
	HVGICDistributorRegGICDIpriorityr224 HVGICDistributorReg = 0x780
	HVGICDistributorRegGICDIpriorityr225 HVGICDistributorReg = 0x784
	HVGICDistributorRegGICDIpriorityr226 HVGICDistributorReg = 0x788
	HVGICDistributorRegGICDIpriorityr227 HVGICDistributorReg = 0x78c
	HVGICDistributorRegGICDIpriorityr228 HVGICDistributorReg = 0x790
	HVGICDistributorRegGICDIpriorityr229 HVGICDistributorReg = 0x794
	HVGICDistributorRegGICDIpriorityr23  HVGICDistributorReg = 0x45c
	HVGICDistributorRegGICDIpriorityr230 HVGICDistributorReg = 0x798
	HVGICDistributorRegGICDIpriorityr231 HVGICDistributorReg = 0x79c
	HVGICDistributorRegGICDIpriorityr232 HVGICDistributorReg = 0x7a0
	HVGICDistributorRegGICDIpriorityr233 HVGICDistributorReg = 0x7a4
	HVGICDistributorRegGICDIpriorityr234 HVGICDistributorReg = 0x7a8
	HVGICDistributorRegGICDIpriorityr235 HVGICDistributorReg = 0x7ac
	HVGICDistributorRegGICDIpriorityr236 HVGICDistributorReg = 0x7b0
	HVGICDistributorRegGICDIpriorityr237 HVGICDistributorReg = 0x7b4
	HVGICDistributorRegGICDIpriorityr238 HVGICDistributorReg = 0x7b8
	HVGICDistributorRegGICDIpriorityr239 HVGICDistributorReg = 0x7bc
	HVGICDistributorRegGICDIpriorityr24  HVGICDistributorReg = 0x460
	HVGICDistributorRegGICDIpriorityr240 HVGICDistributorReg = 0x7c0
	HVGICDistributorRegGICDIpriorityr241 HVGICDistributorReg = 0x7c4
	HVGICDistributorRegGICDIpriorityr242 HVGICDistributorReg = 0x7c8
	HVGICDistributorRegGICDIpriorityr243 HVGICDistributorReg = 0x7cc
	HVGICDistributorRegGICDIpriorityr244 HVGICDistributorReg = 0x7d0
	HVGICDistributorRegGICDIpriorityr245 HVGICDistributorReg = 0x7d4
	HVGICDistributorRegGICDIpriorityr246 HVGICDistributorReg = 0x7d8
	HVGICDistributorRegGICDIpriorityr247 HVGICDistributorReg = 0x7dc
	HVGICDistributorRegGICDIpriorityr248 HVGICDistributorReg = 0x7e0
	HVGICDistributorRegGICDIpriorityr249 HVGICDistributorReg = 0x7e4
	HVGICDistributorRegGICDIpriorityr25  HVGICDistributorReg = 0x464
	HVGICDistributorRegGICDIpriorityr250 HVGICDistributorReg = 0x7e8
	HVGICDistributorRegGICDIpriorityr251 HVGICDistributorReg = 0x7ec
	HVGICDistributorRegGICDIpriorityr252 HVGICDistributorReg = 0x7f0
	HVGICDistributorRegGICDIpriorityr253 HVGICDistributorReg = 0x7f4
	HVGICDistributorRegGICDIpriorityr254 HVGICDistributorReg = 0x7f8
	HVGICDistributorRegGICDIpriorityr26  HVGICDistributorReg = 0x468
	HVGICDistributorRegGICDIpriorityr27  HVGICDistributorReg = 0x46c
	HVGICDistributorRegGICDIpriorityr28  HVGICDistributorReg = 0x470
	HVGICDistributorRegGICDIpriorityr29  HVGICDistributorReg = 0x474
	HVGICDistributorRegGICDIpriorityr3   HVGICDistributorReg = 0x40c
	HVGICDistributorRegGICDIpriorityr30  HVGICDistributorReg = 0x478
	HVGICDistributorRegGICDIpriorityr31  HVGICDistributorReg = 0x47c
	HVGICDistributorRegGICDIpriorityr32  HVGICDistributorReg = 0x480
	HVGICDistributorRegGICDIpriorityr33  HVGICDistributorReg = 0x484
	HVGICDistributorRegGICDIpriorityr34  HVGICDistributorReg = 0x488
	HVGICDistributorRegGICDIpriorityr35  HVGICDistributorReg = 0x48c
	HVGICDistributorRegGICDIpriorityr36  HVGICDistributorReg = 0x490
	HVGICDistributorRegGICDIpriorityr37  HVGICDistributorReg = 0x494
	HVGICDistributorRegGICDIpriorityr38  HVGICDistributorReg = 0x498
	HVGICDistributorRegGICDIpriorityr39  HVGICDistributorReg = 0x49c
	HVGICDistributorRegGICDIpriorityr4   HVGICDistributorReg = 0x410
	HVGICDistributorRegGICDIpriorityr40  HVGICDistributorReg = 0x4a0
	HVGICDistributorRegGICDIpriorityr41  HVGICDistributorReg = 0x4a4
	HVGICDistributorRegGICDIpriorityr42  HVGICDistributorReg = 0x4a8
	HVGICDistributorRegGICDIpriorityr43  HVGICDistributorReg = 0x4ac
	HVGICDistributorRegGICDIpriorityr44  HVGICDistributorReg = 0x4b0
	HVGICDistributorRegGICDIpriorityr45  HVGICDistributorReg = 0x4b4
	HVGICDistributorRegGICDIpriorityr46  HVGICDistributorReg = 0x4b8
	HVGICDistributorRegGICDIpriorityr47  HVGICDistributorReg = 0x4bc
	HVGICDistributorRegGICDIpriorityr48  HVGICDistributorReg = 0x4c0
	HVGICDistributorRegGICDIpriorityr49  HVGICDistributorReg = 0x4c4
	HVGICDistributorRegGICDIpriorityr5   HVGICDistributorReg = 0x414
	HVGICDistributorRegGICDIpriorityr50  HVGICDistributorReg = 0x4c8
	HVGICDistributorRegGICDIpriorityr51  HVGICDistributorReg = 0x4cc
	HVGICDistributorRegGICDIpriorityr52  HVGICDistributorReg = 0x4d0
	HVGICDistributorRegGICDIpriorityr53  HVGICDistributorReg = 0x4d4
	HVGICDistributorRegGICDIpriorityr54  HVGICDistributorReg = 0x4d8
	HVGICDistributorRegGICDIpriorityr55  HVGICDistributorReg = 0x4dc
	HVGICDistributorRegGICDIpriorityr56  HVGICDistributorReg = 0x4e0
	HVGICDistributorRegGICDIpriorityr57  HVGICDistributorReg = 0x4e4
	HVGICDistributorRegGICDIpriorityr58  HVGICDistributorReg = 0x4e8
	HVGICDistributorRegGICDIpriorityr59  HVGICDistributorReg = 0x4ec
	HVGICDistributorRegGICDIpriorityr6   HVGICDistributorReg = 0x418
	HVGICDistributorRegGICDIpriorityr60  HVGICDistributorReg = 0x4f0
	HVGICDistributorRegGICDIpriorityr61  HVGICDistributorReg = 0x4f4
	HVGICDistributorRegGICDIpriorityr62  HVGICDistributorReg = 0x4f8
	HVGICDistributorRegGICDIpriorityr63  HVGICDistributorReg = 0x4fc
	HVGICDistributorRegGICDIpriorityr64  HVGICDistributorReg = 0x500
	HVGICDistributorRegGICDIpriorityr65  HVGICDistributorReg = 0x504
	HVGICDistributorRegGICDIpriorityr66  HVGICDistributorReg = 0x508
	HVGICDistributorRegGICDIpriorityr67  HVGICDistributorReg = 0x50c
	HVGICDistributorRegGICDIpriorityr68  HVGICDistributorReg = 0x510
	HVGICDistributorRegGICDIpriorityr69  HVGICDistributorReg = 0x514
	HVGICDistributorRegGICDIpriorityr7   HVGICDistributorReg = 0x41c
	HVGICDistributorRegGICDIpriorityr70  HVGICDistributorReg = 0x518
	HVGICDistributorRegGICDIpriorityr71  HVGICDistributorReg = 0x51c
	HVGICDistributorRegGICDIpriorityr72  HVGICDistributorReg = 0x520
	HVGICDistributorRegGICDIpriorityr73  HVGICDistributorReg = 0x524
	HVGICDistributorRegGICDIpriorityr74  HVGICDistributorReg = 0x528
	HVGICDistributorRegGICDIpriorityr75  HVGICDistributorReg = 0x52c
	HVGICDistributorRegGICDIpriorityr76  HVGICDistributorReg = 0x530
	HVGICDistributorRegGICDIpriorityr77  HVGICDistributorReg = 0x534
	HVGICDistributorRegGICDIpriorityr78  HVGICDistributorReg = 0x538
	HVGICDistributorRegGICDIpriorityr79  HVGICDistributorReg = 0x53c
	HVGICDistributorRegGICDIpriorityr8   HVGICDistributorReg = 0x420
	HVGICDistributorRegGICDIpriorityr80  HVGICDistributorReg = 0x540
	HVGICDistributorRegGICDIpriorityr81  HVGICDistributorReg = 0x544
	HVGICDistributorRegGICDIpriorityr82  HVGICDistributorReg = 0x548
	HVGICDistributorRegGICDIpriorityr83  HVGICDistributorReg = 0x54c
	HVGICDistributorRegGICDIpriorityr84  HVGICDistributorReg = 0x550
	HVGICDistributorRegGICDIpriorityr85  HVGICDistributorReg = 0x554
	HVGICDistributorRegGICDIpriorityr86  HVGICDistributorReg = 0x558
	HVGICDistributorRegGICDIpriorityr87  HVGICDistributorReg = 0x55c
	HVGICDistributorRegGICDIpriorityr88  HVGICDistributorReg = 0x560
	HVGICDistributorRegGICDIpriorityr89  HVGICDistributorReg = 0x564
	HVGICDistributorRegGICDIpriorityr9   HVGICDistributorReg = 0x424
	HVGICDistributorRegGICDIpriorityr90  HVGICDistributorReg = 0x568
	HVGICDistributorRegGICDIpriorityr91  HVGICDistributorReg = 0x56c
	HVGICDistributorRegGICDIpriorityr92  HVGICDistributorReg = 0x570
	HVGICDistributorRegGICDIpriorityr93  HVGICDistributorReg = 0x574
	HVGICDistributorRegGICDIpriorityr94  HVGICDistributorReg = 0x578
	HVGICDistributorRegGICDIpriorityr95  HVGICDistributorReg = 0x57c
	HVGICDistributorRegGICDIpriorityr96  HVGICDistributorReg = 0x580
	HVGICDistributorRegGICDIpriorityr97  HVGICDistributorReg = 0x584
	HVGICDistributorRegGICDIpriorityr98  HVGICDistributorReg = 0x588
	HVGICDistributorRegGICDIpriorityr99  HVGICDistributorReg = 0x58c
	HVGICDistributorRegGICDIrouter100    HVGICDistributorReg = 0x6320
	HVGICDistributorRegGICDIrouter1000   HVGICDistributorReg = 0x7f40
	HVGICDistributorRegGICDIrouter1001   HVGICDistributorReg = 0x7f48
	HVGICDistributorRegGICDIrouter1002   HVGICDistributorReg = 0x7f50
	HVGICDistributorRegGICDIrouter1003   HVGICDistributorReg = 0x7f58
	HVGICDistributorRegGICDIrouter1004   HVGICDistributorReg = 0x7f60
	HVGICDistributorRegGICDIrouter1005   HVGICDistributorReg = 0x7f68
	HVGICDistributorRegGICDIrouter1006   HVGICDistributorReg = 0x7f70
	HVGICDistributorRegGICDIrouter1007   HVGICDistributorReg = 0x7f78
	HVGICDistributorRegGICDIrouter1008   HVGICDistributorReg = 0x7f80
	HVGICDistributorRegGICDIrouter1009   HVGICDistributorReg = 0x7f88
	HVGICDistributorRegGICDIrouter101    HVGICDistributorReg = 0x6328
	HVGICDistributorRegGICDIrouter1010   HVGICDistributorReg = 0x7f90
	HVGICDistributorRegGICDIrouter1011   HVGICDistributorReg = 0x7f98
	HVGICDistributorRegGICDIrouter1012   HVGICDistributorReg = 0x7fa0
	HVGICDistributorRegGICDIrouter1013   HVGICDistributorReg = 0x7fa8
	HVGICDistributorRegGICDIrouter1014   HVGICDistributorReg = 0x7fb0
	HVGICDistributorRegGICDIrouter1015   HVGICDistributorReg = 0x7fb8
	HVGICDistributorRegGICDIrouter1016   HVGICDistributorReg = 0x7fc0
	HVGICDistributorRegGICDIrouter1017   HVGICDistributorReg = 0x7fc8
	HVGICDistributorRegGICDIrouter1018   HVGICDistributorReg = 0x7fd0
	HVGICDistributorRegGICDIrouter1019   HVGICDistributorReg = 0x7fd8
	HVGICDistributorRegGICDIrouter102    HVGICDistributorReg = 0x6330
	HVGICDistributorRegGICDIrouter103    HVGICDistributorReg = 0x6338
	HVGICDistributorRegGICDIrouter104    HVGICDistributorReg = 0x6340
	HVGICDistributorRegGICDIrouter105    HVGICDistributorReg = 0x6348
	HVGICDistributorRegGICDIrouter106    HVGICDistributorReg = 0x6350
	HVGICDistributorRegGICDIrouter107    HVGICDistributorReg = 0x6358
	HVGICDistributorRegGICDIrouter108    HVGICDistributorReg = 0x6360
	HVGICDistributorRegGICDIrouter109    HVGICDistributorReg = 0x6368
	HVGICDistributorRegGICDIrouter110    HVGICDistributorReg = 0x6370
	HVGICDistributorRegGICDIrouter111    HVGICDistributorReg = 0x6378
	HVGICDistributorRegGICDIrouter112    HVGICDistributorReg = 0x6380
	HVGICDistributorRegGICDIrouter113    HVGICDistributorReg = 0x6388
	HVGICDistributorRegGICDIrouter114    HVGICDistributorReg = 0x6390
	HVGICDistributorRegGICDIrouter115    HVGICDistributorReg = 0x6398
	HVGICDistributorRegGICDIrouter116    HVGICDistributorReg = 0x63a0
	HVGICDistributorRegGICDIrouter117    HVGICDistributorReg = 0x63a8
	HVGICDistributorRegGICDIrouter118    HVGICDistributorReg = 0x63b0
	HVGICDistributorRegGICDIrouter119    HVGICDistributorReg = 0x63b8
	HVGICDistributorRegGICDIrouter120    HVGICDistributorReg = 0x63c0
	HVGICDistributorRegGICDIrouter121    HVGICDistributorReg = 0x63c8
	HVGICDistributorRegGICDIrouter122    HVGICDistributorReg = 0x63d0
	HVGICDistributorRegGICDIrouter123    HVGICDistributorReg = 0x63d8
	HVGICDistributorRegGICDIrouter124    HVGICDistributorReg = 0x63e0
	HVGICDistributorRegGICDIrouter125    HVGICDistributorReg = 0x63e8
	HVGICDistributorRegGICDIrouter126    HVGICDistributorReg = 0x63f0
	HVGICDistributorRegGICDIrouter127    HVGICDistributorReg = 0x63f8
	HVGICDistributorRegGICDIrouter128    HVGICDistributorReg = 0x6400
	HVGICDistributorRegGICDIrouter129    HVGICDistributorReg = 0x6408
	HVGICDistributorRegGICDIrouter130    HVGICDistributorReg = 0x6410
	HVGICDistributorRegGICDIrouter131    HVGICDistributorReg = 0x6418
	HVGICDistributorRegGICDIrouter132    HVGICDistributorReg = 0x6420
	HVGICDistributorRegGICDIrouter133    HVGICDistributorReg = 0x6428
	HVGICDistributorRegGICDIrouter134    HVGICDistributorReg = 0x6430
	HVGICDistributorRegGICDIrouter135    HVGICDistributorReg = 0x6438
	HVGICDistributorRegGICDIrouter136    HVGICDistributorReg = 0x6440
	HVGICDistributorRegGICDIrouter137    HVGICDistributorReg = 0x6448
	HVGICDistributorRegGICDIrouter138    HVGICDistributorReg = 0x6450
	HVGICDistributorRegGICDIrouter139    HVGICDistributorReg = 0x6458
	HVGICDistributorRegGICDIrouter140    HVGICDistributorReg = 0x6460
	HVGICDistributorRegGICDIrouter141    HVGICDistributorReg = 0x6468
	HVGICDistributorRegGICDIrouter142    HVGICDistributorReg = 0x6470
	HVGICDistributorRegGICDIrouter143    HVGICDistributorReg = 0x6478
	HVGICDistributorRegGICDIrouter144    HVGICDistributorReg = 0x6480
	HVGICDistributorRegGICDIrouter145    HVGICDistributorReg = 0x6488
	HVGICDistributorRegGICDIrouter146    HVGICDistributorReg = 0x6490
	HVGICDistributorRegGICDIrouter147    HVGICDistributorReg = 0x6498
	HVGICDistributorRegGICDIrouter148    HVGICDistributorReg = 0x64a0
	HVGICDistributorRegGICDIrouter149    HVGICDistributorReg = 0x64a8
	HVGICDistributorRegGICDIrouter150    HVGICDistributorReg = 0x64b0
	HVGICDistributorRegGICDIrouter151    HVGICDistributorReg = 0x64b8
	HVGICDistributorRegGICDIrouter152    HVGICDistributorReg = 0x64c0
	HVGICDistributorRegGICDIrouter153    HVGICDistributorReg = 0x64c8
	HVGICDistributorRegGICDIrouter154    HVGICDistributorReg = 0x64d0
	HVGICDistributorRegGICDIrouter155    HVGICDistributorReg = 0x64d8
	HVGICDistributorRegGICDIrouter156    HVGICDistributorReg = 0x64e0
	HVGICDistributorRegGICDIrouter157    HVGICDistributorReg = 0x64e8
	HVGICDistributorRegGICDIrouter158    HVGICDistributorReg = 0x64f0
	HVGICDistributorRegGICDIrouter159    HVGICDistributorReg = 0x64f8
	HVGICDistributorRegGICDIrouter160    HVGICDistributorReg = 0x6500
	HVGICDistributorRegGICDIrouter161    HVGICDistributorReg = 0x6508
	HVGICDistributorRegGICDIrouter162    HVGICDistributorReg = 0x6510
	HVGICDistributorRegGICDIrouter163    HVGICDistributorReg = 0x6518
	HVGICDistributorRegGICDIrouter164    HVGICDistributorReg = 0x6520
	HVGICDistributorRegGICDIrouter165    HVGICDistributorReg = 0x6528
	HVGICDistributorRegGICDIrouter166    HVGICDistributorReg = 0x6530
	HVGICDistributorRegGICDIrouter167    HVGICDistributorReg = 0x6538
	HVGICDistributorRegGICDIrouter168    HVGICDistributorReg = 0x6540
	HVGICDistributorRegGICDIrouter169    HVGICDistributorReg = 0x6548
	HVGICDistributorRegGICDIrouter170    HVGICDistributorReg = 0x6550
	HVGICDistributorRegGICDIrouter171    HVGICDistributorReg = 0x6558
	HVGICDistributorRegGICDIrouter172    HVGICDistributorReg = 0x6560
	HVGICDistributorRegGICDIrouter173    HVGICDistributorReg = 0x6568
	HVGICDistributorRegGICDIrouter174    HVGICDistributorReg = 0x6570
	HVGICDistributorRegGICDIrouter175    HVGICDistributorReg = 0x6578
	HVGICDistributorRegGICDIrouter176    HVGICDistributorReg = 0x6580
	HVGICDistributorRegGICDIrouter177    HVGICDistributorReg = 0x6588
	HVGICDistributorRegGICDIrouter178    HVGICDistributorReg = 0x6590
	HVGICDistributorRegGICDIrouter179    HVGICDistributorReg = 0x6598
	HVGICDistributorRegGICDIrouter180    HVGICDistributorReg = 0x65a0
	HVGICDistributorRegGICDIrouter181    HVGICDistributorReg = 0x65a8
	HVGICDistributorRegGICDIrouter182    HVGICDistributorReg = 0x65b0
	HVGICDistributorRegGICDIrouter183    HVGICDistributorReg = 0x65b8
	HVGICDistributorRegGICDIrouter184    HVGICDistributorReg = 0x65c0
	HVGICDistributorRegGICDIrouter185    HVGICDistributorReg = 0x65c8
	HVGICDistributorRegGICDIrouter186    HVGICDistributorReg = 0x65d0
	HVGICDistributorRegGICDIrouter187    HVGICDistributorReg = 0x65d8
	HVGICDistributorRegGICDIrouter188    HVGICDistributorReg = 0x65e0
	HVGICDistributorRegGICDIrouter189    HVGICDistributorReg = 0x65e8
	HVGICDistributorRegGICDIrouter190    HVGICDistributorReg = 0x65f0
	HVGICDistributorRegGICDIrouter191    HVGICDistributorReg = 0x65f8
	HVGICDistributorRegGICDIrouter192    HVGICDistributorReg = 0x6600
	HVGICDistributorRegGICDIrouter193    HVGICDistributorReg = 0x6608
	HVGICDistributorRegGICDIrouter194    HVGICDistributorReg = 0x6610
	HVGICDistributorRegGICDIrouter195    HVGICDistributorReg = 0x6618
	HVGICDistributorRegGICDIrouter196    HVGICDistributorReg = 0x6620
	HVGICDistributorRegGICDIrouter197    HVGICDistributorReg = 0x6628
	HVGICDistributorRegGICDIrouter198    HVGICDistributorReg = 0x6630
	HVGICDistributorRegGICDIrouter199    HVGICDistributorReg = 0x6638
	HVGICDistributorRegGICDIrouter200    HVGICDistributorReg = 0x6640
	HVGICDistributorRegGICDIrouter201    HVGICDistributorReg = 0x6648
	HVGICDistributorRegGICDIrouter202    HVGICDistributorReg = 0x6650
	HVGICDistributorRegGICDIrouter203    HVGICDistributorReg = 0x6658
	HVGICDistributorRegGICDIrouter204    HVGICDistributorReg = 0x6660
	HVGICDistributorRegGICDIrouter205    HVGICDistributorReg = 0x6668
	HVGICDistributorRegGICDIrouter206    HVGICDistributorReg = 0x6670
	HVGICDistributorRegGICDIrouter207    HVGICDistributorReg = 0x6678
	HVGICDistributorRegGICDIrouter208    HVGICDistributorReg = 0x6680
	HVGICDistributorRegGICDIrouter209    HVGICDistributorReg = 0x6688
	HVGICDistributorRegGICDIrouter210    HVGICDistributorReg = 0x6690
	HVGICDistributorRegGICDIrouter211    HVGICDistributorReg = 0x6698
	HVGICDistributorRegGICDIrouter212    HVGICDistributorReg = 0x66a0
	HVGICDistributorRegGICDIrouter213    HVGICDistributorReg = 0x66a8
	HVGICDistributorRegGICDIrouter214    HVGICDistributorReg = 0x66b0
	HVGICDistributorRegGICDIrouter215    HVGICDistributorReg = 0x66b8
	HVGICDistributorRegGICDIrouter216    HVGICDistributorReg = 0x66c0
	HVGICDistributorRegGICDIrouter217    HVGICDistributorReg = 0x66c8
	HVGICDistributorRegGICDIrouter218    HVGICDistributorReg = 0x66d0
	HVGICDistributorRegGICDIrouter219    HVGICDistributorReg = 0x66d8
	HVGICDistributorRegGICDIrouter220    HVGICDistributorReg = 0x66e0
	HVGICDistributorRegGICDIrouter221    HVGICDistributorReg = 0x66e8
	HVGICDistributorRegGICDIrouter222    HVGICDistributorReg = 0x66f0
	HVGICDistributorRegGICDIrouter223    HVGICDistributorReg = 0x66f8
	HVGICDistributorRegGICDIrouter224    HVGICDistributorReg = 0x6700
	HVGICDistributorRegGICDIrouter225    HVGICDistributorReg = 0x6708
	HVGICDistributorRegGICDIrouter226    HVGICDistributorReg = 0x6710
	HVGICDistributorRegGICDIrouter227    HVGICDistributorReg = 0x6718
	HVGICDistributorRegGICDIrouter228    HVGICDistributorReg = 0x6720
	HVGICDistributorRegGICDIrouter229    HVGICDistributorReg = 0x6728
	HVGICDistributorRegGICDIrouter230    HVGICDistributorReg = 0x6730
	HVGICDistributorRegGICDIrouter231    HVGICDistributorReg = 0x6738
	HVGICDistributorRegGICDIrouter232    HVGICDistributorReg = 0x6740
	HVGICDistributorRegGICDIrouter233    HVGICDistributorReg = 0x6748
	HVGICDistributorRegGICDIrouter234    HVGICDistributorReg = 0x6750
	HVGICDistributorRegGICDIrouter235    HVGICDistributorReg = 0x6758
	HVGICDistributorRegGICDIrouter236    HVGICDistributorReg = 0x6760
	HVGICDistributorRegGICDIrouter237    HVGICDistributorReg = 0x6768
	HVGICDistributorRegGICDIrouter238    HVGICDistributorReg = 0x6770
	HVGICDistributorRegGICDIrouter239    HVGICDistributorReg = 0x6778
	HVGICDistributorRegGICDIrouter240    HVGICDistributorReg = 0x6780
	HVGICDistributorRegGICDIrouter241    HVGICDistributorReg = 0x6788
	HVGICDistributorRegGICDIrouter242    HVGICDistributorReg = 0x6790
	HVGICDistributorRegGICDIrouter243    HVGICDistributorReg = 0x6798
	HVGICDistributorRegGICDIrouter244    HVGICDistributorReg = 0x67a0
	HVGICDistributorRegGICDIrouter245    HVGICDistributorReg = 0x67a8
	HVGICDistributorRegGICDIrouter246    HVGICDistributorReg = 0x67b0
	HVGICDistributorRegGICDIrouter247    HVGICDistributorReg = 0x67b8
	HVGICDistributorRegGICDIrouter248    HVGICDistributorReg = 0x67c0
	HVGICDistributorRegGICDIrouter249    HVGICDistributorReg = 0x67c8
	HVGICDistributorRegGICDIrouter250    HVGICDistributorReg = 0x67d0
	HVGICDistributorRegGICDIrouter251    HVGICDistributorReg = 0x67d8
	HVGICDistributorRegGICDIrouter252    HVGICDistributorReg = 0x67e0
	HVGICDistributorRegGICDIrouter253    HVGICDistributorReg = 0x67e8
	HVGICDistributorRegGICDIrouter254    HVGICDistributorReg = 0x67f0
	HVGICDistributorRegGICDIrouter255    HVGICDistributorReg = 0x67f8
	HVGICDistributorRegGICDIrouter256    HVGICDistributorReg = 0x6800
	HVGICDistributorRegGICDIrouter257    HVGICDistributorReg = 0x6808
	HVGICDistributorRegGICDIrouter258    HVGICDistributorReg = 0x6810
	HVGICDistributorRegGICDIrouter259    HVGICDistributorReg = 0x6818
	HVGICDistributorRegGICDIrouter260    HVGICDistributorReg = 0x6820
	HVGICDistributorRegGICDIrouter261    HVGICDistributorReg = 0x6828
	HVGICDistributorRegGICDIrouter262    HVGICDistributorReg = 0x6830
	HVGICDistributorRegGICDIrouter263    HVGICDistributorReg = 0x6838
	HVGICDistributorRegGICDIrouter264    HVGICDistributorReg = 0x6840
	HVGICDistributorRegGICDIrouter265    HVGICDistributorReg = 0x6848
	HVGICDistributorRegGICDIrouter266    HVGICDistributorReg = 0x6850
	HVGICDistributorRegGICDIrouter267    HVGICDistributorReg = 0x6858
	HVGICDistributorRegGICDIrouter268    HVGICDistributorReg = 0x6860
	HVGICDistributorRegGICDIrouter269    HVGICDistributorReg = 0x6868
	HVGICDistributorRegGICDIrouter270    HVGICDistributorReg = 0x6870
	HVGICDistributorRegGICDIrouter271    HVGICDistributorReg = 0x6878
	HVGICDistributorRegGICDIrouter272    HVGICDistributorReg = 0x6880
	HVGICDistributorRegGICDIrouter273    HVGICDistributorReg = 0x6888
	HVGICDistributorRegGICDIrouter274    HVGICDistributorReg = 0x6890
	HVGICDistributorRegGICDIrouter275    HVGICDistributorReg = 0x6898
	HVGICDistributorRegGICDIrouter276    HVGICDistributorReg = 0x68a0
	HVGICDistributorRegGICDIrouter277    HVGICDistributorReg = 0x68a8
	HVGICDistributorRegGICDIrouter278    HVGICDistributorReg = 0x68b0
	HVGICDistributorRegGICDIrouter279    HVGICDistributorReg = 0x68b8
	HVGICDistributorRegGICDIrouter280    HVGICDistributorReg = 0x68c0
	HVGICDistributorRegGICDIrouter281    HVGICDistributorReg = 0x68c8
	HVGICDistributorRegGICDIrouter282    HVGICDistributorReg = 0x68d0
	HVGICDistributorRegGICDIrouter283    HVGICDistributorReg = 0x68d8
	HVGICDistributorRegGICDIrouter284    HVGICDistributorReg = 0x68e0
	HVGICDistributorRegGICDIrouter285    HVGICDistributorReg = 0x68e8
	HVGICDistributorRegGICDIrouter286    HVGICDistributorReg = 0x68f0
	HVGICDistributorRegGICDIrouter287    HVGICDistributorReg = 0x68f8
	HVGICDistributorRegGICDIrouter288    HVGICDistributorReg = 0x6900
	HVGICDistributorRegGICDIrouter289    HVGICDistributorReg = 0x6908
	HVGICDistributorRegGICDIrouter290    HVGICDistributorReg = 0x6910
	HVGICDistributorRegGICDIrouter291    HVGICDistributorReg = 0x6918
	HVGICDistributorRegGICDIrouter292    HVGICDistributorReg = 0x6920
	HVGICDistributorRegGICDIrouter293    HVGICDistributorReg = 0x6928
	HVGICDistributorRegGICDIrouter294    HVGICDistributorReg = 0x6930
	HVGICDistributorRegGICDIrouter295    HVGICDistributorReg = 0x6938
	HVGICDistributorRegGICDIrouter296    HVGICDistributorReg = 0x6940
	HVGICDistributorRegGICDIrouter297    HVGICDistributorReg = 0x6948
	HVGICDistributorRegGICDIrouter298    HVGICDistributorReg = 0x6950
	HVGICDistributorRegGICDIrouter299    HVGICDistributorReg = 0x6958
	HVGICDistributorRegGICDIrouter300    HVGICDistributorReg = 0x6960
	HVGICDistributorRegGICDIrouter301    HVGICDistributorReg = 0x6968
	HVGICDistributorRegGICDIrouter302    HVGICDistributorReg = 0x6970
	HVGICDistributorRegGICDIrouter303    HVGICDistributorReg = 0x6978
	HVGICDistributorRegGICDIrouter304    HVGICDistributorReg = 0x6980
	HVGICDistributorRegGICDIrouter305    HVGICDistributorReg = 0x6988
	HVGICDistributorRegGICDIrouter306    HVGICDistributorReg = 0x6990
	HVGICDistributorRegGICDIrouter307    HVGICDistributorReg = 0x6998
	HVGICDistributorRegGICDIrouter308    HVGICDistributorReg = 0x69a0
	HVGICDistributorRegGICDIrouter309    HVGICDistributorReg = 0x69a8
	HVGICDistributorRegGICDIrouter310    HVGICDistributorReg = 0x69b0
	HVGICDistributorRegGICDIrouter311    HVGICDistributorReg = 0x69b8
	HVGICDistributorRegGICDIrouter312    HVGICDistributorReg = 0x69c0
	HVGICDistributorRegGICDIrouter313    HVGICDistributorReg = 0x69c8
	HVGICDistributorRegGICDIrouter314    HVGICDistributorReg = 0x69d0
	HVGICDistributorRegGICDIrouter315    HVGICDistributorReg = 0x69d8
	HVGICDistributorRegGICDIrouter316    HVGICDistributorReg = 0x69e0
	HVGICDistributorRegGICDIrouter317    HVGICDistributorReg = 0x69e8
	HVGICDistributorRegGICDIrouter318    HVGICDistributorReg = 0x69f0
	HVGICDistributorRegGICDIrouter319    HVGICDistributorReg = 0x69f8
	HVGICDistributorRegGICDIrouter32     HVGICDistributorReg = 0x6100
	HVGICDistributorRegGICDIrouter320    HVGICDistributorReg = 0x6a00
	HVGICDistributorRegGICDIrouter321    HVGICDistributorReg = 0x6a08
	HVGICDistributorRegGICDIrouter322    HVGICDistributorReg = 0x6a10
	HVGICDistributorRegGICDIrouter323    HVGICDistributorReg = 0x6a18
	HVGICDistributorRegGICDIrouter324    HVGICDistributorReg = 0x6a20
	HVGICDistributorRegGICDIrouter325    HVGICDistributorReg = 0x6a28
	HVGICDistributorRegGICDIrouter326    HVGICDistributorReg = 0x6a30
	HVGICDistributorRegGICDIrouter327    HVGICDistributorReg = 0x6a38
	HVGICDistributorRegGICDIrouter328    HVGICDistributorReg = 0x6a40
	HVGICDistributorRegGICDIrouter329    HVGICDistributorReg = 0x6a48
	HVGICDistributorRegGICDIrouter33     HVGICDistributorReg = 0x6108
	HVGICDistributorRegGICDIrouter330    HVGICDistributorReg = 0x6a50
	HVGICDistributorRegGICDIrouter331    HVGICDistributorReg = 0x6a58
	HVGICDistributorRegGICDIrouter332    HVGICDistributorReg = 0x6a60
	HVGICDistributorRegGICDIrouter333    HVGICDistributorReg = 0x6a68
	HVGICDistributorRegGICDIrouter334    HVGICDistributorReg = 0x6a70
	HVGICDistributorRegGICDIrouter335    HVGICDistributorReg = 0x6a78
	HVGICDistributorRegGICDIrouter336    HVGICDistributorReg = 0x6a80
	HVGICDistributorRegGICDIrouter337    HVGICDistributorReg = 0x6a88
	HVGICDistributorRegGICDIrouter338    HVGICDistributorReg = 0x6a90
	HVGICDistributorRegGICDIrouter339    HVGICDistributorReg = 0x6a98
	HVGICDistributorRegGICDIrouter34     HVGICDistributorReg = 0x6110
	HVGICDistributorRegGICDIrouter340    HVGICDistributorReg = 0x6aa0
	HVGICDistributorRegGICDIrouter341    HVGICDistributorReg = 0x6aa8
	HVGICDistributorRegGICDIrouter342    HVGICDistributorReg = 0x6ab0
	HVGICDistributorRegGICDIrouter343    HVGICDistributorReg = 0x6ab8
	HVGICDistributorRegGICDIrouter344    HVGICDistributorReg = 0x6ac0
	HVGICDistributorRegGICDIrouter345    HVGICDistributorReg = 0x6ac8
	HVGICDistributorRegGICDIrouter346    HVGICDistributorReg = 0x6ad0
	HVGICDistributorRegGICDIrouter347    HVGICDistributorReg = 0x6ad8
	HVGICDistributorRegGICDIrouter348    HVGICDistributorReg = 0x6ae0
	HVGICDistributorRegGICDIrouter349    HVGICDistributorReg = 0x6ae8
	HVGICDistributorRegGICDIrouter35     HVGICDistributorReg = 0x6118
	HVGICDistributorRegGICDIrouter350    HVGICDistributorReg = 0x6af0
	HVGICDistributorRegGICDIrouter351    HVGICDistributorReg = 0x6af8
	HVGICDistributorRegGICDIrouter352    HVGICDistributorReg = 0x6b00
	HVGICDistributorRegGICDIrouter353    HVGICDistributorReg = 0x6b08
	HVGICDistributorRegGICDIrouter354    HVGICDistributorReg = 0x6b10
	HVGICDistributorRegGICDIrouter355    HVGICDistributorReg = 0x6b18
	HVGICDistributorRegGICDIrouter356    HVGICDistributorReg = 0x6b20
	HVGICDistributorRegGICDIrouter357    HVGICDistributorReg = 0x6b28
	HVGICDistributorRegGICDIrouter358    HVGICDistributorReg = 0x6b30
	HVGICDistributorRegGICDIrouter359    HVGICDistributorReg = 0x6b38
	HVGICDistributorRegGICDIrouter36     HVGICDistributorReg = 0x6120
	HVGICDistributorRegGICDIrouter360    HVGICDistributorReg = 0x6b40
	HVGICDistributorRegGICDIrouter361    HVGICDistributorReg = 0x6b48
	HVGICDistributorRegGICDIrouter362    HVGICDistributorReg = 0x6b50
	HVGICDistributorRegGICDIrouter363    HVGICDistributorReg = 0x6b58
	HVGICDistributorRegGICDIrouter364    HVGICDistributorReg = 0x6b60
	HVGICDistributorRegGICDIrouter365    HVGICDistributorReg = 0x6b68
	HVGICDistributorRegGICDIrouter366    HVGICDistributorReg = 0x6b70
	HVGICDistributorRegGICDIrouter367    HVGICDistributorReg = 0x6b78
	HVGICDistributorRegGICDIrouter368    HVGICDistributorReg = 0x6b80
	HVGICDistributorRegGICDIrouter369    HVGICDistributorReg = 0x6b88
	HVGICDistributorRegGICDIrouter37     HVGICDistributorReg = 0x6128
	HVGICDistributorRegGICDIrouter370    HVGICDistributorReg = 0x6b90
	HVGICDistributorRegGICDIrouter371    HVGICDistributorReg = 0x6b98
	HVGICDistributorRegGICDIrouter372    HVGICDistributorReg = 0x6ba0
	HVGICDistributorRegGICDIrouter373    HVGICDistributorReg = 0x6ba8
	HVGICDistributorRegGICDIrouter374    HVGICDistributorReg = 0x6bb0
	HVGICDistributorRegGICDIrouter375    HVGICDistributorReg = 0x6bb8
	HVGICDistributorRegGICDIrouter376    HVGICDistributorReg = 0x6bc0
	HVGICDistributorRegGICDIrouter377    HVGICDistributorReg = 0x6bc8
	HVGICDistributorRegGICDIrouter378    HVGICDistributorReg = 0x6bd0
	HVGICDistributorRegGICDIrouter379    HVGICDistributorReg = 0x6bd8
	HVGICDistributorRegGICDIrouter38     HVGICDistributorReg = 0x6130
	HVGICDistributorRegGICDIrouter380    HVGICDistributorReg = 0x6be0
	HVGICDistributorRegGICDIrouter381    HVGICDistributorReg = 0x6be8
	HVGICDistributorRegGICDIrouter382    HVGICDistributorReg = 0x6bf0
	HVGICDistributorRegGICDIrouter383    HVGICDistributorReg = 0x6bf8
	HVGICDistributorRegGICDIrouter384    HVGICDistributorReg = 0x6c00
	HVGICDistributorRegGICDIrouter385    HVGICDistributorReg = 0x6c08
	HVGICDistributorRegGICDIrouter386    HVGICDistributorReg = 0x6c10
	HVGICDistributorRegGICDIrouter387    HVGICDistributorReg = 0x6c18
	HVGICDistributorRegGICDIrouter388    HVGICDistributorReg = 0x6c20
	HVGICDistributorRegGICDIrouter389    HVGICDistributorReg = 0x6c28
	HVGICDistributorRegGICDIrouter39     HVGICDistributorReg = 0x6138
	HVGICDistributorRegGICDIrouter390    HVGICDistributorReg = 0x6c30
	HVGICDistributorRegGICDIrouter391    HVGICDistributorReg = 0x6c38
	HVGICDistributorRegGICDIrouter392    HVGICDistributorReg = 0x6c40
	HVGICDistributorRegGICDIrouter393    HVGICDistributorReg = 0x6c48
	HVGICDistributorRegGICDIrouter394    HVGICDistributorReg = 0x6c50
	HVGICDistributorRegGICDIrouter395    HVGICDistributorReg = 0x6c58
	HVGICDistributorRegGICDIrouter396    HVGICDistributorReg = 0x6c60
	HVGICDistributorRegGICDIrouter397    HVGICDistributorReg = 0x6c68
	HVGICDistributorRegGICDIrouter398    HVGICDistributorReg = 0x6c70
	HVGICDistributorRegGICDIrouter399    HVGICDistributorReg = 0x6c78
	HVGICDistributorRegGICDIrouter40     HVGICDistributorReg = 0x6140
	HVGICDistributorRegGICDIrouter400    HVGICDistributorReg = 0x6c80
	HVGICDistributorRegGICDIrouter401    HVGICDistributorReg = 0x6c88
	HVGICDistributorRegGICDIrouter402    HVGICDistributorReg = 0x6c90
	HVGICDistributorRegGICDIrouter403    HVGICDistributorReg = 0x6c98
	HVGICDistributorRegGICDIrouter404    HVGICDistributorReg = 0x6ca0
	HVGICDistributorRegGICDIrouter405    HVGICDistributorReg = 0x6ca8
	HVGICDistributorRegGICDIrouter406    HVGICDistributorReg = 0x6cb0
	HVGICDistributorRegGICDIrouter407    HVGICDistributorReg = 0x6cb8
	HVGICDistributorRegGICDIrouter408    HVGICDistributorReg = 0x6cc0
	HVGICDistributorRegGICDIrouter409    HVGICDistributorReg = 0x6cc8
	HVGICDistributorRegGICDIrouter41     HVGICDistributorReg = 0x6148
	HVGICDistributorRegGICDIrouter410    HVGICDistributorReg = 0x6cd0
	HVGICDistributorRegGICDIrouter411    HVGICDistributorReg = 0x6cd8
	HVGICDistributorRegGICDIrouter412    HVGICDistributorReg = 0x6ce0
	HVGICDistributorRegGICDIrouter413    HVGICDistributorReg = 0x6ce8
	HVGICDistributorRegGICDIrouter414    HVGICDistributorReg = 0x6cf0
	HVGICDistributorRegGICDIrouter415    HVGICDistributorReg = 0x6cf8
	HVGICDistributorRegGICDIrouter416    HVGICDistributorReg = 0x6d00
	HVGICDistributorRegGICDIrouter417    HVGICDistributorReg = 0x6d08
	HVGICDistributorRegGICDIrouter418    HVGICDistributorReg = 0x6d10
	HVGICDistributorRegGICDIrouter419    HVGICDistributorReg = 0x6d18
	HVGICDistributorRegGICDIrouter42     HVGICDistributorReg = 0x6150
	HVGICDistributorRegGICDIrouter420    HVGICDistributorReg = 0x6d20
	HVGICDistributorRegGICDIrouter421    HVGICDistributorReg = 0x6d28
	HVGICDistributorRegGICDIrouter422    HVGICDistributorReg = 0x6d30
	HVGICDistributorRegGICDIrouter423    HVGICDistributorReg = 0x6d38
	HVGICDistributorRegGICDIrouter424    HVGICDistributorReg = 0x6d40
	HVGICDistributorRegGICDIrouter425    HVGICDistributorReg = 0x6d48
	HVGICDistributorRegGICDIrouter426    HVGICDistributorReg = 0x6d50
	HVGICDistributorRegGICDIrouter427    HVGICDistributorReg = 0x6d58
	HVGICDistributorRegGICDIrouter428    HVGICDistributorReg = 0x6d60
	HVGICDistributorRegGICDIrouter429    HVGICDistributorReg = 0x6d68
	HVGICDistributorRegGICDIrouter43     HVGICDistributorReg = 0x6158
	HVGICDistributorRegGICDIrouter430    HVGICDistributorReg = 0x6d70
	HVGICDistributorRegGICDIrouter431    HVGICDistributorReg = 0x6d78
	HVGICDistributorRegGICDIrouter432    HVGICDistributorReg = 0x6d80
	HVGICDistributorRegGICDIrouter433    HVGICDistributorReg = 0x6d88
	HVGICDistributorRegGICDIrouter434    HVGICDistributorReg = 0x6d90
	HVGICDistributorRegGICDIrouter435    HVGICDistributorReg = 0x6d98
	HVGICDistributorRegGICDIrouter436    HVGICDistributorReg = 0x6da0
	HVGICDistributorRegGICDIrouter437    HVGICDistributorReg = 0x6da8
	HVGICDistributorRegGICDIrouter438    HVGICDistributorReg = 0x6db0
	HVGICDistributorRegGICDIrouter439    HVGICDistributorReg = 0x6db8
	HVGICDistributorRegGICDIrouter44     HVGICDistributorReg = 0x6160
	HVGICDistributorRegGICDIrouter440    HVGICDistributorReg = 0x6dc0
	HVGICDistributorRegGICDIrouter441    HVGICDistributorReg = 0x6dc8
	HVGICDistributorRegGICDIrouter442    HVGICDistributorReg = 0x6dd0
	HVGICDistributorRegGICDIrouter443    HVGICDistributorReg = 0x6dd8
	HVGICDistributorRegGICDIrouter444    HVGICDistributorReg = 0x6de0
	HVGICDistributorRegGICDIrouter445    HVGICDistributorReg = 0x6de8
	HVGICDistributorRegGICDIrouter446    HVGICDistributorReg = 0x6df0
	HVGICDistributorRegGICDIrouter447    HVGICDistributorReg = 0x6df8
	HVGICDistributorRegGICDIrouter448    HVGICDistributorReg = 0x6e00
	HVGICDistributorRegGICDIrouter449    HVGICDistributorReg = 0x6e08
	HVGICDistributorRegGICDIrouter45     HVGICDistributorReg = 0x6168
	HVGICDistributorRegGICDIrouter450    HVGICDistributorReg = 0x6e10
	HVGICDistributorRegGICDIrouter451    HVGICDistributorReg = 0x6e18
	HVGICDistributorRegGICDIrouter452    HVGICDistributorReg = 0x6e20
	HVGICDistributorRegGICDIrouter453    HVGICDistributorReg = 0x6e28
	HVGICDistributorRegGICDIrouter454    HVGICDistributorReg = 0x6e30
	HVGICDistributorRegGICDIrouter455    HVGICDistributorReg = 0x6e38
	HVGICDistributorRegGICDIrouter456    HVGICDistributorReg = 0x6e40
	HVGICDistributorRegGICDIrouter457    HVGICDistributorReg = 0x6e48
	HVGICDistributorRegGICDIrouter458    HVGICDistributorReg = 0x6e50
	HVGICDistributorRegGICDIrouter459    HVGICDistributorReg = 0x6e58
	HVGICDistributorRegGICDIrouter46     HVGICDistributorReg = 0x6170
	HVGICDistributorRegGICDIrouter460    HVGICDistributorReg = 0x6e60
	HVGICDistributorRegGICDIrouter461    HVGICDistributorReg = 0x6e68
	HVGICDistributorRegGICDIrouter462    HVGICDistributorReg = 0x6e70
	HVGICDistributorRegGICDIrouter463    HVGICDistributorReg = 0x6e78
	HVGICDistributorRegGICDIrouter464    HVGICDistributorReg = 0x6e80
	HVGICDistributorRegGICDIrouter465    HVGICDistributorReg = 0x6e88
	HVGICDistributorRegGICDIrouter466    HVGICDistributorReg = 0x6e90
	HVGICDistributorRegGICDIrouter467    HVGICDistributorReg = 0x6e98
	HVGICDistributorRegGICDIrouter468    HVGICDistributorReg = 0x6ea0
	HVGICDistributorRegGICDIrouter469    HVGICDistributorReg = 0x6ea8
	HVGICDistributorRegGICDIrouter47     HVGICDistributorReg = 0x6178
	HVGICDistributorRegGICDIrouter470    HVGICDistributorReg = 0x6eb0
	HVGICDistributorRegGICDIrouter471    HVGICDistributorReg = 0x6eb8
	HVGICDistributorRegGICDIrouter472    HVGICDistributorReg = 0x6ec0
	HVGICDistributorRegGICDIrouter473    HVGICDistributorReg = 0x6ec8
	HVGICDistributorRegGICDIrouter474    HVGICDistributorReg = 0x6ed0
	HVGICDistributorRegGICDIrouter475    HVGICDistributorReg = 0x6ed8
	HVGICDistributorRegGICDIrouter476    HVGICDistributorReg = 0x6ee0
	HVGICDistributorRegGICDIrouter477    HVGICDistributorReg = 0x6ee8
	HVGICDistributorRegGICDIrouter478    HVGICDistributorReg = 0x6ef0
	HVGICDistributorRegGICDIrouter479    HVGICDistributorReg = 0x6ef8
	HVGICDistributorRegGICDIrouter48     HVGICDistributorReg = 0x6180
	HVGICDistributorRegGICDIrouter480    HVGICDistributorReg = 0x6f00
	HVGICDistributorRegGICDIrouter481    HVGICDistributorReg = 0x6f08
	HVGICDistributorRegGICDIrouter482    HVGICDistributorReg = 0x6f10
	HVGICDistributorRegGICDIrouter483    HVGICDistributorReg = 0x6f18
	HVGICDistributorRegGICDIrouter484    HVGICDistributorReg = 0x6f20
	HVGICDistributorRegGICDIrouter485    HVGICDistributorReg = 0x6f28
	HVGICDistributorRegGICDIrouter486    HVGICDistributorReg = 0x6f30
	HVGICDistributorRegGICDIrouter487    HVGICDistributorReg = 0x6f38
	HVGICDistributorRegGICDIrouter488    HVGICDistributorReg = 0x6f40
	HVGICDistributorRegGICDIrouter489    HVGICDistributorReg = 0x6f48
	HVGICDistributorRegGICDIrouter49     HVGICDistributorReg = 0x6188
	HVGICDistributorRegGICDIrouter490    HVGICDistributorReg = 0x6f50
	HVGICDistributorRegGICDIrouter491    HVGICDistributorReg = 0x6f58
	HVGICDistributorRegGICDIrouter492    HVGICDistributorReg = 0x6f60
	HVGICDistributorRegGICDIrouter493    HVGICDistributorReg = 0x6f68
	HVGICDistributorRegGICDIrouter494    HVGICDistributorReg = 0x6f70
	HVGICDistributorRegGICDIrouter495    HVGICDistributorReg = 0x6f78
	HVGICDistributorRegGICDIrouter496    HVGICDistributorReg = 0x6f80
	HVGICDistributorRegGICDIrouter497    HVGICDistributorReg = 0x6f88
	HVGICDistributorRegGICDIrouter498    HVGICDistributorReg = 0x6f90
	HVGICDistributorRegGICDIrouter499    HVGICDistributorReg = 0x6f98
	HVGICDistributorRegGICDIrouter50     HVGICDistributorReg = 0x6190
	HVGICDistributorRegGICDIrouter500    HVGICDistributorReg = 0x6fa0
	HVGICDistributorRegGICDIrouter501    HVGICDistributorReg = 0x6fa8
	HVGICDistributorRegGICDIrouter502    HVGICDistributorReg = 0x6fb0
	HVGICDistributorRegGICDIrouter503    HVGICDistributorReg = 0x6fb8
	HVGICDistributorRegGICDIrouter504    HVGICDistributorReg = 0x6fc0
	HVGICDistributorRegGICDIrouter505    HVGICDistributorReg = 0x6fc8
	HVGICDistributorRegGICDIrouter506    HVGICDistributorReg = 0x6fd0
	HVGICDistributorRegGICDIrouter507    HVGICDistributorReg = 0x6fd8
	HVGICDistributorRegGICDIrouter508    HVGICDistributorReg = 0x6fe0
	HVGICDistributorRegGICDIrouter509    HVGICDistributorReg = 0x6fe8
	HVGICDistributorRegGICDIrouter51     HVGICDistributorReg = 0x6198
	HVGICDistributorRegGICDIrouter510    HVGICDistributorReg = 0x6ff0
	HVGICDistributorRegGICDIrouter511    HVGICDistributorReg = 0x6ff8
	HVGICDistributorRegGICDIrouter512    HVGICDistributorReg = 0x7000
	HVGICDistributorRegGICDIrouter513    HVGICDistributorReg = 0x7008
	HVGICDistributorRegGICDIrouter514    HVGICDistributorReg = 0x7010
	HVGICDistributorRegGICDIrouter515    HVGICDistributorReg = 0x7018
	HVGICDistributorRegGICDIrouter516    HVGICDistributorReg = 0x7020
	HVGICDistributorRegGICDIrouter517    HVGICDistributorReg = 0x7028
	HVGICDistributorRegGICDIrouter518    HVGICDistributorReg = 0x7030
	HVGICDistributorRegGICDIrouter519    HVGICDistributorReg = 0x7038
	HVGICDistributorRegGICDIrouter52     HVGICDistributorReg = 0x61a0
	HVGICDistributorRegGICDIrouter520    HVGICDistributorReg = 0x7040
	HVGICDistributorRegGICDIrouter521    HVGICDistributorReg = 0x7048
	HVGICDistributorRegGICDIrouter522    HVGICDistributorReg = 0x7050
	HVGICDistributorRegGICDIrouter523    HVGICDistributorReg = 0x7058
	HVGICDistributorRegGICDIrouter524    HVGICDistributorReg = 0x7060
	HVGICDistributorRegGICDIrouter525    HVGICDistributorReg = 0x7068
	HVGICDistributorRegGICDIrouter526    HVGICDistributorReg = 0x7070
	HVGICDistributorRegGICDIrouter527    HVGICDistributorReg = 0x7078
	HVGICDistributorRegGICDIrouter528    HVGICDistributorReg = 0x7080
	HVGICDistributorRegGICDIrouter529    HVGICDistributorReg = 0x7088
	HVGICDistributorRegGICDIrouter53     HVGICDistributorReg = 0x61a8
	HVGICDistributorRegGICDIrouter530    HVGICDistributorReg = 0x7090
	HVGICDistributorRegGICDIrouter531    HVGICDistributorReg = 0x7098
	HVGICDistributorRegGICDIrouter532    HVGICDistributorReg = 0x70a0
	HVGICDistributorRegGICDIrouter533    HVGICDistributorReg = 0x70a8
	HVGICDistributorRegGICDIrouter534    HVGICDistributorReg = 0x70b0
	HVGICDistributorRegGICDIrouter535    HVGICDistributorReg = 0x70b8
	HVGICDistributorRegGICDIrouter536    HVGICDistributorReg = 0x70c0
	HVGICDistributorRegGICDIrouter537    HVGICDistributorReg = 0x70c8
	HVGICDistributorRegGICDIrouter538    HVGICDistributorReg = 0x70d0
	HVGICDistributorRegGICDIrouter539    HVGICDistributorReg = 0x70d8
	HVGICDistributorRegGICDIrouter54     HVGICDistributorReg = 0x61b0
	HVGICDistributorRegGICDIrouter540    HVGICDistributorReg = 0x70e0
	HVGICDistributorRegGICDIrouter541    HVGICDistributorReg = 0x70e8
	HVGICDistributorRegGICDIrouter542    HVGICDistributorReg = 0x70f0
	HVGICDistributorRegGICDIrouter543    HVGICDistributorReg = 0x70f8
	HVGICDistributorRegGICDIrouter544    HVGICDistributorReg = 0x7100
	HVGICDistributorRegGICDIrouter545    HVGICDistributorReg = 0x7108
	HVGICDistributorRegGICDIrouter546    HVGICDistributorReg = 0x7110
	HVGICDistributorRegGICDIrouter547    HVGICDistributorReg = 0x7118
	HVGICDistributorRegGICDIrouter548    HVGICDistributorReg = 0x7120
	HVGICDistributorRegGICDIrouter549    HVGICDistributorReg = 0x7128
	HVGICDistributorRegGICDIrouter55     HVGICDistributorReg = 0x61b8
	HVGICDistributorRegGICDIrouter550    HVGICDistributorReg = 0x7130
	HVGICDistributorRegGICDIrouter551    HVGICDistributorReg = 0x7138
	HVGICDistributorRegGICDIrouter552    HVGICDistributorReg = 0x7140
	HVGICDistributorRegGICDIrouter553    HVGICDistributorReg = 0x7148
	HVGICDistributorRegGICDIrouter554    HVGICDistributorReg = 0x7150
	HVGICDistributorRegGICDIrouter555    HVGICDistributorReg = 0x7158
	HVGICDistributorRegGICDIrouter556    HVGICDistributorReg = 0x7160
	HVGICDistributorRegGICDIrouter557    HVGICDistributorReg = 0x7168
	HVGICDistributorRegGICDIrouter558    HVGICDistributorReg = 0x7170
	HVGICDistributorRegGICDIrouter559    HVGICDistributorReg = 0x7178
	HVGICDistributorRegGICDIrouter56     HVGICDistributorReg = 0x61c0
	HVGICDistributorRegGICDIrouter560    HVGICDistributorReg = 0x7180
	HVGICDistributorRegGICDIrouter561    HVGICDistributorReg = 0x7188
	HVGICDistributorRegGICDIrouter562    HVGICDistributorReg = 0x7190
	HVGICDistributorRegGICDIrouter563    HVGICDistributorReg = 0x7198
	HVGICDistributorRegGICDIrouter564    HVGICDistributorReg = 0x71a0
	HVGICDistributorRegGICDIrouter565    HVGICDistributorReg = 0x71a8
	HVGICDistributorRegGICDIrouter566    HVGICDistributorReg = 0x71b0
	HVGICDistributorRegGICDIrouter567    HVGICDistributorReg = 0x71b8
	HVGICDistributorRegGICDIrouter568    HVGICDistributorReg = 0x71c0
	HVGICDistributorRegGICDIrouter569    HVGICDistributorReg = 0x71c8
	HVGICDistributorRegGICDIrouter57     HVGICDistributorReg = 0x61c8
	HVGICDistributorRegGICDIrouter570    HVGICDistributorReg = 0x71d0
	HVGICDistributorRegGICDIrouter571    HVGICDistributorReg = 0x71d8
	HVGICDistributorRegGICDIrouter572    HVGICDistributorReg = 0x71e0
	HVGICDistributorRegGICDIrouter573    HVGICDistributorReg = 0x71e8
	HVGICDistributorRegGICDIrouter574    HVGICDistributorReg = 0x71f0
	HVGICDistributorRegGICDIrouter575    HVGICDistributorReg = 0x71f8
	HVGICDistributorRegGICDIrouter576    HVGICDistributorReg = 0x7200
	HVGICDistributorRegGICDIrouter577    HVGICDistributorReg = 0x7208
	HVGICDistributorRegGICDIrouter578    HVGICDistributorReg = 0x7210
	HVGICDistributorRegGICDIrouter579    HVGICDistributorReg = 0x7218
	HVGICDistributorRegGICDIrouter58     HVGICDistributorReg = 0x61d0
	HVGICDistributorRegGICDIrouter580    HVGICDistributorReg = 0x7220
	HVGICDistributorRegGICDIrouter581    HVGICDistributorReg = 0x7228
	HVGICDistributorRegGICDIrouter582    HVGICDistributorReg = 0x7230
	HVGICDistributorRegGICDIrouter583    HVGICDistributorReg = 0x7238
	HVGICDistributorRegGICDIrouter584    HVGICDistributorReg = 0x7240
	HVGICDistributorRegGICDIrouter585    HVGICDistributorReg = 0x7248
	HVGICDistributorRegGICDIrouter586    HVGICDistributorReg = 0x7250
	HVGICDistributorRegGICDIrouter587    HVGICDistributorReg = 0x7258
	HVGICDistributorRegGICDIrouter588    HVGICDistributorReg = 0x7260
	HVGICDistributorRegGICDIrouter589    HVGICDistributorReg = 0x7268
	HVGICDistributorRegGICDIrouter59     HVGICDistributorReg = 0x61d8
	HVGICDistributorRegGICDIrouter590    HVGICDistributorReg = 0x7270
	HVGICDistributorRegGICDIrouter591    HVGICDistributorReg = 0x7278
	HVGICDistributorRegGICDIrouter592    HVGICDistributorReg = 0x7280
	HVGICDistributorRegGICDIrouter593    HVGICDistributorReg = 0x7288
	HVGICDistributorRegGICDIrouter594    HVGICDistributorReg = 0x7290
	HVGICDistributorRegGICDIrouter595    HVGICDistributorReg = 0x7298
	HVGICDistributorRegGICDIrouter596    HVGICDistributorReg = 0x72a0
	HVGICDistributorRegGICDIrouter597    HVGICDistributorReg = 0x72a8
	HVGICDistributorRegGICDIrouter598    HVGICDistributorReg = 0x72b0
	HVGICDistributorRegGICDIrouter599    HVGICDistributorReg = 0x72b8
	HVGICDistributorRegGICDIrouter60     HVGICDistributorReg = 0x61e0
	HVGICDistributorRegGICDIrouter600    HVGICDistributorReg = 0x72c0
	HVGICDistributorRegGICDIrouter601    HVGICDistributorReg = 0x72c8
	HVGICDistributorRegGICDIrouter602    HVGICDistributorReg = 0x72d0
	HVGICDistributorRegGICDIrouter603    HVGICDistributorReg = 0x72d8
	HVGICDistributorRegGICDIrouter604    HVGICDistributorReg = 0x72e0
	HVGICDistributorRegGICDIrouter605    HVGICDistributorReg = 0x72e8
	HVGICDistributorRegGICDIrouter606    HVGICDistributorReg = 0x72f0
	HVGICDistributorRegGICDIrouter607    HVGICDistributorReg = 0x72f8
	HVGICDistributorRegGICDIrouter608    HVGICDistributorReg = 0x7300
	HVGICDistributorRegGICDIrouter609    HVGICDistributorReg = 0x7308
	HVGICDistributorRegGICDIrouter61     HVGICDistributorReg = 0x61e8
	HVGICDistributorRegGICDIrouter610    HVGICDistributorReg = 0x7310
	HVGICDistributorRegGICDIrouter611    HVGICDistributorReg = 0x7318
	HVGICDistributorRegGICDIrouter612    HVGICDistributorReg = 0x7320
	HVGICDistributorRegGICDIrouter613    HVGICDistributorReg = 0x7328
	HVGICDistributorRegGICDIrouter614    HVGICDistributorReg = 0x7330
	HVGICDistributorRegGICDIrouter615    HVGICDistributorReg = 0x7338
	HVGICDistributorRegGICDIrouter616    HVGICDistributorReg = 0x7340
	HVGICDistributorRegGICDIrouter617    HVGICDistributorReg = 0x7348
	HVGICDistributorRegGICDIrouter618    HVGICDistributorReg = 0x7350
	HVGICDistributorRegGICDIrouter619    HVGICDistributorReg = 0x7358
	HVGICDistributorRegGICDIrouter62     HVGICDistributorReg = 0x61f0
	HVGICDistributorRegGICDIrouter620    HVGICDistributorReg = 0x7360
	HVGICDistributorRegGICDIrouter621    HVGICDistributorReg = 0x7368
	HVGICDistributorRegGICDIrouter622    HVGICDistributorReg = 0x7370
	HVGICDistributorRegGICDIrouter623    HVGICDistributorReg = 0x7378
	HVGICDistributorRegGICDIrouter624    HVGICDistributorReg = 0x7380
	HVGICDistributorRegGICDIrouter625    HVGICDistributorReg = 0x7388
	HVGICDistributorRegGICDIrouter626    HVGICDistributorReg = 0x7390
	HVGICDistributorRegGICDIrouter627    HVGICDistributorReg = 0x7398
	HVGICDistributorRegGICDIrouter628    HVGICDistributorReg = 0x73a0
	HVGICDistributorRegGICDIrouter629    HVGICDistributorReg = 0x73a8
	HVGICDistributorRegGICDIrouter63     HVGICDistributorReg = 0x61f8
	HVGICDistributorRegGICDIrouter630    HVGICDistributorReg = 0x73b0
	HVGICDistributorRegGICDIrouter631    HVGICDistributorReg = 0x73b8
	HVGICDistributorRegGICDIrouter632    HVGICDistributorReg = 0x73c0
	HVGICDistributorRegGICDIrouter633    HVGICDistributorReg = 0x73c8
	HVGICDistributorRegGICDIrouter634    HVGICDistributorReg = 0x73d0
	HVGICDistributorRegGICDIrouter635    HVGICDistributorReg = 0x73d8
	HVGICDistributorRegGICDIrouter636    HVGICDistributorReg = 0x73e0
	HVGICDistributorRegGICDIrouter637    HVGICDistributorReg = 0x73e8
	HVGICDistributorRegGICDIrouter638    HVGICDistributorReg = 0x73f0
	HVGICDistributorRegGICDIrouter639    HVGICDistributorReg = 0x73f8
	HVGICDistributorRegGICDIrouter64     HVGICDistributorReg = 0x6200
	HVGICDistributorRegGICDIrouter640    HVGICDistributorReg = 0x7400
	HVGICDistributorRegGICDIrouter641    HVGICDistributorReg = 0x7408
	HVGICDistributorRegGICDIrouter642    HVGICDistributorReg = 0x7410
	HVGICDistributorRegGICDIrouter643    HVGICDistributorReg = 0x7418
	HVGICDistributorRegGICDIrouter644    HVGICDistributorReg = 0x7420
	HVGICDistributorRegGICDIrouter645    HVGICDistributorReg = 0x7428
	HVGICDistributorRegGICDIrouter646    HVGICDistributorReg = 0x7430
	HVGICDistributorRegGICDIrouter647    HVGICDistributorReg = 0x7438
	HVGICDistributorRegGICDIrouter648    HVGICDistributorReg = 0x7440
	HVGICDistributorRegGICDIrouter649    HVGICDistributorReg = 0x7448
	HVGICDistributorRegGICDIrouter65     HVGICDistributorReg = 0x6208
	HVGICDistributorRegGICDIrouter650    HVGICDistributorReg = 0x7450
	HVGICDistributorRegGICDIrouter651    HVGICDistributorReg = 0x7458
	HVGICDistributorRegGICDIrouter652    HVGICDistributorReg = 0x7460
	HVGICDistributorRegGICDIrouter653    HVGICDistributorReg = 0x7468
	HVGICDistributorRegGICDIrouter654    HVGICDistributorReg = 0x7470
	HVGICDistributorRegGICDIrouter655    HVGICDistributorReg = 0x7478
	HVGICDistributorRegGICDIrouter656    HVGICDistributorReg = 0x7480
	HVGICDistributorRegGICDIrouter657    HVGICDistributorReg = 0x7488
	HVGICDistributorRegGICDIrouter658    HVGICDistributorReg = 0x7490
	HVGICDistributorRegGICDIrouter659    HVGICDistributorReg = 0x7498
	HVGICDistributorRegGICDIrouter66     HVGICDistributorReg = 0x6210
	HVGICDistributorRegGICDIrouter660    HVGICDistributorReg = 0x74a0
	HVGICDistributorRegGICDIrouter661    HVGICDistributorReg = 0x74a8
	HVGICDistributorRegGICDIrouter662    HVGICDistributorReg = 0x74b0
	HVGICDistributorRegGICDIrouter663    HVGICDistributorReg = 0x74b8
	HVGICDistributorRegGICDIrouter664    HVGICDistributorReg = 0x74c0
	HVGICDistributorRegGICDIrouter665    HVGICDistributorReg = 0x74c8
	HVGICDistributorRegGICDIrouter666    HVGICDistributorReg = 0x74d0
	HVGICDistributorRegGICDIrouter667    HVGICDistributorReg = 0x74d8
	HVGICDistributorRegGICDIrouter668    HVGICDistributorReg = 0x74e0
	HVGICDistributorRegGICDIrouter669    HVGICDistributorReg = 0x74e8
	HVGICDistributorRegGICDIrouter67     HVGICDistributorReg = 0x6218
	HVGICDistributorRegGICDIrouter670    HVGICDistributorReg = 0x74f0
	HVGICDistributorRegGICDIrouter671    HVGICDistributorReg = 0x74f8
	HVGICDistributorRegGICDIrouter672    HVGICDistributorReg = 0x7500
	HVGICDistributorRegGICDIrouter673    HVGICDistributorReg = 0x7508
	HVGICDistributorRegGICDIrouter674    HVGICDistributorReg = 0x7510
	HVGICDistributorRegGICDIrouter675    HVGICDistributorReg = 0x7518
	HVGICDistributorRegGICDIrouter676    HVGICDistributorReg = 0x7520
	HVGICDistributorRegGICDIrouter677    HVGICDistributorReg = 0x7528
	HVGICDistributorRegGICDIrouter678    HVGICDistributorReg = 0x7530
	HVGICDistributorRegGICDIrouter679    HVGICDistributorReg = 0x7538
	HVGICDistributorRegGICDIrouter68     HVGICDistributorReg = 0x6220
	HVGICDistributorRegGICDIrouter680    HVGICDistributorReg = 0x7540
	HVGICDistributorRegGICDIrouter681    HVGICDistributorReg = 0x7548
	HVGICDistributorRegGICDIrouter682    HVGICDistributorReg = 0x7550
	HVGICDistributorRegGICDIrouter683    HVGICDistributorReg = 0x7558
	HVGICDistributorRegGICDIrouter684    HVGICDistributorReg = 0x7560
	HVGICDistributorRegGICDIrouter685    HVGICDistributorReg = 0x7568
	HVGICDistributorRegGICDIrouter686    HVGICDistributorReg = 0x7570
	HVGICDistributorRegGICDIrouter687    HVGICDistributorReg = 0x7578
	HVGICDistributorRegGICDIrouter688    HVGICDistributorReg = 0x7580
	HVGICDistributorRegGICDIrouter689    HVGICDistributorReg = 0x7588
	HVGICDistributorRegGICDIrouter69     HVGICDistributorReg = 0x6228
	HVGICDistributorRegGICDIrouter690    HVGICDistributorReg = 0x7590
	HVGICDistributorRegGICDIrouter691    HVGICDistributorReg = 0x7598
	HVGICDistributorRegGICDIrouter692    HVGICDistributorReg = 0x75a0
	HVGICDistributorRegGICDIrouter693    HVGICDistributorReg = 0x75a8
	HVGICDistributorRegGICDIrouter694    HVGICDistributorReg = 0x75b0
	HVGICDistributorRegGICDIrouter695    HVGICDistributorReg = 0x75b8
	HVGICDistributorRegGICDIrouter696    HVGICDistributorReg = 0x75c0
	HVGICDistributorRegGICDIrouter697    HVGICDistributorReg = 0x75c8
	HVGICDistributorRegGICDIrouter698    HVGICDistributorReg = 0x75d0
	HVGICDistributorRegGICDIrouter699    HVGICDistributorReg = 0x75d8
	HVGICDistributorRegGICDIrouter70     HVGICDistributorReg = 0x6230
	HVGICDistributorRegGICDIrouter700    HVGICDistributorReg = 0x75e0
	HVGICDistributorRegGICDIrouter701    HVGICDistributorReg = 0x75e8
	HVGICDistributorRegGICDIrouter702    HVGICDistributorReg = 0x75f0
	HVGICDistributorRegGICDIrouter703    HVGICDistributorReg = 0x75f8
	HVGICDistributorRegGICDIrouter704    HVGICDistributorReg = 0x7600
	HVGICDistributorRegGICDIrouter705    HVGICDistributorReg = 0x7608
	HVGICDistributorRegGICDIrouter706    HVGICDistributorReg = 0x7610
	HVGICDistributorRegGICDIrouter707    HVGICDistributorReg = 0x7618
	HVGICDistributorRegGICDIrouter708    HVGICDistributorReg = 0x7620
	HVGICDistributorRegGICDIrouter709    HVGICDistributorReg = 0x7628
	HVGICDistributorRegGICDIrouter71     HVGICDistributorReg = 0x6238
	HVGICDistributorRegGICDIrouter710    HVGICDistributorReg = 0x7630
	HVGICDistributorRegGICDIrouter711    HVGICDistributorReg = 0x7638
	HVGICDistributorRegGICDIrouter712    HVGICDistributorReg = 0x7640
	HVGICDistributorRegGICDIrouter713    HVGICDistributorReg = 0x7648
	HVGICDistributorRegGICDIrouter714    HVGICDistributorReg = 0x7650
	HVGICDistributorRegGICDIrouter715    HVGICDistributorReg = 0x7658
	HVGICDistributorRegGICDIrouter716    HVGICDistributorReg = 0x7660
	HVGICDistributorRegGICDIrouter717    HVGICDistributorReg = 0x7668
	HVGICDistributorRegGICDIrouter718    HVGICDistributorReg = 0x7670
	HVGICDistributorRegGICDIrouter719    HVGICDistributorReg = 0x7678
	HVGICDistributorRegGICDIrouter72     HVGICDistributorReg = 0x6240
	HVGICDistributorRegGICDIrouter720    HVGICDistributorReg = 0x7680
	HVGICDistributorRegGICDIrouter721    HVGICDistributorReg = 0x7688
	HVGICDistributorRegGICDIrouter722    HVGICDistributorReg = 0x7690
	HVGICDistributorRegGICDIrouter723    HVGICDistributorReg = 0x7698
	HVGICDistributorRegGICDIrouter724    HVGICDistributorReg = 0x76a0
	HVGICDistributorRegGICDIrouter725    HVGICDistributorReg = 0x76a8
	HVGICDistributorRegGICDIrouter726    HVGICDistributorReg = 0x76b0
	HVGICDistributorRegGICDIrouter727    HVGICDistributorReg = 0x76b8
	HVGICDistributorRegGICDIrouter728    HVGICDistributorReg = 0x76c0
	HVGICDistributorRegGICDIrouter729    HVGICDistributorReg = 0x76c8
	HVGICDistributorRegGICDIrouter73     HVGICDistributorReg = 0x6248
	HVGICDistributorRegGICDIrouter730    HVGICDistributorReg = 0x76d0
	HVGICDistributorRegGICDIrouter731    HVGICDistributorReg = 0x76d8
	HVGICDistributorRegGICDIrouter732    HVGICDistributorReg = 0x76e0
	HVGICDistributorRegGICDIrouter733    HVGICDistributorReg = 0x76e8
	HVGICDistributorRegGICDIrouter734    HVGICDistributorReg = 0x76f0
	HVGICDistributorRegGICDIrouter735    HVGICDistributorReg = 0x76f8
	HVGICDistributorRegGICDIrouter736    HVGICDistributorReg = 0x7700
	HVGICDistributorRegGICDIrouter737    HVGICDistributorReg = 0x7708
	HVGICDistributorRegGICDIrouter738    HVGICDistributorReg = 0x7710
	HVGICDistributorRegGICDIrouter739    HVGICDistributorReg = 0x7718
	HVGICDistributorRegGICDIrouter74     HVGICDistributorReg = 0x6250
	HVGICDistributorRegGICDIrouter740    HVGICDistributorReg = 0x7720
	HVGICDistributorRegGICDIrouter741    HVGICDistributorReg = 0x7728
	HVGICDistributorRegGICDIrouter742    HVGICDistributorReg = 0x7730
	HVGICDistributorRegGICDIrouter743    HVGICDistributorReg = 0x7738
	HVGICDistributorRegGICDIrouter744    HVGICDistributorReg = 0x7740
	HVGICDistributorRegGICDIrouter745    HVGICDistributorReg = 0x7748
	HVGICDistributorRegGICDIrouter746    HVGICDistributorReg = 0x7750
	HVGICDistributorRegGICDIrouter747    HVGICDistributorReg = 0x7758
	HVGICDistributorRegGICDIrouter748    HVGICDistributorReg = 0x7760
	HVGICDistributorRegGICDIrouter749    HVGICDistributorReg = 0x7768
	HVGICDistributorRegGICDIrouter75     HVGICDistributorReg = 0x6258
	HVGICDistributorRegGICDIrouter750    HVGICDistributorReg = 0x7770
	HVGICDistributorRegGICDIrouter751    HVGICDistributorReg = 0x7778
	HVGICDistributorRegGICDIrouter752    HVGICDistributorReg = 0x7780
	HVGICDistributorRegGICDIrouter753    HVGICDistributorReg = 0x7788
	HVGICDistributorRegGICDIrouter754    HVGICDistributorReg = 0x7790
	HVGICDistributorRegGICDIrouter755    HVGICDistributorReg = 0x7798
	HVGICDistributorRegGICDIrouter756    HVGICDistributorReg = 0x77a0
	HVGICDistributorRegGICDIrouter757    HVGICDistributorReg = 0x77a8
	HVGICDistributorRegGICDIrouter758    HVGICDistributorReg = 0x77b0
	HVGICDistributorRegGICDIrouter759    HVGICDistributorReg = 0x77b8
	HVGICDistributorRegGICDIrouter76     HVGICDistributorReg = 0x6260
	HVGICDistributorRegGICDIrouter760    HVGICDistributorReg = 0x77c0
	HVGICDistributorRegGICDIrouter761    HVGICDistributorReg = 0x77c8
	HVGICDistributorRegGICDIrouter762    HVGICDistributorReg = 0x77d0
	HVGICDistributorRegGICDIrouter763    HVGICDistributorReg = 0x77d8
	HVGICDistributorRegGICDIrouter764    HVGICDistributorReg = 0x77e0
	HVGICDistributorRegGICDIrouter765    HVGICDistributorReg = 0x77e8
	HVGICDistributorRegGICDIrouter766    HVGICDistributorReg = 0x77f0
	HVGICDistributorRegGICDIrouter767    HVGICDistributorReg = 0x77f8
	HVGICDistributorRegGICDIrouter768    HVGICDistributorReg = 0x7800
	HVGICDistributorRegGICDIrouter769    HVGICDistributorReg = 0x7808
	HVGICDistributorRegGICDIrouter77     HVGICDistributorReg = 0x6268
	HVGICDistributorRegGICDIrouter770    HVGICDistributorReg = 0x7810
	HVGICDistributorRegGICDIrouter771    HVGICDistributorReg = 0x7818
	HVGICDistributorRegGICDIrouter772    HVGICDistributorReg = 0x7820
	HVGICDistributorRegGICDIrouter773    HVGICDistributorReg = 0x7828
	HVGICDistributorRegGICDIrouter774    HVGICDistributorReg = 0x7830
	HVGICDistributorRegGICDIrouter775    HVGICDistributorReg = 0x7838
	HVGICDistributorRegGICDIrouter776    HVGICDistributorReg = 0x7840
	HVGICDistributorRegGICDIrouter777    HVGICDistributorReg = 0x7848
	HVGICDistributorRegGICDIrouter778    HVGICDistributorReg = 0x7850
	HVGICDistributorRegGICDIrouter779    HVGICDistributorReg = 0x7858
	HVGICDistributorRegGICDIrouter78     HVGICDistributorReg = 0x6270
	HVGICDistributorRegGICDIrouter780    HVGICDistributorReg = 0x7860
	HVGICDistributorRegGICDIrouter781    HVGICDistributorReg = 0x7868
	HVGICDistributorRegGICDIrouter782    HVGICDistributorReg = 0x7870
	HVGICDistributorRegGICDIrouter783    HVGICDistributorReg = 0x7878
	HVGICDistributorRegGICDIrouter784    HVGICDistributorReg = 0x7880
	HVGICDistributorRegGICDIrouter785    HVGICDistributorReg = 0x7888
	HVGICDistributorRegGICDIrouter786    HVGICDistributorReg = 0x7890
	HVGICDistributorRegGICDIrouter787    HVGICDistributorReg = 0x7898
	HVGICDistributorRegGICDIrouter788    HVGICDistributorReg = 0x78a0
	HVGICDistributorRegGICDIrouter789    HVGICDistributorReg = 0x78a8
	HVGICDistributorRegGICDIrouter79     HVGICDistributorReg = 0x6278
	HVGICDistributorRegGICDIrouter790    HVGICDistributorReg = 0x78b0
	HVGICDistributorRegGICDIrouter791    HVGICDistributorReg = 0x78b8
	HVGICDistributorRegGICDIrouter792    HVGICDistributorReg = 0x78c0
	HVGICDistributorRegGICDIrouter793    HVGICDistributorReg = 0x78c8
	HVGICDistributorRegGICDIrouter794    HVGICDistributorReg = 0x78d0
	HVGICDistributorRegGICDIrouter795    HVGICDistributorReg = 0x78d8
	HVGICDistributorRegGICDIrouter796    HVGICDistributorReg = 0x78e0
	HVGICDistributorRegGICDIrouter797    HVGICDistributorReg = 0x78e8
	HVGICDistributorRegGICDIrouter798    HVGICDistributorReg = 0x78f0
	HVGICDistributorRegGICDIrouter799    HVGICDistributorReg = 0x78f8
	HVGICDistributorRegGICDIrouter80     HVGICDistributorReg = 0x6280
	HVGICDistributorRegGICDIrouter800    HVGICDistributorReg = 0x7900
	HVGICDistributorRegGICDIrouter801    HVGICDistributorReg = 0x7908
	HVGICDistributorRegGICDIrouter802    HVGICDistributorReg = 0x7910
	HVGICDistributorRegGICDIrouter803    HVGICDistributorReg = 0x7918
	HVGICDistributorRegGICDIrouter804    HVGICDistributorReg = 0x7920
	HVGICDistributorRegGICDIrouter805    HVGICDistributorReg = 0x7928
	HVGICDistributorRegGICDIrouter806    HVGICDistributorReg = 0x7930
	HVGICDistributorRegGICDIrouter807    HVGICDistributorReg = 0x7938
	HVGICDistributorRegGICDIrouter808    HVGICDistributorReg = 0x7940
	HVGICDistributorRegGICDIrouter809    HVGICDistributorReg = 0x7948
	HVGICDistributorRegGICDIrouter81     HVGICDistributorReg = 0x6288
	HVGICDistributorRegGICDIrouter810    HVGICDistributorReg = 0x7950
	HVGICDistributorRegGICDIrouter811    HVGICDistributorReg = 0x7958
	HVGICDistributorRegGICDIrouter812    HVGICDistributorReg = 0x7960
	HVGICDistributorRegGICDIrouter813    HVGICDistributorReg = 0x7968
	HVGICDistributorRegGICDIrouter814    HVGICDistributorReg = 0x7970
	HVGICDistributorRegGICDIrouter815    HVGICDistributorReg = 0x7978
	HVGICDistributorRegGICDIrouter816    HVGICDistributorReg = 0x7980
	HVGICDistributorRegGICDIrouter817    HVGICDistributorReg = 0x7988
	HVGICDistributorRegGICDIrouter818    HVGICDistributorReg = 0x7990
	HVGICDistributorRegGICDIrouter819    HVGICDistributorReg = 0x7998
	HVGICDistributorRegGICDIrouter82     HVGICDistributorReg = 0x6290
	HVGICDistributorRegGICDIrouter820    HVGICDistributorReg = 0x79a0
	HVGICDistributorRegGICDIrouter821    HVGICDistributorReg = 0x79a8
	HVGICDistributorRegGICDIrouter822    HVGICDistributorReg = 0x79b0
	HVGICDistributorRegGICDIrouter823    HVGICDistributorReg = 0x79b8
	HVGICDistributorRegGICDIrouter824    HVGICDistributorReg = 0x79c0
	HVGICDistributorRegGICDIrouter825    HVGICDistributorReg = 0x79c8
	HVGICDistributorRegGICDIrouter826    HVGICDistributorReg = 0x79d0
	HVGICDistributorRegGICDIrouter827    HVGICDistributorReg = 0x79d8
	HVGICDistributorRegGICDIrouter828    HVGICDistributorReg = 0x79e0
	HVGICDistributorRegGICDIrouter829    HVGICDistributorReg = 0x79e8
	HVGICDistributorRegGICDIrouter83     HVGICDistributorReg = 0x6298
	HVGICDistributorRegGICDIrouter830    HVGICDistributorReg = 0x79f0
	HVGICDistributorRegGICDIrouter831    HVGICDistributorReg = 0x79f8
	HVGICDistributorRegGICDIrouter832    HVGICDistributorReg = 0x7a00
	HVGICDistributorRegGICDIrouter833    HVGICDistributorReg = 0x7a08
	HVGICDistributorRegGICDIrouter834    HVGICDistributorReg = 0x7a10
	HVGICDistributorRegGICDIrouter835    HVGICDistributorReg = 0x7a18
	HVGICDistributorRegGICDIrouter836    HVGICDistributorReg = 0x7a20
	HVGICDistributorRegGICDIrouter837    HVGICDistributorReg = 0x7a28
	HVGICDistributorRegGICDIrouter838    HVGICDistributorReg = 0x7a30
	HVGICDistributorRegGICDIrouter839    HVGICDistributorReg = 0x7a38
	HVGICDistributorRegGICDIrouter84     HVGICDistributorReg = 0x62a0
	HVGICDistributorRegGICDIrouter840    HVGICDistributorReg = 0x7a40
	HVGICDistributorRegGICDIrouter841    HVGICDistributorReg = 0x7a48
	HVGICDistributorRegGICDIrouter842    HVGICDistributorReg = 0x7a50
	HVGICDistributorRegGICDIrouter843    HVGICDistributorReg = 0x7a58
	HVGICDistributorRegGICDIrouter844    HVGICDistributorReg = 0x7a60
	HVGICDistributorRegGICDIrouter845    HVGICDistributorReg = 0x7a68
	HVGICDistributorRegGICDIrouter846    HVGICDistributorReg = 0x7a70
	HVGICDistributorRegGICDIrouter847    HVGICDistributorReg = 0x7a78
	HVGICDistributorRegGICDIrouter848    HVGICDistributorReg = 0x7a80
	HVGICDistributorRegGICDIrouter849    HVGICDistributorReg = 0x7a88
	HVGICDistributorRegGICDIrouter85     HVGICDistributorReg = 0x62a8
	HVGICDistributorRegGICDIrouter850    HVGICDistributorReg = 0x7a90
	HVGICDistributorRegGICDIrouter851    HVGICDistributorReg = 0x7a98
	HVGICDistributorRegGICDIrouter852    HVGICDistributorReg = 0x7aa0
	HVGICDistributorRegGICDIrouter853    HVGICDistributorReg = 0x7aa8
	HVGICDistributorRegGICDIrouter854    HVGICDistributorReg = 0x7ab0
	HVGICDistributorRegGICDIrouter855    HVGICDistributorReg = 0x7ab8
	HVGICDistributorRegGICDIrouter856    HVGICDistributorReg = 0x7ac0
	HVGICDistributorRegGICDIrouter857    HVGICDistributorReg = 0x7ac8
	HVGICDistributorRegGICDIrouter858    HVGICDistributorReg = 0x7ad0
	HVGICDistributorRegGICDIrouter859    HVGICDistributorReg = 0x7ad8
	HVGICDistributorRegGICDIrouter86     HVGICDistributorReg = 0x62b0
	HVGICDistributorRegGICDIrouter860    HVGICDistributorReg = 0x7ae0
	HVGICDistributorRegGICDIrouter861    HVGICDistributorReg = 0x7ae8
	HVGICDistributorRegGICDIrouter862    HVGICDistributorReg = 0x7af0
	HVGICDistributorRegGICDIrouter863    HVGICDistributorReg = 0x7af8
	HVGICDistributorRegGICDIrouter864    HVGICDistributorReg = 0x7b00
	HVGICDistributorRegGICDIrouter865    HVGICDistributorReg = 0x7b08
	HVGICDistributorRegGICDIrouter866    HVGICDistributorReg = 0x7b10
	HVGICDistributorRegGICDIrouter867    HVGICDistributorReg = 0x7b18
	HVGICDistributorRegGICDIrouter868    HVGICDistributorReg = 0x7b20
	HVGICDistributorRegGICDIrouter869    HVGICDistributorReg = 0x7b28
	HVGICDistributorRegGICDIrouter87     HVGICDistributorReg = 0x62b8
	HVGICDistributorRegGICDIrouter870    HVGICDistributorReg = 0x7b30
	HVGICDistributorRegGICDIrouter871    HVGICDistributorReg = 0x7b38
	HVGICDistributorRegGICDIrouter872    HVGICDistributorReg = 0x7b40
	HVGICDistributorRegGICDIrouter873    HVGICDistributorReg = 0x7b48
	HVGICDistributorRegGICDIrouter874    HVGICDistributorReg = 0x7b50
	HVGICDistributorRegGICDIrouter875    HVGICDistributorReg = 0x7b58
	HVGICDistributorRegGICDIrouter876    HVGICDistributorReg = 0x7b60
	HVGICDistributorRegGICDIrouter877    HVGICDistributorReg = 0x7b68
	HVGICDistributorRegGICDIrouter878    HVGICDistributorReg = 0x7b70
	HVGICDistributorRegGICDIrouter879    HVGICDistributorReg = 0x7b78
	HVGICDistributorRegGICDIrouter88     HVGICDistributorReg = 0x62c0
	HVGICDistributorRegGICDIrouter880    HVGICDistributorReg = 0x7b80
	HVGICDistributorRegGICDIrouter881    HVGICDistributorReg = 0x7b88
	HVGICDistributorRegGICDIrouter882    HVGICDistributorReg = 0x7b90
	HVGICDistributorRegGICDIrouter883    HVGICDistributorReg = 0x7b98
	HVGICDistributorRegGICDIrouter884    HVGICDistributorReg = 0x7ba0
	HVGICDistributorRegGICDIrouter885    HVGICDistributorReg = 0x7ba8
	HVGICDistributorRegGICDIrouter886    HVGICDistributorReg = 0x7bb0
	HVGICDistributorRegGICDIrouter887    HVGICDistributorReg = 0x7bb8
	HVGICDistributorRegGICDIrouter888    HVGICDistributorReg = 0x7bc0
	HVGICDistributorRegGICDIrouter889    HVGICDistributorReg = 0x7bc8
	HVGICDistributorRegGICDIrouter89     HVGICDistributorReg = 0x62c8
	HVGICDistributorRegGICDIrouter890    HVGICDistributorReg = 0x7bd0
	HVGICDistributorRegGICDIrouter891    HVGICDistributorReg = 0x7bd8
	HVGICDistributorRegGICDIrouter892    HVGICDistributorReg = 0x7be0
	HVGICDistributorRegGICDIrouter893    HVGICDistributorReg = 0x7be8
	HVGICDistributorRegGICDIrouter894    HVGICDistributorReg = 0x7bf0
	HVGICDistributorRegGICDIrouter895    HVGICDistributorReg = 0x7bf8
	HVGICDistributorRegGICDIrouter896    HVGICDistributorReg = 0x7c00
	HVGICDistributorRegGICDIrouter897    HVGICDistributorReg = 0x7c08
	HVGICDistributorRegGICDIrouter898    HVGICDistributorReg = 0x7c10
	HVGICDistributorRegGICDIrouter899    HVGICDistributorReg = 0x7c18
	HVGICDistributorRegGICDIrouter90     HVGICDistributorReg = 0x62d0
	HVGICDistributorRegGICDIrouter900    HVGICDistributorReg = 0x7c20
	HVGICDistributorRegGICDIrouter901    HVGICDistributorReg = 0x7c28
	HVGICDistributorRegGICDIrouter902    HVGICDistributorReg = 0x7c30
	HVGICDistributorRegGICDIrouter903    HVGICDistributorReg = 0x7c38
	HVGICDistributorRegGICDIrouter904    HVGICDistributorReg = 0x7c40
	HVGICDistributorRegGICDIrouter905    HVGICDistributorReg = 0x7c48
	HVGICDistributorRegGICDIrouter906    HVGICDistributorReg = 0x7c50
	HVGICDistributorRegGICDIrouter907    HVGICDistributorReg = 0x7c58
	HVGICDistributorRegGICDIrouter908    HVGICDistributorReg = 0x7c60
	HVGICDistributorRegGICDIrouter909    HVGICDistributorReg = 0x7c68
	HVGICDistributorRegGICDIrouter91     HVGICDistributorReg = 0x62d8
	HVGICDistributorRegGICDIrouter910    HVGICDistributorReg = 0x7c70
	HVGICDistributorRegGICDIrouter911    HVGICDistributorReg = 0x7c78
	HVGICDistributorRegGICDIrouter912    HVGICDistributorReg = 0x7c80
	HVGICDistributorRegGICDIrouter913    HVGICDistributorReg = 0x7c88
	HVGICDistributorRegGICDIrouter914    HVGICDistributorReg = 0x7c90
	HVGICDistributorRegGICDIrouter915    HVGICDistributorReg = 0x7c98
	HVGICDistributorRegGICDIrouter916    HVGICDistributorReg = 0x7ca0
	HVGICDistributorRegGICDIrouter917    HVGICDistributorReg = 0x7ca8
	HVGICDistributorRegGICDIrouter918    HVGICDistributorReg = 0x7cb0
	HVGICDistributorRegGICDIrouter919    HVGICDistributorReg = 0x7cb8
	HVGICDistributorRegGICDIrouter92     HVGICDistributorReg = 0x62e0
	HVGICDistributorRegGICDIrouter920    HVGICDistributorReg = 0x7cc0
	HVGICDistributorRegGICDIrouter921    HVGICDistributorReg = 0x7cc8
	HVGICDistributorRegGICDIrouter922    HVGICDistributorReg = 0x7cd0
	HVGICDistributorRegGICDIrouter923    HVGICDistributorReg = 0x7cd8
	HVGICDistributorRegGICDIrouter924    HVGICDistributorReg = 0x7ce0
	HVGICDistributorRegGICDIrouter925    HVGICDistributorReg = 0x7ce8
	HVGICDistributorRegGICDIrouter926    HVGICDistributorReg = 0x7cf0
	HVGICDistributorRegGICDIrouter927    HVGICDistributorReg = 0x7cf8
	HVGICDistributorRegGICDIrouter928    HVGICDistributorReg = 0x7d00
	HVGICDistributorRegGICDIrouter929    HVGICDistributorReg = 0x7d08
	HVGICDistributorRegGICDIrouter93     HVGICDistributorReg = 0x62e8
	HVGICDistributorRegGICDIrouter930    HVGICDistributorReg = 0x7d10
	HVGICDistributorRegGICDIrouter931    HVGICDistributorReg = 0x7d18
	HVGICDistributorRegGICDIrouter932    HVGICDistributorReg = 0x7d20
	HVGICDistributorRegGICDIrouter933    HVGICDistributorReg = 0x7d28
	HVGICDistributorRegGICDIrouter934    HVGICDistributorReg = 0x7d30
	HVGICDistributorRegGICDIrouter935    HVGICDistributorReg = 0x7d38
	HVGICDistributorRegGICDIrouter936    HVGICDistributorReg = 0x7d40
	HVGICDistributorRegGICDIrouter937    HVGICDistributorReg = 0x7d48
	HVGICDistributorRegGICDIrouter938    HVGICDistributorReg = 0x7d50
	HVGICDistributorRegGICDIrouter939    HVGICDistributorReg = 0x7d58
	HVGICDistributorRegGICDIrouter94     HVGICDistributorReg = 0x62f0
	HVGICDistributorRegGICDIrouter940    HVGICDistributorReg = 0x7d60
	HVGICDistributorRegGICDIrouter941    HVGICDistributorReg = 0x7d68
	HVGICDistributorRegGICDIrouter942    HVGICDistributorReg = 0x7d70
	HVGICDistributorRegGICDIrouter943    HVGICDistributorReg = 0x7d78
	HVGICDistributorRegGICDIrouter944    HVGICDistributorReg = 0x7d80
	HVGICDistributorRegGICDIrouter945    HVGICDistributorReg = 0x7d88
	HVGICDistributorRegGICDIrouter946    HVGICDistributorReg = 0x7d90
	HVGICDistributorRegGICDIrouter947    HVGICDistributorReg = 0x7d98
	HVGICDistributorRegGICDIrouter948    HVGICDistributorReg = 0x7da0
	HVGICDistributorRegGICDIrouter949    HVGICDistributorReg = 0x7da8
	HVGICDistributorRegGICDIrouter95     HVGICDistributorReg = 0x62f8
	HVGICDistributorRegGICDIrouter950    HVGICDistributorReg = 0x7db0
	HVGICDistributorRegGICDIrouter951    HVGICDistributorReg = 0x7db8
	HVGICDistributorRegGICDIrouter952    HVGICDistributorReg = 0x7dc0
	HVGICDistributorRegGICDIrouter953    HVGICDistributorReg = 0x7dc8
	HVGICDistributorRegGICDIrouter954    HVGICDistributorReg = 0x7dd0
	HVGICDistributorRegGICDIrouter955    HVGICDistributorReg = 0x7dd8
	HVGICDistributorRegGICDIrouter956    HVGICDistributorReg = 0x7de0
	HVGICDistributorRegGICDIrouter957    HVGICDistributorReg = 0x7de8
	HVGICDistributorRegGICDIrouter958    HVGICDistributorReg = 0x7df0
	HVGICDistributorRegGICDIrouter959    HVGICDistributorReg = 0x7df8
	HVGICDistributorRegGICDIrouter96     HVGICDistributorReg = 0x6300
	HVGICDistributorRegGICDIrouter960    HVGICDistributorReg = 0x7e00
	HVGICDistributorRegGICDIrouter961    HVGICDistributorReg = 0x7e08
	HVGICDistributorRegGICDIrouter962    HVGICDistributorReg = 0x7e10
	HVGICDistributorRegGICDIrouter963    HVGICDistributorReg = 0x7e18
	HVGICDistributorRegGICDIrouter964    HVGICDistributorReg = 0x7e20
	HVGICDistributorRegGICDIrouter965    HVGICDistributorReg = 0x7e28
	HVGICDistributorRegGICDIrouter966    HVGICDistributorReg = 0x7e30
	HVGICDistributorRegGICDIrouter967    HVGICDistributorReg = 0x7e38
	HVGICDistributorRegGICDIrouter968    HVGICDistributorReg = 0x7e40
	HVGICDistributorRegGICDIrouter969    HVGICDistributorReg = 0x7e48
	HVGICDistributorRegGICDIrouter97     HVGICDistributorReg = 0x6308
	HVGICDistributorRegGICDIrouter970    HVGICDistributorReg = 0x7e50
	HVGICDistributorRegGICDIrouter971    HVGICDistributorReg = 0x7e58
	HVGICDistributorRegGICDIrouter972    HVGICDistributorReg = 0x7e60
	HVGICDistributorRegGICDIrouter973    HVGICDistributorReg = 0x7e68
	HVGICDistributorRegGICDIrouter974    HVGICDistributorReg = 0x7e70
	HVGICDistributorRegGICDIrouter975    HVGICDistributorReg = 0x7e78
	HVGICDistributorRegGICDIrouter976    HVGICDistributorReg = 0x7e80
	HVGICDistributorRegGICDIrouter977    HVGICDistributorReg = 0x7e88
	HVGICDistributorRegGICDIrouter978    HVGICDistributorReg = 0x7e90
	HVGICDistributorRegGICDIrouter979    HVGICDistributorReg = 0x7e98
	HVGICDistributorRegGICDIrouter98     HVGICDistributorReg = 0x6310
	HVGICDistributorRegGICDIrouter980    HVGICDistributorReg = 0x7ea0
	HVGICDistributorRegGICDIrouter981    HVGICDistributorReg = 0x7ea8
	HVGICDistributorRegGICDIrouter982    HVGICDistributorReg = 0x7eb0
	HVGICDistributorRegGICDIrouter983    HVGICDistributorReg = 0x7eb8
	HVGICDistributorRegGICDIrouter984    HVGICDistributorReg = 0x7ec0
	HVGICDistributorRegGICDIrouter985    HVGICDistributorReg = 0x7ec8
	HVGICDistributorRegGICDIrouter986    HVGICDistributorReg = 0x7ed0
	HVGICDistributorRegGICDIrouter987    HVGICDistributorReg = 0x7ed8
	HVGICDistributorRegGICDIrouter988    HVGICDistributorReg = 0x7ee0
	HVGICDistributorRegGICDIrouter989    HVGICDistributorReg = 0x7ee8
	HVGICDistributorRegGICDIrouter99     HVGICDistributorReg = 0x6318
	HVGICDistributorRegGICDIrouter990    HVGICDistributorReg = 0x7ef0
	HVGICDistributorRegGICDIrouter991    HVGICDistributorReg = 0x7ef8
	HVGICDistributorRegGICDIrouter992    HVGICDistributorReg = 0x7f00
	HVGICDistributorRegGICDIrouter993    HVGICDistributorReg = 0x7f08
	HVGICDistributorRegGICDIrouter994    HVGICDistributorReg = 0x7f10
	HVGICDistributorRegGICDIrouter995    HVGICDistributorReg = 0x7f18
	HVGICDistributorRegGICDIrouter996    HVGICDistributorReg = 0x7f20
	HVGICDistributorRegGICDIrouter997    HVGICDistributorReg = 0x7f28
	HVGICDistributorRegGICDIrouter998    HVGICDistributorReg = 0x7f30
	HVGICDistributorRegGICDIrouter999    HVGICDistributorReg = 0x7f38
	HVGICDistributorRegGICDIsactiver0    HVGICDistributorReg = 0x300
	HVGICDistributorRegGICDIsactiver1    HVGICDistributorReg = 0x304
	HVGICDistributorRegGICDIsactiver10   HVGICDistributorReg = 0x328
	HVGICDistributorRegGICDIsactiver11   HVGICDistributorReg = 0x32c
	HVGICDistributorRegGICDIsactiver12   HVGICDistributorReg = 0x330
	HVGICDistributorRegGICDIsactiver13   HVGICDistributorReg = 0x334
	HVGICDistributorRegGICDIsactiver14   HVGICDistributorReg = 0x338
	HVGICDistributorRegGICDIsactiver15   HVGICDistributorReg = 0x33c
	HVGICDistributorRegGICDIsactiver16   HVGICDistributorReg = 0x340
	HVGICDistributorRegGICDIsactiver17   HVGICDistributorReg = 0x344
	HVGICDistributorRegGICDIsactiver18   HVGICDistributorReg = 0x348
	HVGICDistributorRegGICDIsactiver19   HVGICDistributorReg = 0x34c
	HVGICDistributorRegGICDIsactiver2    HVGICDistributorReg = 0x308
	HVGICDistributorRegGICDIsactiver20   HVGICDistributorReg = 0x350
	HVGICDistributorRegGICDIsactiver21   HVGICDistributorReg = 0x354
	HVGICDistributorRegGICDIsactiver22   HVGICDistributorReg = 0x358
	HVGICDistributorRegGICDIsactiver23   HVGICDistributorReg = 0x35c
	HVGICDistributorRegGICDIsactiver24   HVGICDistributorReg = 0x360
	HVGICDistributorRegGICDIsactiver25   HVGICDistributorReg = 0x364
	HVGICDistributorRegGICDIsactiver26   HVGICDistributorReg = 0x368
	HVGICDistributorRegGICDIsactiver27   HVGICDistributorReg = 0x36c
	HVGICDistributorRegGICDIsactiver28   HVGICDistributorReg = 0x370
	HVGICDistributorRegGICDIsactiver29   HVGICDistributorReg = 0x374
	HVGICDistributorRegGICDIsactiver3    HVGICDistributorReg = 0x30c
	HVGICDistributorRegGICDIsactiver30   HVGICDistributorReg = 0x378
	HVGICDistributorRegGICDIsactiver31   HVGICDistributorReg = 0x37c
	HVGICDistributorRegGICDIsactiver4    HVGICDistributorReg = 0x310
	HVGICDistributorRegGICDIsactiver5    HVGICDistributorReg = 0x314
	HVGICDistributorRegGICDIsactiver6    HVGICDistributorReg = 0x318
	HVGICDistributorRegGICDIsactiver7    HVGICDistributorReg = 0x31c
	HVGICDistributorRegGICDIsactiver8    HVGICDistributorReg = 0x320
	HVGICDistributorRegGICDIsactiver9    HVGICDistributorReg = 0x324
	HVGICDistributorRegGICDIsenabler0    HVGICDistributorReg = 0x100
	HVGICDistributorRegGICDIsenabler1    HVGICDistributorReg = 0x104
	HVGICDistributorRegGICDIsenabler10   HVGICDistributorReg = 0x128
	HVGICDistributorRegGICDIsenabler11   HVGICDistributorReg = 0x12c
	HVGICDistributorRegGICDIsenabler12   HVGICDistributorReg = 0x130
	HVGICDistributorRegGICDIsenabler13   HVGICDistributorReg = 0x134
	HVGICDistributorRegGICDIsenabler14   HVGICDistributorReg = 0x138
	HVGICDistributorRegGICDIsenabler15   HVGICDistributorReg = 0x13c
	HVGICDistributorRegGICDIsenabler16   HVGICDistributorReg = 0x140
	HVGICDistributorRegGICDIsenabler17   HVGICDistributorReg = 0x144
	HVGICDistributorRegGICDIsenabler18   HVGICDistributorReg = 0x148
	HVGICDistributorRegGICDIsenabler19   HVGICDistributorReg = 0x14c
	HVGICDistributorRegGICDIsenabler2    HVGICDistributorReg = 0x108
	HVGICDistributorRegGICDIsenabler20   HVGICDistributorReg = 0x150
	HVGICDistributorRegGICDIsenabler21   HVGICDistributorReg = 0x154
	HVGICDistributorRegGICDIsenabler22   HVGICDistributorReg = 0x158
	HVGICDistributorRegGICDIsenabler23   HVGICDistributorReg = 0x15c
	HVGICDistributorRegGICDIsenabler24   HVGICDistributorReg = 0x160
	HVGICDistributorRegGICDIsenabler25   HVGICDistributorReg = 0x164
	HVGICDistributorRegGICDIsenabler26   HVGICDistributorReg = 0x168
	HVGICDistributorRegGICDIsenabler27   HVGICDistributorReg = 0x16c
	HVGICDistributorRegGICDIsenabler28   HVGICDistributorReg = 0x170
	HVGICDistributorRegGICDIsenabler29   HVGICDistributorReg = 0x174
	HVGICDistributorRegGICDIsenabler3    HVGICDistributorReg = 0x10c
	HVGICDistributorRegGICDIsenabler30   HVGICDistributorReg = 0x178
	HVGICDistributorRegGICDIsenabler31   HVGICDistributorReg = 0x17c
	HVGICDistributorRegGICDIsenabler4    HVGICDistributorReg = 0x110
	HVGICDistributorRegGICDIsenabler5    HVGICDistributorReg = 0x114
	HVGICDistributorRegGICDIsenabler6    HVGICDistributorReg = 0x118
	HVGICDistributorRegGICDIsenabler7    HVGICDistributorReg = 0x11c
	HVGICDistributorRegGICDIsenabler8    HVGICDistributorReg = 0x120
	HVGICDistributorRegGICDIsenabler9    HVGICDistributorReg = 0x124
	HVGICDistributorRegGICDIspendr0      HVGICDistributorReg = 0x200
	HVGICDistributorRegGICDIspendr1      HVGICDistributorReg = 0x204
	HVGICDistributorRegGICDIspendr10     HVGICDistributorReg = 0x228
	HVGICDistributorRegGICDIspendr11     HVGICDistributorReg = 0x22c
	HVGICDistributorRegGICDIspendr12     HVGICDistributorReg = 0x230
	HVGICDistributorRegGICDIspendr13     HVGICDistributorReg = 0x234
	HVGICDistributorRegGICDIspendr14     HVGICDistributorReg = 0x238
	HVGICDistributorRegGICDIspendr15     HVGICDistributorReg = 0x23c
	HVGICDistributorRegGICDIspendr16     HVGICDistributorReg = 0x240
	HVGICDistributorRegGICDIspendr17     HVGICDistributorReg = 0x244
	HVGICDistributorRegGICDIspendr18     HVGICDistributorReg = 0x248
	HVGICDistributorRegGICDIspendr19     HVGICDistributorReg = 0x24c
	HVGICDistributorRegGICDIspendr2      HVGICDistributorReg = 0x208
	HVGICDistributorRegGICDIspendr20     HVGICDistributorReg = 0x250
	HVGICDistributorRegGICDIspendr21     HVGICDistributorReg = 0x254
	HVGICDistributorRegGICDIspendr22     HVGICDistributorReg = 0x258
	HVGICDistributorRegGICDIspendr23     HVGICDistributorReg = 0x25c
	HVGICDistributorRegGICDIspendr24     HVGICDistributorReg = 0x260
	HVGICDistributorRegGICDIspendr25     HVGICDistributorReg = 0x264
	HVGICDistributorRegGICDIspendr26     HVGICDistributorReg = 0x268
	HVGICDistributorRegGICDIspendr27     HVGICDistributorReg = 0x26c
	HVGICDistributorRegGICDIspendr28     HVGICDistributorReg = 0x270
	HVGICDistributorRegGICDIspendr29     HVGICDistributorReg = 0x274
	HVGICDistributorRegGICDIspendr3      HVGICDistributorReg = 0x20c
	HVGICDistributorRegGICDIspendr30     HVGICDistributorReg = 0x278
	HVGICDistributorRegGICDIspendr31     HVGICDistributorReg = 0x27c
	HVGICDistributorRegGICDIspendr4      HVGICDistributorReg = 0x210
	HVGICDistributorRegGICDIspendr5      HVGICDistributorReg = 0x214
	HVGICDistributorRegGICDIspendr6      HVGICDistributorReg = 0x218
	HVGICDistributorRegGICDIspendr7      HVGICDistributorReg = 0x21c
	HVGICDistributorRegGICDIspendr8      HVGICDistributorReg = 0x220
	HVGICDistributorRegGICDIspendr9      HVGICDistributorReg = 0x224
	HVGICDistributorRegGICDPidr2         HVGICDistributorReg = 0xffe8
	HVGICDistributorRegGICDTyper         HVGICDistributorReg = 0x4
)

func (e HVGICDistributorReg) String() string {
	switch e {
	case HVGICDistributorRegGICDCtlr:
		return "HVGICDistributorRegGICDCtlr"
	case HVGICDistributorRegGICDIcactiver0:
		return "HVGICDistributorRegGICDIcactiver0"
	case HVGICDistributorRegGICDIcactiver1:
		return "HVGICDistributorRegGICDIcactiver1"
	case HVGICDistributorRegGICDIcactiver10:
		return "HVGICDistributorRegGICDIcactiver10"
	case HVGICDistributorRegGICDIcactiver11:
		return "HVGICDistributorRegGICDIcactiver11"
	case HVGICDistributorRegGICDIcactiver12:
		return "HVGICDistributorRegGICDIcactiver12"
	case HVGICDistributorRegGICDIcactiver13:
		return "HVGICDistributorRegGICDIcactiver13"
	case HVGICDistributorRegGICDIcactiver14:
		return "HVGICDistributorRegGICDIcactiver14"
	case HVGICDistributorRegGICDIcactiver15:
		return "HVGICDistributorRegGICDIcactiver15"
	case HVGICDistributorRegGICDIcactiver16:
		return "HVGICDistributorRegGICDIcactiver16"
	case HVGICDistributorRegGICDIcactiver17:
		return "HVGICDistributorRegGICDIcactiver17"
	case HVGICDistributorRegGICDIcactiver18:
		return "HVGICDistributorRegGICDIcactiver18"
	case HVGICDistributorRegGICDIcactiver19:
		return "HVGICDistributorRegGICDIcactiver19"
	case HVGICDistributorRegGICDIcactiver2:
		return "HVGICDistributorRegGICDIcactiver2"
	case HVGICDistributorRegGICDIcactiver20:
		return "HVGICDistributorRegGICDIcactiver20"
	case HVGICDistributorRegGICDIcactiver21:
		return "HVGICDistributorRegGICDIcactiver21"
	case HVGICDistributorRegGICDIcactiver22:
		return "HVGICDistributorRegGICDIcactiver22"
	case HVGICDistributorRegGICDIcactiver23:
		return "HVGICDistributorRegGICDIcactiver23"
	case HVGICDistributorRegGICDIcactiver24:
		return "HVGICDistributorRegGICDIcactiver24"
	case HVGICDistributorRegGICDIcactiver25:
		return "HVGICDistributorRegGICDIcactiver25"
	case HVGICDistributorRegGICDIcactiver26:
		return "HVGICDistributorRegGICDIcactiver26"
	case HVGICDistributorRegGICDIcactiver27:
		return "HVGICDistributorRegGICDIcactiver27"
	case HVGICDistributorRegGICDIcactiver28:
		return "HVGICDistributorRegGICDIcactiver28"
	case HVGICDistributorRegGICDIcactiver29:
		return "HVGICDistributorRegGICDIcactiver29"
	case HVGICDistributorRegGICDIcactiver3:
		return "HVGICDistributorRegGICDIcactiver3"
	case HVGICDistributorRegGICDIcactiver30:
		return "HVGICDistributorRegGICDIcactiver30"
	case HVGICDistributorRegGICDIcactiver31:
		return "HVGICDistributorRegGICDIcactiver31"
	case HVGICDistributorRegGICDIcactiver4:
		return "HVGICDistributorRegGICDIcactiver4"
	case HVGICDistributorRegGICDIcactiver5:
		return "HVGICDistributorRegGICDIcactiver5"
	case HVGICDistributorRegGICDIcactiver6:
		return "HVGICDistributorRegGICDIcactiver6"
	case HVGICDistributorRegGICDIcactiver7:
		return "HVGICDistributorRegGICDIcactiver7"
	case HVGICDistributorRegGICDIcactiver8:
		return "HVGICDistributorRegGICDIcactiver8"
	case HVGICDistributorRegGICDIcactiver9:
		return "HVGICDistributorRegGICDIcactiver9"
	case HVGICDistributorRegGICDIcenabler0:
		return "HVGICDistributorRegGICDIcenabler0"
	case HVGICDistributorRegGICDIcenabler1:
		return "HVGICDistributorRegGICDIcenabler1"
	case HVGICDistributorRegGICDIcenabler10:
		return "HVGICDistributorRegGICDIcenabler10"
	case HVGICDistributorRegGICDIcenabler11:
		return "HVGICDistributorRegGICDIcenabler11"
	case HVGICDistributorRegGICDIcenabler12:
		return "HVGICDistributorRegGICDIcenabler12"
	case HVGICDistributorRegGICDIcenabler13:
		return "HVGICDistributorRegGICDIcenabler13"
	case HVGICDistributorRegGICDIcenabler14:
		return "HVGICDistributorRegGICDIcenabler14"
	case HVGICDistributorRegGICDIcenabler15:
		return "HVGICDistributorRegGICDIcenabler15"
	case HVGICDistributorRegGICDIcenabler16:
		return "HVGICDistributorRegGICDIcenabler16"
	case HVGICDistributorRegGICDIcenabler17:
		return "HVGICDistributorRegGICDIcenabler17"
	case HVGICDistributorRegGICDIcenabler18:
		return "HVGICDistributorRegGICDIcenabler18"
	case HVGICDistributorRegGICDIcenabler19:
		return "HVGICDistributorRegGICDIcenabler19"
	case HVGICDistributorRegGICDIcenabler2:
		return "HVGICDistributorRegGICDIcenabler2"
	case HVGICDistributorRegGICDIcenabler20:
		return "HVGICDistributorRegGICDIcenabler20"
	case HVGICDistributorRegGICDIcenabler21:
		return "HVGICDistributorRegGICDIcenabler21"
	case HVGICDistributorRegGICDIcenabler22:
		return "HVGICDistributorRegGICDIcenabler22"
	case HVGICDistributorRegGICDIcenabler23:
		return "HVGICDistributorRegGICDIcenabler23"
	case HVGICDistributorRegGICDIcenabler24:
		return "HVGICDistributorRegGICDIcenabler24"
	case HVGICDistributorRegGICDIcenabler25:
		return "HVGICDistributorRegGICDIcenabler25"
	case HVGICDistributorRegGICDIcenabler26:
		return "HVGICDistributorRegGICDIcenabler26"
	case HVGICDistributorRegGICDIcenabler27:
		return "HVGICDistributorRegGICDIcenabler27"
	case HVGICDistributorRegGICDIcenabler28:
		return "HVGICDistributorRegGICDIcenabler28"
	case HVGICDistributorRegGICDIcenabler29:
		return "HVGICDistributorRegGICDIcenabler29"
	case HVGICDistributorRegGICDIcenabler3:
		return "HVGICDistributorRegGICDIcenabler3"
	case HVGICDistributorRegGICDIcenabler30:
		return "HVGICDistributorRegGICDIcenabler30"
	case HVGICDistributorRegGICDIcenabler31:
		return "HVGICDistributorRegGICDIcenabler31"
	case HVGICDistributorRegGICDIcenabler4:
		return "HVGICDistributorRegGICDIcenabler4"
	case HVGICDistributorRegGICDIcenabler5:
		return "HVGICDistributorRegGICDIcenabler5"
	case HVGICDistributorRegGICDIcenabler6:
		return "HVGICDistributorRegGICDIcenabler6"
	case HVGICDistributorRegGICDIcenabler7:
		return "HVGICDistributorRegGICDIcenabler7"
	case HVGICDistributorRegGICDIcenabler8:
		return "HVGICDistributorRegGICDIcenabler8"
	case HVGICDistributorRegGICDIcenabler9:
		return "HVGICDistributorRegGICDIcenabler9"
	case HVGICDistributorRegGICDIcfgr0:
		return "HVGICDistributorRegGICDIcfgr0"
	case HVGICDistributorRegGICDIcfgr1:
		return "HVGICDistributorRegGICDIcfgr1"
	case HVGICDistributorRegGICDIcfgr10:
		return "HVGICDistributorRegGICDIcfgr10"
	case HVGICDistributorRegGICDIcfgr11:
		return "HVGICDistributorRegGICDIcfgr11"
	case HVGICDistributorRegGICDIcfgr12:
		return "HVGICDistributorRegGICDIcfgr12"
	case HVGICDistributorRegGICDIcfgr13:
		return "HVGICDistributorRegGICDIcfgr13"
	case HVGICDistributorRegGICDIcfgr14:
		return "HVGICDistributorRegGICDIcfgr14"
	case HVGICDistributorRegGICDIcfgr15:
		return "HVGICDistributorRegGICDIcfgr15"
	case HVGICDistributorRegGICDIcfgr16:
		return "HVGICDistributorRegGICDIcfgr16"
	case HVGICDistributorRegGICDIcfgr17:
		return "HVGICDistributorRegGICDIcfgr17"
	case HVGICDistributorRegGICDIcfgr18:
		return "HVGICDistributorRegGICDIcfgr18"
	case HVGICDistributorRegGICDIcfgr19:
		return "HVGICDistributorRegGICDIcfgr19"
	case HVGICDistributorRegGICDIcfgr2:
		return "HVGICDistributorRegGICDIcfgr2"
	case HVGICDistributorRegGICDIcfgr20:
		return "HVGICDistributorRegGICDIcfgr20"
	case HVGICDistributorRegGICDIcfgr21:
		return "HVGICDistributorRegGICDIcfgr21"
	case HVGICDistributorRegGICDIcfgr22:
		return "HVGICDistributorRegGICDIcfgr22"
	case HVGICDistributorRegGICDIcfgr23:
		return "HVGICDistributorRegGICDIcfgr23"
	case HVGICDistributorRegGICDIcfgr24:
		return "HVGICDistributorRegGICDIcfgr24"
	case HVGICDistributorRegGICDIcfgr25:
		return "HVGICDistributorRegGICDIcfgr25"
	case HVGICDistributorRegGICDIcfgr26:
		return "HVGICDistributorRegGICDIcfgr26"
	case HVGICDistributorRegGICDIcfgr27:
		return "HVGICDistributorRegGICDIcfgr27"
	case HVGICDistributorRegGICDIcfgr28:
		return "HVGICDistributorRegGICDIcfgr28"
	case HVGICDistributorRegGICDIcfgr29:
		return "HVGICDistributorRegGICDIcfgr29"
	case HVGICDistributorRegGICDIcfgr3:
		return "HVGICDistributorRegGICDIcfgr3"
	case HVGICDistributorRegGICDIcfgr30:
		return "HVGICDistributorRegGICDIcfgr30"
	case HVGICDistributorRegGICDIcfgr31:
		return "HVGICDistributorRegGICDIcfgr31"
	case HVGICDistributorRegGICDIcfgr32:
		return "HVGICDistributorRegGICDIcfgr32"
	case HVGICDistributorRegGICDIcfgr33:
		return "HVGICDistributorRegGICDIcfgr33"
	case HVGICDistributorRegGICDIcfgr34:
		return "HVGICDistributorRegGICDIcfgr34"
	case HVGICDistributorRegGICDIcfgr35:
		return "HVGICDistributorRegGICDIcfgr35"
	case HVGICDistributorRegGICDIcfgr36:
		return "HVGICDistributorRegGICDIcfgr36"
	case HVGICDistributorRegGICDIcfgr37:
		return "HVGICDistributorRegGICDIcfgr37"
	case HVGICDistributorRegGICDIcfgr38:
		return "HVGICDistributorRegGICDIcfgr38"
	case HVGICDistributorRegGICDIcfgr39:
		return "HVGICDistributorRegGICDIcfgr39"
	case HVGICDistributorRegGICDIcfgr4:
		return "HVGICDistributorRegGICDIcfgr4"
	case HVGICDistributorRegGICDIcfgr40:
		return "HVGICDistributorRegGICDIcfgr40"
	case HVGICDistributorRegGICDIcfgr41:
		return "HVGICDistributorRegGICDIcfgr41"
	case HVGICDistributorRegGICDIcfgr42:
		return "HVGICDistributorRegGICDIcfgr42"
	case HVGICDistributorRegGICDIcfgr43:
		return "HVGICDistributorRegGICDIcfgr43"
	case HVGICDistributorRegGICDIcfgr44:
		return "HVGICDistributorRegGICDIcfgr44"
	case HVGICDistributorRegGICDIcfgr45:
		return "HVGICDistributorRegGICDIcfgr45"
	case HVGICDistributorRegGICDIcfgr46:
		return "HVGICDistributorRegGICDIcfgr46"
	case HVGICDistributorRegGICDIcfgr47:
		return "HVGICDistributorRegGICDIcfgr47"
	case HVGICDistributorRegGICDIcfgr48:
		return "HVGICDistributorRegGICDIcfgr48"
	case HVGICDistributorRegGICDIcfgr49:
		return "HVGICDistributorRegGICDIcfgr49"
	case HVGICDistributorRegGICDIcfgr5:
		return "HVGICDistributorRegGICDIcfgr5"
	case HVGICDistributorRegGICDIcfgr50:
		return "HVGICDistributorRegGICDIcfgr50"
	case HVGICDistributorRegGICDIcfgr51:
		return "HVGICDistributorRegGICDIcfgr51"
	case HVGICDistributorRegGICDIcfgr52:
		return "HVGICDistributorRegGICDIcfgr52"
	case HVGICDistributorRegGICDIcfgr53:
		return "HVGICDistributorRegGICDIcfgr53"
	case HVGICDistributorRegGICDIcfgr54:
		return "HVGICDistributorRegGICDIcfgr54"
	case HVGICDistributorRegGICDIcfgr55:
		return "HVGICDistributorRegGICDIcfgr55"
	case HVGICDistributorRegGICDIcfgr56:
		return "HVGICDistributorRegGICDIcfgr56"
	case HVGICDistributorRegGICDIcfgr57:
		return "HVGICDistributorRegGICDIcfgr57"
	case HVGICDistributorRegGICDIcfgr58:
		return "HVGICDistributorRegGICDIcfgr58"
	case HVGICDistributorRegGICDIcfgr59:
		return "HVGICDistributorRegGICDIcfgr59"
	case HVGICDistributorRegGICDIcfgr6:
		return "HVGICDistributorRegGICDIcfgr6"
	case HVGICDistributorRegGICDIcfgr60:
		return "HVGICDistributorRegGICDIcfgr60"
	case HVGICDistributorRegGICDIcfgr61:
		return "HVGICDistributorRegGICDIcfgr61"
	case HVGICDistributorRegGICDIcfgr62:
		return "HVGICDistributorRegGICDIcfgr62"
	case HVGICDistributorRegGICDIcfgr63:
		return "HVGICDistributorRegGICDIcfgr63"
	case HVGICDistributorRegGICDIcfgr7:
		return "HVGICDistributorRegGICDIcfgr7"
	case HVGICDistributorRegGICDIcfgr8:
		return "HVGICDistributorRegGICDIcfgr8"
	case HVGICDistributorRegGICDIcfgr9:
		return "HVGICDistributorRegGICDIcfgr9"
	case HVGICDistributorRegGICDIcpendr0:
		return "HVGICDistributorRegGICDIcpendr0"
	case HVGICDistributorRegGICDIcpendr1:
		return "HVGICDistributorRegGICDIcpendr1"
	case HVGICDistributorRegGICDIcpendr10:
		return "HVGICDistributorRegGICDIcpendr10"
	case HVGICDistributorRegGICDIcpendr11:
		return "HVGICDistributorRegGICDIcpendr11"
	case HVGICDistributorRegGICDIcpendr12:
		return "HVGICDistributorRegGICDIcpendr12"
	case HVGICDistributorRegGICDIcpendr13:
		return "HVGICDistributorRegGICDIcpendr13"
	case HVGICDistributorRegGICDIcpendr14:
		return "HVGICDistributorRegGICDIcpendr14"
	case HVGICDistributorRegGICDIcpendr15:
		return "HVGICDistributorRegGICDIcpendr15"
	case HVGICDistributorRegGICDIcpendr16:
		return "HVGICDistributorRegGICDIcpendr16"
	case HVGICDistributorRegGICDIcpendr17:
		return "HVGICDistributorRegGICDIcpendr17"
	case HVGICDistributorRegGICDIcpendr18:
		return "HVGICDistributorRegGICDIcpendr18"
	case HVGICDistributorRegGICDIcpendr19:
		return "HVGICDistributorRegGICDIcpendr19"
	case HVGICDistributorRegGICDIcpendr2:
		return "HVGICDistributorRegGICDIcpendr2"
	case HVGICDistributorRegGICDIcpendr20:
		return "HVGICDistributorRegGICDIcpendr20"
	case HVGICDistributorRegGICDIcpendr21:
		return "HVGICDistributorRegGICDIcpendr21"
	case HVGICDistributorRegGICDIcpendr22:
		return "HVGICDistributorRegGICDIcpendr22"
	case HVGICDistributorRegGICDIcpendr23:
		return "HVGICDistributorRegGICDIcpendr23"
	case HVGICDistributorRegGICDIcpendr24:
		return "HVGICDistributorRegGICDIcpendr24"
	case HVGICDistributorRegGICDIcpendr25:
		return "HVGICDistributorRegGICDIcpendr25"
	case HVGICDistributorRegGICDIcpendr26:
		return "HVGICDistributorRegGICDIcpendr26"
	case HVGICDistributorRegGICDIcpendr27:
		return "HVGICDistributorRegGICDIcpendr27"
	case HVGICDistributorRegGICDIcpendr28:
		return "HVGICDistributorRegGICDIcpendr28"
	case HVGICDistributorRegGICDIcpendr29:
		return "HVGICDistributorRegGICDIcpendr29"
	case HVGICDistributorRegGICDIcpendr3:
		return "HVGICDistributorRegGICDIcpendr3"
	case HVGICDistributorRegGICDIcpendr30:
		return "HVGICDistributorRegGICDIcpendr30"
	case HVGICDistributorRegGICDIcpendr31:
		return "HVGICDistributorRegGICDIcpendr31"
	case HVGICDistributorRegGICDIcpendr4:
		return "HVGICDistributorRegGICDIcpendr4"
	case HVGICDistributorRegGICDIcpendr5:
		return "HVGICDistributorRegGICDIcpendr5"
	case HVGICDistributorRegGICDIcpendr6:
		return "HVGICDistributorRegGICDIcpendr6"
	case HVGICDistributorRegGICDIcpendr7:
		return "HVGICDistributorRegGICDIcpendr7"
	case HVGICDistributorRegGICDIcpendr8:
		return "HVGICDistributorRegGICDIcpendr8"
	case HVGICDistributorRegGICDIcpendr9:
		return "HVGICDistributorRegGICDIcpendr9"
	case HVGICDistributorRegGICDIgroupr0:
		return "HVGICDistributorRegGICDIgroupr0"
	case HVGICDistributorRegGICDIgroupr1:
		return "HVGICDistributorRegGICDIgroupr1"
	case HVGICDistributorRegGICDIgroupr10:
		return "HVGICDistributorRegGICDIgroupr10"
	case HVGICDistributorRegGICDIgroupr11:
		return "HVGICDistributorRegGICDIgroupr11"
	case HVGICDistributorRegGICDIgroupr12:
		return "HVGICDistributorRegGICDIgroupr12"
	case HVGICDistributorRegGICDIgroupr13:
		return "HVGICDistributorRegGICDIgroupr13"
	case HVGICDistributorRegGICDIgroupr14:
		return "HVGICDistributorRegGICDIgroupr14"
	case HVGICDistributorRegGICDIgroupr15:
		return "HVGICDistributorRegGICDIgroupr15"
	case HVGICDistributorRegGICDIgroupr16:
		return "HVGICDistributorRegGICDIgroupr16"
	case HVGICDistributorRegGICDIgroupr17:
		return "HVGICDistributorRegGICDIgroupr17"
	case HVGICDistributorRegGICDIgroupr18:
		return "HVGICDistributorRegGICDIgroupr18"
	case HVGICDistributorRegGICDIgroupr19:
		return "HVGICDistributorRegGICDIgroupr19"
	case HVGICDistributorRegGICDIgroupr2:
		return "HVGICDistributorRegGICDIgroupr2"
	case HVGICDistributorRegGICDIgroupr20:
		return "HVGICDistributorRegGICDIgroupr20"
	case HVGICDistributorRegGICDIgroupr21:
		return "HVGICDistributorRegGICDIgroupr21"
	case HVGICDistributorRegGICDIgroupr22:
		return "HVGICDistributorRegGICDIgroupr22"
	case HVGICDistributorRegGICDIgroupr23:
		return "HVGICDistributorRegGICDIgroupr23"
	case HVGICDistributorRegGICDIgroupr24:
		return "HVGICDistributorRegGICDIgroupr24"
	case HVGICDistributorRegGICDIgroupr25:
		return "HVGICDistributorRegGICDIgroupr25"
	case HVGICDistributorRegGICDIgroupr26:
		return "HVGICDistributorRegGICDIgroupr26"
	case HVGICDistributorRegGICDIgroupr27:
		return "HVGICDistributorRegGICDIgroupr27"
	case HVGICDistributorRegGICDIgroupr28:
		return "HVGICDistributorRegGICDIgroupr28"
	case HVGICDistributorRegGICDIgroupr29:
		return "HVGICDistributorRegGICDIgroupr29"
	case HVGICDistributorRegGICDIgroupr3:
		return "HVGICDistributorRegGICDIgroupr3"
	case HVGICDistributorRegGICDIgroupr30:
		return "HVGICDistributorRegGICDIgroupr30"
	case HVGICDistributorRegGICDIgroupr31:
		return "HVGICDistributorRegGICDIgroupr31"
	case HVGICDistributorRegGICDIgroupr4:
		return "HVGICDistributorRegGICDIgroupr4"
	case HVGICDistributorRegGICDIgroupr5:
		return "HVGICDistributorRegGICDIgroupr5"
	case HVGICDistributorRegGICDIgroupr6:
		return "HVGICDistributorRegGICDIgroupr6"
	case HVGICDistributorRegGICDIgroupr7:
		return "HVGICDistributorRegGICDIgroupr7"
	case HVGICDistributorRegGICDIgroupr8:
		return "HVGICDistributorRegGICDIgroupr8"
	case HVGICDistributorRegGICDIgroupr9:
		return "HVGICDistributorRegGICDIgroupr9"
	case HVGICDistributorRegGICDIpriorityr0:
		return "HVGICDistributorRegGICDIpriorityr0"
	case HVGICDistributorRegGICDIpriorityr1:
		return "HVGICDistributorRegGICDIpriorityr1"
	case HVGICDistributorRegGICDIpriorityr10:
		return "HVGICDistributorRegGICDIpriorityr10"
	case HVGICDistributorRegGICDIpriorityr100:
		return "HVGICDistributorRegGICDIpriorityr100"
	case HVGICDistributorRegGICDIpriorityr101:
		return "HVGICDistributorRegGICDIpriorityr101"
	case HVGICDistributorRegGICDIpriorityr102:
		return "HVGICDistributorRegGICDIpriorityr102"
	case HVGICDistributorRegGICDIpriorityr103:
		return "HVGICDistributorRegGICDIpriorityr103"
	case HVGICDistributorRegGICDIpriorityr104:
		return "HVGICDistributorRegGICDIpriorityr104"
	case HVGICDistributorRegGICDIpriorityr105:
		return "HVGICDistributorRegGICDIpriorityr105"
	case HVGICDistributorRegGICDIpriorityr106:
		return "HVGICDistributorRegGICDIpriorityr106"
	case HVGICDistributorRegGICDIpriorityr107:
		return "HVGICDistributorRegGICDIpriorityr107"
	case HVGICDistributorRegGICDIpriorityr108:
		return "HVGICDistributorRegGICDIpriorityr108"
	case HVGICDistributorRegGICDIpriorityr109:
		return "HVGICDistributorRegGICDIpriorityr109"
	case HVGICDistributorRegGICDIpriorityr11:
		return "HVGICDistributorRegGICDIpriorityr11"
	case HVGICDistributorRegGICDIpriorityr110:
		return "HVGICDistributorRegGICDIpriorityr110"
	case HVGICDistributorRegGICDIpriorityr111:
		return "HVGICDistributorRegGICDIpriorityr111"
	case HVGICDistributorRegGICDIpriorityr112:
		return "HVGICDistributorRegGICDIpriorityr112"
	case HVGICDistributorRegGICDIpriorityr113:
		return "HVGICDistributorRegGICDIpriorityr113"
	case HVGICDistributorRegGICDIpriorityr114:
		return "HVGICDistributorRegGICDIpriorityr114"
	case HVGICDistributorRegGICDIpriorityr115:
		return "HVGICDistributorRegGICDIpriorityr115"
	case HVGICDistributorRegGICDIpriorityr116:
		return "HVGICDistributorRegGICDIpriorityr116"
	case HVGICDistributorRegGICDIpriorityr117:
		return "HVGICDistributorRegGICDIpriorityr117"
	case HVGICDistributorRegGICDIpriorityr118:
		return "HVGICDistributorRegGICDIpriorityr118"
	case HVGICDistributorRegGICDIpriorityr119:
		return "HVGICDistributorRegGICDIpriorityr119"
	case HVGICDistributorRegGICDIpriorityr12:
		return "HVGICDistributorRegGICDIpriorityr12"
	case HVGICDistributorRegGICDIpriorityr120:
		return "HVGICDistributorRegGICDIpriorityr120"
	case HVGICDistributorRegGICDIpriorityr121:
		return "HVGICDistributorRegGICDIpriorityr121"
	case HVGICDistributorRegGICDIpriorityr122:
		return "HVGICDistributorRegGICDIpriorityr122"
	case HVGICDistributorRegGICDIpriorityr123:
		return "HVGICDistributorRegGICDIpriorityr123"
	case HVGICDistributorRegGICDIpriorityr124:
		return "HVGICDistributorRegGICDIpriorityr124"
	case HVGICDistributorRegGICDIpriorityr125:
		return "HVGICDistributorRegGICDIpriorityr125"
	case HVGICDistributorRegGICDIpriorityr126:
		return "HVGICDistributorRegGICDIpriorityr126"
	case HVGICDistributorRegGICDIpriorityr127:
		return "HVGICDistributorRegGICDIpriorityr127"
	case HVGICDistributorRegGICDIpriorityr128:
		return "HVGICDistributorRegGICDIpriorityr128"
	case HVGICDistributorRegGICDIpriorityr129:
		return "HVGICDistributorRegGICDIpriorityr129"
	case HVGICDistributorRegGICDIpriorityr13:
		return "HVGICDistributorRegGICDIpriorityr13"
	case HVGICDistributorRegGICDIpriorityr130:
		return "HVGICDistributorRegGICDIpriorityr130"
	case HVGICDistributorRegGICDIpriorityr131:
		return "HVGICDistributorRegGICDIpriorityr131"
	case HVGICDistributorRegGICDIpriorityr132:
		return "HVGICDistributorRegGICDIpriorityr132"
	case HVGICDistributorRegGICDIpriorityr133:
		return "HVGICDistributorRegGICDIpriorityr133"
	case HVGICDistributorRegGICDIpriorityr134:
		return "HVGICDistributorRegGICDIpriorityr134"
	case HVGICDistributorRegGICDIpriorityr135:
		return "HVGICDistributorRegGICDIpriorityr135"
	case HVGICDistributorRegGICDIpriorityr136:
		return "HVGICDistributorRegGICDIpriorityr136"
	case HVGICDistributorRegGICDIpriorityr137:
		return "HVGICDistributorRegGICDIpriorityr137"
	case HVGICDistributorRegGICDIpriorityr138:
		return "HVGICDistributorRegGICDIpriorityr138"
	case HVGICDistributorRegGICDIpriorityr139:
		return "HVGICDistributorRegGICDIpriorityr139"
	case HVGICDistributorRegGICDIpriorityr14:
		return "HVGICDistributorRegGICDIpriorityr14"
	case HVGICDistributorRegGICDIpriorityr140:
		return "HVGICDistributorRegGICDIpriorityr140"
	case HVGICDistributorRegGICDIpriorityr141:
		return "HVGICDistributorRegGICDIpriorityr141"
	case HVGICDistributorRegGICDIpriorityr142:
		return "HVGICDistributorRegGICDIpriorityr142"
	case HVGICDistributorRegGICDIpriorityr143:
		return "HVGICDistributorRegGICDIpriorityr143"
	case HVGICDistributorRegGICDIpriorityr144:
		return "HVGICDistributorRegGICDIpriorityr144"
	case HVGICDistributorRegGICDIpriorityr145:
		return "HVGICDistributorRegGICDIpriorityr145"
	case HVGICDistributorRegGICDIpriorityr146:
		return "HVGICDistributorRegGICDIpriorityr146"
	case HVGICDistributorRegGICDIpriorityr147:
		return "HVGICDistributorRegGICDIpriorityr147"
	case HVGICDistributorRegGICDIpriorityr148:
		return "HVGICDistributorRegGICDIpriorityr148"
	case HVGICDistributorRegGICDIpriorityr149:
		return "HVGICDistributorRegGICDIpriorityr149"
	case HVGICDistributorRegGICDIpriorityr15:
		return "HVGICDistributorRegGICDIpriorityr15"
	case HVGICDistributorRegGICDIpriorityr150:
		return "HVGICDistributorRegGICDIpriorityr150"
	case HVGICDistributorRegGICDIpriorityr151:
		return "HVGICDistributorRegGICDIpriorityr151"
	case HVGICDistributorRegGICDIpriorityr152:
		return "HVGICDistributorRegGICDIpriorityr152"
	case HVGICDistributorRegGICDIpriorityr153:
		return "HVGICDistributorRegGICDIpriorityr153"
	case HVGICDistributorRegGICDIpriorityr154:
		return "HVGICDistributorRegGICDIpriorityr154"
	case HVGICDistributorRegGICDIpriorityr155:
		return "HVGICDistributorRegGICDIpriorityr155"
	case HVGICDistributorRegGICDIpriorityr156:
		return "HVGICDistributorRegGICDIpriorityr156"
	case HVGICDistributorRegGICDIpriorityr157:
		return "HVGICDistributorRegGICDIpriorityr157"
	case HVGICDistributorRegGICDIpriorityr158:
		return "HVGICDistributorRegGICDIpriorityr158"
	case HVGICDistributorRegGICDIpriorityr159:
		return "HVGICDistributorRegGICDIpriorityr159"
	case HVGICDistributorRegGICDIpriorityr16:
		return "HVGICDistributorRegGICDIpriorityr16"
	case HVGICDistributorRegGICDIpriorityr160:
		return "HVGICDistributorRegGICDIpriorityr160"
	case HVGICDistributorRegGICDIpriorityr161:
		return "HVGICDistributorRegGICDIpriorityr161"
	case HVGICDistributorRegGICDIpriorityr162:
		return "HVGICDistributorRegGICDIpriorityr162"
	case HVGICDistributorRegGICDIpriorityr163:
		return "HVGICDistributorRegGICDIpriorityr163"
	case HVGICDistributorRegGICDIpriorityr164:
		return "HVGICDistributorRegGICDIpriorityr164"
	case HVGICDistributorRegGICDIpriorityr165:
		return "HVGICDistributorRegGICDIpriorityr165"
	case HVGICDistributorRegGICDIpriorityr166:
		return "HVGICDistributorRegGICDIpriorityr166"
	case HVGICDistributorRegGICDIpriorityr167:
		return "HVGICDistributorRegGICDIpriorityr167"
	case HVGICDistributorRegGICDIpriorityr168:
		return "HVGICDistributorRegGICDIpriorityr168"
	case HVGICDistributorRegGICDIpriorityr169:
		return "HVGICDistributorRegGICDIpriorityr169"
	case HVGICDistributorRegGICDIpriorityr17:
		return "HVGICDistributorRegGICDIpriorityr17"
	case HVGICDistributorRegGICDIpriorityr170:
		return "HVGICDistributorRegGICDIpriorityr170"
	case HVGICDistributorRegGICDIpriorityr171:
		return "HVGICDistributorRegGICDIpriorityr171"
	case HVGICDistributorRegGICDIpriorityr172:
		return "HVGICDistributorRegGICDIpriorityr172"
	case HVGICDistributorRegGICDIpriorityr173:
		return "HVGICDistributorRegGICDIpriorityr173"
	case HVGICDistributorRegGICDIpriorityr174:
		return "HVGICDistributorRegGICDIpriorityr174"
	case HVGICDistributorRegGICDIpriorityr175:
		return "HVGICDistributorRegGICDIpriorityr175"
	case HVGICDistributorRegGICDIpriorityr176:
		return "HVGICDistributorRegGICDIpriorityr176"
	case HVGICDistributorRegGICDIpriorityr177:
		return "HVGICDistributorRegGICDIpriorityr177"
	case HVGICDistributorRegGICDIpriorityr178:
		return "HVGICDistributorRegGICDIpriorityr178"
	case HVGICDistributorRegGICDIpriorityr179:
		return "HVGICDistributorRegGICDIpriorityr179"
	case HVGICDistributorRegGICDIpriorityr18:
		return "HVGICDistributorRegGICDIpriorityr18"
	case HVGICDistributorRegGICDIpriorityr180:
		return "HVGICDistributorRegGICDIpriorityr180"
	case HVGICDistributorRegGICDIpriorityr181:
		return "HVGICDistributorRegGICDIpriorityr181"
	case HVGICDistributorRegGICDIpriorityr182:
		return "HVGICDistributorRegGICDIpriorityr182"
	case HVGICDistributorRegGICDIpriorityr183:
		return "HVGICDistributorRegGICDIpriorityr183"
	case HVGICDistributorRegGICDIpriorityr184:
		return "HVGICDistributorRegGICDIpriorityr184"
	case HVGICDistributorRegGICDIpriorityr185:
		return "HVGICDistributorRegGICDIpriorityr185"
	case HVGICDistributorRegGICDIpriorityr186:
		return "HVGICDistributorRegGICDIpriorityr186"
	case HVGICDistributorRegGICDIpriorityr187:
		return "HVGICDistributorRegGICDIpriorityr187"
	case HVGICDistributorRegGICDIpriorityr188:
		return "HVGICDistributorRegGICDIpriorityr188"
	case HVGICDistributorRegGICDIpriorityr189:
		return "HVGICDistributorRegGICDIpriorityr189"
	case HVGICDistributorRegGICDIpriorityr19:
		return "HVGICDistributorRegGICDIpriorityr19"
	case HVGICDistributorRegGICDIpriorityr190:
		return "HVGICDistributorRegGICDIpriorityr190"
	case HVGICDistributorRegGICDIpriorityr191:
		return "HVGICDistributorRegGICDIpriorityr191"
	case HVGICDistributorRegGICDIpriorityr192:
		return "HVGICDistributorRegGICDIpriorityr192"
	case HVGICDistributorRegGICDIpriorityr193:
		return "HVGICDistributorRegGICDIpriorityr193"
	case HVGICDistributorRegGICDIpriorityr194:
		return "HVGICDistributorRegGICDIpriorityr194"
	case HVGICDistributorRegGICDIpriorityr195:
		return "HVGICDistributorRegGICDIpriorityr195"
	case HVGICDistributorRegGICDIpriorityr196:
		return "HVGICDistributorRegGICDIpriorityr196"
	case HVGICDistributorRegGICDIpriorityr197:
		return "HVGICDistributorRegGICDIpriorityr197"
	case HVGICDistributorRegGICDIpriorityr198:
		return "HVGICDistributorRegGICDIpriorityr198"
	case HVGICDistributorRegGICDIpriorityr199:
		return "HVGICDistributorRegGICDIpriorityr199"
	case HVGICDistributorRegGICDIpriorityr2:
		return "HVGICDistributorRegGICDIpriorityr2"
	case HVGICDistributorRegGICDIpriorityr20:
		return "HVGICDistributorRegGICDIpriorityr20"
	case HVGICDistributorRegGICDIpriorityr200:
		return "HVGICDistributorRegGICDIpriorityr200"
	case HVGICDistributorRegGICDIpriorityr201:
		return "HVGICDistributorRegGICDIpriorityr201"
	case HVGICDistributorRegGICDIpriorityr202:
		return "HVGICDistributorRegGICDIpriorityr202"
	case HVGICDistributorRegGICDIpriorityr203:
		return "HVGICDistributorRegGICDIpriorityr203"
	case HVGICDistributorRegGICDIpriorityr204:
		return "HVGICDistributorRegGICDIpriorityr204"
	case HVGICDistributorRegGICDIpriorityr205:
		return "HVGICDistributorRegGICDIpriorityr205"
	case HVGICDistributorRegGICDIpriorityr206:
		return "HVGICDistributorRegGICDIpriorityr206"
	case HVGICDistributorRegGICDIpriorityr207:
		return "HVGICDistributorRegGICDIpriorityr207"
	case HVGICDistributorRegGICDIpriorityr208:
		return "HVGICDistributorRegGICDIpriorityr208"
	case HVGICDistributorRegGICDIpriorityr209:
		return "HVGICDistributorRegGICDIpriorityr209"
	case HVGICDistributorRegGICDIpriorityr21:
		return "HVGICDistributorRegGICDIpriorityr21"
	case HVGICDistributorRegGICDIpriorityr210:
		return "HVGICDistributorRegGICDIpriorityr210"
	case HVGICDistributorRegGICDIpriorityr211:
		return "HVGICDistributorRegGICDIpriorityr211"
	case HVGICDistributorRegGICDIpriorityr212:
		return "HVGICDistributorRegGICDIpriorityr212"
	case HVGICDistributorRegGICDIpriorityr213:
		return "HVGICDistributorRegGICDIpriorityr213"
	case HVGICDistributorRegGICDIpriorityr214:
		return "HVGICDistributorRegGICDIpriorityr214"
	case HVGICDistributorRegGICDIpriorityr215:
		return "HVGICDistributorRegGICDIpriorityr215"
	case HVGICDistributorRegGICDIpriorityr216:
		return "HVGICDistributorRegGICDIpriorityr216"
	case HVGICDistributorRegGICDIpriorityr217:
		return "HVGICDistributorRegGICDIpriorityr217"
	case HVGICDistributorRegGICDIpriorityr218:
		return "HVGICDistributorRegGICDIpriorityr218"
	case HVGICDistributorRegGICDIpriorityr219:
		return "HVGICDistributorRegGICDIpriorityr219"
	case HVGICDistributorRegGICDIpriorityr22:
		return "HVGICDistributorRegGICDIpriorityr22"
	case HVGICDistributorRegGICDIpriorityr220:
		return "HVGICDistributorRegGICDIpriorityr220"
	case HVGICDistributorRegGICDIpriorityr221:
		return "HVGICDistributorRegGICDIpriorityr221"
	case HVGICDistributorRegGICDIpriorityr222:
		return "HVGICDistributorRegGICDIpriorityr222"
	case HVGICDistributorRegGICDIpriorityr223:
		return "HVGICDistributorRegGICDIpriorityr223"
	case HVGICDistributorRegGICDIpriorityr224:
		return "HVGICDistributorRegGICDIpriorityr224"
	case HVGICDistributorRegGICDIpriorityr225:
		return "HVGICDistributorRegGICDIpriorityr225"
	case HVGICDistributorRegGICDIpriorityr226:
		return "HVGICDistributorRegGICDIpriorityr226"
	case HVGICDistributorRegGICDIpriorityr227:
		return "HVGICDistributorRegGICDIpriorityr227"
	case HVGICDistributorRegGICDIpriorityr228:
		return "HVGICDistributorRegGICDIpriorityr228"
	case HVGICDistributorRegGICDIpriorityr229:
		return "HVGICDistributorRegGICDIpriorityr229"
	case HVGICDistributorRegGICDIpriorityr23:
		return "HVGICDistributorRegGICDIpriorityr23"
	case HVGICDistributorRegGICDIpriorityr230:
		return "HVGICDistributorRegGICDIpriorityr230"
	case HVGICDistributorRegGICDIpriorityr231:
		return "HVGICDistributorRegGICDIpriorityr231"
	case HVGICDistributorRegGICDIpriorityr232:
		return "HVGICDistributorRegGICDIpriorityr232"
	case HVGICDistributorRegGICDIpriorityr233:
		return "HVGICDistributorRegGICDIpriorityr233"
	case HVGICDistributorRegGICDIpriorityr234:
		return "HVGICDistributorRegGICDIpriorityr234"
	case HVGICDistributorRegGICDIpriorityr235:
		return "HVGICDistributorRegGICDIpriorityr235"
	case HVGICDistributorRegGICDIpriorityr236:
		return "HVGICDistributorRegGICDIpriorityr236"
	case HVGICDistributorRegGICDIpriorityr237:
		return "HVGICDistributorRegGICDIpriorityr237"
	case HVGICDistributorRegGICDIpriorityr238:
		return "HVGICDistributorRegGICDIpriorityr238"
	case HVGICDistributorRegGICDIpriorityr239:
		return "HVGICDistributorRegGICDIpriorityr239"
	case HVGICDistributorRegGICDIpriorityr24:
		return "HVGICDistributorRegGICDIpriorityr24"
	case HVGICDistributorRegGICDIpriorityr240:
		return "HVGICDistributorRegGICDIpriorityr240"
	case HVGICDistributorRegGICDIpriorityr241:
		return "HVGICDistributorRegGICDIpriorityr241"
	case HVGICDistributorRegGICDIpriorityr242:
		return "HVGICDistributorRegGICDIpriorityr242"
	case HVGICDistributorRegGICDIpriorityr243:
		return "HVGICDistributorRegGICDIpriorityr243"
	case HVGICDistributorRegGICDIpriorityr244:
		return "HVGICDistributorRegGICDIpriorityr244"
	case HVGICDistributorRegGICDIpriorityr245:
		return "HVGICDistributorRegGICDIpriorityr245"
	case HVGICDistributorRegGICDIpriorityr246:
		return "HVGICDistributorRegGICDIpriorityr246"
	case HVGICDistributorRegGICDIpriorityr247:
		return "HVGICDistributorRegGICDIpriorityr247"
	case HVGICDistributorRegGICDIpriorityr248:
		return "HVGICDistributorRegGICDIpriorityr248"
	case HVGICDistributorRegGICDIpriorityr249:
		return "HVGICDistributorRegGICDIpriorityr249"
	case HVGICDistributorRegGICDIpriorityr25:
		return "HVGICDistributorRegGICDIpriorityr25"
	case HVGICDistributorRegGICDIpriorityr250:
		return "HVGICDistributorRegGICDIpriorityr250"
	case HVGICDistributorRegGICDIpriorityr251:
		return "HVGICDistributorRegGICDIpriorityr251"
	case HVGICDistributorRegGICDIpriorityr252:
		return "HVGICDistributorRegGICDIpriorityr252"
	case HVGICDistributorRegGICDIpriorityr253:
		return "HVGICDistributorRegGICDIpriorityr253"
	case HVGICDistributorRegGICDIpriorityr254:
		return "HVGICDistributorRegGICDIpriorityr254"
	case HVGICDistributorRegGICDIpriorityr26:
		return "HVGICDistributorRegGICDIpriorityr26"
	case HVGICDistributorRegGICDIpriorityr27:
		return "HVGICDistributorRegGICDIpriorityr27"
	case HVGICDistributorRegGICDIpriorityr28:
		return "HVGICDistributorRegGICDIpriorityr28"
	case HVGICDistributorRegGICDIpriorityr29:
		return "HVGICDistributorRegGICDIpriorityr29"
	case HVGICDistributorRegGICDIpriorityr3:
		return "HVGICDistributorRegGICDIpriorityr3"
	case HVGICDistributorRegGICDIpriorityr30:
		return "HVGICDistributorRegGICDIpriorityr30"
	case HVGICDistributorRegGICDIpriorityr31:
		return "HVGICDistributorRegGICDIpriorityr31"
	case HVGICDistributorRegGICDIpriorityr32:
		return "HVGICDistributorRegGICDIpriorityr32"
	case HVGICDistributorRegGICDIpriorityr33:
		return "HVGICDistributorRegGICDIpriorityr33"
	case HVGICDistributorRegGICDIpriorityr34:
		return "HVGICDistributorRegGICDIpriorityr34"
	case HVGICDistributorRegGICDIpriorityr35:
		return "HVGICDistributorRegGICDIpriorityr35"
	case HVGICDistributorRegGICDIpriorityr36:
		return "HVGICDistributorRegGICDIpriorityr36"
	case HVGICDistributorRegGICDIpriorityr37:
		return "HVGICDistributorRegGICDIpriorityr37"
	case HVGICDistributorRegGICDIpriorityr38:
		return "HVGICDistributorRegGICDIpriorityr38"
	case HVGICDistributorRegGICDIpriorityr39:
		return "HVGICDistributorRegGICDIpriorityr39"
	case HVGICDistributorRegGICDIpriorityr4:
		return "HVGICDistributorRegGICDIpriorityr4"
	case HVGICDistributorRegGICDIpriorityr40:
		return "HVGICDistributorRegGICDIpriorityr40"
	case HVGICDistributorRegGICDIpriorityr41:
		return "HVGICDistributorRegGICDIpriorityr41"
	case HVGICDistributorRegGICDIpriorityr42:
		return "HVGICDistributorRegGICDIpriorityr42"
	case HVGICDistributorRegGICDIpriorityr43:
		return "HVGICDistributorRegGICDIpriorityr43"
	case HVGICDistributorRegGICDIpriorityr44:
		return "HVGICDistributorRegGICDIpriorityr44"
	case HVGICDistributorRegGICDIpriorityr45:
		return "HVGICDistributorRegGICDIpriorityr45"
	case HVGICDistributorRegGICDIpriorityr46:
		return "HVGICDistributorRegGICDIpriorityr46"
	case HVGICDistributorRegGICDIpriorityr47:
		return "HVGICDistributorRegGICDIpriorityr47"
	case HVGICDistributorRegGICDIpriorityr48:
		return "HVGICDistributorRegGICDIpriorityr48"
	case HVGICDistributorRegGICDIpriorityr49:
		return "HVGICDistributorRegGICDIpriorityr49"
	case HVGICDistributorRegGICDIpriorityr5:
		return "HVGICDistributorRegGICDIpriorityr5"
	case HVGICDistributorRegGICDIpriorityr50:
		return "HVGICDistributorRegGICDIpriorityr50"
	case HVGICDistributorRegGICDIpriorityr51:
		return "HVGICDistributorRegGICDIpriorityr51"
	case HVGICDistributorRegGICDIpriorityr52:
		return "HVGICDistributorRegGICDIpriorityr52"
	case HVGICDistributorRegGICDIpriorityr53:
		return "HVGICDistributorRegGICDIpriorityr53"
	case HVGICDistributorRegGICDIpriorityr54:
		return "HVGICDistributorRegGICDIpriorityr54"
	case HVGICDistributorRegGICDIpriorityr55:
		return "HVGICDistributorRegGICDIpriorityr55"
	case HVGICDistributorRegGICDIpriorityr56:
		return "HVGICDistributorRegGICDIpriorityr56"
	case HVGICDistributorRegGICDIpriorityr57:
		return "HVGICDistributorRegGICDIpriorityr57"
	case HVGICDistributorRegGICDIpriorityr58:
		return "HVGICDistributorRegGICDIpriorityr58"
	case HVGICDistributorRegGICDIpriorityr59:
		return "HVGICDistributorRegGICDIpriorityr59"
	case HVGICDistributorRegGICDIpriorityr6:
		return "HVGICDistributorRegGICDIpriorityr6"
	case HVGICDistributorRegGICDIpriorityr60:
		return "HVGICDistributorRegGICDIpriorityr60"
	case HVGICDistributorRegGICDIpriorityr61:
		return "HVGICDistributorRegGICDIpriorityr61"
	case HVGICDistributorRegGICDIpriorityr62:
		return "HVGICDistributorRegGICDIpriorityr62"
	case HVGICDistributorRegGICDIpriorityr63:
		return "HVGICDistributorRegGICDIpriorityr63"
	case HVGICDistributorRegGICDIpriorityr64:
		return "HVGICDistributorRegGICDIpriorityr64"
	case HVGICDistributorRegGICDIpriorityr65:
		return "HVGICDistributorRegGICDIpriorityr65"
	case HVGICDistributorRegGICDIpriorityr66:
		return "HVGICDistributorRegGICDIpriorityr66"
	case HVGICDistributorRegGICDIpriorityr67:
		return "HVGICDistributorRegGICDIpriorityr67"
	case HVGICDistributorRegGICDIpriorityr68:
		return "HVGICDistributorRegGICDIpriorityr68"
	case HVGICDistributorRegGICDIpriorityr69:
		return "HVGICDistributorRegGICDIpriorityr69"
	case HVGICDistributorRegGICDIpriorityr7:
		return "HVGICDistributorRegGICDIpriorityr7"
	case HVGICDistributorRegGICDIpriorityr70:
		return "HVGICDistributorRegGICDIpriorityr70"
	case HVGICDistributorRegGICDIpriorityr71:
		return "HVGICDistributorRegGICDIpriorityr71"
	case HVGICDistributorRegGICDIpriorityr72:
		return "HVGICDistributorRegGICDIpriorityr72"
	case HVGICDistributorRegGICDIpriorityr73:
		return "HVGICDistributorRegGICDIpriorityr73"
	case HVGICDistributorRegGICDIpriorityr74:
		return "HVGICDistributorRegGICDIpriorityr74"
	case HVGICDistributorRegGICDIpriorityr75:
		return "HVGICDistributorRegGICDIpriorityr75"
	case HVGICDistributorRegGICDIpriorityr76:
		return "HVGICDistributorRegGICDIpriorityr76"
	case HVGICDistributorRegGICDIpriorityr77:
		return "HVGICDistributorRegGICDIpriorityr77"
	case HVGICDistributorRegGICDIpriorityr78:
		return "HVGICDistributorRegGICDIpriorityr78"
	case HVGICDistributorRegGICDIpriorityr79:
		return "HVGICDistributorRegGICDIpriorityr79"
	case HVGICDistributorRegGICDIpriorityr8:
		return "HVGICDistributorRegGICDIpriorityr8"
	case HVGICDistributorRegGICDIpriorityr80:
		return "HVGICDistributorRegGICDIpriorityr80"
	case HVGICDistributorRegGICDIpriorityr81:
		return "HVGICDistributorRegGICDIpriorityr81"
	case HVGICDistributorRegGICDIpriorityr82:
		return "HVGICDistributorRegGICDIpriorityr82"
	case HVGICDistributorRegGICDIpriorityr83:
		return "HVGICDistributorRegGICDIpriorityr83"
	case HVGICDistributorRegGICDIpriorityr84:
		return "HVGICDistributorRegGICDIpriorityr84"
	case HVGICDistributorRegGICDIpriorityr85:
		return "HVGICDistributorRegGICDIpriorityr85"
	case HVGICDistributorRegGICDIpriorityr86:
		return "HVGICDistributorRegGICDIpriorityr86"
	case HVGICDistributorRegGICDIpriorityr87:
		return "HVGICDistributorRegGICDIpriorityr87"
	case HVGICDistributorRegGICDIpriorityr88:
		return "HVGICDistributorRegGICDIpriorityr88"
	case HVGICDistributorRegGICDIpriorityr89:
		return "HVGICDistributorRegGICDIpriorityr89"
	case HVGICDistributorRegGICDIpriorityr9:
		return "HVGICDistributorRegGICDIpriorityr9"
	case HVGICDistributorRegGICDIpriorityr90:
		return "HVGICDistributorRegGICDIpriorityr90"
	case HVGICDistributorRegGICDIpriorityr91:
		return "HVGICDistributorRegGICDIpriorityr91"
	case HVGICDistributorRegGICDIpriorityr92:
		return "HVGICDistributorRegGICDIpriorityr92"
	case HVGICDistributorRegGICDIpriorityr93:
		return "HVGICDistributorRegGICDIpriorityr93"
	case HVGICDistributorRegGICDIpriorityr94:
		return "HVGICDistributorRegGICDIpriorityr94"
	case HVGICDistributorRegGICDIpriorityr95:
		return "HVGICDistributorRegGICDIpriorityr95"
	case HVGICDistributorRegGICDIpriorityr96:
		return "HVGICDistributorRegGICDIpriorityr96"
	case HVGICDistributorRegGICDIpriorityr97:
		return "HVGICDistributorRegGICDIpriorityr97"
	case HVGICDistributorRegGICDIpriorityr98:
		return "HVGICDistributorRegGICDIpriorityr98"
	case HVGICDistributorRegGICDIpriorityr99:
		return "HVGICDistributorRegGICDIpriorityr99"
	case HVGICDistributorRegGICDIrouter100:
		return "HVGICDistributorRegGICDIrouter100"
	case HVGICDistributorRegGICDIrouter1000:
		return "HVGICDistributorRegGICDIrouter1000"
	case HVGICDistributorRegGICDIrouter1001:
		return "HVGICDistributorRegGICDIrouter1001"
	case HVGICDistributorRegGICDIrouter1002:
		return "HVGICDistributorRegGICDIrouter1002"
	case HVGICDistributorRegGICDIrouter1003:
		return "HVGICDistributorRegGICDIrouter1003"
	case HVGICDistributorRegGICDIrouter1004:
		return "HVGICDistributorRegGICDIrouter1004"
	case HVGICDistributorRegGICDIrouter1005:
		return "HVGICDistributorRegGICDIrouter1005"
	case HVGICDistributorRegGICDIrouter1006:
		return "HVGICDistributorRegGICDIrouter1006"
	case HVGICDistributorRegGICDIrouter1007:
		return "HVGICDistributorRegGICDIrouter1007"
	case HVGICDistributorRegGICDIrouter1008:
		return "HVGICDistributorRegGICDIrouter1008"
	case HVGICDistributorRegGICDIrouter1009:
		return "HVGICDistributorRegGICDIrouter1009"
	case HVGICDistributorRegGICDIrouter101:
		return "HVGICDistributorRegGICDIrouter101"
	case HVGICDistributorRegGICDIrouter1010:
		return "HVGICDistributorRegGICDIrouter1010"
	case HVGICDistributorRegGICDIrouter1011:
		return "HVGICDistributorRegGICDIrouter1011"
	case HVGICDistributorRegGICDIrouter1012:
		return "HVGICDistributorRegGICDIrouter1012"
	case HVGICDistributorRegGICDIrouter1013:
		return "HVGICDistributorRegGICDIrouter1013"
	case HVGICDistributorRegGICDIrouter1014:
		return "HVGICDistributorRegGICDIrouter1014"
	case HVGICDistributorRegGICDIrouter1015:
		return "HVGICDistributorRegGICDIrouter1015"
	case HVGICDistributorRegGICDIrouter1016:
		return "HVGICDistributorRegGICDIrouter1016"
	case HVGICDistributorRegGICDIrouter1017:
		return "HVGICDistributorRegGICDIrouter1017"
	case HVGICDistributorRegGICDIrouter1018:
		return "HVGICDistributorRegGICDIrouter1018"
	case HVGICDistributorRegGICDIrouter1019:
		return "HVGICDistributorRegGICDIrouter1019"
	case HVGICDistributorRegGICDIrouter102:
		return "HVGICDistributorRegGICDIrouter102"
	case HVGICDistributorRegGICDIrouter103:
		return "HVGICDistributorRegGICDIrouter103"
	case HVGICDistributorRegGICDIrouter104:
		return "HVGICDistributorRegGICDIrouter104"
	case HVGICDistributorRegGICDIrouter105:
		return "HVGICDistributorRegGICDIrouter105"
	case HVGICDistributorRegGICDIrouter106:
		return "HVGICDistributorRegGICDIrouter106"
	case HVGICDistributorRegGICDIrouter107:
		return "HVGICDistributorRegGICDIrouter107"
	case HVGICDistributorRegGICDIrouter108:
		return "HVGICDistributorRegGICDIrouter108"
	case HVGICDistributorRegGICDIrouter109:
		return "HVGICDistributorRegGICDIrouter109"
	case HVGICDistributorRegGICDIrouter110:
		return "HVGICDistributorRegGICDIrouter110"
	case HVGICDistributorRegGICDIrouter111:
		return "HVGICDistributorRegGICDIrouter111"
	case HVGICDistributorRegGICDIrouter112:
		return "HVGICDistributorRegGICDIrouter112"
	case HVGICDistributorRegGICDIrouter113:
		return "HVGICDistributorRegGICDIrouter113"
	case HVGICDistributorRegGICDIrouter114:
		return "HVGICDistributorRegGICDIrouter114"
	case HVGICDistributorRegGICDIrouter115:
		return "HVGICDistributorRegGICDIrouter115"
	case HVGICDistributorRegGICDIrouter116:
		return "HVGICDistributorRegGICDIrouter116"
	case HVGICDistributorRegGICDIrouter117:
		return "HVGICDistributorRegGICDIrouter117"
	case HVGICDistributorRegGICDIrouter118:
		return "HVGICDistributorRegGICDIrouter118"
	case HVGICDistributorRegGICDIrouter119:
		return "HVGICDistributorRegGICDIrouter119"
	case HVGICDistributorRegGICDIrouter120:
		return "HVGICDistributorRegGICDIrouter120"
	case HVGICDistributorRegGICDIrouter121:
		return "HVGICDistributorRegGICDIrouter121"
	case HVGICDistributorRegGICDIrouter122:
		return "HVGICDistributorRegGICDIrouter122"
	case HVGICDistributorRegGICDIrouter123:
		return "HVGICDistributorRegGICDIrouter123"
	case HVGICDistributorRegGICDIrouter124:
		return "HVGICDistributorRegGICDIrouter124"
	case HVGICDistributorRegGICDIrouter125:
		return "HVGICDistributorRegGICDIrouter125"
	case HVGICDistributorRegGICDIrouter126:
		return "HVGICDistributorRegGICDIrouter126"
	case HVGICDistributorRegGICDIrouter127:
		return "HVGICDistributorRegGICDIrouter127"
	case HVGICDistributorRegGICDIrouter128:
		return "HVGICDistributorRegGICDIrouter128"
	case HVGICDistributorRegGICDIrouter129:
		return "HVGICDistributorRegGICDIrouter129"
	case HVGICDistributorRegGICDIrouter130:
		return "HVGICDistributorRegGICDIrouter130"
	case HVGICDistributorRegGICDIrouter131:
		return "HVGICDistributorRegGICDIrouter131"
	case HVGICDistributorRegGICDIrouter132:
		return "HVGICDistributorRegGICDIrouter132"
	case HVGICDistributorRegGICDIrouter133:
		return "HVGICDistributorRegGICDIrouter133"
	case HVGICDistributorRegGICDIrouter134:
		return "HVGICDistributorRegGICDIrouter134"
	case HVGICDistributorRegGICDIrouter135:
		return "HVGICDistributorRegGICDIrouter135"
	case HVGICDistributorRegGICDIrouter136:
		return "HVGICDistributorRegGICDIrouter136"
	case HVGICDistributorRegGICDIrouter137:
		return "HVGICDistributorRegGICDIrouter137"
	case HVGICDistributorRegGICDIrouter138:
		return "HVGICDistributorRegGICDIrouter138"
	case HVGICDistributorRegGICDIrouter139:
		return "HVGICDistributorRegGICDIrouter139"
	case HVGICDistributorRegGICDIrouter140:
		return "HVGICDistributorRegGICDIrouter140"
	case HVGICDistributorRegGICDIrouter141:
		return "HVGICDistributorRegGICDIrouter141"
	case HVGICDistributorRegGICDIrouter142:
		return "HVGICDistributorRegGICDIrouter142"
	case HVGICDistributorRegGICDIrouter143:
		return "HVGICDistributorRegGICDIrouter143"
	case HVGICDistributorRegGICDIrouter144:
		return "HVGICDistributorRegGICDIrouter144"
	case HVGICDistributorRegGICDIrouter145:
		return "HVGICDistributorRegGICDIrouter145"
	case HVGICDistributorRegGICDIrouter146:
		return "HVGICDistributorRegGICDIrouter146"
	case HVGICDistributorRegGICDIrouter147:
		return "HVGICDistributorRegGICDIrouter147"
	case HVGICDistributorRegGICDIrouter148:
		return "HVGICDistributorRegGICDIrouter148"
	case HVGICDistributorRegGICDIrouter149:
		return "HVGICDistributorRegGICDIrouter149"
	case HVGICDistributorRegGICDIrouter150:
		return "HVGICDistributorRegGICDIrouter150"
	case HVGICDistributorRegGICDIrouter151:
		return "HVGICDistributorRegGICDIrouter151"
	case HVGICDistributorRegGICDIrouter152:
		return "HVGICDistributorRegGICDIrouter152"
	case HVGICDistributorRegGICDIrouter153:
		return "HVGICDistributorRegGICDIrouter153"
	case HVGICDistributorRegGICDIrouter154:
		return "HVGICDistributorRegGICDIrouter154"
	case HVGICDistributorRegGICDIrouter155:
		return "HVGICDistributorRegGICDIrouter155"
	case HVGICDistributorRegGICDIrouter156:
		return "HVGICDistributorRegGICDIrouter156"
	case HVGICDistributorRegGICDIrouter157:
		return "HVGICDistributorRegGICDIrouter157"
	case HVGICDistributorRegGICDIrouter158:
		return "HVGICDistributorRegGICDIrouter158"
	case HVGICDistributorRegGICDIrouter159:
		return "HVGICDistributorRegGICDIrouter159"
	case HVGICDistributorRegGICDIrouter160:
		return "HVGICDistributorRegGICDIrouter160"
	case HVGICDistributorRegGICDIrouter161:
		return "HVGICDistributorRegGICDIrouter161"
	case HVGICDistributorRegGICDIrouter162:
		return "HVGICDistributorRegGICDIrouter162"
	case HVGICDistributorRegGICDIrouter163:
		return "HVGICDistributorRegGICDIrouter163"
	case HVGICDistributorRegGICDIrouter164:
		return "HVGICDistributorRegGICDIrouter164"
	case HVGICDistributorRegGICDIrouter165:
		return "HVGICDistributorRegGICDIrouter165"
	case HVGICDistributorRegGICDIrouter166:
		return "HVGICDistributorRegGICDIrouter166"
	case HVGICDistributorRegGICDIrouter167:
		return "HVGICDistributorRegGICDIrouter167"
	case HVGICDistributorRegGICDIrouter168:
		return "HVGICDistributorRegGICDIrouter168"
	case HVGICDistributorRegGICDIrouter169:
		return "HVGICDistributorRegGICDIrouter169"
	case HVGICDistributorRegGICDIrouter170:
		return "HVGICDistributorRegGICDIrouter170"
	case HVGICDistributorRegGICDIrouter171:
		return "HVGICDistributorRegGICDIrouter171"
	case HVGICDistributorRegGICDIrouter172:
		return "HVGICDistributorRegGICDIrouter172"
	case HVGICDistributorRegGICDIrouter173:
		return "HVGICDistributorRegGICDIrouter173"
	case HVGICDistributorRegGICDIrouter174:
		return "HVGICDistributorRegGICDIrouter174"
	case HVGICDistributorRegGICDIrouter175:
		return "HVGICDistributorRegGICDIrouter175"
	case HVGICDistributorRegGICDIrouter176:
		return "HVGICDistributorRegGICDIrouter176"
	case HVGICDistributorRegGICDIrouter177:
		return "HVGICDistributorRegGICDIrouter177"
	case HVGICDistributorRegGICDIrouter178:
		return "HVGICDistributorRegGICDIrouter178"
	case HVGICDistributorRegGICDIrouter179:
		return "HVGICDistributorRegGICDIrouter179"
	case HVGICDistributorRegGICDIrouter180:
		return "HVGICDistributorRegGICDIrouter180"
	case HVGICDistributorRegGICDIrouter181:
		return "HVGICDistributorRegGICDIrouter181"
	case HVGICDistributorRegGICDIrouter182:
		return "HVGICDistributorRegGICDIrouter182"
	case HVGICDistributorRegGICDIrouter183:
		return "HVGICDistributorRegGICDIrouter183"
	case HVGICDistributorRegGICDIrouter184:
		return "HVGICDistributorRegGICDIrouter184"
	case HVGICDistributorRegGICDIrouter185:
		return "HVGICDistributorRegGICDIrouter185"
	case HVGICDistributorRegGICDIrouter186:
		return "HVGICDistributorRegGICDIrouter186"
	case HVGICDistributorRegGICDIrouter187:
		return "HVGICDistributorRegGICDIrouter187"
	case HVGICDistributorRegGICDIrouter188:
		return "HVGICDistributorRegGICDIrouter188"
	case HVGICDistributorRegGICDIrouter189:
		return "HVGICDistributorRegGICDIrouter189"
	case HVGICDistributorRegGICDIrouter190:
		return "HVGICDistributorRegGICDIrouter190"
	case HVGICDistributorRegGICDIrouter191:
		return "HVGICDistributorRegGICDIrouter191"
	case HVGICDistributorRegGICDIrouter192:
		return "HVGICDistributorRegGICDIrouter192"
	case HVGICDistributorRegGICDIrouter193:
		return "HVGICDistributorRegGICDIrouter193"
	case HVGICDistributorRegGICDIrouter194:
		return "HVGICDistributorRegGICDIrouter194"
	case HVGICDistributorRegGICDIrouter195:
		return "HVGICDistributorRegGICDIrouter195"
	case HVGICDistributorRegGICDIrouter196:
		return "HVGICDistributorRegGICDIrouter196"
	case HVGICDistributorRegGICDIrouter197:
		return "HVGICDistributorRegGICDIrouter197"
	case HVGICDistributorRegGICDIrouter198:
		return "HVGICDistributorRegGICDIrouter198"
	case HVGICDistributorRegGICDIrouter199:
		return "HVGICDistributorRegGICDIrouter199"
	case HVGICDistributorRegGICDIrouter200:
		return "HVGICDistributorRegGICDIrouter200"
	case HVGICDistributorRegGICDIrouter201:
		return "HVGICDistributorRegGICDIrouter201"
	case HVGICDistributorRegGICDIrouter202:
		return "HVGICDistributorRegGICDIrouter202"
	case HVGICDistributorRegGICDIrouter203:
		return "HVGICDistributorRegGICDIrouter203"
	case HVGICDistributorRegGICDIrouter204:
		return "HVGICDistributorRegGICDIrouter204"
	case HVGICDistributorRegGICDIrouter205:
		return "HVGICDistributorRegGICDIrouter205"
	case HVGICDistributorRegGICDIrouter206:
		return "HVGICDistributorRegGICDIrouter206"
	case HVGICDistributorRegGICDIrouter207:
		return "HVGICDistributorRegGICDIrouter207"
	case HVGICDistributorRegGICDIrouter208:
		return "HVGICDistributorRegGICDIrouter208"
	case HVGICDistributorRegGICDIrouter209:
		return "HVGICDistributorRegGICDIrouter209"
	case HVGICDistributorRegGICDIrouter210:
		return "HVGICDistributorRegGICDIrouter210"
	case HVGICDistributorRegGICDIrouter211:
		return "HVGICDistributorRegGICDIrouter211"
	case HVGICDistributorRegGICDIrouter212:
		return "HVGICDistributorRegGICDIrouter212"
	case HVGICDistributorRegGICDIrouter213:
		return "HVGICDistributorRegGICDIrouter213"
	case HVGICDistributorRegGICDIrouter214:
		return "HVGICDistributorRegGICDIrouter214"
	case HVGICDistributorRegGICDIrouter215:
		return "HVGICDistributorRegGICDIrouter215"
	case HVGICDistributorRegGICDIrouter216:
		return "HVGICDistributorRegGICDIrouter216"
	case HVGICDistributorRegGICDIrouter217:
		return "HVGICDistributorRegGICDIrouter217"
	case HVGICDistributorRegGICDIrouter218:
		return "HVGICDistributorRegGICDIrouter218"
	case HVGICDistributorRegGICDIrouter219:
		return "HVGICDistributorRegGICDIrouter219"
	case HVGICDistributorRegGICDIrouter220:
		return "HVGICDistributorRegGICDIrouter220"
	case HVGICDistributorRegGICDIrouter221:
		return "HVGICDistributorRegGICDIrouter221"
	case HVGICDistributorRegGICDIrouter222:
		return "HVGICDistributorRegGICDIrouter222"
	case HVGICDistributorRegGICDIrouter223:
		return "HVGICDistributorRegGICDIrouter223"
	case HVGICDistributorRegGICDIrouter224:
		return "HVGICDistributorRegGICDIrouter224"
	case HVGICDistributorRegGICDIrouter225:
		return "HVGICDistributorRegGICDIrouter225"
	case HVGICDistributorRegGICDIrouter226:
		return "HVGICDistributorRegGICDIrouter226"
	case HVGICDistributorRegGICDIrouter227:
		return "HVGICDistributorRegGICDIrouter227"
	case HVGICDistributorRegGICDIrouter228:
		return "HVGICDistributorRegGICDIrouter228"
	case HVGICDistributorRegGICDIrouter229:
		return "HVGICDistributorRegGICDIrouter229"
	case HVGICDistributorRegGICDIrouter230:
		return "HVGICDistributorRegGICDIrouter230"
	case HVGICDistributorRegGICDIrouter231:
		return "HVGICDistributorRegGICDIrouter231"
	case HVGICDistributorRegGICDIrouter232:
		return "HVGICDistributorRegGICDIrouter232"
	case HVGICDistributorRegGICDIrouter233:
		return "HVGICDistributorRegGICDIrouter233"
	case HVGICDistributorRegGICDIrouter234:
		return "HVGICDistributorRegGICDIrouter234"
	case HVGICDistributorRegGICDIrouter235:
		return "HVGICDistributorRegGICDIrouter235"
	case HVGICDistributorRegGICDIrouter236:
		return "HVGICDistributorRegGICDIrouter236"
	case HVGICDistributorRegGICDIrouter237:
		return "HVGICDistributorRegGICDIrouter237"
	case HVGICDistributorRegGICDIrouter238:
		return "HVGICDistributorRegGICDIrouter238"
	case HVGICDistributorRegGICDIrouter239:
		return "HVGICDistributorRegGICDIrouter239"
	case HVGICDistributorRegGICDIrouter240:
		return "HVGICDistributorRegGICDIrouter240"
	case HVGICDistributorRegGICDIrouter241:
		return "HVGICDistributorRegGICDIrouter241"
	case HVGICDistributorRegGICDIrouter242:
		return "HVGICDistributorRegGICDIrouter242"
	case HVGICDistributorRegGICDIrouter243:
		return "HVGICDistributorRegGICDIrouter243"
	case HVGICDistributorRegGICDIrouter244:
		return "HVGICDistributorRegGICDIrouter244"
	case HVGICDistributorRegGICDIrouter245:
		return "HVGICDistributorRegGICDIrouter245"
	case HVGICDistributorRegGICDIrouter246:
		return "HVGICDistributorRegGICDIrouter246"
	case HVGICDistributorRegGICDIrouter247:
		return "HVGICDistributorRegGICDIrouter247"
	case HVGICDistributorRegGICDIrouter248:
		return "HVGICDistributorRegGICDIrouter248"
	case HVGICDistributorRegGICDIrouter249:
		return "HVGICDistributorRegGICDIrouter249"
	case HVGICDistributorRegGICDIrouter250:
		return "HVGICDistributorRegGICDIrouter250"
	case HVGICDistributorRegGICDIrouter251:
		return "HVGICDistributorRegGICDIrouter251"
	case HVGICDistributorRegGICDIrouter252:
		return "HVGICDistributorRegGICDIrouter252"
	case HVGICDistributorRegGICDIrouter253:
		return "HVGICDistributorRegGICDIrouter253"
	case HVGICDistributorRegGICDIrouter254:
		return "HVGICDistributorRegGICDIrouter254"
	case HVGICDistributorRegGICDIrouter255:
		return "HVGICDistributorRegGICDIrouter255"
	case HVGICDistributorRegGICDIrouter256:
		return "HVGICDistributorRegGICDIrouter256"
	case HVGICDistributorRegGICDIrouter257:
		return "HVGICDistributorRegGICDIrouter257"
	case HVGICDistributorRegGICDIrouter258:
		return "HVGICDistributorRegGICDIrouter258"
	case HVGICDistributorRegGICDIrouter259:
		return "HVGICDistributorRegGICDIrouter259"
	case HVGICDistributorRegGICDIrouter260:
		return "HVGICDistributorRegGICDIrouter260"
	case HVGICDistributorRegGICDIrouter261:
		return "HVGICDistributorRegGICDIrouter261"
	case HVGICDistributorRegGICDIrouter262:
		return "HVGICDistributorRegGICDIrouter262"
	case HVGICDistributorRegGICDIrouter263:
		return "HVGICDistributorRegGICDIrouter263"
	case HVGICDistributorRegGICDIrouter264:
		return "HVGICDistributorRegGICDIrouter264"
	case HVGICDistributorRegGICDIrouter265:
		return "HVGICDistributorRegGICDIrouter265"
	case HVGICDistributorRegGICDIrouter266:
		return "HVGICDistributorRegGICDIrouter266"
	case HVGICDistributorRegGICDIrouter267:
		return "HVGICDistributorRegGICDIrouter267"
	case HVGICDistributorRegGICDIrouter268:
		return "HVGICDistributorRegGICDIrouter268"
	case HVGICDistributorRegGICDIrouter269:
		return "HVGICDistributorRegGICDIrouter269"
	case HVGICDistributorRegGICDIrouter270:
		return "HVGICDistributorRegGICDIrouter270"
	case HVGICDistributorRegGICDIrouter271:
		return "HVGICDistributorRegGICDIrouter271"
	case HVGICDistributorRegGICDIrouter272:
		return "HVGICDistributorRegGICDIrouter272"
	case HVGICDistributorRegGICDIrouter273:
		return "HVGICDistributorRegGICDIrouter273"
	case HVGICDistributorRegGICDIrouter274:
		return "HVGICDistributorRegGICDIrouter274"
	case HVGICDistributorRegGICDIrouter275:
		return "HVGICDistributorRegGICDIrouter275"
	case HVGICDistributorRegGICDIrouter276:
		return "HVGICDistributorRegGICDIrouter276"
	case HVGICDistributorRegGICDIrouter277:
		return "HVGICDistributorRegGICDIrouter277"
	case HVGICDistributorRegGICDIrouter278:
		return "HVGICDistributorRegGICDIrouter278"
	case HVGICDistributorRegGICDIrouter279:
		return "HVGICDistributorRegGICDIrouter279"
	case HVGICDistributorRegGICDIrouter280:
		return "HVGICDistributorRegGICDIrouter280"
	case HVGICDistributorRegGICDIrouter281:
		return "HVGICDistributorRegGICDIrouter281"
	case HVGICDistributorRegGICDIrouter282:
		return "HVGICDistributorRegGICDIrouter282"
	case HVGICDistributorRegGICDIrouter283:
		return "HVGICDistributorRegGICDIrouter283"
	case HVGICDistributorRegGICDIrouter284:
		return "HVGICDistributorRegGICDIrouter284"
	case HVGICDistributorRegGICDIrouter285:
		return "HVGICDistributorRegGICDIrouter285"
	case HVGICDistributorRegGICDIrouter286:
		return "HVGICDistributorRegGICDIrouter286"
	case HVGICDistributorRegGICDIrouter287:
		return "HVGICDistributorRegGICDIrouter287"
	case HVGICDistributorRegGICDIrouter288:
		return "HVGICDistributorRegGICDIrouter288"
	case HVGICDistributorRegGICDIrouter289:
		return "HVGICDistributorRegGICDIrouter289"
	case HVGICDistributorRegGICDIrouter290:
		return "HVGICDistributorRegGICDIrouter290"
	case HVGICDistributorRegGICDIrouter291:
		return "HVGICDistributorRegGICDIrouter291"
	case HVGICDistributorRegGICDIrouter292:
		return "HVGICDistributorRegGICDIrouter292"
	case HVGICDistributorRegGICDIrouter293:
		return "HVGICDistributorRegGICDIrouter293"
	case HVGICDistributorRegGICDIrouter294:
		return "HVGICDistributorRegGICDIrouter294"
	case HVGICDistributorRegGICDIrouter295:
		return "HVGICDistributorRegGICDIrouter295"
	case HVGICDistributorRegGICDIrouter296:
		return "HVGICDistributorRegGICDIrouter296"
	case HVGICDistributorRegGICDIrouter297:
		return "HVGICDistributorRegGICDIrouter297"
	case HVGICDistributorRegGICDIrouter298:
		return "HVGICDistributorRegGICDIrouter298"
	case HVGICDistributorRegGICDIrouter299:
		return "HVGICDistributorRegGICDIrouter299"
	case HVGICDistributorRegGICDIrouter300:
		return "HVGICDistributorRegGICDIrouter300"
	case HVGICDistributorRegGICDIrouter301:
		return "HVGICDistributorRegGICDIrouter301"
	case HVGICDistributorRegGICDIrouter302:
		return "HVGICDistributorRegGICDIrouter302"
	case HVGICDistributorRegGICDIrouter303:
		return "HVGICDistributorRegGICDIrouter303"
	case HVGICDistributorRegGICDIrouter304:
		return "HVGICDistributorRegGICDIrouter304"
	case HVGICDistributorRegGICDIrouter305:
		return "HVGICDistributorRegGICDIrouter305"
	case HVGICDistributorRegGICDIrouter306:
		return "HVGICDistributorRegGICDIrouter306"
	case HVGICDistributorRegGICDIrouter307:
		return "HVGICDistributorRegGICDIrouter307"
	case HVGICDistributorRegGICDIrouter308:
		return "HVGICDistributorRegGICDIrouter308"
	case HVGICDistributorRegGICDIrouter309:
		return "HVGICDistributorRegGICDIrouter309"
	case HVGICDistributorRegGICDIrouter310:
		return "HVGICDistributorRegGICDIrouter310"
	case HVGICDistributorRegGICDIrouter311:
		return "HVGICDistributorRegGICDIrouter311"
	case HVGICDistributorRegGICDIrouter312:
		return "HVGICDistributorRegGICDIrouter312"
	case HVGICDistributorRegGICDIrouter313:
		return "HVGICDistributorRegGICDIrouter313"
	case HVGICDistributorRegGICDIrouter314:
		return "HVGICDistributorRegGICDIrouter314"
	case HVGICDistributorRegGICDIrouter315:
		return "HVGICDistributorRegGICDIrouter315"
	case HVGICDistributorRegGICDIrouter316:
		return "HVGICDistributorRegGICDIrouter316"
	case HVGICDistributorRegGICDIrouter317:
		return "HVGICDistributorRegGICDIrouter317"
	case HVGICDistributorRegGICDIrouter318:
		return "HVGICDistributorRegGICDIrouter318"
	case HVGICDistributorRegGICDIrouter319:
		return "HVGICDistributorRegGICDIrouter319"
	case HVGICDistributorRegGICDIrouter32:
		return "HVGICDistributorRegGICDIrouter32"
	case HVGICDistributorRegGICDIrouter320:
		return "HVGICDistributorRegGICDIrouter320"
	case HVGICDistributorRegGICDIrouter321:
		return "HVGICDistributorRegGICDIrouter321"
	case HVGICDistributorRegGICDIrouter322:
		return "HVGICDistributorRegGICDIrouter322"
	case HVGICDistributorRegGICDIrouter323:
		return "HVGICDistributorRegGICDIrouter323"
	case HVGICDistributorRegGICDIrouter324:
		return "HVGICDistributorRegGICDIrouter324"
	case HVGICDistributorRegGICDIrouter325:
		return "HVGICDistributorRegGICDIrouter325"
	case HVGICDistributorRegGICDIrouter326:
		return "HVGICDistributorRegGICDIrouter326"
	case HVGICDistributorRegGICDIrouter327:
		return "HVGICDistributorRegGICDIrouter327"
	case HVGICDistributorRegGICDIrouter328:
		return "HVGICDistributorRegGICDIrouter328"
	case HVGICDistributorRegGICDIrouter329:
		return "HVGICDistributorRegGICDIrouter329"
	case HVGICDistributorRegGICDIrouter33:
		return "HVGICDistributorRegGICDIrouter33"
	case HVGICDistributorRegGICDIrouter330:
		return "HVGICDistributorRegGICDIrouter330"
	case HVGICDistributorRegGICDIrouter331:
		return "HVGICDistributorRegGICDIrouter331"
	case HVGICDistributorRegGICDIrouter332:
		return "HVGICDistributorRegGICDIrouter332"
	case HVGICDistributorRegGICDIrouter333:
		return "HVGICDistributorRegGICDIrouter333"
	case HVGICDistributorRegGICDIrouter334:
		return "HVGICDistributorRegGICDIrouter334"
	case HVGICDistributorRegGICDIrouter335:
		return "HVGICDistributorRegGICDIrouter335"
	case HVGICDistributorRegGICDIrouter336:
		return "HVGICDistributorRegGICDIrouter336"
	case HVGICDistributorRegGICDIrouter337:
		return "HVGICDistributorRegGICDIrouter337"
	case HVGICDistributorRegGICDIrouter338:
		return "HVGICDistributorRegGICDIrouter338"
	case HVGICDistributorRegGICDIrouter339:
		return "HVGICDistributorRegGICDIrouter339"
	case HVGICDistributorRegGICDIrouter34:
		return "HVGICDistributorRegGICDIrouter34"
	case HVGICDistributorRegGICDIrouter340:
		return "HVGICDistributorRegGICDIrouter340"
	case HVGICDistributorRegGICDIrouter341:
		return "HVGICDistributorRegGICDIrouter341"
	case HVGICDistributorRegGICDIrouter342:
		return "HVGICDistributorRegGICDIrouter342"
	case HVGICDistributorRegGICDIrouter343:
		return "HVGICDistributorRegGICDIrouter343"
	case HVGICDistributorRegGICDIrouter344:
		return "HVGICDistributorRegGICDIrouter344"
	case HVGICDistributorRegGICDIrouter345:
		return "HVGICDistributorRegGICDIrouter345"
	case HVGICDistributorRegGICDIrouter346:
		return "HVGICDistributorRegGICDIrouter346"
	case HVGICDistributorRegGICDIrouter347:
		return "HVGICDistributorRegGICDIrouter347"
	case HVGICDistributorRegGICDIrouter348:
		return "HVGICDistributorRegGICDIrouter348"
	case HVGICDistributorRegGICDIrouter349:
		return "HVGICDistributorRegGICDIrouter349"
	case HVGICDistributorRegGICDIrouter35:
		return "HVGICDistributorRegGICDIrouter35"
	case HVGICDistributorRegGICDIrouter350:
		return "HVGICDistributorRegGICDIrouter350"
	case HVGICDistributorRegGICDIrouter351:
		return "HVGICDistributorRegGICDIrouter351"
	case HVGICDistributorRegGICDIrouter352:
		return "HVGICDistributorRegGICDIrouter352"
	case HVGICDistributorRegGICDIrouter353:
		return "HVGICDistributorRegGICDIrouter353"
	case HVGICDistributorRegGICDIrouter354:
		return "HVGICDistributorRegGICDIrouter354"
	case HVGICDistributorRegGICDIrouter355:
		return "HVGICDistributorRegGICDIrouter355"
	case HVGICDistributorRegGICDIrouter356:
		return "HVGICDistributorRegGICDIrouter356"
	case HVGICDistributorRegGICDIrouter357:
		return "HVGICDistributorRegGICDIrouter357"
	case HVGICDistributorRegGICDIrouter358:
		return "HVGICDistributorRegGICDIrouter358"
	case HVGICDistributorRegGICDIrouter359:
		return "HVGICDistributorRegGICDIrouter359"
	case HVGICDistributorRegGICDIrouter36:
		return "HVGICDistributorRegGICDIrouter36"
	case HVGICDistributorRegGICDIrouter360:
		return "HVGICDistributorRegGICDIrouter360"
	case HVGICDistributorRegGICDIrouter361:
		return "HVGICDistributorRegGICDIrouter361"
	case HVGICDistributorRegGICDIrouter362:
		return "HVGICDistributorRegGICDIrouter362"
	case HVGICDistributorRegGICDIrouter363:
		return "HVGICDistributorRegGICDIrouter363"
	case HVGICDistributorRegGICDIrouter364:
		return "HVGICDistributorRegGICDIrouter364"
	case HVGICDistributorRegGICDIrouter365:
		return "HVGICDistributorRegGICDIrouter365"
	case HVGICDistributorRegGICDIrouter366:
		return "HVGICDistributorRegGICDIrouter366"
	case HVGICDistributorRegGICDIrouter367:
		return "HVGICDistributorRegGICDIrouter367"
	case HVGICDistributorRegGICDIrouter368:
		return "HVGICDistributorRegGICDIrouter368"
	case HVGICDistributorRegGICDIrouter369:
		return "HVGICDistributorRegGICDIrouter369"
	case HVGICDistributorRegGICDIrouter37:
		return "HVGICDistributorRegGICDIrouter37"
	case HVGICDistributorRegGICDIrouter370:
		return "HVGICDistributorRegGICDIrouter370"
	case HVGICDistributorRegGICDIrouter371:
		return "HVGICDistributorRegGICDIrouter371"
	case HVGICDistributorRegGICDIrouter372:
		return "HVGICDistributorRegGICDIrouter372"
	case HVGICDistributorRegGICDIrouter373:
		return "HVGICDistributorRegGICDIrouter373"
	case HVGICDistributorRegGICDIrouter374:
		return "HVGICDistributorRegGICDIrouter374"
	case HVGICDistributorRegGICDIrouter375:
		return "HVGICDistributorRegGICDIrouter375"
	case HVGICDistributorRegGICDIrouter376:
		return "HVGICDistributorRegGICDIrouter376"
	case HVGICDistributorRegGICDIrouter377:
		return "HVGICDistributorRegGICDIrouter377"
	case HVGICDistributorRegGICDIrouter378:
		return "HVGICDistributorRegGICDIrouter378"
	case HVGICDistributorRegGICDIrouter379:
		return "HVGICDistributorRegGICDIrouter379"
	case HVGICDistributorRegGICDIrouter38:
		return "HVGICDistributorRegGICDIrouter38"
	case HVGICDistributorRegGICDIrouter380:
		return "HVGICDistributorRegGICDIrouter380"
	case HVGICDistributorRegGICDIrouter381:
		return "HVGICDistributorRegGICDIrouter381"
	case HVGICDistributorRegGICDIrouter382:
		return "HVGICDistributorRegGICDIrouter382"
	case HVGICDistributorRegGICDIrouter383:
		return "HVGICDistributorRegGICDIrouter383"
	case HVGICDistributorRegGICDIrouter384:
		return "HVGICDistributorRegGICDIrouter384"
	case HVGICDistributorRegGICDIrouter385:
		return "HVGICDistributorRegGICDIrouter385"
	case HVGICDistributorRegGICDIrouter386:
		return "HVGICDistributorRegGICDIrouter386"
	case HVGICDistributorRegGICDIrouter387:
		return "HVGICDistributorRegGICDIrouter387"
	case HVGICDistributorRegGICDIrouter388:
		return "HVGICDistributorRegGICDIrouter388"
	case HVGICDistributorRegGICDIrouter389:
		return "HVGICDistributorRegGICDIrouter389"
	case HVGICDistributorRegGICDIrouter39:
		return "HVGICDistributorRegGICDIrouter39"
	case HVGICDistributorRegGICDIrouter390:
		return "HVGICDistributorRegGICDIrouter390"
	case HVGICDistributorRegGICDIrouter391:
		return "HVGICDistributorRegGICDIrouter391"
	case HVGICDistributorRegGICDIrouter392:
		return "HVGICDistributorRegGICDIrouter392"
	case HVGICDistributorRegGICDIrouter393:
		return "HVGICDistributorRegGICDIrouter393"
	case HVGICDistributorRegGICDIrouter394:
		return "HVGICDistributorRegGICDIrouter394"
	case HVGICDistributorRegGICDIrouter395:
		return "HVGICDistributorRegGICDIrouter395"
	case HVGICDistributorRegGICDIrouter396:
		return "HVGICDistributorRegGICDIrouter396"
	case HVGICDistributorRegGICDIrouter397:
		return "HVGICDistributorRegGICDIrouter397"
	case HVGICDistributorRegGICDIrouter398:
		return "HVGICDistributorRegGICDIrouter398"
	case HVGICDistributorRegGICDIrouter399:
		return "HVGICDistributorRegGICDIrouter399"
	case HVGICDistributorRegGICDIrouter40:
		return "HVGICDistributorRegGICDIrouter40"
	case HVGICDistributorRegGICDIrouter400:
		return "HVGICDistributorRegGICDIrouter400"
	case HVGICDistributorRegGICDIrouter401:
		return "HVGICDistributorRegGICDIrouter401"
	case HVGICDistributorRegGICDIrouter402:
		return "HVGICDistributorRegGICDIrouter402"
	case HVGICDistributorRegGICDIrouter403:
		return "HVGICDistributorRegGICDIrouter403"
	case HVGICDistributorRegGICDIrouter404:
		return "HVGICDistributorRegGICDIrouter404"
	case HVGICDistributorRegGICDIrouter405:
		return "HVGICDistributorRegGICDIrouter405"
	case HVGICDistributorRegGICDIrouter406:
		return "HVGICDistributorRegGICDIrouter406"
	case HVGICDistributorRegGICDIrouter407:
		return "HVGICDistributorRegGICDIrouter407"
	case HVGICDistributorRegGICDIrouter408:
		return "HVGICDistributorRegGICDIrouter408"
	case HVGICDistributorRegGICDIrouter409:
		return "HVGICDistributorRegGICDIrouter409"
	case HVGICDistributorRegGICDIrouter41:
		return "HVGICDistributorRegGICDIrouter41"
	case HVGICDistributorRegGICDIrouter410:
		return "HVGICDistributorRegGICDIrouter410"
	case HVGICDistributorRegGICDIrouter411:
		return "HVGICDistributorRegGICDIrouter411"
	case HVGICDistributorRegGICDIrouter412:
		return "HVGICDistributorRegGICDIrouter412"
	case HVGICDistributorRegGICDIrouter413:
		return "HVGICDistributorRegGICDIrouter413"
	case HVGICDistributorRegGICDIrouter414:
		return "HVGICDistributorRegGICDIrouter414"
	case HVGICDistributorRegGICDIrouter415:
		return "HVGICDistributorRegGICDIrouter415"
	case HVGICDistributorRegGICDIrouter416:
		return "HVGICDistributorRegGICDIrouter416"
	case HVGICDistributorRegGICDIrouter417:
		return "HVGICDistributorRegGICDIrouter417"
	case HVGICDistributorRegGICDIrouter418:
		return "HVGICDistributorRegGICDIrouter418"
	case HVGICDistributorRegGICDIrouter419:
		return "HVGICDistributorRegGICDIrouter419"
	case HVGICDistributorRegGICDIrouter42:
		return "HVGICDistributorRegGICDIrouter42"
	case HVGICDistributorRegGICDIrouter420:
		return "HVGICDistributorRegGICDIrouter420"
	case HVGICDistributorRegGICDIrouter421:
		return "HVGICDistributorRegGICDIrouter421"
	case HVGICDistributorRegGICDIrouter422:
		return "HVGICDistributorRegGICDIrouter422"
	case HVGICDistributorRegGICDIrouter423:
		return "HVGICDistributorRegGICDIrouter423"
	case HVGICDistributorRegGICDIrouter424:
		return "HVGICDistributorRegGICDIrouter424"
	case HVGICDistributorRegGICDIrouter425:
		return "HVGICDistributorRegGICDIrouter425"
	case HVGICDistributorRegGICDIrouter426:
		return "HVGICDistributorRegGICDIrouter426"
	case HVGICDistributorRegGICDIrouter427:
		return "HVGICDistributorRegGICDIrouter427"
	case HVGICDistributorRegGICDIrouter428:
		return "HVGICDistributorRegGICDIrouter428"
	case HVGICDistributorRegGICDIrouter429:
		return "HVGICDistributorRegGICDIrouter429"
	case HVGICDistributorRegGICDIrouter43:
		return "HVGICDistributorRegGICDIrouter43"
	case HVGICDistributorRegGICDIrouter430:
		return "HVGICDistributorRegGICDIrouter430"
	case HVGICDistributorRegGICDIrouter431:
		return "HVGICDistributorRegGICDIrouter431"
	case HVGICDistributorRegGICDIrouter432:
		return "HVGICDistributorRegGICDIrouter432"
	case HVGICDistributorRegGICDIrouter433:
		return "HVGICDistributorRegGICDIrouter433"
	case HVGICDistributorRegGICDIrouter434:
		return "HVGICDistributorRegGICDIrouter434"
	case HVGICDistributorRegGICDIrouter435:
		return "HVGICDistributorRegGICDIrouter435"
	case HVGICDistributorRegGICDIrouter436:
		return "HVGICDistributorRegGICDIrouter436"
	case HVGICDistributorRegGICDIrouter437:
		return "HVGICDistributorRegGICDIrouter437"
	case HVGICDistributorRegGICDIrouter438:
		return "HVGICDistributorRegGICDIrouter438"
	case HVGICDistributorRegGICDIrouter439:
		return "HVGICDistributorRegGICDIrouter439"
	case HVGICDistributorRegGICDIrouter44:
		return "HVGICDistributorRegGICDIrouter44"
	case HVGICDistributorRegGICDIrouter440:
		return "HVGICDistributorRegGICDIrouter440"
	case HVGICDistributorRegGICDIrouter441:
		return "HVGICDistributorRegGICDIrouter441"
	case HVGICDistributorRegGICDIrouter442:
		return "HVGICDistributorRegGICDIrouter442"
	case HVGICDistributorRegGICDIrouter443:
		return "HVGICDistributorRegGICDIrouter443"
	case HVGICDistributorRegGICDIrouter444:
		return "HVGICDistributorRegGICDIrouter444"
	case HVGICDistributorRegGICDIrouter445:
		return "HVGICDistributorRegGICDIrouter445"
	case HVGICDistributorRegGICDIrouter446:
		return "HVGICDistributorRegGICDIrouter446"
	case HVGICDistributorRegGICDIrouter447:
		return "HVGICDistributorRegGICDIrouter447"
	case HVGICDistributorRegGICDIrouter448:
		return "HVGICDistributorRegGICDIrouter448"
	case HVGICDistributorRegGICDIrouter449:
		return "HVGICDistributorRegGICDIrouter449"
	case HVGICDistributorRegGICDIrouter45:
		return "HVGICDistributorRegGICDIrouter45"
	case HVGICDistributorRegGICDIrouter450:
		return "HVGICDistributorRegGICDIrouter450"
	case HVGICDistributorRegGICDIrouter451:
		return "HVGICDistributorRegGICDIrouter451"
	case HVGICDistributorRegGICDIrouter452:
		return "HVGICDistributorRegGICDIrouter452"
	case HVGICDistributorRegGICDIrouter453:
		return "HVGICDistributorRegGICDIrouter453"
	case HVGICDistributorRegGICDIrouter454:
		return "HVGICDistributorRegGICDIrouter454"
	case HVGICDistributorRegGICDIrouter455:
		return "HVGICDistributorRegGICDIrouter455"
	case HVGICDistributorRegGICDIrouter456:
		return "HVGICDistributorRegGICDIrouter456"
	case HVGICDistributorRegGICDIrouter457:
		return "HVGICDistributorRegGICDIrouter457"
	case HVGICDistributorRegGICDIrouter458:
		return "HVGICDistributorRegGICDIrouter458"
	case HVGICDistributorRegGICDIrouter459:
		return "HVGICDistributorRegGICDIrouter459"
	case HVGICDistributorRegGICDIrouter46:
		return "HVGICDistributorRegGICDIrouter46"
	case HVGICDistributorRegGICDIrouter460:
		return "HVGICDistributorRegGICDIrouter460"
	case HVGICDistributorRegGICDIrouter461:
		return "HVGICDistributorRegGICDIrouter461"
	case HVGICDistributorRegGICDIrouter462:
		return "HVGICDistributorRegGICDIrouter462"
	case HVGICDistributorRegGICDIrouter463:
		return "HVGICDistributorRegGICDIrouter463"
	case HVGICDistributorRegGICDIrouter464:
		return "HVGICDistributorRegGICDIrouter464"
	case HVGICDistributorRegGICDIrouter465:
		return "HVGICDistributorRegGICDIrouter465"
	case HVGICDistributorRegGICDIrouter466:
		return "HVGICDistributorRegGICDIrouter466"
	case HVGICDistributorRegGICDIrouter467:
		return "HVGICDistributorRegGICDIrouter467"
	case HVGICDistributorRegGICDIrouter468:
		return "HVGICDistributorRegGICDIrouter468"
	case HVGICDistributorRegGICDIrouter469:
		return "HVGICDistributorRegGICDIrouter469"
	case HVGICDistributorRegGICDIrouter47:
		return "HVGICDistributorRegGICDIrouter47"
	case HVGICDistributorRegGICDIrouter470:
		return "HVGICDistributorRegGICDIrouter470"
	case HVGICDistributorRegGICDIrouter471:
		return "HVGICDistributorRegGICDIrouter471"
	case HVGICDistributorRegGICDIrouter472:
		return "HVGICDistributorRegGICDIrouter472"
	case HVGICDistributorRegGICDIrouter473:
		return "HVGICDistributorRegGICDIrouter473"
	case HVGICDistributorRegGICDIrouter474:
		return "HVGICDistributorRegGICDIrouter474"
	case HVGICDistributorRegGICDIrouter475:
		return "HVGICDistributorRegGICDIrouter475"
	case HVGICDistributorRegGICDIrouter476:
		return "HVGICDistributorRegGICDIrouter476"
	case HVGICDistributorRegGICDIrouter477:
		return "HVGICDistributorRegGICDIrouter477"
	case HVGICDistributorRegGICDIrouter478:
		return "HVGICDistributorRegGICDIrouter478"
	case HVGICDistributorRegGICDIrouter479:
		return "HVGICDistributorRegGICDIrouter479"
	case HVGICDistributorRegGICDIrouter48:
		return "HVGICDistributorRegGICDIrouter48"
	case HVGICDistributorRegGICDIrouter480:
		return "HVGICDistributorRegGICDIrouter480"
	case HVGICDistributorRegGICDIrouter481:
		return "HVGICDistributorRegGICDIrouter481"
	case HVGICDistributorRegGICDIrouter482:
		return "HVGICDistributorRegGICDIrouter482"
	case HVGICDistributorRegGICDIrouter483:
		return "HVGICDistributorRegGICDIrouter483"
	case HVGICDistributorRegGICDIrouter484:
		return "HVGICDistributorRegGICDIrouter484"
	case HVGICDistributorRegGICDIrouter485:
		return "HVGICDistributorRegGICDIrouter485"
	case HVGICDistributorRegGICDIrouter486:
		return "HVGICDistributorRegGICDIrouter486"
	case HVGICDistributorRegGICDIrouter487:
		return "HVGICDistributorRegGICDIrouter487"
	case HVGICDistributorRegGICDIrouter488:
		return "HVGICDistributorRegGICDIrouter488"
	case HVGICDistributorRegGICDIrouter489:
		return "HVGICDistributorRegGICDIrouter489"
	case HVGICDistributorRegGICDIrouter49:
		return "HVGICDistributorRegGICDIrouter49"
	case HVGICDistributorRegGICDIrouter490:
		return "HVGICDistributorRegGICDIrouter490"
	case HVGICDistributorRegGICDIrouter491:
		return "HVGICDistributorRegGICDIrouter491"
	case HVGICDistributorRegGICDIrouter492:
		return "HVGICDistributorRegGICDIrouter492"
	case HVGICDistributorRegGICDIrouter493:
		return "HVGICDistributorRegGICDIrouter493"
	case HVGICDistributorRegGICDIrouter494:
		return "HVGICDistributorRegGICDIrouter494"
	case HVGICDistributorRegGICDIrouter495:
		return "HVGICDistributorRegGICDIrouter495"
	case HVGICDistributorRegGICDIrouter496:
		return "HVGICDistributorRegGICDIrouter496"
	case HVGICDistributorRegGICDIrouter497:
		return "HVGICDistributorRegGICDIrouter497"
	case HVGICDistributorRegGICDIrouter498:
		return "HVGICDistributorRegGICDIrouter498"
	case HVGICDistributorRegGICDIrouter499:
		return "HVGICDistributorRegGICDIrouter499"
	case HVGICDistributorRegGICDIrouter50:
		return "HVGICDistributorRegGICDIrouter50"
	case HVGICDistributorRegGICDIrouter500:
		return "HVGICDistributorRegGICDIrouter500"
	case HVGICDistributorRegGICDIrouter501:
		return "HVGICDistributorRegGICDIrouter501"
	case HVGICDistributorRegGICDIrouter502:
		return "HVGICDistributorRegGICDIrouter502"
	case HVGICDistributorRegGICDIrouter503:
		return "HVGICDistributorRegGICDIrouter503"
	case HVGICDistributorRegGICDIrouter504:
		return "HVGICDistributorRegGICDIrouter504"
	case HVGICDistributorRegGICDIrouter505:
		return "HVGICDistributorRegGICDIrouter505"
	case HVGICDistributorRegGICDIrouter506:
		return "HVGICDistributorRegGICDIrouter506"
	case HVGICDistributorRegGICDIrouter507:
		return "HVGICDistributorRegGICDIrouter507"
	case HVGICDistributorRegGICDIrouter508:
		return "HVGICDistributorRegGICDIrouter508"
	case HVGICDistributorRegGICDIrouter509:
		return "HVGICDistributorRegGICDIrouter509"
	case HVGICDistributorRegGICDIrouter51:
		return "HVGICDistributorRegGICDIrouter51"
	case HVGICDistributorRegGICDIrouter510:
		return "HVGICDistributorRegGICDIrouter510"
	case HVGICDistributorRegGICDIrouter511:
		return "HVGICDistributorRegGICDIrouter511"
	case HVGICDistributorRegGICDIrouter512:
		return "HVGICDistributorRegGICDIrouter512"
	case HVGICDistributorRegGICDIrouter513:
		return "HVGICDistributorRegGICDIrouter513"
	case HVGICDistributorRegGICDIrouter514:
		return "HVGICDistributorRegGICDIrouter514"
	case HVGICDistributorRegGICDIrouter515:
		return "HVGICDistributorRegGICDIrouter515"
	case HVGICDistributorRegGICDIrouter516:
		return "HVGICDistributorRegGICDIrouter516"
	case HVGICDistributorRegGICDIrouter517:
		return "HVGICDistributorRegGICDIrouter517"
	case HVGICDistributorRegGICDIrouter518:
		return "HVGICDistributorRegGICDIrouter518"
	case HVGICDistributorRegGICDIrouter519:
		return "HVGICDistributorRegGICDIrouter519"
	case HVGICDistributorRegGICDIrouter52:
		return "HVGICDistributorRegGICDIrouter52"
	case HVGICDistributorRegGICDIrouter520:
		return "HVGICDistributorRegGICDIrouter520"
	case HVGICDistributorRegGICDIrouter521:
		return "HVGICDistributorRegGICDIrouter521"
	case HVGICDistributorRegGICDIrouter522:
		return "HVGICDistributorRegGICDIrouter522"
	case HVGICDistributorRegGICDIrouter523:
		return "HVGICDistributorRegGICDIrouter523"
	case HVGICDistributorRegGICDIrouter524:
		return "HVGICDistributorRegGICDIrouter524"
	case HVGICDistributorRegGICDIrouter525:
		return "HVGICDistributorRegGICDIrouter525"
	case HVGICDistributorRegGICDIrouter526:
		return "HVGICDistributorRegGICDIrouter526"
	case HVGICDistributorRegGICDIrouter527:
		return "HVGICDistributorRegGICDIrouter527"
	case HVGICDistributorRegGICDIrouter528:
		return "HVGICDistributorRegGICDIrouter528"
	case HVGICDistributorRegGICDIrouter529:
		return "HVGICDistributorRegGICDIrouter529"
	case HVGICDistributorRegGICDIrouter53:
		return "HVGICDistributorRegGICDIrouter53"
	case HVGICDistributorRegGICDIrouter530:
		return "HVGICDistributorRegGICDIrouter530"
	case HVGICDistributorRegGICDIrouter531:
		return "HVGICDistributorRegGICDIrouter531"
	case HVGICDistributorRegGICDIrouter532:
		return "HVGICDistributorRegGICDIrouter532"
	case HVGICDistributorRegGICDIrouter533:
		return "HVGICDistributorRegGICDIrouter533"
	case HVGICDistributorRegGICDIrouter534:
		return "HVGICDistributorRegGICDIrouter534"
	case HVGICDistributorRegGICDIrouter535:
		return "HVGICDistributorRegGICDIrouter535"
	case HVGICDistributorRegGICDIrouter536:
		return "HVGICDistributorRegGICDIrouter536"
	case HVGICDistributorRegGICDIrouter537:
		return "HVGICDistributorRegGICDIrouter537"
	case HVGICDistributorRegGICDIrouter538:
		return "HVGICDistributorRegGICDIrouter538"
	case HVGICDistributorRegGICDIrouter539:
		return "HVGICDistributorRegGICDIrouter539"
	case HVGICDistributorRegGICDIrouter54:
		return "HVGICDistributorRegGICDIrouter54"
	case HVGICDistributorRegGICDIrouter540:
		return "HVGICDistributorRegGICDIrouter540"
	case HVGICDistributorRegGICDIrouter541:
		return "HVGICDistributorRegGICDIrouter541"
	case HVGICDistributorRegGICDIrouter542:
		return "HVGICDistributorRegGICDIrouter542"
	case HVGICDistributorRegGICDIrouter543:
		return "HVGICDistributorRegGICDIrouter543"
	case HVGICDistributorRegGICDIrouter544:
		return "HVGICDistributorRegGICDIrouter544"
	case HVGICDistributorRegGICDIrouter545:
		return "HVGICDistributorRegGICDIrouter545"
	case HVGICDistributorRegGICDIrouter546:
		return "HVGICDistributorRegGICDIrouter546"
	case HVGICDistributorRegGICDIrouter547:
		return "HVGICDistributorRegGICDIrouter547"
	case HVGICDistributorRegGICDIrouter548:
		return "HVGICDistributorRegGICDIrouter548"
	case HVGICDistributorRegGICDIrouter549:
		return "HVGICDistributorRegGICDIrouter549"
	case HVGICDistributorRegGICDIrouter55:
		return "HVGICDistributorRegGICDIrouter55"
	case HVGICDistributorRegGICDIrouter550:
		return "HVGICDistributorRegGICDIrouter550"
	case HVGICDistributorRegGICDIrouter551:
		return "HVGICDistributorRegGICDIrouter551"
	case HVGICDistributorRegGICDIrouter552:
		return "HVGICDistributorRegGICDIrouter552"
	case HVGICDistributorRegGICDIrouter553:
		return "HVGICDistributorRegGICDIrouter553"
	case HVGICDistributorRegGICDIrouter554:
		return "HVGICDistributorRegGICDIrouter554"
	case HVGICDistributorRegGICDIrouter555:
		return "HVGICDistributorRegGICDIrouter555"
	case HVGICDistributorRegGICDIrouter556:
		return "HVGICDistributorRegGICDIrouter556"
	case HVGICDistributorRegGICDIrouter557:
		return "HVGICDistributorRegGICDIrouter557"
	case HVGICDistributorRegGICDIrouter558:
		return "HVGICDistributorRegGICDIrouter558"
	case HVGICDistributorRegGICDIrouter559:
		return "HVGICDistributorRegGICDIrouter559"
	case HVGICDistributorRegGICDIrouter56:
		return "HVGICDistributorRegGICDIrouter56"
	case HVGICDistributorRegGICDIrouter560:
		return "HVGICDistributorRegGICDIrouter560"
	case HVGICDistributorRegGICDIrouter561:
		return "HVGICDistributorRegGICDIrouter561"
	case HVGICDistributorRegGICDIrouter562:
		return "HVGICDistributorRegGICDIrouter562"
	case HVGICDistributorRegGICDIrouter563:
		return "HVGICDistributorRegGICDIrouter563"
	case HVGICDistributorRegGICDIrouter564:
		return "HVGICDistributorRegGICDIrouter564"
	case HVGICDistributorRegGICDIrouter565:
		return "HVGICDistributorRegGICDIrouter565"
	case HVGICDistributorRegGICDIrouter566:
		return "HVGICDistributorRegGICDIrouter566"
	case HVGICDistributorRegGICDIrouter567:
		return "HVGICDistributorRegGICDIrouter567"
	case HVGICDistributorRegGICDIrouter568:
		return "HVGICDistributorRegGICDIrouter568"
	case HVGICDistributorRegGICDIrouter569:
		return "HVGICDistributorRegGICDIrouter569"
	case HVGICDistributorRegGICDIrouter57:
		return "HVGICDistributorRegGICDIrouter57"
	case HVGICDistributorRegGICDIrouter570:
		return "HVGICDistributorRegGICDIrouter570"
	case HVGICDistributorRegGICDIrouter571:
		return "HVGICDistributorRegGICDIrouter571"
	case HVGICDistributorRegGICDIrouter572:
		return "HVGICDistributorRegGICDIrouter572"
	case HVGICDistributorRegGICDIrouter573:
		return "HVGICDistributorRegGICDIrouter573"
	case HVGICDistributorRegGICDIrouter574:
		return "HVGICDistributorRegGICDIrouter574"
	case HVGICDistributorRegGICDIrouter575:
		return "HVGICDistributorRegGICDIrouter575"
	case HVGICDistributorRegGICDIrouter576:
		return "HVGICDistributorRegGICDIrouter576"
	case HVGICDistributorRegGICDIrouter577:
		return "HVGICDistributorRegGICDIrouter577"
	case HVGICDistributorRegGICDIrouter578:
		return "HVGICDistributorRegGICDIrouter578"
	case HVGICDistributorRegGICDIrouter579:
		return "HVGICDistributorRegGICDIrouter579"
	case HVGICDistributorRegGICDIrouter58:
		return "HVGICDistributorRegGICDIrouter58"
	case HVGICDistributorRegGICDIrouter580:
		return "HVGICDistributorRegGICDIrouter580"
	case HVGICDistributorRegGICDIrouter581:
		return "HVGICDistributorRegGICDIrouter581"
	case HVGICDistributorRegGICDIrouter582:
		return "HVGICDistributorRegGICDIrouter582"
	case HVGICDistributorRegGICDIrouter583:
		return "HVGICDistributorRegGICDIrouter583"
	case HVGICDistributorRegGICDIrouter584:
		return "HVGICDistributorRegGICDIrouter584"
	case HVGICDistributorRegGICDIrouter585:
		return "HVGICDistributorRegGICDIrouter585"
	case HVGICDistributorRegGICDIrouter586:
		return "HVGICDistributorRegGICDIrouter586"
	case HVGICDistributorRegGICDIrouter587:
		return "HVGICDistributorRegGICDIrouter587"
	case HVGICDistributorRegGICDIrouter588:
		return "HVGICDistributorRegGICDIrouter588"
	case HVGICDistributorRegGICDIrouter589:
		return "HVGICDistributorRegGICDIrouter589"
	case HVGICDistributorRegGICDIrouter59:
		return "HVGICDistributorRegGICDIrouter59"
	case HVGICDistributorRegGICDIrouter590:
		return "HVGICDistributorRegGICDIrouter590"
	case HVGICDistributorRegGICDIrouter591:
		return "HVGICDistributorRegGICDIrouter591"
	case HVGICDistributorRegGICDIrouter592:
		return "HVGICDistributorRegGICDIrouter592"
	case HVGICDistributorRegGICDIrouter593:
		return "HVGICDistributorRegGICDIrouter593"
	case HVGICDistributorRegGICDIrouter594:
		return "HVGICDistributorRegGICDIrouter594"
	case HVGICDistributorRegGICDIrouter595:
		return "HVGICDistributorRegGICDIrouter595"
	case HVGICDistributorRegGICDIrouter596:
		return "HVGICDistributorRegGICDIrouter596"
	case HVGICDistributorRegGICDIrouter597:
		return "HVGICDistributorRegGICDIrouter597"
	case HVGICDistributorRegGICDIrouter598:
		return "HVGICDistributorRegGICDIrouter598"
	case HVGICDistributorRegGICDIrouter599:
		return "HVGICDistributorRegGICDIrouter599"
	case HVGICDistributorRegGICDIrouter60:
		return "HVGICDistributorRegGICDIrouter60"
	case HVGICDistributorRegGICDIrouter600:
		return "HVGICDistributorRegGICDIrouter600"
	case HVGICDistributorRegGICDIrouter601:
		return "HVGICDistributorRegGICDIrouter601"
	case HVGICDistributorRegGICDIrouter602:
		return "HVGICDistributorRegGICDIrouter602"
	case HVGICDistributorRegGICDIrouter603:
		return "HVGICDistributorRegGICDIrouter603"
	case HVGICDistributorRegGICDIrouter604:
		return "HVGICDistributorRegGICDIrouter604"
	case HVGICDistributorRegGICDIrouter605:
		return "HVGICDistributorRegGICDIrouter605"
	case HVGICDistributorRegGICDIrouter606:
		return "HVGICDistributorRegGICDIrouter606"
	case HVGICDistributorRegGICDIrouter607:
		return "HVGICDistributorRegGICDIrouter607"
	case HVGICDistributorRegGICDIrouter608:
		return "HVGICDistributorRegGICDIrouter608"
	case HVGICDistributorRegGICDIrouter609:
		return "HVGICDistributorRegGICDIrouter609"
	case HVGICDistributorRegGICDIrouter61:
		return "HVGICDistributorRegGICDIrouter61"
	case HVGICDistributorRegGICDIrouter610:
		return "HVGICDistributorRegGICDIrouter610"
	case HVGICDistributorRegGICDIrouter611:
		return "HVGICDistributorRegGICDIrouter611"
	case HVGICDistributorRegGICDIrouter612:
		return "HVGICDistributorRegGICDIrouter612"
	case HVGICDistributorRegGICDIrouter613:
		return "HVGICDistributorRegGICDIrouter613"
	case HVGICDistributorRegGICDIrouter614:
		return "HVGICDistributorRegGICDIrouter614"
	case HVGICDistributorRegGICDIrouter615:
		return "HVGICDistributorRegGICDIrouter615"
	case HVGICDistributorRegGICDIrouter616:
		return "HVGICDistributorRegGICDIrouter616"
	case HVGICDistributorRegGICDIrouter617:
		return "HVGICDistributorRegGICDIrouter617"
	case HVGICDistributorRegGICDIrouter618:
		return "HVGICDistributorRegGICDIrouter618"
	case HVGICDistributorRegGICDIrouter619:
		return "HVGICDistributorRegGICDIrouter619"
	case HVGICDistributorRegGICDIrouter62:
		return "HVGICDistributorRegGICDIrouter62"
	case HVGICDistributorRegGICDIrouter620:
		return "HVGICDistributorRegGICDIrouter620"
	case HVGICDistributorRegGICDIrouter621:
		return "HVGICDistributorRegGICDIrouter621"
	case HVGICDistributorRegGICDIrouter622:
		return "HVGICDistributorRegGICDIrouter622"
	case HVGICDistributorRegGICDIrouter623:
		return "HVGICDistributorRegGICDIrouter623"
	case HVGICDistributorRegGICDIrouter624:
		return "HVGICDistributorRegGICDIrouter624"
	case HVGICDistributorRegGICDIrouter625:
		return "HVGICDistributorRegGICDIrouter625"
	case HVGICDistributorRegGICDIrouter626:
		return "HVGICDistributorRegGICDIrouter626"
	case HVGICDistributorRegGICDIrouter627:
		return "HVGICDistributorRegGICDIrouter627"
	case HVGICDistributorRegGICDIrouter628:
		return "HVGICDistributorRegGICDIrouter628"
	case HVGICDistributorRegGICDIrouter629:
		return "HVGICDistributorRegGICDIrouter629"
	case HVGICDistributorRegGICDIrouter63:
		return "HVGICDistributorRegGICDIrouter63"
	case HVGICDistributorRegGICDIrouter630:
		return "HVGICDistributorRegGICDIrouter630"
	case HVGICDistributorRegGICDIrouter631:
		return "HVGICDistributorRegGICDIrouter631"
	case HVGICDistributorRegGICDIrouter632:
		return "HVGICDistributorRegGICDIrouter632"
	case HVGICDistributorRegGICDIrouter633:
		return "HVGICDistributorRegGICDIrouter633"
	case HVGICDistributorRegGICDIrouter634:
		return "HVGICDistributorRegGICDIrouter634"
	case HVGICDistributorRegGICDIrouter635:
		return "HVGICDistributorRegGICDIrouter635"
	case HVGICDistributorRegGICDIrouter636:
		return "HVGICDistributorRegGICDIrouter636"
	case HVGICDistributorRegGICDIrouter637:
		return "HVGICDistributorRegGICDIrouter637"
	case HVGICDistributorRegGICDIrouter638:
		return "HVGICDistributorRegGICDIrouter638"
	case HVGICDistributorRegGICDIrouter639:
		return "HVGICDistributorRegGICDIrouter639"
	case HVGICDistributorRegGICDIrouter64:
		return "HVGICDistributorRegGICDIrouter64"
	case HVGICDistributorRegGICDIrouter640:
		return "HVGICDistributorRegGICDIrouter640"
	case HVGICDistributorRegGICDIrouter641:
		return "HVGICDistributorRegGICDIrouter641"
	case HVGICDistributorRegGICDIrouter642:
		return "HVGICDistributorRegGICDIrouter642"
	case HVGICDistributorRegGICDIrouter643:
		return "HVGICDistributorRegGICDIrouter643"
	case HVGICDistributorRegGICDIrouter644:
		return "HVGICDistributorRegGICDIrouter644"
	case HVGICDistributorRegGICDIrouter645:
		return "HVGICDistributorRegGICDIrouter645"
	case HVGICDistributorRegGICDIrouter646:
		return "HVGICDistributorRegGICDIrouter646"
	case HVGICDistributorRegGICDIrouter647:
		return "HVGICDistributorRegGICDIrouter647"
	case HVGICDistributorRegGICDIrouter648:
		return "HVGICDistributorRegGICDIrouter648"
	case HVGICDistributorRegGICDIrouter649:
		return "HVGICDistributorRegGICDIrouter649"
	case HVGICDistributorRegGICDIrouter65:
		return "HVGICDistributorRegGICDIrouter65"
	case HVGICDistributorRegGICDIrouter650:
		return "HVGICDistributorRegGICDIrouter650"
	case HVGICDistributorRegGICDIrouter651:
		return "HVGICDistributorRegGICDIrouter651"
	case HVGICDistributorRegGICDIrouter652:
		return "HVGICDistributorRegGICDIrouter652"
	case HVGICDistributorRegGICDIrouter653:
		return "HVGICDistributorRegGICDIrouter653"
	case HVGICDistributorRegGICDIrouter654:
		return "HVGICDistributorRegGICDIrouter654"
	case HVGICDistributorRegGICDIrouter655:
		return "HVGICDistributorRegGICDIrouter655"
	case HVGICDistributorRegGICDIrouter656:
		return "HVGICDistributorRegGICDIrouter656"
	case HVGICDistributorRegGICDIrouter657:
		return "HVGICDistributorRegGICDIrouter657"
	case HVGICDistributorRegGICDIrouter658:
		return "HVGICDistributorRegGICDIrouter658"
	case HVGICDistributorRegGICDIrouter659:
		return "HVGICDistributorRegGICDIrouter659"
	case HVGICDistributorRegGICDIrouter66:
		return "HVGICDistributorRegGICDIrouter66"
	case HVGICDistributorRegGICDIrouter660:
		return "HVGICDistributorRegGICDIrouter660"
	case HVGICDistributorRegGICDIrouter661:
		return "HVGICDistributorRegGICDIrouter661"
	case HVGICDistributorRegGICDIrouter662:
		return "HVGICDistributorRegGICDIrouter662"
	case HVGICDistributorRegGICDIrouter663:
		return "HVGICDistributorRegGICDIrouter663"
	case HVGICDistributorRegGICDIrouter664:
		return "HVGICDistributorRegGICDIrouter664"
	case HVGICDistributorRegGICDIrouter665:
		return "HVGICDistributorRegGICDIrouter665"
	case HVGICDistributorRegGICDIrouter666:
		return "HVGICDistributorRegGICDIrouter666"
	case HVGICDistributorRegGICDIrouter667:
		return "HVGICDistributorRegGICDIrouter667"
	case HVGICDistributorRegGICDIrouter668:
		return "HVGICDistributorRegGICDIrouter668"
	case HVGICDistributorRegGICDIrouter669:
		return "HVGICDistributorRegGICDIrouter669"
	case HVGICDistributorRegGICDIrouter67:
		return "HVGICDistributorRegGICDIrouter67"
	case HVGICDistributorRegGICDIrouter670:
		return "HVGICDistributorRegGICDIrouter670"
	case HVGICDistributorRegGICDIrouter671:
		return "HVGICDistributorRegGICDIrouter671"
	case HVGICDistributorRegGICDIrouter672:
		return "HVGICDistributorRegGICDIrouter672"
	case HVGICDistributorRegGICDIrouter673:
		return "HVGICDistributorRegGICDIrouter673"
	case HVGICDistributorRegGICDIrouter674:
		return "HVGICDistributorRegGICDIrouter674"
	case HVGICDistributorRegGICDIrouter675:
		return "HVGICDistributorRegGICDIrouter675"
	case HVGICDistributorRegGICDIrouter676:
		return "HVGICDistributorRegGICDIrouter676"
	case HVGICDistributorRegGICDIrouter677:
		return "HVGICDistributorRegGICDIrouter677"
	case HVGICDistributorRegGICDIrouter678:
		return "HVGICDistributorRegGICDIrouter678"
	case HVGICDistributorRegGICDIrouter679:
		return "HVGICDistributorRegGICDIrouter679"
	case HVGICDistributorRegGICDIrouter68:
		return "HVGICDistributorRegGICDIrouter68"
	case HVGICDistributorRegGICDIrouter680:
		return "HVGICDistributorRegGICDIrouter680"
	case HVGICDistributorRegGICDIrouter681:
		return "HVGICDistributorRegGICDIrouter681"
	case HVGICDistributorRegGICDIrouter682:
		return "HVGICDistributorRegGICDIrouter682"
	case HVGICDistributorRegGICDIrouter683:
		return "HVGICDistributorRegGICDIrouter683"
	case HVGICDistributorRegGICDIrouter684:
		return "HVGICDistributorRegGICDIrouter684"
	case HVGICDistributorRegGICDIrouter685:
		return "HVGICDistributorRegGICDIrouter685"
	case HVGICDistributorRegGICDIrouter686:
		return "HVGICDistributorRegGICDIrouter686"
	case HVGICDistributorRegGICDIrouter687:
		return "HVGICDistributorRegGICDIrouter687"
	case HVGICDistributorRegGICDIrouter688:
		return "HVGICDistributorRegGICDIrouter688"
	case HVGICDistributorRegGICDIrouter689:
		return "HVGICDistributorRegGICDIrouter689"
	case HVGICDistributorRegGICDIrouter69:
		return "HVGICDistributorRegGICDIrouter69"
	case HVGICDistributorRegGICDIrouter690:
		return "HVGICDistributorRegGICDIrouter690"
	case HVGICDistributorRegGICDIrouter691:
		return "HVGICDistributorRegGICDIrouter691"
	case HVGICDistributorRegGICDIrouter692:
		return "HVGICDistributorRegGICDIrouter692"
	case HVGICDistributorRegGICDIrouter693:
		return "HVGICDistributorRegGICDIrouter693"
	case HVGICDistributorRegGICDIrouter694:
		return "HVGICDistributorRegGICDIrouter694"
	case HVGICDistributorRegGICDIrouter695:
		return "HVGICDistributorRegGICDIrouter695"
	case HVGICDistributorRegGICDIrouter696:
		return "HVGICDistributorRegGICDIrouter696"
	case HVGICDistributorRegGICDIrouter697:
		return "HVGICDistributorRegGICDIrouter697"
	case HVGICDistributorRegGICDIrouter698:
		return "HVGICDistributorRegGICDIrouter698"
	case HVGICDistributorRegGICDIrouter699:
		return "HVGICDistributorRegGICDIrouter699"
	case HVGICDistributorRegGICDIrouter70:
		return "HVGICDistributorRegGICDIrouter70"
	case HVGICDistributorRegGICDIrouter700:
		return "HVGICDistributorRegGICDIrouter700"
	case HVGICDistributorRegGICDIrouter701:
		return "HVGICDistributorRegGICDIrouter701"
	case HVGICDistributorRegGICDIrouter702:
		return "HVGICDistributorRegGICDIrouter702"
	case HVGICDistributorRegGICDIrouter703:
		return "HVGICDistributorRegGICDIrouter703"
	case HVGICDistributorRegGICDIrouter704:
		return "HVGICDistributorRegGICDIrouter704"
	case HVGICDistributorRegGICDIrouter705:
		return "HVGICDistributorRegGICDIrouter705"
	case HVGICDistributorRegGICDIrouter706:
		return "HVGICDistributorRegGICDIrouter706"
	case HVGICDistributorRegGICDIrouter707:
		return "HVGICDistributorRegGICDIrouter707"
	case HVGICDistributorRegGICDIrouter708:
		return "HVGICDistributorRegGICDIrouter708"
	case HVGICDistributorRegGICDIrouter709:
		return "HVGICDistributorRegGICDIrouter709"
	case HVGICDistributorRegGICDIrouter71:
		return "HVGICDistributorRegGICDIrouter71"
	case HVGICDistributorRegGICDIrouter710:
		return "HVGICDistributorRegGICDIrouter710"
	case HVGICDistributorRegGICDIrouter711:
		return "HVGICDistributorRegGICDIrouter711"
	case HVGICDistributorRegGICDIrouter712:
		return "HVGICDistributorRegGICDIrouter712"
	case HVGICDistributorRegGICDIrouter713:
		return "HVGICDistributorRegGICDIrouter713"
	case HVGICDistributorRegGICDIrouter714:
		return "HVGICDistributorRegGICDIrouter714"
	case HVGICDistributorRegGICDIrouter715:
		return "HVGICDistributorRegGICDIrouter715"
	case HVGICDistributorRegGICDIrouter716:
		return "HVGICDistributorRegGICDIrouter716"
	case HVGICDistributorRegGICDIrouter717:
		return "HVGICDistributorRegGICDIrouter717"
	case HVGICDistributorRegGICDIrouter718:
		return "HVGICDistributorRegGICDIrouter718"
	case HVGICDistributorRegGICDIrouter719:
		return "HVGICDistributorRegGICDIrouter719"
	case HVGICDistributorRegGICDIrouter72:
		return "HVGICDistributorRegGICDIrouter72"
	case HVGICDistributorRegGICDIrouter720:
		return "HVGICDistributorRegGICDIrouter720"
	case HVGICDistributorRegGICDIrouter721:
		return "HVGICDistributorRegGICDIrouter721"
	case HVGICDistributorRegGICDIrouter722:
		return "HVGICDistributorRegGICDIrouter722"
	case HVGICDistributorRegGICDIrouter723:
		return "HVGICDistributorRegGICDIrouter723"
	case HVGICDistributorRegGICDIrouter724:
		return "HVGICDistributorRegGICDIrouter724"
	case HVGICDistributorRegGICDIrouter725:
		return "HVGICDistributorRegGICDIrouter725"
	case HVGICDistributorRegGICDIrouter726:
		return "HVGICDistributorRegGICDIrouter726"
	case HVGICDistributorRegGICDIrouter727:
		return "HVGICDistributorRegGICDIrouter727"
	case HVGICDistributorRegGICDIrouter728:
		return "HVGICDistributorRegGICDIrouter728"
	case HVGICDistributorRegGICDIrouter729:
		return "HVGICDistributorRegGICDIrouter729"
	case HVGICDistributorRegGICDIrouter73:
		return "HVGICDistributorRegGICDIrouter73"
	case HVGICDistributorRegGICDIrouter730:
		return "HVGICDistributorRegGICDIrouter730"
	case HVGICDistributorRegGICDIrouter731:
		return "HVGICDistributorRegGICDIrouter731"
	case HVGICDistributorRegGICDIrouter732:
		return "HVGICDistributorRegGICDIrouter732"
	case HVGICDistributorRegGICDIrouter733:
		return "HVGICDistributorRegGICDIrouter733"
	case HVGICDistributorRegGICDIrouter734:
		return "HVGICDistributorRegGICDIrouter734"
	case HVGICDistributorRegGICDIrouter735:
		return "HVGICDistributorRegGICDIrouter735"
	case HVGICDistributorRegGICDIrouter736:
		return "HVGICDistributorRegGICDIrouter736"
	case HVGICDistributorRegGICDIrouter737:
		return "HVGICDistributorRegGICDIrouter737"
	case HVGICDistributorRegGICDIrouter738:
		return "HVGICDistributorRegGICDIrouter738"
	case HVGICDistributorRegGICDIrouter739:
		return "HVGICDistributorRegGICDIrouter739"
	case HVGICDistributorRegGICDIrouter74:
		return "HVGICDistributorRegGICDIrouter74"
	case HVGICDistributorRegGICDIrouter740:
		return "HVGICDistributorRegGICDIrouter740"
	case HVGICDistributorRegGICDIrouter741:
		return "HVGICDistributorRegGICDIrouter741"
	case HVGICDistributorRegGICDIrouter742:
		return "HVGICDistributorRegGICDIrouter742"
	case HVGICDistributorRegGICDIrouter743:
		return "HVGICDistributorRegGICDIrouter743"
	case HVGICDistributorRegGICDIrouter744:
		return "HVGICDistributorRegGICDIrouter744"
	case HVGICDistributorRegGICDIrouter745:
		return "HVGICDistributorRegGICDIrouter745"
	case HVGICDistributorRegGICDIrouter746:
		return "HVGICDistributorRegGICDIrouter746"
	case HVGICDistributorRegGICDIrouter747:
		return "HVGICDistributorRegGICDIrouter747"
	case HVGICDistributorRegGICDIrouter748:
		return "HVGICDistributorRegGICDIrouter748"
	case HVGICDistributorRegGICDIrouter749:
		return "HVGICDistributorRegGICDIrouter749"
	case HVGICDistributorRegGICDIrouter75:
		return "HVGICDistributorRegGICDIrouter75"
	case HVGICDistributorRegGICDIrouter750:
		return "HVGICDistributorRegGICDIrouter750"
	case HVGICDistributorRegGICDIrouter751:
		return "HVGICDistributorRegGICDIrouter751"
	case HVGICDistributorRegGICDIrouter752:
		return "HVGICDistributorRegGICDIrouter752"
	case HVGICDistributorRegGICDIrouter753:
		return "HVGICDistributorRegGICDIrouter753"
	case HVGICDistributorRegGICDIrouter754:
		return "HVGICDistributorRegGICDIrouter754"
	case HVGICDistributorRegGICDIrouter755:
		return "HVGICDistributorRegGICDIrouter755"
	case HVGICDistributorRegGICDIrouter756:
		return "HVGICDistributorRegGICDIrouter756"
	case HVGICDistributorRegGICDIrouter757:
		return "HVGICDistributorRegGICDIrouter757"
	case HVGICDistributorRegGICDIrouter758:
		return "HVGICDistributorRegGICDIrouter758"
	case HVGICDistributorRegGICDIrouter759:
		return "HVGICDistributorRegGICDIrouter759"
	case HVGICDistributorRegGICDIrouter76:
		return "HVGICDistributorRegGICDIrouter76"
	case HVGICDistributorRegGICDIrouter760:
		return "HVGICDistributorRegGICDIrouter760"
	case HVGICDistributorRegGICDIrouter761:
		return "HVGICDistributorRegGICDIrouter761"
	case HVGICDistributorRegGICDIrouter762:
		return "HVGICDistributorRegGICDIrouter762"
	case HVGICDistributorRegGICDIrouter763:
		return "HVGICDistributorRegGICDIrouter763"
	case HVGICDistributorRegGICDIrouter764:
		return "HVGICDistributorRegGICDIrouter764"
	case HVGICDistributorRegGICDIrouter765:
		return "HVGICDistributorRegGICDIrouter765"
	case HVGICDistributorRegGICDIrouter766:
		return "HVGICDistributorRegGICDIrouter766"
	case HVGICDistributorRegGICDIrouter767:
		return "HVGICDistributorRegGICDIrouter767"
	case HVGICDistributorRegGICDIrouter768:
		return "HVGICDistributorRegGICDIrouter768"
	case HVGICDistributorRegGICDIrouter769:
		return "HVGICDistributorRegGICDIrouter769"
	case HVGICDistributorRegGICDIrouter77:
		return "HVGICDistributorRegGICDIrouter77"
	case HVGICDistributorRegGICDIrouter770:
		return "HVGICDistributorRegGICDIrouter770"
	case HVGICDistributorRegGICDIrouter771:
		return "HVGICDistributorRegGICDIrouter771"
	case HVGICDistributorRegGICDIrouter772:
		return "HVGICDistributorRegGICDIrouter772"
	case HVGICDistributorRegGICDIrouter773:
		return "HVGICDistributorRegGICDIrouter773"
	case HVGICDistributorRegGICDIrouter774:
		return "HVGICDistributorRegGICDIrouter774"
	case HVGICDistributorRegGICDIrouter775:
		return "HVGICDistributorRegGICDIrouter775"
	case HVGICDistributorRegGICDIrouter776:
		return "HVGICDistributorRegGICDIrouter776"
	case HVGICDistributorRegGICDIrouter777:
		return "HVGICDistributorRegGICDIrouter777"
	case HVGICDistributorRegGICDIrouter778:
		return "HVGICDistributorRegGICDIrouter778"
	case HVGICDistributorRegGICDIrouter779:
		return "HVGICDistributorRegGICDIrouter779"
	case HVGICDistributorRegGICDIrouter78:
		return "HVGICDistributorRegGICDIrouter78"
	case HVGICDistributorRegGICDIrouter780:
		return "HVGICDistributorRegGICDIrouter780"
	case HVGICDistributorRegGICDIrouter781:
		return "HVGICDistributorRegGICDIrouter781"
	case HVGICDistributorRegGICDIrouter782:
		return "HVGICDistributorRegGICDIrouter782"
	case HVGICDistributorRegGICDIrouter783:
		return "HVGICDistributorRegGICDIrouter783"
	case HVGICDistributorRegGICDIrouter784:
		return "HVGICDistributorRegGICDIrouter784"
	case HVGICDistributorRegGICDIrouter785:
		return "HVGICDistributorRegGICDIrouter785"
	case HVGICDistributorRegGICDIrouter786:
		return "HVGICDistributorRegGICDIrouter786"
	case HVGICDistributorRegGICDIrouter787:
		return "HVGICDistributorRegGICDIrouter787"
	case HVGICDistributorRegGICDIrouter788:
		return "HVGICDistributorRegGICDIrouter788"
	case HVGICDistributorRegGICDIrouter789:
		return "HVGICDistributorRegGICDIrouter789"
	case HVGICDistributorRegGICDIrouter79:
		return "HVGICDistributorRegGICDIrouter79"
	case HVGICDistributorRegGICDIrouter790:
		return "HVGICDistributorRegGICDIrouter790"
	case HVGICDistributorRegGICDIrouter791:
		return "HVGICDistributorRegGICDIrouter791"
	case HVGICDistributorRegGICDIrouter792:
		return "HVGICDistributorRegGICDIrouter792"
	case HVGICDistributorRegGICDIrouter793:
		return "HVGICDistributorRegGICDIrouter793"
	case HVGICDistributorRegGICDIrouter794:
		return "HVGICDistributorRegGICDIrouter794"
	case HVGICDistributorRegGICDIrouter795:
		return "HVGICDistributorRegGICDIrouter795"
	case HVGICDistributorRegGICDIrouter796:
		return "HVGICDistributorRegGICDIrouter796"
	case HVGICDistributorRegGICDIrouter797:
		return "HVGICDistributorRegGICDIrouter797"
	case HVGICDistributorRegGICDIrouter798:
		return "HVGICDistributorRegGICDIrouter798"
	case HVGICDistributorRegGICDIrouter799:
		return "HVGICDistributorRegGICDIrouter799"
	case HVGICDistributorRegGICDIrouter80:
		return "HVGICDistributorRegGICDIrouter80"
	case HVGICDistributorRegGICDIrouter800:
		return "HVGICDistributorRegGICDIrouter800"
	case HVGICDistributorRegGICDIrouter801:
		return "HVGICDistributorRegGICDIrouter801"
	case HVGICDistributorRegGICDIrouter802:
		return "HVGICDistributorRegGICDIrouter802"
	case HVGICDistributorRegGICDIrouter803:
		return "HVGICDistributorRegGICDIrouter803"
	case HVGICDistributorRegGICDIrouter804:
		return "HVGICDistributorRegGICDIrouter804"
	case HVGICDistributorRegGICDIrouter805:
		return "HVGICDistributorRegGICDIrouter805"
	case HVGICDistributorRegGICDIrouter806:
		return "HVGICDistributorRegGICDIrouter806"
	case HVGICDistributorRegGICDIrouter807:
		return "HVGICDistributorRegGICDIrouter807"
	case HVGICDistributorRegGICDIrouter808:
		return "HVGICDistributorRegGICDIrouter808"
	case HVGICDistributorRegGICDIrouter809:
		return "HVGICDistributorRegGICDIrouter809"
	case HVGICDistributorRegGICDIrouter81:
		return "HVGICDistributorRegGICDIrouter81"
	case HVGICDistributorRegGICDIrouter810:
		return "HVGICDistributorRegGICDIrouter810"
	case HVGICDistributorRegGICDIrouter811:
		return "HVGICDistributorRegGICDIrouter811"
	case HVGICDistributorRegGICDIrouter812:
		return "HVGICDistributorRegGICDIrouter812"
	case HVGICDistributorRegGICDIrouter813:
		return "HVGICDistributorRegGICDIrouter813"
	case HVGICDistributorRegGICDIrouter814:
		return "HVGICDistributorRegGICDIrouter814"
	case HVGICDistributorRegGICDIrouter815:
		return "HVGICDistributorRegGICDIrouter815"
	case HVGICDistributorRegGICDIrouter816:
		return "HVGICDistributorRegGICDIrouter816"
	case HVGICDistributorRegGICDIrouter817:
		return "HVGICDistributorRegGICDIrouter817"
	case HVGICDistributorRegGICDIrouter818:
		return "HVGICDistributorRegGICDIrouter818"
	case HVGICDistributorRegGICDIrouter819:
		return "HVGICDistributorRegGICDIrouter819"
	case HVGICDistributorRegGICDIrouter82:
		return "HVGICDistributorRegGICDIrouter82"
	case HVGICDistributorRegGICDIrouter820:
		return "HVGICDistributorRegGICDIrouter820"
	case HVGICDistributorRegGICDIrouter821:
		return "HVGICDistributorRegGICDIrouter821"
	case HVGICDistributorRegGICDIrouter822:
		return "HVGICDistributorRegGICDIrouter822"
	case HVGICDistributorRegGICDIrouter823:
		return "HVGICDistributorRegGICDIrouter823"
	case HVGICDistributorRegGICDIrouter824:
		return "HVGICDistributorRegGICDIrouter824"
	case HVGICDistributorRegGICDIrouter825:
		return "HVGICDistributorRegGICDIrouter825"
	case HVGICDistributorRegGICDIrouter826:
		return "HVGICDistributorRegGICDIrouter826"
	case HVGICDistributorRegGICDIrouter827:
		return "HVGICDistributorRegGICDIrouter827"
	case HVGICDistributorRegGICDIrouter828:
		return "HVGICDistributorRegGICDIrouter828"
	case HVGICDistributorRegGICDIrouter829:
		return "HVGICDistributorRegGICDIrouter829"
	case HVGICDistributorRegGICDIrouter83:
		return "HVGICDistributorRegGICDIrouter83"
	case HVGICDistributorRegGICDIrouter830:
		return "HVGICDistributorRegGICDIrouter830"
	case HVGICDistributorRegGICDIrouter831:
		return "HVGICDistributorRegGICDIrouter831"
	case HVGICDistributorRegGICDIrouter832:
		return "HVGICDistributorRegGICDIrouter832"
	case HVGICDistributorRegGICDIrouter833:
		return "HVGICDistributorRegGICDIrouter833"
	case HVGICDistributorRegGICDIrouter834:
		return "HVGICDistributorRegGICDIrouter834"
	case HVGICDistributorRegGICDIrouter835:
		return "HVGICDistributorRegGICDIrouter835"
	case HVGICDistributorRegGICDIrouter836:
		return "HVGICDistributorRegGICDIrouter836"
	case HVGICDistributorRegGICDIrouter837:
		return "HVGICDistributorRegGICDIrouter837"
	case HVGICDistributorRegGICDIrouter838:
		return "HVGICDistributorRegGICDIrouter838"
	case HVGICDistributorRegGICDIrouter839:
		return "HVGICDistributorRegGICDIrouter839"
	case HVGICDistributorRegGICDIrouter84:
		return "HVGICDistributorRegGICDIrouter84"
	case HVGICDistributorRegGICDIrouter840:
		return "HVGICDistributorRegGICDIrouter840"
	case HVGICDistributorRegGICDIrouter841:
		return "HVGICDistributorRegGICDIrouter841"
	case HVGICDistributorRegGICDIrouter842:
		return "HVGICDistributorRegGICDIrouter842"
	case HVGICDistributorRegGICDIrouter843:
		return "HVGICDistributorRegGICDIrouter843"
	case HVGICDistributorRegGICDIrouter844:
		return "HVGICDistributorRegGICDIrouter844"
	case HVGICDistributorRegGICDIrouter845:
		return "HVGICDistributorRegGICDIrouter845"
	case HVGICDistributorRegGICDIrouter846:
		return "HVGICDistributorRegGICDIrouter846"
	case HVGICDistributorRegGICDIrouter847:
		return "HVGICDistributorRegGICDIrouter847"
	case HVGICDistributorRegGICDIrouter848:
		return "HVGICDistributorRegGICDIrouter848"
	case HVGICDistributorRegGICDIrouter849:
		return "HVGICDistributorRegGICDIrouter849"
	case HVGICDistributorRegGICDIrouter85:
		return "HVGICDistributorRegGICDIrouter85"
	case HVGICDistributorRegGICDIrouter850:
		return "HVGICDistributorRegGICDIrouter850"
	case HVGICDistributorRegGICDIrouter851:
		return "HVGICDistributorRegGICDIrouter851"
	case HVGICDistributorRegGICDIrouter852:
		return "HVGICDistributorRegGICDIrouter852"
	case HVGICDistributorRegGICDIrouter853:
		return "HVGICDistributorRegGICDIrouter853"
	case HVGICDistributorRegGICDIrouter854:
		return "HVGICDistributorRegGICDIrouter854"
	case HVGICDistributorRegGICDIrouter855:
		return "HVGICDistributorRegGICDIrouter855"
	case HVGICDistributorRegGICDIrouter856:
		return "HVGICDistributorRegGICDIrouter856"
	case HVGICDistributorRegGICDIrouter857:
		return "HVGICDistributorRegGICDIrouter857"
	case HVGICDistributorRegGICDIrouter858:
		return "HVGICDistributorRegGICDIrouter858"
	case HVGICDistributorRegGICDIrouter859:
		return "HVGICDistributorRegGICDIrouter859"
	case HVGICDistributorRegGICDIrouter86:
		return "HVGICDistributorRegGICDIrouter86"
	case HVGICDistributorRegGICDIrouter860:
		return "HVGICDistributorRegGICDIrouter860"
	case HVGICDistributorRegGICDIrouter861:
		return "HVGICDistributorRegGICDIrouter861"
	case HVGICDistributorRegGICDIrouter862:
		return "HVGICDistributorRegGICDIrouter862"
	case HVGICDistributorRegGICDIrouter863:
		return "HVGICDistributorRegGICDIrouter863"
	case HVGICDistributorRegGICDIrouter864:
		return "HVGICDistributorRegGICDIrouter864"
	case HVGICDistributorRegGICDIrouter865:
		return "HVGICDistributorRegGICDIrouter865"
	case HVGICDistributorRegGICDIrouter866:
		return "HVGICDistributorRegGICDIrouter866"
	case HVGICDistributorRegGICDIrouter867:
		return "HVGICDistributorRegGICDIrouter867"
	case HVGICDistributorRegGICDIrouter868:
		return "HVGICDistributorRegGICDIrouter868"
	case HVGICDistributorRegGICDIrouter869:
		return "HVGICDistributorRegGICDIrouter869"
	case HVGICDistributorRegGICDIrouter87:
		return "HVGICDistributorRegGICDIrouter87"
	case HVGICDistributorRegGICDIrouter870:
		return "HVGICDistributorRegGICDIrouter870"
	case HVGICDistributorRegGICDIrouter871:
		return "HVGICDistributorRegGICDIrouter871"
	case HVGICDistributorRegGICDIrouter872:
		return "HVGICDistributorRegGICDIrouter872"
	case HVGICDistributorRegGICDIrouter873:
		return "HVGICDistributorRegGICDIrouter873"
	case HVGICDistributorRegGICDIrouter874:
		return "HVGICDistributorRegGICDIrouter874"
	case HVGICDistributorRegGICDIrouter875:
		return "HVGICDistributorRegGICDIrouter875"
	case HVGICDistributorRegGICDIrouter876:
		return "HVGICDistributorRegGICDIrouter876"
	case HVGICDistributorRegGICDIrouter877:
		return "HVGICDistributorRegGICDIrouter877"
	case HVGICDistributorRegGICDIrouter878:
		return "HVGICDistributorRegGICDIrouter878"
	case HVGICDistributorRegGICDIrouter879:
		return "HVGICDistributorRegGICDIrouter879"
	case HVGICDistributorRegGICDIrouter88:
		return "HVGICDistributorRegGICDIrouter88"
	case HVGICDistributorRegGICDIrouter880:
		return "HVGICDistributorRegGICDIrouter880"
	case HVGICDistributorRegGICDIrouter881:
		return "HVGICDistributorRegGICDIrouter881"
	case HVGICDistributorRegGICDIrouter882:
		return "HVGICDistributorRegGICDIrouter882"
	case HVGICDistributorRegGICDIrouter883:
		return "HVGICDistributorRegGICDIrouter883"
	case HVGICDistributorRegGICDIrouter884:
		return "HVGICDistributorRegGICDIrouter884"
	case HVGICDistributorRegGICDIrouter885:
		return "HVGICDistributorRegGICDIrouter885"
	case HVGICDistributorRegGICDIrouter886:
		return "HVGICDistributorRegGICDIrouter886"
	case HVGICDistributorRegGICDIrouter887:
		return "HVGICDistributorRegGICDIrouter887"
	case HVGICDistributorRegGICDIrouter888:
		return "HVGICDistributorRegGICDIrouter888"
	case HVGICDistributorRegGICDIrouter889:
		return "HVGICDistributorRegGICDIrouter889"
	case HVGICDistributorRegGICDIrouter89:
		return "HVGICDistributorRegGICDIrouter89"
	case HVGICDistributorRegGICDIrouter890:
		return "HVGICDistributorRegGICDIrouter890"
	case HVGICDistributorRegGICDIrouter891:
		return "HVGICDistributorRegGICDIrouter891"
	case HVGICDistributorRegGICDIrouter892:
		return "HVGICDistributorRegGICDIrouter892"
	case HVGICDistributorRegGICDIrouter893:
		return "HVGICDistributorRegGICDIrouter893"
	case HVGICDistributorRegGICDIrouter894:
		return "HVGICDistributorRegGICDIrouter894"
	case HVGICDistributorRegGICDIrouter895:
		return "HVGICDistributorRegGICDIrouter895"
	case HVGICDistributorRegGICDIrouter896:
		return "HVGICDistributorRegGICDIrouter896"
	case HVGICDistributorRegGICDIrouter897:
		return "HVGICDistributorRegGICDIrouter897"
	case HVGICDistributorRegGICDIrouter898:
		return "HVGICDistributorRegGICDIrouter898"
	case HVGICDistributorRegGICDIrouter899:
		return "HVGICDistributorRegGICDIrouter899"
	case HVGICDistributorRegGICDIrouter90:
		return "HVGICDistributorRegGICDIrouter90"
	case HVGICDistributorRegGICDIrouter900:
		return "HVGICDistributorRegGICDIrouter900"
	case HVGICDistributorRegGICDIrouter901:
		return "HVGICDistributorRegGICDIrouter901"
	case HVGICDistributorRegGICDIrouter902:
		return "HVGICDistributorRegGICDIrouter902"
	case HVGICDistributorRegGICDIrouter903:
		return "HVGICDistributorRegGICDIrouter903"
	case HVGICDistributorRegGICDIrouter904:
		return "HVGICDistributorRegGICDIrouter904"
	case HVGICDistributorRegGICDIrouter905:
		return "HVGICDistributorRegGICDIrouter905"
	case HVGICDistributorRegGICDIrouter906:
		return "HVGICDistributorRegGICDIrouter906"
	case HVGICDistributorRegGICDIrouter907:
		return "HVGICDistributorRegGICDIrouter907"
	case HVGICDistributorRegGICDIrouter908:
		return "HVGICDistributorRegGICDIrouter908"
	case HVGICDistributorRegGICDIrouter909:
		return "HVGICDistributorRegGICDIrouter909"
	case HVGICDistributorRegGICDIrouter91:
		return "HVGICDistributorRegGICDIrouter91"
	case HVGICDistributorRegGICDIrouter910:
		return "HVGICDistributorRegGICDIrouter910"
	case HVGICDistributorRegGICDIrouter911:
		return "HVGICDistributorRegGICDIrouter911"
	case HVGICDistributorRegGICDIrouter912:
		return "HVGICDistributorRegGICDIrouter912"
	case HVGICDistributorRegGICDIrouter913:
		return "HVGICDistributorRegGICDIrouter913"
	case HVGICDistributorRegGICDIrouter914:
		return "HVGICDistributorRegGICDIrouter914"
	case HVGICDistributorRegGICDIrouter915:
		return "HVGICDistributorRegGICDIrouter915"
	case HVGICDistributorRegGICDIrouter916:
		return "HVGICDistributorRegGICDIrouter916"
	case HVGICDistributorRegGICDIrouter917:
		return "HVGICDistributorRegGICDIrouter917"
	case HVGICDistributorRegGICDIrouter918:
		return "HVGICDistributorRegGICDIrouter918"
	case HVGICDistributorRegGICDIrouter919:
		return "HVGICDistributorRegGICDIrouter919"
	case HVGICDistributorRegGICDIrouter92:
		return "HVGICDistributorRegGICDIrouter92"
	case HVGICDistributorRegGICDIrouter920:
		return "HVGICDistributorRegGICDIrouter920"
	case HVGICDistributorRegGICDIrouter921:
		return "HVGICDistributorRegGICDIrouter921"
	case HVGICDistributorRegGICDIrouter922:
		return "HVGICDistributorRegGICDIrouter922"
	case HVGICDistributorRegGICDIrouter923:
		return "HVGICDistributorRegGICDIrouter923"
	case HVGICDistributorRegGICDIrouter924:
		return "HVGICDistributorRegGICDIrouter924"
	case HVGICDistributorRegGICDIrouter925:
		return "HVGICDistributorRegGICDIrouter925"
	case HVGICDistributorRegGICDIrouter926:
		return "HVGICDistributorRegGICDIrouter926"
	case HVGICDistributorRegGICDIrouter927:
		return "HVGICDistributorRegGICDIrouter927"
	case HVGICDistributorRegGICDIrouter928:
		return "HVGICDistributorRegGICDIrouter928"
	case HVGICDistributorRegGICDIrouter929:
		return "HVGICDistributorRegGICDIrouter929"
	case HVGICDistributorRegGICDIrouter93:
		return "HVGICDistributorRegGICDIrouter93"
	case HVGICDistributorRegGICDIrouter930:
		return "HVGICDistributorRegGICDIrouter930"
	case HVGICDistributorRegGICDIrouter931:
		return "HVGICDistributorRegGICDIrouter931"
	case HVGICDistributorRegGICDIrouter932:
		return "HVGICDistributorRegGICDIrouter932"
	case HVGICDistributorRegGICDIrouter933:
		return "HVGICDistributorRegGICDIrouter933"
	case HVGICDistributorRegGICDIrouter934:
		return "HVGICDistributorRegGICDIrouter934"
	case HVGICDistributorRegGICDIrouter935:
		return "HVGICDistributorRegGICDIrouter935"
	case HVGICDistributorRegGICDIrouter936:
		return "HVGICDistributorRegGICDIrouter936"
	case HVGICDistributorRegGICDIrouter937:
		return "HVGICDistributorRegGICDIrouter937"
	case HVGICDistributorRegGICDIrouter938:
		return "HVGICDistributorRegGICDIrouter938"
	case HVGICDistributorRegGICDIrouter939:
		return "HVGICDistributorRegGICDIrouter939"
	case HVGICDistributorRegGICDIrouter94:
		return "HVGICDistributorRegGICDIrouter94"
	case HVGICDistributorRegGICDIrouter940:
		return "HVGICDistributorRegGICDIrouter940"
	case HVGICDistributorRegGICDIrouter941:
		return "HVGICDistributorRegGICDIrouter941"
	case HVGICDistributorRegGICDIrouter942:
		return "HVGICDistributorRegGICDIrouter942"
	case HVGICDistributorRegGICDIrouter943:
		return "HVGICDistributorRegGICDIrouter943"
	case HVGICDistributorRegGICDIrouter944:
		return "HVGICDistributorRegGICDIrouter944"
	case HVGICDistributorRegGICDIrouter945:
		return "HVGICDistributorRegGICDIrouter945"
	case HVGICDistributorRegGICDIrouter946:
		return "HVGICDistributorRegGICDIrouter946"
	case HVGICDistributorRegGICDIrouter947:
		return "HVGICDistributorRegGICDIrouter947"
	case HVGICDistributorRegGICDIrouter948:
		return "HVGICDistributorRegGICDIrouter948"
	case HVGICDistributorRegGICDIrouter949:
		return "HVGICDistributorRegGICDIrouter949"
	case HVGICDistributorRegGICDIrouter95:
		return "HVGICDistributorRegGICDIrouter95"
	case HVGICDistributorRegGICDIrouter950:
		return "HVGICDistributorRegGICDIrouter950"
	case HVGICDistributorRegGICDIrouter951:
		return "HVGICDistributorRegGICDIrouter951"
	case HVGICDistributorRegGICDIrouter952:
		return "HVGICDistributorRegGICDIrouter952"
	case HVGICDistributorRegGICDIrouter953:
		return "HVGICDistributorRegGICDIrouter953"
	case HVGICDistributorRegGICDIrouter954:
		return "HVGICDistributorRegGICDIrouter954"
	case HVGICDistributorRegGICDIrouter955:
		return "HVGICDistributorRegGICDIrouter955"
	case HVGICDistributorRegGICDIrouter956:
		return "HVGICDistributorRegGICDIrouter956"
	case HVGICDistributorRegGICDIrouter957:
		return "HVGICDistributorRegGICDIrouter957"
	case HVGICDistributorRegGICDIrouter958:
		return "HVGICDistributorRegGICDIrouter958"
	case HVGICDistributorRegGICDIrouter959:
		return "HVGICDistributorRegGICDIrouter959"
	case HVGICDistributorRegGICDIrouter96:
		return "HVGICDistributorRegGICDIrouter96"
	case HVGICDistributorRegGICDIrouter960:
		return "HVGICDistributorRegGICDIrouter960"
	case HVGICDistributorRegGICDIrouter961:
		return "HVGICDistributorRegGICDIrouter961"
	case HVGICDistributorRegGICDIrouter962:
		return "HVGICDistributorRegGICDIrouter962"
	case HVGICDistributorRegGICDIrouter963:
		return "HVGICDistributorRegGICDIrouter963"
	case HVGICDistributorRegGICDIrouter964:
		return "HVGICDistributorRegGICDIrouter964"
	case HVGICDistributorRegGICDIrouter965:
		return "HVGICDistributorRegGICDIrouter965"
	case HVGICDistributorRegGICDIrouter966:
		return "HVGICDistributorRegGICDIrouter966"
	case HVGICDistributorRegGICDIrouter967:
		return "HVGICDistributorRegGICDIrouter967"
	case HVGICDistributorRegGICDIrouter968:
		return "HVGICDistributorRegGICDIrouter968"
	case HVGICDistributorRegGICDIrouter969:
		return "HVGICDistributorRegGICDIrouter969"
	case HVGICDistributorRegGICDIrouter97:
		return "HVGICDistributorRegGICDIrouter97"
	case HVGICDistributorRegGICDIrouter970:
		return "HVGICDistributorRegGICDIrouter970"
	case HVGICDistributorRegGICDIrouter971:
		return "HVGICDistributorRegGICDIrouter971"
	case HVGICDistributorRegGICDIrouter972:
		return "HVGICDistributorRegGICDIrouter972"
	case HVGICDistributorRegGICDIrouter973:
		return "HVGICDistributorRegGICDIrouter973"
	case HVGICDistributorRegGICDIrouter974:
		return "HVGICDistributorRegGICDIrouter974"
	case HVGICDistributorRegGICDIrouter975:
		return "HVGICDistributorRegGICDIrouter975"
	case HVGICDistributorRegGICDIrouter976:
		return "HVGICDistributorRegGICDIrouter976"
	case HVGICDistributorRegGICDIrouter977:
		return "HVGICDistributorRegGICDIrouter977"
	case HVGICDistributorRegGICDIrouter978:
		return "HVGICDistributorRegGICDIrouter978"
	case HVGICDistributorRegGICDIrouter979:
		return "HVGICDistributorRegGICDIrouter979"
	case HVGICDistributorRegGICDIrouter98:
		return "HVGICDistributorRegGICDIrouter98"
	case HVGICDistributorRegGICDIrouter980:
		return "HVGICDistributorRegGICDIrouter980"
	case HVGICDistributorRegGICDIrouter981:
		return "HVGICDistributorRegGICDIrouter981"
	case HVGICDistributorRegGICDIrouter982:
		return "HVGICDistributorRegGICDIrouter982"
	case HVGICDistributorRegGICDIrouter983:
		return "HVGICDistributorRegGICDIrouter983"
	case HVGICDistributorRegGICDIrouter984:
		return "HVGICDistributorRegGICDIrouter984"
	case HVGICDistributorRegGICDIrouter985:
		return "HVGICDistributorRegGICDIrouter985"
	case HVGICDistributorRegGICDIrouter986:
		return "HVGICDistributorRegGICDIrouter986"
	case HVGICDistributorRegGICDIrouter987:
		return "HVGICDistributorRegGICDIrouter987"
	case HVGICDistributorRegGICDIrouter988:
		return "HVGICDistributorRegGICDIrouter988"
	case HVGICDistributorRegGICDIrouter989:
		return "HVGICDistributorRegGICDIrouter989"
	case HVGICDistributorRegGICDIrouter99:
		return "HVGICDistributorRegGICDIrouter99"
	case HVGICDistributorRegGICDIrouter990:
		return "HVGICDistributorRegGICDIrouter990"
	case HVGICDistributorRegGICDIrouter991:
		return "HVGICDistributorRegGICDIrouter991"
	case HVGICDistributorRegGICDIrouter992:
		return "HVGICDistributorRegGICDIrouter992"
	case HVGICDistributorRegGICDIrouter993:
		return "HVGICDistributorRegGICDIrouter993"
	case HVGICDistributorRegGICDIrouter994:
		return "HVGICDistributorRegGICDIrouter994"
	case HVGICDistributorRegGICDIrouter995:
		return "HVGICDistributorRegGICDIrouter995"
	case HVGICDistributorRegGICDIrouter996:
		return "HVGICDistributorRegGICDIrouter996"
	case HVGICDistributorRegGICDIrouter997:
		return "HVGICDistributorRegGICDIrouter997"
	case HVGICDistributorRegGICDIrouter998:
		return "HVGICDistributorRegGICDIrouter998"
	case HVGICDistributorRegGICDIrouter999:
		return "HVGICDistributorRegGICDIrouter999"
	case HVGICDistributorRegGICDIsactiver0:
		return "HVGICDistributorRegGICDIsactiver0"
	case HVGICDistributorRegGICDIsactiver1:
		return "HVGICDistributorRegGICDIsactiver1"
	case HVGICDistributorRegGICDIsactiver10:
		return "HVGICDistributorRegGICDIsactiver10"
	case HVGICDistributorRegGICDIsactiver11:
		return "HVGICDistributorRegGICDIsactiver11"
	case HVGICDistributorRegGICDIsactiver12:
		return "HVGICDistributorRegGICDIsactiver12"
	case HVGICDistributorRegGICDIsactiver13:
		return "HVGICDistributorRegGICDIsactiver13"
	case HVGICDistributorRegGICDIsactiver14:
		return "HVGICDistributorRegGICDIsactiver14"
	case HVGICDistributorRegGICDIsactiver15:
		return "HVGICDistributorRegGICDIsactiver15"
	case HVGICDistributorRegGICDIsactiver16:
		return "HVGICDistributorRegGICDIsactiver16"
	case HVGICDistributorRegGICDIsactiver17:
		return "HVGICDistributorRegGICDIsactiver17"
	case HVGICDistributorRegGICDIsactiver18:
		return "HVGICDistributorRegGICDIsactiver18"
	case HVGICDistributorRegGICDIsactiver19:
		return "HVGICDistributorRegGICDIsactiver19"
	case HVGICDistributorRegGICDIsactiver2:
		return "HVGICDistributorRegGICDIsactiver2"
	case HVGICDistributorRegGICDIsactiver20:
		return "HVGICDistributorRegGICDIsactiver20"
	case HVGICDistributorRegGICDIsactiver21:
		return "HVGICDistributorRegGICDIsactiver21"
	case HVGICDistributorRegGICDIsactiver22:
		return "HVGICDistributorRegGICDIsactiver22"
	case HVGICDistributorRegGICDIsactiver23:
		return "HVGICDistributorRegGICDIsactiver23"
	case HVGICDistributorRegGICDIsactiver24:
		return "HVGICDistributorRegGICDIsactiver24"
	case HVGICDistributorRegGICDIsactiver25:
		return "HVGICDistributorRegGICDIsactiver25"
	case HVGICDistributorRegGICDIsactiver26:
		return "HVGICDistributorRegGICDIsactiver26"
	case HVGICDistributorRegGICDIsactiver27:
		return "HVGICDistributorRegGICDIsactiver27"
	case HVGICDistributorRegGICDIsactiver28:
		return "HVGICDistributorRegGICDIsactiver28"
	case HVGICDistributorRegGICDIsactiver29:
		return "HVGICDistributorRegGICDIsactiver29"
	case HVGICDistributorRegGICDIsactiver3:
		return "HVGICDistributorRegGICDIsactiver3"
	case HVGICDistributorRegGICDIsactiver30:
		return "HVGICDistributorRegGICDIsactiver30"
	case HVGICDistributorRegGICDIsactiver31:
		return "HVGICDistributorRegGICDIsactiver31"
	case HVGICDistributorRegGICDIsactiver4:
		return "HVGICDistributorRegGICDIsactiver4"
	case HVGICDistributorRegGICDIsactiver5:
		return "HVGICDistributorRegGICDIsactiver5"
	case HVGICDistributorRegGICDIsactiver6:
		return "HVGICDistributorRegGICDIsactiver6"
	case HVGICDistributorRegGICDIsactiver7:
		return "HVGICDistributorRegGICDIsactiver7"
	case HVGICDistributorRegGICDIsactiver8:
		return "HVGICDistributorRegGICDIsactiver8"
	case HVGICDistributorRegGICDIsactiver9:
		return "HVGICDistributorRegGICDIsactiver9"
	case HVGICDistributorRegGICDIsenabler0:
		return "HVGICDistributorRegGICDIsenabler0"
	case HVGICDistributorRegGICDIsenabler1:
		return "HVGICDistributorRegGICDIsenabler1"
	case HVGICDistributorRegGICDIsenabler10:
		return "HVGICDistributorRegGICDIsenabler10"
	case HVGICDistributorRegGICDIsenabler11:
		return "HVGICDistributorRegGICDIsenabler11"
	case HVGICDistributorRegGICDIsenabler12:
		return "HVGICDistributorRegGICDIsenabler12"
	case HVGICDistributorRegGICDIsenabler13:
		return "HVGICDistributorRegGICDIsenabler13"
	case HVGICDistributorRegGICDIsenabler14:
		return "HVGICDistributorRegGICDIsenabler14"
	case HVGICDistributorRegGICDIsenabler15:
		return "HVGICDistributorRegGICDIsenabler15"
	case HVGICDistributorRegGICDIsenabler16:
		return "HVGICDistributorRegGICDIsenabler16"
	case HVGICDistributorRegGICDIsenabler17:
		return "HVGICDistributorRegGICDIsenabler17"
	case HVGICDistributorRegGICDIsenabler18:
		return "HVGICDistributorRegGICDIsenabler18"
	case HVGICDistributorRegGICDIsenabler19:
		return "HVGICDistributorRegGICDIsenabler19"
	case HVGICDistributorRegGICDIsenabler2:
		return "HVGICDistributorRegGICDIsenabler2"
	case HVGICDistributorRegGICDIsenabler20:
		return "HVGICDistributorRegGICDIsenabler20"
	case HVGICDistributorRegGICDIsenabler21:
		return "HVGICDistributorRegGICDIsenabler21"
	case HVGICDistributorRegGICDIsenabler22:
		return "HVGICDistributorRegGICDIsenabler22"
	case HVGICDistributorRegGICDIsenabler23:
		return "HVGICDistributorRegGICDIsenabler23"
	case HVGICDistributorRegGICDIsenabler24:
		return "HVGICDistributorRegGICDIsenabler24"
	case HVGICDistributorRegGICDIsenabler25:
		return "HVGICDistributorRegGICDIsenabler25"
	case HVGICDistributorRegGICDIsenabler26:
		return "HVGICDistributorRegGICDIsenabler26"
	case HVGICDistributorRegGICDIsenabler27:
		return "HVGICDistributorRegGICDIsenabler27"
	case HVGICDistributorRegGICDIsenabler28:
		return "HVGICDistributorRegGICDIsenabler28"
	case HVGICDistributorRegGICDIsenabler29:
		return "HVGICDistributorRegGICDIsenabler29"
	case HVGICDistributorRegGICDIsenabler3:
		return "HVGICDistributorRegGICDIsenabler3"
	case HVGICDistributorRegGICDIsenabler30:
		return "HVGICDistributorRegGICDIsenabler30"
	case HVGICDistributorRegGICDIsenabler31:
		return "HVGICDistributorRegGICDIsenabler31"
	case HVGICDistributorRegGICDIsenabler4:
		return "HVGICDistributorRegGICDIsenabler4"
	case HVGICDistributorRegGICDIsenabler5:
		return "HVGICDistributorRegGICDIsenabler5"
	case HVGICDistributorRegGICDIsenabler6:
		return "HVGICDistributorRegGICDIsenabler6"
	case HVGICDistributorRegGICDIsenabler7:
		return "HVGICDistributorRegGICDIsenabler7"
	case HVGICDistributorRegGICDIsenabler8:
		return "HVGICDistributorRegGICDIsenabler8"
	case HVGICDistributorRegGICDIsenabler9:
		return "HVGICDistributorRegGICDIsenabler9"
	case HVGICDistributorRegGICDIspendr0:
		return "HVGICDistributorRegGICDIspendr0"
	case HVGICDistributorRegGICDIspendr1:
		return "HVGICDistributorRegGICDIspendr1"
	case HVGICDistributorRegGICDIspendr10:
		return "HVGICDistributorRegGICDIspendr10"
	case HVGICDistributorRegGICDIspendr11:
		return "HVGICDistributorRegGICDIspendr11"
	case HVGICDistributorRegGICDIspendr12:
		return "HVGICDistributorRegGICDIspendr12"
	case HVGICDistributorRegGICDIspendr13:
		return "HVGICDistributorRegGICDIspendr13"
	case HVGICDistributorRegGICDIspendr14:
		return "HVGICDistributorRegGICDIspendr14"
	case HVGICDistributorRegGICDIspendr15:
		return "HVGICDistributorRegGICDIspendr15"
	case HVGICDistributorRegGICDIspendr16:
		return "HVGICDistributorRegGICDIspendr16"
	case HVGICDistributorRegGICDIspendr17:
		return "HVGICDistributorRegGICDIspendr17"
	case HVGICDistributorRegGICDIspendr18:
		return "HVGICDistributorRegGICDIspendr18"
	case HVGICDistributorRegGICDIspendr19:
		return "HVGICDistributorRegGICDIspendr19"
	case HVGICDistributorRegGICDIspendr2:
		return "HVGICDistributorRegGICDIspendr2"
	case HVGICDistributorRegGICDIspendr20:
		return "HVGICDistributorRegGICDIspendr20"
	case HVGICDistributorRegGICDIspendr21:
		return "HVGICDistributorRegGICDIspendr21"
	case HVGICDistributorRegGICDIspendr22:
		return "HVGICDistributorRegGICDIspendr22"
	case HVGICDistributorRegGICDIspendr23:
		return "HVGICDistributorRegGICDIspendr23"
	case HVGICDistributorRegGICDIspendr24:
		return "HVGICDistributorRegGICDIspendr24"
	case HVGICDistributorRegGICDIspendr25:
		return "HVGICDistributorRegGICDIspendr25"
	case HVGICDistributorRegGICDIspendr26:
		return "HVGICDistributorRegGICDIspendr26"
	case HVGICDistributorRegGICDIspendr27:
		return "HVGICDistributorRegGICDIspendr27"
	case HVGICDistributorRegGICDIspendr28:
		return "HVGICDistributorRegGICDIspendr28"
	case HVGICDistributorRegGICDIspendr29:
		return "HVGICDistributorRegGICDIspendr29"
	case HVGICDistributorRegGICDIspendr3:
		return "HVGICDistributorRegGICDIspendr3"
	case HVGICDistributorRegGICDIspendr30:
		return "HVGICDistributorRegGICDIspendr30"
	case HVGICDistributorRegGICDIspendr31:
		return "HVGICDistributorRegGICDIspendr31"
	case HVGICDistributorRegGICDIspendr4:
		return "HVGICDistributorRegGICDIspendr4"
	case HVGICDistributorRegGICDIspendr5:
		return "HVGICDistributorRegGICDIspendr5"
	case HVGICDistributorRegGICDIspendr6:
		return "HVGICDistributorRegGICDIspendr6"
	case HVGICDistributorRegGICDIspendr7:
		return "HVGICDistributorRegGICDIspendr7"
	case HVGICDistributorRegGICDIspendr8:
		return "HVGICDistributorRegGICDIspendr8"
	case HVGICDistributorRegGICDIspendr9:
		return "HVGICDistributorRegGICDIspendr9"
	case HVGICDistributorRegGICDPidr2:
		return "HVGICDistributorRegGICDPidr2"
	case HVGICDistributorRegGICDTyper:
		return "HVGICDistributorRegGICDTyper"
	default:
		return fmt.Sprintf("HVGICDistributorReg(%d)", e)
	}
}

type HVGICIccReg uint

const (
	HVGICIccRegAp0r0El1   HVGICIccReg = 0xc644
	HVGICIccRegAp1r0El1   HVGICIccReg = 0xc648
	HVGICIccRegBpr0El1    HVGICIccReg = 0xc643
	HVGICIccRegBpr1El1    HVGICIccReg = 0xc663
	HVGICIccRegCtlrEl1    HVGICIccReg = 0xc664
	HVGICIccRegIgrpen0El1 HVGICIccReg = 0xc666
	HVGICIccRegIgrpen1El1 HVGICIccReg = 0xc667
	HVGICIccRegPmrEl1     HVGICIccReg = 0xc230
	HVGICIccRegRprEl1     HVGICIccReg = 0xc65b
	HVGICIccRegSreEl1     HVGICIccReg = 0xc665
	HVGICIccRegSreEl2     HVGICIccReg = 0xe64d
)

func (e HVGICIccReg) String() string {
	switch e {
	case HVGICIccRegAp0r0El1:
		return "HVGICIccRegAp0r0El1"
	case HVGICIccRegAp1r0El1:
		return "HVGICIccRegAp1r0El1"
	case HVGICIccRegBpr0El1:
		return "HVGICIccRegBpr0El1"
	case HVGICIccRegBpr1El1:
		return "HVGICIccRegBpr1El1"
	case HVGICIccRegCtlrEl1:
		return "HVGICIccRegCtlrEl1"
	case HVGICIccRegIgrpen0El1:
		return "HVGICIccRegIgrpen0El1"
	case HVGICIccRegIgrpen1El1:
		return "HVGICIccRegIgrpen1El1"
	case HVGICIccRegPmrEl1:
		return "HVGICIccRegPmrEl1"
	case HVGICIccRegRprEl1:
		return "HVGICIccRegRprEl1"
	case HVGICIccRegSreEl1:
		return "HVGICIccRegSreEl1"
	case HVGICIccRegSreEl2:
		return "HVGICIccRegSreEl2"
	default:
		return fmt.Sprintf("HVGICIccReg(%d)", e)
	}
}

type HVGICIchReg uint

const (
	HVGICIchRegAp0r0El2 HVGICIchReg = 0xe640
	HVGICIchRegAp1r0El2 HVGICIchReg = 0xe648
	HVGICIchRegEisrEl2  HVGICIchReg = 0xe65b
	HVGICIchRegElrsrEl2 HVGICIchReg = 0xe65d
	HVGICIchRegHcrEl2   HVGICIchReg = 0xe658
	HVGICIchRegLr0El2   HVGICIchReg = 0xe660
	HVGICIchRegLr10El2  HVGICIchReg = 0xe66a
	HVGICIchRegLr11El2  HVGICIchReg = 0xe66b
	HVGICIchRegLr12El2  HVGICIchReg = 0xe66c
	HVGICIchRegLr13El2  HVGICIchReg = 0xe66d
	HVGICIchRegLr14El2  HVGICIchReg = 0xe66e
	HVGICIchRegLr15El2  HVGICIchReg = 0xe66f
	HVGICIchRegLr1El2   HVGICIchReg = 0xe661
	HVGICIchRegLr2El2   HVGICIchReg = 0xe662
	HVGICIchRegLr3El2   HVGICIchReg = 0xe663
	HVGICIchRegLr4El2   HVGICIchReg = 0xe664
	HVGICIchRegLr5El2   HVGICIchReg = 0xe665
	HVGICIchRegLr6El2   HVGICIchReg = 0xe666
	HVGICIchRegLr7El2   HVGICIchReg = 0xe667
	HVGICIchRegLr8El2   HVGICIchReg = 0xe668
	HVGICIchRegLr9El2   HVGICIchReg = 0xe669
	HVGICIchRegMisrEl2  HVGICIchReg = 0xe65a
	HVGICIchRegVmcrEl2  HVGICIchReg = 0xe65f
	HVGICIchRegVtrEl2   HVGICIchReg = 0xe659
)

func (e HVGICIchReg) String() string {
	switch e {
	case HVGICIchRegAp0r0El2:
		return "HVGICIchRegAp0r0El2"
	case HVGICIchRegAp1r0El2:
		return "HVGICIchRegAp1r0El2"
	case HVGICIchRegEisrEl2:
		return "HVGICIchRegEisrEl2"
	case HVGICIchRegElrsrEl2:
		return "HVGICIchRegElrsrEl2"
	case HVGICIchRegHcrEl2:
		return "HVGICIchRegHcrEl2"
	case HVGICIchRegLr0El2:
		return "HVGICIchRegLr0El2"
	case HVGICIchRegLr10El2:
		return "HVGICIchRegLr10El2"
	case HVGICIchRegLr11El2:
		return "HVGICIchRegLr11El2"
	case HVGICIchRegLr12El2:
		return "HVGICIchRegLr12El2"
	case HVGICIchRegLr13El2:
		return "HVGICIchRegLr13El2"
	case HVGICIchRegLr14El2:
		return "HVGICIchRegLr14El2"
	case HVGICIchRegLr15El2:
		return "HVGICIchRegLr15El2"
	case HVGICIchRegLr1El2:
		return "HVGICIchRegLr1El2"
	case HVGICIchRegLr2El2:
		return "HVGICIchRegLr2El2"
	case HVGICIchRegLr3El2:
		return "HVGICIchRegLr3El2"
	case HVGICIchRegLr4El2:
		return "HVGICIchRegLr4El2"
	case HVGICIchRegLr5El2:
		return "HVGICIchRegLr5El2"
	case HVGICIchRegLr6El2:
		return "HVGICIchRegLr6El2"
	case HVGICIchRegLr7El2:
		return "HVGICIchRegLr7El2"
	case HVGICIchRegLr8El2:
		return "HVGICIchRegLr8El2"
	case HVGICIchRegLr9El2:
		return "HVGICIchRegLr9El2"
	case HVGICIchRegMisrEl2:
		return "HVGICIchRegMisrEl2"
	case HVGICIchRegVmcrEl2:
		return "HVGICIchRegVmcrEl2"
	case HVGICIchRegVtrEl2:
		return "HVGICIchRegVtrEl2"
	default:
		return fmt.Sprintf("HVGICIchReg(%d)", e)
	}
}

type HVGICIcvReg uint

const (
	HVGICIcvRegAp0r0El1   HVGICIcvReg = 0xc644
	HVGICIcvRegAp1r0El1   HVGICIcvReg = 0xc648
	HVGICIcvRegBpr0El1    HVGICIcvReg = 0xc643
	HVGICIcvRegBpr1El1    HVGICIcvReg = 0xc663
	HVGICIcvRegCtlrEl1    HVGICIcvReg = 0xc664
	HVGICIcvRegIgrpen0El1 HVGICIcvReg = 0xc666
	HVGICIcvRegIgrpen1El1 HVGICIcvReg = 0xc667
	HVGICIcvRegPmrEl1     HVGICIcvReg = 0xc230
	HVGICIcvRegRprEl1     HVGICIcvReg = 0xc65b
	HVGICIcvRegSreEl1     HVGICIcvReg = 0xc665
)

func (e HVGICIcvReg) String() string {
	switch e {
	case HVGICIcvRegAp0r0El1:
		return "HVGICIcvRegAp0r0El1"
	case HVGICIcvRegAp1r0El1:
		return "HVGICIcvRegAp1r0El1"
	case HVGICIcvRegBpr0El1:
		return "HVGICIcvRegBpr0El1"
	case HVGICIcvRegBpr1El1:
		return "HVGICIcvRegBpr1El1"
	case HVGICIcvRegCtlrEl1:
		return "HVGICIcvRegCtlrEl1"
	case HVGICIcvRegIgrpen0El1:
		return "HVGICIcvRegIgrpen0El1"
	case HVGICIcvRegIgrpen1El1:
		return "HVGICIcvRegIgrpen1El1"
	case HVGICIcvRegPmrEl1:
		return "HVGICIcvRegPmrEl1"
	case HVGICIcvRegRprEl1:
		return "HVGICIcvRegRprEl1"
	case HVGICIcvRegSreEl1:
		return "HVGICIcvRegSreEl1"
	default:
		return fmt.Sprintf("HVGICIcvReg(%d)", e)
	}
}

type HVGICIntid uint

const (
	HVGICIntEl1PhysicalTimer HVGICIntid = 30
	HVGICIntEl1VirtualTimer  HVGICIntid = 27
	HVGICIntEl2PhysicalTimer HVGICIntid = 26
	// HVGICIntMaintenance: A register Hypervisor uses to signal virtual Interrupts (vIRQs) that the framework sends to guests running at exception level 2 (EL2).
	HVGICIntMaintenance HVGICIntid = 25
	// HVGICIntPerformanceMonitor: A register the framework uses to count GIC related events.
	HVGICIntPerformanceMonitor HVGICIntid = 23
)

func (e HVGICIntid) String() string {
	switch e {
	case HVGICIntEl1PhysicalTimer:
		return "HVGICIntEl1PhysicalTimer"
	case HVGICIntEl1VirtualTimer:
		return "HVGICIntEl1VirtualTimer"
	case HVGICIntEl2PhysicalTimer:
		return "HVGICIntEl2PhysicalTimer"
	case HVGICIntMaintenance:
		return "HVGICIntMaintenance"
	case HVGICIntPerformanceMonitor:
		return "HVGICIntPerformanceMonitor"
	default:
		return fmt.Sprintf("HVGICIntid(%d)", e)
	}
}

type HVGICMsiReg uint

const (
	HVGICRegGicmSetSpiNsr HVGICMsiReg = 0x40
	HVGICRegGicmTyper     HVGICMsiReg = 0x8
)

func (e HVGICMsiReg) String() string {
	switch e {
	case HVGICRegGicmSetSpiNsr:
		return "HVGICRegGicmSetSpiNsr"
	case HVGICRegGicmTyper:
		return "HVGICRegGicmTyper"
	default:
		return fmt.Sprintf("HVGICMsiReg(%d)", e)
	}
}

type HVGICRedistributorReg uint

const (
	HVGICRedistributorRegGICRIcactiver0  HVGICRedistributorReg = 0x10380
	HVGICRedistributorRegGICRIcenabler0  HVGICRedistributorReg = 0x10180
	HVGICRedistributorRegGICRIcfgr0      HVGICRedistributorReg = 0x10c00
	HVGICRedistributorRegGICRIcfgr1      HVGICRedistributorReg = 0x10c04
	HVGICRedistributorRegGICRIcpendr0    HVGICRedistributorReg = 0x10280
	HVGICRedistributorRegGICRIgroupr0    HVGICRedistributorReg = 0x10080
	HVGICRedistributorRegGICRIpriorityr0 HVGICRedistributorReg = 0x10400
	HVGICRedistributorRegGICRIpriorityr1 HVGICRedistributorReg = 0x10404
	HVGICRedistributorRegGICRIpriorityr2 HVGICRedistributorReg = 0x10408
	HVGICRedistributorRegGICRIpriorityr3 HVGICRedistributorReg = 0x1040c
	HVGICRedistributorRegGICRIpriorityr4 HVGICRedistributorReg = 0x10410
	HVGICRedistributorRegGICRIpriorityr5 HVGICRedistributorReg = 0x10414
	HVGICRedistributorRegGICRIpriorityr6 HVGICRedistributorReg = 0x10418
	HVGICRedistributorRegGICRIpriorityr7 HVGICRedistributorReg = 0x1041c
	HVGICRedistributorRegGICRIsactiver0  HVGICRedistributorReg = 0x10300
	HVGICRedistributorRegGICRIsenabler0  HVGICRedistributorReg = 0x10100
	HVGICRedistributorRegGICRIspendr0    HVGICRedistributorReg = 0x10200
	HVGICRedistributorRegGICRPidr2       HVGICRedistributorReg = 0xffe8
	HVGICRedistributorRegGICRTyper       HVGICRedistributorReg = 0x8
)

func (e HVGICRedistributorReg) String() string {
	switch e {
	case HVGICRedistributorRegGICRIcactiver0:
		return "HVGICRedistributorRegGICRIcactiver0"
	case HVGICRedistributorRegGICRIcenabler0:
		return "HVGICRedistributorRegGICRIcenabler0"
	case HVGICRedistributorRegGICRIcfgr0:
		return "HVGICRedistributorRegGICRIcfgr0"
	case HVGICRedistributorRegGICRIcfgr1:
		return "HVGICRedistributorRegGICRIcfgr1"
	case HVGICRedistributorRegGICRIcpendr0:
		return "HVGICRedistributorRegGICRIcpendr0"
	case HVGICRedistributorRegGICRIgroupr0:
		return "HVGICRedistributorRegGICRIgroupr0"
	case HVGICRedistributorRegGICRIpriorityr0:
		return "HVGICRedistributorRegGICRIpriorityr0"
	case HVGICRedistributorRegGICRIpriorityr1:
		return "HVGICRedistributorRegGICRIpriorityr1"
	case HVGICRedistributorRegGICRIpriorityr2:
		return "HVGICRedistributorRegGICRIpriorityr2"
	case HVGICRedistributorRegGICRIpriorityr3:
		return "HVGICRedistributorRegGICRIpriorityr3"
	case HVGICRedistributorRegGICRIpriorityr4:
		return "HVGICRedistributorRegGICRIpriorityr4"
	case HVGICRedistributorRegGICRIpriorityr5:
		return "HVGICRedistributorRegGICRIpriorityr5"
	case HVGICRedistributorRegGICRIpriorityr6:
		return "HVGICRedistributorRegGICRIpriorityr6"
	case HVGICRedistributorRegGICRIpriorityr7:
		return "HVGICRedistributorRegGICRIpriorityr7"
	case HVGICRedistributorRegGICRIsactiver0:
		return "HVGICRedistributorRegGICRIsactiver0"
	case HVGICRedistributorRegGICRIsenabler0:
		return "HVGICRedistributorRegGICRIsenabler0"
	case HVGICRedistributorRegGICRIspendr0:
		return "HVGICRedistributorRegGICRIspendr0"
	case HVGICRedistributorRegGICRPidr2:
		return "HVGICRedistributorRegGICRPidr2"
	case HVGICRedistributorRegGICRTyper:
		return "HVGICRedistributorRegGICRTyper"
	default:
		return fmt.Sprintf("HVGICRedistributorReg(%d)", e)
	}
}

type HVIPAGranule uint

const (
	// HVIPAGranule16kb: # Discussion
	HVIPAGranule16kb HVIPAGranule = 1
	// HVIPAGranule4kb: # Discussion
	HVIPAGranule4kb HVIPAGranule = 0
)

func (e HVIPAGranule) String() string {
	switch e {
	case HVIPAGranule16kb:
		return "HVIPAGranule16kb"
	case HVIPAGranule4kb:
		return "HVIPAGranule4kb"
	default:
		return fmt.Sprintf("HVIPAGranule(%d)", e)
	}
}

type HVInterruptType uint

const (
	// HVInterruptTypeFiq: ARM Fast Interrupt Request.
	HVInterruptTypeFiq HVInterruptType = 1
	// HVInterruptTypeIrq: ARM Interrupt Request.
	HVInterruptTypeIrq HVInterruptType = 0
)

func (e HVInterruptType) String() string {
	switch e {
	case HVInterruptTypeFiq:
		return "HVInterruptTypeFiq"
	case HVInterruptTypeIrq:
		return "HVInterruptTypeIrq"
	default:
		return fmt.Sprintf("HVInterruptType(%d)", e)
	}
}

type HVReg uint

const (
	// HVRegCpsr: The value that identifies the current program status register (CPSR).
	HVRegCpsr HVReg = 34
	// HVRegFP: The value that identifies the frame pointer (FP).
	HVRegFP HVReg = 29
	// HVRegFpcr: The value that identifies the floating-point control register (FPCR).
	HVRegFpcr HVReg = 32
	// HVRegFpsr: The value that identifies the floating-point status register (FPSR).
	HVRegFpsr HVReg = 33
	// HVRegLr: The value that identifies the link register (LR).
	HVRegLr HVReg = 30
	// HVRegPC: The value that identifies the program counter (PC).
	HVRegPC HVReg = 31
	// HVRegX0: The value that identifies register X0.
	HVRegX0 HVReg = 0
	// HVRegX1: The value that identifies register X1.
	HVRegX1 HVReg = 1
	// HVRegX10: The value that identifies register X10.
	HVRegX10 HVReg = 10
	// HVRegX11: The value that identifies register X11.
	HVRegX11 HVReg = 11
	// HVRegX12: The value that identifies register X12.
	HVRegX12 HVReg = 12
	// HVRegX13: The value that identifies register X13.
	HVRegX13 HVReg = 13
	// HVRegX14: The value that identifies register X14.
	HVRegX14 HVReg = 14
	// HVRegX15: The value that identifies register X15.
	HVRegX15 HVReg = 15
	// HVRegX16: The value that identifies register X16.
	HVRegX16 HVReg = 16
	// HVRegX17: The value that identifies register X17.
	HVRegX17 HVReg = 17
	// HVRegX18: The value that identifies register X18.
	HVRegX18 HVReg = 18
	// HVRegX19: The value that identifies register X19.
	HVRegX19 HVReg = 19
	// HVRegX2: The value that identifies register X2.
	HVRegX2 HVReg = 2
	// HVRegX20: The value that identifies register X20.
	HVRegX20 HVReg = 20
	// HVRegX21: The value that identifies register X21.
	HVRegX21 HVReg = 21
	// HVRegX22: The value that identifies register X22.
	HVRegX22 HVReg = 22
	// HVRegX23: The value that identifies register X23.
	HVRegX23 HVReg = 23
	// HVRegX24: The value that identifies register X24.
	HVRegX24 HVReg = 24
	// HVRegX25: The value that identifies register X25.
	HVRegX25 HVReg = 25
	// HVRegX26: The value that identifies register X26.
	HVRegX26 HVReg = 26
	// HVRegX27: The value that identifies register X27.
	HVRegX27 HVReg = 27
	// HVRegX28: The value that identifies register X28.
	HVRegX28 HVReg = 28
	// HVRegX29: The value that identifies register X29.
	HVRegX29 HVReg = 29
	// HVRegX3: The value that identifies register X3.
	HVRegX3 HVReg = 3
	// HVRegX30: The value that identifies register X30.
	HVRegX30 HVReg = 30
	// HVRegX4: The value that identifies register X4.
	HVRegX4 HVReg = 4
	// HVRegX5: The value that identifies register X5.
	HVRegX5 HVReg = 5
	// HVRegX6: The value that identifies register X6.
	HVRegX6 HVReg = 6
	// HVRegX7: The value that identifies register X7.
	HVRegX7 HVReg = 7
	// HVRegX8: The value that identifies register X8.
	HVRegX8 HVReg = 8
	// HVRegX9: The value that identifies register X9.
	HVRegX9 HVReg = 9
)

func (e HVReg) String() string {
	switch e {
	case HVRegCpsr:
		return "HVRegCpsr"
	case HVRegFP:
		return "HVRegFP"
	case HVRegFpcr:
		return "HVRegFpcr"
	case HVRegFpsr:
		return "HVRegFpsr"
	case HVRegLr:
		return "HVRegLr"
	case HVRegPC:
		return "HVRegPC"
	case HVRegX0:
		return "HVRegX0"
	case HVRegX1:
		return "HVRegX1"
	case HVRegX10:
		return "HVRegX10"
	case HVRegX11:
		return "HVRegX11"
	case HVRegX12:
		return "HVRegX12"
	case HVRegX13:
		return "HVRegX13"
	case HVRegX14:
		return "HVRegX14"
	case HVRegX15:
		return "HVRegX15"
	case HVRegX16:
		return "HVRegX16"
	case HVRegX17:
		return "HVRegX17"
	case HVRegX18:
		return "HVRegX18"
	case HVRegX19:
		return "HVRegX19"
	case HVRegX2:
		return "HVRegX2"
	case HVRegX20:
		return "HVRegX20"
	case HVRegX21:
		return "HVRegX21"
	case HVRegX22:
		return "HVRegX22"
	case HVRegX23:
		return "HVRegX23"
	case HVRegX24:
		return "HVRegX24"
	case HVRegX25:
		return "HVRegX25"
	case HVRegX26:
		return "HVRegX26"
	case HVRegX27:
		return "HVRegX27"
	case HVRegX28:
		return "HVRegX28"
	case HVRegX3:
		return "HVRegX3"
	case HVRegX4:
		return "HVRegX4"
	case HVRegX5:
		return "HVRegX5"
	case HVRegX6:
		return "HVRegX6"
	case HVRegX7:
		return "HVRegX7"
	case HVRegX8:
		return "HVRegX8"
	case HVRegX9:
		return "HVRegX9"
	default:
		return fmt.Sprintf("HVReg(%d)", e)
	}
}

type HVReturn uint

const (
	// HVBadArgument: The operation was unsuccessful because the function call had an invalid argument.
	HVBadArgument HVReturn = 0xfae94003
	// HVBusy: The operation was unsuccessful because the owning resource was busy.
	HVBusy HVReturn = 0xfae94002
	// HVDenied: The system didn’t allow the requested operation.
	HVDenied HVReturn = 0xfae94007
	// HVError: The operation was unsuccessful.
	HVError HVReturn = 0xfae94001
	HVFault HVReturn = 0xfae94008
	// HVNoDevice: The operation was unsuccessful because no VM or vCPU was available.
	HVNoDevice HVReturn = 0xfae94006
	// HVNoResources: The operation was unsuccessful because the host had no resources available to complete the request.
	HVNoResources HVReturn = 0xfae94005
	// HVSuccess: The operation completed successfully.
	HVSuccess HVReturn = 0
	// HVUnsupported: The operation requested isn’t supported by the hypervisor.
	HVUnsupported HVReturn = 0xfae9400f
)

func (e HVReturn) String() string {
	switch e {
	case HVBadArgument:
		return "HVBadArgument"
	case HVBusy:
		return "HVBusy"
	case HVDenied:
		return "HVDenied"
	case HVError:
		return "HVError"
	case HVFault:
		return "HVFault"
	case HVNoDevice:
		return "HVNoDevice"
	case HVNoResources:
		return "HVNoResources"
	case HVSuccess:
		return "HVSuccess"
	case HVUnsupported:
		return "HVUnsupported"
	default:
		return fmt.Sprintf("HVReturn(%d)", e)
	}
}

type HVSIMDFPReg uint

const (
	// HVSIMDFPRegQ0: The value representing SIMD register Q0.
	HVSIMDFPRegQ0 HVSIMDFPReg = 0
	// HVSIMDFPRegQ1: The value representing SIMD register Q1.
	HVSIMDFPRegQ1 HVSIMDFPReg = 1
	// HVSIMDFPRegQ10: The value representing SIMD register Q10.
	HVSIMDFPRegQ10 HVSIMDFPReg = 10
	// HVSIMDFPRegQ11: The value representing SIMD register Q11.
	HVSIMDFPRegQ11 HVSIMDFPReg = 11
	// HVSIMDFPRegQ12: The value representing SIMD register Q12.
	HVSIMDFPRegQ12 HVSIMDFPReg = 12
	// HVSIMDFPRegQ13: The value representing SIMD register Q13.
	HVSIMDFPRegQ13 HVSIMDFPReg = 13
	// HVSIMDFPRegQ14: The value representing SIMD register Q14.
	HVSIMDFPRegQ14 HVSIMDFPReg = 14
	// HVSIMDFPRegQ15: The value representing SIMD register Q15.
	HVSIMDFPRegQ15 HVSIMDFPReg = 15
	// HVSIMDFPRegQ16: The value representing SIMD register Q16.
	HVSIMDFPRegQ16 HVSIMDFPReg = 16
	// HVSIMDFPRegQ17: The value representing SIMD register Q17.
	HVSIMDFPRegQ17 HVSIMDFPReg = 17
	// HVSIMDFPRegQ18: The value representing SIMD register Q18.
	HVSIMDFPRegQ18 HVSIMDFPReg = 18
	// HVSIMDFPRegQ19: The value representing SIMD register Q19.
	HVSIMDFPRegQ19 HVSIMDFPReg = 19
	// HVSIMDFPRegQ2: The value representing SIMD register Q2.
	HVSIMDFPRegQ2 HVSIMDFPReg = 2
	// HVSIMDFPRegQ20: The value representing SIMD register Q20.
	HVSIMDFPRegQ20 HVSIMDFPReg = 20
	// HVSIMDFPRegQ21: The value representing SIMD register Q21.
	HVSIMDFPRegQ21 HVSIMDFPReg = 21
	// HVSIMDFPRegQ22: The value representing SIMD register Q22.
	HVSIMDFPRegQ22 HVSIMDFPReg = 22
	// HVSIMDFPRegQ23: The value representing SIMD register Q23.
	HVSIMDFPRegQ23 HVSIMDFPReg = 23
	// HVSIMDFPRegQ24: The value representing SIMD register Q24.
	HVSIMDFPRegQ24 HVSIMDFPReg = 24
	// HVSIMDFPRegQ25: The value representing SIMD register Q25.
	HVSIMDFPRegQ25 HVSIMDFPReg = 25
	// HVSIMDFPRegQ26: The value representing SIMD register Q26.
	HVSIMDFPRegQ26 HVSIMDFPReg = 26
	// HVSIMDFPRegQ27: The value representing SIMD register Q27.
	HVSIMDFPRegQ27 HVSIMDFPReg = 27
	// HVSIMDFPRegQ28: The value representing SIMD register Q28.
	HVSIMDFPRegQ28 HVSIMDFPReg = 28
	// HVSIMDFPRegQ29: The value representing SIMD register Q29.
	HVSIMDFPRegQ29 HVSIMDFPReg = 29
	// HVSIMDFPRegQ3: The value representing SIMD register Q3.
	HVSIMDFPRegQ3 HVSIMDFPReg = 3
	// HVSIMDFPRegQ30: The value representing SIMD register Q30.
	HVSIMDFPRegQ30 HVSIMDFPReg = 30
	// HVSIMDFPRegQ31: The value representing SIMD register Q31.
	HVSIMDFPRegQ31 HVSIMDFPReg = 31
	// HVSIMDFPRegQ4: The value representing SIMD register Q4.
	HVSIMDFPRegQ4 HVSIMDFPReg = 4
	// HVSIMDFPRegQ5: The value representing SIMD register Q5.
	HVSIMDFPRegQ5 HVSIMDFPReg = 5
	// HVSIMDFPRegQ6: The value representing SIMD register Q6.
	HVSIMDFPRegQ6 HVSIMDFPReg = 6
	// HVSIMDFPRegQ7: The value representing SIMD register Q7.
	HVSIMDFPRegQ7 HVSIMDFPReg = 7
	// HVSIMDFPRegQ8: The value representing SIMD register Q8.
	HVSIMDFPRegQ8 HVSIMDFPReg = 8
	// HVSIMDFPRegQ9: The value representing SIMD register Q9.
	HVSIMDFPRegQ9 HVSIMDFPReg = 9
)

func (e HVSIMDFPReg) String() string {
	switch e {
	case HVSIMDFPRegQ0:
		return "HVSIMDFPRegQ0"
	case HVSIMDFPRegQ1:
		return "HVSIMDFPRegQ1"
	case HVSIMDFPRegQ10:
		return "HVSIMDFPRegQ10"
	case HVSIMDFPRegQ11:
		return "HVSIMDFPRegQ11"
	case HVSIMDFPRegQ12:
		return "HVSIMDFPRegQ12"
	case HVSIMDFPRegQ13:
		return "HVSIMDFPRegQ13"
	case HVSIMDFPRegQ14:
		return "HVSIMDFPRegQ14"
	case HVSIMDFPRegQ15:
		return "HVSIMDFPRegQ15"
	case HVSIMDFPRegQ16:
		return "HVSIMDFPRegQ16"
	case HVSIMDFPRegQ17:
		return "HVSIMDFPRegQ17"
	case HVSIMDFPRegQ18:
		return "HVSIMDFPRegQ18"
	case HVSIMDFPRegQ19:
		return "HVSIMDFPRegQ19"
	case HVSIMDFPRegQ2:
		return "HVSIMDFPRegQ2"
	case HVSIMDFPRegQ20:
		return "HVSIMDFPRegQ20"
	case HVSIMDFPRegQ21:
		return "HVSIMDFPRegQ21"
	case HVSIMDFPRegQ22:
		return "HVSIMDFPRegQ22"
	case HVSIMDFPRegQ23:
		return "HVSIMDFPRegQ23"
	case HVSIMDFPRegQ24:
		return "HVSIMDFPRegQ24"
	case HVSIMDFPRegQ25:
		return "HVSIMDFPRegQ25"
	case HVSIMDFPRegQ26:
		return "HVSIMDFPRegQ26"
	case HVSIMDFPRegQ27:
		return "HVSIMDFPRegQ27"
	case HVSIMDFPRegQ28:
		return "HVSIMDFPRegQ28"
	case HVSIMDFPRegQ29:
		return "HVSIMDFPRegQ29"
	case HVSIMDFPRegQ3:
		return "HVSIMDFPRegQ3"
	case HVSIMDFPRegQ30:
		return "HVSIMDFPRegQ30"
	case HVSIMDFPRegQ31:
		return "HVSIMDFPRegQ31"
	case HVSIMDFPRegQ4:
		return "HVSIMDFPRegQ4"
	case HVSIMDFPRegQ5:
		return "HVSIMDFPRegQ5"
	case HVSIMDFPRegQ6:
		return "HVSIMDFPRegQ6"
	case HVSIMDFPRegQ7:
		return "HVSIMDFPRegQ7"
	case HVSIMDFPRegQ8:
		return "HVSIMDFPRegQ8"
	case HVSIMDFPRegQ9:
		return "HVSIMDFPRegQ9"
	default:
		return fmt.Sprintf("HVSIMDFPReg(%d)", e)
	}
}

type HVSMEPReg uint

const (
	HVSMEPReg0  HVSMEPReg = 0
	HVSMEPReg1  HVSMEPReg = 1
	HVSMEPReg10 HVSMEPReg = 10
	HVSMEPReg11 HVSMEPReg = 11
	HVSMEPReg12 HVSMEPReg = 12
	HVSMEPReg13 HVSMEPReg = 13
	HVSMEPReg14 HVSMEPReg = 14
	HVSMEPReg15 HVSMEPReg = 15
	HVSMEPReg2  HVSMEPReg = 2
	HVSMEPReg3  HVSMEPReg = 3
	HVSMEPReg4  HVSMEPReg = 4
	HVSMEPReg5  HVSMEPReg = 5
	HVSMEPReg6  HVSMEPReg = 6
	HVSMEPReg7  HVSMEPReg = 7
	HVSMEPReg8  HVSMEPReg = 8
	HVSMEPReg9  HVSMEPReg = 9
)

func (e HVSMEPReg) String() string {
	switch e {
	case HVSMEPReg0:
		return "HVSMEPReg0"
	case HVSMEPReg1:
		return "HVSMEPReg1"
	case HVSMEPReg10:
		return "HVSMEPReg10"
	case HVSMEPReg11:
		return "HVSMEPReg11"
	case HVSMEPReg12:
		return "HVSMEPReg12"
	case HVSMEPReg13:
		return "HVSMEPReg13"
	case HVSMEPReg14:
		return "HVSMEPReg14"
	case HVSMEPReg15:
		return "HVSMEPReg15"
	case HVSMEPReg2:
		return "HVSMEPReg2"
	case HVSMEPReg3:
		return "HVSMEPReg3"
	case HVSMEPReg4:
		return "HVSMEPReg4"
	case HVSMEPReg5:
		return "HVSMEPReg5"
	case HVSMEPReg6:
		return "HVSMEPReg6"
	case HVSMEPReg7:
		return "HVSMEPReg7"
	case HVSMEPReg8:
		return "HVSMEPReg8"
	case HVSMEPReg9:
		return "HVSMEPReg9"
	default:
		return fmt.Sprintf("HVSMEPReg(%d)", e)
	}
}

type HVSMEZReg uint

const (
	HVSMEZReg0  HVSMEZReg = 0
	HVSMEZReg1  HVSMEZReg = 1
	HVSMEZReg10 HVSMEZReg = 10
	HVSMEZReg11 HVSMEZReg = 11
	HVSMEZReg12 HVSMEZReg = 12
	HVSMEZReg13 HVSMEZReg = 13
	HVSMEZReg14 HVSMEZReg = 14
	HVSMEZReg15 HVSMEZReg = 15
	HVSMEZReg16 HVSMEZReg = 16
	HVSMEZReg17 HVSMEZReg = 17
	HVSMEZReg18 HVSMEZReg = 18
	HVSMEZReg19 HVSMEZReg = 19
	HVSMEZReg2  HVSMEZReg = 2
	HVSMEZReg20 HVSMEZReg = 20
	HVSMEZReg21 HVSMEZReg = 21
	HVSMEZReg22 HVSMEZReg = 22
	HVSMEZReg23 HVSMEZReg = 23
	HVSMEZReg24 HVSMEZReg = 24
	HVSMEZReg25 HVSMEZReg = 25
	HVSMEZReg26 HVSMEZReg = 26
	HVSMEZReg27 HVSMEZReg = 27
	HVSMEZReg28 HVSMEZReg = 28
	HVSMEZReg29 HVSMEZReg = 29
	HVSMEZReg3  HVSMEZReg = 3
	HVSMEZReg30 HVSMEZReg = 30
	HVSMEZReg31 HVSMEZReg = 31
	HVSMEZReg4  HVSMEZReg = 4
	HVSMEZReg5  HVSMEZReg = 5
	HVSMEZReg6  HVSMEZReg = 6
	HVSMEZReg7  HVSMEZReg = 7
	HVSMEZReg8  HVSMEZReg = 8
	HVSMEZReg9  HVSMEZReg = 9
)

func (e HVSMEZReg) String() string {
	switch e {
	case HVSMEZReg0:
		return "HVSMEZReg0"
	case HVSMEZReg1:
		return "HVSMEZReg1"
	case HVSMEZReg10:
		return "HVSMEZReg10"
	case HVSMEZReg11:
		return "HVSMEZReg11"
	case HVSMEZReg12:
		return "HVSMEZReg12"
	case HVSMEZReg13:
		return "HVSMEZReg13"
	case HVSMEZReg14:
		return "HVSMEZReg14"
	case HVSMEZReg15:
		return "HVSMEZReg15"
	case HVSMEZReg16:
		return "HVSMEZReg16"
	case HVSMEZReg17:
		return "HVSMEZReg17"
	case HVSMEZReg18:
		return "HVSMEZReg18"
	case HVSMEZReg19:
		return "HVSMEZReg19"
	case HVSMEZReg2:
		return "HVSMEZReg2"
	case HVSMEZReg20:
		return "HVSMEZReg20"
	case HVSMEZReg21:
		return "HVSMEZReg21"
	case HVSMEZReg22:
		return "HVSMEZReg22"
	case HVSMEZReg23:
		return "HVSMEZReg23"
	case HVSMEZReg24:
		return "HVSMEZReg24"
	case HVSMEZReg25:
		return "HVSMEZReg25"
	case HVSMEZReg26:
		return "HVSMEZReg26"
	case HVSMEZReg27:
		return "HVSMEZReg27"
	case HVSMEZReg28:
		return "HVSMEZReg28"
	case HVSMEZReg29:
		return "HVSMEZReg29"
	case HVSMEZReg3:
		return "HVSMEZReg3"
	case HVSMEZReg30:
		return "HVSMEZReg30"
	case HVSMEZReg31:
		return "HVSMEZReg31"
	case HVSMEZReg4:
		return "HVSMEZReg4"
	case HVSMEZReg5:
		return "HVSMEZReg5"
	case HVSMEZReg6:
		return "HVSMEZReg6"
	case HVSMEZReg7:
		return "HVSMEZReg7"
	case HVSMEZReg8:
		return "HVSMEZReg8"
	case HVSMEZReg9:
		return "HVSMEZReg9"
	default:
		return fmt.Sprintf("HVSMEZReg(%d)", e)
	}
}

type HVSysReg uint

const (
	HVSysRegActlrEl1 HVSysReg = 0xc081
	// HVSysRegAfsr0El1: The value that represents the system register AFSR0_EL1.
	HVSysRegAfsr0El1 HVSysReg = 0xc288
	// HVSysRegAfsr1El1: The value that represents the system register AFSR1_EL1.
	HVSysRegAfsr1El1 HVSysReg = 0xc289
	// HVSysRegAmairEl1: The value that represents the system register AMAIR_EL1.
	HVSysRegAmairEl1 HVSysReg = 0xc518
	// HVSysRegApdakeyhiEl1: The value that represents the system register APDAKEYHI_EL1.
	HVSysRegApdakeyhiEl1 HVSysReg = 0xc111
	// HVSysRegApdakeyloEl1: The value that represents the system register APDAKEYLO_E1.
	HVSysRegApdakeyloEl1 HVSysReg = 0xc110
	// HVSysRegApdbkeyhiEl1: The value that represents the system register ADPBKEYHI_EL1.
	HVSysRegApdbkeyhiEl1 HVSysReg = 0xc113
	// HVSysRegApdbkeyloEl1: The value that represents the system register APDBKEYLO_E1.
	HVSysRegApdbkeyloEl1 HVSysReg = 0xc112
	// HVSysRegApgakeyhiEl1: The value that represents the system register AOGAKEYHI_EL1.
	HVSysRegApgakeyhiEl1 HVSysReg = 0xc119
	// HVSysRegApgakeyloEl1: The value that represents the system register APGAKEYLO_REL1.
	HVSysRegApgakeyloEl1 HVSysReg = 0xc118
	// HVSysRegApiakeyhiEl1: The value that represents the system register APIAKEYHI_EL1.
	HVSysRegApiakeyhiEl1 HVSysReg = 0xc109
	// HVSysRegApiakeyloEl1: The value that represents the system register APIAKEYLO_EL1.
	HVSysRegApiakeyloEl1 HVSysReg = 0xc108
	// HVSysRegApibkeyhiEl1: The value that represents the system register APIBKEYHI_EL1.
	HVSysRegApibkeyhiEl1 HVSysReg = 0xc10b
	// HVSysRegApibkeyloEl1: The value that represents the system register APIBKEYLO_EL1.
	HVSysRegApibkeyloEl1 HVSysReg = 0xc10a
	HVSysRegCnthctlEl2   HVSysReg = 0xe708
	HVSysRegCnthpCtlEl2  HVSysReg = 0xe711
	HVSysRegCnthpCvalEl2 HVSysReg = 0xe712
	HVSysRegCnthpTvalEl2 HVSysReg = 0xe710
	// HVSysRegCntkctlEl1: The value that represents the system register CNTKCTL_EL1.
	HVSysRegCntkctlEl1  HVSysReg = 0xc708
	HVSysRegCntpCtlEl0  HVSysReg = 0xdf11
	HVSysRegCntpCvalEl0 HVSysReg = 0xdf12
	HVSysRegCntpTvalEl0 HVSysReg = 0xdf10
	// HVSysRegCntvCtlEl0: The value that represents the system register CNTV_CRTL_EL0.
	HVSysRegCntvCtlEl0 HVSysReg = 0xdf19
	// HVSysRegCntvCvalEl0: The value that represents the system register CNTV_CVAL_EL0.
	HVSysRegCntvCvalEl0 HVSysReg = 0xdf1a
	HVSysRegCntvoffEl2  HVSysReg = 0xe703
	// HVSysRegContextidrEl1: The value that represents the system register CONTEXTIDR_EL1.
	HVSysRegContextidrEl1 HVSysReg = 0xc681
	// HVSysRegCpacrEl1: The value that represents the system register CPACR_EL1.
	HVSysRegCpacrEl1 HVSysReg = 0xc082
	HVSysRegCptrEl2  HVSysReg = 0xe08a
	// HVSysRegCsselrEl1: The value that represents the system register CSSELR_EL1.
	HVSysRegCsselrEl1 HVSysReg = 0xd000
	// HVSysRegDbgbcr0El1: The value that represents the system register DBGBCR0_EL1.
	HVSysRegDbgbcr0El1 HVSysReg = 0x8005
	// HVSysRegDbgbcr10El1: The value that represents the system register DBGBCR10_EL1.
	HVSysRegDbgbcr10El1 HVSysReg = 0x8055
	// HVSysRegDbgbcr11El1: The value that represents the system register DBGBCR11_EL1.
	HVSysRegDbgbcr11El1 HVSysReg = 0x805d
	// HVSysRegDbgbcr12El1: The value that represents the system register DBGBCR12_EL1.
	HVSysRegDbgbcr12El1 HVSysReg = 0x8065
	// HVSysRegDbgbcr13El1: The value that represents the system register DBGBCR13_EL1.
	HVSysRegDbgbcr13El1 HVSysReg = 0x806d
	// HVSysRegDbgbcr14El1: The value that represents the system register DBGBCR14_EL1.
	HVSysRegDbgbcr14El1 HVSysReg = 0x8075
	// HVSysRegDbgbcr15El1: The value that represents the system register DBGBCR15_EL1.
	HVSysRegDbgbcr15El1 HVSysReg = 0x807d
	// HVSysRegDbgbcr1El1: The value that represents the system register DBGBCR1_EL1.
	HVSysRegDbgbcr1El1 HVSysReg = 0x800d
	// HVSysRegDbgbcr2El1: The value that represents the system register DBGBCR1_EL1.
	HVSysRegDbgbcr2El1 HVSysReg = 0x8015
	// HVSysRegDbgbcr3El1: The value that represents the system register DBGBCR3_EL1.
	HVSysRegDbgbcr3El1 HVSysReg = 0x801d
	// HVSysRegDbgbcr4El1: The value that represents the system register DBGBCR4_EL1.
	HVSysRegDbgbcr4El1 HVSysReg = 0x8025
	// HVSysRegDbgbcr5El1: The value that represents the system register DBGBCR5_EL1.
	HVSysRegDbgbcr5El1 HVSysReg = 0x802d
	// HVSysRegDbgbcr6El1: The value that represents the system register DBGBCR6_EL1.
	HVSysRegDbgbcr6El1 HVSysReg = 0x8035
	// HVSysRegDbgbcr7El1: The value that represents the system register DBGBCR7_EL1.
	HVSysRegDbgbcr7El1 HVSysReg = 0x803d
	// HVSysRegDbgbcr8El1: The value that represents the system register DBGBCR8_EL1.
	HVSysRegDbgbcr8El1 HVSysReg = 0x8045
	// HVSysRegDbgbcr9El1: The value that represents the system register DBGBCR9_EL1.
	HVSysRegDbgbcr9El1 HVSysReg = 0x804d
	// HVSysRegDbgbvr0El1: The value that represents the system register DBGBVR0_EL1.
	HVSysRegDbgbvr0El1 HVSysReg = 0x8004
	// HVSysRegDbgbvr10El1: The value that represents the system register DBGBVR10_EL1.
	HVSysRegDbgbvr10El1 HVSysReg = 0x8054
	// HVSysRegDbgbvr11El1: The value that represents the system register DBGBVR11_EL1.
	HVSysRegDbgbvr11El1 HVSysReg = 0x805c
	// HVSysRegDbgbvr12El1: The value that represents the system register DBGBVR12_EL1.
	HVSysRegDbgbvr12El1 HVSysReg = 0x8064
	// HVSysRegDbgbvr13El1: The value that represents the system register DBGBVR13_EL1.
	HVSysRegDbgbvr13El1 HVSysReg = 0x806c
	// HVSysRegDbgbvr14El1: The value that represents the system register DBGBVR14_EL1.
	HVSysRegDbgbvr14El1 HVSysReg = 0x8074
	// HVSysRegDbgbvr15El1: The value that represents the system register DBGBVR15_EL1.
	HVSysRegDbgbvr15El1 HVSysReg = 0x807c
	// HVSysRegDbgbvr1El1: The value that represents the system register DBGBVR1_EL1.
	HVSysRegDbgbvr1El1 HVSysReg = 0x800c
	// HVSysRegDbgbvr2El1: The value that represents the system register DBGBVR2_EL1.
	HVSysRegDbgbvr2El1 HVSysReg = 0x8014
	// HVSysRegDbgbvr3El1: The value that represents the system register DBGBVR3_EL1.
	HVSysRegDbgbvr3El1 HVSysReg = 0x801c
	// HVSysRegDbgbvr4El1: The value that represents the system register DBGBVR4_EL1.
	HVSysRegDbgbvr4El1 HVSysReg = 0x8024
	// HVSysRegDbgbvr5El1: The value that represents the system register DBGBVR5_EL1.
	HVSysRegDbgbvr5El1 HVSysReg = 0x802c
	// HVSysRegDbgbvr6El1: The value that represents the system register DBGBVR6_EL1.
	HVSysRegDbgbvr6El1 HVSysReg = 0x8034
	// HVSysRegDbgbvr7El1: The value that represents the system register DBGBVR7_EL1.
	HVSysRegDbgbvr7El1 HVSysReg = 0x803c
	// HVSysRegDbgbvr8El1: The value that represents the system register DBGBVR8_EL1.
	HVSysRegDbgbvr8El1 HVSysReg = 0x8044
	// HVSysRegDbgbvr9El1: The value that represents the system register DBGBVR9_EL1.
	HVSysRegDbgbvr9El1 HVSysReg = 0x804c
	// HVSysRegDbgwcr0El1: The value that represents the system register DBGWCR0_EL1.
	HVSysRegDbgwcr0El1 HVSysReg = 0x8007
	// HVSysRegDbgwcr10El1: The value that represents the system register DBGWCR10_EL1.
	HVSysRegDbgwcr10El1 HVSysReg = 0x8057
	// HVSysRegDbgwcr11El1: The value that represents the system register DBGWCR11_EL1.
	HVSysRegDbgwcr11El1 HVSysReg = 0x805f
	// HVSysRegDbgwcr12El1: The value that represents the system register DBGWCR12_EL1.
	HVSysRegDbgwcr12El1 HVSysReg = 0x8067
	// HVSysRegDbgwcr13El1: The value that represents the system register DBGWCR13_EL1.
	HVSysRegDbgwcr13El1 HVSysReg = 0x806f
	// HVSysRegDbgwcr14El1: The value that represents the system register DBGWCR14_EL1.
	HVSysRegDbgwcr14El1 HVSysReg = 0x8077
	// HVSysRegDbgwcr15El1: The value that represents the system register DBGWCR15_EL1.
	HVSysRegDbgwcr15El1 HVSysReg = 0x807f
	// HVSysRegDbgwcr1El1: The value that represents the system register DBGWCR1_EL1.
	HVSysRegDbgwcr1El1 HVSysReg = 0x800f
	// HVSysRegDbgwcr2El1: The value that represents the system register DBGWCR2_EL1.
	HVSysRegDbgwcr2El1 HVSysReg = 0x8017
	// HVSysRegDbgwcr3El1: The value that represents the system register DBGWCR3_EL1.
	HVSysRegDbgwcr3El1 HVSysReg = 0x801f
	// HVSysRegDbgwcr4El1: The value that represents the system register DBGWCR4_EL1.
	HVSysRegDbgwcr4El1 HVSysReg = 0x8027
	// HVSysRegDbgwcr5El1: The value that represents the system register DBGWCR5_EL1.
	HVSysRegDbgwcr5El1 HVSysReg = 0x802f
	// HVSysRegDbgwcr6El1: The value that represents the system register DBGWCR6_EL1.
	HVSysRegDbgwcr6El1 HVSysReg = 0x8037
	// HVSysRegDbgwcr7El1: The value that represents the system register DBGWCR7_EL1.
	HVSysRegDbgwcr7El1 HVSysReg = 0x803f
	// HVSysRegDbgwcr8El1: The value that represents the system register DBGWCR8_EL1.
	HVSysRegDbgwcr8El1 HVSysReg = 0x8047
	// HVSysRegDbgwcr9El1: The value that represents the system register DBGWCR9_EL1.
	HVSysRegDbgwcr9El1 HVSysReg = 0x804f
	// HVSysRegDbgwvr0El1: The value that represents the system register DBGWVR0_EL1.
	HVSysRegDbgwvr0El1 HVSysReg = 0x8006
	// HVSysRegDbgwvr10El1: The value that represents the system register DBGWVR10_EL1.
	HVSysRegDbgwvr10El1 HVSysReg = 0x8056
	// HVSysRegDbgwvr11El1: The value that represents the system register DBGWVR11_EL1.
	HVSysRegDbgwvr11El1 HVSysReg = 0x805e
	// HVSysRegDbgwvr12El1: The value that represents the system register DBGWVR12_EL1.
	HVSysRegDbgwvr12El1 HVSysReg = 0x8066
	// HVSysRegDbgwvr13El1: The value that represents the system register DBGWVR13_EL1.
	HVSysRegDbgwvr13El1 HVSysReg = 0x806e
	// HVSysRegDbgwvr14El1: The value that represents the system register DBGWVR14_EL1.
	HVSysRegDbgwvr14El1 HVSysReg = 0x8076
	// HVSysRegDbgwvr15El1: The value that represents the system register DBGWVR15_EL1.
	HVSysRegDbgwvr15El1 HVSysReg = 0x807e
	// HVSysRegDbgwvr1El1: The value that represents the system register DBGWVR1_EL1.
	HVSysRegDbgwvr1El1 HVSysReg = 0x800e
	// HVSysRegDbgwvr2El1: The value that represents the system register DBGWVR2_EL1.
	HVSysRegDbgwvr2El1 HVSysReg = 0x8016
	// HVSysRegDbgwvr3El1: The value that represents the system register DBGWVR3_EL1.
	HVSysRegDbgwvr3El1 HVSysReg = 0x801e
	// HVSysRegDbgwvr4El1: The value that represents the system register DBGWVR4_EL1.
	HVSysRegDbgwvr4El1 HVSysReg = 0x8026
	// HVSysRegDbgwvr5El1: The value that represents the system register DBGWVR5_EL1.
	HVSysRegDbgwvr5El1 HVSysReg = 0x802e
	// HVSysRegDbgwvr6El1: The value that represents the system register DBGWVR6_EL1.
	HVSysRegDbgwvr6El1 HVSysReg = 0x8036
	// HVSysRegDbgwvr7El1: The value that represents the system register DBGWVR7_EL1.
	HVSysRegDbgwvr7El1 HVSysReg = 0x803e
	// HVSysRegDbgwvr8El1: The value that represents the system register DBGWVR8_EL1.
	HVSysRegDbgwvr8El1 HVSysReg = 0x8046
	// HVSysRegDbgwvr9El1: The value that represents the system register DBGWVR9_EL1.
	HVSysRegDbgwvr9El1 HVSysReg = 0x804e
	// HVSysRegElrEl1: The value that represents the system register ELR_EL1.
	HVSysRegElrEl1 HVSysReg = 0xc201
	HVSysRegElrEl2 HVSysReg = 0xe201
	// HVSysRegEsrEl1: The value that represents the system register ESR_EL1.
	HVSysRegEsrEl1 HVSysReg = 0xc290
	HVSysRegEsrEl2 HVSysReg = 0xe290
	// HVSysRegFarEl1: The value that represents the system register FAR_EL1.
	HVSysRegFarEl1   HVSysReg = 0xc300
	HVSysRegFarEl2   HVSysReg = 0xe300
	HVSysRegHcrEl2   HVSysReg = 0xe088
	HVSysRegHpfarEl2 HVSysReg = 0xe304
	// HVSysRegIDAa64dfr0El1: The value that describes the AArch64 Debug Feature Register 0.
	HVSysRegIDAa64dfr0El1 HVSysReg = 0xc028
	// HVSysRegIDAa64dfr1El1: The value that describes the AArch64 Debug Feature Register 1.
	HVSysRegIDAa64dfr1El1 HVSysReg = 0xc029
	// HVSysRegIDAa64isar0El1: The value that describes the AArch64 Instruction Set Attribute Register 0.
	HVSysRegIDAa64isar0El1 HVSysReg = 0xc030
	// HVSysRegIDAa64isar1El1: The value that describes the AArch64 Instruction Set Attribute Register 1.
	HVSysRegIDAa64isar1El1 HVSysReg = 0xc031
	// HVSysRegIDAa64mmfr0El1: The value that describes the AArch64 Memory Model Feature Register 0.
	HVSysRegIDAa64mmfr0El1 HVSysReg = 0xc038
	// HVSysRegIDAa64mmfr1El1: The value that describes the AArch64 Memory Model Feature Register 1.
	HVSysRegIDAa64mmfr1El1 HVSysReg = 0xc039
	// HVSysRegIDAa64mmfr2El1: The value that describes the AArch64 Memory Model Feature Register 2.
	HVSysRegIDAa64mmfr2El1 HVSysReg = 0xc03a
	// HVSysRegIDAa64pfr0El1: The value that describes the AArch64 Processor Feature Register 0.
	HVSysRegIDAa64pfr0El1 HVSysReg = 0xc020
	// HVSysRegIDAa64pfr1El1: The value that describes the AArch64 Processor Feature Register 1.
	HVSysRegIDAa64pfr1El1  HVSysReg = 0xc021
	HVSysRegIDAa64smfr0El1 HVSysReg = 0xc025
	HVSysRegIDAa64zfr0El1  HVSysReg = 0xc024
	// HVSysRegMairEl1: The value that represents the system register MAIR_EL1.
	HVSysRegMairEl1 HVSysReg = 0xc510
	HVSysRegMairEl2 HVSysReg = 0xe510
	// HVSysRegMdccintEl1: The value that represents the system register MDCCINT_EL1.
	HVSysRegMdccintEl1 HVSysReg = 0x8010
	HVSysRegMdcrEl2    HVSysReg = 0xe089
	// HVSysRegMdscrEl1: The value that represents the system register MDSCR_EL0.
	HVSysRegMdscrEl1 HVSysReg = 0x8012
	// HVSysRegMidrEl1: The value that represents the system register MIDR_EL1.
	HVSysRegMidrEl1 HVSysReg = 0xc000
	// HVSysRegMpidrEl1: The value that represents the system register MPIDR_EL1.
	HVSysRegMpidrEl1 HVSysReg = 0xc005
	// HVSysRegParEl1: The value that represents the system register PAR_EL1.
	HVSysRegParEl1 HVSysReg = 0xc3a0
	// HVSysRegSctlrEl1: The value that represents the system register SCTLR_EL1.
	HVSysRegSctlrEl1   HVSysReg = 0xc080
	HVSysRegSctlrEl2   HVSysReg = 0xe080
	HVSysRegScxtnumEl0 HVSysReg = 0xde87
	HVSysRegScxtnumEl1 HVSysReg = 0xc687
	HVSysRegSmcrEl1    HVSysReg = 0xc096
	HVSysRegSmpriEl1   HVSysReg = 0xc094
	// HVSysRegSpEl0: The value that represents the system register SP_EL0.
	HVSysRegSpEl0 HVSysReg = 0xc208
	// HVSysRegSpEl1: The value that represents the system register SP_EL1.
	HVSysRegSpEl1 HVSysReg = 0xe208
	HVSysRegSpEl2 HVSysReg = 0xf208
	// HVSysRegSpsrEl1: The value that represents the system register SPSR_EL1.
	HVSysRegSpsrEl1 HVSysReg = 0xc200
	HVSysRegSpsrEl2 HVSysReg = 0xe200
	// HVSysRegTcrEl1: The value that represents the system register TCR_EL1.
	HVSysRegTcrEl1    HVSysReg = 0xc102
	HVSysRegTcrEl2    HVSysReg = 0xe102
	HVSysRegTpidr2El0 HVSysReg = 0xde85
	// HVSysRegTpidrEl0: The value that represents the system register TPIDR_EL0.
	HVSysRegTpidrEl0 HVSysReg = 0xde82
	// HVSysRegTpidrEl1: The value that represents the system register TPIDR_EL1.
	HVSysRegTpidrEl1 HVSysReg = 0xc684
	HVSysRegTpidrEl2 HVSysReg = 0xe682
	// HVSysRegTpidrroEl0: The value that represents the system register TPIDRRO_EL0.
	HVSysRegTpidrroEl0 HVSysReg = 0xde83
	// HVSysRegTtbr0El1: The value that represents the system register TTBR0_EL1.
	HVSysRegTtbr0El1 HVSysReg = 0xc100
	HVSysRegTtbr0El2 HVSysReg = 0xe100
	// HVSysRegTtbr1El1: The value that represents the system register TTBR1_EL1.
	HVSysRegTtbr1El1 HVSysReg = 0xc101
	HVSysRegTtbr1El2 HVSysReg = 0xe101
	// HVSysRegVbarEl1: The value that represents the system register VBAR_EL1.
	HVSysRegVbarEl1   HVSysReg = 0xc600
	HVSysRegVbarEl2   HVSysReg = 0xe600
	HVSysRegVmpidrEl2 HVSysReg = 0xe005
	HVSysRegVpidrEl2  HVSysReg = 0xe000
	HVSysRegVtcrEl2   HVSysReg = 0xe10a
	HVSysRegVttbrEl2  HVSysReg = 0xe108
)

func (e HVSysReg) String() string {
	switch e {
	case HVSysRegActlrEl1:
		return "HVSysRegActlrEl1"
	case HVSysRegAfsr0El1:
		return "HVSysRegAfsr0El1"
	case HVSysRegAfsr1El1:
		return "HVSysRegAfsr1El1"
	case HVSysRegAmairEl1:
		return "HVSysRegAmairEl1"
	case HVSysRegApdakeyhiEl1:
		return "HVSysRegApdakeyhiEl1"
	case HVSysRegApdakeyloEl1:
		return "HVSysRegApdakeyloEl1"
	case HVSysRegApdbkeyhiEl1:
		return "HVSysRegApdbkeyhiEl1"
	case HVSysRegApdbkeyloEl1:
		return "HVSysRegApdbkeyloEl1"
	case HVSysRegApgakeyhiEl1:
		return "HVSysRegApgakeyhiEl1"
	case HVSysRegApgakeyloEl1:
		return "HVSysRegApgakeyloEl1"
	case HVSysRegApiakeyhiEl1:
		return "HVSysRegApiakeyhiEl1"
	case HVSysRegApiakeyloEl1:
		return "HVSysRegApiakeyloEl1"
	case HVSysRegApibkeyhiEl1:
		return "HVSysRegApibkeyhiEl1"
	case HVSysRegApibkeyloEl1:
		return "HVSysRegApibkeyloEl1"
	case HVSysRegCnthctlEl2:
		return "HVSysRegCnthctlEl2"
	case HVSysRegCnthpCtlEl2:
		return "HVSysRegCnthpCtlEl2"
	case HVSysRegCnthpCvalEl2:
		return "HVSysRegCnthpCvalEl2"
	case HVSysRegCnthpTvalEl2:
		return "HVSysRegCnthpTvalEl2"
	case HVSysRegCntkctlEl1:
		return "HVSysRegCntkctlEl1"
	case HVSysRegCntpCtlEl0:
		return "HVSysRegCntpCtlEl0"
	case HVSysRegCntpCvalEl0:
		return "HVSysRegCntpCvalEl0"
	case HVSysRegCntpTvalEl0:
		return "HVSysRegCntpTvalEl0"
	case HVSysRegCntvCtlEl0:
		return "HVSysRegCntvCtlEl0"
	case HVSysRegCntvCvalEl0:
		return "HVSysRegCntvCvalEl0"
	case HVSysRegCntvoffEl2:
		return "HVSysRegCntvoffEl2"
	case HVSysRegContextidrEl1:
		return "HVSysRegContextidrEl1"
	case HVSysRegCpacrEl1:
		return "HVSysRegCpacrEl1"
	case HVSysRegCptrEl2:
		return "HVSysRegCptrEl2"
	case HVSysRegCsselrEl1:
		return "HVSysRegCsselrEl1"
	case HVSysRegDbgbcr0El1:
		return "HVSysRegDbgbcr0El1"
	case HVSysRegDbgbcr10El1:
		return "HVSysRegDbgbcr10El1"
	case HVSysRegDbgbcr11El1:
		return "HVSysRegDbgbcr11El1"
	case HVSysRegDbgbcr12El1:
		return "HVSysRegDbgbcr12El1"
	case HVSysRegDbgbcr13El1:
		return "HVSysRegDbgbcr13El1"
	case HVSysRegDbgbcr14El1:
		return "HVSysRegDbgbcr14El1"
	case HVSysRegDbgbcr15El1:
		return "HVSysRegDbgbcr15El1"
	case HVSysRegDbgbcr1El1:
		return "HVSysRegDbgbcr1El1"
	case HVSysRegDbgbcr2El1:
		return "HVSysRegDbgbcr2El1"
	case HVSysRegDbgbcr3El1:
		return "HVSysRegDbgbcr3El1"
	case HVSysRegDbgbcr4El1:
		return "HVSysRegDbgbcr4El1"
	case HVSysRegDbgbcr5El1:
		return "HVSysRegDbgbcr5El1"
	case HVSysRegDbgbcr6El1:
		return "HVSysRegDbgbcr6El1"
	case HVSysRegDbgbcr7El1:
		return "HVSysRegDbgbcr7El1"
	case HVSysRegDbgbcr8El1:
		return "HVSysRegDbgbcr8El1"
	case HVSysRegDbgbcr9El1:
		return "HVSysRegDbgbcr9El1"
	case HVSysRegDbgbvr0El1:
		return "HVSysRegDbgbvr0El1"
	case HVSysRegDbgbvr10El1:
		return "HVSysRegDbgbvr10El1"
	case HVSysRegDbgbvr11El1:
		return "HVSysRegDbgbvr11El1"
	case HVSysRegDbgbvr12El1:
		return "HVSysRegDbgbvr12El1"
	case HVSysRegDbgbvr13El1:
		return "HVSysRegDbgbvr13El1"
	case HVSysRegDbgbvr14El1:
		return "HVSysRegDbgbvr14El1"
	case HVSysRegDbgbvr15El1:
		return "HVSysRegDbgbvr15El1"
	case HVSysRegDbgbvr1El1:
		return "HVSysRegDbgbvr1El1"
	case HVSysRegDbgbvr2El1:
		return "HVSysRegDbgbvr2El1"
	case HVSysRegDbgbvr3El1:
		return "HVSysRegDbgbvr3El1"
	case HVSysRegDbgbvr4El1:
		return "HVSysRegDbgbvr4El1"
	case HVSysRegDbgbvr5El1:
		return "HVSysRegDbgbvr5El1"
	case HVSysRegDbgbvr6El1:
		return "HVSysRegDbgbvr6El1"
	case HVSysRegDbgbvr7El1:
		return "HVSysRegDbgbvr7El1"
	case HVSysRegDbgbvr8El1:
		return "HVSysRegDbgbvr8El1"
	case HVSysRegDbgbvr9El1:
		return "HVSysRegDbgbvr9El1"
	case HVSysRegDbgwcr0El1:
		return "HVSysRegDbgwcr0El1"
	case HVSysRegDbgwcr10El1:
		return "HVSysRegDbgwcr10El1"
	case HVSysRegDbgwcr11El1:
		return "HVSysRegDbgwcr11El1"
	case HVSysRegDbgwcr12El1:
		return "HVSysRegDbgwcr12El1"
	case HVSysRegDbgwcr13El1:
		return "HVSysRegDbgwcr13El1"
	case HVSysRegDbgwcr14El1:
		return "HVSysRegDbgwcr14El1"
	case HVSysRegDbgwcr15El1:
		return "HVSysRegDbgwcr15El1"
	case HVSysRegDbgwcr1El1:
		return "HVSysRegDbgwcr1El1"
	case HVSysRegDbgwcr2El1:
		return "HVSysRegDbgwcr2El1"
	case HVSysRegDbgwcr3El1:
		return "HVSysRegDbgwcr3El1"
	case HVSysRegDbgwcr4El1:
		return "HVSysRegDbgwcr4El1"
	case HVSysRegDbgwcr5El1:
		return "HVSysRegDbgwcr5El1"
	case HVSysRegDbgwcr6El1:
		return "HVSysRegDbgwcr6El1"
	case HVSysRegDbgwcr7El1:
		return "HVSysRegDbgwcr7El1"
	case HVSysRegDbgwcr8El1:
		return "HVSysRegDbgwcr8El1"
	case HVSysRegDbgwcr9El1:
		return "HVSysRegDbgwcr9El1"
	case HVSysRegDbgwvr0El1:
		return "HVSysRegDbgwvr0El1"
	case HVSysRegDbgwvr10El1:
		return "HVSysRegDbgwvr10El1"
	case HVSysRegDbgwvr11El1:
		return "HVSysRegDbgwvr11El1"
	case HVSysRegDbgwvr12El1:
		return "HVSysRegDbgwvr12El1"
	case HVSysRegDbgwvr13El1:
		return "HVSysRegDbgwvr13El1"
	case HVSysRegDbgwvr14El1:
		return "HVSysRegDbgwvr14El1"
	case HVSysRegDbgwvr15El1:
		return "HVSysRegDbgwvr15El1"
	case HVSysRegDbgwvr1El1:
		return "HVSysRegDbgwvr1El1"
	case HVSysRegDbgwvr2El1:
		return "HVSysRegDbgwvr2El1"
	case HVSysRegDbgwvr3El1:
		return "HVSysRegDbgwvr3El1"
	case HVSysRegDbgwvr4El1:
		return "HVSysRegDbgwvr4El1"
	case HVSysRegDbgwvr5El1:
		return "HVSysRegDbgwvr5El1"
	case HVSysRegDbgwvr6El1:
		return "HVSysRegDbgwvr6El1"
	case HVSysRegDbgwvr7El1:
		return "HVSysRegDbgwvr7El1"
	case HVSysRegDbgwvr8El1:
		return "HVSysRegDbgwvr8El1"
	case HVSysRegDbgwvr9El1:
		return "HVSysRegDbgwvr9El1"
	case HVSysRegElrEl1:
		return "HVSysRegElrEl1"
	case HVSysRegElrEl2:
		return "HVSysRegElrEl2"
	case HVSysRegEsrEl1:
		return "HVSysRegEsrEl1"
	case HVSysRegEsrEl2:
		return "HVSysRegEsrEl2"
	case HVSysRegFarEl1:
		return "HVSysRegFarEl1"
	case HVSysRegFarEl2:
		return "HVSysRegFarEl2"
	case HVSysRegHcrEl2:
		return "HVSysRegHcrEl2"
	case HVSysRegHpfarEl2:
		return "HVSysRegHpfarEl2"
	case HVSysRegIDAa64dfr0El1:
		return "HVSysRegIDAa64dfr0El1"
	case HVSysRegIDAa64dfr1El1:
		return "HVSysRegIDAa64dfr1El1"
	case HVSysRegIDAa64isar0El1:
		return "HVSysRegIDAa64isar0El1"
	case HVSysRegIDAa64isar1El1:
		return "HVSysRegIDAa64isar1El1"
	case HVSysRegIDAa64mmfr0El1:
		return "HVSysRegIDAa64mmfr0El1"
	case HVSysRegIDAa64mmfr1El1:
		return "HVSysRegIDAa64mmfr1El1"
	case HVSysRegIDAa64mmfr2El1:
		return "HVSysRegIDAa64mmfr2El1"
	case HVSysRegIDAa64pfr0El1:
		return "HVSysRegIDAa64pfr0El1"
	case HVSysRegIDAa64pfr1El1:
		return "HVSysRegIDAa64pfr1El1"
	case HVSysRegIDAa64smfr0El1:
		return "HVSysRegIDAa64smfr0El1"
	case HVSysRegIDAa64zfr0El1:
		return "HVSysRegIDAa64zfr0El1"
	case HVSysRegMairEl1:
		return "HVSysRegMairEl1"
	case HVSysRegMairEl2:
		return "HVSysRegMairEl2"
	case HVSysRegMdccintEl1:
		return "HVSysRegMdccintEl1"
	case HVSysRegMdcrEl2:
		return "HVSysRegMdcrEl2"
	case HVSysRegMdscrEl1:
		return "HVSysRegMdscrEl1"
	case HVSysRegMidrEl1:
		return "HVSysRegMidrEl1"
	case HVSysRegMpidrEl1:
		return "HVSysRegMpidrEl1"
	case HVSysRegParEl1:
		return "HVSysRegParEl1"
	case HVSysRegSctlrEl1:
		return "HVSysRegSctlrEl1"
	case HVSysRegSctlrEl2:
		return "HVSysRegSctlrEl2"
	case HVSysRegScxtnumEl0:
		return "HVSysRegScxtnumEl0"
	case HVSysRegScxtnumEl1:
		return "HVSysRegScxtnumEl1"
	case HVSysRegSmcrEl1:
		return "HVSysRegSmcrEl1"
	case HVSysRegSmpriEl1:
		return "HVSysRegSmpriEl1"
	case HVSysRegSpEl0:
		return "HVSysRegSpEl0"
	case HVSysRegSpEl1:
		return "HVSysRegSpEl1"
	case HVSysRegSpEl2:
		return "HVSysRegSpEl2"
	case HVSysRegSpsrEl1:
		return "HVSysRegSpsrEl1"
	case HVSysRegSpsrEl2:
		return "HVSysRegSpsrEl2"
	case HVSysRegTcrEl1:
		return "HVSysRegTcrEl1"
	case HVSysRegTcrEl2:
		return "HVSysRegTcrEl2"
	case HVSysRegTpidr2El0:
		return "HVSysRegTpidr2El0"
	case HVSysRegTpidrEl0:
		return "HVSysRegTpidrEl0"
	case HVSysRegTpidrEl1:
		return "HVSysRegTpidrEl1"
	case HVSysRegTpidrEl2:
		return "HVSysRegTpidrEl2"
	case HVSysRegTpidrroEl0:
		return "HVSysRegTpidrroEl0"
	case HVSysRegTtbr0El1:
		return "HVSysRegTtbr0El1"
	case HVSysRegTtbr0El2:
		return "HVSysRegTtbr0El2"
	case HVSysRegTtbr1El1:
		return "HVSysRegTtbr1El1"
	case HVSysRegTtbr1El2:
		return "HVSysRegTtbr1El2"
	case HVSysRegVbarEl1:
		return "HVSysRegVbarEl1"
	case HVSysRegVbarEl2:
		return "HVSysRegVbarEl2"
	case HVSysRegVmpidrEl2:
		return "HVSysRegVmpidrEl2"
	case HVSysRegVpidrEl2:
		return "HVSysRegVpidrEl2"
	case HVSysRegVtcrEl2:
		return "HVSysRegVtcrEl2"
	case HVSysRegVttbrEl2:
		return "HVSysRegVttbrEl2"
	default:
		return fmt.Sprintf("HVSysReg(%d)", e)
	}
}

type HVVmExitinfo uint

const (
	HVVmExitinfoAPICAccessRead HVVmExitinfo = 7
	HVVmExitinfoIOAPICEoi      HVVmExitinfo = 4
	HVVmExitinfoInitAp         HVVmExitinfo = 2
	HVVmExitinfoInjectExcp     HVVmExitinfo = 5
	HVVmExitinfoSmi            HVVmExitinfo = 6
	HVVmExitinfoStartupAp      HVVmExitinfo = 3
	HVVmExitinfoVmx            HVVmExitinfo = 1
)

func (e HVVmExitinfo) String() string {
	switch e {
	case HVVmExitinfoAPICAccessRead:
		return "HVVmExitinfoAPICAccessRead"
	case HVVmExitinfoIOAPICEoi:
		return "HVVmExitinfoIOAPICEoi"
	case HVVmExitinfoInitAp:
		return "HVVmExitinfoInitAp"
	case HVVmExitinfoInjectExcp:
		return "HVVmExitinfoInjectExcp"
	case HVVmExitinfoSmi:
		return "HVVmExitinfoSmi"
	case HVVmExitinfoStartupAp:
		return "HVVmExitinfoStartupAp"
	case HVVmExitinfoVmx:
		return "HVVmExitinfoVmx"
	default:
		return fmt.Sprintf("HVVmExitinfo(%d)", e)
	}
}

type HVVmxCapability uint

const (
	// HVVmxCapBasic: Field ID for basic VMX capabilities.
	HVVmxCapBasic HVVmxCapability = 5
	// HVVmxCapCr0Fixed0: Field ID for CR0 allowed, zero-bits VMX capability.
	HVVmxCapCr0Fixed0 HVVmxCapability = 11
	// HVVmxCapCr0Fixed1: Field ID for CR0 allowed, one-bits VMX capability.
	HVVmxCapCr0Fixed1 HVVmxCapability = 12
	// HVVmxCapCr4Fixed0: Fields ID for CR4 allowed, zero-bits VMX capability.
	HVVmxCapCr4Fixed0 HVVmxCapability = 13
	// HVVmxCapCr4Fixed1: Field ID for CR4 allowed, one-bits VMX capability.
	HVVmxCapCr4Fixed1 HVVmxCapability = 14
	// HVVmxCapEntry: Field ID for VM entry capabilities.
	HVVmxCapEntry HVVmxCapability = 3
	// HVVmxCapEptVpidCap: Field ID for EPT/VPID VMX capabilities.
	HVVmxCapEptVpidCap HVVmxCapability = 16
	// HVVmxCapExit: Field ID for VM exit capabilities.
	HVVmxCapExit HVVmxCapability = 4
	// HVVmxCapMisc: Field ID for miscellaneous VMX capabilities.
	HVVmxCapMisc HVVmxCapability = 10
	// HVVmxCapPinbased: Field ID for pin-based capabilities.
	HVVmxCapPinbased HVVmxCapability = 0
	// HVVmxCapPreemptionTimer: Field ID for preemption timer frequency.
	HVVmxCapPreemptionTimer HVVmxCapability = 32
	// HVVmxCapProcbased: Field ID for primary proc-based capabilities.
	HVVmxCapProcbased HVVmxCapability = 1
	// HVVmxCapProcbased2: Field ID for secondary proc-based capabilities.
	HVVmxCapProcbased2 HVVmxCapability = 2
	// HVVmxCapTrueEntry: Field ID for hardware VM-entry VMX capabilities.
	HVVmxCapTrueEntry HVVmxCapability = 8
	// HVVmxCapTrueExit: Field ID for hardware VM-exit VMX capabilities.
	HVVmxCapTrueExit HVVmxCapability = 9
	// HVVmxCapTruePinbased: Field ID for hardware pin-based VMX capabilities.
	HVVmxCapTruePinbased HVVmxCapability = 6
	// HVVmxCapTrueProcbased: Field ID for primary process-based VMX capabilities.
	HVVmxCapTrueProcbased HVVmxCapability = 7
	// HVVmxCapVmcsEnum: Field ID for VMCS enumeration capability.
	HVVmxCapVmcsEnum HVVmxCapability = 15
)

func (e HVVmxCapability) String() string {
	switch e {
	case HVVmxCapBasic:
		return "HVVmxCapBasic"
	case HVVmxCapCr0Fixed0:
		return "HVVmxCapCr0Fixed0"
	case HVVmxCapCr0Fixed1:
		return "HVVmxCapCr0Fixed1"
	case HVVmxCapCr4Fixed0:
		return "HVVmxCapCr4Fixed0"
	case HVVmxCapCr4Fixed1:
		return "HVVmxCapCr4Fixed1"
	case HVVmxCapEntry:
		return "HVVmxCapEntry"
	case HVVmxCapEptVpidCap:
		return "HVVmxCapEptVpidCap"
	case HVVmxCapExit:
		return "HVVmxCapExit"
	case HVVmxCapMisc:
		return "HVVmxCapMisc"
	case HVVmxCapPinbased:
		return "HVVmxCapPinbased"
	case HVVmxCapPreemptionTimer:
		return "HVVmxCapPreemptionTimer"
	case HVVmxCapProcbased:
		return "HVVmxCapProcbased"
	case HVVmxCapProcbased2:
		return "HVVmxCapProcbased2"
	case HVVmxCapTrueEntry:
		return "HVVmxCapTrueEntry"
	case HVVmxCapTrueExit:
		return "HVVmxCapTrueExit"
	case HVVmxCapTruePinbased:
		return "HVVmxCapTruePinbased"
	case HVVmxCapTrueProcbased:
		return "HVVmxCapTrueProcbased"
	case HVVmxCapVmcsEnum:
		return "HVVmxCapVmcsEnum"
	default:
		return fmt.Sprintf("HVVmxCapability(%d)", e)
	}
}

type HVX86Reg uint

const (
	// HVX86Cr0: The value that identifies the x86 control-register CR0.
	HVX86Cr0 HVX86Reg = 36
	// HVX86Cr1: The value that identifies the x86 control-register CR1.
	HVX86Cr1 HVX86Reg = 37
	// HVX86Cr2: The value that identifies the x86 control-register CR2.
	HVX86Cr2 HVX86Reg = 38
	// HVX86Cr3: The value that identifies the x86 control-register CR3.
	HVX86Cr3 HVX86Reg = 39
	// HVX86Cr4: The value that identifies the x86 control-register CR4.
	HVX86Cr4 HVX86Reg = 40
	// HVX86Cs: The value that identifies the x86 code-segment register.
	HVX86Cs HVX86Reg = 18
	// HVX86Dr0: The value that identifies the x86 debug-register DR0.
	HVX86Dr0 HVX86Reg = 41
	// HVX86Dr1: The value that identifies the x86 debug-register DR1.
	HVX86Dr1 HVX86Reg = 42
	// HVX86Dr2: The value that identifies the x86 debug-register DR2.
	HVX86Dr2 HVX86Reg = 43
	// HVX86Dr3: The value that identifies the x86 debug-register DR3.
	HVX86Dr3 HVX86Reg = 44
	// HVX86Dr4: The value that identifies the x86 debug-register DR4.
	HVX86Dr4 HVX86Reg = 45
	// HVX86Dr5: The value that identifies the x86 debug-register DR5.
	HVX86Dr5 HVX86Reg = 46
	// HVX86Dr6: The value that identifies the x86 debug-register DR6.
	HVX86Dr6 HVX86Reg = 47
	// HVX86Dr7: The value that identifies the x86 debug-register DR7.
	HVX86Dr7 HVX86Reg = 48
	// HVX86Ds: The value that identifies the x86 data-segment register.
	HVX86Ds HVX86Reg = 20
	// HVX86Es: The value that identifies the x86 segment register ES.
	HVX86Es HVX86Reg = 21
	// HVX86Fs: The value that identifies the x86 segment register FS.
	HVX86Fs HVX86Reg = 22
	// HVX86GdtBase: The value that identifies the x86 global descriptor, table-base register.
	HVX86GdtBase HVX86Reg = 26
	// HVX86GdtLimit: The value that identifies the x86 global descriptor, table-limit register.
	HVX86GdtLimit HVX86Reg = 27
	// HVX86Gs: The value that identifies the x86 segment register GS.
	HVX86Gs HVX86Reg = 23
	// HVX86IdtBase: The value that identifies the x86 interrupt descriptor, table-base register.
	HVX86IdtBase HVX86Reg = 24
	// HVX86IdtLimit: The value that identifies the x86 interrupt descriptor, table-base register.
	HVX86IdtLimit HVX86Reg = 25
	// HVX86LdtAr: The value that identifies the x86 local descriptor table, access-rights register.
	HVX86LdtAr HVX86Reg = 31
	// HVX86LdtBase: The value that identifies the x86 local descriptor, table-base register.
	HVX86LdtBase HVX86Reg = 29
	// HVX86LdtLimit: The value that identifies the x86 local descriptor, table-limit register.
	HVX86LdtLimit HVX86Reg = 30
	// HVX86Ldtr: The value that identifies the x86 local descriptor, table register.
	HVX86Ldtr HVX86Reg = 28
	// HVX86R10: The value that identifies the x86 general-purpose register R10.
	HVX86R10 HVX86Reg = 12
	// HVX86R11: The value that identifies the x86 general-purpose register R11.
	HVX86R11 HVX86Reg = 13
	// HVX86R12: The value that identifies the x86 general-purpose register R12.
	HVX86R12 HVX86Reg = 14
	// HVX86R13: The value that identifies the x86 general-purpose register R13.
	HVX86R13 HVX86Reg = 15
	// HVX86R14: The value that identifies the x86 general-purpose register R14.
	HVX86R14 HVX86Reg = 16
	// HVX86R15: The value that identifies the x86 general-purpose register R15.
	HVX86R15 HVX86Reg = 17
	// HVX86R8: The value that identifies the x86 general-purpose register R8.
	HVX86R8 HVX86Reg = 10
	// HVX86R9: The value that identifies the x86 general-purpose register R9.
	HVX86R9 HVX86Reg = 11
	// HVX86Rax: The value that identifies the x86 accumulator register.
	HVX86Rax HVX86Reg = 2
	// HVX86Rbp: The value that identifies the x86 base pointer register.
	HVX86Rbp HVX86Reg = 9
	// HVX86Rbx: The value that identifies the x86 base register.
	HVX86Rbx HVX86Reg = 5
	// HVX86Rcx: The value that identifies the x86 counter register.
	HVX86Rcx HVX86Reg = 3
	// HVX86Rdi: The value that identifies the x86 destination index register.
	HVX86Rdi HVX86Reg = 7
	// HVX86Rdx: The value that identifies the x86 data register.
	HVX86Rdx HVX86Reg = 4
	// HVX86RegistersMax: The value that identifies the maximum value of x86 register constants.
	HVX86RegistersMax HVX86Reg = 51
	// HVX86Rflags: The value that identifies the x86 status register.
	HVX86Rflags HVX86Reg = 1
	// HVX86Rip: The value that identifies the x86 instruction pointer register.
	HVX86Rip HVX86Reg = 0
	// HVX86Rsi: The value that identifies the x86 source index register.
	HVX86Rsi HVX86Reg = 6
	// HVX86Rsp: The value that identifies the x86 stack pointer register.
	HVX86Rsp HVX86Reg = 8
	// HVX86Ss: The value that identifies the x86 stack-segment register.
	HVX86Ss HVX86Reg = 19
	// HVX86Tpr: The value that identifies the x86 task-priority register.
	HVX86Tpr HVX86Reg = 49
	// HVX86Tr: The value that identifies the x86 task register.
	HVX86Tr HVX86Reg = 32
	// HVX86TssAr: The value that identifies the x86 task-state, segment-access, rights register.
	HVX86TssAr HVX86Reg = 35
	// HVX86TssBase: The value that identifies the x86 task-state, segment-base register.
	HVX86TssBase HVX86Reg = 33
	// HVX86TssLimit: The value that identifies the x86 task state segment limit register.
	HVX86TssLimit HVX86Reg = 34
	// HVX86Xcr0: The value that identifies the x86 extended-control register.
	HVX86Xcr0 HVX86Reg = 50
)

func (e HVX86Reg) String() string {
	switch e {
	case HVX86Cr0:
		return "HVX86Cr0"
	case HVX86Cr1:
		return "HVX86Cr1"
	case HVX86Cr2:
		return "HVX86Cr2"
	case HVX86Cr3:
		return "HVX86Cr3"
	case HVX86Cr4:
		return "HVX86Cr4"
	case HVX86Cs:
		return "HVX86Cs"
	case HVX86Dr0:
		return "HVX86Dr0"
	case HVX86Dr1:
		return "HVX86Dr1"
	case HVX86Dr2:
		return "HVX86Dr2"
	case HVX86Dr3:
		return "HVX86Dr3"
	case HVX86Dr4:
		return "HVX86Dr4"
	case HVX86Dr5:
		return "HVX86Dr5"
	case HVX86Dr6:
		return "HVX86Dr6"
	case HVX86Dr7:
		return "HVX86Dr7"
	case HVX86Ds:
		return "HVX86Ds"
	case HVX86Es:
		return "HVX86Es"
	case HVX86Fs:
		return "HVX86Fs"
	case HVX86GdtBase:
		return "HVX86GdtBase"
	case HVX86GdtLimit:
		return "HVX86GdtLimit"
	case HVX86Gs:
		return "HVX86Gs"
	case HVX86IdtBase:
		return "HVX86IdtBase"
	case HVX86IdtLimit:
		return "HVX86IdtLimit"
	case HVX86LdtAr:
		return "HVX86LdtAr"
	case HVX86LdtBase:
		return "HVX86LdtBase"
	case HVX86LdtLimit:
		return "HVX86LdtLimit"
	case HVX86Ldtr:
		return "HVX86Ldtr"
	case HVX86R10:
		return "HVX86R10"
	case HVX86R11:
		return "HVX86R11"
	case HVX86R12:
		return "HVX86R12"
	case HVX86R13:
		return "HVX86R13"
	case HVX86R14:
		return "HVX86R14"
	case HVX86R15:
		return "HVX86R15"
	case HVX86R8:
		return "HVX86R8"
	case HVX86R9:
		return "HVX86R9"
	case HVX86Rax:
		return "HVX86Rax"
	case HVX86Rbp:
		return "HVX86Rbp"
	case HVX86Rbx:
		return "HVX86Rbx"
	case HVX86Rcx:
		return "HVX86Rcx"
	case HVX86Rdi:
		return "HVX86Rdi"
	case HVX86Rdx:
		return "HVX86Rdx"
	case HVX86RegistersMax:
		return "HVX86RegistersMax"
	case HVX86Rflags:
		return "HVX86Rflags"
	case HVX86Rip:
		return "HVX86Rip"
	case HVX86Rsi:
		return "HVX86Rsi"
	case HVX86Rsp:
		return "HVX86Rsp"
	case HVX86Ss:
		return "HVX86Ss"
	case HVX86Tpr:
		return "HVX86Tpr"
	case HVX86Tr:
		return "HVX86Tr"
	case HVX86TssAr:
		return "HVX86TssAr"
	case HVX86TssBase:
		return "HVX86TssBase"
	case HVX86TssLimit:
		return "HVX86TssLimit"
	case HVX86Xcr0:
		return "HVX86Xcr0"
	default:
		return fmt.Sprintf("HVX86Reg(%d)", e)
	}
}

type HVAllocate uint

const (
	HVAllocateDefault HVAllocate = 0
)

func (e HVAllocate) String() string {
	switch e {
	case HVAllocateDefault:
		return "HVAllocateDefault"
	default:
		return fmt.Sprintf("HVAllocate(%d)", e)
	}
}

type HVCap uint

const (
	// HVCapAddrspacemax: A value that indicates the maximum number of available address spaces.
	HVCapAddrspacemax HVCap = 1
	// HVCapVcpumax: A value that indicates the maximum number of available vCPUs.
	HVCapVcpumax HVCap = 0
)

func (e HVCap) String() string {
	switch e {
	case HVCapAddrspacemax:
		return "HVCapAddrspacemax"
	case HVCapVcpumax:
		return "HVCapVcpumax"
	default:
		return fmt.Sprintf("HVCap(%d)", e)
	}
}

type HVDeadline uint

const (
	// HVDeadlineForever: The value that indicates a vCPU deadline that never expires.
	HVDeadlineForever HVDeadline = 0
)

func (e HVDeadline) String() string {
	switch e {
	case HVDeadlineForever:
		return "HVDeadlineForever"
	default:
		return fmt.Sprintf("HVDeadline(%d)", e)
	}
}

type HVIon uint

const (
	// HVIonAnySize: The value that represents a request for notifications of an I/O result of any size.
	HVIonAnySize HVIon = 4
	// HVIonAnyValue: The value that represents a request for notifications of an I/O result that contains any value.
	HVIonAnyValue HVIon = 2
	// HVIonExitFull: The value that represents a request for notifications if the I/O queue is full.
	HVIonExitFull HVIon = 8
	// HVIonNone: The value that represents a request for no notifications.
	HVIonNone HVIon = 0
)

func (e HVIon) String() string {
	switch e {
	case HVIonAnySize:
		return "HVIonAnySize"
	case HVIonAnyValue:
		return "HVIonAnyValue"
	case HVIonExitFull:
		return "HVIonExitFull"
	case HVIonNone:
		return "HVIonNone"
	default:
		return fmt.Sprintf("HVIon(%d)", e)
	}
}

type HVMemory uint

const (
	// HVMemoryExec: The value that represents the memory-execute permission.
	HVMemoryExec         HVMemory = 4
	HVMemoryMaxprot      HVMemory = 16
	HVMemoryMaxprotExec  HVMemory = 128
	HVMemoryMaxprotRead  HVMemory = 32
	HVMemoryMaxprotUexec HVMemory = 256
	HVMemoryMaxprotWrite HVMemory = 64
	// HVMemoryRead: The value that represents the memory-read permission.
	HVMemoryRead  HVMemory = 1
	HVMemoryUexec HVMemory = 8
	// HVMemoryWrite: The value that represents the memory-write permission.
	HVMemoryWrite HVMemory = 2
)

func (e HVMemory) String() string {
	switch e {
	case HVMemoryExec:
		return "HVMemoryExec"
	case HVMemoryMaxprot:
		return "HVMemoryMaxprot"
	case HVMemoryMaxprotExec:
		return "HVMemoryMaxprotExec"
	case HVMemoryMaxprotRead:
		return "HVMemoryMaxprotRead"
	case HVMemoryMaxprotUexec:
		return "HVMemoryMaxprotUexec"
	case HVMemoryMaxprotWrite:
		return "HVMemoryMaxprotWrite"
	case HVMemoryRead:
		return "HVMemoryRead"
	case HVMemoryUexec:
		return "HVMemoryUexec"
	case HVMemoryWrite:
		return "HVMemoryWrite"
	default:
		return fmt.Sprintf("HVMemory(%d)", e)
	}
}

type HVMsr uint

const (
	// HVMsrIa32APmc0: The value that represents support for address performance-counter register 0.
	HVMsrIa32APmc0 HVMsr = 0x4c1
	// HVMsrIa32APmc7: The value that represents support for address performance-counter register 7.
	HVMsrIa32APmc7 HVMsr = 1217
	// HVMsrIa32ArchCapabilities: The value that represents the Model-Specific Register (MSR) that you use to enumerate processor capabilities.
	HVMsrIa32ArchCapabilities HVMsr = 0x10a
	// HVMsrIa32Cstar: The value that represents the address of IA-32e Mode System Call Target Address.
	HVMsrIa32Cstar HVMsr = 0xc0000083
	// HVMsrIa32Debugctl: The value that represents the address of the Debug Control Register.
	HVMsrIa32Debugctl HVMsr = 0x1d9
	// HVMsrIa32Efer: The value that represents the address of the Entended Feature Enable Register (EFER).
	HVMsrIa32Efer HVMsr = 0xc0000080
	// HVMsrIa32FixedCtr0: The value that represents the address of Fixed-Function Performance Counter Register 0.
	HVMsrIa32FixedCtr0 HVMsr = 0x309
	// HVMsrIa32FixedCtr1: The value that represents the address of Fixed-Function Performance Counter Register 1.
	HVMsrIa32FixedCtr1 HVMsr = 0x30a
	// HVMsrIa32FixedCtr2: The value that represents the address of Fixed-Function Performance Counter Register 2.
	HVMsrIa32FixedCtr2 HVMsr = 0x30b
	// HVMsrIa32FixedCtr3: The value that represents the address of Fixed-Function Performance Counter Register 3.
	HVMsrIa32FixedCtr3 HVMsr = 0x30c
	// HVMsrIa32FixedCtrCtrl: The value that represents the address of the Fixed-Function Counter Control Register.
	HVMsrIa32FixedCtrCtrl HVMsr = 0x38d
	// HVMsrIa32FlushCmd: The value that represents the address of the Flush Command Register.
	HVMsrIa32FlushCmd HVMsr = 0x10b
	// HVMsrIa32Fmask: The value that represents the address of the System Call Flag Mask (FMASK) Register.
	HVMsrIa32Fmask HVMsr = 0xc0000084
	// HVMsrIa32FsBase: The value that represents the address of the map for the base address of the FS segment register.
	HVMsrIa32FsBase HVMsr = 0xc0000100
	// HVMsrIa32GsBase: The value that represents the address of the map for the base address of the GS segment register.
	HVMsrIa32GsBase HVMsr = 0xc0000101
	// HVMsrIa32KernelGsBase: The value that represents the address swap target for the base address of the GS segment register.
	HVMsrIa32KernelGsBase HVMsr = 0xc0000102
	// HVMsrIa32Lstar: The value that represents the address of the IA-32e Mode System Call Target Address.
	HVMsrIa32Lstar HVMsr = 0xc0000082
	// HVMsrIa32PerfGlobalCtrl: The value that represents the address of the Global Performance Counter Control Register.
	HVMsrIa32PerfGlobalCtrl HVMsr = 0x38f
	// HVMsrIa32PerfGlobalInuse: The value that represents the address of the register that indicates whether the core performance monitor interface is in use.
	HVMsrIa32PerfGlobalInuse HVMsr = 0x392
	// HVMsrIa32PerfGlobalStatus: The value that represents the address of the Global Performance Status Register.
	HVMsrIa32PerfGlobalStatus HVMsr = 0x38e
	// HVMsrIa32PerfGlobalStatusReset: The value that represents the address of the Global Performance Counter Overflow Reset Control Register.
	HVMsrIa32PerfGlobalStatusReset HVMsr = 0x390
	// HVMsrIa32PerfGlobalStatusSet: The value that represents the address of the Global Performance Counter Overflow Set Control Register.
	HVMsrIa32PerfGlobalStatusSet HVMsr = 0x391
	// HVMsrIa32Perfevntsel0: The value that represents the address of Performance Event Select Counter 0.
	HVMsrIa32Perfevntsel0 HVMsr = 0x186
	// HVMsrIa32Perfevntsel7: The value that represents the address of Performance Event Select Counter 7.
	HVMsrIa32Perfevntsel7 HVMsr = 390
	// HVMsrIa32Pmc0: The value that represents the address of Performance Counter Register 0.
	HVMsrIa32Pmc0 HVMsr = 0xc1
	// HVMsrIa32Pmc7: The value that represents the address of Performance Counter Register 7.
	HVMsrIa32Pmc7 HVMsr = 193
	// HVMsrIa32PredCmd: The value that represents the address of the Prediction Command Register.
	HVMsrIa32PredCmd HVMsr = 0x49
	// HVMsrIa32SpecCtrl: The value that represents the address of Speculation Control Register.
	HVMsrIa32SpecCtrl HVMsr = 0x48
	// HVMsrIa32Star: The value that represents the address of the System Call Target Address Register.
	HVMsrIa32Star HVMsr = 0xc0000081
	// HVMsrIa32SysenterCs: The value that represents the address of the CS Register target for Current Privilege Level (CPL) 0 code.
	HVMsrIa32SysenterCs HVMsr = 0x174
	// HVMsrIa32SysenterEip: The value that represents the address of the Extended Instruction Pointer (EIP) Register target for Current Privilege Level (CPL) 0 code.
	HVMsrIa32SysenterEip HVMsr = 0x176
	// HVMsrIa32SysenterEsp: The value that represents the address of the Extended Stack Pointer (ESP) Register target for Current Privilege Level (CPL) 0 code.
	HVMsrIa32SysenterEsp HVMsr = 0x175
	// HVMsrIa32Tsc: The value that represents the address of the Time-Stamp Counter Register.
	HVMsrIa32Tsc HVMsr = 0x10
	// HVMsrIa32TscAux: The value that represents the address of the Auxiliary Time-Stamp Counter Register.
	HVMsrIa32TscAux HVMsr = 0xc0000103
	// HVMsrIa32Xss: The value that represents the address of the Extended Supervisors State Mask (XSS) Register.
	HVMsrIa32Xss HVMsr = 0xda0
	// HVMsrLastbranch0FromIP: The value that represents the address of the Last Branch Record 0 from Instruction Pointer (IP) register.
	HVMsrLastbranch0FromIP HVMsr = 0x680
	// HVMsrLastbranch0ToIP: The value that represents the address of the Last Branch Record 0 to Instruction Pointer (IP) register.
	HVMsrLastbranch0ToIP HVMsr = 0x6c0
	// HVMsrLastbranch31FromIP: The value that represents the address of the Last Branch Record 31 from Instruction Pointer (IP) register.
	HVMsrLastbranch31FromIP HVMsr = 1664
	// HVMsrLastbranch31ToIP: The value that represents the address of the Last Branch Record 31 to Instruction Pointer (IP) register.
	HVMsrLastbranch31ToIP HVMsr = 1728
	// HVMsrLastbranchInfo0: The value that represents the address of the Last Branch Record 0 additional information register.
	HVMsrLastbranchInfo0 HVMsr = 0xdc0
	// HVMsrLastbranchInfo31: The value that represents the address of the Last Branch Record 31 additional information register.
	HVMsrLastbranchInfo31 HVMsr = 3520
	// HVMsrLastbranchTos: The value that represents the address of the Last Branch Record Top of Stack (TOS) Register.
	HVMsrLastbranchTos HVMsr = 0x1c9
	// HVMsrLastintFromIP: The value that represents the address of the Last Interrupt from Instruction Pointer (IP) Register.
	HVMsrLastintFromIP HVMsr = 0x1dd
	// HVMsrLastintToIP: The value that represents the address of the Last Interrupt to Instruction Pointer (IP) Register.
	HVMsrLastintToIP HVMsr = 0x1de
	// HVMsrLbrSelect: The value that represents the address of the Last Branch Record Filtering Select Register.
	HVMsrLbrSelect HVMsr = 0x1c8
	// HVMsrNone: The Model-Specific Register (MSR) no-access permission.
	HVMsrNone HVMsr = 0
	// HVMsrPerfMetrics: The value that represents the address of the Performance Metrics Register.
	HVMsrPerfMetrics HVMsr = 0x329
	// HVMsrRead: The Model-Specific Register (MSR) read permission.
	HVMsrRead HVMsr = 1
	// HVMsrWrite: The Model-Specific Register (MSR) write permission.
	HVMsrWrite HVMsr = 2
)

func (e HVMsr) String() string {
	switch e {
	case HVMsrIa32APmc0:
		return "HVMsrIa32APmc0"
	case HVMsrIa32ArchCapabilities:
		return "HVMsrIa32ArchCapabilities"
	case HVMsrIa32Cstar:
		return "HVMsrIa32Cstar"
	case HVMsrIa32Debugctl:
		return "HVMsrIa32Debugctl"
	case HVMsrIa32Efer:
		return "HVMsrIa32Efer"
	case HVMsrIa32FixedCtr0:
		return "HVMsrIa32FixedCtr0"
	case HVMsrIa32FixedCtr1:
		return "HVMsrIa32FixedCtr1"
	case HVMsrIa32FixedCtr2:
		return "HVMsrIa32FixedCtr2"
	case HVMsrIa32FixedCtr3:
		return "HVMsrIa32FixedCtr3"
	case HVMsrIa32FixedCtrCtrl:
		return "HVMsrIa32FixedCtrCtrl"
	case HVMsrIa32FlushCmd:
		return "HVMsrIa32FlushCmd"
	case HVMsrIa32Fmask:
		return "HVMsrIa32Fmask"
	case HVMsrIa32FsBase:
		return "HVMsrIa32FsBase"
	case HVMsrIa32GsBase:
		return "HVMsrIa32GsBase"
	case HVMsrIa32KernelGsBase:
		return "HVMsrIa32KernelGsBase"
	case HVMsrIa32Lstar:
		return "HVMsrIa32Lstar"
	case HVMsrIa32PerfGlobalCtrl:
		return "HVMsrIa32PerfGlobalCtrl"
	case HVMsrIa32PerfGlobalInuse:
		return "HVMsrIa32PerfGlobalInuse"
	case HVMsrIa32PerfGlobalStatus:
		return "HVMsrIa32PerfGlobalStatus"
	case HVMsrIa32PerfGlobalStatusReset:
		return "HVMsrIa32PerfGlobalStatusReset"
	case HVMsrIa32PerfGlobalStatusSet:
		return "HVMsrIa32PerfGlobalStatusSet"
	case HVMsrIa32Perfevntsel0:
		return "HVMsrIa32Perfevntsel0"
	case HVMsrIa32Pmc0:
		return "HVMsrIa32Pmc0"
	case HVMsrIa32PredCmd:
		return "HVMsrIa32PredCmd"
	case HVMsrIa32SpecCtrl:
		return "HVMsrIa32SpecCtrl"
	case HVMsrIa32Star:
		return "HVMsrIa32Star"
	case HVMsrIa32SysenterCs:
		return "HVMsrIa32SysenterCs"
	case HVMsrIa32SysenterEip:
		return "HVMsrIa32SysenterEip"
	case HVMsrIa32SysenterEsp:
		return "HVMsrIa32SysenterEsp"
	case HVMsrIa32Tsc:
		return "HVMsrIa32Tsc"
	case HVMsrIa32TscAux:
		return "HVMsrIa32TscAux"
	case HVMsrIa32Xss:
		return "HVMsrIa32Xss"
	case HVMsrLastbranch0FromIP:
		return "HVMsrLastbranch0FromIP"
	case HVMsrLastbranch0ToIP:
		return "HVMsrLastbranch0ToIP"
	case HVMsrLastbranchInfo0:
		return "HVMsrLastbranchInfo0"
	case HVMsrLastbranchTos:
		return "HVMsrLastbranchTos"
	case HVMsrLastintFromIP:
		return "HVMsrLastintFromIP"
	case HVMsrLastintToIP:
		return "HVMsrLastintToIP"
	case HVMsrLbrSelect:
		return "HVMsrLbrSelect"
	case HVMsrNone:
		return "HVMsrNone"
	case HVMsrPerfMetrics:
		return "HVMsrPerfMetrics"
	case HVMsrRead:
		return "HVMsrRead"
	case HVMsrWrite:
		return "HVMsrWrite"
	default:
		return fmt.Sprintf("HVMsr(%d)", e)
	}
}

type HVShadowVmcs uint

const (
	// HVShadowVmcsNone: The value that indicates no access to the shadow VMCS fields.
	HVShadowVmcsNone HVShadowVmcs = 0
	// HVShadowVmcsRead: The value that indicates read access to the shadow VMCS fields.
	HVShadowVmcsRead HVShadowVmcs = 1
	// HVShadowVmcsWrite: The value that indicates read access to the write access shadow VMCS fields.
	HVShadowVmcsWrite HVShadowVmcs = 2
)

func (e HVShadowVmcs) String() string {
	switch e {
	case HVShadowVmcsNone:
		return "HVShadowVmcsNone"
	case HVShadowVmcsRead:
		return "HVShadowVmcsRead"
	case HVShadowVmcsWrite:
		return "HVShadowVmcsWrite"
	default:
		return fmt.Sprintf("HVShadowVmcs(%d)", e)
	}
}

type HVVCPU uint

const (
	// HVVCPUAccelRdpmc: Instructs the kernel, when set, to handle RDPMC VM exits directly rather than passing them to user space.
	HVVCPUAccelRdpmc HVVCPU = 1
	// HVVCPUDefault: The default vCPU creation behavior.
	HVVCPUDefault HVVCPU = 0
	// HVVCPUTscRelative: The value that represents the relative offset the system should add to the hypervisor TSC clock.
	HVVCPUTscRelative HVVCPU = 2
)

func (e HVVCPU) String() string {
	switch e {
	case HVVCPUAccelRdpmc:
		return "HVVCPUAccelRdpmc"
	case HVVCPUDefault:
		return "HVVCPUDefault"
	case HVVCPUTscRelative:
		return "HVVCPUTscRelative"
	default:
		return fmt.Sprintf("HVVCPU(%d)", e)
	}
}

type HVVm uint

const (
	HVVmAccelAPIC HVVm = 1024
	// HVVmDefault: The default VM creation behavior.
	HVVmDefault            HVVm = 0
	HVVmMitigationAEnable  HVVm = 2
	HVVmMitigationBEnable  HVVm = 4
	HVVmMitigationCEnable  HVVm = 8
	HVVmMitigationDEnable  HVVm = 16
	HVVmMitigationEEnable  HVVm = 64
	HVVmSpecifyMitigations HVVm = 1
)

func (e HVVm) String() string {
	switch e {
	case HVVmAccelAPIC:
		return "HVVmAccelAPIC"
	case HVVmDefault:
		return "HVVmDefault"
	case HVVmMitigationAEnable:
		return "HVVmMitigationAEnable"
	case HVVmMitigationBEnable:
		return "HVVmMitigationBEnable"
	case HVVmMitigationCEnable:
		return "HVVmMitigationCEnable"
	case HVVmMitigationDEnable:
		return "HVVmMitigationDEnable"
	case HVVmMitigationEEnable:
		return "HVVmMitigationEEnable"
	case HVVmSpecifyMitigations:
		return "HVVmSpecifyMitigations"
	default:
		return fmt.Sprintf("HVVm(%d)", e)
	}
}

type HVVmSpace uint

const (
	// HVVmSpaceDefault: The value that represents the default VM address space.
	HVVmSpaceDefault HVVmSpace = 0
)

func (e HVVmSpace) String() string {
	switch e {
	case HVVmSpaceDefault:
		return "HVVmSpaceDefault"
	default:
		return fmt.Sprintf("HVVmSpace(%d)", e)
	}
}

type HVVmx uint

const (
	// HVVmxInfoMsrIa32ArchCapabilities: The value of the IA32 architecture capabilities model specific register.
	HVVmxInfoMsrIa32ArchCapabilities HVVmx = 0
	// HVVmxInfoMsrIa32PerfCapabilities: The value of the IA32 performance capabilities model specific register.
	HVVmxInfoMsrIa32PerfCapabilities HVVmx = 1
	// HVVmxNeedMsrIa32SpecCtrl: The bitmask of the required fields of the IA32 Speculation Control model specific register.
	HVVmxNeedMsrIa32SpecCtrl HVVmx = 8
	// HVVmxValidMsrIa32Debugctl: The bitmask of the IA32 Debug-Control model specific register.
	HVVmxValidMsrIa32Debugctl HVVmx = 6
	// HVVmxValidMsrIa32FixedCtrCtrl: The bitmask fo the supported fields of the Fixed-Function-Counter Control Register.
	HVVmxValidMsrIa32FixedCtrCtrl HVVmx = 3
	// HVVmxValidMsrIa32PerfGlobalCtrl: The bitmask of the supported fields of the IA32 Global-Counter Control Facility Register.
	HVVmxValidMsrIa32PerfGlobalCtrl HVVmx = 4
	// HVVmxValidMsrIa32PerfGlobalStatus: The bitmast of the supported fields of the Global-Counter-Control Status model specific register.
	HVVmxValidMsrIa32PerfGlobalStatus HVVmx = 5
	// HVVmxValidMsrIa32Perfevntsel: The bitmask of the supported fields of the IA32 Performance-Event Selection Mode model specific register.
	HVVmxValidMsrIa32Perfevntsel HVVmx = 2
	// HVVmxValidMsrIa32SpecCtrl: The bitmask of the suppported fields of the Speculation Control model specific register.
	HVVmxValidMsrIa32SpecCtrl HVVmx = 7
)

func (e HVVmx) String() string {
	switch e {
	case HVVmxInfoMsrIa32ArchCapabilities:
		return "HVVmxInfoMsrIa32ArchCapabilities"
	case HVVmxInfoMsrIa32PerfCapabilities:
		return "HVVmxInfoMsrIa32PerfCapabilities"
	case HVVmxNeedMsrIa32SpecCtrl:
		return "HVVmxNeedMsrIa32SpecCtrl"
	case HVVmxValidMsrIa32Debugctl:
		return "HVVmxValidMsrIa32Debugctl"
	case HVVmxValidMsrIa32FixedCtrCtrl:
		return "HVVmxValidMsrIa32FixedCtrCtrl"
	case HVVmxValidMsrIa32PerfGlobalCtrl:
		return "HVVmxValidMsrIa32PerfGlobalCtrl"
	case HVVmxValidMsrIa32PerfGlobalStatus:
		return "HVVmxValidMsrIa32PerfGlobalStatus"
	case HVVmxValidMsrIa32Perfevntsel:
		return "HVVmxValidMsrIa32Perfevntsel"
	case HVVmxValidMsrIa32SpecCtrl:
		return "HVVmxValidMsrIa32SpecCtrl"
	default:
		return fmt.Sprintf("HVVmx(%d)", e)
	}
}

type IrqInfo uint

const (
	// IrqInfoErrorValid: The value that indicates the error associated with the interrupt is valid and is readable from the VMCS.
	IrqInfoErrorValid IrqInfo = 2048
	// IrqInfoExtIrq: The value that represents an external interrupt.
	IrqInfoExtIrq IrqInfo = 0
	// IrqInfoHardExc: The value that represents a hardware exception.
	IrqInfoHardExc IrqInfo = 768
	// IrqInfoNmi: The value that represents a non-maskable-interrupt.
	IrqInfoNmi IrqInfo = 512
	// IrqInfoPrivSoftExc: The value that represents a privileged software exception.
	IrqInfoPrivSoftExc IrqInfo = 1280
	// IrqInfoSoftExc: The value that represents a software exception interrupt.
	IrqInfoSoftExc IrqInfo = 1536
	// IrqInfoSoftIrq: The value that represents a software interrupt.
	IrqInfoSoftIrq IrqInfo = 1024
	// IrqInfoTypeMask: The value that represents the interrupt mask.
	IrqInfoTypeMask IrqInfo = 1792
	// IrqInfoValid: The value that represents the interrupt is valid.
	IrqInfoValid      IrqInfo = 2147483648
	IrqInfoVectorMask IrqInfo = 255
)

func (e IrqInfo) String() string {
	switch e {
	case IrqInfoErrorValid:
		return "IrqInfoErrorValid"
	case IrqInfoExtIrq:
		return "IrqInfoExtIrq"
	case IrqInfoHardExc:
		return "IrqInfoHardExc"
	case IrqInfoNmi:
		return "IrqInfoNmi"
	case IrqInfoPrivSoftExc:
		return "IrqInfoPrivSoftExc"
	case IrqInfoSoftExc:
		return "IrqInfoSoftExc"
	case IrqInfoSoftIrq:
		return "IrqInfoSoftIrq"
	case IrqInfoTypeMask:
		return "IrqInfoTypeMask"
	case IrqInfoValid:
		return "IrqInfoValid"
	case IrqInfoVectorMask:
		return "IrqInfoVectorMask"
	default:
		return fmt.Sprintf("IrqInfo(%d)", e)
	}
}

type KhvIon uint

const (
	KHVIonAnySize  KhvIon = 4
	KHVIonAnyValue KhvIon = 2
	KHVIonExitFull KhvIon = 8
	KHVIonNone     KhvIon = 0
)

func (e KhvIon) String() string {
	switch e {
	case KHVIonAnySize:
		return "KHVIonAnySize"
	case KHVIonAnyValue:
		return "KHVIonAnyValue"
	case KHVIonExitFull:
		return "KHVIonExitFull"
	case KHVIonNone:
		return "KHVIonNone"
	default:
		return fmt.Sprintf("KhvIon(%d)", e)
	}
}

type PinBasedIntr uint

const (
	// CPUBased2APICRegVirt: This value controls whether the logical processor virtualizes certain advanced programmable interrupt controller (APIC) accesses.
	CPUBased2APICRegVirt PinBasedIntr = 256
	// CPUBased2DescTable: The value that controls whether executions of descriptor table instructions cause VM exits.
	CPUBased2DescTable PinBasedIntr = 4
	// CPUBased2EnclsExitMap: The value that controls whether executions of Enclave Instruction Leaf Functions (ENCLS) cause examination of the ENCLS-exiting bitmap to determine whether the instruction causes a VM exit.
	CPUBased2EnclsExitMap PinBasedIntr = 32768
	// CPUBased2EnclvExitMap: The value that controls whether executions of an enclave VMM function instruction (ENCLV) checks the ENCLV-exiting bitmap to determine whether the instruction causes a VM exit.
	CPUBased2EnclvExitMap PinBasedIntr = 268435456
	// CPUBased2Ept: The value that controls enabling extended page tables (EPT).
	CPUBased2Ept PinBasedIntr = 2
	// CPUBased2EptModeBasedExec: The value that controls whether to base extended page table (EPT) execute permissions on whether access to a linear address is supervisor or user mode.
	CPUBased2EptModeBasedExec PinBasedIntr = 4194304
	// CPUBased2EptSubpageWrite: The value that controls whether extended page table (EPT) write permissions specify granularity of 128 bytes.
	CPUBased2EptSubpageWrite PinBasedIntr = 8388608
	// CPUBased2EptVe: The value that controls whether extended page table (EPT) violations cause virtualization exceptions instead of VM exits.
	CPUBased2EptVe PinBasedIntr = 262144
	// CPUBased2Invpcid: The value that controls whether any execution of the Invalidate Process-Context Identifier instruction (INVPCID) causes an invalid opcode exception.
	CPUBased2Invpcid PinBasedIntr = 4096
	// CPUBased2PauseLoop: The value that controls whether a series of executions of the PAUSE instruction can cause a VM exit.
	CPUBased2PauseLoop PinBasedIntr = 1024
	// CPUBased2Pml: The value that controls whether an access to a guest-physical address that sets an extended page table (EPT) dirty bit also adds an entry to the page-modification log.
	CPUBased2Pml PinBasedIntr = 131072
	// CPUBased2PtConcealVmx: The value that controls whether the processor trace facility suppresses information that the processor was in VMX non-root operation.
	CPUBased2PtConcealVmx PinBasedIntr = 524288
	// CPUBased2PtGuestPhysical: The value that controls whether to treat all output addresses used by Intel Processor Trace as guest-physical addresses and translated using the extended page table.
	CPUBased2PtGuestPhysical PinBasedIntr = 16777216
	// CPUBased2Rdrand: The value that controls whether executions of the hardware random number generator instruction (RDRAND) cause VM exits.
	CPUBased2Rdrand PinBasedIntr = 2048
	// CPUBased2Rdseed: The value that controls whether executions of random number generator instructions (RDSEED) cause VM exits.
	CPUBased2Rdseed PinBasedIntr = 65536
	// CPUBased2Rdtscp: The value that controls whether any execution of read timestamp-counter and processor ID instruction (RDTSCP) causes an invalid-opcode exception.
	CPUBased2Rdtscp PinBasedIntr = 8
	// CPUBased2TscScaling: The value that controls whether the execution of various read time stamp counters and read model-specific registers that read from the IA32 timestamp counter model specific register return a value modified by the TSC multiplier field.
	CPUBased2TscScaling PinBasedIntr = 33554432
	// CPUBased2Unrestricted: The value that controls whether guest software may run in unpaged protected mode or in real address mode.
	CPUBased2Unrestricted PinBasedIntr = 128
	// CPUBased2UserWaitPause: The value that controls whether any execution of TPAUSE, UMONITOR, or UMWAIT instrucitons generate an illegal opcode exception.
	CPUBased2UserWaitPause PinBasedIntr = 67108864
	// CPUBased2VirtIntrDelivery: The value that enables evaluation and delivery of pending virtual interrupts and emulation of writes to the APIC registers that control interrupt prioritization.
	CPUBased2VirtIntrDelivery PinBasedIntr = 512
	// CPUBased2VirtualAPIC: The value that controls whether the logical processor provides special treatment for access to the Advanced Programmable Interrupt Controller (APIC).
	CPUBased2VirtualAPIC PinBasedIntr = 1
	// CPUBased2VmcsShadow: The value that controls whether execution of VMREAD and VMWRITE in VMX non-root operation may access a shadow VMCS instead of causing a VM exit.
	CPUBased2VmcsShadow PinBasedIntr = 16384
	// CPUBased2Vmfunc: The value that enables use of the “Invoke VM function” (VMFUNC) instruction in VMX non-root operation.
	CPUBased2Vmfunc PinBasedIntr = 8192
	// CPUBased2Vpid: The value that controls the association of cached translations of linear addresses with a virtual processor identifier (VPID).
	CPUBased2Vpid PinBasedIntr = 32
	// CPUBased2Wbinvd: The value that controls whether executions of the Invalidate Cache with Writeback instruction (WBINVD) cause VM exits.
	CPUBased2Wbinvd PinBasedIntr = 64
	// CPUBased2X2apic: The value that controls the logical processor’s treatment of reading/writing of Model Specific Registers to APIC MSRs.
	CPUBased2X2apic PinBasedIntr = 16
	// CPUBased2XsavesXrstors: The value that controls whether any execution of save or restore state instructions (XSAVES or XRSTORS) causes an invalid opcode exception.
	CPUBased2XsavesXrstors PinBasedIntr = 1048576
	// CPUBasedCr3Load: The value that controls whether executions of MOV to Control Register 3 (CR3) cause VM exits.
	CPUBasedCr3Load PinBasedIntr = 32768
	// CPUBasedCr3Store: The value that controls whether executions of MOV from Control Register 3 (CR3) cause VM exits.
	CPUBasedCr3Store PinBasedIntr = 65536
	// CPUBasedCr8Load: The value that controls whether executions of MOV to Control Register 8 (CR8) cause VM exits.
	CPUBasedCr8Load PinBasedIntr = 524288
	// CPUBasedCr8Store: The value that controls whether executions of MOV from Control Register 8 (CR8) cause VM exits.
	CPUBasedCr8Store PinBasedIntr = 1048576
	// CPUBasedHlt: The value that controls whether the execution of HALT instructions cause VM exits.
	CPUBasedHlt PinBasedIntr = 128
	// CPUBasedIOBitmaps: The value that controls whether to use I/O bitmaps to restrict executions of I/O instructions.
	CPUBasedIOBitmaps PinBasedIntr = 33554432
	// CPUBasedInvlpg: The value that controls whether the execution of invalid page instructions (INVLPG) cause VM exits.
	CPUBasedInvlpg PinBasedIntr = 512
	// CPUBasedIrqWnd: The value that controls whether a VM exits at the beginning of any instruction where there’s no blocking of interrupts and the interrupt flag is 1.
	CPUBasedIrqWnd PinBasedIntr = 4
	// CPUBasedMonitor: The value that controls whether executions of the Set Up Monitor Address instruction (MONITOR) cause VM exits.
	CPUBasedMonitor PinBasedIntr = 536870912
	// CPUBasedMovDr: The value that controls whether executions of MOV to or from Debug Registers (DR) cause VM exits.
	CPUBasedMovDr PinBasedIntr = 8388608
	// CPUBasedMsrBitmaps: The value that controls use of whether Model Specific Register (MSR) bitmaps to control execution of the read-from and write-to MSR instructions.
	CPUBasedMsrBitmaps PinBasedIntr = 268435456
	// CPUBasedMtf: The value that controls enabling the monitor trap flag debugging feature.
	CPUBasedMtf PinBasedIntr = 134217728
	// CPUBasedMwait: The value that controls whether the execution of Monitor Wait instructions (MWAIT) cause VM exits.
	CPUBasedMwait PinBasedIntr = 1024
	// CPUBasedPause: The value that controls whether executions of spin-wait loop (PAUSE) instruction causes VM exits.
	CPUBasedPause PinBasedIntr = 1073741824
	// CPUBasedRdpmc: The value that controls whether the execution of Read Performance Monitoring Counters instructions (RDPMC) cause VM exits.
	CPUBasedRdpmc PinBasedIntr = 2048
	// CPUBasedRdtsc: The value that controls whether the execution of Read Timestamp-Counter instructions (RDTSC) cause VM exits.
	CPUBasedRdtsc PinBasedIntr = 4096
	// CPUBasedSecondaryCtls: The value that conntrols use of the secondary processor-based VM-execution controls.
	CPUBasedSecondaryCtls PinBasedIntr = 2147483648
	// CPUBasedTprShadow: The value that controls enabling Task Priority Register (TPR) virtualization and other APIC-virtualization features.
	CPUBasedTprShadow PinBasedIntr = 2097152
	// CPUBasedTscOffset: The value that controls whether reading the timestamp-counter MSRs changes depending on the value of the timestamp-counter offset field.
	CPUBasedTscOffset PinBasedIntr = 8
	// CPUBasedUncondIO: The value that controls whether executions of various I/O instructions cause VM exits.
	CPUBasedUncondIO PinBasedIntr = 16777216
	// CPUBasedVirtualNmiWnd: The value that controls if a VM exit occurs at the beginning of any instruction if there’s no virtual-NMI blocking.
	CPUBasedVirtualNmiWnd PinBasedIntr = 4194304
	// PinBasedIntrValue: The value that controls whether external interrupts cause VM exits.
	PinBasedIntrValue PinBasedIntr = 1
	// PinBasedNmi: The value that controls whether external non-maskable interrupts cause VM exits.
	PinBasedNmi PinBasedIntr = 8
	// PinBasedPostedIntr: The value that controls whether the processor gives special treatment to interrupts with posted-interrupt notification vectors.
	PinBasedPostedIntr PinBasedIntr = 128
	// PinBasedPreemptionTimer: The value that controls whether the VMX-preemption timer counts down in VMX non-root operation.
	PinBasedPreemptionTimer PinBasedIntr = 64
	// PinBasedVirtualNmi: The value that controls blocking of non-maskable interrupts.
	PinBasedVirtualNmi PinBasedIntr = 32
	// VmentryDeactivateDualMonitor: The value that controls whether the treatment of SMIs and system-management mode (SMM) is in effect after the VM entry.
	VmentryDeactivateDualMonitor PinBasedIntr = 2048
	// VmentryGuestIa32e: The value that controls whether the logical processor is in IA-32e mode after VM entry.
	VmentryGuestIa32e PinBasedIntr = 512
	// VmentryLoadCetState: The value that controls whether to load CET-related model specific registers and SPP on VM exit.
	VmentryLoadCetState PinBasedIntr = 1048576
	// VmentryLoadDbgControls: The value that controls whetherto load Debug Register 7 and the IA32_DEBUGCTL model specific register (MSR) on VM entry.
	VmentryLoadDbgControls PinBasedIntr = 4
	// VmentryLoadEfer: The value that determines whether to load the IA32_EFER model specific register on VM entry.
	VmentryLoadEfer PinBasedIntr = 32768
	// VmentryLoadIa32Bndcfgs: The value that controls whether to load the IA32_BNDCFGS model specific register on VM entry.
	VmentryLoadIa32Bndcfgs PinBasedIntr = 65536
	// VmentryLoadIa32Pat: The value that controls whether to load the IA32_PAT model specific register on VM entry.
	VmentryLoadIa32Pat PinBasedIntr = 16384
	// VmentryLoadIa32PerfGlobalCtrl: The value that controls whether to load the IA32_PERF_GLOBAL_CTRL model specific register on VM entry.
	VmentryLoadIa32PerfGlobalCtrl PinBasedIntr = 8192
	// VmentryLoadIa32RtitCtl: The value that controls whether to clear the IA32_RTIT_CTL model specific register (MSR) on VM exit.
	VmentryLoadIa32RtitCtl PinBasedIntr = 262144
	VmentryLoadPkrs        PinBasedIntr = 4194304
	// VmentryPtConcealVmx: The value that controls whether the Intel Processor Trace produces a paging information packet (PIP) on a VM entry or a VMCS packet on a VM entry that returns from system-management mode.
	VmentryPtConcealVmx PinBasedIntr = 131072
	// VmentrySmm: The value that controls whether the logical processor is in system-management mode (SMM) after VM entry.
	VmentrySmm PinBasedIntr = 1024
	// VmexitAckIntr: The value that controls whether the logical processor sends an acknowledgement to the interrupt controller when the VM exits.
	VmexitAckIntr PinBasedIntr = 32768
	// VmexitClearIa32Bndcfgs: The value that controls whether to clear the IA32_BNDCFGS model specific register on VM exit.
	VmexitClearIa32Bndcfgs PinBasedIntr = 8388608
	// VmexitClearIa32RtitCtl: The value that controls whether to clear the IA32_RTIT_CTL model specific register (MSR) on VM exit.
	VmexitClearIa32RtitCtl PinBasedIntr = 33554432
	// VmexitHostIa32e: This value controls, on processors that support Intel 64 architecture, whether a logical processor is in 64-bit mode after the next VM exit.
	VmexitHostIa32e PinBasedIntr = 512
	// VmexitLoadCetState: The value that controls whether to load CET-related MSRs and SPP on VM exit.
	VmexitLoadCetState PinBasedIntr = 268435456
	// VmexitLoadEfer: The value that controls whether to load the IA32_EFER MSR on VM exit.
	VmexitLoadEfer PinBasedIntr = 2097152
	// VmexitLoadIa32Pat: The value that controls whether to load the IA32_EFER mode specific register on VM exit.
	VmexitLoadIa32Pat PinBasedIntr = 524288
	// VmexitLoadIa32PerfGlobalCtrl: The value that controls whether to load the IA32_PERF_GLOBAL_CTRL model specific register on VM exit.
	VmexitLoadIa32PerfGlobalCtrl PinBasedIntr = 4096
	VmexitLoadPkrs               PinBasedIntr = 536870912
	// VmexitPtConcealVmx: The value that controls whether the Intel Processor Trace produces a paging information packet on VM exit or a VMCS packet on SMM VM exit.
	VmexitPtConcealVmx PinBasedIntr = 16777216
	// VmexitSaveDbgControls: Thievalue that controls whether to save debug register 7 DR7 and the IA32 debug control DEBUGCTL MSR on VM exit.
	VmexitSaveDbgControls PinBasedIntr = 4
	// VmexitSaveEfer: The value that controls whether to save the IA32_EFER MSR on VM exit.
	VmexitSaveEfer PinBasedIntr = 1048576
	// VmexitSaveIa32Pat: The value that controls whether to save the IA32_EFER model specific register on VM exit.
	VmexitSaveIa32Pat PinBasedIntr = 262144
	// VmexitSaveVmxTimer: The value that controls whether to save the value of the VMX-preemption timer on VM exit.
	VmexitSaveVmxTimer      PinBasedIntr = 4194304
	VmxEptVpidAdvVmexitInfo PinBasedIntr = 4194304
	// VmxEptVpidSupportAd: The value that controls if extended page tables (EPT) support accessed and dirty flags.
	VmxEptVpidSupportAd PinBasedIntr = 2097152
	// VmxEptVpidSupportExonly: The value that controls whether extended page tables (EPT) support execute-only translations.
	VmxEptVpidSupportExonly PinBasedIntr = 1
)

func (e PinBasedIntr) String() string {
	switch e {
	case CPUBased2APICRegVirt:
		return "CPUBased2APICRegVirt"
	case CPUBased2DescTable:
		return "CPUBased2DescTable"
	case CPUBased2EnclsExitMap:
		return "CPUBased2EnclsExitMap"
	case CPUBased2EnclvExitMap:
		return "CPUBased2EnclvExitMap"
	case CPUBased2Ept:
		return "CPUBased2Ept"
	case CPUBased2EptModeBasedExec:
		return "CPUBased2EptModeBasedExec"
	case CPUBased2EptSubpageWrite:
		return "CPUBased2EptSubpageWrite"
	case CPUBased2EptVe:
		return "CPUBased2EptVe"
	case CPUBased2Invpcid:
		return "CPUBased2Invpcid"
	case CPUBased2PauseLoop:
		return "CPUBased2PauseLoop"
	case CPUBased2Pml:
		return "CPUBased2Pml"
	case CPUBased2PtConcealVmx:
		return "CPUBased2PtConcealVmx"
	case CPUBased2PtGuestPhysical:
		return "CPUBased2PtGuestPhysical"
	case CPUBased2Rdrand:
		return "CPUBased2Rdrand"
	case CPUBased2Rdseed:
		return "CPUBased2Rdseed"
	case CPUBased2Rdtscp:
		return "CPUBased2Rdtscp"
	case CPUBased2TscScaling:
		return "CPUBased2TscScaling"
	case CPUBased2Unrestricted:
		return "CPUBased2Unrestricted"
	case CPUBased2UserWaitPause:
		return "CPUBased2UserWaitPause"
	case CPUBased2VirtIntrDelivery:
		return "CPUBased2VirtIntrDelivery"
	case CPUBased2VirtualAPIC:
		return "CPUBased2VirtualAPIC"
	case CPUBased2VmcsShadow:
		return "CPUBased2VmcsShadow"
	case CPUBased2Vmfunc:
		return "CPUBased2Vmfunc"
	case CPUBased2Vpid:
		return "CPUBased2Vpid"
	case CPUBased2Wbinvd:
		return "CPUBased2Wbinvd"
	case CPUBased2X2apic:
		return "CPUBased2X2apic"
	case CPUBased2XsavesXrstors:
		return "CPUBased2XsavesXrstors"
	case CPUBasedMonitor:
		return "CPUBasedMonitor"
	case CPUBasedMtf:
		return "CPUBasedMtf"
	case CPUBasedPause:
		return "CPUBasedPause"
	case CPUBasedSecondaryCtls:
		return "CPUBasedSecondaryCtls"
	case CPUBasedTprShadow:
		return "CPUBasedTprShadow"
	default:
		return fmt.Sprintf("PinBasedIntr(%d)", e)
	}
}

type Vmcs uint

const (
	VmcsCtrlAPICAccess            Vmcs = 0x2014
	VmcsCtrlCPUBased              Vmcs = 0x4002
	VmcsCtrlCPUBased2             Vmcs = 0x401e
	VmcsCtrlCr0Mask               Vmcs = 0x6000
	VmcsCtrlCr0Shadow             Vmcs = 0x6004
	VmcsCtrlCr3Count              Vmcs = 0x400a
	VmcsCtrlCr3Value0             Vmcs = 0x6008
	VmcsCtrlCr3Value1             Vmcs = 0x600a
	VmcsCtrlCr3Value2             Vmcs = 0x600c
	VmcsCtrlCr3Value3             Vmcs = 0x600e
	VmcsCtrlCr4Mask               Vmcs = 0x6002
	VmcsCtrlCr4Shadow             Vmcs = 0x6006
	VmcsCtrlEnclsExitingBitmap    Vmcs = 0x202e
	VmcsCtrlEnclvExitingBitmap    Vmcs = 0x2036
	VmcsCtrlEoiExitBitmap0        Vmcs = 0x201c
	VmcsCtrlEoiExitBitmap1        Vmcs = 0x201e
	VmcsCtrlEoiExitBitmap2        Vmcs = 0x2020
	VmcsCtrlEoiExitBitmap3        Vmcs = 0x2022
	VmcsCtrlEptp                  Vmcs = 0x201a
	VmcsCtrlEptpIndex             Vmcs = 0x4
	VmcsCtrlEptpListAddr          Vmcs = 0x2024
	VmcsCtrlExcBitmap             Vmcs = 0x4004
	VmcsCtrlExecutiveVmcsPtr      Vmcs = 0x200c
	VmcsCtrlIOBitmapA             Vmcs = 0x2000
	VmcsCtrlIOBitmapB             Vmcs = 0x2002
	VmcsCtrlMsrBitmaps            Vmcs = 0x2004
	VmcsCtrlPfErrorMask           Vmcs = 0x4006
	VmcsCtrlPfErrorMatch          Vmcs = 0x4008
	VmcsCtrlPinBased              Vmcs = 0x4000
	VmcsCtrlPleGap                Vmcs = 0x4020
	VmcsCtrlPleWindow             Vmcs = 0x4022
	VmcsCtrlPmlAddr               Vmcs = 0x200e
	VmcsCtrlPostedIntDescAddr     Vmcs = 0x2016
	VmcsCtrlPostedIntNVector      Vmcs = 0x2
	VmcsCtrlSppTable              Vmcs = 0x2030
	VmcsCtrlTprThreshold          Vmcs = 0x401c
	VmcsCtrlTscMultiplier         Vmcs = 0x2032
	VmcsCtrlTscOffset             Vmcs = 0x2010
	VmcsCtrlVirtExcInfoAddr       Vmcs = 0x202a
	VmcsCtrlVirtualAPIC           Vmcs = 0x2012
	VmcsCtrlVmentryControls       Vmcs = 0x4012
	VmcsCtrlVmentryExcError       Vmcs = 0x4018
	VmcsCtrlVmentryInstrLen       Vmcs = 0x401a
	VmcsCtrlVmentryIrqInfo        Vmcs = 0x4016
	VmcsCtrlVmentryMsrLoadAddr    Vmcs = 0x200a
	VmcsCtrlVmentryMsrLoadCount   Vmcs = 0x4014
	VmcsCtrlVmexitControls        Vmcs = 0x400c
	VmcsCtrlVmexitMsrLoadAddr     Vmcs = 0x2008
	VmcsCtrlVmexitMsrLoadCount    Vmcs = 0x4010
	VmcsCtrlVmexitMsrStoreAddr    Vmcs = 0x2006
	VmcsCtrlVmexitMsrStoreCount   Vmcs = 0x400e
	VmcsCtrlVmfuncCtrl            Vmcs = 0x2018
	VmcsCtrlVmreadBitmapAddr      Vmcs = 0x2026
	VmcsCtrlVmwriteBitmapAddr     Vmcs = 0x2028
	VmcsCtrlXssExitingBitmap      Vmcs = 0x202c
	VmcsGuestActivityState        Vmcs = 0x4826
	VmcsGuestCr0                  Vmcs = 0x6800
	VmcsGuestCr3                  Vmcs = 0x6802
	VmcsGuestCr4                  Vmcs = 0x6804
	VmcsGuestCs                   Vmcs = 0x802
	VmcsGuestCsAr                 Vmcs = 0x4816
	VmcsGuestCsBase               Vmcs = 0x6808
	VmcsGuestCsLimit              Vmcs = 0x4802
	VmcsGuestDebugExc             Vmcs = 0x6822
	VmcsGuestDr7                  Vmcs = 0x681a
	VmcsGuestDs                   Vmcs = 0x806
	VmcsGuestDsAr                 Vmcs = 0x481a
	VmcsGuestDsBase               Vmcs = 0x680c
	VmcsGuestDsLimit              Vmcs = 0x4806
	VmcsGuestEs                   Vmcs = 0x800
	VmcsGuestEsAr                 Vmcs = 0x4814
	VmcsGuestEsBase               Vmcs = 0x6806
	VmcsGuestEsLimit              Vmcs = 0x4800
	VmcsGuestFs                   Vmcs = 0x808
	VmcsGuestFsAr                 Vmcs = 0x481c
	VmcsGuestFsBase               Vmcs = 0x680e
	VmcsGuestFsLimit              Vmcs = 0x4808
	VmcsGuestGdtrBase             Vmcs = 0x6816
	VmcsGuestGdtrLimit            Vmcs = 0x4810
	VmcsGuestGs                   Vmcs = 0x80a
	VmcsGuestGsAr                 Vmcs = 0x481e
	VmcsGuestGsBase               Vmcs = 0x6810
	VmcsGuestGsLimit              Vmcs = 0x480a
	VmcsGuestIa32Bndcfgs          Vmcs = 0x2812
	VmcsGuestIa32Debugctl         Vmcs = 0x2802
	VmcsGuestIa32Efer             Vmcs = 0x2806
	VmcsGuestIa32IntrSspTableAddr Vmcs = 0x682c
	VmcsGuestIa32Pat              Vmcs = 0x2804
	VmcsGuestIa32PerfGlobalCtrl   Vmcs = 0x2808
	VmcsGuestIa32Pkrs             Vmcs = 0x2818
	VmcsGuestIa32RtitCtl          Vmcs = 0x2814
	VmcsGuestIa32SCet             Vmcs = 0x6828
	VmcsGuestIa32SysenterCs       Vmcs = 0x482a
	VmcsGuestIdtrBase             Vmcs = 0x6818
	VmcsGuestIdtrLimit            Vmcs = 0x4812
	VmcsGuestIgnoreIrq            Vmcs = 18468
	VmcsGuestIntStatus            Vmcs = 0x810
	VmcsGuestInterruptibility     Vmcs = 0x4824
	VmcsGuestLdtr                 Vmcs = 0x80c
	VmcsGuestLdtrAr               Vmcs = 0x4820
	VmcsGuestLdtrBase             Vmcs = 0x6812
	VmcsGuestLdtrLimit            Vmcs = 0x480c
	VmcsGuestLinkPointer          Vmcs = 0x2800
	VmcsGuestPdpte0               Vmcs = 0x280a
	VmcsGuestPdpte1               Vmcs = 0x280c
	VmcsGuestPdpte2               Vmcs = 0x280e
	VmcsGuestPdpte3               Vmcs = 0x2810
	VmcsGuestPhysicalAddress      Vmcs = 0x2400
	VmcsGuestRflags               Vmcs = 0x6820
	VmcsGuestRip                  Vmcs = 0x681e
	VmcsGuestRsp                  Vmcs = 0x681c
	VmcsGuestSmbase               Vmcs = 0x4828
	VmcsGuestSs                   Vmcs = 0x804
	VmcsGuestSsAr                 Vmcs = 0x4818
	VmcsGuestSsBase               Vmcs = 0x680a
	VmcsGuestSsLimit              Vmcs = 0x4804
	VmcsGuestSsp                  Vmcs = 0x682a
	VmcsGuestSysenterEip          Vmcs = 0x6826
	VmcsGuestSysenterEsp          Vmcs = 0x6824
	VmcsGuestTr                   Vmcs = 0x80e
	VmcsGuestTrAr                 Vmcs = 0x4822
	VmcsGuestTrBase               Vmcs = 0x6814
	VmcsGuestTrLimit              Vmcs = 0x480e
	VmcsGuestVmxTimerValue        Vmcs = 0x482e
	VmcsGuestpmlIndex             Vmcs = 0x812
	VmcsHostCr0                   Vmcs = 0x6c00
	VmcsHostCr3                   Vmcs = 0x6c02
	VmcsHostCr4                   Vmcs = 0x6c04
	VmcsHostCs                    Vmcs = 0xc02
	VmcsHostDs                    Vmcs = 0xc06
	VmcsHostEs                    Vmcs = 0xc00
	VmcsHostFs                    Vmcs = 0xc08
	VmcsHostFsBase                Vmcs = 0x6c06
	VmcsHostGdtrBase              Vmcs = 0x6c0c
	VmcsHostGs                    Vmcs = 0xc0a
	VmcsHostGsBase                Vmcs = 0x6c08
	VmcsHostIa32Efer              Vmcs = 0x2c02
	VmcsHostIa32IntrSspTableAddr  Vmcs = 0x6c1c
	VmcsHostIa32Pat               Vmcs = 0x2c00
	VmcsHostIa32PerfGlobalCtrl    Vmcs = 0x2c04
	VmcsHostIa32Pkrs              Vmcs = 0x2c06
	VmcsHostIa32SCet              Vmcs = 0x6c18
	VmcsHostIa32SysenterCs        Vmcs = 0x4c00
	VmcsHostIa32SysenterEip       Vmcs = 0x6c12
	VmcsHostIa32SysenterEsp       Vmcs = 0x6c10
	VmcsHostIdtrBase              Vmcs = 0x6c0e
	VmcsHostRip                   Vmcs = 0x6c16
	VmcsHostRsp                   Vmcs = 0x6c14
	VmcsHostSs                    Vmcs = 0xc04
	VmcsHostSsp                   Vmcs = 0x6c1a
	VmcsHostTr                    Vmcs = 0xc0c
	VmcsHostTrBase                Vmcs = 0x6c0a
	VmcsInvalid                   Vmcs = 27904
	VmcsMax                       Vmcs = 0x6d00
	VmcsRoExitQualific            Vmcs = 0x6400
	VmcsRoExitReason              Vmcs = 0x4402
	VmcsRoGuestLinAddr            Vmcs = 0x640a
	VmcsRoIORcx                   Vmcs = 0x6402
	VmcsRoIORdi                   Vmcs = 0x6406
	VmcsRoIORip                   Vmcs = 0x6408
	VmcsRoIORsi                   Vmcs = 0x6404
	VmcsRoIdtVectorError          Vmcs = 0x440a
	VmcsRoIdtVectorInfo           Vmcs = 0x4408
	VmcsRoInstrError              Vmcs = 0x4400
	VmcsRoVmexitInstrLen          Vmcs = 0x440c
	VmcsRoVmexitIrqError          Vmcs = 0x4406
	VmcsRoVmexitIrqInfo           Vmcs = 0x4404
	VmcsRoVmxInstrInfo            Vmcs = 0x440e
	VmcsVpid                      Vmcs = 0
)

func (e Vmcs) String() string {
	switch e {
	case VmcsCtrlAPICAccess:
		return "VmcsCtrlAPICAccess"
	case VmcsCtrlCPUBased:
		return "VmcsCtrlCPUBased"
	case VmcsCtrlCPUBased2:
		return "VmcsCtrlCPUBased2"
	case VmcsCtrlCr0Mask:
		return "VmcsCtrlCr0Mask"
	case VmcsCtrlCr0Shadow:
		return "VmcsCtrlCr0Shadow"
	case VmcsCtrlCr3Count:
		return "VmcsCtrlCr3Count"
	case VmcsCtrlCr3Value0:
		return "VmcsCtrlCr3Value0"
	case VmcsCtrlCr3Value1:
		return "VmcsCtrlCr3Value1"
	case VmcsCtrlCr3Value2:
		return "VmcsCtrlCr3Value2"
	case VmcsCtrlCr3Value3:
		return "VmcsCtrlCr3Value3"
	case VmcsCtrlCr4Mask:
		return "VmcsCtrlCr4Mask"
	case VmcsCtrlCr4Shadow:
		return "VmcsCtrlCr4Shadow"
	case VmcsCtrlEnclsExitingBitmap:
		return "VmcsCtrlEnclsExitingBitmap"
	case VmcsCtrlEnclvExitingBitmap:
		return "VmcsCtrlEnclvExitingBitmap"
	case VmcsCtrlEoiExitBitmap0:
		return "VmcsCtrlEoiExitBitmap0"
	case VmcsCtrlEoiExitBitmap1:
		return "VmcsCtrlEoiExitBitmap1"
	case VmcsCtrlEoiExitBitmap2:
		return "VmcsCtrlEoiExitBitmap2"
	case VmcsCtrlEoiExitBitmap3:
		return "VmcsCtrlEoiExitBitmap3"
	case VmcsCtrlEptp:
		return "VmcsCtrlEptp"
	case VmcsCtrlEptpIndex:
		return "VmcsCtrlEptpIndex"
	case VmcsCtrlEptpListAddr:
		return "VmcsCtrlEptpListAddr"
	case VmcsCtrlExcBitmap:
		return "VmcsCtrlExcBitmap"
	case VmcsCtrlExecutiveVmcsPtr:
		return "VmcsCtrlExecutiveVmcsPtr"
	case VmcsCtrlIOBitmapA:
		return "VmcsCtrlIOBitmapA"
	case VmcsCtrlIOBitmapB:
		return "VmcsCtrlIOBitmapB"
	case VmcsCtrlMsrBitmaps:
		return "VmcsCtrlMsrBitmaps"
	case VmcsCtrlPfErrorMask:
		return "VmcsCtrlPfErrorMask"
	case VmcsCtrlPfErrorMatch:
		return "VmcsCtrlPfErrorMatch"
	case VmcsCtrlPinBased:
		return "VmcsCtrlPinBased"
	case VmcsCtrlPleGap:
		return "VmcsCtrlPleGap"
	case VmcsCtrlPleWindow:
		return "VmcsCtrlPleWindow"
	case VmcsCtrlPmlAddr:
		return "VmcsCtrlPmlAddr"
	case VmcsCtrlPostedIntDescAddr:
		return "VmcsCtrlPostedIntDescAddr"
	case VmcsCtrlPostedIntNVector:
		return "VmcsCtrlPostedIntNVector"
	case VmcsCtrlSppTable:
		return "VmcsCtrlSppTable"
	case VmcsCtrlTprThreshold:
		return "VmcsCtrlTprThreshold"
	case VmcsCtrlTscMultiplier:
		return "VmcsCtrlTscMultiplier"
	case VmcsCtrlTscOffset:
		return "VmcsCtrlTscOffset"
	case VmcsCtrlVirtExcInfoAddr:
		return "VmcsCtrlVirtExcInfoAddr"
	case VmcsCtrlVirtualAPIC:
		return "VmcsCtrlVirtualAPIC"
	case VmcsCtrlVmentryControls:
		return "VmcsCtrlVmentryControls"
	case VmcsCtrlVmentryExcError:
		return "VmcsCtrlVmentryExcError"
	case VmcsCtrlVmentryInstrLen:
		return "VmcsCtrlVmentryInstrLen"
	case VmcsCtrlVmentryIrqInfo:
		return "VmcsCtrlVmentryIrqInfo"
	case VmcsCtrlVmentryMsrLoadAddr:
		return "VmcsCtrlVmentryMsrLoadAddr"
	case VmcsCtrlVmentryMsrLoadCount:
		return "VmcsCtrlVmentryMsrLoadCount"
	case VmcsCtrlVmexitControls:
		return "VmcsCtrlVmexitControls"
	case VmcsCtrlVmexitMsrLoadAddr:
		return "VmcsCtrlVmexitMsrLoadAddr"
	case VmcsCtrlVmexitMsrLoadCount:
		return "VmcsCtrlVmexitMsrLoadCount"
	case VmcsCtrlVmexitMsrStoreAddr:
		return "VmcsCtrlVmexitMsrStoreAddr"
	case VmcsCtrlVmexitMsrStoreCount:
		return "VmcsCtrlVmexitMsrStoreCount"
	case VmcsCtrlVmfuncCtrl:
		return "VmcsCtrlVmfuncCtrl"
	case VmcsCtrlVmreadBitmapAddr:
		return "VmcsCtrlVmreadBitmapAddr"
	case VmcsCtrlVmwriteBitmapAddr:
		return "VmcsCtrlVmwriteBitmapAddr"
	case VmcsCtrlXssExitingBitmap:
		return "VmcsCtrlXssExitingBitmap"
	case VmcsGuestActivityState:
		return "VmcsGuestActivityState"
	case VmcsGuestCr0:
		return "VmcsGuestCr0"
	case VmcsGuestCr3:
		return "VmcsGuestCr3"
	case VmcsGuestCr4:
		return "VmcsGuestCr4"
	case VmcsGuestCs:
		return "VmcsGuestCs"
	case VmcsGuestCsAr:
		return "VmcsGuestCsAr"
	case VmcsGuestCsBase:
		return "VmcsGuestCsBase"
	case VmcsGuestCsLimit:
		return "VmcsGuestCsLimit"
	case VmcsGuestDebugExc:
		return "VmcsGuestDebugExc"
	case VmcsGuestDr7:
		return "VmcsGuestDr7"
	case VmcsGuestDs:
		return "VmcsGuestDs"
	case VmcsGuestDsAr:
		return "VmcsGuestDsAr"
	case VmcsGuestDsBase:
		return "VmcsGuestDsBase"
	case VmcsGuestDsLimit:
		return "VmcsGuestDsLimit"
	case VmcsGuestEs:
		return "VmcsGuestEs"
	case VmcsGuestEsAr:
		return "VmcsGuestEsAr"
	case VmcsGuestEsBase:
		return "VmcsGuestEsBase"
	case VmcsGuestEsLimit:
		return "VmcsGuestEsLimit"
	case VmcsGuestFs:
		return "VmcsGuestFs"
	case VmcsGuestFsAr:
		return "VmcsGuestFsAr"
	case VmcsGuestFsBase:
		return "VmcsGuestFsBase"
	case VmcsGuestFsLimit:
		return "VmcsGuestFsLimit"
	case VmcsGuestGdtrBase:
		return "VmcsGuestGdtrBase"
	case VmcsGuestGdtrLimit:
		return "VmcsGuestGdtrLimit"
	case VmcsGuestGs:
		return "VmcsGuestGs"
	case VmcsGuestGsAr:
		return "VmcsGuestGsAr"
	case VmcsGuestGsBase:
		return "VmcsGuestGsBase"
	case VmcsGuestGsLimit:
		return "VmcsGuestGsLimit"
	case VmcsGuestIa32Bndcfgs:
		return "VmcsGuestIa32Bndcfgs"
	case VmcsGuestIa32Debugctl:
		return "VmcsGuestIa32Debugctl"
	case VmcsGuestIa32Efer:
		return "VmcsGuestIa32Efer"
	case VmcsGuestIa32IntrSspTableAddr:
		return "VmcsGuestIa32IntrSspTableAddr"
	case VmcsGuestIa32Pat:
		return "VmcsGuestIa32Pat"
	case VmcsGuestIa32PerfGlobalCtrl:
		return "VmcsGuestIa32PerfGlobalCtrl"
	case VmcsGuestIa32Pkrs:
		return "VmcsGuestIa32Pkrs"
	case VmcsGuestIa32RtitCtl:
		return "VmcsGuestIa32RtitCtl"
	case VmcsGuestIa32SCet:
		return "VmcsGuestIa32SCet"
	case VmcsGuestIa32SysenterCs:
		return "VmcsGuestIa32SysenterCs"
	case VmcsGuestIdtrBase:
		return "VmcsGuestIdtrBase"
	case VmcsGuestIdtrLimit:
		return "VmcsGuestIdtrLimit"
	case VmcsGuestIgnoreIrq:
		return "VmcsGuestIgnoreIrq"
	case VmcsGuestIntStatus:
		return "VmcsGuestIntStatus"
	case VmcsGuestLdtr:
		return "VmcsGuestLdtr"
	case VmcsGuestLdtrAr:
		return "VmcsGuestLdtrAr"
	case VmcsGuestLdtrBase:
		return "VmcsGuestLdtrBase"
	case VmcsGuestLdtrLimit:
		return "VmcsGuestLdtrLimit"
	case VmcsGuestLinkPointer:
		return "VmcsGuestLinkPointer"
	case VmcsGuestPdpte0:
		return "VmcsGuestPdpte0"
	case VmcsGuestPdpte1:
		return "VmcsGuestPdpte1"
	case VmcsGuestPdpte2:
		return "VmcsGuestPdpte2"
	case VmcsGuestPdpte3:
		return "VmcsGuestPdpte3"
	case VmcsGuestPhysicalAddress:
		return "VmcsGuestPhysicalAddress"
	case VmcsGuestRflags:
		return "VmcsGuestRflags"
	case VmcsGuestRip:
		return "VmcsGuestRip"
	case VmcsGuestRsp:
		return "VmcsGuestRsp"
	case VmcsGuestSmbase:
		return "VmcsGuestSmbase"
	case VmcsGuestSs:
		return "VmcsGuestSs"
	case VmcsGuestSsAr:
		return "VmcsGuestSsAr"
	case VmcsGuestSsBase:
		return "VmcsGuestSsBase"
	case VmcsGuestSsLimit:
		return "VmcsGuestSsLimit"
	case VmcsGuestSsp:
		return "VmcsGuestSsp"
	case VmcsGuestSysenterEip:
		return "VmcsGuestSysenterEip"
	case VmcsGuestSysenterEsp:
		return "VmcsGuestSysenterEsp"
	case VmcsGuestTr:
		return "VmcsGuestTr"
	case VmcsGuestTrAr:
		return "VmcsGuestTrAr"
	case VmcsGuestTrBase:
		return "VmcsGuestTrBase"
	case VmcsGuestTrLimit:
		return "VmcsGuestTrLimit"
	case VmcsGuestVmxTimerValue:
		return "VmcsGuestVmxTimerValue"
	case VmcsGuestpmlIndex:
		return "VmcsGuestpmlIndex"
	case VmcsHostCr0:
		return "VmcsHostCr0"
	case VmcsHostCr3:
		return "VmcsHostCr3"
	case VmcsHostCr4:
		return "VmcsHostCr4"
	case VmcsHostCs:
		return "VmcsHostCs"
	case VmcsHostDs:
		return "VmcsHostDs"
	case VmcsHostEs:
		return "VmcsHostEs"
	case VmcsHostFs:
		return "VmcsHostFs"
	case VmcsHostFsBase:
		return "VmcsHostFsBase"
	case VmcsHostGdtrBase:
		return "VmcsHostGdtrBase"
	case VmcsHostGs:
		return "VmcsHostGs"
	case VmcsHostGsBase:
		return "VmcsHostGsBase"
	case VmcsHostIa32Efer:
		return "VmcsHostIa32Efer"
	case VmcsHostIa32IntrSspTableAddr:
		return "VmcsHostIa32IntrSspTableAddr"
	case VmcsHostIa32Pat:
		return "VmcsHostIa32Pat"
	case VmcsHostIa32PerfGlobalCtrl:
		return "VmcsHostIa32PerfGlobalCtrl"
	case VmcsHostIa32Pkrs:
		return "VmcsHostIa32Pkrs"
	case VmcsHostIa32SCet:
		return "VmcsHostIa32SCet"
	case VmcsHostIa32SysenterCs:
		return "VmcsHostIa32SysenterCs"
	case VmcsHostIa32SysenterEip:
		return "VmcsHostIa32SysenterEip"
	case VmcsHostIa32SysenterEsp:
		return "VmcsHostIa32SysenterEsp"
	case VmcsHostIdtrBase:
		return "VmcsHostIdtrBase"
	case VmcsHostRip:
		return "VmcsHostRip"
	case VmcsHostRsp:
		return "VmcsHostRsp"
	case VmcsHostSs:
		return "VmcsHostSs"
	case VmcsHostSsp:
		return "VmcsHostSsp"
	case VmcsHostTr:
		return "VmcsHostTr"
	case VmcsHostTrBase:
		return "VmcsHostTrBase"
	case VmcsInvalid:
		return "VmcsInvalid"
	case VmcsRoExitQualific:
		return "VmcsRoExitQualific"
	case VmcsRoExitReason:
		return "VmcsRoExitReason"
	case VmcsRoGuestLinAddr:
		return "VmcsRoGuestLinAddr"
	case VmcsRoIORcx:
		return "VmcsRoIORcx"
	case VmcsRoIORdi:
		return "VmcsRoIORdi"
	case VmcsRoIORip:
		return "VmcsRoIORip"
	case VmcsRoIORsi:
		return "VmcsRoIORsi"
	case VmcsRoIdtVectorError:
		return "VmcsRoIdtVectorError"
	case VmcsRoIdtVectorInfo:
		return "VmcsRoIdtVectorInfo"
	case VmcsRoInstrError:
		return "VmcsRoInstrError"
	case VmcsRoVmexitInstrLen:
		return "VmcsRoVmexitInstrLen"
	case VmcsRoVmexitIrqError:
		return "VmcsRoVmexitIrqError"
	case VmcsRoVmexitIrqInfo:
		return "VmcsRoVmexitIrqInfo"
	case VmcsRoVmxInstrInfo:
		return "VmcsRoVmxInstrInfo"
	case VmcsVpid:
		return "VmcsVpid"
	default:
		return fmt.Sprintf("Vmcs(%d)", e)
	}
}

type VmxBasicTrue uint

const (
	// VmxBasicTrueCtls: This bit field, in the value returned by the IA32_VMX_BASIC model specific register, determines if it’s possible to disable any VMX controls.
	VmxBasicTrueCtls VmxBasicTrue = 36028797018963968
)

func (e VmxBasicTrue) String() string {
	switch e {
	case VmxBasicTrueCtls:
		return "VmxBasicTrueCtls"
	default:
		return fmt.Sprintf("VmxBasicTrue(%d)", e)
	}
}

type VmxReason uint

const (
	// VmxReasonAPICAccess: The guest attempted to access memory at a physical address on the APIC-access page and the “virtualize APIC accesses” VM-execution control was 1.
	VmxReasonAPICAccess VmxReason = 44
	// VmxReasonAPICWrite: The guest completed a write to the virtual-APIC page that requires virtualization by VMM software.
	VmxReasonAPICWrite VmxReason = 56
	// VmxReasonCpuid: The guest software attempted to execute the CPUID instruction.
	VmxReasonCpuid VmxReason = 10
	// VmxReasonEncls: The guest attempted to execute an unsupported ENCLS instruction.
	VmxReasonEncls VmxReason = 60
	// VmxReasonEptInvept: The guest attempted to execute the Invalidate cached Extended Page Table (INVEPT) instruction.
	VmxReasonEptInvept VmxReason = 50
	// VmxReasonEptMisconfig: An attempt to access memory with a guest-physical address encountered a misconfigured Extended Page Table (EPT) paging-structure entry.
	VmxReasonEptMisconfig VmxReason = 49
	// VmxReasonEptViolation: The configuration of the Extended Page Table (EPT) paging structures disallowed an attempt to access memory with a guest-physical address.
	VmxReasonEptViolation VmxReason = 48
	// VmxReasonExcNmi: VMX exit due to an exception or non-maskable interrupt (NMI).
	VmxReasonExcNmi VmxReason = 0
	// VmxReasonGdtrIdtr: The guest attempted to execute LGDT, LIDT, SGDT, or SIDT instructions and the “descriptor-table exiting” VM-execution control was 1.
	VmxReasonGdtrIdtr VmxReason = 46
	// VmxReasonGetsec: The guest attempted to execute GETSEC instruction.
	VmxReasonGetsec VmxReason = 11
	// VmxReasonHlt: The guest attempted to execute HLT and the “HLT exiting” VM-execution control was 1.
	VmxReasonHlt VmxReason = 12
	// VmxReasonIO: Guest attempted to execute an I/O instruction.
	VmxReasonIO VmxReason = 30
	// VmxReasonIOSmi: VMX exited due to an I/O SMM Interrupt.
	VmxReasonIOSmi VmxReason = 5
	// VmxReasonInit: VMX exit due to an INIT signal.
	VmxReasonInit VmxReason = 3
	// VmxReasonInvd: The guest attempted to execute Invalidate Caches (INVD) instruction.
	VmxReasonInvd VmxReason = 13
	// VmxReasonInvlpg: The guest attempted to execute the Invalidate TLB Entry (INVLPG) instruction and the “INVLPG exiting” VM-execution control was 1.
	VmxReasonInvlpg VmxReason = 14
	// VmxReasonInvpcid: The guest attempted to execute an INVPCID instruction and the “enable INVPCID” and “INVLPG exiting” VM-execution controls were both 1.
	VmxReasonInvpcid VmxReason = 58
	// VmxReasonInvvpid: The guest attempted to execute the INVVPID instruction.
	VmxReasonInvvpid VmxReason = 53
	// VmxReasonIrq: An external interrupt arrived and the “external-interrupt exiting” VM-execution control was 1.
	VmxReasonIrq VmxReason = 1
	// VmxReasonIrqWnd: VMX exited due to an Interrupt Window.
	VmxReasonIrqWnd VmxReason = 7
	// VmxReasonLdtrTr: The guest attempted to execute LLDT, LTR, SLDT, or STR instructions and the “descriptor-table exiting” VM-execution control was 1.
	VmxReasonLdtrTr VmxReason = 47
	// VmxReasonMonitor: The guest attempted to execute MONITOR and the “MONITOR exiting” VM-execution control was 1.
	VmxReasonMonitor VmxReason = 39
	// VmxReasonMovCr: The guest attempted to access one of the CR0, CR3, CR4 or CR8 control registers.
	VmxReasonMovCr VmxReason = 28
	// VmxReasonMovDr: The guest attempted a MOV to or from a debug register and the “MOV-DR exiting” VM-execution control was 1.
	VmxReasonMovDr VmxReason = 29
	// VmxReasonMtf: VM exit occurred due to the setting of the monitor trap flag (MTF) or injection of a pending MTF VM exit.
	VmxReasonMtf VmxReason = 37
	// VmxReasonMwait: The guest attempted to execute an MWAIT instruction and the “MWAIT exiting” VM-execution control was 1.
	VmxReasonMwait VmxReason = 36
	// VmxReasonOtherSmi: An SMI arrived and caused an SMM VM exit.
	VmxReasonOtherSmi VmxReason = 6
	// VmxReasonPause: The guest attempted to execute PAUSE when the VM-execution control was 1 or exceeded the execition time window.
	VmxReasonPause   VmxReason = 40
	VmxReasonPmlFull VmxReason = 62
	// VmxReasonRdmsr: The guest attempted to execute RDMSR.
	VmxReasonRdmsr VmxReason = 31
	// VmxReasonRdpmc: The guest attempted to execute read performance monitoring counters (RDPMC) instruction and the “RDPMC exiting” VM-execution control was 1.
	VmxReasonRdpmc VmxReason = 15
	// VmxReasonRdrand: The guest software attempted to execute RDRAND instruction and the “RDRAND exiting” VM-execution control was 1.
	VmxReasonRdrand VmxReason = 57
	// VmxReasonRdseed: The guest attempted to execute RDSEED and the “RDSEED exiting” VM-execution control was 1.
	VmxReasonRdseed VmxReason = 61
	// VmxReasonRdtsc: The guest attempted to execute read time stamp counter (RDTSC) instruction and the “RDTSC exiting” VM-execution control was 1.
	VmxReasonRdtsc VmxReason = 16
	// VmxReasonRdtscp: The guest attempted to execute an RDTSCP instruction and the “enable RDTSCP” and “RDTSC exiting” VM-execution controls were both 1.
	VmxReasonRdtscp VmxReason = 51
	// VmxReasonRsm: The guest software attempted to execute a return from system management mode (RSM) instuction in system-management mode.
	VmxReasonRsm VmxReason = 17
	// VmxReasonSipi: VMS exit due to startup (IPI).
	VmxReasonSipi VmxReason = 4
	// VmxReasonSppEvent: The processor attempted to determine an access’s sub-page write permission and encountered an SPP miss or an SPP misconfiguration.
	VmxReasonSppEvent VmxReason = 66
	// VmxReasonTask: The guest attempted a task switch.
	VmxReasonTask VmxReason = 9
	// VmxReasonTpause: The guest attempted to execute a TPAUSE instuction and both the “enable user wait and pause” and “RDTSC exiting” VM-execution controls were both 1.
	VmxReasonTpause VmxReason = 68
	// VmxReasonTprThreshold: The logical processor determined that the value of the byte at offset 080H on the virtual-APIC page was below the required TPR threshold.
	VmxReasonTprThreshold VmxReason = 43
	// VmxReasonTripleFault: VMX exit due to a triple fault.
	VmxReasonTripleFault VmxReason = 2
	// VmxReasonUmwait: The guest attempted to execute a UMWAIT instruction and both the “enable user wait and pause” and “RDTSC exiting” VM-execution controls were both 1.
	VmxReasonUmwait VmxReason = 67
	// VmxReasonVirtualNmiWnd: At the beginning of an instruction, there was no virtual-NMI blocking.
	VmxReasonVirtualNmiWnd VmxReason = 8
	// VmxReasonVirtualizedEoi: The system performed EOI virtualization for a virtual interrupt whose vector indexed a bit set in the EOIexit bitmap.
	VmxReasonVirtualizedEoi VmxReason = 45
	// VmxReasonVmcall: The execution of VMCALL by either by the guest or the executive monitor casued an ordinary VM exit or an SMM VM exit, respectively.
	VmxReasonVmcall VmxReason = 18
	// VmxReasonVmclear: The guest attempted to execute VMCLEAR.
	VmxReasonVmclear VmxReason = 19
	// VmxReasonVmentryGuest: VM entry failed one of the entry checks.
	VmxReasonVmentryGuest VmxReason = 33
	// VmxReasonVmentryMc: A machine-check event occurred during VM entry.
	VmxReasonVmentryMc VmxReason = 41
	// VmxReasonVmentryMsr: A VM entry failed in an attempt to load model specific registers.
	VmxReasonVmentryMsr VmxReason = 34
	// VmxReasonVmfunc: The guest called a VM function and the VM function either wasn’t enabled or generated a function-specific condition causing a VM exit.
	VmxReasonVmfunc VmxReason = 59
	// VmxReasonVmlaunch: The guest attempted to execute VMLAUNCH.
	VmxReasonVmlaunch VmxReason = 20
	// VmxReasonVmoff: The guest attempted to execute VMXOFF.
	VmxReasonVmoff VmxReason = 26
	// VmxReasonVmon: The guest attempted to execute VMXON.
	VmxReasonVmon VmxReason = 27
	// VmxReasonVmptrld: The guest attempted to execute VMPTRLD.
	VmxReasonVmptrld VmxReason = 21
	// VmxReasonVmptrst: The guest attempted to execute VMPTRST.
	VmxReasonVmptrst VmxReason = 22
	// VmxReasonVmread: The guest attempted to execute VMREAD.
	VmxReasonVmread VmxReason = 23
	// VmxReasonVmresume: The guest attempted to execute VMRESUME.
	VmxReasonVmresume VmxReason = 24
	// VmxReasonVmwrite: The guest attempted to execute VMWRITE.
	VmxReasonVmwrite VmxReason = 25
	// VmxReasonVmxTimerExpired: The preemption timer counted down to zero.
	VmxReasonVmxTimerExpired VmxReason = 52
	// VmxReasonWbinvd: The guest attempted to execute WBINVD and the “WBINVD exiting” VM-execution control was 1.
	VmxReasonWbinvd VmxReason = 54
	// VmxReasonWrmsr: The guest attempted to execute WRMSR.
	VmxReasonWrmsr VmxReason = 32
	// VmxReasonXrstors: The guest attempted to execute XRSTORS which wasn’t allowed in the current configuration.
	VmxReasonXrstors VmxReason = 64
	// VmxReasonXsaves: The guest attempted to execute XSAVES which wasn’t allowed in the current configuration.
	VmxReasonXsaves VmxReason = 63
	// VmxReasonXsetbv: The guest attempted to execute XSETBV.
	VmxReasonXsetbv VmxReason = 55
)

func (e VmxReason) String() string {
	switch e {
	case VmxReasonAPICAccess:
		return "VmxReasonAPICAccess"
	case VmxReasonAPICWrite:
		return "VmxReasonAPICWrite"
	case VmxReasonCpuid:
		return "VmxReasonCpuid"
	case VmxReasonEncls:
		return "VmxReasonEncls"
	case VmxReasonEptInvept:
		return "VmxReasonEptInvept"
	case VmxReasonEptMisconfig:
		return "VmxReasonEptMisconfig"
	case VmxReasonEptViolation:
		return "VmxReasonEptViolation"
	case VmxReasonExcNmi:
		return "VmxReasonExcNmi"
	case VmxReasonGdtrIdtr:
		return "VmxReasonGdtrIdtr"
	case VmxReasonGetsec:
		return "VmxReasonGetsec"
	case VmxReasonHlt:
		return "VmxReasonHlt"
	case VmxReasonIO:
		return "VmxReasonIO"
	case VmxReasonIOSmi:
		return "VmxReasonIOSmi"
	case VmxReasonInit:
		return "VmxReasonInit"
	case VmxReasonInvd:
		return "VmxReasonInvd"
	case VmxReasonInvlpg:
		return "VmxReasonInvlpg"
	case VmxReasonInvpcid:
		return "VmxReasonInvpcid"
	case VmxReasonInvvpid:
		return "VmxReasonInvvpid"
	case VmxReasonIrq:
		return "VmxReasonIrq"
	case VmxReasonIrqWnd:
		return "VmxReasonIrqWnd"
	case VmxReasonLdtrTr:
		return "VmxReasonLdtrTr"
	case VmxReasonMonitor:
		return "VmxReasonMonitor"
	case VmxReasonMovCr:
		return "VmxReasonMovCr"
	case VmxReasonMovDr:
		return "VmxReasonMovDr"
	case VmxReasonMtf:
		return "VmxReasonMtf"
	case VmxReasonMwait:
		return "VmxReasonMwait"
	case VmxReasonOtherSmi:
		return "VmxReasonOtherSmi"
	case VmxReasonPause:
		return "VmxReasonPause"
	case VmxReasonPmlFull:
		return "VmxReasonPmlFull"
	case VmxReasonRdmsr:
		return "VmxReasonRdmsr"
	case VmxReasonRdpmc:
		return "VmxReasonRdpmc"
	case VmxReasonRdrand:
		return "VmxReasonRdrand"
	case VmxReasonRdseed:
		return "VmxReasonRdseed"
	case VmxReasonRdtsc:
		return "VmxReasonRdtsc"
	case VmxReasonRdtscp:
		return "VmxReasonRdtscp"
	case VmxReasonRsm:
		return "VmxReasonRsm"
	case VmxReasonSipi:
		return "VmxReasonSipi"
	case VmxReasonSppEvent:
		return "VmxReasonSppEvent"
	case VmxReasonTask:
		return "VmxReasonTask"
	case VmxReasonTpause:
		return "VmxReasonTpause"
	case VmxReasonTprThreshold:
		return "VmxReasonTprThreshold"
	case VmxReasonTripleFault:
		return "VmxReasonTripleFault"
	case VmxReasonUmwait:
		return "VmxReasonUmwait"
	case VmxReasonVirtualNmiWnd:
		return "VmxReasonVirtualNmiWnd"
	case VmxReasonVirtualizedEoi:
		return "VmxReasonVirtualizedEoi"
	case VmxReasonVmcall:
		return "VmxReasonVmcall"
	case VmxReasonVmclear:
		return "VmxReasonVmclear"
	case VmxReasonVmentryGuest:
		return "VmxReasonVmentryGuest"
	case VmxReasonVmentryMc:
		return "VmxReasonVmentryMc"
	case VmxReasonVmentryMsr:
		return "VmxReasonVmentryMsr"
	case VmxReasonVmfunc:
		return "VmxReasonVmfunc"
	case VmxReasonVmlaunch:
		return "VmxReasonVmlaunch"
	case VmxReasonVmoff:
		return "VmxReasonVmoff"
	case VmxReasonVmon:
		return "VmxReasonVmon"
	case VmxReasonVmptrld:
		return "VmxReasonVmptrld"
	case VmxReasonVmptrst:
		return "VmxReasonVmptrst"
	case VmxReasonVmread:
		return "VmxReasonVmread"
	case VmxReasonVmresume:
		return "VmxReasonVmresume"
	case VmxReasonVmwrite:
		return "VmxReasonVmwrite"
	case VmxReasonVmxTimerExpired:
		return "VmxReasonVmxTimerExpired"
	case VmxReasonWbinvd:
		return "VmxReasonWbinvd"
	case VmxReasonWrmsr:
		return "VmxReasonWrmsr"
	case VmxReasonXrstors:
		return "VmxReasonXrstors"
	case VmxReasonXsaves:
		return "VmxReasonXsaves"
	case VmxReasonXsetbv:
		return "VmxReasonXsetbv"
	default:
		return fmt.Sprintf("VmxReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Hypervisor/hv_boot_state
type HVBootState uint32

const (
	HVBsInit    HVBootState = 0
	HVBsRunning HVBootState = 2
	HVBsSipi    HVBootState = 1
)

func (e HVBootState) String() string {
	switch e {
	case HVBsInit:
		return "HVBsInit"
	case HVBsRunning:
		return "HVBsRunning"
	case HVBsSipi:
		return "HVBsSipi"
	default:
		return fmt.Sprintf("HVBootState(%d)", e)
	}
}
