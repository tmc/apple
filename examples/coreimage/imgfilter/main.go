// Command imgfilter processes images using Apple's CoreImage framework.
//
// Usage:
//
//	imgfilter <command> [flags] <image> [-o output]
//
// Commands:
//
//	blur         Apply Gaussian blur (-radius N)
//	sharpen      Apply sharpening (-sharpness N)
//	resize       Scale image (-scale N)
//	crop         Crop to rectangle (-x -y -w -h)
//	grayscale    Convert to grayscale
//	qrgen        Generate QR code from text
//
// Image can be a file path or "-" for stdin. Output via -o flag (default: stdout as PNG).
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if len(os.Args) < 2 {
		usage()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	var err error
	switch cmd {
	case "blur":
		err = runBlur(args)
	case "sharpen":
		err = runSharpen(args)
	case "resize":
		err = runResize(args)
	case "crop":
		err = runCrop(args)
	case "grayscale":
		err = runGrayscale(args)
	case "qrgen":
		err = runQRGen(args)
	case "-h", "-help", "--help", "help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "imgfilter: unknown command %q\n", cmd)
		usage()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "imgfilter %s: %v\n", cmd, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: imgfilter <command> [flags] <image> [-o output]

Commands:
  blur         Apply Gaussian blur (-radius N)
  sharpen      Apply sharpening (-sharpness N)
  resize       Scale image (-scale N)
  crop         Crop to rectangle (-x -y -w -h)
  grayscale    Convert to grayscale
  qrgen        Generate QR code from text

Image can be a file path or "-" for stdin.
Output: -o file.png (default: stdout as PNG).
`)
	os.Exit(2)
}

// --- Commands ---

func runBlur(args []string) error {
	fs := flag.NewFlagSet("blur", flag.ExitOnError)
	radius := fs.Float64("radius", 10, "blur radius (sigma)")
	output := fs.String("o", "", "output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: imgfilter blur [-radius N] [-o output] <image>")
	}

	img, cleanup, err := loadCIImage(fs.Arg(0))
	if err != nil {
		return err
	}
	if cleanup != nil {
		defer cleanup()
	}

	blurred := img.ImageByApplyingGaussianBlurWithSigma(*radius)
	return saveCIImage(blurred, *output)
}

func runSharpen(args []string) error {
	fs := flag.NewFlagSet("sharpen", flag.ExitOnError)
	sharpness := fs.Float64("sharpness", 0.4, "sharpness amount (0.0-2.0)")
	output := fs.String("o", "", "output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: imgfilter sharpen [-sharpness N] [-o output] <image>")
	}

	img, cleanup, err := loadCIImage(fs.Arg(0))
	if err != nil {
		return err
	}
	if cleanup != nil {
		defer cleanup()
	}

	filter := coreimage.NewFilterWithName("CISharpenLuminance")
	if filter.ID == 0 {
		return fmt.Errorf("CISharpenLuminance filter not available")
	}
	filter.SetValueForKey(img, "inputImage")
	num := foundation.NewNumberWithDouble(*sharpness)
	filter.SetValueForKey(num, "inputSharpness")

	outputImage := filter.OutputImage()
	if outputImage == nil {
		return fmt.Errorf("filter produced no output")
	}

	// Crop to original extent since filtering may change bounds.
	extent := img.Extent()
	result := outputImage.ImageByCroppingToRect(extent)

	return saveCIImage(result, *output)
}

func runResize(args []string) error {
	fs := flag.NewFlagSet("resize", flag.ExitOnError)
	scale := fs.Float64("scale", 0.5, "scale factor")
	output := fs.String("o", "", "output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: imgfilter resize [-scale N] [-o output] <image>")
	}

	img, cleanup, err := loadCIImage(fs.Arg(0))
	if err != nil {
		return err
	}
	if cleanup != nil {
		defer cleanup()
	}

	transform := corefoundation.CGAffineTransform{
		A: *scale, D: *scale,
	}
	resized := img.ImageByApplyingTransform(transform)
	return saveCIImage(resized, *output)
}

func runCrop(args []string) error {
	fs := flag.NewFlagSet("crop", flag.ExitOnError)
	x := fs.Float64("x", 0, "crop origin x")
	y := fs.Float64("y", 0, "crop origin y")
	w := fs.Float64("w", 0, "crop width (0 = image width)")
	h := fs.Float64("h", 0, "crop height (0 = image height)")
	output := fs.String("o", "", "output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: imgfilter crop [-x N] [-y N] [-w N] [-h N] [-o output] <image>")
	}

	img, cleanup, err := loadCIImage(fs.Arg(0))
	if err != nil {
		return err
	}
	if cleanup != nil {
		defer cleanup()
	}

	extent := img.Extent()
	cw, ch := *w, *h
	if cw == 0 {
		cw = extent.Size.Width
	}
	if ch == 0 {
		ch = extent.Size.Height
	}

	rect := corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: *x, Y: *y},
		Size:   corefoundation.CGSize{Width: cw, Height: ch},
	}
	cropped := img.ImageByCroppingToRect(rect)
	return saveCIImage(cropped, *output)
}

func runGrayscale(args []string) error {
	fs := flag.NewFlagSet("grayscale", flag.ExitOnError)
	output := fs.String("o", "", "output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: imgfilter grayscale [-o output] <image>")
	}

	img, cleanup, err := loadCIImage(fs.Arg(0))
	if err != nil {
		return err
	}
	if cleanup != nil {
		defer cleanup()
	}

	gray := img.ImageByApplyingFilter("CIPhotoEffectMono")
	return saveCIImage(gray, *output)
}

func runQRGen(args []string) error {
	fs := flag.NewFlagSet("qrgen", flag.ExitOnError)
	output := fs.String("o", "", "output file (default: stdout)")
	scale := fs.Float64("scale", 10, "scale factor for QR code")
	fs.Parse(args)

	var text string
	if fs.NArg() > 0 && fs.Arg(0) != "-" {
		text = fs.Arg(0)
	} else {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("read stdin: %w", err)
		}
		text = string(data)
	}
	if text == "" {
		return fmt.Errorf("no text provided")
	}

	filter := coreimage.NewFilterWithName("CIQRCodeGenerator")
	if filter.ID == 0 {
		return fmt.Errorf("CIQRCodeGenerator filter not available")
	}

	// Set message data (NSUTF8StringEncoding = 4).
	msgStr := foundation.GetNSStringClass().StringWithString(text)
	dataObj := msgStr.DataUsingEncoding(4)
	filter.SetValueForKey(dataObj, "inputMessage")

	// Set correction level.
	filter.SetValueForKey(foundation.NewStringWithString("M"), "inputCorrectionLevel")

	outputImage := filter.OutputImage()
	if outputImage == nil {
		return fmt.Errorf("filter produced no output")
	}

	// Scale up the QR code (it's tiny by default).
	transform := corefoundation.CGAffineTransform{
		A: *scale, D: *scale,
	}
	scaled := outputImage.ImageByApplyingTransform(transform)

	return saveCIImage(scaled, *output)
}

// --- Image I/O ---

func loadCIImage(path string) (coreimage.CIImage, func(), error) {
	if path == "-" {
		return loadCIImageStdin()
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return coreimage.CIImage{}, nil, fmt.Errorf("resolve path: %w", err)
	}
	url := foundation.NewURLFileURLWithPath(absPath)
	img := coreimage.NewImageWithContentsOfURL(url)
	if img.ID == 0 {
		return coreimage.CIImage{}, nil, fmt.Errorf("cannot load image: %s", path)
	}
	return img, nil, nil
}

func loadCIImageStdin() (coreimage.CIImage, func(), error) {
	f, err := os.CreateTemp("", "imgfilter-*.png")
	if err != nil {
		return coreimage.CIImage{}, nil, fmt.Errorf("create temp file: %w", err)
	}
	tmpPath := f.Name()

	if _, err := io.Copy(f, os.Stdin); err != nil {
		f.Close()
		os.Remove(tmpPath)
		return coreimage.CIImage{}, nil, fmt.Errorf("read stdin: %w", err)
	}
	f.Close()

	absPath, _ := filepath.Abs(tmpPath)
	url := foundation.NewURLFileURLWithPath(absPath)
	img := coreimage.NewImageWithContentsOfURL(url)
	if img.ID == 0 {
		os.Remove(tmpPath)
		return coreimage.CIImage{}, nil, fmt.Errorf("cannot load image from stdin")
	}
	return img, func() { os.Remove(tmpPath) }, nil
}

func saveCIImage(img coreimage.ICIImage, output string) error {
	ctx := coreimage.NewCIContext()
	colorSpace := coregraphics.CGColorSpaceCreateDeviceRGB()

	if output == "" || output == "-" {
		// Write to stdout as PNG via temp file.
		f, err := os.CreateTemp("", "imgfilter-out-*.png")
		if err != nil {
			return fmt.Errorf("create temp file: %w", err)
		}
		tmpPath := f.Name()
		f.Close()

		url := foundation.NewURLFileURLWithPath(tmpPath)
		// kCIFormatRGBA8 = 24
		ok, err := ctx.WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(img, url, 24, colorSpace, nil)
		if err != nil {
			os.Remove(tmpPath)
			return fmt.Errorf("write png: %w", err)
		}
		if !ok {
			os.Remove(tmpPath)
			return fmt.Errorf("write png failed")
		}

		data, err := os.ReadFile(tmpPath)
		os.Remove(tmpPath)
		if err != nil {
			return fmt.Errorf("read temp: %w", err)
		}
		os.Stdout.Write(data)
		return nil
	}

	absPath, _ := filepath.Abs(output)
	url := foundation.NewURLFileURLWithPath(absPath)

	ext := strings.ToLower(filepath.Ext(output))
	switch ext {
	case ".jpg", ".jpeg":
		ok, err := ctx.WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError(img, url, colorSpace, nil)
		if err != nil {
			return fmt.Errorf("write jpeg: %w", err)
		}
		if !ok {
			return fmt.Errorf("write jpeg failed")
		}
	default:
		ok, err := ctx.WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(img, url, 24, colorSpace, nil)
		if err != nil {
			return fmt.Errorf("write png: %w", err)
		}
		if !ok {
			return fmt.Errorf("write png failed")
		}
	}

	return nil
}
