package interactions

type StartupMenu struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

type NotesMenu struct {
    choices []string
    cursor int
    selected map[int]struct{}
}
