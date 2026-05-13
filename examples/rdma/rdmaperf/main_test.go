package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/tmc/apple/rdma"
)

func TestParseSize(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"1", 1},
		{"64K", 64 * 1024},
		{"1M", 1024 * 1024},
	}
	for _, tt := range tests {
		if got := parseSize(tt.in); got != tt.want {
			t.Fatalf("parseSize(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestPatternRoundTrip(t *testing.T) {
	for _, pattern := range []string{"stream", "pingpong", "duplex"} {
		if got := patternName(patternID(pattern)); got != pattern {
			t.Fatalf("patternName(patternID(%q)) = %q", pattern, got)
		}
	}
}

func TestSummarizeLatency(t *testing.T) {
	got := summarizeLatency([]time.Duration{
		10 * time.Microsecond,
		20 * time.Microsecond,
		30 * time.Microsecond,
		40 * time.Microsecond,
	})
	if got == nil {
		t.Fatal("summarizeLatency returned nil")
	}
	if got.Count != 4 {
		t.Fatalf("Count = %d, want 4", got.Count)
	}
	if got.Min != "10µs" || got.P50 != "20µs" || got.Max != "40µs" {
		t.Fatalf("summary = %#v", got)
	}
}

func TestTCPNetwork(t *testing.T) {
	tests := []struct {
		addr string
		want string
	}{
		{"127.0.0.1:9000", "tcp4"},
		{"0.0.0.0:9000", "tcp4"},
		{":9000", "tcp"},
		{"[::1]:9000", "tcp"},
		{"localhost:9000", "tcp"},
	}
	for _, tt := range tests {
		if got := tcpNetwork(tt.addr); got != tt.want {
			t.Fatalf("tcpNetwork(%q) = %q, want %q", tt.addr, got, tt.want)
		}
	}
}

func TestRDMAReadinessNames(t *testing.T) {
	if got := rdmaNetInterface("rdma_en2"); got != "en2" {
		t.Fatalf("rdmaNetInterface = %q, want en2", got)
	}
	if got := portStateName(1); got != "PORT_DOWN" {
		t.Fatalf("portStateName = %q, want PORT_DOWN", got)
	}
	if got := linkLayerName(100); got != "Thunderbolt" {
		t.Fatalf("linkLayerName = %q, want Thunderbolt", got)
	}
	if got := mtuBytes(5); got != 4096 {
		t.Fatalf("mtuBytes = %d, want 4096", got)
	}
}

func TestRTRAttrUsesLocalGIDIndex(t *testing.T) {
	r := rdmaResources{gidIndex: 0, port: ibvPortAttr{ActiveMTU: 5}}
	remote := rdmaPeerInfo{
		LID:       1,
		QPN:       2,
		PSN:       3,
		GID:       "fe800000000000000000000000000001",
		UseGlobal: true,
		ActiveMTU: 5,
	}
	attr, err := r.rtrAttr(remote)
	if err != nil {
		t.Fatal(err)
	}
	if attr.PathMTU != 5 {
		t.Fatalf("PathMTU = %d, want active MTU 5", attr.PathMTU)
	}
	if attr.AHAttr.GRH.SGIDIndex != 0 {
		t.Fatalf("SGIDIndex = %d, want local gid index 0", attr.AHAttr.GRH.SGIDIndex)
	}
	if attr.AHAttr.GRH.HopLimit != 1 {
		t.Fatalf("HopLimit = %d, want 1", attr.AHAttr.GRH.HopLimit)
	}
	if attr.AHAttr.IsGlobal != 1 {
		t.Fatalf("IsGlobal = %d, want 1", attr.AHAttr.IsGlobal)
	}
	if attr.AHAttr.GRH.DGID != (rdma.IbvGID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}) {
		t.Fatalf("DGID = %x, want remote gid", attr.AHAttr.GRH.DGID)
	}
}

func TestRTRAttrRejectsInvalidLocalGIDIndex(t *testing.T) {
	for _, index := range []int{-1, 256} {
		r := rdmaResources{gidIndex: index}
		_, err := r.rtrAttr(rdmaPeerInfo{
			GID:       "fe800000000000000000000000000001",
			UseGlobal: true,
		})
		if err == nil {
			t.Fatalf("rtrAttr accepted out-of-range local gid index %d", index)
		}
	}
}

func TestRTRAttrLeavesGRHZeroForLocalOnlyRoute(t *testing.T) {
	r := rdmaResources{gidIndex: 7}
	attr, err := r.rtrAttr(rdmaPeerInfo{
		LID:       1,
		QPN:       2,
		PSN:       3,
		UseGlobal: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	if attr.AHAttr.IsGlobal != 0 {
		t.Fatalf("IsGlobal = %d, want 0", attr.AHAttr.IsGlobal)
	}
	if attr.AHAttr.GRH != (rdma.IbvGlobalRoute{}) {
		t.Fatalf("GRH = %#v, want zero", attr.AHAttr.GRH)
	}
}

func TestNegotiatedPathMTU(t *testing.T) {
	tests := []struct {
		name   string
		local  int32
		remote int32
		want   int32
	}{
		{"same active", 5, 5, 5},
		{"remote lower", 5, 4, 4},
		{"local lower", 3, 5, 3},
		{"missing remote", 5, 0, rdma.IBV_MTU_1024},
		{"missing local", 0, 5, rdma.IBV_MTU_1024},
		{"invalid", 99, 99, rdma.IBV_MTU_1024},
	}
	for _, tt := range tests {
		if got := negotiatedPathMTU(tt.local, tt.remote); got != tt.want {
			t.Fatalf("%s: negotiatedPathMTU(%d, %d) = %d, want %d", tt.name, tt.local, tt.remote, got, tt.want)
		}
	}
}

func TestSelectRouteGID(t *testing.T) {
	gid0 := rdma.IbvGID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	gid1 := rdma.IbvGID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}
	ipv4 := rdma.IbvGID{10: 0xff, 11: 0xff, 12: 192, 13: 0, 14: 2, 15: 1}
	tests := []struct {
		name      string
		gids      []routeGID
		preferred int
		linkLayer uint8
		wantIndex int
		wantOK    bool
	}{
		{
			name: "auto prefers ipv4 mapped",
			gids: []routeGID{
				{index: 1, gid: gid1},
				{index: 4, gid: ipv4},
			},
			preferred: -1,
			linkLayer: 100,
			wantIndex: 4,
			wantOK:    true,
		},
		{
			name:      "auto rejects thunderbolt ipv4 mapped index zero",
			gids:      []routeGID{{index: 0, gid: ipv4}},
			preferred: -1,
			linkLayer: 100,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name: "auto falls back to apple thunderbolt index one",
			gids: []routeGID{
				{index: 0, gid: gid0},
				{index: 1, gid: gid1},
			},
			preferred: -1,
			linkLayer: 100,
			wantIndex: 1,
			wantOK:    true,
		},
		{
			name: "explicit index overrides auto",
			gids: []routeGID{
				{index: 0, gid: gid0},
				{index: 1, gid: gid1},
			},
			preferred: 0,
			linkLayer: 100,
			wantIndex: 0,
			wantOK:    true,
		},
		{
			name:      "explicit missing index fails",
			gids:      []routeGID{{index: 1, gid: gid1}},
			preferred: 2,
			linkLayer: 100,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name:      "auto rejects apple thunderbolt index zero only",
			gids:      []routeGID{{index: 0, gid: gid0}},
			preferred: -1,
			linkLayer: 100,
			wantIndex: -1,
			wantOK:    false,
		},
		{
			name:      "auto permits non-thunderbolt first nonzero fallback",
			gids:      []routeGID{{index: 0, gid: gid0}},
			preferred: -1,
			linkLayer: 0,
			wantIndex: 0,
			wantOK:    true,
		},
	}
	for _, tt := range tests {
		_, gotIndex, gotOK := selectRouteGID(tt.gids, tt.preferred, tt.linkLayer)
		if gotOK != tt.wantOK || gotIndex != tt.wantIndex {
			t.Fatalf("%s: selectRouteGID index=%d ok=%v, want index=%d ok=%v", tt.name, gotIndex, gotOK, tt.wantIndex, tt.wantOK)
		}
	}
}

func TestValidateGIDIndexFlag(t *testing.T) {
	tests := []struct {
		index   int
		wantErr bool
	}{
		{-2, true},
		{-1, false},
		{0, false},
		{255, false},
		{256, true},
	}
	for _, tt := range tests {
		err := validateGIDIndexFlag(tt.index)
		if (err != nil) != tt.wantErr {
			t.Fatalf("validateGIDIndexFlag(%d) err=%v, wantErr=%v", tt.index, err, tt.wantErr)
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
		if got := routeGIDScanLimit(tt.tableLen); got != tt.want {
			t.Fatalf("routeGIDScanLimit(%d) = %d, want %d", tt.tableLen, got, tt.want)
		}
	}
}

func TestRDMAPingpongRequiresRTRConfirm(t *testing.T) {
	if err := checkRDMAPingpongOptIn(false); err == nil {
		t.Fatal("checkRDMAPingpongOptIn(false) succeeded")
	}
	if err := checkRDMAPingpongOptIn(true); err != nil {
		t.Fatalf("checkRDMAPingpongOptIn(true) = %v", err)
	}
}

func TestUsageWarnsBeforeRDMAPingpongExamples(t *testing.T) {
	old := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stderr = w
	usage()
	w.Close()
	os.Stderr = old
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	text := string(out)
	warn := strings.Index(text, "rdma-pingpong drives QP INIT->RTR")
	example := strings.Index(text, "rdmaperf rdma-pingpong -listen")
	if warn < 0 {
		t.Fatalf("usage output missing RDMA warning:\n%s", text)
	}
	if example < 0 {
		t.Fatalf("usage output missing RDMA examples:\n%s", text)
	}
	if warn > example {
		t.Fatalf("usage warning appears after RDMA example:\n%s", text)
	}
	for _, line := range strings.Split(text[example:], "\n") {
		if !strings.Contains(line, "rdmaperf rdma-pingpong") {
			continue
		}
		if !strings.Contains(line, "-setup-timeout 12s") {
			t.Fatalf("RDMA example is not bounded by setup timeout: %q", line)
		}
	}
}

func TestRDMAPingpongOptInGateRunsFirst(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	accepted := make(chan bool, 1)
	go func() {
		if tcp, ok := ln.(*net.TCPListener); ok {
			_ = tcp.SetDeadline(time.Now().Add(250 * time.Millisecond))
		}
		c, err := ln.Accept()
		if err == nil {
			c.Close()
			accepted <- true
			return
		}
		accepted <- false
	}()

	cmd := exec.Command(os.Args[0], "-test.run=TestRDMAPingpongOptInGateRunsFirstHelper")
	cmd.Env = append(os.Environ(),
		"RDMAPERF_HELPER=1",
		"RDMAPERF_ADDR="+ln.Addr().String(),
	)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("helper succeeded, want opt-in failure; output:\n%s", out)
	}
	if !strings.Contains(string(out), "rdma-pingpong drives QP INIT->RTR") {
		t.Fatalf("helper output missing opt-in error:\n%s", out)
	}
	if <-accepted {
		t.Fatal("rdma-pingpong dialed before checking -allow-rtr")
	}
}

func TestRDMAPingpongOptInGateRunsFirstHelper(t *testing.T) {
	if os.Getenv("RDMAPERF_HELPER") != "1" {
		return
	}
	rdmaPingpong([]string{"-addr", os.Getenv("RDMAPERF_ADDR")})
}

func TestExchangeRDMAPeerInfoSymmetric(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	left := rdmaPeerInfo{Name: "left", LID: 1, QPN: 2, PSN: 3, GIDIndex: 1, GID: "00000000000000000000ffffac1ffd01", UseGlobal: true}
	right := rdmaPeerInfo{Name: "right", LID: 4, QPN: 5, PSN: 6, GIDIndex: 1, GID: "00000000000000000000ffffac1ffd02", UseGlobal: true}
	errc := make(chan error, 2)
	var gotLeft, gotRight rdmaPeerInfo
	go func() {
		var err error
		gotLeft, err = exchangeRDMAPeerInfo(newRDMAControlConn(a), left, time.Second)
		errc <- err
	}()
	go func() {
		var err error
		gotRight, err = exchangeRDMAPeerInfo(newRDMAControlConn(b), right, time.Second)
		errc <- err
	}()
	for i := 0; i < 2; i++ {
		if err := <-errc; err != nil {
			t.Fatal(err)
		}
	}
	if gotLeft != right {
		t.Fatalf("left saw %#v, want %#v", gotLeft, right)
	}
	if gotRight != left {
		t.Fatalf("right saw %#v, want %#v", gotRight, left)
	}
}

func TestExchangeRDMAPeerInfoRejectsMissingQPN(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	valid := rdmaPeerInfo{Name: "valid", LID: 1, QPN: 2, PSN: 3, GIDIndex: 1, GID: "00000000000000000000ffffac1ffd01", UseGlobal: true}
	empty := rdmaPeerInfo{Name: "empty"}
	errc := make(chan error, 2)
	var got rdmaPeerInfo
	go func() {
		var err error
		got, err = exchangeRDMAPeerInfo(newRDMAControlConn(a), valid, time.Second)
		errc <- err
	}()
	go func() {
		_, err := exchangeRDMAPeerInfo(newRDMAControlConn(b), empty, time.Second)
		errc <- err
	}()

	var sawInvalid bool
	for i := 0; i < 2; i++ {
		err := <-errc
		if err == nil {
			continue
		}
		if strings.Contains(err.Error(), "invalid remote rdma info") && strings.Contains(err.Error(), "missing qpn") {
			sawInvalid = true
			continue
		}
		t.Fatalf("unexpected error: %v", err)
	}
	if !sawInvalid {
		t.Fatal("missing qpn was accepted")
	}
	if got.Name != "empty" {
		t.Fatalf("got peer info %#v, want decoded empty peer", got)
	}
}

func TestRDMAControlConnReusesDecoderAcrossStages(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	client := newRDMAControlConn(a)
	server := newRDMAControlConn(b)
	clientInfo := rdmaPeerInfo{Name: "client", LID: 1, QPN: 2, PSN: 3, GIDIndex: 1, GID: "00000000000000000000ffffac1ffd01", UseGlobal: true}
	serverInfo := rdmaPeerInfo{Name: "server", LID: 1, QPN: 4, PSN: 5, GIDIndex: 1, GID: "00000000000000000000ffffac1ffd02", UseGlobal: true}
	errc := make(chan error, 2)
	go func() {
		var res rdmaBenchResult
		if err := runRDMAControlHello(client, &res, "pre-resource-control", "client", time.Second); err != nil {
			errc <- err
			return
		}
		if err := runRDMAControlHello(client, &res, "post-resource-control", "client", time.Second); err != nil {
			errc <- err
			return
		}
		got, err := exchangeRDMAPeerInfo(client, clientInfo, time.Second)
		if err != nil {
			errc <- err
			return
		}
		if got != serverInfo {
			errc <- fmt.Errorf("client saw %#v, want %#v", got, serverInfo)
			return
		}
		errc <- exchangeRDMAReady(client, nil, time.Second)
	}()
	go func() {
		var res rdmaBenchResult
		if err := runRDMAControlHello(server, &res, "pre-resource-control", "server", time.Second); err != nil {
			errc <- err
			return
		}
		if err := runRDMAControlHello(server, &res, "post-resource-control", "server", time.Second); err != nil {
			errc <- err
			return
		}
		got, err := exchangeRDMAPeerInfo(server, serverInfo, time.Second)
		if err != nil {
			errc <- err
			return
		}
		if got != clientInfo {
			errc <- fmt.Errorf("server saw %#v, want %#v", got, clientInfo)
			return
		}
		errc <- exchangeRDMAReady(server, nil, time.Second)
	}()
	for i := 0; i < 2; i++ {
		if err := <-errc; err != nil {
			t.Fatal(err)
		}
	}
}

func TestExchangeRDMAControlHelloSymmetric(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	left := rdmaControlHello{Stage: "pre-resource-control", Role: "client"}
	right := rdmaControlHello{Stage: "pre-resource-control", Role: "server"}
	errc := make(chan error, 2)
	var gotLeft, gotRight rdmaControlHello
	go func() {
		var err error
		gotLeft, err = exchangeRDMAControlHello(newRDMAControlConn(a), left, time.Second)
		errc <- err
	}()
	go func() {
		var err error
		gotRight, err = exchangeRDMAControlHello(newRDMAControlConn(b), right, time.Second)
		errc <- err
	}()
	for i := 0; i < 2; i++ {
		if err := <-errc; err != nil {
			t.Fatal(err)
		}
	}
	if gotLeft != right {
		t.Fatalf("left saw %#v, want %#v", gotLeft, right)
	}
	if gotRight != left {
		t.Fatalf("right saw %#v, want %#v", gotRight, left)
	}
}

func TestRunRDMAControlHelloRecordsFailure(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	var res rdmaBenchResult
	err := runRDMAControlHello(newRDMAControlConn(a), &res, "post-resource-control", "client", 10*time.Millisecond)
	if err == nil {
		t.Fatal("runRDMAControlHello succeeded with silent peer")
	}
	if res.Stage != "post-resource-control" {
		t.Fatalf("stage = %q, want post-resource-control", res.Stage)
	}
	if len(res.Control) != 1 {
		t.Fatalf("control events = %d, want 1", len(res.Control))
	}
	event := res.Control[0]
	if event.Stage != "post-resource-control" || event.OK {
		t.Fatalf("control event = %#v, want failed post-resource-control", event)
	}
	if !strings.Contains(event.Error, "receive remote control hello") {
		t.Fatalf("control error = %q, want receive remote control hello", event.Error)
	}
}

func TestRunRDMAControlHelloRejectsWrongStage(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	errc := make(chan error, 1)
	go func() {
		_, err := exchangeRDMAControlHello(newRDMAControlConn(b), rdmaControlHello{Stage: "exchange-rdma-info", Role: "server"}, time.Second)
		errc <- err
	}()
	var res rdmaBenchResult
	err := runRDMAControlHello(newRDMAControlConn(a), &res, "pre-resource-control", "client", time.Second)
	if err == nil {
		t.Fatal("runRDMAControlHello accepted wrong remote stage")
	}
	if sendErr := <-errc; sendErr != nil {
		t.Fatalf("peer exchange error: %v", sendErr)
	}
	if len(res.Control) != 1 {
		t.Fatalf("control events = %d, want 1", len(res.Control))
	}
	event := res.Control[0]
	if event.OK {
		t.Fatalf("control event = %#v, want failed wrong-stage event", event)
	}
	if !strings.Contains(event.Error, `stage "exchange-rdma-info"`) {
		t.Fatalf("control error = %q, want wrong stage", event.Error)
	}
}

func TestRunRDMAControlHelloRejectsSameRole(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	errc := make(chan error, 1)
	go func() {
		_, err := exchangeRDMAControlHello(newRDMAControlConn(b), rdmaControlHello{Stage: "pre-resource-control", Role: "client"}, time.Second)
		errc <- err
	}()
	var res rdmaBenchResult
	err := runRDMAControlHello(newRDMAControlConn(a), &res, "pre-resource-control", "client", time.Second)
	if err == nil {
		t.Fatal("runRDMAControlHello accepted same remote role")
	}
	if sendErr := <-errc; sendErr != nil {
		t.Fatalf("peer exchange error: %v", sendErr)
	}
	if len(res.Control) != 1 {
		t.Fatalf("control events = %d, want 1", len(res.Control))
	}
	event := res.Control[0]
	if event.OK {
		t.Fatalf("control event = %#v, want failed same-role event", event)
	}
	if !strings.Contains(event.Error, `role "client"`) {
		t.Fatalf("control error = %q, want same role", event.Error)
	}
}

func TestRunRDMAControlHelloRejectsUnknownRole(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	errc := make(chan error, 1)
	go func() {
		_, err := exchangeRDMAControlHello(newRDMAControlConn(b), rdmaControlHello{Stage: "pre-resource-control", Role: "observer"}, time.Second)
		errc <- err
	}()
	var res rdmaBenchResult
	err := runRDMAControlHello(newRDMAControlConn(a), &res, "pre-resource-control", "client", time.Second)
	if err == nil {
		t.Fatal("runRDMAControlHello accepted unknown remote role")
	}
	if sendErr := <-errc; sendErr != nil {
		t.Fatalf("peer exchange error: %v", sendErr)
	}
	if len(res.Control) != 1 {
		t.Fatalf("control events = %d, want 1", len(res.Control))
	}
	event := res.Control[0]
	if event.OK {
		t.Fatalf("control event = %#v, want failed unknown-role event", event)
	}
	if !strings.Contains(event.Error, `role "observer", want "server"`) {
		t.Fatalf("control error = %q, want unknown role", event.Error)
	}
}

func TestExchangeRDMAPeerInfoTimeout(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	_, err := exchangeRDMAPeerInfo(newRDMAControlConn(a), rdmaPeerInfo{Name: "left"}, 10*time.Millisecond)
	if err == nil {
		t.Fatal("exchangeRDMAPeerInfo succeeded with silent peer")
	}
	if !strings.Contains(err.Error(), "receive remote rdma info") {
		t.Fatalf("error = %v, want receive remote rdma info", err)
	}
}

func TestExchangeRDMAReadyTimeout(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	err := exchangeRDMAReady(newRDMAControlConn(a), nil, 10*time.Millisecond)
	if err == nil {
		t.Fatal("exchangeRDMAReady succeeded with silent peer")
	}
	if !strings.Contains(err.Error(), "receive rdma ready status") {
		t.Fatalf("error = %v, want receive rdma ready status", err)
	}
}

func TestExchangeRDMAReadyReportsRemoteFailure(t *testing.T) {
	a, b := net.Pipe()
	defer a.Close()
	defer b.Close()

	errc := make(chan error, 2)
	go func() {
		errc <- exchangeRDMAReady(newRDMAControlConn(a), nil, time.Second)
	}()
	go func() {
		errc <- exchangeRDMAReady(newRDMAControlConn(b), errRDMASetupTimeout, time.Second)
	}()
	errs := []error{<-errc, <-errc}
	var sawRemote bool
	for _, err := range errs {
		if err == nil {
			continue
		}
		if strings.Contains(err.Error(), "remote rdma setup failed") {
			sawRemote = true
			continue
		}
		t.Fatalf("unexpected ready error: %v", err)
	}
	if !sawRemote {
		t.Fatal("did not observe remote setup failure")
	}
}

func TestRDMAPingpongResultKeepsStageAndLocalOnExchangeFailure(t *testing.T) {
	res := rdmaBenchResult{
		Stage: "exchange-rdma-info",
		Local: rdmaPeerInfo{
			Name:      "rdma_en1",
			LID:       1,
			QPN:       2304,
			PSN:       7,
			GIDIndex:  1,
			GID:       "00000000000000000000ffffac1ffd01",
			UseGlobal: true,
		},
		Error: "receive remote rdma info: i/o timeout",
	}
	data, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	for _, want := range []string{
		`"stage":"exchange-rdma-info"`,
		`"qpn":2304`,
		`"gid_index":1`,
		`"receive remote rdma info`,
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("result JSON %s missing %s", text, want)
		}
	}
}

func TestRunTCPLocalhost(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 3; i++ {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			serveConn(c, true)
		}
	}()

	stream := runTCP(ln.Addr().String(), "stream", 1024, 50*time.Millisecond)
	if stream.Error != "" {
		t.Fatalf("stream error: %v", stream.Error)
	}
	if stream.Bytes == 0 || stream.BytesPerSec == 0 {
		t.Fatalf("stream result missing bytes/rate: %#v", stream)
	}

	ping := runTCP(ln.Addr().String(), "pingpong", 64, 50*time.Millisecond)
	if ping.Error != "" {
		t.Fatalf("pingpong error: %v", ping.Error)
	}
	if ping.Latency == nil || ping.Latency.Count == 0 {
		t.Fatalf("pingpong missing latency: %#v", ping)
	}

	duplex := runTCP(ln.Addr().String(), "duplex", 1024, 50*time.Millisecond)
	if duplex.Error != "" {
		t.Fatalf("duplex error: %v", duplex.Error)
	}
	if duplex.Bytes == 0 || duplex.BytesPerSec == 0 {
		t.Fatalf("duplex result missing bytes/rate: %#v", duplex)
	}

	<-done
}
