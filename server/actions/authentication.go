package actions

import (
	"athena/server/database/store"
	"athena/server/security"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseData struct {
	NameFirst    string `json:"nameFirst"`
	NameLast     string `json:"nameLast"`
	NameInitials string `json:"nameInitials"`
	RoleName     string `json:"roleName"`
	SessionToken string `json:"sessionToken"`
}

func Login(message *ClientMessage) *ServerMessage {
	var data LoginData

	if !decodeMessageData(message, &data) {
		return notAcceptable(message)
	}

	var user = store.GetUserPasswordHashByContactEmail(data.Email)

	if user == nil {
		return notAuthenticated(message)
	}

	if !security.VerifyPassword(data.Password, user.PasswordHash) {
		return notAuthenticated(message)
	}

	store.SetUserSessionToken(user.Id, security.CreateToken())

	user = store.GetUserById(user.Id)

	message.Client.AccountId = user.AccountId
	message.Client.TeamId = user.TeamId
	message.Client.UserId = user.Id
	message.Client.UserRole = user.RoleValue

	return success(message, &LoginResponseData{
		NameFirst:    user.NameFirst,
		NameLast:     user.NameLast,
		NameInitials: user.NameInitials,
		RoleName:     user.RoleName,
		SessionToken: user.SessionToken,
	})
}

func Logout(message *ClientMessage) *ServerMessage {
	var user = store.GetUserById(message.Client.UserId)

	if user == nil {
		return notAuthenticated(message)
	}

	store.SetUserSessionToken(user.Id, "")

	message.Client.AccountId = 0
	message.Client.TeamId = 0
	message.Client.UserId = 0
	message.Client.UserRole = 0

	return success(message, nil)
}

func ResumeLogin(message *ClientMessage) *ServerMessage {
	var user = store.GetUserBySessionToken(message.Token)

	if user == nil {
		return notAuthenticated(message)
	}

	return success(message, &LoginResponseData{
		NameFirst:    user.NameFirst,
		NameLast:     user.NameLast,
		NameInitials: user.NameInitials,
		RoleName:     user.RoleName,
		SessionToken: user.SessionToken,
	})
}
