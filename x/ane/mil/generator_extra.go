//go:build darwin

package mil

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

// buildInfoHeader is the MIL program header expected by the ANE compiler.
// Version strings match coremltools 9.0 output format.
const buildInfoHeader = "program(1.3)\n" +
	"[buildInfo = dict<string, string>({{\"coremlc-component-MIL\", \"3510.2.1\"}, {\"coremlc-version\", \"3505.4.1\"}, {\"coremltools-component-milinternal\", \"\"}, {\"coremltools-version\", \"9.0\"}})]\n"

// GenConvFP16 generates a minimal fp16 conv graph matching the working training path.
func GenConvFP16(inCh, outCh, spatial int) string {
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n"+
			"        string c_pad_type = const()[name = string(\"c_pad_type\"), val = string(\"valid\")];\n"+
			"        tensor<int32, [2]> c_strides = const()[name = string(\"c_strides\"), val = tensor<int32, [2]>([1, 1])];\n"+
			"        tensor<int32, [4]> c_pad = const()[name = string(\"c_pad\"), val = tensor<int32, [4]>([0, 0, 0, 0])];\n"+
			"        tensor<int32, [2]> c_dilations = const()[name = string(\"c_dilations\"), val = tensor<int32, [2]>([1, 1])];\n"+
			"        int32 c_groups = const()[name = string(\"c_groups\"), val = int32(1)];\n"+
			"        tensor<fp16, [%d, %d, 1, 1]> W = const()[name = string(\"W\"), val = tensor<fp16, [%d, %d, 1, 1]>(BLOBFILE(path = string(\"@model_path/weights/weight.bin\"), offset = uint64(64)))];\n"+
			"        tensor<fp16, [1, %d, 1, %d]> y = conv(dilations = c_dilations, groups = c_groups, pad = c_pad, pad_type = c_pad_type, strides = c_strides, weight = W, x = x)[name = string(\"conv\")];\n"+
			"    } -> (y);\n"+
			"}\n",
		inCh, spatial,
		outCh, inCh, outCh, inCh,
		outCh, spatial,
	)
}

// GenConvDynamicFP16 generates an fp16 conv graph with runtime-provided weights.
func GenConvDynamicFP16(inCh, outCh, spatial int) string {
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x, tensor<fp16, [%d, %d, 1, 1]> W) {\n"+
			"        string c_pad_type = const()[name = string(\"c_pad_type\"), val = string(\"valid\")];\n"+
			"        tensor<int32, [2]> c_strides = const()[name = string(\"c_strides\"), val = tensor<int32, [2]>([1, 1])];\n"+
			"        tensor<int32, [4]> c_pad = const()[name = string(\"c_pad\"), val = tensor<int32, [4]>([0, 0, 0, 0])];\n"+
			"        tensor<int32, [2]> c_dilations = const()[name = string(\"c_dilations\"), val = tensor<int32, [2]>([1, 1])];\n"+
			"        int32 c_groups = const()[name = string(\"c_groups\"), val = int32(1)];\n"+
			"        tensor<fp16, [1, %d, 1, %d]> y = conv(dilations = c_dilations, groups = c_groups, pad = c_pad, pad_type = c_pad_type, strides = c_strides, weight = W, x = x)[name = string(\"conv\")];\n"+
			"    } -> (y);\n"+
			"}\n",
		inCh, spatial, outCh, inCh, outCh, spatial,
	)
}

// GenClassifierForward generates a classifier projection kernel.
// The embedding matrix is baked as a conv weight tensor with shape [vocab, dim, 1, 1].
func GenClassifierForward(dim, vocab, seq int) string {
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n"+
			"        string pt = const()[name=string(\"pt\"), val=string(\"valid\")];\n"+
			"        tensor<int32, [2]> st = const()[name=string(\"st\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        tensor<int32, [4]> pd = const()[name=string(\"pd\"), val=tensor<int32, [4]>([0,0,0,0])];\n"+
			"        tensor<int32, [2]> dl = const()[name=string(\"dl\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        int32 gr = const()[name=string(\"gr\"), val=int32(1)];\n"+
			"        tensor<fp16, [%d,%d,1,1]> We = const()[name=string(\"We\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/embed.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=We,x=x)[name=string(\"cls\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		dim, seq,
		vocab, dim, vocab, dim,
		vocab, seq,
	)
}

// GenClassifierBackward generates the classifier backward kernel.
// It multiplies baked transpose(embed) by dlogits to produce dx.
func GenClassifierBackward(dim, vocab, seq int) string {
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> dl) {\n"+
			"        tensor<int32, [3]> sh3 = const()[name=string(\"sh3\"), val=tensor<int32, [3]>([1,%d,%d])];\n"+
			"        tensor<fp16, [1,%d,%d]> dl3 = reshape(shape=sh3,x=dl)[name=string(\"rdl\")];\n"+
			"        tensor<fp16, [1,%d,%d]> Wet = const()[name=string(\"Wet\"), val=tensor<fp16, [1,%d,%d]>(BLOBFILE(path=string(\"@model_path/weights/embed_t.bin\"), offset=uint64(64)))];\n"+
			"        bool bF = const()[name=string(\"bF\"), val=bool(false)];\n"+
			"        tensor<fp16, [1,%d,%d]> dx3 = matmul(transpose_x=bF,transpose_y=bF,x=Wet,y=dl3)[name=string(\"mm\")];\n"+
			"        tensor<int32, [4]> sh4 = const()[name=string(\"sh4\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = reshape(shape=sh4,x=dx3)[name=string(\"out\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		vocab, seq,
		vocab, seq,
		vocab, seq,
		dim, vocab, dim, vocab,
		dim, seq,
		dim, seq,
		dim, seq,
	)
}

// GenSoftmaxVocab generates a softmax over the channel dimension.
func GenSoftmaxVocab(vocab, seq int) string {
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n"+
			"        int32 ax = const()[name=string(\"ax\"), val=int32(1)];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = softmax(axis=ax,x=x)[name=string(\"sm\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		vocab, seq,
		vocab, seq,
	)
}

// GenFinalRMSNorm generates a final-layer RMSNorm kernel with baked weights.
func GenFinalRMSNorm(dim, seq int) string {
	invd := 1.0 / float64(dim)
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n"+
			"        tensor<fp16, [1,%d,1,%d]> sq = mul(x=x,y=x)[name=string(\"sq\")];\n"+
			"        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([1])];\n"+
			"        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss = reduce_sum(x=sq,axes=rax,keep_dims=kd)[name=string(\"ss\")];\n"+
			"        fp16 invd = const()[name=string(\"invd\"), val=fp16(%f)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss2 = mul(x=ss,y=invd)[name=string(\"ss2\")];\n"+
			"        fp16 eps = const()[name=string(\"eps\"), val=fp16(0.00001)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss3 = add(x=ss2,y=eps)[name=string(\"ss3\")];\n"+
			"        fp16 nhalf = const()[name=string(\"nhalf\"), val=fp16(-0.5)];\n"+
			"        tensor<fp16, [1,1,1,%d]> rrms = pow(x=ss3,y=nhalf)[name=string(\"rrms\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xr = mul(x=x,y=rrms)[name=string(\"xr\")];\n"+
			"        tensor<fp16, [1,%d,1,1]> rw = const()[name=string(\"rw\"), val=tensor<fp16, [1,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/rms_w.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = mul(x=xr,y=rw)[name=string(\"out\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		dim, seq,
		dim, seq,
		seq,
		invd,
		seq,
		seq,
		seq,
		dim, seq,
		dim, dim,
		dim, seq,
	)
}

