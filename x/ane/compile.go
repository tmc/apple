//go:build darwin

package ane

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

// Compile compiles a model and returns a ready-to-evaluate Model.
func (c *Client) Compile(opts CompileOptions) (*Model, error) {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return nil, fmt.Errorf("ane: client closed")
	}
	c.mu.Unlock()

	if opts.QoS == 0 {
		opts.QoS = 21
	}

	switch opts.ModelType {
	case ModelTypeMIL:
		k, err := compileMIL(c, opts)
		if err != nil {
			return nil, err
		}
		c.compiles.Add(1)
		return k, nil
	case ModelTypePackage:
		k, err := compilePackage(c, opts)
		if err != nil {
			return nil, err
		}
		c.compiles.Add(1)
		return k, nil
	default:
		return nil, fmt.Errorf("ane: unknown model type %d", opts.ModelType)
	}
}

func compileMIL(c *Client, opts CompileOptions) (*Model, error) {
	// Build the network text as NSData.
	networkText := foundation.NewDataFromBytes(opts.MILText)

	// Build the weights dictionary. Weightless MIL still expects an empty
	// dictionary here; passing nil can fail descriptor creation for ops like
	// softmax.
	weightFiles, err := compileWeightFiles(opts)
	if err != nil {
		return nil, &ANEError{Op: "compile", Err: err}
	}
	weightsDict := buildWeightsDict(weightFiles)

	// Create the model descriptor.
	descClass := appleneuralengine.GetANEInMemoryModelDescriptorClass()
	descObj := descClass.ModelWithMILTextWeightsOptionsPlist(networkText, weightsDict, nil)
	if descObj == nil || descObj.GetID() == 0 {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("failed to create model descriptor")}
	}

	// Create the in-memory model from the descriptor.
	modelClass := appleneuralengine.GetANEInMemoryModelClass()
	modelObj := modelClass.InMemoryModelWithDescriptor(descObj)
	if modelObj == nil || modelObj.GetID() == 0 {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("failed to create in-memory model")}
	}
	model := appleneuralengine.ANEInMemoryModelFromID(modelObj.GetID())

	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	// Pre-populate the temp directory that the Espresso IR translator expects.
	if err := prepopulateTempDir(model, opts.MILText, weightFiles); err != nil {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("pre-populate temp dir: %w", err)}
	}

	emptyOpts := foundation.NewNSMutableDictionary()

	// Compile.
	ok, err := model.CompileWithQoSOptionsError(opts.QoS, emptyOpts)
	if err != nil || !ok {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("compile failed: %w", err)}
	}

	// Re-set mask before load; compile preserves it but load clears it.
	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	// Load (with one retry after 100ms).
	ok, err = model.LoadWithQoSOptionsError(opts.QoS, emptyOpts)
	if err != nil || !ok {
		time.Sleep(100 * time.Millisecond)
		if opts.PerfStatsMask != 0 {
			model.SetPerfStatsMask(opts.PerfStatsMask)
		}
		ok, err = model.LoadWithQoSOptionsError(opts.QoS, emptyOpts)
		if err != nil || !ok {
			return nil, &ANEError{Op: "load", Err: fmt.Errorf("load failed: %w", err)}
		}
	}

	// Re-set mask after load on both the wrapper and inner model.
	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
		if inner := model.Model(); inner != nil {
			inner.SetPerfStatsMask(opts.PerfStatsMask)
		}
	}

	// Parse model attributes to get the compiled tensor layouts.
	aneModel := model.Model()
	if aneModel == nil {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("in-memory model has no underlying ANEModel after compile")}
	}
	inputLayouts, outputLayouts, err := parseModelLayouts(aneModel.ModelAttributes())
	if err != nil {
		model.UnloadWithQoSError(opts.QoS)
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("parse model layouts: %w", err)}
	}

	// Create surfaces and request using model-driven layouts.
	request, inputs, outputs, err := createRequestAndSurfaces(inputLayouts, outputLayouts)
	if err != nil {
		model.UnloadWithQoSError(opts.QoS)
		return nil, err
	}

	// Map IOSurfaces.
	ok, err = model.MapIOSurfacesWithRequestCacheInferenceError(request, true)
	if err != nil || !ok {
		ok, err = model.MapIOSurfacesWithRequestCacheInferenceError(request, false)
		if err != nil || !ok {
			model.UnloadWithQoSError(opts.QoS)
			return nil, &ANEError{Op: "map", Err: fmt.Errorf("map IOSurfaces failed: %w", err)}
		}
	}

	m := &Model{
		client:        c,
		modelType:     ModelTypeMIL,
		qos:           opts.QoS,
		inMemModel:    model,
		request:       request,
		inputs:        inputs,
		outputs:       outputs,
		perfStatsMask: opts.PerfStatsMask,
		inputLayouts:  inputLayouts,
		outputLayouts: outputLayouts,
		mapped:        true,
	}
	runtime.SetFinalizer(m, (*Model).Close)
	return m, nil
}

