#include <metal_stdlib>
using namespace metal;

struct ScatterParams {
    uint channels;
    uint width;
    uint height;
    uint elem_size;
    uint row_stride;
    uint plane_stride;
    uint count;
};

kernel void scatter_contiguous_to_planar(
    const device float* src [[buffer(0)]],
    device char* dst [[buffer(1)]],
    constant ScatterParams& params [[buffer(2)]],
    uint idx [[thread_position_in_grid]])
{
    if (idx >= params.count) return;

    if (params.width == 1 && params.height == 1) {
        *((device float*)(dst + idx * params.plane_stride)) = src[idx];
        return;
    }

    uint row_size = params.channels * params.width;
    uint h = idx / row_size;
    uint rem = idx % row_size;
    uint c = rem / params.width;
    uint w = rem % params.width;
    uint byte_offset = c * params.plane_stride + h * params.row_stride + w * params.elem_size;
    *((device float*)(dst + byte_offset)) = src[idx];
}
