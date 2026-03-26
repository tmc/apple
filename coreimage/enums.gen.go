// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/ECCVersion-swift.enum
type CIDataMatrixCodeECCVersion int

const (
	// CIDataMatrixCodeECCVersion000: Indicates error correction using convolutional code error correction with no data protection.
	CIDataMatrixCodeECCVersion000 CIDataMatrixCodeECCVersion = 0
	// CIDataMatrixCodeECCVersion050: Indicates 1/4 of the symbol is dedicated to convolutional code error correction.
	CIDataMatrixCodeECCVersion050 CIDataMatrixCodeECCVersion = 50
	// CIDataMatrixCodeECCVersion080: Indicates 1/3 of the symbol is dedicated to convolutional code error correction.
	CIDataMatrixCodeECCVersion080 CIDataMatrixCodeECCVersion = 80
	// CIDataMatrixCodeECCVersion100: Indicates 1/2 of the symbol is dedicated to convolutional code error correction.
	CIDataMatrixCodeECCVersion100 CIDataMatrixCodeECCVersion = 100
	// CIDataMatrixCodeECCVersion140: Indicates 3/4 of the symbol is dedicated to convolutional code error correction.
	CIDataMatrixCodeECCVersion140 CIDataMatrixCodeECCVersion = 140
	// CIDataMatrixCodeECCVersion200: Indicates error correction using Reed-Solomon error correction.
	CIDataMatrixCodeECCVersion200 CIDataMatrixCodeECCVersion = 200
)

func (e CIDataMatrixCodeECCVersion) String() string {
	switch e {
	case CIDataMatrixCodeECCVersion000:
		return "CIDataMatrixCodeECCVersion000"
	case CIDataMatrixCodeECCVersion050:
		return "CIDataMatrixCodeECCVersion050"
	case CIDataMatrixCodeECCVersion080:
		return "CIDataMatrixCodeECCVersion080"
	case CIDataMatrixCodeECCVersion100:
		return "CIDataMatrixCodeECCVersion100"
	case CIDataMatrixCodeECCVersion140:
		return "CIDataMatrixCodeECCVersion140"
	case CIDataMatrixCodeECCVersion200:
		return "CIDataMatrixCodeECCVersion200"
	default:
		return fmt.Sprintf("CIDataMatrixCodeECCVersion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/ErrorCorrectionLevel-swift.enum
type CIQRCodeErrorCorrectionLevel int

const (
	// CIQRCodeErrorCorrectionLevelH: Indicates that approximately 65% of the symbol data is dedicated to error correction.
	CIQRCodeErrorCorrectionLevelH CIQRCodeErrorCorrectionLevel = 3
	// CIQRCodeErrorCorrectionLevelL: Indicates that approximately 20% of the symbol data is dedicated to error correction.
	CIQRCodeErrorCorrectionLevelL CIQRCodeErrorCorrectionLevel = 0
	// CIQRCodeErrorCorrectionLevelM: Indicates that approximately 37% of the symbol data is dedicated to error correction.
	CIQRCodeErrorCorrectionLevelM CIQRCodeErrorCorrectionLevel = 1
	// CIQRCodeErrorCorrectionLevelQ: Indicates that approximately 55% of the symbol data is dedicated to error correction.
	CIQRCodeErrorCorrectionLevelQ CIQRCodeErrorCorrectionLevel = 2
)

func (e CIQRCodeErrorCorrectionLevel) String() string {
	switch e {
	case CIQRCodeErrorCorrectionLevelH:
		return "CIQRCodeErrorCorrectionLevelH"
	case CIQRCodeErrorCorrectionLevelL:
		return "CIQRCodeErrorCorrectionLevelL"
	case CIQRCodeErrorCorrectionLevelM:
		return "CIQRCodeErrorCorrectionLevelM"
	case CIQRCodeErrorCorrectionLevelQ:
		return "CIQRCodeErrorCorrectionLevelQ"
	default:
		return fmt.Sprintf("CIQRCodeErrorCorrectionLevel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIRenderDestinationAlphaMode
type CIRenderDestinationAlphaMode uint

const (
	// CIRenderDestinationAlphaNone: Designates a destination with no alpha compositing.
	CIRenderDestinationAlphaNone CIRenderDestinationAlphaMode = 0
	// CIRenderDestinationAlphaPremultiplied: Designates a destination that expects premultiplied alpha values.
	CIRenderDestinationAlphaPremultiplied CIRenderDestinationAlphaMode = 1
	// CIRenderDestinationAlphaUnpremultiplied: Designates a destination that expects non-premultiplied alpha values.
	CIRenderDestinationAlphaUnpremultiplied CIRenderDestinationAlphaMode = 2
)

func (e CIRenderDestinationAlphaMode) String() string {
	switch e {
	case CIRenderDestinationAlphaNone:
		return "CIRenderDestinationAlphaNone"
	case CIRenderDestinationAlphaPremultiplied:
		return "CIRenderDestinationAlphaPremultiplied"
	case CIRenderDestinationAlphaUnpremultiplied:
		return "CIRenderDestinationAlphaUnpremultiplied"
	default:
		return fmt.Sprintf("CIRenderDestinationAlphaMode(%d)", e)
	}
}

