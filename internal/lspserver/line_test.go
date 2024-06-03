package lspserver

import (
	"fmt"
	"testing"

	"math/rand/v2"

	"github.com/stretchr/testify/assert"
)

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

func TestEmptyFileShouldBeParsedAsEmpty(t *testing.T) {
	line := ParseLine("", FakedCategoriesProvider())
	_, ok := IsEmpty(line)
	assert.True(t, ok, "Should be parsed as empty")
}

func TestAShouldParseLineForFreeDayAsValid(t *testing.T) {
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

/** TODO
wrongly parsed:
- not enough fields
- wrong time format given
- minutes not 0,15,30,45

warnings in success:
- jira ticket prefix wrongly typed JIRRA, JIR, JRIA


*/
