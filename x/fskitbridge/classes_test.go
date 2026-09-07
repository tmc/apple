package fskitbridge

import (
	"testing"

	"github.com/tmc/apple/objc"
)

func TestRegisterClassesMissingNames(t *testing.T) {
	tests := []struct {
		name string
		cfg  ClassConfig
	}{
		{"all missing", ClassConfig{}},
		{"missing file system", ClassConfig{VolumeName: "V", ItemName: "I"}},
		{"missing volume", ClassConfig{FileSystemName: "F", ItemName: "I"}},
		{"missing item", ClassConfig{FileSystemName: "F", VolumeName: "V"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := RegisterClasses(tt.cfg); err == nil {
				t.Fatal("RegisterClasses() error = nil, want error")
			}
		})
	}
}

func TestRegisterClassesAttachExistingFileSystemMethods(t *testing.T) {
	const (
		fileSystemName = "FSKitBridgeAttachTestFileSystem"
		volumeName     = "FSKitBridgeAttachTestVolume"
		itemName       = "FSKitBridgeAttachTestItem"
	)
	fileSystem, err := objc.RegisterClass(fileSystemName, objc.GetClass("NSObject"), nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	selector := objc.Sel("fskitBridgeAttachTest")
	cfg := ClassConfig{
		FileSystemName:     fileSystemName,
		VolumeName:         volumeName,
		ItemName:           itemName,
		ExistingFileSystem: fileSystem,
		FileSystemProtocols: []string{
			"FSUnaryFileSystemOperations",
		},
		FileSystemMethods: []objc.MethodDef{
			{Cmd: selector, Fn: func(objc.ID, objc.SEL) {}},
		},
	}
	if _, err := RegisterClasses(cfg); err != nil {
		t.Fatal(err)
	}
	instance := objc.Send[objc.ID](objc.ID(fileSystem), objc.Sel("new"))
	if objc.RespondsToSelector(instance, selector) {
		t.Fatal("existing class has method without AttachExistingFileSystemMethods")
	}
	protocol := objc.GetProtocol("FSUnaryFileSystemOperations")
	if protocol == nil {
		t.Fatal("missing FSUnaryFileSystemOperations protocol")
	}
	if objc.Send[bool](instance, objc.Sel("conformsToProtocol:"), protocol) {
		t.Fatal("existing class has protocol without AttachExistingFileSystemMethods")
	}
	cfg.AttachExistingFileSystemMethods = true
	if _, err := RegisterClasses(cfg); err != nil {
		t.Fatal(err)
	}
	if !objc.RespondsToSelector(instance, selector) {
		t.Fatal("existing class does not have attached method")
	}
	if !objc.Send[bool](instance, objc.Sel("conformsToProtocol:"), protocol) {
		t.Fatal("existing class does not have attached protocol")
	}
}
