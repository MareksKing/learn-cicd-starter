package auth

import (
	"net/http"
	"testing"
)


func TestGetApiKeyNoAuth(t *testing.T) {
	header	:= make(http.Header)
	header.Set("Key", "nothing")

	got, err := GetAPIKey(header)
	if got != "" {
		t.Fatal("Expected no return string")
	}
	if err != ErrNoAuthHeaderIncluded {
		t.Fatal("Error was expected")
	}
}

func TestGetApiKeyWrongAuthLength(t *testing.T) {
	header	:= make(http.Header)
	header.Set("Authorization", "short")

	got, err := GetAPIKey(header)
	if got != "" {
		t.Fatal("Expected no return string")
	}
	if err.Error() != "malformed authorization header" {
		t.Fatal("Error was expected")
	}
}

func TestGetApiKeyNoApiKeyProvided(t *testing.T) {
	header	:= make(http.Header)
	header.Set("Authorization", "aseotu : hasecuhasoetuh")

	got, err := GetAPIKey(header)
	if got != "" {
		t.Fatal("Expected no return string")
	}
	if err.Error() != "malformed authorization header" {
		t.Fatal("Error was expected")
	}
	
}

func TestGetApiKeyCorrectHeader(t *testing.T) {
	header	:= make(http.Header)
	header.Set("Authorization", "ApiKey super-secret-api-key")

	got, err := GetAPIKey(header)
	
	if got != "super-secret-api-key" {
		t.Fatal("Api key doesn't match")
	}
	if err != nil{
		t.Fatal("error should be nil")
	}
}
