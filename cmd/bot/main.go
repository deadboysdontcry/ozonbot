package main

import (
  "log"
  "os"

  "gitlab.ozon.dev/deadboysdontcry/cool-notes-bot/internal/handlers"
  "gitlab.ozon.dev/deadboysdontcry/cool-notes-bot/internal/commander"
)

func main() {
  cmd, err := commander.Init(os.Getenv("ApiKey")) 
  if err != nil {
    log.Panic(err)
  }
  handlers.AddHandlers(cmd) 
  if  err = cmd.Run(); err != nil {
    log.Panic(err)
  }
}
