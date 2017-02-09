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

	//�����T��  case *linebot.TextMessage:

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {

			   case "�A�n":
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+"�i�H��ڲ�Ѯ@!")).Do();
				 case "�ߦw":
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("�ߦw�A���A���n��!")).Do();
				 case "�A�O��":
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("�A���k�B��")).Do();
				 case "���ѦY�ƻ�":
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("�Z�ҥ��\�ܾA�X�A")).Do();
				 case "���^��":
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("�L�O�ڤڴ˶�?")).Do();
				 case "���ڿ�":
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("�ڷ|�N���A")).Do();
				 case "�ڥi�H������?":
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("�A�i�H�� �A�O�� ���ѦY�ƻ� ���ڿ� �A�n")).Do();
		
				
				default:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("�A��"+message.Text+"�ڤ�����")).Do();
				
				 err != nil {
					log.Print(err)
				}
			}
		}
	}
}
