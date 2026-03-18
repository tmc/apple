// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMatrix] class.
var (
	_NSMatrixClass     NSMatrixClass
	_NSMatrixClassOnce sync.Once
)

func getNSMatrixClass() NSMatrixClass {
	_NSMatrixClassOnce.Do(func() {
		_NSMatrixClass = NSMatrixClass{class: objc.GetClass("NSMatrix")}
	})
	return _NSMatrixClass
}

// GetNSMatrixClass returns the class object for NSMatrix.
func GetNSMatrixClass() NSMatrixClass {
	return getNSMatrixClass()
}

type NSMatrixClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMatrixClass) Alloc() NSMatrix {
	rv := objc.Send[NSMatrix](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A legacy interface for grouping radio buttons or other types of cells
// together.
//
// # Overview
// 
// [NSMatrix] uses flipped coordinates by default. The cells in an [NSMatrix]
// object are numbered by row and column, each starting with 0; for example,
// the top left [NSCell] would be at (0, 0), and the [NSCell] that’s second
// down and third across would be at (1, 2).
// 
// The [NSMatrix] class has the notion of a single selected cell, which is the
// cell that was most recently clicked or that was so designated by a
// [NSMatrix.SelectCellAtRowColumn] or [NSMatrix.SelectCellWithTag] message. The selected cell
// is the cell chosen for action messages except for [PerformClick]
// ([NSCell]), which is assigned to the key cell. (The key cell is generally
// identical to the selected cell, but can be given click focus while leaving
// the selected cell unchanged.) If the user has selected multiple cells, the
// selected cell is the one lowest and furthest to the right in the matrix of
// cells.
//
// # Initializing an NSMatrix Object
//
//   - [NSMatrix.InitWithFrameModeCellClassNumberOfRowsNumberOfColumns]: Initializes and returns a newly allocated matrix of the specified size using cells of the given class.
//   - [NSMatrix.InitWithFrameModePrototypeNumberOfRowsNumberOfColumns]: Initializes and returns a newly allocated matrix of the specified size using the given cell as a prototype.
//
// # Configuring the Matrix Object
//
//   - [NSMatrix.Mode]: The selection mode of the receiver.
//   - [NSMatrix.SetMode]
//   - [NSMatrix.AllowsEmptySelection]: A Boolean that indicates whether a radio-mode matrix supports an empty selection.
//   - [NSMatrix.SetAllowsEmptySelection]
//   - [NSMatrix.SelectionByRect]: A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor.
//   - [NSMatrix.SetSelectionByRect]
//
// # Managing the Cell Class
//
//   - [NSMatrix.Prototype]: The prototype cell that’s copied whenever the matrix creates a new cell.
//   - [NSMatrix.SetPrototype]
//
// # Laying Out the Cells of the Matrix
//
//   - [NSMatrix.AddColumn]: Adds a new column of cells to the right of the last column.
//   - [NSMatrix.AddColumnWithCells]: Adds a new column of cells to the right of the last column, using the given cells.
//   - [NSMatrix.AddRow]: Adds a new row of cells below the last row.
//   - [NSMatrix.AddRowWithCells]: Adds a new row of cells below the last row, using the specified cells.
//   - [NSMatrix.CellFrameAtRowColumn]: Returns the frame rectangle of the cell that would be drawn at the specified location.
//   - [NSMatrix.CellSize]: The size of each cell in the matrix.
//   - [NSMatrix.SetCellSize]
//   - [NSMatrix.GetNumberOfRowsColumns]: Obtains the number of rows and columns in the receiver.
//   - [NSMatrix.InsertColumn]: Inserts a new column of cells at the specified location.
//   - [NSMatrix.InsertColumnWithCells]: Inserts a new column of cells before the specified column, using the given cells.
//   - [NSMatrix.InsertRow]: Inserts a new row of cells before the specified row.
//   - [NSMatrix.InsertRowWithCells]: Inserts a new row of cells before the specified row, using the given cells.
//   - [NSMatrix.IntercellSpacing]: The vertical and horizontal spacing between cells in the matrix.
//   - [NSMatrix.SetIntercellSpacing]
//   - [NSMatrix.MakeCellAtRowColumn]: Creates a new cell at the location specified by the given row and column in the receiver.
//   - [NSMatrix.NumberOfColumns]: The number of columns in the matrix.
//   - [NSMatrix.NumberOfRows]: The number of rows in the matrix.
//   - [NSMatrix.PutCellAtRowColumn]: Replaces the cell at the specified row and column with the new cell.
//   - [NSMatrix.RemoveColumn]: Removes the specified column at from the receiver.
//   - [NSMatrix.RemoveRow]: Removes the specified row from the receiver.
//   - [NSMatrix.RenewRowsColumns]: Changes the number of rows and columns in the receiver.
//   - [NSMatrix.SortUsingFunctionContext]: Sorts the receiver’s cells in ascending order as defined by the specified comparison function.
//   - [NSMatrix.SortUsingSelector]: Sorts the receiver’s cells in ascending order as defined by the comparison method.
//
// # Auto Layout Sizing
//
//   - [NSMatrix.AutorecalculatesCellSize]: A Boolean that indicates whether the matrix auto-recalculates its cell size.
//   - [NSMatrix.SetAutorecalculatesCellSize]
//
// # Finding Matrix Coordinates
//
//   - [NSMatrix.GetRowColumnForPoint]: Indicates whether the specified point lies within one of the cells of the matrix and returns the location of the cell within which the point lies.
//   - [NSMatrix.GetRowColumnOfCell]: Searches the receiver for the specified cell and returns the row and column of the cell
//
// # Managing Attributes of Individual Cells
//
//   - [NSMatrix.SetStateAtRowColumn]: Sets the state of the cell at specified location.
//   - [NSMatrix.SetToolTipForCell]: Sets the tooltip for the cell.
//   - [NSMatrix.ToolTipForCell]: Returns the tooltip for the specified cell.
//
// # Selecting and Deselecting Cells
//
//   - [NSMatrix.SelectCellAtRowColumn]: Selects the cell at the specified row and column within the receiver.
//   - [NSMatrix.SelectCellWithTag]: Selects the last cell with the given tag.
//   - [NSMatrix.KeyCell]: The cell that will be clicked when the user presses the Space bar.
//   - [NSMatrix.SetKeyCell]
//   - [NSMatrix.SetSelectionFromToAnchorHighlight]: Programmatically selects a range of cells.
//   - [NSMatrix.DeselectAllCells]: Deselects all cells in the receiver and, if necessary, redisplays the receiver.
//   - [NSMatrix.DeselectSelectedCell]: Deselects the selected cell or cells.
//
// # Finding Cells
//
//   - [NSMatrix.SelectedCells]: An array containing all of the matrix’s highlighted cells plus its selected cell.
//   - [NSMatrix.SelectedColumn]: The column number of the selected cell.
//   - [NSMatrix.SelectedRow]: The row number of the selected cell.
//   - [NSMatrix.CellAtRowColumn]: Returns the cell at the specified row and column.
//   - [NSMatrix.CellWithTag]: Searches the receiver and returns the last cell matching the specified tag.
//   - [NSMatrix.Cells]: An array containing the cells of the matrix.
//
// # Modifying Graphics Attributes
//
//   - [NSMatrix.BackgroundColor]: The background color of the matrix (the space between the cells).
//   - [NSMatrix.SetBackgroundColor]
//   - [NSMatrix.CellBackgroundColor]: The background color of the matrix’s cells.
//   - [NSMatrix.SetCellBackgroundColor]
//   - [NSMatrix.DrawsBackground]: A Boolean that indicates whether the matrix draws its background.
//   - [NSMatrix.SetDrawsBackground]
//   - [NSMatrix.DrawsCellBackground]: A Boolean that indicates whether the matrix draws the background within each of its cells.
//   - [NSMatrix.SetDrawsCellBackground]
//
// # Editing Text in Cells
//
//   - [NSMatrix.SelectText]: Selects text in the currently selected cell or in the key cell.
//   - [NSMatrix.SelectTextAtRowColumn]: Selects the text in the cell at the specified location and returns the cell.
//   - [NSMatrix.TextShouldBeginEditing]: Requests permission to begin editing text.
//   - [NSMatrix.TextDidBeginEditing]: Invoked when there’s a change in the text after the receiver gains first responder status.
//   - [NSMatrix.TextDidChange]: Invoked when a key-down event or paste operation occurs that changes the receiver’s contents.
//   - [NSMatrix.TextShouldEndEditing]: Requests permission to end editing.
//   - [NSMatrix.TextDidEndEditing]: Invoked when text editing ends.
//
// # Setting Tab Key Behavior
//
//   - [NSMatrix.TabKeyTraversesCells]: A Boolean that indicates whether pressing the Tab key advances the key cell to the next selectable cell.
//   - [NSMatrix.SetTabKeyTraversesCells]
//
// # Managing the Delegate
//
//   - [NSMatrix.Delegate]: The delegate for messages from the field editor.
//   - [NSMatrix.SetDelegate]
//
// # Resizing the Matrix and Its Cells
//
//   - [NSMatrix.AutosizesCells]: A Boolean that indicates whether the cell sizes change when the receiver is resized.
//   - [NSMatrix.SetAutosizesCells]
//   - [NSMatrix.SetValidateSize]: Specifies whether the receiver’s size information is validated.
//   - [NSMatrix.SizeToCells]: Changes the width and the height of the receiver’s frame so it exactly contains the cells.
//
// # Scrolling Cells in the Matrix
//
//   - [NSMatrix.SetScrollable]: Specifies whether the cells in the matrix are scrollable.
//   - [NSMatrix.ScrollCellToVisibleAtRowColumn]: Scrolls the receiver so the specified cell is visible.
//
// # Displaying and Highlighting Cells
//
//   - [NSMatrix.DrawCellAtRowColumn]: Displays the cell at the specified row and column.
//   - [NSMatrix.HighlightCellAtRowColumn]: Highlights or unhighlights the cell at the specified row and column location.
//
// # Managing and Sending Action Messages
//
//   - [NSMatrix.SendAction]: If the selected cell has both an action and a target, sends its action to its target.
//   - [NSMatrix.SendActionToForAllCells]: Iterates through the cells in the receiver, sending the specified selector to an object for each cell.
//   - [NSMatrix.DoubleAction]: The action sent to the target of the receiver when the user double-clicks a cell.
//   - [NSMatrix.SetDoubleAction]
//   - [NSMatrix.SendDoubleAction]: Sends the double-click action message to the target of the receiver.
//
// # Handling Event and Action Messages
//
//   - [NSMatrix.MouseDownFlags]: The flags in effect at the mouse-down event that started the current tracking session.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix
type NSMatrix struct {
	NSControl
}

// NSMatrixFromID constructs a [NSMatrix] from an objc.ID.
//
// A legacy interface for grouping radio buttons or other types of cells
// together.
func NSMatrixFromID(id objc.ID) NSMatrix {
	return NSMatrix{NSControl: NSControlFromID(id)}
}
// NOTE: NSMatrix adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMatrix] class.
//
// # Initializing an NSMatrix Object
//
//   - [INSMatrix.InitWithFrameModeCellClassNumberOfRowsNumberOfColumns]: Initializes and returns a newly allocated matrix of the specified size using cells of the given class.
//   - [INSMatrix.InitWithFrameModePrototypeNumberOfRowsNumberOfColumns]: Initializes and returns a newly allocated matrix of the specified size using the given cell as a prototype.
//
// # Configuring the Matrix Object
//
//   - [INSMatrix.Mode]: The selection mode of the receiver.
//   - [INSMatrix.SetMode]
//   - [INSMatrix.AllowsEmptySelection]: A Boolean that indicates whether a radio-mode matrix supports an empty selection.
//   - [INSMatrix.SetAllowsEmptySelection]
//   - [INSMatrix.SelectionByRect]: A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor.
//   - [INSMatrix.SetSelectionByRect]
//
// # Managing the Cell Class
//
//   - [INSMatrix.Prototype]: The prototype cell that’s copied whenever the matrix creates a new cell.
//   - [INSMatrix.SetPrototype]
//
// # Laying Out the Cells of the Matrix
//
//   - [INSMatrix.AddColumn]: Adds a new column of cells to the right of the last column.
//   - [INSMatrix.AddColumnWithCells]: Adds a new column of cells to the right of the last column, using the given cells.
//   - [INSMatrix.AddRow]: Adds a new row of cells below the last row.
//   - [INSMatrix.AddRowWithCells]: Adds a new row of cells below the last row, using the specified cells.
//   - [INSMatrix.CellFrameAtRowColumn]: Returns the frame rectangle of the cell that would be drawn at the specified location.
//   - [INSMatrix.CellSize]: The size of each cell in the matrix.
//   - [INSMatrix.SetCellSize]
//   - [INSMatrix.GetNumberOfRowsColumns]: Obtains the number of rows and columns in the receiver.
//   - [INSMatrix.InsertColumn]: Inserts a new column of cells at the specified location.
//   - [INSMatrix.InsertColumnWithCells]: Inserts a new column of cells before the specified column, using the given cells.
//   - [INSMatrix.InsertRow]: Inserts a new row of cells before the specified row.
//   - [INSMatrix.InsertRowWithCells]: Inserts a new row of cells before the specified row, using the given cells.
//   - [INSMatrix.IntercellSpacing]: The vertical and horizontal spacing between cells in the matrix.
//   - [INSMatrix.SetIntercellSpacing]
//   - [INSMatrix.MakeCellAtRowColumn]: Creates a new cell at the location specified by the given row and column in the receiver.
//   - [INSMatrix.NumberOfColumns]: The number of columns in the matrix.
//   - [INSMatrix.NumberOfRows]: The number of rows in the matrix.
//   - [INSMatrix.PutCellAtRowColumn]: Replaces the cell at the specified row and column with the new cell.
//   - [INSMatrix.RemoveColumn]: Removes the specified column at from the receiver.
//   - [INSMatrix.RemoveRow]: Removes the specified row from the receiver.
//   - [INSMatrix.RenewRowsColumns]: Changes the number of rows and columns in the receiver.
//   - [INSMatrix.SortUsingFunctionContext]: Sorts the receiver’s cells in ascending order as defined by the specified comparison function.
//   - [INSMatrix.SortUsingSelector]: Sorts the receiver’s cells in ascending order as defined by the comparison method.
//
// # Auto Layout Sizing
//
//   - [INSMatrix.AutorecalculatesCellSize]: A Boolean that indicates whether the matrix auto-recalculates its cell size.
//   - [INSMatrix.SetAutorecalculatesCellSize]
//
// # Finding Matrix Coordinates
//
//   - [INSMatrix.GetRowColumnForPoint]: Indicates whether the specified point lies within one of the cells of the matrix and returns the location of the cell within which the point lies.
//   - [INSMatrix.GetRowColumnOfCell]: Searches the receiver for the specified cell and returns the row and column of the cell
//
// # Managing Attributes of Individual Cells
//
//   - [INSMatrix.SetStateAtRowColumn]: Sets the state of the cell at specified location.
//   - [INSMatrix.SetToolTipForCell]: Sets the tooltip for the cell.
//   - [INSMatrix.ToolTipForCell]: Returns the tooltip for the specified cell.
//
// # Selecting and Deselecting Cells
//
//   - [INSMatrix.SelectCellAtRowColumn]: Selects the cell at the specified row and column within the receiver.
//   - [INSMatrix.SelectCellWithTag]: Selects the last cell with the given tag.
//   - [INSMatrix.KeyCell]: The cell that will be clicked when the user presses the Space bar.
//   - [INSMatrix.SetKeyCell]
//   - [INSMatrix.SetSelectionFromToAnchorHighlight]: Programmatically selects a range of cells.
//   - [INSMatrix.DeselectAllCells]: Deselects all cells in the receiver and, if necessary, redisplays the receiver.
//   - [INSMatrix.DeselectSelectedCell]: Deselects the selected cell or cells.
//
// # Finding Cells
//
//   - [INSMatrix.SelectedCells]: An array containing all of the matrix’s highlighted cells plus its selected cell.
//   - [INSMatrix.SelectedColumn]: The column number of the selected cell.
//   - [INSMatrix.SelectedRow]: The row number of the selected cell.
//   - [INSMatrix.CellAtRowColumn]: Returns the cell at the specified row and column.
//   - [INSMatrix.CellWithTag]: Searches the receiver and returns the last cell matching the specified tag.
//   - [INSMatrix.Cells]: An array containing the cells of the matrix.
//
// # Modifying Graphics Attributes
//
//   - [INSMatrix.BackgroundColor]: The background color of the matrix (the space between the cells).
//   - [INSMatrix.SetBackgroundColor]
//   - [INSMatrix.CellBackgroundColor]: The background color of the matrix’s cells.
//   - [INSMatrix.SetCellBackgroundColor]
//   - [INSMatrix.DrawsBackground]: A Boolean that indicates whether the matrix draws its background.
//   - [INSMatrix.SetDrawsBackground]
//   - [INSMatrix.DrawsCellBackground]: A Boolean that indicates whether the matrix draws the background within each of its cells.
//   - [INSMatrix.SetDrawsCellBackground]
//
// # Editing Text in Cells
//
//   - [INSMatrix.SelectText]: Selects text in the currently selected cell or in the key cell.
//   - [INSMatrix.SelectTextAtRowColumn]: Selects the text in the cell at the specified location and returns the cell.
//   - [INSMatrix.TextShouldBeginEditing]: Requests permission to begin editing text.
//   - [INSMatrix.TextDidBeginEditing]: Invoked when there’s a change in the text after the receiver gains first responder status.
//   - [INSMatrix.TextDidChange]: Invoked when a key-down event or paste operation occurs that changes the receiver’s contents.
//   - [INSMatrix.TextShouldEndEditing]: Requests permission to end editing.
//   - [INSMatrix.TextDidEndEditing]: Invoked when text editing ends.
//
// # Setting Tab Key Behavior
//
//   - [INSMatrix.TabKeyTraversesCells]: A Boolean that indicates whether pressing the Tab key advances the key cell to the next selectable cell.
//   - [INSMatrix.SetTabKeyTraversesCells]
//
// # Managing the Delegate
//
//   - [INSMatrix.Delegate]: The delegate for messages from the field editor.
//   - [INSMatrix.SetDelegate]
//
// # Resizing the Matrix and Its Cells
//
//   - [INSMatrix.AutosizesCells]: A Boolean that indicates whether the cell sizes change when the receiver is resized.
//   - [INSMatrix.SetAutosizesCells]
//   - [INSMatrix.SetValidateSize]: Specifies whether the receiver’s size information is validated.
//   - [INSMatrix.SizeToCells]: Changes the width and the height of the receiver’s frame so it exactly contains the cells.
//
// # Scrolling Cells in the Matrix
//
//   - [INSMatrix.SetScrollable]: Specifies whether the cells in the matrix are scrollable.
//   - [INSMatrix.ScrollCellToVisibleAtRowColumn]: Scrolls the receiver so the specified cell is visible.
//
// # Displaying and Highlighting Cells
//
//   - [INSMatrix.DrawCellAtRowColumn]: Displays the cell at the specified row and column.
//   - [INSMatrix.HighlightCellAtRowColumn]: Highlights or unhighlights the cell at the specified row and column location.
//
// # Managing and Sending Action Messages
//
//   - [INSMatrix.SendAction]: If the selected cell has both an action and a target, sends its action to its target.
//   - [INSMatrix.SendActionToForAllCells]: Iterates through the cells in the receiver, sending the specified selector to an object for each cell.
//   - [INSMatrix.DoubleAction]: The action sent to the target of the receiver when the user double-clicks a cell.
//   - [INSMatrix.SetDoubleAction]
//   - [INSMatrix.SendDoubleAction]: Sends the double-click action message to the target of the receiver.
//
// # Handling Event and Action Messages
//
//   - [INSMatrix.MouseDownFlags]: The flags in effect at the mouse-down event that started the current tracking session.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix
type INSMatrix interface {
	INSControl
	NSUserInterfaceValidations
	NSViewToolTipOwner

	// Topic: Initializing an NSMatrix Object

	// Initializes and returns a newly allocated matrix of the specified size using cells of the given class.
	InitWithFrameModeCellClassNumberOfRowsNumberOfColumns(frameRect corefoundation.CGRect, mode NSMatrixMode, factoryId objc.Class, rowsHigh int, colsWide int) NSMatrix
	// Initializes and returns a newly allocated matrix of the specified size using the given cell as a prototype.
	InitWithFrameModePrototypeNumberOfRowsNumberOfColumns(frameRect corefoundation.CGRect, mode NSMatrixMode, cell INSCell, rowsHigh int, colsWide int) NSMatrix

	// Topic: Configuring the Matrix Object

	// The selection mode of the receiver.
	Mode() NSMatrixMode
	SetMode(value NSMatrixMode)
	// A Boolean that indicates whether a radio-mode matrix supports an empty selection.
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	// A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor.
	SelectionByRect() bool
	SetSelectionByRect(value bool)

	// Topic: Managing the Cell Class

	// The prototype cell that’s copied whenever the matrix creates a new cell.
	Prototype() INSCell
	SetPrototype(value INSCell)

	// Topic: Laying Out the Cells of the Matrix

	// Adds a new column of cells to the right of the last column.
	AddColumn()
	// Adds a new column of cells to the right of the last column, using the given cells.
	AddColumnWithCells(newCells []NSCell)
	// Adds a new row of cells below the last row.
	AddRow()
	// Adds a new row of cells below the last row, using the specified cells.
	AddRowWithCells(newCells []NSCell)
	// Returns the frame rectangle of the cell that would be drawn at the specified location.
	CellFrameAtRowColumn(row int, col int) corefoundation.CGRect
	// The size of each cell in the matrix.
	CellSize() corefoundation.CGSize
	SetCellSize(value corefoundation.CGSize)
	// Obtains the number of rows and columns in the receiver.
	GetNumberOfRowsColumns(rowCount unsafe.Pointer, colCount unsafe.Pointer)
	// Inserts a new column of cells at the specified location.
	InsertColumn(column int)
	// Inserts a new column of cells before the specified column, using the given cells.
	InsertColumnWithCells(column int, newCells []NSCell)
	// Inserts a new row of cells before the specified row.
	InsertRow(row int)
	// Inserts a new row of cells before the specified row, using the given cells.
	InsertRowWithCells(row int, newCells []NSCell)
	// The vertical and horizontal spacing between cells in the matrix.
	IntercellSpacing() corefoundation.CGSize
	SetIntercellSpacing(value corefoundation.CGSize)
	// Creates a new cell at the location specified by the given row and column in the receiver.
	MakeCellAtRowColumn(row int, col int) INSCell
	// The number of columns in the matrix.
	NumberOfColumns() int
	// The number of rows in the matrix.
	NumberOfRows() int
	// Replaces the cell at the specified row and column with the new cell.
	PutCellAtRowColumn(newCell INSCell, row int, col int)
	// Removes the specified column at from the receiver.
	RemoveColumn(col int)
	// Removes the specified row from the receiver.
	RemoveRow(row int)
	// Changes the number of rows and columns in the receiver.
	RenewRowsColumns(newRows int, newCols int)
	// Sorts the receiver’s cells in ascending order as defined by the specified comparison function.
	SortUsingFunctionContext(compare objectivec.IObject, context unsafe.Pointer)
	// Sorts the receiver’s cells in ascending order as defined by the comparison method.
	SortUsingSelector(comparator objc.SEL)

	// Topic: Auto Layout Sizing

	// A Boolean that indicates whether the matrix auto-recalculates its cell size.
	AutorecalculatesCellSize() bool
	SetAutorecalculatesCellSize(value bool)

	// Topic: Finding Matrix Coordinates

	// Indicates whether the specified point lies within one of the cells of the matrix and returns the location of the cell within which the point lies.
	GetRowColumnForPoint(point corefoundation.CGPoint) (int, int, bool)
	// Searches the receiver for the specified cell and returns the row and column of the cell
	GetRowColumnOfCell(cell NSCell) (int, int, bool)

	// Topic: Managing Attributes of Individual Cells

	// Sets the state of the cell at specified location.
	SetStateAtRowColumn(value int, row int, col int)
	// Sets the tooltip for the cell.
	SetToolTipForCell(toolTipString string, cell INSCell)
	// Returns the tooltip for the specified cell.
	ToolTipForCell(cell INSCell) string

	// Topic: Selecting and Deselecting Cells

	// Selects the cell at the specified row and column within the receiver.
	SelectCellAtRowColumn(row int, col int)
	// Selects the last cell with the given tag.
	SelectCellWithTag(tag int) bool
	// The cell that will be clicked when the user presses the Space bar.
	KeyCell() INSCell
	SetKeyCell(value INSCell)
	// Programmatically selects a range of cells.
	SetSelectionFromToAnchorHighlight(startPos int, endPos int, anchorPos int, lit bool)
	// Deselects all cells in the receiver and, if necessary, redisplays the receiver.
	DeselectAllCells()
	// Deselects the selected cell or cells.
	DeselectSelectedCell()

	// Topic: Finding Cells

	// An array containing all of the matrix’s highlighted cells plus its selected cell.
	SelectedCells() []NSCell
	// The column number of the selected cell.
	SelectedColumn() int
	// The row number of the selected cell.
	SelectedRow() int
	// Returns the cell at the specified row and column.
	CellAtRowColumn(row int, col int) INSCell
	// Searches the receiver and returns the last cell matching the specified tag.
	CellWithTag(tag int) INSCell
	// An array containing the cells of the matrix.
	Cells() []NSCell

	// Topic: Modifying Graphics Attributes

	// The background color of the matrix (the space between the cells).
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// The background color of the matrix’s cells.
	CellBackgroundColor() INSColor
	SetCellBackgroundColor(value INSColor)
	// A Boolean that indicates whether the matrix draws its background.
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	// A Boolean that indicates whether the matrix draws the background within each of its cells.
	DrawsCellBackground() bool
	SetDrawsCellBackground(value bool)

	// Topic: Editing Text in Cells

	// Selects text in the currently selected cell or in the key cell.
	SelectText(sender objectivec.IObject)
	// Selects the text in the cell at the specified location and returns the cell.
	SelectTextAtRowColumn(row int, col int) INSCell
	// Requests permission to begin editing text.
	TextShouldBeginEditing(textObject INSText) bool
	// Invoked when there’s a change in the text after the receiver gains first responder status.
	TextDidBeginEditing(notification foundation.NSNotification)
	// Invoked when a key-down event or paste operation occurs that changes the receiver’s contents.
	TextDidChange(notification foundation.NSNotification)
	// Requests permission to end editing.
	TextShouldEndEditing(textObject INSText) bool
	// Invoked when text editing ends.
	TextDidEndEditing(notification foundation.NSNotification)

	// Topic: Setting Tab Key Behavior

	// A Boolean that indicates whether pressing the Tab key advances the key cell to the next selectable cell.
	TabKeyTraversesCells() bool
	SetTabKeyTraversesCells(value bool)

	// Topic: Managing the Delegate

	// The delegate for messages from the field editor.
	Delegate() NSMatrixDelegate
	SetDelegate(value NSMatrixDelegate)

	// Topic: Resizing the Matrix and Its Cells

	// A Boolean that indicates whether the cell sizes change when the receiver is resized.
	AutosizesCells() bool
	SetAutosizesCells(value bool)
	// Specifies whether the receiver’s size information is validated.
	SetValidateSize(flag bool)
	// Changes the width and the height of the receiver’s frame so it exactly contains the cells.
	SizeToCells()

	// Topic: Scrolling Cells in the Matrix

	// Specifies whether the cells in the matrix are scrollable.
	SetScrollable(flag bool)
	// Scrolls the receiver so the specified cell is visible.
	ScrollCellToVisibleAtRowColumn(row int, col int)

	// Topic: Displaying and Highlighting Cells

	// Displays the cell at the specified row and column.
	DrawCellAtRowColumn(row int, col int)
	// Highlights or unhighlights the cell at the specified row and column location.
	HighlightCellAtRowColumn(flag bool, row int, col int)

	// Topic: Managing and Sending Action Messages

	// If the selected cell has both an action and a target, sends its action to its target.
	SendAction() bool
	// Iterates through the cells in the receiver, sending the specified selector to an object for each cell.
	SendActionToForAllCells(selector objc.SEL, object objectivec.IObject, flag bool)
	// The action sent to the target of the receiver when the user double-clicks a cell.
	DoubleAction() objc.SEL
	SetDoubleAction(value objc.SEL)
	// Sends the double-click action message to the target of the receiver.
	SendDoubleAction()

	// Topic: Handling Event and Action Messages

	// The flags in effect at the mouse-down event that started the current tracking session.
	MouseDownFlags() int
}

// Init initializes the instance.
func (m NSMatrix) Init() NSMatrix {
	rv := objc.Send[NSMatrix](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMatrix) Autorelease() NSMatrix {
	rv := objc.Send[NSMatrix](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMatrix creates a new NSMatrix instance.
func NewNSMatrix() NSMatrix {
	class := getNSMatrixClass()
	rv := objc.Send[NSMatrix](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewMatrixWithCoder(coder foundation.INSCoder) NSMatrix {
	instance := getNSMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMatrixFromID(rv)
}

// Initializes a newly allocated matrix with the specified frame.
//
// frameRect: The frame with which to initialize the matrix.
//
// # Return Value
// 
// The [NSMatrix], initialized with default parameters. The new [NSMatrix]
// contains no rows or columns. The default mode is [NSRadioModeMatrix]. The
// default cell class is [NSActionCell].
//
// # Discussion
// 
// .
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/init(frame:)
func NewMatrixWithFrame(frameRect corefoundation.CGRect) NSMatrix {
	instance := getNSMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSMatrixFromID(rv)
}

// Initializes and returns a newly allocated matrix of the specified size
// using cells of the given class.
//
// frameRect: The matrix’s frame.
//
// mode: The tracking mode for the matrix; this can be one of the modes described in
// [NSMatrix.Mode].
// //
// [NSMatrix.Mode]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum
//
// factoryId: The class to use for any cells that the matrix creates and uses. This can
// be obtained by sending a `class` message to the desired subclass of
// [NSCell].
//
// rowsHigh: The number of rows in the matrix.
//
// colsWide: The number of columns in the matrix.
//
// # Return Value
// 
// The initialized instance of [NSMatrix].
//
// # Discussion
// 
// This method is the designated initializer for matrices that add cells by
// creating instances of an [NSCell] subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/init(frame:mode:cellClass:numberOfRows:numberOfColumns:)
func NewMatrixWithFrameModeCellClassNumberOfRowsNumberOfColumns(frameRect corefoundation.CGRect, mode NSMatrixMode, factoryId objc.Class, rowsHigh int, colsWide int) NSMatrix {
	instance := getNSMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:mode:cellClass:numberOfRows:numberOfColumns:"), frameRect, mode, factoryId, rowsHigh, colsWide)
	return NSMatrixFromID(rv)
}

// Initializes and returns a newly allocated matrix of the specified size
// using the given cell as a prototype.
//
// frameRect: The matrix’s frame.
//
// mode: The tracking mode for the matrix; this can be one of the modes described in
// [NSMatrix.Mode].
// //
// [NSMatrix.Mode]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum
//
// cell: An instance of a subclass of [NSCell], which the new matrix copies when it
// creates new cells.
//
// rowsHigh: The number of rows in the matrix.
//
// colsWide: The number of columns in the matrix.
//
// # Discussion
// 
// This method is the designated initializer for matrices that add cells by
// copying an instance of an [NSCell] subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/init(frame:mode:prototype:numberOfRows:numberOfColumns:)
func NewMatrixWithFrameModePrototypeNumberOfRowsNumberOfColumns(frameRect corefoundation.CGRect, mode NSMatrixMode, cell INSCell, rowsHigh int, colsWide int) NSMatrix {
	instance := getNSMatrixClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:mode:prototype:numberOfRows:numberOfColumns:"), frameRect, mode, cell, rowsHigh, colsWide)
	return NSMatrixFromID(rv)
}

// Initializes and returns a newly allocated matrix of the specified size
// using cells of the given class.
//
// frameRect: The matrix’s frame.
//
// mode: The tracking mode for the matrix; this can be one of the modes described in
// [NSMatrix.Mode].
// //
// [NSMatrix.Mode]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum
//
// factoryId: The class to use for any cells that the matrix creates and uses. This can
// be obtained by sending a `class` message to the desired subclass of
// [NSCell].
//
// rowsHigh: The number of rows in the matrix.
//
// colsWide: The number of columns in the matrix.
//
// # Return Value
// 
// The initialized instance of [NSMatrix].
//
// # Discussion
// 
// This method is the designated initializer for matrices that add cells by
// creating instances of an [NSCell] subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/init(frame:mode:cellClass:numberOfRows:numberOfColumns:)
func (m NSMatrix) InitWithFrameModeCellClassNumberOfRowsNumberOfColumns(frameRect corefoundation.CGRect, mode NSMatrixMode, factoryId objc.Class, rowsHigh int, colsWide int) NSMatrix {
	rv := objc.Send[NSMatrix](m.ID, objc.Sel("initWithFrame:mode:cellClass:numberOfRows:numberOfColumns:"), frameRect, mode, factoryId, rowsHigh, colsWide)
	return rv
}

// Initializes and returns a newly allocated matrix of the specified size
// using the given cell as a prototype.
//
// frameRect: The matrix’s frame.
//
// mode: The tracking mode for the matrix; this can be one of the modes described in
// [NSMatrix.Mode].
// //
// [NSMatrix.Mode]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum
//
// cell: An instance of a subclass of [NSCell], which the new matrix copies when it
// creates new cells.
//
// rowsHigh: The number of rows in the matrix.
//
// colsWide: The number of columns in the matrix.
//
// # Discussion
// 
// This method is the designated initializer for matrices that add cells by
// copying an instance of an [NSCell] subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/init(frame:mode:prototype:numberOfRows:numberOfColumns:)
func (m NSMatrix) InitWithFrameModePrototypeNumberOfRowsNumberOfColumns(frameRect corefoundation.CGRect, mode NSMatrixMode, cell INSCell, rowsHigh int, colsWide int) NSMatrix {
	rv := objc.Send[NSMatrix](m.ID, objc.Sel("initWithFrame:mode:prototype:numberOfRows:numberOfColumns:"), frameRect, mode, cell, rowsHigh, colsWide)
	return rv
}

// Adds a new column of cells to the right of the last column.
//
// # Discussion
// 
// This method raises an [NSRangeException] if there are 0 rows or 0 columns.
// This method creates new cells as needed with [CellAtRowColumn]. Use
// [RenewRowsColumns] to add new cells to an empty matrix.
// 
// If the number of rows or columns in the receiver has been changed with
// [RenewRowsColumns], new cells are created only if they are needed. This
// fact allows you to grow and shrink an [NSMatrix] without repeatedly
// creating and freeing the cells.
// 
// This method redraws the receiver. Your code may need to send [SizeToCells]
// after sending this method to resize the receiver to fit the newly added
// cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/addColumn()
func (m NSMatrix) AddColumn() {
	objc.Send[objc.ID](m.ID, objc.Sel("addColumn"))
}

// Adds a new column of cells to the right of the last column, using the given
// cells.
//
// newCells: An array of objects to use when filling the new column starting with the
// object at index 0. Each object in should be an instance of [NSCell] or one
// of its subclasses (usually [NSActionCell]). The array should have a
// sufficient number of cells to fill the entire column. Extra cells are
// ignored, unless the matrix is empty. In that case, a matrix is created with
// one column and enough rows for all the elements of `newCells`.
//
// # Discussion
// 
// This method redraws the receiver. Your code may need to send [SizeToCells]
// after sending this method to resize the receiver to fit the newly added
// cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/addColumn(with:)
func (m NSMatrix) AddColumnWithCells(newCells []NSCell) {
	objc.Send[objc.ID](m.ID, objc.Sel("addColumnWithCells:"), objectivec.IObjectSliceToNSArray(newCells))
}

// Adds a new row of cells below the last row.
//
// # Discussion
// 
// New cells are created as needed with [CellAtRowColumn]. This method raises
// an [NSRangeException] if there are 0 rows or 0 columns. Use
// [RenewRowsColumns] to add new cells to an empty matrix.
// 
// If the number of rows or columns in the receiver has been changed with
// [RenewRowsColumns], then new cells are created only if they are needed.
// This fact allows you to grow and shrink an NSMatrix without repeatedly
// creating and freeing the cells.
// 
// This method redraws the receiver. Your code may need to send [SizeToCells]
// after sending this method to resize the receiver to fit the newly added
// cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/addRow()
func (m NSMatrix) AddRow() {
	objc.Send[objc.ID](m.ID, objc.Sel("addRow"))
}

// Adds a new row of cells below the last row, using the specified cells.
//
// newCells: An array of objects to use to fill the new row, starting with the object at
// index 0. Each object should be an instance of [NSCell] or one of its
// subclasses (usually [NSActionCell]). The array should contain a sufficient
// number of cells to fill the entire row. Extra cells are ignored, unless the
// matrix is empty. In that case, a matrix is created with one row and enough
// columns for all the elements of `newCells`.
//
// # Discussion
// 
// This method redraws the receiver. Your code may need to send [SizeToCells]
// after sending this method to resize the receiver to fit the newly added
// cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/addRow(with:)
func (m NSMatrix) AddRowWithCells(newCells []NSCell) {
	objc.Send[objc.ID](m.ID, objc.Sel("addRowWithCells:"), objectivec.IObjectSliceToNSArray(newCells))
}

// Returns the frame rectangle of the cell that would be drawn at the
// specified location.
//
// row: The row of the cell.
//
// col: The column of the cell.
//
// # Return Value
// 
// The frame rectangle of the cell (whether or not the specified cell actually
// exists).
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/cellFrame(atRow:column:)
func (m NSMatrix) CellFrameAtRowColumn(row int, col int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m.ID, objc.Sel("cellFrameAtRow:column:"), row, col)
	return corefoundation.CGRect(rv)
}

// Obtains the number of rows and columns in the receiver.
//
// rowCount: On return, the number of rows in the matrix.
//
// colCount: On return, the number of columns in the matrix.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/getNumberOfRows(_:columns:)
func (m NSMatrix) GetNumberOfRowsColumns(rowCount unsafe.Pointer, colCount unsafe.Pointer) {
	objc.Send[objc.ID](m.ID, objc.Sel("getNumberOfRows:columns:"), rowCount, colCount)
}

// Inserts a new column of cells at the specified location.
//
// column: The number of the column before which the new column is inserted. If
// `column` is greater than the number of columns in the receiver, enough
// columns are created to expand the receiver to be `column` columns wide.
//
// # Discussion
// 
// New cells are created if needed with [CellAtRowColumn]. This method redraws
// the receiver. Your code may need to send [SizeToCells] after sending this
// method to resize the receiver to fit the newly added cells.
// 
// If the number of rows or columns in the receiver has been changed with
// [RenewRowsColumns], new cells are created only if they’re needed. This
// fact allows you to grow and shrink an [NSMatrix] without repeatedly
// creating and freeing the cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/insertColumn(_:)
func (m NSMatrix) InsertColumn(column int) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertColumn:"), column)
}

// Inserts a new column of cells before the specified column, using the given
// cells.
//
// column: The column at which to insert the new cells.
//
// newCells: An array of objects to use to fill the new column, starting with the object
// at index 0. Each object should be an instance of [NSCell] or one of its
// subclasses (usually [NSActionCell]).
//
// # Discussion
// 
// If `column` is greater than the number of columns in the receiver, enough
// columns are created to expand the receiver to be `column` columns wide.
// `newCells` should either be empty or contain a sufficient number of cells
// to fill each new column. If `newCells` is `nil` or an array with no
// elements, the call is equivalent to calling [InsertColumn]. Extra cells are
// ignored, unless the matrix is empty. In that case, a matrix is created with
// one column and enough rows for all the elements of `newCells`.
// 
// This method redraws the receiver. Your code may need to send [SizeToCells]
// after sending this method to resize the receiver to fit the newly added
// cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/insertColumn(_:with:)
func (m NSMatrix) InsertColumnWithCells(column int, newCells []NSCell) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertColumn:withCells:"), column, objectivec.IObjectSliceToNSArray(newCells))
}

// Inserts a new row of cells before the specified row.
//
// row: The location at which to insert the new row. If this is greater than the
// number of rows in the receiver, enough rows are created to expand the
// receiver to be `row` rows high.
//
// # Discussion
// 
// New cells are created if needed with [CellAtRowColumn]. This method redraws
// the receiver. Your code may need to send [SizeToCells] after sending this
// method to resize the receiver to fit the newly added cells.
// 
// If the number of rows or columns in the receiver has been changed with
// [RenewRowsColumns], then new cells are created only if they’re needed.
// This fact allows you to grow and shrink an [NSMatrix] without repeatedly
// creating and freeing the cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/insertRow(_:)
func (m NSMatrix) InsertRow(row int) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertRow:"), row)
}

// Inserts a new row of cells before the specified row, using the given cells.
//
// row: The location at which to insert the new row.
//
// newCells: An array of objects to use when filling the new row, starting with the
// object at index 0. Each object in `newCells` should be an instance of
// [NSCell] or one of its subclasses (usually [NSActionCell]).
//
// # Discussion
// 
// If `row` is greater than the number of rows in the receiver, enough rows
// are created to expand the receiver to be `row` rows high. `newCells` should
// either be empty or contain a sufficient number of cells to fill each new
// row. If `newCells` is `nil` or an array with no elements, the call is
// equivalent to calling [InsertRow]. Extra cells are ignored, unless the
// matrix is empty. In that case, a matrix is created with one row and enough
// columns for all the elements of `newCells`.
// 
// This method redraws the receiver. Your code may need to send [SizeToCells]
// after sending this method to resize the receiver to fit the newly added
// cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/insertRow(_:with:)
func (m NSMatrix) InsertRowWithCells(row int, newCells []NSCell) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertRow:withCells:"), row, objectivec.IObjectSliceToNSArray(newCells))
}

// Creates a new cell at the location specified by the given row and column in
// the receiver.
//
// row: The row in which to create the new cell.
//
// col: The column in which to create the new cell.
//
// # Return Value
// 
// The newly created cell.
//
// # Discussion
// 
// If the receiver has a prototype cell, it’s copied to create the new cell.
// If not, and if the receiver has a cell class set, it allocates and
// initializes (with `init`) an instance of that class. If the receiver
// hasn’t had either a prototype cell or a cell class set, [NSMatrix]
// creates an [NSActionCell].
// 
// Your code should never invoke this method directly; it’s used by [AddRow]
// and other methods when a cell must be created. It may be overridden to
// provide more specific initialization of cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/makeCell(atRow:column:)
func (m NSMatrix) MakeCellAtRowColumn(row int, col int) INSCell {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("makeCellAtRow:column:"), row, col)
	return NSCellFromID(rv)
}

// Replaces the cell at the specified row and column with the new cell.
//
// newCell: The cell to insert into the matrix.
//
// row: The row in which to place the new cell.
//
// col: The column in which to place the new cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/putCell(_:atRow:column:)
func (m NSMatrix) PutCellAtRowColumn(newCell INSCell, row int, col int) {
	objc.Send[objc.ID](m.ID, objc.Sel("putCell:atRow:column:"), newCell, row, col)
}

// Removes the specified column at from the receiver.
//
// col: The column to remove.
//
// # Discussion
// 
// The column’s cells are autoreleased. This method redraws the receiver.
// Your code should normally send [SizeToCells] after invoking this method to
// resize the receiver so it fits the reduced cell count.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/removeColumn(_:)
func (m NSMatrix) RemoveColumn(col int) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeColumn:"), col)
}

