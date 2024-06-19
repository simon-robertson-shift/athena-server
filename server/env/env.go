package env

import "os"

func GetDatabaseFilePath() string {
	return os.Getenv("ATHENA_DATABASE_FILE_PATH")
}

func GetDatabaseVersionDirectory() string {
	return os.Getenv("ATHENA_DATABASE_VERSION_DIRECTORY")
}

func GetDatabaseVersionFile() string {
	return os.Getenv("ATHENA_DATABASE_VERSION_FILE")
}

func GetSocketServerAddress() string {
	return os.Getenv("ATHENA_SOCKET_SERVER_ADDRESS")
}

func GetSocketServerPath() string {
	return os.Getenv("ATHENA_SOCKET_SERVER_PATH")
}

func GetTursoFilePath() string {
	return os.Getenv("ATHENA_TURSO_FILE")
}

func GetTursoToken() string {
	return os.Getenv("ATHENA_TURSO_TOKEN")
}

func GetTursoURL() string {
	return os.Getenv("ATHENA_TURSO_URL")
}
