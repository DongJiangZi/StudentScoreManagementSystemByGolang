package config

type Config struct {
	DataFilePath string
	LogFilePath  string
}

func Load() *Config {
	return &Config{
		DataFilePath: "data/students.json",
		LogFilePath:  "logs/app.log",
	}
}
