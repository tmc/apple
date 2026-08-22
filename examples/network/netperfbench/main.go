// Command netperfbench measures TCP echo latency and throughput over three
// transports so they can be compared on identical work:
//
//	nw     Network.framework through the Go bindings in package network
//	std    the Go standard library net package
//	swift  Network.framework from Swift (see swift/netperfbench.swift)
//
// Every implementation speaks the same wire protocol — the client sends a
// fixed-size payload and reads the same number of bytes back — so a client
// of one kind can be pointed at a server of another. Pairing a Go client
// with a Swift server (and the reverse) separates the cost of the bindings
// from the cost of Network.framework itself.
//
// Run both roles at once:
//
//	netperfbench -role both -impl nw -size 4096 -n 20000
//
// Or split them, on one host or two:
//
//	netperfbench -role server -impl nw -port 51000
//	netperfbench -role client -impl std -addr 127.0.0.1:51000
//
// With -json the client writes one JSON object of results to stdout, which
// is what run.sh consumes to build its comparison table. The object carries
// the configuration, the resolved network path, and the machine it was
// measured on, so a saved result stays interpretable without the command
// line that produced it.
//
// A measurement that silently ran over the wrong path is worse than no
// measurement, so -require-interface and -forbid-loopback turn the path into
// an assertion:
//
//	netperfbench -role client -addr 10.0.0.2:51000 -require-interface en0
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"slices"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	role      = flag.String("role", "both", "server, client, or both")
	impl      = flag.String("impl", "nw", "transport: nw (Network.framework) or std (net package)")
	addr      = flag.String("addr", "127.0.0.1:51000", "server address for -role client")
	port      = flag.String("port", "51000", "listen port for -role server")
	size      = flag.Int("size", 4096, "payload bytes per round trip")
	count     = flag.Int("n", 10000, "round trips to measure")
	warmup    = flag.Int("warmup", 200, "unmeasured round trips before measuring")
	inflight  = flag.Int("inflight", 1, "messages sent before waiting for any echo")
	repeat    = flag.Int("repeat", 1, "measurement repetitions; the median run is reported")
	label     = flag.String("label", "", "name for this run in the report (default: impl)")
	asJSON    = flag.Bool("json", false, "write results as JSON")
	recvBatch = flag.Bool("recv-batch", false, "receive one whole in-flight batch (experimental)")

	requireInterface = flag.String("require-interface", "", "fail unless the connection used this interface (nw only)")
	forbidLoopback   = flag.Bool("forbid-loopback", false, "fail if the connection used loopback (nw only)")

	cpuProfile = flag.String("cpuprofile", "", "write a CPU profile of the measured run to this file")
	memProfile = flag.String("memprofile", "", "write a heap profile after the measured run to this file")
)

// echoServer accepts connections and echoes every byte back.
type echoServer interface {
	// Port returns the port the server bound to, which matters when the
	// caller asked for port 0.
	Port() string
	Close()
}

// echoClient performs one round trip at a time on a single connection.
type echoClient interface {
	// RoundTrip sends n copies of buf and reads all n echoes back.
	RoundTrip(buf []byte, n int) error
	Close()
}

// pathReporter is implemented by clients that can say which network path
// they actually got. A benchmark that meant to measure a physical link and
// silently ran over loopback is worse than no measurement, so the answer is
// recorded in the result and can be asserted with -require-interface.
type pathReporter interface {
	Path() *pathInfo
}

// pathInfo describes the resolved path of a connection.
type pathInfo struct {
	Status       string   `json:"status"`
	UsesLoopback bool     `json:"uses_loopback"`
	UsesWifi     bool     `json:"uses_wifi"`
	UsesWired    bool     `json:"uses_wired"`
	Interfaces   []string `json:"interfaces"`
}

func (p *pathInfo) validate() error {
	if *requireInterface == "" && !*forbidLoopback {
		return nil
	}
	if p == nil {
		return fmt.Errorf("path assertion requested but the %s transport reports no path", *impl)
	}
	if *forbidLoopback && p.UsesLoopback {
		return fmt.Errorf("connection used loopback (interfaces %v)", p.Interfaces)
	}
	if *requireInterface != "" && !slices.Contains(p.Interfaces, *requireInterface) {
		return fmt.Errorf("connection did not use interface %q (used %v)", *requireInterface, p.Interfaces)
	}
	return nil
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("netperfbench: ")
	flag.Parse()
	enableSignposts()

	if *size <= 0 {
		log.Fatal("-size must be positive")
	}
	if *impl != "nw" && *impl != "std" {
		log.Fatalf("unknown -impl %q", *impl)
	}

	switch *role {
	case "server":
		srv, err := serve(*impl, *port)
		if err != nil {
			log.Fatal(err)
		}
		defer srv.Close()
		fmt.Fprintf(os.Stderr, "listening on port %s (%s)\n", srv.Port(), *impl)
		waitForSignal()
	case "client":
		if err := runClient(*addr); err != nil {
			log.Fatal(err)
		}
	case "both":
		srv, err := serve(*impl, "0")
		if err != nil {
			log.Fatal(err)
		}
		defer srv.Close()
		if err := runClient("127.0.0.1:" + srv.Port()); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown -role %q", *role)
	}
}

