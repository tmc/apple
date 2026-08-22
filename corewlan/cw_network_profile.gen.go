// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CWNetworkProfile] class.
var (
	_CWNetworkProfileClass     CWNetworkProfileClass
	_CWNetworkProfileClassOnce sync.Once
)

func getCWNetworkProfileClass() CWNetworkProfileClass {
	_CWNetworkProfileClassOnce.Do(func() {
		_CWNetworkProfileClass = CWNetworkProfileClass{class: objc.GetClass("CWNetworkProfile")}
	})
	return _CWNetworkProfileClass
}

// GetCWNetworkProfileClass returns the class object for CWNetworkProfile.
func GetCWNetworkProfileClass() CWNetworkProfileClass {
	return getCWNetworkProfileClass()
}

type CWNetworkProfileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CWNetworkProfileClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CWNetworkProfileClass) Alloc() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Encapsulates an immutable network profile entry.
//
// # Getting a network profile
//
//   - [CWNetworkProfile.InitWithNetworkProfile]: Creates and returns a CWNetworkProfile object initialized with the given CWNetworkProfile object.
//
// # Comparing network profiles
//
//   - [CWNetworkProfile.IsEqualToNetworkProfile]: Determine CWNetworkProfile object equality.
//
// # Instance Properties
//
//   - [CWNetworkProfile.Security]: The security mode for the network profile.
//   - [CWNetworkProfile.Ssid]: The service set identifier (SSID) for the network profile, encoded as a string.
//   - [CWNetworkProfile.SsidData]: The service set identifier (SSID) for the network profile, returned as data.
//
// # Initializers
//
//   - [CWNetworkProfile.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile
type CWNetworkProfile struct {
	objectivec.Object
}

// CWNetworkProfileFromID constructs a [CWNetworkProfile] from an objc.ID.
//
// Encapsulates an immutable network profile entry.
func CWNetworkProfileFromID(id objc.ID) CWNetworkProfile {
	return CWNetworkProfile{objectivec.Object{ID: id}}
}

// NOTE: CWNetworkProfile adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CWNetworkProfile] class.
//
// # Getting a network profile
//
//   - [ICWNetworkProfile.InitWithNetworkProfile]: Creates and returns a CWNetworkProfile object initialized with the given CWNetworkProfile object.
//
// # Comparing network profiles
//
//   - [ICWNetworkProfile.IsEqualToNetworkProfile]: Determine CWNetworkProfile object equality.
//
// # Instance Properties
//
//   - [ICWNetworkProfile.Security]: The security mode for the network profile.
//   - [ICWNetworkProfile.Ssid]: The service set identifier (SSID) for the network profile, encoded as a string.
//   - [ICWNetworkProfile.SsidData]: The service set identifier (SSID) for the network profile, returned as data.
//
// # Initializers
//
//   - [ICWNetworkProfile.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile
type ICWNetworkProfile interface {
	objectivec.IObject

	// Topic: Getting a network profile

	// Creates and returns a CWNetworkProfile object initialized with the given CWNetworkProfile object.
	InitWithNetworkProfile(networkProfile ICWNetworkProfile) CWNetworkProfile

	// Topic: Comparing network profiles

	// Determine CWNetworkProfile object equality.
	IsEqualToNetworkProfile(networkProfile ICWNetworkProfile) bool

	// Topic: Instance Properties

	// The security mode for the network profile.
	Security() CWSecurity
	// The service set identifier (SSID) for the network profile, encoded as a string.
	Ssid() string
	// The service set identifier (SSID) for the network profile, returned as data.
	SsidData() foundation.NSData

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CWNetworkProfile

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CWNetworkProfile) Init() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CWNetworkProfile) Autorelease() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWNetworkProfile creates a new CWNetworkProfile instance.
func NewCWNetworkProfile() CWNetworkProfile {
	class := getCWNetworkProfileClass()
	rv := objc.Send[CWNetworkProfile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/init(coder:)
func NewCWNetworkProfileWithCoder(coder foundation.INSCoder) CWNetworkProfile {
	instance := getCWNetworkProfileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CWNetworkProfileFromID(rv)
}

// Creates and returns a CWNetworkProfile object initialized with the given
// CWNetworkProfile object.
//
// networkProfile: The CWNetworkProfile object to use to initialize a new CWNetworkProfile
// object.
//
// # Return Value
//
// A CWNetworkProfile object.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/init(networkProfile:)
func NewCWNetworkProfileWithNetworkProfile(networkProfile ICWNetworkProfile) CWNetworkProfile {
	instance := getCWNetworkProfileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetworkProfile:"), networkProfile)
	return CWNetworkProfileFromID(rv)
}

// Creates and returns a CWNetworkProfile object initialized with the given
// CWNetworkProfile object.
//
// networkProfile: The CWNetworkProfile object to use to initialize a new CWNetworkProfile
// object.
//
// # Return Value
//
// A CWNetworkProfile object.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/init(networkProfile:)
func (c CWNetworkProfile) InitWithNetworkProfile(networkProfile ICWNetworkProfile) CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](c.ID, objc.Sel("initWithNetworkProfile:"), networkProfile)
	return rv
}

// Determine CWNetworkProfile object equality.
//
// networkProfile: The CWNetworkProfile object with which to compare the receiver.
//
// # Return Value
//
// YES if the objects are equal.
//
// # Discussion
//
// CWNetwork objects are considered equal if their corresponding ssidData and
// securityType properties are equal.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/isEqual(to:)
func (c CWNetworkProfile) IsEqualToNetworkProfile(networkProfile ICWNetworkProfile) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEqualToNetworkProfile:"), networkProfile)
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/init(coder:)
func (c CWNetworkProfile) InitWithCoder(coder foundation.INSCoder) CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CWNetworkProfile) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The security mode for the network profile.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/security
func (c CWNetworkProfile) Security() CWSecurity {
	rv := objc.Send[CWSecurity](c.ID, objc.Sel("security"))
	return CWSecurity(rv)
}

// The service set identifier (SSID) for the network profile, encoded as a
// string.
//
// # Discussion
//
// If the SSID can not be encoded as a valid UTF-8 or WinLatin1 string, this
// method returns nil.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/ssid
func (c CWNetworkProfile) Ssid() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ssid"))
	return foundation.NSStringFromID(rv).String()
}

// The service set identifier (SSID) for the network profile, returned as
// data.
//
// # Discussion
//
// The SSID is 1-32 octets.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/ssidData
func (c CWNetworkProfile) SsidData() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ssidData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
