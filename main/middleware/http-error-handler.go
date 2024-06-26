package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

type HttpErrorHandler struct {
}

func NewHttpErrorHandler() *HttpErrorHandler {
	return &HttpErrorHandler{}
}

func (h HttpErrorHandler) HandleError(err error, c echo.Context) {
	log.Printf("Error: %v\n", err)
	if !c.Response().Committed {
		err := h.MapError(err, c)
		if err != nil {
			log.Printf("Error while sending response: %v\n", err)
		}
	} else {
		log.Printf("Response already committed\n")
	}
}

func (h HttpErrorHandler) MapError(err error, c echo.Context) error {
	switch err.Error() {
	case "pokemon not found":
		return c.JSON(404, ErrorMessage{"Pokemon not found"})
	case "error while fetching pokemon":
		return c.JSON(400, ErrorMessage{"Error while fetching pokemon"})
	case "error while deleting pokemon":
		return c.JSON(400, ErrorMessage{"Error while deleting pokemon"})
	case "pokemon already exists":
		return c.JSON(400, ErrorMessage{"Pokemon already exists"})
	case "user already exists":
		return c.JSON(400, ErrorMessage{"User already exists"})
	case "error while creating pokemon":
		return c.JSON(400, ErrorMessage{"Database constraint violated while creating pokemon"})
	default:
		return c.JSON(500, ErrorMessage{"Internal server error"})
	}
}