func serve(impl, port string) (echoServer, error) {
	if impl == "std" {
		return serveStd(port)
	}
	return serveNW(port)
}

func dial(impl, addr string, inflight int) (echoClient, error) {
	if impl == "std" {
		return dialStd(addr)
	}
	return dialNW(addr, inflight)
}

func runClient(addr string) error {
	c, err := dial(*impl, addr, max(*inflight, 1))
	if err != nil {
		return err
	}
	defer c.Close()

	// A non-constant payload keeps any compression or page-dedup shortcut
	// out of the measurement.
	buf := make([]byte, *size)
	for i := range buf {
		buf[i] = byte(i*7 + 13)
	}

	for range *warmup {
		if err := c.RoundTrip(buf, *inflight); err != nil {
			return fmt.Errorf("warmup: %w", err)
		}
	}

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			return err
		}
		defer pprof.StopCPUProfile()
	}

	runs := make([]run, 0, *repeat)
	for range max(*repeat, 1) {
		latencies := make([]time.Duration, 0, *count)
		cpu0 := readCPUTime()
		start := time.Now()
		for range *count {
			t0 := time.Now()
			if err := c.RoundTrip(buf, *inflight); err != nil {
				return err
			}
			latencies = append(latencies, time.Since(t0))
		}
		elapsed := time.Since(start)
		runs = append(runs, summarize(latencies, *size, elapsed, readCPUTime().sub(cpu0)))
	}

	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			return err
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			return err
		}
	}

	warnIfLoaded()

	var path *pathInfo
	if pr, ok := c.(pathReporter); ok {
		path = pr.Path()
	}
	if err := path.validate(); err != nil {
		return err
	}
	return report(collect(runs, path))
}

// run holds the measurements of one repetition. A latency is the time for
// one batch of -inflight messages to go out and come back, so a batch moves
// 2*size*inflight bytes and PerMsg divides the batch latency by the number
// of messages in it.
type run struct {
	Count      int     `json:"round_trips"`
	Messages   int     `json:"messages"`
	PerMsg     float64 `json:"p50_us_per_message"`
	MsgRate    float64 `json:"messages_per_sec"`
	Elapsed    float64 `json:"elapsed_sec"`
	Throughput float64 `json:"throughput_mbps"` // MB/s, both directions
	RPS        float64 `json:"round_trips_per_sec"`
	Min        float64 `json:"min_us"`
	P50        float64 `json:"p50_us"`
	P90        float64 `json:"p90_us"`
	P99        float64 `json:"p99_us"`
	Max        float64 `json:"max_us"`
	Mean       float64 `json:"mean_us"`
	StdDev     float64 `json:"stddev_us"`

	// CPU accounting. CPUPerMsg is the work a message costs; PerMsg minus
	// CPUPerMsg is what it spent waiting.
	CPU       cpuTime `json:"cpu"`
	CPUPerMsg float64 `json:"cpu_us_per_message"`
	Busy      float64 `json:"cpu_busy_fraction"`
}

// result is the full record of a client invocation: the median run's
// numbers inline, every run for inspection, and enough configuration and
// environment to interpret them without the command line that produced it.
type result struct {
	Label     string    `json:"label"`
	Impl      string    `json:"impl"`
	Size      int       `json:"payload_bytes"`
	Inflight  int       `json:"inflight"`
	RecvBatch bool      `json:"receive_batch"`
	Warmup    int       `json:"warmup_round_trips"`
	Repeat    int       `json:"repetitions"`
	Path      *pathInfo `json:"path,omitempty"`
	Env       env       `json:"env"`
	run
	Runs []run `json:"runs,omitempty"`
}

// collect picks the median repetition by p50 latency as the headline
// result. A median is resistant to the one repetition that caught a
// scheduler hiccup in a way a mean of means is not.
func collect(runs []run, path *pathInfo) result {
	byP50 := make([]run, len(runs))
	copy(byP50, runs)
	sort.Slice(byP50, func(i, j int) bool { return byP50[i].P50 < byP50[j].P50 })

	name := *label
	if name == "" {
		name = *impl
	}
	r := result{
		Label:     name,
		Impl:      *impl,
		Size:      *size,
		Inflight:  max(*inflight, 1),
		RecvBatch: *recvBatch,
		Warmup:    *warmup,
		Repeat:    len(runs),
		Path:      path,
		Env:       captureEnv(),
		run:       byP50[len(byP50)/2],
	}
	if len(runs) > 1 {
		r.Runs = runs
	}
	return r
}

