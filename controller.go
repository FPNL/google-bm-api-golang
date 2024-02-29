package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	log.Println("body is ", string(body))
	log.Println("header is ", r.Header)

	var msg UserMessage
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Println("error is ", err)
		return
	}

	if msg.Secret != "" && msg.ClientToken != "" {
		w.WriteHeader(200)
		_, err = w.Write([]byte(msg.Secret))
		if err != nil {
			log.Println("error is ", err)
			return
		}
		return
	}

	conversationId := msg.ConversationId

	if msg.Message.Text != "" {
		message := msg.Message.Text

		// Log message parameters
		log.Println("conversationId: ", conversationId)
		log.Println("message: ", message)

		echoMessage(message, conversationId)
	} else if msg.SuggestionResponse.PostbackData != "" {
		message := msg.SuggestionResponse.PostbackData

		// Log message parameters
		log.Println("conversationId: ", conversationId)
		log.Println("message: ", message)

		echoMessage(message, conversationId)
	} else if msg.UserStatus.IsTyping {
		log.Println("User is typing")
	} else if msg.UserStatus.RequestedLiveAgent {
		log.Println("User requested transfer to live agent")
	}

	w.WriteHeader(200)
}
