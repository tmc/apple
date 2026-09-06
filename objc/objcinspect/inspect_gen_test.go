//go:build darwin

package objcinspect

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
	"unsafe"

	basepurego "github.com/ebitengine/purego"
	purego "github.com/ebitengine/purego/objc"
)

func TestLooksLikeObject_RejectsSmallIntegers(t *testing.T) {
	if LooksLikeObject(purego.ID(40)) {
		t.Errorf("LooksLikeObject(40) = true, want false")
	}
}

func TestCheck_MissingSelectorReturnsErrSelectorNotFound(t *testing.T) {
	ensureLibObjC()

	nsobject := purego.GetClass("NSObject")
	if nsobject == 0 {
		t.Skip("NSObject class unavailable")
	}
	obj := purego.ID(purego.ID(nsobject).Send(purego.RegisterName("new")))
	defer purego.ID(obj).Send(purego.RegisterName("release"))

	bogusSel := purego.RegisterName("nonExistentSelectorThatDoesNotExist:")
	err := Check(obj, bogusSel, nil, 123)
	if err == nil {
		t.Fatal("expected error for non-existent selector, got nil")
	}
	if !errors.Is(err, ErrSelectorNotFound) {
		t.Fatalf("expected ErrSelectorNotFound, got: %v", err)
	}
}

func TestCheck_ValidCallReturnsNil(t *testing.T) {
	ensureLibObjC()

	nsobject := purego.GetClass("NSObject")
	if nsobject == 0 {
		t.Skip("NSObject class unavailable")
	}
	obj := purego.ID(purego.ID(nsobject).Send(purego.RegisterName("new")))
	defer purego.ID(obj).Send(purego.RegisterName("release"))

	isEqualSel := purego.RegisterName("isEqual:")
	err := Check(obj, isEqualSel, reflect.TypeOf(true), obj)
	if err != nil {
		t.Fatalf("Check(obj, isEqual:, BOOL, obj) failed: %v", err)
	}
}

func TestCheck_BOOLAcceptanceForKindChar(t *testing.T) {
	tyChar := Type{Encoding: "c", Kind: KindChar}
	if !matchType(tyChar, reflect.TypeOf(true)) {
		t.Fatalf("matchType(KindChar, bool) = false, want true for x86_64 BOOL compatibility")
	}
	tyUChar := Type{Encoding: "C", Kind: KindUChar}
	if !matchType(tyUChar, reflect.TypeOf(true)) {
		t.Fatalf("matchType(KindUChar, bool) = false, want true for x86_64 BOOL compatibility")
	}
}

func TestCheck_ReturnTypeMismatchReturnsErr(t *testing.T) {
	ensureLibObjC()

	nsobject := purego.GetClass("NSObject")
	if nsobject == 0 {
		t.Skip("NSObject class unavailable")
	}
	obj := purego.ID(purego.ID(nsobject).Send(purego.RegisterName("new")))
	defer purego.ID(obj).Send(purego.RegisterName("release"))

	isEqualSel := purego.RegisterName("isEqual:")
	err := Check(obj, isEqualSel, reflect.TypeOf(float64(0)), obj)
	if err == nil {
		t.Fatal("expected return type mismatch error, got nil")
	}
	if !strings.Contains(err.Error(), "wrong register file") {
		t.Errorf("expected register file error context, got: %v", err)
	}
}

func TestCheck_ArgTypeMismatchReturnsErr(t *testing.T) {
	ensureLibObjC()

	nsobject := purego.GetClass("NSObject")
	if nsobject == 0 {
		t.Skip("NSObject class unavailable")
	}
	obj := purego.ID(purego.ID(nsobject).Send(purego.RegisterName("new")))
	defer purego.ID(obj).Send(purego.RegisterName("release"))

	isEqualSel := purego.RegisterName("isEqual:")
	err := Check(obj, isEqualSel, reflect.TypeOf(true), 3.14)
	if err == nil {
		t.Fatal("expected arg type mismatch error, got nil")
	}
}

// countQuotedFieldNames counts the field-name quotes in a struct or union
// body by walking it the way the field loop does: read a quote if one is
// there, then consume one type. It parses in field context, so a quoted string
// after '@' is resolved the same way parseType resolves it; the oracle checks
// that the quote walk and the field walk stay in step, not that the ambiguity
// is resolved correctly. The round-trip oracle is what independently checks
// the latter.
func countQuotedFieldNames(fieldsBody string) int {
	fp := &parser{s: fieldsBody, inFields: true}
	count := 0
	for fp.i < len(fp.s) {
		if fp.s[fp.i] == '"' {
			count++
			fp.i++
			for fp.i < len(fp.s) && fp.s[fp.i] != '"' {
				fp.i++
			}
			if fp.i < len(fp.s) && fp.s[fp.i] == '"' {
				fp.i++
			}
		}
		if fp.i >= len(fp.s) || fp.s[fp.i] == '}' || fp.s[fp.i] == ')' {
			break
		}
		if fp.s[fp.i] != '"' {
			_, err := fp.parseType()
			if err != nil {
				break
			}
		}
	}
	return count
}

