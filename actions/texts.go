package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"russell_editor/models"
)

type TextsResource struct {
	buffalo.Resource
}

func (t TextsResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	textID := c.Param("text_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := models.LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var text models.Text
	for _, tx := range course.Texts {
		if tx.ID == textID {
			text = tx
		}
	}
	c.Set("text", text)
	return c.Render(http.StatusOK, r.HTML("texts/edit.plush.html"))
}
