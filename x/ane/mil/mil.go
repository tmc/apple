//go:build darwin

package mil

import (
	"encoding/binary"
	"fmt"
	"math"
)

// buildInfo matches the format expected by the ANE MIL parser.
const buildInfo = `[buildInfo = dict<string, string>({{"coremlc-component-MIL", "3510.2.1"}, {"coremlc-version", "3505.4.1"}, {"coremltools-component-milinternal", ""}, {"coremltools-version", "9.0"}})]`

// GenConv generates a MIL text for a 1×1 convolution kernel with fp16 internal computation.
// inCh and outCh are channel counts; spatial is the spatial dimension (1 for vectors).
func GenConv(inCh, outCh, spatial int) string {
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>(tensor<fp32, [1, %d, 1, %d]> x) {
        string c_pad_type = const()[name = string("c_pad_type"), val = string("valid")];
        tensor<int32, [2]> c_strides = const()[name = string("c_strides"), val = tensor<int32, [2]>([1, 1])];
        tensor<int32, [4]> c_pad = const()[name = string("c_pad"), val = tensor<int32, [4]>([0, 0, 0, 0])];
        tensor<int32, [2]> c_dilations = const()[name = string("c_dilations"), val = tensor<int32, [2]>([1, 1])];
        int32 c_groups = const()[name = string("c_groups"), val = int32(1)];
        string to_fp16 = const()[name = string("to_fp16"), val = string("fp16")];
        tensor<fp16, [1, %d, 1, %d]> x16 = cast(dtype = to_fp16, x = x)[name = string("cast_in")];
        tensor<fp16, [%d, %d, 1, 1]> W = const()[name = string("W"), val = tensor<fp16, [%d, %d, 1, 1]>(BLOBFILE(path = string("@model_path/weights/weight.bin"), offset = uint64(64)))];
        tensor<fp16, [1, %d, 1, %d]> y16 = conv(dilations = c_dilations, groups = c_groups, pad = c_pad, pad_type = c_pad_type, strides = c_strides, weight = W, x = x16)[name = string("conv")];
        string to_fp32 = const()[name = string("to_fp32"), val = string("fp32")];
        tensor<fp32, [1, %d, 1, %d]> y = cast(dtype = to_fp32, x = y16)[name = string("cast_out")];
    } -> (y);
}
`, buildInfo,
		inCh, spatial, // input shape
		inCh, spatial, // x16 shape
		outCh, inCh, // weight shape
		outCh, inCh, // weight shape (repeated in BLOBFILE ref)
		outCh, spatial, // y16 shape
		outCh, spatial, // y shape
	)
}

// GenConvFP32 generates a MIL text for a 1×1 convolution with fp32 weights (no casting).
func GenConvFP32(inCh, outCh, spatial int) string {
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>(tensor<fp32, [1, %d, 1, %d]> x) {
        string c_pad_type = const()[name = string("c_pad_type"), val = string("valid")];
        tensor<int32, [2]> c_strides = const()[name = string("c_strides"), val = tensor<int32, [2]>([1, 1])];
        tensor<int32, [4]> c_pad = const()[name = string("c_pad"), val = tensor<int32, [4]>([0, 0, 0, 0])];
        tensor<int32, [2]> c_dilations = const()[name = string("c_dilations"), val = tensor<int32, [2]>([1, 1])];
        int32 c_groups = const()[name = string("c_groups"), val = int32(1)];
        tensor<fp16, [%d, %d, 1, 1]> W = const()[name = string("W"), val = tensor<fp16, [%d, %d, 1, 1]>(BLOBFILE(path = string("@model_path/weights/weight.bin"), offset = uint64(64)))];
        string to_fp16 = const()[name = string("to_fp16"), val = string("fp16")];
        tensor<fp16, [1, %d, 1, %d]> x16 = cast(dtype = to_fp16, x = x)[name = string("cast_in")];
        tensor<fp16, [1, %d, 1, %d]> y16 = conv(dilations = c_dilations, groups = c_groups, pad = c_pad, pad_type = c_pad_type, strides = c_strides, weight = W, x = x16)[name = string("conv")];
        string to_fp32 = const()[name = string("to_fp32"), val = string("fp32")];
        tensor<fp32, [1, %d, 1, %d]> y = cast(dtype = to_fp32, x = y16)[name = string("cast_out")];
    } -> (y);
}
`, buildInfo,
		inCh, spatial, // input
		outCh, inCh, // weight shape
		outCh, inCh, // weight BLOBFILE ref
		inCh, spatial, // x16
		outCh, spatial, // y16
		outCh, spatial, // y
	)
}

// GenMatmul generates a MIL text for a matrix multiplication as a 1×1 convolution.
// This is equivalent to y = x @ W^T for [batch, inCh] -> [batch, outCh].
func GenMatmul(inCh, outCh, spatial int) string {
	return GenConv(inCh, outCh, spatial)
}

