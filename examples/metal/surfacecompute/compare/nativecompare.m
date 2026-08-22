// nativecompare is the Objective-C twin of the Go surfacecompute consumer:
// the same row-sum Metal kernel over an IOSurface-backed R32Float texture,
// timed the same way (best of N reps, whole-surface bytes / wall time).
//
// It exists to answer one question: what does driving Metal from Go via
// purego/objc.Send cost relative to native objc_msgSend? The GPU work is
// identical by construction, so large surfaces should converge and small
// surfaces expose the host-side dispatch overhead.
//
//   clang -O2 -fobjc-arc -framework Metal -framework IOSurface \
//         -framework CoreFoundation -o nativecompare nativecompare.m
//   ./nativecompare 4096 4096 10
#import <Metal/Metal.h>
#import <IOSurface/IOSurface.h>
#import <mach/mach_time.h>
#import <stdio.h>
#import <stdlib.h>

static NSString *const kKernelSource = @""
    "#include <metal_stdlib>\n"
    "using namespace metal;\n"
    "kernel void sumrows(texture2d<float, access::read> tex [[texture(0)]],\n"
    "                    device float *partials [[buffer(0)]],\n"
    "                    uint row [[thread_position_in_grid]]) {\n"
    "    if (row >= tex.get_height()) { return; }\n"
    "    float acc = 0.0f;\n"
    "    const uint w = tex.get_width();\n"
    "    for (uint x = 0; x < w; x++) { acc += tex.read(uint2(x, row)).r; }\n"
    "    partials[row] = acc;\n"
    "}\n";

static double now_sec(void) {
    static mach_timebase_info_data_t tb;
    if (tb.denom == 0) { mach_timebase_info(&tb); }
    return (double)mach_absolute_time() * tb.numer / tb.denom / 1e9;
}

int main(int argc, char **argv) {
    int width = argc > 1 ? atoi(argv[1]) : 4096;
    int height = argc > 2 ? atoi(argv[2]) : 4096;
    int reps = argc > 3 ? atoi(argv[3]) : 10;
    long bytes = (long)width * height * 4;

    @autoreleasepool {
        NSDictionary *props = @{
            (id)kIOSurfaceWidth : @(width),
            (id)kIOSurfaceHeight : @(height),
            (id)kIOSurfaceBytesPerElement : @4,
            (id)kIOSurfaceBytesPerRow : @(width * 4),
            (id)kIOSurfacePixelFormat : @(0x4c303066), // 'L00f'
        };
        IOSurfaceRef surf = IOSurfaceCreate((__bridge CFDictionaryRef)props);
        if (!surf) { fprintf(stderr, "IOSurfaceCreate failed\n"); return 1; }

        IOSurfaceLock(surf, 0, NULL);
        float *data = IOSurfaceGetBaseAddress(surf);
        double want = 0;
        for (long i = 0; i < (long)width * height; i++) {
            data[i] = (float)(i % 251) * 0.5f;
            want += data[i];
        }
        IOSurfaceUnlock(surf, 0, NULL);

        id<MTLDevice> device = MTLCreateSystemDefaultDevice();
        if (!device) { fprintf(stderr, "no Metal device\n"); return 1; }

        MTLTextureDescriptor *desc = [MTLTextureDescriptor
            texture2DDescriptorWithPixelFormat:MTLPixelFormatR32Float
                                         width:width
                                        height:height
                                     mipmapped:NO];
        desc.usage = MTLTextureUsageShaderRead;
        desc.storageMode = MTLStorageModeShared;
        id<MTLTexture> tex = [device newTextureWithDescriptor:desc
                                                    iosurface:surf
                                                        plane:0];
        if (!tex) { fprintf(stderr, "texture from IOSurface failed\n"); return 1; }

        NSError *err = nil;
        id<MTLLibrary> lib = [device newLibraryWithSource:kKernelSource
                                                  options:nil
                                                    error:&err];
        if (!lib) { fprintf(stderr, "compile: %s\n", err.description.UTF8String); return 1; }
        id<MTLComputePipelineState> pso =
            [device newComputePipelineStateWithFunction:[lib newFunctionWithName:@"sumrows"]
                                                  error:&err];
        if (!pso) { fprintf(stderr, "pipeline: %s\n", err.description.UTF8String); return 1; }

        id<MTLCommandQueue> queue = [device newCommandQueue];
        id<MTLBuffer> partials = [device newBufferWithLength:height * 4
                                                     options:MTLResourceStorageModeShared];

        NSUInteger tg = pso.maxTotalThreadsPerThreadgroup;
        if (tg > (NSUInteger)height) { tg = height; }

        double best = 1e18, got = 0;
        for (int r = 0; r < reps; r++) {
            double start = now_sec();
            id<MTLCommandBuffer> cb = [queue commandBuffer];
            id<MTLComputeCommandEncoder> enc = [cb computeCommandEncoder];
            [enc setComputePipelineState:pso];
            [enc setTexture:tex atIndex:0];
            [enc setBuffer:partials offset:0 atIndex:0];
            [enc dispatchThreads:MTLSizeMake(height, 1, 1)
                threadsPerThreadgroup:MTLSizeMake(tg, 1, 1)];
            [enc endEncoding];
            [cb commit];
            [cb waitUntilCompleted];
            float *p = partials.contents;
            got = 0;
            for (int i = 0; i < height; i++) { got += p[i]; }
            double d = now_sec() - start;
            if (d < best) { best = d; }
        }

        double tol = (want > 1 ? want : 1) * 1e-7 * width;
        if (got < want - tol || got > want + tol) {
            fprintf(stderr, "checksum mismatch: got %.0f want %.0f\n", got, want);
            return 1;
        }
        printf("objc gpu %8.2f GiB/s  (%ld MiB in %.3fms)\n",
               bytes / best / (1 << 30), bytes >> 20, best * 1e3);
        CFRelease(surf);
    }
    return 0;
}
