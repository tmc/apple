// Code generated from Apple documentation. DO NOT EDIT.

package metal

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

const (

	MTLAttributeStrideStatic uint = 9223372036854775807

	MTLBufferLayoutStrideDynamic uint = 9223372036854775807
)

var MTL4CommandQueueErrorDomain foundation.NSErrorDomain


var MTLBinaryArchiveDomain foundation.NSErrorDomain


var MTLCaptureErrorDomain foundation.NSErrorDomain

var MTLCommandBufferEncoderInfoErrorKey foundation.NSErrorUserInfoKey

var MTLCommandBufferErrorDomain foundation.NSErrorDomain

var mTLCommonCounterClipperInvocations MTLCommonCounter

var mTLCommonCounterClipperPrimitivesOut MTLCommonCounter

var mTLCommonCounterComputeKernelInvocations MTLCommonCounter

var mTLCommonCounterFragmentCycles MTLCommonCounter

var mTLCommonCounterFragmentInvocations MTLCommonCounter

var mTLCommonCounterFragmentsPassed MTLCommonCounter

var mTLCommonCounterPostTessellationVertexCycles MTLCommonCounter

var mTLCommonCounterPostTessellationVertexInvocations MTLCommonCounter

var mTLCommonCounterRenderTargetWriteCycles MTLCommonCounter

var mTLCommonCounterSetStageUtilization MTLCommonCounterSet

var mTLCommonCounterSetStatistic MTLCommonCounterSet

var mTLCommonCounterSetTimestamp MTLCommonCounterSet

var mTLCommonCounterTessellationCycles MTLCommonCounter

var mTLCommonCounterTessellationInputPatches MTLCommonCounter

var mTLCommonCounterTimestamp MTLCommonCounter

var mTLCommonCounterTotalCycles MTLCommonCounter

var mTLCommonCounterVertexCycles MTLCommonCounter

var mTLCommonCounterVertexInvocations MTLCommonCounter

var MTLCounterErrorDomain foundation.NSErrorDomain

var MTLDeviceRemovalRequestedNotification MTLDeviceNotificationName

var MTLDeviceWasAddedNotification MTLDeviceNotificationName

var MTLDeviceWasRemovedNotification MTLDeviceNotificationName

var MTLDynamicLibraryDomain foundation.NSErrorDomain

var MTLIOErrorDomain foundation.NSErrorDomain

var MTLLibraryErrorDomain foundation.NSErrorDomain

var MTLLogStateErrorDomain foundation.NSErrorDomain

var MTLTensorDomain foundation.NSErrorDomain

var DeviceCertificationiPhonePerformanceGaming NSDeviceCertification

var ProcessInfoPerformanceProfileDidChangeNotification foundation.NSNotificationName

var processPerformanceProfileDefault NSProcessPerformanceProfile

