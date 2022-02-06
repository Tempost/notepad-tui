package database

import (
    "os"
	"gorm.io/gorm"
)

type UserNotes struct {
	gorm.Model
	UserName string
	// consider converting slice of notes into maps
	ProjectNotes  []ProjectNote
	PersonalNotes []PersonalNote
	ClassNotes    []ClassNote
	Scratchpad
}

type ProjectNote struct {
	gorm.Model
	Name string
	// do not store this as a pointer in the DB
	note *os.File // check if file is actually a markdown file
}

type PersonalNote struct {
	gorm.Model
	Name string
	note *os.File
}

type ClassNote struct {
	gorm.Model
	Name string
	note *os.File
}

type Scratchpad struct {
	gorm.Model
	note *os.File
}
