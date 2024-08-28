package app

import (
	"aws-lambda/api"
	"aws-lambda/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	db := database.NewDynamoDBClient()
	ApiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler,
	}
}
