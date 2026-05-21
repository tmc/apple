//go:build darwin && arm64

package arm64

import (
	"fmt"

	"github.com/tmc/apple/hypervisor"
)

const (
	ExceptionClassDataAbortLowerEL = 0x24
	ExceptionClassDataAbortSameEL  = 0x25
	ExceptionClassHVC64            = 0x16
	ExceptionClassSystemReg        = 0x18
)

// RegInfo names a Hypervisor.framework register used by state or trap logic.
type RegInfo[T any] struct {
	Name     string
	Reg      T
	ReadOnly bool
}

// DataAbortAccess describes the access encoded in a data-abort syndrome.
type DataAbortAccess struct {
	Size  int
	Reg   hypervisor.HVReg
	Write bool
}

// SystemRegAccess describes the access encoded in a system-register syndrome.
type SystemRegAccess struct {
	Op0    uint64
	Op1    uint64
	CRn    uint64
	CRm    uint64
	Op2    uint64
	Rt     hypervisor.HVReg
	SysReg hypervisor.HVSysReg
	Name   string
	Read   bool
	Write  bool
}

var coreRegs = []RegInfo[hypervisor.HVReg]{
	{"x0", hypervisor.HVRegX0, false},
	{"x1", hypervisor.HVRegX1, false},
	{"x2", hypervisor.HVRegX2, false},
	{"x3", hypervisor.HVRegX3, false},
	{"x4", hypervisor.HVRegX4, false},
	{"x5", hypervisor.HVRegX5, false},
	{"x6", hypervisor.HVRegX6, false},
	{"x7", hypervisor.HVRegX7, false},
	{"x8", hypervisor.HVRegX8, false},
	{"x9", hypervisor.HVRegX9, false},
	{"x10", hypervisor.HVRegX10, false},
	{"x11", hypervisor.HVRegX11, false},
	{"x12", hypervisor.HVRegX12, false},
	{"x13", hypervisor.HVRegX13, false},
	{"x14", hypervisor.HVRegX14, false},
	{"x15", hypervisor.HVRegX15, false},
	{"x16", hypervisor.HVRegX16, false},
	{"x17", hypervisor.HVRegX17, false},
	{"x18", hypervisor.HVRegX18, false},
	{"x19", hypervisor.HVRegX19, false},
	{"x20", hypervisor.HVRegX20, false},
	{"x21", hypervisor.HVRegX21, false},
	{"x22", hypervisor.HVRegX22, false},
	{"x23", hypervisor.HVRegX23, false},
	{"x24", hypervisor.HVRegX24, false},
	{"x25", hypervisor.HVRegX25, false},
	{"x26", hypervisor.HVRegX26, false},
	{"x27", hypervisor.HVRegX27, false},
	{"x28", hypervisor.HVRegX28, false},
	{"x29", hypervisor.HVRegX29, false},
	{"x30", hypervisor.HVRegX30, false},
	{"pc", hypervisor.HVRegPC, false},
	{"pstate", hypervisor.HVRegCpsr, false},
	{"fpcr", hypervisor.HVRegFpcr, false},
	{"fpsr", hypervisor.HVRegFpsr, false},
}