var processPerformanceProfileSustained NSProcessPerformanceProfile

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "MTL4CommandQueueErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTL4CommandQueueErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "MTLBinaryArchiveDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLBinaryArchiveDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCaptureErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLCaptureErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommandBufferEncoderInfoErrorKey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLCommandBufferEncoderInfoErrorKey = foundation.NSErrorUserInfoKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommandBufferErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLCommandBufferErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterClipperInvocations"); err == nil && ptr != 0 {
		mTLCommonCounterClipperInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterClipperPrimitivesOut"); err == nil && ptr != 0 {
		mTLCommonCounterClipperPrimitivesOut = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterComputeKernelInvocations"); err == nil && ptr != 0 {
		mTLCommonCounterComputeKernelInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterFragmentCycles"); err == nil && ptr != 0 {
		mTLCommonCounterFragmentCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterFragmentInvocations"); err == nil && ptr != 0 {
		mTLCommonCounterFragmentInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterFragmentsPassed"); err == nil && ptr != 0 {
		mTLCommonCounterFragmentsPassed = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterPostTessellationVertexCycles"); err == nil && ptr != 0 {
		mTLCommonCounterPostTessellationVertexCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterPostTessellationVertexInvocations"); err == nil && ptr != 0 {
		mTLCommonCounterPostTessellationVertexInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterRenderTargetWriteCycles"); err == nil && ptr != 0 {
		mTLCommonCounterRenderTargetWriteCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterSetStageUtilization"); err == nil && ptr != 0 {
		mTLCommonCounterSetStageUtilization = *(*MTLCommonCounterSet)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterSetStatistic"); err == nil && ptr != 0 {
		mTLCommonCounterSetStatistic = *(*MTLCommonCounterSet)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterSetTimestamp"); err == nil && ptr != 0 {
		mTLCommonCounterSetTimestamp = *(*MTLCommonCounterSet)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterTessellationCycles"); err == nil && ptr != 0 {
		mTLCommonCounterTessellationCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterTessellationInputPatches"); err == nil && ptr != 0 {
		mTLCommonCounterTessellationInputPatches = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterTimestamp"); err == nil && ptr != 0 {
		mTLCommonCounterTimestamp = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterTotalCycles"); err == nil && ptr != 0 {
		mTLCommonCounterTotalCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterVertexCycles"); err == nil && ptr != 0 {
		mTLCommonCounterVertexCycles = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCommonCounterVertexInvocations"); err == nil && ptr != 0 {
		mTLCommonCounterVertexInvocations = *(*MTLCommonCounter)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLCounterErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLCounterErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLDeviceRemovalRequestedNotification"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLDeviceRemovalRequestedNotification = MTLDeviceNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLDeviceWasAddedNotification"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLDeviceWasAddedNotification = MTLDeviceNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLDeviceWasRemovedNotification"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLDeviceWasRemovedNotification = MTLDeviceNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLDynamicLibraryDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLDynamicLibraryDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLIOErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLIOErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLLibraryErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLLibraryErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLLogStateErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTLLogStateErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTLTensorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
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
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				ProcessInfoPerformanceProfileDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProcessPerformanceProfileDefault"); err == nil && ptr != 0 {
		processPerformanceProfileDefault = *(*NSProcessPerformanceProfile)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "NSProcessPerformanceProfileSustained"); err == nil && ptr != 0 {
		processPerformanceProfileSustained = *(*NSProcessPerformanceProfile)(unsafe.Pointer(ptr))
	}

}

type MTLCommonCounterSetValues struct{}

// MTLCommonCounterSets provides typed accessors for [MTLCommonCounterSet] constants.
var MTLCommonCounterSets MTLCommonCounterSetValues

// StageUtilization returns The common name for the counter set that contains hardware utilization measurements from various render stages.
func (MTLCommonCounterSetValues) StageUtilization() MTLCommonCounterSet { return mTLCommonCounterSetStageUtilization }

// Statistic returns The common name for the counter set that contains GPU workload statistics.
func (MTLCommonCounterSetValues) Statistic() MTLCommonCounterSet { return mTLCommonCounterSetStatistic }

// Timestamp returns The common name for the counter set that contains the timestamp counter.
func (MTLCommonCounterSetValues) Timestamp() MTLCommonCounterSet { return mTLCommonCounterSetTimestamp }


type MTLCommonCounterValues struct{}

// MTLCommonCounters provides typed accessors for [MTLCommonCounter] constants.
var MTLCommonCounters MTLCommonCounterValues

// ClipperInvocations returns The common name for the counter that tracks the number of primitives a render pass sends to the clip stage.
func (MTLCommonCounterValues) ClipperInvocations() MTLCommonCounter { return mTLCommonCounterClipperInvocations }

// ClipperPrimitivesOut returns The common name for the counter that tracks the number of primitives the clip stage produces during a render pass.
func (MTLCommonCounterValues) ClipperPrimitivesOut() MTLCommonCounter { return mTLCommonCounterClipperPrimitivesOut }

