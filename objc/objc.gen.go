// Code generated from internal/generator/templates/runtime/objc.txtar by applegen. DO NOT EDIT.

// Package objc provides cached Objective-C runtime helpers.
//
// This package wraps purego/objc to provide selector caching for better performance.
package objc

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	purego "github.com/ebitengine/purego"
	pobjc "github.com/ebitengine/purego/objc"
)

// Type aliases for convenience
type (
	ID        = pobjc.ID
	SEL       = pobjc.SEL
	Class     = pobjc.Class
	Block     = pobjc.Block
	Protocol  = pobjc.Protocol
	MethodDef = pobjc.MethodDef
	FieldDef  = pobjc.FieldDef
	IMP       = pobjc.IMP
)

// IDGetter is implemented by types that wrap an Objective-C object ID.
// This allows objc.Send to automatically extract the ID from wrapper types.
type IDGetter interface {
	GetID() ID
}

// CArrayArg marks a Go slice argument that should be passed to Objective-C as
// a pointer to contiguous C array storage.
//
// Use CArray at call sites for APIs that take `*T` and a paired `count`
// parameter.
type CArrayArg struct {
	value any
}

// CArray marks a slice argument for C-array pointer conversion in Send.
func CArray(v any) CArrayArg {
	return CArrayArg{value: v}
}

// IDFrom converts a raw pointer to an ID.
func IDFrom(ptr unsafe.Pointer) ID {
	return ID(uintptr(ptr))
}

// IDValueAt loads an Objective-C object ID stored at a symbol address.
//
// Dynamic libraries commonly export Objective-C object constants as pointers to
// storage containing the real object ID. Dlsym returns the storage address, so
// callers must load the pointer-sized value stored there.
func IDValueAt(addr uintptr) ID {
	if addr == 0 {
		return 0
	}
	return *(*ID)(unsafe.Pointer(addr))
}

// ValueAt loads a value of type T from a symbol address.
//
// Dlsym returns the address of a global in a loaded Mach-O image, so reading an
// exported constant means dereferencing that address at the constant's type.
// The memory is not Go-managed and does not move, which is what makes the
// uintptr round trip safe here — go vet cannot prove that, so it flags the
// conversion. Doing it once, here, keeps the generated packages from repeating
// an unsafe conversion several hundred times.
//
// The zero value is returned for a nil address.
func ValueAt[T any](addr uintptr) T {
	if addr == 0 {
		var zero T
		return zero
	}
	return *(*T)(unsafe.Pointer(addr))
}

// NewBlock creates an Objective-C block from a Go function.
// The Go function must take a Block as its first argument.
// Use Block.Release() to free the block when it is no longer in use.
//
// Method type encodings from method_getTypeEncoding render block arguments
// as bare "@?", with no inner signature. Extended signatures ("@?<...>")
// carrying parameter types come from protocol extended method types
// (_protocol_getMethodTypeEncoding) or block descriptors.
func NewBlock(fn any) Block {
	return pobjc.NewBlock(fn)
}

var nsErrorBlockSignature = []byte(`v@?@"NSError"` + "\x00")

type blockDescriptorWithSignature struct {
	_         uintptr
	size      uintptr
	_         uintptr
	dispose   uintptr
	signature *byte
}

type blockLayoutWithSignature struct {
	isa        uintptr
	flags      uint32
	_          uint32
	invoke     uintptr
	descriptor *blockDescriptorWithSignature
}

// SetBlockSignature sets the Objective-C runtime signature for block.
// The signature storage must outlive block.
//
// See NewBlock for details on standard ("@?") vs extended ("@?<...>")
// block type encodings.
func SetBlockSignature(block Block, signature []byte) bool {
	if block == 0 || len(signature) == 0 {
		return false
	}
	layout := (*blockLayoutWithSignature)(unsafe.Pointer(uintptr(block)))
	if layout.descriptor == nil {
		return false
	}
	layout.descriptor.signature = &signature[0]
	return true
}

