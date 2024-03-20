package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"russell_editor/models"
)

type AssessmentsResource struct {
	buffalo.Resource
}

func (a AssessmentsResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	assessmentID := c.Param("assessment_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := models.LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var assessment models.Assessment
	for _, at := range course.AssessmentTypes {
		for _, as := range at.Assessments {
			if as.ID == assessmentID {
				assessment = as
			}
		}
	}
	c.Set("assessment", assessment)
	return c.Render(http.StatusOK, r.HTML("assessments/edit.plush.html"))
}
