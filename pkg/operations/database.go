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
	// Generating a filename for the database based on the current users name
	filename := os.Getenv("USER") + ".db"

	// TODO: Once we get to were we can "ship" a binary determine where the data should go
	filepath := "./data/" + filename

	// check if our database exists, otherwise skip code inside if scope and open the database
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		os.Create(filepath)

		if err := OpenDatabase(filepath); err != nil {
			return err
		}

		user := UserNotes{
			UserName: os.Getenv("USER"),
		}

		err := db.AutoMigrate(
			&UserNotes{},
			&ProjectNote{},
			&PersonalNote{},
			&ClassNote{},
			&Scratchpad{},
		)

		if err != nil {
			return err
		}

		// Save our newly created user to the database
		db.Save(&user)
	}

	return OpenDatabase(filepath)
}

// Pretty straight forward, we are just opening up our database and storing the handle in the db global variable
// Requires a variable to store the error
func OpenDatabase(filepath string) error {
	var err error
	db, err = gorm.Open(sqlite.Open(filepath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

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
	note, err := os.Create("./edits/" + name + ".md")
	if err != nil {
		log.Fatalln("[ ERROR ] Could not create file: ", err)
	}
	p.Name = name
	p.note = note
}

// Create new markdownfile -> projectname_personal.md
func (p *PersonalNote) CreateNewPersonalNote(name string) {
	note, err := os.Create("./edits/" + name + ".md")
	if err != nil {
		log.Fatalln("[ ERROR ] Could not create file: ", err)
	}
	p.Name = name
	p.note = note
}

// Create new markdownfile -> projectname_class.md
func (c *ClassNote) CreateNewClassNote(name string) {
	note, err := os.Create("./edits/" + name + ".md")
	if err != nil {
		log.Fatalln("[ ERROR ] Could not create file: ", err)
	}
	c.Name = name
	c.note = note
}

// TODO(Cody): when go generics is released in stable consider changing this up
// Creating a new note that the user wishes to edit
// Requires a variable to handle errors
func (u *UserNotes) CreateNewNote(noteType string, name string) (error, *os.File) {
	noteType = strings.ToLower(noteType)
	var note *os.File
	switch noteType {
	case "project":
		{
			var p ProjectNote
			p.CreateNewProjectNote(name)
			u.ProjectNotes = append(u.ProjectNotes, p)
		}
	case "personal":
		{
			var p PersonalNote
			p.CreateNewPersonalNote(name)
			u.PersonalNotes = append(u.PersonalNotes, p)
		}
	case "class":
		{
			var c ClassNote
			c.CreateNewClassNote(name)
			u.ClassNotes = append(u.ClassNotes, c)
		}
	default:
		{
			err := errors.New("[ ERROR ] Something went wrong, unable to create new note as type:")
			return fmt.Errorf("%v %s", err, noteType), nil
		}
	}

	return nil, note
}

// NOTE(Cody): It can't be this simple can it?
func (u *UserNotes) SaveEditsToDatabase() {
	db.Save(&u)
}
