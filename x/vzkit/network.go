package vzkit

import (
	"fmt"
	"strings"

	"github.com/ebitengine/purego"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/vmnet"
)

// NetworkMode represents the type of network configuration.
type NetworkMode string

const (
	NetworkModeNAT     NetworkMode = "nat"
	NetworkModeBridged NetworkMode = "bridged"
	NetworkModeVMNet   NetworkMode = "vmnet"
	NetworkModeNone    NetworkMode = "none"
)

// NetworkConfig holds network configuration settings.
type NetworkConfig struct {
	Mode      NetworkMode
	Interface string // For bridged mode: the host interface name (e.g., "en0")
}

// ParseNetworkMode parses a network mode string.
//
// Formats:
//   - "nat" - NAT networking (default)
//   - "bridged:en0" - Bridged to specific interface
//   - "vmnet" - VMNet shared networking
//   - "none" - No networking
func ParseNetworkMode(s string) (NetworkConfig, error) {
	if s == "" || s == "nat" {
		return NetworkConfig{Mode: NetworkModeNAT}, nil
	}
	if s == "none" {
		return NetworkConfig{Mode: NetworkModeNone}, nil
	}
	if s == "vmnet" {
		return NetworkConfig{Mode: NetworkModeVMNet}, nil
	}
	if strings.HasPrefix(s, "bridged:") {
		iface := strings.TrimPrefix(s, "bridged:")
		if iface == "" {
			return NetworkConfig{}, fmt.Errorf("bridged mode requires interface name (e.g., bridged:en0)")
		}
		return NetworkConfig{Mode: NetworkModeBridged, Interface: iface}, nil
	}
	return NetworkConfig{}, fmt.Errorf("unknown network mode: %s (use nat, bridged:<iface>, vmnet, or none)", s)
}

// vmnet operating mode constants (not yet in the generated vmnet bindings).
const (
	vmnetSharedMode vmnet.Operating_modes_t = 1001
	vmnetSuccess    vmnet.Vmnet_return_t    = 1000
)

// vmnet_network_configuration_create and vmnet_network_create are macOS 26+
// functions needed to create a vmnet network for VZVmnetNetworkDeviceAttachment.
var (
	_vmnetNetworkConfigurationCreate func(mode vmnet.Vmnet_mode_t, status *vmnet.Vmnet_return_t) vmnet.Vmnet_network_configuration_ref
	_vmnetNetworkCreate              func(config vmnet.Vmnet_network_configuration_ref, status *vmnet.Vmnet_return_t) vmnet.Vmnet_network_ref
	vmnetFuncsLoaded                 bool
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
	purego.RegisterFunc(&_vmnetNetworkConfigurationCreate, sym1)
	purego.RegisterFunc(&_vmnetNetworkCreate, sym2)
	vmnetFuncsLoaded = true
}

// createVMNetNetwork creates a shared-mode vmnet network.
func createVMNetNetwork() (vmnet.Vmnet_network_ref, error) {
	if !vmnetFuncsLoaded {
		return 0, fmt.Errorf("vmnet network APIs unavailable (requires macOS 26+)")
	}
	var status vmnet.Vmnet_return_t
	config := _vmnetNetworkConfigurationCreate(vmnetSharedMode, &status)
	if config == 0 || status != vmnetSuccess {
		return 0, fmt.Errorf("create vmnet network configuration (status %d)", status)
	}
	network := _vmnetNetworkCreate(config, &status)
	if network == 0 || status != vmnetSuccess {
		return 0, fmt.Errorf("create vmnet network (status %d)", status)
	}
	return network, nil
}

// CreateNetworkAttachment creates a network device attachment based on config.
func CreateNetworkAttachment(config NetworkConfig) (vz.VZNetworkDeviceAttachment, error) {
	switch config.Mode {
	case NetworkModeNAT:
		attachment := vz.NewVZNATNetworkDeviceAttachment()
		if attachment.ID == 0 {
			return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("create NAT attachment")
		}
		return vz.VZNetworkDeviceAttachmentFromID(attachment.ID), nil

	case NetworkModeBridged:
		interfaces := vz.GetVZBridgedNetworkInterfaceClass().NetworkInterfaces()
		if len(interfaces) == 0 {
			return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("no bridged network interfaces available")
		}
		var target vz.VZBridgedNetworkInterface
		for _, iface := range interfaces {
			if iface.Identifier() == config.Interface {
				target = iface
				break
			}
		}
		if target.ID == 0 {
			var available []string
			for _, iface := range interfaces {
				available = append(available, iface.Identifier())
			}
			return vz.VZNetworkDeviceAttachment{}, fmt.Errorf(
				"interface '%s' not found; available: %s",
				config.Interface, strings.Join(available, ", "))
		}
		attachment := vz.NewBridgedNetworkDeviceAttachmentWithInterface(&target)
		if attachment.ID == 0 {
			return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("create bridged attachment for %s", config.Interface)
		}
		return attachment.VZNetworkDeviceAttachment, nil

	case NetworkModeVMNet:
		network, err := createVMNetNetwork()
		if err != nil {
			return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("vmnet: %w", err)
		}
		attachment := vz.NewVmnetNetworkDeviceAttachmentWithNetwork(network)
		if attachment.ID == 0 {
			return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("create vmnet attachment")
		}
		return attachment.VZNetworkDeviceAttachment, nil

	case NetworkModeNone:
		return vz.VZNetworkDeviceAttachment{}, nil

	default:
		return vz.VZNetworkDeviceAttachment{}, fmt.Errorf("unsupported network mode: %s", config.Mode)
	}
}

// CreateNetworkDevice creates a VZVirtioNetworkDeviceConfiguration with a random MAC.
func CreateNetworkDevice(config NetworkConfig) (vz.VZVirtioNetworkDeviceConfiguration, error) {
	if config.Mode == NetworkModeNone {
		return vz.VZVirtioNetworkDeviceConfiguration{}, nil
	}

	attachment, err := CreateNetworkAttachment(config)
	if err != nil {
		return vz.VZVirtioNetworkDeviceConfiguration{}, err
	}

	networkConfig := vz.NewVZVirtioNetworkDeviceConfiguration()
	if networkConfig.ID == 0 {
		return networkConfig, fmt.Errorf("create network configuration")
	}
	networkConfig.SetAttachment(&attachment)

	macAddr := vz.GetVZMACAddressClass().RandomLocallyAdministeredAddress()
	if macAddr.ID != 0 {
		networkConfig.SetMACAddress(&macAddr)
	}

	return networkConfig, nil
}
