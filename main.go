package main

import (
	"github.com/MarkTBSS/go-kbtg-challenge_8/configs"
	"github.com/MarkTBSS/go-kbtg-challenge_8/postgres"
	"github.com/MarkTBSS/go-kbtg-challenge_8/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/MarkTBSS/go-kbtg-challenge_8/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	configs.LoadConfig()
	databaseInstance, err := postgres.New()
	if err != nil {
		panic(err)
	}
	echoInstance := echo.New()
	echoInstance.GET("/swagger/*", echoSwagger.WrapHandler)
	handler := wallet.New(databaseInstance)
	echoInstance.GET("/api/v1/wallets", handler.WalletHandler)
	echoInstance.GET("/api/v1/wallets/query", handler.QueryParamHandler)
	echoInstance.GET("/api/v1/users/:id/wallets", handler.PathParamHandler)
	echoInstance.POST("/api/v1/wallets", handler.BindingDataHandler)
	echoInstance.PUT("/api/v1/wallets", handler.UpdateHandler)
	echoInstance.DELETE("/api/v1/users/:id/wallets", handler.DeleteHandler)
	echoInstance.Logger.Fatal(echoInstance.Start(":1323"))
}
