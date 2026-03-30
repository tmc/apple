// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioFile] class.
var (
	_AVAudioFileClass     AVAudioFileClass
	_AVAudioFileClassOnce sync.Once
)

func getAVAudioFileClass() AVAudioFileClass {
	_AVAudioFileClassOnce.Do(func() {
		_AVAudioFileClass = AVAudioFileClass{class: objc.GetClass("AVAudioFile")}
	})
	return _AVAudioFileClass
}

// GetAVAudioFileClass returns the class object for AVAudioFile.
func GetAVAudioFileClass() AVAudioFileClass {
	return getAVAudioFileClass()
}

type AVAudioFileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioFileClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioFileClass) Alloc() AVAudioFile {
	rv := objc.Send[AVAudioFile](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an audio file that the system can open for
// reading or writing.
//
// # Overview
//
// Regardless of the file format, you read and write using [AVAudioPCMBuffer]
// objects. These objects contain samples as [AVAudioCommonFormat] that the
// framework refers to as the file’s processing format. You convert to and
// from using the file’s actual format.
//
// Reads and writes are always sequential. Random access is possible by
// setting the [AVAudioFile.FramePosition] property.
//
// # Creating an Audio File
//
//   - [AVAudioFile.InitForReadingError]: Opens a file for reading using the standard, deinterleaved floating point format.
//   - [AVAudioFile.InitForReadingCommonFormatInterleavedError]: Opens a file for reading using the specified processing format.
//   - [AVAudioFile.InitForWritingSettingsError]: Opens a file for writing using the specified settings.
//   - [AVAudioFile.InitForWritingSettingsCommonFormatInterleavedError]: Opens a file for writing using a specified processing format and settings.
//
// # Reading and Writing the Audio Buffer
//
//   - [AVAudioFile.ReadIntoBufferError]: Reads an entire audio buffer.
//   - [AVAudioFile.ReadIntoBufferFrameCountError]: Reads a portion of an audio buffer using the number of frames you specify.
//   - [AVAudioFile.WriteFromBufferError]: Writes an audio buffer sequentially.
//   - [AVAudioFile.Close]: Closes the audio file.
//
// # Getting Audio File Properties
//
//   - [AVAudioFile.Url]: The location of the audio file.
//   - [AVAudioFile.FileFormat]: The on-disk format of the file.
//   - [AVAudioFile.ProcessingFormat]: The processing format of the file.
//   - [AVAudioFile.Length]: The number of sample frames in the file.
//   - [AVAudioFile.FramePosition]: The position in the file where the next read or write operation occurs.
//   - [AVAudioFile.SetFramePosition]
//   - [AVAudioFile.AVAudioFileTypeKey]: A string that indicates the audio file type.
//   - [AVAudioFile.IsOpen]: A Boolean value that indicates whether the file is open.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile
//
// [AVAudioCommonFormat]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat
type AVAudioFile struct {
	objectivec.Object
}

// AVAudioFileFromID constructs a [AVAudioFile] from an objc.ID.
//
// An object that represents an audio file that the system can open for
// reading or writing.
func AVAudioFileFromID(id objc.ID) AVAudioFile {
	return AVAudioFile{objectivec.Object{ID: id}}
}

// NOTE: AVAudioFile adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioFile] class.
//
// # Creating an Audio File
//
//   - [IAVAudioFile.InitForReadingError]: Opens a file for reading using the standard, deinterleaved floating point format.
//   - [IAVAudioFile.InitForReadingCommonFormatInterleavedError]: Opens a file for reading using the specified processing format.
//   - [IAVAudioFile.InitForWritingSettingsError]: Opens a file for writing using the specified settings.
//   - [IAVAudioFile.InitForWritingSettingsCommonFormatInterleavedError]: Opens a file for writing using a specified processing format and settings.
//
// # Reading and Writing the Audio Buffer
//
//   - [IAVAudioFile.ReadIntoBufferError]: Reads an entire audio buffer.
//   - [IAVAudioFile.ReadIntoBufferFrameCountError]: Reads a portion of an audio buffer using the number of frames you specify.
//   - [IAVAudioFile.WriteFromBufferError]: Writes an audio buffer sequentially.
//   - [IAVAudioFile.Close]: Closes the audio file.
//
// # Getting Audio File Properties
//
//   - [IAVAudioFile.Url]: The location of the audio file.
//   - [IAVAudioFile.FileFormat]: The on-disk format of the file.
//   - [IAVAudioFile.ProcessingFormat]: The processing format of the file.
//   - [IAVAudioFile.Length]: The number of sample frames in the file.
//   - [IAVAudioFile.FramePosition]: The position in the file where the next read or write operation occurs.
//   - [IAVAudioFile.SetFramePosition]
//   - [IAVAudioFile.AVAudioFileTypeKey]: A string that indicates the audio file type.
//   - [IAVAudioFile.IsOpen]: A Boolean value that indicates whether the file is open.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile
type IAVAudioFile interface {
	objectivec.IObject

	// Topic: Creating an Audio File

	// Opens a file for reading using the standard, deinterleaved floating point format.
	InitForReadingError(fileURL foundation.INSURL) (AVAudioFile, error)
	// Opens a file for reading using the specified processing format.
	InitForReadingCommonFormatInterleavedError(fileURL foundation.INSURL, format AVAudioCommonFormat, interleaved bool) (AVAudioFile, error)
	// Opens a file for writing using the specified settings.
	InitForWritingSettingsError(fileURL foundation.INSURL, settings foundation.INSDictionary) (AVAudioFile, error)
	// Opens a file for writing using a specified processing format and settings.
	InitForWritingSettingsCommonFormatInterleavedError(fileURL foundation.INSURL, settings foundation.INSDictionary, format AVAudioCommonFormat, interleaved bool) (AVAudioFile, error)

	// Topic: Reading and Writing the Audio Buffer

	// Reads an entire audio buffer.
	ReadIntoBufferError(buffer IAVAudioPCMBuffer) (bool, error)
	// Reads a portion of an audio buffer using the number of frames you specify.
	ReadIntoBufferFrameCountError(buffer IAVAudioPCMBuffer, frames AVAudioFrameCount) (bool, error)
	// Writes an audio buffer sequentially.
	WriteFromBufferError(buffer IAVAudioPCMBuffer) (bool, error)
	// Closes the audio file.
	Close()

	// Topic: Getting Audio File Properties

	// The location of the audio file.
	Url() foundation.INSURL
	// The on-disk format of the file.
	FileFormat() IAVAudioFormat
	// The processing format of the file.
	ProcessingFormat() IAVAudioFormat
	// The number of sample frames in the file.
	Length() AVAudioFramePosition
	// The position in the file where the next read or write operation occurs.
	FramePosition() AVAudioFramePosition
	SetFramePosition(value AVAudioFramePosition)
	// A string that indicates the audio file type.
	AVAudioFileTypeKey() string
	// A Boolean value that indicates whether the file is open.
	IsOpen() bool
}

// Init initializes the instance.
func (a AVAudioFile) Init() AVAudioFile {
	rv := objc.Send[AVAudioFile](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioFile) Autorelease() AVAudioFile {
	rv := objc.Send[AVAudioFile](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioFile creates a new AVAudioFile instance.
func NewAVAudioFile() AVAudioFile {
	class := getAVAudioFileClass()
	rv := objc.Send[AVAudioFile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Opens a file for reading using the specified processing format.
//
// fileURL: The file to read.
//
// format: The processing format to use when reading from the file.
//
// interleaved: The Boolean value that indicates whether to use an interleaved processing
// format.
//
// # Return Value
//
// A new [AVAudioFile] instance you use for reading.
//
// # Discussion
//
// The processing format refers to the buffers it reads from the file. The
// system reads the content and converts from the file format to the
// processing format. The processing format must be at the same sample rate as
// the actual file contents, and must be linear PCM. The interleaved parameter
// determines whether the processing buffer is in an interleaved float format.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forReading:commonFormat:interleaved:)
func NewAudioFileForReadingCommonFormatInterleavedError(fileURL foundation.INSURL, format AVAudioCommonFormat, interleaved bool) (AVAudioFile, error) {
	var errorPtr objc.ID
	instance := getAVAudioFileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForReading:commonFormat:interleaved:error:"), fileURL, format, interleaved, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil
}

// Opens a file for reading using the standard, deinterleaved floating point
// format.
//
// fileURL: The file to read.
//
// # Return Value
//
// A new [AVAudioFile] instance you use for reading.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forReading:)
func NewAudioFileForReadingError(fileURL foundation.INSURL) (AVAudioFile, error) {
	var errorPtr objc.ID
	instance := getAVAudioFileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForReading:error:"), fileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil
}

// Opens a file for writing using a specified processing format and settings.
//
// fileURL: The path at which to create the file.
//
// settings: The format of the file to create.
//
// format: The processing format to use when writing to the file.
//
// interleaved: The Boolean value that indicates whether to use an interleaved processing
// format.
//
// # Return Value
//
// A new [AVAudioFile] instance for writing.
//
// # Discussion
//
// This method infers the file type to create from the file extension of
// `fileURL`, and overwrites a file at the specified URL if a file exists.
//
// For more information about the `settings` parameter, see the [Settings]
// property in the [AVAudioRecorder] class.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forWriting:settings:commonFormat:interleaved:)
func NewAudioFileForWritingSettingsCommonFormatInterleavedError(fileURL foundation.INSURL, settings foundation.INSDictionary, format AVAudioCommonFormat, interleaved bool) (AVAudioFile, error) {
	var errorPtr objc.ID
	instance := getAVAudioFileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForWriting:settings:commonFormat:interleaved:error:"), fileURL, settings, format, interleaved, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil
}

// Opens a file for writing using the specified settings.
//
// fileURL: The path of the file to create for writing.
//
// settings: The format of the file to create.
//
// # Return Value
//
// A new [AVAudioFile] instance for writing.
//
// # Discussion
//
// This method infers the file type to create from the file extension of
// `fileURL`, and overwrites a file at the specified URL if a file exists.
//
// The file opens for writing using the standard format
// [AVAudioPCMFormatFloat32]. For more information about the `settings`
// parameter, see the [Settings] property in the [AVAudioRecorder] class.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forWriting:settings:)
func NewAudioFileForWritingSettingsError(fileURL foundation.INSURL, settings foundation.INSDictionary) (AVAudioFile, error) {
	var errorPtr objc.ID
	instance := getAVAudioFileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForWriting:settings:error:"), fileURL, settings, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil
}

// Opens a file for reading using the standard, deinterleaved floating point
// format.
//
// fileURL: The file to read.
//
// # Return Value
//
// A new [AVAudioFile] instance you use for reading.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forReading:)
func (a AVAudioFile) InitForReadingError(fileURL foundation.INSURL) (AVAudioFile, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initForReading:error:"), fileURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil

}

// Opens a file for reading using the specified processing format.
//
// fileURL: The file to read.
//
// format: The processing format to use when reading from the file.
//
// interleaved: The Boolean value that indicates whether to use an interleaved processing
// format.
//
// # Return Value
//
// A new [AVAudioFile] instance you use for reading.
//
// # Discussion
//
// The processing format refers to the buffers it reads from the file. The
// system reads the content and converts from the file format to the
// processing format. The processing format must be at the same sample rate as
// the actual file contents, and must be linear PCM. The interleaved parameter
// determines whether the processing buffer is in an interleaved float format.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forReading:commonFormat:interleaved:)
func (a AVAudioFile) InitForReadingCommonFormatInterleavedError(fileURL foundation.INSURL, format AVAudioCommonFormat, interleaved bool) (AVAudioFile, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initForReading:commonFormat:interleaved:error:"), fileURL, format, interleaved, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil

}

// Opens a file for writing using the specified settings.
//
// fileURL: The path of the file to create for writing.
//
// settings: The format of the file to create.
//
// # Return Value
//
// A new [AVAudioFile] instance for writing.
//
// # Discussion
//
// This method infers the file type to create from the file extension of
// `fileURL`, and overwrites a file at the specified URL if a file exists.
//
// The file opens for writing using the standard format
// [AVAudioPCMFormatFloat32]. For more information about the `settings`
// parameter, see the [Settings] property in the [AVAudioRecorder] class.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forWriting:settings:)
func (a AVAudioFile) InitForWritingSettingsError(fileURL foundation.INSURL, settings foundation.INSDictionary) (AVAudioFile, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initForWriting:settings:error:"), fileURL, settings, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil

}

// Opens a file for writing using a specified processing format and settings.
//
// fileURL: The path at which to create the file.
//
// settings: The format of the file to create.
//
// format: The processing format to use when writing to the file.
//
// interleaved: The Boolean value that indicates whether to use an interleaved processing
// format.
//
// # Return Value
//
// A new [AVAudioFile] instance for writing.
//
// # Discussion
//
// This method infers the file type to create from the file extension of
// `fileURL`, and overwrites a file at the specified URL if a file exists.
//
// For more information about the `settings` parameter, see the [Settings]
// property in the [AVAudioRecorder] class.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/init(forWriting:settings:commonFormat:interleaved:)
func (a AVAudioFile) InitForWritingSettingsCommonFormatInterleavedError(fileURL foundation.INSURL, settings foundation.INSDictionary, format AVAudioCommonFormat, interleaved bool) (AVAudioFile, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initForWriting:settings:commonFormat:interleaved:error:"), fileURL, settings, format, interleaved, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioFile{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioFileFromID(rv), nil

}

// Reads an entire audio buffer.
//
// buffer: The buffer from which to read the file. Its format must match the file’s
// processing format.
//
// # Discussion
//
// When reading sequentially from the [FramePosition] property, the method
// attempts to fill the buffer to its capacity. On return, the buffer’s
// [Length] property indicates the number of sample frames it successfully
// reads.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/read(into:)
func (a AVAudioFile) ReadIntoBufferError(buffer IAVAudioPCMBuffer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("readIntoBuffer:error:"), buffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readIntoBuffer:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Reads a portion of an audio buffer using the number of frames you specify.
//
// buffer: The buffer from which to read the file. Its format must match the file’s
// processing format.
//
// frames: The number of frames to read.
//
// # Discussion
//
// You use this method to read fewer frames than the buffer’s
// `frameCapacity`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/read(into:frameCount:)
func (a AVAudioFile) ReadIntoBufferFrameCountError(buffer IAVAudioPCMBuffer, frames AVAudioFrameCount) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("readIntoBuffer:frameCount:error:"), buffer, frames, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readIntoBuffer:frameCount:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Writes an audio buffer sequentially.
//
// buffer: The buffer from which to write to the file. Its format must match the
// file’s processing format.
//
// # Discussion
//
// The buffer’s [FrameLength] signifies how much of the buffer the method
// writes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/write(from:)
func (a AVAudioFile) WriteFromBufferError(buffer IAVAudioPCMBuffer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("writeFromBuffer:error:"), buffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeFromBuffer:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Closes the audio file.
//
// # Discussion
//
// Calling this method closes the underlying file, if open. It’s normally
// unnecessary to close a file opened for reading because it’s automatically
// closed when released. It’s only necessary to close a file opened for
// writing in order to achieve specific control over when the file’s header
// is updated.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/close()
func (a AVAudioFile) Close() {
	objc.Send[objc.ID](a.ID, objc.Sel("close"))
}

// The location of the audio file.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/url
func (a AVAudioFile) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The on-disk format of the file.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/fileFormat
func (a AVAudioFile) FileFormat() IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fileFormat"))
	return AVAudioFormatFromID(objc.ID(rv))
}

// The processing format of the file.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/processingFormat
func (a AVAudioFile) ProcessingFormat() IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("processingFormat"))
	return AVAudioFormatFromID(objc.ID(rv))
}

// The number of sample frames in the file.
//
// # Discussion
//
// This can be computationally expensive to compute for the first time.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/length
func (a AVAudioFile) Length() AVAudioFramePosition {
	rv := objc.Send[AVAudioFramePosition](a.ID, objc.Sel("length"))
	return AVAudioFramePosition(rv)
}

// The position in the file where the next read or write operation occurs.
//
// # Discussion
//
// Set the `framePosition` property to perform a seek before a read or write.
// A read or write operation advances the frame position value by the number
// of frames it reads or writes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/framePosition
func (a AVAudioFile) FramePosition() AVAudioFramePosition {
	rv := objc.Send[AVAudioFramePosition](a.ID, objc.Sel("framePosition"))
	return AVAudioFramePosition(rv)
}
func (a AVAudioFile) SetFramePosition(value AVAudioFramePosition) {
	objc.Send[struct{}](a.ID, objc.Sel("setFramePosition:"), value)
}

// A string that indicates the audio file type.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiofiletypekey
func (a AVAudioFile) AVAudioFileTypeKey() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioFileTypeKey"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the file is open.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFile/isOpen
func (a AVAudioFile) IsOpen() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isOpen"))
	return rv
}
