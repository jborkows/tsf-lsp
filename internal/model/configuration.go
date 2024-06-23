package model

import (
	"errors"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type ProjectRootDirectory = string

const configuration_file_name = "configuration.yaml"

var project_root = ProjectRootDirectory("")

type Config struct {
	Categories []Category `yaml:"categories"`
}

type CustomBool bool

// UnmarshalYAML implements the yaml.Unmarshaler interface for CustomBool.
func (cb *CustomBool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var text string
	if err := unmarshal(&text); err != nil {
		return err
	}
	text = strings.ToLower(text)
	switch text {
	case "yes":
		*cb = CustomBool(true)
	case "no":
		*cb = CustomBool(false)
	default:
		return errors.New("invalid value for boolean")
	}
	return nil
}

type Category struct {
	Name      string     `yaml:"name"`
	Free      CustomBool `yaml:"free"`
	PaidExtra CustomBool `yaml:"extraPaid"`
}

type CategoryProvider interface {
	categories() ([]Category, error)
}

var categoryProvider *CategoryProvider = nil

type fileBasedCategoryProvider struct {
	categoryValues []Category
}

func (self *fileBasedCategoryProvider) categories() ([]Category, error) {
	return self.categoryValues, nil
}

func readConfig() *Config {

	if project_root == "" {
		panic("root directory was not initialized, please open some timesheet file")
	}
	data, err := os.ReadFile(filepath.Join(project_root, configuration_file_name))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	var configuration Config
	withoutTabs := strings.ReplaceAll(string(data), "\t", " ")
	err = yaml.Unmarshal([]byte(withoutTabs), &configuration)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &configuration
}

func ApplicationCategoryProvider() CategoryProvider {
	if categoryProvider != nil {
		return *categoryProvider
	}

	config := readConfig()
	provider := fileBasedCategoryProvider{categoryValues: config.Categories}
	return &provider
}

func FilePathToUri(filePath string) string {
	filePath = filepath.ToSlash(filePath)
	fileURL := url.URL{
		Scheme: "file",
		Path:   filePath,
	}
	fileURI := fileURL.String()
	return fileURI
}

func FindProjectRoot(uri string) ProjectRootDirectory {
	if project_root != "" {
		return project_root
	}
	parsedUri, err := url.Parse(uri)
	if err != nil {
		log.Fatalf("Cannot parse URI: %s -> %v", uri, err)
		return ""
	}
	file := parsedUri.Path
	project_root = findRootConfigFile(filepath.Dir(file), 5)
	return project_root

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return !info.IsDir()
}

func findRootConfigFile(currentPath string, numberOfRetries int) ProjectRootDirectory {

	if fileExists(filepath.Join(currentPath, configuration_file_name)) {
		return currentPath
	}

	if numberOfRetries < 1 {
		log.Fatalf("Cannot find %s, please define it at project root", currentPath)
		return ""
	}

	return findRootConfigFile(filepath.Join(currentPath, ".."), numberOfRetries-1)
}
