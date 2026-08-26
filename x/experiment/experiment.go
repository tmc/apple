package experiment

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"
)

// Schema is the receipt schema identifier written into every receipt.
const Schema = "mlxperf.receipt/v1"

// ArmKind classifies an arm's role in an experiment.
type ArmKind string

const (
	// ArmBaseline is the immutable reference arm.
	ArmBaseline ArmKind = "baseline"
	// ArmCandidate is the arm under test.
	ArmCandidate ArmKind = "candidate"
	// ArmDuplicate is an identical duplicate control used to estimate
	// harness noise. An experiment without one may not conclude performance.
	ArmDuplicate ArmKind = "duplicate"
)

// Verdict records the correctness outcome of one sample.
type Verdict string

const (
	VerdictPass Verdict = "pass"
	VerdictFail Verdict = "fail"
)

// Receipt verdicts.
const (
	// VerdictFaster means the candidate beat the baseline beyond the noise
	// measured by the duplicate control.
	VerdictFaster = "faster"
	// VerdictSlower means the candidate lost to the baseline.
	VerdictSlower = "slower"
	// VerdictNoWinner means no performance conclusion is possible: the
	// claimed delta does not exceed the control gap.
	VerdictNoWinner = "no-winner"
	// VerdictRefused means a gate refused the experiment; there is no score.
	VerdictRefused = "refused"
)

// Workload identifies what was measured. Any change to identity, work units,
// or byte accounting creates a new workload; receipts for different
// workloads are never comparable.
type Workload struct {
	ID        string           `json:"id"`
	WorkUnits int64            `json:"work_units"`
	Bytes     map[string]int64 `json:"bytes,omitempty"`
}

// Sample is one measured arm execution with its order and environment
// observations.
type Sample struct {
	Arm       string           `json:"arm"`
	Kind      ArmKind          `json:"kind"`
	Round     int              `json:"round"`
	Order     int              `json:"order"`
	Elapsed   time.Duration    `json:"elapsed"`
	WorkUnits int64            `json:"work_units"`
	Bytes     map[string]int64 `json:"bytes,omitempty"`
	Output    string           `json:"output,omitempty"`
	Verdict   Verdict          `json:"verdict"`
	Note      string           `json:"note,omitempty"`
}

// Refusal is a typed reason the arena declined to rank or promote. A refusal
// replaces a score; it never annotates one.
type Refusal struct {
	Reason string `json:"reason"`
	Detail string `json:"detail,omitempty"`
}

func (r Refusal) Error() string {
	if r.Detail == "" {
		return r.Reason
	}
	return fmt.Sprintf("%s: %s", r.Reason, r.Detail)
}

// Refusal reasons. These are contract, not prose: analysis code switches on
// them.
const (
	RefusalUnequalWork         = "unequal-work"
	RefusalCorrectnessFailed   = "correctness-failed"
	RefusalMissingControl      = "missing-duplicate-control"
	RefusalImpossibleByteRate  = "impossible-byte-rate"
	RefusalContaminatedHost    = "contaminated-host"
	RefusalIdenticalArms       = "identical-arms"
	RefusalInsufficientSamples = "insufficient-samples"
	RefusalIdentityDrift       = "identity-drift"
)

// Receipt is the canonical result of one measured experiment.
type Receipt struct {
	Schema    string    `json:"schema"`
	CreatedAt time.Time `json:"created_at"`
	Workload  Workload  `json:"workload"`
	Label     string    `json:"label,omitempty"`
	Samples   []Sample  `json:"samples"`
	Refusals  []Refusal `json:"refusals,omitempty"`
	Verdict   string    `json:"verdict"`
	Winner    string    `json:"winner,omitempty"`
	Digest    string    `json:"digest"`
}

// Refused reports whether any refusal was recorded.
func (r *Receipt) Refused() bool { return len(r.Refusals) > 0 }

// HasRefusal reports whether a refusal with the given reason was recorded.
func (r *Receipt) HasRefusal(reason string) bool {
	for _, ref := range r.Refusals {
		if ref.Reason == reason {
			return true
		}
	}
	return false
}

// Seal computes the digest over the canonical form of the receipt (with the
// digest field cleared) and stores it.
func (r *Receipt) Seal() error {
	sum, err := r.digest()
	if err != nil {
		return err
	}
	r.Digest = sum
	return nil
}

func (r *Receipt) digest() (string, error) {
	saved := r.Digest
	r.Digest = ""
	defer func() { r.Digest = saved }()
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:]), nil
}

// Verify checks that a sealed receipt still matches its digest.
func (r *Receipt) Verify() error {
	if r.Schema != Schema {
		return fmt.Errorf("receipt schema %q, want %q", r.Schema, Schema)
	}
	if r.Digest == "" {
		return errors.New("receipt is not sealed")
	}
	want, err := r.digest()
	if err != nil {
		return err
	}
	if want != r.Digest {
		return fmt.Errorf("receipt digest mismatch: have %s, computed %s", r.Digest, want)
	}
	return nil
}

// Marshal canonicalizes and seals the receipt, then encodes it as JSON.
func (r *Receipt) Marshal() ([]byte, error) {
	sort.Slice(r.Samples, func(i, j int) bool { return r.Samples[i].Order < r.Samples[j].Order })
	if err := r.Seal(); err != nil {
		return nil, err
	}
	return json.MarshalIndent(r, "", "  ")
}

// ParseReceipt decodes a JSON-encoded receipt and verifies its schema and
// digest. A receipt that fails verification is returned alongside the error
// so callers can retain the tampered evidence.
func ParseReceipt(b []byte) (*Receipt, error) {
	var r Receipt
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}
	if err := r.Verify(); err != nil {
		return &r, err
	}
	return &r, nil
}
