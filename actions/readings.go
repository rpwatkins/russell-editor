package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"russell_editor/models"
)

type ReadingsResource struct {
	buffalo.Resource
}

func (rd ReadingsResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	readingID := c.Param("reading_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := models.LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var reading models.Reading
	for _, p := range course.Parts {
		for _, w := range p.Weeks {
			for _, rdg := range w.Readings {
				if rdg.ID == readingID {
					reading = rdg
				}
			}
		}
	}
	c.Set("reading", reading)
	return c.Render(http.StatusOK, r.HTML("readings/edit.plush.html"))
}
