// Code generated from Apple documentation. DO NOT EDIT.

package vmnet

import (
	"unsafe"
	"github.com/ebitengine/purego"
)

var vmnet_allocate_mac_address_key int8

var vmnet_enable_checksum_offload_key int8

var vmnet_enable_isolation_key int8

var vmnet_enable_tso_key int8

var vmnet_enable_virtio_header_key int8

var vmnet_end_address_key int8

var vmnet_estimated_packets_available_key int8

var vmnet_host_ip_address_key int8

var vmnet_host_ipv6_address_key int8

var vmnet_host_subnet_mask_key int8

var vmnet_interface_id_key int8

var vmnet_mac_address_key int8

var vmnet_max_packet_size_key int8

var vmnet_mtu_key int8

var vmnet_nat66_prefix_key int8

var vmnet_network_identifier_key int8

var vmnet_operation_mode_key int8

var vmnet_read_max_packets_key int8

var vmnet_shared_interface_name_key int8

var vmnet_start_address_key int8

var vmnet_subnet_mask_key int8

var vmnet_write_max_packets_key int8

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_allocate_mac_address_key"); err == nil && ptr != 0 {
		vmnet_allocate_mac_address_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_enable_checksum_offload_key"); err == nil && ptr != 0 {
		vmnet_enable_checksum_offload_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_enable_isolation_key"); err == nil && ptr != 0 {
		vmnet_enable_isolation_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_enable_tso_key"); err == nil && ptr != 0 {
		vmnet_enable_tso_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_enable_virtio_header_key"); err == nil && ptr != 0 {
		vmnet_enable_virtio_header_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_end_address_key"); err == nil && ptr != 0 {
		vmnet_end_address_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_estimated_packets_available_key"); err == nil && ptr != 0 {
		vmnet_estimated_packets_available_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_host_ip_address_key"); err == nil && ptr != 0 {
		vmnet_host_ip_address_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_host_ipv6_address_key"); err == nil && ptr != 0 {
		vmnet_host_ipv6_address_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_host_subnet_mask_key"); err == nil && ptr != 0 {
		vmnet_host_subnet_mask_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_interface_id_key"); err == nil && ptr != 0 {
		vmnet_interface_id_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_mac_address_key"); err == nil && ptr != 0 {
		vmnet_mac_address_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_max_packet_size_key"); err == nil && ptr != 0 {
		vmnet_max_packet_size_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_mtu_key"); err == nil && ptr != 0 {
		vmnet_mtu_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_nat66_prefix_key"); err == nil && ptr != 0 {
		vmnet_nat66_prefix_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_network_identifier_key"); err == nil && ptr != 0 {
		vmnet_network_identifier_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_operation_mode_key"); err == nil && ptr != 0 {
		vmnet_operation_mode_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_read_max_packets_key"); err == nil && ptr != 0 {
		vmnet_read_max_packets_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_shared_interface_name_key"); err == nil && ptr != 0 {
		vmnet_shared_interface_name_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_start_address_key"); err == nil && ptr != 0 {
		vmnet_start_address_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_subnet_mask_key"); err == nil && ptr != 0 {
		vmnet_subnet_mask_key = *(*int8)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "vmnet_write_max_packets_key"); err == nil && ptr != 0 {
		vmnet_write_max_packets_key = *(*int8)(unsafe.Pointer(ptr))
	}

}

