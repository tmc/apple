package dynamicmatmul

import (
	"fmt"
	"sync"
	"time"

	"github.com/tmc/apple/x/ane/mil"
	"github.com/tmc/apple/x/ane/model"
)

const defaultQoS = uint32(21)

// Options configures executor creation.
type Options struct {
	QoS uint32

	// TileOut forces output-channel tiling when > 0.
	//
	// Each tile compiles a separate kernel with output width <= TileOut.
	// When TileOut == 0, New first tries a single full-width kernel and
	// falls back to tiling only if full-width compilation fails.
	TileOut int
}

// EvalStats reports per-eval ANE timing.
type EvalStats struct {
	HWExecutionNS uint64
	Metrics       map[string]float64
}

// Executor evaluates row-major y = x*w with runtime-provided weights.
//
// x has shape [batch, inDim], w has shape [inDim, outDim], and the result has
// shape [batch, outDim]. The executor compiles once for a fixed shape and
// rewrites the packed input surface on each evaluation.
type Executor struct {
	mu sync.Mutex

	batch  int
	inDim  int
	outDim int

	tiles        []tile
	prevOneHot   []int
	touchedRows  []int
	rowSeen      []bool
	weightsReady bool
}

type tile struct {
	outOffset int
	outDim    int
	k         *model.Kernel

	inputPacked  []float32
	outputPacked []float32
}

// New compiles a dynamic matmul kernel for the provided shape.
func New(batch, inDim, outDim int, opts Options) (*Executor, error) {
	if batch <= 0 || inDim <= 0 || outDim <= 0 {
		return nil, fmt.Errorf("dynamic matmul: invalid shape batch=%d inDim=%d outDim=%d", batch, inDim, outDim)
	}
	qos := opts.QoS
	if qos == 0 {
		qos = defaultQoS
	}

	tiles, err := compileTiles(batch, inDim, outDim, qos, opts.TileOut)
	if err != nil {
		return nil, err
	}

	return &Executor{
		batch:      batch,
		inDim:      inDim,
		outDim:     outDim,
		tiles:      tiles,
		prevOneHot: initPrevOneHot(batch),
		rowSeen:    make([]bool, inDim),
	}, nil
}

// Close releases the compiled kernel.
func (e *Executor) Close() {
	if e == nil {
		return
	}
	e.mu.Lock()
	tiles := e.tiles
	e.tiles = nil
	e.mu.Unlock()
	for i := range tiles {
		if tiles[i].k != nil {
			tiles[i].k.Close()
		}
	}
}

