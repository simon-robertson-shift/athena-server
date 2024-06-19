package clients

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	AccountId int64
	UserId    int64
	UserRole  int64
	Socket    *websocket.Conn
}

var connectionList = make([]*Connection, 0, 1000)
var connectionListMutex sync.Mutex
var disconnectionCount = 0
var disconnectionThreshold = 100

func Connect(socket *websocket.Conn) *Connection {
	connection := &Connection{
		Socket: socket,
	}

	connectionListMutex.Lock()

	added := false

	for i := 0; i < len(connectionList); i++ {
		if connectionList[i] != nil {
			continue
		}

		connectionList[i] = connection
		added = true
		break
	}

	if !added {
		connectionList = append(connectionList, connection)
	}

	connectionListMutex.Unlock()

	return connection
}

func Disconnect(connection *Connection) {
	connectionListMutex.Lock()

	for i := 0; i < len(connectionList); i++ {
		if connectionList[i] != connection {
			continue
		}

		connection.Socket.Close()

		connectionList[i] = nil
		disconnectionCount++
		break
	}

	if disconnectionCount >= disconnectionThreshold {
		nextConnectionList := make([]*Connection, 0, cap(connectionList))

		for i := 0; i < len(connectionList); i++ {
			value := connectionList[i]

			if value == nil {
				continue
			}

			nextConnectionList = append(nextConnectionList, value)
		}

		connectionList = nextConnectionList
		disconnectionCount = 0
	}

	connectionListMutex.Unlock()
}

// Sends a message to a specific client.
func SendMessage(message interface{}, client *Connection) {
	bytes, _ := json.Marshal(message)

	if err := client.Socket.WriteMessage(websocket.TextMessage, bytes); err != nil {
		Disconnect(client)
	}
}

// Sends a message to all clients within a specific account.
func SendMessageToAccountClients(message interface{}, client *Connection) {
	bytes, _ := json.Marshal(message)

	connectionListMutex.Lock()

	for i := 0; i < len(connectionList); i++ {
		connection := connectionList[i]

		if connection == nil || connection.AccountId != client.AccountId {
			continue
		}

		if err := connection.Socket.WriteMessage(websocket.TextMessage, bytes); err != nil {
			connectionList[i] = nil
			disconnectionCount++
		}
	}

	connectionListMutex.Unlock()
}

// Sends a message to all clients connected to the server.
func SendMessageToClients(message interface{}) {
	bytes, _ := json.Marshal(message)

	connectionListMutex.Lock()

	for i := 0; i < len(connectionList); i++ {
		connection := connectionList[i]

		if connection == nil {
			continue
		}

		if err := connection.Socket.WriteMessage(websocket.TextMessage, bytes); err != nil {
			connectionList[i] = nil
			disconnectionCount++
		}
	}

	connectionListMutex.Unlock()
}
