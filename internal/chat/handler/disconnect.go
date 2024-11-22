package handler

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"incognous/internal/chat/connection"
)

type DisconnectHandler struct {
	connectionManager *connection.Manager
}

func NewDisconnectHandler(connectionManager *connection.Manager) *DisconnectHandler {
	return &DisconnectHandler{connectionManager: connectionManager}
}

func (h *DisconnectHandler) Handle(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	connectionID := request.RequestContext.ConnectionID
	h.connectionManager.RemoveConnection(connectionID)

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}