// Removes the specified row from the receiver.
//
// row: The row to remove.
//
// # Discussion
// 
// The row’s cells are autoreleased. This method redraws the receiver. Your
// code should normally send [SizeToCells] after invoking this method to
// resize the receiver so it fits the reduced cell count.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/removeRow(_:)
func (m NSMatrix) RemoveRow(row int) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeRow:"), row)
}

// Changes the number of rows and columns in the receiver.
//
// newRows: The new number of rows in the matrix.
//
// newCols: The new number of columns in the matrix.
//
// # Discussion
// 
// This method uses the same cells as before, creating new cells only if the
// new size is larger; it never frees cells. It doesn’t redisplay the
// receiver. Your code should normally send [SizeToCells] after invoking this
// method to resize the receiver so it fits the changed cell arrangement. This
// method deselects all cells in the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/renewRows(_:columns:)
func (m NSMatrix) RenewRowsColumns(newRows int, newCols int) {
	objc.Send[objc.ID](m.ID, objc.Sel("renewRows:columns:"), newRows, newCols)
}

// Sorts the receiver’s cells in ascending order as defined by the specified
// comparison function.
//
// compare: The function to use when comparing cells. The comparison function is used
// to compare two elements at a time and should return [NSOrderedAscending] if
// the first element is smaller than the second, [NSOrderedDescending] if the
// first element is larger than the second, and [NSOrderedSame] if the
// elements are equal.
//
// context: Context passed to the comparison function as its third argument, each time
// its called. This allows the comparison to be based on some outside
// parameter, such as whether character sorting is case-sensitive or
// case-insensitive.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/sort(using:context:)
func (m NSMatrix) SortUsingFunctionContext(compare objectivec.IObject, context unsafe.Pointer) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortUsingFunction:context:"), compare, context)
}

