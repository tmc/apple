//go:build darwin && arm64

package rdma

import (
	"os/exec"
	"path/filepath"
	"testing"
	"unsafe"
)

// TestGoTypesMatchAppleABI compares the Go declarations against the C records
// they claim to match.
//
// TestAppleHeaderABI measures the same records, but compares them to constants
// written beside it, so it proves only that the SDK has not moved. It cannot
// fail because of anything in this package: every Go type here could be the
// wrong size and it would still pass. This is the half that reads the Go side.
//
// The C values are taken from the oracle rather than repeated here, so there is
// one place a number can be wrong and it is the compiler's.
func TestGoTypesMatchAppleABI(t *testing.T) {
	oracle := runABIOracle(t)

	sizes := map[string]uintptr{
		"union ibv_gid":           unsafe.Sizeof(IbvGID{}),
		"struct ibv_qp_cap":       unsafe.Sizeof(IbvQPCap{}),
		"struct ibv_qp_init_attr": unsafe.Sizeof(IbvQPInitAttr{}),
		"struct ibv_global_route": unsafe.Sizeof(IbvGlobalRoute{}),
		"struct ibv_ah_attr":      unsafe.Sizeof(IbvAHAttr{}),
		"struct ibv_qp_attr":      unsafe.Sizeof(IbvQPAttr{}),
		"struct ibv_sge":          unsafe.Sizeof(IbvSGE{}),
		"struct ibv_send_wr":      unsafe.Sizeof(IbvSendWR{}),
		"struct ibv_recv_wr":      unsafe.Sizeof(IbvRecvWR{}),
		"struct ibv_wc":           unsafe.Sizeof(IbvWC{}),
		"struct ibv_port_attr":    unsafe.Sizeof(IbvPortAttr{}),
	}
	for cName, got := range sizes {
		want, ok := oracle[cName+".size"]
		if !ok {
			t.Errorf("ABI oracle does not measure %s", cName)
			continue
		}
		if got != want {
			t.Errorf("sizeof(%s) = %d in Go, %d in C", cName, got, want)
		}
	}

	var (
		initAttr IbvQPInitAttr
		ahAttr   IbvAHAttr
		qpAttr   IbvQPAttr
		sendWR   IbvSendWR
		portAttr IbvPortAttr
	)
	offsets := map[string]uintptr{
		"struct ibv_qp_init_attr.send_cq":      unsafe.Offsetof(initAttr.SendCQ),
		"struct ibv_qp_init_attr.cap":          unsafe.Offsetof(initAttr.Cap),
		"struct ibv_qp_init_attr.qp_type":      unsafe.Offsetof(initAttr.QPType),
		"struct ibv_ah_attr.dlid":              unsafe.Offsetof(ahAttr.DLID),
		"struct ibv_qp_attr.cap":               unsafe.Offsetof(qpAttr.Cap),
		"struct ibv_qp_attr.ah_attr":           unsafe.Offsetof(qpAttr.AHAttr),
		"struct ibv_qp_attr.alt_ah_attr":       unsafe.Offsetof(qpAttr.AltAHAttr),
		"struct ibv_qp_attr.pkey_index":        unsafe.Offsetof(qpAttr.PKeyIndex),
		"struct ibv_qp_attr.port_num":          unsafe.Offsetof(qpAttr.PortNum),
		"struct ibv_send_wr.wr":                unsafe.Offsetof(sendWR.WR),
		"struct ibv_port_attr.active_mtu":      unsafe.Offsetof(portAttr.ActiveMTU),
		"struct ibv_port_attr.gid_tbl_len":     unsafe.Offsetof(portAttr.GIDTblLen),
		"struct ibv_port_attr.lid":             unsafe.Offsetof(portAttr.LID),
		"struct ibv_port_attr.link_layer":      unsafe.Offsetof(portAttr.LinkLayer),
		"struct ibv_port_attr.port_cap_flags2": unsafe.Offsetof(portAttr.PortCapFlags2),
		// These three are not struct fields on the Go side: the opaque handles
		// are uintptr, and the accessors reach into them by constant. A wrong
		// constant here is a wrong read of live memory, so the oracle checks it
		// exactly as it checks a field.
		"struct ibv_qp.qp_num": ibvQPNumOffset,
		"struct ibv_mr.lkey":   ibvMRLKeyOffset,
		"struct ibv_mr.rkey":   ibvMRRKeyOffset,
	}
	for name, got := range offsets {
		want, ok := oracle[name]
		if !ok {
			t.Errorf("ABI oracle does not measure %s", name)
			continue
		}
		if got != want {
			t.Errorf("offsetof(%s) = %d in Go, %d in C", name, got, want)
		}
	}
}

// runABIOracle compiles and runs the C program that reports the SDK's own
// measurements, and returns them keyed the way it prints them.
func runABIOracle(t *testing.T) map[string]uintptr {
	t.Helper()
	clang := xcrun(t, "--sdk", "macosx", "--find", "clang")
	sdk := xcrun(t, "--sdk", "macosx", "--show-sdk-path")
	exe := filepath.Join(t.TempDir(), "abi-oracle")
	cmd := exec.Command(clang, "-isysroot", sdk, "-std=c11", "-Werror", "-o", exe, filepath.Join("testdata", "abi_oracle.c"))
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("compile ABI oracle: %v\n%s", err, out)
	}
	out, err := exec.Command(exe).Output()
	if err != nil {
		t.Fatalf("run ABI oracle: %v", err)
	}
	return parseABIOracle(t, out)
}
