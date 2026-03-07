// Command ocr extracts text from an image file using the Vision framework.
//
// Usage: ocr <image-file>
package main

import (
	"fmt"
	"os"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/vision"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ocr <image-file>\n")
		os.Exit(1)
	}
	path := os.Args[1]
	url := foundation.NewURLFileURLWithPath(path)

	handler := vision.NewImageRequestHandlerWithURLOptions(url, nil)
	request := vision.NewVNRecognizeTextRequest()

	ok, err := handler.PerformRequestsError([]vision.VNRequest{request.VNRequest})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if !ok {
		fmt.Fprintf(os.Stderr, "request failed\n")
		os.Exit(1)
	}

	results := request.Results()
	for _, obs := range results {
		textObs := vision.VNRecognizedTextObservationFromID(obs.ID)
		candidates := textObs.TopCandidates(1)
		for _, c := range candidates {
			fmt.Printf("[%.2f] %s\n", c.Confidence(), c.String())
		}
	}
}
