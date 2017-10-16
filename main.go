package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "context"
	//"time"
	// "encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	// defer cancel()
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
			if message.Text == "莊于鋅" {
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("低能兒")).Do(); err != nil {
						log.Print(err)
					}
			}
			if message.Text == "抽" {
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("抽三小")).Do(); err != nil {
						log.Print(err)
					}
			}
			if message.Text == "追不到" {
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("YEAH")).Do(); err != nil {
						log.Print(err)
					}
					if message.Text == "台大" {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("游采青")).Do(); err != nil {
								log.Print(err)
							}
			}
		}
	}
}
}
