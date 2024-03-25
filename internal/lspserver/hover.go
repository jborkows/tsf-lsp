package lspserver

import (
	"fmt"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func (s *State) Hover(id int, uri string, position Position) HoverResponse {
	// In real life, this would look up the type in our type analysis code...

	document := s.Documents[uri]
	line := document.lines[position.Line]

	return HoverResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: HoverResult{
			Contents: fmt.Sprintf("Line: %d, Characters: %d", position.Line+1, len(line)),
		},
	}
}
