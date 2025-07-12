package handlers

import (
	"filemanager/internal/filemanager"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Files []os.DirEntry
}

func Page(c echo.Context) error {
	urlPath := c.Request().URL.Path

	if urlPath == "/favicon.ico" {
		return c.NoContent(http.StatusNoContent)
	}

	files, _ := filemanager.List(urlPath)

	data := &Data{files}
	return c.Render(http.StatusOK, "index.html", data)
}
