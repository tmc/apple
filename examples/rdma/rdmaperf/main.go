// Command rdmaperf measures TCP and RDMA readiness paths.
package main

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/tmc/apple/rdma"
	xrdma "github.com/tmc/apple/x/rdma"
)

const headerSize = 8

type result struct {
	Mode        string          `json:"mode"`
	Pattern     string          `json:"pattern,omitempty"`
	Addr        string          `json:"addr,omitempty"`
	LocalAddr   string          `json:"local_addr,omitempty"`
	RemoteAddr  string          `json:"remote_addr,omitempty"`
	Duration    string          `json:"duration,omitempty"`
	Elapsed     string          `json:"elapsed,omitempty"`
	Size        int             `json:"size,omitempty"`
	Bytes       uint64          `json:"bytes,omitempty"`
	Messages    uint64          `json:"messages,omitempty"`
	BytesPerSec float64         `json:"bytes_per_sec,omitempty"`
	MsgsPerSec  float64         `json:"msgs_per_sec,omitempty"`
	Latency     *latencySummary `json:"latency,omitempty"`
	RDMA        *rdmaSummary    `json:"rdma,omitempty"`
	Error       string          `json:"error,omitempty"`
}

type rdmaBenchResult struct {
	Mode           string             `json:"mode"`
	Role           string             `json:"role"`
	Addr           string             `json:"addr,omitempty"`
	Device         string             `json:"device,omitempty"`
	Stage          string             `json:"stage,omitempty"`
	Control        []rdmaControlEvent `json:"control,omitempty"`
	Size           int                `json:"size"`
	Iterations     int                `json:"iterations"`
	Elapsed        string             `json:"elapsed,omitempty"`
	Bytes          uint64             `json:"bytes,omitempty"`
	BytesPerSec    float64            `json:"bytes_per_sec,omitempty"`
	MessagesPerSec float64            `json:"messages_per_sec,omitempty"`
	Latency        *latencySummary    `json:"latency,omitempty"`
	Local          rdmaPeerInfo       `json:"local"`
	Remote         rdmaPeerInfo       `json:"remote"`
	Error          string             `json:"error,omitempty"`
}

type latencySummary struct {
	Count int64  `json:"count"`
	Min   string `json:"min"`
	P50   string `json:"p50"`
	P95   string `json:"p95"`
	P99   string `json:"p99"`
	Max   string `json:"max"`
}

type rdmaSummary struct {
	Available bool              `json:"available"`
	Devices   []rdmaDevice      `json:"devices,omitempty"`
	Steps     []rdmaStep        `json:"steps,omitempty"`
	Notes     []string          `json:"notes,omitempty"`
	Fields    map[string]string `json:"fields,omitempty"`
}

type rdmaDevice struct {
	Index        int    `json:"index"`
	Name         string `json:"name"`
	NetInterface string `json:"net_interface,omitempty"`
	Handle       string `json:"handle"`
}

type rdmaStep struct {
	Name   string         `json:"name"`
	OK     bool           `json:"ok"`
	Return int            `json:"return,omitempty"`
	Handle string         `json:"handle,omitempty"`
	Fields map[string]any `json:"fields,omitempty"`
	Error  string         `json:"error,omitempty"`
}

type rdmaPeerInfo struct {
	Name      string `json:"name,omitempty"`
	LID       uint16 `json:"lid"`
	QPN       uint32 `json:"qpn"`
	PSN       uint32 `json:"psn"`
	GIDIndex  int    `json:"gid_index"`
	GID       string `json:"gid"`
	UseGlobal bool   `json:"use_global"`
	ActiveMTU int32  `json:"active_mtu,omitempty"`
}

type rdmaReadyStatus struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type rdmaControlHello struct {
	Stage string `json:"stage"`
	Role  string `json:"role"`
}

type rdmaControlEvent struct {
	Stage       string `json:"stage"`
	OK          bool   `json:"ok"`
	RemoteStage string `json:"remote_stage,omitempty"`
	RemoteRole  string `json:"remote_role,omitempty"`
	Error       string `json:"error,omitempty"`
}

type rdmaControlConn struct {
	conn net.Conn
	enc  *json.Encoder
	dec  *json.Decoder
}

func newRDMAControlConn(c net.Conn) *rdmaControlConn {
	return &rdmaControlConn{
		conn: c,
		enc:  json.NewEncoder(c),
		dec:  json.NewDecoder(c),
	}
}

var errRDMASetupTimeout = errors.New("rdma setup timeout")

