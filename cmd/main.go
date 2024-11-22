package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"incognous/handler"
)

func Handler(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.RequestContext.RouteKey {
	case "$connect":
		return handler.ConnectHandler(ctx, request)
	case "$disconnect":
		return handler.DisconnectHandler(ctx, request)
	case "message":
		return handler.MessageHandler(ctx, request)
	default:
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
