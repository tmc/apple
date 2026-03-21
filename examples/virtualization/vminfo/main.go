// Command vminfo queries Apple Virtualization framework capabilities.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/tmc/apple/virtualization"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	vmClass := virtualization.GetVZVirtualMachineClass()
	cfgClass := virtualization.GetVZVirtualMachineConfigurationClass()

	supported := vmClass.Supported()
	minCPU := cfgClass.MinimumAllowedCPUCount()
	maxCPU := cfgClass.MaximumAllowedCPUCount()
	minMem := cfgClass.MinimumAllowedMemorySize()
	maxMem := cfgClass.MaximumAllowedMemorySize()

	info := map[string]any{
		"supported":       supported,
		"min_cpu_count":   minCPU,
		"max_cpu_count":   maxCPU,
		"min_memory_bytes": minMem,
		"max_memory_bytes": maxMem,
		"min_memory_mb":   minMem / (1024 * 1024),
		"max_memory_mb":   maxMem / (1024 * 1024),
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(info); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
