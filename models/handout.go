package models

// Handout is a course handout
type Handout struct {
	ID          string
	Name        string
	URL         string
	Revision    string
	LastMod     string
	Description string
	Syllabus    string
}
