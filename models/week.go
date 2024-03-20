package models

// Week contains data for a week in a course
type Week struct {
	ID          string
	Name        string
	BeginDate   string
	EndDate     string
	Prefix      string
	Description string
	Assessment  string
	Filesets    []Fileset
	Readings    []Reading
}
