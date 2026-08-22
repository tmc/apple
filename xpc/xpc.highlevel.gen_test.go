// Code generated from Apple documentation for XPC. DO NOT EDIT.

package xpc

import (
	"context"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"unsafe"

	"github.com/tmc/apple/dispatch"
)

func TestBorrowedHandleRoundTrip(t *testing.T) {
	const handle = uintptr(0x1234)
	// PeerRequirement is deliberately absent: unlike the other six types,
	// PeerRequirementFromHandle retains and returns an owned reference, so a
	// fake handle would be retained (and later released) by the native side.
	tests := []struct {
		name string
		got  uintptr
	}{
		{"Endpoint", EndpointFromHandle(handle).Handle()},
		{"RichError", RichErrorFromHandle(handle).Handle()},
		{"Listener", ListenerFromHandle(handle).Handle()},
		{"Session", SessionFromHandle(handle).Handle()},
		{"ReceivedMessage", ReceivedMessageFromHandle(handle).Handle()},
		{"IncomingSessionRequest", IncomingSessionRequestFromHandle(handle).Handle()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.got != handle {
				t.Fatalf("Handle() = %#x, want %#x", test.got, handle)
			}
		})
	}
}

func TestNilPointerWrapperHandle(t *testing.T) {
	var listener *Listener
	if got := listener.Handle(); got != 0 {
		t.Fatalf("nil Listener.Handle() = %#x, want 0", got)
	}
	var session *Session
	if got := session.Handle(); got != 0 {
		t.Fatalf("nil Session.Handle() = %#x, want 0", got)
	}
}

// TestCodecRoundTripHonoursXPCTags guards the two halves of the codec against
// drifting apart. Encoding read xpc tags while decoding fell through to
// encoding/json, which does not, so a tagged field encoded under its tag name
// and decoded into nothing, keeping its zero value with no error reported.
func TestCodecRoundTripHonoursXPCTags(t *testing.T) {
	type payload struct {
		First  int64  `xpc:"firstNumber"`
		Second int64  `xpc:"secondNumber"`
		Name   string `json:"name"`
		Plain  bool
	}
	want := payload{First: 23, Second: 19, Name: "add", Plain: true}

	dict, err := encodeMessage(want)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	for _, key := range []string{"firstNumber", "secondNumber", "name", "Plain"} {
		if _, ok := dict[key]; !ok {
			t.Errorf("encoded dictionary missing key %q: %v", key, dict)
		}
	}

	var got payload
	if err := decodeMessage(ReceivedMessage{decoded: dict}, &got); err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if got != want {
		t.Errorf("round trip = %+v, want %+v", got, want)
	}
}

// TestCodecDecodeWidensIntegers checks that XPC's int64 lands in narrower Go
// fields, and that a value above 2^53 survives, which it would not if the
// decode path routed through a float.
func TestCodecDecodeWidensIntegers(t *testing.T) {
	type payload struct {
		Small int   `xpc:"small"`
		Big   int64 `xpc:"big"`
	}
	dict := Dictionary{"small": int64(7), "big": int64(9007199254740993)}

	var got payload
	if err := decodeMessage(ReceivedMessage{decoded: dict}, &got); err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if got.Small != 7 {
		t.Errorf("Small = %d, want 7", got.Small)
	}
	if got.Big != 9007199254740993 {
		t.Errorf("Big = %d, want 9007199254740993", got.Big)
	}
}

// canonicalDictionary renders d as sorted, type-tagged lines. It is the golden
// format for the codec tests. It depends on nothing but the standard library
// and on Dictionary itself, so it cannot drift with the codec it measures.
//
// Keys are sorted; nested dictionaries and maps are flattened with ".";
// slice elements are indexed ".0", ".1"; the type column is %T and the value
// column is %v, except []byte, which is hex. An empty dictionary renders as
// the single line "<empty>".
func canonicalDictionary(d Dictionary) string {
	var lines []string
	var walk func(prefix string, v any)
	emit := func(key string, v any) {
		if b, ok := v.([]byte); ok {
			lines = append(lines, fmt.Sprintf("%s\t%T\t%x", key, v, b))
			return
		}
		lines = append(lines, fmt.Sprintf("%s\t%T\t%v", key, v, v))
	}
	walk = func(prefix string, v any) {
		switch x := v.(type) {
		case Dictionary:
			walkMap(prefix, map[string]any(x), walk)
			return
		case map[string]any:
			walkMap(prefix, x, walk)
			return
		case []byte:
			emit(prefix, x)
			return
		case []any:
			if len(x) == 0 {
				lines = append(lines, prefix+"\t[]interface {}\t<empty>")
				return
			}
			for i, e := range x {
				walk(fmt.Sprintf("%s.%d", prefix, i), e)
			}
			return
		}
		emit(prefix, v)
	}
	walkMap("", map[string]any(d), walk)
	if len(lines) == 0 {
		return "<empty>\n"
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n") + "\n"
}

// walkMap visits m's entries in sorted key order, joining keys with ".".
func walkMap(prefix string, m map[string]any, walk func(string, any)) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		key := k
		if prefix != "" {
			key = prefix + "." + k
		}
		walk(key, m[k])
	}
}

type goldenMarshaler struct{}

func (goldenMarshaler) MarshalXPC() (Dictionary, error) {
	return Dictionary{"marshaler": "xpc"}, nil
}

type goldenTextMarshaler struct{}

func (goldenTextMarshaler) MarshalText() ([]byte, error) { return []byte("text"), nil }

type goldenBinaryMarshaler struct{}

func (goldenBinaryMarshaler) MarshalBinary() ([]byte, error) { return []byte("binary"), nil }

type goldenJSONMarshaler struct{}

func (goldenJSONMarshaler) MarshalJSON() ([]byte, error) {
	return []byte(`{"n":1,"s":"x"}`), nil
}

type goldenTagged struct {
	A string `xpc:"xpcName"`
	B int64  `xpc:"xpcNum"`
}

type goldenJSONTagged struct {
	A string `json:"jsonName"`
	B int64  `json:"jsonNum"`
}

type goldenUntagged struct {
	A string
	B int64
}

type goldenNested struct {
	Inner goldenTagged `xpc:"inner"`
	Flag  bool         `xpc:"flag"`
}

