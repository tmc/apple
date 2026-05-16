package network

import (
	"testing"
)

func newTestBrowser(t testing.TB) NWBrowser {
	t.Helper()

	descriptor := NWBrowseDescriptorCreateBonjourService("_appledocs._tcp", "local")
	if descriptor.ID == 0 {
		t.Fatal("NWBrowseDescriptorCreateBonjourService returned nil")
	}
	t.Cleanup(descriptor.Release)

	parameters := NWParametersCreate()
	if parameters.ID == 0 {
		t.Fatal("NWParametersCreate returned nil")
	}
	t.Cleanup(parameters.Release)

	browser := NWBrowserCreate(descriptor, parameters)
	if browser.ID == 0 {
		t.Fatal("NWBrowserCreate returned nil")
	}
	t.Cleanup(browser.Release)
	return browser
}

func TestBrowserStateChangedHandlerIsRetained(t *testing.T) {
	browser := newTestBrowser(t)
	key := networkAsyncBlockKey{
		owner:  browser.ID,
		setter: "nw_browser_set_state_changed_handler:0",
	}
	t.Cleanup(func() { clearNetworkAsyncBlock(browser.ID, key.setter) })

	NWBrowserSetStateChangedHandler(browser, func(NWBrowserState, NWError) {})

	networkAsyncBlockMu.Lock()
	block := networkAsyncBlocks[key]
	networkAsyncBlockMu.Unlock()
	if block == 0 {
		t.Fatal("browser state changed handler block was not retained")
	}
}

func TestBrowserStateChangedHandlerReplacementKeepsSingleRegistration(t *testing.T) {
	browser := newTestBrowser(t)
	key := networkAsyncBlockKey{
		owner:  browser.ID,
		setter: "nw_browser_set_state_changed_handler:0",
	}
	t.Cleanup(func() { clearNetworkAsyncBlock(browser.ID, key.setter) })

	NWBrowserSetStateChangedHandler(browser, func(NWBrowserState, NWError) {})

	networkAsyncBlockMu.Lock()
	first := networkAsyncBlocks[key]
	firstCount := len(networkAsyncBlocks)
	networkAsyncBlockMu.Unlock()
	if first == 0 {
		t.Fatal("first retained block missing")
	}

	NWBrowserSetStateChangedHandler(browser, func(NWBrowserState, NWError) {})

	networkAsyncBlockMu.Lock()
	second := networkAsyncBlocks[key]
	secondCount := len(networkAsyncBlocks)
	networkAsyncBlockMu.Unlock()
	if second == 0 {
		t.Fatal("second retained block missing")
	}
	if second == first {
		t.Fatal("setter replacement reused the old retained block")
	}
	if secondCount != firstCount {
		t.Fatalf("retained block count = %d, want %d", secondCount, firstCount)
	}
}
