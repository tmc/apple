// Command keychain demonstrates macOS Security framework keychain operations
// using the generated security and corefoundation bindings.
//
// It adds a password to the keychain, retrieves it, verifies the round-trip,
// and cleans up.
//
// Usage:
//
//	keychain
package main

import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/security"
)

const kCFStringEncodingUTF8 = uint32(0x08000100)

// Keychain attribute constants loaded as raw CFTypeRef pointers.
// The generated security global vars resolve these to Go strings,
// but SecItemAdd/SecItemCopyMatching/SecItemDelete need the original
// CFStringRef pointers as dictionary keys.
var (
	kSecClass                unsafe.Pointer
	kSecClassGenericPassword unsafe.Pointer
	kSecAttrAccount          unsafe.Pointer
	kSecAttrService          unsafe.Pointer
	kSecValueData            unsafe.Pointer
	kSecReturnData           unsafe.Pointer
	kSecMatchLimit           unsafe.Pointer
	kSecMatchLimitOne        unsafe.Pointer
	kCFBooleanTrue           unsafe.Pointer
)

func init() {
	runtime.LockOSThread()

	secLib, err := purego.Dlopen("/System/Library/Frameworks/Security.framework/Security", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "keychain: load Security: %v\n", err)
		os.Exit(1)
	}
	cfLib, err := purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "keychain: load CoreFoundation: %v\n", err)
		os.Exit(1)
	}

	loadSym := func(lib uintptr, name string) unsafe.Pointer {
		sym, err := purego.Dlsym(lib, name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "keychain: load %s: %v\n", name, err)
			os.Exit(1)
		}
		return *(*unsafe.Pointer)(unsafe.Pointer(sym))
	}

	kSecClass = loadSym(secLib, "kSecClass")
	kSecClassGenericPassword = loadSym(secLib, "kSecClassGenericPassword")
	kSecAttrAccount = loadSym(secLib, "kSecAttrAccount")
	kSecAttrService = loadSym(secLib, "kSecAttrService")
	kSecValueData = loadSym(secLib, "kSecValueData")
	kSecReturnData = loadSym(secLib, "kSecReturnData")
	kSecMatchLimit = loadSym(secLib, "kSecMatchLimit")
	kSecMatchLimitOne = loadSym(secLib, "kSecMatchLimitOne")
	kCFBooleanTrue = loadSym(cfLib, "kCFBooleanTrue")
}

func main() {
	service := "com.example.apple-go-keychain"
	account := "testuser"
	password := "MySecretPassword123!"

	serviceStr := corefoundation.CFStringCreateWithCString(0, service, kCFStringEncodingUTF8)
	accountStr := corefoundation.CFStringCreateWithCString(0, account, kCFStringEncodingUTF8)
	passwordData := cfDataFromBytes([]byte(password))
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(serviceStr))
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(accountStr))
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(passwordData))

	// Delete any existing item first.
	deleteQuery := cfDict(
		kSecClass, kSecClassGenericPassword,
		unsafe.Pointer(kSecAttrService), unsafe.Pointer(serviceStr),
		unsafe.Pointer(kSecAttrAccount), unsafe.Pointer(accountStr),
	)
	security.SecItemDelete(corefoundation.CFDictionaryRef(deleteQuery))
	corefoundation.CFRelease(corefoundation.CFTypeRef(deleteQuery))

	// Add password to keychain.
	fmt.Printf("Adding password for %s/%s\n", service, account)
	addAttrs := cfDict(
		kSecClass, kSecClassGenericPassword,
		unsafe.Pointer(kSecAttrService), unsafe.Pointer(serviceStr),
		unsafe.Pointer(kSecAttrAccount), unsafe.Pointer(accountStr),
		unsafe.Pointer(kSecValueData), unsafe.Pointer(passwordData),
	)
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(addAttrs))

	status := security.SecItemAdd(corefoundation.CFDictionaryRef(addAttrs), nil)
	if status != 0 {
		fmt.Fprintf(os.Stderr, "SecItemAdd failed: OSStatus %d\n", status)
		os.Exit(1)
	}
	fmt.Println("  added")

	// Retrieve password from keychain.
	fmt.Println("Retrieving password")
	retrieveQuery := cfDict(
		kSecClass, kSecClassGenericPassword,
		unsafe.Pointer(kSecAttrService), unsafe.Pointer(serviceStr),
		unsafe.Pointer(kSecAttrAccount), unsafe.Pointer(accountStr),
		unsafe.Pointer(kSecReturnData), kCFBooleanTrue,
		unsafe.Pointer(kSecMatchLimit), kSecMatchLimitOne,
	)
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(retrieveQuery))

	var result corefoundation.CFTypeRef
	status = security.SecItemCopyMatching(corefoundation.CFDictionaryRef(retrieveQuery), &result)
	if status != 0 {
		fmt.Fprintf(os.Stderr, "SecItemCopyMatching failed: OSStatus %d\n", status)
		os.Exit(1)
	}
	defer corefoundation.CFRelease(result)

	retrieved := cfDataToBytes(corefoundation.CFDataRef(result))
	fmt.Printf("  retrieved: %s\n", retrieved)
	if string(retrieved) == password {
		fmt.Println("  round-trip OK")
	} else {
		fmt.Fprintf(os.Stderr, "  mismatch: got %q, want %q\n", retrieved, password)
		os.Exit(1)
	}

	// Clean up.
	fmt.Println("Deleting keychain item")
	cleanupQuery := cfDict(
		kSecClass, kSecClassGenericPassword,
		unsafe.Pointer(kSecAttrService), unsafe.Pointer(serviceStr),
		unsafe.Pointer(kSecAttrAccount), unsafe.Pointer(accountStr),
	)
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(cleanupQuery))

	status = security.SecItemDelete(corefoundation.CFDictionaryRef(cleanupQuery))
	if status != 0 {
		fmt.Fprintf(os.Stderr, "SecItemDelete failed: OSStatus %d\n", status)
		os.Exit(1)
	}
	fmt.Println("  deleted")
}

// cfDict builds a CFDictionary from alternating key/value unsafe.Pointer pairs.
func cfDict(pairs ...unsafe.Pointer) unsafe.Pointer {
	n := len(pairs) / 2
	keys := make([]unsafe.Pointer, n)
	vals := make([]unsafe.Pointer, n)
	for i := 0; i < n; i++ {
		keys[i] = pairs[i*2]
		vals[i] = pairs[i*2+1]
	}
	cbs := corefoundation.KCFTypeDictionaryKeyCallBacks
	vcbs := corefoundation.KCFTypeDictionaryValueCallBacks
	return unsafe.Pointer(corefoundation.CFDictionaryCreate(0,
		unsafe.Pointer(&keys[0]),
		unsafe.Pointer(&vals[0]),
		n,
		&cbs,
		&vcbs,
	))
}

// cfDataFromBytes creates a CFData from a Go byte slice.
func cfDataFromBytes(b []byte) corefoundation.CFDataRef {
	if len(b) == 0 {
		return corefoundation.CFDataCreate(0, nil, 0)
	}
	return corefoundation.CFDataCreate(0, &b[0], len(b))
}

// cfDataToBytes converts a CFData to a Go byte slice.
func cfDataToBytes(data corefoundation.CFDataRef) []byte {
	length := corefoundation.CFDataGetLength(data)
	if length == 0 {
		return nil
	}
	ptr := corefoundation.CFDataGetBytePtr(data)
	return unsafe.Slice(ptr, length)
}
