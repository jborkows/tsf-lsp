package lsp

import (
	"fmt"
	"log"
	"strings"
)

func getDiagnosticsForFile(text string) []Diagnostic {
	diagnostics := []Diagnostic{}
	forbiddenWords := []string{"dupa", "fuck"}

	for row, line := range strings.Split(text, "\n") {
		for _, word := range forbiddenWords {
			lowered := strings.ToLower(line)
			if strings.Contains(lowered, word) {
				idx := strings.Index(lowered, word)
				log.Printf("Found '%s' in '%s' at %d", line, word, idx)
				diagnostics = append(diagnostics, Diagnostic{
					Range:    LineRange(row, idx, idx+len(word)),
					Severity: 1,
					Source:   "Words",
					Message:  fmt.Sprintf("Please do not use %s", word),
				})
			}
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

func RegisterDiagnostics(state *State, responseWriter func(any)) func() error {
	workChannel := make(chan func(), 20)

	//create pool of 5 workers waiting for work till the channel is closed
	for i := 0; i < 5; i++ {
		go func() {
			for work := range workChannel {
				work()
			}
		}()
	}

	closer := func() error {
		close(workChannel)
		return nil
	}

	state.AddOpenListener(func(uri string, content string) {
		workChannel <- func() {
			result := produceDiagnostics(uri, content)
			if result != nil {
				responseWriter(result)
			}
		}
	})

	state.AddChangeListener(func(uri string, content string) {
		workChannel <- func() {
			result := produceDiagnostics(uri, content)
			if result != nil {
				responseWriter(result)
			}
		}
	})

	return closer
}
