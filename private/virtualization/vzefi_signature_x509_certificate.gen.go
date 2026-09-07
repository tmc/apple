// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/security"
)

// The class instance for the [VZEFISignatureX509Certificate] class.
var (
	_VZEFISignatureX509CertificateClass     VZEFISignatureX509CertificateClass
	_VZEFISignatureX509CertificateClassOnce sync.Once
)

func getVZEFISignatureX509CertificateClass() VZEFISignatureX509CertificateClass {
	_VZEFISignatureX509CertificateClassOnce.Do(func() {
		_VZEFISignatureX509CertificateClass = VZEFISignatureX509CertificateClass{class: objc.GetClass("VZEFISignatureX509Certificate")}
	})
	return _VZEFISignatureX509CertificateClass
}

// GetVZEFISignatureX509CertificateClass returns the class object for VZEFISignatureX509Certificate.
func GetVZEFISignatureX509CertificateClass() VZEFISignatureX509CertificateClass {
	return getVZEFISignatureX509CertificateClass()
}

type VZEFISignatureX509CertificateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZEFISignatureX509CertificateClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEFISignatureX509CertificateClass) Alloc() VZEFISignatureX509Certificate {
	rv := objc.SendIfResponds[VZEFISignatureX509Certificate](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZEFISignatureX509Certificate.Certificate]
//   - [VZEFISignatureX509Certificate.InitWithCertificate]
type VZEFISignatureX509Certificate struct {
	VZEFISignature
}

// VZEFISignatureX509CertificateFromID constructs a [VZEFISignatureX509Certificate] from an objc.ID.
func VZEFISignatureX509CertificateFromID(id objc.ID) VZEFISignatureX509Certificate {
	return VZEFISignatureX509Certificate{VZEFISignature: VZEFISignatureFromID(id)}
}

// Ensure VZEFISignatureX509Certificate implements IVZEFISignatureX509Certificate.
var _ IVZEFISignatureX509Certificate = VZEFISignatureX509Certificate{}

// An interface definition for the [VZEFISignatureX509Certificate] class.
//
// # Methods
//
//   - [IVZEFISignatureX509Certificate.Certificate]
//   - [IVZEFISignatureX509Certificate.InitWithCertificate]
type IVZEFISignatureX509Certificate interface {
	IVZEFISignature

	// Topic: Methods

	Certificate() security.SecCertificateRef
	InitWithCertificate(certificate security.SecCertificateRef) VZEFISignatureX509Certificate
}

// Init initializes the instance.
func (v VZEFISignatureX509Certificate) Init() VZEFISignatureX509Certificate {
	rv := objc.SendIfResponds[VZEFISignatureX509Certificate](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZEFISignatureX509Certificate) Autorelease() VZEFISignatureX509Certificate {
	rv := objc.SendIfResponds[VZEFISignatureX509Certificate](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEFISignatureX509Certificate creates a new VZEFISignatureX509Certificate instance.
func NewVZEFISignatureX509Certificate() VZEFISignatureX509Certificate {
	class := getVZEFISignatureX509CertificateClass()
	rv := objc.SendIfResponds[VZEFISignatureX509Certificate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZEFISignatureX509CertificateWithCertificate(certificate security.SecCertificateRef) VZEFISignatureX509Certificate {
	instance := getVZEFISignatureX509CertificateClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCertificate:"), certificate)
	return VZEFISignatureX509CertificateFromID(rv)
}

func (v VZEFISignatureX509Certificate) InitWithCertificate(certificate security.SecCertificateRef) VZEFISignatureX509Certificate {
	rv := objc.SendIfResponds[VZEFISignatureX509Certificate](v.ID, objc.Sel("initWithCertificate:"), certificate)
	return rv
}

func (v VZEFISignatureX509Certificate) Certificate() security.SecCertificateRef {
	rv := objc.SendIfResponds[security.SecCertificateRef](v.ID, objc.Sel("certificate"))
	return security.SecCertificateRef(rv)
}
