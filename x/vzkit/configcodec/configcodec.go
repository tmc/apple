package configcodec

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
)

// Format identifies a Virtualization configuration format.
type Format uint64

// DefaultFormat is the framework default configuration format.
const DefaultFormat Format = 0

// Encoder wraps the private configuration encoder.
type Encoder struct {
	raw pvz.VZVirtualMachineConfigurationEncoder
}

// NewEncoder creates a configuration encoder.
func NewEncoder() Encoder {
	return Encoder{raw: pvz.NewVZVirtualMachineConfigurationEncoder()}
}

// NewEncoderWithBaseURL creates a configuration encoder rooted at baseURL.
func NewEncoderWithBaseURL(url foundation.INSURL) Encoder {
	return Encoder{raw: pvz.NewVZVirtualMachineConfigurationEncoderWithBaseURL(url)}
}

// NewEncoderWithBasePath creates a configuration encoder rooted at a file URL
// for basePath.
func NewEncoderWithBasePath(basePath string) Encoder {
	if basePath == "" {
		return NewEncoder()
	}
	return NewEncoderWithBaseURL(foundation.NewURLFileURLWithPath(basePath))
}

func (e Encoder) valid() error {
	if e.raw.ID == 0 {
		return fmt.Errorf("create configuration encoder")
	}
	return nil
}

// Data encodes the configuration into opaque data.
func (e Encoder) Data(configuration objectivec.IObject, format uint64) (objectivec.IObject, error) {
	if err := e.valid(); err != nil {
		return nil, err
	}
	return e.raw.DataWithConfigurationFormatError(configuration, format)
}

// Bytes encodes the configuration and returns the raw bytes from the
// resulting NSData.
func (e Encoder) Bytes(configuration objectivec.IObject, format Format) ([]byte, error) {
	data, err := e.Data(configuration, uint64(format))
	if err != nil {
		return nil, err
	}
	if data == nil || data.GetID() == 0 {
		return nil, fmt.Errorf("encode configuration: empty data")
	}
	nsdata := foundation.NSDataFromID(data.GetID())
	length := int(nsdata.Length())
	if length == 0 {
		return []byte{}, nil
	}
	bytes := unsafe.Slice((*byte)(nsdata.Bytes()), length)
	if len(bytes) == 0 {
		return nil, fmt.Errorf("encode configuration: empty data")
	}
	out := make([]byte, len(bytes))
	copy(out, bytes)
	return out, nil
}

// Encode encodes configuration to bytes using the default encoder.
func Encode(configuration objectivec.IObject, format Format) ([]byte, error) {
	return NewEncoder().Bytes(configuration, format)
}

// EncodeAtBasePath encodes configuration with a file URL base path.
func EncodeAtBasePath(basePath string, configuration objectivec.IObject, format Format) ([]byte, error) {
	return NewEncoderWithBasePath(basePath).Bytes(configuration, format)
}

// Decoder wraps the private configuration decoder.
type Decoder struct {
	raw pvz.VZVirtualMachineConfigurationDecoder
}

// NewDecoder creates a configuration decoder.
func NewDecoder() Decoder {
	return Decoder{raw: pvz.NewVZVirtualMachineConfigurationDecoder()}
}

// NewDecoderWithBaseURL creates a decoder rooted at baseURL.
func NewDecoderWithBaseURL(url foundation.INSURL) Decoder {
	return Decoder{raw: pvz.NewVZVirtualMachineConfigurationDecoderWithBaseURL(url)}
}

// NewDecoderWithBasePath creates a decoder rooted at a file URL for basePath.
func NewDecoderWithBasePath(basePath string) Decoder {
	if basePath == "" {
		return NewDecoder()
	}
	return NewDecoderWithBaseURL(foundation.NewURLFileURLWithPath(basePath))
}

func (d Decoder) valid() error {
	if d.raw.ID == 0 {
		return fmt.Errorf("create configuration decoder")
	}
	return nil
}

// Configuration decodes configuration data.
func (d Decoder) Configuration(data objectivec.IObject, format unsafe.Pointer) (objectivec.IObject, error) {
	if err := d.valid(); err != nil {
		return nil, err
	}
	return d.raw.ConfigurationFromDataFormatError(data, format)
}

// Decode decodes configuration bytes and returns the VM configuration object.
func (d Decoder) Decode(data []byte, format Format) (pvz.VZVirtualMachineConfiguration, error) {
	if err := d.valid(); err != nil {
		return pvz.VZVirtualMachineConfiguration{}, err
	}
	blob := foundation.NewDataWithBytesLength(data)
	formatValue := uint64(format)
	decoded, err := d.raw.ConfigurationFromDataFormatError(blob, unsafe.Pointer(&formatValue))
	if err != nil {
		return pvz.VZVirtualMachineConfiguration{}, err
	}
	if decoded == nil || decoded.GetID() == 0 {
		return pvz.VZVirtualMachineConfiguration{}, fmt.Errorf("decode configuration: empty result")
	}
	return pvz.VZVirtualMachineConfigurationFromID(decoded.GetID()), nil
}

// Decode decodes bytes into a VM configuration using the default decoder.
func Decode(data []byte, format Format) (pvz.VZVirtualMachineConfiguration, error) {
	return NewDecoder().Decode(data, format)
}

// DecodeAtBasePath decodes bytes using a decoder rooted at basePath.
func DecodeAtBasePath(basePath string, data []byte, format Format) (pvz.VZVirtualMachineConfiguration, error) {
	return NewDecoderWithBasePath(basePath).Decode(data, format)
}

