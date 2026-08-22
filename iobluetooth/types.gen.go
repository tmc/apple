// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"encoding/binary"
	"unsafe"

	"github.com/tmc/apple/kernel"
)

// C struct types

// BluetoothAFHHostChannelClassification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothAFHHostChannelClassification
type BluetoothAFHHostChannelClassification struct {
	Data [10]uint8
}

// BluetoothAFHResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothAFHResults
type BluetoothAFHResults struct {
	Handle kernel.BluetoothConnectionHandle
	Mode   kernel.BluetoothAFHMode
	AfhMap [10]uint8
}

// BluetoothDeviceAddress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothDeviceAddress
type BluetoothDeviceAddress struct {
	Data [6]uint8
}

// BluetoothEnhancedSynchronousConnectionInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothEnhancedSynchronousConnectionInfo
type BluetoothEnhancedSynchronousConnectionInfo struct {
	TransmitBandWidth                 kernel.BluetoothHCITransmitBandwidth
	ReceiveBandWidth                  kernel.BluetoothHCIReceiveBandwidth
	TransmitCodingFormat              kernel.BluetoothHCITransmitCodingFormat
	ReceiveCodingFormat               kernel.BluetoothHCIReceiveCodingFormat
	TransmitCodecFrameSize            kernel.BluetoothHCITransmitCodecFrameSize
	ReceiveCodecFrameSize             kernel.BluetoothHCIReceiveCodecFrameSize
	InputBandwidth                    kernel.BluetoothHCIInputBandwidth
	OutputBandwidth                   kernel.BluetoothHCIOutputBandwidth
	InputCodingFormat                 kernel.BluetoothHCIInputCodingFormat
	OutputCodingFormat                kernel.BluetoothHCIOutputCodingFormat
	InputCodedDataSize                kernel.BluetoothHCIInputCodedDataSize
	OutputCodedDataSize               kernel.BluetoothHCIOutputCodedDataSize
	InputPCMDataFormat                kernel.BluetoothHCIInputPCMDataFormat
	OutputPCMDataFormat               kernel.BluetoothHCIOutputPCMDataFormat
	InputPCMSampelPayloadMSBPosition  kernel.BluetoothHCIInputPCMSamplePayloadMSBPosition
	OutputPCMSampelPayloadMSBPosition kernel.BluetoothHCIOutputPCMSamplePayloadMSBPosition
	InputDataPath                     kernel.BluetoothHCIInputDataPath
	OutputDataPath                    kernel.BluetoothHCIOutputDataPath
	InputTransportUnitSize            kernel.BluetoothHCIInputTransportUnitSize
	OutputTransportUnitSize           kernel.BluetoothHCIOutputTransportUnitSize
	MaxLatency                        kernel.BluetoothHCIMaxLatency
	VoiceSetting                      kernel.BluetoothHCIVoiceSetting
	RetransmissionEffort              kernel.BluetoothHCIRetransmissionEffort
	PacketType                        kernel.BluetoothPacketType
}

// BluetoothEventFilterCondition
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothEventFilterCondition
type BluetoothEventFilterCondition struct {
	Data [7]uint8
}

// BluetoothHCIAcceptSynchronousConnectionRequestParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIAcceptSynchronousConnectionRequestParams
type BluetoothHCIAcceptSynchronousConnectionRequestParams struct {
	TransmitBandwidth    uint32
	ReceiveBandwidth     uint32
	MaxLatency           uint16
	ContentFormat        uint16
	RetransmissionEffort uint8
	PacketType           uint16
}

// BluetoothHCIAutomaticFlushTimeoutInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIAutomaticFlushTimeoutInfo
type BluetoothHCIAutomaticFlushTimeoutInfo struct {
	Handle  kernel.BluetoothConnectionHandle
	Timeout kernel.BluetoothHCIAutomaticFlushTimeout
}

// BluetoothHCIBufferSize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIBufferSize
type BluetoothHCIBufferSize struct {
	ACLDataPacketLength    uint16
	SCODataPacketLength    uint8
	TotalNumACLDataPackets uint16
	TotalNumSCODataPackets uint16
}

// BluetoothHCICurrentInquiryAccessCodes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCICurrentInquiryAccessCodes
type BluetoothHCICurrentInquiryAccessCodes struct {
	Count kernel.BluetoothHCIInquiryAccessCodeCount
	Codes *BluetoothHCIInquiryAccessCode
}

// BluetoothHCICurrentInquiryAccessCodesForWrite
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCICurrentInquiryAccessCodesForWrite
type BluetoothHCICurrentInquiryAccessCodesForWrite struct {
	Count kernel.BluetoothHCIInquiryAccessCodeCount
	Codes [192]uint8
}

// BluetoothHCIEncryptionKeySizeInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEncryptionKeySizeInfo
type BluetoothHCIEncryptionKeySizeInfo struct {
	Handle  kernel.BluetoothConnectionHandle
	KeySize kernel.BluetoothHCIEncryptionKeySize
}

// BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams
type BluetoothHCIEnhancedAcceptSynchronousConnectionRequestParams struct {
	TransmitBandwidth                 uint32
	ReceiveBandwidth                  uint32
	TransmitCodingFormat              uint64
	ReceiveCodingFormat               uint64
	TransmitCodecFrameSize            uint16
	ReceiveCodecFrameSize             uint16
	InputBandwidth                    uint32
	OutputBandwidth                   uint32
	InputCodingFormat                 uint64
	OutputCodingFormat                uint64
	InputCodedDataSize                uint16
	OutputCodedDataSize               uint16
	InputPCMDataFormat                uint8
	OutputPCMDataFormat               uint8
	InputPCMSamplePayloadMSBPosition  uint8
	OutputPCMSamplePayloadMSBPosition uint8
	InputDataPath                     uint8
	OutputDataPath                    uint8
	InputTransportUnitSize            uint8
	OutputTransportUnitSize           uint8
	MaxLatency                        uint16
	PacketType                        uint16
	RetransmissionEffort              uint8
}

// BluetoothHCIEnhancedSetupSynchronousConnectionParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEnhancedSetupSynchronousConnectionParams
type BluetoothHCIEnhancedSetupSynchronousConnectionParams struct {
	TransmitBandwidth                 uint32
	ReceiveBandwidth                  uint32
	TransmitCodingFormat              uint64
	ReceiveCodingFormat               uint64
	TransmitCodecFrameSize            uint16
	ReceiveCodecFrameSize             uint16
	InputBandwidth                    uint32
	OutputBandwidth                   uint32
	InputCodingFormat                 uint64
	OutputCodingFormat                uint64
	InputCodedDataSize                uint16
	OutputCodedDataSize               uint16
	InputPCMDataFormat                uint8
	OutputPCMDataFormat               uint8
	InputPCMSamplePayloadMSBPosition  uint8
	OutputPCMSamplePayloadMSBPosition uint8
	InputDataPath                     uint8
	OutputDataPath                    uint8
	InputTransportUnitSize            uint8
	OutputTransportUnitSize           uint8
	MaxLatency                        uint16
	PacketType                        uint16
	RetransmissionEffort              uint8
}

// BluetoothHCIEventAuthenticationCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventAuthenticationCompleteResults
type BluetoothHCIEventAuthenticationCompleteResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
}

// BluetoothHCIEventChangeConnectionLinkKeyCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventChangeConnectionLinkKeyCompleteResults
type BluetoothHCIEventChangeConnectionLinkKeyCompleteResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
}

// BluetoothHCIEventConnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventConnectionCompleteResults
type BluetoothHCIEventConnectionCompleteResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	DeviceAddress    BluetoothDeviceAddress
	LinkType         kernel.BluetoothLinkType
	EncryptionMode   kernel.BluetoothHCIEncryptionMode
}

// BluetoothHCIEventConnectionPacketTypeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventConnectionPacketTypeResults
type BluetoothHCIEventConnectionPacketTypeResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	PacketType       kernel.BluetoothPacketType
}

// BluetoothHCIEventConnectionRequestResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventConnectionRequestResults
type BluetoothHCIEventConnectionRequestResults struct {
	DeviceAddress BluetoothDeviceAddress
	ClassOfDevice kernel.BluetoothClassOfDevice
	LinkType      kernel.BluetoothLinkType
}

// BluetoothHCIEventDataBufferOverflowResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventDataBufferOverflowResults
type BluetoothHCIEventDataBufferOverflowResults struct {
	LinkType kernel.BluetoothLinkType
}

// BluetoothHCIEventDisconnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventDisconnectionCompleteResults
type BluetoothHCIEventDisconnectionCompleteResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	Reason           kernel.BluetoothReasonCode
}

// BluetoothHCIEventEncryptionChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventEncryptionChangeResults
type BluetoothHCIEventEncryptionChangeResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	Enable           kernel.BluetoothEncryptionEnable
}

// BluetoothHCIEventEncryptionKeyRefreshCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventEncryptionKeyRefreshCompleteResults
type BluetoothHCIEventEncryptionKeyRefreshCompleteResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
}

// BluetoothHCIEventFlowSpecificationData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventFlowSpecificationData
type BluetoothHCIEventFlowSpecificationData struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	Flags            uint8
	FlowDirection    uint8
	ServiceType      uint8
	TokenRate        uint32
	TokenBucketSize  uint32
	PeakBandwidth    uint32
	AccessLatency    uint32
}

// BluetoothHCIEventFlushOccurredResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventFlushOccurredResults
type BluetoothHCIEventFlushOccurredResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
}

// BluetoothHCIEventHardwareErrorResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventHardwareErrorResults
type BluetoothHCIEventHardwareErrorResults struct {
	Error kernel.BluetoothHCIStatus
}

// BluetoothHCIEventLEConnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEConnectionCompleteResults
type BluetoothHCIEventLEConnectionCompleteResults struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [17]byte
}

// ConnectionHandle returns the ConnectionHandle field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) ConnectionHandle() kernel.BluetoothConnectionHandle {
	return *(*kernel.BluetoothConnectionHandle)(unsafe.Pointer(&s.storage[0]))
}

// SetConnectionHandle updates the ConnectionHandle field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) SetConnectionHandle(v kernel.BluetoothConnectionHandle) {
	*(*kernel.BluetoothConnectionHandle)(unsafe.Pointer(&s.storage[0])) = v
}

// Role returns the Role field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) Role() uint8 {
	return uint8(s.storage[2])
}

