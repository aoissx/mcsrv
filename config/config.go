package config

import (
	_ "embed"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/aoissx/mcsrv/model"
)

//go:embed mcsrv.toml
var defaultConfig []byte

const (
	// ConfigFile is config file name.
	ConfigFile = "mcsrv.toml"
)

func GetConfig() (model.ConfigModel, error) {
	config := model.ConfigModel{}
	_, err := toml.DecodeFile(ConfigFile, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func SaveConfig(config model.ConfigModel) error {
	file, err := os.Create(ConfigFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file).Encode(config)
	if encoder != nil {
		return encoder
	}

	return nil
}

func SaveDefaultConfig() error {
	file, err := os.Create(ConfigFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(defaultConfig)
	if err != nil {
		return err
	}

	return nil
}

func CheckConfigFile() bool {
	_, err := os.Stat(ConfigFile)
	return err == nil
}

// show config data
func ShowConfig() {
	config, err := GetConfig()
	if err != nil {
		panic(err)
	}

	LogInfo("Show config data.")
	LogInfo("Server name: " + config.Server.Name)
	LogInfo("Server version: " + config.Server.Version)
}
