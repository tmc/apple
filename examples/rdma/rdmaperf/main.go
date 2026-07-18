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
	"runtime/debug"
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
	Mode           string                `json:"mode"`
	Pattern        string                `json:"pattern,omitempty"`
	Addr           string                `json:"addr,omitempty"`
	LocalAddr      string                `json:"local_addr,omitempty"`
	RemoteAddr     string                `json:"remote_addr,omitempty"`
	Duration       string                `json:"duration,omitempty"`
	Elapsed        string                `json:"elapsed,omitempty"`
	Size           int                   `json:"size,omitempty"`
	Bytes          uint64                `json:"bytes,omitempty"`
	Messages       uint64                `json:"messages,omitempty"`
	BytesPerSec    float64               `json:"bytes_per_sec,omitempty"`
	MsgsPerSec     float64               `json:"msgs_per_sec,omitempty"`
	Latency        *latencySummary       `json:"latency,omitempty"`
	RDMA           *rdmaSummary          `json:"rdma,omitempty"`
	RCCapability   *rcCapabilityResult   `json:"rc_capability,omitempty"`
	RKeyCapability *rkeyCapabilityResult `json:"rkey_capability,omitempty"`
	PDLifecycle    *pdLifecycleResult    `json:"pd_lifecycle,omitempty"`
	Error          string                `json:"error,omitempty"`
}

