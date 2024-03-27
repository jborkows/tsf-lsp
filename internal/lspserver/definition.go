package lspserver

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func (s *State) DefinitionLocation(uri string, position Position) *Location {

	location, err := s.findDefinition(uri, position)
	if err != nil {
		log.Printf("Error finding definition: %v", err)
		return nil
	}

	log.Printf("Found definition at %v", location)
	return location

}
func (s *State) findDefinition(uri string, position Position) (*Location, error) {
	log.Printf("Finding definition for %s at %v", uri, position)
	document := s.Documents[uri]
	//read the word to find from position.Character to end of line or first space

	wordToFind, err := document.word(position)
	if err != nil {
		return nil, fmt.Errorf("Error getting word: %w", err)
	}

	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("Error parsing URI: %w", err)
	}
	filePath := u.Path
	// If on Windows, remove the leading slash if it looks like a drive letter is present
	if runtime.GOOS == "windows" && len(filePath) > 0 && filePath[1] == ':' {
		filePath = strings.TrimPrefix(filePath, "/")
	}
	directory := path.Dir(filePath)
	definitions := path.Join(directory, "keywords.txt")
	if _, err := os.Stat(definitions); os.IsNotExist(err) {
		return nil, fmt.Errorf("No definitions file found")
	}

	// Read the definitions filePath
	file, err := os.Open(definitions)
	if err != nil {
		return nil, fmt.Errorf("Error opening definitions file: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	wordToFind = strings.ToLower(wordToFind)

	log.Printf("Word to find: %s", wordToFind)
	for scanner.Scan() {
		lineNumber++
		line := strings.ToLower(scanner.Text())
		index := strings.Index(line, wordToFind)
		if index != -1 {

			absPath, err := filepath.Abs(definitions)
			if err != nil {
				return nil, fmt.Errorf("Error getting absolute path: %w", err)
			}
			// Create a URL struct
			u := &url.URL{
				Scheme: "file",
				Path:   absPath,
			}
			uri := u.String()
			return &Location{
				URI: uri,
				Range: Range{
					Start: Position{
						Line:      lineNumber - 1,
						Character: index,
					},
					End: Position{
						Line:      lineNumber - 1,
						Character: index + len(wordToFind),
					},
				},
			}, nil

		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading definitions file: %w", err)
	}

	return nil, fmt.Errorf("No definition found")
}
