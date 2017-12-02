package main

import (
	"net/http"

	"com.example/dbcom"
	"com.example/service"
	"github.com/labstack/echo"
)

func main() {
	db := dbcom.InitDB("snacks.db")
	dbcom.Create(db)
	e := echo.New()
	e.GET("/", getRoot)
	e.GET("/snacks", service.ListSnacksHandler(db))
	e.PUT("/snacks", service.SaveSnacksHandler(db))
	e.DELETE("/snacks/:name", service.RemoveSnacksHandler(db))
	e.DELETE("/snacks", service.ClearSnacksHandler(db))
	e.Logger.Fatal(e.Start(":8080"))
}

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Go!!")
}
