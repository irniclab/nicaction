package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	// فیلدهای مورد نیاز برای تنظیمات
}

func LoadConfig(configFilePath string) (*Config, error) {
	config := &Config{}

	// خواندن فایل تنظیمات و تبدیل آن به یک شیء Config
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(configFile, config); err != nil {
		return nil, err
	}

	return config, nil
}
