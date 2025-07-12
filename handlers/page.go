package handlers

import (
	"filemanager/internal/filemanager"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Files   []os.DirEntry
	UrlPath string
}

func Page(c echo.Context) error {
	urlPath := c.Request().URL.Path

	if urlPath == "/favicon.ico" {
		return c.NoContent(http.StatusNoContent)
	}
	fmt.Println(urlPath)
	files, _ := filemanager.List(urlPath)

	data := &Data{files, urlPath}
	return c.Render(http.StatusOK, "index.html", data)
}
