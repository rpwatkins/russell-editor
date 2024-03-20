package actions

import (
	"strings"

	"github.com/gobuffalo/buffalo"
)

// SetActiveMenu attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetActiveMenu(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		// returns the
		activeMenuItem := func() string {
			routeInfo := c.Value("current_route").(buffalo.RouteInfo)
			route := getResourceName(routeInfo)

			switch route {

			case "HomeHandler":
				return "home"
			case "CoursesResource":
				courseID := c.Param("course_id")
				return courseID
			case "PartsResource":
				courseID := c.Param("course_id")
				return courseID

			default:
				return ""
			}
		}

		// set default active menu note (this can be overridden in route handlers)
		a := activeMenuItem()
		c.Set("activeMenuItem", a)

		return next(c)
	}
}

// getResourceName gets a resource's name from the request routeInfo struct
func getResourceName(route buffalo.RouteInfo) string {

	pts := strings.Split(route.HandlerName, ".")
	res := pts[1]
	return res

}