// Sorts the receiver’s cells in ascending order as defined by the
// comparison method.
//
// comparator: The comparison method.
//
// # Discussion
// 
// The comparator message is sent to each object in the matrix and has as its
// single argument another object in the array. The comparison method is used
// to compare two elements at a time and should return [NSOrderedAscending] if
// the receiver is smaller than the argument, [NSOrderedDescending] if the
// receiver is larger than the argument, and [NSOrderedSame] if they are
// equal.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/sort(using:)
func (m NSMatrix) SortUsingSelector(comparator objc.SEL) {
	objc.Send[objc.ID](m.ID, objc.Sel("sortUsingSelector:"), comparator)
}

// Indicates whether the specified point lies within one of the cells of the
// matrix and returns the location of the cell within which the point lies.
//
// row: On return, the row of the cell containing the specified point.
//
// col: On return, the column of the cell containing the specified point.
//
// point: The point to locate; this point should be in the coordinate system of the
// receiver.
//
// # Return Value
// 
// [true] if the point lies within one of the cells in the receiver; [false]
// if the point falls outside the bounds of the receiver or lies within an
// intercell spacing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/getRow(_:column:for:)
func (m NSMatrix) GetRowColumnForPoint(point corefoundation.CGPoint) (int, int, bool) {
	var row int
	var col int
	rv := objc.Send[bool](m.ID, objc.Sel("getRow:column:forPoint:"), unsafe.Pointer(&row), unsafe.Pointer(&col), point)
	return row, col, rv
}

