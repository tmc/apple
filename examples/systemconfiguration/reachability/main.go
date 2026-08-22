// Command reachability reports whether a host is reachable with the current
// network configuration, using the SystemConfiguration framework.
//
// Usage: reachability <hostname>
package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/tmc/apple/systemconfiguration"
)

// flagNames maps each reachability flag to its short name, in report order.
var flagNames = []struct {
	flag systemconfiguration.SCNetworkReachabilityFlags
	name string
}{
	{systemconfiguration.KSCNetworkReachabilityFlagsTransientConnection, "transient-connection"},
	{systemconfiguration.KSCNetworkReachabilityFlagsReachable, "reachable"},
	{systemconfiguration.KSCNetworkReachabilityFlagsConnectionRequired, "connection-required"},
	{systemconfiguration.KSCNetworkReachabilityFlagsConnectionOnTraffic, "connection-on-traffic"},
	{systemconfiguration.KSCNetworkReachabilityFlagsInterventionRequired, "intervention-required"},
	{systemconfiguration.KSCNetworkReachabilityFlagsConnectionOnDemand, "connection-on-demand"},
	{systemconfiguration.KSCNetworkReachabilityFlagsIsLocalAddress, "is-local-address"},
	{systemconfiguration.KSCNetworkReachabilityFlagsIsDirect, "is-direct"},
	{systemconfiguration.KSCNetworkReachabilityFlagsIsWWAN, "is-wwan"},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: reachability <hostname>\n")
		os.Exit(1)
	}
	host := os.Args[1]

	target := systemconfiguration.SCNetworkReachabilityCreateWithName(0, host)
	if target == 0 {
		fmt.Fprintf(os.Stderr, "cannot create reachability target for %s: %s\n", host, lastError())
		os.Exit(1)
	}

	var flags systemconfiguration.SCNetworkReachabilityFlags
	if !systemconfiguration.SCNetworkReachabilityGetFlags(target, &flags) {
		fmt.Fprintf(os.Stderr, "cannot get reachability flags for %s: %s\n", host, lastError())
		os.Exit(1)
	}

	reachable := flags&systemconfiguration.KSCNetworkReachabilityFlagsReachable != 0 &&
		flags&systemconfiguration.KSCNetworkReachabilityFlagsConnectionRequired == 0
	fmt.Printf("host:      %s\n", host)
	fmt.Printf("reachable: %t\n", reachable)
	fmt.Printf("flags:     0x%08x", uint32(flags))
	for _, f := range flagNames {
		if flags&f.flag != 0 {
			fmt.Printf(" %s", f.name)
		}
	}
	fmt.Println()
}

// lastError returns the message for the most recent SystemConfiguration error.
func lastError() string {
	code := systemconfiguration.SCError()
	msg := systemconfiguration.SCErrorString(code)
	if msg == nil {
		return fmt.Sprintf("error %d", code)
	}
	return goString(msg)
}

// goString converts a NUL-terminated C string to a Go string.
func goString(p *byte) string {
	var b []byte
	for q := p; *q != 0; q = (*byte)(unsafe.Add(unsafe.Pointer(q), 1)) {
		b = append(b, *q)
	}
	return string(b)
}
