//go:build darwin

package irecovery

import (
	"context"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
)

// Device identifies an Apple USB DFU or recovery endpoint. Unknown serial fields
// remain zero; Open refuses to select an endpoint without an exact nonzero ECID.
type Device struct {
	ECID      uint64
	CPID      uint32
	ProductID uint16
	Serial    string
	Mode      string
}

type library struct {
	image      uintptr
	ctx        uintptr
	init       func(*uintptr) int32
	exit       func(uintptr)
	list       func(uintptr, *uintptr) int64
	freeList   func(uintptr, int32)
	descriptor func(uintptr, *[18]byte) int32
	open       func(uintptr, *uintptr) int32
	close      func(uintptr)
	serial     func(uintptr, uint8, *byte, int32) int32
	control    func(uintptr, uint8, uint8, uint16, uint16, *byte, uint16, uint32) int32
	bulk       func(uintptr, uint8, *byte, int32, *int32, uint32) int32
	claim      func(uintptr, int32) int32
	release    func(uintptr, int32) int32
}

func load(path string) (l *library, err error) {
	if path == "" {
		return nil, fmt.Errorf("libusb library path is required")
	}
	image, err := purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_LOCAL)
	if err != nil {
		return nil, fmt.Errorf("load libusb: %w", err)
	}
	l = &library{image: image}
	defer func() {
		if err != nil {
			purego.Dlclose(image)
		}
	}()
	for _, s := range []struct {
		name string
		fn   any
	}{
		{"libusb_init", &l.init}, {"libusb_exit", &l.exit}, {"libusb_get_device_list", &l.list}, {"libusb_free_device_list", &l.freeList},
		{"libusb_get_device_descriptor", &l.descriptor}, {"libusb_open", &l.open}, {"libusb_close", &l.close}, {"libusb_get_string_descriptor_ascii", &l.serial},
		{"libusb_control_transfer", &l.control}, {"libusb_bulk_transfer", &l.bulk}, {"libusb_claim_interface", &l.claim}, {"libusb_release_interface", &l.release},
	} {
		var p uintptr
		p, err = purego.Dlsym(image, s.name)
		if err != nil {
			return nil, fmt.Errorf("resolve %s: %w", s.name, err)
		}
		purego.RegisterFunc(s.fn, p)
	}
	if code := l.init(&l.ctx); code != 0 {
		return nil, fmt.Errorf("initialize libusb: %d", code)
	}
	return l, nil
}
func (l *library) dispose() { l.exit(l.ctx); purego.Dlclose(l.image) }

// Discover lists readable Apple DFU/recovery endpoints. A candidate that cannot
// be opened or identified causes an error instead of silently disappearing.
func Discover(ctx context.Context, path string) ([]Device, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	l, err := load(path)
	if err != nil {
		return nil, err
	}
	defer l.dispose()
	var result []Device
	err = l.each(ctx, func(d Device, h uintptr) (bool, error) { result = append(result, d); return false, nil })
	return result, err
}
func (l *library) each(ctx context.Context, visit func(Device, uintptr) (bool, error)) error {
	var list uintptr
	count := l.list(l.ctx, &list)
	if count < 0 {
		return fmt.Errorf("list USB devices: %d", count)
	}
	defer l.freeList(list, 1)
	for i := int64(0); i < count; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		dev := *(*uintptr)(unsafe.Pointer(list + uintptr(i)*unsafe.Sizeof(uintptr(0))))
		var desc [18]byte
		if code := l.descriptor(dev, &desc); code != 0 {
			return fmt.Errorf("read USB descriptor: %d", code)
		}
		if binary.LittleEndian.Uint16(desc[8:]) != 0x05ac {
			continue
		}
		pid := binary.LittleEndian.Uint16(desc[10:])
		mode := ""
		if pid == 0x1227 {
			mode = "dfu"
		} else if pid >= 0x1280 && pid <= 0x1283 {
			mode = "recovery"
		} else {
			continue
		}
		var h uintptr
		if code := l.open(dev, &h); code != 0 {
			return fmt.Errorf("open Apple USB %04x: %d", pid, code)
		}
		var b [1024]byte
		n := l.serial(h, desc[16], &b[0], int32(len(b)))
		if n < 0 {
			l.close(h)
			return fmt.Errorf("read Apple USB serial: %d", n)
		}
		d, err := parseSerial(string(b[:n]))
		if err != nil {
			l.close(h)
			return err
		}
		d.ProductID = pid
		d.Mode = mode
		if err := ctx.Err(); err != nil {
			l.close(h)
			return err
		}
		keep, err := visit(d, h)
		if !keep {
			l.close(h)
		}
		if err != nil {
			return err
		}
		if keep {
			return nil
		}
	}
	return nil
}