// codecGoldenCases is the fixture set for TestCodecEncodeGoldens. There is one
// case per branch of encodeMessage, so a deleted branch cannot pass. Every
// case with wantErr false must have a matching file in testdata/codec, and
// every file in that directory must have a matching case;
// TestCodecGoldenFilesAreExhaustive enforces both directions so a deleted case
// cannot silently retire a golden.
var codecGoldenCases = []struct {
	name    string
	value   any
	wantErr bool
}{
	{name: "nil", value: nil},
	{name: "dictionary", value: Dictionary{"a": int64(1), "b": "two"}},
	{name: "map_string_any", value: map[string]any{"a": int64(1), "b": true}},
	{name: "map_string_int64", value: map[string]int64{"a": 1, "b": 2}},
	{name: "map_int_key", value: map[int]string{1: "x"}, wantErr: true},
	{name: "pointer_struct", value: &goldenTagged{A: "p", B: 2}},
	{name: "nil_pointer", value: (*goldenTagged)(nil)},
	{name: "struct_xpc_tag", value: goldenTagged{A: "s", B: 3}},
	{name: "struct_json_tag", value: goldenJSONTagged{A: "j", B: 4}},
	{name: "struct_untagged", value: goldenUntagged{A: "u", B: 5}},
	{name: "struct_nested", value: goldenNested{Inner: goldenTagged{A: "in", B: 6}, Flag: true}},
	{name: "slice_int64", value: []int64{7, 8}},
	{name: "array_int64", value: [2]int64{9, 10}},
	{name: "string", value: "add"},
	{name: "int64", value: int64(23)},
	{name: "bool", value: true},
	{name: "bytes", value: []byte("hello")},
	{name: "marshaler", value: goldenMarshaler{}},
	{name: "text_marshaler", value: goldenTextMarshaler{}},
	{name: "binary_marshaler", value: goldenBinaryMarshaler{}},
	{name: "json_marshaler", value: goldenJSONMarshaler{}},
	{name: "wide_int64", value: int64(9007199254740993)},
}

// codecGoldenDir is where the hand-written goldens live. The files are not
// emitted from any template: a generated expectation would regenerate itself
// alongside a wire change and the gate would pass while the wire moved.
const codecGoldenDir = "testdata/codec"

// TestCodecEncodeGoldens pins the wire shape produced by encodeMessage. There
// is deliberately no -update flag: writing a golden is a hand edit, so a wire
// change always appears in review as a diff to testdata/codec.
func TestCodecEncodeGoldens(t *testing.T) {
	for _, c := range codecGoldenCases {
		t.Run(c.name, func(t *testing.T) {
			dict, err := encodeMessage(c.value)
			if c.wantErr {
				if err == nil {
					t.Fatalf("encodeMessage(%s) = %v, want error", c.name, dict)
				}
				return
			}
			if err != nil {
				t.Fatalf("encodeMessage(%s): %v", c.name, err)
			}
			path := filepath.Join("testdata", "codec", c.name+".golden")
			want, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("read golden: %v (write it by hand; there is no update flag)", err)
			}
			if got := canonicalDictionary(dict); got != string(want) {
				t.Errorf("wire shape changed for %s\n--- got ---\n%s--- want (%s) ---\n%s", c.name, got, path, want)
			}
		})
	}
}

// TestCodecRoundTripGoldens checks that decodeMessage inverts encodeMessage
// for the fixtures that have an inverse. The marshaler, scalar, slice and
// error fixtures are absent by construction: encodeMessage is lossy for them
// (a TextMarshaler has no Unmarshaler, a scalar becomes a "value" envelope),
// so a round trip would assert something the format does not promise.
func TestCodecRoundTripGoldens(t *testing.T) {
	cases := []struct {
		name string
		want any
		into func() any
	}{
		{"struct_xpc_tag", goldenTagged{A: "s", B: 3}, func() any { return new(goldenTagged) }},
		{"struct_json_tag", goldenJSONTagged{A: "j", B: 4}, func() any { return new(goldenJSONTagged) }},
		{"struct_untagged", goldenUntagged{A: "u", B: 5}, func() any { return new(goldenUntagged) }},
		{"wide_int64_field", goldenTagged{A: "w", B: 9007199254740993}, func() any { return new(goldenTagged) }},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			dict, err := encodeMessage(c.want)
			if err != nil {
				t.Fatalf("encodeMessage: %v", err)
			}
			dst := c.into()
			msg := ReceivedMessage{decoded: dict}
			if err := decodeMessage(msg, dst); err != nil {
				t.Fatalf("decodeMessage: %v", err)
			}
			got := reflect.ValueOf(dst).Elem().Interface()
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("round trip = %#v, want %#v", got, c.want)
			}
		})
	}
}

// TestCodecGoldenFilesAreExhaustive runs both directions: a case with no file
// fails, and a file with no case fails. Deleting a fixture to silence a
// failure therefore leaves an orphan file that fails in its place.
func TestCodecGoldenFilesAreExhaustive(t *testing.T) {
	wantFiles := map[string]bool{}
	for _, c := range codecGoldenCases {
		if c.wantErr {
			continue
		}
		wantFiles[c.name+".golden"] = true
	}
	entries, err := os.ReadDir(codecGoldenDir)
	if err != nil {
		t.Fatalf("read %s: %v", codecGoldenDir, err)
	}
	haveFiles := map[string]bool{}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".golden") {
			continue
		}
		haveFiles[e.Name()] = true
	}
	for name := range wantFiles {
		if !haveFiles[name] {
			t.Errorf("case has no golden file: %s/%s", codecGoldenDir, name)
		}
	}
	for name := range haveFiles {
		if !wantFiles[name] {
			t.Errorf("golden file has no case: %s/%s", codecGoldenDir, name)
		}
	}
}

// TestCreateFlags pins the bit values against <xpc/listener.h> and
// <xpc/session.h>. These options were declared but never passed for long
// enough that setting them did nothing at all.
func TestCreateFlags(t *testing.T) {
	if got := (ListenerOptions{}).flags(); got != 0 {
		t.Errorf("zero ListenerOptions flags = %#x, want 0", got)
	}
	if got := (ListenerOptions{Inactive: true, ForceMach: true}).flags(); got != 0x3 {
		t.Errorf("inactive|mach = %#x, want 0x3", got)
	}
	if got := (ListenerOptions{ForceXPCService: true}).flags(); got != 0x4 {
		t.Errorf("xpcservice = %#x, want 0x4", got)
	}
	if got := (SessionOptions{}).flags(); got != 0 {
		t.Errorf("zero SessionOptions flags = %#x, want 0", got)
	}
	if got := (SessionOptions{Inactive: true, Privileged: true}).flags(); got != 0x3 {
		t.Errorf("inactive|privileged = %#x, want 0x3", got)
	}
}

// TestTypeSymbolsMatchLiveObjects catches a silent decode failure. The XPC
// type objects are exported as underscore-prefixed C symbols and the type
// value is the symbol's ADDRESS, per
//
//	#define XPC_TYPE_INT64 (&_xpc_type_int64)
//
// Looking up the unprefixed name, or dereferencing the address, leaves every
// comparison in rawObjectToValue false, so every value decodes through the
// copy_description fallback as a string. Comparing against the type of a real
// object is what discriminates: merely checking that the lookup is non-nil and
// distinct passes even when the address is dereferenced.
func TestTypeSymbolsMatchLiveObjects(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("XPC framework unavailable")
	}
	tests := []struct {
		symbol string
		object unsafe.Pointer
	}{
		{"xpc_type_bool", raw_xpc_bool_create(true)},
		{"xpc_type_int64", raw_xpc_int64_create(42)},
		{"xpc_type_uint64", raw_xpc_uint64_create(42)},
		{"xpc_type_double", raw_xpc_double_create(1.5)},
		{"xpc_type_string", raw_xpc_string_create("hello")},
		{"xpc_type_dictionary", raw_xpc_dictionary_create_empty()},
	}
	for _, test := range tests {
		t.Run(test.symbol, func(t *testing.T) {
			if test.object == nil {
				t.Fatalf("could not create a %s object", test.symbol)
			}
			defer releaseRaw(test.object)
			want := raw_xpc_get_type(test.object)
			if got := xpcTypeSymbol(test.symbol); got != want {
				t.Errorf("xpcTypeSymbol(%q) = %p, but a live object reports type %p",
					test.symbol, got, want)
			}
		})
	}
}

