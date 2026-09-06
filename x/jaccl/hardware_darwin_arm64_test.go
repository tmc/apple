//go:build darwin && arm64

package jaccl

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/tmc/apple/rdma"
	xrdma "github.com/tmc/apple/x/rdma"
)

// TestHardwareEvidence is an opt-in, externally coordinated fabric gate.
// Start one test process per rank with identical group/coordinator settings.
// It reports UNMEASURED when that fabric setup is not explicitly requested.
func TestHardwareEvidence(t *testing.T) {
	if os.Getenv("JACCL_HARDWARE") != "1" {
		t.Log("UNMEASURED: set JACCL_HARDWARE=1 for an externally coordinated fabric run")
		return
	}
	rank, err := hardwareEnvInt("JACCL_HARDWARE_RANK")
	if err != nil {
		t.Fatal(err)
	}
	size, err := hardwareEnvIntDefault("JACCL_HARDWARE_SIZE", 2)
	if err != nil || size < 2 || rank < 0 || rank >= size {
		t.Fatalf("JACCL_HARDWARE_SIZE=%d rank=%d: %v", size, rank, err)
	}
	port, err := hardwareEnvInt("JACCL_HARDWARE_PORT")
	if err != nil || port <= 0 || port > 255 {
		t.Fatalf("JACCL_HARDWARE_PORT: %v", err)
	}
	groupID := os.Getenv("JACCL_HARDWARE_GROUP")
	coordinator := os.Getenv("JACCL_HARDWARE_COORDINATOR")
	source := os.Getenv("JACCL_HARDWARE_SOURCE")
	if groupID == "" || coordinator == "" || source == "" {
		t.Fatal("set JACCL_HARDWARE_GROUP, JACCL_HARDWARE_COORDINATOR, and JACCL_HARDWARE_SOURCE")
	}
	if err := hardwareSafetyCheck(coordinator, os.Getenv("JACCL_IBV_DEVICES")); err != nil {
		t.Fatal(err)
	}
	// The full all-dtype corpus performs 51 coordinated operations. Leave room
	// for a physical two-host fabric to complete both admission and completion
	// agreements for each one.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	topology := Mesh(size)
	topologyName := "mesh"
	if os.Getenv("JACCL_HARDWARE_RING") == "1" {
		topology = Ring(size)
		topologyName = "ring"
	}
	group, err := Open(ctx, Config{Rank: rank, Size: size, GroupID: groupID, Coordinator: coordinator, Device: os.Getenv("JACCL_HARDWARE_DEVICE"), Port: uint8(port), Topology: topology})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := group.Close(); err != nil {
			t.Errorf("teardown: %v", err)
		}
	}()
	backend, ok := group.backend.(*nativeBackend)
	if !ok {
		t.Fatal("hardware group did not return native backend")
	}
	device := backend.primaryDevice()
	if device == nil {
		t.Fatal("hardware evidence requires exactly one local RDMA device")
	}
	binaryHash, err := hardwareBinaryHash()
	if err != nil {
		t.Fatal(err)
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		t.Logf("receipt build=%s binary_sha256=%s source=%s topology=%s device=%s port=%d gid_index=%d mtu=%d destinations=%v", info.Main.Version, binaryHash, source, topologyName, device.name, backend.cfg.Port, device.route.Index, device.port.ActiveMTU, backend.destinations)
	} else if rank == 1 {
		t.Logf("receipt build=unknown binary_sha256=%s source=%s topology=%s device=%s port=%d gid_index=%d mtu=%d destinations=%v", binaryHash, source, topologyName, device.name, backend.cfg.Port, device.route.Index, device.port.ActiveMTU, backend.destinations)
	}
	payload := []byte("jaccl hardware evidence payload")
	if rank == 0 {
		if err := group.Send(ctx, 1, payload); err != nil {
			t.Fatal(err)
		}
	} else {
		got := make([]byte, len(payload))
		if err := group.Recv(ctx, 0, got); err != nil {
			t.Fatal(err)
		}
		if gotHash, wantHash := sha256.Sum256(got), sha256.Sum256(payload); gotHash != wantHash {
			t.Fatalf("payload hash = %x, want %x", gotHash, wantHash)
		}
		t.Logf("payload_sha256=%x", sha256.Sum256(got))
	}

	gathered := make([]byte, size*4)
	gatherInput := []byte{byte(rank), byte(rank + 10), byte(rank + 20), byte(rank + 30)}
	if err := group.AllGather(ctx, gathered, gatherInput); err != nil {
		t.Fatal(err)
	}
	wantGathered := make([]byte, len(gathered))
	for peer := 0; peer < size; peer++ {
		copy(wantGathered[peer*4:], []byte{byte(peer), byte(peer + 10), byte(peer + 20), byte(peer + 30)})
	}
	if gotHash, wantHash := sha256.Sum256(gathered), sha256.Sum256(wantGathered); gotHash != wantHash {
		t.Fatalf("all-gather hash = %x, want %x", gotHash, wantHash)
	}
	t.Logf("all_gather_sha256=%x", sha256.Sum256(gathered))

	reduceInput := int32Bytes(int32(rank+1), int32(5-rank))
	for _, test := range []struct {
		name string
		call func([]byte) error
		want []byte
	}{
		{"sum", func(dst []byte) error { return group.AllSum(ctx, dst, reduceInput, Int32) }, int32Bytes(int32(size*(size+1)/2), int32(5*size-size*(size-1)/2))},
		{"max", func(dst []byte) error { return group.AllMax(ctx, dst, reduceInput, Int32) }, int32Bytes(int32(size), 5)},
		{"min", func(dst []byte) error { return group.AllMin(ctx, dst, reduceInput, Int32) }, int32Bytes(1, int32(6-size))},
	} {
		reduced := make([]byte, len(reduceInput))
		if err := test.call(reduced); err != nil {
			t.Fatalf("all-reduce %s: %v", test.name, err)
		}
		if gotHash, wantHash := sha256.Sum256(reduced), sha256.Sum256(test.want); gotHash != wantHash {
			t.Fatalf("all-reduce %s hash = %x, want %x", test.name, gotHash, wantHash)
		}
		t.Logf("all_reduce_int32_%s_sha256=%x", test.name, sha256.Sum256(reduced))
	}
	floatInput := float32Bytes(float32(rank+1), float32(10-rank))
	for _, test := range []struct {
		name string
		call func([]byte) error
		want []byte
	}{
		{"sum", func(dst []byte) error { return group.AllSum(ctx, dst, floatInput, Float32) }, float32Bytes(float32(size*(size+1)/2), float32(10*size-size*(size-1)/2))},
		{"max", func(dst []byte) error { return group.AllMax(ctx, dst, floatInput, Float32) }, float32Bytes(float32(size), 10)},
		{"min", func(dst []byte) error { return group.AllMin(ctx, dst, floatInput, Float32) }, float32Bytes(1, float32(11-size))},
	} {
		reduced := make([]byte, len(floatInput))
		if err := test.call(reduced); err != nil {
			t.Fatalf("float32 all-reduce %s: %v", test.name, err)
		}
		if gotHash, wantHash := sha256.Sum256(reduced), sha256.Sum256(test.want); gotHash != wantHash {
			t.Fatalf("float32 all-reduce %s hash = %x, want %x", test.name, gotHash, wantHash)
		}
		t.Logf("all_reduce_float32_%s_sha256=%x", test.name, sha256.Sum256(reduced))
	}
	allDTypes, err := hardwareAllDTypes(ctx, group, rank, size)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("all_reduce_all_dtypes_sha256=%x", sha256.Sum256(allDTypes))
}

