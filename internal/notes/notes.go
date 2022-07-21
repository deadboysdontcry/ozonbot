package notes

import (
  "time"
  "github.com/pkg/errors"
)

//notes - slice of Notes, it is sorted by note.time
//each index of the Note is it's id
type Notes struct {
  notes []Note // sorted by time
}

//returns note time by id, if id is wrong returns error
func (n *Notes) GetTime(id int) (time.Time, error) {
  node, err := n.getNote(id) 
  if err != nil {
    return time.Time{}, err
  }
  return node.GetTime(), nil
}

//returns note body by id, if id is wrong returns error
func (n *Notes) GetBody(id int) (string, error) {
  node, err := n.getNote(id) 
  if err != nil {
    return "", err
  }
  return node.GetBody(), nil
}

//returns note title by id, if id is wrong returns error
func (n *Notes) GetTitle(id int) (string, error) {
  node, err := n.getNote(id) 
  if err != nil {
    return "", err
  }
  return node.GetTitle(), nil
}

//deletes note form the slice by id, if id is wrong returns error
func (n *Notes) DeleteNote(id int) error {
  if len(n.notes) <= id {
    return errors.New("bad id")
  }
  n.notes = append(n.notes[:id], n.notes[id + 1:]...)
  return nil
}

//adds new node 
func (n *Notes) AddNote(title, body string) {
  n.addNote(CreateNote(title, body))
}

//updates body of existing note by id, return error if id is wrong
func (n *Notes) UpdateBody(id int, newBody string) error {
  if len(n.notes) <= id {
    return errors.New("bad id")
  }
  updatedNote := n.notes[id]
  updatedNote.SetNoteBody(newBody)
  updatedNote.SetCurrentTime()
  _ = n.DeleteNote(id)
  n.addNote(updatedNote)
  return nil
}

//returns 2 slices of ids and titles, each id is id of the string with same index
func (n *Notes) GetAllIdsAndTitles() ([]int, []string) {
  var ids []int
  var titles []string
  for id, note := range n.notes {
    ids = append(ids, id)
    titles = append(titles, note.GetTitle())
  }
  return ids, titles
}

func (n* Notes) getNote(id int) (*Note, error) {
   if len(n.notes) <= id {
    return nil, errors.New("bad id")
  }
  return &n.notes[id], nil;
}

func (n *Notes) addNote(note Note) {
  n.notes = append(n.notes, note)
}


