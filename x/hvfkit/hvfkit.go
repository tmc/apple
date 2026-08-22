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
	handle       hypervisor.HVVmConfig
	Capabilities Capabilities
	IPABits      uint32
	IPAGranule   hypervisor.HVIPAGranule
	EL2Enabled   bool
	Created      bool
}

// ConfigOption configures a VM configuration.
type ConfigOption func(*Config) error

// VCPUConfig is an opaque Hypervisor.framework vCPU configuration.
type VCPUConfig struct {
	handle hypervisor.HVVCPUConfig
}

// GICConfig is an opaque Hypervisor.framework GIC configuration.
type GICConfig struct {
	handle hypervisor.HVGICConfig
}

// VCPU is a Hypervisor.framework vCPU and its HVF-owned exit area.
type VCPU struct {
	ID   hypervisor.HVVCPU
	Exit *hypervisor.HVVCPUExit
}

// QueryCapabilities returns host Hypervisor.framework capabilities.
func QueryCapabilities() (Capabilities, error) {
	var caps Capabilities
	if err := call("hv_vm_get_max_vcpu_count", func() hypervisor.HVReturn {
		return hypervisor.HVVmGetMaxVCPUCount(&caps.MaxVCPUCount)
	}); err != nil {
		return Capabilities{}, err
	}
	if err := call("hv_vm_config_get_max_ipa_size", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigGetMaxIPASize(&caps.MaxIPABits)
	}); err != nil {
		return Capabilities{}, err
	}
	if err := call("hv_vm_config_get_el2_supported", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigGetEl2Supported(&caps.EL2Supported)
	}); err != nil {
		return Capabilities{}, err
	}
	return caps, nil
}

// NewConfig creates and inspects a Hypervisor.framework VM configuration.
func NewConfig(opts ...ConfigOption) (cfg *Config, err error) {
	handle, err := create("hv_vm_config_create", hypervisor.HVVmConfigCreate)
	if err != nil {
		return nil, err
	}
	if handle == nil {
		return nil, fmt.Errorf("hv_vm_config_create returned nil")
	}
	cfg = &Config{handle: handle}
	defer func() {
		if err != nil {
			err = joinReleaseError(err, cfg.Release())
			cfg = nil
		}
	}()
	if cfg.Capabilities, err = QueryCapabilities(); err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}
	if err := cfg.refresh(); err != nil {
		return nil, err
	}
	return cfg, nil
}

// WithIPASize sets the physical address size for a VM configuration.
func WithIPASize(bits uint32) ConfigOption {
	return func(c *Config) error {
		return c.setIPASize(bits)
	}
}

// WithIPAGranule sets the IPA granule for a VM configuration.
func WithIPAGranule(granule hypervisor.HVIPAGranule) ConfigOption {
	return func(c *Config) error {
		return c.SetIPAGranule(granule)
	}
}

// WithEL2 configures whether a VM configuration enables EL2.
func WithEL2(enabled bool) ConfigOption {
	return func(c *Config) error {
		return c.setEL2(enabled)
	}
}

// DefaultIPASize returns the default IPA size for new VM configurations.
func DefaultIPASize() (uint32, error) {
	var bits uint32
	err := call("hv_vm_config_get_default_ipa_size", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigGetDefaultIPASize(&bits)
	})
	return bits, err
}

// DefaultIPAGranule returns the default IPA granule for new VM configurations.
func DefaultIPAGranule() (hypervisor.HVIPAGranule, error) {
	var granule hypervisor.HVIPAGranule
	err := call("hv_vm_config_get_default_ipa_granule", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigGetDefaultIPAGranule(&granule)
	})
	return granule, err
}

// GetIPAGranule reads a VM configuration's IPA granule.
func (c *Config) GetIPAGranule() (hypervisor.HVIPAGranule, error) {
	var granule hypervisor.HVIPAGranule
	err := call("hv_vm_config_get_ipa_granule", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigGetIPAGranule(c.handle, &granule)
	})
	return granule, err
}

// SetIPAGranule writes a VM configuration's IPA granule.
func (c *Config) SetIPAGranule(granule hypervisor.HVIPAGranule) error {
	return call("hv_vm_config_set_ipa_granule", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigSetIPAGranule(c.handle, granule)
	})
}

