//go:build darwin

package espresso

import (
	"unsafe"

	"github.com/tmc/apple/private/espresso"
)

// pass wraps an Espresso pass class.
type pass struct {
	name string
	p    espresso.IEspressoCustomPass
}

func (p *pass) Name() string { return p.name }

func (p *pass) Run(n *Network) bool {
	// The underlying Espresso pass expects a raw pointer to the espresso_net_t struct,
	// which is the ObjC object returned by Net(). We pass the object ID as the pointer.
	netPtr := n.net.Net().GetID()
	//nolint:unsafeptr // objc ID is a valid pointer to the ObjC object
	return p.p.RunOnNetwork(*(*unsafe.Pointer)(unsafe.Pointer(&netPtr)))
}

// Fusion passes.

func FuseConvBatchNorm() Pass {
	return &pass{"fuse_conv_batchnorm", espresso.NewEspressoPass_fuse_conv_batchnorm()}
}

func FuseAddReLU() Pass {
	return &pass{"fuse_add_and_relu", espresso.NewEspressoPass_fuse_add_and_relu()}
}

func FuseGELU() Pass {
	return &pass{"fuse_gelu_with_erf", espresso.NewEspressoPass_fuse_gelu_with_erf()}
}

func FuseFastGELU() Pass {
	return &pass{"fuse_fast_gelu", espresso.NewEspressoPass_fuse_fast_gelu_1()}
}

func FusePad() Pass {
	return &pass{"fuse_pad", espresso.NewEspressoPass_fuse_pad()}
}

func FuseAffineScale() Pass {
	return &pass{"fuse_affine_scale", espresso.NewEspressoPass_fuse_affine_scale()}
}

func FuseGRUActivation() Pass {
	return &pass{"fuse_gru_activation", espresso.NewEspressoPass_fuse_gru_activation()}
}

// Strength reduction passes.

func ReduceBatchMatMulToInnerProduct() Pass {
	return &pass{"strength_reduction_batch_matmul_to_inner_product", espresso.NewEspressoPass_strength_reduction_batch_matmul_to_inner_product()}
}

func ReduceGatherToSlice() Pass {
	return &pass{"strength_reduction_gather_to_slice", espresso.NewEspressoPass_strength_reduction_gather_to_slice()}
}

func ReduceGatherToLookup() Pass {
	return &pass{"strength_reduction_gather_to_lookup", espresso.NewEspressoPass_strength_reduction_gather_to_lookup()}
}

func RemoveIdentityTransposes() Pass {
	return &pass{"strength_reduction_remove_identity_transposes", espresso.NewEspressoPass_strength_reduction_remove_identity_transposes()}
}

func RemoveReshapeChain() Pass {
	return &pass{"remove_reshape_chain", espresso.NewEspressoPass_remove_reshape_chain()}
}

func ReduceReshapeToFlatten() Pass {
	return &pass{"strength_reduction_reshape_to_flatten", espresso.NewEspressoPass_strength_reduction_reshape_to_flatten()}
}

// Quantization passes.

func QuantizeGathers() Pass {
	return &pass{"quantize_gathers", espresso.NewEspressoPass_quantize_gathers()}
}

func TransformQuantizeKernel() Pass {
	return &pass{"transform_quantize_kernel", espresso.NewEspressoPass_transform_quantize_kernel()}
}

func TransformStaticQuantizeKernel() Pass {
	return &pass{"transform_static_quantize_kernel", espresso.NewEspressoPass_transform_static_quantize_kernel()}
}

// Normalization passes.

func MergePyTorchLayerNorm() Pass {
	return &pass{"merge_pytorch_layernorm", espresso.NewEspressoPass_merge_pytorch_layernorm()}
}

func MergeTFLayerNorm() Pass {
	return &pass{"merge_tf_layernorm", espresso.NewEspressoPass_merge_tf_layernorm()}
}

func MergeChannelNorm() Pass {
	return &pass{"merge_channel_norm", espresso.NewEspressoPass_merge_channel_norm()}
}

func MergeTFNormalization() Pass {
	return &pass{"merge_tf_normalization", espresso.NewEspressoPass_merge_tf_normalization()}
}

// Layout passes.

func RemoveNHWCNCHWTransposes() Pass {
	return &pass{"remove_nhwc_nchw_transposes", espresso.NewEspressoPass_remove_nhwc_nchw_transposes()}
}

func RemoveSeqWWSeqTransposes() Pass {
	return &pass{"remove_seqw_wseq_transposes", espresso.NewEspressoPass_remove_seqw_wseq_transposes()}
}

// Other passes.

func FoldConstants() Pass {
	return &pass{"fold_constants", espresso.NewEspressoPass_fold_constants()}
}

func FindSharedWeights() Pass {
	return &pass{"find_shared_weights", espresso.NewEspressoPass_find_shared_weights()}
}

func CompressSegments() Pass {
	return &pass{"compress_segments", espresso.NewEspressoPass_compress_segments()}
}

func TransposeInnerProductWeights() Pass {
	return &pass{"transpose_inner_product_weights", espresso.NewEspressoPass_transpose_inner_product_weights()}
}

// Optimize applies a standard set of optimization passes to the network.
func Optimize(n *Network) error {
	return OptimizeWith(n,
		FoldConstants(),
		FuseConvBatchNorm(),
		FuseAddReLU(),
		FuseGELU(),
		FuseAffineScale(),
		FusePad(),
		RemoveIdentityTransposes(),
		RemoveReshapeChain(),
		ReduceBatchMatMulToInnerProduct(),
		ReduceGatherToSlice(),
		RemoveNHWCNCHWTransposes(),
		FindSharedWeights(),
		CompressSegments(),
	)
}

// OptimizeWith applies the given passes in order.
func OptimizeWith(n *Network, passes ...Pass) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	if n.closed {
		return ErrClosed
	}
	for _, p := range passes {
		p.Run(n)
	}
	return nil
}
