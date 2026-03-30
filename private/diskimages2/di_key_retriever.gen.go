// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIKeyRetriever] class.
var (
	_DIKeyRetrieverClass     DIKeyRetrieverClass
	_DIKeyRetrieverClassOnce sync.Once
)

func getDIKeyRetrieverClass() DIKeyRetrieverClass {
	_DIKeyRetrieverClassOnce.Do(func() {
		_DIKeyRetrieverClass = DIKeyRetrieverClass{class: objc.GetClass("DIKeyRetriever")}
	})
	return _DIKeyRetrieverClass
}

// GetDIKeyRetrieverClass returns the class object for DIKeyRetriever.
func GetDIKeyRetrieverClass() DIKeyRetrieverClass {
	return getDIKeyRetrieverClass()
}

type DIKeyRetrieverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIKeyRetrieverClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIKeyRetrieverClass) Alloc() DIKeyRetriever {
	rv := objc.Send[DIKeyRetriever](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever
type DIKeyRetriever struct {
	objectivec.Object
}

// DIKeyRetrieverFromID constructs a [DIKeyRetriever] from an objc.ID.
func DIKeyRetrieverFromID(id objc.ID) DIKeyRetriever {
	return DIKeyRetriever{objectivec.Object{ID: id}}
}

// Ensure DIKeyRetriever implements IDIKeyRetriever.
var _ IDIKeyRetriever = DIKeyRetriever{}

// An interface definition for the [DIKeyRetriever] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever
type IDIKeyRetriever interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DIKeyRetriever) Init() DIKeyRetriever {
	rv := objc.Send[DIKeyRetriever](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIKeyRetriever) Autorelease() DIKeyRetriever {
	rv := objc.Send[DIKeyRetriever](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIKeyRetriever creates a new DIKeyRetriever instance.
func NewDIKeyRetriever() DIKeyRetriever {
	class := getDIKeyRetrieverClass()
	rv := objc.Send[DIKeyRetriever](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/KKMSKeyWithURL:destKey:destKeySize:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) KKMSKeyWithURLDestKeyDestKeySizeError(url foundation.INSURL, key string, size uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("KKMSKeyWithURL:destKey:destKeySize:error:"), url, unsafe.Pointer(unsafe.StringData(key+"\x00")), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("KKMSKeyWithURL:destKey:destKeySize:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/WKMSKeyWithURL:authData:destKey:destKeySize:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) WKMSKeyWithURLAuthDataDestKeyDestKeySizeError(url foundation.INSURL, data unsafe.Pointer, key string, size uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("WKMSKeyWithURL:authData:destKey:destKeySize:error:"), url, data, unsafe.Pointer(unsafe.StringData(key+"\x00")), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("WKMSKeyWithURL:authData:destKey:destKeySize:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/decryptKeyWithData:destKey:destKeySize:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) DecryptKeyWithDataDestKeyDestKeySizeError(data objectivec.IObject, key string, size uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("decryptKeyWithData:destKey:destKeySize:error:"), data, unsafe.Pointer(unsafe.StringData(key+"\x00")), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("decryptKeyWithData:destKey:destKeySize:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/getRequestWithURL:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) GetRequestWithURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("getRequestWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/hintFormat:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) HintFormat(format objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("hintFormat:"), format)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/jsonResponseWithRequest:session:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) JsonResponseWithRequestSessionError(request objectivec.IObject, session objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("jsonResponseWithRequest:session:error:"), request, session, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/newACEndpointWithEnvironment:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) NewACEndpointWithEnvironmentError(environment objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("newACEndpointWithEnvironment:error:"), environment, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/newDawTokenWithError:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) NewDawTokenWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("newDawTokenWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/newEnvWithDictionary:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) NewEnvWithDictionaryError(dictionary objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("newEnvWithDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/newSessionWithError:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) NewSessionWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("newSessionWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/newUrl:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) NewUrl(url foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("newUrl:"), url)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/postRequestWithURL:session:data:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) PostRequestWithURLSessionDataError(url foundation.INSURL, session objectivec.IObject, data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("postRequestWithURL:session:data:error:"), url, session, data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/requestSynchronousDataWithRequest:session:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) RequestSynchronousDataWithRequestSessionError(request objectivec.IObject, session objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("requestSynchronousDataWithRequest:session:error:"), request, session, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIKeyRetriever/writeHexKeyToBuffer:hexKey:error:
func (_DIKeyRetrieverClass DIKeyRetrieverClass) WriteHexKeyToBufferHexKeyError(buffer string, key objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIKeyRetrieverClass.class), objc.Sel("writeHexKeyToBuffer:hexKey:error:"), unsafe.Pointer(unsafe.StringData(buffer+"\x00")), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeHexKeyToBuffer:hexKey:error: returned NO with nil NSError")
	}
	return rv, nil

}
