package rdma

import "fmt"

// Capability describes the evidence level for one Apple Thunderbolt RDMA
// feature.
type Capability string

const (
	// CapabilityUnknown reports that this binding has no primary provider
	// observation for the feature.
	CapabilityUnknown Capability = "unknown"

	// CapabilityConfirmed reports an observed working provider path.
	CapabilityConfirmed Capability = "confirmed"

	// CapabilityRejected reports an observed provider rejection for the
	// documented device, configuration, and run.
	CapabilityRejected Capability = "rejected"

	// CapabilityObservedZero reports an observed zero value where a nonzero
	// capability value would be required by the documented path.
	CapabilityObservedZero Capability = "observed_zero"
)

// Capabilities describes the evidence currently available for Apple's
// Thunderbolt RDMA provider. It does not infer unsupported features from the
// absence of binding constants.
type Capabilities struct {
	QueuePairUC Capability
	SendRecv    Capability
	QueuePairRC Capability
	RDMARead    Capability
	RDMAWrite   Capability
	RemoteKey   Capability
}

// AppleThunderboltCapabilities reports the current evidence for Apple's
// Thunderbolt RDMA provider features.
//
// UC queue pairs and SEND/RECV are confirmed by the working collective path.
// On 2026-07-17, a guarded RC create on rdma_en3 returned a nil QP with
// EOPNOTSUPP (errno 102), so RC is rejected for that device, configuration,
// and run. A guarded memory registration on the same device requested remote
// read/write and returned rkey 0x0. That is an observed key value, not an
// attempted or rejected one-sided operation. READ and WRITE remain unknown.
func AppleThunderboltCapabilities() Capabilities {
	return Capabilities{
		QueuePairUC: CapabilityConfirmed,
		SendRecv:    CapabilityConfirmed,
		QueuePairRC: CapabilityRejected,
		RDMARead:    CapabilityUnknown,
		RDMAWrite:   CapabilityUnknown,
		RemoteKey:   CapabilityObservedZero,
	}
}

// ErrRDMAReadUnavailable reports that this binding does not expose RDMA READ.
// It does not make a claim about whether the provider accepts it.
var ErrRDMAReadUnavailable = fmt.Errorf("%w: RDMA READ is not exposed by this binding", ErrUnsupportedOperation)