// SetRole updates the Role field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) SetRole(v uint8) {
	s.storage[2] = uint8(v)
}

// PeerAddressType returns the PeerAddressType field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) PeerAddressType() uint8 {
	return uint8(s.storage[3])
}

// SetPeerAddressType updates the PeerAddressType field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) SetPeerAddressType(v uint8) {
	s.storage[3] = uint8(v)
}

// PeerAddress returns the PeerAddress field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) PeerAddress() BluetoothDeviceAddress {
	return *(*BluetoothDeviceAddress)(unsafe.Pointer(&s.storage[4]))
}

// SetPeerAddress updates the PeerAddress field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) SetPeerAddress(v BluetoothDeviceAddress) {
	*(*BluetoothDeviceAddress)(unsafe.Pointer(&s.storage[4])) = v
}

// ConnInterval returns the ConnInterval field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) ConnInterval() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[10:12]))
}

// SetConnInterval updates the ConnInterval field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) SetConnInterval(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[10:12], uint16(v))
}

// ConnLatency returns the ConnLatency field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) ConnLatency() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[12:14]))
}

// SetConnLatency updates the ConnLatency field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) SetConnLatency(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[12:14], uint16(v))
}

// SupervisionTimeout returns the SupervisionTimeout field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) SupervisionTimeout() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[14:16]))
}

// SetSupervisionTimeout updates the SupervisionTimeout field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) SetSupervisionTimeout(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[14:16], uint16(v))
}

// MasterClockAccuracy returns the MasterClockAccuracy field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) MasterClockAccuracy() uint8 {
	return uint8(s.storage[16])
}

// SetMasterClockAccuracy updates the MasterClockAccuracy field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionCompleteResults) SetMasterClockAccuracy(v uint8) {
	s.storage[16] = uint8(v)
}

// BluetoothHCIEventLEConnectionUpdateCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEConnectionUpdateCompleteResults
type BluetoothHCIEventLEConnectionUpdateCompleteResults struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [8]byte
}

// ConnectionHandle returns the ConnectionHandle field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionUpdateCompleteResults) ConnectionHandle() kernel.BluetoothConnectionHandle {
	return *(*kernel.BluetoothConnectionHandle)(unsafe.Pointer(&s.storage[0]))
}

// SetConnectionHandle updates the ConnectionHandle field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionUpdateCompleteResults) SetConnectionHandle(v kernel.BluetoothConnectionHandle) {
	*(*kernel.BluetoothConnectionHandle)(unsafe.Pointer(&s.storage[0])) = v
}

// ConnInterval returns the ConnInterval field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionUpdateCompleteResults) ConnInterval() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetConnInterval updates the ConnInterval field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionUpdateCompleteResults) SetConnInterval(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// ConnLatency returns the ConnLatency field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionUpdateCompleteResults) ConnLatency() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[4:6]))
}

// SetConnLatency updates the ConnLatency field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionUpdateCompleteResults) SetConnLatency(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[4:6], uint16(v))
}

// SupervisionTimeout returns the SupervisionTimeout field from the record's packed storage.
func (s *BluetoothHCIEventLEConnectionUpdateCompleteResults) SupervisionTimeout() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[6:8]))
}

// SetSupervisionTimeout updates the SupervisionTimeout field in the record's packed storage.
func (s *BluetoothHCIEventLEConnectionUpdateCompleteResults) SetSupervisionTimeout(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[6:8], uint16(v))
}

// BluetoothHCIEventLEEnhancedConnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEEnhancedConnectionCompleteResults
type BluetoothHCIEventLEEnhancedConnectionCompleteResults struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [29]byte
}

// ConnectionHandle returns the ConnectionHandle field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) ConnectionHandle() kernel.BluetoothConnectionHandle {
	return *(*kernel.BluetoothConnectionHandle)(unsafe.Pointer(&s.storage[0]))
}

// SetConnectionHandle updates the ConnectionHandle field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetConnectionHandle(v kernel.BluetoothConnectionHandle) {
	*(*kernel.BluetoothConnectionHandle)(unsafe.Pointer(&s.storage[0])) = v
}

// Role returns the Role field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) Role() uint8 {
	return uint8(s.storage[2])
}

// SetRole updates the Role field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetRole(v uint8) {
	s.storage[2] = uint8(v)
}

// PeerAddressType returns the PeerAddressType field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) PeerAddressType() uint8 {
	return uint8(s.storage[3])
}

// SetPeerAddressType updates the PeerAddressType field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetPeerAddressType(v uint8) {
	s.storage[3] = uint8(v)
}

// PeerAddress returns the PeerAddress field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) PeerAddress() BluetoothDeviceAddress {
	return *(*BluetoothDeviceAddress)(unsafe.Pointer(&s.storage[4]))
}

// SetPeerAddress updates the PeerAddress field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetPeerAddress(v BluetoothDeviceAddress) {
	*(*BluetoothDeviceAddress)(unsafe.Pointer(&s.storage[4])) = v
}

// LocalResolvablePrivateAddress returns the LocalResolvablePrivateAddress field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) LocalResolvablePrivateAddress() BluetoothDeviceAddress {
	return *(*BluetoothDeviceAddress)(unsafe.Pointer(&s.storage[10]))
}

