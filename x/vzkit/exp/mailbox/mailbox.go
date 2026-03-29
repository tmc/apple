package mailbox

import (
	"fmt"

	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
)

// Config describes a mailbox device configuration.
type Config struct {
	Attachment objectivec.IObject
}

// NewAttachment creates a mailbox attachment object.
func NewAttachment() (pvz.VZMailboxDeviceAttachment, error) {
	attachment := pvz.NewVZMailboxDeviceAttachment()
	if attachment.ID == 0 {
		return attachment, fmt.Errorf("create mailbox attachment")
	}
	return attachment, nil
}

// NewConfig creates a mailbox device configuration and applies cfg.
func NewConfig(cfg Config) (pvz.VZMailboxDeviceConfiguration, error) {
	device := pvz.NewVZMailboxDeviceConfiguration()
	if device.ID == 0 {
		return device, fmt.Errorf("create mailbox device configuration")
	}
	if cfg.Attachment != nil && cfg.Attachment.GetID() != 0 {
		attachment := pvz.VZMailboxDeviceAttachmentFromID(cfg.Attachment.GetID())
		device.SetAttachment(&attachment)
	}
	return device, nil
}

// Validate checks the mailbox configuration before use.
func Validate(cfg pvz.VZMailboxDeviceConfiguration) error {
	ok, err := cfg.ValidateWithError()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("validate mailbox configuration")
	}
	return nil
}
