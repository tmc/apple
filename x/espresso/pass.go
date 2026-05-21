//go:build darwin

package espresso

import "github.com/tmc/apple/private/espresso"

// pass wraps an Espresso pass class.
type pass struct {
	name string
	p    espresso.IEspressoCustomPass
}

func (p *pass) Name() string { return p.name }

func (p *pass) Run(n *Network) bool {
	// The underlying Espresso pass expects a raw pointer to the espresso_net_t struct,
	// which is the ObjC object returned by Net(). We pass the object ID as the pointer.
	return p.p.RunOnNetwork(n.net.Net())
}

// Fusion passes.

func FuseConvBatchNorm() Pass {
	return &pass{"fuse_conv_batchnorm", espresso.NewEspressoPassFuseConvBatchnorm()}
}

func FuseAddReLU() Pass {
	return &pass{"fuse_add_and_relu", espresso.NewEspressoPassFuseAddAndRelu()}
}

func FuseGELU() Pass {
	return &pass{"fuse_gelu_with_erf", espresso.NewEspressoPassFuseGeluWithErf()}
}

func FuseFastGELU() Pass {
	return &pass{"fuse_fast_gelu", espresso.NewEspressoPassFuseFastGelu1()}
}

func FusePad() Pass {
	return &pass{"fuse_pad", espresso.NewEspressoPassFusePad()}
}

func FuseAffineScale() Pass {
	return &pass{"fuse_affine_scale", espresso.NewEspressoPassFuseAffineScale()}
}

func FuseGRUActivation() Pass {
	return &pass{"fuse_gru_activation", espresso.NewEspressoPassFuseGruActivation()}
}

// Strength reduction passes.

func ReduceBatchMatMulToInnerProduct() Pass {
	return &pass{"strength_reduction_batch_matmul_to_inner_product", espresso.NewEspressoPassStrengthReductionBatchMatmulToInnerProduct()}
}

func ReduceGatherToSlice() Pass {
	return &pass{"strength_reduction_gather_to_slice", espresso.NewEspressoPassStrengthReductionGatherToSlice()}
}

func ReduceGatherToLookup() Pass {
	return &pass{"strength_reduction_gather_to_lookup", espresso.NewEspressoPassStrengthReductionGatherToLookup()}
}

func RemoveIdentityTransposes() Pass {
	return &pass{"strength_reduction_remove_identity_transposes", espresso.NewEspressoPassStrengthReductionRemoveIdentityTransposes()}
}

func RemoveReshapeChain() Pass {
	return &pass{"remove_reshape_chain", espresso.NewEspressoPassRemoveReshapeChain()}
}

func ReduceReshapeToFlatten() Pass {
	return &pass{"strength_reduction_reshape_to_flatten", espresso.NewEspressoPassStrengthReductionReshapeToFlatten()}
}

// Quantization passes.

func QuantizeGathers() Pass {
	return &pass{"quantize_gathers", espresso.NewEspressoPassQuantizeGathers()}
}

func TransformQuantizeKernel() Pass {
	return &pass{"transform_quantize_kernel", espresso.NewEspressoPassTransformQuantizeKernel()}
}

func TransformStaticQuantizeKernel() Pass {
	return &pass{"transform_static_quantize_kernel", espresso.NewEspressoPassTransformStaticQuantizeKernel()}
}

// Normalization passes.

func MergePyTorchLayerNorm() Pass {
	return &pass{"merge_pytorch_layernorm", espresso.NewEspressoPassMergePytorchLayernorm()}
}

func MergeTFLayerNorm() Pass {
	return &pass{"merge_tf_layernorm", espresso.NewEspressoPassMergeTfLayernorm()}
}

func MergeChannelNorm() Pass {
	return &pass{"merge_channel_norm", espresso.NewEspressoPassMergeChannelNorm()}
}

func MergeTFNormalization() Pass {
	return &pass{"merge_tf_normalization", espresso.NewEspressoPassMergeTfNormalization()}
}

// Layout passes.

func RemoveNHWCNCHWTransposes() Pass {
	return &pass{"remove_nhwc_nchw_transposes", espresso.NewEspressoPassRemoveNhwcNchwTransposes()}
}

func RemoveSeqWWSeqTransposes() Pass {
	return &pass{"remove_seqw_wseq_transposes", espresso.NewEspressoPassRemoveSeqwWseqTransposes()}
}

// Other passes.

func FoldConstants() Pass {
	return &pass{"fold_constants", espresso.NewEspressoPassFoldConstants()}
}

func FindSharedWeights() Pass {
	return &pass{"find_shared_weights", espresso.NewEspressoPassFindSharedWeights()}
}

func CompressSegments() Pass {
	return &pass{"compress_segments", espresso.NewEspressoPassCompressSegments()}
}

func TransposeInnerProductWeights() Pass {
	return &pass{"transpose_inner_product_weights", espresso.NewEspressoPassTransposeInnerProductWeights()}
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
