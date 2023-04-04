package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/irniclab/nicaction/types"
)

// LoadConfig بارگذاری تنظیمات از فایل
func LoadConfig(path string) (types.Config, error) {
	var conf types.Config

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

// SaveConfig ذخیره تغییرات تنظیمات در فایل
func SaveConfig(path string, conf types.Config) error {
	bytes, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