// TestDecodedValueTypes is the end-to-end consequence: with the type lookup
// wrong, each of these arrives as a description string instead of its Go type.
func TestDecodedValueTypes(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("XPC framework unavailable")
	}
	dict := Dictionary{"i": int64(-5), "u": uint64(5), "f": 1.5, "s": "x", "b": true}
	raw, err := dictionaryToRawObject(dict)
	if err != nil {
		t.Fatalf("dictionaryToRawObject: %v", err)
	}
	defer releaseRaw(raw)
	got, err := rawObjectToDictionary(raw)
	if err != nil {
		t.Fatalf("rawObjectToDictionary: %v", err)
	}
	for key, want := range dict {
		if got[key] != want {
			t.Errorf("%q = %#v (%T), want %#v (%T)", key, got[key], got[key], want, want)
		}
	}
}

// blockCount reports how many callback blocks newXPCBlock has registered.
// Every block is retained for the process's life (purego callbacks cannot be
// unregistered), so this counter only grows and is the direct measure of the
// exhaustible resource: purego's callback table holds a fixed 2000 entries and
// panics when full, from a native callback thread, which aborts the process.
// newXPCBlock appends exactly three keepalive entries per registered callback
// (the Go func, the block literal, and its descriptor), so the division is
// exact rather than a heuristic; the sensitivity step in
// TestApplierBlockSharesOneCallback fails loudly if that stops being true.
func blockCount() int {
	xpcBlockKeepaliveMu.Lock()
	defer xpcBlockKeepaliveMu.Unlock()
	return len(xpcBlockKeepalive) / 3
}

// decodeCount reports how many raw container decodes have run. Every
// rawObjectToDictionary and rawArrayToSlice call takes exactly one applier
// token, and tokens are never reused, so the token counter counts decodes.
// It replaces blockCount as the memo's instrument: decodes no longer register
// callbacks, so blockCount can no longer see one.
func decodeCount() uint64 {
	xpcApplierMu.Lock()
	defer xpcApplierMu.Unlock()
	return xpcApplierNext
}

// liveApplierTokens reports how many applier tokens are currently held. Both
// apply functions are synchronous and release their token before returning, so
// outside a decode this must be zero; a nonzero value at rest is a leaked
// token, the failure mode the token table trades for the callback-table one.
func liveApplierTokens() int {
	xpcApplierMu.Lock()
	defer xpcApplierMu.Unlock()
	return len(xpcApplierState)
}

// replyCount reports how many reply blocks have been built. Tokens are never
// reused, so the counter counts calls made on the cancellable-context path and
// is the denominator for the callback cost of that path.
func replyCount() uint64 {
	xpcReplyMu.Lock()
	defer xpcReplyMu.Unlock()
	return xpcReplyNext
}

// liveReplyTokens reports how many reply tokens are currently held: calls
// whose reply has not yet arrived, plus any whose reply never will. This is
// the resource the callback table was traded for, so it is the one to watch —
// unlike the callback table it does not abort the process when it grows, which
// also means nothing else reports it.
func liveReplyTokens() int {
	xpcReplyMu.Lock()
	defer xpcReplyMu.Unlock()
	return len(xpcReplyState)
}

// TestReplyBlockSharesOneCallback measures the callback cost of the
// cancellable-context call path without needing a peer to answer: building the
// block is where the registration happened, so building many of them is the
// measurement. It is the offline half of the live bracket, which exercises the
// same path against a real service.
func TestReplyBlockSharesOneCallback(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("XPC framework unavailable")
	}

	// Sensitivity (positive control): blockCount must still see a real
	// registration, in units of one. Without it, the zero below could mean
	// the instrument is blind rather than that the path is clean.
	before := blockCount()
	if _, err := newXPCBlock(func(_ uintptr) {}); err != nil {
		t.Fatalf("newXPCBlock: %v", err)
	}
	if spent := blockCount() - before; spent != 1 {
		t.Fatalf("one newXPCBlock registered %d blocks, want 1: blockCount is in the wrong unit", spent)
	}

	// Warm the shared trampoline so its one-off registration is not charged
	// to the loop below.
	before = blockCount()
	xpcReplyTrampoline()
	if spent := blockCount() - before; spent != 0 {
		t.Errorf("the reply trampoline registered %d newXPCBlock entries, want 0: "+
			"it must not go through newXPCBlock", spent)
	}

	const n = 4000 // twice purego's 2000-entry callback table
	before = blockCount()
	beforeCalls := replyCount()
	blocks := make([]*xpcReplyBlock, 0, n)
	for i := 0; i < n; i++ {
		blk, err := newXPCReplyBlock(func(Dictionary, error) {})
		if err != nil {
			t.Fatalf("newXPCReplyBlock %d: %v", i, err)
		}
		blocks = append(blocks, blk)
	}
	if built := replyCount() - beforeCalls; built != n {
		t.Fatalf("built %d reply blocks, want %d: the denominator is wrong", built, n)
	}
	if spent := blockCount() - before; spent != 0 {
		t.Errorf("%d reply blocks cost %d callback-table entries, want 0: the table holds 2000 "+
			"and panics from a native thread when full, which aborts the process", n, spent)
	}

	// The tokens are held until the trampoline runs. Drain them the way a
	// reply does and require the table to come back to where it started:
	// a token that is not freed is the leak this trade introduced.
	live := liveReplyTokens()
	if live < n {
		t.Fatalf("only %d of %d reply tokens are held: they are being dropped before the reply arrives", live, n)
	}
	for _, blk := range blocks {
		if e := xpcReplyTake(blk.token); e == nil {
			t.Fatalf("token %d was already gone", blk.token)
		}
		if e := xpcReplyTake(blk.token); e != nil {
			t.Fatalf("token %d survived being taken: a second invocation would call the Go closure twice", blk.token)
		}
	}
	if got := liveReplyTokens(); got != live-n {
		t.Errorf("after draining %d tokens, %d are held, want %d", n, got, live-n)
	}
}

