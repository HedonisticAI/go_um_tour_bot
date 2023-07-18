package main

import (
	"log"
	"strings"

	get_pair "github.com/HedonisticAI/go_um_tour_bot/http"
	"github.com/HedonisticAI/go_um_tour_bot/map_char"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	Char := map_char.Fill("unmatched_characters.json")
	Map := map_char.Fill("unmatched_maps.json")
	Auth := make(map[int64]get_pair.User)
	bot, err := tgbotapi.NewBotAPI("6383506273:AAEiwcrIqt6A272co6C24pSNZISqrUpf6UI")
	if err != nil {
		panic(err)
	}
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {

		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msg.ReplyToMessageID = update.Message.MessageID
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /auth and /draft.\n"
			msg.Text += "usage of auth - '/auth <challonge login> <challonge password> <tournament>'\n"
			msg.Text += "usage of draft - after you used /auth you can just type '/draft'"
		case "auth":
			params := strings.Split(update.Message.Text, " ")
			if len(params) != 4 {
				msg.Text = "Please enter correct params"
				break
			}
			Auth[update.Message.From.ID] = get_pair.User{params[1], params[2], params[3]}
			log.Printf("User with %d id %s Username %s pwd %s Tourmanent id found\n", update.Message.From.ID, Auth[update.Message.From.ID].Username, Auth[update.Message.From.ID].Pwd, Auth[update.Message.From.ID].TournamentID)
			msg.Text = "User with this params authorized:\n" + Auth[update.Message.From.ID].Username + "\n" + Auth[update.Message.From.ID].Pwd + "\n" + Auth[update.Message.From.ID].TournamentID

		case "draft":
			if val, ok := Auth[update.Message.From.ID]; ok {
				log.Printf("User with %s name requested draft\n", val.Username)
				msg.Text = val.Draft(Char, Map)
			} else {
				msg.Text = "Please use auth command first"
			}

		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
	bot.Debug = true
}
