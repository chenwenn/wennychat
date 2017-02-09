// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
				
				
		case *linebot.TextMessage:
				
				log.Println(message.Text)
				if message.Text=="狗" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("豬隊友就是在說你吧!")).Do();	
				err != nil {
					log.Print(err)
				}
				}
				if message.Text=="晚安" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("晚安!祝你有個好夢~")).Do();	
				err != nil {
					log.Print(err)
				}
				}
				
					if message.Text=="你好" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+"!可以跟我聊天ㄛ!")).Do();	
				err != nil {
					log.Print(err)
				}
				}
					if message.Text=="你是誰" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("FBI偵探"+message.Texta)).Do();	
				err != nil {
					log.Print(err)
				}
				}
						if message.Text=="蔡英文" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("他是歐巴森嗎")).Do();	
				err != nil {
					log.Print(err)
				}
				}
					if message.Text=="今天吃甚麼" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("廁所用餐很適合你")).Do();	
				err != nil {
					log.Print(err)
				}
				}
				if message.Text=="給我錢" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我會燒給你")).Do();	
				err != nil {
					log.Print(err)
				}
				}
				if message.Text=="我可以說什麼?" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你可以說 你是誰 今天吃甚麼 給我錢 你好 等等...")).Do();	
				err != nil {
					log.Print(err)
				}
				}

				
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你說"+message.Text+"我不明白")).Do();	
				err != nil {
					log.Print(err)
				}
				
				
				
				
			}
		}
		
}
}