// GenFinalRMSNormDynamic generates a final-layer RMSNorm kernel with runtime-provided weights.
func GenFinalRMSNormDynamic(dim, seq int) string {
	invd := 1.0 / float64(dim)
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x, tensor<fp16, [1, %d, 1, 1]> rw) {\n"+
			"        tensor<fp16, [1,%d,1,%d]> sq = mul(x=x,y=x)[name=string(\"sq\")];\n"+
			"        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([1])];\n"+
			"        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss = reduce_sum(x=sq,axes=rax,keep_dims=kd)[name=string(\"ss\")];\n"+
			"        fp16 invd = const()[name=string(\"invd\"), val=fp16(%f)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss2 = mul(x=ss,y=invd)[name=string(\"ss2\")];\n"+
			"        fp16 eps = const()[name=string(\"eps\"), val=fp16(0.00001)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss3 = add(x=ss2,y=eps)[name=string(\"ss3\")];\n"+
			"        fp16 nhalf = const()[name=string(\"nhalf\"), val=fp16(-0.5)];\n"+
			"        tensor<fp16, [1,1,1,%d]> rrms = pow(x=ss3,y=nhalf)[name=string(\"rrms\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xr = mul(x=x,y=rrms)[name=string(\"xr\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = mul(x=xr,y=rw)[name=string(\"out\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		dim, seq, dim,
		dim, seq,
		seq,
		invd,
		seq,
		seq,
		seq,
		dim, seq,
		dim, seq,
	)
}

// GenRMSNormBackward generates the dx half of RMSNorm backward with baked weights.
// The input is concat(dy, x) along the channel dimension; dw remains a cheap CPU reduction.
func GenRMSNormBackward(dim, seq int) string {
	invd := 1.0 / float64(dim)
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> inp) {\n"+
			"        tensor<int32, [4]> sz = const()[name=string(\"sz\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n"+
			"        tensor<int32, [4]> b0 = const()[name=string(\"b0\"), val=tensor<int32, [4]>([0,0,0,0])];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dy = slice_by_size(x=inp,begin=b0,size=sz)[name=string(\"sdy\")];\n"+
			"        tensor<int32, [4]> b1 = const()[name=string(\"b1\"), val=tensor<int32, [4]>([0,%d,0,0])];\n"+
			"        tensor<fp16, [1,%d,1,%d]> x = slice_by_size(x=inp,begin=b1,size=sz)[name=string(\"sx\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> sq = mul(x=x,y=x)[name=string(\"sq\")];\n"+
			"        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([1])];\n"+
			"        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss = reduce_sum(x=sq,axes=rax,keep_dims=kd)[name=string(\"ss\")];\n"+
			"        fp16 invd = const()[name=string(\"invd\"), val=fp16(%f)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss2 = mul(x=ss,y=invd)[name=string(\"ss2\")];\n"+
			"        fp16 eps = const()[name=string(\"eps\"), val=fp16(0.00001)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss3 = add(x=ss2,y=eps)[name=string(\"ss3\")];\n"+
			"        fp16 nhalf = const()[name=string(\"nhalf\"), val=fp16(-0.5)];\n"+
			"        tensor<fp16, [1,1,1,%d]> rrms = pow(x=ss3,y=nhalf)[name=string(\"rrms\")];\n"+
			"        tensor<fp16, [1,%d,1,1]> w = const()[name=string(\"w\"), val=tensor<fp16, [1,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/rms_w.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dyw = mul(x=dy,y=w)[name=string(\"dyw\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dywx = mul(x=dyw,y=x)[name=string(\"dywx\")];\n"+
			"        tensor<fp16, [1,1,1,%d]> dot_sum = reduce_sum(x=dywx,axes=rax,keep_dims=kd)[name=string(\"ds\")];\n"+
			"        tensor<fp16, [1,1,1,%d]> dot_sc = mul(x=dot_sum,y=invd)[name=string(\"dsc\")];\n"+
			"        tensor<fp16, [1,1,1,%d]> rrms2 = mul(x=rrms,y=rrms)[name=string(\"rr2\")];\n"+
			"        tensor<fp16, [1,1,1,%d]> coeff = mul(x=dot_sc,y=rrms2)[name=string(\"cof\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xc = mul(x=x,y=coeff)[name=string(\"xc\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> diff = sub(x=dyw,y=xc)[name=string(\"dif\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = mul(x=diff,y=rrms)[name=string(\"out\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		2*dim, seq,
		dim, seq,
		dim, seq,
		dim,
		dim, seq,
		dim, seq,
		seq,
		invd,
		seq,
		seq,
		seq,
		dim, dim,
		dim, seq,
		dim, seq,
		seq,
		seq,
		seq,
		seq,
		dim, seq,
		dim, seq,
		dim, seq,
	)
}

// GenRMSNormBackwardDynamic generates the dx half of RMSNorm backward with runtime-provided weights.
// The input is concat(dy, x) along the channel dimension; dw remains a cheap CPU reduction.
func GenRMSNormBackwardDynamic(dim, seq int) string {
	invd := 1.0 / float64(dim)
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> inp, tensor<fp16, [1, %d, 1, 1]> w) {\n"+
			"        tensor<int32, [4]> sz = const()[name=string(\"sz\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n"+
			"        tensor<int32, [4]> b0 = const()[name=string(\"b0\"), val=tensor<int32, [4]>([0,0,0,0])];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dy = slice_by_size(x=inp,begin=b0,size=sz)[name=string(\"sdy\")];\n"+
			"        tensor<int32, [4]> b1 = const()[name=string(\"b1\"), val=tensor<int32, [4]>([0,%d,0,0])];\n"+
			"        tensor<fp16, [1,%d,1,%d]> x = slice_by_size(x=inp,begin=b1,size=sz)[name=string(\"sx\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> sq = mul(x=x,y=x)[name=string(\"sq\")];\n"+
			"        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([1])];\n"+
			"        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss = reduce_sum(x=sq,axes=rax,keep_dims=kd)[name=string(\"ss\")];\n"+
			"        fp16 invd = const()[name=string(\"invd\"), val=fp16(%f)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss2 = mul(x=ss,y=invd)[name=string(\"ss2\")];\n"+
			"        fp16 eps = const()[name=string(\"eps\"), val=fp16(0.00001)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss3 = add(x=ss2,y=eps)[name=string(\"ss3\")];\n"+
			"        fp16 nhalf = const()[name=string(\"nhalf\"), val=fp16(-0.5)];\n"+
			"        tensor<fp16, [1,1,1,%d]> rrms = pow(x=ss3,y=nhalf)[name=string(\"rrms\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dyw = mul(x=dy,y=w)[name=string(\"dyw\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dywx = mul(x=dyw,y=x)[name=string(\"dywx\")];\n"+
			"        tensor<fp16, [1,1,1,%d]> dot_sum = reduce_sum(x=dywx,axes=rax,keep_dims=kd)[name=string(\"ds\")];\n"+
			"        tensor<fp16, [1,1,1,%d]> dot_sc = mul(x=dot_sum,y=invd)[name=string(\"dsc\")];\n"+
			"        tensor<fp16, [1,1,1,%d]> rrms2 = mul(x=rrms,y=rrms)[name=string(\"rr2\")];\n"+
			"        tensor<fp16, [1,1,1,%d]> coeff = mul(x=dot_sc,y=rrms2)[name=string(\"cof\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xc = mul(x=x,y=coeff)[name=string(\"xc\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> diff = sub(x=dyw,y=xc)[name=string(\"dif\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = mul(x=diff,y=rrms)[name=string(\"out\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		2*dim, seq, dim,
		dim, seq,
		dim, seq,
		dim,
		dim, seq,
		dim, seq,
		seq,
		invd,
		seq,
		seq,
		seq,
		dim, seq,
		dim, seq,
		seq,
		seq,
		seq,
		seq,
		dim, seq,
		dim, seq,
		dim, seq,
	)
}

