package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbAddress string

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	dbAddress = "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=" + os.Getenv("DB_SSLMODE")
}

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

type MenuType string

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasigoreng",
			Price:     15000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Babi Guling",
			OrderCode: "babiguling",
			Price:     30000,
			Type:      MenuTypeFood,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name:      "Arak Bali",
			OrderCode: "arakbali",
			Price:     14000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Es Teh",
			OrderCode: "esteh",
			Price:     3000,
			Type:      MenuTypeDrink,
		},
	}

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&MenuItem{})

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}

func getMenu(c echo.Context) error {
	menuType := c.FormValue("type")

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func main() {
	seedDB()
	e := echo.New()
	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
