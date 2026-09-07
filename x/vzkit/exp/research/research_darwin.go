//go:build darwin

package research

import (
	"bytes"
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	privatevz "github.com/tmc/apple/private/virtualization"
	vz "github.com/tmc/apple/virtualization"
)

// ProbeHardwareModel checks the Objective-C signatures used to construct the
// research hardware model. It does not instantiate a VM or prove host permission.
func ProbeHardwareModel() error {
	descriptor := objc.GetClass("_VZMacHardwareModelDescriptor")
	model := objc.GetClass("VZMacHardwareModel")
	if descriptor == 0 || model == 0 {
		return fmt.Errorf("ios hardware model classes unavailable")
	}
	for _, method := range []struct{ selector, argument string }{
		{"setPlatformVersion:", "I"}, {"setBoardID:", "I"}, {"setISA:", "q"},
	} {
		m := objectivec.Class_getInstanceMethod(descriptor, objectivec.SEL(objc.Sel(method.selector)))
		if err := checkMethod(m, "v", []string{"@", ":", method.argument}); err != nil {
			return fmt.Errorf("ios descriptor %s: %w", method.selector, err)
		}
	}
	m := objectivec.Class_getClassMethod(model, objectivec.SEL(objc.Sel("_hardwareModelWithDescriptor:")))
	if err := checkMethod(m, "@", []string{"@", ":", "@"}); err != nil {
		return fmt.Errorf("ios hardware model factory: %w", err)
	}
	return nil
}

func checkMethod(m objectivec.Method, result string, args []string) error {
	if m == 0 {
		return fmt.Errorf("selector unavailable")
	}
	if got := objectivec.Method_getNumberOfArguments(m); got != uint32(len(args)) {
		return fmt.Errorf("argument count %d, want %d", got, len(args))
	}
	var buf [256]byte
	objectivec.Method_getReturnType(m, &buf[0], uintptr(len(buf)))
	if got := encodingString(buf[:]); got != result {
		return fmt.Errorf("return encoding %q, want %q", got, result)
	}
	for i, want := range args {
		clear(buf[:])
		objectivec.Method_getArgumentType(m, uint32(i), &buf[0], uintptr(len(buf)))
		if got := encodingString(buf[:]); got != want {
			return fmt.Errorf("argument %d encoding %q, want %q", i, got, want)
		}
	}
	return nil
}

func encodingString(buf []byte) string {
	if i := bytes.IndexByte(buf, 0); i >= 0 {
		return string(buf[:i])
	}
	return string(buf)
}

// NewHardwareModel constructs a model from a private hardware descriptor.
// Call on the runtime's main thread inside an autorelease pool. The caller owns
// the returned reference and must Release it. Signature checks precede private
// calls; successful construction still does not qualify a firmware or VM graph.
func NewHardwareModel(config HardwareDescriptor) (vz.VZMacHardwareModel, error) {
	if config.PlatformVersion == 0 || config.BoardID == 0 || config.ISA == 0 {
		return vz.VZMacHardwareModel{}, fmt.Errorf("incomplete hardware descriptor")
	}
	if err := ProbeHardwareModel(); err != nil {
		return vz.VZMacHardwareModel{}, err
	}
	descriptor := privatevz.NewVZMacHardwareModelDescriptor()
	if descriptor.ID == 0 {
		return vz.VZMacHardwareModel{}, fmt.Errorf("create ios hardware descriptor: nil object")
	}
	defer descriptor.Release()
	descriptor.SetPlatformVersion(config.PlatformVersion)
	descriptor.SetBoardID(config.BoardID)
	descriptor.SetISA(config.ISA)
	object, err := privatevz.GetVZMacHardwareModelClass().HardwareModelWithDescriptor(descriptor)
	if err != nil {
		return vz.VZMacHardwareModel{}, fmt.Errorf("create ios hardware model: %w", err)
	}
	if object == nil || object.GetID() == 0 {
		return vz.VZMacHardwareModel{}, fmt.Errorf("create ios hardware model: nil object")
	}
	model := vz.VZMacHardwareModelFromID(object.GetID())
	if !model.IsSupported() {
		return vz.VZMacHardwareModel{}, fmt.Errorf("ios hardware model unsupported by host")
	}
	model.Retain()
	return model, nil
}

// HardwareDescriptor selects private platform characteristics. Callers choose
// firmware-compatible values; no descriptor is qualified by this package.
type HardwareDescriptor struct {
	PlatformVersion uint32
	BoardID         uint32
	ISA             int64
}

// ECID reads the private chip identifier, checking its Objective-C ABI first.
func ECID(machine vz.VZMacMachineIdentifier) (uint64, error) {
	if machine.ID == 0 {
		return 0, fmt.Errorf("nil machine identifier")
	}
	method := objectivec.Class_getInstanceMethod(objc.GetClass("VZMacMachineIdentifier"), objectivec.SEL(objc.Sel("_ECID")))
	if err := checkMethod(method, "Q", []string{"@", ":"}); err != nil {
		return 0, err
	}
	value, err := privatevz.VZMacMachineIdentifierFromID(machine.ID).ECID()
	if err != nil {
		return 0, err
	}
	if value == 0 {
		return 0, fmt.Errorf("machine identifier has zero ecid")
	}
	return value, nil
}

// BootOptions selects private DFU and iBoot stop points.
type BootOptions struct{ ForceDFU, StopInIBootStage1, StopInIBootStage2 bool }

// ConfigureStart applies private boot options to a caller-owned VZ start object.
// It neither starts a VM nor changes its dispatch queue or lifetime.
func ConfigureStart(start vz.VZMacOSVirtualMachineStartOptions, options BootOptions) error {
	if start.ID == 0 {
		return fmt.Errorf("nil start options")
	}
	private := privatevz.VZMacOSVirtualMachineStartOptionsFromID(start.ID)
	for _, setting := range []struct {
		set   func(bool) error
		value bool
	}{
		{private.SetForceDFU, options.ForceDFU}, {private.SetStopInIBootStage1, options.StopInIBootStage1}, {private.SetStopInIBootStage2, options.StopInIBootStage2},
	} {
		if err := setting.set(setting.value); err != nil {
			return fmt.Errorf("set research boot option: %w", err)
		}
	}
	return nil
}
