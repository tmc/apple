package vzkit

import (
	vz "github.com/tmc/apple/virtualization"
	networkx "github.com/tmc/apple/x/vzkit/network"
)

// NetworkMode represents the type of network configuration.
type NetworkMode = networkx.Mode

const (
	NetworkModeNAT      = networkx.ModeNAT
	NetworkModeBridged  = networkx.ModeBridged
	NetworkModeHostOnly = networkx.ModeHostOnly
	NetworkModeVMNet    = networkx.ModeVMNet
	NetworkModeNone     = networkx.ModeNone
)

// NetworkConfig holds network configuration settings.
type NetworkConfig = networkx.Config

// ParseNetworkMode parses a network mode string.
func ParseNetworkMode(s string) (NetworkConfig, error) { return networkx.Parse(s) }

// LookupBridgedNetworkInterface resolves a bridged interface by identifier.
func LookupBridgedNetworkInterface(identifier string) (vz.VZBridgedNetworkInterface, error) {
	return networkx.LookupBridgedNetworkInterface(identifier)
}

// CreateNATAttachment creates a NAT attachment.
func CreateNATAttachment(isolate bool) (vz.VZNetworkDeviceAttachment, error) {
	return networkx.CreateNATAttachment(isolate)
}

// CreateBridgedAttachment creates a bridged attachment.
func CreateBridgedAttachment(identifier string, macNAT bool) (vz.VZNetworkDeviceAttachment, error) {
	return networkx.CreateBridgedAttachment(identifier, macNAT)
}

// CreateHostOnlyNetworkAttachment creates a host-only attachment.
func CreateHostOnlyNetworkAttachment() (vz.VZNetworkDeviceAttachment, error) {
	return networkx.CreateHostOnlyNetworkAttachment()
}

// VhostUserNetworkConfig describes a private vhost-user network attachment.
type VhostUserNetworkConfig = networkx.VhostUserConfig

// DefaultVhostUserNetworkConfig returns a config populated with framework defaults.
func DefaultVhostUserNetworkConfig(interfaceName string) VhostUserNetworkConfig {
	return networkx.DefaultVhostUserConfig(interfaceName)
}

// CreateVhostUserNetworkAttachment creates a private vhost-user attachment.
func CreateVhostUserNetworkAttachment(cfg VhostUserNetworkConfig) (vz.VZNetworkDeviceAttachment, error) {
	return networkx.CreateVhostUserNetworkAttachment(cfg)
}

// CreateNetworkAttachment creates a network device attachment based on config.
func CreateNetworkAttachment(config NetworkConfig) (vz.VZNetworkDeviceAttachment, error) {
	return networkx.CreateAttachment(config)
}

// CreateNetworkDevice creates a VZVirtioNetworkDeviceConfiguration with a random MAC.
func CreateNetworkDevice(config NetworkConfig) (vz.VZVirtioNetworkDeviceConfiguration, error) {
	return networkx.CreateDevice(config)
}