// Searches the receiver for the specified cell and returns the row and column
// of the cell
//
// row: On return, the row in which the cell is located.
//
// col: On return, the column in which the cell is located.
//
// cell: The cell to locate within the matrix.
//
// # Return Value
// 
// [true] if the cell is one of the cells in the receiver, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// .
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/getRow(_:column:of:)
func (m NSMatrix) GetRowColumnOfCell(cell NSCell) (int, int, bool) {
	var row int
	var col int
	rv := objc.Send[bool](m.ID, objc.Sel("getRow:column:ofCell:"), unsafe.Pointer(&row), unsafe.Pointer(&col), cell)
	return row, col, rv
}

// Sets the state of the cell at specified location.
//
// value: The value to assign to the cell.
//
// row: The row in which the cell is located.
//
// col: The column in which the cell is located.
//
// # Discussion
// 
// For radio-mode matrices, if `value` is nonzero the specified cell is
// selected before its state is set to `value`. If `value` is 0 and the
// receiver is a radio-mode matrix, the currently selected cell is deselected
// (providing that empty selection is allowed).
// 
// This method redraws the affected cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/setState(_:atRow:column:)
func (m NSMatrix) SetStateAtRowColumn(value int, row int, col int) {
	objc.Send[objc.ID](m.ID, objc.Sel("setState:atRow:column:"), value, row, col)
}