func compileWeightFiles(opts CompileOptions) ([]WeightFile, error) {
	files := make([]WeightFile, 0, len(opts.WeightFiles)+1)
	if len(opts.WeightBlob) > 0 {
		path := opts.WeightPath
		if path == "" {
			path = "@model_path/weights/weight.bin"
		}
		files = append(files, WeightFile{Path: path, Blob: opts.WeightBlob})
	}
	files = append(files, opts.WeightFiles...)

	seen := make(map[string]struct{}, len(files))
	for i := range files {
		if files[i].Path == "" {
			return nil, fmt.Errorf("weight file path is empty")
		}
		if _, ok := seen[files[i].Path]; ok {
			return nil, fmt.Errorf("duplicate weight path %q", files[i].Path)
		}
		seen[files[i].Path] = struct{}{}
	}
	return files, nil
}

func compilePackage(c *Client, opts CompileOptions) (*Model, error) {
	if opts.PackagePath == "" {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("PackagePath is required for ModelTypePackage")}
	}

	modelKey := opts.ModelKey
	if modelKey == "" {
		modelKey = "s"
	}

	// Create the model from the package path.
	modelURL := foundation.NewURLFileURLWithPath(opts.PackagePath)
	modelClass := appleneuralengine.GetANEModelClass()
	keyObj := foundation.NewStringWithString(modelKey)
	modelObj := modelClass.ModelAtURLKey(modelURL, keyObj)
	if modelObj == nil || modelObj.GetID() == 0 {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("failed to create model from %s", opts.PackagePath)}
	}
	model := appleneuralengine.ANEModelFromID(modelObj.GetID())

	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	aneClient := packageClient(c.aneClient)
	emptyOpts := foundation.NewNSMutableDictionary()

	// Compile.
	ok, err := aneClient.CompileModelOptionsQosError(model, emptyOpts, opts.QoS)
	if err != nil || !ok {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("compile failed: %w", err)}
	}

	// Re-set mask before load.
	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	// Load (with one retry after 100ms).
	ok, err = aneClient.LoadModelOptionsQosError(model, emptyOpts, opts.QoS)
	if err != nil || !ok {
		time.Sleep(100 * time.Millisecond)
		if opts.PerfStatsMask != 0 {
			model.SetPerfStatsMask(opts.PerfStatsMask)
		}
		ok, err = aneClient.LoadModelOptionsQosError(model, emptyOpts, opts.QoS)
		if err != nil || !ok {
			return nil, &ANEError{Op: "load", Err: fmt.Errorf("load failed: %w", err)}
		}
	}

	// Re-set mask after load.
	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	// Parse model attributes to get the compiled tensor layouts.
	inputLayouts, outputLayouts, err := parseModelLayouts(model.ModelAttributes())
	if err != nil {
		aneClient.UnloadModelOptionsQosError(model, nil, opts.QoS)
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("parse model layouts: %w", err)}
	}

	// Create surfaces and request.
	request, inputs, outputs, err := createRequestAndSurfaces(inputLayouts, outputLayouts)
	if err != nil {
		aneClient.UnloadModelOptionsQosError(model, nil, opts.QoS)
		return nil, err
	}

	// Map IOSurfaces.
	ok, err = aneClient.MapIOSurfacesWithModelRequestCacheInferenceError(model, request, true)
	if err != nil || !ok {
		ok, err = aneClient.MapIOSurfacesWithModelRequestCacheInferenceError(model, request, false)
		if err != nil || !ok {
			aneClient.UnloadModelOptionsQosError(model, nil, opts.QoS)
			return nil, &ANEError{Op: "map", Err: fmt.Errorf("map IOSurfaces failed: %w", err)}
		}
	}

	m := &Model{
		client:        c,
		modelType:     ModelTypePackage,
		qos:           opts.QoS,
		aneModel:      model,
		aneClient:     aneClient,
		request:       request,
		inputs:        inputs,
		outputs:       outputs,
		perfStatsMask: opts.PerfStatsMask,
		inputLayouts:  inputLayouts,
		outputLayouts: outputLayouts,
		mapped:        true,
	}
	runtime.SetFinalizer(m, (*Model).Close)
	return m, nil
}