// Conn owns a libusb context and a claimed interface. Its zero value is closed.
// Close waits for any in-flight transfer; callers must close every successful Open.
type Conn struct {
	mu      sync.Mutex
	library *library
	handle  uintptr
	info    Device
}

// Open claims interface zero of the endpoint with the exact requested ECID.
// It never detaches a kernel driver automatically.
func Open(ctx context.Context, path string, ecid uint64) (*Conn, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if ecid == 0 {
		return nil, fmt.Errorf("nonzero ECID is required")
	}
	l, err := load(path)
	if err != nil {
		return nil, err
	}
	var conn *Conn
	err = l.each(ctx, func(d Device, h uintptr) (bool, error) {
		if d.ECID != ecid {
			return false, nil
		}
		if code := l.claim(h, 0); code != 0 {
			return false, fmt.Errorf("claim recovery interface: %d", code)
		}
		conn = &Conn{library: l, handle: h, info: d}
		return true, nil
	})
	if err != nil || conn == nil {
		l.dispose()
		if err == nil {
			err = fmt.Errorf("recovery endpoint for ECID %x not found", ecid)
		}
		return nil, err
	}
	return conn, nil
}

// Close releases the interface, device handle and libusb context. It is idempotent.
func (c *Conn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.library == nil {
		return nil
	}
	code := c.library.release(c.handle, 0)
	c.library.close(c.handle)
	c.library.dispose()
	c.library = nil
	c.handle = 0
	if code != 0 {
		return fmt.Errorf("release recovery interface: %d", code)
	}
	return nil
}

// Control performs one USB control transfer. For an IN request, data receives the
// bytes. For an OUT request, data is transmitted. Errors do not imply no mutation.
func (c *Conn) Control(ctx context.Context, kind, request uint8, value, index uint16, data []byte) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	timeout, err := transferTimeout(ctx)
	if err != nil {
		return 0, err
	}
	if c.library == nil {
		return 0, fmt.Errorf("recovery connection is closed")
	}
	if len(data) > 65535 {
		return 0, fmt.Errorf("USB control payload exceeds 65535 bytes")
	}
	n := c.library.control(c.handle, kind, request, value, index, unsafe.SliceData(data), uint16(len(data)), timeout)
	if n < 0 {
		if ctx.Err() != nil {
			return 0, ctx.Err()
		}
		return 0, fmt.Errorf("USB control transfer: %d", n)
	}
	return int(n), nil
}

// BulkWrite performs one bulk OUT transfer, returning any partial byte count.
func (c *Conn) BulkWrite(ctx context.Context, endpoint uint8, data []byte) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	timeout, err := transferTimeout(ctx)
	if err != nil {
		return 0, err
	}
	if c.library == nil {
		return 0, fmt.Errorf("recovery connection is closed")
	}
	if endpoint == 0 || endpoint&0x80 != 0 || len(data) > 1<<20 {
		return 0, fmt.Errorf("invalid bulk OUT endpoint or payload")
	}
	var n int32
	code := c.library.bulk(c.handle, endpoint, unsafe.SliceData(data), int32(len(data)), &n, timeout)
	if code != 0 {
		if ctx.Err() != nil {
			return int(n), ctx.Err()
		}
		return int(n), fmt.Errorf("USB bulk transfer: %d", code)
	}
	return int(n), nil
}
func transferTimeout(ctx context.Context) (uint32, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	d := time.Second
	if deadline, ok := ctx.Deadline(); ok {
		d = min(d, time.Until(deadline))
	}
	if d <= 0 {
		return 0, context.DeadlineExceeded
	}
	return uint32(max(1, (d+time.Millisecond-1)/time.Millisecond)), nil
}
func parseSerial(serial string) (Device, error) {
	d := Device{Serial: serial}
	seen := map[string]bool{}
	for _, field := range strings.Fields(serial) {
		key, value, ok := strings.Cut(field, ":")
		if !ok || (key != "ECID" && key != "CPID") {
			continue
		}
		if seen[key] {
			return d, fmt.Errorf("duplicate USB serial field %s", key)
		}
		seen[key] = true
		bits := 64
		if key == "CPID" {
			bits = 32
		}
		n, err := strconv.ParseUint(value, 16, bits)
		if err != nil {
			return d, fmt.Errorf("parse USB %s: %w", key, err)
		}
		if key == "ECID" {
			d.ECID = n
		} else {
			d.CPID = uint32(n)
		}
	}
	return d, nil
}
