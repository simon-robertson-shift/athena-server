package server

import (
	"athena/server/actions"
	"athena/server/clients"
	"athena/server/env"
	"net/http"

	"github.com/gorilla/websocket"
)

var connectionUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(req *http.Request) bool {
		// TODO
		return true
	},
}

func Start() {
	socketServerAddress := env.GetSocketServerAddress()
	socketServerPath := env.GetSocketServerPath()

	http.HandleFunc(socketServerPath, handleRequest)
	http.ListenAndServe(socketServerAddress, nil)
}

func handleRequest(response http.ResponseWriter, request *http.Request) {
	socket, err := connectionUpgrader.Upgrade(response, request, nil)

	if err != nil {
		return
	}

	handleConnectionRequests(clients.Connect(socket))
}

func handleConnectionRequests(connection *clients.Connection) {
	for {
		messageType, bytes, err := connection.Socket.ReadMessage()

		if err != nil {
			break
		}

		if messageType != websocket.TextMessage {
			continue
		}

		request, err := actions.CreateClientMessage(connection, bytes)

		if err != nil {
			continue
		}

		response := actions.Execute(request)

		if response == nil {
			panic("action execution response must not be nil")
		}

		clients.SendMessage(response, connection)
	}

	clients.Disconnect(connection)
}
