package main

import "testing"

func TestGenerateID(t *testing.T) {
	id, err := GenerateID()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(id) != 8 {
		t.Errorf("expected id length 8, got %d", err)
	}
}
