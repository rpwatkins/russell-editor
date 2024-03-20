package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"russell_editor/models"
)

type FilesetsResource struct {
	buffalo.Resource
}

func (f FilesetsResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	filesetID := c.Param("fileset_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := models.LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var fileset models.Fileset
	for _, p := range course.Parts {
		for _, w := range p.Weeks {
			for _, fs := range w.Filesets {
				if fs.ID == filesetID {
					fileset = fs
				}
			}
		}
	}
	c.Set("fileset", fileset)
	return c.Render(http.StatusOK, r.HTML("filesets/edit.plush.html"))
}
