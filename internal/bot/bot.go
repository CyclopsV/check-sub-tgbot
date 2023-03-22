package bot

import (
	"fmt"
	"log"

	"github.com/CyclopsV/check-sub-tgbot/configs"
	"github.com/NicoNex/echotron/v3"
)

type Bot struct {
	configs configs.Config
	api     echotron.API
}

func New(config configs.Config) *Bot {
	api := echotron.NewAPI(config.TgBotToken)
	botInfo, err := api.GetMe()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", botInfo.Result)

	return &Bot{
		configs: config,
		api:     api,
	}

}

func (b *Bot) Run() {
	for u := range echotron.PollingUpdates(b.configs.TgBotToken) {
		if u.Message == nil {
			continue
		}
		if b.isCommand(u.Message.Entities) {
			answer := fmt.Sprintln("Призвание данного бота заключается лишь в одном: защита данного чата от нападков других ботов, путем требования подписки, для общения в данном чате.\n\nВсе вопросы касательно бота направлять <a href=\"https://t.me/CyclopsV\">@CyclopsV</a>")
			b.api.SendMessage(answer, u.ChatID(), &echotron.MessageOptions{ParseMode: "html"})
		}

		if b.checkMember(u) {
			b.sendWarn(u)
		}
	}
}

func (b *Bot) isCommand(ent []*echotron.MessageEntity) bool {
	for _, el := range ent {
		if el.Type == "bot_command" {
			return true
		}
	}
	return false
}

func (b *Bot) checkMember(u *echotron.Update) bool {
	checkChat, _ := b.api.GetChatMember(b.configs.ChatID, u.Message.From.ID)
	checkChannel, _ := b.api.GetChatMember(b.configs.ChannelID, u.Message.From.ID)
	return checkChannel.Result.Status == "left" || checkChat.Result.Status == "left"
}

func (b *Bot) sendWarn(u *echotron.Update) {
	b.api.DeleteMessage(b.configs.ChatID, u.Message.ID)
	answer := fmt.Sprintf("Ишь какой неприличный, <a href=\"tg://user?id=%v\">@%v</a>, иди подпишись на <a href=\"https://t.me/RolePlay_Games\">канал</a> чата, а потом возвращайся и общайся", u.Message.From.ID, u.Message.From.Username)
	b.api.SendMessage(answer, u.ChatID(), &echotron.MessageOptions{ParseMode: "html"})
}
