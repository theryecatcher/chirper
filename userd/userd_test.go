package userd

import (
	"testing"
)

func TestUserD_Create(t *testing.T) {
	_, err := New(&Config{})
	if err != nil {
		t.Fatalf("Unexpected error {%+v}", err)
	}
}