// Sets the tooltip for the cell.
//
// toolTipString: The string to use as the cell’s tooltip (or help tag).
//
// cell: The cell to which to assign the tooltip.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/setToolTip(_:for:)
func (m NSMatrix) SetToolTipForCell(toolTipString string, cell INSCell) {
	objc.Send[objc.ID](m.ID, objc.Sel("setToolTip:forCell:"), objc.String(toolTipString), cell)
}

// Returns the tooltip for the specified cell.
//
// cell: The cell for which to return the tooltip.
//
// # Return Value
// 
// The string used as the cell’s tooltip.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/toolTip(for:)
func (m NSMatrix) ToolTipForCell(cell INSCell) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("toolTipForCell:"), cell)
	return foundation.NSStringFromID(rv).String()
}

// Selects the cell at the specified row and column within the receiver.
//
// row: The row of the cell to select.
//
// col: The column of the cell to select.
//
// # Discussion
// 
// If the specified cell is an editable text cell, its text is selected. If
// either `row` or `column` is –1, then the current selection is cleared
// (unless the receiver is an [NSRadioModeMatrix] and doesn’t allow empty
// selection). This method redraws the affected cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/selectCell(atRow:column:)
func (m NSMatrix) SelectCellAtRowColumn(row int, col int) {
	objc.Send[objc.ID](m.ID, objc.Sel("selectCellAtRow:column:"), row, col)
}

