package lspserver

import "strings"

func ParseLine(input string) FileLine {
	trimmed := strings.TrimSpace(input)
	if len(trimmed) == 0 {
		return &EmptyLine{}
	}
	splitted := strings.Split(trimmed, "\t")
	category := CategoryFrom(splitted[0])
	if FreeDays().IsFreeDay(category) {
		if len(splitted) > 1 {
			return &InValidLine{
				ErrorsOnLine: []ErrorOnLine{
					{
						Message:      "Free day should not contain any additional info",
						StartingFrom: len(category),
						FinishedAt:   len(input),
					},
				},
			}
		}

		return &ValidLine{
			Category:    category,
			Description: nil,
			Hours:       8,
			Minutes:     0,
		}
	}

	return &EmptyLine{}
}

type FileLine interface {
	lineType() string
}

type ValidLine struct {
	Category    Category
	Description *string
	Hours       int
	Minutes     int
}

func (self *ValidLine) lineType() string {
	return "valid"
}

type EmptyLine struct {
}

func (self *EmptyLine) lineType() string {
	return "empty"
}

type ErrorOnLine struct {
	Message      string
	StartingFrom int
	FinishedAt   int
}

type InValidLine struct {
	ErrorsOnLine []ErrorOnLine
}

func (self *InValidLine) lineType() string {
	return "invalid"
}

func IsValid(fileLine FileLine) (*ValidLine, bool) {
	if fileLine.lineType() != "valid" {
		return nil, false
	}

	return fileLine.(*ValidLine), true
}

func IsEmpty(fileLine FileLine) (*EmptyLine, bool) {
	if fileLine.lineType() != "empty" {
		return nil, false
	}

	return fileLine.(*EmptyLine), true
}

func IsInValid(fileLine FileLine) (*InValidLine, bool) {
	if fileLine.lineType() != "invalid" {
		return nil, false
	}

	return fileLine.(*InValidLine), true
}
