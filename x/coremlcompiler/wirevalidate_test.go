package coremlcompiler

import "testing"

// TestValidateModelInterface covers the ModelDescription preconditions Core ML
// enforces at load time and the oneof exclusivity generated protobuf code
// would enforce for us.
func TestValidateModelInterface(t *testing.T) {
	arr := func() *ArrayFeatureType {
		return &ArrayFeatureType{Shape: []int64{1}, DataType: ArrayDataTypeFloat16}
	}
	ok := FeatureDescription{Name: "x", Type: &FeatureType{MultiArrayType: arr()}}
	tests := []struct {
		name    string
		model   *Model
		wantErr bool
	}{
		{
			name:  "single function ok",
			model: &Model{SpecVersion: 8, Description: ModelDescription{Inputs: []FeatureDescription{ok}}},
		},
		{
			name:    "spec version unset",
			model:   &Model{Description: ModelDescription{Inputs: []FeatureDescription{ok}}},
			wantErr: true,
		},
		{
			name: "functions and model level features both set",
			model: &Model{SpecVersion: 9, Description: ModelDescription{
				Functions:           []FunctionDescription{{Name: "main", Inputs: []FeatureDescription{ok}}},
				DefaultFunctionName: "main",
				Inputs:              []FeatureDescription{ok},
			}},
			wantErr: true,
		},
		{
			name: "default function name not a function",
			model: &Model{SpecVersion: 9, Description: ModelDescription{
				Functions:           []FunctionDescription{{Name: "main"}},
				DefaultFunctionName: "other",
			}},
			wantErr: true,
		},
		{
			name:    "empty feature name",
			model:   &Model{SpecVersion: 8, Description: ModelDescription{Inputs: []FeatureDescription{{Type: &FeatureType{Int64Type: true}}}}},
			wantErr: true,
		},
		{
			name:    "feature with no type",
			model:   &Model{SpecVersion: 8, Description: ModelDescription{Inputs: []FeatureDescription{{Name: "x"}}}},
			wantErr: true,
		},
		{
			name: "two oneof members set",
			model: &Model{SpecVersion: 8, Description: ModelDescription{Inputs: []FeatureDescription{
				{Name: "x", Type: &FeatureType{Int64Type: true, StringType: true}},
			}}},
			wantErr: true,
		},
		{
			name: "image without color space",
			model: &Model{SpecVersion: 8, Description: ModelDescription{Inputs: []FeatureDescription{
				{Name: "img", Type: &FeatureType{ImageType: &ImageFeatureType{Width: 2, Height: 2}}},
			}}},
			wantErr: true,
		},
		{
			name: "image with rgb color space",
			model: &Model{SpecVersion: 8, Description: ModelDescription{Inputs: []FeatureDescription{
				{Name: "img", Type: &FeatureType{ImageType: &ImageFeatureType{Width: 2, Height: 2, ColorSpace: ColorSpaceRGB}}},
			}}},
		},
		{
			name: "state feature must be fp16 or int8",
			model: &Model{SpecVersion: 9, Description: ModelDescription{States: []FeatureDescription{
				{Name: "s", Type: &FeatureType{StateArrayType: &ArrayFeatureType{Shape: []int64{1}, DataType: ArrayDataTypeFloat32}}},
			}}},
			wantErr: true,
		},
		{
			name: "state feature fp16 ok",
			model: &Model{SpecVersion: 9, Description: ModelDescription{States: []FeatureDescription{
				{Name: "s", Type: &FeatureType{StateArrayType: arr()}},
			}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateModelInterface(tt.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateModelInterface = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestEncodeEmptyTensorValueSetsField checks that an empty (zero-element)
// tensor const still sets its oneof field, as coremltools' SetInParent does.
