package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/dduuddeekk/go-restaurant-app/internal/database"
	"github.com/dduuddeekk/go-restaurant-app/internal/delivery/rest"
	"github.com/dduuddeekk/go-restaurant-app/internal/logger"
	mRepo "github.com/dduuddeekk/go-restaurant-app/internal/repository/menu"
	oRepo "github.com/dduuddeekk/go-restaurant-app/internal/repository/order"
	uRepo "github.com/dduuddeekk/go-restaurant-app/internal/repository/user"
	"github.com/dduuddeekk/go-restaurant-app/internal/tracing"
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
	logger.Init()
	url := os.Getenv("TRACE_URL")
	tracing.Init(url)
	e := echo.New()

	db := database.GetDB(dbAddress)
	secret := os.Getenv("SECRET")
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Second)
	if err != nil {
		panic(err)
	}

	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}
