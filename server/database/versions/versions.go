package versions

import (
	"athena/server/database"
	"athena/server/env"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type VersionFile struct {
	path    string
	version byte
}

func (o *VersionFile) loadStatements() []string {
	bytes, err := os.ReadFile(o.path)

	if err != nil {
		panic(err)
	}

	statements := strings.Split(string(bytes), ";")

	for i, statement := range statements {
		statements[i] = strings.TrimSpace(statement)
	}

	return statements
}

func Sync() {
	data := database.GetDatabaseInstance()
	currentVersion := loadCurrentVersion()
	versions := loadVersions()

	for _, file := range versions {
		if file.version <= currentVersion {
			continue
		}

		statements := file.loadStatements()

		transaction, err := data.Begin()

		if err != nil {
			panic(err)
		}

		for _, statement := range statements {
			if len(statement) == 0 {
				continue
			}

			_, err := transaction.Exec(statement)

			if err != nil {
				panic(err)
			}
		}

		if err := transaction.Commit(); err != nil {
			transaction.Rollback()
			panic(err)
		}

		currentVersion = file.version

		saveCurrentVersion(currentVersion)

		log.Println("Updated the database to version", currentVersion)
	}
}

func loadCurrentVersion() byte {
	bytes, err := os.ReadFile(env.GetDatabaseVersionFile())

	if err != nil {
		return 0
	}

	return bytes[0]
}

func saveCurrentVersion(version byte) {
	err := os.WriteFile(env.GetDatabaseVersionFile(), []byte{version}, 0666)

	if err != nil {
		panic(err)
	}
}

func loadVersions() []*VersionFile {
	files, err := os.ReadDir(env.GetDatabaseVersionDirectory())

	if err != nil {
		panic(err)
	}

	list := make([]*VersionFile, 0, 20)

	for _, file := range files {
		name := file.Name()

		if !strings.HasSuffix(name, ".sql") {
			continue
		}

		version, err := strconv.Atoi(name[0:3])

		if err != nil {
			continue
		}

		list = append(list, &VersionFile{
			path:    path.Join(env.GetDatabaseVersionDirectory(), name),
			version: byte(version),
		})
	}

	return list
}
