package rdma

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"unsafe"

	"github.com/ebitengine/purego"
)

func TestRDMACall3ABI(t *testing.T) {
	const word = unsafe.Sizeof(uintptr(0))
	if got, want := unsafe.Sizeof(rdmaCall3Args{}), 25*word; got != want {
		t.Fatalf("rdmaCall3Args size = %d, want %d", got, want)
	}
	var args rdmaCall3Args
	for _, tt := range []struct {
		name string
		got  uintptr
		want uintptr
	}{
		{"a1", unsafe.Offsetof(args.a1) / word, 1},
		{"f1", unsafe.Offsetof(args.f1) / word, 16},
		{"arm64R8", unsafe.Offsetof(args.arm64R8) / word, 24},
	} {
		if tt.got != tt.want {
			t.Fatalf("rdmaCall3Args %s offset = %d words, want %d", tt.name, tt.got, tt.want)
		}
	}
	if rdmaSyscall15XABI0 == 0 {
		t.Fatal("purego syscall15X trampoline is unavailable")
	}
}

func TestRDMACall3DeepNest(t *testing.T) {
	roundtrip := buildDatapathTestSymbol(t, "roundtrip")
	store := buildDatapathTestSymbol(t, "store")

	const depth = 6
	var scratch int64
	var pin runtime.Pinner
	defer pin.Unpin()

	var nest func(int)
	nest = func(level int) {
		sentinel := int64(0x1000000000000000) | int64(level)
		cell := new(int64)
		*cell = sentinel
		pin.Pin(cell)

		cb := purego.NewCallback(func() {
			for i := int64(0); i < 8; i++ {
				rdmaCall3(store, uintptr(unsafe.Pointer(&scratch)), uintptr(int64(level)*1000+i+1), 0)
			}
			runtime.KeepAlive(&scratch)
			if level > 0 {
				nest(level - 1)
			}
		})

		got := rdmaCall3(roundtrip, cb, uintptr(unsafe.Pointer(cell)), 0)
		runtime.KeepAlive(cell)
		if int64(got) != sentinel {
			t.Fatalf("level %d: Call3 read %#x, want %#x", level, uint64(got), uint64(sentinel))
		}
	}
	nest(depth)
}

func TestDatapathWrappersCall3(t *testing.T) {
	fn := buildDatapathTestFunction(t)

	var wc IbvWC
	poller := IbvCQPoller{cq: 7, fnPtr: fn}
	if got := poller.Poll(3, &wc); got != 10 {
		t.Fatalf("Poll = %d, want 10", got)
	}
	if got := wc.WRID; got != 10 {
		t.Fatalf("Poll WRID = %d, want 10", got)
	}

	var send IbvSendWR
	var badSend *IbvSendWR
	poster := IbvQPPoster{qp: 7, sendPtr: fn, recvPtr: fn}
	if got := poster.PostSend(&send, &badSend); got != 7 {
		t.Fatalf("PostSend = %d, want 7", got)
	}
	if badSend != &send {
		t.Fatalf("PostSend bad WR = %p, want %p", badSend, &send)
	}

	var recv IbvRecvWR
	var badRecv *IbvRecvWR
	if got := poster.PostRecv(&recv, &badRecv); got != 7 {
		t.Fatalf("PostRecv = %d, want 7", got)
	}
	if badRecv != (*IbvRecvWR)(unsafe.Pointer(&recv)) {
		t.Fatalf("PostRecv bad WR = %p, want %p", badRecv, &recv)
	}
}

func TestDatapathWrappersCall3Allocs(t *testing.T) {
	fn := buildDatapathTestFunction(t)
	var wc IbvWC
	poller := IbvCQPoller{cq: 7, fnPtr: fn}
	if allocs := testing.AllocsPerRun(1000, func() { _ = poller.Poll(3, &wc) }); allocs != 0 {
		t.Fatalf("Poll allocs = %v, want 0", allocs)
	}
}

func buildDatapathTestFunction(tb testing.TB) uintptr {
	return buildDatapathTestSymbol(tb, "datapath")
}

func buildDatapathTestSymbol(tb testing.TB, symbol string) uintptr {
	tb.Helper()
	dir := tb.TempDir()
	src := filepath.Join(dir, "datapath.c")
	lib := filepath.Join(dir, "libdatapath.dylib")
	const code = `
#include <stdint.h>
int datapath(uintptr_t handle, uintptr_t n, void *out) {
	if (n > 16) {
		*(void **)out = (void *)n;
		return (int)handle;
	}
  uint64_t value = handle + n;
  *(uint64_t *)out = value;
  return (int)value;
}
typedef void (*cb_t)(void);
int64_t roundtrip(cb_t cb, int64_t *p, uintptr_t ignored) {
  int64_t saved = *p;
  cb();
  (void)saved;
  return *p;
}
void store(int64_t *p, int64_t v, uintptr_t ignored) { *p = v; }
`
	if err := os.WriteFile(src, []byte(code), 0o644); err != nil {
		tb.Fatal(err)
	}
	if out, err := exec.Command("cc", "-dynamiclib", "-o", lib, src).CombinedOutput(); err != nil {
		tb.Fatalf("cc: %v\n%s", err, out)
	}
	h, err := purego.Dlopen(lib, purego.RTLD_NOW|purego.RTLD_LOCAL)
	if err != nil {
		tb.Fatalf("Dlopen: %v", err)
	}
	tb.Cleanup(func() { _ = purego.Dlclose(h) })
	fn, err := purego.Dlsym(h, symbol)
	if err != nil {
		tb.Fatalf("Dlsym(%s): %v", symbol, err)
	}
	runtime.KeepAlive(lib)
	return fn
}
