package store

import (
	"athena/server/database"
)

type User struct {
	Id           int64
	AccountId    int64
	TeamId       int64
	NameFirst    string
	NameLast     string
	NameInitials string
	ContactEmail string
	ContactPhone string
	RoleName     string
	RoleValue    int64
	PasswordHash string
	SessionToken string
	CreatedAt    int64
	UpdatedAt    int64
}

func GetUserById(id int64) *User {
	return database.Rows[User](`
		SELECT
			id,
			account_id,
			team_id,
			name_first,
			name_last,
			name_initials,
			contact_email,
			contact_phone,
			role_name,
			role_value,
			password_hash,
			session_token,
			created_at,
			updated_at
		FROM
			users
		WHERE
			id = ?
		LIMIT 1
	`, id).First(func(row *User, fields database.RowFieldsReceiver) {
		fields(
			&row.Id,
			&row.AccountId,
			&row.TeamId,
			&row.NameFirst,
			&row.NameLast,
			&row.NameInitials,
			&row.ContactEmail,
			&row.ContactPhone,
			&row.RoleName,
			&row.RoleValue,
			&row.PasswordHash,
			&row.SessionToken,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
	})
}

func GetUserBySessionToken(sessionToken string) *User {
	return database.Rows[User](`
		SELECT
			id,
			account_id,
			team_id,
			name_first,
			name_last,
			name_initials,
			contact_email,
			contact_phone,
			role_name,
			role_value,
			password_hash,
			session_token,
			created_at,
			updated_at
		FROM
			users
		WHERE
			session_token = ?
		LIMIT 1
	`, sessionToken).First(func(row *User, fields database.RowFieldsReceiver) {
		fields(
			&row.Id,
			&row.AccountId,
			&row.TeamId,
			&row.NameFirst,
			&row.NameLast,
			&row.NameInitials,
			&row.ContactEmail,
			&row.ContactPhone,
			&row.RoleName,
			&row.RoleValue,
			&row.PasswordHash,
			&row.SessionToken,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
	})
}

func GetUserPasswordHashByContactEmail(contactEmail string) *User {
	return database.Rows[User](`
		SELECT
			id,
			password_hash
		FROM
			users
		WHERE
			contact_email = ?
		LIMIT 1
	`, contactEmail).First(func(row *User, fields database.RowFieldsReceiver) {
		fields(
			&row.Id,
			&row.PasswordHash,
		)
	})
}

func SetUserSessionToken(id int64, sessionToken string) {
	database.Exec(`
		UPDATE
			users
		SET
			session_token = ?
		WHERE
			id = ?
		LIMIT 1
	`, sessionToken, id)
}