var sysRegs = []RegInfo[hypervisor.HVSysReg]{
	{"mpidr_el1", hypervisor.HVSysRegMpidrEl1, true},
	{"sctlr_el1", hypervisor.HVSysRegSctlrEl1, false},
	{"cpacr_el1", hypervisor.HVSysRegCpacrEl1, false},
	{"apiakeylo_el1", hypervisor.HVSysRegApiakeyloEl1, false},
	{"apiakeyhi_el1", hypervisor.HVSysRegApiakeyhiEl1, false},
	{"apibkeylo_el1", hypervisor.HVSysRegApibkeyloEl1, false},
	{"apibkeyhi_el1", hypervisor.HVSysRegApibkeyhiEl1, false},
	{"apdakeylo_el1", hypervisor.HVSysRegApdakeyloEl1, false},
	{"apdakeyhi_el1", hypervisor.HVSysRegApdakeyhiEl1, false},
	{"apdbkeylo_el1", hypervisor.HVSysRegApdbkeyloEl1, false},
	{"apdbkeyhi_el1", hypervisor.HVSysRegApdbkeyhiEl1, false},
	{"apgakeylo_el1", hypervisor.HVSysRegApgakeyloEl1, false},
	{"apgakeyhi_el1", hypervisor.HVSysRegApgakeyhiEl1, false},
	{"ttbr0_el1", hypervisor.HVSysRegTtbr0El1, false},
	{"ttbr1_el1", hypervisor.HVSysRegTtbr1El1, false},
	{"tcr_el1", hypervisor.HVSysRegTcrEl1, false},
	{"cntkctl_el1", hypervisor.HVSysRegCntkctlEl1, false},
	{"esr_el1", hypervisor.HVSysRegEsrEl1, false},
	{"far_el1", hypervisor.HVSysRegFarEl1, false},
	{"mair_el1", hypervisor.HVSysRegMairEl1, false},
	{"amair_el1", hypervisor.HVSysRegAmairEl1, false},
	{"vbar_el1", hypervisor.HVSysRegVbarEl1, false},
	{"contextidr_el1", hypervisor.HVSysRegContextidrEl1, false},
	{"sp_el0", hypervisor.HVSysRegSpEl0, false},
	{"sp_el1", hypervisor.HVSysRegSpEl1, false},
	{"elr_el1", hypervisor.HVSysRegElrEl1, false},
	{"spsr_el1", hypervisor.HVSysRegSpsrEl1, false},
	{"tpidr_el0", hypervisor.HVSysRegTpidrEl0, false},
	{"tpidrro_el0", hypervisor.HVSysRegTpidrroEl0, false},
	{"tpidr_el1", hypervisor.HVSysRegTpidrEl1, false},
}

var timerRegs = []RegInfo[hypervisor.HVSysReg]{
	{"cntv_ctl_el0", hypervisor.HVSysRegCntvCtlEl0, false},
	{"cntv_cval_el0", hypervisor.HVSysRegCntvCvalEl0, false},
	{"cntp_ctl_el0", hypervisor.HVSysRegCntpCtlEl0, false},
	{"cntp_cval_el0", hypervisor.HVSysRegCntpCvalEl0, false},
}

var simdFPRegs = []RegInfo[hypervisor.HVSIMDFPReg]{
	{"q0", hypervisor.HVSIMDFPRegQ0, false},
	{"q1", hypervisor.HVSIMDFPRegQ1, false},
	{"q2", hypervisor.HVSIMDFPRegQ2, false},
	{"q3", hypervisor.HVSIMDFPRegQ3, false},
	{"q4", hypervisor.HVSIMDFPRegQ4, false},
	{"q5", hypervisor.HVSIMDFPRegQ5, false},
	{"q6", hypervisor.HVSIMDFPRegQ6, false},
	{"q7", hypervisor.HVSIMDFPRegQ7, false},
	{"q8", hypervisor.HVSIMDFPRegQ8, false},
	{"q9", hypervisor.HVSIMDFPRegQ9, false},
	{"q10", hypervisor.HVSIMDFPRegQ10, false},
	{"q11", hypervisor.HVSIMDFPRegQ11, false},
	{"q12", hypervisor.HVSIMDFPRegQ12, false},
	{"q13", hypervisor.HVSIMDFPRegQ13, false},
	{"q14", hypervisor.HVSIMDFPRegQ14, false},
	{"q15", hypervisor.HVSIMDFPRegQ15, false},
	{"q16", hypervisor.HVSIMDFPRegQ16, false},
	{"q17", hypervisor.HVSIMDFPRegQ17, false},
	{"q18", hypervisor.HVSIMDFPRegQ18, false},
	{"q19", hypervisor.HVSIMDFPRegQ19, false},
	{"q20", hypervisor.HVSIMDFPRegQ20, false},
	{"q21", hypervisor.HVSIMDFPRegQ21, false},
	{"q22", hypervisor.HVSIMDFPRegQ22, false},
	{"q23", hypervisor.HVSIMDFPRegQ23, false},
	{"q24", hypervisor.HVSIMDFPRegQ24, false},
	{"q25", hypervisor.HVSIMDFPRegQ25, false},
	{"q26", hypervisor.HVSIMDFPRegQ26, false},
	{"q27", hypervisor.HVSIMDFPRegQ27, false},
	{"q28", hypervisor.HVSIMDFPRegQ28, false},
	{"q29", hypervisor.HVSIMDFPRegQ29, false},
	{"q30", hypervisor.HVSIMDFPRegQ30, false},
	{"q31", hypervisor.HVSIMDFPRegQ31, false},
}

