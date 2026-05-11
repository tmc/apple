// Command rdmaperf measures TCP and RDMA readiness paths.
package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/tmc/apple/rdma"
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
`)
}

func serve(args []string) {
	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	listenAddr := fs.String("listen", ":9000", "listen address")
	jsonOut := fs.Bool("json", false, "print JSON connection summaries")
	fs.Parse(args)

	ln, err := net.Listen("tcp", *listenAddr)
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
	var hdr [headerSize]byte
	if _, err := io.ReadFull(c, hdr[:]); err != nil {
		return
	}
	size := int(binary.BigEndian.Uint32(hdr[:4]))
	pattern := patternName(binary.BigEndian.Uint32(hdr[4:]))
	buf := make([]byte, size)
	res := result{
		Mode:       "serve",
		Pattern:    pattern,
		LocalAddr:  c.LocalAddr().String(),
		RemoteAddr: c.RemoteAddr().String(),
		Size:       size,
	}
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
	c, err := net.Dial("tcp", addr)
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
	fs.Parse(args)

	res := result{Mode: "rdma-probe", RDMA: probeRDMA()}
	if *jsonOut {
		writeJSON(res)
	} else {
		printResult(res)
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
		ctx, err := rdma.Ibv_open_device(dev.Handle)
		sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_open_device", i), OK: err == nil && ctx != 0, Handle: hexHandle(ctx), Error: errString(err)})
		if ctx == 0 {
			continue
		}
		buf := make([]byte, 4096)
		rc, err := rdma.Ibv_query_device(ctx, uintptr(unsafe.Pointer(unsafe.SliceData(buf))))
		sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_query_device", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)})
		portBuf := make([]byte, unsafe.Sizeof(ibvPortAttr{}))
		rc, err = rdma.Ibv_query_port(ctx, 1, uintptr(unsafe.Pointer(unsafe.SliceData(portBuf))))
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
			rc, err = rdma.Ibv_close_device(ctx)
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_close_device", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)})
			continue
		}
		pd, err := rdma.Ibv_alloc_pd(ctx)
		sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_alloc_pd", i), OK: err == nil && pd != 0, Handle: hexHandle(pd), Error: nilReturnError(err, pd, "protection domain")})
		cq, err := rdma.Ibv_create_cq(ctx, 16, 0, 0, 0)
		sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_create_cq", i), OK: err == nil && cq != 0, Handle: hexHandle(cq), Error: nilReturnError(err, cq, "completion queue")})
		if cq != 0 {
			rc, err := rdma.Ibv_destroy_cq(cq)
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_destroy_cq", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)})
		}
		if pd != 0 {
			rc, err := rdma.Ibv_dealloc_pd(pd)
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_dealloc_pd", i), OK: err == nil && rc == 0, Return: rc, Error: errOrCode(err, rc)})
		} else {
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_reg_mr", i), Error: "skipped because ibv_alloc_pd failed"})
			sum.Steps = append(sum.Steps, rdmaStep{Name: fmt.Sprintf("device[%d].ibv_create_qp", i), Error: "skipped because ibv_alloc_pd failed"})
		}
		rc, err = rdma.Ibv_close_device(ctx)
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
