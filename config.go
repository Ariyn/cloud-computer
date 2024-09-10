package cloud_computer

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Redis struct {
		Host string
	}
}

var configPath string

func init() {
	values := GetArgument("config")
	if len(values) != 0 {
		configPath = values[0]
	} else {
		configPath = "examples/config.json"
	}
}

func ReadConfig() (c Config, err error) {
	f, err := os.OpenFile(configPath, os.O_RDONLY, 0644)
	if err != nil {
		return
	}

	configData, err := io.ReadAll(f)
	if err != nil {
		return
	}

	err = json.Unmarshal(configData, &c)
	return
}