// Selects the last cell with the given tag.
//
// tag: The tag of the cell to select.
//
// # Return Value
// 
// [true] if the receiver contains a cell whose tag matches `anInt`, or
// [false] if no such cell exists
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the matrix has at least one cell whose tag is equal to `anInt`, the last
// cell (when viewing the matrix as a row-ordered array) is selected. If the
// specified cell is an editable text cell, its text is selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/selectCell(withTag:)
func (m NSMatrix) SelectCellWithTag(tag int) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("selectCellWithTag:"), tag)
	return rv
}

// Programmatically selects a range of cells.
//
// startPos: The position of the cell that marks where the user would have pressed the
// mouse button.
//
// endPos: The position of the cell that marks where the user would have released the
// mouse button.
//
// anchorPos: The position of the cell to treat as the last cell the user would have
// selected. To simulate Shift-dragging (continuous selection) `anchorPos`
// should be the `endPos` used in the last method call. To simulate
// Command-dragging (discontinuous selection), `anchorPos` should be the same
// as this method call’s `startPos`.
//
// lit: [true] if cells selected by this method should be highlighted.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// `startPos`, `endPos`, and `anchorPos` are cell positions, counting from 0
// at the upper left cell of the receiver, in row order. For example, the
// third cell in the top row would be number 2.
// 
// To simulate dragging without a modifier key, deselecting anything that was
// selected before, call [DeselectAllCells] before calling this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/setSelectionFrom(_:to:anchor:highlight:)
func (m NSMatrix) SetSelectionFromToAnchorHighlight(startPos int, endPos int, anchorPos int, lit bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setSelectionFrom:to:anchor:highlight:"), startPos, endPos, anchorPos, lit)
}

// Deselects all cells in the receiver and, if necessary, redisplays the
// receiver.
//
// # Discussion
// 
// If the selection mode is [NSRadioModeMatrix] and empty selection is not
// allowed, this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/deselectAllCells()
func (m NSMatrix) DeselectAllCells() {
	objc.Send[objc.ID](m.ID, objc.Sel("deselectAllCells"))
}

// Deselects the selected cell or cells.
//
// # Discussion
// 
// If the selection mode is [NSRadioModeMatrix] and empty selection is not
// allowed, or if nothing is currently selected, this method does nothing.
// This method doesn’t redisplay the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/deselectSelectedCell()
func (m NSMatrix) DeselectSelectedCell() {
	objc.Send[objc.ID](m.ID, objc.Sel("deselectSelectedCell"))
}

// Returns the cell at the specified row and column.
//
// row: The number of the row containing the cell to return.
//
// col: The number of the column containing the cell to return.
//
// # Return Value
// 
// The [NSCell] object at the specified row and column location specified, or
// `nil` if either `row` or `column` is outside the bounds of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/cell(atRow:column:)
func (m NSMatrix) CellAtRowColumn(row int, col int) INSCell {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("cellAtRow:column:"), row, col)
	return NSCellFromID(rv)
}

// Searches the receiver and returns the last cell matching the specified tag.
//
// tag: The tag of the cell to return.
//
// # Return Value
// 
// The last (when viewing the matrix as a row-ordered array) [NSCell] object
// that has a tag matching `anInt`, or `nil` if no such cell exists
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/cell(withTag:)
func (m NSMatrix) CellWithTag(tag int) INSCell {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("cellWithTag:"), tag)
	return NSCellFromID(rv)
}

// Selects text in the currently selected cell or in the key cell.
//
// # Discussion
// 
// If the currently selected cell is editable and enabled, its text is
// selected. Otherwise, the key cell is selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/selectText(_:)
func (m NSMatrix) SelectText(sender objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("selectText:"), sender)
}

// Selects the text in the cell at the specified location and returns the
// cell.
//
// row: The row containing the text to select.
//
// col: The column containing the text to select.
//
// # Return Value
// 
// If it is both editable and selectable, the cell at the specified row and
// column. If the cell at the specified location, is either not editable or
// not selectable, this method does nothing and returns nil. If `row` and
// `column` indicate a cell that is outside the receiver, this method does
// nothing and returns the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/selectText(atRow:column:)
func (m NSMatrix) SelectTextAtRowColumn(row int, col int) INSCell {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("selectTextAtRow:column:"), row, col)
	return NSCellFromID(rv)
}

// Requests permission to begin editing text.
//
// textObject: The text object requesting permission to begin editing.
//
// # Return Value
// 
// [true] if the text object should proceed to make changes. If the delegate
// returns [false], the text object abandons the editing operation.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The default behavior of this method is to return the value obtained from
// [ControlTextShouldBeginEditing], unless the delegate doesn’t respond to
// that method, in which case it returns [true], thereby allowing text editing
// to proceed.
// 
// This method is invoked to let the [NSTextField] respond to impending
// changes to its text. This method’s default behavior is to send
// [ControlTextShouldBeginEditing] to the receiver’s delegate (passing the
// receiver and `textObject` as parameters).
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/textShouldBeginEditing(_:)
func (m NSMatrix) TextShouldBeginEditing(textObject INSText) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("textShouldBeginEditing:"), textObject)
	return rv
}

// Invoked when there’s a change in the text after the receiver gains first
// responder status.
//
// notification: The [textDidBeginEditingNotification] notification.
// //
// [textDidBeginEditingNotification]: https://developer.apple.com/documentation/AppKit/NSControl/textDidBeginEditingNotification
//
// # Discussion
// 
// This method’s default behavior is to post an
// [textDidBeginEditingNotification] along with the receiving object to the
// default notification center. The posted notification’s user info contains
// the contents of notification’s user info dictionary, plus an additional
// key-value pair. The additional key is “[NSFieldEditor]”; the value for
// this key is the text object that began editing.
//
// [textDidBeginEditingNotification]: https://developer.apple.com/documentation/AppKit/NSControl/textDidBeginEditingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/textDidBeginEditing(_:)
func (m NSMatrix) TextDidBeginEditing(notification foundation.NSNotification) {
	objc.Send[objc.ID](m.ID, objc.Sel("textDidBeginEditing:"), notification)
}

// Invoked when a key-down event or paste operation occurs that changes the
// receiver’s contents.
//
// notification: The [textDidChangeNotification] notification.
// //
// [textDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSControl/textDidChangeNotification
//
// # Discussion
// 
// This method’s default behavior is to pass this message on to the selected
// cell (if the selected cell responds to ``) and then to post an
// [textDidChangeNotification] along with the receiving object to the default
// notification center. The posted notification’s user info contains the
// contents of notification’s user info dictionary, plus an additional
// key-value pair. The additional key is “[NSFieldEditor]”; the value for
// this key is the text object that changed.
//
// [textDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSControl/textDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/textDidChange(_:)
func (m NSMatrix) TextDidChange(notification foundation.NSNotification) {
	objc.Send[objc.ID](m.ID, objc.Sel("textDidChange:"), notification)
}

