package coremlcompiler

import (
	"encoding/hex"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// unknownFieldsModel is testdata/gen_unknownfields.py's output: a model
// carrying fields this package does not model. Kept as hex so the
// repository holds no binary fixture.
const unknownFieldsModel = "080712530a210a01781208616e20696e7075741a122a100a010410a08004fa01" +
	"060a0408011010520e0a01791a092a070a010410a080045a017992030e0a0178" +
	"1a092a070a010410a08004a206090a07666978747572655001b21f020801"

// TestUnknownFieldsSurviveReencode checks that fields this package does not
// model survive decodeModel followed by EncodeModel. The fixture is written by
// coremltools (testdata/gen_unknownfields.py), not by our encoder, so the test
// can observe fields our encoder cannot produce.
func TestUnknownFieldsSurviveReencode(t *testing.T) {
	data, err := hex.DecodeString(unknownFieldsModel)
	if err != nil {
		t.Fatal(err)
	}
	m, err := decodeModel(data)
	if err != nil {
		t.Fatal(err)
	}
	got := collectFields(t, data)
	want := collectFields(t, EncodeModel(m))

	// Spot-check the fields the Go types drop, so a change that stops
	// preserving them names the field rather than a diff.
	for _, path := range []string{
		"10",      // Model.isUpdatable
		"2/11",    // ModelDescription.predictedFeatureName
		"2/50",    // ModelDescription.trainingInput
		"2/1/2",   // FeatureDescription.shortDescription
		"2/1/3/5", // ArrayFeatureType, holding the unmodeled shapeRange
	} {
		if !hasPath(got, path) {
			t.Fatalf("fixture lacks field %s; regenerate unknownFieldsModel", path)
		}
		if !hasPath(want, path) {
			t.Errorf("field %s lost in re-encode", path)
		}
	}
	for path, vals := range got {
		if strings.Join(want[path], ",") != strings.Join(vals, ",") {
			t.Errorf("field %s: got %v, want %v", path, want[path], vals)
		}
	}
	for path := range want {
		if _, ok := got[path]; !ok {
			t.Errorf("field %s: invented by re-encode", path)
		}
	}
}

// hasPath reports whether fields records path or anything nested under it.
func hasPath(fields map[string][]string, path string) bool {
	for p := range fields {
		if p == path || strings.HasPrefix(p, path+"/") {
			return true
		}
	}
	return false
}

// collectFields walks protobuf wire bytes and returns every leaf value keyed by
// its slash-separated field-number path. Length-delimited values are descended
// into when they parse as a message; the same heuristic applies to both sides
// of a comparison.
func collectFields(t *testing.T, data []byte) map[string][]string {
	out := make(map[string][]string)
	var walk func(prefix string, data []byte)
	walk = func(prefix string, data []byte) {
		r := newProtoReader(data)
		for !r.done() {
			field, wire, err := r.readTag()
			if err != nil {
				t.Fatalf("walk %s: %v", prefix, err)
			}
			path := prefix + strconv.Itoa(field)
			switch wire {
			case wireBytes:
				b, err := r.readBytes()
				if err != nil {
					t.Fatalf("walk %s: %v", path, err)
				}
				if isMessage(b) {
					walk(path+"/", b)
					continue
				}
				out[path] = append(out[path], hex.EncodeToString(b))
			case wireVarint:
				v, err := r.readVarint()
				if err != nil {
					t.Fatalf("walk %s: %v", path, err)
				}
				out[path] = append(out[path], strconv.FormatUint(v, 10))
			default:
				if err := r.skip(wire); err != nil {
					t.Fatalf("walk %s: %v", path, err)
				}
				out[path] = append(out[path], "opaque")
			}
		}
	}
	walk("", data)
	for _, vals := range out {
		sort.Strings(vals)
	}
	return out
}

// isMessage reports whether b parses cleanly as protobuf wire format.
// Empty input is not treated as a message: it carries no fields to compare.
func isMessage(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	r := newProtoReader(b)
	for !r.done() {
		field, wire, err := r.readTag()
		if err != nil || field == 0 {
			return false
		}
		if err := r.skip(wire); err != nil {
			return false
		}
	}
	return true
}
