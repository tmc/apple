// Command rdmainfo probes Apple's RDMA userspace bindings.
package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/rdma"
)

type output struct {
	Available bool          `json:"available"`
	State     string        `json:"state,omitempty"`
	Devices   []deviceInfo  `json:"devices,omitempty"`
	Steps     []stepResult  `json:"steps,omitempty"`
	Notes     []string      `json:"notes,omitempty"`
	Symbols   []symbolGroup `json:"symbols,omitempty"`
}

type deviceInfo struct {
	Index        int    `json:"index"`
	Name         string `json:"name"`
	NetInterface string `json:"net_interface,omitempty"`
	Handle       string `json:"handle"`
}

type stepResult struct {
	Name    string         `json:"name"`
	OK      bool           `json:"ok"`
	Return  int            `json:"return,omitempty"`
	Handle  string         `json:"handle,omitempty"`
	Bytes   int            `json:"bytes,omitempty"`
	Preview string         `json:"preview,omitempty"`
	Fields  map[string]any `json:"fields,omitempty"`
	Error   string         `json:"error,omitempty"`
}

type symbolGroup struct {
	Name    string   `json:"name"`
	Symbols []string `json:"symbols"`
}

func main() {
	if len(os.Args) == 1 {
		status(nil)
		return
	}

	switch os.Args[1] {
	case "status":
		status(os.Args[2:])
	case "list":
		list(os.Args[2:])
	case "features":
		features(os.Args[2:])
	case "scan":
		scan(os.Args[2:])
	case "matrix":
		matrix(os.Args[2:])
	case "exercise":
		exercise(os.Args[2:])
	case "open":
		openDevice(os.Args[2:])
	case "query":
		query(os.Args[2:])
	case "lifecycle":
		lifecycle(os.Args[2:])
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "rdmainfo: unknown command %q\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `usage: rdmainfo <command> [options]

Commands:
  status      Report whether librdma and probe symbols are available.
  list        List RDMA devices.
  features    Show generated symbols and what this tool exercises.
  scan        Run safe bring-up checks: status, list, open, query.
  matrix      Run scan-style checks across all devices and selected ports.
  exercise    Exercise every generated verb that is safe on this host.
  open        Open and close one RDMA device.
  query       Open a device and call ibv_query_device / ibv_query_port.
  lifecycle   Open a device, allocate PD, create CQ, optionally register memory.

Common options:
  -json              Print JSON instead of text for commands that report data.
  -require           Exit non-zero if RDMA is unavailable or no device is present.
  -device N          Select device index from list output.
  -name substring    Select first device whose name contains substring.

Default behavior treats "librdma unavailable" and "no RDMA devices" as normal
states. Use -require for CI or hardware bring-up once RDMA should be enabled.
`)
}

func status(args []string) {
	fs := flag.NewFlagSet("status", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	require := fs.Bool("require", false, "exit non-zero if unavailable or empty")
	fs.Parse(args)

	out := output{Available: rdma.Available()}
	if !out.Available {
		out.State = "unavailable"
		out.Notes = append(out.Notes, "/usr/lib/librdma.dylib or required probe symbols are not present")
		printOutput(out, *jsonOut)
		if *require {
			os.Exit(1)
		}
		return
	}
	devs, err := rdma.Devices()
	if err != nil {
		fatalf("rdma devices: %v", err)
	}
	out.Devices = deviceInfos(devs)
	if len(devs) == 0 {
		out.State = "available-no-devices"
		out.Notes = append(out.Notes, "librdma loaded, but no RDMA devices are reported")
		printOutput(out, *jsonOut)
		if *require {
			os.Exit(1)
		}
		return
	}
	out.State = "devices-present"
	printOutput(out, *jsonOut)
}

func list(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	require := fs.Bool("require", false, "exit non-zero if unavailable or empty")
	fs.Parse(args)

	devs, ok := devicesOrExplain(*require, *jsonOut)
	if !ok {
		return
	}
	out := output{Available: true, Devices: deviceInfos(devs)}
	if len(devs) == 0 {
		out.State = "available-no-devices"
		out.Notes = append(out.Notes, "no RDMA devices reported")
		printOutput(out, *jsonOut)
		return
	}
	out.State = "devices-present"
	printOutput(out, *jsonOut)
}

func features(args []string) {
	fs := flag.NewFlagSet("features", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	fs.Parse(args)

	out := output{
		Available: rdma.Available(),
		Symbols: []symbolGroup{
			{"high-level helpers", []string{"Available", "Devices"}},
			{"device discovery", []string{"Ibv_get_device_list", "Ibv_free_device_list", "Ibv_get_device_name"}},
			{"context lifecycle", []string{"Ibv_open_device", "Ibv_close_device"}},
			{"resource lifecycle", []string{"Ibv_alloc_pd", "Ibv_dealloc_pd", "Ibv_create_cq", "Ibv_destroy_cq"}},
			{"queries", []string{"Ibv_query_device", "Ibv_query_port"}},
			{"manual/advanced", []string{"Ibv_reg_mr", "Ibv_dereg_mr", "Ibv_create_qp", "Ibv_modify_qp", "Ibv_destroy_qp"}},
		},
		Notes: []string{
			"QP creation/modification is listed but not exercised because generated attr structs are not available yet.",
			"MR registration is opt-in in lifecycle with -register-memory.",
			"Pointer-returning verbs may fail by returning nil; lifecycle reports errno when librdma sets it.",
		},
	}
	if *jsonOut {
		writeJSON(out)
		return
	}
	fmt.Printf("librdma available: %v\n\n", out.Available)
	for _, group := range out.Symbols {
		fmt.Println(group.Name + ":")
		for _, sym := range group.Symbols {
			fmt.Println("  " + sym)
		}
	}
	fmt.Println()
	for _, note := range out.Notes {
		fmt.Println("note: " + note)
	}
}

func scan(args []string) {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	selector := addDeviceFlags(fs)
	jsonOut := fs.Bool("json", false, "print JSON")
	require := fs.Bool("require", false, "exit non-zero if unavailable or empty")
	port := fs.Int("port", 1, "port number for ibv_query_port")
	attrSize := fs.Int("attr-size", 4096, "raw byte buffer size for query calls")
	fs.Parse(args)
	portNum := parsePortFlag("port", *port)

	out := output{Available: rdma.Available()}
	dev, ok := pickDevice(selector, *require, *jsonOut, &out)
	if !ok {
		printOutput(out, *jsonOut)
		return
	}
	out.Devices = []deviceInfo{deviceInfoFrom(dev, selector.index)}
	ctx, ok := openContext(dev, *require, &out)
	if !ok {
		printOutput(out, *jsonOut)
		return
	}

	runQuery(ctx, portNum, *attrSize, *attrSize, 16, &out)
	closeContext(ctx, &out)
	printOutput(out, *jsonOut)
}

func matrix(args []string) {
	fs := flag.NewFlagSet("matrix", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	require := fs.Bool("require", false, "exit non-zero if unavailable or empty")
	portsText := fs.String("ports", "1,2", "comma-separated port numbers to query")
	attrSize := fs.Int("attr-size", 4096, "raw byte buffer size for query calls")
	preview := fs.Int("preview", 16, "bytes of successful query buffers to print")
	fs.Parse(args)

	out := output{Available: rdma.Available()}
	devs, ok := devicesOrExplain(*require, *jsonOut)
	if !ok {
		out.State = "unavailable"
		printOutput(out, *jsonOut)
		return
	}
	out.Devices = deviceInfos(devs)
	if len(devs) == 0 {
		out.State = "available-no-devices"
		out.Notes = append(out.Notes, "no RDMA devices reported")
		printOutput(out, *jsonOut)
		return
	}
	ports := parsePorts(*portsText)
	for i, dev := range devs {
		ctx, ok := openContext(dev, *require, &out)
		if !ok {
			continue
		}
		out.Steps = append(out.Steps, stepResult{
			Name:   fmt.Sprintf("device[%d].name", i),
			OK:     true,
			Handle: dev.Name,
		})
		appendQueryDevice(ctx, fmt.Sprintf("device[%d].ibv_query_device", i), *attrSize, *preview, &out)
		for _, port := range ports {
			appendQueryPort(ctx, port, fmt.Sprintf("device[%d].port[%d].ibv_query_port", i, port), *attrSize, *preview, &out)
		}
		closeContext(ctx, &out)
	}
	printOutput(out, *jsonOut)
}

func exercise(args []string) {
	fs := flag.NewFlagSet("exercise", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "print JSON")
	require := fs.Bool("require", false, "exit non-zero if unavailable or empty")
	portsText := fs.String("ports", "1", "comma-separated port numbers to query")
	attrSize := fs.Int("attr-size", 4096, "raw byte buffer size for query calls")
	preview := fs.Int("preview", 0, "bytes of successful query buffers to print")
	cqe := fs.Int("cqe", 16, "completion queue entries")
	registerMemory := fs.Bool("register-memory", false, "also register a small memory region when PD allocation works")
	mrBytes := fs.Int("mr-bytes", 4096, "bytes to allocate when -register-memory is set")
	access := fs.Int("mr-access", 0, "ibv_reg_mr access flags")
	fs.Parse(args)

	out := output{Available: rdma.Available()}
	devs, ok := devicesOrExplain(*require, *jsonOut)
	if !ok {
		out.State = "unavailable"
		printOutput(out, *jsonOut)
		return
	}
	out.Devices = deviceInfos(devs)
	if len(devs) == 0 {
		out.State = "available-no-devices"
		out.Notes = append(out.Notes, "no RDMA devices reported")
		printOutput(out, *jsonOut)
		return
	}
	out.State = "devices-present"
	ports := parsePorts(*portsText)
	for i, dev := range devs {
		ctx, ok := openContext(dev, *require, &out)
		if !ok {
			continue
		}
		out.Steps = append(out.Steps, stepResult{
			Name:   fmt.Sprintf("device[%d].name", i),
			OK:     true,
			Handle: dev.Name,
		})
		appendQueryDevice(ctx, fmt.Sprintf("device[%d].ibv_query_device", i), *attrSize, *preview, &out)
		active := false
		for _, port := range ports {
			step := queryPort(ctx, port, fmt.Sprintf("device[%d].port[%d].ibv_query_port", i, port), *attrSize, *preview)
			if step.OK && step.Fields["state"] == int32(4) {
				active = true
			}
			out.Steps = append(out.Steps, step)
		}
		if !active {
			out.Steps = append(out.Steps, stepResult{
				Name:  fmt.Sprintf("device[%d].rdma_datapath_ready", i),
				Error: "skipped resource verbs because no queried port is PORT_ACTIVE",
			})
			closeContext(ctx, &out)
			continue
		}
		pd, pdOK := allocPD(ctx, &out)
		cq, cqOK := createCQ(ctx, *cqe, &out)
		if *registerMemory && pdOK {
			registerMR(pd, *mrBytes, *access, &out)
		} else if *registerMemory {
			out.Steps = append(out.Steps, stepResult{
				Name:  "ibv_reg_mr",
				Error: "skipped because ibv_alloc_pd did not return a protection domain",
			})
		}
		if !pdOK {
			out.Steps = append(out.Steps, stepResult{
				Name:  "ibv_create_qp",
				Error: "skipped because ibv_alloc_pd did not return a protection domain",
			})
			out.Steps = append(out.Steps, stepResult{
				Name:  "ibv_modify_qp",
				Error: "skipped because no queue pair was created",
			})
			out.Steps = append(out.Steps, stepResult{
				Name:  "ibv_destroy_qp",
				Error: "skipped because no queue pair was created",
			})
		}
		if cqOK {
			destroyCQ(cq, &out)
		}
		if pdOK {
			deallocPD(pd, &out)
		}
		closeContext(ctx, &out)
	}
	out.Notes = append(out.Notes,
		"device discovery uses Ibv_get_device_list, Ibv_get_device_name, and Ibv_free_device_list through rdma.Devices",
		"QP creation requires a protection domain and ibv_qp_init_attr; this command reports it as skipped when PD allocation fails",
	)
	printOutput(out, *jsonOut)
}

func openDevice(args []string) {
	fs := flag.NewFlagSet("open", flag.ExitOnError)
	selector := addDeviceFlags(fs)
	jsonOut := fs.Bool("json", false, "print JSON")
	require := fs.Bool("require", false, "exit non-zero if unavailable or empty")
	fs.Parse(args)

	out := output{Available: rdma.Available()}
	dev, ok := pickDevice(selector, *require, *jsonOut, &out)
	if !ok {
		printOutput(out, *jsonOut)
		return
	}
	out.Devices = []deviceInfo{deviceInfoFrom(dev, selector.index)}
	ctx, ok := openContext(dev, *require, &out)
	if ok {
		closeContext(ctx, &out)
	}
	printOutput(out, *jsonOut)
}

func query(args []string) {
	fs := flag.NewFlagSet("query", flag.ExitOnError)
	selector := addDeviceFlags(fs)
	jsonOut := fs.Bool("json", false, "print JSON")
	require := fs.Bool("require", false, "exit non-zero if unavailable or empty")
	port := fs.Int("port", 1, "port number for ibv_query_port")
	deviceAttrSize := fs.Int("device-attr-size", 4096, "raw byte buffer size for ibv_query_device")
	portAttrSize := fs.Int("port-attr-size", 4096, "raw byte buffer size for ibv_query_port")
	preview := fs.Int("preview", 32, "bytes of successful query buffers to print")
	fs.Parse(args)
	portNum := parsePortFlag("port", *port)

	out := output{Available: rdma.Available()}
	dev, ok := pickDevice(selector, *require, *jsonOut, &out)
	if !ok {
		printOutput(out, *jsonOut)
		return
	}
	out.Devices = []deviceInfo{deviceInfoFrom(dev, selector.index)}
	ctx, ok := openContext(dev, *require, &out)
	if !ok {
		printOutput(out, *jsonOut)
		return
	}

	runQuery(ctx, portNum, *deviceAttrSize, *portAttrSize, *preview, &out)
	closeContext(ctx, &out)
	printOutput(out, *jsonOut)
}

func lifecycle(args []string) {
	fs := flag.NewFlagSet("lifecycle", flag.ExitOnError)
	selector := addDeviceFlags(fs)
	jsonOut := fs.Bool("json", false, "print JSON")
	require := fs.Bool("require", false, "exit non-zero if unavailable or empty")
	cqe := fs.Int("cqe", 16, "completion queue entries")
	registerMemory := fs.Bool("register-memory", false, "also register a small memory region")
	mrBytes := fs.Int("mr-bytes", 4096, "bytes to allocate when -register-memory is set")
	access := fs.Int("mr-access", 0, "ibv_reg_mr access flags")
	fs.Parse(args)

	out := output{Available: rdma.Available()}
	dev, ok := pickDevice(selector, *require, *jsonOut, &out)
	if !ok {
		printOutput(out, *jsonOut)
		return
	}
	out.Devices = []deviceInfo{deviceInfoFrom(dev, selector.index)}
	ctx, ok := openContext(dev, *require, &out)
	if !ok {
		printOutput(out, *jsonOut)
		return
	}

	step := queryPort(ctx, 1, "ibv_query_port", int(unsafe.Sizeof(ibvPortAttr{})), 0)
	out.Steps = append(out.Steps, step)
	if !step.OK || step.Fields["state"] != int32(4) {
		out.Steps = append(out.Steps, stepResult{
			Name:  "rdma_datapath_ready",
			Error: "skipped resource verbs because port 1 is not PORT_ACTIVE",
		})
		closeContext(ctx, &out)
		printOutput(out, *jsonOut)
		return
	}

	pd, pdOK := allocPD(ctx, &out)
	cq, cqOK := createCQ(ctx, *cqe, &out)

	if *registerMemory && pdOK {
		registerMR(pd, *mrBytes, *access, &out)
	} else if *registerMemory {
		out.Steps = append(out.Steps, stepResult{
			Name:  "ibv_reg_mr",
			Error: "skipped because ibv_alloc_pd did not return a protection domain",
		})
	}
	if cqOK {
		destroyCQ(cq, &out)
	}
	if pdOK {
		deallocPD(pd, &out)
	}
	closeContext(ctx, &out)
	printOutput(out, *jsonOut)
}

type deviceSelector struct {
	index int
	name  string
}

type ibvDeviceAttr struct {
	FWVer                 [64]byte
	NodeGUID              uint64
	SysImageGUID          uint64
	MaxMRSize             uint64
	PageSizeCap           uint64
	VendorID              uint32
	VendorPartID          uint32
	HWVer                 uint32
	MaxQP                 int32
	MaxQPWR               int32
	DeviceCapFlags        uint32
	MaxSGE                int32
	MaxSGERD              int32
	MaxCQ                 int32
	MaxCQE                int32
	MaxMR                 int32
	MaxPD                 int32
	MaxQPRDAtom           int32
	MaxEERDAtom           int32
	MaxResRDAtom          int32
	MaxQPInitRDAtom       int32
	MaxEEInitRDAtom       int32
	AtomicCap             int32
	MaxEE                 int32
	MaxRDD                int32
	MaxMW                 int32
	MaxRawIPv6QP          int32
	MaxRawEthYQP          int32
	MaxMcastGrp           int32
	MaxMcastQPAttach      int32
	MaxTotalMcastQPAttach int32
	MaxAH                 int32
	MaxFMR                int32
	MaxMapPerFMR          int32
	MaxSRQ                int32
	MaxSRQWR              int32
	MaxSRQSGE             int32
	MaxPKeys              uint16
	LocalCAAckDelay       uint8
	PhysPortCnt           uint8
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

func addDeviceFlags(fs *flag.FlagSet) *deviceSelector {
	s := &deviceSelector{}
	fs.IntVar(&s.index, "device", 0, "device index from list")
	fs.StringVar(&s.name, "name", "", "select first device whose name contains substring")
	return s
}

func devicesOrExplain(require, jsonOut bool) ([]rdma.Device, bool) {
	if !rdma.Available() {
		if jsonOut {
			writeJSON(output{
				Available: false,
				State:     "unavailable",
				Notes:     []string{"librdma or required probe symbols are not present"},
			})
		} else {
			fmt.Println("RDMA unavailable: librdma or required probe symbols are not present")
		}
		if require {
			os.Exit(1)
		}
		return nil, false
	}
	devs, err := rdma.Devices()
	if err != nil {
		fatalf("rdma devices: %v", err)
	}
	if len(devs) == 0 && require {
		fmt.Fprintln(os.Stderr, "rdmainfo: no RDMA devices reported")
		os.Exit(1)
	}
	return devs, true
}

func pickDevice(selector *deviceSelector, require, jsonOut bool, out *output) (rdma.Device, bool) {
	devs, ok := devicesOrExplain(require, jsonOut)
	if !ok {
		out.State = "unavailable"
		return rdma.Device{}, false
	}
	out.Devices = deviceInfos(devs)
	if len(devs) == 0 {
		out.State = "available-no-devices"
		out.Notes = append(out.Notes, "no RDMA devices reported")
		return rdma.Device{}, false
	}
	if selector.name != "" {
		for i, dev := range devs {
			if strings.Contains(dev.Name, selector.name) {
				selector.index = i
				return dev, true
			}
		}
		fatalf("no RDMA device name contains %q", selector.name)
	}
	if selector.index < 0 || selector.index >= len(devs) {
		fatalf("device index %d out of range [0,%d)", selector.index, len(devs))
	}
	return devs[selector.index], true
}

func openContext(dev rdma.Device, require bool, out *output) (rdma.RDMAContext, bool) {
	ctx, err := rdma.Ibv_open_device(dev.Handle)
	step := stepResult{Name: "ibv_open_device", OK: err == nil && ctx != 0, Handle: hexHandle(ctx)}
	if err != nil {
		step.Error = err.Error()
		out.Steps = append(out.Steps, step)
		fatalf("ibv_open_device: %v", err)
	}
	if ctx == 0 {
		step.Error = "returned nil context"
		out.Steps = append(out.Steps, step)
		if require {
			os.Exit(1)
		}
		return 0, false
	}
	out.Steps = append(out.Steps, step)
	return ctx, true
}

func closeContext(ctx rdma.RDMAContext, out *output) {
	rc, err := rdma.Ibv_close_device(ctx)
	out.Steps = append(out.Steps, resultStep("ibv_close_device", rc, err))
}

func allocPD(ctx rdma.RDMAContext, out *output) (rdma.RDMAPD, bool) {
	setErrno(0)
	pd, err := rdma.Ibv_alloc_pd(ctx)
	callErrno := errno()
	step := stepResult{Name: "ibv_alloc_pd", OK: err == nil && pd != 0, Handle: hexHandle(pd)}
	if err != nil {
		step.Error = err.Error()
		out.Steps = append(out.Steps, step)
		return 0, false
	}
	if pd == 0 {
		step.Error = "returned nil protection domain"
		if callErrno != 0 {
			step.Error += ": " + errnoName(callErrno)
		}
		out.Steps = append(out.Steps, step)
		return 0, false
	}
	out.Steps = append(out.Steps, step)
	return pd, true
}

func deallocPD(pd rdma.RDMAPD, out *output) {
	rc, err := rdma.Ibv_dealloc_pd(pd)
	out.Steps = append(out.Steps, resultStep("ibv_dealloc_pd", rc, err))
}

func createCQ(ctx rdma.RDMAContext, cqe int, out *output) (rdma.RDMACQ, bool) {
	setErrno(0)
	cq, err := rdma.Ibv_create_cq(ctx, cqe, 0, 0, 0)
	callErrno := errno()
	step := stepResult{Name: "ibv_create_cq", OK: err == nil && cq != 0, Handle: hexHandle(cq)}
	if err != nil {
		step.Error = err.Error()
		out.Steps = append(out.Steps, step)
		return 0, false
	}
	if cq == 0 {
		step.Error = "returned nil completion queue"
		if callErrno != 0 {
			step.Error += ": " + errnoName(callErrno)
		}
		out.Steps = append(out.Steps, step)
		return 0, false
	}
	out.Steps = append(out.Steps, step)
	return cq, true
}

func destroyCQ(cq rdma.RDMACQ, out *output) {
	rc, err := rdma.Ibv_destroy_cq(cq)
	out.Steps = append(out.Steps, resultStep("ibv_destroy_cq", rc, err))
}

func registerMR(pd rdma.RDMAPD, n int, access int, out *output) {
	if n <= 0 {
		out.Steps = append(out.Steps, stepResult{Name: "ibv_reg_mr", Error: "mr-bytes must be positive"})
		return
	}
	buf := make([]byte, n)
	setErrno(0)
	mr, err := rdma.Ibv_reg_mr(pd, uintptr(unsafe.Pointer(unsafe.SliceData(buf))), uintptr(len(buf)), access)
	callErrno := errno()
	step := stepResult{Name: "ibv_reg_mr", OK: err == nil && mr != 0, Handle: hexHandle(mr), Bytes: len(buf)}
	if err != nil {
		step.Error = err.Error()
		out.Steps = append(out.Steps, step)
		return
	}
	if mr == 0 {
		step.Error = "returned nil memory region"
		if callErrno != 0 {
			step.Error += ": " + errnoName(callErrno)
		}
		out.Steps = append(out.Steps, step)
		return
	}
	out.Steps = append(out.Steps, step)
	rc, err := rdma.Ibv_dereg_mr(mr)
	out.Steps = append(out.Steps, resultStep("ibv_dereg_mr", rc, err))
}

func runQuery(ctx rdma.RDMAContext, port uint8, deviceAttrSize, portAttrSize, preview int, out *output) {
	appendQueryDevice(ctx, "ibv_query_device", deviceAttrSize, preview, out)
	appendQueryPort(ctx, port, "ibv_query_port", portAttrSize, preview, out)
}

func appendQueryDevice(ctx rdma.RDMAContext, name string, size, preview int, out *output) {
	if size == 0 {
		return
	}
	buf := queryBuffer(size, int(unsafe.Sizeof(ibvDeviceAttr{})), "device attr")
	rc, err := rdma.Ibv_query_device(ctx, uintptr(unsafe.Pointer(unsafe.SliceData(buf))))
	step := queryStep(name, rc, err, buf, preview)
	if step.OK {
		step.Fields = deviceAttrFields(buf)
	}
	out.Steps = append(out.Steps, step)
}

func appendQueryPort(ctx rdma.RDMAContext, port uint8, name string, size, preview int, out *output) {
	out.Steps = append(out.Steps, queryPort(ctx, port, name, size, preview))
}

func queryPort(ctx rdma.RDMAContext, port uint8, name string, size, preview int) stepResult {
	if size == 0 {
		return stepResult{Name: name}
	}
	buf := queryBuffer(size, int(unsafe.Sizeof(ibvPortAttr{})), "port attr")
	rc, err := rdma.Ibv_query_port(ctx, port, uintptr(unsafe.Pointer(unsafe.SliceData(buf))))
	step := queryStep(name, rc, err, buf, preview)
	if step.OK {
		step.Fields = portAttrFields(buf)
	}
	return step
}

func queryBuffer(size, min int, name string) []byte {
	if size < 0 {
		fatalf("%s buffer size must be non-negative", name)
	}
	if size == 0 {
		return nil
	}
	if size > 0 && size < min {
		fatalf("%s buffer size %d is smaller than required %d", name, size, min)
	}
	return make([]byte, size)
}

func queryStep(name string, rc int, err error, buf []byte, preview int) stepResult {
	step := resultStep(name, rc, err)
	step.Bytes = len(buf)
	if err == nil && rc == 0 && preview > 0 {
		if preview > len(buf) {
			preview = len(buf)
		}
		step.Preview = hex.EncodeToString(buf[:preview])
	}
	return step
}

func deviceAttrFields(buf []byte) map[string]any {
	if len(buf) < int(unsafe.Sizeof(ibvDeviceAttr{})) {
		return nil
	}
	attr := (*ibvDeviceAttr)(unsafe.Pointer(unsafe.SliceData(buf)))
	return map[string]any{
		"fw_ver":           strings.TrimRight(string(attr.FWVer[:]), "\x00"),
		"node_guid":        fmt.Sprintf("0x%016x", attr.NodeGUID),
		"sys_image_guid":   fmt.Sprintf("0x%016x", attr.SysImageGUID),
		"max_mr_size":      attr.MaxMRSize,
		"page_size_cap":    attr.PageSizeCap,
		"vendor_id":        attr.VendorID,
		"vendor_part_id":   attr.VendorPartID,
		"hw_ver":           attr.HWVer,
		"max_qp":           attr.MaxQP,
		"max_qp_wr":        attr.MaxQPWR,
		"device_cap_flags": fmt.Sprintf("0x%x", attr.DeviceCapFlags),
		"max_sge":          attr.MaxSGE,
		"max_cq":           attr.MaxCQ,
		"max_cqe":          attr.MaxCQE,
		"max_mr":           attr.MaxMR,
		"max_pd":           attr.MaxPD,
		"atomic_cap":       attr.AtomicCap,
		"max_pkeys":        attr.MaxPKeys,
		"phys_port_cnt":    attr.PhysPortCnt,
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
		"max_mtu":          attr.MaxMTU,
		"max_mtu_bytes":    mtuBytes(attr.MaxMTU),
		"active_mtu":       attr.ActiveMTU,
		"active_mtu_bytes": mtuBytes(attr.ActiveMTU),
		"gid_tbl_len":      attr.GIDTblLen,
		"port_cap_flags":   fmt.Sprintf("0x%x", attr.PortCapFlags),
		"max_msg_sz":       attr.MaxMsgSZ,
		"bad_pkey_cntr":    attr.BadPKeyCntr,
		"qkey_viol_cntr":   attr.QKeyViolCntr,
		"pkey_tbl_len":     attr.PKeyTblLen,
		"lid":              attr.LID,
		"sm_lid":           attr.SMLID,
		"active_width":     attr.ActiveWidth,
		"active_speed":     attr.ActiveSpeed,
		"phys_state":       attr.PhysState,
		"link_layer":       attr.LinkLayer,
		"link_layer_name":  linkLayerName(attr.LinkLayer),
		"flags":            fmt.Sprintf("0x%x", attr.Flags),
		"port_cap_flags2":  fmt.Sprintf("0x%x", attr.PortCapFlags2),
	}
}

func resultStep(name string, rc int, err error) stepResult {
	step := stepResult{Name: name, OK: err == nil && rc == 0, Return: rc}
	if err != nil {
		step.Error = err.Error()
	} else if rc != 0 {
		step.Error = errnoName(rc)
	}
	return step
}

func deviceInfos(devs []rdma.Device) []deviceInfo {
	out := make([]deviceInfo, 0, len(devs))
	for i, dev := range devs {
		out = append(out, deviceInfoFrom(dev, i))
	}
	return out
}

func deviceInfoFrom(dev rdma.Device, index int) deviceInfo {
	return deviceInfo{
		Index:        index,
		Name:         dev.Name,
		NetInterface: rdmaNetInterface(dev.Name),
		Handle:       hexHandle(dev.Handle),
	}
}

func printOutput(out output, jsonOut bool) {
	if jsonOut {
		writeJSON(out)
		return
	}
	if out.State != "" {
		fmt.Println("state:", out.State)
	}
	fmt.Println("librdma available:", out.Available)
	if len(out.Devices) > 0 {
		fmt.Println("devices:")
		for _, dev := range out.Devices {
			if dev.NetInterface != "" {
				fmt.Printf("  %d\t%s\t%s\t%s\n", dev.Index, dev.Name, dev.NetInterface, dev.Handle)
			} else {
				fmt.Printf("  %d\t%s\t%s\n", dev.Index, dev.Name, dev.Handle)
			}
		}
	}
	for _, step := range out.Steps {
		printStep(step)
	}
	for _, note := range out.Notes {
		fmt.Println("note:", note)
	}
}

func printStep(step stepResult) {
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
	if step.Bytes != 0 {
		fmt.Printf(" bytes=%d", step.Bytes)
	}
	if step.Preview != "" {
		fmt.Printf(" preview=%s", step.Preview)
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

func sortedFieldNames(fields map[string]any) []string {
	names := make([]string, 0, len(fields))
	for name := range fields {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
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

func writeJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fatalf("json: %v", err)
	}
}

func hexHandle[T ~uintptr](v T) string {
	return fmt.Sprintf("0x%x", uintptr(v))
}

func parsePorts(text string) []uint8 {
	var ports []uint8
	for _, field := range strings.Split(text, ",") {
		field = strings.TrimSpace(field)
		if field == "" {
			continue
		}
		port, err := strconv.Atoi(field)
		if err != nil {
			fatalf("invalid port %q", field)
		}
		ports = append(ports, parsePortFlag("port", port))
	}
	if len(ports) == 0 {
		fatalf("no ports selected")
	}
	return ports
}

func parsePortFlag(name string, port int) uint8 {
	if port < 0 || port > 255 {
		fatalf("%s %d out of range [0,255]", name, port)
	}
	return uint8(port)
}

var errorPointer func() unsafe.Pointer

func errno() int {
	initErrno()
	if errorPointer == nil {
		return 0
	}
	return int(*(*int32)(errorPointer()))
}

func setErrno(v int32) {
	initErrno()
	if errorPointer != nil {
		*(*int32)(errorPointer()) = v
	}
}

func initErrno() {
	if errorPointer != nil {
		return
	}
	sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "__error")
	if err != nil || sym == 0 {
		return
	}
	purego.RegisterFunc(&errorPointer, sym)
}

func errnoName(rc int) string {
	switch rc {
	case 1:
		return "EPERM"
	case 2:
		return "ENOENT"
	case 5:
		return "EIO"
	case 6:
		return "ENXIO"
	case 12:
		return "ENOMEM"
	case 13:
		return "EACCES"
	case 16:
		return "EBUSY"
	case 19:
		return "ENODEV"
	case 22:
		return "EINVAL"
	case 38:
		return "ENOSYS"
	case 45:
		return "EOPNOTSUPP"
	case 95:
		return "ENOTSUP"
	default:
		return fmt.Sprintf("errno %d", rc)
	}
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "rdmainfo: "+format+"\n", args...)
	os.Exit(1)
}