// parseModelLayouts extracts tensor layouts from the compiled model's attributes.
// The attributes dictionary contains NetworkStatusList with LiveInputList and
// LiveOutputList, each entry having Channels, Width, Height, PlaneStride,
// RowStride, and Type fields.
func parseModelLayouts(attrs foundation.INSDictionary) ([]TensorLayout, []TensorLayout, error) {
	if attrs == nil || attrs.GetID() == 0 {
		return nil, nil, fmt.Errorf("model attributes are nil")
	}

	// Get NetworkStatusList (NSArray of NSDictionary).
	netListObj := dictGet(attrs, "NetworkStatusList")
	if netListObj == 0 {
		return nil, nil, fmt.Errorf("no NetworkStatusList in model attributes")
	}
	netList := foundation.NSArrayFromID(netListObj)
	if netList.Count() == 0 {
		return nil, nil, fmt.Errorf("empty NetworkStatusList")
	}

	// Use the first procedure's layout.
	proc := netList.ObjectAtIndex(0)
	inputSymbols, outputSymbols, err := parseProcedureSymbolMaps(attrs)
	if err != nil {
		return nil, nil, err
	}

	// Parse LiveInputList.
	inputListID := dictGet(proc, "LiveInputList")
	if inputListID == 0 {
		return nil, nil, fmt.Errorf("no LiveInputList in model attributes")
	}
	inputList := foundation.NSArrayFromID(inputListID)
	inputLayouts := make([]TensorLayout, inputList.Count())
	for i := uint(0); i < inputList.Count(); i++ {
		entry := inputList.ObjectAtIndex(i)
		l := parseTensorEntry(entry, inputSymbols)
		if err := validateLayout(l); err != nil {
			return nil, nil, fmt.Errorf("%w: input[%d]: %v", ErrUnsupportedLayout, i, err)
		}
		inputLayouts[i] = l
	}

	// Parse LiveOutputList.
	outputListID := dictGet(proc, "LiveOutputList")
	if outputListID == 0 {
		return nil, nil, fmt.Errorf("no LiveOutputList in model attributes")
	}
	outputList := foundation.NSArrayFromID(outputListID)
	outputLayouts := make([]TensorLayout, outputList.Count())
	for i := uint(0); i < outputList.Count(); i++ {
		entry := outputList.ObjectAtIndex(i)
		l := parseTensorEntry(entry, outputSymbols)
		if err := validateLayout(l); err != nil {
			return nil, nil, fmt.Errorf("%w: output[%d]: %v", ErrUnsupportedLayout, i, err)
		}
		outputLayouts[i] = l
	}

	return inputLayouts, outputLayouts, nil
}