// TestApplierBlockSharesOneCallback measures the callback cost of decoding.
//
// Before the applier blocks were hoisted, rawObjectToDictionary and
// rawArrayToSlice each called newXPCBlock inside the function body, so every
// decode — and every nested dictionary or array within it — burned one entry of
// purego's fixed 2000-entry callback table. A service that decoded a request
// per message therefore died of callback exhaustion after a few hundred round
// trips. The block literal is now an ordinary Go allocation carrying a token,
// dispatching to one shared callback per block signature.
func TestApplierBlockSharesOneCallback(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("XPC framework unavailable")
	}

	// Sensitivity (positive control): blockCount must still see a real
	// callback registration, in units of one. Without this the zero measured
	// below could mean the instrument is blind.
	before := blockCount()
	if _, err := newXPCBlock(func(_ uintptr) {}); err != nil {
		t.Fatalf("newXPCBlock: %v", err)
	}
	if spent := blockCount() - before; spent != 1 {
		t.Fatalf("one newXPCBlock registered %d blocks, want 1: blockCount is in the wrong unit", spent)
	}

	// Warm the shared trampolines so their one-off registration is not
	// attributed to the loop below. Each is created at most once per process.
	before = blockCount()
	xpcDictApplier()
	xpcArrayApplier()
	if spent := blockCount() - before; spent != 0 {
		t.Errorf("the shared trampolines registered %d newXPCBlock entries, want 0: "+
			"they must not go through newXPCBlock", spent)
	}

	// A message with a nested dictionary and a nested array: under the old
	// code one decode of this cost three callbacks, not one.
	nested := Dictionary{
		"op":    "add",
		"inner": Dictionary{"a": int64(1), "b": "two"},
		"list":  []any{int64(1), "x", true},
	}
	raw, err := dictionaryToRawObject(nested)
	if err != nil {
		t.Fatalf("dictionaryToRawObject: %v", err)
	}
	defer releaseRaw(raw)

	const decodes = 500
	beforeBlocks := blockCount()
	beforeDecodes := decodeCount()
	for i := 0; i < decodes; i++ {
		got, err := rawObjectToDictionary(raw)
		if err != nil {
			t.Fatalf("decode %d: %v", i, err)
		}
		// Verify the decode really happened, through the shared
		// trampoline, including both nested container kinds. A decode
		// that silently produced nothing would also register no
		// callbacks.
		if got["op"] != "add" {
			t.Fatalf("decode %d: op = %#v, want \"add\"", i, got["op"])
		}
		inner, ok := got["inner"].(Dictionary)
		if !ok || inner["b"] != "two" {
			t.Fatalf("decode %d: inner = %#v", i, got["inner"])
		}
		list, ok := got["list"].([]any)
		if !ok || len(list) != 3 || list[1] != "x" {
			t.Fatalf("decode %d: list = %#v", i, got["list"])
		}
	}
	if spent := blockCount() - beforeBlocks; spent != 0 {
		t.Errorf("%d decodes registered %d callback blocks, want 0 (%.2f per decode): "+
			"the appliers are still per-call", decodes, spent, float64(spent)/decodes)
	}
	// Three containers per decode: the outer dictionary, the inner
	// dictionary, and the array. This is the count that used to be the
	// callback count.
	if got, want := decodeCount()-beforeDecodes, uint64(decodes*3); got != want {
		t.Errorf("%d decodes consumed %d applier tokens, want %d", decodes, got, want)
	}
	if live := liveApplierTokens(); live != 0 {
		t.Errorf("%d applier tokens still held after all decodes returned, want 0: tokens are leaking", live)
	}
}

// TestReceivedMessageDictionaryMemo measures the memo instead of asserting it:
// it counts raw decodes across repeated Dictionary calls on one raw-backed
// message.
//
// The instrument is decodeCount, not blockCount: since the appliers were
// hoisted onto a shared trampoline a decode registers no callback, so
// blockCount is blind to one. See TestApplierBlockSharesOneCallback.
//
// The unmemoized arm is the negative control. It is the same message with a nil
// memo, and it must consume one decode per call; without it, a Dictionary that
// had stopped decoding entirely would also make the memoized arm pass.
func TestReceivedMessageDictionaryMemo(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("XPC framework unavailable")
	}
	want := Dictionary{"op": "add", "n": int64(7)}
	raw, err := dictionaryToRawObject(want)
	if err != nil {
		t.Fatalf("dictionaryToRawObject: %v", err)
	}
	defer releaseRaw(raw)

	const calls = 8

	// Calibration: one decode of this flat dictionary costs exactly one
	// applier token. If it does not, every count below is in the wrong unit
	// and the verdict would be about the instrument.
	before := decodeCount()
	calibration := ReceivedMessage{raw: raw}
	if got := calibration.Dictionary(); got["op"] != "add" {
		t.Fatalf("calibration decode = %#v, want %#v", got, want)
	}
	if perDecode := decodeCount() - before; perDecode != 1 {
		t.Fatalf("one decode consumed %d applier tokens, want 1: decodeCount is in the wrong unit", perDecode)
	}

	memoized := ReceivedMessage{raw: raw, memo: &messageMemo{}}
	before = decodeCount()
	for i := 0; i < calls; i++ {
		got := memoized.Dictionary()
		if got["op"] != "add" || got["n"] != int64(7) {
			t.Fatalf("call %d decoded %#v, want %#v", i, got, want)
		}
		got["op"] = "mutated by the caller"
	}
	if spent := decodeCount() - before; spent != 1 {
		t.Errorf("%d Dictionary calls on a memoized message ran %d decodes, want 1", calls, spent)
	}

	// The copy shares the memo, which is why memo is a pointer: value methods
	// mean a ReceivedMessage is copied on every call and every assignment.
	before = decodeCount()
	copied := memoized
	if got := copied.Dictionary(); got["op"] != "add" {
		t.Errorf("a copy of a memoized message decoded %#v; the cached body was not shared or was mutated", got)
	}
	if spent := decodeCount() - before; spent != 0 {
		t.Errorf("a copy of a memoized message ran %d decodes, want 0", spent)
	}

	unmemoized := ReceivedMessage{raw: raw}
	before = decodeCount()
	for i := 0; i < calls; i++ {
		if got := unmemoized.Dictionary(); got["op"] != "add" {
			t.Fatalf("unmemoized call %d decoded %#v, want %#v", i, got, want)
		}
	}
	if spent := decodeCount() - before; spent != calls {
		t.Errorf("%d Dictionary calls with no memo ran %d decodes, want %d; "+
			"the counter is not sensitive to a decode, so the memoized arm above proves nothing",
			calls, spent, calls)
	}
}

// requireSymbol registers a synthetic "missing symbol" so the tests below can
// prove that argument validation runs before symbol lookup on runtimes that
// lack a symbol. The previous entry, if any, is restored.
func requireSymbol(symbol string) func() {
	old, had := rawSymbolLookupErrors[symbol]
	rawSymbolLookupErrors[symbol] = errors.New("synthetic missing symbol for test")
	return func() {
		if had {
			rawSymbolLookupErrors[symbol] = old
		} else {
			delete(rawSymbolLookupErrors, symbol)
		}
	}
}

// TestPeerRequirementConstructorValidation covers the argument checks that run
// before any native call: an empty signingIdentifier in either SignedAs
// constructor, an empty entitlement, a nil lwcr, and an unsupported
// entitlement value, each named in the error.
func TestPeerRequirementConstructorValidation(t *testing.T) {
	if _, err := NewSameTeamSignedAsRequirement(""); err == nil || !strings.Contains(err.Error(), "signingIdentifier") {
		t.Errorf("NewSameTeamSignedAsRequirement(\"\") = %v, want an error naming signingIdentifier", err)
	}
	if _, err := NewPlatformBinarySignedAsRequirement(""); err == nil || !strings.Contains(err.Error(), "signingIdentifier") {
		t.Errorf("NewPlatformBinarySignedAsRequirement(\"\") = %v, want an error naming signingIdentifier", err)
	}
	if _, err := NewEntitlementExistsRequirement(""); err == nil || !strings.Contains(err.Error(), "entitlement") {
		t.Errorf("NewEntitlementExistsRequirement(\"\") = %v, want an error naming entitlement", err)
	}
	if _, err := NewEntitlementMatchesRequirement("", "x"); err == nil || !strings.Contains(err.Error(), "entitlement") {
		t.Errorf("NewEntitlementMatchesRequirement(\"\", \"x\") = %v, want an error naming entitlement", err)
	}
	if _, err := NewEntitlementMatchesRequirement("com.example.test", 1.5); err == nil || !strings.Contains(err.Error(), "float64") {
		t.Errorf("NewEntitlementMatchesRequirement(..., 1.5) = %v, want an error naming float64", err)
	}
	if _, err := NewLightweightCodeRequirement(nil); err == nil || !strings.Contains(err.Error(), "lwcr") {
		t.Errorf("NewLightweightCodeRequirement(nil) = %v, want an error naming lwcr", err)
	}
}

