package handlers

import (
  "strings"
  "strconv"
 
  "github.com/pkg/errors"
  "gitlab.ozon.dev/deadboysdontcry/cool-notes-bot/internal/commander"
  "gitlab.ozon.dev/deadboysdontcry/cool-notes-bot/internal/storage"
)

const (
	helpCmd = "help"
	listCmd = "list"
	addCmd  = "add"
  updateCmd = "update"
  deleteCmd = "delete"
  getCmd = "get"
)

var WrongParametrsCount = errors.New("wrong params count")

func addNote(userId int, data string) string {

  params := strings.Split(data, " ")
  if len(params) <  2 {
    return WrongParametrsCount.Error()
  }
  title := params[0]
  body := strings.Join(params[1:], " ")
  storage.AddNote(userId, title, body)  

  return "note added successfully"
}

func updateNote(userId int, data string) string {
  params := strings.Split(data, " ")
  if len(params) < 2 {
    return WrongParametrsCount.Error()
  }
  noteId, err := strconv.Atoi(params[0])
  if err != nil {
    return errors.New("required number").Error()
  }
  body := strings.Join(params[1:], " ")
  err = storage.UpdateNote(userId, noteId, body) 
  if err != nil {
    return err.Error()
  } 
  return "note updated successfully"
}

func deleteNote(userId int, data string) string {
  params := strings.Split(data, " ")
  if len(params) != 1 {
    return errors.Wrap(WrongParametrsCount, "required 1 params").Error()
  }
  noteId, err := strconv.Atoi(params[0])
  if err != nil {
    return errors.New("required number").Error()
  }
  err = storage.DeleteNote(userId, noteId)
  if err != nil {
    return err.Error()
  }
  return "note deleted successfully"
}

func getTitleAndBody(userId int, data string) string {
  params := strings.Split(data, " ")
  if len(params) != 1 {
    return errors.Wrap(WrongParametrsCount, "required 1 params").Error()
  }
  noteId, err := strconv.Atoi(params[0])
  if err != nil {
    return errors.New("required number").Error()
  }
  var result, title, body string
  title, err = storage.GetTitle(userId, noteId)
  if err != nil {
    return err.Error()
  }
  body, err = storage.GetBody(userId, noteId)
  if err != nil {
    return err.Error()
  }
 
  result = title + ":\n" + body
  return result
}

func getAllNotes(userId int, data string) string {
  return "id:title :\n" + storage.List(userId)
}

func help(userId int, data string) string {
  return "/help - list commands\n" +
		"/list - list of all notes\n" +
		"/add <title> <body> - add new note\n" +
    "/update <id> <new body> - update note body with new body\n" + 
    "/delete <id> - delete note\n" +
    "/get <id> - title and body\n" 

}


//AddHandlers registers handlers for Commander.route
func AddHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, help)
	c.RegisterHandler(listCmd, getAllNotes)
	c.RegisterHandler(updateCmd, updateNote) 
	c.RegisterHandler(deleteCmd, deleteNote) 
	c.RegisterHandler(getCmd, getTitleAndBody) 
	c.RegisterHandler(addCmd, addNote) 
}

