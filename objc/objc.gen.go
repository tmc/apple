// Code generated from Apple documentation by applegen. DO NOT EDIT.

// Package objc provides cached Objective-C runtime helpers.
//
// This package wraps purego/objc to provide selector caching for better performance.
package objc

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	basepurego "github.com/ebitengine/purego"
	purego "github.com/ebitengine/purego/objc"
)

// Type aliases for convenience
type (
	ID        = purego.ID
	SEL       = purego.SEL
	Class     = purego.Class
	Block     = purego.Block
	Protocol  = purego.Protocol
	MethodDef = purego.MethodDef
	FieldDef  = purego.FieldDef
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

// NewBlock creates an Objective-C block from a Go function.
// The Go function must take a Block as its first argument.
// Use Block.Release() to free the block when it is no longer in use.
func NewBlock(fn any) Block {
	return purego.NewBlock(fn)
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

// GoStringPtr converts a nullable C string to a heap-backed *string.
func GoStringPtr(cstr *byte) *string {
	if cstr == nil {
		return nil
	}
	s := GoString(cstr)
	return &s
}

var (
	selCache sync.Map // map[string]purego.SEL
)

// Sel returns a cached selector for the given name.
// This avoids the global lock in purego.RegisterName on repeated calls.
func Sel(name string) SEL {
	if sel, ok := selCache.Load(name); ok {
		return sel.(SEL)
	}
	sel := purego.RegisterName(name)
	selCache.Store(name, sel)
	return sel
}

// NSArrayToSlice converts an NSArray ID into a []ID by calling count and objectAtIndex:.
func NSArrayToSlice(array ID) []ID {
	if array == 0 {
		return nil
	}
	count := purego.Send[uint](array, RegisterName("count"))
	if count == 0 {
		return nil
	}
	sel := RegisterName("objectAtIndex:")
	result := make([]ID, count)
	for i := uint(0); i < count; i++ {
		result[i] = purego.Send[ID](array, sel, i)
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
// function that avoids reflect.MakeFunc — zero allocations, ~10x faster.
//
// Otherwise Send falls back to purego.Send[T] with full argument processing:
// IDGetter extraction, nil→ID(0), CArrayArg conversion, and NSArray→[]ID.
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
				var result any = ID(rv)
				return result.(T)
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
		arrayID := purego.Send[ID](id, sel, args...)
		runtime.KeepAlive(keepAlive)
		var result any = NSArrayToSlice(arrayID)
		return result.(T)
	}
	ret := purego.Send[T](id, sel, args...)
	runtime.KeepAlive(keepAlive)
	return ret
}

// tryFastArgs attempts to convert all args to uintptr values.
// Returns the uintptr slice and true if all args are uintptr-castable,
// or nil, false if any arg requires the slow path.
func tryFastArgs(args []any) ([]uintptr, bool) {
	var buf [8]uintptr
	uargs := buf[:len(args)]
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
		r, _, _ := basepurego.SyscallN(objcMsgSendAddr, all...)
		return r
	}
}

// Pre-registered typed msgSend functions for the zero-allocation fast path.
// On arm64/amd64, purego's RegisterFunc produces direct assembly stubs when
// all parameters are uintptr, avoiding reflect.MakeFunc entirely.
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
	libobjcHandle, err = basepurego.Dlopen("/usr/lib/libobjc.A.dylib", basepurego.RTLD_LAZY|basepurego.RTLD_GLOBAL)
	if err != nil {
		return // non-darwin or libobjc unavailable
	}
	objcMsgSendAddr, err = basepurego.Dlsym(libobjcHandle, "objc_msgSend")
	if err != nil {
		return
	}
	basepurego.RegisterFunc(&msgSend0, objcMsgSendAddr)
	basepurego.RegisterFunc(&msgSend1, objcMsgSendAddr)
	basepurego.RegisterFunc(&msgSend2, objcMsgSendAddr)
	basepurego.RegisterFunc(&msgSend3, objcMsgSendAddr)
	basepurego.RegisterFunc(&msgSend4, objcMsgSendAddr)
	basepurego.RegisterFunc(&msgSend5, objcMsgSendAddr)
	basepurego.RegisterFunc(&msgSend6, objcMsgSendAddr)
	basepurego.RegisterFunc(&msgSend7, objcMsgSendAddr)
	basepurego.RegisterFunc(&msgSend8, objcMsgSendAddr)
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

// GetClass returns the class with the given name.
func GetClass(name string) Class {
	if name == "" {
		return 0
	}
	if cls := purego.GetClass(name); cls != 0 {
		return cls
	}
	// Runtime class names in private frameworks sometimes differ only by
	// a leading underscore. Try both forms as a compatibility fallback.
	if name[0] == '_' {
		return purego.GetClass(name[1:])
	}
	return purego.GetClass("_" + name)
}

// GetProtocol returns the protocol with the given name, or nil if not found.
func GetProtocol(name string) *Protocol {
	return purego.GetProtocol(name)
}

// RegisterName registers a selector with the Objective-C runtime.
// This is the same as Sel but without caching - use Sel for repeated calls.
func RegisterName(name string) SEL {
	return purego.RegisterName(name)
}

// RegisterClass registers a new Objective-C class with the runtime.
// The class inherits from superClass and implements the given protocols.
// ivars defines instance variables, methods defines the class methods.
func RegisterClass(name string, superClass Class, protocols []*Protocol, ivars []FieldDef, methods []MethodDef) (Class, error) {
	return purego.RegisterClass(name, superClass, protocols, ivars, methods)
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

// String converts a Go string to an NSString object.
// This must be called before passing Go strings to Objective-C methods that expect NSString*.
// The returned ID is autoreleased.
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

// ErrUnrecognizedSelector is returned when an object does not respond to a selector.
var ErrUnrecognizedSelector = &UnrecognizedSelectorError{}

// UnrecognizedSelectorError indicates an object does not respond to a selector.
type UnrecognizedSelectorError struct {
	Selector string
}

func (e *UnrecognizedSelectorError) Error() string {
	if e.Selector != "" {
		return "unrecognized selector: " + e.Selector
	}
	return "unrecognized selector"
}

// RespondsToSelector checks if an object responds to the given selector.
// This is the safe way to check before calling a method.
func RespondsToSelector(id ID, sel SEL) bool {
	if id == 0 {
		return false
	}
	return Send[bool](id, Sel("respondsToSelector:"), sel)
}

// SafeSend calls a selector only if the object responds to it.
// Returns the zero value and ErrUnrecognizedSelector if the selector is not recognized.
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

// selName attempts to get the name of a selector for error messages.
// Returns empty string if it can't be determined.
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
	libobjc, err = basepurego.Dlopen("libobjc.A.dylib", basepurego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	basepurego.RegisterLibFunc(&class_addMethod, libobjc, "class_addMethod")
	basepurego.RegisterLibFunc(&objc_registerClassPair, libobjc, "objc_registerClassPair")
	basepurego.RegisterLibFunc(&objc_autoreleasePoolPush, libobjc, "objc_autoreleasePoolPush")
	basepurego.RegisterLibFunc(&objc_autoreleasePoolPop, libobjc, "objc_autoreleasePoolPop")
}

// AutoreleasePool executes fn within an Objective-C autorelease pool.
// Any autoreleased objects created during fn are released when fn returns.
func AutoreleasePool(fn func()) {
	ensureLibObjC()
	pool := objc_autoreleasePoolPush()
	defer objc_autoreleasePoolPop(pool)
	fn()
}

// RegisterClassPair registers a class pair with the runtime.
func RegisterClassPair(cls Class) {
	ensureLibObjC()
	objc_registerClassPair(cls)
}

// AddMethod adds a new method to a class.
func AddMethod(cls Class, sel SEL, impl any, types string) bool {
	ensureLibObjC()
	var imp uintptr
	switch v := impl.(type) {
	case uintptr:
		imp = v
	case func():
		// Should use NewCallback, but here we expect already created callback IMP
		panic("AddMethod expects uintptr IMP (use purego.NewCallback)")
	default:
		panic(fmt.Sprintf("AddMethod expects uintptr IMP, got %T", impl))
	}
	return class_addMethod(cls, sel, imp, types)
}

// SendWithError calls a selector handles the NSError** pattern.
// It assumes the method accepts an NSError** as its last argument.
// It automatically appends the error pointer to the arguments.
func SendWithError[T any](id ID, sel SEL, args ...any) (T, error) {
	var err ID
	args = append(args, &err)
	ret := Send[T](id, sel, args...)
	if err != 0 {
		return ret, fmt.Errorf("Objective-C error: %v", err)
	}
	return ret, nil
}

// CallWithError calls a void-returning selector and handles the NSError** pattern.
func CallWithError(id ID, sel SEL, args ...any) error {
	_, err := SendWithError[ID](id, sel, args...)
	return err
}
