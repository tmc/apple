// Todo-app is a multi-window todo list built entirely in Go.
//
// It demonstrates dynamic UI updates, text input, scroll views, and
// multi-window management using the generated AppKit bindings. A
// preferences window lets the user configure the maximum number of items.
package main

import (
	"fmt"
	"runtime"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// Application state.
var (
	app           appkit.NSApplication
	mainWindow    appkit.NSWindow
	prefsWindow   appkit.NSWindow
	todoInput     appkit.NSTextField
	addBtn        appkit.NSButton
	todoListView  appkit.NSView
	scrollView    appkit.NSScrollView
	statusLabel   appkit.NSTextField
	clearBtn      appkit.NSButton
	maxItemsField appkit.NSTextField
	emptyLabel    appkit.NSTextField
	todos         []string
	todoItemViews []appkit.NSView
	maxTodoItems  = 10
	isPrefsOpen   bool
)

// Layout constants following Apple HIG spacing guidelines.
const (
	windowWidth  = 480
	windowHeight = 520
	margin       = 20.0
	rowHeight    = 36.0
	rowSpacing   = 1.0
)

func init() { runtime.LockOSThread() }

func main() {
	app = appkit.GetNSApplicationClass().SharedApplication()
	app.SetActivationPolicy(appkit.NSApplicationActivationPolicyRegular)
	app.SetMainMenu(buildMenuBar())

	delegate := appkit.NewNSApplicationDelegate(appkit.NSApplicationDelegateConfig{
		ShouldTerminateAfterLastWindowClosed: func(_ appkit.NSApplication) bool { return true },
		DidFinishLaunching: func(_ foundation.NSNotification) {
			buildMainWindow()
			mainWindow.MakeKeyAndOrderFront(nil)
			app.Activate()
			todoInput.SelectText(nil)
		},
	})
	app.SetDelegate(delegate)
	app.Run()
}

// addTodo adds the current input text as a new todo item.
func addTodo() {
	text := todoInput.StringValue()
	if text == "" || len(todos) >= maxTodoItems {
		return
	}
	todos = append(todos, text)
	todoInput.SetStringValue("")
	refreshTodoList()
	todoInput.SelectText(nil)
}

// buildMainWindow creates the primary todo list window.
func buildMainWindow() {
	mainWindow = appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 0, Y: 0},
			Size:   corefoundation.CGSize{Width: windowWidth, Height: windowHeight},
		},
		appkit.NSWindowStyleMaskTitled|
			appkit.NSWindowStyleMaskClosable|
			appkit.NSWindowStyleMaskMiniaturizable|
			appkit.NSWindowStyleMaskResizable,
		appkit.NSBackingStoreBuffered, false,
	)
	mainWindow.SetTitle("Todo List")
	mainWindow.SetMinSize(corefoundation.CGSize{Width: 360, Height: 300})
	mainWindow.SetTitlebarSeparatorStyle(appkit.NSTitlebarSeparatorStyleLine)
	mainWindow.SetFrameAutosaveName("TodoListMain")
	mainWindow.Center()

	contentView := mainWindow.ContentView().(appkit.NSView)
	colors := appkit.GetNSColorClass()
	font := appkit.GetNSFontClass()
	contentWidth := windowWidth - 2*margin

	// --- Input area (top) ---
	inputAreaY := windowHeight - 80.0

	todoInput = appkit.NewNSTextField()
	todoInput.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: margin, Y: inputAreaY},
		Size:   corefoundation.CGSize{Width: contentWidth - 80, Height: 28},
	})
	todoInput.SetPlaceholderString("What needs to be done?")
	todoInput.SetFont(font.SystemFontOfSize(13))
	todoInput.SetActionHandler(func() { addTodo() })
	contentView.AddSubview(todoInput)

	addBtn = appkit.NewButtonWithTitleTargetAction("Add", nil, 0)
	addBtn.SetBezelStyle(appkit.NSBezelStylePush)
	addBtn.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: windowWidth - margin - 70, Y: inputAreaY - 1},
		Size:   corefoundation.CGSize{Width: 70, Height: 30},
	})
	addBtn.SetToolTip("Add a new todo item")
	addBtn.SetActionHandler(func() { addTodo() })
	contentView.AddSubview(addBtn)

	// --- Separator below input ---
	separator := appkit.NewBoxWithFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: margin, Y: inputAreaY - 12},
		Size:   corefoundation.CGSize{Width: contentWidth, Height: 1},
	})
	separator.SetBoxType(appkit.NSBoxSeparator)
	contentView.AddSubview(separator)

	// --- Scroll view for the list (main content area) ---
	listTop := inputAreaY - 20.0
	listBottom := 52.0
	listHeight := listTop - listBottom

	scrollView = appkit.GetNSScrollViewClass().Alloc().Init()
	scrollView.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: margin, Y: listBottom},
		Size:   corefoundation.CGSize{Width: contentWidth, Height: listHeight},
	})
	scrollView.SetHasVerticalScroller(true)
	scrollView.SetBorderType(appkit.NSNoBorder)
	scrollView.SetDrawsBackground(false)

	todoListView = appkit.GetNSViewClass().Alloc().InitWithFrame(
		corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 0, Y: 0},
			Size:   corefoundation.CGSize{Width: contentWidth, Height: listHeight},
		},
	)
	scrollView.SetDocumentView(todoListView)
	contentView.AddSubview(scrollView)

	// Empty state label (shown when no todos).
	emptyLabel = appkit.NewTextFieldLabelWithString("No items yet")
	emptyLabel.SetFont(font.SystemFontOfSize(15))
	emptyLabel.SetTextColor(colors.TertiaryLabelColor())
	emptyLabel.SetAlignment(appkit.NSTextAlignmentCenter)
	emptyLabel.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: margin, Y: listBottom + listHeight/2 - 12},
		Size:   corefoundation.CGSize{Width: contentWidth, Height: 24},
	})
	contentView.AddSubview(emptyLabel)

	// --- Bottom bar: status + actions ---
	bottomSep := appkit.NewBoxWithFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: margin, Y: listBottom - 4},
		Size:   corefoundation.CGSize{Width: contentWidth, Height: 1},
	})
	bottomSep.SetBoxType(appkit.NSBoxSeparator)
	contentView.AddSubview(bottomSep)

	statusLabel = appkit.NewTextFieldLabelWithString(statusText())
	statusLabel.SetFont(font.SystemFontOfSize(11))
	statusLabel.SetTextColor(colors.SecondaryLabelColor())
	statusLabel.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: margin, Y: 16},
		Size:   corefoundation.CGSize{Width: 200, Height: 16},
	})
	contentView.AddSubview(statusLabel)

	prefsBtn := appkit.NewButtonWithTitleTargetAction("Settings...", nil, 0)
	prefsBtn.SetBezelStyle(appkit.NSBezelStyleAccessoryBarAction)
	prefsBtn.SetFont(font.SystemFontOfSize(11))
	prefsBtn.SetToolTip("Configure maximum number of items")
	prefsBtn.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: windowWidth - margin - 152, Y: 12},
		Size:   corefoundation.CGSize{Width: 76, Height: 24},
	})
	prefsBtn.SetActionHandler(func() {
		if isPrefsOpen {
			prefsWindow.MakeKeyAndOrderFront(nil)
			return
		}
		buildPreferencesWindow()
	})
	contentView.AddSubview(prefsBtn)

	clearBtn = appkit.NewButtonWithTitleTargetAction("Clear All", nil, 0)
	clearBtn.SetBezelStyle(appkit.NSBezelStyleAccessoryBarAction)
	clearBtn.SetFont(font.SystemFontOfSize(11))
	clearBtn.SetToolTip("Remove all items")
	clearBtn.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: windowWidth - margin - 70, Y: 12},
		Size:   corefoundation.CGSize{Width: 70, Height: 24},
	})
	clearBtn.SetActionHandler(confirmClearAll)
	contentView.AddSubview(clearBtn)

	refreshTodoList()
}

