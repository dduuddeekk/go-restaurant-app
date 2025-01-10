package main

import (
	"fmt"
	"os"

	"github.com/dduuddeekk/go-restaurant-app/internal/database"
	"github.com/dduuddeekk/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/dduuddeekk/go-restaurant-app/internal/repository/menu"
	rRepo "github.com/dduuddeekk/go-restaurant-app/internal/repository/order"
	rUsecase "github.com/dduuddeekk/go-restaurant-app/internal/usecase/resto"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)
	menuRepo := mRepo.GetRepository(db)
	orderRepo := rRepo.GetRepository(db)
	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo)
	h := rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}
