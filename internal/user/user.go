package user

import (
  "github.com/pkg/errors"
)

  
var BadName = errors.New("bad name")

type User struct {
  id int
}

//constructs new user 
func CreareUser(Id int) User {
  newUser:= User{}
  newUser.SetId(Id)
  return newUser
}

// id setter
func (u *User) SetId(id int){
  u.id = id 
}

//id getter
func (u *User) GetUserId() int {
  return u.id
}
