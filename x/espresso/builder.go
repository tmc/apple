package espresso

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"maps"
	"math"
)

// ElementwiseOp identifies an Espresso elementwise operation by its numeric code.
type ElementwiseOp int

// Espresso elementwise operation codes.
const (
	ElementwiseAdd  ElementwiseOp = 0
	ElementwiseMul  ElementwiseOp = 1
	ElementwiseMax  ElementwiseOp = 2
	ElementwiseMin  ElementwiseOp = 3
	ElementwiseSub  ElementwiseOp = 5
	ElementwiseDiv  ElementwiseOp = 6
	ElementwiseSqrt ElementwiseOp = 10
	ElementwiseExp  ElementwiseOp = 11
	ElementwiseRsqrt ElementwiseOp = 12
	ElementwiseLog  ElementwiseOp = 13
	ElementwisePow  ElementwiseOp = 14
	ElementwiseAbs  ElementwiseOp = 15
)

// Builder constructs Espresso IR JSON (model.espresso.net + model.espresso.shape)
// and optional weight binaries for inference on Apple's Espresso framework.
type Builder struct {
	formatVersion int
	layers        []layerDef
	weightIdx     int
	weightData    [][]byte
}

type layerDef struct {
	Name       string
	Type       string
	Bottom     []string
	Top        []string
	Attrs      map[string]any
	WeightIdx  int // -1 if no weights
	BiasIdx    int // -1 if no bias
}

// LayerOption configures a layer definition.
type LayerOption func(*layerDef)

// WithBottom sets the input blob names for a layer.
func WithBottom(names ...string) LayerOption {
	return func(l *layerDef) { l.Bottom = names }
}

// WithTop sets the output blob names for a layer.
func WithTop(names ...string) LayerOption {
	return func(l *layerDef) { l.Top = names }
}

// WithWeights assigns float32 weight data to a layer.
func WithWeights(data []float32) LayerOption {
	return func(l *layerDef) {
		l.Attrs["_weight_data"] = data
	}
}

// WithBias assigns float32 bias data to a layer.
func WithBias(data []float32) LayerOption {
	return func(l *layerDef) {
		l.Attrs["_bias_data"] = data
	}
}

// WithPad sets padding [top, bottom, left, right].
func WithPad(t, b, l, r int) LayerOption {
	return func(ld *layerDef) {
		ld.Attrs["pad_t"] = t
		ld.Attrs["pad_b"] = b
		ld.Attrs["pad_l"] = l
		ld.Attrs["pad_r"] = r
	}
}

// WithStride sets stride [height, width].
func WithStride(h, w int) LayerOption {
	return func(l *layerDef) {
		l.Attrs["stride_x"] = w
		l.Attrs["stride_y"] = h
	}
}

// WithGroups sets the convolution groups.
func WithGroups(g int) LayerOption {
	return func(l *layerDef) { l.Attrs["n_groups"] = g }
}

// WithHasReLU enables fused ReLU on the layer.
func WithHasReLU(v bool) LayerOption {
	return func(l *layerDef) {
		if v {
			l.Attrs["has_relu"] = 1
			l.Attrs["fused_relu"] = 1
		}
	}
}

// NewBuilder creates a new Espresso IR builder with format version 200.
func NewBuilder() *Builder {
	return &Builder{
		formatVersion: 200,
	}
}

func (b *Builder) addLayer(name, typ string, opts []LayerOption) *Builder {
	l := layerDef{
		Name:      name,
		Type:      typ,
		WeightIdx: -1,
		BiasIdx:   -1,
		Attrs:     make(map[string]any),
	}
	for _, o := range opts {
		o(&l)
	}
	// Auto-connect: if no bottom specified, use previous layer's top.
	if len(l.Bottom) == 0 && len(b.layers) > 0 {
		prev := b.layers[len(b.layers)-1]
		l.Bottom = prev.Top
	}
	// Auto-name top if not specified.
	if len(l.Top) == 0 {
		l.Top = []string{name}
	}
	b.layers = append(b.layers, l)
	return b
}

// InnerProduct adds a fully connected layer.
func (b *Builder) InnerProduct(name string, nOutput int, opts ...LayerOption) *Builder {
	opts = append([]LayerOption{func(l *layerDef) {
		l.Attrs["nO"] = nOutput
	}}, opts...)
	return b.addLayer(name, "inner_product", opts)
}

// Convolution adds a convolution layer.
func (b *Builder) Convolution(name string, nOutput, kernelH, kernelW int, opts ...LayerOption) *Builder {
	opts = append([]LayerOption{func(l *layerDef) {
		l.Attrs["nO"] = nOutput
		l.Attrs["Nx"] = kernelW
		l.Attrs["Ny"] = kernelH
	}}, opts...)
	return b.addLayer(name, "convolution", opts)
}

// Activation adds an activation layer.
func (b *Builder) Activation(name, mode string, opts ...LayerOption) *Builder {
	opts = append([]LayerOption{func(l *layerDef) {
		l.Attrs["mode"] = mode
	}}, opts...)
	return b.addLayer(name, "activation", opts)
}

// Softmax adds a softmax layer.
func (b *Builder) Softmax(name string, opts ...LayerOption) *Builder {
	return b.addLayer(name, "softmax", opts)
}

