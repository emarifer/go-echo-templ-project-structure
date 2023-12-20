package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/emarifer/go-templ-project-structure/services"
	"github.com/emarifer/go-templ-project-structure/views/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type UserService interface {
	GetAllUsers() ([]services.User, error)
	GetUserById(id int) (services.User, error)
}

func New(us UserService) *UserHandler {
	return &UserHandler{
		UserService: us,
	}
}

type UserHandler struct {
	UserService UserService
}

func (uh *UserHandler) HandlerShowUsers(c echo.Context) error {
	udata, err := uh.UserService.GetAllUsers()
	if err != nil {
		// fmt.Println(err)
		return err
	}

	si := user.ShowIndex("| Home", user.Show(udata))

	return uh.View(c, si)
}

func (uh *UserHandler) HandlerShowUserById(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}

	udata, err := uh.UserService.GetUserById(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return err
	}

	di := user.DetailsIndex(
		fmt.Sprintf(
			"| User details %s",
			cases.Title(language.English).String(udata.Username),
		),
		user.Details(tz, udata),
	)

	// return c.JSON(http.StatusOK, udata)
	return uh.View(c, di)
}

func (uh *UserHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
