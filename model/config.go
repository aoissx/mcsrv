package model

type RestartConfig struct {
	Enable bool     `toml:"enable"`
	Time   []string `toml:"time"`
}

type ServerConfig struct {
	Name      string        `toml:"name"`
	Version   string        `toml:"version"`
	Restart   RestartConfig `toml:"restart"`
	MaxMemory int           `toml:"max_memory"`
	MinMemory int           `toml:"min_memory"`
}

type ConfigModel struct {
	Server ServerConfig `toml:"server"`
}
