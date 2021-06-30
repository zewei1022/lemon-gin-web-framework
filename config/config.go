package config

type Config struct {
	Server   Server
	Database Database
	Redis    Redis
	JWT      JWT
	Zap      Zap
}