// ComputeKernelInvocations returns The common name for the counter that tracks the number of times a pass invokes any compute kernel.
func (MTLCommonCounterValues) ComputeKernelInvocations() MTLCommonCounter { return mTLCommonCounterComputeKernelInvocations }

// FragmentCycles returns The common name for the counter that tracks the number of cycles the GPU uses to run fragment shaders during a pass.
func (MTLCommonCounterValues) FragmentCycles() MTLCommonCounter { return mTLCommonCounterFragmentCycles }

// FragmentInvocations returns The common name for the counter that tracks the number of times a render pass calls fragment shaders.
func (MTLCommonCounterValues) FragmentInvocations() MTLCommonCounter { return mTLCommonCounterFragmentInvocations }

// FragmentsPassed returns The common name for the counter that tracks the number of fragments a render pass sends to the visibility and blend stages.
func (MTLCommonCounterValues) FragmentsPassed() MTLCommonCounter { return mTLCommonCounterFragmentsPassed }

// PostTessellationVertexCycles returns The common name for the counter that tracks the number of cycles the GPU uses to run post-tessellation vertex shaders during a pass.
func (MTLCommonCounterValues) PostTessellationVertexCycles() MTLCommonCounter { return mTLCommonCounterPostTessellationVertexCycles }

// PostTessellationVertexInvocations returns The common name for the counter that tracks the number of vertices a render pass sends to a post-tessellation vertex shader.
func (MTLCommonCounterValues) PostTessellationVertexInvocations() MTLCommonCounter { return mTLCommonCounterPostTessellationVertexInvocations }

// RenderTargetWriteCycles returns The common name for the counter that tracks the number of cycles the GPU uses to write data to render targets during a render pass.
func (MTLCommonCounterValues) RenderTargetWriteCycles() MTLCommonCounter { return mTLCommonCounterRenderTargetWriteCycles }

// TessellationCycles returns The common name for the counter that tracks the number of cycles the GPU uses to run the tessellation stage during a pass.
func (MTLCommonCounterValues) TessellationCycles() MTLCommonCounter { return mTLCommonCounterTessellationCycles }

// TessellationInputPatches returns The common name for the counter that tracks the number of tessellation patches a render pass sends to the tessellation stage.
func (MTLCommonCounterValues) TessellationInputPatches() MTLCommonCounter { return mTLCommonCounterTessellationInputPatches }

// Timestamp returns The common name for the counter that tracks the current time.
func (MTLCommonCounterValues) Timestamp() MTLCommonCounter { return mTLCommonCounterTimestamp }

// TotalCycles returns The common name for the counter that tracks the total number of cycles the GPU uses to run a pass.
func (MTLCommonCounterValues) TotalCycles() MTLCommonCounter { return mTLCommonCounterTotalCycles }

// VertexCycles returns The common name for the counter that tracks the number of cycles the GPU uses to run vertex shaders during a pass.
func (MTLCommonCounterValues) VertexCycles() MTLCommonCounter { return mTLCommonCounterVertexCycles }

// VertexInvocations returns The common name for the counter that tracks the number of times a render pass calls any vertex shader.
func (MTLCommonCounterValues) VertexInvocations() MTLCommonCounter { return mTLCommonCounterVertexInvocations }


type NSProcessPerformanceProfileValues struct{}

// NSProcessPerformanceProfiles provides typed accessors for [NSProcessPerformanceProfile] constants.
var NSProcessPerformanceProfiles NSProcessPerformanceProfileValues

// Default returns The default performance profile for a device.
func (NSProcessPerformanceProfileValues) Default() NSProcessPerformanceProfile { return processPerformanceProfileDefault }

// Sustained returns The performance profile for a device representing sustained performance.
func (NSProcessPerformanceProfileValues) Sustained() NSProcessPerformanceProfile { return processPerformanceProfileSustained }