// SetLocalResolvablePrivateAddress updates the LocalResolvablePrivateAddress field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetLocalResolvablePrivateAddress(v BluetoothDeviceAddress) {
	*(*BluetoothDeviceAddress)(unsafe.Pointer(&s.storage[10])) = v
}

// PeerResolvablePrivateAddress returns the PeerResolvablePrivateAddress field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) PeerResolvablePrivateAddress() BluetoothDeviceAddress {
	return *(*BluetoothDeviceAddress)(unsafe.Pointer(&s.storage[16]))
}

// SetPeerResolvablePrivateAddress updates the PeerResolvablePrivateAddress field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetPeerResolvablePrivateAddress(v BluetoothDeviceAddress) {
	*(*BluetoothDeviceAddress)(unsafe.Pointer(&s.storage[16])) = v
}

// ConnInterval returns the ConnInterval field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) ConnInterval() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[22:24]))
}

// SetConnInterval updates the ConnInterval field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetConnInterval(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[22:24], uint16(v))
}

// ConnLatency returns the ConnLatency field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) ConnLatency() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[24:26]))
}

// SetConnLatency updates the ConnLatency field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetConnLatency(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[24:26], uint16(v))
}

// SupervisionTimeout returns the SupervisionTimeout field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SupervisionTimeout() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[26:28]))
}

// SetSupervisionTimeout updates the SupervisionTimeout field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetSupervisionTimeout(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[26:28], uint16(v))
}

// MasterClockAccuracy returns the MasterClockAccuracy field from the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) MasterClockAccuracy() uint8 {
	return uint8(s.storage[28])
}

// SetMasterClockAccuracy updates the MasterClockAccuracy field in the record's packed storage.
func (s *BluetoothHCIEventLEEnhancedConnectionCompleteResults) SetMasterClockAccuracy(v uint8) {
	s.storage[28] = uint8(v)
}

// BluetoothHCIEventLELongTermKeyRequestResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLELongTermKeyRequestResults
type BluetoothHCIEventLELongTermKeyRequestResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	RandomNumber     [8]uint8
	Ediv             uint16
}

// BluetoothHCIEventLEMetaResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEMetaResults
type BluetoothHCIEventLEMetaResults struct {
	Length uint8
	Data   [255]uint8
}

// BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults
type BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [10]byte
}

// ConnectionHandle returns the ConnectionHandle field from the record's packed storage.
func (s *BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults) ConnectionHandle() kernel.BluetoothConnectionHandle {
	return *(*kernel.BluetoothConnectionHandle)(unsafe.Pointer(&s.storage[0]))
}

// SetConnectionHandle updates the ConnectionHandle field in the record's packed storage.
func (s *BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults) SetConnectionHandle(v kernel.BluetoothConnectionHandle) {
	*(*kernel.BluetoothConnectionHandle)(unsafe.Pointer(&s.storage[0])) = v
}

// UsedFeatures returns the UsedFeatures field from the record's packed storage.
func (s *BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults) UsedFeatures() BluetoothHCISupportedFeatures {
	return *(*BluetoothHCISupportedFeatures)(unsafe.Pointer(&s.storage[2]))
}

// SetUsedFeatures updates the UsedFeatures field in the record's packed storage.
func (s *BluetoothHCIEventLEReadRemoteUsedFeaturesCompleteResults) SetUsedFeatures(v BluetoothHCISupportedFeatures) {
	*(*BluetoothHCISupportedFeatures)(unsafe.Pointer(&s.storage[2])) = v
}

// BluetoothHCIEventLinkKeyNotificationResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventLinkKeyNotificationResults
type BluetoothHCIEventLinkKeyNotificationResults struct {
	DeviceAddress BluetoothDeviceAddress
	LinkKey       BluetoothKey
	KeyType       kernel.BluetoothKeyType
}

// BluetoothHCIEventMasterLinkKeyCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventMasterLinkKeyCompleteResults
type BluetoothHCIEventMasterLinkKeyCompleteResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	KeyFlag          kernel.BluetoothKeyFlag
}

// BluetoothHCIEventMaxSlotsChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventMaxSlotsChangeResults
type BluetoothHCIEventMaxSlotsChangeResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	MaxSlots         kernel.BluetoothMaxSlots
}

// BluetoothHCIEventModeChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventModeChangeResults
type BluetoothHCIEventModeChangeResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	Mode             kernel.BluetoothHCIConnectionMode
	ModeInterval     kernel.BluetoothHCIModeInterval
}

// BluetoothHCIEventPageScanModeChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventPageScanModeChangeResults
type BluetoothHCIEventPageScanModeChangeResults struct {
	DeviceAddress BluetoothDeviceAddress
	PageScanMode  kernel.BluetoothPageScanMode
}

// BluetoothHCIEventPageScanRepetitionModeChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventPageScanRepetitionModeChangeResults
type BluetoothHCIEventPageScanRepetitionModeChangeResults struct {
	DeviceAddress          BluetoothDeviceAddress
	PageScanRepetitionMode kernel.BluetoothPageScanRepetitionMode
}

// BluetoothHCIEventQoSSetupCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventQoSSetupCompleteResults
type BluetoothHCIEventQoSSetupCompleteResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	SetupParams      BluetoothHCIQualityOfServiceSetupParams
}

// BluetoothHCIEventQoSViolationResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventQoSViolationResults
type BluetoothHCIEventQoSViolationResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
}

// BluetoothHCIEventReadClockOffsetResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadClockOffsetResults
type BluetoothHCIEventReadClockOffsetResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	ClockOffset      kernel.BluetoothClockOffset
}

// BluetoothHCIEventReadExtendedFeaturesResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadExtendedFeaturesResults
type BluetoothHCIEventReadExtendedFeaturesResults struct {
	ConnectionHandle      kernel.BluetoothConnectionHandle
	SupportedFeaturesInfo BluetoothHCIExtendedFeaturesInfo
}

// BluetoothHCIEventReadRemoteExtendedFeaturesResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadRemoteExtendedFeaturesResults
type BluetoothHCIEventReadRemoteExtendedFeaturesResults struct {
	Error            kernel.BluetoothHCIStatus
	ConnectionHandle kernel.BluetoothConnectionHandle
	Page             kernel.BluetoothHCIPageNumber
	MaxPage          kernel.BluetoothHCIPageNumber
	LmpFeatures      BluetoothHCISupportedFeatures
}

// BluetoothHCIEventReadRemoteSupportedFeaturesResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadRemoteSupportedFeaturesResults
type BluetoothHCIEventReadRemoteSupportedFeaturesResults struct {
	Error            kernel.BluetoothHCIStatus
	ConnectionHandle kernel.BluetoothConnectionHandle
	LmpFeatures      BluetoothHCISupportedFeatures
}

// BluetoothHCIEventReadRemoteVersionInfoResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadRemoteVersionInfoResults
type BluetoothHCIEventReadRemoteVersionInfoResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	LmpVersion       kernel.BluetoothLMPVersion
	ManufacturerName kernel.BluetoothManufacturerName
	LmpSubversion    kernel.BluetoothLMPSubversion
}

// BluetoothHCIEventReadSupportedFeaturesResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReadSupportedFeaturesResults
type BluetoothHCIEventReadSupportedFeaturesResults struct {
	ConnectionHandle  kernel.BluetoothConnectionHandle
	SupportedFeatures BluetoothHCISupportedFeatures
}

// BluetoothHCIEventRemoteNameRequestResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventRemoteNameRequestResults
type BluetoothHCIEventRemoteNameRequestResults struct {
	DeviceAddress BluetoothDeviceAddress
	DeviceName    [248]uint8
}

// BluetoothHCIEventReturnLinkKeysResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventReturnLinkKeysResults
type BluetoothHCIEventReturnLinkKeysResults struct {
	NumLinkKeys   uint8
	DeviceAddress BluetoothDeviceAddress
	LinkKey       BluetoothKey
}

// BluetoothHCIEventRoleChangeResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventRoleChangeResults
type BluetoothHCIEventRoleChangeResults struct {
	ConnectionHandle kernel.BluetoothConnectionHandle
	DeviceAddress    BluetoothDeviceAddress
	Role             kernel.BluetoothRole
}

// BluetoothHCIEventSimplePairingCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventSimplePairingCompleteResults
type BluetoothHCIEventSimplePairingCompleteResults struct {
	DeviceAddress BluetoothDeviceAddress
}

// BluetoothHCIEventSniffSubratingResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventSniffSubratingResults
type BluetoothHCIEventSniffSubratingResults struct {
	ConnectionHandle   kernel.BluetoothConnectionHandle
	MaxTransmitLatency uint16
	MaxReceiveLatency  uint16
	MinRemoteTimeout   uint16
	MinLocalTimeout    uint16
}

// BluetoothHCIEventSynchronousConnectionChangedResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventSynchronousConnectionChangedResults
type BluetoothHCIEventSynchronousConnectionChangedResults struct {
	ConnectionHandle     kernel.BluetoothConnectionHandle
	TransmissionInterval uint8
	RetransmissionWindow uint8
	ReceivePacketLength  uint16
	TransmitPacketLength uint16
}

// BluetoothHCIEventSynchronousConnectionCompleteResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventSynchronousConnectionCompleteResults
type BluetoothHCIEventSynchronousConnectionCompleteResults struct {
	ConnectionHandle     kernel.BluetoothConnectionHandle
	DeviceAddress        BluetoothDeviceAddress
	LinkType             kernel.BluetoothLinkType
	TransmissionInterval uint8
	RetransmissionWindow uint8
	ReceivePacketLength  uint16
	TransmitPacketLength uint16
	AirMode              kernel.BluetoothAirMode
}

// BluetoothHCIEventVendorSpecificResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIEventVendorSpecificResults
type BluetoothHCIEventVendorSpecificResults struct {
	Length uint8
	Data   [255]uint8
}

// BluetoothHCIExtendedFeaturesInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIExtendedFeaturesInfo
type BluetoothHCIExtendedFeaturesInfo struct {
	Page    kernel.BluetoothHCIPageNumber
	MaxPage kernel.BluetoothHCIPageNumber
	Data    [8]uint8
}

// BluetoothHCIExtendedInquiryResponse
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIExtendedInquiryResponse
type BluetoothHCIExtendedInquiryResponse struct {
	Data [240]uint8
}

