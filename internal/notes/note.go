package notes

import (
  "time"
)

//time - latest node update time
//title - note title
//body -- note body
type Note struct {
   time time.Time
   title string
   body string
}

//constructs new note
func CreateNote(title, body string) Note {
  note := Note{}
  note.SetCurrentTime()
  note.SetNoteTitle(title)
  note.SetNoteBody(body)
  return note
}

//title setter
func (n *Note) SetNoteTitle(title string) {
  if len(title) == 0 {
    n.title = "no title"
  } else {
    n.title = title
  }
}

//body setter
func (n *Note) SetNoteBody(body string) {
  if len(body) == 0 {
    n.title = "no body"
  } else {
    n.body = body
  }
}

//sets Note.time
func (n *Note) SetCurrentTime() {
  n.time = time.Now()
}

//time getter
func (n *Note) GetTime() time.Time {
  return n.time
}

//title getter
func (n *Note) GetTitle() string {
  return n.title
}

//body getter
func (n *Note) GetBody() string {
  return n.body
}
