package api

import (
	"fmt"
	"io"
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
		fmt.Printf("please set environment variable '%s'\n", INVOICE_PDF_URL)
	}
	uri := url + params.Code

	// Get the data
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(ctx.Response(), resp.Body)

	return nil
}
