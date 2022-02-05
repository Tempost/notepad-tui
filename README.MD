# Little notpad TUI making use of database to save notes

## Keep track of project notes, class notes, or things that need to get done!
    - GORM -- we can store our .md files in the database! in a BLOB column
        -- files <100kib are much faster to store this way than relying on IO to do it
        -- https://sqlite.org/appfileformat.html
    - bubbletea
    - go!
    - markdown
    - open notes in $EDITOR env
    - when .md files update push them to database, first just manually select save notes

## Notes
    - Consider using Glamour to allow viewing of notes in the terminal vs view AND edit in editor
        -- https://github.com/charmbracelet/glamour
