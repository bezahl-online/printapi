package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
)

func (a *API) Print(ctx echo.Context) error {
	var request PrintJSONRequestBody
	var err error
	if err = ctx.Bind(&request); err != nil {
		return err
	}
	print(request.PdfUrl)
	if err != nil {
		return SendError(ctx, http.StatusNotFound, fmt.Sprintf("%s", err.Error()))
	}
	return SendStatus(ctx, http.StatusOK, "OK")
}

var URL = "api.mms.bezahl.online/invoicePDF?"

func print(URL string) error {
	filePath := "test.pdf"
	err := downloadFile(filePath, URL)
	if err != nil {
		return err
	}
	cmd := exec.Command("lp", "-o", "media=A4", "-d", "DCPL3550CDW", filePath)
	s, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	cmd.Start()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(s)
	if err != nil {
		return err
	}
	fmt.Print(buf.String())
	cmd.Wait()
	return nil
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
