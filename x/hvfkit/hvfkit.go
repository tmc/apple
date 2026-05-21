//go:build darwin && arm64

package hvfkit

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/hypervisor"
)

// Capabilities describes host Hypervisor.framework limits.
type Capabilities struct {
	MaxVCPUCount uint32
	MaxIPABits   uint32
	EL2Supported bool
}

// Config is an opaque Hypervisor.framework VM configuration.
type Config struct {
	Handle       unsafe.Pointer
	Capabilities Capabilities
	IPABits      uint32
	IPAGranule   hypervisor.HVIPAGranule
	EL2Enabled   bool
	Created      bool
}

// QueryCapabilities returns host Hypervisor.framework capabilities.
func QueryCapabilities() (Capabilities, error) {
	var caps Capabilities
	if err := call("hv_vm_get_max_vcpu_count", func() int32 {
		return hypervisor.HVVmGetMaxVCPUCount(&caps.MaxVCPUCount)
	}); err != nil {
		return Capabilities{}, err
	}
	if err := call("hv_vm_config_get_max_ipa_size", func() int32 {
		return hypervisor.HVVmConfigGetMaxIPASize(&caps.MaxIPABits)
	}); err != nil {
		return Capabilities{}, err
	}
	if err := call("hv_vm_config_get_el2_supported", func() int32 {
		return hypervisor.HVVmConfigGetEl2Supported(&caps.EL2Supported)
	}); err != nil {
		return Capabilities{}, err
	}
	return caps, nil
}

// NewConfig creates and inspects a Hypervisor.framework VM configuration.
func NewConfig(ipaBits uint32, enableEL2 bool, createVM bool) (*Config, error) {
	handle, err := create("hv_vm_config_create", hypervisor.HVVmConfigCreate)
	if err != nil {
		return nil, err
	}
	if handle == nil {
		return nil, fmt.Errorf("hv_vm_config_create returned nil")
	}
	if ipaBits > 0 {
		if err := call("hv_vm_config_set_ipa_size", func() int32 {
			return hypervisor.HVVmConfigSetIPASize(handle, ipaBits)
		}); err != nil {
			return nil, err
		}
	}
	if enableEL2 {
		if err := call("hv_vm_config_set_el2_enabled", func() int32 {
			return hypervisor.HVVmConfigSetEl2Enabled(handle, true)
		}); err != nil {
			return nil, err
		}
	}
	cfg := &Config{Handle: handle}
	if err := call("hv_vm_config_get_ipa_size", func() int32 {
		return hypervisor.HVVmConfigGetIPASize(handle, &cfg.IPABits)
	}); err != nil {
		return nil, err
	}
	if err := call("hv_vm_config_get_ipa_granule", func() int32 {
		return hypervisor.HVVmConfigGetIPAGranule(handle, &cfg.IPAGranule)
	}); err != nil {
		return nil, err
	}
	if err := call("hv_vm_config_get_el2_enabled", func() int32 {
		return hypervisor.HVVmConfigGetEl2Enabled(handle, &cfg.EL2Enabled)
	}); err != nil {
		return nil, err
	}
	if createVM {
		if err := CreateVM(handle); err != nil {
			return nil, err
		}
		cfg.Created = true
	}
	return cfg, nil
}

// DefaultIPASize returns the default IPA size for new VM configurations.
func DefaultIPASize() (uint32, error) {
	var bits uint32
	err := call("hv_vm_config_get_default_ipa_size", func() int32 {
		return hypervisor.HVVmConfigGetDefaultIPASize(&bits)
	})
	return bits, err
}

// DefaultIPAGranule returns the default IPA granule for new VM configurations.
func DefaultIPAGranule() (hypervisor.HVIPAGranule, error) {
	var granule hypervisor.HVIPAGranule
	err := call("hv_vm_config_get_default_ipa_granule", func() int32 {
		return hypervisor.HVVmConfigGetDefaultIPAGranule(&granule)
	})
	return granule, err
}

// GetConfigIPAGranule reads a VM configuration's IPA granule.
func GetConfigIPAGranule(config unsafe.Pointer) (hypervisor.HVIPAGranule, error) {
	var granule hypervisor.HVIPAGranule
	err := call("hv_vm_config_get_ipa_granule", func() int32 {
		return hypervisor.HVVmConfigGetIPAGranule(config, &granule)
	})
	return granule, err
}

