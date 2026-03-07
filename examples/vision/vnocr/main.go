// Command vnocr extracts text from images using the Vision framework.
//
// Usage:
//
//	vnocr [flags] <image-file>
//	cat image.png | vnocr [flags] -
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/vision"
)

type textResult struct {
	Text       string  `json:"text"`
	Confidence float32 `json:"confidence"`
	BBox       bbox    `json:"bbox"`
}

type bbox struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

func main() {
	jsonOut := flag.Bool("j", false, "JSON output")
	level := flag.String("level", "accurate", "recognition level: fast or accurate")
	lang := flag.String("lang", "", "recognition language hint (e.g. en)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: vnocr [flags] <image-file>\n")
		fmt.Fprintf(os.Stderr, "       cat image.png | vnocr [flags] -\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(2)
	}

	handler, err := makeHandler(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	request := vision.NewVNRecognizeTextRequest()
	switch *level {
	case "fast":
		request.SetRecognitionLevel(vision.VNRequestTextRecognitionLevelFast)
	case "accurate":
		request.SetRecognitionLevel(vision.VNRequestTextRecognitionLevelAccurate)
	default:
		fmt.Fprintf(os.Stderr, "error: unknown level %q (use fast or accurate)\n", *level)
		os.Exit(2)
	}
	if *lang != "" {
		request.SetRecognitionLanguages([]string{*lang})
	}

	ok, perr := handler.PerformRequestsError([]vision.VNRequest{request.VNImageBasedRequest.VNRequest})
	if perr != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", perr)
		os.Exit(2)
	}
	if !ok {
		fmt.Fprintf(os.Stderr, "request failed\n")
		os.Exit(2)
	}

	results := request.Results()
	if len(results) == 0 {
		os.Exit(1)
	}

	var items []textResult
	for _, obs := range results {
		textObs := vision.VNRecognizedTextObservationFromID(obs.ID)
		bb := textObs.BoundingBox()
		candidates := textObs.TopCandidates(1)
		for _, c := range candidates {
			items = append(items, textResult{
				Text:       c.String(),
				Confidence: float32(c.Confidence()),
				BBox: bbox{
					X: bb.Origin.X,
					Y: bb.Origin.Y,
					W: bb.Size.Width,
					H: bb.Size.Height,
				},
			})
		}
	}

	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(items); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(2)
		}
		return
	}

	for _, item := range items {
		fmt.Println(item.Text)
	}
}

func makeHandler(path string) (vision.VNImageRequestHandler, error) {
	if path == "-" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return vision.VNImageRequestHandler{}, fmt.Errorf("reading stdin: %w", err)
		}
		nsData := foundation.NewDataWithBytesLength(data)
		return vision.NewImageRequestHandlerWithDataOptions(nsData, nil), nil
	}
	url := foundation.NewURLFileURLWithPath(path)
	return vision.NewImageRequestHandlerWithURLOptions(url, nil), nil
}
