package lspserver

import (
	"fmt"
	"log"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func (s *State) CodeActions(position Position, uri string) ([]CodeAction, error) {
	actions, err := s.findCodeActions(uri, position)
	if err != nil {
		return nil, fmt.Errorf("Error finding code actions: %w", err)
	}
	log.Printf("Found code actions: %v", actions)
	return actions, nil
}

func (s *State) findCodeActions(uri string, position Position) ([]CodeAction, error) {

	return nil, nil
}
