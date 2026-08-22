// Code generated from Apple documentation. DO NOT EDIT.

package systemconfiguration

import (
	"unsafe"
)

// SCBondInterfaceRef is the reference to an object that represents an Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterface
type SCBondInterfaceRef uintptr

// SCBondStatusRef is the reference to an object that represents the status of an Ethernet bond interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatus
type SCBondStatusRef uintptr

// SCDynamicStoreCallBack is callback used when notification of changes made to the dynamic store is delivered.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCallBack
type SCDynamicStoreCallBack = func(SCDynamicStoreRef, uintptr, unsafe.Pointer)

// SCDynamicStoreRef is the handle to an open dynamic store session with the system configuration daemon.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStore
type SCDynamicStoreRef uintptr

// SCNetworkConnectionCallBack is the type of callback function used when a status event is delivered.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCallBack
type SCNetworkConnectionCallBack = func(connection SCNetworkConnectionRef, status SCNetworkConnectionStatus, info unsafe.Pointer)

// SCNetworkConnectionFlags is flags that indicate whether the specified network node name or address is reachable, whether a connection is required, and whether some user intervention may be required when establishing a connection.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionFlags
type SCNetworkConnectionFlags = uint32

// SCNetworkConnectionRef is the handle to manage a connection-oriented service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnection
type SCNetworkConnectionRef uintptr

// SCNetworkInterfaceRef is the reference to an object that represents a network interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterface
type SCNetworkInterfaceRef uintptr

// SCNetworkProtocolRef is the reference to an object that represents a network protocol.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocol
type SCNetworkProtocolRef uintptr

// SCNetworkReachabilityCallBack is type of callback function used when the reachability of a network address or name changes.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCallBack
type SCNetworkReachabilityCallBack = func(target SCNetworkReachabilityRef, flags uint, info unsafe.Pointer)

// SCNetworkReachabilityRef is the handle to a network address or name.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachability
type SCNetworkReachabilityRef uintptr

// SCNetworkServiceRef is the reference to an object that represents a network service.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkService
type SCNetworkServiceRef uintptr

// SCNetworkSetRef is the reference to an object that represents a network set.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSet
type SCNetworkSetRef uintptr

// SCPreferencesCallBack is type of the callback function used when the preferences have been updated or applied.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCallBack
type SCPreferencesCallBack = func(prefs SCPreferencesRef, notificationType SCPreferencesNotification, info unsafe.Pointer)

// SCPreferencesRef is the handle to an open preferences session for accessing system configuration preferences.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCPreferences
type SCPreferencesRef uintptr

// SCVLANInterfaceRef is the reference to an object that represents a virtual LAN (VLAN) interface.
//
// See: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterface
type SCVLANInterfaceRef uintptr
