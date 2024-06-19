package database

import (
	"athena/server/env"
	"database/sql"
	"time"

	"github.com/tursodatabase/go-libsql"
)

var databaseConnector *libsql.Connector
var databaseInstance *sql.DB

func Connect(development bool) (err error) {
	if development {
		databaseInstance, err = sql.Open("libsql", env.GetDatabaseFilePath())
		return err
	}

	tursoFilePath := env.GetTursoFilePath()
	tursoToken := env.GetTursoToken()
	tursoURL := env.GetTursoURL()

	databaseConnector, err = libsql.NewEmbeddedReplicaConnector(tursoFilePath, tursoURL,
		libsql.WithAuthToken(tursoToken),
		libsql.WithSyncInterval(time.Hour),
	)

	if err != nil {
		return err
	}

	databaseInstance = sql.OpenDB(databaseConnector)
	return
}

func Disconnect() {
	if databaseInstance != nil {
		databaseInstance.Close()
	}

	if databaseConnector != nil {
		databaseConnector.Close()
	}
}

func GetDatabaseInstance() *sql.DB {
	return databaseInstance
}

// Uses sql.Exec() but will panic in the event of a critical error.
func Exec(statement string, args ...any) (insertedId int64, updatedRows int64) {
	data := GetDatabaseInstance()
	result, err := data.Exec(statement, args...)

	if err != nil {
		panic(err)
	}

	if value, err := result.LastInsertId(); err == nil {
		insertedId = value
	}

	if value, err := result.RowsAffected(); err == nil {
		updatedRows = value
	}

	return
}

// Uses sql.Query() but will panic in the event of a critical error.
func Query(statement string, args ...any) (rows *sql.Rows) {
	data := GetDatabaseInstance()
	rows, err := data.Query(statement, args...)

	if err != nil {
		panic(err)
	}

	return
}

type RowCallback[T any] func(row *T, fields RowFieldsReceiver)

type RowFieldsReceiver func(fields ...any)

type RowScanner[T any] struct {
	rows *sql.Rows
}

func (o *RowScanner[T]) Every(callback RowCallback[T]) []*T {
	list := make([]*T, 0, 20)

	for o.rows.Next() {
		var row T

		callback(&row, func(fields ...any) {
			err := o.rows.Scan(fields...)

			if err != nil {
				panic(err)
			}

			list = append(list, &row)
		})
	}

	o.rows.Close()

	return list
}

func (o *RowScanner[T]) First(callback RowCallback[T]) *T {
	var row T
	var ref *T = nil

	for o.rows.Next() {
		callback(&row, func(fields ...any) {
			err := o.rows.Scan(fields...)

			if err != nil {
				panic(err)
			}
		})

		ref = &row
		break
	}

	o.rows.Close()

	return ref
}

func Rows[T any](statement string, args ...any) *RowScanner[T] {
	return &RowScanner[T]{
		rows: Query(statement, args...),
	}
}
