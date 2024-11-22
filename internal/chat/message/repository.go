package message

import (
	"context"

	"incognous/internal/platform/dynamodb"
	"incognous/message"
)

type Repository interface {
	SaveMessage(ctx context.Context, message *message.Message) error
	GetMessagesByChatID(ctx context.Context, chatID string) ([]*message.Message, error)
}

type DynamoDBRepository struct {
	db *dynamodb.Client
}

func NewDynamoDBRepository(db *dynamodb.Client) *DynamoDBRepository {
	return &DynamoDBRepository{db: db}
}

func (r *DynamoDBRepository) SaveMessage(ctx context.Context, msg *message.Message) error {
	return r.db.SaveMessage(ctx, msg)
}

func (r *DynamoDBRepository) GetMessagesByChatID(ctx context.Context, chatID string) ([]*message.Message, error) {
	return r.db.GetMessagesByChatID(ctx, chatID)
}