// parseTensorEntry extracts a TensorLayout from a single LiveInputList/LiveOutputList entry.
func parseTensorEntry(entry objectivec.IObject, symbolIndices map[string]int) TensorLayout {
	channels := dictGetInt(entry, "Channels")
	width := dictGetInt(entry, "Width")
	height := dictGetInt(entry, "Height")
	planeStride := dictGetInt(entry, "PlaneStride")
	rowStride := dictGetInt(entry, "RowStride")
	typeName := dictGetString(entry, "Type")
	name := dictGetString(entry, "Name")
	symbol := dictGetString(entry, "Symbol")

	elemSize := 2 // fp16 default
	if typeName == "Float32" {
		elemSize = 4
	}

	symbolIndex := -1
	switch {
	case symbol != "":
		if idx, ok := symbolIndices[symbol]; ok {
			symbolIndex = idx
		}
	case name != "":
		if idx, ok := symbolIndices[name]; ok {
			symbolIndex = idx
		}
	}

	return TensorLayout{
		Channels:    channels,
		Width:       width,
		Height:      height,
		ElemSize:    elemSize,
		RowStride:   rowStride,
		PlaneStride: planeStride,
		Name:        name,
		Symbol:      symbol,
		SymbolIndex: symbolIndex,
	}
}

func parseProcedureSymbolMaps(attrs foundation.INSDictionary) (map[string]int, map[string]int, error) {
	descID := dictGet(attrs, "ANEFModelDescription")
	if descID == 0 {
		return nil, nil, nil
	}
	desc := objectivec.ObjectFromID(descID)
	return symbolPositions(stringArrayForKey(desc, "kANEFModelInputSymbolsArrayKey")),
		symbolPositions(stringArrayForKey(desc, "kANEFModelOutputSymbolsArrayKey")),
		nil
}

func stringArrayForKey(dict objectivec.IObject, key string) []string {
	id := dictGet(dict, key)
	if id == 0 {
		return nil
	}
	arr := foundation.NSArrayFromID(id)
	out := make([]string, 0, arr.Count())
	for i := uint(0); i < arr.Count(); i++ {
		out = append(out, foundation.NSStringFromID(arr.ObjectAtIndex(i).GetID()).String())
	}
	return out
}

func symbolPositions(names []string) map[string]int {
	out := make(map[string]int, len(names))
	for i, name := range names {
		out[name] = i
	}
	return out
}

// dictGet returns the objc.ID for a string key in a dictionary, or 0 if not found.
func dictGet(dict objectivec.IObject, key string) objc.ID {
	return objc.Send[objc.ID](dict.GetID(), objc.Sel("objectForKey:"), objc.String(key))
}

// dictGetInt returns the int value for a string key in a dictionary, or 0 if not found.
func dictGetInt(dict objectivec.IObject, key string) int {
	id := dictGet(dict, key)
	if id == 0 {
		return 0
	}
	return foundation.NSNumberFromID(id).IntValue()
}

// dictGetString returns the string value for a string key in a dictionary, or "" if not found.
func dictGetString(dict objectivec.IObject, key string) string {
	id := dictGet(dict, key)
	if id == 0 {
		return ""
	}
	return foundation.NSStringFromID(id).UTF8String()
}

// newPerfStats creates a valid ANEPerformanceStats via the class factory.
func newPerfStats() appleneuralengine.ANEPerformanceStats {
	cls := appleneuralengine.GetANEPerformanceStatsClass()
	obj := cls.StatsWithHardwareExecutionNS(0)
	return appleneuralengine.ANEPerformanceStatsFromID(obj.GetID())
}

