package provider
package provider

import (
	"testing"
)

func TestProvider(t *testing.T) {
	provider := New("test")()
	if provider == nil {
		t.Fatal("Expected provider to be non-nil")
	}
}
