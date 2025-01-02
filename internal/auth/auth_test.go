package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name string
		headerName string
		headerValue string
		result string
		er string
	} {
		{ name: "normal header", headerName: "Authorization", headerValue: "ApiKey jwt", result: "jwt" },
		{ name: "invalid header name", headerName: "Authoriation", headerValue: "ApiKey jwt" },
		{ name: "invalid header value - wrong leading key", headerName: "Authorization", headerValue: "Bearer jwt" },
		{ name: "invalid header value - no leading key", headerName: "Authorization", headerValue: "jwt", result: "" },
	}

	for _, tc := range tests {
		header := http.Header{}
		header.Set(tc.headerName, tc.headerValue)
	
		got, _ := GetAPIKey(header)
	
		want := tc.result
	
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, want, got)
		}
	}
}