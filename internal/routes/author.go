package routes

import (
	"deneme/internal/model"
	"github.com/labstack/echo/v4"
)

func CreateAuthor(c echo.Context) error {
	req := new(model.Author)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(400, err)
	}

	// Author Store Insert Data
	AuthorStore.CreateAuthor(*req)

	return c.JSON(200, req)
}

func ListAuthor(c echo.Context) error {

	// Author Store Get All Data
	resp := AuthorStore.ListAuthor()

	return c.JSON(200, resp)
}
