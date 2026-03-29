// Package customvirtio provides experimental custom Virtio device helpers.
//
// It is intended to wrap private custom Virtio device configuration types,
// including PCI identifiers, queue counts, feature bits, provider hookup, and
// plugin naming for device implementations outside the built-in Virtio set.
//
// The package should expose the minimum surface needed to describe a custom
// device without hiding important low-level configuration choices.
package customvirtio
