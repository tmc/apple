// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/foundation"
"github.com/tmc/apple/objectivec"
)

// Updates the dynamic aspect ratio of the device.
//
// dynamicAspectRatio: The new [AVCaptureDevice.AspectRatio] the device should output.
// //
// [AVCaptureDevice.AspectRatio]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/AspectRatio
//
// handler: A block called by the device when `dynamicAspectRatio` is set to the value
// specified. If you call [SetDynamicAspectRatioCompletionHandler] multiple
// times, the completion handlers are called in FIFO order. The block receives
// a timestamp which matches that of the first buffer to which all settings
// have been applied. Note that the timestamp is synchronized to the device
// clock, and thus must be converted to the [SynchronizationClock] prior to
// comparison with the timestamps of buffers delivered via an
// [AVCaptureVideoDataOutput]. You may pass `nil` for the `handler` parameter
// if you do not need to know when the operation completes.
//
// # Discussion
// 
// This is the only way of setting [DynamicAspectRatio]. This method throws an
// [NSInvalidArgumentException] if `dynamicAspectRatio` is not a supported
// aspect ratio found in the device’s activeFormat’s
// [SupportedDynamicAspectRatios]. This method throws an [NSGenericException]
// if you call it without first obtaining exclusive access to the device using
// [LockForConfiguration].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setDynamicAspectRatio(_:completionHandler:)
func (c AVCaptureDevice) SetDynamicAspectRatioCompletionHandler(dynamicAspectRatio objectivec.IObject, handler CMTimeErrorHandler) {
	_block1, _cleanup1 := NewCMTimeErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](c.ID, objc.Sel("setDynamicAspectRatio:completionHandler:"), dynamicAspectRatio, _block1)
}
// Smoothly ends a zoom transition in progress.
//
// # Discussion
// 
// Calling this method is equivalent to calling
// [RampToVideoZoomFactorWithRate] with a rate of zero. If a zoom transition
// is in progress, the transition slows to a stop (instead of stopping
// abruptly).
// 
// Before calling this method, you must call [LockForConfiguration] to acquire
// exclusive access to the device’s configuration properties. If you
// don’t, calling this method raises an exception. When you finish
// configuring the device, call [UnlockForConfiguration] to release the lock
// and allow other devices to configure the settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cancelVideoZoomRamp()
func (c AVCaptureDevice) CancelVideoZoomRamp() {
objc.Send[objc.ID](c.ID, objc.Sel("cancelVideoZoomRamp"))
}
// Converts device-specific white balance RGB gain values to
// device-independent chromaticity values.
//
// whiteBalanceGains: The white balance gain values. You can’t specify a value of
// [currentWhiteBalanceGains].
// //
// [currentWhiteBalanceGains]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/currentWhiteBalanceGains
//
// # Return Value
// 
// A structure that contains device-independent values.
//
// # Discussion
// 
// Call this method to convert device-specific white balance RGB gain values
// to device-independent chromaticity (little x, little y) values.
// 
// Each change in the structure supports values between `1.0` and
// [MaxWhiteBalanceGain]. This method throws an exception if you specify an
// unsupported value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/chromaticityValues(for:)
func (c AVCaptureDevice) ChromaticityValuesForDeviceWhiteBalanceGains(whiteBalanceGains AVCaptureWhiteBalanceGains) AVCaptureWhiteBalanceChromaticityValues {
rv := objc.Send[AVCaptureWhiteBalanceChromaticityValues](c.ID, objc.Sel("chromaticityValuesForDeviceWhiteBalanceGains:"), whiteBalanceGains)
return AVCaptureWhiteBalanceChromaticityValues(rv)
}
// Converts device-independent chromaticity values to device-specific white
// balance RGB gain values.
//
// chromaticityValues: The chromaticity values for which to get white balance RGB gain values.
//
// # Return Value
// 
// A structure that contains device-specific RGB gain values.
//
// # Discussion
// 
// This property specifies the current red, green, and blue gain values used
// for white balance. You can use the values to adjust color casts for a given
// scene.
// 
// Each channel supports values between `1.0` and -[MaxWhiteBalanceGain].
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/deviceWhiteBalanceGains(for:)-9gdtw
func (c AVCaptureDevice) DeviceWhiteBalanceGainsForChromaticityValues(chromaticityValues AVCaptureWhiteBalanceChromaticityValues) AVCaptureWhiteBalanceGains {
rv := objc.Send[AVCaptureWhiteBalanceGains](c.ID, objc.Sel("deviceWhiteBalanceGainsForChromaticityValues:"), chromaticityValues)
return AVCaptureWhiteBalanceGains(rv)
}
// Converts device-independent temperature and tint values to device-specific
// white balance RGB gain values.
//
// tempAndTintValues: An [AVCaptureDevice.WhiteBalanceTemperatureAndTintValues] structure
// containing the temperature and tint values.
// //
// [AVCaptureDevice.WhiteBalanceTemperatureAndTintValues]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceTemperatureAndTintValues
//
// # Return Value
// 
// A fully populated [AVCaptureDevice.WhiteBalanceGains] structure containing
// device-specific RGB gain values.
//
// [AVCaptureDevice.WhiteBalanceGains]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceGains
//
// # Discussion
// 
// Call this method to convert device-independent temperature and tint values
// to device-specific RGB white balance gain values.
// 
// You may pass any temperature and tint values and corresponding white
// balance gains will be produced. Note, though, that some temperature and
// tint combinations yield out-of-range device RGB values that will cause an
// exception to be thrown if passed directly to
// [SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler]. Be
// sure to verify that the red, green, and blue gain values are within the
// range of [MaxWhiteBalanceGain]].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/deviceWhiteBalanceGains(for:)-3wtsa
func (c AVCaptureDevice) DeviceWhiteBalanceGainsForTemperatureAndTintValues(tempAndTintValues AVCaptureWhiteBalanceTemperatureAndTintValues) AVCaptureWhiteBalanceGains {
rv := objc.Send[AVCaptureWhiteBalanceGains](c.ID, objc.Sel("deviceWhiteBalanceGainsForTemperatureAndTintValues:"), tempAndTintValues)
return AVCaptureWhiteBalanceGains(rv)
}
// Begins a smooth transition from the current zoom factor to another.
//
// factor: The new magnification factor.
//
// rate: The rate at which to transition to the new magnification factor, specified
// in powers of two per second.
//
// # Discussion
// 
// Allowed values for `factor` range from `1.0` (full field of view) to the
// [VideoMaxZoomFactor] specified by the active capture format.
// 
// During a ramp, the zoom factor changes at an exponential rate, but this
// yields a visually linear transition. The `rate` parameter controls the
// speed of this transition independent of direction; for example, a value of
// `1.0` causes zoom factor to double every second if zooming in (that’s, if
// the specified `factor` is greater than the current [VideoZoomFactor]) or
// halve every second if zooming out.
// 
// Before calling this method, you must call [LockForConfiguration] to acquire
// exclusive access to the device’s configuration properties. If you
// don’t, calling this method raises an exception. When you finish
// configuring the device, call [UnlockForConfiguration] to release the lock
// and allow other devices to configure the settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/ramp(toVideoZoomFactor:withRate:)
func (c AVCaptureDevice) RampToVideoZoomFactorWithRate(factor float64, rate float32) {
objc.Send[objc.ID](c.ID, objc.Sel("rampToVideoZoomFactor:withRate:"), factor, rate)
}
// Sets the exposure mode to a custom state, and locks exposure duration and
// ISO at explicit values.
//
// duration: The exposure duration.
// 
// Pass a value of [currentExposureDuration] to leave the current exposure
// duration unchanged.
// 
// Changes made to the exposure duration may result in changes to
// [ActiveVideoMinFrameDuration] or [ActiveVideoMaxFrameDuration].
// //
// [currentExposureDuration]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/currentExposureDuration
//
// ISO: The exposure ISO value.
// 
// Pass a value of [currentISO] to leave the current ISO unchanged.
// //
// [currentISO]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/currentISO
//
// handler: A callback the system invokes when the adjustment to the exposure duration
// and ISO is complete. If you call this method multiple times, the system
// calls the completion handlers in FIFO order.
// 
// The system passes a time value that matches that of the first buffer to
// which its applied all settings. It synchronizes the timestamp to the device
// clock, and you must convert the timestamp to the [SynchronizationClock]
// prior to comparison with the timestamps of buffers delivered through an
// [AVCaptureVideoDataOutput].
// 
// You can pass `nil` for this parameter if you don’t require this
// information.
//
// duration is a [coremedia.CMTime].
//
// # Discussion
// 
// This method throws an exception if you set the exposure duration or ISO
// values to an unsupported level
// 
// Before changing exposure mode, you must call [LockForConfiguration] to
// acquire exclusive access to the device’s configuration properties.
// Otherwise, setting the value of this property raises an exception. When you
// finish configuring the device, call [UnlockForConfiguration] to release the
// lock and allow other devices to configure the settings.
// 
// When using [AVCapturePhotoOutput] to capture photos, the
// [PhotoQualityPrioritization] property of [AVCapturePhotoSettings] defaults
// to [CapturePhotoQualityPrioritizationBalanced], which allows photo capture
// to temporarily override the capture device’s exposure duration and ISO if
// the scene is dark enough to require multi-image fusion to improve quality.
// To ensure that the system honors the device exposure duration and ISO
// values while in [CaptureExposureModeCustom] or [CaptureExposureModeLocked]
// mode, you must photo quality prioritization to
// [CapturePhotoQualityPrioritizationSpeed].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setExposureModeCustom(duration:iso:completionHandler:)
func (c AVCaptureDevice) SetExposureModeCustomWithDurationISOCompletionHandler(duration objectivec.IObject, ISO float32, handler CMTimeHandler) {
	_block2, _cleanup2 := NewCMTimeBlock(handler)
	defer _cleanup2()
	objc.Send[objc.ID](c.ID, objc.Sel("setExposureModeCustomWithDuration:ISO:completionHandler:"), duration, ISO, _block2)
}
// Sets the bias to apply to the target exposure value.
//
// bias: The bias to apply to the exposure target value.
//
// handler: A callback the system invokes when the adjustment to the exposure target
// bias is complete. If you call this method multiple times, the system calls
// the completion handlers in FIFO order.
// 
// The system passes a time value that matches that of the first buffer to
// which its applied all settings. It synchronizes the timestamp to the device
// clock, and you must convert the timestamp to the [SynchronizationClock]
// prior to comparison with the timestamps of buffers delivered through an
// [AVCaptureVideoDataOutput].
// 
// You can pass `nil` for this parameter if you don’t require this
// information.
//
// # Discussion
// 
// Before changing the value the lens position, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setExposureTargetBias(_:completionHandler:)
func (c AVCaptureDevice) SetExposureTargetBiasCompletionHandler(bias float32, handler CMTimeHandler) {
	_block1, _cleanup1 := NewCMTimeBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](c.ID, objc.Sel("setExposureTargetBias:completionHandler:"), bias, _block1)
}
// Locks the lens position at the specified value, and sets the focus mode to
// a locked state.
//
// lensPosition: The lens position. Pass a value of [currentLensPosition] to leave the
// current lens position unchanged.
// //
// [currentLensPosition]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/currentLensPosition
//
// handler: A callback the system invokes when the adjustment to the lens position is
// complete and the [FocusMode] set to a locked state. If you call this method
// multiple times, the system calls the completion handlers in FIFO order.
// 
// The system passes a time value that matches that of the first buffer to
// which its applied all settings. It synchronizes the timestamp to the device
// clock, and you must convert the timestamp to the [SynchronizationClock]
// prior to comparison with the timestamps of buffers delivered through an
// [AVCaptureVideoDataOutput].
// 
// You can pass `nil` for this parameter if you don’t require this
// information.
//
// # Discussion
// 
// Calling this method is the only way to set the value of the [LensPosition]
// property. This method throws an exception if you set the value to an
// unsupported level.
// 
// Before changing the value the lens position, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setFocusModeLocked(lensPosition:completionHandler:)
func (c AVCaptureDevice) SetFocusModeLockedWithLensPositionCompletionHandler(lensPosition float32, handler CMTimeHandler) {
	_block1, _cleanup1 := NewCMTimeBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](c.ID, objc.Sel("setFocusModeLockedWithLensPosition:completionHandler:"), lensPosition, _block1)
}
// Sets the white balance to locked mode with the specified white balance
// gains.
//
// whiteBalanceGains: The white balance gains to set. Pass a value of [currentWhiteBalanceGains]
// to leave the current white balance unchanged.
// //
// [currentWhiteBalanceGains]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/currentWhiteBalanceGains
//
// handler: A callback the system invokes when the adjustment to the white balance is
// complete and the [WhiteBalanceMode] set to a locked state. If you call this
// method multiple times, the system calls the completion handlers in FIFO
// order.
// 
// The system passes a time value that matches that of the first buffer to
// which its applied all settings. It synchronizes the timestamp to the device
// clock, and you must convert the timestamp to the [SynchronizationClock]
// prior to comparison with the timestamps of buffers delivered through an
// [AVCaptureVideoDataOutput].
// 
// You can pass `nil` for this parameter if you don’t require this
// information.
//
// # Discussion
// 
// Each channel in the white balance gains structure supports values between
// `1.0` and [MaxWhiteBalanceGain]. Setting a channel value outside this range
// generates an exception.
// 
// The system normalizes gain values to the minimum channel value to avoid
// brightness changes (for example, `R:2 G:2 B:4` normalizes to `R:1 G:1
// B:2`).
// 
// Before changing the value the white balance gains, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setWhiteBalanceModeLocked(with:completionHandler:)
func (c AVCaptureDevice) SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler(whiteBalanceGains AVCaptureWhiteBalanceGains, handler CMTimeHandler) {
	_block1, _cleanup1 := NewCMTimeBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](c.ID, objc.Sel("setWhiteBalanceModeLockedWithDeviceWhiteBalanceGains:completionHandler:"), whiteBalanceGains, _block1)
}
// Sets white balance to locked mode with explicit temperature and tint
// values.
//
// whiteBalanceTemperatureAndTintValues: The white balance temperature and tint values, as computed from
// [TemperatureAndTintValuesForDeviceWhiteBalanceGains] method,
// [AVCaptureDevice.WhiteBalanceTemperatureAndTintValues] presets or manual
// input.
// //
// [AVCaptureDevice.WhiteBalanceTemperatureAndTintValues]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceTemperatureAndTintValues
//
// handler: A block to be called when white balance values have been set to the values
// specified and [WhiteBalanceMode] is set to
// [AVCaptureWhiteBalanceModeLocked]. If
// [SetWhiteBalanceModeLockedWithDeviceWhiteBalanceTemperatureAndTintValuesCompletionHandler]
// is called multiple times, the completion handlers are called in FIFO order.
// The block receives a timestamp which matches that of the first buffer to
// which all settings have been applied. Note that the timestamp is
// synchronized to the device clock, and thus must be converted to the
// [SynchronizationClock] prior to comparison with the timestamps of buffers
// delivered via an [AVCaptureVideoDataOutput]. This parameter may be `nil` if
// synchronization is not required.
//
// # Discussion
// 
// This method takes a [AVCaptureDevice.WhiteBalanceTemperatureAndTintValues]
// struct and applies the appropriate [AVCaptureDevice.WhiteBalanceGains].
// This method throws an [NSRangeException] if any of the values are set to an
// unsupported level. This method throws an [NSGenericException] if called
// without first obtaining exclusive access to the device using
// [LockForConfiguration].
//
// [AVCaptureDevice.WhiteBalanceGains]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceGains
// [AVCaptureDevice.WhiteBalanceTemperatureAndTintValues]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceTemperatureAndTintValues
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setWhiteBalanceModeLocked(whiteBalanceTemperatureAndTintValues:handler:)
func (c AVCaptureDevice) SetWhiteBalanceModeLockedWithDeviceWhiteBalanceTemperatureAndTintValuesCompletionHandler(whiteBalanceTemperatureAndTintValues AVCaptureWhiteBalanceTemperatureAndTintValues, handler CMTimeHandler) {
	_block1, _cleanup1 := NewCMTimeBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](c.ID, objc.Sel("setWhiteBalanceModeLockedWithDeviceWhiteBalanceTemperatureAndTintValues:completionHandler:"), whiteBalanceTemperatureAndTintValues, _block1)
}
// Converts device-specific white balance RGB gain values to
// device-independent temperature and tint values.
//
// whiteBalanceGains: The white balance gain values. You can’t specify a value of
// [currentWhiteBalanceGains].
// //
// [currentWhiteBalanceGains]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/currentWhiteBalanceGains
//
// # Return Value
// 
// A structure that contains device-independent values.
//
// # Discussion
// 
// Each change in the structure supports values between `1.0` and
// [MaxWhiteBalanceGain]. This method throws an exception if you specify an
// unsupported value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/temperatureAndTintValues(for:)
func (c AVCaptureDevice) TemperatureAndTintValuesForDeviceWhiteBalanceGains(whiteBalanceGains AVCaptureWhiteBalanceGains) AVCaptureWhiteBalanceTemperatureAndTintValues {
rv := objc.Send[AVCaptureWhiteBalanceTemperatureAndTintValues](c.ID, objc.Sel("temperatureAndTintValuesForDeviceWhiteBalanceGains:"), whiteBalanceGains)
return AVCaptureWhiteBalanceTemperatureAndTintValues(rv)
}

