package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

// Course is the struct that holds toml data
type Course struct {
	Project           string
	CourseTitle       string
	CourseNumber      string
	CourseCode        string
	CourseDescription string
	Revision          string
	LastModified      string
	SyllabusURL       string
	Semester          string
	Parts             []Part
	AssessmentTypes   []AssessmentType
	Readings          []Reading
	Handouts          []Handout
	Texts             []Text
	Policies          []Policy
}

type Courses []Course

// LoadCourse loads course toml data into a Course struct
func LoadCourse(dataLocation, project string) (Course, error) {

	var course Course

	// file
	fileName := fmt.Sprintf("%s.toml", project)
	path := filepath.Join(dataLocation, "data", "courses", fileName)

	// load bytes
	tomlBytes, err := os.ReadFile(path)
	if err != nil {
		return course, err
	}

	// unmarshall toml data to course struct
	if err := toml.Unmarshal([]byte(tomlBytes), &course); err != nil {
		return course, err
	}

	// return course toml
	return course, nil
}

func LoadCourses(dataLocation string) (Courses, error) {

	var courses Courses
	courseList, err := ListCourses(dataLocation)
	if err != nil {
		return courses, err
	}

	for _, c := range courseList {
		course, err := LoadCourse(dataLocation, c)
		if err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func ListCourses(cwd string) ([]string, error) {
	var list []string

	dataPath := filepath.Join(cwd, "data", "courses")
	dirList, err := os.ReadDir(dataPath)
	if err != nil {
		return list, err
	}

	for _, f := range dirList {
		if !f.IsDir() {
			p := strings.Replace(f.Name(), ".toml", "", 1)
			list = append(list, p)
		}
	}
	return list, nil
}

// Write writes the Course struct to toml data file
func (c *Course) Write(dataLocation string, course string) error {

	// marshall course to toml
	b, err := toml.Marshal(c)
	if err != nil {
		return err
	}

	// get file path
	fileName := fmt.Sprintf("%s.toml", course)
	name := filepath.Join(dataLocation, "data", "courses", fileName)

	// write file to path
	if err := os.WriteFile(name, b, 0666); err != nil {
		return err
	}
	return nil
}
