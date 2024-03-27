package lspserver

import (
	"fmt"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func (s *State) Hover(uri string, position Position) HoverResult {
	// In real life, this would look up the type in our type analysis code...

	document := s.Documents[uri]
	line := document.lines[position.Line]

	return HoverResult{
		Contents: fmt.Sprintf("Line: %d, Characters: %d", position.Line+1, len(line)),
	}
}
