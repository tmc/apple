// Code generated from Apple documentation. DO NOT EDIT.

package metal

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// uint values.
const (

	//
	// See: https://developer.apple.com/documentation/Metal/MTLAttributeStrideStatic
	MTLAttributeStrideStatic uint = 9223372036854775807

	//
	// See: https://developer.apple.com/documentation/Metal/MTLBufferLayoutStrideDynamic
	MTLBufferLayoutStrideDynamic uint = 9223372036854775807
)

var (
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueueErrorDomain
	MTL4CommandQueueErrorDomain foundation.NSErrorDomain
	// MTLBinaryArchiveDomain is the domain for Metal binary archive errors.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveDomain
	MTLBinaryArchiveDomain foundation.NSErrorDomain
	// MTLCaptureErrorDomain is the error domain for capture errors.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCaptureErrorDomain
	MTLCaptureErrorDomain foundation.NSErrorDomain
	// MTLCommandBufferEncoderInfoErrorKey is a key to a command buffer error’s user information dictionary that retrieves additional information about a GPU’s runtime error.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferEncoderInfoErrorKey
	MTLCommandBufferEncoderInfoErrorKey foundation.NSErrorUserInfoKey
	// MTLCommandBufferErrorDomain is the domain for Metal command buffer errors.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferErrorDomain
	MTLCommandBufferErrorDomain foundation.NSErrorDomain
	// MTLCounterErrorDomain is the domain for Metal counter sample buffer errors.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCounterErrorDomain
	MTLCounterErrorDomain foundation.NSErrorDomain
	// See: https://developer.apple.com/documentation/Metal/MTLDeviceErrorDomain
	MTLDeviceErrorDomain foundation.NSErrorDomain
	// MTLDynamicLibraryDomain is the domain for Metal dynamic library errors.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryDomain
	MTLDynamicLibraryDomain foundation.NSErrorDomain
	// MTLIOErrorDomain is the domain for input/output command queue errors.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIOErrorDomain
	MTLIOErrorDomain foundation.NSErrorDomain
	// MTLLibraryErrorDomain is the error domain for Metal libraries.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibraryErrorDomain
	MTLLibraryErrorDomain foundation.NSErrorDomain
	// See: https://developer.apple.com/documentation/Metal/MTLLogStateErrorDomain
	MTLLogStateErrorDomain foundation.NSErrorDomain
	// MTLTensorDomain is an error domain for errors that pertain to creating a tensor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensorDomain
	MTLTensorDomain foundation.NSErrorDomain
	// See: https://developer.apple.com/documentation/Metal/NSProcessInfoPerformanceProfileDidChangeNotification
	ProcessInfoPerformanceProfileDidChangeNotification foundation.NSNotificationName
)

var (
)

var (
)

var (
	// MTLDeviceRemovalRequestedNotification is a notification that Metal sends to observers when a person requests to remove a GPU device from the system.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDeviceNotificationName/removalRequested
	MTLDeviceRemovalRequestedNotification MTLDeviceNotificationName
	// MTLDeviceWasAddedNotification is a notification that Metal sends to observers when the system adds a GPU device.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDeviceNotificationName/wasAdded
	MTLDeviceWasAddedNotification MTLDeviceNotificationName
	// MTLDeviceWasRemovedNotification is a notification that Metal sends to observers when the system removes a GPU device.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLDeviceNotificationName/wasRemoved
	MTLDeviceWasRemovedNotification MTLDeviceNotificationName
)

var (
	// DeviceCertificationiPhonePerformanceGaming is the performance gaming tier for iPhone.
	//
	// See: https://developer.apple.com/documentation/Metal/NSDeviceCertification/iPhonePerformanceGaming
	DeviceCertificationiPhonePerformanceGaming NSDeviceCertification
)

