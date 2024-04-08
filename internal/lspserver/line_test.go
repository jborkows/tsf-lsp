package lspserver

import (
	"fmt"
	"testing"

	"math/rand/v2"

	"github.com/stretchr/testify/assert"
)

func TestEmptyFileShouldBeParsedAsEmpty(t *testing.T) {
	line := ParseLine("", FakedCategoriesProvider())
	_, ok := IsEmpty(line)
	assert.True(t, ok, "Should be parsed as empty")
}

func TestShouldParseLineForFreeDayAsValid(t *testing.T) {
	days := FreeDays().FreeDaysCategories()
	day := days[rand.IntN(len(days))]
	line := ParseLine(string(day), FakedCategoriesProvider())
	valid, ok := IsValid(line)
	assert.True(t, ok, fmt.Sprintf("should be valid but was %#v", line))
	assert.Equal(t, day, valid.Category)
	assert.Equal(t, 8, valid.Hours)
	assert.Equal(t, 0, valid.Minutes)
	assert.Nil(t, valid.Description)
	assert.Nil(t, valid.Warnings)
}

func TestShouldFreeDayShouldNotContainAnyOtherInfo(t *testing.T) {
	days := FreeDays().FreeDaysCategories()
	day := days[rand.IntN(len(days))]
	inputText := fmt.Sprintf("%s\taaaaa\t1h 30m", day)
	line := ParseLine(inputText, FakedCategoriesProvider())
	_, ok := IsValid(line)
	assert.False(t, ok, fmt.Sprintf("should not be valid but was %#v", line))
	invalid, ok := IsInValid(line)

	assert.True(t, ok, fmt.Sprintf("should be invalid but was %#v", line))
	assert.Equal(t, 1, len(invalid.ErrorsOnLine))
	anError := invalid.ErrorsOnLine[0]
	assert.Equal(t, ErrorOnLine{MessageOnLine: MessageOnLine{
		Message:      "Free day should not contain any additional info",
		StartingFrom: len(string(day)),
		FinishedAt:   len(inputText),
	},
	}, anError)
}

type projectCategoriesProvider struct {
	categories []Category
}

func (self projectCategoriesProvider) ProjectCategories() []Category {
	return self.categories
}

func FakedCategoriesProvider(names ...string) ProjectCategoriesProvider {
	categories := make([]Category, len(names))
	for _, name := range names {
		categories = append(categories, CategoryFrom(name))
	}
	return projectCategoriesProvider{
		categories,
	}

}

func TestShouldParseProjectCategory(t *testing.T) {
	line := ParseLine(fmt.Sprintf("%s\tonly hours\t 2h", "Project A"), FakedCategoriesProvider("Project A"))
	valid, ok := IsValid(line)
	assert.True(t, ok, fmt.Sprintf("should be valid but was %#v", line))
	assert.Equal(t, CategoryFrom("Project A"), valid.Category)
	assert.Equal(t, 2, valid.Hours)
	assert.Equal(t, 0, valid.Minutes)
	assert.Equal(t, "only hours", valid.Description)
	assert.Nil(t, valid.Warnings)

	line = ParseLine(fmt.Sprintf("%s\thours with minutes\t 1h 30m", "Project A"), FakedCategoriesProvider("Project A"))
	valid, ok = IsValid(line)
	assert.True(t, ok, fmt.Sprintf("should be valid but was %#v", line))
	assert.Equal(t, CategoryFrom("Project A"), valid.Category)
	assert.Equal(t, 1, valid.Hours)
	assert.Equal(t, 30, valid.Minutes)
	assert.Equal(t, "hours with minutes", valid.Description)
	assert.Nil(t, valid.Warnings)

	line = ParseLine(fmt.Sprintf("%s\tonly minutes\t30m", "Project A"), FakedCategoriesProvider("Project A"))
	valid, ok = IsValid(line)
	assert.True(t, ok, fmt.Sprintf("should be valid but was %#v", line))
	assert.Equal(t, CategoryFrom("Project A"), valid.Category)
	assert.Equal(t, 0, valid.Hours)
	assert.Equal(t, 30, valid.Minutes)
	assert.Equal(t, "only minutes", valid.Description)
	assert.Nil(t, valid.Warnings)
}

/** TODO
wrongly parsed:
- not enough fields
- wrong time format given
- minutes not 0,15,30,45

warnings in success:
- jira ticket prefix wrongly typed JIRRA, JIR, JRIA


*/