// Eval computes y = x*w and returns a new output slice.
func (e *Executor) Eval(x, w []float32) ([]float32, error) {
	dst := make([]float32, e.outputLen())
	_, err := e.evalInto(dst, x, w, false)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// EvalWithStats computes y = x*w and returns a new output slice plus timing.
func (e *Executor) EvalWithStats(x, w []float32) ([]float32, EvalStats, error) {
	dst := make([]float32, e.outputLen())
	st, err := e.evalInto(dst, x, w, true)
	if err != nil {
		return nil, st, err
	}
	return dst, st, nil
}

// EvalInto computes y = x*w into dst.
func (e *Executor) EvalInto(dst, x, w []float32) (EvalStats, error) {
	return e.evalInto(dst, x, w, true)
}

func (e *Executor) evalInto(dst, x, w []float32, collectMetrics bool) (EvalStats, error) {
	if e == nil {
		return EvalStats{}, fmt.Errorf("dynamic matmul: executor is nil")
	}
	if err := e.validateIO(dst, x, w); err != nil {
		return EvalStats{}, err
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.tiles) == 0 {
		return EvalStats{}, fmt.Errorf("dynamic matmul: executor is closed")
	}

	var hwNS uint64
	var metrics map[string]float64
	for i := range e.tiles {
		tile := &e.tiles[i]
		packInputTile(tile.inputPacked, x, w, e.batch, e.inDim, e.outDim, tile.outOffset, tile.outDim)
		if err := tile.k.WriteInputF32(0, tile.inputPacked); err != nil {
			return EvalStats{}, fmt.Errorf("dynamic matmul: write input tile %d: %w", i, err)
		}
		evalStart := time.Now()
		tileHW := uint64(0)
		if collectMetrics {
			st, err := tile.k.EvalWithStats()
			if err != nil {
				return EvalStats{}, fmt.Errorf("dynamic matmul: eval tile %d: %w", i, err)
			}
			tileHW = st.HWExecutionNS
			metrics = addEvalMetrics(metrics, st.Metrics)
		} else {
			if err := tile.k.Eval(); err != nil {
				return EvalStats{}, fmt.Errorf("dynamic matmul: eval tile %d: %w", i, err)
			}
		}
		if err := tile.k.ReadOutputF32(0, tile.outputPacked); err != nil {
			return EvalStats{}, fmt.Errorf("dynamic matmul: read output tile %d: %w", i, err)
		}
		unpackOutputTile(dst, tile.outputPacked, e.batch, e.outDim, tile.outOffset, tile.outDim)
		if tileHW != 0 {
			hwNS += tileHW
		} else {
			hwNS += uint64(time.Since(evalStart).Nanoseconds())
		}
	}
	return EvalStats{HWExecutionNS: hwNS, Metrics: metrics}, nil
}

// EvalCF evaluates a channel-first input tensor against previously primed
// weights and leaves the output resident in the ANE output surfaces.
//
// xCF has shape [inDim, batch].
func (e *Executor) EvalCF(xCF []float32) (EvalStats, error) {
	if e == nil {
		return EvalStats{}, fmt.Errorf("dynamic matmul: executor is nil")
	}
	if len(xCF) != e.inDim*e.batch {
		return EvalStats{}, fmt.Errorf("dynamic matmul: channel-first input length=%d want=%d", len(xCF), e.inDim*e.batch)
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.tiles) == 0 {
		return EvalStats{}, fmt.Errorf("dynamic matmul: executor is closed")
	}
	if !e.weightsReady {
		return EvalStats{}, fmt.Errorf("dynamic matmul: weights are not primed")
	}
	return e.evalCFLocked(xCF, true)
}

// EvalCFHW evaluates a channel-first input tensor against previously primed
// weights and returns only aggregate hardware execution time.
func (e *Executor) EvalCFHW(xCF []float32) (uint64, error) {
	if e == nil {
		return 0, fmt.Errorf("dynamic matmul: executor is nil")
	}
	if len(xCF) != e.inDim*e.batch {
		return 0, fmt.Errorf("dynamic matmul: channel-first input length=%d want=%d", len(xCF), e.inDim*e.batch)
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.tiles) == 0 {
		return 0, fmt.Errorf("dynamic matmul: executor is closed")
	}
	if !e.weightsReady {
		return 0, fmt.Errorf("dynamic matmul: weights are not primed")
	}
	st, err := e.evalCFLocked(xCF, false)
	if err != nil {
		return 0, err
	}
	return st.HWExecutionNS, nil
}

func (e *Executor) evalCFLocked(xCF []float32, collectMetrics bool) (EvalStats, error) {
	var hwNS uint64
	var metrics map[string]float64
	for i := range e.tiles {
		tile := &e.tiles[i]
		stageChannelFirstActivations(tile.inputPacked, xCF, e.batch, e.inDim, tile.outDim)
		if err := tile.k.WriteInputF32(0, tile.inputPacked); err != nil {
			return EvalStats{}, fmt.Errorf("dynamic matmul: write channel-first input tile %d: %w", i, err)
		}
		evalStart := time.Now()
		var tileHW uint64
		if collectMetrics {
			st, err := tile.k.EvalWithStats()
			if err != nil {
				return EvalStats{}, fmt.Errorf("dynamic matmul: eval channel-first tile %d: %w", i, err)
			}
			tileHW = st.HWExecutionNS
			metrics = addEvalMetrics(metrics, st.Metrics)
		} else {
			if err := tile.k.Eval(); err != nil {
				return EvalStats{}, fmt.Errorf("dynamic matmul: eval channel-first tile %d: %w", i, err)
			}
		}
		if tileHW != 0 {
			hwNS += tileHW
		} else {
			hwNS += uint64(time.Since(evalStart).Nanoseconds())
		}
	}
	return EvalStats{HWExecutionNS: hwNS, Metrics: metrics}, nil
}

// ReadOutputCF reads the last evaluated output tensor in channel-first order.
//
// dstCF has shape [outDim, batch].
func (e *Executor) ReadOutputCF(dstCF []float32) error {
	if e == nil {
		return fmt.Errorf("dynamic matmul: executor is nil")
	}
	if len(dstCF) != e.outputLen() {
		return fmt.Errorf("dynamic matmul: channel-first output length=%d want=%d", len(dstCF), e.outputLen())
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.tiles) == 0 {
		return fmt.Errorf("dynamic matmul: executor is closed")
	}

	for i := range e.tiles {
		tile := &e.tiles[i]
		if err := tile.k.ReadOutputF32(0, tile.outputPacked); err != nil {
			return fmt.Errorf("dynamic matmul: read channel-first output tile %d: %w", i, err)
		}
		copyOutputTileCF(dstCF, tile.outputPacked, e.batch, e.outDim, tile.outOffset, tile.outDim)
	}
	return nil
}

// EvalCFIOInto evaluates a channel-first input tensor into a channel-first
// output tensor using previously primed weights.
//
// xCF has shape [inDim, batch] and dstCF has shape [outDim, batch].
func (e *Executor) EvalCFIOInto(dstCF, xCF []float32) (EvalStats, error) {
	st, err := e.EvalCF(xCF)
	if err != nil {
		return EvalStats{}, err
	}
	if err := e.ReadOutputCF(dstCF); err != nil {
		return EvalStats{}, err
	}
	return st, nil
}

// EvalCFIOIntoHW evaluates a channel-first input tensor into a channel-first
// output tensor using previously primed weights and returns only aggregate
// hardware execution time.
func (e *Executor) EvalCFIOIntoHW(dstCF, xCF []float32) (uint64, error) {
	hwNS, err := e.EvalCFHW(xCF)
	if err != nil {
		return 0, err
	}
	if err := e.ReadOutputCF(dstCF); err != nil {
		return 0, err
	}
	return hwNS, nil
}

// PrimeWeightsIO copies the full IO-layout weight matrix into the cached ANE
// input buffers. wIO must be laid out as [inDim, outDim].
func (e *Executor) PrimeWeightsIO(wIO []float32) error {
	if e == nil {
		return fmt.Errorf("dynamic matmul: executor is nil")
	}
	if len(wIO) != e.inDim*e.outDim {
		return fmt.Errorf("dynamic matmul: io weight length=%d want=%d", len(wIO), e.inDim*e.outDim)
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.tiles) == 0 {
		return fmt.Errorf("dynamic matmul: executor is closed")
	}
	for i := range e.tiles {
		primeTileWeights(&e.tiles[i], wIO, e.batch, e.inDim, e.outDim)
		if err := writeFullTileInput(&e.tiles[i]); err != nil {
			return err
		}
	}
	e.weightsReady = true
	return nil
}

// UpdateWeightsIORows patches a subset of rows in the cached IO-layout weight
// buffers. rows contains input-channel row ids in [0, inDim).
func (e *Executor) UpdateWeightsIORows(wIO []float32, rows []int) error {
	if e == nil {
		return fmt.Errorf("dynamic matmul: executor is nil")
	}
	if len(wIO) != e.inDim*e.outDim {
		return fmt.Errorf("dynamic matmul: io weight length=%d want=%d", len(wIO), e.inDim*e.outDim)
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.tiles) == 0 {
		return fmt.Errorf("dynamic matmul: executor is closed")
	}
	if !e.weightsReady {
		return fmt.Errorf("dynamic matmul: weights are not primed")
	}
	for _, row := range rows {
		if row < 0 || row >= e.inDim {
			return fmt.Errorf("dynamic matmul: weight row %d out of range [0,%d)", row, e.inDim)
		}
	}
	for i := range e.tiles {
		updateTileWeightRows(&e.tiles[i], wIO, rows, e.batch, e.outDim)
		if err := writeTileRows(&e.tiles[i], rows); err != nil {
			return err
		}
	}
	return nil
}

// EvalOneHotIOInto computes y = x*w for one-hot activations encoded by xs and
// a previously primed IO-layout weight matrix.
//
// xs holds at most batch token ids in [0, inDim). Position t selects the input
// row for batch element t. Remaining batch positions are treated as zero input.
func (e *Executor) EvalOneHotIOInto(dst []float32, xs []int) (EvalStats, error) {
	return e.evalOneHotIOInto(dst, xs, true)
}

// EvalOneHotIOIntoHW computes y = x*w for one-hot activations encoded by xs
// and a previously primed IO-layout weight matrix, returning only aggregate
// hardware execution time.
func (e *Executor) EvalOneHotIOIntoHW(dst []float32, xs []int) (uint64, error) {
	st, err := e.evalOneHotIOInto(dst, xs, false)
	if err != nil {
		return 0, err
	}
	return st.HWExecutionNS, nil
}

func (e *Executor) evalOneHotIOInto(dst []float32, xs []int, collectMetrics bool) (EvalStats, error) {
	if e == nil {
		return EvalStats{}, fmt.Errorf("dynamic matmul: executor is nil")
	}
	if len(xs) > e.batch {
		return EvalStats{}, fmt.Errorf("dynamic matmul: one-hot batch len=%d want <= %d", len(xs), e.batch)
	}
	logicalBatch := e.batch
	switch len(dst) {
	case e.outputLen():
		logicalBatch = e.batch
	case len(xs) * e.outDim:
		logicalBatch = len(xs)
	default:
		return EvalStats{}, fmt.Errorf("dynamic matmul: output length=%d want=%d or %d", len(dst), e.outputLen(), len(xs)*e.outDim)
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.tiles) == 0 {
		return EvalStats{}, fmt.Errorf("dynamic matmul: executor is closed")
	}
	if !e.weightsReady {
		return EvalStats{}, fmt.Errorf("dynamic matmul: weights are not primed")
	}
	for _, x := range xs {
		if x < 0 || x >= e.inDim {
			return EvalStats{}, fmt.Errorf("dynamic matmul: one-hot index %d out of range [0,%d)", x, e.inDim)
		}
	}

	rows := e.touchedOneHotRows(xs)
	defer e.clearTouchedRows(rows)

	var hwNS uint64
	var metrics map[string]float64
	for i := range e.tiles {
		tile := &e.tiles[i]
		stageOneHotActivations(tile.inputPacked, e.prevOneHot, xs, e.batch, tile.outDim)
		if err := writeTileRows(tile, rows); err != nil {
			return EvalStats{}, fmt.Errorf("dynamic matmul: write one-hot tile %d: %w", i, err)
		}
		evalStart := time.Now()
		tileHW := uint64(0)
		if collectMetrics {
			st, err := tile.k.EvalWithStats()
			if err != nil {
				return EvalStats{}, fmt.Errorf("dynamic matmul: eval one-hot tile %d: %w", i, err)
			}
			tileHW = st.HWExecutionNS
			metrics = addEvalMetrics(metrics, st.Metrics)
		} else {
			if err := tile.k.Eval(); err != nil {
				return EvalStats{}, fmt.Errorf("dynamic matmul: eval one-hot tile %d: %w", i, err)
			}
		}
		if err := tile.k.ReadOutputF32(0, tile.outputPacked); err != nil {
			return EvalStats{}, fmt.Errorf("dynamic matmul: read one-hot tile %d: %w", i, err)
		}
		unpackOutputTileRows(dst, tile.outputPacked, logicalBatch, e.batch, e.outDim, tile.outOffset, tile.outDim)
		if tileHW != 0 {
			hwNS += tileHW
		} else {
			hwNS += uint64(time.Since(evalStart).Nanoseconds())
		}
	}
	updatePrevOneHot(e.prevOneHot, xs)
	return EvalStats{HWExecutionNS: hwNS, Metrics: metrics}, nil
}

func addEvalMetrics(dst, src map[string]float64) map[string]float64 {
	if len(src) == 0 {
		return dst
	}
	if dst == nil {
		dst = make(map[string]float64, len(src))
	}
	for k, v := range src {
		dst[k] += v
	}
	return dst
}

// CopyOutputToInput copies the last evaluated output tensor into a destination
// kernel input without converting through Go-managed float buffers.
func (e *Executor) CopyOutputToInput(dst *model.Kernel, dstInput, dstChannel int) error {
	if e == nil {
		return fmt.Errorf("dynamic matmul: executor is nil")
	}
	if dst == nil {
		return fmt.Errorf("dynamic matmul: destination kernel is nil")
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	if len(e.tiles) == 0 {
		return fmt.Errorf("dynamic matmul: executor is closed")
	}

	for i := range e.tiles {
		tile := &e.tiles[i]
		if err := model.CopyOutputChannelsToInput(dst, dstInput, dstChannel+tile.outOffset, tile.k, 0, 0, tile.outDim); err != nil {
			return fmt.Errorf("dynamic matmul: copy output tile %d: %w", i, err)
		}
	}
	return nil
}

func (e *Executor) validateIO(dst, x, w []float32) error {
	if got, want := len(x), e.batch*e.inDim; got != want {
		return fmt.Errorf("dynamic matmul: input length=%d want=%d", got, want)
	}
	if got, want := len(w), e.inDim*e.outDim; got != want {
		return fmt.Errorf("dynamic matmul: weight length=%d want=%d", got, want)
	}
	if got, want := len(dst), e.outputLen(); got != want {
		return fmt.Errorf("dynamic matmul: output length=%d want=%d", got, want)
	}
	return nil
}

func (e *Executor) outputLen() int {
	if e == nil {
		return 0
	}
	return e.batch * e.outDim
}

func packInput(dst, x, w []float32, batch, inDim, outDim int) {
	packInputTile(dst, x, w, batch, inDim, outDim, 0, outDim)
}

func packInputTile(dst, x, w []float32, batch, inDim, fullOutDim, outOffset, tileOutDim int) {
	rowWidth := batch + tileOutDim
	for d := 0; d < inDim; d++ {
		row := dst[d*rowWidth : (d+1)*rowWidth]
		for t := 0; t < batch; t++ {
			row[t] = x[t*inDim+d]
		}
		copy(row[batch:], w[d*fullOutDim+outOffset:d*fullOutDim+outOffset+tileOutDim])
	}
}

func stageChannelFirstActivations(dst, xCF []float32, batch, inDim, tileOutDim int) {
	rowWidth := batch + tileOutDim
	for d := 0; d < inDim; d++ {
		row := dst[d*rowWidth : (d+1)*rowWidth]
		copy(row[:batch], xCF[d*batch:(d+1)*batch])
	}
}

func stageOneHotActivations(dst []float32, prev, cur []int, batch, tileOutDim int) {
	rowWidth := batch + tileOutDim
	for t, idx := range prev {
		if idx >= 0 {
			dst[idx*rowWidth+t] = 0
		}
	}
	for t, idx := range cur {
		dst[idx*rowWidth+t] = 1
	}
}

func unpackOutput(dst, src []float32, batch, outDim int) {
	unpackOutputTile(dst, src, batch, outDim, 0, outDim)
}

func unpackOutputTile(dst, src []float32, batch, fullOutDim, outOffset, tileOutDim int) {
	unpackOutputTileRows(dst, src, batch, batch, fullOutDim, outOffset, tileOutDim)
}

func unpackOutputTileRows(dst, src []float32, rows, batch, fullOutDim, outOffset, tileOutDim int) {
	for c := 0; c < tileOutDim; c++ {
		row := src[c*batch : (c+1)*batch]
		for t := 0; t < rows; t++ {
			dst[t*fullOutDim+outOffset+c] = row[t]
		}
	}
}

func copyOutputTileCF(dst, src []float32, batch, fullOutDim, outOffset, tileOutDim int) {
	for c := 0; c < tileOutDim; c++ {
		copy(dst[(outOffset+c)*batch:(outOffset+c+1)*batch], src[c*batch:(c+1)*batch])
	}
}

func compileTiles(batch, inDim, outDim int, qos uint32, tileOut int) ([]tile, error) {
	if tileOut > 0 {
		tiles, err := compileTiled(batch, inDim, outDim, qos, tileOut)
		if err != nil {
			return nil, fmt.Errorf("dynamic matmul: compile tiled kernels: %w", err)
		}
		return tiles, nil
	}

	full, err := compileTile(batch, inDim, outDim, qos, 0)
	if err == nil {
		return []tile{full}, nil
	}
	fullErr := err

	var lastErr error
	for _, cand := range defaultTileCandidates(outDim) {
		tiles, terr := compileTiled(batch, inDim, outDim, qos, cand)
		if terr == nil {
			return tiles, nil
		}
		lastErr = terr
	}
	if lastErr == nil {
		lastErr = fullErr
	}
	return nil, fmt.Errorf("dynamic matmul: full compile failed: %w; tile fallback failed: %v", fullErr, lastErr)
}

func compileTiled(batch, inDim, outDim int, qos uint32, tileOut int) ([]tile, error) {
	if tileOut <= 0 {
		return nil, fmt.Errorf("tileOut=%d must be > 0", tileOut)
	}
	var tiles []tile
	for off := 0; off < outDim; off += tileOut {
		thisOut := tileOut
		if rem := outDim - off; rem < thisOut {
			thisOut = rem
		}
		t, err := compileTile(batch, inDim, thisOut, qos, off)
		if err != nil {
			for i := range tiles {
				if tiles[i].k != nil {
					tiles[i].k.Close()
				}
			}
			return nil, err
		}
		tiles = append(tiles, t)
	}
	return tiles, nil
}

func compileTile(batch, inDim, outDim int, qos uint32, outOffset int) (tile, error) {
	k, err := model.Compile(model.CompileOptions{
		MILText:     mil.GenDynamicMatmul(inDim, outDim, batch),
		SharedModel: true,
		QoS:         qos,
	})
	if err != nil {
		return tile{}, err
	}
	return tile{
		outOffset:    outOffset,
		outDim:       outDim,
		k:            k,
		inputPacked:  make([]float32, inDim*(batch+outDim)),
		outputPacked: make([]float32, outDim*batch),
	}, nil
}

func primeTileWeights(tile *tile, wIO []float32, batch, inDim, fullOutDim int) {
	rowWidth := batch + tile.outDim
	for in := 0; in < inDim; in++ {
		row := tile.inputPacked[in*rowWidth : (in+1)*rowWidth]
		copy(row[batch:], wIO[in*fullOutDim+tile.outOffset:in*fullOutDim+tile.outOffset+tile.outDim])
	}
}

func updateTileWeightRows(tile *tile, wIO []float32, rows []int, batch, fullOutDim int) {
	rowWidth := batch + tile.outDim
	for _, in := range rows {
		row := tile.inputPacked[in*rowWidth : (in+1)*rowWidth]
		copy(row[batch:], wIO[in*fullOutDim+tile.outOffset:in*fullOutDim+tile.outOffset+tile.outDim])
	}
}

func defaultTileCandidates(outDim int) []int {
	base := []int{384, 256, 128, 64}
	out := make([]int, 0, len(base)+1)
	seen := make(map[int]bool, len(base)+1)
	if outDim > 0 {
		c := minInt(outDim, 512)
		if c > 0 && c < outDim {
			out = append(out, c)
			seen[c] = true
		}
	}
	for _, c := range base {
		if c <= 0 || c >= outDim || seen[c] {
			continue
		}
		out = append(out, c)
		seen[c] = true
	}
	if len(out) == 0 && outDim > 1 {
		out = append(out, outDim-1)
	}
	return out
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initPrevOneHot(batch int) []int {
	prev := make([]int, batch)
	for i := range prev {
		prev[i] = -1
	}
	return prev
}

func updatePrevOneHot(prev []int, cur []int) {
	for i := range prev {
		prev[i] = -1
	}
	for i, x := range cur {
		prev[i] = x
	}
}

func (e *Executor) touchedOneHotRows(cur []int) []int {
	rows := e.touchedRows[:0]
	for _, x := range e.prevOneHot {
		if x >= 0 && !e.rowSeen[x] {
			e.rowSeen[x] = true
			rows = append(rows, x)
		}
	}
	for _, x := range cur {
		if x >= 0 && !e.rowSeen[x] {
			e.rowSeen[x] = true
			rows = append(rows, x)
		}
	}
	e.touchedRows = rows
	return rows
}

func (e *Executor) clearTouchedRows(rows []int) {
	for _, x := range rows {
		e.rowSeen[x] = false
	}
}
