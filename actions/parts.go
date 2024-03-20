package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
)

type PartsResource struct {
	buffalo.Resource
}

func (p PartsResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	partID := c.Param("part_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var part Part
	for _, p := range course.Parts {
		if p.ID == partID {
			part = p
		}
	}
	c.Set("part", part)
	return c.Render(http.StatusOK, r.HTML("parts/edit.plush.html"))
}
