package actions

import (
	"fmt"
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/plush/v4"
	"russell_editor/models"
)

// plush templates custom helpers
func addPlushHelpers() {

	if err := plush.Helpers.Add("courseMenu", courseMenu); err != nil {
		fmt.Print("error loading plush helper")
	}

}

func courseMenu(courses models.Courses) models.Courses {

	folder := envy.Get("RUSSELL_SITE_LOCATION", "")

	courses, err := models.LoadCourses(folder)
	if err != nil {
		log.Fatal(err)
	}
	return courses
}
