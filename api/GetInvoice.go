package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

const INVOICE_PDF_URL = "INVOICE_PDF_URL"

func (a *API) GetInvoice(ctx echo.Context, params GetInvoiceParams) error {
	url := "http://localhost:8090/invoice_pdf?code="
	if len(os.Getenv(INVOICE_PDF_URL)) > 0 {
		url = os.Getenv(INVOICE_PDF_URL)
	} else {
		fmt.Printf("please set environment variable '%s'", INVOICE_PDF_URL)
	}
	uri := url + params.Code

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
