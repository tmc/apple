package main

import (
	"math"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
)

// canvasSize is the side of the square drawing surface, in pixels. It is a
// whole multiple of imageSize so that downsampling is an exact box filter.
const canvasSize = 280

// brushRadius is the stroke half-width in canvas pixels. MNIST digits are
// drawn with a stroke roughly 2 of 28 pixels wide, which is this radius once
// the canvas is scaled down.
const brushRadius = 10

// canvas is an 8-bit grayscale drawing surface: 0 is background and 255 is
// full ink, matching MNIST's white-digit-on-black convention.
type canvas struct {
	pixels []byte
	// last stroke point, so drags paint a connected line rather than dots.
	lastX, lastY float64
	drawing      bool
}

func newCanvas() *canvas {
	return &canvas{pixels: make([]byte, canvasSize*canvasSize)}
}

// clear erases the canvas.
func (c *canvas) clear() {
	for i := range c.pixels {
		c.pixels[i] = 0
	}
	c.drawing = false
}

// isEmpty reports whether nothing has been drawn.
func (c *canvas) isEmpty() bool {
	for _, p := range c.pixels {
		if p != 0 {
			return false
		}
	}
	return true
}

// beginStroke starts a new stroke at (x, y) in canvas coordinates.
func (c *canvas) beginStroke(x, y float64) {
	c.lastX, c.lastY = x, y
	c.drawing = true
	c.stamp(x, y)
}

// extendStroke paints from the previous point to (x, y). Points are
// interpolated because mouse events arrive far apart during a fast drag.
func (c *canvas) extendStroke(x, y float64) {
	if !c.drawing {
		c.beginStroke(x, y)
		return
	}
	dx, dy := x-c.lastX, y-c.lastY
	steps := int(max(math.Abs(dx), math.Abs(dy))) + 1
	for i := 1; i <= steps; i++ {
		t := float64(i) / float64(steps)
		c.stamp(c.lastX+dx*t, c.lastY+dy*t)
	}
	c.lastX, c.lastY = x, y
}

// endStroke lifts the pen.
func (c *canvas) endStroke() { c.drawing = false }

// stamp paints one soft round dab centered at (x, y). The edge falls off
// linearly over the last pixel so downsampled strokes are not jagged.
func (c *canvas) stamp(x, y float64) {
	minX := clampInt(int(x)-brushRadius-1, 0, canvasSize-1)
	maxX := clampInt(int(x)+brushRadius+1, 0, canvasSize-1)
	minY := clampInt(int(y)-brushRadius-1, 0, canvasSize-1)
	maxY := clampInt(int(y)+brushRadius+1, 0, canvasSize-1)

	for py := minY; py <= maxY; py++ {
		for px := minX; px <= maxX; px++ {
			d := math.Hypot(float64(px)-x, float64(py)-y)
			if d > brushRadius {
				continue
			}
			v := 255.0
			if edge := float64(brushRadius) - d; edge < 1 {
				v = 255 * edge
			}
			i := py*canvasSize + px
			if b := byte(v); b > c.pixels[i] {
				c.pixels[i] = b
			}
		}
	}
}

// paint replaces the canvas with an imageSize x imageSize input, scaled up to
// canvas resolution. It exists so the self test can drive the same path a
// mouse stroke does.
func (c *canvas) paint(src []float32) {
	const scale = canvasSize / imageSize
	for y := 0; y < canvasSize; y++ {
		for x := 0; x < canvasSize; x++ {
			c.pixels[y*canvasSize+x] = byte(255 * src[(y/scale)*imageSize+(x/scale)])
		}
	}
	c.drawing = false
}

// image returns the canvas as an NSImage suitable for an NSImageView.
func (c *canvas) image() appkit.NSImage {
	space := coregraphics.CGColorSpaceCreateDeviceGray()
	defer coregraphics.CGColorSpaceRelease(space)

	// kCGImageAlphaNone is 0: one 8-bit gray component per pixel.
	ctx := coregraphics.CGBitmapContextCreate(
		unsafe.Pointer(&c.pixels[0]), canvasSize, canvasSize, 8, canvasSize, space, coregraphics.CGBitmapInfo(0))
	if ctx == 0 {
		return appkit.NSImage{}
	}
	defer coregraphics.CGContextRelease(ctx)

	cg := coregraphics.CGBitmapContextCreateImage(ctx)
	if cg == 0 {
		return appkit.NSImage{}
	}
	defer coregraphics.CGImageRelease(cg)

	return appkit.NewImageWithCGImageSize(cg, corefoundation.CGSize{Width: canvasSize, Height: canvasSize})
}

// normalized returns the canvas as the imageSize*imageSize input the network
// expects, with values in [0, 1].
//
// MNIST is not raw pixels: each digit was cropped to its bounding box, scaled
// so the longer side is 20 pixels, and placed in a 28x28 field centered on its
// center of mass. Skipping this step makes hand-drawn input land far outside
// the training distribution, so the classifier answers confidently and wrongly.
func (c *canvas) normalized() []float32 {
	out := make([]float32, imageSize*imageSize)

	minX, minY, maxX, maxY := canvasSize, canvasSize, -1, -1
	for y := 0; y < canvasSize; y++ {
		for x := 0; x < canvasSize; x++ {
			if c.pixels[y*canvasSize+x] == 0 {
				continue
			}
			minX, maxX = min(minX, x), max(maxX, x)
			minY, maxY = min(minY, y), max(maxY, y)
		}
	}
	if maxX < 0 {
		return out // nothing drawn
	}

	// Scale the bounding box into a 20x20 box, preserving aspect ratio.
	const fit = 20
	boxW, boxH := maxX-minX+1, maxY-minY+1
	scale := float64(fit) / float64(max(boxW, boxH))
	dstW, dstH := max(1, int(float64(boxW)*scale)), max(1, int(float64(boxH)*scale))

	// Box-filter the cropped region down to dstW x dstH.
	small := make([]float64, dstW*dstH)
	for dy := 0; dy < dstH; dy++ {
		for dx := 0; dx < dstW; dx++ {
			x0 := minX + dx*boxW/dstW
			x1 := max(x0+1, minX+(dx+1)*boxW/dstW)
			y0 := minY + dy*boxH/dstH
			y1 := max(y0+1, minY+(dy+1)*boxH/dstH)
			sum, n := 0.0, 0
			for y := y0; y < y1 && y < canvasSize; y++ {
				for x := x0; x < x1 && x < canvasSize; x++ {
					sum += float64(c.pixels[y*canvasSize+x])
					n++
				}
			}
			if n > 0 {
				small[dy*dstW+dx] = sum / float64(n) / 255
			}
		}
	}

	// Center the scaled digit on its center of mass.
	var mass, cx, cy float64
	for y := 0; y < dstH; y++ {
		for x := 0; x < dstW; x++ {
			v := small[y*dstW+x]
			mass += v
			cx += v * float64(x)
			cy += v * float64(y)
		}
	}
	offX, offY := (imageSize-dstW)/2, (imageSize-dstH)/2
	if mass > 0 {
		offX = clampInt(imageSize/2-int(cx/mass), 0, imageSize-dstW)
		offY = clampInt(imageSize/2-int(cy/mass), 0, imageSize-dstH)
	}
	for y := 0; y < dstH; y++ {
		for x := 0; x < dstW; x++ {
			out[(y+offY)*imageSize+(x+offX)] = float32(small[y*dstW+x])
		}
	}
	return out
}

func clampInt(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