// CreateVM creates a process VM with config.
func (c *Config) CreateVM() error {
	if err := call("hv_vm_create", func() hypervisor.HVReturn {
		return hypervisor.HVVmCreate(c.handle)
	}); err != nil {
		return err
	}
	c.Created = true
	return nil
}

func (c *Config) setIPASize(bits uint32) error {
	if bits == 0 {
		return nil
	}
	return call("hv_vm_config_set_ipa_size", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigSetIPASize(c.handle, bits)
	})
}

func (c *Config) setEL2(enabled bool) error {
	return call("hv_vm_config_set_el2_enabled", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigSetEl2Enabled(c.handle, enabled)
	})
}

func (c *Config) refresh() error {
	if err := call("hv_vm_config_get_ipa_size", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigGetIPASize(c.handle, &c.IPABits)
	}); err != nil {
		return err
	}
	if err := call("hv_vm_config_get_ipa_granule", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigGetIPAGranule(c.handle, &c.IPAGranule)
	}); err != nil {
		return err
	}
	return call("hv_vm_config_get_el2_enabled", func() hypervisor.HVReturn {
		return hypervisor.HVVmConfigGetEl2Enabled(c.handle, &c.EL2Enabled)
	})
}

// Release releases the retained VM configuration object.
func (c *Config) Release() error {
	if c == nil || c.handle == nil {
		return nil
	}
	handle := c.handle
	if err := osRelease(handle); err != nil {
		return err
	}
	c.handle = nil
	return nil
}

// Close destroys the process VM if this config created it, then releases the config.
func (c *Config) Close() error {
	if c == nil {
		return nil
	}
	var err error
	if c.Created {
		err = DestroyVM()
		c.Created = false
	}
	return joinReleaseError(err, c.Release())
}

// DestroyVM destroys the process VM.
func DestroyVM() error {
	return call("hv_vm_destroy", hypervisor.HVVmDestroy)
}

// MapMemory maps host memory into guest physical memory.
func MapMemory(addr unsafe.Pointer, ipa uint64, size uintptr, flags uint64) error {
	return call("hv_vm_map", func() hypervisor.HVReturn {
		return hypervisor.HVVmMap(addr, ipa, size, flags)
	})
}

// UnmapMemory removes a guest physical memory mapping.
func UnmapMemory(ipa uint64, size uintptr) error {
	return call("hv_vm_unmap", func() hypervisor.HVReturn {
		return hypervisor.HVVmUnmap(ipa, size)
	})
}

// AllocateMemory allocates Hypervisor.framework-owned memory.
func AllocateMemory(size uintptr, flags uint64) (unsafe.Pointer, error) {
	var addr unsafe.Pointer
	err := call("hv_vm_allocate", func() hypervisor.HVReturn {
		return hypervisor.HVVmAllocate(unsafe.Pointer(&addr), size, flags)
	})
	return addr, err
}

// DeallocateMemory releases Hypervisor.framework-owned memory.
func DeallocateMemory(addr unsafe.Pointer, size uintptr) error {
	return call("hv_vm_deallocate", func() hypervisor.HVReturn {
		return hypervisor.HVVmDeallocate(addr, size)
	})
}

// ProtectMemory changes guest physical memory permissions.
func ProtectMemory(ipa uint64, size uintptr, flags uint64) error {
	return call("hv_vm_protect", func() hypervisor.HVReturn {
		return hypervisor.HVVmProtect(ipa, size, flags)
	})
}

// NewVCPUConfig creates a Hypervisor.framework vCPU configuration.
func NewVCPUConfig() (*VCPUConfig, error) {
	handle, err := create("hv_vcpu_config_create", hypervisor.HVVCPUConfigCreate)
	if err != nil {
		return nil, err
	}
	if handle == nil {
		return nil, fmt.Errorf("hv_vcpu_config_create returned nil")
	}
	return &VCPUConfig{handle: handle}, nil
}

// Release releases the retained vCPU configuration object.
func (c *VCPUConfig) Release() error {
	if c == nil || c.handle == nil {
		return nil
	}
	handle := c.handle
	if err := osRelease(handle); err != nil {
		return err
	}
	c.handle = nil
	return nil
}