// SaveOptions wraps private VM save options.
type SaveOptions struct {
	raw pvz.VZVirtualMachineSaveOptions
}

// NewSaveOptions creates default save options.
func NewSaveOptions() SaveOptions {
	return SaveOptions{raw: pvz.NewVZVirtualMachineSaveOptions()}
}

// NewSaveOptionsFrom creates save options and applies the requested flags.
func NewSaveOptionsFrom(compress, encrypt bool) (SaveOptions, error) {
	o := NewSaveOptions()
	if o.raw.ID == 0 {
		return o, fmt.Errorf("create save options")
	}
	o.SetCompress(compress)
	o.SetEncrypt(encrypt)
	return o, nil
}

// Compress reports whether compression is enabled.
func (o SaveOptions) Compress() bool {
	return o.raw.Compress()
}

// SetCompress toggles compression.
func (o SaveOptions) SetCompress(v bool) {
	o.raw.SetCompress(v)
}

// Encrypt reports whether encryption is enabled.
func (o SaveOptions) Encrypt() bool {
	return o.raw.Encrypt()
}

// SetEncrypt toggles encryption.
func (o SaveOptions) SetEncrypt(v bool) {
	o.raw.SetEncrypt(v)
}

// StartConfig describes the delegated exception classes to install on a VM.
type StartConfig struct {
	DelegatedExceptionClasses []objectivec.IObject
}

// StartOptions wraps private VM start options.
type StartOptions struct {
	raw pvz.VZVirtualMachineStartOptions
}

// NewStartOptions creates default start options.
func NewStartOptions() StartOptions {
	return StartOptions{raw: pvz.NewVZVirtualMachineStartOptions()}
}

// NewStartOptionsFrom creates start options and applies the requested classes.
func NewStartOptionsFrom(cfg StartConfig) (StartOptions, error) {
	o := NewStartOptions()
	if o.raw.ID == 0 {
		return o, fmt.Errorf("create start options")
	}
	if len(cfg.DelegatedExceptionClasses) > 0 {
		classes := makeObjectArray(cfg.DelegatedExceptionClasses)
		if classes != nil && classes.GetID() != 0 {
			o.SetDelegatedExceptionClasses(classes)
		}
	}
	return o, nil
}

// DelegatedExceptionClasses returns the delegated exception class list.
func (o StartOptions) DelegatedExceptionClasses() foundation.INSArray {
	rv := objc.Send[objc.ID](o.raw.ID, objc.Sel("_delegatedExceptionClasses"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// SetDelegatedExceptionClasses sets the delegated exception classes.
func (o StartOptions) SetDelegatedExceptionClasses(classes objectivec.IObject) {
	o.raw.SetDelegatedExceptionClasses(classes)
}

// MacStartConfig describes macOS-specific start options.
type MacStartConfig struct {
	ForceDFU          bool
	StopInIBootStage1 bool
	StopInIBootStage2 bool
}

// MacStartOptions wraps private macOS-specific VM start options.
type MacStartOptions struct {
	raw pvz.VZMacOSVirtualMachineStartOptions
}

// NewMacStartOptions creates macOS start options.
func NewMacStartOptions() MacStartOptions {
	return MacStartOptions{raw: pvz.NewVZMacOSVirtualMachineStartOptions()}
}

// NewMacStartOptionsFrom creates macOS start options and applies the requested flags.
func NewMacStartOptionsFrom(cfg MacStartConfig) (MacStartOptions, error) {
	o := NewMacStartOptions()
	if o.raw.ID == 0 {
		return o, fmt.Errorf("create macOS start options")
	}
	o.SetForceDFU(cfg.ForceDFU)
	o.SetStopInIBootStage1(cfg.StopInIBootStage1)
	o.SetStopInIBootStage2(cfg.StopInIBootStage2)
	return o, nil
}

// ForceDFU reports whether DFU start is enabled.
func (o MacStartOptions) ForceDFU() bool {
	return objc.Send[bool](o.raw.ID, objc.Sel("_forceDFU"))
}

// SetForceDFU toggles DFU start.
func (o MacStartOptions) SetForceDFU(v bool) {
	o.raw.SetForceDFU(v)
}

// StopInIBootStage1 reports whether stage 1 iBoot stop is enabled.
func (o MacStartOptions) StopInIBootStage1() bool {
	return objc.Send[bool](o.raw.ID, objc.Sel("_stopInIBootStage1"))
}

// SetStopInIBootStage1 toggles stage 1 iBoot stop.
func (o MacStartOptions) SetStopInIBootStage1(v bool) {
	o.raw.SetStopInIBootStage1(v)
}

// StopInIBootStage2 reports whether stage 2 iBoot stop is enabled.
func (o MacStartOptions) StopInIBootStage2() bool {
	return objc.Send[bool](o.raw.ID, objc.Sel("_stopInIBootStage2"))
}

// SetStopInIBootStage2 toggles stage 2 iBoot stop.
func (o MacStartOptions) SetStopInIBootStage2(v bool) {
	o.raw.SetStopInIBootStage2(v)
}

// ErrUnsupportedFormat reports a missing encoding/decoding format.
var ErrUnsupportedFormat = fmt.Errorf("unsupported configuration format")

func makeObjectArray(values []objectivec.IObject) objectivec.IObject {
	if len(values) == 0 {
		return nil
	}
	array := foundation.NewMutableArrayWithCapacity(uint(len(values)))
	for _, value := range values {
		if value == nil || value.GetID() == 0 {
			continue
		}
		array.AddObject(value)
	}
	if array.GetID() == 0 {
		return nil
	}
	return array
}
