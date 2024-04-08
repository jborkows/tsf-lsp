package lspserver

import (
	"fmt"
	"testing"

	"math/rand/v2"

	"github.com/stretchr/testify/assert"
)

func TestEmptyFileShouldBeParsedAsEmpty(t *testing.T) {
	line := ParseLine("")
	_, ok := IsEmpty(line)
	assert.True(t, ok, "Should be parsed as empty")
}

func TestShouldParseLineForFreeDayAsValid(t *testing.T) {
	days := FreeDays().FreeDaysCategories()
	day := days[rand.IntN(len(days))]
	line := ParseLine(string(day))
	valid, ok := IsValid(line)
	assert.True(t, ok, fmt.Sprintf("should be valid but was %#v", line))
	assert.Equal(t, day, valid.Category)
	assert.Equal(t, 8, valid.Hours)
	assert.Equal(t, 0, valid.Minutes)
	assert.Nil(t, valid.Description)
}

func TestShouldFreeDayShouldNotContainAnyOtherInfo(t *testing.T) {
	days := FreeDays().FreeDaysCategories()
	day := days[rand.IntN(len(days))]
	inputText := fmt.Sprintf("%s\taaaaa\t1h 30m", day)
	line := ParseLine(inputText)
	_, ok := IsValid(line)
	assert.False(t, ok, fmt.Sprintf("should not be valid but was %#v", line))
	invalid, ok := IsInValid(line)

	assert.True(t, ok, fmt.Sprintf("should be invalid but was %#v", line))
	assert.Equal(t, 1, len(invalid.ErrorsOnLine))
	anError := invalid.ErrorsOnLine[0]
	assert.Equal(t, ErrorOnLine{
		Message:      "Free day should not contain any additional info",
		StartingFrom: len(string(day)),
		FinishedAt:   len(inputText),
	}, anError)
}
