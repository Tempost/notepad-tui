package database

// this package will cover all/majority of the database functions
// consider storting types/interfaces for the database in a seperate file

// when opening different note types load them into memory
// write any changed to the DB upon leaving the note types menu

// when editing a note
//  -- Pull markdown file from database
//  -- Create a temp file in edits directory
//  -- Push all files in edits directory to database when user saves notes from TUI

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// global variable to avoid passing in a db handle to every function that works on the database
var db *gorm.DB

// On runtime of the application check for a database and it will open
// if no database is found this will create a new database, populate with defaults and open
// The defaults are just one blank note for each note type
func CheckForDatabase() error {
	filename := os.Getenv("USER") + ".db"
	// Once we get to were we can "ship" a binary determine where the data should go
	filepath := "../../data/" + filename

	// check if our file exists, otherwise skip code inside if scope and open the database
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		os.Create(filepath)

		err := OpenDatabase(filepath)
		if err != nil {
			return err
		}

		db.AutoMigrate(
			&UserNotes{},
			&ProjectNote{},
			&PersonalNote{},
			&ClassNote{},
			&Scratchpad{},
		)
	}

	return OpenDatabase(filepath)
}

func OpenDatabase(filepath string) error {
	var err error
	db, err = gorm.Open(sqlite.Open(filepath))
	if err != nil {
		return fmt.Errorf("[ ERROR ]: %v", err)
	}
	return nil
}

// Create new markdownfile -> projectname_project.md
// Push to edits directory
// Handle the case where a Note has the same name as one in the database or in the yet to be
// saved notes
func (p *ProjectNote) CreateNewProjectNote(name string) {
	note, err := GenNewFile(name)
	if err != nil {
		log.Fatalln("[ ERROR ] Could not create file: ", err)
	}
	p.note = note
}

// Create new markdownfile -> projectname_project.md
func (p *PersonalNote) CreateNewPersonalNote(name string) {
	note, err := GenNewFile(name)
	if err != nil {
		log.Fatalln("[ ERROR ] Could not create file: ", err)
	}
	p.note = note
}

// Create new markdownfile -> projectname_project.md
func (c *ClassNote) CreateNewClassNote(name string) {
	note, err := GenNewFile(name)
	if err != nil {
		log.Fatalln("[ ERROR ] Could not create file: ", err)
	}
	c.note = note
}

func GenNewFile(name string) (*os.File, error) {
	filepath := "../../edits/" + name + ".md"
	return os.Create(filepath)
}

// when go generics is released in stable consider changing this up
func (u *UserNotes) CreateNewNote(noteType string, name string) error {
	noteType = strings.ToLower(noteType)

	switch noteType {
	case "project":
		{
			var p ProjectNote
			p.CreateNewProjectNote(name)
		}
	case "personal":
		{
			var p PersonalNote
			p.CreateNewPersonalNote(name)
		}
	case "class":
		{
			var c ClassNote
			c.CreateNewClassNote(name)
		}
	default:
		{
			err := errors.New("[ ERROR ] Something went wrong, unable to create new note as type:")
			return fmt.Errorf("%v %s", err, noteType)
		}
	}

	return nil
}

func (u *UserNotes) SaveEditsToDatabase() {
    panic("Not implmented yet")
}
