package ocr

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"sort"
	"strings"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/vision"
)

// TextObservation holds a recognized text region from OCR.
type TextObservation struct {
	Text        string
	Confidence  float32
	BoundingBox corefoundation.CGRect // normalized coordinates (0-1), origin at bottom-left
	Center      image.Point           // center in screen pixel coordinates
}

// Service performs text recognition using Apple's Vision framework.
type Service struct {
	verbose bool
}

// NewService creates a new OCR service.
func NewService(verbose bool) *Service {
	return &Service{verbose: verbose}
}

// RecognizeText runs OCR on an image and returns all recognized text observations.
func (s *Service) RecognizeText(img image.Image) ([]TextObservation, error) {
	if img == nil {
		return nil, fmt.Errorf("nil image")
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("encode png: %w", err)
	}
	data := foundation.NewDataFromBytes(buf.Bytes())
	if data.ID == 0 {
		return nil, fmt.Errorf("create NSData from image bytes")
	}

	return s.recognizeTextInData(data, img.Bounds().Dx(), img.Bounds().Dy())
}

// FindText searches for text on screen and returns its center pixel coordinates.
// When multiple observations contain the needle, it prefers:
//  1. Exact full-text matches (observation text equals needle, case-insensitive)
//  2. Shorter observation text (closer to the needle itself)
//  3. Observations closer to the bottom of the screen
//
// Returns found=false if the text is not visible.
func (s *Service) FindText(img image.Image, needle string) (x, y float64, found bool) {
	return s.FindTextWithOptions(img, needle, SearchOptions{})
}

// FindTextWithOptions searches for text with optional region and ranking preferences.
func (s *Service) FindTextWithOptions(img image.Image, needle string, opts SearchOptions) (x, y float64, found bool) {
	observations, err := s.RecognizeText(img)
	if err != nil {
		if s.verbose {
			fmt.Printf("[ocr] error: %v\n", err)
		}
		return 0, 0, false
	}

	best, ok := BestMatch(observations, needle, opts, img.Bounds())
	if !ok {
		return 0, 0, false
	}
	if s.verbose {
		fmt.Printf("[ocr] found %q in %q at (%d, %d)\n", needle, best.Text, best.Center.X, best.Center.Y)
	}
	return float64(best.Center.X), float64(best.Center.Y), true
}

// FindTextNormalized searches for text and returns its center in normalized
// coordinates (0-1, top-left origin) suitable for mouse input commands.
func (s *Service) FindTextNormalized(img image.Image, needle string) (x, y float64, found bool) {
	return s.FindTextNormalizedWithOptions(img, needle, SearchOptions{})
}

// FindTextNormalizedWithOptions searches for text with options and returns the
// center in normalized coordinates (0-1, top-left origin).
func (s *Service) FindTextNormalizedWithOptions(img image.Image, needle string, opts SearchOptions) (x, y float64, found bool) {
	px, py, found := s.FindTextWithOptions(img, needle, opts)
	if !found {
		return 0, 0, false
	}
	bounds := img.Bounds()
	return px / float64(bounds.Dx()), py / float64(bounds.Dy()), true
}

// AllText returns all recognized text as a single string with lines separated by newlines.
func (s *Service) AllText(img image.Image) string {
	observations, err := s.RecognizeText(img)
	if err != nil {
		return ""
	}
	var lines []string
	for _, obs := range observations {
		lines = append(lines, obs.Text)
	}
	return strings.Join(lines, "\n")
}

// BestMatch finds the best matching observation for needle among the given
// observations, applying search options and ranking heuristics.
func BestMatch(observations []TextObservation, needle string, opts SearchOptions, bounds image.Rectangle) (TextObservation, bool) {
	needle = strings.ToLower(needle)

	var matches []TextObservation
	for _, obs := range observations {
		if !observationInRegion(obs, bounds, opts.Region) {
			continue
		}
		if strings.Contains(strings.ToLower(obs.Text), needle) {
			matches = append(matches, obs)
		}
	}
	if len(matches) == 0 {
		return TextObservation{}, false
	}
	if len(matches) == 1 {
		return matches[0], true
	}

	sort.Slice(matches, func(i, j int) bool {
		ti := strings.ToLower(matches[i].Text)
		tj := strings.ToLower(matches[j].Text)
		exactI := ti == needle
		exactJ := tj == needle
		if exactI != exactJ {
			return exactI
		}
		if len(ti) != len(tj) {
			return len(ti) < len(tj)
		}
		if opts.PreferTop {
			if matches[i].Center.Y != matches[j].Center.Y {
				return matches[i].Center.Y < matches[j].Center.Y
			}
		} else if matches[i].Center.Y != matches[j].Center.Y {
			return matches[i].Center.Y > matches[j].Center.Y
		}
		return matches[i].Center.X < matches[j].Center.X
	})
	return matches[0], true
}

func (s *Service) recognizeTextInData(data foundation.INSData, width, height int) ([]TextObservation, error) {
	handler := vision.NewImageRequestHandlerWithDataOptions(data, nil)

	request := vision.NewVNRecognizeTextRequest()
	request.SetRecognitionLevel(vision.VNRequestTextRecognitionLevelAccurate)
	request.SetUsesLanguageCorrection(true)

	ok, err := handler.PerformRequestsError([]vision.VNRequest{
		vision.VNRequestFromID(request.ID),
	})
	if err != nil {
		return nil, fmt.Errorf("perform requests: %w", err)
	}
	if !ok {
		return nil, fmt.Errorf("perform requests returned false")
	}

	return s.extractResults(request, width, height), nil
}

func (s *Service) extractResults(request vision.VNRecognizeTextRequest, width, height int) []TextObservation {
	observations := request.Results()
	var results []TextObservation
	for _, obs := range observations {
		textObs := vision.VNRecognizedTextObservationFromID(obs.ID)
		candidates := textObs.TopCandidates(1)
		if len(candidates) == 0 {
			continue
		}
		candidate := candidates[0]
		bbox := textObs.BoundingBox()

		centerX := int((bbox.Origin.X + bbox.Size.Width/2) * float64(width))
		centerY := int((1 - bbox.Origin.Y - bbox.Size.Height/2) * float64(height))

		result := TextObservation{
			Text:        candidate.String(),
			Confidence:  float32(candidate.Confidence()),
			BoundingBox: bbox,
			Center:      image.Point{X: centerX, Y: centerY},
		}
		results = append(results, result)

		if s.verbose {
			fmt.Printf("[ocr] [%.2f] %q at (%d,%d)\n", result.Confidence, result.Text, centerX, centerY)
		}
	}
	return results
}
