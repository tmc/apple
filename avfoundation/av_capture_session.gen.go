// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureSession] class.
var (
	_AVCaptureSessionClass     AVCaptureSessionClass
	_AVCaptureSessionClassOnce sync.Once
)

func getAVCaptureSessionClass() AVCaptureSessionClass {
	_AVCaptureSessionClassOnce.Do(func() {
		_AVCaptureSessionClass = AVCaptureSessionClass{class: objc.GetClass("AVCaptureSession")}
	})
	return _AVCaptureSessionClass
}

// GetAVCaptureSessionClass returns the class object for AVCaptureSession.
func GetAVCaptureSessionClass() AVCaptureSessionClass {
	return getAVCaptureSessionClass()
}

type AVCaptureSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureSessionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureSessionClass) Alloc() AVCaptureSession {
	rv := objc.Send[AVCaptureSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that configures capture behavior and coordinates the flow of data
// from input devices to capture outputs.
//
// # Overview
// 
// To perform real-time capture, you instantiate a capture session and add
// appropriate inputs and outputs. The following code fragment illustrates how
// to configure a capture device to record audio.
// 
// Call the [AVCaptureSession.StartRunning] method to start the flow of data from the inputs to
// the outputs, and call the [AVCaptureSession.StopRunning] method to stop the flow.
// 
// You use the [AVCaptureSession.SessionPreset] property to customize the quality level,
// bitrate, or other settings for the output. Most common capture
// configurations are available through session presets; however, some
// specialized options (such as high frame rate) require directly setting a
// capture format on an [AVCaptureDevice] instance.
//
// # Configuring a session
//
//   - [AVCaptureSession.BeginConfiguration]: Marks the beginning of changes to a running capture session’s configuration to perform in a single atomic update.
//   - [AVCaptureSession.CommitConfiguration]: Commits one or more changes to a running capture session’s configuration in a single atomic update.
//
// # Setting a session preset
//
//   - [AVCaptureSession.CanSetSessionPreset]: Determines whether you can configure a capture session with the specified preset.
//   - [AVCaptureSession.SessionPreset]: A preset value that indicates the quality level or bit rate of the output.
//   - [AVCaptureSession.SetSessionPreset]
//
// # Configuring inputs
//
//   - [AVCaptureSession.Inputs]: The inputs that provide media data to a capture session.
//   - [AVCaptureSession.CanAddInput]: Determines whether you can add an input to a session.
//   - [AVCaptureSession.AddInput]: Adds a capture input to the session.
//   - [AVCaptureSession.RemoveInput]: Removes an input from the session.
//
// # Configuring outputs
//
//   - [AVCaptureSession.Outputs]: The output destinations to which a captures session sends its data.
//   - [AVCaptureSession.CanAddOutput]: Determines whether you can add an output to a session.
//   - [AVCaptureSession.AddOutput]: Adds an output to the capture session.
//   - [AVCaptureSession.RemoveOutput]: Removes an output from a capture session.
//
// # Connecting inputs and outputs
//
//   - [AVCaptureSession.Connections]: The connections between inputs and outputs that a capture session contains.
//   - [AVCaptureSession.AddConnection]: Adds a connection to the capture session.
//   - [AVCaptureSession.CanAddConnection]: Determines whether a you can add a connection to a capture session.
//   - [AVCaptureSession.AddInputWithNoConnections]: Adds a capture input to a session without forming any connections.
//   - [AVCaptureSession.AddOutputWithNoConnections]: Adds a capture output to the session without forming any connections.
//   - [AVCaptureSession.RemoveConnection]: Removes a capture connection from the session.
//
// # Configuring deferred start
//
//   - [AVCaptureSession.ManualDeferredStartSupported]: A [BOOL] value that indicates whether the session supports manually running deferred start.
//   - [AVCaptureSession.AutomaticallyRunsDeferredStart]: A Boolean value that indicates whether deferred start runs automatically.
//   - [AVCaptureSession.SetAutomaticallyRunsDeferredStart]
//   - [AVCaptureSession.RunDeferredStartWhenNeeded]: Tells the session to run deferred start when appropriate.
//   - [AVCaptureSession.DeferredStartDelegate]: A delegate object that observes events about deferred start.
//   - [AVCaptureSession.DeferredStartDelegateCallbackQueue]: The dispatch queue on which the session calls deferred start delegate methods.
//   - [AVCaptureSession.SetDeferredStartDelegateDeferredStartDelegateCallbackQueue]: Sets a delegate object for the session to call when performing deferred start.
//
// # Configuring capture controls
//
//   - [AVCaptureSession.SupportsControls]: A Boolean value that indicates whether a capture session supports controls.
//   - [AVCaptureSession.MaxControlsCount]: The maximum number of controls a capture session supports.
//   - [AVCaptureSession.Controls]: The controls that allow configuring the camera system from device hardware.
//   - [AVCaptureSession.CanAddControl]: Returns a Boolean value that indicates whether a capture session add the specified control.
//   - [AVCaptureSession.AddControl]: Adds a control to a capture session.
//   - [AVCaptureSession.RemoveControl]: Removes a control from a capture session.
//   - [AVCaptureSession.SetControlsDelegateQueue]: Sets a delegate object for the system to call when it activates and presents controls.
//   - [AVCaptureSession.ControlsDelegate]: A delegate object that observes changes to the state of capture controls.
//   - [AVCaptureSession.ControlsDelegateCallbackQueue]: The dispatch queue on which the system calls controls delegate methods.
//
// # Managing the session life cycle
//
//   - [AVCaptureSession.StartRunning]: Starts the flow of data through the capture pipeline.
//   - [AVCaptureSession.StopRunning]: Stops the flow of data through the capture pipeline.
//
// # Observing session state
//
//   - [AVCaptureSession.Running]: A Boolean value that indicates whether the capture session is in a running state.
//
// # Synchronizing output
//
//   - [AVCaptureSession.SynchronizationClock]: A clock to use for output synchronization.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession
type AVCaptureSession struct {
	objectivec.Object
}

// AVCaptureSessionFromID constructs a [AVCaptureSession] from an objc.ID.
//
// An object that configures capture behavior and coordinates the flow of data
// from input devices to capture outputs.
func AVCaptureSessionFromID(id objc.ID) AVCaptureSession {
	return AVCaptureSession{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureSession] class.
//
// # Configuring a session
//
//   - [IAVCaptureSession.BeginConfiguration]: Marks the beginning of changes to a running capture session’s configuration to perform in a single atomic update.
//   - [IAVCaptureSession.CommitConfiguration]: Commits one or more changes to a running capture session’s configuration in a single atomic update.
//
// # Setting a session preset
//
//   - [IAVCaptureSession.CanSetSessionPreset]: Determines whether you can configure a capture session with the specified preset.
//   - [IAVCaptureSession.SessionPreset]: A preset value that indicates the quality level or bit rate of the output.
//   - [IAVCaptureSession.SetSessionPreset]
//
// # Configuring inputs
//
//   - [IAVCaptureSession.Inputs]: The inputs that provide media data to a capture session.
//   - [IAVCaptureSession.CanAddInput]: Determines whether you can add an input to a session.
//   - [IAVCaptureSession.AddInput]: Adds a capture input to the session.
//   - [IAVCaptureSession.RemoveInput]: Removes an input from the session.
//
// # Configuring outputs
//
//   - [IAVCaptureSession.Outputs]: The output destinations to which a captures session sends its data.
//   - [IAVCaptureSession.CanAddOutput]: Determines whether you can add an output to a session.
//   - [IAVCaptureSession.AddOutput]: Adds an output to the capture session.
//   - [IAVCaptureSession.RemoveOutput]: Removes an output from a capture session.
//
// # Connecting inputs and outputs
//
//   - [IAVCaptureSession.Connections]: The connections between inputs and outputs that a capture session contains.
//   - [IAVCaptureSession.AddConnection]: Adds a connection to the capture session.
//   - [IAVCaptureSession.CanAddConnection]: Determines whether a you can add a connection to a capture session.
//   - [IAVCaptureSession.AddInputWithNoConnections]: Adds a capture input to a session without forming any connections.
//   - [IAVCaptureSession.AddOutputWithNoConnections]: Adds a capture output to the session without forming any connections.
//   - [IAVCaptureSession.RemoveConnection]: Removes a capture connection from the session.
//
// # Configuring deferred start
//
//   - [IAVCaptureSession.ManualDeferredStartSupported]: A [BOOL] value that indicates whether the session supports manually running deferred start.
//   - [IAVCaptureSession.AutomaticallyRunsDeferredStart]: A Boolean value that indicates whether deferred start runs automatically.
//   - [IAVCaptureSession.SetAutomaticallyRunsDeferredStart]
//   - [IAVCaptureSession.RunDeferredStartWhenNeeded]: Tells the session to run deferred start when appropriate.
//   - [IAVCaptureSession.DeferredStartDelegate]: A delegate object that observes events about deferred start.
//   - [IAVCaptureSession.DeferredStartDelegateCallbackQueue]: The dispatch queue on which the session calls deferred start delegate methods.
//   - [IAVCaptureSession.SetDeferredStartDelegateDeferredStartDelegateCallbackQueue]: Sets a delegate object for the session to call when performing deferred start.
//
// # Configuring capture controls
//
//   - [IAVCaptureSession.SupportsControls]: A Boolean value that indicates whether a capture session supports controls.
//   - [IAVCaptureSession.MaxControlsCount]: The maximum number of controls a capture session supports.
//   - [IAVCaptureSession.Controls]: The controls that allow configuring the camera system from device hardware.
//   - [IAVCaptureSession.CanAddControl]: Returns a Boolean value that indicates whether a capture session add the specified control.
//   - [IAVCaptureSession.AddControl]: Adds a control to a capture session.
//   - [IAVCaptureSession.RemoveControl]: Removes a control from a capture session.
//   - [IAVCaptureSession.SetControlsDelegateQueue]: Sets a delegate object for the system to call when it activates and presents controls.
//   - [IAVCaptureSession.ControlsDelegate]: A delegate object that observes changes to the state of capture controls.
//   - [IAVCaptureSession.ControlsDelegateCallbackQueue]: The dispatch queue on which the system calls controls delegate methods.
//
// # Managing the session life cycle
//
//   - [IAVCaptureSession.StartRunning]: Starts the flow of data through the capture pipeline.
//   - [IAVCaptureSession.StopRunning]: Stops the flow of data through the capture pipeline.
//
// # Observing session state
//
//   - [IAVCaptureSession.Running]: A Boolean value that indicates whether the capture session is in a running state.
//
// # Synchronizing output
//
//   - [IAVCaptureSession.SynchronizationClock]: A clock to use for output synchronization.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession
type IAVCaptureSession interface {
	objectivec.IObject

	// Topic: Configuring a session

	// Marks the beginning of changes to a running capture session’s configuration to perform in a single atomic update.
	BeginConfiguration()
	// Commits one or more changes to a running capture session’s configuration in a single atomic update.
	CommitConfiguration()

	// Topic: Setting a session preset

	// Determines whether you can configure a capture session with the specified preset.
	CanSetSessionPreset(preset AVCaptureSessionPreset) bool
	// A preset value that indicates the quality level or bit rate of the output.
	SessionPreset() AVCaptureSessionPreset
	SetSessionPreset(value AVCaptureSessionPreset)

	// Topic: Configuring inputs

	// The inputs that provide media data to a capture session.
	Inputs() []AVCaptureInput
	// Determines whether you can add an input to a session.
	CanAddInput(input IAVCaptureInput) bool
	// Adds a capture input to the session.
	AddInput(input IAVCaptureInput)
	// Removes an input from the session.
	RemoveInput(input IAVCaptureInput)

	// Topic: Configuring outputs

	// The output destinations to which a captures session sends its data.
	Outputs() []AVCaptureOutput
	// Determines whether you can add an output to a session.
	CanAddOutput(output IAVCaptureOutput) bool
	// Adds an output to the capture session.
	AddOutput(output IAVCaptureOutput)
	// Removes an output from a capture session.
	RemoveOutput(output IAVCaptureOutput)

	// Topic: Connecting inputs and outputs

	// The connections between inputs and outputs that a capture session contains.
	Connections() []AVCaptureConnection
	// Adds a connection to the capture session.
	AddConnection(connection IAVCaptureConnection)
	// Determines whether a you can add a connection to a capture session.
	CanAddConnection(connection IAVCaptureConnection) bool
	// Adds a capture input to a session without forming any connections.
	AddInputWithNoConnections(input IAVCaptureInput)
	// Adds a capture output to the session without forming any connections.
	AddOutputWithNoConnections(output IAVCaptureOutput)
	// Removes a capture connection from the session.
	RemoveConnection(connection IAVCaptureConnection)

	// Topic: Configuring deferred start

	// A [BOOL] value that indicates whether the session supports manually running deferred start.
	ManualDeferredStartSupported() bool
	// A Boolean value that indicates whether deferred start runs automatically.
	AutomaticallyRunsDeferredStart() bool
	SetAutomaticallyRunsDeferredStart(value bool)
	// Tells the session to run deferred start when appropriate.
	RunDeferredStartWhenNeeded()
	// A delegate object that observes events about deferred start.
	DeferredStartDelegate() AVCaptureSessionDeferredStartDelegate
	// The dispatch queue on which the session calls deferred start delegate methods.
	DeferredStartDelegateCallbackQueue() dispatch.Queue
	// Sets a delegate object for the session to call when performing deferred start.
	SetDeferredStartDelegateDeferredStartDelegateCallbackQueue(deferredStartDelegate AVCaptureSessionDeferredStartDelegate, deferredStartDelegateCallbackQueue dispatch.Queue)

	// Topic: Configuring capture controls

	// A Boolean value that indicates whether a capture session supports controls.
	SupportsControls() bool
	// The maximum number of controls a capture session supports.
	MaxControlsCount() int
	// The controls that allow configuring the camera system from device hardware.
	Controls() []AVCaptureControl
	// Returns a Boolean value that indicates whether a capture session add the specified control.
	CanAddControl(control IAVCaptureControl) bool
	// Adds a control to a capture session.
	AddControl(control IAVCaptureControl)
	// Removes a control from a capture session.
	RemoveControl(control IAVCaptureControl)
	// Sets a delegate object for the system to call when it activates and presents controls.
	SetControlsDelegateQueue(controlsDelegate AVCaptureSessionControlsDelegate, controlsDelegateCallbackQueue dispatch.Queue)
	// A delegate object that observes changes to the state of capture controls.
	ControlsDelegate() AVCaptureSessionControlsDelegate
	// The dispatch queue on which the system calls controls delegate methods.
	ControlsDelegateCallbackQueue() dispatch.Queue

	// Topic: Managing the session life cycle

	// Starts the flow of data through the capture pipeline.
	StartRunning()
	// Stops the flow of data through the capture pipeline.
	StopRunning()

	// Topic: Observing session state

	// A Boolean value that indicates whether the capture session is in a running state.
	Running() bool

	// Topic: Synchronizing output

	// A clock to use for output synchronization.
	SynchronizationClock() uintptr
}

// Init initializes the instance.
func (c AVCaptureSession) Init() AVCaptureSession {
	rv := objc.Send[AVCaptureSession](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureSession) Autorelease() AVCaptureSession {
	rv := objc.Send[AVCaptureSession](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureSession creates a new AVCaptureSession instance.
func NewAVCaptureSession() AVCaptureSession {
	class := getAVCaptureSessionClass()
	rv := objc.Send[AVCaptureSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Marks the beginning of changes to a running capture session’s
// configuration to perform in a single atomic update.
//
// # Discussion
// 
// Call this method and [CommitConfiguration] to batch multiple configuration
// operations on a running session into an atomic update.
// 
// After you call this method, you can add or remove outputs, alter the
// [SessionPreset], or configure individual capture input or output
// properties. The session configuration doesn’t change until you invoke
// [CommitConfiguration], at which the system updates all settings. You can
// nest [BeginConfiguration] and [CommitConfiguration] pairs, and the system
// applies the changes when you call the outermost commit.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/beginConfiguration()
func (c AVCaptureSession) BeginConfiguration() {
	objc.Send[objc.ID](c.ID, objc.Sel("beginConfiguration"))
}
// Commits one or more changes to a running capture session’s configuration
// in a single atomic update.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/commitConfiguration()
func (c AVCaptureSession) CommitConfiguration() {
	objc.Send[objc.ID](c.ID, objc.Sel("commitConfiguration"))
}
// Determines whether you can configure a capture session with the specified
// preset.
//
// preset: A preset value to test.
//
// # Return Value
// 
// [true] if the capture session supports the preset; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to determine whether the capture session, in its current
// I/O configuration, supports a particular preset. You can only set a preset
// that returns [true] as the capture session’s [SessionPreset] property
// value.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canSetSessionPreset(_:)
func (c AVCaptureSession) CanSetSessionPreset(preset AVCaptureSessionPreset) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canSetSessionPreset:"), objc.String(string(preset)))
	return rv
}
// Determines whether you can add an input to a session.
//
// input: An input to add to the session.
//
// # Return Value
// 
// [true] if you can add the input to the session; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method returns [false] if you can’t add an input to a capture
// session. This occurs, for example, if you attempt to add the input to a
// session twice, or if the input already belongs to another capture session.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddInput(_:)
func (c AVCaptureSession) CanAddInput(input IAVCaptureInput) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canAddInput:"), input)
	return rv
}
// Adds a capture input to the session.
//
// input: An input to add to the session.
//
// # Discussion
// 
// It’s only valid to call this method if [CanAddInput] returns [true].
// 
// You can invoke this method while the session is running.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addInput(_:)
func (c AVCaptureSession) AddInput(input IAVCaptureInput) {
	objc.Send[objc.ID](c.ID, objc.Sel("addInput:"), input)
}
// Removes an input from the session.
//
// input: An input to remove from the capture session.
//
// # Discussion
// 
// You can invoke this method while the session is running.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeInput(_:)
func (c AVCaptureSession) RemoveInput(input IAVCaptureInput) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeInput:"), input)
}
// Determines whether you can add an output to a session.
//
// output: An output to add to the session.
//
// # Return Value
// 
// [true] if you can add the output; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// In iOS and Mac Catalyst, the system imposes the following limitations on
// the combinations of outputs a capture session may contain:
// 
// - An app may add only a single output of a particular type. For apps that
// link against iOS 16 or later, this restriction no longer applies to
// [AVCaptureVideoDataOutput]. - Prior to iOS 16, you can add an
// [AVCaptureVideoDataOutput] and an [AVCaptureMovieFileOutput] to the same
// session, but only one may have its connection active. If you attempt to
// enable both connections, the system chooses the movie file output as the
// active connection and disables the video data output’s connection. For
// apps that link against iOS 16 or later, this restriction no longer exists.
// - Similarly, prior to iOS 16, you can add an [AVCaptureAudioDataOutput] and
// an [AVCaptureMovieFileOutput] to the same session, but only one may have
// its connection active. If you attempt to enable both connections, the
// system chooses the movie file output and disables the audio data output’s
// connection. For apps that link against iOS 16 or later, this restriction no
// longer exists. - An app can’t add an [AVCapturePhotoOutput] and
// [AVCaptureStillImageOutput] to the same session.
//
// [AVCaptureStillImageOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddOutput(_:)
func (c AVCaptureSession) CanAddOutput(output IAVCaptureOutput) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canAddOutput:"), output)
	return rv
}
// Adds an output to the capture session.
//
// output: An output to add to the session.
//
// # Discussion
// 
// You can only add an output to a session using this method if [CanAddOutput]
// returns [true].
// 
// You can invoke this method while the session is running.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addOutput(_:)
func (c AVCaptureSession) AddOutput(output IAVCaptureOutput) {
	objc.Send[objc.ID](c.ID, objc.Sel("addOutput:"), output)
}
// Removes an output from a capture session.
//
// output: An output to remove from the capture session.
//
// # Discussion
// 
// You can call this method while the session is running.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeOutput(_:)
func (c AVCaptureSession) RemoveOutput(output IAVCaptureOutput) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeOutput:"), output)
}
// Adds a connection to the capture session.
//
// connection: The capture connection to add to the session.
//
// # Discussion
// 
// You can only add a capture connection to a session using this method if
// [CanAddConnection] returns [true].
// 
// When using [AddInput] or [AddOutput], the session automatically forms
// connections between all compatible inputs and outputs. Manually adding
// connections is only necessary when adding an input or output with no
// connections.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addConnection(_:)
func (c AVCaptureSession) AddConnection(connection IAVCaptureConnection) {
	objc.Send[objc.ID](c.ID, objc.Sel("addConnection:"), connection)
}
// Determines whether a you can add a connection to a capture session.
//
// connection: A connect object to test.
//
// # Return Value
// 
// [true] if you can add the connection; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddConnection(_:)
func (c AVCaptureSession) CanAddConnection(connection IAVCaptureConnection) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canAddConnection:"), connection)
	return rv
}
// Adds a capture input to a session without forming any connections.
//
// input: The capture input to add to the session.
//
// # Discussion
// 
// You can call this method while the session is running.
// 
// In most cases, use the [AddInput] method to add new inputs to a session.
// Call this method if you require fine-grained control over which inputs
// connect to which outputs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addInputWithNoConnections(_:)
func (c AVCaptureSession) AddInputWithNoConnections(input IAVCaptureInput) {
	objc.Send[objc.ID](c.ID, objc.Sel("addInputWithNoConnections:"), input)
}
// Adds a capture output to the session without forming any connections.
//
// output: The capture output to add to the session.
//
// # Discussion
// 
// You can call this method while the session is running.
// 
// In most cases, use the [AddOutput] method to add new outputs to a session.
// Call this method if you require fine-grained control over which inputs
// connect to which outputs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addOutputWithNoConnections(_:)
func (c AVCaptureSession) AddOutputWithNoConnections(output IAVCaptureOutput) {
	objc.Send[objc.ID](c.ID, objc.Sel("addOutputWithNoConnections:"), output)
}
// Removes a capture connection from the session.
//
// connection: The capture connection to remove from the session.
//
// # Discussion
// 
// You can call this method while the session is running.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeConnection(_:)
func (c AVCaptureSession) RemoveConnection(connection IAVCaptureConnection) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeConnection:"), connection)
}
// Tells the session to run deferred start when appropriate.
//
// # Discussion
// 
// For best perceived startup performance, call this after displaying the
// first frame, so that deferred start processing doesn’t interfere with
// other initialization operations. For example, if using a [CAMetalLayer] to
// draw camera frames, add a `presentHandler` (using
// doc://com.apple.documentation/metal/mtldrawable/addpresentedhandler) to the
// first drawable and call [RunDeferredStartWhenNeeded] from there.
// 
// If one or more outputs need to start to perform a capture operation, and
// [RunDeferredStartWhenNeeded] has not run yet, the session runs the deferred
// start on your app’s behalf. Only call this method once for each
// configuration commit - after the first call, subsequent calls to
// [RunDeferredStartWhenNeeded] have no effect. The deferred start runs
// asynchronously, so this method returns immediately.
//
// [CAMetalLayer]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/runDeferredStartWhenNeeded()
func (c AVCaptureSession) RunDeferredStartWhenNeeded() {
	objc.Send[objc.ID](c.ID, objc.Sel("runDeferredStartWhenNeeded"))
}
// Sets a delegate object for the session to call when performing deferred
// start.
//
// deferredStartDelegate: An object conforming to the [AVCaptureSessionDeferredStartDelegate]
// protocol that receives events about deferred start.
//
// deferredStartDelegateCallbackQueue: A dispatch queue on which deferredStart delegate methods are called.
//
// # Discussion
// 
// This delegate receives a call to the [SessionWillRunDeferredStart] method
// when deferred start is about to run. It is non-blocking, so by the time
// this method is called, the deferred start may already be underway. If you
// want your app to perform initialization (potentially) concurrently with
// deferred start (e.g. user-facing camera features that are not needed to
// display the first preview frame, but are available to the user as soon as
// possible) it may be done in the delegate’s [SessionWillRunDeferredStart]
// method. To wait until deferred start is finished to perform some remaining
// initialization work, use the [SessionDidRunDeferredStart] method instead.
// 
// The delegate receives a call to the [SessionDidRunDeferredStart] method
// when the deferred start finishes running. This allows you to run
// less-critical application initialization code. For example, if you’ve
// deferred an [AVCapturePhotoOutput] by setting its [DeferredStartEnabled]
// property to `true`, and you’d like to do some app-specific initialization
// related to still capture, here might be a good place to put it.
// 
// If the delegate is non-nil, the session still calls the
// [SessionWillRunDeferredStart] and [SessionDidRunDeferredStart] methods
// regardless of the value of the session’s [AutomaticallyRunsDeferredStart]
// property.
// 
// To minimize the capture session’s startup latency, defer all unnecessary
// work until after the session starts. This delegate provides callbacks for
// you to schedule deferred work without impacting session startup
// performance.
// 
// To perform initialization prior to deferred start but after the user
// interface displays, set [AutomaticallyRunsDeferredStart] to `false`, and
// then run the custom initialization prior to calling
// [RunDeferredStartWhenNeeded].
// 
// If [DeferredStartDelegate] is not [NULL], the session throws an exception
// if [DeferredStartDelegateCallbackQueue] is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/setDeferredStartDelegate(_:deferredStartDelegateCallbackQueue:)
func (c AVCaptureSession) SetDeferredStartDelegateDeferredStartDelegateCallbackQueue(deferredStartDelegate AVCaptureSessionDeferredStartDelegate, deferredStartDelegateCallbackQueue dispatch.Queue) {
	objc.Send[objc.ID](c.ID, objc.Sel("setDeferredStartDelegate:deferredStartDelegateCallbackQueue:"), deferredStartDelegate, uintptr(deferredStartDelegateCallbackQueue.Handle()))
}
// Returns a Boolean value that indicates whether a capture session add the
// specified control.
//
// control: The capture control to add.
//
// # Return Value
// 
// [true] if the capture session can add the control; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Call this method to determine whether you can successfully add a control to
// a capture session using the [AddControl] method. A capture session may not
// be able to add a control due to its current session configuration or if
// unsupported by the host platform.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddControl(_:)
func (c AVCaptureSession) CanAddControl(control IAVCaptureControl) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canAddControl:"), control)
	return rv
}
// Adds a control to a capture session.
//
// control: The capture control to add.
//
// # Discussion
// 
// A capture session may not be able to add a control due to configuration
// reasons or limits of the host platform. Before calling this method,
// determine whether you can successfully add a control by calling the capture
// session’s [CanAddControl] method.
// 
// You may call this method while the session is running.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addControl(_:)
func (c AVCaptureSession) AddControl(control IAVCaptureControl) {
	objc.Send[objc.ID](c.ID, objc.Sel("addControl:"), control)
}
// Removes a control from a capture session.
//
// control: The control to remove.
//
// # Discussion
// 
// You may call this method while the session is running.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeControl(_:)
func (c AVCaptureSession) RemoveControl(control IAVCaptureControl) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeControl:"), control)
}
// Sets a delegate object for the system to call when it activates and
// presents controls.
//
// controlsDelegate: An object that adopts the controls delegate protocol.
//
// controlsDelegateCallbackQueue: A serial dispatch queue on which to call the delegate methods. You must
// specify a serial queue to ensure callbacks occur in order.
// 
// This argument must not be `nil` unless the `controlsDelegate` argument is
// also `nil;` otherwise, the system throws an [InvalidArgumentException].
//
// # Discussion
// 
// People interact with capture controls by performing specific gestures to
// enable their visibility. Specify a delegate to for the system to call when
// it presents and dismisses controls. The system calls the delegate’s
// methods on the specified callback queue.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/setControlsDelegate(_:queue:)
func (c AVCaptureSession) SetControlsDelegateQueue(controlsDelegate AVCaptureSessionControlsDelegate, controlsDelegateCallbackQueue dispatch.Queue) {
	objc.Send[objc.ID](c.ID, objc.Sel("setControlsDelegate:queue:"), controlsDelegate, uintptr(controlsDelegateCallbackQueue.Handle()))
}
// Starts the flow of data through the capture pipeline.
//
// # Discussion
// 
// Call this method to start the flow of data from the capture session’s
// inputs to its outputs. This method is synchronous and blocks until the
// session starts running or it fails, which it reports by posting an
// [runtimeErrorNotification] notification.
//
// [runtimeErrorNotification]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/runtimeErrorNotification
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/startRunning()
func (c AVCaptureSession) StartRunning() {
	objc.Send[objc.ID](c.ID, objc.Sel("startRunning"))
}
// Stops the flow of data through the capture pipeline.
//
// # Discussion
// 
// Call this method to stop the flow of data from the inputs to the outputs
// connected to the capture session. This method is synchronous and blocks
// until the session stops running completely.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/stopRunning()
func (c AVCaptureSession) StopRunning() {
	objc.Send[objc.ID](c.ID, objc.Sel("stopRunning"))
}