// SetConfigIPAGranule writes a VM configuration's IPA granule.
func SetConfigIPAGranule(config unsafe.Pointer, granule hypervisor.HVIPAGranule) error {
	return call("hv_vm_config_set_ipa_granule", func() int32 {
		return hypervisor.HVVmConfigSetIPAGranule(config, granule)
	})
}

// CreateVM creates a process VM with config.
func CreateVM(config unsafe.Pointer) error {
	return call("hv_vm_create", func() int32 {
		return hypervisor.HVVmCreate(config)
	})
}

// DestroyVM destroys the process VM.
func DestroyVM() error {
	return call("hv_vm_destroy", hypervisor.HVVmDestroy)
}

// MapMemory maps host memory into guest physical memory.
func MapMemory(addr unsafe.Pointer, ipa uint64, size uintptr, flags uint64) error {
	return call("hv_vm_map", func() int32 {
		return hypervisor.HVVmMap(addr, ipa, size, flags)
	})
}

// UnmapMemory removes a guest physical memory mapping.
func UnmapMemory(ipa uint64, size uintptr) error {
	return call("hv_vm_unmap", func() int32 {
		return hypervisor.HVVmUnmap(ipa, size)
	})
}

// AllocateMemory allocates Hypervisor.framework-owned memory.
func AllocateMemory(addr unsafe.Pointer, size uintptr, flags uint64) error {
	return call("hv_vm_allocate", func() int32 {
		return hypervisor.HVVmAllocate(addr, size, flags)
	})
}

// DeallocateMemory releases Hypervisor.framework-owned memory.
func DeallocateMemory(addr unsafe.Pointer, size uintptr) error {
	return call("hv_vm_deallocate", func() int32 {
		return hypervisor.HVVmDeallocate(addr, size)
	})
}

// ProtectMemory changes guest physical memory permissions.
func ProtectMemory(ipa uint64, size uintptr, flags uint64) error {
	return call("hv_vm_protect", func() int32 {
		return hypervisor.HVVmProtect(ipa, size, flags)
	})
}

// NewVCPUConfig creates a Hypervisor.framework vCPU configuration.
func NewVCPUConfig() (unsafe.Pointer, error) {
	config, err := create("hv_vcpu_config_create", hypervisor.HVVCPUConfigCreate)
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, fmt.Errorf("hv_vcpu_config_create returned nil")
	}
	return config, nil
}

// GetVCPUConfigFeatureReg reads a feature register from a vCPU configuration.
func GetVCPUConfigFeatureReg(config unsafe.Pointer, reg hypervisor.HVFeatureReg) (uint64, error) {
	var value uint64
	err := call("hv_vcpu_config_get_feature_reg", func() int32 {
		return hypervisor.HVVCPUConfigGetFeatureReg(config, reg, &value)
	})
	return value, err
}

// CreateVCPU creates a vCPU and returns the HVF-owned exit area.
func CreateVCPU(exit **hypervisor.HVVCPUExit, config unsafe.Pointer) (uint64, error) {
	var vcpu uint64
	err := call("hv_vcpu_create", func() int32 {
		return hypervisor.HVVCPUCreate(&vcpu, exit, config)
	})
	return vcpu, err
}

// DestroyVCPU destroys a vCPU.
func DestroyVCPU(id uint64) error {
	return call("hv_vcpu_destroy", func() int32 {
		return hypervisor.HVVCPUDestroy(id)
	})
}

// RunVCPU enters a vCPU until the next exit.
func RunVCPU(id uint64) error {
	return call("hv_vcpu_run", func() int32 {
		return hypervisor.HVVCPURun(id)
	})
}

// ExitVCPUs requests exits from the supplied vCPUs.
func ExitVCPUs(ids []uint64) error {
	return call("hv_vcpus_exit", func() int32 {
		if len(ids) == 0 {
			return hypervisor.HVVcpusExit(nil, 0)
		}
		return hypervisor.HVVcpusExit(&ids[0], uint32(len(ids)))
	})
}

// GetExecTime returns the time a vCPU has spent executing.
func GetExecTime(vcpu uint64) (uint64, error) {
	var tm uint64
	err := call("hv_vcpu_get_exec_time", func() int32 {
		return hypervisor.HVVCPUGetExecTime(vcpu, &tm)
	})
	return tm, err
}

