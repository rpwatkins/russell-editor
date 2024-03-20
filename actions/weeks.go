package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"russell_editor/models"
)

type WeeksResource struct {
	buffalo.Resource
}

func (w WeeksResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	weekID := c.Param("week_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := models.LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var week models.Week
	for _, pt := range course.Parts {
		for _, wk := range pt.Weeks {
			if wk.ID == weekID {
				week = wk
			}
		}

	}
	c.Set("week", week)
	return c.Render(http.StatusOK, r.HTML("weeks/edit.plush.html"))
}
