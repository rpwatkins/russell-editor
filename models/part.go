package models

// Part is a course part
type Part struct {
	ID    string
	Name  string
	Weeks []Week
}