// GetPendingInterrupt reports whether an interrupt type is pending.
func GetPendingInterrupt(vcpu uint64, typ hypervisor.HVInterruptType) (bool, error) {
	var pending bool
	err := call("hv_vcpu_get_pending_interrupt", func() int32 {
		return hypervisor.HVVCPUGetPendingInterrupt(vcpu, typ, &pending)
	})
	return pending, err
}

// SetPendingInterrupt changes whether an interrupt type is pending.
func SetPendingInterrupt(vcpu uint64, typ hypervisor.HVInterruptType, pending bool) error {
	return call("hv_vcpu_set_pending_interrupt", func() int32 {
		return hypervisor.HVVCPUSetPendingInterrupt(vcpu, typ, pending)
	})
}

// GetReg reads a general-purpose vCPU register.
func GetReg(vcpu uint64, reg hypervisor.HVReg) (uint64, error) {
	var value uint64
	err := call("hv_vcpu_get_reg", func() int32 {
		return hypervisor.HVVCPUGetReg(vcpu, reg, &value)
	})
	return value, err
}

// SetReg writes a general-purpose vCPU register.
func SetReg(vcpu uint64, reg hypervisor.HVReg, value uint64) error {
	return call("hv_vcpu_set_reg", func() int32 {
		return hypervisor.HVVCPUSetReg(vcpu, reg, value)
	})
}

// GetSysReg reads a vCPU system register.
func GetSysReg(vcpu uint64, reg hypervisor.HVSysReg) (uint64, error) {
	var value uint64
	err := call("hv_vcpu_get_sys_reg", func() int32 {
		return hypervisor.HVVCPUGetSysReg(vcpu, reg, &value)
	})
	return value, err
}

// SetSysReg writes a vCPU system register.
func SetSysReg(vcpu uint64, reg hypervisor.HVSysReg, value uint64) error {
	return call("hv_vcpu_set_sys_reg", func() int32 {
		return hypervisor.HVVCPUSetSysReg(vcpu, reg, value)
	})
}

// GetSIMDFPReg reads a SIMD/FP register.
func GetSIMDFPReg(vcpu uint64, reg hypervisor.HVSIMDFPReg) ([16]byte, error) {
	var value [16]byte
	err := call("hv_vcpu_get_simd_fp_reg", func() int32 {
		return hypervisor.HVVCPUGetSIMDFPReg(vcpu, reg, &value)
	})
	return value, err
}

// SetSIMDFPReg writes a SIMD/FP register.
func SetSIMDFPReg(vcpu uint64, reg hypervisor.HVSIMDFPReg, value [16]byte) error {
	return call("hv_vcpu_set_simd_fp_reg", func() int32 {
		return hypervisor.HVVCPUSetSIMDFPReg(vcpu, reg, value)
	})
}

// GetVtimerMask reports whether the virtual timer is masked.
func GetVtimerMask(vcpu uint64) (bool, error) {
	var masked bool
	err := call("hv_vcpu_get_vtimer_mask", func() int32 {
		return hypervisor.HVVCPUGetVtimerMask(vcpu, &masked)
	})
	return masked, err
}

// GetVtimerOffset reads the virtual timer offset.
func GetVtimerOffset(vcpu uint64) (uint64, error) {
	var offset uint64
	err := call("hv_vcpu_get_vtimer_offset", func() int32 {
		return hypervisor.HVVCPUGetVtimerOffset(vcpu, &offset)
	})
	return offset, err
}

// SetVtimerMask masks or unmasks the virtual timer.
func SetVtimerMask(vcpu uint64, masked bool) error {
	return call("hv_vcpu_set_vtimer_mask", func() int32 {
		return hypervisor.HVVCPUSetVtimerMask(vcpu, masked)
	})
}

// SetVtimerOffset writes the virtual timer offset.
func SetVtimerOffset(vcpu uint64, offset uint64) error {
	return call("hv_vcpu_set_vtimer_offset", func() int32 {
		return hypervisor.HVVCPUSetVtimerOffset(vcpu, offset)
	})
}