var gicDistributorRegs = []RegInfo[hypervisor.HVGICDistributorReg]{
	{"gicd_ctlr", hypervisor.HVGICDistributorRegGICDCtlr, false},
	{"gicd_igroupr0", hypervisor.HVGICDistributorRegGICDIgroupr0, false},
	{"gicd_igroupr1", hypervisor.HVGICDistributorRegGICDIgroupr1, false},
	{"gicd_isenabler0", hypervisor.HVGICDistributorRegGICDIsenabler0, false},
	{"gicd_isenabler1", hypervisor.HVGICDistributorRegGICDIsenabler1, false},
	{"gicd_ispendr0", hypervisor.HVGICDistributorRegGICDIspendr0, false},
	{"gicd_ispendr1", hypervisor.HVGICDistributorRegGICDIspendr1, false},
	{"gicd_isactiver0", hypervisor.HVGICDistributorRegGICDIsactiver0, false},
	{"gicd_isactiver1", hypervisor.HVGICDistributorRegGICDIsactiver1, false},
	{"gicd_ipriorityr0", hypervisor.HVGICDistributorRegGICDIpriorityr0, false},
	{"gicd_ipriorityr1", hypervisor.HVGICDistributorRegGICDIpriorityr1, false},
	{"gicd_ipriorityr2", hypervisor.HVGICDistributorRegGICDIpriorityr2, false},
	{"gicd_ipriorityr3", hypervisor.HVGICDistributorRegGICDIpriorityr3, false},
	{"gicd_ipriorityr4", hypervisor.HVGICDistributorRegGICDIpriorityr4, false},
	{"gicd_ipriorityr5", hypervisor.HVGICDistributorRegGICDIpriorityr5, false},
	{"gicd_ipriorityr6", hypervisor.HVGICDistributorRegGICDIpriorityr6, false},
	{"gicd_ipriorityr7", hypervisor.HVGICDistributorRegGICDIpriorityr7, false},
	{"gicd_ipriorityr8", hypervisor.HVGICDistributorRegGICDIpriorityr8, false},
	{"gicd_ipriorityr9", hypervisor.HVGICDistributorRegGICDIpriorityr9, false},
	{"gicd_ipriorityr10", hypervisor.HVGICDistributorRegGICDIpriorityr10, false},
	{"gicd_ipriorityr11", hypervisor.HVGICDistributorRegGICDIpriorityr11, false},
	{"gicd_ipriorityr12", hypervisor.HVGICDistributorRegGICDIpriorityr12, false},
	{"gicd_ipriorityr13", hypervisor.HVGICDistributorRegGICDIpriorityr13, false},
	{"gicd_ipriorityr14", hypervisor.HVGICDistributorRegGICDIpriorityr14, false},
	{"gicd_ipriorityr15", hypervisor.HVGICDistributorRegGICDIpriorityr15, false},
	{"gicd_icfgr0", hypervisor.HVGICDistributorRegGICDIcfgr0, false},
	{"gicd_icfgr1", hypervisor.HVGICDistributorRegGICDIcfgr1, false},
	{"gicd_icfgr2", hypervisor.HVGICDistributorRegGICDIcfgr2, false},
	{"gicd_icfgr3", hypervisor.HVGICDistributorRegGICDIcfgr3, false},
}

var gicRedistributorRegs = []RegInfo[hypervisor.HVGICRedistributorReg]{
	{"gicr_igroupr0", hypervisor.HVGICRedistributorRegGICRIgroupr0, false},
	{"gicr_isenabler0", hypervisor.HVGICRedistributorRegGICRIsenabler0, false},
	{"gicr_ispendr0", hypervisor.HVGICRedistributorRegGICRIspendr0, false},
	{"gicr_isactiver0", hypervisor.HVGICRedistributorRegGICRIsactiver0, false},
	{"gicr_ipriorityr0", hypervisor.HVGICRedistributorRegGICRIpriorityr0, false},
	{"gicr_ipriorityr1", hypervisor.HVGICRedistributorRegGICRIpriorityr1, false},
	{"gicr_ipriorityr2", hypervisor.HVGICRedistributorRegGICRIpriorityr2, false},
	{"gicr_ipriorityr3", hypervisor.HVGICRedistributorRegGICRIpriorityr3, false},
	{"gicr_ipriorityr4", hypervisor.HVGICRedistributorRegGICRIpriorityr4, false},
	{"gicr_ipriorityr5", hypervisor.HVGICRedistributorRegGICRIpriorityr5, false},
	{"gicr_ipriorityr6", hypervisor.HVGICRedistributorRegGICRIpriorityr6, false},
	{"gicr_ipriorityr7", hypervisor.HVGICRedistributorRegGICRIpriorityr7, false},
	{"gicr_icfgr0", hypervisor.HVGICRedistributorRegGICRIcfgr0, false},
	{"gicr_icfgr1", hypervisor.HVGICRedistributorRegGICRIcfgr1, false},
}