// BluetoothHCIExtendedInquiryResult
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIExtendedInquiryResult
type BluetoothHCIExtendedInquiryResult struct {
	NumberOfReponses        uint8
	DeviceAddress           BluetoothDeviceAddress
	PageScanRepetitionMode  kernel.BluetoothPageScanRepetitionMode
	Reserved                uint8
	ClassOfDevice           kernel.BluetoothClassOfDevice
	ClockOffset             kernel.BluetoothClockOffset
	RSSIValue               kernel.BluetoothHCIRSSIValue
	ExtendedInquiryResponse BluetoothHCIExtendedInquiryResponse
}

// BluetoothHCIFailedContactInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIFailedContactInfo
type BluetoothHCIFailedContactInfo struct {
	Count  kernel.BluetoothHCIFailedContactCount
	Handle kernel.BluetoothConnectionHandle
}

// BluetoothHCIInquiryAccessCode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryAccessCode
type BluetoothHCIInquiryAccessCode struct {
	Data [3]uint8
}

// BluetoothHCIInquiryResult
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryResult
type BluetoothHCIInquiryResult struct {
	DeviceAddress          BluetoothDeviceAddress
	PageScanRepetitionMode kernel.BluetoothPageScanRepetitionMode
	PageScanPeriodMode     kernel.BluetoothHCIPageScanPeriodMode
	PageScanMode           kernel.BluetoothHCIPageScanMode
	ClassOfDevice          kernel.BluetoothClassOfDevice
	ClockOffset            kernel.BluetoothClockOffset
}

// BluetoothHCIInquiryResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryResults
type BluetoothHCIInquiryResults struct {
	Results [50]BluetoothHCIInquiryResult
	Count   uint32
}

// BluetoothHCIInquiryWithRSSIResult
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryWithRSSIResult
type BluetoothHCIInquiryWithRSSIResult struct {
	DeviceAddress          BluetoothDeviceAddress
	PageScanRepetitionMode kernel.BluetoothPageScanRepetitionMode
	Reserved               uint8
	ClassOfDevice          kernel.BluetoothClassOfDevice
	ClockOffset            kernel.BluetoothClockOffset
	RSSIValue              kernel.BluetoothHCIRSSIValue
}

// BluetoothHCIInquiryWithRSSIResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIInquiryWithRSSIResults
type BluetoothHCIInquiryWithRSSIResults struct {
	Results [50]BluetoothHCIInquiryWithRSSIResult
	Count   uint32
}

// BluetoothHCILEBufferSize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCILEBufferSize
type BluetoothHCILEBufferSize struct {
	ACLDataPacketLength    uint16
	TotalNumACLDataPackets uint8
}

// BluetoothHCILinkPolicySettingsInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCILinkPolicySettingsInfo
type BluetoothHCILinkPolicySettingsInfo struct {
	Settings kernel.BluetoothHCILinkPolicySettings
	Handle   kernel.BluetoothConnectionHandle
}

// BluetoothHCILinkQualityInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCILinkQualityInfo
type BluetoothHCILinkQualityInfo struct {
	Handle       kernel.BluetoothConnectionHandle
	QualityValue kernel.BluetoothHCILinkQuality
}

// BluetoothHCILinkSupervisionTimeout
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCILinkSupervisionTimeout
type BluetoothHCILinkSupervisionTimeout struct {
	Handle  kernel.BluetoothConnectionHandle
	Timeout uint16
}

// BluetoothHCIQualityOfServiceSetupParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIQualityOfServiceSetupParams
type BluetoothHCIQualityOfServiceSetupParams struct {
	Flags          uint8
	ServiceType    uint8
	TokenRate      uint32
	PeakBandwidth  uint32
	Latency        uint32
	DelayVariation uint32
}

// BluetoothHCIRSSIInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIRSSIInfo
type BluetoothHCIRSSIInfo struct {
	Handle    kernel.BluetoothConnectionHandle
	RSSIValue kernel.BluetoothHCIRSSIValue
}

// BluetoothHCIReadExtendedInquiryResponseResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIReadExtendedInquiryResponseResults
type BluetoothHCIReadExtendedInquiryResponseResults struct {
	OutFECRequired          kernel.BluetoothHCIFECRequired
	ExtendedInquiryResponse BluetoothHCIExtendedInquiryResponse
}

// BluetoothHCIReadLMPHandleResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIReadLMPHandleResults
type BluetoothHCIReadLMPHandleResults struct {
	Handle     kernel.BluetoothConnectionHandle
	Lmp_handle kernel.BluetoothLMPHandle
	Reserved   uint32
}

// BluetoothHCIReadLocalOOBDataResults
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIReadLocalOOBDataResults
type BluetoothHCIReadLocalOOBDataResults struct {
	Hash       BluetoothHCISimplePairingOOBData
	Randomizer BluetoothHCISimplePairingOOBData
}

// BluetoothHCIRequestCallbackInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIRequestCallbackInfo
type BluetoothHCIRequestCallbackInfo struct {
	UserCallback   uint64
	UserRefCon     uint64
	InternalRefCon uint64
	AsyncIDRefCon  uint64
	Reserved       uint64
}

// BluetoothHCIRoleInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIRoleInfo
type BluetoothHCIRoleInfo struct {
	Role   uint8
	Handle kernel.BluetoothConnectionHandle
}

// BluetoothHCIScanActivity
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIScanActivity
type BluetoothHCIScanActivity struct {
	ScanInterval uint16
	ScanWindow   uint16
}