// SetNSErrorBlockSignature sets block's signature to a single NSError argument.
func SetNSErrorBlockSignature(block Block) bool {
	return SetBlockSignature(block, nsErrorBlockSignature)
}

// GoString converts a C string (*byte from UTF8String) to a Go string.
// This is needed because UTF8String returns const char*, not an ObjC object.
func GoString(cstr *byte) string {
	if cstr == nil {
		return ""
	}
	// Find the length by scanning for null terminator
	ptr := unsafe.Pointer(cstr)
	length := 0
	for *(*byte)(unsafe.Add(ptr, length)) != 0 {
		length++
	}
	// Convert to Go string
	return string(unsafe.Slice(cstr, length))
}

// BytesPointer returns a C pointer to b's backing array, or nil when b is
// empty.
//
// unsafe.SliceData returns a non-nil pointer for an empty but non-nil slice,
// and for a nil slice it returns nil only by accident of the current
// implementation. Callees that take a byte buffer treat NULL as "no bytes";
// handing them a pointer to zero-sized storage instead makes them read past
// the end of it. Every generated call site that passes a []byte to
// Objective-C goes through here so the empty case is spelled NULL exactly
// once.
func BytesPointer(b []byte) unsafe.Pointer {
	if len(b) == 0 {
		return nil
	}
	return unsafe.Pointer(unsafe.SliceData(b))
}

// ErrInitFailed reports that an Objective-C initializer returned nil without
// filling in its NSError out-parameter. Wrapping that nil would hand the
// caller an object that only fails once it is used, far from the cause.
var ErrInitFailed = errors.New("objc: initializer returned nil")

// GoStringPtr converts a nullable C string to a heap-backed *string.
func GoStringPtr(cstr *byte) *string {
	if cstr == nil {
		return nil
	}
	s := GoString(cstr)
	return &s
}

var (
	selCache sync.Map // map[string]pobjc.SEL
)

// Sel returns a cached selector for the given name.
// This avoids the global lock in pobjc.RegisterName on repeated calls.
func Sel(name string) SEL {
	if sel, ok := selCache.Load(name); ok {
		return sel.(SEL)
	}
	sel := pobjc.RegisterName(name)
	selCache.Store(name, sel)
	return sel
}

// NSArrayToSlice converts an NSArray ID into a []ID by calling count and objectAtIndex:.
func NSArrayToSlice(array ID) []ID {
	if array == 0 {
		return nil
	}
	count := pobjc.Send[uint](array, RegisterName("count"))
	if count == 0 {
		return nil
	}
	sel := RegisterName("objectAtIndex:")
	result := make([]ID, count)
	for i := uint(0); i < count; i++ {
		result[i] = pobjc.Send[ID](array, sel, i)
	}
	return result
}

// ConvertSlice maps []ID to []T using a converter function.
func ConvertSlice[T any](ids []ID, convert func(ID) T) []T {
	if len(ids) == 0 {
		return nil
	}
	result := make([]T, len(ids))
	for i, id := range ids {
		result[i] = convert(id)
	}
	return result
}

// ConvertSliceToStrings maps []ID to []string via IDToString.
func ConvertSliceToStrings(ids []ID) []string {
	return ConvertSlice(ids, IDToString)
}

