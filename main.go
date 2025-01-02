package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasigoreng",
			Price:     15000,
		},
		{
			Name:      "Babi Guling",
			OrderCode: "babiguling",
			Price:     30000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getDrinkMenu(c echo.Context) error {
	drinkMenu := []MenuItem{
		{
			Name:      "Arak Bali",
			OrderCode: "arakbali",
			Price:     14000,
		},
		{
			Name:      "Es Teh",
			OrderCode: "esteh",
			Price:     3000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}

func main() {
	e := echo.New()
	// localhost:14045/menu/food
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