// TestEntitlementValueConversion pins the dedicated conversion for
// entitlement_matches_value: exactly bool, string, and signed integers that
// int64 represents exactly, everything else rejected by name.
func TestEntitlementValueConversion(t *testing.T) {
	type namedInt int64
	type namedBool bool
	for _, value := range []any{
		true, "x",
		int(1), int8(1), int16(1), int32(1), int64(1),
		namedInt(1), namedBool(true),
	} {
		obj, err := entitlementValueToRawObject(value)
		if err != nil {
			t.Errorf("entitlementValueToRawObject(%T) rejected: %v", value, err)
			continue
		}
		if obj == nil && frameworkHandle != 0 {
			t.Errorf("entitlementValueToRawObject(%T) = nil object", value)
		}
		releaseRaw(obj)
	}
	for _, value := range []any{
		nil, uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		float32(1.5), 1.5, []byte("x"), []int{1}, Dictionary{},
	} {
		obj, err := entitlementValueToRawObject(value)
		if err == nil {
			releaseRaw(obj)
			t.Errorf("entitlementValueToRawObject(%T) accepted, want rejection", value)
			continue
		}
		if !strings.Contains(err.Error(), fmt.Sprintf("%T", value)) {
			t.Errorf("entitlementValueToRawObject(%T) error %q does not name the type", value, err)
		}
	}
}

// TestPeerRequirementValidationBeforeSymbolLookup proves a bad argument
// reports the argument error even when the target symbol is missing, while a
// valid argument on the same runtime reports the unavailable-symbol error.
func TestPeerRequirementValidationBeforeSymbolLookup(t *testing.T) {
	restore := requireSymbol("xpc_peer_requirement_create_entitlement_exists")
	defer restore()

	if _, err := NewEntitlementExistsRequirement(""); err == nil || !strings.Contains(err.Error(), "entitlement") {
		t.Errorf("empty entitlement with missing symbol = %v, want the argument error", err)
	}
	if _, err := NewEntitlementExistsRequirement("com.example.test"); err == nil || !strings.Contains(err.Error(), "unavailable") {
		t.Errorf("valid entitlement with missing symbol = %v, want the unavailable-symbol error", err)
	}

	if _, err := NewSameTeamSignedAsRequirement(""); err == nil || !strings.Contains(err.Error(), "signingIdentifier") {
		t.Errorf("empty signingIdentifier with missing symbol = %v, want the argument error", err)
	}
	restore2 := requireSymbol("xpc_peer_requirement_create_team_identity")
	defer restore2()
	if _, err := NewSameTeamSignedAsRequirement("com.example.signed"); err == nil || !strings.Contains(err.Error(), "unavailable") {
		t.Errorf("valid signingIdentifier with missing symbol = %v, want the unavailable-symbol error", err)
	}
}

// TestReceivedMessageMethodSet pins the public method set of ReceivedMessage.
// One reply per received message is structural: the only reply path is the
// unexported sendReplyDictionary, called once by the block in
// SetIncomingMessageHandler. Re-adding a public Reply would restore the
// two-replies hazard, so it must break a test loudly rather than compile.
func TestReceivedMessageMethodSet(t *testing.T) {
	typ := reflect.TypeOf(ReceivedMessage{})
	var got []string
	for i := 0; i < typ.NumMethod(); i++ {
		got = append(got, typ.Method(i).Name)
	}
	sort.Strings(got)
	want := []string{"Decode", "Dictionary", "Handle", "SenderSatisfies"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReceivedMessage methods = %v, want %v", got, want)
	}
}

// TestCallRejectsNilSession and TestNotifyRejectsNilSession cover the nil
// guards on the send surface, which no test reached before.
func TestCallRejectsNilSession(t *testing.T) {
	var s *Session
	if _, err := s.Call(context.Background(), Dictionary{}); err == nil || !strings.Contains(err.Error(), "session is nil") {
		t.Errorf("Call on a nil session = %v, want the nil-session error", err)
	}
	if _, err := s.CallDictionary(context.Background(), Dictionary{}); err == nil || !strings.Contains(err.Error(), "session is nil") {
		t.Errorf("CallDictionary on a nil session = %v, want the nil-session error", err)
	}
}

func TestNotifyRejectsNilSession(t *testing.T) {
	var s *Session
	if err := s.Notify(Dictionary{}); err == nil || !strings.Contains(err.Error(), "session is nil") {
		t.Errorf("Notify on a nil session = %v, want the nil-session error", err)
	}
	if err := s.NotifyDictionary(Dictionary{}); err == nil || !strings.Contains(err.Error(), "session is nil") {
		t.Errorf("NotifyDictionary on a nil session = %v, want the nil-session error", err)
	}
}

// TestCallHonoursExpiredContext proves ctx is checked before any raw symbol is
// looked up: with the send symbol poisoned, an already-cancelled context still
// reports the context error rather than the unavailable-symbol error.
func TestCallHonoursExpiredContext(t *testing.T) {
	restore := requireSymbol("xpc_session_send_message_with_reply_sync")
	defer restore()

	s := &Session{}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := s.CallDictionary(ctx, Dictionary{}); !errors.Is(err, context.Canceled) {
		t.Errorf("CallDictionary with a cancelled context = %v, want context.Canceled", err)
	}
	if _, err := s.CallDictionary(context.Background(), Dictionary{}); err == nil || !strings.Contains(err.Error(), "unavailable") {
		t.Errorf("CallDictionary with a live context = %v, want the unavailable-symbol error", err)
	}
}

// TestOptionsRequirementValidationBeforeSymbolLookup proves the options paths
// validate a nonnil requirement before any create symbol is looked up.
func TestOptionsRequirementValidationBeforeSymbolLookup(t *testing.T) {
	restore := requireSymbol("xpc_session_create_mach_service")
	defer restore()

	if _, err := DialMachService("dev.tmc.apple.xpc.test", SessionOptions{Requirement: &PeerRequirement{}}); err == nil || !strings.Contains(err.Error(), "peer requirement") {
		t.Errorf("closed requirement with missing symbol = %v, want the requirement error", err)
	}
	if _, err := DialMachService("dev.tmc.apple.xpc.test", SessionOptions{}); err == nil || !strings.Contains(err.Error(), "unavailable") {
		t.Errorf("nil requirement with missing symbol = %v, want the unavailable-symbol error", err)
	}
}