// FeatureReg reads a feature register from a vCPU configuration.
func (c *VCPUConfig) FeatureReg(reg hypervisor.HVFeatureReg) (uint64, error) {
	var value uint64
	err := call("hv_vcpu_config_get_feature_reg", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUConfigGetFeatureReg(c.handle, reg, &value)
	})
	return value, err
}

// CreateVCPU creates a vCPU and returns the HVF-owned exit area.
func CreateVCPU(config *VCPUConfig) (*VCPU, error) {
	var id hypervisor.HVVCPU
	var exit *hypervisor.HVVCPUExit
	var handle hypervisor.HVVCPUConfig
	if config != nil {
		handle = config.handle
	}
	err := call("hv_vcpu_create", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUCreate(&id, &exit, handle)
	})
	if err != nil {
		return nil, err
	}
	return &VCPU{ID: id, Exit: exit}, nil
}

// DestroyVCPU destroys a vCPU.
func DestroyVCPU(id hypervisor.HVVCPU) error {
	return call("hv_vcpu_destroy", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUDestroy(id)
	})
}

// RunVCPU enters a vCPU until the next exit.
func RunVCPU(id hypervisor.HVVCPU) error {
	return call("hv_vcpu_run", func() hypervisor.HVReturn {
		return hypervisor.HVVCPURun(id)
	})
}

// ExitVCPUs requests exits from the supplied vCPUs.
func ExitVCPUs(ids []hypervisor.HVVCPU) error {
	return call("hv_vcpus_exit", func() hypervisor.HVReturn {
		if len(ids) == 0 {
			return hypervisor.HVVcpusExit(nil, 0)
		}
		return hypervisor.HVVcpusExit(&ids[0], uint32(len(ids)))
	})
}

// GetExecTime returns the time a vCPU has spent executing.
func GetExecTime(vcpu hypervisor.HVVCPU) (uint64, error) {
	var tm uint64
	err := call("hv_vcpu_get_exec_time", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUGetExecTime(vcpu, &tm)
	})
	return tm, err
}

// GetPendingInterrupt reports whether an interrupt type is pending.
func GetPendingInterrupt(vcpu hypervisor.HVVCPU, typ hypervisor.HVInterruptType) (bool, error) {
	var pending bool
	err := call("hv_vcpu_get_pending_interrupt", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUGetPendingInterrupt(vcpu, typ, &pending)
	})
	return pending, err
}

// SetPendingInterrupt changes whether an interrupt type is pending.
func SetPendingInterrupt(vcpu hypervisor.HVVCPU, typ hypervisor.HVInterruptType, pending bool) error {
	return call("hv_vcpu_set_pending_interrupt", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUSetPendingInterrupt(vcpu, typ, pending)
	})
}

// GetReg reads a general-purpose vCPU register.
func GetReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVReg) (uint64, error) {
	var value uint64
	err := call("hv_vcpu_get_reg", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUGetReg(vcpu, reg, &value)
	})
	return value, err
}

// SetReg writes a general-purpose vCPU register.
func SetReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVReg, value uint64) error {
	return call("hv_vcpu_set_reg", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUSetReg(vcpu, reg, value)
	})
}

// GetSysReg reads a vCPU system register.
func GetSysReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVSysReg) (uint64, error) {
	var value uint64
	err := call("hv_vcpu_get_sys_reg", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUGetSysReg(vcpu, reg, &value)
	})
	return value, err
}

// SetSysReg writes a vCPU system register.
func SetSysReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVSysReg, value uint64) error {
	return call("hv_vcpu_set_sys_reg", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUSetSysReg(vcpu, reg, value)
	})
}

// GetSIMDFPReg reads a SIMD/FP register.
func GetSIMDFPReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVSIMDFPReg) ([16]byte, error) {
	var value [16]byte
	err := call("hv_vcpu_get_simd_fp_reg", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUGetSIMDFPReg(vcpu, reg, &value)
	})
	return value, err
}

// SetSIMDFPReg writes a SIMD/FP register.
func SetSIMDFPReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVSIMDFPReg, value [16]byte) error {
	return call("hv_vcpu_set_simd_fp_reg", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUSetSIMDFPReg(vcpu, reg, value)
	})
}