// GetTrapDebugExceptions reports whether debug exceptions trap to the host.
func GetTrapDebugExceptions(vcpu uint64) (bool, error) {
	var enabled bool
	err := call("hv_vcpu_get_trap_debug_exceptions", func() int32 {
		return hypervisor.HVVCPUGetTrapDebugExceptions(vcpu, &enabled)
	})
	return enabled, err
}

// SetTrapDebugExceptions configures whether debug exceptions trap to the host.
func SetTrapDebugExceptions(vcpu uint64, enabled bool) error {
	return call("hv_vcpu_set_trap_debug_exceptions", func() int32 {
		return hypervisor.HVVCPUSetTrapDebugExceptions(vcpu, enabled)
	})
}

// GetTrapDebugRegAccesses reports whether debug register accesses trap to the host.
func GetTrapDebugRegAccesses(vcpu uint64) (bool, error) {
	var enabled bool
	err := call("hv_vcpu_get_trap_debug_reg_accesses", func() int32 {
		return hypervisor.HVVCPUGetTrapDebugRegAccesses(vcpu, &enabled)
	})
	return enabled, err
}

// SetTrapDebugRegAccesses configures whether debug register accesses trap to the host.
func SetTrapDebugRegAccesses(vcpu uint64, enabled bool) error {
	return call("hv_vcpu_set_trap_debug_reg_accesses", func() int32 {
		return hypervisor.HVVCPUSetTrapDebugRegAccesses(vcpu, enabled)
	})
}

// CreateGIC configures and creates a GICv3 device.
func CreateGIC(distributorBase, redistributorBase uint64) error {
	config, err := NewGICConfig(distributorBase, redistributorBase)
	if err != nil {
		return err
	}
	return CreateGICWithConfig(config)
}

// NewGICConfig creates a GICv3 configuration with distributor bases set.
func NewGICConfig(distributorBase, redistributorBase uint64) (unsafe.Pointer, error) {
	config, err := create("hv_gic_config_create", hypervisor.HVGICConfigCreate)
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, fmt.Errorf("hv_gic_config_create returned nil")
	}
	if err := call("hv_gic_config_set_distributor_base", func() int32 {
		return hypervisor.HVGICConfigSetDistributorBase(config, distributorBase)
	}); err != nil {
		return nil, err
	}
	if err := call("hv_gic_config_set_redistributor_base", func() int32 {
		return hypervisor.HVGICConfigSetRedistributorBase(config, redistributorBase)
	}); err != nil {
		return nil, err
	}
	return config, nil
}

// SetGICConfigMSIInterruptRange configures the GIC MSI interrupt range.
func SetGICConfigMSIInterruptRange(config unsafe.Pointer, base, count uint32) error {
	return call("hv_gic_config_set_msi_interrupt_range", func() int32 {
		return hypervisor.HVGICConfigSetMsiInterruptRange(config, base, count)
	})
}

// SetGICConfigMSIRegionBase configures the GIC MSI region base.
func SetGICConfigMSIRegionBase(config unsafe.Pointer, base uint64) error {
	return call("hv_gic_config_set_msi_region_base", func() int32 {
		return hypervisor.HVGICConfigSetMsiRegionBase(config, base)
	})
}

// CreateGICWithConfig creates a GICv3 device from config.
func CreateGICWithConfig(config unsafe.Pointer) error {
	return call("hv_gic_create", func() int32 {
		return hypervisor.HVGICCreate(config)
	})
}

// GetGICDistributorBaseAlignment returns the required distributor base alignment.
func GetGICDistributorBaseAlignment() (uintptr, error) {
	var alignment uintptr
	err := call("hv_gic_get_distributor_base_alignment", func() int32 {
		return hypervisor.HVGICGetDistributorBaseAlignment(&alignment)
	})
	return alignment, err
}

// GetGICDistributorSize returns the GIC distributor region size.
func GetGICDistributorSize() (uintptr, error) {
	var size uintptr
	err := call("hv_gic_get_distributor_size", func() int32 {
		return hypervisor.HVGICGetDistributorSize(&size)
	})
	return size, err
}

// GetGICRedistributorBaseAlignment returns the required redistributor base alignment.
func GetGICRedistributorBaseAlignment() (uintptr, error) {
	var alignment uintptr
	err := call("hv_gic_get_redistributor_base_alignment", func() int32 {
		return hypervisor.HVGICGetRedistributorBaseAlignment(&alignment)
	})
	return alignment, err
}