// BluetoothHCISetupSynchronousConnectionParams
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISetupSynchronousConnectionParams
type BluetoothHCISetupSynchronousConnectionParams struct {
	TransmitBandwidth    uint32
	ReceiveBandwidth     uint32
	MaxLatency           uint16
	VoiceSetting         uint16
	RetransmissionEffort uint8
	PacketType           uint16
}

// BluetoothHCISimplePairingOOBData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISimplePairingOOBData
type BluetoothHCISimplePairingOOBData struct {
	Data [16]uint8
}

// BluetoothHCIStoredLinkKeysInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIStoredLinkKeysInfo
type BluetoothHCIStoredLinkKeysInfo struct {
	NumLinkKeysRead               uint16
	MaxNumLinkKeysAllowedInDevice uint16
}

// BluetoothHCISupportedCommands
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISupportedCommands
type BluetoothHCISupportedCommands struct {
	Data [64]uint8
}

// BluetoothHCISupportedFeatures
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCISupportedFeatures
type BluetoothHCISupportedFeatures struct {
	Data [8]uint8
}

// BluetoothHCITransmitPowerLevelInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCITransmitPowerLevelInfo
type BluetoothHCITransmitPowerLevelInfo struct {
	Handle kernel.BluetoothConnectionHandle
	Level  kernel.BluetoothHCITransmitPowerLevel
}

// BluetoothHCIVersionInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothHCIVersionInfo
type BluetoothHCIVersionInfo struct {
	ManufacturerName kernel.BluetoothManufacturerName
	LmpVersion       kernel.BluetoothLMPVersion
	LmpSubVersion    kernel.BluetoothLMPSubversion
	HciVersion       uint8
	HciRevision      uint16
}

// BluetoothIOCapabilityResponse
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothIOCapabilityResponse
type BluetoothIOCapabilityResponse struct {
	DeviceAddress              BluetoothDeviceAddress
	IoCapability               kernel.BluetoothIOCapability
	OOBDataPresence            kernel.BluetoothOOBDataPresence
	AuthenticationRequirements kernel.BluetoothAuthenticationRequirements
}

// BluetoothIRK
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothIRK
type BluetoothIRK struct {
	Data [16]uint8
}

// BluetoothKey
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothKey
type BluetoothKey struct {
	Data [16]uint8
}

// BluetoothKeypressNotification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothKeypressNotification
type BluetoothKeypressNotification struct {
	DeviceAddress    BluetoothDeviceAddress
	NotificationType kernel.BluetoothKeypressNotificationType
}

// BluetoothL2CAPQualityOfServiceOptions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPQualityOfServiceOptions
type BluetoothL2CAPQualityOfServiceOptions struct {
	Flags           uint8
	ServiceType     uint8
	TokenRate       uint32
	TokenBucketSize uint32
	PeakBandwidth   uint32
	Latency         uint32
	DelayVariation  uint32
}

// BluetoothL2CAPRetransmissionAndFlowControlOptions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothL2CAPRetransmissionAndFlowControlOptions
type BluetoothL2CAPRetransmissionAndFlowControlOptions struct {
	Flags                 uint8
	TxWindowSize          uint8
	MaxTransmit           uint8
	RetransmissionTimeout uint16
	MonitorTimeout        uint16
	MaxPDUPayloadSize     uint16
}

// BluetoothPINCode
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothPINCode
type BluetoothPINCode struct {
	Data [16]uint8
}

// BluetoothReadClockInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothReadClockInfo
type BluetoothReadClockInfo struct {
	Handle   kernel.BluetoothConnectionHandle
	Clock    uint32
	Accuracy uint16
}

// BluetoothRemoteHostSupportedFeaturesNotification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothRemoteHostSupportedFeaturesNotification
type BluetoothRemoteHostSupportedFeaturesNotification struct {
	DeviceAddress         BluetoothDeviceAddress
	HostSupportedFeatures BluetoothHCISupportedFeatures
}

// BluetoothSetEventMask
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothSetEventMask
type BluetoothSetEventMask struct {
	Data [8]uint8
}

// BluetoothSynchronousConnectionInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothSynchronousConnectionInfo
type BluetoothSynchronousConnectionInfo struct {
	TransmitBandWidth    kernel.BluetoothHCITransmitBandwidth
	ReceiveBandWidth     kernel.BluetoothHCIReceiveBandwidth
	MaxLatency           kernel.BluetoothHCIMaxLatency
	VoiceSetting         kernel.BluetoothHCIVoiceSetting
	RetransmissionEffort kernel.BluetoothHCIRetransmissionEffort
	PacketType           kernel.BluetoothPacketType
}

// BluetoothTransportInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothTransportInfo
type BluetoothTransportInfo struct {
	ProductID              uint32
	VendorID               uint32
	Type                   uint32
	ProductName            [35]int8
	VendorName             [35]int8
	TotalDataBytesSent     uint64
	TotalSCOBytesSent      uint64
	TotalDataBytesReceived uint64
	TotalSCOBytesReceived  uint64
}

// BluetoothUserConfirmationRequest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothUserConfirmationRequest
type BluetoothUserConfirmationRequest struct {
	DeviceAddress BluetoothDeviceAddress
	NumericValue  kernel.BluetoothNumericValue
}

