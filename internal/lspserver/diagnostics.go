package lspserver

import (
	"fmt"
	"log"
	"strings"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func getDiagnosticsForFile(text string) []Diagnostic {
	diagnostics := []Diagnostic{}
	forbiddenWords := []string{"dupa", "fuck"}

	for row, line := range strings.Split(text, "\n") {
		for _, word := range forbiddenWords {
			lowered := strings.ToLower(line)
			idx := strings.Index(lowered, word)
			if idx < 0 {
				continue
			}
			log.Printf("Found '%s' in '%s' at %d", line, word, idx)
			diagnostics = append(diagnostics, Diagnostic{
				Range:    LineRange(row, idx, idx+len(word)),
				Severity: 1,
				Source:   "Words",
				Message:  fmt.Sprintf("Please do not use %s", word),
			})
		}

	}

	return diagnostics
}

func produceDiagnostics(uri string, content string) *PublishDiagnosticsNotification {
	diagnostics := getDiagnosticsForFile(content)
	if len(diagnostics) > 1 {
		log.Printf("Sending diagnostics %v", diagnostics)
	}
	return &PublishDiagnosticsNotification{
		Notification: Notification{
			RPC:    "2.0",
			Method: "textDocument/publishDiagnostics",
		},
		Params: PublishDiagnosticsParams{
			URI:         uri,
			Diagnostics: diagnostics,
		},
	}
}
