package config

type Config struct {
	RepositoryType string
	DataFilePath   string
	LogFilePath    string
}

func Load() *Config {
	return &Config{
		RepositoryType: "json",
		DataFilePath:   "data/students.json",
		LogFilePath:    "logs/app.log",
	}
}