// createRequestAndSurfaces creates IOSurfaces from model-driven layouts
// and wraps them in an ANERequest.
func createRequestAndSurfaces(inputLayouts, outputLayouts []TensorLayout) (appleneuralengine.ANERequest, []coregraphics.IOSurfaceRef, []coregraphics.IOSurfaceRef, error) {
	if len(inputLayouts) == 0 || len(outputLayouts) == 0 {
		return appleneuralengine.ANERequest{}, nil, nil, &ANEError{Op: "map", Err: fmt.Errorf("input and output layouts must be specified")}
	}

	ioClass := appleneuralengine.GetANEIOSurfaceObjectClass()

	// Create input IOSurfaces.
	inputs := make([]coregraphics.IOSurfaceRef, len(inputLayouts))
	inputArr := foundation.NewNSMutableArray()
	inputIdxArr := foundation.NewNSMutableArray()
	for i, layout := range inputLayouts {
		ref, err := createSurfaceForLayout(layout)
		if err != nil {
			return appleneuralengine.ANERequest{}, nil, nil, &ANEError{Op: "map", Err: err}
		}
		inputs[i] = ref
		wrapped := ioClass.ObjectWithIOSurface(ref)
		inputArr.AddObject(wrapped)
		symbolIndex := i
		if layout.SymbolIndex >= 0 {
			symbolIndex = layout.SymbolIndex
		}
		inputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(symbolIndex))
	}

	// Create output IOSurfaces.
	outputs := make([]coregraphics.IOSurfaceRef, len(outputLayouts))
	outputArr := foundation.NewNSMutableArray()
	outputIdxArr := foundation.NewNSMutableArray()
	for i, layout := range outputLayouts {
		ref, err := createSurfaceForLayout(layout)
		if err != nil {
			return appleneuralengine.ANERequest{}, nil, nil, &ANEError{Op: "map", Err: err}
		}
		outputs[i] = ref
		wrapped := ioClass.ObjectWithIOSurface(ref)
		outputArr.AddObject(wrapped)
		symbolIndex := i
		if layout.SymbolIndex >= 0 {
			symbolIndex = layout.SymbolIndex
		}
		outputIdxArr.AddObject(foundation.GetNSNumberClass().NumberWithInt(symbolIndex))
	}

	procIdx := foundation.GetNSNumberClass().NumberWithInt(0)
	txnHandle := foundation.GetNSNumberClass().NumberWithUnsignedLongLong(1)

	// Build the request.
	reqClass := appleneuralengine.GetANERequestClass()
	reqObj := reqClass.RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(
		inputArr,
		inputIdxArr,
		outputArr,
		outputIdxArr,
		nil, // weightsBuffer
		nil, // perfStats — attached per-eval via SetPerfStats
		procIdx,
		nil, // sharedEvents — attached per-eval
		txnHandle,
	)
	if reqObj == nil || reqObj.GetID() == 0 {
		reqObj = reqClass.RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndex(
			inputArr,
			inputIdxArr,
			outputArr,
			outputIdxArr,
			nil, // weightsBuffer
			nil, // perfStats — attached per-eval via SetPerfStats
			procIdx,
		)
	}
	if reqObj == nil || reqObj.GetID() == 0 {
		return appleneuralengine.ANERequest{}, nil, nil, &ANEError{Op: "map", Err: fmt.Errorf("failed to create request")}
	}
	request := appleneuralengine.ANERequestFromID(reqObj.GetID())
	if request.TransactionHandle().GetID() == 0 {
		request.SetTransactionHandle(txnHandle)
	}

	return request, inputs, outputs, nil
}

