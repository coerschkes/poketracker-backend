package main

import (
	"github.com/labstack/echo/v4"
	"poketracker-backend/main/external"
)

func main() {
	e := echo.New()
	fb := external.NewFirebaseAuthenticator()
	e.POST("/validate", fb.Validate)

	e.Logger.Fatal(e.Start(":1323"))
}
