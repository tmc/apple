package signpost

import (
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// dsoHandle returns a pointer to a Mach-O image header to pass as the "dso"
// argument of _os_signpost_emit_with_name_impl. In C this is &__dso_handle,
// the header of the image containing the call. There is no linked
// __dso_handle in a purego build, so we use the main executable's header from
// _dyld_get_image_header(0), which the logging system accepts.
var (
	dsoOnce   sync.Once
	dsoHeader unsafe.Pointer

	dyldGetImageHeader func(index uint32) unsafe.Pointer
)

func dsoHandle() unsafe.Pointer {
	dsoOnce.Do(func() {
		sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "_dyld_get_image_header")
		if err != nil || sym == 0 {
			return
		}
		purego.RegisterFunc(&dyldGetImageHeader, sym)
		dsoHeader = dyldGetImageHeader(0) // index 0 is the main executable
	})
	return dsoHeader
}