// TestHardwareRouteProbe reports local route candidates without claiming a
// fabric result. It is separate from TestHardwareEvidence so absence of a
// route remains UNMEASURED instead of a collective pass or failure.
func TestHardwareRouteProbe(t *testing.T) {
	if os.Getenv("JACCL_HARDWARE_PROBE") != "1" {
		t.Log("UNMEASURED: set JACCL_HARDWARE_PROBE=1 to inspect local route candidates")
		return
	}
	port, err := hardwareEnvIntDefault("JACCL_HARDWARE_PORT", 1)
	if err != nil || port <= 0 || port > 255 {
		t.Fatalf("JACCL_HARDWARE_PORT: %v", err)
	}
	list, err := rdma.OpenDeviceList()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := list.Close(); err != nil {
			t.Errorf("close device list: %v", err)
		}
	}()
	var candidates int
	for _, device := range list.Devices() {
		context, err := device.Open()
		if err != nil {
			t.Logf("route device=%s port=%d open: %v", device.Name, port, err)
			continue
		}
		var attr rdma.IbvPortAttr
		rc, err := rdma.IbvQueryPortAttr(context, uint8(port), &attr)
		if err != nil || rc != 0 {
			t.Logf("route device=%s port=%d query: %v", device.Name, port, nativeProviderRC("ibv_query_port", rc, err))
		} else {
			t.Logf("route device=%s port=%d state=%d phys_state=%d link_layer=%d gid_table_len=%d mtu=%d", device.Name, port, attr.State, attr.PhysState, attr.LinkLayer, attr.GIDTblLen, attr.ActiveMTU)
			gids := make([]xrdma.RouteGID, 0, xrdma.RouteGIDScanLimit(int32(attr.GIDTblLen)))
			for index := 0; index < cap(gids); index++ {
				var gid rdma.IbvGID
				rc, err := rdma.IbvQueryGidInto(context, uint8(port), index, &gid)
				if err != nil || rc != 0 {
					t.Logf("route device=%s port=%d gid_index=%d query: %v", device.Name, port, index, nativeProviderRC("ibv_query_gid", rc, err))
					continue
				}
				gids = append(gids, xrdma.RouteGID{Index: index, GID: gid})
				t.Logf("route device=%s port=%d gid_index=%d gid=%x nonzero=%t ipv4_mapped=%t", device.Name, port, index, gid, !xrdma.IsZeroGID(gid), xrdma.IsIPv4MappedGID(gid))
			}
			if attr.State == rdma.IBV_PORT_ACTIVE {
				if route, ok := xrdma.SelectRouteGID(gids, -1, attr.LinkLayer); ok {
					candidates++
					t.Logf("route candidate device=%s port=%d gid_index=%d gid=%x mtu=%d", device.Name, port, route.Index, route.GID, attr.ActiveMTU)
				}
			}
		}
		if rc, err := rdma.IbvCloseDevice(context); err != nil || rc != 0 {
			t.Errorf("close route probe device %s: %v", device.Name, nativeProviderRC("ibv_close_device", int(rc), err))
		}
	}
	if candidates == 0 {
		t.Log("UNMEASURED: no safe active local RDMA route")
	}
}

