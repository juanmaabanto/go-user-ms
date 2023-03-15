package ports

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/juanmaabanto/go-user-ms/internal/app"
	"github.com/juanmaabanto/go-user-ms/internal/app/command"
	"github.com/juanmaabanto/go-user-ms/internal/app/query"
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

// ReadUser godoc
// @Summary Get a user by Id.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string  true  "User Id"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/users/{id} [get]
func (h HttpServer) ReadUser(c echo.Context) error {
	item := query.GetUserById{Id: c.Param("id")}

	result, err := h.app.Queries.GetUserById.Handle(c.Request().Context(), item)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, result)
}

// UpdateUser godoc
// @Summary Update a user.
// @Tags Users
// @Accept json
// @Produce json
// @Param command body command.UpdateUser true "Object to be modified."
// @Success 204
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/users [patch]
func (h HttpServer) UpdateUser(c echo.Context) error {
	item := command.UpdateUser{}

	if err := c.Bind(&item); err != nil {
		panic(err)
	}

	if err := c.Validate(item); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return c.JSON(http.StatusUnprocessableEntity, Simple(validationErrors))
	}

	err := h.app.Commands.UpdateUser.Handle(c.Request().Context(), item)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

// DeleteUser godoc
// @Summary Delete a user by Id.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string  true  "User Id"
// @Success 204
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/users/{id} [delete]
func (h HttpServer) DeleteUser(c echo.Context) error {
	item := command.DeleteUser{Id: c.Param("id")}

	err := h.app.Commands.DeleteUser.Handle(c.Request().Context(), item)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusNoContent, nil)
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
