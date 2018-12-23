package functional

import (
	"bytes"
	"fmt"
	"github.com/vovanada/microservices-test/shared/gateway/errors"
	"github.com/vovanada/microservices-test/test-pack/fixtures"
	"gopkg.in/resty.v1"
	"os"
	"testing"
)

func TestImport_InvalidFile(t *testing.T) {

	testURL := os.Getenv(TestURL)

	if testURL == "" {
		t.Fatalf("env variable[%s] is empty", TestURL)
	}

	errResp := &errors.Error{}

	fileReader := bytes.NewReader(fixtures.ImportInvalidFile)

	resp, err := resty.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFileReader("file", "file.txt", fileReader).
		SetError(&errResp).
		Post(fmt.Sprintf("%s/api/ports/import", testURL))

	if err != nil {
		t.Fatalf("failed to send request, %s", err)
	}

	if !resp.IsError() {
		t.Fatalf("response without err, %v", resp)
	}

	if errResp == nil || errResp.Message != "Invalid request" {
		t.Fatalf("error response is not equal, %v", errResp)
	}
}
