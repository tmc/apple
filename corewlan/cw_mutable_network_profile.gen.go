// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CWMutableNetworkProfile] class.
var (
	_CWMutableNetworkProfileClass     CWMutableNetworkProfileClass
	_CWMutableNetworkProfileClassOnce sync.Once
)

func getCWMutableNetworkProfileClass() CWMutableNetworkProfileClass {
	_CWMutableNetworkProfileClassOnce.Do(func() {
		_CWMutableNetworkProfileClass = CWMutableNetworkProfileClass{class: objc.GetClass("CWMutableNetworkProfile")}
	})
	return _CWMutableNetworkProfileClass
}

// GetCWMutableNetworkProfileClass returns the class object for CWMutableNetworkProfile.
func GetCWMutableNetworkProfileClass() CWMutableNetworkProfileClass {
	return getCWMutableNetworkProfileClass()
}

type CWMutableNetworkProfileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CWMutableNetworkProfileClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CWMutableNetworkProfileClass) Alloc() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Encapsulates a mutable network profile entry.
//
// # Overview
//
// Use this class to change profile properties. To commit Wi-Fi network
// profile changes, use [CWMutableConfiguration.NetworkProfiles] and
// [CWInterface.CommitConfigurationAuthorizationError].
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile
type CWMutableNetworkProfile struct {
	CWNetworkProfile
}

// CWMutableNetworkProfileFromID constructs a [CWMutableNetworkProfile] from an objc.ID.
//
// Encapsulates a mutable network profile entry.
func CWMutableNetworkProfileFromID(id objc.ID) CWMutableNetworkProfile {
	return CWMutableNetworkProfile{CWNetworkProfile: CWNetworkProfileFromID(id)}
}

// NOTE: CWMutableNetworkProfile adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CWMutableNetworkProfile] class.
//
// See: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile
type ICWMutableNetworkProfile interface {
	ICWNetworkProfile
}

// Init initializes the instance.
func (c CWMutableNetworkProfile) Init() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CWMutableNetworkProfile) Autorelease() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWMutableNetworkProfile creates a new CWMutableNetworkProfile instance.
func NewCWMutableNetworkProfile() CWMutableNetworkProfile {
	class := getCWMutableNetworkProfileClass()
	rv := objc.Send[CWMutableNetworkProfile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/init(coder:)
func NewCWMutableNetworkProfileWithCoder(coder foundation.INSCoder) CWMutableNetworkProfile {
	instance := getCWMutableNetworkProfileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CWMutableNetworkProfileFromID(rv)
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
func NewCWMutableNetworkProfileWithNetworkProfile(networkProfile ICWNetworkProfile) CWMutableNetworkProfile {
	instance := getCWMutableNetworkProfileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetworkProfile:"), networkProfile)
	return CWMutableNetworkProfileFromID(rv)
}
