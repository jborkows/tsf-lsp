package lspserver

import (
	"log"
	"strings"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

type FileContent struct {
	lines []string
}
type State struct {
	// Map of file names to contents
	Documents   map[string]FileContent
	notifier    NotificationSender
	workChannel chan func()
}

func (s *State) Close() {
	close(s.workChannel)
}

func fileContentFromText(text string) FileContent {
	return FileContent{
		lines: strings.Split(text, "\n"),
	}
}

func (s *State) OpenDocument(uri, text string) {

	lines := fileContentFromText(text)
	s.Documents[uri] = lines
	s.workChannel <- func() {
		result := produceDiagnostics(uri, &lines)
		log.Printf("Got diagnostics %v", result)
		s.notifier(result)
	}
}

func (s *State) UpdateDocument(uri, text string) {
	lines := fileContentFromText(text)
	s.Documents[uri] = lines

	s.workChannel <- func() {
		result := produceDiagnostics(uri, &lines)
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
	initiatilizeWorkers(workChannel)
	return &State{
		Documents:   map[string]FileContent{},
		notifier:    notificationSender,
		workChannel: workChannel,
	}

}

func initiatilizeWorkers(workChannel chan func()) {
	for i := 0; i < 5; i++ {
		go func() {
			for work := range workChannel {
				work()
			}
		}()
	}
}