type ibvPortAttr struct {
	State         int32
	MaxMTU        int32
	ActiveMTU     int32
	GIDTblLen     int32
	PortCapFlags  uint32
	MaxMsgSZ      uint32
	BadPKeyCntr   uint32
	QKeyViolCntr  uint32
	PKeyTblLen    uint16
	LID           uint16
	SMLID         uint16
	LMC           uint8
	MaxVLNum      uint8
	SMSL          uint8
	SubnetTimeout uint8
	InitTypeReply uint8
	ActiveWidth   uint8
	ActiveSpeed   uint8
	PhysState     uint8
	LinkLayer     uint8
	Flags         uint8
	PortCapFlags2 uint16
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "serve":
		serve(os.Args[2:])
	case "tcp":
		tcpClient(os.Args[2:])
	case "sweep":
		sweep(os.Args[2:])
	case "rdma-probe":
		rdmaProbe(os.Args[2:])
	case "rdma-pingpong":
		rdmaPingpong(os.Args[2:])
	case "rdma-port-state":
		rdmaPortState(os.Args[2:])
	case "interfaces":
		interfaces(os.Args[2:])
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "rdmaperf: unknown command %q\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `usage: rdmaperf <command> [options]

Commands:
  serve       Serve TCP benchmark connections.
  tcp         Run one TCP benchmark against a server.
  sweep       Run tcp across common payload sizes.
  rdma-probe  Exercise RDMA discovery/open/query/resource readiness.
  rdma-pingpong
              Run RDMA SEND/RECV ping-pong using TCP only for setup exchange.
  interfaces  List local interface addresses useful for -listen and -addr.

Patterns:
  stream      Client sends payloads for -duration; server counts bytes.
  pingpong    Client sends one payload and waits for echo; reports latency.
  duplex      Client and server send simultaneously; reports aggregate bytes.

Examples:
  rdmaperf serve -listen 169.254.14.228:9000
  rdmaperf tcp -addr 169.254.x.y:9000 -pattern stream -size 1M -duration 30s
  rdmaperf sweep -addr 169.254.x.y:9000 -pattern pingpong -duration 10s
  rdmaperf rdma-probe -json

Warning:
  rdma-pingpong drives QP INIT->RTR and can wedge Apple Thunderbolt RDMA ports.
  Run rdmainfo preflight first and use -allow-rtr for one bounded attempt only.
  See README.md for the outer timeout wrapper to use on macOS.

RDMA examples:
  rdmaperf rdma-pingpong -listen :9100 -setup-timeout 12s -allow-rtr
  rdmaperf rdma-pingpong -addr 169.254.x.y:9100 -setup-timeout 12s -allow-rtr
`)
}

func serve(args []string) {
	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	listenAddr := fs.String("listen", ":9000", "listen address")
	jsonOut := fs.Bool("json", false, "print JSON connection summaries")
	fs.Parse(args)

	ln, err := listenTCP(*listenAddr)
	if err != nil {
		fatalf("listen: %v", err)
	}
	fmt.Fprintf(os.Stderr, "rdmaperf serving on %s\n", ln.Addr())
	for {
		c, err := ln.Accept()
		if err != nil {
			fatalf("accept: %v", err)
		}
		go serveConn(c, *jsonOut)
	}
}

func serveConn(c net.Conn, jsonOut bool) {
	defer c.Close()
	start := time.Now()
	res := result{
		Mode:       "serve",
		LocalAddr:  c.LocalAddr().String(),
		RemoteAddr: c.RemoteAddr().String(),
	}
	var hdr [headerSize]byte
	if _, err := io.ReadFull(c, hdr[:]); err != nil {
		res.Error = "read header: " + err.Error()
		finishResult(&res, time.Since(start))
		if jsonOut {
			writeJSON(res)
		} else {
			printResult(res)
		}
		return
	}
	size := int(binary.BigEndian.Uint32(hdr[:4]))
	pattern := patternName(binary.BigEndian.Uint32(hdr[4:]))
	buf := make([]byte, size)
	res.Pattern = pattern
	res.Size = size
	switch pattern {
	case "stream":
		n, _ := io.CopyBuffer(io.Discard, c, buf)
		res.Bytes = uint64(n)
	case "pingpong":
		for {
			if _, err := io.ReadFull(c, buf); err != nil {
				break
			}
			res.Messages++
			res.Bytes += uint64(len(buf))
			if _, err := c.Write(buf); err != nil {
				break
			}
		}
	case "duplex":
		var sent atomic.Uint64
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				n, err := c.Write(buf)
				sent.Add(uint64(n))
				if err != nil {
					return
				}
			}
		}()
		n, _ := io.CopyBuffer(io.Discard, c, buf)
		res.Bytes = uint64(n) + sent.Load()
		_ = c.Close()
		wg.Wait()
	default:
		res.Error = "unknown pattern"
	}
	finishResult(&res, time.Since(start))
	if jsonOut {
		writeJSON(res)
	} else {
		printResult(res)
	}
}

func tcpClient(args []string) {
	fs := flag.NewFlagSet("tcp", flag.ExitOnError)
	addr := fs.String("addr", "", "server address")
	pattern := fs.String("pattern", "stream", "stream, pingpong, or duplex")
	sizeText := fs.String("size", "1M", "payload size")
	durationText := fs.String("duration", "10s", "benchmark duration")
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)

	if *addr == "" {
		fatalf("-addr is required")
	}
	size := parseSize(*sizeText)
	duration := parseDuration(*durationText)
	res := runTCP(*addr, *pattern, size, duration)
	if *jsonOut {
		writeJSON(res)
	} else {
		printResult(res)
	}
	if res.Error != "" {
		os.Exit(1)
	}
}

func sweep(args []string) {
	fs := flag.NewFlagSet("sweep", flag.ExitOnError)
	addr := fs.String("addr", "", "server address")
	pattern := fs.String("pattern", "stream", "stream, pingpong, or duplex")
	sizesText := fs.String("sizes", "1,64,4K,64K,1M,16M", "comma-separated payload sizes")
	durationText := fs.String("duration", "10s", "duration per size")
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)

	if *addr == "" {
		fatalf("-addr is required")
	}
	duration := parseDuration(*durationText)
	var results []result
	for _, field := range strings.Split(*sizesText, ",") {
		field = strings.TrimSpace(field)
		if field == "" {
			continue
		}
		results = append(results, runTCP(*addr, *pattern, parseSize(field), duration))
	}
	if *jsonOut {
		writeJSON(results)
		return
	}
	for _, res := range results {
		printResult(res)
	}
}

func runTCP(addr, pattern string, size int, duration time.Duration) result {
	res := result{
		Mode:     "tcp",
		Pattern:  pattern,
		Addr:     addr,
		Duration: duration.String(),
		Size:     size,
	}
	c, err := dialTCP(addr)
	if err != nil {
		res.Error = err.Error()
		return res
	}
	defer c.Close()
	res.LocalAddr = c.LocalAddr().String()
	res.RemoteAddr = c.RemoteAddr().String()
	if err := writeHeader(c, size, pattern); err != nil {
		res.Error = err.Error()
		return res
	}
	payload := makePayload(size)
	start := time.Now()
	deadline := start.Add(duration)
	switch pattern {
	case "stream":
		for time.Now().Before(deadline) {
			n, err := c.Write(payload)
			res.Bytes += uint64(n)
			res.Messages++
			if err != nil {
				res.Error = err.Error()
				break
			}
		}
	case "pingpong":
		samples := make([]time.Duration, 0, 4096)
		reply := make([]byte, len(payload))
		for time.Now().Before(deadline) {
			t0 := time.Now()
			if _, err := c.Write(payload); err != nil {
				res.Error = err.Error()
				break
			}
			if _, err := io.ReadFull(c, reply); err != nil {
				res.Error = err.Error()
				break
			}
			samples = append(samples, time.Since(t0))
			res.Bytes += uint64(len(payload) * 2)
			res.Messages++
		}
		res.Latency = summarizeLatency(samples)
	case "duplex":
		var sent atomic.Uint64
		var recv atomic.Uint64
		var msgs atomic.Uint64
		errc := make(chan error, 2)
		done := make(chan struct{})
		go func() {
			for {
				select {
				case <-done:
					errc <- nil
					return
				default:
				}
				n, err := c.Write(payload)
				sent.Add(uint64(n))
				msgs.Add(1)
				if err != nil {
					errc <- err
					return
				}
			}
		}()
		go func() {
			buf := make([]byte, size)
			for {
				select {
				case <-done:
					errc <- nil
					return
				default:
				}
				n, err := c.Read(buf)
				recv.Add(uint64(n))
				if err != nil {
					errc <- err
					return
				}
			}
		}()
		time.Sleep(duration)
		close(done)
		_ = c.Close()
		for i := 0; i < 2; i++ {
			if err := <-errc; err != nil && !isClosedConnError(err) {
				res.Error = err.Error()
			}
		}
		res.Bytes = recv.Load() + sent.Load()
		res.Messages = msgs.Load()
	default:
		res.Error = "unknown pattern"
	}
	finishResult(&res, time.Since(start))
	return res
}

func listenTCP(addr string) (net.Listener, error) {
	return net.Listen(tcpNetwork(addr), addr)
}

func dialTCP(addr string) (net.Conn, error) {
	return net.Dial(tcpNetwork(addr), addr)
}

func tcpNetwork(addr string) string {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return "tcp"
	}
	ip := net.ParseIP(host)
	if ip != nil && ip.To4() != nil {
		return "tcp4"
	}
	return "tcp"
}

func finishResult(res *result, elapsed time.Duration) {
	res.Elapsed = elapsed.String()
	if elapsed > 0 {
		res.BytesPerSec = float64(res.Bytes) / elapsed.Seconds()
		res.MsgsPerSec = float64(res.Messages) / elapsed.Seconds()
	}
}

func rdmaProbe(args []string) {
	fs := flag.NewFlagSet("rdma-probe", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	timeout := fs.Duration("timeout", 10*time.Second, "maximum time to wait for RDMA provider probe calls")
	fs.Parse(args)

	stopWatchdog := startRDMAWatchdog("rdma-probe", *timeout)
	defer stopWatchdog()
	res := result{Mode: "rdma-probe", RDMA: probeRDMA()}
	if *jsonOut {
		writeJSON(res)
	} else {
		printResult(res)
	}
}

func startRDMAWatchdog(name string, timeout time.Duration) func() {
	if timeout <= 0 {
		return func() {}
	}
	timer := time.AfterFunc(timeout, func() {
		fmt.Fprintf(os.Stderr, "rdmaperf: %s watchdog expired after %s\n", name, timeout)
		os.Exit(124)
	})
	return func() {
		timer.Stop()
	}
}

func rdmaPingpong(args []string) {
	fs := flag.NewFlagSet("rdma-pingpong", flag.ExitOnError)
	listenAddr := fs.String("listen", "", "listen address for server role")
	addr := fs.String("addr", "", "server address for client role")
	deviceName := fs.String("name", "", "select first RDMA device whose name contains substring")
	deviceIndex := fs.Int("device", -1, "RDMA device index")
	gidIndex := fs.Int("gid-index", -1, "source GID index for GRH; -1 auto-selects, 0..255 selects explicitly")
	sizeText := fs.String("size", "64", "payload size")
	iters := fs.Int("iters", 10000, "ping-pong iterations")
	setupTimeout := fs.Duration("setup-timeout", 5*time.Second, "maximum time to wait for local QP setup")
	jsonOut := fs.Bool("json", false, "print JSON")
	allowRTR := fs.Bool("allow-rtr", false, "allow QP INIT->RTR transition; may wedge Apple Thunderbolt RDMA ports")
	fs.Parse(args)

	if (*listenAddr == "") == (*addr == "") {
		fatalf("exactly one of -listen or -addr is required")
	}
	if err := validateGIDIndexFlag(*gidIndex); err != nil {
		fatalf("%v", err)
	}
	if *iters <= 0 {
		fatalf("-iters must be positive")
	}
	if err := checkRDMAPingpongOptIn(*allowRTR); err != nil {
		fatalf("%v", err)
	}
	size := parseSize(*sizeText)
	var res rdmaBenchResult
	if *listenAddr != "" {
		res = runRDMAPingpongServer(*listenAddr, *deviceName, *deviceIndex, *gidIndex, size, *iters, *setupTimeout)
	} else {
		res = runRDMAPingpongClient(*addr, *deviceName, *deviceIndex, *gidIndex, size, *iters, *setupTimeout)
	}
	if *jsonOut {
		writeJSON(res)
	} else {
		printRDMABench(res)
	}
	if res.Error != "" {
		os.Exit(1)
	}
}

func checkRDMAPingpongOptIn(allowRTR bool) error {
	return xrdma.RequireRTRAttemptAllowed(allowRTR)
}

func validateGIDIndexFlag(index int) error {
	if index == -1 {
		return nil
	}
	if index < 0 {
		return fmt.Errorf("-gid-index must be -1 or between 0 and 255")
	}
	if index > 255 {
		return fmt.Errorf("-gid-index must be between 0 and 255")
	}
	return nil
}

type rdmaResources struct {
	dev      rdma.Device
	ctx      rdma.RDMAContext
	pd       rdma.RDMAPD
	cq       rdma.RDMACQ
	qp       rdma.RDMAQP
	mr       rdma.RDMAMR
	poller   rdma.IbvCQPoller
	poster   rdma.IbvQPPoster
	mapBuf   []byte
	buf      []byte
	port     ibvPortAttr
	gid      rdma.IbvGID
	gidIndex int
	psn      uint32
}

type routeGID struct {
	index int
	gid   rdma.IbvGID
}

func runRDMAPingpongServer(listenAddr, deviceName string, deviceIndex, gidIndex, size, iters int, setupTimeout time.Duration) rdmaBenchResult {
	res := rdmaBenchResult{Mode: "rdma-pingpong", Role: "server", Addr: listenAddr, Size: size, Iterations: iters}
	ln, err := listenTCP(listenAddr)
	if err != nil {
		res.Error = err.Error()
		return res
	}
	defer ln.Close()
	c, err := ln.Accept()
	if err != nil {
		res.Error = err.Error()
		return res
	}
	defer c.Close()
	res.Addr = c.LocalAddr().String()
	if err := runRDMAPingpong(c, false, deviceName, deviceIndex, gidIndex, size, iters, setupTimeout, &res); err != nil {
		res.Error = err.Error()
	}
	return res
}

func runRDMAPingpongClient(addr, deviceName string, deviceIndex, gidIndex, size, iters int, setupTimeout time.Duration) rdmaBenchResult {
	res := rdmaBenchResult{Mode: "rdma-pingpong", Role: "client", Addr: addr, Size: size, Iterations: iters}
	c, err := dialTCP(addr)
	if err != nil {
		res.Error = err.Error()
		return res
	}
	defer c.Close()
	if err := runRDMAPingpong(c, true, deviceName, deviceIndex, gidIndex, size, iters, setupTimeout, &res); err != nil {
		res.Error = err.Error()
	}
	return res
}

func runRDMAPingpong(c net.Conn, client bool, deviceName string, deviceIndex, gidIndex, size, iters int, setupTimeout time.Duration, res *rdmaBenchResult) error {
	defer c.SetDeadline(time.Time{})
	role := "server"
	if client {
		role = "client"
	}
	control := newRDMAControlConn(c)
	if err := runRDMAControlHello(control, res, "pre-resource-control", role, setupTimeout); err != nil {
		return err
	}
	res.Stage = "open-rdma-resources"
	r, err := openRDMAResourcesWithTimeout(deviceName, deviceIndex, gidIndex, size, setupTimeout)
	if err != nil {
		return err
	}
	closeResources := true
	defer func() {
		if closeResources {
			r.close()
		}
	}()

	local := r.peerInfo()
	res.Device = r.dev.Name
	res.Local = local
	if err := runRDMAControlHello(control, res, "post-resource-control", role, setupTimeout); err != nil {
		return err
	}
	res.Stage = "exchange-rdma-info"
	remote, err := exchangeRDMAPeerInfo(control, local, setupTimeout)
	res.Remote = remote
	if err != nil {
		return err
	}

	res.Stage = "connect-rdma"
	connectErr := r.connectWithTimeout(remote, setupTimeout)
	if errors.Is(connectErr, errRDMASetupTimeout) {
		closeResources = false
	}
	// This ready exchange is the post-RTS barrier: neither side posts datapath
	// work until both QPs have completed the INIT->RTR->RTS transition.
	res.Stage = "exchange-rdma-ready"
	if err := exchangeRDMAReady(control, connectErr, setupTimeout); err != nil {
		return err
	}
	if connectErr != nil {
		return connectErr
	}
	_ = c.SetDeadline(time.Time{})
	res.Stage = "datapath"
	if client {
		err = runRDMAPingpongClientLoop(r, iters, res)
	} else {
		err = runRDMAPingpongServerLoop(r, iters, res)
	}
	if err != nil {
		return err
	}
	res.Stage = "done"
	return nil
}

func runRDMAControlHello(c *rdmaControlConn, res *rdmaBenchResult, stage, role string, timeout time.Duration) error {
	res.Stage = stage
	remote, err := exchangeRDMAControlHello(c, rdmaControlHello{Stage: stage, Role: role}, timeout)
	event := rdmaControlEvent{Stage: stage, OK: err == nil}
	if err != nil {
		event.Error = err.Error()
		res.Control = append(res.Control, event)
		return fmt.Errorf("%s: %w", stage, err)
	}
	event.RemoteStage = remote.Stage
	event.RemoteRole = remote.Role
	if remote.Stage != stage {
		event.OK = false
		event.Error = fmt.Sprintf("remote control hello stage %q, want %q", remote.Stage, stage)
		res.Control = append(res.Control, event)
		return fmt.Errorf("%s: %s", stage, event.Error)
	}
	wantRole := "client"
	if role == "client" {
		wantRole = "server"
	}
	if remote.Role != wantRole {
		event.OK = false
		event.Error = fmt.Sprintf("remote control hello role %q, want %q", remote.Role, wantRole)
		res.Control = append(res.Control, event)
		return fmt.Errorf("%s: %s", stage, event.Error)
	}
	res.Control = append(res.Control, event)
	return nil
}

func exchangeRDMAControlHello(c *rdmaControlConn, local rdmaControlHello, timeout time.Duration) (rdmaControlHello, error) {
	if err := setControlDeadline(c, timeout); err != nil {
		return rdmaControlHello{}, err
	}
	sendc := make(chan error, 1)
	go func() {
		sendc <- c.enc.Encode(local)
	}()
	var remote rdmaControlHello
	recvErr := c.dec.Decode(&remote)
	sendErr := <-sendc
	if recvErr != nil {
		return rdmaControlHello{}, fmt.Errorf("receive remote control hello: %w", recvErr)
	}
	if sendErr != nil {
		return rdmaControlHello{}, fmt.Errorf("send local control hello: %w", sendErr)
	}
	return remote, nil
}

func exchangeRDMAPeerInfo(c *rdmaControlConn, local rdmaPeerInfo, timeout time.Duration) (rdmaPeerInfo, error) {
	if err := setControlDeadline(c, timeout); err != nil {
		return rdmaPeerInfo{}, err
	}
	sendc := make(chan error, 1)
	go func() {
		sendc <- c.enc.Encode(local)
	}()
	var remote rdmaPeerInfo
	recvErr := c.dec.Decode(&remote)
	sendErr := <-sendc
	if recvErr != nil {
		return rdmaPeerInfo{}, fmt.Errorf("receive remote rdma info: %w", recvErr)
	}
	if sendErr != nil {
		return rdmaPeerInfo{}, fmt.Errorf("send local rdma info: %w", sendErr)
	}
	if err := validateRDMAPeerInfo(remote); err != nil {
		return remote, fmt.Errorf("invalid remote rdma info: %w", err)
	}
	return remote, nil
}

func validateRDMAPeerInfo(info rdmaPeerInfo) error {
	if info.QPN == 0 {
		return fmt.Errorf("missing qpn")
	}
	if info.GIDIndex < 0 {
		return fmt.Errorf("invalid gid index %d", info.GIDIndex)
	}
	if !info.UseGlobal {
		return nil
	}
	if _, err := parseGID(info.GID); err != nil {
		return fmt.Errorf("invalid gid: %w", err)
	}
	return nil
}

func exchangeRDMAReady(c *rdmaControlConn, localErr error, timeout time.Duration) error {
	if err := setControlDeadline(c, timeout); err != nil {
		return err
	}
	local := rdmaReadyStatus{OK: localErr == nil}
	if localErr != nil {
		local.Error = localErr.Error()
	}
	var remote rdmaReadyStatus
	sendc := make(chan error, 1)
	go func() {
		sendc <- c.enc.Encode(local)
	}()
	recvErr := c.dec.Decode(&remote)
	sendErr := <-sendc
	if recvErr != nil {
		return fmt.Errorf("receive rdma ready status: %w", recvErr)
	}
	if sendErr != nil {
		return fmt.Errorf("send rdma ready status: %w", sendErr)
	}
	if !remote.OK {
		return fmt.Errorf("remote rdma setup failed: %s", remote.Error)
	}
	return nil
}

func setControlDeadline(c *rdmaControlConn, timeout time.Duration) error {
	if timeout <= 0 {
		return c.conn.SetDeadline(time.Time{})
	}
	return c.conn.SetDeadline(time.Now().Add(timeout + 2*time.Second))
}

func openRDMAResources(deviceName string, deviceIndex, gidIndex, size int, probeTimeout time.Duration) (*rdmaResources, error) {
	if !rdma.Available() {
		return nil, fmt.Errorf("rdma unavailable")
	}
	devs, err := rdmaDevices()
	if err != nil {
		return nil, fmt.Errorf("rdma devices: %w", err)
	}
	if len(devs) == 0 {
		return nil, fmt.Errorf("no rdma devices")
	}
	dev, err := selectRDMADevice(devs, deviceName, deviceIndex, probeTimeout)
	if err != nil {
		return nil, err
	}
	ctx, err := dev.Open()
	if err != nil {
		return nil, fmt.Errorf("ibv_open_device: %w", err)
	}
	r := &rdmaResources{dev: dev, ctx: ctx, gidIndex: -1, psn: 7}
	success := false
	defer func() {
		if !success {
			r.close()
		}
	}()

	if err = r.queryPortAndGID(gidIndex); err != nil {
		return nil, err
	}
	if r.port.State != 4 {
		return nil, inactivePortError(dev.Name, r.port.State, deviceName != "" || deviceIndex >= 0)
	}
	r.pd, err = rdma.IbvAllocPd(ctx)
	if err != nil || r.pd == 0 {
		return nil, fmt.Errorf("ibv_alloc_pd: %s", nilReturnError(err, r.pd, "protection domain"))
	}
	r.cq, err = rdma.IbvCreateCq(ctx, 64, 0, 0, 0)
	if err != nil || r.cq == 0 {
		return nil, fmt.Errorf("ibv_create_cq: %s", nilReturnError(err, r.cq, "completion queue"))
	}
	r.buf, r.mapBuf, err = rdmaBuffer(size)
	if err != nil {
		return nil, err
	}
	r.mr, err = rdma.IbvRegMr(r.pd, uintptr(unsafe.Pointer(unsafe.SliceData(r.buf))), uintptr(len(r.buf)), rdma.IBV_ACCESS_LOCAL_WRITE|rdma.IBV_ACCESS_REMOTE_READ|rdma.IBV_ACCESS_REMOTE_WRITE)
	runtime.KeepAlive(r.buf)
	if err != nil || r.mr == 0 {
		return nil, fmt.Errorf("ibv_reg_mr: %s", nilReturnError(err, r.mr, "memory region"))
	}
	init := rdma.IbvQPInitAttr{
		SendCQ: r.cq,
		RecvCQ: r.cq,
		Cap: rdma.IbvQPCap{
			MaxSendWR:  32,
			MaxRecvWR:  32,
			MaxSendSGE: 1,
			MaxRecvSGE: 1,
		},
		QPType:   rdma.IBV_QPT_UC,
		SQSigAll: 0,
	}
	r.qp, err = rdma.IbvCreateQpAttr(r.pd, &init)
	if err != nil || r.qp == 0 {
		return nil, fmt.Errorf("ibv_create_qp: %s", nilReturnError(err, r.qp, "queue pair"))
	}
	r.poller, err = rdma.NewIbvCQPoller(r.cq)
	if err != nil {
		return nil, err
	}
	r.poster, err = rdma.NewIbvQPPoster(r.qp)
	if err != nil {
		return nil, err
	}
	success = true
	return r, nil
}

func inactivePortError(name string, state int32, explicit bool) error {
	msg := fmt.Sprintf("%s port 1 is %s, need PORT_ACTIVE", name, portStateName(state))
	if explicit {
		msg += "; selected explicitly, try omitting -name/-device to auto-select an active RDMA device"
	}
	return fmt.Errorf("%s", msg)
}

func selectRDMADevice(devs []rdma.Device, name string, index int, probeTimeout time.Duration) (rdma.Device, error) {
	if name != "" {
		for _, dev := range devs {
			if strings.Contains(dev.Name, name) {
				return dev, nil
			}
		}
		return rdma.Device{}, fmt.Errorf("no rdma device name contains %q", name)
	}
	if index >= 0 {
		if index >= len(devs) {
			return rdma.Device{}, fmt.Errorf("rdma device index %d out of range [0,%d)", index, len(devs))
		}
		return devs[index], nil
	}
	var states []string
	for _, dev := range devs {
		state, err := queryRDMADevicePortState(dev.Name, probeTimeout)
		if err != nil {
			states = append(states, fmt.Sprintf("%s: %v", dev.Name, err))
			continue
		}
		if state == 4 {
			return dev, nil
		}
		states = append(states, fmt.Sprintf("%s: %s", dev.Name, portStateName(state)))
	}
	return rdma.Device{}, fmt.Errorf("no PORT_ACTIVE RDMA device found; probed %s", strings.Join(states, ", "))
}

func rdmaDevices() ([]rdma.Device, error) {
	devs, err := rdma.Devices()
	if err == nil && len(devs) == 0 {
		devs, err = rdma.Devices()
	}
	return devs, err
}

func queryRDMADevicePortState(name string, timeout time.Duration) (int32, error) {
	if timeout <= 0 {
		return queryRDMADevicePortStateLocal(name)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Args[0], "rdma-port-state", "-name", name)
	out, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return 0, fmt.Errorf("query_port timed out after %s", timeout)
	}
	var got struct {
		State int32  `json:"state"`
		Error string `json:"error,omitempty"`
	}
	if err := json.Unmarshal(out, &got); err != nil {
		return 0, fmt.Errorf("decode port state: %w", err)
	}
	if err != nil && got.Error == "" {
		return 0, err
	}
	if got.Error != "" {
		return 0, fmt.Errorf("%s", got.Error)
	}
	return got.State, nil
}

func queryRDMADevicePortStateLocal(name string) (int32, error) {
	devs, err := rdmaDevices()
	if err != nil {
		return 0, fmt.Errorf("rdma devices: %w", err)
	}
	for _, dev := range devs {
		if dev.Name != name {
			continue
		}
		ctx, err := dev.Open()
		if err != nil || ctx == 0 {
			return 0, fmt.Errorf("open failed: %v", err)
		}
		var attr ibvPortAttr
		rc, qerr := rdma.IbvQueryPort(ctx, 1, uintptr(unsafe.Pointer(&attr)))
		runtime.KeepAlive(&attr)
		_, _ = rdma.IbvCloseDevice(ctx)
		if qerr != nil || rc != 0 {
			return 0, fmt.Errorf("query_port failed: %s", errOrCode(qerr, rc))
		}
		return attr.State, nil
	}
	return 0, fmt.Errorf("not found")
}

func rdmaPortState(args []string) {
	fs := flag.NewFlagSet("rdma-port-state", flag.ExitOnError)
	name := fs.String("name", "", "RDMA device name")
	fs.Parse(args)
	if *name == "" {
		fatalf("-name is required")
	}
	var out struct {
		State int32  `json:"state"`
		Error string `json:"error,omitempty"`
	}
	state, err := queryRDMADevicePortStateLocal(*name)
	if err != nil {
		out.Error = err.Error()
		writeJSON(out)
		os.Exit(1)
	}
	out.State = state
	writeJSON(out)
}

func openRDMAResourcesWithTimeout(deviceName string, deviceIndex, gidIndex, size int, timeout time.Duration) (*rdmaResources, error) {
	if timeout <= 0 {
		return openRDMAResources(deviceName, deviceIndex, gidIndex, size, timeout)
	}
	type result struct {
		res *rdmaResources
		err error
	}
	done := make(chan result, 1)
	go func() {
		res, err := openRDMAResources(deviceName, deviceIndex, gidIndex, size, timeout)
		done <- result{res: res, err: err}
	}()
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case got := <-done:
		return got.res, got.err
	case <-timer.C:
		return nil, fmt.Errorf("%w opening rdma resources after %s", errRDMASetupTimeout, timeout)
	}
}

func (r *rdmaResources) queryPortAndGID(preferredGIDIndex int) error {
	rc, err := rdma.IbvQueryPort(r.ctx, 1, uintptr(unsafe.Pointer(&r.port)))
	runtime.KeepAlive(&r.port)
	if err != nil || rc != 0 {
		return fmt.Errorf("ibv_query_port: %s", errOrCode(err, rc))
	}
	if preferredGIDIndex >= 0 {
		gid, err := r.queryGID(preferredGIDIndex)
		if err != nil {
			return err
		}
		r.gid = gid
		r.gidIndex = preferredGIDIndex
		return nil
	}
	limit := routeGIDScanLimit(r.port.GIDTblLen)
	gids := make([]routeGID, 0, 2)
	for i := 0; i < limit; i++ {
		var gid rdma.IbvGID
		rc, err := rdma.IbvQueryGidInto(r.ctx, 1, i, &gid)
		if err != nil || rc != 0 {
			if len(gids) > 0 {
				break
			}
			return fmt.Errorf("ibv_query_gid[%d]: %s", i, errOrCode(err, rc))
		}
		if !isZeroGID(gid) {
			gids = append(gids, routeGID{index: i, gid: gid})
		}
	}
	gid, index, ok := selectRouteGID(gids, preferredGIDIndex, r.port.LinkLayer)
	if !ok {
		if preferredGIDIndex >= 0 {
			return fmt.Errorf("gid index %d is not available or is zero", preferredGIDIndex)
		}
		return fmt.Errorf("auto-select route gid: no safe nonzero gid found")
	}
	r.gid = gid
	r.gidIndex = index
	return nil
}

func (r *rdmaResources) queryGID(index int) (rdma.IbvGID, error) {
	var gid rdma.IbvGID
	rc, err := rdma.IbvQueryGidInto(r.ctx, 1, index, &gid)
	if err != nil || rc != 0 {
		return gid, fmt.Errorf("ibv_query_gid[%d]: %s", index, errOrCode(err, rc))
	}
	if isZeroGID(gid) {
		return gid, fmt.Errorf("gid index %d is zero", index)
	}
	return gid, nil
}

func selectRouteGID(gids []routeGID, preferred int, linkLayer uint8) (rdma.IbvGID, int, bool) {
	entries := make([]xrdma.RouteGID, 0, len(gids))
	for _, entry := range gids {
		entries = append(entries, xrdma.RouteGID{Index: entry.index, GID: entry.gid})
	}
	entry, ok := xrdma.SelectRouteGID(entries, preferred, linkLayer)
	if !ok {
		return rdma.IbvGID{}, -1, false
	}
	return entry.GID, entry.Index, true
}

func routeGIDScanLimit(tableLen int32) int {
	return xrdma.RouteGIDScanLimit(tableLen)
}

func (r *rdmaResources) peerInfo() rdmaPeerInfo {
	return rdmaPeerInfo{
		Name:      r.dev.Name,
		LID:       r.port.LID,
		QPN:       rdma.Ibv_qp_num(r.qp),
		PSN:       r.psn,
		GIDIndex:  r.gidIndex,
		GID:       fmt.Sprintf("%x", r.gid[:]),
		UseGlobal: !isZeroGID(r.gid),
		ActiveMTU: r.port.ActiveMTU,
	}
}

func (r *rdmaResources) connect(remote rdmaPeerInfo) error {
	init := rdma.IbvQPAttr{
		QPState:       rdma.IBV_QPS_INIT,
		PKeyIndex:     0,
		PortNum:       1,
		QPAccessFlags: rdma.IBV_ACCESS_LOCAL_WRITE | rdma.IBV_ACCESS_REMOTE_READ | rdma.IBV_ACCESS_REMOTE_WRITE,
	}
	rc, err := rdma.IbvModifyQpAttr(r.qp, &init, rdma.IBV_QP_STATE|rdma.IBV_QP_PKEY_INDEX|rdma.IBV_QP_PORT|rdma.IBV_QP_ACCESS_FLAGS)
	if err != nil || rc != 0 {
		return fmt.Errorf("modify qp INIT: %s", errOrCode(err, rc))
	}

	rtr, err := r.rtrAttr(remote)
	if err != nil {
		return err
	}
	rc, err = rdma.IbvModifyQpAttr(r.qp, &rtr, rdma.IBV_QP_STATE|rdma.IBV_QP_AV|rdma.IBV_QP_PATH_MTU|rdma.IBV_QP_DEST_QPN|rdma.IBV_QP_RQ_PSN)
	if err != nil || rc != 0 {
		return fmt.Errorf("modify qp RTR: %s", errOrCode(err, rc))
	}

	rts := rdma.IbvQPAttr{
		QPState: rdma.IBV_QPS_RTS,
		SQPSN:   r.psn,
	}
	rc, err = rdma.IbvModifyQpAttr(r.qp, &rts, rdma.IBV_QP_STATE|rdma.IBV_QP_SQ_PSN)
	if err != nil || rc != 0 {
		return fmt.Errorf("modify qp RTS: %s", errOrCode(err, rc))
	}
	return nil
}

func (r *rdmaResources) rtrAttr(remote rdmaPeerInfo) (rdma.IbvQPAttr, error) {
	attr := rdma.IbvQPAttr{
		QPState:   rdma.IBV_QPS_RTR,
		PathMTU:   negotiatedPathMTU(r.port.ActiveMTU, remote.ActiveMTU),
		RQPSN:     remote.PSN,
		DestQPNum: remote.QPN,
		AHAttr: rdma.IbvAHAttr{
			DLID:     remote.LID,
			PortNum:  1,
			IsGlobal: boolByte(remote.UseGlobal),
		},
	}
	if !remote.UseGlobal {
		return attr, nil
	}
	if r.gidIndex < 0 || r.gidIndex > 255 {
		return attr, fmt.Errorf("local gid index %d out of uint8 range", r.gidIndex)
	}
	gid, err := parseGID(remote.GID)
	if err != nil {
		return attr, err
	}
	attr.AHAttr.GRH.DGID = gid
	attr.AHAttr.GRH.SGIDIndex = uint8(r.gidIndex)
	attr.AHAttr.GRH.HopLimit = 1
	return attr, nil
}

func negotiatedPathMTU(local, remote int32) int32 {
	if mtuBytes(local) == 0 {
		local = rdma.IBV_MTU_1024
	}
	if mtuBytes(remote) == 0 {
		remote = rdma.IBV_MTU_1024
	}
	if local < remote {
		return local
	}
	return remote
}

func (r *rdmaResources) connectWithTimeout(remote rdmaPeerInfo, timeout time.Duration) error {
	if timeout <= 0 {
		return r.connect(remote)
	}
	done := make(chan error, 1)
	go func() {
		done <- r.connect(remote)
	}()
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case err := <-done:
		return err
	case <-timer.C:
		return fmt.Errorf("%w after %s", errRDMASetupTimeout, timeout)
	}
}

func runRDMAPingpongClientLoop(r *rdmaResources, iters int, res *rdmaBenchResult) error {
	samples := make([]time.Duration, 0, iters)
	start := time.Now()
	for i := 0; i < iters; i++ {
		if err := r.postRecv(uint64(i)); err != nil {
			return err
		}
		t0 := time.Now()
		if err := r.postSend(uint64(i)); err != nil {
			return err
		}
		if err := r.poll(2, 5*time.Second); err != nil {
			return err
		}
		samples = append(samples, time.Since(t0))
	}
	finishRDMABench(res, time.Since(start), uint64(iters*len(r.buf)*2), uint64(iters))
	res.Latency = summarizeLatency(samples)
	return nil
}

func runRDMAPingpongServerLoop(r *rdmaResources, iters int, res *rdmaBenchResult) error {
	start := time.Now()
	for i := 0; i < iters; i++ {
		if err := r.postRecv(uint64(i)); err != nil {
			return err
		}
		if err := r.poll(1, 5*time.Second); err != nil {
			return err
		}
		if err := r.postSend(uint64(i)); err != nil {
			return err
		}
		if err := r.poll(1, 5*time.Second); err != nil {
			return err
		}
	}
	finishRDMABench(res, time.Since(start), uint64(iters*len(r.buf)*2), uint64(iters))
	return nil
}

func (r *rdmaResources) postRecv(id uint64) error {
	sge := rdma.IbvSGE{Addr: uint64(uintptr(unsafe.Pointer(unsafe.SliceData(r.buf)))), Length: uint32(len(r.buf)), LKey: rdma.Ibv_mr_lkey(r.mr)}
	wr := rdma.IbvRecvWR{WRID: id, SGList: &sge, NumSGE: 1}
	var bad *rdma.IbvRecvWR
	rc := r.poster.PostRecv(&wr, &bad)
	if rc != 0 {
		return fmt.Errorf("ibv_post_recv: %s", errOrCode(nil, rc))
	}
	return nil
}

func (r *rdmaResources) postSend(id uint64) error {
	sge := rdma.IbvSGE{Addr: uint64(uintptr(unsafe.Pointer(unsafe.SliceData(r.buf)))), Length: uint32(len(r.buf)), LKey: rdma.Ibv_mr_lkey(r.mr)}
	wr := rdma.IbvSendWR{WRID: id, SGList: &sge, NumSGE: 1, Opcode: rdma.IBV_WR_SEND, SendFlags: rdma.IBV_SEND_SIGNALED}
	var bad *rdma.IbvSendWR
	rc := r.poster.PostSend(&wr, &bad)
	if rc != 0 {
		return fmt.Errorf("ibv_post_send: %s", errOrCode(nil, rc))
	}
	return nil
}

func (r *rdmaResources) poll(want int, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	var got int
	for got < want {
		var wc rdma.IbvWC
		n := r.poller.Poll(1, &wc)
		if n < 0 {
			return fmt.Errorf("ibv_poll_cq: rc=%d", n)
		}
		if n == 0 {
			if time.Now().After(deadline) {
				return fmt.Errorf("timed out polling cq after %s", timeout)
			}
			runtimeYield()
			continue
		}
		if wc.Status != rdma.IBV_WC_SUCCESS {
			return fmt.Errorf("work completion status=%d opcode=%d vendor_err=%d", wc.Status, wc.Opcode, wc.VendorErr)
		}
		got += n
	}
	return nil
}

func (r *rdmaResources) close() {
	if r.qp != 0 {
		_, _ = rdma.IbvDestroyQp(r.qp)
	}
	if r.mr != 0 {
		_, _ = rdma.IbvDeregMr(r.mr)
	}
	if r.mapBuf != nil {
		_ = syscall.Munmap(r.mapBuf)
	}
	if r.cq != 0 {
		_, _ = rdma.IbvDestroyCq(r.cq)
	}
	if r.pd != 0 {
		_, _ = rdma.IbvDeallocPd(r.pd)
	}
	if r.ctx != 0 {
		_, _ = rdma.IbvCloseDevice(r.ctx)
	}
}

func finishRDMABench(res *rdmaBenchResult, elapsed time.Duration, bytes, messages uint64) {
	res.Elapsed = elapsed.String()
	res.Bytes = bytes
	if elapsed > 0 {
		res.BytesPerSec = float64(bytes) / elapsed.Seconds()
		res.MessagesPerSec = float64(messages) / elapsed.Seconds()
	}
}

func probeRDMA() *rdmaSummary {
	sum := &rdmaSummary{Available: rdma.Available()}
	if !sum.Available {
		sum.Notes = append(sum.Notes, "librdma or required symbols are unavailable")
		return sum
	}
	devs, err := rdma.Devices()
	if len(devs) == 0 && err == nil {
		devs, err = rdma.Devices()
		if len(devs) > 0 {
			sum.Notes = append(sum.Notes, "initial device-list call was empty; retry returned devices")
		}
	}
	if err != nil {
		sum.Notes = append(sum.Notes, "rdma devices: "+err.Error())
		return sum
	}
	if len(devs) == 0 {
		sum.Notes = append(sum.Notes, "librdma reported no RDMA devices")
	}
	for i, dev := range devs {
		sum.Devices = append(sum.Devices, rdmaDevice{Index: i, Name: dev.Name, NetInterface: rdmaNetInterface(dev.Name), Handle: hexHandle(dev.Handle)})
		ctx, err := dev.Open()
		sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_open_device", i), OK: err == nil && ctx != 0, Handle: hexHandle(ctx), Error: errString(err)})
		if ctx == 0 {
			continue
		}
		buf := make([]byte, 4096)
		rc, err := rdma.IbvQueryDevice(ctx, uintptr(unsafe.Pointer(unsafe.SliceData(buf))))
		runtime.KeepAlive(buf)
		sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_query_device", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)})
		portBuf := make([]byte, unsafe.Sizeof(ibvPortAttr{}))
		rc, err = rdma.IbvQueryPort(ctx, 1, uintptr(unsafe.Pointer(unsafe.SliceData(portBuf))))
		runtime.KeepAlive(portBuf)
		portStep := rdmaStep{Name: fmt.Sprintf("device[%d].ibv_query_port", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)}
		if portStep.OK {
			portStep.Fields = portAttrFields(portBuf)
		}
		sum.Steps = append(sum.Steps, portStep)
		if !portStep.OK || portStep.Fields["state"] != int32(4) {
			sum.Steps = append(sum.Steps, rdmaStep{
				Name:  fmt.Sprintf("device[%d].rdma_datapath_ready", i),
				Error: "skipped resource verbs because port 1 is not PORT_ACTIVE",
			})
			rc, err = rdma.IbvCloseDevice(ctx)
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_close_device", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)})
			continue
		}
		pd, err := rdma.IbvAllocPd(ctx)
		sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_alloc_pd", i), OK: err == nil && pd != 0, Handle: hexHandle(pd), Error: nilReturnError(err, pd, "protection domain")})
		cq, err := rdma.IbvCreateCq(ctx, 16, 0, 0, 0)
		sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_create_cq", i), OK: err == nil && cq != 0, Handle: hexHandle(cq), Error: nilReturnError(err, cq, "completion queue")})
		if cq != 0 {
			rc, err := rdma.IbvDestroyCq(cq)
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_destroy_cq", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)})
		}
		if pd != 0 {
			rc, err := rdma.IbvDeallocPd(pd)
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_dealloc_pd", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)})
		} else {
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_reg_mr", i), Error: "skipped because ibv_alloc_pd failed"})
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_create_qp", i), Error: "skipped because ibv_alloc_pd failed"})
		}
		rc, err = rdma.IbvCloseDevice(ctx)
		sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_close_device", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)})
	}
	sum.Notes = append(sum.Notes, "true RDMA datapath benchmarking requires successful PD, CQ, MR, and QP lifecycle")
	return sum
}

func interfaces(args []string) {
	fs := flag.NewFlagSet("interfaces", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)
	type iface struct {
		Name  string   `json:"name"`
		Flags string   `json:"flags"`
		Addrs []string `json:"addrs,omitempty"`
	}
	var out []iface
	ifaces, err := net.Interfaces()
	if err != nil {
		fatalf("interfaces: %v", err)
	}
	for _, ni := range ifaces {
		addrs, _ := ni.Addrs()
		item := iface{Name: ni.Name, Flags: ni.Flags.String()}
		for _, addr := range addrs {
			item.Addrs = append(item.Addrs, addr.String())
		}
		out = append(out, item)
	}
	if *jsonOut {
		writeJSON(out)
		return
	}
	for _, item := range out {
		fmt.Printf("%s\t%s", item.Name, item.Flags)
		if len(item.Addrs) > 0 {
			fmt.Printf("\t%s", strings.Join(item.Addrs, ", "))
		}
		fmt.Println()
	}
}

func writeHeader(w io.Writer, size int, pattern string) error {
	var hdr [headerSize]byte
	binary.BigEndian.PutUint32(hdr[:4], uint32(size))
	binary.BigEndian.PutUint32(hdr[4:], patternID(pattern))
	_, err := w.Write(hdr[:])
	return err
}

func patternID(pattern string) uint32 {
	switch pattern {
	case "stream":
		return 1
	case "pingpong":
		return 2
	case "duplex":
		return 3
	default:
		fatalf("unknown pattern %q", pattern)
	}
	return 0
}

func patternName(id uint32) string {
	switch id {
	case 1:
		return "stream"
	case 2:
		return "pingpong"
	case 3:
		return "duplex"
	default:
		return "unknown"
	}
}

func makePayload(size int) []byte {
	if size <= 0 {
		fatalf("size must be positive")
	}
	buf := make([]byte, size)
	for i := range buf {
		buf[i] = byte(i)
	}
	return buf
}

func parseSize(text string) int {
	text = strings.TrimSpace(strings.ToUpper(text))
	mult := 1
	switch {
	case strings.HasSuffix(text, "K"):
		mult = 1024
		text = strings.TrimSuffix(text, "K")
	case strings.HasSuffix(text, "M"):
		mult = 1024 * 1024
		text = strings.TrimSuffix(text, "M")
	case strings.HasSuffix(text, "G"):
		mult = 1024 * 1024 * 1024
		text = strings.TrimSuffix(text, "G")
	}
	var n int
	if _, err := fmt.Sscanf(text, "%d", &n); err != nil || n <= 0 {
		fatalf("invalid size %q", text)
	}
	return n * mult
}

func parseDuration(text string) time.Duration {
	d, err := time.ParseDuration(text)
	if err != nil || d <= 0 {
		fatalf("invalid duration %q", text)
	}
	return d
}

func summarizeLatency(samples []time.Duration) *latencySummary {
	if len(samples) == 0 {
		return nil
	}
	sort.Slice(samples, func(i, j int) bool { return samples[i] < samples[j] })
	return &latencySummary{
		Count: int64(len(samples)),
		Min:   samples[0].String(),
		P50:   percentile(samples, 50).String(),
		P95:   percentile(samples, 95).String(),
		P99:   percentile(samples, 99).String(),
		Max:   samples[len(samples)-1].String(),
	}
}

func percentile(samples []time.Duration, p int) time.Duration {
	idx := (len(samples)*p + 99) / 100
	if idx <= 0 {
		idx = 1
	}
	if idx > len(samples) {
		idx = len(samples)
	}
	return samples[idx-1]
}

func printResult(res result) {
	if res.RDMA != nil {
		printRDMA(res.RDMA)
		return
	}
	if res.Error != "" {
		fmt.Printf("%s %s addr=%s local=%s remote=%s size=%d error=%s\n", res.Mode, res.Pattern, res.Addr, res.LocalAddr, res.RemoteAddr, res.Size, res.Error)
		return
	}
	fmt.Printf("%s %s addr=%s local=%s remote=%s size=%d elapsed=%s bytes=%d bytes/s=%.0f msgs/s=%.0f\n",
		res.Mode, res.Pattern, res.Addr, res.LocalAddr, res.RemoteAddr, res.Size, res.Elapsed, res.Bytes, res.BytesPerSec, res.MsgsPerSec)
	if res.Latency != nil {
		fmt.Printf("latency count=%d min=%s p50=%s p95=%s p99=%s max=%s\n",
			res.Latency.Count, res.Latency.Min, res.Latency.P50, res.Latency.P95, res.Latency.P99, res.Latency.Max)
	}
}

func printRDMABench(res rdmaBenchResult) {
	if res.Error != "" {
		fmt.Printf("%s %s addr=%s device=%s size=%d iters=%d error=%s\n", res.Mode, res.Role, res.Addr, res.Device, res.Size, res.Iterations, res.Error)
		return
	}
	fmt.Printf("%s %s addr=%s device=%s size=%d iters=%d elapsed=%s bytes/s=%.0f msgs/s=%.0f\n",
		res.Mode, res.Role, res.Addr, res.Device, res.Size, res.Iterations, res.Elapsed, res.BytesPerSec, res.MessagesPerSec)
	fmt.Printf("local lid=%d qpn=%d psn=%d gid_index=%d gid=%s global=%v\n",
		res.Local.LID, res.Local.QPN, res.Local.PSN, res.Local.GIDIndex, res.Local.GID, res.Local.UseGlobal)
	fmt.Printf("remote lid=%d qpn=%d psn=%d gid_index=%d gid=%s global=%v\n",
		res.Remote.LID, res.Remote.QPN, res.Remote.PSN, res.Remote.GIDIndex, res.Remote.GID, res.Remote.UseGlobal)
	if res.Latency != nil {
		fmt.Printf("latency count=%d min=%s p50=%s p95=%s p99=%s max=%s\n",
			res.Latency.Count, res.Latency.Min, res.Latency.P50, res.Latency.P95, res.Latency.P99, res.Latency.Max)
	}
}

func printRDMA(sum *rdmaSummary) {
	fmt.Println("rdma available:", sum.Available)
	for _, dev := range sum.Devices {
		if dev.NetInterface != "" {
			fmt.Printf("device %d %s %s %s\n", dev.Index, dev.Name, dev.NetInterface, dev.Handle)
		} else {
			fmt.Printf("device %d %s %s\n", dev.Index, dev.Name, dev.Handle)
		}
	}
	for _, step := range sum.Steps {
		status := "ok"
		if !step.OK {
			status = "failed"
		}
		fmt.Printf("%s: %s", step.Name, status)
		if step.Return != 0 {
			fmt.Printf(" rc=%d", step.Return)
		}
		if step.Handle != "" && step.Handle != "0x0" {
			fmt.Printf(" handle=%s", step.Handle)
		}
		if step.Error != "" {
			fmt.Printf(" error=%s", step.Error)
		}
		fmt.Println()
		if len(step.Fields) > 0 {
			for _, name := range sortedFieldNames(step.Fields) {
				fmt.Printf("  %s=%v\n", name, step.Fields[name])
			}
		}
	}
	for _, note := range sum.Notes {
		fmt.Println("note:", note)
	}
}

func rdmaBuffer(n int) ([]byte, []byte, error) {
	if n <= 0 {
		return nil, nil, fmt.Errorf("size must be positive")
	}
	page := os.Getpagesize()
	size := roundUp(n, page)
	buf, err := syscall.Mmap(-1, 0, size, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		return nil, nil, fmt.Errorf("allocate mmap buffer: %w", err)
	}
	return buf[:n], buf, nil
}

func roundUp(n, unit int) int {
	if unit <= 0 {
		return n
	}
	if rem := n % unit; rem != 0 {
		n += unit - rem
	}
	return n
}

func isZeroGID(gid rdma.IbvGID) bool {
	return xrdma.IsZeroGID(gid)
}

func isIPv4MappedGID(gid rdma.IbvGID) bool {
	return xrdma.IsIPv4MappedGID(gid)
}

func parseGID(text string) (rdma.IbvGID, error) {
	var gid rdma.IbvGID
	b, err := hex.DecodeString(text)
	if err != nil {
		return gid, fmt.Errorf("parse gid: %w", err)
	}
	if len(b) != len(gid) {
		return gid, fmt.Errorf("parse gid: got %d bytes, want %d", len(b), len(gid))
	}
	copy(gid[:], b)
	return gid, nil
}

func boolByte(v bool) uint8 {
	if v {
		return 1
	}
	return 0
}

func runtimeYield() {
	runtime.Gosched()
}

func writeJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fatalf("json: %v", err)
	}
}

func portAttrFields(buf []byte) map[string]any {
	if len(buf) < int(unsafe.Sizeof(ibvPortAttr{})) {
		return nil
	}
	attr := (*ibvPortAttr)(unsafe.Pointer(unsafe.SliceData(buf)))
	return map[string]any{
		"state":            attr.State,
		"state_name":       portStateName(attr.State),
		"active_mtu":       attr.ActiveMTU,
		"active_mtu_bytes": mtuBytes(attr.ActiveMTU),
		"gid_tbl_len":      attr.GIDTblLen,
		"lid":              attr.LID,
		"link_layer":       attr.LinkLayer,
		"link_layer_name":  linkLayerName(attr.LinkLayer),
		"phys_state":       attr.PhysState,
	}
}

func sortedFieldNames(fields map[string]any) []string {
	names := make([]string, 0, len(fields))
	for name := range fields {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func errOrCode(err error, rc int) string {
	if err != nil {
		return err.Error()
	}
	if rc != 0 {
		if rc == 16 {
			return xrdma.ErrnoText(rc)
		}
		return fmt.Sprintf("errno %d", rc)
	}
	return ""
}

func nilReturnError[T ~uintptr](err error, v T, name string) string {
	if err != nil {
		return err.Error()
	}
	if v == 0 {
		return "returned nil " + name
	}
	return ""
}

func isClosedConnError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, net.ErrClosed) || errors.Is(err, io.EOF) {
		return true
	}
	msg := err.Error()
	return strings.Contains(msg, "use of closed network connection") ||
		strings.Contains(msg, "connection reset by peer") ||
		strings.Contains(msg, "broken pipe")
}

func rdmaNetInterface(name string) string {
	return strings.TrimPrefix(name, "rdma_")
}

func portStateName(state int32) string {
	switch state {
	case 1:
		return "PORT_DOWN"
	case 2:
		return "PORT_INIT"
	case 3:
		return "PORT_ARMED"
	case 4:
		return "PORT_ACTIVE"
	case 5:
		return "PORT_ACTIVE_DEFER"
	default:
		return fmt.Sprintf("PORT_%d", state)
	}
}

func linkLayerName(layer uint8) string {
	switch layer {
	case 100:
		return "Thunderbolt"
	case 1:
		return "InfiniBand"
	case 2:
		return "Ethernet"
	default:
		return fmt.Sprintf("link_layer_%d", layer)
	}
}

func mtuBytes(mtu int32) int {
	switch mtu {
	case 1:
		return 256
	case 2:
		return 512
	case 3:
		return 1024
	case 4:
		return 2048
	case 5:
		return 4096
	default:
		return 0
	}
}

func hexHandle[T ~uintptr](v T) string {
	return fmt.Sprintf("0x%x", uintptr(v))
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "rdmaperf: "+format+"\n", args...)
	os.Exit(1)
}
