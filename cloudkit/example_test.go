//go:build darwin

package cloudkit_test

import (
	"fmt"

	"github.com/tmc/apple/cloudkit"
)

func ExampleCKAccountStatus() {
	status := cloudkit.CKAccountStatusAvailable
	fmt.Println(status)
	// Output:
	// CKAccountStatusAvailable
}

func ExampleCKDatabaseScope() {
	scope := cloudkit.CKDatabaseScopePublic
	fmt.Println(scope)
	// Output:
	// CKDatabaseScopePublic
}

func ExampleCKRecordID() {
	zoneID := cloudkit.GetCKRecordZoneIDClass().Alloc().InitWithZoneNameOwnerName("custom-zone", "owner-42")
	recordID := cloudkit.GetCKRecordIDClass().Alloc().InitWithRecordNameZoneID("record-123", zoneID)

	fmt.Println("RecordName:", recordID.RecordName())
	fmt.Println("ZoneName:", recordID.ZoneID().ZoneName())
	fmt.Println("OwnerName:", recordID.ZoneID().OwnerName())

	// Output:
	// RecordName: record-123
	// ZoneName: custom-zone
	// OwnerName: owner-42
}

func ExampleCKErrorCode() {
	code := cloudkit.CKErrorNetworkUnavailable
	fmt.Println(code)
	// Output:
	// CKErrorNetworkUnavailable
}
