// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphExecutableSerializationDescriptor] class.
var (
	_MPSGraphExecutableSerializationDescriptorClass     MPSGraphExecutableSerializationDescriptorClass
	_MPSGraphExecutableSerializationDescriptorClassOnce sync.Once
)

func getMPSGraphExecutableSerializationDescriptorClass() MPSGraphExecutableSerializationDescriptorClass {
	_MPSGraphExecutableSerializationDescriptorClassOnce.Do(func() {
		_MPSGraphExecutableSerializationDescriptorClass = MPSGraphExecutableSerializationDescriptorClass{class: objc.GetClass("MPSGraphExecutableSerializationDescriptor")}
	})
	return _MPSGraphExecutableSerializationDescriptorClass
}

// GetMPSGraphExecutableSerializationDescriptorClass returns the class object for MPSGraphExecutableSerializationDescriptor.
func GetMPSGraphExecutableSerializationDescriptorClass() MPSGraphExecutableSerializationDescriptorClass {
	return getMPSGraphExecutableSerializationDescriptorClass()
}

type MPSGraphExecutableSerializationDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphExecutableSerializationDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphExecutableSerializationDescriptorClass) Alloc() MPSGraphExecutableSerializationDescriptor {
	rv := objc.Send[MPSGraphExecutableSerializationDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that consists of all the levers to serialize an executable.
//
// # Instance Properties
//
//   - [MPSGraphExecutableSerializationDescriptor.Append]: Flag to append to an existing .mpsgraphpackage if found at provided url.
//   - [MPSGraphExecutableSerializationDescriptor.SetAppend]
//   - [MPSGraphExecutableSerializationDescriptor.DeploymentPlatform]: The deployment platform used to serialize the executable.
//   - [MPSGraphExecutableSerializationDescriptor.SetDeploymentPlatform]
//   - [MPSGraphExecutableSerializationDescriptor.MinimumDeploymentTarget]: The minimum deployment target to serialize the executable.
//   - [MPSGraphExecutableSerializationDescriptor.SetMinimumDeploymentTarget]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor
type MPSGraphExecutableSerializationDescriptor struct {
	MPSGraphObject
}

// MPSGraphExecutableSerializationDescriptorFromID constructs a [MPSGraphExecutableSerializationDescriptor] from an objc.ID.
//
// A class that consists of all the levers to serialize an executable.
func MPSGraphExecutableSerializationDescriptorFromID(id objc.ID) MPSGraphExecutableSerializationDescriptor {
	return MPSGraphExecutableSerializationDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphExecutableSerializationDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphExecutableSerializationDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphExecutableSerializationDescriptor.Append]: Flag to append to an existing .mpsgraphpackage if found at provided url.
//   - [IMPSGraphExecutableSerializationDescriptor.SetAppend]
//   - [IMPSGraphExecutableSerializationDescriptor.DeploymentPlatform]: The deployment platform used to serialize the executable.
//   - [IMPSGraphExecutableSerializationDescriptor.SetDeploymentPlatform]
//   - [IMPSGraphExecutableSerializationDescriptor.MinimumDeploymentTarget]: The minimum deployment target to serialize the executable.
//   - [IMPSGraphExecutableSerializationDescriptor.SetMinimumDeploymentTarget]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor
type IMPSGraphExecutableSerializationDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// Flag to append to an existing .mpsgraphpackage if found at provided url.
	Append() bool
	SetAppend(value bool)
	// The deployment platform used to serialize the executable.
	DeploymentPlatform() MPSGraphDeploymentPlatform
	SetDeploymentPlatform(value MPSGraphDeploymentPlatform)
	// The minimum deployment target to serialize the executable.
	MinimumDeploymentTarget() string
	SetMinimumDeploymentTarget(value string)
}

// Init initializes the instance.
func (g MPSGraphExecutableSerializationDescriptor) Init() MPSGraphExecutableSerializationDescriptor {
	rv := objc.Send[MPSGraphExecutableSerializationDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphExecutableSerializationDescriptor) Autorelease() MPSGraphExecutableSerializationDescriptor {
	rv := objc.Send[MPSGraphExecutableSerializationDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphExecutableSerializationDescriptor creates a new MPSGraphExecutableSerializationDescriptor instance.
func NewMPSGraphExecutableSerializationDescriptor() MPSGraphExecutableSerializationDescriptor {
	class := getMPSGraphExecutableSerializationDescriptorClass()
	rv := objc.Send[MPSGraphExecutableSerializationDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Flag to append to an existing .mpsgraphpackage if found at provided url.
//
// # Discussion
//
// If false, the exisiting .mpsgraphpackage will be overwritten.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/append
func (g MPSGraphExecutableSerializationDescriptor) Append() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("append"))
	return rv
}
func (g MPSGraphExecutableSerializationDescriptor) SetAppend(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setAppend:"), value)
}

// The deployment platform used to serialize the executable.
//
// # Discussion
//
// Defaults to the current platform.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/deploymentPlatform
func (g MPSGraphExecutableSerializationDescriptor) DeploymentPlatform() MPSGraphDeploymentPlatform {
	rv := objc.Send[MPSGraphDeploymentPlatform](g.ID, objc.Sel("deploymentPlatform"))
	return MPSGraphDeploymentPlatform(rv)
}
func (g MPSGraphExecutableSerializationDescriptor) SetDeploymentPlatform(value MPSGraphDeploymentPlatform) {
	objc.Send[struct{}](g.ID, objc.Sel("setDeploymentPlatform:"), value)
}

// The minimum deployment target to serialize the executable.
//
// # Discussion
//
// If not set, the package created will target the latest version of the
// `deploymentPlatform` set.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/minimumDeploymentTarget
func (g MPSGraphExecutableSerializationDescriptor) MinimumDeploymentTarget() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("minimumDeploymentTarget"))
	return foundation.NSStringFromID(rv).String()
}
func (g MPSGraphExecutableSerializationDescriptor) SetMinimumDeploymentTarget(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setMinimumDeploymentTarget:"), objc.String(value))
}