// Elementwise adds an elementwise operation layer with a string operation name.
func (b *Builder) Elementwise(name, operation string, opts ...LayerOption) *Builder {
	opts = append([]LayerOption{func(l *layerDef) {
		l.Attrs["operation"] = operation
	}}, opts...)
	return b.addLayer(name, "elementwise", opts)
}

// ElementwiseOped adds an elementwise operation layer with a typed operation code.
func (b *Builder) ElementwiseOped(name string, op ElementwiseOp, opts ...LayerOption) *Builder {
	opts = append([]LayerOption{func(l *layerDef) {
		l.Attrs["operation"] = int(op)
	}}, opts...)
	return b.addLayer(name, "elementwise", opts)
}

// BatchNorm adds a batch normalization layer.
func (b *Builder) BatchNorm(name string, channels int, opts ...LayerOption) *Builder {
	opts = append([]LayerOption{func(l *layerDef) {
		l.Attrs["n_channels"] = channels
	}}, opts...)
	return b.addLayer(name, "batch_norm", opts)
}

// Concat adds a concatenation layer.
func (b *Builder) Concat(name string, opts ...LayerOption) *Builder {
	return b.addLayer(name, "concat", opts)
}

// Reshape adds a reshape layer.
func (b *Builder) Reshape(name string, shape [4]int, opts ...LayerOption) *Builder {
	opts = append([]LayerOption{func(l *layerDef) {
		l.Attrs["shape"] = []int{shape[0], shape[1], shape[2], shape[3]}
	}}, opts...)
	return b.addLayer(name, "reshape", opts)
}

// Pooling adds a pooling layer.
func (b *Builder) Pooling(name, mode string, kernelH, kernelW int, opts ...LayerOption) *Builder {
	opts = append([]LayerOption{func(l *layerDef) {
		l.Attrs["mode"] = mode
		l.Attrs["Nx"] = kernelW
		l.Attrs["Ny"] = kernelH
	}}, opts...)
	return b.addLayer(name, "pooling", opts)
}

// Build produces the Espresso IR JSON files (net and shape).
func (b *Builder) Build() (netJSON []byte, shapeJSON []byte, err error) {
	if len(b.layers) == 0 {
		return nil, nil, fmt.Errorf("espresso builder: no layers")
	}

	net, shape := b.buildMaps()
	netJSON, err = json.MarshalIndent(net, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("espresso builder: marshal net: %w", err)
	}
	shapeJSON, err = json.MarshalIndent(shape, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("espresso builder: marshal shape: %w", err)
	}
	return netJSON, shapeJSON, nil
}

// BuildWithWeights produces IR JSON and a weight binary.
func (b *Builder) BuildWithWeights() (netJSON, shapeJSON, weightsBin []byte, err error) {
	netJSON, shapeJSON, err = b.Build()
	if err != nil {
		return nil, nil, nil, err
	}

	if len(b.weightData) == 0 {
		return netJSON, shapeJSON, nil, nil
	}

	// Concatenate all weight segments.
	total := 0
	for _, d := range b.weightData {
		total += len(d)
	}
	weightsBin = make([]byte, total)
	off := 0
	for _, d := range b.weightData {
		copy(weightsBin[off:], d)
		off += len(d)
	}
	return netJSON, shapeJSON, weightsBin, nil
}

func (b *Builder) buildMaps() (net map[string]any, shape map[string]any) {
	layers := make(map[string]any)

	// Assign weight indices and collect weight data.
	b.weightData = nil
	b.weightIdx = 0
	for i := range b.layers {
		l := &b.layers[i]
		if wdata, ok := l.Attrs["_weight_data"]; ok {
			l.WeightIdx = b.weightIdx
			b.weightIdx++
			b.weightData = append(b.weightData, float32ToBytes(wdata.([]float32)))
			delete(l.Attrs, "_weight_data")
		}
		if bdata, ok := l.Attrs["_bias_data"]; ok {
			l.BiasIdx = b.weightIdx
			b.weightIdx++
			b.weightData = append(b.weightData, float32ToBytes(bdata.([]float32)))
			delete(l.Attrs, "_bias_data")
		}
	}

	for _, l := range b.layers {
		entry := map[string]any{
			"type":   l.Type,
			"bottom": l.Bottom,
			"top":    l.Top,
		}
		if l.WeightIdx >= 0 {
			entry["blob_weights"] = l.WeightIdx
		}
		if l.BiasIdx >= 0 {
			entry["blob_biases"] = l.BiasIdx
		}
		maps.Copy(entry, l.Attrs)
		layers[l.Name] = entry
	}

	net = map[string]any{
		"format_version": b.formatVersion,
		"layers":         layers,
	}

	// Build shape dict: map layer names to empty shape placeholders.
	// The actual shapes are computed by the Espresso runtime.
	shapeEntries := make(map[string]any)
	for _, l := range b.layers {
		for _, t := range l.Top {
			shapeEntries[t] = map[string]any{}
		}
	}
	shape = map[string]any{
		"layer_shapes": shapeEntries,
	}

	return net, shape
}

func float32ToBytes(data []float32) []byte {
	buf := make([]byte, len(data)*4)
	for i, v := range data {
		binary.LittleEndian.PutUint32(buf[i*4:], math.Float32bits(v))
	}
	return buf
}
