// Command vndetect detects objects in images using the Vision framework.
//
// Usage:
//
//	vndetect [flags] <subcommand> <image-file>
//	cat image.png | vndetect [flags] <subcommand> -
//
// Subcommands: faces, barcodes, classify, rectangles, humans, animals
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/vision"
)

type result struct {
	Type       string   `json:"type,omitempty"`
	Confidence float32  `json:"confidence"`
	BBox       *bbox    `json:"bbox,omitempty"`
	Label      string   `json:"label,omitempty"`
	Labels     []label  `json:"labels,omitempty"`
	Payload    string   `json:"payload,omitempty"`
	Symbology  string   `json:"symbology,omitempty"`
	Corners    *corners `json:"corners,omitempty"`
}

type bbox struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

type label struct {
	Identifier string  `json:"identifier"`
	Confidence float32 `json:"confidence"`
}

type corners struct {
	TopLeft     point `json:"top_left"`
	TopRight    point `json:"top_right"`
	BottomLeft  point `json:"bottom_left"`
	BottomRight point `json:"bottom_right"`
}

type point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

var subcommands = map[string]bool{
	"faces": true, "barcodes": true, "classify": true,
	"rectangles": true, "humans": true, "animals": true,
}

func main() {
	jsonOut := flag.Bool("j", false, "JSON output")
	minConf := flag.Float64("min", 0.0, "minimum confidence threshold")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: vndetect [flags] <subcommand> <image-file>\n")
		fmt.Fprintf(os.Stderr, "       cat image.png | vndetect [flags] <subcommand> -\n\n")
		fmt.Fprintf(os.Stderr, "subcommands: faces, barcodes, classify, rectangles, humans, animals\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(2)
	}

	sub := flag.Arg(0)
	if !subcommands[sub] {
		fmt.Fprintf(os.Stderr, "error: unknown subcommand %q\n", sub)
		flag.Usage()
		os.Exit(2)
	}

	handler, err := makeHandler(flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	results, err := detect(sub, handler, float32(*minConf))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
	if len(results) == 0 {
		os.Exit(1)
	}

	if *jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(results); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(2)
		}
		return
	}

	for i, r := range results {
		printResult(i, r)
	}
}

func detect(sub string, handler vision.VNImageRequestHandler, minConf float32) ([]result, error) {
	switch sub {
	case "faces":
		return detectFaces(handler, minConf)
	case "barcodes":
		return detectBarcodes(handler, minConf)
	case "classify":
		return classifyImage(handler, minConf)
	case "rectangles":
		return detectRectangles(handler, minConf)
	case "humans":
		return detectHumans(handler, minConf)
	case "animals":
		return recognizeAnimals(handler, minConf)
	default:
		return nil, fmt.Errorf("unknown subcommand %q", sub)
	}
}

func detectFaces(handler vision.VNImageRequestHandler, minConf float32) ([]result, error) {
	request := vision.NewVNDetectFaceRectanglesRequest()
	if err := perform(handler, request.VNImageBasedRequest.VNRequest); err != nil {
		return nil, err
	}
	var results []result
	for _, obs := range request.Results() {
		face := vision.VNFaceObservationFromID(obs.ID)
		conf := float32(face.Confidence())
		if conf < minConf {
			continue
		}
		bb := face.BoundingBox()
		results = append(results, result{
			Type:       "face",
			Confidence: conf,
			BBox:       makeBBox(bb),
		})
	}
	return results, nil
}

func detectBarcodes(handler vision.VNImageRequestHandler, minConf float32) ([]result, error) {
	request := vision.NewVNDetectBarcodesRequest()
	if err := perform(handler, request.VNImageBasedRequest.VNRequest); err != nil {
		return nil, err
	}
	var results []result
	for _, obs := range request.Results() {
		bc := vision.VNBarcodeObservationFromID(obs.ID)
		conf := float32(bc.Confidence())
		if conf < minConf {
			continue
		}
		bb := bc.BoundingBox()
		results = append(results, result{
			Type:       "barcode",
			Confidence: conf,
			BBox:       makeBBox(bb),
			Payload:    bc.PayloadStringValue(),
			Symbology:  string(bc.Symbology()),
		})
	}
	return results, nil
}

func classifyImage(handler vision.VNImageRequestHandler, minConf float32) ([]result, error) {
	request := vision.NewVNClassifyImageRequest()
	if err := perform(handler, request.VNImageBasedRequest.VNRequest); err != nil {
		return nil, err
	}
	var results []result
	for _, obs := range request.Results() {
		cls := vision.VNClassificationObservationFromID(obs.ID)
		conf := float32(cls.Confidence())
		if conf < minConf {
			continue
		}
		results = append(results, result{
			Type:       "classification",
			Confidence: conf,
			Label:      cls.Identifier(),
		})
	}
	return results, nil
}