// GetVtimerMask reports whether the virtual timer is masked.
func GetVtimerMask(vcpu hypervisor.HVVCPU) (bool, error) {
	var masked bool
	err := call("hv_vcpu_get_vtimer_mask", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUGetVtimerMask(vcpu, &masked)
	})
	return masked, err
}

// GetVtimerOffset reads the virtual timer offset.
func GetVtimerOffset(vcpu hypervisor.HVVCPU) (uint64, error) {
	var offset uint64
	err := call("hv_vcpu_get_vtimer_offset", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUGetVtimerOffset(vcpu, &offset)
	})
	return offset, err
}

// SetVtimerMask masks or unmasks the virtual timer.
func SetVtimerMask(vcpu hypervisor.HVVCPU, masked bool) error {
	return call("hv_vcpu_set_vtimer_mask", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUSetVtimerMask(vcpu, masked)
	})
}

// SetVtimerOffset writes the virtual timer offset.
func SetVtimerOffset(vcpu hypervisor.HVVCPU, offset uint64) error {
	return call("hv_vcpu_set_vtimer_offset", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUSetVtimerOffset(vcpu, offset)
	})
}

// GetTrapDebugExceptions reports whether debug exceptions trap to the host.
func GetTrapDebugExceptions(vcpu hypervisor.HVVCPU) (bool, error) {
	var enabled bool
	err := call("hv_vcpu_get_trap_debug_exceptions", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUGetTrapDebugExceptions(vcpu, &enabled)
	})
	return enabled, err
}

// SetTrapDebugExceptions configures whether debug exceptions trap to the host.
func SetTrapDebugExceptions(vcpu hypervisor.HVVCPU, enabled bool) error {
	return call("hv_vcpu_set_trap_debug_exceptions", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUSetTrapDebugExceptions(vcpu, enabled)
	})
}

// GetTrapDebugRegAccesses reports whether debug register accesses trap to the host.
func GetTrapDebugRegAccesses(vcpu hypervisor.HVVCPU) (bool, error) {
	var enabled bool
	err := call("hv_vcpu_get_trap_debug_reg_accesses", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUGetTrapDebugRegAccesses(vcpu, &enabled)
	})
	return enabled, err
}

// SetTrapDebugRegAccesses configures whether debug register accesses trap to the host.
func SetTrapDebugRegAccesses(vcpu hypervisor.HVVCPU, enabled bool) error {
	return call("hv_vcpu_set_trap_debug_reg_accesses", func() hypervisor.HVReturn {
		return hypervisor.HVVCPUSetTrapDebugRegAccesses(vcpu, enabled)
	})
}

// CreateGIC configures and creates a GICv3 device.
func CreateGIC(distributorBase, redistributorBase uint64) error {
	config, err := NewGICConfig(distributorBase, redistributorBase)
	if err != nil {
		return err
	}
	defer func() { _ = config.Release() }()
	return config.Create()
}

// NewGICConfig creates a GICv3 configuration with distributor bases set.
func NewGICConfig(distributorBase, redistributorBase uint64) (cfg *GICConfig, err error) {
	handle, err := create("hv_gic_config_create", hypervisor.HVGICConfigCreate)
	if err != nil {
		return nil, err
	}
	if handle == nil {
		return nil, fmt.Errorf("hv_gic_config_create returned nil")
	}
	cfg = &GICConfig{handle: handle}
	defer func() {
		if err != nil {
			err = joinReleaseError(err, cfg.Release())
			cfg = nil
		}
	}()
	if err := call("hv_gic_config_set_distributor_base", func() hypervisor.HVReturn {
		return hypervisor.HVGICConfigSetDistributorBase(cfg.handle, distributorBase)
	}); err != nil {
		return nil, err
	}
	if err := call("hv_gic_config_set_redistributor_base", func() hypervisor.HVReturn {
		return hypervisor.HVGICConfigSetRedistributorBase(cfg.handle, redistributorBase)
	}); err != nil {
		return nil, err
	}
	return cfg, nil
}

// SetMSIInterruptRange configures the GIC MSI interrupt range.
func (c *GICConfig) SetMSIInterruptRange(base, count uint32) error {
	return call("hv_gic_config_set_msi_interrupt_range", func() hypervisor.HVReturn {
		return hypervisor.HVGICConfigSetMsiInterruptRange(c.handle, base, count)
	})
}