// confirmClearAll shows a confirmation alert before clearing all items.
func confirmClearAll() {
	if len(todos) == 0 {
		return
	}
	alert := appkit.NewNSAlert()
	alert.SetMessageText("Clear all items?")
	alert.SetInformativeText(fmt.Sprintf("This will remove %s from the list. This cannot be undone.", statusText()))
	alert.SetAlertStyle(appkit.NSAlertStyleWarning)
	alert.AddButtonWithTitle("Clear All")
	alert.AddButtonWithTitle("Cancel")
	if alert.RunModal() == 1000 { // NSAlertFirstButtonReturn
		todos = nil
		refreshTodoList()
	}
}

// buildPreferencesWindow creates the settings window.
func buildPreferencesWindow() {
	prefsWindow = appkit.NewWindowWithContentRectStyleMaskBackingDefer(
		corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 0, Y: 0},
			Size:   corefoundation.CGSize{Width: 340, Height: 150},
		},
		appkit.NSWindowStyleMaskTitled|appkit.NSWindowStyleMaskClosable,
		appkit.NSBackingStoreBuffered, false,
	)
	prefsWindow.SetTitle("Preferences")

	prefsDelegate := appkit.NewNSWindowDelegate(appkit.NSWindowDelegateConfig{
		WillClose: func(_ foundation.NSNotification) {
			isPrefsOpen = false
		},
	})
	prefsWindow.SetDelegate(prefsDelegate)

	contentView := prefsWindow.ContentView().(appkit.NSView)
	font := appkit.GetNSFontClass()

	maxLabel := appkit.NewTextFieldLabelWithString("Maximum todo items:")
	maxLabel.SetFont(font.SystemFontOfSize(13))
	maxLabel.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 20, Y: 90},
		Size:   corefoundation.CGSize{Width: 180, Height: 20},
	})
	contentView.AddSubview(maxLabel)

	maxItemsField = appkit.NewNSTextField()
	maxItemsField.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 205, Y: 88},
		Size:   corefoundation.CGSize{Width: 60, Height: 24},
	})
	maxItemsField.SetStringValue(fmt.Sprintf("%d", maxTodoItems))
	contentView.AddSubview(maxItemsField)

	rangeLabel := appkit.NewTextFieldLabelWithString("1 - 50")
	rangeLabel.SetFont(font.SystemFontOfSize(11))
	rangeLabel.SetTextColor(appkit.GetNSColorClass().SecondaryLabelColor())
	rangeLabel.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 270, Y: 91},
		Size:   corefoundation.CGSize{Width: 50, Height: 16},
	})
	contentView.AddSubview(rangeLabel)

	sep := appkit.NewBoxWithFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 20, Y: 60},
		Size:   corefoundation.CGSize{Width: 300, Height: 1},
	})
	sep.SetBoxType(appkit.NSBoxSeparator)
	contentView.AddSubview(sep)

	cancelBtn := appkit.NewButtonWithTitleTargetAction("Cancel", nil, 0)
	cancelBtn.SetBezelStyle(appkit.NSBezelStylePush)
	cancelBtn.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 160, Y: 16},
		Size:   corefoundation.CGSize{Width: 80, Height: 28},
	})
	cancelBtn.SetKeyEquivalent("\033")
	cancelBtn.SetActionHandler(func() {
		prefsWindow.OrderOut(nil)
		isPrefsOpen = false
	})
	contentView.AddSubview(cancelBtn)

	saveBtn := appkit.NewButtonWithTitleTargetAction("Save", nil, 0)
	saveBtn.SetBezelStyle(appkit.NSBezelStylePush)
	saveBtn.SetKeyEquivalent("\r")
	saveBtn.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 246, Y: 16},
		Size:   corefoundation.CGSize{Width: 80, Height: 28},
	})
	saveBtn.SetActionHandler(func() {
		var n int
		fmt.Sscanf(maxItemsField.StringValue(), "%d", &n)
		if n >= 1 && n <= 50 {
			maxTodoItems = n
			refreshTodoList()
		}
		prefsWindow.OrderOut(nil)
		isPrefsOpen = false
	})
	contentView.AddSubview(saveBtn)

	isPrefsOpen = true
	prefsWindow.Center()
	prefsWindow.MakeKeyAndOrderFront(nil)
}

