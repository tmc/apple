// Command call3smoke is a fixture for TestRDMACall3ProductionLink. It exists
// only to be built into a binary whose symbol table can then be inspected:
// linking rdma into a main package is the only way to observe that purego's
// syscallXABI0 trampoline lands in initialized read-only data, which a library
// test cannot establish about itself.
//
// It lives under testdata so it is not part of ./... and ships no command.
package main

import (
	"fmt"

	_ "github.com/tmc/apple/rdma"
)

func main() {
	// TestRDMACall3ProductionLink asserts on this exact line, so that a binary
	// which links but traps at startup is distinguishable from one that runs.
	fmt.Println("rdma Call3 trampoline: available")
}
