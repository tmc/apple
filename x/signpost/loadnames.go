package signpost

import (
	"debug/macho"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"

	"github.com/tmc/apple/x/internal/oslogabi"
)

// namesAnchorSymbol is the exported symbol signpostnames -dylib places at the
// start of the __oslogstring section; its runtime address is the base for
// every string offset in the image.
const namesAnchorSymbol = "signpost_names"

// namesImage is a loaded Mach-O image holding pooled signpost strings. Every
// emit resolves its name and format against a single image, whose header is
// passed as the dso argument.
type namesImage struct {
	dso  unsafe.Pointer
	pool map[string]*byte
}

var (
	imagesMu sync.Mutex
	images   []*namesImage
)

// LoadNames loads a names dylib built by signpostnames -dylib and makes its
// strings available for signpost decoding. It is the opt-in path for builds
// whose linker cannot carry a __TEXT,__oslogstring section (internal linking,
// CGO_ENABLED=0); externally linked builds get the same effect from a .syso
// with no runtime step. Names are resolved against the executable's own
// section first, then against loaded images in LoadNames order.
func LoadNames(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	f, err := macho.Open(abs)
	if err != nil {
		return fmt.Errorf("load signpost names %s: %w", path, err)
	}
	sect := f.Section("__oslogstring")
	if sect == nil {
		f.Close()
		return fmt.Errorf("load signpost names %s: no __TEXT,__oslogstring section", path)
	}
	data, err := sect.Data()
	f.Close()
	if err != nil {
		return fmt.Errorf("load signpost names %s: %w", path, err)
	}

	h, err := purego.Dlopen(abs, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("load signpost names %s: %w", path, err)
	}
	anchor, err := purego.Dlsym(h, namesAnchorSymbol)
	if err != nil {
		return fmt.Errorf("load signpost names %s: no %s anchor (rebuild with signpostnames -dylib): %w", path, namesAnchorSymbol, err)
	}
	var info dlInfo
	if dladdrFunc()(anchor, &info) == 0 || info.fbase == nil {
		return fmt.Errorf("load signpost names %s: dladdr failed", path)
	}

	// The anchor sits at the section start, so the section's file bytes
	// mirror memory at the anchor and each string's offset is its pointer.
	// dladdr's saddr is the anchor as a live pointer, which keeps the
	// arithmetic off uintptr.
	base := info.saddr
	pool := make(map[string]*byte, 8)
	for off := 0; off < len(data); {
		end := off
		for end < len(data) && data[end] != 0 {
			end++
		}
		if end == len(data) {
			break
		}
		name := string(data[off:end])
		if _, ok := pool[name]; !ok {
			pool[name] = (*byte)(unsafe.Add(base, off))
		}
		off = end + 1
	}

	imagesMu.Lock()
	images = append(images, &namesImage{dso: info.fbase, pool: pool})
	imagesMu.Unlock()
	return nil
}

// lookup finds in-image pointers for name and format, preferring a single
// image holding both (the executable's __oslogstring section first, then
// LoadNames images). The two strings must resolve against the same dso, so
// when no image has both, an image holding just the format is still used:
// the format decodes and the heap-passed name renders as "<missing name>",
// which beats losing the message too. nameOK reports whether the name will
// decode; nil pointers mean the caller must pass heap copies.
func lookup(name, format string) (namePtr, fmtPtr *byte, dso unsafe.Pointer, nameOK bool) {
	type candidate struct {
		pool map[string]*byte
		dso  unsafe.Pointer
	}
	var cands []candidate
	if p := namePool(); p != nil {
		cands = append(cands, candidate{p, dsoHandle()})
	}
	imagesMu.Lock()
	for _, img := range images {
		cands = append(cands, candidate{img.pool, img.dso})
	}
	imagesMu.Unlock()
	for _, c := range cands {
		if np, fp := c.pool[name], c.pool[format]; np != nil && fp != nil {
			return np, fp, c.dso, true
		}
	}
	for _, c := range cands {
		if fp := c.pool[format]; fp != nil {
			return nil, fp, c.dso, false
		}
	}
	return nil, nil, nil, false
}

var warnOnce sync.Once

// warnUnpooled reports, once per process, the first signpost name that will
// emit and pair but not decode. Loud beats silent: the failure otherwise
// surfaces only as "<missing name>" in trace output nobody is watching.
func warnUnpooled(name string) {
	warnOnce.Do(func() {
		fmt.Fprintf(os.Stderr, "signpost: name %q is not in the binary's __oslogstring pool; "+
			"it will emit and pair but decode as \"<missing name>\". "+
			"Generate a pool with signpostnames (external linking), or "+
			"signpostnames -dylib plus signpost.LoadNames for CGO_ENABLED=0 builds.\n", name)
	})
}

// dsoHandle returns the main executable's image header; see
// oslogabi.DSOHandle.
func dsoHandle() unsafe.Pointer { return oslogabi.DSOHandle() }

// dlInfo mirrors Dl_info from <dlfcn.h>.
type dlInfo struct {
	fname *byte
	fbase unsafe.Pointer
	sname *byte
	saddr unsafe.Pointer
}

var (
	dladdrOnce sync.Once
	dladdr     func(addr uintptr, info *dlInfo) int32
)

func dladdrFunc() func(addr uintptr, info *dlInfo) int32 {
	dladdrOnce.Do(func() {
		if sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "dladdr"); err == nil && sym != 0 {
			purego.RegisterFunc(&dladdr, sym)
		}
	})
	if dladdr == nil {
		return func(uintptr, *dlInfo) int32 { return 0 }
	}
	return dladdr
}