// refreshTodoList rebuilds the todo list UI from the current state.
func refreshTodoList() {
	for _, v := range todoItemViews {
		v.RemoveFromSuperview()
	}
	todoItemViews = nil

	font := appkit.GetNSFontClass()
	colors := appkit.GetNSColorClass()

	if len(todos) == 0 {
		emptyLabel.SetStringValue("No items yet")
	} else {
		emptyLabel.SetStringValue("")
	}

	addBtn.SetEnabled(len(todos) < maxTodoItems)
	clearBtn.SetEnabled(len(todos) > 0)

	contentWidth := windowWidth - 2*margin
	y := 0.0

	for i := len(todos) - 1; i >= 0; i-- {
		idx := i
		text := todos[i]

		row := appkit.GetNSViewClass().Alloc().InitWithFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 0, Y: y},
			Size:   corefoundation.CGSize{Width: contentWidth, Height: rowHeight},
		})

		// Alternating row background.
		if i%2 == 0 {
			bg := appkit.GetNSBoxClass().Alloc().Init()
			bg.SetFrame(corefoundation.CGRect{
				Origin: corefoundation.CGPoint{X: 0, Y: 0},
				Size:   corefoundation.CGSize{Width: contentWidth, Height: rowHeight},
			})
			bg.SetBoxType(appkit.NSBoxCustom)
			bg.SetTitlePosition(appkit.NSNoTitle)
			bg.SetBorderWidth(0)
			bg.SetFillColor(colors.ControlBackgroundColor())
			bg.SetCornerRadius(0)
			row.AddSubview(bg)
		}

		// Item number.
		num := appkit.NewTextFieldLabelWithString(fmt.Sprintf("%d", i+1))
		num.SetFont(font.MonospacedDigitSystemFontOfSizeWeight(12, appkit.NSFontWeights.Regular()))
		num.SetTextColor(colors.TertiaryLabelColor())
		num.SetAlignment(appkit.NSTextAlignmentRight)
		num.SetFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 4, Y: 8},
			Size:   corefoundation.CGSize{Width: 24, Height: 20},
		})
		row.AddSubview(num)

		// Item text.
		label := appkit.NewTextFieldLabelWithString(text)
		label.SetFont(font.SystemFontOfSize(13))
		label.SetFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: 36, Y: 8},
			Size:   corefoundation.CGSize{Width: contentWidth - 110, Height: 20},
		})
		row.AddSubview(label)

		// Remove button.
		del := appkit.NewButtonWithTitleTargetAction("Remove", nil, 0)
		del.SetBezelStyle(appkit.NSBezelStyleAccessoryBarAction)
		del.SetFont(font.SystemFontOfSize(11))
		del.SetFrame(corefoundation.CGRect{
			Origin: corefoundation.CGPoint{X: contentWidth - 68, Y: 6},
			Size:   corefoundation.CGSize{Width: 62, Height: 24},
		})
		del.SetActionHandler(func() {
			if idx < len(todos) {
				todos = append(todos[:idx], todos[idx+1:]...)
				refreshTodoList()
			}
		})
		row.AddSubview(del)

		todoListView.AddSubview(row)
		todoItemViews = append(todoItemViews, row)
		y += rowHeight + rowSpacing
	}

	// Resize document view to fit content.
	h := float64(len(todos)) * (rowHeight + rowSpacing)
	listFrame := scrollView.Frame()
	if h < listFrame.Size.Height {
		h = listFrame.Size.Height
	}
	todoListView.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 0, Y: 0},
		Size:   corefoundation.CGSize{Width: contentWidth, Height: h},
	})
	statusLabel.SetStringValue(statusText())
}

