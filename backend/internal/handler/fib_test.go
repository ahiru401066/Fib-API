package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFibHandler(t *testing.T) {
	testcase := []struct {
		name       string
		method     string
		url        string
		wantStatus int
		wantBody   string 
	}{
		{
			name:       "valid input",
			method:     "GET",
			url:        "/fib?n=5",
			wantStatus: http.StatusOK,
			wantBody:   `"result":"5"`,
		},
		{
			name:       "invalid method",
			method:     "POST",
			url:        "/fib?n=5",
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"BadRequest"`,
		},
		{
			name:       "invalid query param",
			method:     "GET",
			url:        "/fib?n=abc",
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"BadRequest"`,
		},
		{
			name:       "negative number",
			method:     "GET",
			url:        "/fib?n=-2",
			wantStatus: http.StatusBadRequest,
			wantBody:   `"message":"BadRequest"`,
		},

	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.url, nil)
			w := httptest.NewRecorder()

			FibHandler(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tc.wantStatus {
				t.Errorf("status %dを期待したが、%dだった", tc.wantStatus, resp.StatusCode)
			}

			body := w.Body.String()
			if !strings.Contains(body, tc.wantBody) {
				t.Errorf("%sを期待したが、%sだった",tc.wantBody, body)
			}
		})
	}
}
