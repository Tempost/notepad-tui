package database

// this package will cover all/majority of the database functions
// consider storting types/interfaces for the database in a seperate file

// when opening different note types load them into memory
// write any changed to the DB upon leaving the note types menu

// when editing a note
//  -- Pull markdown file from database
//  -- Create a temp file in edits directory
//  -- Push all files in edits directory to database when use saves notes from TUI

import (
	"fmt"
    "errors"
	"os"
    "strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserNotes struct {
	gorm.Model
	UserName string
	// consider converting slice of notes into maps
	ProjectNotes  []ProjectNote
	PersonalNotes []PersonalNote
	ClassNotes    []ClassNote
	Scratchpads   []Scratchpad
}

type ProjectNote struct {
    gorm.Model
	Name string
    // do not store this as a pointer in the DB
    notes *os.File // check if file is actually a markdown file 
}

type PersonalNote struct {
    gorm.Model
	Name string
    notes *os.File
}

type ClassNote struct {
    gorm.Model
	Name string
    notes *os.File
}

type Scratchpad struct {
    gorm.Model
    notes *os.File
}

var bd *gorm.DB

// On runtime of the application check for a database and it will open
// if no database is found this will create a new database, populate with defaults and open
// The defaults are just one blank note for each note type
func CheckForDatabase() (*gorm.DB, error) {
    filename := os.Getenv("USER") + ".db"
    filepath := "../../data/" + filename

    // check if our file exists, otherwise skip code inside if scope and open the database
    if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
        os.Create(filepath) 
        return OpenDatabase(filepath)
    }
    
    return OpenDatabase(filepath)
}

func OpenDatabase(filepath string) (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(filepath))
    if err != nil {
        return nil, fmt.Errorf("[ ERROR ]: %v", err)
    }
    return db, nil
}

func CreateNewProjectNote() {
	panic("[ INFO ] Not implmented yet. Consider working on this")
}

func CreateNewPersonalNote() {
	panic("[ INFO ] Not implmented yet. Consider working on this")
}

func CreateNewClassNote() {
	panic("[ INFO ] Not implmented yet. Consider working on this")
}

// when go generics is released in stable consider changing this up
func (u *UserNotes) CreateNewNote(noteType string) error {
    noteType = strings.ToLower(noteType)

    switch noteType {
        case "projectnote": {
            CreateNewProjectNote()
        }
        case "personal": {
            CreateNewPersonalNote()
        }
        case "class": {
            CreateNewClassNote()
        }
        default: {
            err := errors.New("[ ERROR ] Something went wrong, unable to create new note as type:") 
            return fmt.Errorf("%v %s", err, noteType)
        }
    } 

    return errors.New("[ ERROR ] Something went really wrong, broke out of switch statement. This should never happen.") 
}
