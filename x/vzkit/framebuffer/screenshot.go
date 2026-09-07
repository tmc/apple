package framebuffer

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
)

// CaptureImage takes a screenshot of the display and returns it as an
// image.Image. It drives -[VZGraphicsDisplay _takeScreenshotWithCompletionHandler:],
// whose real completion ABI is void(^)(CGImageRef, NSError*), and converts the
// returned CGImage into an *image.RGBA.
//
// The VM must be running and the display must have painted at least one frame;
// otherwise the framework reports an error through the completion handler.
func (d Display) CaptureImage(ctx context.Context) (image.Image, error) {
	type result struct {
		img uintptr
		err error
	}
	done := make(chan result, 1)
	cleanup, err := d.raw.TakeScreenshotWithImageCompletionHandler(func(img uintptr, err error) {
		// Retain the CGImage so it outlives the completion block; the caller
		// releases it after conversion.
		if img != 0 {
			coregraphics.CGImageRetain(coregraphics.CGImageRef(img))
		}
		done <- result{img: img, err: err}
	})
	if err != nil {
		return nil, err
	}
	select {
	case r := <-done:
		cleanup()
		if r.err != nil {
			if r.img != 0 {
				coregraphics.CGImageRelease(coregraphics.CGImageRef(r.img))
			}
			return nil, r.err
		}
		if r.img == 0 {
			return nil, errors.New("screenshot returned a nil image")
		}
		defer coregraphics.CGImageRelease(coregraphics.CGImageRef(r.img))
		return cgImageToRGBA(coregraphics.CGImageRef(r.img))
	case <-ctx.Done():
		// The block may still fire; leak it rather than release under a race.
		return nil, ctx.Err()
	}
}

// CapturePNG captures the display and writes it to path as a PNG.
func (d Display) CapturePNG(ctx context.Context, path string) error {
	img, err := d.CaptureImage(ctx)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return err
	}
	return f.Close()
}

// cgImageToRGBA draws a CGImage into a bitmap context backed by an image.RGBA
// and returns the result. The bitmap is 8 bits per component, RGBA byte order,
// premultiplied alpha, which matches image.RGBA exactly.
func cgImageToRGBA(img coregraphics.CGImageRef) (*image.RGBA, error) {
	w := int(coregraphics.CGImageGetWidth(img))
	h := int(coregraphics.CGImageGetHeight(img))
	if w <= 0 || h <= 0 {
		return nil, fmt.Errorf("screenshot image has zero size %dx%d", w, h)
	}
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))

	space := coregraphics.CGColorSpaceCreateDeviceRGB()
	if space == 0 {
		return nil, errors.New("create device RGB color space")
	}
	defer coregraphics.CGColorSpaceRelease(space)

	bitmapInfo := coregraphics.CGBitmapInfo(coregraphics.KCGImageAlphaPremultipliedLast) | coregraphics.KCGBitmapByteOrder32Big
	ctx := coregraphics.CGBitmapContextCreate(
		unsafe.Pointer(&rgba.Pix[0]),
		uintptr(w), uintptr(h),
		8, uintptr(rgba.Stride),
		space, bitmapInfo,
	)
	if ctx == 0 {
		return nil, errors.New("create bitmap context")
	}
	defer coregraphics.CGContextRelease(ctx)

	coregraphics.CGContextDrawImage(ctx, corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 0, Y: 0},
		Size:   corefoundation.CGSize{Width: float64(w), Height: float64(h)},
	}, img)
	return rgba, nil
}
