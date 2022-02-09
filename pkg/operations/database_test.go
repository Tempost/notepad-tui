package database

import (
	"os"
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
	if _, err := u.CreateNewNote("project", "New Work"); err != nil {
		t.Fail()
	}
}

func TestPersonalNote(t *testing.T) {
	var u UserNotes
	if _, err := u.CreateNewNote("personal", "Todo"); err != nil {
		t.Fail()
	}
}

func TestClassNote(t *testing.T) {
	var u UserNotes
	if _, err := u.CreateNewNote("class", "Compiler"); err != nil {
		t.Fail()
	}
}

func TestErrorNote(t *testing.T) {
	var u UserNotes
	if _, err := u.CreateNewNote("peal", "blank"); err != nil {
		t.Log(err)
	}
}

func TestNewUserCreated(t *testing.T) {
	filepath := "../../data/" + os.Getenv("USER") + ".db"
	if err := OpenDatabase(filepath); err != nil {
		t.Fail()
	}

	var user UserNotes
	db.First(&user)

	if user.UserName != "tempost" {
		t.Fail()
	}

	temp, _ := db.DB()

	defer temp.Close()
}
