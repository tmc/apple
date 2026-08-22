package main

import (
	"fmt"

	"github.com/tmc/apple/hypervisor"
)

func main() {
	var maxVCPU uint32
	if ok := call("hv_vm_get_max_vcpu_count", func() int32 {
		return int32(hypervisor.HVVmGetMaxVCPUCount(&maxVCPU))
	}); !ok {
		return
	}
	fmt.Printf("max vCPUs: %d\n", maxVCPU)

	var maxIPA uint32
	if call("hv_vm_config_get_max_ipa_size", func() int32 {
		return int32(hypervisor.HVVmConfigGetMaxIPASize(&maxIPA))
	}) {
		fmt.Printf("max IPA size: %d bits\n", maxIPA)
	}
}

func call(name string, fn func() int32) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s unavailable: %v\n", name, r)
			ok = false
		}
	}()
	ret := fn()
	if ret != int32(hypervisor.HVSuccess) {
		fmt.Printf("%s failed: %s (%d)\n", name, hypervisor.HVReturn(ret), ret)
		return false
	}
	return true
}
