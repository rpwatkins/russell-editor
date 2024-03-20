package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"russell_editor/models"
)

type PoliciesResource struct {
	buffalo.Resource
}

func (p PoliciesResource) Edit(c buffalo.Context) error {
	project := c.Param("course_id")
	policyID := c.Param("policy_id")
	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	course, err := models.LoadCourse(folder, project)
	if err != nil {
		return err
	}
	c.Set("course", course)

	var policy models.Policy
	for _, pl := range course.Policies {
		if pl.ID == policyID {
			policy = pl
		}
	}
	c.Set("policy", policy)
	return c.Render(http.StatusOK, r.HTML("policies/edit.plush.html"))
}
