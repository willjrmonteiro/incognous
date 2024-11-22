package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"incognous/internal/chat/connection"
	"incognous/internal/chat/handler"
	"incognous/internal/chat/message"
	"incognous/internal/platform/dynamodb"
)

func main() {
	connectionManager := connection.NewManager()
	dbClient := dynamodb.NewClient()
	messageRepository := message.NewDynamoDBRepository(dbClient)

	connectHandler := handler.NewConnectHandler(connectionManager)
	disconnectHandler := handler.NewDisconnectHandler(connectionManager)
	messageHandler := handler.NewMessageHandler(connectionManager, messageRepository)

	lambda.Start(func(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.RequestContext.RouteKey {
		case "$connect":
			return connectHandler.Handle(ctx, request)
		case "$disconnect":
			return disconnectHandler.Handle(ctx, request)
		case "message":
			return messageHandler.Handle(ctx, request)
		default:
			return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound}, nil
		}
	})
}
