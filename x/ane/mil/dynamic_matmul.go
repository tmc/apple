//go:build darwin

package mil

import "fmt"

const dynamicMatmulBuildInfo = `[buildInfo = dict<string, string>({{"coremlc-component-MIL", "3510.2.1"}, {"coremlc-version", "3505.4.1"}, {"coremltools-component-milinternal", ""}, {"coremltools-version", "9.0"}})]`

// GenDynamicMatmul generates a weightless MIL graph for y = x*w.
//
// The single input tensor packs activations and weights together as
// [1, inCh, 1, batch+outCh] fp32:
//   - spatial [0:batch] contains activations laid out as [inCh, batch]
//   - spatial [batch:batch+outCh] contains weights laid out as [inCh, outCh]
//
// The output tensor is [1, outCh, 1, batch] fp32.
func GenDynamicMatmul(inCh, outCh, batch int) string {
	spatial := batch + outCh
	return fmt.Sprintf(`program(1.3)
%s
{
    func main<ios18>(tensor<fp32, [1, %d, 1, %d]> x) {
        string to16 = const()[name = string("to16"), val = string("fp16")];
        tensor<fp16, [1, %d, 1, %d]> xh = cast(dtype = to16, x = x)[name = string("cin")];

        tensor<int32, [4]> act_begin = const()[name = string("act_begin"), val = tensor<int32, [4]>([0, 0, 0, 0])];
        tensor<int32, [4]> act_size = const()[name = string("act_size"), val = tensor<int32, [4]>([1, %d, 1, %d])];
        tensor<fp16, [1, %d, 1, %d]> act = slice_by_size(x = xh, begin = act_begin, size = act_size)[name = string("act")];

        tensor<int32, [4]> wt_begin = const()[name = string("wt_begin"), val = tensor<int32, [4]>([0, 0, 0, %d])];
        tensor<int32, [4]> wt_size = const()[name = string("wt_size"), val = tensor<int32, [4]>([1, %d, 1, %d])];
        tensor<fp16, [1, %d, 1, %d]> wt = slice_by_size(x = xh, begin = wt_begin, size = wt_size)[name = string("wt")];

        tensor<int32, [4]> act_shape = const()[name = string("act_shape"), val = tensor<int32, [4]>([1, 1, %d, %d])];
        tensor<fp16, [1, 1, %d, %d]> act2 = reshape(shape = act_shape, x = act)[name = string("act2")];

        tensor<int32, [4]> perm = const()[name = string("perm"), val = tensor<int32, [4]>([0, 1, 3, 2])];
        tensor<fp16, [1, 1, %d, %d]> act3 = transpose(perm = perm, x = act2)[name = string("act3")];

        tensor<int32, [4]> wt_shape = const()[name = string("wt_shape"), val = tensor<int32, [4]>([1, 1, %d, %d])];
        tensor<fp16, [1, 1, %d, %d]> W = reshape(shape = wt_shape, x = wt)[name = string("W")];

        bool bfalse = const()[name = string("bfalse"), val = bool(false)];
        tensor<fp16, [1, 1, %d, %d]> yh = matmul(transpose_x = bfalse, transpose_y = bfalse, x = act3, y = W)[name = string("mm")];

        tensor<fp16, [1, 1, %d, %d]> yt = transpose(perm = perm, x = yh)[name = string("yt")];

        tensor<int32, [4]> out_shape = const()[name = string("out_shape"), val = tensor<int32, [4]>([1, %d, 1, %d])];
        tensor<fp16, [1, %d, 1, %d]> yr = reshape(shape = out_shape, x = yt)[name = string("yr")];

        string to32 = const()[name = string("to32"), val = string("fp32")];
        tensor<fp32, [1, %d, 1, %d]> y = cast(dtype = to32, x = yr)[name = string("cout")];
    } -> (y);
}
`, dynamicMatmulBuildInfo,
		inCh, spatial,
		inCh, spatial,
		inCh, batch,
		inCh, batch,
		batch,
		inCh, outCh,
		inCh, outCh,
		inCh, batch,
		inCh, batch,
		batch, inCh,
		inCh, outCh,
		inCh, outCh,
		batch, outCh,
		outCh, batch,
		outCh, batch,
		outCh, batch,
		outCh, batch,
	)
}
