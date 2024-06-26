package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"russell_editor/models"
)

type PartsResource struct {
	buffalo.Resource
}

func (p PartsResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	partID := c.Param("part_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := models.LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var part models.Part
	for _, pt := range course.Parts {
		if pt.ID == partID {
			part = pt
		}
	}
	c.Set("part", part)
	return c.Render(http.StatusOK, r.HTML("parts/edit.plush.html"))
}