var gicICCRegs = []RegInfo[hypervisor.HVGICIccReg]{
	{"icc_pmr_el1", hypervisor.HVGICIccRegPmrEl1, false},
	{"icc_bpr0_el1", hypervisor.HVGICIccRegBpr0El1, false},
	{"icc_bpr1_el1", hypervisor.HVGICIccRegBpr1El1, false},
	{"icc_ctlr_el1", hypervisor.HVGICIccRegCtlrEl1, false},
	{"icc_sre_el1", hypervisor.HVGICIccRegSreEl1, false},
	{"icc_ap0r0_el1", hypervisor.HVGICIccRegAp0r0El1, false},
	{"icc_ap1r0_el1", hypervisor.HVGICIccRegAp1r0El1, false},
	{"icc_igrpen0_el1", hypervisor.HVGICIccRegIgrpen0El1, false},
	{"icc_igrpen1_el1", hypervisor.HVGICIccRegIgrpen1El1, false},
}

var gicICHRegs = []RegInfo[hypervisor.HVGICIchReg]{
	{"ich_ap0r0_el2", hypervisor.HVGICIchRegAp0r0El2, false},
	{"ich_ap1r0_el2", hypervisor.HVGICIchRegAp1r0El2, false},
	{"ich_eisr_el2", hypervisor.HVGICIchRegEisrEl2, false},
	{"ich_elrsr_el2", hypervisor.HVGICIchRegElrsrEl2, false},
	{"ich_hcr_el2", hypervisor.HVGICIchRegHcrEl2, false},
	{"ich_lr0_el2", hypervisor.HVGICIchRegLr0El2, false},
	{"ich_lr1_el2", hypervisor.HVGICIchRegLr1El2, false},
	{"ich_lr2_el2", hypervisor.HVGICIchRegLr2El2, false},
	{"ich_lr3_el2", hypervisor.HVGICIchRegLr3El2, false},
	{"ich_lr4_el2", hypervisor.HVGICIchRegLr4El2, false},
	{"ich_lr5_el2", hypervisor.HVGICIchRegLr5El2, false},
	{"ich_lr6_el2", hypervisor.HVGICIchRegLr6El2, false},
	{"ich_lr7_el2", hypervisor.HVGICIchRegLr7El2, false},
	{"ich_lr8_el2", hypervisor.HVGICIchRegLr8El2, false},
	{"ich_lr9_el2", hypervisor.HVGICIchRegLr9El2, false},
	{"ich_lr10_el2", hypervisor.HVGICIchRegLr10El2, false},
	{"ich_lr11_el2", hypervisor.HVGICIchRegLr11El2, false},
	{"ich_lr12_el2", hypervisor.HVGICIchRegLr12El2, false},
	{"ich_lr13_el2", hypervisor.HVGICIchRegLr13El2, false},
	{"ich_lr14_el2", hypervisor.HVGICIchRegLr14El2, false},
	{"ich_lr15_el2", hypervisor.HVGICIchRegLr15El2, false},
	{"ich_misr_el2", hypervisor.HVGICIchRegMisrEl2, false},
	{"ich_vmcr_el2", hypervisor.HVGICIchRegVmcrEl2, false},
	{"ich_vtr_el2", hypervisor.HVGICIchRegVtrEl2, false},
}

