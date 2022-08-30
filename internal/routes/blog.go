package routes

import (
	"deneme/internal/model"
	"github.com/labstack/echo/v4"
)

func CreateBlogPost(c echo.Context) error {
	req := new(model.CreateBlogPost)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(400, err)
	}

	author := AuthorStore.GetAuthor(req.AuthorName)

	post := &model.BlogPost{
		Title: req.Title,
		Text:  req.Text,
		Author: struct {
			Name string `json:"name"`
			Bio  string `json:"bio"`
		}{Name: author.Name, Bio: author.Bio},
		Photos: req.Photos,
	}

	// BlogPost Store Insert Data
	BlogStore.CreateBlogPost(*post)

	return c.JSON(200, req)
}

func ListBlogPost(c echo.Context) error {

	// Blog Store Get All Data
	resp := BlogStore.ListBlogPost()

	return c.JSON(200, resp)
}

func GetBlogPost(c echo.Context) error {
	id := c.Param("id")

	resp := BlogStore.GetBlogPost(id)

	return c.JSON(200, resp)
}
