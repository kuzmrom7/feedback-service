package config

type Configurations struct {
	Server   *ServerConfigurations
	Database *DatabaseConfigurations
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	DBName     string
	DBPort     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
}
