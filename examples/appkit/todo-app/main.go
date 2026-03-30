// Command todo-app builds a small multi-window todo list in AppKit.
//
// It is intentionally larger than hello-window and counter. Read those first
// if you want the smallest possible AppKit starting point.
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
	margin     = 20.0
	rowHeight  = 36.0
	rowSpacing = 1.0
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
			Size:   corefoundation.CGSize{Width: 480, Height: 520},
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
	mainWindow.Center()

	contentView := mainWindow.ContentView().(appkit.NSView)
	colors := appkit.GetNSColorClass()
	font := appkit.GetNSFontClass()

	// --- Input area (top): text field + Add button ---
	todoInput = appkit.NewNSTextField()
	todoInput.SetPlaceholderString("What needs to be done?")
	todoInput.SetFont(font.SystemFontOfSize(13))
	todoInput.SetActionHandler(func() { addTodo() })
	todoInput.SetTranslatesAutoresizingMaskIntoConstraints(false)

	addBtn = appkit.NewButtonWithTitleTargetAction("Add", nil, 0)
	addBtn.SetBezelStyle(appkit.NSBezelStylePush)
	addBtn.SetToolTip("Add a new todo item")
	addBtn.SetActionHandler(func() { addTodo() })
	addBtn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	addBtn.WidthAnchor().ConstraintEqualToConstant(70).SetActive(true)

	inputStack := appkit.GetNSStackViewClass().Alloc().Init()
	inputStack.SetOrientation(appkit.NSUserInterfaceLayoutOrientationHorizontal)
	inputStack.SetSpacing(8)
	inputStack.AddArrangedSubview(todoInput)
	inputStack.AddArrangedSubview(addBtn)
	inputStack.SetTranslatesAutoresizingMaskIntoConstraints(false)
	// Let the text field stretch while the button stays fixed width.
	todoInput.SetContentHuggingPriorityForOrientation(appkit.NSLayoutPriority(200), appkit.NSLayoutConstraintOrientationHorizontal)

	// --- Separator below input ---
	separator := appkit.GetNSBoxClass().Alloc().Init()
	separator.SetBoxType(appkit.NSBoxSeparator)
	separator.SetTranslatesAutoresizingMaskIntoConstraints(false)

	// --- Scroll view for the list ---
	scrollView = appkit.GetNSScrollViewClass().Alloc().Init()
	scrollView.SetHasVerticalScroller(true)
	scrollView.SetBorderType(appkit.NSNoBorder)
	scrollView.SetDrawsBackground(false)
	scrollView.SetTranslatesAutoresizingMaskIntoConstraints(false)

	todoListView = appkit.GetNSViewClass().Alloc().Init()
	todoListView.SetFrame(corefoundation.CGRect{
		Origin: corefoundation.CGPoint{X: 0, Y: 0},
		Size:   corefoundation.CGSize{Width: 440, Height: 400},
	})
	scrollView.SetDocumentView(todoListView)

	// Empty state label (centered over scroll view).
	emptyLabel = appkit.NewTextFieldLabelWithString("No items yet")
	emptyLabel.SetFont(font.SystemFontOfSize(15))
	emptyLabel.SetTextColor(colors.TertiaryLabelColor())
	emptyLabel.SetAlignment(appkit.NSTextAlignmentCenter)
	emptyLabel.SetTranslatesAutoresizingMaskIntoConstraints(false)

	// --- Bottom separator ---
	bottomSep := appkit.GetNSBoxClass().Alloc().Init()
	bottomSep.SetBoxType(appkit.NSBoxSeparator)
	bottomSep.SetTranslatesAutoresizingMaskIntoConstraints(false)

	// --- Bottom bar: status + Settings + Clear All ---
	statusLabel = appkit.NewTextFieldLabelWithString(statusText())
	statusLabel.SetFont(font.SystemFontOfSize(11))
	statusLabel.SetTextColor(colors.SecondaryLabelColor())
	statusLabel.SetTranslatesAutoresizingMaskIntoConstraints(false)
	statusLabel.SetContentHuggingPriorityForOrientation(appkit.NSLayoutPriority(200), appkit.NSLayoutConstraintOrientationHorizontal)

	prefsBtn := appkit.NewButtonWithTitleTargetAction("Settings...", nil, 0)
	prefsBtn.SetBezelStyle(appkit.NSBezelStyleAccessoryBarAction)
	prefsBtn.SetFont(font.SystemFontOfSize(11))
	prefsBtn.SetToolTip("Configure maximum number of items")
	prefsBtn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	prefsBtn.SetActionHandler(func() {
		if isPrefsOpen {
			prefsWindow.MakeKeyAndOrderFront(nil)
			return
		}
		buildPreferencesWindow()
	})

	clearBtn = appkit.NewButtonWithTitleTargetAction("Clear All", nil, 0)
	clearBtn.SetBezelStyle(appkit.NSBezelStyleAccessoryBarAction)
	clearBtn.SetFont(font.SystemFontOfSize(11))
	clearBtn.SetToolTip("Remove all items")
	clearBtn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	clearBtn.SetActionHandler(confirmClearAll)

	bottomStack := appkit.GetNSStackViewClass().Alloc().Init()
	bottomStack.SetOrientation(appkit.NSUserInterfaceLayoutOrientationHorizontal)
	bottomStack.SetSpacing(8)
	bottomStack.AddArrangedSubview(statusLabel)
	bottomStack.AddArrangedSubview(prefsBtn)
	bottomStack.AddArrangedSubview(clearBtn)
	bottomStack.SetTranslatesAutoresizingMaskIntoConstraints(false)

	// Add all subviews to content view.
	contentView.AddSubview(inputStack)
	contentView.AddSubview(separator)
	contentView.AddSubview(scrollView)
	contentView.AddSubview(emptyLabel)
	contentView.AddSubview(bottomSep)
	contentView.AddSubview(bottomStack)

	// Layout constraints.
	// Input area pinned to top.
	inputStack.TopAnchor().ConstraintEqualToAnchorConstant(contentView.TopAnchor(), margin).SetActive(true)
	inputStack.LeadingAnchor().ConstraintEqualToAnchorConstant(contentView.LeadingAnchor(), margin).SetActive(true)
	inputStack.TrailingAnchor().ConstraintEqualToAnchorConstant(contentView.TrailingAnchor(), -margin).SetActive(true)

	// Separator below input.
	separator.TopAnchor().ConstraintEqualToAnchorConstant(inputStack.BottomAnchor(), 12).SetActive(true)
	separator.LeadingAnchor().ConstraintEqualToAnchorConstant(contentView.LeadingAnchor(), margin).SetActive(true)
	separator.TrailingAnchor().ConstraintEqualToAnchorConstant(contentView.TrailingAnchor(), -margin).SetActive(true)

	// Scroll view fills middle area.
	scrollView.TopAnchor().ConstraintEqualToAnchorConstant(separator.BottomAnchor(), 8).SetActive(true)
	scrollView.LeadingAnchor().ConstraintEqualToAnchorConstant(contentView.LeadingAnchor(), margin).SetActive(true)
	scrollView.TrailingAnchor().ConstraintEqualToAnchorConstant(contentView.TrailingAnchor(), -margin).SetActive(true)
	scrollView.BottomAnchor().ConstraintEqualToAnchorConstant(bottomSep.TopAnchor(), -8).SetActive(true)

	// Empty label centered over scroll view.
	emptyLabel.CenterXAnchor().ConstraintEqualToAnchor(scrollView.CenterXAnchor()).SetActive(true)
	emptyLabel.CenterYAnchor().ConstraintEqualToAnchor(scrollView.CenterYAnchor()).SetActive(true)

	// Bottom separator.
	bottomSep.LeadingAnchor().ConstraintEqualToAnchorConstant(contentView.LeadingAnchor(), margin).SetActive(true)
	bottomSep.TrailingAnchor().ConstraintEqualToAnchorConstant(contentView.TrailingAnchor(), -margin).SetActive(true)
	bottomSep.BottomAnchor().ConstraintEqualToAnchorConstant(bottomStack.TopAnchor(), -8).SetActive(true)

	// Bottom bar pinned to bottom.
	bottomStack.LeadingAnchor().ConstraintEqualToAnchorConstant(contentView.LeadingAnchor(), margin).SetActive(true)
	bottomStack.TrailingAnchor().ConstraintEqualToAnchorConstant(contentView.TrailingAnchor(), -margin).SetActive(true)
	bottomStack.BottomAnchor().ConstraintEqualToAnchorConstant(contentView.BottomAnchor(), -12).SetActive(true)

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

	// Settings row: label + field + range hint.
	maxLabel := appkit.NewTextFieldLabelWithString("Maximum todo items:")
	maxLabel.SetFont(font.SystemFontOfSize(13))

	maxItemsField = appkit.NewNSTextField()
	maxItemsField.SetStringValue(fmt.Sprintf("%d", maxTodoItems))
	maxItemsField.SetTranslatesAutoresizingMaskIntoConstraints(false)
	maxItemsField.WidthAnchor().ConstraintEqualToConstant(60).SetActive(true)

	rangeLabel := appkit.NewTextFieldLabelWithString("1 - 50")
	rangeLabel.SetFont(font.SystemFontOfSize(11))
	rangeLabel.SetTextColor(appkit.GetNSColorClass().SecondaryLabelColor())

	settingsRow := appkit.GetNSStackViewClass().Alloc().Init()
	settingsRow.SetOrientation(appkit.NSUserInterfaceLayoutOrientationHorizontal)
	settingsRow.SetSpacing(8)
	settingsRow.AddArrangedSubview(maxLabel)
	settingsRow.AddArrangedSubview(maxItemsField)
	settingsRow.AddArrangedSubview(rangeLabel)

	// Separator.
	sep := appkit.GetNSBoxClass().Alloc().Init()
	sep.SetBoxType(appkit.NSBoxSeparator)

	// Button row: Cancel + Save.
	cancelBtn := appkit.NewButtonWithTitleTargetAction("Cancel", nil, 0)
	cancelBtn.SetBezelStyle(appkit.NSBezelStylePush)
	cancelBtn.SetKeyEquivalent("\033")
	cancelBtn.SetActionHandler(func() {
		prefsWindow.OrderOut(nil)
		isPrefsOpen = false
	})

	saveBtn := appkit.NewButtonWithTitleTargetAction("Save", nil, 0)
	saveBtn.SetBezelStyle(appkit.NSBezelStylePush)
	saveBtn.SetKeyEquivalent("\r")
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

	// Spacer pushes buttons to the right.
	spacer := appkit.GetNSViewClass().Alloc().Init()
	spacer.SetTranslatesAutoresizingMaskIntoConstraints(false)
	spacer.SetContentHuggingPriorityForOrientation(appkit.NSLayoutPriority(1), appkit.NSLayoutConstraintOrientationHorizontal)

	buttonRow := appkit.GetNSStackViewClass().Alloc().Init()
	buttonRow.SetOrientation(appkit.NSUserInterfaceLayoutOrientationHorizontal)
	buttonRow.SetSpacing(8)
	buttonRow.AddArrangedSubview(spacer)
	buttonRow.AddArrangedSubview(cancelBtn)
	buttonRow.AddArrangedSubview(saveBtn)

	// Vertical stack: settings + separator + buttons.
	outer := appkit.GetNSStackViewClass().Alloc().Init()
	outer.SetOrientation(appkit.NSUserInterfaceLayoutOrientationVertical)
	outer.SetSpacing(16)
	outer.SetEdgeInsets(foundation.NSEdgeInsets{Top: 20, Left: 20, Bottom: 16, Right: 20})
	outer.AddArrangedSubview(settingsRow)
	outer.AddArrangedSubview(sep)
	outer.AddArrangedSubview(buttonRow)
	outer.SetTranslatesAutoresizingMaskIntoConstraints(false)

	contentView.AddSubview(outer)

	outer.LeadingAnchor().ConstraintEqualToAnchor(contentView.LeadingAnchor()).SetActive(true)
	outer.TrailingAnchor().ConstraintEqualToAnchor(contentView.TrailingAnchor()).SetActive(true)
	outer.TopAnchor().ConstraintEqualToAnchor(contentView.TopAnchor()).SetActive(true)
	outer.BottomAnchor().ConstraintEqualToAnchor(contentView.BottomAnchor()).SetActive(true)

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

	// Get the scroll view's visible width for row sizing.
	contentWidth := scrollView.Frame().Size.Width
	if contentWidth < 100 {
		contentWidth = 440
	}

	y := 0.0
	for i := len(todos) - 1; i >= 0; i-- {
		idx := i
		text := todos[i]

		row := appkit.GetNSViewClass().Alloc().Init()
		row.SetFrame(corefoundation.CGRect{
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
		num.SetFont(font.MonospacedDigitSystemFontOfSizeWeight(12, appkit.NSFontWeights.Regular))
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

	// Window menu (Cmd+W).
	windowMenuItem := appkit.NewMenuItemWithTitleActionKeyEquivalent("Window", 0, "")
	windowMenu := appkit.NewNSMenu()
	windowMenu.SetTitle("Window")
	windowMenu.AddItemWithTitleActionKeyEquivalent("Close", objc.Sel("performClose:"), "w")
	windowMenu.AddItemWithTitleActionKeyEquivalent("Minimize", objc.Sel("performMiniaturize:"), "m")
	windowMenuItem.SetSubmenu(windowMenu)
	menuBar.AddItem(windowMenuItem)

	return menuBar
}
