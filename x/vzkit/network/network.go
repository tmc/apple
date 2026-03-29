package network

import (
	"fmt"
	"strings"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	privvz "github.com/tmc/apple/private/virtualization"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/vmnet"
)

// Mode describes a network attachment mode.
type Mode string

const (
	ModeNAT      Mode = "nat"
	ModeBridged  Mode = "bridged"
	ModeHostOnly Mode = "host-only"
	ModeVMNet    Mode = "vmnet"
	ModeNone     Mode = "none"
)

// Config describes a VM network device configuration.
type Config struct {
	Mode      Mode
	Interface string
}

// Parse parses a network mode string.
func Parse(s string) (Config, error) {
	if s == "" || s == "nat" {
		return Config{Mode: ModeNAT}, nil
	}
	if s == "none" {
		return Config{Mode: ModeNone}, nil
	}
	if s == "host-only" {
		return Config{Mode: ModeHostOnly}, nil
	}
	if s == "vmnet" {
		return Config{Mode: ModeVMNet}, nil
	}
	if strings.HasPrefix(s, "bridged:") {
		iface := strings.TrimPrefix(s, "bridged:")
		if iface == "" {
			return Config{}, fmt.Errorf("bridged mode requires interface name (e.g., bridged:en0)")
		}
		return Config{Mode: ModeBridged, Interface: iface}, nil
	}
	return Config{}, fmt.Errorf("unknown network mode: %s (use nat, bridged:<iface>, host-only, vmnet, or none)", s)
}

const (
	vmnetSharedMode vmnet.Operating_modes_t = 1001
	vmnetSuccess    vmnet.Vmnet_return_t    = 1000
)

var (
	vmnetNetworkConfigurationCreate func(mode vmnet.Vmnet_mode_t, status *vmnet.Vmnet_return_t) vmnet.Vmnet_network_configuration_ref
	vmnetNetworkCreate              func(config vmnet.Vmnet_network_configuration_ref, status *vmnet.Vmnet_return_t) vmnet.Vmnet_network_ref
	vmnetFuncsLoaded                bool
)

func init() {
	handle, err := purego.Dlopen("/System/Library/Frameworks/vmnet.framework/vmnet", purego.RTLD_LAZY)
	if err != nil {
		return
	}
	sym1, err := purego.Dlsym(handle, "vmnet_network_configuration_create")
	if err != nil {
		return
	}
	sym2, err := purego.Dlsym(handle, "vmnet_network_create")
	if err != nil {
		return
	}
	purego.RegisterFunc(&vmnetNetworkConfigurationCreate, sym1)
	purego.RegisterFunc(&vmnetNetworkCreate, sym2)
	vmnetFuncsLoaded = true
}

func createVMNetNetwork() (vmnet.Vmnet_network_ref, error) {
	if !vmnetFuncsLoaded {
		return 0, fmt.Errorf("vmnet network APIs unavailable (requires macOS 26+)")
	}
	var status vmnet.Vmnet_return_t
	config := vmnetNetworkConfigurationCreate(vmnetSharedMode, &status)
	if config == 0 || status != vmnetSuccess {
		return 0, fmt.Errorf("create vmnet network configuration (status %d)", status)
	}
	network := vmnetNetworkCreate(config, &status)
	if network == 0 || status != vmnetSuccess {
		return 0, fmt.Errorf("create vmnet network (status %d)", status)
	}
	return network, nil
}

// LookupBridgedNetworkInterface resolves a bridged interface by identifier.
func LookupBridgedNetworkInterface(identifier string) (vz.VZBridgedNetworkInterface, error) {
	interfaces := vz.GetVZBridgedNetworkInterfaceClass().NetworkInterfaces()
	if len(interfaces) == 0 {
		return vz.VZBridgedNetworkInterface{}, fmt.Errorf("no bridged network interfaces available")
	}
	for _, iface := range interfaces {
		if iface.Identifier() == identifier {
			return iface, nil
		}
	}

	var available []string
	for _, iface := range interfaces {
		available = append(available, iface.Identifier())
	}
	return vz.VZBridgedNetworkInterface{}, fmt.Errorf("interface %q not found; available: %s", identifier, strings.Join(available, ", "))
}

// CreateNATAttachment creates a NAT attachment.
func CreateNATAttachment(isolate bool) (vz.VZNetworkDeviceAttachment, error) {
	attachment := vz.NewVZNATNetworkDeviceAttachment()
	if attachment.ID == 0 {
		return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("create NAT attachment")
	}
	if isolate {
		privvz.VZNATNetworkDeviceAttachmentFromID(attachment.ID).SetInterfaceIsolationEnabled(true)
	}
	return vz.VZNetworkDeviceAttachmentFromID(attachment.ID), nil
}

// CreateBridgedAttachment creates a bridged attachment for the named interface.
func CreateBridgedAttachment(identifier string, macNAT bool) (vz.VZNetworkDeviceAttachment, error) {
	target, err := LookupBridgedNetworkInterface(identifier)
	if err != nil {
		return vz.VZNetworkDeviceAttachment{}, err
	}
	attachment := vz.NewBridgedNetworkDeviceAttachmentWithInterface(&target)
	if attachment.ID == 0 {
		return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("create bridged attachment for %s", identifier)
	}
	if macNAT {
		privvz.VZBridgedNetworkDeviceAttachmentFromID(attachment.ID).SetMacNatEnabled(true)
	}
	return attachment.VZNetworkDeviceAttachment, nil
}

