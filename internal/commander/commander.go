package commander

import (
  "github.com/pkg/errors"
  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
  startingOffset = 1
  timeout = 60
)

//template of the response function
type CmdHandler func(int, string) string

var UnknownCommand = errors.New("unknown command")

// commander - structure that handles user messages
//route - for easy handling of responses
type Commander struct {
  bot *tgbotapi.BotAPI
  route map[string]CmdHandler
}

// Init creates a new Commander instance.
func Init(apikey string) (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(apikey)
	if err != nil {
    return nil, err
	}
	bot.Debug = true

  return &Commander{
    bot: bot,
    route: make(map[string]CmdHandler),
  }, nil
}

// Run processes user mesages
func (c* Commander)Run() error{
  u := tgbotapi.NewUpdate(startingOffset)
	u.Timeout = timeout

	updates := c.bot.GetUpdatesChan(u)
  for update := range updates {
		if update.Message == nil {
			continue
		}
    
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			if f, ok := c.route[cmd]; ok {
				msg.Text = f(int(update.Message.From.ID), update.Message.CommandArguments())
			} else {
				msg.Text = "unknown command, type /help to see all commands"
			}
		} else {
      msg.Text = "enter some command pls=)"
    }
		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil

}

//RegisterHandler added new funcs for route
func (c *Commander) RegisterHandler(cmd string, f CmdHandler) {
	c.route[cmd] = f
}

