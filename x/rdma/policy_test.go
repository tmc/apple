package rdma

import (
	"errors"
	"strings"
	"testing"
)

func TestSelectRouteGID(t *testing.T) {
	gid0 := GID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	gid1 := GID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}
	ipv4 := GID{10: 0xff, 11: 0xff, 12: 192, 13: 0, 14: 2, 15: 1}
	tests := []struct {
		name      string
		gids      []RouteGID
		preferred int
		linkLayer uint8
		wantIndex int
		wantOK    bool
	}{
		{
			name: "auto prefers ipv4 mapped",
			gids: []RouteGID{
				{Index: 1, GID: gid1},
				{Index: 4, GID: ipv4},
			},
			preferred: -1,
			linkLayer: LinkLayerThunderbolt,
			wantIndex: 4,
			wantOK:    true,
		},
		{
			name:      "auto rejects thunderbolt ipv4 mapped index zero",
			gids:      []RouteGID{{Index: 0, GID: ipv4}},
			preferred: -1,
			linkLayer: LinkLayerThunderbolt,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name: "auto falls back to apple thunderbolt index one",
			gids: []RouteGID{
				{Index: 0, GID: gid0},
				{Index: 1, GID: gid1},
			},
			preferred: -1,
			linkLayer: LinkLayerThunderbolt,
			wantIndex: 1,
			wantOK:    true,
		},
		{
			name: "explicit index overrides auto",
			gids: []RouteGID{
				{Index: 0, GID: gid0},
				{Index: 1, GID: gid1},
			},
			preferred: 0,
			linkLayer: LinkLayerThunderbolt,
			wantIndex: 0,
			wantOK:    true,
		},
		{
			name:      "explicit missing index fails",
			gids:      []RouteGID{{Index: 1, GID: gid1}},
			preferred: 2,
			linkLayer: LinkLayerThunderbolt,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name:      "auto rejects apple thunderbolt index zero only",
			gids:      []RouteGID{{Index: 0, GID: gid0}},
			preferred: -1,
			linkLayer: LinkLayerThunderbolt,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name:      "auto rejects zero index one",
			gids:      []RouteGID{{Index: 1}},
			preferred: -1,
			linkLayer: LinkLayerThunderbolt,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name:      "auto permits non-thunderbolt first nonzero fallback",
			gids:      []RouteGID{{Index: 0, GID: gid0}},
			preferred: -1,
			linkLayer: 0,
			wantIndex: 0,
			wantOK:    true,
		},
	}
	for _, tt := range tests {
		got, gotOK := SelectRouteGID(tt.gids, tt.preferred, tt.linkLayer)
		if gotOK != tt.wantOK || got.Index != tt.wantIndex {
			t.Fatalf("%s: SelectRouteGID index=%d ok=%v, want index=%d ok=%v", tt.name, got.Index, gotOK, tt.wantIndex, tt.wantOK)
		}
	}
}

func TestSelectRouteGIDInfo(t *testing.T) {
	tests := []struct {
		name      string
		gids      []GIDInfo
		linkLayer uint8
		wantIndex int
		wantOK    bool
	}{
		{
			name: "prefers ipv4 mapped",
			gids: []GIDInfo{
				{Index: 1, Raw: "fe800000000000000000000000000001"},
				{Index: 3, Raw: "00000000000000000000ffffc0000201", IPv4Mapped: true},
			},
			linkLayer: LinkLayerThunderbolt,
			wantIndex: 3,
			wantOK:    true,
		},
		{
			name:      "blocks thunderbolt ipv4 mapped index zero",
			gids:      []GIDInfo{{Index: 0, Raw: "00000000000000000000ffffc0000201", IPv4Mapped: true}},
			linkLayer: LinkLayerThunderbolt,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name: "falls back to apple thunderbolt index one",
			gids: []GIDInfo{
				{Index: 0, Raw: "fe800000000000000000000000000001"},
				{Index: 1, Raw: "fe800000000000000000000000000002"},
			},
			linkLayer: LinkLayerThunderbolt,
			wantIndex: 1,
			wantOK:    true,
		},
		{
			name:      "blocks thunderbolt first nonzero fallback",
			gids:      []GIDInfo{{Index: 2, Raw: "fe800000000000000000000000000002"}},
			linkLayer: LinkLayerThunderbolt,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name:      "blocks zero index one",
			gids:      []GIDInfo{{Index: 1, Raw: "00000000000000000000000000000000"}},
			linkLayer: LinkLayerThunderbolt,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name:      "permits non-thunderbolt first nonzero fallback",
			gids:      []GIDInfo{{Index: 2, Raw: "fe800000000000000000000000000002"}},
			linkLayer: 0,
			wantIndex: 2,
			wantOK:    true,
		},
		{
			name:      "empty",
			gids:      nil,
			linkLayer: LinkLayerThunderbolt,
			wantIndex: -1,
			wantOK:    false,
		},
	}
	for _, tt := range tests {
		got, gotOK := SelectRouteGIDInfo(tt.gids, tt.linkLayer)
		if gotOK != tt.wantOK || got.Index != tt.wantIndex {
			t.Fatalf("%s: SelectRouteGIDInfo index=%d ok=%v, want index=%d ok=%v", tt.name, got.Index, gotOK, tt.wantIndex, tt.wantOK)
		}
	}
}

