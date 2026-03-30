// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

// C struct types

// AVAudio3DAngularOrientation - A structure that represents the angular orientation of the listener in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DAngularOrientation
type AVAudio3DAngularOrientation struct {
	Yaw   float32 // The side-to-side movement of the listener’s head.
	Pitch float32 // The up-and-down movement of the listener’s head.
	Roll  float32 // The tilt of the listener’s head.

}

// AVAudio3DPoint - A structure that represents a point in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DPoint
type AVAudio3DPoint struct {
	X float32 // The location on the x-axis, in meters.
	Y float32 // The location on the y-axis, in meters.
	Z float32 // The location on the z-axis, in meters.

}

// AVAudio3DVectorOrientation - A structure that represents two orthogonal vectors that describe the orientation of the listener in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DVectorOrientation
type AVAudio3DVectorOrientation struct {
	Forward AVAudio3DVector // The forward vector points in the direction that the listener faces.
	Up      AVAudio3DVector // The up vector is orthogonal to the forward vector and points upward from the listener’s head.

}

// AVAudioConverterPrimeInfo - Priming information for audio conversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverterPrimeInfo
type AVAudioConverterPrimeInfo struct {
	LeadingFrames  AVAudioFrameCount // The number of leading (previous) input frames the converter requires to perform a high-quality conversion.
	TrailingFrames AVAudioFrameCount // The number of trailing input frames, past the end input frame, the converter requires to perform a high-quality conversion.

}

// AVAudioVoiceProcessingOtherAudioDuckingConfiguration - The configuration of ducking non-voice audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioVoiceProcessingOtherAudioDuckingConfiguration
type AVAudioVoiceProcessingOtherAudioDuckingConfiguration struct {
	EnableAdvancedDucking bool                                         // Enables advanced ducking which ducks other audio based on the presence of voice activity from local and remote chat participants.
	DuckingLevel          AVAudioVoiceProcessingOtherAudioDuckingLevel // The ducking level of other audio.

}

// AVBeatRange - A specific time range within a music track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVBeatRange-c.struct
type AVBeatRange struct {
	Start  float64
	Length float64
}
