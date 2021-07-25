package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *API) GetInvoice(ctx echo.Context, params GetInvoiceParams) error {
	uri := "http://localhost:8090/invoice_pdf?code=" + params.Code

	// Get the data
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the buffer // FIXME: proxy stream directly
	var b []byte = make([]byte, 10000000)
	nr, err := resp.Body.Read(b)

	// send buffer
	err = ctx.Blob(200, "application/pdf", b[:nr])
	if err != nil {
		return SendError(ctx, http.StatusNotFound, fmt.Sprintf("%s", err.Error()))
	}
	return nil
}
