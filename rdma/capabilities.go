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
}

// AppleThunderboltCapabilities reports the current evidence for Apple's
// Thunderbolt RDMA provider features.
//
// UC queue pairs and SEND/RECV are confirmed by the working collective path.
// RC, READ, and WRITE are unknown: this binding does not exercise or expose
// them, but that absence is not evidence of provider rejection.
func AppleThunderboltCapabilities() Capabilities {
	return Capabilities{
		QueuePairUC: CapabilityConfirmed,
		SendRecv:    CapabilityConfirmed,
		QueuePairRC: CapabilityUnknown,
		RDMARead:    CapabilityUnknown,
		RDMAWrite:   CapabilityUnknown,
	}
}

// ErrRDMAReadUnavailable reports that this binding does not expose RDMA READ.
// It does not make a claim about whether the provider accepts it.
var ErrRDMAReadUnavailable = fmt.Errorf("%w: RDMA READ is not exposed by this binding", ErrUnsupportedOperation)
