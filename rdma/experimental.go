package rdma

// IBV_QPT_RCExperimental is the Apple verbs ABI value for a reliable-connected
// queue pair. It is exported solely for the one-shot RC capability probe.
//
// Its presence does not claim that Apple's Thunderbolt RDMA provider accepts
// RC queue pairs. Production code must not use it until a separately recorded
// provider observation establishes that capability.
const IBV_QPT_RCExperimental = 2