type rdmaBenchResult struct {
	Mode           string             `json:"mode"`
	Role           string             `json:"role"`
	Addr           string             `json:"addr,omitempty"`
	Commit         string             `json:"commit,omitempty"`
	Host           string             `json:"host,omitempty"`
	Command        string             `json:"command,omitempty"`
	Device         string             `json:"device,omitempty"`
	DevicePair     string             `json:"device_pair,omitempty"`
	Stage          string             `json:"stage,omitempty"`
	SetupTimeout   string             `json:"setup_timeout,omitempty"`
	GateEnv        string             `json:"gate_env,omitempty"`
	FailureClass   string             `json:"failure_class,omitempty"`
	FirstError     string             `json:"first_error,omitempty"`
	NoRetry        bool               `json:"no_retry"`
	DatapathClaim  bool               `json:"datapath_claim"`
	Control        []rdmaControlEvent `json:"control,omitempty"`
	Size           int                `json:"size"`
	Iterations     int                `json:"iterations"`
	MRCount        int                `json:"mr_count,omitempty"`
	PDsPerRound    int                `json:"pds_per_round,omitempty"`
	QPsPerRound    int                `json:"qps_per_round,omitempty"`
	RoundsDone     int                `json:"rounds_done,omitempty"`
	Outcome        string             `json:"outcome,omitempty"`
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

type rcCapabilityResult struct {
	Device       string `json:"device,omitempty"`
	Outcome      string `json:"outcome"`
	Attempts     int    `json:"attempts"`
	NoRTR        bool   `json:"no_rtr"`
	NoData       bool   `json:"no_data"`
	CreateErrno  int    `json:"create_errno,omitempty"`
	CreateError  string `json:"create_error,omitempty"`
	DestroyError string `json:"destroy_error,omitempty"`
}

type rkeyCapabilityResult struct {
	Device          string `json:"device,omitempty"`
	Outcome         string `json:"outcome"`
	Attempts        int    `json:"attempts"`
	NoQP            bool   `json:"no_qp"`
	NoRTR           bool   `json:"no_rtr"`
	NoData          bool   `json:"no_data"`
	Addr            string `json:"addr,omitempty"`
	LKey            string `json:"lkey,omitempty"`
	RKey            string `json:"rkey,omitempty"`
	RegisterErrno   int    `json:"register_errno,omitempty"`
	RegisterError   string `json:"register_error,omitempty"`
	DeregisterError string `json:"deregister_error,omitempty"`
}

type pdLifecycleResult struct {
	Device        string `json:"device,omitempty"`
	Mode          string `json:"mode"`
	Outcome       string `json:"outcome"`
	Cycles        int    `json:"cycles"`
	AllocPerCycle int    `json:"alloc_per_cycle"`
	MaxAlloc      int    `json:"max_alloc,omitempty"`
	RoundsDone    int    `json:"rounds_done"`
	Allocations   int    `json:"allocations"`
	Deallocations int    `json:"deallocations"`
	NoMR          bool   `json:"no_mr"`
	NoQP          bool   `json:"no_qp"`
	NoRTR         bool   `json:"no_rtr"`
	NoData        bool   `json:"no_data"`
	AllocateErrno int    `json:"allocate_errno,omitempty"`
	AllocateError string `json:"allocate_error,omitempty"`
	DeallocErrno  int    `json:"dealloc_errno,omitempty"`
	DeallocError  string `json:"dealloc_error,omitempty"`
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
	case "rdma-rc-capability":
		rdmaRCCapability(os.Args[2:])
	case "rdma-rkey-capability":
		rdmaRKeyCapability(os.Args[2:])
	case "rdma-pd-lifecycle":
		rdmaPDLifecycle(os.Args[2:])
	case "rdma-pingpong":
		rdmaPingpong(os.Args[2:])
	case "rdma-lifecycle-probe":
		rdmaLifecycleProbe(os.Args[2:])
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
  rdma-rc-capability
              One guarded RC-QP create/destroy capability probe; no RTR or data.
  rdma-rkey-capability
              One guarded MR registration capability probe; no QP, RTR, or data.
  rdma-pd-lifecycle
              Guarded PD alloc/dealloc/realloc lifecycle probe; no MR, QP, RTR, or data.
  rdma-pingpong
              Run RDMA SEND/RECV ping-pong using TCP only for setup exchange.
  rdma-lifecycle-probe
              Two-rank, no-data QP setup/teardown reclamation probe.
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

const rcCapabilityConfirmEnv = "CONFIRM_RDMA_RC_CAPABILITY"
const rcCapabilityConfirmValue = "one-shot-qp-create"

const rkeyCapabilityConfirmEnv = "CONFIRM_RDMA_RKEY_CAPABILITY"
const rkeyCapabilityConfirmValue = "one-shot-mr-register"
const rkeyCapabilityBytes = 4096

const pdLifecycleConfirmEnv = "CONFIRM_RDMA_PD_LIFECYCLE"
const pdLifecycleConfirmValue = "one-shot-pd-lifecycle"
const (
	maxPDLifecycleCycles   = 32
	maxPDAllocPerCycle     = 32
	maxPDExhaustionAlloc   = 64
	defaultPDAllocPerCycle = 11
)

func rdmaRCCapability(args []string) {
	fs := flag.NewFlagSet("rdma-rc-capability", flag.ExitOnError)
	deviceName := fs.String("name", "", "select first RDMA device whose name contains substring")
	deviceIndex := fs.Int("device", -1, "select RDMA device index")
	timeout := fs.Duration("timeout", 10*time.Second, "watchdog limit for the one provider attempt")
	allow := fs.Bool("allow-rc-probe", false, "acknowledge one RC queue-pair create/destroy attempt")
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)

	if err := validateRCCapabilityTimeout(*timeout); err != nil {
		fatalf("%v", err)
	}
	if err := requireRCCapabilityProbeAllowed(*allow); err != nil {
		fatalf("%v", err)
	}
	stopWatchdog := startRDMAWatchdog("rdma RC capability probe", *timeout)
	res := result{Mode: "rdma-rc-capability", RCCapability: probeRCCapability(*deviceName, *deviceIndex)}
	stopWatchdog()
	if res.RCCapability.Outcome == "inconclusive" {
		res.Error = res.RCCapability.CreateError
	}
	if res.RCCapability.DestroyError != "" {
		res.Error = res.RCCapability.DestroyError
	}
	if *jsonOut {
		writeJSON(res)
	} else {
		printResult(res)
	}
	if res.Error != "" {
		os.Exit(1)
	}
}

func requireRCCapabilityProbeAllowed(allow bool) error {
	if !allow {
		return fmt.Errorf("refusing RC capability probe: pass -allow-rc-probe for one create/destroy attempt")
	}
	if os.Getenv(rcCapabilityConfirmEnv) != rcCapabilityConfirmValue {
		return fmt.Errorf("refusing RC capability probe: set %s=%s", rcCapabilityConfirmEnv, rcCapabilityConfirmValue)
	}
	return nil
}

func validateRCCapabilityTimeout(timeout time.Duration) error {
	if timeout <= 0 {
		return fmt.Errorf("-timeout must be positive")
	}
	return nil
}

func rdmaRKeyCapability(args []string) {
	fs := flag.NewFlagSet("rdma-rkey-capability", flag.ExitOnError)
	deviceName := fs.String("name", "", "select first RDMA device whose name contains substring")
	deviceIndex := fs.Int("device", -1, "select RDMA device index")
	timeout := fs.Duration("timeout", 10*time.Second, "watchdog limit for the one provider attempt")
	allow := fs.Bool("allow-rkey-probe", false, "acknowledge one memory-register/deregister attempt")
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)

	if err := validateRCCapabilityTimeout(*timeout); err != nil {
		fatalf("%v", err)
	}
	if err := requireRKeyCapabilityProbeAllowed(*allow); err != nil {
		fatalf("%v", err)
	}
	stopWatchdog := startRDMAWatchdog("rdma rkey capability probe", *timeout)
	res := result{Mode: "rdma-rkey-capability", RKeyCapability: probeRKeyCapability(*deviceName, *deviceIndex)}
	stopWatchdog()
	if res.RKeyCapability.Outcome == "inconclusive" {
		res.Error = res.RKeyCapability.RegisterError
	}
	if res.RKeyCapability.DeregisterError != "" {
		res.Error = res.RKeyCapability.DeregisterError
	}
	if *jsonOut {
		writeJSON(res)
	} else {
		printResult(res)
	}
	if res.Error != "" {
		os.Exit(1)
	}
}

func requireRKeyCapabilityProbeAllowed(allow bool) error {
	if !allow {
		return fmt.Errorf("refusing rkey capability probe: pass -allow-rkey-probe for one memory-register/deregister attempt")
	}
	if os.Getenv(rkeyCapabilityConfirmEnv) != rkeyCapabilityConfirmValue {
		return fmt.Errorf("refusing rkey capability probe: set %s=%s", rkeyCapabilityConfirmEnv, rkeyCapabilityConfirmValue)
	}
	return nil
}

