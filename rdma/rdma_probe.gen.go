// Code generated from Apple documentation for rdma. DO NOT EDIT.

package rdma

// Device describes an RDMA device returned by [Devices].
type Device struct {
	Handle RDMADevice
	Name   string

	list *DeviceList
}

// Available reports whether librdma and the generated probe symbols are available.
func Available() bool {
	return frameworkHandle != 0 &&
		_ibvGetDeviceList != nil &&
		_ibvFreeDeviceList != nil &&
		_ibvGetDeviceName != nil
}

// Devices returns the RDMA devices currently reported by librdma.
//
// It only calls exported lifecycle/probe symbols. It does not enable RDMA and
// may return an empty slice on systems where RDMA over Thunderbolt is disabled.
func Devices() ([]Device, error) {
	list, err := OpenDeviceList()
	if err != nil {
		return nil, err
	}
	return list.Devices(), nil
}
