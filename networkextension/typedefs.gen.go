// Code generated from Apple documentation. DO NOT EDIT.

package networkextension

import (
	"unsafe"
	"github.com/tmc/apple/objectivec"
)

// NEFilterPacketHandler is a type for Swift closures or ObjectiveC blocks that make filtering decisions about network packets.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketHandler
type NEFilterPacketHandler = func(NEFilterPacketContext, objectivec.Object, NETrafficDirection, unsafe.Pointer, uint32) NEFilterPacketProviderVerdict

// See: https://developer.apple.com/documentation/NetworkExtension/NWEndpointArray
type NWEndpointArray = []objectivec.Object