// GenIdentity generates a MIL text for an identity operation (1×1 conv with identity weights).
func GenIdentity(channels, spatial int) string {
	return GenConv(channels, channels, spatial)
}

// GenIdentityFP16IO generates a MIL text for an identity operation with fp16 I/O
// (no fp32 casts, the ANE reads and writes fp16 directly).
func GenIdentityFP16IO(channels, spatial int) string {
	return GenConvFP16IO(channels, channels, spatial)
}

// GenConvFP16IO generates a MIL text for a 1×1 convolution with fp16 I/O (no casts).
func GenConvFP16IO(inCh, outCh, spatial int) string {
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {
        string c_pad_type = const()[name = string("c_pad_type"), val = string("valid")];
        tensor<int32, [2]> c_strides = const()[name = string("c_strides"), val = tensor<int32, [2]>([1, 1])];
        tensor<int32, [4]> c_pad = const()[name = string("c_pad"), val = tensor<int32, [4]>([0, 0, 0, 0])];
        tensor<int32, [2]> c_dilations = const()[name = string("c_dilations"), val = tensor<int32, [2]>([1, 1])];
        int32 c_groups = const()[name = string("c_groups"), val = int32(1)];
        tensor<fp16, [%d, %d, 1, 1]> W = const()[name = string("W"), val = tensor<fp16, [%d, %d, 1, 1]>(BLOBFILE(path = string("@model_path/weights/weight.bin"), offset = uint64(64)))];
        tensor<fp16, [1, %d, 1, %d]> y = conv(dilations = c_dilations, groups = c_groups, pad = c_pad, pad_type = c_pad_type, strides = c_strides, weight = W, x = x)[name = string("conv")];
    } -> (y);
}
`, buildInfo,
		inCh, spatial, // input shape
		outCh, inCh, // weight shape
		outCh, inCh, // weight shape (repeated)
		outCh, spatial, // output shape
	)
}

// GenScaleFP16IO generates a MIL text for a simple multiplication (1 channel, S spatial).
// Each spatial element is multiplied by the scalar weight.
func GenScaleFP16IO(spatial int) string {
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>(tensor<fp16, [1, 1, 1, %d]> x) {
        string c_pad_type = const()[name = string("c_pad_type"), val = string("valid")];
        tensor<int32, [2]> c_strides = const()[name = string("c_strides"), val = tensor<int32, [2]>([1, 1])];
        tensor<int32, [4]> c_pad = const()[name = string("c_pad"), val = tensor<int32, [4]>([0, 0, 0, 0])];
        tensor<int32, [2]> c_dilations = const()[name = string("c_dilations"), val = tensor<int32, [2]>([1, 1])];
        int32 c_groups = const()[name = string("c_groups"), val = int32(1)];
        tensor<fp16, [1, 1, 1, 1]> W = const()[name = string("W"), val = tensor<fp16, [1, 1, 1, 1]>(BLOBFILE(path = string("@model_path/weights/weight.bin"), offset = uint64(64)))];
        tensor<fp16, [1, 1, 1, %d]> y = conv(dilations = c_dilations, groups = c_groups, pad = c_pad, pad_type = c_pad_type, strides = c_strides, weight = W, x = x)[name = string("conv")];
    } -> (y);
}
`, buildInfo, spatial, spatial)
}

// BuildWeightBlob constructs the binary weight blob for MIL compilation.
//
// The blob layout matches the ANE's expected format:
//   - Bytes 0-63:   File header (0x01 at offset 0, 0x02 at offset 4)
//   - Bytes 64-127: Chunk header (0xDEADBEEF magic, data size, data offset)
//   - Bytes 128+:   FP16 weight data
//
// weights must have exactly outCh*inCh elements (OIHW layout, H=W=1).
func BuildWeightBlob(weights []float32, outCh, inCh int) ([]byte, error) {
	expected := outCh * inCh
	if len(weights) != expected {
		return nil, fmt.Errorf("mil: weight count %d != outCh*inCh (%d*%d = %d)", len(weights), outCh, inCh, expected)
	}

	fp16Data := make([]byte, len(weights)*2)
	for i, w := range weights {
		binary.LittleEndian.PutUint16(fp16Data[i*2:], float32ToFP16(w))
	}

	const fileHeaderSize = 64
	const chunkHeaderSize = 64
	dataOffset := fileHeaderSize + chunkHeaderSize
	totalSize := dataOffset + len(fp16Data)

	buf := make([]byte, totalSize)

	// File header (64 bytes).
	buf[0] = 0x01 // file magic byte 1
	buf[4] = 0x02 // file magic byte 2

	// Chunk header at offset 64.
	off := fileHeaderSize
	binary.LittleEndian.PutUint32(buf[off:], 0xDEADBEEF)              // chunk magic
	buf[off+4] = 0x01                                                 // chunk version
	binary.LittleEndian.PutUint32(buf[off+8:], uint32(len(fp16Data))) // data size
	binary.LittleEndian.PutUint32(buf[off+16:], uint32(dataOffset))   // absolute data offset

	// FP16 weight data at offset 128.
	copy(buf[dataOffset:], fp16Data)

	return buf, nil
}