// A preset value that indicates the quality level or bit rate of the output.
//
// # Discussion
// 
// Specify a preset value to configure a capture session’s format and
// settings. The default preset is [high], which produces high-quality video
// and audio output, but you can specify any preset value that returns [true]
// for a call to [CanSetSessionPreset].
// 
// You can set this value while the session is running.
//
// [high]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/high
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/sessionPreset
func (c AVCaptureSession) SessionPreset() AVCaptureSessionPreset {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sessionPreset"))
	return AVCaptureSessionPreset(foundation.NSStringFromID(rv).String())
}
func (c AVCaptureSession) SetSessionPreset(value AVCaptureSessionPreset) {
	objc.Send[struct{}](c.ID, objc.Sel("setSessionPreset:"), objc.String(string(value)))
}
// The inputs that provide media data to a capture session.
//
// # Discussion
// 
// You add new inputs to a capture session by callings its [AddInput] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/inputs
func (c AVCaptureSession) Inputs() []AVCaptureInput {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("inputs"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureInput {
		return AVCaptureInputFromID(id)
	})
}
// The output destinations to which a captures session sends its data.
//
// # Discussion
// 
// You add new outputs to a capture session by calling its [AddOutput] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/outputs
func (c AVCaptureSession) Outputs() []AVCaptureOutput {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("outputs"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureOutput {
		return AVCaptureOutputFromID(id)
	})
}
// The connections between inputs and outputs that a capture session contains.
//
// # Discussion
// 
// A capture session automatically forms connections between inputs and
// outputs when you call the [AddInput] or [AddOutput] methods. You can
// explicitly add connections to a session by calling the [AddConnection]
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/connections
func (c AVCaptureSession) Connections() []AVCaptureConnection {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("connections"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureConnection {
		return AVCaptureConnectionFromID(id)
	})
}
// A [BOOL] value that indicates whether the session supports manually running
// deferred start.
//
// # Discussion
// 
// Deferred Start is a feature that allows you to control, on a per-output
// basis, whether output objects start when or after the session is started.
// The session defers starting an output when its `deferredStartEnabled`
// property is set to `true`, and starts it after the session is started.
// 
// You can only set the [AutomaticallyRunsDeferredStart] property value to
// `false` if the session supports manual deferred start.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isManualDeferredStartSupported
func (c AVCaptureSession) ManualDeferredStartSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isManualDeferredStartSupported"))
	return rv
}
// A Boolean value that indicates whether deferred start runs automatically.
//
// # Discussion
// 
// Deferred Start is a feature that allows you to control, on a per-output
// basis, whether output objects start when or after the session is started.
// The session defers starting an output when its [DeferredStartEnabled]
// property is set to `true`, and starts it after the session is started.
// 
// When this value is `true`, [AVCaptureSession] automatically runs deferred
// start. If only [AVCaptureVideoPreviewLayer] objects have
// [DeferredStartEnabled] set to `false`, the session runs deferred start a
// short time after displaying the first frame. If there are [AVCaptureOutput]
// objects that have [DeferredStartEnabled] set to `false`, then the session
// waits until each output that provides streaming data to your app sends its
// first frame.
// 
// If you set this value to `false`, call [RunDeferredStartWhenNeeded] to
// indicate when to run deferred start.
// 
// By default, for apps that are linked on or after iOS 26, this value is
// `true`.
// 
// If [ManualDeferredStartSupported] is `false`, setting this property value
// to false results in the session throwing an invalid argument exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyRunsDeferredStart
func (c AVCaptureSession) AutomaticallyRunsDeferredStart() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("automaticallyRunsDeferredStart"))
	return rv
}
func (c AVCaptureSession) SetAutomaticallyRunsDeferredStart(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyRunsDeferredStart:"), value)
}
// A delegate object that observes events about deferred start.
//
// # Discussion
// 
// Call the [SetDeferredStartDelegateDeferredStartDelegateCallbackQueue]
// method to set the deferred start delegate for a session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/deferredStartDelegate
func (c AVCaptureSession) DeferredStartDelegate() AVCaptureSessionDeferredStartDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("deferredStartDelegate"))
	return AVCaptureSessionDeferredStartDelegateObjectFromID(rv)
}
// The dispatch queue on which the session calls deferred start delegate
// methods.
//
// # Discussion
// 
// Call the [SetDeferredStartDelegateDeferredStartDelegateCallbackQueue]
// method to specify the dispatch queue on which to call the deferred start
// delegate methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/deferredStartDelegateCallbackQueue
func (c AVCaptureSession) DeferredStartDelegateCallbackQueue() dispatch.Queue {
	rv := objc.Send[uintptr](c.ID, objc.Sel("deferredStartDelegateCallbackQueue"))
	return dispatch.QueueFromHandle(rv)
}
// A Boolean value that indicates whether a capture session supports controls.
//
// # Discussion
// 
// A capture session supports controls only on platforms that provide the
// required hardware.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/supportsControls
func (c AVCaptureSession) SupportsControls() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supportsControls"))
	return rv
}
// The maximum number of controls a capture session supports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/maxControlsCount
func (c AVCaptureSession) MaxControlsCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("maxControlsCount"))
	return rv
}
// The controls that allow configuring the camera system from device hardware.
//
// # Discussion
// 
// You modify the contents of this array by calling the [AddControl] and
// [RemoveControl] methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/controls
func (c AVCaptureSession) Controls() []AVCaptureControl {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("controls"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureControl {
		return AVCaptureControlFromID(id)
	})
}
// A delegate object that observes changes to the state of capture controls.
//
// # Discussion
// 
// Call the [SetControlsDelegateQueue] method to set the controls delegate for
// a session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/controlsDelegate
func (c AVCaptureSession) ControlsDelegate() AVCaptureSessionControlsDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("controlsDelegate"))
	return AVCaptureSessionControlsDelegateObjectFromID(rv)
}
// The dispatch queue on which the system calls controls delegate methods.
//
// # Discussion
// 
// Call the [SetControlsDelegateQueue] method to specify the dispatch queue on
// which to call the controls delegate methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/controlsDelegateCallbackQueue
func (c AVCaptureSession) ControlsDelegateCallbackQueue() dispatch.Queue {
	rv := objc.Send[uintptr](c.ID, objc.Sel("controlsDelegateCallbackQueue"))
	return dispatch.QueueFromHandle(rv)
}
// A Boolean value that indicates whether the capture session is in a running
// state.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isRunning
func (c AVCaptureSession) Running() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isRunning"))
	return rv
}
// A clock to use for output synchronization.
//
// # Discussion
// 
// All capture output sample buffer timestamps are on the synchronization
// clock’s timebase. Use this clock in conjunction with the clock from an
// [AVCaptureInputPort] object to synchronize capture output with external
// data sources such as Core Motion samples.
// 
// The example below shows how to reverse synchronize the output timestamps to
// the original timestamps in the
// [CaptureOutputDidOutputSampleBufferFromConnection] method:
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/synchronizationClock
func (c AVCaptureSession) SynchronizationClock() uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("synchronizationClock"))
	return rv
}

