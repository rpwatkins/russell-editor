package models

// AssessmentType is an assessment type
type AssessmentType struct {
	ID          string
	Name        string
	Assessments []Assessment
}
