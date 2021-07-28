package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestGetInvoice(t *testing.T) {
	receiptCode := "99999120210727120121"

	result := testutil.NewRequest().Get(
		fmt.Sprintf("/invoice?code=%s", receiptCode)).WithAcceptJson().Go(t, e)
	if assert.Equal(t, http.StatusOK, result.Code()) {
		want := "%PDF"
		got := fmt.Sprintf("%s", result.Recorder.Body.String())
		assert.Equal(t, want, got[:4])
	}
}
