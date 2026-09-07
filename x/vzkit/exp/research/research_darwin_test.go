//go:build darwin

package research

import (
	"fmt"
	vz "github.com/tmc/apple/virtualization"
	"testing"
)

func TestNilInputs(t *testing.T) {
	if _, err := ECID(vz.VZMacMachineIdentifier{}); err == nil {
		t.Fatal("nil identity accepted")
	}
	if err := ConfigureStart(vz.VZMacOSVirtualMachineStartOptions{}, BootOptions{}); err == nil {
		t.Fatal("nil options accepted")
	}
	if _, err := NewHardwareModel(HardwareDescriptor{}); err == nil {
		t.Fatal("empty descriptor accepted")
	}
}
func ExampleNewHardwareModel() {
	_, err := NewHardwareModel(HardwareDescriptor{})
	fmt.Println(err)
	// Output: incomplete hardware descriptor
}
func ExampleECID() {
	_, err := ECID(vz.VZMacMachineIdentifier{})
	fmt.Println(err)
	// Output: nil machine identifier
}
func ExampleConfigureStart() {
	err := ConfigureStart(vz.VZMacOSVirtualMachineStartOptions{}, BootOptions{})
	fmt.Println(err)
	// Output: nil start options
}
