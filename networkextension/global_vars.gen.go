// Code generated from Apple documentation. DO NOT EDIT.

package networkextension

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// NEAppProxyErrorDomain is the domain used for app proxy errors.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyErrorDomain
	NEAppProxyErrorDomain string
	// NEDNSProxyConfigurationDidChangeNotification is a notification that is posted when the DNS proxy configuration changes.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyConfigurationDidChangeNotification
	NEDNSProxyConfigurationDidChangeNotification string
	// NEDNSProxyErrorDomain is the DNS proxy error domain.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyErrorDomain
	NEDNSProxyErrorDomain string
	// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsConfigurationDidChangeNotification
	NEDNSSettingsConfigurationDidChangeNotification string
	// NEDNSSettingsErrorDomain is the domain for errors resulting from calls to the DNS settings manager.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NEDNSSettingsErrorDomain
	NEDNSSettingsErrorDomain string
	// NEFilterConfigurationDidChangeNotification is posted after the filter configuration stored in the Network Extension preferences changes.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterConfigurationDidChangeNotification
	NEFilterConfigurationDidChangeNotification string
	// NEFilterErrorDomain is the domain for errors resulting from calls to the filter manager.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterErrorDomain
	NEFilterErrorDomain string
	// See: https://developer.apple.com/documentation/NetworkExtension/NERelayClientErrorDomain
	NERelayClientErrorDomain string
	// See: https://developer.apple.com/documentation/NetworkExtension/NERelayConfigurationDidChangeNotification
	NERelayConfigurationDidChangeNotification string
	// NERelayErrorDomain is the domain for errors resulting from calls to the relay manager.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NERelayErrorDomain
	NERelayErrorDomain string
	// NETunnelProviderErrorDomain is the domain used for Tunnel Provider errors.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderErrorDomain
	NETunnelProviderErrorDomain string
	// NEVPNConfigurationChangeNotification is posted after the VPN configuration stored in the Network Extension preferences changes.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConfigurationChangeNotification
	NEVPNConfigurationChangeNotification string
	// NEVPNConnectionErrorDomain is the domain for errors resulting from VPN connection calls.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionErrorDomain
	NEVPNConnectionErrorDomain string
	// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionStartOptionPassword
	NEVPNConnectionStartOptionPassword string
	// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnectionStartOptionUsername
	NEVPNConnectionStartOptionUsername string
	// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNErrorDomain
	NEVPNErrorDomain string
	// NEVPNStatusDidChangeNotification is posted when the status of the VPN connection changes.
	//
	// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatusDidChangeNotification
	NEVPNStatusDidChangeNotification string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEAppProxyErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEAppProxyErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEDNSProxyConfigurationDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEDNSProxyConfigurationDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEDNSProxyErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEDNSProxyErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEDNSSettingsConfigurationDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEDNSSettingsConfigurationDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEDNSSettingsErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEDNSSettingsErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEFilterConfigurationDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEFilterConfigurationDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEFilterErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEFilterErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NERelayClientErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NERelayClientErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NERelayConfigurationDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NERelayConfigurationDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NERelayErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NERelayErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NETunnelProviderErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NETunnelProviderErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEVPNConfigurationChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEVPNConfigurationChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEVPNConnectionErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEVPNConnectionErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEVPNConnectionStartOptionPassword"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEVPNConnectionStartOptionPassword = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEVPNConnectionStartOptionUsername"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEVPNConnectionStartOptionUsername = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEVPNErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEVPNErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NEVPNStatusDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				NEVPNStatusDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

}
