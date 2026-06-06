package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type dummyData struct {
	Foo string `json:"foo"`
}

func TestSendSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	data := dummyData{Foo: "bar"}

	SendSuccess(w, http.StatusCreated, data)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status 201, got %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var resp Response
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if resp.Error != nil {
		t.Errorf("expected error to be nil, got %+v", resp.Error)
	}

	// Because resp.Data is interface{}, let's map inspect it
	innerMap, ok := resp.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("expected data to be map, got %T", resp.Data)
	}
	if innerMap["foo"] != "bar" {
		t.Errorf("expected foo to be 'bar', got %v", innerMap["foo"])
	}
}

func TestSendError(t *testing.T) {
	w := httptest.NewRecorder()

	SendError(w, http.StatusBadRequest, "invalid_request", "Something was invalid")

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var resp Response
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if resp.Data != nil {
		t.Errorf("expected data to be nil, got %+v", resp.Data)
	}

	if resp.Error == nil {
		t.Fatal("expected error to be non-nil")
	}

	if resp.Error.Code != "invalid_request" {
		t.Errorf("expected error code 'invalid_request', got '%s'", resp.Error.Code)
	}

	if resp.Error.Message != "Something was invalid" {
		t.Errorf("expected error message 'Something was invalid', got '%s'", resp.Error.Message)
	}
}