// Returns the relative extrinsic matrix from one capture device to another.
//
// fromDevice: The capture device that represents the source camera.
//
// toDevice: The capture device that represents the destination camera.
//
// # Return Value
// 
// A [Data] containing a [matrix_float4x3] matrix, which is a column major
// with 3 rows and 4 columns.
//
// [Data]: https://developer.apple.com/documentation/Foundation/Data
// [matrix_float4x3]: https://developer.apple.com/documentation/simd/matrix_float4x3
//
// # Discussion
// 
// The extrinsic matrix consists of a unitless 3x3 rotation matrix (R) on the
// left and a translation (t) 3x1 column vector on the right, whose units are
// millimeters. The matrix expresses the destination camera’s extrinsics
// relative to the source camera. If `X_from` is a 3D point in the source
// camera’s coordinate system, you project it into the destination
// camera’s coordinate system with `X_to = [R | t] * X_from`.
// 
// Only physical cameras for which factory calibrations exist provide an
// extrinsic matrix. Virtual device cameras return `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/extrinsicMatrix(from:to:)
func (_AVCaptureDeviceClass AVCaptureDeviceClass) ExtrinsicMatrixFromDeviceToDevice(fromDevice IAVCaptureDevice, toDevice IAVCaptureDevice) foundation.NSData {
rv := objc.Send[objc.ID](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("extrinsicMatrixFromDevice:toDevice:"), fromDevice, toDevice)
return foundation.NSDataFromID(rv)
}