// TestUpstreamHardwareEvidence records the standalone JACCL arm of a matched
// fabric comparison. Run it separately from TestHardwareEvidence on the same
// topology, retaining both receipts for byte-level comparison.
func TestUpstreamHardwareEvidence(t *testing.T) {
	if os.Getenv("JACCL_UPSTREAM_HARDWARE") != "1" {
		t.Log("UNMEASURED: set JACCL_UPSTREAM_HARDWARE=1 for the standalone JACCL fabric arm")
		return
	}
	binaryPath := os.Getenv("JACCL_UPSTREAM_HARDWARE_BINARY")
	if binaryPath == "" {
		t.Fatal("set JACCL_UPSTREAM_HARDWARE_BINARY to upstream_hardware_oracle")
	}
	rank, err := hardwareEnvInt("JACCL_HARDWARE_RANK")
	if err != nil {
		t.Fatal(err)
	}
	size, err := hardwareEnvIntDefault("JACCL_HARDWARE_SIZE", 2)
	if err != nil || size < 2 || rank < 0 || rank >= size {
		t.Fatalf("JACCL_HARDWARE_SIZE=%d rank=%d: %v", size, rank, err)
	}
	if devices := os.Getenv("JACCL_IBV_DEVICES"); devices == "" {
		t.Fatal("set JACCL_IBV_DEVICES for standalone JACCL")
	}
	coordinator := os.Getenv("JACCL_COORDINATOR")
	if coordinator == "" {
		coordinator = os.Getenv("JACCL_HARDWARE_COORDINATOR")
		if coordinator == "" {
			t.Fatal("set JACCL_COORDINATOR or JACCL_HARDWARE_COORDINATOR")
		}
	}
	if err := hardwareSafetyCheck(coordinator, os.Getenv("JACCL_IBV_DEVICES")); err != nil {
		t.Fatal(err)
	}
	if configuredRank, ok := os.LookupEnv("JACCL_RANK"); ok && configuredRank != strconv.Itoa(rank) {
		t.Fatalf("JACCL_RANK %q, want JACCL_HARDWARE_RANK %d", configuredRank, rank)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, binaryPath)
	cmd.Env = os.Environ()
	if _, ok := os.LookupEnv("JACCL_RANK"); !ok {
		cmd.Env = append(cmd.Env, "JACCL_RANK="+strconv.Itoa(rank))
	}
	if os.Getenv("JACCL_COORDINATOR") == "" {
		cmd.Env = append(cmd.Env, "JACCL_COORDINATOR="+coordinator)
	}
	if os.Getenv("JACCL_HARDWARE_RING") == "1" && os.Getenv("JACCL_RING") == "" {
		cmd.Env = append(cmd.Env, "JACCL_RING=1")
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("run standalone JACCL hardware oracle: %v\n%s", err, output)
	}
	payloadHash := fmt.Sprintf("%x", sha256.Sum256([]byte("jaccl hardware evidence payload")))
	gathered := make([]byte, size*4)
	for peer := 0; peer < size; peer++ {
		copy(gathered[peer*4:], []byte{byte(peer), byte(peer + 10), byte(peer + 20), byte(peer + 30)})
	}
	gatherHash := fmt.Sprintf("%x", sha256.Sum256(gathered))
	reduceSumHash := fmt.Sprintf("%x", sha256.Sum256(int32Bytes(int32(size*(size+1)/2), int32(5*size-size*(size-1)/2))))
	reduceMaxHash := fmt.Sprintf("%x", sha256.Sum256(int32Bytes(int32(size), 5)))
	reduceMinHash := fmt.Sprintf("%x", sha256.Sum256(int32Bytes(1, int32(6-size))))
	floatSumHash := fmt.Sprintf("%x", sha256.Sum256(float32Bytes(float32(size*(size+1)/2), float32(10*size-size*(size-1)/2))))
	floatMaxHash := fmt.Sprintf("%x", sha256.Sum256(float32Bytes(float32(size), 10)))
	floatMinHash := fmt.Sprintf("%x", sha256.Sum256(float32Bytes(1, float32(11-size))))
	allDTypes, err := hardwareAllDTypesExpected(size)
	if err != nil {
		t.Fatal(err)
	}
	allDTypesHash := fmt.Sprintf("%x", sha256.Sum256(allDTypes))
	for _, value := range []string{
		"implementation=upstream-jaccl rank=" + strconv.Itoa(rank) + " size=" + strconv.Itoa(size),
		"payload_sha256=" + payloadHash,
		"all_gather_sha256=" + gatherHash,
		"all_reduce_int32_sum_sha256=" + reduceSumHash,
		"all_reduce_int32_max_sha256=" + reduceMaxHash,
		"all_reduce_int32_min_sha256=" + reduceMinHash,
		"all_reduce_float32_sum_sha256=" + floatSumHash,
		"all_reduce_float32_max_sha256=" + floatMaxHash,
		"all_reduce_float32_min_sha256=" + floatMinHash,
		"all_reduce_all_dtypes_sha256=" + allDTypesHash,
	} {
		if !strings.Contains(string(output), value) {
			t.Fatalf("standalone JACCL receipt missing %q: %s", value, output)
		}
	}
	binaryHash, err := hardwareFileHash(binaryPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("receipt upstream_binary=%s upstream_binary_sha256=%s output=%s", binaryPath, binaryHash, strings.TrimSpace(string(output)))
}

var hardwareDTypes = []DType{
	Bool,
	Int8,
	Int16,
	Int32,
	Int64,
	UInt8,
	UInt16,
	UInt32,
	UInt64,
	Float16,
	BFloat16,
	Float32,
	Float64,
	Complex64,
}

func TestHardwareAllDTypesExpected(t *testing.T) {
	for _, size := range []int{2, 4, 8} {
		receipt, err := hardwareAllDTypesExpected(size)
		if err != nil {
			t.Fatalf("size %d: %v", size, err)
		}
		var want int
		for _, dtype := range hardwareDTypes {
			width, err := dtype.Size()
			if err != nil {
				t.Fatalf("size %d dtype %d: %v", size, dtype, err)
			}
			want += 3 * width
		}
		if len(receipt) != want {
			t.Fatalf("size %d receipt length = %d, want %d", size, len(receipt), want)
		}
	}
}

func hardwareAllDTypes(ctx context.Context, group *Group, rank, size int) ([]byte, error) {
	want, err := hardwareAllDTypesExpected(size)
	if err != nil {
		return nil, err
	}
	receipt := make([]byte, 0, len(want))
	for _, dtype := range hardwareDTypes {
		input, err := hardwareDTypeBytes(dtype, rank+1)
		if err != nil {
			return nil, err
		}
		for _, test := range []struct {
			name string
			call func([]byte) error
		}{
			{"sum", func(dst []byte) error { return group.AllSum(ctx, dst, input, dtype) }},
			{"max", func(dst []byte) error { return group.AllMax(ctx, dst, input, dtype) }},
			{"min", func(dst []byte) error { return group.AllMin(ctx, dst, input, dtype) }},
		} {
			got := make([]byte, len(input))
			if err := test.call(got); err != nil {
				return nil, fmt.Errorf("all-dtype %d %s: %w", dtype, test.name, err)
			}
			receipt = append(receipt, got...)
		}
	}
	if !bytes.Equal(receipt, want) {
		return nil, fmt.Errorf("all-dtype receipt differs")
	}
	return receipt, nil
}

func hardwareAllDTypesExpected(size int) ([]byte, error) {
	sum := size * (size + 1) / 2
	receipt := make([]byte, 0, 3*(1+1+2+4+8+1+2+4+8+2+2+4+8+8))
	for _, dtype := range hardwareDTypes {
		for _, value := range []int{sum, size, 1} {
			data, err := hardwareDTypeBytes(dtype, value)
			if err != nil {
				return nil, err
			}
			receipt = append(receipt, data...)
		}
	}
	return receipt, nil
}

func hardwareDTypeBytes(dtype DType, value int) ([]byte, error) {
	switch dtype {
	case Bool:
		return []byte{1}, nil
	case Int8:
		return []byte{byte(int8(value))}, nil
	case Int16:
		return int16Bytes(int16(value)), nil
	case Int32:
		return int32Bytes(int32(value)), nil
	case Int64:
		return int64Bytes(int64(value)), nil
	case UInt8:
		return []byte{byte(value)}, nil
	case UInt16:
		return uint16Bytes(uint16(value)), nil
	case UInt32:
		return uint32Bytes(uint32(value)), nil
	case UInt64:
		return uint64Bytes(uint64(value)), nil
	case Float16:
		return uint16Bytes(float32ToHalf(float32(value))), nil
	case BFloat16:
		return uint16Bytes(float32ToBFloat16(float32(value))), nil
	case Float32:
		return float32Bytes(float32(value)), nil
	case Float64:
		return float64Bytes(float64(value)), nil
	case Complex64:
		return complex64Bytes(complex(float32(value), 0)), nil
	default:
		return nil, fmt.Errorf("hardware dtype %d", dtype)
	}
}

func hardwareBinaryHash() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("locate test binary: %w", err)
	}
	return hardwareFileHash(path)
}

func hardwareFileHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read binary %s: %w", path, err)
	}
	return fmt.Sprintf("%x", sha256.Sum256(data)), nil
}

func hardwareEnvInt(name string) (int, error) {
	value := os.Getenv(name)
	if value == "" {
		return 0, fmt.Errorf("set %s", name)
	}
	return strconv.Atoi(value)
}

func hardwareEnvIntDefault(name string, fallback int) (int, error) {
	if os.Getenv(name) == "" {
		return fallback, nil
	}
	return hardwareEnvInt(name)
}

// hardwareSafetyCheck rejects known single-host layouts before either hardware
// harness opens an RDMA device. A timeout cannot interrupt an IOKit call stuck
// in kernel state, so this check protects the known destructive configuration.
func hardwareSafetyCheck(coordinator, devicesPath string) error {
	if filepath.Base(devicesPath) == "devices-local.json" {
		return fmt.Errorf("refuse hardware run with loopback devices-local.json")
	}
	host, _, err := net.SplitHostPort(coordinator)
	if err != nil {
		return nil
	}
	if host == "localhost" || net.ParseIP(host).IsLoopback() {
		return fmt.Errorf("refuse hardware run with loopback coordinator %q", coordinator)
	}
	return nil
}

func TestHardwareSafetyCheck(t *testing.T) {
	tests := []struct {
		name, coordinator, devicesPath string
		wantErr                        bool
	}{
		{"two host", "192.168.0.1:34807", "devices.json", false},
		{"ipv4 loopback", "127.0.0.1:34807", "devices.json", true},
		{"ipv6 loopback", "[::1]:34807", "devices.json", true},
		{"localhost", "localhost:34807", "devices.json", true},
		{"local device fixture", "192.168.0.1:34807", "devices-local.json", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := hardwareSafetyCheck(test.coordinator, test.devicesPath)
			if (err != nil) != test.wantErr {
				t.Fatalf("hardwareSafetyCheck(%q, %q) = %v, want error %t", test.coordinator, test.devicesPath, err, test.wantErr)
			}
		})
	}
}