func TestRouteGIDScanLimit(t *testing.T) {
	tests := []struct {
		tableLen int32
		want     int
	}{
		{0, 8},
		{1, 1},
		{8, 8},
		{1024, 8},
	}
	for _, tt := range tests {
		if got := RouteGIDScanLimit(tt.tableLen); got != tt.want {
			t.Fatalf("RouteGIDScanLimit(%d) = %d, want %d", tt.tableLen, got, tt.want)
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
		if got := PreflightGIDScanLimit(tt.tableLen, tt.requested); got != tt.want {
			t.Fatalf("PreflightGIDScanLimit(%d, %d) = %d, want %d", tt.tableLen, tt.requested, got, tt.want)
		}
	}
}

func TestDerivePreflightSafety(t *testing.T) {
	routeIndex := 1
	base := PreflightReport{
		Devices: []PreflightDevice{{
			Name:          "rdma_en3",
			State:         PortActive,
			LinkLayer:     LinkLayerThunderbolt,
			RouteGIDIndex: &routeIndex,
		}},
		IORegistry: map[string]int{
			IORegistryPeerInterface:  3,
			IORegistryXDomainService: 8,
		},
		RecentLog: []string{"AppleThunderboltRDMA context allocation/query/free"},
	}
	tests := []struct {
		name string
		edit func(*PreflightReport)
		want bool
	}{
		{"ready", nil, true},
		{"no active thunderbolt", func(r *PreflightReport) { r.Devices[0].State = 1 }, false},
		{"no peer interface", func(r *PreflightReport) { r.IORegistry[IORegistryPeerInterface] = 0 }, false},
		{"no xdomain", func(r *PreflightReport) { r.IORegistry[IORegistryXDomainService] = 0 }, false},
		{"rtr failure log", func(r *PreflightReport) { r.RecentLog = []string{"Failed INIT->RTR kernel transition"} }, false},
		{"unicode rtr failure log", func(r *PreflightReport) { r.RecentLog = []string{"Failed INIT→RTR kernel transition"} }, false},
		{"no route gid", func(r *PreflightReport) { r.Devices[0].RouteGIDIndex = nil }, false},
		{"route gid index zero", func(r *PreflightReport) {
			index := 0
			r.Devices[0].RouteGIDIndex = &index
		}, false},
	}
	for _, tt := range tests {
		report := base
		report.Devices = append([]PreflightDevice(nil), base.Devices...)
		report.IORegistry = map[string]int{
			IORegistryPeerInterface:  base.IORegistry[IORegistryPeerInterface],
			IORegistryXDomainService: base.IORegistry[IORegistryXDomainService],
		}
		report.RecentLog = append([]string(nil), base.RecentLog...)
		if tt.edit != nil {
			tt.edit(&report)
		}
		got, reasons := DerivePreflightSafety(report)
		if got != tt.want {
			t.Fatalf("%s: safe = %v, want %v, reasons=%v", tt.name, got, tt.want, reasons)
		}
		if !tt.want && len(reasons) == 0 {
			t.Fatalf("%s: no reasons returned", tt.name)
		}
	}
}

func TestErrnoText(t *testing.T) {
	got := ErrnoText(16)
	if !strings.Contains(got, "EBUSY") || !strings.Contains(got, "reboot") {
		t.Fatalf("ErrnoText(16) = %q", got)
	}
	if got := ErrnoName(60); got != "ETIMEDOUT" {
		t.Fatalf("ErrnoName(60) = %q, want ETIMEDOUT", got)
	}
}

func TestRequireRTRAttemptAllowed(t *testing.T) {
	if err := RequireRTRAttemptAllowed(true); err != nil {
		t.Fatalf("RequireRTRAttemptAllowed(true) = %v", err)
	}
	err := RequireRTRAttemptAllowed(false)
	if !errors.Is(err, ErrRTRUnsafe) {
		t.Fatalf("errors.Is(%v, ErrRTRUnsafe) = false", err)
	}
	if !strings.Contains(err.Error(), "one bounded attempt") {
		t.Fatalf("RequireRTRAttemptAllowed(false) = %q", err)
	}
}
