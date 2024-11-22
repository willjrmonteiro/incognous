package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"incognous/message"
)

type Client struct {
	db *dynamodb.DynamoDB
}

func NewClient() *Client {
	sess, _ := session.NewSession()
	return &Client{
		db: dynamodb.New(sess),
	}
}

func (c *Client) SaveMessage(ctx context.Context, msg *message.Message) error {
	av, err := dynamodbattribute.MarshalMap(msg)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("chat-table"),
	}

	_, err = c.db.PutItemWithContext(ctx, input)
	return err
}

func (c *Client) GetMessagesByChatID(ctx context.Context, chatID string) ([]*message.Message, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("chat-table"),
		KeyConditionExpression: aws.String("ChatID = :chatID"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":chatID": {S: aws.String(chatID)},
		},
	}

	result, err := c.db.QueryWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	var messages []*message.Message
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
