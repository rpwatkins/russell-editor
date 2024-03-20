package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"russell_editor/models"
)

type AssessmentTypesResource struct {
	buffalo.Resource
}

func (a AssessmentTypesResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	assessmentTypeID := c.Param("assessment_type_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := models.LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var assessmentType models.AssessmentType
	for _, at := range course.AssessmentTypes {
		if at.ID == assessmentTypeID {
			assessmentType = at
		}
	}
	c.Set("assessmentType", assessmentType)
	return c.Render(http.StatusOK, r.HTML("assessment_types/edit.plush.html"))
}
