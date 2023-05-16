package echo

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	ydbAccess "github.com/makropus/telegram-echo-bot/internal/yandex-cloud"
)

var (
	db *sql.DB

	bot  *tgbotapi.BotAPI
	user *tgbotapi.User

	scream bool
)

func HandleWebHookRequest(rw http.ResponseWriter, req *http.Request) {
	db = ydbAccess.ConnectToYDB()

	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("ERROR FATAL AUTH: %s", err)
	}
	updateCh := bot.ListenForWebhookRespReqFormat(rw, req)
	update := <-updateCh
	handleUpdate(&update)
}

func handleUpdate(update *tgbotapi.Update) {
	if update != nil && update.Message != nil {
		handleMessage(update.Message)
	}
}

func handleMessage(message *tgbotapi.Message) {
	user = message.From
	text := message.Text

	scream = getScreamingStatus(int(user.ID))

	if user == nil {
		return
	}

	log.Printf("INFO: [%d:%s] wrote %s\n", user.ID, user.FirstName, text)

	var err error
	if strings.HasPrefix(text, "/") {
		handleCommand(text)
	} else if len(text) > 0 {
		if scream {
			text = strings.ToUpper(text)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		//msg.Entities = message.Entities
		_, err = bot.Send(msg)
	}

	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}
	saveScreamingStatus(user.ID, user.UserName, scream)
}

func handleCommand(command string) {
	switch command {
	case "/scream":
		scream = true
	case "/whisper":
		scream = false
	}
}

func getScreamingStatus(userID int) bool {
	var status bool
	query := fmt.Sprintf(
		"SELECT scream FROM `telegramEchoBot/screamStatus` WHERE id=%d",
		userID)
	// for some reason db.QueryRow("some query ?", arg) doesnt insert arg in place of '?'
	err := db.QueryRow(query).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Printf("ERROR READ DB: %s\n", err)
	}
	return status
}

func saveScreamingStatus(userID int64, username string, status bool) {
	query := fmt.Sprintf(
		"REPLACE INTO `telegramEchoBot/screamStatus` (id, username, scream) VALUES (%d, \"%s\", %t)",
		userID, username, status)
	log.Println(`SAVE QUERY: ` + query)
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("ERROR WRITE DB:: %s\n", err)
	}
}
