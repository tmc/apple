// Code generated from internal/generator/templates/runtime/objc.txtar by applegen. DO NOT EDIT.

//go:build darwin

package objcinspect

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	basepurego "github.com/ebitengine/purego"
	purego "github.com/ebitengine/purego/objc"
)

var (
	ErrSelectorNotFound = errors.New("selector not found")

	libobjcOnce             sync.Once
	libobjc                 uintptr
	object_getClass         func(id purego.ID) purego.Class
	class_getInstanceMethod func(cls purego.Class, sel purego.SEL) uintptr
	method_getTypeEncoding  func(method uintptr) *byte
	methodCache             sync.Map
)

// LooksLikeObject reports whether id plausibly points to an Objective-C object.
//
// LooksLikeObject NEVER dereferences memory. A false return value is conclusive
// (id is definitely not a valid object), while a true return value indicates id is in
// a plausible address range (> 0x1000, 8-byte aligned) or is an arm64 tagged pointer.
func LooksLikeObject(id purego.ID) bool {
	return looksLikeObjectAddr(uintptr(id))
}

func ensureLibObjC() {
	libobjcOnce.Do(func() {
		var err error
		libobjc, err = basepurego.Dlopen("libobjc.A.dylib", basepurego.RTLD_GLOBAL)
		if err != nil {
			panic(fmt.Sprintf("objcinspect: failed to dlopen libobjc: %v", err))
		}
		basepurego.RegisterLibFunc(&object_getClass, libobjc, "object_getClass")
		basepurego.RegisterLibFunc(&class_getInstanceMethod, libobjc, "class_getInstanceMethod")
		basepurego.RegisterLibFunc(&method_getTypeEncoding, libobjc, "method_getTypeEncoding")
	})
}

type methodCacheKey struct {
	cls purego.Class
	sel purego.SEL
}

type cachedMethodInfo struct {
	sig Signature
	err error
}

// Check validates that sending sel to id with args matches the runtime method signature.
// It returns an error if id is not an object, if sel is not implemented by id's class,
// or if return/argument type families mismatch.
//
// Check has three deliberate limits:
//
// Matching is family-level, not width-level. An integer argument satisfies any
// integer parameter, and reflect.Bool satisfies signed char and _Bool alike,
// because BOOL encodes as 'B' on arm64 but 'c' on x86_64. A width mismatch
// truncates a value; a float/int crossing reads it from the wrong register file,
// which is the failure Check exists to catch.
//
// A struct argument satisfies an object parameter, so a generated Go wrapper
// struct passes where an id is expected.
//
// Qualifiers are not checked. The parser consumes r, n, N, o, O, R and V
// without recording them, so Check cannot distinguish const char * from
// char *, or an out parameter from an in one.
//
// Check does NOT invoke the method or call objc.Send.
func Check(id purego.ID, sel purego.SEL, want reflect.Type, args ...any) error {
	if !LooksLikeObject(id) {
		return fmt.Errorf("invalid receiver ID: %v", id)
	}

	ensureLibObjC()

	cls := object_getClass(id)
	if cls == 0 {
		return fmt.Errorf("invalid class for receiver ID: %v", id)
	}

	key := methodCacheKey{cls: cls, sel: sel}
	var info cachedMethodInfo
	if val, ok := methodCache.Load(key); ok {
		info = val.(cachedMethodInfo)
	} else {
		method := class_getInstanceMethod(cls, sel)
		if method == 0 {
			info = cachedMethodInfo{err: fmt.Errorf("%w: selector %d not found on class 0x%x", ErrSelectorNotFound, sel, uintptr(cls))}
		} else {
			encPtr := method_getTypeEncoding(method)
			if encPtr == nil {
				info = cachedMethodInfo{err: fmt.Errorf("nil type encoding for selector %v", sel)}
			} else {
				enc := gostring(encPtr)
				sig, err := ParseSignature(enc)
				if err != nil {
					info = cachedMethodInfo{err: fmt.Errorf("parse method encoding %q: %w", enc, err)}
				} else {
					info = cachedMethodInfo{sig: sig}
				}
			}
		}
		methodCache.Store(key, info)
	}

	if info.err != nil {
		return info.err
	}

	expectedArgCount := 0
	if len(info.sig.Args) >= 2 {
		expectedArgCount = len(info.sig.Args) - 2
	}

	if len(args) != expectedArgCount {
		return fmt.Errorf("argument count mismatch: got %d args, selector expects %d", len(args), expectedArgCount)
	}

	// Validate return type if specified
	if want != nil && want.Kind() != reflect.Invalid {
		if info.sig.Return.Kind == KindVoid {
			return fmt.Errorf("return type mismatch: method returns void, want %v", want)
		}
		if !matchType(info.sig.Return, want) {
			extra := ""
			if (isFloatKind(info.sig.Return.Kind) && isIntReflect(want.Kind())) || (isIntKind(info.sig.Return.Kind) && isFloatReflect(want.Kind())) {
				extra = " (float/int mismatch: value would be read from wrong register file)"
			}
			return fmt.Errorf("return type mismatch: method returns %v (%s), want %v%s", info.sig.Return.Kind, info.sig.Return.Encoding, want, extra)
		}
	}

	// Validate argument types
	for i, arg := range args {
		targetArg := info.sig.Args[i+2]
		var argType reflect.Type
		if arg != nil {
			argType = reflect.TypeOf(arg)
		}
		if !matchType(targetArg, argType) {
			extra := ""
			if argType != nil {
				if (isFloatKind(targetArg.Kind) && isIntReflect(argType.Kind())) || (isIntKind(targetArg.Kind) && isFloatReflect(argType.Kind())) {
					extra = " (float/int mismatch: value would be read from wrong register file)"
				}
			}
			return fmt.Errorf("arg %d type mismatch: method expects %v (%s), got %v%s", i, targetArg.Kind, targetArg.Encoding, argType, extra)
		}
	}

	return nil
}