func rdmaPDLifecycle(args []string) {
	fs := flag.NewFlagSet("rdma-pd-lifecycle", flag.ExitOnError)
	deviceName := fs.String("name", "", "select first RDMA device whose name contains substring")
	deviceIndex := fs.Int("device", -1, "select RDMA device index")
	mode := fs.String("mode", "reclaim", "probe mode: reclaim or exhaust")
	cycles := fs.Int("cycles", 1, "allocation/deallocation rounds in reclaim mode (1..32)")
	allocPerCycle := fs.Int("alloc-per-cycle", defaultPDAllocPerCycle, "PD allocations per reclaim round (1..32)")
	maxAlloc := fs.Int("max-alloc", 16, "allocation cap in exhaust mode (1..64)")
	timeout := fs.Duration("timeout", 60*time.Second, "whole-probe watchdog limit")
	opTimeout := fs.Duration("op-timeout", 2*time.Second, "watchdog limit for each PD allocation or deallocation")
	allow := fs.Bool("allow-pd-lifecycle-probe", false, "acknowledge bounded protection-domain lifecycle probe")
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)

	if err := validatePDLifecycleProbe(*mode, *cycles, *allocPerCycle, *maxAlloc, *timeout, *opTimeout); err != nil {
		fatalf("%v", err)
	}
	if err := requirePDLifecycleProbeAllowed(*allow); err != nil {
		fatalf("%v", err)
	}
	stopWatchdog := startRDMAWatchdog("rdma PD lifecycle probe", *timeout)
	res := result{Mode: "rdma-pd-lifecycle", PDLifecycle: probePDLifecycle(*deviceName, *deviceIndex, *mode, *cycles, *allocPerCycle, *maxAlloc, *opTimeout)}
	stopWatchdog()
	if !successfulPDLifecycleOutcome(res.PDLifecycle.Outcome) {
		res.Error = res.PDLifecycle.AllocateError
		if res.PDLifecycle.DeallocError != "" {
			res.Error = res.PDLifecycle.DeallocError
		}
	}
	if *jsonOut {
		writeJSON(res)
	} else {
		printResult(res)
	}
	if res.Error != "" {
		os.Exit(1)
	}
}

func requirePDLifecycleProbeAllowed(allow bool) error {
	if !allow {
		return fmt.Errorf("refusing PD lifecycle probe: pass -allow-pd-lifecycle-probe for bounded alloc/dealloc/realloc")
	}
	if os.Getenv(pdLifecycleConfirmEnv) != pdLifecycleConfirmValue {
		return fmt.Errorf("refusing PD lifecycle probe: set %s=%s", pdLifecycleConfirmEnv, pdLifecycleConfirmValue)
	}
	return nil
}

func validatePDLifecycleProbe(mode string, cycles, allocPerCycle, maxAlloc int, timeout, opTimeout time.Duration) error {
	if mode != "reclaim" && mode != "exhaust" {
		return fmt.Errorf("-mode must be reclaim or exhaust")
	}
	if cycles < 1 || cycles > maxPDLifecycleCycles {
		return fmt.Errorf("-cycles must be in [1,%d]", maxPDLifecycleCycles)
	}
	if allocPerCycle < 1 || allocPerCycle > maxPDAllocPerCycle {
		return fmt.Errorf("-alloc-per-cycle must be in [1,%d]", maxPDAllocPerCycle)
	}
	if maxAlloc < 1 || maxAlloc > maxPDExhaustionAlloc {
		return fmt.Errorf("-max-alloc must be in [1,%d]", maxPDExhaustionAlloc)
	}
	if err := validateRCCapabilityTimeout(timeout); err != nil {
		return err
	}
	if err := validateRCCapabilityTimeout(opTimeout); err != nil {
		return fmt.Errorf("-op-timeout: %w", err)
	}
	if opTimeout > timeout {
		return fmt.Errorf("-op-timeout must not exceed -timeout")
	}
	return nil
}

