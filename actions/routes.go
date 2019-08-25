package actions

import (
	"github.com/gobuffalo/buffalo"
)

func routes(app *buffalo.App) {
		app.GET("/", HomeHandler)
		app.POST("/charges", CreateChargeHandler)
}