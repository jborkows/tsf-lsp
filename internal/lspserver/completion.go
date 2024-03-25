package lspserver

import (
	"fmt"
	"strings"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func (s *State) Completion(id int, uri string, position Position) CompletionResponse {
	completions, err := s.findCompletions(uri, position)
	if err != nil {
		return CompletionResponse{
			Response: Response{
				RPC: "2.0",
				ID:  &id,
			},
			Result: nil,
		}
	}
	return CompletionResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: completions,
	}
}
func (s *State) findCompletions(uri string, position Position) ([]CompletionItem, error) {

	document := s.Documents[uri]
	wordToComplete, err := document.word(position)

	if err != nil {
		return nil, fmt.Errorf("Error getting word: %w", err)
	}
	wordToComplete = strings.ToLower(wordToComplete)
	if strings.Contains("dupa", wordToComplete) {
		completions := []CompletionItem{
			{
				Label:  "dupa",
				Detail: "dupa dupa dupa",
			},
			{
				Label:         "dupa i kamieni kupa",
				Detail:        "dupa i kamieni kupa",
				Documentation: "podstawa panstwa",
			},
			{
				Label:         "pupka",
				Detail:        "pupka niemowlaka",
				Documentation: "troche kultury",
			},
		}
		return completions, nil
	} else {
		//they are also filtered on client...
		return []CompletionItem{
			{
				Label:         "to sie uzupelni",
				Detail:        "to jest opis",
				Documentation: "to jest dokumentacja",
			},
			{
				Label:         "this will be completed",
				Detail:        "this is a description",
				Documentation: "this is a documentation",
			},
		}, nil
	}

}
