package bot

import (
	"errors"
	"fmt"
	"gym-tracker/app/calculos"
	"gym-tracker/app/series"
	"gym-tracker/app/set"
	"gym-tracker/app/user"
	config "gym-tracker/infra"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type Telegram struct {
	bot           *tgbotapi.BotAPI
	seriesService series.Service
	setService    set.Service
	userService   user.Service
	updates       tgbotapi.UpdatesChannel
	states        State
}

type Bot interface {
	Chat()
}

func NewTelegram(config config.Telegram, serieService series.Service, service set.Service, userService user.Service) Telegram {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = config.Debug
	log.Printf("Bot autorizado como %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(config.NewUpdated)
	u.Timeout = config.Timeout

	return Telegram{
		bot:           bot,
		updates:       bot.GetUpdatesChan(u),
		seriesService: serieService,
		setService:    service,
		states:        NewState(),
		userService:   userService,
	}
}

func (t Telegram) Chat() {
	for update := range t.updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		text := update.Message.Text

		if _, ok := t.states.state[chatID]; !ok {
			t.states.state[chatID] = StateIdle
		}
		userExtracted := extractUserFromMessage(update.Message)
		if strings.HasPrefix(text, "/") {
			t.handleCommand(chatID, text, userExtracted)
			continue
		}

		t.handleState(chatID, text, userExtracted)
	}
}

func (t Telegram) handleState(chatID int64, text string, userExtracted user.User) {
	switch t.states.state[chatID] {
	case StateWatingSerie:
		serieDTO := series.NewSeriesDTO(text, userExtracted.ID)
		_, err := t.seriesService.CreateSeries(serieDTO.ToEntity())
		if err != nil {
			log.Println(err)
			t.bot.Send(tgbotapi.NewMessage(chatID, "Erro ao salvar a serie!"))
			return
		}
		t.bot.Send(tgbotapi.NewMessage(chatID, "Serie Salva com sucesso!"))
		t.states.NextState(chatID)
		t.bot.Send(tgbotapi.NewMessage(chatID, "Digite: peso tempo (ex: 80 60)"))
	default:
		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			t.bot.Send(tgbotapi.NewMessage(chatID, "Formato inválido. Use: peso tempo e reps"))
			return
		}

		weight, weightErr := strconv.Atoi(parts[0])
		time, timeErr := strconv.Atoi(parts[1])
		reps, repsErr := strconv.Atoi(parts[2])
		if weightErr != nil || timeErr != nil || repsErr != nil {
			t.bot.Send(tgbotapi.NewMessage(chatID, "Valores inválidos"))
			return
		}

		actualSerie, err := t.seriesService.ActualSerie()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			t.bot.Send(tgbotapi.NewMessage(chatID, "Cadastre uma serie primeiro"))
		}
		setDTO := set.NewDTO(actualSerie.ID, weight, time, reps)

		if err := t.setService.AddSet(setDTO); err != nil {
			t.bot.Send(tgbotapi.NewMessage(chatID, "Erro ao salvar set"))
			return
		}

		t.bot.Send(tgbotapi.NewMessage(chatID, "Set salvo com sucesso"))
	}
}

func (t Telegram) handleCommand(chatID int64, text string, userExtracted user.User) {
	initialMessage := fmt.Sprintf("Bem-vindo ao bot\ncommandos: \n%s\n%s", StartSeries, FinishSeries)

	switch text {
	case string(Start):
		if t.userService.IsNew(userExtracted.ChatID) {
			_, err := t.userService.CreateUser(userExtracted)
			if err != nil {
				log.Fatalln("Error to save user")
			}
		}
		t.bot.Send(tgbotapi.NewMessage(chatID, initialMessage))
	case string(StartSeries):
		t.states.NextState(chatID)
		t.bot.Send(tgbotapi.NewMessage(chatID, "Digite o nome da série"))
	case string(FinishSeries):
		t.states.NextState(chatID)
		serieID, _ := t.seriesService.FinalizeSerie()
		t.setService.FinalizeSet(serieID)
		t.bot.Send(tgbotapi.NewMessage(chatID, "Série finalizada com sucesso"))
	case string(Calculate):
		t.states.NextState(chatID)
		var tvl calculos.TrainingVolumeLoad
		allSeries := t.seriesService.GetALlSeriesByChatID(chatID)
		result := tvl.Calculate(allSeries)
		fmt.Println(result)
	}
}

func extractUserFromMessage(message *tgbotapi.Message) user.User {
	var fullName strings.Builder
	if message.From.FirstName != "" {
		fullName.WriteString(message.From.FirstName)
	}
	if message.From.LastName != "" {
		fullName.WriteString(" ")
		fullName.WriteString(message.From.LastName)
	}

	return user.User{
		FullName: fullName.String(),
		ChatID:   message.Chat.ID,
	}
}
