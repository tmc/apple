package experiment

import (
	"encoding/json"
	"testing"
)

func TestReceiptDigestStable(t *testing.T) {
	r := receipt(t)
	if err := r.Seal(); err != nil {
		t.Fatal(err)
	}
	want := r.Digest

	// Re-sealing the unchanged receipt is a no-op.
	if err := r.Seal(); err != nil {
		t.Fatal(err)
	}
	if r.Digest != want {
		t.Fatalf("re-seal changed digest: %s != %s", r.Digest, want)
	}

	// Any content change invalidates the digest.
	r.Verdict = VerdictFaster
	if err := r.Verify(); err == nil {
		t.Fatal("Verify accepted an edited receipt")
	}
}

func TestReceiptRoundTrip(t *testing.T) {
	r := receipt(t)
	b, err := r.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	got, err := ParseReceipt(b)
	if err != nil {
		t.Fatalf("ParseReceipt: %v", err)
	}
	if got.Digest != r.Digest || got.Verdict != r.Verdict || len(got.Samples) != len(r.Samples) {
		t.Fatalf("round trip mismatch: %+v", got)
	}

	// Tampering with any field breaks verification.
	var tampered map[string]any
	if err := json.Unmarshal(b, &tampered); err != nil {
		t.Fatal(err)
	}
	tampered["verdict"] = VerdictRefused
	b2, _ := json.Marshal(tampered)
	if _, err := ParseReceipt(b2); err == nil {
		t.Fatal("ParseReceipt accepted a tampered receipt")
	}
}

func TestParseReceiptWrongSchema(t *testing.T) {
	r := receipt(t)
	r.Schema = "something.else/v9"
	if err := r.Seal(); err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := ParseReceipt(b); err == nil {
		t.Fatal("ParseReceipt accepted a foreign schema")
	}
}

func TestHasRefusal(t *testing.T) {
	r := &Receipt{Refusals: []Refusal{{Reason: RefusalUnequalWork}}}
	if !r.Refused() || !r.HasRefusal(RefusalUnequalWork) {
		t.Fatal("refusal not recorded")
	}
	if r.HasRefusal(RefusalCorrectnessFailed) {
		t.Fatal("phantom refusal")
	}
}

func receipt(t *testing.T) *Receipt {
	t.Helper()
	return &Receipt{
		Schema: Schema,
		Workload: Workload{
			ID:        "matmul/64x64/fp32",
			WorkUnits: 4096,
			Bytes:     map[string]int64{"device": 32768},
		},
		Samples: []Sample{
			{Arm: "baseline", Kind: ArmBaseline, Round: 0, Order: 0, Elapsed: 100, WorkUnits: 4096, Verdict: VerdictPass},
			{Arm: "candidate", Kind: ArmCandidate, Round: 0, Order: 1, Elapsed: 90, WorkUnits: 4096, Verdict: VerdictPass},
			{Arm: "dup", Kind: ArmDuplicate, Round: 0, Order: 2, Elapsed: 101, WorkUnits: 4096, Verdict: VerdictPass},
		},
		Verdict: VerdictNoWinner,
	}
}