// TestPeerRequirementFromHandleOwnership covers the owned-reference contract:
// zero yields nil, a nonzero handle is retained and preserved, Close is
// idempotent, and Handle reads zero after Close.
func TestPeerRequirementFromHandleOwnership(t *testing.T) {
	if got := PeerRequirementFromHandle(0); got != nil {
		t.Fatalf("PeerRequirementFromHandle(0) = %p, want nil", got)
	}

	req, err := NewEntitlementExistsRequirement("com.example.test")
	if err != nil {
		if strings.Contains(err.Error(), "unavailable") || strings.Contains(err.Error(), "framework") {
			t.Skipf("requirement constructors unavailable: %v", err)
		}
		t.Fatalf("NewEntitlementExistsRequirement: %v", err)
	}
	handle := req.Handle()
	if handle == 0 {
		t.Fatal("fresh requirement has a zero handle")
	}

	retained := PeerRequirementFromHandle(handle)
	if retained == nil || retained.Handle() != handle {
		t.Fatalf("PeerRequirementFromHandle(%#x) = %v, want the same handle", handle, retained)
	}
	// The retained reference is independent: closing the original leaves it
	// holding its own reference, and both Closes are nil.
	if err := req.Close(); err != nil {
		t.Fatalf("Close original: %v", err)
	}
	if got := req.Handle(); got != 0 {
		t.Errorf("Handle after Close = %#x, want 0", got)
	}
	if err := req.Close(); err != nil {
		t.Errorf("second Close: %v, want nil", err)
	}
	if retained.Handle() != handle {
		t.Errorf("retained handle = %#x after original Close, want %#x", retained.Handle(), handle)
	}
	if err := retained.Close(); err != nil {
		t.Fatalf("Close retained: %v", err)
	}
	if got := retained.Handle(); got != 0 {
		t.Errorf("retained Handle after Close = %#x, want 0", got)
	}
	if err := retained.Close(); err != nil {
		t.Errorf("second Close on retained: %v, want nil", err)
	}
}

// TestNewRequirementNilRaw covers newRequirement's nil-result handling: with
// and without an accompanying rich error.
func TestNewRequirementNilRaw(t *testing.T) {
	if rawSymbolError("xpc_peer_requirement_create_team_identity") != nil {
		t.Skip("requirement constructor symbol unavailable")
	}
	if _, err := newRequirement(func(richErr *unsafe.Pointer) unsafe.Pointer {
		return nil
	}, "xpc_peer_requirement_create_team_identity"); err == nil || err.Error() != "xpc: failed to create peer requirement" {
		t.Errorf("nil raw without rich error = %v, want the generic failure", err)
	}

	// A real rich error: creating an active session to a name nothing vends
	// fails at creation with error_out set, which is a deterministic source
	// of a genuine xpc_rich_error_t.
	if rawSymbolError("xpc_session_create_mach_service") != nil {
		t.Skip("xpc_session_create_mach_service unavailable")
	}
	var richErr unsafe.Pointer
	if raw := raw_xpc_session_create_mach_service(
		fmt.Sprintf("dev.tmc.apple.xpc.absent.%d", os.Getpid()), nil, 0, &richErr); raw != nil {
		releaseRaw(raw)
		t.Skip("active session create to an absent service unexpectedly succeeded")
	}
	if richErr == nil {
		t.Skip("active session create produced no rich error")
	}
	if _, err := newRequirement(func(out *unsafe.Pointer) unsafe.Pointer {
		*out = richErr
		return nil
	}, "xpc_peer_requirement_create_team_identity"); err == nil {
		t.Fatal("nil raw with rich error = nil, want a RichError")
	} else {
		var rich RichError
		if !errors.As(err, &rich) {
			t.Errorf("nil raw with rich error = %T %v, want a RichError", err, err)
		}
	}
}

// TestOptionsRequirementNilAndClosed covers the options paths: a nil
// requirement is the valid "no restriction" sentinel, and a nonnil zero or
// closed requirement is refused by both ListenerOptions and SessionOptions.
func TestOptionsRequirementNilAndClosed(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("XPC framework unavailable")
	}
	name := fmt.Sprintf("dev.tmc.apple.xpc.ordinary.%d", os.Getpid())
	accept := func(req IncomingSessionRequest) IncomingDecision { return req.Reject("no") }

	// nil requirement: both paths must proceed past validation. Session
	// creation to a name nothing vends is lazy when inactive, so an inactive
	// dial succeeds; the listener create is likewise lazy. Either way, the
	// error (if any) must be a native one, never a requirement validation
	// error. Neither object is cancelled: xpc_session_cancel requires an
	// activated session, and these are never activated.
	session, err := DialMachService(name, SessionOptions{Inactive: true})
	_ = session
	if err != nil {
		t.Errorf("nil requirement dial refused: %v", err)
	}
	listener, err := NewServiceListener(name, ListenerOptions{}, accept)
	if err != nil {
		if strings.Contains(err.Error(), "peer requirement") {
			t.Errorf("nil requirement listener refused by validation: %v", err)
		}
	} else {
		listener.Cancel()
	}

	// nonnil zero or closed requirement: refused before any native create.
	if _, err := DialMachService(name, SessionOptions{Requirement: &PeerRequirement{}}); err == nil || !strings.Contains(err.Error(), "peer requirement") {
		t.Errorf("zero requirement dial = %v, want a requirement error", err)
	}
	if _, err := NewServiceListener(name, ListenerOptions{Requirement: &PeerRequirement{}}, accept); err == nil || !strings.Contains(err.Error(), "peer requirement") {
		t.Errorf("zero requirement listener = %v, want a requirement error", err)
	}
}

// TestSetPeerRequirementLifecycle proves the inactive-only contract: nil,
// zero, and closed requirements are refused; an active session and a second
// installation are refused before native code runs; a handle-derived session
// is refused. The guard cases inject a missing xpc_session_set_peer_requirement
// symbol so a removed guard surfaces as the unavailable-symbol error instead
// of reaching native code: the assertions below then fail cleanly rather than
// trapping the test binary, which is the point of the canary.
func TestSetPeerRequirementLifecycle(t *testing.T) {
	session := &Session{raw: pointerFromHandle(0x1234)}

	if err := session.SetPeerRequirement(nil); err == nil || !strings.Contains(err.Error(), "nil") {
		t.Errorf("SetPeerRequirement(nil) = %v, want a nil-requirement error", err)
	}
	if err := session.SetPeerRequirement(&PeerRequirement{}); err == nil || !strings.Contains(err.Error(), "closed") {
		t.Errorf("SetPeerRequirement(zero) = %v, want a closed-requirement error", err)
	}

	// A real closed requirement, when one can be built.
	if closed, cerr := NewEntitlementExistsRequirement("com.example.closed"); cerr == nil {
		if err := closed.Close(); err != nil {
			t.Fatalf("Close: %v", err)
		}
		if err := session.SetPeerRequirement(closed); err == nil || !strings.Contains(err.Error(), "closed") {
			t.Errorf("SetPeerRequirement(closed) = %v, want a closed-requirement error", err)
		}
	}

	// An open requirement for the refusals that occur after the requirement
	// check itself: active session, duplicate installation, handle-derived
	// session. The injected missing symbol makes each refusal prove it fired
	// before native code: with the guard removed, the call would return the
	// unavailable-symbol error instead and the assertion would fail without
	// trapping.
	req, err := NewEntitlementExistsRequirement("com.example.test")
	if err != nil {
		if strings.Contains(err.Error(), "unavailable") || strings.Contains(err.Error(), "framework") {
			req = &PeerRequirement{raw: pointerFromHandle(0x5678)}
		} else {
			t.Fatalf("NewEntitlementExistsRequirement: %v", err)
		}
	}
	defer req.Close()
	restore := requireSymbol("xpc_session_set_peer_requirement")
	defer restore()

	active := &Session{raw: pointerFromHandle(0x1234), active: true}
	if err := active.SetPeerRequirement(req); err == nil || !strings.Contains(err.Error(), "inactive") {
		t.Errorf("SetPeerRequirement on active session = %v, want an inactive-only error", err)
	}
	double := &Session{raw: pointerFromHandle(0x1234), requirementSet: true}
	if err := double.SetPeerRequirement(req); err == nil || !strings.Contains(err.Error(), "already set") {
		t.Errorf("second SetPeerRequirement = %v, want an already-set error", err)
	}
	derived := SessionFromHandle(0x1234)
	if err := derived.SetPeerRequirement(req); err == nil || !strings.Contains(err.Error(), "handle-derived") {
		t.Errorf("SetPeerRequirement on handle-derived session = %v, want a handle-derived error", err)
	}
}

