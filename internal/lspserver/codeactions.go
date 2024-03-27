package lspserver

import (
	"fmt"
	"log"
	"strings"

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
	document := s.Documents[uri]
	wordToComplete, err := document.word(position)

	if err != nil {
		return nil, err
	}
	wordToComplete = strings.ToLower(wordToComplete)
	if strings.Contains("dupa", wordToComplete) {
		log.Printf("Found word to complete: %s", wordToComplete)
		actions := []CodeAction{
			{
				Title: "dupa",
				Kind:  "quickfix",
				Edit: &WorkspaceEdit{
					Changes: map[string][]TextEdit{
						uri: {
							{
								Range: Range{
									Start: Position{
										Line:      position.Line,
										Character: position.Character,
									},
									End: Position{
										Line:      position.Line,
										Character: position.Character + len(wordToComplete),
									},
								},
								NewText: "kupa",
							},
						},
					},
				},
			},
			{
				Title: "dupa i kamieni kupa",
				Kind:  "refactor",
				Command: &Command{
					Command:   "some_command",
					Title:     "Execute on server",
					Arguments: "aaaa",
				},
			},
		}
		return actions, nil
	}
	return nil, nil
}
