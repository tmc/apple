// Package storagehotplug provides runtime storage device management.
//
// It is intended to wrap the private storage device APIs used to create live
// devices, swap attachments, resize backing media, and detach or replace disks
// on a running virtual machine.
//
// The package should focus on:
//
//	Constructing live storage devices from attachments
//	Swapping attachments with completion handlers
//	Safe helpers for common hot-plug and hot-swap flows
//
// This package complements the static storage configuration helpers in the
// parent vzkit package.
package storagehotplug
