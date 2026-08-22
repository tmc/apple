// Code generated from Apple documentation for Hypervisor. DO NOT EDIT.

package hypervisor

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("Hypervisor: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("Hypervisor: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("Hypervisor: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("Hypervisor: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _hVGICConfigCreate func() unsafe.Pointer
var _hVGICConfigCreateErr error

func tryHVGICConfigCreate() (unsafe.Pointer, error) {
	if _hVGICConfigCreate == nil {
		return nil, symbolCallError("hv_gic_config_create", "15.0", _hVGICConfigCreateErr)
	}
	return _hVGICConfigCreate(), nil
}

// HVGICConfigCreate creates a generic interrupt controller (GIC) configuration object.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_create()
func HVGICConfigCreate() unsafe.Pointer {
	result, callErr := tryHVGICConfigCreate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICConfigSetDistributorBase func(config unsafe.Pointer, distributor_base_address HVIPA) HVReturn
var _hVGICConfigSetDistributorBaseErr error

func tryHVGICConfigSetDistributorBase(config unsafe.Pointer, distributor_base_address HVIPA) (HVReturn, error) {
	if _hVGICConfigSetDistributorBase == nil {
		return *new(HVReturn), symbolCallError("hv_gic_config_set_distributor_base", "15.0", _hVGICConfigSetDistributorBaseErr)
	}
	return _hVGICConfigSetDistributorBase(config, distributor_base_address), nil
}

// HVGICConfigSetDistributorBase sets the generic interrupt controller (GIC) distributor region’s base address.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_set_distributor_base(_:_:)
func HVGICConfigSetDistributorBase(config unsafe.Pointer, distributor_base_address HVIPA) HVReturn {
	result, callErr := tryHVGICConfigSetDistributorBase(config, distributor_base_address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICConfigSetMsiInterruptRange func(config unsafe.Pointer, msi_intid_base uint32, msi_intid_count uint32) HVReturn
var _hVGICConfigSetMsiInterruptRangeErr error

func tryHVGICConfigSetMsiInterruptRange(config unsafe.Pointer, msi_intid_base uint32, msi_intid_count uint32) (HVReturn, error) {
	if _hVGICConfigSetMsiInterruptRange == nil {
		return *new(HVReturn), symbolCallError("hv_gic_config_set_msi_interrupt_range", "15.0", _hVGICConfigSetMsiInterruptRangeErr)
	}
	return _hVGICConfigSetMsiInterruptRange(config, msi_intid_base, msi_intid_count), nil
}

// HVGICConfigSetMsiInterruptRange sets the range of message signaled interrupts (MSIs) the generic interrupt controller supports.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_set_msi_interrupt_range(_:_:_:)
func HVGICConfigSetMsiInterruptRange(config unsafe.Pointer, msi_intid_base uint32, msi_intid_count uint32) HVReturn {
	result, callErr := tryHVGICConfigSetMsiInterruptRange(config, msi_intid_base, msi_intid_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICConfigSetMsiRegionBase func(config unsafe.Pointer, msi_region_base_address HVIPA) HVReturn
var _hVGICConfigSetMsiRegionBaseErr error

func tryHVGICConfigSetMsiRegionBase(config unsafe.Pointer, msi_region_base_address HVIPA) (HVReturn, error) {
	if _hVGICConfigSetMsiRegionBase == nil {
		return *new(HVReturn), symbolCallError("hv_gic_config_set_msi_region_base", "15.0", _hVGICConfigSetMsiRegionBaseErr)
	}
	return _hVGICConfigSetMsiRegionBase(config, msi_region_base_address), nil
}

// HVGICConfigSetMsiRegionBase sets the generic interrupt controllers message signaled interrupts (MSIs) region base address.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_set_msi_region_base(_:_:)
func HVGICConfigSetMsiRegionBase(config unsafe.Pointer, msi_region_base_address HVIPA) HVReturn {
	result, callErr := tryHVGICConfigSetMsiRegionBase(config, msi_region_base_address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICConfigSetRedistributorBase func(config unsafe.Pointer, redistributor_base_address HVIPA) HVReturn
var _hVGICConfigSetRedistributorBaseErr error

func tryHVGICConfigSetRedistributorBase(config unsafe.Pointer, redistributor_base_address HVIPA) (HVReturn, error) {
	if _hVGICConfigSetRedistributorBase == nil {
		return *new(HVReturn), symbolCallError("hv_gic_config_set_redistributor_base", "15.0", _hVGICConfigSetRedistributorBaseErr)
	}
	return _hVGICConfigSetRedistributorBase(config, redistributor_base_address), nil
}

// HVGICConfigSetRedistributorBase sets the generic interrupt controller (GIC) redistributor region base address.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_config_set_redistributor_base(_:_:)
func HVGICConfigSetRedistributorBase(config unsafe.Pointer, redistributor_base_address HVIPA) HVReturn {
	result, callErr := tryHVGICConfigSetRedistributorBase(config, redistributor_base_address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICCreate func(gic_config unsafe.Pointer) HVReturn
var _hVGICCreateErr error

func tryHVGICCreate(gic_config unsafe.Pointer) (HVReturn, error) {
	if _hVGICCreate == nil {
		return *new(HVReturn), symbolCallError("hv_gic_create", "15.0", _hVGICCreateErr)
	}
	return _hVGICCreate(gic_config), nil
}

// HVGICCreate creates a generic interrupt controller (GIC) v3 device for a VM configuration.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_create(_:)
func HVGICCreate(gic_config unsafe.Pointer) HVReturn {
	result, callErr := tryHVGICCreate(gic_config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetDistributorBaseAlignment func(distributor_base_alignment *uintptr) HVReturn
var _hVGICGetDistributorBaseAlignmentErr error

func tryHVGICGetDistributorBaseAlignment(distributor_base_alignment *uintptr) (HVReturn, error) {
	if _hVGICGetDistributorBaseAlignment == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_distributor_base_alignment", "15.0", _hVGICGetDistributorBaseAlignmentErr)
	}
	return _hVGICGetDistributorBaseAlignment(distributor_base_alignment), nil
}

// HVGICGetDistributorBaseAlignment gets the alignment for the base address of the generic interrupt controller (GIC) distributor region, in bytes.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_distributor_base_alignment(_:)
func HVGICGetDistributorBaseAlignment(distributor_base_alignment *uintptr) HVReturn {
	result, callErr := tryHVGICGetDistributorBaseAlignment(distributor_base_alignment)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetDistributorReg func(reg HVGICDistributorReg, value *uint64) HVReturn
var _hVGICGetDistributorRegErr error

func tryHVGICGetDistributorReg(reg HVGICDistributorReg, value *uint64) (HVReturn, error) {
	if _hVGICGetDistributorReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_distributor_reg", "15.0", _hVGICGetDistributorRegErr)
	}
	return _hVGICGetDistributorReg(reg, value), nil
}

// HVGICGetDistributorReg reads a generic interrupt controller (GIC) distributor register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_distributor_reg(_:_:)
func HVGICGetDistributorReg(reg HVGICDistributorReg, value *uint64) HVReturn {
	result, callErr := tryHVGICGetDistributorReg(reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetDistributorSize func(distributor_size *uintptr) HVReturn
var _hVGICGetDistributorSizeErr error

func tryHVGICGetDistributorSize(distributor_size *uintptr) (HVReturn, error) {
	if _hVGICGetDistributorSize == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_distributor_size", "15.0", _hVGICGetDistributorSizeErr)
	}
	return _hVGICGetDistributorSize(distributor_size), nil
}

// HVGICGetDistributorSize gets the size of the generic interrupt controller (GIC) distributor region, in bytes.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_distributor_size(_:)
func HVGICGetDistributorSize(distributor_size *uintptr) HVReturn {
	result, callErr := tryHVGICGetDistributorSize(distributor_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetIccReg func(vcpu HVVCPU, reg HVGICIccReg, value *uint64) HVReturn
var _hVGICGetIccRegErr error

func tryHVGICGetIccReg(vcpu HVVCPU, reg HVGICIccReg, value *uint64) (HVReturn, error) {
	if _hVGICGetIccReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_icc_reg", "15.0", _hVGICGetIccRegErr)
	}
	return _hVGICGetIccReg(vcpu, reg, value), nil
}

// HVGICGetIccReg reads a generic interrupt controller’s ICC CPU system register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_icc_reg(_:_:_:)
func HVGICGetIccReg(vcpu HVVCPU, reg HVGICIccReg, value *uint64) HVReturn {
	result, callErr := tryHVGICGetIccReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetIchReg func(vcpu HVVCPU, reg HVGICIchReg, value *uint64) HVReturn
var _hVGICGetIchRegErr error

func tryHVGICGetIchReg(vcpu HVVCPU, reg HVGICIchReg, value *uint64) (HVReturn, error) {
	if _hVGICGetIchReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_ich_reg", "15.0", _hVGICGetIchRegErr)
	}
	return _hVGICGetIchReg(vcpu, reg, value), nil
}

// HVGICGetIchReg reads a generic interrupt controller’s (GIC) ICH virtualization control system register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_ich_reg(_:_:_:)
func HVGICGetIchReg(vcpu HVVCPU, reg HVGICIchReg, value *uint64) HVReturn {
	result, callErr := tryHVGICGetIchReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetIcvReg func(vcpu HVVCPU, reg HVGICIcvReg, value *uint64) HVReturn
var _hVGICGetIcvRegErr error

func tryHVGICGetIcvReg(vcpu HVVCPU, reg HVGICIcvReg, value *uint64) (HVReturn, error) {
	if _hVGICGetIcvReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_icv_reg", "15.0", _hVGICGetIcvRegErr)
	}
	return _hVGICGetIcvReg(vcpu, reg, value), nil
}

// HVGICGetIcvReg writes a generic interrupt controller’s (GIC) ICV system register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_icv_reg(_:_:_:)
func HVGICGetIcvReg(vcpu HVVCPU, reg HVGICIcvReg, value *uint64) HVReturn {
	result, callErr := tryHVGICGetIcvReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetIntid func(interrupt HVGICIntid, intid *uint32) HVReturn
var _hVGICGetIntidErr error

func tryHVGICGetIntid(interrupt HVGICIntid, intid *uint32) (HVReturn, error) {
	if _hVGICGetIntid == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_intid", "15.0", _hVGICGetIntidErr)
	}
	return _hVGICGetIntid(interrupt, intid), nil
}

// HVGICGetIntid gets the interrupt ID for reserved interrupts.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_intid(_:_:)
func HVGICGetIntid(interrupt HVGICIntid, intid *uint32) HVReturn {
	result, callErr := tryHVGICGetIntid(interrupt, intid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetMsiReg func(reg HVGICMsiReg, value *uint64) HVReturn
var _hVGICGetMsiRegErr error

func tryHVGICGetMsiReg(reg HVGICMsiReg, value *uint64) (HVReturn, error) {
	if _hVGICGetMsiReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_msi_reg", "15.0", _hVGICGetMsiRegErr)
	}
	return _hVGICGetMsiReg(reg, value), nil
}

// HVGICGetMsiReg reads a generic interrupt controller (GIC) distributor message signaled interrupt (MSI) register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_msi_reg(_:_:)
func HVGICGetMsiReg(reg HVGICMsiReg, value *uint64) HVReturn {
	result, callErr := tryHVGICGetMsiReg(reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetMsiRegionBaseAlignment func(msi_region_base_alignment *uintptr) HVReturn
var _hVGICGetMsiRegionBaseAlignmentErr error

func tryHVGICGetMsiRegionBaseAlignment(msi_region_base_alignment *uintptr) (HVReturn, error) {
	if _hVGICGetMsiRegionBaseAlignment == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_msi_region_base_alignment", "15.0", _hVGICGetMsiRegionBaseAlignmentErr)
	}
	return _hVGICGetMsiRegionBaseAlignment(msi_region_base_alignment), nil
}

// HVGICGetMsiRegionBaseAlignment gets the alignment, in bytes, for the base address of the generic interrupt controller’s message signaled interrupts (MSI) region.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_msi_region_base_alignment(_:)
func HVGICGetMsiRegionBaseAlignment(msi_region_base_alignment *uintptr) HVReturn {
	result, callErr := tryHVGICGetMsiRegionBaseAlignment(msi_region_base_alignment)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetMsiRegionSize func(msi_region_size *uintptr) HVReturn
var _hVGICGetMsiRegionSizeErr error

func tryHVGICGetMsiRegionSize(msi_region_size *uintptr) (HVReturn, error) {
	if _hVGICGetMsiRegionSize == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_msi_region_size", "15.0", _hVGICGetMsiRegionSizeErr)
	}
	return _hVGICGetMsiRegionSize(msi_region_size), nil
}

// HVGICGetMsiRegionSize gets the size in bytes of the generic interrupt controller’s (GIC) message signaled interrupts (MSI) region.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_msi_region_size(_:)
func HVGICGetMsiRegionSize(msi_region_size *uintptr) HVReturn {
	result, callErr := tryHVGICGetMsiRegionSize(msi_region_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetRedistributorBase func(vcpu HVVCPU, redistributor_base_address *HVIPA) HVReturn
var _hVGICGetRedistributorBaseErr error

func tryHVGICGetRedistributorBase(vcpu HVVCPU, redistributor_base_address *HVIPA) (HVReturn, error) {
	if _hVGICGetRedistributorBase == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_redistributor_base", "15.0", _hVGICGetRedistributorBaseErr)
	}
	return _hVGICGetRedistributorBase(vcpu, redistributor_base_address), nil
}

// HVGICGetRedistributorBase gets the redistributor base guest physical address for the given vCPU.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_base(_:_:)
func HVGICGetRedistributorBase(vcpu HVVCPU, redistributor_base_address *HVIPA) HVReturn {
	result, callErr := tryHVGICGetRedistributorBase(vcpu, redistributor_base_address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetRedistributorBaseAlignment func(redistributor_base_alignment *uintptr) HVReturn
var _hVGICGetRedistributorBaseAlignmentErr error

func tryHVGICGetRedistributorBaseAlignment(redistributor_base_alignment *uintptr) (HVReturn, error) {
	if _hVGICGetRedistributorBaseAlignment == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_redistributor_base_alignment", "15.0", _hVGICGetRedistributorBaseAlignmentErr)
	}
	return _hVGICGetRedistributorBaseAlignment(redistributor_base_alignment), nil
}

// HVGICGetRedistributorBaseAlignment gets the alignment for the base address of the generic interrupt controller (GIC) redistributor region, in bytes.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_base_alignment(_:)
func HVGICGetRedistributorBaseAlignment(redistributor_base_alignment *uintptr) HVReturn {
	result, callErr := tryHVGICGetRedistributorBaseAlignment(redistributor_base_alignment)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetRedistributorReg func(vcpu HVVCPU, reg HVGICRedistributorReg, value *uint64) HVReturn
var _hVGICGetRedistributorRegErr error

func tryHVGICGetRedistributorReg(vcpu HVVCPU, reg HVGICRedistributorReg, value *uint64) (HVReturn, error) {
	if _hVGICGetRedistributorReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_redistributor_reg", "15.0", _hVGICGetRedistributorRegErr)
	}
	return _hVGICGetRedistributorReg(vcpu, reg, value), nil
}

// HVGICGetRedistributorReg read a generic interrupt controller (GIC) redistributor register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_reg(_:_:_:)
func HVGICGetRedistributorReg(vcpu HVVCPU, reg HVGICRedistributorReg, value *uint64) HVReturn {
	result, callErr := tryHVGICGetRedistributorReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetRedistributorRegionSize func(redistributor_region_size *uintptr) HVReturn
var _hVGICGetRedistributorRegionSizeErr error

func tryHVGICGetRedistributorRegionSize(redistributor_region_size *uintptr) (HVReturn, error) {
	if _hVGICGetRedistributorRegionSize == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_redistributor_region_size", "15.0", _hVGICGetRedistributorRegionSizeErr)
	}
	return _hVGICGetRedistributorRegionSize(redistributor_region_size), nil
}

// HVGICGetRedistributorRegionSize gets the total size in bytes of the generic interrupt controller (GIC) redistributor region.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_region_size(_:)
func HVGICGetRedistributorRegionSize(redistributor_region_size *uintptr) HVReturn {
	result, callErr := tryHVGICGetRedistributorRegionSize(redistributor_region_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetRedistributorSize func(redistributor_size *uintptr) HVReturn
var _hVGICGetRedistributorSizeErr error

func tryHVGICGetRedistributorSize(redistributor_size *uintptr) (HVReturn, error) {
	if _hVGICGetRedistributorSize == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_redistributor_size", "15.0", _hVGICGetRedistributorSizeErr)
	}
	return _hVGICGetRedistributorSize(redistributor_size), nil
}

// HVGICGetRedistributorSize gets the size in bytes of a single generic interrupt controller (GIC) redistributor.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_redistributor_size(_:)
func HVGICGetRedistributorSize(redistributor_size *uintptr) HVReturn {
	result, callErr := tryHVGICGetRedistributorSize(redistributor_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICGetSpiInterruptRange func(spi_intid_base *uint32, spi_intid_count *uint32) HVReturn
var _hVGICGetSpiInterruptRangeErr error

func tryHVGICGetSpiInterruptRange(spi_intid_base *uint32, spi_intid_count *uint32) (HVReturn, error) {
	if _hVGICGetSpiInterruptRange == nil {
		return *new(HVReturn), symbolCallError("hv_gic_get_spi_interrupt_range", "15.0", _hVGICGetSpiInterruptRangeErr)
	}
	return _hVGICGetSpiInterruptRange(spi_intid_base, spi_intid_count), nil
}

// HVGICGetSpiInterruptRange gets the range of shared peripheral interrupts (SPIs) the generic interrupt controller supports.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_get_spi_interrupt_range(_:_:)
func HVGICGetSpiInterruptRange(spi_intid_base *uint32, spi_intid_count *uint32) HVReturn {
	result, callErr := tryHVGICGetSpiInterruptRange(spi_intid_base, spi_intid_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICReset func() HVReturn
var _hVGICResetErr error

func tryHVGICReset() (HVReturn, error) {
	if _hVGICReset == nil {
		return *new(HVReturn), symbolCallError("hv_gic_reset", "15.0", _hVGICResetErr)
	}
	return _hVGICReset(), nil
}

// HVGICReset resets the generic interrupt controller (GIC) device.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_reset()
func HVGICReset() HVReturn {
	result, callErr := tryHVGICReset()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICSendMsi func(address HVIPA, intid uint32) HVReturn
var _hVGICSendMsiErr error

func tryHVGICSendMsi(address HVIPA, intid uint32) (HVReturn, error) {
	if _hVGICSendMsi == nil {
		return *new(HVReturn), symbolCallError("hv_gic_send_msi", "15.0", _hVGICSendMsiErr)
	}
	return _hVGICSendMsi(address, intid), nil
}

// HVGICSendMsi sends a message signaled interrupt (MSI).
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_send_msi(_:_:)
func HVGICSendMsi(address HVIPA, intid uint32) HVReturn {
	result, callErr := tryHVGICSendMsi(address, intid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICSetDistributorReg func(reg HVGICDistributorReg, value uint64) HVReturn
var _hVGICSetDistributorRegErr error

func tryHVGICSetDistributorReg(reg HVGICDistributorReg, value uint64) (HVReturn, error) {
	if _hVGICSetDistributorReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_set_distributor_reg", "15.0", _hVGICSetDistributorRegErr)
	}
	return _hVGICSetDistributorReg(reg, value), nil
}

// HVGICSetDistributorReg writes the provided value to a generic interrupt controller (GIC) distributor register you specify.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_distributor_reg(_:_:)
func HVGICSetDistributorReg(reg HVGICDistributorReg, value uint64) HVReturn {
	result, callErr := tryHVGICSetDistributorReg(reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICSetIccReg func(vcpu HVVCPU, reg HVGICIccReg, value uint64) HVReturn
var _hVGICSetIccRegErr error

func tryHVGICSetIccReg(vcpu HVVCPU, reg HVGICIccReg, value uint64) (HVReturn, error) {
	if _hVGICSetIccReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_set_icc_reg", "15.0", _hVGICSetIccRegErr)
	}
	return _hVGICSetIccReg(vcpu, reg, value), nil
}

// HVGICSetIccReg writes to a generic interrupt controller (GIC) ICC cpu system register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_icc_reg(_:_:_:)
func HVGICSetIccReg(vcpu HVVCPU, reg HVGICIccReg, value uint64) HVReturn {
	result, callErr := tryHVGICSetIccReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICSetIchReg func(vcpu HVVCPU, reg HVGICIchReg, value uint64) HVReturn
var _hVGICSetIchRegErr error

func tryHVGICSetIchReg(vcpu HVVCPU, reg HVGICIchReg, value uint64) (HVReturn, error) {
	if _hVGICSetIchReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_set_ich_reg", "15.0", _hVGICSetIchRegErr)
	}
	return _hVGICSetIchReg(vcpu, reg, value), nil
}

// HVGICSetIchReg writes to a generic interrupt controller (GIC) ICH virtualization control system register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_ich_reg(_:_:_:)
func HVGICSetIchReg(vcpu HVVCPU, reg HVGICIchReg, value uint64) HVReturn {
	result, callErr := tryHVGICSetIchReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICSetIcvReg func(vcpu HVVCPU, reg HVGICIcvReg, value uint64) HVReturn
var _hVGICSetIcvRegErr error

func tryHVGICSetIcvReg(vcpu HVVCPU, reg HVGICIcvReg, value uint64) (HVReturn, error) {
	if _hVGICSetIcvReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_set_icv_reg", "15.0", _hVGICSetIcvRegErr)
	}
	return _hVGICSetIcvReg(vcpu, reg, value), nil
}

// HVGICSetIcvReg writes to a generic interrupt controller (GIC) ICV system register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_icv_reg(_:_:_:)
func HVGICSetIcvReg(vcpu HVVCPU, reg HVGICIcvReg, value uint64) HVReturn {
	result, callErr := tryHVGICSetIcvReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICSetMsiReg func(reg HVGICMsiReg, value uint64) HVReturn
var _hVGICSetMsiRegErr error

func tryHVGICSetMsiReg(reg HVGICMsiReg, value uint64) (HVReturn, error) {
	if _hVGICSetMsiReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_set_msi_reg", "15.0", _hVGICSetMsiRegErr)
	}
	return _hVGICSetMsiReg(reg, value), nil
}

// HVGICSetMsiReg writes to a generic interrupt controller distributor message signaled interrupt (MSI) register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_msi_reg(_:_:)
func HVGICSetMsiReg(reg HVGICMsiReg, value uint64) HVReturn {
	result, callErr := tryHVGICSetMsiReg(reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICSetRedistributorReg func(vcpu HVVCPU, reg HVGICRedistributorReg, value uint64) HVReturn
var _hVGICSetRedistributorRegErr error

func tryHVGICSetRedistributorReg(vcpu HVVCPU, reg HVGICRedistributorReg, value uint64) (HVReturn, error) {
	if _hVGICSetRedistributorReg == nil {
		return *new(HVReturn), symbolCallError("hv_gic_set_redistributor_reg", "15.0", _hVGICSetRedistributorRegErr)
	}
	return _hVGICSetRedistributorReg(vcpu, reg, value), nil
}

// HVGICSetRedistributorReg writes to a GIC redistributor register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_redistributor_reg(_:_:_:)
func HVGICSetRedistributorReg(vcpu HVVCPU, reg HVGICRedistributorReg, value uint64) HVReturn {
	result, callErr := tryHVGICSetRedistributorReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICSetSpi func(intid uint32, level bool) HVReturn
var _hVGICSetSpiErr error

func tryHVGICSetSpi(intid uint32, level bool) (HVReturn, error) {
	if _hVGICSetSpi == nil {
		return *new(HVReturn), symbolCallError("hv_gic_set_spi", "15.0", _hVGICSetSpiErr)
	}
	return _hVGICSetSpi(intid, level), nil
}

// HVGICSetSpi triggers a shared peripheral interrupt (SPI).
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_spi(_:_:)
func HVGICSetSpi(intid uint32, level bool) HVReturn {
	result, callErr := tryHVGICSetSpi(intid, level)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICSetState func(gic_state_data unsafe.Pointer, gic_state_size uintptr) HVReturn
var _hVGICSetStateErr error

func tryHVGICSetState(gic_state_data unsafe.Pointer, gic_state_size uintptr) (HVReturn, error) {
	if _hVGICSetState == nil {
		return *new(HVReturn), symbolCallError("hv_gic_set_state", "15.0", _hVGICSetStateErr)
	}
	return _hVGICSetState(gic_state_data, gic_state_size), nil
}

// HVGICSetState sets the state of a generic interrupt controller (GIC) device.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_set_state(_:_:)
func HVGICSetState(gic_state_data unsafe.Pointer, gic_state_size uintptr) HVReturn {
	result, callErr := tryHVGICSetState(gic_state_data, gic_state_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICStateCreate func() unsafe.Pointer
var _hVGICStateCreateErr error

func tryHVGICStateCreate() (unsafe.Pointer, error) {
	if _hVGICStateCreate == nil {
		return nil, symbolCallError("hv_gic_state_create", "15.0", _hVGICStateCreateErr)
	}
	return _hVGICStateCreate(), nil
}

// HVGICStateCreate create a generic interrupt controller (GIC) state object.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_state_create()
func HVGICStateCreate() unsafe.Pointer {
	result, callErr := tryHVGICStateCreate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICStateGetData func(state unsafe.Pointer, gic_state_data unsafe.Pointer) HVReturn
var _hVGICStateGetDataErr error

func tryHVGICStateGetData(state unsafe.Pointer, gic_state_data unsafe.Pointer) (HVReturn, error) {
	if _hVGICStateGetData == nil {
		return *new(HVReturn), symbolCallError("hv_gic_state_get_data", "15.0", _hVGICStateGetDataErr)
	}
	return _hVGICStateGetData(state, gic_state_data), nil
}

// HVGICStateGetData gets the state data for generic interrupt controller (GIC).
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_state_get_data(_:_:)
func HVGICStateGetData(state unsafe.Pointer, gic_state_data unsafe.Pointer) HVReturn {
	result, callErr := tryHVGICStateGetData(state, gic_state_data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVGICStateGetSize func(state unsafe.Pointer, gic_state_size *uintptr) HVReturn
var _hVGICStateGetSizeErr error

func tryHVGICStateGetSize(state unsafe.Pointer, gic_state_size *uintptr) (HVReturn, error) {
	if _hVGICStateGetSize == nil {
		return *new(HVReturn), symbolCallError("hv_gic_state_get_size", "15.0", _hVGICStateGetSizeErr)
	}
	return _hVGICStateGetSize(state, gic_state_size), nil
}

// HVGICStateGetSize gets the size of the buffer required for generic interrupt controller (GIC) state.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_gic_state_get_size(_:_:)
func HVGICStateGetSize(state unsafe.Pointer, gic_state_size *uintptr) HVReturn {
	result, callErr := tryHVGICStateGetSize(state, gic_state_size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVSMEConfigGetMaxSvlBytes func(value *uintptr) HVReturn
var _hVSMEConfigGetMaxSvlBytesErr error

func tryHVSMEConfigGetMaxSvlBytes(value *uintptr) (HVReturn, error) {
	if _hVSMEConfigGetMaxSvlBytes == nil {
		return *new(HVReturn), symbolCallError("hv_sme_config_get_max_svl_bytes", "15.2", _hVSMEConfigGetMaxSvlBytesErr)
	}
	return _hVSMEConfigGetMaxSvlBytes(value), nil
}

// HVSMEConfigGetMaxSvlBytes.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_sme_config_get_max_svl_bytes(_:)
func HVSMEConfigGetMaxSvlBytes(value *uintptr) HVReturn {
	result, callErr := tryHVSMEConfigGetMaxSvlBytes(value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUConfigCreate func() unsafe.Pointer
var _hVVCPUConfigCreateErr error

func tryHVVCPUConfigCreate() (unsafe.Pointer, error) {
	if _hVVCPUConfigCreate == nil {
		return nil, symbolCallError("hv_vcpu_config_create", "11.0", _hVVCPUConfigCreateErr)
	}
	return _hVVCPUConfigCreate(), nil
}

// HVVCPUConfigCreate creates a vCPU configuration object.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_config_create()
func HVVCPUConfigCreate() unsafe.Pointer {
	result, callErr := tryHVVCPUConfigCreate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUConfigGetCcsidrEl1SysRegValues func(config unsafe.Pointer, cache_type HVCacheType, values uint64) HVReturn
var _hVVCPUConfigGetCcsidrEl1SysRegValuesErr error

func tryHVVCPUConfigGetCcsidrEl1SysRegValues(config unsafe.Pointer, cache_type HVCacheType, values uint64) (HVReturn, error) {
	if _hVVCPUConfigGetCcsidrEl1SysRegValues == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_config_get_ccsidr_el1_sys_reg_values", "11.0", _hVVCPUConfigGetCcsidrEl1SysRegValuesErr)
	}
	return _hVVCPUConfigGetCcsidrEl1SysRegValues(config, cache_type, values), nil
}

// HVVCPUConfigGetCcsidrEl1SysRegValues returns the Cache Size ID Register (CCSIDR_EL1) values for the vCPU configuration and cache type you specify.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_config_get_ccsidr_el1_sys_reg_values(_:_:_:)
func HVVCPUConfigGetCcsidrEl1SysRegValues(config unsafe.Pointer, cache_type HVCacheType, values uint64) HVReturn {
	result, callErr := tryHVVCPUConfigGetCcsidrEl1SysRegValues(config, cache_type, values)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUConfigGetFeatureReg func(config unsafe.Pointer, feature_reg HVFeatureReg, value *uint64) HVReturn
var _hVVCPUConfigGetFeatureRegErr error

func tryHVVCPUConfigGetFeatureReg(config unsafe.Pointer, feature_reg HVFeatureReg, value *uint64) (HVReturn, error) {
	if _hVVCPUConfigGetFeatureReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_config_get_feature_reg", "11.0", _hVVCPUConfigGetFeatureRegErr)
	}
	return _hVVCPUConfigGetFeatureReg(config, feature_reg, value), nil
}

// HVVCPUConfigGetFeatureReg gets the value of a feature register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_config_get_feature_reg(_:_:_:)
func HVVCPUConfigGetFeatureReg(config unsafe.Pointer, feature_reg HVFeatureReg, value *uint64) HVReturn {
	result, callErr := tryHVVCPUConfigGetFeatureReg(config, feature_reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUCreate func(vcpu *HVVCPU, exit **HVVCPUExit, config unsafe.Pointer) HVReturn
var _hVVCPUCreateErr error

func tryHVVCPUCreate(vcpu *HVVCPU, exit **HVVCPUExit, config unsafe.Pointer) (HVReturn, error) {
	if _hVVCPUCreate == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_create", "11.0", _hVVCPUCreateErr)
	}
	return _hVVCPUCreate(vcpu, exit, config), nil
}

// HVVCPUCreate creates a vCPU instance for the current thread.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_create(_:_:_:)
func HVVCPUCreate(vcpu *HVVCPU, exit **HVVCPUExit, config unsafe.Pointer) HVReturn {
	result, callErr := tryHVVCPUCreate(vcpu, exit, config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUDestroy func(vcpu HVVCPU) HVReturn
var _hVVCPUDestroyErr error

func tryHVVCPUDestroy(vcpu HVVCPU) (HVReturn, error) {
	if _hVVCPUDestroy == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_destroy", "11.0", _hVVCPUDestroyErr)
	}
	return _hVVCPUDestroy(vcpu), nil
}

// HVVCPUDestroy destroys the vCPU instance associated with the current thread.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_destroy(_:)
func HVVCPUDestroy(vcpu HVVCPU) HVReturn {
	result, callErr := tryHVVCPUDestroy(vcpu)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetExecTime func(vcpu HVVCPU, time *uint64) HVReturn
var _hVVCPUGetExecTimeErr error

func tryHVVCPUGetExecTime(vcpu HVVCPU, time *uint64) (HVReturn, error) {
	if _hVVCPUGetExecTime == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_exec_time", "11.0", _hVVCPUGetExecTimeErr)
	}
	return _hVVCPUGetExecTime(vcpu, time), nil
}

// HVVCPUGetExecTime returns, by reference, the cumulative execution time of a vCPU, in nanoseconds.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_exec_time(_:_:)
func HVVCPUGetExecTime(vcpu HVVCPU, time *uint64) HVReturn {
	result, callErr := tryHVVCPUGetExecTime(vcpu, time)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetPendingInterrupt func(vcpu HVVCPU, type_ HVInterruptType, pending *bool) HVReturn
var _hVVCPUGetPendingInterruptErr error

func tryHVVCPUGetPendingInterrupt(vcpu HVVCPU, type_ HVInterruptType, pending *bool) (HVReturn, error) {
	if _hVVCPUGetPendingInterrupt == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_pending_interrupt", "11.0", _hVVCPUGetPendingInterruptErr)
	}
	return _hVVCPUGetPendingInterrupt(vcpu, type_, pending), nil
}

// HVVCPUGetPendingInterrupt gets pending interrupts for a vCPU.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_pending_interrupt(_:_:_:)
func HVVCPUGetPendingInterrupt(vcpu HVVCPU, type_ HVInterruptType, pending *bool) HVReturn {
	result, callErr := tryHVVCPUGetPendingInterrupt(vcpu, type_, pending)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetReg func(vcpu HVVCPU, reg HVReg, value *uint64) HVReturn
var _hVVCPUGetRegErr error

func tryHVVCPUGetReg(vcpu HVVCPU, reg HVReg, value *uint64) (HVReturn, error) {
	if _hVVCPUGetReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_reg", "11.0", _hVVCPUGetRegErr)
	}
	return _hVVCPUGetReg(vcpu, reg, value), nil
}

// HVVCPUGetReg gets the current value of a vCPU register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_reg(_:_:_:)
func HVVCPUGetReg(vcpu HVVCPU, reg HVReg, value *uint64) HVReturn {
	result, callErr := tryHVVCPUGetReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetSIMDFPReg func(vcpu HVVCPU, reg HVSIMDFPReg, value *[16]byte) HVReturn
var _hVVCPUGetSIMDFPRegErr error

func tryHVVCPUGetSIMDFPReg(vcpu HVVCPU, reg HVSIMDFPReg, value *[16]byte) (HVReturn, error) {
	if _hVVCPUGetSIMDFPReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_simd_fp_reg", "11.0", _hVVCPUGetSIMDFPRegErr)
	}
	return _hVVCPUGetSIMDFPReg(vcpu, reg, value), nil
}

// HVVCPUGetSIMDFPReg gets the current value of a vCPU SIMD and FP register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_simd_fp_reg(_:_:_:)
func HVVCPUGetSIMDFPReg(vcpu HVVCPU, reg HVSIMDFPReg, value *[16]byte) HVReturn {
	result, callErr := tryHVVCPUGetSIMDFPReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetSMEPReg func(vcpu HVVCPU, reg HVSMEPReg, value *byte, length uintptr) HVReturn
var _hVVCPUGetSMEPRegErr error

func tryHVVCPUGetSMEPReg(vcpu HVVCPU, reg HVSMEPReg, value []byte, length uintptr) (HVReturn, error) {
	if _hVVCPUGetSMEPReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_sme_p_reg", "15.2", _hVVCPUGetSMEPRegErr)
	}
	return _hVVCPUGetSMEPReg(vcpu, reg, unsafe.SliceData(value), length), nil
}

// HVVCPUGetSMEPReg returns the value of a vCPU P predicate register in streaming Scalable Vector Extension (SVE) mode.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_p_reg(_:_:_:_:)
func HVVCPUGetSMEPReg(vcpu HVVCPU, reg HVSMEPReg, value []byte, length uintptr) HVReturn {
	result, callErr := tryHVVCPUGetSMEPReg(vcpu, reg, value, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetSMEState func(vcpu HVVCPU, sme_state *HVVCPUSMEState) HVReturn
var _hVVCPUGetSMEStateErr error

func tryHVVCPUGetSMEState(vcpu HVVCPU, sme_state *HVVCPUSMEState) (HVReturn, error) {
	if _hVVCPUGetSMEState == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_sme_state", "15.2", _hVVCPUGetSMEStateErr)
	}
	return _hVVCPUGetSMEState(vcpu, sme_state), nil
}

// HVVCPUGetSMEState gets the current Scalable Matrix Extension (SME) state.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_state(_:_:)
func HVVCPUGetSMEState(vcpu HVVCPU, sme_state *HVVCPUSMEState) HVReturn {
	result, callErr := tryHVVCPUGetSMEState(vcpu, sme_state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetSMEZReg func(vcpu HVVCPU, reg HVSMEZReg, value *byte, length uintptr) HVReturn
var _hVVCPUGetSMEZRegErr error

func tryHVVCPUGetSMEZReg(vcpu HVVCPU, reg HVSMEZReg, value []byte, length uintptr) (HVReturn, error) {
	if _hVVCPUGetSMEZReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_sme_z_reg", "15.2", _hVVCPUGetSMEZRegErr)
	}
	return _hVVCPUGetSMEZReg(vcpu, reg, unsafe.SliceData(value), length), nil
}

// HVVCPUGetSMEZReg returns the value of a vCPU Z vector register in streaming Scalable Vector Extension (SVE) mode.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_z_reg(_:_:_:_:)
func HVVCPUGetSMEZReg(vcpu HVVCPU, reg HVSMEZReg, value []byte, length uintptr) HVReturn {
	result, callErr := tryHVVCPUGetSMEZReg(vcpu, reg, value, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetSMEZaReg func(vcpu HVVCPU, value *byte, length uintptr) HVReturn
var _hVVCPUGetSMEZaRegErr error

func tryHVVCPUGetSMEZaReg(vcpu HVVCPU, value []byte, length uintptr) (HVReturn, error) {
	if _hVVCPUGetSMEZaReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_sme_za_reg", "15.2", _hVVCPUGetSMEZaRegErr)
	}
	return _hVVCPUGetSMEZaReg(vcpu, unsafe.SliceData(value), length), nil
}

// HVVCPUGetSMEZaReg returns the value of the vCPU ZA matrix register in streaming Scalable Vector Extension (SVE) mode.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_za_reg(_:_:_:)
func HVVCPUGetSMEZaReg(vcpu HVVCPU, value []byte, length uintptr) HVReturn {
	result, callErr := tryHVVCPUGetSMEZaReg(vcpu, value, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetSMEZt0Reg func(vcpu HVVCPU, value *[64]byte) HVReturn
var _hVVCPUGetSMEZt0RegErr error

func tryHVVCPUGetSMEZt0Reg(vcpu HVVCPU, value *[64]byte) (HVReturn, error) {
	if _hVVCPUGetSMEZt0Reg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_sme_zt0_reg", "15.2", _hVVCPUGetSMEZt0RegErr)
	}
	return _hVVCPUGetSMEZt0Reg(vcpu, value), nil
}

// HVVCPUGetSMEZt0Reg returns the current value of the vCPU ZT0 register in streaming Scalable Vector Extension (SVE) mode.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sme_zt0_reg(_:_:)
func HVVCPUGetSMEZt0Reg(vcpu HVVCPU, value *[64]byte) HVReturn {
	result, callErr := tryHVVCPUGetSMEZt0Reg(vcpu, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetSysReg func(vcpu HVVCPU, reg HVSysReg, value *uint64) HVReturn
var _hVVCPUGetSysRegErr error

func tryHVVCPUGetSysReg(vcpu HVVCPU, reg HVSysReg, value *uint64) (HVReturn, error) {
	if _hVVCPUGetSysReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_sys_reg", "11.0", _hVVCPUGetSysRegErr)
	}
	return _hVVCPUGetSysReg(vcpu, reg, value), nil
}

// HVVCPUGetSysReg gets the current value of a vCPU system register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_sys_reg(_:_:_:)
func HVVCPUGetSysReg(vcpu HVVCPU, reg HVSysReg, value *uint64) HVReturn {
	result, callErr := tryHVVCPUGetSysReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetTrapDebugExceptions func(vcpu HVVCPU, value *bool) HVReturn
var _hVVCPUGetTrapDebugExceptionsErr error

func tryHVVCPUGetTrapDebugExceptions(vcpu HVVCPU, value *bool) (HVReturn, error) {
	if _hVVCPUGetTrapDebugExceptions == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_trap_debug_exceptions", "11.0", _hVVCPUGetTrapDebugExceptionsErr)
	}
	return _hVVCPUGetTrapDebugExceptions(vcpu, value), nil
}

// HVVCPUGetTrapDebugExceptions gets whether debug exceptions exit the guest.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_trap_debug_exceptions(_:_:)
func HVVCPUGetTrapDebugExceptions(vcpu HVVCPU, value *bool) HVReturn {
	result, callErr := tryHVVCPUGetTrapDebugExceptions(vcpu, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetTrapDebugRegAccesses func(vcpu HVVCPU, value *bool) HVReturn
var _hVVCPUGetTrapDebugRegAccessesErr error

func tryHVVCPUGetTrapDebugRegAccesses(vcpu HVVCPU, value *bool) (HVReturn, error) {
	if _hVVCPUGetTrapDebugRegAccesses == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_trap_debug_reg_accesses", "11.0", _hVVCPUGetTrapDebugRegAccessesErr)
	}
	return _hVVCPUGetTrapDebugRegAccesses(vcpu, value), nil
}

// HVVCPUGetTrapDebugRegAccesses gets whether debug-register accesses exit the guest.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_trap_debug_reg_accesses(_:_:)
func HVVCPUGetTrapDebugRegAccesses(vcpu HVVCPU, value *bool) HVReturn {
	result, callErr := tryHVVCPUGetTrapDebugRegAccesses(vcpu, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetVtimerMask func(vcpu HVVCPU, vtimer_is_masked *bool) HVReturn
var _hVVCPUGetVtimerMaskErr error

func tryHVVCPUGetVtimerMask(vcpu HVVCPU, vtimer_is_masked *bool) (HVReturn, error) {
	if _hVVCPUGetVtimerMask == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_vtimer_mask", "11.0", _hVVCPUGetVtimerMaskErr)
	}
	return _hVVCPUGetVtimerMask(vcpu, vtimer_is_masked), nil
}

// HVVCPUGetVtimerMask gets the virtual timer mask.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_vtimer_mask(_:_:)
func HVVCPUGetVtimerMask(vcpu HVVCPU, vtimer_is_masked *bool) HVReturn {
	result, callErr := tryHVVCPUGetVtimerMask(vcpu, vtimer_is_masked)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUGetVtimerOffset func(vcpu HVVCPU, vtimer_offset *uint64) HVReturn
var _hVVCPUGetVtimerOffsetErr error

func tryHVVCPUGetVtimerOffset(vcpu HVVCPU, vtimer_offset *uint64) (HVReturn, error) {
	if _hVVCPUGetVtimerOffset == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_get_vtimer_offset", "11.0", _hVVCPUGetVtimerOffsetErr)
	}
	return _hVVCPUGetVtimerOffset(vcpu, vtimer_offset), nil
}

// HVVCPUGetVtimerOffset returns the vTimer offset for the vCPU ID you specify.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_get_vtimer_offset(_:_:)
func HVVCPUGetVtimerOffset(vcpu HVVCPU, vtimer_offset *uint64) HVReturn {
	result, callErr := tryHVVCPUGetVtimerOffset(vcpu, vtimer_offset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPURun func(vcpu HVVCPU) HVReturn
var _hVVCPURunErr error

func tryHVVCPURun(vcpu HVVCPU) (HVReturn, error) {
	if _hVVCPURun == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_run", "11.0", _hVVCPURunErr)
	}
	return _hVVCPURun(vcpu), nil
}

// HVVCPURun starts the execution of a vCPU.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_run(_:)
func HVVCPURun(vcpu HVVCPU) HVReturn {
	result, callErr := tryHVVCPURun(vcpu)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetPendingInterrupt func(vcpu HVVCPU, type_ HVInterruptType, pending bool) HVReturn
var _hVVCPUSetPendingInterruptErr error

func tryHVVCPUSetPendingInterrupt(vcpu HVVCPU, type_ HVInterruptType, pending bool) (HVReturn, error) {
	if _hVVCPUSetPendingInterrupt == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_pending_interrupt", "11.0", _hVVCPUSetPendingInterruptErr)
	}
	return _hVVCPUSetPendingInterrupt(vcpu, type_, pending), nil
}

// HVVCPUSetPendingInterrupt sets pending interrupts for a vCPU.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_pending_interrupt(_:_:_:)
func HVVCPUSetPendingInterrupt(vcpu HVVCPU, type_ HVInterruptType, pending bool) HVReturn {
	result, callErr := tryHVVCPUSetPendingInterrupt(vcpu, type_, pending)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetReg func(vcpu HVVCPU, reg HVReg, value uint64) HVReturn
var _hVVCPUSetRegErr error

func tryHVVCPUSetReg(vcpu HVVCPU, reg HVReg, value uint64) (HVReturn, error) {
	if _hVVCPUSetReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_reg", "11.0", _hVVCPUSetRegErr)
	}
	return _hVVCPUSetReg(vcpu, reg, value), nil
}

// HVVCPUSetReg sets the value of a vCPU register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_reg(_:_:_:)
func HVVCPUSetReg(vcpu HVVCPU, reg HVReg, value uint64) HVReturn {
	result, callErr := tryHVVCPUSetReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetSIMDFPReg func(vcpu HVVCPU, reg HVSIMDFPReg, value *[16]byte) HVReturn
var _hVVCPUSetSIMDFPRegErr error

func tryHVVCPUSetSIMDFPReg(vcpu HVVCPU, reg HVSIMDFPReg, value [16]byte) (HVReturn, error) {
	if _hVVCPUSetSIMDFPReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_simd_fp_reg", "11.0", _hVVCPUSetSIMDFPRegErr)
	}
	return _hVVCPUSetSIMDFPReg(vcpu, reg, &value), nil
}

// HVVCPUSetSIMDFPReg sets the value of a vCPU SIMD&FP register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_simd_fp_reg(_:_:_:)
func HVVCPUSetSIMDFPReg(vcpu HVVCPU, reg HVSIMDFPReg, value [16]byte) HVReturn {
	result, callErr := tryHVVCPUSetSIMDFPReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetSMEPReg func(vcpu HVVCPU, reg HVSMEPReg, value *byte, length uintptr) HVReturn
var _hVVCPUSetSMEPRegErr error

func tryHVVCPUSetSMEPReg(vcpu HVVCPU, reg HVSMEPReg, value []byte, length uintptr) (HVReturn, error) {
	if _hVVCPUSetSMEPReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_sme_p_reg", "15.2", _hVVCPUSetSMEPRegErr)
	}
	return _hVVCPUSetSMEPReg(vcpu, reg, unsafe.SliceData(value), length), nil
}

// HVVCPUSetSMEPReg sets the value of a vCPU P predicate register in streaming Scalable Vector Extension (SVE) mode.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_p_reg(_:_:_:_:)
func HVVCPUSetSMEPReg(vcpu HVVCPU, reg HVSMEPReg, value []byte, length uintptr) HVReturn {
	result, callErr := tryHVVCPUSetSMEPReg(vcpu, reg, value, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetSMEState func(vcpu HVVCPU, sme_state *HVVCPUSMEState) HVReturn
var _hVVCPUSetSMEStateErr error

func tryHVVCPUSetSMEState(vcpu HVVCPU, sme_state *HVVCPUSMEState) (HVReturn, error) {
	if _hVVCPUSetSMEState == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_sme_state", "15.2", _hVVCPUSetSMEStateErr)
	}
	return _hVVCPUSetSMEState(vcpu, sme_state), nil
}

// HVVCPUSetSMEState sets the SME state consisting of the streaming Scalable Vector Extension (SVE) mode and ZA storage enable.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_state(_:_:)
func HVVCPUSetSMEState(vcpu HVVCPU, sme_state *HVVCPUSMEState) HVReturn {
	result, callErr := tryHVVCPUSetSMEState(vcpu, sme_state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetSMEZReg func(vcpu HVVCPU, reg HVSMEZReg, value *byte, length uintptr) HVReturn
var _hVVCPUSetSMEZRegErr error

func tryHVVCPUSetSMEZReg(vcpu HVVCPU, reg HVSMEZReg, value []byte, length uintptr) (HVReturn, error) {
	if _hVVCPUSetSMEZReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_sme_z_reg", "15.2", _hVVCPUSetSMEZRegErr)
	}
	return _hVVCPUSetSMEZReg(vcpu, reg, unsafe.SliceData(value), length), nil
}

// HVVCPUSetSMEZReg sets the value of a vCPU Z vector register in streaming Scalable Vector Extension (SVE) mode.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_z_reg(_:_:_:_:)
func HVVCPUSetSMEZReg(vcpu HVVCPU, reg HVSMEZReg, value []byte, length uintptr) HVReturn {
	result, callErr := tryHVVCPUSetSMEZReg(vcpu, reg, value, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetSMEZaReg func(vcpu HVVCPU, value *byte, length uintptr) HVReturn
var _hVVCPUSetSMEZaRegErr error

func tryHVVCPUSetSMEZaReg(vcpu HVVCPU, value []byte, length uintptr) (HVReturn, error) {
	if _hVVCPUSetSMEZaReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_sme_za_reg", "15.2", _hVVCPUSetSMEZaRegErr)
	}
	return _hVVCPUSetSMEZaReg(vcpu, unsafe.SliceData(value), length), nil
}

// HVVCPUSetSMEZaReg sets the value of the vCPU ZA matrix register in streaming Scalable Vector Extension (SVE) mode.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_za_reg(_:_:_:)
func HVVCPUSetSMEZaReg(vcpu HVVCPU, value []byte, length uintptr) HVReturn {
	result, callErr := tryHVVCPUSetSMEZaReg(vcpu, value, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetSMEZt0Reg func(vcpu HVVCPU, value *[64]byte) HVReturn
var _hVVCPUSetSMEZt0RegErr error

func tryHVVCPUSetSMEZt0Reg(vcpu HVVCPU, value *[64]byte) (HVReturn, error) {
	if _hVVCPUSetSMEZt0Reg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_sme_zt0_reg", "15.2", _hVVCPUSetSMEZt0RegErr)
	}
	return _hVVCPUSetSMEZt0Reg(vcpu, value), nil
}

// HVVCPUSetSMEZt0Reg sets the value of the vCPU ZT0 register in streaming Scalable Vector Extension (SVE) mode.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sme_zt0_reg(_:_:)
func HVVCPUSetSMEZt0Reg(vcpu HVVCPU, value *[64]byte) HVReturn {
	result, callErr := tryHVVCPUSetSMEZt0Reg(vcpu, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetSysReg func(vcpu HVVCPU, reg HVSysReg, value uint64) HVReturn
var _hVVCPUSetSysRegErr error

func tryHVVCPUSetSysReg(vcpu HVVCPU, reg HVSysReg, value uint64) (HVReturn, error) {
	if _hVVCPUSetSysReg == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_sys_reg", "11.0", _hVVCPUSetSysRegErr)
	}
	return _hVVCPUSetSysReg(vcpu, reg, value), nil
}

// HVVCPUSetSysReg sets the value of a vCPU system register.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_sys_reg(_:_:_:)
func HVVCPUSetSysReg(vcpu HVVCPU, reg HVSysReg, value uint64) HVReturn {
	result, callErr := tryHVVCPUSetSysReg(vcpu, reg, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetTrapDebugExceptions func(vcpu HVVCPU, value bool) HVReturn
var _hVVCPUSetTrapDebugExceptionsErr error

func tryHVVCPUSetTrapDebugExceptions(vcpu HVVCPU, value bool) (HVReturn, error) {
	if _hVVCPUSetTrapDebugExceptions == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_trap_debug_exceptions", "11.0", _hVVCPUSetTrapDebugExceptionsErr)
	}
	return _hVVCPUSetTrapDebugExceptions(vcpu, value), nil
}

// HVVCPUSetTrapDebugExceptions sets whether debug exceptions exit the guest.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_trap_debug_exceptions(_:_:)
func HVVCPUSetTrapDebugExceptions(vcpu HVVCPU, value bool) HVReturn {
	result, callErr := tryHVVCPUSetTrapDebugExceptions(vcpu, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetTrapDebugRegAccesses func(vcpu HVVCPU, value bool) HVReturn
var _hVVCPUSetTrapDebugRegAccessesErr error

func tryHVVCPUSetTrapDebugRegAccesses(vcpu HVVCPU, value bool) (HVReturn, error) {
	if _hVVCPUSetTrapDebugRegAccesses == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_trap_debug_reg_accesses", "11.0", _hVVCPUSetTrapDebugRegAccessesErr)
	}
	return _hVVCPUSetTrapDebugRegAccesses(vcpu, value), nil
}

// HVVCPUSetTrapDebugRegAccesses sets whether debug-register accesses exit the guest.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_trap_debug_reg_accesses(_:_:)
func HVVCPUSetTrapDebugRegAccesses(vcpu HVVCPU, value bool) HVReturn {
	result, callErr := tryHVVCPUSetTrapDebugRegAccesses(vcpu, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetVtimerMask func(vcpu HVVCPU, vtimer_is_masked bool) HVReturn
var _hVVCPUSetVtimerMaskErr error

func tryHVVCPUSetVtimerMask(vcpu HVVCPU, vtimer_is_masked bool) (HVReturn, error) {
	if _hVVCPUSetVtimerMask == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_vtimer_mask", "11.0", _hVVCPUSetVtimerMaskErr)
	}
	return _hVVCPUSetVtimerMask(vcpu, vtimer_is_masked), nil
}

// HVVCPUSetVtimerMask sets or clears the virtual timer mask.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_vtimer_mask(_:_:)
func HVVCPUSetVtimerMask(vcpu HVVCPU, vtimer_is_masked bool) HVReturn {
	result, callErr := tryHVVCPUSetVtimerMask(vcpu, vtimer_is_masked)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVCPUSetVtimerOffset func(vcpu HVVCPU, vtimer_offset uint64) HVReturn
var _hVVCPUSetVtimerOffsetErr error

func tryHVVCPUSetVtimerOffset(vcpu HVVCPU, vtimer_offset uint64) (HVReturn, error) {
	if _hVVCPUSetVtimerOffset == nil {
		return *new(HVReturn), symbolCallError("hv_vcpu_set_vtimer_offset", "11.0", _hVVCPUSetVtimerOffsetErr)
	}
	return _hVVCPUSetVtimerOffset(vcpu, vtimer_offset), nil
}

// HVVCPUSetVtimerOffset sets the vTimer offset to a value that you provide.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpu_set_vtimer_offset(_:_:)
func HVVCPUSetVtimerOffset(vcpu HVVCPU, vtimer_offset uint64) HVReturn {
	result, callErr := tryHVVCPUSetVtimerOffset(vcpu, vtimer_offset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVcpusExit func(vcpus *HVVCPU, vcpu_count uint32) HVReturn
var _hVVcpusExitErr error

func tryHVVcpusExit(vcpus *HVVCPU, vcpu_count uint32) (HVReturn, error) {
	if _hVVcpusExit == nil {
		return *new(HVReturn), symbolCallError("hv_vcpus_exit", "11.0", _hVVcpusExitErr)
	}
	return _hVVcpusExit(vcpus, vcpu_count), nil
}

// HVVcpusExit forces an immediate exit of a set of vCPUs of the VM.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vcpus_exit(_:_:)
func HVVcpusExit(vcpus *HVVCPU, vcpu_count uint32) HVReturn {
	result, callErr := tryHVVcpusExit(vcpus, vcpu_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmAllocate func(uvap unsafe.Pointer, size uintptr, flags HVAllocateFlags) HVReturn
var _hVVmAllocateErr error

func tryHVVmAllocate(uvap unsafe.Pointer, size uintptr, flags HVAllocateFlags) (HVReturn, error) {
	if _hVVmAllocate == nil {
		return *new(HVReturn), symbolCallError("hv_vm_allocate", "12.1", _hVVmAllocateErr)
	}
	return _hVVmAllocate(uvap, size, flags), nil
}

// HVVmAllocate.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_allocate(_:_:_:)
func HVVmAllocate(uvap unsafe.Pointer, size uintptr, flags HVAllocateFlags) HVReturn {
	result, callErr := tryHVVmAllocate(uvap, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigCreate func() unsafe.Pointer
var _hVVmConfigCreateErr error

func tryHVVmConfigCreate() (unsafe.Pointer, error) {
	if _hVVmConfigCreate == nil {
		return nil, symbolCallError("hv_vm_config_create", "13.0", _hVVmConfigCreateErr)
	}
	return _hVVmConfigCreate(), nil
}

// HVVmConfigCreate creates a virtual machine configuration object.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_create()
func HVVmConfigCreate() unsafe.Pointer {
	result, callErr := tryHVVmConfigCreate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigGetDefaultIPAGranule func(granule *HVIPAGranule) HVReturn
var _hVVmConfigGetDefaultIPAGranuleErr error

func tryHVVmConfigGetDefaultIPAGranule(granule *HVIPAGranule) (HVReturn, error) {
	if _hVVmConfigGetDefaultIPAGranule == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_get_default_ipa_granule", "26.0", _hVVmConfigGetDefaultIPAGranuleErr)
	}
	return _hVVmConfigGetDefaultIPAGranule(granule), nil
}

// HVVmConfigGetDefaultIPAGranule.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_default_ipa_granule(_:)
func HVVmConfigGetDefaultIPAGranule(granule *HVIPAGranule) HVReturn {
	result, callErr := tryHVVmConfigGetDefaultIPAGranule(granule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigGetDefaultIPASize func(ipa_bit_length *uint32) HVReturn
var _hVVmConfigGetDefaultIPASizeErr error

func tryHVVmConfigGetDefaultIPASize(ipa_bit_length *uint32) (HVReturn, error) {
	if _hVVmConfigGetDefaultIPASize == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_get_default_ipa_size", "13.0", _hVVmConfigGetDefaultIPASizeErr)
	}
	return _hVVmConfigGetDefaultIPASize(ipa_bit_length), nil
}

// HVVmConfigGetDefaultIPASize.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_default_ipa_size(_:)
func HVVmConfigGetDefaultIPASize(ipa_bit_length *uint32) HVReturn {
	result, callErr := tryHVVmConfigGetDefaultIPASize(ipa_bit_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigGetEl2Enabled func(config unsafe.Pointer, el2_enabled *bool) HVReturn
var _hVVmConfigGetEl2EnabledErr error

func tryHVVmConfigGetEl2Enabled(config unsafe.Pointer, el2_enabled *bool) (HVReturn, error) {
	if _hVVmConfigGetEl2Enabled == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_get_el2_enabled", "15.0", _hVVmConfigGetEl2EnabledErr)
	}
	return _hVVmConfigGetEl2Enabled(config, el2_enabled), nil
}

// HVVmConfigGetEl2Enabled return a status value that indicates whether the VM configuration enables support for Exception Level 2 (EL2).
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_el2_enabled(_:_:)
func HVVmConfigGetEl2Enabled(config unsafe.Pointer, el2_enabled *bool) HVReturn {
	result, callErr := tryHVVmConfigGetEl2Enabled(config, el2_enabled)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigGetEl2Supported func(el2_supported *bool) HVReturn
var _hVVmConfigGetEl2SupportedErr error

func tryHVVmConfigGetEl2Supported(el2_supported *bool) (HVReturn, error) {
	if _hVVmConfigGetEl2Supported == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_get_el2_supported", "15.0", _hVVmConfigGetEl2SupportedErr)
	}
	return _hVVmConfigGetEl2Supported(el2_supported), nil
}

// HVVmConfigGetEl2Supported returns a status value that indicates whether the current platform supports Exception Level 2 (EL2).
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_el2_supported(_:)
func HVVmConfigGetEl2Supported(el2_supported *bool) HVReturn {
	result, callErr := tryHVVmConfigGetEl2Supported(el2_supported)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigGetIPAGranule func(config unsafe.Pointer, granule *HVIPAGranule) HVReturn
var _hVVmConfigGetIPAGranuleErr error

func tryHVVmConfigGetIPAGranule(config unsafe.Pointer, granule *HVIPAGranule) (HVReturn, error) {
	if _hVVmConfigGetIPAGranule == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_get_ipa_granule", "26.0", _hVVmConfigGetIPAGranuleErr)
	}
	return _hVVmConfigGetIPAGranule(config, granule), nil
}

// HVVmConfigGetIPAGranule.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_ipa_granule(_:_:)
func HVVmConfigGetIPAGranule(config unsafe.Pointer, granule *HVIPAGranule) HVReturn {
	result, callErr := tryHVVmConfigGetIPAGranule(config, granule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigGetIPASize func(config unsafe.Pointer, ipa_bit_length *uint32) HVReturn
var _hVVmConfigGetIPASizeErr error

func tryHVVmConfigGetIPASize(config unsafe.Pointer, ipa_bit_length *uint32) (HVReturn, error) {
	if _hVVmConfigGetIPASize == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_get_ipa_size", "13.0", _hVVmConfigGetIPASizeErr)
	}
	return _hVVmConfigGetIPASize(config, ipa_bit_length), nil
}

// HVVmConfigGetIPASize.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_ipa_size(_:_:)
func HVVmConfigGetIPASize(config unsafe.Pointer, ipa_bit_length *uint32) HVReturn {
	result, callErr := tryHVVmConfigGetIPASize(config, ipa_bit_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigGetMaxIPASize func(ipa_bit_length *uint32) HVReturn
var _hVVmConfigGetMaxIPASizeErr error

func tryHVVmConfigGetMaxIPASize(ipa_bit_length *uint32) (HVReturn, error) {
	if _hVVmConfigGetMaxIPASize == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_get_max_ipa_size", "13.0", _hVVmConfigGetMaxIPASizeErr)
	}
	return _hVVmConfigGetMaxIPASize(ipa_bit_length), nil
}

// HVVmConfigGetMaxIPASize.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_get_max_ipa_size(_:)
func HVVmConfigGetMaxIPASize(ipa_bit_length *uint32) HVReturn {
	result, callErr := tryHVVmConfigGetMaxIPASize(ipa_bit_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigSetEl2Enabled func(config unsafe.Pointer, el2_enabled bool) HVReturn
var _hVVmConfigSetEl2EnabledErr error

func tryHVVmConfigSetEl2Enabled(config unsafe.Pointer, el2_enabled bool) (HVReturn, error) {
	if _hVVmConfigSetEl2Enabled == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_set_el2_enabled", "15.0", _hVVmConfigSetEl2EnabledErr)
	}
	return _hVVmConfigSetEl2Enabled(config, el2_enabled), nil
}

// HVVmConfigSetEl2Enabled sets whether the specified VM configuration enables support for Exception Level 2 (EL2).
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_set_el2_enabled(_:_:)
func HVVmConfigSetEl2Enabled(config unsafe.Pointer, el2_enabled bool) HVReturn {
	result, callErr := tryHVVmConfigSetEl2Enabled(config, el2_enabled)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigSetIPAGranule func(config unsafe.Pointer, granule HVIPAGranule) HVReturn
var _hVVmConfigSetIPAGranuleErr error

func tryHVVmConfigSetIPAGranule(config unsafe.Pointer, granule HVIPAGranule) (HVReturn, error) {
	if _hVVmConfigSetIPAGranule == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_set_ipa_granule", "26.0", _hVVmConfigSetIPAGranuleErr)
	}
	return _hVVmConfigSetIPAGranule(config, granule), nil
}

// HVVmConfigSetIPAGranule.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_set_ipa_granule(_:_:)
func HVVmConfigSetIPAGranule(config unsafe.Pointer, granule HVIPAGranule) HVReturn {
	result, callErr := tryHVVmConfigSetIPAGranule(config, granule)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmConfigSetIPASize func(config unsafe.Pointer, ipa_bit_length uint32) HVReturn
var _hVVmConfigSetIPASizeErr error

func tryHVVmConfigSetIPASize(config unsafe.Pointer, ipa_bit_length uint32) (HVReturn, error) {
	if _hVVmConfigSetIPASize == nil {
		return *new(HVReturn), symbolCallError("hv_vm_config_set_ipa_size", "13.0", _hVVmConfigSetIPASizeErr)
	}
	return _hVVmConfigSetIPASize(config, ipa_bit_length), nil
}

// HVVmConfigSetIPASize.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_config_set_ipa_size(_:_:)
func HVVmConfigSetIPASize(config unsafe.Pointer, ipa_bit_length uint32) HVReturn {
	result, callErr := tryHVVmConfigSetIPASize(config, ipa_bit_length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmCreate func(config unsafe.Pointer) HVReturn
var _hVVmCreateErr error

func tryHVVmCreate(config unsafe.Pointer) (HVReturn, error) {
	if _hVVmCreate == nil {
		return *new(HVReturn), symbolCallError("hv_vm_create", "11.0", _hVVmCreateErr)
	}
	return _hVVmCreate(config), nil
}

// HVVmCreate creates a VM instance for the current process.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_create(_:)
func HVVmCreate(config unsafe.Pointer) HVReturn {
	result, callErr := tryHVVmCreate(config)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmDeallocate func(uva unsafe.Pointer, size uintptr) HVReturn
var _hVVmDeallocateErr error

func tryHVVmDeallocate(uva unsafe.Pointer, size uintptr) (HVReturn, error) {
	if _hVVmDeallocate == nil {
		return *new(HVReturn), symbolCallError("hv_vm_deallocate", "12.1", _hVVmDeallocateErr)
	}
	return _hVVmDeallocate(uva, size), nil
}

// HVVmDeallocate.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_deallocate(_:_:)
func HVVmDeallocate(uva unsafe.Pointer, size uintptr) HVReturn {
	result, callErr := tryHVVmDeallocate(uva, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmDestroy func() HVReturn
var _hVVmDestroyErr error

func tryHVVmDestroy() (HVReturn, error) {
	if _hVVmDestroy == nil {
		return *new(HVReturn), symbolCallError("hv_vm_destroy", "11.0", _hVVmDestroyErr)
	}
	return _hVVmDestroy(), nil
}

// HVVmDestroy destroys the VM instance associated with the current process.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_destroy()
func HVVmDestroy() HVReturn {
	result, callErr := tryHVVmDestroy()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmGetMaxVCPUCount func(max_vcpu_count *uint32) HVReturn
var _hVVmGetMaxVCPUCountErr error

func tryHVVmGetMaxVCPUCount(max_vcpu_count *uint32) (HVReturn, error) {
	if _hVVmGetMaxVCPUCount == nil {
		return *new(HVReturn), symbolCallError("hv_vm_get_max_vcpu_count", "11.0", _hVVmGetMaxVCPUCountErr)
	}
	return _hVVmGetMaxVCPUCount(max_vcpu_count), nil
}

// HVVmGetMaxVCPUCount returns the maximum number of vCPUs that the hypervisor supports.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_get_max_vcpu_count(_:)
func HVVmGetMaxVCPUCount(max_vcpu_count *uint32) HVReturn {
	result, callErr := tryHVVmGetMaxVCPUCount(max_vcpu_count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmMap func(addr unsafe.Pointer, ipa HVIPA, size uintptr, flags HVMemoryFlags) HVReturn
var _hVVmMapErr error

func tryHVVmMap(addr unsafe.Pointer, ipa HVIPA, size uintptr, flags HVMemoryFlags) (HVReturn, error) {
	if _hVVmMap == nil {
		return *new(HVReturn), symbolCallError("hv_vm_map", "11.0", _hVVmMapErr)
	}
	return _hVVmMap(addr, ipa, size, flags), nil
}

// HVVmMap maps a region in the virtual address space of the current process into the guest physical address space of the VM.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_map(_:_:_:_:)
func HVVmMap(addr unsafe.Pointer, ipa HVIPA, size uintptr, flags HVMemoryFlags) HVReturn {
	result, callErr := tryHVVmMap(addr, ipa, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmProtect func(ipa HVIPA, size uintptr, flags HVMemoryFlags) HVReturn
var _hVVmProtectErr error

func tryHVVmProtect(ipa HVIPA, size uintptr, flags HVMemoryFlags) (HVReturn, error) {
	if _hVVmProtect == nil {
		return *new(HVReturn), symbolCallError("hv_vm_protect", "11.0", _hVVmProtectErr)
	}
	return _hVVmProtect(ipa, size, flags), nil
}

// HVVmProtect modifies the permissions of a region in the guest physical address space of the VM.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_protect(_:_:_:)
func HVVmProtect(ipa HVIPA, size uintptr, flags HVMemoryFlags) HVReturn {
	result, callErr := tryHVVmProtect(ipa, size, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hVVmUnmap func(ipa HVIPA, size uintptr) HVReturn
var _hVVmUnmapErr error

func tryHVVmUnmap(ipa HVIPA, size uintptr) (HVReturn, error) {
	if _hVVmUnmap == nil {
		return *new(HVReturn), symbolCallError("hv_vm_unmap", "11.0", _hVVmUnmapErr)
	}
	return _hVVmUnmap(ipa, size), nil
}

// HVVmUnmap unmaps a region in the guest physical address space of the VM.
//
// See: https://developer.apple.com/documentation/Hypervisor/hv_vm_unmap(_:_:)
func HVVmUnmap(ipa HVIPA, size uintptr) HVReturn {
	result, callErr := tryHVVmUnmap(ipa, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_hVGICConfigCreate, &_hVGICConfigCreateErr, frameworkHandle, "hv_gic_config_create", "15.0")
	registerFunc(&_hVGICConfigSetDistributorBase, &_hVGICConfigSetDistributorBaseErr, frameworkHandle, "hv_gic_config_set_distributor_base", "15.0")
	registerFunc(&_hVGICConfigSetMsiInterruptRange, &_hVGICConfigSetMsiInterruptRangeErr, frameworkHandle, "hv_gic_config_set_msi_interrupt_range", "15.0")
	registerFunc(&_hVGICConfigSetMsiRegionBase, &_hVGICConfigSetMsiRegionBaseErr, frameworkHandle, "hv_gic_config_set_msi_region_base", "15.0")
	registerFunc(&_hVGICConfigSetRedistributorBase, &_hVGICConfigSetRedistributorBaseErr, frameworkHandle, "hv_gic_config_set_redistributor_base", "15.0")
	registerFunc(&_hVGICCreate, &_hVGICCreateErr, frameworkHandle, "hv_gic_create", "15.0")
	registerFunc(&_hVGICGetDistributorBaseAlignment, &_hVGICGetDistributorBaseAlignmentErr, frameworkHandle, "hv_gic_get_distributor_base_alignment", "15.0")
	registerFunc(&_hVGICGetDistributorReg, &_hVGICGetDistributorRegErr, frameworkHandle, "hv_gic_get_distributor_reg", "15.0")
	registerFunc(&_hVGICGetDistributorSize, &_hVGICGetDistributorSizeErr, frameworkHandle, "hv_gic_get_distributor_size", "15.0")
	registerFunc(&_hVGICGetIccReg, &_hVGICGetIccRegErr, frameworkHandle, "hv_gic_get_icc_reg", "15.0")
	registerFunc(&_hVGICGetIchReg, &_hVGICGetIchRegErr, frameworkHandle, "hv_gic_get_ich_reg", "15.0")
	registerFunc(&_hVGICGetIcvReg, &_hVGICGetIcvRegErr, frameworkHandle, "hv_gic_get_icv_reg", "15.0")
	registerFunc(&_hVGICGetIntid, &_hVGICGetIntidErr, frameworkHandle, "hv_gic_get_intid", "15.0")
	registerFunc(&_hVGICGetMsiReg, &_hVGICGetMsiRegErr, frameworkHandle, "hv_gic_get_msi_reg", "15.0")
	registerFunc(&_hVGICGetMsiRegionBaseAlignment, &_hVGICGetMsiRegionBaseAlignmentErr, frameworkHandle, "hv_gic_get_msi_region_base_alignment", "15.0")
	registerFunc(&_hVGICGetMsiRegionSize, &_hVGICGetMsiRegionSizeErr, frameworkHandle, "hv_gic_get_msi_region_size", "15.0")
	registerFunc(&_hVGICGetRedistributorBase, &_hVGICGetRedistributorBaseErr, frameworkHandle, "hv_gic_get_redistributor_base", "15.0")
	registerFunc(&_hVGICGetRedistributorBaseAlignment, &_hVGICGetRedistributorBaseAlignmentErr, frameworkHandle, "hv_gic_get_redistributor_base_alignment", "15.0")
	registerFunc(&_hVGICGetRedistributorReg, &_hVGICGetRedistributorRegErr, frameworkHandle, "hv_gic_get_redistributor_reg", "15.0")
	registerFunc(&_hVGICGetRedistributorRegionSize, &_hVGICGetRedistributorRegionSizeErr, frameworkHandle, "hv_gic_get_redistributor_region_size", "15.0")
	registerFunc(&_hVGICGetRedistributorSize, &_hVGICGetRedistributorSizeErr, frameworkHandle, "hv_gic_get_redistributor_size", "15.0")
	registerFunc(&_hVGICGetSpiInterruptRange, &_hVGICGetSpiInterruptRangeErr, frameworkHandle, "hv_gic_get_spi_interrupt_range", "15.0")
	registerFunc(&_hVGICReset, &_hVGICResetErr, frameworkHandle, "hv_gic_reset", "15.0")
	registerFunc(&_hVGICSendMsi, &_hVGICSendMsiErr, frameworkHandle, "hv_gic_send_msi", "15.0")
	registerFunc(&_hVGICSetDistributorReg, &_hVGICSetDistributorRegErr, frameworkHandle, "hv_gic_set_distributor_reg", "15.0")
	registerFunc(&_hVGICSetIccReg, &_hVGICSetIccRegErr, frameworkHandle, "hv_gic_set_icc_reg", "15.0")
	registerFunc(&_hVGICSetIchReg, &_hVGICSetIchRegErr, frameworkHandle, "hv_gic_set_ich_reg", "15.0")
	registerFunc(&_hVGICSetIcvReg, &_hVGICSetIcvRegErr, frameworkHandle, "hv_gic_set_icv_reg", "15.0")
	registerFunc(&_hVGICSetMsiReg, &_hVGICSetMsiRegErr, frameworkHandle, "hv_gic_set_msi_reg", "15.0")
	registerFunc(&_hVGICSetRedistributorReg, &_hVGICSetRedistributorRegErr, frameworkHandle, "hv_gic_set_redistributor_reg", "15.0")
	registerFunc(&_hVGICSetSpi, &_hVGICSetSpiErr, frameworkHandle, "hv_gic_set_spi", "15.0")
	registerFunc(&_hVGICSetState, &_hVGICSetStateErr, frameworkHandle, "hv_gic_set_state", "15.0")
	registerFunc(&_hVGICStateCreate, &_hVGICStateCreateErr, frameworkHandle, "hv_gic_state_create", "15.0")
	registerFunc(&_hVGICStateGetData, &_hVGICStateGetDataErr, frameworkHandle, "hv_gic_state_get_data", "15.0")
	registerFunc(&_hVGICStateGetSize, &_hVGICStateGetSizeErr, frameworkHandle, "hv_gic_state_get_size", "15.0")
	registerFunc(&_hVSMEConfigGetMaxSvlBytes, &_hVSMEConfigGetMaxSvlBytesErr, frameworkHandle, "hv_sme_config_get_max_svl_bytes", "15.2")
	registerFunc(&_hVVCPUConfigCreate, &_hVVCPUConfigCreateErr, frameworkHandle, "hv_vcpu_config_create", "11.0")
	registerFunc(&_hVVCPUConfigGetCcsidrEl1SysRegValues, &_hVVCPUConfigGetCcsidrEl1SysRegValuesErr, frameworkHandle, "hv_vcpu_config_get_ccsidr_el1_sys_reg_values", "11.0")
	registerFunc(&_hVVCPUConfigGetFeatureReg, &_hVVCPUConfigGetFeatureRegErr, frameworkHandle, "hv_vcpu_config_get_feature_reg", "11.0")
	registerFunc(&_hVVCPUCreate, &_hVVCPUCreateErr, frameworkHandle, "hv_vcpu_create", "11.0")
	registerFunc(&_hVVCPUDestroy, &_hVVCPUDestroyErr, frameworkHandle, "hv_vcpu_destroy", "11.0")
	registerFunc(&_hVVCPUGetExecTime, &_hVVCPUGetExecTimeErr, frameworkHandle, "hv_vcpu_get_exec_time", "11.0")
	registerFunc(&_hVVCPUGetPendingInterrupt, &_hVVCPUGetPendingInterruptErr, frameworkHandle, "hv_vcpu_get_pending_interrupt", "11.0")
	registerFunc(&_hVVCPUGetReg, &_hVVCPUGetRegErr, frameworkHandle, "hv_vcpu_get_reg", "11.0")
	registerFunc(&_hVVCPUGetSIMDFPReg, &_hVVCPUGetSIMDFPRegErr, frameworkHandle, "hv_vcpu_get_simd_fp_reg", "11.0")
	registerFunc(&_hVVCPUGetSMEPReg, &_hVVCPUGetSMEPRegErr, frameworkHandle, "hv_vcpu_get_sme_p_reg", "15.2")
	registerFunc(&_hVVCPUGetSMEState, &_hVVCPUGetSMEStateErr, frameworkHandle, "hv_vcpu_get_sme_state", "15.2")
	registerFunc(&_hVVCPUGetSMEZReg, &_hVVCPUGetSMEZRegErr, frameworkHandle, "hv_vcpu_get_sme_z_reg", "15.2")
	registerFunc(&_hVVCPUGetSMEZaReg, &_hVVCPUGetSMEZaRegErr, frameworkHandle, "hv_vcpu_get_sme_za_reg", "15.2")
	registerFunc(&_hVVCPUGetSMEZt0Reg, &_hVVCPUGetSMEZt0RegErr, frameworkHandle, "hv_vcpu_get_sme_zt0_reg", "15.2")
	registerFunc(&_hVVCPUGetSysReg, &_hVVCPUGetSysRegErr, frameworkHandle, "hv_vcpu_get_sys_reg", "11.0")
	registerFunc(&_hVVCPUGetTrapDebugExceptions, &_hVVCPUGetTrapDebugExceptionsErr, frameworkHandle, "hv_vcpu_get_trap_debug_exceptions", "11.0")
	registerFunc(&_hVVCPUGetTrapDebugRegAccesses, &_hVVCPUGetTrapDebugRegAccessesErr, frameworkHandle, "hv_vcpu_get_trap_debug_reg_accesses", "11.0")
	registerFunc(&_hVVCPUGetVtimerMask, &_hVVCPUGetVtimerMaskErr, frameworkHandle, "hv_vcpu_get_vtimer_mask", "11.0")
	registerFunc(&_hVVCPUGetVtimerOffset, &_hVVCPUGetVtimerOffsetErr, frameworkHandle, "hv_vcpu_get_vtimer_offset", "11.0")
	registerFunc(&_hVVCPURun, &_hVVCPURunErr, frameworkHandle, "hv_vcpu_run", "11.0")
	registerFunc(&_hVVCPUSetPendingInterrupt, &_hVVCPUSetPendingInterruptErr, frameworkHandle, "hv_vcpu_set_pending_interrupt", "11.0")
	registerFunc(&_hVVCPUSetReg, &_hVVCPUSetRegErr, frameworkHandle, "hv_vcpu_set_reg", "11.0")
	registerFunc(&_hVVCPUSetSIMDFPReg, &_hVVCPUSetSIMDFPRegErr, frameworkHandle, "hv_vcpu_set_simd_fp_reg", "11.0")
	registerFunc(&_hVVCPUSetSMEPReg, &_hVVCPUSetSMEPRegErr, frameworkHandle, "hv_vcpu_set_sme_p_reg", "15.2")
	registerFunc(&_hVVCPUSetSMEState, &_hVVCPUSetSMEStateErr, frameworkHandle, "hv_vcpu_set_sme_state", "15.2")
	registerFunc(&_hVVCPUSetSMEZReg, &_hVVCPUSetSMEZRegErr, frameworkHandle, "hv_vcpu_set_sme_z_reg", "15.2")
	registerFunc(&_hVVCPUSetSMEZaReg, &_hVVCPUSetSMEZaRegErr, frameworkHandle, "hv_vcpu_set_sme_za_reg", "15.2")
	registerFunc(&_hVVCPUSetSMEZt0Reg, &_hVVCPUSetSMEZt0RegErr, frameworkHandle, "hv_vcpu_set_sme_zt0_reg", "15.2")
	registerFunc(&_hVVCPUSetSysReg, &_hVVCPUSetSysRegErr, frameworkHandle, "hv_vcpu_set_sys_reg", "11.0")
	registerFunc(&_hVVCPUSetTrapDebugExceptions, &_hVVCPUSetTrapDebugExceptionsErr, frameworkHandle, "hv_vcpu_set_trap_debug_exceptions", "11.0")
	registerFunc(&_hVVCPUSetTrapDebugRegAccesses, &_hVVCPUSetTrapDebugRegAccessesErr, frameworkHandle, "hv_vcpu_set_trap_debug_reg_accesses", "11.0")
	registerFunc(&_hVVCPUSetVtimerMask, &_hVVCPUSetVtimerMaskErr, frameworkHandle, "hv_vcpu_set_vtimer_mask", "11.0")
	registerFunc(&_hVVCPUSetVtimerOffset, &_hVVCPUSetVtimerOffsetErr, frameworkHandle, "hv_vcpu_set_vtimer_offset", "11.0")
	registerFunc(&_hVVcpusExit, &_hVVcpusExitErr, frameworkHandle, "hv_vcpus_exit", "11.0")
	registerFunc(&_hVVmAllocate, &_hVVmAllocateErr, frameworkHandle, "hv_vm_allocate", "12.1")
	registerFunc(&_hVVmConfigCreate, &_hVVmConfigCreateErr, frameworkHandle, "hv_vm_config_create", "13.0")
	registerFunc(&_hVVmConfigGetDefaultIPAGranule, &_hVVmConfigGetDefaultIPAGranuleErr, frameworkHandle, "hv_vm_config_get_default_ipa_granule", "26.0")
	registerFunc(&_hVVmConfigGetDefaultIPASize, &_hVVmConfigGetDefaultIPASizeErr, frameworkHandle, "hv_vm_config_get_default_ipa_size", "13.0")
	registerFunc(&_hVVmConfigGetEl2Enabled, &_hVVmConfigGetEl2EnabledErr, frameworkHandle, "hv_vm_config_get_el2_enabled", "15.0")
	registerFunc(&_hVVmConfigGetEl2Supported, &_hVVmConfigGetEl2SupportedErr, frameworkHandle, "hv_vm_config_get_el2_supported", "15.0")
	registerFunc(&_hVVmConfigGetIPAGranule, &_hVVmConfigGetIPAGranuleErr, frameworkHandle, "hv_vm_config_get_ipa_granule", "26.0")
	registerFunc(&_hVVmConfigGetIPASize, &_hVVmConfigGetIPASizeErr, frameworkHandle, "hv_vm_config_get_ipa_size", "13.0")
	registerFunc(&_hVVmConfigGetMaxIPASize, &_hVVmConfigGetMaxIPASizeErr, frameworkHandle, "hv_vm_config_get_max_ipa_size", "13.0")
	registerFunc(&_hVVmConfigSetEl2Enabled, &_hVVmConfigSetEl2EnabledErr, frameworkHandle, "hv_vm_config_set_el2_enabled", "15.0")
	registerFunc(&_hVVmConfigSetIPAGranule, &_hVVmConfigSetIPAGranuleErr, frameworkHandle, "hv_vm_config_set_ipa_granule", "26.0")
	registerFunc(&_hVVmConfigSetIPASize, &_hVVmConfigSetIPASizeErr, frameworkHandle, "hv_vm_config_set_ipa_size", "13.0")
	registerFunc(&_hVVmCreate, &_hVVmCreateErr, frameworkHandle, "hv_vm_create", "11.0")
	registerFunc(&_hVVmDeallocate, &_hVVmDeallocateErr, frameworkHandle, "hv_vm_deallocate", "12.1")
	registerFunc(&_hVVmDestroy, &_hVVmDestroyErr, frameworkHandle, "hv_vm_destroy", "11.0")
	registerFunc(&_hVVmGetMaxVCPUCount, &_hVVmGetMaxVCPUCountErr, frameworkHandle, "hv_vm_get_max_vcpu_count", "11.0")
	registerFunc(&_hVVmMap, &_hVVmMapErr, frameworkHandle, "hv_vm_map", "11.0")
	registerFunc(&_hVVmProtect, &_hVVmProtectErr, frameworkHandle, "hv_vm_protect", "11.0")
	registerFunc(&_hVVmUnmap, &_hVVmUnmapErr, frameworkHandle, "hv_vm_unmap", "11.0")
}