// Send calls objc_msgSend with the given arguments.
//
// When all arguments are uintptr-sized primitives (ID, SEL, Class, uintptr, bool,
// or integer types) and T is ID or struct{}, Send uses a pre-registered typed
// function instead of the general argument-processing path: 300ns and 4
// allocations against 520ns and 8 allocations under -tags objc_slowpath, so
// roughly 1.7x. The pre-registered functions still allocate; purego's
// RegisterFunc calls reflect.New on every invocation.
//
// Otherwise Send falls back to purego/objc with full argument processing:
// IDGetter extraction, nil→ID(0), CArrayArg conversion, and NSArray→[]ID.
//
// Send keeps values in args alive until the Objective-C call returns, and
// that reaches through unsafe.Pointer arguments: an unsafe.Pointer is a
// reference the garbage collector can see, so keeping one alive also keeps
// alive the storage it points into, including when it points into the
// interior of a larger object. A caller passing
// unsafe.Pointer(unsafe.SliceData(b)) — which is what BytesPointer returns —
// therefore does not need its own runtime.KeepAlive(b); the backing array
// survives the call. TestSendKeepsSliceAliveThroughDerivedPointer holds this
// property down, and fails if the KeepAlive below is removed.
//
// What Send cannot protect is storage behind a uintptr. A uintptr is a plain
// integer that the collector does not trace, so a caller must never convert a
// pointer to uintptr before handing it to Send; pass the unsafe.Pointer.
func Send[T any](id ID, sel SEL, args ...any) T {
	// Fast path: when T is ID or struct{} and all args are uintptr-castable,
	// use the pre-registered typed msgSendN functions directly.
	if objcMsgSendAddr != 0 && len(args) <= 8 {
		var zero T
		tType := reflect.TypeOf(&zero).Elem()
		tKind := tType.Kind()
		isVoidStruct := tKind == reflect.Struct && tType.Size() == 0
		if tKind == reflect.Uintptr || isVoidStruct {
			if uargs, ok := tryFastArgs(args); ok {
				rv := fastSend(id, sel, uargs)
				// Keep args alive through the msgSend call so GC does not
				// collect backing storage behind unsafe.Pointer arguments
				// (e.g. []byte → unsafe.SliceData, string → unsafe.StringData).
				runtime.KeepAlive(args)
				if isVoidStruct {
					// void return — return zero value of T (struct{}{})
					return zero
				}
				// T is uintptr-kinded (ID, SEL, Class, uintptr, …). Convert the
				// raw uintptr result to T by reinterpreting its bits; a plain
				// result.(T) assertion would panic because the boxed dynamic
				// type (ID) differs from named types like SEL or uintptr.
				return *(*T)(unsafe.Pointer(&rv))
			}
		}
	}

	// Slow path: full argument processing.
	selector := selName(sel)
	cArrayArgs := cArrayArgIndexes(selector, len(args))
	keepAlive := make([]any, 0, len(args))

	// Process args to extract IDs from object wrappers and handle nil
	for i, arg := range args {
		if arg == nil {
			// Convert nil to objc nil (ID 0)
			args[i] = ID(0)
		} else if carray, ok := arg.(CArrayArg); ok {
			if converted, holder, ok := toCArrayArg(carray.value); ok {
				args[i] = converted
				if holder != nil {
					keepAlive = append(keepAlive, holder)
				}
			} else {
				panic(fmt.Sprintf("objc.CArray: unsupported argument type %T", carray.value))
			}
		} else if _, ok := cArrayArgs[i]; ok {
			// Compatibility fallback for generated methods that still pass
			// count-paired slices directly without objc.CArray(...).
			if converted, holder, ok := toCArrayArg(arg); ok {
				args[i] = converted
				if holder != nil {
					keepAlive = append(keepAlive, holder)
				}
			}
		} else if getter, ok := arg.(IDGetter); ok {
			args[i] = getter.GetID()
		}
	}
	var zero T
	if reflect.TypeOf(&zero).Elem().Kind() == reflect.Slice {
		arrayID := pobjc.Send[ID](id, sel, args...)
		runtime.KeepAlive(keepAlive)
		var result any = NSArrayToSlice(arrayID)
		return result.(T)
	}
	ret := pobjc.Send[T](id, sel, args...)
	runtime.KeepAlive(keepAlive)
	return ret
}