// prepopulateTempDir writes the MIL text and weight blobs to the temp directory
// that the ANE compiler's Espresso IR translator expects to find them at.
func prepopulateTempDir(model appleneuralengine.ANEInMemoryModel, milText []byte, weightFiles []WeightFile) error {
	hexID := model.HexStringIdentifier()
	if hexID == "" {
		return fmt.Errorf("model has no hex identifier")
	}

	tmpDir := filepath.Join(os.TempDir(), hexID)

	weightsDir := filepath.Join(tmpDir, "weights")
	if err := os.MkdirAll(weightsDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", weightsDir, err)
	}

	milPath := filepath.Join(tmpDir, "model.mil")
	if err := os.WriteFile(milPath, milText, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", milPath, err)
	}

	for _, wf := range weightFiles {
		relPath := wf.Path
		if len(relPath) > 12 && relPath[:12] == "@model_path/" {
			relPath = relPath[12:]
		}
		blobPath := filepath.Join(tmpDir, relPath)
		blobDir := filepath.Dir(blobPath)
		if err := os.MkdirAll(blobDir, 0o755); err != nil {
			return fmt.Errorf("mkdir %s: %w", blobDir, err)
		}
		if err := os.WriteFile(blobPath, wf.Blob, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", blobPath, err)
		}
	}

	return nil
}

// validateLayout checks that a TensorLayout has valid dimensions and alignment.
func validateLayout(l TensorLayout) error {
	if l.Channels <= 0 || l.Width <= 0 || l.Height <= 0 {
		return fmt.Errorf("invalid dimensions: C=%d H=%d W=%d", l.Channels, l.Height, l.Width)
	}
	if l.ElemSize != 2 && l.ElemSize != 4 {
		return fmt.Errorf("unsupported element size %d (want 2 or 4)", l.ElemSize)
	}
	if l.RowStride%64 != 0 {
		return fmt.Errorf("RowStride %d is not 64-byte aligned", l.RowStride)
	}
	minStride := RowStrideFor(l.Width, l.ElemSize)
	if l.RowStride < minStride {
		return fmt.Errorf("RowStride %d < minimum %d for Width=%d ElemSize=%d", l.RowStride, minStride, l.Width, l.ElemSize)
	}
	expectedPlane := l.Height * l.RowStride
	if l.PlaneStride < expectedPlane {
		return fmt.Errorf("PlaneStride %d < Height*RowStride (%d*%d = %d)", l.PlaneStride, l.Height, l.RowStride, expectedPlane)
	}
	return nil
}

// CompileWithStats compiles a model and returns a Model along with compilation timing.
func (c *Client) CompileWithStats(opts CompileOptions) (*Model, CompileStats, error) {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return nil, CompileStats{}, fmt.Errorf("ane: client closed")
	}
	c.mu.Unlock()

	if opts.QoS == 0 {
		opts.QoS = 21
	}

	totalStart := time.Now()
	var cs CompileStats

	switch opts.ModelType {
	case ModelTypeMIL:
		k, err := compileMILWithStats(c, opts, &cs)
		if err != nil {
			return nil, CompileStats{}, err
		}
		cs.TotalNS = time.Since(totalStart).Nanoseconds()
		c.compiles.Add(1)
		return k, cs, nil
	case ModelTypePackage:
		k, err := compilePackageWithStats(c, opts, &cs)
		if err != nil {
			return nil, CompileStats{}, err
		}
		cs.TotalNS = time.Since(totalStart).Nanoseconds()
		c.compiles.Add(1)
		return k, cs, nil
	default:
		return nil, CompileStats{}, fmt.Errorf("ane: unknown model type %d", opts.ModelType)
	}
}

