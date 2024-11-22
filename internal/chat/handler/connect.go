package handler

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"incognous/internal/chat/connection"
)

type ConnectHandler struct {
	connectionManager *connection.Manager
}

func NewConnectHandler(connectionManager *connection.Manager) *ConnectHandler {
	return &ConnectHandler{connectionManager: connectionManager}
}

func (h *ConnectHandler) Handle(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	connectionID := request.RequestContext.ConnectionID
	endpoint := request.RequestContext.DomainName + "/" + request.RequestContext.Stage

	conn := connection.NewConnection(connectionID, endpoint)
	h.connectionManager.AddConnection(conn)

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}
