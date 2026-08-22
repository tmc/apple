//go:build darwin

package diskarbitration_test

import (
	"fmt"

	"github.com/tmc/apple/diskarbitration"
)

func ExampleDASessionGetTypeID() {
	typeID := diskarbitration.DASessionGetTypeID()
	if typeID != 0 {
		fmt.Println("DASession type ID retrieved")
	}

	// Output:
	// DASession type ID retrieved
}

func ExampleDADiskGetTypeID() {
	typeID := diskarbitration.DADiskGetTypeID()
	if typeID != 0 {
		fmt.Println("DADisk type ID retrieved")
	}

	// Output:
	// DADisk type ID retrieved
}

func ExampleDASessionCreate() {
	session := diskarbitration.DASessionCreate(0)
	if session != 0 {
		fmt.Println("DASession created successfully")
	}

	// Output:
	// DASession created successfully
}

func ExampleDADissenterCreate() {
	dissenter := diskarbitration.DADissenterCreate(0, 0, 0)
	status := diskarbitration.DADissenterGetStatus(dissenter)
	fmt.Printf("Dissenter status: %d\n", status)

	// Output:
	// Dissenter status: 0
}
