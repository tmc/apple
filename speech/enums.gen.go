// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/Speech/SFSpeechError/Code
type SFSpeechErrorCode int

const (
	// SFSpeechErrorCodeAudioReadFailed: The audio file could not be read.
	SFSpeechErrorCodeAudioReadFailed SFSpeechErrorCode = 2
	// SFSpeechErrorCodeInternalServiceError: There was an internal error.
	SFSpeechErrorCodeInternalServiceError SFSpeechErrorCode = 1
	// SFSpeechErrorCodeMalformedSupplementalModel: The custom language model file was malformed.
	SFSpeechErrorCodeMalformedSupplementalModel SFSpeechErrorCode = 8
	// SFSpeechErrorCodeMissingParameter: A required parameter is missing/nil.
	SFSpeechErrorCodeMissingParameter SFSpeechErrorCode = 13
	// SFSpeechErrorCodeTimeout: The operation timed out.
	SFSpeechErrorCodeTimeout SFSpeechErrorCode = 12
	// SFSpeechErrorCodeUndefinedTemplateClassName: The custom language model templates were malformed.
	SFSpeechErrorCodeUndefinedTemplateClassName SFSpeechErrorCode = 7
)

func (e SFSpeechErrorCode) String() string {
	switch e {
	case SFSpeechErrorCodeAudioReadFailed:
		return "SFSpeechErrorCodeAudioReadFailed"
	case SFSpeechErrorCodeInternalServiceError:
		return "SFSpeechErrorCodeInternalServiceError"
	case SFSpeechErrorCodeMalformedSupplementalModel:
		return "SFSpeechErrorCodeMalformedSupplementalModel"
	case SFSpeechErrorCodeMissingParameter:
		return "SFSpeechErrorCodeMissingParameter"
	case SFSpeechErrorCodeTimeout:
		return "SFSpeechErrorCodeTimeout"
	case SFSpeechErrorCodeUndefinedTemplateClassName:
		return "SFSpeechErrorCodeUndefinedTemplateClassName"
	default:
		return fmt.Sprintf("SFSpeechErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint
type SFSpeechRecognitionTaskHint int

const (
	// SFSpeechRecognitionTaskHintConfirmation: A task that uses captured speech for short, confirmation-style requests.
	SFSpeechRecognitionTaskHintConfirmation SFSpeechRecognitionTaskHint = 3
	// SFSpeechRecognitionTaskHintDictation: A task that uses captured speech for text entry.
	SFSpeechRecognitionTaskHintDictation SFSpeechRecognitionTaskHint = 1
	// SFSpeechRecognitionTaskHintSearch: A task that uses captured speech to specify search terms.
	SFSpeechRecognitionTaskHintSearch SFSpeechRecognitionTaskHint = 2
	// SFSpeechRecognitionTaskHintUnspecified: An unspecified type of task.
	SFSpeechRecognitionTaskHintUnspecified SFSpeechRecognitionTaskHint = 0
)

func (e SFSpeechRecognitionTaskHint) String() string {
	switch e {
	case SFSpeechRecognitionTaskHintConfirmation:
		return "SFSpeechRecognitionTaskHintConfirmation"
	case SFSpeechRecognitionTaskHintDictation:
		return "SFSpeechRecognitionTaskHintDictation"
	case SFSpeechRecognitionTaskHintSearch:
		return "SFSpeechRecognitionTaskHintSearch"
	case SFSpeechRecognitionTaskHintUnspecified:
		return "SFSpeechRecognitionTaskHintUnspecified"
	default:
		return fmt.Sprintf("SFSpeechRecognitionTaskHint(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskState
type SFSpeechRecognitionTaskState int

const (
	// SFSpeechRecognitionTaskStateCanceling: Delivery of recognition results has finished, but audio recording may be ongoing.
	SFSpeechRecognitionTaskStateCanceling SFSpeechRecognitionTaskState = 3
	// SFSpeechRecognitionTaskStateCompleted: Delivery of recognition requests has finished and audio recording has stopped.
	SFSpeechRecognitionTaskStateCompleted SFSpeechRecognitionTaskState = 4
	// SFSpeechRecognitionTaskStateFinishing: Audio recording has stopped, but delivery of recognition results may continue.
	SFSpeechRecognitionTaskStateFinishing SFSpeechRecognitionTaskState = 2
	// SFSpeechRecognitionTaskStateRunning: Speech recognition (potentially including audio recording) is in progress.
	SFSpeechRecognitionTaskStateRunning SFSpeechRecognitionTaskState = 1
	// SFSpeechRecognitionTaskStateStarting: Speech recognition (potentially including audio recording) has not yet started.
	SFSpeechRecognitionTaskStateStarting SFSpeechRecognitionTaskState = 0
)

func (e SFSpeechRecognitionTaskState) String() string {
	switch e {
	case SFSpeechRecognitionTaskStateCanceling:
		return "SFSpeechRecognitionTaskStateCanceling"
	case SFSpeechRecognitionTaskStateCompleted:
		return "SFSpeechRecognitionTaskStateCompleted"
	case SFSpeechRecognitionTaskStateFinishing:
		return "SFSpeechRecognitionTaskStateFinishing"
	case SFSpeechRecognitionTaskStateRunning:
		return "SFSpeechRecognitionTaskStateRunning"
	case SFSpeechRecognitionTaskStateStarting:
		return "SFSpeechRecognitionTaskStateStarting"
	default:
		return fmt.Sprintf("SFSpeechRecognitionTaskState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerAuthorizationStatus
type SFSpeechRecognizerAuthorizationStatus int

const (
	// SFSpeechRecognizerAuthorizationStatusAuthorized: The user granted your app’s request to perform speech recognition.
	SFSpeechRecognizerAuthorizationStatusAuthorized SFSpeechRecognizerAuthorizationStatus = 3
	// SFSpeechRecognizerAuthorizationStatusDenied: The user denied your app’s request to perform speech recognition.
	SFSpeechRecognizerAuthorizationStatusDenied SFSpeechRecognizerAuthorizationStatus = 1
	// SFSpeechRecognizerAuthorizationStatusNotDetermined: The app’s authorization status has not yet been determined.
	SFSpeechRecognizerAuthorizationStatusNotDetermined SFSpeechRecognizerAuthorizationStatus = 0
	// SFSpeechRecognizerAuthorizationStatusRestricted: The device prevents your app from performing speech recognition.
	SFSpeechRecognizerAuthorizationStatusRestricted SFSpeechRecognizerAuthorizationStatus = 2
)

func (e SFSpeechRecognizerAuthorizationStatus) String() string {
	switch e {
	case SFSpeechRecognizerAuthorizationStatusAuthorized:
		return "SFSpeechRecognizerAuthorizationStatusAuthorized"
	case SFSpeechRecognizerAuthorizationStatusDenied:
		return "SFSpeechRecognizerAuthorizationStatusDenied"
	case SFSpeechRecognizerAuthorizationStatusNotDetermined:
		return "SFSpeechRecognizerAuthorizationStatusNotDetermined"
	case SFSpeechRecognizerAuthorizationStatusRestricted:
		return "SFSpeechRecognizerAuthorizationStatusRestricted"
	default:
		return fmt.Sprintf("SFSpeechRecognizerAuthorizationStatus(%d)", e)
	}
}
