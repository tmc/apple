//go:build darwin

package network_test

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/network"
)

func ExampleGetNWEndpointClass() {
	cls := network.GetNWEndpointClass()
	name := objc.GoString(objectivec.Class_getName(cls.Class()))
	fmt.Println("class name:", name)
	fmt.Println("supports secure coding:", cls.SupportsSecureCoding())
	// Output:
	// class name: NWEndpoint
	// supports secure coding: true
}
