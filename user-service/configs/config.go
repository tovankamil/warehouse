package configs

import (
	"github.com/spf13/viper"
)

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`
}

type SqlDb struct {
	Host              string `json:"host"`
	Port              string `json:"port"`
	User              string `json:"user"`
	Password          string `json:"password"`
	DBName            string `json:"name"`
	MaxIdleConnection int    `json:"max_idle_connection"`
	MaxOpenConnection int    `json:"max_open_connection"`
}

type Redis struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type RabbitMQ struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Supabase struct {
	Url    string `json:"url"`
	ApiKey string `json:"api_key"`
	Bucket string `json:"bucket"`
}

type Config struct {
	App      App      `json:"app"`
	SqlDb    SqlDb    `json:"sql_db"`
	Redis    Redis    `json:"redis"`
	RabbitMQ RabbitMQ `json:"rabbitmq"`
	Supabase Supabase `json:"supabase"`
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort: viper.GetString("APP_PORT"),
			AppEnv:  viper.GetString("APP_ENV"),
		},
		SqlDb: SqlDb{
			Host:              viper.GetString("DATABASE_HOST"),
			Port:              viper.GetString("DATABASE_PORT"),
			User:              viper.GetString("DATABASE_USER"),
			Password:          viper.GetString("DATABASE_PASSWORD"),
			DBName:            viper.GetString("DATABASE_NAME"),
			MaxIdleConnection: viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
			MaxOpenConnection: viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
		},
		Redis: Redis{
			Host: viper.GetString("REDIS_HOST"),
			Port: viper.GetString("REDIS_PORT"),
		},
		RabbitMQ: RabbitMQ{
			Host:     viper.GetString("RABBITMQ_HOST"),
			Port:     viper.GetString("RABBITMQ_PORT"),
			User:     viper.GetString("RABBITMQ_USER"),
			Password: viper.GetString("RABBITMQ_PASSWORD"),
		},
		Supabase: Supabase{
			Url:    viper.GetString("SUPABASE_URL"),
			ApiKey: viper.GetString("SUPABASE_API_KEY"),
			Bucket: viper.GetString("SUPABASE_BUCKET"),
		},
	}
}
