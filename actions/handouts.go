package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"russell_editor/models"
)

type HandoutsResource struct {
	buffalo.Resource
}

func (h HandoutsResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	handoutID := c.Param("handout_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := models.LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var handout models.Handout
	for _, h := range course.Handouts {
		if p.ID == partID {
			part = p
		}
	}
	c.Set("part", part)
	return c.Render(http.StatusOK, r.HTML("parts/edit.plush.html"))
}