// TestSetTargetQueueLifecycle covers Change 2's session method: zero and
// repeated replacement are accepted while inactive, and active and
// handle-derived sessions are refused before native code runs. As in
// TestSetPeerRequirementLifecycle, the guard cases inject a missing symbol so
// a removed guard fails the assertion instead of trapping.
func TestSetTargetQueueLifecycle(t *testing.T) {
	// The refusals fire before native code, so fake sessions suffice.
	restore := requireSymbol("xpc_session_set_target_queue")
	defer restore()
	active := &Session{raw: pointerFromHandle(0x1234), active: true}
	if err := active.SetTargetQueue(dispatch.Queue{}); err == nil || !strings.Contains(err.Error(), "inactive") {
		t.Errorf("SetTargetQueue on active session = %v, want an inactive-only error", err)
	}
	derived := SessionFromHandle(0x1234)
	if err := derived.SetTargetQueue(dispatch.Queue{}); err == nil || !strings.Contains(err.Error(), "handle-derived") {
		t.Errorf("SetTargetQueue on handle-derived session = %v, want a handle-derived error", err)
	}
	restore()

	// Acceptance needs a real inactive session, which dialing a name nothing
	// vends provides: inactive creation is lazy and does not connect until
	// activation, so it succeeds where an active create fails eagerly. The
	// session cannot be activated (its service does not exist), so the
	// post-activation refusals are covered by the fake active session above
	// and by TestSetPeerRequirementLifecycle.
	if frameworkHandle == 0 {
		t.Skip("XPC framework unavailable")
	}
	if rawSymbolError("xpc_session_set_target_queue") != nil {
		t.Skip("xpc_session_set_target_queue unavailable")
	}
	name := fmt.Sprintf("dev.tmc.apple.xpc.ordinary.%d", os.Getpid())
	session, err := DialMachService(name, SessionOptions{Inactive: true})
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	// Not cancelled: xpc_session_cancel requires an activated session, and
	// this one is never activated. Dropping it releases the reference.

	queue := dispatch.QueueCreate(fmt.Sprintf("xpc.target.%d", os.Getpid()))
	if err := session.SetTargetQueue(dispatch.Queue{}); err != nil {
		t.Errorf("SetTargetQueue(zero) = %v, want nil", err)
	}
	if err := session.SetTargetQueue(queue); err != nil {
		t.Errorf("SetTargetQueue(queue) = %v, want nil", err)
	}
	if err := session.SetTargetQueue(dispatch.Queue{}); err != nil {
		t.Errorf("repeated SetTargetQueue(zero) = %v, want nil", err)
	}
}

// --- raw-symbol reachability ---------------------------------------------
//
// rawReachability, rawAllCalledSymbols and rawReachUnresolved are emitted by
// the generator into xpc.rawreach.gen.go. The tests below are the gates that
// keep those sets honest: coverage of every raw call, an allowlist for every
// call edge the analysis could not follow, and a canary that proves the
// availability guards actually fire.

var errPoisonedCanary = errors.New("xpc: poisoned by canary")

// poisonRawSymbol makes name look unavailable for the duration of the test.
// rawSymbolLookupErrors is process-global, so no test in this group may call
// t.Parallel.
func poisonRawSymbol(t *testing.T, name string) {
	t.Helper()
	prev, had := rawSymbolLookupErrors[name]
	rawSymbolLookupErrors[name] = errPoisonedCanary
	t.Cleanup(func() {
		if had {
			rawSymbolLookupErrors[name] = prev
		} else {
			delete(rawSymbolLookupErrors, name)
		}
	})
}

// rawReachUnion is the union of every entry point's reachable symbol set.
func rawReachUnion() map[string]bool {
	union := map[string]bool{}
	for _, syms := range rawReachability {
		for _, s := range syms {
			union[s] = true
		}
	}
	return union
}

// rawReachCoverageGaps returns the symbols in called that no entry point
// reaches. It is shared by the coverage gate and its mutation control so both
// measure with the same instrument.
func rawReachCoverageGaps(called []string) []string {
	union := rawReachUnion()
	var missing []string
	for _, s := range called {
		if !union[s] {
			missing = append(missing, s)
		}
	}
	sort.Strings(missing)
	return missing
}

func TestRawReachCoversEveryRawCall(t *testing.T) {
	if len(rawAllCalledSymbols) == 0 {
		t.Fatal("rawAllCalledSymbols is empty: the analysis measured nothing, so a zero here would come from not looking")
	}
	missing := rawReachCoverageGaps(rawAllCalledSymbols)
	t.Logf("raw reach coverage: %d symbols called, %d entry points, %d reachable, %d unreachable",
		len(rawAllCalledSymbols), len(rawReachability), len(rawReachUnion()), len(missing))
	for _, s := range missing {
		t.Errorf("raw symbol %s is called but is reachable from no exported entry point: its availability guard, if any, is not derived from the emitted set", s)
	}
}

func TestRawReachCoverageIsSensitive(t *testing.T) {
	fabricated := append(append([]string(nil), rawAllCalledSymbols...), "xpc_symbol_that_does_not_exist")
	missing := rawReachCoverageGaps(fabricated)
	if len(missing) == 0 {
		t.Fatalf("mutation control failed: a fabricated symbol was reported as covered (compared %d symbols)", len(fabricated))
	}
	t.Logf("mutation control: compared %d symbols, fabricated symbol correctly reported missing", len(fabricated))
}

func TestRawReachUnresolvedIsAllowlisted(t *testing.T) {
	t.Logf("raw reach: %d call edge(s) the generator could not follow", len(rawReachUnresolved))
	for _, e := range rawReachUnresolved {
		key := e.Entry + "|" + e.Expr
		if _, ok := rawReachAllowed[key]; !ok {
			t.Errorf("unfollowable call edge %s (%s:%d, %s) is not in rawReachAllowed: the reachable set for %s may be an UNDER-estimate; write down why that is acceptable", key, e.File, e.Line, e.Kind, e.Entry)
		}
	}
}

