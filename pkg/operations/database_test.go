package database

import (
	"testing"
)

func TestCheckForDatabase(t *testing.T) {
	_, err := CheckForDatabase()
	if err != nil {
		t.Errorf("[ Error ] %v", err)
	}
}
