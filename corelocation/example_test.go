//go:build darwin

package corelocation_test

import (
	"fmt"

	"github.com/tmc/apple/corelocation"
)

func ExampleCLLocation() {
	loc1 := corelocation.NewLocationWithLatitudeLongitude(37.7749, -122.4194)
	loc2 := corelocation.NewLocationWithLatitudeLongitude(37.7749, -122.4194)

	coord := loc1.Coordinate()
	dist := loc1.DistanceFromLocation(loc2)

	fmt.Printf("Lat: %.4f, Lon: %.4f\n", coord.Latitude, coord.Longitude)
	fmt.Printf("Distance: %.0f meters\n", dist)

	// Output:
	// Lat: 37.7749, Lon: -122.4194
	// Distance: 0 meters
}

func ExampleCLLocationCoordinate2D() {
	coord := corelocation.CLLocationCoordinate2D{
		Latitude:  37.7749,
		Longitude: -122.4194,
	}
	fmt.Printf("Latitude: %.4f, Longitude: %.4f\n", coord.Latitude, coord.Longitude)

	// Output:
	// Latitude: 37.7749, Longitude: -122.4194
}

func ExampleCLCircularGeographicCondition() {
	center := corelocation.CLLocationCoordinate2D{Latitude: 37.7749, Longitude: -122.4194}
	cond := corelocation.NewCircularGeographicConditionWithCenterRadius(center, 500.0)
	fmt.Printf("Center Lat: %.4f, Radius: %.0f meters\n", cond.Center().Latitude, cond.Radius())

	// Output:
	// Center Lat: 37.7749, Radius: 500 meters
}
