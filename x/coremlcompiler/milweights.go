package coremlcompiler

import (
	"fmt"
	"path"
	"strings"
)

// A BlobLayout selects how a program's weights are distributed across files.
// Both layouts hold the same MIL Blob Storage v2 format written by
// [WriteMILBlob]; they differ only in how many blobs go in a file.
type BlobLayout int

const (
	// BlobLayoutSingleFile puts every tensor in one file, each at its own
	// offset. This is the layout coremltools produces and the one a
	// .mlpackage normally carries.
	BlobLayoutSingleFile BlobLayout = iota

	// BlobLayoutFilePerConst gives each tensor its own file, named for the
	// const that reads it. Every tensor is then the sole entry in its file and
	// so always sits at offset 64.
	BlobLayoutFilePerConst
)

func (l BlobLayout) String() string {
	switch l {
	case BlobLayoutSingleFile:
		return "single-file"
	case BlobLayoutFilePerConst:
		return "file-per-const"
	}
	return fmt.Sprintf("BlobLayout(%d)", int(l))
}

// DefaultWeightFileName is the file [BlobLayoutSingleFile] writes to, the name
// coremltools gives a .mlpackage's weight blob.
const DefaultWeightFileName = "weights/weight.bin"

// DefaultWeightDir is the directory [BlobLayoutFilePerConst] writes into.
const DefaultWeightDir = "weights"

// A WeightTensor is one tensor to place in a program's weights.
//
// Name is the MIL const that reads it, and names the file under
// [BlobLayoutFilePerConst]. Data is the tensor's raw little-endian payload;
// [Float16Bytes] produces it for an fp16 tensor. NumElements is required only
// for sub-byte element types, where the trailing partial byte has to be
// reported in the blob metadata.
type WeightTensor struct {
	Name        string
	DType       BlobDataType
	Data        []byte
	NumElements int
}

// A BlobRef locates one tensor's blob: the BLOBFILE path and offset the const
// reading it must name.
type BlobRef struct {
	Path   string
	Offset uint64
}

// BuildWeights places tensors into weight files according to layout. It
// returns the files, ready for [WriteWeightRoot], and the [BlobRef] for each
// tensor keyed by name, to be used as the BLOBFILE path and offset of the
// const that reads it.
//
// Tensor order is preserved, so the returned files are byte-for-byte
// reproducible for a given input.
func BuildWeights(tensors []WeightTensor, layout BlobLayout) ([]WeightFile, map[string]BlobRef, error) {
	seen := make(map[string]bool, len(tensors))
	for i, t := range tensors {
		if t.Name == "" {
			return nil, nil, fmt.Errorf("coremlcompiler: weight tensor %d has no name", i)
		}
		if seen[t.Name] {
			return nil, nil, fmt.Errorf("coremlcompiler: weight tensor %q appears twice", t.Name)
		}
		seen[t.Name] = true
	}

	refs := make(map[string]BlobRef, len(tensors))
	switch layout {
	case BlobLayoutSingleFile:
		entries := make([]BlobEntry, len(tensors))
		for i, t := range tensors {
			entries[i] = BlobEntry{DType: t.DType, Data: t.Data, NumElements: t.NumElements}
		}
		data, offsets := WriteMILBlob(entries)
		blobPath := ModelPathPrefix + DefaultWeightFileName
		for i, t := range tensors {
			refs[t.Name] = BlobRef{Path: blobPath, Offset: offsets[i]}
		}
		if len(tensors) == 0 {
			return nil, refs, nil
		}
		return []WeightFile{{Path: blobPath, Blob: data}}, refs, nil

	case BlobLayoutFilePerConst:
		files := make([]WeightFile, 0, len(tensors))
		for _, t := range tensors {
			name, err := weightFileName(t.Name)
			if err != nil {
				return nil, nil, err
			}
			data, offsets := WriteMILBlob([]BlobEntry{{DType: t.DType, Data: t.Data, NumElements: t.NumElements}})
			blobPath := ModelPathPrefix + path.Join(DefaultWeightDir, name)
			files = append(files, WeightFile{Path: blobPath, Blob: data})
			refs[t.Name] = BlobRef{Path: blobPath, Offset: offsets[0]}
		}
		return files, refs, nil
	}
	return nil, nil, fmt.Errorf("coremlcompiler: unknown blob layout %d", int(layout))
}

// weightFileName turns a const name into the file it gets under
// BlobLayoutFilePerConst. The name reaches the filesystem, so a name that
// would leave the weights directory is refused rather than sanitized: a
// silently renamed file is one the MIL text no longer points at.
func weightFileName(constName string) (string, error) {
	if strings.ContainsAny(constName, `/\`) || strings.Contains(constName, "..") {
		return "", fmt.Errorf("coremlcompiler: weight tensor %q cannot name a file", constName)
	}
	return constName + ".bin", nil
}
