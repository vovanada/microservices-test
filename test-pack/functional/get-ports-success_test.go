package functional

import (
	"fmt"
	"github.com/vovanada/microservices-test/shared/gateway/errors"
	"gopkg.in/resty.v1"
	"os"
	"testing"
)

func TestGetPorts_Success(t *testing.T) {

	testURL := os.Getenv(TestURL)

	if testURL == "" {
		t.Fatalf("env variable[%s] is empty", TestURL)
	}

	errResp := &errors.Error{}

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetError(&errResp).
		Get(fmt.Sprintf("%s/api/ports", testURL))

	if err != nil {
		t.Fatalf("failed to send request, %s", err)
	}

	if !resp.IsSuccess() {
		t.Fatalf("errors, %v", resp)
	}

	if errResp != nil && errResp.Message != "" {
		t.Fatalf("error response is not nil, %v", errResp)
	}
}
