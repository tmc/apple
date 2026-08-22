//go:build darwin

package systemconfiguration_test

import (
	"fmt"

	"github.com/tmc/apple/systemconfiguration"
)

func ExampleKSCStatus() {
	status := systemconfiguration.KSCStatusOK
	fmt.Println(status)

	reachabilityStatus := systemconfiguration.KSCStatusReachabilityUnknown
	fmt.Println(reachabilityStatus)

	// Output:
	// KSCStatusOK
	// KSCStatusReachabilityUnknown
}

func ExampleSCNetworkReachabilityFlags() {
	flags := systemconfiguration.KSCNetworkReachabilityFlagsReachable
	fmt.Println(flags)

	wwanFlags := systemconfiguration.KSCNetworkReachabilityFlagsIsWWAN
	fmt.Println(wwanFlags)

	// Output:
	// KSCNetworkReachabilityFlagsReachable
	// KSCNetworkReachabilityFlagsIsWWAN
}

func ExampleSCNetworkReachabilityGetTypeID() {
	typeID := systemconfiguration.SCNetworkReachabilityGetTypeID()
	if typeID != 0 {
		fmt.Println("SCNetworkReachability type ID retrieved")
	}

	// Output:
	// SCNetworkReachability type ID retrieved
}
