// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKFetchRecordZoneChangesConfiguration] class.
var (
	_CKFetchRecordZoneChangesConfigurationClass     CKFetchRecordZoneChangesConfigurationClass
	_CKFetchRecordZoneChangesConfigurationClassOnce sync.Once
)

func getCKFetchRecordZoneChangesConfigurationClass() CKFetchRecordZoneChangesConfigurationClass {
	_CKFetchRecordZoneChangesConfigurationClassOnce.Do(func() {
		_CKFetchRecordZoneChangesConfigurationClass = CKFetchRecordZoneChangesConfigurationClass{class: objc.GetClass("CKFetchRecordZoneChangesConfiguration")}
	})
	return _CKFetchRecordZoneChangesConfigurationClass
}

// GetCKFetchRecordZoneChangesConfigurationClass returns the class object for CKFetchRecordZoneChangesConfiguration.
func GetCKFetchRecordZoneChangesConfigurationClass() CKFetchRecordZoneChangesConfigurationClass {
	return getCKFetchRecordZoneChangesConfigurationClass()
}

type CKFetchRecordZoneChangesConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKFetchRecordZoneChangesConfigurationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKFetchRecordZoneChangesConfigurationClass) Alloc() CKFetchRecordZoneChangesConfiguration {
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A configuration object that describes the information to fetch from a
// record zone.
//
// # Accessing a Zone Change Configuration
//
//   - [CKFetchRecordZoneChangesConfiguration.PreviousServerChangeToken]: The server change token.
//   - [CKFetchRecordZoneChangesConfiguration.SetPreviousServerChangeToken]
//   - [CKFetchRecordZoneChangesConfiguration.ResultsLimit]: The maximum number of records that CloudKit retrieves when fetching zone changes.
//   - [CKFetchRecordZoneChangesConfiguration.SetResultsLimit]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration
type CKFetchRecordZoneChangesConfiguration struct {
	objectivec.Object
}

// CKFetchRecordZoneChangesConfigurationFromID constructs a [CKFetchRecordZoneChangesConfiguration] from an objc.ID.
//
// A configuration object that describes the information to fetch from a
// record zone.
func CKFetchRecordZoneChangesConfigurationFromID(id objc.ID) CKFetchRecordZoneChangesConfiguration {
	return CKFetchRecordZoneChangesConfiguration{objectivec.Object{ID: id}}
}

// NOTE: CKFetchRecordZoneChangesConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKFetchRecordZoneChangesConfiguration] class.
//
// # Accessing a Zone Change Configuration
//
//   - [ICKFetchRecordZoneChangesConfiguration.PreviousServerChangeToken]: The server change token.
//   - [ICKFetchRecordZoneChangesConfiguration.SetPreviousServerChangeToken]
//   - [ICKFetchRecordZoneChangesConfiguration.ResultsLimit]: The maximum number of records that CloudKit retrieves when fetching zone changes.
//   - [ICKFetchRecordZoneChangesConfiguration.SetResultsLimit]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration
type ICKFetchRecordZoneChangesConfiguration interface {
	objectivec.IObject

	// Topic: Accessing a Zone Change Configuration

	// The server change token.
	PreviousServerChangeToken() ICKServerChangeToken
	SetPreviousServerChangeToken(value ICKServerChangeToken)
	// The maximum number of records that CloudKit retrieves when fetching zone changes.
	ResultsLimit() uint
	SetResultsLimit(value uint)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKFetchRecordZoneChangesConfiguration) Init() CKFetchRecordZoneChangesConfiguration {
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKFetchRecordZoneChangesConfiguration) Autorelease() CKFetchRecordZoneChangesConfiguration {
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordZoneChangesConfiguration creates a new CKFetchRecordZoneChangesConfiguration instance.
func NewCKFetchRecordZoneChangesConfiguration() CKFetchRecordZoneChangesConfiguration {
	class := getCKFetchRecordZoneChangesConfigurationClass()
	rv := objc.Send[CKFetchRecordZoneChangesConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKFetchRecordZoneChangesConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The server change token.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/previousServerChangeToken
func (c CKFetchRecordZoneChangesConfiguration) PreviousServerChangeToken() ICKServerChangeToken {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("previousServerChangeToken"))
	return CKServerChangeTokenFromID(objc.ID(rv))
}
func (c CKFetchRecordZoneChangesConfiguration) SetPreviousServerChangeToken(value ICKServerChangeToken) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}

// The maximum number of records that CloudKit retrieves when fetching zone
// changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/ZoneConfiguration/resultsLimit
func (c CKFetchRecordZoneChangesConfiguration) ResultsLimit() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("resultsLimit"))
	return rv
}
func (c CKFetchRecordZoneChangesConfiguration) SetResultsLimit(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setResultsLimit:"), value)
}
