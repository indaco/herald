package herald

// ListKind determines whether a list is rendered as unordered or ordered.
type ListKind int

const (
	// Unordered renders list items with bullet characters.
	Unordered ListKind = iota
	// Ordered renders list items with sequential numbers.
	Ordered
)

// ListItem represents a single entry in a nested list. It may contain
// children to create hierarchical lists.
type ListItem struct {
	Text     string
	Children []ListItem
	Kind     ListKind // how Children are rendered (UL vs OL)
}

// Item creates a leaf ListItem with no children.
func Item(text string) ListItem {
	return ListItem{Text: text}
}

// Items converts multiple strings into a slice of leaf ListItems.
func Items(texts ...string) []ListItem {
	items := make([]ListItem, len(texts))
	for i, t := range texts {
		items[i] = ListItem{Text: t}
	}
	return items
}

// ItemWithChildren creates a ListItem whose children are rendered as an
// unordered sub-list.
func ItemWithChildren(text string, children ...ListItem) ListItem {
	return ListItem{Text: text, Children: children, Kind: Unordered}
}

// ItemWithOLChildren creates a ListItem whose children are rendered as an
// ordered sub-list.
func ItemWithOLChildren(text string, children ...ListItem) ListItem {
	return ListItem{Text: text, Children: children, Kind: Ordered}
}
