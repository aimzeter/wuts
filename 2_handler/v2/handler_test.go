package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/aimzeter/wuts/2_handler/v2"
	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string

		wantCode int
		wantBody string
	}{
		{
			name:     "GET /participants/32710412345 should return 200",
			method:   http.MethodGet,
			path:     "/participants/32710412345",
			wantCode: 200,
			wantBody: "Showing participant with NIK 32710412345",
		},
		{
			name:     "POST /participants/32710412345 should return 405",
			method:   http.MethodPost,
			path:     "/participants/32710412345",
			wantCode: 405,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			srv := handler.NewServer()
			r := httptest.NewRequest(tc.method, tc.path, nil)
			w := httptest.NewRecorder()

			// When
			srv.ServeHTTP(w, r)

			// Then
			assert.Equal(t, tc.wantCode, w.Code)
			if tc.wantBody != "" {
				assert.Equal(t, tc.wantBody, w.Body.String())
			}
		})
	}
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name    string
		method  string
		path    string
		payload string

		wantCode int
		wantBody string
	}{
		{
			name:   "POST /participants with valid payload should return 201",
			method: http.MethodPost,
			path:   "/participants",
			payload: `
				{
					"nik": "32710412345",
					"agree": true
				}
			`,
			wantCode: 201,
			wantBody: "Participant with NIK 32710412345 successfully registered",
		},
		{
			name:   "GET /participants should return 405",
			method: http.MethodGet,
			path:   "/participants",
			payload: `
				{
					"nik": "32710412345",
					"agree": true
				}
			`,
			wantCode: 405,
		},
		{
			name:   "POST /participants with invalid payload should return 400",
			method: http.MethodPost,
			path:   "/participants",
			payload: `
				{
					"nik": "32710412345"
					"agree": true
				}
			`,
			wantCode: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			srv := handler.NewServer()
			r := httptest.NewRequest(tc.method, tc.path, bytes.NewBufferString(tc.payload))
			w := httptest.NewRecorder()

			// When
			srv.ServeHTTP(w, r)

			// Then
			assert.Equal(t, tc.wantCode, w.Code)
			if tc.wantBody != "" {
				assert.Equal(t, tc.wantBody, w.Body.String())
			}
		})
	}
}
