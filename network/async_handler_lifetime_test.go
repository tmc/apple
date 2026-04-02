package network

import (
	"testing"
)

func newTestBrowser(t testing.TB) Nw_browser_t {
	t.Helper()

	descriptor := Nw_browse_descriptor_create_bonjour_service("_appledocs._tcp", "local")
	if descriptor.ID == 0 {
		t.Fatal("Nw_browse_descriptor_create_bonjour_service returned nil")
	}
	t.Cleanup(descriptor.Release)

	parameters := Nw_parameters_create()
	if parameters.ID == 0 {
		t.Fatal("Nw_parameters_create returned nil")
	}
	t.Cleanup(parameters.Release)

	browser := Nw_browser_create(descriptor, parameters)
	if browser.ID == 0 {
		t.Fatal("Nw_browser_create returned nil")
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

	Nw_browser_set_state_changed_handler(browser, func(NwBrowserState, Nw_error_t) {})

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

	Nw_browser_set_state_changed_handler(browser, func(NwBrowserState, Nw_error_t) {})

	networkAsyncBlockMu.Lock()
	first := networkAsyncBlocks[key]
	firstCount := len(networkAsyncBlocks)
	networkAsyncBlockMu.Unlock()
	if first == 0 {
		t.Fatal("first retained block missing")
	}

	Nw_browser_set_state_changed_handler(browser, func(NwBrowserState, Nw_error_t) {})

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
