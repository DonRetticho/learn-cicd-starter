package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name        string
		headerKey   string
		headerValue string
		expectedKey string
		expectedErr string
	}{
		{name: "empty header", headerKey: "", headerValue: "", expectedKey: "", expectedErr: "no authorization header included"},
		{name: "malformed header", headerKey: "Authorization", headerValue: "Bearer 12345", expectedKey: "", expectedErr: "malformed authorization header"},
		{name: "successfull header", headerKey: "Authorization", headerValue: "ApiKey 12345", expectedKey: "12345", expectedErr: ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}
			if tc.headerKey != "" {
				headers.Add(tc.headerKey, tc.headerValue)
			}

			got, err := GetAPIKey(headers)

			if tc.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error containing %q, but got nil", tc.expectedErr)
					return
				}
				if !strings.Contains(err.Error(), tc.expectedErr) {
					t.Errorf("expected error %q, bot got %q", tc.expectedErr, err.Error())
					return
				}
				return
			}
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
				return
			}
			if got != tc.expectedKey {
				t.Errorf("expected key %q, but got %q", tc.expectedKey, got)
			}
		})
	}
}