func startRDMAWatchdog(name string, timeout time.Duration) func() {
	if timeout <= 0 {
		return func() {}
	}
	timer := time.AfterFunc(timeout, func() {
		fmt.Fprintf(os.Stderr, "rdmaperf: %s watchdog expired after %s; provider may be wedged for this boot, stop live RDMA probes and reboot before retrying\n", name, timeout)
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
	zeroDLIDWhenGlobal := fs.Bool("zero-dlid-when-global", false, "set DLID=0 when RTR AH uses a global route")
	grhHopLimit := fs.Int("grh-hop-limit", 0, "GRH hop limit override, 0 uses default")
	grhTrafficClass := fs.Int("grh-traffic-class", 0, "GRH traffic class")
	grhFlowLabel := fs.Uint("grh-flow-label", 0, "GRH flow label")
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
	policy, err := rdmaRTRPolicyFromFlags(*zeroDLIDWhenGlobal, *grhHopLimit, *grhTrafficClass, *grhFlowLabel)
	if err != nil {
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
		res = runRDMAPingpongServer(*listenAddr, *deviceName, *deviceIndex, *gidIndex, policy, size, *iters, *setupTimeout)
	} else {
		res = runRDMAPingpongClient(*addr, *deviceName, *deviceIndex, *gidIndex, policy, size, *iters, *setupTimeout)
	}
	finalizeRDMABenchResult(&res, *setupTimeout, *allowRTR)
	if *jsonOut {
		writeJSON(res)
	} else {
		printRDMABench(res)
	}
	if res.Error != "" {
		os.Exit(1)
	}
}

const lifecycleProbeConfirmEnv = "CONFIRM_RDMA_LIFECYCLE_LEAK"
const lifecycleProbeConfirmValue = "one-shot-lifecycle"

func rdmaLifecycleProbe(args []string) {
	fs := flag.NewFlagSet("rdma-lifecycle-probe", flag.ExitOnError)
	listenAddr := fs.String("listen", "", "listen address for rank 0")
	addr := fs.String("addr", "", "rank 0 address for rank 1")
	deviceName := fs.String("name", "", "select RDMA device")
	deviceIndex := fs.Int("device", -1, "select RDMA device index")
	rounds := fs.Int("rounds", 2, "setup/teardown rounds (1..3)")
	mrs := fs.Int("mrs", 2, "memory regions per round (1..4)")
	timeout := fs.Duration("timeout", 20*time.Second, "whole-probe watchdog limit")
	setupTimeout := fs.Duration("setup-timeout", 5*time.Second, "per-round setup watchdog limit")
	allow := fs.Bool("allow-lifecycle-probe", false, "acknowledge two-rank lifecycle probe")
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)
	if (*listenAddr == "") == (*addr == "") || *rounds < 1 || *rounds > 3 || *mrs < 1 || *mrs > 4 || *timeout <= 0 || *setupTimeout <= 0 {
		fatalf("require exactly one of -listen/-addr, -rounds in [1,3], -mrs in [1,4], and positive timeouts")
	}
	if !*allow || os.Getenv(lifecycleProbeConfirmEnv) != lifecycleProbeConfirmValue {
		fatalf("refusing lifecycle probe: pass -allow-lifecycle-probe and set %s=%s", lifecycleProbeConfirmEnv, lifecycleProbeConfirmValue)
	}
	if err := xrdma.RequireRTRAttemptAllowed(true); err != nil {
		fatalf("lifecycle probe RTR gate: %v", err)
	}
	stop := startRDMAWatchdog("rdma lifecycle probe", *timeout)
	var res rdmaBenchResult
	if *listenAddr != "" {
		res = runLifecycleServer(*listenAddr, *deviceName, *deviceIndex, *rounds, *mrs, *setupTimeout)
	} else {
		res = runLifecycleClient(*addr, *deviceName, *deviceIndex, *rounds, *mrs, *setupTimeout)
	}
	stop()
	res.Mode, res.Iterations, res.DatapathClaim = "rdma-lifecycle-probe", *rounds, false
	res.MRCount, res.PDsPerRound, res.QPsPerRound = *mrs, 1, 1
	if res.Error == "" && res.RoundsDone == *rounds {
		res.Outcome = "reclaimed"
	} else if res.Error != "" {
		res.Outcome = "failed"
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

func runLifecycleServer(listen, name string, index, rounds, mrs int, timeout time.Duration) rdmaBenchResult {
	res := newRDMABenchResult("rank0", listen, 0, rounds)
	ln, err := listenTCP(listen)
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
	for i := 0; i < rounds; i++ {
		if err := runRDMAPingpong(c, false, name, index, -1, xrdma.RTRPolicy{}, 4096, 0, mrs, timeout, &res); err != nil {
			res.Error = fmt.Sprintf("round %d: %v", i+1, err)
			return res
		}
		res.RoundsDone++
	}
	res.Stage = "done"
	return res
}

func runLifecycleClient(addr, name string, index, rounds, mrs int, timeout time.Duration) rdmaBenchResult {
	res := newRDMABenchResult("rank1", addr, 0, rounds)
	c, err := dialTCP(addr)
	if err != nil {
		res.Error = err.Error()
		return res
	}
	defer c.Close()
	for i := 0; i < rounds; i++ {
		if err := runRDMAPingpong(c, true, name, index, -1, xrdma.RTRPolicy{}, 4096, 0, mrs, timeout, &res); err != nil {
			res.Error = fmt.Sprintf("round %d: %v", i+1, err)
			return res
		}
		res.RoundsDone++
	}
	res.Stage = "done"
	return res
}

func rdmaRTRPolicyFromFlags(zeroDLIDWhenGlobal bool, hopLimit, trafficClass int, flowLabel uint) (xrdma.RTRPolicy, error) {
	if hopLimit < 0 || hopLimit > 255 {
		return xrdma.RTRPolicy{}, fmt.Errorf("-grh-hop-limit must be between 0 and 255")
	}
	if trafficClass < 0 || trafficClass > 255 {
		return xrdma.RTRPolicy{}, fmt.Errorf("-grh-traffic-class must be between 0 and 255")
	}
	if flowLabel > 0xfffff {
		return xrdma.RTRPolicy{}, fmt.Errorf("-grh-flow-label must be between 0 and 1048575")
	}
	return xrdma.RTRPolicy{
		ZeroDLIDWhenGlobal: zeroDLIDWhenGlobal,
		HopLimit:           uint8(hopLimit),
		TrafficClass:       uint8(trafficClass),
		FlowLabel:          uint32(flowLabel),
	}, nil
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
	dev       rdma.Device
	ctx       rdma.RDMAContext
	pd        rdma.RDMAPD
	cq        rdma.RDMACQ
	qp        rdma.RDMAQP
	mr        rdma.RDMAMR
	extraMRs  []rdma.RDMAMR
	extraMaps [][]byte
	poller    rdma.IbvCQPoller
	poster    rdma.IbvQPPoster
	mapBuf    []byte
	buf       []byte
	port      ibvPortAttr
	gid       rdma.IbvGID
	gidIndex  int
	psn       uint32
}

type routeGID struct {
	index int
	gid   rdma.IbvGID
}

func runRDMAPingpongServer(listenAddr, deviceName string, deviceIndex, gidIndex int, policy xrdma.RTRPolicy, size, iters int, setupTimeout time.Duration) rdmaBenchResult {
	res := newRDMABenchResult("server", listenAddr, size, iters)
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
	if err := runRDMAPingpong(c, false, deviceName, deviceIndex, gidIndex, policy, size, iters, 1, setupTimeout, &res); err != nil {
		res.Error = err.Error()
	}
	return res
}

func runRDMAPingpongClient(addr, deviceName string, deviceIndex, gidIndex int, policy xrdma.RTRPolicy, size, iters int, setupTimeout time.Duration) rdmaBenchResult {
	res := newRDMABenchResult("client", addr, size, iters)
	c, err := dialTCP(addr)
	if err != nil {
		res.Error = err.Error()
		return res
	}
	defer c.Close()
	if err := runRDMAPingpong(c, true, deviceName, deviceIndex, gidIndex, policy, size, iters, 1, setupTimeout, &res); err != nil {
		res.Error = err.Error()
	}
	return res
}

func newRDMABenchResult(role, addr string, size, iters int) rdmaBenchResult {
	host, _ := os.Hostname()
	return rdmaBenchResult{
		Mode:       "rdma-pingpong",
		Role:       role,
		Addr:       addr,
		Commit:     vcsRevision(),
		Host:       host,
		Command:    strings.Join(os.Args, " "),
		GateEnv:    "allow-rtr=false",
		NoRetry:    true,
		Size:       size,
		Iterations: iters,
	}
}

func finalizeRDMABenchResult(res *rdmaBenchResult, setupTimeout time.Duration, allowRTR bool) {
	res.SetupTimeout = setupTimeout.String()
	res.GateEnv = fmt.Sprintf("allow-rtr=%t", allowRTR)
	res.NoRetry = true
	res.DatapathClaim = res.Stage == "done" && res.Error == ""
	if res.Device != "" && res.Remote.Name != "" {
		res.DevicePair = res.Device + "<->" + res.Remote.Name
	}
	if res.Error != "" {
		res.FirstError = firstLine(res.Error)
		res.FailureClass = classifyRDMABenchFailure(res.Error)
	}
}

func vcsRevision() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			return setting.Value
		}
	}
	return ""
}

