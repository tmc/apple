// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import "github.com/tmc/apple/corefoundation"

func cgBitmapInfoMake(alpha CGImageAlphaInfo, component CGImageComponentInfo, byteOrder CGImageByteOrderInfo, pixelFormat CGImagePixelFormatInfo) CGBitmapInfo {
	return CGBitmapInfo(alpha) | CGBitmapInfo(component) | CGBitmapInfo(byteOrder) | CGBitmapInfo(pixelFormat)
}

func cgPointMake(x, y float64) corefoundation.CGPoint {
	return corefoundation.CGPoint{X: x, Y: y}
}

func cgRectMake(x, y, width, height float64) corefoundation.CGRect {
	return corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: x, Y: y},
		Size:   corefoundation.CGSize{Width: width, Height: height},
	}
}

func cgSizeMake(width, height float64) corefoundation.CGSize {
	return corefoundation.CGSize{Width: width, Height: height}
}

func cgVectorMake(dx, dy float64) corefoundation.CGVector {
	return corefoundation.CGVector{Dx: dx, Dy: dy}
}

