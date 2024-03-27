package lspserver

import (
	. "github.com/jborkows/tsf-lsp/internal/lsp"
)

func (state *State) Color(uri string) []ColorInformation {
	var colors []ColorInformation

	colors = append(colors, ColorInformation{
		Range: Range{
			Start: Position{
				Line:      0,
				Character: 0,
			},
			End: Position{
				Line:      0,
				Character: 5,
			},
		},
		Color: Color{
			Red:   1.0,
			Green: 0.0,
			Blue:  0.0,
			Alpha: 1.0,
		},
	}, ColorInformation{
		Range: Range{
			Start: Position{
				Line:      1,
				Character: 5,
			},
			End: Position{
				Line:      1,
				Character: 10,
			},
		},
		Color: Color{
			Red:   0.8,
			Green: 1.0,
			Blue:  0.2,
			Alpha: 0.5,
		},
	},
	)
	return colors
}
