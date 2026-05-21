//go:build darwin && arm64

package arm64

import (
	"github.com/tmc/apple/hypervisor"
	"testing"
)

func TestRegisterLookups(t *testing.T) {
	tests := []struct {
		name string
		ok   bool
	}{
		{"core x0", hasCoreReg("x0")},
		{"core pstate", hasCoreReg("pstate")},
		{"sys sctlr_el1", hasSysReg("sctlr_el1")},
		{"timer cntv_ctl_el0", hasTimerReg("cntv_ctl_el0")},
		{"simd q31", hasSIMDFPReg("q31")},
		{"dist gicd_ctlr", hasGICDistributorReg("gicd_ctlr")},
		{"redist gicr_isenabler0", hasGICRedistributorReg("gicr_isenabler0")},
		{"icc icc_pmr_el1", hasGICICCReg("icc_pmr_el1")},
	}
	for _, tt := range tests {
		if !tt.ok {
			t.Fatalf("%s missing", tt.name)
		}
	}
	if _, ok := LookupSysReg("mpidr_el1"); ok {
		t.Fatal("LookupSysReg accepted read-only mpidr_el1")
	}
}

func TestRegisterSlicesAreCopies(t *testing.T) {
	regs := CoreRegs()
	regs[0].Name = "changed"
	regs = CoreRegs()
	if regs[0].Name != "x0" {
		t.Fatalf("CoreRegs()[0].Name = %q, want x0", regs[0].Name)
	}
}

func TestDecodeSystemReg(t *testing.T) {
	access, ok := DecodeSystemReg(0x62313802)
	if !ok {
		t.Fatal("DecodeSystemReg = false")
	}
	if !access.Write || access.Read {
		t.Fatalf("access read=%v write=%v, want write", access.Read, access.Write)
	}
	if access.Rt != hypervisor.HVRegX0 {
		t.Fatalf("access Rt = %s, want %s", RegName(access.Rt), RegName(hypervisor.HVRegX0))
	}
	if access.SysReg != hypervisor.HVSysRegCnthctlEl2 || access.Name != "cnthctl_el2" {
		t.Fatalf("access sysreg = %s/%#x, want cnthctl_el2/%#x", access.Name, access.SysReg, hypervisor.HVSysRegCnthctlEl2)
	}

	lorr := SysRegEncoding(3, 0, 10, 4, 3)
	if _, ok := EmulatedSysReg(lorr); ok {
		t.Fatal("LORC_EL1 unexpectedly emulated")
	}
	ich := SysRegEncoding(3, 4, 12, 11, 0)
	if reg, ok := LookupGICICHReg(ich); !ok || reg != hypervisor.HVGICIchRegHcrEl2 {
		t.Fatalf("LookupGICICHReg(ICH_HCR_EL2) = %#x, %v", reg, ok)
	}
	if name := SysRegEncodingName(ich); name != "ich_hcr_el2" {
		t.Fatalf("SysRegEncodingName(ICH_HCR_EL2) = %q", name)
	}
}

func TestDecodeDataAbort(t *testing.T) {
	syndrome := uint64(ExceptionClassDataAbortLowerEL<<26) | (1 << 24) | (2 << 22) | (5 << 16) | (1 << 6)
	got, ok := DecodeDataAbort(syndrome)
	if !ok {
		t.Fatal("DecodeDataAbort returned ok=false")
	}
	if got.Size != 4 || got.Reg != hypervisor.HVRegX5 || !got.Write {
		t.Fatalf("DecodeDataAbort = %+v", got)
	}
	if SyndromeIsDataAbort(ExceptionClassHVC64 << 26) {
		t.Fatal("SyndromeIsDataAbort accepted non-data-abort syndrome")
	}
}

func hasCoreReg(name string) bool {
	_, ok := LookupCoreReg(name)
	return ok
}

func hasSysReg(name string) bool {
	_, ok := LookupSysReg(name)
	return ok
}

func hasTimerReg(name string) bool {
	_, ok := LookupTimerReg(name)
	return ok
}

func hasSIMDFPReg(name string) bool {
	_, ok := LookupSIMDFPReg(name)
	return ok
}

func hasGICDistributorReg(name string) bool {
	_, ok := LookupGICDistributorReg(name)
	return ok
}

func hasGICRedistributorReg(name string) bool {
	_, ok := LookupGICRedistributorReg(name)
	return ok
}

func hasGICICCReg(name string) bool {
	_, ok := LookupGICICCReg(name)
	return ok
}
