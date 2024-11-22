package repository_test

import (
	"testing"

	"incognous/internal/chat/message"
	"incognous/internal/chat/message/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessageRepository(t *testing.T) {
	repo := repository.NewMessageRepository()

	t.Run("SaveMessage", func(t *testing.T) {
		msg := &message.Message{
			ID:      "1",
			Content: "Hello, World!",
		}
		err := repo.SaveMessage(msg)
		require.NoError(t, err)

		savedMsg, err := repo.GetMessage("1")
		require.NoError(t, err)
		assert.Equal(t, msg, savedMsg)
	})

	t.Run("GetMessage_NotFound", func(t *testing.T) {
		_, err := repo.GetMessage("nonexistent")
		assert.Error(t, err)
	})
}