var (
)
func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTL4CommandQueueErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTL4CommandQueueErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLBinaryArchiveDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLBinaryArchiveDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCaptureErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLCaptureErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommandBufferEncoderInfoErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLCommandBufferEncoderInfoErrorKey = foundation.NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommandBufferErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLCommandBufferErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterClipperInvocations"); err == nil && ptr != 0 {
		MTLCommonCounters.ClipperInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterClipperPrimitivesOut"); err == nil && ptr != 0 {
		MTLCommonCounters.ClipperPrimitivesOut = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterComputeKernelInvocations"); err == nil && ptr != 0 {
		MTLCommonCounters.ComputeKernelInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterFragmentCycles"); err == nil && ptr != 0 {
		MTLCommonCounters.FragmentCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterFragmentInvocations"); err == nil && ptr != 0 {
		MTLCommonCounters.FragmentInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterFragmentsPassed"); err == nil && ptr != 0 {
		MTLCommonCounters.FragmentsPassed = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterPostTessellationVertexCycles"); err == nil && ptr != 0 {
		MTLCommonCounters.PostTessellationVertexCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterPostTessellationVertexInvocations"); err == nil && ptr != 0 {
		MTLCommonCounters.PostTessellationVertexInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterRenderTargetWriteCycles"); err == nil && ptr != 0 {
		MTLCommonCounters.RenderTargetWriteCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterSetStageUtilization"); err == nil && ptr != 0 {
		MTLCommonCounterSets.StageUtilization = *(*MTLCommonCounterSet)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterSetStatistic"); err == nil && ptr != 0 {
		MTLCommonCounterSets.Statistic = *(*MTLCommonCounterSet)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterSetTimestamp"); err == nil && ptr != 0 {
		MTLCommonCounterSets.Timestamp = *(*MTLCommonCounterSet)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterTessellationCycles"); err == nil && ptr != 0 {
		MTLCommonCounters.TessellationCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterTessellationInputPatches"); err == nil && ptr != 0 {
		MTLCommonCounters.TessellationInputPatches = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterTimestamp"); err == nil && ptr != 0 {
		MTLCommonCounters.Timestamp = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterTotalCycles"); err == nil && ptr != 0 {
		MTLCommonCounters.TotalCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterVertexCycles"); err == nil && ptr != 0 {
		MTLCommonCounters.VertexCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterVertexInvocations"); err == nil && ptr != 0 {
		MTLCommonCounters.VertexInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCounterErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLCounterErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLDeviceErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLDeviceErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLDeviceRemovalRequestedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLDeviceRemovalRequestedNotification = MTLDeviceNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLDeviceWasAddedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLDeviceWasAddedNotification = MTLDeviceNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLDeviceWasRemovedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLDeviceWasRemovedNotification = MTLDeviceNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLDynamicLibraryDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLDynamicLibraryDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLIOErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLIOErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLLibraryErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLLibraryErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLLogStateErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLLogStateErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLTensorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLTensorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSDeviceCertificationiPhonePerformanceGaming"); err == nil && ptr != 0 {
		DeviceCertificationiPhonePerformanceGaming = *(*NSDeviceCertification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProcessInfoPerformanceProfileDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProcessInfoPerformanceProfileDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProcessPerformanceProfileDefault"); err == nil && ptr != 0 {
		NSProcessPerformanceProfiles.Default = *(*NSProcessPerformanceProfile)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProcessPerformanceProfileSustained"); err == nil && ptr != 0 {
		NSProcessPerformanceProfiles.Sustained = *(*NSProcessPerformanceProfile)(unsafe.Pointer(ptr))
	}

}

// MTLCommonCounterSets provides typed accessors for [MTLCommonCounterSet] constants.
var MTLCommonCounterSets struct {
	// StageUtilization: The common name for the counter set that contains hardware utilization measurements from various render stages.
	StageUtilization MTLCommonCounterSet
	// Statistic: The common name for the counter set that contains GPU workload statistics.
	Statistic MTLCommonCounterSet
	// Timestamp: The common name for the counter set that contains the timestamp counter.
	Timestamp MTLCommonCounterSet
}

// MTLCommonCounters provides typed accessors for [MTLCommonCounter] constants.
var MTLCommonCounters struct {
	// ClipperInvocations: The common name for the counter that tracks the number of primitives a render pass sends to the clip stage.
	ClipperInvocations MTLCommonCounter
	// ClipperPrimitivesOut: The common name for the counter that tracks the number of primitives the clip stage produces during a render pass.
	ClipperPrimitivesOut MTLCommonCounter
	// ComputeKernelInvocations: The common name for the counter that tracks the number of times a pass invokes any compute kernel.
	ComputeKernelInvocations MTLCommonCounter
	// FragmentCycles: The common name for the counter that tracks the number of cycles the GPU uses to run fragment shaders during a pass.
	FragmentCycles MTLCommonCounter
	// FragmentInvocations: The common name for the counter that tracks the number of times a render pass calls fragment shaders.
	FragmentInvocations MTLCommonCounter
	// FragmentsPassed: The common name for the counter that tracks the number of fragments a render pass sends to the visibility and blend stages.
	FragmentsPassed MTLCommonCounter
	// PostTessellationVertexCycles: The common name for the counter that tracks the number of cycles the GPU uses to run post-tessellation vertex shaders during a pass.
	PostTessellationVertexCycles MTLCommonCounter
	// PostTessellationVertexInvocations: The common name for the counter that tracks the number of vertices a render pass sends to a post-tessellation vertex shader.
	PostTessellationVertexInvocations MTLCommonCounter
	// RenderTargetWriteCycles: The common name for the counter that tracks the number of cycles the GPU uses to write data to render targets during a render pass.
	RenderTargetWriteCycles MTLCommonCounter
	// TessellationCycles: The common name for the counter that tracks the number of cycles the GPU uses to run the tessellation stage during a pass.
	TessellationCycles MTLCommonCounter
	// TessellationInputPatches: The common name for the counter that tracks the number of tessellation patches a render pass sends to the tessellation stage.
	TessellationInputPatches MTLCommonCounter
	// Timestamp: The common name for the counter that tracks the current time.
	Timestamp MTLCommonCounter
	// TotalCycles: The common name for the counter that tracks the total number of cycles the GPU uses to run a pass.
	TotalCycles MTLCommonCounter
	// VertexCycles: The common name for the counter that tracks the number of cycles the GPU uses to run vertex shaders during a pass.
	VertexCycles MTLCommonCounter
	// VertexInvocations: The common name for the counter that tracks the number of times a render pass calls any vertex shader.
	VertexInvocations MTLCommonCounter
}

// NSProcessPerformanceProfiles provides typed accessors for [NSProcessPerformanceProfile] constants.
var NSProcessPerformanceProfiles struct {
	// Default: The default performance profile for a device.
	Default NSProcessPerformanceProfile
	// Sustained: The performance profile for a device representing sustained performance.
	Sustained NSProcessPerformanceProfile
}