func compileMILWithStats(c *Client, opts CompileOptions, cs *CompileStats) (*Model, error) {
	networkText := foundation.NewDataFromBytes(opts.MILText)

	weightFiles, err := compileWeightFiles(opts)
	if err != nil {
		return nil, &ANEError{Op: "compile", Err: err}
	}
	weightsDict := buildWeightsDict(weightFiles)

	descClass := appleneuralengine.GetANEInMemoryModelDescriptorClass()
	descObj := descClass.ModelWithMILTextWeightsOptionsPlist(networkText, weightsDict, nil)
	if descObj == nil || descObj.GetID() == 0 {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("failed to create model descriptor")}
	}

	modelClass := appleneuralengine.GetANEInMemoryModelClass()
	modelObj := modelClass.InMemoryModelWithDescriptor(descObj)
	if modelObj == nil || modelObj.GetID() == 0 {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("failed to create in-memory model")}
	}
	model := appleneuralengine.ANEInMemoryModelFromID(modelObj.GetID())

	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	if err := prepopulateTempDir(model, opts.MILText, weightFiles); err != nil {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("pre-populate temp dir: %w", err)}
	}

	emptyOpts := foundation.NewNSMutableDictionary()

	// Compile with timing.
	compileStart := time.Now()
	ok, err := model.CompileWithQoSOptionsError(opts.QoS, emptyOpts)
	if err != nil || !ok {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("compile failed: %w", err)}
	}
	cs.CompileNS = time.Since(compileStart).Nanoseconds()

	// Re-set mask before load; load clears it.
	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	// Load with timing (with one retry after 100ms).
	loadStart := time.Now()
	ok, err = model.LoadWithQoSOptionsError(opts.QoS, emptyOpts)
	if err != nil || !ok {
		time.Sleep(100 * time.Millisecond)
		if opts.PerfStatsMask != 0 {
			model.SetPerfStatsMask(opts.PerfStatsMask)
		}
		ok, err = model.LoadWithQoSOptionsError(opts.QoS, emptyOpts)
		if err != nil || !ok {
			return nil, &ANEError{Op: "load", Err: fmt.Errorf("load failed: %w", err)}
		}
	}
	cs.LoadNS = time.Since(loadStart).Nanoseconds()

	// Re-set mask after load on both wrapper and inner model.
	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
		if inner := model.Model(); inner != nil {
			inner.SetPerfStatsMask(opts.PerfStatsMask)
		}
	}

	aneModel := model.Model()
	if aneModel == nil {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("in-memory model has no underlying ANEModel after compile")}
	}
	inputLayouts, outputLayouts, err := parseModelLayouts(aneModel.ModelAttributes())
	if err != nil {
		model.UnloadWithQoSError(opts.QoS)
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("parse model layouts: %w", err)}
	}

	request, inputs, outputs, err := createRequestAndSurfaces(inputLayouts, outputLayouts)
	if err != nil {
		model.UnloadWithQoSError(opts.QoS)
		return nil, err
	}

	ok, err = model.MapIOSurfacesWithRequestCacheInferenceError(request, true)
	if err != nil || !ok {
		ok, err = model.MapIOSurfacesWithRequestCacheInferenceError(request, false)
		if err != nil || !ok {
			model.UnloadWithQoSError(opts.QoS)
			return nil, &ANEError{Op: "map", Err: fmt.Errorf("map IOSurfaces failed: %w", err)}
		}
	}

	m := &Model{
		client:        c,
		modelType:     ModelTypeMIL,
		qos:           opts.QoS,
		inMemModel:    model,
		request:       request,
		inputs:        inputs,
		outputs:       outputs,
		perfStatsMask: opts.PerfStatsMask,
		inputLayouts:  inputLayouts,
		outputLayouts: outputLayouts,
		mapped:        true,
	}
	runtime.SetFinalizer(m, (*Model).Close)
	return m, nil
}

