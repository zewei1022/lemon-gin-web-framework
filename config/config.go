package config

type Config struct {
	Server   Server
	Database Database
	Redis    Redis
	Mongodb  Mongodb
	JWT      JWT
	Zap      Zap
}
