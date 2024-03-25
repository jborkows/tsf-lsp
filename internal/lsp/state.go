package lsp

import (
	"fmt"
	"log"
	"strings"
)

type State struct {
	// Map of file names to contents
	Documents   map[string]string
	notifier    NotificationSender
	workChannel chan func()
}

func (s *State) Close() {
	close(s.workChannel)
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
	s.workChannel <- func() {
		result := produceDiagnostics(uri, text)
		log.Printf("Got diagnostics %v", result)
		s.notifier(result)
	}
}

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

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text

	s.workChannel <- func() {
		result := produceDiagnostics(uri, text)
		if result != nil {
			s.notifier(result)
		}
	}
}

func LineRange(line, start, end int) Range {
	return Range{
		Start: Position{
			Line:      line,
			Character: start,
		},
		End: Position{
			Line:      line,
			Character: end,
		},
	}
}

type NotificationSender func(v any)

func ProvideState(notificationSender NotificationSender) *State {

	workChannel := make(chan func(), 20)

	//create pool of 5 workers waiting for work till the channel is closed
	for i := 0; i < 5; i++ {
		go func() {
			for work := range workChannel {
				work()
			}
		}()
	}
	return &State{
		Documents:   map[string]string{},
		notifier:    notificationSender,
		workChannel: workChannel,
	}

}