// BluetoothUserPasskeyNotification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/BluetoothUserPasskeyNotification
type BluetoothUserPasskeyNotification struct {
	DeviceAddress BluetoothDeviceAddress
	Passkey       kernel.BluetoothPasskey
}

// IOBluetoothDeviceSearchAttributes - Structure used to search for particular devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceSearchAttributes
type IOBluetoothDeviceSearchAttributes struct {
	Options              IOBluetoothDeviceSearchOptions
	MaxResults           uint32
	DeviceAttributeCount uint32
	AttributeList        *IOBluetoothDeviceSearchDeviceAttributes
}

// IOBluetoothDeviceSearchDeviceAttributes - Structure used to search for particular devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceSearchDeviceAttributes
type IOBluetoothDeviceSearchDeviceAttributes struct {
	Address           BluetoothDeviceAddress
	Name              [248]uint8
	ServiceClassMajor kernel.BluetoothServiceClassMajor
	DeviceClassMajor  kernel.BluetoothDeviceClassMajor
	DeviceClassMinor  kernel.BluetoothDeviceClassMinor
}

// IOBluetoothL2CAPChannelDataBlock
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelDataBlock
type IOBluetoothL2CAPChannelDataBlock struct {
	DataPtr  unsafe.Pointer
	DataSize uintptr
}

// IOBluetoothL2CAPChannelEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannelEvent
type IOBluetoothL2CAPChannelEvent struct {
	EventType IOBluetoothL2CAPChannelEventType
	U         [4]uint64
	Status    int32
}

// OBEXAbortCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAbortCommandData
type OBEXAbortCommandData struct {
	HeaderDataPtr    unsafe.Pointer
	HeaderDataLength uintptr
}

// OBEXAbortCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXAbortCommandResponseData
type OBEXAbortCommandResponseData struct {
	ServerResponseOpCode OBEXOpCode
	HeaderDataPtr        unsafe.Pointer
	HeaderDataLength     uintptr
}

// OBEXConnectCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXConnectCommandData
type OBEXConnectCommandData struct {
	HeaderDataPtr    unsafe.Pointer
	HeaderDataLength uintptr
	MaxPacketSize    OBEXMaxPacketLength
	Version          OBEXVersion
	Flags            OBEXFlags
}

// OBEXConnectCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXConnectCommandResponseData
type OBEXConnectCommandResponseData struct {
	ServerResponseOpCode OBEXOpCode
	HeaderDataPtr        unsafe.Pointer
	HeaderDataLength     uintptr
	MaxPacketSize        OBEXMaxPacketLength
	Version              OBEXVersion
	Flags                OBEXFlags
}

// OBEXDisconnectCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXDisconnectCommandData
type OBEXDisconnectCommandData struct {
	HeaderDataPtr    unsafe.Pointer
	HeaderDataLength uintptr
}

// OBEXDisconnectCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXDisconnectCommandResponseData
type OBEXDisconnectCommandResponseData struct {
	ServerResponseOpCode OBEXOpCode
	HeaderDataPtr        unsafe.Pointer
	HeaderDataLength     uintptr
}

// OBEXErrorData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXErrorData
type OBEXErrorData struct {
	Error      OBEXError
	DataPtr    unsafe.Pointer
	DataLength uintptr
}

// OBEXGetCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXGetCommandData
type OBEXGetCommandData struct {
	HeaderDataPtr    unsafe.Pointer
	HeaderDataLength uintptr
}

// OBEXGetCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXGetCommandResponseData
type OBEXGetCommandResponseData struct {
	ServerResponseOpCode OBEXOpCode
	HeaderDataPtr        unsafe.Pointer
	HeaderDataLength     uintptr
}

// OBEXPutCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXPutCommandData
type OBEXPutCommandData struct {
	HeaderDataPtr      unsafe.Pointer
	HeaderDataLength   uintptr
	BodyDataLeftToSend uintptr
}

// OBEXPutCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXPutCommandResponseData
type OBEXPutCommandResponseData struct {
	ServerResponseOpCode OBEXOpCode
	HeaderDataPtr        unsafe.Pointer
	HeaderDataLength     uintptr
}

// OBEXSessionEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSessionEvent
type OBEXSessionEvent struct {
	Type             OBEXSessionEventType
	Session          OBEXSessionRef
	RefCon           unsafe.Pointer
	IsEndOfEventData bool
	Reserved1        unsafe.Pointer
	Reserved2        unsafe.Pointer
	U                [4]uint64
}

// OBEXSetPathCommandData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSetPathCommandData
type OBEXSetPathCommandData struct {
	HeaderDataPtr    unsafe.Pointer
	HeaderDataLength uintptr
	Flags            OBEXFlags
	Constants        OBEXConstants
}

// OBEXSetPathCommandResponseData - Part of the OBEXSessionEvent structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXSetPathCommandResponseData
type OBEXSetPathCommandResponseData struct {
	ServerResponseOpCode OBEXOpCode
	HeaderDataPtr        unsafe.Pointer
	HeaderDataLength     uintptr
	Flags                OBEXFlags
	Constants            OBEXConstants
}

// OBEXTransportEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXTransportEvent
type OBEXTransportEvent struct {
	Type       OBEXTransportEventType
	Status     OBEXError
	DataPtr    unsafe.Pointer
	DataLength uintptr
}
