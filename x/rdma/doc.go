// Package rdma provides operational helpers for Apple's generated RDMA
// bindings. It also provides deterministic teardown for a caller-owned
// resource chain.
//
// The package keeps policy that is not part of the C verbs ABI out of
// github.com/tmc/apple/rdma. It does not open devices, transition queue pairs,
// or post work requests; callers use it to classify errors, select route GIDs,
// and decide whether a read-only preflight permits one bounded RTR attempt.
package rdma