// GenFFNForward generates a fused FFN block with baked W1/W2/W3 weights.
// It computes W2(silu(W1(x)) * W3(x)).
func GenFFNForward(dim, hidden, seq int) string {
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n"+
			"        string pt = const()[name=string(\"pt\"), val=string(\"valid\")];\n"+
			"        tensor<int32, [2]> st = const()[name=string(\"st\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        tensor<int32, [4]> pd = const()[name=string(\"pd\"), val=tensor<int32, [4]>([0,0,0,0])];\n"+
			"        tensor<int32, [2]> dl = const()[name=string(\"dl\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        int32 gr = const()[name=string(\"gr\"), val=int32(1)];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W1 = const()[name=string(\"W1\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w1.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W3 = const()[name=string(\"W3\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w3.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W2 = const()[name=string(\"W2\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w2.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> h1 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W1,x=x)[name=string(\"c1\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> h3 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W3,x=x)[name=string(\"c3\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> sig = sigmoid(x=h1)[name=string(\"sg\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> silu = mul(x=h1,y=sig)[name=string(\"si\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> gate = mul(x=silu,y=h3)[name=string(\"gt\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W2,x=gate)[name=string(\"c2\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		dim, seq,
		hidden, dim, hidden, dim,
		hidden, dim, hidden, dim,
		dim, hidden, dim, hidden,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		dim, seq,
	)
}

// GenFFNForwardRMS generates the full FFN block with internal RMSNorm and the
// final residual-free output only.
func GenFFNForwardRMS(dim, hidden, seq int) string {
	invd := 1.0 / float64(dim)
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n"+
			"        tensor<fp16, [1,%d,1,%d]> sq = mul(x=x,y=x)[name=string(\"sq\")];\n"+
			"        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([1])];\n"+
			"        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss = reduce_sum(x=sq,axes=rax,keep_dims=kd)[name=string(\"ss\")];\n"+
			"        fp16 invd = const()[name=string(\"invd\"), val=fp16(%f)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss2 = mul(x=ss,y=invd)[name=string(\"ss2\")];\n"+
			"        fp16 eps = const()[name=string(\"eps\"), val=fp16(0.00001)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss3 = add(x=ss2,y=eps)[name=string(\"ss3\")];\n"+
			"        fp16 nhalf = const()[name=string(\"nhalf\"), val=fp16(-0.5)];\n"+
			"        tensor<fp16, [1,1,1,%d]> rrms = pow(x=ss3,y=nhalf)[name=string(\"rrms\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xr = mul(x=x,y=rrms)[name=string(\"xr\")];\n"+
			"        tensor<fp16, [1,%d,1,1]> rw = const()[name=string(\"rw\"), val=tensor<fp16, [1,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/rms2.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xn = mul(x=xr,y=rw)[name=string(\"xn\")];\n"+
			"        string pt = const()[name=string(\"pt\"), val=string(\"valid\")];\n"+
			"        tensor<int32, [2]> st = const()[name=string(\"st\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        tensor<int32, [4]> pd = const()[name=string(\"pd\"), val=tensor<int32, [4]>([0,0,0,0])];\n"+
			"        tensor<int32, [2]> dl = const()[name=string(\"dl\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        int32 gr = const()[name=string(\"gr\"), val=int32(1)];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W1 = const()[name=string(\"W1\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w1.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W3 = const()[name=string(\"W3\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w3.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W2 = const()[name=string(\"W2\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w2.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> h1 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W1,x=xn)[name=string(\"c1\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> h3 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W3,x=xn)[name=string(\"c3\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> sig = sigmoid(x=h1)[name=string(\"sg\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> silu = mul(x=h1,y=sig)[name=string(\"si\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> gate = mul(x=silu,y=h3)[name=string(\"gt\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W2,x=gate)[name=string(\"c2\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		dim, seq,
		dim, seq,
		seq,
		invd,
		seq,
		seq,
		seq,
		dim, seq,
		dim, dim,
		dim, seq,
		hidden, dim, hidden, dim,
		hidden, dim, hidden, dim,
		dim, hidden, dim, hidden,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		dim, seq,
	)
}

// GenFFNForwardTaps generates a fused FFN block that also returns intermediates.
// The output layout is concat(out, h1, h3) along the channel dimension.
func GenFFNForwardTaps(dim, hidden, seq int) string {
	invd := 1.0 / float64(dim)
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n"+
			"        tensor<fp16, [1,%d,1,%d]> sq = mul(x=x,y=x)[name=string(\"sq\")];\n"+
			"        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([1])];\n"+
			"        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss = reduce_sum(x=sq,axes=rax,keep_dims=kd)[name=string(\"ss\")];\n"+
			"        fp16 invd = const()[name=string(\"invd\"), val=fp16(%f)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss2 = mul(x=ss,y=invd)[name=string(\"ss2\")];\n"+
			"        fp16 eps = const()[name=string(\"eps\"), val=fp16(0.00001)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss3 = add(x=ss2,y=eps)[name=string(\"ss3\")];\n"+
			"        fp16 nhalf = const()[name=string(\"nhalf\"), val=fp16(-0.5)];\n"+
			"        tensor<fp16, [1,1,1,%d]> rrms = pow(x=ss3,y=nhalf)[name=string(\"rrms\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xr = mul(x=x,y=rrms)[name=string(\"xr\")];\n"+
			"        tensor<fp16, [1,%d,1,1]> rw = const()[name=string(\"rw\"), val=tensor<fp16, [1,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/rms2.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xn = mul(x=xr,y=rw)[name=string(\"xn\")];\n"+
			"        string pt = const()[name=string(\"pt\"), val=string(\"valid\")];\n"+
			"        tensor<int32, [2]> st = const()[name=string(\"st\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        tensor<int32, [4]> pd = const()[name=string(\"pd\"), val=tensor<int32, [4]>([0,0,0,0])];\n"+
			"        tensor<int32, [2]> dl = const()[name=string(\"dl\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        int32 gr = const()[name=string(\"gr\"), val=int32(1)];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W1 = const()[name=string(\"W1\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w1.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W3 = const()[name=string(\"W3\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w3.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W2 = const()[name=string(\"W2\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w2.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> h1 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W1,x=xn)[name=string(\"c1\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> h3 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W3,x=xn)[name=string(\"c3\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> sig = sigmoid(x=h1)[name=string(\"sg\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> silu = mul(x=h1,y=sig)[name=string(\"si\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> gate = mul(x=silu,y=h3)[name=string(\"gt\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W2,x=gate)[name=string(\"c2\")];\n"+
			"        int32 cax = const()[name=string(\"cax\"), val=int32(1)];\n"+
			"        bool cid = const()[name=string(\"cid\"), val=bool(false)];\n"+
			"        tensor<fp16, [1,%d,1,%d]> taps = concat(axis=cax,interleave=cid,values=(out,h1,h3))[name=string(\"cat\")];\n"+
			"    } -> (taps);\n"+
			"}\n",
		dim, seq,
		dim, seq,
		seq,
		invd,
		seq,
		seq,
		seq,
		dim, seq,
		dim, dim,
		dim, seq,
		hidden, dim, hidden, dim,
		hidden, dim, hidden, dim,
		dim, hidden, dim, hidden,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		dim, seq,
		dim+2*hidden, seq,
	)
}

