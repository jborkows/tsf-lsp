package lspserver

import (
	"fmt"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func (s *State) Completion(uri string, position Position) ([]CompletionItem, error) {
	completions, err := s.findCompletions(uri, position)
	if err != nil {
		return nil, fmt.Errorf("Error finding completions: %w", err)
	}
	return completions, nil
}
func (s *State) findCompletions(uri string, position Position) ([]CompletionItem, error) {
	completions := []CompletionItem{}
	return completions, nil
}