// BuildIdentityWeightBlob builds weights for an identity convolution (I matrix).
func BuildIdentityWeightBlob(channels int) ([]byte, error) {
	weights := make([]float32, channels*channels)
	for i := range channels {
		weights[i*channels+i] = 1.0
	}
	return BuildWeightBlob(weights, channels, channels)
}

func float32ToFP16(f float32) uint16 {
	b := math.Float32bits(f)
	sign := (b >> 16) & 0x8000
	exp := int((b>>23)&0xFF) - 127 + 15
	frac := b & 0x7FFFFF

	switch {
	case exp <= 0:
		return uint16(sign)
	case exp >= 31:
		return uint16(sign | 0x7C00)
	default:
		return uint16(sign | uint32(exp)<<10 | (frac >> 13))
	}
}

// GenRMSNorm generates a MIL text for the overflow-safe 11-op RMSNorm decomposition.
//
// The 11-op sequence prevents fp16 overflow (values >256 cause CPU fallback):
//
//	abs → reduce_max → maximum(1e-6) → real_div → square → reduce_mean →
//	add(eps) → sqrt → mul(safe_max) → real_div → mul(weight)
//
// The program takes a single fp16 tensor input [1, channels, 1, spatial] and produces
// the same shape output. The weight vector is loaded from a BLOBFILE.
func GenRMSNorm(channels, spatial int, eps float64) string {
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {
        tensor<fp16, [1, %d, 1, %d]> abs_x = abs(x = x)[name = string("abs")];
        tensor<fp16, [1, 1, 1, %d]> max_abs = reduce_max(axes = tensor<int32, [1]>([1]), keep_dims = true, x = abs_x)[name = string("reduce_max")];
        fp16 c_eps_floor = const()[name = string("c_eps_floor"), val = fp16(1e-6)];
        tensor<fp16, [1, 1, 1, %d]> safe_max = maximum(x = max_abs, y = c_eps_floor)[name = string("safe_max")];
        tensor<fp16, [1, %d, 1, %d]> x_normed = real_div(x = x, y = safe_max)[name = string("norm_div")];
        tensor<fp16, [1, %d, 1, %d]> x_sq = square(x = x_normed)[name = string("square")];
        tensor<fp16, [1, 1, 1, %d]> mean_sq = reduce_mean(axes = tensor<int32, [1]>([1]), keep_dims = true, x = x_sq)[name = string("reduce_mean")];
        fp16 c_eps = const()[name = string("c_eps"), val = fp16(%e)];
        tensor<fp16, [1, 1, 1, %d]> mean_sq_eps = add(x = mean_sq, y = c_eps)[name = string("add_eps")];
        tensor<fp16, [1, 1, 1, %d]> rms_normed = sqrt(x = mean_sq_eps)[name = string("sqrt")];
        tensor<fp16, [1, 1, 1, %d]> rms = mul(x = rms_normed, y = safe_max)[name = string("mul_rescale")];
        tensor<fp16, [1, %d, 1, %d]> x_div_rms = real_div(x = x, y = rms)[name = string("rms_div")];
        tensor<fp16, [%d]> W = const()[name = string("W"), val = tensor<fp16, [%d]>(BLOBFILE(path = string("@model_path/weights/weight.bin"), offset = uint64(64)))];
        tensor<fp16, [1, %d, 1, %d]> y = mul(x = x_div_rms, y = W)[name = string("mul_weight")];
    } -> (y);
}
`, buildInfo,
		channels, spatial, // input shape
		channels, spatial, // abs shape
		spatial,           // reduce_max output
		spatial,           // safe_max
		channels, spatial, // x_normed
		channels, spatial, // x_sq
		spatial,           // mean_sq
		eps,               // eps value
		spatial,           // mean_sq_eps
		spatial,           // rms_normed (sqrt output)
		spatial,           // rms (rescaled by safe_max)
		channels, spatial, // x_div_rms
		channels,          // weight shape
		channels,          // weight shape (BLOBFILE)
		channels, spatial, // y
	)
}

// BuildWeightBlobV1 constructs a binary weight blob for a flat 1D weight vector.
// Unlike BuildWeightBlob which reshapes to OIHW, this stores raw 1D data
// suitable for RMSNorm weight vectors and other non-convolution weights.
func BuildWeightBlobV1(data []float32) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("mil: empty weight data")
	}

	fp16Data := make([]byte, len(data)*2)
	for i, w := range data {
		binary.LittleEndian.PutUint16(fp16Data[i*2:], float32ToFP16(w))
	}

	const fileHeaderSize = 64
	const chunkHeaderSize = 64
	dataOffset := fileHeaderSize + chunkHeaderSize
	totalSize := dataOffset + len(fp16Data)

	buf := make([]byte, totalSize)

	// File header (64 bytes).
	buf[0] = 0x01
	buf[4] = 0x02

	// Chunk header at offset 64.
	off := fileHeaderSize
	binary.LittleEndian.PutUint32(buf[off:], 0xDEADBEEF)
	buf[off+4] = 0x01
	binary.LittleEndian.PutUint32(buf[off+8:], uint32(len(fp16Data)))
	binary.LittleEndian.PutUint32(buf[off+16:], uint32(dataOffset))

	// FP16 weight data at offset 128.
	copy(buf[dataOffset:], fp16Data)

	return buf, nil
}

// GenSDPA generates a MIL text for scaled dot-product attention.
// Inputs: Q, K, V of shape [1, nHeads, seqLen, headDim]. Scale is 1/sqrt(headDim).
func GenSDPA(headDim, nHeads, seqLen int) string {
	scale := 1.0 / math.Sqrt(float64(headDim))
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>(tensor<fp16, [1, %d, %d, %d]> Q, tensor<fp16, [1, %d, %d, %d]> K, tensor<fp16, [1, %d, %d, %d]> V) {
        fp16 c_scale = const()[name = string("c_scale"), val = fp16(%e)];
        tensor<fp16, [1, %d, %d, %d]> y = scaled_dot_product_attention(query = Q, key = K, value = V, scale = c_scale)[name = string("sdpa")];
    } -> (y);
}
`, buildInfo,
		nHeads, seqLen, headDim, // Q shape
		nHeads, seqLen, headDim, // K shape
		nHeads, seqLen, headDim, // V shape
		scale,                   // scale value
		nHeads, seqLen, headDim, // output shape
	)
}

