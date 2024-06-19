package actions

import (
	"athena/server/clients"
	"encoding/json"
)

type Action func(message *ClientMessage) *ServerMessage

type ClientMessage struct {
	Id     int             `json:"id"`
	Type   int             `json:"type"`
	Data   json.RawMessage `json:"data"`
	Token  string          `json:"token"`
	Client *clients.Connection
}

type ServerMessage struct {
	Id   int             `json:"id"`
	Type int             `json:"type"`
	Data json.RawMessage `json:"data"`
}

const (
	RoleValue_Observer      = 1
	RoleValue_Producer      = 2
	RoleValue_Manager       = 3
	RoleValue_Owner         = 4
	RoleValue_Administrator = 5
)

const (
	MessageType_Ping = 1
	MessageType_Pong = 2

	MessageType_Login       = 10
	MessageType_Logout      = 11
	MessageType_ResumeLogin = 12

	MessageType_Success          = 200
	MessageType_NotAcceptable    = 210
	MessageType_NotAuthenticated = 211
	MessageType_NotAuthorized    = 212
	MessageType_NotAvailable     = 213
)

var clientActions = map[int]Action{
	MessageType_Login:       Login,
	MessageType_Logout:      Logout,
	MessageType_ResumeLogin: ResumeLogin,
}

func CreateClientMessage(client *clients.Connection, bytes []byte) (message *ClientMessage, err error) {
	message = &ClientMessage{}
	err = json.Unmarshal(bytes, message)

	if err == nil {
		message.Client = client
	}

	return
}

func CreateServerMessage(clientMessage *ClientMessage, messageType int, messageData any) *ServerMessage {
	message := &ServerMessage{
		Id:   clientMessage.Id,
		Type: messageType,
	}

	if messageData != nil {
		bytes, _ := json.Marshal(messageData)
		message.Data = bytes
	}

	return message
}

func CreateServerMessageForAccount(clientMessage *ClientMessage, messageType int, messageData any) *ServerMessage {
	response := CreateServerMessage(clientMessage, messageType, messageData)
	response.Id = 0
	return response
}

func Execute(message *ClientMessage) *ServerMessage {
	if message.Type == MessageType_Ping {
		return CreateServerMessage(message, MessageType_Pong, nil)
	}

	action, found := clientActions[message.Type]

	if !found {
		return notAcceptable(message)
	}

	return action(message)
}

func decodeMessageData(message *ClientMessage, destination any) bool {
	if message.Data != nil {
		return json.Unmarshal(message.Data, destination) == nil
	}

	return false
}

func isAuthorized(message *ClientMessage, roleValue int64) bool {
	return message.Client.UserRole >= roleValue
}

func success(message *ClientMessage, data any) *ServerMessage {
	return CreateServerMessage(message, MessageType_Success, data)
}

func notAcceptable(message *ClientMessage) *ServerMessage {
	return CreateServerMessage(message, MessageType_NotAcceptable, nil)
}

func notAuthenticated(message *ClientMessage) *ServerMessage {
	return CreateServerMessage(message, MessageType_NotAuthenticated, nil)
}

func notAuthorized(message *ClientMessage) *ServerMessage {
	return CreateServerMessage(message, MessageType_NotAuthorized, nil)
}

func notAvailable(message *ClientMessage) *ServerMessage {
	return CreateServerMessage(message, MessageType_NotAvailable, nil)
}
