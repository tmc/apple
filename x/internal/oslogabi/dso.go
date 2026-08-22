package oslogabi

import (
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// DSOHandle returns a Mach-O image header to pass as the "dso" argument of
// _os_log_impl and _os_signpost_emit_with_name_impl. In C this is
// &__dso_handle, the header of the calling image; there is none in a purego
// build, so the main executable's header from _dyld_get_image_header(0) is
// used, which the logging system accepts.
var (
	dsoOnce   sync.Once
	dsoHeader unsafe.Pointer

	dyldGetImageHeader func(index uint32) unsafe.Pointer
)

func DSOHandle() unsafe.Pointer {
	dsoOnce.Do(func() {
		sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "_dyld_get_image_header")
		if err != nil || sym == 0 {
			return
		}
		purego.RegisterFunc(&dyldGetImageHeader, sym)
		dsoHeader = dyldGetImageHeader(0)
	})
	return dsoHeader
}
