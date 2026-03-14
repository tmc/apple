//go:build darwin

package telemetry

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/private/appleneuralengine"
	"github.com/tmc/apple/x/ane"
)

// ProbeClientInfo returns capability information about the ANE connection.
func ProbeClientInfo(c *ane.Client) ClientInfo {
	var ci ClientInfo
	client := appleneuralengine.ANEClientFromID(objc.ID(c.ClientObjcID()))
	vc := client.VirtualClient()
	if vc != nil {
		probeVirtualClientInfo(vc, &ci)
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
func ProbeCacheInfo(c *ane.Client) CacheInfo {
	var ci CacheInfo
	func() {
		defer func() { recover() }()
		mgr := appleneuralengine.NewANEModelCacheManager()
		if mgr.ID == 0 {
			return
		}
		dirURL := mgr.CacheDir()
		if dirURL == nil || dirURL.GetID() == 0 {
			return
		}
		nsURL := foundation.NSURLFromID(dirURL.GetID())
		ci.CacheDir = nsURL.Path()
		ci.CacheDirKnown = true
	}()
	return ci
}

// Snapshot captures a point-in-time snapshot of the host and ANE environment.
func Snapshot(c *ane.Client) ClientSnapshot {
	return ClientSnapshot{
		Device: c.Info(),
		Client: ProbeClientInfo(c),
		Cache:  ProbeCacheInfo(c),
	}
}
