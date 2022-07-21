package storage

import (
  "strconv"
  "gitlab.ozon.dev/deadboysdontcry/cool-notes-bot/internal/notes"
  "gitlab.ozon.dev/deadboysdontcry/cool-notes-bot/internal/user"
)

type Storage struct {
  
  idUser map[int]*user.User
  data map[int]*notes.Notes
}

var storage Storage

func init() {
  storage.idUser = make(map[int]*user.User)
  storage.data = make(map[int]*notes.Notes)
}

//creates new note, construct new user, if this is his first request 
func AddNote(userId int, title, body string) {
  if !isExists(userId) {
    addUser(userId)
  }
  storage.data[userId].AddNote(title, body)
}


//updates existing note, construct new user, if this is his first request 
//returns error if noteId is wrong
func UpdateNote(userId, noteId int,body string) error {
  if !isExists(userId) {
    addUser(userId)
  }
  err :=  storage.data[userId].UpdateBody(noteId, body)
  if err != nil {
    return err
  }
  return nil
}

//deletes existing note, construct new user, if this is his first request 
//returns error if noteId is wrong
func DeleteNote(userId, noteId int) error {
  if !isExists(userId) {
    addUser(userId)
  }
  err :=  storage.data[userId].DeleteNote(noteId)
  if err != nil {
    return err
  }
  return nil
}

//returns note time, construct new user, if this is his first request 
//returns error if noteId is wrong
func GetTime(userId int, noteId int) (string, error) {
   if !isExists(userId) {
    addUser(userId)
  }
  userNotes, _  := storage.data[userId] 
  noteTime, err := userNotes.GetTime(noteId)
  if err != nil {
    return "", err
  }
  str := strconv.Itoa(noteTime.Day()) + " " + noteTime.Month().String() + strconv.Itoa(noteTime.Year())
  return str, nil
}

//returns note title, construct new user, if this is his first request 
//returns error if noteId is wrong
func GetTitle(userId int, noteId int) (string, error) {
  if !isExists(userId) {
    addUser(userId)
  }
  userNotes, _  := storage.data[userId] 
  title, err := userNotes.GetTitle(noteId)
  if err != nil {
    return "", err
  }
  return title, nil
}

//returns note body, construct new user, if this is his first request 
//returns error if noteId is wrong
func GetBody(userId int, noteId int) (string, error) {
  if !isExists(userId) {
    addUser(userId)
  }
  userNotes, _  := storage.data[userId] 
  body, err := userNotes.GetBody(noteId)
  if err != nil {
    return "", err
  }
  return body, nil
}

//returns string of all ids and titles, which belong to user with that userId
//construct new user, if this is his first request 
func List(userId int) string {
  if !isExists(userId) {
    addUser(userId)
  }
  userNotes, _  := storage.data[userId] 
  ids, titles := userNotes.GetAllIdsAndTitles()
  var str string
  for indx := range ids {
    str += strconv.Itoa(ids[indx])
    str += ": "
    str += titles[indx];
    str += "\n"
  }
  return str; 
}

func isExists(userId int) bool {
  _, ok := storage.idUser[userId]
  return ok
}

func addUser(userId int) {
  new_user := user.CreareUser(userId)
  storage.idUser[userId] = &new_user
  storage.data[userId] = &notes.Notes{}
}