// Requests permission to end editing.
//
// textObject: The text object requesting permission to end editing.
//
// # Return Value
// 
// [true] if the text object should proceed to finish editing and resign first
// responder status. If the delegate returns [false], the text object selects
// all of its text and remains the first responder.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The [NSMatrix] method returns [false] if the text field contains invalid
// contents; otherwise it returns the value passed back from
// [ControlTextShouldEndEditing].
// 
// This method is invoked to let the [NSTextField] respond to impending loss
// of first-responder status. This method’s default behavior checks the text
// field for validity; providing that the field contents are deemed valid, and
// providing that the delegate responds, [ControlTextShouldEndEditing] is sent
// to the receiver’s delegate (passing the receiver and `textObject` as
// parameters).
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/textShouldEndEditing(_:)
func (m NSMatrix) TextShouldEndEditing(textObject INSText) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("textShouldEndEditing:"), textObject)
	return rv
}

// Invoked when text editing ends.
//
// notification: The [textDidEndEditingNotification] notification.
// //
// [textDidEndEditingNotification]: https://developer.apple.com/documentation/AppKit/NSControl/textDidEndEditingNotification
//
// # Discussion
// 
// This method’s default behavior is to post an
// [textDidEndEditingNotification] along with the receiving object to the
// default notification center. The posted notification’s user info contains
// the contents of notification’s user info dictionary, plus an additional
// key-value pair. The additional key is “[NSFieldEditor]”; the value for
// this key is the text object that began editing. After posting the
// notification, [NSMatrix] sends an [EndEditing] message to the selected
// cell, draws and makes the selected cell key, and then takes the appropriate
// action based on which key was used to end editing (Return, Tab, or
// Back-Tab).
//
// [textDidEndEditingNotification]: https://developer.apple.com/documentation/AppKit/NSControl/textDidEndEditingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/textDidEndEditing(_:)
func (m NSMatrix) TextDidEndEditing(notification foundation.NSNotification) {
	objc.Send[objc.ID](m.ID, objc.Sel("textDidEndEditing:"), notification)
}

// Specifies whether the receiver’s size information is validated.
//
// flag: [true] to assume that the size information in the receiver is correct. If
// `flag` is [false], the [NSControl] method [calcSize()] will be invoked
// before any further drawing is done.
// //
// [calcSize()]: https://developer.apple.com/documentation/AppKit/NSControl/calcSize()
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/setValidateSize(_:)
func (m NSMatrix) SetValidateSize(flag bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setValidateSize:"), flag)
}

// Changes the width and the height of the receiver’s frame so it exactly
// contains the cells.
//
// # Discussion
// 
// This method does not redraw the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/sizeToCells()
func (m NSMatrix) SizeToCells() {
	objc.Send[objc.ID](m.ID, objc.Sel("sizeToCells"))
}

// Specifies whether the cells in the matrix are scrollable.
//
// flag: [true] to make all the cells in the receiver scrollable, so the text they
// contain scrolls to remain in view if the user types past the edge of the
// cell. If `flag` is [false], all cells are made nonscrolling. The prototype
// cell, if there is one, is also set accordingly
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/setScrollable(_:)
func (m NSMatrix) SetScrollable(flag bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setScrollable:"), flag)
}

// Scrolls the receiver so the specified cell is visible.
//
// row: The row of the cell to make visible.
//
// col: The column of the cell to make visible.
//
// # Discussion
// 
// This method scrolls if the receiver is in a scrolling view and `row` and
// `column` represent a valid cell within the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/scrollCellToVisible(atRow:column:)
func (m NSMatrix) ScrollCellToVisibleAtRowColumn(row int, col int) {
	objc.Send[objc.ID](m.ID, objc.Sel("scrollCellToVisibleAtRow:column:"), row, col)
}

// Displays the cell at the specified row and column.
//
// row: The row containing the cell to draw.
//
// col: The column containing the cell to draw.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/drawCell(atRow:column:)
func (m NSMatrix) DrawCellAtRowColumn(row int, col int) {
	objc.Send[objc.ID](m.ID, objc.Sel("drawCellAtRow:column:"), row, col)
}

// Highlights or unhighlights the cell at the specified row and column
// location.
//
// flag: [true] to highlight the cell; [false] to unhighlight the cell.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// row: The row containing the cell.
//
// col: The column containing the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/highlightCell(_:atRow:column:)
func (m NSMatrix) HighlightCellAtRowColumn(flag bool, row int, col int) {
	objc.Send[objc.ID](m.ID, objc.Sel("highlightCell:atRow:column:"), flag, row, col)
}

// If the selected cell has both an action and a target, sends its action to
// its target.
//
// # Return Value
// 
// [true] if an action was successfully sent to a target. If the selected cell
// is disabled, this method does nothing and returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the cell has an action but no target, its action is sent to the target
// of the receiver. If the cell doesn’t have an action, or if there is no
// selected cell, the receiver sends its own action to its target.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/sendAction()
func (m NSMatrix) SendAction() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("sendAction"))
	return rv
}

// Iterates through the cells in the receiver, sending the specified selector
// to an object for each cell.
//
// selector: The selector to send to the object for each cell. This must represent a
// method that takes a single argument: the id of the current cell in the
// iteration. `aSelector`’s return value must be a BOOL. If `aSelector`
// returns [false] for any cell, [NSMatrix] terminates immediately, without
// sending the message for the remaining cells. If it returns [true],
// [NSMatrix] proceeds to the next cell.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// object: The object that is sent the selector for each cell in the matrix.
//
// flag: [true] if the method should iterate through all cells in the matrix;
// [false] if it should iterate through just the selected cells in the matrix.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Iteration begins with the cell in the upper-left corner of the receiver,
// proceeding through the appropriate entries in the first row, then on to the
// next.
// 
// This method is not invoked to send action messages to target objects in
// response to mouse-down events in the receiver. Instead, you can invoke it
// if you want to have multiple cells in an [NSMatrix] interact with an
// object. For example, you could use it to verify the titles in a list of
// items or to enable a series of radio buttons based on their purpose in
// relation to `anObject`.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/sendAction(_:to:forAllCells:)
func (m NSMatrix) SendActionToForAllCells(selector objc.SEL, object objectivec.IObject, flag bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("sendAction:to:forAllCells:"), selector, object, flag)
}

// Sends the double-click action message to the target of the receiver.
//
// # Discussion
// 
// If the receiver doesn’t have a double-click action, the double-click
// action message of the selected cell (as returned by [selectedCell]) is sent
// to the selected cell’s target. Finally, if the selected cell also has no
// action, then the single-click action of the receiver is sent to the target
// of the receiver.
// 
// If the selected cell is disabled, this method does nothing.
// 
// Your code shouldn’t invoke this method; it’s sent in response to a
// double-click event in the [NSMatrix]. Override it if you need to change the
// search order for an action to send.
//
// [selectedCell]: https://developer.apple.com/documentation/AppKit/NSMatrix/selectedCell
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/sendDoubleAction()
func (m NSMatrix) SendDoubleAction() {
	objc.Send[objc.ID](m.ID, objc.Sel("sendDoubleAction"))
}

// Returns a Boolean value that indicates whether the sender should be
// enabled.
//
// item: The user interface item to validate. You can send `anItem` the [Action] and
// [Tag] messages.
//
// # Return Value
// 
// [true] if the user interface item should be enabled, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations/validateUserInterfaceItem(_:)
func (m NSMatrix) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// Returns the tool tip string to be displayed due to the cursor pausing at
// location `point` within the tool tip rectangle identified by `tag` in the
// view `view`.
//
// # Discussion
// 
// `userData` is additional information provided by the creator of the tool
// tip rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewToolTipOwner/view(_:stringForToolTip:point:userData:)
func (m NSMatrix) ViewStringForToolTipPointUserData(view INSView, tag objectivec.IObject, point corefoundation.CGPoint, data unsafe.Pointer) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("view:stringForToolTip:point:userData:"), view, tag, point, data)
	return foundation.NSStringFromID(rv).String()
}

// The selection mode of the receiver.
//
// # Discussion
// 
// See [NSMatrix.Mode] for possible values.
//
// [NSMatrix.Mode]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/mode-swift.property
func (m NSMatrix) Mode() NSMatrixMode {
	rv := objc.Send[NSMatrixMode](m.ID, objc.Sel("mode"))
	return NSMatrixMode(rv)
}
func (m NSMatrix) SetMode(value NSMatrixMode) {
	objc.Send[struct{}](m.ID, objc.Sel("setMode:"), value)
}

// A Boolean that indicates whether a radio-mode matrix supports an empty
// selection.
//
// # Discussion
// 
// When the value of this property is [true], the matrix allows one or zero
// cells to be selected. When the value of this property is [true], the matrix
// allows one and only one cell (not zero cells) to be selected. This setting
// has effect only in the [NSRadioModeMatrix] selection mode.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/allowsEmptySelection
func (m NSMatrix) AllowsEmptySelection() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowsEmptySelection"))
	return rv
}
func (m NSMatrix) SetAllowsEmptySelection(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowsEmptySelection:"), value)
}

// A Boolean that indicates whether the user can select a rectangle of cells
// in the receiver by dragging the cursor.
//
// # Discussion
// 
// When the value of this property is [true], the matrix allows the user to
// select a rectangle of cells by dragging. When the value of this property is
// [false], selection in the matrix is on a row-by-row basis. The default
// value of this property is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/isSelectionByRect
func (m NSMatrix) SelectionByRect() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isSelectionByRect"))
	return rv
}
func (m NSMatrix) SetSelectionByRect(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSelectionByRect:"), value)
}

