// Code generated from Apple documentation for vmnet. DO NOT EDIT.

package vmnet
import (
	"github.com/tmc/apple/objectivec"
)

// C struct types
// Vmpktdesc - Describes a packet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vmnet/vmpktdesc
type Vmpktdesc struct {
	Vm_pkt_size int // The size of the packet, in bytes.
	Vm_pkt_iovcnt uint32 // The number of packet buffers in `vm_pkt_iov`.
	Vm_flags uint32 // Option flags. Should be set to `0` on read.
	Vm_pkt_iov objectivec.IObject // An array of packet buffers.

}