func isFloatKind(k Kind) bool {
	return k == KindFloat || k == KindDouble
}

func isIntKind(k Kind) bool {
	switch k {
	case KindChar, KindUChar, KindShort, KindUShort, KindInt, KindUInt, KindLong, KindULong, KindLongLong, KindULongLong, KindBool:
		return true
	default:
		return false
	}
}

func isFloatReflect(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
}

func isIntReflect(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Bool:
		return true
	default:
		return false
	}
}

func matchType(expected Type, got reflect.Type) bool {
	if got == nil {
		return expected.Kind == KindId || expected.Kind == KindClass || expected.Kind == KindSEL || expected.Kind == KindPointer || expected.Kind == KindBlock
	}
	k := got.Kind()
	switch expected.Kind {
	case KindVoid:
		return k == reflect.Invalid
	case KindId, KindClass, KindSEL, KindPointer, KindBlock, KindProtocol:
		return k == reflect.Uintptr || k == reflect.Pointer || k == reflect.UnsafePointer || k == reflect.Interface || k == reflect.Struct
	case KindBool:
		return k == reflect.Bool || k == reflect.Uint8 || k == reflect.Int || k == reflect.Int32
	case KindChar, KindShort, KindInt, KindLong, KindLongLong:
		// On x86_64 macOS, BOOL is typedef'd to signed char ('c'), so reflect.Bool is accepted here as well as for KindBool ('B').
		return k == reflect.Bool || k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 || k == reflect.Uintptr
	case KindUChar, KindUShort, KindUInt, KindULong, KindULongLong:
		// On x86_64 macOS, BOOL is typedef'd to unsigned char ('C') in some contexts, so reflect.Bool is accepted here.
		return k == reflect.Bool || k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 || k == reflect.Uintptr || k == reflect.Int
	case KindFloat, KindDouble:
		return k == reflect.Float32 || k == reflect.Float64
	case KindStruct, KindUnion:
		return k == reflect.Struct || k == reflect.Array || k == reflect.Slice || k == reflect.Uintptr
	default:
		return true
	}
}

func gostring(ptr *byte) string {
	if ptr == nil {
		return ""
	}
	var b []byte
	for p := ptr; *p != 0; p = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) {
		b = append(b, *p)
	}
	return string(b)
}
