package lspserver

import (
	// "fmt"
	// "unicode"
	//
	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func (content *FileContent) word(position Position) (string, error) {
	// line := position.Line
	// character := position.Character
	// if line < 0 || line >= len(content.lines) {
	// 	return "", fmt.Errorf("Line %d out of range", line)
	// }
	// if character < 0 || character > len(content.lines[line]) {
	// 	return "", fmt.Errorf("Character %d out of range", character)
	// }
	// runes := []rune(content.lines[line])
	// newStart := startFromIndex(character, runes)
	// for i := newStart; i < len(runes); i++ {
	// 	if unicode.IsSpace(runes[i]) {
	// 		return string(runes[newStart:i]), nil
	// 	}
	// }
	// return string(runes[newStart:]), nil
	return "", nil
}

// func startFromIndex(character int, runes []rune) int {
// 	newStart := character
// 	if newStart >= len(runes) {
// 		newStart = len(runes) - 1
// 	}
// 	if newStart <= 0 {
// 		return 0
// 	}
// 	for i := newStart; i >= 0; i-- {
// 		newStart = i
// 		if unicode.IsSpace(runes[i]) {
// 			newStart++
// 			break
// 		}
// 	}
// 	newStart++
// 	if newStart > len(runes) {
// 		newStart = len(runes)
// 	}
// 	return newStart
// }
