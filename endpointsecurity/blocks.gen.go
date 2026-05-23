// Code generated from Apple documentation. DO NOT EDIT.

package endpointsecurity

import (
	"github.com/tmc/apple/objc"
)

// EsHandlerBlock handles A block that handles a message received from Endpoint Security.

// NewEsHandlerBlock wraps a Go [EsHandlerBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewEsHandlerBlock(handler EsHandlerBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive *Es_client_t, extra0 *Es_message_t) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}