// tryFastArgs attempts to convert all args to uintptr values.
// Returns the uintptr slice and true if all args are uintptr-castable,
// or nil, false if any arg requires the slow path.
func tryFastArgs(args []any) ([]uintptr, bool) {
	uargs := make([]uintptr, len(args))
	for i, arg := range args {
		switch v := arg.(type) {
		case ID:
			uargs[i] = uintptr(v)
		case SEL:
			uargs[i] = uintptr(v)
		case Class:
			uargs[i] = uintptr(v)
		case uintptr:
			uargs[i] = v
		case bool:
			if v {
				uargs[i] = 1
			}
		case int:
			uargs[i] = uintptr(v)
		case int8:
			uargs[i] = uintptr(v)
		case int16:
			uargs[i] = uintptr(v)
		case int32:
			uargs[i] = uintptr(v)
		case int64:
			uargs[i] = uintptr(v)
		case uint:
			uargs[i] = uintptr(v)
		case uint8:
			uargs[i] = uintptr(v)
		case uint16:
			uargs[i] = uintptr(v)
		case uint32:
			uargs[i] = uintptr(v)
		case uint64:
			uargs[i] = uintptr(v)
		case unsafe.Pointer:
			uargs[i] = uintptr(v)
		default:
			return nil, false
		}
	}
	return uargs, true
}

