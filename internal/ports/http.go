package ports

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/juanmaabanto/go-user-ms/internal/app"
	"github.com/juanmaabanto/go-user-ms/internal/app/command"
	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

// CreateUser godoc
// @Summary Create a new user.
// @Tags Users
// @Accept json
// @Produce json
// @Param command body command.CreateUser true "Object to be created."
// @Success 201 {string} string "Id of the created object"
// @Failure 400 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/users [post]
func (h HttpServer) CreateUser(c echo.Context) error {
	item := command.CreateUser{}

	if err := c.Bind(&item); err != nil {
		panic(err)
	}

	if err := c.Validate(item); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusUnprocessableEntity, Simple(validationErrors))
	}

	id, err := h.app.Commands.CreateUser.Handle(c.Request().Context(), item)

	if err != nil {
		panic(err)
	}

	c.Response().Header().Set("location", c.Request().URL.String()+"/"+id)

	return c.JSON(http.StatusCreated, id)
}

func Simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}