// GetGICRedistributorSize returns one GIC redistributor region size.
func GetGICRedistributorSize() (uintptr, error) {
	var size uintptr
	err := call("hv_gic_get_redistributor_size", func() int32 {
		return hypervisor.HVGICGetRedistributorSize(&size)
	})
	return size, err
}

// GetGICRedistributorRegionSize returns the full redistributor region size.
func GetGICRedistributorRegionSize() (uintptr, error) {
	var size uintptr
	err := call("hv_gic_get_redistributor_region_size", func() int32 {
		return hypervisor.HVGICGetRedistributorRegionSize(&size)
	})
	return size, err
}

// GetGICRedistributorBase returns a vCPU redistributor base address.
func GetGICRedistributorBase(vcpu uint64) (uint64, error) {
	var base uint64
	err := call("hv_gic_get_redistributor_base", func() int32 {
		return hypervisor.HVGICGetRedistributorBase(vcpu, &base)
	})
	return base, err
}

// ResetGIC resets the current GIC state.
func ResetGIC() error {
	return call("hv_gic_reset", hypervisor.HVGICReset)
}

// SetSPI sets a shared peripheral interrupt level.
func SetSPI(intid uint32, level bool) error {
	return call("hv_gic_set_spi", func() int32 {
		return hypervisor.HVGICSetSpi(intid, level)
	})
}

// GetGICIntID returns the concrete GIC interrupt ID for a reserved interrupt.
func GetGICIntID(interrupt hypervisor.HVGICIntid) (uint32, error) {
	var intid uint32
	err := call("hv_gic_get_intid", func() int32 {
		return hypervisor.HVGICGetIntid(interrupt, &intid)
	})
	return intid, err
}

// GetGICDistributorReg reads a GIC distributor register.
func GetGICDistributorReg(reg hypervisor.HVGICDistributorReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_distributor_reg", func() int32 {
		return hypervisor.HVGICGetDistributorReg(reg, &value)
	})
	return value, err
}

// SetGICDistributorReg writes a GIC distributor register.
func SetGICDistributorReg(reg hypervisor.HVGICDistributorReg, value uint64) error {
	return call("hv_gic_set_distributor_reg", func() int32 {
		return hypervisor.HVGICSetDistributorReg(reg, value)
	})
}

// GetGICRedistributorReg reads a GIC redistributor register.
func GetGICRedistributorReg(vcpu uint64, reg hypervisor.HVGICRedistributorReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_redistributor_reg", func() int32 {
		return hypervisor.HVGICGetRedistributorReg(vcpu, reg, &value)
	})
	return value, err
}

// SetGICRedistributorReg writes a GIC redistributor register.
func SetGICRedistributorReg(vcpu uint64, reg hypervisor.HVGICRedistributorReg, value uint64) error {
	return call("hv_gic_set_redistributor_reg", func() int32 {
		return hypervisor.HVGICSetRedistributorReg(vcpu, reg, value)
	})
}

// GetGICICCReg reads a GIC CPU-interface register.
func GetGICICCReg(vcpu uint64, reg hypervisor.HVGICIccReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_icc_reg", func() int32 {
		return hypervisor.HVGICGetIccReg(vcpu, reg, &value)
	})
	return value, err
}

// SetGICICCReg writes a GIC CPU-interface register.
func SetGICICCReg(vcpu uint64, reg hypervisor.HVGICIccReg, value uint64) error {
	return call("hv_gic_set_icc_reg", func() int32 {
		return hypervisor.HVGICSetIccReg(vcpu, reg, value)
	})
}

// GetGICICHReg reads a GIC hypervisor-control register.
func GetGICICHReg(vcpu uint64, reg hypervisor.HVGICIchReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_ich_reg", func() int32 {
		return hypervisor.HVGICGetIchReg(vcpu, reg, &value)
	})
	return value, err
}

// SetGICICHReg writes a GIC hypervisor-control register.
func SetGICICHReg(vcpu uint64, reg hypervisor.HVGICIchReg, value uint64) error {
	return call("hv_gic_set_ich_reg", func() int32 {
		return hypervisor.HVGICSetIchReg(vcpu, reg, value)
	})
}

