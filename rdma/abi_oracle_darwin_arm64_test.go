//go:build darwin && arm64

package rdma

import (
	"bufio"
	"bytes"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestAppleHeaderABI(t *testing.T) {
	testdata := filepath.Join("testdata", "abi_oracle.c")
	clang := xcrun(t, "--sdk", "macosx", "--find", "clang")
	sdk := xcrun(t, "--sdk", "macosx", "--show-sdk-path")
	exe := filepath.Join(t.TempDir(), "abi-oracle")
	cmd := exec.Command(clang, "-isysroot", sdk, "-std=c11", "-Werror", "-o", exe, testdata)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("compile ABI oracle: %v\n%s", err, out)
	}
	out, err := exec.Command(exe).Output()
	if err != nil {
		t.Fatalf("run ABI oracle: %v", err)
	}
	got := parseABIOracle(t, out)
	for name, want := range map[string]uintptr{
		"union ibv_gid.size":                   16,
		"struct ibv_qp_cap.size":               20,
		"struct ibv_qp_init_attr.size":         64,
		"struct ibv_global_route.size":         24,
		"struct ibv_ah_attr.size":              32,
		"struct ibv_qp_attr.size":              144,
		"struct ibv_sge.size":                  16,
		"struct ibv_send_wr.size":              128,
		"struct ibv_recv_wr.size":              32,
		"struct ibv_wc.size":                   48,
		"struct ibv_port_attr.size":            52,
		"struct ibv_qp_init_attr.send_cq":      8,
		"struct ibv_qp_init_attr.cap":          32,
		"struct ibv_qp_init_attr.qp_type":      52,
		"struct ibv_ah_attr.dlid":              24,
		"struct ibv_qp_attr.cap":               36,
		"struct ibv_qp_attr.ah_attr":           56,
		"struct ibv_qp_attr.alt_ah_attr":       88,
		"struct ibv_qp_attr.pkey_index":        120,
		"struct ibv_qp_attr.port_num":          129,
		"struct ibv_qp_attr.rate_limit":        136,
		"struct ibv_send_wr.wr":                40,
		"struct ibv_port_attr.gid_tbl_len":     12,
		"struct ibv_port_attr.lid":             34,
		"struct ibv_port_attr.link_layer":      46,
		"struct ibv_port_attr.port_cap_flags2": 48,
	} {
		if got[name] != want {
			t.Errorf("%s = %d, want %d", name, got[name], want)
		}
	}
}

func xcrun(t *testing.T, args ...string) string {
	t.Helper()
	out, err := exec.Command("xcrun", args...).Output()
	if err != nil {
		t.Fatalf("xcrun %s: %v", strings.Join(args, " "), err)
	}
	return strings.TrimSpace(string(out))
}

func parseABIOracle(t *testing.T, out []byte) map[string]uintptr {
	t.Helper()
	got := make(map[string]uintptr)
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		name, value, ok := strings.Cut(line, "=")
		if !ok {
			t.Fatalf("malformed ABI oracle line %q", line)
		}
		n, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			t.Fatalf("parse ABI oracle %q: %v", line, err)
		}
		got[name] = uintptr(n)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("read ABI oracle: %v", err)
	}
	if len(got) == 0 {
		t.Fatal("ABI oracle produced no values")
	}
	return got
}
