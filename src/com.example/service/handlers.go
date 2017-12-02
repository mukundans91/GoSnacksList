package service

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

type Result map[string]interface{}

func ListSnacksHandler(db *sql.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, ListSnacks(db))
	}

}

func SaveSnacksHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var snack Snack
		c.Bind(&snack)

		_, err := SaveSnack(db, snack)
		if err == nil {
			return c.JSON(http.StatusCreated, Result{"name": snack.Name, "quantity": snack.Quantity})
		} else {
			return err
		}
	}
}

func RemoveSnacksHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := DeleteSnack(db, c.Param("name"))
		if err == nil {
			return c.JSON(http.StatusOK, Result{"name": c.Param("name"), "quantity": 0})
		} else {
			return err
		}

	}
}

func ClearSnacksHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		del, err := DeleteAll(db)
		if err == nil {
			return c.JSON(http.StatusOK, "Cleared list of "+string(del)+" items")
		} else {
			return err
		}
	}
}