// fastSend dispatches to the appropriate pre-registered msgSendN function.
func fastSend(id ID, sel SEL, args []uintptr) uintptr {
	switch len(args) {
	case 0:
		return msgSend0(uintptr(id), uintptr(sel))
	case 1:
		return msgSend1(uintptr(id), uintptr(sel), args[0])
	case 2:
		return msgSend2(uintptr(id), uintptr(sel), args[0], args[1])
	case 3:
		return msgSend3(uintptr(id), uintptr(sel), args[0], args[1], args[2])
	case 4:
		return msgSend4(uintptr(id), uintptr(sel), args[0], args[1], args[2], args[3])
	case 5:
		return msgSend5(uintptr(id), uintptr(sel), args[0], args[1], args[2], args[3], args[4])
	case 6:
		return msgSend6(uintptr(id), uintptr(sel), args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		return msgSend7(uintptr(id), uintptr(sel), args[0], args[1], args[2], args[3], args[4], args[5], args[6])
	case 8:
		return msgSend8(uintptr(id), uintptr(sel), args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	default:
		all := make([]uintptr, 0, 2+len(args))
		all = append(all, uintptr(id), uintptr(sel))
		all = append(all, args...)
		r, _, _ := purego.SyscallN(objcMsgSendAddr, all...)
		return r
	}
}

// Pre-registered typed msgSend functions for the fast path. These are not
// allocation-free: purego's RegisterFunc closure calls reflect.New on every
// invocation, which profiles at roughly 60% of the fast path's allocated
// bytes. purego.SyscallN measures about 4x cheaper (96ns and 1 allocation
// against 386ns and 4 allocations at zero arguments) and is a candidate
// replacement, but it is untested under -race and with a Go callback on the
// stack.
var (
	msgSend0        func(id, sel uintptr) uintptr
	msgSend1        func(id, sel, a1 uintptr) uintptr
	msgSend2        func(id, sel, a1, a2 uintptr) uintptr
	msgSend3        func(id, sel, a1, a2, a3 uintptr) uintptr
	msgSend4        func(id, sel, a1, a2, a3, a4 uintptr) uintptr
	msgSend5        func(id, sel, a1, a2, a3, a4, a5 uintptr) uintptr
	msgSend6        func(id, sel, a1, a2, a3, a4, a5, a6 uintptr) uintptr
	msgSend7        func(id, sel, a1, a2, a3, a4, a5, a6, a7 uintptr) uintptr
	msgSend8        func(id, sel, a1, a2, a3, a4, a5, a6, a7, a8 uintptr) uintptr
	objcMsgSendAddr uintptr
)

func initFastSend() {
	var libobjcHandle uintptr
	var err error
	libobjcHandle, err = purego.Dlopen("/usr/lib/libobjc.A.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return // non-darwin or libobjc unavailable
	}
	objcMsgSendAddr, err = purego.Dlsym(libobjcHandle, "objc_msgSend")
	if err != nil {
		return
	}
	purego.RegisterFunc(&msgSend0, objcMsgSendAddr)
	purego.RegisterFunc(&msgSend1, objcMsgSendAddr)
	purego.RegisterFunc(&msgSend2, objcMsgSendAddr)
	purego.RegisterFunc(&msgSend3, objcMsgSendAddr)
	purego.RegisterFunc(&msgSend4, objcMsgSendAddr)
	purego.RegisterFunc(&msgSend5, objcMsgSendAddr)
	purego.RegisterFunc(&msgSend6, objcMsgSendAddr)
	purego.RegisterFunc(&msgSend7, objcMsgSendAddr)
	purego.RegisterFunc(&msgSend8, objcMsgSendAddr)
}

func cArrayArgIndexes(selector string, argc int) map[int]struct{} {
	if selector == "" || !strings.Contains(selector, "count:") || argc < 2 {
		return nil
	}
	parts := strings.Split(strings.TrimSuffix(selector, ":"), ":")
	out := map[int]struct{}{}
	for i, part := range parts {
		if part != "count" || i == 0 || i > argc {
			continue
		}
		out[i-1] = struct{}{}
	}
	return out
}

func toCArrayArg(arg any) (converted any, holder any, ok bool) {
	switch v := arg.(type) {
	case unsafe.Pointer:
		return v, nil, true
	case []ID:
		if len(v) == 0 {
			return unsafe.Pointer(nil), nil, true
		}
		return unsafe.Pointer(unsafe.SliceData(v)), v, true
	case []Class:
		if len(v) == 0 {
			return unsafe.Pointer(nil), nil, true
		}
		return unsafe.Pointer(unsafe.SliceData(v)), v, true
	case []string:
		if len(v) == 0 {
			return unsafe.Pointer(nil), nil, true
		}
		ids := make([]ID, len(v))
		for i := range v {
			ids[i] = String(v[i])
		}
		return unsafe.Pointer(unsafe.SliceData(ids)), ids, true
	}

	rv := reflect.ValueOf(arg)
	if !rv.IsValid() {
		return nil, nil, false
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return unsafe.Pointer(nil), nil, true
		}
		return unsafe.Pointer(rv.Pointer()), arg, true
	}
	if rv.Kind() != reflect.Slice {
		return nil, nil, false
	}
	if rv.Len() == 0 {
		return unsafe.Pointer(nil), nil, true
	}

	ids := make([]ID, rv.Len())
	allIDGetters := true
	for i := range ids {
		elem := rv.Index(i).Interface()
		if elem == nil {
			ids[i] = 0
			continue
		}
		getter, ok := elem.(IDGetter)
		if !ok {
			allIDGetters = false
			break
		}
		ids[i] = getter.GetID()
	}
	if allIDGetters {
		return unsafe.Pointer(unsafe.SliceData(ids)), ids, true
	}

	ptr := unsafe.Pointer(rv.Pointer())
	if ptr == nil {
		return nil, nil, false
	}
	return ptr, arg, true
}

// GetClass returns the class with the exact given name.
func GetClass(name string) Class {
	return pobjc.GetClass(name)
}

// GetProtocol returns the protocol with the given name, or nil if not found.
func GetProtocol(name string) *Protocol {
	return pobjc.GetProtocol(name)
}

// RegisterName registers a selector with the Objective-C runtime.
// This is the same as Sel but without caching - use Sel for repeated calls.
func RegisterName(name string) SEL {
	return pobjc.RegisterName(name)
}

// RegisterClass registers a new Objective-C class with the runtime.
// The class inherits from superClass and implements the given protocols.
// ivars defines instance variables, methods defines the class methods.
func RegisterClass(name string, superClass Class, protocols []*Protocol, ivars []FieldDef, methods []MethodDef) (Class, error) {
	return pobjc.RegisterClass(name, superClass, protocols, ivars, methods)
}

// SendSuper sends sel to the superclass of id's class.
//
// SendSuper derives the starting class from id's dynamic class, so it is only
// correct for methods on leaf classes. A method inherited by a further
// subclass would recurse into itself.
func SendSuper[T any](id ID, sel SEL, args ...any) T {
	return pobjc.SendSuper[T](id, sel, args...)
}

var (
	nsStringClass     Class
	selStringWithUTF8 SEL
	initOnce          sync.Once
)

func initStringHelpers() {
	nsStringClass = GetClass("NSString")
	selStringWithUTF8 = Sel("stringWithUTF8String:")
}

// String converts a Go string to an autoreleased NSString object.
//
// Callers on a thread without a run loop should enclose work that creates
// strings in AutoreleasePool. Otherwise autoreleased strings can accumulate
// for the lifetime of the thread.
func String(s string) ID {
	initOnce.Do(initStringHelpers)
	return Send[ID](ID(nsStringClass), selStringWithUTF8, s)
}

// IDToString converts an NSString ID to a Go string.
func IDToString(id ID) string {
	if id == 0 {
		return ""
	}
	// Call UTF8String
	cstr := Send[*byte](id, Sel("UTF8String"))
	return GoString(cstr)
}

// IDToStringPtr converts a nullable NSString ID to a heap-backed *string.
func IDToStringPtr(id ID) *string {
	if id == 0 {
		return nil
	}
	s := IDToString(id)
	return &s
}

// ErrUnrecognizedSelector is the sentinel error for an unavailable selector.
var ErrUnrecognizedSelector = errors.New("unrecognized selector")

// UnrecognizedSelectorError records a selector an object does not respond to.
type UnrecognizedSelectorError struct {
	Selector string
}

func (e *UnrecognizedSelectorError) Error() string {
	if e.Selector != "" {
		return ErrUnrecognizedSelector.Error() + ": " + e.Selector
	}
	return ErrUnrecognizedSelector.Error()
}

func (e *UnrecognizedSelectorError) Unwrap() error {
	return ErrUnrecognizedSelector
}

// RespondsToSelector checks if an object responds to the given selector.
// This is the safe way to check before calling a method.
func RespondsToSelector(id ID, sel SEL) bool {
	if id == 0 {
		return false
	}
	return Send[bool](id, Sel("respondsToSelector:"), sel)
}

// SendIfResponds calls a selector if the receiver responds to it and returns
// the zero value if it does not.
//
// It exists for bindings to private frameworks, where a selector may simply be
// absent: Apple ships no compatibility guarantee for them, and a selector that
// exists on one macOS build can be gone on the next. Sending to a receiver that
// does not respond raises NSInvalidArgumentException, which for a Go process is
// not an exception but an abort -- there is no recover, and the process dies
// inside the Objective-C runtime with a stack that does not name the caller.
//
// Returning the zero value is deliberately quiet, and that is a real tradeoff:
// a caller cannot tell "the selector is missing" from "the call returned zero".
// Use [SafeSend] where the caller can act on the difference. This variant is
// for generated bindings, whose signatures are fixed by the Objective-C method
// and have nowhere to put an error, and where the alternative is not a better
// error but a dead process.
//
// The receiver may be an instance or a class: NSObject implements
// respondsToSelector: on both, so a class object correctly reports its class
// methods.
func SendIfResponds[T any](id ID, sel SEL, args ...any) T {
	var zero T
	if !RespondsToSelector(id, sel) {
		return zero
	}
	return Send[T](id, sel, args...)
}

// SafeSend calls a selector only if the object responds to it.
// Returns the zero value and an error matching ErrUnrecognizedSelector if the
// selector is not recognized.
// This prevents NSInvalidArgumentException crashes from unrecognized selectors.
func SafeSend[T any](id ID, sel SEL, args ...any) (T, error) {
	var zero T
	if !RespondsToSelector(id, sel) {
		return zero, &UnrecognizedSelectorError{Selector: selName(sel)}
	}
	return Send[T](id, sel, args...), nil
}

// MustSend calls a selector and panics with a clear error if the object doesn't respond.
// Use this when you expect the selector to always exist but want a clearer panic message
// than the NSInvalidArgumentException.
func MustSend[T any](id ID, sel SEL, args ...any) T {
	if !RespondsToSelector(id, sel) {
		panic(&UnrecognizedSelectorError{Selector: selName(sel)})
	}
	return Send[T](id, sel, args...)
}

// selName returns the cached name of sel for error messages.
// Selectors not previously passed to Sel have no cached name.
func selName(sel SEL) string {
	// Iterate through cache to find the name
	var name string
	selCache.Range(func(key, value any) bool {
		if value.(SEL) == sel {
			name = key.(string)
			return false
		}
		return true
	})
	return name
}

var (
	class_addMethod          func(Class, SEL, uintptr, string) bool
	objc_registerClassPair   func(Class)
	objc_autoreleasePoolPush func() uintptr
	objc_autoreleasePoolPop  func(uintptr)
	libobjc                  uintptr
)

func ensureLibObjC() {
	if libobjc != 0 {
		return
	}
	var err error
	libobjc, err = purego.Dlopen("libobjc.A.dylib", purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	purego.RegisterLibFunc(&class_addMethod, libobjc, "class_addMethod")
	purego.RegisterLibFunc(&objc_registerClassPair, libobjc, "objc_registerClassPair")
	purego.RegisterLibFunc(&objc_autoreleasePoolPush, libobjc, "objc_autoreleasePoolPush")
	purego.RegisterLibFunc(&objc_autoreleasePoolPop, libobjc, "objc_autoreleasePoolPop")
}

// AutoreleasePool executes fn within an Objective-C autorelease pool.
// Any autoreleased objects created during fn are released when fn returns.
//
// AutoreleasePool pins the calling goroutine to its OS thread (via
// runtime.LockOSThread) for the duration of fn because Objective-C
// autorelease pools are thread-affine and must be popped on the thread
// that pushed them.
func AutoreleasePool(fn func()) {
	ensureLibObjC()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	pool := objc_autoreleasePoolPush()
	defer objc_autoreleasePoolPop(pool)
	fn()
}

// RegisterClassPair registers a class pair with the runtime.
func RegisterClassPair(cls Class) {
	ensureLibObjC()
	objc_registerClassPair(cls)
}

// NewIMP returns an Objective-C method implementation for fn.
//
// fn must take (ID, SEL) as its first two arguments, because that is how the
// runtime calls every method: the receiver and the selector arrive in the
// first two registers whether or not the Go function declares them. A function
// missing that prefix reads the receiver as its first declared argument and
// every later argument shifts, which corrupts the call silently rather than
// failing. NewIMP panics instead.
//
// The returned implementation is never deallocated.
func NewIMP(fn any) IMP {
	return pobjc.NewIMP(fn)
}

// AddMethod adds a new method to a class.
// impl is an Objective-C method implementation, such as the result of NewIMP.
func AddMethod(cls Class, sel SEL, impl IMP, types string) bool {
	ensureLibObjC()
	return class_addMethod(cls, sel, uintptr(impl), types)
}

// ObjCError is an Objective-C error returned through an NSError** parameter.
// ID is the underlying NSError object.
type ObjCError struct {
	ID ID
}

// Error returns the localized error description with its domain and code.
func (e ObjCError) Error() string {
	if e.ID == 0 {
		return "Objective-C error"
	}
	description := IDToString(Send[ID](e.ID, Sel("localizedDescription")))
	domain := IDToString(Send[ID](e.ID, Sel("domain")))
	code := Send[int](e.ID, Sel("code"))
	if domain == "" {
		if description != "" {
			return description
		}
		return "Objective-C error"
	}
	if description == "" {
		return fmt.Sprintf("%s error %d", domain, code)
	}
	return fmt.Sprintf("%s (%s error %d)", description, domain, code)
}

// SendWithError calls a selector that uses the NSError** pattern.
// It assumes the method accepts an NSError** as its last argument.
// It automatically appends the error pointer to the arguments.
func SendWithError[T any](id ID, sel SEL, args ...any) (T, error) {
	var err ID
	args = append(args, &err)
	ret := Send[T](id, sel, args...)
	if err != 0 {
		return ret, ObjCError{ID: err}
	}
	return ret, nil
}

// CallWithError calls a void-returning selector and handles the NSError** pattern.
func CallWithError(id ID, sel SEL, args ...any) error {
	_, err := SendWithError[ID](id, sel, args...)
	return err
}