func firstLine(s string) string {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		return s[:i]
	}
	return s
}

func classifyRDMABenchFailure(s string) string {
	switch {
	case strings.Contains(s, "rdma setup timeout"):
		return string(rdma.FailureProviderTimeout)
	case strings.Contains(s, "errno 60") || strings.Contains(s, "ETIMEDOUT") || strings.Contains(s, "i/o timeout"):
		return "timeout"
	case strings.Contains(s, "nil provider result") || strings.Contains(s, "provider returned nil"):
		return string(rdma.FailureNilProviderResult)
	case strings.Contains(s, "provider returned negative status"):
		return string(rdma.FailureNegativeProviderReturn)
	case strings.Contains(s, "provider returned status"):
		return string(rdma.FailureProviderStatus)
	case strings.Contains(s, "no RDMA device") || strings.Contains(s, "no rdma device"):
		return string(rdma.FailureNoDevice)
	case strings.Contains(s, "rdma rtr unsafe"):
		return "rtr_refused"
	case strings.Contains(s, "work completion protection error"):
		return "completion_protection"
	case strings.Contains(s, "work completion failure"):
		return "completion_failure"
	default:
		return "error"
	}
}

func runRDMAPingpong(c net.Conn, client bool, deviceName string, deviceIndex, gidIndex int, policy xrdma.RTRPolicy, size, iters, mrCount int, setupTimeout time.Duration, res *rdmaBenchResult) error {
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
	r, err := openRDMAResources(deviceName, deviceIndex, gidIndex, size, setupTimeout)
	if err != nil {
		return err
	}
	defer r.close()
	if err := addLifecycleMRs(r, mrCount-1, setupTimeout); err != nil {
		return err
	}

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
	// IbvModifyQp can wedge inside the provider. Do not run it in a goroutine
	// and return on a timer: that would leave the call operating on resources
	// this function may close. The watchdog terminates the process if the
	// synchronous attempt does not return; it is containment, not cancellation.
	stopWatchdog := startRDMAWatchdog("rdma QP setup", setupTimeout)
	connectErr := r.connect(remote, policy)
	stopWatchdog()
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

func addLifecycleMRs(r *rdmaResources, count int, timeout time.Duration) error {
	for range count {
		buf, mapBuf, err := rdmaBuffer(len(r.buf))
		if err != nil {
			return err
		}
		stop := startRDMAWatchdog("ibv_reg_mr", timeout)
		mr, err := rdma.IbvRegMr(r.pd, uintptr(unsafe.Pointer(unsafe.SliceData(buf))), uintptr(len(buf)), rdma.IBV_ACCESS_LOCAL_WRITE|rdma.IBV_ACCESS_REMOTE_READ|rdma.IBV_ACCESS_REMOTE_WRITE)
		stop()
		runtime.KeepAlive(buf)
		if err != nil || mr == 0 {
			_ = syscall.Munmap(mapBuf)
			return fmt.Errorf("ibv_reg_mr: %s", nilReturnError(err, mr, "memory region"))
		}
		r.extraMRs = append(r.extraMRs, mr)
		r.extraMaps = append(r.extraMaps, mapBuf)
	}
	return nil
}

func probeRCCapability(deviceName string, deviceIndex int) *rcCapabilityResult {
	result := &rcCapabilityResult{Outcome: "inconclusive", NoRTR: true, NoData: true}
	if !rdma.Available() {
		result.CreateError = "rdma unavailable"
		return result
	}
	devs, err := rdmaDevices()
	if err != nil {
		result.CreateError = fmt.Sprintf("rdma devices: %v", err)
		return result
	}
	dev, err := selectRCCapabilityDevice(devs, deviceName, deviceIndex)
	if err != nil {
		result.CreateError = err.Error()
		return result
	}
	result.Device = dev.Name
	ctx, err := dev.Open()
	if err != nil || ctx == 0 {
		result.CreateError = "ibv_open_device: " + nilReturnError(err, ctx, "context")
		return result
	}
	defer rdma.IbvCloseDevice(ctx)
	pd, err := rdma.IbvAllocPd(ctx)
	if err != nil || pd == 0 {
		result.CreateError = "ibv_alloc_pd: " + nilReturnError(err, pd, "protection domain")
		return result
	}
	defer rdma.IbvDeallocPd(pd)
	cq, err := rdma.IbvCreateCq(ctx, 1, 0, 0, 0)
	if err != nil || cq == 0 {
		result.CreateError = "ibv_create_cq: " + nilReturnError(err, cq, "completion queue")
		return result
	}
	defer rdma.IbvDestroyCq(cq)

	result.Attempts = 1
	qp, err := rdma.IbvCreateQpAttr(pd, &rdma.IbvQPInitAttr{
		SendCQ: cq,
		RecvCQ: cq,
		Cap: rdma.IbvQPCap{
			MaxSendWR:  1,
			MaxRecvWR:  1,
			MaxSendSGE: 1,
			MaxRecvSGE: 1,
		},
		QPType: rdma.IBV_QPT_RCExperimental,
	})
	result.Outcome, result.CreateErrno, result.CreateError = classifyRCCapabilityCreate(qp, err)
	if qp != 0 {
		rc, destroyErr := rdma.IbvDestroyQp(qp)
		if destroyErr != nil || rc != 0 {
			result.DestroyError = errOrCode(destroyErr, rc)
		}
	}
	return result
}

func probeRKeyCapability(deviceName string, deviceIndex int) *rkeyCapabilityResult {
	result := &rkeyCapabilityResult{Outcome: "inconclusive", NoQP: true, NoRTR: true, NoData: true}
	if !rdma.Available() {
		result.RegisterError = "rdma unavailable"
		return result
	}
	devs, err := rdmaDevices()
	if err != nil {
		result.RegisterError = fmt.Sprintf("rdma devices: %v", err)
		return result
	}
	dev, err := selectRCCapabilityDevice(devs, deviceName, deviceIndex)
	if err != nil {
		result.RegisterError = err.Error()
		return result
	}
	result.Device = dev.Name
	ctx, err := dev.Open()
	if err != nil || ctx == 0 {
		result.RegisterError = "ibv_open_device: " + nilReturnError(err, ctx, "context")
		return result
	}
	defer rdma.IbvCloseDevice(ctx)
	pd, err := rdma.IbvAllocPd(ctx)
	if err != nil || pd == 0 {
		result.RegisterError = "ibv_alloc_pd: " + nilReturnError(err, pd, "protection domain")
		return result
	}
	defer rdma.IbvDeallocPd(pd)
	buf, mapBuf, err := rdmaBuffer(rkeyCapabilityBytes)
	if err != nil {
		result.RegisterError = err.Error()
		return result
	}
	defer syscall.Munmap(mapBuf)

	result.Attempts = 1
	mr, err := rdma.IbvRegMr(pd, uintptr(unsafe.Pointer(unsafe.SliceData(buf))), uintptr(len(buf)), rdma.IBV_ACCESS_LOCAL_WRITE|rdma.IBV_ACCESS_REMOTE_READ|rdma.IBV_ACCESS_REMOTE_WRITE)
	runtime.KeepAlive(buf)
	result.Outcome, result.RegisterErrno, result.RegisterError = classifyRKeyCapabilityRegistration(mr, err)
	if mr != 0 {
		result.Addr = fmt.Sprintf("0x%x", uintptr(unsafe.Pointer(unsafe.SliceData(buf))))
		result.LKey = fmt.Sprintf("0x%x", rdma.Ibv_mr_lkey(mr))
		result.RKey = fmt.Sprintf("0x%x", rdma.Ibv_mr_rkey(mr))
		if rdma.Ibv_mr_rkey(mr) == 0 {
			result.Outcome = "zero"
		} else {
			result.Outcome = "nonzero"
		}
		rc, deregisterErr := rdma.IbvDeregMr(mr)
		if deregisterErr != nil || rc != 0 {
			result.DeregisterError = errOrCode(deregisterErr, rc)
		}
	}
	return result
}

// probePDLifecycle probes either PD exhaustion or same-process reclamation. It
// stops at the first unexpected provider result and never retries that call.
func probePDLifecycle(deviceName string, deviceIndex int, mode string, cycles, allocPerCycle, maxAlloc int, opTimeout time.Duration) *pdLifecycleResult {
	result := &pdLifecycleResult{Mode: mode, Outcome: "inconclusive", Cycles: cycles, AllocPerCycle: allocPerCycle, NoMR: true, NoQP: true, NoRTR: true, NoData: true}
	if !rdma.Available() {
		result.AllocateError = "rdma unavailable"
		return result
	}
	devs, err := rdmaDevices()
	if err != nil {
		result.AllocateError = fmt.Sprintf("rdma devices: %v", err)
		return result
	}
	dev, err := selectRCCapabilityDevice(devs, deviceName, deviceIndex)
	if err != nil {
		result.AllocateError = err.Error()
		return result
	}
	result.Device = dev.Name
	ctx, err := dev.Open()
	if err != nil || ctx == 0 {
		result.AllocateError = "ibv_open_device: " + nilReturnError(err, ctx, "context")
		return result
	}
	defer rdma.IbvCloseDevice(ctx)

	if mode == "exhaust" {
		result.MaxAlloc = maxAlloc
		return probePDExhaustion(ctx, result, maxAlloc, opTimeout)
	}
	for cycle := 0; cycle < cycles; cycle++ {
		pds, ok := allocatePDs(ctx, result, allocPerCycle, opTimeout)
		if !ok {
			if cycle > 0 {
				result.Outcome = "reclamation_failed"
			} else {
				result.Outcome = "allocation_failed"
			}
			return result
		}
		if !deallocatePDs(pds, result, opTimeout) {
			result.Outcome = "deallocation_failed"
			return result
		}
		result.RoundsDone++
	}
	result.Outcome = "reclaimed"
	return result
}

func probePDExhaustion(ctx rdma.RDMAContext, result *pdLifecycleResult, maxAlloc int, opTimeout time.Duration) *pdLifecycleResult {
	pds, ok := allocatePDs(ctx, result, maxAlloc, opTimeout)
	if ok {
		result.Outcome = "limit_not_reached"
	} else {
		result.Outcome = "exhausted"
	}
	if !deallocatePDs(pds, result, opTimeout) {
		result.Outcome = "deallocation_failed"
	}
	return result
}

func allocatePDs(ctx rdma.RDMAContext, result *pdLifecycleResult, n int, opTimeout time.Duration) ([]rdma.RDMAPD, bool) {
	pds := make([]rdma.RDMAPD, 0, n)
	for range n {
		stop := startRDMAWatchdog("ibv_alloc_pd", opTimeout)
		pd, err := rdma.IbvAllocPd(ctx)
		stop()
		if err != nil || pd == 0 {
			result.AllocateErrno, result.AllocateError = providerErrno(err), "ibv_alloc_pd: "+nilReturnError(err, pd, "protection domain")
			return pds, false
		}
		result.Allocations++
		pds = append(pds, pd)
	}
	return pds, true
}

func deallocatePDs(pds []rdma.RDMAPD, result *pdLifecycleResult, opTimeout time.Duration) bool {
	for _, pd := range pds {
		stop := startRDMAWatchdog("ibv_dealloc_pd", opTimeout)
		rc, err := rdma.IbvDeallocPd(pd)
		stop()
		if err != nil || rc != 0 {
			result.DeallocErrno, result.DeallocError = providerErrno(err), "ibv_dealloc_pd: "+errOrCode(err, rc)
			return false
		}
		result.Deallocations++
	}
	return true
}

func successfulPDLifecycleOutcome(outcome string) bool {
	return outcome == "reclaimed" || outcome == "exhausted"
}

func providerErrno(err error) int {
	var providerErr *rdma.ProviderError
	if errors.As(err, &providerErr) && providerErr.ErrnoSet {
		return providerErr.Errno
	}
	return 0
}

func classifyRKeyCapabilityRegistration(mr rdma.RDMAMR, err error) (outcome string, errno int, detail string) {
	if mr != 0 {
		return "registered", 0, ""
	}
	if err == nil {
		return "inconclusive", 0, "ibv_reg_mr returned nil memory region"
	}
	var providerErr *rdma.ProviderError
	if errors.As(err, &providerErr) && providerErr.ErrnoSet {
		errno = providerErr.Errno
	}
	return "inconclusive", errno, err.Error()
}

func selectRCCapabilityDevice(devs []rdma.Device, name string, index int) (rdma.Device, error) {
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
	if len(devs) == 0 {
		return rdma.Device{}, fmt.Errorf("no rdma devices")
	}
	return devs[0], nil
}

func classifyRCCapabilityCreate(qp rdma.RDMAQP, err error) (outcome string, errno int, detail string) {
	if qp != 0 {
		return "supported", 0, ""
	}
	if err == nil {
		return "inconclusive", 0, "ibv_create_qp returned nil queue pair"
	}
	var providerErr *rdma.ProviderError
	if errors.As(err, &providerErr) && providerErr.ErrnoSet {
		errno = providerErr.Errno
		if rdma.IsUnsupportedErrno(errno) {
			return "rejected", errno, err.Error()
		}
	}
	return "inconclusive", errno, err.Error()
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
		if r.port.LinkLayer == xrdma.LinkLayerThunderbolt {
			return fmt.Errorf("auto-select route gid: no automatic-safe Thunderbolt route gid found; index 0 requires explicit -gid-index 0")
		}
		return fmt.Errorf("auto-select route gid: no nonzero gid found")
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

func (r *rdmaResources) connect(remote rdmaPeerInfo, policy xrdma.RTRPolicy) error {
	init := rdma.IbvQPAttr{
		QPState:       rdma.IBV_QPS_INIT,
		PKeyIndex:     0,
		PortNum:       1,
		QPAccessFlags: rdma.IBV_ACCESS_LOCAL_WRITE | rdma.IBV_ACCESS_REMOTE_READ | rdma.IBV_ACCESS_REMOTE_WRITE,
	}
	initMask := rdma.IBV_QP_STATE | rdma.IBV_QP_PKEY_INDEX | rdma.IBV_QP_PORT | rdma.IBV_QP_ACCESS_FLAGS
	rc, err := rdma.IbvModifyQpAttr(r.qp, &init, initMask)
	if err != nil || rc != 0 {
		return rdma.NewModifyQPError(r.qp, &init, initMask, rc, err)
	}

	rtr, rtrMask, err := r.rtrAttr(remote, policy)
	if err != nil {
		return err
	}
	rc, err = rdma.IbvModifyQpAttr(r.qp, &rtr, rtrMask)
	if err != nil || rc != 0 {
		return rdma.NewModifyQPError(r.qp, &rtr, rtrMask, rc, err)
	}

	rts := rdma.IbvQPAttr{
		QPState: rdma.IBV_QPS_RTS,
		SQPSN:   r.psn,
	}
	rtsMask := rdma.IBV_QP_STATE | rdma.IBV_QP_SQ_PSN
	rc, err = rdma.IbvModifyQpAttr(r.qp, &rts, rtsMask)
	if err != nil || rc != 0 {
		return rdma.NewModifyQPError(r.qp, &rts, rtsMask, rc, err)
	}
	return nil
}

func (r *rdmaResources) rtrAttr(remote rdmaPeerInfo, policy xrdma.RTRPolicy) (rdma.IbvQPAttr, int, error) {
	var gid rdma.IbvGID
	if remote.UseGlobal {
		var err error
		gid, err = parseGID(remote.GID)
		if err != nil {
			return rdma.IbvQPAttr{}, 0, err
		}
	}
	return xrdma.RTRAttr(xrdma.LocalQP{
		PortNum:   1,
		GIDIndex:  r.gidIndex,
		ActiveMTU: r.port.ActiveMTU,
		LinkLayer: r.port.LinkLayer,
	}, xrdma.RemoteQP{
		Name:      remote.Name,
		LID:       remote.LID,
		QPN:       remote.QPN,
		PSN:       remote.PSN,
		GIDIndex:  remote.GIDIndex,
		GID:       gid,
		UseGlobal: remote.UseGlobal,
		ActiveMTU: remote.ActiveMTU,
	}, policy)
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
			return fmt.Errorf("work completion %s status=%d opcode=%d vendor_err=%d", rdma.ClassifyCompletionStatus(wc.Status), wc.Status, wc.Opcode, wc.VendorErr)
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
	for _, mr := range r.extraMRs {
		_, _ = rdma.IbvDeregMr(mr)
	}
	for _, mapBuf := range r.extraMaps {
		_ = syscall.Munmap(mapBuf)
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
	if res.RCCapability != nil {
		cap := res.RCCapability
		fmt.Printf("%s device=%s outcome=%s attempts=%d no_rtr=%v no_data=%v", res.Mode, cap.Device, cap.Outcome, cap.Attempts, cap.NoRTR, cap.NoData)
		if cap.CreateErrno != 0 {
			fmt.Printf(" create_errno=%s", rdma.ErrnoText(cap.CreateErrno))
		}
		if cap.CreateError != "" {
			fmt.Printf(" create_error=%s", cap.CreateError)
		}
		if cap.DestroyError != "" {
			fmt.Printf(" destroy_error=%s", cap.DestroyError)
		}
		fmt.Println()
		return
	}
	if res.RKeyCapability != nil {
		cap := res.RKeyCapability
		fmt.Printf("%s device=%s outcome=%s attempts=%d no_qp=%v no_rtr=%v no_data=%v", res.Mode, cap.Device, cap.Outcome, cap.Attempts, cap.NoQP, cap.NoRTR, cap.NoData)
		for _, field := range []struct {
			name  string
			value string
		}{{"addr", cap.Addr}, {"lkey", cap.LKey}, {"rkey", cap.RKey}, {"register_error", cap.RegisterError}, {"deregister_error", cap.DeregisterError}} {
			if field.value != "" {
				fmt.Printf(" %s=%s", field.name, field.value)
			}
		}
		if cap.RegisterErrno != 0 {
			fmt.Printf(" register_errno=%s", rdma.ErrnoText(cap.RegisterErrno))
		}
		fmt.Println()
		return
	}
	if res.PDLifecycle != nil {
		cap := res.PDLifecycle
		fmt.Printf("%s device=%s mode=%s outcome=%s cycles=%d alloc_per_cycle=%d max_alloc=%d rounds_done=%d allocations=%d deallocations=%d no_mr=%v no_qp=%v no_rtr=%v no_data=%v", res.Mode, cap.Device, cap.Mode, cap.Outcome, cap.Cycles, cap.AllocPerCycle, cap.MaxAlloc, cap.RoundsDone, cap.Allocations, cap.Deallocations, cap.NoMR, cap.NoQP, cap.NoRTR, cap.NoData)
		if cap.AllocateErrno != 0 {
			fmt.Printf(" allocate_errno=%s", rdma.ErrnoText(cap.AllocateErrno))
		}
		if cap.AllocateError != "" {
			fmt.Printf(" allocate_error=%s", cap.AllocateError)
		}
		if cap.DeallocErrno != 0 {
			fmt.Printf(" dealloc_errno=%s", rdma.ErrnoText(cap.DeallocErrno))
		}
		if cap.DeallocError != "" {
			fmt.Printf(" dealloc_error=%s", cap.DeallocError)
		}
		fmt.Println()
		return
	}
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
		return rdma.ErrnoText(rc)
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
	for _, a := range args {
		if err, ok := a.(error); ok {
			if hint := xrdma.ResourceExhaustionHint(err); hint != "" {
				fmt.Fprintf(os.Stderr, "rdmaperf: hint: %s\n", hint)
			}
		}
	}
	os.Exit(1)
}