// The prototype cell that’s copied whenever the matrix creates a new cell.
//
// # Discussion
// 
// When the value of this property is `nil`, there is no prototype cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/prototype
func (m NSMatrix) Prototype() INSCell {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("prototype"))
	return NSCellFromID(objc.ID(rv))
}
func (m NSMatrix) SetPrototype(value INSCell) {
	objc.Send[struct{}](m.ID, objc.Sel("setPrototype:"), value)
}

// The size of each cell in the matrix.
//
// # Discussion
// 
// This value corresponds to both the width and height of each cell, because
// all cells in an [NSMatrix] are the same size.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/cellSize
func (m NSMatrix) CellSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("cellSize"))
	return corefoundation.CGSize(rv)
}
func (m NSMatrix) SetCellSize(value corefoundation.CGSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setCellSize:"), value)
}

// The vertical and horizontal spacing between cells in the matrix.
//
// # Discussion
// 
// By default, both values are `1.0` in the matrix’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/intercellSpacing
func (m NSMatrix) IntercellSpacing() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("intercellSpacing"))
	return corefoundation.CGSize(rv)
}
func (m NSMatrix) SetIntercellSpacing(value corefoundation.CGSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setIntercellSpacing:"), value)
}

// The number of columns in the matrix.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/numberOfColumns
func (m NSMatrix) NumberOfColumns() int {
	rv := objc.Send[int](m.ID, objc.Sel("numberOfColumns"))
	return rv
}

// The number of rows in the matrix.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/numberOfRows
func (m NSMatrix) NumberOfRows() int {
	rv := objc.Send[int](m.ID, objc.Sel("numberOfRows"))
	return rv
}

// A Boolean that indicates whether the matrix auto-recalculates its cell
// size.
//
// # Discussion
// 
// When the value of this property is [true], auto-recalculation occurs. The
// matrix will adjust its [CellSize] to accommodate its largest cell. Changing
// the `cellSize` does not directly affect the frame of the matrix; however it
// does affect the intrinsic content size, which may cause the receiver to
// resize under Auto Layout. When using Auto Layout, you typically want this
// to be set to [true].
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/autorecalculatesCellSize
func (m NSMatrix) AutorecalculatesCellSize() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("autorecalculatesCellSize"))
	return rv
}
func (m NSMatrix) SetAutorecalculatesCellSize(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAutorecalculatesCellSize:"), value)
}

// The cell that will be clicked when the user presses the Space bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/keyCell
func (m NSMatrix) KeyCell() INSCell {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("keyCell"))
	return NSCellFromID(objc.ID(rv))
}
func (m NSMatrix) SetKeyCell(value INSCell) {
	objc.Send[struct{}](m.ID, objc.Sel("setKeyCell:"), value)
}

// An array containing all of the matrix’s highlighted cells plus its
// selected cell.
//
// # Discussion
// 
// See the class description for a discussion of the selected cell.
// 
// As an alternative to using [SetSelectionFromToAnchorHighlight] for
// programmatically making discontiguous selections of cells in a matrix, you
// could first set the single selected cell and then set subsequent cells to
// be highlighted; afterwards you can access [SelectedCells] to obtain the
// selection of cells.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/selectedCells
func (m NSMatrix) SelectedCells() []NSCell {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("selectedCells"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSCell {
		return NSCellFromID(id)
	})
}

// The column number of the selected cell.
//
// # Discussion
// 
// When the value of this property is `–1`, no cells are selected. If cells
// in multiple columns are selected, the value of this property is the number
// of the last (rightmost) column containing a selected cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/selectedColumn
func (m NSMatrix) SelectedColumn() int {
	rv := objc.Send[int](m.ID, objc.Sel("selectedColumn"))
	return rv
}

// The row number of the selected cell.
//
// # Discussion
// 
// When the value of this property is `–1`, no cells are selected. If cells
// in multiple rows are selected, the value of this property is the number of
// the last row containing a selected cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/selectedRow
func (m NSMatrix) SelectedRow() int {
	rv := objc.Send[int](m.ID, objc.Sel("selectedRow"))
	return rv
}

// An array containing the cells of the matrix.
//
// # Discussion
// 
// The cells in the array are row-ordered; that is, the first row of cells
// appears first in the array, followed by the second row, and so forth.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/cells
func (m NSMatrix) Cells() []NSCell {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("cells"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSCell {
		return NSCellFromID(id)
	})
}

// The background color of the matrix (the space between the cells).
//
// # Discussion
// 
// The value of this property is the background color used to fill the space
// between cells or the space behind any non-opaque cells. The default
// background color is the color returned by the [NSColor] method
// [controlColor].
//
// [controlColor]: https://developer.apple.com/documentation/AppKit/NSColor/controlColor
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/backgroundColor
func (m NSMatrix) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (m NSMatrix) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](m.ID, objc.Sel("setBackgroundColor:"), value)
}

// The background color of the matrix’s cells.
//
// # Discussion
// 
// The value of this property is the background color used to fill the space
// behind non-opaque cells. The default background color is the color returned
// by the [NSColor] method [controlColor].
//
// [controlColor]: https://developer.apple.com/documentation/AppKit/NSColor/controlColor
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/cellBackgroundColor
func (m NSMatrix) CellBackgroundColor() INSColor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("cellBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (m NSMatrix) SetCellBackgroundColor(value INSColor) {
	objc.Send[struct{}](m.ID, objc.Sel("setCellBackgroundColor:"), value)
}

// A Boolean that indicates whether the matrix draws its background.
//
// # Discussion
// 
// When the value of this property is [true], the matrix draws its background
// (the space between the cells).
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/drawsBackground
func (m NSMatrix) DrawsBackground() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("drawsBackground"))
	return rv
}
func (m NSMatrix) SetDrawsBackground(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setDrawsBackground:"), value)
}

// A Boolean that indicates whether the matrix draws the background within
// each of its cells.
//
// # Discussion
// 
// When the value of this property is [true], the matrix draws the cell
// background.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/drawsCellBackground
func (m NSMatrix) DrawsCellBackground() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("drawsCellBackground"))
	return rv
}
func (m NSMatrix) SetDrawsCellBackground(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setDrawsCellBackground:"), value)
}

// A Boolean that indicates whether pressing the Tab key advances the key cell
// to the next selectable cell.
//
// # Discussion
// 
// When the value of this property is [true], pressing the Tab key should
// advance the key cell to the next selectable cell in the receiver. When the
// value of this property is [false] or if there aren’t any selectable cells
// after the current one, the next view in the window becomes key when the
// user presses the Tab key.
// 
// Pressing Shift-Tab causes the key cell to advance in the opposite direction
// (if the value of this property is [false], or if there aren’t any
// selectable cells before the current one, the previous view in the window
// becomes key).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/tabKeyTraversesCells
func (m NSMatrix) TabKeyTraversesCells() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("tabKeyTraversesCells"))
	return rv
}
func (m NSMatrix) SetTabKeyTraversesCells(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setTabKeyTraversesCells:"), value)
}

// The delegate for messages from the field editor.
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/delegate
func (m NSMatrix) Delegate() NSMatrixDelegate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("delegate"))
	return NSMatrixDelegateObjectFromID(rv)
}
func (m NSMatrix) SetDelegate(value NSMatrixDelegate) {
	objc.Send[struct{}](m.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean that indicates whether the cell sizes change when the receiver is
// resized.
//
// # Discussion
// 
// When the value of this property is [true], whenever the matrix is resized,
// the sizes of the cells change in proportion, keeping the intercell space
// constant. This property verifies that the cell sizes and intercell spacing
// add up to the exact size of the matrix, adjusting the size of the cells and
// updating the matrix if they don’t. When the value of this property is
// [false], then the intercell spacing and cell size remain constant.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/autosizesCells
func (m NSMatrix) AutosizesCells() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("autosizesCells"))
	return rv
}
func (m NSMatrix) SetAutosizesCells(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAutosizesCells:"), value)
}

// The action sent to the target of the receiver when the user double-clicks a
// cell.
//
// # Discussion
// 
// The double-click action of an [NSMatrix] is sent after the appropriate
// single-click action (for the [NSCell] clicked, or for the [NSMatrix] if the
// [NSCell] doesn’t have its own action). If there is no double-click action
// and the [NSMatrix] doesn’t ignore multiple clicks, the single-click
// action is sent twice. If the value of this property is a non-`nil`
// selector, this property also sets `ignoresMultiClick` to [false];
// otherwise, it leaves `ignoresMultiClick` unchanged.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/doubleAction
func (m NSMatrix) DoubleAction() objc.SEL {
	rv := objc.Send[objc.SEL](m.ID, objc.Sel("doubleAction"))
	return rv
}
func (m NSMatrix) SetDoubleAction(value objc.SEL) {
	objc.Send[struct{}](m.ID, objc.Sel("setDoubleAction:"), value)
}

// The flags in effect at the mouse-down event that started the current
// tracking session.
//
// # Discussion
// 
// The [NSMatrix] [MouseDown] method obtains these flags by sending a
// [ModifierFlags] message to the event passed into [MouseDown]. Use this
// property if you want to access these flags. This property is valid only
// during tracking; it isn’t useful if the target of the receiver initiates
// another tracking loop as part of its action method (as a cell that pops up
// a pop-up list does, for example).
//
// See: https://developer.apple.com/documentation/AppKit/NSMatrix/mouseDownFlags
func (m NSMatrix) MouseDownFlags() int {
	rv := objc.Send[int](m.ID, objc.Sel("mouseDownFlags"))
	return rv
}

			// Protocol methods for NSUserInterfaceValidations
			

			// Protocol methods for NSViewToolTipOwner
			

