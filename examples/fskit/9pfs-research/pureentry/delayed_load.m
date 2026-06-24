#import <objc/runtime.h>
#include <mach-o/dyld.h>
#include <mach-o/loader.h>
#include <mach-o/nlist.h>
#include <pthread.h>
#include <string.h>
#include <unistd.h>

static int cstr_eq(const char *a, const char *b) {
	while (*a == *b) { if (*a == 0) return 1; a++; b++; }
	return 0;
}

static unsigned long fileoff_to_vmaddr(const struct mach_header_64 *mh, uint32_t fileoff) {
	const char *p = (const char *)(mh + 1);
	for (uint32_t i = 0; i < mh->ncmds; i++) {
		const struct load_command *lc = (const struct load_command *)p;
		if (lc->cmd == LC_SEGMENT_64) {
			const struct segment_command_64 *seg = (const struct segment_command_64 *)p;
			if ((uint64_t)fileoff >= seg->fileoff && (uint64_t)fileoff < seg->fileoff + seg->filesize)
				return seg->vmaddr + ((uint64_t)fileoff - seg->fileoff);
		}
		p += lc->cmdsize;
	}
	return 0;
}

static void *find_main_symbol(const char *want) {
	const struct mach_header_64 *mh = (const struct mach_header_64 *)_dyld_get_image_header(0);
	long slide = _dyld_get_image_vmaddr_slide(0);
	const char *p = (const char *)(mh + 1);
	const struct symtab_command *symtab = 0;
	for (uint32_t i = 0; i < mh->ncmds; i++) {
		const struct load_command *lc = (const struct load_command *)p;
		if (lc->cmd == LC_SYMTAB) { symtab = (const struct symtab_command *)p; break; }
		p += lc->cmdsize;
	}
	if (!symtab) return 0;
	unsigned long symVM = fileoff_to_vmaddr(mh, symtab->symoff);
	unsigned long strVM = fileoff_to_vmaddr(mh, symtab->stroff);
	if (!symVM || !strVM) return 0;
	const struct nlist_64 *syms = (const struct nlist_64 *)(symVM + slide);
	const char *strs = (const char *)(strVM + slide);
	for (uint32_t i = 0; i < symtab->nsyms; i++)
		if (syms[i].n_un.n_strx && cstr_eq(strs + syms[i].n_un.n_strx, want))
			return (void *)(syms[i].n_value + slide);
	return 0;
}

static void *go_thread(void *arg) {
	(void)arg;
	static char name[] = "NinePFSExtension";
	static char *argv[] = { name, 0 };
	void (*rt0)(int, char **) = find_main_symbol("__rt0_arm64_darwin");
	if (rt0) rt0(1, argv);
	return 0;
}

static void start_go_once(void) {
	static int started;
	if (__sync_bool_compare_and_swap(&started, 0, 1)) {
		pthread_t thread;
		pthread_create(&thread, 0, go_thread, 0);
	}
}

void NinePFSDelayedLoadResource(void *self, id resource, id options, void (^reply)(id, id)) {
	Class cls = objc_getClass("NinePFileSystem");
	SEL sel = sel_registerName("loadResource:options:replyHandler:");
	IMP original = class_getMethodImplementation(cls, sel);
	start_go_once();
	for (int i = 0; i < 200; i++) {
		IMP current = class_getMethodImplementation(cls, sel);
		if (current != original) {
			((void (*)(id, SEL, id, id, id))current)((__bridge id)self, sel, resource, options, reply);
			return;
		}
		usleep(10000);
	}
	reply(nil, nil);
}