func TestRawReachEveryEntryPointHasACanary(t *testing.T) {
	seen := map[string]int{}
	for _, c := range canaries {
		seen[c.entry]++
	}
	proven, unmeasured := 0, 0
	for entry := range rawReachability {
		switch {
		case seen[entry] == 1:
			proven++
		case seen[entry] > 1:
			t.Errorf("entry point %s has %d canaries, want exactly 1", entry, seen[entry])
		case canaryUnmeasured[entry] != "":
			unmeasured++
		default:
			t.Errorf("entry point %s has no canary and no recorded reason: its availability guard is UNMEASURED and nobody has written down why", entry)
		}
	}
	for _, c := range canaries {
		if _, ok := rawReachability[c.entry]; !ok {
			t.Errorf("canary names entry point %s, which the generator does not emit", c.entry)
		}
	}
	t.Logf("canary coverage: %d entry points, %d canaries, %d UNMEASURED (recorded)", len(rawReachability), proven, unmeasured)
}

func TestPoisonedSymbolCanary(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("XPC framework not loaded: requireRawSymbols short-circuits before any per-symbol check")
	}
	if len(canaries) == 0 {
		t.Fatal("no canaries: a green run here would prove nothing")
	}
	fired := 0
	for _, c := range canaries {
		t.Run(c.entry, func(t *testing.T) {
			syms, ok := rawReachability[c.entry]
			if !ok {
				t.Fatalf("no reachable set emitted for %s", c.entry)
			}
			found := false
			for _, s := range syms {
				if s == c.poison {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("poison symbol %q is not in the reachable set for %s: a typo would make this row pass vacuously", c.poison, c.entry)
			}
			// Negative control: unpoisoned, the canary must not report the
			// poison error. Otherwise the positive result below proves nothing.
			if err := c.call(t); errors.Is(err, errPoisonedCanary) {
				t.Fatalf("negative control failed: %s reported the poison error while unpoisoned", c.entry)
			}
			poisonRawSymbol(t, c.poison)
			err := c.call(t)
			if !errors.Is(err, errPoisonedCanary) {
				t.Fatalf("guard did not fire for %s: the availability check on %s is dead (got %v)", c.entry, c.poison, err)
			}
			fired++
		})
	}
	t.Logf("poisoned-symbol canary: %d/%d rows fired", fired, len(canaries))
}

// exportedPackageNames parses the package directory (excluding test files) and
// returns the exported top-level func, type, and method names. Methods are
// reported as "Type.Method".
func exportedPackageNames(t *testing.T) map[string]bool {
	t.Helper()
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", func(fi os.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, 0)
	if err != nil {
		t.Fatalf("parse package dir: %v", err)
	}
	names := make(map[string]bool)
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				switch d := decl.(type) {
				case *ast.FuncDecl:
					if !d.Name.IsExported() {
						continue
					}
					if d.Recv == nil || len(d.Recv.List) == 0 {
						names[d.Name.Name] = true
						continue
					}
					recv := d.Recv.List[0].Type
					if star, ok := recv.(*ast.StarExpr); ok {
						recv = star.X
					}
					if ident, ok := recv.(*ast.Ident); ok {
						names[ident.Name+"."+d.Name.Name] = true
					}
				case *ast.GenDecl:
					for _, spec := range d.Specs {
						switch s := spec.(type) {
						case *ast.TypeSpec:
							if s.Name.IsExported() {
								names[s.Name.Name] = true
							}
						case *ast.ValueSpec:
							for _, n := range s.Names {
								if n.IsExported() {
									names[n.Name] = true
								}
							}
						}
					}
				}
			}
		}
	}
	return names
}

// TestSwiftOmissionLedgerMatchesExportedSurface checks the generated Swift
// omission ledger against the package's actual exported surface, in both
// directions: an omitted member must not be shipped, and every ledger entry
// must have a deliberate Go-name mapping in the table below.
//
// An empty ledger is a hard failure. That is the assertion that catches a regen
// run which had no framework documentation (stale parsed cache), a state in
// which every other gate in this package still passes.
func TestSwiftOmissionLedgerMatchesExportedSurface(t *testing.T) {
	if len(xpcSwiftOmissions) == 0 {
		t.Fatal("swift omission ledger is empty; regen ran without framework docs (stale parsed cache) — see U19")
	}

	// The Go name each omitted member would carry if we shipped it. Keys are
	// ledger identifiers with any disambiguating "-hash" suffix removed.
	want := map[string]string{
		"doc://com.apple.xpc/documentation/XPC/XPCListener/InitializationOptions":                                                         "ListenerInitializationOptions",
		"doc://com.apple.xpc/documentation/XPC/XPCListener/InitializationOptions/inactive":                                                "ListenerInitializationOptionsInactive",
		"doc://com.apple.xpc/documentation/XPC/XPCListener/InitializationOptions/none":                                                    "ListenerInitializationOptionsNone",
		"doc://com.apple.xpc/documentation/XPC/XPCListener/endpoint":                                                                      "Listener.Endpoint",
		"doc://com.apple.xpc/documentation/XPC/XPCReceivedMessage/expectsReply":                                                           "ReceivedMessage.ExpectsReply",
		"doc://com.apple.xpc/documentation/XPC/XPCReceivedMessage/handoffReply(to:_:)":                                                    "ReceivedMessage.HandoffReply",
		"doc://com.apple.xpc/documentation/XPC/XPCReceivedMessage/isSync":                                                                 "ReceivedMessage.IsSync",
		"doc://com.apple.xpc/documentation/XPC/XPCSession/init(endpoint:targetQueue:options:cancellationHandler:)":                        "NewSessionFromEndpoint",
		"doc://com.apple.xpc/documentation/XPC/XPCSession/init(endpoint:targetQueue:options:incomingMessageHandler:cancellationHandler:)": "NewSessionFromEndpoint",
	}

	exported := exportedPackageNames(t)

	// Mutation control for the AST scan: without it, a scan that returned zero
	// names would pass the loop below vacuously by never finding anything.
	mustExist := []string{
		"Session.Cancel",
		"Listener.Activate",
		"ReceivedMessage.Decode",
		"Session.SetTargetQueue",
	}
	for _, name := range mustExist {
		if !exported[name] {
			t.Fatalf("AST scan did not find %q, which is defined in this package: the scan is broken, so every ledger check below would pass vacuously (scan found %d exported names)", name, len(exported))
		}
	}

	for _, om := range xpcSwiftOmissions {
		id := om.Identifier
		if i := strings.LastIndex(id, ")"); i >= 0 {
			id = id[:i+1]
		}
		goName, ok := want[id]
		if !ok {
			t.Errorf("omission %q has no Go-name mapping in this test's table; a new omission must be mapped deliberately so the table cannot fall behind classifyMember", om.Identifier)
			continue
		}
		if exported[goName] {
			t.Errorf("omission %q claims %s is not shipped, but %s is exported by this package; the ledger is stale", om.Identifier, goName, goName)
		}
	}
	t.Logf("ledger: %d omissions checked against %d exported names", len(xpcSwiftOmissions), len(exported))
}
