// Command whereami prints the device's current location using Core Location.
//
// It requests "when in use" authorization, starts location updates, and waits
// for the first fix. Location access requires user approval; a command-line
// binary that has never been approved reports a denied or restricted status
// instead of a fix.
//
// Usage: whereami [-timeout d]
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/corelocation"
	"github.com/tmc/apple/foundation"
)

func statusString(s corelocation.CLAuthorizationStatus) string {
	switch s {
	case corelocation.KCLAuthorizationStatusNotDetermined:
		return "not determined"
	case corelocation.KCLAuthorizationStatusRestricted:
		return "restricted"
	case corelocation.KCLAuthorizationStatusDenied:
		return "denied"
	case corelocation.KCLAuthorizationStatusAuthorizedAlways:
		return "authorized always"
	case corelocation.KCLAuthorizationStatusAuthorizedWhenInUse:
		return "authorized when in use"
	}
	return fmt.Sprintf("unknown (%d)", int(s))
}

func main() {
	timeout := flag.Duration("timeout", 15*time.Second, "how long to wait for a location fix")
	flag.Parse()

	if !corelocation.GetCLLocationManagerClass().LocationServicesEnabled() {
		fmt.Fprintf(os.Stderr, "location services are disabled; enable them in System Settings > Privacy & Security > Location Services\n")
		os.Exit(1)
	}

	manager := corelocation.NewCLLocationManager()
	manager.SetDesiredAccuracy(corelocation.KCLLocationAccuracyBest)

	var lastErr string
	delegate := corelocation.NewCLLocationManagerDelegate(corelocation.CLLocationManagerDelegateConfig{
		LocationManagerDidChangeAuthorization: func(m corelocation.CLLocationManager) {
			fmt.Fprintf(os.Stderr, "authorization: %s\n", statusString(m.AuthorizationStatus()))
		},
		LocationManagerDidFailWithError: func(_ corelocation.CLLocationManager, err foundation.NSError) {
			lastErr = err.Description()
		},
	})
	manager.SetDelegate(delegate)

	manager.RequestWhenInUseAuthorization()
	manager.StartUpdatingLocation()
	defer manager.StopUpdatingLocation()

	// Pump the run loop so Core Location can deliver callbacks, polling the
	// manager's location property until it has a fix or we time out.
	deadline := time.Now().Add(*timeout)
	for time.Now().Before(deadline) {
		corefoundation.CFRunLoopRunInMode(corefoundation.KCFRunLoopDefaultMode, 0.1, false)
		loc := manager.Location()
		if loc.GetID() == 0 {
			continue
		}
		l := corelocation.CLLocationFromID(loc.GetID())
		c := l.Coordinate()
		fmt.Printf("latitude:   %.6f\n", c.Latitude)
		fmt.Printf("longitude:  %.6f\n", c.Longitude)
		fmt.Printf("altitude:   %.1f m\n", l.Altitude())
		fmt.Printf("horizontal: ±%.1f m\n", l.HorizontalAccuracy())
		fmt.Printf("vertical:   ±%.1f m\n", l.VerticalAccuracy())
		fmt.Printf("timestamp:  %s\n", l.Timestamp().Description())
		return
	}

	status := manager.AuthorizationStatus()
	switch status {
	case corelocation.KCLAuthorizationStatusDenied, corelocation.KCLAuthorizationStatusRestricted:
		fmt.Fprintf(os.Stderr, "no location: authorization %s; approve this binary in System Settings > Privacy & Security > Location Services\n", statusString(status))
	case corelocation.KCLAuthorizationStatusNotDetermined:
		fmt.Fprintf(os.Stderr, "no location: authorization still not determined; a command-line binary may never be prompted\n")
	default:
		fmt.Fprintf(os.Stderr, "no location fix within %v\n", *timeout)
	}
	if lastErr != "" {
		fmt.Fprintf(os.Stderr, "last error: %s\n", lastErr)
	}
	os.Exit(1)
}