// CreateHostOnlyNetworkAttachment creates a host-only attachment.
func CreateHostOnlyNetworkAttachment() (vz.VZNetworkDeviceAttachment, error) {
	attachment := privvz.NewVZHostOnlyNetworkDeviceAttachment()
	if attachment.ID == 0 {
		return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("create host-only attachment")
	}
	return vz.VZNetworkDeviceAttachmentFromID(attachment.ID), nil
}

// VhostUserConfig describes a private vhost-user network attachment.
type VhostUserConfig struct {
	Interface                   string
	XPCService                  int64
	MaximumTransmissionUnit     int64
	HostChecksumOffload         int64
	GuestChecksumOffload        int64
	HostTCPSegmentationOffload  int64
	GuestTCPSegmentationOffload int64
}

// DefaultVhostUserConfig returns a config populated with framework defaults.
func DefaultVhostUserConfig(interfaceName string) VhostUserConfig {
	class := privvz.GetVZVhostUserNetworkDeviceAttachmentClass()
	defaultOffload := class.DefaultOffloadMode()
	return VhostUserConfig{
		Interface:                   interfaceName,
		MaximumTransmissionUnit:     class.DefaultMaximumTransmissionUnit(),
		HostChecksumOffload:         defaultOffload,
		GuestChecksumOffload:        defaultOffload,
		HostTCPSegmentationOffload:  defaultOffload,
		GuestTCPSegmentationOffload: defaultOffload,
	}
}

// CreateVhostUserNetworkAttachment creates a private vhost-user attachment.
func CreateVhostUserNetworkAttachment(cfg VhostUserConfig) (vz.VZNetworkDeviceAttachment, error) {
	if cfg.Interface == "" {
		return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("vhost-user interface is required")
	}

	class := privvz.GetVZVhostUserNetworkDeviceAttachmentClass()
	if cfg.MaximumTransmissionUnit == 0 {
		cfg.MaximumTransmissionUnit = class.DefaultMaximumTransmissionUnit()
	}
	if cfg.HostChecksumOffload == 0 &&
		cfg.GuestChecksumOffload == 0 &&
		cfg.HostTCPSegmentationOffload == 0 &&
		cfg.GuestTCPSegmentationOffload == 0 {
		defaultOffload := class.DefaultOffloadMode()
		cfg.HostChecksumOffload = defaultOffload
		cfg.GuestChecksumOffload = defaultOffload
		cfg.HostTCPSegmentationOffload = defaultOffload
		cfg.GuestTCPSegmentationOffload = defaultOffload
	}

	iface := objectivec.Object{ID: objc.String(cfg.Interface)}
	var (
		attachment privvz.VZVhostUserNetworkDeviceAttachment
		err        error
	)
	if cfg.XPCService != 0 {
		attachment, err = privvz.NewVZVhostUserNetworkDeviceAttachmentWithInterfaceXpcServiceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError(
			iface,
			cfg.XPCService,
			cfg.MaximumTransmissionUnit,
			cfg.HostChecksumOffload,
			cfg.GuestChecksumOffload,
			cfg.HostTCPSegmentationOffload,
			cfg.GuestTCPSegmentationOffload,
		)
	} else {
		attachment, err = privvz.NewVZVhostUserNetworkDeviceAttachmentWithInterfaceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError(
			iface,
			cfg.MaximumTransmissionUnit,
			cfg.HostChecksumOffload,
			cfg.GuestChecksumOffload,
			cfg.HostTCPSegmentationOffload,
			cfg.GuestTCPSegmentationOffload,
		)
	}
	if err != nil {
		return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("create vhost-user attachment: %w", err)
	}
	return vz.VZNetworkDeviceAttachmentFromID(attachment.ID), nil
}

// CreateAttachment creates a network device attachment based on cfg.
func CreateAttachment(cfg Config) (vz.VZNetworkDeviceAttachment, error) {
	switch cfg.Mode {
	case ModeNAT:
		return CreateNATAttachment(false)
	case ModeBridged:
		return CreateBridgedAttachment(cfg.Interface, false)
	case ModeHostOnly:
		return CreateHostOnlyNetworkAttachment()
	case ModeVMNet:
		network, err := createVMNetNetwork()
		if err != nil {
			return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("vmnet: %w", err)
		}
		attachment := vz.NewVmnetNetworkDeviceAttachmentWithNetwork(network)
		if attachment.ID == 0 {
			return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("create vmnet attachment")
		}
		return attachment.VZNetworkDeviceAttachment, nil
	case ModeNone:
		return vz.VZNetworkDeviceAttachment{}, nil
	default:
		return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("unsupported network mode: %s", cfg.Mode)
	}
}

// CreateDevice creates a Virtio network device with a random MAC address.
func CreateDevice(cfg Config) (vz.VZVirtioNetworkDeviceConfiguration, error) {
	if cfg.Mode == ModeNone {
		return vz.VZVirtioNetworkDeviceConfiguration{}, nil
	}
	attachment, err := CreateAttachment(cfg)
	if err != nil {
		return vz.VZVirtioNetworkDeviceConfiguration{}, err
	}
	device := vz.NewVZVirtioNetworkDeviceConfiguration()
	if device.ID == 0 {
		return device, fmt.Errorf("create network configuration")
	}
	device.SetAttachment(&attachment)
	macAddr := vz.GetVZMACAddressClass().RandomLocallyAdministeredAddress()
	if macAddr.ID != 0 {
		device.SetMACAddress(&macAddr)
	}
	return device, nil
}
