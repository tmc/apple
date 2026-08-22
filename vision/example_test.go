//go:build darwin

package vision_test

import (
	"fmt"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/vision"
)

func ExampleVNRecognizeTextRequest() {
	req := vision.NewVNRecognizeTextRequest()
	req.SetRecognitionLevel(vision.VNRequestTextRecognitionLevelAccurate)
	req.SetUsesLanguageCorrection(true)
	req.SetMinimumTextHeight(0.05)
	req.SetCustomWords([]string{"Apple", "Vision"})

	fmt.Println("RecognitionLevel:", req.RecognitionLevel())
	fmt.Println("UsesLanguageCorrection:", req.UsesLanguageCorrection())
	fmt.Printf("MinimumTextHeight: %.2f\n", req.MinimumTextHeight())
	fmt.Println("CustomWords:", req.CustomWords())

	// Output:
	// RecognitionLevel: VNRequestTextRecognitionLevelAccurate
	// UsesLanguageCorrection: true
	// MinimumTextHeight: 0.05
	// CustomWords: [Apple Vision]
}

func ExampleVNDetectFaceRectanglesRequest() {
	req := vision.NewVNDetectFaceRectanglesRequest()
	req.SetRevision(1)
	req.SetRegionOfInterest(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 0.1, Y: 0.1},
		Size:   corefoundation.CGSize{Width: 0.8, Height: 0.8},
	})

	fmt.Println("Revision:", req.Revision())
	roi := req.RegionOfInterest()
	fmt.Printf("RegionOfInterest: (%.1f, %.1f, %.1f, %.1f)\n", roi.Origin.X, roi.Origin.Y, roi.Size.Width, roi.Size.Height)

	// Output:
	// Revision: 1
	// RegionOfInterest: (0.1, 0.1, 0.8, 0.8)
}