// SetMSIRegionBase configures the GIC MSI region base.
func (c *GICConfig) SetMSIRegionBase(base uint64) error {
	return call("hv_gic_config_set_msi_region_base", func() hypervisor.HVReturn {
		return hypervisor.HVGICConfigSetMsiRegionBase(c.handle, base)
	})
}

// Create creates a GICv3 device from config.
func (c *GICConfig) Create() error {
	return call("hv_gic_create", func() hypervisor.HVReturn {
		return hypervisor.HVGICCreate(c.handle)
	})
}

// Release releases the retained GIC configuration object.
func (c *GICConfig) Release() error {
	if c == nil || c.handle == nil {
		return nil
	}
	handle := c.handle
	if err := osRelease(handle); err != nil {
		return err
	}
	c.handle = nil
	return nil
}

// GetGICDistributorBaseAlignment returns the required distributor base alignment.
func GetGICDistributorBaseAlignment() (uintptr, error) {
	var alignment uintptr
	err := call("hv_gic_get_distributor_base_alignment", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetDistributorBaseAlignment(&alignment)
	})
	return alignment, err
}

// GetGICDistributorSize returns the GIC distributor region size.
func GetGICDistributorSize() (uintptr, error) {
	var size uintptr
	err := call("hv_gic_get_distributor_size", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetDistributorSize(&size)
	})
	return size, err
}

// GetGICRedistributorBaseAlignment returns the required redistributor base alignment.
func GetGICRedistributorBaseAlignment() (uintptr, error) {
	var alignment uintptr
	err := call("hv_gic_get_redistributor_base_alignment", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetRedistributorBaseAlignment(&alignment)
	})
	return alignment, err
}

// GetGICRedistributorSize returns one GIC redistributor region size.
func GetGICRedistributorSize() (uintptr, error) {
	var size uintptr
	err := call("hv_gic_get_redistributor_size", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetRedistributorSize(&size)
	})
	return size, err
}

// GetGICRedistributorRegionSize returns the full redistributor region size.
func GetGICRedistributorRegionSize() (uintptr, error) {
	var size uintptr
	err := call("hv_gic_get_redistributor_region_size", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetRedistributorRegionSize(&size)
	})
	return size, err
}

// GetGICRedistributorBase returns a vCPU redistributor base address.
func GetGICRedistributorBase(vcpu hypervisor.HVVCPU) (uint64, error) {
	var base uint64
	err := call("hv_gic_get_redistributor_base", func() hypervisor.HVReturn {
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
	return call("hv_gic_set_spi", func() hypervisor.HVReturn {
		return hypervisor.HVGICSetSpi(intid, level)
	})
}

// GetGICIntID returns the concrete GIC interrupt ID for a reserved interrupt.
func GetGICIntID(interrupt hypervisor.HVGICIntid) (uint32, error) {
	var intid uint32
	err := call("hv_gic_get_intid", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetIntid(interrupt, &intid)
	})
	return intid, err
}

// GetGICDistributorReg reads a GIC distributor register.
func GetGICDistributorReg(reg hypervisor.HVGICDistributorReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_distributor_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetDistributorReg(reg, &value)
	})
	return value, err
}

// SetGICDistributorReg writes a GIC distributor register.
func SetGICDistributorReg(reg hypervisor.HVGICDistributorReg, value uint64) error {
	return call("hv_gic_set_distributor_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICSetDistributorReg(reg, value)
	})
}

// GetGICRedistributorReg reads a GIC redistributor register.
func GetGICRedistributorReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVGICRedistributorReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_redistributor_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetRedistributorReg(vcpu, reg, &value)
	})
	return value, err
}

// SetGICRedistributorReg writes a GIC redistributor register.
func SetGICRedistributorReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVGICRedistributorReg, value uint64) error {
	return call("hv_gic_set_redistributor_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICSetRedistributorReg(vcpu, reg, value)
	})
}

// GetGICICCReg reads a GIC CPU-interface register.
func GetGICICCReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVGICIccReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_icc_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetIccReg(vcpu, reg, &value)
	})
	return value, err
}

// SetGICICCReg writes a GIC CPU-interface register.
func SetGICICCReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVGICIccReg, value uint64) error {
	return call("hv_gic_set_icc_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICSetIccReg(vcpu, reg, value)
	})
}