var el2SysRegs = []RegInfo[hypervisor.HVSysReg]{
	{"sctlr_el2", hypervisor.HVSysRegSctlrEl2, false},
	{"hcr_el2", hypervisor.HVSysRegHcrEl2, false},
	{"cptr_el2", hypervisor.HVSysRegCptrEl2, false},
	{"cnthctl_el2", hypervisor.HVSysRegCnthctlEl2, false},
	{"cntvoff_el2", hypervisor.HVSysRegCntvoffEl2, false},
	{"elr_el2", hypervisor.HVSysRegElrEl2, false},
	{"esr_el2", hypervisor.HVSysRegEsrEl2, false},
	{"far_el2", hypervisor.HVSysRegFarEl2, false},
	{"mair_el2", hypervisor.HVSysRegMairEl2, false},
	{"mdcr_el2", hypervisor.HVSysRegMdcrEl2, false},
	{"sp_el2", hypervisor.HVSysRegSpEl2, false},
	{"spsr_el2", hypervisor.HVSysRegSpsrEl2, false},
	{"tcr_el2", hypervisor.HVSysRegTcrEl2, false},
	{"tpidr_el2", hypervisor.HVSysRegTpidrEl2, false},
	{"ttbr0_el2", hypervisor.HVSysRegTtbr0El2, false},
	{"ttbr1_el2", hypervisor.HVSysRegTtbr1El2, false},
	{"vbar_el2", hypervisor.HVSysRegVbarEl2, false},
	{"vmpidr_el2", hypervisor.HVSysRegVmpidrEl2, false},
	{"vpidr_el2", hypervisor.HVSysRegVpidrEl2, false},
	{"vttbr_el2", hypervisor.HVSysRegVttbrEl2, false},
}

func cloneRegs[T any](regs []RegInfo[T]) []RegInfo[T] {
	return append([]RegInfo[T](nil), regs...)
}

// CoreRegs returns the named general-purpose register set.
func CoreRegs() []RegInfo[hypervisor.HVReg] { return cloneRegs(coreRegs) }

// SysRegs returns the named EL1 system-register set.
func SysRegs() []RegInfo[hypervisor.HVSysReg] { return cloneRegs(sysRegs) }

// TimerRegs returns the named timer register set.
func TimerRegs() []RegInfo[hypervisor.HVSysReg] { return cloneRegs(timerRegs) }

// SIMDFPRegs returns the named SIMD/FP register set.
func SIMDFPRegs() []RegInfo[hypervisor.HVSIMDFPReg] { return cloneRegs(simdFPRegs) }

// GICDistributorRegs returns the named GIC distributor register set.
func GICDistributorRegs() []RegInfo[hypervisor.HVGICDistributorReg] {
	return cloneRegs(gicDistributorRegs)
}

// GICRedistributorRegs returns the named GIC redistributor register set.
func GICRedistributorRegs() []RegInfo[hypervisor.HVGICRedistributorReg] {
	return cloneRegs(gicRedistributorRegs)
}

// GICICCRegs returns the named GIC CPU-interface register set.
func GICICCRegs() []RegInfo[hypervisor.HVGICIccReg] { return cloneRegs(gicICCRegs) }

// GICICHRegs returns the named GIC hypervisor-control register set.
func GICICHRegs() []RegInfo[hypervisor.HVGICIchReg] { return cloneRegs(gicICHRegs) }

// EL2SysRegs returns the named EL2 system-register set.
func EL2SysRegs() []RegInfo[hypervisor.HVSysReg] { return cloneRegs(el2SysRegs) }

// LookupCoreReg returns the general-purpose register with name.
func LookupCoreReg(name string) (hypervisor.HVReg, bool) {
	for _, reg := range coreRegs {
		if reg.Name == name {
			return reg.Reg, true
		}
	}
	return 0, false
}

// RegName returns a stable name for a general-purpose register.
func RegName(reg hypervisor.HVReg) string {
	for _, item := range coreRegs {
		if item.Reg == reg {
			return item.Name
		}
	}
	return fmt.Sprintf("reg_%#x", uint64(reg))
}

// LookupSysReg returns the writable EL1 system register with name.
func LookupSysReg(name string) (hypervisor.HVSysReg, bool) {
	for _, reg := range sysRegs {
		if reg.Name == name && !reg.ReadOnly {
			return reg.Reg, true
		}
	}
	return 0, false
}

// LookupTimerReg returns the timer register with name.
func LookupTimerReg(name string) (hypervisor.HVSysReg, bool) {
	for _, reg := range timerRegs {
		if reg.Name == name {
			return reg.Reg, true
		}
	}
	return 0, false
}

// LookupSIMDFPReg returns the SIMD/FP register with name.
func LookupSIMDFPReg(name string) (hypervisor.HVSIMDFPReg, bool) {
	for _, reg := range simdFPRegs {
		if reg.Name == name {
			return reg.Reg, true
		}
	}
	return 0, false
}

// LookupGICDistributorReg returns the GIC distributor register with name.
func LookupGICDistributorReg(name string) (hypervisor.HVGICDistributorReg, bool) {
	for _, reg := range gicDistributorRegs {
		if reg.Name == name {
			return reg.Reg, true
		}
	}
	return 0, false
}

