// Package irecovery discovers Apple DFU and recovery USB endpoints and provides
// control and bulk transfers through libusb. It does not use usbmuxd or own a VM.
// Callers supply a libusb shared-library path and select a device by ECID.
// Transfers are serialized, have a maximum one-second native timeout, and are
// never retried automatically because a partial write may have changed the device.
package irecovery
