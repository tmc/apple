package main

import (
	"os"
	"syscall"
	"testing"
	"unsafe"
)

func TestParsePorts(t *testing.T) {
	got := parsePorts("1, 2,255")
	want := []uint8{1, 2, 255}
	if len(got) != len(want) {
		t.Fatalf("parsePorts length = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("parsePorts[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestQueryBuffer(t *testing.T) {
	min := int(unsafe.Sizeof(ibvDeviceAttr{}))
	buf := queryBuffer(min, min, "device attr")
	if len(buf) != min {
		t.Fatalf("queryBuffer length = %d, want %d", len(buf), min)
	}
	if got := queryBuffer(0, min, "device attr"); got != nil {
		t.Fatalf("queryBuffer zero = %#v, want nil", got)
	}
}

func TestQueryStepPreview(t *testing.T) {
	step := queryStep("ibv_query_device", 0, nil, []byte{0, 1, 2, 3}, 2)
	if !step.OK {
		t.Fatal("queryStep OK = false, want true")
	}
	if step.Preview != "0001" {
		t.Fatalf("queryStep preview = %q, want 0001", step.Preview)
	}
}

func TestRDMAReadinessNames(t *testing.T) {
	if got := rdmaNetInterface("rdma_en3"); got != "en3" {
		t.Fatalf("rdmaNetInterface = %q, want en3", got)
	}
	if got := portStateName(4); got != "PORT_ACTIVE" {
		t.Fatalf("portStateName = %q, want PORT_ACTIVE", got)
	}
	if got := linkLayerName(100); got != "Thunderbolt" {
		t.Fatalf("linkLayerName = %q, want Thunderbolt", got)
	}
	if got := mtuBytes(5); got != 4096 {
		t.Fatalf("mtuBytes = %d, want 4096", got)
	}
}

func TestErrnoNamePreservesEvidenceCodes(t *testing.T) {
	tests := []struct {
		errno int
		want  string
	}{
		{22, "EINVAL"},
		{60, "errno 60 (ETIMEDOUT)"},
		{96, "errno 96 (EPROTONOSUPPORT)"},
	}
	for _, tt := range tests {
		if got := errnoName(tt.errno); got != tt.want {
			t.Fatalf("errnoName(%d) = %q, want %q", tt.errno, got, tt.want)
		}
	}
}

func TestDerivePreflightSafety(t *testing.T) {
	routeIndex := 1
	base := preflightReport{
		Devices: []preflightDevice{{
			Name:          "rdma_en3",
			State:         4,
			LinkLayer:     100,
			RouteGIDIndex: &routeIndex,
		}},
		IORegistry: map[string]int{
			"AppleThunderboltRDMAPeerInterface": 3,
			"IOThunderboltXDomainService":       8,
		},
		RecentLog: []string{"AppleThunderboltRDMA context allocation/query/free"},
	}
	tests := []struct {
		name string
		edit func(*preflightReport)
		want bool
	}{
		{"ready", nil, true},
		{"no active thunderbolt", func(r *preflightReport) { r.Devices[0].State = 1 }, false},
		{"no peer interface", func(r *preflightReport) { r.IORegistry["AppleThunderboltRDMAPeerInterface"] = 0 }, false},
		{"no xdomain", func(r *preflightReport) { r.IORegistry["IOThunderboltXDomainService"] = 0 }, false},
		{"rtr failure log", func(r *preflightReport) { r.RecentLog = []string{"Failed INIT->RTR kernel transition"} }, false},
		{"unicode rtr failure log", func(r *preflightReport) { r.RecentLog = []string{"Failed INIT→RTR kernel transition"} }, false},
		{"no route gid", func(r *preflightReport) { r.Devices[0].RouteGIDIndex = nil }, false},
		{"route gid index zero", func(r *preflightReport) {
			index := 0
			r.Devices[0].RouteGIDIndex = &index
		}, false},
	}
	for _, tt := range tests {
		report := base
		report.Devices = append([]preflightDevice(nil), base.Devices...)
		report.IORegistry = map[string]int{
			"AppleThunderboltRDMAPeerInterface": base.IORegistry["AppleThunderboltRDMAPeerInterface"],
			"IOThunderboltXDomainService":       base.IORegistry["IOThunderboltXDomainService"],
		}
		report.RecentLog = append([]string(nil), base.RecentLog...)
		if tt.edit != nil {
			tt.edit(&report)
		}
		got, reasons := derivePreflightSafety(report)
		if got != tt.want {
			t.Fatalf("%s: safe = %v, want %v, reasons=%v", tt.name, got, tt.want, reasons)
		}
		if !tt.want && len(reasons) == 0 {
			t.Fatalf("%s: no reasons returned", tt.name)
		}
	}
}

func TestPreflightGIDSelection(t *testing.T) {
	tests := []struct {
		name      string
		gids      []gidInfo
		linkLayer uint8
		wantIndex int
		wantOK    bool
	}{
		{
			name: "prefers ipv4 mapped",
			gids: []gidInfo{
				{Index: 1, Raw: "fe800000000000000000000000000001"},
				{Index: 3, Raw: "00000000000000000000ffffc0000201", IPv4Mapped: true},
			},
			linkLayer: 100,
			wantIndex: 3,
			wantOK:    true,
		},
		{
			name: "blocks thunderbolt ipv4 mapped index zero",
			gids: []gidInfo{
				{Index: 0, Raw: "00000000000000000000ffffc0000201", IPv4Mapped: true},
			},
			linkLayer: 100,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name: "falls back to apple thunderbolt index one",
			gids: []gidInfo{
				{Index: 0, Raw: "fe800000000000000000000000000001"},
				{Index: 1, Raw: "fe800000000000000000000000000002"},
			},
			linkLayer: 100,
			wantIndex: 1,
			wantOK:    true,
		},
		{
			name: "blocks thunderbolt first nonzero fallback",
			gids: []gidInfo{
				{Index: 2, Raw: "fe800000000000000000000000000002"},
			},
			linkLayer: 100,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name: "permits non-thunderbolt first nonzero fallback",
			gids: []gidInfo{
				{Index: 2, Raw: "fe800000000000000000000000000002"},
			},
			linkLayer: 0,
			wantIndex: 2,
			wantOK:    true,
		},
		{
			name:      "empty",
			gids:      nil,
			linkLayer: 100,
			wantIndex: -1,
			wantOK:    false,
		},
	}
	for _, tt := range tests {
		_, gotIndex, gotOK := selectRouteGIDInfo(tt.gids, tt.linkLayer)
		if gotOK != tt.wantOK || gotIndex != tt.wantIndex {
			t.Fatalf("%s: selectRouteGIDInfo index=%d ok=%v, want index=%d ok=%v", tt.name, gotIndex, gotOK, tt.wantIndex, tt.wantOK)
		}
	}
}

func TestPreflightGIDScanLimit(t *testing.T) {
	tests := []struct {
		tableLen  int32
		requested int
		want      int
	}{
		{0, 2, 2},
		{1, 2, 1},
		{8, 16, 8},
		{1024, 16, 8},
		{1024, -1, 0},
	}
	for _, tt := range tests {
		if got := preflightGIDScanLimit(tt.tableLen, tt.requested); got != tt.want {
			t.Fatalf("preflightGIDScanLimit(%d, %d) = %d, want %d", tt.tableLen, tt.requested, got, tt.want)
		}
	}
}

func TestCountLinesContaining(t *testing.T) {
	text := "AppleThunderboltRDMAPeerInterface one\nother AppleThunderboltRDMAPeerInterface\nnope\n"
	if got := countLinesContaining(text, "AppleThunderboltRDMAPeerInterface"); got != 2 {
		t.Fatalf("countLinesContaining = %d, want 2", got)
	}
}

func TestRDMABuffer(t *testing.T) {
	buf, mapBuf, err := rdmaBuffer(64)
	if err != nil {
		t.Fatalf("rdmaBuffer: %v", err)
	}
	defer syscall.Munmap(mapBuf)
	if len(buf) != 64 {
		t.Fatalf("rdmaBuffer length = %d, want 64", len(buf))
	}
	if len(mapBuf)%os.Getpagesize() != 0 {
		t.Fatalf("rdmaBuffer mapping length = %d, want page multiple", len(mapBuf))
	}
	if len(mapBuf) < os.Getpagesize() {
		t.Fatalf("rdmaBuffer mapping length = %d, want at least one page", len(mapBuf))
	}
	if uintptr(unsafe.Pointer(unsafe.SliceData(buf)))%uintptr(os.Getpagesize()) != 0 {
		t.Fatalf("rdmaBuffer is not page aligned")
	}
}

func TestRoundUp(t *testing.T) {
	tests := []struct {
		name string
		n    int
		unit int
		want int
	}{
		{"exact", 4096, 4096, 4096},
		{"partial", 4097, 4096, 8192},
		{"zero unit", 5, 0, 5},
	}
	for _, tt := range tests {
		if got := roundUp(tt.n, tt.unit); got != tt.want {
			t.Fatalf("%s: roundUp(%d, %d) = %d, want %d", tt.name, tt.n, tt.unit, got, tt.want)
		}
	}
}
