// Code generated from Apple documentation for Foundation. DO NOT EDIT.
//go:build ios
// +build ios

package foundation

import (
	"github.com/tmc/apple/objc"
)

// # Return Value
//
// # May return nil if there is no deletion record available for this sensor
//
// # Discussion
//
// # Returns a sensor stream that contains deletion records of the sensor
//
// This sensor stream should only be used for fetching. All other operations
// will be ignored. Deletion records share the recording and authorization
// state with their parent sensor.
//
// See: https://developer.apple.com/documentation/Foundation/NSString/sr_sensorForDeletionRecordsFromSensor()
func (s NSString) Sr_sensorForDeletionRecordsFromSensor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("sr_sensorForDeletionRecordsFromSensor"))
	return unsafe.Pointer(rv)
}