// TestLiveCorpusSweep parses every method and ivar type encoding reachable
// from objc_copyClassList on the running system and applies three oracles.
//
// Full consumption: an ivar encoding must be consumed to the last byte.
// Method encodings are consumed by construction, since ParseSignature loops
// until the input is exhausted.
//
// Quoted-name count: a struct whose body carries N quoted field names must
// parse into exactly N fields. This is the oracle that catches a parser
// reading the letters of a name as type codes.
//
// Round-trip, and this one is narrower than it sounds. It applies to struct
// encodings only, and it compares formatType(ty) against ty.Encoding, which is
// the exact source substring that struct was parsed from — not against the
// method or ivar encoding that contained it. Nothing is normalized on either
// side. Method-level qualifiers and stack offsets simply fall outside every
// struct substring, so the parser dropping r, n, N, o, O, R and V cannot show
// up here: "*@:r*" does not round-trip to itself, and the sweep never asks it
// to. A qualifier written inside a struct body would be caught, and the sweep
// reports no mismatches.
//
// The sweep's blind spot is extended block encodings. method_getTypeEncoding
// renders a block argument as a bare "@?"; the "@?<...>" form lives in
// protocol extended method types and block descriptors, neither of which this
// sweep walks. However many hundreds of thousands of encodings pass here, they
// give ParseBlockEncoding zero coverage. Its only guard is the fixture test.
func TestLiveCorpusSweep(t *testing.T) {
	ensureLibObjC()
	for _, fw := range []string{
		"/System/Library/Frameworks/Foundation.framework/Foundation",
		"/System/Library/Frameworks/AppKit.framework/AppKit",
		"/System/Library/Frameworks/CoreGraphics.framework/CoreGraphics",
		"/System/Library/Frameworks/Metal.framework/Metal",
		"/System/Library/Frameworks/AVFoundation.framework/AVFoundation",
		"/System/Library/Frameworks/CoreData.framework/CoreData",
		"/System/Library/Frameworks/QuartzCore.framework/QuartzCore",
	} {
		basepurego.Dlopen(fw, basepurego.RTLD_LAZY|basepurego.RTLD_GLOBAL)
	}

	// These return malloc'd C arrays. Declaring them as *uintptr rather than
	// uintptr lets unsafe.Slice take the pointer directly, with no uintptr to
	// unsafe.Pointer round trip for vet to flag.
	var objc_copyClassList func(outCount *uint32) *uintptr
	var class_copyMethodList func(cls uintptr, outCount *uint32) *uintptr
	var class_copyIvarList func(cls uintptr, outCount *uint32) *uintptr
	var ivar_getTypeEncoding func(ivar uintptr) *byte

	basepurego.RegisterLibFunc(&objc_copyClassList, libobjc, "objc_copyClassList")
	basepurego.RegisterLibFunc(&class_copyMethodList, libobjc, "class_copyMethodList")
	basepurego.RegisterLibFunc(&class_copyIvarList, libobjc, "class_copyIvarList")
	basepurego.RegisterLibFunc(&ivar_getTypeEncoding, libobjc, "ivar_getTypeEncoding")

	var n uint32
	clsPtr := objc_copyClassList(&n)
	if clsPtr == nil || n == 0 {
		t.Fatal("objc_copyClassList returned no classes")
	}
	classes := unsafe.Slice(clsPtr, int(n))

	var totalMethods, methodFail, totalIvars, ivarFail, structCount, structQuotedMismatch, roundTripFail int
	var methodFailSamples, ivarFailSamples, quotedMismatchSamples, roundTripSamples []string

	// Accumulate every struct definition the sweep walks past, so that a tag
	// spelled opaquely in one encoding can be resolved against a spelling that
	// carries fields from elsewhere in the same process.
	reg := NewRegistry()
	opaqueTags := map[string]bool{}

	var checkStruct func(ty Type)
	checkStruct = func(ty Type) {
		if ty.Kind == KindStruct && ty.Name != "" {
			if len(ty.Fields) > 0 {
				reg.Register(ty)
			} else {
				opaqueTags[ty.Name] = true
			}
		}
		if ty.Kind == KindStruct && strings.Contains(ty.Encoding, "=") {
			structCount++
			inner := ty.Encoding[1 : len(ty.Encoding)-1]
			eq := strings.IndexByte(inner, '=')
			if eq >= 0 {
				fieldsBody := inner[eq+1:]
				quoteCount := countQuotedFieldNames(fieldsBody)
				if quoteCount > 0 && quoteCount != len(ty.Fields) {
					structQuotedMismatch++
					if len(quotedMismatchSamples) < 5 {
						quotedMismatchSamples = append(quotedMismatchSamples, fmt.Sprintf("%s: quoted=%d fields=%d", ty.Encoding, quoteCount, len(ty.Fields)))
					}
				}
			}
			reencoded := formatType(ty)
			if reencoded != ty.Encoding {
				roundTripFail++
				if len(roundTripSamples) < 5 {
					roundTripSamples = append(roundTripSamples, fmt.Sprintf("%s: got %s", ty.Encoding, reencoded))
				}
			}
		}
		for _, f := range ty.Fields {
			checkStruct(f)
		}
		if ty.Elem != nil {
			checkStruct(*ty.Elem)
		}
	}

	for _, cls := range classes {
		var mn uint32
		if mPtr := class_copyMethodList(cls, &mn); mPtr != nil {
			for _, m := range unsafe.Slice(mPtr, int(mn)) {
				encPtr := method_getTypeEncoding(m)
				if encPtr == nil {
					continue
				}
				enc := gostring(encPtr)
				totalMethods++
				sig, err := ParseSignature(enc)
				if err != nil {
					methodFail++
					if len(methodFailSamples) < 5 {
						methodFailSamples = append(methodFailSamples, fmt.Sprintf("%q: %v", enc, err))
					}
					continue
				}
				checkStruct(sig.Return)
				for _, arg := range sig.Args {
					checkStruct(arg)
				}
			}
		}
		var in uint32
		if iPtr := class_copyIvarList(cls, &in); iPtr != nil {
			for _, iv := range unsafe.Slice(iPtr, int(in)) {
				ep := ivar_getTypeEncoding(iv)
				if ep == nil {
					continue
				}
				enc := gostring(ep)
				if enc == "" {
					continue
				}
				totalIvars++
				p := &parser{s: enc}
				ty, err := p.parseType()
				p.skipQualifiersAndOffsets()
				if err != nil || p.i < len(p.s) {
					ivarFail++
					if len(ivarFailSamples) < 10 {
						ivarFailSamples = append(ivarFailSamples, fmt.Sprintf("%q (parsed %d/%d): %v", enc, p.i, len(p.s), err))
					}
					continue
				}
				checkStruct(ty)
			}
		}
	}

	// Tags seen both ways: opaque in at least one encoding, defined in another.
	// Resolving one must produce that tag's definition, so the resolved type
	// has to keep the name it was asked about and come back with fields.
	//
	// It does not have to come back with a size. Resolve rewrites the tag it
	// was given and the fields below it, but a definition reached that way can
	// itself contain an opaque field or a bitfield, and NaturalSize answers -1
	// for both. Counting the sized ones separately keeps that visible instead
	// of asserting a size the encoding never promised.
	var resolvedTags []string
	sized := 0
	for tag := range opaqueTags {
		opaque := Type{Kind: KindStruct, Name: tag, Encoding: "{" + tag + "=}"}
		if opaque.NaturalSize() != -1 {
			t.Errorf("opaque %s: NaturalSize = %d, want -1", tag, opaque.NaturalSize())
		}
		r := reg.Resolve(opaque)
		if len(r.Fields) == 0 {
			continue
		}
		if r.Name != tag {
			t.Errorf("resolved %s: Name = %q, want %q", tag, r.Name, tag)
		}
		if n := r.NaturalSize(); n > 0 {
			sized++
		} else if n != -1 {
			t.Errorf("resolved %s: NaturalSize = %d, want -1 or positive", tag, n)
		}
		resolvedTags = append(resolvedTags, tag)
	}
	sort.Strings(resolvedTags)

	t.Logf("Live sweep summary: methods=%d (failures=%d), ivars=%d (failures=%d), structs checked=%d, quoted mismatch=%d, roundtrip mismatch=%d",
		totalMethods, methodFail, totalIvars, ivarFail, structCount, structQuotedMismatch, roundTripFail)
	t.Logf("Registry: %d definitions, %d opaque tags, %d resolved across encodings (%d of those gain a size)",
		reg.Len(), len(opaqueTags), len(resolvedTags), sized)
	t.Logf("Registry resolved: %s", strings.Join(resolvedTags, " "))

	for _, s := range methodFailSamples {
		t.Logf("method fail sample: %s", s)
	}
	for _, s := range ivarFailSamples {
		t.Logf("ivar fail sample: %s", s)
	}
	for _, s := range quotedMismatchSamples {
		t.Logf("quoted mismatch sample: %s", s)
	}
	for _, s := range roundTripSamples {
		t.Logf("roundtrip sample: %s", s)
	}

	if methodFail > 0 || ivarFail > 0 || structQuotedMismatch > 0 || roundTripFail > 0 {
		t.Fatalf("Live sweep failed: methodFail=%d ivarFail=%d quotedMismatch=%d roundTripFail=%d", methodFail, ivarFail, structQuotedMismatch, roundTripFail)
	}
}
