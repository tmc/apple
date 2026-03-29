package ocr

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

// Region describes a normalized screen rectangle (0-1, top-left origin).
type Region struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

// SearchOptions controls OCR text matching behavior.
type SearchOptions struct {
	Region    *Region
	PreferTop bool
}

var menuBarRegion = Region{
	MinX: 0,
	MinY: 0,
	MaxX: 1,
	MaxY: 0.30,
}

// MenuSearchOptions returns options tuned for menu bar targeting.
func MenuSearchOptions() SearchOptions {
	region := menuBarRegion
	return SearchOptions{
		Region:    &region,
		PreferTop: true,
	}
}

// ParseSearchOptions parses a region selector for OCR commands.
// Supported selectors:
//   - "" / "screen" / "full": whole screen
//   - "menu" / "menubar": top menu bar strip
//   - "x1,y1,x2,y2": normalized rectangle coordinates
func ParseSearchOptions(regionSpec string) (SearchOptions, error) {
	spec := strings.TrimSpace(strings.ToLower(regionSpec))
	switch spec {
	case "", "screen", "full":
		return SearchOptions{}, nil
	case "menu", "menubar", "top-menu":
		return MenuSearchOptions(), nil
	}

	parts := strings.Split(spec, ",")
	if len(parts) != 4 {
		return SearchOptions{}, fmt.Errorf("invalid region %q (want menu or x1,y1,x2,y2)", regionSpec)
	}

	values := make([]float64, 0, 4)
	for _, part := range parts {
		v, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
		if err != nil {
			return SearchOptions{}, fmt.Errorf("invalid region %q: %w", regionSpec, err)
		}
		values = append(values, v)
	}

	region := Region{
		MinX: values[0],
		MinY: values[1],
		MaxX: values[2],
		MaxY: values[3],
	}
	if err := validateRegion(region); err != nil {
		return SearchOptions{}, fmt.Errorf("invalid region %q: %w", regionSpec, err)
	}
	return SearchOptions{Region: &region}, nil
}

func validateRegion(region Region) error {
	if region.MinX < 0 || region.MinX > 1 || region.MaxX < 0 || region.MaxX > 1 {
		return fmt.Errorf("x coordinates must be within [0,1]")
	}
	if region.MinY < 0 || region.MinY > 1 || region.MaxY < 0 || region.MaxY > 1 {
		return fmt.Errorf("y coordinates must be within [0,1]")
	}
	if region.MinX >= region.MaxX {
		return fmt.Errorf("x1 must be less than x2")
	}
	if region.MinY >= region.MaxY {
		return fmt.Errorf("y1 must be less than y2")
	}
	return nil
}

func observationInRegion(obs TextObservation, bounds image.Rectangle, region *Region) bool {
	if region == nil {
		return true
	}
	width := bounds.Dx()
	height := bounds.Dy()
	if width == 0 || height == 0 {
		return false
	}
	x := float64(obs.Center.X-bounds.Min.X) / float64(width)
	y := float64(obs.Center.Y-bounds.Min.Y) / float64(height)
	return x >= region.MinX && x <= region.MaxX && y >= region.MinY && y <= region.MaxY
}
