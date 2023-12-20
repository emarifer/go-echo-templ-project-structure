package main

import (
	"net/http"

	"github.com/emarifer/go-templ-project-structure/db"
	"github.com/emarifer/go-templ-project-structure/handlers"
	"github.com/emarifer/go-templ-project-structure/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// In production, the name of the database
// would be obtained from an .env file
const dbName = "user_data.db"

func main() {
	app := echo.New()

	app.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	app.Static("/", "assets")
	app.Use(middleware.Logger())

	// We redirect the root route to the "/user" route
	app.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/user")
	})

	uStore, err := db.NewUserStore(dbName)
	if err != nil {
		app.Logger.Fatalf("failed to create store: %s", err)
	}

	us := services.NewServicesUser(services.User{}, uStore)

	h := handlers.New(us)

	handlers.SetupRoutes(app, h)

	app.Logger.Fatal(app.Start(":5000"))
}

/*
https://www.reddit.com/r/golang/comments/17d12wk/using_echo_with_ahtempl/

https://templ.guide/project-structure/project-structure
https://github.com/a-h/templ/tree/main/examples/counter

https://echo.labstack.com/docs

https://www.reddit.com/r/golang/comments/vggekd/whats_the_usage_of_error_in_echo_framework/
https://echo.labstack.com/docs/error-handling
https://medium.com/@emretiryaki3/custom-error-handling-for-echo-web-framework-go-152992ab9cce
https://sorcererxw.com/en/articles/go-echo-error-handing
*/
