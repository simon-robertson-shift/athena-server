package actions

import (
	"athena/server/clients"
	"encoding/json"
)

type Action func(request *Request) *Response

type Request struct {
	Id     int             `json:"id"`
	Type   int             `json:"type"`
	Data   json.RawMessage `json:"data"`
	Token  string          `json:"token"`
	Client *clients.Connection
}

type Response struct {
	Id   int             `json:"id"`
	Type int             `json:"type"`
	Data json.RawMessage `json:"data"`
}

const (
	RequestType_Ping = 1

	RequestType_Login  = 10
	RequestType_Logout = 11
)

const (
	ResponseType_AccountNotification = 0

	ResponseType_Pong    = 1
	ResponseType_Success = 2

	ResponseType_NotAcceptable    = 10
	ResponseType_NotAuthenticated = 11
	ResponseType_NotAuthorized    = 12
	ResponseType_NotAvailable     = 13
)

var requestActions = map[int]Action{
	// RequestType_Login: Login,
}

func CreateRequest(client *clients.Connection, bytes []byte) (request *Request, err error) {
	request = &Request{}
	err = json.Unmarshal(bytes, request)

	if err == nil {
		request.Client = client
	}

	return
}

func CreateResponse(request *Request, responseType int, responseData interface{}) *Response {
	response := &Response{
		Id:   request.Id,
		Type: responseType,
	}

	if responseData != nil {
		bytes, _ := json.Marshal(responseData)
		response.Data = bytes
	}

	return response
}

func CreateResponseForAccount(request *Request, responseType int, responseData interface{}) *Response {
	response := CreateResponse(request, responseType, responseData)
	response.Id = 0
	return response
}

func Execute(request *Request) *Response {
	if request.Type == RequestType_Ping {
		return CreateResponse(request, ResponseType_Pong, nil)
	}

	action, found := requestActions[request.Type]

	if !found {
		return notAcceptable(request)
	}

	return action(request)
}

func decodeRequestData(request *Request, data interface{}) bool {
	if request.Data != nil {
		return json.Unmarshal(request.Data, data) == nil
	}

	return false
}

func notAcceptable(request *Request) *Response {
	return CreateResponse(request, ResponseType_NotAcceptable, nil)
}

func notAuthenticated(request *Request) *Response {
	return CreateResponse(request, ResponseType_NotAuthenticated, nil)
}

func notAuthorized(request *Request) *Response {
	return CreateResponse(request, ResponseType_NotAuthorized, nil)
}

func notAvailable(request *Request) *Response {
	return CreateResponse(request, ResponseType_NotAvailable, nil)
}
