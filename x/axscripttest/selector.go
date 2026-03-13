package axscripttest

import (
	"fmt"
	"strings"

	"github.com/tmc/apple/x/axuiautomation"
)

type selector struct {
	role  string
	title string
	id    string
}

// parseSelector extracts role=, title=, and id= pairs from args.
// It returns the parsed selector and any remaining (non-selector) args.
func parseSelector(args []string) (selector, []string, error) {
	var sel selector
	var rest []string
	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "role="):
			sel.role = strings.TrimPrefix(arg, "role=")
		case strings.HasPrefix(arg, "title="):
			sel.title = strings.TrimPrefix(arg, "title=")
		case strings.HasPrefix(arg, "id="):
			sel.id = strings.TrimPrefix(arg, "id=")
		default:
			rest = append(rest, arg)
		}
	}
	if sel.role == "" && sel.title == "" && sel.id == "" {
		return sel, rest, fmt.Errorf("no selector specified (use role=, title=, or id=)")
	}
	return sel, rest, nil
}

func (sel selector) query(app *axuiautomation.Application) *axuiautomation.ElementQuery {
	q := app.Descendants()
	if sel.role != "" {
		q = q.ByRole(sel.role)
	}
	if sel.title != "" {
		q = q.ByTitle(sel.title)
	}
	if sel.id != "" {
		q = q.ByIdentifier(sel.id)
	}
	return q
}

func (sel selector) find(app *axuiautomation.Application) (*axuiautomation.Element, error) {
	elem := sel.query(app).First()
	if elem == nil {
		return nil, fmt.Errorf("element not found: %s", sel)
	}
	return elem, nil
}

func (sel selector) String() string {
	var parts []string
	if sel.role != "" {
		parts = append(parts, "role="+sel.role)
	}
	if sel.title != "" {
		parts = append(parts, "title="+sel.title)
	}
	if sel.id != "" {
		parts = append(parts, "id="+sel.id)
	}
	return strings.Join(parts, " ")
}
