package oslog

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"

	"github.com/tmc/apple/x/internal/oslogabi"
)

// Type is an os_log_type_t: the level a message is logged at.
type Type uint8

const (
	TypeDefault Type = 0x00 // OS_LOG_TYPE_DEFAULT
	TypeInfo    Type = 0x01 // OS_LOG_TYPE_INFO
	TypeDebug   Type = 0x02 // OS_LOG_TYPE_DEBUG
	TypeError   Type = 0x10 // OS_LOG_TYPE_ERROR
	TypeFault   Type = 0x11 // OS_LOG_TYPE_FAULT
)

var (
	symsOnce sync.Once

	osLogCreate func(subsystem, category *byte) uintptr
	osLogImpl   func(dso unsafe.Pointer, log uintptr, typ uint8, format, buf *byte, size uint32)
	osLogEnab   func(log uintptr, typ uint8) bool
)

func loadSymbols() {
	symsOnce.Do(func() {
		if s, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_log_create"); err == nil && s != 0 {
			purego.RegisterFunc(&osLogCreate, s)
		}
		if s, err := purego.Dlsym(purego.RTLD_DEFAULT, "_os_log_impl"); err == nil && s != 0 {
			purego.RegisterFunc(&osLogImpl, s)
		}
		if s, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_log_type_enabled"); err == nil && s != 0 {
			purego.RegisterFunc(&osLogEnab, s)
		}
	})
}

// Logger writes messages to a single os_log handle. Create one with [New]. It
// is safe for concurrent use. The zero value is not usable.
type Logger struct {
	handle uintptr
}

// New returns a Logger that writes under the given subsystem (reverse-DNS, e.g.
// "com.example.app") and category. New never returns nil; if the os_log symbols
// cannot be resolved the Logger's methods are no-ops.
func New(subsystem, category string) *Logger {
	loadSymbols()
	l := &Logger{}
	if osLogCreate != nil {
		sub := cstring(subsystem)
		cat := cstring(category)
		l.handle = osLogCreate(&sub[0], &cat[0])
		runtime.KeepAlive(sub)
		runtime.KeepAlive(cat)
	}
	return l
}

// Enabled reports whether messages of the given type are being recorded.
func (l *Logger) Enabled(t Type) bool {
	if l == nil || l.handle == 0 || osLogEnab == nil {
		return false
	}
	return osLogEnab(l.handle, uint8(t))
}

// Debug logs at OS_LOG_TYPE_DEBUG. Debug messages are the most verbose and are
// discarded unless explicitly enabled for the subsystem.
func (l *Logger) Debug(format string, args ...any) { l.Log(TypeDebug, format, args...) }

// Info logs at OS_LOG_TYPE_INFO.
func (l *Logger) Info(format string, args ...any) { l.Log(TypeInfo, format, args...) }

// Default logs at OS_LOG_TYPE_DEFAULT, the standard level.
func (l *Logger) Default(format string, args ...any) { l.Log(TypeDefault, format, args...) }

// Error logs at OS_LOG_TYPE_ERROR.
func (l *Logger) Error(format string, args ...any) { l.Log(TypeError, format, args...) }

// Fault logs at OS_LOG_TYPE_FAULT, for a bug in program execution.
func (l *Logger) Fault(format string, args ...any) { l.Log(TypeFault, format, args...) }

// Log writes a formatted message at the given type. The format string uses
// os_log specifiers, not Go's: %d/%u/%x (and l-prefixed 64-bit forms), %p, %s,
// and %@ (rendered as %s). Wrap a specifier as %{public}… or %{private}… to
// set its visibility; the default matches os_log (scalars public, strings
// private on release builds).
func (l *Logger) Log(t Type, format string, args ...any) {
	if l == nil || l.handle == 0 || osLogImpl == nil {
		return
	}
	if osLogEnab != nil && !osLogEnab(l.handle, uint8(t)) {
		return
	}
	buf, pins := oslogabi.Encode(format, args)
	cfmt := cstring(format)
	osLogImpl(oslogabi.DSOHandle(), l.handle, uint8(t), &cfmt[0], bufPtr(buf), uint32(len(buf)))
	runtime.KeepAlive(cfmt)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(pins)
}

func bufPtr(b []byte) *byte {
	if len(b) == 0 {
		return &emptyByte
	}
	return &b[0]
}

var emptyByte byte

// cstring returns a NUL-terminated copy of s. Callers keep the slice alive.
func cstring(s string) []byte {
	b := make([]byte, len(s)+1)
	copy(b, s)
	return b
}
