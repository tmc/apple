#import "NinePFileSystem.h"
#import <dispatch/dispatch.h>
#import <os/log.h>
#include <stdio.h>

extern int NinePFSConfigureFileSystem(void *);
extern int NinePFSInit(void);
extern void NinePFSProbeResource(void *, void *, void *);
extern void NinePFSLoadResource(void *, void *, void *, void *);
extern void NinePFSUnloadResource(void *, void *, void *, void *);

static void NinePFSObjCLog(NSString *message) {
	NSLog(@"9pfs: %@", message);
	FILE *f = fopen("/tmp/9pfs-extension.log", "a");
	if (f != NULL) {
		fprintf(f, "9pfs: %s\n", message.UTF8String);
		fclose(f);
	}
}

void NinePFSLogCString(const char *message) {
	if (message == NULL) {
		return;
	}
	os_log_with_type(OS_LOG_DEFAULT, OS_LOG_TYPE_DEFAULT, "9pfs: %{public}s", message);
	NinePFSObjCLog([NSString stringWithUTF8String:message]);
}

@implementation NinePFileSystem

- (void)didFinishLoading {
	NinePFSObjCLog(@"didFinishLoading");
#ifndef NINEPFS_OBJC_CONTROL
	int status = NinePFSInit();
	NinePFSObjCLog(status == 0 ? @"didFinishLoading Go init ok" : @"didFinishLoading Go init failed");
#endif
}

- (void)probeResource:(FSResource *)resource replyHandler:(void (^)(FSProbeResult * _Nullable, NSError * _Nullable))reply {
	NinePFSObjCLog(@"ObjC probeResource begin");
#ifdef NINEPFS_OBJC_CONTROL
	FSContainerIdentifier *containerID = [[FSContainerIdentifier alloc] initWithUUID:[NSUUID UUID]];
	reply([FSProbeResult usableProbeResultWithName:@"9pfs" containerID:containerID], nil);
	NinePFSObjCLog(@"ObjC probeResource control reply");
#else
	void (^replyCopy)(FSProbeResult * _Nullable, NSError * _Nullable) = [reply copy];
	NinePFSObjCLog(@"ObjC probeResource Go begin");
	NinePFSProbeResource((__bridge void *)self, (__bridge void *)resource, (__bridge void *)replyCopy);
	NinePFSObjCLog(@"ObjC probeResource Go end");
#endif
}

- (void)loadResource:(FSResource *)resource options:(FSTaskOptions *)options replyHandler:(void (^)(FSVolume * _Nullable, NSError * _Nullable))reply {
	NinePFSObjCLog(@"ObjC loadResource begin");
#ifdef NINEPFS_OBJC_CONTROL
	NSError *error = [NSError errorWithDomain:FSKitErrorDomain code:FSErrorResourceUnusable userInfo:@{
		NSLocalizedDescriptionKey: @"9pfs ObjC control reached loadResource"
	}];
	reply(nil, error);
	NinePFSObjCLog(@"ObjC loadResource control reply");
#else
	void (^replyCopy)(FSVolume * _Nullable, NSError * _Nullable) = [reply copy];
	NinePFSObjCLog(@"ObjC loadResource Go begin");
	NinePFSLoadResource((__bridge void *)self, (__bridge void *)resource, (__bridge void *)options, (__bridge void *)replyCopy);
	NinePFSObjCLog(@"ObjC loadResource Go end");
#endif
}

- (void)unloadResource:(FSResource *)resource options:(FSTaskOptions *)options replyHandler:(void (^)(NSError * _Nullable))reply {
	NinePFSObjCLog(@"ObjC unloadResource begin");
#ifdef NINEPFS_OBJC_CONTROL
	reply(nil);
	NinePFSObjCLog(@"ObjC unloadResource control reply");
#else
	void (^replyCopy)(NSError * _Nullable) = [reply copy];
	NinePFSObjCLog(@"ObjC unloadResource Go begin");
	NinePFSUnloadResource((__bridge void *)self, (__bridge void *)resource, (__bridge void *)options, (__bridge void *)replyCopy);
	NinePFSObjCLog(@"ObjC unloadResource Go end");
#endif
}

@end

@implementation NinePFSItem
@end

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wprotocol"
#pragma clang diagnostic ignored "-Wobjc-protocol-property-synthesis"
@implementation NinePFSVolume
@end
#pragma clang diagnostic pop

void NinePFSInvokeErrorBlock(id block, id error) {
	((void (^)(id))block)(error);
}

void NinePFSInvokeObjectErrorBlock(id block, id object, id error) {
	((void (^)(id, id))block)(object, error);
}

void NinePFSInvokeItemNameErrorBlock(id block, id item, id name, id error) {
	((void (^)(id, id, id))block)(item, name, error);
}

void NinePFSInvokeVerifierErrorBlock(id block, unsigned long long verifier, id error) {
	((void (^)(unsigned long long, id))block)(verifier, error);
}

void NinePFSInvokeBoolErrorBlock(id block, bool value, id error) {
	((void (^)(bool, id))block)(value, error);
}

void NinePFSInvokeSizeErrorBlock(id block, size_t size, id error) {
	((void (^)(size_t, id))block)(size, error);
}
