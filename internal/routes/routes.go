package routes

import (
	"deneme/internal/store/author"
	"deneme/internal/store/blog"
	"github.com/labstack/echo/v4"
)

var (
	AuthorStore *author.Store
	BlogStore   *blog.Store
)

func RegisterRoutes(e *echo.Echo) {
	AuthorStore = author.CreateAuthorStore()
	BlogStore = blog.CreateBlogStore()

	e.POST("/author", CreateAuthor)
	e.GET("/author-list", ListAuthor)

	e.POST("/blog-post", CreateBlogPost)
	e.GET("/post-list", ListBlogPost)
	e.GET("/blog-post/:id", GetBlogPost)
}
