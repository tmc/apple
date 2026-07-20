// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

import (
	"unsafe"
)

// C struct types

// SCDynamicStoreContext - Structure containing user-specified data and callbacks for a dynamic store session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreContext
type SCDynamicStoreContext struct {
	Version         int                                 // The version number of the structure type being passed in as a parameter to the [SCDynamicStore] creation function (such as [SCDynamicStoreCreate(_:_:_:_:)](<doc://com.apple.systemconfiguration/documentation/SystemConfiguration/SCDynamicStoreCreate(_:_:_:_:)>)). This structure is version `0`.
	Info            unsafe.Pointer                      // A C pointer to a user-specified block of data.
	CopyDescription func(unsafe.Pointer) uintptr        // The callback used to provide a description of the `info` field.
	Release         func(unsafe.Pointer)                // The callback used to remove a retain previously added for the `info` field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value of this parameter can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // The callback used to add a retain for the `info` field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value of this parameter can be [NULL].

}

// SCNetworkConnectionContext - A structure containing user-specified data and callbacks for a network connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionContext
type SCNetworkConnectionContext struct {
	Version         int                                 // The version number of the structure type being passed in as a parameter to the [SCNetworkConnectionCreateWithServiceID(_:_:_:_:)](<doc://com.apple.systemconfiguration/documentation/SystemConfiguration/SCNetworkConnectionCreateWithServiceID(_:_:_:_:)>) function. This structure is version `0`.
	Info            unsafe.Pointer                      // A C pointer to a user-specified block of data.
	CopyDescription func(unsafe.Pointer) uintptr        // The callback used to provide a description of the `info` field.
	Release         func(unsafe.Pointer)                // The calllback used to remove a retain previously added for the info field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value may be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // The callback used to add a retain for the info field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value may be [NULL].

}

// SCNetworkReachabilityContext - Structure containing user-specified data and callbacks used with [SCNetworkReachabilitySetCallback(_:_:_:)](<doc://com.apple.systemconfiguration/documentation/SystemConfiguration/SCNetworkReachabilitySetCallback(_:_:_:)>).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityContext
type SCNetworkReachabilityContext struct {
	Version         int                                 // The version number of the structure type being passed in as a parameter to an [SCDynamicStore] creation function. This structure is version `0`.
	Info            unsafe.Pointer                      // A C pointer to a user-specified block of data.
	CopyDescription func(unsafe.Pointer) uintptr        // The callback used to provide a description of the `info` field.
	Release         func(unsafe.Pointer)                // The callback used to remove a retain previously added for the info field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value can be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // The callback used to add a retain for the info field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value can be [NULL].

}

// SCPreferencesContext - A structure containing user-specified data and callbacks for accessing system configuration preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesContext
type SCPreferencesContext struct {
	Version         int                                 // The version number of the structure type being passed in as a parameter to [SCPreferencesSetCallback(_:_:_:)](<doc://com.apple.systemconfiguration/documentation/SystemConfiguration/SCPreferencesSetCallback(_:_:_:)>). This structure is version `0`.
	Info            unsafe.Pointer                      // A C pointer to a user-specified block of data.
	CopyDescription func(unsafe.Pointer) uintptr        // The callback used to provide a description of the `info` field.
	Release         func(unsafe.Pointer)                // The calllback used to remove a retain previously added for the `info` field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value may be [NULL].
	Retain          func(unsafe.Pointer) unsafe.Pointer // The callback used to add a retain for the `info` field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value may be [NULL].

}
