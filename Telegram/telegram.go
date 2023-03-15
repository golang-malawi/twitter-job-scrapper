package telegram

import (
	"fmt"
	"os"
	"github.com/preciousnyasulu/twitter-job-scrapper/twitter"
	"github.com/enescakir/emoji"
	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"strings"
)

// var searchButton = telegrambot.NewInlineKeyboardMarkup(
// 	telegrambot.NewInlineKeyboardRow(
// 		telegrambot.NewInlineKeyboardButtonData("Search " + emoji.Parse(":mag:"), "Search"),
// 	),
// )

func GetUpdate() {
	godotenv.Load(".env")
	bot, err := telegrambot.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := telegrambot.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		
		if update.Message == nil {
			continue
		}
		msg := telegrambot.NewMessage(update.Message.Chat.ID, "Rep: "+update.Message.Text)
		switch update.Message.Command() {
		case "start":
			// txt := "Welcome " + update.Message.From.FirstName + ", lets try to find you a job " + emoji.Parse(":smiley:") + ".n\  Type some phrases that can help us find a job for you."
			txt := fmt.Sprintf("welcome %s, lets try ti find you a job %s. \n\n%sType some phrases that can help us find a job for you e.g Developer, remote, malawi",update.Message.From.FirstName,emoji.Parse(":smiley:"),emoji.Parse(":writing_hand:"))
			
			msg = telegrambot.NewMessage(update.Message.Chat.ID, txt)
			msg.ReplyMarkup = searchButton

		default:
			usertext := update.Message.Text
			//validate text from user
			if len(usertext) < 5 {
				msg = telegrambot.NewMessage(update.Message.Chat.ID, "That was not a valid command, try again.")
			}

			//send each found tweet to the user
			tweets := twitter.GetTweets(usertext)
			for _,value := range tweets{
				result := strings.Trim(fmt.Sprintf("%s",value) , "{}")
				msg = telegrambot.NewMessage(update.Message.Chat.ID,result)
				msg.ReplyToMessageID = update.Message.MessageID
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}
		}

		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

}