// GenFFNBackward generates the backward half of the fused FFN block.
// Input layout is concat(dffn, h1, h3); output layout is concat(dx, dh1, dh3).
func GenFFNBackward(dim, hidden, seq int) string {
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n"+
			"        string pt = const()[name=string(\"pt\"), val=string(\"valid\")];\n"+
			"        tensor<int32, [2]> st = const()[name=string(\"st\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        tensor<int32, [4]> pd = const()[name=string(\"pd\"), val=tensor<int32, [4]>([0,0,0,0])];\n"+
			"        tensor<int32, [2]> dl = const()[name=string(\"dl\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        int32 gr = const()[name=string(\"gr\"), val=int32(1)];\n"+
			"        tensor<int32, [4]> bd = const()[name=string(\"bd\"), val=tensor<int32, [4]>([0,0,0,0])];\n"+
			"        tensor<int32, [4]> sd = const()[name=string(\"sd\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dffn = slice_by_size(x=x,begin=bd,size=sd)[name=string(\"s0\")];\n"+
			"        tensor<int32, [4]> b1 = const()[name=string(\"b1\"), val=tensor<int32, [4]>([0,%d,0,0])];\n"+
			"        tensor<int32, [4]> s1 = const()[name=string(\"s1\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n"+
			"        tensor<fp16, [1,%d,1,%d]> h1 = slice_by_size(x=x,begin=b1,size=s1)[name=string(\"s1x\")];\n"+
			"        tensor<int32, [4]> b3 = const()[name=string(\"b3\"), val=tensor<int32, [4]>([0,%d,0,0])];\n"+
			"        tensor<fp16, [1,%d,1,%d]> h3 = slice_by_size(x=x,begin=b3,size=s1)[name=string(\"s3x\")];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W2t = const()[name=string(\"W2t\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w2t.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dsilu = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W2t,x=dffn)[name=string(\"cw2\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> sig = sigmoid(x=h1)[name=string(\"sg\")];\n"+
			"        fp16 one = const()[name=string(\"one\"), val=fp16(1.0)];\n"+
			"        tensor<fp16, [1,%d,1,%d]> oms = sub(x=one,y=sig)[name=string(\"oms\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> homs = mul(x=h1,y=oms)[name=string(\"homs\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> brk = add(x=one,y=homs)[name=string(\"brk\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dsd = mul(x=sig,y=brk)[name=string(\"dsd\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> t1 = mul(x=dsilu,y=h3)[name=string(\"t1\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dh1 = mul(x=t1,y=dsd)[name=string(\"dh1\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> slh = mul(x=h1,y=sig)[name=string(\"slh\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dh3 = mul(x=dsilu,y=slh)[name=string(\"dh3\")];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W1t = const()[name=string(\"W1t\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w1t.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [%d,%d,1,1]> W3t = const()[name=string(\"W3t\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/w3t.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dx1 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W1t,x=dh1)[name=string(\"cw1\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dx3 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W3t,x=dh3)[name=string(\"cw3\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> dx = add(x=dx1,y=dx3)[name=string(\"adx\")];\n"+
			"        int32 cax = const()[name=string(\"cax\"), val=int32(1)];\n"+
			"        bool cid = const()[name=string(\"cid\"), val=bool(false)];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = concat(axis=cax,interleave=cid,values=(dx,dh1,dh3))[name=string(\"cat\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		dim+2*hidden, seq,
		dim, seq,
		dim, seq,
		dim,
		hidden, seq,
		hidden, seq,
		dim+hidden,
		hidden, seq,
		hidden, dim, hidden, dim,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		dim, hidden, dim, hidden,
		dim, hidden, dim, hidden,
		dim, seq,
		dim, seq,
		dim, seq,
		dim+2*hidden, seq,
	)
}

// BuildVectorWeightBlob builds a single-baked-weight blob for a 1D fp16 tensor.
func BuildVectorWeightBlob(weights []float32) ([]byte, error) {
	return BuildWeightBlob(weights, 1, len(weights))
}

// BuildFP16Blob builds a generic fp16 MIL BLOBFILE payload from row-major data.
func BuildFP16Blob(data []float32) ([]byte, error) {
	fp16Data := make([]byte, len(data)*2)
	for i, v := range data {
		binary.LittleEndian.PutUint16(fp16Data[i*2:], float32ToFP16(v))
	}
	return buildMILBlob(fp16Data), nil
}

// BuildCausalMaskBlob builds the upper-triangular fp16 causal mask used by SDPA.
func BuildCausalMaskBlob(seq int) ([]byte, error) {
	if seq <= 0 {
		return nil, fmt.Errorf("seq=%d must be > 0", seq)
	}
	mask := make([]float32, seq*seq)
	for t := 0; t < seq; t++ {
		for t2 := 0; t2 < seq; t2++ {
			if t2 > t {
				mask[t*seq+t2] = -65504
			}
		}
	}
	return BuildFP16Blob(mask)
}

// BuildRoPECosSinBlobs builds fp16 cosine/sine tables for RoPE.
// Each output has shape [1,1,seq,headDim], flattened row-major.
func BuildRoPECosSinBlobs(seq, headDim int) ([]byte, []byte, error) {
	if seq <= 0 {
		return nil, nil, fmt.Errorf("seq=%d must be > 0", seq)
	}
	if headDim <= 0 || headDim%2 != 0 {
		return nil, nil, fmt.Errorf("headDim=%d must be even and > 0", headDim)
	}
	half := headDim / 2
	cosTbl := make([]float32, seq*headDim)
	sinTbl := make([]float32, seq*headDim)
	for pos := 0; pos < seq; pos++ {
		row := pos * headDim
		for i := 0; i < half; i++ {
			freq := float64(pos) / math.Pow(10000, float64(2*i)/float64(headDim))
			c := float32(math.Cos(freq))
			s := float32(math.Sin(freq))
			even := row + 2*i
			cosTbl[even] = c
			cosTbl[even+1] = c
			sinTbl[even] = s
			sinTbl[even+1] = s
		}
	}
	cosBlob, err := BuildFP16Blob(cosTbl)
	if err != nil {
		return nil, nil, err
	}
	sinBlob, err := BuildFP16Blob(sinTbl)
	if err != nil {
		return nil, nil, err
	}
	return cosBlob, sinBlob, nil
}

// BuildTransposedWeightBlob builds a baked-weight blob for the transpose of a row-major matrix.
// The input matrix is [rows, cols] row-major; the baked tensor is [cols, rows].
func BuildTransposedWeightBlob(weights []float32, rows, cols int) ([]byte, error) {
	if rows < 0 || cols < 0 {
		return nil, fmt.Errorf("negative dimensions rows=%d cols=%d", rows, cols)
	}
	if len(weights) != rows*cols {
		return nil, fmt.Errorf("weight count=%d want=%d", len(weights), rows*cols)
	}
	transposed := make([]float32, len(weights))
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			transposed[c*rows+r] = weights[r*cols+c]
		}
	}
	return BuildWeightBlob(transposed, cols, rows)
}