func compilePackageWithStats(c *Client, opts CompileOptions, cs *CompileStats) (*Model, error) {
	if opts.PackagePath == "" {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("PackagePath is required for ModelTypePackage")}
	}

	modelKey := opts.ModelKey
	if modelKey == "" {
		modelKey = "s"
	}

	modelURL := foundation.NewURLFileURLWithPath(opts.PackagePath)
	modelClass := appleneuralengine.GetANEModelClass()
	keyObj := foundation.NewStringWithString(modelKey)
	modelObj := modelClass.ModelAtURLKey(modelURL, keyObj)
	if modelObj == nil || modelObj.GetID() == 0 {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("failed to create model from %s", opts.PackagePath)}
	}
	model := appleneuralengine.ANEModelFromID(modelObj.GetID())

	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	aneClient := packageClient(c.aneClient)
	emptyOpts := foundation.NewNSMutableDictionary()

	// Compile with timing.
	compileStart := time.Now()
	ok, err := aneClient.CompileModelOptionsQosError(model, emptyOpts, opts.QoS)
	if err != nil || !ok {
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("compile failed: %w", err)}
	}
	cs.CompileNS = time.Since(compileStart).Nanoseconds()

	// Re-set mask before load.
	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	// Load with timing (with one retry after 100ms).
	loadStart := time.Now()
	ok, err = aneClient.LoadModelOptionsQosError(model, emptyOpts, opts.QoS)
	if err != nil || !ok {
		time.Sleep(100 * time.Millisecond)
		if opts.PerfStatsMask != 0 {
			model.SetPerfStatsMask(opts.PerfStatsMask)
		}
		ok, err = aneClient.LoadModelOptionsQosError(model, emptyOpts, opts.QoS)
		if err != nil || !ok {
			return nil, &ANEError{Op: "load", Err: fmt.Errorf("load failed: %w", err)}
		}
	}
	cs.LoadNS = time.Since(loadStart).Nanoseconds()

	// Re-set mask after load.
	if opts.PerfStatsMask != 0 {
		model.SetPerfStatsMask(opts.PerfStatsMask)
	}

	inputLayouts, outputLayouts, err := parseModelLayouts(model.ModelAttributes())
	if err != nil {
		aneClient.UnloadModelOptionsQosError(model, nil, opts.QoS)
		return nil, &ANEError{Op: "compile", Err: fmt.Errorf("parse model layouts: %w", err)}
	}

	request, inputs, outputs, err := createRequestAndSurfaces(inputLayouts, outputLayouts)
	if err != nil {
		aneClient.UnloadModelOptionsQosError(model, nil, opts.QoS)
		return nil, err
	}

	ok, err = aneClient.MapIOSurfacesWithModelRequestCacheInferenceError(model, request, true)
	if err != nil || !ok {
		ok, err = aneClient.MapIOSurfacesWithModelRequestCacheInferenceError(model, request, false)
		if err != nil || !ok {
			aneClient.UnloadModelOptionsQosError(model, nil, opts.QoS)
			return nil, &ANEError{Op: "map", Err: fmt.Errorf("map IOSurfaces failed: %w", err)}
		}
	}

	m := &Model{
		client:        c,
		modelType:     ModelTypePackage,
		qos:           opts.QoS,
		aneModel:      model,
		aneClient:     aneClient,
		request:       request,
		inputs:        inputs,
		outputs:       outputs,
		perfStatsMask: opts.PerfStatsMask,
		inputLayouts:  inputLayouts,
		outputLayouts: outputLayouts,
		mapped:        true,
	}
	runtime.SetFinalizer(m, (*Model).Close)
	return m, nil
}

func packageClient(fallback appleneuralengine.ANEClient) appleneuralengine.ANEClient {
	if c := appleneuralengine.GetANEClientClass().SharedConnection(); c != nil && c.ID != 0 {
		return *c
	}
	return fallback
}

// buildWeightsDict creates an NSDictionary
// { path: { "offset": 0, "data": blob }, ... }.
func buildWeightsDict(files []WeightFile) objectivec.IObject {
	outer := foundation.NewNSMutableDictionary()
	for _, wf := range files {
		data := foundation.NewDataFromBytes(wf.Blob)
		offset := foundation.GetNSNumberClass().NumberWithInt(0)

		inner := foundation.NewNSMutableDictionary()
		inner.SetObjectForKey(offset, foundation.NSCopyingObject{Object: objectivec.Object{ID: objc.String("offset")}})
		inner.SetObjectForKey(data, foundation.NSCopyingObject{Object: objectivec.Object{ID: objc.String("data")}})
		outer.SetObjectForKey(inner, foundation.NSCopyingObject{Object: objectivec.Object{ID: objc.String(wf.Path)}})
	}

	return outer
}
