package ocr

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"time"
)

// DebugEntry represents a single OCR debug log entry.
type DebugEntry struct {
	Timestamp    string            `json:"timestamp"`
	Step         string            `json:"step,omitempty"`
	Observations []TextObservation `json:"observations"`
	ScreenState  string            `json:"screen_state,omitempty"`
}

// SaveDebugScreenshot saves a screenshot with OCR bounding boxes overlaid.
// It writes both a PNG image and a JSON file with the observations.
func SaveDebugScreenshot(img image.Image, observations []TextObservation, dir, prefix string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create debug dir: %w", err)
	}

	ts := time.Now().Format("20060102-150405")
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	overlay := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			overlay.Set(x, y, img.At(x+bounds.Min.X, y+bounds.Min.Y))
		}
	}

	green := color.RGBA{R: 0, G: 255, B: 0, A: 200}
	for _, obs := range observations {
		x1 := int(obs.BoundingBox.Origin.X * float64(w))
		y1 := int((1 - obs.BoundingBox.Origin.Y - obs.BoundingBox.Size.Height) * float64(h))
		x2 := int((obs.BoundingBox.Origin.X + obs.BoundingBox.Size.Width) * float64(w))
		y2 := int((1 - obs.BoundingBox.Origin.Y) * float64(h))

		DrawRect(overlay, x1, y1, x2, y2, green)
	}

	imgPath := filepath.Join(dir, fmt.Sprintf("%s-%s.png", prefix, ts))
	f, err := os.Create(imgPath)
	if err != nil {
		return fmt.Errorf("create debug image: %w", err)
	}
	defer f.Close()
	if err := png.Encode(f, overlay); err != nil {
		return fmt.Errorf("encode debug image: %w", err)
	}

	entry := DebugEntry{
		Timestamp:    ts,
		Step:         prefix,
		Observations: observations,
	}
	jsonPath := filepath.Join(dir, fmt.Sprintf("%s-%s.json", prefix, ts))
	jsonData, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal debug json: %w", err)
	}
	if err := os.WriteFile(jsonPath, jsonData, 0644); err != nil {
		return fmt.Errorf("write debug json: %w", err)
	}

	return nil
}

// DrawRect draws a 2px rectangle outline on an RGBA image.
func DrawRect(img *image.RGBA, x1, y1, x2, y2 int, c color.RGBA) {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	if x1 < 0 {
		x1 = 0
	}
	if y1 < 0 {
		y1 = 0
	}
	if x2 >= w {
		x2 = w - 1
	}
	if y2 >= h {
		y2 = h - 1
	}

	for x := x1; x <= x2; x++ {
		img.SetRGBA(x, y1, c)
		img.SetRGBA(x, y2, c)
		if y1+1 < h {
			img.SetRGBA(x, y1+1, c)
		}
		if y2-1 >= 0 {
			img.SetRGBA(x, y2-1, c)
		}
	}
	for y := y1; y <= y2; y++ {
		img.SetRGBA(x1, y, c)
		img.SetRGBA(x2, y, c)
		if x1+1 < w {
			img.SetRGBA(x1+1, y, c)
		}
		if x2-1 >= 0 {
			img.SetRGBA(x2-1, y, c)
		}
	}
}