// GetGICICVReg reads a GIC virtual CPU-interface register.
func GetGICICVReg(vcpu uint64, reg hypervisor.HVGICIcvReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_icv_reg", func() int32 {
		return hypervisor.HVGICGetIcvReg(vcpu, reg, &value)
	})
	return value, err
}

// SetGICICVReg writes a GIC virtual CPU-interface register.
func SetGICICVReg(vcpu uint64, reg hypervisor.HVGICIcvReg, value uint64) error {
	return call("hv_gic_set_icv_reg", func() int32 {
		return hypervisor.HVGICSetIcvReg(vcpu, reg, value)
	})
}

// GetGICMSIReg reads a GIC MSI register.
func GetGICMSIReg(reg hypervisor.HVGICMsiReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_msi_reg", func() int32 {
		return hypervisor.HVGICGetMsiReg(reg, &value)
	})
	return value, err
}

// SetGICMSIReg writes a GIC MSI register.
func SetGICMSIReg(reg hypervisor.HVGICMsiReg, value uint64) error {
	return call("hv_gic_set_msi_reg", func() int32 {
		return hypervisor.HVGICSetMsiReg(reg, value)
	})
}

// GetGICMSIRegionBaseAlignment returns the required MSI region base alignment.
func GetGICMSIRegionBaseAlignment() (uintptr, error) {
	var alignment uintptr
	err := call("hv_gic_get_msi_region_base_alignment", func() int32 {
		return hypervisor.HVGICGetMsiRegionBaseAlignment(&alignment)
	})
	return alignment, err
}

// GetGICMSIRegionSize returns the GIC MSI region size.
func GetGICMSIRegionSize() (uintptr, error) {
	var size uintptr
	err := call("hv_gic_get_msi_region_size", func() int32 {
		return hypervisor.HVGICGetMsiRegionSize(&size)
	})
	return size, err
}

// GetGICSPIInterruptRange returns the available shared peripheral interrupt range.
func GetGICSPIInterruptRange() (base, count uint32, err error) {
	err = call("hv_gic_get_spi_interrupt_range", func() int32 {
		return hypervisor.HVGICGetSpiInterruptRange(&base, &count)
	})
	return base, count, err
}

// SendGICMSI sends a GIC MSI interrupt.
func SendGICMSI(address uint64, intid uint32) error {
	return call("hv_gic_send_msi", func() int32 {
		return hypervisor.HVGICSendMsi(address, intid)
	})
}

// GetGICStateData returns the Hypervisor.framework GIC state blob.
func GetGICStateData() ([]byte, error) {
	state, err := create("hv_gic_state_create", hypervisor.HVGICStateCreate)
	if err != nil {
		return nil, err
	}
	if state == nil {
		return nil, fmt.Errorf("hv_gic_state_create returned nil")
	}
	var size uintptr
	if err := call("hv_gic_state_get_size", func() int32 {
		return hypervisor.HVGICStateGetSize(state, &size)
	}); err != nil {
		return nil, err
	}
	if size == 0 {
		return nil, nil
	}
	data := make([]byte, size)
	if err := call("hv_gic_state_get_data", func() int32 {
		return hypervisor.HVGICStateGetData(state, unsafe.Pointer(&data[0]))
	}); err != nil {
		return nil, err
	}
	return data, nil
}

// SetGICStateData restores a Hypervisor.framework GIC state blob.
func SetGICStateData(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	return call("hv_gic_set_state", func() int32 {
		return hypervisor.HVGICSetState(unsafe.Pointer(&data[0]), uintptr(len(data)))
	})
}

// call wraps a Hypervisor.framework return-code function.
func call(name string, fn func() int32) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoveredError(name, r)
		}
	}()
	ret := fn()
	if ret != int32(hypervisor.HVSuccess) {
		return fmt.Errorf("%s failed: %s (%d)", name, hypervisor.HVReturn(ret), ret)
	}
	return nil
}

// create wraps a Hypervisor.framework object factory.
func create(name string, fn func() unsafe.Pointer) (handle unsafe.Pointer, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoveredError(name, r)
		}
	}()
	return fn(), nil
}

func recoveredError(name string, r any) error {
	if err, ok := r.(error); ok {
		return fmt.Errorf("%s: %w", name, err)
	}
	return fmt.Errorf("%s: %v", name, r)
}