// GenReadState generates a MIL text for reading a named state buffer.
// This is used for iOS 18+ stateful inference (e.g., KV cache on ANE).
func GenReadState(name string, shape [4]int) string {
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>() {
        tensor<fp16, [%d, %d, %d, %d]> y = read_state(name = string("%s"))[name = string("read_%s")];
    } -> (y);
}
`, buildInfo,
		shape[0], shape[1], shape[2], shape[3],
		name, name,
	)
}

// GenUpdateState generates a MIL text for updating a named state buffer.
// This emits the coreml_update_state op for iOS 18+ stateful inference.
func GenUpdateState(name string, shape [4]int) string {
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>(tensor<fp16, [%d, %d, %d, %d]> value) {
        tensor<fp16, [%d, %d, %d, %d]> y = coreml_update_state(state = string("%s"), value = value)[name = string("update_%s")];
    } -> (y);
}
`, buildInfo,
		shape[0], shape[1], shape[2], shape[3], // input shape
		shape[0], shape[1], shape[2], shape[3], // output shape
		name, name,
	)
}

// GenGQAExpand generates a MIL text for expanding KV heads for grouped-query attention.
// It tiles the KV tensor along the head dimension by a factor of qHeads/kvHeads.
func GenGQAExpand(kvHeads, qHeads, headDim, seqLen int) string {
	if kvHeads == 0 {
		kvHeads = 1
	}
	repeatFactor := qHeads / kvHeads
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>(tensor<fp16, [1, %d, %d, %d]> x) {
        tensor<int32, [4]> c_reps = const()[name = string("c_reps"), val = tensor<int32, [4]>([1, %d, 1, 1])];
        tensor<fp16, [1, %d, %d, %d]> y = tile(reps = c_reps, x = x)[name = string("tile")];
    } -> (y);
}
`, buildInfo,
		kvHeads, seqLen, headDim, // input shape [1, kvHeads, seqLen, headDim]
		repeatFactor,            // tile repeat factor on head dim
		qHeads, seqLen, headDim, // output shape [1, qHeads, seqLen, headDim]
	)
}

func fp16ToFloat32(h uint16) float32 {
	sign := uint32(h>>15) & 1
	exp := uint32(h>>10) & 0x1F
	frac := uint32(h) & 0x3FF

	switch {
	case exp == 0:
		if frac == 0 {
			return math.Float32frombits(sign << 31)
		}
		for frac&0x400 == 0 {
			frac <<= 1
			exp--
		}
		exp++
		frac &= 0x3FF
		fallthrough
	case exp < 31:
		return math.Float32frombits(sign<<31 | (exp+127-15)<<23 | frac<<13)
	default:
		return math.Float32frombits(sign<<31 | 0x7F800000 | frac<<13)
	}
}
