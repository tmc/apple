// Copyright 2026 The apple Authors.

//go:build xpclive

// Adjudication of two behaviours that testdata/codec golden files pin without
// judging: struct_nested.golden (structToDictionary does not recurse) and
// json_marshaler.golden (the types a number takes on the json.Marshaler path).
// A golden file records what the code does; only a real connection says what
// that costs.
//
// Both tests use the service's "mirror" op, which replies with the dictionary
// that actually arrived. The reply is therefore evidence about the wire, not
// about the client's own encoder agreeing with itself.
package xpc_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/tmc/apple/xpc"
)

// nestedInner is the shape struct_nested.golden pins: a struct field whose
// value is another struct.
type nestedInner struct {
	Name  string `xpc:"in"`
	Count int64  `xpc:"n"`
}

type nestedOuter struct {
	Op    string      `xpc:"op"`
	Inner nestedInner `xpc:"inner"`
}

// TestLiveNestedStructIsRejectedNotCorrupted determines what libxpc does with
// the raw Go struct value that structToDictionary hands to
// writeDictionaryToRaw.
func TestLiveNestedStructIsRejectedNotCorrupted(t *testing.T) {
	session := dial(t, serviceName)

	// Positive control: the same nesting expressed as a Dictionary crosses the
	// wire intact, so a failure below is about the struct path and not about
	// nesting.
	msg, err := session.CallDictionary(context.Background(), xpc.Dictionary{
		"op":    "mirror",
		"inner": xpc.Dictionary{"in": "six", "n": int64(6)},
	})
	if err != nil {
		t.Fatalf("control: nested Dictionary: %v", err)
	}
	inner, ok := msg.Dictionary()["inner"].(xpc.Dictionary)
	if !ok {
		t.Fatalf("control: inner = %#v, want a Dictionary: nesting itself does not work, so nothing below is about structs", msg.Dictionary()["inner"])
	}
	if inner["in"] != "six" || inner["n"] != int64(6) {
		t.Fatalf("control: inner = %#v, want {in:six n:6}", inner)
	}

	// The behaviour under adjudication.
	_, err = session.Call(context.Background(), nestedOuter{Op: "mirror", Inner: nestedInner{Name: "six", Count: 6}})
	if err == nil {
		t.Fatalf("Call with a nested struct field succeeded; the golden file says the raw struct reaches writeDictionaryToRaw, so it either crossed the wire or was silently dropped")
	}
	if !strings.Contains(err.Error(), "unsupported dictionary value type") {
		t.Fatalf("Call with a nested struct field failed with %v, want the unsupported-type error", err)
	}
	t.Logf("nested struct is refused at encode time, not corrupted: %v", err)
}

// jsonNumbers marshals, through json.Marshaler, the three numbers that
// distinguish a float64 wire type from an exact one: the first integer float64
// cannot represent, an unsigned value above math.MaxInt64, and a number that is
// not an integer at all.
type jsonNumbers struct{}

const (
	bigInt   = int64(1)<<53 + 1  // 9007199254740993, the first integer float64 rounds
	hugeUint = uint64(1)<<63 + 1 // 9223372036854775809, above math.MaxInt64
)

func (jsonNumbers) MarshalJSON() ([]byte, error) {
	return []byte(`{"op":"mirror","n":9007199254740993,"u":9223372036854775809,"f":1.5,"neg":-9007199254740993}`), nil
}

// TestLiveJSONMarshalerCarriesInt64 tests the 2^53 boundary across a real
// connection rather than reasoning about float64.
//
// The json.Marshaler path used to decode into an any with json.Unmarshal, which
// makes every number a float64: 9007199254740993 arrived as 9007199254740992,
// off by one, with no error anywhere. The ordinary path had already been proven
// to carry int64 exactly (TestCodecDecodeWidensIntegers, and the control
// below), so the package contradicted itself about a property it cares about.
// json_marshaler.golden now records int64 and this test is what says the wire
// agrees.
func TestLiveJSONMarshalerCarriesInt64(t *testing.T) {
	session := dial(t, serviceName)

	// Positive control: the same integer sent as an int64 arrives exactly, so a
	// failure below belongs to the json.Marshaler path and not to XPC.
	msg, err := session.CallDictionary(context.Background(), xpc.Dictionary{"op": "mirror", "n": bigInt})
	if err != nil {
		t.Fatalf("control: %v", err)
	}
	if got := msg.Dictionary()["n"]; got != bigInt {
		t.Fatalf("control: n = %#v (%T), want int64(%d): XPC itself loses the value, so nothing below is about the codec", got, got, bigInt)
	}

	// Negative control for the instrument: a float64 with the rounded value is
	// what the defect produced. If the assertion below could not tell the two
	// apart it would pass on the broken tree too.
	if any(float64(bigInt)) == any(bigInt) {
		t.Fatalf("float64(%d) compares equal to int64(%d) through an any: the assertions below cannot see the defect", bigInt, bigInt)
	}

	msg, err = session.Call(context.Background(), jsonNumbers{})
	if err != nil {
		t.Fatalf("Call(jsonNumbers): %v", err)
	}
	got := msg.Dictionary()

	if got["n"] != bigInt {
		t.Errorf("n = %#v (%T), want int64(%d): the json.Marshaler path is still rounding through float64", got["n"], got["n"], bigInt)
	}
	if got["neg"] != -bigInt {
		t.Errorf("neg = %#v (%T), want int64(%d)", got["neg"], got["neg"], -bigInt)
	}
	// Above math.MaxInt64 the exact wire type is uint64. XPC reports its own
	// integer types, so a uint64 comes back as one.
	if got["u"] != hugeUint {
		t.Errorf("u = %#v (%T), want uint64(%d)", got["u"], got["u"], hugeUint)
	}
	if got["f"] != 1.5 {
		t.Errorf("f = %#v (%T), want float64(1.5): a non-integer must stay a double", got["f"], got["f"])
	}
	t.Logf("json.Marshaler path on the wire: n=%#v neg=%#v u=%#v f=%#v", got["n"], got["neg"], got["u"], got["f"])

	// The same exactness seen by a caller who decodes rather than inspects.
	var out struct {
		N   int64 `xpc:"n"`
		Neg int64 `xpc:"neg"`
	}
	if err := msg.Decode(&out); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if out.N != bigInt || out.Neg != -bigInt {
		t.Errorf("decoded N=%d Neg=%d, want %d and %d", out.N, out.Neg, bigInt, -bigInt)
	}

	// For the record: this is what json.Unmarshal into an any does to the same
	// bytes, which is what the path used to do.
	var lossy any
	if err := json.Unmarshal([]byte(`{"n":9007199254740993}`), &lossy); err != nil {
		t.Fatal(err)
	}
	f := lossy.(map[string]any)["n"].(float64)
	if int64(f) == bigInt {
		t.Fatalf("json.Unmarshal into an any preserved %d: the defect this test guards cannot recur, so the test is not a gate", bigInt)
	}
	t.Logf("json.Unmarshal into an any still yields float64 %.0f (off by %d); decodeJSONPayload uses a Decoder with UseNumber instead",
		f, bigInt-int64(f))
}
