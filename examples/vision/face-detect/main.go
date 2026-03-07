// Face-detect uses the Vision framework to detect faces in an image.
//
// Usage: face-detect <image-path>
package main

import (
	"fmt"
	"os"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/vision"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <image-path>\n", os.Args[0])
		os.Exit(1)
	}

	imageURL := foundation.NewURLFileURLWithPath(os.Args[1])
	handler := vision.NewImageRequestHandlerWithURLOptions(imageURL, nil)
	request := vision.NewVNDetectFaceRectanglesRequest()

	ok, err := handler.PerformRequestsError([]vision.VNRequest{request.VNImageBasedRequest.VNRequest})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	if !ok {
		fmt.Fprintf(os.Stderr, "request failed\n")
		os.Exit(1)
	}

	observations := request.Results()
	fmt.Printf("Detected %d face(s)\n", len(observations))
	for i, obs := range observations {
		face := vision.VNFaceObservationFromID(obs.ID)
		bb := face.BoundingBox()
		conf := face.Confidence()
		fmt.Printf("  Face %d: origin=(%.3f, %.3f) size=(%.3f x %.3f) confidence=%.2f\n",
			i+1, bb.Origin.X, bb.Origin.Y, bb.Size.Width, bb.Size.Height, conf)
	}
}
