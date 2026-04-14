package coremlcompiler

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const defaultMILSpecVersion = 8

// CompileMILText compiles an already-emitted mlprogram MIL text into a
// compiled bundle at outputPath.
//
// desc names the model inputs, outputs, and states stored in coremldata.bin
// and metadata.json. weightRoot, when non-empty, is copied into the bundle
// preserving relative paths, so BLOBFILE paths like "@model_path/weights/..."
// resolve correctly at runtime.
func CompileMILText(milText string, specVersion int32, desc ModelDescription, weightRoot, outputPath string) error {
	if milText == "" {
		return fmt.Errorf("coremlcompiler: mil text is empty")
	}
	if specVersion == 0 {
		specVersion = defaultMILSpecVersion
	}
	model := &Model{
		SpecVersion:    specVersion,
		Description:    desc,
		MLProgram:      &Program{},
		descriptionRaw: encodeModelDescription(desc),
	}
	return compileMILTextModel(model, milText, weightRoot, outputPath)
}

func compileMILTextModel(model *Model, milText, weightRoot, outputPath string) error {
	if err := os.MkdirAll(outputPath, 0o755); err != nil {
		return fmt.Errorf("coremlcompiler: mkdir output: %w", err)
	}

	milPath := filepath.Join(outputPath, "model.mil")
	if err := os.WriteFile(milPath, []byte(milText), 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write model.mil: %w", err)
	}

	if weightRoot != "" {
		if err := copyDirContents(weightRoot, outputPath); err != nil {
			return fmt.Errorf("coremlcompiler: copy weights: %w", err)
		}
	}

	coremldata := buildCoreMLData(model)
	cdPath := filepath.Join(outputPath, "coremldata.bin")
	if err := os.WriteFile(cdPath, coremldata, 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write coremldata.bin: %w", err)
	}
	if err := writeAnalyticsCoreMLData(outputPath, EncodeModel(model)); err != nil {
		return fmt.Errorf("coremlcompiler: write analytics coremldata.bin: %w", err)
	}

	metaJSON, err := buildMetadataJSON(model)
	if err != nil {
		return fmt.Errorf("coremlcompiler: build metadata.json: %w", err)
	}
	metaPath := filepath.Join(outputPath, "metadata.json")
	if err := os.WriteFile(metaPath, metaJSON, 0o644); err != nil {
		return fmt.Errorf("coremlcompiler: write metadata.json: %w", err)
	}

	return nil
}

func copyDirContents(srcRoot, dstRoot string) error {
	return filepath.WalkDir(srcRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == srcRoot {
			return nil
		}
		rel, err := filepath.Rel(srcRoot, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstRoot, rel)
		if d.IsDir() {
			return os.MkdirAll(dstPath, 0o755)
		}
		if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
			return err
		}
		return copyFile(path, dstPath)
	})
}

func writeAnalyticsCoreMLData(outputPath string, modelProto []byte) error {
	analyticsDir := filepath.Join(outputPath, "analytics")
	if err := os.MkdirAll(analyticsDir, 0o755); err != nil {
		return err
	}
	data := buildAnalyticsCoreMLData(modelProto)
	return os.WriteFile(filepath.Join(analyticsDir, "coremldata.bin"), data, 0o644)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}
