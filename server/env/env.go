package env

import "os"

func GetAiKey() string {
	return os.Getenv("ATHENA_AI_KEY")
}

func GetAiOrganization() string {
	return os.Getenv("ATHENA_AI_ORGANIZATION")
}

func GetAiProject() string {
	return os.Getenv("ATHENA_AI_PROJECT")
}

// Only used in the development environment.
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

// Only used in the production environment.
func GetTursoFilePath() string {
	return os.Getenv("ATHENA_TURSO_FILE")
}

// Only used in the production environment.
func GetTursoToken() string {
	return os.Getenv("ATHENA_TURSO_TOKEN")
}

// Only used in the production environment.
func GetTursoURL() string {
	return os.Getenv("ATHENA_TURSO_URL")
}
