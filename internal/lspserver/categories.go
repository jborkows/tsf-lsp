package lspserver

import "strings"

type Category (string)

func CategoryFrom(value string) Category {
	return Category(strings.ToLower(strings.TrimSpace(value)))
}

type FreeDaysProvider interface {
	FreeDaysCategories() []Category
	IsFreeDay(category Category) bool
}

type ProjectCategoriesProvider interface {
	ProjectCategories() []Category
}

type freeDay struct {
	freeDays []Category
}

func FreeDays() FreeDaysProvider {
	return freeInstance
}

var freeInstance *freeDay

func init() {

	categories := make([]Category, 2)
	categories[0] = CategoryFrom("Holiday")
	categories[1] = CategoryFrom("Free")
	freeInstance = &freeDay{
		freeDays: categories,
	}
}

func (self *freeDay) FreeDaysCategories() []Category {
	return self.freeDays
}

func (self *freeDay) IsFreeDay(category Category) bool {

	for _, free := range self.freeDays {

		if free == category {
			return true
		}
	}
	return false
}