// env records what the measurement was taken on, so a number found later
// in a log can still be placed.
type env struct {
	Time       string `json:"time"`
	Host       string `json:"host"`
	GoVersion  string `json:"go_version"`
	GOMAXPROCS int    `json:"gomaxprocs"`
	NumCPU     int    `json:"num_cpu"`
	OS         string `json:"os_version,omitempty"`
	Kernel     string `json:"kernel,omitempty"`
	Commit     string `json:"git_commit,omitempty"`
	Dirty      int    `json:"git_dirty_files"`
	LoadAvg    string `json:"load_average,omitempty"`
}

func captureEnv() env {
	host, _ := os.Hostname()
	e := env{
		Time:       time.Now().UTC().Format(time.RFC3339),
		Host:       host,
		GoVersion:  runtime.Version(),
		GOMAXPROCS: runtime.GOMAXPROCS(0),
		NumCPU:     runtime.NumCPU(),
		OS:         cmdOutput("sw_vers", "-productVersion"),
		Kernel:     cmdOutput("uname", "-r"),
		Commit:     cmdOutput("git", "rev-parse", "HEAD"),
		LoadAvg:    cmdOutput("sysctl", "-n", "vm.loadavg"),
	}
	if dirty := cmdOutput("git", "status", "--porcelain"); dirty != "" {
		e.Dirty = strings.Count(dirty, "\n") + 1
	}
	return e
}

// warnIfLoaded says so when the machine was busy enough that the numbers
// describe the scheduler more than the transport. The load average is
// recorded in the result either way; this just makes it hard to miss.
func warnIfLoaded() {
	// sysctl reports vm.loadavg as "{ 1.23 4.56 7.89 }".
	fields := strings.Fields(strings.Trim(cmdOutput("sysctl", "-n", "vm.loadavg"), "{} "))
	if len(fields) == 0 {
		return
	}
	load, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return
	}
	if load > float64(runtime.NumCPU())/2 {
		fmt.Fprintf(os.Stderr, "netperfbench: load average %.1f on %d CPUs — treat these numbers as noise\n",
			load, runtime.NumCPU())
	}
}

func cmdOutput(name string, args ...string) string {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func summarize(latencies []time.Duration, size int, elapsed time.Duration, cpu cpuTime) run {
	sorted := make([]time.Duration, len(latencies))
	copy(sorted, latencies)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	us := func(d time.Duration) float64 { return float64(d) / float64(time.Microsecond) }
	pct := func(p float64) float64 {
		if len(sorted) == 0 {
			return 0
		}
		i := int(p * float64(len(sorted)-1))
		return us(sorted[i])
	}

	var sum float64
	for _, d := range latencies {
		sum += us(d)
	}
	mean := sum / float64(len(latencies))
	var variance float64
	for _, d := range latencies {
		diff := us(d) - mean
		variance += diff * diff
	}
	variance /= float64(len(latencies))

	n := max(*inflight, 1)
	messages := len(latencies) * n
	bytes := float64(messages) * float64(size) * 2
	return run{
		Count:      len(latencies),
		Messages:   messages,
		PerMsg:     pct(0.50) / float64(n),
		MsgRate:    float64(messages) / elapsed.Seconds(),
		Elapsed:    elapsed.Seconds(),
		Throughput: bytes / elapsed.Seconds() / (1 << 20),
		RPS:        float64(len(latencies)) / elapsed.Seconds(),
		Min:        pct(0),
		P50:        pct(0.50),
		P90:        pct(0.90),
		P99:        pct(0.99),
		Max:        us(sorted[len(sorted)-1]),
		Mean:       mean,
		StdDev:     math.Sqrt(variance),
		CPU:        cpu,
		CPUPerMsg:  cpu.total() / float64(messages),
		Busy:       cpu.busy(elapsed),
	}
}

func report(r result) error {
	if *asJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(r)
	}
	fmt.Printf("%s: %d batches of %d x %d bytes in %.2fs\n", r.Label, r.Count, r.Inflight, r.Size, r.Elapsed)
	fmt.Printf("  latency us   min %.1f  p50 %.1f  p90 %.1f  p99 %.1f  max %.1f  mean %.1f ±%.1f\n",
		r.Min, r.P50, r.P90, r.P99, r.Max, r.Mean, r.StdDev)
	fmt.Printf("  %.1f us/message, %.0f messages/sec, %.1f MB/s\n", r.PerMsg, r.MsgRate, r.Throughput)
	fmt.Printf("  cpu %.1f us/message (%.1f user, %.1f sys), %.2f cores busy, %.1f us/message waiting\n",
		r.CPUPerMsg, r.CPU.User/float64(r.Messages), r.CPU.Sys/float64(r.Messages),
		r.Busy, r.PerMsg-r.CPUPerMsg)
	if r.Repeat > 1 {
		fmt.Printf("  median of %d repetitions\n", r.Repeat)
	}
	if r.Path != nil {
		fmt.Printf("  path %s via %s\n", r.Path.Status, strings.Join(r.Path.Interfaces, ","))
	}
	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