func buildMILBlob(fp16Data []byte) []byte {
	const fileHeaderSize = 64
	const chunkHeaderSize = 64
	dataOffset := fileHeaderSize + chunkHeaderSize
	totalSize := dataOffset + len(fp16Data)

	buf := make([]byte, totalSize)
	buf[0] = 0x01
	buf[4] = 0x02

	off := fileHeaderSize
	binary.LittleEndian.PutUint32(buf[off:], 0xDEADBEEF)
	buf[off+4] = 0x01
	binary.LittleEndian.PutUint32(buf[off+8:], uint32(len(fp16Data)))
	binary.LittleEndian.PutUint32(buf[off+16:], uint32(dataOffset))
	copy(buf[dataOffset:], fp16Data)
	return buf
}

// GenSDPAForward generates the fused attention forward block and returns x2 only.
func GenSDPAForward(dim, heads, seq int) string {
	headDim := dim / heads
	scale := 1.0 / math.Sqrt(float64(headDim))

	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", buildInfoHeader)
	fmt.Fprintf(&b, "    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> sq = mul(x=x,y=x)[name=string(\"sq\")];\n", dim, seq)
	b.WriteString("        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([1])];\n")
	b.WriteString("        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,1,%d]> ss = reduce_sum(x=sq,axes=rax,keep_dims=kd)[name=string(\"ss\")];\n", seq)
	fmt.Fprintf(&b, "        fp16 invd = const()[name=string(\"invd\"), val=fp16(%f)];\n", 1.0/float64(dim))
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,1,%d]> ss2 = mul(x=ss,y=invd)[name=string(\"ss2\")];\n", seq)
	b.WriteString("        fp16 eps = const()[name=string(\"eps\"), val=fp16(0.00001)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,1,%d]> ss3 = add(x=ss2,y=eps)[name=string(\"ss3\")];\n", seq)
	b.WriteString("        fp16 nhalf = const()[name=string(\"nhalf\"), val=fp16(-0.5)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,1,%d]> rrms = pow(x=ss3,y=nhalf)[name=string(\"rrms\")];\n", seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> xr = mul(x=x,y=rrms)[name=string(\"xr\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,1]> rw = const()[name=string(\"rw\"), val=tensor<fp16, [1,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/rms1.bin\"), offset=uint64(64)))];\n", dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> xn = mul(x=xr,y=rw)[name=string(\"xn\")];\n", dim, seq)
	appendConvConsts(&b)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wq = const()[name=string(\"Wq\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wq.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wk = const()[name=string(\"Wk\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wk.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wv = const()[name=string(\"Wv\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wv.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wo = const()[name=string(\"Wo\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wo.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> qf = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wq,x=xn)[name=string(\"cq\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> kf = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wk,x=xn)[name=string(\"ck\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> vf = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wv,x=xn)[name=string(\"cv\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> qsh = const()[name=string(\"qsh\"), val=tensor<int32, [4]>([1,%d,%d,%d])];\n", heads, headDim, seq)
	b.WriteString("        tensor<int32, [4]> pm = const()[name=string(\"pm\"), val=tensor<int32, [4]>([0,1,3,2])];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> q4 = reshape(shape=qsh,x=qf)[name=string(\"rq\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> q = transpose(perm=pm,x=q4)[name=string(\"tq\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> k4 = reshape(shape=qsh,x=kf)[name=string(\"rk\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> k = transpose(perm=pm,x=k4)[name=string(\"tk\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> v4 = reshape(shape=qsh,x=vf)[name=string(\"rv\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> v = transpose(perm=pm,x=v4)[name=string(\"tv\")];\n", heads, seq, headDim)
	b.WriteString("        bool tx = const()[name=string(\"tx\"), val=bool(false)];\n")
	b.WriteString("        bool ty = const()[name=string(\"ty\"), val=bool(true)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> sc1 = matmul(transpose_x=tx,transpose_y=ty,x=q,y=k)[name=string(\"mm1\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        fp16 scv = const()[name=string(\"scv\"), val=fp16(%f)];\n", scale)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> sc2 = mul(x=sc1,y=scv)[name=string(\"scl\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,%d,%d]> cm = const()[name=string(\"cm\"), val=tensor<fp16, [1,1,%d,%d]>(BLOBFILE(path=string(\"@model_path/weights/mask.bin\"), offset=uint64(64)))];\n", seq, seq, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> ms = add(x=sc2,y=cm)[name=string(\"msk\")];\n", heads, seq, seq)
	b.WriteString("        int32 sax = const()[name=string(\"sax\"), val=int32(-1)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> aw = softmax(axis=sax,x=ms)[name=string(\"sm\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> a4 = matmul(transpose_x=tx,transpose_y=tx,x=aw,y=v)[name=string(\"mm2\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> at = transpose(perm=pm,x=a4)[name=string(\"ta\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> os = const()[name=string(\"os\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> af = reshape(shape=os,x=at)[name=string(\"ra\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> oo = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wo,x=af)[name=string(\"co\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> out = add(x=x,y=oo)[name=string(\"x2\")];\n", dim, seq)
	b.WriteString("    } -> (out);\n}\n")
	return b.String()
}

// GenQKVForwardRMS generates the RMSNorm plus QKV projection block.
// Output layout is concat(q, k, v) along the channel dimension.
func GenQKVForwardRMS(dim, seq int) string {
	invd := 1.0 / float64(dim)
	return fmt.Sprintf(
		buildInfoHeader+
			"{\n"+
			"    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n"+
			"        tensor<fp16, [1,%d,1,%d]> sq = mul(x=x,y=x)[name=string(\"sq\")];\n"+
			"        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([1])];\n"+
			"        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss = reduce_sum(x=sq,axes=rax,keep_dims=kd)[name=string(\"ss\")];\n"+
			"        fp16 invd = const()[name=string(\"invd\"), val=fp16(%f)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss2 = mul(x=ss,y=invd)[name=string(\"ss2\")];\n"+
			"        fp16 eps = const()[name=string(\"eps\"), val=fp16(0.00001)];\n"+
			"        tensor<fp16, [1,1,1,%d]> ss3 = add(x=ss2,y=eps)[name=string(\"ss3\")];\n"+
			"        fp16 nhalf = const()[name=string(\"nhalf\"), val=fp16(-0.5)];\n"+
			"        tensor<fp16, [1,1,1,%d]> rrms = pow(x=ss3,y=nhalf)[name=string(\"rrms\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xr = mul(x=x,y=rrms)[name=string(\"xr\")];\n"+
			"        tensor<fp16, [1,%d,1,1]> rw = const()[name=string(\"rw\"), val=tensor<fp16, [1,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/rms1.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> xn = mul(x=xr,y=rw)[name=string(\"xn\")];\n"+
			"        string pt = const()[name=string(\"pt\"), val=string(\"valid\")];\n"+
			"        tensor<int32, [2]> st = const()[name=string(\"st\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        tensor<int32, [4]> pd = const()[name=string(\"pd\"), val=tensor<int32, [4]>([0,0,0,0])];\n"+
			"        tensor<int32, [2]> dl = const()[name=string(\"dl\"), val=tensor<int32, [2]>([1,1])];\n"+
			"        int32 gr = const()[name=string(\"gr\"), val=int32(1)];\n"+
			"        tensor<fp16, [%d,%d,1,1]> Wq = const()[name=string(\"Wq\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wq.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [%d,%d,1,1]> Wk = const()[name=string(\"Wk\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wk.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [%d,%d,1,1]> Wv = const()[name=string(\"Wv\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wv.bin\"), offset=uint64(64)))];\n"+
			"        tensor<fp16, [1,%d,1,%d]> q = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wq,x=xn)[name=string(\"cq\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> k = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wk,x=xn)[name=string(\"ck\")];\n"+
			"        tensor<fp16, [1,%d,1,%d]> v = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wv,x=xn)[name=string(\"cv\")];\n"+
			"        int32 cax = const()[name=string(\"cax\"), val=int32(1)];\n"+
			"        bool cid = const()[name=string(\"cid\"), val=bool(false)];\n"+
			"        tensor<fp16, [1,%d,1,%d]> out = concat(axis=cax,interleave=cid,values=(q,k,v))[name=string(\"cat\")];\n"+
			"    } -> (out);\n"+
			"}\n",
		dim, seq,
		dim, seq,
		seq,
		invd,
		seq,
		seq,
		seq,
		dim, seq,
		dim, dim,
		dim, seq,
		dim, dim, dim, dim,
		dim, dim, dim, dim,
		dim, dim, dim, dim,
		dim, seq,
		dim, seq,
		dim, seq,
		3*dim, seq,
	)
}

// GenSDPAApplyForward generates the attention application block.
// Input0 is x and input1 is concat(q, k, v). Output layout is concat(x2, attn).
func GenSDPAApplyForward(dim, heads, seq int) string {
	headDim := dim / heads
	scale := 1.0 / math.Sqrt(float64(headDim))

	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", buildInfoHeader)
	fmt.Fprintf(&b, "    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x, tensor<fp16, [1, %d, 1, %d]> qkv) {\n", dim, seq, 3*dim, seq)
	appendConvConsts(&b)
	fmt.Fprintf(&b, "        tensor<int32, [4]> sz = const()[name=string(\"sz\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", dim, seq)
	b.WriteString("        tensor<int32, [4]> b0 = const()[name=string(\"b0\"), val=tensor<int32, [4]>([0,0,0,0])];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> qf = slice_by_size(x=qkv,begin=b0,size=sz)[name=string(\"sq\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b1 = const()[name=string(\"b1\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> kf = slice_by_size(x=qkv,begin=b1,size=sz)[name=string(\"sk\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b2 = const()[name=string(\"b2\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", 2*dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> vf = slice_by_size(x=qkv,begin=b2,size=sz)[name=string(\"sv\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> qsh = const()[name=string(\"qsh\"), val=tensor<int32, [4]>([1,%d,%d,%d])];\n", heads, headDim, seq)
	b.WriteString("        tensor<int32, [4]> pm = const()[name=string(\"pm\"), val=tensor<int32, [4]>([0,1,3,2])];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> q4 = reshape(shape=qsh,x=qf)[name=string(\"rq\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> q = transpose(perm=pm,x=q4)[name=string(\"tq\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> k4 = reshape(shape=qsh,x=kf)[name=string(\"rk\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> k = transpose(perm=pm,x=k4)[name=string(\"tk\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> v4 = reshape(shape=qsh,x=vf)[name=string(\"rv\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> v = transpose(perm=pm,x=v4)[name=string(\"tv\")];\n", heads, seq, headDim)
	b.WriteString("        bool tx = const()[name=string(\"tx\"), val=bool(false)];\n")
	b.WriteString("        bool ty = const()[name=string(\"ty\"), val=bool(true)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> sc1 = matmul(transpose_x=tx,transpose_y=ty,x=q,y=k)[name=string(\"mm1\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        fp16 scv = const()[name=string(\"scv\"), val=fp16(%f)];\n", scale)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> sc2 = mul(x=sc1,y=scv)[name=string(\"scl\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,%d,%d]> cm = const()[name=string(\"cm\"), val=tensor<fp16, [1,1,%d,%d]>(BLOBFILE(path=string(\"@model_path/weights/mask.bin\"), offset=uint64(64)))];\n", seq, seq, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> ms = add(x=sc2,y=cm)[name=string(\"msk\")];\n", heads, seq, seq)
	b.WriteString("        int32 sax = const()[name=string(\"sax\"), val=int32(-1)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> aw = softmax(axis=sax,x=ms)[name=string(\"sm\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> a4 = matmul(transpose_x=tx,transpose_y=tx,x=aw,y=v)[name=string(\"mm2\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> at = transpose(perm=pm,x=a4)[name=string(\"ta\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> os = const()[name=string(\"os\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> af = reshape(shape=os,x=at)[name=string(\"ra\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wo = const()[name=string(\"Wo\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wo.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> oo = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wo,x=af)[name=string(\"co\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> x2 = add(x=x,y=oo)[name=string(\"x2\")];\n", dim, seq)
	b.WriteString("        int32 cax = const()[name=string(\"cax\"), val=int32(1)];\n")
	b.WriteString("        bool cid = const()[name=string(\"cid\"), val=bool(false)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> out = concat(axis=cax,interleave=cid,values=(x2,af))[name=string(\"cat\")];\n", 2*dim, seq)
	b.WriteString("    } -> (out);\n}\n")
	return b.String()
}

// GenSDPAForwardTaps generates the fused attention forward block with taps.
// Output layout is concat(x2, q, k, v, attn) along the channel dimension.
func GenSDPAForwardTaps(dim, heads, seq int) string {
	headDim := dim / heads
	scale := 1.0 / math.Sqrt(float64(headDim))

	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", buildInfoHeader)
	fmt.Fprintf(&b, "    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> sq = mul(x=x,y=x)[name=string(\"sq\")];\n", dim, seq)
	b.WriteString("        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([1])];\n")
	b.WriteString("        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,1,%d]> ss = reduce_sum(x=sq,axes=rax,keep_dims=kd)[name=string(\"ss\")];\n", seq)
	fmt.Fprintf(&b, "        fp16 invd = const()[name=string(\"invd\"), val=fp16(%f)];\n", 1.0/float64(dim))
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,1,%d]> ss2 = mul(x=ss,y=invd)[name=string(\"ss2\")];\n", seq)
	b.WriteString("        fp16 eps = const()[name=string(\"eps\"), val=fp16(0.00001)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,1,%d]> ss3 = add(x=ss2,y=eps)[name=string(\"ss3\")];\n", seq)
	b.WriteString("        fp16 nhalf = const()[name=string(\"nhalf\"), val=fp16(-0.5)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,1,%d]> rrms = pow(x=ss3,y=nhalf)[name=string(\"rrms\")];\n", seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> xr = mul(x=x,y=rrms)[name=string(\"xr\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,1]> rw = const()[name=string(\"rw\"), val=tensor<fp16, [1,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/rms1.bin\"), offset=uint64(64)))];\n", dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> xn = mul(x=xr,y=rw)[name=string(\"xn\")];\n", dim, seq)
	appendConvConsts(&b)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wq = const()[name=string(\"Wq\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wq.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wk = const()[name=string(\"Wk\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wk.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wv = const()[name=string(\"Wv\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wv.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wo = const()[name=string(\"Wo\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wo.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> qf = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wq,x=xn)[name=string(\"cq\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> kf = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wk,x=xn)[name=string(\"ck\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> vf = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wv,x=xn)[name=string(\"cv\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> qsh = const()[name=string(\"qsh\"), val=tensor<int32, [4]>([1,%d,%d,%d])];\n", heads, headDim, seq)
	b.WriteString("        tensor<int32, [4]> pm = const()[name=string(\"pm\"), val=tensor<int32, [4]>([0,1,3,2])];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> q4 = reshape(shape=qsh,x=qf)[name=string(\"rq\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> q = transpose(perm=pm,x=q4)[name=string(\"tq\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> k4 = reshape(shape=qsh,x=kf)[name=string(\"rk\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> k = transpose(perm=pm,x=k4)[name=string(\"tk\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> v4 = reshape(shape=qsh,x=vf)[name=string(\"rv\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> v = transpose(perm=pm,x=v4)[name=string(\"tv\")];\n", heads, seq, headDim)
	b.WriteString("        bool tx = const()[name=string(\"tx\"), val=bool(false)];\n")
	b.WriteString("        bool ty = const()[name=string(\"ty\"), val=bool(true)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> sc1 = matmul(transpose_x=tx,transpose_y=ty,x=q,y=k)[name=string(\"mm1\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        fp16 scv = const()[name=string(\"scv\"), val=fp16(%f)];\n", scale)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> sc2 = mul(x=sc1,y=scv)[name=string(\"scl\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,%d,%d]> cm = const()[name=string(\"cm\"), val=tensor<fp16, [1,1,%d,%d]>(BLOBFILE(path=string(\"@model_path/weights/mask.bin\"), offset=uint64(64)))];\n", seq, seq, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> ms = add(x=sc2,y=cm)[name=string(\"msk\")];\n", heads, seq, seq)
	b.WriteString("        int32 sax = const()[name=string(\"sax\"), val=int32(-1)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> aw = softmax(axis=sax,x=ms)[name=string(\"sm\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> a4 = matmul(transpose_x=tx,transpose_y=tx,x=aw,y=v)[name=string(\"mm2\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> at = transpose(perm=pm,x=a4)[name=string(\"ta\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> os = const()[name=string(\"os\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> af = reshape(shape=os,x=at)[name=string(\"ra\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> oo = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wo,x=af)[name=string(\"co\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> x2 = add(x=x,y=oo)[name=string(\"x2\")];\n", dim, seq)
	b.WriteString("        int32 cax = const()[name=string(\"cax\"), val=int32(1)];\n")
	b.WriteString("        bool cid = const()[name=string(\"cid\"), val=bool(false)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> out = concat(axis=cax,interleave=cid,values=(x2,qf,kf,vf,af))[name=string(\"cat\")];\n", 5*dim, seq)
	b.WriteString("    } -> (out);\n}\n")
	return b.String()
}

// GenQKVBackward generates the fused QKV backward kernel.
// Input layout is concat(dq, dk, dv); output layout is dx.
func GenQKVBackward(dim, heads, seq int) string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", buildInfoHeader)
	fmt.Fprintf(&b, "    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n", 3*dim, seq)
	appendConvConsts(&b)
	fmt.Fprintf(&b, "        tensor<int32, [4]> sz = const()[name=string(\"sz\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", dim, seq)
	b.WriteString("        tensor<int32, [4]> b0 = const()[name=string(\"b0\"), val=tensor<int32, [4]>([0,0,0,0])];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dq = slice_by_size(x=x,begin=b0,size=sz)[name=string(\"s0\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b1 = const()[name=string(\"b1\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dk = slice_by_size(x=x,begin=b1,size=sz)[name=string(\"s1\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b2 = const()[name=string(\"b2\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", 2*dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dv = slice_by_size(x=x,begin=b2,size=sz)[name=string(\"s2\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wqt = const()[name=string(\"Wqt\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wqt.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wkt = const()[name=string(\"Wkt\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wkt.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wvt = const()[name=string(\"Wvt\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wvt.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dxq = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wqt,x=dq)[name=string(\"cq\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dxk = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wkt,x=dk)[name=string(\"ck\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dxv = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wvt,x=dv)[name=string(\"cv\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dxqk = add(x=dxq,y=dxk)[name=string(\"aqk\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> out = add(x=dxqk,y=dxv)[name=string(\"out\")];\n", dim, seq)
	b.WriteString("    } -> (out);\n}\n")
	return b.String()
}

// GenSDPABackward1 generates the first SDPA backward kernel plus Wo^T.
// Input layout is concat(q, k, v, dx2); output layout is concat(dv, probs, dp).
func GenSDPABackward1(dim, heads, seq int) string {
	headDim := dim / heads
	scoreCh := heads * seq
	scale := 1.0 / math.Sqrt(float64(headDim))

	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", buildInfoHeader)
	fmt.Fprintf(&b, "    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n", 4*dim, seq)
	appendConvConsts(&b)
	fmt.Fprintf(&b, "        tensor<int32, [4]> sz = const()[name=string(\"sz\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", dim, seq)
	b.WriteString("        tensor<int32, [4]> b0 = const()[name=string(\"b0\"), val=tensor<int32, [4]>([0,0,0,0])];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> qf = slice_by_size(x=x,begin=b0,size=sz)[name=string(\"s0\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b1 = const()[name=string(\"b1\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> kf = slice_by_size(x=x,begin=b1,size=sz)[name=string(\"s1\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b2 = const()[name=string(\"b2\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", 2*dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> vf = slice_by_size(x=x,begin=b2,size=sz)[name=string(\"s2\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b3 = const()[name=string(\"b3\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", 3*dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dx2f = slice_by_size(x=x,begin=b3,size=sz)[name=string(\"s3\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [%d,%d,1,1]> Wot = const()[name=string(\"Wot\"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string(\"@model_path/weights/wot.bin\"), offset=uint64(64)))];\n", dim, dim, dim, dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> df = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=Wot,x=dx2f)[name=string(\"cwo\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> rsh = const()[name=string(\"rsh\"), val=tensor<int32, [4]>([1,%d,%d,%d])];\n", heads, headDim, seq)
	b.WriteString("        tensor<int32, [4]> pm = const()[name=string(\"pm\"), val=tensor<int32, [4]>([0,1,3,2])];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> qr = reshape(shape=rsh,x=qf)[name=string(\"rq\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> q = transpose(perm=pm,x=qr)[name=string(\"tq\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> kr = reshape(shape=rsh,x=kf)[name=string(\"rk\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> k = transpose(perm=pm,x=kr)[name=string(\"tk\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> vr = reshape(shape=rsh,x=vf)[name=string(\"rv\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> v = transpose(perm=pm,x=vr)[name=string(\"tv\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dr = reshape(shape=rsh,x=df)[name=string(\"rd\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> da = transpose(perm=pm,x=dr)[name=string(\"td\")];\n", heads, seq, headDim)
	b.WriteString("        bool bF = const()[name=string(\"bF\"), val=bool(false)];\n")
	b.WriteString("        bool bT = const()[name=string(\"bT\"), val=bool(true)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> sc1 = matmul(transpose_x=bF,transpose_y=bT,x=q,y=k)[name=string(\"mm1\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        fp16 scv = const()[name=string(\"scv\"), val=fp16(%f)];\n", scale)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> sc2 = mul(x=sc1,y=scv)[name=string(\"scl\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,1,%d,%d]> cm = const()[name=string(\"cm\"), val=tensor<fp16, [1,1,%d,%d]>(BLOBFILE(path=string(\"@model_path/weights/mask.bin\"), offset=uint64(64)))];\n", seq, seq, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> ms = add(x=sc2,y=cm)[name=string(\"msk\")];\n", heads, seq, seq)
	b.WriteString("        int32 sax = const()[name=string(\"sax\"), val=int32(-1)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> probs = softmax(axis=sax,x=ms)[name=string(\"sm\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dv4 = matmul(transpose_x=bT,transpose_y=bF,x=probs,y=da)[name=string(\"dv\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dp4 = matmul(transpose_x=bF,transpose_y=bT,x=da,y=v)[name=string(\"dp\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dvt = transpose(perm=pm,x=dv4)[name=string(\"dvt\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> dvs = const()[name=string(\"dvs\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dvf = reshape(shape=dvs,x=dvt)[name=string(\"dvf\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> scs = const()[name=string(\"scs\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", scoreCh, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> pf = reshape(shape=scs,x=probs)[name=string(\"pf\")];\n", scoreCh, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dpf = reshape(shape=scs,x=dp4)[name=string(\"dpf\")];\n", scoreCh, seq)
	b.WriteString("        int32 cax = const()[name=string(\"cax\"), val=int32(1)];\n")
	b.WriteString("        bool cid = const()[name=string(\"cid\"), val=bool(false)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> out = concat(axis=cax,interleave=cid,values=(dvf,pf,dpf))[name=string(\"cat\")];\n", dim+2*scoreCh, seq)
	b.WriteString("    } -> (out);\n}\n")
	return b.String()
}

// GenSDPABackward2 generates the second SDPA backward kernel.
// Input layout is concat(probs, dp, q, k); output layout is concat(dq, dk).
func GenSDPABackward2(dim, heads, seq int) string {
	headDim := dim / heads
	scoreCh := heads * seq
	scale := 1.0 / math.Sqrt(float64(headDim))

	var b strings.Builder
	fmt.Fprintf(&b, "%s{\n", buildInfoHeader)
	fmt.Fprintf(&b, "    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {\n", 2*scoreCh+2*dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> sz_sc = const()[name=string(\"szsc\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", scoreCh, seq)
	b.WriteString("        tensor<int32, [4]> b0 = const()[name=string(\"b0\"), val=tensor<int32, [4]>([0,0,0,0])];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> pf = slice_by_size(x=x,begin=b0,size=sz_sc)[name=string(\"s0\")];\n", scoreCh, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b1 = const()[name=string(\"b1\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", scoreCh)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dpf = slice_by_size(x=x,begin=b1,size=sz_sc)[name=string(\"s1\")];\n", scoreCh, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> sz_d = const()[name=string(\"szd\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b2 = const()[name=string(\"b2\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", 2*scoreCh)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> qf = slice_by_size(x=x,begin=b2,size=sz_d)[name=string(\"s2\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> b3 = const()[name=string(\"b3\"), val=tensor<int32, [4]>([0,%d,0,0])];\n", 2*scoreCh+dim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> kf = slice_by_size(x=x,begin=b3,size=sz_d)[name=string(\"s3\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> ssh = const()[name=string(\"ssh\"), val=tensor<int32, [4]>([1,%d,%d,%d])];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> probs = reshape(shape=ssh,x=pf)[name=string(\"rp\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dp = reshape(shape=ssh,x=dpf)[name=string(\"rdp\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> rsh = const()[name=string(\"rsh\"), val=tensor<int32, [4]>([1,%d,%d,%d])];\n", heads, headDim, seq)
	b.WriteString("        tensor<int32, [4]> pm = const()[name=string(\"pm\"), val=tensor<int32, [4]>([0,1,3,2])];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> qr = reshape(shape=rsh,x=qf)[name=string(\"rq\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> q = transpose(perm=pm,x=qr)[name=string(\"tq\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> kr = reshape(shape=rsh,x=kf)[name=string(\"rk\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> k = transpose(perm=pm,x=kr)[name=string(\"tk\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> pdp = mul(x=probs,y=dp)[name=string(\"pdp\")];\n", heads, seq, seq)
	b.WriteString("        tensor<int32, [1]> rax = const()[name=string(\"rax\"), val=tensor<int32, [1]>([-1])];\n")
	b.WriteString("        bool kd = const()[name=string(\"kd\"), val=bool(true)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,1]> spdp = reduce_sum(x=pdp,axes=rax,keep_dims=kd)[name=string(\"rs\")];\n", heads, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dps = sub(x=dp,y=spdp)[name=string(\"dps\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> ds0 = mul(x=probs,y=dps)[name=string(\"ds0\")];\n", heads, seq, seq)
	fmt.Fprintf(&b, "        fp16 scv = const()[name=string(\"scv\"), val=fp16(%f)];\n", scale)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> ds = mul(x=ds0,y=scv)[name=string(\"ds\")];\n", heads, seq, seq)
	b.WriteString("        bool bF = const()[name=string(\"bF\"), val=bool(false)];\n")
	b.WriteString("        bool bT = const()[name=string(\"bT\"), val=bool(true)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dq4 = matmul(transpose_x=bF,transpose_y=bF,x=ds,y=k)[name=string(\"dq\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dk4 = matmul(transpose_x=bT,transpose_y=bF,x=ds,y=q)[name=string(\"dk\")];\n", heads, seq, headDim)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dqt = transpose(perm=pm,x=dq4)[name=string(\"dqt\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,%d,%d]> dkt = transpose(perm=pm,x=dk4)[name=string(\"dkt\")];\n", heads, headDim, seq)
	fmt.Fprintf(&b, "        tensor<int32, [4]> fs = const()[name=string(\"fs\"), val=tensor<int32, [4]>([1,%d,1,%d])];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dqf = reshape(shape=fs,x=dqt)[name=string(\"dqf\")];\n", dim, seq)
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> dkf = reshape(shape=fs,x=dkt)[name=string(\"dkf\")];\n", dim, seq)
	b.WriteString("        int32 cax = const()[name=string(\"cax\"), val=int32(1)];\n")
	b.WriteString("        bool cid = const()[name=string(\"cid\"), val=bool(false)];\n")
	fmt.Fprintf(&b, "        tensor<fp16, [1,%d,1,%d]> out = concat(axis=cax,interleave=cid,values=(dqf,dkf))[name=string(\"cat\")];\n", 2*dim, seq)
	b.WriteString("    } -> (out);\n}\n")
	return b.String()
}

func appendConvConsts(b *strings.Builder) {
	b.WriteString("        string pt = const()[name=string(\"pt\"), val=string(\"valid\")];\n")
	b.WriteString("        tensor<int32, [2]> st = const()[name=string(\"st\"), val=tensor<int32, [2]>([1,1])];\n")
	b.WriteString("        tensor<int32, [4]> pd = const()[name=string(\"pd\"), val=tensor<int32, [4]>([0,0,0,0])];\n")
	b.WriteString("        tensor<int32, [2]> dl = const()[name=string(\"dl\"), val=tensor<int32, [2]>([1,1])];\n")
	b.WriteString("        int32 gr = const()[name=string(\"gr\"), val=int32(1)];\n")
}
