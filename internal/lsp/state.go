package lsp

type State struct {
	// Map of file names to contents
	Documents       map[string]string
	openListeners   []func(uri string, content string)
	changeListeners []func(uri string, content string)
}

func (s *State) AddOpenListener(listener func(uri string, content string)) {
	s.openListeners = append(s.openListeners, listener)
}

func (s *State) AddChangeListener(listener func(uri string, content string)) {
	s.changeListeners = append(s.changeListeners, listener)
}

func newState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
	for _, listener := range s.openListeners {
		listener(uri, text)
	}

}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
	for _, listener := range s.changeListeners {
		listener(uri, text)
	}

}

func LineRange(line, start, end int) Range {
	return Range{
		Start: Position{
			Line:      line,
			Character: start,
		},
		End: Position{
			Line:      line,
			Character: end,
		},
	}
}

var state State

func ProvideState() *State {
	return &state
}

func init() {
	state = newState()
}
