package model

import (
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanReadTestObject(t *testing.T) {
	assert.NotEmpty(t, test_project_directory())
}

func test_project_directory() string {
	return os.Getenv("MY_TEST_PROJECT_DIRECTORY")
}

var counter int32 = 0

func create_some_fake_file(t *testing.T) string {

	entryDirectory := filepath.Join(test_project_directory(), "aaaa", "bbbb", "cccc", "dddd")
	error := os.MkdirAll(entryDirectory, 0700)
	if error != nil {
		t.Fatal("Cannot create folder")
	}
	value := atomic.AddInt32(&counter, 1)
	testfile := filepath.Join(entryDirectory, fmt.Sprintf("aaa%d", value))
	_, error = os.Create(testfile)
	if error != nil {
		t.Fatal("Cannot create file")
	}
	return testfile
}

func TestFindProjectRoot(t *testing.T) {
	filesUri := FilePathToUri(create_some_fake_file(t))
	t.Logf("Created %s", filesUri)
	root := FindProjectRoot(filesUri)
	assert.Equal(t, test_project_directory(), root)
}

/** Will not work since project directory is set only once... -> it will stay*/
// func TestReadingConfigurationWithoutOpeningFile(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			if r != "root directory was not initialized, please open some timesheet file" {
// 				t.Errorf("Unexpected panic message: %v", r)
// 			}
// 		} else {
// 			t.Errorf("Should ended with error")
// 		}
// 	}()
// 	provider := ApplicationCategoryProvider()
// 	_, _ = provider.categories()
// }

func TestCanReadConfiguration(t *testing.T) {

	filesUri := FilePathToUri(create_some_fake_file(t))
	_ = FindProjectRoot(filesUri)
	provider := ApplicationCategoryProvider()
	categories, error := provider.categories()
	if error != nil {
		t.Fatalf("Cannot get categories: %v", error)
	}

	assert.Contains(t, categories, Category{
		Name:      "holiday",
		Free:      true,
		PaidExtra: true,
	})

	assert.Contains(t, categories, Category{
		Name:      "overtime work",
		Free:      false,
		PaidExtra: true,
	})

	assert.Contains(t, categories, Category{
		Name:      "normal work",
		Free:      false,
		PaidExtra: false,
	})
}
