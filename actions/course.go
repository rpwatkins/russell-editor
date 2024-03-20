package actions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

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
func (c *Course) Write(cwd string, course string) error {

	// marshall course to toml
	b, err := toml.Marshal(c)
	if err != nil {
		return err
	}

	// get file path
	fileName := fmt.Sprintf("%s.toml", course)
	name := filepath.Join(cwd, "data", "courses", fileName)

	// write file to path
	if err := os.WriteFile(name, b, 0666); err != nil {
		return err
	}
	return nil
}

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
	Parts             []Part
	AssessmentTypes   []AssessmentType
	Readings          []Reading
	Handouts          []Handout
	Texts             []Text
	Policies          []Policy
	Semester          string
}

type Courses []Course

// Part is a course part
type Part struct {
	Name  string
	Weeks []Week
}

// AssessmentType is an assessment type
type AssessmentType struct {
	Name        string
	Assessments []Assessment
}

// Week contains data for a week in a course
type Week struct {
	Name        string
	BeginDate   string
	EndDate     string
	Prefix      string
	Description string
	Assessment  string
	Filesets    []Fileset
	Readings    []Reading
}

// Fileset contains the files for a week
type Fileset struct {
	Pres        string
	Name        string
	Revision    string
	LastMod     string
	Description string
}

// Reading is a course reading
type Reading struct {
	URL         string
	Authors     string
	Name        string
	Pages       string
	Description string
}

// Handout is a course handout
type Handout struct {
	Name        string
	URL         string
	Revision    string
	LastMod     string
	Description string
	Syllabus    string
}

// Assessment is an assessment
type Assessment struct {
	Name    string
	DueDate string
	Percent string
	Pdf     string
}

// Text contains data on textbooks
type Text struct {
	Name      string
	Authors   string
	Year      string
	Edition   string
	Publisher string
	URL       string
}

// Policy contains data for policies
type Policy struct {
	Name     string
	Content  string
	LastMod  string
	Revision string
}
