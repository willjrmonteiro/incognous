package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"incognous/internal/chat/connection"
	"incognous/internal/chat/message"
	"incognous/message"
)

type MessageHandler struct {
	connectionManager *connection.Manager
	messageRepository message.Repository
}

func NewMessageHandler(
	connectionManager *connection.Manager,
	messageRepository message.Repository,
) *MessageHandler {
	return &MessageHandler{
		connectionManager: connectionManager,
		messageRepository: messageRepository,
	}
}

func (h *MessageHandler) Handle(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody MessageRequestBody
	if err := json.Unmarshal([]byte(request.Body), &reqBody); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	msg := message.NewMessage(reqBody.ChatID, reqBody.SenderID, reqBody.Content)

	if err := h.messageRepository.SaveMessage(ctx, msg); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	messageBytes, _ := json.Marshal(msg)
	h.connectionManager.BroadcastMessage(reqBody.ChatID, messageBytes)

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}
