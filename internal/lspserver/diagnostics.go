package lspserver

import (
	"log"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func getDiagnosticsForFile(text *FileContent) []Diagnostic {
	diagnostics := []Diagnostic{}

	if len(text.lines) == 0 {
		return diagnostics
	}
	if len(text.lines) == 1 {
		return diagnostics
	}
	return diagnostics
}

func produceDiagnostics(uri string, content *FileContent) *PublishDiagnosticsNotification {
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
