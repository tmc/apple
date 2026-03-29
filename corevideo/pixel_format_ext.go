package corevideo

// Corrected FourCharCode values for pixel format constants.
// The generated enums.gen.go has these as 0 due to a codegen bug
// with anonymous enum FourCC parsing (fixed in appledocs d5f96cd08d,
// pending regeneration).
const (
	// CVPixelFormatOneComponent16Half is kCVPixelFormatType_OneComponent16Half ('L00h').
	// 16-bit one component IEEE half-precision float.
	CVPixelFormatOneComponent16Half uint32 = 0x4C303068

	// CVPixelFormatOneComponent32Float is kCVPixelFormatType_OneComponent32Float ('L00f').
	// 32-bit one component IEEE float.
	CVPixelFormatOneComponent32Float uint32 = 0x4C303066
)
