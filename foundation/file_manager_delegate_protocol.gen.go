// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The interface a file manager’s delegate uses to intervene during operations or if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate
type NSFileManagerDelegate interface {
	objectivec.IObject
}

// NSFileManagerDelegateObject wraps an existing Objective-C object that conforms to the NSFileManagerDelegate protocol.
type NSFileManagerDelegateObject struct {
	objectivec.Object
}

func (o NSFileManagerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileManagerDelegateObjectFromID constructs a [NSFileManagerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileManagerDelegateObjectFromID(id objc.ID) NSFileManagerDelegateObject {
	return NSFileManagerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate if the file manager should move the specified item to the
// new URL.
//
// fileManager: The file manager object that is attempting to move the file or directory.
//
// srcURL: The URL of the file or directory that the file manager wants to move.
//
// dstURL: The URL specifying the new location for the file or directory.
//
// # Return Value
//
// true if the item should be moved or false if it should not be moved. If you
// do not implement this method, the file manager assumes a response of true.
//
// # Discussion
//
// This method is called only once for the item being moved, regardless of
// whether the item is a file, directory, or symbolic link.
//
// This method performs the same task as the
// [FileManagerShouldMoveItemAtPathToPath] method and is preferred over that
// method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldMoveItemAt:to:)
func (o NSFileManagerDelegateObject) FileManagerShouldMoveItemAtURLToURL(fileManager INSFileManager, srcURL INSURL, dstURL INSURL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldMoveItemAtURL:toURL:"), fileManager, srcURL, dstURL)
	return rv
}

// Asks the delegate if the file manager should move the specified item to the
// new path.
//
// fileManager: The file manager object that is attempting to move the file or directory.
//
// srcPath: The path to the file or directory that the file manager wants to move.
//
// dstPath: The new path for the file or directory.
//
// # Return Value
//
// true if the operation should proceed, otherwise false. If you do not
// implement this method, the file manager assumes a response of true.
//
// # Discussion
//
// This method is called only once for the item being moved, regardless of
// whether the item is a file, directory, or symbolic link.
//
// This method performs the same task as the
// [FileManagerShouldMoveItemAtURLToURL] method, which is preferred over this
// method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldMoveItemAtPath:toPath:)
func (o NSFileManagerDelegateObject) FileManagerShouldMoveItemAtPathToPath(fileManager INSFileManager, srcPath string, dstPath string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldMoveItemAtPath:toPath:"), fileManager, objc.String(srcPath), objc.String(dstPath))
	return rv
}

// Asks the delegate if the move operation should continue after an error
// occurs while moving the item at the specified URL.
//
// fileManager: The file manager object that attempted to move the item.
//
// error: The error that occurred while trying to move the item in `srcURL`.
//
// srcURL: The URL of the file or directory that the file manager tried to move.
//
// dstURL: The URL of the intended destination for the item in `srcURL`.
//
// # Return Value
//
// true if the operation should proceed or false if it should be aborted. If
// you do not implement this method, the file manager assumes a response of
// false.
//
// # Discussion
//
// The file manager calls this method when there is a problem moving the item
// to the specified location. If you return true, the file manager proceeds to
// remove the item from its current location as if the move operation had
// completed successfully.
//
// This method performs the same task as the
// [FileManagerShouldProceedAfterErrorMovingItemAtPathToPath] method and is
// preferred over that method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldProceedAfterError:movingItemAt:to:)
func (o NSFileManagerDelegateObject) FileManagerShouldProceedAfterErrorMovingItemAtURLToURL(fileManager INSFileManager, error_ INSError, srcURL INSURL, dstURL INSURL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:movingItemAtURL:toURL:"), fileManager, error_, srcURL, dstURL)
	return rv
}

// Asks the delegate if the move operation should continue after an error
// occurs while moving the item at the specified path.
//
// fileManager: The file manager object that attempted to move the item.
//
// error: The error that occurred while trying to move the item in `srcPath`.
//
// srcPath: The path of the file or directory that the file manager tried to move.
//
// dstPath: The path of the intended destination for the item in `srcPath`.
//
// # Return Value
//
// true if the operation should proceed or false if it should be aborted. If
// you do not implement this method, the file manager assumes a response of
// false.
//
// # Discussion
//
// The file manager calls this method when there is a problem moving the item
// to the specified location. If you return true, the file manager proceeds to
// remove the item from its current location as if the move operation had
// completed successfully.
//
// This method performs the same task as the
// [FileManagerShouldProceedAfterErrorMovingItemAtURLToURL] method, which is
// preferred over this method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldProceedAfterError:movingItemAtPath:toPath:)
func (o NSFileManagerDelegateObject) FileManagerShouldProceedAfterErrorMovingItemAtPathToPath(fileManager INSFileManager, error_ INSError, srcPath string, dstPath string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:movingItemAtPath:toPath:"), fileManager, error_, objc.String(srcPath), objc.String(dstPath))
	return rv
}

// Asks the delegate if the file manager should copy the specified item to the
// new URL.
//
// fileManager: The file manager object that is attempting to copy the file or directory.
//
// srcURL: The URL of the file or directory that the file manager wants to copy.
//
// dstURL: The URL specifying the location for the copied file or directory.
//
// # Return Value
//
// true if the item should be copied or false if the file manager should stop
// copying items associated with the current operation. If you do not
// implement this method, the file manager assumes a response of true.
//
// # Discussion
//
// This method is called once for each item that needs to be copied. Thus, for
// a directory, this method is called once for the directory and once for each
// item in the directory.
//
// This method performs the same task as the
// [FileManagerShouldCopyItemAtPathToPath] method and is preferred over that
// method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldCopyItemAt:to:)
func (o NSFileManagerDelegateObject) FileManagerShouldCopyItemAtURLToURL(fileManager INSFileManager, srcURL INSURL, dstURL INSURL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldCopyItemAtURL:toURL:"), fileManager, srcURL, dstURL)
	return rv
}

// Asks the delegate if the file manager should copy the specified item to the
// new path.
//
// fileManager: The file manager object that is attempting to copy the file or directory.
//
// srcPath: The path to the file or directory that the file manager wants to copy.
//
// dstPath: The new path for the copied file or directory.
//
// # Return Value
//
// true if the item should be copied or false if the file manager should stop
// copying items associated with the current operation. If you do not
// implement this method, the file manager assumes a response of true.
//
// # Discussion
//
// This method is called once for each item that needs to be copied. Thus, for
// a directory, this method is called once for the directory and once for each
// item in the directory.
//
// This method performs the same task as the
// [FileManagerShouldCopyItemAtURLToURL] method, which is preferred over this
// method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldCopyItemAtPath:toPath:)
func (o NSFileManagerDelegateObject) FileManagerShouldCopyItemAtPathToPath(fileManager INSFileManager, srcPath string, dstPath string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldCopyItemAtPath:toPath:"), fileManager, objc.String(srcPath), objc.String(dstPath))
	return rv
}

// Asks the delegate if the move operation should continue after an error
// occurs while copying the item at the specified URL.
//
// fileManager: The file manager object that attempted to copy the item.
//
// error: The error that occurred during the attempt to copy.
//
// srcURL: The URL or a file or directory that `fileManager` is attempting to copy.
//
// dstURL: The URL or a file or directory to which `fileManager` is attempting to
// copy.
//
// # Return Value
//
// true if the operation should proceed or false if it should be aborted. If
// you do not implement this method, the file manager assumes a response of
// false.
//
// # Discussion
//
// The file manager calls this method when there is a problem copying the item
// to the specified location. If you return true, the file manager continues
// copying any other items and ignores the error.
//
// This method performs the same task as the
// [FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath] method and is
// preferred over that method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldProceedAfterError:copyingItemAt:to:)
func (o NSFileManagerDelegateObject) FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager INSFileManager, error_ INSError, srcURL INSURL, dstURL INSURL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:copyingItemAtURL:toURL:"), fileManager, error_, srcURL, dstURL)
	return rv
}

// Asks the delegate if the move operation should continue after an error
// occurs while copying the item at the specified path.
//
// fileManager: The [NSFileManager] object that sent this message.
//
// error: The error that occurred during the attempt to copy.
//
// srcPath: The path or a file or directory that `fileManager` is attempting to copy.
//
// dstPath: The path or a file or directory to which `fileManager` is attempting to
// copy.
//
// # Return Value
//
// true if the operation should proceed or false if it should be aborted. If
// you do not implement this method, the file manager assumes a response of
// false.
//
// # Discussion
//
// The file manager calls this method when there is a problem copying the item
// to the specified location. If you return true, the file manager continues
// copying any other items and ignores the error.
//
// This method performs the same task as the
// [FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL] method, which is
// preferred over this method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldProceedAfterError:copyingItemAtPath:toPath:)
func (o NSFileManagerDelegateObject) FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(fileManager INSFileManager, error_ INSError, srcPath string, dstPath string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:copyingItemAtPath:toPath:"), fileManager, error_, objc.String(srcPath), objc.String(dstPath))
	return rv
}

// Asks the delegate whether the item at the specified URL should be deleted.
//
// fileManager: The file manager object that is attempting to remove the file or directory.
//
// URL: The URL indicating the file or directory that the file manager is
// attempting to delete.
//
// # Return Value
//
// true if the specified item should be removed or false if it should not be
// removed.
//
// # Discussion
//
// Removed items are deleted immediately and not placed in the Trash. If the
// specified item is a directory, returning false prevents both the directory
// and its children from being deleted.
//
// This method performs the same task as the
// [FileManagerShouldRemoveItemAtPath] method and is preferred over that
// method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldRemoveItemAt:)
func (o NSFileManagerDelegateObject) FileManagerShouldRemoveItemAtURL(fileManager INSFileManager, URL INSURL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldRemoveItemAtURL:"), fileManager, URL)
	return rv
}

// Asks the delegate whether the item at the specified path should be deleted.
//
// fileManager: The file manager object that is attempting to remove the file or directory.
//
// path: The path to the file or directory that the file manager is attempting to
// delete.
//
// # Return Value
//
// true if the specified item should be deleted or false if it should not be
// deleted.
//
// # Discussion
//
// Removed items are deleted immediately and not placed in the Trash. If the
// specified item is a directory, returning false prevents both the directory
// and its children from being deleted.
//
// This method performs the same task as the
// [FileManagerShouldRemoveItemAtURL] method, which is preferred over this
// method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldRemoveItemAtPath:)
func (o NSFileManagerDelegateObject) FileManagerShouldRemoveItemAtPath(fileManager INSFileManager, path string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldRemoveItemAtPath:"), fileManager, objc.String(path))
	return rv
}

// Asks the delegate if the operation should continue after an error occurs
// while removing the item at the specified URL.
//
// fileManager: The file manager object that attempted to remove the item.
//
// error: The error that occurred while attempting to remove the item at [URL].
//
// URL: The URL for the file or directory that the file manager tried to delete.
//
// # Return Value
//
// true if the operation should proceed or false if it should be aborted. If
// you do not implement this method, the file manager assumes a response of
// false.
//
// # Discussion
//
// The file manager calls this method when there is a problem deleting the
// item to the specified location. If you return true, the file manager
// continues deleting any remaining items and ignores the error.
//
// This method performs the same task as the
// [FileManagerShouldProceedAfterErrorRemovingItemAtPath] method and is
// preferred over that method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldProceedAfterError:removingItemAt:)
func (o NSFileManagerDelegateObject) FileManagerShouldProceedAfterErrorRemovingItemAtURL(fileManager INSFileManager, error_ INSError, URL INSURL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:removingItemAtURL:"), fileManager, error_, URL)
	return rv
}

// Asks the delegate if the operation should continue after an error occurs
// while removing the item at the specified path.
//
// fileManager: The file manager object that attempted to remove the item.
//
// error: The error that occurred during the attempt to copy.
//
// path: The path for the file or directory that the file manager tried to delete.
//
// # Return Value
//
// true if the operation should proceed or false if it should be aborted. If
// you do not implement this method, the file manager assumes a response of
// false.
//
// # Discussion
//
// The file manager calls this method when there is a problem deleting the
// item to the specified location. If you return true, the file manager
// continues deleting any remaining items and ignores the error.
//
// This method performs the same task as the
// [FileManagerShouldProceedAfterErrorRemovingItemAtURL] method, which is
// preferred over this method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldProceedAfterError:removingItemAtPath:)
func (o NSFileManagerDelegateObject) FileManagerShouldProceedAfterErrorRemovingItemAtPath(fileManager INSFileManager, error_ INSError, path string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:removingItemAtPath:"), fileManager, error_, objc.String(path))
	return rv
}

// Asks the delegate if a hard link should be created between the items at the
// two URLs.
//
// fileManager: The file manager object that is attempting to create the link.
//
// srcURL: The URL identifying the new hard link to be created.
//
// dstURL: The URL identifying the destination of the link.
//
// # Return Value
//
// true if the link should be created or false if it should not be created.
//
// # Discussion
//
// If the item specified by `destURL` is a directory, returning false prevents
// links from being created to both the directory and its children.
//
// This method performs the same task as the
// [FileManagerShouldLinkItemAtPathToPath] method and is preferred over that
// method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldLinkItemAt:to:)
func (o NSFileManagerDelegateObject) FileManagerShouldLinkItemAtURLToURL(fileManager INSFileManager, srcURL INSURL, dstURL INSURL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldLinkItemAtURL:toURL:"), fileManager, srcURL, dstURL)
	return rv
}

// Asks the delegate if a hard link should be created between the items at the
// two paths.
//
// fileManager: The file manager object that is attempting to create the link.
//
// srcPath: The path or a file or directory that `fileManager` is about to attempt to
// link.
//
// dstPath: The path or a file or directory to which `fileManager` is about to attempt
// to link.
//
// # Return Value
//
// true if the operation should proceed, otherwise false.
//
// # Discussion
//
// If the item specified by `destURL` is a directory, returning false prevents
// links from being created to both the directory and its children.
//
// This method performs the same task as the
// [FileManagerShouldLinkItemAtURLToURL] method, which is preferred over this
// method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldLinkItemAtPath:toPath:)
func (o NSFileManagerDelegateObject) FileManagerShouldLinkItemAtPathToPath(fileManager INSFileManager, srcPath string, dstPath string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldLinkItemAtPath:toPath:"), fileManager, objc.String(srcPath), objc.String(dstPath))
	return rv
}

// Asks the delegate if the operation should continue after an error occurs
// while linking to the item at the specified URL.
//
// fileManager: The file manager object that attempted to create the link.
//
// error: The error that occurred during the link attempt.
//
// srcURL: The URL of the attempted link location.
//
// dstURL: The URL of the file or directory that was the destination of the hard link.
//
// # Return Value
//
// true if the operation should proceed or false if it should be aborted. If
// you do not implement this method, the file manager assumes a response of
// false.
//
// # Discussion
//
// The file manager calls this method when there is a problem creating a hard
// link to the item at the specified location. If you return true, the file
// manager continues creating any other links associated with the current
// operation and ignores the error.
//
// This method performs the same task as the
// [FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath] method and is
// preferred over that method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldProceedAfterError:linkingItemAt:to:)
func (o NSFileManagerDelegateObject) FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(fileManager INSFileManager, error_ INSError, srcURL INSURL, dstURL INSURL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:linkingItemAtURL:toURL:"), fileManager, error_, srcURL, dstURL)
	return rv
}

// Asks the delegate if the operation should continue after an error occurs
// while linking to the item at the specified path.
//
// fileManager: The file manager object that attempted to create the link.
//
// error: The error that occurred during the link attempt.
//
// srcPath: The path to the attempted link location.
//
// dstPath: The path to the file or directory that was the destination of the hard
// link.
//
// # Return Value
//
// true if the operation should proceed or false if it should be aborted. If
// you do not implement this method, the file manager assumes a response of
// false.
//
// # Discussion
//
// The file manager calls this method when there is a problem creating a hard
// link to the item at the specified location. If you return true, the file
// manager continues creating any other links associated with the current
// operation and ignores the error.
//
// This method performs the same task as the
// [FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL] method, which is
// preferred over this method in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/Foundation/FileManagerDelegate/fileManager(_:shouldProceedAfterError:linkingItemAtPath:toPath:)
func (o NSFileManagerDelegateObject) FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(fileManager INSFileManager, error_ INSError, srcPath string, dstPath string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:linkingItemAtPath:toPath:"), fileManager, error_, objc.String(srcPath), objc.String(dstPath))
	return rv
}

// NSFileManagerDelegateConfig holds optional typed callbacks for [NSFileManagerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate
type NSFileManagerDelegateConfig struct {

	// Other Methods
	// FileManagerShouldMoveItemAtURLToURL — Asks the delegate if the file manager should move the specified item to the new URL.
	FileManagerShouldMoveItemAtURLToURL func(fileManager NSFileManager, srcURL objectivec.Object, dstURL objectivec.Object) bool
	// FileManagerShouldProceedAfterErrorMovingItemAtURLToURL — Asks the delegate if the move operation should continue after an error occurs while moving the item at the specified URL.
	FileManagerShouldProceedAfterErrorMovingItemAtURLToURL func(fileManager NSFileManager, error_ objectivec.Object, srcURL objectivec.Object, dstURL objectivec.Object) bool
	// FileManagerShouldCopyItemAtURLToURL — Asks the delegate if the file manager should copy the specified item to the new URL.
	FileManagerShouldCopyItemAtURLToURL func(fileManager NSFileManager, srcURL objectivec.Object, dstURL objectivec.Object) bool
	// FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL — Asks the delegate if the move operation should continue after an error occurs while copying the item at the specified URL.
	FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL func(fileManager NSFileManager, error_ objectivec.Object, srcURL objectivec.Object, dstURL objectivec.Object) bool
	// FileManagerShouldRemoveItemAtURL — Asks the delegate whether the item at the specified URL should be deleted.
	FileManagerShouldRemoveItemAtURL func(fileManager NSFileManager, URL objectivec.Object) bool
	// FileManagerShouldProceedAfterErrorRemovingItemAtURL — Asks the delegate if the operation should continue after an error occurs while removing the item at the specified URL.
	FileManagerShouldProceedAfterErrorRemovingItemAtURL func(fileManager NSFileManager, error_ objectivec.Object, URL objectivec.Object) bool
	// FileManagerShouldLinkItemAtURLToURL — Asks the delegate if a hard link should be created between the items at the two URLs.
	FileManagerShouldLinkItemAtURLToURL func(fileManager NSFileManager, srcURL objectivec.Object, dstURL objectivec.Object) bool
	// FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL — Asks the delegate if the operation should continue after an error occurs while linking to the item at the specified URL.
	FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL func(fileManager NSFileManager, error_ objectivec.Object, srcURL objectivec.Object, dstURL objectivec.Object) bool
}

// NewNSFileManagerDelegate creates an Objective-C object implementing the [NSFileManagerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSFileManagerDelegateObject] satisfies the [NSFileManagerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate
func NewNSFileManagerDelegate(config NSFileManagerDelegateConfig) NSFileManagerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSFileManagerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.FileManagerShouldMoveItemAtURLToURL != nil {
		fn := config.FileManagerShouldMoveItemAtURLToURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("fileManager:shouldMoveItemAtURL:toURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, fileManagerID objc.ID, srcURLID objc.ID, dstURLID objc.ID) bool {
				fileManager := NSFileManagerFromID(fileManagerID)
				srcURL := objectivec.ObjectFromID(srcURLID)
				dstURL := objectivec.ObjectFromID(dstURLID)
				return fn(fileManager, srcURL, dstURL)
			},
		})
	}

	if config.FileManagerShouldProceedAfterErrorMovingItemAtURLToURL != nil {
		fn := config.FileManagerShouldProceedAfterErrorMovingItemAtURLToURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("fileManager:shouldProceedAfterError:movingItemAtURL:toURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, fileManagerID objc.ID, error_ID objc.ID, srcURLID objc.ID, dstURLID objc.ID) bool {
				fileManager := NSFileManagerFromID(fileManagerID)
				error_ := objectivec.ObjectFromID(error_ID)
				srcURL := objectivec.ObjectFromID(srcURLID)
				dstURL := objectivec.ObjectFromID(dstURLID)
				return fn(fileManager, error_, srcURL, dstURL)
			},
		})
	}

	if config.FileManagerShouldCopyItemAtURLToURL != nil {
		fn := config.FileManagerShouldCopyItemAtURLToURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("fileManager:shouldCopyItemAtURL:toURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, fileManagerID objc.ID, srcURLID objc.ID, dstURLID objc.ID) bool {
				fileManager := NSFileManagerFromID(fileManagerID)
				srcURL := objectivec.ObjectFromID(srcURLID)
				dstURL := objectivec.ObjectFromID(dstURLID)
				return fn(fileManager, srcURL, dstURL)
			},
		})
	}

	if config.FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL != nil {
		fn := config.FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("fileManager:shouldProceedAfterError:copyingItemAtURL:toURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, fileManagerID objc.ID, error_ID objc.ID, srcURLID objc.ID, dstURLID objc.ID) bool {
				fileManager := NSFileManagerFromID(fileManagerID)
				error_ := objectivec.ObjectFromID(error_ID)
				srcURL := objectivec.ObjectFromID(srcURLID)
				dstURL := objectivec.ObjectFromID(dstURLID)
				return fn(fileManager, error_, srcURL, dstURL)
			},
		})
	}

	if config.FileManagerShouldRemoveItemAtURL != nil {
		fn := config.FileManagerShouldRemoveItemAtURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("fileManager:shouldRemoveItemAtURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, fileManagerID objc.ID, URLID objc.ID) bool {
				fileManager := NSFileManagerFromID(fileManagerID)
				URL := objectivec.ObjectFromID(URLID)
				return fn(fileManager, URL)
			},
		})
	}

	if config.FileManagerShouldProceedAfterErrorRemovingItemAtURL != nil {
		fn := config.FileManagerShouldProceedAfterErrorRemovingItemAtURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("fileManager:shouldProceedAfterError:removingItemAtURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, fileManagerID objc.ID, error_ID objc.ID, URLID objc.ID) bool {
				fileManager := NSFileManagerFromID(fileManagerID)
				error_ := objectivec.ObjectFromID(error_ID)
				URL := objectivec.ObjectFromID(URLID)
				return fn(fileManager, error_, URL)
			},
		})
	}

	if config.FileManagerShouldLinkItemAtURLToURL != nil {
		fn := config.FileManagerShouldLinkItemAtURLToURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("fileManager:shouldLinkItemAtURL:toURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, fileManagerID objc.ID, srcURLID objc.ID, dstURLID objc.ID) bool {
				fileManager := NSFileManagerFromID(fileManagerID)
				srcURL := objectivec.ObjectFromID(srcURLID)
				dstURL := objectivec.ObjectFromID(dstURLID)
				return fn(fileManager, srcURL, dstURL)
			},
		})
	}

	if config.FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL != nil {
		fn := config.FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("fileManager:shouldProceedAfterError:linkingItemAtURL:toURL:"),
			Fn: func(self objc.ID, _cmd objc.SEL, fileManagerID objc.ID, error_ID objc.ID, srcURLID objc.ID, dstURLID objc.ID) bool {
				fileManager := NSFileManagerFromID(fileManagerID)
				error_ := objectivec.ObjectFromID(error_ID)
				srcURL := objectivec.ObjectFromID(srcURLID)
				dstURL := objectivec.ObjectFromID(dstURLID)
				return fn(fileManager, error_, srcURL, dstURL)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSFileManagerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSFileManagerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSFileManagerDelegateObjectFromID(instance)
}