// LookupGICRedistributorReg returns the GIC redistributor register with name.
func LookupGICRedistributorReg(name string) (hypervisor.HVGICRedistributorReg, bool) {
	for _, reg := range gicRedistributorRegs {
		if reg.Name == name {
			return reg.Reg, true
		}
	}
	return 0, false
}

// LookupGICICCReg returns the GIC CPU-interface register with name.
func LookupGICICCReg(name string) (hypervisor.HVGICIccReg, bool) {
	for _, reg := range gicICCRegs {
		if reg.Name == name {
			return reg.Reg, true
		}
	}
	return 0, false
}

// LookupGICICHReg returns the GIC hypervisor-control register encoded as sysreg.
func LookupGICICHReg(sysreg hypervisor.HVSysReg) (hypervisor.HVGICIchReg, bool) {
	for _, reg := range gicICHRegs {
		if hypervisor.HVSysReg(reg.Reg) == sysreg {
			return reg.Reg, true
		}
	}
	return 0, false
}

// SysRegEncoding encodes a system-register tuple as a Hypervisor.framework register.
func SysRegEncoding(op0, op1, crn, crm, op2 uint64) hypervisor.HVSysReg {
	return hypervisor.HVSysReg(op0<<14 | op1<<11 | crn<<7 | crm<<3 | op2)
}

// EmulatedSysReg reports whether reg is named by the register tables.
func EmulatedSysReg(reg hypervisor.HVSysReg) (hypervisor.HVSysReg, bool) {
	for _, item := range sysRegs {
		if item.Reg == reg {
			return item.Reg, true
		}
	}
	for _, item := range timerRegs {
		if item.Reg == reg {
			return item.Reg, true
		}
	}
	for _, item := range el2SysRegs {
		if item.Reg == reg {
			return item.Reg, true
		}
	}
	return 0, false
}

// SysRegEncodingName returns a stable name for a system register encoding.
func SysRegEncodingName(reg hypervisor.HVSysReg) string {
	for _, item := range sysRegs {
		if item.Reg == reg {
			return item.Name
		}
	}
	for _, item := range timerRegs {
		if item.Reg == reg {
			return item.Name
		}
	}
	for _, item := range el2SysRegs {
		if item.Reg == reg {
			return item.Name
		}
	}
	for _, item := range gicICHRegs {
		if hypervisor.HVSysReg(item.Reg) == reg {
			return item.Name
		}
	}
	return fmt.Sprintf("sysreg_%#x", uint64(reg))
}

// SyndromeEC returns the exception class encoded in syndrome.
func SyndromeEC(syndrome uint64) uint64 {
	return (syndrome >> 26) & 0x3f
}

// SyndromeIsDataAbort reports whether syndrome is a data-abort exception.
func SyndromeIsDataAbort(syndrome uint64) bool {
	ec := SyndromeEC(syndrome)
	return ec == ExceptionClassDataAbortLowerEL || ec == ExceptionClassDataAbortSameEL
}

// DecodeSystemReg decodes a trapped system-register access syndrome.
func DecodeSystemReg(syndrome uint64) (SystemRegAccess, bool) {
	if SyndromeEC(syndrome) != ExceptionClassSystemReg {
		return SystemRegAccess{}, false
	}
	iss := syndrome & 0x1ffffff
	access := SystemRegAccess{
		Op0:   (iss >> 20) & 0x3,
		Op1:   (iss >> 14) & 0x7,
		CRn:   (iss >> 10) & 0xf,
		CRm:   (iss >> 1) & 0xf,
		Op2:   (iss >> 17) & 0x7,
		Rt:    hypervisor.HVReg((iss >> 5) & 0x1f),
		Read:  iss&1 != 0,
		Write: iss&1 == 0,
	}
	access.SysReg = SysRegEncoding(access.Op0, access.Op1, access.CRn, access.CRm, access.Op2)
	access.Name = SysRegEncodingName(access.SysReg)
	return access, true
}

// DecodeDataAbort decodes a data-abort access syndrome.
func DecodeDataAbort(syndrome uint64) (DataAbortAccess, bool) {
	if !SyndromeIsDataAbort(syndrome) || syndrome&(1<<24) == 0 {
		return DataAbortAccess{}, false
	}
	size := 1 << ((syndrome >> 22) & 0x3)
	reg := hypervisor.HVReg((syndrome >> 16) & 0x1f)
	return DataAbortAccess{
		Size:  int(size),
		Reg:   reg,
		Write: syndrome&(1<<6) != 0,
	}, true
}
