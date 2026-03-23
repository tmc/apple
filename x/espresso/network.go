//go:build darwin

package espresso

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"sort"
	"strings"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/espresso"
)

// Network is a loaded Espresso neural network ready for execution.
type Network struct {
	net      espresso.EspressoNetwork
	ctx      *Context
	binder   espresso.EspressoDataFrameExecutor
	storExec espresso.EspressoDataFrameStorageExecutor
	irJSON   []byte // original IR JSON for introspection

	mu     sync.Mutex
	closed bool
}

type networkConfig struct {
	computePath     int
	binSerializerID []byte
}

// NetworkOption configures network loading.
type NetworkOption func(*networkConfig)

// WithComputePath sets the compute path for network loading.
func WithComputePath(path int) NetworkOption {
	return func(c *networkConfig) {
		c.computePath = path
	}
}

// WithBinSerializerID sets the binary serializer ID for weight loading.
func WithBinSerializerID(id []byte) NetworkOption {
	return func(c *networkConfig) {
		c.binSerializerID = id
	}
}

// LoadNetwork loads a network from Espresso IR (JSON bytes) with optional binary weights.
func (c *Context) LoadNetwork(ir []byte, opts ...NetworkOption) (*Network, error) {
	if c.isClosed() {
		return nil, ErrClosed
	}

	var cfg networkConfig
	for _, o := range opts {
		o(&cfg)
	}

	// The generated binding passes []byte as const char * via
	// unsafe.Pointer(unsafe.SliceData(...)). The C++ parser expects a
	// null-terminated string for the JSON IR. Go slices are not
	// null-terminated, so we append a zero byte.
	irZ := append(ir, 0)

	var net espresso.EspressoNetwork
	if cfg.binSerializerID != nil {
		net = espresso.NewEspressoNetworkWithJSFileBinSerializerIdContextComputePath(
			irZ, cfg.binSerializerID, c.ctx, cfg.computePath,
		)
	} else {
		net = espresso.NewEspressoNetworkWithJSFileContextComputePath(
			irZ, c.ctx, cfg.computePath,
		)
	}

	if net.GetID() == 0 {
		return nil, ErrNilNetwork
	}

	// The constructor can return a non-nil ObjC ID even when IR parsing
	// fails internally, leaving members uninitialized. Verify the network
	// populated its internal state before exposing it to callers.
	if !networkInitialized(net) {
		return nil, &Error{Op: "load", Err: fmt.Errorf("network object created but internal state is nil (IR parse may have failed)")}
	}

	binder := espresso.NewEspressoDataFrameExecutor()
	storExec := espresso.NewEspressoDataFrameStorageExecutor()

	n := &Network{
		net:      net,
		ctx:      c,
		binder:   binder,
		storExec: storExec,
		irJSON:   ir,
	}
	runtime.SetFinalizer(n, (*Network).Close)
	return n, nil
}

// LoadNetworkFromFile loads a network from an Espresso net.json file.
// If no WithBinSerializerID option is given, it searches for sidecar weight files.
func (c *Context) LoadNetworkFromFile(path string, opts ...NetworkOption) (*Network, error) {
	ir, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("espresso load: %w", err)
	}

	// Auto-resolve weight file if no explicit bin serializer ID.
	var hasBinSerializer bool
	for _, o := range opts {
		var cfg networkConfig
		o(&cfg)
		if cfg.binSerializerID != nil {
			hasBinSerializer = true
			break
		}
	}
	if !hasBinSerializer {
		if wf := resolveWeightFile(path); wf != "" {
			weights, err := os.ReadFile(wf)
			if err == nil {
				opts = append(opts, WithBinSerializerID(weights))
			}
		}
	}

	return c.LoadNetwork(ir, opts...)
}

// resolveWeightFile searches for sidecar .bin files next to a net.json file.
// Search order: stem.weights.bin, weights.bin, stem.bin, first .bin in directory.
func resolveWeightFile(jsonPath string) string {
	dir := filepath.Dir(jsonPath)
	stem := strings.TrimSuffix(filepath.Base(jsonPath), filepath.Ext(jsonPath))

	candidates := []string{
		filepath.Join(dir, stem+".weights.bin"),
		filepath.Join(dir, "weights.bin"),
		filepath.Join(dir, stem+".bin"),
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}

	// Fallback: first .bin in same directory.
	entries, err := os.ReadDir(dir)
	if err != nil {
		return ""
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".bin") {
			return filepath.Join(dir, e.Name())
		}
	}
	return ""
}

// LayerCount returns the number of layers in the network.
func (n *Network) LayerCount() int {
	return int(n.net.Layers_size())
}

// Close releases network resources.
func (n *Network) Close() error {
	n.mu.Lock()
	defer n.mu.Unlock()
	if n.closed {
		return nil
	}
	n.closed = true
	n.binder.FreeTemporaryResources()
	n.net.Wipe_layers_blobs()
	runtime.SetFinalizer(n, nil)
	return nil
}

// LayerNames returns the names of all layers in topological order.
// Falls back to sorted keys from the IR JSON if available.
func (n *Network) LayerNames() []string {
	if len(n.irJSON) == 0 {
		return nil
	}
	var ir struct {
		Layers map[string]json.RawMessage `json:"layers"`
	}
	if err := json.Unmarshal(n.irJSON, &ir); err != nil {
		return nil
	}
	names := make([]string, 0, len(ir.Layers))
	for k := range ir.Layers {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

// HasLayer reports whether the network contains a layer with the given name.
func (n *Network) HasLayer(name string) bool {
	return slices.Contains(n.LayerNames(), name)
}

// networkInitialized checks that an EspressoNetwork has populated internal
// state after construction. The ObjC init can succeed (non-nil ID) but
// leave members uninitialized when IR parsing fails silently — accessing
// those members via objc_msgSend dereferences nil and causes a fatal
// SIGSEGV (not recoverable in Go).
//
// EspressoNetwork has a single ivar: _net (shared_ptr<Espresso::net>) at
// offset 8. The shared_ptr layout is {__ptr_, __cntrl_} — two pointers.
// If init fails, __ptr_ (the first pointer in the shared_ptr, at the ivar
// offset) will be nil. We read it directly via the ObjC runtime API,
// avoiding objc_msgSend entirely.
func networkInitialized(net espresso.EspressoNetwork) bool {
	id := net.GetID()
	if id == 0 {
		return false
	}
	obj := objectivec.Object{ID: id}
	cls := objectivec.Object_getClass(obj)
	if uintptr(cls) == 0 {
		return false
	}

	// Look up the _net ivar (shared_ptr<Espresso::net>) on EspressoNetwork.
	// This reads class metadata, not instance memory, so it's always safe.
	ivar := objectivec.Class_getInstanceVariable(cls, "_net")
	if ivar == 0 {
		// Ivar not found — can't validate; assume OK to avoid
		// breaking on future framework versions that rename the ivar.
		return true
	}

	// Read the first pointer in the shared_ptr (__ptr_) at the ivar's offset.
	// This is a direct memory read from alloc'd memory — no objc_msgSend.
	offset := objectivec.Ivar_getOffset(ivar)
	ptr := *(*uintptr)(unsafe.Add(unsafe.Pointer(id), offset))
	return ptr != 0
}
