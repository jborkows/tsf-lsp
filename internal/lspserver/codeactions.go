package lspserver

import (
	"log"
	"strings"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func (s *State) CodeAction(id int, position Position, uri string) TextDocumentCodeActionResponse {
	actions, err := s.findCodeActions(uri, position)
	if err != nil {
		log.Printf("Error finding code actions: %v", err)
		return TextDocumentCodeActionResponse{
			Response: Response{
				RPC: "2.0",
				ID:  &id,
			},
			Result: nil,
		}
	}
	log.Printf("Found code actions: %v", actions)
	return TextDocumentCodeActionResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
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
					Command:   "some command",
					Title:     "Execute on server",
					Arguments: "aaaa",
				},
			},
		}
		return actions, nil
	}
	return nil, nil
}