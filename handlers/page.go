package handlers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Files   []os.DirEntry
	UrlPath string
}

type Error struct {
	Message string `json:"message"`
}

func Page(c echo.Context) error {
	urlPath := c.Request().URL.Path

	if urlPath == "/favicon.ico" {
		return c.NoContent(http.StatusNoContent)
	}
	//fmt.Println(urlPath)
	files, _ := os.ReadDir(urlPath)

	data := &Data{files, urlPath}
	return c.Render(http.StatusOK, "index.html", data)
}

func New(c echo.Context) error {

	name := c.QueryParam("name")
	isFolder, boolError := strconv.ParseBool(c.QueryParam("isFolder"))

	if boolError != nil {
		return c.JSON(http.StatusBadRequest, &Error{"invalid parameter"})
	}

	if isFolder {
		err := os.Mkdir(name, 0755)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{err.Error()})
		}
	} else {
		_, err := os.Create(name)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, &Error{err.Error()})
		}
	}
	return c.JSON(http.StatusOK, nil)
}

func Remove(c echo.Context) error {
	path := c.QueryParam("path")

	err := os.RemoveAll(path)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Error{err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}