func statusText() string {
	n := len(todos)
	if n == 0 {
		return "No items"
	}
	if n == 1 {
		return "1 item"
	}
	return fmt.Sprintf("%d items", n)
}

func buildMenuBar() appkit.NSMenu {
	menuBar := appkit.NewNSMenu()

	// App menu.
	appMenuItem := appkit.NewMenuItemWithTitleActionKeyEquivalent("", 0, "")
	appMenu := appkit.NewNSMenu()
	appMenu.AddItemWithTitleActionKeyEquivalent("About Todo List", objc.Sel("orderFrontStandardAboutPanel:"), "")
	appMenu.AddItem(appkit.GetNSMenuItemClass().SeparatorItem())
	prefsItem := appkit.NewMenuItemWithTitleActionKeyEquivalent("Settings...", 0, ",")
	prefsItem.SetActionHandler(func() {
		if isPrefsOpen {
			prefsWindow.MakeKeyAndOrderFront(nil)
			return
		}
		buildPreferencesWindow()
	})
	appMenu.AddItem(prefsItem)
	appMenu.AddItem(appkit.GetNSMenuItemClass().SeparatorItem())
	appMenu.AddItemWithTitleActionKeyEquivalent("Quit Todo List", objc.Sel("terminate:"), "q")
	appMenuItem.SetSubmenu(appMenu)
	menuBar.AddItem(appMenuItem)

	// Edit menu (for cut/copy/paste in text fields).
	editMenuItem := appkit.NewMenuItemWithTitleActionKeyEquivalent("Edit", 0, "")
	editMenu := appkit.NewNSMenu()
	editMenu.SetTitle("Edit")
	editMenu.AddItemWithTitleActionKeyEquivalent("Undo", objc.Sel("undo:"), "z")
	editMenu.AddItemWithTitleActionKeyEquivalent("Redo", objc.Sel("redo:"), "Z")
	editMenu.AddItem(appkit.GetNSMenuItemClass().SeparatorItem())
	editMenu.AddItemWithTitleActionKeyEquivalent("Cut", objc.Sel("cut:"), "x")
	editMenu.AddItemWithTitleActionKeyEquivalent("Copy", objc.Sel("copy:"), "c")
	editMenu.AddItemWithTitleActionKeyEquivalent("Paste", objc.Sel("paste:"), "v")
	editMenu.AddItemWithTitleActionKeyEquivalent("Select All", objc.Sel("selectAll:"), "a")
	editMenuItem.SetSubmenu(editMenu)
	menuBar.AddItem(editMenuItem)

	return menuBar
}
