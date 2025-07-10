// https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		name        string
		headerKey   string
		headerValue string
		want        string
		wantError   bool
	}{
		"ok":               {headerKey: "Authorization", headerValue: "ApiKey 123456789", want: "123456789_", wantError: false},
		"malformed header": {headerKey: "Authorization", headerValue: "Token 123456789", want: "", wantError: true},
		"no auth header":   {headerKey: "Nope", headerValue: "", want: "", wantError: true},
	}

	for name, tc := range tests {
		header := http.Header{}
		header.Set(tc.headerKey, tc.headerValue)

		got, err := GetAPIKey(header)

		if tc.wantError {
			if err == nil {
				t.Fatalf("[%s] expected error but got none", name)
			}
		} else {
			if err != nil {
				t.Fatalf("[%s] error: %v", name, err.Error())
			}
		}

		if got != tc.want {
			t.Fatalf("[%s] expected: %v, got: %v", name, tc.want, got)
		}

	}
}
