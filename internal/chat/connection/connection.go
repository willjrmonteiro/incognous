package handler

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"incognous/connection"
)

var connectionManager = connection.NewManager()

// ConnectHandler lida com as conexões WebSocket.
func ConnectHandler(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	connectionID := request.RequestContext.ConnectionID
	endpoint := request.RequestContext.DomainName + "/" + request.RequestContext.Stage

	conn := connection.NewConnection(connectionID, endpoint)
	connectionManager.AddConnection(conn)

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}

func main() {
	lambda.Start(ConnectHandler)
}
