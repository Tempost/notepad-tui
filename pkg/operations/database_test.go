package database

import (
	"testing"
)

func TestCheckForDatabase(t *testing.T) {
	err := CheckForDatabase()
	if err != nil {
		t.Errorf("[ Error ] %v", err)
	}
}

func TestProjectNewNote(t *testing.T) {
	var u UserNotes
	if err := u.CreateNewNote("project", "New Work"); err != nil {
		t.Fail()
	}
}

func TestPersonalNote(t *testing.T) {
	var u UserNotes
	if err := u.CreateNewNote("personal", "Todo"); err != nil {
		t.Fail()
	}
}

func TestClassNote(t *testing.T) {
	var u UserNotes
	if err := u.CreateNewNote("class", "Compiler"); err != nil {
		t.Fail()
	}
}

func TestErrorNote(t *testing.T) {
	var u UserNotes
	if err := u.CreateNewNote("peal", "blank"); err != nil {
		t.Log(err)
	}
}

func TestSaveFunction(t *testing.T) {
    
}
