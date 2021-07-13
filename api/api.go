package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// API is the api interface type
type API struct{}

var e *echo.Echo = echo.New()

func init() {
	fmt.Println("init")
	server := &API{}
	RegisterHandlers(e, server)
}

// GetTest returns status ok
func (a *API) GetTest(ctx echo.Context) error {
	if err := SendStatus(ctx, http.StatusOK, "OK"); err != nil {
		return err
	}
	return nil
}

// SendStatus function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func SendStatus(ctx echo.Context, code int, message string) error {
	statusError := Status{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, statusError)
	return err
}

// SendError function wraps sending of an error in the Error format
// returns sent error message or new error if sending fails
func SendError(ctx echo.Context, code int, message string) error {
	status := Status{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, status)
	if err != nil {
		return err
	}
	return fmt.Errorf(status.Message)
}
