package coremlcompiler

import "encoding/binary"

// coremldata.bin format (magic=502, mlprogram):
//
//	[0x00] uint32 LE  magic (502)
//	[0x04] uint32 LE  specificationVersion
//	[0x08] [20]byte   zeros (reserved)
//	[0x1c] uint32 LE  compute type string length (7)
//	[0x20] [4]byte    zeros
//	[0x24] [7]byte    compute type string ("generic")
//	[0x2b] byte       spec format indicator
//	[0x2c] [20]byte   zeros (pad to 0x40)
//	[0x40] [11]byte   zeros
//	[0x4b] uint16 LE  ModelDescription proto length
//	[0x4d] [6]byte    zeros (pad to 0x53)
//	[0x53] []byte     ModelDescription protobuf
//	[EOF-16] uint32 LE magic (repeated)
//	[EOF-12] [12]byte  zeros

const (
	coremldataMagic      = 502
	coremldataHeaderLen  = 0x53 // fixed header size before proto data
	coremldataTrailerLen = 16
)

// buildCoreMLData constructs the coremldata.bin binary from a parsed model.
func buildCoreMLData(m *Model) []byte {
	computeType := "generic"
	protoLen := len(m.descriptionRaw)

	totalSize := coremldataHeaderLen + protoLen + coremldataTrailerLen
	buf := make([]byte, totalSize)

	// [0x00] Magic.
	binary.LittleEndian.PutUint32(buf[0x00:], coremldataMagic)

	// [0x04] Spec version.
	binary.LittleEndian.PutUint32(buf[0x04:], uint32(m.SpecVersion))

	// [0x1c] Compute type string length.
	binary.LittleEndian.PutUint32(buf[0x1c:], uint32(len(computeType)))

	// [0x24] Compute type string (no null terminator needed).
	copy(buf[0x24:], computeType)

	// [0x2b] Spec format indicator (observed: 0x06 for spec=5 mlprogram).
	buf[0x2b] = specFormatByte(m.SpecVersion)

	// [0x4b] Proto data length as uint16 LE.
	binary.LittleEndian.PutUint16(buf[0x4b:], uint16(protoLen))

	// [0x53] ModelDescription protobuf bytes.
	copy(buf[coremldataHeaderLen:], m.descriptionRaw)

	// Trailer: magic repeated + zeros.
	binary.LittleEndian.PutUint32(buf[totalSize-coremldataTrailerLen:], coremldataMagic)

	return buf
}

// specFormatByte returns the byte at offset 0x2b based on the spec version.
// Observed values: 0x06 for spec 5 (mlprogram), 0x01 for spec 1 (NeuralNetwork).
func specFormatByte(specVersion int32) byte {
	if specVersion >= 9 {
		return 0x09
	}
	if specVersion >= 5 {
		return 0x06
	}
	return 0x01
}
