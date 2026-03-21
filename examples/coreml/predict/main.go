// Command predict loads a CoreML model and prints its description as JSON.
//
// It accepts both compiled (.mlmodelc) and uncompiled (.mlmodel) model files.
// Uncompiled models are compiled on the fly before loading.
//
// Usage:
//
//	predict <path-to-model>
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

func init() {
	runtime.LockOSThread()
}

type featureInfo struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Optional bool   `json:"optional"`
	// MultiArray constraint details, if applicable.
	Shape    []int  `json:"shape,omitempty"`
	DataType string `json:"data_type,omitempty"`
	// Image constraint details, if applicable.
	PixelsWide int    `json:"pixels_wide,omitempty"`
	PixelsHigh int    `json:"pixels_high,omitempty"`
	PixelFormat string `json:"pixel_format,omitempty"`
}

type modelInfo struct {
	Path                    string            `json:"path"`
	Inputs                  []featureInfo     `json:"inputs"`
	Outputs                 []featureInfo     `json:"outputs"`
	PredictedFeatureName    string            `json:"predicted_feature_name,omitempty"`
	PredictedProbabilities  string            `json:"predicted_probabilities_name,omitempty"`
	IsUpdatable             bool              `json:"is_updatable"`
	Metadata                map[string]string `json:"metadata,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: predict <model-path>\n")
		os.Exit(1)
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "predict: %v\n", err)
		os.Exit(1)
	}
}

func run(modelPath string) error {
	absPath, err := filepath.Abs(modelPath)
	if err != nil {
		return fmt.Errorf("resolve path: %w", err)
	}

	url := foundation.NewURLFileURLWithPath(absPath)

	// If the model is uncompiled (.mlmodel), compile it first.
	if strings.HasSuffix(absPath, ".mlmodel") {
		compiledURL, err := coreml.GetMLModelClass().CompileModelAtURLError(url)
		if err != nil {
			return fmt.Errorf("compile model: %w", err)
		}
		url = compiledURL
	}

	model, err := coreml.NewModelWithContentsOfURLError(url)
	if err != nil {
		return fmt.Errorf("load model: %w", err)
	}

	desc := model.ModelDescription().(coreml.MLModelDescription)

	info := modelInfo{
		Path:                   absPath,
		PredictedFeatureName:   desc.PredictedFeatureName(),
		PredictedProbabilities: desc.PredictedProbabilitiesName(),
		IsUpdatable:            desc.IsUpdatable(),
	}

	info.Inputs = describeFeatures(desc.InputDescriptionsByName().(foundation.NSDictionary))
	info.Outputs = describeFeatures(desc.OutputDescriptionsByName().(foundation.NSDictionary))
	info.Metadata = extractMetadata(desc.Metadata().(foundation.NSDictionary))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(info)
}

func describeFeatures(dict foundation.NSDictionary) []featureInfo {
	keys := dict.AllKeys()
	features := make([]featureInfo, 0, len(keys))
	for _, k := range keys {
		name := foundation.NSStringFromID(k.GetID()).String()
		val := dict.ObjectForKey(k)
		fd := coreml.MLFeatureDescriptionFromID(val.GetID())

		fi := featureInfo{
			Name:     name,
			Type:     featureTypeName(fd.Type()),
			Optional: fd.Optional(),
		}

		switch fd.Type() {
		case coreml.MLFeatureTypeMultiArray:
			c := fd.MultiArrayConstraint().(coreml.MLMultiArrayConstraint)
			if c.ID != 0 {
				for _, n := range c.Shape() {
					fi.Shape = append(fi.Shape, n.IntValue())
				}
				fi.DataType = c.DataType().String()
			}
		case coreml.MLFeatureTypeImage:
			c := fd.ImageConstraint().(coreml.MLImageConstraint)
			if c.ID != 0 {
				fi.PixelsWide = c.PixelsWide()
				fi.PixelsHigh = c.PixelsHigh()
				fi.PixelFormat = fmt.Sprintf("0x%x", c.PixelFormatType())
			}
		}

		features = append(features, fi)
	}
	return features
}

func extractMetadata(dict foundation.NSDictionary) map[string]string {
	if dict.Count() == 0 {
		return nil
	}
	m := make(map[string]string)
	for _, k := range dict.AllKeys() {
		key := foundation.NSStringFromID(k.GetID()).String()
		val := dict.ObjectForKey(k)
		desc := objc.Send[objc.ID](val.GetID(), objc.Sel("description"))
		m[key] = foundation.NSStringFromID(desc).String()
	}
	return m
}

func featureTypeName(t coreml.MLFeatureType) string {
	switch t {
	case coreml.MLFeatureTypeInt64:
		return "Int64"
	case coreml.MLFeatureTypeDouble:
		return "Double"
	case coreml.MLFeatureTypeString:
		return "String"
	case coreml.MLFeatureTypeImage:
		return "Image"
	case coreml.MLFeatureTypeMultiArray:
		return "MultiArray"
	case coreml.MLFeatureTypeDictionary:
		return "Dictionary"
	case coreml.MLFeatureTypeSequence:
		return "Sequence"
	case coreml.MLFeatureTypeState:
		return "State"
	default:
		return fmt.Sprintf("Unknown(%d)", t)
	}
}
