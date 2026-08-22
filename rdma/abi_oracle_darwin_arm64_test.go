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
		"struct ibv_context.ops":               8,
		"struct ibv_context_ops.poll_cq":       88,
		"struct ibv_context_ops.post_send":     200,
		"struct ibv_context_ops.post_recv":     208,
		"struct ibv_qp.qp_num":                 ibvQPNumOffset,
		"struct ibv_mr.lkey":                   ibvMRLKeyOffset,
		"struct ibv_mr.rkey":                   ibvMRRKeyOffset,
		"struct ibv_comp_channel.fd":           8,
		"struct ibv_cq.channel":                8,
		"struct ibv_port_attr.active_mtu":      8,
		"struct ibv_port_attr.gid_tbl_len":     12,
		"struct ibv_port_attr.lid":             34,
		"struct ibv_port_attr.link_layer":      46,
		"struct ibv_port_attr.port_cap_flags2": 48,
		"IBV_QPT_RC":                           2,
		"IBV_QPT_UC":                           3,
		"IBV_WR_RDMA_WRITE":                    0,
		"IBV_WR_SEND":                          2,
		"IBV_WR_RDMA_READ":                     4,
		"IBV_ACCESS_LOCAL_WRITE":               1,
		"IBV_ACCESS_REMOTE_WRITE":              2,
		"IBV_ACCESS_REMOTE_READ":               4,
		"IBV_WC_SUCCESS":                       0,
		"IBV_WC_LOC_PROT_ERR":                  4,
		"IBV_WC_LOC_ACCESS_ERR":                8,
		"IBV_WC_REM_ACCESS_ERR":                10,
		"IBV_QP_STATE":                         1,
		"IBV_QP_ACCESS_FLAGS":                  8,
		"IBV_QP_PKEY_INDEX":                    16,
		"IBV_QP_PORT":                          32,
		"IBV_QP_AV":                            128,
		"IBV_QP_PATH_MTU":                      256,
		"IBV_QP_RQ_PSN":                        4096,
		"IBV_QP_SQ_PSN":                        65536,
		"IBV_QP_DEST_QPN":                      1048576,
		"IBV_SEND_SIGNALED":                    2,
		"ENOTSUP":                              45,
	} {
		value, ok := got[name]
		if !ok {
			t.Errorf("ABI oracle omitted %s", name)
			continue
		}
		if value != want {
			t.Errorf("%s = %d, want %d", name, value, want)
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
