package api

import (
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestPrint(t *testing.T) {
	testurl := "http://localhost:8090/invoice_pdf?code=99999120210712170257"
	testurl = "https://www.orimi.com/pdf-test.pdf"
	request := PrintJSONRequestBody{
		PdfUrl: testurl,
	}
	result := testutil.NewRequest().Post("/print").WithJsonBody(request).WithAcceptJson().Go(t, e)
	if assert.Equal(t, http.StatusOK, result.Code()) {

	}
}
