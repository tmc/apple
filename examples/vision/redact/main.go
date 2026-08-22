// Command redact detects faces and text in an image with the Vision framework
// and writes a PNG copy with those regions obscured using Core Image.
//
// Usage: redact [-mode pixelate|box] [-scale n] [-pad f] [-faces] [-text] <input-image> <output.png>
//
// By default both faces and text are redacted. Passing -faces or -text limits
// redaction to that kind of region.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/vision"
)

func main() {
	mode := flag.String("mode", "pixelate", "redaction mode: pixelate or box")
	scale := flag.Float64("scale", 24, "pixel block size for -mode pixelate")
	pad := flag.Float64("pad", 0.02, "extra margin around each region, as a fraction of image size")
	faces := flag.Bool("faces", false, "redact faces only")
	text := flag.Bool("text", false, "redact text only")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: redact [flags] <input-image> <output.png>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}
	if *mode != "pixelate" && *mode != "box" {
		fmt.Fprintf(os.Stderr, "redact: unknown mode %q\n", *mode)
		os.Exit(1)
	}
	if !*faces && !*text {
		*faces, *text = true, true
	}

	if err := run(flag.Arg(0), flag.Arg(1), *mode, *scale, *pad, *faces, *text); err != nil {
		fmt.Fprintf(os.Stderr, "redact: %v\n", err)
		os.Exit(1)
	}
}

func run(inPath, outPath, mode string, scale, pad float64, wantFaces, wantText bool) error {
	inURL := foundation.NewURLFileURLWithPath(inPath)

	base := coreimage.NewImageWithContentsOfURL(inURL)
	if base.ID == 0 {
		return fmt.Errorf("cannot read image %s", inPath)
	}
	extent := base.Extent()
	if extent.Size.Width <= 0 || extent.Size.Height <= 0 {
		return fmt.Errorf("image %s has empty extent", inPath)
	}

	regions, err := detect(inURL, wantFaces, wantText)
	if err != nil {
		return err
	}
	if len(regions) == 0 {
		fmt.Fprintf(os.Stderr, "redact: no regions detected; writing an unmodified copy\n")
	}

	// Vision reports normalized rects with a bottom-left origin, which matches
	// the Core Image coordinate space, so only a scale to pixels is needed.
	var cover coreimage.ICIImage
	switch mode {
	case "box":
		cover = coreimage.NewImageWithColor(coreimage.NewColorWithRedGreenBlue(0, 0, 0))
	case "pixelate":
		params := foundation.NewDictionaryWithObjectForKey(
			foundation.NewNumberWithDouble(scale),
			foundation.NewStringWithString("inputScale"),
		)
		cover = base.ImageByApplyingFilterWithInputParameters("CIPixellate", params)
	}

	out := coreimage.ICIImage(base)
	for _, r := range regions {
		rect := denormalize(r, extent, pad)
		out = cover.ImageByCroppingToRect(rect).ImageByCompositingOverImage(out)
	}
	out = out.ImageByCroppingToRect(extent)

	ctx := coreimage.NewCIContext()
	colorSpace := base.ColorSpace()
	if colorSpace == 0 {
		colorSpace = coregraphics.CGColorSpaceCreateDeviceRGB()
	}
	outURL := foundation.NewURLFileURLWithPath(outPath)
	if _, err := ctx.WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(
		out, outURL, coreimage.CIFormat(coreimage.KCIFormatRGBA8), colorSpace, nil); err != nil {
		return fmt.Errorf("writing %s: %w", outPath, err)
	}

	fmt.Printf("redacted %d region(s) -> %s\n", len(regions), outPath)
	return nil
}

// detect returns the normalized bounding boxes of the requested region kinds.
func detect(imageURL foundation.NSURL, wantFaces, wantText bool) ([]corefoundation.CGRect, error) {
	handler := vision.NewImageRequestHandlerWithURLOptions(imageURL, nil)

	var requests []vision.VNRequest
	faceReq := vision.NewVNDetectFaceRectanglesRequest()
	textReq := vision.NewVNRecognizeTextRequest()
	if wantFaces {
		requests = append(requests, faceReq.VNImageBasedRequest.VNRequest)
	}
	if wantText {
		requests = append(requests, textReq.VNImageBasedRequest.VNRequest)
	}

	ok, err := handler.PerformRequestsError(requests)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("vision request failed")
	}

	var rects []corefoundation.CGRect
	if wantFaces {
		for _, obs := range faceReq.Results() {
			rects = append(rects, vision.VNFaceObservationFromID(obs.ID).BoundingBox())
		}
	}
	if wantText {
		for _, obs := range textReq.Results() {
			rects = append(rects, vision.VNRecognizedTextObservationFromID(obs.ID).BoundingBox())
		}
	}
	return rects, nil
}

// denormalize converts a normalized Vision rect to pixel coordinates within
// extent, grown by pad (a fraction of the image size) and clipped to extent.
func denormalize(r, extent corefoundation.CGRect, pad float64) corefoundation.CGRect {
	w, h := extent.Size.Width, extent.Size.Height
	x0 := extent.Origin.X + r.Origin.X*w - pad*w
	y0 := extent.Origin.Y + r.Origin.Y*h - pad*h
	x1 := extent.Origin.X + (r.Origin.X+r.Size.Width)*w + pad*w
	y1 := extent.Origin.Y + (r.Origin.Y+r.Size.Height)*h + pad*h

	x0 = clamp(x0, extent.Origin.X, extent.Origin.X+w)
	x1 = clamp(x1, extent.Origin.X, extent.Origin.X+w)
	y0 = clamp(y0, extent.Origin.Y, extent.Origin.Y+h)
	y1 = clamp(y1, extent.Origin.Y, extent.Origin.Y+h)

	return corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: x0, Y: y0},
		Size:   corefoundation.CGSize{Width: x1 - x0, Height: y1 - y0},
	}
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
