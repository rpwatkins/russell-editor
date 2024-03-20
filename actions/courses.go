package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
)

type CoursesResource struct {
	buffalo.Resource
}

func (cr CoursesResource) Show(c buffalo.Context) error {
	project := c.Param("course_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	return c.Render(http.StatusOK, r.HTML("courses/index.plush.html"))
}

func (cr CoursesResource) Edit(c buffalo.Context) error {

	return c.Render(http.StatusOK, r.HTML("courses/edit.plush.html"))

}