// A Boolean value that indicates whether the device consists of two or more
// physical devices.
//
// # Discussion
// 
// Examples of virtual devices are:
// 
// - The dual camera, which supports seamless switching between wide-angle and
// telephoto cameras while zooming and generating depth data from the
// disparities between the points of view of the physical cameras. - The
// TrueDepth camera, which generates depth data from disparities between YUV
// and infrared cameras pointed in the same direction.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVirtualDevice
func (c AVCaptureDevice) VirtualDevice() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isVirtualDevice"))
		return rv
}
// An array of physical devices that make up a virtual device.
//
// # Discussion
// 
// The value of this property is an empty array when called on a device whose
// [VirtualDevice] property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/constituentDevices
func (c AVCaptureDevice) ConstituentDevices() []AVCaptureDevice {
rv := objc.Send[[]objc.ID](c.ID, objc.Sel("constituentDevices"))
return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureDevice {
	return AVCaptureDeviceFromID(id)
})
}
// A Boolean value that indicates whether the device monitors the subject area
// for changes.
//
// # Discussion
// 
// The value of this property indicates whether the device monitors the video
// subject area for changes, such as lighting changes, substantial movement,
// and so on. If you enable subject area change monitoring, the capture device
// object sends an [subjectAreaDidChangeNotification] whenever it detects a
// change to the subject area. You can observe this notification and take
// action such as focusing, adjusting exposure, and so on.
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you’re finished configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// This property is key-value observable.
//
// [subjectAreaDidChangeNotification]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/subjectAreaDidChangeNotification
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSubjectAreaChangeMonitoringEnabled
func (c AVCaptureDevice) SubjectAreaChangeMonitoringEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isSubjectAreaChangeMonitoringEnabled"))
		return rv
}
func (c AVCaptureDevice) SetSubjectAreaChangeMonitoringEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setSubjectAreaChangeMonitoringEnabled:"), value)
}
// A monitor owned by the device that recommends an optimal framing based on
// the content in the scene.
//
// # Discussion
// 
// An ultra wide camera device that supports dynamic aspect ratio
// configuration may also support “smart framing monitoring”. If this
// property returns non `nil`, you may use it to listen for framing
// recommendations by configuring its [enabledFramings] and calling
// [startMonitoring()]. The smart framing monitor only makes recommendations
// when the current [ActiveFormat] supports smart framing (see
// [SmartFramingSupported]).
//
// [enabledFramings]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor/enabledFramings
// [startMonitoring()]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor/startMonitoring()
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/smartFramingMonitor
func (c AVCaptureDevice) SmartFramingMonitor() objc.ID {
rv := objc.Send[objc.ID](c.ID, objc.Sel("smartFramingMonitor"))
		return rv
}
// A key-value observable property indicating the current aspect ratio for a
// device.
//
// # Discussion
// 
// This property is initialized to the first [AVCaptureDevice.AspectRatio]
// listed in the device’s activeFormat’s [SupportedDynamicAspectRatios]
// property. If the activeFormat’s [SupportedDynamicAspectRatios] is an
// empty array, this property returns nil.
//
// [AVCaptureDevice.AspectRatio]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/AspectRatio
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dynamicAspectRatio
func (c AVCaptureDevice) DynamicAspectRatio() objc.ID {
rv := objc.Send[objc.ID](c.ID, objc.Sel("dynamicAspectRatio"))
		return rv
}
// A key-value observable property describing the output dimensions of the
// video buffer based on the device’s dynamic aspect ratio.
//
// # Discussion
// 
// If the device’s activeFormat’s [SupportedDynamicAspectRatios] is an
// empty array, this property returns {0,0}.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dynamicDimensions
func (c AVCaptureDevice) DynamicDimensions() objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("dynamicDimensions"))
return objectivec.Object{ID: rv}
}
// A value that indicates the capture device’s current system pressure
// state.
//
// # Discussion
// 
// This property indicates whether the capture device is currently in an
// elevated system pressure condition. When system pressure reaches a
// [shutdown] state, the capture device can’t continue to provide input, and
// the capture session becomes interrupted until the pressured state abates.
// 
// You can effectively mitigate system pressure by lowering the device’s
// [ActiveVideoMinFrameDuration] in response to changes in the system pressure
// state. Implement frame rate throttling to bring system pressure down if
// your capture use case can tolerate a reduced frame rate.
//
// [shutdown]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class/Level-swift.struct/shutdown
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPressureState-swift.property
func (c AVCaptureDevice) SystemPressureState() objc.ID {
rv := objc.Send[objc.ID](c.ID, objc.Sel("systemPressureState"))
		return rv
}
// The nominal 35mm equivalent focal length of the capture device’s lens.
//
// # Discussion
// 
// This value represents a nominal measurement of the device’s field of
// view, expressed as a 35mm equivalent focal length, measured diagonally. The
// value is similar to the [FocalLengthIn35mmFormat] EXIF entry (see
// [kCGImagePropertyExifFocalLenIn35mmFilm]) for a photo captured using the
// device’s format where [HighestPhotoQualitySupported] is `true` or when
// you’ve configured the session with the [photo] preset.
// 
// This property value is `0` for virtual devices and external cameras.
//
// [kCGImagePropertyExifFocalLenIn35mmFilm]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifFocalLenIn35mmFilm
// [photo]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/photo
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/nominalFocalLengthIn35mmFilm
func (c AVCaptureDevice) NominalFocalLengthIn35mmFilm() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("nominalFocalLengthIn35mmFilm"))
		return rv
}
// The current exposure ISO value.
//
// # Discussion
// 
// This property controls the sensor’s sensitivity to light by applying a
// gain value to the signal. This value is between the active format’s
// [MinISO] and [MaxISO] values. Higher values result in noisier images.
// 
// To set the ISO, call the
// [SetExposureModeCustomWithDurationISOCompletionHandler] method.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/iso
func (c AVCaptureDevice) ISO() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("ISO"))
		return rv
}
// The currently active depth data format of the capture device.
//
// # Discussion
// 
// You must obtain exclusive access to the device by calling
// [LockForConfiguration] before setting this property value.
// 
// You can set this property only to formats present in the active format’s
// [SupportedDepthDataFormats] array. Attempting to set an unsupported format
// throws an exception.
// 
// You can’t set the frame rate of depth data directly. Instead, the system
// synchronizes the depth data frame rate to the device’s
// [ActiveVideoMinFrameDuration] and [ActiveVideoMaxFrameDuration] values. It
// may match the device’s current frame rate, or lower, if the system
// can’t produce depth data fast enough for the active video frame rate.
// 
// Delivery of depth data to a [AVCaptureDepthDataOutput] may increase the
// system load, resulting in a reduced video frame rate for thermal
// sustainability.
// 
// On devices where depth data isn’t supported, this property value is
// `nil`.
// 
// This property is key-value observable.
//
// [AVCaptureDepthDataOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeDepthDataFormat
func (c AVCaptureDevice) ActiveDepthDataFormat() IAVCaptureDeviceFormat {
rv := objc.Send[objc.ID](c.ID, objc.Sel("activeDepthDataFormat"))
return AVCaptureDeviceFormatFromID(rv)
}
func (c AVCaptureDevice) SetActiveDepthDataFormat(value IAVCaptureDeviceFormat) {
objc.Send[struct{}](c.ID, objc.Sel("setActiveDepthDataFormat:"), value)
}
// The minimum frame duration of depth data.
//
// # Discussion
// 
// Use this property to set an upper limit to the frame rate at which the
// system produces depth data. Lowering the depth data frame rate typically
// lowers power consumption which increases the time the camera can run before
// it reaches an elevated system pressure state. Setting a value outside the
// active depth data format’s supported frame rate range produces an
// exception.
// 
// Setting this property to [invalid] resets it to the active depth data
// format’s default minimum frame duration. Setting this property to
// [positiveInfinity] results in a depth data frame rate of `0`.
// 
// This value gets reset whenever either the active video format or the active
// depth data format changes.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
// [positiveInfinity]: https://developer.apple.com/documentation/CoreMedia/CMTime/positiveInfinity
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeDepthDataMinFrameDuration
func (c AVCaptureDevice) ActiveDepthDataMinFrameDuration() objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("activeDepthDataMinFrameDuration"))
return objectivec.Object{ID: rv}
}
func (c AVCaptureDevice) SetActiveDepthDataMinFrameDuration(value objectivec.IObject) {
objc.Send[struct{}](c.ID, objc.Sel("setActiveDepthDataMinFrameDuration:"), value)
}
// The maximum exposure duration, in seconds, defined in the autoexposure
// algorithm.
//
// # Discussion
// 
// When you set the exposureMode to [CaptureExposureModeAutoExpose] or
// [CaptureExposureModeContinuousAutoExposure], the autoexposure algorithm
// picks a default maximum exposure duration that’s tuned for the current
// configuration, balancing low light image quality with motion preservation.
// By querying or key-value observing this property, you can determine the
// current maximum exposure duration in use.
// 
// You may also override the default value by setting this property to a value
// between the format’s [MinExposureDuration] and [MaxExposureDuration]
// values. The system throws an exception if you pass an out-of-bounds
// exposure value.
// 
// Setting the property to the special value of [invalid] resets the
// autoexposure maximum duration to the device’s default for your current
// configuration. When the device’s [ActiveFormat] or the capture
// session’s [SessionPreset] changes, this property resets to the default
// max exposure duration for the new format or session preset.
// 
// On some devices, the auto exposure algorithm picks a different maximum
// exposure duration for a given format depending on whether you used the
// [SessionPreset] or [ActiveFormat] APIs to set to set the format. To ensure
// uniform default handling of maximum exposure duration, set the value of a
// capture input’s [UnifiedAutoExposureDefaultsEnabled] property to [true].
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeMaxExposureDuration
func (c AVCaptureDevice) ActiveMaxExposureDuration() objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("activeMaxExposureDuration"))
return objectivec.Object{ID: rv}
}
func (c AVCaptureDevice) SetActiveMaxExposureDuration(value objectivec.IObject) {
objc.Send[struct{}](c.ID, objc.Sel("setActiveMaxExposureDuration:"), value)
}
// A value that controls the allowable range for automatic focusing.
//
// # Discussion
// 
// By default, a device capable of hardware focusing attempts to focus on
// objects at any distance. If you expect to focus primarily on near or far
// objects, set a range restriction to increase the speed and reduce the power
// consumption of automatic focusing, and to reduce the chance of focusing
// ambiguities.
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/autoFocusRangeRestriction-swift.property
func (c AVCaptureDevice) AutoFocusRangeRestriction() AVCaptureAutoFocusRangeRestriction {
rv := objc.Send[AVCaptureAutoFocusRangeRestriction](c.ID, objc.Sel("autoFocusRangeRestriction"))
		return AVCaptureAutoFocusRangeRestriction(rv)
}
func (c AVCaptureDevice) SetAutoFocusRangeRestriction(value AVCaptureAutoFocusRangeRestriction) {
objc.Send[struct{}](c.ID, objc.Sel("setAutoFocusRangeRestriction:"), value)
}
// A Boolean value that indicates whether the device supports focus range
// restrictions.
//
// # Discussion
// 
// Focus range restriction is available only on compatible devices. If this
// property’s value is [false], setting the value of
// [AutoFocusRangeRestriction] raises an exception.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAutoFocusRangeRestrictionSupported
func (c AVCaptureDevice) AutoFocusRangeRestrictionSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isAutoFocusRangeRestrictionSupported"))
		return rv
}
// A Boolean value that indicates whether the device automatically adjusts
// face-driven autoexposure.
//
// # Discussion
// 
// The value of this property defaults to [true] for devices that support
// autoexposure. If your app requires explicitly setting the state of
// [FaceDrivenAutoExposureEnabled], set this value to [false].
// 
// To set this property value, you must call the device’s
// [LockForConfiguration] method to obtain exclusive access to configure it.
// Otherwise, attempting to set a value raises an exception. When you finish
// configuring the device, call [UnlockForConfiguration] to release the lock.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/automaticallyAdjustsFaceDrivenAutoExposureEnabled
func (c AVCaptureDevice) AutomaticallyAdjustsFaceDrivenAutoExposureEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("automaticallyAdjustsFaceDrivenAutoExposureEnabled"))
		return rv
}
func (c AVCaptureDevice) SetAutomaticallyAdjustsFaceDrivenAutoExposureEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyAdjustsFaceDrivenAutoExposureEnabled:"), value)
}
// A Boolean value that indicates whether the device automatically adjusts
// face-driven autofocus.
//
// # Discussion
// 
// The value of this property defaults to [true] for devices that support auto
// focus. If your app requires explicitly setting the state of
// [FaceDrivenAutoFocusEnabled], set this value to [false].
// 
// To set this property value, you must call the device’s
// [LockForConfiguration] method to obtain exclusive access to configure it.
// Otherwise, attempting to set a value raises an exception. When you finish
// configuring the device, call [UnlockForConfiguration] to release the lock.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/automaticallyAdjustsFaceDrivenAutoFocusEnabled
func (c AVCaptureDevice) AutomaticallyAdjustsFaceDrivenAutoFocusEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("automaticallyAdjustsFaceDrivenAutoFocusEnabled"))
		return rv
}
func (c AVCaptureDevice) SetAutomaticallyAdjustsFaceDrivenAutoFocusEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyAdjustsFaceDrivenAutoFocusEnabled:"), value)
}
// A Boolean value that indicates whether the device automatically manages the
// state of high dynamic range (HDR) video streaming.
//
// # Discussion
// 
// By default, this value is `true`, and a capture device automatically
// enables [VideoHDREnabled] if it’s a good fit for the active format.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/automaticallyAdjustsVideoHDREnabled
func (c AVCaptureDevice) AutomaticallyAdjustsVideoHDREnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("automaticallyAdjustsVideoHDREnabled"))
		return rv
}
func (c AVCaptureDevice) SetAutomaticallyAdjustsVideoHDREnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyAdjustsVideoHDREnabled:"), value)
}
// A Boolean value that indicates whether the capture device automatically
// switches to low-light boost mode when necessary.
//
// # Discussion
// 
// The default value is [false]. When it’s [true], the device may engage a
// special low-light boost mode to improve image quality. It switches, at its
// discretion, to a special boost mode under low light, and back to normal
// operation when the scene becomes sufficiently lit.
// 
// Setting a value for this property throws an exception if the value of
// [LowLightBoostSupported] is false.
// 
// A capture device that supports this feature may only engage boost mode for
// certain source formats or resolutions. The switch between normal operation
// and low light boost mode may drop one or more video frames.
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// This property is key-value observable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/automaticallyEnablesLowLightBoostWhenAvailable
func (c AVCaptureDevice) AutomaticallyEnablesLowLightBoostWhenAvailable() bool {
rv := objc.Send[bool](c.ID, objc.Sel("automaticallyEnablesLowLightBoostWhenAvailable"))
		return rv
}
func (c AVCaptureDevice) SetAutomaticallyEnablesLowLightBoostWhenAvailable(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyEnablesLowLightBoostWhenAvailable:"), value)
}
// The current device-specific RGB white balance gain values in use.
//
// # Discussion
// 
// This property specifies the current red, green, and blue gain values used
// for white balance. You can use the values to adjust color casts for a given
// scene. Each channel supports values between 1.0 and -[MaxWhiteBalanceGain].
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/deviceWhiteBalanceGains
func (c AVCaptureDevice) DeviceWhiteBalanceGains() AVCaptureWhiteBalanceGains {
rv := objc.Send[AVCaptureWhiteBalanceGains](c.ID, objc.Sel("deviceWhiteBalanceGains"))
		return AVCaptureWhiteBalanceGains(rv)
}
// The video zoom factor at which a dual camera device can automatically
// switch between cameras.
//
// # Discussion
// 
// A dual camera device (see [builtInDualCamera]) contains both wide-angle and
// telephoto cameras.
// 
// This property’s value is the zoom factor at which the zoomed field of
// view from the wide-angle camera matches the full field of view from the
// telephoto camera. When the [VideoZoomFactor] setting meets or exceeds this
// value, the device can automatically chooses which camera provides output
// imagery (or automatically combine imagery from both to create final output)
// based on scene conditions. For zoom factors below this value, the device
// always uses imagery from the wide-angle camera.
// 
// On a single-camera device, this value is always `1.0`.
//
// [builtInDualCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInDualCamera
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dualCameraSwitchOverVideoZoomFactor
func (c AVCaptureDevice) DualCameraSwitchOverVideoZoomFactor() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("dualCameraSwitchOverVideoZoomFactor"))
		return rv
}
// The length of time over which exposure takes place.
//
// # Discussion
// 
// The exposure duration is between the active format’s
// [MinExposureDuration] and [MaxExposureDuration].
// 
// To set the exposure duration, call the
// [SetExposureModeCustomWithDurationISOCompletionHandler] method.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureDuration
func (c AVCaptureDevice) ExposureDuration() objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("exposureDuration"))
return objectivec.Object{ID: rv}
}
// The bias to apply to the target exposure value, in exposure value (EV)
// units.
//
// # Discussion
// 
// When the device exposure mode is
// [CaptureExposureModeContinuousAutoExposure] or [CaptureExposureModeLocked],
// the bias affects both metering ([ExposureTargetOffset]), and the actual
// exposure level ([ExposureDuration] and [ISO]). When the exposure mode is
// [CaptureExposureModeCustom], it only affects metering.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureTargetBias
func (c AVCaptureDevice) ExposureTargetBias() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("exposureTargetBias"))
		return rv
}
// The metered exposure level’s offset from the target exposure value, in
// exposure value (EV) units.
//
// # Discussion
// 
// The value of property indicates the difference between the metered exposure
// level of the current scene and the target exposure value.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureTargetOffset
func (c AVCaptureDevice) ExposureTargetOffset() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("exposureTargetOffset"))
		return rv
}
// A Boolean value that indicates whether the device has face-driven
// autoexposure enabled.
//
// # Discussion
// 
// Face-driven autoexposure takes a subject’s face into account when
// performing automatic exposure adjustments. Enabling this feature can better
// expose subjects with darker skin tones or those who are backlit. For apps
// that link against iOS 15.4 or later, the value of this property defaults to
// [true] for devices that support autoexposure.
// 
// Before setting a value for this property, perform the following:
// 
// - Obtain exclusive access to the device by calling its
// [LockForConfiguration] method. - Set the value of the device’s
// [AutomaticallyAdjustsFaceDrivenAutoExposureEnabled] property to [false].
// 
// Attempting to set a value before performing these steps results in an
// exception.
// 
// When you finish configuring the device, unlock it by calling its
// [UnlockForConfiguration] method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFaceDrivenAutoExposureEnabled
func (c AVCaptureDevice) FaceDrivenAutoExposureEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isFaceDrivenAutoExposureEnabled"))
		return rv
}
func (c AVCaptureDevice) SetFaceDrivenAutoExposureEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setFaceDrivenAutoExposureEnabled:"), value)
}
// A Boolean value that indicates whether the device has face-driven autofocus
// enabled.
//
// # Discussion
// 
// Face-driven auto focus takes a subject’s face into account when adjusting
// auto focus. For apps that link against iOS 15.4 or later, the value of this
// property defaults to [true] for devices that support auto focus.
// 
// Before setting a value for this property, perform the following:
// 
// - Obtain exclusive access to the device by calling its
// [LockForConfiguration] method. - Set the value of the device’s
// [AutomaticallyAdjustsFaceDrivenAutoFocusEnabled] property to [false].
// 
// Attempting to set a value before performing these steps results in an
// exception.
// 
// When you finish configuring the device, unlock it by calling its
// [UnlockForConfiguration] method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFaceDrivenAutoFocusEnabled
func (c AVCaptureDevice) FaceDrivenAutoFocusEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isFaceDrivenAutoFocusEnabled"))
		return rv
}
func (c AVCaptureDevice) SetFaceDrivenAutoFocusEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setFaceDrivenAutoFocusEnabled:"), value)
}
// A Boolean value that indicates whether the flash is currently active.
//
// # Discussion
// 
// When the flash is active, it flashes when capturing a photo.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFlashActive
func (c AVCaptureDevice) FlashActive() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isFlashActive"))
		return rv
}
// A Boolean value that indicates whether geometric distortion correction is
// enabled for this device.
//
// # Discussion
// 
// When the device supports geometric distortion correction (GDC), the default
// value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isGeometricDistortionCorrectionEnabled
func (c AVCaptureDevice) GeometricDistortionCorrectionEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isGeometricDistortionCorrectionEnabled"))
		return rv
}
func (c AVCaptureDevice) SetGeometricDistortionCorrectionEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setGeometricDistortionCorrectionEnabled:"), value)
}
// A Boolean value that indicates whether this device supports geometric
// distortion correction.
//
// # Discussion
// 
// Some devices benefit from geometric distortion correction (GDC), such as
// devices with a very wide field of view. GDC lessens the fisheye effect at
// the outer edge of the frame at the cost of losing a small amount of the
// horizontal field of view. When you enable GDC, the device upscales the
// corrected image to the original image size.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isGeometricDistortionCorrectionSupported
func (c AVCaptureDevice) GeometricDistortionCorrectionSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isGeometricDistortionCorrectionSupported"))
		return rv
}
// A Boolean value that indicates whether the device should use global tone
// mapping.
//
// # Discussion
// 
// Tone mapping is a technique used to map the pixel levels in high dynamic
// range images to a reduced dynamic range (such as mapping from 16-bit to
// 8-bit), while still retaining an appearance as close to the original image
// as possible. Normally the active camera uses adaptive, local tone curves to
// preserve the highest image quality and adapt quickly to changing lighting
// conditions.
// 
// When this property value is true, the tone map adjusts dynamically
// depending on the current scene and applies to all pixels in an image. You
// can only enable this setting if the device’s active format’s
// [GlobalToneMappingSupported] property returns [true]. If set to its default
// value of [false], the framework may apply different tone maps to different
// pixels in an image.
// 
// This property resets to its default value of [false] under the following
// conditions:
// 
// - You change the device’s active format. - You add the device’s input
// to a session. - You change the capture session’s preset value.
// 
// Key-value observe this property to observe automatic changes to its value.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isGlobalToneMappingEnabled
func (c AVCaptureDevice) GlobalToneMappingEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isGlobalToneMappingEnabled"))
		return rv
}
func (c AVCaptureDevice) SetGlobalToneMappingEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setGlobalToneMappingEnabled:"), value)
}
// The current device-specific white balance values required for a neutral
// gray white point.
//
// # Discussion
// 
// This property specifies the current red, green, and blue gain values
// derived from the current scene to deliver a neutral (or gray world) white
// point for white balance.
// 
// Gray world values assume you’ve placed a neutral subject (for example, a
// gray card) in the middle of the subject area, and fills the center 50% of
// the frame. Apps can read these values and apply them to the device using
// [SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler].
// 
// Each change supports values between `1.0` and [MaxWhiteBalanceGain]. You
// can read the value at any time, regardless of white balance mode.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/grayWorldDeviceWhiteBalanceGains
func (c AVCaptureDevice) GrayWorldDeviceWhiteBalanceGains() AVCaptureWhiteBalanceGains {
rv := objc.Send[AVCaptureWhiteBalanceGains](c.ID, objc.Sel("grayWorldDeviceWhiteBalanceGains"))
		return AVCaptureWhiteBalanceGains(rv)
}
// The size of the lens diaphragm.
//
// # Discussion
// 
// The value of this property is a float indicating the size (the `f` number)
// of the lens diaphragm.
// 
// This value doesn’t change.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/lensAperture
func (c AVCaptureDevice) LensAperture() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("lensAperture"))
		return rv
}
// The current focus position of the lens.
//
// # Discussion
// 
// A lens position value doesn’t correspond to an exact physical distance,
// nor does it represent a consistent focus distance from device to device.
// 
// The range of possible positions is `0.0` to `1.0`, with `0.0` being the
// shortest distance at which the lens can focus and `1.0` the furthest. Note
// that `1.0` doesn’t represent focus at infinity. The default value is
// `1.0`.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/lensPosition
func (c AVCaptureDevice) LensPosition() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("lensPosition"))
		return rv
}
// A Boolean value that indicates whether the device supports locking focus to
// a specific lens position.
//
// # Discussion
// 
// If this property’s value is [false], calling the
// [SetFocusModeLockedWithLensPositionCompletionHandler] method with a lens
// position value other than [currentLensPosition] raises an exception.
//
// [currentLensPosition]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/currentLensPosition
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isLockingFocusWithCustomLensPositionSupported
func (c AVCaptureDevice) LockingFocusWithCustomLensPositionSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isLockingFocusWithCustomLensPositionSupported"))
		return rv
}
// A Boolean value that indicates whether the device supports locking white
// balance to specific gain values.
//
// # Discussion
// 
// If the value is [false], calling the
// [SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler]
// method with a white balance gains value other than
// [currentWhiteBalanceGains] throws an exception.
//
// [currentWhiteBalanceGains]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/currentWhiteBalanceGains
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isLockingWhiteBalanceWithCustomDeviceGainsSupported
func (c AVCaptureDevice) LockingWhiteBalanceWithCustomDeviceGainsSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isLockingWhiteBalanceWithCustomDeviceGainsSupported"))
		return rv
}
// A Boolean value that indicates whether the capture device’s low light
// boost feature is in an enabled state.
//
// # Discussion
// 
// The value of this property indicates whether the capture device currently
// enhancing images to improve quality due to low light conditions. When this
// property is [true], the capture device has switched into a special mode in
// which it perceives more light in images.
// 
// This property is key-value observable.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isLowLightBoostEnabled
func (c AVCaptureDevice) LowLightBoostEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isLowLightBoostEnabled"))
		return rv
}
// A Boolean value that indicates whether the capture device supports boosting
// images in low-light conditions.
//
// # Discussion
// 
// You can set the capture device’s
// [AutomaticallyEnablesLowLightBoostWhenAvailable] property only if this
// property is [true].
// 
// This property is key-value observable.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isLowLightBoostSupported
func (c AVCaptureDevice) LowLightBoostSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isLowLightBoostSupported"))
		return rv
}
// The maximum zoom factor allowed in the current capture configuration.
//
// # Discussion
// 
// On single-camera devices, this value is always equal to the device
// format’s [VideoMaxZoomFactor] value. On a dual-camera device, the allowed
// range of video zoom factors can change if the device is delivering depth
// data to one or more capture outputs.
// 
// Setting the [VideoZoomFactor] property to (or calling the
// [RampToVideoZoomFactorWithRate] method with) a value greater than the
// device format’s [VideoMaxZoomFactor] value always raises an exception.
// Setting the video zoom factor to a value between the maximum available zoom
// factor and the device format’s maximum clamps the zoom setting to the
// maximum available value.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/maxAvailableVideoZoomFactor
func (c AVCaptureDevice) MaxAvailableVideoZoomFactor() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("maxAvailableVideoZoomFactor"))
		return rv
}
// The maximum supported exposure bias, in exposure value (EV) units.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/maxExposureTargetBias
func (c AVCaptureDevice) MaxExposureTargetBias() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("maxExposureTargetBias"))
		return rv
}
// The maximum supported value to which you can set a color channel.
//
// # Discussion
// 
// This property doesn’t change for the life of the object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/maxWhiteBalanceGain
func (c AVCaptureDevice) MaxWhiteBalanceGain() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("maxWhiteBalanceGain"))
		return rv
}
// The minimum zoom factor allowed in the current capture configuration.
//
// # Discussion
// 
// On single-camera devices, this value is always `1.0`. On a dual-camera
// device, the allowed range of video zoom factors can change if the device is
// delivering depth data to one or more capture outputs.
// 
// Setting the [VideoZoomFactor] property to (or calling the
// [RampToVideoZoomFactorWithRate] method with) a value less than `1.0` always
// raises an exception. Setting the video zoom factor to a value between `1.0`
// and the minimum available zoom factor clamps the zoom setting to the
// minimum.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minAvailableVideoZoomFactor
func (c AVCaptureDevice) MinAvailableVideoZoomFactor() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("minAvailableVideoZoomFactor"))
		return rv
}
// The minimum supported exposure bias, in exposure value (EV) units.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minExposureTargetBias
func (c AVCaptureDevice) MinExposureTargetBias() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("minExposureTargetBias"))
		return rv
}
// A Boolean value that indicates whether a zoom transition is in progress.
//
// # Discussion
// 
// Key-value observe this property to determine when a zoom transitions begins
// or ends.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isRampingVideoZoom
func (c AVCaptureDevice) RampingVideoZoom() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isRampingVideoZoom"))
		return rv
}
// A Boolean value that indicates whether smooth autofocus is in an enabled
// state on the device.
//
// # Discussion
// 
// On capable devices, you can enable a focusing mode in which the camera
// makes lens movements more slowly. This mode make focus transitions less
// visually intrusive, a behavior that you may want for video capture.
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSmoothAutoFocusEnabled
func (c AVCaptureDevice) SmoothAutoFocusEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isSmoothAutoFocusEnabled"))
		return rv
}
func (c AVCaptureDevice) SetSmoothAutoFocusEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setSmoothAutoFocusEnabled:"), value)
}
// A Boolean value that indicates whether the device supports smooth
// autofocus.
//
// # Discussion
// 
// The smooth focusing mode is available only on compatible devices. If this
// property’s value is [false], setting the value of
// [SmoothAutoFocusEnabled] to [true] raises an exception.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSmoothAutoFocusSupported
func (c AVCaptureDevice) SmoothAutoFocusSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isSmoothAutoFocusSupported"))
		return rv
}
// A Boolean value that indicates whether the device streams high dynamic
// range video buffers, also known as extended dynamic range (EDR).
//
// # Discussion
// 
// The device ignores the value of this property when [ActiveColorSpace] is
// HLG BT2020 color space because HDR is effectively always on and can’t be
// disabled.
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// Note that setting this property may cause a lengthy reconfiguration of the
// device, similar to setting a new active format or capture session preset.
// If you’re setting either the active format or the [SessionPreset] this
// property, you should bracket these operations with [BeginConfiguration] and
// session [CommitConfiguration] to minimize reconfiguration time.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVideoHDREnabled
func (c AVCaptureDevice) VideoHDREnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isVideoHDREnabled"))
		return rv
}
func (c AVCaptureDevice) SetVideoHDREnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setVideoHDREnabled:"), value)
}
// A value that controls the cropping and enlargement of images captured by
// the device.
//
// # Discussion
// 
// This value is a multiplier. For example, a value of `2.0` doubles the size
// of an image’s subject (and halves the field of view). Allowed values
// range from `1.0` (full field of view) to the value of the active format’s
// [VideoMaxZoomFactor] property. Setting the value of this property jumps
// immediately to the new zoom factor. For a smooth transition, use the
// [RampToVideoZoomFactorWithRate] method.
// 
// The device achieves a zoom effect by cropping around the center of the
// image captured by the sensor. At low zoom factors, the cropped images is
// equal to or larger than the output size. At higher zoom factors, the device
// must scale the cropped image up to the output size, resulting in a loss of
// image quality. The active format’s [VideoZoomFactorUpscaleThreshold]
// property indicates the factors at which upscaling occurs.
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/videoZoomFactor
func (c AVCaptureDevice) VideoZoomFactor() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("videoZoomFactor"))
		return rv
}
func (c AVCaptureDevice) SetVideoZoomFactor(value float64) {
objc.Send[struct{}](c.ID, objc.Sel("setVideoZoomFactor:"), value)
}
// An array of video zoom factors at or above which a virtual device, such as
// the dual camera, may switch to its next constituent device.
//
// # Discussion
// 
// This property contains zoom factors at which the field of view of one
// constituent device matches the full field of view of the next constituent
// device. The number of switched-over video zoom factors is always one fewer
// than the count of the [ConstituentDevices] property. These factors progress
// in the same order as the devices listed in that property.
// 
// The value of this property is an empty array for nonvirtual devices.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/virtualDeviceSwitchOverVideoZoomFactors
func (c AVCaptureDevice) VirtualDeviceSwitchOverVideoZoomFactors() []foundation.NSNumber {
rv := objc.Send[[]objc.ID](c.ID, objc.Sel("virtualDeviceSwitchOverVideoZoomFactors"))
return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
	return foundation.NSNumberFromID(id)
})
}

