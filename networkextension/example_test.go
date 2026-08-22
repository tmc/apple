//go:build darwin

package networkextension_test

import (
	"fmt"

	"github.com/tmc/apple/networkextension"
)

func ExampleNEDNSProtocol() {
	https := networkextension.NEDNSProtocolHTTPS
	fmt.Println(https)

	tls := networkextension.NEDNSProtocolTLS
	fmt.Println(tls)

	// Output:
	// NEDNSProtocolHTTPS
	// NEDNSProtocolTLS
}

func ExampleNEFilterAction() {
	allow := networkextension.NEFilterActionAllow
	fmt.Println(allow)

	drop := networkextension.NEFilterActionDrop
	fmt.Println(drop)

	// Output:
	// NEFilterActionAllow
	// NEFilterActionDrop
}

func ExampleGetNEVPNManagerClass() {
	class := networkextension.GetNEVPNManagerClass()
	if class.Class() != 0 {
		fmt.Println("NEVPNManager class loaded")
	}

	// Output:
	// NEVPNManager class loaded
}
