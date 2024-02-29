package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"

	"github.com/FPNL/google-bm-api-golang/businessmessages/v1"
)

var representative = &businessmessages.BusinessMessagesRepresentative{
	AvatarImage:        "https://storage.googleapis.com/sample-avatars-for-bm/bot-avatar.jpg",
	DisplayName:        "Echo Bot",
	RepresentativeType: "BOT",
}

func echoMessage(message string, conversationId string) error {
	s, err := businessmessages.NewService(context.TODO(), opts...)
	if err != nil {
		log.Printf("businessmessages.NewService: %v\n", err)
		return fmt.Errorf("businessmessages.NewService: %v", err)
	}

	eventCall := s.Conversations.Events.Create(
		"conversations/"+conversationId,
		&businessmessages.BusinessMessagesEvent{
			EventType:      "TYPING_STARTED",
			Representative: representative,
		},
	)
	eventCall.EventId(uuid.New().String())
	event, err := eventCall.Do()
	if err != nil {
		log.Printf("s.Conversations.Events.Create: %v\n", err)
		return fmt.Errorf("s.Conversations.Events.Create: %v", err)
	}

	log.Println(event)

	msgCall := s.Conversations.Messages.Create(
		"conversations/"+conversationId,
		&businessmessages.BusinessMessagesMessage{
			ContainsRichText: false,
			Fallback:         "fallback message: hello world",
			Image:            nil,
			MessageId:        uuid.New().String(),
			// Name:             "",
			Representative: representative,
			// RichCard:         nil,
			// Suggestions:      nil,
			Text: message,
		},
	)

	msg, err := msgCall.Do()
	if err != nil {
		log.Printf("s.Conversations.Messages.Create: %v\n", err)
		return fmt.Errorf("s.Conversations.Messages.Create: %v", err)
	}
	log.Println(msg)

	eventCall = s.Conversations.Events.Create(
		"conversations/"+conversationId,
		&businessmessages.BusinessMessagesEvent{
			EventType:      "TYPING_STOPPED",
			Representative: representative,
		},
	)

	eventCall.EventId(uuid.New().String())

	event, err = eventCall.Do()
	if err != nil {
		log.Printf("s.Conversations.Events.Create: %v\n", err)
		return fmt.Errorf("s.Conversations.Events.Create: %v", err)
	}

	log.Println(event)

	return nil
}
