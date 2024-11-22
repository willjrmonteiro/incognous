package connection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionManager(t *testing.T) {
	t.Run("TestCreateConnection", func(t *testing.T) {
		manager := NewConnectionManager()
		conn := manager.CreateConnection("test-connection")
		assert.NotNil(t, conn)
		assert.Equal(t, "test-connection", conn.ID)
	})

	t.Run("TestGetConnection", func(t *testing.T) {
		manager := NewConnectionManager()
		manager.CreateConnection("test-connection")
		conn := manager.GetConnection("test-connection")
		assert.NotNil(t, conn)
		assert.Equal(t, "test-connection", conn.ID)
	})

	t.Run("TestDeleteConnection", func(t *testing.T) {
		manager := NewConnectionManager()
		manager.CreateConnection("test-connection")
		manager.DeleteConnection("test-connection")
		conn := manager.GetConnection("test-connection")
		assert.Nil(t, conn)
	})
}
