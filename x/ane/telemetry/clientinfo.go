//go:build darwin

package telemetry

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/private/appleneuralengine"
	"github.com/tmc/apple/x/ane"
)

// ProbeClientInfo returns capability information about the ANE connection.
func ProbeClientInfo(c *ane.Client) ClientInfo {
	var ci ClientInfo
	client := appleneuralengine.ANEClientFromID(objc.ID(c.ClientObjcID()))
	vc := client.VirtualClient()
	if vc != nil && vc.GetID() != 0 {
		avc := appleneuralengine.ANEVirtualClientFromID(vc.GetID())
		probeVirtualClientInfo(&avc, &ci)
		return ci
	}
	probeDeviceInfoFallback(&ci)
	return ci
}

func probeVirtualClientInfo(vc *appleneuralengine.ANEVirtualClient, ci *ClientInfo) {
	func() {
		defer func() { recover() }()
		ci.NumANEs = vc.NumANEs()
		ci.NumANEsKnown = true
	}()
	func() {
		defer func() { recover() }()
		ci.NumCores = vc.NumANECores()
		ci.NumCoresKnown = true
	}()
	func() {
		defer func() { recover() }()
		if obj := vc.AneArchitectureTypeStr(); obj != nil {
			ci.ArchitectureStr = descriptionString(obj.GetID())
			ci.ArchitectureStrKnown = true
		}
	}()
	func() {
		defer func() { recover() }()
		ci.BoardType = vc.AneBoardtype()
		ci.BoardTypeKnown = true
	}()
	func() {
		defer func() { recover() }()
		ci.CapabilityMask = vc.NegotiatedCapabilityMask()
		ci.CapabilityMaskKnown = true
	}()
	func() {
		defer func() { recover() }()
		ci.DataInterfaceVersion = uint64(vc.NegotiatedDataInterfaceVersion())
		ci.DataInterfaceVersionKnown = true
	}()
	func() {
		defer func() { recover() }()
		ci.PrecompiledBinarySupported = vc.ValidateEnvironmentForPrecompiledBinarySupport()
		ci.PrecompiledBinarySupportedKnown = true
	}()
}

func probeDeviceInfoFallback(ci *ClientInfo) {
	cls := appleneuralengine.GetANEDeviceInfoClass()
	func() {
		defer func() { recover() }()
		ci.NumANEs = cls.NumANEs()
		ci.NumANEsKnown = true
	}()
	func() {
		defer func() { recover() }()
		ci.NumCores = cls.NumANECores()
		ci.NumCoresKnown = true
	}()
	func() {
		defer func() { recover() }()
		if obj := cls.AneArchitectureType(); obj != nil {
			ci.ArchitectureStr = descriptionString(obj.GetID())
			ci.ArchitectureStrKnown = true
		}
	}()
	func() {
		defer func() { recover() }()
		ci.BoardType = cls.AneBoardType()
		ci.BoardTypeKnown = true
	}()
}

// ProbeCacheInfo returns ANE model cache information.
// Note: _ANEModelCacheManager is an XPC-only class that cannot be
// instantiated from client processes, so cache info is unavailable.
func ProbeCacheInfo(c *ane.Client) CacheInfo {
	return CacheInfo{}
}

// Snapshot captures a point-in-time snapshot of the host and ANE environment.
func Snapshot(c *ane.Client) ClientSnapshot {
	return ClientSnapshot{
		Device: c.Info(),
		Client: ProbeClientInfo(c),
		Cache:  ProbeCacheInfo(c),
	}
}