// GetGICICHReg reads a GIC hypervisor-control register.
func GetGICICHReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVGICIchReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_ich_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetIchReg(vcpu, reg, &value)
	})
	return value, err
}

// SetGICICHReg writes a GIC hypervisor-control register.
func SetGICICHReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVGICIchReg, value uint64) error {
	return call("hv_gic_set_ich_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICSetIchReg(vcpu, reg, value)
	})
}

// GetGICICVReg reads a GIC virtual CPU-interface register.
func GetGICICVReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVGICIcvReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_icv_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetIcvReg(vcpu, reg, &value)
	})
	return value, err
}

// SetGICICVReg writes a GIC virtual CPU-interface register.
func SetGICICVReg(vcpu hypervisor.HVVCPU, reg hypervisor.HVGICIcvReg, value uint64) error {
	return call("hv_gic_set_icv_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICSetIcvReg(vcpu, reg, value)
	})
}

// GetGICMSIReg reads a GIC MSI register.
func GetGICMSIReg(reg hypervisor.HVGICMsiReg) (uint64, error) {
	var value uint64
	err := call("hv_gic_get_msi_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetMsiReg(reg, &value)
	})
	return value, err
}

// SetGICMSIReg writes a GIC MSI register.
func SetGICMSIReg(reg hypervisor.HVGICMsiReg, value uint64) error {
	return call("hv_gic_set_msi_reg", func() hypervisor.HVReturn {
		return hypervisor.HVGICSetMsiReg(reg, value)
	})
}

// GetGICMSIRegionBaseAlignment returns the required MSI region base alignment.
func GetGICMSIRegionBaseAlignment() (uintptr, error) {
	var alignment uintptr
	err := call("hv_gic_get_msi_region_base_alignment", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetMsiRegionBaseAlignment(&alignment)
	})
	return alignment, err
}

// GetGICMSIRegionSize returns the GIC MSI region size.
func GetGICMSIRegionSize() (uintptr, error) {
	var size uintptr
	err := call("hv_gic_get_msi_region_size", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetMsiRegionSize(&size)
	})
	return size, err
}

// GetGICSPIInterruptRange returns the available shared peripheral interrupt range.
func GetGICSPIInterruptRange() (base, count uint32, err error) {
	err = call("hv_gic_get_spi_interrupt_range", func() hypervisor.HVReturn {
		return hypervisor.HVGICGetSpiInterruptRange(&base, &count)
	})
	return base, count, err
}

// SendGICMSI sends a GIC MSI interrupt.
func SendGICMSI(address uint64, intid uint32) error {
	return call("hv_gic_send_msi", func() hypervisor.HVReturn {
		return hypervisor.HVGICSendMsi(address, intid)
	})
}

// GetGICStateData returns the Hypervisor.framework GIC state blob.
func GetGICStateData() (data []byte, err error) {
	state, err := create("hv_gic_state_create", hypervisor.HVGICStateCreate)
	if err != nil {
		return nil, err
	}
	if state == nil {
		return nil, fmt.Errorf("hv_gic_state_create returned nil")
	}
	defer func() {
		err = joinReleaseError(err, osRelease(state))
	}()
	var size uintptr
	if err := call("hv_gic_state_get_size", func() hypervisor.HVReturn {
		return hypervisor.HVGICStateGetSize(state, &size)
	}); err != nil {
		return nil, err
	}
	if size == 0 {
		return nil, nil
	}
	data = make([]byte, size)
	if err := call("hv_gic_state_get_data", func() hypervisor.HVReturn {
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
	return call("hv_gic_set_state", func() hypervisor.HVReturn {
		return hypervisor.HVGICSetState(unsafe.Pointer(&data[0]), uintptr(len(data)))
	})
}

// call wraps a Hypervisor.framework return-code function.
func call(name string, fn func() hypervisor.HVReturn) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoveredError(name, r)
		}
	}()
	ret := fn()
	if ret != hypervisor.HVSuccess {
		return fmt.Errorf("%s failed: %s (%d)", name, ret, int32(ret))
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

func joinReleaseError(err, releaseErr error) error {
	if err == nil {
		return releaseErr
	}
	if releaseErr == nil {
		return err
	}
	return fmt.Errorf("%w; release: %v", err, releaseErr)
}
