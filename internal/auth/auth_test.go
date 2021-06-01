package auth

import (
	"strings"
	"testing"
)

func TestGetAuth(t *testing.T) {
	get := GetAuth()

	if strings.Contains(get, "not") {
		t.Errorf("expected not to find 'not', got %v", get)
	}
}