func detectRectangles(handler vision.VNImageRequestHandler, minConf float32) ([]result, error) {
	request := vision.NewVNDetectRectanglesRequest()
	if err := perform(handler, request.VNImageBasedRequest.VNRequest); err != nil {
		return nil, err
	}
	var results []result
	for _, obs := range request.Results() {
		rect := vision.VNRectangleObservationFromID(obs.ID)
		conf := float32(rect.Confidence())
		if conf < minConf {
			continue
		}
		bb := rect.BoundingBox()
		tl := rect.TopLeft()
		tr := rect.TopRight()
		bl := rect.BottomLeft()
		br := rect.BottomRight()
		results = append(results, result{
			Type:       "rectangle",
			Confidence: conf,
			BBox:       makeBBox(bb),
			Corners: &corners{
				TopLeft:     point{tl.X, tl.Y},
				TopRight:    point{tr.X, tr.Y},
				BottomLeft:  point{bl.X, bl.Y},
				BottomRight: point{br.X, br.Y},
			},
		})
	}
	return results, nil
}

func detectHumans(handler vision.VNImageRequestHandler, minConf float32) ([]result, error) {
	request := vision.NewVNDetectHumanRectanglesRequest()
	if err := perform(handler, request.VNImageBasedRequest.VNRequest); err != nil {
		return nil, err
	}
	var results []result
	for _, obs := range request.Results() {
		human := vision.VNDetectedObjectObservationFromID(obs.ID)
		conf := float32(human.Confidence())
		if conf < minConf {
			continue
		}
		bb := human.BoundingBox()
		results = append(results, result{
			Type:       "human",
			Confidence: conf,
			BBox:       makeBBox(bb),
		})
	}
	return results, nil
}

func recognizeAnimals(handler vision.VNImageRequestHandler, minConf float32) ([]result, error) {
	request := vision.NewVNRecognizeAnimalsRequest()
	if err := perform(handler, request.VNImageBasedRequest.VNRequest); err != nil {
		return nil, err
	}
	var results []result
	for _, obs := range request.Results() {
		animal := vision.VNRecognizedObjectObservationFromID(obs.ID)
		conf := float32(animal.Confidence())
		if conf < minConf {
			continue
		}
		bb := animal.BoundingBox()
		var labels []label
		for _, l := range animal.Labels() {
			labels = append(labels, label{
				Identifier: l.Identifier(),
				Confidence: float32(l.Confidence()),
			})
		}
		results = append(results, result{
			Type:       "animal",
			Confidence: conf,
			BBox:       makeBBox(bb),
			Labels:     labels,
		})
	}
	return results, nil
}

func perform(handler vision.VNImageRequestHandler, request vision.VNRequest) error {
	ok, err := handler.PerformRequestsError([]vision.VNRequest{request})
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("request failed")
	}
	return nil
}

func makeBBox(r corefoundation.CGRect) *bbox {
	return &bbox{
		X: r.Origin.X,
		Y: r.Origin.Y,
		W: r.Size.Width,
		H: r.Size.Height,
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

func printResult(i int, r result) {
	switch r.Type {
	case "face", "human":
		fmt.Printf("  %s %d: origin=(%.3f, %.3f) size=(%.3f x %.3f) confidence=%.2f\n",
			r.Type, i+1, r.BBox.X, r.BBox.Y, r.BBox.W, r.BBox.H, r.Confidence)
	case "barcode":
		fmt.Printf("  barcode %d: %s [%s] confidence=%.2f\n",
			i+1, r.Payload, r.Symbology, r.Confidence)
	case "classification":
		fmt.Printf("  %s: %.4f\n", r.Label, r.Confidence)
	case "rectangle":
		fmt.Printf("  rect %d: origin=(%.3f, %.3f) size=(%.3f x %.3f) confidence=%.2f\n",
			i+1, r.BBox.X, r.BBox.Y, r.BBox.W, r.BBox.H, r.Confidence)
	case "animal":
		labels := ""
		for j, l := range r.Labels {
			if j > 0 {
				labels += ", "
			}
			labels += fmt.Sprintf("%s(%.2f)", l.Identifier, l.Confidence)
		}
		fmt.Printf("  animal %d: origin=(%.3f, %.3f) size=(%.3f x %.3f) labels=[%s] confidence=%.2f\n",
			i+1, r.BBox.X, r.BBox.Y, r.BBox.W, r.BBox.H, labels, r.Confidence)
	}
}
