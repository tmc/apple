package rdma

import "fmt"

// Capabilities describes the verbs supported by Apple's Thunderbolt RDMA
// provider. It is not a claim about every provider implementing libibverbs.
type Capabilities struct {
	QueuePairUC bool
	SendRecv    bool
	RDMAWrite   bool
	RDMARead    bool
}

// AppleThunderboltCapabilities reports the verbs supported by Apple's
// Thunderbolt RDMA provider.
//
// The provider accepts UC queue pairs and supports SEND/RECV and RDMA WRITE.
// It rejects RC queue pairs, so RDMA READ is unavailable.
func AppleThunderboltCapabilities() Capabilities {
	return Capabilities{
		QueuePairUC: true,
		SendRecv:    true,
		RDMAWrite:   true,
	}
}

// ErrRDMAReadUnsupported reports that Apple's Thunderbolt RDMA provider does
// not support RDMA READ. It wraps ErrUnsupportedOperation.
var ErrRDMAReadUnsupported = fmt.Errorf("%w: RDMA READ requires an RC queue pair, which Apple's Thunderbolt RDMA provider rejects", ErrUnsupportedOperation)
