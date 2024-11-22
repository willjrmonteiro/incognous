package connection

import (
	"sync"
)

type Manager struct {
	connections map[string]*Connection
	mutex       sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		connections: make(map[string]*Connection),
	}
}

func (m *Manager) AddConnection(conn *Connection) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.connections[conn.ConnectionID] = conn
}

func (m *Manager) RemoveConnection(connectionID string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.connections, connectionID)
}

func (m *Manager) GetConnectionsByChatID(chatID string) []*Connection {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	connections := make([]*Connection, 0)
	for _, conn := range m.connections {
		if conn.ChatID == chatID {
			connections = append(connections, conn)
		}
	}
	return connections
}

func (m *Manager) BroadcastMessage(chatID string, message []byte) error {
	connections := m.GetConnectionsByChatID(chatID)
	for _, conn := range connections {
		conn.SendMessage(message)
	}
	return nil
}
