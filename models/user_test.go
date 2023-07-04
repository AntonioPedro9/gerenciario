package models

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	user := User{
		ID:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password123",
		CreatedAt: time.Now(),
	}

	if user.ID != 1 {
		t.Errorf("Expected ID to be 1, but got %d", user.ID)
	}

	if user.Name != "John Doe" {
		t.Errorf("Expected Name to be 'John Doe', but got '%s'", user.Name)
	}

	if user.Email != "john@example.com" {
		t.Errorf("Expected Email to be 'john@example.com', but got '%s'", user.Email)
	}

	if user.Password != "password123" {
		t.Errorf("Expected Password to be 'password123', but got '%s'", user.Password)
	}

	if user.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be a non-zero time")
	}
}
